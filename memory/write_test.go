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

package memory

import (
	"bytes"
	"testing"

	"android.googlesource.com/platform/tools/gpu/binary/endian"
	"android.googlesource.com/platform/tools/gpu/device"
)

func TestWrite(t *testing.T) {
	arch := device.Architecture{
		PointerAlignment: 4,
		PointerSize:      4,
		IntegerSize:      4,
		ByteOrder:        endian.Big,
	}

	values := []interface{}{
		uint8(0x12), int8(0x12),
		uint16(0x1234), int16(0x1234),
		uint32(0x12345678), int32(0x12345678),
		Pointer(0x87654321),
		[]uint8{0x10, 0x20, 0x30},
		"hello",
	}

	buf := &bytes.Buffer{}
	w := endian.Writer(buf, arch.ByteOrder)
	Write(w, arch, values)

	got := buf.Bytes()
	expected := []byte{
		0x12, 0x12, // uint8(0x12), int8(0x12),
		0x12, 0x34, 0x12, 0x34, // uint16(0x1234), int16(0x1234),
		0x12, 0x34, 0x56, 0x78, // uint32(0x12345678),
		0x12, 0x34, 0x56, 0x78, // int32(0x12345678),
		0x87, 0x65, 0x43, 0x21, // Pointer(0x87654321),
		0x10, 0x20, 0x30, // []uint8{0x10, 0x20, 0x30},
		'h', 'e', 'l', 'l', 'o', 0, // "hello"
	}

	if !bytes.Equal(got, expected) {
		t.Errorf("Bytes were not as expected.\nExpected: % x\nGot:      % x", expected, got)
	}
}
