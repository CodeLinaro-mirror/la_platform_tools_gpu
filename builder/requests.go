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

package builder

import (
	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/service"
)

// GetState records the parameters of a service.GetState RPC request.
type GetState struct {
	binary.Generate
	Capture service.CaptureId
	After   atom.ID
}

// GetHierarchy records the parameters of a service.GetHierarchy RPC request.
type GetHierarchy struct {
	binary.Generate
	Capture service.CaptureId
	Context atom.ContextID
}

// GetMemoryInfo records the parameters of a service.GetMemoryInfo RPC request.
type GetMemoryInfo struct {
	binary.Generate
	Capture service.CaptureId
	After   atom.ID
	Range   memory.Range
}

// GetFramebufferColor records the parameters of a service.GetFramebufferColor RPC request.
type GetFramebufferColor struct {
	binary.Generate
	Capture  service.CaptureId
	Context  atom.ContextID
	Device   service.DeviceId
	After    atom.ID
	Settings service.RenderSettings
}

// GetFramebufferDepth records the parameters of a service.GetFramebufferDepth RPC request.
type GetFramebufferDepth struct {
	binary.Generate
	Capture service.CaptureId
	Context atom.ContextID
	Device  service.DeviceId
	After   atom.ID
}

// ReplaceAtom records the parameters of a service.ReplaceAtom RPC request.
type ReplaceAtom struct {
	binary.Generate
	Capture service.CaptureId
	Atom    atom.ID
	Type    atom.TypeID
	Data    service.Binary
}

// GetTimingInfo records the parameters of a service.GetTimingInfo RPC request.
type GetTimingInfo struct {
	binary.Generate
	Capture    service.CaptureId
	Context    atom.ContextID
	Device     service.DeviceId
	TimingMask service.TimingMask
}

// PrerenderFramebuffers records the parameters of a service.PrerenderFramebuffers RPC request.
type PrerenderFramebuffers struct {
	binary.Generate
	Device  service.DeviceId
	Capture service.CaptureId
	AtomIDs []uint64
	Width   uint32
	Height  uint32
}

// RenderFramebufferDepth records the parameters of an internal RenderFramebufferDepth request.
type RenderFramebufferDepth struct {
	binary.Generate
	Capture           service.CaptureId
	Context           atom.ContextID
	Device            service.DeviceId
	After             atom.ID
	FramebufferWidth  uint32
	FramebufferHeight uint32
}

// RenderFramebufferColor records the parameters of an internal RenderFramebufferColor request.
type RenderFramebufferColor struct {
	binary.Generate
	Capture   service.CaptureId
	Context   atom.ContextID
	Device    service.DeviceId
	After     atom.ID
	Width     uint32
	Height    uint32
	Wireframe bool
}

// captures records the parameters of an internal request.
type captures struct {
	binary.Generate
	ids service.CaptureIdArray
}

// getCaptureFramebufferDimensions records the parameters of an internal request.
type getCaptureFramebufferDimensions struct {
	binary.Generate
	Capture service.CaptureId
	Context atom.ContextID
}

// atomFramebufferDimensions records the parameters of an internal resource for getCaptureFramebufferDimensions.
type atomFramebufferDimensions struct {
	binary.Generate
	From   atom.ID
	Width  uint32
	Height uint32
}

// captureFramebufferDimensions records the parameters of an internal resource for getCaptureFramebufferDimensions.
type captureFramebufferDimensions struct {
	binary.Generate
	Dimensions []atomFramebufferDimensions
}
