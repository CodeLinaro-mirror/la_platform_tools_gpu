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
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/replay/builder"
)

// Replayer is the interface that wraps the basic Replay method.
type Replayer interface {
	// Replay issues replay operations to the replay builder b for the given atom
	// with identifier id, and graphics API state s. If the replay action will
	// have an effect on the graphics driver state, then the call to Replay should
	// also apply the corresponding changes to the state s. If postback is true
	// then the replay instructions should include postback of all outputs of the
	// action.
	Replay(id atom.ID, s *gfxapi.State, b *builder.Builder, postback bool)
}

// Replay issues replay operations to the replay builder b for the given atom a
// with identifier id, and graphics API state s. If replaying the Atom will have
// an effect on the graphics driver state, then the call to Replay will also
// apply the corresponding changes to the state s. If postback is true then
// the replay instructions should include postback of all outputs of the Atom.
func Replay(id atom.ID, a atom.Atom, s *gfxapi.State, b *builder.Builder, postback bool) {
	switch a := a.(type) {
	case *atom.Observation:
		b.Observation(a.Range, a.ResourceID)
	case Replayer:
		a.Replay(id, s, b, postback)
	default:
		a.Mutate(s)
	}
}
