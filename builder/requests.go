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

//go:generate codergen -go=requests_binary.go requests.go

package builder

import (
	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/service"
)

// GetState records the parameters of a service.GetState RPC request.
type GetState struct {
	Capture service.CaptureId
	Context atom.ContextId
	After   atom.Id
}

// GetHierarchy records the parameters of a service.GetHierarchy RPC request.
type GetHierarchy struct {
	Capture service.CaptureId
	Context atom.ContextId
}

// GetMemoryInfo records the parameters of a service.GetMemoryInfo RPC request.
type GetMemoryInfo struct {
	Capture service.CaptureId
	Context atom.ContextId
	After   atom.Id
	Range   memory.Range
}

// GetFramebufferColor records the parameters of a service.GetFramebufferColor RPC request.
type GetFramebufferColor struct {
	Capture  service.CaptureId
	Context  atom.ContextId
	Device   service.DeviceId
	After    atom.Id
	Settings service.RenderSettings
}

// GetFramebufferDepth records the parameters of a service.GetFramebufferDepth RPC request.
type GetFramebufferDepth struct {
	Capture service.CaptureId
	Context atom.ContextId
	Device  service.DeviceId
	After   atom.Id
}

// ReplaceAtom records the parameters of a service.ReplaceAtom RPC request.
type ReplaceAtom struct {
	Capture service.CaptureId
	Atom    atom.Id
	Type    atom.TypeId
	Data    service.Binary
}

// GetTimingInfo records the parameters of a service.GetTimingInfo RPC request.
type GetTimingInfo struct {
	Capture    service.CaptureId
	Context    atom.ContextId
	Device     service.DeviceId
	TimingMask service.TimingMask
}

// PrerenderFramebuffers records the parameters of a service.PrerenderFramebuffers RPC request.
type PrerenderFramebuffers struct {
	Device  service.DeviceId
	Capture service.CaptureId
	AtomIds []uint64
	Width   uint32
	Height  uint32
}

// RenderFramebufferDepth records the parameters of an internal RenderFramebufferDepth request.
type RenderFramebufferDepth struct {
	Capture           service.CaptureId
	Context           atom.ContextId
	Device            service.DeviceId
	After             atom.Id
	FramebufferWidth  uint32
	FramebufferHeight uint32
}

// RenderFramebufferColor records the parameters of an internal RenderFramebufferColor request.
type RenderFramebufferColor struct {
	Capture   service.CaptureId
	Context   atom.ContextId
	Device    service.DeviceId
	After     atom.Id
	Width     uint32
	Height    uint32
	Wireframe bool
}

// getCaptureFramebufferDimensions records the parameters of an internal request.
type getCaptureFramebufferDimensions struct {
	Capture service.CaptureId
	Context atom.ContextId
}

// atomFramebufferDimensions records the parameters of an internal resource for getCaptureFramebufferDimensions.
type atomFramebufferDimensions struct {
	From   atom.Id
	Width  uint32
	Height uint32
}

// captureFramebufferDimensions records the parameters of an internal resource for getCaptureFramebufferDimensions.
type captureFramebufferDimensions struct {
	Dimensions []atomFramebufferDimensions
}
