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

type delegate interface {
	Error(...interface{})
	Log(...interface{})
}

// Testing returns a Logger that writes to t's log methods.
func Testing(t delegate) Logger {
	out := make(chan interface{}, 64)
	go func() {
		for o := range out {
			switch o := o.(type) {
			case Entry:
				switch o.Kind {
				case Error:
					t.Error(o.String())
				case Warning, Info:
					t.Log(o.String())
				}

			case FlushRequest:
				close(o)
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
