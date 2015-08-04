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

// Package service is the definition of the RPC GPU debugger service exposed by the server.
//
// It is not the actual implementation of the service functionality.
package service

import (
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/schema"
	"android.googlesource.com/platform/tools/gpu/image"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/service/path"
)

// binary: java.source = adt/idea/android/src
// binary: java.package = com.android.tools.idea.editors.gfxtrace.service
// binary: java.indent = "  "
// binary: java.member_prefix = my
// binary: service = Service

type Service interface {
	path.Service
	// The GetSchema returns the type and constant schema descriptions for all
	// objects used in the api.
	// This includes all the types included in or referenced from the atom stream.
	GetSchema(l log.Logger) (Schema, error)

	// ImportCapture imports capture data emitted by the graphics spy, returning
	// the new capture identifier.
	ImportCapture(name string, Data []uint8, l log.Logger) (*path.Capture, error)

	// GetCaptures returns the full list of capture identifiers avaliable on the
	// server.
	GetCaptures(l log.Logger) ([]*path.Capture, error)

	// GetDevices returns the full list of replay devices avaliable to the server.
	// These include local replay devices and any connected Android devices.
	// This list may change over time, as devices are connected and disconnected.
	GetDevices(l log.Logger) ([]*path.Device, error)

	// GetFramebufferColor returns the ImageInfo identifier describing the bound
	// color buffer for the given device, immediately following the atom after.
	// The provided RenderSettings structure can be used to adjust maximum desired
	// dimensions of the image, as well as applying debug visualizations.
	GetFramebufferColor(device *path.Device, after *path.Atom, settings RenderSettings, l log.Logger) (*path.ImageInfo, error)

	// GetFramebufferDepth returns the ImageInfo identifier describing the bound
	// depth buffer for the given device, immediately following the atom after.
	GetFramebufferDepth(device *path.Device, after *path.Atom, l log.Logger) (*path.ImageInfo, error)

	// GetTimingInfo performs timings of the given capture on the given device,
	// returning an identifier to the results.
	// This function is experimental and will change signature.
	GetTimingInfo(device *path.Device, capture *path.Capture, flags TimingFlags, l log.Logger) (*path.TimingInfo, error)

	// PrerenderFramebuffers renders the framebuffer contents after each of the
	// given atoms of interest in the given capture on the given device for the
	// given graphics API, resized to fit within the given dimensions while keeping
	// the respective framebuffers original aspect ratio. This function doesn't
	// return any data, as it is used to pre-populate the cache of framebuffer
	// thumbnails that later get queried by the client. This function is
	// experimental and may change signature.
	PrerenderFramebuffers(device *path.Device, capture *path.Capture, api ApiID, width uint32, height uint32, atomIDs []uint64, l log.Logger) error
}

type ApiID binary.ID
type AtomsID binary.ID

// TimingFlags is a bitfield describing what should be timed.
// This is experimental and will change in the near future.
type TimingFlags int

const (
	TimingCPU         TimingFlags = 0 // Time using CPU timers (default).
	TimingGPU         TimingFlags = 1 // Time using GPU timers
	TimingPerCommand  TimingFlags = 2 // Time individual commands.
	TimingPerDrawCall TimingFlags = 4 // Time each draw call.
	TimingPerFrame    TimingFlags = 8 // Time each frame.
)

// Schema holds all the rtti information for dynamic types in the service.
type Schema struct {
	binary.Generate
	Classes   []*schema.Class      // The set of classes in the schema
	Constants []schema.ConstantSet // All the constants the schema includes
}

// Capture describes single capture file held by the server.
type Capture struct {
	binary.Generate
	Name  string  // Name given to the capture. e.g. "KittyWorld"
	Atoms AtomsID // The path to the stream of atoms in this capture.
	Apis  []ApiID // List of graphics APIs used by this capture.
}

// Report describes all warnings and errors found by a capture.
type Report struct {
	binary.Generate
	Items []ReportItem
}

// ReportItem represents an entry in a report.
type ReportItem struct {
	binary.Generate
	Severity log.Severity // The severity of the report item.
	Message  string       // The message for the item.
	Atom     uint64       // The index of the atom the item refers to.
}

// MemoryInfo describes the state of a range of memory at a specific point in
// the atom stream.
type MemoryInfo struct {
	binary.Generate
	Data     []uint8          // The memory values for the span.
	Reads    memory.RangeList // The Data-relative ranges that were read-from at the specified atom.
	Writes   memory.RangeList // The Data-relative ranges that were written-to at the specified atom.
	Observed memory.RangeList // The Data-relative ranges that have been observed.
}

// ImageInfo describes an image, such as a texture or framebuffer at a specific
// point in the atom stream.
type ImageInfo struct {
	binary.Generate
	Format image.Format // The format of the image.
	Width  uint32       // The width of the image in pixels.
	Height uint32       // The height of the image in pixels.
	Data   *path.Blob   // The pixel data of the image.
}

// TimingInfo holds the results of a resolved GetTimingInfo request.
// This is experimental and will change in the near future.
type TimingInfo struct {
	binary.Generate
	PerCommand  []AtomTimer      // The timing results of each command.
	PerDrawCall []AtomRangeTimer // The timing results of each draw call.
	PerFrame    []AtomRangeTimer // The timing results of each frame.
}

// AtomTimer holds the timing information for a single atom.
// This is experimental and will change in the near future.
type AtomTimer struct {
	binary.Generate
	AtomID      uint64 // The atom that was timed.
	Nanoseconds uint64 // The time taken for that atom.
}

// AtomRangeTimer holds the timing information for a range of contiguous atoms.
// This is experimental and will change in the near future.
type AtomRangeTimer struct {
	binary.Generate
	FromAtomID  uint64 // The first atom in the range that was timed.
	ToAtomID    uint64 // The last atom in the range that was timed.
	Nanoseconds uint64 // The time taken for all atoms in the range.
}

// WireframeMode is an enumerator of wireframe modes that can be used by
// RenderSettings.
type WireframeMode int

const (
	// NoWireframe indicates that nothing should be drawn in wireframe.
	NoWireframe = WireframeMode(iota)

	// WireframeOverlay indicates that the single draw call should be overlayed
	// with the wireframe of the mesh.
	WireframeOverlay

	// AllWireframe indicates that all draw calls should be displayed in wireframe.
	AllWireframe
)

// RenderSettings contains settings and flags to be used in replaying and
// returning a bound render target's color buffer.
type RenderSettings struct {
	binary.Generate
	MaxWidth      uint32        // The desired maximum width of the image. The returned image may be larger than this.
	MaxHeight     uint32        // The desired minimum height of the image. The returned image may be larger than this.
	WireframeMode WireframeMode // The wireframe mode to use when rendering.
}
