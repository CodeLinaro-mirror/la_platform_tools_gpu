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

func checkEntry(t *testing.T, expectedMessage string, expectedKind log.Kind, actual log.Entry) {
	test.AssertEquals(t, expectedMessage, actual.Message)
	test.AssertEquals(t, expectedKind, actual.Kind)
}

func TestLogAdapterAddNoOverflow(t *testing.T) {
	l := CreateLogAdapter(10, callOnUI)
	test.AssertEquals(t, 0, l.Count())

	log.Infof(l.Logger(), "I %v", 0)
	log.Warningf(l.Logger(), "W %v", 1)
	l.Logger().Errorf("E %v", 2)
	l.Logger().Flush()

	test.AssertEquals(t, 3, l.Count())

	checkEntry(t, "I 0", log.Info, l.Entry(0))
	test.AssertEquals(t, 0, l.ItemIndex(0))

	checkEntry(t, "W 1", log.Warning, l.Entry(1))
	test.AssertEquals(t, 1, l.ItemIndex(1))

	checkEntry(t, "E 2", log.Error, l.Entry(2))
	test.AssertEquals(t, 2, l.ItemIndex(2))
}

func TestLogAdapterClearThenAddNoOverflow(t *testing.T) {
	l := CreateLogAdapter(10, callOnUI)
	log.Infof(l.Logger(), "OLD I")
	log.Warningf(l.Logger(), "OLD W")
	l.Logger().Errorf("OLD E")
	l.Logger().Flush()

	l.Clear()

	test.AssertEquals(t, 0, l.Count())

	log.Infof(l.Logger(), "I %v", 0)
	log.Warningf(l.Logger(), "W %v", 1)
	l.Logger().Errorf("E %v", 2)
	l.Logger().Flush()

	checkEntry(t, "I 0", log.Info, l.Entry(0))
	checkEntry(t, "W 1", log.Warning, l.Entry(1))
	checkEntry(t, "E 2", log.Error, l.Entry(2))
}

func TestLogAdapterAddOverflow(t *testing.T) {
	l := CreateLogAdapter(5, callOnUI)
	log.Infof(l.Logger(), "I %v", 0)
	log.Warningf(l.Logger(), "W %v", 1)
	l.Logger().Errorf("E %v", 2)
	log.Infof(l.Logger(), "I %v", 3)
	log.Warningf(l.Logger(), "W %v", 4)
	l.Logger().Errorf("E %v", 5)
	log.Infof(l.Logger(), "I %v", 6)
	log.Warningf(l.Logger(), "W %v", 7)
	l.Logger().Errorf("E %v", 8)
	l.Logger().Flush()

	test.AssertEquals(t, 5, l.Count())

	checkEntry(t, "W 4", log.Warning, l.Entry(0))
	test.AssertEquals(t, 0, l.ItemIndex(4))

	checkEntry(t, "E 5", log.Error, l.Entry(1))
	test.AssertEquals(t, 1, l.ItemIndex(5))

	checkEntry(t, "I 6", log.Info, l.Entry(2))
	test.AssertEquals(t, 2, l.ItemIndex(6))

	checkEntry(t, "W 7", log.Warning, l.Entry(3))
	test.AssertEquals(t, 3, l.ItemIndex(7))

	checkEntry(t, "E 8", log.Error, l.Entry(4))
	test.AssertEquals(t, 4, l.ItemIndex(8))
}

func TestLogAdapterClearThenTestAddOverflow(t *testing.T) {
	l := CreateLogAdapter(5, callOnUI)

	log.Infof(l.Logger(), "OLD I")
	log.Warningf(l.Logger(), "OLD W")
	l.Logger().Errorf("OLD E")
	log.Infof(l.Logger(), "OLD I")
	log.Warningf(l.Logger(), "OLD W")
	l.Logger().Errorf("OLD E")
	l.Logger().Flush()

	l.Clear()

	test.AssertEquals(t, 0, l.Count())

	log.Infof(l.Logger(), "I %v", 0)
	log.Warningf(l.Logger(), "W %v", 1)
	l.Logger().Errorf("E %v", 2)
	log.Infof(l.Logger(), "I %v", 3)
	log.Warningf(l.Logger(), "W %v", 4)
	l.Logger().Errorf("E %v", 5)
	log.Infof(l.Logger(), "I %v", 6)
	log.Warningf(l.Logger(), "W %v", 7)
	l.Logger().Errorf("E %v", 8)
	l.Logger().Flush()

	test.AssertEquals(t, 5, l.Count())

	checkEntry(t, "W 4", log.Warning, l.Entry(0))
	checkEntry(t, "E 5", log.Error, l.Entry(1))
	checkEntry(t, "I 6", log.Info, l.Entry(2))
	checkEntry(t, "W 7", log.Warning, l.Entry(3))
	checkEntry(t, "E 8", log.Error, l.Entry(4))
}
