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

package binary

import (
	"bytes"
	"testing"
)

var testByteSequence = []byte{
	//        LSB        MSB
	0x00, // (0) 00000000
	0x01, // (1) 10000000
	0x02, // (2) 01000000
	0x10, // (3) 00001000
	0x80, // (4) 00000001
	0x3A, // (5) 01011100
	0x02, // (6) 01000000
	0x3C, // (7) 00111100
}

var testBitSequence = []struct {
	bits, count, pos uint32
}{
	// (0) 0₀0₁0 0 0 0 0 0
	{bits: 0, count: 2, pos: 2},
	// (0) 0 0 0₀0₁0₂0 0 0
	{bits: 0, count: 3, pos: 5},
	// (0) 0 0 0 0 0 0₀0₁0₂
	{bits: 0, count: 3, pos: 8},
	// (1) 1₀0₁0₂0 0 0 0 0
	{bits: 1, count: 3, pos: 11},
	// (1) 1 0 0 0₀0₁0₂0₃0₄
	// (2) 0₅1₆0₇0₈0₉0 0 0
	{bits: 64, count: 10, pos: 21},
	// (2) 0 1 0 0 0 0₀0₁0₂
	// (3) 0₃0₄0₅0₆1₇0₈0₉0
	{bits: 128, count: 10, pos: 31},
	// (3) 0 0 0 1 0 0 0 0₀
	// (4) 0₁0₂0₃0₄0₅0₆0₇1₈
	// (5) 0₉1ₐ0 1 1 1 0 0
	{bits: 1280, count: 11, pos: 42},
	// (5) 0 1 0₀1₁1₂1₃0₄0₅
	// (6) 0₆1₇0₈0₉0ₐ0 0 0
	{bits: 142, count: 11, pos: 53},
	// (6) 0 1 0 0 0 0₀0₁0₂
	// (7) 0₃0₄1 1 1 1 0 0
	{bits: 0, count: 5, pos: 58},
	// (7) 0 0 1₀1₁1₂1 0 0
	{bits: 7, count: 3, pos: 61},
	// (7) 0 0 1 1 1 1₀0₁0₂
	{bits: 1, count: 3, pos: 64},
}

func TestBitStreamReadBit(t *testing.T) {
	data := []byte{
		//    LSB        MSB
		0x00, // 00000000
		0x01, // 10000000
		0x02, // 01000000
		0x10, // 00001000
		0x80, // 00000001
	}
	bits := []uint32{
		0, 0, 0, 0, 0, 0, 0, 0,
		1, 0, 0, 0, 0, 0, 0, 0,
		0, 1, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 1, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 1,
	}
	bs := BitStream{Data: data}
	for i, expected := range bits {
		actual := bs.ReadBit()
		if expected != actual {
			t.Errorf("Bit %d was not as expected. Expected %v, got %v", i, expected, actual)
		}
	}
}

func TestBitStreamRead(t *testing.T) {
	bs := BitStream{Data: testByteSequence}
	for i, c := range testBitSequence {
		if bits := bs.Read(c.count); c.bits != bits {
			t.Errorf("Bits returned from segment %d were not as expected. Expected %v, got %v", i, c.bits, bits)
		}
		if pos := bs.ReadPos; c.pos != pos {
			t.Errorf("Read position after segment %d was not as expected. Expected %v, got %v", i, c.pos, pos)
		}
	}
}

func TestBitStreamZeroRemainder(t *testing.T) {
	data := []byte{
		//        LSB        MSB
		0x00, // (0) 00000000
		0x01, // (1) 10000000
		0x02, // (2) 01000000
		0x10, // (3) 00001000
	}
	bs := BitStream{Data: data}
	expected := uint32(0x10020100)
	if got := bs.Read(32); got != expected {
		t.Errorf("Bits returned were not as expected. Expected %v, got %v", expected, got)
	}
}

func TestBitStreamWriteBit(t *testing.T) {
	expected := []byte{
		//    LSB        MSB
		0x00, // 00000000
		0x01, // 10000000
		0x02, // 01000000
		0x10, // 00001000
		0x80, // 00000001
	}
	bits := []uint32{
		0, 0, 0, 0, 0, 0, 0, 0,
		1, 0, 0, 0, 0, 0, 0, 0,
		0, 1, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 1, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 1,
	}
	bs := BitStream{}
	for _, b := range bits {
		bs.WriteBit(b)
	}

	if got := bs.Data; !bytes.Equal(got, expected) {
		t.Errorf("Bytes written was not as expected. Expected %v, got %v", expected, got)
	}
}

func TestBitStreamWrite(t *testing.T) {
	bs := BitStream{}
	for i, c := range testBitSequence {
		bs.Write(c.bits, c.count)
		if pos := bs.WritePos; c.pos != pos {
			t.Errorf("Write position after segment %d was not as expected. Expected %v, got %v", i, c.pos, pos)
		}
	}
	expected := testByteSequence
	if got := bs.Data; !bytes.Equal(got, expected) {
		t.Errorf("Bytes written was not as expected. Expected %v, got %v", expected, got)
	}
}

func BenchmarkBitstreamRead(b *testing.B) {
	data := make([]byte, b.N*32)
	bs := BitStream{Data: data}
	for i := 0; i < b.N; i++ {
		bs.Read(uint32(i & 31))
	}
}
