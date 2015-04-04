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

package endian

import (
	"bytes"
	"reflect"
	"testing"
)

func TestReadWrite(t *testing.T) {
	for _, e := range []struct {
		name   string
		values interface{}
		data   []byte
	}{
		{"Bool",
			[]bool{true, false},
			[]byte{1, 0},
		},
		{"Int8",
			[]int8{0, 127, -128, -1},
			[]byte{0x00, 0x7f, 0x80, 0xff},
		},
		{"Uint8",
			[]uint8{0x00, 0x7f, 0x80, 0xff},
			[]byte{0x00, 0x7f, 0x80, 0xff},
		},

		{"Int16",
			[]int16{0, 32767, -32768, -1},
			[]byte{
				0x00, 0x00,
				0xff, 0x7f,
				0x00, 0x80,
				0xff, 0xff,
			}},

		{"Uint16",
			[]uint16{0, 0xbeef, 0xc0de},
			[]byte{
				0x00, 0x00,
				0xef, 0xbe,
				0xde, 0xc0,
			}},

		{"Int32",
			[]int32{0, 2147483647, -2147483648, -1},
			[]byte{
				0x00, 0x00, 0x00, 0x00,
				0xff, 0xff, 0xff, 0x7f,
				0x00, 0x00, 0x00, 0x80,
				0xff, 0xff, 0xff, 0xff,
			}},

		{"Uint32",
			[]uint32{0, 0x01234567, 0x10abcdef},
			[]byte{
				0x00, 0x00, 0x00, 0x00,
				0x67, 0x45, 0x23, 0x01,
				0xef, 0xcd, 0xab, 0x10,
			}},

		{"Int64",
			[]int64{0, 9223372036854775807, -9223372036854775808, -1},
			[]byte{
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80,
				0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
			}},

		{"Uint64",
			[]uint64{0, 0x0123456789abcdef, 0xfedcba9876543210},
			[]byte{
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0xef, 0xcd, 0xab, 0x89, 0x67, 0x45, 0x23, 0x01,
				0x10, 0x32, 0x54, 0x76, 0x98, 0xba, 0xdc, 0xfe,
			}},

		{"Float32",
			[]float32{0, 1, 64.5},
			[]byte{
				0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x80, 0x3f,
				0x00, 0x00, 0x81, 0x42,
			}},

		{"Float64",
			[]float64{0, 1, 64.5},
			[]byte{
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0, 0x3f,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x20, 0x50, 0x40,
			}},

		{"String",
			[]string{
				"Hello",
				"",
				"World",
				"こんにちは世界",
			},
			[]byte{
				'H', 'e', 'l', 'l', 'o', 0x00,
				0x00,
				'W', 'o', 'r', 'l', 'd', 0x00,
				0xe3, 0x81, 0x93, 0xe3, 0x82, 0x93, 0xe3, 0x81, 0xab, 0xe3, 0x81, 0xa1, 0xe3, 0x81, 0xaf, 0xe4, 0xb8, 0x96, 0xe7, 0x95, 0x8c, 0x00,
			}},
	} {
		b := &bytes.Buffer{}
		rv := reflect.ValueOf(Reader(b, Little))
		r := rv.MethodByName(e.name)
		wv := reflect.ValueOf(Writer(b, Little))
		w := wv.MethodByName(e.name)
		s := reflect.ValueOf(e.values)
		for i := 0; i < s.Len(); i++ {
			w.Call([]reflect.Value{s.Index(i)})
		}
		if !bytes.Equal(e.data, b.Bytes()) {
			t.Errorf(`%v gave unexpected bytes.
Expected: %# x
Got:      %# x`, e.name, e.data, b.Bytes())
		}
		for i := 0; i < s.Len(); i++ {
			expected := s.Index(i)
			result := r.Call(nil)
			got := result[0]
			err := result[1]
			if !err.IsNil() {
				t.Errorf("%v %d gave unexpected error: %v", i, err)
			}
			if !reflect.DeepEqual(expected.Interface(), got.Interface()) {
				t.Errorf("%v %d gave unexpected value. Expected: %v, got: %v", e.name, i, expected.Interface(), got.Interface())
			}

		}
	}
}
