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

package maker

import (
	"path/filepath"
	"strings"
)

// Path returns the '/' delimited file-path formed from the joined path
// segments, without correcting non-forward slash delimiters found in any path
// segment.
func Path(path ...string) string {
	return strings.Join(path, "/")
}

// OSPath returns the OS delimited file-path formed from the joined path
// segments.
func OSPath(path ...string) (string, error) {
	ospath := filepath.FromSlash(Path(path...))
	abs, err := filepath.Abs(ospath)
	if err != nil {
		return abs, err
	}
	return filepath.Clean(abs), nil
}

// CommonPath returns the '/' delimited file-path formed from the joined path
// segments, correcting non-forward slash delimiters found in any path segment.
func CommonPath(path ...string) string {
	return filepath.ToSlash(Path(path...))
}

func PathSplit(path ...string) (dir, base string) {
	return filepath.Split(Path(path...))
}
