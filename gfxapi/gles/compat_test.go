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

package gles_test

import (
	"strings"
	"testing"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary/endian"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/device"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/gfxapi/gles"
	"android.googlesource.com/platform/tools/gpu/gfxapi/gles/glsl/ast"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/service"
)

var (
	compat     = gles.VisibleForTestingCompat
	getContext = gles.VisibleForTestingGetContext
	glslCompat = gles.VisibleForTestingGlSlCompat
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

func newContextInfo(a device.Architecture, d database.Database, l log.Logger, width, height int, preserveBuffersOnSwap bool) atom.Atom {
	names := []gles.GLenum{}
	offsets := []uint32{}
	sizes := []uint32{}
	data := ""
	for name, value := range map[gles.GLenum]string{
		gles.GLenum_GL_VERSION: "OpenGL ES 2.0",
	} {
		names = append(names, name)
		offsets = append(offsets, uint32(len(data)))
		sizes = append(sizes, uint32(len(value)))
		data = data + value
	}

	return gles.NewContextInfo(
		uint32(len(names)),
		p(0x10000),
		p(0x20000),
		p(0x30000),
		p(0x40000),
		gles.GLsizei(width),
		gles.GLsizei(height),
		gles.GLenum_GL_RGB565,
		gles.GLenum_GL_DEPTH_COMPONENT16,
		gles.GLenum_GL_STENCIL_INDEX8,
		true,
		preserveBuffersOnSwap).
		AddRead(atom.Data(a, d, l, p(0x10000), names)).
		AddRead(atom.Data(a, d, l, p(0x20000), offsets)).
		AddRead(atom.Data(a, d, l, p(0x30000), sizes)).
		AddRead(atom.Data(a, d, l, p(0x40000), data))
}

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

	shaderType := gles.GLenum_GL_VERTEX_SHADER
	if c.lang == ast.LangFragmentShader {
		shaderType = gles.GLenum_GL_FRAGMENT_SHADER
	}

	mw := &mockWriter{}
	for _, a := range []atom.Atom{
		gles.NewEglCreateContext(memory.Nullptr, memory.Nullptr, memory.Nullptr, memory.Nullptr, memory.Nullptr),
		gles.NewEglMakeCurrent(memory.Nullptr, memory.Nullptr, memory.Nullptr, memory.Nullptr, 0),
		newContextInfo(a, d, l, 64, 64, true),
		gles.NewGlCreateShader(shaderType, 0x10),
		gles.NewGlShaderSource(0x10, 1, p(0x100000), p(0x100010)).
			AddRead(atom.Data(a, d, l, p(0x100000), p(0x100020))).
			AddRead(atom.Data(a, d, l, p(0x100010), int32(len(c.source)))).
			AddRead(atom.Data(a, d, l, p(0x100020), c.source)),
	} {
		transform.Transform(atom.NoID, a, mw)
	}

	// Find the output glShaderSource atom.
	var cmd *gles.GlShaderSource
	for _, a := range mw.atoms {
		if a, ok := a.(*gles.GlShaderSource); ok {
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

func TestGlVertexAttribPointerCompatTest(t *testing.T) {
	d, l := database.NewInMemory(nil), log.Testing(t)
	a := device.Architecture{
		PointerAlignment: 4,
		PointerSize:      4,
		IntegerSize:      4,
		ByteOrder:        endian.Little,
	}

	device := &service.Device{Version: OpenGL_3_0}
	transform, err := compat(device, d, l)
	if err != nil {
		log.E(l, "Failed to create compatability transform: %v", err)
		return
	}

	positions := []float32{-1., -1., 1., -1., -1., 1., 1., 1.}
	indices := []uint16{0, 1, 2, 1, 2, 3}
	mw := &mockWriter{}
	for _, a := range []atom.Atom{
		gles.NewEglCreateContext(memory.Nullptr, memory.Nullptr, memory.Nullptr, memory.Nullptr, memory.Nullptr),
		gles.NewEglMakeCurrent(memory.Nullptr, memory.Nullptr, memory.Nullptr, memory.Nullptr, 0),
		newContextInfo(a, d, l, 64, 64, true),
		gles.NewGlEnableVertexAttribArray(0),
		gles.NewGlVertexAttribPointer(0, 2, gles.GLenum_GL_FLOAT, gles.GLboolean(0), 8, p(0x100000)).
			AddRead(atom.Data(a, d, l, p(0x100000), positions)),
		gles.NewGlDrawElements(gles.GLenum_GL_TRIANGLES, gles.GLsizei(len(indices)), gles.GLenum_GL_UNSIGNED_SHORT, p(0x200000)).
			AddRead(atom.Data(a, d, l, p(0x200000), indices)),
	} {
		transform.Transform(atom.NoID, a, mw)
	}

	// Find glDrawElements and check it is using a buffer instead of client's memory now
	s := gfxapi.NewState()
	for _, a := range mw.atoms {
		a.Mutate(s, d, l)
		if _, ok := a.(*gles.GlDrawElements); ok {
			ctx := getContext(s)
			vao := ctx.Instances.VertexArrays[ctx.BoundVertexArray]
			array := vao.VertexAttributeArrays[0]
			binding := vao.VertexBufferBindings[array.Binding]
			if binding.Buffer != 0 && array.Pointer.Address == 0 {
				return // Success
			} else {
				t.Error("glDrawElements does not source vertex data from buffer.")
				return
			}
		}
	}

	t.Error("glDrawElements atom not found.")
	return
}

func TestShaderCompat(t *testing.T) {
	for _, test := range glslCompatTests {
		glShaderSourceCompatTest(test).run(t)
	}
}
