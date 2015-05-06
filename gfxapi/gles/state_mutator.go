////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

package gles

import (
	"fmt"
	"log"
	"reflect"

	"android.googlesource.com/platform/tools/gpu/gfxapi"
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

func (ϟa *ReplayCreateRenderer) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := ReplayCreateRenderer{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying replayCreateRenderer expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *ReplayBindRenderer) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := ReplayBindRenderer{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying replayBindRenderer expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *BackbufferInfo) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := BackbufferInfo{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                                                 // ContextPtr
	GetContext_0_result := context                                                                               // ContextPtr
	ctx := GetContext_0_result                                                                                   // ContextPtr
	backbuffer := ctx.Instances.Framebuffers.Get(FramebufferId(uint32(0)))                                       // FramebufferPtr
	color_id := RenderbufferId(backbuffer.Attachments.Get(FramebufferAttachment_GL_COLOR_ATTACHMENT0).Object)    // RenderbufferId
	color_buffer := ctx.Instances.Renderbuffers.Get(color_id)                                                    // RenderbufferPtr
	depth_id := RenderbufferId(backbuffer.Attachments.Get(FramebufferAttachment_GL_DEPTH_ATTACHMENT).Object)     // RenderbufferId
	depth_buffer := ctx.Instances.Renderbuffers.Get(depth_id)                                                    // RenderbufferPtr
	stencil_id := RenderbufferId(backbuffer.Attachments.Get(FramebufferAttachment_GL_STENCIL_ATTACHMENT).Object) // RenderbufferId
	stencil_buffer := ctx.Instances.Renderbuffers.Get(stencil_id)                                                // RenderbufferPtr
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
	_, _, _, _, _, _, _, _, _, _ = context, GetContext_0_result, ctx, backbuffer, color_id, color_buffer, depth_id, depth_buffer, stencil_id, stencil_buffer
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying backbufferInfo expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *StartTimer) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := StartTimer{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying startTimer expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *StopTimer) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := StopTimer{}
	ϟa.Result = ϟa.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying stopTimer expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *FlushPostBuffer) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := FlushPostBuffer{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying flushPostBuffer expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *EglInitialize) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := EglInitialize{}
	ϟa.Major = ϟa.Major
	ϟa.Minor = ϟa.Minor
	ϟa.Result = ϟa.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying eglInitialize expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *EglCreateContext) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := EglCreateContext{}
	context := EGLContext(ϟa.Result) // EGLContext
	identifier := ϟc.NextContextID   // ContextID
	ϟc.NextContextID = (ϟc.NextContextID) + (ContextID(uint32(1)))
	ctx := func() *Context {
		s := &Context{}
		s.Init()
		return s
	}() // ContextPtr
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
	}() // FramebufferPtr
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
	CreateContext_1_result := ctx // ContextPtr
	ϟc.EGLContexts[context] = CreateContext_1_result
	ϟa.Result = context
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_1_result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying eglCreateContext expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *EglMakeCurrent) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := EglMakeCurrent{}
	SetContext_2_context := ϟc.EGLContexts.Get(ϟa.Context) // ContextPtr
	ϟc.Contexts[ϟc.CurrentThread] = SetContext_2_context
	ϟa.Result = ϟa.Result
	_ = SetContext_2_context
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying eglMakeCurrent expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *EglSwapBuffers) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := EglSwapBuffers{}
	ϟa.Result = ϟa.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying eglSwapBuffers expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *EglQuerySurface) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := EglQuerySurface{}
	ϟa.Value = ϟa.Value
	ϟa.Result = ϟa.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying eglQuerySurface expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlXCreateContext) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlXCreateContext{}
	context := GLXContext(ϟa.Result) // GLXContext
	identifier := ϟc.NextContextID   // ContextID
	ϟc.NextContextID = (ϟc.NextContextID) + (ContextID(uint32(1)))
	ctx := func() *Context {
		s := &Context{}
		s.Init()
		return s
	}() // ContextPtr
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
	}() // FramebufferPtr
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
	CreateContext_3_result := ctx // ContextPtr
	ϟc.GLXContexts[context] = CreateContext_3_result
	ϟa.Result = context
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_3_result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glXCreateContext expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlXCreateNewContext) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlXCreateNewContext{}
	context := GLXContext(ϟa.Result) // GLXContext
	identifier := ϟc.NextContextID   // ContextID
	ϟc.NextContextID = (ϟc.NextContextID) + (ContextID(uint32(1)))
	ctx := func() *Context {
		s := &Context{}
		s.Init()
		return s
	}() // ContextPtr
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
	}() // FramebufferPtr
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
	CreateContext_4_result := ctx // ContextPtr
	ϟc.GLXContexts[context] = CreateContext_4_result
	ϟa.Result = context
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_4_result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glXCreateNewContext expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlXMakeContextCurrent) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlXMakeContextCurrent{}
	SetContext_5_context := ϟc.GLXContexts.Get(ϟa.Ctx) // ContextPtr
	ϟc.Contexts[ϟc.CurrentThread] = SetContext_5_context
	_ = SetContext_5_context
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glXMakeContextCurrent expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlXSwapBuffers) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlXSwapBuffers{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glXSwapBuffers expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *WglCreateContext) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := WglCreateContext{}
	context := HGLRC(ϟa.Result)    // HGLRC
	identifier := ϟc.NextContextID // ContextID
	ϟc.NextContextID = (ϟc.NextContextID) + (ContextID(uint32(1)))
	ctx := func() *Context {
		s := &Context{}
		s.Init()
		return s
	}() // ContextPtr
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
	}() // FramebufferPtr
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
	CreateContext_6_result := ctx // ContextPtr
	ϟc.WGLContexts[context] = CreateContext_6_result
	ϟa.Result = context
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_6_result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying wglCreateContext expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *WglCreateContextAttribsARB) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := WglCreateContextAttribsARB{}
	context := HGLRC(ϟa.Result)    // HGLRC
	identifier := ϟc.NextContextID // ContextID
	ϟc.NextContextID = (ϟc.NextContextID) + (ContextID(uint32(1)))
	ctx := func() *Context {
		s := &Context{}
		s.Init()
		return s
	}() // ContextPtr
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
	}() // FramebufferPtr
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
	CreateContext_7_result := ctx // ContextPtr
	ϟc.WGLContexts[context] = CreateContext_7_result
	ϟa.Result = context
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_7_result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying wglCreateContextAttribsARB expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *WglMakeCurrent) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := WglMakeCurrent{}
	SetContext_8_context := ϟc.WGLContexts.Get(ϟa.Hglrc) // ContextPtr
	ϟc.Contexts[ϟc.CurrentThread] = SetContext_8_context
	ϟa.Result = ϟa.Result
	_ = SetContext_8_context
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying wglMakeCurrent expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *WglSwapBuffers) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := WglSwapBuffers{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying wglSwapBuffers expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *CGLCreateContext) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := CGLCreateContext{}
	context := CGLContextObj(ϟa.Ctx) // CGLContextObj
	identifier := ϟc.NextContextID   // ContextID
	ϟc.NextContextID = (ϟc.NextContextID) + (ContextID(uint32(1)))
	ctx := func() *Context {
		s := &Context{}
		s.Init()
		return s
	}() // ContextPtr
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
	}() // FramebufferPtr
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
	CreateContext_9_result := ctx // ContextPtr
	ϟc.CGLContexts[context] = CreateContext_9_result
	ϟa.Ctx = context
	ϟa.Result = ϟa.Result
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_9_result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying CGLCreateContext expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *CGLSetCurrentContext) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := CGLSetCurrentContext{}
	SetContext_10_context := ϟc.CGLContexts.Get(ϟa.Ctx) // ContextPtr
	ϟc.Contexts[ϟc.CurrentThread] = SetContext_10_context
	ϟa.Result = ϟa.Result
	_ = SetContext_10_context
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying CGLSetCurrentContext expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlEnableClientState) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlEnableClientState{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_11_result := context              // ContextPtr
	ctx := GetContext_11_result                  // ContextPtr
	ctx.Capabilities[Capability(ϟa.Type)] = true
	_, _, _ = context, GetContext_11_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glEnableClientState expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDisableClientState) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlDisableClientState{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_12_result := context              // ContextPtr
	ctx := GetContext_12_result                  // ContextPtr
	ctx.Capabilities[Capability(ϟa.Type)] = false
	_, _, _ = context, GetContext_12_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDisableClientState expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetProgramBinaryOES) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGetProgramBinaryOES{}
	ϟa.BytesWritten = ϟa.BytesWritten
	ϟa.BinaryFormat = ϟa.BinaryFormat
	ϟa.Binary = ϟa.Binary
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetProgramBinaryOES expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlProgramBinaryOES) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlProgramBinaryOES{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glProgramBinaryOES expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlStartTilingQCOM) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlStartTilingQCOM{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glStartTilingQCOM expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlEndTilingQCOM) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlEndTilingQCOM{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glEndTilingQCOM expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDiscardFramebufferEXT) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlDiscardFramebufferEXT{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDiscardFramebufferEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlInsertEventMarkerEXT) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlInsertEventMarkerEXT{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glInsertEventMarkerEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlPushGroupMarkerEXT) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlPushGroupMarkerEXT{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glPushGroupMarkerEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlPopGroupMarkerEXT) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlPopGroupMarkerEXT{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glPopGroupMarkerEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlTexStorage1DEXT) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlTexStorage1DEXT{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glTexStorage1DEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlTexStorage2DEXT) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlTexStorage2DEXT{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glTexStorage2DEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlTexStorage3DEXT) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlTexStorage3DEXT{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glTexStorage3DEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlTextureStorage1DEXT) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlTextureStorage1DEXT{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glTextureStorage1DEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlTextureStorage2DEXT) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlTextureStorage2DEXT{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glTextureStorage2DEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlTextureStorage3DEXT) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlTextureStorage3DEXT{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glTextureStorage3DEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGenVertexArraysOES) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGenVertexArraysOES{}
	ϟo.Arrays = make(VertexArrayIdArray, ϟa.Count)
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_13_result := context              // ContextPtr
	ctx := GetContext_13_result                  // ContextPtr
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		id := VertexArrayId(ϟa.Arrays[i]) // VertexArrayId
		ctx.Instances.VertexArrays[id] = func() *VertexArray {
			s := &VertexArray{}
			s.Init()
			return s
		}()
		ϟa.Arrays[i] = id
		_ = id
	}
	_, _, _ = context, GetContext_13_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGenVertexArraysOES expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlBindVertexArrayOES) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlBindVertexArrayOES{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_14_result := context              // ContextPtr
	ctx := GetContext_14_result                  // ContextPtr
	if !(ctx.Instances.VertexArrays.Contains(ϟa.Array)) {
		ctx.Instances.VertexArrays[ϟa.Array] = func() *VertexArray {
			s := &VertexArray{}
			s.Init()
			return s
		}()
	}
	ctx.BoundVertexArray = ϟa.Array
	_, _, _ = context, GetContext_14_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glBindVertexArrayOES expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDeleteVertexArraysOES) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlDeleteVertexArraysOES{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_15_result := context              // ContextPtr
	ctx := GetContext_15_result                  // ContextPtr
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		ctx.Instances.VertexArrays.Delete(ϟa.Arrays[i])
	}
	_, _, _ = context, GetContext_15_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDeleteVertexArraysOES expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlIsVertexArrayOES) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlIsVertexArrayOES{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_16_result := context              // ContextPtr
	ctx := GetContext_16_result                  // ContextPtr
	ϟa.Result = ctx.Instances.VertexArrays.Contains(ϟa.Array)
	_, _, _ = context, GetContext_16_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glIsVertexArrayOES expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlEGLImageTargetTexture2DOES) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlEGLImageTargetTexture2DOES{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glEGLImageTargetTexture2DOES expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlEGLImageTargetRenderbufferStorageOES) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlEGLImageTargetRenderbufferStorageOES{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glEGLImageTargetRenderbufferStorageOES expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetGraphicsResetStatusEXT) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGetGraphicsResetStatusEXT{}
	ϟa.Result = ϟa.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetGraphicsResetStatusEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlBindAttribLocation) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlBindAttribLocation{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_17_result := context              // ContextPtr
	ctx := GetContext_17_result                  // ContextPtr
	p := ctx.Instances.Programs.Get(ϟa.Program)  // ProgramPtr
	p.AttributeBindings[ϟa.Name] = ϟa.Location
	_, _, _, _ = context, GetContext_17_result, ctx, p
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glBindAttribLocation expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlBlendFunc) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlBlendFunc{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_18_result := context              // ContextPtr
	ctx := GetContext_18_result                  // ContextPtr
	ctx.Blending.SrcRgbBlendFactor = ϟa.SrcFactor
	ctx.Blending.SrcAlphaBlendFactor = ϟa.SrcFactor
	ctx.Blending.DstRgbBlendFactor = ϟa.DstFactor
	ctx.Blending.DstAlphaBlendFactor = ϟa.DstFactor
	_, _, _ = context, GetContext_18_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glBlendFunc expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlBlendFuncSeparate) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlBlendFuncSeparate{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_19_result := context              // ContextPtr
	ctx := GetContext_19_result                  // ContextPtr
	ctx.Blending.SrcRgbBlendFactor = ϟa.SrcFactorRgb
	ctx.Blending.DstRgbBlendFactor = ϟa.DstFactorRgb
	ctx.Blending.SrcAlphaBlendFactor = ϟa.SrcFactorAlpha
	ctx.Blending.DstAlphaBlendFactor = ϟa.DstFactorAlpha
	_, _, _ = context, GetContext_19_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glBlendFuncSeparate expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlBlendEquation) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlBlendEquation{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_20_result := context              // ContextPtr
	ctx := GetContext_20_result                  // ContextPtr
	ctx.Blending.BlendEquationRgb = ϟa.Equation
	ctx.Blending.BlendEquationAlpha = ϟa.Equation
	_, _, _ = context, GetContext_20_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glBlendEquation expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlBlendEquationSeparate) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlBlendEquationSeparate{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_21_result := context              // ContextPtr
	ctx := GetContext_21_result                  // ContextPtr
	ctx.Blending.BlendEquationRgb = ϟa.Rgb
	ctx.Blending.BlendEquationAlpha = ϟa.Alpha
	_, _, _ = context, GetContext_21_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glBlendEquationSeparate expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlBlendColor) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlBlendColor{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_22_result := context              // ContextPtr
	ctx := GetContext_22_result                  // ContextPtr
	ctx.Blending.BlendColor = func() Color {
		s := Color{}
		s.Init()
		s.Red = ϟa.Red
		s.Green = ϟa.Green
		s.Blue = ϟa.Blue
		s.Alpha = ϟa.Alpha
		return s
	}()
	_, _, _ = context, GetContext_22_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glBlendColor expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlEnableVertexAttribArray) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlEnableVertexAttribArray{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_23_result := context              // ContextPtr
	ctx := GetContext_23_result                  // ContextPtr
	ctx.VertexAttributeArrays.Get(ϟa.Location).Enabled = true
	_, _, _ = context, GetContext_23_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glEnableVertexAttribArray expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDisableVertexAttribArray) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlDisableVertexAttribArray{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_24_result := context              // ContextPtr
	ctx := GetContext_24_result                  // ContextPtr
	ctx.VertexAttributeArrays.Get(ϟa.Location).Enabled = false
	_, _, _ = context, GetContext_24_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDisableVertexAttribArray expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlVertexAttribPointer) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlVertexAttribPointer{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)    // ContextPtr
	GetContext_25_result := context                 // ContextPtr
	ctx := GetContext_25_result                     // ContextPtr
	a := ctx.VertexAttributeArrays.Get(ϟa.Location) // VertexAttributeArrayPtr
	a.Size = uint32(ϟa.Size)
	a.Type = ϟa.Type
	a.Normalized = ϟa.Normalized
	a.Stride = ϟa.Stride
	a.Pointer = memory.Pointer(ϟa.Data)
	a.Buffer = ctx.BoundBuffers.Get(BufferTarget_GL_ARRAY_BUFFER)
	_, _, _, _ = context, GetContext_25_result, ctx, a
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glVertexAttribPointer expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetActiveAttrib) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGetActiveAttrib{}
	ϟa.BufferBytesWritten = ϟa.BufferBytesWritten
	ϟa.VectorCount = ϟa.VectorCount
	ϟa.Type = ϟa.Type
	ϟa.Name = ϟa.Name
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetActiveAttrib expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetActiveUniform) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGetActiveUniform{}
	ϟa.BufferBytesWritten = ϟa.BufferBytesWritten
	ϟa.Size = ϟa.Size
	ϟa.Type = ϟa.Type
	ϟa.Name = ϟa.Name
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetActiveUniform expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetError) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGetError{}
	ϟa.Result = ϟa.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetError expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetProgramiv) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGetProgramiv{}
	ϟo.Value = make(S32Array, int32(1))
	ϟa.Value = ϟa.Value
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetProgramiv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetShaderiv) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGetShaderiv{}
	ϟo.Value = make(S32Array, int32(1))
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_26_result := context              // ContextPtr
	ctx := GetContext_26_result                  // ContextPtr
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // ShaderPtr
	ϟa.Value[int32(0)] = func() (result int32) {
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
			return strlen(s.InfoLog)
		case ShaderParameter_GL_SHADER_SOURCE_LENGTH:
			return strlen(s.Source[int32(0)])
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.Parameter, ϟa))
			return result
		}
	}()
	_, _, _, _ = context, GetContext_26_result, ctx, s
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetShaderiv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetUniformLocation) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGetUniformLocation{}
	ϟa.Result = ϟa.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetUniformLocation expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetAttribLocation) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGetAttribLocation{}
	ϟa.Result = ϟa.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetAttribLocation expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlPixelStorei) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlPixelStorei{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_27_result := context              // ContextPtr
	ctx := GetContext_27_result                  // ContextPtr
	ctx.PixelStorage[ϟa.Parameter] = ϟa.Value
	_, _, _ = context, GetContext_27_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glPixelStorei expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlTexParameteri) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlTexParameteri{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                     // ContextPtr
	GetContext_28_result := context                                  // ContextPtr
	ctx := GetContext_28_result                                      // ContextPtr
	id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(ϟa.Target) // TextureId
	t := ctx.Instances.Textures.Get(id)                              // TexturePtr
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
		// TODO: better unmatched handling
		v := ϟa.Parameter
		log.Printf("Error: Missing switch case handler for value %T %v", v, v)
	}
	_, _, _, _, _ = context, GetContext_28_result, ctx, id, t
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glTexParameteri expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlTexParameterf) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlTexParameterf{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                     // ContextPtr
	GetContext_29_result := context                                  // ContextPtr
	ctx := GetContext_29_result                                      // ContextPtr
	id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(ϟa.Target) // TextureId
	t := ctx.Instances.Textures.Get(id)                              // TexturePtr
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
		// TODO: better unmatched handling
		v := ϟa.Parameter
		log.Printf("Error: Missing switch case handler for value %T %v", v, v)
	}
	_, _, _, _, _ = context, GetContext_29_result, ctx, id, t
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glTexParameterf expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetTexParameteriv) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGetTexParameteriv{}
	ϟo.Values = make(S32Array, int32(1))
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                     // ContextPtr
	GetContext_30_result := context                                  // ContextPtr
	ctx := GetContext_30_result                                      // ContextPtr
	id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(ϟa.Target) // TextureId
	t := ctx.Instances.Textures.Get(id)                              // TexturePtr
	ϟa.Values[int32(0)] = func() (result int32) {
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
	}()
	_, _, _, _, _ = context, GetContext_30_result, ctx, id, t
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetTexParameteriv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetTexParameterfv) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGetTexParameterfv{}
	ϟo.Values = make(F32Array, int32(1))
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                     // ContextPtr
	GetContext_31_result := context                                  // ContextPtr
	ctx := GetContext_31_result                                      // ContextPtr
	id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(ϟa.Target) // TextureId
	t := ctx.Instances.Textures.Get(id)                              // TexturePtr
	ϟa.Values[int32(0)] = func() (result float32) {
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
	}()
	_, _, _, _, _ = context, GetContext_31_result, ctx, id, t
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetTexParameterfv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniform1i) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlUniform1i{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)            // ContextPtr
	GetContext_32_result := context                         // ContextPtr
	ctx := GetContext_32_result                             // ContextPtr
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // ProgramPtr
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = ShaderUniformType_GL_INT
	uniform.Value.S32 = ϟa.Value
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _ = context, GetContext_32_result, ctx, program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniform1i expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniform2i) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlUniform2i{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)            // ContextPtr
	GetContext_33_result := context                         // ContextPtr
	ctx := GetContext_33_result                             // ContextPtr
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // ProgramPtr
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
	_, _, _, _, _ = context, GetContext_33_result, ctx, program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniform2i expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniform3i) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlUniform3i{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)            // ContextPtr
	GetContext_34_result := context                         // ContextPtr
	ctx := GetContext_34_result                             // ContextPtr
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // ProgramPtr
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
	_, _, _, _, _ = context, GetContext_34_result, ctx, program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniform3i expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniform4i) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlUniform4i{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)            // ContextPtr
	GetContext_35_result := context                         // ContextPtr
	ctx := GetContext_35_result                             // ContextPtr
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // ProgramPtr
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
	_, _, _, _, _ = context, GetContext_35_result, ctx, program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniform4i expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniform1iv) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlUniform1iv{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)            // ContextPtr
	GetContext_36_result := context                         // ContextPtr
	ctx := GetContext_36_result                             // ContextPtr
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // ProgramPtr
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = ShaderUniformType_GL_INT
	uniform.Value.S32 = ϟa.Value[int32(0)]
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _ = context, GetContext_36_result, ctx, program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniform1iv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniform2iv) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlUniform2iv{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)            // ContextPtr
	GetContext_37_result := context                         // ContextPtr
	ctx := GetContext_37_result                             // ContextPtr
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // ProgramPtr
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = ShaderUniformType_GL_INT_VEC2
	uniform.Value.Vec2i = func() Vec2i {
		s := Vec2i{}
		s.Init()
		s.X = ϟa.Value[int32(0)]
		s.Y = ϟa.Value[int32(1)]
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _ = context, GetContext_37_result, ctx, program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniform2iv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniform3iv) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlUniform3iv{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)            // ContextPtr
	GetContext_38_result := context                         // ContextPtr
	ctx := GetContext_38_result                             // ContextPtr
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // ProgramPtr
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = ShaderUniformType_GL_INT_VEC3
	uniform.Value.Vec3i = func() Vec3i {
		s := Vec3i{}
		s.Init()
		s.X = ϟa.Value[int32(0)]
		s.Y = ϟa.Value[int32(1)]
		s.Z = ϟa.Value[int32(2)]
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _ = context, GetContext_38_result, ctx, program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniform3iv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniform4iv) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlUniform4iv{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)            // ContextPtr
	GetContext_39_result := context                         // ContextPtr
	ctx := GetContext_39_result                             // ContextPtr
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // ProgramPtr
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = ShaderUniformType_GL_INT_VEC4
	uniform.Value.Vec4i = func() Vec4i {
		s := Vec4i{}
		s.Init()
		s.X = ϟa.Value[int32(0)]
		s.Y = ϟa.Value[int32(1)]
		s.Z = ϟa.Value[int32(2)]
		s.W = ϟa.Value[int32(3)]
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _ = context, GetContext_39_result, ctx, program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniform4iv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniform1f) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlUniform1f{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)            // ContextPtr
	GetContext_40_result := context                         // ContextPtr
	ctx := GetContext_40_result                             // ContextPtr
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // ProgramPtr
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT
	uniform.Value.F32 = ϟa.Value
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _ = context, GetContext_40_result, ctx, program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniform1f expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniform2f) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlUniform2f{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)            // ContextPtr
	GetContext_41_result := context                         // ContextPtr
	ctx := GetContext_41_result                             // ContextPtr
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // ProgramPtr
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
	_, _, _, _, _ = context, GetContext_41_result, ctx, program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniform2f expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniform3f) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlUniform3f{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)            // ContextPtr
	GetContext_42_result := context                         // ContextPtr
	ctx := GetContext_42_result                             // ContextPtr
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // ProgramPtr
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
	_, _, _, _, _ = context, GetContext_42_result, ctx, program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniform3f expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniform4f) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlUniform4f{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)            // ContextPtr
	GetContext_43_result := context                         // ContextPtr
	ctx := GetContext_43_result                             // ContextPtr
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // ProgramPtr
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
	_, _, _, _, _ = context, GetContext_43_result, ctx, program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniform4f expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniform1fv) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlUniform1fv{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)            // ContextPtr
	GetContext_44_result := context                         // ContextPtr
	ctx := GetContext_44_result                             // ContextPtr
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // ProgramPtr
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT
	uniform.Value.F32 = ϟa.Value[int32(0)]
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _ = context, GetContext_44_result, ctx, program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniform1fv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniform2fv) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlUniform2fv{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)            // ContextPtr
	GetContext_45_result := context                         // ContextPtr
	ctx := GetContext_45_result                             // ContextPtr
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // ProgramPtr
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT_VEC2
	uniform.Value.Vec2f = func() Vec2f {
		s := Vec2f{}
		s.Init()
		s.X = ϟa.Value[int32(0)]
		s.Y = ϟa.Value[int32(1)]
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _ = context, GetContext_45_result, ctx, program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniform2fv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniform3fv) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlUniform3fv{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)            // ContextPtr
	GetContext_46_result := context                         // ContextPtr
	ctx := GetContext_46_result                             // ContextPtr
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // ProgramPtr
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT_VEC3
	uniform.Value.Vec3f = func() Vec3f {
		s := Vec3f{}
		s.Init()
		s.X = ϟa.Value[int32(0)]
		s.Y = ϟa.Value[int32(1)]
		s.Z = ϟa.Value[int32(2)]
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _ = context, GetContext_46_result, ctx, program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniform3fv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniform4fv) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlUniform4fv{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)            // ContextPtr
	GetContext_47_result := context                         // ContextPtr
	ctx := GetContext_47_result                             // ContextPtr
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // ProgramPtr
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT_VEC4
	uniform.Value.Vec4f = func() Vec4f {
		s := Vec4f{}
		s.Init()
		s.X = ϟa.Value[int32(0)]
		s.Y = ϟa.Value[int32(1)]
		s.Z = ϟa.Value[int32(2)]
		s.W = ϟa.Value[int32(3)]
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _ = context, GetContext_47_result, ctx, program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniform4fv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniformMatrix2fv) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlUniformMatrix2fv{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)            // ContextPtr
	GetContext_48_result := context                         // ContextPtr
	ctx := GetContext_48_result                             // ContextPtr
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // ProgramPtr
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT_MAT2
	uniform.Value.Mat2f = func() Mat2f {
		s := Mat2f{}
		s.Init()
		s.Col0 = func() Vec2f {
			s := Vec2f{}
			s.Init()
			s.X = ϟa.Values[int32(0)]
			s.Y = ϟa.Values[int32(1)]
			return s
		}()
		s.Col1 = func() Vec2f {
			s := Vec2f{}
			s.Init()
			s.X = ϟa.Values[int32(0)]
			s.Y = ϟa.Values[int32(1)]
			return s
		}()
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _ = context, GetContext_48_result, ctx, program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniformMatrix2fv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniformMatrix3fv) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlUniformMatrix3fv{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)            // ContextPtr
	GetContext_49_result := context                         // ContextPtr
	ctx := GetContext_49_result                             // ContextPtr
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // ProgramPtr
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT_MAT3
	uniform.Value.Mat3f = func() Mat3f {
		s := Mat3f{}
		s.Init()
		s.Col0 = func() Vec3f {
			s := Vec3f{}
			s.Init()
			s.X = ϟa.Values[int32(0)]
			s.Y = ϟa.Values[int32(1)]
			s.Z = ϟa.Values[int32(2)]
			return s
		}()
		s.Col1 = func() Vec3f {
			s := Vec3f{}
			s.Init()
			s.X = ϟa.Values[int32(3)]
			s.Y = ϟa.Values[int32(4)]
			s.Z = ϟa.Values[int32(5)]
			return s
		}()
		s.Col2 = func() Vec3f {
			s := Vec3f{}
			s.Init()
			s.X = ϟa.Values[int32(6)]
			s.Y = ϟa.Values[int32(7)]
			s.Z = ϟa.Values[int32(8)]
			return s
		}()
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _ = context, GetContext_49_result, ctx, program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniformMatrix3fv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniformMatrix4fv) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlUniformMatrix4fv{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)            // ContextPtr
	GetContext_50_result := context                         // ContextPtr
	ctx := GetContext_50_result                             // ContextPtr
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // ProgramPtr
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Value.Mat4f = func() Mat4f {
		s := Mat4f{}
		s.Init()
		s.Col0 = func() Vec4f {
			s := Vec4f{}
			s.Init()
			s.X = ϟa.Values[int32(0)]
			s.Y = ϟa.Values[int32(1)]
			s.Z = ϟa.Values[int32(2)]
			s.W = ϟa.Values[int32(3)]
			return s
		}()
		s.Col1 = func() Vec4f {
			s := Vec4f{}
			s.Init()
			s.X = ϟa.Values[int32(4)]
			s.Y = ϟa.Values[int32(5)]
			s.Z = ϟa.Values[int32(6)]
			s.W = ϟa.Values[int32(7)]
			return s
		}()
		s.Col2 = func() Vec4f {
			s := Vec4f{}
			s.Init()
			s.X = ϟa.Values[int32(8)]
			s.Y = ϟa.Values[int32(9)]
			s.Z = ϟa.Values[int32(10)]
			s.W = ϟa.Values[int32(11)]
			return s
		}()
		s.Col3 = func() Vec4f {
			s := Vec4f{}
			s.Init()
			s.X = ϟa.Values[int32(12)]
			s.Y = ϟa.Values[int32(13)]
			s.Z = ϟa.Values[int32(14)]
			s.W = ϟa.Values[int32(15)]
			return s
		}()
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _ = context, GetContext_50_result, ctx, program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniformMatrix4fv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetUniformfv) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGetUniformfv{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetUniformfv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetUniformiv) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGetUniformiv{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetUniformiv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlVertexAttrib1f) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlVertexAttrib1f{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glVertexAttrib1f expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlVertexAttrib2f) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlVertexAttrib2f{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glVertexAttrib2f expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlVertexAttrib3f) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlVertexAttrib3f{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glVertexAttrib3f expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlVertexAttrib4f) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlVertexAttrib4f{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glVertexAttrib4f expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlVertexAttrib1fv) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlVertexAttrib1fv{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glVertexAttrib1fv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlVertexAttrib2fv) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlVertexAttrib2fv{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glVertexAttrib2fv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlVertexAttrib3fv) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlVertexAttrib3fv{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glVertexAttrib3fv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlVertexAttrib4fv) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlVertexAttrib4fv{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glVertexAttrib4fv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetShaderPrecisionFormat) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGetShaderPrecisionFormat{}
	ϟo.Range = make(S32Array, int32(2))
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetShaderPrecisionFormat expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDepthMask) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlDepthMask{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_51_result := context              // ContextPtr
	ctx := GetContext_51_result                  // ContextPtr
	ctx.Rasterizing.DepthMask = ϟa.Enabled
	_, _, _ = context, GetContext_51_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDepthMask expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDepthFunc) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlDepthFunc{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_52_result := context              // ContextPtr
	ctx := GetContext_52_result                  // ContextPtr
	ctx.Rasterizing.DepthTestFunction = ϟa.Function
	_, _, _ = context, GetContext_52_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDepthFunc expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDepthRangef) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlDepthRangef{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_53_result := context              // ContextPtr
	ctx := GetContext_53_result                  // ContextPtr
	ctx.Rasterizing.DepthNear = ϟa.Near
	ctx.Rasterizing.DepthFar = ϟa.Far
	_, _, _ = context, GetContext_53_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDepthRangef expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlColorMask) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlColorMask{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_54_result := context              // ContextPtr
	ctx := GetContext_54_result                  // ContextPtr
	ctx.Rasterizing.ColorMaskRed = ϟa.Red
	ctx.Rasterizing.ColorMaskGreen = ϟa.Green
	ctx.Rasterizing.ColorMaskBlue = ϟa.Blue
	ctx.Rasterizing.ColorMaskAlpha = ϟa.Alpha
	_, _, _ = context, GetContext_54_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glColorMask expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlStencilMask) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlStencilMask{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_55_result := context              // ContextPtr
	ctx := GetContext_55_result                  // ContextPtr
	ctx.Rasterizing.StencilMask[FaceMode_GL_FRONT] = ϟa.Mask
	ctx.Rasterizing.StencilMask[FaceMode_GL_BACK] = ϟa.Mask
	_, _, _ = context, GetContext_55_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glStencilMask expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlStencilMaskSeparate) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlStencilMaskSeparate{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_56_result := context              // ContextPtr
	ctx := GetContext_56_result                  // ContextPtr
	switch ϟa.Face {
	case FaceMode_GL_FRONT:
		ctx.Rasterizing.StencilMask[FaceMode_GL_FRONT] = ϟa.Mask
	case FaceMode_GL_BACK:
		ctx.Rasterizing.StencilMask[FaceMode_GL_BACK] = ϟa.Mask
	case FaceMode_GL_FRONT_AND_BACK:
		ctx.Rasterizing.StencilMask[FaceMode_GL_FRONT] = ϟa.Mask
		ctx.Rasterizing.StencilMask[FaceMode_GL_BACK] = ϟa.Mask
	default:
		// TODO: better unmatched handling
		v := ϟa.Face
		log.Printf("Error: Missing switch case handler for value %T %v", v, v)
	}
	_, _, _ = context, GetContext_56_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glStencilMaskSeparate expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlStencilFuncSeparate) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlStencilFuncSeparate{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glStencilFuncSeparate expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlStencilOpSeparate) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlStencilOpSeparate{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glStencilOpSeparate expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlFrontFace) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlFrontFace{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_57_result := context              // ContextPtr
	ctx := GetContext_57_result                  // ContextPtr
	ctx.Rasterizing.FrontFace = ϟa.Orientation
	_, _, _ = context, GetContext_57_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glFrontFace expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlViewport) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlViewport{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_58_result := context              // ContextPtr
	ctx := GetContext_58_result                  // ContextPtr
	ctx.Rasterizing.Viewport = func() Rect {
		s := Rect{}
		s.Init()
		s.X = ϟa.X
		s.Y = ϟa.Y
		s.Width = ϟa.Width
		s.Height = ϟa.Height
		return s
	}()
	_, _, _ = context, GetContext_58_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glViewport expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlScissor) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlScissor{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_59_result := context              // ContextPtr
	ctx := GetContext_59_result                  // ContextPtr
	ctx.Rasterizing.Scissor = func() Rect {
		s := Rect{}
		s.Init()
		s.X = ϟa.X
		s.Y = ϟa.Y
		s.Width = ϟa.Width
		s.Height = ϟa.Height
		return s
	}()
	_, _, _ = context, GetContext_59_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glScissor expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlActiveTexture) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlActiveTexture{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_60_result := context              // ContextPtr
	ctx := GetContext_60_result                  // ContextPtr
	ctx.ActiveTextureUnit = ϟa.Unit
	if !(ctx.TextureUnits.Contains(ϟa.Unit)) {
		ctx.TextureUnits[ϟa.Unit] = ctx.TextureUnits.Get(ϟa.Unit)
	}
	_, _, _ = context, GetContext_60_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glActiveTexture expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGenTextures) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGenTextures{}
	ϟo.Textures = make(TextureIdArray, ϟa.Count)
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_61_result := context              // ContextPtr
	ctx := GetContext_61_result                  // ContextPtr
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		id := TextureId(ϟa.Textures[i]) // TextureId
		ctx.Instances.Textures[id] = func() *Texture {
			s := &Texture{}
			s.Init()
			return s
		}()
		ϟa.Textures[i] = id
		_ = id
	}
	_, _, _ = context, GetContext_61_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGenTextures expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDeleteTextures) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlDeleteTextures{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_62_result := context              // ContextPtr
	ctx := GetContext_62_result                  // ContextPtr
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		ctx.Instances.Textures.Delete(ϟa.Textures[i])
	}
	_, _, _ = context, GetContext_62_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDeleteTextures expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlIsTexture) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlIsTexture{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_63_result := context              // ContextPtr
	ctx := GetContext_63_result                  // ContextPtr
	ϟa.Result = ctx.Instances.Textures.Contains(ϟa.Texture)
	_, _, _ = context, GetContext_63_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glIsTexture expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlBindTexture) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlBindTexture{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_64_result := context              // ContextPtr
	ctx := GetContext_64_result                  // ContextPtr
	if !(ctx.Instances.Textures.Contains(ϟa.Texture)) {
		ctx.Instances.Textures[ϟa.Texture] = func() *Texture {
			s := &Texture{}
			s.Init()
			return s
		}()
	}
	ctx.TextureUnits.Get(ctx.ActiveTextureUnit)[ϟa.Target] = ϟa.Texture
	_, _, _ = context, GetContext_64_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glBindTexture expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlTexImage2D) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlTexImage2D{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_65_result := context              // ContextPtr
	ctx := GetContext_65_result                  // ContextPtr
	switch ϟa.Target {
	case TextureImageTarget_GL_TEXTURE_2D:
		id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_2D) // TextureId
		t := ctx.Instances.Textures.Get(id)                                                // TexturePtr
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.Width
			s.Height = ϟa.Height
			s.Size = imageSize(uint32(ϟa.Width), uint32(ϟa.Height), ϟa.Format, ϟa.Type)
			s.Format = ImageTexelFormat(ϟa.Format)
			return s
		}() // Image
		if ((ctx.BoundBuffers.Get(BufferTarget_GL_PIXEL_UNPACK_BUFFER)) == (BufferId(uint32(0)))) && ((ϟa.Data) != (TexturePointer(memory.Pointer(0)))) {
			l.Data.Write(ϟs.Memory.Slice(memory.Range{
				Base: memory.Pointer(ϟa.Data),
				Size: uint64(l.Size),
			}))
		}
		t.Texture2D[ϟa.Level] = l
		t.Kind = TextureKind_TEXTURE2D
		t.Format = ImageTexelFormat(ϟa.Format)
		_, _, _ = id, t, l
	case TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X, TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Z, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z:
		id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_CUBE_MAP) // TextureId
		t := ctx.Instances.Textures.Get(id)                                                      // TexturePtr
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.Width
			s.Height = ϟa.Height
			s.Size = imageSize(uint32(ϟa.Width), uint32(ϟa.Height), ϟa.Format, ϟa.Type)
			s.Format = ImageTexelFormat(ϟa.Format)
			return s
		}() // Image
		if ((ctx.BoundBuffers.Get(BufferTarget_GL_PIXEL_UNPACK_BUFFER)) == (BufferId(uint32(0)))) && ((ϟa.Data) != (TexturePointer(memory.Pointer(0)))) {
			l.Data.Write(ϟs.Memory.Slice(memory.Range{
				Base: memory.Pointer(ϟa.Data),
				Size: uint64(l.Size),
			}))
		}
		cube := t.Cubemap.Get(ϟa.Level) // CubemapLevel
		cube.Faces[CubeMapImageTarget(ϟa.Target)] = l
		t.Cubemap[ϟa.Level] = cube
		t.Kind = TextureKind_CUBEMAP
		t.Format = ImageTexelFormat(ϟa.Format)
		_, _, _, _ = id, t, l, cube
	default:
		// TODO: better unmatched handling
		v := ϟa.Target
		log.Printf("Error: Missing switch case handler for value %T %v", v, v)
	}
	_, _, _ = context, GetContext_65_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glTexImage2D expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlTexSubImage2D) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlTexSubImage2D{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_66_result := context              // ContextPtr
	ctx := GetContext_66_result                  // ContextPtr
	switch ϟa.Target {
	case TextureImageTarget_GL_TEXTURE_2D:
		id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_2D) // TextureId
		t := ctx.Instances.Textures.Get(id)                                                // TexturePtr
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.Width
			s.Height = ϟa.Height
			s.Size = imageSize(uint32(ϟa.Width), uint32(ϟa.Height), ϟa.Format, ϟa.Type)
			s.Format = ImageTexelFormat(ϟa.Format)
			return s
		}() // Image
		if ((ctx.BoundBuffers.Get(BufferTarget_GL_PIXEL_UNPACK_BUFFER)) == (BufferId(uint32(0)))) && ((ϟa.Data) != (TexturePointer(memory.Pointer(0)))) {
			l.Data.Write(ϟs.Memory.Slice(memory.Range{
				Base: memory.Pointer(ϟa.Data),
				Size: uint64(l.Size),
			}))
		}
		t.Texture2D[ϟa.Level] = l
		t.Kind = TextureKind_TEXTURE2D
		t.Format = ImageTexelFormat(ϟa.Format)
		_, _, _ = id, t, l
	case TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X, TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Z, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z:
		id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_CUBE_MAP) // TextureId
		t := ctx.Instances.Textures.Get(id)                                                      // TexturePtr
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.Width
			s.Height = ϟa.Height
			s.Size = imageSize(uint32(ϟa.Width), uint32(ϟa.Height), ϟa.Format, ϟa.Type)
			s.Format = ImageTexelFormat(ϟa.Format)
			return s
		}() // Image
		if ((ctx.BoundBuffers.Get(BufferTarget_GL_PIXEL_UNPACK_BUFFER)) == (BufferId(uint32(0)))) && ((ϟa.Data) != (TexturePointer(memory.Pointer(0)))) {
			l.Data.Write(ϟs.Memory.Slice(memory.Range{
				Base: memory.Pointer(ϟa.Data),
				Size: uint64(l.Size),
			}))
		}
		cube := t.Cubemap.Get(ϟa.Level) // CubemapLevel
		cube.Faces[CubeMapImageTarget(ϟa.Target)] = l
		t.Cubemap[ϟa.Level] = cube
		t.Kind = TextureKind_CUBEMAP
		t.Format = ImageTexelFormat(ϟa.Format)
		_, _, _, _ = id, t, l, cube
	default:
		// TODO: better unmatched handling
		v := ϟa.Target
		log.Printf("Error: Missing switch case handler for value %T %v", v, v)
	}
	_, _, _ = context, GetContext_66_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glTexSubImage2D expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlCopyTexImage2D) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlCopyTexImage2D{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glCopyTexImage2D expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlCopyTexSubImage2D) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlCopyTexSubImage2D{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glCopyTexSubImage2D expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlCompressedTexImage2D) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlCompressedTexImage2D{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_67_result := context              // ContextPtr
	ctx := GetContext_67_result                  // ContextPtr
	switch ϟa.Target {
	case TextureImageTarget_GL_TEXTURE_2D:
		id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_2D) // TextureId
		t := ctx.Instances.Textures.Get(id)                                                // TexturePtr
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.Width
			s.Height = ϟa.Height
			s.Size = uint32(ϟa.ImageSize)
			s.Format = ImageTexelFormat(ϟa.Format)
			return s
		}() // Image
		if ((ctx.BoundBuffers.Get(BufferTarget_GL_PIXEL_UNPACK_BUFFER)) == (BufferId(uint32(0)))) && ((ϟa.Data) != (TexturePointer(memory.Pointer(0)))) {
			l.Data.Write(ϟs.Memory.Slice(memory.Range{
				Base: memory.Pointer(ϟa.Data),
				Size: uint64(l.Size),
			}))
		}
		t.Texture2D[ϟa.Level] = l
		t.Kind = TextureKind_TEXTURE2D
		t.Format = ImageTexelFormat(ϟa.Format)
		_, _, _ = id, t, l
	case TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X, TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Z, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z:
		id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_CUBE_MAP) // TextureId
		t := ctx.Instances.Textures.Get(id)                                                      // TexturePtr
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.Width
			s.Height = ϟa.Height
			s.Size = uint32(ϟa.ImageSize)
			s.Format = ImageTexelFormat(ϟa.Format)
			return s
		}() // Image
		if ((ctx.BoundBuffers.Get(BufferTarget_GL_PIXEL_UNPACK_BUFFER)) == (BufferId(uint32(0)))) && ((ϟa.Data) != (TexturePointer(memory.Pointer(0)))) {
			l.Data.Write(ϟs.Memory.Slice(memory.Range{
				Base: memory.Pointer(ϟa.Data),
				Size: uint64(l.Size),
			}))
		}
		cube := t.Cubemap.Get(ϟa.Level) // CubemapLevel
		cube.Faces[CubeMapImageTarget(ϟa.Target)] = l
		t.Cubemap[ϟa.Level] = cube
		t.Kind = TextureKind_CUBEMAP
		t.Format = ImageTexelFormat(ϟa.Format)
		_, _, _, _ = id, t, l, cube
	default:
		// TODO: better unmatched handling
		v := ϟa.Target
		log.Printf("Error: Missing switch case handler for value %T %v", v, v)
	}
	_, _, _ = context, GetContext_67_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glCompressedTexImage2D expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlCompressedTexSubImage2D) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlCompressedTexSubImage2D{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glCompressedTexSubImage2D expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGenerateMipmap) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGenerateMipmap{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGenerateMipmap expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlReadPixels) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlReadPixels{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glReadPixels expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGenFramebuffers) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGenFramebuffers{}
	ϟo.Framebuffers = make(FramebufferIdArray, ϟa.Count)
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_68_result := context              // ContextPtr
	ctx := GetContext_68_result                  // ContextPtr
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		id := FramebufferId(ϟa.Framebuffers[i]) // FramebufferId
		ctx.Instances.Framebuffers[id] = func() *Framebuffer {
			s := &Framebuffer{}
			s.Init()
			return s
		}()
		ϟa.Framebuffers[i] = id
		_ = id
	}
	_, _, _ = context, GetContext_68_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGenFramebuffers expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlBindFramebuffer) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlBindFramebuffer{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_69_result := context              // ContextPtr
	ctx := GetContext_69_result                  // ContextPtr
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
	_, _, _ = context, GetContext_69_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glBindFramebuffer expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlCheckFramebufferStatus) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlCheckFramebufferStatus{}
	ϟa.Result = ϟa.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glCheckFramebufferStatus expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDeleteFramebuffers) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlDeleteFramebuffers{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_70_result := context              // ContextPtr
	ctx := GetContext_70_result                  // ContextPtr
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		ctx.Instances.Framebuffers.Delete(ϟa.Framebuffers[i])
	}
	_, _, _ = context, GetContext_70_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDeleteFramebuffers expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlIsFramebuffer) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlIsFramebuffer{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_71_result := context              // ContextPtr
	ctx := GetContext_71_result                  // ContextPtr
	ϟa.Result = ctx.Instances.Framebuffers.Contains(ϟa.Framebuffer)
	_, _, _ = context, GetContext_71_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glIsFramebuffer expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGenRenderbuffers) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGenRenderbuffers{}
	ϟo.Renderbuffers = make(RenderbufferIdArray, ϟa.Count)
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_72_result := context              // ContextPtr
	ctx := GetContext_72_result                  // ContextPtr
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		id := RenderbufferId(ϟa.Renderbuffers[i]) // RenderbufferId
		ctx.Instances.Renderbuffers[id] = func() *Renderbuffer {
			s := &Renderbuffer{}
			s.Init()
			return s
		}()
		ϟa.Renderbuffers[i] = id
		_ = id
	}
	_, _, _ = context, GetContext_72_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGenRenderbuffers expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlBindRenderbuffer) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlBindRenderbuffer{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_73_result := context              // ContextPtr
	ctx := GetContext_73_result                  // ContextPtr
	if !(ctx.Instances.Renderbuffers.Contains(ϟa.Renderbuffer)) {
		ctx.Instances.Renderbuffers[ϟa.Renderbuffer] = func() *Renderbuffer {
			s := &Renderbuffer{}
			s.Init()
			return s
		}()
	}
	ctx.BoundRenderbuffers[ϟa.Target] = ϟa.Renderbuffer
	_, _, _ = context, GetContext_73_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glBindRenderbuffer expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlRenderbufferStorage) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlRenderbufferStorage{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_74_result := context              // ContextPtr
	ctx := GetContext_74_result                  // ContextPtr
	id := ctx.BoundRenderbuffers.Get(ϟa.Target)  // RenderbufferId
	rb := ctx.Instances.Renderbuffers.Get(id)    // RenderbufferPtr
	rb.Format = ϟa.Format
	rb.Width = ϟa.Width
	rb.Height = ϟa.Height
	_, _, _, _, _ = context, GetContext_74_result, ctx, id, rb
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glRenderbufferStorage expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDeleteRenderbuffers) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlDeleteRenderbuffers{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_75_result := context              // ContextPtr
	ctx := GetContext_75_result                  // ContextPtr
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		ctx.Instances.Renderbuffers.Delete(ϟa.Renderbuffers[i])
	}
	_, _, _ = context, GetContext_75_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDeleteRenderbuffers expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlIsRenderbuffer) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlIsRenderbuffer{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_76_result := context              // ContextPtr
	ctx := GetContext_76_result                  // ContextPtr
	ϟa.Result = ctx.Instances.Renderbuffers.Contains(ϟa.Renderbuffer)
	_, _, _ = context, GetContext_76_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glIsRenderbuffer expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetRenderbufferParameteriv) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGetRenderbufferParameteriv{}
	ϟo.Values = make(S32Array, int32(1))
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_77_result := context              // ContextPtr
	ctx := GetContext_77_result                  // ContextPtr
	id := ctx.BoundRenderbuffers.Get(ϟa.Target)  // RenderbufferId
	rb := ctx.Instances.Renderbuffers.Get(id)    // RenderbufferPtr
	ϟa.Values[int32(0)] = func() (result int32) {
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
	}()
	_, _, _, _, _ = context, GetContext_77_result, ctx, id, rb
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetRenderbufferParameteriv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGenBuffers) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGenBuffers{}
	ϟo.Buffers = make(BufferIdArray, ϟa.Count)
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_78_result := context              // ContextPtr
	ctx := GetContext_78_result                  // ContextPtr
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		id := BufferId(ϟa.Buffers[i]) // BufferId
		ctx.Instances.Buffers[id] = func() *Buffer {
			s := &Buffer{}
			s.Init()
			return s
		}()
		ϟa.Buffers[i] = id
		_ = id
	}
	_, _, _ = context, GetContext_78_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGenBuffers expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlBindBuffer) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlBindBuffer{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_79_result := context              // ContextPtr
	ctx := GetContext_79_result                  // ContextPtr
	if !(ctx.Instances.Buffers.Contains(ϟa.Buffer)) {
		ctx.Instances.Buffers[ϟa.Buffer] = func() *Buffer {
			s := &Buffer{}
			s.Init()
			return s
		}()
	}
	ctx.BoundBuffers[ϟa.Target] = ϟa.Buffer
	_, _, _ = context, GetContext_79_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glBindBuffer expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlBufferData) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlBufferData{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_80_result := context              // ContextPtr
	ctx := GetContext_80_result                  // ContextPtr
	id := ctx.BoundBuffers.Get(ϟa.Target)        // BufferId
	b := ctx.Instances.Buffers.Get(id)           // BufferPtr
	if (ϟa.Data) != (BufferDataPointer(memory.Pointer(0))) {
		b.Data.Write(ϟs.Memory.Slice(memory.Range{
			Base: memory.Pointer(ϟa.Data),
			Size: uint64(uint32(ϟa.Size)),
		}))
	}
	b.Size = ϟa.Size
	b.Usage = ϟa.Usage
	_, _, _, _, _ = context, GetContext_80_result, ctx, id, b
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glBufferData expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlBufferSubData) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlBufferSubData{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glBufferSubData expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDeleteBuffers) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlDeleteBuffers{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_81_result := context              // ContextPtr
	ctx := GetContext_81_result                  // ContextPtr
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		ctx.Instances.Buffers.Delete(ϟa.Buffers[i])
	}
	_, _, _ = context, GetContext_81_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDeleteBuffers expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlIsBuffer) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlIsBuffer{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_82_result := context              // ContextPtr
	ctx := GetContext_82_result                  // ContextPtr
	ϟa.Result = ctx.Instances.Buffers.Contains(ϟa.Buffer)
	_, _, _ = context, GetContext_82_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glIsBuffer expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetBufferParameteriv) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGetBufferParameteriv{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_83_result := context              // ContextPtr
	ctx := GetContext_83_result                  // ContextPtr
	id := ctx.BoundBuffers.Get(ϟa.Target)        // BufferId
	b := ctx.Instances.Buffers.Get(id)           // BufferPtr
	ϟa.Value = func() (result int32) {
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
	}()
	_, _, _, _, _ = context, GetContext_83_result, ctx, id, b
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetBufferParameteriv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlCreateShader) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlCreateShader{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_84_result := context              // ContextPtr
	ctx := GetContext_84_result                  // ContextPtr
	id := ShaderId(ϟa.Result)                    // ShaderId
	ctx.Instances.Shaders[id] = func() *Shader {
		s := &Shader{}
		s.Init()
		return s
	}()
	s := ctx.Instances.Shaders.Get(id) // ShaderPtr
	s.Type = ϟa.Type
	ϟa.Result = id
	_, _, _, _, _ = context, GetContext_84_result, ctx, id, s
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glCreateShader expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDeleteShader) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlDeleteShader{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_85_result := context              // ContextPtr
	ctx := GetContext_85_result                  // ContextPtr
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // ShaderPtr
	s.Deletable = true
	ctx.Instances.Shaders.Delete(ϟa.Shader)
	_, _, _, _ = context, GetContext_85_result, ctx, s
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDeleteShader expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlShaderSource) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlShaderSource{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_86_result := context              // ContextPtr
	ctx := GetContext_86_result                  // ContextPtr
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // ShaderPtr
	s.Source = ϟa.Source
	_, _, _, _ = context, GetContext_86_result, ctx, s
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glShaderSource expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlShaderBinary) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlShaderBinary{}
	for i := int32(int32(0)); i < ϟa.Count; i++ {
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glShaderBinary expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetShaderInfoLog) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGetShaderInfoLog{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_87_result := context              // ContextPtr
	ctx := GetContext_87_result                  // ContextPtr
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // ShaderPtr
	min_88_a := ϟa.BufferLength                  // s32
	min_88_b := strlen(s.InfoLog)                // s32
	min_88_result := func() (result int32) {
		switch (min_88_a) < (min_88_b) {
		case true:
			return min_88_a
		case false:
			return min_88_b
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (min_88_a) < (min_88_b), ϟa))
			return result
		}
	}() // s32
	ϟa.StringLengthWritten = min_88_result
	ϟa.Info = substr(s.InfoLog, int32(0), ϟa.StringLengthWritten)
	_, _, _, _, _, _, _ = context, GetContext_87_result, ctx, s, min_88_a, min_88_b, min_88_result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetShaderInfoLog expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetShaderSource) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGetShaderSource{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_89_result := context              // ContextPtr
	ctx := GetContext_89_result                  // ContextPtr
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // ShaderPtr
	min_90_a := ϟa.BufferLength                  // s32
	min_90_b := strlen(s.Source[int32(0)])       // s32
	min_90_result := func() (result int32) {
		switch (min_90_a) < (min_90_b) {
		case true:
			return min_90_a
		case false:
			return min_90_b
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (min_90_a) < (min_90_b), ϟa))
			return result
		}
	}() // s32
	ϟa.StringLengthWritten = min_90_result
	ϟa.Source = substr(s.Source[int32(0)], int32(0), ϟa.StringLengthWritten)
	_, _, _, _, _, _, _ = context, GetContext_89_result, ctx, s, min_90_a, min_90_b, min_90_result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetShaderSource expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlReleaseShaderCompiler) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlReleaseShaderCompiler{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glReleaseShaderCompiler expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlCompileShader) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlCompileShader{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glCompileShader expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlIsShader) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlIsShader{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_91_result := context              // ContextPtr
	ctx := GetContext_91_result                  // ContextPtr
	ϟa.Result = ctx.Instances.Shaders.Contains(ϟa.Shader)
	_, _, _ = context, GetContext_91_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glIsShader expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlCreateProgram) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlCreateProgram{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_92_result := context              // ContextPtr
	ctx := GetContext_92_result                  // ContextPtr
	id := ProgramId(ϟa.Result)                   // ProgramId
	ctx.Instances.Programs[id] = func() *Program {
		s := &Program{}
		s.Init()
		return s
	}()
	ϟa.Result = id
	_, _, _, _ = context, GetContext_92_result, ctx, id
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glCreateProgram expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDeleteProgram) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlDeleteProgram{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_93_result := context              // ContextPtr
	ctx := GetContext_93_result                  // ContextPtr
	ctx.Instances.Programs.Delete(ϟa.Program)
	_, _, _ = context, GetContext_93_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDeleteProgram expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlAttachShader) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlAttachShader{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_94_result := context              // ContextPtr
	ctx := GetContext_94_result                  // ContextPtr
	p := ctx.Instances.Programs.Get(ϟa.Program)  // ProgramPtr
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // ShaderPtr
	p.Shaders[s.Type] = ϟa.Shader
	_, _, _, _, _ = context, GetContext_94_result, ctx, p, s
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glAttachShader expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDetachShader) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlDetachShader{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_95_result := context              // ContextPtr
	ctx := GetContext_95_result                  // ContextPtr
	p := ctx.Instances.Programs.Get(ϟa.Program)  // ProgramPtr
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // ShaderPtr
	p.Shaders.Delete(s.Type)
	_, _, _, _, _ = context, GetContext_95_result, ctx, p, s
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDetachShader expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetAttachedShaders) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGetAttachedShaders{}
	ϟo.Shaders = make(ShaderIdArray, ϟa.BufferLength)
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_96_result := context              // ContextPtr
	ctx := GetContext_96_result                  // ContextPtr
	p := ctx.Instances.Programs.Get(ϟa.Program)  // ProgramPtr
	min_97_a := ϟa.BufferLength                  // s32
	min_97_b := int32(len(p.Shaders))            // s32
	min_97_result := func() (result int32) {
		switch (min_97_a) < (min_97_b) {
		case true:
			return min_97_a
		case false:
			return min_97_b
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (min_97_a) < (min_97_b), ϟa))
			return result
		}
	}() // s32
	ϟa.ShadersLengthWritten = min_97_result
	ϟa.Shaders = p.Shaders.Range()
	_, _, _, _, _, _, _ = context, GetContext_96_result, ctx, p, min_97_a, min_97_b, min_97_result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetAttachedShaders expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlLinkProgram) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlLinkProgram{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glLinkProgram expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetProgramInfoLog) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGetProgramInfoLog{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_98_result := context              // ContextPtr
	ctx := GetContext_98_result                  // ContextPtr
	p := ctx.Instances.Programs.Get(ϟa.Program)  // ProgramPtr
	min_99_a := ϟa.BufferLength                  // s32
	min_99_b := strlen(p.InfoLog)                // s32
	min_99_result := func() (result int32) {
		switch (min_99_a) < (min_99_b) {
		case true:
			return min_99_a
		case false:
			return min_99_b
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (min_99_a) < (min_99_b), ϟa))
			return result
		}
	}() // s32
	ϟa.StringLengthWritten = min_99_result
	ϟa.Info = substr(p.InfoLog, int32(0), ϟa.StringLengthWritten)
	_, _, _, _, _, _, _ = context, GetContext_98_result, ctx, p, min_99_a, min_99_b, min_99_result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetProgramInfoLog expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUseProgram) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlUseProgram{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_100_result := context             // ContextPtr
	ctx := GetContext_100_result                 // ContextPtr
	ctx.BoundProgram = ϟa.Program
	_, _, _ = context, GetContext_100_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUseProgram expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlIsProgram) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlIsProgram{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_101_result := context             // ContextPtr
	ctx := GetContext_101_result                 // ContextPtr
	ϟa.Result = ctx.Instances.Programs.Contains(ϟa.Program)
	_, _, _ = context, GetContext_101_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glIsProgram expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlValidateProgram) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlValidateProgram{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glValidateProgram expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlClearColor) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlClearColor{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_102_result := context             // ContextPtr
	ctx := GetContext_102_result                 // ContextPtr
	ctx.Clearing.ClearColor = func() Color {
		s := Color{}
		s.Init()
		s.Red = ϟa.R
		s.Green = ϟa.G
		s.Blue = ϟa.B
		s.Alpha = ϟa.A
		return s
	}()
	_, _, _ = context, GetContext_102_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glClearColor expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlClearDepthf) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlClearDepthf{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_103_result := context             // ContextPtr
	ctx := GetContext_103_result                 // ContextPtr
	ctx.Clearing.ClearDepth = ϟa.Depth
	_, _, _ = context, GetContext_103_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glClearDepthf expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlClearStencil) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlClearStencil{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_104_result := context             // ContextPtr
	ctx := GetContext_104_result                 // ContextPtr
	ctx.Clearing.ClearStencil = ϟa.Stencil
	_, _, _ = context, GetContext_104_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glClearStencil expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlClear) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlClear{}
	if (ClearMask_GL_COLOR_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glClear expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlCullFace) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlCullFace{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_105_result := context             // ContextPtr
	ctx := GetContext_105_result                 // ContextPtr
	ctx.Rasterizing.CullFace = ϟa.Mode
	_, _, _ = context, GetContext_105_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glCullFace expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlPolygonOffset) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlPolygonOffset{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_106_result := context             // ContextPtr
	ctx := GetContext_106_result                 // ContextPtr
	ctx.Rasterizing.PolygonOffsetUnits = ϟa.Units
	ctx.Rasterizing.PolygonOffsetFactor = ϟa.ScaleFactor
	_, _, _ = context, GetContext_106_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glPolygonOffset expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlLineWidth) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlLineWidth{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_107_result := context             // ContextPtr
	ctx := GetContext_107_result                 // ContextPtr
	ctx.Rasterizing.LineWidth = ϟa.Width
	_, _, _ = context, GetContext_107_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glLineWidth expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlSampleCoverage) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlSampleCoverage{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_108_result := context             // ContextPtr
	ctx := GetContext_108_result                 // ContextPtr
	ctx.Rasterizing.SampleCoverageValue = ϟa.Value
	ctx.Rasterizing.SampleCoverageInvert = ϟa.Invert
	_, _, _ = context, GetContext_108_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glSampleCoverage expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlHint) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlHint{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_109_result := context             // ContextPtr
	ctx := GetContext_109_result                 // ContextPtr
	ctx.GenerateMipmapHint = ϟa.Mode
	_, _, _ = context, GetContext_109_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glHint expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlFramebufferRenderbuffer) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlFramebufferRenderbuffer{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_110_result := context             // ContextPtr
	ctx := GetContext_110_result                 // ContextPtr
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
	framebuffer := ctx.Instances.Framebuffers.Get(framebufferId)        // FramebufferPtr
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
	_, _, _, _, _, _, _ = context, GetContext_110_result, ctx, target, framebufferId, framebuffer, attachment
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glFramebufferRenderbuffer expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlFramebufferTexture2D) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlFramebufferTexture2D{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_111_result := context             // ContextPtr
	ctx := GetContext_111_result                 // ContextPtr
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
	framebuffer := ctx.Instances.Framebuffers.Get(framebufferId)        // FramebufferPtr
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
	_, _, _, _, _, _, _ = context, GetContext_111_result, ctx, target, framebufferId, framebuffer, attachment
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glFramebufferTexture2D expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetFramebufferAttachmentParameteriv) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGetFramebufferAttachmentParameteriv{}
	ϟo.Value = make(S32Array, int32(1))
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_112_result := context             // ContextPtr
	ctx := GetContext_112_result                 // ContextPtr
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
	framebuffer := ctx.Instances.Framebuffers.Get(framebufferId) // FramebufferPtr
	a := framebuffer.Attachments.Get(ϟa.Attachment)              // FramebufferAttachmentInfo
	ϟa.Value[int32(0)] = func() (result int32) {
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
	}()
	_, _, _, _, _, _, _ = context, GetContext_112_result, ctx, target, framebufferId, framebuffer, a
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetFramebufferAttachmentParameteriv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDrawElements) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlDrawElements{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                     // ContextPtr
	GetContext_113_result := context                                 // ContextPtr
	ctx := GetContext_113_result                                     // ContextPtr
	id := ctx.BoundBuffers.Get(BufferTarget_GL_ELEMENT_ARRAY_BUFFER) // BufferId
	index_data := func() (result memory.Pointer) {
		switch (id) != (BufferId(uint32(0))) {
		case true:
			return memoryOffset(ctx.Instances.Buffers.Get(id).Data, uint64(ϟa.Indices))
		case false:
			return memory.Pointer(ϟa.Indices)
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (id) != (BufferId(uint32(0))), ϟa))
			return result
		}
	}() // VoidArray
	IndexSize_114_indices_type := ϟa.IndicesType // IndicesType
	IndexSize_114_result := func() (result uint32) {
		switch IndexSize_114_indices_type {
		case IndicesType_GL_UNSIGNED_BYTE:
			return uint32(1)
		case IndicesType_GL_UNSIGNED_SHORT:
			return uint32(2)
		case IndicesType_GL_UNSIGNED_INT:
			return uint32(4)
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", IndexSize_114_indices_type, ϟa))
			return result
		}
	}() // u32
	read(index_data, uint32(0), (uint32(ϟa.ElementCount))*(IndexSize_114_result))
	first := minIndex(index_data, ϟa.IndicesType, uint32(ϟa.ElementCount)) // u32
	last := maxIndex(index_data, ϟa.IndicesType, uint32(ϟa.ElementCount))  // u32
	ReadVertexArrays_115_ctx := ctx                                        // ContextPtr
	ReadVertexArrays_115_first_index := first                              // u32
	ReadVertexArrays_115_last_index := last                                // u32
	for i := int32(int32(0)); i < int32(len(ReadVertexArrays_115_ctx.VertexAttributeArrays)); i++ {
		arr := ReadVertexArrays_115_ctx.VertexAttributeArrays.Get(AttributeLocation(i)) // VertexAttributeArrayPtr
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
				read(arr.Pointer, offset, elsize)
				_ = offset
			}
			_, _, _, _ = vertexAttribTypeSize_116_t, vertexAttribTypeSize_116_result, elsize, elstride
		}
		_ = arr
	}
	_, _, _, _, _, _, _, _, _, _, _, _ = context, GetContext_113_result, ctx, id, index_data, IndexSize_114_indices_type, IndexSize_114_result, first, last, ReadVertexArrays_115_ctx, ReadVertexArrays_115_first_index, ReadVertexArrays_115_last_index
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDrawElements expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDrawArrays) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlDrawArrays{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                   // ContextPtr
	GetContext_117_result := context                               // ContextPtr
	ctx := GetContext_117_result                                   // ContextPtr
	last_index := (ϟa.FirstIndex) + ((ϟa.IndexCount) - (int32(1))) // s32
	ReadVertexArrays_118_ctx := ctx                                // ContextPtr
	ReadVertexArrays_118_first_index := uint32(ϟa.FirstIndex)      // u32
	ReadVertexArrays_118_last_index := uint32(last_index)          // u32
	for i := int32(int32(0)); i < int32(len(ReadVertexArrays_118_ctx.VertexAttributeArrays)); i++ {
		arr := ReadVertexArrays_118_ctx.VertexAttributeArrays.Get(AttributeLocation(i)) // VertexAttributeArrayPtr
		if (arr.Enabled) && ((arr.Buffer) == (BufferId(uint32(0)))) {
			vertexAttribTypeSize_119_t := arr.Type // VertexAttribType
			vertexAttribTypeSize_119_result := func() (result uint32) {
				switch vertexAttribTypeSize_119_t {
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
					panic(fmt.Errorf("Unmatched switch(%v) in atom %T", vertexAttribTypeSize_119_t, ϟa))
					return result
				}
			}() // u32
			elsize := (vertexAttribTypeSize_119_result) * (arr.Size) // u32
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
			for v := uint32(ReadVertexArrays_118_first_index); v < (ReadVertexArrays_118_last_index)+(uint32(1)); v++ {
				offset := (elstride) * (v) // u32
				read(arr.Pointer, offset, elsize)
				_ = offset
			}
			_, _, _, _ = vertexAttribTypeSize_119_t, vertexAttribTypeSize_119_result, elsize, elstride
		}
		_ = arr
	}
	_, _, _, _, _, _, _ = context, GetContext_117_result, ctx, last_index, ReadVertexArrays_118_ctx, ReadVertexArrays_118_first_index, ReadVertexArrays_118_last_index
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDrawArrays expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlFlush) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlFlush{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glFlush expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlFinish) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlFinish{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glFinish expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetBooleanv) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGetBooleanv{}
	ϟo.Values = make(BoolArray, stateVariableSize(ϟa.Param))
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_120_result := context             // ContextPtr
	ctx := GetContext_120_result                 // ContextPtr
	switch ϟa.Param {
	case StateVariable_GL_BLEND:
		ϟa.Values[int32(0)] = ctx.Capabilities.Get(Capability_GL_BLEND)
	case StateVariable_GL_CULL_FACE:
		ϟa.Values[int32(0)] = ctx.Capabilities.Get(Capability_GL_CULL_FACE)
	case StateVariable_GL_DEPTH_TEST:
		ϟa.Values[int32(0)] = ctx.Capabilities.Get(Capability_GL_DEPTH_TEST)
	case StateVariable_GL_DITHER:
		ϟa.Values[int32(0)] = ctx.Capabilities.Get(Capability_GL_DITHER)
	case StateVariable_GL_POLYGON_OFFSET_FILL:
		ϟa.Values[int32(0)] = ctx.Capabilities.Get(Capability_GL_POLYGON_OFFSET_FILL)
	case StateVariable_GL_SAMPLE_ALPHA_TO_COVERAGE:
		ϟa.Values[int32(0)] = ctx.Capabilities.Get(Capability_GL_SAMPLE_ALPHA_TO_COVERAGE)
	case StateVariable_GL_SAMPLE_COVERAGE:
		ϟa.Values[int32(0)] = ctx.Capabilities.Get(Capability_GL_SAMPLE_COVERAGE)
	case StateVariable_GL_SCISSOR_TEST:
		ϟa.Values[int32(0)] = ctx.Capabilities.Get(Capability_GL_SCISSOR_TEST)
	case StateVariable_GL_STENCIL_TEST:
		ϟa.Values[int32(0)] = ctx.Capabilities.Get(Capability_GL_STENCIL_TEST)
	case StateVariable_GL_DEPTH_WRITEMASK:
		ϟa.Values[int32(0)] = ctx.Rasterizing.DepthMask
	case StateVariable_GL_COLOR_WRITEMASK:
		ϟa.Values[int32(0)] = ctx.Rasterizing.ColorMaskRed
		ϟa.Values[int32(1)] = ctx.Rasterizing.ColorMaskGreen
		ϟa.Values[int32(2)] = ctx.Rasterizing.ColorMaskBlue
		ϟa.Values[int32(3)] = ctx.Rasterizing.ColorMaskAlpha
	case StateVariable_GL_SAMPLE_COVERAGE_INVERT:
		ϟa.Values[int32(0)] = ctx.Rasterizing.SampleCoverageInvert
	case StateVariable_GL_SHADER_COMPILER:
		ϟa.Values[int32(0)] = ϟa.Values[int32(0)]
	case StateVariable_GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT:
		ϟa.Values[int32(0)] = ϟa.Values[int32(0)]
	default:
		// TODO: better unmatched handling
		v := ϟa.Param
		log.Printf("Error: Missing switch case handler for value %T %v", v, v)
	}
	_, _, _ = context, GetContext_120_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetBooleanv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetFloatv) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGetFloatv{}
	ϟo.Values = make(F32Array, stateVariableSize(ϟa.Param))
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_121_result := context             // ContextPtr
	ctx := GetContext_121_result                 // ContextPtr
	switch ϟa.Param {
	case StateVariable_GL_DEPTH_RANGE:
		ϟa.Values[int32(0)] = ctx.Rasterizing.DepthNear
		ϟa.Values[int32(1)] = ctx.Rasterizing.DepthFar
	case StateVariable_GL_LINE_WIDTH:
		ϟa.Values[int32(0)] = ctx.Rasterizing.LineWidth
	case StateVariable_GL_POLYGON_OFFSET_FACTOR:
		ϟa.Values[int32(0)] = ctx.Rasterizing.PolygonOffsetFactor
	case StateVariable_GL_POLYGON_OFFSET_UNITS:
		ϟa.Values[int32(0)] = ctx.Rasterizing.PolygonOffsetUnits
	case StateVariable_GL_SAMPLE_COVERAGE_VALUE:
		ϟa.Values[int32(0)] = ctx.Rasterizing.SampleCoverageValue
	case StateVariable_GL_COLOR_CLEAR_VALUE:
		ϟa.Values[int32(0)] = ctx.Clearing.ClearColor.Red
		ϟa.Values[int32(1)] = ctx.Clearing.ClearColor.Green
		ϟa.Values[int32(2)] = ctx.Clearing.ClearColor.Blue
		ϟa.Values[int32(3)] = ctx.Clearing.ClearColor.Alpha
	case StateVariable_GL_DEPTH_CLEAR_VALUE:
		ϟa.Values[int32(0)] = ctx.Clearing.ClearDepth
	case StateVariable_GL_ALIASED_LINE_WIDTH_RANGE:
		ϟa.Values[int32(0)] = ϟa.Values[int32(0)]
		ϟa.Values[int32(1)] = ϟa.Values[int32(1)]
	case StateVariable_GL_ALIASED_POINT_SIZE_RANGE:
		ϟa.Values[int32(0)] = ϟa.Values[int32(0)]
		ϟa.Values[int32(1)] = ϟa.Values[int32(1)]
	case StateVariable_GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT:
		ϟa.Values[int32(0)] = ϟa.Values[int32(0)]
	default:
		// TODO: better unmatched handling
		v := ϟa.Param
		log.Printf("Error: Missing switch case handler for value %T %v", v, v)
	}
	_, _, _ = context, GetContext_121_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetFloatv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetIntegerv) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGetIntegerv{}
	ϟo.Values = make(S32Array, stateVariableSize(ϟa.Param))
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_122_result := context             // ContextPtr
	ctx := GetContext_122_result                 // ContextPtr
	switch ϟa.Param {
	case StateVariable_GL_ACTIVE_TEXTURE:
		ϟa.Values[int32(0)] = int32(ctx.ActiveTextureUnit)
	case StateVariable_GL_ARRAY_BUFFER_BINDING:
		ϟa.Values[int32(0)] = int32(ctx.BoundBuffers.Get(BufferTarget_GL_ARRAY_BUFFER))
	case StateVariable_GL_ELEMENT_ARRAY_BUFFER_BINDING:
		ϟa.Values[int32(0)] = int32(ctx.BoundBuffers.Get(BufferTarget_GL_ELEMENT_ARRAY_BUFFER))
	case StateVariable_GL_BLEND_SRC_ALPHA:
		ϟa.Values[int32(0)] = int32(ctx.Blending.SrcAlphaBlendFactor)
	case StateVariable_GL_BLEND_SRC_RGB:
		ϟa.Values[int32(0)] = int32(ctx.Blending.SrcRgbBlendFactor)
	case StateVariable_GL_BLEND_DST_ALPHA:
		ϟa.Values[int32(0)] = int32(ctx.Blending.DstAlphaBlendFactor)
	case StateVariable_GL_BLEND_DST_RGB:
		ϟa.Values[int32(0)] = int32(ctx.Blending.DstRgbBlendFactor)
	case StateVariable_GL_BLEND_EQUATION_RGB:
		ϟa.Values[int32(0)] = int32(ctx.Blending.BlendEquationRgb)
	case StateVariable_GL_BLEND_EQUATION_ALPHA:
		ϟa.Values[int32(0)] = int32(ctx.Blending.BlendEquationAlpha)
	case StateVariable_GL_BLEND_COLOR:
		ϟa.Values[int32(0)] = int32(ctx.Blending.BlendColor.Red)
		ϟa.Values[int32(1)] = int32(ctx.Blending.BlendColor.Green)
		ϟa.Values[int32(2)] = int32(ctx.Blending.BlendColor.Blue)
		ϟa.Values[int32(3)] = int32(ctx.Blending.BlendColor.Alpha)
	case StateVariable_GL_DEPTH_FUNC:
		ϟa.Values[int32(0)] = int32(ctx.Rasterizing.DepthTestFunction)
	case StateVariable_GL_DEPTH_CLEAR_VALUE:
		ϟa.Values[int32(0)] = int32(ctx.Clearing.ClearDepth)
	case StateVariable_GL_STENCIL_WRITEMASK:
		ϟa.Values[int32(0)] = int32(ctx.Rasterizing.StencilMask.Get(FaceMode_GL_FRONT))
	case StateVariable_GL_STENCIL_BACK_WRITEMASK:
		ϟa.Values[int32(0)] = int32(ctx.Rasterizing.StencilMask.Get(FaceMode_GL_BACK))
	case StateVariable_GL_VIEWPORT:
		ϟa.Values[int32(0)] = ctx.Rasterizing.Viewport.X
		ϟa.Values[int32(1)] = ctx.Rasterizing.Viewport.Y
		ϟa.Values[int32(2)] = ctx.Rasterizing.Viewport.Width
		ϟa.Values[int32(3)] = ctx.Rasterizing.Viewport.Height
	case StateVariable_GL_SCISSOR_BOX:
		ϟa.Values[int32(0)] = ctx.Rasterizing.Scissor.X
		ϟa.Values[int32(1)] = ctx.Rasterizing.Scissor.Y
		ϟa.Values[int32(2)] = ctx.Rasterizing.Scissor.Width
		ϟa.Values[int32(3)] = ctx.Rasterizing.Scissor.Height
	case StateVariable_GL_FRONT_FACE:
		ϟa.Values[int32(0)] = int32(ctx.Rasterizing.FrontFace)
	case StateVariable_GL_CULL_FACE_MODE:
		ϟa.Values[int32(0)] = int32(ctx.Rasterizing.CullFace)
	case StateVariable_GL_STENCIL_CLEAR_VALUE:
		ϟa.Values[int32(0)] = ctx.Clearing.ClearStencil
	case StateVariable_GL_FRAMEBUFFER_BINDING:
		ϟa.Values[int32(0)] = int32(ctx.BoundFramebuffers.Get(FramebufferTarget_GL_FRAMEBUFFER))
	case StateVariable_GL_READ_FRAMEBUFFER_BINDING:
		ϟa.Values[int32(0)] = int32(ctx.BoundFramebuffers.Get(FramebufferTarget_GL_READ_FRAMEBUFFER))
	case StateVariable_GL_RENDERBUFFER_BINDING:
		ϟa.Values[int32(0)] = int32(ctx.BoundRenderbuffers.Get(RenderbufferTarget_GL_RENDERBUFFER))
	case StateVariable_GL_CURRENT_PROGRAM:
		ϟa.Values[int32(0)] = int32(ctx.BoundProgram)
	case StateVariable_GL_TEXTURE_BINDING_2D:
		ϟa.Values[int32(0)] = int32(ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_2D))
	case StateVariable_GL_TEXTURE_BINDING_CUBE_MAP:
		ϟa.Values[int32(0)] = int32(ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_CUBE_MAP))
	case StateVariable_GL_GENERATE_MIPMAP_HINT:
		ϟa.Values[int32(0)] = int32(ctx.GenerateMipmapHint)
	case StateVariable_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS:
		ϟa.Values[int32(0)] = ϟa.Values[int32(0)]
	case StateVariable_GL_MAX_CUBE_MAP_TEXTURE_SIZE:
		ϟa.Values[int32(0)] = ϟa.Values[int32(0)]
	case StateVariable_GL_MAX_FRAGMENT_UNIFORM_VECTORS:
		ϟa.Values[int32(0)] = ϟa.Values[int32(0)]
	case StateVariable_GL_MAX_RENDERBUFFER_SIZE:
		ϟa.Values[int32(0)] = ϟa.Values[int32(0)]
	case StateVariable_GL_MAX_TEXTURE_IMAGE_UNITS:
		ϟa.Values[int32(0)] = ϟa.Values[int32(0)]
	case StateVariable_GL_MAX_TEXTURE_SIZE:
		ϟa.Values[int32(0)] = ϟa.Values[int32(0)]
	case StateVariable_GL_MAX_VARYING_VECTORS:
		ϟa.Values[int32(0)] = ϟa.Values[int32(0)]
	case StateVariable_GL_MAX_VERTEX_ATTRIBS:
		ϟa.Values[int32(0)] = ϟa.Values[int32(0)]
	case StateVariable_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS:
		ϟa.Values[int32(0)] = ϟa.Values[int32(0)]
	case StateVariable_GL_MAX_VERTEX_UNIFORM_VECTORS:
		ϟa.Values[int32(0)] = ϟa.Values[int32(0)]
	case StateVariable_GL_MAX_VIEWPORT_DIMS:
		ϟa.Values[int32(0)] = ϟa.Values[int32(0)]
		ϟa.Values[int32(1)] = ϟa.Values[int32(1)]
	case StateVariable_GL_NUM_COMPRESSED_TEXTURE_FORMATS:
		ϟa.Values[int32(0)] = ϟa.Values[int32(0)]
	case StateVariable_GL_NUM_SHADER_BINARY_FORMATS:
		ϟa.Values[int32(0)] = ϟa.Values[int32(0)]
	case StateVariable_GL_PACK_ALIGNMENT:
		ϟa.Values[int32(0)] = ctx.PixelStorage.Get(PixelStoreParameter_GL_PACK_ALIGNMENT)
	case StateVariable_GL_UNPACK_ALIGNMENT:
		ϟa.Values[int32(0)] = ctx.PixelStorage.Get(PixelStoreParameter_GL_UNPACK_ALIGNMENT)
	case StateVariable_GL_ALPHA_BITS:
		ϟa.Values[int32(0)] = ϟa.Values[int32(0)]
	case StateVariable_GL_BLUE_BITS:
		ϟa.Values[int32(0)] = ϟa.Values[int32(0)]
	case StateVariable_GL_GREEN_BITS:
		ϟa.Values[int32(0)] = ϟa.Values[int32(0)]
	case StateVariable_GL_RED_BITS:
		ϟa.Values[int32(0)] = ϟa.Values[int32(0)]
	case StateVariable_GL_DEPTH_BITS:
		ϟa.Values[int32(0)] = ϟa.Values[int32(0)]
	case StateVariable_GL_SAMPLE_BUFFERS:
		ϟa.Values[int32(0)] = ϟa.Values[int32(0)]
	case StateVariable_GL_SAMPLES:
		ϟa.Values[int32(0)] = ϟa.Values[int32(0)]
	case StateVariable_GL_SHADER_BINARY_FORMATS:
		ϟa.Values = ϟa.Values
	case StateVariable_GL_COMPRESSED_TEXTURE_FORMATS:
		ϟa.Values = ϟa.Values
	case StateVariable_GL_STENCIL_BITS:
		ϟa.Values[int32(0)] = ϟa.Values[int32(0)]
	case StateVariable_GL_SUBPIXEL_BITS:
		ϟa.Values[int32(0)] = ϟa.Values[int32(0)]
	case StateVariable_GL_IMPLEMENTATION_COLOR_READ_FORMAT:
		ϟa.Values[int32(0)] = ϟa.Values[int32(0)]
	case StateVariable_GL_IMPLEMENTATION_COLOR_READ_TYPE:
		ϟa.Values[int32(0)] = ϟa.Values[int32(0)]
	case StateVariable_GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT:
		ϟa.Values[int32(0)] = ϟa.Values[int32(0)]
	case StateVariable_GL_GPU_DISJOINT_EXT:
		ϟa.Values[int32(0)] = ϟa.Values[int32(0)]
	default:
		// TODO: better unmatched handling
		v := ϟa.Param
		log.Printf("Error: Missing switch case handler for value %T %v", v, v)
	}
	_, _, _ = context, GetContext_122_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetIntegerv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetString) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGetString{}
	ϟa.Result = ϟa.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetString expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlEnable) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlEnable{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_123_result := context             // ContextPtr
	ctx := GetContext_123_result                 // ContextPtr
	ctx.Capabilities[ϟa.Capability] = true
	_, _, _ = context, GetContext_123_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glEnable expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDisable) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlDisable{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_124_result := context             // ContextPtr
	ctx := GetContext_124_result                 // ContextPtr
	ctx.Capabilities[ϟa.Capability] = false
	_, _, _ = context, GetContext_124_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDisable expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlIsEnabled) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlIsEnabled{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_125_result := context             // ContextPtr
	ctx := GetContext_125_result                 // ContextPtr
	ϟa.Result = ctx.Capabilities.Get(ϟa.Capability)
	_, _, _ = context, GetContext_125_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glIsEnabled expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlMapBufferRange) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlMapBufferRange{}
	ϟa.Result = ϟa.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glMapBufferRange expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUnmapBuffer) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlUnmapBuffer{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUnmapBuffer expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlInvalidateFramebuffer) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlInvalidateFramebuffer{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glInvalidateFramebuffer expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlRenderbufferStorageMultisample) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlRenderbufferStorageMultisample{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glRenderbufferStorageMultisample expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlBlitFramebuffer) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlBlitFramebuffer{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glBlitFramebuffer expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGenQueries) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGenQueries{}
	ϟo.Queries = make(QueryIdArray, ϟa.Count)
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_126_result := context             // ContextPtr
	ctx := GetContext_126_result                 // ContextPtr
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		id := QueryId(ϟa.Queries[i]) // QueryId
		ctx.Instances.Queries[id] = func() *Query {
			s := &Query{}
			s.Init()
			return s
		}()
		ϟa.Queries[i] = id
		_ = id
	}
	_, _, _ = context, GetContext_126_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGenQueries expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlBeginQuery) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlBeginQuery{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glBeginQuery expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlEndQuery) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlEndQuery{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glEndQuery expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDeleteQueries) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlDeleteQueries{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_127_result := context             // ContextPtr
	ctx := GetContext_127_result                 // ContextPtr
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		ctx.Instances.Queries.Delete(ϟa.Queries[i])
	}
	_, _, _ = context, GetContext_127_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDeleteQueries expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlIsQuery) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlIsQuery{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_128_result := context             // ContextPtr
	ctx := GetContext_128_result                 // ContextPtr
	ϟa.Result = ctx.Instances.Queries.Contains(ϟa.Query)
	_, _, _ = context, GetContext_128_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glIsQuery expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetQueryiv) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGetQueryiv{}
	ϟa.Value = ϟa.Value
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetQueryiv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetQueryObjectuiv) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGetQueryObjectuiv{}
	ϟa.Value = ϟa.Value
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetQueryObjectuiv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGenQueriesEXT) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGenQueriesEXT{}
	ϟo.Queries = make(QueryIdArray, ϟa.Count)
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_129_result := context             // ContextPtr
	ctx := GetContext_129_result                 // ContextPtr
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		id := QueryId(ϟa.Queries[i]) // QueryId
		ctx.Instances.Queries[id] = func() *Query {
			s := &Query{}
			s.Init()
			return s
		}()
		ϟa.Queries[i] = id
		_ = id
	}
	_, _, _ = context, GetContext_129_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGenQueriesEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlBeginQueryEXT) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlBeginQueryEXT{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glBeginQueryEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlEndQueryEXT) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlEndQueryEXT{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glEndQueryEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDeleteQueriesEXT) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlDeleteQueriesEXT{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_130_result := context             // ContextPtr
	ctx := GetContext_130_result                 // ContextPtr
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		ctx.Instances.Queries.Delete(ϟa.Queries[i])
	}
	_, _, _ = context, GetContext_130_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDeleteQueriesEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlIsQueryEXT) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlIsQueryEXT{}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // ContextPtr
	GetContext_131_result := context             // ContextPtr
	ctx := GetContext_131_result                 // ContextPtr
	ϟa.Result = ctx.Instances.Queries.Contains(ϟa.Query)
	_, _, _ = context, GetContext_131_result, ctx
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glIsQueryEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlQueryCounterEXT) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlQueryCounterEXT{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glQueryCounterEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetQueryivEXT) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGetQueryivEXT{}
	ϟa.Value = ϟa.Value
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetQueryivEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetQueryObjectivEXT) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGetQueryObjectivEXT{}
	ϟa.Value = ϟa.Value
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetQueryObjectivEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetQueryObjectuivEXT) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGetQueryObjectuivEXT{}
	ϟa.Value = ϟa.Value
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetQueryObjectuivEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetQueryObjecti64vEXT) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGetQueryObjecti64vEXT{}
	ϟa.Value = ϟa.Value
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetQueryObjecti64vEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetQueryObjectui64vEXT) Mutate(ϟs *gfxapi.State) error {
	ϟc := getState(ϟs)
	ϟo := GlGetQueryObjectui64vEXT{}
	ϟa.Value = ϟa.Value
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetQueryObjectui64vEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
