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
	"io"
	"net"

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
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/service"
	"android.googlesource.com/platform/tools/gpu/service/path"
)

type rpcServer struct {
	service.Resolver
	ReplayManager *replay.Manager
}

func (s rpcServer) ListenAndServe(addr string, mtu int, logger log.Logger) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	for {
		if conn, err := listener.Accept(); err == nil {
			service.BindServer(conn, conn, mtu, log.Fork(logger), &s)
		} else {
			return err
		}
	}
}

// Compliance with the service.RPC interface.

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
func (s rpcServer) Import(name string, data []uint8, l log.Logger) (service.CaptureId, error) {
	atoms := []atom.Atom{}
	d := cyclic.Decoder(vle.Reader(bytes.NewBuffer(data)))
	for {
		if obj, err := d.Object(); err != nil {
			if err != io.EOF {
				log.Warningf(l, "Decode of capture errored after decoding %d atoms: %v", len(atoms), err)
			}
			break
		} else {
			atoms = append(atoms, obj.(atom.Atom))
		}
	}
	if len(atoms) == 0 {
		return service.CaptureId{}, nil
	}
	id, err := builder.ImportCapture(name, atoms, s.Database, l)
	if err != nil {
		return service.CaptureId{}, err
	}
	return id, nil
}

// GetCaptures returns the full list of capture identifiers avaliable on the server.
func (s rpcServer) GetCaptures(l log.Logger) ([]service.CaptureId, error) {
	return builder.Captures(s.Database, l)
}

// GetDevices returns the full list of replay devices avaliable to the server.
// These include local replay devices and any connected Android devices.
// This list may change over time, as devices are connected and disconnected.
func (s rpcServer) GetDevices(l log.Logger) ([]service.DeviceId, error) {
	devices := s.ReplayManager.Devices()
	ids := make([]service.DeviceId, len(devices))
	for i, d := range devices {
		ids[i] = d.ID()
	}
	return ids, nil
}

// GetMemoryInfo returns the MemoryInfo identifier describing the memory state
// for the given capture and memory range, immediately following the atom after.
func (s rpcServer) GetMemoryInfo(
	after *path.Atom,
	rng memory.Range,
	l log.Logger) (service.MemoryInfoId, error) {

	if err := after.Validate(); err != nil {
		return service.MemoryInfoId{}, err
	}
	id, err := database.Store(&builder.GetMemoryInfo{
		After: after,
		Range: memory.Range{Base: rng.Base, Size: rng.Size},
	}, s.Database, l)
	return service.MemoryInfoId{ID: id}, err
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
	l log.Logger) (service.ImageInfoId, error) {

	if err := device.Validate(); err != nil {
		return service.ImageInfoId{}, err
	}
	if err := after.Validate(); err != nil {
		return service.ImageInfoId{}, err
	}
	id, err := database.Store(&builder.GetFramebufferColor{
		Device:   device,
		After:    after,
		Settings: settings,
	}, s.Database, l)
	return service.ImageInfoId{ID: id}, err
}

// GetFramebufferDepth returns the ImageInfo identifier describing the bound
// depth buffer for the given device, capture and graphics API immediately
// following the atom after.
func (s rpcServer) GetFramebufferDepth(
	device *path.Device,
	after *path.Atom,
	l log.Logger) (service.ImageInfoId, error) {

	if err := device.Validate(); err != nil {
		return service.ImageInfoId{}, err
	}
	if err := after.Validate(); err != nil {
		return service.ImageInfoId{}, err
	}
	id, err := database.Store(&builder.GetFramebufferDepth{
		Device: device,
		After:  after,
	}, s.Database, l)
	return service.ImageInfoId{ID: id}, err
}

// GetTimingInfo performs timings of the given capture on the given device and
// capture, returning an identifier to the results.
// This function is experimental and will change signature.
func (s rpcServer) GetTimingInfo(
	device *path.Device,
	capture *path.Capture,
	flags service.TimingFlags,
	l log.Logger) (service.TimingInfoId, error) {

	if err := device.Validate(); err != nil {
		return service.TimingInfoId{}, err
	}
	if err := capture.Validate(); err != nil {
		return service.TimingInfoId{}, err
	}
	id, err := database.Store(&builder.GetTimingInfo{
		Device:  device,
		Capture: capture,
		Flags:   flags,
	}, s.Database, l)
	return service.TimingInfoId{ID: id}, err
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
	apiID service.ApiId,
	width, height uint32,
	atomIDs []uint64,
	l log.Logger) (service.BinaryId, error) {

	if err := device.Validate(); err != nil {
		return service.BinaryId{}, err
	}
	if err := capture.Validate(); err != nil {
		return service.BinaryId{}, err
	}
	id, err := database.Store(&builder.PrerenderFramebuffers{
		Device:  device,
		Capture: capture,
		API:     apiID,
		Width:   width,
		Height:  height,
		AtomIDs: atomIDs,
	}, s.Database, l)
	return service.BinaryId{ID: id}, err
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

// ResolveBinary resolves and returns the byte array associated with id.
func (s rpcServer) ResolveBinary(id service.BinaryId, l log.Logger) (res []uint8, err error) {
	if out, err := s.Database.Resolve(id.ID, l); err == nil {
		res = out.([]uint8)
	}
	return res, err
}
