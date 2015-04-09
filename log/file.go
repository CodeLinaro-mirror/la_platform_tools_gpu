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
	"path/filepath"
)

// File creates a new Logger that will write messages to the specified file path.
// If a file exists at the specified path, then this file will be overwritten.
func File(path string) (Logger, error) {
	os.MkdirAll(filepath.Dir(path), 0755)
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	out := make(chan interface{}, 64)
	go func() {
		defer file.Close()
		for t := range out {
			switch t := t.(type) {
			case Entry:
				file.WriteString(t.String())

			case FlushRequest:
				file.Sync()
				close(t)
			}
		}
	}()
	nextUid := uint32(1)
	return &channel{
		uid:     0,
		nextUid: &nextUid,
		scope:   "",
		out:     out,
	}, nil
}
