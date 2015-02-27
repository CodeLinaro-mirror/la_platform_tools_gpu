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

package gfxapi

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/service"
)

// Image represents pixel data from an api query.
// The exact format of the data depends on the query that generated it.
type Image struct {
	Data  []byte // The pixel data for the image
	Error error  // The error that occurred generating the image if there was one.
}

// CallTiming represents the call timing information for a replay.
type CallTiming struct {
	TimingInfo service.TimingInfo // The timing data.
	Error      error              // The error that occurred generating the timing, if there was one.
}

// API is the common interface to a graphics programming api.
type API interface {
	// Name returns the official name of the api.
	Name() string
	// Schema returns the programmatic description of the api, as used by clients.
	Schema() service.Schema
	// InitialState builds and returns a clean state block for the api.
	InitialState() State
	// StateMutator returns an object that can be used to emulate the state changes caused by api commands.
	StateMutator(State) atom.Writer
	// ColorBuffer is used to request the color buffer at a particular point in a capture.
	ColorBuffer(ctx *replay.Context, mgr *replay.Manager, after atom.ID, width, height uint32, wireframe bool) <-chan Image
	// DepthBuffer is used to request the depth buffer at a particular point in a capture.
	DepthBuffer(ctx *replay.Context, mgr *replay.Manager, after atom.ID) <-chan Image
	// TimeCalls requests the timing information for a capture.
	TimeCalls(ctx *replay.Context, mgr *replay.Manager, mask service.TimingMask) <-chan CallTiming
}

var apis map[string]API = make(map[string]API)

// Register adds an api to the understood set.
// It is illegal to register the same name twice.
func Register(api API) {
	if _, present := apis[api.Name()]; present {
		panic(fmt.Errorf("API name %s registered more than once", api.Name()))
	}
	apis[api.Name()] = api
}

// Find looks up a graphics API by name.
// If the name has not been registered, it returns nil.
func Find(name string) API {
	return apis[name]
}
