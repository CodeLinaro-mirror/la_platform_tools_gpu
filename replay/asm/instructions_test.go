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
	"bytes"
	"testing"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/replay/opcode"
	"android.googlesource.com/platform/tools/gpu/replay/value"
	"android.googlesource.com/platform/tools/gpu/replay/vm"
)

type testPtrResolver struct{}

func (testPtrResolver) TranslateTemporaryPointer(ptr uint64) uint64 { return ptr }
func (testPtrResolver) TranslateCapturePointer(ptr uint64) uint64   { return ptr }

func check(t *testing.T, Instructions []Instruction, expected ...interface{}) {
	buf := &bytes.Buffer{}
	b := binary.NewEncoder(buf)
	for _, Instruction := range Instructions {
		Instruction.Encode(testPtrResolver{}, b)
	}
	gotOpcodes, err := opcode.Disassemble(buf)
	if err != nil {
		panic(err)
	}
	opcode.CheckDisassembly(t, gotOpcodes, expected...)
}

func TestCall(t *testing.T) {
	check(t,
		[]Instruction{
			Call{false, 0x1234},
			Call{true, 0x5678},
		},
		opcode.Call{PushReturn: false, FunctionID: 0x1234},
		opcode.Call{PushReturn: true, FunctionID: 0x5678},
	)
}

func TestPush_UnsignedNoExpand(t *testing.T) {
	check(t,
		[]Instruction{
			Push{value.U32(0xaaaaa)}, // Repeating pattern of 1010
			Push{value.U32(0x55555)}, // Repeating pattern of 0101
		},
		opcode.PushI{DataType: vm.TypeUint32, Value: 0xaaaaa},
		opcode.PushI{DataType: vm.TypeUint32, Value: 0x55555},
	)
}

func TestPush_UnsignedOneExpand(t *testing.T) {
	check(t,
		[]Instruction{
			Push{value.U32(0x100000)},   // One bit beyond what can fit in a PushI
			Push{value.U32(0x4000000)},  // One bit beyond what can fit in a Extend payload
			Push{value.U32(0xaaaaaaaa)}, // 1010101010...
			Push{value.U32(0x55555555)}, // 0101010101...
		},
		opcode.PushI{DataType: vm.TypeUint32, Value: 0},
		opcode.Extend{Value: 0x100000},

		opcode.PushI{DataType: vm.TypeUint32, Value: 1},
		opcode.Extend{Value: 0},

		opcode.PushI{DataType: vm.TypeUint32, Value: 0x2a},
		opcode.Extend{Value: 0x2aaaaaa},

		opcode.PushI{DataType: vm.TypeUint32, Value: 0x15},
		opcode.Extend{Value: 0x1555555},
	)
}

func TestPush_SignedPositiveNoExpand(t *testing.T) {
	check(t,
		[]Instruction{
			Push{value.S32(0x2aaaa)}, // 0010101010...
			Push{value.S32(0x55555)}, // 1010101010...
		},
		opcode.PushI{DataType: vm.TypeInt32, Value: 0x2aaaa},
		opcode.PushI{DataType: vm.TypeInt32, Value: 0x55555},
	)
}

func TestPush_SignedPositiveOneExpand(t *testing.T) {
	check(t,
		[]Instruction{
			Push{value.S32(0x80000)},    // One bit beyond what can fit in a PushI
			Push{value.S32(0x4000000)},  // One bit beyond what can fit in a Extend payload
			Push{value.S32(0x2aaaaaaa)}, // 0010101010...
			Push{value.S32(0x55555555)}, // 0101010101...
		},
		opcode.PushI{DataType: vm.TypeInt32, Value: 0},
		opcode.Extend{Value: 0x80000},

		opcode.PushI{DataType: vm.TypeInt32, Value: 1},
		opcode.Extend{Value: 0},

		opcode.PushI{DataType: vm.TypeInt32, Value: 0x0a},
		opcode.Extend{Value: 0x2aaaaaa},

		opcode.PushI{DataType: vm.TypeInt32, Value: 0x15},
		opcode.Extend{Value: 0x1555555},
	)
}

