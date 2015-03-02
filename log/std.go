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
	"os"
)

// Std returns a Logger that writes to stdout and stderr.
func Std() Logger {
	out := make(chan interface{}, 64)
	go func() {
		for t := range out {
			switch t := t.(type) {
			case Entry:
				if t.Kind == Error {
					os.Stderr.WriteString(t.String())
				} else {
					os.Stdout.WriteString(t.String())
				}
			case FlushRequest:
				os.Stderr.Sync()
				os.Stdout.Sync()
			}
		}
	}()
	nextUid := uint32(1)
	return &channel{
		uid:     0,
		nextUid: &nextUid,
		scope:   "",
		out:     out,
	}
}
