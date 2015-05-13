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

package opcode

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/replay/protocol"
)

func bit(bits, idx uint32) bool {
	if bits&(1<<idx) != 0 {
		return true
	} else {
		return false
	}
}

func setBit(bits, idx uint32, v bool) uint32 {
	if v {
		return bits | (1 << idx)
	} else {
		return bits & ^(1 << idx)
	}
}

// ┏━━┯━━┯━━┯━━┯━━┯━━┳━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┓
// ┃c │c │c │c │c │c ┃0 │0 │0 │0 │0 │0 │0 │0 │0 │0 │0 │0 │0 │0 │0 │0 │0 │0 │0 │0 │0 │0 │0 │0 │0 │0 ┃
// ┃ ₅│ ₄│ ₃│ ₂│ ₁│ ₀┃₂₅│₂₄│₂₃│₂₂│₂₁│₂₀│₁₉│₁₈│₁₇│₁₆│₁₅│₁₄│₁₃│₁₂│₁₁│₁₀│ ₉│ ₈│ ₇│ ₆│ ₅│ ₄│ ₃│ ₂│ ₁│ ₀┃
// ┡━━┿━━┿━━┿━━┿━━┿━━╇━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┩
// │₃₁│₃₀│₂₉│₂₈│₂₇│₂₆│₂₅│₂₄│₂₃│₂₂│₂₁│₂₀│₁₉│₁₈│₁₇│₁₆│₁₅│₁₄│₁₃│₁₂│₁₁│₁₀│ ₉│ ₈│ ₇│ ₆│ ₅│ ₄│ ₃│ ₂│ ₁│ ₀│
// └──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┘
func packC(c uint32) uint32 {
	if c >= 0x3f {
		panic(fmt.Errorf("c exceeds 6 bits (0x%x)", c))
	}
	return c << 26
}

// ┏━━┯━━┯━━┯━━┯━━┯━━┳━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┓
// ┃c │c │c │c │c │c ┃x │x │x │x │x │x │x │x │x │x │x │x │x │x │x │x │x │x │x │x │x │x │x │x │x │x ┃
// ┃ ₅│ ₄│ ₃│ ₂│ ₁│ ₀┃₂₅│₂₄│₂₃│₂₂│₂₁│₂₀│₁₉│₁₈│₁₇│₁₆│₁₅│₁₄│₁₃│₁₂│₁₁│₁₀│ ₉│ ₈│ ₇│ ₆│ ₅│ ₄│ ₃│ ₂│ ₁│ ₀┃
// ┡━━┿━━┿━━┿━━┿━━┿━━╇━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┩
// │₃₁│₃₀│₂₉│₂₈│₂₇│₂₆│₂₅│₂₄│₂₃│₂₂│₂₁│₂₀│₁₉│₁₈│₁₇│₁₆│₁₅│₁₄│₁₃│₁₂│₁₁│₁₀│ ₉│ ₈│ ₇│ ₆│ ₅│ ₄│ ₃│ ₂│ ₁│ ₀│
// └──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┘
func packCX(c uint32, x uint32) uint32 {
	if x > 0x3ffffff {
		panic(fmt.Errorf("x exceeds 26 bits (0x%x)", x))
	}
	return packC(c) | x
}

// ┏━━┯━━┯━━┯━━┯━━┯━━┳━━┯━━┯━━┯━━┯━━┯━━┳━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┯━━┓
// ┃c │c │c │c │c │c ┃y │y │y │y │y │y ┃z │z │z │z │z │z │z │z │z │z │z │z │z │z │z │z │z │z │z │z ┃
// ┃ ₅│ ₄│ ₃│ ₂│ ₁│ ₀┃ ₅│ ₄│ ₃│ ₂│ ₁│ ₀┃₁₉│₁₈│₁₇│₁₆│₁₅│₁₄│₁₃│₁₂│₁₁│₁₀│ ₉│ ₈│ ₇│ ₆│ ₅│ ₄│ ₃│ ₂│ ₁│ ₀┃
// ┡━━┿━━┿━━┿━━┿━━┿━━╇━━┿━━┿━━┿━━┿━━┿━━╇━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┿━━┩
// │₃₁│₃₀│₂₉│₂₈│₂₇│₂₆│₂₅│₂₄│₂₃│₂₂│₂₁│₂₀│₁₉│₁₈│₁₇│₁₆│₁₅│₁₄│₁₃│₁₂│₁₁│₁₀│ ₉│ ₈│ ₇│ ₆│ ₅│ ₄│ ₃│ ₂│ ₁│ ₀│
// └──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┴──┘
func packCYZ(c uint32, y uint32, z uint32) uint32 {
	if y > 0x3f {
		panic(fmt.Errorf("y exceeds 6 bits (0x%x)", y))
	}
	if z > 0xfffff {
		panic(fmt.Errorf("z exceeds 20 bits (0x%x)", z))
	}
	return packC(c) | (y << 20) | z
}

func unpackC(i uint32) uint32 { return i >> 26 }
func unpackX(i uint32) uint32 { return i & 0x3ffffff }
func unpackY(i uint32) uint32 { return (i >> 20) & 0x3f }
func unpackZ(i uint32) uint32 { return i & 0xfffff }

// Call represents the CALL virtual machine opcode.
type Call struct {
	PushReturn bool   // Should the return value be pushed onto the stack?
	FunctionID uint16 // The function identifier to call.
}

func (c Call) Encode(e binary.Encoder) error {
	return e.Uint32(packCX(protocol.OpCall, setBit(uint32(c.FunctionID), 24, c.PushReturn)))
}

