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
	"testing"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/endian"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/device"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/replay/builder"
	"android.googlesource.com/platform/tools/gpu/replay/opcode"
	"android.googlesource.com/platform/tools/gpu/replay/protocol"
)

type write struct {
	at  memory.Pointer
	src memory.Slice
}

type expected struct {
	opcodes   []interface{}
	resources []binary.ID
	constants []byte
}

type test struct {
	writes   []write
	atoms    []atom.Atom
	expected expected
}

func check(t *testing.T, a device.Architecture, d database.Database, l log.Logger, test test) {
	b := builder.New(a)
	s := gfxapi.NewState()

	// TODO: Test for different capture / replay architectures
	s.Architecture = a

	for _, w := range test.writes {
		s.Memory[memory.ApplicationPool].Write(w.at, w.src)
	}

	for i, a := range test.atoms {
		func() {
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("Panic replaying atom %d %T: %v\n%v", i, a, err, s)
					panic(err)
				}
			}()
			replay.Replay(atom.ID(i), a, s, d, l, b)
		}()
	}

	payload, _, err := b.Build(log.Nop{})
	if err != nil {
		t.Errorf("Failed to build opcodes: %v", err)
	}

	ops := bytes.NewBuffer(payload.Opcodes)
	gotOpcodes, err := opcode.Disassemble(ops, endian.Little)
	if err != nil {
		t.Errorf("Failed to disassemble opcodes: %v", err)
	}
	opcode.CheckDisassembly(t, gotOpcodes, test.expected.opcodes...)

	checkResource(t, payload.Resources, test.expected.resources)

	if !bytes.Equal(payload.Constants, test.expected.constants) {
		t.Errorf("Constant buffer was not as expected.\nGot:      % .2x\nExpected: % .2x",
			payload.Constants, test.expected.constants)
	}
}

