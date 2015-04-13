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
		&test.AtomA{ID: 20, Context: 1},
		&test.AtomA{ID: 50},
		&test.AtomA{ID: 90},
		&test.AtomA{ID: 70, AtomFlags: 2},
		&test.AtomA{ID: 80},
		&test.AtomA{ID: 00},
		&test.AtomA{ID: 60},
		&test.AtomB{ID: 40},
	)

	transform := &Injector{}
	transform.Inject(30, 20, &test.AtomA{ID: 20, Context: 1})
	transform.Inject(90, 70, &test.AtomA{ID: 70, AtomFlags: 2})
	transform.Inject(90, 80, &test.AtomA{ID: 80})
	transform.Inject(60, 40, &test.AtomB{ID: 40})

	transform.Inject(40, 0, &test.AtomA{Context: 0xdead}) // Should not be injected

	checkTransform(t, transform, inputs, expected)
}
