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

import (
	"fmt"
	"time"
)

const fmtTimestamp = "%.2d:%.2d:%.2d.%.3d"
const fmtString = "#%.4d %v: %s%s"

// Entry is a single Info, Warning or Error message written to the chan passed to Channel.
type Entry struct {
	Kind      Kind      // The Entry kind
	Message   string    // The message text
	Scope     string    // The scope of the message
	Context   uint32    // The context identifier for the message
	Timestamp time.Time // The time at which the message was raised
}

// String returns the string representation of the entry with timestamp.
func (e *Entry) String() string {
	h, m, s := e.Timestamp.Clock()
	return fmt.Sprintf(fmtTimestamp+fmtString+"\n",
		h, m, s, e.Timestamp.Nanosecond()/1000000,
		e.Context, e.Kind, e.Scope, e.Message)
}

// StringNoTimestamp String returns the string representation of the entry without the timestamp.
func (e *Entry) StringNoTimestamp() string {
	return fmt.Sprintf(fmtString+"\n",
		e.Context, e.Kind, e.Scope, e.Message)
}

// Kind defines an entry's message kind - either Info, Warning or Error.
type Kind int

const (
	Info Kind = iota
	Warning
	Error
)