func checkResource(t *testing.T, got []protocol.ResourceInfo, expected []binary.ID) {
	matched := len(got) == len(expected)
	if matched {
		for i, expected := range expected {
			g, e := got[i].ID, expected.String()
			if g != e {
				matched = false
				break
			}
		}
	}
	if !matched {
		t.Errorf("Resources were not as expected:")
		c := max(len(got), len(expected))
		for i := 0; i < c; i++ {
			g, e := "<none>", "<none>"
			if i < len(got) {
				g = got[i].ID
			}
			if i < len(expected) {
				e = expected[i].String()
			}
			if g == e {
				t.Errorf("  %d: %v", i, g)
			} else {
				t.Errorf("* %d: %v ---  EXPECTED: %v", i, g, e)
			}
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func TestOperationsOpCall_NoIn_NoOut(t *testing.T) {
	d, l := database.InMemory(), log.Testing(t)
	a := device.Architecture{
		PointerAlignment: 4,
		PointerSize:      4,
		IntegerSize:      4,
		ByteOrder:        endian.Little,
	}
	check(t, a, d, l, test{
		atoms: []atom.Atom{
			NewCmdVoid(),
		},
		expected: expected{
			opcodes: []interface{}{
				opcode.Label{Value: 0},
				opcode.Call{FunctionID: funcInfoCmdVoid.ID},
			},
		},
	})
}

func TestOperationsOpCall_Unknowns(t *testing.T) {
	d, l := database.InMemory(), log.Testing(t)
	a := device.Architecture{
		PointerAlignment: 4,
		PointerSize:      4,
		IntegerSize:      8,
		ByteOrder:        endian.Little,
	}

	check(t, a, d, l, test{
		atoms: []atom.Atom{
			NewCmdUnknownRet(10),
			NewCmdUnknownWritePtr(0x200000).
				AddRead(atom.Data(a, d, l, 0x200000, int(100))).
				AddWrite(atom.Data(a, d, l, 0x200000, int(200))),
			NewCmdUnknownWriteSlice(0x100000).
				AddRead(atom.Data(a, d, l, 0x100000, []int{0, 1, 2, 3, 4})).
				AddWrite(atom.Data(a, d, l, 0x100000, []int{5, 6, 7, 8, 9})),
		},
		expected: expected{
			opcodes: []interface{}{
				opcode.Label{Value: 0},
				opcode.Call{FunctionID: funcInfoCmdUnknownRet.ID},

				opcode.Label{Value: 1},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 8 * 5},
				opcode.Call{FunctionID: funcInfoCmdUnknownWritePtr.ID},

				opcode.Label{Value: 2},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
				opcode.Call{FunctionID: funcInfoCmdUnknownWriteSlice.ID},
			},
		},
	})
}

func TestOperationsOpCall_SingleInputArg(t *testing.T) {
	d, l := database.InMemory(), log.Testing(t)
	a := device.Architecture{
		PointerAlignment: 4,
		PointerSize:      4,
		IntegerSize:      4,
		ByteOrder:        endian.Little,
	}
	check(t, a, d, l, test{
		atoms: []atom.Atom{
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
		},
		expected: expected{
			opcodes: []interface{}{
				opcode.Label{Value: 0},
				opcode.PushI{DataType: protocol.TypeUint8, Value: 20},
				opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidU8.ID},

				opcode.Label{Value: 1},
				opcode.PushI{DataType: protocol.TypeInt8, Value: 0xfffec},
				opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidS8.ID},

				opcode.Label{Value: 2},
				opcode.PushI{DataType: protocol.TypeUint16, Value: 200},
				opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidU16.ID},

				opcode.Label{Value: 3},
				opcode.PushI{DataType: protocol.TypeInt16, Value: 0xfff38},
				opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidS16.ID},

				opcode.Label{Value: 4},
				opcode.PushI{DataType: protocol.TypeFloat, Value: 0x7f},
				opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidF32.ID},

				opcode.Label{Value: 5},
				opcode.PushI{DataType: protocol.TypeUint32, Value: 2000},
				opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidU32.ID},

				opcode.Label{Value: 6},
				opcode.PushI{DataType: protocol.TypeInt32, Value: 0xff830},
				opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidS32.ID},

				opcode.Label{Value: 7},
				opcode.PushI{DataType: protocol.TypeDouble, Value: 0x3ff},
				opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidF64.ID},

				opcode.Label{Value: 8},
				opcode.PushI{DataType: protocol.TypeUint64, Value: 20000},
				opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidU64.ID},

				opcode.Label{Value: 9},
				opcode.PushI{DataType: protocol.TypeInt64, Value: 0xfb1e0},
				opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidS64.ID},

				opcode.Label{Value: 10},
				opcode.PushI{DataType: protocol.TypeBool, Value: 1},
				opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidBool.ID},

				opcode.Label{Value: 11},
				opcode.PushI{DataType: protocol.TypeConstantPointer, Value: 0x00},
				opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoidString.ID},
			},
			constants: []byte{'h', 'e', 'l', 'l', 'o', 0},
		},
	})
}

func TestOperationsOpCall_3_Strings(t *testing.T) {
	d, l := database.InMemory(), log.Testing(t)
	a := device.Architecture{
		PointerAlignment: 4,
		PointerSize:      4,
		IntegerSize:      4,
		ByteOrder:        endian.Little,
	}
	check(t, a, d, l, test{
		atoms: []atom.Atom{
			NewCmdVoid3Strings("hello", "world", "hello"),
		},
		expected: expected{
			opcodes: []interface{}{
				opcode.Label{Value: 0},
				opcode.PushI{DataType: protocol.TypeConstantPointer, Value: 0x00},
				opcode.PushI{DataType: protocol.TypeConstantPointer, Value: 0x08},
				opcode.PushI{DataType: protocol.TypeConstantPointer, Value: 0x00},
				opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoid3Strings.ID},
			},
			constants: []byte{
				/* 0x00 */ 'h', 'e', 'l', 'l', 'o', 0x00, 0x00, 0x00,
				/* 0x08 */ 'w', 'o', 'r', 'l', 'd', 0x00,
			},
		},
	})
}

