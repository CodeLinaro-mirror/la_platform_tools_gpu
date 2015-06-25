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
	"net"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary/cyclic"
	"android.googlesource.com/platform/tools/gpu/binary/vle"
	"android.googlesource.com/platform/tools/gpu/builder"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/service"
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
			service.BindServer(conn, conn, mtu, logger.Fork(), &s)
		} else {
			return err
		}
	}
}

// Compliance with the service.RPC interface.

// The GetSchema returns the type and constant schema descriptions for all
// objects used in the api.
// This includes all the types included in or referenced from the atom stream.
func (s rpcServer) GetSchema(l log.Logger) (service.ClassPtrArray, error) {
	return nil, nil
}

// Import imports capture data emitted by the graphics spy, returning the new
// capture identifier.
func (s rpcServer) Import(name string, data service.U8Array, l log.Logger) (service.CaptureId, error) {
	atoms := atom.List{}
	if err := atoms.Decode(cyclic.Decoder(vle.Reader(bytes.NewBuffer(data)))); err != nil {
		if len(atoms) == 0 {
			return service.CaptureId{}, err
		}
		l.Warningf("Decode of capture errored after decoding %d atoms: %v", len(atoms), err)
	}
	id, err := builder.ImportCapture(name, atoms, s.Database, l)
	if err != nil {
		return service.CaptureId{}, err
	}
	return id, nil
}

// GetCaptures returns the full list of capture identifiers avaliable on the server.
func (s rpcServer) GetCaptures(l log.Logger) (service.CaptureIdArray, error) {
	return builder.Captures(s.Database, l)
}

// GetDevices returns the full list of replay devices avaliable to the server.
// These include local replay devices and any connected Android devices.
// This list may change over time, as devices are connected and disconnected.
func (s rpcServer) GetDevices(l log.Logger) (service.DeviceIdArray, error) {
	devices := s.ReplayManager.Devices()
	ids := make(service.DeviceIdArray, len(devices))
	for i, d := range devices {
		ids[i] = d.ID()
	}
	return ids, nil
}

// GetState returns an identifier to a binary blob containing the graphics state
// immediately following the atom after.
// The binary blob can be fetched with a call to ResolveBinary, and decoded
// using the capture's schema.
func (s rpcServer) GetState(
	captureID service.CaptureId,
	at uint64,
	l log.Logger) (service.BinaryId, error) {

	id, err := database.Store(&builder.GetState{
		Capture: captureID,
		After:   atom.ID(at),
	}, s.Database, l)
	return service.BinaryId{ID: id}, err
}

// GetHierarchy returns the atom hierarchy identifier for the given capture.
// Currently there is only one hierarchy per capture, but this is likely to
// change in the future.
func (s rpcServer) GetHierarchy(
	captureID service.CaptureId,
	l log.Logger) (service.HierarchyId, error) {

	id, err := database.Store(&builder.GetHierarchy{
		Capture: captureID,
	}, s.Database, l)
	return service.HierarchyId{ID: id}, err
}

// GetMemoryInfo returns the MemoryInfo identifier describing the memory state
// for the given capture and memory range, immediately following the atom after.
func (s rpcServer) GetMemoryInfo(
	captureID service.CaptureId,
	after uint64,
	rng service.MemoryRange,
	l log.Logger) (service.MemoryInfoId, error) {

	id, err := database.Store(&builder.GetMemoryInfo{
		Capture: captureID,
		After:   atom.ID(after),
		Range:   memory.Range{Base: memory.Pointer(rng.Base), Size: rng.Size},
	}, s.Database, l)
	return service.MemoryInfoId{ID: id}, err
}

// GetFramebufferColor returns the ImageInfo identifier describing the bound
// color buffer for the given device, capture and graphics API immediately
// following the atom after. The provided RenderSettings structure can be used
// to adjust maximum desired dimensions of the image, as well as applying debug
// visualizations.
func (s rpcServer) GetFramebufferColor(
	deviceID service.DeviceId,
	captureID service.CaptureId,
	apiID service.ApiId,
	after uint64,
	settings service.RenderSettings,
	l log.Logger) (service.ImageInfoId, error) {

	id, err := database.Store(&builder.GetFramebufferColor{
		Device:   deviceID,
		Capture:  captureID,
		API:      apiID,
		After:    atom.ID(after),
		Settings: settings,
	}, s.Database, l)
	return service.ImageInfoId{ID: id}, err
}

// GetFramebufferDepth returns the ImageInfo identifier describing the bound
// depth buffer for the given device, capture and graphics API immediately
// following the atom after.
func (s rpcServer) GetFramebufferDepth(
	deviceID service.DeviceId,
	captureID service.CaptureId,
	apiID service.ApiId,
	after uint64,
	l log.Logger) (service.ImageInfoId, error) {

	id, err := database.Store(&builder.GetFramebufferDepth{
		Device:  deviceID,
		Capture: captureID,
		API:     apiID,
		After:   atom.ID(after),
	}, s.Database, l)
	return service.ImageInfoId{ID: id}, err
}

// GetTimingInfo performs timings of the given capture on the given device and
// capture, returning an identifier to the results.
// This function is experimental and will change signature.
func (s rpcServer) GetTimingInfo(
	deviceID service.DeviceId,
	captureID service.CaptureId,
	mask service.TimingMask,
	l log.Logger) (service.TimingInfoId, error) {

	id, err := database.Store(&builder.GetTimingInfo{
		Device:     deviceID,
		Capture:    captureID,
		TimingMask: mask,
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
	deviceID service.DeviceId,
	captureID service.CaptureId,
	apiID service.ApiId,
	width, height uint32,
	atomIDs service.U64Array,
	l log.Logger) (service.BinaryId, error) {

	id, err := database.Store(&builder.PrerenderFramebuffers{
		Device:  deviceID,
		Capture: captureID,
		API:     apiID,
		Width:   width,
		Height:  height,
		AtomIDs: atomIDs,
	}, s.Database, l)
	return service.BinaryId{ID: id}, err
}

// ReplaceAtom creates and new capture based on an existing capture, but with
// a single atom replaced.
func (s rpcServer) ReplaceAtom(
	capture service.CaptureId,
	atomID uint64,
	atomType uint16,
	data service.Binary,
	l log.Logger) (service.CaptureId, error) {

	id, err := database.Store(&builder.ReplaceAtom{
		Capture: capture,
		Atom:    atom.ID(atomID),
		Type:    atom.TypeID(atomType),
		Data:    data,
	}, s.Database, l)
	return service.CaptureId{ID: id}, err
}
