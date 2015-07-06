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

package check

import (
	"reflect"
	"testing"
)

// SlicesEqual checks the array or slice of got matches expected.
// If the arrays or slices are equal then true is returned.
// If any differences are found then these are logged to t, the test fails and
// false is returned.
func SlicesEqual(t *testing.T, got interface{}, expected interface{}) (equal bool) {
	return slicesEqual(t, got, expected, func(a, b interface{}) bool { return a == b })
}

// SlicesDeepEqual checks the array or slice of got matches expected using a
// deep-equal comparison.
// If the arrays or slices are equal then true is returned.
// If any differences are found then these are logged to t, the test fails and
// false is returned.
func SlicesDeepEqual(t *testing.T, got interface{}, expected interface{}) (equal bool) {
	return slicesEqual(t, got, expected, reflect.DeepEqual)
}

func slicesEqual(t *testing.T, got interface{}, expected interface{}, same func(a, b interface{}) bool) (equal bool) {
	vg, ve := reflect.ValueOf(got), reflect.ValueOf(expected)
	cg, ce := vg.Len(), ve.Len()
	equal = cg == ce
	if equal {
		for i := 0; i < cg; i++ {
			g, e := vg.Index(i).Interface(), ve.Index(i).Interface()
			if !same(g, e) {
				equal = false
				break
			}
		}
	}

	if !equal {
		for i := 0; i < cg || i < ce; i++ {
			var g, e interface{}
			g, e = "<missing>", "<missing>"
			if i < cg {
				g = vg.Index(i).Interface()
			}
			if i < ce {
				e = ve.Index(i).Interface()
			}
			if same(g, e) {
				t.Logf("  %d: %T %+v", i, g, g)
			} else {
				t.Logf("* %d: %T %+v ---  EXPECTED: %T %+v", i, g, g, e, e)
			}
		}

		t.Fail()
	}

	return equal
}
