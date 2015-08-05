// Copyright (C) 2015 The Android Open Source Project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"sync"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/cyclic"
	"android.googlesource.com/platform/tools/gpu/binary/registry"
	"android.googlesource.com/platform/tools/gpu/binary/schema"
	"android.googlesource.com/platform/tools/gpu/binary/vle"
	"android.googlesource.com/platform/tools/gpu/builder"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi/all"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/service"
	"android.googlesource.com/platform/tools/gpu/service/path"
)

type rpcServer struct {
	database.Database
	ReplayManager *replay.Manager
}

// closer, provide connection closing which knows when the last connection
// has gone away.
type closer struct {
	numConn *sync.WaitGroup
	conn    net.Conn
}

func (c closer) Close() error {
	if c.numConn == nil {
		return fmt.Errorf("Multiple calls to Close")
	}
	c.numConn.Done()
	defer func() {
		c.numConn = nil
		c.conn = nil
	}()
	return c.conn.Close()
}

func (s rpcServer) ListenAndServe(addr string, mtu int, logger log.Logger, shutdownOnDisconnect bool) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Errorf(logger, "Error binding to port: %v: %v", addr, err)
		return err
	}

	shutdown := false
	var newCloser func(conn net.Conn) io.Closer
	if shutdownOnDisconnect {
		var numConn *sync.WaitGroup
		newCloser = func(conn net.Conn) io.Closer {
			if numConn == nil {
				// First call to new Closer.
				numConn = &sync.WaitGroup{}
				numConn.Add(1)

				// Wait for the number of connection to fall to zero and
				// then Close() the listener.
				go func() {
					numConn.Wait()
					shutdown = true
					if err := listener.Close(); err != nil {
						log.Errorf(logger, "Closing listener failed: %v", err)
					}
				}()
			} else {
				numConn.Add(1)
			}
			return closer{numConn: numConn, conn: conn}
		}
	} else {
		newCloser = func(conn net.Conn) io.Closer { return conn }
	}

	for !shutdown {
		if conn, err := listener.Accept(); err == nil {
			service.BindServer(conn, conn, newCloser(conn), mtu, log.Fork(logger), &s)
		} else {
			if shutdown {
				log.Infof(logger, "Shutdown requested")
				return nil
			} else {
				log.Errorf(logger, "Error accepting connection: %v", err)
				return err
			}
		}
	}
	return nil
}

// Compliance with the service.Service interface.

// The GetSchema returns the type and constant schema descriptions for all
// objects used in the api.
// This includes all the types included in or referenced from the atom stream.
func (s rpcServer) GetSchema(l log.Logger) (service.Schema, error) {
	result := service.Schema{}
	result.Classes = make([]*schema.Class, 0, registry.Global.Count())
	all.GraphicsNamespace.Visit(func(c binary.Class) {
		class := schema.Lookup(c.ID())
		if class != nil {
			result.Classes = append(result.Classes, class)
		}
	})
	all.VisitConstantSets(func(c schema.ConstantSet) {
		result.Constants = append(result.Constants, c)
	})
	return result, nil
}

// Import imports capture data emitted by the graphics spy, returning the new
// capture identifier.
func (s rpcServer) ImportCapture(name string, data []uint8, l log.Logger) (*path.Capture, error) {
	list := atom.NewList()
	d := cyclic.Decoder(vle.Reader(bytes.NewBuffer(data)))
	for {
		if obj, err := d.Object(); err != nil {
			if err != io.EOF {
				log.Warningf(l, "Decode of capture errored after decoding %d atoms: %v", len(list.Atoms), err)
			}
			break
		} else {
			list.Atoms = append(list.Atoms, obj.(atom.Atom))
		}
	}
	if len(list.Atoms) == 0 {
		return nil, nil
	}
	return builder.ImportCapture(name, list, s.Database, l)
}

// GetCaptures returns the full list of capture identifiers avaliable on the server.
func (s rpcServer) GetCaptures(l log.Logger) ([]*path.Capture, error) {
	return builder.Captures(s.Database, l)
}

// GetDevices returns the full list of replay devices avaliable to the server.
// These include local replay devices and any connected Android devices.
// This list may change over time, as devices are connected and disconnected.
func (s rpcServer) GetDevices(l log.Logger) ([]*path.Device, error) {
	devices := s.ReplayManager.Devices()
	paths := make([]*path.Device, len(devices))
	for i, d := range devices {
		paths[i] = &path.Device{ID: d.ID()}
	}
	return paths, nil
}

