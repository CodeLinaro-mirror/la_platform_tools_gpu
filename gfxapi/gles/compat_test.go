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
	"strings"
	"testing"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary/endian"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/device"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/gfxapi/gles/glsl/ast"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/service"
)

type mockWriter struct {
	atoms []atom.Atom
}

func (m *mockWriter) Write(id atom.ID, a atom.Atom) {
	m.atoms = append(m.atoms, a)
}

func p(addr uint64) memory.Pointer {
	return memory.Pointer{Address: addr, Pool: memory.ApplicationPool}
}

type glShaderSourceCompatTest glslCompatTest

func (c glShaderSourceCompatTest) run(t *testing.T) {
	d, l := database.NewInMemory(nil), log.Testing(t)
	a := device.Architecture{
		PointerAlignment: 4,
		PointerSize:      4,
		IntegerSize:      4,
		ByteOrder:        endian.Little,
	}

	device := &service.Device{Version: c.target}
	transform, err := compat(device, d, l)
	if err != nil {
		log.E(l, "Failed to create compatability transform: %v", err)
		return
	}

	shaderType := GLenum_GL_VERTEX_SHADER
	if c.lang == ast.LangFragmentShader {
		shaderType = GLenum_GL_FRAGMENT_SHADER
	}

	mw := &mockWriter{}
	for _, a := range []atom.Atom{
		NewEglCreateContext(memory.Nullptr, memory.Nullptr, memory.Nullptr, memory.Nullptr, memory.Nullptr),
		NewEglMakeCurrent(memory.Nullptr, memory.Nullptr, memory.Nullptr, memory.Nullptr, 0),
		NewContextInfo("", "", "", "OpenGL ES 2.0", 64, 64, GLenum_GL_RGB565, GLenum_GL_DEPTH_COMPONENT16, GLenum_GL_STENCIL_INDEX8, true, true),
		NewGlCreateShader(shaderType, 0x10),
		NewGlShaderSource(0x10, 1, p(0x100000), p(0x100010)).
			AddRead(atom.Data(a, d, l, p(0x100000), p(0x100020))).
			AddRead(atom.Data(a, d, l, p(0x100010), int32(len(c.source)))).
			AddRead(atom.Data(a, d, l, p(0x100020), c.source)),
	} {
		transform.Transform(atom.NoID, a, mw)
	}

	// Find the output glShaderSource atom.
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

	srcPtr := cmd.Source.Read(cmd, s, d, l, nil) // 0'th glShaderSource string pointer
	got := strings.TrimRight(string(srcPtr.StringSlice(s, d, l).Read(cmd, s, d, l, nil)), "\x00")

	expected, err := glslCompat(c.source, c.lang, &service.Device{Version: c.target})
	if err != nil {
		t.Errorf("Unexpected error returned by glslCompat: %v", err)
	}
	if got != expected {
		t.Errorf("Converting to target '%s' produced unexpected output.\nGot:\n%s\nExpected:\n%v",
			c.target, got, expected)
	}
}

func TestShaderCompat(t *testing.T) {
	for _, test := range glslCompatTests {
		glShaderSourceCompatTest(test).run(t)
	}
}
