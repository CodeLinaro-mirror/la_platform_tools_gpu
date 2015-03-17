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

package memory

import (
	"bytes"
	"testing"

	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/log"
)

type testData []byte

func (d testData) Get(db database.Database, logger log.Logger) ([]byte, error) {
	return []byte(d), nil
}
func (d testData) Size() uint64 {
	return uint64(len(d))
}

func check(t *testing.T, expected, got []byte) {
	if !bytes.Equal(expected, got) {
		t.Errorf("Bytes did not match.\nExpected: %v\nGot:      %v", expected, got)
	}
}

func TestSlice(t *testing.T) {
	data := testData{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, test := range []struct {
		rng      Range
		expected []byte
	}{
		{Range{Base: 0, Size: 10}, []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{Range{Base: 3, Size: 3}, []byte{3, 4, 5}},
		{Range{Base: 6, Size: 3}, []byte{6, 7, 8}},
	} {
		got, err := slice(data, test.rng).Get(nil, nil)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
		check(t, test.expected, got)
	}
}

// Write layout:
//      0    1    2    3    4    5
//   ╔════╤════╤════╤════╤════╗────┐
// 0 ║ 10 │ 11 │ 12 │ 13 │ 14 ║    │
//   ╚════╧════╧════╧════╧════╝────┘
//   ╔════╤════╤════╤════╤════╗────┐
// 1 ║ 20 │ 21 │ 22 │ 23 │ 24 ║    │
//   ╚════╧════╧════╧════╧════╝────┘
//   ╔════╤════╗────┬────┬────┬────┐
// 2 ║ 30 │ 31 ║    │    │    │    │
//   ╚════╧════╝────┴────┴────┴────┘
//   ╔════╤════╤════╤════╤════╤════╗
// 3 ║ 40 │ 41 │ 42 │ 43 │ 44 │ 45 ║
//   ╚════╧════╧════╧════╧════╧════╝
func TestMemoryWriteDirect(t *testing.T) {
	m := Memory{}
	for _, test := range []struct {
		data     testData
		expected []byte
	}{
		{testData{10, 11, 12, 13, 14}, []byte{10, 11, 12, 13, 14, 0}},
		{testData{20, 21, 22, 23, 24}, []byte{20, 21, 22, 23, 24, 0}},
		{testData{30, 31}, []byte{30, 31, 22, 23, 24, 0}},
		{testData{40, 41, 42, 43, 44, 45}, []byte{40, 41, 42, 43, 44, 45}},
	} {
		m.Write(test.data)
		got, err := m.Slice(Range{Base: 0, Size: 6}).Get(nil, nil)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
		check(t, test.expected, got)
	}
}

// Write layout:
//      0    1    2    3    4    5    6    7    8    9   10   11
//   ┌────╔════╤════╤════╗────┬────┬────┬────┬────┬────┬────┬────┐
// 0 │    ║ 10 │ 11 │ 12 ║    │    │    │    │    │    │    │    │
//   └────╚════╧════╧════╝────┴────┴────┴────┴────┴────┴────┴────┘
//   ┌────┬────┬────┬────┬────┬────┬────╔════╤════╤════╤════╗────┐
// 1 │    │    │    │    │    │    │    ║ 20 │ 21 │ 22 │ 23 ║    │
//   └────┴────┴────┴────┴────┴────┴────╚════╧════╧════╧════╝────┘
//   ┌────┬────╔════╤════╗────┬────┬────┬────┬────┬────┬────┬────┐
// 2 │    │    ║ 30 │ 31 ║    │    │    │    │    │    │    │    │
//   └────┴────╚════╧════╝────┴────┴────┴────┴────┴────┴────┴────┘
//   ┌────┬────╔════╤════╤════╗────┬────┬────┬────┬────┬────┬────┐
// 3 │    │    ║ 40 │ 41 │ 42 ║    │    │    │    │    │    │    │
//   └────┴────╚════╧════╧════╝────┴────┴────┴────┴────┴────┴────┘
//   ┌────┬────┬────┬────┬────┬────┬────┬────╔════╗────┬────┬────┐
// 4 │    │    │    │    │    │    │    │    ║ 50 ║    │    │    │
//   └────┴────┴────┴────┴────┴────┴────┴────╚════╝────┴────┴────┘
//
func TestMemoryWriteSliced(t *testing.T) {
	m := Memory{}
	m.Slice(Range{Base: 1, Size: 3}).Write(testData{10, 11, 12})
	m.Slice(Range{Base: 7, Size: 4}).Write(testData{20, 21, 22, 23})
	m.Slice(Range{Base: 2, Size: 2}).Write(testData{30, 31})
	m.Slice(Range{Base: 2, Size: 3}).Write(testData{40, 41, 42})
	m.Slice(Range{Base: 8, Size: 1}).Write(testData{50})

	for _, test := range []struct {
		rng      Range
		expected []byte
	}{
		{Range{Base: 0, Size: 12}, []byte{0, 10, 40, 41, 42, 00, 00, 20, 50, 22, 23, 00}},
		{Range{Base: 1, Size: 10}, []byte{10, 40, 41, 42, 00, 00, 20, 50, 22, 23}},
		{Range{Base: 2, Size: 3}, []byte{40, 41, 42}},
		{Range{Base: 5, Size: 2}, []byte{0, 0}},
	} {
		got, err := m.Slice(test.rng).Get(nil, nil)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
		check(t, test.expected, got)
	}
}
