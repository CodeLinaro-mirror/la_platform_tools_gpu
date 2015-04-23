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
	"android.googlesource.com/platform/tools/gpu/gfxapi/gles/glsl"
	"android.googlesource.com/platform/tools/gpu/gfxapi/gles/glsl/ast"
)

type mockWriter struct {
	atoms []atom.Atom
}

func (m *mockWriter) Write(id atom.ID, a atom.Atom) {
	m.atoms = append(m.atoms, a)
}

func runTest(t *testing.T, src string, expected string) {
	s := precisionStrip()
	tree, err := glsl.Parse(expected, ast.LangVertexShader)
	if len(err) > 0 {
		t.Errorf("Unexpected error parsing the expected output: %s.", err[0])
	}
	expected = fmt.Sprint(glsl.Formatter(tree))

	mw := &mockWriter{}

	s.Transform(0, NewGlShaderSource(0, 1, []string{src}, nil), mw)

	if len(mw.atoms) != 1 {
		t.Error("Unexpected number of Write calls: got %d, expected 1.", len(mw.atoms))
	}

	cmd, ok := mw.atoms[0].(*GlShaderSource)
	if !ok {
		t.Errorf("Received wrong kind of atom: got `%T`, expected `%T`.",
			mw.atoms[0], &GlShaderSource{})
		return
	}

	if cmd.Count != 1 {
		t.Errorf("Unexpected number of lines: got %d, expected 1.", cmd.Count)
		return
	}
	if cmd.Source[0] != expected {
		t.Errorf("Received unexpected string: got `%s`, expected `%s`.",
			cmd.Source[0], expected)
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
	s := precisionStrip()
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
