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

package test

import (
	"bytes"
	"encoding/binary"
	"testing"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/replay/builder"
	"android.googlesource.com/platform/tools/gpu/replay/opcode"
	"android.googlesource.com/platform/tools/gpu/replay/protocol"
)

func check(t *testing.T, ptrSize, ptrAlignment int, wantOutput bool, atoms []atom.Atom, opcodes []interface{}, constants []byte) {
	b := builder.New(ptrSize, ptrAlignment, binary.LittleEndian)
	r := newReplayWriter(b)
	for _, atom := range atoms {
		r.Write(0, atom, wantOutput)
	}

	payload, _ := b.Build(log.Nop{})

	ops := bytes.NewBuffer(payload.Opcodes.Data)
	gotOpcodes, err := opcode.Disassemble(ops, binary.LittleEndian)
	if err != nil {
		t.Errorf("Failed to disassemble opcodes: %v", err)
	}
	opcode.CheckDisassembly(t, gotOpcodes, opcodes...)

	gotConstants := payload.Constants
	if !bytes.Equal(gotConstants.Data, constants) {
		t.Errorf("Constant buffer was not as expected.\nGot:      % .2x\nExpected: % .2x", gotConstants, constants)
	}
}

func TestOperationsOpCall_NoIn_NoOut(t *testing.T) {
	check(t, 4 /* pointer size */, 4 /* pointer alignment */, false /* wantOutput */, []atom.Atom{
		NewCmdVoid(),
	}, []interface{}{
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoid.ID},
	}, []byte{})
}

func TestOperationsOpCall_SingleInputArg(t *testing.T) {
	check(t, 4 /* pointer size */, 4 /* pointer alignment */, false /* wantOutput */, []atom.Atom{
		NewCmdVoidU8(20),
		NewCmdVoidS8(-20),
		NewCmdVoidU16(200),
		NewCmdVoidS16(-200),
		NewCmdVoidF32(1.0),
		NewCmdVoidU32(2000),
		NewCmdVoidS32(-2000),
		NewCmdVoidF64(1.0),
		NewCmdVoidU64(20000),
		NewCmdVoidS64(-20000),
		NewCmdVoidBool(true),
		NewCmdVoidString("hello"),
	}, []interface{}{
		opcode.PushI{DataType: protocol.TypeUint8, Value: 20},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidU8.ID},

		opcode.PushI{DataType: protocol.TypeInt8, Value: 0xfffec},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidS8.ID},

		opcode.PushI{DataType: protocol.TypeUint16, Value: 200},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidU16.ID},

		opcode.PushI{DataType: protocol.TypeInt16, Value: 0xfff38},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidS16.ID},

		opcode.PushI{DataType: protocol.TypeFloat, Value: 0x7f},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidF32.ID},

		opcode.PushI{DataType: protocol.TypeUint32, Value: 2000},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidU32.ID},

		opcode.PushI{DataType: protocol.TypeInt32, Value: 0xff830},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidS32.ID},

		opcode.PushI{DataType: protocol.TypeDouble, Value: 0x3ff},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidF64.ID},

		opcode.PushI{DataType: protocol.TypeUint64, Value: 20000},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidU64.ID},

		opcode.PushI{DataType: protocol.TypeInt64, Value: 0xfb1e0},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidS64.ID},

		opcode.PushI{DataType: protocol.TypeBool, Value: 1},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidBool.ID},

		opcode.PushI{DataType: protocol.TypeConstantPointer, Value: 0x00},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidString.ID},
	}, []byte{'h', 'e', 'l', 'l', 'o', 0})
}

func TestOperationsOpCall_3_Strings(t *testing.T) {
	check(t, 4 /* pointer size */, 4 /* pointer alignment */, false /* wantOutput */, []atom.Atom{
		NewCmdVoid3Strings("hello", "world", "hello"),
	}, []interface{}{
		opcode.PushI{DataType: protocol.TypeConstantPointer, Value: 0x00},
		opcode.PushI{DataType: protocol.TypeConstantPointer, Value: 0x08},
		opcode.PushI{DataType: protocol.TypeConstantPointer, Value: 0x00},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoid3Strings.ID},
	}, []byte{
		/* 0x00 */ 'h', 'e', 'l', 'l', 'o', 0x00, 0x00, 0x00,
		/* 0x08 */ 'w', 'o', 'r', 'l', 'd', 0x00,
	})
}

