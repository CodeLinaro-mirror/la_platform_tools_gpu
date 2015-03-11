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

package atom

import (
	"bytes"
	"reflect"
	"testing"

	"android.googlesource.com/platform/tools/gpu/binary"
)

var testList = List{
	&testAtomA{Context: 0x10, Int32: 100},
	&testAtomB{Context: 0x20, Bool: true},
	&testAtomC{Context: 0x10, String: "Pizza"},
}
var testData = []byte{
	0x06,       // Atom 0: Size (+2 for type)
	0x0a,       // Atom 0: Type
	0x10,       // Atom 0: Context
	0x80, 0xc8, // Atom 0: Data

	0x05, // Atom 1: Size (+2 for type)
	0x14, // Atom 1: Type
	0x20, // Atom 1: Context
	0x01, // Atom 1: Data

	0x0a,                          // Atom 2: Size (+2 for type)
	0x1e,                          // Atom 2: Type
	0x10,                          // Atom 2: Context
	0x05, 'P', 'i', 'z', 'z', 'a', // Atom 2: Data

	0x00, // EOS 0
}

func TestAtomListEncode(t *testing.T) {
	buf := &bytes.Buffer{}
	enc := binary.NewEncoder(buf)
	err := testList.Encode(enc)
	if err != nil {
		t.Errorf("Encode returned unexpected error: %v", err)
	}
	got := buf.Bytes()
	if !bytes.Equal(testData, got) {
		t.Errorf("Encoded data was not as expected.\nExpected: % x\nGot:      % x", testData, got)
	}
}

func TestAtomListDecode(t *testing.T) {
	list := List{}
	err := list.Decode(binary.NewDecoder(bytes.NewBuffer(testData)))
	if err != nil {
		t.Errorf("Decode returned unexpected error: %v", err)
	}
	if !reflect.DeepEqual(testList, list) {
		t.Errorf("Decoded list was not as expected.\nExpected: %#v\nGot:      %#v", testList, list)
	}
}

type writeRecord struct {
	id   ID
	atom Atom
}
type writeRecordList []writeRecord

func (t *writeRecordList) Write(id ID, atom Atom) { *t = append(*t, writeRecord{id, atom}) }

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func TestAtomListWriteTo(t *testing.T) {
	expected := writeRecordList{
		writeRecord{0, &testAtomA{Context: 0x10, Int32: 100}},
		writeRecord{1, &testAtomB{Context: 0x20, Bool: true}},
		writeRecord{3, &EOS{Context: 0x20}},
		writeRecord{2, &testAtomC{Context: 0x10, String: "Pizza"}},
		writeRecord{4, &EOS{Context: 0x10}},
	}
	got := writeRecordList{}
	testList.WriteTo(&got)

	matched := len(expected) == len(got)

	if matched {
		for i := range expected {
			e, g := expected[i], got[i]
			if e.id != g.id || !reflect.DeepEqual(g.atom, e.atom) {
				matched = false
				break
			}
		}
	}

	if !matched {
		c := max(len(expected), len(got))
		for i := 0; i < c; i++ {
			if i > len(got) {
				t.Errorf("(%d) Expected: %#v Got: <nothing>", i, expected[i])
				continue
			}
			e, g := expected[i], got[i]
			if e.id != g.id || !reflect.DeepEqual(g.atom, e.atom) {
				t.Errorf("(%d) Expected: %#v Got: %#v", i, e, g)
				continue
			}
			t.Logf("(%d) Matched: %#v", i, g)
		}
	}
}
