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
	"io"
	"os"
	"path/filepath"
)

type syncer interface {
	Sync() error
}

func doSync(w io.Writer) {
	if w != nil {
		if s, ok := w.(syncer); ok {
			s.Sync()
		}
	}
}

// Writer returns a Logger that writes to the supplied streams.
func Writer(info, warn, err io.Writer, done func()) Logger {
	out := make(chan interface{}, 64)
	go func() {
		if done != nil {
			defer done()
		}
		for t := range out {
			switch t := t.(type) {
			case Entry:
				switch {
				case t.Severity <= Error:
					io.WriteString(err, t.String())
				case t.Severity <= Warning:
					io.WriteString(warn, t.String())
				default:
					io.WriteString(info, t.String())
				}
			case FlushRequest:
				doSync(info)
				doSync(warn)
				doSync(err)
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
	}
}

// File creates a new Logger that will write messages to the specified file path.
// If a file exists at the specified path, then this file will be overwritten.
func File(path string) (Logger, error) {
	os.MkdirAll(filepath.Dir(path), 0755)
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return Writer(file, file, file, func() { file.Close() }), nil
}

// Std returns a Logger that writes to stdout and stderr.
func Std() Logger {
	return Writer(os.Stdout, os.Stdout, os.Stderr, nil)
}
