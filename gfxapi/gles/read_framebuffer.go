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
	"math/rand"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/replay/builder"
	"android.googlesource.com/platform/tools/gpu/replay/value"
	"android.googlesource.com/platform/tools/gpu/service"
)

type readFramebuffer struct {
	state      *gfxapi.State
	database   database.Database
	logger     log.Logger
	injections map[atom.ID][]func(out atom.Writer)
}

func newReadFramebuffer(d database.Database, l log.Logger) *readFramebuffer {
	return &readFramebuffer{
		state:      gfxapi.NewState(),
		database:   d,
		logger:     l,
		injections: make(map[atom.ID][]func(out atom.Writer)),
	}
}

func (t *readFramebuffer) Transform(id atom.ID, a atom.Atom, out atom.Writer) {
	if err := a.Mutate(t.state, t.database, t.logger); err != nil {
		log.Errorf(t.logger, "%v", err)
	}
	out.Write(id, a)
	if r, ok := t.injections[id]; ok {
		for _, injection := range r {
			injection(out)
		}
		delete(t.injections, id)
	}
}

func (t *readFramebuffer) Flush(out atom.Writer) {}

func (t *readFramebuffer) Depth(id atom.ID, device *service.Device, img chan replay.Image) {
	t.injections[id] = append(t.injections[id], func(out atom.Writer) {
		s, d, l := t.state, t.database, t.logger
		arch := s.Architecture
		c := getContext(s)
		version, _ := ParseVersion(device.Version)

		colorW, colorH, err := getState(s).getFramebufferAttachmentSize(gfxapi.FramebufferAttachmentColor)
		if err != nil {
			log.Errorf(l, "%v", err)
			return
		}
		depthW, depthH, err := getState(s).getFramebufferAttachmentSize(gfxapi.FramebufferAttachmentDepth)
		if err != nil {
			log.Errorf(l, "%v", err)
			return
		}

		const (
			uTextureLocation      UniformLocation   = 0
			aScreenCoordsLocation AttributeLocation = 0
			vsSource                                = `
				#version 110
				precision highp float;

				attribute vec2 aScreenCoords;
				varying vec2 vTexCoords;

				void main() {
					vTexCoords = aScreenCoords / 2.0 + vec2(0.5, 0.5);
					gl_Position = vec4(aScreenCoords.xy, 0.0, 1.0);
				}`
			fsSource = `
				#version 110
				precision highp float;

				uniform sampler2D uTexture;
				varying vec2 vTexCoords;

				vec4 float2rgba(float f) {
					vec4 v = vec4(f, 0.0, 0.0, 0.0) + fract(f * vec4(0.0, 255.0, 65025.0, 16581375.0));
					return v - vec4(v.yzw, 0.0) / 255.0;
				}

				void main() {
					float v = texture2D(uTexture, vTexCoords).r;
					gl_FragColor = float2rgba(v);
				}`
		)
		var (
			origProgramID         = c.BoundProgram
			origRenderbufferID    = c.BoundRenderbuffers[GLenum_GL_RENDERBUFFER]
			origReadFramebufferID = c.BoundFramebuffers[GLenum_GL_READ_FRAMEBUFFER]
			origDrawFramebufferID = c.BoundFramebuffers[GLenum_GL_DRAW_FRAMEBUFFER]
			origTextureID         = c.TextureUnits[c.ActiveTextureUnit].Bindings[GLenum_GL_TEXTURE_2D]
			origArrayBufferID     = c.BoundBuffers[GLenum_GL_ARRAY_BUFFER]
			origActiveTextureUnit = int32(c.ActiveTextureUnit - GLenum_GL_TEXTURE0)

			inW  = int32(depthW)
			inH  = int32(depthH)
			outW = int32(colorW)
			outH = int32(colorH)
		)

		// Generate new unused object IDs.
		renderbufferID := RenderbufferId(newUnusedID(func(x uint32) bool { _, ok := c.Instances.Renderbuffers[RenderbufferId(x)]; return ok }))
		framebufferID := FramebufferId(newUnusedID(func(x uint32) bool { _, ok := c.Instances.Framebuffers[FramebufferId(x)]; return ok }))
		textureID := TextureId(newUnusedID(func(x uint32) bool { _, ok := c.Instances.Textures[TextureId(x)]; return ok }))
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
		t.glDisable(GLenum_GL_ALPHA_TEST)
		t.bindOrSaveVertexArray(version, arrayID, aScreenCoordsLocation)
		out.Write(atom.NoID, NewGlEnableVertexAttribArray(aScreenCoordsLocation))

		writeEach(out,
			// Setup new framebuffer/renderbuffer.
			NewGlGenFramebuffers(1, memory.Tmp).
				AddRead(atom.Data(arch, d, l, memory.Tmp, framebufferID)),
			NewGlGenRenderbuffers(1, memory.Tmp).
				AddRead(atom.Data(arch, d, l, memory.Tmp, renderbufferID)),
			NewGlBindRenderbuffer(GLenum_GL_RENDERBUFFER, renderbufferID),
			NewGlRenderbufferStorage(GLenum_GL_RENDERBUFFER, GLenum_GL_RGBA8, GLsizei(outW), GLsizei(outH)),
			NewGlBindFramebuffer(GLenum_GL_DRAW_FRAMEBUFFER, framebufferID),
			NewGlFramebufferRenderbuffer(GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_COLOR_ATTACHMENT0, GLenum_GL_RENDERBUFFER, renderbufferID),

			// Setup depth texture.
			NewGlGenTextures(1, memory.Tmp).
				AddRead(atom.Data(arch, d, l, memory.Tmp, textureID)),
			NewGlBindTexture(GLenum_GL_TEXTURE_2D, textureID),
			NewGlTexImage2D(GLenum_GL_TEXTURE_2D, 0, GLint(GLenum_GL_DEPTH24_STENCIL8), GLsizei(outW), GLsizei(outH), 0, GLenum_GL_DEPTH_STENCIL, GLenum_GL_UNSIGNED_INT_24_8, memory.Nullptr),
			NewGlTexParameteri(GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_MIN_FILTER, GLint(GLenum_GL_NEAREST)),
			NewGlTexParameteri(GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_MAG_FILTER, GLint(GLenum_GL_NEAREST)),
			NewGlTexParameteri(GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_WRAP_S, GLint(GLenum_GL_CLAMP_TO_EDGE)),
			NewGlTexParameteri(GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_WRAP_T, GLint(GLenum_GL_CLAMP_TO_EDGE)),

			// Blit depth attachment.
			NewGlFramebufferTexture2D(GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_DEPTH_ATTACHMENT, GLenum_GL_TEXTURE_2D, textureID, 0),
			NewGlBlitFramebuffer(0, 0, GLint(inW), GLint(inH), 0, 0, GLint(outW), GLint(outH), GLbitfield_GL_DEPTH_BUFFER_BIT, GLenum_GL_NEAREST),
			NewGlFramebufferTexture2D(GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_DEPTH_ATTACHMENT, GLenum_GL_TEXTURE_2D, TextureId(0), 0),
		)

		// Create the shader program
		writeEach(out,
			BuildProgram(arch, d, l, vertexShaderID, fragmentShaderID, programID, vsSource, fsSource)...)

		writeEach(out,
			NewGlBindAttribLocation(programID, aScreenCoordsLocation, "aScreenCoords"),
			NewGlLinkProgram(programID),
			NewGlUseProgram(programID),
			NewGlBindTexture(GLenum_GL_TEXTURE_2D, textureID),
			NewGlGetUniformLocation(programID, "uTexture", uTextureLocation),
			NewGlUniform1i(uTextureLocation, GLint(origActiveTextureUnit)),
			NewGlGenBuffers(1, memory.Tmp).
				AddWrite(atom.Data(arch, d, l, memory.Tmp, bufferID)),
			NewGlBindBuffer(GLenum_GL_ARRAY_BUFFER, bufferID),
			NewGlBufferData(GLenum_GL_ARRAY_BUFFER, GLsizeiptr(4*len(positions)), memory.Tmp, GLenum_GL_STATIC_DRAW).
				AddRead(atom.Data(s.Architecture, d, l, memory.Tmp, positions)),
			NewGlVertexAttribPointer(aScreenCoordsLocation, 2, GLenum_GL_FLOAT, GLboolean(0), 0, memory.Nullptr),
			NewGlDrawArrays(GLenum_GL_TRIANGLE_STRIP, 0, 4),

			// Bind new framebuffer for reading.
			NewGlBindFramebuffer(GLenum_GL_READ_FRAMEBUFFER, framebufferID),
		)

		postColorData(s, outW, outH, out, img)

		// Restore conditionally changed state.
		t.revert()

		writeEach(out,
			// Restore buffer/vertexAttrib state.
			NewGlBindBuffer(GLenum_GL_ARRAY_BUFFER, origArrayBufferID),
			NewGlDeleteBuffers(1, memory.Tmp).
				AddRead(atom.Data(arch, d, l, memory.Tmp, bufferID)),

			// Restore texture state.
			NewGlBindTexture(GLenum_GL_TEXTURE_2D, origTextureID),
			NewGlDeleteTextures(1, memory.Tmp).
				AddRead(atom.Data(arch, d, l, memory.Tmp, textureID)),

			// Restore framebuffer/renderbuffer state.
			NewGlBindRenderbuffer(GLenum_GL_RENDERBUFFER, origRenderbufferID),
			NewGlBindFramebuffer(GLenum_GL_READ_FRAMEBUFFER, origReadFramebufferID),
			NewGlBindFramebuffer(GLenum_GL_DRAW_FRAMEBUFFER, origDrawFramebufferID),
			NewGlDeleteRenderbuffers(1, memory.Tmp).
				AddRead(atom.Data(arch, d, l, memory.Tmp, renderbufferID)),
			NewGlDeleteFramebuffers(1, memory.Tmp).
				AddRead(atom.Data(arch, d, l, memory.Tmp, framebufferID)),

			// Restore program state.
			NewGlUseProgram(origProgramID),
			NewGlDeleteProgram(programID),
			NewGlDeleteShader(vertexShaderID),
			NewGlDeleteShader(fragmentShaderID),
		)
	})
}

