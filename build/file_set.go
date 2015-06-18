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

package build

// FileSet is a list of Files with no duplicates.
type FileSet []File

// Files returns a FileSet formed from the list of files.
func Files(files ...File) FileSet {
	return FileSet(files)
}

// Append returns a new FileSet with files appended to fs. Duplicates will not
// be included in the returned FileSet.
func (fs FileSet) Append(files ...File) FileSet {
	m := fs.toMap()
	out := append(FileSet{}, fs...)
	for _, file := range files {
		if _, found := m[file]; !found {
			out = append(out, file)
		}
	}
	return out
}

// Filter returns the list of files in this FileSet that matches any pattern in
// patterns.
func (fs FileSet) Filter(patterns ...string) FileSet {
	out := make(FileSet, 0, len(fs))
	for _, file := range fs {
		if file.Matches(patterns...) {
			out = append(out, file)
		}
	}
	return out
}

// Exclude returns the list of files in this FileSet that does not match any
// pattern in patterns.
func (fs FileSet) Exclude(patterns ...string) FileSet {
	out := make(FileSet, 0, len(fs))
nextfile:
	for _, file := range fs {
		if file.Matches(patterns...) {
			continue nextfile
		}
		out = append(out, file)
	}
	return out
}

func (fs FileSet) toMap() map[File]struct{} {
	m := make(map[File]struct{}, len(fs))
	for _, file := range fs {
		m[file] = struct{}{}
	}
	return m
}
