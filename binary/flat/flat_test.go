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
	"testing"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/pod"
	"android.googlesource.com/platform/tools/gpu/binary/test"
	"android.googlesource.com/platform/tools/gpu/binary/vle"
)

func TestValue(t *testing.T) {
	for _, entry := range []test.Entry{
		{
			Name:   "One",
			Values: []interface{}{test.ObjectA},
			Data: []byte{
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'A',
			},
		},
		{
			Name:   "Repeat",
			Values: []interface{}{test.ObjectA, test.ObjectA},
			Data: []byte{
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'A',
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'A',
			},
		},
		{
			Name:   "Many",
			Values: []interface{}{test.ObjectA, test.ObjectB, test.ObjectA},
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
			Values: []interface{}{nil},
			Data:   test.Bytes{}.ID(binary.ID{}).Data,
		},
		{
			Name:   "One",
			Values: []interface{}{test.ObjectA},
			Data: test.Bytes{}.ID(test.TypeAID).Add(
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'A',
			).Data,
		},
		{
			Name:   "Repeat",
			Values: []interface{}{test.ObjectA, test.ObjectA},
			Data: test.Bytes{}.ID(test.TypeAID).Add(
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'A',
			).ID(test.TypeAID).Add(
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'A',
			).Data,
		},
		{
			Name:   "Many",
			Values: []interface{}{test.ObjectA, test.ObjectB, test.ObjectA, nil},
			Data: test.Bytes{}.ID(test.TypeAID).Add(
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'A',
			).ID(test.TypeBID).Add(
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'B',
			).ID(test.TypeAID).Add(
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'A',
			).ID(binary.ID{}).Data,
		},
		{
			Name:   "POD",
			Values: []interface{}{true, int8(0x1), uint32(0x10), true, int8(0x1)},
			Data: test.Bytes{}.ID((*pod.Bool)(nil).Class().ID()).Add(0x01).
				ID((*pod.Int8)(nil).Class().ID()).Add(0x01).
				ID((*pod.Uint32)(nil).Class().ID()).Add(0x10).
				ID((*pod.Bool)(nil).Class().ID()).Add(0x01).
				ID((*pod.Int8)(nil).Class().ID()).Add(0x01).
				Data,
		},
	} {
		b := &bytes.Buffer{}
		test.EncodeObject(t, entry, Encoder(vle.Writer(b)), b)
		r := bytes.NewReader(entry.Data)
		test.DecodeObject(t, entry, Decoder(vle.Reader(r)), r)
	}
}

func TestUnknownTypeError(t *testing.T) {
	d := Decoder(vle.Reader(bytes.NewBuffer([]byte{
		0x0C, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x09,
		'B', 'a', 'd', 'O', 'b', 'j', 'e', 'c', 't',
	})))
	if _, err := d.Object(); err == nil {
		t.Errorf("Expected error decoding unknown type")
	}
}
