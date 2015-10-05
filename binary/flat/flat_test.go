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
