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

package atom_test

import (
	"bytes"
	"reflect"
	"testing"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/atom/test"
	"android.googlesource.com/platform/tools/gpu/binary/cyclic"
	"android.googlesource.com/platform/tools/gpu/binary/objects"
	bt "android.googlesource.com/platform/tools/gpu/binary/test"
	"android.googlesource.com/platform/tools/gpu/binary/vle"
)

var testList = atom.List{
	&test.AtomA{},
	&test.AtomB{Bool: true},
	&test.AtomC{String: "Pizza"},
}

var testData = bt.Bytes{}.Add(
	0x03, // Atom 0: object sid + encoded
	0x03, // Atom 0: type sid + encoded
).ID(test.AtomAID).Add(
	0x00, // Atom 0: Id
	0x00, // Atom 0: Flags
	0x05, // Atom 1: object sid + encoded
	0x05, // Atom 1: sid + encoded
).ID(test.AtomBID).Add(
	0x00, // Atom 1: Id
	0x01, // Atom 1: Bool
	0x07, // Atom 2: object sid + encoded
	0x07, // Atom 2: sid + encoded
).ID(test.AtomCID).Add(
	0x05, 'P', 'i', 'z', 'z', 'a', // Atom 2: String
	0x09, // Terminator: object sid + encoded
	0x09, // Terminator: sid + encoded
).ID(objects.TerminatorID).Data

func TestAtomListEncode(t *testing.T) {
	buf := &bytes.Buffer{}
	enc := cyclic.Encoder(vle.Writer(buf))
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
	list := atom.List{}
	err := list.Decode(cyclic.Decoder(vle.Reader(bytes.NewBuffer(testData))))
	if err != nil {
		t.Errorf("Decode returned unexpected error: %v", err)
	}
	if !reflect.DeepEqual(testList, list) {
		t.Errorf("Decoded list was not as expected.\nExpected: %#v\nGot:      %#v", testList, list)
	}
}

type writeRecord struct {
	id   atom.ID
	atom atom.Atom
}
type writeRecordList []writeRecord

func (t *writeRecordList) Write(id atom.ID, atom atom.Atom) { *t = append(*t, writeRecord{id, atom}) }

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func TestAtomListWriteTo(t *testing.T) {
	expected := writeRecordList{
		writeRecord{0, &test.AtomA{}},
		writeRecord{1, &test.AtomB{Bool: true}},
		writeRecord{2, &test.AtomC{String: "Pizza"}},
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
