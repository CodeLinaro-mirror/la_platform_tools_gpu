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

package transform

import (
	"testing"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/atom/test"
)

func TestSkipDrawCalls(t *testing.T) {
	inputs := list(
		&test.AtomA{ID: 2},
		&test.AtomA{ID: 4, AtomFlags: atom.DrawCall},
		&test.AtomA{ID: 3},
		&test.AtomA{ID: 7, AtomFlags: atom.DrawCall},
		&test.AtomA{ID: 8, AtomFlags: atom.DrawCall},
		&test.AtomA{ID: 6},
		&test.AtomA{ID: 13, AtomFlags: atom.DrawCall},
		&test.AtomA{ID: 18},
		&test.AtomA{ID: 5},
		&test.AtomA{ID: 12},
		&test.AtomA{ID: 10, AtomFlags: atom.DrawCall},
		&test.AtomA{ID: 11, AtomFlags: atom.EndOfFrame},

		&test.AtomA{ID: 25},
		&test.AtomA{ID: 24, AtomFlags: atom.DrawCall},
		&test.AtomA{ID: 23},
		&test.AtomA{ID: 11, AtomFlags: atom.EndOfFrame},

		&test.AtomA{ID: 0},
		&test.AtomA{ID: 1, AtomFlags: atom.DrawCall},
		&test.AtomA{ID: 15},
		&test.AtomA{ID: 14, AtomFlags: atom.DrawCall},
		&test.AtomA{ID: 9},
		&test.AtomA{ID: 16, AtomFlags: atom.DrawCall},
		&test.AtomA{ID: 17, AtomFlags: atom.DrawCall},
		&test.AtomA{ID: 19},
		&test.AtomA{ID: 26, AtomFlags: atom.DrawCall},
		&test.AtomA{ID: 19},
		&test.AtomA{ID: 27, AtomFlags: atom.DrawCall},
		&test.AtomA{ID: 21, AtomFlags: atom.EndOfFrame},
		atomAtomID{&atom.EOS{}, 200},
	)

	transform := &SkipDrawCalls{}
	transform.Draw(7)
	transform.Draw(14)
	transform.Draw(17)
	checkTransform(t, transform, inputs, list(
		&test.AtomA{ID: 2},
		&test.AtomA{ID: 4, AtomFlags: atom.DrawCall},
		&test.AtomA{ID: 3},
		&test.AtomA{ID: 7, AtomFlags: atom.DrawCall}, // <---
		// &test.AtomA{ID: 8, AtomFlags: atom.DrawCall},
		&test.AtomA{ID: 6},
		// &test.AtomA{ID: 13, AtomFlags: atom.DrawCall},
		&test.AtomA{ID: 18},
		&test.AtomA{ID: 5},
		&test.AtomA{ID: 12},
		//&test.AtomA{ID: 10, AtomFlags: atom.DrawCall},
		&test.AtomA{ID: 11, AtomFlags: atom.EndOfFrame},

		&test.AtomA{ID: 25},
		// &test.AtomA{ID: 24, AtomFlags: atom.DrawCall},
		&test.AtomA{ID: 23},
		&test.AtomA{ID: 11, AtomFlags: atom.EndOfFrame},

		&test.AtomA{ID: 0},
		&test.AtomA{ID: 1, AtomFlags: atom.DrawCall},
		&test.AtomA{ID: 15},
		&test.AtomA{ID: 14, AtomFlags: atom.DrawCall}, // <---
		&test.AtomA{ID: 9},
		&test.AtomA{ID: 16, AtomFlags: atom.DrawCall},
		&test.AtomA{ID: 17, AtomFlags: atom.DrawCall}, // <---
		&test.AtomA{ID: 19},
		// &test.AtomA{ID: 26, AtomFlags: atom.DrawCall},
		&test.AtomA{ID: 19},
		// &test.AtomA{ID: 27, AtomFlags: atom.DrawCall},
		&test.AtomA{ID: 21, AtomFlags: atom.EndOfFrame},
		atomAtomID{&atom.EOS{}, 200},
	))
}
