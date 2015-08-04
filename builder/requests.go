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
	"android.googlesource.com/platform/tools/gpu/image"
	"android.googlesource.com/platform/tools/gpu/service"
	"android.googlesource.com/platform/tools/gpu/service/path"
)

// GetState records the parameters of a service.GetState RPC request.
type GetState struct {
	binary.Generate
	After *path.Atom
}

// GetHierarchy records the parameters of a service.GetHierarchy RPC request.
type GetHierarchy struct {
	binary.Generate
	Capture *path.Capture
}

// GetFramebufferColor records the parameters of a service.GetFramebufferColor RPC request.
type GetFramebufferColor struct {
	binary.Generate
	Device   *path.Device
	After    *path.Atom
	Settings service.RenderSettings
}

// GetFramebufferDepth records the parameters of a service.GetFramebufferDepth RPC request.
type GetFramebufferDepth struct {
	binary.Generate
	Device *path.Device
	After  *path.Atom
}

// GetTimingInfo records the parameters of a service.GetTimingInfo RPC request.
type GetTimingInfo struct {
	binary.Generate
	Device  *path.Device
	Capture *path.Capture
	Flags   service.TimingFlags
}

// PrerenderFramebuffers records the parameters of a service.PrerenderFramebuffers RPC request.
type PrerenderFramebuffers struct {
	binary.Generate
	Device  *path.Device
	Capture *path.Capture
	API     service.ApiID
	AtomIDs []uint64
	Width   uint32
	Height  uint32
}

// RenderFramebufferDepth records the parameters of an internal RenderFramebufferDepth request.
type RenderFramebufferDepth struct {
	binary.Generate
	Device            *path.Device
	After             *path.Atom
	FramebufferWidth  uint32
	FramebufferHeight uint32
}

// RenderFramebufferColor records the parameters of an internal RenderFramebufferColor request.
type RenderFramebufferColor struct {
	binary.Generate
	Device        *path.Device
	After         *path.Atom
	Width         uint32
	Height        uint32
	WireframeMode service.WireframeMode
}

// ConvertImage is a request to decode a compressed texture.
type ConvertImage struct {
	binary.Generate
	Data       binary.ID
	Width      int
	Height     int
	FormatFrom image.Format
	FormatTo   image.Format
}

// BuildReport generates a service.Report for the given capture.
type BuildReport struct {
	binary.Generate
	Capture *path.Capture
}

// Get resolves the object, value or memory at Path.
type Get struct {
	binary.Generate
	Path path.Path
}

// Set creates a copy of the capture referenced by Path, but with the object,
// value or memory at p replaced with v. The path returned is identical to Path,
// but with the base changed to refer to the new capture.
type Set struct {
	binary.Generate
	Path  path.Path
	Value interface{}
}

// Follow returns the path to the object that the value at Path links to.
// If the value at Path does not link to anything then nil is returned.
type Follow struct {
	binary.Generate
	Path path.Path
}

// getCaptureFramebufferDimensions records the parameters of an internal request.
type getCaptureFramebufferDimensions struct {
	binary.Generate
	Capture *path.Capture
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