func TestOperationsOpCall_3_Arrays(t *testing.T) {
	check(t, 8 /* pointer size */, 8 /* pointer alignment */, false /* wantOutput */, []atom.Atom{
		NewCmdVoid3Arrays(S8Array{1, 2, 3}, StringArray{"hello", "world", ":D"}, BoolArray{true, false, true}),
	}, []interface{}{
		opcode.PushI{DataType: protocol.TypeConstantPointer, Value: 0x00}, // a

		opcode.PushI{DataType: protocol.TypeConstantPointer, Value: 0x08},
		opcode.PushI{DataType: protocol.TypeConstantPointer, Value: 0x10},
		opcode.PushI{DataType: protocol.TypeConstantPointer, Value: 0x18},
		opcode.StoreV{Address: 0x10},
		opcode.StoreV{Address: 0x08},
		opcode.StoreV{Address: 0x00},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x00},

		opcode.PushI{DataType: protocol.TypeConstantPointer, Value: 0x20},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoid3Arrays.ID},
	}, []byte{
		/* 0x00 */ 0x01, 0x02, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00,
		/* 0x08 */ 'h', 'e', 'l', 'l', 'o', 0x00, 0x00, 0x00,
		/* 0x10 */ 'w', 'o', 'r', 'l', 'd', 0x00, 0x00, 0x00,
		/* 0x18 */ ':', 'D', 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		/* 0x20 */ 0x01, 0x00, 0x01,
	})
}

func TestOperationsOpCall_ArrayOfStrings_32bitOS(t *testing.T) {
	check(t, 4 /* pointer size */, 4 /* pointer alignment */, false /* wantOutput */, []atom.Atom{
		NewCmdVoidArrayOfStrings(StringArray{"an", "array", "of", "strings"}),
	}, []interface{}{
		opcode.PushI{DataType: protocol.TypeConstantPointer, Value: 0x00},
		opcode.PushI{DataType: protocol.TypeConstantPointer, Value: 0x04},
		opcode.PushI{DataType: protocol.TypeConstantPointer, Value: 0x0c},
		opcode.PushI{DataType: protocol.TypeConstantPointer, Value: 0x10},
		opcode.StoreV{Address: 0x0c},
		opcode.StoreV{Address: 0x08},
		opcode.StoreV{Address: 0x04},
		opcode.StoreV{Address: 0x00},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidArrayOfStrings.ID},
	}, []byte{
		/* 0x00 */ 'a', 'n', 0x00, 0x00,
		/* 0x04 */ 'a', 'r', 'r', 'a', 'y', 0x00, 0x00, 0x00,
		/* 0x0c */ 'o', 'f', 0x00, 0x00,
		/* 0x10 */ 's', 't', 'r', 'i', 'n', 'g', 's', 0x00,
	})
}

func TestOperationsOpCall_ArrayOfStrings_64bitOS(t *testing.T) {
	check(t, 8 /* pointer size */, 8 /* pointer alignment */, false /* wantOutput */, []atom.Atom{
		NewCmdVoidArrayOfStrings(StringArray{"an", "array", "of", "strings"}),
	}, []interface{}{
		opcode.PushI{DataType: protocol.TypeConstantPointer, Value: 0x00},
		opcode.PushI{DataType: protocol.TypeConstantPointer, Value: 0x08},
		opcode.PushI{DataType: protocol.TypeConstantPointer, Value: 0x10},
		opcode.PushI{DataType: protocol.TypeConstantPointer, Value: 0x18},
		opcode.StoreV{Address: 0x18},
		opcode.StoreV{Address: 0x10},
		opcode.StoreV{Address: 0x08},
		opcode.StoreV{Address: 0x00},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidArrayOfStrings.ID},
	}, []byte{
		/* 0x00 */ 'a', 'n', 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		/* 0x08 */ 'a', 'r', 'r', 'a', 'y', 0x00, 0x00, 0x00,
		/* 0x10 */ 'o', 'f', 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		/* 0x18 */ 's', 't', 'r', 'i', 'n', 'g', 's', 0x00,
	})
}

