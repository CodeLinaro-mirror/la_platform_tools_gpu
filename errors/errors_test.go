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

package errors

import (
	"strings"
	"testing"
)

func c() error {
	return New("Boo!")
}

func b() error {
	return c()
}

func a() error {
	return b()
}

func TestError(t *testing.T) {
	expected :=
		`Boo!
 • android.googlesource.com/platform/tools/gpu/errors_test.go:20 errors.c
 • android.googlesource.com/platform/tools/gpu/errors_test.go:24 errors.b
 • android.googlesource.com/platform/tools/gpu/errors_test.go:28 errors.a
 • android.googlesource.com/platform/tools/gpu/errors_test.go:40 errors.TestError`
	got := a().Error()
	if strings.HasPrefix(got, expected) {
		t.Errorf("Error() returned unexpected string.\nExpected:\n%v\nGot:\n%v", expected, got)
	}
}
