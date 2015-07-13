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

import "reflect"

// Equal returns true if the paths a and b are equal, otherwise false.
func Equal(a, b Path) bool {
	if (a == nil) && (b == nil) {
		return true // both nil
	}
	if (a == nil) || (b == nil) {
		return false // one nil
	}
	va, vb := reflect.ValueOf(a), reflect.ValueOf(b)
	if va.Type() != vb.Type() {
		return false // different types
	}
	na, nb := va.IsNil(), vb.IsNil()
	if na && nb {
		return true // both boxed nil
	}
	if na != nb {
		return false // one boxed nil
	}
	return a.Path() == b.Path()
}
