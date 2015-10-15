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
	"errors"
	"testing"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/endian"
	"android.googlesource.com/platform/tools/gpu/check"
	"android.googlesource.com/platform/tools/gpu/device"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/replay/asm"
	"android.googlesource.com/platform/tools/gpu/replay/protocol"
	"android.googlesource.com/platform/tools/gpu/replay/value"
)

func TestCommitAtom(t *testing.T) {
	for _, test := range []struct {
		name     string
		f        func(*Builder)
		expected []asm.Instruction
	}{
		{
			"Call with used return value",
			func(b *Builder) {
				b.BeginAtom(10)
				b.Push(value.U8(1))
				b.Call(FunctionInfo{123, protocol.TypeUint8, 1})
				b.Store(value.AbsolutePointer(0x10000))
				b.CommitAtom()
			},
			[]asm.Instruction{
				asm.Label{Value: 10},
				asm.Push{Value: value.U8(1)},
				asm.Call{PushReturn: true, FunctionID: 123},
				asm.Store{Destination: value.AbsolutePointer(0x10000)},
			},
		},
		{
			"Call with unused return value",
			func(b *Builder) {
				b.BeginAtom(10)
				b.Push(value.U8(1))
				b.Call(FunctionInfo{123, protocol.TypeUint8, 1})
				b.CommitAtom()
			},
			[]asm.Instruction{
				asm.Label{Value: 10},
				asm.Push{Value: value.U8(1)},
				asm.Call{PushReturn: false, FunctionID: 123},
			},
		},
		{
			"Remove unused push",
			func(b *Builder) {
				b.BeginAtom(10)
				b.Push(value.U32(12))
				b.CommitAtom()
			},
			[]asm.Instruction{
				asm.Label{Value: 10},
			},
		},
		{
			"Unused pushes",
			func(b *Builder) {
				b.BeginAtom(10)
				b.Push(value.U32(12))
				b.Push(value.U32(34))
				b.Push(value.U32(56))
				b.CommitAtom()
			},
			[]asm.Instruction{
				asm.Label{Value: 10},
			},
		},
		{
			"Unused clone",
			func(b *Builder) {
				b.BeginAtom(10)
				b.Call(FunctionInfo{123, protocol.TypeUint8, 0})
				b.Clone(0)
				b.CommitAtom()
			},
			[]asm.Instruction{
				asm.Label{Value: 10},
				asm.Call{PushReturn: true, FunctionID: 123},
				asm.Pop{Count: 1},
			},
		},
		{
			"Unused clone",
			func(b *Builder) {
				b.BeginAtom(10)
				b.Call(FunctionInfo{123, protocol.TypeUint8, 0})
				b.Clone(0)
				b.Store(value.AbsolutePointer(0x10000))
				b.CommitAtom()
			},
			[]asm.Instruction{
				asm.Label{Value: 10},
				asm.Call{PushReturn: true, FunctionID: 123},
				asm.Nop{},
				asm.Store{Destination: value.AbsolutePointer(0x10000)},
			},
		},
		{
			"Unused clone of return value",
			func(b *Builder) {
				b.BeginAtom(10)
				b.Call(FunctionInfo{123, protocol.TypeUint8, 0})
				b.Clone(0)
				b.CommitAtom()
			},
			[]asm.Instruction{
				asm.Label{Value: 10},
				// TODO: Could be optimised further.
				asm.Call{PushReturn: true, FunctionID: 123},
				asm.Pop{Count: 1},
			},
		},
		{
			"Use one of three return values",
			func(b *Builder) {
				b.BeginAtom(10)
				b.Call(FunctionInfo{123, protocol.TypeUint8, 0})
				b.Call(FunctionInfo{123, protocol.TypeUint8, 0})
				b.Call(FunctionInfo{123, protocol.TypeUint8, 0})
				b.Clone(1)
				b.Store(value.AbsolutePointer(0x10000))
				b.CommitAtom()
			},
			[]asm.Instruction{
				asm.Label{Value: 10},
				asm.Call{PushReturn: false, FunctionID: 123},
				asm.Call{PushReturn: true, FunctionID: 123},
				asm.Call{PushReturn: false, FunctionID: 123},
				asm.Nop{},
				asm.Store{Destination: value.AbsolutePointer(0x10000)},
			},
		},
	} {
		b := New(device.Architecture{
			PointerAlignment: 4,
			PointerSize:      4,
			IntegerSize:      4,
			ByteOrder:        endian.Little,
		})
		test.f(b)
		if !check.SlicesEqual(t, b.instructions, test.expected) {
			t.Errorf("Test '%s' failed:", test.name)
		}
	}
}

