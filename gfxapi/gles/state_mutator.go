////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

package gles

import (
	"log"
	"reflect"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/gfxapi/state"
	"android.googlesource.com/platform/tools/gpu/memory"
)

func getState(a atom.Atom, s *state.State) *State {
	id := a.ContextID()
	return s.Contexts[id].(*State)
}
func (ϟa *Init) Mutate(ϟs *state.State) error {
	ϟc := &State{}
	ϟc.Init()
	ϟs.Contexts[ϟa.ContextID()] = ϟc
	ϟo := Init_Out{}
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
	backbufferColorId := RenderbufferId(4294967295)   // RenderbufferId
	backbufferDepthId := RenderbufferId(4294967294)   // RenderbufferId
	backbufferStencilId := RenderbufferId(4294967293) // RenderbufferId
	backbufferColor := func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}() // RenderbufferRef
	backbufferColor.Width = ϟa.In.Width
	backbufferColor.Height = ϟa.In.Height
	backbufferColor.Format = ϟa.In.ColorFmt
	ϟc.Instances.Renderbuffers[backbufferColorId] = backbufferColor
	backbufferDepth := func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}() // RenderbufferRef
	backbufferDepth.Width = ϟa.In.Width
	backbufferDepth.Height = ϟa.In.Height
	backbufferDepth.Format = ϟa.In.DepthFmt
	ϟc.Instances.Renderbuffers[backbufferDepthId] = backbufferDepth
	backbufferStencil := func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}() // RenderbufferRef
	backbufferStencil.Width = ϟa.In.Width
	backbufferStencil.Height = ϟa.In.Height
	backbufferStencil.Format = ϟa.In.StencilFmt
	ϟc.Instances.Renderbuffers[backbufferStencilId] = backbufferStencil
	backbuffer := func() *Framebuffer {
		s := &Framebuffer{}
		s.Init()
		return s
	}() // FramebufferRef
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
	ϟc.Rasterizing.Scissor.Width = ϟa.In.Width
	ϟc.Rasterizing.Scissor.Height = ϟa.In.Height
	ϟc.Rasterizing.StencilMask[FaceMode_GL_FRONT] = 4294967295
	ϟc.Rasterizing.StencilMask[FaceMode_GL_BACK] = 4294967295
	ϟc.Rasterizing.Viewport.Width = ϟa.In.Width
	ϟc.Rasterizing.Viewport.Height = ϟa.In.Height
	ϟc.PixelStorage[PixelStoreParameter_GL_PACK_ALIGNMENT] = 4
	ϟc.PixelStorage[PixelStoreParameter_GL_UNPACK_ALIGNMENT] = 4
	for i := int32(0); i < 64; i++ {
		ϟc.VertexAttributeArrays[AttributeLocation(i)] = func() *VertexAttributeArray {
			s := &VertexAttributeArray{}
			s.Init()
			return s
		}()
	}
	_, _, _, _, _, _, _ = backbufferColorId, backbufferDepthId, backbufferStencilId, backbufferColor, backbufferDepth, backbufferStencil, backbuffer
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying init expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *StartTimer) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := StartTimer_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying startTimer expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *StopTimer) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := StopTimer_Out{}
	ϟo.Result = ϟa.Out.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying stopTimer expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *FlushPostBuffer) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := FlushPostBuffer_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying flushPostBuffer expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *EglCreateContext) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := EglCreateContext_Out{}
	ϟo.Version = ϟa.Out.Version
	ϟo.Context = ϟa.Out.Context
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying eglCreateContext expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *EglMakeCurrent) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := EglMakeCurrent_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying eglMakeCurrent expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *EglSwapBuffers) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := EglSwapBuffers_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying eglSwapBuffers expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlEnableClientState) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlEnableClientState_Out{}
	ϟc.Capabilities[Capability(ϟa.In.Type)] = true
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glEnableClientState expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlDisableClientState) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDisableClientState_Out{}
	ϟc.Capabilities[Capability(ϟa.In.Type)] = false
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glDisableClientState expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGetProgramBinaryOES) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetProgramBinaryOES_Out{}
	ϟo.BytesWritten = ϟa.Out.BytesWritten
	ϟo.BinaryFormat = ϟa.Out.BinaryFormat
	ϟo.Binary = ϟa.Out.Binary
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGetProgramBinaryOES expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlProgramBinaryOES) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlProgramBinaryOES_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glProgramBinaryOES expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlStartTilingQCOM) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlStartTilingQCOM_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glStartTilingQCOM expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlEndTilingQCOM) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlEndTilingQCOM_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glEndTilingQCOM expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlDiscardFramebufferEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDiscardFramebufferEXT_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glDiscardFramebufferEXT expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlInsertEventMarkerEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlInsertEventMarkerEXT_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glInsertEventMarkerEXT expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlPushGroupMarkerEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlPushGroupMarkerEXT_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glPushGroupMarkerEXT expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlPopGroupMarkerEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlPopGroupMarkerEXT_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glPopGroupMarkerEXT expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlTexStorage1DEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlTexStorage1DEXT_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glTexStorage1DEXT expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlTexStorage2DEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlTexStorage2DEXT_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glTexStorage2DEXT expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlTexStorage3DEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlTexStorage3DEXT_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glTexStorage3DEXT expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlTextureStorage1DEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlTextureStorage1DEXT_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glTextureStorage1DEXT expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlTextureStorage2DEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlTextureStorage2DEXT_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glTextureStorage2DEXT expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlTextureStorage3DEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlTextureStorage3DEXT_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glTextureStorage3DEXT expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGenVertexArraysOES) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGenVertexArraysOES_Out{}
	ϟo.Arrays = make(VertexArrayIdArray, ϟa.In.Count)
	for i := int32(0); i < ϟa.In.Count; i++ {
		id := VertexArrayId(ϟa.Out.Arrays[i]) // VertexArrayId
		ϟc.Instances.VertexArrays[id] = func() *VertexArray {
			s := &VertexArray{}
			s.Init()
			return s
		}()
		ϟo.Arrays[i] = id
		_ = id
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGenVertexArraysOES expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlBindVertexArrayOES) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlBindVertexArrayOES_Out{}
	if !(ϟc.Instances.VertexArrays.Contains(ϟa.In.Array)) {
		ϟc.Instances.VertexArrays[ϟa.In.Array] = func() *VertexArray {
			s := &VertexArray{}
			s.Init()
			return s
		}()
	}
	ϟc.BoundVertexArray = ϟa.In.Array
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glBindVertexArrayOES expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlDeleteVertexArraysOES) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDeleteVertexArraysOES_Out{}
	for i := int32(0); i < ϟa.In.Count; i++ {
		ϟc.Instances.VertexArrays.Delete(ϟa.In.Arrays[i])
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glDeleteVertexArraysOES expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlIsVertexArrayOES) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlIsVertexArrayOES_Out{}
	ϟo.Result = ϟc.Instances.VertexArrays.Contains(ϟa.In.Array)
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glIsVertexArrayOES expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlEGLImageTargetTexture2DOES) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlEGLImageTargetTexture2DOES_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glEGLImageTargetTexture2DOES expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlEGLImageTargetRenderbufferStorageOES) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlEGLImageTargetRenderbufferStorageOES_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glEGLImageTargetRenderbufferStorageOES expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGetGraphicsResetStatusEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetGraphicsResetStatusEXT_Out{}
	ϟo.Result = ϟa.Out.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGetGraphicsResetStatusEXT expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlBindAttribLocation) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlBindAttribLocation_Out{}
	p := ϟc.Instances.Programs.Get(ϟa.In.Program) // ProgramRef
	p.AttributeBindings[ϟa.In.Name] = ϟa.In.Location
	_ = p
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glBindAttribLocation expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlBlendFunc) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlBlendFunc_Out{}
	ϟc.Blending.SrcRgbBlendFactor = ϟa.In.SrcFactor
	ϟc.Blending.SrcAlphaBlendFactor = ϟa.In.SrcFactor
	ϟc.Blending.DstRgbBlendFactor = ϟa.In.DstFactor
	ϟc.Blending.DstAlphaBlendFactor = ϟa.In.DstFactor
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glBlendFunc expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlBlendFuncSeparate) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlBlendFuncSeparate_Out{}
	ϟc.Blending.SrcRgbBlendFactor = ϟa.In.SrcFactorRgb
	ϟc.Blending.DstRgbBlendFactor = ϟa.In.DstFactorRgb
	ϟc.Blending.SrcAlphaBlendFactor = ϟa.In.SrcFactorAlpha
	ϟc.Blending.DstAlphaBlendFactor = ϟa.In.DstFactorAlpha
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glBlendFuncSeparate expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlBlendEquation) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlBlendEquation_Out{}
	ϟc.Blending.BlendEquationRgb = ϟa.In.Equation
	ϟc.Blending.BlendEquationAlpha = ϟa.In.Equation
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glBlendEquation expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlBlendEquationSeparate) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlBlendEquationSeparate_Out{}
	ϟc.Blending.BlendEquationRgb = ϟa.In.Rgb
	ϟc.Blending.BlendEquationAlpha = ϟa.In.Alpha
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glBlendEquationSeparate expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlBlendColor) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlBlendColor_Out{}
	ϟc.Blending.BlendColor = func() Color {
		s := Color{}
		s.Init()
		s.Red = ϟa.In.Red
		s.Green = ϟa.In.Green
		s.Blue = ϟa.In.Blue
		s.Alpha = ϟa.In.Alpha
		return s
	}()
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glBlendColor expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlEnableVertexAttribArray) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlEnableVertexAttribArray_Out{}
	a := ϟc.VertexAttributeArrays.Get(ϟa.In.Location) // VertexAttributeArrayRef
	a.Enabled = true
	_ = a
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glEnableVertexAttribArray expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlDisableVertexAttribArray) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDisableVertexAttribArray_Out{}
	a := ϟc.VertexAttributeArrays.Get(ϟa.In.Location) // VertexAttributeArrayRef
	a.Enabled = false
	_ = a
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glDisableVertexAttribArray expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlVertexAttribPointer) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlVertexAttribPointer_Out{}
	a := ϟc.VertexAttributeArrays.Get(ϟa.In.Location) // VertexAttributeArrayRef
	a.Size = ϟa.In.Size
	a.Type = ϟa.In.Type
	a.Normalized = ϟa.In.Normalized
	a.Stride = ϟa.In.Stride
	a.Data = memory.Pointer(ϟa.In.Data)
	_ = a
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glVertexAttribPointer expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGetActiveAttrib) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetActiveAttrib_Out{}
	ϟo.BufferBytesWritten = ϟa.Out.BufferBytesWritten
	ϟo.VectorCount = ϟa.Out.VectorCount
	ϟo.Type = ϟa.Out.Type
	ϟo.Name = ϟa.Out.Name
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGetActiveAttrib expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGetActiveUniform) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetActiveUniform_Out{}
	ϟo.BufferBytesWritten = ϟa.Out.BufferBytesWritten
	ϟo.Size = ϟa.Out.Size
	ϟo.Type = ϟa.Out.Type
	ϟo.Name = ϟa.Out.Name
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGetActiveUniform expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGetError) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetError_Out{}
	ϟo.Result = ϟa.Out.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGetError expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGetProgramiv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetProgramiv_Out{}
	ϟo.Value = make(S32Array, 1)
	ϟo.Value = ϟa.Out.Value
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGetProgramiv expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGetShaderiv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetShaderiv_Out{}
	ϟo.Value = make(S32Array, 1)
	s := ϟc.Instances.Shaders.Get(ϟa.In.Shader) // ShaderRef
	ϟo.Value[0] = func() (result int32) {
		switch ϟa.In.Parameter {
		case ShaderParameter_GL_SHADER_TYPE:
			return int32(s.Type)
		case ShaderParameter_GL_DELETE_STATUS:
			return func() (result int32) {
				switch s.Deletable {
				case true:
					return 1
				case false:
					return 0
				default:
					// TODO: better unmatched handling
					log.Panicf("Unmatched switch in capture")
					return result
				}
			}()
		case ShaderParameter_GL_COMPILE_STATUS:
			return func() (result int32) {
				switch s.Compiled {
				case true:
					return 1
				case false:
					return 0
				default:
					// TODO: better unmatched handling
					log.Panicf("Unmatched switch in capture")
					return result
				}
			}()
		case ShaderParameter_GL_INFO_LOG_LENGTH:
			return strlen(s.InfoLog)
		case ShaderParameter_GL_SHADER_SOURCE_LENGTH:
			return strlen(s.Source[0])
		default:
			// TODO: better unmatched handling
			log.Panicf("Unmatched switch in capture")
			return result
		}
	}()
	_ = s
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGetShaderiv expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGetUniformLocation) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetUniformLocation_Out{}
	ϟo.Result = ϟa.Out.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGetUniformLocation expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGetAttribLocation) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetAttribLocation_Out{}
	ϟo.Result = ϟa.Out.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGetAttribLocation expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlPixelStorei) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlPixelStorei_Out{}
	ϟc.PixelStorage[ϟa.In.Parameter] = ϟa.In.Value
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glPixelStorei expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlTexParameteri) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlTexParameteri_Out{}
	id := ϟc.TextureUnits.Get(ϟc.ActiveTextureUnit).Get(ϟa.In.Target) // TextureId
	t := ϟc.Instances.Textures.Get(id)                                // TextureRef
	switch ϟa.In.Parameter {
	case TextureParameter_GL_TEXTURE_MAG_FILTER:
		t.MagFilter = TextureFilterMode(ϟa.In.Value)
	case TextureParameter_GL_TEXTURE_MIN_FILTER:
		t.MinFilter = TextureFilterMode(ϟa.In.Value)
	case TextureParameter_GL_TEXTURE_WRAP_S:
		t.WrapS = TextureWrapMode(ϟa.In.Value)
	case TextureParameter_GL_TEXTURE_WRAP_T:
		t.WrapT = TextureWrapMode(ϟa.In.Value)
	case TextureParameter_GL_TEXTURE_MAX_ANISOTROPY_EXT:
		t.MaxAnisotropy = float32(ϟa.In.Value)
	case TextureParameter_GL_TEXTURE_SWIZZLE_R:
		t.SwizzleR = TexelComponent(ϟa.In.Value)
	case TextureParameter_GL_TEXTURE_SWIZZLE_G:
		t.SwizzleG = TexelComponent(ϟa.In.Value)
	case TextureParameter_GL_TEXTURE_SWIZZLE_B:
		t.SwizzleB = TexelComponent(ϟa.In.Value)
	case TextureParameter_GL_TEXTURE_SWIZZLE_A:
		t.SwizzleA = TexelComponent(ϟa.In.Value)
	default:
		// TODO: better unmatched handling
		v := ϟa.In.Parameter
		log.Printf("Error: Missing switch case handler for value %T %v", v, v)
	}
	_, _ = id, t
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glTexParameteri expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlTexParameterf) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlTexParameterf_Out{}
	id := ϟc.TextureUnits.Get(ϟc.ActiveTextureUnit).Get(ϟa.In.Target) // TextureId
	t := ϟc.Instances.Textures.Get(id)                                // TextureRef
	switch ϟa.In.Parameter {
	case TextureParameter_GL_TEXTURE_MAG_FILTER:
		t.MagFilter = TextureFilterMode(ϟa.In.Value)
	case TextureParameter_GL_TEXTURE_MIN_FILTER:
		t.MinFilter = TextureFilterMode(ϟa.In.Value)
	case TextureParameter_GL_TEXTURE_WRAP_S:
		t.WrapS = TextureWrapMode(ϟa.In.Value)
	case TextureParameter_GL_TEXTURE_WRAP_T:
		t.WrapT = TextureWrapMode(ϟa.In.Value)
	case TextureParameter_GL_TEXTURE_MAX_ANISOTROPY_EXT:
		t.MaxAnisotropy = ϟa.In.Value
	case TextureParameter_GL_TEXTURE_SWIZZLE_R:
		t.SwizzleR = TexelComponent(ϟa.In.Value)
	case TextureParameter_GL_TEXTURE_SWIZZLE_G:
		t.SwizzleG = TexelComponent(ϟa.In.Value)
	case TextureParameter_GL_TEXTURE_SWIZZLE_B:
		t.SwizzleB = TexelComponent(ϟa.In.Value)
	case TextureParameter_GL_TEXTURE_SWIZZLE_A:
		t.SwizzleA = TexelComponent(ϟa.In.Value)
	default:
		// TODO: better unmatched handling
		v := ϟa.In.Parameter
		log.Printf("Error: Missing switch case handler for value %T %v", v, v)
	}
	_, _ = id, t
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glTexParameterf expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGetTexParameteriv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetTexParameteriv_Out{}
	ϟo.Values = make(S32Array, 1)
	id := ϟc.TextureUnits.Get(ϟc.ActiveTextureUnit).Get(ϟa.In.Target) // TextureId
	t := ϟc.Instances.Textures.Get(id)                                // TextureRef
	ϟo.Values[0] = func() (result int32) {
		switch ϟa.In.Parameter {
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
			log.Panicf("Unmatched switch in capture")
			return result
		}
	}()
	_, _ = id, t
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGetTexParameteriv expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGetTexParameterfv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetTexParameterfv_Out{}
	ϟo.Values = make(F32Array, 1)
	id := ϟc.TextureUnits.Get(ϟc.ActiveTextureUnit).Get(ϟa.In.Target) // TextureId
	t := ϟc.Instances.Textures.Get(id)                                // TextureRef
	ϟo.Values[0] = func() (result float32) {
		switch ϟa.In.Parameter {
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
			log.Panicf("Unmatched switch in capture")
			return result
		}
	}()
	_, _ = id, t
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGetTexParameterfv expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlUniform1i) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniform1i_Out{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramRef
	uniform := program.Uniforms.Get(ϟa.In.Location)       // Uniform
	uniform.Type = ShaderUniformType_GL_INT
	uniform.Value.S32 = ϟa.In.Value
	program.Uniforms[ϟa.In.Location] = uniform
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glUniform1i expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlUniform2i) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniform2i_Out{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramRef
	uniform := program.Uniforms.Get(ϟa.In.Location)       // Uniform
	uniform.Type = ShaderUniformType_GL_INT_VEC2
	uniform.Value.Vec2i = func() Vec2i {
		s := Vec2i{}
		s.Init()
		s.X = ϟa.In.Value0
		s.Y = ϟa.In.Value1
		return s
	}()
	program.Uniforms[ϟa.In.Location] = uniform
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glUniform2i expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlUniform3i) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniform3i_Out{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramRef
	uniform := program.Uniforms.Get(ϟa.In.Location)       // Uniform
	uniform.Type = ShaderUniformType_GL_INT_VEC3
	uniform.Value.Vec3i = func() Vec3i {
		s := Vec3i{}
		s.Init()
		s.X = ϟa.In.Value0
		s.Y = ϟa.In.Value1
		s.Z = ϟa.In.Value2
		return s
	}()
	program.Uniforms[ϟa.In.Location] = uniform
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glUniform3i expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlUniform4i) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniform4i_Out{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramRef
	uniform := program.Uniforms.Get(ϟa.In.Location)       // Uniform
	uniform.Type = ShaderUniformType_GL_INT_VEC4
	uniform.Value.Vec4i = func() Vec4i {
		s := Vec4i{}
		s.Init()
		s.X = ϟa.In.Value0
		s.Y = ϟa.In.Value1
		s.Z = ϟa.In.Value2
		s.W = ϟa.In.Value3
		return s
	}()
	program.Uniforms[ϟa.In.Location] = uniform
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glUniform4i expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlUniform1iv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniform1iv_Out{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramRef
	uniform := program.Uniforms.Get(ϟa.In.Location)       // Uniform
	uniform.Type = ShaderUniformType_GL_INT
	uniform.Value.S32 = ϟa.In.Value[0]
	program.Uniforms[ϟa.In.Location] = uniform
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glUniform1iv expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlUniform2iv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniform2iv_Out{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramRef
	uniform := program.Uniforms.Get(ϟa.In.Location)       // Uniform
	uniform.Type = ShaderUniformType_GL_INT_VEC2
	uniform.Value.Vec2i = func() Vec2i {
		s := Vec2i{}
		s.Init()
		s.X = ϟa.In.Value[0]
		s.Y = ϟa.In.Value[1]
		return s
	}()
	program.Uniforms[ϟa.In.Location] = uniform
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glUniform2iv expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlUniform3iv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniform3iv_Out{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramRef
	uniform := program.Uniforms.Get(ϟa.In.Location)       // Uniform
	uniform.Type = ShaderUniformType_GL_INT_VEC3
	uniform.Value.Vec3i = func() Vec3i {
		s := Vec3i{}
		s.Init()
		s.X = ϟa.In.Value[0]
		s.Y = ϟa.In.Value[1]
		s.Z = ϟa.In.Value[2]
		return s
	}()
	program.Uniforms[ϟa.In.Location] = uniform
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glUniform3iv expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlUniform4iv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniform4iv_Out{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramRef
	uniform := program.Uniforms.Get(ϟa.In.Location)       // Uniform
	uniform.Type = ShaderUniformType_GL_INT_VEC4
	uniform.Value.Vec4i = func() Vec4i {
		s := Vec4i{}
		s.Init()
		s.X = ϟa.In.Value[0]
		s.Y = ϟa.In.Value[1]
		s.Z = ϟa.In.Value[2]
		s.W = ϟa.In.Value[3]
		return s
	}()
	program.Uniforms[ϟa.In.Location] = uniform
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glUniform4iv expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlUniform1f) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniform1f_Out{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramRef
	uniform := program.Uniforms.Get(ϟa.In.Location)       // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT
	uniform.Value.F32 = ϟa.In.Value
	program.Uniforms[ϟa.In.Location] = uniform
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glUniform1f expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlUniform2f) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniform2f_Out{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramRef
	uniform := program.Uniforms.Get(ϟa.In.Location)       // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT_VEC2
	uniform.Value.Vec2f = func() Vec2f {
		s := Vec2f{}
		s.Init()
		s.X = ϟa.In.Value0
		s.Y = ϟa.In.Value1
		return s
	}()
	program.Uniforms[ϟa.In.Location] = uniform
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glUniform2f expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlUniform3f) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniform3f_Out{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramRef
	uniform := program.Uniforms.Get(ϟa.In.Location)       // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT_VEC3
	uniform.Value.Vec3f = func() Vec3f {
		s := Vec3f{}
		s.Init()
		s.X = ϟa.In.Value0
		s.Y = ϟa.In.Value1
		s.Z = ϟa.In.Value2
		return s
	}()
	program.Uniforms[ϟa.In.Location] = uniform
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glUniform3f expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlUniform4f) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniform4f_Out{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramRef
	uniform := program.Uniforms.Get(ϟa.In.Location)       // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT_VEC4
	uniform.Value.Vec4f = func() Vec4f {
		s := Vec4f{}
		s.Init()
		s.X = ϟa.In.Value0
		s.Y = ϟa.In.Value1
		s.Z = ϟa.In.Value2
		s.W = ϟa.In.Value3
		return s
	}()
	program.Uniforms[ϟa.In.Location] = uniform
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glUniform4f expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlUniform1fv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniform1fv_Out{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramRef
	uniform := program.Uniforms.Get(ϟa.In.Location)       // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT
	uniform.Value.F32 = ϟa.In.Value[0]
	program.Uniforms[ϟa.In.Location] = uniform
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glUniform1fv expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlUniform2fv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniform2fv_Out{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramRef
	uniform := program.Uniforms.Get(ϟa.In.Location)       // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT_VEC2
	uniform.Value.Vec2f = func() Vec2f {
		s := Vec2f{}
		s.Init()
		s.X = ϟa.In.Value[0]
		s.Y = ϟa.In.Value[1]
		return s
	}()
	program.Uniforms[ϟa.In.Location] = uniform
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glUniform2fv expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlUniform3fv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniform3fv_Out{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramRef
	uniform := program.Uniforms.Get(ϟa.In.Location)       // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT_VEC3
	uniform.Value.Vec3f = func() Vec3f {
		s := Vec3f{}
		s.Init()
		s.X = ϟa.In.Value[0]
		s.Y = ϟa.In.Value[1]
		s.Z = ϟa.In.Value[2]
		return s
	}()
	program.Uniforms[ϟa.In.Location] = uniform
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glUniform3fv expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlUniform4fv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniform4fv_Out{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramRef
	uniform := program.Uniforms.Get(ϟa.In.Location)       // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT_VEC4
	uniform.Value.Vec4f = func() Vec4f {
		s := Vec4f{}
		s.Init()
		s.X = ϟa.In.Value[0]
		s.Y = ϟa.In.Value[1]
		s.Z = ϟa.In.Value[2]
		s.W = ϟa.In.Value[3]
		return s
	}()
	program.Uniforms[ϟa.In.Location] = uniform
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glUniform4fv expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlUniformMatrix2fv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniformMatrix2fv_Out{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramRef
	uniform := program.Uniforms.Get(ϟa.In.Location)       // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT_MAT2
	uniform.Value.Mat2f = func() Mat2f {
		s := Mat2f{}
		s.Init()
		s.Col0 = func() Vec2f {
			s := Vec2f{}
			s.Init()
			s.X = ϟa.In.Values[0]
			s.Y = ϟa.In.Values[1]
			return s
		}()
		s.Col1 = func() Vec2f {
			s := Vec2f{}
			s.Init()
			s.X = ϟa.In.Values[0]
			s.Y = ϟa.In.Values[1]
			return s
		}()
		return s
	}()
	program.Uniforms[ϟa.In.Location] = uniform
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glUniformMatrix2fv expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlUniformMatrix3fv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniformMatrix3fv_Out{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramRef
	uniform := program.Uniforms.Get(ϟa.In.Location)       // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT_MAT3
	uniform.Value.Mat3f = func() Mat3f {
		s := Mat3f{}
		s.Init()
		s.Col0 = func() Vec3f {
			s := Vec3f{}
			s.Init()
			s.X = ϟa.In.Values[0]
			s.Y = ϟa.In.Values[1]
			s.Z = ϟa.In.Values[2]
			return s
		}()
		s.Col1 = func() Vec3f {
			s := Vec3f{}
			s.Init()
			s.X = ϟa.In.Values[3]
			s.Y = ϟa.In.Values[4]
			s.Z = ϟa.In.Values[5]
			return s
		}()
		s.Col2 = func() Vec3f {
			s := Vec3f{}
			s.Init()
			s.X = ϟa.In.Values[6]
			s.Y = ϟa.In.Values[7]
			s.Z = ϟa.In.Values[8]
			return s
		}()
		return s
	}()
	program.Uniforms[ϟa.In.Location] = uniform
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glUniformMatrix3fv expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlUniformMatrix4fv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUniformMatrix4fv_Out{}
	program := ϟc.Instances.Programs.Get(ϟc.BoundProgram) // ProgramRef
	uniform := program.Uniforms.Get(ϟa.In.Location)       // Uniform
	uniform.Value.Mat4f = func() Mat4f {
		s := Mat4f{}
		s.Init()
		s.Col0 = func() Vec4f {
			s := Vec4f{}
			s.Init()
			s.X = ϟa.In.Values[0]
			s.Y = ϟa.In.Values[1]
			s.Z = ϟa.In.Values[2]
			s.W = ϟa.In.Values[3]
			return s
		}()
		s.Col1 = func() Vec4f {
			s := Vec4f{}
			s.Init()
			s.X = ϟa.In.Values[4]
			s.Y = ϟa.In.Values[5]
			s.Z = ϟa.In.Values[6]
			s.W = ϟa.In.Values[7]
			return s
		}()
		s.Col2 = func() Vec4f {
			s := Vec4f{}
			s.Init()
			s.X = ϟa.In.Values[8]
			s.Y = ϟa.In.Values[9]
			s.Z = ϟa.In.Values[10]
			s.W = ϟa.In.Values[11]
			return s
		}()
		s.Col3 = func() Vec4f {
			s := Vec4f{}
			s.Init()
			s.X = ϟa.In.Values[12]
			s.Y = ϟa.In.Values[13]
			s.Z = ϟa.In.Values[14]
			s.W = ϟa.In.Values[15]
			return s
		}()
		return s
	}()
	program.Uniforms[ϟa.In.Location] = uniform
	_, _ = program, uniform
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glUniformMatrix4fv expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGetUniformfv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetUniformfv_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGetUniformfv expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGetUniformiv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetUniformiv_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGetUniformiv expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlVertexAttrib1f) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlVertexAttrib1f_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glVertexAttrib1f expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlVertexAttrib2f) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlVertexAttrib2f_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glVertexAttrib2f expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlVertexAttrib3f) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlVertexAttrib3f_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glVertexAttrib3f expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlVertexAttrib4f) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlVertexAttrib4f_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glVertexAttrib4f expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlVertexAttrib1fv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlVertexAttrib1fv_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glVertexAttrib1fv expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlVertexAttrib2fv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlVertexAttrib2fv_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glVertexAttrib2fv expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlVertexAttrib3fv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlVertexAttrib3fv_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glVertexAttrib3fv expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlVertexAttrib4fv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlVertexAttrib4fv_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glVertexAttrib4fv expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGetShaderPrecisionFormat) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetShaderPrecisionFormat_Out{}
	ϟo.Range = make(S32Array, 2)
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGetShaderPrecisionFormat expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlDepthMask) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDepthMask_Out{}
	ϟc.Rasterizing.DepthMask = ϟa.In.Enabled
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glDepthMask expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlDepthFunc) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDepthFunc_Out{}
	ϟc.Rasterizing.DepthTestFunction = ϟa.In.Function
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glDepthFunc expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlDepthRangef) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDepthRangef_Out{}
	ϟc.Rasterizing.DepthNear = ϟa.In.Near
	ϟc.Rasterizing.DepthFar = ϟa.In.Far
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glDepthRangef expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlColorMask) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlColorMask_Out{}
	ϟc.Rasterizing.ColorMaskRed = ϟa.In.Red
	ϟc.Rasterizing.ColorMaskGreen = ϟa.In.Green
	ϟc.Rasterizing.ColorMaskBlue = ϟa.In.Blue
	ϟc.Rasterizing.ColorMaskAlpha = ϟa.In.Alpha
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glColorMask expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlStencilMask) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlStencilMask_Out{}
	ϟc.Rasterizing.StencilMask[FaceMode_GL_FRONT] = ϟa.In.Mask
	ϟc.Rasterizing.StencilMask[FaceMode_GL_BACK] = ϟa.In.Mask
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glStencilMask expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlStencilMaskSeparate) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlStencilMaskSeparate_Out{}
	switch ϟa.In.Face {
	case FaceMode_GL_FRONT:
		ϟc.Rasterizing.StencilMask[FaceMode_GL_FRONT] = ϟa.In.Mask
	case FaceMode_GL_BACK:
		ϟc.Rasterizing.StencilMask[FaceMode_GL_BACK] = ϟa.In.Mask
	case FaceMode_GL_FRONT_AND_BACK:
		ϟc.Rasterizing.StencilMask[FaceMode_GL_FRONT] = ϟa.In.Mask
		ϟc.Rasterizing.StencilMask[FaceMode_GL_BACK] = ϟa.In.Mask
	default:
		// TODO: better unmatched handling
		v := ϟa.In.Face
		log.Printf("Error: Missing switch case handler for value %T %v", v, v)
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glStencilMaskSeparate expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlStencilFuncSeparate) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlStencilFuncSeparate_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glStencilFuncSeparate expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlStencilOpSeparate) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlStencilOpSeparate_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glStencilOpSeparate expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlFrontFace) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlFrontFace_Out{}
	ϟc.Rasterizing.FrontFace = ϟa.In.Orientation
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glFrontFace expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlViewport) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlViewport_Out{}
	ϟc.Rasterizing.Viewport = func() Rect {
		s := Rect{}
		s.Init()
		s.X = ϟa.In.X
		s.Y = ϟa.In.Y
		s.Width = ϟa.In.Width
		s.Height = ϟa.In.Height
		return s
	}()
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glViewport expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlScissor) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlScissor_Out{}
	ϟc.Rasterizing.Scissor = func() Rect {
		s := Rect{}
		s.Init()
		s.X = ϟa.In.X
		s.Y = ϟa.In.Y
		s.Width = ϟa.In.Width
		s.Height = ϟa.In.Height
		return s
	}()
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glScissor expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlActiveTexture) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlActiveTexture_Out{}
	ϟc.ActiveTextureUnit = ϟa.In.Unit
	if !(ϟc.TextureUnits.Contains(ϟa.In.Unit)) {
		ϟc.TextureUnits[ϟa.In.Unit] = ϟc.TextureUnits.Get(ϟa.In.Unit)
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glActiveTexture expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGenTextures) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGenTextures_Out{}
	ϟo.Textures = make(TextureIdArray, ϟa.In.Count)
	for i := int32(0); i < ϟa.In.Count; i++ {
		id := TextureId(ϟa.Out.Textures[i]) // TextureId
		ϟc.Instances.Textures[id] = func() *Texture {
			s := &Texture{}
			s.Init()
			return s
		}()
		ϟo.Textures[i] = id
		_ = id
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGenTextures expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlDeleteTextures) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDeleteTextures_Out{}
	for i := int32(0); i < ϟa.In.Count; i++ {
		ϟc.Instances.Textures.Delete(ϟa.In.Textures[i])
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glDeleteTextures expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlIsTexture) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlIsTexture_Out{}
	ϟo.Result = ϟc.Instances.Textures.Contains(ϟa.In.Texture)
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glIsTexture expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlBindTexture) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlBindTexture_Out{}
	if !(ϟc.Instances.Textures.Contains(ϟa.In.Texture)) {
		ϟc.Instances.Textures[ϟa.In.Texture] = func() *Texture {
			s := &Texture{}
			s.Init()
			return s
		}()
	}
	ϟc.TextureUnits.Get(ϟc.ActiveTextureUnit)[ϟa.In.Target] = ϟa.In.Texture
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glBindTexture expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlTexImage2D) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlTexImage2D_Out{}
	switch ϟa.In.Target {
	case TextureImageTarget_GL_TEXTURE_2D:
		id := ϟc.TextureUnits.Get(ϟc.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_2D) // TextureId
		t := ϟc.Instances.Textures.Get(id)                                               // TextureRef
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.In.Width
			s.Height = ϟa.In.Height
			s.Size = imageSize(ϟa.In.Width, ϟa.In.Height, ϟa.In.Format, ϟa.In.Type)
			s.Format = ImageTexelFormat(ϟa.In.Format)
			return s
		}() // Image
		l.Data.Write(ϟs.Memory.Slice(memory.Range{
			Base: memory.Pointer(ϟa.In.Data),
			Size: uint64(l.Size),
		}))
		t.Texture2D[ϟa.In.Level] = l
		t.Kind = TextureKind_TEXTURE2D
		t.Format = ImageTexelFormat(ϟa.In.Format)
		_, _, _ = id, t, l
	case TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X, TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Z, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z:
		id := ϟc.TextureUnits.Get(ϟc.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_CUBE_MAP) // TextureId
		t := ϟc.Instances.Textures.Get(id)                                                     // TextureRef
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.In.Width
			s.Height = ϟa.In.Height
			s.Size = imageSize(ϟa.In.Width, ϟa.In.Height, ϟa.In.Format, ϟa.In.Type)
			s.Format = ImageTexelFormat(ϟa.In.Format)
			return s
		}() // Image
		cube := t.Cubemap.Get(ϟa.In.Level) // CubemapLevel
		cube.Faces[CubeMapImageTarget(ϟa.In.Target)] = l
		t.Cubemap[ϟa.In.Level] = cube
		t.Kind = TextureKind_CUBEMAP
		t.Format = ImageTexelFormat(ϟa.In.Format)
		_, _, _, _ = id, t, l, cube
	default:
		// TODO: better unmatched handling
		v := ϟa.In.Target
		log.Printf("Error: Missing switch case handler for value %T %v", v, v)
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glTexImage2D expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlTexSubImage2D) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlTexSubImage2D_Out{}
	switch ϟa.In.Target {
	case TextureImageTarget_GL_TEXTURE_2D:
		id := ϟc.TextureUnits.Get(ϟc.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_2D) // TextureId
		t := ϟc.Instances.Textures.Get(id)                                               // TextureRef
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.In.Width
			s.Height = ϟa.In.Height
			s.Size = imageSize(ϟa.In.Width, ϟa.In.Height, ϟa.In.Format, ϟa.In.Type)
			s.Format = ImageTexelFormat(ϟa.In.Format)
			return s
		}() // Image
		l.Data.Write(ϟs.Memory.Slice(memory.Range{
			Base: memory.Pointer(ϟa.In.Data),
			Size: uint64(l.Size),
		}))
		t.Texture2D[ϟa.In.Level] = l
		t.Kind = TextureKind_TEXTURE2D
		t.Format = ImageTexelFormat(ϟa.In.Format)
		_, _, _ = id, t, l
	case TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X, TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Z, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z:
		id := ϟc.TextureUnits.Get(ϟc.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_CUBE_MAP) // TextureId
		t := ϟc.Instances.Textures.Get(id)                                                     // TextureRef
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.In.Width
			s.Height = ϟa.In.Height
			s.Size = imageSize(ϟa.In.Width, ϟa.In.Height, ϟa.In.Format, ϟa.In.Type)
			s.Format = ImageTexelFormat(ϟa.In.Format)
			return s
		}() // Image
		l.Data.Write(ϟs.Memory.Slice(memory.Range{
			Base: memory.Pointer(ϟa.In.Data),
			Size: uint64(l.Size),
		}))
		cube := t.Cubemap.Get(ϟa.In.Level) // CubemapLevel
		cube.Faces[CubeMapImageTarget(ϟa.In.Target)] = l
		t.Cubemap[ϟa.In.Level] = cube
		t.Kind = TextureKind_CUBEMAP
		t.Format = ImageTexelFormat(ϟa.In.Format)
		_, _, _, _ = id, t, l, cube
	default:
		// TODO: better unmatched handling
		v := ϟa.In.Target
		log.Printf("Error: Missing switch case handler for value %T %v", v, v)
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glTexSubImage2D expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlCopyTexImage2D) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlCopyTexImage2D_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glCopyTexImage2D expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlCopyTexSubImage2D) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlCopyTexSubImage2D_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glCopyTexSubImage2D expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlCompressedTexImage2D) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlCompressedTexImage2D_Out{}
	switch ϟa.In.Target {
	case TextureImageTarget_GL_TEXTURE_2D:
		id := ϟc.TextureUnits.Get(ϟc.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_2D) // TextureId
		t := ϟc.Instances.Textures.Get(id)                                               // TextureRef
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.In.Width
			s.Height = ϟa.In.Height
			s.Size = ϟa.In.ImageSize
			s.Format = ImageTexelFormat(ϟa.In.Format)
			return s
		}() // Image
		l.Data.Write(ϟs.Memory.Slice(memory.Range{
			Base: memory.Pointer(ϟa.In.Data),
			Size: uint64(l.Size),
		}))
		t.Texture2D[ϟa.In.Level] = l
		t.Kind = TextureKind_TEXTURE2D
		t.Format = ImageTexelFormat(ϟa.In.Format)
		_, _, _ = id, t, l
	case TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X, TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Z, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z:
		id := ϟc.TextureUnits.Get(ϟc.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_CUBE_MAP) // TextureId
		t := ϟc.Instances.Textures.Get(id)                                                     // TextureRef
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.In.Width
			s.Height = ϟa.In.Height
			s.Size = ϟa.In.ImageSize
			s.Format = ImageTexelFormat(ϟa.In.Format)
			return s
		}() // Image
		l.Data.Write(ϟs.Memory.Slice(memory.Range{
			Base: memory.Pointer(ϟa.In.Data),
			Size: uint64(l.Size),
		}))
		cube := t.Cubemap.Get(ϟa.In.Level) // CubemapLevel
		cube.Faces[CubeMapImageTarget(ϟa.In.Target)] = l
		t.Cubemap[ϟa.In.Level] = cube
		t.Kind = TextureKind_CUBEMAP
		t.Format = ImageTexelFormat(ϟa.In.Format)
		_, _, _, _ = id, t, l, cube
	default:
		// TODO: better unmatched handling
		v := ϟa.In.Target
		log.Printf("Error: Missing switch case handler for value %T %v", v, v)
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glCompressedTexImage2D expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlCompressedTexSubImage2D) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlCompressedTexSubImage2D_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glCompressedTexSubImage2D expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGenerateMipmap) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGenerateMipmap_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGenerateMipmap expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlReadPixels) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlReadPixels_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glReadPixels expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGenFramebuffers) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGenFramebuffers_Out{}
	ϟo.Framebuffers = make(FramebufferIdArray, ϟa.In.Count)
	for i := int32(0); i < ϟa.In.Count; i++ {
		id := FramebufferId(ϟa.Out.Framebuffers[i]) // FramebufferId
		ϟc.Instances.Framebuffers[id] = func() *Framebuffer {
			s := &Framebuffer{}
			s.Init()
			return s
		}()
		ϟo.Framebuffers[i] = id
		_ = id
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGenFramebuffers expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlBindFramebuffer) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlBindFramebuffer_Out{}
	if !(ϟc.Instances.Framebuffers.Contains(ϟa.In.Framebuffer)) {
		ϟc.Instances.Framebuffers[ϟa.In.Framebuffer] = func() *Framebuffer {
			s := &Framebuffer{}
			s.Init()
			return s
		}()
	}
	if (ϟa.In.Target) == (FramebufferTarget_GL_FRAMEBUFFER) {
		ϟc.BoundFramebuffers[FramebufferTarget_GL_READ_FRAMEBUFFER] = ϟa.In.Framebuffer
		ϟc.BoundFramebuffers[FramebufferTarget_GL_DRAW_FRAMEBUFFER] = ϟa.In.Framebuffer
	} else {
		ϟc.BoundFramebuffers[ϟa.In.Target] = ϟa.In.Framebuffer
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glBindFramebuffer expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlCheckFramebufferStatus) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlCheckFramebufferStatus_Out{}
	ϟo.Result = ϟa.Out.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glCheckFramebufferStatus expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlDeleteFramebuffers) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDeleteFramebuffers_Out{}
	for i := int32(0); i < ϟa.In.Count; i++ {
		ϟc.Instances.Framebuffers.Delete(ϟa.In.Framebuffers[i])
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glDeleteFramebuffers expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlIsFramebuffer) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlIsFramebuffer_Out{}
	ϟo.Result = ϟc.Instances.Framebuffers.Contains(ϟa.In.Framebuffer)
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glIsFramebuffer expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGenRenderbuffers) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGenRenderbuffers_Out{}
	ϟo.Renderbuffers = make(RenderbufferIdArray, ϟa.In.Count)
	for i := int32(0); i < ϟa.In.Count; i++ {
		id := RenderbufferId(ϟa.Out.Renderbuffers[i]) // RenderbufferId
		ϟc.Instances.Renderbuffers[id] = func() *Renderbuffer {
			s := &Renderbuffer{}
			s.Init()
			return s
		}()
		ϟo.Renderbuffers[i] = id
		_ = id
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGenRenderbuffers expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlBindRenderbuffer) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlBindRenderbuffer_Out{}
	if !(ϟc.Instances.Renderbuffers.Contains(ϟa.In.Renderbuffer)) {
		ϟc.Instances.Renderbuffers[ϟa.In.Renderbuffer] = func() *Renderbuffer {
			s := &Renderbuffer{}
			s.Init()
			return s
		}()
	}
	ϟc.BoundRenderbuffers[ϟa.In.Target] = ϟa.In.Renderbuffer
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glBindRenderbuffer expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlRenderbufferStorage) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlRenderbufferStorage_Out{}
	id := ϟc.BoundRenderbuffers.Get(ϟa.In.Target) // RenderbufferId
	rb := ϟc.Instances.Renderbuffers.Get(id)      // RenderbufferRef
	rb.Format = ϟa.In.Format
	rb.Width = ϟa.In.Width
	rb.Height = ϟa.In.Height
	_, _ = id, rb
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glRenderbufferStorage expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlDeleteRenderbuffers) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDeleteRenderbuffers_Out{}
	for i := int32(0); i < ϟa.In.Count; i++ {
		ϟc.Instances.Renderbuffers.Delete(ϟa.In.Renderbuffers[i])
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glDeleteRenderbuffers expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlIsRenderbuffer) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlIsRenderbuffer_Out{}
	ϟo.Result = ϟc.Instances.Renderbuffers.Contains(ϟa.In.Renderbuffer)
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glIsRenderbuffer expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGetRenderbufferParameteriv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetRenderbufferParameteriv_Out{}
	ϟo.Values = make(S32Array, 1)
	id := ϟc.BoundRenderbuffers.Get(ϟa.In.Target) // RenderbufferId
	rb := ϟc.Instances.Renderbuffers.Get(id)      // RenderbufferRef
	ϟo.Values[0] = func() (result int32) {
		switch ϟa.In.Parameter {
		case RenderbufferParameter_GL_RENDERBUFFER_WIDTH:
			return rb.Width
		case RenderbufferParameter_GL_RENDERBUFFER_HEIGHT:
			return rb.Height
		case RenderbufferParameter_GL_RENDERBUFFER_INTERNAL_FORMAT:
			return int32(rb.Format)
		default:
			// TODO: better unmatched handling
			log.Panicf("Unmatched switch in capture")
			return result
		}
	}()
	_, _ = id, rb
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGetRenderbufferParameteriv expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGenBuffers) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGenBuffers_Out{}
	ϟo.Buffers = make(BufferIdArray, ϟa.In.Count)
	for i := int32(0); i < ϟa.In.Count; i++ {
		id := BufferId(ϟa.Out.Buffers[i]) // BufferId
		ϟc.Instances.Buffers[id] = func() *Buffer {
			s := &Buffer{}
			s.Init()
			return s
		}()
		ϟo.Buffers[i] = id
		_ = id
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGenBuffers expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlBindBuffer) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlBindBuffer_Out{}
	if !(ϟc.Instances.Buffers.Contains(ϟa.In.Buffer)) {
		ϟc.Instances.Buffers[ϟa.In.Buffer] = func() *Buffer {
			s := &Buffer{}
			s.Init()
			return s
		}()
	}
	ϟc.BoundBuffers[ϟa.In.Target] = ϟa.In.Buffer
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glBindBuffer expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlBufferData) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlBufferData_Out{}
	id := ϟc.BoundBuffers.Get(ϟa.In.Target) // BufferId
	b := ϟc.Instances.Buffers.Get(id)       // BufferRef
	b.Data.Write(ϟs.Memory.Slice(memory.Range{
		Base: memory.Pointer(ϟa.In.Data),
		Size: uint64(ϟa.In.Size),
	}))
	b.Size = ϟa.In.Size
	b.Usage = ϟa.In.Usage
	_, _ = id, b
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glBufferData expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlBufferSubData) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlBufferSubData_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glBufferSubData expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlDeleteBuffers) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDeleteBuffers_Out{}
	for i := int32(0); i < ϟa.In.Count; i++ {
		ϟc.Instances.Buffers.Delete(ϟa.In.Buffers[i])
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glDeleteBuffers expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlIsBuffer) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlIsBuffer_Out{}
	ϟo.Result = ϟc.Instances.Buffers.Contains(ϟa.In.Buffer)
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glIsBuffer expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGetBufferParameteriv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetBufferParameteriv_Out{}
	id := ϟc.BoundBuffers.Get(ϟa.In.Target) // BufferId
	b := ϟc.Instances.Buffers.Get(id)       // BufferRef
	ϟo.Value = func() (result int32) {
		switch ϟa.In.Parameter {
		case BufferParameter_GL_BUFFER_SIZE:
			return b.Size
		case BufferParameter_GL_BUFFER_USAGE:
			return int32(b.Usage)
		default:
			// TODO: better unmatched handling
			log.Panicf("Unmatched switch in capture")
			return result
		}
	}()
	_, _ = id, b
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGetBufferParameteriv expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlCreateShader) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlCreateShader_Out{}
	id := ShaderId(ϟa.Out.Result) // ShaderId
	ϟc.Instances.Shaders[id] = func() *Shader {
		s := &Shader{}
		s.Init()
		return s
	}()
	s := ϟc.Instances.Shaders.Get(id) // ShaderRef
	s.Type = ϟa.In.Type
	ϟo.Result = id
	_, _ = id, s
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glCreateShader expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlDeleteShader) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDeleteShader_Out{}
	s := ϟc.Instances.Shaders.Get(ϟa.In.Shader) // ShaderRef
	s.Deletable = true
	ϟc.Instances.Shaders.Delete(ϟa.In.Shader)
	_ = s
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glDeleteShader expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlShaderSource) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlShaderSource_Out{}
	s := ϟc.Instances.Shaders.Get(ϟa.In.Shader) // ShaderRef
	s.Source = ϟa.In.Source
	_ = s
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glShaderSource expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlShaderBinary) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlShaderBinary_Out{}
	for i := int32(0); i < ϟa.In.Count; i++ {
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glShaderBinary expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGetShaderInfoLog) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetShaderInfoLog_Out{}
	s := ϟc.Instances.Shaders.Get(ϟa.In.Shader) // ShaderRef
	min_0_a := ϟa.In.BufferLength               // s32
	min_0_b := strlen(s.InfoLog)                // s32
	min_0_result := func() (result int32) {
		switch (min_0_a) < (min_0_b) {
		case true:
			return min_0_a
		case false:
			return min_0_b
		default:
			// TODO: better unmatched handling
			log.Panicf("Unmatched switch in capture")
			return result
		}
	}() // s32
	ϟo.StringLengthWritten = min_0_result
	ϟo.Info = substr(s.InfoLog, 0, ϟo.StringLengthWritten)
	_, _, _, _ = s, min_0_a, min_0_b, min_0_result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGetShaderInfoLog expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGetShaderSource) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetShaderSource_Out{}
	s := ϟc.Instances.Shaders.Get(ϟa.In.Shader) // ShaderRef
	min_1_a := ϟa.In.BufferLength               // s32
	min_1_b := strlen(s.Source[0])              // s32
	min_1_result := func() (result int32) {
		switch (min_1_a) < (min_1_b) {
		case true:
			return min_1_a
		case false:
			return min_1_b
		default:
			// TODO: better unmatched handling
			log.Panicf("Unmatched switch in capture")
			return result
		}
	}() // s32
	ϟo.StringLengthWritten = min_1_result
	ϟo.Source = substr(s.Source[0], 0, ϟo.StringLengthWritten)
	_, _, _, _ = s, min_1_a, min_1_b, min_1_result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGetShaderSource expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlReleaseShaderCompiler) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlReleaseShaderCompiler_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glReleaseShaderCompiler expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlCompileShader) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlCompileShader_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glCompileShader expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlIsShader) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlIsShader_Out{}
	ϟo.Result = ϟc.Instances.Shaders.Contains(ϟa.In.Shader)
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glIsShader expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlCreateProgram) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlCreateProgram_Out{}
	id := ProgramId(ϟa.Out.Result) // ProgramId
	ϟc.Instances.Programs[id] = func() *Program {
		s := &Program{}
		s.Init()
		return s
	}()
	ϟo.Result = id
	_ = id
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glCreateProgram expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlDeleteProgram) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDeleteProgram_Out{}
	ϟc.Instances.Programs.Delete(ϟa.In.Program)
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glDeleteProgram expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlAttachShader) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlAttachShader_Out{}
	p := ϟc.Instances.Programs.Get(ϟa.In.Program) // ProgramRef
	s := ϟc.Instances.Shaders.Get(ϟa.In.Shader)   // ShaderRef
	p.Shaders[s.Type] = ϟa.In.Shader
	_, _ = p, s
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glAttachShader expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlDetachShader) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDetachShader_Out{}
	p := ϟc.Instances.Programs.Get(ϟa.In.Program) // ProgramRef
	s := ϟc.Instances.Shaders.Get(ϟa.In.Shader)   // ShaderRef
	p.Shaders.Delete(s.Type)
	_, _ = p, s
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glDetachShader expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGetAttachedShaders) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetAttachedShaders_Out{}
	ϟo.Shaders = make(ShaderIdArray, ϟa.In.BufferLength)
	p := ϟc.Instances.Programs.Get(ϟa.In.Program) // ProgramRef
	min_2_a := ϟa.In.BufferLength                 // s32
	min_2_b := int32(len(p.Shaders))              // s32
	min_2_result := func() (result int32) {
		switch (min_2_a) < (min_2_b) {
		case true:
			return min_2_a
		case false:
			return min_2_b
		default:
			// TODO: better unmatched handling
			log.Panicf("Unmatched switch in capture")
			return result
		}
	}() // s32
	ϟo.ShadersLengthWritten = min_2_result
	ϟo.Shaders = p.Shaders.Range()
	_, _, _, _ = p, min_2_a, min_2_b, min_2_result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGetAttachedShaders expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlLinkProgram) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlLinkProgram_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glLinkProgram expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGetProgramInfoLog) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetProgramInfoLog_Out{}
	p := ϟc.Instances.Programs.Get(ϟa.In.Program) // ProgramRef
	min_3_a := ϟa.In.BufferLength                 // s32
	min_3_b := strlen(p.InfoLog)                  // s32
	min_3_result := func() (result int32) {
		switch (min_3_a) < (min_3_b) {
		case true:
			return min_3_a
		case false:
			return min_3_b
		default:
			// TODO: better unmatched handling
			log.Panicf("Unmatched switch in capture")
			return result
		}
	}() // s32
	ϟo.StringLengthWritten = min_3_result
	ϟo.Info = substr(p.InfoLog, 0, ϟo.StringLengthWritten)
	_, _, _, _ = p, min_3_a, min_3_b, min_3_result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGetProgramInfoLog expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlUseProgram) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUseProgram_Out{}
	ϟc.BoundProgram = ϟa.In.Program
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glUseProgram expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlIsProgram) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlIsProgram_Out{}
	ϟo.Result = ϟc.Instances.Programs.Contains(ϟa.In.Program)
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glIsProgram expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlValidateProgram) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlValidateProgram_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glValidateProgram expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlClearColor) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlClearColor_Out{}
	ϟc.Clearing.ClearColor = func() Color {
		s := Color{}
		s.Init()
		s.Red = ϟa.In.R
		s.Green = ϟa.In.G
		s.Blue = ϟa.In.B
		s.Alpha = ϟa.In.A
		return s
	}()
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glClearColor expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlClearDepthf) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlClearDepthf_Out{}
	ϟc.Clearing.ClearDepth = ϟa.In.Depth
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glClearDepthf expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlClearStencil) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlClearStencil_Out{}
	ϟc.Clearing.ClearStencil = ϟa.In.Stencil
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glClearStencil expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlClear) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlClear_Out{}
	if (ClearMask_GL_COLOR_BUFFER_BIT)&(ϟa.In.Mask) != 0 {
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glClear expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlCullFace) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlCullFace_Out{}
	ϟc.Rasterizing.CullFace = ϟa.In.Mode
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glCullFace expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlPolygonOffset) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlPolygonOffset_Out{}
	ϟc.Rasterizing.PolygonOffsetUnits = ϟa.In.Units
	ϟc.Rasterizing.PolygonOffsetFactor = ϟa.In.ScaleFactor
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glPolygonOffset expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlLineWidth) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlLineWidth_Out{}
	ϟc.Rasterizing.LineWidth = ϟa.In.Width
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glLineWidth expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlSampleCoverage) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlSampleCoverage_Out{}
	ϟc.Rasterizing.SampleCoverageValue = ϟa.In.Value
	ϟc.Rasterizing.SampleCoverageInvert = ϟa.In.Invert
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glSampleCoverage expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlHint) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlHint_Out{}
	ϟc.GenerateMipmapHint = ϟa.In.Mode
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glHint expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlFramebufferRenderbuffer) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlFramebufferRenderbuffer_Out{}
	target := func() (result FramebufferTarget) {
		switch ϟa.In.FramebufferTarget {
		case FramebufferTarget_GL_FRAMEBUFFER:
			return FramebufferTarget_GL_DRAW_FRAMEBUFFER
		case FramebufferTarget_GL_DRAW_FRAMEBUFFER:
			return FramebufferTarget_GL_DRAW_FRAMEBUFFER
		case FramebufferTarget_GL_READ_FRAMEBUFFER:
			return FramebufferTarget_GL_READ_FRAMEBUFFER
		default:
			// TODO: better unmatched handling
			log.Panicf("Unmatched switch in capture")
			return result
		}
	}() // FramebufferTarget
	framebufferId := ϟc.BoundFramebuffers.Get(target)                      // FramebufferId
	framebuffer := ϟc.Instances.Framebuffers.Get(framebufferId)            // FramebufferRef
	attachment := framebuffer.Attachments.Get(ϟa.In.FramebufferAttachment) // FramebufferAttachmentInfo
	if (ϟa.In.Renderbuffer) == (ϟc.Internals.NilRenderbuffer) {
		attachment.Type = FramebufferAttachmentType_GL_NONE
	} else {
		attachment.Type = FramebufferAttachmentType_GL_RENDERBUFFER
	}
	attachment.Object = uint32(ϟa.In.Renderbuffer)
	attachment.TextureLevel = 0
	attachment.CubeMapFace = CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X
	framebuffer.Attachments[ϟa.In.FramebufferAttachment] = attachment
	_, _, _, _ = target, framebufferId, framebuffer, attachment
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glFramebufferRenderbuffer expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlFramebufferTexture2D) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlFramebufferTexture2D_Out{}
	target := func() (result FramebufferTarget) {
		switch ϟa.In.FramebufferTarget {
		case FramebufferTarget_GL_FRAMEBUFFER:
			return FramebufferTarget_GL_DRAW_FRAMEBUFFER
		case FramebufferTarget_GL_DRAW_FRAMEBUFFER:
			return FramebufferTarget_GL_DRAW_FRAMEBUFFER
		case FramebufferTarget_GL_READ_FRAMEBUFFER:
			return FramebufferTarget_GL_READ_FRAMEBUFFER
		default:
			// TODO: better unmatched handling
			log.Panicf("Unmatched switch in capture")
			return result
		}
	}() // FramebufferTarget
	framebufferId := ϟc.BoundFramebuffers.Get(target)                      // FramebufferId
	framebuffer := ϟc.Instances.Framebuffers.Get(framebufferId)            // FramebufferRef
	attachment := framebuffer.Attachments.Get(ϟa.In.FramebufferAttachment) // FramebufferAttachmentInfo
	if (ϟa.In.Texture) == (ϟc.Internals.NilTexture) {
		attachment.Type = FramebufferAttachmentType_GL_NONE
		attachment.Object = uint32(ϟc.Internals.NilTexture)
		attachment.TextureLevel = 0
		attachment.CubeMapFace = CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X
	} else {
		attachment.Type = FramebufferAttachmentType_GL_TEXTURE
		attachment.Object = uint32(ϟa.In.Texture)
		attachment.TextureLevel = ϟa.In.Level
		attachment.CubeMapFace = func() (result CubeMapImageTarget) {
			switch ϟa.In.TextureTarget {
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
				log.Panicf("Unmatched switch in capture")
				return result
			}
		}()
	}
	framebuffer.Attachments[ϟa.In.FramebufferAttachment] = attachment
	_, _, _, _ = target, framebufferId, framebuffer, attachment
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glFramebufferTexture2D expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGetFramebufferAttachmentParameteriv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetFramebufferAttachmentParameteriv_Out{}
	ϟo.Value = make(S32Array, 1)
	target := func() (result FramebufferTarget) {
		switch ϟa.In.FramebufferTarget {
		case FramebufferTarget_GL_FRAMEBUFFER:
			return FramebufferTarget_GL_DRAW_FRAMEBUFFER
		case FramebufferTarget_GL_DRAW_FRAMEBUFFER:
			return FramebufferTarget_GL_DRAW_FRAMEBUFFER
		case FramebufferTarget_GL_READ_FRAMEBUFFER:
			return FramebufferTarget_GL_READ_FRAMEBUFFER
		default:
			// TODO: better unmatched handling
			log.Panicf("Unmatched switch in capture")
			return result
		}
	}() // FramebufferTarget
	framebufferId := ϟc.BoundFramebuffers.Get(target)           // FramebufferId
	framebuffer := ϟc.Instances.Framebuffers.Get(framebufferId) // FramebufferRef
	a := framebuffer.Attachments.Get(ϟa.In.Attachment)          // FramebufferAttachmentInfo
	ϟo.Value[0] = func() (result int32) {
		switch ϟa.In.Parameter {
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
			log.Panicf("Unmatched switch in capture")
			return result
		}
	}()
	_, _, _, _ = target, framebufferId, framebuffer, a
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGetFramebufferAttachmentParameteriv expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlDrawElements) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDrawElements_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glDrawElements expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlDrawArrays) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDrawArrays_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glDrawArrays expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlFlush) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlFlush_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glFlush expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlFinish) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlFinish_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glFinish expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGetBooleanv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetBooleanv_Out{}
	ϟo.Values = make(BoolArray, stateVariableSize(ϟa.In.Param))
	switch ϟa.In.Param {
	case StateVariable_GL_BLEND:
		ϟo.Values[0] = ϟc.Capabilities.Get(Capability_GL_BLEND)
	case StateVariable_GL_CULL_FACE:
		ϟo.Values[0] = ϟc.Capabilities.Get(Capability_GL_CULL_FACE)
	case StateVariable_GL_DEPTH_TEST:
		ϟo.Values[0] = ϟc.Capabilities.Get(Capability_GL_DEPTH_TEST)
	case StateVariable_GL_DITHER:
		ϟo.Values[0] = ϟc.Capabilities.Get(Capability_GL_DITHER)
	case StateVariable_GL_POLYGON_OFFSET_FILL:
		ϟo.Values[0] = ϟc.Capabilities.Get(Capability_GL_POLYGON_OFFSET_FILL)
	case StateVariable_GL_SAMPLE_ALPHA_TO_COVERAGE:
		ϟo.Values[0] = ϟc.Capabilities.Get(Capability_GL_SAMPLE_ALPHA_TO_COVERAGE)
	case StateVariable_GL_SAMPLE_COVERAGE:
		ϟo.Values[0] = ϟc.Capabilities.Get(Capability_GL_SAMPLE_COVERAGE)
	case StateVariable_GL_SCISSOR_TEST:
		ϟo.Values[0] = ϟc.Capabilities.Get(Capability_GL_SCISSOR_TEST)
	case StateVariable_GL_STENCIL_TEST:
		ϟo.Values[0] = ϟc.Capabilities.Get(Capability_GL_STENCIL_TEST)
	case StateVariable_GL_DEPTH_WRITEMASK:
		ϟo.Values[0] = ϟc.Rasterizing.DepthMask
	case StateVariable_GL_COLOR_WRITEMASK:
		ϟo.Values[0] = ϟc.Rasterizing.ColorMaskRed
		ϟo.Values[1] = ϟc.Rasterizing.ColorMaskGreen
		ϟo.Values[2] = ϟc.Rasterizing.ColorMaskBlue
		ϟo.Values[3] = ϟc.Rasterizing.ColorMaskAlpha
	case StateVariable_GL_SAMPLE_COVERAGE_INVERT:
		ϟo.Values[0] = ϟc.Rasterizing.SampleCoverageInvert
	case StateVariable_GL_SHADER_COMPILER:
		ϟo.Values[0] = ϟa.Out.Values[0]
	case StateVariable_GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT:
		ϟo.Values[0] = ϟa.Out.Values[0]
	default:
		// TODO: better unmatched handling
		v := ϟa.In.Param
		log.Printf("Error: Missing switch case handler for value %T %v", v, v)
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGetBooleanv expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGetFloatv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetFloatv_Out{}
	ϟo.Values = make(F32Array, stateVariableSize(ϟa.In.Param))
	switch ϟa.In.Param {
	case StateVariable_GL_DEPTH_RANGE:
		ϟo.Values[0] = ϟc.Rasterizing.DepthNear
		ϟo.Values[1] = ϟc.Rasterizing.DepthFar
	case StateVariable_GL_LINE_WIDTH:
		ϟo.Values[0] = ϟc.Rasterizing.LineWidth
	case StateVariable_GL_POLYGON_OFFSET_FACTOR:
		ϟo.Values[0] = ϟc.Rasterizing.PolygonOffsetFactor
	case StateVariable_GL_POLYGON_OFFSET_UNITS:
		ϟo.Values[0] = ϟc.Rasterizing.PolygonOffsetUnits
	case StateVariable_GL_SAMPLE_COVERAGE_VALUE:
		ϟo.Values[0] = ϟc.Rasterizing.SampleCoverageValue
	case StateVariable_GL_COLOR_CLEAR_VALUE:
		ϟo.Values[0] = ϟc.Clearing.ClearColor.Red
		ϟo.Values[1] = ϟc.Clearing.ClearColor.Green
		ϟo.Values[2] = ϟc.Clearing.ClearColor.Blue
		ϟo.Values[3] = ϟc.Clearing.ClearColor.Alpha
	case StateVariable_GL_DEPTH_CLEAR_VALUE:
		ϟo.Values[0] = ϟc.Clearing.ClearDepth
	case StateVariable_GL_ALIASED_LINE_WIDTH_RANGE:
		ϟo.Values[0] = ϟa.Out.Values[0]
		ϟo.Values[1] = ϟa.Out.Values[1]
	case StateVariable_GL_ALIASED_POINT_SIZE_RANGE:
		ϟo.Values[0] = ϟa.Out.Values[0]
		ϟo.Values[1] = ϟa.Out.Values[1]
	case StateVariable_GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT:
		ϟo.Values[0] = ϟa.Out.Values[0]
	default:
		// TODO: better unmatched handling
		v := ϟa.In.Param
		log.Printf("Error: Missing switch case handler for value %T %v", v, v)
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGetFloatv expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGetIntegerv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetIntegerv_Out{}
	ϟo.Values = make(S32Array, stateVariableSize(ϟa.In.Param))
	switch ϟa.In.Param {
	case StateVariable_GL_ACTIVE_TEXTURE:
		ϟo.Values[0] = int32(ϟc.ActiveTextureUnit)
	case StateVariable_GL_ARRAY_BUFFER_BINDING:
		ϟo.Values[0] = int32(ϟc.BoundBuffers.Get(BufferTarget_GL_ARRAY_BUFFER))
	case StateVariable_GL_ELEMENT_ARRAY_BUFFER_BINDING:
		ϟo.Values[0] = int32(ϟc.BoundBuffers.Get(BufferTarget_GL_ELEMENT_ARRAY_BUFFER))
	case StateVariable_GL_BLEND_SRC_ALPHA:
		ϟo.Values[0] = int32(ϟc.Blending.SrcAlphaBlendFactor)
	case StateVariable_GL_BLEND_SRC_RGB:
		ϟo.Values[0] = int32(ϟc.Blending.SrcRgbBlendFactor)
	case StateVariable_GL_BLEND_DST_ALPHA:
		ϟo.Values[0] = int32(ϟc.Blending.DstAlphaBlendFactor)
	case StateVariable_GL_BLEND_DST_RGB:
		ϟo.Values[0] = int32(ϟc.Blending.DstRgbBlendFactor)
	case StateVariable_GL_BLEND_EQUATION_RGB:
		ϟo.Values[0] = int32(ϟc.Blending.BlendEquationRgb)
	case StateVariable_GL_BLEND_EQUATION_ALPHA:
		ϟo.Values[0] = int32(ϟc.Blending.BlendEquationAlpha)
	case StateVariable_GL_BLEND_COLOR:
		ϟo.Values[0] = int32(ϟc.Blending.BlendColor.Red)
		ϟo.Values[1] = int32(ϟc.Blending.BlendColor.Green)
		ϟo.Values[2] = int32(ϟc.Blending.BlendColor.Blue)
		ϟo.Values[3] = int32(ϟc.Blending.BlendColor.Alpha)
	case StateVariable_GL_DEPTH_FUNC:
		ϟo.Values[0] = int32(ϟc.Rasterizing.DepthTestFunction)
	case StateVariable_GL_DEPTH_CLEAR_VALUE:
		ϟo.Values[0] = int32(ϟc.Clearing.ClearDepth)
	case StateVariable_GL_STENCIL_WRITEMASK:
		ϟo.Values[0] = int32(ϟc.Rasterizing.StencilMask.Get(FaceMode_GL_FRONT))
	case StateVariable_GL_STENCIL_BACK_WRITEMASK:
		ϟo.Values[0] = int32(ϟc.Rasterizing.StencilMask.Get(FaceMode_GL_BACK))
	case StateVariable_GL_VIEWPORT:
		ϟo.Values[0] = ϟc.Rasterizing.Viewport.X
		ϟo.Values[1] = ϟc.Rasterizing.Viewport.Y
		ϟo.Values[2] = ϟc.Rasterizing.Viewport.Width
		ϟo.Values[3] = ϟc.Rasterizing.Viewport.Height
	case StateVariable_GL_SCISSOR_BOX:
		ϟo.Values[0] = ϟc.Rasterizing.Scissor.X
		ϟo.Values[1] = ϟc.Rasterizing.Scissor.Y
		ϟo.Values[2] = ϟc.Rasterizing.Scissor.Width
		ϟo.Values[3] = ϟc.Rasterizing.Scissor.Height
	case StateVariable_GL_FRONT_FACE:
		ϟo.Values[0] = int32(ϟc.Rasterizing.FrontFace)
	case StateVariable_GL_CULL_FACE_MODE:
		ϟo.Values[0] = int32(ϟc.Rasterizing.CullFace)
	case StateVariable_GL_STENCIL_CLEAR_VALUE:
		ϟo.Values[0] = ϟc.Clearing.ClearStencil
	case StateVariable_GL_FRAMEBUFFER_BINDING:
		ϟo.Values[0] = int32(ϟc.BoundFramebuffers.Get(FramebufferTarget_GL_FRAMEBUFFER))
	case StateVariable_GL_READ_FRAMEBUFFER_BINDING:
		ϟo.Values[0] = int32(ϟc.BoundFramebuffers.Get(FramebufferTarget_GL_READ_FRAMEBUFFER))
	case StateVariable_GL_RENDERBUFFER_BINDING:
		ϟo.Values[0] = int32(ϟc.BoundRenderbuffers.Get(RenderbufferTarget_GL_RENDERBUFFER))
	case StateVariable_GL_CURRENT_PROGRAM:
		ϟo.Values[0] = int32(ϟc.BoundProgram)
	case StateVariable_GL_TEXTURE_BINDING_2D:
		ϟo.Values[0] = int32(ϟc.TextureUnits.Get(ϟc.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_2D))
	case StateVariable_GL_TEXTURE_BINDING_CUBE_MAP:
		ϟo.Values[0] = int32(ϟc.TextureUnits.Get(ϟc.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_CUBE_MAP))
	case StateVariable_GL_GENERATE_MIPMAP_HINT:
		ϟo.Values[0] = int32(ϟc.GenerateMipmapHint)
	case StateVariable_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS:
		ϟo.Values[0] = ϟa.Out.Values[0]
	case StateVariable_GL_MAX_CUBE_MAP_TEXTURE_SIZE:
		ϟo.Values[0] = ϟa.Out.Values[0]
	case StateVariable_GL_MAX_FRAGMENT_UNIFORM_VECTORS:
		ϟo.Values[0] = ϟa.Out.Values[0]
	case StateVariable_GL_MAX_RENDERBUFFER_SIZE:
		ϟo.Values[0] = ϟa.Out.Values[0]
	case StateVariable_GL_MAX_TEXTURE_IMAGE_UNITS:
		ϟo.Values[0] = ϟa.Out.Values[0]
	case StateVariable_GL_MAX_TEXTURE_SIZE:
		ϟo.Values[0] = ϟa.Out.Values[0]
	case StateVariable_GL_MAX_VARYING_VECTORS:
		ϟo.Values[0] = ϟa.Out.Values[0]
	case StateVariable_GL_MAX_VERTEX_ATTRIBS:
		ϟo.Values[0] = ϟa.Out.Values[0]
	case StateVariable_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS:
		ϟo.Values[0] = ϟa.Out.Values[0]
	case StateVariable_GL_MAX_VERTEX_UNIFORM_VECTORS:
		ϟo.Values[0] = ϟa.Out.Values[0]
	case StateVariable_GL_MAX_VIEWPORT_DIMS:
		ϟo.Values[0] = ϟa.Out.Values[0]
		ϟo.Values[1] = ϟa.Out.Values[1]
	case StateVariable_GL_NUM_COMPRESSED_TEXTURE_FORMATS:
		ϟo.Values[0] = ϟa.Out.Values[0]
	case StateVariable_GL_NUM_SHADER_BINARY_FORMATS:
		ϟo.Values[0] = ϟa.Out.Values[0]
	case StateVariable_GL_PACK_ALIGNMENT:
		ϟo.Values[0] = ϟc.PixelStorage.Get(PixelStoreParameter_GL_PACK_ALIGNMENT)
	case StateVariable_GL_UNPACK_ALIGNMENT:
		ϟo.Values[0] = ϟc.PixelStorage.Get(PixelStoreParameter_GL_UNPACK_ALIGNMENT)
	case StateVariable_GL_ALPHA_BITS:
		ϟo.Values[0] = ϟa.Out.Values[0]
	case StateVariable_GL_BLUE_BITS:
		ϟo.Values[0] = ϟa.Out.Values[0]
	case StateVariable_GL_GREEN_BITS:
		ϟo.Values[0] = ϟa.Out.Values[0]
	case StateVariable_GL_RED_BITS:
		ϟo.Values[0] = ϟa.Out.Values[0]
	case StateVariable_GL_DEPTH_BITS:
		ϟo.Values[0] = ϟa.Out.Values[0]
	case StateVariable_GL_SAMPLE_BUFFERS:
		ϟo.Values[0] = ϟa.Out.Values[0]
	case StateVariable_GL_SAMPLES:
		ϟo.Values[0] = ϟa.Out.Values[0]
	case StateVariable_GL_SHADER_BINARY_FORMATS:
		ϟo.Values = ϟa.Out.Values
	case StateVariable_GL_COMPRESSED_TEXTURE_FORMATS:
		ϟo.Values = ϟa.Out.Values
	case StateVariable_GL_STENCIL_BITS:
		ϟo.Values[0] = ϟa.Out.Values[0]
	case StateVariable_GL_SUBPIXEL_BITS:
		ϟo.Values[0] = ϟa.Out.Values[0]
	case StateVariable_GL_IMPLEMENTATION_COLOR_READ_FORMAT:
		ϟo.Values[0] = ϟa.Out.Values[0]
	case StateVariable_GL_IMPLEMENTATION_COLOR_READ_TYPE:
		ϟo.Values[0] = ϟa.Out.Values[0]
	case StateVariable_GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT:
		ϟo.Values[0] = ϟa.Out.Values[0]
	case StateVariable_GL_GPU_DISJOINT_EXT:
		ϟo.Values[0] = ϟa.Out.Values[0]
	default:
		// TODO: better unmatched handling
		v := ϟa.In.Param
		log.Printf("Error: Missing switch case handler for value %T %v", v, v)
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGetIntegerv expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGetString) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetString_Out{}
	ϟo.Result = ϟa.Out.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGetString expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlEnable) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlEnable_Out{}
	ϟc.Capabilities[ϟa.In.Capability] = true
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glEnable expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlDisable) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDisable_Out{}
	ϟc.Capabilities[ϟa.In.Capability] = false
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glDisable expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlIsEnabled) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlIsEnabled_Out{}
	ϟo.Result = ϟc.Capabilities.Get(ϟa.In.Capability)
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glIsEnabled expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlMapBufferRange) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlMapBufferRange_Out{}
	ϟo.Result = ϟa.Out.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glMapBufferRange expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlUnmapBuffer) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlUnmapBuffer_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glUnmapBuffer expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlInvalidateFramebuffer) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlInvalidateFramebuffer_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glInvalidateFramebuffer expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlRenderbufferStorageMultisample) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlRenderbufferStorageMultisample_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glRenderbufferStorageMultisample expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlBlitFramebuffer) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlBlitFramebuffer_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glBlitFramebuffer expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGenQueries) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGenQueries_Out{}
	ϟo.Queries = make(QueryIdArray, ϟa.In.Count)
	for i := int32(0); i < ϟa.In.Count; i++ {
		id := QueryId(ϟa.Out.Queries[i]) // QueryId
		ϟc.Instances.Queries[id] = func() *Query {
			s := &Query{}
			s.Init()
			return s
		}()
		ϟo.Queries[i] = id
		_ = id
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGenQueries expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlBeginQuery) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlBeginQuery_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glBeginQuery expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlEndQuery) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlEndQuery_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glEndQuery expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlDeleteQueries) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDeleteQueries_Out{}
	for i := int32(0); i < ϟa.In.Count; i++ {
		ϟc.Instances.Queries.Delete(ϟa.In.Queries[i])
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glDeleteQueries expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlIsQuery) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlIsQuery_Out{}
	ϟo.Result = ϟc.Instances.Queries.Contains(ϟa.In.Query)
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glIsQuery expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGetQueryiv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetQueryiv_Out{}
	ϟo.Value = ϟa.Out.Value
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGetQueryiv expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGetQueryObjectuiv) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetQueryObjectuiv_Out{}
	ϟo.Value = ϟa.Out.Value
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGetQueryObjectuiv expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGenQueriesEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGenQueriesEXT_Out{}
	ϟo.Queries = make(QueryIdArray, ϟa.In.Count)
	for i := int32(0); i < ϟa.In.Count; i++ {
		id := QueryId(ϟa.Out.Queries[i]) // QueryId
		ϟc.Instances.Queries[id] = func() *Query {
			s := &Query{}
			s.Init()
			return s
		}()
		ϟo.Queries[i] = id
		_ = id
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGenQueriesEXT expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlBeginQueryEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlBeginQueryEXT_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glBeginQueryEXT expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlEndQueryEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlEndQueryEXT_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glEndQueryEXT expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlDeleteQueriesEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlDeleteQueriesEXT_Out{}
	for i := int32(0); i < ϟa.In.Count; i++ {
		ϟc.Instances.Queries.Delete(ϟa.In.Queries[i])
	}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glDeleteQueriesEXT expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlIsQueryEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlIsQueryEXT_Out{}
	ϟo.Result = ϟc.Instances.Queries.Contains(ϟa.In.Query)
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glIsQueryEXT expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlQueryCounterEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlQueryCounterEXT_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glQueryCounterEXT expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGetQueryivEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetQueryivEXT_Out{}
	ϟo.Value = ϟa.Out.Value
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGetQueryivEXT expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGetQueryObjectivEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetQueryObjectivEXT_Out{}
	ϟo.Value = ϟa.Out.Value
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGetQueryObjectivEXT expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGetQueryObjectuivEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetQueryObjectuivEXT_Out{}
	ϟo.Value = ϟa.Out.Value
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGetQueryObjectuivEXT expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGetQueryObjecti64vEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetQueryObjecti64vEXT_Out{}
	ϟo.Value = ϟa.Out.Value
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGetQueryObjecti64vEXT expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *GlGetQueryObjectui64vEXT) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := GlGetQueryObjectui64vEXT_Out{}
	ϟo.Value = ϟa.Out.Value
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying glGetQueryObjectui64vEXT expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
