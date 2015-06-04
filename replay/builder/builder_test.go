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
	"testing"

	"android.googlesource.com/platform/tools/gpu/binary/endian"
	"android.googlesource.com/platform/tools/gpu/device"
	"android.googlesource.com/platform/tools/gpu/replay/asm"
	"android.googlesource.com/platform/tools/gpu/replay/protocol"
	"android.googlesource.com/platform/tools/gpu/replay/value"
)

func TestEndAtom(t *testing.T) {
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
				b.EndAtom()
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
				b.EndAtom()
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
				b.EndAtom()
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
				b.EndAtom()
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
				b.EndAtom()
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
				b.EndAtom()
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
				b.EndAtom()
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
				b.EndAtom()
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
		if !asm.Check(t, b.instructions, test.expected) {
			t.Errorf("Test '%s' failed:", test.name)
		}
	}
}
