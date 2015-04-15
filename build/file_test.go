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

import "testing"

func TestFileContains(t *testing.T) {
	for _, test := range []struct {
		dir      File
		file     File
		expected bool
	}{
		{
			dir:      File("foo").Join("bar"),
			file:     File("foo").Join("bar", "cat"),
			expected: true,
		}, {
			dir:      File("foo").Join("bar"),
			file:     File("foo").Join("bar"),
			expected: false,
		}, {
			dir:      File("foo").Join("bar"),
			file:     File("foo").Join("nom"),
			expected: false,
		}, {
			dir:      File("foo").Join("bar", "cat"),
			file:     File("foo").Join("bar"),
			expected: false,
		},
	} {
		got := test.dir.Contains(test.file)
		if got != test.expected {
			t.Errorf("File('%s').Contains('%s') returned %v, expected %v.",
				test.dir, test.file, got, test.expected)
		}
	}
}
