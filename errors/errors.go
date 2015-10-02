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

// Package errors provides helper methods for constructing errors that contain
// the stack trace.
package errors

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

// Error holds an error message and callstack at the point of generation.
type Error struct {
	Message   string
	Callstack []uintptr
}

// New creates and returns new Error.
func New(msg string) Error {
	var buf [128]uintptr
	pc := buf[:runtime.Callers(2, buf[:])]
	return Error{Message: msg, Callstack: pc}
}

// Newf creates and returns new Error with the printf-format string.
func Newf(msg string, args ...interface{}) Error {
	var buf [128]uintptr
	pc := buf[:runtime.Callers(2, buf[:])]
	return Error{Message: fmt.Sprintf(msg, args...), Callstack: pc}
}

// Error implements the error interface.
func (e Error) Error() string {
	lines := make([]string, len(e.Callstack))
	for i, c := range e.Callstack {
		fun := runtime.FuncForPC(c)
		abs, line := fun.FileLine(c)
		pkg, funcName := path.Split(fun.Name())
		fileName := path.Base(abs)
		lines[i] = fmt.Sprintf(" • %v%v:%v %v", pkg, fileName, line, funcName)
	}
	return fmt.Sprintf("%v\n%v", e.Message, strings.Join(lines, "\n"))
}
