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

// Package log provides a hierarchical logger interface and implementations of the interface.
package log

// Logger is the interface for types that implement a hierarchical message logger.
type Logger interface {
	// Info writes an information message to the logger. These message types are intented to be used
	// for non-critial, expected events. Arguments are handled in the manner of fmt.Printf.
	Infof(msg string, args ...interface{})

	// Warning writes a warning message to the logger. This is intented to be used for unexpected but
	// non-critical events. Arguments are handled in the manner of fmt.Printf.
	Warningf(msg string, args ...interface{})

	// Error writes an error message to the logger. This is intented to be used for unexpected and
	// critical error events. Arguments are handled in the manner of fmt.Printf.
	Errorf(msg string, args ...interface{})

	// Enter creates a new logger scoped within the existing logger. This can be used to produce
	// hierarchical log messages.
	Enter(name string) Logger

	// Fork creates a new logger with the same scope as the existing logger, but with a new context
	// identifier. It is good practice to fork logs before passing to another goroutine so that
	// messages can be associated with their goroutine of execution.
	Fork() Logger

	// Flush ensures that any pending messages are written by the logger.
	Flush()

	// Close closes the logger, automatically flushing any remaining messages.
	// After calling Close, no other methods can be called on the logger.
	Close()
}
