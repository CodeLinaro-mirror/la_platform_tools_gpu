/*
 * Copyright 2015, The Android Open Source Project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package opcode

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/replay/vm"
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
		panic("c exceeds 6 bits")
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
		panic("x exceeds 26 bits")
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
		panic("y exceeds 6 bits")
	}
	if z > 0xfffff {
		panic("z exceeds 20 bits")
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
	FunctionId uint16 // The function identifier to call.
}

func (c Call) Encode(e *binary.Encoder) error {
	return e.Uint32(packCX(vm.OpCall, setBit(uint32(c.FunctionId), 24, c.PushReturn)))
}

// PushI represents the PUSH_I virtual machine opcode.
type PushI struct {
	DataType vm.Type // The value type to push.
	Value    uint32  // The value to push packed into the low 20 bits.
}

func (c PushI) Encode(e *binary.Encoder) error {
	return e.Uint32(packCYZ(vm.OpPushI, uint32(c.DataType), c.Value))
}

// LoadC represents the LOAD_C virtual machine opcode.
type LoadC struct {
	DataType vm.Type // The value type to load.
	Address  uint32  // The pointer to the value in constant address-space.
}

func (c LoadC) Encode(e *binary.Encoder) error {
	return e.Uint32(packCYZ(vm.OpLoadC, uint32(c.DataType), c.Address))
}

// LoadV represents the LOAD_V virtual machine opcode.
type LoadV struct {
	DataType vm.Type // The value type to load.
	Address  uint32  // The pointer to the value in volatile address-space.
}

func (c LoadV) Encode(e *binary.Encoder) error {
	return e.Uint32(packCYZ(vm.OpLoadV, uint32(c.DataType), c.Address))
}

// Load represents the LOAD virtual machine opcode.
type Load struct {
	DataType vm.Type // The value types to load.
}

func (c Load) Encode(e *binary.Encoder) error {
	return e.Uint32(packCYZ(vm.OpLoad, uint32(c.DataType), 0))
}

// Pop represents the POP virtual machine opcode.
type Pop struct {
	Count uint32 // Number of elements to pop from the top of the stack.
}

func (c Pop) Encode(e *binary.Encoder) error {
	return e.Uint32(packCX(vm.OpPop, c.Count))
}

// StoreV represents the STORE_V virtual machine opcode.
type StoreV struct {
	Address uint32 // Pointer in volatile address-space.
}

func (c StoreV) Encode(e *binary.Encoder) error {
	return e.Uint32(packCX(vm.OpStoreV, c.Address))
}

// Store represents the STORE virtual machine opcode.
type Store struct{}

func (c Store) Encode(e *binary.Encoder) error {
	return e.Uint32(packC(vm.OpStore))
}

// Resource represents the RESOURCE virtual machine opcode.
type Resource struct {
	Id uint32 // The index of the resource identifier.
}

func (c Resource) Encode(e *binary.Encoder) error {
	return e.Uint32(packCX(vm.OpResource, c.Id))
}

// Post represents the POST virtual machine opcode.
type Post struct{}

func (c Post) Encode(e *binary.Encoder) error {
	return e.Uint32(packC(vm.OpPost))
}

// Copy represents the COPY virtual machine opcode.
type Copy struct {
	Count uint32 // Number of bytes to copy.
}

func (c Copy) Encode(e *binary.Encoder) error {
	return e.Uint32(packCX(vm.OpCopy, c.Count))
}

// Clone represents the CLONE virtual machine opcode.
type Clone struct {
	Index uint32 // Index of element from top of stack to clone.
}

func (c Clone) Encode(e *binary.Encoder) error {
	return e.Uint32(packCX(vm.OpClone, c.Index))
}

// Strcpy represents the STRCPY virtual machine opcode.
type Strcpy struct {
	MaxSize uint32 // Maximum size in bytes to copy.
}

func (c Strcpy) Encode(e *binary.Encoder) error {
	return e.Uint32(packCX(vm.OpStrcpy, c.MaxSize))
}

// Extend represents the EXTEND virtual machine opcode.
type Extend struct {
	Value uint32 // 26 bit value to extend the top of the stack by.
}

func (c Extend) Encode(e *binary.Encoder) error {
	return e.Uint32(packCX(vm.OpExtend, c.Value))
}

// Decode returns the opcode decoded from decoder d.
func Decode(d *binary.Decoder) (interface{}, error) {
	i, err := d.Uint32()
	if err != nil {
		return nil, err
	}
	code := unpackC(i)
	switch code {
	case vm.OpCall:
		return Call{PushReturn: bit(i, 24), FunctionId: uint16(unpackX(i))}, nil
	case vm.OpPushI:
		return PushI{DataType: vm.Type(unpackY(i)), Value: unpackZ(i)}, nil
	case vm.OpLoadC:
		return LoadC{DataType: vm.Type(unpackY(i)), Address: unpackZ(i)}, nil
	case vm.OpLoadV:
		return LoadV{DataType: vm.Type(unpackY(i)), Address: unpackZ(i)}, nil
	case vm.OpLoad:
		return Load{DataType: vm.Type(unpackY(i))}, nil
	case vm.OpPop:
		return Pop{Count: unpackX(i)}, nil
	case vm.OpStoreV:
		return StoreV{Address: unpackX(i)}, nil
	case vm.OpStore:
		return Store{}, nil
	case vm.OpResource:
		return Resource{Id: unpackX(i)}, nil
	case vm.OpPost:
		return Post{}, nil
	case vm.OpCopy:
		return Copy{Count: unpackX(i)}, nil
	case vm.OpClone:
		return Clone{Index: unpackX(i)}, nil
	case vm.OpStrcpy:
		return Strcpy{MaxSize: unpackX(i)}, nil
	case vm.OpExtend:
		return Extend{Value: unpackX(i)}, nil
	default:
		return nil, fmt.Errorf("Unknown opcode with code %v", code)
	}
}
