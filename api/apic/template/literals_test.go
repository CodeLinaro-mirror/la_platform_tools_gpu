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

package template

import "testing"

func TestBitposWithNoBitsSet(t *testing.T) {
	in := uint32(0)
	expected := -1
	got := (&Functions{}).Bitpos(in)
	if got != expected {
		t.Errorf("Bitpos(%#v) returned unexpected value. Expected: %#v, got: %#v", in, expected, got)
	}
}

func TestBitposWithOneBitSet(t *testing.T) {
	for pos := uint(0); pos < 32; pos++ {
		in := uint32(1 << pos)
		expected := int(pos)
		got := (&Functions{}).Bitpos(in)
		if got != expected {
			t.Errorf("Bitpos(%#v) returned unexpected value. Expected: %#v, got: %#v", in, expected, got)
		}
	}
}

func TestBitposWithTwoBitsSet(t *testing.T) {
	in := uint32(0x8080)
	expected := -1
	got := (&Functions{}).Bitpos(in)
	if got != expected {
		t.Errorf("Bitpos(%#v) returned unexpected value. Expected: %#v, got: %#v", in, expected, got)
	}
}