func TestRevertAtom(t *testing.T) {
	for _, test := range []struct {
		name     string
		f        func(*Builder)
		expected []asm.Instruction
	}{
		{
			"Revert atom",
			func(b *Builder) {
				b.BeginAtom(10)
				b.Push(value.U8(1))
				b.Call(FunctionInfo{123, protocol.TypeUint8, 1})
				b.Store(value.AbsolutePointer(0x10000))
				b.RevertAtom(nil)
			},
			[]asm.Instruction{},
		},
		{
			"Commit atom, revert atom",
			func(b *Builder) {
				b.BeginAtom(10)
				b.Push(value.U8(1))
				b.Call(FunctionInfo{123, protocol.TypeUint8, 1})
				b.Store(value.AbsolutePointer(0x10000))
				b.CommitAtom()
				b.BeginAtom(20)
				b.Push(value.U8(2))
				b.Call(FunctionInfo{234, protocol.TypeUint8, 1})
				b.Store(value.AbsolutePointer(0x10000))
				b.RevertAtom(nil)
			},
			[]asm.Instruction{
				asm.Label{Value: 10},
				asm.Push{Value: value.U8(1)},
				asm.Call{PushReturn: true, FunctionID: 123},
				asm.Store{Destination: value.AbsolutePointer(0x10000)},
			},
		},
	} {
		b := New(device.Architecture{
			PointerAlignment: 4,
			PointerSize:      4,
			IntegerSize:      4,
			ByteOrder:        endian.Little,
		})
		test.f(b)
		if !check.SlicesEqual(t, b.instructions, test.expected) {
			t.Errorf("Test '%s' failed:", test.name)
		}
	}
}

func TestRevertPostbackAtom(t *testing.T) {
	expectedErr := errors.New("Oh noes!")
	postbackErr := error(nil)
	postback := Postback(func(d binary.Decoder, err error) error {
		if d != nil {
			t.Errorf("Unexpected decoder passed to postback on RevertAtom")
		}
		postbackErr = err
		return nil
	})

	for _, test := range []struct {
		name     string
		f        func(*Builder)
		expected []asm.Instruction
	}{
		{
			"Revert postback atom",
			func(b *Builder) {
				b.BeginAtom(10)
				b.Post(value.AbsolutePointer(0x10000), 100, postback)
				b.RevertAtom(expectedErr)
			},
			[]asm.Instruction{},
		},
	} {
		b := New(device.Architecture{
			PointerAlignment: 4,
			PointerSize:      4,
			IntegerSize:      4,
			ByteOrder:        endian.Little,
		})
		test.f(b)
		if !check.SlicesEqual(t, b.instructions, test.expected) {
			t.Errorf("Test '%s' failed:", test.name)
		}
	}
	if postbackErr != expectedErr {
		t.Errorf("Postback was not informed of RevertAtom")
	}
}

func TestMapMemory(t *testing.T) {
	for _, test := range []struct {
		name     string
		f        func(*Builder)
		expected []asm.Instruction
	}{
		{
			"No mapping",
			func(b *Builder) {
				b.BeginAtom(10)
				b.Push(value.ObservedPointer(0x100004))
				b.Call(FunctionInfo{123, protocol.TypeVolatilePointer, 1})
				b.CommitAtom()
			},
			[]asm.Instruction{
				asm.Label{Value: 10},
				asm.Push{Value: value.ObservedPointer(0x100004)},
				asm.Call{FunctionID: 123},
			},
		},
		{
			"MapMemory",
			func(b *Builder) {
				b.BeginAtom(10)
				b.Call(FunctionInfo{100, protocol.TypeAbsolutePointer, 0})
				b.MapMemory(memory.Range{Base: 0x100000, Size: 0x10})
				b.CommitAtom()

				b.BeginAtom(20)
				b.Push(value.ObservedPointer(0x100004))
				b.Call(FunctionInfo{123, protocol.TypeVoid, 1})
				b.CommitAtom()
			},
			[]asm.Instruction{
				asm.Label{Value: 10},
				asm.Call{FunctionID: 100, PushReturn: true},
				asm.Store{Destination: value.VolatilePointer(0x0)}, // mapping-storage address

				asm.Label{Value: 20},
				asm.Load{DataType: protocol.TypeAbsolutePointer, Source: value.VolatilePointer(0x0)}, // mapping-storage address
				asm.Push{Value: value.AbsolutePointer(0x4)},
				asm.Add{Count: 2},
				asm.Call{FunctionID: 123},
			},
		},
		{
			"UnmapMemory",
			func(b *Builder) {
				b.BeginAtom(10)
				b.Call(FunctionInfo{100, protocol.TypeAbsolutePointer, 0})
				b.MapMemory(memory.Range{Base: 0x100000, Size: 0x10})
				b.CommitAtom()

				b.BeginAtom(20)
				b.UnmapMemory(memory.Range{Base: 0x100000, Size: 0x10})
				b.CommitAtom()

				b.BeginAtom(30)
				b.Push(value.ObservedPointer(0x100004))
				b.Call(FunctionInfo{123, protocol.TypeVoid, 1})
				b.CommitAtom()
			},
			[]asm.Instruction{
				asm.Label{Value: 10},
				asm.Call{FunctionID: 100, PushReturn: true},
				asm.Store{Destination: value.VolatilePointer(0x0)}, // mapping-storage address

				asm.Label{Value: 20},

				asm.Label{Value: 30},
				asm.Push{Value: value.ObservedPointer(0x100004)},
				asm.Call{FunctionID: 123},
			},
		},
	} {
		b := New(device.Architecture{
			PointerAlignment: 4,
			PointerSize:      4,
			IntegerSize:      4,
			ByteOrder:        endian.Little,
		})
		test.f(b)
		if !check.SlicesEqual(t, b.instructions, test.expected) {
			t.Errorf("Test '%s' failed:", test.name)
		}
	}
}