func (t *readFramebuffer) Color(id atom.ID, width, height uint32, img chan replay.Image) {
	t.injections[id] = append(t.injections[id], func(out atom.Writer) {
		s, d, l := t.state, t.database, t.logger
		arch := s.Architecture
		c := getContext(s)

		colorW, colorH, err := getState(s).getFramebufferAttachmentSize(gfxapi.FramebufferAttachmentColor)
		if err != nil {
			log.Errorf(l, "%v", err)
			return
		}

		var (
			origRenderbufferID    = c.BoundRenderbuffers[GLenum_GL_RENDERBUFFER]
			origReadFramebufferID = c.BoundFramebuffers[GLenum_GL_READ_FRAMEBUFFER]
			origDrawFramebufferID = c.BoundFramebuffers[GLenum_GL_DRAW_FRAMEBUFFER]

			inW  = int32(colorW)
			inH  = int32(colorH)
			outW = int32(width)
			outH = int32(height)
		)

		if inW == outW && inH == outH {
			postColorData(s, outW, outH, out, img)
		} else {
			// Generate new unused object IDs.
			renderbufferID := RenderbufferId(newUnusedID(func(x uint32) bool { _, ok := c.Instances.Renderbuffers[RenderbufferId(x)]; return ok }))
			framebufferID := FramebufferId(newUnusedID(func(x uint32) bool { _, ok := c.Instances.Framebuffers[FramebufferId(x)]; return ok }))

			ctx := getContext(s)
			origScissor := ctx.Rasterizing.Scissor

			writeEach(out,
				NewGlScissor(0, 0, GLsizei(colorW), GLsizei(colorH)),
				NewGlGenFramebuffers(1, memory.Tmp).
					AddRead(atom.Data(arch, d, l, memory.Tmp, framebufferID)),
				NewGlBindFramebuffer(GLenum_GL_DRAW_FRAMEBUFFER, framebufferID),
				NewGlGenRenderbuffers(1, memory.Tmp).
					AddRead(atom.Data(arch, d, l, memory.Tmp, renderbufferID)),
				NewGlBindRenderbuffer(GLenum_GL_RENDERBUFFER, renderbufferID),
				NewGlRenderbufferStorage(GLenum_GL_RENDERBUFFER, GLenum_GL_RGBA8, GLsizei(outW), GLsizei(outH)),
				NewGlFramebufferRenderbuffer(GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_COLOR_ATTACHMENT0, GLenum_GL_RENDERBUFFER, renderbufferID),
				NewGlBlitFramebuffer(0, 0, GLint(inW), GLint(inH), 0, 0, GLint(outW), GLint(outH), GLbitfield_GL_COLOR_BUFFER_BIT, GLenum_GL_LINEAR),
				NewGlBindFramebuffer(GLenum_GL_READ_FRAMEBUFFER, framebufferID),
			)

			postColorData(s, outW, outH, out, img)

			writeEach(out,
				NewGlBindRenderbuffer(GLenum_GL_RENDERBUFFER, origRenderbufferID),
				NewGlBindFramebuffer(GLenum_GL_READ_FRAMEBUFFER, origReadFramebufferID),
				NewGlBindFramebuffer(GLenum_GL_DRAW_FRAMEBUFFER, origDrawFramebufferID),
				NewGlDeleteRenderbuffers(1, memory.Tmp).
					AddRead(atom.Data(arch, d, l, memory.Tmp, renderbufferID)),
				NewGlDeleteFramebuffers(1, memory.Tmp).
					AddRead(atom.Data(arch, d, l, memory.Tmp, framebufferID)),
				NewGlScissor(origScissor.X, origScissor.Y, origScissor.Width, origScissor.Height),
			)
		}
	})
}