func TestOperationsOpCall_ReturnValue_DontWantOutput(t *testing.T) {
	check(t, 4 /* pointer size */, 4 /* pointer alignment */, false /* wantOutput */, []atom.Atom{
		NewCmdU8(20),
		NewCmdS8(-20),
		NewCmdU16(200),
		NewCmdS16(-200),
		NewCmdF32(1.0),
		NewCmdU32(2000),
		NewCmdS32(-2000),
		NewCmdF64(1.0),
		NewCmdU64(20000),
		NewCmdS64(-20000),
		NewCmdBool(true),
	}, []interface{}{
		opcode.Call{PushReturn: true, FunctionID: funcInfoCmdU8.ID}, opcode.StoreV{Address: 0},
		opcode.Call{PushReturn: true, FunctionID: funcInfoCmdS8.ID}, opcode.StoreV{Address: 0},
		opcode.Call{PushReturn: true, FunctionID: funcInfoCmdU16.ID}, opcode.StoreV{Address: 0},
		opcode.Call{PushReturn: true, FunctionID: funcInfoCmdS16.ID}, opcode.StoreV{Address: 0},
		opcode.Call{PushReturn: true, FunctionID: funcInfoCmdF32.ID}, opcode.StoreV{Address: 0},
		opcode.Call{PushReturn: true, FunctionID: funcInfoCmdU32.ID}, opcode.StoreV{Address: 0},
		opcode.Call{PushReturn: true, FunctionID: funcInfoCmdS32.ID}, opcode.StoreV{Address: 0},
		opcode.Call{PushReturn: true, FunctionID: funcInfoCmdF64.ID}, opcode.StoreV{Address: 0},
		opcode.Call{PushReturn: true, FunctionID: funcInfoCmdU64.ID}, opcode.StoreV{Address: 0},
		opcode.Call{PushReturn: true, FunctionID: funcInfoCmdS64.ID}, opcode.StoreV{Address: 0},
		opcode.Call{PushReturn: true, FunctionID: funcInfoCmdBool.ID}, opcode.StoreV{Address: 0},
	}, []byte{})
}

func TestOperationsOpCall_ReturnValue_WantOutput(t *testing.T) {
	check(t, 4 /* pointer size */, 4 /* pointer alignment */, true /* wantOutput */, []atom.Atom{
		NewCmdU8(20),
		NewCmdS8(-20),
		NewCmdU16(200),
		NewCmdS16(-200),
		NewCmdF32(1.0),
		NewCmdU32(2000),
		NewCmdS32(-2000),
		NewCmdF64(1.0),
		NewCmdU64(20000),
		NewCmdS64(-20000),
		NewCmdBool(true),
	}, []interface{}{
		opcode.Call{PushReturn: true, FunctionID: funcInfoCmdU8.ID}, opcode.StoreV{Address: 0},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.PushI{DataType: protocol.TypeUint32, Value: 1},
		opcode.Post{},

		opcode.Call{PushReturn: true, FunctionID: funcInfoCmdS8.ID}, opcode.StoreV{Address: 0},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.PushI{DataType: protocol.TypeUint32, Value: 1},
		opcode.Post{},

		opcode.Call{PushReturn: true, FunctionID: funcInfoCmdU16.ID}, opcode.StoreV{Address: 0},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.PushI{DataType: protocol.TypeUint32, Value: 2},
		opcode.Post{},

		opcode.Call{PushReturn: true, FunctionID: funcInfoCmdS16.ID}, opcode.StoreV{Address: 0},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.PushI{DataType: protocol.TypeUint32, Value: 2},
		opcode.Post{},

		opcode.Call{PushReturn: true, FunctionID: funcInfoCmdF32.ID}, opcode.StoreV{Address: 0},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.PushI{DataType: protocol.TypeUint32, Value: 4},
		opcode.Post{},

		opcode.Call{PushReturn: true, FunctionID: funcInfoCmdU32.ID}, opcode.StoreV{Address: 0},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.PushI{DataType: protocol.TypeUint32, Value: 4},
		opcode.Post{},

		opcode.Call{PushReturn: true, FunctionID: funcInfoCmdS32.ID}, opcode.StoreV{Address: 0},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.PushI{DataType: protocol.TypeUint32, Value: 4},
		opcode.Post{},

		opcode.Call{PushReturn: true, FunctionID: funcInfoCmdF64.ID}, opcode.StoreV{Address: 0},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.PushI{DataType: protocol.TypeUint32, Value: 8},
		opcode.Post{},

		opcode.Call{PushReturn: true, FunctionID: funcInfoCmdU64.ID}, opcode.StoreV{Address: 0},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.PushI{DataType: protocol.TypeUint32, Value: 8},
		opcode.Post{},

		opcode.Call{PushReturn: true, FunctionID: funcInfoCmdS64.ID}, opcode.StoreV{Address: 0},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.PushI{DataType: protocol.TypeUint32, Value: 8},
		opcode.Post{},

		opcode.Call{PushReturn: true, FunctionID: funcInfoCmdBool.ID}, opcode.StoreV{Address: 0},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.PushI{DataType: protocol.TypeUint32, Value: 1},
		opcode.Post{},
	}, []byte{})
}

