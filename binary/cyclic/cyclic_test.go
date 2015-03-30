// Copyright (C) 2014 The Android Open Source Project
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

package cyclic

import (
	"bytes"
	"io/ioutil"
	"reflect"
	"testing"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/vle"
)

type testObjectA struct {
	binary.Generate
	data string
}

type testObjectB struct {
	binary.Generate
	data string
}

type testObjectC struct {
	binary.Generate "disable"
	data            string
}

var testObjA = &testObjectA{data: "ObjectA"}
var testObjB = &testObjectB{data: "ObjectB"}
var testObjC = &testObjectC{data: "ObjectC"}

func TestEncodeDecode(t *testing.T) {
	for _, v := range []struct {
		name   string
		values []binary.Object
		data   []byte
	}{
		{"Nil",
			[]binary.Object{nil},
			[]byte{0},
		},
		{"One",
			[]binary.Object{testObjA},
			[]byte{
				0x01,
				0x47, 0x30, 0xec, 0x1c, 0xa6, 0x6f, 0x5f, 0xc7, 0xc0, 0x64, 0x3a, 0xc1, 0xd6, 0x45, 0x06, 0x06, 0xb1, 0x59, 0x1b, 0x97,
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'A',
			},
		},
		{"Repeat",
			[]binary.Object{testObjA, testObjA},
			[]byte{
				0x01,
				0x47, 0x30, 0xec, 0x1c, 0xa6, 0x6f, 0x5f, 0xc7, 0xc0, 0x64, 0x3a, 0xc1, 0xd6, 0x45, 0x06, 0x06, 0xb1, 0x59, 0x1b, 0x97,
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'A',

				0x01,
			},
		},
		{"Many",
			[]binary.Object{testObjA, testObjB, testObjA, nil},
			[]byte{
				0x01,
				0x47, 0x30, 0xec, 0x1c, 0xa6, 0x6f, 0x5f, 0xc7, 0xc0, 0x64, 0x3a, 0xc1, 0xd6, 0x45, 0x06, 0x06, 0xb1, 0x59, 0x1b, 0x97,
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'A',

				0x02,
				0xc8, 0x39, 0xdf, 0x69, 0x4e, 0x4e, 0xce, 0x03, 0xed, 0x8b, 0x41, 0x01, 0xe2, 0x46, 0x52, 0x88, 0x8e, 0x78, 0xc8, 0x67,
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'B',

				0x01,

				0x00,
			},
		},
	} {
		b := &bytes.Buffer{}
		e := Encoder(vle.Writer(b))
		d := Decoder(vle.Reader(b))
		for i, o := range v.values {
			if err := e.Object(o); err != nil {
				t.Errorf("%v[%v] encode gave unexpected error: %v", v.name, i, err)
			}
		}
		if !bytes.Equal(v.data, b.Bytes()) {
			t.Errorf(`%v gave unexpected bytes.
Expected: %# x
Got:      %# x`, v.name, v.data, b.Bytes())
			for i, o := range v.values {
				if got, err := d.Object(); err != nil {
					t.Errorf("%v[%v] decode gave unexpected error: %v", v.name, i, err)
				} else if !reflect.DeepEqual(o, got) {
					t.Errorf("%v[%v] unexpected object. Expected: %v, got: %v", v.name, i, o, got)
				}
			}
		}
	}
}

func TestUnknownTypeError(t *testing.T) {
	e := Encoder(vle.Writer(ioutil.Discard))
	if err := e.Object(testObjC); err == nil {
		t.Errorf("Expected error encoding unknown type")
	}
	d := Decoder(vle.Reader(bytes.NewBuffer([]byte{
		0x01,
		0x0C, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x07,
		'O', 'b', 'j', 'e', 'c', 't', 'C',
	})))
	if _, err := d.Object(); err == nil {
		t.Errorf("Expected error decoding unknown type")
	}
}
