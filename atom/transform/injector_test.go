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

func TestInjector(t *testing.T) {
	inputs := list(
		&test.AtomA{ID: 10},
		&test.AtomA{ID: 30},
		&test.AtomA{ID: 50},
		&test.AtomA{ID: 90},
		&test.AtomA{ID: 00},
		&test.AtomA{ID: 60},
	)
	expected := list(
		&test.AtomA{ID: 10},
		&test.AtomA{ID: 30},
		&test.AtomA{ID: atom.NoID, AtomFlags: 1},
		&test.AtomA{ID: 50},
		&test.AtomA{ID: 90},
		&test.AtomA{ID: atom.NoID, AtomFlags: 2},
		&test.AtomA{ID: atom.NoID, AtomFlags: 3},
		&test.AtomA{ID: 00},
		&test.AtomA{ID: 60},
		&test.AtomB{ID: atom.NoID},
	)

	transform := &Injector{}
	transform.Inject(30, &test.AtomA{ID: atom.NoID, AtomFlags: 1})
	transform.Inject(90, &test.AtomA{ID: atom.NoID, AtomFlags: 2})
	transform.Inject(90, &test.AtomA{ID: atom.NoID, AtomFlags: 3})
	transform.Inject(60, &test.AtomB{ID: atom.NoID})

	transform.Inject(40, &test.AtomA{ID: 100, AtomFlags: 5}) // Should not be injected

	checkTransform(t, transform, inputs, expected)
}
