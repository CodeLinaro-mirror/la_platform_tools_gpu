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

package service

import (
	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/schema"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/service/path"
)

type RPC interface {
	// The GetSchema returns the type and constant schema descriptions for all
	// objects used in the api.
	// This includes all the types included in or referenced from the atom stream.
	GetSchema(l log.Logger) (Schema, error)

	// Import imports capture data emitted by the graphics spy, returning the new
	// capture identifier.
	Import(name string, Data []uint8, l log.Logger) (CaptureId, error)

	// GetCaptures returns the full list of capture identifiers avaliable on the
	// server.
	GetCaptures(l log.Logger) ([]CaptureId, error)

	// GetDevices returns the full list of replay devices avaliable to the server.
	// These include local replay devices and any connected Android devices.
	// This list may change over time, as devices are connected and disconnected.
	GetDevices(l log.Logger) ([]DeviceId, error)

	// GetMemoryInfo returns the MemoryInfo identifier describing the memory state
	// for the given capture and range, immediately following the atom
	// after.
	GetMemoryInfo(after *path.Atom, rng memory.Range, l log.Logger) (MemoryInfoId, error)

	// GetFramebufferColor returns the ImageInfo identifier describing the bound
	// color buffer for the given device, immediately following the atom after.
	// The provided RenderSettings structure can be used to adjust maximum desired
	// dimensions of the image, as well as applying debug visualizations.
	GetFramebufferColor(device *path.Device, after *path.Atom, settings RenderSettings, l log.Logger) (ImageInfoId, error)

	// GetFramebufferDepth returns the ImageInfo identifier describing the bound
	// depth buffer for the given device, immediately following the atom after.
	GetFramebufferDepth(device *path.Device, after *path.Atom, l log.Logger) (ImageInfoId, error)

	// GetTimingInfo performs timings of the given capture on the given device,
	// returning an identifier to the results.
	// This function is experimental and will change signature.
	GetTimingInfo(device *path.Device, capture *path.Capture, flags TimingFlags, l log.Logger) (TimingInfoId, error)

	// PrerenderFramebuffers renders the framebuffer contents after each of the
	// given atoms of interest in the given capture on the given device for the
	// given graphics API, resized to fit within the given dimensions while keeping
	// the respective framebuffers original aspect ratio. This function doesn't
	// return any data, as it is used to pre-populate the cache of framebuffer
	// thumbnails that later get queried by the client. This function is
	// experimental and may change signature.
	PrerenderFramebuffers(device *path.Device, capture *path.Capture, api ApiId, width uint32, height uint32, atomIds []uint64, l log.Logger) (BinaryId, error)

	// Get resolves and returns the object, value or memory at the path p.
	Get(p path.Path, l log.Logger) (interface{}, error)

	// Set creates a copy of the capture referenced by p, but with the object, value
	// or memory at p replaced with v. The path returned is identical to p, but with
	// the base changed to refer to the new capture.
	Set(p path.Path, v interface{}, l log.Logger) (path.Path, error)

	ResolveAtomStream(id AtomStreamId, l log.Logger) (AtomStream, error)
	ResolveBinary(id BinaryId, l log.Logger) ([]uint8, error)
	ResolveCapture(id CaptureId, l log.Logger) (Capture, error)
	ResolveDevice(id DeviceId, l log.Logger) (Device, error)
	ResolveImageInfo(id ImageInfoId, l log.Logger) (ImageInfo, error)
	ResolveMemoryInfo(id MemoryInfoId, l log.Logger) (MemoryInfo, error)
	ResolveTimingInfo(id TimingInfoId, l log.Logger) (TimingInfo, error)
}

// Handle ApiId
type ApiId struct {
	binary.Generate
	ID binary.ID
}

// Handle AtomStreamId
type AtomStreamId struct {
	binary.Generate
	ID binary.ID
}

// Handle BinaryId
type BinaryId struct {
	binary.Generate
	ID binary.ID
}

// Handle CaptureId
type CaptureId struct {
	binary.Generate
	ID binary.ID
}

// Handle DeviceId
type DeviceId struct {
	binary.Generate
	ID binary.ID
}

// Handle ImageInfoId
type ImageInfoId struct {
	binary.Generate
	ID binary.ID
}

// Handle MemoryInfoId
type MemoryInfoId struct {
	binary.Generate
	ID binary.ID
}

// Handle TimingInfoId
type TimingInfoId struct {
	binary.Generate
	ID binary.ID
}

// Enum Severity
type Severity int

