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

package asm

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/replay/opcode"
	"android.googlesource.com/platform/tools/gpu/replay/protocol"
	"android.googlesource.com/platform/tools/gpu/replay/value"
)

// Instruction is the interface of all instruction types.
//
// Encode writes the instruction's opcodes to the binary encoder e, translating
// any pointers to their final, resolved addresses using the PointerResolver r.
// An instruction can produce zero, one or many opcodes.
type Instruction interface {
	Encode(r value.PointerResolver, e *binary.Encoder) error
}

func encodePush(t protocol.Type, v uint64, e *binary.Encoder) error {
	mask19 := uint64(0x7ffff)
	mask20 := uint64(0xfffff)
	mask26 := uint64(0x3ffffff)
	mask45 := uint64(0x1fffffffffff)
	mask46 := uint64(0x3fffffffffff)
	mask52 := uint64(0xfffffffffffff)

	//     ▏60       ▏50       ▏40       ▏30       ▏20       ▏10
	// ○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○●●●●●●●●●●●●●●●●●●● mask19
	// ○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○●●●●●●●●●●●●●●●●●●●● mask20
	// ○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○●●●●●●●●●●●●●●●●●●●●●●●●●● mask26
	// ○○○○○○○○○○○○○○○○○○○●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●● mask45
	// ○○○○○○○○○○○○○○○○○○●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●● mask46
	// ○○○○○○○○○○○○●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●● mask52
	//                                            ▕      PUSHI 20     ▕
	//                                      ▕         EXTEND 26       ▕

	switch t {
	case protocol.TypeFloat:
		push := opcode.PushI{DataType: t, Value: uint32(v >> 23)}
		if err := push.Encode(e); err != nil {
			return err
		}
		if v&0x7fffff != 0 {
			return opcode.Extend{Value: uint32(v & 0x7fffff)}.Encode(e)
		}
	case protocol.TypeDouble:
		push := opcode.PushI{DataType: t, Value: uint32(v >> 52)}
		if err := push.Encode(e); err != nil {
			return err
		}
		v &= mask52
		if v != 0 {
			ext := opcode.Extend{Value: uint32(v >> 26)}
			if err := ext.Encode(e); err != nil {
				return err
			}
			return opcode.Extend{Value: uint32(v & mask26)}.Encode(e)
		}
	case protocol.TypeInt8, protocol.TypeInt16, protocol.TypeInt32, protocol.TypeInt64:
		// Signed PUSHI types are sign-extended
		switch {
		case v&^mask19 == 0:
			// ○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒
			//                                            ▕      PUSHI 20     ▕
			return opcode.PushI{DataType: t, Value: uint32(v)}.Encode(e)
		case v&^mask19 == ^mask19:
			// ●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒
			//                                            ▕      PUSHI 20     ▕
			return opcode.PushI{DataType: t, Value: uint32(v & mask20)}.Encode(e)
		case v&^mask45 == 0:
			// ○○○○○○○○○○○○○○○○○○○◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒
			//                  ▕      PUSHI 20     ▕         EXTEND 26       ▕
			push := opcode.PushI{DataType: t, Value: uint32(v >> 26)}
			if err := push.Encode(e); err != nil {
				return err
			}
			return opcode.Extend{uint32(v & mask26)}.Encode(e)
		case v&^mask45 == ^mask45:
			// ●●●●●●●●●●●●●●●●●●●◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒
			//                  ▕      PUSHI 20     ▕         EXTEND 26       ▕
			push := opcode.PushI{DataType: t, Value: uint32((v >> 26) & mask20)}
			if err := push.Encode(e); err != nil {
				return err
			}
			return opcode.Extend{uint32(v & mask26)}.Encode(e)
		default:
			// ◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒
			//▕  PUSHI 12 ▕         EXTEND 26       ▕         EXTEND 26       ▕
			push := opcode.PushI{DataType: t, Value: uint32(v >> 52)}
			if err := push.Encode(e); err != nil {
				return err
			}
			ext := opcode.Extend{uint32((v >> 26) & mask26)}
			if err := ext.Encode(e); err != nil {
				return err
			}
			return opcode.Extend{uint32(v & mask26)}.Encode(e)
		}
	case protocol.TypeBool,
		protocol.TypeUint8, protocol.TypeUint16, protocol.TypeUint32, protocol.TypeUint64,
		protocol.TypeAbsolutePointer, protocol.TypeConstantPointer, protocol.TypeVolatilePointer:
		switch {
		case v&^mask20 == 0:
			// ○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒
			//                                            ▕      PUSHI 20     ▕
			return opcode.PushI{DataType: t, Value: uint32(v)}.Encode(e)
		case v&^mask46 == 0:
			// ○○○○○○○○○○○○○○○○○○◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒
			//                  ▕      PUSHI 20     ▕         EXTEND 26       ▕
			push := opcode.PushI{DataType: t, Value: uint32(v >> 26)}
			if err := push.Encode(e); err != nil {
				return err
			}
			return opcode.Extend{uint32(v & mask26)}.Encode(e)
		default:
			// ◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒◒
			//▕  PUSHI 12 ▕         EXTEND 26       ▕         EXTEND 26       ▕
			push := opcode.PushI{DataType: t, Value: uint32(v >> 52)}
			if err := push.Encode(e); err != nil {
				return err
			}
			ext := opcode.Extend{uint32((v >> 26) & mask26)}
			if err := ext.Encode(e); err != nil {
				return err
			}
			return opcode.Extend{uint32(v & mask26)}.Encode(e)
		}
	}
	return fmt.Errorf("Cannot push value type %s", t.String())
}