func TestOperationsOpCall_3_In_Arrays(t *testing.T) {
	d, l := database.InMemory(), log.Testing(t)
	a := device.Architecture{
		PointerAlignment: 8,
		PointerSize:      4,
		IntegerSize:      8,
		ByteOrder:        endian.Little,
	}

	aRng, aID := atom.Data(a, d, l, 0x40000+5* /* sizeof(u8)  */ 1, []uint8{
		5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
	})
	bRng, bID := atom.Data(a, d, l, 0x50000+5* /* sizeof(u8*) */ 4, []memory.Pointer{
		5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
	})
	cRng, cID := atom.Data(a, d, l, 0x60000+5* /* sizeof(int) */ 8, []int{
		5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
	})

	check(t, a, d, l, test{
		atoms: []atom.Atom{
			NewCmdVoid3InArrays(0x40000, 0x50000, 0x60000).
				AddRead(aRng, aID).
				AddRead(bRng, bID).
				AddRead(cRng, cID),
		},
		expected: expected{
			//   ┌────┬────┬────┬────┬────╔════╤════╤════╤════╤════╤════╤════╤════╤════╤════╗
			// b │0x10│0x14│0x18│0x1c│0x20║0x24│0x28│0x2c│0x30│0x34│0x38│0x3c│0x40│0x44│0x48║
			//   └────┴────┴────┴────┴────╚════╧════╧════╧════╧════╧════╧════╧════╧════╧════╝
			//   ┌────┬────┬────┬────┬────╔════╤════╤════╤════╤════╤════╤════╤════╤════╤════╗
			// c │0x50│0x58│0x60│0x68│0x70║0x78│0x80│0x88│0x90│0x98│0xa0│0xa8│0xb0│0xb8│0xc0║
			//   └────┴────┴────┴────┴────╚════╧════╧════╧════╧════╧════╧════╧════╧════╧════╝
			//   ┌────┬────┬────┬────┬────╔════╤════╤════╤════╤════╤════╤════╤════╤════╤════╗
			// a │0x00│0x01│0x02│0x03│0x04║0x05│0x06│0x07│0x08│0x09│0x0a│0x0b│0x0c│0x0d│0x0e║
			//   └────┴────┴────┴────┴────╚════╧════╧════╧════╧════╧════╧════╧════╧════╧════╝
			resources: []binary.ID{bID, cID, aID},
			opcodes: []interface{}{
				opcode.Label{Value: 0},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x24},
				opcode.Resource{ID: 0},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x78},
				opcode.Resource{ID: 1},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x5},
				opcode.Resource{ID: 2},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x00},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x10},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x50},
				opcode.Call{PushReturn: false, FunctionID: funcInfoCmdVoid3InArrays.ID},
			},
		},
	})
}

