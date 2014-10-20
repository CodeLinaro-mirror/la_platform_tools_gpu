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

package parse

import (
	"fmt"
)

// RuneEOL is the rune that marks the end of a line.
const RuneEOL = '\n'

// A Token represents the smallest consumed unit input.
type Token struct {
	Runes []rune // The full rune array for the string this token is from.
	Start int    // The start of the token in the full rune array.
	End   int    // One past the end of the token.
}

// Format implements fmt.Formatter writing the start end and value of the token.
func (t Token) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, "%d:%d:%s", t.Start, t.End, t.String())
}

// String returns the string form of the rune range the token represents.
func (t Token) String() string {
	if t.Start >= t.End {
		return ""
	}
	return string(t.Runes[t.Start:t.End])
}

// Cursor is used to calculate the line and column of the start of the token.
// It may be very expensive to call, and is intended to be used sparingly in
// producing human readable error messages only.
func (t Token) Cursor() (line int, column int) {
	line = 1
	column = 1
	for _, r := range t.Runes[:t.Start] {
		if r == RuneEOL {
			line++
			column = 0
		}
		column++
	}
	return line, column
}

// Len returns the length of the token in runes.
func (t Token) Len() int {
	return t.End - t.Start
}
