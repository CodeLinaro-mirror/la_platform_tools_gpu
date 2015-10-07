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
	"reflect"
	"testing"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/test"
	"android.googlesource.com/platform/tools/gpu/binary/vle"
)

func EncodeObject(t *testing.T, entry test.Entry, e binary.Encoder, buf *bytes.Buffer) {
	for i, o := range entry.Values {
		e.Object(o)
		if e.Error() != nil {
			t.Errorf("%v[%v] Object gave unexpected error: %v", entry.Name, i, e.Error())
		}
	}
	test.VerifyData(t, entry, buf)
}

func DecodeObject(t *testing.T, entry test.Entry, d binary.Decoder, reader *bytes.Reader) {
	for i, o := range entry.Values {
		got := d.Object()
		if d.Error() != nil {
			t.Errorf("%v[%v] Object gave unexpected error: %v", entry.Name, i, d.Error())
		} else if !reflect.DeepEqual(o, got) {
			t.Errorf("%v[%v] unexpected object. Expected: %+v, got: %+v", entry.Name, i, o, got)
		}
	}
}

func TestObject(t *testing.T) {
	for _, entry := range []test.Entry{
		{
			Name:   "Nil",
			Values: []binary.Object{nil},
			Data:   []byte{0},
		},
		{
			Name:   "One",
			Values: []binary.Object{test.ObjectA},
			Data: test.Bytes{}.Add(
				0x03, // object sid + encoded
				0x03, // type sid + encoded
			).Add(test.EntityA...).Add(
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'A',
			).Data,
		},
		{
			Name:   "Repeat",
			Values: []binary.Object{test.ObjectA, test.ObjectA},
			Data: test.Bytes{}.Add(
				0x03, // object sid + encoded
				0x03, // type sid + encoded
			).Add(test.EntityA...).Add(
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'A',

				0x02, // repeated object sid
			).Data,
		},
		{
			Name:   "Many",
			Values: []binary.Object{test.ObjectA, test.ObjectB, test.ObjectA, nil},
			Data: test.Bytes{}.Add(
				0x03, // object sid + encoded
				0x03, // type sid + encoded
			).Add(test.EntityA...).Add(
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'A',

				0x05, // object sid + encoded
				0x05, // type sid + encoded
			).Add(test.EntityB...).Add(
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'B',

				0x02, // repeated object sid

				0x00, // nil object sid
			).Data,
		},
	} {
		b := &bytes.Buffer{}
		EncodeObject(t, entry, Encoder(vle.Writer(b)), b)
		r := bytes.NewReader(entry.Data)
		DecodeObject(t, entry, Decoder(vle.Reader(r)), r)
	}
}

func TestUnknownTypeError(t *testing.T) {
	d := Decoder(vle.Reader(bytes.NewBuffer([]byte{
		0x03, // object sid + encoded
		0x03, // type sid + encoded
		0x0C, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x09,
		'B', 'a', 'd', 'O', 'b', 'j', 'e', 'c', 't',
	})))
	d.Object()
	if d.Error() == nil {
		t.Errorf("Expected error decoding unknown type")
	}
}