func TestOperationsOpCall_SinglePointerElementRead(t *testing.T) {
	d, l := database.InMemory(), log.Testing(t)
	a := device.Architecture{
		PointerAlignment: 4,
		PointerSize:      4,
		IntegerSize:      4,
		ByteOrder:        endian.Little,
	}
	p := memory.Pointer(0x100000)
	rng1, id1 := atom.Data(a, d, l, p, []byte{
		0x01,
	})
	rng2, id2 := atom.Data(a, d, l, p, []byte{
		0x01, 0x23,
	})
	rng4, id4 := atom.Data(a, d, l, p, []byte{
		0x01, 0x23, 0x45, 0x67,
	})
	rng8, id8 := atom.Data(a, d, l, p, []byte{
		0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef,
	})
	check(t, a, d, l, test{
		atoms: []atom.Atom{
			NewCmdVoidReadBool(p).AddRead(rng1, id1),
			NewCmdVoidReadU8(p).AddRead(rng1, id1),
			NewCmdVoidReadS8(p).AddRead(rng1, id1),
			NewCmdVoidReadU16(p).AddRead(rng2, id2),
			NewCmdVoidReadS16(p).AddRead(rng2, id2),
			NewCmdVoidReadF32(p).AddRead(rng4, id4),
			NewCmdVoidReadU32(p).AddRead(rng4, id4),
			NewCmdVoidReadS32(p).AddRead(rng4, id4),
			NewCmdVoidReadF64(p).AddRead(rng8, id8),
			NewCmdVoidReadU64(p).AddRead(rng8, id8),
			NewCmdVoidReadS64(p).AddRead(rng8, id8),

			NewCmdVoidReadS32(p),  // Uses previous observations
			NewCmdVoidReadBool(p), // Uses previous observations
		},
		expected: expected{
			resources: []binary.ID{id1, id2, id4, id8},
			opcodes: []interface{}{
				opcode.Label{Value: 0},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
				opcode.Resource{ID: 0},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
				opcode.Call{FunctionID: funcInfoCmdVoidReadBool.ID},

				opcode.Label{Value: 1},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
				opcode.Resource{ID: 0},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
				opcode.Call{FunctionID: funcInfoCmdVoidReadU8.ID},

				opcode.Label{Value: 2},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
				opcode.Resource{ID: 0},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
				opcode.Call{FunctionID: funcInfoCmdVoidReadS8.ID},

				opcode.Label{Value: 3},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
				opcode.Resource{ID: 1},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
				opcode.Call{FunctionID: funcInfoCmdVoidReadU16.ID},

				opcode.Label{Value: 4},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
				opcode.Resource{ID: 1},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
				opcode.Call{FunctionID: funcInfoCmdVoidReadS16.ID},

				opcode.Label{Value: 5},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
				opcode.Resource{ID: 2},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
				opcode.Call{FunctionID: funcInfoCmdVoidReadF32.ID},

				opcode.Label{Value: 6},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
				opcode.Resource{ID: 2},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
				opcode.Call{FunctionID: funcInfoCmdVoidReadU32.ID},

				opcode.Label{Value: 7},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
				opcode.Resource{ID: 2},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
				opcode.Call{FunctionID: funcInfoCmdVoidReadS32.ID},

				opcode.Label{Value: 8},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
				opcode.Resource{ID: 3},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
				opcode.Call{FunctionID: funcInfoCmdVoidReadF64.ID},

				opcode.Label{Value: 9},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
				opcode.Resource{ID: 3},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
				opcode.Call{FunctionID: funcInfoCmdVoidReadU64.ID},

				opcode.Label{Value: 10},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
				opcode.Resource{ID: 3},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
				opcode.Call{FunctionID: funcInfoCmdVoidReadS64.ID},

				opcode.Label{Value: 11},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
				opcode.Resource{ID: 2},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
				opcode.Call{FunctionID: funcInfoCmdVoidReadS32.ID},

				opcode.Label{Value: 12},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
				opcode.Resource{ID: 0},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0},
				opcode.Call{FunctionID: funcInfoCmdVoidReadBool.ID},
			},
		},
	})
}

func TestOperationsOpCall_MultiplePointerElementReads(t *testing.T) {
	d, l := database.InMemory(), log.Testing(t)
	a := device.Architecture{
		PointerAlignment: 16,
		PointerSize:      4,
		IntegerSize:      4,
		ByteOrder:        endian.Little,
	}
	aRng, aID := atom.Data(a, d, l, 0x100000, float32(10))
	bRng, bID := atom.Data(a, d, l, 0x200000, uint16(20))
	cRng, cID := atom.Data(a, d, l, 0x300000, false)
	check(t, a, d, l, test{
		atoms: []atom.Atom{
			NewCmdVoidReadPtrs(0x100000, 0x200000, 0x300000).
				AddRead(aRng, aID).
				AddRead(bRng, bID).
				AddRead(cRng, cID),
		},
		expected: expected{
			resources: []binary.ID{aID, bID, cID},
			opcodes: []interface{}{
				opcode.Label{Value: 0},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x00},
				opcode.Resource{ID: 0},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x10},
				opcode.Resource{ID: 1},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x20},
				opcode.Resource{ID: 2},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x00},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x10},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x20},
				opcode.Call{FunctionID: funcInfoCmdVoidReadPtrs.ID},
			},
		},
	})
}