func TestPush_SignedNegativeNoExpand(t *testing.T) {
	check(t,
		[]Instruction{
			Push{value.S32(-1)},
			Push{value.S32(-0x55556)}, // Repeating pattern of 1010
			Push{value.S32(-0x2aaab)}, // Repeating pattern of 0101
		},
		opcode.PushI{DataType: vm.TypeInt32, Value: 0xfffff},
		opcode.PushI{DataType: vm.TypeInt32, Value: 0xaaaaa},
		opcode.PushI{DataType: vm.TypeInt32, Value: 0xd5555},
	)
}

func TestPush_SignedNegativeOneExpand(t *testing.T) {
	check(t,
		[]Instruction{
			Push{value.S32(-0x100001)},  // One bit beyond what can fit in a PushI
			Push{value.S32(-0x4000001)}, // One bit beyond what can fit in a Extend payload

			Push{value.S32(-0x2aaaaaab)}, // 110101010...
			Push{value.S32(-0x55555556)}, // 101010101...
		},
		opcode.PushI{DataType: vm.TypeInt32, Value: 0xfffff},
		opcode.Extend{Value: 0x03efffff},

		opcode.PushI{DataType: vm.TypeInt32, Value: 0xffffe},
		opcode.Extend{Value: 0x03ffffff},

		opcode.PushI{DataType: vm.TypeInt32, Value: 0xffff5},
		opcode.Extend{Value: 0x1555555},

		opcode.PushI{DataType: vm.TypeInt32, Value: 0xfffea},
		opcode.Extend{Value: 0x2aaaaaa},
	)
}

func TestPush_FloatNoExpand(t *testing.T) {
	check(t,
		[]Instruction{
			Push{value.F32(-2.0)},
			Push{value.F32(-1.0)},
			Push{value.F32(-0.5)},
			Push{value.F32(0)},
			Push{value.F32(0.5)},
			Push{value.F32(1.0)},
			Push{value.F32(2.0)},
		},
		opcode.PushI{DataType: vm.TypeFloat, Value: 0x180},
		opcode.PushI{DataType: vm.TypeFloat, Value: 0x17f},
		opcode.PushI{DataType: vm.TypeFloat, Value: 0x17e},
		opcode.PushI{DataType: vm.TypeFloat, Value: 0x000},
		opcode.PushI{DataType: vm.TypeFloat, Value: 0x07e},
		opcode.PushI{DataType: vm.TypeFloat, Value: 0x07f},
		opcode.PushI{DataType: vm.TypeFloat, Value: 0x080},
	)
}

func TestPush_FloatOneExpand(t *testing.T) {
	check(t,
		[]Instruction{
			Push{value.F32(-3)},
			Push{value.F32(-1.75)},
			Push{value.F32(1.75)},
			Push{value.F32(3)},
		},
		opcode.PushI{DataType: vm.TypeFloat, Value: 0x180},
		opcode.Extend{Value: 0x400000},

		opcode.PushI{DataType: vm.TypeFloat, Value: 0x17F},
		opcode.Extend{Value: 0x600000},

		opcode.PushI{DataType: vm.TypeFloat, Value: 0x07F},
		opcode.Extend{Value: 0x600000},

		opcode.PushI{DataType: vm.TypeFloat, Value: 0x080},
		opcode.Extend{Value: 0x400000},
	)
}

func TestPush_DoubleNoExpand(t *testing.T) {
	check(t,
		[]Instruction{
			Push{value.F64(-2.0)},
			Push{value.F64(-1.0)},
			Push{value.F64(-0.5)},
			Push{value.F64(0)},
			Push{value.F64(0.5)},
			Push{value.F64(1.0)},
			Push{value.F64(2.0)},
		},
		opcode.PushI{DataType: vm.TypeDouble, Value: 0xc00},
		opcode.PushI{DataType: vm.TypeDouble, Value: 0xbff},
		opcode.PushI{DataType: vm.TypeDouble, Value: 0xbfe},
		opcode.PushI{DataType: vm.TypeDouble, Value: 0x000},
		opcode.PushI{DataType: vm.TypeDouble, Value: 0x3fe},
		opcode.PushI{DataType: vm.TypeDouble, Value: 0x3ff},
		opcode.PushI{DataType: vm.TypeDouble, Value: 0x400},
	)
}

