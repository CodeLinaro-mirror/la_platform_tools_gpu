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

func TestInjector(t *testing.T) {
	inputs := list(
		testAtom{ID: 10},
		testAtom{ID: 30},
		testAtom{ID: 50},
		testAtom{ID: 90},
		testAtom{ID: 00},
		testAtom{ID: 60},
		atomAtomID{&atom.EOS{}, 0},
	)
	expected := list(
		testAtom{ID: 10},
		testAtom{ID: 30},
		testAtom{ID: 20, Context: 1},
		testAtom{ID: 50},
		testAtom{ID: 90},
		testAtom{ID: 70, AtomFlags: 2},
		testAtom{ID: 80},
		testAtom{ID: 00},
		testAtom{ID: 60},
		testAtom{ID: 40, Type: 3},
		atomAtomID{&atom.EOS{}, 0},
	)

	transform := &Injector{}
	transform.Inject(30, 20, testAtom{ID: 20, Context: 1})
	transform.Inject(90, 70, testAtom{ID: 70, AtomFlags: 2})
	transform.Inject(90, 80, testAtom{ID: 80})
	transform.Inject(60, 40, testAtom{ID: 40, Type: 3})

	transform.Inject(40, 0, testAtom{Context: 0xdead}) // Should not be injected

	checkTransform(t, transform, inputs, expected)
}
