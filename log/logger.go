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

package log

// Logger is the interface for types that implement a hierarchical message logger.
type Logger interface {
	Interface

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
