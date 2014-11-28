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

// BitStream provides methods for reading and writing bits to a slice of bytes.
// Bits are packed in a least-significant-bit to most-significant-bit order.
type BitStream struct {
	Data     []byte // The byte slice containing the bits
	ReadPos  uint32 // The current read offset from the start of the Data slice (in bits)
	WritePos uint32 // The current write offset from the start of the Data slice (in bits)
}

// ReadBit reads a single bit from the BitStream, incrementing ReadPos by one.
func (s *BitStream) ReadBit() uint32 {
	ReadPos := s.ReadPos
	s.ReadPos = ReadPos + 1
	return (uint32(s.Data[ReadPos/8]) >> (ReadPos % 8)) & 1
}

// WriteBit writes a single bit to the BitStream, incrementing WritePos by one.
func (s *BitStream) WriteBit(bit uint32) {
	WritePos := s.WritePos
	b := WritePos / 8
	if b == uint32(len(s.Data)) {
		s.Data = append(s.Data, byte(0))
	}
	s.Data[b] |= byte((bit & 1) << (WritePos % 8))
	s.WritePos = WritePos + 1
}

// Read reads the specified number of bits from the BitStream, increamenting the ReadPos by the
// specified number of bits and returning the bits packed into a uint32. The bits are packed into
// the uint32 from LSB to MSB.
func (s *BitStream) Read(count uint32) uint32 {
	byteIdx := s.ReadPos / 8
	bitIdx := s.ReadPos & 7

	// Start
	val := uint32(s.Data[byteIdx]) >> bitIdx
	readCount := 8 - bitIdx
	if count <= readCount {
		s.ReadPos += count
		return val & ((1 << count) - 1)
	}
	s.ReadPos += readCount
	byteIdx++
	bitIdx = 0

	// Whole bytes
	for ; readCount+7 < count; readCount += 8 {
		val |= uint32(s.Data[byteIdx]) << readCount
		byteIdx++
		s.ReadPos += 8
	}

	// Remainder
	rem := count - readCount
	if rem > 0 {
		val |= (uint32(s.Data[byteIdx]) & ((1 << rem) - 1)) << readCount
		s.ReadPos += rem
	}
	return val
}

// Write writes the specified number of bits from the packed uint32, increamenting the WritePos by
// the specified number of bits. The bits are read from the uint32 from LSB to MSB.
func (s *BitStream) Write(bits, count uint32) {
	// TODO: Slowest implementation ever? Quite possibly. Optimise
	for i := uint32(0); i < count; i++ {
		s.WriteBit(bits >> i)
	}
}
