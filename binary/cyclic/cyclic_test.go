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
	"testing"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/test"
	"android.googlesource.com/platform/tools/gpu/binary/vle"
)

func TestValue(t *testing.T) {
	for _, entry := range []test.Entry{
		{
			Name:   "One",
			Values: []binary.Object{test.ObjectA},
			Data: []byte{
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'A',
			},
		},
		{
			Name:   "Repeat",
			Values: []binary.Object{test.ObjectA, test.ObjectA},
			Data: []byte{
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'A',
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'A',
			},
		},
		{
			Name:   "Many",
			Values: []binary.Object{test.ObjectA, test.ObjectB, test.ObjectA},
			Data: []byte{
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'A',
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'B',
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'A',
			},
		},
	} {
		b := &bytes.Buffer{}
		test.EncodeValue(t, entry, Encoder(vle.Writer(b)), b)
		r := bytes.NewReader(entry.Data)
		test.DecodeValue(t, entry, Decoder(vle.Reader(r)), r)
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
			Data: []byte{
				0x03, // object sid + encoded
				0x03, // type sid + encoded
				0xa4, 0xbe, 0x00, 0x04, 0x4c, 0x84, 0x76, 0x86, 0xdc, 0x77, 0x63, 0x6d, 0x19, 0xdd, 0x63, 0x33, 0x17, 0x38, 0xbf, 0x24,
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'A',
			},
		},
		{
			Name:   "Repeat",
			Values: []binary.Object{test.ObjectA, test.ObjectA},
			Data: []byte{
				0x03, // object sid + encoded
				0x03, // type sid + encoded
				0xa4, 0xbe, 0x00, 0x04, 0x4c, 0x84, 0x76, 0x86, 0xdc, 0x77, 0x63, 0x6d, 0x19, 0xdd, 0x63, 0x33, 0x17, 0x38, 0xbf, 0x24,
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'A',

				0x02, // repeated object sid
			},
		},
		{
			Name:   "Many",
			Values: []binary.Object{test.ObjectA, test.ObjectB, test.ObjectA, nil},
			Data: []byte{
				0x03, // object sid + encoded
				0x03, // type sid + encoded
				0xa4, 0xbe, 0x00, 0x04, 0x4c, 0x84, 0x76, 0x86, 0xdc, 0x77, 0x63, 0x6d, 0x19, 0xdd, 0x63, 0x33, 0x17, 0x38, 0xbf, 0x24,
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'A',

				0x05, // object sid + encoded
				0x05, // type sid + encoded
				0x73, 0xbd, 0xff, 0x55, 0x9c, 0xc4, 0x5b, 0xe3, 0xaf, 0x72, 0xfd, 0xb6, 0x97, 0xfb, 0x0e, 0xe1, 0x8d, 0x19, 0xa9, 0x67,
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'B',

				0x02, // repeated object sid

				0x00, // nil object sid
			},
		},
	} {
		b := &bytes.Buffer{}
		test.EncodeObject(t, entry, Encoder(vle.Writer(b)), b)
		r := bytes.NewReader(entry.Data)
		test.DecodeObject(t, entry, Decoder(vle.Reader(r)), r)
	}
}

func TestUnknownTypeError(t *testing.T) {
	e := Encoder(vle.Writer(ioutil.Discard))
	if err := e.Object(test.BadObject); err == nil {
		t.Errorf("Expected error encoding unknown type")
	}
	d := Decoder(vle.Reader(bytes.NewBuffer([]byte{
		0x03, // object sid + encoded
		0x03, // type sid + encoded
		0x0C, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x09,
		'B', 'a', 'd', 'O', 'b', 'j', 'e', 'c', 't',
	})))
	if _, err := d.Object(); err == nil {
		t.Errorf("Expected error decoding unknown type")
	}
}
