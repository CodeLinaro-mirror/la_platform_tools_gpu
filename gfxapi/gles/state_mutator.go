////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

package gles

import (
	"log"
	"reflect"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/memory"
)

type StateMutator struct {
	State          *state
	ValidateOutput bool
}

var _ atom.Writer = StateMutator{} // Interface compliance test
func (m StateMutator) Write(ψ atom.ID, Θ atom.Atom) {
	switch ω := Θ.(type) {
	case *Init:
		Σ := Init_Out{}
		m.State.Instances.Buffers[m.State.Internals.NilBuffer] = func() *Buffer {
			s := &Buffer{}
			s.Init()
			return s
		}()
		m.State.Instances.Textures[m.State.Internals.NilTexture] = func() *Texture {
			s := &Texture{}
			s.Init()
			return s
		}()
		m.State.Instances.Renderbuffers[m.State.Internals.NilRenderbuffer] = func() *Renderbuffer {
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
		backbufferColor.Width = ω.In.Width
		backbufferColor.Height = ω.In.Height
		backbufferColor.Format = ω.In.ColorFmt
		m.State.Instances.Renderbuffers[backbufferColorId] = backbufferColor
		backbufferDepth := func() *Renderbuffer {
			s := &Renderbuffer{}
			s.Init()
			return s
		}() // RenderbufferRef
		backbufferDepth.Width = ω.In.Width
		backbufferDepth.Height = ω.In.Height
		backbufferDepth.Format = ω.In.DepthFmt
		m.State.Instances.Renderbuffers[backbufferDepthId] = backbufferDepth
		backbufferStencil := func() *Renderbuffer {
			s := &Renderbuffer{}
			s.Init()
			return s
		}() // RenderbufferRef
		backbufferStencil.Width = ω.In.Width
		backbufferStencil.Height = ω.In.Height
		backbufferStencil.Format = ω.In.StencilFmt
		m.State.Instances.Renderbuffers[backbufferStencilId] = backbufferStencil
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
		m.State.Instances.Framebuffers[m.State.Internals.Backbuffer] = backbuffer
		m.State.BoundFramebuffers[FramebufferTarget_GL_FRAMEBUFFER] = m.State.Internals.Backbuffer
		m.State.Rasterizing.Scissor.Width = ω.In.Width
		m.State.Rasterizing.Scissor.Height = ω.In.Height
		m.State.Rasterizing.StencilMask[FaceMode_GL_FRONT] = 4294967295
		m.State.Rasterizing.StencilMask[FaceMode_GL_BACK] = 4294967295
		m.State.Rasterizing.Viewport.Width = ω.In.Width
		m.State.Rasterizing.Viewport.Height = ω.In.Height
		m.State.PixelStorage[PixelStoreParameter_GL_PACK_ALIGNMENT] = 4
		m.State.PixelStorage[PixelStoreParameter_GL_UNPACK_ALIGNMENT] = 4
		for i := int32(0); i < 64; i++ {
			m.State.VertexAttributeArrays[AttributeLocation(i)] = func() *VertexAttributeArray {
				s := &VertexAttributeArray{}
				s.Init()
				return s
			}()
		}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying init expected %v got %v", ω.Out, Σ)
		}
	case *StartTimer:
		Σ := StartTimer_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying startTimer expected %v got %v", ω.Out, Σ)
		}
	case *StopTimer:
		Σ := StopTimer_Out{}
		Σ.Result = ω.Out.Result
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying stopTimer expected %v got %v", ω.Out, Σ)
		}
	case *FlushPostBuffer:
		Σ := FlushPostBuffer_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying flushPostBuffer expected %v got %v", ω.Out, Σ)
		}
	case *EglCreateContext:
		Σ := EglCreateContext_Out{}
		Σ.Version = ω.Out.Version
		Σ.Context = ω.Out.Context
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying eglCreateContext expected %v got %v", ω.Out, Σ)
		}
	case *EglMakeCurrent:
		Σ := EglMakeCurrent_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying eglMakeCurrent expected %v got %v", ω.Out, Σ)
		}
	case *EglSwapBuffers:
		Σ := EglSwapBuffers_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying eglSwapBuffers expected %v got %v", ω.Out, Σ)
		}
	case *GlEnableClientState:
		Σ := GlEnableClientState_Out{}
		m.State.Capabilities[Capability(ω.In.Type)] = true
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glEnableClientState expected %v got %v", ω.Out, Σ)
		}
	case *GlDisableClientState:
		Σ := GlDisableClientState_Out{}
		m.State.Capabilities[Capability(ω.In.Type)] = false
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glDisableClientState expected %v got %v", ω.Out, Σ)
		}
	case *GlGetProgramBinaryOES:
		Σ := GlGetProgramBinaryOES_Out{}
		Σ.BytesWritten = ω.Out.BytesWritten
		Σ.BinaryFormat = ω.Out.BinaryFormat
		Σ.Binary = ω.Out.Binary
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGetProgramBinaryOES expected %v got %v", ω.Out, Σ)
		}
	case *GlProgramBinaryOES:
		Σ := GlProgramBinaryOES_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glProgramBinaryOES expected %v got %v", ω.Out, Σ)
		}
	case *GlStartTilingQCOM:
		Σ := GlStartTilingQCOM_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glStartTilingQCOM expected %v got %v", ω.Out, Σ)
		}
	case *GlEndTilingQCOM:
		Σ := GlEndTilingQCOM_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glEndTilingQCOM expected %v got %v", ω.Out, Σ)
		}
	case *GlDiscardFramebufferEXT:
		Σ := GlDiscardFramebufferEXT_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glDiscardFramebufferEXT expected %v got %v", ω.Out, Σ)
		}
	case *GlInsertEventMarkerEXT:
		Σ := GlInsertEventMarkerEXT_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glInsertEventMarkerEXT expected %v got %v", ω.Out, Σ)
		}
	case *GlPushGroupMarkerEXT:
		Σ := GlPushGroupMarkerEXT_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glPushGroupMarkerEXT expected %v got %v", ω.Out, Σ)
		}
	case *GlPopGroupMarkerEXT:
		Σ := GlPopGroupMarkerEXT_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glPopGroupMarkerEXT expected %v got %v", ω.Out, Σ)
		}
	case *GlTexStorage1DEXT:
		Σ := GlTexStorage1DEXT_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glTexStorage1DEXT expected %v got %v", ω.Out, Σ)
		}
	case *GlTexStorage2DEXT:
		Σ := GlTexStorage2DEXT_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glTexStorage2DEXT expected %v got %v", ω.Out, Σ)
		}
	case *GlTexStorage3DEXT:
		Σ := GlTexStorage3DEXT_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glTexStorage3DEXT expected %v got %v", ω.Out, Σ)
		}
	case *GlTextureStorage1DEXT:
		Σ := GlTextureStorage1DEXT_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glTextureStorage1DEXT expected %v got %v", ω.Out, Σ)
		}
	case *GlTextureStorage2DEXT:
		Σ := GlTextureStorage2DEXT_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glTextureStorage2DEXT expected %v got %v", ω.Out, Σ)
		}
	case *GlTextureStorage3DEXT:
		Σ := GlTextureStorage3DEXT_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glTextureStorage3DEXT expected %v got %v", ω.Out, Σ)
		}
	case *GlGenVertexArraysOES:
		Σ := GlGenVertexArraysOES_Out{}
		Σ.Arrays = make(VertexArrayIdArray, ω.In.Count)
		for i := int32(0); i < ω.In.Count; i++ {
			id := VertexArrayId(ω.Out.Arrays[i]) // VertexArrayId
			m.State.Instances.VertexArrays[id] = func() *VertexArray {
				s := &VertexArray{}
				s.Init()
				return s
			}()
			Σ.Arrays[i] = id
		}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGenVertexArraysOES expected %v got %v", ω.Out, Σ)
		}
	case *GlBindVertexArrayOES:
		Σ := GlBindVertexArrayOES_Out{}
		if (m.State.Instances.VertexArrays.Contains(ω.In.Array)) == (false) {
			m.State.Instances.VertexArrays[ω.In.Array] = func() *VertexArray {
				s := &VertexArray{}
				s.Init()
				return s
			}()
		}
		m.State.BoundVertexArray = ω.In.Array
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glBindVertexArrayOES expected %v got %v", ω.Out, Σ)
		}
	case *GlDeleteVertexArraysOES:
		Σ := GlDeleteVertexArraysOES_Out{}
		for i := int32(0); i < ω.In.Count; i++ {
			m.State.Instances.VertexArrays.Delete(ω.In.Arrays[i])
		}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glDeleteVertexArraysOES expected %v got %v", ω.Out, Σ)
		}
	case *GlIsVertexArrayOES:
		Σ := GlIsVertexArrayOES_Out{}
		Σ.Result = m.State.Instances.VertexArrays.Contains(ω.In.Array)
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glIsVertexArrayOES expected %v got %v", ω.Out, Σ)
		}
	case *GlEGLImageTargetTexture2DOES:
		Σ := GlEGLImageTargetTexture2DOES_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glEGLImageTargetTexture2DOES expected %v got %v", ω.Out, Σ)
		}
	case *GlEGLImageTargetRenderbufferStorageOES:
		Σ := GlEGLImageTargetRenderbufferStorageOES_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glEGLImageTargetRenderbufferStorageOES expected %v got %v", ω.Out, Σ)
		}
	case *GlGetGraphicsResetStatusEXT:
		Σ := GlGetGraphicsResetStatusEXT_Out{}
		Σ.Result = ω.Out.Result
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGetGraphicsResetStatusEXT expected %v got %v", ω.Out, Σ)
		}
	case *GlBindAttribLocation:
		Σ := GlBindAttribLocation_Out{}
		p := m.State.Instances.Programs.Get(ω.In.Program) // ProgramRef
		p.AttributeBindings[ω.In.Name] = ω.In.Location
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glBindAttribLocation expected %v got %v", ω.Out, Σ)
		}
	case *GlBlendFunc:
		Σ := GlBlendFunc_Out{}
		m.State.Blending.SrcRgbBlendFactor = ω.In.SrcFactor
		m.State.Blending.SrcAlphaBlendFactor = ω.In.SrcFactor
		m.State.Blending.DstRgbBlendFactor = ω.In.DstFactor
		m.State.Blending.DstAlphaBlendFactor = ω.In.DstFactor
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glBlendFunc expected %v got %v", ω.Out, Σ)
		}
	case *GlBlendFuncSeparate:
		Σ := GlBlendFuncSeparate_Out{}
		m.State.Blending.SrcRgbBlendFactor = ω.In.SrcFactorRgb
		m.State.Blending.DstRgbBlendFactor = ω.In.DstFactorRgb
		m.State.Blending.SrcAlphaBlendFactor = ω.In.SrcFactorAlpha
		m.State.Blending.DstAlphaBlendFactor = ω.In.DstFactorAlpha
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glBlendFuncSeparate expected %v got %v", ω.Out, Σ)
		}
	case *GlBlendEquation:
		Σ := GlBlendEquation_Out{}
		m.State.Blending.BlendEquationRgb = ω.In.Equation
		m.State.Blending.BlendEquationAlpha = ω.In.Equation
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glBlendEquation expected %v got %v", ω.Out, Σ)
		}
	case *GlBlendEquationSeparate:
		Σ := GlBlendEquationSeparate_Out{}
		m.State.Blending.BlendEquationRgb = ω.In.Rgb
		m.State.Blending.BlendEquationAlpha = ω.In.Alpha
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glBlendEquationSeparate expected %v got %v", ω.Out, Σ)
		}
	case *GlBlendColor:
		Σ := GlBlendColor_Out{}
		m.State.Blending.BlendColor = func() Color {
			s := Color{}
			s.Init()
			s.Red = ω.In.Red
			s.Green = ω.In.Green
			s.Blue = ω.In.Blue
			s.Alpha = ω.In.Alpha
			return s
		}()
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glBlendColor expected %v got %v", ω.Out, Σ)
		}
	case *GlEnableVertexAttribArray:
		Σ := GlEnableVertexAttribArray_Out{}
		a := m.State.VertexAttributeArrays.Get(ω.In.Location) // VertexAttributeArrayRef
		a.Enabled = true
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glEnableVertexAttribArray expected %v got %v", ω.Out, Σ)
		}
	case *GlDisableVertexAttribArray:
		Σ := GlDisableVertexAttribArray_Out{}
		a := m.State.VertexAttributeArrays.Get(ω.In.Location) // VertexAttributeArrayRef
		a.Enabled = false
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glDisableVertexAttribArray expected %v got %v", ω.Out, Σ)
		}
	case *GlVertexAttribPointer:
		Σ := GlVertexAttribPointer_Out{}
		a := m.State.VertexAttributeArrays.Get(ω.In.Location) // VertexAttributeArrayRef
		a.Size = ω.In.Size
		a.Type = ω.In.Type
		a.Normalized = ω.In.Normalized
		a.Stride = ω.In.Stride
		a.Data = memory.Pointer(ω.In.Data)
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glVertexAttribPointer expected %v got %v", ω.Out, Σ)
		}
	case *GlGetActiveAttrib:
		Σ := GlGetActiveAttrib_Out{}
		Σ.BufferBytesWritten = ω.Out.BufferBytesWritten
		Σ.VectorCount = ω.Out.VectorCount
		Σ.Type = ω.Out.Type
		Σ.Name = ω.Out.Name
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGetActiveAttrib expected %v got %v", ω.Out, Σ)
		}
	case *GlGetActiveUniform:
		Σ := GlGetActiveUniform_Out{}
		Σ.BufferBytesWritten = ω.Out.BufferBytesWritten
		Σ.Size = ω.Out.Size
		Σ.Type = ω.Out.Type
		Σ.Name = ω.Out.Name
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGetActiveUniform expected %v got %v", ω.Out, Σ)
		}
	case *GlGetError:
		Σ := GlGetError_Out{}
		Σ.Result = ω.Out.Result
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGetError expected %v got %v", ω.Out, Σ)
		}
	case *GlGetProgramiv:
		Σ := GlGetProgramiv_Out{}
		Σ.Value = make(S32Array, 1)
		Σ.Value = ω.Out.Value
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGetProgramiv expected %v got %v", ω.Out, Σ)
		}
	case *GlGetShaderiv:
		Σ := GlGetShaderiv_Out{}
		Σ.Value = make(S32Array, 1)
		s := m.State.Instances.Shaders.Get(ω.In.Shader) // ShaderRef
		Σ.Value[0] = func() (result int32) {
			switch ω.In.Parameter {
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
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGetShaderiv expected %v got %v", ω.Out, Σ)
		}
	case *GlGetUniformLocation:
		Σ := GlGetUniformLocation_Out{}
		Σ.Result = ω.Out.Result
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGetUniformLocation expected %v got %v", ω.Out, Σ)
		}
	case *GlGetAttribLocation:
		Σ := GlGetAttribLocation_Out{}
		Σ.Result = ω.Out.Result
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGetAttribLocation expected %v got %v", ω.Out, Σ)
		}
	case *GlPixelStorei:
		Σ := GlPixelStorei_Out{}
		m.State.PixelStorage[ω.In.Parameter] = ω.In.Value
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glPixelStorei expected %v got %v", ω.Out, Σ)
		}
	case *GlTexParameteri:
		Σ := GlTexParameteri_Out{}
		id := m.State.TextureUnits.Get(m.State.ActiveTextureUnit).Get(ω.In.Target) // TextureId
		t := m.State.Instances.Textures.Get(id)                                    // TextureRef
		switch ω.In.Parameter {
		case TextureParameter_GL_TEXTURE_MAG_FILTER:
			t.MagFilter = TextureFilterMode(ω.In.Value)
		case TextureParameter_GL_TEXTURE_MIN_FILTER:
			t.MinFilter = TextureFilterMode(ω.In.Value)
		case TextureParameter_GL_TEXTURE_WRAP_S:
			t.WrapS = TextureWrapMode(ω.In.Value)
		case TextureParameter_GL_TEXTURE_WRAP_T:
			t.WrapT = TextureWrapMode(ω.In.Value)
		case TextureParameter_GL_TEXTURE_MAX_ANISOTROPY_EXT:
			t.MaxAnisotropy = float32(ω.In.Value)
		case TextureParameter_GL_TEXTURE_SWIZZLE_R:
			t.SwizzleR = TexelComponent(ω.In.Value)
		case TextureParameter_GL_TEXTURE_SWIZZLE_G:
			t.SwizzleG = TexelComponent(ω.In.Value)
		case TextureParameter_GL_TEXTURE_SWIZZLE_B:
			t.SwizzleB = TexelComponent(ω.In.Value)
		case TextureParameter_GL_TEXTURE_SWIZZLE_A:
			t.SwizzleA = TexelComponent(ω.In.Value)
		default:
			// TODO: better unmatched handling
			v := ω.In.Parameter
			log.Printf("Error: Missing switch case handler for value %T %v", v, v)
		}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glTexParameteri expected %v got %v", ω.Out, Σ)
		}
	case *GlTexParameterf:
		Σ := GlTexParameterf_Out{}
		id := m.State.TextureUnits.Get(m.State.ActiveTextureUnit).Get(ω.In.Target) // TextureId
		t := m.State.Instances.Textures.Get(id)                                    // TextureRef
		switch ω.In.Parameter {
		case TextureParameter_GL_TEXTURE_MAG_FILTER:
			t.MagFilter = TextureFilterMode(ω.In.Value)
		case TextureParameter_GL_TEXTURE_MIN_FILTER:
			t.MinFilter = TextureFilterMode(ω.In.Value)
		case TextureParameter_GL_TEXTURE_WRAP_S:
			t.WrapS = TextureWrapMode(ω.In.Value)
		case TextureParameter_GL_TEXTURE_WRAP_T:
			t.WrapT = TextureWrapMode(ω.In.Value)
		case TextureParameter_GL_TEXTURE_MAX_ANISOTROPY_EXT:
			t.MaxAnisotropy = ω.In.Value
		case TextureParameter_GL_TEXTURE_SWIZZLE_R:
			t.SwizzleR = TexelComponent(ω.In.Value)
		case TextureParameter_GL_TEXTURE_SWIZZLE_G:
			t.SwizzleG = TexelComponent(ω.In.Value)
		case TextureParameter_GL_TEXTURE_SWIZZLE_B:
			t.SwizzleB = TexelComponent(ω.In.Value)
		case TextureParameter_GL_TEXTURE_SWIZZLE_A:
			t.SwizzleA = TexelComponent(ω.In.Value)
		default:
			// TODO: better unmatched handling
			v := ω.In.Parameter
			log.Printf("Error: Missing switch case handler for value %T %v", v, v)
		}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glTexParameterf expected %v got %v", ω.Out, Σ)
		}
	case *GlGetTexParameteriv:
		Σ := GlGetTexParameteriv_Out{}
		Σ.Values = make(S32Array, 1)
		id := m.State.TextureUnits.Get(m.State.ActiveTextureUnit).Get(ω.In.Target) // TextureId
		t := m.State.Instances.Textures.Get(id)                                    // TextureRef
		Σ.Values[0] = func() (result int32) {
			switch ω.In.Parameter {
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
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGetTexParameteriv expected %v got %v", ω.Out, Σ)
		}
	case *GlGetTexParameterfv:
		Σ := GlGetTexParameterfv_Out{}
		Σ.Values = make(F32Array, 1)
		id := m.State.TextureUnits.Get(m.State.ActiveTextureUnit).Get(ω.In.Target) // TextureId
		t := m.State.Instances.Textures.Get(id)                                    // TextureRef
		Σ.Values[0] = func() (result float32) {
			switch ω.In.Parameter {
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
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGetTexParameterfv expected %v got %v", ω.Out, Σ)
		}
	case *GlUniform1i:
		Σ := GlUniform1i_Out{}
		program := m.State.Instances.Programs.Get(m.State.BoundProgram) // ProgramRef
		uniform := program.Uniforms.Get(ω.In.Location)                  // Uniform
		uniform.Type = ShaderUniformType_GL_INT
		uniform.Value.S32 = ω.In.Value
		program.Uniforms[ω.In.Location] = uniform
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glUniform1i expected %v got %v", ω.Out, Σ)
		}
	case *GlUniform2i:
		Σ := GlUniform2i_Out{}
		program := m.State.Instances.Programs.Get(m.State.BoundProgram) // ProgramRef
		uniform := program.Uniforms.Get(ω.In.Location)                  // Uniform
		uniform.Type = ShaderUniformType_GL_INT_VEC2
		uniform.Value.Vec2i = func() Vec2i {
			s := Vec2i{}
			s.Init()
			s.X = ω.In.Value0
			s.Y = ω.In.Value1
			return s
		}()
		program.Uniforms[ω.In.Location] = uniform
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glUniform2i expected %v got %v", ω.Out, Σ)
		}
	case *GlUniform3i:
		Σ := GlUniform3i_Out{}
		program := m.State.Instances.Programs.Get(m.State.BoundProgram) // ProgramRef
		uniform := program.Uniforms.Get(ω.In.Location)                  // Uniform
		uniform.Type = ShaderUniformType_GL_INT_VEC3
		uniform.Value.Vec3i = func() Vec3i {
			s := Vec3i{}
			s.Init()
			s.X = ω.In.Value0
			s.Y = ω.In.Value1
			s.Z = ω.In.Value2
			return s
		}()
		program.Uniforms[ω.In.Location] = uniform
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glUniform3i expected %v got %v", ω.Out, Σ)
		}
	case *GlUniform4i:
		Σ := GlUniform4i_Out{}
		program := m.State.Instances.Programs.Get(m.State.BoundProgram) // ProgramRef
		uniform := program.Uniforms.Get(ω.In.Location)                  // Uniform
		uniform.Type = ShaderUniformType_GL_INT_VEC4
		uniform.Value.Vec4i = func() Vec4i {
			s := Vec4i{}
			s.Init()
			s.X = ω.In.Value0
			s.Y = ω.In.Value1
			s.Z = ω.In.Value2
			s.W = ω.In.Value3
			return s
		}()
		program.Uniforms[ω.In.Location] = uniform
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glUniform4i expected %v got %v", ω.Out, Σ)
		}
	case *GlUniform1iv:
		Σ := GlUniform1iv_Out{}
		program := m.State.Instances.Programs.Get(m.State.BoundProgram) // ProgramRef
		uniform := program.Uniforms.Get(ω.In.Location)                  // Uniform
		uniform.Type = ShaderUniformType_GL_INT
		uniform.Value.S32 = ω.In.Value[0]
		program.Uniforms[ω.In.Location] = uniform
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glUniform1iv expected %v got %v", ω.Out, Σ)
		}
	case *GlUniform2iv:
		Σ := GlUniform2iv_Out{}
		program := m.State.Instances.Programs.Get(m.State.BoundProgram) // ProgramRef
		uniform := program.Uniforms.Get(ω.In.Location)                  // Uniform
		uniform.Type = ShaderUniformType_GL_INT_VEC2
		uniform.Value.Vec2i = func() Vec2i {
			s := Vec2i{}
			s.Init()
			s.X = ω.In.Value[0]
			s.Y = ω.In.Value[1]
			return s
		}()
		program.Uniforms[ω.In.Location] = uniform
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glUniform2iv expected %v got %v", ω.Out, Σ)
		}
	case *GlUniform3iv:
		Σ := GlUniform3iv_Out{}
		program := m.State.Instances.Programs.Get(m.State.BoundProgram) // ProgramRef
		uniform := program.Uniforms.Get(ω.In.Location)                  // Uniform
		uniform.Type = ShaderUniformType_GL_INT_VEC3
		uniform.Value.Vec3i = func() Vec3i {
			s := Vec3i{}
			s.Init()
			s.X = ω.In.Value[0]
			s.Y = ω.In.Value[1]
			s.Z = ω.In.Value[2]
			return s
		}()
		program.Uniforms[ω.In.Location] = uniform
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glUniform3iv expected %v got %v", ω.Out, Σ)
		}
	case *GlUniform4iv:
		Σ := GlUniform4iv_Out{}
		program := m.State.Instances.Programs.Get(m.State.BoundProgram) // ProgramRef
		uniform := program.Uniforms.Get(ω.In.Location)                  // Uniform
		uniform.Type = ShaderUniformType_GL_INT_VEC4
		uniform.Value.Vec4i = func() Vec4i {
			s := Vec4i{}
			s.Init()
			s.X = ω.In.Value[0]
			s.Y = ω.In.Value[1]
			s.Z = ω.In.Value[2]
			s.W = ω.In.Value[3]
			return s
		}()
		program.Uniforms[ω.In.Location] = uniform
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glUniform4iv expected %v got %v", ω.Out, Σ)
		}
	case *GlUniform1f:
		Σ := GlUniform1f_Out{}
		program := m.State.Instances.Programs.Get(m.State.BoundProgram) // ProgramRef
		uniform := program.Uniforms.Get(ω.In.Location)                  // Uniform
		uniform.Type = ShaderUniformType_GL_FLOAT
		uniform.Value.F32 = ω.In.Value
		program.Uniforms[ω.In.Location] = uniform
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glUniform1f expected %v got %v", ω.Out, Σ)
		}
	case *GlUniform2f:
		Σ := GlUniform2f_Out{}
		program := m.State.Instances.Programs.Get(m.State.BoundProgram) // ProgramRef
		uniform := program.Uniforms.Get(ω.In.Location)                  // Uniform
		uniform.Type = ShaderUniformType_GL_FLOAT_VEC2
		uniform.Value.Vec2f = func() Vec2f {
			s := Vec2f{}
			s.Init()
			s.X = ω.In.Value0
			s.Y = ω.In.Value1
			return s
		}()
		program.Uniforms[ω.In.Location] = uniform
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glUniform2f expected %v got %v", ω.Out, Σ)
		}
	case *GlUniform3f:
		Σ := GlUniform3f_Out{}
		program := m.State.Instances.Programs.Get(m.State.BoundProgram) // ProgramRef
		uniform := program.Uniforms.Get(ω.In.Location)                  // Uniform
		uniform.Type = ShaderUniformType_GL_FLOAT_VEC3
		uniform.Value.Vec3f = func() Vec3f {
			s := Vec3f{}
			s.Init()
			s.X = ω.In.Value0
			s.Y = ω.In.Value1
			s.Z = ω.In.Value2
			return s
		}()
		program.Uniforms[ω.In.Location] = uniform
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glUniform3f expected %v got %v", ω.Out, Σ)
		}
	case *GlUniform4f:
		Σ := GlUniform4f_Out{}
		program := m.State.Instances.Programs.Get(m.State.BoundProgram) // ProgramRef
		uniform := program.Uniforms.Get(ω.In.Location)                  // Uniform
		uniform.Type = ShaderUniformType_GL_FLOAT_VEC4
		uniform.Value.Vec4f = func() Vec4f {
			s := Vec4f{}
			s.Init()
			s.X = ω.In.Value0
			s.Y = ω.In.Value1
			s.Z = ω.In.Value2
			s.W = ω.In.Value3
			return s
		}()
		program.Uniforms[ω.In.Location] = uniform
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glUniform4f expected %v got %v", ω.Out, Σ)
		}
	case *GlUniform1fv:
		Σ := GlUniform1fv_Out{}
		program := m.State.Instances.Programs.Get(m.State.BoundProgram) // ProgramRef
		uniform := program.Uniforms.Get(ω.In.Location)                  // Uniform
		uniform.Type = ShaderUniformType_GL_FLOAT
		uniform.Value.F32 = ω.In.Value[0]
		program.Uniforms[ω.In.Location] = uniform
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glUniform1fv expected %v got %v", ω.Out, Σ)
		}
	case *GlUniform2fv:
		Σ := GlUniform2fv_Out{}
		program := m.State.Instances.Programs.Get(m.State.BoundProgram) // ProgramRef
		uniform := program.Uniforms.Get(ω.In.Location)                  // Uniform
		uniform.Type = ShaderUniformType_GL_FLOAT_VEC2
		uniform.Value.Vec2f = func() Vec2f {
			s := Vec2f{}
			s.Init()
			s.X = ω.In.Value[0]
			s.Y = ω.In.Value[1]
			return s
		}()
		program.Uniforms[ω.In.Location] = uniform
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glUniform2fv expected %v got %v", ω.Out, Σ)
		}
	case *GlUniform3fv:
		Σ := GlUniform3fv_Out{}
		program := m.State.Instances.Programs.Get(m.State.BoundProgram) // ProgramRef
		uniform := program.Uniforms.Get(ω.In.Location)                  // Uniform
		uniform.Type = ShaderUniformType_GL_FLOAT_VEC3
		uniform.Value.Vec3f = func() Vec3f {
			s := Vec3f{}
			s.Init()
			s.X = ω.In.Value[0]
			s.Y = ω.In.Value[1]
			s.Z = ω.In.Value[2]
			return s
		}()
		program.Uniforms[ω.In.Location] = uniform
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glUniform3fv expected %v got %v", ω.Out, Σ)
		}
	case *GlUniform4fv:
		Σ := GlUniform4fv_Out{}
		program := m.State.Instances.Programs.Get(m.State.BoundProgram) // ProgramRef
		uniform := program.Uniforms.Get(ω.In.Location)                  // Uniform
		uniform.Type = ShaderUniformType_GL_FLOAT_VEC4
		uniform.Value.Vec4f = func() Vec4f {
			s := Vec4f{}
			s.Init()
			s.X = ω.In.Value[0]
			s.Y = ω.In.Value[1]
			s.Z = ω.In.Value[2]
			s.W = ω.In.Value[3]
			return s
		}()
		program.Uniforms[ω.In.Location] = uniform
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glUniform4fv expected %v got %v", ω.Out, Σ)
		}
	case *GlUniformMatrix2fv:
		Σ := GlUniformMatrix2fv_Out{}
		program := m.State.Instances.Programs.Get(m.State.BoundProgram) // ProgramRef
		uniform := program.Uniforms.Get(ω.In.Location)                  // Uniform
		uniform.Type = ShaderUniformType_GL_FLOAT_MAT2
		uniform.Value.Mat2f = func() Mat2f {
			s := Mat2f{}
			s.Init()
			s.Col0 = func() Vec2f {
				s := Vec2f{}
				s.Init()
				s.X = ω.In.Values[0]
				s.Y = ω.In.Values[1]
				return s
			}()
			s.Col1 = func() Vec2f {
				s := Vec2f{}
				s.Init()
				s.X = ω.In.Values[0]
				s.Y = ω.In.Values[1]
				return s
			}()
			return s
		}()
		program.Uniforms[ω.In.Location] = uniform
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glUniformMatrix2fv expected %v got %v", ω.Out, Σ)
		}
	case *GlUniformMatrix3fv:
		Σ := GlUniformMatrix3fv_Out{}
		program := m.State.Instances.Programs.Get(m.State.BoundProgram) // ProgramRef
		uniform := program.Uniforms.Get(ω.In.Location)                  // Uniform
		uniform.Type = ShaderUniformType_GL_FLOAT_MAT3
		uniform.Value.Mat3f = func() Mat3f {
			s := Mat3f{}
			s.Init()
			s.Col0 = func() Vec3f {
				s := Vec3f{}
				s.Init()
				s.X = ω.In.Values[0]
				s.Y = ω.In.Values[1]
				s.Z = ω.In.Values[2]
				return s
			}()
			s.Col1 = func() Vec3f {
				s := Vec3f{}
				s.Init()
				s.X = ω.In.Values[3]
				s.Y = ω.In.Values[4]
				s.Z = ω.In.Values[5]
				return s
			}()
			s.Col2 = func() Vec3f {
				s := Vec3f{}
				s.Init()
				s.X = ω.In.Values[6]
				s.Y = ω.In.Values[7]
				s.Z = ω.In.Values[8]
				return s
			}()
			return s
		}()
		program.Uniforms[ω.In.Location] = uniform
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glUniformMatrix3fv expected %v got %v", ω.Out, Σ)
		}
	case *GlUniformMatrix4fv:
		Σ := GlUniformMatrix4fv_Out{}
		program := m.State.Instances.Programs.Get(m.State.BoundProgram) // ProgramRef
		uniform := program.Uniforms.Get(ω.In.Location)                  // Uniform
		uniform.Value.Mat4f = func() Mat4f {
			s := Mat4f{}
			s.Init()
			s.Col0 = func() Vec4f {
				s := Vec4f{}
				s.Init()
				s.X = ω.In.Values[0]
				s.Y = ω.In.Values[1]
				s.Z = ω.In.Values[2]
				s.W = ω.In.Values[3]
				return s
			}()
			s.Col1 = func() Vec4f {
				s := Vec4f{}
				s.Init()
				s.X = ω.In.Values[4]
				s.Y = ω.In.Values[5]
				s.Z = ω.In.Values[6]
				s.W = ω.In.Values[7]
				return s
			}()
			s.Col2 = func() Vec4f {
				s := Vec4f{}
				s.Init()
				s.X = ω.In.Values[8]
				s.Y = ω.In.Values[9]
				s.Z = ω.In.Values[10]
				s.W = ω.In.Values[11]
				return s
			}()
			s.Col3 = func() Vec4f {
				s := Vec4f{}
				s.Init()
				s.X = ω.In.Values[12]
				s.Y = ω.In.Values[13]
				s.Z = ω.In.Values[14]
				s.W = ω.In.Values[15]
				return s
			}()
			return s
		}()
		program.Uniforms[ω.In.Location] = uniform
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glUniformMatrix4fv expected %v got %v", ω.Out, Σ)
		}
	case *GlGetUniformfv:
		Σ := GlGetUniformfv_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGetUniformfv expected %v got %v", ω.Out, Σ)
		}
	case *GlGetUniformiv:
		Σ := GlGetUniformiv_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGetUniformiv expected %v got %v", ω.Out, Σ)
		}
	case *GlVertexAttrib1f:
		Σ := GlVertexAttrib1f_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glVertexAttrib1f expected %v got %v", ω.Out, Σ)
		}
	case *GlVertexAttrib2f:
		Σ := GlVertexAttrib2f_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glVertexAttrib2f expected %v got %v", ω.Out, Σ)
		}
	case *GlVertexAttrib3f:
		Σ := GlVertexAttrib3f_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glVertexAttrib3f expected %v got %v", ω.Out, Σ)
		}
	case *GlVertexAttrib4f:
		Σ := GlVertexAttrib4f_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glVertexAttrib4f expected %v got %v", ω.Out, Σ)
		}
	case *GlVertexAttrib1fv:
		Σ := GlVertexAttrib1fv_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glVertexAttrib1fv expected %v got %v", ω.Out, Σ)
		}
	case *GlVertexAttrib2fv:
		Σ := GlVertexAttrib2fv_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glVertexAttrib2fv expected %v got %v", ω.Out, Σ)
		}
	case *GlVertexAttrib3fv:
		Σ := GlVertexAttrib3fv_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glVertexAttrib3fv expected %v got %v", ω.Out, Σ)
		}
	case *GlVertexAttrib4fv:
		Σ := GlVertexAttrib4fv_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glVertexAttrib4fv expected %v got %v", ω.Out, Σ)
		}
	case *GlGetShaderPrecisionFormat:
		Σ := GlGetShaderPrecisionFormat_Out{}
		Σ.Range = make(S32Array, 2)
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGetShaderPrecisionFormat expected %v got %v", ω.Out, Σ)
		}
	case *GlDepthMask:
		Σ := GlDepthMask_Out{}
		m.State.Rasterizing.DepthMask = ω.In.Enabled
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glDepthMask expected %v got %v", ω.Out, Σ)
		}
	case *GlDepthFunc:
		Σ := GlDepthFunc_Out{}
		m.State.Rasterizing.DepthTestFunction = ω.In.Function
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glDepthFunc expected %v got %v", ω.Out, Σ)
		}
	case *GlDepthRangef:
		Σ := GlDepthRangef_Out{}
		m.State.Rasterizing.DepthNear = ω.In.Near
		m.State.Rasterizing.DepthFar = ω.In.Far
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glDepthRangef expected %v got %v", ω.Out, Σ)
		}
	case *GlColorMask:
		Σ := GlColorMask_Out{}
		m.State.Rasterizing.ColorMaskRed = ω.In.Red
		m.State.Rasterizing.ColorMaskGreen = ω.In.Green
		m.State.Rasterizing.ColorMaskBlue = ω.In.Blue
		m.State.Rasterizing.ColorMaskAlpha = ω.In.Alpha
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glColorMask expected %v got %v", ω.Out, Σ)
		}
	case *GlStencilMask:
		Σ := GlStencilMask_Out{}
		m.State.Rasterizing.StencilMask[FaceMode_GL_FRONT] = ω.In.Mask
		m.State.Rasterizing.StencilMask[FaceMode_GL_BACK] = ω.In.Mask
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glStencilMask expected %v got %v", ω.Out, Σ)
		}
	case *GlStencilMaskSeparate:
		Σ := GlStencilMaskSeparate_Out{}
		switch ω.In.Face {
		case FaceMode_GL_FRONT:
			m.State.Rasterizing.StencilMask[FaceMode_GL_FRONT] = ω.In.Mask
		case FaceMode_GL_BACK:
			m.State.Rasterizing.StencilMask[FaceMode_GL_BACK] = ω.In.Mask
		case FaceMode_GL_FRONT_AND_BACK:
			m.State.Rasterizing.StencilMask[FaceMode_GL_FRONT] = ω.In.Mask
			m.State.Rasterizing.StencilMask[FaceMode_GL_BACK] = ω.In.Mask
		default:
			// TODO: better unmatched handling
			v := ω.In.Face
			log.Printf("Error: Missing switch case handler for value %T %v", v, v)
		}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glStencilMaskSeparate expected %v got %v", ω.Out, Σ)
		}
	case *GlStencilFuncSeparate:
		Σ := GlStencilFuncSeparate_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glStencilFuncSeparate expected %v got %v", ω.Out, Σ)
		}
	case *GlStencilOpSeparate:
		Σ := GlStencilOpSeparate_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glStencilOpSeparate expected %v got %v", ω.Out, Σ)
		}
	case *GlFrontFace:
		Σ := GlFrontFace_Out{}
		m.State.Rasterizing.FrontFace = ω.In.Orientation
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glFrontFace expected %v got %v", ω.Out, Σ)
		}
	case *GlViewport:
		Σ := GlViewport_Out{}
		m.State.Rasterizing.Viewport = func() Rect {
			s := Rect{}
			s.Init()
			s.X = ω.In.X
			s.Y = ω.In.Y
			s.Width = ω.In.Width
			s.Height = ω.In.Height
			return s
		}()
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glViewport expected %v got %v", ω.Out, Σ)
		}
	case *GlScissor:
		Σ := GlScissor_Out{}
		m.State.Rasterizing.Scissor = func() Rect {
			s := Rect{}
			s.Init()
			s.X = ω.In.X
			s.Y = ω.In.Y
			s.Width = ω.In.Width
			s.Height = ω.In.Height
			return s
		}()
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glScissor expected %v got %v", ω.Out, Σ)
		}
	case *GlActiveTexture:
		Σ := GlActiveTexture_Out{}
		m.State.ActiveTextureUnit = ω.In.Unit
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glActiveTexture expected %v got %v", ω.Out, Σ)
		}
	case *GlGenTextures:
		Σ := GlGenTextures_Out{}
		Σ.Textures = make(TextureIdArray, ω.In.Count)
		for i := int32(0); i < ω.In.Count; i++ {
			id := TextureId(ω.Out.Textures[i]) // TextureId
			m.State.Instances.Textures[id] = func() *Texture {
				s := &Texture{}
				s.Init()
				return s
			}()
			Σ.Textures[i] = id
		}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGenTextures expected %v got %v", ω.Out, Σ)
		}
	case *GlDeleteTextures:
		Σ := GlDeleteTextures_Out{}
		for i := int32(0); i < ω.In.Count; i++ {
			m.State.Instances.Textures.Delete(ω.In.Textures[i])
		}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glDeleteTextures expected %v got %v", ω.Out, Σ)
		}
	case *GlIsTexture:
		Σ := GlIsTexture_Out{}
		Σ.Result = m.State.Instances.Textures.Contains(ω.In.Texture)
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glIsTexture expected %v got %v", ω.Out, Σ)
		}
	case *GlBindTexture:
		Σ := GlBindTexture_Out{}
		if (m.State.Instances.Textures.Contains(ω.In.Texture)) == (false) {
			m.State.Instances.Textures[ω.In.Texture] = func() *Texture {
				s := &Texture{}
				s.Init()
				return s
			}()
		}
		m.State.TextureUnits.Get(m.State.ActiveTextureUnit)[ω.In.Target] = ω.In.Texture
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glBindTexture expected %v got %v", ω.Out, Σ)
		}
	case *GlTexImage2D:
		Σ := GlTexImage2D_Out{}
		switch ω.In.Target {
		case TextureImageTarget_GL_TEXTURE_2D:
			id := m.State.TextureUnits.Get(m.State.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_2D) // TextureId
			t := m.State.Instances.Textures.Get(id)                                                    // TextureRef
			l := func() Image {
				s := Image{}
				s.Init()
				s.Width = ω.In.Width
				s.Height = ω.In.Height
				s.Size = imageSize(ω.In.Width, ω.In.Height, ω.In.Format, ω.In.Type)
				s.Format = ImageTexelFormat(ω.In.Format)
				return s
			}() // Image
			l.Data.Write(m.State.Mem.Slice(memory.Range{
				Base: memory.Pointer(ω.In.Data),
				Size: uint64(l.Size),
			}))
			t.Texture2D[ω.In.Level] = l
			t.Kind = TextureKind_TEXTURE2D
			t.Format = ImageTexelFormat(ω.In.Format)
		case TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X, TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Z, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z:
			id := m.State.TextureUnits.Get(m.State.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_CUBE_MAP) // TextureId
			t := m.State.Instances.Textures.Get(id)                                                          // TextureRef
			l := func() Image {
				s := Image{}
				s.Init()
				s.Width = ω.In.Width
				s.Height = ω.In.Height
				s.Size = imageSize(ω.In.Width, ω.In.Height, ω.In.Format, ω.In.Type)
				s.Format = ImageTexelFormat(ω.In.Format)
				return s
			}() // Image
			cube := t.Cubemap.Get(ω.In.Level) // CubemapLevel
			cube.Faces[CubeMapImageTarget(ω.In.Target)] = l
			t.Cubemap[ω.In.Level] = cube
			t.Kind = TextureKind_CUBEMAP
			t.Format = ImageTexelFormat(ω.In.Format)
		default:
			// TODO: better unmatched handling
			v := ω.In.Target
			log.Printf("Error: Missing switch case handler for value %T %v", v, v)
		}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glTexImage2D expected %v got %v", ω.Out, Σ)
		}
	case *GlTexSubImage2D:
		Σ := GlTexSubImage2D_Out{}
		switch ω.In.Target {
		case TextureImageTarget_GL_TEXTURE_2D:
			id := m.State.TextureUnits.Get(m.State.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_2D) // TextureId
			t := m.State.Instances.Textures.Get(id)                                                    // TextureRef
			l := func() Image {
				s := Image{}
				s.Init()
				s.Width = ω.In.Width
				s.Height = ω.In.Height
				s.Size = imageSize(ω.In.Width, ω.In.Height, ω.In.Format, ω.In.Type)
				s.Format = ImageTexelFormat(ω.In.Format)
				return s
			}() // Image
			l.Data.Write(m.State.Mem.Slice(memory.Range{
				Base: memory.Pointer(ω.In.Data),
				Size: uint64(l.Size),
			}))
			t.Texture2D[ω.In.Level] = l
			t.Kind = TextureKind_TEXTURE2D
			t.Format = ImageTexelFormat(ω.In.Format)
		case TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X, TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Z, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z:
			id := m.State.TextureUnits.Get(m.State.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_CUBE_MAP) // TextureId
			t := m.State.Instances.Textures.Get(id)                                                          // TextureRef
			l := func() Image {
				s := Image{}
				s.Init()
				s.Width = ω.In.Width
				s.Height = ω.In.Height
				s.Size = imageSize(ω.In.Width, ω.In.Height, ω.In.Format, ω.In.Type)
				s.Format = ImageTexelFormat(ω.In.Format)
				return s
			}() // Image
			l.Data.Write(m.State.Mem.Slice(memory.Range{
				Base: memory.Pointer(ω.In.Data),
				Size: uint64(l.Size),
			}))
			cube := t.Cubemap.Get(ω.In.Level) // CubemapLevel
			cube.Faces[CubeMapImageTarget(ω.In.Target)] = l
			t.Cubemap[ω.In.Level] = cube
			t.Kind = TextureKind_CUBEMAP
			t.Format = ImageTexelFormat(ω.In.Format)
		default:
			// TODO: better unmatched handling
			v := ω.In.Target
			log.Printf("Error: Missing switch case handler for value %T %v", v, v)
		}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glTexSubImage2D expected %v got %v", ω.Out, Σ)
		}
	case *GlCopyTexImage2D:
		Σ := GlCopyTexImage2D_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glCopyTexImage2D expected %v got %v", ω.Out, Σ)
		}
	case *GlCopyTexSubImage2D:
		Σ := GlCopyTexSubImage2D_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glCopyTexSubImage2D expected %v got %v", ω.Out, Σ)
		}
	case *GlCompressedTexImage2D:
		Σ := GlCompressedTexImage2D_Out{}
		switch ω.In.Target {
		case TextureImageTarget_GL_TEXTURE_2D:
			id := m.State.TextureUnits.Get(m.State.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_2D) // TextureId
			t := m.State.Instances.Textures.Get(id)                                                    // TextureRef
			l := func() Image {
				s := Image{}
				s.Init()
				s.Width = ω.In.Width
				s.Height = ω.In.Height
				s.Size = ω.In.ImageSize
				s.Format = ImageTexelFormat(ω.In.Format)
				return s
			}() // Image
			l.Data.Write(m.State.Mem.Slice(memory.Range{
				Base: memory.Pointer(ω.In.Data),
				Size: uint64(l.Size),
			}))
			t.Texture2D[ω.In.Level] = l
			t.Kind = TextureKind_TEXTURE2D
			t.Format = ImageTexelFormat(ω.In.Format)
		case TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X, TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Z, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z:
			id := m.State.TextureUnits.Get(m.State.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_CUBE_MAP) // TextureId
			t := m.State.Instances.Textures.Get(id)                                                          // TextureRef
			l := func() Image {
				s := Image{}
				s.Init()
				s.Width = ω.In.Width
				s.Height = ω.In.Height
				s.Size = ω.In.ImageSize
				s.Format = ImageTexelFormat(ω.In.Format)
				return s
			}() // Image
			l.Data.Write(m.State.Mem.Slice(memory.Range{
				Base: memory.Pointer(ω.In.Data),
				Size: uint64(l.Size),
			}))
			cube := t.Cubemap.Get(ω.In.Level) // CubemapLevel
			cube.Faces[CubeMapImageTarget(ω.In.Target)] = l
			t.Cubemap[ω.In.Level] = cube
			t.Kind = TextureKind_CUBEMAP
			t.Format = ImageTexelFormat(ω.In.Format)
		default:
			// TODO: better unmatched handling
			v := ω.In.Target
			log.Printf("Error: Missing switch case handler for value %T %v", v, v)
		}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glCompressedTexImage2D expected %v got %v", ω.Out, Σ)
		}
	case *GlCompressedTexSubImage2D:
		Σ := GlCompressedTexSubImage2D_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glCompressedTexSubImage2D expected %v got %v", ω.Out, Σ)
		}
	case *GlGenerateMipmap:
		Σ := GlGenerateMipmap_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGenerateMipmap expected %v got %v", ω.Out, Σ)
		}
	case *GlReadPixels:
		Σ := GlReadPixels_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glReadPixels expected %v got %v", ω.Out, Σ)
		}
	case *GlGenFramebuffers:
		Σ := GlGenFramebuffers_Out{}
		Σ.Framebuffers = make(FramebufferIdArray, ω.In.Count)
		for i := int32(0); i < ω.In.Count; i++ {
			id := FramebufferId(ω.Out.Framebuffers[i]) // FramebufferId
			m.State.Instances.Framebuffers[id] = func() *Framebuffer {
				s := &Framebuffer{}
				s.Init()
				return s
			}()
			Σ.Framebuffers[i] = id
		}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGenFramebuffers expected %v got %v", ω.Out, Σ)
		}
	case *GlBindFramebuffer:
		Σ := GlBindFramebuffer_Out{}
		if (m.State.Instances.Framebuffers.Contains(ω.In.Framebuffer)) == (false) {
			m.State.Instances.Framebuffers[ω.In.Framebuffer] = func() *Framebuffer {
				s := &Framebuffer{}
				s.Init()
				return s
			}()
		}
		m.State.BoundFramebuffers[ω.In.Target] = ω.In.Framebuffer
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glBindFramebuffer expected %v got %v", ω.Out, Σ)
		}
	case *GlCheckFramebufferStatus:
		Σ := GlCheckFramebufferStatus_Out{}
		Σ.Result = ω.Out.Result
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glCheckFramebufferStatus expected %v got %v", ω.Out, Σ)
		}
	case *GlDeleteFramebuffers:
		Σ := GlDeleteFramebuffers_Out{}
		for i := int32(0); i < ω.In.Count; i++ {
			m.State.Instances.Framebuffers.Delete(ω.In.Framebuffers[i])
		}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glDeleteFramebuffers expected %v got %v", ω.Out, Σ)
		}
	case *GlIsFramebuffer:
		Σ := GlIsFramebuffer_Out{}
		Σ.Result = m.State.Instances.Framebuffers.Contains(ω.In.Framebuffer)
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glIsFramebuffer expected %v got %v", ω.Out, Σ)
		}
	case *GlGenRenderbuffers:
		Σ := GlGenRenderbuffers_Out{}
		Σ.Renderbuffers = make(RenderbufferIdArray, ω.In.Count)
		for i := int32(0); i < ω.In.Count; i++ {
			id := RenderbufferId(ω.Out.Renderbuffers[i]) // RenderbufferId
			m.State.Instances.Renderbuffers[id] = func() *Renderbuffer {
				s := &Renderbuffer{}
				s.Init()
				return s
			}()
			Σ.Renderbuffers[i] = id
		}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGenRenderbuffers expected %v got %v", ω.Out, Σ)
		}
	case *GlBindRenderbuffer:
		Σ := GlBindRenderbuffer_Out{}
		if (m.State.Instances.Renderbuffers.Contains(ω.In.Renderbuffer)) == (false) {
			m.State.Instances.Renderbuffers[ω.In.Renderbuffer] = func() *Renderbuffer {
				s := &Renderbuffer{}
				s.Init()
				return s
			}()
		}
		m.State.BoundRenderbuffers[ω.In.Target] = ω.In.Renderbuffer
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glBindRenderbuffer expected %v got %v", ω.Out, Σ)
		}
	case *GlRenderbufferStorage:
		Σ := GlRenderbufferStorage_Out{}
		id := m.State.BoundRenderbuffers.Get(ω.In.Target) // RenderbufferId
		rb := m.State.Instances.Renderbuffers.Get(id)     // RenderbufferRef
		rb.Format = ω.In.Format
		rb.Width = ω.In.Width
		rb.Height = ω.In.Height
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glRenderbufferStorage expected %v got %v", ω.Out, Σ)
		}
	case *GlDeleteRenderbuffers:
		Σ := GlDeleteRenderbuffers_Out{}
		for i := int32(0); i < ω.In.Count; i++ {
			m.State.Instances.Renderbuffers.Delete(ω.In.Renderbuffers[i])
		}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glDeleteRenderbuffers expected %v got %v", ω.Out, Σ)
		}
	case *GlIsRenderbuffer:
		Σ := GlIsRenderbuffer_Out{}
		Σ.Result = m.State.Instances.Renderbuffers.Contains(ω.In.Renderbuffer)
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glIsRenderbuffer expected %v got %v", ω.Out, Σ)
		}
	case *GlGetRenderbufferParameteriv:
		Σ := GlGetRenderbufferParameteriv_Out{}
		Σ.Values = make(S32Array, 1)
		id := m.State.BoundRenderbuffers.Get(ω.In.Target) // RenderbufferId
		rb := m.State.Instances.Renderbuffers.Get(id)     // RenderbufferRef
		Σ.Values[0] = func() (result int32) {
			switch ω.In.Parameter {
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
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGetRenderbufferParameteriv expected %v got %v", ω.Out, Σ)
		}
	case *GlGenBuffers:
		Σ := GlGenBuffers_Out{}
		Σ.Buffers = make(BufferIdArray, ω.In.Count)
		for i := int32(0); i < ω.In.Count; i++ {
			id := BufferId(ω.Out.Buffers[i]) // BufferId
			m.State.Instances.Buffers[id] = func() *Buffer {
				s := &Buffer{}
				s.Init()
				return s
			}()
			Σ.Buffers[i] = id
		}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGenBuffers expected %v got %v", ω.Out, Σ)
		}
	case *GlBindBuffer:
		Σ := GlBindBuffer_Out{}
		if (m.State.Instances.Buffers.Contains(ω.In.Buffer)) == (false) {
			m.State.Instances.Buffers[ω.In.Buffer] = func() *Buffer {
				s := &Buffer{}
				s.Init()
				return s
			}()
		}
		m.State.BoundBuffers[ω.In.Target] = ω.In.Buffer
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glBindBuffer expected %v got %v", ω.Out, Σ)
		}
	case *GlBufferData:
		Σ := GlBufferData_Out{}
		id := m.State.BoundBuffers.Get(ω.In.Target) // BufferId
		b := m.State.Instances.Buffers.Get(id)      // BufferRef
		b.Data.Write(m.State.Mem.Slice(memory.Range{
			Base: memory.Pointer(ω.In.Data),
			Size: uint64(ω.In.Size),
		}))
		b.Size = ω.In.Size
		b.Usage = ω.In.Usage
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glBufferData expected %v got %v", ω.Out, Σ)
		}
	case *GlBufferSubData:
		Σ := GlBufferSubData_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glBufferSubData expected %v got %v", ω.Out, Σ)
		}
	case *GlDeleteBuffers:
		Σ := GlDeleteBuffers_Out{}
		for i := int32(0); i < ω.In.Count; i++ {
			m.State.Instances.Buffers.Delete(ω.In.Buffers[i])
		}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glDeleteBuffers expected %v got %v", ω.Out, Σ)
		}
	case *GlIsBuffer:
		Σ := GlIsBuffer_Out{}
		Σ.Result = m.State.Instances.Buffers.Contains(ω.In.Buffer)
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glIsBuffer expected %v got %v", ω.Out, Σ)
		}
	case *GlGetBufferParameteriv:
		Σ := GlGetBufferParameteriv_Out{}
		id := m.State.BoundBuffers.Get(ω.In.Target) // BufferId
		b := m.State.Instances.Buffers.Get(id)      // BufferRef
		Σ.Value = func() (result int32) {
			switch ω.In.Parameter {
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
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGetBufferParameteriv expected %v got %v", ω.Out, Σ)
		}
	case *GlCreateShader:
		Σ := GlCreateShader_Out{}
		id := ShaderId(ω.Out.Result) // ShaderId
		m.State.Instances.Shaders[id] = func() *Shader {
			s := &Shader{}
			s.Init()
			return s
		}()
		s := m.State.Instances.Shaders.Get(id) // ShaderRef
		s.Type = ω.In.Type
		Σ.Result = id
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glCreateShader expected %v got %v", ω.Out, Σ)
		}
	case *GlDeleteShader:
		Σ := GlDeleteShader_Out{}
		s := m.State.Instances.Shaders.Get(ω.In.Shader) // ShaderRef
		s.Deletable = true
		m.State.Instances.Shaders.Delete(ω.In.Shader)
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glDeleteShader expected %v got %v", ω.Out, Σ)
		}
	case *GlShaderSource:
		Σ := GlShaderSource_Out{}
		s := m.State.Instances.Shaders.Get(ω.In.Shader) // ShaderRef
		s.Source = ω.In.Source
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glShaderSource expected %v got %v", ω.Out, Σ)
		}
	case *GlShaderBinary:
		Σ := GlShaderBinary_Out{}
		for i := int32(0); i < ω.In.Count; i++ {
		}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glShaderBinary expected %v got %v", ω.Out, Σ)
		}
	case *GlGetShaderInfoLog:
		Σ := GlGetShaderInfoLog_Out{}
		s := m.State.Instances.Shaders.Get(ω.In.Shader) // ShaderRef
		min_0_a := ω.In.BufferLength                    // s32
		min_0_b := strlen(s.InfoLog)                    // s32
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
		Σ.StringLengthWritten = min_0_result
		Σ.Info = substr(s.InfoLog, 0, Σ.StringLengthWritten)
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGetShaderInfoLog expected %v got %v", ω.Out, Σ)
		}
	case *GlGetShaderSource:
		Σ := GlGetShaderSource_Out{}
		s := m.State.Instances.Shaders.Get(ω.In.Shader) // ShaderRef
		min_1_a := ω.In.BufferLength                    // s32
		min_1_b := strlen(s.Source[0])                  // s32
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
		Σ.StringLengthWritten = min_1_result
		Σ.Source = substr(s.Source[0], 0, Σ.StringLengthWritten)
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGetShaderSource expected %v got %v", ω.Out, Σ)
		}
	case *GlReleaseShaderCompiler:
		Σ := GlReleaseShaderCompiler_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glReleaseShaderCompiler expected %v got %v", ω.Out, Σ)
		}
	case *GlCompileShader:
		Σ := GlCompileShader_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glCompileShader expected %v got %v", ω.Out, Σ)
		}
	case *GlIsShader:
		Σ := GlIsShader_Out{}
		Σ.Result = m.State.Instances.Shaders.Contains(ω.In.Shader)
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glIsShader expected %v got %v", ω.Out, Σ)
		}
	case *GlCreateProgram:
		Σ := GlCreateProgram_Out{}
		id := ProgramId(ω.Out.Result) // ProgramId
		m.State.Instances.Programs[id] = func() *Program {
			s := &Program{}
			s.Init()
			return s
		}()
		Σ.Result = id
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glCreateProgram expected %v got %v", ω.Out, Σ)
		}
	case *GlDeleteProgram:
		Σ := GlDeleteProgram_Out{}
		m.State.Instances.Programs.Delete(ω.In.Program)
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glDeleteProgram expected %v got %v", ω.Out, Σ)
		}
	case *GlAttachShader:
		Σ := GlAttachShader_Out{}
		p := m.State.Instances.Programs.Get(ω.In.Program) // ProgramRef
		s := m.State.Instances.Shaders.Get(ω.In.Shader)   // ShaderRef
		p.Shaders[s.Type] = ω.In.Shader
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glAttachShader expected %v got %v", ω.Out, Σ)
		}
	case *GlDetachShader:
		Σ := GlDetachShader_Out{}
		p := m.State.Instances.Programs.Get(ω.In.Program) // ProgramRef
		s := m.State.Instances.Shaders.Get(ω.In.Shader)   // ShaderRef
		p.Shaders.Delete(s.Type)
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glDetachShader expected %v got %v", ω.Out, Σ)
		}
	case *GlGetAttachedShaders:
		Σ := GlGetAttachedShaders_Out{}
		Σ.Shaders = make(ShaderIdArray, ω.In.BufferLength)
		p := m.State.Instances.Programs.Get(ω.In.Program) // ProgramRef
		min_2_a := ω.In.BufferLength                      // s32
		min_2_b := int32(len(p.Shaders))                  // s32
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
		Σ.ShadersLengthWritten = min_2_result
		Σ.Shaders = p.Shaders.Range()
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGetAttachedShaders expected %v got %v", ω.Out, Σ)
		}
	case *GlLinkProgram:
		Σ := GlLinkProgram_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glLinkProgram expected %v got %v", ω.Out, Σ)
		}
	case *GlGetProgramInfoLog:
		Σ := GlGetProgramInfoLog_Out{}
		p := m.State.Instances.Programs.Get(ω.In.Program) // ProgramRef
		min_3_a := ω.In.BufferLength                      // s32
		min_3_b := strlen(p.InfoLog)                      // s32
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
		Σ.StringLengthWritten = min_3_result
		Σ.Info = substr(p.InfoLog, 0, Σ.StringLengthWritten)
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGetProgramInfoLog expected %v got %v", ω.Out, Σ)
		}
	case *GlUseProgram:
		Σ := GlUseProgram_Out{}
		m.State.BoundProgram = ω.In.Program
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glUseProgram expected %v got %v", ω.Out, Σ)
		}
	case *GlIsProgram:
		Σ := GlIsProgram_Out{}
		Σ.Result = m.State.Instances.Programs.Contains(ω.In.Program)
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glIsProgram expected %v got %v", ω.Out, Σ)
		}
	case *GlValidateProgram:
		Σ := GlValidateProgram_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glValidateProgram expected %v got %v", ω.Out, Σ)
		}
	case *GlClearColor:
		Σ := GlClearColor_Out{}
		m.State.Clearing.ClearColor = func() Color {
			s := Color{}
			s.Init()
			s.Red = ω.In.R
			s.Green = ω.In.G
			s.Blue = ω.In.B
			s.Alpha = ω.In.A
			return s
		}()
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glClearColor expected %v got %v", ω.Out, Σ)
		}
	case *GlClearDepthf:
		Σ := GlClearDepthf_Out{}
		m.State.Clearing.ClearDepth = ω.In.Depth
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glClearDepthf expected %v got %v", ω.Out, Σ)
		}
	case *GlClearStencil:
		Σ := GlClearStencil_Out{}
		m.State.Clearing.ClearStencil = ω.In.Stencil
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glClearStencil expected %v got %v", ω.Out, Σ)
		}
	case *GlClear:
		Σ := GlClear_Out{}
		if (ClearMask_GL_COLOR_BUFFER_BIT)&(ω.In.Mask) != 0 {
		}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glClear expected %v got %v", ω.Out, Σ)
		}
	case *GlCullFace:
		Σ := GlCullFace_Out{}
		m.State.Rasterizing.CullFace = ω.In.Mode
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glCullFace expected %v got %v", ω.Out, Σ)
		}
	case *GlPolygonOffset:
		Σ := GlPolygonOffset_Out{}
		m.State.Rasterizing.PolygonOffsetUnits = ω.In.Units
		m.State.Rasterizing.PolygonOffsetFactor = ω.In.ScaleFactor
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glPolygonOffset expected %v got %v", ω.Out, Σ)
		}
	case *GlLineWidth:
		Σ := GlLineWidth_Out{}
		m.State.Rasterizing.LineWidth = ω.In.Width
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glLineWidth expected %v got %v", ω.Out, Σ)
		}
	case *GlSampleCoverage:
		Σ := GlSampleCoverage_Out{}
		m.State.Rasterizing.SampleCoverageValue = ω.In.Value
		m.State.Rasterizing.SampleCoverageInvert = ω.In.Invert
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glSampleCoverage expected %v got %v", ω.Out, Σ)
		}
	case *GlHint:
		Σ := GlHint_Out{}
		m.State.GenerateMipmapHint = ω.In.Mode
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glHint expected %v got %v", ω.Out, Σ)
		}
	case *GlFramebufferRenderbuffer:
		Σ := GlFramebufferRenderbuffer_Out{}
		framebufferId := m.State.BoundFramebuffers.Get(ω.In.FramebufferTarget) // FramebufferId
		framebuffer := m.State.Instances.Framebuffers.Get(framebufferId)       // FramebufferRef
		attachment := framebuffer.Attachments.Get(ω.In.FramebufferAttachment)  // FramebufferAttachmentInfo
		if (ω.In.Renderbuffer) == (m.State.Internals.NilRenderbuffer) {
			attachment.Type = FramebufferAttachmentType_GL_NONE
		} else {
			attachment.Type = FramebufferAttachmentType_GL_RENDERBUFFER
		}
		attachment.Object = uint32(ω.In.Renderbuffer)
		attachment.TextureLevel = 0
		attachment.CubeMapFace = CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		framebuffer.Attachments[ω.In.FramebufferAttachment] = attachment
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glFramebufferRenderbuffer expected %v got %v", ω.Out, Σ)
		}
	case *GlFramebufferTexture2D:
		Σ := GlFramebufferTexture2D_Out{}
		framebufferId := m.State.BoundFramebuffers.Get(ω.In.FramebufferTarget) // FramebufferId
		framebuffer := m.State.Instances.Framebuffers.Get(framebufferId)       // FramebufferRef
		attachment := framebuffer.Attachments.Get(ω.In.FramebufferAttachment)  // FramebufferAttachmentInfo
		if (ω.In.Texture) == (m.State.Internals.NilTexture) {
			attachment.Type = FramebufferAttachmentType_GL_NONE
			attachment.Object = uint32(m.State.Internals.NilTexture)
			attachment.TextureLevel = 0
			attachment.CubeMapFace = CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		}
		if (ω.In.Texture) != (m.State.Internals.NilTexture) {
			attachment.Type = FramebufferAttachmentType_GL_TEXTURE
			attachment.Object = uint32(ω.In.Texture)
			attachment.TextureLevel = ω.In.Level
			attachment.CubeMapFace = func() (result CubeMapImageTarget) {
				switch ω.In.TextureTarget {
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
		framebuffer.Attachments[ω.In.FramebufferAttachment] = attachment
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glFramebufferTexture2D expected %v got %v", ω.Out, Σ)
		}
	case *GlGetFramebufferAttachmentParameteriv:
		Σ := GlGetFramebufferAttachmentParameteriv_Out{}
		Σ.Value = make(S32Array, 1)
		framebufferId := m.State.BoundFramebuffers.Get(ω.In.Target)      // FramebufferId
		framebuffer := m.State.Instances.Framebuffers.Get(framebufferId) // FramebufferRef
		a := framebuffer.Attachments.Get(ω.In.Attachment)                // FramebufferAttachmentInfo
		Σ.Value[0] = func() (result int32) {
			switch ω.In.Parameter {
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
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGetFramebufferAttachmentParameteriv expected %v got %v", ω.Out, Σ)
		}
	case *GlDrawElements:
		Σ := GlDrawElements_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glDrawElements expected %v got %v", ω.Out, Σ)
		}
	case *GlDrawArrays:
		Σ := GlDrawArrays_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glDrawArrays expected %v got %v", ω.Out, Σ)
		}
	case *GlFlush:
		Σ := GlFlush_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glFlush expected %v got %v", ω.Out, Σ)
		}
	case *GlFinish:
		Σ := GlFinish_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glFinish expected %v got %v", ω.Out, Σ)
		}
	case *GlGetBooleanv:
		Σ := GlGetBooleanv_Out{}
		Σ.Values = make(BoolArray, stateVariableSize(ω.In.Param))
		switch ω.In.Param {
		case StateVariable_GL_BLEND:
			Σ.Values[0] = m.State.Capabilities.Get(Capability_GL_BLEND)
		case StateVariable_GL_CULL_FACE:
			Σ.Values[0] = m.State.Capabilities.Get(Capability_GL_CULL_FACE)
		case StateVariable_GL_DEPTH_TEST:
			Σ.Values[0] = m.State.Capabilities.Get(Capability_GL_DEPTH_TEST)
		case StateVariable_GL_DITHER:
			Σ.Values[0] = m.State.Capabilities.Get(Capability_GL_DITHER)
		case StateVariable_GL_POLYGON_OFFSET_FILL:
			Σ.Values[0] = m.State.Capabilities.Get(Capability_GL_POLYGON_OFFSET_FILL)
		case StateVariable_GL_SAMPLE_ALPHA_TO_COVERAGE:
			Σ.Values[0] = m.State.Capabilities.Get(Capability_GL_SAMPLE_ALPHA_TO_COVERAGE)
		case StateVariable_GL_SAMPLE_COVERAGE:
			Σ.Values[0] = m.State.Capabilities.Get(Capability_GL_SAMPLE_COVERAGE)
		case StateVariable_GL_SCISSOR_TEST:
			Σ.Values[0] = m.State.Capabilities.Get(Capability_GL_SCISSOR_TEST)
		case StateVariable_GL_STENCIL_TEST:
			Σ.Values[0] = m.State.Capabilities.Get(Capability_GL_STENCIL_TEST)
		case StateVariable_GL_DEPTH_WRITEMASK:
			Σ.Values[0] = m.State.Rasterizing.DepthMask
		case StateVariable_GL_COLOR_WRITEMASK:
			Σ.Values[0] = m.State.Rasterizing.ColorMaskRed
			Σ.Values[1] = m.State.Rasterizing.ColorMaskGreen
			Σ.Values[2] = m.State.Rasterizing.ColorMaskBlue
			Σ.Values[3] = m.State.Rasterizing.ColorMaskAlpha
		case StateVariable_GL_SAMPLE_COVERAGE_INVERT:
			Σ.Values[0] = m.State.Rasterizing.SampleCoverageInvert
		case StateVariable_GL_SHADER_COMPILER:
			Σ.Values[0] = ω.Out.Values[0]
		case StateVariable_GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT:
			Σ.Values[0] = ω.Out.Values[0]
		default:
			// TODO: better unmatched handling
			v := ω.In.Param
			log.Printf("Error: Missing switch case handler for value %T %v", v, v)
		}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGetBooleanv expected %v got %v", ω.Out, Σ)
		}
	case *GlGetFloatv:
		Σ := GlGetFloatv_Out{}
		Σ.Values = make(F32Array, stateVariableSize(ω.In.Param))
		switch ω.In.Param {
		case StateVariable_GL_DEPTH_RANGE:
			Σ.Values[0] = m.State.Rasterizing.DepthNear
			Σ.Values[1] = m.State.Rasterizing.DepthFar
		case StateVariable_GL_LINE_WIDTH:
			Σ.Values[0] = m.State.Rasterizing.LineWidth
		case StateVariable_GL_POLYGON_OFFSET_FACTOR:
			Σ.Values[0] = m.State.Rasterizing.PolygonOffsetFactor
		case StateVariable_GL_POLYGON_OFFSET_UNITS:
			Σ.Values[0] = m.State.Rasterizing.PolygonOffsetUnits
		case StateVariable_GL_SAMPLE_COVERAGE_VALUE:
			Σ.Values[0] = m.State.Rasterizing.SampleCoverageValue
		case StateVariable_GL_COLOR_CLEAR_VALUE:
			Σ.Values[0] = m.State.Clearing.ClearColor.Red
			Σ.Values[1] = m.State.Clearing.ClearColor.Green
			Σ.Values[2] = m.State.Clearing.ClearColor.Blue
			Σ.Values[3] = m.State.Clearing.ClearColor.Alpha
		case StateVariable_GL_DEPTH_CLEAR_VALUE:
			Σ.Values[0] = m.State.Clearing.ClearDepth
		case StateVariable_GL_ALIASED_LINE_WIDTH_RANGE:
			Σ.Values[0] = ω.Out.Values[0]
			Σ.Values[1] = ω.Out.Values[1]
		case StateVariable_GL_ALIASED_POINT_SIZE_RANGE:
			Σ.Values[0] = ω.Out.Values[0]
			Σ.Values[1] = ω.Out.Values[1]
		case StateVariable_GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT:
			Σ.Values[0] = ω.Out.Values[0]
		default:
			// TODO: better unmatched handling
			v := ω.In.Param
			log.Printf("Error: Missing switch case handler for value %T %v", v, v)
		}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGetFloatv expected %v got %v", ω.Out, Σ)
		}
	case *GlGetIntegerv:
		Σ := GlGetIntegerv_Out{}
		Σ.Values = make(S32Array, stateVariableSize(ω.In.Param))
		switch ω.In.Param {
		case StateVariable_GL_ACTIVE_TEXTURE:
			Σ.Values[0] = int32(m.State.ActiveTextureUnit)
		case StateVariable_GL_ARRAY_BUFFER_BINDING:
			Σ.Values[0] = int32(m.State.BoundBuffers.Get(BufferTarget_GL_ARRAY_BUFFER))
		case StateVariable_GL_ELEMENT_ARRAY_BUFFER_BINDING:
			Σ.Values[0] = int32(m.State.BoundBuffers.Get(BufferTarget_GL_ELEMENT_ARRAY_BUFFER))
		case StateVariable_GL_BLEND_SRC_ALPHA:
			Σ.Values[0] = int32(m.State.Blending.SrcAlphaBlendFactor)
		case StateVariable_GL_BLEND_SRC_RGB:
			Σ.Values[0] = int32(m.State.Blending.SrcRgbBlendFactor)
		case StateVariable_GL_BLEND_DST_ALPHA:
			Σ.Values[0] = int32(m.State.Blending.DstAlphaBlendFactor)
		case StateVariable_GL_BLEND_DST_RGB:
			Σ.Values[0] = int32(m.State.Blending.DstRgbBlendFactor)
		case StateVariable_GL_BLEND_EQUATION_RGB:
			Σ.Values[0] = int32(m.State.Blending.BlendEquationRgb)
		case StateVariable_GL_BLEND_EQUATION_ALPHA:
			Σ.Values[0] = int32(m.State.Blending.BlendEquationAlpha)
		case StateVariable_GL_BLEND_COLOR:
			Σ.Values[0] = int32(m.State.Blending.BlendColor.Red)
			Σ.Values[1] = int32(m.State.Blending.BlendColor.Green)
			Σ.Values[2] = int32(m.State.Blending.BlendColor.Blue)
			Σ.Values[3] = int32(m.State.Blending.BlendColor.Alpha)
		case StateVariable_GL_DEPTH_FUNC:
			Σ.Values[0] = int32(m.State.Rasterizing.DepthTestFunction)
		case StateVariable_GL_DEPTH_CLEAR_VALUE:
			Σ.Values[0] = int32(m.State.Clearing.ClearDepth)
		case StateVariable_GL_STENCIL_WRITEMASK:
			Σ.Values[0] = int32(m.State.Rasterizing.StencilMask.Get(FaceMode_GL_FRONT))
		case StateVariable_GL_STENCIL_BACK_WRITEMASK:
			Σ.Values[0] = int32(m.State.Rasterizing.StencilMask.Get(FaceMode_GL_BACK))
		case StateVariable_GL_VIEWPORT:
			Σ.Values[0] = m.State.Rasterizing.Viewport.X
			Σ.Values[1] = m.State.Rasterizing.Viewport.Y
			Σ.Values[2] = m.State.Rasterizing.Viewport.Width
			Σ.Values[3] = m.State.Rasterizing.Viewport.Height
		case StateVariable_GL_SCISSOR_BOX:
			Σ.Values[0] = m.State.Rasterizing.Scissor.X
			Σ.Values[1] = m.State.Rasterizing.Scissor.Y
			Σ.Values[2] = m.State.Rasterizing.Scissor.Width
			Σ.Values[3] = m.State.Rasterizing.Scissor.Height
		case StateVariable_GL_FRONT_FACE:
			Σ.Values[0] = int32(m.State.Rasterizing.FrontFace)
		case StateVariable_GL_CULL_FACE_MODE:
			Σ.Values[0] = int32(m.State.Rasterizing.CullFace)
		case StateVariable_GL_STENCIL_CLEAR_VALUE:
			Σ.Values[0] = m.State.Clearing.ClearStencil
		case StateVariable_GL_FRAMEBUFFER_BINDING:
			Σ.Values[0] = int32(m.State.BoundFramebuffers.Get(FramebufferTarget_GL_FRAMEBUFFER))
		case StateVariable_GL_READ_FRAMEBUFFER_BINDING:
			Σ.Values[0] = int32(m.State.BoundFramebuffers.Get(FramebufferTarget_GL_READ_FRAMEBUFFER))
		case StateVariable_GL_RENDERBUFFER_BINDING:
			Σ.Values[0] = int32(m.State.BoundRenderbuffers.Get(RenderbufferTarget_GL_RENDERBUFFER))
		case StateVariable_GL_CURRENT_PROGRAM:
			Σ.Values[0] = int32(m.State.BoundProgram)
		case StateVariable_GL_TEXTURE_BINDING_2D:
			Σ.Values[0] = int32(m.State.TextureUnits.Get(m.State.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_2D))
		case StateVariable_GL_TEXTURE_BINDING_CUBE_MAP:
			Σ.Values[0] = int32(m.State.TextureUnits.Get(m.State.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_CUBE_MAP))
		case StateVariable_GL_GENERATE_MIPMAP_HINT:
			Σ.Values[0] = int32(m.State.GenerateMipmapHint)
		case StateVariable_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS:
			Σ.Values[0] = ω.Out.Values[0]
		case StateVariable_GL_MAX_CUBE_MAP_TEXTURE_SIZE:
			Σ.Values[0] = ω.Out.Values[0]
		case StateVariable_GL_MAX_FRAGMENT_UNIFORM_VECTORS:
			Σ.Values[0] = ω.Out.Values[0]
		case StateVariable_GL_MAX_RENDERBUFFER_SIZE:
			Σ.Values[0] = ω.Out.Values[0]
		case StateVariable_GL_MAX_TEXTURE_IMAGE_UNITS:
			Σ.Values[0] = ω.Out.Values[0]
		case StateVariable_GL_MAX_TEXTURE_SIZE:
			Σ.Values[0] = ω.Out.Values[0]
		case StateVariable_GL_MAX_VARYING_VECTORS:
			Σ.Values[0] = ω.Out.Values[0]
		case StateVariable_GL_MAX_VERTEX_ATTRIBS:
			Σ.Values[0] = ω.Out.Values[0]
		case StateVariable_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS:
			Σ.Values[0] = ω.Out.Values[0]
		case StateVariable_GL_MAX_VERTEX_UNIFORM_VECTORS:
			Σ.Values[0] = ω.Out.Values[0]
		case StateVariable_GL_MAX_VIEWPORT_DIMS:
			Σ.Values[0] = ω.Out.Values[0]
			Σ.Values[1] = ω.Out.Values[1]
		case StateVariable_GL_NUM_COMPRESSED_TEXTURE_FORMATS:
			Σ.Values[0] = ω.Out.Values[0]
		case StateVariable_GL_NUM_SHADER_BINARY_FORMATS:
			Σ.Values[0] = ω.Out.Values[0]
		case StateVariable_GL_PACK_ALIGNMENT:
			Σ.Values[0] = m.State.PixelStorage.Get(PixelStoreParameter_GL_PACK_ALIGNMENT)
		case StateVariable_GL_UNPACK_ALIGNMENT:
			Σ.Values[0] = m.State.PixelStorage.Get(PixelStoreParameter_GL_UNPACK_ALIGNMENT)
		case StateVariable_GL_ALPHA_BITS:
			Σ.Values[0] = ω.Out.Values[0]
		case StateVariable_GL_BLUE_BITS:
			Σ.Values[0] = ω.Out.Values[0]
		case StateVariable_GL_GREEN_BITS:
			Σ.Values[0] = ω.Out.Values[0]
		case StateVariable_GL_RED_BITS:
			Σ.Values[0] = ω.Out.Values[0]
		case StateVariable_GL_DEPTH_BITS:
			Σ.Values[0] = ω.Out.Values[0]
		case StateVariable_GL_SAMPLE_BUFFERS:
			Σ.Values[0] = ω.Out.Values[0]
		case StateVariable_GL_SAMPLES:
			Σ.Values[0] = ω.Out.Values[0]
		case StateVariable_GL_SHADER_BINARY_FORMATS:
			Σ.Values = ω.Out.Values
		case StateVariable_GL_COMPRESSED_TEXTURE_FORMATS:
			Σ.Values = ω.Out.Values
		case StateVariable_GL_STENCIL_BITS:
			Σ.Values[0] = ω.Out.Values[0]
		case StateVariable_GL_SUBPIXEL_BITS:
			Σ.Values[0] = ω.Out.Values[0]
		case StateVariable_GL_IMPLEMENTATION_COLOR_READ_FORMAT:
			Σ.Values[0] = ω.Out.Values[0]
		case StateVariable_GL_IMPLEMENTATION_COLOR_READ_TYPE:
			Σ.Values[0] = ω.Out.Values[0]
		case StateVariable_GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT:
			Σ.Values[0] = ω.Out.Values[0]
		case StateVariable_GL_GPU_DISJOINT_EXT:
			Σ.Values[0] = ω.Out.Values[0]
		default:
			// TODO: better unmatched handling
			v := ω.In.Param
			log.Printf("Error: Missing switch case handler for value %T %v", v, v)
		}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGetIntegerv expected %v got %v", ω.Out, Σ)
		}
	case *GlGetString:
		Σ := GlGetString_Out{}
		Σ.Result = ω.Out.Result
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGetString expected %v got %v", ω.Out, Σ)
		}
	case *GlEnable:
		Σ := GlEnable_Out{}
		m.State.Capabilities[ω.In.Capability] = true
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glEnable expected %v got %v", ω.Out, Σ)
		}
	case *GlDisable:
		Σ := GlDisable_Out{}
		m.State.Capabilities[ω.In.Capability] = false
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glDisable expected %v got %v", ω.Out, Σ)
		}
	case *GlIsEnabled:
		Σ := GlIsEnabled_Out{}
		Σ.Result = m.State.Capabilities.Get(ω.In.Capability)
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glIsEnabled expected %v got %v", ω.Out, Σ)
		}
	case *GlMapBufferRange:
		Σ := GlMapBufferRange_Out{}
		Σ.Result = ω.Out.Result
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glMapBufferRange expected %v got %v", ω.Out, Σ)
		}
	case *GlUnmapBuffer:
		Σ := GlUnmapBuffer_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glUnmapBuffer expected %v got %v", ω.Out, Σ)
		}
	case *GlInvalidateFramebuffer:
		Σ := GlInvalidateFramebuffer_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glInvalidateFramebuffer expected %v got %v", ω.Out, Σ)
		}
	case *GlRenderbufferStorageMultisample:
		Σ := GlRenderbufferStorageMultisample_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glRenderbufferStorageMultisample expected %v got %v", ω.Out, Σ)
		}
	case *GlBlitFramebuffer:
		Σ := GlBlitFramebuffer_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glBlitFramebuffer expected %v got %v", ω.Out, Σ)
		}
	case *GlGenQueries:
		Σ := GlGenQueries_Out{}
		Σ.Queries = make(QueryIdArray, ω.In.Count)
		for i := int32(0); i < ω.In.Count; i++ {
			id := QueryId(ω.Out.Queries[i]) // QueryId
			m.State.Instances.Queries[id] = func() *Query {
				s := &Query{}
				s.Init()
				return s
			}()
			Σ.Queries[i] = id
		}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGenQueries expected %v got %v", ω.Out, Σ)
		}
	case *GlBeginQuery:
		Σ := GlBeginQuery_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glBeginQuery expected %v got %v", ω.Out, Σ)
		}
	case *GlEndQuery:
		Σ := GlEndQuery_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glEndQuery expected %v got %v", ω.Out, Σ)
		}
	case *GlDeleteQueries:
		Σ := GlDeleteQueries_Out{}
		for i := int32(0); i < ω.In.Count; i++ {
			m.State.Instances.Queries.Delete(ω.In.Queries[i])
		}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glDeleteQueries expected %v got %v", ω.Out, Σ)
		}
	case *GlIsQuery:
		Σ := GlIsQuery_Out{}
		Σ.Result = m.State.Instances.Queries.Contains(ω.In.Query)
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glIsQuery expected %v got %v", ω.Out, Σ)
		}
	case *GlGetQueryiv:
		Σ := GlGetQueryiv_Out{}
		Σ.Value = ω.Out.Value
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGetQueryiv expected %v got %v", ω.Out, Σ)
		}
	case *GlGetQueryObjectuiv:
		Σ := GlGetQueryObjectuiv_Out{}
		Σ.Value = ω.Out.Value
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGetQueryObjectuiv expected %v got %v", ω.Out, Σ)
		}
	case *GlGenQueriesEXT:
		Σ := GlGenQueriesEXT_Out{}
		Σ.Queries = make(QueryIdArray, ω.In.Count)
		for i := int32(0); i < ω.In.Count; i++ {
			id := QueryId(ω.Out.Queries[i]) // QueryId
			m.State.Instances.Queries[id] = func() *Query {
				s := &Query{}
				s.Init()
				return s
			}()
			Σ.Queries[i] = id
		}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGenQueriesEXT expected %v got %v", ω.Out, Σ)
		}
	case *GlBeginQueryEXT:
		Σ := GlBeginQueryEXT_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glBeginQueryEXT expected %v got %v", ω.Out, Σ)
		}
	case *GlEndQueryEXT:
		Σ := GlEndQueryEXT_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glEndQueryEXT expected %v got %v", ω.Out, Σ)
		}
	case *GlDeleteQueriesEXT:
		Σ := GlDeleteQueriesEXT_Out{}
		for i := int32(0); i < ω.In.Count; i++ {
			m.State.Instances.Queries.Delete(ω.In.Queries[i])
		}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glDeleteQueriesEXT expected %v got %v", ω.Out, Σ)
		}
	case *GlIsQueryEXT:
		Σ := GlIsQueryEXT_Out{}
		Σ.Result = m.State.Instances.Queries.Contains(ω.In.Query)
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glIsQueryEXT expected %v got %v", ω.Out, Σ)
		}
	case *GlQueryCounterEXT:
		Σ := GlQueryCounterEXT_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glQueryCounterEXT expected %v got %v", ω.Out, Σ)
		}
	case *GlGetQueryivEXT:
		Σ := GlGetQueryivEXT_Out{}
		Σ.Value = ω.Out.Value
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGetQueryivEXT expected %v got %v", ω.Out, Σ)
		}
	case *GlGetQueryObjectivEXT:
		Σ := GlGetQueryObjectivEXT_Out{}
		Σ.Value = ω.Out.Value
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGetQueryObjectivEXT expected %v got %v", ω.Out, Σ)
		}
	case *GlGetQueryObjectuivEXT:
		Σ := GlGetQueryObjectuivEXT_Out{}
		Σ.Value = ω.Out.Value
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGetQueryObjectuivEXT expected %v got %v", ω.Out, Σ)
		}
	case *GlGetQueryObjecti64vEXT:
		Σ := GlGetQueryObjecti64vEXT_Out{}
		Σ.Value = ω.Out.Value
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGetQueryObjecti64vEXT expected %v got %v", ω.Out, Σ)
		}
	case *GlGetQueryObjectui64vEXT:
		Σ := GlGetQueryObjectui64vEXT_Out{}
		Σ.Value = ω.Out.Value
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying glGetQueryObjectui64vEXT expected %v got %v", ω.Out, Σ)
		}
	case *memory.Observation:
		m.State.Mem.Slice(ω.Range).Write(memory.ResourceData(ω.ResourceID, ω.Range.Size))
	}
}
