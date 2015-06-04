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

package asm

import "testing"

// Check is a test helper function that checks the list of got instructions
// matches those in expected. If any differences are found then these are logged
// to t, and the test fails.
func Check(t *testing.T, got []Instruction, expected []Instruction) (matched bool) {
	matched = len(got) == len(expected)
	if matched {
		for i := range expected {
			if got[i] != expected[i] {
				matched = false
				break
			}
		}
	}
	if !matched {
		t.Errorf("Instructions were not as expected:")
		c := len(got)
		if c < len(expected) {
			c = len(expected)
		}
		for i := 0; i < c; i++ {
			g, e := Instruction(nil), Instruction(nil)
			if i < len(got) {
				g = got[i]
			}
			if i < len(expected) {
				e = expected[i]
			}
			if g == e {
				t.Errorf("  %d: %T%+v", i, g, g)
			} else {
				t.Errorf("* %d: %T%+v ---  EXPECTED: %T%+v", i, g, g, e, e)
			}
		}
	}
	return
}