// GetFramebufferColor returns the ImageInfo identifier describing the bound
// color buffer for the given device, capture and graphics API immediately
// following the atom after. The provided RenderSettings structure can be used
// to adjust maximum desired dimensions of the image, as well as applying debug
// visualizations.
func (s rpcServer) GetFramebufferColor(
	device *path.Device,
	after *path.Atom,
	settings service.RenderSettings,
	l log.Logger) (*path.ImageInfo, error) {

	if err := device.Validate(); err != nil {
		return nil, err
	}
	if err := after.Validate(); err != nil {
		return nil, err
	}
	id, err := database.Store(&builder.GetFramebufferColor{
		Device:   device,
		After:    after,
		Settings: settings,
	}, s.Database, l)
	return &path.ImageInfo{ID: id}, err
}

// GetFramebufferDepth returns the ImageInfo identifier describing the bound
// depth buffer for the given device, capture and graphics API immediately
// following the atom after.
func (s rpcServer) GetFramebufferDepth(
	device *path.Device,
	after *path.Atom,
	l log.Logger) (*path.ImageInfo, error) {

	if err := device.Validate(); err != nil {
		return nil, err
	}
	if err := after.Validate(); err != nil {
		return nil, err
	}
	id, err := database.Store(&builder.GetFramebufferDepth{
		Device: device,
		After:  after,
	}, s.Database, l)
	return &path.ImageInfo{ID: id}, err
}

// GetTimingInfo performs timings of the given capture on the given device and
// capture, returning an identifier to the results.
// This function is experimental and will change signature.
func (s rpcServer) GetTimingInfo(
	device *path.Device,
	capture *path.Capture,
	flags service.TimingFlags,
	l log.Logger) (*path.TimingInfo, error) {

	if err := device.Validate(); err != nil {
		return nil, err
	}
	if err := capture.Validate(); err != nil {
		return nil, err
	}
	id, err := database.Store(&builder.GetTimingInfo{
		Device:  device,
		Capture: capture,
		Flags:   flags,
	}, s.Database, l)
	return &path.TimingInfo{ID: id}, err
}

// PrerenderFramebuffers renders the framebuffer contents after each of the
// given atoms of interest in the given capture on the given device, resized to
// fit within the given dimensions while keeping the respective framebuffers
// original aspect ratio. This function doesn't return any data, as it is used
// to pre-populate the cache of framebuffer thumbnails that later get queried by
// the client.
// This function is experimental and may change signature.
func (s rpcServer) PrerenderFramebuffers(
	device *path.Device,
	capture *path.Capture,
	apiID service.ApiID,
	width, height uint32,
	atomIndicies []uint64,
	l log.Logger) error {

	if err := device.Validate(); err != nil {
		return err
	}
	if err := capture.Validate(); err != nil {
		return err
	}
	_, err := database.Build(&builder.PrerenderFramebuffers{
		Device:       device,
		Capture:      capture,
		API:          apiID,
		Width:        width,
		Height:       height,
		AtomIndicies: atomIndicies,
	}, s.Database, l)
	return err
}

// Get resolves and returns the object, value or memory at the path p.
func (s rpcServer) Get(p path.Path, l log.Logger) (interface{}, error) {
	if err := p.Validate(); err != nil {
		return nil, err
	}
	return database.Build(&builder.Get{Path: p}, s.Database, l)
}

// Set creates a copy of the capture referenced by p, but with the object, value
// or memory at p replaced with v. The path returned is identical to p, but with
// the base changed to refer to the new capture.
func (s rpcServer) Set(p path.Path, v interface{}, l log.Logger) (path.Path, error) {
	if err := p.Validate(); err != nil {
		return nil, err
	}
	res, err := database.Build(&builder.Set{Path: p, Value: v}, s.Database, l)
	if err != nil {
		return nil, err
	}
	return res.(path.Path), nil
}

// Follow returns the path to the object that the value at p links to.
// If the value at p does not link to anything then nil is returned.
func (s rpcServer) Follow(p path.Path, l log.Logger) (path.Path, error) {
	if err := p.Validate(); err != nil {
		return nil, err
	}
	res, err := database.Build(&builder.Follow{Path: p}, s.Database, l)
	if err != nil {
		return nil, err
	}
	return res.(path.Path), nil
}
