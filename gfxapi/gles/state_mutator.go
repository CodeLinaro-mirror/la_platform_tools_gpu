////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

package gles

import (
	"fmt"
	"log"
	"reflect"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/gfxapi/state"
	"android.googlesource.com/platform/tools/gpu/memory"
)

func getState(a atom.Atom, s *state.State) *State {
	id := a.ContextID()
	if state, ok := s.Contexts[id].(*State); ok {
		return state
	} else {
		panic(fmt.Errorf("State for atom %T with context id %d was %T, expected *gles.State",
			a, id, s.Contexts[id]))
	}
}

func (ϟa *Init) Mutate(ϟs *state.State) error {
	ϟc := &State{}
	ϟc.Init()
	ϟs.Contexts[ϟa.ContextID()] = ϟc
	ϟo := Init{}
	ϟc.Instances.Buffers[ϟc.Internals.NilBuffer] = func() *Buffer {
		s := &Buffer{}
		s.Init()
		return s
	}()
	ϟc.Instances.Textures[ϟc.Internals.NilTexture] = func() *Texture {
		s := &Texture{}
		s.Init()
		return s
	}()
	ϟc.Instances.Renderbuffers[ϟc.Internals.NilRenderbuffer] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	backbufferColorId := RenderbufferId(uint32(4294967295))   // RenderbufferId
	backbufferDepthId := RenderbufferId(uint32(4294967294))   // RenderbufferId
	backbufferStencilId := RenderbufferId(uint32(4294967293)) // RenderbufferId
	backbufferColor := func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}() // RenderbufferPtr
	backbufferColor.Width = ϟa.Width
	backbufferColor.Height = ϟa.Height
	backbufferColor.Format = ϟa.ColorFmt
	ϟc.Instances.Renderbuffers[backbufferColorId] = backbufferColor
	backbufferDepth := func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}() // RenderbufferPtr
	backbufferDepth.Width = ϟa.Width
	backbufferDepth.Height = ϟa.Height
	backbufferDepth.Format = ϟa.DepthFmt
	ϟc.Instances.Renderbuffers[backbufferDepthId] = backbufferDepth
	backbufferStencil := func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}() // RenderbufferPtr
	backbufferStencil.Width = ϟa.Width
	backbufferStencil.Height = ϟa.Height
	backbufferStencil.Format = ϟa.StencilFmt
	ϟc.Instances.Renderbuffers[backbufferStencilId] = backbufferStencil
	backbuffer := func() *Framebuffer {
		s := &Framebuffer{}
		s.Init()
		return s
	}() // FramebufferPtr
	backbuffer.Attachments[FramebufferAttachment_GL_COLOR_ATTACHMENT0] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(backbufferColorId)
		s.Type = FramebufferAttachmentType_GL_RENDERBUFFER
		s.CubeMapFace = CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	backbuffer.Attachments[FramebufferAttachment_GL_DEPTH_ATTACHMENT] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(backbufferDepthId)
		s.Type = FramebufferAttachmentType_GL_RENDERBUFFER
		s.CubeMapFace = CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	backbuffer.Attachments[FramebufferAttachment_GL_STENCIL_ATTACHMENT] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(backbufferStencilId)
		s.Type = FramebufferAttachmentType_GL_RENDERBUFFER
		s.CubeMapFace = CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	ϟc.Instances.Framebuffers[ϟc.Internals.Backbuffer] = backbuffer
	ϟc.BoundFramebuffers[FramebufferTarget_GL_DRAW_FRAMEBUFFER] = ϟc.Internals.Backbuffer
	ϟc.BoundFramebuffers[FramebufferTarget_GL_READ_FRAMEBUFFER] = ϟc.Internals.Backbuffer
	ϟc.Rasterizing.Scissor.Width = ϟa.Width
	ϟc.Rasterizing.Scissor.Height = ϟa.Height
	ϟc.Rasterizing.StencilMask[FaceMode_GL_FRONT] = uint32(4294967295)
	ϟc.Rasterizing.StencilMask[FaceMode_GL_BACK] = uint32(4294967295)
	ϟc.Rasterizing.Viewport.Width = ϟa.Width
	ϟc.Rasterizing.Viewport.Height = ϟa.Height
	ϟc.PixelStorage[PixelStoreParameter_GL_PACK_ALIGNMENT] = int32(4)
	ϟc.PixelStorage[PixelStoreParameter_GL_UNPACK_ALIGNMENT] = int32(4)
	for i := int32(int32(0)); i < int32(64); i++ {
		ϟc.VertexAttributeArrays[AttributeLocation(i)] = func() *VertexAttributeArray {
			s := &VertexAttributeArray{}
			s.Init()
			return s
		}()
	}
	_, _, _, _, _, _, _ = backbufferColorId, backbufferDepthId, backbufferStencilId, backbufferColor, backbufferDepth, backbufferStencil, backbuffer
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying init expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *StartTimer) Mutate(ϟs *state.State) error {
	ϟc := &State{}
	ϟo := StartTimer{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying startTimer expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *StopTimer) Mutate(ϟs *state.State) error {
	ϟc := &State{}
	ϟo := StopTimer{}
	ϟa.Result = ϟa.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying stopTimer expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *FlushPostBuffer) Mutate(ϟs *state.State) error {
	ϟc := &State{}
	ϟo := FlushPostBuffer{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying flushPostBuffer expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *EglInitialize) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := EglInitialize{}
	ϟa.Major = ϟa.Major
	ϟa.Minor = ϟa.Minor
	ϟa.Result = ϟa.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying eglInitialize expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *EglCreateContext) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := EglCreateContext{}
	ϟa.Result = ϟa.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying eglCreateContext expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *EglMakeCurrent) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := EglMakeCurrent{}
	ϟa.Result = ϟa.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying eglMakeCurrent expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *EglSwapBuffers) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := EglSwapBuffers{}
	ϟa.Result = ϟa.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying eglSwapBuffers expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *WglCreateContext) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := WglCreateContext{}
	ϟa.Result = ϟa.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying wglCreateContext expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *WglMakeCurrent) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := WglMakeCurrent{}
	ϟa.Result = ϟa.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying wglMakeCurrent expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *WglSwapBuffers) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := WglSwapBuffers{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying wglSwapBuffers expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *CGLCreateContext) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CGLCreateContext{}
	ϟa.Result = ϟa.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying CGLCreateContext expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlEnableClientState) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlEnableClientState{}
	ϟc.Capabilities[Capability(ϟa.Type)] = true
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glEnableClientState expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDisableClientState) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDisableClientState{}
	ϟc.Capabilities[Capability(ϟa.Type)] = false
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDisableClientState expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetProgramBinaryOES) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetProgramBinaryOES{}
	ϟa.BytesWritten = ϟa.BytesWritten
	ϟa.BinaryFormat = ϟa.BinaryFormat
	ϟa.Binary = ϟa.Binary
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetProgramBinaryOES expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlProgramBinaryOES) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlProgramBinaryOES{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glProgramBinaryOES expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlStartTilingQCOM) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlStartTilingQCOM{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glStartTilingQCOM expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlEndTilingQCOM) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlEndTilingQCOM{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glEndTilingQCOM expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDiscardFramebufferEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDiscardFramebufferEXT{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDiscardFramebufferEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlInsertEventMarkerEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlInsertEventMarkerEXT{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glInsertEventMarkerEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlPushGroupMarkerEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlPushGroupMarkerEXT{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glPushGroupMarkerEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlPopGroupMarkerEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlPopGroupMarkerEXT{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glPopGroupMarkerEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlTexStorage1DEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlTexStorage1DEXT{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glTexStorage1DEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlTexStorage2DEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlTexStorage2DEXT{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glTexStorage2DEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlTexStorage3DEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlTexStorage3DEXT{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glTexStorage3DEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlTextureStorage1DEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlTextureStorage1DEXT{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glTextureStorage1DEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlTextureStorage2DEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlTextureStorage2DEXT{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glTextureStorage2DEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlTextureStorage3DEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlTextureStorage3DEXT{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glTextureStorage3DEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGenVertexArraysOES) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGenVertexArraysOES{}
	ϟo.Arrays = make(VertexArrayIdArray, ϟa.Count)
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		id := VertexArrayId(ϟa.Arrays[i]) // VertexArrayId
		ϟc.Instances.VertexArrays[id] = func() *VertexArray {
			s := &VertexArray{}
			s.Init()
			return s
		}()
		ϟa.Arrays[i] = id
		_ = id
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGenVertexArraysOES expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlBindVertexArrayOES) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlBindVertexArrayOES{}
	if !(ϟc.Instances.VertexArrays.Contains(ϟa.Array)) {
		ϟc.Instances.VertexArrays[ϟa.Array] = func() *VertexArray {
			s := &VertexArray{}
			s.Init()
			return s
		}()
	}
	ϟc.BoundVertexArray = ϟa.Array
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glBindVertexArrayOES expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDeleteVertexArraysOES) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDeleteVertexArraysOES{}
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		ϟc.Instances.VertexArrays.Delete(ϟa.Arrays[i])
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDeleteVertexArraysOES expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlIsVertexArrayOES) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlIsVertexArrayOES{}
	ϟa.Result = ϟc.Instances.VertexArrays.Contains(ϟa.Array)
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glIsVertexArrayOES expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlEGLImageTargetTexture2DOES) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlEGLImageTargetTexture2DOES{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glEGLImageTargetTexture2DOES expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlEGLImageTargetRenderbufferStorageOES) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlEGLImageTargetRenderbufferStorageOES{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glEGLImageTargetRenderbufferStorageOES expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetGraphicsResetStatusEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetGraphicsResetStatusEXT{}
	ϟa.Result = ϟa.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetGraphicsResetStatusEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlBindAttribLocation) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlBindAttribLocation{}
	p := ϟc.Instances.Programs.Get(ϟa.Program) // ProgramPtr
	p.AttributeBindings[ϟa.Name] = ϟa.Location
	_ = p
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glBindAttribLocation expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlBlendFunc) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlBlendFunc{}
	ϟc.Blending.SrcRgbBlendFactor = ϟa.SrcFactor
	ϟc.Blending.SrcAlphaBlendFactor = ϟa.SrcFactor
	ϟc.Blending.DstRgbBlendFactor = ϟa.DstFactor
	ϟc.Blending.DstAlphaBlendFactor = ϟa.DstFactor
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glBlendFunc expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlBlendFuncSeparate) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlBlendFuncSeparate{}
	ϟc.Blending.SrcRgbBlendFactor = ϟa.SrcFactorRgb
	ϟc.Blending.DstRgbBlendFactor = ϟa.DstFactorRgb
	ϟc.Blending.SrcAlphaBlendFactor = ϟa.SrcFactorAlpha
	ϟc.Blending.DstAlphaBlendFactor = ϟa.DstFactorAlpha
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glBlendFuncSeparate expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlBlendEquation) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlBlendEquation{}
	ϟc.Blending.BlendEquationRgb = ϟa.Equation
	ϟc.Blending.BlendEquationAlpha = ϟa.Equation
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glBlendEquation expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlBlendEquationSeparate) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlBlendEquationSeparate{}
	ϟc.Blending.BlendEquationRgb = ϟa.Rgb
	ϟc.Blending.BlendEquationAlpha = ϟa.Alpha
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glBlendEquationSeparate expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlBlendColor) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlBlendColor{}
	ϟc.Blending.BlendColor = func() Color {
		s := Color{}
		s.Init()
		s.Red = ϟa.Red
		s.Green = ϟa.Green
		s.Blue = ϟa.Blue
		s.Alpha = ϟa.Alpha
		return s
	}()
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glBlendColor expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlEnableVertexAttribArray) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlEnableVertexAttribArray{}
	a := ϟc.VertexAttributeArrays.Get(ϟa.Location) // VertexAttributeArrayPtr
	a.Enabled = true
	_ = a
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glEnableVertexAttribArray expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDisableVertexAttribArray) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDisableVertexAttribArray{}
	a := ϟc.VertexAttributeArrays.Get(ϟa.Location) // VertexAttributeArrayPtr
	a.Enabled = false
	_ = a
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDisableVertexAttribArray expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlVertexAttribPointer) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlVertexAttribPointer{}
	a := ϟc.VertexAttributeArrays.Get(ϟa.Location) // VertexAttributeArrayPtr
	a.Size = uint32(ϟa.Size)
	a.Type = ϟa.Type
	a.Normalized = ϟa.Normalized
	a.Stride = ϟa.Stride
	a.Pointer = memory.Pointer(ϟa.Data)
	a.Buffer = ϟc.BoundBuffers.Get(BufferTarget_GL_ARRAY_BUFFER)
	_ = a
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glVertexAttribPointer expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetActiveAttrib) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
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
func (ϟa *GlGetActiveUniform) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
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
func (ϟa *GlGetError) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetError{}
	ϟa.Result = ϟa.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetError expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetProgramiv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetProgramiv{}
	ϟo.Value = make(S32Array, int32(1))
	ϟa.Value = ϟa.Value
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetProgramiv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetShaderiv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetShaderiv{}
	ϟo.Value = make(S32Array, int32(1))
	s := ϟc.Instances.Shaders.Get(ϟa.Shader) // ShaderPtr
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
	_ = s
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetShaderiv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetUniformLocation) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetUniformLocation{}
	ϟa.Result = ϟa.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetUniformLocation expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetAttribLocation) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetAttribLocation{}
	ϟa.Result = ϟa.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetAttribLocation expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlPixelStorei) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlPixelStorei{}
	ϟc.PixelStorage[ϟa.Parameter] = ϟa.Value
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glPixelStorei expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlTexParameteri) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlTexParameteri{}
	id := ϟc.TextureUnits.Get(ϟc.ActiveTextureUnit).Get(ϟa.Target) // TextureId
	t := ϟc.Instances.Textures.Get(id)                             // TexturePtr
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
	_, _ = id, t
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glTexParameteri expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlTexParameterf) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlTexParameterf{}
	id := ϟc.TextureUnits.Get(ϟc.ActiveTextureUnit).Get(ϟa.Target) // TextureId
	t := ϟc.Instances.Textures.Get(id)                             // TexturePtr
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
	_, _ = id, t
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glTexParameterf expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetTexParameteriv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetTexParameteriv{}
	ϟo.Values = make(S32Array, int32(1))
	id := ϟc.TextureUnits.Get(ϟc.ActiveTextureUnit).Get(ϟa.Target) // TextureId
	t := ϟc.Instances.Textures.Get(id)                             // TexturePtr
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
	_, _ = id, t
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetTexParameteriv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetTexParameterfv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetTexParameterfv{}
	ϟo.Values = make(F32Array, int32(1))
	id := ϟc.TextureUnits.Get(ϟc.ActiveTextureUnit).Get(ϟa.Target) // TextureId
	t := ϟc.Instances.Textures.Get(id)                             // TexturePtr
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
	_, _ = id, t
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetTexParameterfv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniform1i) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniform1i{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramPtr
	uniform := program.Uniforms.Get(ϟa.Location)          // Uniform
	uniform.Type = ShaderUniformType_GL_INT
	uniform.Value.S32 = ϟa.Value
	program.Uniforms[ϟa.Location] = uniform
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniform1i expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniform2i) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniform2i{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramPtr
	uniform := program.Uniforms.Get(ϟa.Location)          // Uniform
	uniform.Type = ShaderUniformType_GL_INT_VEC2
	uniform.Value.Vec2i = func() Vec2i {
		s := Vec2i{}
		s.Init()
		s.X = ϟa.Value0
		s.Y = ϟa.Value1
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniform2i expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniform3i) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniform3i{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramPtr
	uniform := program.Uniforms.Get(ϟa.Location)          // Uniform
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
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniform3i expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniform4i) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniform4i{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramPtr
	uniform := program.Uniforms.Get(ϟa.Location)          // Uniform
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
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniform4i expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniform1iv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniform1iv{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramPtr
	uniform := program.Uniforms.Get(ϟa.Location)          // Uniform
	uniform.Type = ShaderUniformType_GL_INT
	uniform.Value.S32 = ϟa.Value[int32(0)]
	program.Uniforms[ϟa.Location] = uniform
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniform1iv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniform2iv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniform2iv{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramPtr
	uniform := program.Uniforms.Get(ϟa.Location)          // Uniform
	uniform.Type = ShaderUniformType_GL_INT_VEC2
	uniform.Value.Vec2i = func() Vec2i {
		s := Vec2i{}
		s.Init()
		s.X = ϟa.Value[int32(0)]
		s.Y = ϟa.Value[int32(1)]
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniform2iv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniform3iv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniform3iv{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramPtr
	uniform := program.Uniforms.Get(ϟa.Location)          // Uniform
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
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniform3iv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniform4iv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniform4iv{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramPtr
	uniform := program.Uniforms.Get(ϟa.Location)          // Uniform
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
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniform4iv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniform1f) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniform1f{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramPtr
	uniform := program.Uniforms.Get(ϟa.Location)          // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT
	uniform.Value.F32 = ϟa.Value
	program.Uniforms[ϟa.Location] = uniform
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniform1f expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniform2f) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniform2f{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramPtr
	uniform := program.Uniforms.Get(ϟa.Location)          // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT_VEC2
	uniform.Value.Vec2f = func() Vec2f {
		s := Vec2f{}
		s.Init()
		s.X = ϟa.Value0
		s.Y = ϟa.Value1
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniform2f expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniform3f) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniform3f{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramPtr
	uniform := program.Uniforms.Get(ϟa.Location)          // Uniform
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
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniform3f expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniform4f) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniform4f{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramPtr
	uniform := program.Uniforms.Get(ϟa.Location)          // Uniform
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
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniform4f expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniform1fv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniform1fv{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramPtr
	uniform := program.Uniforms.Get(ϟa.Location)          // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT
	uniform.Value.F32 = ϟa.Value[int32(0)]
	program.Uniforms[ϟa.Location] = uniform
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniform1fv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniform2fv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniform2fv{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramPtr
	uniform := program.Uniforms.Get(ϟa.Location)          // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT_VEC2
	uniform.Value.Vec2f = func() Vec2f {
		s := Vec2f{}
		s.Init()
		s.X = ϟa.Value[int32(0)]
		s.Y = ϟa.Value[int32(1)]
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniform2fv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniform3fv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniform3fv{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramPtr
	uniform := program.Uniforms.Get(ϟa.Location)          // Uniform
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
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniform3fv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniform4fv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniform4fv{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramPtr
	uniform := program.Uniforms.Get(ϟa.Location)          // Uniform
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
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniform4fv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniformMatrix2fv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniformMatrix2fv{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramPtr
	uniform := program.Uniforms.Get(ϟa.Location)          // Uniform
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
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniformMatrix2fv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniformMatrix3fv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniformMatrix3fv{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramPtr
	uniform := program.Uniforms.Get(ϟa.Location)          // Uniform
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
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniformMatrix3fv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUniformMatrix4fv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniformMatrix4fv{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramPtr
	uniform := program.Uniforms.Get(ϟa.Location)          // Uniform
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
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUniformMatrix4fv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetUniformfv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetUniformfv{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetUniformfv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetUniformiv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetUniformiv{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetUniformiv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlVertexAttrib1f) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlVertexAttrib1f{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glVertexAttrib1f expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlVertexAttrib2f) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlVertexAttrib2f{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glVertexAttrib2f expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlVertexAttrib3f) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlVertexAttrib3f{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glVertexAttrib3f expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlVertexAttrib4f) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlVertexAttrib4f{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glVertexAttrib4f expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlVertexAttrib1fv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlVertexAttrib1fv{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glVertexAttrib1fv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlVertexAttrib2fv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlVertexAttrib2fv{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glVertexAttrib2fv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlVertexAttrib3fv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlVertexAttrib3fv{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glVertexAttrib3fv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlVertexAttrib4fv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlVertexAttrib4fv{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glVertexAttrib4fv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetShaderPrecisionFormat) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetShaderPrecisionFormat{}
	ϟo.Range = make(S32Array, int32(2))
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetShaderPrecisionFormat expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDepthMask) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDepthMask{}
	ϟc.Rasterizing.DepthMask = ϟa.Enabled
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDepthMask expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDepthFunc) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDepthFunc{}
	ϟc.Rasterizing.DepthTestFunction = ϟa.Function
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDepthFunc expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDepthRangef) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDepthRangef{}
	ϟc.Rasterizing.DepthNear = ϟa.Near
	ϟc.Rasterizing.DepthFar = ϟa.Far
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDepthRangef expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlColorMask) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlColorMask{}
	ϟc.Rasterizing.ColorMaskRed = ϟa.Red
	ϟc.Rasterizing.ColorMaskGreen = ϟa.Green
	ϟc.Rasterizing.ColorMaskBlue = ϟa.Blue
	ϟc.Rasterizing.ColorMaskAlpha = ϟa.Alpha
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glColorMask expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlStencilMask) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlStencilMask{}
	ϟc.Rasterizing.StencilMask[FaceMode_GL_FRONT] = ϟa.Mask
	ϟc.Rasterizing.StencilMask[FaceMode_GL_BACK] = ϟa.Mask
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glStencilMask expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlStencilMaskSeparate) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlStencilMaskSeparate{}
	switch ϟa.Face {
	case FaceMode_GL_FRONT:
		ϟc.Rasterizing.StencilMask[FaceMode_GL_FRONT] = ϟa.Mask
	case FaceMode_GL_BACK:
		ϟc.Rasterizing.StencilMask[FaceMode_GL_BACK] = ϟa.Mask
	case FaceMode_GL_FRONT_AND_BACK:
		ϟc.Rasterizing.StencilMask[FaceMode_GL_FRONT] = ϟa.Mask
		ϟc.Rasterizing.StencilMask[FaceMode_GL_BACK] = ϟa.Mask
	default:
		// TODO: better unmatched handling
		v := ϟa.Face
		log.Printf("Error: Missing switch case handler for value %T %v", v, v)
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glStencilMaskSeparate expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlStencilFuncSeparate) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlStencilFuncSeparate{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glStencilFuncSeparate expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlStencilOpSeparate) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlStencilOpSeparate{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glStencilOpSeparate expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlFrontFace) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlFrontFace{}
	ϟc.Rasterizing.FrontFace = ϟa.Orientation
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glFrontFace expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlViewport) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlViewport{}
	ϟc.Rasterizing.Viewport = func() Rect {
		s := Rect{}
		s.Init()
		s.X = ϟa.X
		s.Y = ϟa.Y
		s.Width = ϟa.Width
		s.Height = ϟa.Height
		return s
	}()
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glViewport expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlScissor) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlScissor{}
	ϟc.Rasterizing.Scissor = func() Rect {
		s := Rect{}
		s.Init()
		s.X = ϟa.X
		s.Y = ϟa.Y
		s.Width = ϟa.Width
		s.Height = ϟa.Height
		return s
	}()
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glScissor expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlActiveTexture) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlActiveTexture{}
	ϟc.ActiveTextureUnit = ϟa.Unit
	if !(ϟc.TextureUnits.Contains(ϟa.Unit)) {
		ϟc.TextureUnits[ϟa.Unit] = ϟc.TextureUnits.Get(ϟa.Unit)
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glActiveTexture expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGenTextures) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGenTextures{}
	ϟo.Textures = make(TextureIdArray, ϟa.Count)
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		id := TextureId(ϟa.Textures[i]) // TextureId
		ϟc.Instances.Textures[id] = func() *Texture {
			s := &Texture{}
			s.Init()
			return s
		}()
		ϟa.Textures[i] = id
		_ = id
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGenTextures expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDeleteTextures) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDeleteTextures{}
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		ϟc.Instances.Textures.Delete(ϟa.Textures[i])
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDeleteTextures expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlIsTexture) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlIsTexture{}
	ϟa.Result = ϟc.Instances.Textures.Contains(ϟa.Texture)
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glIsTexture expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlBindTexture) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlBindTexture{}
	if !(ϟc.Instances.Textures.Contains(ϟa.Texture)) {
		ϟc.Instances.Textures[ϟa.Texture] = func() *Texture {
			s := &Texture{}
			s.Init()
			return s
		}()
	}
	ϟc.TextureUnits.Get(ϟc.ActiveTextureUnit)[ϟa.Target] = ϟa.Texture
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glBindTexture expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlTexImage2D) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlTexImage2D{}
	switch ϟa.Target {
	case TextureImageTarget_GL_TEXTURE_2D:
		id := ϟc.TextureUnits.Get(ϟc.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_2D) // TextureId
		t := ϟc.Instances.Textures.Get(id)                                               // TexturePtr
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.Width
			s.Height = ϟa.Height
			s.Size = imageSize(uint32(ϟa.Width), uint32(ϟa.Height), ϟa.Format, ϟa.Type)
			s.Format = ImageTexelFormat(ϟa.Format)
			return s
		}() // Image
		if (ϟa.Data) != (TexturePointer(memory.Pointer(0))) {
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
		id := ϟc.TextureUnits.Get(ϟc.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_CUBE_MAP) // TextureId
		t := ϟc.Instances.Textures.Get(id)                                                     // TexturePtr
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.Width
			s.Height = ϟa.Height
			s.Size = imageSize(uint32(ϟa.Width), uint32(ϟa.Height), ϟa.Format, ϟa.Type)
			s.Format = ImageTexelFormat(ϟa.Format)
			return s
		}() // Image
		if (ϟa.Data) != (TexturePointer(memory.Pointer(0))) {
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
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glTexImage2D expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlTexSubImage2D) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlTexSubImage2D{}
	switch ϟa.Target {
	case TextureImageTarget_GL_TEXTURE_2D:
		id := ϟc.TextureUnits.Get(ϟc.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_2D) // TextureId
		t := ϟc.Instances.Textures.Get(id)                                               // TexturePtr
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.Width
			s.Height = ϟa.Height
			s.Size = imageSize(uint32(ϟa.Width), uint32(ϟa.Height), ϟa.Format, ϟa.Type)
			s.Format = ImageTexelFormat(ϟa.Format)
			return s
		}() // Image
		l.Data.Write(ϟs.Memory.Slice(memory.Range{
			Base: memory.Pointer(ϟa.Data),
			Size: uint64(l.Size),
		}))
		t.Texture2D[ϟa.Level] = l
		t.Kind = TextureKind_TEXTURE2D
		t.Format = ImageTexelFormat(ϟa.Format)
		_, _, _ = id, t, l
	case TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X, TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Z, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z:
		id := ϟc.TextureUnits.Get(ϟc.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_CUBE_MAP) // TextureId
		t := ϟc.Instances.Textures.Get(id)                                                     // TexturePtr
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.Width
			s.Height = ϟa.Height
			s.Size = imageSize(uint32(ϟa.Width), uint32(ϟa.Height), ϟa.Format, ϟa.Type)
			s.Format = ImageTexelFormat(ϟa.Format)
			return s
		}() // Image
		l.Data.Write(ϟs.Memory.Slice(memory.Range{
			Base: memory.Pointer(ϟa.Data),
			Size: uint64(l.Size),
		}))
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
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glTexSubImage2D expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlCopyTexImage2D) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlCopyTexImage2D{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glCopyTexImage2D expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlCopyTexSubImage2D) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlCopyTexSubImage2D{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glCopyTexSubImage2D expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlCompressedTexImage2D) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlCompressedTexImage2D{}
	switch ϟa.Target {
	case TextureImageTarget_GL_TEXTURE_2D:
		id := ϟc.TextureUnits.Get(ϟc.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_2D) // TextureId
		t := ϟc.Instances.Textures.Get(id)                                               // TexturePtr
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.Width
			s.Height = ϟa.Height
			s.Size = uint32(ϟa.ImageSize)
			s.Format = ImageTexelFormat(ϟa.Format)
			return s
		}() // Image
		l.Data.Write(ϟs.Memory.Slice(memory.Range{
			Base: memory.Pointer(ϟa.Data),
			Size: uint64(l.Size),
		}))
		t.Texture2D[ϟa.Level] = l
		t.Kind = TextureKind_TEXTURE2D
		t.Format = ImageTexelFormat(ϟa.Format)
		_, _, _ = id, t, l
	case TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X, TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Z, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z:
		id := ϟc.TextureUnits.Get(ϟc.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_CUBE_MAP) // TextureId
		t := ϟc.Instances.Textures.Get(id)                                                     // TexturePtr
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.Width
			s.Height = ϟa.Height
			s.Size = uint32(ϟa.ImageSize)
			s.Format = ImageTexelFormat(ϟa.Format)
			return s
		}() // Image
		l.Data.Write(ϟs.Memory.Slice(memory.Range{
			Base: memory.Pointer(ϟa.Data),
			Size: uint64(l.Size),
		}))
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
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glCompressedTexImage2D expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlCompressedTexSubImage2D) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlCompressedTexSubImage2D{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glCompressedTexSubImage2D expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGenerateMipmap) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGenerateMipmap{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGenerateMipmap expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlReadPixels) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlReadPixels{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glReadPixels expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGenFramebuffers) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGenFramebuffers{}
	ϟo.Framebuffers = make(FramebufferIdArray, ϟa.Count)
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		id := FramebufferId(ϟa.Framebuffers[i]) // FramebufferId
		ϟc.Instances.Framebuffers[id] = func() *Framebuffer {
			s := &Framebuffer{}
			s.Init()
			return s
		}()
		ϟa.Framebuffers[i] = id
		_ = id
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGenFramebuffers expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlBindFramebuffer) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlBindFramebuffer{}
	if !(ϟc.Instances.Framebuffers.Contains(ϟa.Framebuffer)) {
		ϟc.Instances.Framebuffers[ϟa.Framebuffer] = func() *Framebuffer {
			s := &Framebuffer{}
			s.Init()
			return s
		}()
	}
	if (ϟa.Target) == (FramebufferTarget_GL_FRAMEBUFFER) {
		ϟc.BoundFramebuffers[FramebufferTarget_GL_READ_FRAMEBUFFER] = ϟa.Framebuffer
		ϟc.BoundFramebuffers[FramebufferTarget_GL_DRAW_FRAMEBUFFER] = ϟa.Framebuffer
	} else {
		ϟc.BoundFramebuffers[ϟa.Target] = ϟa.Framebuffer
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glBindFramebuffer expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlCheckFramebufferStatus) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlCheckFramebufferStatus{}
	ϟa.Result = ϟa.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glCheckFramebufferStatus expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDeleteFramebuffers) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDeleteFramebuffers{}
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		ϟc.Instances.Framebuffers.Delete(ϟa.Framebuffers[i])
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDeleteFramebuffers expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlIsFramebuffer) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlIsFramebuffer{}
	ϟa.Result = ϟc.Instances.Framebuffers.Contains(ϟa.Framebuffer)
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glIsFramebuffer expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGenRenderbuffers) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGenRenderbuffers{}
	ϟo.Renderbuffers = make(RenderbufferIdArray, ϟa.Count)
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		id := RenderbufferId(ϟa.Renderbuffers[i]) // RenderbufferId
		ϟc.Instances.Renderbuffers[id] = func() *Renderbuffer {
			s := &Renderbuffer{}
			s.Init()
			return s
		}()
		ϟa.Renderbuffers[i] = id
		_ = id
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGenRenderbuffers expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlBindRenderbuffer) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlBindRenderbuffer{}
	if !(ϟc.Instances.Renderbuffers.Contains(ϟa.Renderbuffer)) {
		ϟc.Instances.Renderbuffers[ϟa.Renderbuffer] = func() *Renderbuffer {
			s := &Renderbuffer{}
			s.Init()
			return s
		}()
	}
	ϟc.BoundRenderbuffers[ϟa.Target] = ϟa.Renderbuffer
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glBindRenderbuffer expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlRenderbufferStorage) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlRenderbufferStorage{}
	id := ϟc.BoundRenderbuffers.Get(ϟa.Target) // RenderbufferId
	rb := ϟc.Instances.Renderbuffers.Get(id)   // RenderbufferPtr
	rb.Format = ϟa.Format
	rb.Width = ϟa.Width
	rb.Height = ϟa.Height
	_, _ = id, rb
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glRenderbufferStorage expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDeleteRenderbuffers) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDeleteRenderbuffers{}
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		ϟc.Instances.Renderbuffers.Delete(ϟa.Renderbuffers[i])
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDeleteRenderbuffers expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlIsRenderbuffer) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlIsRenderbuffer{}
	ϟa.Result = ϟc.Instances.Renderbuffers.Contains(ϟa.Renderbuffer)
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glIsRenderbuffer expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetRenderbufferParameteriv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetRenderbufferParameteriv{}
	ϟo.Values = make(S32Array, int32(1))
	id := ϟc.BoundRenderbuffers.Get(ϟa.Target) // RenderbufferId
	rb := ϟc.Instances.Renderbuffers.Get(id)   // RenderbufferPtr
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
	_, _ = id, rb
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetRenderbufferParameteriv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGenBuffers) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGenBuffers{}
	ϟo.Buffers = make(BufferIdArray, ϟa.Count)
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		id := BufferId(ϟa.Buffers[i]) // BufferId
		ϟc.Instances.Buffers[id] = func() *Buffer {
			s := &Buffer{}
			s.Init()
			return s
		}()
		ϟa.Buffers[i] = id
		_ = id
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGenBuffers expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlBindBuffer) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlBindBuffer{}
	if !(ϟc.Instances.Buffers.Contains(ϟa.Buffer)) {
		ϟc.Instances.Buffers[ϟa.Buffer] = func() *Buffer {
			s := &Buffer{}
			s.Init()
			return s
		}()
	}
	ϟc.BoundBuffers[ϟa.Target] = ϟa.Buffer
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glBindBuffer expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlBufferData) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlBufferData{}
	id := ϟc.BoundBuffers.Get(ϟa.Target) // BufferId
	b := ϟc.Instances.Buffers.Get(id)    // BufferPtr
	if (ϟa.Data) != (BufferDataPointer(memory.Pointer(0))) {
		b.Data.Write(ϟs.Memory.Slice(memory.Range{
			Base: memory.Pointer(ϟa.Data),
			Size: uint64(uint32(ϟa.Size)),
		}))
	}
	b.Size = ϟa.Size
	b.Usage = ϟa.Usage
	_, _ = id, b
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glBufferData expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlBufferSubData) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlBufferSubData{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glBufferSubData expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDeleteBuffers) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDeleteBuffers{}
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		ϟc.Instances.Buffers.Delete(ϟa.Buffers[i])
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDeleteBuffers expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlIsBuffer) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlIsBuffer{}
	ϟa.Result = ϟc.Instances.Buffers.Contains(ϟa.Buffer)
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glIsBuffer expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetBufferParameteriv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetBufferParameteriv{}
	id := ϟc.BoundBuffers.Get(ϟa.Target) // BufferId
	b := ϟc.Instances.Buffers.Get(id)    // BufferPtr
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
	_, _ = id, b
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetBufferParameteriv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlCreateShader) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlCreateShader{}
	id := ShaderId(ϟa.Result) // ShaderId
	ϟc.Instances.Shaders[id] = func() *Shader {
		s := &Shader{}
		s.Init()
		return s
	}()
	s := ϟc.Instances.Shaders.Get(id) // ShaderPtr
	s.Type = ϟa.Type
	ϟa.Result = id
	_, _ = id, s
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glCreateShader expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDeleteShader) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDeleteShader{}
	s := ϟc.Instances.Shaders.Get(ϟa.Shader) // ShaderPtr
	s.Deletable = true
	ϟc.Instances.Shaders.Delete(ϟa.Shader)
	_ = s
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDeleteShader expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlShaderSource) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlShaderSource{}
	s := ϟc.Instances.Shaders.Get(ϟa.Shader) // ShaderPtr
	s.Source = ϟa.Source
	_ = s
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glShaderSource expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlShaderBinary) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlShaderBinary{}
	for i := int32(int32(0)); i < ϟa.Count; i++ {
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glShaderBinary expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetShaderInfoLog) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetShaderInfoLog{}
	s := ϟc.Instances.Shaders.Get(ϟa.Shader) // ShaderPtr
	min_0_a := ϟa.BufferLength               // s32
	min_0_b := strlen(s.InfoLog)             // s32
	min_0_result := func() (result int32) {
		switch (min_0_a) < (min_0_b) {
		case true:
			return min_0_a
		case false:
			return min_0_b
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (min_0_a) < (min_0_b), ϟa))
			return result
		}
	}() // s32
	ϟa.StringLengthWritten = min_0_result
	ϟa.Info = substr(s.InfoLog, int32(0), ϟa.StringLengthWritten)
	_, _, _, _ = s, min_0_a, min_0_b, min_0_result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetShaderInfoLog expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetShaderSource) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetShaderSource{}
	s := ϟc.Instances.Shaders.Get(ϟa.Shader) // ShaderPtr
	min_1_a := ϟa.BufferLength               // s32
	min_1_b := strlen(s.Source[int32(0)])    // s32
	min_1_result := func() (result int32) {
		switch (min_1_a) < (min_1_b) {
		case true:
			return min_1_a
		case false:
			return min_1_b
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (min_1_a) < (min_1_b), ϟa))
			return result
		}
	}() // s32
	ϟa.StringLengthWritten = min_1_result
	ϟa.Source = substr(s.Source[int32(0)], int32(0), ϟa.StringLengthWritten)
	_, _, _, _ = s, min_1_a, min_1_b, min_1_result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetShaderSource expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlReleaseShaderCompiler) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlReleaseShaderCompiler{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glReleaseShaderCompiler expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlCompileShader) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlCompileShader{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glCompileShader expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlIsShader) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlIsShader{}
	ϟa.Result = ϟc.Instances.Shaders.Contains(ϟa.Shader)
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glIsShader expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlCreateProgram) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlCreateProgram{}
	id := ProgramId(ϟa.Result) // ProgramId
	ϟc.Instances.Programs[id] = func() *Program {
		s := &Program{}
		s.Init()
		return s
	}()
	ϟa.Result = id
	_ = id
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glCreateProgram expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDeleteProgram) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDeleteProgram{}
	ϟc.Instances.Programs.Delete(ϟa.Program)
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDeleteProgram expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlAttachShader) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlAttachShader{}
	p := ϟc.Instances.Programs.Get(ϟa.Program) // ProgramPtr
	s := ϟc.Instances.Shaders.Get(ϟa.Shader)   // ShaderPtr
	p.Shaders[s.Type] = ϟa.Shader
	_, _ = p, s
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glAttachShader expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDetachShader) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDetachShader{}
	p := ϟc.Instances.Programs.Get(ϟa.Program) // ProgramPtr
	s := ϟc.Instances.Shaders.Get(ϟa.Shader)   // ShaderPtr
	p.Shaders.Delete(s.Type)
	_, _ = p, s
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDetachShader expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetAttachedShaders) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetAttachedShaders{}
	ϟo.Shaders = make(ShaderIdArray, ϟa.BufferLength)
	p := ϟc.Instances.Programs.Get(ϟa.Program) // ProgramPtr
	min_2_a := ϟa.BufferLength                 // s32
	min_2_b := int32(len(p.Shaders))           // s32
	min_2_result := func() (result int32) {
		switch (min_2_a) < (min_2_b) {
		case true:
			return min_2_a
		case false:
			return min_2_b
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (min_2_a) < (min_2_b), ϟa))
			return result
		}
	}() // s32
	ϟa.ShadersLengthWritten = min_2_result
	ϟa.Shaders = p.Shaders.Range()
	_, _, _, _ = p, min_2_a, min_2_b, min_2_result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetAttachedShaders expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlLinkProgram) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlLinkProgram{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glLinkProgram expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetProgramInfoLog) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetProgramInfoLog{}
	p := ϟc.Instances.Programs.Get(ϟa.Program) // ProgramPtr
	min_3_a := ϟa.BufferLength                 // s32
	min_3_b := strlen(p.InfoLog)               // s32
	min_3_result := func() (result int32) {
		switch (min_3_a) < (min_3_b) {
		case true:
			return min_3_a
		case false:
			return min_3_b
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (min_3_a) < (min_3_b), ϟa))
			return result
		}
	}() // s32
	ϟa.StringLengthWritten = min_3_result
	ϟa.Info = substr(p.InfoLog, int32(0), ϟa.StringLengthWritten)
	_, _, _, _ = p, min_3_a, min_3_b, min_3_result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetProgramInfoLog expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUseProgram) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUseProgram{}
	ϟc.BoundProgram = ϟa.Program
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUseProgram expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlIsProgram) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlIsProgram{}
	ϟa.Result = ϟc.Instances.Programs.Contains(ϟa.Program)
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glIsProgram expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlValidateProgram) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlValidateProgram{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glValidateProgram expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlClearColor) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlClearColor{}
	ϟc.Clearing.ClearColor = func() Color {
		s := Color{}
		s.Init()
		s.Red = ϟa.R
		s.Green = ϟa.G
		s.Blue = ϟa.B
		s.Alpha = ϟa.A
		return s
	}()
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glClearColor expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlClearDepthf) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlClearDepthf{}
	ϟc.Clearing.ClearDepth = ϟa.Depth
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glClearDepthf expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlClearStencil) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlClearStencil{}
	ϟc.Clearing.ClearStencil = ϟa.Stencil
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glClearStencil expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlClear) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlClear{}
	if (ClearMask_GL_COLOR_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glClear expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlCullFace) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlCullFace{}
	ϟc.Rasterizing.CullFace = ϟa.Mode
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glCullFace expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlPolygonOffset) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlPolygonOffset{}
	ϟc.Rasterizing.PolygonOffsetUnits = ϟa.Units
	ϟc.Rasterizing.PolygonOffsetFactor = ϟa.ScaleFactor
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glPolygonOffset expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlLineWidth) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlLineWidth{}
	ϟc.Rasterizing.LineWidth = ϟa.Width
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glLineWidth expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlSampleCoverage) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlSampleCoverage{}
	ϟc.Rasterizing.SampleCoverageValue = ϟa.Value
	ϟc.Rasterizing.SampleCoverageInvert = ϟa.Invert
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glSampleCoverage expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlHint) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlHint{}
	ϟc.GenerateMipmapHint = ϟa.Mode
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glHint expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlFramebufferRenderbuffer) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlFramebufferRenderbuffer{}
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
	framebufferId := ϟc.BoundFramebuffers.Get(target)                   // FramebufferId
	framebuffer := ϟc.Instances.Framebuffers.Get(framebufferId)         // FramebufferPtr
	attachment := framebuffer.Attachments.Get(ϟa.FramebufferAttachment) // FramebufferAttachmentInfo
	if (ϟa.Renderbuffer) == (ϟc.Internals.NilRenderbuffer) {
		attachment.Type = FramebufferAttachmentType_GL_NONE
	} else {
		attachment.Type = FramebufferAttachmentType_GL_RENDERBUFFER
	}
	attachment.Object = uint32(ϟa.Renderbuffer)
	attachment.TextureLevel = int32(0)
	attachment.CubeMapFace = CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X
	framebuffer.Attachments[ϟa.FramebufferAttachment] = attachment
	_, _, _, _ = target, framebufferId, framebuffer, attachment
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glFramebufferRenderbuffer expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlFramebufferTexture2D) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlFramebufferTexture2D{}
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
	framebufferId := ϟc.BoundFramebuffers.Get(target)                   // FramebufferId
	framebuffer := ϟc.Instances.Framebuffers.Get(framebufferId)         // FramebufferPtr
	attachment := framebuffer.Attachments.Get(ϟa.FramebufferAttachment) // FramebufferAttachmentInfo
	if (ϟa.Texture) == (ϟc.Internals.NilTexture) {
		attachment.Type = FramebufferAttachmentType_GL_NONE
		attachment.Object = uint32(ϟc.Internals.NilTexture)
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
	_, _, _, _ = target, framebufferId, framebuffer, attachment
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glFramebufferTexture2D expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetFramebufferAttachmentParameteriv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetFramebufferAttachmentParameteriv{}
	ϟo.Value = make(S32Array, int32(1))
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
	framebufferId := ϟc.BoundFramebuffers.Get(target)           // FramebufferId
	framebuffer := ϟc.Instances.Framebuffers.Get(framebufferId) // FramebufferPtr
	a := framebuffer.Attachments.Get(ϟa.Attachment)             // FramebufferAttachmentInfo
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
	_, _, _, _ = target, framebufferId, framebuffer, a
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetFramebufferAttachmentParameteriv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDrawElements) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDrawElements{}
	id := ϟc.BoundBuffers.Get(BufferTarget_GL_ELEMENT_ARRAY_BUFFER) // BufferId
	index_data := func() (result memory.Pointer) {
		switch (id) != (ϟc.Internals.NilBuffer) {
		case true:
			return memoryOffset(ϟc.Instances.Buffers.Get(id).Data, uint64(ϟa.Indices))
		case false:
			return memory.Pointer(ϟa.Indices)
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (id) != (ϟc.Internals.NilBuffer), ϟa))
			return result
		}
	}() // VoidArray
	IndexSize_4_indices_type := ϟa.IndicesType // IndicesType
	IndexSize_4_result := func() (result uint32) {
		switch IndexSize_4_indices_type {
		case IndicesType_GL_UNSIGNED_BYTE:
			return uint32(1)
		case IndicesType_GL_UNSIGNED_SHORT:
			return uint32(2)
		case IndicesType_GL_UNSIGNED_INT:
			return uint32(4)
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", IndexSize_4_indices_type, ϟa))
			return result
		}
	}() // u32
	read(index_data, uint32(0), (uint32(ϟa.ElementCount))*(IndexSize_4_result))
	first := minIndex(index_data, ϟa.IndicesType, uint32(ϟa.ElementCount)) // u32
	last := maxIndex(index_data, ϟa.IndicesType, uint32(ϟa.ElementCount))  // u32
	ReadVertexArrays_5_first_index := first                                // u32
	ReadVertexArrays_5_index_count := (last) - (first)                     // u32
	for i := int32(int32(0)); i < int32(len(ϟc.VertexAttributeArrays)); i++ {
		arr := ϟc.VertexAttributeArrays.Get(AttributeLocation(i)) // VertexAttributeArrayPtr
		if arr.Enabled {
			vertexAttribTypeSize_6_t := arr.Type // VertexAttribType
			vertexAttribTypeSize_6_result := func() (result uint32) {
				switch vertexAttribTypeSize_6_t {
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
					panic(fmt.Errorf("Unmatched switch(%v) in atom %T", vertexAttribTypeSize_6_t, ϟa))
					return result
				}
			}() // u32
			elsize := (vertexAttribTypeSize_6_result) * (arr.Size) // u32
			size := (elsize) * (ReadVertexArrays_5_index_count)    // u32
			offset := (elsize) * (ReadVertexArrays_5_first_index)  // u32
			if (arr.Buffer) == (BufferId(uint32(0))) {
				read(arr.Pointer, offset, size)
			}
			_, _, _, _, _ = vertexAttribTypeSize_6_t, vertexAttribTypeSize_6_result, elsize, size, offset
		}
		_ = arr
	}
	_, _, _, _, _, _, _, _ = id, index_data, IndexSize_4_indices_type, IndexSize_4_result, first, last, ReadVertexArrays_5_first_index, ReadVertexArrays_5_index_count
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDrawElements expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDrawArrays) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDrawArrays{}
	ReadVertexArrays_7_first_index := uint32(ϟa.FirstIndex) // u32
	ReadVertexArrays_7_index_count := uint32(ϟa.IndexCount) // u32
	for i := int32(int32(0)); i < int32(len(ϟc.VertexAttributeArrays)); i++ {
		arr := ϟc.VertexAttributeArrays.Get(AttributeLocation(i)) // VertexAttributeArrayPtr
		if arr.Enabled {
			vertexAttribTypeSize_8_t := arr.Type // VertexAttribType
			vertexAttribTypeSize_8_result := func() (result uint32) {
				switch vertexAttribTypeSize_8_t {
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
					panic(fmt.Errorf("Unmatched switch(%v) in atom %T", vertexAttribTypeSize_8_t, ϟa))
					return result
				}
			}() // u32
			elsize := (vertexAttribTypeSize_8_result) * (arr.Size) // u32
			size := (elsize) * (ReadVertexArrays_7_index_count)    // u32
			offset := (elsize) * (ReadVertexArrays_7_first_index)  // u32
			if (arr.Buffer) == (BufferId(uint32(0))) {
				read(arr.Pointer, offset, size)
			}
			_, _, _, _, _ = vertexAttribTypeSize_8_t, vertexAttribTypeSize_8_result, elsize, size, offset
		}
		_ = arr
	}
	_, _ = ReadVertexArrays_7_first_index, ReadVertexArrays_7_index_count
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDrawArrays expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlFlush) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlFlush{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glFlush expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlFinish) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlFinish{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glFinish expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetBooleanv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetBooleanv{}
	ϟo.Values = make(BoolArray, stateVariableSize(ϟa.Param))
	switch ϟa.Param {
	case StateVariable_GL_BLEND:
		ϟa.Values[int32(0)] = ϟc.Capabilities.Get(Capability_GL_BLEND)
	case StateVariable_GL_CULL_FACE:
		ϟa.Values[int32(0)] = ϟc.Capabilities.Get(Capability_GL_CULL_FACE)
	case StateVariable_GL_DEPTH_TEST:
		ϟa.Values[int32(0)] = ϟc.Capabilities.Get(Capability_GL_DEPTH_TEST)
	case StateVariable_GL_DITHER:
		ϟa.Values[int32(0)] = ϟc.Capabilities.Get(Capability_GL_DITHER)
	case StateVariable_GL_POLYGON_OFFSET_FILL:
		ϟa.Values[int32(0)] = ϟc.Capabilities.Get(Capability_GL_POLYGON_OFFSET_FILL)
	case StateVariable_GL_SAMPLE_ALPHA_TO_COVERAGE:
		ϟa.Values[int32(0)] = ϟc.Capabilities.Get(Capability_GL_SAMPLE_ALPHA_TO_COVERAGE)
	case StateVariable_GL_SAMPLE_COVERAGE:
		ϟa.Values[int32(0)] = ϟc.Capabilities.Get(Capability_GL_SAMPLE_COVERAGE)
	case StateVariable_GL_SCISSOR_TEST:
		ϟa.Values[int32(0)] = ϟc.Capabilities.Get(Capability_GL_SCISSOR_TEST)
	case StateVariable_GL_STENCIL_TEST:
		ϟa.Values[int32(0)] = ϟc.Capabilities.Get(Capability_GL_STENCIL_TEST)
	case StateVariable_GL_DEPTH_WRITEMASK:
		ϟa.Values[int32(0)] = ϟc.Rasterizing.DepthMask
	case StateVariable_GL_COLOR_WRITEMASK:
		ϟa.Values[int32(0)] = ϟc.Rasterizing.ColorMaskRed
		ϟa.Values[int32(1)] = ϟc.Rasterizing.ColorMaskGreen
		ϟa.Values[int32(2)] = ϟc.Rasterizing.ColorMaskBlue
		ϟa.Values[int32(3)] = ϟc.Rasterizing.ColorMaskAlpha
	case StateVariable_GL_SAMPLE_COVERAGE_INVERT:
		ϟa.Values[int32(0)] = ϟc.Rasterizing.SampleCoverageInvert
	case StateVariable_GL_SHADER_COMPILER:
		ϟa.Values[int32(0)] = ϟa.Values[int32(0)]
	case StateVariable_GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT:
		ϟa.Values[int32(0)] = ϟa.Values[int32(0)]
	default:
		// TODO: better unmatched handling
		v := ϟa.Param
		log.Printf("Error: Missing switch case handler for value %T %v", v, v)
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetBooleanv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetFloatv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetFloatv{}
	ϟo.Values = make(F32Array, stateVariableSize(ϟa.Param))
	switch ϟa.Param {
	case StateVariable_GL_DEPTH_RANGE:
		ϟa.Values[int32(0)] = ϟc.Rasterizing.DepthNear
		ϟa.Values[int32(1)] = ϟc.Rasterizing.DepthFar
	case StateVariable_GL_LINE_WIDTH:
		ϟa.Values[int32(0)] = ϟc.Rasterizing.LineWidth
	case StateVariable_GL_POLYGON_OFFSET_FACTOR:
		ϟa.Values[int32(0)] = ϟc.Rasterizing.PolygonOffsetFactor
	case StateVariable_GL_POLYGON_OFFSET_UNITS:
		ϟa.Values[int32(0)] = ϟc.Rasterizing.PolygonOffsetUnits
	case StateVariable_GL_SAMPLE_COVERAGE_VALUE:
		ϟa.Values[int32(0)] = ϟc.Rasterizing.SampleCoverageValue
	case StateVariable_GL_COLOR_CLEAR_VALUE:
		ϟa.Values[int32(0)] = ϟc.Clearing.ClearColor.Red
		ϟa.Values[int32(1)] = ϟc.Clearing.ClearColor.Green
		ϟa.Values[int32(2)] = ϟc.Clearing.ClearColor.Blue
		ϟa.Values[int32(3)] = ϟc.Clearing.ClearColor.Alpha
	case StateVariable_GL_DEPTH_CLEAR_VALUE:
		ϟa.Values[int32(0)] = ϟc.Clearing.ClearDepth
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
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetFloatv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetIntegerv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetIntegerv{}
	ϟo.Values = make(S32Array, stateVariableSize(ϟa.Param))
	switch ϟa.Param {
	case StateVariable_GL_ACTIVE_TEXTURE:
		ϟa.Values[int32(0)] = int32(ϟc.ActiveTextureUnit)
	case StateVariable_GL_ARRAY_BUFFER_BINDING:
		ϟa.Values[int32(0)] = int32(ϟc.BoundBuffers.Get(BufferTarget_GL_ARRAY_BUFFER))
	case StateVariable_GL_ELEMENT_ARRAY_BUFFER_BINDING:
		ϟa.Values[int32(0)] = int32(ϟc.BoundBuffers.Get(BufferTarget_GL_ELEMENT_ARRAY_BUFFER))
	case StateVariable_GL_BLEND_SRC_ALPHA:
		ϟa.Values[int32(0)] = int32(ϟc.Blending.SrcAlphaBlendFactor)
	case StateVariable_GL_BLEND_SRC_RGB:
		ϟa.Values[int32(0)] = int32(ϟc.Blending.SrcRgbBlendFactor)
	case StateVariable_GL_BLEND_DST_ALPHA:
		ϟa.Values[int32(0)] = int32(ϟc.Blending.DstAlphaBlendFactor)
	case StateVariable_GL_BLEND_DST_RGB:
		ϟa.Values[int32(0)] = int32(ϟc.Blending.DstRgbBlendFactor)
	case StateVariable_GL_BLEND_EQUATION_RGB:
		ϟa.Values[int32(0)] = int32(ϟc.Blending.BlendEquationRgb)
	case StateVariable_GL_BLEND_EQUATION_ALPHA:
		ϟa.Values[int32(0)] = int32(ϟc.Blending.BlendEquationAlpha)
	case StateVariable_GL_BLEND_COLOR:
		ϟa.Values[int32(0)] = int32(ϟc.Blending.BlendColor.Red)
		ϟa.Values[int32(1)] = int32(ϟc.Blending.BlendColor.Green)
		ϟa.Values[int32(2)] = int32(ϟc.Blending.BlendColor.Blue)
		ϟa.Values[int32(3)] = int32(ϟc.Blending.BlendColor.Alpha)
	case StateVariable_GL_DEPTH_FUNC:
		ϟa.Values[int32(0)] = int32(ϟc.Rasterizing.DepthTestFunction)
	case StateVariable_GL_DEPTH_CLEAR_VALUE:
		ϟa.Values[int32(0)] = int32(ϟc.Clearing.ClearDepth)
	case StateVariable_GL_STENCIL_WRITEMASK:
		ϟa.Values[int32(0)] = int32(ϟc.Rasterizing.StencilMask.Get(FaceMode_GL_FRONT))
	case StateVariable_GL_STENCIL_BACK_WRITEMASK:
		ϟa.Values[int32(0)] = int32(ϟc.Rasterizing.StencilMask.Get(FaceMode_GL_BACK))
	case StateVariable_GL_VIEWPORT:
		ϟa.Values[int32(0)] = ϟc.Rasterizing.Viewport.X
		ϟa.Values[int32(1)] = ϟc.Rasterizing.Viewport.Y
		ϟa.Values[int32(2)] = ϟc.Rasterizing.Viewport.Width
		ϟa.Values[int32(3)] = ϟc.Rasterizing.Viewport.Height
	case StateVariable_GL_SCISSOR_BOX:
		ϟa.Values[int32(0)] = ϟc.Rasterizing.Scissor.X
		ϟa.Values[int32(1)] = ϟc.Rasterizing.Scissor.Y
		ϟa.Values[int32(2)] = ϟc.Rasterizing.Scissor.Width
		ϟa.Values[int32(3)] = ϟc.Rasterizing.Scissor.Height
	case StateVariable_GL_FRONT_FACE:
		ϟa.Values[int32(0)] = int32(ϟc.Rasterizing.FrontFace)
	case StateVariable_GL_CULL_FACE_MODE:
		ϟa.Values[int32(0)] = int32(ϟc.Rasterizing.CullFace)
	case StateVariable_GL_STENCIL_CLEAR_VALUE:
		ϟa.Values[int32(0)] = ϟc.Clearing.ClearStencil
	case StateVariable_GL_FRAMEBUFFER_BINDING:
		ϟa.Values[int32(0)] = int32(ϟc.BoundFramebuffers.Get(FramebufferTarget_GL_FRAMEBUFFER))
	case StateVariable_GL_READ_FRAMEBUFFER_BINDING:
		ϟa.Values[int32(0)] = int32(ϟc.BoundFramebuffers.Get(FramebufferTarget_GL_READ_FRAMEBUFFER))
	case StateVariable_GL_RENDERBUFFER_BINDING:
		ϟa.Values[int32(0)] = int32(ϟc.BoundRenderbuffers.Get(RenderbufferTarget_GL_RENDERBUFFER))
	case StateVariable_GL_CURRENT_PROGRAM:
		ϟa.Values[int32(0)] = int32(ϟc.BoundProgram)
	case StateVariable_GL_TEXTURE_BINDING_2D:
		ϟa.Values[int32(0)] = int32(ϟc.TextureUnits.Get(ϟc.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_2D))
	case StateVariable_GL_TEXTURE_BINDING_CUBE_MAP:
		ϟa.Values[int32(0)] = int32(ϟc.TextureUnits.Get(ϟc.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_CUBE_MAP))
	case StateVariable_GL_GENERATE_MIPMAP_HINT:
		ϟa.Values[int32(0)] = int32(ϟc.GenerateMipmapHint)
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
		ϟa.Values[int32(0)] = ϟc.PixelStorage.Get(PixelStoreParameter_GL_PACK_ALIGNMENT)
	case StateVariable_GL_UNPACK_ALIGNMENT:
		ϟa.Values[int32(0)] = ϟc.PixelStorage.Get(PixelStoreParameter_GL_UNPACK_ALIGNMENT)
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
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetIntegerv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetString) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetString{}
	ϟa.Result = ϟa.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetString expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlEnable) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlEnable{}
	ϟc.Capabilities[ϟa.Capability] = true
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glEnable expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDisable) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDisable{}
	ϟc.Capabilities[ϟa.Capability] = false
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDisable expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlIsEnabled) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlIsEnabled{}
	ϟa.Result = ϟc.Capabilities.Get(ϟa.Capability)
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glIsEnabled expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlMapBufferRange) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlMapBufferRange{}
	ϟa.Result = ϟa.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glMapBufferRange expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlUnmapBuffer) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUnmapBuffer{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glUnmapBuffer expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlInvalidateFramebuffer) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlInvalidateFramebuffer{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glInvalidateFramebuffer expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlRenderbufferStorageMultisample) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlRenderbufferStorageMultisample{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glRenderbufferStorageMultisample expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlBlitFramebuffer) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlBlitFramebuffer{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glBlitFramebuffer expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGenQueries) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGenQueries{}
	ϟo.Queries = make(QueryIdArray, ϟa.Count)
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		id := QueryId(ϟa.Queries[i]) // QueryId
		ϟc.Instances.Queries[id] = func() *Query {
			s := &Query{}
			s.Init()
			return s
		}()
		ϟa.Queries[i] = id
		_ = id
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGenQueries expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlBeginQuery) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlBeginQuery{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glBeginQuery expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlEndQuery) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlEndQuery{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glEndQuery expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDeleteQueries) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDeleteQueries{}
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		ϟc.Instances.Queries.Delete(ϟa.Queries[i])
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDeleteQueries expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlIsQuery) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlIsQuery{}
	ϟa.Result = ϟc.Instances.Queries.Contains(ϟa.Query)
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glIsQuery expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetQueryiv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetQueryiv{}
	ϟa.Value = ϟa.Value
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetQueryiv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetQueryObjectuiv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetQueryObjectuiv{}
	ϟa.Value = ϟa.Value
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetQueryObjectuiv expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGenQueriesEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGenQueriesEXT{}
	ϟo.Queries = make(QueryIdArray, ϟa.Count)
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		id := QueryId(ϟa.Queries[i]) // QueryId
		ϟc.Instances.Queries[id] = func() *Query {
			s := &Query{}
			s.Init()
			return s
		}()
		ϟa.Queries[i] = id
		_ = id
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGenQueriesEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlBeginQueryEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlBeginQueryEXT{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glBeginQueryEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlEndQueryEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlEndQueryEXT{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glEndQueryEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlDeleteQueriesEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDeleteQueriesEXT{}
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		ϟc.Instances.Queries.Delete(ϟa.Queries[i])
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glDeleteQueriesEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlIsQueryEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlIsQueryEXT{}
	ϟa.Result = ϟc.Instances.Queries.Contains(ϟa.Query)
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glIsQueryEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlQueryCounterEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlQueryCounterEXT{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glQueryCounterEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetQueryivEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetQueryivEXT{}
	ϟa.Value = ϟa.Value
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetQueryivEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetQueryObjectivEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetQueryObjectivEXT{}
	ϟa.Value = ϟa.Value
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetQueryObjectivEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetQueryObjectuivEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetQueryObjectuivEXT{}
	ϟa.Value = ϟa.Value
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetQueryObjectuivEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetQueryObjecti64vEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetQueryObjecti64vEXT{}
	ϟa.Value = ϟa.Value
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetQueryObjecti64vEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
func (ϟa *GlGetQueryObjectui64vEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetQueryObjectui64vEXT{}
	ϟa.Value = ϟa.Value
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa, ϟo) {
		log.Printf("Applying glGetQueryObjectui64vEXT expected %v got %v", ϟa, ϟo)
	}
	return nil
}
