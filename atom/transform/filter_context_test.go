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

func TestContextFilter(t *testing.T) {
	inputs := list(
		testAtom{ID: 10, Context: 1},
		testAtom{ID: 30, Context: 2},
		testAtom{ID: 20, Context: 3},
		testAtom{ID: 50, Context: 1},
		testAtom{ID: 90, Context: 2},
		testAtom{ID: 70, Context: 3},
		testAtom{ID: 80, Context: 2},
		testAtom{ID: 00, Context: 3},
		testAtom{ID: 60, Context: 3},
		testAtom{ID: 40, Context: 3},
		atomAtomID{&atom.EOS{}, 999},
	)

	checkTransform(t, ContextFilter(1), inputs, list(
		testAtom{ID: 10, Context: 1},
		testAtom{ID: 50, Context: 1},
		atomAtomID{&atom.EOS{}, 999},
	))

	checkTransform(t, ContextFilter(2), inputs, list(
		testAtom{ID: 30, Context: 2},
		testAtom{ID: 90, Context: 2},
		testAtom{ID: 80, Context: 2},
		atomAtomID{&atom.EOS{}, 999},
	))

	checkTransform(t, ContextFilter(3), inputs, list(
		testAtom{ID: 20, Context: 3},
		testAtom{ID: 70, Context: 3},
		testAtom{ID: 00, Context: 3},
		testAtom{ID: 60, Context: 3},
		testAtom{ID: 40, Context: 3},
		atomAtomID{&atom.EOS{}, 999},
	))
}
