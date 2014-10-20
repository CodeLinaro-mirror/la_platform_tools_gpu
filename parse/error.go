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
	"errors"
	"fmt"
	"runtime"
)

var (
	// ParseErrorLimit is the maximum number of errors before a parse is aborted.
	ParseErrorLimit = 10
	// AbortParse is paniced when a parse cannot continue. It is recovered at the
	// top level, to allow the errors to be cleanly returned to the caller.
	AbortParse = errors.New("abort")
)

// Error represents the information that us useful in debugging a parse failure.
type Error struct {
	// At is the parse fragment that was being processed when the error was encountered.
	At Fragment
	// Message is the message associated with the error.
	Message string
	// Stack is the captured stack trace at the point the error was noticed.
	Stack []byte
}

// ErrorList is a convenience type for managing lists of errors.
type ErrorList []Error

func (err Error) Error() string {
	return err.Message
}

func (err Error) Format(f fmt.State, c rune) {
	line, column := err.At.Token().Cursor()
	fmt.Fprintf(f, "%v:%v: %s", line, column, err.Message)
}

func (l *ErrorList) Add(r *Reader, at Fragment, message string, args ...interface{}) {
	if len(*l) >= ParseErrorLimit {
		panic(AbortParse)
	}
	err := Error{At: at}
	if at == nil || at.Token().Len() == 0 {
		invalid := &fragment{}
		if r != nil {
			invalid.SetToken(r.GuessNextToken())
			err.At = invalid
		}
	}
	if len(args) > 0 {
		err.Message = fmt.Sprintf(message, args...)
	} else {
		err.Message = message
	}
	var stack [1 << 16]byte
	size := runtime.Stack(stack[:], false)
	err.Stack = make([]byte, size)
	copy(err.Stack, stack[:size])
	*l = append(*l, err)
}