// Nop is a no-operation Instruction. Instructions of this type do nothing.
type Nop struct{}

func (Nop) Encode(r value.PointerResolver, e *binary.Encoder) error {
	return nil
}

// Call is an Instruction to call a VM registered function.
// This instruction will pop the parameters from the VM stack starting with the
// first parameter. If PushReturn is true, then the return value of the function
// call will be pushed to the top of the VM stack.
type Call struct {
	PushReturn bool   // If true, the return value is pushed to the VM stack.
	FunctionID uint16 // The function id registered with the VM to invoke.
}

func (a Call) Encode(r value.PointerResolver, e *binary.Encoder) error {
	return opcode.Call{
		PushReturn: a.PushReturn,
		FunctionID: a.FunctionID,
	}.Encode(e)
}

// Push is an Instruction to push Value to the top of the VM stack.
type Push struct {
	Value value.Value // The value to push on to the VM stack.
}

func (a Push) Encode(r value.PointerResolver, e *binary.Encoder) error {
	return encodePush(a.Value.Type(), a.Value.Get(r), e)
}

// Pop is an Instruction that discards Count values from the top of the VM
// stack.
type Pop struct {
	Count uint32 // Number of values to discard from the top of the VM stack.
}

func (a Pop) Encode(r value.PointerResolver, e *binary.Encoder) error {
	return opcode.Pop{Count: a.Count}.Encode(e)
}

// Copy is an Instruction that pops the source address and then the target
// address from the top of the VM stack, and then copies Count bytes from
// source to target.
type Copy struct {
	Count uint64 // Number of bytes to copy.
}

func (a Copy) Encode(r value.PointerResolver, e *binary.Encoder) error {
	return opcode.Copy{Count: uint32(a.Count)}.Encode(e)
}

// Clone is an Instruction that makes a copy of the the n-th element from the
// top of the VM stack and pushes the copy to the top of the VM stack.
type Clone struct {
	Index int
}

func (a Clone) Encode(r value.PointerResolver, e *binary.Encoder) error {
	return opcode.Clone{Index: uint32(a.Index)}.Encode(e)
}

// Load is an Instruction that loads the value of type DataType from pointer
// Source and pushes the loaded value to the top of the VM stack.
type Load struct {
	DataType protocol.Type
	Source   value.Pointer
}

func (a Load) Encode(r value.PointerResolver, e *binary.Encoder) error {
	addr := a.Source.Get(r)
	switch a.Source.(type) {
	case value.ConstantPointer:
		if addr < 0x100000 {
			return opcode.LoadC{DataType: a.DataType, Address: uint32(addr)}.Encode(e)
		} else {
			if err := encodePush(a.Source.Type(), a.Source.Get(r), e); err != nil {
				return err
			}
			return opcode.Load{DataType: a.DataType}.Encode(e)
		}
	case value.VolatilePointer, value.VolatileTemporaryPointer, value.VolatileCapturePointer:
		if addr < 0x100000 {
			return opcode.LoadV{DataType: a.DataType, Address: uint32(addr)}.Encode(e)
		} else {
			if err := encodePush(a.Source.Type(), a.Source.Get(r), e); err != nil {
				return err
			}
			return opcode.Load{DataType: a.DataType}.Encode(e)
		}
	default:
		return fmt.Errorf("Unsupported load source type %T", a.Source)
	}
}

// Store is an Instruction that pops the value from the top of the VM stack and
// writes the value to Destination.
type Store struct {
	Destination value.Pointer
}

func (a Store) Encode(r value.PointerResolver, e *binary.Encoder) error {
	addr := a.Destination.Get(r)
	if addr < 0x3ffffff {
		return opcode.StoreV{Address: uint32(addr)}.Encode(e)
	} else {
		if err := encodePush(a.Destination.Type(), a.Destination.Get(r), e); err != nil {
			return err
		}
		return opcode.Store{}.Encode(e)
	}
}

// Strcpy is an Instruction that pops the source address then the target address
// from the top of the VM stack, and then copies at most MaxCount-1 bytes from
// source to target. If the MaxCount is greater than the source string length,
// then the target will be padded with 0s. The destination buffer will always be
// 0-terminated.
type Strcpy struct {
	MaxCount uint64
}

func (a Strcpy) Encode(r value.PointerResolver, e *binary.Encoder) error {
	return opcode.Strcpy{
		MaxSize: uint32(a.MaxCount),
	}.Encode(e)
}

// Resource is an Instruction that loads the resource with index Index of Size
// bytes and writes the resource to Destination.
type Resource struct {
	Index       uint32
	Destination memory.Pointer
}

func (a Resource) Encode(r value.PointerResolver, e *binary.Encoder) error {
	ptr := value.VolatileCapturePointer(a.Destination)
	if err := encodePush(ptr.Type(), ptr.Get(r), e); err != nil {
		return err
	}
	return opcode.Resource{
		ID: a.Index,
	}.Encode(e)
}

// Post is an Instruction that posts Size bytes from Source to the server.
type Post struct {
	Source value.Pointer
	Size   uint64
}

func (a Post) Encode(r value.PointerResolver, e *binary.Encoder) error {
	if err := encodePush(a.Source.Type(), a.Source.Get(r), e); err != nil {
		return err
	}
	if err := encodePush(protocol.TypeUint32, a.Size, e); err != nil {
		return err
	}
	return opcode.Post{}.Encode(e)
}