func TestOperationsOpCall_SinglePointerElementWrite(t *testing.T) {
	d, l := database.InMemory(), log.Testing(t)
	a := device.Architecture{
		PointerAlignment: 4,
		PointerSize:      4,
		IntegerSize:      4,
		ByteOrder:        endian.Little,
	}
	check(t, a, d, l, test{
		atoms: []atom.Atom{
			NewCmdVoidWriteU8(0x100000).
				AddWrite(atom.Data(a, d, l, 0x100000, uint8(1))),
			NewCmdVoidWriteS8(0x200000).
				AddWrite(atom.Data(a, d, l, 0x200000, int8(1))),
			NewCmdVoidWriteU16(0x300000).
				AddWrite(atom.Data(a, d, l, 0x300000, uint16(1))),
			NewCmdVoidWriteS16(0x400000).
				AddWrite(atom.Data(a, d, l, 0x400000, int16(1))),
			NewCmdVoidWriteF32(0x500000).
				AddWrite(atom.Data(a, d, l, 0x500000, float32(1))),
			NewCmdVoidWriteU32(0x600000).
				AddWrite(atom.Data(a, d, l, 0x600000, uint32(1))),
			NewCmdVoidWriteS32(0x700000).
				AddWrite(atom.Data(a, d, l, 0x700000, int32(1))),
			NewCmdVoidWriteF64(0x800000).
				AddWrite(atom.Data(a, d, l, 0x800000, float64(1))),
			NewCmdVoidWriteU64(0x900000).
				AddWrite(atom.Data(a, d, l, 0x900000, uint64(1))),
			NewCmdVoidWriteS64(0xa00000).
				AddWrite(atom.Data(a, d, l, 0xa00000, int64(1))),
			NewCmdVoidWriteBool(0xb00000).
				AddWrite(atom.Data(a, d, l, 0xb00000, bool(true))),
		},
		expected: expected{
			opcodes: []interface{}{
				opcode.Label{Value: 0},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x00},
				opcode.Call{FunctionID: funcInfoCmdVoidWriteU8.ID},
				opcode.Label{Value: 1},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x04},
				opcode.Call{FunctionID: funcInfoCmdVoidWriteS8.ID},
				opcode.Label{Value: 2},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x08},
				opcode.Call{FunctionID: funcInfoCmdVoidWriteU16.ID},
				opcode.Label{Value: 3},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x0c},
				opcode.Call{FunctionID: funcInfoCmdVoidWriteS16.ID},
				opcode.Label{Value: 4},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x10},
				opcode.Call{FunctionID: funcInfoCmdVoidWriteF32.ID},
				opcode.Label{Value: 5},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x14},
				opcode.Call{FunctionID: funcInfoCmdVoidWriteU32.ID},
				opcode.Label{Value: 6},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x18},
				opcode.Call{FunctionID: funcInfoCmdVoidWriteS32.ID},
				opcode.Label{Value: 7},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x1c},
				opcode.Call{FunctionID: funcInfoCmdVoidWriteF64.ID},
				opcode.Label{Value: 8},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x24},
				opcode.Call{FunctionID: funcInfoCmdVoidWriteU64.ID},
				opcode.Label{Value: 9},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x2c},
				opcode.Call{FunctionID: funcInfoCmdVoidWriteS64.ID},
				opcode.Label{Value: 10},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x34},
				opcode.Call{FunctionID: funcInfoCmdVoidWriteBool.ID},
			},
		},
	})
}

func TestOperationsOpCall_MultiplePointerElementWrites(t *testing.T) {
	d, l := database.InMemory(), log.Testing(t)
	a := device.Architecture{
		PointerAlignment: 16,
		PointerSize:      4,
		IntegerSize:      4,
		ByteOrder:        endian.Little,
	}
	check(t, a, d, l, test{
		atoms: []atom.Atom{
			NewCmdVoidWritePtrs(0x100000, 0x200000, 0x300000),
		},
		expected: expected{
			opcodes: []interface{}{
				opcode.Label{Value: 0},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x00},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x10},
				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: 0x20},
				opcode.Call{FunctionID: funcInfoCmdVoidWritePtrs.ID},
			},
		},
	})
}

