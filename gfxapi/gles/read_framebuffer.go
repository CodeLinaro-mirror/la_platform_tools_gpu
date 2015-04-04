package gles

import (
	"bytes"
	"math/rand"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/endian"
	"android.googlesource.com/platform/tools/gpu/binary/flat"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/database/store"
	"android.googlesource.com/platform/tools/gpu/gfxapi/state"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/replay/builder"
	"android.googlesource.com/platform/tools/gpu/replay/value"
)

// These must conform to the replay.Replayer interface.
var _ = replay.Replayer(readFramebufferDepth{})
var _ = replay.Replayer(readFramebufferColor{})

// readFramebufferDepth is an atom used to postback the content of the currently
// bound framebuffer's depth attachment.
type readFramebufferDepth struct {
	binary.Generate `disable:"true"`
	contextID       atom.ContextID
	database        database.Database
}

func (a readFramebufferDepth) ContextID() atom.ContextID { return a.contextID }
func (a readFramebufferDepth) TypeID() atom.TypeID       { return 0 }
func (a readFramebufferDepth) Flags() atom.Flags         { return 0 }

func (a readFramebufferDepth) Replay(id atom.ID, s *state.State, b *builder.Builder, wantOutput bool) {
	defer b.EndAtom()

	c := getState(a, s)
	cid := a.ContextID()
	colorW, colorH, err := c.GetFramebufferAttachmentSize(state.FramebufferAttachmentColor)
	if err != nil {
		return
	}
	depthW, depthH, err := c.GetFramebufferAttachmentSize(state.FramebufferAttachmentDepth)
	if err != nil {
		return
	}

	const (
		uTextureLocation      UniformLocation   = 0
		aScreenCoordsLocation AttributeLocation = 0

		// TODO: Add a way to allocate memory from a free range for injected observations.
		positionsAddr memory.Pointer = 0xffffffffdeadbead
		indicesAddr   memory.Pointer = 0xffffffffdeadbeef

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
		origRenderbufferID       = c.BoundRenderbuffers[RenderbufferTarget_GL_RENDERBUFFER]
		origReadFramebufferID    = c.BoundFramebuffers[FramebufferTarget_GL_READ_FRAMEBUFFER]
		origDrawFramebufferID    = c.BoundFramebuffers[FramebufferTarget_GL_DRAW_FRAMEBUFFER]
		origTextureID            = c.TextureUnits[c.ActiveTextureUnit][TextureTarget_GL_TEXTURE_2D]
		origArrayBufferID        = c.BoundBuffers[BufferTarget_GL_ARRAY_BUFFER]
		origElementArrayBufferID = c.BoundBuffers[BufferTarget_GL_ELEMENT_ARRAY_BUFFER]
		origActiveTextureUnit    = int32(c.ActiveTextureUnit - TextureUnit_GL_TEXTURE0)

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

	// Map vertex attrib and indice resources.
	var buffer bytes.Buffer
	enc := flat.Encoder(endian.Writer(&buffer, endian.Little))
	for _, f := range []float32{-1., -1., 1., -1., -1., 1., 1., 1.} {
		enc.Float32(f)
	}
	positionsData := store.Blob{Data: buffer.Bytes()}
	positionsDataId, err := a.database.Store(&positionsData, log.Nop{})
	if err != nil {
		return
	}
	indicesData := store.Blob{Data: []byte{0, 1, 2, 3}}
	indicesDataId, err := a.database.Store(&indicesData, log.Nop{})
	if err != nil {
		return
	}

	// Temporarily change rasterizing/blending state and enable VAP 0.
	undoList := []atom.Atom{}
	for _, cap := range []Capability{
		Capability_GL_BLEND,
		Capability_GL_DEPTH_TEST,
		Capability_GL_STENCIL_TEST,
		Capability_GL_CULL_FACE,
	} {
		capability := cap
		if c.Capabilities[capability] {
			replayNoPost(id, s, b, NewGlDisable(cid, capability))
			undoList = append(undoList, NewGlEnable(cid, capability))
		}
	}
	if !c.VertexAttributeArrays[aScreenCoordsLocation].Enabled {
		replayNoPost(id, s, b, NewGlEnableVertexAttribArray(cid, aScreenCoordsLocation))
		undoList = append(undoList, NewGlDisableVertexAttribArray(cid, aScreenCoordsLocation))
	}

	replayNoPost(id, s, b,
		// Setup new framebuffer/renderbuffer.
		NewGlGenFramebuffers(cid, 1, FramebufferIdArray{framebufferID}),
		NewGlBindFramebuffer(cid, FramebufferTarget_GL_DRAW_FRAMEBUFFER, framebufferID),
		NewGlGenRenderbuffers(cid, 1, RenderbufferIdArray{renderbufferID}),
		NewGlBindRenderbuffer(cid, RenderbufferTarget_GL_RENDERBUFFER, renderbufferID),
		NewGlRenderbufferStorage(cid, RenderbufferTarget_GL_RENDERBUFFER, RenderbufferFormat_GL_RGBA8, outW, outH),
		NewGlFramebufferRenderbuffer(cid, FramebufferTarget_GL_DRAW_FRAMEBUFFER, FramebufferAttachment_GL_COLOR_ATTACHMENT0, RenderbufferTarget_GL_RENDERBUFFER, renderbufferID),

		// Setup depth texture.
		NewGlGenTextures(cid, 1, TextureIdArray{textureID}),
		NewGlBindTexture(cid, TextureTarget_GL_TEXTURE_2D, textureID),
		NewGlTexImage2D(cid, TextureImageTarget_GL_TEXTURE_2D, 0, TexelFormat_GL_DEPTH24_STENCIL8, outW, outH, 0, TexelFormat_GL_DEPTH_STENCIL, TexelType_GL_UNSIGNED_INT_24_8, TexturePointer(0)),
		NewGlTexParameteri(cid, TextureTarget_GL_TEXTURE_2D, TextureParameter_GL_TEXTURE_MIN_FILTER, int32(TextureFilterMode_GL_NEAREST)),
		NewGlTexParameteri(cid, TextureTarget_GL_TEXTURE_2D, TextureParameter_GL_TEXTURE_MAG_FILTER, int32(TextureFilterMode_GL_NEAREST)),
		NewGlTexParameteri(cid, TextureTarget_GL_TEXTURE_2D, TextureParameter_GL_TEXTURE_WRAP_S, int32(TextureWrapMode_GL_CLAMP_TO_EDGE)),
		NewGlTexParameteri(cid, TextureTarget_GL_TEXTURE_2D, TextureParameter_GL_TEXTURE_WRAP_T, int32(TextureWrapMode_GL_CLAMP_TO_EDGE)),

		// Blit depth attachment.
		NewGlFramebufferTexture2D(cid, FramebufferTarget_GL_DRAW_FRAMEBUFFER, FramebufferAttachment_GL_DEPTH_ATTACHMENT, TextureImageTarget_GL_TEXTURE_2D, textureID, 0),
		NewGlBlitFramebuffer(cid, 0, 0, inW, inH, 0, 0, outW, outH, ClearMask_GL_DEPTH_BUFFER_BIT, TextureFilterMode_GL_NEAREST),
		NewGlFramebufferTexture2D(cid, FramebufferTarget_GL_DRAW_FRAMEBUFFER, FramebufferAttachment_GL_DEPTH_ATTACHMENT, TextureImageTarget_GL_TEXTURE_2D, TextureId(0), 0),

		// Bind new framebuffer.
		NewGlBindFramebuffer(cid, FramebufferTarget_GL_READ_FRAMEBUFFER, framebufferID),

		// Render depth texture to framebuffer color attachment.
		NewGlClear(cid, ClearMask_GL_COLOR_BUFFER_BIT),
		NewGlCreateShader(cid, ShaderType_GL_VERTEX_SHADER, vertexShaderID),
		NewGlShaderSource(cid, vertexShaderID, 1, StringArray{vertexShaderSource}, S32Array{int32(len(vertexShaderSource))}),
		NewGlCompileShader(cid, vertexShaderID),
		NewGlCreateShader(cid, ShaderType_GL_FRAGMENT_SHADER, fragmentShaderID),
		NewGlShaderSource(cid, fragmentShaderID, 1, StringArray{fragmentShaderSource}, S32Array{int32(len(fragmentShaderSource))}),
		NewGlCompileShader(cid, fragmentShaderID),
		NewGlCreateProgram(cid, programID),
		NewGlAttachShader(cid, programID, vertexShaderID),
		NewGlAttachShader(cid, programID, fragmentShaderID),
		NewGlBindAttribLocation(cid, programID, aScreenCoordsLocation, "aScreenCoords"),
		NewGlLinkProgram(cid, programID),
		NewGlUseProgram(cid, programID),
		NewGlBindTexture(cid, TextureTarget_GL_TEXTURE_2D, textureID),
		NewGlGetUniformLocation(cid, programID, "uTexture", uTextureLocation),
		NewGlUniform1i(cid, uTextureLocation, origActiveTextureUnit),
		NewGlBindBuffer(cid, BufferTarget_GL_ARRAY_BUFFER, 0),
		NewGlBindBuffer(cid, BufferTarget_GL_ELEMENT_ARRAY_BUFFER, 0),
		NewGlVertexAttribPointer(cid, aScreenCoordsLocation, 2, VertexAttribType_GL_FLOAT, false, 0, VertexPointer(positionsAddr)),
		&atom.Observation{Context: cid, Range: memory.Range{Base: positionsAddr, Size: 8 * 4}, ResourceID: positionsDataId},
		&atom.Observation{Context: cid, Range: memory.Range{Base: indicesAddr, Size: 4}, ResourceID: indicesDataId},
		NewGlDrawElements(cid, DrawMode_GL_TRIANGLE_STRIP, 4, IndicesType_GL_UNSIGNED_BYTE, IndicesPointer(indicesAddr)),
	)

	postColorData(id, b, s, outW, outH, cid)

	// Restore conditionally changed state.
	replayNoPost(id, s, b, undoList...)

	replayNoPost(id, s, b,
		// Restore buffer/vertexAttrib state.
		NewGlBindBuffer(cid, BufferTarget_GL_ELEMENT_ARRAY_BUFFER, origElementArrayBufferID),
		NewGlBindBuffer(cid, BufferTarget_GL_ARRAY_BUFFER, origArrayBufferID),
		// Note: we're not restoring the original VertexAttribPointer as we may re-enter an inconsistent state, which would abort the current replay batch.
		// NewGlVertexAttribPointer(cid, aScreenCoordsLocation, origVertexAttrib.Size, origVertexAttrib.Type, origVertexAttrib.Normalized, origVertexAttrib.Stride, VertexPointer(origVertexAttrib.Data)),

		// Restore texture state.
		NewGlBindTexture(cid, TextureTarget_GL_TEXTURE_2D, origTextureID),
		NewGlDeleteTextures(cid, 1, TextureIdArray{textureID}),

		// Restore framebuffer/renderbuffer state.
		NewGlBindRenderbuffer(cid, RenderbufferTarget_GL_RENDERBUFFER, origRenderbufferID),
		NewGlBindFramebuffer(cid, FramebufferTarget_GL_READ_FRAMEBUFFER, origReadFramebufferID),
		NewGlBindFramebuffer(cid, FramebufferTarget_GL_DRAW_FRAMEBUFFER, origDrawFramebufferID),
		NewGlDeleteRenderbuffers(cid, 1, RenderbufferIdArray{renderbufferID}),
		NewGlDeleteFramebuffers(cid, 1, FramebufferIdArray{framebufferID}),
	)
}

// readFramebufferColor is an atom used to postback the content of the currently
// bound framebuffer's color attachment.
type readFramebufferColor struct {
	binary.Generate `disable:"true"`
	contextID       atom.ContextID
	width, height   uint32
}

func (a readFramebufferColor) ContextID() atom.ContextID { return a.contextID }
func (a readFramebufferColor) TypeID() atom.TypeID       { return 0 }
func (a readFramebufferColor) Flags() atom.Flags         { return 0 }

func (a readFramebufferColor) Replay(id atom.ID, s *state.State, b *builder.Builder, wantOutput bool) {
	defer b.EndAtom()

	c := getState(a, s)
	cid := a.ContextID()
	colorW, colorH, err := c.GetFramebufferAttachmentSize(state.FramebufferAttachmentColor)
	if err != nil {
		return
	}

	var (
		origRenderbufferID    = c.BoundRenderbuffers[RenderbufferTarget_GL_RENDERBUFFER]
		origReadFramebufferID = c.BoundFramebuffers[FramebufferTarget_GL_READ_FRAMEBUFFER]
		origDrawFramebufferID = c.BoundFramebuffers[FramebufferTarget_GL_DRAW_FRAMEBUFFER]

		inW  = int32(colorW)
		inH  = int32(colorH)
		outW = int32(a.width)
		outH = int32(a.height)
	)

	// Generate new unused object IDs.
	renderbufferID := RenderbufferId(newUnusedID(func(x uint32) bool { _, ok := c.Instances.Renderbuffers[RenderbufferId(x)]; return ok }))
	framebufferID := FramebufferId(newUnusedID(func(x uint32) bool { _, ok := c.Instances.Framebuffers[FramebufferId(x)]; return ok }))

	if inW == outW && inH == outH {
		postColorData(id, b, s, outW, outH, cid)
	} else {
		replayNoPost(id, s, b,
			NewGlGenFramebuffers(cid, 1, FramebufferIdArray{framebufferID}),
			NewGlBindFramebuffer(cid, FramebufferTarget_GL_DRAW_FRAMEBUFFER, framebufferID),
			NewGlGenRenderbuffers(cid, 1, RenderbufferIdArray{renderbufferID}),
			NewGlBindRenderbuffer(cid, RenderbufferTarget_GL_RENDERBUFFER, renderbufferID),
			NewGlRenderbufferStorage(cid, RenderbufferTarget_GL_RENDERBUFFER, RenderbufferFormat_GL_RGBA8, outW, outH),
			NewGlFramebufferRenderbuffer(cid, FramebufferTarget_GL_DRAW_FRAMEBUFFER, FramebufferAttachment_GL_COLOR_ATTACHMENT0, RenderbufferTarget_GL_RENDERBUFFER, renderbufferID),
			NewGlBlitFramebuffer(cid, 0, 0, inW, inH, 0, 0, outW, outH, ClearMask_GL_COLOR_BUFFER_BIT, TextureFilterMode_GL_LINEAR),
			NewGlBindFramebuffer(cid, FramebufferTarget_GL_READ_FRAMEBUFFER, framebufferID),
		)

		postColorData(id, b, s, outW, outH, cid)

		replayNoPost(id, s, b,
			NewGlBindRenderbuffer(cid, RenderbufferTarget_GL_RENDERBUFFER, origRenderbufferID),
			NewGlBindFramebuffer(cid, FramebufferTarget_GL_READ_FRAMEBUFFER, origReadFramebufferID),
			NewGlBindFramebuffer(cid, FramebufferTarget_GL_DRAW_FRAMEBUFFER, origDrawFramebufferID),
			NewGlDeleteRenderbuffers(cid, 1, RenderbufferIdArray{renderbufferID}),
			NewGlDeleteFramebuffers(cid, 1, FramebufferIdArray{framebufferID}),
		)
	}
}

func postColorData(id atom.ID, b *builder.Builder, s *state.State, width, height int32, cid atom.ContextID) {
	c := s.Contexts[cid].(*State)
	origPackAlignment := c.PixelStorage[PixelStoreParameter_GL_PACK_ALIGNMENT]
	if origPackAlignment != 1 {
		replayNoPost(id, s, b, NewGlPixelStorei(cid, PixelStoreParameter_GL_PACK_ALIGNMENT, 1))
		defer replayNoPost(id, s, b, NewGlPixelStorei(cid, PixelStoreParameter_GL_PACK_ALIGNMENT, origPackAlignment))
	}

	imageSize := uint64(width * height * 4)
	addr := b.AllocateTemporaryMemory(imageSize)

	// glReadPixels(0, 0, width, height, GL_RGBA, GL_UNSIGNED_BYTE, addr)
	b.Push(value.S32(0))
	b.Push(value.S32(0))
	b.Push(value.S32(width))
	b.Push(value.S32(height))
	b.Push(value.U32(TexelFormat_GL_RGBA))
	b.Push(value.U32(TexelType_GL_UNSIGNED_BYTE))
	b.Push(addr)
	b.CallNoPush(funcInfoGlReadPixels)

	b.Post(addr, imageSize, id,
		func(d binary.Decoder) (interface{}, error) {
			buf := make([]byte, imageSize)
			err := d.Data(buf)
			return buf, err
		},
	)
}

func replayNoPost(id atom.ID, s *state.State, b *builder.Builder, atoms ...atom.Atom) {
	for _, a := range atoms {
		replay.Replay(id, a, s, b, false)
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
