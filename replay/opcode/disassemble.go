/*
 * Copyright 2015, The Android Open Source Project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package opcode

import (
	"io"
	"testing"

	"android.googlesource.com/platform/tools/gpu/binary"
)

// Disassemble disassembles and returns the stream of encoded Opcodes from r,
// stopping once an EOF is reached.
func Disassemble(r io.Reader) ([]interface{}, error) {
	d := binary.NewDecoder(r)
	opcodes := []interface{}{}
	for {
		opcode, err := Decode(d)
		switch err {
		case nil:
			opcodes = append(opcodes, opcode)
		case io.EOF:
			return opcodes, nil
		default:
			return nil, err
		}
	}
}

// CheckDisassembly is a test helper function that checks the list of got
// opcodes matches those in expected. If any differences are found then these
// are logged to t, and the test fails.
func CheckDisassembly(t *testing.T, got []interface{}, expected ...interface{}) {
	matched := len(got) == len(expected)
	if matched {
		for i := range got {
			if got[i] != expected[i] {
				matched = false
				break
			}
		}
	}

	if !matched {
		for i := 0; i < len(got) || i < len(expected); i++ {
			var e, g interface{}
			if i < len(expected) {
				e = expected[i]
			}
			if i < len(got) {
				g = got[i]
			}

			if e == g {
				t.Logf("  %d: %T%+v", i, g, g)
			} else {
				t.Logf("* %d: %T%+v ---  EXPECTED: %T%+v", i, g, g, e, e)
			}
		}

		t.Fail()
	}
}
