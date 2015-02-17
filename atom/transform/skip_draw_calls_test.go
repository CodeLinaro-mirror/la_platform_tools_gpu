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
)

func TestSkipDrawCalls(t *testing.T) {
	inputs := list(
		testAtom{ID: 2},
		testAtom{ID: 4, AtomFlags: atom.DrawCall},
		testAtom{ID: 3},
		testAtom{ID: 7, AtomFlags: atom.DrawCall},
		testAtom{ID: 8, AtomFlags: atom.DrawCall},
		testAtom{ID: 6},
		testAtom{ID: 13, AtomFlags: atom.DrawCall},
		testAtom{ID: 18},
		testAtom{ID: 5},
		testAtom{ID: 12},
		testAtom{ID: 10, AtomFlags: atom.DrawCall},
		testAtom{ID: 11, AtomFlags: atom.EndOfFrame},

		testAtom{ID: 25},
		testAtom{ID: 24, AtomFlags: atom.DrawCall},
		testAtom{ID: 23},
		testAtom{ID: 11, AtomFlags: atom.EndOfFrame},

		testAtom{ID: 0},
		testAtom{ID: 1, AtomFlags: atom.DrawCall},
		testAtom{ID: 15},
		testAtom{ID: 14, AtomFlags: atom.DrawCall},
		testAtom{ID: 9},
		testAtom{ID: 16, AtomFlags: atom.DrawCall},
		testAtom{ID: 17, AtomFlags: atom.DrawCall},
		testAtom{ID: 19},
		testAtom{ID: 26, AtomFlags: atom.DrawCall},
		testAtom{ID: 19},
		testAtom{ID: 27, AtomFlags: atom.DrawCall},
		testAtom{ID: 21, AtomFlags: atom.EndOfFrame},
		atomAtomID{&atom.EOS{}, 200},
	)

	transform := &SkipDrawCalls{}
	transform.Draw(7)
	transform.Draw(14)
	transform.Draw(17)
	checkTransform(t, transform, inputs, list(
		testAtom{ID: 2},
		testAtom{ID: 4, AtomFlags: atom.DrawCall},
		testAtom{ID: 3},
		testAtom{ID: 7, AtomFlags: atom.DrawCall}, // <---
		// testAtom{ID: 8, AtomFlags: atom.DrawCall},
		testAtom{ID: 6},
		// testAtom{ID: 13, AtomFlags: atom.DrawCall},
		testAtom{ID: 18},
		testAtom{ID: 5},
		testAtom{ID: 12},
		//testAtom{ID: 10, AtomFlags: atom.DrawCall},
		testAtom{ID: 11, AtomFlags: atom.EndOfFrame},

		testAtom{ID: 25},
		// testAtom{ID: 24, AtomFlags: atom.DrawCall},
		testAtom{ID: 23},
		testAtom{ID: 11, AtomFlags: atom.EndOfFrame},

		testAtom{ID: 0},
		testAtom{ID: 1, AtomFlags: atom.DrawCall},
		testAtom{ID: 15},
		testAtom{ID: 14, AtomFlags: atom.DrawCall}, // <---
		testAtom{ID: 9},
		testAtom{ID: 16, AtomFlags: atom.DrawCall},
		testAtom{ID: 17, AtomFlags: atom.DrawCall}, // <---
		testAtom{ID: 19},
		// testAtom{ID: 26, AtomFlags: atom.DrawCall},
		testAtom{ID: 19},
		// testAtom{ID: 27, AtomFlags: atom.DrawCall},
		testAtom{ID: 21, AtomFlags: atom.EndOfFrame},
		atomAtomID{&atom.EOS{}, 200},
	))
}
