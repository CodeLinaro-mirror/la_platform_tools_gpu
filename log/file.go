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

// File implements the Logger interface, writing all messages out to a text file.
type File struct {
	channel
}

// NewFile creates a new File that will write messages to the specified file path.
// If a file exists at the specified path, then this file will be overwritten.
func NewFile(path string) (*File, error) {
	os.MkdirAll(filepath.Dir(path), 0755)
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	out := make(chan Entry, 64)
	go func() {
		defer file.Close()
		for entry := range out {
			file.WriteString(entry.String())
		}
	}()
	nextUid := uint32(1)
	return &File{
		channel: channel{
			uid:     0,
			nextUid: &nextUid,
			scope:   "",
			out:     out,
		},
	}, nil
}

// Close closes the file. Writing messages to the File after it has been closed may deadlock the
// program.
func (f *File) Close() {
	close(f.out)
}
