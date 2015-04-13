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

package cpp

import (
	"testing"

	"android.googlesource.com/platform/tools/gpu/build"
)

func eq(a, b build.FileSet) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestParseDeps(t *testing.T) {
	for i, test := range []struct {
		depfile  string
		expected build.FileSet
		fail     bool
	}{
		{
			depfile:  `/dev/gpu/pkg/osx-x64-release/foo/bar.o: /dev/gpu/src/foo/A.cpp`,
			expected: build.FileSet{"/dev/gpu/src/foo/A.cpp"},
		}, {
			depfile: `/dev/gpu/pkg/osx-x64-release/foo/bar.o: \
  /dev/gpu/src/foo/A.cpp \
  /dev/gpu/src/foo/B.h /dev/gpu/src/foo/C.h \
  /dev/gpu/src/foo/D.h
`,
			expected: build.FileSet{
				"/dev/gpu/src/foo/A.cpp", "/dev/gpu/src/foo/B.h",
				"/dev/gpu/src/foo/C.h", "/dev/gpu/src/foo/D.h",
			},
		}, {
			depfile: `c:\dev\gpu\pkg\foo\bar.o: \
  c:\dev\gpu\pkg\foo\A.cpp \
  c:\dev\gpu\pkg\foo\B.h c:\dev\gpu\pkg\foo\C.h \
  c:\dev\gpu\pkg\foo\D.h
`,
			expected: build.FileSet{
				`c:\dev\gpu\pkg\foo\A.cpp`, `c:\dev\gpu\pkg\foo\B.h`,
				`c:\dev\gpu\pkg\foo\C.h`, `c:\dev\gpu\pkg\foo\D.h`,
			},
		}, {
			depfile: ``,
			fail:    true,
		}, {
			depfile: `/dev/gpu/pkg/osx-x64-release/foo/bar.o /dev/gpu/src/foo/A.cpp`,
			fail:    true,
		}, {
			depfile: `/dev/gpu/pkg/osx-x64-release/foo/bar.o: a : b`,
			fail:    true,
		},
	} {
		deps, err := parseDeps(test.depfile)
		switch {
		case test.fail && err == nil:
			t.Errorf("%d Expected error, got none.", i)

		case !test.fail && err != nil:
			t.Errorf("%d Parse failure: %v", i, err)

		case !eq(deps, test.expected):
			t.Errorf("Dependencies were not as expected.\nExpected: %v\nGot:      %v",
				test.expected, deps)
		}
	}
}
