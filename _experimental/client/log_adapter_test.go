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

package client

import (
	"testing"

	"android.googlesource.com/platform/tools/gpu/log"
	test "github.com/google/gxui/testing"
)

func callOnUI(f func()) bool {
	f()
	return true
}

func checkEntry(t *testing.T, expectedMessage string, expectedSeverity log.Severity, actual log.Entry) {
	test.AssertEquals(t, expectedMessage, actual.Message)
	test.AssertEquals(t, expectedSeverity, actual.Severity)
}

func TestLogAdapterAddNoOverflow(t *testing.T) {
	a := CreateLogAdapter(10, callOnUI)
	l := a.Logger()

	test.AssertEquals(t, 0, a.Count())

	log.Infof(l, "I %v", 0)
	log.Warningf(l, "W %v", 1)
	log.Errorf(l, "E %v", 2)
	l.Flush()

	test.AssertEquals(t, 3, a.Count())

	checkEntry(t, "I 0", log.Info, a.Entry(0))
	test.AssertEquals(t, 0, a.ItemIndex(0))

	checkEntry(t, "W 1", log.Warning, a.Entry(1))
	test.AssertEquals(t, 1, a.ItemIndex(1))

	checkEntry(t, "E 2", log.Error, a.Entry(2))
	test.AssertEquals(t, 2, a.ItemIndex(2))
}

func TestLogAdapterClearThenAddNoOverflow(t *testing.T) {
	a := CreateLogAdapter(10, callOnUI)
	l := a.Logger()

	log.Infof(l, "OLD I")
	log.Warningf(l, "OLD W")
	log.Errorf(l, "OLD E")
	l.Flush()

	a.Clear()

	test.AssertEquals(t, 0, a.Count())

	log.Infof(l, "I %v", 0)
	log.Warningf(l, "W %v", 1)
	log.Errorf(l, "E %v", 2)
	l.Flush()

	checkEntry(t, "I 0", log.Info, a.Entry(0))
	checkEntry(t, "W 1", log.Warning, a.Entry(1))
	checkEntry(t, "E 2", log.Error, a.Entry(2))
}

func TestLogAdapterAddOverflow(t *testing.T) {
	a := CreateLogAdapter(5, callOnUI)
	l := a.Logger()

	log.Infof(l, "I %v", 0)
	log.Warningf(l, "W %v", 1)
	log.Errorf(l, "E %v", 2)
	log.Infof(l, "I %v", 3)
	log.Warningf(l, "W %v", 4)
	log.Errorf(l, "E %v", 5)
	log.Infof(l, "I %v", 6)
	log.Warningf(l, "W %v", 7)
	log.Errorf(l, "E %v", 8)
	l.Flush()

	test.AssertEquals(t, 5, a.Count())

	checkEntry(t, "W 4", log.Warning, a.Entry(0))
	test.AssertEquals(t, 0, a.ItemIndex(4))

	checkEntry(t, "E 5", log.Error, a.Entry(1))
	test.AssertEquals(t, 1, a.ItemIndex(5))

	checkEntry(t, "I 6", log.Info, a.Entry(2))
	test.AssertEquals(t, 2, a.ItemIndex(6))

	checkEntry(t, "W 7", log.Warning, a.Entry(3))
	test.AssertEquals(t, 3, a.ItemIndex(7))

	checkEntry(t, "E 8", log.Error, a.Entry(4))
	test.AssertEquals(t, 4, a.ItemIndex(8))
}

func TestLogAdapterClearThenTestAddOverflow(t *testing.T) {
	a := CreateLogAdapter(5, callOnUI)
	l := a.Logger()

	log.Infof(l, "OLD I")
	log.Warningf(l, "OLD W")
	log.Errorf(l, "OLD E")
	log.Infof(l, "OLD I")
	log.Warningf(l, "OLD W")
	log.Errorf(l, "OLD E")
	l.Flush()

	a.Clear()

	test.AssertEquals(t, 0, a.Count())

	log.Infof(l, "I %v", 0)
	log.Warningf(l, "W %v", 1)
	log.Errorf(l, "E %v", 2)
	log.Infof(l, "I %v", 3)
	log.Warningf(l, "W %v", 4)
	log.Errorf(l, "E %v", 5)
	log.Infof(l, "I %v", 6)
	log.Warningf(l, "W %v", 7)
	log.Errorf(l, "E %v", 8)
	l.Flush()

	test.AssertEquals(t, 5, a.Count())

	checkEntry(t, "W 4", log.Warning, a.Entry(0))
	checkEntry(t, "E 5", log.Error, a.Entry(1))
	checkEntry(t, "I 6", log.Info, a.Entry(2))
	checkEntry(t, "W 7", log.Warning, a.Entry(3))
	checkEntry(t, "E 8", log.Error, a.Entry(4))
}
