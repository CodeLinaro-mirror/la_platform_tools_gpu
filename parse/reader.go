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
	"bytes"
	"unicode"
	"unicode/utf8"
)

// Reader is the interface to an object that converts a rune array into tokens.
type Reader struct {
	runes  []rune // The string being parsed.
	offset int    // The start of the current token.
	cursor int    // The offset of the next unparsed rune.
}

func (r *Reader) setData(data string) {
	r.runes = bytes.Runes([]byte(data))
}

// Token peeks at the current scanned token value. It does not consume anything.
func (r *Reader) Token() Token {
	return Token{Runes: r.runes, Start: r.offset, End: r.cursor}
}

// Consume consumes the current token.
func (r *Reader) Consume() Token {
	tok := r.Token()
	r.offset = r.cursor
	return tok
}

// Advance moves the cursor one rune forward.
func (r *Reader) Advance() {
	if r.cursor < len(r.runes) {
		r.cursor++
	}
}

// Rollback sets the cursor back to the last consume point.
func (r *Reader) Rollback() {
	r.cursor = r.offset
}

// IsEOF returns true when the cursor is at the end of the input.
func (r *Reader) IsEOF() bool {
	return r.cursor >= len(r.runes)
}

// GuessNextToken attempts to do a general purpose consume of a single
// arbitrary token from the stream. It is used by error handlers to indicate
// where the error occurred. It guarantees that if the stream is not finished,
// it will consume at least one character.
func (r *Reader) GuessNextToken() Token {
	if r.cursor == r.offset {
		r.Space()
		switch {
		case r.AlphaNumeric():
		case r.Numeric():
		case r.NotSpace():
		default:
			r.Advance()
		}
	}
	return r.Consume()
}

// Peek returns the next rune without advancing the cursor.
func (r *Reader) Peek() rune {
	if r.cursor >= len(r.runes) {
		return utf8.RuneError
	}
	return r.runes[r.cursor]
}

// Rune advances and returns true if the next rune after the cursor matches value.
func (r *Reader) Rune(value rune) bool {
	if r.cursor >= len(r.runes) || r.runes[r.cursor] != value {
		return false
	}
	r.cursor++
	return true
}

// SeekRune advances the cursor until either the value is found or the end
// of stream is reached.
// It returns true if it found value, false otherwise.
func (r *Reader) SeekRune(value rune) bool {
	for i := r.cursor; i < len(r.runes); i++ {
		if r.runes[i] == value {
			r.cursor = i
			return true
		}
	}
	return false
}

// String checks to see if value occurs at cursor, if it does, it advances the
// cursor past it and returns true.
func (r *Reader) String(value string) bool {
	end := r.cursor + len(value)
	if end > len(r.runes) {
		return false
	}
	for i, v := range value {
		if r.runes[r.cursor+i] != v {
			return false
		}
	}
	r.cursor = end
	return true
}

// Space skips over any whitespace, returning true if it advanced the cursor.
func (r *Reader) Space() bool {
	i := r.cursor
	for ; i < len(r.runes); i++ {
		r := r.runes[i]
		if r == RuneEOL || !unicode.IsSpace(r) {
			break
		}
	}
	if i == r.cursor {
		return false
	}
	r.cursor = i
	return true
}

// Space skips over any non whitespace, returning true if it advanced the cursor.
func (r *Reader) NotSpace() bool {
	i := r.cursor
	for ; i < len(r.runes); i++ {
		if unicode.IsSpace(r.runes[i]) {
			break
		}
	}
	if i == r.cursor {
		return false
	}
	r.cursor = i
	return true
}

// Numeric tries to move past the common number pattern, returning true if
// found and false if not.
func (r *Reader) Numeric() bool {
	i := r.cursor
	if i < len(r.runes) && (r.runes[i] == '+' || r.runes[i] == '-') {
		i++
	}
	if i < len(r.runes) && unicode.IsDigit(r.runes[i]) {
		for i++; i < len(r.runes); i++ {
			r := r.runes[i]
			if r != '-' && r != '+' && r != '.' && !unicode.IsLetter(r) && !unicode.IsDigit(r) {
				break
			}
		}
	}

	if i == r.cursor {
		return false
	}
	r.cursor = i
	return true
}

// AlphaNumeric moves past anything that starts with a letter or underscore,
// and consists of letters, numbers or underscores. It returns true if the
// pattern was matched, false otherwise.
func (r *Reader) AlphaNumeric() bool {
	i := r.cursor
	if i >= len(r.runes) {
		return false
	}
	next := r.runes[r.cursor]
	if next == '_' || unicode.IsLetter(next) {
		for i++; i < len(r.runes); i++ {
			next := r.runes[i]
			if next != '_' && !unicode.IsLetter(next) && !unicode.IsDigit(next) {
				break
			}
		}
	}
	if i == r.cursor {
		return false
	}
	r.cursor = i
	return true
}

// NewReader creates a new reader which reads from the supplied string.
func NewReader(data string) *Reader {
	r := &Reader{}
	r.setData(data)
	return r
}
