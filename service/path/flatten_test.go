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

package path

import (
	"testing"

	"android.googlesource.com/platform/tools/gpu/check"
)

func TestFlatten(t *testing.T) {
	a := &Capture{}
	b := a.Atoms()
	c := b.Index(10)
	d := c.StateAfter()
	e := d.Field("x")
	f := e.ArrayIndex(1)
	g := f.MapIndex(20)

	check.SlicesEqual(t, Flatten(f), []Path{a, b, c, d, e, f})    // Even
	check.SlicesEqual(t, Flatten(g), []Path{a, b, c, d, e, f, g}) // Odd
}
