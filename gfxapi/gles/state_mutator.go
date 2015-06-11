////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

package gles

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
)

func getState(s *gfxapi.State) *State {
	api := API()
	if state, ok := s.APIs[api].(*State); ok {
		return state
	} else {
		if s.APIs == nil {
			s.APIs = make(map[gfxapi.API]interface{})
		}
		state = &State{}
		state.Init()
		s.APIs[api] = state
		return state
	}
}

func (ϟa *EglInitialize) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	if (ϟa.Major) != (EGLintᵖ{}) {
		ϟa.Major.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Major.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
	}
	if (ϟa.Major) != (EGLintᵖ{}) {
		ϟa.Minor.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Minor.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
	}
	ϟa.Result = ϟa.Result
	return nil
}
func (ϟa *EglCreateContext) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	context := EGLContext(ϟa.Result) // EGLContext
	identifier := ϟc.NextContextID   // ContextID
	ϟc.NextContextID = (ϟc.NextContextID) + (ContextID(uint32(1)))
	ctx := func() *Context {
		s := &Context{}
		s.Init()
		return s
	}() // Contextʳ
	ctx.Identifier = identifier
	ctx.Instances.Buffers[BufferId(uint32(0))] = func() *Buffer {
		s := &Buffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Textures[TextureId(uint32(0))] = func() *Texture {
		s := &Texture{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[RenderbufferId(uint32(0))] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	color_id := RenderbufferId(uint32(4294967295))   // RenderbufferId
	depth_id := RenderbufferId(uint32(4294967294))   // RenderbufferId
	stencil_id := RenderbufferId(uint32(4294967293)) // RenderbufferId
	ctx.Instances.Renderbuffers[color_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[depth_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[stencil_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	backbuffer := func() *Framebuffer {
		s := &Framebuffer{}
		s.Init()
		return s
	}() // Framebufferʳ
	backbuffer.Attachments[FramebufferAttachment_GL_COLOR_ATTACHMENT0] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(color_id)
		s.Type = FramebufferAttachmentType_GL_RENDERBUFFER
		s.CubeMapFace = CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	backbuffer.Attachments[FramebufferAttachment_GL_DEPTH_ATTACHMENT] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(depth_id)
		s.Type = FramebufferAttachmentType_GL_RENDERBUFFER
		s.CubeMapFace = CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	backbuffer.Attachments[FramebufferAttachment_GL_STENCIL_ATTACHMENT] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(stencil_id)
		s.Type = FramebufferAttachmentType_GL_RENDERBUFFER
		s.CubeMapFace = CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	ctx.Instances.Framebuffers[FramebufferId(uint32(0))] = backbuffer
	ctx.BoundFramebuffers[FramebufferTarget_GL_DRAW_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.BoundFramebuffers[FramebufferTarget_GL_READ_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.Rasterizing.StencilMask[FaceMode_GL_FRONT] = uint32(4294967295)
	ctx.Rasterizing.StencilMask[FaceMode_GL_BACK] = uint32(4294967295)
	ctx.PixelStorage[PixelStoreParameter_GL_PACK_ALIGNMENT] = int32(4)
	ctx.PixelStorage[PixelStoreParameter_GL_UNPACK_ALIGNMENT] = int32(4)
	for i := int32(int32(0)); i < int32(64); i++ {
		ctx.VertexAttributeArrays[AttributeLocation(i)] = func() *VertexAttributeArray {
			s := &VertexAttributeArray{}
			s.Init()
			return s
		}()
	}
	CreateContext_0_result := ctx // Contextʳ
	ϟc.EGLContexts[context] = CreateContext_0_result
	ϟa.Result = context
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_0_result
	return nil
}
func (ϟa *EglMakeCurrent) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	SetContext_1_context := ϟc.EGLContexts.Get(ϟa.Context) // Contextʳ
	ϟc.Contexts[ϟc.CurrentThread] = SetContext_1_context
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = SetContext_1_context
	return nil
}
func (ϟa *EglSwapBuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	return nil
}
func (ϟa *EglQuerySurface) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
	ϟa.Result = ϟa.Result
	return nil
}
func (ϟa *GlXCreateContext) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	context := GLXContext(ϟa.Result) // GLXContext
	identifier := ϟc.NextContextID   // ContextID
	ϟc.NextContextID = (ϟc.NextContextID) + (ContextID(uint32(1)))
	ctx := func() *Context {
		s := &Context{}
		s.Init()
		return s
	}() // Contextʳ
	ctx.Identifier = identifier
	ctx.Instances.Buffers[BufferId(uint32(0))] = func() *Buffer {
		s := &Buffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Textures[TextureId(uint32(0))] = func() *Texture {
		s := &Texture{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[RenderbufferId(uint32(0))] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	color_id := RenderbufferId(uint32(4294967295))   // RenderbufferId
	depth_id := RenderbufferId(uint32(4294967294))   // RenderbufferId
	stencil_id := RenderbufferId(uint32(4294967293)) // RenderbufferId
	ctx.Instances.Renderbuffers[color_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[depth_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[stencil_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	backbuffer := func() *Framebuffer {
		s := &Framebuffer{}
		s.Init()
		return s
	}() // Framebufferʳ
	backbuffer.Attachments[FramebufferAttachment_GL_COLOR_ATTACHMENT0] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(color_id)
		s.Type = FramebufferAttachmentType_GL_RENDERBUFFER
		s.CubeMapFace = CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	backbuffer.Attachments[FramebufferAttachment_GL_DEPTH_ATTACHMENT] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(depth_id)
		s.Type = FramebufferAttachmentType_GL_RENDERBUFFER
		s.CubeMapFace = CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	backbuffer.Attachments[FramebufferAttachment_GL_STENCIL_ATTACHMENT] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(stencil_id)
		s.Type = FramebufferAttachmentType_GL_RENDERBUFFER
		s.CubeMapFace = CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	ctx.Instances.Framebuffers[FramebufferId(uint32(0))] = backbuffer
	ctx.BoundFramebuffers[FramebufferTarget_GL_DRAW_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.BoundFramebuffers[FramebufferTarget_GL_READ_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.Rasterizing.StencilMask[FaceMode_GL_FRONT] = uint32(4294967295)
	ctx.Rasterizing.StencilMask[FaceMode_GL_BACK] = uint32(4294967295)
	ctx.PixelStorage[PixelStoreParameter_GL_PACK_ALIGNMENT] = int32(4)
	ctx.PixelStorage[PixelStoreParameter_GL_UNPACK_ALIGNMENT] = int32(4)
	for i := int32(int32(0)); i < int32(64); i++ {
		ctx.VertexAttributeArrays[AttributeLocation(i)] = func() *VertexAttributeArray {
			s := &VertexAttributeArray{}
			s.Init()
			return s
		}()
	}
	CreateContext_2_result := ctx // Contextʳ
	ϟc.GLXContexts[context] = CreateContext_2_result
	ϟa.Result = context
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_2_result
	return nil
}
func (ϟa *GlXCreateNewContext) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	context := GLXContext(ϟa.Result) // GLXContext
	identifier := ϟc.NextContextID   // ContextID
	ϟc.NextContextID = (ϟc.NextContextID) + (ContextID(uint32(1)))
	ctx := func() *Context {
		s := &Context{}
		s.Init()
		return s
	}() // Contextʳ
	ctx.Identifier = identifier
	ctx.Instances.Buffers[BufferId(uint32(0))] = func() *Buffer {
		s := &Buffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Textures[TextureId(uint32(0))] = func() *Texture {
		s := &Texture{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[RenderbufferId(uint32(0))] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	color_id := RenderbufferId(uint32(4294967295))   // RenderbufferId
	depth_id := RenderbufferId(uint32(4294967294))   // RenderbufferId
	stencil_id := RenderbufferId(uint32(4294967293)) // RenderbufferId
	ctx.Instances.Renderbuffers[color_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[depth_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[stencil_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	backbuffer := func() *Framebuffer {
		s := &Framebuffer{}
		s.Init()
		return s
	}() // Framebufferʳ
	backbuffer.Attachments[FramebufferAttachment_GL_COLOR_ATTACHMENT0] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(color_id)
		s.Type = FramebufferAttachmentType_GL_RENDERBUFFER
		s.CubeMapFace = CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	backbuffer.Attachments[FramebufferAttachment_GL_DEPTH_ATTACHMENT] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(depth_id)
		s.Type = FramebufferAttachmentType_GL_RENDERBUFFER
		s.CubeMapFace = CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	backbuffer.Attachments[FramebufferAttachment_GL_STENCIL_ATTACHMENT] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(stencil_id)
		s.Type = FramebufferAttachmentType_GL_RENDERBUFFER
		s.CubeMapFace = CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	ctx.Instances.Framebuffers[FramebufferId(uint32(0))] = backbuffer
	ctx.BoundFramebuffers[FramebufferTarget_GL_DRAW_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.BoundFramebuffers[FramebufferTarget_GL_READ_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.Rasterizing.StencilMask[FaceMode_GL_FRONT] = uint32(4294967295)
	ctx.Rasterizing.StencilMask[FaceMode_GL_BACK] = uint32(4294967295)
	ctx.PixelStorage[PixelStoreParameter_GL_PACK_ALIGNMENT] = int32(4)
	ctx.PixelStorage[PixelStoreParameter_GL_UNPACK_ALIGNMENT] = int32(4)
	for i := int32(int32(0)); i < int32(64); i++ {
		ctx.VertexAttributeArrays[AttributeLocation(i)] = func() *VertexAttributeArray {
			s := &VertexAttributeArray{}
			s.Init()
			return s
		}()
	}
	CreateContext_3_result := ctx // Contextʳ
	ϟc.GLXContexts[context] = CreateContext_3_result
	ϟa.Result = context
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_3_result
	return nil
}
func (ϟa *GlXMakeContextCurrent) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	SetContext_4_context := ϟc.GLXContexts.Get(ϟa.Ctx) // Contextʳ
	ϟc.Contexts[ϟc.CurrentThread] = SetContext_4_context
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = SetContext_4_context
	return nil
}
func (ϟa *GlXSwapBuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *WglCreateContext) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	context := HGLRC(ϟa.Result)    // HGLRC
	identifier := ϟc.NextContextID // ContextID
	ϟc.NextContextID = (ϟc.NextContextID) + (ContextID(uint32(1)))
	ctx := func() *Context {
		s := &Context{}
		s.Init()
		return s
	}() // Contextʳ
	ctx.Identifier = identifier
	ctx.Instances.Buffers[BufferId(uint32(0))] = func() *Buffer {
		s := &Buffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Textures[TextureId(uint32(0))] = func() *Texture {
		s := &Texture{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[RenderbufferId(uint32(0))] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	color_id := RenderbufferId(uint32(4294967295))   // RenderbufferId
	depth_id := RenderbufferId(uint32(4294967294))   // RenderbufferId
	stencil_id := RenderbufferId(uint32(4294967293)) // RenderbufferId
	ctx.Instances.Renderbuffers[color_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[depth_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[stencil_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	backbuffer := func() *Framebuffer {
		s := &Framebuffer{}
		s.Init()
		return s
	}() // Framebufferʳ
	backbuffer.Attachments[FramebufferAttachment_GL_COLOR_ATTACHMENT0] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(color_id)
		s.Type = FramebufferAttachmentType_GL_RENDERBUFFER
		s.CubeMapFace = CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	backbuffer.Attachments[FramebufferAttachment_GL_DEPTH_ATTACHMENT] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(depth_id)
		s.Type = FramebufferAttachmentType_GL_RENDERBUFFER
		s.CubeMapFace = CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	backbuffer.Attachments[FramebufferAttachment_GL_STENCIL_ATTACHMENT] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(stencil_id)
		s.Type = FramebufferAttachmentType_GL_RENDERBUFFER
		s.CubeMapFace = CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	ctx.Instances.Framebuffers[FramebufferId(uint32(0))] = backbuffer
	ctx.BoundFramebuffers[FramebufferTarget_GL_DRAW_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.BoundFramebuffers[FramebufferTarget_GL_READ_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.Rasterizing.StencilMask[FaceMode_GL_FRONT] = uint32(4294967295)
	ctx.Rasterizing.StencilMask[FaceMode_GL_BACK] = uint32(4294967295)
	ctx.PixelStorage[PixelStoreParameter_GL_PACK_ALIGNMENT] = int32(4)
	ctx.PixelStorage[PixelStoreParameter_GL_UNPACK_ALIGNMENT] = int32(4)
	for i := int32(int32(0)); i < int32(64); i++ {
		ctx.VertexAttributeArrays[AttributeLocation(i)] = func() *VertexAttributeArray {
			s := &VertexAttributeArray{}
			s.Init()
			return s
		}()
	}
	CreateContext_5_result := ctx // Contextʳ
	ϟc.WGLContexts[context] = CreateContext_5_result
	ϟa.Result = context
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_5_result
	return nil
}
func (ϟa *WglCreateContextAttribsARB) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	context := HGLRC(ϟa.Result)    // HGLRC
	identifier := ϟc.NextContextID // ContextID
	ϟc.NextContextID = (ϟc.NextContextID) + (ContextID(uint32(1)))
	ctx := func() *Context {
		s := &Context{}
		s.Init()
		return s
	}() // Contextʳ
	ctx.Identifier = identifier
	ctx.Instances.Buffers[BufferId(uint32(0))] = func() *Buffer {
		s := &Buffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Textures[TextureId(uint32(0))] = func() *Texture {
		s := &Texture{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[RenderbufferId(uint32(0))] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	color_id := RenderbufferId(uint32(4294967295))   // RenderbufferId
	depth_id := RenderbufferId(uint32(4294967294))   // RenderbufferId
	stencil_id := RenderbufferId(uint32(4294967293)) // RenderbufferId
	ctx.Instances.Renderbuffers[color_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[depth_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[stencil_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	backbuffer := func() *Framebuffer {
		s := &Framebuffer{}
		s.Init()
		return s
	}() // Framebufferʳ
	backbuffer.Attachments[FramebufferAttachment_GL_COLOR_ATTACHMENT0] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(color_id)
		s.Type = FramebufferAttachmentType_GL_RENDERBUFFER
		s.CubeMapFace = CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	backbuffer.Attachments[FramebufferAttachment_GL_DEPTH_ATTACHMENT] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(depth_id)
		s.Type = FramebufferAttachmentType_GL_RENDERBUFFER
		s.CubeMapFace = CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	backbuffer.Attachments[FramebufferAttachment_GL_STENCIL_ATTACHMENT] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(stencil_id)
		s.Type = FramebufferAttachmentType_GL_RENDERBUFFER
		s.CubeMapFace = CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	ctx.Instances.Framebuffers[FramebufferId(uint32(0))] = backbuffer
	ctx.BoundFramebuffers[FramebufferTarget_GL_DRAW_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.BoundFramebuffers[FramebufferTarget_GL_READ_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.Rasterizing.StencilMask[FaceMode_GL_FRONT] = uint32(4294967295)
	ctx.Rasterizing.StencilMask[FaceMode_GL_BACK] = uint32(4294967295)
	ctx.PixelStorage[PixelStoreParameter_GL_PACK_ALIGNMENT] = int32(4)
	ctx.PixelStorage[PixelStoreParameter_GL_UNPACK_ALIGNMENT] = int32(4)
	for i := int32(int32(0)); i < int32(64); i++ {
		ctx.VertexAttributeArrays[AttributeLocation(i)] = func() *VertexAttributeArray {
			s := &VertexAttributeArray{}
			s.Init()
			return s
		}()
	}
	CreateContext_6_result := ctx // Contextʳ
	ϟc.WGLContexts[context] = CreateContext_6_result
	ϟa.Result = context
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_6_result
	return nil
}
func (ϟa *WglMakeCurrent) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	SetContext_7_context := ϟc.WGLContexts.Get(ϟa.Hglrc) // Contextʳ
	ϟc.Contexts[ϟc.CurrentThread] = SetContext_7_context
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = SetContext_7_context
	return nil
}
func (ϟa *WglSwapBuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *CGLCreateContext) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	context := CGLContextObj(ϟa.Ctx.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl)) // CGLContextObj
	identifier := ϟc.NextContextID                                                                         // ContextID
	ϟc.NextContextID = (ϟc.NextContextID) + (ContextID(uint32(1)))
	ctx := func() *Context {
		s := &Context{}
		s.Init()
		return s
	}() // Contextʳ
	ctx.Identifier = identifier
	ctx.Instances.Buffers[BufferId(uint32(0))] = func() *Buffer {
		s := &Buffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Textures[TextureId(uint32(0))] = func() *Texture {
		s := &Texture{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[RenderbufferId(uint32(0))] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	color_id := RenderbufferId(uint32(4294967295))   // RenderbufferId
	depth_id := RenderbufferId(uint32(4294967294))   // RenderbufferId
	stencil_id := RenderbufferId(uint32(4294967293)) // RenderbufferId
	ctx.Instances.Renderbuffers[color_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[depth_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[stencil_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	backbuffer := func() *Framebuffer {
		s := &Framebuffer{}
		s.Init()
		return s
	}() // Framebufferʳ
	backbuffer.Attachments[FramebufferAttachment_GL_COLOR_ATTACHMENT0] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(color_id)
		s.Type = FramebufferAttachmentType_GL_RENDERBUFFER
		s.CubeMapFace = CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	backbuffer.Attachments[FramebufferAttachment_GL_DEPTH_ATTACHMENT] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(depth_id)
		s.Type = FramebufferAttachmentType_GL_RENDERBUFFER
		s.CubeMapFace = CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	backbuffer.Attachments[FramebufferAttachment_GL_STENCIL_ATTACHMENT] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(stencil_id)
		s.Type = FramebufferAttachmentType_GL_RENDERBUFFER
		s.CubeMapFace = CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	ctx.Instances.Framebuffers[FramebufferId(uint32(0))] = backbuffer
	ctx.BoundFramebuffers[FramebufferTarget_GL_DRAW_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.BoundFramebuffers[FramebufferTarget_GL_READ_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.Rasterizing.StencilMask[FaceMode_GL_FRONT] = uint32(4294967295)
	ctx.Rasterizing.StencilMask[FaceMode_GL_BACK] = uint32(4294967295)
	ctx.PixelStorage[PixelStoreParameter_GL_PACK_ALIGNMENT] = int32(4)
	ctx.PixelStorage[PixelStoreParameter_GL_UNPACK_ALIGNMENT] = int32(4)
	for i := int32(int32(0)); i < int32(64); i++ {
		ctx.VertexAttributeArrays[AttributeLocation(i)] = func() *VertexAttributeArray {
			s := &VertexAttributeArray{}
			s.Init()
			return s
		}()
	}
	CreateContext_8_result := ctx // Contextʳ
	ϟc.CGLContexts[context] = CreateContext_8_result
	ϟa.Ctx.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(context, ϟs)
	ϟa.Result = ϟa.Result
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_8_result
	return nil
}
func (ϟa *CGLSetCurrentContext) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	SetContext_9_context := ϟc.CGLContexts.Get(ϟa.Ctx) // Contextʳ
	ϟc.Contexts[ϟc.CurrentThread] = SetContext_9_context
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = SetContext_9_context
	return nil
}
func (ϟa *GlEnableClientState) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_10_result := context              // Contextʳ
	ctx := GetContext_10_result                  // Contextʳ
	ctx.Capabilities[Capability(ϟa.Type)] = true
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_10_result, ctx
	return nil
}
func (ϟa *GlDisableClientState) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_11_result := context              // Contextʳ
	ctx := GetContext_11_result                  // Contextʳ
	ctx.Capabilities[Capability(ϟa.Type)] = false
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_11_result, ctx
	return nil
}
func (ϟa *GlGetProgramBinaryOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	l := int32(ϟa.BytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl)) // s32
	ϟa.BytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(l, ϟs)
	ϟa.BinaryFormat.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.BinaryFormat.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
	_ = l
	return nil
}
func (ϟa *GlProgramBinaryOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlStartTilingQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlEndTilingQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlDiscardFramebufferEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlInsertEventMarkerEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	if (ϟa.Length) > (int32(0)) {
		_ = string(ϟa.Marker.StringSlice(ϟs, ϟd, ϟl, false).Read(ϟs, ϟd, ϟl))
	} else {
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlPushGroupMarkerEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	if (ϟa.Length) > (int32(0)) {
		_ = string(ϟa.Marker.StringSlice(ϟs, ϟd, ϟl, false).Read(ϟs, ϟd, ϟl))
	} else {
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlPopGroupMarkerEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlTexStorage1DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlTexStorage2DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlTexStorage3DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlTextureStorage1DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlTextureStorage2DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlTextureStorage3DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlGenVertexArraysOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	a := ϟa.Arrays.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // VertexArrayIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                 // Contextʳ
	GetContext_12_result := context                              // Contextʳ
	ctx := GetContext_12_result                                  // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		id := VertexArrayId(ϟa.Arrays.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).Read(ϟs, ϟd, ϟl)) // VertexArrayId
		ctx.Instances.VertexArrays[id] = func() *VertexArray {
			s := &VertexArray{}
			s.Init()
			return s
		}()
		a.Index(uint64(i), ϟs).Write(id, ϟs)
		_ = id
	}
	_, _, _, _ = a, context, GetContext_12_result, ctx
	return nil
}
func (ϟa *GlBindVertexArrayOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_13_result := context              // Contextʳ
	ctx := GetContext_13_result                  // Contextʳ
	if !(ctx.Instances.VertexArrays.Contains(ϟa.Array)) {
		ctx.Instances.VertexArrays[ϟa.Array] = func() *VertexArray {
			s := &VertexArray{}
			s.Init()
			return s
		}()
	}
	ctx.BoundVertexArray = ϟa.Array
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_13_result, ctx
	return nil
}
func (ϟa *GlDeleteVertexArraysOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                 // Contextʳ
	GetContext_14_result := context                              // Contextʳ
	ctx := GetContext_14_result                                  // Contextʳ
	a := ϟa.Arrays.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // VertexArrayIdˢ
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		ctx.Instances.VertexArrays[a.Index(uint64(i), ϟs).Read(ϟs, ϟd, ϟl)] = (*VertexArray)(nil)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = context, GetContext_14_result, ctx, a
	return nil
}
func (ϟa *GlIsVertexArrayOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_15_result := context              // Contextʳ
	ctx := GetContext_15_result                  // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ctx.Instances.VertexArrays.Contains(ϟa.Array)
	_, _, _ = context, GetContext_15_result, ctx
	return nil
}
func (ϟa *GlEGLImageTargetTexture2DOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlEGLImageTargetRenderbufferStorageOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlGetGraphicsResetStatusEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	return nil
}
func (ϟa *GlBindAttribLocation) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_16_result := context              // Contextʳ
	ctx := GetContext_16_result                  // Contextʳ
	p := ctx.Instances.Programs.Get(ϟa.Program)  // Programʳ
	p.AttributeBindings[ϟa.Name] = ϟa.Location
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = context, GetContext_16_result, ctx, p
	return nil
}
func (ϟa *GlBlendFunc) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_17_result := context              // Contextʳ
	ctx := GetContext_17_result                  // Contextʳ
	ctx.Blending.SrcRgbBlendFactor = ϟa.SrcFactor
	ctx.Blending.SrcAlphaBlendFactor = ϟa.SrcFactor
	ctx.Blending.DstRgbBlendFactor = ϟa.DstFactor
	ctx.Blending.DstAlphaBlendFactor = ϟa.DstFactor
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_17_result, ctx
	return nil
}
func (ϟa *GlBlendFuncSeparate) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_18_result := context              // Contextʳ
	ctx := GetContext_18_result                  // Contextʳ
	ctx.Blending.SrcRgbBlendFactor = ϟa.SrcFactorRgb
	ctx.Blending.DstRgbBlendFactor = ϟa.DstFactorRgb
	ctx.Blending.SrcAlphaBlendFactor = ϟa.SrcFactorAlpha
	ctx.Blending.DstAlphaBlendFactor = ϟa.DstFactorAlpha
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_18_result, ctx
	return nil
}
func (ϟa *GlBlendEquation) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_19_result := context              // Contextʳ
	ctx := GetContext_19_result                  // Contextʳ
	ctx.Blending.BlendEquationRgb = ϟa.Equation
	ctx.Blending.BlendEquationAlpha = ϟa.Equation
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_19_result, ctx
	return nil
}
func (ϟa *GlBlendEquationSeparate) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_20_result := context              // Contextʳ
	ctx := GetContext_20_result                  // Contextʳ
	ctx.Blending.BlendEquationRgb = ϟa.Rgb
	ctx.Blending.BlendEquationAlpha = ϟa.Alpha
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_20_result, ctx
	return nil
}
func (ϟa *GlBlendColor) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_21_result := context              // Contextʳ
	ctx := GetContext_21_result                  // Contextʳ
	ctx.Blending.BlendColor = func() Color {
		s := Color{}
		s.Init()
		s.Red = ϟa.Red
		s.Green = ϟa.Green
		s.Blue = ϟa.Blue
		s.Alpha = ϟa.Alpha
		return s
	}()
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_21_result, ctx
	return nil
}
func (ϟa *GlEnableVertexAttribArray) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_22_result := context              // Contextʳ
	ctx := GetContext_22_result                  // Contextʳ
	ctx.VertexAttributeArrays.Get(ϟa.Location).Enabled = true
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_22_result, ctx
	return nil
}
func (ϟa *GlDisableVertexAttribArray) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_23_result := context              // Contextʳ
	ctx := GetContext_23_result                  // Contextʳ
	ctx.VertexAttributeArrays.Get(ϟa.Location).Enabled = false
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_23_result, ctx
	return nil
}
func (ϟa *GlVertexAttribPointer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)    // Contextʳ
	GetContext_24_result := context                 // Contextʳ
	ctx := GetContext_24_result                     // Contextʳ
	a := ctx.VertexAttributeArrays.Get(ϟa.Location) // VertexAttributeArrayʳ
	a.Size = uint32(ϟa.Size)
	a.Type = ϟa.Type
	a.Normalized = ϟa.Normalized
	a.Stride = ϟa.Stride
	a.Pointer = ϟa.Data
	a.Buffer = ctx.BoundBuffers.Get(BufferTarget_GL_ARRAY_BUFFER)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = context, GetContext_24_result, ctx, a
	return nil
}
func (ϟa *GlGetActiveAttrib) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	l := int32(ϟa.BufferBytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl)) // s32
	ϟa.BufferBytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(l, ϟs)
	ϟa.VectorCount.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(int32(ϟa.VectorCount.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl)), ϟs)
	ϟa.Type.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ShaderAttribType(ϟa.Type.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl)), ϟs)
	_ = l
	return nil
}
func (ϟa *GlGetActiveUniform) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	l := int32(ϟa.BufferBytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl)) // s32
	ϟa.BufferBytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(l, ϟs)
	ϟa.VectorCount.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(int32(ϟa.VectorCount.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl)), ϟs)
	ϟa.Type.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ShaderUniformType(ϟa.Type.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl)), ϟs)
	_ = l
	return nil
}
func (ϟa *GlGetError) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	return nil
}
func (ϟa *GlGetProgramiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
	return nil
}
func (ϟa *GlGetShaderiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_25_result := context              // Contextʳ
	ctx := GetContext_25_result                  // Contextʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // Shaderʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(func() (result int32) {
		switch ϟa.Parameter {
		case ShaderParameter_GL_SHADER_TYPE:
			return int32(s.Type)
		case ShaderParameter_GL_DELETE_STATUS:
			return func() (result int32) {
				switch s.Deletable {
				case true:
					return int32(1)
				case false:
					return int32(0)
				default:
					// TODO: better unmatched handling
					panic(fmt.Errorf("Unmatched switch(%v) in atom %T", s.Deletable, ϟa))
					return result
				}
			}()
		case ShaderParameter_GL_COMPILE_STATUS:
			return func() (result int32) {
				switch s.Compiled {
				case true:
					return int32(1)
				case false:
					return int32(0)
				default:
					// TODO: better unmatched handling
					panic(fmt.Errorf("Unmatched switch(%v) in atom %T", s.Compiled, ϟa))
					return result
				}
			}()
		case ShaderParameter_GL_INFO_LOG_LENGTH:
			return int32(s.InfoLog.Count)
		case ShaderParameter_GL_SHADER_SOURCE_LENGTH:
			return int32(len(s.Source))
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.Parameter, ϟa))
			return result
		}
	}(), ϟs)
	_, _, _, _ = context, GetContext_25_result, ctx, s
	return nil
}
func (ϟa *GlGetUniformLocation) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	return nil
}
func (ϟa *GlGetAttribLocation) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	return nil
}
func (ϟa *GlPixelStorei) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_26_result := context              // Contextʳ
	ctx := GetContext_26_result                  // Contextʳ
	ctx.PixelStorage[ϟa.Parameter] = ϟa.Value
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_26_result, ctx
	return nil
}
func (ϟa *GlTexParameteri) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                     // Contextʳ
	GetContext_27_result := context                                  // Contextʳ
	ctx := GetContext_27_result                                      // Contextʳ
	id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(ϟa.Target) // TextureId
	t := ctx.Instances.Textures.Get(id)                              // Textureʳ
	switch ϟa.Parameter {
	case TextureParameter_GL_TEXTURE_MAG_FILTER:
		t.MagFilter = TextureFilterMode(ϟa.Value)
	case TextureParameter_GL_TEXTURE_MIN_FILTER:
		t.MinFilter = TextureFilterMode(ϟa.Value)
	case TextureParameter_GL_TEXTURE_WRAP_S:
		t.WrapS = TextureWrapMode(ϟa.Value)
	case TextureParameter_GL_TEXTURE_WRAP_T:
		t.WrapT = TextureWrapMode(ϟa.Value)
	case TextureParameter_GL_TEXTURE_MAX_ANISOTROPY_EXT:
		t.MaxAnisotropy = float32(ϟa.Value)
	case TextureParameter_GL_TEXTURE_SWIZZLE_R:
		t.SwizzleR = TexelComponent(ϟa.Value)
	case TextureParameter_GL_TEXTURE_SWIZZLE_G:
		t.SwizzleG = TexelComponent(ϟa.Value)
	case TextureParameter_GL_TEXTURE_SWIZZLE_B:
		t.SwizzleB = TexelComponent(ϟa.Value)
	case TextureParameter_GL_TEXTURE_SWIZZLE_A:
		t.SwizzleA = TexelComponent(ϟa.Value)
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = context, GetContext_27_result, ctx, id, t
	return nil
}
func (ϟa *GlTexParameterf) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                     // Contextʳ
	GetContext_28_result := context                                  // Contextʳ
	ctx := GetContext_28_result                                      // Contextʳ
	id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(ϟa.Target) // TextureId
	t := ctx.Instances.Textures.Get(id)                              // Textureʳ
	switch ϟa.Parameter {
	case TextureParameter_GL_TEXTURE_MAG_FILTER:
		t.MagFilter = TextureFilterMode(ϟa.Value)
	case TextureParameter_GL_TEXTURE_MIN_FILTER:
		t.MinFilter = TextureFilterMode(ϟa.Value)
	case TextureParameter_GL_TEXTURE_WRAP_S:
		t.WrapS = TextureWrapMode(ϟa.Value)
	case TextureParameter_GL_TEXTURE_WRAP_T:
		t.WrapT = TextureWrapMode(ϟa.Value)
	case TextureParameter_GL_TEXTURE_MAX_ANISOTROPY_EXT:
		t.MaxAnisotropy = ϟa.Value
	case TextureParameter_GL_TEXTURE_SWIZZLE_R:
		t.SwizzleR = TexelComponent(ϟa.Value)
	case TextureParameter_GL_TEXTURE_SWIZZLE_G:
		t.SwizzleG = TexelComponent(ϟa.Value)
	case TextureParameter_GL_TEXTURE_SWIZZLE_B:
		t.SwizzleB = TexelComponent(ϟa.Value)
	case TextureParameter_GL_TEXTURE_SWIZZLE_A:
		t.SwizzleA = TexelComponent(ϟa.Value)
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = context, GetContext_28_result, ctx, id, t
	return nil
}
func (ϟa *GlGetTexParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                     // Contextʳ
	GetContext_29_result := context                                  // Contextʳ
	ctx := GetContext_29_result                                      // Contextʳ
	id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(ϟa.Target) // TextureId
	t := ctx.Instances.Textures.Get(id)                              // Textureʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Values.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(func() (result int32) {
		switch ϟa.Parameter {
		case TextureParameter_GL_TEXTURE_MAG_FILTER:
			return int32(t.MagFilter)
		case TextureParameter_GL_TEXTURE_MIN_FILTER:
			return int32(t.MinFilter)
		case TextureParameter_GL_TEXTURE_WRAP_S:
			return int32(t.WrapS)
		case TextureParameter_GL_TEXTURE_WRAP_T:
			return int32(t.WrapT)
		case TextureParameter_GL_TEXTURE_MAX_ANISOTROPY_EXT:
			return int32(t.MaxAnisotropy)
		case TextureParameter_GL_TEXTURE_SWIZZLE_R:
			return int32(t.SwizzleR)
		case TextureParameter_GL_TEXTURE_SWIZZLE_G:
			return int32(t.SwizzleG)
		case TextureParameter_GL_TEXTURE_SWIZZLE_B:
			return int32(t.SwizzleB)
		case TextureParameter_GL_TEXTURE_SWIZZLE_A:
			return int32(t.SwizzleA)
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.Parameter, ϟa))
			return result
		}
	}(), ϟs)
	_, _, _, _, _ = context, GetContext_29_result, ctx, id, t
	return nil
}
func (ϟa *GlGetTexParameterfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                     // Contextʳ
	GetContext_30_result := context                                  // Contextʳ
	ctx := GetContext_30_result                                      // Contextʳ
	id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(ϟa.Target) // TextureId
	t := ctx.Instances.Textures.Get(id)                              // Textureʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Values.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(func() (result float32) {
		switch ϟa.Parameter {
		case TextureParameter_GL_TEXTURE_MAG_FILTER:
			return float32(t.MagFilter)
		case TextureParameter_GL_TEXTURE_MIN_FILTER:
			return float32(t.MinFilter)
		case TextureParameter_GL_TEXTURE_WRAP_S:
			return float32(t.WrapS)
		case TextureParameter_GL_TEXTURE_WRAP_T:
			return float32(t.WrapT)
		case TextureParameter_GL_TEXTURE_MAX_ANISOTROPY_EXT:
			return t.MaxAnisotropy
		case TextureParameter_GL_TEXTURE_SWIZZLE_R:
			return float32(t.SwizzleR)
		case TextureParameter_GL_TEXTURE_SWIZZLE_G:
			return float32(t.SwizzleG)
		case TextureParameter_GL_TEXTURE_SWIZZLE_B:
			return float32(t.SwizzleB)
		case TextureParameter_GL_TEXTURE_SWIZZLE_A:
			return float32(t.SwizzleA)
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.Parameter, ϟa))
			return result
		}
	}(), ϟs)
	_, _, _, _, _ = context, GetContext_30_result, ctx, id, t
	return nil
}
func (ϟa *GlUniform1i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)            // Contextʳ
	GetContext_31_result := context                         // Contextʳ
	ctx := GetContext_31_result                             // Contextʳ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = ShaderUniformType_GL_INT
	uniform.Value.S32 = ϟa.Value
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = context, GetContext_31_result, ctx, program, uniform
	return nil
}
func (ϟa *GlUniform2i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)            // Contextʳ
	GetContext_32_result := context                         // Contextʳ
	ctx := GetContext_32_result                             // Contextʳ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = ShaderUniformType_GL_INT_VEC2
	uniform.Value.Vec2i = func() Vec2i {
		s := Vec2i{}
		s.Init()
		s.X = ϟa.Value0
		s.Y = ϟa.Value1
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = context, GetContext_32_result, ctx, program, uniform
	return nil
}
func (ϟa *GlUniform3i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)            // Contextʳ
	GetContext_33_result := context                         // Contextʳ
	ctx := GetContext_33_result                             // Contextʳ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = ShaderUniformType_GL_INT_VEC3
	uniform.Value.Vec3i = func() Vec3i {
		s := Vec3i{}
		s.Init()
		s.X = ϟa.Value0
		s.Y = ϟa.Value1
		s.Z = ϟa.Value2
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = context, GetContext_33_result, ctx, program, uniform
	return nil
}
func (ϟa *GlUniform4i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)            // Contextʳ
	GetContext_34_result := context                         // Contextʳ
	ctx := GetContext_34_result                             // Contextʳ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = ShaderUniformType_GL_INT_VEC4
	uniform.Value.Vec4i = func() Vec4i {
		s := Vec4i{}
		s.Init()
		s.X = ϟa.Value0
		s.Y = ϟa.Value1
		s.Z = ϟa.Value2
		s.W = ϟa.Value3
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = context, GetContext_34_result, ctx, program, uniform
	return nil
}
func (ϟa *GlUniform1iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)            // Contextʳ
	GetContext_35_result := context                         // Contextʳ
	ctx := GetContext_35_result                             // Contextʳ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = ShaderUniformType_GL_INT
	uniform.Value.S32 = ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = context, GetContext_35_result, ctx, program, uniform
	return nil
}
func (ϟa *GlUniform2iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                             // Contextʳ
	GetContext_36_result := context                                          // Contextʳ
	ctx := GetContext_36_result                                              // Contextʳ
	v := ϟa.Value.Slice(uint64(int32(0)), uint64((ϟa.Count)*(int32(2))), ϟs) // S32ˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                  // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                             // Uniform
	uniform.Type = ShaderUniformType_GL_INT_VEC2
	uniform.Value.Vec2i = func() Vec2i {
		s := Vec2i{}
		s.Init()
		s.X = v.Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl)
		s.Y = v.Index(uint64(1), ϟs).Read(ϟs, ϟd, ϟl)
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = context, GetContext_36_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform3iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                             // Contextʳ
	GetContext_37_result := context                                          // Contextʳ
	ctx := GetContext_37_result                                              // Contextʳ
	v := ϟa.Value.Slice(uint64(int32(0)), uint64((ϟa.Count)*(int32(3))), ϟs) // S32ˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                  // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                             // Uniform
	uniform.Type = ShaderUniformType_GL_INT_VEC3
	uniform.Value.Vec3i = func() Vec3i {
		s := Vec3i{}
		s.Init()
		s.X = v.Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl)
		s.Y = v.Index(uint64(1), ϟs).Read(ϟs, ϟd, ϟl)
		s.Z = v.Index(uint64(2), ϟs).Read(ϟs, ϟd, ϟl)
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = context, GetContext_37_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform4iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                             // Contextʳ
	GetContext_38_result := context                                          // Contextʳ
	ctx := GetContext_38_result                                              // Contextʳ
	v := ϟa.Value.Slice(uint64(int32(0)), uint64((ϟa.Count)*(int32(4))), ϟs) // S32ˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                  // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                             // Uniform
	uniform.Type = ShaderUniformType_GL_INT_VEC4
	uniform.Value.Vec4i = func() Vec4i {
		s := Vec4i{}
		s.Init()
		s.X = v.Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl)
		s.Y = v.Index(uint64(1), ϟs).Read(ϟs, ϟd, ϟl)
		s.Z = v.Index(uint64(2), ϟs).Read(ϟs, ϟd, ϟl)
		s.W = v.Index(uint64(3), ϟs).Read(ϟs, ϟd, ϟl)
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = context, GetContext_38_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform1f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)            // Contextʳ
	GetContext_39_result := context                         // Contextʳ
	ctx := GetContext_39_result                             // Contextʳ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT
	uniform.Value.F32 = ϟa.Value
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = context, GetContext_39_result, ctx, program, uniform
	return nil
}
func (ϟa *GlUniform2f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)            // Contextʳ
	GetContext_40_result := context                         // Contextʳ
	ctx := GetContext_40_result                             // Contextʳ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT_VEC2
	uniform.Value.Vec2f = func() Vec2f {
		s := Vec2f{}
		s.Init()
		s.X = ϟa.Value0
		s.Y = ϟa.Value1
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = context, GetContext_40_result, ctx, program, uniform
	return nil
}
func (ϟa *GlUniform3f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)            // Contextʳ
	GetContext_41_result := context                         // Contextʳ
	ctx := GetContext_41_result                             // Contextʳ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT_VEC3
	uniform.Value.Vec3f = func() Vec3f {
		s := Vec3f{}
		s.Init()
		s.X = ϟa.Value0
		s.Y = ϟa.Value1
		s.Z = ϟa.Value2
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = context, GetContext_41_result, ctx, program, uniform
	return nil
}
func (ϟa *GlUniform4f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)            // Contextʳ
	GetContext_42_result := context                         // Contextʳ
	ctx := GetContext_42_result                             // Contextʳ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT_VEC4
	uniform.Value.Vec4f = func() Vec4f {
		s := Vec4f{}
		s.Init()
		s.X = ϟa.Value0
		s.Y = ϟa.Value1
		s.Z = ϟa.Value2
		s.W = ϟa.Value3
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = context, GetContext_42_result, ctx, program, uniform
	return nil
}
func (ϟa *GlUniform1fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                // Contextʳ
	GetContext_43_result := context                             // Contextʳ
	ctx := GetContext_43_result                                 // Contextʳ
	v := ϟa.Value.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // F32ˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)     // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT
	uniform.Value.F32 = v.Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = context, GetContext_43_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                             // Contextʳ
	GetContext_44_result := context                                          // Contextʳ
	ctx := GetContext_44_result                                              // Contextʳ
	v := ϟa.Value.Slice(uint64(int32(0)), uint64((ϟa.Count)*(int32(2))), ϟs) // F32ˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                  // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                             // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT_VEC2
	uniform.Value.Vec2f = func() Vec2f {
		s := Vec2f{}
		s.Init()
		s.X = v.Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl)
		s.Y = v.Index(uint64(1), ϟs).Read(ϟs, ϟd, ϟl)
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = context, GetContext_44_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                             // Contextʳ
	GetContext_45_result := context                                          // Contextʳ
	ctx := GetContext_45_result                                              // Contextʳ
	v := ϟa.Value.Slice(uint64(int32(0)), uint64((ϟa.Count)*(int32(3))), ϟs) // F32ˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                  // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                             // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT_VEC3
	uniform.Value.Vec3f = func() Vec3f {
		s := Vec3f{}
		s.Init()
		s.X = v.Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl)
		s.Y = v.Index(uint64(1), ϟs).Read(ϟs, ϟd, ϟl)
		s.Z = v.Index(uint64(2), ϟs).Read(ϟs, ϟd, ϟl)
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = context, GetContext_45_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                             // Contextʳ
	GetContext_46_result := context                                          // Contextʳ
	ctx := GetContext_46_result                                              // Contextʳ
	v := ϟa.Value.Slice(uint64(int32(0)), uint64((ϟa.Count)*(int32(4))), ϟs) // F32ˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                  // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                             // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT_VEC4
	uniform.Value.Vec4f = func() Vec4f {
		s := Vec4f{}
		s.Init()
		s.X = v.Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl)
		s.Y = v.Index(uint64(1), ϟs).Read(ϟs, ϟd, ϟl)
		s.Z = v.Index(uint64(2), ϟs).Read(ϟs, ϟd, ϟl)
		s.W = v.Index(uint64(3), ϟs).Read(ϟs, ϟd, ϟl)
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = context, GetContext_46_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniformMatrix2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                              // Contextʳ
	GetContext_47_result := context                                           // Contextʳ
	ctx := GetContext_47_result                                               // Contextʳ
	v := ϟa.Values.Slice(uint64(int32(0)), uint64((ϟa.Count)*(int32(4))), ϟs) // F32ˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                   // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                              // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT_MAT2
	uniform.Value.Mat2f = func() Mat2f {
		s := Mat2f{}
		s.Init()
		s.Col0 = func() Vec2f {
			s := Vec2f{}
			s.Init()
			s.X = v.Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl)
			s.Y = v.Index(uint64(1), ϟs).Read(ϟs, ϟd, ϟl)
			return s
		}()
		s.Col1 = func() Vec2f {
			s := Vec2f{}
			s.Init()
			s.X = v.Index(uint64(3), ϟs).Read(ϟs, ϟd, ϟl)
			s.Y = v.Index(uint64(4), ϟs).Read(ϟs, ϟd, ϟl)
			return s
		}()
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = context, GetContext_47_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniformMatrix3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                              // Contextʳ
	GetContext_48_result := context                                           // Contextʳ
	ctx := GetContext_48_result                                               // Contextʳ
	v := ϟa.Values.Slice(uint64(int32(0)), uint64((ϟa.Count)*(int32(9))), ϟs) // F32ˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                   // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                              // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT_MAT3
	uniform.Value.Mat3f = func() Mat3f {
		s := Mat3f{}
		s.Init()
		s.Col0 = func() Vec3f {
			s := Vec3f{}
			s.Init()
			s.X = v.Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl)
			s.Y = v.Index(uint64(1), ϟs).Read(ϟs, ϟd, ϟl)
			s.Z = v.Index(uint64(2), ϟs).Read(ϟs, ϟd, ϟl)
			return s
		}()
		s.Col1 = func() Vec3f {
			s := Vec3f{}
			s.Init()
			s.X = v.Index(uint64(3), ϟs).Read(ϟs, ϟd, ϟl)
			s.Y = v.Index(uint64(4), ϟs).Read(ϟs, ϟd, ϟl)
			s.Z = v.Index(uint64(5), ϟs).Read(ϟs, ϟd, ϟl)
			return s
		}()
		s.Col2 = func() Vec3f {
			s := Vec3f{}
			s.Init()
			s.X = v.Index(uint64(6), ϟs).Read(ϟs, ϟd, ϟl)
			s.Y = v.Index(uint64(7), ϟs).Read(ϟs, ϟd, ϟl)
			s.Z = v.Index(uint64(8), ϟs).Read(ϟs, ϟd, ϟl)
			return s
		}()
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = context, GetContext_48_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniformMatrix4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                               // Contextʳ
	GetContext_49_result := context                                            // Contextʳ
	ctx := GetContext_49_result                                                // Contextʳ
	v := ϟa.Values.Slice(uint64(int32(0)), uint64((ϟa.Count)*(int32(16))), ϟs) // F32ˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                    // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                               // Uniform
	uniform.Value.Mat4f = func() Mat4f {
		s := Mat4f{}
		s.Init()
		s.Col0 = func() Vec4f {
			s := Vec4f{}
			s.Init()
			s.X = v.Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl)
			s.Y = v.Index(uint64(1), ϟs).Read(ϟs, ϟd, ϟl)
			s.Z = v.Index(uint64(2), ϟs).Read(ϟs, ϟd, ϟl)
			s.W = v.Index(uint64(3), ϟs).Read(ϟs, ϟd, ϟl)
			return s
		}()
		s.Col1 = func() Vec4f {
			s := Vec4f{}
			s.Init()
			s.X = v.Index(uint64(4), ϟs).Read(ϟs, ϟd, ϟl)
			s.Y = v.Index(uint64(5), ϟs).Read(ϟs, ϟd, ϟl)
			s.Z = v.Index(uint64(6), ϟs).Read(ϟs, ϟd, ϟl)
			s.W = v.Index(uint64(7), ϟs).Read(ϟs, ϟd, ϟl)
			return s
		}()
		s.Col2 = func() Vec4f {
			s := Vec4f{}
			s.Init()
			s.X = v.Index(uint64(8), ϟs).Read(ϟs, ϟd, ϟl)
			s.Y = v.Index(uint64(9), ϟs).Read(ϟs, ϟd, ϟl)
			s.Z = v.Index(uint64(10), ϟs).Read(ϟs, ϟd, ϟl)
			s.W = v.Index(uint64(11), ϟs).Read(ϟs, ϟd, ϟl)
			return s
		}()
		s.Col3 = func() Vec4f {
			s := Vec4f{}
			s.Init()
			s.X = v.Index(uint64(12), ϟs).Read(ϟs, ϟd, ϟl)
			s.Y = v.Index(uint64(13), ϟs).Read(ϟs, ϟd, ϟl)
			s.Z = v.Index(uint64(14), ϟs).Read(ϟs, ϟd, ϟl)
			s.W = v.Index(uint64(15), ϟs).Read(ϟs, ϟd, ϟl)
			return s
		}()
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = context, GetContext_49_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlGetUniformfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlGetUniformiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlVertexAttrib1f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlVertexAttrib2f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlVertexAttrib3f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlVertexAttrib4f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlVertexAttrib1fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlVertexAttrib2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlVertexAttrib3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlVertexAttrib4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlGetShaderPrecisionFormat) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Precision.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Precision.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
	return nil
}
func (ϟa *GlDepthMask) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_50_result := context              // Contextʳ
	ctx := GetContext_50_result                  // Contextʳ
	ctx.Rasterizing.DepthMask = ϟa.Enabled
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_50_result, ctx
	return nil
}
func (ϟa *GlDepthFunc) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_51_result := context              // Contextʳ
	ctx := GetContext_51_result                  // Contextʳ
	ctx.Rasterizing.DepthTestFunction = ϟa.Function
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_51_result, ctx
	return nil
}
func (ϟa *GlDepthRangef) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_52_result := context              // Contextʳ
	ctx := GetContext_52_result                  // Contextʳ
	ctx.Rasterizing.DepthNear = ϟa.Near
	ctx.Rasterizing.DepthFar = ϟa.Far
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_52_result, ctx
	return nil
}
func (ϟa *GlColorMask) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_53_result := context              // Contextʳ
	ctx := GetContext_53_result                  // Contextʳ
	ctx.Rasterizing.ColorMaskRed = ϟa.Red
	ctx.Rasterizing.ColorMaskGreen = ϟa.Green
	ctx.Rasterizing.ColorMaskBlue = ϟa.Blue
	ctx.Rasterizing.ColorMaskAlpha = ϟa.Alpha
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_53_result, ctx
	return nil
}
func (ϟa *GlStencilMask) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_54_result := context              // Contextʳ
	ctx := GetContext_54_result                  // Contextʳ
	ctx.Rasterizing.StencilMask[FaceMode_GL_FRONT] = ϟa.Mask
	ctx.Rasterizing.StencilMask[FaceMode_GL_BACK] = ϟa.Mask
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_54_result, ctx
	return nil
}
func (ϟa *GlStencilMaskSeparate) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_55_result := context              // Contextʳ
	ctx := GetContext_55_result                  // Contextʳ
	switch ϟa.Face {
	case FaceMode_GL_FRONT:
		ctx.Rasterizing.StencilMask[FaceMode_GL_FRONT] = ϟa.Mask
	case FaceMode_GL_BACK:
		ctx.Rasterizing.StencilMask[FaceMode_GL_BACK] = ϟa.Mask
	case FaceMode_GL_FRONT_AND_BACK:
		ctx.Rasterizing.StencilMask[FaceMode_GL_FRONT] = ϟa.Mask
		ctx.Rasterizing.StencilMask[FaceMode_GL_BACK] = ϟa.Mask
	default:
		v := ϟa.Face
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_55_result, ctx
	return nil
}
func (ϟa *GlStencilFuncSeparate) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlStencilOpSeparate) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlFrontFace) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_56_result := context              // Contextʳ
	ctx := GetContext_56_result                  // Contextʳ
	ctx.Rasterizing.FrontFace = ϟa.Orientation
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_56_result, ctx
	return nil
}
func (ϟa *GlViewport) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_57_result := context              // Contextʳ
	ctx := GetContext_57_result                  // Contextʳ
	ctx.Rasterizing.Viewport = func() Rect {
		s := Rect{}
		s.Init()
		s.X = ϟa.X
		s.Y = ϟa.Y
		s.Width = ϟa.Width
		s.Height = ϟa.Height
		return s
	}()
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_57_result, ctx
	return nil
}
func (ϟa *GlScissor) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_58_result := context              // Contextʳ
	ctx := GetContext_58_result                  // Contextʳ
	ctx.Rasterizing.Scissor = func() Rect {
		s := Rect{}
		s.Init()
		s.X = ϟa.X
		s.Y = ϟa.Y
		s.Width = ϟa.Width
		s.Height = ϟa.Height
		return s
	}()
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_58_result, ctx
	return nil
}
func (ϟa *GlActiveTexture) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_59_result := context              // Contextʳ
	ctx := GetContext_59_result                  // Contextʳ
	ctx.ActiveTextureUnit = ϟa.Unit
	if !(ctx.TextureUnits.Contains(ϟa.Unit)) {
		ctx.TextureUnits[ϟa.Unit] = ctx.TextureUnits.Get(ϟa.Unit)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_59_result, ctx
	return nil
}
func (ϟa *GlGenTextures) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	t := ϟa.Textures.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // TextureIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                   // Contextʳ
	GetContext_60_result := context                                // Contextʳ
	ctx := GetContext_60_result                                    // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		id := TextureId(ϟa.Textures.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).Read(ϟs, ϟd, ϟl)) // TextureId
		ctx.Instances.Textures[id] = func() *Texture {
			s := &Texture{}
			s.Init()
			return s
		}()
		t.Index(uint64(i), ϟs).Write(id, ϟs)
		_ = id
	}
	_, _, _, _ = t, context, GetContext_60_result, ctx
	return nil
}
func (ϟa *GlDeleteTextures) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	t := ϟa.Textures.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // TextureIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                   // Contextʳ
	GetContext_61_result := context                                // Contextʳ
	ctx := GetContext_61_result                                    // Contextʳ
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		ctx.Instances.Textures[t.Index(uint64(i), ϟs).Read(ϟs, ϟd, ϟl)] = (*Texture)(nil)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = t, context, GetContext_61_result, ctx
	return nil
}
func (ϟa *GlIsTexture) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_62_result := context              // Contextʳ
	ctx := GetContext_62_result                  // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ctx.Instances.Textures.Contains(ϟa.Texture)
	_, _, _ = context, GetContext_62_result, ctx
	return nil
}
func (ϟa *GlBindTexture) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_63_result := context              // Contextʳ
	ctx := GetContext_63_result                  // Contextʳ
	if !(ctx.Instances.Textures.Contains(ϟa.Texture)) {
		ctx.Instances.Textures[ϟa.Texture] = func() *Texture {
			s := &Texture{}
			s.Init()
			return s
		}()
	}
	ctx.TextureUnits.Get(ctx.ActiveTextureUnit)[ϟa.Target] = ϟa.Texture
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_63_result, ctx
	return nil
}
func (ϟa *GlTexImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_64_result := context              // Contextʳ
	ctx := GetContext_64_result                  // Contextʳ
	switch ϟa.Target {
	case TextureImageTarget_GL_TEXTURE_2D:
		id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_2D) // TextureId
		t := ctx.Instances.Textures.Get(id)                                                // Textureʳ
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.Width
			s.Height = ϟa.Height
			s.Size = externs{ϟs, ϟd, ϟl}.imageSize(uint32(ϟa.Width), uint32(ϟa.Height), ϟa.Format, ϟa.Type)
			s.Format = ImageTexelFormat(ϟa.Format)
			return s
		}() // Image
		if (ϟa.Data) != (TexturePointer(Voidᵖ{})) {
			if (ctx.BoundBuffers.Get(BufferTarget_GL_PIXEL_UNPACK_BUFFER)) == (BufferId(uint32(0))) {
				l.Data = U8ᵖ(ϟa.Data).Slice(uint64(uint32(0)), uint64(l.Size), ϟs).Clone(ϟs)
			}
		} else {
			l.Data = MakeU8ˢ(uint64(l.Size), ϟs)
		}
		t.Texture2D[ϟa.Level] = l
		t.Kind = TextureKind_TEXTURE2D
		t.Format = ImageTexelFormat(ϟa.Format)
		_, _, _ = id, t, l
	case TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X, TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Z, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z:
		id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_CUBE_MAP) // TextureId
		t := ctx.Instances.Textures.Get(id)                                                      // Textureʳ
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.Width
			s.Height = ϟa.Height
			s.Size = externs{ϟs, ϟd, ϟl}.imageSize(uint32(ϟa.Width), uint32(ϟa.Height), ϟa.Format, ϟa.Type)
			s.Format = ImageTexelFormat(ϟa.Format)
			return s
		}() // Image
		if (ϟa.Data) != (TexturePointer(Voidᵖ{})) {
			if (ctx.BoundBuffers.Get(BufferTarget_GL_PIXEL_UNPACK_BUFFER)) == (BufferId(uint32(0))) {
				l.Data = U8ᵖ(ϟa.Data).Slice(uint64(uint32(0)), uint64(l.Size), ϟs).Clone(ϟs)
			}
		} else {
			l.Data = MakeU8ˢ(uint64(l.Size), ϟs)
		}
		cube := t.Cubemap.Get(ϟa.Level) // CubemapLevel
		cube.Faces[CubeMapImageTarget(ϟa.Target)] = l
		t.Cubemap[ϟa.Level] = cube
		t.Kind = TextureKind_CUBEMAP
		t.Format = ImageTexelFormat(ϟa.Format)
		_, _, _, _ = id, t, l, cube
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_64_result, ctx
	return nil
}
func (ϟa *GlTexSubImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_65_result := context              // Contextʳ
	ctx := GetContext_65_result                  // Contextʳ
	switch ϟa.Target {
	case TextureImageTarget_GL_TEXTURE_2D:
		id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_2D) // TextureId
		t := ctx.Instances.Textures.Get(id)                                                // Textureʳ
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.Width
			s.Height = ϟa.Height
			s.Size = externs{ϟs, ϟd, ϟl}.imageSize(uint32(ϟa.Width), uint32(ϟa.Height), ϟa.Format, ϟa.Type)
			s.Format = ImageTexelFormat(ϟa.Format)
			return s
		}() // Image
		if ((ctx.BoundBuffers.Get(BufferTarget_GL_PIXEL_UNPACK_BUFFER)) == (BufferId(uint32(0)))) && ((ϟa.Data) != (TexturePointer(Voidᵖ{}))) {
			l.Data = U8ᵖ(ϟa.Data).Slice(uint64(uint32(0)), uint64(l.Size), ϟs).Clone(ϟs)
		}
		t.Texture2D[ϟa.Level] = l
		t.Kind = TextureKind_TEXTURE2D
		t.Format = ImageTexelFormat(ϟa.Format)
		_, _, _ = id, t, l
	case TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X, TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Z, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z:
		id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_CUBE_MAP) // TextureId
		t := ctx.Instances.Textures.Get(id)                                                      // Textureʳ
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.Width
			s.Height = ϟa.Height
			s.Size = externs{ϟs, ϟd, ϟl}.imageSize(uint32(ϟa.Width), uint32(ϟa.Height), ϟa.Format, ϟa.Type)
			s.Format = ImageTexelFormat(ϟa.Format)
			return s
		}() // Image
		if ((ctx.BoundBuffers.Get(BufferTarget_GL_PIXEL_UNPACK_BUFFER)) == (BufferId(uint32(0)))) && ((ϟa.Data) != (TexturePointer(Voidᵖ{}))) {
			l.Data = U8ᵖ(ϟa.Data).Slice(uint64(uint32(0)), uint64(l.Size), ϟs).Clone(ϟs)
		}
		cube := t.Cubemap.Get(ϟa.Level) // CubemapLevel
		cube.Faces[CubeMapImageTarget(ϟa.Target)] = l
		t.Cubemap[ϟa.Level] = cube
		t.Kind = TextureKind_CUBEMAP
		t.Format = ImageTexelFormat(ϟa.Format)
		_, _, _, _ = id, t, l, cube
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_65_result, ctx
	return nil
}
func (ϟa *GlCopyTexImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlCopyTexSubImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlCompressedTexImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_66_result := context              // Contextʳ
	ctx := GetContext_66_result                  // Contextʳ
	switch ϟa.Target {
	case TextureImageTarget_GL_TEXTURE_2D:
		id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_2D) // TextureId
		t := ctx.Instances.Textures.Get(id)                                                // Textureʳ
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.Width
			s.Height = ϟa.Height
			s.Size = uint32(ϟa.ImageSize)
			s.Format = ImageTexelFormat(ϟa.Format)
			return s
		}() // Image
		if ((ctx.BoundBuffers.Get(BufferTarget_GL_PIXEL_UNPACK_BUFFER)) == (BufferId(uint32(0)))) && ((ϟa.Data) != (TexturePointer(Voidᵖ{}))) {
			l.Data = U8ᵖ(ϟa.Data).Slice(uint64(uint32(0)), uint64(l.Size), ϟs).Clone(ϟs)
		}
		t.Texture2D[ϟa.Level] = l
		t.Kind = TextureKind_TEXTURE2D
		t.Format = ImageTexelFormat(ϟa.Format)
		_, _, _ = id, t, l
	case TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X, TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Z, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z:
		id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_CUBE_MAP) // TextureId
		t := ctx.Instances.Textures.Get(id)                                                      // Textureʳ
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.Width
			s.Height = ϟa.Height
			s.Size = uint32(ϟa.ImageSize)
			s.Format = ImageTexelFormat(ϟa.Format)
			return s
		}() // Image
		if ((ctx.BoundBuffers.Get(BufferTarget_GL_PIXEL_UNPACK_BUFFER)) == (BufferId(uint32(0)))) && ((ϟa.Data) != (TexturePointer(Voidᵖ{}))) {
			l.Data = U8ᵖ(ϟa.Data).Slice(uint64(uint32(0)), uint64(l.Size), ϟs).Clone(ϟs)
		}
		cube := t.Cubemap.Get(ϟa.Level) // CubemapLevel
		cube.Faces[CubeMapImageTarget(ϟa.Target)] = l
		t.Cubemap[ϟa.Level] = cube
		t.Kind = TextureKind_CUBEMAP
		t.Format = ImageTexelFormat(ϟa.Format)
		_, _, _, _ = id, t, l, cube
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_66_result, ctx
	return nil
}
func (ϟa *GlCompressedTexSubImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlGenerateMipmap) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlReadPixels) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlGenFramebuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	f := ϟa.Framebuffers.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // FramebufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                       // Contextʳ
	GetContext_67_result := context                                    // Contextʳ
	ctx := GetContext_67_result                                        // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		id := FramebufferId(ϟa.Framebuffers.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).Read(ϟs, ϟd, ϟl)) // FramebufferId
		ctx.Instances.Framebuffers[id] = func() *Framebuffer {
			s := &Framebuffer{}
			s.Init()
			return s
		}()
		f.Index(uint64(i), ϟs).Write(id, ϟs)
		_ = id
	}
	_, _, _, _ = f, context, GetContext_67_result, ctx
	return nil
}
func (ϟa *GlBindFramebuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_68_result := context              // Contextʳ
	ctx := GetContext_68_result                  // Contextʳ
	if !(ctx.Instances.Framebuffers.Contains(ϟa.Framebuffer)) {
		ctx.Instances.Framebuffers[ϟa.Framebuffer] = func() *Framebuffer {
			s := &Framebuffer{}
			s.Init()
			return s
		}()
	}
	if (ϟa.Target) == (FramebufferTarget_GL_FRAMEBUFFER) {
		ctx.BoundFramebuffers[FramebufferTarget_GL_READ_FRAMEBUFFER] = ϟa.Framebuffer
		ctx.BoundFramebuffers[FramebufferTarget_GL_DRAW_FRAMEBUFFER] = ϟa.Framebuffer
	} else {
		ctx.BoundFramebuffers[ϟa.Target] = ϟa.Framebuffer
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_68_result, ctx
	return nil
}
func (ϟa *GlCheckFramebufferStatus) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	return nil
}
func (ϟa *GlDeleteFramebuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	f := ϟa.Framebuffers.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // FramebufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                       // Contextʳ
	GetContext_69_result := context                                    // Contextʳ
	ctx := GetContext_69_result                                        // Contextʳ
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		ctx.Instances.Framebuffers[f.Index(uint64(i), ϟs).Read(ϟs, ϟd, ϟl)] = (*Framebuffer)(nil)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = f, context, GetContext_69_result, ctx
	return nil
}
func (ϟa *GlIsFramebuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_70_result := context              // Contextʳ
	ctx := GetContext_70_result                  // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ctx.Instances.Framebuffers.Contains(ϟa.Framebuffer)
	_, _, _ = context, GetContext_70_result, ctx
	return nil
}
func (ϟa *GlGenRenderbuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	r := ϟa.Renderbuffers.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // RenderbufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                        // Contextʳ
	GetContext_71_result := context                                     // Contextʳ
	ctx := GetContext_71_result                                         // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		id := RenderbufferId(ϟa.Renderbuffers.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).Read(ϟs, ϟd, ϟl)) // RenderbufferId
		ctx.Instances.Renderbuffers[id] = func() *Renderbuffer {
			s := &Renderbuffer{}
			s.Init()
			return s
		}()
		r.Index(uint64(i), ϟs).Write(id, ϟs)
		_ = id
	}
	_, _, _, _ = r, context, GetContext_71_result, ctx
	return nil
}
func (ϟa *GlBindRenderbuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_72_result := context              // Contextʳ
	ctx := GetContext_72_result                  // Contextʳ
	if !(ctx.Instances.Renderbuffers.Contains(ϟa.Renderbuffer)) {
		ctx.Instances.Renderbuffers[ϟa.Renderbuffer] = func() *Renderbuffer {
			s := &Renderbuffer{}
			s.Init()
			return s
		}()
	}
	ctx.BoundRenderbuffers[ϟa.Target] = ϟa.Renderbuffer
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_72_result, ctx
	return nil
}
func (ϟa *GlRenderbufferStorage) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_73_result := context              // Contextʳ
	ctx := GetContext_73_result                  // Contextʳ
	id := ctx.BoundRenderbuffers.Get(ϟa.Target)  // RenderbufferId
	rb := ctx.Instances.Renderbuffers.Get(id)    // Renderbufferʳ
	rb.Format = ϟa.Format
	rb.Width = ϟa.Width
	rb.Height = ϟa.Height
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = context, GetContext_73_result, ctx, id, rb
	return nil
}
func (ϟa *GlDeleteRenderbuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	r := ϟa.Renderbuffers.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // RenderbufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                        // Contextʳ
	GetContext_74_result := context                                     // Contextʳ
	ctx := GetContext_74_result                                         // Contextʳ
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		ctx.Instances.Renderbuffers[r.Index(uint64(i), ϟs).Read(ϟs, ϟd, ϟl)] = (*Renderbuffer)(nil)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = r, context, GetContext_74_result, ctx
	return nil
}
func (ϟa *GlIsRenderbuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_75_result := context              // Contextʳ
	ctx := GetContext_75_result                  // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ctx.Instances.Renderbuffers.Contains(ϟa.Renderbuffer)
	_, _, _ = context, GetContext_75_result, ctx
	return nil
}
func (ϟa *GlGetRenderbufferParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_76_result := context              // Contextʳ
	ctx := GetContext_76_result                  // Contextʳ
	id := ctx.BoundRenderbuffers.Get(ϟa.Target)  // RenderbufferId
	rb := ctx.Instances.Renderbuffers.Get(id)    // Renderbufferʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Values.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(func() (result int32) {
		switch ϟa.Parameter {
		case RenderbufferParameter_GL_RENDERBUFFER_WIDTH:
			return rb.Width
		case RenderbufferParameter_GL_RENDERBUFFER_HEIGHT:
			return rb.Height
		case RenderbufferParameter_GL_RENDERBUFFER_INTERNAL_FORMAT:
			return int32(rb.Format)
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.Parameter, ϟa))
			return result
		}
	}(), ϟs)
	_, _, _, _, _ = context, GetContext_76_result, ctx, id, rb
	return nil
}
func (ϟa *GlGenBuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	b := ϟa.Buffers.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // BufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                  // Contextʳ
	GetContext_77_result := context                               // Contextʳ
	ctx := GetContext_77_result                                   // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		id := BufferId(ϟa.Buffers.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).Read(ϟs, ϟd, ϟl)) // BufferId
		ctx.Instances.Buffers[id] = func() *Buffer {
			s := &Buffer{}
			s.Init()
			return s
		}()
		b.Index(uint64(i), ϟs).Write(id, ϟs)
		_ = id
	}
	_, _, _, _ = b, context, GetContext_77_result, ctx
	return nil
}
func (ϟa *GlBindBuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_78_result := context              // Contextʳ
	ctx := GetContext_78_result                  // Contextʳ
	if !(ctx.Instances.Buffers.Contains(ϟa.Buffer)) {
		ctx.Instances.Buffers[ϟa.Buffer] = func() *Buffer {
			s := &Buffer{}
			s.Init()
			return s
		}()
	}
	ctx.BoundBuffers[ϟa.Target] = ϟa.Buffer
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_78_result, ctx
	return nil
}
func (ϟa *GlBufferData) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_79_result := context              // Contextʳ
	ctx := GetContext_79_result                  // Contextʳ
	id := ctx.BoundBuffers.Get(ϟa.Target)        // BufferId
	b := ctx.Instances.Buffers.Get(id)           // Bufferʳ
	b.Data = func() (result U8ˢ) {
		switch (ϟa.Data) != (BufferDataPointer(Voidᵖ{})) {
		case true:
			return U8ᵖ(ϟa.Data).Slice(uint64(int32(0)), uint64(ϟa.Size), ϟs).Clone(ϟs)
		case false:
			return MakeU8ˢ(uint64(ϟa.Size), ϟs)
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (ϟa.Data) != (BufferDataPointer(Voidᵖ{})), ϟa))
			return result
		}
	}()
	b.Size = ϟa.Size
	b.Usage = ϟa.Usage
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = context, GetContext_79_result, ctx, id, b
	return nil
}
func (ϟa *GlBufferSubData) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlDeleteBuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	b := ϟa.Buffers.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // BufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                  // Contextʳ
	GetContext_80_result := context                               // Contextʳ
	ctx := GetContext_80_result                                   // Contextʳ
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		ctx.Instances.Buffers[b.Index(uint64(i), ϟs).Read(ϟs, ϟd, ϟl)] = (*Buffer)(nil)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = b, context, GetContext_80_result, ctx
	return nil
}
func (ϟa *GlIsBuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_81_result := context              // Contextʳ
	ctx := GetContext_81_result                  // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ctx.Instances.Buffers.Contains(ϟa.Buffer)
	_, _, _ = context, GetContext_81_result, ctx
	return nil
}
func (ϟa *GlGetBufferParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_82_result := context              // Contextʳ
	ctx := GetContext_82_result                  // Contextʳ
	id := ctx.BoundBuffers.Get(ϟa.Target)        // BufferId
	b := ctx.Instances.Buffers.Get(id)           // Bufferʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(func() (result int32) {
		switch ϟa.Parameter {
		case BufferParameter_GL_BUFFER_SIZE:
			return b.Size
		case BufferParameter_GL_BUFFER_USAGE:
			return int32(b.Usage)
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.Parameter, ϟa))
			return result
		}
	}(), ϟs)
	_, _, _, _, _ = context, GetContext_82_result, ctx, id, b
	return nil
}
func (ϟa *GlCreateShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_83_result := context              // Contextʳ
	ctx := GetContext_83_result                  // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	id := ShaderId(ϟa.Result) // ShaderId
	ctx.Instances.Shaders[id] = func() *Shader {
		s := &Shader{}
		s.Init()
		return s
	}()
	s := ctx.Instances.Shaders.Get(id) // Shaderʳ
	s.Type = ϟa.Type
	ϟa.Result = id
	_, _, _, _, _ = context, GetContext_83_result, ctx, id, s
	return nil
}
func (ϟa *GlDeleteShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_84_result := context              // Contextʳ
	ctx := GetContext_84_result                  // Contextʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // Shaderʳ
	s.Deletable = true
	ctx.Instances.Shaders[ϟa.Shader] = (*Shader)(nil)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = context, GetContext_84_result, ctx, s
	return nil
}
func (ϟa *GlShaderSource) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	sources := ϟa.Source.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // Charᵖˢ
	lengths := ϟa.Length.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // S32ˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                       // Contextʳ
	GetContext_85_result := context                                    // Contextʳ
	ctx := GetContext_85_result                                        // Contextʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)                          // Shaderʳ
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		str := func() (result string) {
			switch ((ϟa.Length) == (S32ᵖ{})) || ((lengths.Index(uint64(i), ϟs).Read(ϟs, ϟd, ϟl)) < (int32(0))) {
			case true:
				return string(sources.Index(uint64(i), ϟs).Read(ϟs, ϟd, ϟl).StringSlice(ϟs, ϟd, ϟl, false).Read(ϟs, ϟd, ϟl))
			case false:
				return string(sources.Index(uint64(i), ϟs).Read(ϟs, ϟd, ϟl).Slice(uint64(int32(0)), uint64(lengths.Index(uint64(i), ϟs).Read(ϟs, ϟd, ϟl)), ϟs).Read(ϟs, ϟd, ϟl))
			default:
				// TODO: better unmatched handling
				panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ((ϟa.Length) == (S32ᵖ{})) || ((lengths.Index(uint64(i), ϟs).Read(ϟs, ϟd, ϟl)) < (int32(0))), ϟa))
				return result
			}
		}() // string
		s.Source += str
		_ = str
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = sources, lengths, context, GetContext_85_result, ctx, s
	return nil
}
func (ϟa *GlShaderBinary) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	for i := int32(int32(0)); i < ϟa.Count; i++ {
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlGetShaderInfoLog) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_86_result := context              // Contextʳ
	ctx := GetContext_86_result                  // Contextʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // Shaderʳ
	min_87_a := ϟa.BufferLength                  // s32
	min_87_b := int32(s.InfoLog.Count)           // s32
	min_87_result := func() (result int32) {
		switch (min_87_a) < (min_87_b) {
		case true:
			return min_87_a
		case false:
			return min_87_b
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (min_87_a) < (min_87_b), ϟa))
			return result
		}
	}() // s32
	l := min_87_result // s32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Info.Slice(uint64(int32(0)), uint64(l), ϟs).Copy(s.InfoLog.Slice(uint64(int32(0)), uint64(l), ϟs), ϟs, ϟd, ϟl)
	ϟa.StringLengthWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(l, ϟs)
	_, _, _, _, _, _, _, _ = context, GetContext_86_result, ctx, s, min_87_a, min_87_b, min_87_result, l
	return nil
}
func (ϟa *GlGetShaderSource) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_88_result := context              // Contextʳ
	ctx := GetContext_88_result                  // Contextʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // Shaderʳ
	min_89_a := ϟa.BufferLength                  // s32
	min_89_b := int32(len(s.Source))             // s32
	min_89_result := func() (result int32) {
		switch (min_89_a) < (min_89_b) {
		case true:
			return min_89_a
		case false:
			return min_89_b
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (min_89_a) < (min_89_b), ϟa))
			return result
		}
	}() // s32
	l := min_89_result // s32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Source.Slice(uint64(int32(0)), uint64(l), ϟs).Copy(MakeCharˢFromString(s.Source, ϟs).Slice(uint64(int32(0)), uint64(l), ϟs), ϟs, ϟd, ϟl)
	ϟa.StringLengthWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(l, ϟs)
	_, _, _, _, _, _, _, _ = context, GetContext_88_result, ctx, s, min_89_a, min_89_b, min_89_result, l
	return nil
}
func (ϟa *GlReleaseShaderCompiler) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlCompileShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlIsShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_90_result := context              // Contextʳ
	ctx := GetContext_90_result                  // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ctx.Instances.Shaders.Contains(ϟa.Shader)
	_, _, _ = context, GetContext_90_result, ctx
	return nil
}
func (ϟa *GlCreateProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_91_result := context              // Contextʳ
	ctx := GetContext_91_result                  // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	id := ProgramId(ϟa.Result) // ProgramId
	ctx.Instances.Programs[id] = func() *Program {
		s := &Program{}
		s.Init()
		return s
	}()
	ϟa.Result = id
	_, _, _, _ = context, GetContext_91_result, ctx, id
	return nil
}
func (ϟa *GlDeleteProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_92_result := context              // Contextʳ
	ctx := GetContext_92_result                  // Contextʳ
	ctx.Instances.Programs[ϟa.Program] = (*Program)(nil)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_92_result, ctx
	return nil
}
func (ϟa *GlAttachShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_93_result := context              // Contextʳ
	ctx := GetContext_93_result                  // Contextʳ
	p := ctx.Instances.Programs.Get(ϟa.Program)  // Programʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // Shaderʳ
	p.Shaders[s.Type] = ϟa.Shader
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = context, GetContext_93_result, ctx, p, s
	return nil
}
func (ϟa *GlDetachShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_94_result := context              // Contextʳ
	ctx := GetContext_94_result                  // Contextʳ
	p := ctx.Instances.Programs.Get(ϟa.Program)  // Programʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // Shaderʳ
	p.Shaders[s.Type] = ShaderId(uint32(0))
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = context, GetContext_94_result, ctx, p, s
	return nil
}
func (ϟa *GlGetAttachedShaders) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_95_result := context              // Contextʳ
	ctx := GetContext_95_result                  // Contextʳ
	p := ctx.Instances.Programs.Get(ϟa.Program)  // Programʳ
	min_96_a := ϟa.BufferLength                  // s32
	min_96_b := int32(len(p.Shaders))            // s32
	min_96_result := func() (result int32) {
		switch (min_96_a) < (min_96_b) {
		case true:
			return min_96_a
		case false:
			return min_96_b
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (min_96_a) < (min_96_b), ϟa))
			return result
		}
	}() // s32
	l := min_96_result // s32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.ShadersLengthWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(l, ϟs)
	_, _, _, _, _, _, _, _ = context, GetContext_95_result, ctx, p, min_96_a, min_96_b, min_96_result, l
	return nil
}
func (ϟa *GlLinkProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlGetProgramInfoLog) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_97_result := context              // Contextʳ
	ctx := GetContext_97_result                  // Contextʳ
	p := ctx.Instances.Programs.Get(ϟa.Program)  // Programʳ
	min_98_a := ϟa.BufferLength                  // s32
	min_98_b := int32(p.InfoLog.Count)           // s32
	min_98_result := func() (result int32) {
		switch (min_98_a) < (min_98_b) {
		case true:
			return min_98_a
		case false:
			return min_98_b
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (min_98_a) < (min_98_b), ϟa))
			return result
		}
	}() // s32
	l := min_98_result // s32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Info.Slice(uint64(int32(0)), uint64(l), ϟs).Copy(p.InfoLog.Slice(uint64(int32(0)), uint64(l), ϟs), ϟs, ϟd, ϟl)
	ϟa.StringLengthWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(l, ϟs)
	_, _, _, _, _, _, _, _ = context, GetContext_97_result, ctx, p, min_98_a, min_98_b, min_98_result, l
	return nil
}
func (ϟa *GlUseProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_99_result := context              // Contextʳ
	ctx := GetContext_99_result                  // Contextʳ
	ctx.BoundProgram = ϟa.Program
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_99_result, ctx
	return nil
}
func (ϟa *GlIsProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_100_result := context             // Contextʳ
	ctx := GetContext_100_result                 // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ctx.Instances.Programs.Contains(ϟa.Program)
	_, _, _ = context, GetContext_100_result, ctx
	return nil
}
func (ϟa *GlValidateProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlClearColor) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_101_result := context             // Contextʳ
	ctx := GetContext_101_result                 // Contextʳ
	ctx.Clearing.ClearColor = func() Color {
		s := Color{}
		s.Init()
		s.Red = ϟa.R
		s.Green = ϟa.G
		s.Blue = ϟa.B
		s.Alpha = ϟa.A
		return s
	}()
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_101_result, ctx
	return nil
}
func (ϟa *GlClearDepthf) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_102_result := context             // Contextʳ
	ctx := GetContext_102_result                 // Contextʳ
	ctx.Clearing.ClearDepth = ϟa.Depth
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_102_result, ctx
	return nil
}
func (ϟa *GlClearStencil) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_103_result := context             // Contextʳ
	ctx := GetContext_103_result                 // Contextʳ
	ctx.Clearing.ClearStencil = ϟa.Stencil
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_103_result, ctx
	return nil
}
func (ϟa *GlClear) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	if (ClearMask_GL_COLOR_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlCullFace) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_104_result := context             // Contextʳ
	ctx := GetContext_104_result                 // Contextʳ
	ctx.Rasterizing.CullFace = ϟa.Mode
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_104_result, ctx
	return nil
}
func (ϟa *GlPolygonOffset) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_105_result := context             // Contextʳ
	ctx := GetContext_105_result                 // Contextʳ
	ctx.Rasterizing.PolygonOffsetUnits = ϟa.Units
	ctx.Rasterizing.PolygonOffsetFactor = ϟa.ScaleFactor
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_105_result, ctx
	return nil
}
func (ϟa *GlLineWidth) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_106_result := context             // Contextʳ
	ctx := GetContext_106_result                 // Contextʳ
	ctx.Rasterizing.LineWidth = ϟa.Width
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_106_result, ctx
	return nil
}
func (ϟa *GlSampleCoverage) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_107_result := context             // Contextʳ
	ctx := GetContext_107_result                 // Contextʳ
	ctx.Rasterizing.SampleCoverageValue = ϟa.Value
	ctx.Rasterizing.SampleCoverageInvert = ϟa.Invert
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_107_result, ctx
	return nil
}
func (ϟa *GlHint) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_108_result := context             // Contextʳ
	ctx := GetContext_108_result                 // Contextʳ
	ctx.GenerateMipmapHint = ϟa.Mode
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_108_result, ctx
	return nil
}
func (ϟa *GlFramebufferRenderbuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_109_result := context             // Contextʳ
	ctx := GetContext_109_result                 // Contextʳ
	target := func() (result FramebufferTarget) {
		switch ϟa.FramebufferTarget {
		case FramebufferTarget_GL_FRAMEBUFFER:
			return FramebufferTarget_GL_DRAW_FRAMEBUFFER
		case FramebufferTarget_GL_DRAW_FRAMEBUFFER:
			return FramebufferTarget_GL_DRAW_FRAMEBUFFER
		case FramebufferTarget_GL_READ_FRAMEBUFFER:
			return FramebufferTarget_GL_READ_FRAMEBUFFER
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.FramebufferTarget, ϟa))
			return result
		}
	}() // FramebufferTarget
	framebufferId := ctx.BoundFramebuffers.Get(target)                  // FramebufferId
	framebuffer := ctx.Instances.Framebuffers.Get(framebufferId)        // Framebufferʳ
	attachment := framebuffer.Attachments.Get(ϟa.FramebufferAttachment) // FramebufferAttachmentInfo
	if (ϟa.Renderbuffer) == (RenderbufferId(uint32(0))) {
		attachment.Type = FramebufferAttachmentType_GL_NONE
	} else {
		attachment.Type = FramebufferAttachmentType_GL_RENDERBUFFER
	}
	attachment.Object = uint32(ϟa.Renderbuffer)
	attachment.TextureLevel = int32(0)
	attachment.CubeMapFace = CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X
	framebuffer.Attachments[ϟa.FramebufferAttachment] = attachment
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = context, GetContext_109_result, ctx, target, framebufferId, framebuffer, attachment
	return nil
}
func (ϟa *GlFramebufferTexture2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_110_result := context             // Contextʳ
	ctx := GetContext_110_result                 // Contextʳ
	target := func() (result FramebufferTarget) {
		switch ϟa.FramebufferTarget {
		case FramebufferTarget_GL_FRAMEBUFFER:
			return FramebufferTarget_GL_DRAW_FRAMEBUFFER
		case FramebufferTarget_GL_DRAW_FRAMEBUFFER:
			return FramebufferTarget_GL_DRAW_FRAMEBUFFER
		case FramebufferTarget_GL_READ_FRAMEBUFFER:
			return FramebufferTarget_GL_READ_FRAMEBUFFER
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.FramebufferTarget, ϟa))
			return result
		}
	}() // FramebufferTarget
	framebufferId := ctx.BoundFramebuffers.Get(target)                  // FramebufferId
	framebuffer := ctx.Instances.Framebuffers.Get(framebufferId)        // Framebufferʳ
	attachment := framebuffer.Attachments.Get(ϟa.FramebufferAttachment) // FramebufferAttachmentInfo
	if (ϟa.Texture) == (TextureId(uint32(0))) {
		attachment.Type = FramebufferAttachmentType_GL_NONE
		attachment.Object = uint32(0)
		attachment.TextureLevel = int32(0)
		attachment.CubeMapFace = CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X
	} else {
		attachment.Type = FramebufferAttachmentType_GL_TEXTURE
		attachment.Object = uint32(ϟa.Texture)
		attachment.TextureLevel = ϟa.Level
		attachment.CubeMapFace = func() (result CubeMapImageTarget) {
			switch ϟa.TextureTarget {
			case TextureImageTarget_GL_TEXTURE_2D:
				return CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X
			case TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X:
				return CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X
			case TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Y:
				return CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Y
			case TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
				return CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Z
			case TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_X:
				return CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_X
			case TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y:
				return CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y
			case TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z:
				return CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z
			default:
				// TODO: better unmatched handling
				panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.TextureTarget, ϟa))
				return result
			}
		}()
	}
	framebuffer.Attachments[ϟa.FramebufferAttachment] = attachment
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = context, GetContext_110_result, ctx, target, framebufferId, framebuffer, attachment
	return nil
}
func (ϟa *GlGetFramebufferAttachmentParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_111_result := context             // Contextʳ
	ctx := GetContext_111_result                 // Contextʳ
	target := func() (result FramebufferTarget) {
		switch ϟa.FramebufferTarget {
		case FramebufferTarget_GL_FRAMEBUFFER:
			return FramebufferTarget_GL_DRAW_FRAMEBUFFER
		case FramebufferTarget_GL_DRAW_FRAMEBUFFER:
			return FramebufferTarget_GL_DRAW_FRAMEBUFFER
		case FramebufferTarget_GL_READ_FRAMEBUFFER:
			return FramebufferTarget_GL_READ_FRAMEBUFFER
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.FramebufferTarget, ϟa))
			return result
		}
	}() // FramebufferTarget
	framebufferId := ctx.BoundFramebuffers.Get(target)           // FramebufferId
	framebuffer := ctx.Instances.Framebuffers.Get(framebufferId) // Framebufferʳ
	a := framebuffer.Attachments.Get(ϟa.Attachment)              // FramebufferAttachmentInfo
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(func() (result int32) {
		switch ϟa.Parameter {
		case FramebufferAttachmentParameter_GL_FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE:
			return int32(a.Type)
		case FramebufferAttachmentParameter_GL_FRAMEBUFFER_ATTACHMENT_OBJECT_NAME:
			return int32(a.Object)
		case FramebufferAttachmentParameter_GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL:
			return a.TextureLevel
		case FramebufferAttachmentParameter_GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE:
			return int32(a.CubeMapFace)
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.Parameter, ϟa))
			return result
		}
	}(), ϟs)
	_, _, _, _, _, _, _ = context, GetContext_111_result, ctx, target, framebufferId, framebuffer, a
	return nil
}
func (ϟa *GlDrawElements) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                     // Contextʳ
	GetContext_112_result := context                                 // Contextʳ
	ctx := GetContext_112_result                                     // Contextʳ
	count := uint32(ϟa.ElementCount)                                 // u32
	id := ctx.BoundBuffers.Get(BufferTarget_GL_ELEMENT_ARRAY_BUFFER) // BufferId
	if (id) != (BufferId(uint32(0))) {
		index_data := ctx.Instances.Buffers.Get(id).Data                                                   // U8ˢ
		offset := uint32(uint64(ϟa.Indices.Address))                                                       // u32
		first := externs{ϟs, ϟd, ϟl}.minIndex(U8ᵖ(index_data.Index(0, ϟs)), ϟa.IndicesType, offset, count) // u32
		last := externs{ϟs, ϟd, ϟl}.maxIndex(U8ᵖ(index_data.Index(0, ϟs)), ϟa.IndicesType, offset, count)  // u32
		ReadVertexArrays_113_ctx := ctx                                                                    // Contextʳ
		ReadVertexArrays_113_first_index := first                                                          // u32
		ReadVertexArrays_113_last_index := last                                                            // u32
		for i := int32(int32(0)); i < int32(len(ReadVertexArrays_113_ctx.VertexAttributeArrays)); i++ {
			arr := ReadVertexArrays_113_ctx.VertexAttributeArrays.Get(AttributeLocation(i)) // VertexAttributeArrayʳ
			if (arr.Enabled) && ((arr.Buffer) == (BufferId(uint32(0)))) {
				vertexAttribTypeSize_114_t := arr.Type // VertexAttribType
				vertexAttribTypeSize_114_result := func() (result uint32) {
					switch vertexAttribTypeSize_114_t {
					case VertexAttribType_GL_BYTE:
						return uint32(1)
					case VertexAttribType_GL_UNSIGNED_BYTE:
						return uint32(1)
					case VertexAttribType_GL_SHORT:
						return uint32(2)
					case VertexAttribType_GL_UNSIGNED_SHORT:
						return uint32(2)
					case VertexAttribType_GL_FIXED:
						return uint32(4)
					case VertexAttribType_GL_FLOAT:
						return uint32(4)
					case VertexAttribType_GL_ARB_half_float_vertex:
						return uint32(2)
					case VertexAttribType_GL_HALF_FLOAT_OES:
						return uint32(2)
					default:
						// TODO: better unmatched handling
						panic(fmt.Errorf("Unmatched switch(%v) in atom %T", vertexAttribTypeSize_114_t, ϟa))
						return result
					}
				}() // u32
				elsize := (vertexAttribTypeSize_114_result) * (arr.Size) // u32
				elstride := func() (result uint32) {
					switch (arr.Stride) == (int32(0)) {
					case true:
						return elsize
					case false:
						return uint32(arr.Stride)
					default:
						// TODO: better unmatched handling
						panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (arr.Stride) == (int32(0)), ϟa))
						return result
					}
				}() // u32
				for v := uint32(ReadVertexArrays_113_first_index); v < (ReadVertexArrays_113_last_index)+(uint32(1)); v++ {
					offset := (elstride) * (v) // u32
					_ = offset
				}
				_, _, _, _ = vertexAttribTypeSize_114_t, vertexAttribTypeSize_114_result, elsize, elstride
			}
			_ = arr
		}
		_, _, _, _, _, _, _ = index_data, offset, first, last, ReadVertexArrays_113_ctx, ReadVertexArrays_113_first_index, ReadVertexArrays_113_last_index
	} else {
		index_data := U8ᵖ(ϟa.Indices)                                                       // U8ᵖ
		first := externs{ϟs, ϟd, ϟl}.minIndex(index_data, ϟa.IndicesType, uint32(0), count) // u32
		last := externs{ϟs, ϟd, ϟl}.maxIndex(index_data, ϟa.IndicesType, uint32(0), count)  // u32
		ReadVertexArrays_115_ctx := ctx                                                     // Contextʳ
		ReadVertexArrays_115_first_index := first                                           // u32
		ReadVertexArrays_115_last_index := last                                             // u32
		for i := int32(int32(0)); i < int32(len(ReadVertexArrays_115_ctx.VertexAttributeArrays)); i++ {
			arr := ReadVertexArrays_115_ctx.VertexAttributeArrays.Get(AttributeLocation(i)) // VertexAttributeArrayʳ
			if (arr.Enabled) && ((arr.Buffer) == (BufferId(uint32(0)))) {
				vertexAttribTypeSize_116_t := arr.Type // VertexAttribType
				vertexAttribTypeSize_116_result := func() (result uint32) {
					switch vertexAttribTypeSize_116_t {
					case VertexAttribType_GL_BYTE:
						return uint32(1)
					case VertexAttribType_GL_UNSIGNED_BYTE:
						return uint32(1)
					case VertexAttribType_GL_SHORT:
						return uint32(2)
					case VertexAttribType_GL_UNSIGNED_SHORT:
						return uint32(2)
					case VertexAttribType_GL_FIXED:
						return uint32(4)
					case VertexAttribType_GL_FLOAT:
						return uint32(4)
					case VertexAttribType_GL_ARB_half_float_vertex:
						return uint32(2)
					case VertexAttribType_GL_HALF_FLOAT_OES:
						return uint32(2)
					default:
						// TODO: better unmatched handling
						panic(fmt.Errorf("Unmatched switch(%v) in atom %T", vertexAttribTypeSize_116_t, ϟa))
						return result
					}
				}() // u32
				elsize := (vertexAttribTypeSize_116_result) * (arr.Size) // u32
				elstride := func() (result uint32) {
					switch (arr.Stride) == (int32(0)) {
					case true:
						return elsize
					case false:
						return uint32(arr.Stride)
					default:
						// TODO: better unmatched handling
						panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (arr.Stride) == (int32(0)), ϟa))
						return result
					}
				}() // u32
				for v := uint32(ReadVertexArrays_115_first_index); v < (ReadVertexArrays_115_last_index)+(uint32(1)); v++ {
					offset := (elstride) * (v) // u32
					_ = offset
				}
				_, _, _, _ = vertexAttribTypeSize_116_t, vertexAttribTypeSize_116_result, elsize, elstride
			}
			_ = arr
		}
		IndexSize_117_indices_type := ϟa.IndicesType // IndicesType
		IndexSize_117_result := func() (result uint32) {
			switch IndexSize_117_indices_type {
			case IndicesType_GL_UNSIGNED_BYTE:
				return uint32(1)
			case IndicesType_GL_UNSIGNED_SHORT:
				return uint32(2)
			case IndicesType_GL_UNSIGNED_INT:
				return uint32(4)
			default:
				// TODO: better unmatched handling
				panic(fmt.Errorf("Unmatched switch(%v) in atom %T", IndexSize_117_indices_type, ϟa))
				return result
			}
		}() // u32
		_, _, _, _, _, _, _, _ = index_data, first, last, ReadVertexArrays_115_ctx, ReadVertexArrays_115_first_index, ReadVertexArrays_115_last_index, IndexSize_117_indices_type, IndexSize_117_result
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = context, GetContext_112_result, ctx, count, id
	return nil
}
func (ϟa *GlDrawArrays) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                   // Contextʳ
	GetContext_118_result := context                               // Contextʳ
	ctx := GetContext_118_result                                   // Contextʳ
	last_index := (ϟa.FirstIndex) + ((ϟa.IndexCount) - (int32(1))) // s32
	ReadVertexArrays_119_ctx := ctx                                // Contextʳ
	ReadVertexArrays_119_first_index := uint32(ϟa.FirstIndex)      // u32
	ReadVertexArrays_119_last_index := uint32(last_index)          // u32
	for i := int32(int32(0)); i < int32(len(ReadVertexArrays_119_ctx.VertexAttributeArrays)); i++ {
		arr := ReadVertexArrays_119_ctx.VertexAttributeArrays.Get(AttributeLocation(i)) // VertexAttributeArrayʳ
		if (arr.Enabled) && ((arr.Buffer) == (BufferId(uint32(0)))) {
			vertexAttribTypeSize_120_t := arr.Type // VertexAttribType
			vertexAttribTypeSize_120_result := func() (result uint32) {
				switch vertexAttribTypeSize_120_t {
				case VertexAttribType_GL_BYTE:
					return uint32(1)
				case VertexAttribType_GL_UNSIGNED_BYTE:
					return uint32(1)
				case VertexAttribType_GL_SHORT:
					return uint32(2)
				case VertexAttribType_GL_UNSIGNED_SHORT:
					return uint32(2)
				case VertexAttribType_GL_FIXED:
					return uint32(4)
				case VertexAttribType_GL_FLOAT:
					return uint32(4)
				case VertexAttribType_GL_ARB_half_float_vertex:
					return uint32(2)
				case VertexAttribType_GL_HALF_FLOAT_OES:
					return uint32(2)
				default:
					// TODO: better unmatched handling
					panic(fmt.Errorf("Unmatched switch(%v) in atom %T", vertexAttribTypeSize_120_t, ϟa))
					return result
				}
			}() // u32
			elsize := (vertexAttribTypeSize_120_result) * (arr.Size) // u32
			elstride := func() (result uint32) {
				switch (arr.Stride) == (int32(0)) {
				case true:
					return elsize
				case false:
					return uint32(arr.Stride)
				default:
					// TODO: better unmatched handling
					panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (arr.Stride) == (int32(0)), ϟa))
					return result
				}
			}() // u32
			for v := uint32(ReadVertexArrays_119_first_index); v < (ReadVertexArrays_119_last_index)+(uint32(1)); v++ {
				offset := (elstride) * (v) // u32
				_ = offset
			}
			_, _, _, _ = vertexAttribTypeSize_120_t, vertexAttribTypeSize_120_result, elsize, elstride
		}
		_ = arr
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = context, GetContext_118_result, ctx, last_index, ReadVertexArrays_119_ctx, ReadVertexArrays_119_first_index, ReadVertexArrays_119_last_index
	return nil
}
func (ϟa *GlFlush) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlFinish) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlGetBooleanv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	v := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs) // Boolˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                                        // Contextʳ
	GetContext_121_result := context                                                                    // Contextʳ
	ctx := GetContext_121_result                                                                        // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	switch ϟa.Param {
	case StateVariable_GL_BLEND:
		v.Index(uint64(0), ϟs).Write(ctx.Capabilities.Get(Capability_GL_BLEND), ϟs)
	case StateVariable_GL_CULL_FACE:
		v.Index(uint64(0), ϟs).Write(ctx.Capabilities.Get(Capability_GL_CULL_FACE), ϟs)
	case StateVariable_GL_DEPTH_TEST:
		v.Index(uint64(0), ϟs).Write(ctx.Capabilities.Get(Capability_GL_DEPTH_TEST), ϟs)
	case StateVariable_GL_DITHER:
		v.Index(uint64(0), ϟs).Write(ctx.Capabilities.Get(Capability_GL_DITHER), ϟs)
	case StateVariable_GL_POLYGON_OFFSET_FILL:
		v.Index(uint64(0), ϟs).Write(ctx.Capabilities.Get(Capability_GL_POLYGON_OFFSET_FILL), ϟs)
	case StateVariable_GL_SAMPLE_ALPHA_TO_COVERAGE:
		v.Index(uint64(0), ϟs).Write(ctx.Capabilities.Get(Capability_GL_SAMPLE_ALPHA_TO_COVERAGE), ϟs)
	case StateVariable_GL_SAMPLE_COVERAGE:
		v.Index(uint64(0), ϟs).Write(ctx.Capabilities.Get(Capability_GL_SAMPLE_COVERAGE), ϟs)
	case StateVariable_GL_SCISSOR_TEST:
		v.Index(uint64(0), ϟs).Write(ctx.Capabilities.Get(Capability_GL_SCISSOR_TEST), ϟs)
	case StateVariable_GL_STENCIL_TEST:
		v.Index(uint64(0), ϟs).Write(ctx.Capabilities.Get(Capability_GL_STENCIL_TEST), ϟs)
	case StateVariable_GL_DEPTH_WRITEMASK:
		v.Index(uint64(0), ϟs).Write(ctx.Rasterizing.DepthMask, ϟs)
	case StateVariable_GL_COLOR_WRITEMASK:
		v.Index(uint64(0), ϟs).Write(ctx.Rasterizing.ColorMaskRed, ϟs)
		v.Index(uint64(1), ϟs).Write(ctx.Rasterizing.ColorMaskGreen, ϟs)
		v.Index(uint64(2), ϟs).Write(ctx.Rasterizing.ColorMaskBlue, ϟs)
		v.Index(uint64(3), ϟs).Write(ctx.Rasterizing.ColorMaskAlpha, ϟs)
	case StateVariable_GL_SAMPLE_COVERAGE_INVERT:
		v.Index(uint64(0), ϟs).Write(ctx.Rasterizing.SampleCoverageInvert, ϟs)
	case StateVariable_GL_SHADER_COMPILER:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
	case StateVariable_GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
	default:
		v := ϟa.Param
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	_, _, _, _ = v, context, GetContext_121_result, ctx
	return nil
}
func (ϟa *GlGetFloatv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	v := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs) // F32ˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                                        // Contextʳ
	GetContext_122_result := context                                                                    // Contextʳ
	ctx := GetContext_122_result                                                                        // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	switch ϟa.Param {
	case StateVariable_GL_DEPTH_RANGE:
		v.Index(uint64(0), ϟs).Write(ctx.Rasterizing.DepthNear, ϟs)
		v.Index(uint64(1), ϟs).Write(ctx.Rasterizing.DepthFar, ϟs)
	case StateVariable_GL_LINE_WIDTH:
		v.Index(uint64(0), ϟs).Write(ctx.Rasterizing.LineWidth, ϟs)
	case StateVariable_GL_POLYGON_OFFSET_FACTOR:
		v.Index(uint64(0), ϟs).Write(ctx.Rasterizing.PolygonOffsetFactor, ϟs)
	case StateVariable_GL_POLYGON_OFFSET_UNITS:
		v.Index(uint64(0), ϟs).Write(ctx.Rasterizing.PolygonOffsetUnits, ϟs)
	case StateVariable_GL_SAMPLE_COVERAGE_VALUE:
		v.Index(uint64(0), ϟs).Write(ctx.Rasterizing.SampleCoverageValue, ϟs)
	case StateVariable_GL_COLOR_CLEAR_VALUE:
		v.Index(uint64(0), ϟs).Write(ctx.Clearing.ClearColor.Red, ϟs)
		v.Index(uint64(1), ϟs).Write(ctx.Clearing.ClearColor.Green, ϟs)
		v.Index(uint64(2), ϟs).Write(ctx.Clearing.ClearColor.Blue, ϟs)
		v.Index(uint64(3), ϟs).Write(ctx.Clearing.ClearColor.Alpha, ϟs)
	case StateVariable_GL_DEPTH_CLEAR_VALUE:
		v.Index(uint64(0), ϟs).Write(ctx.Clearing.ClearDepth, ϟs)
	case StateVariable_GL_ALIASED_LINE_WIDTH_RANGE:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
		v.Index(uint64(1), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(1), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
	case StateVariable_GL_ALIASED_POINT_SIZE_RANGE:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
		v.Index(uint64(1), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(1), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
	case StateVariable_GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
	default:
		v := ϟa.Param
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	_, _, _, _ = v, context, GetContext_122_result, ctx
	return nil
}
func (ϟa *GlGetIntegerv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	v := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs) // S32ˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                                        // Contextʳ
	GetContext_123_result := context                                                                    // Contextʳ
	ctx := GetContext_123_result                                                                        // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	switch ϟa.Param {
	case StateVariable_GL_ACTIVE_TEXTURE:
		v.Index(uint64(0), ϟs).Write(int32(ctx.ActiveTextureUnit), ϟs)
	case StateVariable_GL_ARRAY_BUFFER_BINDING:
		v.Index(uint64(0), ϟs).Write(int32(ctx.BoundBuffers.Get(BufferTarget_GL_ARRAY_BUFFER)), ϟs)
	case StateVariable_GL_ELEMENT_ARRAY_BUFFER_BINDING:
		v.Index(uint64(0), ϟs).Write(int32(ctx.BoundBuffers.Get(BufferTarget_GL_ELEMENT_ARRAY_BUFFER)), ϟs)
	case StateVariable_GL_BLEND_SRC_ALPHA:
		v.Index(uint64(0), ϟs).Write(int32(ctx.Blending.SrcAlphaBlendFactor), ϟs)
	case StateVariable_GL_BLEND_SRC_RGB:
		v.Index(uint64(0), ϟs).Write(int32(ctx.Blending.SrcRgbBlendFactor), ϟs)
	case StateVariable_GL_BLEND_DST_ALPHA:
		v.Index(uint64(0), ϟs).Write(int32(ctx.Blending.DstAlphaBlendFactor), ϟs)
	case StateVariable_GL_BLEND_DST_RGB:
		v.Index(uint64(0), ϟs).Write(int32(ctx.Blending.DstRgbBlendFactor), ϟs)
	case StateVariable_GL_BLEND_EQUATION_RGB:
		v.Index(uint64(0), ϟs).Write(int32(ctx.Blending.BlendEquationRgb), ϟs)
	case StateVariable_GL_BLEND_EQUATION_ALPHA:
		v.Index(uint64(0), ϟs).Write(int32(ctx.Blending.BlendEquationAlpha), ϟs)
	case StateVariable_GL_BLEND_COLOR:
		v.Index(uint64(0), ϟs).Write(int32(ctx.Blending.BlendColor.Red), ϟs)
		v.Index(uint64(1), ϟs).Write(int32(ctx.Blending.BlendColor.Green), ϟs)
		v.Index(uint64(2), ϟs).Write(int32(ctx.Blending.BlendColor.Blue), ϟs)
		v.Index(uint64(3), ϟs).Write(int32(ctx.Blending.BlendColor.Alpha), ϟs)
	case StateVariable_GL_DEPTH_FUNC:
		v.Index(uint64(0), ϟs).Write(int32(ctx.Rasterizing.DepthTestFunction), ϟs)
	case StateVariable_GL_DEPTH_CLEAR_VALUE:
		v.Index(uint64(0), ϟs).Write(int32(ctx.Clearing.ClearDepth), ϟs)
	case StateVariable_GL_STENCIL_WRITEMASK:
		v.Index(uint64(0), ϟs).Write(int32(ctx.Rasterizing.StencilMask.Get(FaceMode_GL_FRONT)), ϟs)
	case StateVariable_GL_STENCIL_BACK_WRITEMASK:
		v.Index(uint64(0), ϟs).Write(int32(ctx.Rasterizing.StencilMask.Get(FaceMode_GL_BACK)), ϟs)
	case StateVariable_GL_VIEWPORT:
		v.Index(uint64(0), ϟs).Write(ctx.Rasterizing.Viewport.X, ϟs)
		v.Index(uint64(1), ϟs).Write(ctx.Rasterizing.Viewport.Y, ϟs)
		v.Index(uint64(2), ϟs).Write(ctx.Rasterizing.Viewport.Width, ϟs)
		v.Index(uint64(3), ϟs).Write(ctx.Rasterizing.Viewport.Height, ϟs)
	case StateVariable_GL_SCISSOR_BOX:
		v.Index(uint64(0), ϟs).Write(ctx.Rasterizing.Scissor.X, ϟs)
		v.Index(uint64(1), ϟs).Write(ctx.Rasterizing.Scissor.Y, ϟs)
		v.Index(uint64(2), ϟs).Write(ctx.Rasterizing.Scissor.Width, ϟs)
		v.Index(uint64(3), ϟs).Write(ctx.Rasterizing.Scissor.Height, ϟs)
	case StateVariable_GL_FRONT_FACE:
		v.Index(uint64(0), ϟs).Write(int32(ctx.Rasterizing.FrontFace), ϟs)
	case StateVariable_GL_CULL_FACE_MODE:
		v.Index(uint64(0), ϟs).Write(int32(ctx.Rasterizing.CullFace), ϟs)
	case StateVariable_GL_STENCIL_CLEAR_VALUE:
		v.Index(uint64(0), ϟs).Write(ctx.Clearing.ClearStencil, ϟs)
	case StateVariable_GL_FRAMEBUFFER_BINDING:
		v.Index(uint64(0), ϟs).Write(int32(ctx.BoundFramebuffers.Get(FramebufferTarget_GL_FRAMEBUFFER)), ϟs)
	case StateVariable_GL_READ_FRAMEBUFFER_BINDING:
		v.Index(uint64(0), ϟs).Write(int32(ctx.BoundFramebuffers.Get(FramebufferTarget_GL_READ_FRAMEBUFFER)), ϟs)
	case StateVariable_GL_RENDERBUFFER_BINDING:
		v.Index(uint64(0), ϟs).Write(int32(ctx.BoundRenderbuffers.Get(RenderbufferTarget_GL_RENDERBUFFER)), ϟs)
	case StateVariable_GL_CURRENT_PROGRAM:
		v.Index(uint64(0), ϟs).Write(int32(ctx.BoundProgram), ϟs)
	case StateVariable_GL_TEXTURE_BINDING_2D:
		v.Index(uint64(0), ϟs).Write(int32(ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_2D)), ϟs)
	case StateVariable_GL_TEXTURE_BINDING_CUBE_MAP:
		v.Index(uint64(0), ϟs).Write(int32(ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_CUBE_MAP)), ϟs)
	case StateVariable_GL_GENERATE_MIPMAP_HINT:
		v.Index(uint64(0), ϟs).Write(int32(ctx.GenerateMipmapHint), ϟs)
	case StateVariable_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl) // any
		v.Index(uint64(0), ϟs).Write(result, ϟs)
		_ = result
	case StateVariable_GL_MAX_CUBE_MAP_TEXTURE_SIZE:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl) // any
		v.Index(uint64(0), ϟs).Write(result, ϟs)
		_ = result
	case StateVariable_GL_MAX_FRAGMENT_UNIFORM_VECTORS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl) // any
		v.Index(uint64(0), ϟs).Write(result, ϟs)
		_ = result
	case StateVariable_GL_MAX_RENDERBUFFER_SIZE:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl) // any
		v.Index(uint64(0), ϟs).Write(result, ϟs)
		_ = result
	case StateVariable_GL_MAX_TEXTURE_IMAGE_UNITS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl) // any
		v.Index(uint64(0), ϟs).Write(result, ϟs)
		_ = result
	case StateVariable_GL_MAX_TEXTURE_SIZE:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl) // any
		v.Index(uint64(0), ϟs).Write(result, ϟs)
		_ = result
	case StateVariable_GL_MAX_VARYING_VECTORS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl) // any
		v.Index(uint64(0), ϟs).Write(result, ϟs)
		_ = result
	case StateVariable_GL_MAX_VERTEX_ATTRIBS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl) // any
		v.Index(uint64(0), ϟs).Write(result, ϟs)
		_ = result
	case StateVariable_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
	case StateVariable_GL_MAX_VERTEX_UNIFORM_VECTORS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl) // any
		v.Index(uint64(0), ϟs).Write(result, ϟs)
		_ = result
	case StateVariable_GL_MAX_VIEWPORT_DIMS:
		max_width := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl)  // any
		max_height := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(1), ϟs).Read(ϟs, ϟd, ϟl) // any
		v.Index(uint64(0), ϟs).Write(max_width, ϟs)
		v.Index(uint64(1), ϟs).Write(max_height, ϟs)
		_, _ = max_width, max_height
	case StateVariable_GL_NUM_COMPRESSED_TEXTURE_FORMATS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl) // any
		v.Index(uint64(0), ϟs).Write(result, ϟs)
		_ = result
	case StateVariable_GL_NUM_SHADER_BINARY_FORMATS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl) // any
		v.Index(uint64(0), ϟs).Write(result, ϟs)
		_ = result
	case StateVariable_GL_PACK_ALIGNMENT:
		v.Index(uint64(0), ϟs).Write(ctx.PixelStorage.Get(PixelStoreParameter_GL_PACK_ALIGNMENT), ϟs)
	case StateVariable_GL_UNPACK_ALIGNMENT:
		v.Index(uint64(0), ϟs).Write(ctx.PixelStorage.Get(PixelStoreParameter_GL_UNPACK_ALIGNMENT), ϟs)
	case StateVariable_GL_ALPHA_BITS:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
	case StateVariable_GL_BLUE_BITS:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
	case StateVariable_GL_GREEN_BITS:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
	case StateVariable_GL_RED_BITS:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
	case StateVariable_GL_DEPTH_BITS:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
	case StateVariable_GL_SAMPLE_BUFFERS:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
	case StateVariable_GL_SAMPLES:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
	case StateVariable_GL_STENCIL_BITS:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
	case StateVariable_GL_SUBPIXEL_BITS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl) // any
		v.Index(uint64(0), ϟs).Write(result, ϟs)
		_ = result
	case StateVariable_GL_IMPLEMENTATION_COLOR_READ_FORMAT:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
	case StateVariable_GL_IMPLEMENTATION_COLOR_READ_TYPE:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
	case StateVariable_GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
	case StateVariable_GL_GPU_DISJOINT_EXT:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
	default:
		v := ϟa.Param
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	_, _, _, _ = v, context, GetContext_123_result, ctx
	return nil
}
func (ϟa *GlGetString) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = Charᵖ{}
	return nil
}
func (ϟa *GlEnable) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_124_result := context             // Contextʳ
	ctx := GetContext_124_result                 // Contextʳ
	ctx.Capabilities[ϟa.Capability] = true
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_124_result, ctx
	return nil
}
func (ϟa *GlDisable) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_125_result := context             // Contextʳ
	ctx := GetContext_125_result                 // Contextʳ
	ctx.Capabilities[ϟa.Capability] = false
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_125_result, ctx
	return nil
}
func (ϟa *GlIsEnabled) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_126_result := context             // Contextʳ
	ctx := GetContext_126_result                 // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ctx.Capabilities.Get(ϟa.Capability)
	_, _, _ = context, GetContext_126_result, ctx
	return nil
}
func (ϟa *GlMapBufferRange) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	return nil
}
func (ϟa *GlUnmapBuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlInvalidateFramebuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlRenderbufferStorageMultisample) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlBlitFramebuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlGenQueries) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	q := ϟa.Queries.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // QueryIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                  // Contextʳ
	GetContext_127_result := context                              // Contextʳ
	ctx := GetContext_127_result                                  // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		id := QueryId(ϟa.Queries.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).Read(ϟs, ϟd, ϟl)) // QueryId
		ctx.Instances.Queries[id] = func() *Query {
			s := &Query{}
			s.Init()
			return s
		}()
		q.Index(uint64(i), ϟs).Write(id, ϟs)
		_ = id
	}
	_, _, _, _ = q, context, GetContext_127_result, ctx
	return nil
}
func (ϟa *GlBeginQuery) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlEndQuery) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlDeleteQueries) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	q := ϟa.Queries.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // QueryIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                  // Contextʳ
	GetContext_128_result := context                              // Contextʳ
	ctx := GetContext_128_result                                  // Contextʳ
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		ctx.Instances.Queries[q.Index(uint64(i), ϟs).Read(ϟs, ϟd, ϟl)] = (*Query)(nil)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = q, context, GetContext_128_result, ctx
	return nil
}
func (ϟa *GlIsQuery) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_129_result := context             // Contextʳ
	ctx := GetContext_129_result                 // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ctx.Instances.Queries.Contains(ϟa.Query)
	_, _, _ = context, GetContext_129_result, ctx
	return nil
}
func (ϟa *GlGetQueryiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
	return nil
}
func (ϟa *GlGetQueryObjectuiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
	return nil
}
func (ϟa *GlGenQueriesEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	q := ϟa.Queries.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // QueryIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                  // Contextʳ
	GetContext_130_result := context                              // Contextʳ
	ctx := GetContext_130_result                                  // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		id := QueryId(ϟa.Queries.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).Read(ϟs, ϟd, ϟl)) // QueryId
		ctx.Instances.Queries[id] = func() *Query {
			s := &Query{}
			s.Init()
			return s
		}()
		q.Index(uint64(i), ϟs).Write(id, ϟs)
		_ = id
	}
	_, _, _, _ = q, context, GetContext_130_result, ctx
	return nil
}
func (ϟa *GlBeginQueryEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlEndQueryEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlDeleteQueriesEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	q := ϟa.Queries.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // QueryIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                  // Contextʳ
	GetContext_131_result := context                              // Contextʳ
	ctx := GetContext_131_result                                  // Contextʳ
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		ctx.Instances.Queries[q.Index(uint64(i), ϟs).Read(ϟs, ϟd, ϟl)] = (*Query)(nil)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = q, context, GetContext_131_result, ctx
	return nil
}
func (ϟa *GlIsQueryEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_132_result := context             // Contextʳ
	ctx := GetContext_132_result                 // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ctx.Instances.Queries.Contains(ϟa.Query)
	_, _, _ = context, GetContext_132_result, ctx
	return nil
}
func (ϟa *GlQueryCounterEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlGetQueryivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
	return nil
}
func (ϟa *GlGetQueryObjectivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
	return nil
}
func (ϟa *GlGetQueryObjectuivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
	return nil
}
func (ϟa *GlGetQueryObjecti64vEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
	return nil
}
func (ϟa *GlGetQueryObjectui64vEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
	return nil
}
func (ϟa *ReplayCreateRenderer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *ReplayBindRenderer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *BackbufferInfo) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                                                 // Contextʳ
	GetContext_133_result := context                                                                             // Contextʳ
	ctx := GetContext_133_result                                                                                 // Contextʳ
	backbuffer := ctx.Instances.Framebuffers.Get(FramebufferId(uint32(0)))                                       // Framebufferʳ
	color_id := RenderbufferId(backbuffer.Attachments.Get(FramebufferAttachment_GL_COLOR_ATTACHMENT0).Object)    // RenderbufferId
	color_buffer := ctx.Instances.Renderbuffers.Get(color_id)                                                    // Renderbufferʳ
	depth_id := RenderbufferId(backbuffer.Attachments.Get(FramebufferAttachment_GL_DEPTH_ATTACHMENT).Object)     // RenderbufferId
	depth_buffer := ctx.Instances.Renderbuffers.Get(depth_id)                                                    // Renderbufferʳ
	stencil_id := RenderbufferId(backbuffer.Attachments.Get(FramebufferAttachment_GL_STENCIL_ATTACHMENT).Object) // RenderbufferId
	stencil_buffer := ctx.Instances.Renderbuffers.Get(stencil_id)                                                // Renderbufferʳ
	color_buffer.Width = ϟa.Width
	color_buffer.Height = ϟa.Height
	color_buffer.Format = ϟa.ColorFmt
	depth_buffer.Width = ϟa.Width
	depth_buffer.Height = ϟa.Height
	depth_buffer.Format = ϟa.DepthFmt
	stencil_buffer.Width = ϟa.Width
	stencil_buffer.Height = ϟa.Height
	stencil_buffer.Format = ϟa.StencilFmt
	if ϟa.ResetViewportScissor {
		ctx.Rasterizing.Scissor.Width = ϟa.Width
		ctx.Rasterizing.Scissor.Height = ϟa.Height
		ctx.Rasterizing.Viewport.Width = ϟa.Width
		ctx.Rasterizing.Viewport.Height = ϟa.Height
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _, _, _ = context, GetContext_133_result, ctx, backbuffer, color_id, color_buffer, depth_id, depth_buffer, stencil_id, stencil_buffer
	return nil
}
func (ϟa *StartTimer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *StopTimer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	return nil
}
func (ϟa *FlushPostBuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
