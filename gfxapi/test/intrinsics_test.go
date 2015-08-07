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
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
)

func p(addr uint64) memory.Pointer {
	return memory.Pointer{Address: addr, Pool: memory.ApplicationPool}
}

func checkBytes(t *testing.T, got, expected []byte) {
	if !bytes.Equal(got, expected) {
		t.Errorf("Data was not as expected.\nGot:      % .2x\nExpected: % .2x", got, expected)
	}
}

func TestClone(t *testing.T) {
	s, d, l := gfxapi.NewState(), database.NewInMemory(nil), log.Testing(t)
	expected := []byte{0x54, 0x33, 0x42, 0x43, 0x46, 0x34, 0x63, 0x24, 0x14, 0x24}
	for _, a := range []atom.Atom{
		NewCmdClone(p(0x1234), 10).
			AddRead(atom.Data(s.Architecture, d, l, p(0x1234), expected)),
	} {
		a.Mutate(s, d, l)
	}
	got := getState(s).U8s.Read(nil, s, d, l, nil)
	checkBytes(t, got, expected)
}

func TestMake(t *testing.T) {
	s, d, l := gfxapi.NewState(), database.NewInMemory(nil), log.Testing(t)
	if s.NextPoolID != 1 {
		t.Errorf("Expected initial NextPoolID to be 1, instead got %d", s.NextPoolID)
	}
	NewCmdMake(10).Mutate(s, d, l)
	if c := getState(s).U8s.Count; c != 10 {
		t.Errorf("Expected buffer count to be 10, instead got %d", c)
	}
	if s.NextPoolID != 2 {
		t.Errorf("Expected initial NextPoolID to be 2, instead got %d", s.NextPoolID)
	}
}

func TestCopy(t *testing.T) {
	s, d, l := gfxapi.NewState(), database.NewInMemory(nil), log.Testing(t)
	expected := []byte{0x54, 0x33, 0x42, 0x43, 0x46, 0x34, 0x63, 0x24, 0x14, 0x24}
	for _, a := range []atom.Atom{
		NewCmdMake(10),
		NewCmdCopy(p(0x1234), 10).
			AddRead(atom.Data(s.Architecture, d, l, p(0x1234), expected)),
	} {
		a.Mutate(s, d, l)
	}
	got := getState(s).U8s.Read(nil, s, d, l, nil)
	checkBytes(t, got, expected)
}

func TestCharsliceToString(t *testing.T) {
	s, d, l := gfxapi.NewState(), database.NewInMemory(nil), log.Testing(t)
	expected := "ħęľĺő ŵōřŀď"
	NewCmdCharsliceToString(p(0x1234), uint32(len(expected))).
		AddRead(atom.Data(s.Architecture, d, l, p(0x1234), expected)).
		Mutate(s, d, l)
	if got := getState(s).Str; got != expected {
		t.Errorf("Data was not as expected.\nGot:      '%s'\nExpected: '%s'", got, expected)
	}
}

func TestCharptrToString(t *testing.T) {
	s, d, l := gfxapi.NewState(), database.NewInMemory(nil), log.Testing(t)
	expected := "ħęľĺő ŵōřŀď"
	NewCmdCharptrToString(p(0x1234)).
		AddRead(atom.Data(s.Architecture, d, l, p(0x1234), expected)).
		Mutate(s, d, l)
	if got := getState(s).Str; got != expected {
		t.Errorf("Data was not as expected.\nGot:      '%s'\nExpected: '%s'", got, expected)
	}
}

func TestSliceCasts(t *testing.T) {
	s, d, l := gfxapi.NewState(), database.NewInMemory(nil), log.Testing(t)
	s.Architecture.IntegerSize = 6 // non-multiple of u16
	NewCmdSliceCasts(p(0x1234), 10).Mutate(s, d, l)
	if got, expected := getState(s).U8s, NewU8ᵖ(0x1234).Slice(0, 20, s); got != expected {
		t.Errorf("U16[] -> U8[] was not as expected.\nGot:      '%s'\nExpected: '%s'", got, expected)
	}
	if got, expected := getState(s).U16s, NewU16ᵖ(0x1234).Slice(0, 10, s); got != expected {
		t.Errorf("U16[] -> U16[] was not as expected.\nGot:      '%s'\nExpected: '%s'", got, expected)
	}
	if got, expected := getState(s).U32s, NewU32ᵖ(0x1234).Slice(0, 5, s); got != expected {
		t.Errorf("U16[] -> U32[] was not as expected.\nGot:      '%s'\nExpected: '%s'", got, expected)
	}
	if got, expected := getState(s).Ints, NewIntᵖ(0x1234).Slice(0, 3, s); got != expected {
		t.Errorf("U16[] -> int[] was not as expected.\nGot:      '%s'\nExpected: '%s'", got, expected)
	}
}