func TestOperationsOpCall_ReturnValue(t *testing.T) {
	d, l := database.InMemory(), log.Testing(t)
	a := device.Architecture{
		PointerAlignment: 4,
		PointerSize:      4,
		IntegerSize:      4,
		ByteOrder:        endian.Little,
	}
	check(t, a, d, l, test{
		atoms: []atom.Atom{
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
			NewCmdString("hello"),
			NewCmdPointer(0x10000),
		},
		expected: expected{
			opcodes: []interface{}{
				opcode.Label{Value: 0},
				opcode.Call{FunctionID: funcInfoCmdU8.ID},
				opcode.Label{Value: 1},
				opcode.Call{FunctionID: funcInfoCmdS8.ID},
				opcode.Label{Value: 2},
				opcode.Call{FunctionID: funcInfoCmdU16.ID},
				opcode.Label{Value: 3},
				opcode.Call{FunctionID: funcInfoCmdS16.ID},
				opcode.Label{Value: 4},
				opcode.Call{FunctionID: funcInfoCmdF32.ID},
				opcode.Label{Value: 5},
				opcode.Call{FunctionID: funcInfoCmdU32.ID},
				opcode.Label{Value: 6},
				opcode.Call{FunctionID: funcInfoCmdS32.ID},
				opcode.Label{Value: 7},
				opcode.Call{FunctionID: funcInfoCmdF64.ID},
				opcode.Label{Value: 8},
				opcode.Call{FunctionID: funcInfoCmdU64.ID},
				opcode.Label{Value: 9},
				opcode.Call{FunctionID: funcInfoCmdS64.ID},
				opcode.Label{Value: 10},
				opcode.Call{FunctionID: funcInfoCmdBool.ID},
				opcode.Label{Value: 11},
				opcode.Call{FunctionID: funcInfoCmdString.ID},
				opcode.Label{Value: 12},
				opcode.Call{FunctionID: funcInfoCmdPointer.ID},
			},
		},
	})
}

func TestOperationsOpCall_3Remapped(t *testing.T) {
	d, l := database.InMemory(), log.Testing(t)
	a := device.Architecture{
		PointerAlignment: 4,
		PointerSize:      4,
		IntegerSize:      4,
		ByteOrder:        endian.Little,
	}
	check(t, a, d, l, test{
		atoms: []atom.Atom{
			NewCmdVoid3Remapped(0x10, 0x20, 0x10),
		},
		expected: expected{
			opcodes: []interface{}{
				opcode.Label{Value: 0},
				// First-seen values get an identical remapping value.
				opcode.PushI{DataType: protocol.TypeUint32, Value: 0x10},
				opcode.Clone{Index: 0},
				opcode.StoreV{Address: 0x0},

				opcode.PushI{DataType: protocol.TypeUint32, Value: 0x20},
				opcode.Clone{Index: 0},
				opcode.StoreV{Address: 0x4},

				// Subsequently-seen values use the remapped value.
				opcode.LoadV{DataType: protocol.TypeUint32, Address: 0x00},

				opcode.Call{FunctionID: funcInfoCmdVoid3Remapped.ID},
			},
		},
	})
}

func TestOperationsOpCall_InArrayOfRemapped(t *testing.T) {
	d, l := database.InMemory(), log.Testing(t)
	a := device.Architecture{
		PointerAlignment: 4,
		PointerSize:      4,
		IntegerSize:      4,
		ByteOrder:        endian.Little,
	}

	rng, id := atom.Data(a, d, l, 0x100000, []remapped{10, 20, 10, 30, 20})

	pbase := uint32(4 * 3) // parameter array base address
	tbase := uint32(0)     // remap table base address

	check(t, a, d, l, test{
		atoms: []atom.Atom{
			NewCmdVoidInArrayOfRemapped(0x100000).
				AddRead(rng, id),
		},
		expected: expected{
			opcodes: []interface{}{
				opcode.Label{Value: 0},

				// 10 --> remap[0], param[0]
				opcode.PushI{DataType: protocol.TypeUint32, Value: 10},
				opcode.Clone{},
				opcode.StoreV{Address: tbase + 4*0},
				opcode.StoreV{Address: pbase + 4*0},

				// 20 --> remap[1], param[1]
				opcode.PushI{DataType: protocol.TypeUint32, Value: 20},
				opcode.Clone{},
				opcode.StoreV{Address: tbase + 4*1},
				opcode.StoreV{Address: pbase + 4*1},

				// remap[0] --> param[2]
				opcode.LoadV{DataType: protocol.TypeUint32, Address: tbase + 4*0},
				opcode.StoreV{Address: pbase + 4*2},

				// 30 --> remap[2], param[3]
				opcode.PushI{DataType: protocol.TypeUint32, Value: 30},
				opcode.Clone{},
				opcode.StoreV{Address: tbase + 4*2},
				opcode.StoreV{Address: pbase + 4*3},

				// remap[1] --> param[4]
				opcode.LoadV{DataType: protocol.TypeUint32, Address: tbase + 4*1},
				opcode.StoreV{Address: pbase + 4*4},

				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: pbase},
				opcode.Call{FunctionID: funcInfoCmdVoidInArrayOfRemapped.ID},
			},
		},
	})
}