func TestPush_DoubleExpand(t *testing.T) {
	check(t,
		[]Instruction{
			Push{value.F64(-3)},
			Push{value.F64(-1.75)},
			Push{value.F64(1) / 3},
			Push{value.F64(1.75)},
			Push{value.F64(3)},
		},
		opcode.PushI{DataType: vm.TypeDouble, Value: 0xc00},
		opcode.Extend{Value: 0x2000000},
		opcode.Extend{Value: 0x0},

		opcode.PushI{DataType: vm.TypeDouble, Value: 0xbff},
		opcode.Extend{Value: 0x3000000},
		opcode.Extend{Value: 0x0},

		opcode.PushI{DataType: vm.TypeDouble, Value: 0x3fd},
		opcode.Extend{Value: 0x1555555},
		opcode.Extend{Value: 0x1555555},

		opcode.PushI{DataType: vm.TypeDouble, Value: 0x3ff},
		opcode.Extend{Value: 0x3000000},
		opcode.Extend{Value: 0x0},

		opcode.PushI{DataType: vm.TypeDouble, Value: 0x400},
		opcode.Extend{Value: 0x2000000},
		opcode.Extend{Value: 0x0},
	)
}
func TestPop(t *testing.T) {
	check(t,
		[]Instruction{
			Pop{10},
		},
		opcode.Pop{Count: 10},
	)
}

func TestCopy(t *testing.T) {
	check(t,
		[]Instruction{
			Copy{100},
		},
		opcode.Copy{Count: 100},
	)
}

func TestClone(t *testing.T) {
	check(t,
		[]Instruction{
			Clone{100},
		},
		opcode.Clone{Index: 100},
	)
}

func TestLoad(t *testing.T) {
	check(t,
		[]Instruction{
			Load{vm.TypeUint16, value.ConstantPointer(0x10)},
			Load{vm.TypeUint16, value.ConstantPointer(0x123456)},

			Load{vm.TypeUint16, value.VolatilePointer(0x10)},
			Load{vm.TypeUint16, value.VolatilePointer(0x123456)},
		},
		opcode.LoadC{DataType: vm.TypeUint16, Address: 0x10},

		opcode.PushI{DataType: vm.TypeConstantPointer, Value: 0},
		opcode.Extend{Value: 0x123456},
		opcode.Load{DataType: vm.TypeUint16},

		opcode.LoadV{DataType: vm.TypeUint16, Address: 0x10},

		opcode.PushI{DataType: vm.TypeVolatilePointer, Value: 0},
		opcode.Extend{Value: 0x123456},
		opcode.Load{DataType: vm.TypeUint16},
	)
}

func TestStore(t *testing.T) {
	check(t,
		[]Instruction{
			Store{value.VolatilePointer(0x10)},
			Store{value.VolatilePointer(0x4000000)},
		},
		opcode.StoreV{Address: 0x10},

		opcode.PushI{DataType: vm.TypeVolatilePointer, Value: 1},
		opcode.Extend{Value: 0},
		opcode.Store{},
	)
}

func TestStrcpy(t *testing.T) {
	check(t,
		[]Instruction{
			Strcpy{0x3000},
		},
		opcode.Strcpy{MaxSize: 0x3000},
	)
}

func TestResource(t *testing.T) {
	check(t,
		[]Instruction{
			Resource{10, 0x10},
			Resource{20, 0x4050607},
		},
		opcode.PushI{DataType: vm.TypeVolatilePointer, Value: 0x10},
		opcode.Resource{10},

		opcode.PushI{DataType: vm.TypeVolatilePointer, Value: 1},
		opcode.Extend{Value: 0x50607},
		opcode.Resource{20},
	)
}

func TestPost(t *testing.T) {
	check(t,
		[]Instruction{
			Post{value.AbsolutePointer(0x10), 0x50},
			Post{value.VolatilePointer(0x4050607), 0x8090a0b},
		},
		opcode.PushI{DataType: vm.TypeAbsolutePointer, Value: 0x10},
		opcode.PushI{DataType: vm.TypeUint32, Value: 0x50},
		opcode.Post{},

		opcode.PushI{DataType: vm.TypeVolatilePointer, Value: 1},
		opcode.Extend{Value: 0x50607},
		opcode.PushI{DataType: vm.TypeUint32, Value: 2},
		opcode.Extend{Value: 0x90a0b},
		opcode.Post{},
	)
}
