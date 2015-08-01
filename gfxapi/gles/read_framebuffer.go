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
)

func readFramebufferDepth(out chan replay.Image) atom.Atom {
	return replay.Custom(func(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error {
		arch := s.Architecture
		c := getContext(s)

		colorW, colorH, err := getState(s).getFramebufferAttachmentSize(gfxapi.FramebufferAttachmentColor)
		if err != nil {
			return err
		}
		depthW, depthH, err := getState(s).getFramebufferAttachmentSize(gfxapi.FramebufferAttachmentDepth)
		if err != nil {
			return err
		}

		const (
			uTextureLocation      UniformLocation   = 0
			aScreenCoordsLocation AttributeLocation = 0

			vertexShaderSource string = `
				precision highp float;
				attribute vec2 aScreenCoords;
				varying vec2 vTexCoords;

				void main() {
					vTexCoords = aScreenCoords / 2. + vec2(0.5, 0.5);
					gl_Position = vec4(aScreenCoords.xy, 0., 1.);
				}`
			fragmentShaderSource string = `
				precision highp float;
				uniform sampler2D uTexture;
				varying vec2 vTexCoords;

				vec4 float2rgba(float f) {
					vec4 v = fract(f * vec4(1., 255., 65025., 16581375.));
					return v - vec4(v.yzw, 0.) / 255.;
				}

				void main() {
					float sample = texture2D(uTexture, vTexCoords).r;
					gl_FragColor = float2rgba(sample);
				}`
		)

		var (
			origProgramID            = c.BoundProgram
			origRenderbufferID       = c.BoundRenderbuffers[GLenum_GL_RENDERBUFFER]
			origReadFramebufferID    = c.BoundFramebuffers[GLenum_GL_READ_FRAMEBUFFER]
			origDrawFramebufferID    = c.BoundFramebuffers[GLenum_GL_DRAW_FRAMEBUFFER]
			origTextureID            = c.TextureUnits[c.ActiveTextureUnit][GLenum_GL_TEXTURE_2D]
			origArrayBufferID        = c.BoundBuffers[GLenum_GL_ARRAY_BUFFER]
			origElementArrayBufferID = c.BoundBuffers[GLenum_GL_ELEMENT_ARRAY_BUFFER]
			origActiveTextureUnit    = int32(c.ActiveTextureUnit - GLenum_GL_TEXTURE0)

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
				NewGlDisable(capability).Replay(i, s, d, l, b)
				undoList = append(undoList, NewGlEnable(capability))
			}
		}
		if !c.VertexAttributeArrays[aScreenCoordsLocation].Enabled {
			NewGlEnableVertexAttribArray(aScreenCoordsLocation).Replay(i, s, d, l, b)
			undoList = append(undoList, NewGlDisableVertexAttribArray(aScreenCoordsLocation))
		}

		replayEach(i, s, d, l, b,
			// Setup new framebuffer/renderbuffer.
			NewGlGenFramebuffers(1, memory.Tmp).
				AddRead(atom.Data(arch, d, l, memory.Tmp, framebufferID)),
			NewGlBindFramebuffer(GLenum_GL_DRAW_FRAMEBUFFER, framebufferID),
			NewGlGenRenderbuffers(1, memory.Tmp).
				AddRead(atom.Data(arch, d, l, memory.Tmp, renderbufferID)),
			NewGlBindRenderbuffer(GLenum_GL_RENDERBUFFER, renderbufferID),
			NewGlRenderbufferStorage(GLenum_GL_RENDERBUFFER, GLenum_GL_RGBA8, outW, outH),
			NewGlFramebufferRenderbuffer(GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_COLOR_ATTACHMENT0, GLenum_GL_RENDERBUFFER, renderbufferID),

			// Setup depth texture.
			NewGlGenTextures(1, memory.Tmp).
				AddRead(atom.Data(arch, d, l, memory.Tmp, textureID)),
			NewGlBindTexture(GLenum_GL_TEXTURE_2D, textureID),
			NewGlTexImage2D(GLenum_GL_TEXTURE_2D, 0, GLenum_GL_DEPTH24_STENCIL8, outW, outH, 0, GLenum_GL_DEPTH_STENCIL, GLenum_GL_UNSIGNED_INT_24_8, memory.Nullptr),
			NewGlTexParameteri(GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_MIN_FILTER, int32(GLenum_GL_NEAREST)),
			NewGlTexParameteri(GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_MAG_FILTER, int32(GLenum_GL_NEAREST)),
			NewGlTexParameteri(GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_WRAP_S, int32(GLenum_GL_CLAMP_TO_EDGE)),
			NewGlTexParameteri(GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_WRAP_T, int32(GLenum_GL_CLAMP_TO_EDGE)),

			// Blit depth attachment.
			NewGlFramebufferTexture2D(GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_DEPTH_ATTACHMENT, GLenum_GL_TEXTURE_2D, textureID, 0),
			NewGlBlitFramebuffer(0, 0, inW, inH, 0, 0, outW, outH, GLbitfield_GL_DEPTH_BUFFER_BIT, GLenum_GL_NEAREST),
			NewGlFramebufferTexture2D(GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_DEPTH_ATTACHMENT, GLenum_GL_TEXTURE_2D, TextureId(0), 0),

			// Bind new framebuffer.
			NewGlBindFramebuffer(GLenum_GL_READ_FRAMEBUFFER, framebufferID),

			// Render depth texture to framebuffer color attachment.
			NewGlClear(GLbitfield_GL_COLOR_BUFFER_BIT),
		)

		// Create the shader program
		replayEach(i, s, d, l, b,
			NewProgram(arch, d, l, vertexShaderID, fragmentShaderID, programID, vertexShaderSource, fragmentShaderSource)...)

		replayEach(i, s, d, l, b,
			NewGlBindAttribLocation(programID, aScreenCoordsLocation, "aScreenCoords"),
			NewGlLinkProgram(programID),
			NewGlUseProgram(programID),
			NewGlBindTexture(GLenum_GL_TEXTURE_2D, textureID),
			NewGlGetUniformLocation(programID, "uTexture", uTextureLocation),
			NewGlUniform1i(uTextureLocation, origActiveTextureUnit),
			NewGlBindBuffer(GLenum_GL_ARRAY_BUFFER, 0),
			NewGlBindBuffer(GLenum_GL_ELEMENT_ARRAY_BUFFER, 0),
			NewGlVertexAttribPointer(aScreenCoordsLocation, 2, GLenum_GL_FLOAT, false, 0, memory.Tmp),
			NewGlDrawArrays(GLenum_GL_TRIANGLE_STRIP, 0, 4).
				AddRead(atom.Data(arch, d, l, memory.Tmp, positions)),
		)

		postColorData(i, s, d, l, b, outW, outH, out)

		// Restore conditionally changed state.
		replayEach(i, s, d, l, b, undoList...)

		replayEach(i, s, d, l, b,
			// Restore buffer/vertexAttrib state.
			NewGlBindBuffer(GLenum_GL_ELEMENT_ARRAY_BUFFER, origElementArrayBufferID),
			NewGlBindBuffer(GLenum_GL_ARRAY_BUFFER, origArrayBufferID),
			// Note: we're not restoring the original VertexAttribPointer as we may re-enter an inconsistent state, which would abort the current replay batch.
			// NewGlVertexAttribPointer(aScreenCoordsLocation, origVertexAttrib.Size, origVertexAttrib.Type, origVertexAttrib.Normalized, origVertexAttrib.Stride, VertexPointer(origVertexAttrib.Data)),

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

		return nil
	})
}

func readFramebufferColor(width, height uint32, out chan replay.Image) atom.Atom {
	return replay.Custom(func(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error {
		arch := s.Architecture
		c := getContext(s)

		colorW, colorH, err := getState(s).getFramebufferAttachmentSize(gfxapi.FramebufferAttachmentColor)
		if err != nil {
			return err
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

		// Generate new unused object IDs.
		renderbufferID := RenderbufferId(newUnusedID(func(x uint32) bool { _, ok := c.Instances.Renderbuffers[RenderbufferId(x)]; return ok }))
		framebufferID := FramebufferId(newUnusedID(func(x uint32) bool { _, ok := c.Instances.Framebuffers[FramebufferId(x)]; return ok }))

		if inW == outW && inH == outH {
			postColorData(i, s, d, l, b, outW, outH, out)
		} else {
			replayEach(i, s, d, l, b,
				NewGlGenFramebuffers(1, memory.Tmp).
					AddRead(atom.Data(arch, d, l, memory.Tmp, framebufferID)),
				NewGlBindFramebuffer(GLenum_GL_DRAW_FRAMEBUFFER, framebufferID),
				NewGlGenRenderbuffers(1, memory.Tmp).
					AddRead(atom.Data(arch, d, l, memory.Tmp, renderbufferID)),
				NewGlBindRenderbuffer(GLenum_GL_RENDERBUFFER, renderbufferID),
				NewGlRenderbufferStorage(GLenum_GL_RENDERBUFFER, GLenum_GL_RGBA8, outW, outH),
				NewGlFramebufferRenderbuffer(GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_COLOR_ATTACHMENT0, GLenum_GL_RENDERBUFFER, renderbufferID),
				NewGlBlitFramebuffer(0, 0, inW, inH, 0, 0, outW, outH, GLbitfield_GL_COLOR_BUFFER_BIT, GLenum_GL_LINEAR),
				NewGlBindFramebuffer(GLenum_GL_READ_FRAMEBUFFER, framebufferID),
			)

			postColorData(i, s, d, l, b, outW, outH, out)

			replayEach(i, s, d, l, b,
				NewGlBindRenderbuffer(GLenum_GL_RENDERBUFFER, origRenderbufferID),
				NewGlBindFramebuffer(GLenum_GL_READ_FRAMEBUFFER, origReadFramebufferID),
				NewGlBindFramebuffer(GLenum_GL_DRAW_FRAMEBUFFER, origDrawFramebufferID),
				NewGlDeleteRenderbuffers(1, memory.Tmp).
					AddRead(atom.Data(arch, d, l, memory.Tmp, renderbufferID)),
				NewGlDeleteFramebuffers(1, memory.Tmp).
					AddRead(atom.Data(arch, d, l, memory.Tmp, framebufferID)),
			)
		}

		return nil
	})
}

func postColorData(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder, width, height int32, img chan<- replay.Image) {
	ctx := getContext(s)
	origPackAlignment := ctx.PixelStorage[GLenum_GL_PACK_ALIGNMENT]
	if origPackAlignment != 1 {
		NewGlPixelStorei(GLenum_GL_PACK_ALIGNMENT, 1).Replay(i, s, d, l, b)
		defer NewGlPixelStorei(GLenum_GL_PACK_ALIGNMENT, origPackAlignment).Replay(i, s, d, l, b)
	}
	if origPackBuffer, ok := ctx.BoundBuffers[GLenum_GL_PIXEL_PACK_BUFFER]; ok && origPackBuffer != 0 {
		NewGlBindBuffer(GLenum_GL_PIXEL_PACK_BUFFER, 0).Replay(i, s, d, l, b)
		defer NewGlBindBuffer(GLenum_GL_PIXEL_PACK_BUFFER, origPackBuffer).Replay(i, s, d, l, b)
	}

	imageSize := uint64(width * height * 4)
	NewGlReadPixels(0, 0, width, height, GLenum_GL_RGBA, GLenum_GL_UNSIGNED_BYTE, memory.Tmp).Replay(i, s, d, l, b)
	b.Post(value.RemappedPointer(memory.Tmp.Address), imageSize, func(d binary.Decoder, err error) error {
		var data []byte
		if err == nil {
			data = make([]byte, imageSize)
			err = d.Data(data)
		}
		if err != nil {
			data = nil
		}
		img <- replay.Image{Data: data, Error: err}
		return err
	})
}

func replayEach(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder, atoms ...atom.Atom) {
	for _, a := range atoms {
		if r, ok := a.(replay.Replayer); ok {
			r.Replay(i, s, d, l, b)
		}
	}
}

func newUnusedID(existenceTest func(uint32) bool) uint32 {
	for {
		x := rand.Uint32()
		if !existenceTest(x) {
			return x
		}
	}
}
