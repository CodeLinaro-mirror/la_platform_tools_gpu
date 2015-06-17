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

package gles

import (
	"fmt"
	"testing"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary/endian"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/device"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/gfxapi/gles/glsl"
	"android.googlesource.com/platform/tools/gpu/gfxapi/gles/glsl/ast"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
)

type mockWriter struct {
	atoms []atom.Atom
}

func (m *mockWriter) Write(id atom.ID, a atom.Atom) {
	m.atoms = append(m.atoms, a)
}

func runTest(t *testing.T, src string, expected string) {
	d, l := database.NewInMemory(nil), log.Testing(t)
	a := device.Architecture{
		PointerAlignment: 4,
		PointerSize:      4,
		IntegerSize:      4,
		ByteOrder:        endian.Little,
	}

	if tree, err := glsl.Parse(expected, ast.LangVertexShader); len(err) == 0 {
		expected = fmt.Sprint(glsl.Formatter(tree))
	} else {
		t.Errorf("Unexpected error parsing the expected output: %s.", err[0])
	}

	in := []atom.Atom{
		NewEglCreateContext(0, 0, 0, 0, 0),
		NewEglMakeCurrent(0, 0, 0, 0, 0),
		NewGlCreateShader(ShaderType_GL_VERTEX_SHADER, 0x10),
		NewGlShaderSource(0x10, 1, 0x100000, 0x100010).
			AddRead(atom.Data(a, d, l, 0x100000, memory.Pointer(0x100020))).
			AddRead(atom.Data(a, d, l, 0x100010, int32(len(src)))).
			AddRead(atom.Data(a, d, l, 0x100020, src)),
	}

	mw := &mockWriter{}
	ps := precisionStrip(d, l)
	for _, a := range in {
		ps.Transform(atom.NoID, a, mw)
	}

	var cmd *GlShaderSource
	for _, a := range mw.atoms {
		if a, ok := a.(*GlShaderSource); ok {
			cmd = a
			break
		}
	}

	if cmd == nil {
		t.Error("Transform did not produce a glShaderSource atom. Atoms produced:")
		for i, a := range mw.atoms {
			t.Errorf("%d %T", i, a)
		}
		return
	}

	if cmd.Count != 1 {
		t.Errorf("Unexpected number of sources: got %d, expected 1.", cmd.Count)
		return
	}

	s := gfxapi.NewState()
	for _, a := range mw.atoms {
		a.Mutate(s, d, l)
	}

	srcPtr := cmd.Source.Read(s, d, l) // 0'th glShaderSource string pointer

	if got := string(srcPtr.StringSlice(s, d, l, false).Read(s, d, l)); got != expected {
		t.Errorf("Received unexpected string at %v: got `%s`, expected `%s`.", srcPtr, got, expected)
		t.Errorf("Application memory pool writes:\n%v", s.Memory[memory.ApplicationPool])
	}
}

func TestStripTop(t *testing.T) {
	runTest(t, "precision highp int;", "")
}

func TestStripNested(t *testing.T) {
	runTest(t, "void main() { precision highp int; }", "void main() { }")
}

func TestStripIf(t *testing.T) {
	runTest(t, "void f(bool a) { if(a) precision highp int; }", "void f(bool a) { if(a); }")
}

func TestStripElse(t *testing.T) {
	runTest(t, "void f(bool a) { if(a) a=true; else precision highp int; }",
		"void f(bool a) { if(a) a=true; else ; }")
}

func TestStripWhile(t *testing.T) {
	runTest(t, "void f(bool a) { while(a) precision highp int; }",
		"void f(bool a) { while(a) ; }")
}

func TestStripDo(t *testing.T) {
	runTest(t, "void f(bool a) { do precision highp int; while(a); }",
		"void f(bool a) { do ; while(a); }")
}

func TestStripFor(t *testing.T) {
	runTest(t, "void f(bool a) { for(int b=0; b<1; ++b) precision highp int; }",
		"void f(bool a) { for(int b=0; b<1; ++b) ; }")
}

func TestStripFunction(t *testing.T) {
	runTest(t, "highp int f(lowp int b);", "int f(int b);")
}

func TestStripVarDecl(t *testing.T) {
	runTest(t, "highp int a;", "int a;")
}

func TestStripConversion(t *testing.T) {
	runTest(t, "int a; int b = lowp int(a);", "int a; int b = int(a);")
}

func TestStripPassthrough(t *testing.T) {
	d, l := database.Database(nil), log.Testing(t)
	s := precisionStrip(d, l)
	mw := &mockWriter{}
	a := &GlGetError{}

	s.Transform(0, a, mw)

	if len(mw.atoms) != 1 {
		t.Error("Unexpected number of Write calls: got %d, expected 1.", len(mw.atoms))
	}

	if mw.atoms[0] != a {
		t.Errorf("Received wrong atom: got `%v`, expected `%v`.", mw.atoms[0], a)
		return
	}
}
