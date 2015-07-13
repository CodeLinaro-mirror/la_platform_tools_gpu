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

// Nop is an implementation of Logger interface that does nothing for Info, Warning, Error and Flush.
// Enter and Fork both return Nops.
type Nop struct{}

// Log does nothing
func (Nop) log(s Severity, msg string, args ...interface{}) {}

// enter returns the same Nop implementation of Logger
func (Nop) enter(name string) Logger { return Nop{} }

// fork returns the same Nop implementation of Logger
func (Nop) fork() Logger { return Nop{} }

// close does nothing
func (Nop) close() {}
