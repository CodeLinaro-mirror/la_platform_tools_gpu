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

func check(t *testing.T, expected, got []byte) {
	if !bytes.Equal(expected, got) {
		t.Errorf("Bytes did not match.\nExpected: %v\nGot:      %v", expected, got)
	}
}

func TestBlobSlice(t *testing.T) {
	data := Blob([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	for _, test := range []struct {
		rng      Range
		expected []byte
	}{
		{Range{Base: 0, Size: 10}, []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{Range{Base: 3, Size: 3}, []byte{3, 4, 5}},
		{Range{Base: 6, Size: 3}, []byte{6, 7, 8}},
	} {
		got, err := data.Slice(test.rng).Get(nil, nil)
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
func TestPoolBlobWriteRead(t *testing.T) {
	p := Pool{}
	for _, test := range []struct {
		data     Slice
		expected []byte
	}{
		{Blob([]byte{10, 11, 12, 13, 14}), []byte{10, 11, 12, 13, 14, 0}},
		{Blob([]byte{20, 21, 22, 23, 24}), []byte{20, 21, 22, 23, 24, 0}},
		{Blob([]byte{30, 31}), []byte{30, 31, 22, 23, 24, 0}},
		{Blob([]byte{40, 41, 42, 43, 44, 45}), []byte{40, 41, 42, 43, 44, 45}},
	} {
		p.Write(0, test.data)
		got, err := p.Slice(Range{Base: 0, Size: 6}).Get(nil, nil)
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
func TestMemoryBlobWriteReadScattered(t *testing.T) {
	p := Pool{}
	p.Write(1, Blob([]byte{10, 11, 12}))
	p.Write(7, Blob([]byte{20, 21, 22, 23}))
	p.Write(2, Blob([]byte{30, 31}))
	p.Write(2, Blob([]byte{40, 41, 42}))
	p.Write(8, Blob([]byte{50}))

	for _, test := range []struct {
		rng      Range
		expected []byte
	}{
		{Range{Base: 0, Size: 12}, []byte{0, 10, 40, 41, 42, 00, 00, 20, 50, 22, 23, 00}},
		{Range{Base: 1, Size: 10}, []byte{10, 40, 41, 42, 00, 00, 20, 50, 22, 23}},
		{Range{Base: 2, Size: 3}, []byte{40, 41, 42}},
		{Range{Base: 5, Size: 2}, []byte{0, 0}},
		{Range{Base: 8, Size: 1}, []byte{50}},
	} {
		got, err := p.Slice(test.rng).Get(nil, nil)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
		check(t, test.expected, got)
	}
}

// Write layout:
//      0    1    2    3    4    5    6    7    8    9   10   11
//   ┌────╔════╤════╤════╗────┬────┬────┬────┬────┬────┬────┬────┐
// 0 │    ║A 10│ 11 │ 12 ║    │    │    │    │    │    │    │    │
//   └────╚════╧════╧════╝────┴────┴────┴────┴────┴────┴────┴────┘
//   ┌────┬────┬────┬────┬────┬────┬────╔════╤════╤════╤════╗────┐
// 1 │    │    │    │    │    │    │    ║B 20│ 21 │ 22 │ 23 ║    │
//   └────┴────┴────┴────┴────┴────┴────╚════╧════╧════╧════╝────┘
//   ┌────┬────╔════╤════╗────┬────┬────┬────┬────┬────┬────┬────┐
// 2 │    │    ║C 30│ 31 ║    │    │    │    │    │    │    │    │
//   └────┴────╚════╧════╝────┴────┴────┴────┴────┴────┴────┴────┘
//   ┌────┬────╔════╤════╤════╗────┬────┬────┬────┬────┬────┬────┐
// 3 │    │    ║D 40│ 41 │ 42 ║    │    │    │    │    │    │    │
//   └────┴────╚════╧════╧════╝────┴────┴────┴────┴────┴────┴────┘
//   ┌────┬────┬────┬────┬────┬────┬────┬────╔════╗────┬────┬────┐
// 4 │    │    │    │    │    │    │    │    ║E 50║    │    │    │
//   └────┴────┴────┴────┴────┴────┴────┴────╚════╝────┴────┴────┘
//
func TestMemoryResourceWriteReadScattered(t *testing.T) {
	d, l := database.NewInMemory(nil), log.Testing(t)
	resA, _ := database.Store([]byte{10, 11, 12}, d, l)
	resB, _ := database.Store([]byte{20, 21, 22, 23}, d, l)
	resC, _ := database.Store([]byte{30, 31}, d, l)
	resD, _ := database.Store([]byte{40, 41, 42}, d, l)
	resE, _ := database.Store([]byte{50}, d, l)

	p := Pool{}
	p.Write(1, Resource(resA, 3))
	p.Write(7, Resource(resB, 4))
	p.Write(2, Resource(resC, 2))
	p.Write(2, Resource(resD, 3))
	p.Write(8, Resource(resE, 1))

	for _, test := range []struct {
		rng      Range
		expected []byte
	}{
		{Range{Base: 0, Size: 12}, []byte{0, 10, 40, 41, 42, 00, 00, 20, 50, 22, 23, 00}},
		{Range{Base: 1, Size: 10}, []byte{10, 40, 41, 42, 00, 00, 20, 50, 22, 23}},
		{Range{Base: 2, Size: 3}, []byte{40, 41, 42}},
		{Range{Base: 5, Size: 2}, []byte{0, 0}},
		{Range{Base: 8, Size: 1}, []byte{50}},
	} {
		slice := p.Slice(test.rng)
		gotData, err := slice.Get(d, l)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
		check(t, test.expected, gotData)
		gotID, err := slice.ResourceID(d, l)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
		expectedID, _ := database.Store(test.expected, d, l)
		if gotID != expectedID {
			t.Errorf("Unexpected ID. Expected: %v, Got: %v", expectedID, gotID)
		}
	}
}