func postColorData(s *gfxapi.State, width, height int32, out atom.Writer, img chan<- replay.Image) {
	ctx := getContext(s)
	origPackAlignment := ctx.PixelStorage[GLenum_GL_PACK_ALIGNMENT]
	if origPackAlignment != 1 {
		out.Write(atom.NoID, NewGlPixelStorei(GLenum_GL_PACK_ALIGNMENT, 1))
		defer out.Write(atom.NoID, NewGlPixelStorei(GLenum_GL_PACK_ALIGNMENT, origPackAlignment))
	}
	if origPackBuffer, ok := ctx.BoundBuffers[GLenum_GL_PIXEL_PACK_BUFFER]; ok && origPackBuffer != 0 {
		out.Write(atom.NoID, NewGlBindBuffer(GLenum_GL_PIXEL_PACK_BUFFER, 0))
		defer out.Write(atom.NoID, NewGlBindBuffer(GLenum_GL_PIXEL_PACK_BUFFER, origPackBuffer))
	}

	imageSize := uint64(width * height * 4)
	out.Write(atom.NoID, NewGlReadPixels(0, 0, GLsizei(width), GLsizei(height), GLenum_GL_RGBA, GLenum_GL_UNSIGNED_BYTE, memory.Tmp))
	out.Write(atom.NoID, replay.Custom(func(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error {
		b.Post(value.RemappedPointer(memory.Tmp.Address), imageSize, func(d binary.Decoder, err error) error {
			var data []byte
			if err == nil {
				data = make([]byte, imageSize)
				err = d.Data(data)
			}
			if err != nil {
				err = fmt.Errorf("Could not read framebuffer data (expected length %d bytes): %v", imageSize, err)
				data = nil
			}
			img <- replay.Image{Data: data, Error: err}
			return err
		})
		return nil
	}))
}

func writeEach(out atom.Writer, atoms ...atom.Atom) {
	for _, a := range atoms {
		out.Write(atom.NoID, a)
	}
}

func newUnusedID(existenceTest func(uint32) bool) uint32 {
	for {
		x := rand.Uint32()
		if !existenceTest(x) && x != 0 {
			return x
		}
	}
}
