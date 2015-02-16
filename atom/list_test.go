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

var expectedList = List{&testAtomA{Int32: 100}, &testAtomB{Bool: true}, &testAtomC{String: "Pizza"}}
var expectedData = []byte{
	0x08, 0x00, // Atom 0: Size (+2 for type)
	byte(testAtomIdA & 0xff), byte(testAtomIdA >> 8), // Atom 0: Type
	0x64, 0x00, 0x00, 0x00, // Atom 0: Data

	0x05, 0x00, // Atom 1: Size (+2 for type)
	byte(testAtomIdB & 0xff), byte(testAtomIdB >> 8), // Atom 1: Type
	0x01, // Atom 1: Data

	0x0d, 0x00, // Atom 2: Size (+2 for type)
	byte(testAtomIdC & 0xff), byte(testAtomIdC >> 8), // Atom 2: Type
	0x05, 0x00, 0x00, 0x00, 'P', 'i', 'z', 'z', 'a', // Atom 2: Data

	0x00, 0x00, // EOS 0
}

func TestAtomListEncode(t *testing.T) {
	buf := &bytes.Buffer{}
	enc := binary.NewEncoder(buf)
	err := expectedList.Encode(enc)
	if err != nil {
		t.Errorf("Encode returned unexpected error: %v", err)
	}
	got := buf.Bytes()
	if !bytes.Equal(expectedData, got) {
		t.Errorf("Encoded data was not as expected.\nExpected: %v\nGot:      %v", expectedData, got)
	}
}

func TestAtomListDecode(t *testing.T) {
	list := List{}
	err := list.Decode(binary.NewDecoder(bytes.NewBuffer(expectedData)))
	if err != nil {
		t.Errorf("Decode returned unexpected error: %v", err)
	}
	if !reflect.DeepEqual(expectedList, list) {
		t.Errorf("Decoded list was not as expected.\nExpected: %#v\nGot:      %#v", expectedList, list)
	}
}
