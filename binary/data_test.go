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

package binary_test

import (
	"bytes"
	"testing"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/cyclic"
	"android.googlesource.com/platform/tools/gpu/binary/vle"
)

func TestDataEncodeDecode(t *testing.T) {
	for _, v := range []struct {
		name string
		data binary.Data
	}{
		{"Basic",
			binary.Data{0x10, 0x20, 0x30, 0xaa, 0xbb, 0xcc},
		},
	} {
		b := &bytes.Buffer{}
		e := cyclic.Encoder(vle.Writer(b))
		d := cyclic.Decoder(vle.Reader(b))
		if err := v.data.Encode(e); err != nil {
			t.Errorf("%v encode gave unexpected error: %v", v.name, err)
		}
		got := binary.Data{}
		if err := got.Decode(d); err != nil {
			t.Errorf("%v decode gave unexpected error: %v", v.name, err)
		}
		if !bytes.Equal(v.data, got) {
			t.Errorf("%v gave unexpected value. Expected: %v, got: %v", v.name, v.data, got)
		}
	}
}