func TestOperationsOpCall_ReturnValueString_DontWantOutput(t *testing.T) {
	check(t, 4 /* pointer size */, 4 /* pointer alignment */, false /* wantOutput */, []atom.Atom{
		NewCmdString("hello"),
	}, []interface{}{
		opcode.Call{PushReturn: true, FunctionID: funcInfoCmdString.ID},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x0},
		opcode.Strcpy{MaxSize: 10},
	}, []byte{})
}

func TestOperationsOpCall_ReturnValueString_WantOutput(t *testing.T) {
	check(t, 4 /* pointer size */, 4 /* pointer alignment */, true /* wantOutput */, []atom.Atom{
		NewCmdString("hello"),
	}, []interface{}{
		opcode.Call{PushReturn: true, FunctionID: funcInfoCmdString.ID},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x0},
		opcode.Strcpy{MaxSize: 10},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.PushI{DataType: protocol.TypeUint32, Value: 10},
		opcode.Post{},
	}, []byte{})
}

func TestOperationsOpCall_ReturnValueArray_DontWantOutput(t *testing.T) {
	check(t, 4 /* pointer size */, 4 /* pointer alignment */, false /* wantOutput */, []atom.Atom{
		NewCmdArrayOfFloat(F32Array{1, 2, 3}),
	}, []interface{}{
		opcode.Call{PushReturn: true, FunctionID: funcInfoCmdArrayOfFloat.ID},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x0},
		opcode.Copy{Count: 40},
	}, []byte{})
}

func TestOperationsOpCall_ReturnValueArray_WantOutput(t *testing.T) {
	check(t, 4 /* pointer size */, 4 /* pointer alignment */, true /* wantOutput */, []atom.Atom{
		NewCmdArrayOfFloat(F32Array{1, 2, 3}),
	}, []interface{}{
		opcode.Call{PushReturn: true, FunctionID: funcInfoCmdArrayOfFloat.ID},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x0},
		opcode.Copy{Count: 40},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.PushI{DataType: protocol.TypeUint32, Value: 40},
		opcode.Post{},
	}, []byte{})
}

func TestOperationsOpCall_ReturnValuePointer_DontWantOutput(t *testing.T) {
	check(t, 4 /* pointer size */, 4 /* pointer alignment */, false /* wantOutput */, []atom.Atom{
		NewCmdPointer(0x100),
	}, []interface{}{
		opcode.Call{PushReturn: true, FunctionID: funcInfoCmdPointer.ID},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x0},
		opcode.Copy{Count: 10},
	}, []byte{})
}

