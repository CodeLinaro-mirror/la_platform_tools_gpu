// Copyright (C) 2014 The Android Open Source Project
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

package resolver

import (
	"fmt"
	"testing"

	"android.googlesource.com/platform/tools/gpu/api/ast"
	"android.googlesource.com/platform/tools/gpu/api/parser"
	"android.googlesource.com/platform/tools/gpu/api/semantic"
)

func TestStaticArrays(t *testing.T) {
	for _, test := range []struct {
		name   string
		source string
		errors []string
	}{
		{
			name: "StaticArray initializer",
			source: `type u32[3] A
cmd void foo() { a := A(1,2,3) }`,
		}, {
			name: "StaticArray initializer. Error: too many values",
			source: `type u32[3] A
cmd void foo() { a := A(1,2,3,4) }`,
			errors: []string{"2:24: expected 3 values, got 4"},
		}, {
			name: "StaticArray initializer. Error: too few values",
			source: `type u32[3] A
cmd void foo() { a := A(1,2) }`,
			errors: []string{"2:24: expected 3 values, got 2"},
		},

		{
			name: "StaticArray index read",
			source: `type u32[3] A
cmd void foo() { a := A(1,2,3) i := a[1] }`,
		}, {
			name: "StaticArray index write",
			source: `type u32[3] A
cmd void foo() { a := A(1,2,3) a[2] = 2 }`,
		}, {
			name: "StaticArray index read. Error: out of bounds",
			source: `type u32[3] A
cmd void foo() { a := A(1,2,3) i := a[3] }`,
			errors: []string{"2:38: array index 3 is out of bounds for u32[3]"},
		},
	} {
		astAPI, errs := parser.Parse(test.source)
		if len(errs) > 0 {
			t.Errorf("Testing '%s' - Unexpected parse errors: %v", test.name, errs)
			continue
		}
		_, errs = Resolve([]*ast.API{astAPI}, map[string]*semantic.API{}, ASTToSemantic{})
		matched := len(errs) == len(test.errors)
		if matched {
			for i, err := range errs {
				if fmt.Sprintf("%s", err) != test.errors[i] {
					matched = false
					break
				}
			}
		}
		if !matched {
			t.Errorf("Testing '%s' - Resolve errors were not as expected:", test.name)
			c := len(errs)
			if c < len(test.errors) {
				c = len(test.errors)
			}
			for i := 0; i < c; i++ {
				expected, got := "<none>", "<none>"
				if i < len(test.errors) {
					expected = test.errors[i]
				}
				if i < len(errs) {
					got = fmt.Sprintf("%s", errs[i])
				}
				if expected != got {
					t.Errorf("(%d) - expected: '%v', got: '%v'", i, expected, got)
				}
			}
		}
	}
}
