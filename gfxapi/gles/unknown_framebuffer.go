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
	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/service"
)

// undefinedFramebuffer adds a transform that will render a pattern into the
// color buffer at the end of each frame.
func undefinedFramebuffer(device *service.Device, d database.Database, l log.Logger) atom.Transformer {
	s := gfxapi.NewState()
	return atom.Transform("DirtyFramebuffer", func(i atom.ID, a atom.Atom, out atom.Writer) {
		a.Mutate(s, d, l)
		out.Write(i, a)
		if a.Flags().IsEndOfFrame() {
			drawUndefinedFramebuffer(a, device, s, d, l, out)
		}
	})
}

func drawUndefinedFramebuffer(a atom.Atom, device *service.Device, s *gfxapi.State, d database.Database, l log.Logger, out atom.Writer) error {
	const (
		aScreenCoordsLocation AttributeLocation = 0

		vertexShaderSource string = `
					precision highp float;
					attribute vec2 aScreenCoords;
					varying vec2 uv;

					void main() {
						uv = aScreenCoords;
						gl_Position = vec4(aScreenCoords.xy, 0., 1.);
					}`
		fragmentShaderSource string = `
					precision highp float;
					varying vec2 uv;

					float F(float a) { return smoothstep(0.0, 0.1, a) * smoothstep(0.4, 0.3, a); }

					void main() {
						vec2 v = uv * 5.0;
						gl_FragColor = vec4(0.8, 0.9, 0.6, 1.0) * F(fract(v.x + v.y));
					}`
	)

	arch := s.Architecture
	version, _ := ParseVersion(device.Version)
	c := getContext(s)

	var (
		origProgramID     = c.BoundProgram
		origArrayBufferID = c.BoundBuffers[GLenum_GL_ARRAY_BUFFER]
	)

	// Generate new unused object IDs.
	programID := ProgramId(newUnusedID(func(x uint32) bool { _, ok := c.Instances.Programs[ProgramId(x)]; return ok }))
	vertexShaderID := ShaderId(newUnusedID(func(x uint32) bool { _, ok := c.Instances.Shaders[ShaderId(x)]; return ok }))
	fragmentShaderID := ShaderId(newUnusedID(func(x uint32) bool {
		_, ok := c.Instances.Shaders[ShaderId(x)]
		return ok || ShaderId(x) == vertexShaderID
	}))
	bufferID := BufferId(newUnusedID(func(x uint32) bool { _, ok := c.Instances.Buffers[BufferId(x)]; return ok }))
	arrayID := VertexArrayId(newUnusedID(func(x uint32) bool { _, ok := c.Instances.VertexArrays[VertexArrayId(x)]; return ok }))

	// 2D vertices positions for a full screen 2D triangle strip.
	positions := []float32{-1., -1., 1., -1., -1., 1., 1., 1.}

	t := tweaker{out: out, state: s, ctx: getContext(s), database: d, logger: l}

	// Temporarily change rasterizing/blending state and enable VAP 0.
	t.glDisable(GLenum_GL_BLEND)
	t.glDisable(GLenum_GL_CULL_FACE)
	t.glDisable(GLenum_GL_DEPTH_TEST)
	t.glDisable(GLenum_GL_SCISSOR_TEST)
	t.glDisable(GLenum_GL_STENCIL_TEST)
	t.bindOrSaveVertexArray(version, arrayID, aScreenCoordsLocation)
	out.Write(atom.NoID, NewGlEnableVertexAttribArray(aScreenCoordsLocation))

	// Create the shader program
	for _, a := range BuildProgram(s.Architecture, d, l, vertexShaderID, fragmentShaderID, programID, vertexShaderSource, fragmentShaderSource) {
		out.Write(atom.NoID, a)
	}

	out.Write(atom.NoID, NewGlBindAttribLocation(programID, aScreenCoordsLocation, "aScreenCoords"))
	out.Write(atom.NoID, NewGlLinkProgram(programID))
	out.Write(atom.NoID, NewGlUseProgram(programID))
	out.Write(atom.NoID, NewGlGenBuffers(1, memory.Tmp).
		AddWrite(atom.Data(arch, d, l, memory.Tmp, bufferID)))
	out.Write(atom.NoID, NewGlBindBuffer(GLenum_GL_ARRAY_BUFFER, bufferID))
	out.Write(atom.NoID, NewGlBufferData(GLenum_GL_ARRAY_BUFFER, GLsizeiptr(4*len(positions)), memory.Tmp, GLenum_GL_STATIC_DRAW).
		AddRead(atom.Data(s.Architecture, d, l, memory.Tmp, positions)))
	out.Write(atom.NoID, NewGlVertexAttribPointer(aScreenCoordsLocation, 2, GLenum_GL_FLOAT, GLboolean(0), 0, memory.Nullptr))
	out.Write(atom.NoID, NewGlDrawArrays(GLenum_GL_TRIANGLE_STRIP, 0, 4))

	t.revert()

	// Restore buffer state.
	out.Write(atom.NoID, NewGlBindBuffer(GLenum_GL_ARRAY_BUFFER, origArrayBufferID))
	out.Write(atom.NoID, NewGlDeleteBuffers(1, memory.Tmp).
		AddRead(atom.Data(arch, d, l, memory.Tmp, bufferID)))

	// Restore program state.
	out.Write(atom.NoID, NewGlUseProgram(origProgramID))
	out.Write(atom.NoID, NewGlDeleteProgram(programID))
	out.Write(atom.NoID, NewGlDeleteShader(vertexShaderID))
	out.Write(atom.NoID, NewGlDeleteShader(fragmentShaderID))

	return nil
}
