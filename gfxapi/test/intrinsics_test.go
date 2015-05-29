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
)

func checkBytes(t *testing.T, got, expected []byte) {
	if !bytes.Equal(got, expected) {
		t.Errorf("Data was not as expected.\nGot:      % .2x\nExpected: % .2x", got, expected)
	}
}

func TestClone(t *testing.T) {
	s, d, l := gfxapi.NewState(), database.InMemory(), log.Testing(t)
	expected := []byte{0x54, 0x33, 0x42, 0x43, 0x46, 0x34, 0x63, 0x24, 0x14, 0x24}
	for _, a := range []atom.Atom{
		NewCmdClone(0x1234, 10).
			AddRead(atom.Data(s.Architecture, d, l, 0x1234, expected)),
	} {
		a.Mutate(s, d, l)
	}
	got := getState(s).Buf.Read(s, d, l)
	checkBytes(t, got, expected)
}

func TestMake(t *testing.T) {
	s, d, l := gfxapi.NewState(), database.InMemory(), log.Testing(t)
	if s.NextPoolID != 1 {
		t.Errorf("Expected initial NextPoolID to be 1, instead got %d", s.NextPoolID)
	}
	NewCmdMake(10).Mutate(s, d, l)
	if c := getState(s).Buf.Count; c != 10 {
		t.Errorf("Expected buffer count to be 10, instead got %d", c)
	}
	if s.NextPoolID != 2 {
		t.Errorf("Expected initial NextPoolID to be 2, instead got %d", s.NextPoolID)
	}
}

func TestCopy(t *testing.T) {
	s, d, l := gfxapi.NewState(), database.InMemory(), log.Testing(t)
	expected := []byte{0x54, 0x33, 0x42, 0x43, 0x46, 0x34, 0x63, 0x24, 0x14, 0x24}
	for _, a := range []atom.Atom{
		NewCmdMake(10),
		NewCmdCopy(0x1234, 10).
			AddRead(atom.Data(s.Architecture, d, l, 0x1234, expected)),
	} {
		a.Mutate(s, d, l)
	}
	got := getState(s).Buf.Read(s, d, l)
	checkBytes(t, got, expected)
}

func TestCharsliceToString(t *testing.T) {
	s, d, l := gfxapi.NewState(), database.InMemory(), log.Testing(t)
	expected := "ħęľĺő ŵōřŀď"
	for _, a := range []atom.Atom{
		NewCmdCharsliceToString(0x1234, uint32(len(expected))).
			AddRead(atom.Data(s.Architecture, d, l, 0x1234, expected)),
	} {
		a.Mutate(s, d, l)
	}
	got := getState(s).Str
	if got != expected {
		t.Errorf("Data was not as expected.\nGot:      '%s'\nExpected: '%s'", got, expected)
	}
}