const (
	SeverityEmergency     Severity = 0 // Errors that indicate a failure of the server, no further data should be trusted.
	SeverityAlert         Severity = 1 // Errors that indicate a failure of the server, possibly recoverable.
	SeverityCritical      Severity = 2 // Errors so severe that no further analysis or replay can be performed beyond this point.
	SeverityError         Severity = 3 // Errors describing problems that should be fixed. Likely to result in undefined behavior.
	SeverityWarning       Severity = 4 // Warnings describe issues that might affect performance or compatibility, but could be ignored.
	SeverityNotice        Severity = 5 // Normal but significant condition message.
	SeverityInformational Severity = 6 // Informational messages, safe to ignore.
	SeverityDebug         Severity = 7 // Verbose, debug-level messages.
)

// The enumerator of image formats. Will expand.
type ImageFormat int

const (
	ImageFormatRGBA8   ImageFormat = 0
	ImageFormatFloat32 ImageFormat = 1
)

// TimingFlags is a bitfield describing what should be timed.
// This is experimental and will change in the near future.
type TimingFlags int

const (
	TimingFlagsTimingCPU         TimingFlags = 0 // Time using CPU timers (default).
	TimingFlagsTimingGPU         TimingFlags = 1 // Time using GPU timers
	TimingFlagsTimingPerCommand  TimingFlags = 2 // Time individual commands.
	TimingFlagsTimingPerDrawCall TimingFlags = 4 // Time each draw call.
	TimingFlagsTimingPerFrame    TimingFlags = 8 // Time each frame.
)

// Schema holds all the rtti information for dynamic types in the service.
type Schema struct {
	binary.Generate
	Classes   []*schema.Class      // The set of classes in the schema
	Constants []schema.ConstantSet // All the constants the schema includes
}

// Device describes replay target avaliable to the server.
type Device struct {
	binary.Generate
	Name             string // The name of the device. e.g. "Bob's phone"
	Model            string // The model of the device. e.g. "Nexus 5"
	OS               string // The operating system of the device. e.g. "Android 5.0"
	PointerSize      uint8  // Size in bytes of a pointer on the device's architecture.
	PointerAlignment uint8  // Alignment in bytes of a pointer on the device's architecture.
	MaxMemorySize    uint64 // The total amount of contiguous memory pre-allocated for replay.
	Extensions       string // Renderer extensions list. e.g. "GL_KHR_debug GL_EXT_sRGB [...]".
	Renderer         string // Renderer name. e.g. "Adreno (TM) 320".
	Vendor           string // Renderer vendor name. e.g. "Qualcomm".
	Version          string // Renderer version. e.g. "OpenGL ES 3.0 V@53.0 AU@  (CL@)".
}

// Capture describes single capture file held by the server.
type Capture struct {
	binary.Generate
	Name  string       // Name given to the capture. e.g. "KittyWorld"
	Atoms AtomStreamId // The identifier of the stream of atoms in this capture.
	Apis  []ApiId      // List of graphics APIs used by this capture.
}

// Report describes all warnings and errors found by a capture.
type Report struct {
	binary.Generate
	Items []ReportItem
}

// ReportItem represents an entry in a report.
type ReportItem struct {
	binary.Generate
	Severity Severity // The severity of the report item.
	Message  string   // The message for the item.
	Atom     uint64   // The index of the atom the item refers to.
}

// AtomStream holds a stream of atoms.
type AtomStream struct {
	binary.Generate
	Atoms []atom.Atom
}

// Hierarchy holds the root to an AtomGroup hierarchy.
type Hierarchy struct {
	binary.Generate
	Root atom.Group
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
	Format ImageFormat // The format of the image.
	Width  uint32      // The width of the image in pixels.
	Height uint32      // The height of the image in pixels.
	Data   BinaryId    // The pixel data of the image.
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
	AtomId      uint64 // The atom that was timed.
	Nanoseconds uint64 // The time taken for that atom.
}

// AtomRangeTimer holds the timing information for a range of contiguous atoms.
// This is experimental and will change in the near future.
type AtomRangeTimer struct {
	binary.Generate
	FromAtomId  uint64 // The first atom in the range that was timed.
	ToAtomId    uint64 // The last atom in the range that was timed.
	Nanoseconds uint64 // The time taken for all atoms in the range.
}

// RenderSettings contains settings and flags to be used in replaying and
// returning a bound render target's color buffer.
type RenderSettings struct {
	binary.Generate
	MaxWidth  uint32 // The desired maximum width of the image. The returned image may be larger than this.
	MaxHeight uint32 // The desired minimum height of the image. The returned image may be larger than this.
	Wireframe bool   // True if the all geometry should be rendered as wireframe.
}