func TestOperationsOpCall_OutArrayOfRemapped(t *testing.T) {
	d, l := database.InMemory(), log.Testing(t)
	a := device.Architecture{
		PointerAlignment: 4,
		PointerSize:      4,
		IntegerSize:      4,
		ByteOrder:        endian.Little,
	}

	pbase := uint32(4 * 3) // parameter array base address
	tbase := uint32(0)     // remap table base address

	check(t, a, d, l, test{
		atoms: []atom.Atom{
			NewCmdVoidOutArrayOfRemapped(0x100000).
				AddWrite(atom.Data(a, d, l, 0x100000, []remapped{10, 20, 10, 30, 20})),
		},
		expected: expected{
			opcodes: []interface{}{
				opcode.Label{Value: 0},

				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: pbase},
				opcode.Call{FunctionID: funcInfoCmdVoidOutArrayOfRemapped.ID},

				// param[0] --> remap[0]
				opcode.LoadV{DataType: protocol.TypeUint32, Address: pbase + 4*0},
				opcode.StoreV{Address: tbase + 4*0},

				// param[1] --> remap[1]
				opcode.LoadV{DataType: protocol.TypeUint32, Address: pbase + 4*1},
				opcode.StoreV{Address: tbase + 4*1},

				// param[3] --> remap[2]
				opcode.LoadV{DataType: protocol.TypeUint32, Address: pbase + 4*3},
				opcode.StoreV{Address: tbase + 4*2},
			},
		},
	})
}

func TestOperationsOpCall_OutArrayOfUnknownRemapped(t *testing.T) {
	d, l := database.InMemory(), log.Testing(t)
	a := device.Architecture{
		PointerAlignment: 4,
		PointerSize:      4,
		IntegerSize:      4,
		ByteOrder:        endian.Little,
	}

	pbase := uint32(4 * 3) // parameter array base address
	tbase := uint32(0)     // remap table base address

	check(t, a, d, l, test{
		atoms: []atom.Atom{
			NewCmdVoidOutArrayOfUnknownRemapped(0x100000).
				AddWrite(atom.Data(a, d, l, 0x100000, []remapped{10, 20, 10, 30, 20})),
		},
		expected: expected{
			opcodes: []interface{}{
				opcode.Label{Value: 0},

				opcode.PushI{DataType: protocol.TypeVolatilePointer, Value: pbase},
				opcode.Call{FunctionID: funcInfoCmdVoidOutArrayOfUnknownRemapped.ID},

				// param[0] --> remap[0]
				opcode.LoadV{DataType: protocol.TypeUint32, Address: pbase + 4*0},
				opcode.StoreV{Address: tbase + 4*0},

				// param[1] --> remap[1]
				opcode.LoadV{DataType: protocol.TypeUint32, Address: pbase + 4*1},
				opcode.StoreV{Address: tbase + 4*1},

				// param[3] --> remap[2]
				opcode.LoadV{DataType: protocol.TypeUint32, Address: pbase + 4*3},
				opcode.StoreV{Address: tbase + 4*2},
			},
		},
	})
}
