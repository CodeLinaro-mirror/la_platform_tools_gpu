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

package validate

import (
	"testing"

	"android.googlesource.com/platform/tools/gpu/api/ast"
	"android.googlesource.com/platform/tools/gpu/api/parser"
	"android.googlesource.com/platform/tools/gpu/api/resolver"
	"android.googlesource.com/platform/tools/gpu/api/semantic"
)

func compile(t *testing.T, source string) *semantic.API {
	parsed, errs := parser.Parse(source)
	for _, err := range errs {
		t.Error(err.Error())
	}
	if len(errs) > 0 {
		return nil
	}

	compiled, errs := resolver.Resolve([]*ast.API{parsed}, nil, resolver.ASTToSemantic{})
	for _, err := range errs {
		t.Error(err.Error())
	}
	if len(errs) > 0 {
		return nil
	}

	return compiled
}

type expectedError struct {
	line   int
	column int
	msg    string
}

type test struct {
	source   string
	expected []expectedError
}

func TestSimpleUnreachable(t *testing.T) {
	expect := func(errs ...expectedError) []expectedError {
		return errs
	}
	unreachable := func(l, c int) expectedError {
		return expectedError{l, c, "Unreachable block"}
	}

	for _, test := range []test{
		// Bool tests
		{`s32 x = 0
      cmd void f() {
        if (true) { x = 1 }
      }`, expect()},

		{`s32 x = 0
      cmd void f() {
        if (true) { x = 1 } else {}
      }`, expect()},

		{`s32 x = 0
      cmd void f() {
        if (false) {} else { x = 2 }
      }`, expect()},

		{`s32 x = 0
      cmd void f() {
        a := false
        if (a) { x = 1 } else { x = 2 }
      }`, expect(unreachable(4, 16))},

		{`s32 x = 0
      cmd void f() {
        a := true
        if (a) { x = 1 } else { x = 2 }
      }`, expect(unreachable(4, 31))},

		{`s32 x = 0
      cmd void f() {
        a := true == false
        if (a) { x = 1 } else { x = 2 }
      }`, expect(unreachable(4, 16))},

		{`s32 x = 0
      cmd void f() {
        a := true && false
        if (a) { x = 1 } else { x = 2 }
      }`, expect(unreachable(4, 16))},

		{`s32 x = 0
      cmd void f() {
        a := true || false
        if (a) { x = 1 } else { x = 2 }
      }`, expect(unreachable(4, 31))},

		{`s32 x = 0
      cmd void f(bool a) {
        if (a) { if (!a) { x = 1 } }
        if (!a) { if (a) { x = 1 } }
        if (a) {} else if (a) { x = 1 }
        if (!a) {} else if (!a) { x = 1 }
      }`, expect(
			unreachable(3, 26),
			unreachable(4, 26),
			unreachable(5, 31),
			unreachable(6, 33),
		)},

		{`s32 x = 0
      cmd void f(bool v) {
        assert(v)
        if (v) { x = 1 } else { x = 2 }
      }`, expect(unreachable(4, 31))},

		// Uint tests
		{`s32 x = 0
      cmd void f(u32 a) {
        if (a > 5) { x = 1 } else { x = 2 }
      }`, expect()},

		{`s32 x = 0
      cmd void f() {
        a := 5 as u32
        if (a > 4) { x = 1 } else { x = 2 }
        if (a > 5) { x = 1 } else { x = 2 }
        if (a > 6) { x = 1 } else { x = 2 }
        if (4 > a) { x = 1 } else { x = 2 }
        if (5 > a) { x = 1 } else { x = 2 }
        if (6 > a) { x = 1 } else { x = 2 }
      }`, expect(
			unreachable(4, 35),
			unreachable(5, 20),
			unreachable(6, 20),
			unreachable(7, 20),
			unreachable(8, 20),
			unreachable(9, 35),
		)},

		{`s32 x = 0
      cmd void f() {
        a := 5 as u32
        if (a >= 4) { x = 1 } else { x = 2 }
        if (a >= 5) { x = 1 } else { x = 2 }
        if (a >= 6) { x = 1 } else { x = 2 }
        if (4 >= a) { x = 1 } else { x = 2 }
        if (5 >= a) { x = 1 } else { x = 2 }
        if (6 >= a) { x = 1 } else { x = 2 }
      }`, expect(
			unreachable(4, 36),
			unreachable(5, 36),
			unreachable(6, 21),
			unreachable(7, 21),
			unreachable(8, 36),
			unreachable(9, 36),
		)},

		{`s32 x = 0
      cmd void f() {
        a := 5 as u32
        if (a < 4) { x = 1 } else { x = 2 }
        if (a < 5) { x = 1 } else { x = 2 }
        if (a < 6) { x = 1 } else { x = 2 }
        if (4 < a) { x = 1 } else { x = 2 }
        if (5 < a) { x = 1 } else { x = 2 }
        if (6 < a) { x = 1 } else { x = 2 }
      }`, expect(
			unreachable(4, 20),
			unreachable(5, 20),
			unreachable(6, 35),
			unreachable(7, 35),
			unreachable(8, 20),
			unreachable(9, 20),
		)},

		{`s32 x = 0
      cmd void f() {
        a := 5 as u32
        if (a <= 4) { x = 1 } else { x = 2 }
        if (a <= 5) { x = 1 } else { x = 2 }
        if (a <= 6) { x = 1 } else { x = 2 }
        if (4 <= a) { x = 1 } else { x = 2 }
        if (5 <= a) { x = 1 } else { x = 2 }
        if (6 <= a) { x = 1 } else { x = 2 }
      }`, expect(
			unreachable(4, 21),
			unreachable(5, 36),
			unreachable(6, 36),
			unreachable(7, 36),
			unreachable(8, 36),
			unreachable(9, 21),
		)},

		{`s32 x = 0
      cmd void f() {
        a := 5 as u32
        if (a == 4) { x = 1 } else { x = 2 }
        if (a == 5) { x = 1 } else { x = 2 }
        if (a == 6) { x = 1 } else { x = 2 }
      }`, expect(
			unreachable(4, 21),
			unreachable(5, 36),
			unreachable(6, 21),
		)},
	} {
		if api := compile(t, test.source); api != nil {
			got := noUnreachables(api)
			ok := true
			if len(got) == len(test.expected) {
				for i := range got {
					l, c := got[i].At.Token().Cursor()
					if l != test.expected[i].line ||
						c != test.expected[i].column ||
						got[i].Message != test.expected[i].msg {
						ok = false
					}
				}
			} else {
				ok = false
			}

			if !ok {
				t.Errorf("Errors were not as expected for:\n%s\n\n Got: %+v, Wanted: %+v", test.source, got, test.expected)
			}
		}
	}
}