// PushI represents the PUSH_I virtual machine opcode.
type PushI struct {
	DataType protocol.Type // The value type to push.
	Value    uint32        // The value to push packed into the low 20 bits.
}

func (c PushI) Encode(e binary.Encoder) error {
	return e.Uint32(packCYZ(protocol.OpPushI, uint32(c.DataType), c.Value))
}

// LoadC represents the LOAD_C virtual machine opcode.
type LoadC struct {
	DataType protocol.Type // The value type to load.
	Address  uint32        // The pointer to the value in constant address-space.
}

func (c LoadC) Encode(e binary.Encoder) error {
	return e.Uint32(packCYZ(protocol.OpLoadC, uint32(c.DataType), c.Address))
}

// LoadV represents the LOAD_V virtual machine opcode.
type LoadV struct {
	DataType protocol.Type // The value type to load.
	Address  uint32        // The pointer to the value in volatile address-space.
}

func (c LoadV) Encode(e binary.Encoder) error {
	return e.Uint32(packCYZ(protocol.OpLoadV, uint32(c.DataType), c.Address))
}

// Load represents the LOAD virtual machine opcode.
type Load struct {
	DataType protocol.Type // The value types to load.
}

func (c Load) Encode(e binary.Encoder) error {
	return e.Uint32(packCYZ(protocol.OpLoad, uint32(c.DataType), 0))
}

// Pop represents the POP virtual machine opcode.
type Pop struct {
	Count uint32 // Number of elements to pop from the top of the stack.
}

func (c Pop) Encode(e binary.Encoder) error {
	return e.Uint32(packCX(protocol.OpPop, c.Count))
}

// StoreV represents the STORE_V virtual machine opcode.
type StoreV struct {
	Address uint32 // Pointer in volatile address-space.
}

func (c StoreV) Encode(e binary.Encoder) error {
	return e.Uint32(packCX(protocol.OpStoreV, c.Address))
}

// Store represents the STORE virtual machine opcode.
type Store struct{}

func (c Store) Encode(e binary.Encoder) error {
	return e.Uint32(packC(protocol.OpStore))
}

// Resource represents the RESOURCE virtual machine opcode.
type Resource struct {
	ID uint32 // The index of the resource identifier.
}

func (c Resource) Encode(e binary.Encoder) error {
	return e.Uint32(packCX(protocol.OpResource, c.ID))
}

// Post represents the POST virtual machine opcode.
type Post struct{}

func (c Post) Encode(e binary.Encoder) error {
	return e.Uint32(packC(protocol.OpPost))
}

// Copy represents the COPY virtual machine opcode.
type Copy struct {
	Count uint32 // Number of bytes to copy.
}

func (c Copy) Encode(e binary.Encoder) error {
	return e.Uint32(packCX(protocol.OpCopy, c.Count))
}

// Clone represents the CLONE virtual machine opcode.
type Clone struct {
	Index uint32 // Index of element from top of stack to clone.
}

func (c Clone) Encode(e binary.Encoder) error {
	return e.Uint32(packCX(protocol.OpClone, c.Index))
}

// Strcpy represents the STRCPY virtual machine opcode.
type Strcpy struct {
	MaxSize uint32 // Maximum size in bytes to copy.
}

func (c Strcpy) Encode(e binary.Encoder) error {
	return e.Uint32(packCX(protocol.OpStrcpy, c.MaxSize))
}

// Extend represents the EXTEND virtual machine opcode.
type Extend struct {
	Value uint32 // 26 bit value to extend the top of the stack by.
}

func (c Extend) Encode(e binary.Encoder) error {
	return e.Uint32(packCX(protocol.OpExtend, c.Value))
}

// Extend represents the LABEL virtual machine opcode.
type Label struct {
	Value uint32 // 26 bit label name.
}

func (c Label) Encode(e binary.Encoder) error {
	return e.Uint32(packCX(protocol.OpLabel, c.Value))
}

// Decode returns the opcode decoded from decoder d.
func Decode(d binary.Decoder) (interface{}, error) {
	i, err := d.Uint32()
	if err != nil {
		return nil, err
	}
	code := unpackC(i)
	switch code {
	case protocol.OpCall:
		return Call{PushReturn: bit(i, 24), FunctionID: uint16(unpackX(i))}, nil
	case protocol.OpPushI:
		return PushI{DataType: protocol.Type(unpackY(i)), Value: unpackZ(i)}, nil
	case protocol.OpLoadC:
		return LoadC{DataType: protocol.Type(unpackY(i)), Address: unpackZ(i)}, nil
	case protocol.OpLoadV:
		return LoadV{DataType: protocol.Type(unpackY(i)), Address: unpackZ(i)}, nil
	case protocol.OpLoad:
		return Load{DataType: protocol.Type(unpackY(i))}, nil
	case protocol.OpPop:
		return Pop{Count: unpackX(i)}, nil
	case protocol.OpStoreV:
		return StoreV{Address: unpackX(i)}, nil
	case protocol.OpStore:
		return Store{}, nil
	case protocol.OpResource:
		return Resource{ID: unpackX(i)}, nil
	case protocol.OpPost:
		return Post{}, nil
	case protocol.OpCopy:
		return Copy{Count: unpackX(i)}, nil
	case protocol.OpClone:
		return Clone{Index: unpackX(i)}, nil
	case protocol.OpStrcpy:
		return Strcpy{MaxSize: unpackX(i)}, nil
	case protocol.OpExtend:
		return Extend{Value: unpackX(i)}, nil
	case protocol.OpLabel:
		return Label{Value: unpackX(i)}, nil
	default:
		return nil, fmt.Errorf("Unknown opcode with code %v", code)
	}
}