func TestOperationsOpCall_ReturnValuePointer_WantOutput(t *testing.T) {
	check(t, 4 /* pointer size */, 4 /* pointer alignment */, true /* wantOutput */, []atom.Atom{
		NewCmdPointer(0x100),
	}, []interface{}{
		opcode.Call{PushReturn: true, FunctionID: funcInfoCmdPointer.ID},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x0},
		opcode.Copy{Count: 10},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.PushI{DataType: protocol.TypeUint32, Value: 10},
		opcode.Post{},
	}, []byte{})
}

func TestOperationsOpCall_SingleOutputArg_DontWantOutput(t *testing.T) {
	check(t, 4 /* pointer size */, 4 /* pointer alignment */, false /* wantOutput */, []atom.Atom{
		NewCmdVoidOutU8(20),
		NewCmdVoidOutS8(-20),
		NewCmdVoidOutU16(200),
		NewCmdVoidOutS16(-200),
		NewCmdVoidOutF32(1.0),
		NewCmdVoidOutU32(2000),
		NewCmdVoidOutS32(-2000),
		NewCmdVoidOutF64(1.0),
		NewCmdVoidOutU64(20000),
		NewCmdVoidOutS64(-20000),
		NewCmdVoidOutBool(true),
		NewCmdVoidOutString("hello"),
	}, []interface{}{
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidOutU8.ID},

		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidOutS8.ID},

		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidOutU16.ID},

		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidOutS16.ID},

		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidOutF32.ID},

		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidOutU32.ID},

		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidOutS32.ID},

		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidOutF64.ID},

		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidOutU64.ID},

		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidOutS64.ID},

		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidOutBool.ID},

		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidOutString.ID},
	}, []byte{})
}

func TestOperationsOpCall_SingleOutputArg_WantOutput(t *testing.T) {
	check(t, 4 /* pointer size */, 4 /* pointer alignment */, true /* wantOutput */, []atom.Atom{
		NewCmdVoidOutU8(20),
		NewCmdVoidOutS8(-20),
		NewCmdVoidOutU16(200),
		NewCmdVoidOutS16(-200),
		NewCmdVoidOutF32(1.0),
		NewCmdVoidOutU32(2000),
		NewCmdVoidOutS32(-2000),
		NewCmdVoidOutF64(1.0),
		NewCmdVoidOutU64(20000),
		NewCmdVoidOutS64(-20000),
		NewCmdVoidOutBool(true),
		NewCmdVoidOutString("hello"),
		NewCmdVoidOutFixedSizeBuffer(0xdeadbeef),
	}, []interface{}{
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidOutU8.ID},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.PushI{DataType: protocol.TypeUint32, Value: 1},
		opcode.Post{},

		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidOutS8.ID},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.PushI{DataType: protocol.TypeUint32, Value: 1},
		opcode.Post{},

		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidOutU16.ID},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.PushI{DataType: protocol.TypeUint32, Value: 2},
		opcode.Post{},

		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidOutS16.ID},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.PushI{DataType: protocol.TypeUint32, Value: 2},
		opcode.Post{},

		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidOutF32.ID},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.PushI{DataType: protocol.TypeUint32, Value: 4},
		opcode.Post{},

		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidOutU32.ID},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.PushI{DataType: protocol.TypeUint32, Value: 4},
		opcode.Post{},

		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidOutS32.ID},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.PushI{DataType: protocol.TypeUint32, Value: 4},
		opcode.Post{},

		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidOutF64.ID},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.PushI{DataType: protocol.TypeUint32, Value: 8},
		opcode.Post{},

		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidOutU64.ID},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.PushI{DataType: protocol.TypeUint32, Value: 8},
		opcode.Post{},

		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidOutS64.ID},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.PushI{DataType: protocol.TypeUint32, Value: 8},
		opcode.Post{},

		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidOutBool.ID},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.PushI{DataType: protocol.TypeUint32, Value: 1},
		opcode.Post{},

		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidOutString.ID},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.PushI{DataType: protocol.TypeUint32, Value: 10},
		opcode.Post{},

		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidOutFixedSizeBuffer.ID},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.PushI{DataType: protocol.TypeUint32, Value: 10},
		opcode.Post{},
	}, []byte{})
}

