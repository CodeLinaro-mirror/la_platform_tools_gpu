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

package replay

import (
	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/service"
)

// WireframeMode is an enumerator of wireframe modes used by QueryColorBuffer.
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

// QueryIssues is the interface implemented by types that can verify the replay
// performs as expected and without errors. The returned chan will receive a
// stream of issues detected while replaying the capture and the chan will close
// after the last issue is sent (if any).
// If the capture includes FramebufferObservation atoms, this also includes
// checking the replayed framebuffer matches (within reasonable error) the
// framebuffer observed at capture time.
type QueryIssues interface {
	QueryIssues(ctx Context, mgr *Manager) <-chan Issue
}

// QueryColorBuffer is the interface implemented by types that can return the
// content of the color buffer at a particular point in a capture.
type QueryColorBuffer interface {
	QueryColorBuffer(ctx Context, mgr *Manager, after atom.ID, width, height uint32, wireframeMode WireframeMode) <-chan Image
}

// QueryDepthBufferer is the interface implemented by types that can return the
// content of the depth buffer at a particular point in a capture.
type QueryDepthBuffer interface {
	QueryDepthBuffer(ctx Context, mgr *Manager, after atom.ID) <-chan Image
}

// QueryCallDurations is the interface implemented by types that can time the
// duration of each call in a capture.
type QueryCallDurations interface {
	QueryCallDurations(ctx Context, mgr *Manager, flags service.TimingFlags) <-chan CallTiming
}

// Issue represents a single replay issue reported by QueryIssues.
type Issue struct {
	Atom  atom.ID // The atom that reported the issue.
	Error error   // The error that occurred generating the timing, if there was one.
}

// CallTiming represents the call timing information for a replay.
type CallTiming struct {
	TimingInfo service.TimingInfo // The timing data.
	Error      error              // The error that occurred generating the timing, if there was one.
}

// Image represents pixel data from an api query.
// The exact format of the data depends on the query that generated it.
type Image struct {
	Data  []byte // The pixel data for the image
	Error error  // The error that occurred generating the image if there was one.
}
