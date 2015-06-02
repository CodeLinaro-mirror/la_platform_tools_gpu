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

// Package replay is used to issue replay requests to replay devices.
package replay

import (
	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service"
)

// Generator is the interface for types that support replay generation.
type Generator interface {
	// ReplayTransforms is called when a replay pass is ready to be sent to the
	// replay device. ReplayTransforms returns an atom transform list that
	// transforms the original, unaltered atom stream into a stream configured for
	// the replay pass. The transforms should satisfy all the specified requests
	// and config.
	ReplayTransforms(
		ctx Context,
		cfg Config,
		requests []Request,
		device *service.Device,
		db database.Database,
		logger log.Logger) atom.Transforms
}

// Context describes the source capture and replay target information used for
// issuing a replay request.
type Context struct {
	DeviceID  service.DeviceId  // The identifier of the device being used for replay.
	CaptureID service.CaptureId // The identifier of the capture that is being replayed.
}

// Config is a user-defined type used to describe the type of replay being
// requested. Replay requests made with configs that have equality (==) will
// likely be batched into the same replay pass. Configs can be used to force
// requests into different replay passes. For example, by issuing requests with
// different configs we can prevent a profiling Request from being issued in the
// same pass as a Request to render all draw calls in wireframe.
type Config interface{}

// Request is a user-defined type that holds information relevant to a single
// replay request. An example Request would be one that informs ReplayTransforms
// to insert a postback of the currently bound render-target content at a
// specific atom.
type Request interface{}
