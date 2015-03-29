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

package flat

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
	data string
}

func (o testObjectC) Encode(e binary.Encoder) error  { return nil }
func (o *testObjectC) Decode(d binary.Decoder) error { return nil }

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
			[]byte{
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
		},
		{"One",
			[]binary.Object{testObjA},
			[]byte{
				0x05, 0x32, 0xe2, 0xea, 0x69, 0x05, 0x2c, 0x2e, 0x64, 0x6b, 0xfa, 0xdf, 0xc5, 0x8b, 0x48, 0x94, 0x92, 0xbb, 0xe2, 0x21,
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'A',
			},
		},
		{"Repeat",
			[]binary.Object{testObjA, testObjA},
			[]byte{
				0x05, 0x32, 0xe2, 0xea, 0x69, 0x05, 0x2c, 0x2e, 0x64, 0x6b, 0xfa, 0xdf, 0xc5, 0x8b, 0x48, 0x94, 0x92, 0xbb, 0xe2, 0x21,
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'A',

				0x05, 0x32, 0xe2, 0xea, 0x69, 0x05, 0x2c, 0x2e, 0x64, 0x6b, 0xfa, 0xdf, 0xc5, 0x8b, 0x48, 0x94, 0x92, 0xbb, 0xe2, 0x21,
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'A',
			},
		},
		{"Many",
			[]binary.Object{testObjA, testObjB, testObjA, nil},
			[]byte{
				0x05, 0x32, 0xe2, 0xea, 0x69, 0x05, 0x2c, 0x2e, 0x64, 0x6b, 0xfa, 0xdf, 0xc5, 0x8b, 0x48, 0x94, 0x92, 0xbb, 0xe2, 0x21,
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'A',

				0x3c, 0x38, 0xd2, 0x5b, 0x79, 0x7e, 0x5f, 0x61, 0xd8, 0xc3, 0xc4, 0xaa, 0xca, 0xcd, 0xc6, 0x73, 0x50, 0x21, 0x96, 0xe5,
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'B',

				0x05, 0x32, 0xe2, 0xea, 0x69, 0x05, 0x2c, 0x2e, 0x64, 0x6b, 0xfa, 0xdf, 0xc5, 0x8b, 0x48, 0x94, 0x92, 0xbb, 0xe2, 0x21,
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'A',

				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
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
		0x0C, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x07,
		'O', 'b', 'j', 'e', 'c', 't', 'C',
	})))
	if _, err := d.Object(); err == nil {
		t.Errorf("Expected error decoding unknown type")
	}
}