func TestOperationsOpCall_3OutputStrings_DontWantOutput(t *testing.T) {
	check(t, 4 /* pointer size */, 4 /* pointer alignment */, false /* wantOutput */, []atom.Atom{
		NewCmdVoidOut3Strings("hello", "world", "hello"),
	}, []interface{}{
		// 0x00: a (byte[0x0f])
		// 0x10: b (byte[0x1f])
		// 0x30: c (byte[0x2f])
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x00},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x10},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x30},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidOut3Strings.ID},
	}, []byte{})
}

func TestOperationsOpCall_3OutputStrings_WantOutput(t *testing.T) {
	check(t, 4 /* pointer size */, 4 /* pointer alignment */, true /* wantOutput */, []atom.Atom{
		NewCmdVoidOut3Strings("hello", "world", "hello"),
	}, []interface{}{
		// 0x00: a (byte[0x0f])
		// 0x10: b (byte[0x1f])
		// 0x30: c (byte[0x2f])
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x00},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x10},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x30},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidOut3Strings.ID},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
		opcode.PushI{DataType: protocol.TypeUint32, Value: 0x5f},
		opcode.Post{},
	}, []byte{})
}

func TestOperationsOpCall_RemappedInputs(t *testing.T) {
	check(t, 4 /* pointer size */, 4 /* pointer alignment */, false /* wantOutput */, []atom.Atom{
		NewCmdVoid3Remapped(0x10, 0x20, 0x10),
	}, []interface{}{
		opcode.PushI{DataType: protocol.TypeUint32, Value: 0x10},
		opcode.Clone{Index: 0},
		opcode.StoreV{Address: 0x0},
		opcode.PushI{DataType: protocol.TypeUint32, Value: 0x20},
		opcode.Clone{Index: 0},
		opcode.StoreV{Address: 0x4},
		opcode.LoadV{DataType: protocol.TypeUint32, Address: 0x00},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoid3Remapped.ID},
	}, []byte{})
}

func TestOperationsOpCall_RemappedOutputs(t *testing.T) {
	check(t, 4 /* pointer size */, 4 /* pointer alignment */, false /* wantOutput */, []atom.Atom{
		NewCmdVoidOut3Remapped(0x10, 0x20, 0x10),
	}, []interface{}{
		// TODO: Sub-optimal output - these could be written straight to the remapped slots.
		// 0x00: id<0x10>
		// 0x04: id<0x20>
		// 0x08: a
		// 0x0c: b
		// 0x10: c
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x08},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x0c},
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x10},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidOut3Remapped.ID},
		opcode.LoadV{DataType: protocol.TypeUint32, Address: 0x08}, // a
		opcode.StoreV{Address: 0x00},                               // id<0x10>
		opcode.LoadV{DataType: protocol.TypeUint32, Address: 0x0c}, // b
		opcode.StoreV{Address: 0x04},                               // id<0x20>
	}, []byte{})
}

func TestOperationsOpCall_RemappedArrayOutput(t *testing.T) {
	check(t, 4 /* pointer size */, 4 /* pointer alignment */, false /* wantOutput */, []atom.Atom{
		NewCmdVoidOutArrayOfRemapped(RemappedArray{0x10, 0x20, 0x10, 0x30, 0x10}),
	}, []interface{}{
		// 0x00: id<0x10>
		// 0x04: id<0x20>
		// 0x08: id<0x30>
		// 0x0c: a[0]
		// 0x10: a[1]
		// 0x14: a[2]
		// 0x18: a[3]
		// 0x1c: a[4]
		opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x0c},
		opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidOutArrayOfRemapped.ID},
		opcode.LoadV{DataType: protocol.TypeUint32, Address: 0x0c}, // a[0]
		opcode.StoreV{Address: 0x00},                               // id<0x10>
		opcode.LoadV{DataType: protocol.TypeUint32, Address: 0x10}, // a[1]
		opcode.StoreV{Address: 0x04},                               // id<0x20>
		opcode.LoadV{DataType: protocol.TypeUint32, Address: 0x18}, // a[3]
		opcode.StoreV{Address: 0x08},                               // id<0x30>
	}, []byte{})
}
