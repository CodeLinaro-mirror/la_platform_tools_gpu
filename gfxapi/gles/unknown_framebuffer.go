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
)

// undefinedFramebuffer adds a transform that will render a pattern into the
// color buffer at the end of each frame.
func undefinedFramebuffer(d database.Database, l log.Logger) atom.Transformer {
	s := gfxapi.NewState()
	return atom.Transform("DirtyFramebuffer", func(i atom.ID, a atom.Atom, out atom.Writer) {
		a.Mutate(s, d, l)
		out.Write(i, a)
		if a.Flags().IsEndOfFrame() {
			drawUndefinedFramebuffer(a, s, d, l, out)
		}
	})
}

func drawUndefinedFramebuffer(a atom.Atom, s *gfxapi.State, d database.Database, l log.Logger, out atom.Writer) error {
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

	c := getContext(s)

	var (
		origProgramID            = c.BoundProgram
		origArrayBufferID        = c.BoundBuffers[GLenum_GL_ARRAY_BUFFER]
		origElementArrayBufferID = c.BoundBuffers[GLenum_GL_ELEMENT_ARRAY_BUFFER]
		origVertexAttrib         = *(c.VertexAttributeArrays[aScreenCoordsLocation])
	)

	// Generate new unused object IDs.
	programID := ProgramId(newUnusedID(func(x uint32) bool { _, ok := c.Instances.Programs[ProgramId(x)]; return ok }))
	vertexShaderID := ShaderId(newUnusedID(func(x uint32) bool { _, ok := c.Instances.Shaders[ShaderId(x)]; return ok }))
	fragmentShaderID := ShaderId(newUnusedID(func(x uint32) bool {
		_, ok := c.Instances.Shaders[ShaderId(x)]
		return ok || ShaderId(x) == vertexShaderID
	}))

	// 2D vertices positions for a full screen 2D triangle strip.
	positions := []float32{-1., -1., 1., -1., -1., 1., 1., 1.}

	// Temporarily change rasterizing/blending state and enable VAP 0.
	undoList := []atom.Atom{}
	for _, cap := range []GLenum{
		GLenum_GL_BLEND,
		GLenum_GL_CULL_FACE,
		GLenum_GL_DEPTH_TEST,
		GLenum_GL_SCISSOR_TEST,
		GLenum_GL_STENCIL_TEST,
	} {
		capability := cap
		if c.Capabilities[capability] {
			out.Write(atom.NoID, NewGlDisable(capability))
			undoList = append(undoList, NewGlEnable(capability))
		}
	}
	if !c.VertexAttributeArrays[aScreenCoordsLocation].Enabled {
		out.Write(atom.NoID, NewGlEnableVertexAttribArray(aScreenCoordsLocation))
		undoList = append(undoList, NewGlDisableVertexAttribArray(aScreenCoordsLocation))
	}

	// Create the shader program
	for _, a := range NewProgram(s.Architecture, d, l, vertexShaderID, fragmentShaderID, programID, vertexShaderSource, fragmentShaderSource) {
		out.Write(atom.NoID, a)
	}

	out.Write(atom.NoID, NewGlBindAttribLocation(programID, aScreenCoordsLocation, "aScreenCoords"))
	out.Write(atom.NoID, NewGlLinkProgram(programID))
	out.Write(atom.NoID, NewGlUseProgram(programID))
	out.Write(atom.NoID, NewGlBindBuffer(GLenum_GL_ARRAY_BUFFER, 0))
	out.Write(atom.NoID, NewGlBindBuffer(GLenum_GL_ELEMENT_ARRAY_BUFFER, 0))
	out.Write(atom.NoID, NewGlVertexAttribPointer(aScreenCoordsLocation, 2, GLenum_GL_FLOAT, false, 0, memory.Tmp))
	out.Write(atom.NoID, NewGlDrawArrays(GLenum_GL_TRIANGLE_STRIP, 0, 4).
		AddRead(atom.Data(s.Architecture, d, l, memory.Tmp, positions)))

	// Restore conditionally changed state.
	for _, a := range undoList {
		out.Write(atom.NoID, a)
	}

	// Restore buffer/vertexAttrib state.
	if origVertexAttrib.Buffer != 0 || origVertexAttrib.Pointer.Address != 0 {
		out.Write(atom.NoID, NewGlBindBuffer(GLenum_GL_ARRAY_BUFFER, origVertexAttrib.Buffer))
		out.Write(atom.NoID, NewGlVertexAttribPointer(aScreenCoordsLocation, GLint(origVertexAttrib.Size), origVertexAttrib.Type, origVertexAttrib.Normalized, origVertexAttrib.Stride, origVertexAttrib.Pointer.Pointer))
	}
	out.Write(atom.NoID, NewGlBindBuffer(GLenum_GL_ELEMENT_ARRAY_BUFFER, origElementArrayBufferID))
	out.Write(atom.NoID, NewGlBindBuffer(GLenum_GL_ARRAY_BUFFER, origArrayBufferID))

	// Restore program state.
	out.Write(atom.NoID, NewGlUseProgram(origProgramID))
	out.Write(atom.NoID, NewGlDeleteProgram(programID))
	out.Write(atom.NoID, NewGlDeleteShader(vertexShaderID))
	out.Write(atom.NoID, NewGlDeleteShader(fragmentShaderID))

	return nil
}
