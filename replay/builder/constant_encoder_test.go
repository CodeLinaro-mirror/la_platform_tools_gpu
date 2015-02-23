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

package builder

import (
	"reflect"
	"testing"

	"android.googlesource.com/platform/tools/gpu/replay/value"
)

func TestConstantEncoderCache(t *testing.T) {
	c := newConstantEncoder(4)

	addr1 := c.writeValues(value.U32(0x1234), value.S16(-1))
	addr2 := c.writeValues(value.U32(0x1234), value.S16(-1))

	if addr1 != addr2 {
		t.Errorf("expected %v got %v", addr1, addr2)
	}
}

func TestConstantEncoderAlignment(t *testing.T) {
	c := newConstantEncoder(8)

	c.writeValues(value.U32(0x1234))
	c.writeValues(value.S16(-1))

	expected := []byte{
		0x34, 0x12, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0xff, 0xff,
	}

	if !reflect.DeepEqual(expected, c.data) {
		t.Errorf("expected %v got %v", expected, c.data)
	}
}
