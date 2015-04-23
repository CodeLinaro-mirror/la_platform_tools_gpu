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

// QueryColorBuffer is the interface implemented by types that can return the
// content of the color buffer at a particular point in a capture.
type QueryColorBuffer interface {
	QueryColorBuffer(ctx *Context, mgr *Manager, after atom.ID, width, height uint32, wireframe bool) <-chan Image
}

// QueryDepthBufferer is the interface implemented by types that can return the
// content of the depth buffer at a particular point in a capture.
type QueryDepthBuffer interface {
	QueryDepthBuffer(ctx *Context, mgr *Manager, after atom.ID) <-chan Image
}

// QueryCallDurations is the interface implemented by types that can time the
// duration of each call in a capture.
type QueryCallDurations interface {
	QueryCallDurations(ctx *Context, mgr *Manager, mask service.TimingMask) <-chan CallTiming
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
