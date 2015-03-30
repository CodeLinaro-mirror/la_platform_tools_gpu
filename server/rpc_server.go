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
	Database      database.Database
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

// Import imports capture data emitted by the graphics spy, returning the new
// capture identifier.
func (s rpcServer) Import(logger log.Logger, name string, data service.U8Array) (service.CaptureId, error) {
	atoms := atom.List{}
	if err := atoms.Decode(cyclic.Decoder(vle.Reader(bytes.NewBuffer(data)))); err != nil {
		return service.CaptureId{}, err
	}
	id, err := builder.ImportCapture(name, atoms, s.Database, logger)
	if err != nil {
		return service.CaptureId{}, err
	}
	return id, nil
}

// GetCaptures returns the full list of capture identifiers avaliable on the server.
func (s rpcServer) GetCaptures(logger log.Logger) (service.CaptureIdArray, error) {
	return builder.Captures(s.Database, logger)
}

// GetDevices returns the full list of replay devices avaliable to the server.
// These include local replay devices and any connected Android devices.
// This list may change over time, as devices are connected and disconnected.
func (s rpcServer) GetDevices(logger log.Logger) (service.DeviceIdArray, error) {
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
func (s rpcServer) GetState(logger log.Logger, captureID service.CaptureId, at uint64) (service.BinaryId, error) {
	id, err := s.Database.StoreRequest(&builder.GetState{
		Capture: captureID,
		After:   atom.ID(at),
	}, logger)
	return service.BinaryId{ID: id}, err
}

// GetHierarchy returns the atom hierarchy identifier for the given capture and context.
// Currently there is only one hierarchy per capture context, but this is likely to change in the future.
func (s rpcServer) GetHierarchy(logger log.Logger, captureID service.CaptureId, ctxID uint32) (service.HierarchyId, error) {
	id, err := s.Database.StoreRequest(&builder.GetHierarchy{
		Capture: captureID,
		Context: atom.ContextID(ctxID),
	}, logger)
	return service.HierarchyId{ID: id}, err
}

// GetMemoryInfo returns the MemoryInfo identifier describing the memory state
// for the given capture, context and range, immediately following the atom after.
func (s rpcServer) GetMemoryInfo(logger log.Logger, captureID service.CaptureId, after uint64, rng service.MemoryRange) (service.MemoryInfoId, error) {
	id, err := s.Database.StoreRequest(&builder.GetMemoryInfo{
		Capture: captureID,
		After:   atom.ID(after),
		Range:   memory.Range{Base: memory.Pointer(rng.Base), Size: rng.Size},
	}, logger)
	return service.MemoryInfoId{ID: id}, err
}

// GetFramebufferColor returns the ImageInfo identifier describing the bound color buffer for the given device,
// capture and context immediately following the atom after. The provided RenderSettings structure can be used
// to adjust maximum desired dimensions of the image, as well as applying debug visualizations.
func (s rpcServer) GetFramebufferColor(logger log.Logger, deviceID service.DeviceId, captureID service.CaptureId, ctxID uint32, after uint64, settings service.RenderSettings) (service.ImageInfoId, error) {
	id, err := s.Database.StoreRequest(&builder.GetFramebufferColor{
		Device:   deviceID,
		Capture:  captureID,
		Context:  atom.ContextID(ctxID),
		After:    atom.ID(after),
		Settings: settings,
	}, logger)
	return service.ImageInfoId{ID: id}, err
}

// GetFramebufferDepth returns the ImageInfo identifier describing the bound depth buffer for the given device,
// capture and context immediately following the atom after.
func (s rpcServer) GetFramebufferDepth(logger log.Logger, deviceID service.DeviceId, captureID service.CaptureId, ctxID uint32, after uint64) (service.ImageInfoId, error) {
	id, err := s.Database.StoreRequest(&builder.GetFramebufferDepth{
		Device:  deviceID,
		Capture: captureID,
		Context: atom.ContextID(ctxID),
		After:   atom.ID(after),
	}, logger)
	return service.ImageInfoId{ID: id}, err
}

// ReplaceAtom creates and new capture based on an existing capture, but with a single atom replaced.
func (s rpcServer) ReplaceAtom(logger log.Logger, capture service.CaptureId, atomID uint64, atomType uint16, data service.Binary) (service.CaptureId, error) {
	id, err := s.Database.StoreRequest(&builder.ReplaceAtom{
		Capture: capture,
		Atom:    atom.ID(atomID),
		Type:    atom.TypeID(atomType),
		Data:    data,
	}, logger)
	return service.CaptureId{ID: id}, err
}

// GetTimingInfo performs timings of the given capture, context on the given device,
// returning an identifier to the results.
// This function is experimental and will change signature.
func (s rpcServer) GetTimingInfo(logger log.Logger, deviceID service.DeviceId, captureID service.CaptureId, ctxID uint32, mask service.TimingMask) (service.TimingInfoId, error) {
	id, err := s.Database.StoreRequest(&builder.GetTimingInfo{
		Device:     deviceID,
		Capture:    captureID,
		Context:    atom.ContextID(ctxID),
		TimingMask: mask,
	}, logger)
	return service.TimingInfoId{ID: id}, err
}

// PrerenderFramebuffers renders the framebuffer contents after each of the given atoms of interest
// in the given capture on the given device, resized to fit within the given dimensions while keeping
// the respective framebuffers original aspect ratio. This function doesn't return any data, as it is
// used to pre-populate the cache of framebuffer thumbnails that later get queried by the client.
// This function is experimental and may change signature.
func (s rpcServer) PrerenderFramebuffers(logger log.Logger, deviceID service.DeviceId, captureID service.CaptureId, width, height uint32, atomIDs service.U64Array) (service.BinaryId, error) {
	id, err := s.Database.StoreRequest(&builder.PrerenderFramebuffers{
		Device:  deviceID,
		Capture: captureID,
		Width:   width,
		Height:  height,
		AtomIDs: atomIDs,
	}, logger)
	return service.BinaryId{ID: id}, err
}

// ResolveAtomStream resolves the given id to a cached AtomStream or builds it on demand before returning it.
// ResolveAtomStream and the following two-step resolution methods provide a way for the client to first check its local cache for a
// potential match, before making a resolution request to the server, that could incur a significant bandwidth and computational cost.
func (s rpcServer) ResolveAtomStream(logger log.Logger, id service.AtomStreamId) (r service.AtomStream, e error) {
	e = s.Database.Load(id.ID, logger, &r)
	return r, e
}

// ResolveBinary resolves the given id to a cached Binary or builds it on demand before returning it.
func (s rpcServer) ResolveBinary(logger log.Logger, id service.BinaryId) (r service.Binary, e error) {
	e = s.Database.Load(id.ID, logger, &r)
	return r, e
}

// ResolveCapture resolves the given id to a cached Capture or builds it on demand before returning it.
func (s rpcServer) ResolveCapture(logger log.Logger, id service.CaptureId) (r service.Capture, e error) {
	e = s.Database.Load(id.ID, logger, &r)
	return r, e
}

// ResolveDevice resolves the given id to a cached Device or builds it on demand before returning it.
func (s rpcServer) ResolveDevice(logger log.Logger, id service.DeviceId) (r service.Device, e error) {
	e = s.Database.Load(id.ID, logger, &r)
	return r, e
}

// ResolveHierarchy resolves the given id to a cached Hierarchy or builds it on demand before returning it.
func (s rpcServer) ResolveHierarchy(logger log.Logger, id service.HierarchyId) (r service.Hierarchy, e error) {
	e = s.Database.Load(id.ID, logger, &r)
	return r, e
}

// ResolveImageInfo resolves the given id to a cached ImageInfo or builds it on demand before returning it.
func (s rpcServer) ResolveImageInfo(logger log.Logger, id service.ImageInfoId) (r service.ImageInfo, e error) {
	e = s.Database.Load(id.ID, logger, &r)
	return r, e
}

// ResolveMemoryInfo resolves the given id to a cached MemoryInfo or builds it on demand before returning it.
func (s rpcServer) ResolveMemoryInfo(logger log.Logger, id service.MemoryInfoId) (r service.MemoryInfo, e error) {
	e = s.Database.Load(id.ID, logger, &r)
	return r, e
}

// ResolveSchema resolves the given id to a cached Schema or builds it on demand before returning it.
func (s rpcServer) ResolveSchema(logger log.Logger, id service.SchemaId) (r service.Schema, e error) {
	e = s.Database.Load(id.ID, logger, &r)
	return r, e
}

// ResolveTimingInfo resolves the given id to a cached TimingInfo or builds it on demand before returning it.
func (s rpcServer) ResolveTimingInfo(logger log.Logger, id service.TimingInfoId) (r service.TimingInfo, e error) {
	e = s.Database.Load(id.ID, logger, &r)
	return r, e
}
