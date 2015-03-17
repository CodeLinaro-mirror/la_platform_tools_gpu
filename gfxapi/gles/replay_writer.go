////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

package gles

import (
	"fmt"
	"io"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/replay/builder"
	"android.googlesource.com/platform/tools/gpu/replay/protocol"
	"android.googlesource.com/platform/tools/gpu/replay/value"
)

type replayer interface {
	replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool)
}

func readBytes(r io.Reader, c uint64) ([]byte, error) {
	b := make([]byte, c)
	_, err := io.ReadFull(r, b)
	return b, err
}
func readString(r io.Reader, c uint64) (string, error) {
	if buf, err := readBytes(r, c); err == nil {
		str := string(buf)
		for i, c := range str {
			if c == 0 {
				return str[:i], nil
			}
		}
		return str, nil
	} else {
		return "", err
	}
}

var funcInfoInit = builder.FunctionInfo{ID: 0, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoStartTimer = builder.FunctionInfo{ID: 1, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoStopTimer = builder.FunctionInfo{ID: 2, ReturnType: protocol.TypeUint64, Parameters: 1}
var funcInfoFlushPostBuffer = builder.FunctionInfo{ID: 3, ReturnType: protocol.TypeVoid, Parameters: 0}
var funcInfoEglCreateContext = builder.FunctionInfo{ID: 4, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoEglMakeCurrent = builder.FunctionInfo{ID: 5, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoEglSwapBuffers = builder.FunctionInfo{ID: 6, ReturnType: protocol.TypeVoid, Parameters: 0}
var funcInfoGlEnableClientState = builder.FunctionInfo{ID: 7, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlDisableClientState = builder.FunctionInfo{ID: 8, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlGetProgramBinaryOES = builder.FunctionInfo{ID: 9, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlProgramBinaryOES = builder.FunctionInfo{ID: 10, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlStartTilingQCOM = builder.FunctionInfo{ID: 11, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlEndTilingQCOM = builder.FunctionInfo{ID: 12, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlDiscardFramebufferEXT = builder.FunctionInfo{ID: 13, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlInsertEventMarkerEXT = builder.FunctionInfo{ID: 14, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlPushGroupMarkerEXT = builder.FunctionInfo{ID: 15, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlPopGroupMarkerEXT = builder.FunctionInfo{ID: 16, ReturnType: protocol.TypeVoid, Parameters: 0}
var funcInfoGlTexStorage1DEXT = builder.FunctionInfo{ID: 17, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlTexStorage2DEXT = builder.FunctionInfo{ID: 18, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlTexStorage3DEXT = builder.FunctionInfo{ID: 19, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlTextureStorage1DEXT = builder.FunctionInfo{ID: 20, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlTextureStorage2DEXT = builder.FunctionInfo{ID: 21, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlTextureStorage3DEXT = builder.FunctionInfo{ID: 22, ReturnType: protocol.TypeVoid, Parameters: 7}
var funcInfoGlGenVertexArraysOES = builder.FunctionInfo{ID: 23, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBindVertexArrayOES = builder.FunctionInfo{ID: 24, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlDeleteVertexArraysOES = builder.FunctionInfo{ID: 25, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlIsVertexArrayOES = builder.FunctionInfo{ID: 26, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlEGLImageTargetTexture2DOES = builder.FunctionInfo{ID: 27, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlEGLImageTargetRenderbufferStorageOES = builder.FunctionInfo{ID: 28, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetGraphicsResetStatusEXT = builder.FunctionInfo{ID: 29, ReturnType: protocol.TypeUint32, Parameters: 0}
var funcInfoGlBindAttribLocation = builder.FunctionInfo{ID: 30, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlBlendFunc = builder.FunctionInfo{ID: 31, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBlendFuncSeparate = builder.FunctionInfo{ID: 32, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlBlendEquation = builder.FunctionInfo{ID: 33, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlBlendEquationSeparate = builder.FunctionInfo{ID: 34, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBlendColor = builder.FunctionInfo{ID: 35, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlEnableVertexAttribArray = builder.FunctionInfo{ID: 36, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlDisableVertexAttribArray = builder.FunctionInfo{ID: 37, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlVertexAttribPointer = builder.FunctionInfo{ID: 38, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlGetActiveAttrib = builder.FunctionInfo{ID: 39, ReturnType: protocol.TypeVoid, Parameters: 7}
var funcInfoGlGetActiveUniform = builder.FunctionInfo{ID: 40, ReturnType: protocol.TypeVoid, Parameters: 7}
var funcInfoGlGetError = builder.FunctionInfo{ID: 41, ReturnType: protocol.TypeUint32, Parameters: 0}
var funcInfoGlGetProgramiv = builder.FunctionInfo{ID: 42, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetShaderiv = builder.FunctionInfo{ID: 43, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetUniformLocation = builder.FunctionInfo{ID: 44, ReturnType: protocol.TypeInt32, Parameters: 2}
var funcInfoGlGetAttribLocation = builder.FunctionInfo{ID: 45, ReturnType: protocol.TypeInt32, Parameters: 2}
var funcInfoGlPixelStorei = builder.FunctionInfo{ID: 46, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlTexParameteri = builder.FunctionInfo{ID: 47, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlTexParameterf = builder.FunctionInfo{ID: 48, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetTexParameteriv = builder.FunctionInfo{ID: 49, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetTexParameterfv = builder.FunctionInfo{ID: 50, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform1i = builder.FunctionInfo{ID: 51, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlUniform2i = builder.FunctionInfo{ID: 52, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform3i = builder.FunctionInfo{ID: 53, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUniform4i = builder.FunctionInfo{ID: 54, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlUniform1iv = builder.FunctionInfo{ID: 55, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform2iv = builder.FunctionInfo{ID: 56, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform3iv = builder.FunctionInfo{ID: 57, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform4iv = builder.FunctionInfo{ID: 58, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform1f = builder.FunctionInfo{ID: 59, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlUniform2f = builder.FunctionInfo{ID: 60, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform3f = builder.FunctionInfo{ID: 61, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUniform4f = builder.FunctionInfo{ID: 62, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlUniform1fv = builder.FunctionInfo{ID: 63, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform2fv = builder.FunctionInfo{ID: 64, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform3fv = builder.FunctionInfo{ID: 65, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform4fv = builder.FunctionInfo{ID: 66, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniformMatrix2fv = builder.FunctionInfo{ID: 67, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUniformMatrix3fv = builder.FunctionInfo{ID: 68, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUniformMatrix4fv = builder.FunctionInfo{ID: 69, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlGetUniformfv = builder.FunctionInfo{ID: 70, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetUniformiv = builder.FunctionInfo{ID: 71, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlVertexAttrib1f = builder.FunctionInfo{ID: 72, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlVertexAttrib2f = builder.FunctionInfo{ID: 73, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlVertexAttrib3f = builder.FunctionInfo{ID: 74, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlVertexAttrib4f = builder.FunctionInfo{ID: 75, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlVertexAttrib1fv = builder.FunctionInfo{ID: 76, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlVertexAttrib2fv = builder.FunctionInfo{ID: 77, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlVertexAttrib3fv = builder.FunctionInfo{ID: 78, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlVertexAttrib4fv = builder.FunctionInfo{ID: 79, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetShaderPrecisionFormat = builder.FunctionInfo{ID: 80, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlDepthMask = builder.FunctionInfo{ID: 81, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlDepthFunc = builder.FunctionInfo{ID: 82, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlDepthRangef = builder.FunctionInfo{ID: 83, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlColorMask = builder.FunctionInfo{ID: 84, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlStencilMask = builder.FunctionInfo{ID: 85, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlStencilMaskSeparate = builder.FunctionInfo{ID: 86, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlStencilFuncSeparate = builder.FunctionInfo{ID: 87, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlStencilOpSeparate = builder.FunctionInfo{ID: 88, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlFrontFace = builder.FunctionInfo{ID: 89, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlViewport = builder.FunctionInfo{ID: 90, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlScissor = builder.FunctionInfo{ID: 91, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlActiveTexture = builder.FunctionInfo{ID: 92, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlGenTextures = builder.FunctionInfo{ID: 93, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlDeleteTextures = builder.FunctionInfo{ID: 94, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlIsTexture = builder.FunctionInfo{ID: 95, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlBindTexture = builder.FunctionInfo{ID: 96, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlTexImage2D = builder.FunctionInfo{ID: 97, ReturnType: protocol.TypeVoid, Parameters: 9}
var funcInfoGlTexSubImage2D = builder.FunctionInfo{ID: 98, ReturnType: protocol.TypeVoid, Parameters: 9}
var funcInfoGlCopyTexImage2D = builder.FunctionInfo{ID: 99, ReturnType: protocol.TypeVoid, Parameters: 8}
var funcInfoGlCopyTexSubImage2D = builder.FunctionInfo{ID: 100, ReturnType: protocol.TypeVoid, Parameters: 8}
var funcInfoGlCompressedTexImage2D = builder.FunctionInfo{ID: 101, ReturnType: protocol.TypeVoid, Parameters: 8}
var funcInfoGlCompressedTexSubImage2D = builder.FunctionInfo{ID: 102, ReturnType: protocol.TypeVoid, Parameters: 9}
var funcInfoGlGenerateMipmap = builder.FunctionInfo{ID: 103, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlReadPixels = builder.FunctionInfo{ID: 104, ReturnType: protocol.TypeVoid, Parameters: 7}
var funcInfoGlGenFramebuffers = builder.FunctionInfo{ID: 105, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBindFramebuffer = builder.FunctionInfo{ID: 106, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlCheckFramebufferStatus = builder.FunctionInfo{ID: 107, ReturnType: protocol.TypeUint32, Parameters: 1}
var funcInfoGlDeleteFramebuffers = builder.FunctionInfo{ID: 108, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlIsFramebuffer = builder.FunctionInfo{ID: 109, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlGenRenderbuffers = builder.FunctionInfo{ID: 110, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBindRenderbuffer = builder.FunctionInfo{ID: 111, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlRenderbufferStorage = builder.FunctionInfo{ID: 112, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlDeleteRenderbuffers = builder.FunctionInfo{ID: 113, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlIsRenderbuffer = builder.FunctionInfo{ID: 114, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlGetRenderbufferParameteriv = builder.FunctionInfo{ID: 115, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGenBuffers = builder.FunctionInfo{ID: 116, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBindBuffer = builder.FunctionInfo{ID: 117, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBufferData = builder.FunctionInfo{ID: 118, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlBufferSubData = builder.FunctionInfo{ID: 119, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlDeleteBuffers = builder.FunctionInfo{ID: 120, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlIsBuffer = builder.FunctionInfo{ID: 121, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlGetBufferParameteriv = builder.FunctionInfo{ID: 122, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlCreateShader = builder.FunctionInfo{ID: 123, ReturnType: protocol.TypeUint32, Parameters: 1}
var funcInfoGlDeleteShader = builder.FunctionInfo{ID: 124, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlShaderSource = builder.FunctionInfo{ID: 125, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlShaderBinary = builder.FunctionInfo{ID: 126, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlGetShaderInfoLog = builder.FunctionInfo{ID: 127, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlGetShaderSource = builder.FunctionInfo{ID: 128, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlReleaseShaderCompiler = builder.FunctionInfo{ID: 129, ReturnType: protocol.TypeVoid, Parameters: 0}
var funcInfoGlCompileShader = builder.FunctionInfo{ID: 130, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlIsShader = builder.FunctionInfo{ID: 131, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlCreateProgram = builder.FunctionInfo{ID: 132, ReturnType: protocol.TypeUint32, Parameters: 0}
var funcInfoGlDeleteProgram = builder.FunctionInfo{ID: 133, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlAttachShader = builder.FunctionInfo{ID: 134, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlDetachShader = builder.FunctionInfo{ID: 135, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetAttachedShaders = builder.FunctionInfo{ID: 136, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlLinkProgram = builder.FunctionInfo{ID: 137, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlGetProgramInfoLog = builder.FunctionInfo{ID: 138, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUseProgram = builder.FunctionInfo{ID: 139, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlIsProgram = builder.FunctionInfo{ID: 140, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlValidateProgram = builder.FunctionInfo{ID: 141, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlClearColor = builder.FunctionInfo{ID: 142, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlClearDepthf = builder.FunctionInfo{ID: 143, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlClearStencil = builder.FunctionInfo{ID: 144, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlClear = builder.FunctionInfo{ID: 145, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlCullFace = builder.FunctionInfo{ID: 146, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlPolygonOffset = builder.FunctionInfo{ID: 147, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlLineWidth = builder.FunctionInfo{ID: 148, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlSampleCoverage = builder.FunctionInfo{ID: 149, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlHint = builder.FunctionInfo{ID: 150, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlFramebufferRenderbuffer = builder.FunctionInfo{ID: 151, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlFramebufferTexture2D = builder.FunctionInfo{ID: 152, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlGetFramebufferAttachmentParameteriv = builder.FunctionInfo{ID: 153, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlDrawElements = builder.FunctionInfo{ID: 154, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlDrawArrays = builder.FunctionInfo{ID: 155, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlFlush = builder.FunctionInfo{ID: 156, ReturnType: protocol.TypeVoid, Parameters: 0}
var funcInfoGlFinish = builder.FunctionInfo{ID: 157, ReturnType: protocol.TypeVoid, Parameters: 0}
var funcInfoGlGetBooleanv = builder.FunctionInfo{ID: 158, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetFloatv = builder.FunctionInfo{ID: 159, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetIntegerv = builder.FunctionInfo{ID: 160, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetString = builder.FunctionInfo{ID: 161, ReturnType: protocol.TypeAbsolutePointer, Parameters: 1}
var funcInfoGlEnable = builder.FunctionInfo{ID: 162, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlDisable = builder.FunctionInfo{ID: 163, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlIsEnabled = builder.FunctionInfo{ID: 164, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlMapBufferRange = builder.FunctionInfo{ID: 165, ReturnType: protocol.TypeAbsolutePointer, Parameters: 4}
var funcInfoGlUnmapBuffer = builder.FunctionInfo{ID: 166, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlInvalidateFramebuffer = builder.FunctionInfo{ID: 167, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlRenderbufferStorageMultisample = builder.FunctionInfo{ID: 168, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlBlitFramebuffer = builder.FunctionInfo{ID: 169, ReturnType: protocol.TypeVoid, Parameters: 10}
var funcInfoGlGenQueries = builder.FunctionInfo{ID: 170, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBeginQuery = builder.FunctionInfo{ID: 171, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlEndQuery = builder.FunctionInfo{ID: 172, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlDeleteQueries = builder.FunctionInfo{ID: 173, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlIsQuery = builder.FunctionInfo{ID: 174, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlGetQueryiv = builder.FunctionInfo{ID: 175, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetQueryObjectuiv = builder.FunctionInfo{ID: 176, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGenQueriesEXT = builder.FunctionInfo{ID: 177, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBeginQueryEXT = builder.FunctionInfo{ID: 178, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlEndQueryEXT = builder.FunctionInfo{ID: 179, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlDeleteQueriesEXT = builder.FunctionInfo{ID: 180, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlIsQueryEXT = builder.FunctionInfo{ID: 181, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlQueryCounterEXT = builder.FunctionInfo{ID: 182, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetQueryivEXT = builder.FunctionInfo{ID: 183, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetQueryObjectivEXT = builder.FunctionInfo{ID: 184, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetQueryObjectuivEXT = builder.FunctionInfo{ID: 185, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetQueryObjecti64vEXT = builder.FunctionInfo{ID: 186, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetQueryObjectui64vEXT = builder.FunctionInfo{ID: 187, ReturnType: protocol.TypeVoid, Parameters: 3}

func (c RenderbufferId) value(b *builder.Builder, ω atom.Atom, s *state) value.Value {
	return value.U32(uint32(c))
}
func (c TextureId) value(b *builder.Builder, ω atom.Atom, s *state) value.Value {
	return value.U32(uint32(c))
}
func (c FramebufferId) value(b *builder.Builder, ω atom.Atom, s *state) value.Value {
	return value.U32(uint32(c))
}
func (c BufferId) value(b *builder.Builder, ω atom.Atom, s *state) value.Value {
	return value.U32(uint32(c))
}
func (c ShaderId) value(b *builder.Builder, ω atom.Atom, s *state) value.Value {
	return value.U32(uint32(c))
}
func (c ProgramId) value(b *builder.Builder, ω atom.Atom, s *state) value.Value {
	return value.U32(uint32(c))
}
func (c VertexArrayId) value(b *builder.Builder, ω atom.Atom, s *state) value.Value {
	return value.U32(uint32(c))
}
func (c QueryId) value(b *builder.Builder, ω atom.Atom, s *state) value.Value {
	return value.U32(uint32(c))
}
func (c UniformLocation) value(b *builder.Builder, ω atom.Atom, s *state) value.Value {
	return value.S32(int32(c))
}
func (c AttributeLocation) value(b *builder.Builder, ω atom.Atom, s *state) value.Value {
	return value.S32(int32(c))
}
func (arr BoolArray) value(b *builder.Builder, ω atom.Atom, s *state) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			b.Push(value.Bool(e))
		}
		return b.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr BufferIdArray) value(b *builder.Builder, ω atom.Atom, s *state) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			if key, remap := e.remap(ω, s); remap {
				loadRemap(b, key, e.value(b, ω, s))
			} else {
				b.Push(e.value(b, ω, s))
			}
		}
		return b.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr DiscardFramebufferAttachmentArray) value(b *builder.Builder, ω atom.Atom, s *state) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			b.Push(value.U32(e))
		}
		return b.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr F32Array) value(b *builder.Builder, ω atom.Atom, s *state) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			b.Push(value.F32(e))
		}
		return b.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr FramebufferAttachmentArray) value(b *builder.Builder, ω atom.Atom, s *state) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			b.Push(value.U32(e))
		}
		return b.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr FramebufferIdArray) value(b *builder.Builder, ω atom.Atom, s *state) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			if key, remap := e.remap(ω, s); remap {
				loadRemap(b, key, e.value(b, ω, s))
			} else {
				b.Push(e.value(b, ω, s))
			}
		}
		return b.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr QueryIdArray) value(b *builder.Builder, ω atom.Atom, s *state) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			if key, remap := e.remap(ω, s); remap {
				loadRemap(b, key, e.value(b, ω, s))
			} else {
				b.Push(e.value(b, ω, s))
			}
		}
		return b.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr RenderbufferIdArray) value(b *builder.Builder, ω atom.Atom, s *state) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			if key, remap := e.remap(ω, s); remap {
				loadRemap(b, key, e.value(b, ω, s))
			} else {
				b.Push(e.value(b, ω, s))
			}
		}
		return b.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr S32Array) value(b *builder.Builder, ω atom.Atom, s *state) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			b.Push(value.S32(e))
		}
		return b.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr ShaderIdArray) value(b *builder.Builder, ω atom.Atom, s *state) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			if key, remap := e.remap(ω, s); remap {
				loadRemap(b, key, e.value(b, ω, s))
			} else {
				b.Push(e.value(b, ω, s))
			}
		}
		return b.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr StringArray) value(b *builder.Builder, ω atom.Atom, s *state) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			b.Push(b.String(e))
		}
		return b.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr TextureIdArray) value(b *builder.Builder, ω atom.Atom, s *state) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			if key, remap := e.remap(ω, s); remap {
				loadRemap(b, key, e.value(b, ω, s))
			} else {
				b.Push(e.value(b, ω, s))
			}
		}
		return b.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr VertexArrayIdArray) value(b *builder.Builder, ω atom.Atom, s *state) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			if key, remap := e.remap(ω, s); remap {
				loadRemap(b, key, e.value(b, ω, s))
			} else {
				b.Push(e.value(b, ω, s))
			}
		}
		return b.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}

type StopTimer_Postback struct {
	Result uint64
}

func (o *StopTimer_Postback) Decode(d *protocol.Decoder) error {
	if v, err := d.Uint64(); err == nil {
		o.Result = v
	} else {
		return err
	}
	return nil
}

type EglCreateContext_Postback struct {
	Version int32
	Context int32
}

func (o *EglCreateContext_Postback) Decode(d *protocol.Decoder) error {
	if v, err := d.Int32(); err == nil {
		o.Version = v
	} else {
		return err
	}
	if v, err := d.Int32(); err == nil {
		o.Context = v
	} else {
		return err
	}
	return nil
}

type GlGetProgramBinaryOES_Postback struct {
	BytesWritten int32
	BinaryFormat uint32
	Binary       []byte
}

func (o *GlGetProgramBinaryOES_Postback) Decode(binary_cnt uint64, d *protocol.Decoder) error {
	if v, err := d.Int32(); err == nil {
		o.BytesWritten = v
	} else {
		return err
	}
	if v, err := d.Uint32(); err == nil {
		o.BinaryFormat = v
	} else {
		return err
	}
	if val, err := readBytes(d, binary_cnt); err == nil {
		o.Binary = val
	} else {
		return err
	}
	return nil
}

type GlGenVertexArraysOES_Postback struct {
	Arrays VertexArrayIdArray
}

func (o *GlGenVertexArraysOES_Postback) Decode(arrays_cnt uint64, d *protocol.Decoder) error {
	o.Arrays = make(VertexArrayIdArray, arrays_cnt)
	for i := range o.Arrays {
		{
			var x uint32
			if v, err := d.Uint32(); err == nil {
				x = v
			} else {
				return err
			}
			o.Arrays[i] = VertexArrayId(x)
		}
	}
	return nil
}

type GlIsVertexArrayOES_Postback struct {
	Result bool
}

func (o *GlIsVertexArrayOES_Postback) Decode(d *protocol.Decoder) error {
	if v, err := d.Bool(); err == nil {
		o.Result = v
	} else {
		return err
	}
	return nil
}

type GlGetGraphicsResetStatusEXT_Postback struct {
	Result ResetStatus
}

func (o *GlGetGraphicsResetStatusEXT_Postback) Decode(d *protocol.Decoder) error {
	if v, err := d.Uint32(); err == nil {
		o.Result = ResetStatus(v)
	} else {
		return err
	}
	return nil
}

type GlGetActiveAttrib_Postback struct {
	BufferBytesWritten int32
	VectorCount        int32
	Type               ShaderAttribType
	Name               string
}

func (o *GlGetActiveAttrib_Postback) Decode(name_cnt uint64, d *protocol.Decoder) error {
	if v, err := d.Int32(); err == nil {
		o.BufferBytesWritten = v
	} else {
		return err
	}
	if v, err := d.Int32(); err == nil {
		o.VectorCount = v
	} else {
		return err
	}
	if v, err := d.Uint32(); err == nil {
		o.Type = ShaderAttribType(v)
	} else {
		return err
	}
	if val, err := readString(d, name_cnt); err == nil {
		o.Name = val
	} else {
		return err
	}
	return nil
}

type GlGetActiveUniform_Postback struct {
	BufferBytesWritten int32
	Size               int32
	Type               ShaderUniformType
	Name               string
}

func (o *GlGetActiveUniform_Postback) Decode(name_cnt uint64, d *protocol.Decoder) error {
	if v, err := d.Int32(); err == nil {
		o.BufferBytesWritten = v
	} else {
		return err
	}
	if v, err := d.Int32(); err == nil {
		o.Size = v
	} else {
		return err
	}
	if v, err := d.Uint32(); err == nil {
		o.Type = ShaderUniformType(v)
	} else {
		return err
	}
	if val, err := readString(d, name_cnt); err == nil {
		o.Name = val
	} else {
		return err
	}
	return nil
}

type GlGetError_Postback struct {
	Result Error
}

func (o *GlGetError_Postback) Decode(d *protocol.Decoder) error {
	if v, err := d.Uint32(); err == nil {
		o.Result = Error(v)
	} else {
		return err
	}
	return nil
}

type GlGetProgramiv_Postback struct {
	Value S32Array
}

func (o *GlGetProgramiv_Postback) Decode(value_cnt uint64, d *protocol.Decoder) error {
	o.Value = make(S32Array, value_cnt)
	for i := range o.Value {
		if v, err := d.Int32(); err == nil {
			o.Value[i] = v
		} else {
			return err
		}
	}
	return nil
}

type GlGetShaderiv_Postback struct {
	Value S32Array
}

func (o *GlGetShaderiv_Postback) Decode(value_cnt uint64, d *protocol.Decoder) error {
	o.Value = make(S32Array, value_cnt)
	for i := range o.Value {
		if v, err := d.Int32(); err == nil {
			o.Value[i] = v
		} else {
			return err
		}
	}
	return nil
}

type GlGetUniformLocation_Postback struct {
	Result UniformLocation
}

func (o *GlGetUniformLocation_Postback) Decode(d *protocol.Decoder) error {
	{
		var x int32
		if v, err := d.Int32(); err == nil {
			x = v
		} else {
			return err
		}
		o.Result = UniformLocation(x)
	}
	return nil
}

type GlGetAttribLocation_Postback struct {
	Result AttributeLocation
}

func (o *GlGetAttribLocation_Postback) Decode(d *protocol.Decoder) error {
	{
		var x int32
		if v, err := d.Int32(); err == nil {
			x = v
		} else {
			return err
		}
		o.Result = AttributeLocation(x)
	}
	return nil
}

type GlGetTexParameteriv_Postback struct {
	Values S32Array
}

func (o *GlGetTexParameteriv_Postback) Decode(values_cnt uint64, d *protocol.Decoder) error {
	o.Values = make(S32Array, values_cnt)
	for i := range o.Values {
		if v, err := d.Int32(); err == nil {
			o.Values[i] = v
		} else {
			return err
		}
	}
	return nil
}

type GlGetTexParameterfv_Postback struct {
	Values F32Array
}

func (o *GlGetTexParameterfv_Postback) Decode(values_cnt uint64, d *protocol.Decoder) error {
	o.Values = make(F32Array, values_cnt)
	for i := range o.Values {
		if v, err := d.Float32(); err == nil {
			o.Values[i] = v
		} else {
			return err
		}
	}
	return nil
}

type GlGetShaderPrecisionFormat_Postback struct {
	Range     S32Array
	Precision int32
}

func (o *GlGetShaderPrecisionFormat_Postback) Decode(range_cnt uint64, d *protocol.Decoder) error {
	o.Range = make(S32Array, range_cnt)
	for i := range o.Range {
		if v, err := d.Int32(); err == nil {
			o.Range[i] = v
		} else {
			return err
		}
	}
	if v, err := d.Int32(); err == nil {
		o.Precision = v
	} else {
		return err
	}
	return nil
}

type GlGenTextures_Postback struct {
	Textures TextureIdArray
}

func (o *GlGenTextures_Postback) Decode(textures_cnt uint64, d *protocol.Decoder) error {
	o.Textures = make(TextureIdArray, textures_cnt)
	for i := range o.Textures {
		{
			var x uint32
			if v, err := d.Uint32(); err == nil {
				x = v
			} else {
				return err
			}
			o.Textures[i] = TextureId(x)
		}
	}
	return nil
}

type GlIsTexture_Postback struct {
	Result bool
}

func (o *GlIsTexture_Postback) Decode(d *protocol.Decoder) error {
	if v, err := d.Bool(); err == nil {
		o.Result = v
	} else {
		return err
	}
	return nil
}

type GlReadPixels_Postback struct {
	Data []byte
}

func (o *GlReadPixels_Postback) Decode(data_cnt uint64, d *protocol.Decoder) error {
	if val, err := readBytes(d, data_cnt); err == nil {
		o.Data = val
	} else {
		return err
	}
	return nil
}

type GlGenFramebuffers_Postback struct {
	Framebuffers FramebufferIdArray
}

func (o *GlGenFramebuffers_Postback) Decode(framebuffers_cnt uint64, d *protocol.Decoder) error {
	o.Framebuffers = make(FramebufferIdArray, framebuffers_cnt)
	for i := range o.Framebuffers {
		{
			var x uint32
			if v, err := d.Uint32(); err == nil {
				x = v
			} else {
				return err
			}
			o.Framebuffers[i] = FramebufferId(x)
		}
	}
	return nil
}

type GlCheckFramebufferStatus_Postback struct {
	Result FramebufferStatus
}

func (o *GlCheckFramebufferStatus_Postback) Decode(d *protocol.Decoder) error {
	if v, err := d.Uint32(); err == nil {
		o.Result = FramebufferStatus(v)
	} else {
		return err
	}
	return nil
}

type GlIsFramebuffer_Postback struct {
	Result bool
}

func (o *GlIsFramebuffer_Postback) Decode(d *protocol.Decoder) error {
	if v, err := d.Bool(); err == nil {
		o.Result = v
	} else {
		return err
	}
	return nil
}

type GlGenRenderbuffers_Postback struct {
	Renderbuffers RenderbufferIdArray
}

func (o *GlGenRenderbuffers_Postback) Decode(renderbuffers_cnt uint64, d *protocol.Decoder) error {
	o.Renderbuffers = make(RenderbufferIdArray, renderbuffers_cnt)
	for i := range o.Renderbuffers {
		{
			var x uint32
			if v, err := d.Uint32(); err == nil {
				x = v
			} else {
				return err
			}
			o.Renderbuffers[i] = RenderbufferId(x)
		}
	}
	return nil
}

type GlIsRenderbuffer_Postback struct {
	Result bool
}

func (o *GlIsRenderbuffer_Postback) Decode(d *protocol.Decoder) error {
	if v, err := d.Bool(); err == nil {
		o.Result = v
	} else {
		return err
	}
	return nil
}

type GlGetRenderbufferParameteriv_Postback struct {
	Values S32Array
}

func (o *GlGetRenderbufferParameteriv_Postback) Decode(values_cnt uint64, d *protocol.Decoder) error {
	o.Values = make(S32Array, values_cnt)
	for i := range o.Values {
		if v, err := d.Int32(); err == nil {
			o.Values[i] = v
		} else {
			return err
		}
	}
	return nil
}

type GlGenBuffers_Postback struct {
	Buffers BufferIdArray
}

func (o *GlGenBuffers_Postback) Decode(buffers_cnt uint64, d *protocol.Decoder) error {
	o.Buffers = make(BufferIdArray, buffers_cnt)
	for i := range o.Buffers {
		{
			var x uint32
			if v, err := d.Uint32(); err == nil {
				x = v
			} else {
				return err
			}
			o.Buffers[i] = BufferId(x)
		}
	}
	return nil
}

type GlIsBuffer_Postback struct {
	Result bool
}

func (o *GlIsBuffer_Postback) Decode(d *protocol.Decoder) error {
	if v, err := d.Bool(); err == nil {
		o.Result = v
	} else {
		return err
	}
	return nil
}

type GlGetBufferParameteriv_Postback struct {
	Value int32
}

func (o *GlGetBufferParameteriv_Postback) Decode(d *protocol.Decoder) error {
	if v, err := d.Int32(); err == nil {
		o.Value = v
	} else {
		return err
	}
	return nil
}

type GlCreateShader_Postback struct {
	Result ShaderId
}

func (o *GlCreateShader_Postback) Decode(d *protocol.Decoder) error {
	{
		var x uint32
		if v, err := d.Uint32(); err == nil {
			x = v
		} else {
			return err
		}
		o.Result = ShaderId(x)
	}
	return nil
}

type GlGetShaderInfoLog_Postback struct {
	StringLengthWritten int32
	Info                string
}

func (o *GlGetShaderInfoLog_Postback) Decode(info_cnt uint64, d *protocol.Decoder) error {
	if v, err := d.Int32(); err == nil {
		o.StringLengthWritten = v
	} else {
		return err
	}
	if val, err := readString(d, info_cnt); err == nil {
		o.Info = val
	} else {
		return err
	}
	return nil
}

type GlGetShaderSource_Postback struct {
	StringLengthWritten int32
	Source              string
}

func (o *GlGetShaderSource_Postback) Decode(source_cnt uint64, d *protocol.Decoder) error {
	if v, err := d.Int32(); err == nil {
		o.StringLengthWritten = v
	} else {
		return err
	}
	if val, err := readString(d, source_cnt); err == nil {
		o.Source = val
	} else {
		return err
	}
	return nil
}

type GlIsShader_Postback struct {
	Result bool
}

func (o *GlIsShader_Postback) Decode(d *protocol.Decoder) error {
	if v, err := d.Bool(); err == nil {
		o.Result = v
	} else {
		return err
	}
	return nil
}

type GlCreateProgram_Postback struct {
	Result ProgramId
}

func (o *GlCreateProgram_Postback) Decode(d *protocol.Decoder) error {
	{
		var x uint32
		if v, err := d.Uint32(); err == nil {
			x = v
		} else {
			return err
		}
		o.Result = ProgramId(x)
	}
	return nil
}

type GlGetAttachedShaders_Postback struct {
	ShadersLengthWritten int32
	Shaders              ShaderIdArray
}

func (o *GlGetAttachedShaders_Postback) Decode(shaders_cnt uint64, d *protocol.Decoder) error {
	if v, err := d.Int32(); err == nil {
		o.ShadersLengthWritten = v
	} else {
		return err
	}
	o.Shaders = make(ShaderIdArray, shaders_cnt)
	for i := range o.Shaders {
		{
			var x uint32
			if v, err := d.Uint32(); err == nil {
				x = v
			} else {
				return err
			}
			o.Shaders[i] = ShaderId(x)
		}
	}
	return nil
}

type GlGetProgramInfoLog_Postback struct {
	StringLengthWritten int32
	Info                string
}

func (o *GlGetProgramInfoLog_Postback) Decode(info_cnt uint64, d *protocol.Decoder) error {
	if v, err := d.Int32(); err == nil {
		o.StringLengthWritten = v
	} else {
		return err
	}
	if val, err := readString(d, info_cnt); err == nil {
		o.Info = val
	} else {
		return err
	}
	return nil
}

type GlIsProgram_Postback struct {
	Result bool
}

func (o *GlIsProgram_Postback) Decode(d *protocol.Decoder) error {
	if v, err := d.Bool(); err == nil {
		o.Result = v
	} else {
		return err
	}
	return nil
}

type GlGetFramebufferAttachmentParameteriv_Postback struct {
	Value S32Array
}

func (o *GlGetFramebufferAttachmentParameteriv_Postback) Decode(value_cnt uint64, d *protocol.Decoder) error {
	o.Value = make(S32Array, value_cnt)
	for i := range o.Value {
		if v, err := d.Int32(); err == nil {
			o.Value[i] = v
		} else {
			return err
		}
	}
	return nil
}

type GlGetBooleanv_Postback struct {
	Values BoolArray
}

func (o *GlGetBooleanv_Postback) Decode(values_cnt uint64, d *protocol.Decoder) error {
	o.Values = make(BoolArray, values_cnt)
	for i := range o.Values {
		if v, err := d.Bool(); err == nil {
			o.Values[i] = v
		} else {
			return err
		}
	}
	return nil
}

type GlGetFloatv_Postback struct {
	Values F32Array
}

func (o *GlGetFloatv_Postback) Decode(values_cnt uint64, d *protocol.Decoder) error {
	o.Values = make(F32Array, values_cnt)
	for i := range o.Values {
		if v, err := d.Float32(); err == nil {
			o.Values[i] = v
		} else {
			return err
		}
	}
	return nil
}

type GlGetIntegerv_Postback struct {
	Values S32Array
}

func (o *GlGetIntegerv_Postback) Decode(values_cnt uint64, d *protocol.Decoder) error {
	o.Values = make(S32Array, values_cnt)
	for i := range o.Values {
		if v, err := d.Int32(); err == nil {
			o.Values[i] = v
		} else {
			return err
		}
	}
	return nil
}

type GlGetString_Postback struct {
	Result string
}

func (o *GlGetString_Postback) Decode(result_cnt uint64, d *protocol.Decoder) error {
	if val, err := readString(d, result_cnt); err == nil {
		o.Result = val
	} else {
		return err
	}
	return nil
}

type GlIsEnabled_Postback struct {
	Result bool
}

func (o *GlIsEnabled_Postback) Decode(d *protocol.Decoder) error {
	if v, err := d.Bool(); err == nil {
		o.Result = v
	} else {
		return err
	}
	return nil
}

type GlMapBufferRange_Postback struct {
	Result []byte
}

func (o *GlMapBufferRange_Postback) Decode(result_cnt uint64, d *protocol.Decoder) error {
	if val, err := readBytes(d, result_cnt); err == nil {
		o.Result = val
	} else {
		return err
	}
	return nil
}

type GlGenQueries_Postback struct {
	Queries QueryIdArray
}

func (o *GlGenQueries_Postback) Decode(queries_cnt uint64, d *protocol.Decoder) error {
	o.Queries = make(QueryIdArray, queries_cnt)
	for i := range o.Queries {
		{
			var x uint32
			if v, err := d.Uint32(); err == nil {
				x = v
			} else {
				return err
			}
			o.Queries[i] = QueryId(x)
		}
	}
	return nil
}

type GlIsQuery_Postback struct {
	Result bool
}

func (o *GlIsQuery_Postback) Decode(d *protocol.Decoder) error {
	if v, err := d.Bool(); err == nil {
		o.Result = v
	} else {
		return err
	}
	return nil
}

type GlGetQueryiv_Postback struct {
	Value int32
}

func (o *GlGetQueryiv_Postback) Decode(d *protocol.Decoder) error {
	if v, err := d.Int32(); err == nil {
		o.Value = v
	} else {
		return err
	}
	return nil
}

type GlGetQueryObjectuiv_Postback struct {
	Value uint32
}

func (o *GlGetQueryObjectuiv_Postback) Decode(d *protocol.Decoder) error {
	if v, err := d.Uint32(); err == nil {
		o.Value = v
	} else {
		return err
	}
	return nil
}

type GlGenQueriesEXT_Postback struct {
	Queries QueryIdArray
}

func (o *GlGenQueriesEXT_Postback) Decode(queries_cnt uint64, d *protocol.Decoder) error {
	o.Queries = make(QueryIdArray, queries_cnt)
	for i := range o.Queries {
		{
			var x uint32
			if v, err := d.Uint32(); err == nil {
				x = v
			} else {
				return err
			}
			o.Queries[i] = QueryId(x)
		}
	}
	return nil
}

type GlIsQueryEXT_Postback struct {
	Result bool
}

func (o *GlIsQueryEXT_Postback) Decode(d *protocol.Decoder) error {
	if v, err := d.Bool(); err == nil {
		o.Result = v
	} else {
		return err
	}
	return nil
}

type GlGetQueryivEXT_Postback struct {
	Value int32
}

func (o *GlGetQueryivEXT_Postback) Decode(d *protocol.Decoder) error {
	if v, err := d.Int32(); err == nil {
		o.Value = v
	} else {
		return err
	}
	return nil
}

type GlGetQueryObjectivEXT_Postback struct {
	Value int32
}

func (o *GlGetQueryObjectivEXT_Postback) Decode(d *protocol.Decoder) error {
	if v, err := d.Int32(); err == nil {
		o.Value = v
	} else {
		return err
	}
	return nil
}

type GlGetQueryObjectuivEXT_Postback struct {
	Value uint32
}

func (o *GlGetQueryObjectuivEXT_Postback) Decode(d *protocol.Decoder) error {
	if v, err := d.Uint32(); err == nil {
		o.Value = v
	} else {
		return err
	}
	return nil
}

type GlGetQueryObjecti64vEXT_Postback struct {
	Value int64
}

func (o *GlGetQueryObjecti64vEXT_Postback) Decode(d *protocol.Decoder) error {
	if v, err := d.Int64(); err == nil {
		o.Value = v
	} else {
		return err
	}
	return nil
}

type GlGetQueryObjectui64vEXT_Postback struct {
	Value uint64
}

func (o *GlGetQueryObjectui64vEXT_Postback) Decode(d *protocol.Decoder) error {
	if v, err := d.Uint64(); err == nil {
		o.Value = v
	} else {
		return err
	}
	return nil
}

type replayWriter struct {
	builder *builder.Builder
	state   *state
}

func newReplayWriter(b *builder.Builder) *replayWriter {
	return &replayWriter{
		builder: b,
		state:   initialState(),
	}
}
func loadRemap(b *builder.Builder, key interface{}, val value.Value) {
	ptr, found := b.Remappings[key]
	if found {
		b.Load(val.Type(), ptr)
	} else {
		ptr = b.AllocateMemory(uint64(val.Type().Size(b.PointerSize())))
		b.Push(val) // We have an input to an unknown id, use the unmapped value.
		b.Clone(0)
		b.Store(ptr)
		b.Remappings[key] = ptr
	}
}
func storeRemap(b *builder.Builder, key interface{}, val value.Pointer, ty protocol.Type) {
	ptr, found := b.Remappings[key]
	if !found {
		ptr = b.AllocateMemory(uint64(ty.Size(b.PointerSize())))
		b.Load(ty, val)
		b.Store(ptr)
		b.Remappings[key] = ptr
	}
}
func (r *replayWriter) Write(id atom.ID, a atom.Atom, wantOutput bool) {
	b := r.builder
	switch ω := a.(type) {
	case *memory.Observation:
		b.Observation(ω.Range, ω.ResourceID)
	case replayer:
		ω.replay(id, r.state, b, wantOutput)
	case *atom.EOS:
	default:
		panic(fmt.Errorf("Unsupported atom type %T for Write", ω))
	}
	b.EndAtom()
}

var _ = replayer(&Init{}) // interface compliance check
func (ω *Init) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.S32(ω.In.Width))
	b.Push(value.S32(ω.In.Height))
	b.Push(value.U32(ω.In.ColorFmt))
	b.Push(value.U32(ω.In.DepthFmt))
	b.Push(value.U32(ω.In.StencilFmt))
	b.CallNoPush(funcInfoInit)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&StartTimer{}) // interface compliance check
func (ω *StartTimer) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U8(ω.In.Index))
	b.CallNoPush(funcInfoStartTimer)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&StopTimer{}) // interface compliance check
func (ω *StopTimer) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{8 /* result */})
	b.Push(value.U8(ω.In.Index))
	b.CallPush(funcInfoStopTimer)
	b.Store(outputs[0])
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := StopTimer_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&FlushPostBuffer{}) // interface compliance check
func (ω *FlushPostBuffer) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.CallNoPush(funcInfoFlushPostBuffer)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&EglCreateContext{}) // interface compliance check
func (ω *EglCreateContext) defaultReplay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{4 /* version */, 4 /* context */})
	b.Push(outputs[0]) // version
	b.Push(outputs[1]) // context
	b.CallNoPush(funcInfoEglCreateContext)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := EglCreateContext_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&EglMakeCurrent{}) // interface compliance check
func (ω *EglMakeCurrent) defaultReplay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.S32(ω.In.Context))
	b.CallNoPush(funcInfoEglMakeCurrent)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&EglSwapBuffers{}) // interface compliance check
func (ω *EglSwapBuffers) defaultReplay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.CallNoPush(funcInfoEglSwapBuffers)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlEnableClientState{}) // interface compliance check
func (ω *GlEnableClientState) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Type))
	b.CallNoPush(funcInfoGlEnableClientState)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlDisableClientState{}) // interface compliance check
func (ω *GlDisableClientState) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Type))
	b.CallNoPush(funcInfoGlDisableClientState)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlGetProgramBinaryOES{}) // interface compliance check
func (ω *GlGetProgramBinaryOES) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	binary_cnt := uint64(ω.In.BufferSize)
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{4 /* bytes_written */, 4 /* binary_format */, binary_cnt /* binary */})
	if key, remap := ω.In.Program.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Program.value(b, ω, s))
	} else {
		b.Push(ω.In.Program.value(b, ω, s))
	}
	b.Push(value.S32(ω.In.BufferSize))
	b.Push(outputs[0]) // bytes_written
	b.Push(outputs[1]) // binary_format
	b.Push(outputs[2]) // binary
	b.CallNoPush(funcInfoGlGetProgramBinaryOES)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGetProgramBinaryOES_Postback{}
			if err := postback.Decode(binary_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlProgramBinaryOES{}) // interface compliance check
func (ω *GlProgramBinaryOES) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Program.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Program.value(b, ω, s))
	} else {
		b.Push(ω.In.Program.value(b, ω, s))
	}
	b.Push(value.U32(ω.In.BinaryFormat))
	b.Push(value.VolatileCapturePointer(uint64(ω.In.Binary)))
	b.Push(value.S32(ω.In.BinarySize))
	b.CallNoPush(funcInfoGlProgramBinaryOES)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlStartTilingQCOM{}) // interface compliance check
func (ω *GlStartTilingQCOM) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.S32(ω.In.X))
	b.Push(value.S32(ω.In.Y))
	b.Push(value.S32(ω.In.Width))
	b.Push(value.S32(ω.In.Height))
	b.Push(value.U32(ω.In.PreserveMask))
	b.CallNoPush(funcInfoGlStartTilingQCOM)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlEndTilingQCOM{}) // interface compliance check
func (ω *GlEndTilingQCOM) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.PreserveMask))
	b.CallNoPush(funcInfoGlEndTilingQCOM)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlDiscardFramebufferEXT{}) // interface compliance check
func (ω *GlDiscardFramebufferEXT) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Target))
	b.Push(value.S32(ω.In.NumAttachments))
	b.Push(ω.In.Attachments.value(b, ω, s))
	b.CallNoPush(funcInfoGlDiscardFramebufferEXT)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlInsertEventMarkerEXT{}) // interface compliance check
func (ω *GlInsertEventMarkerEXT) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.S32(ω.In.Length))
	b.Push(b.String(ω.In.Marker))
	b.CallNoPush(funcInfoGlInsertEventMarkerEXT)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlPushGroupMarkerEXT{}) // interface compliance check
func (ω *GlPushGroupMarkerEXT) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.S32(ω.In.Length))
	b.Push(b.String(ω.In.Marker))
	b.CallNoPush(funcInfoGlPushGroupMarkerEXT)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlPopGroupMarkerEXT{}) // interface compliance check
func (ω *GlPopGroupMarkerEXT) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.CallNoPush(funcInfoGlPopGroupMarkerEXT)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlTexStorage1DEXT{}) // interface compliance check
func (ω *GlTexStorage1DEXT) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Target))
	b.Push(value.S32(ω.In.Levels))
	b.Push(value.U32(ω.In.Format))
	b.Push(value.S32(ω.In.Width))
	b.CallNoPush(funcInfoGlTexStorage1DEXT)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlTexStorage2DEXT{}) // interface compliance check
func (ω *GlTexStorage2DEXT) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Target))
	b.Push(value.S32(ω.In.Levels))
	b.Push(value.U32(ω.In.Format))
	b.Push(value.S32(ω.In.Width))
	b.Push(value.S32(ω.In.Height))
	b.CallNoPush(funcInfoGlTexStorage2DEXT)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlTexStorage3DEXT{}) // interface compliance check
func (ω *GlTexStorage3DEXT) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Target))
	b.Push(value.S32(ω.In.Levels))
	b.Push(value.U32(ω.In.Format))
	b.Push(value.S32(ω.In.Width))
	b.Push(value.S32(ω.In.Height))
	b.Push(value.S32(ω.In.Depth))
	b.CallNoPush(funcInfoGlTexStorage3DEXT)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlTextureStorage1DEXT{}) // interface compliance check
func (ω *GlTextureStorage1DEXT) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Texture.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Texture.value(b, ω, s))
	} else {
		b.Push(ω.In.Texture.value(b, ω, s))
	}
	b.Push(value.U32(ω.In.Target))
	b.Push(value.S32(ω.In.Levels))
	b.Push(value.U32(ω.In.Format))
	b.Push(value.S32(ω.In.Width))
	b.CallNoPush(funcInfoGlTextureStorage1DEXT)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlTextureStorage2DEXT{}) // interface compliance check
func (ω *GlTextureStorage2DEXT) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Texture.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Texture.value(b, ω, s))
	} else {
		b.Push(ω.In.Texture.value(b, ω, s))
	}
	b.Push(value.U32(ω.In.Target))
	b.Push(value.S32(ω.In.Levels))
	b.Push(value.U32(ω.In.Format))
	b.Push(value.S32(ω.In.Width))
	b.Push(value.S32(ω.In.Height))
	b.CallNoPush(funcInfoGlTextureStorage2DEXT)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlTextureStorage3DEXT{}) // interface compliance check
func (ω *GlTextureStorage3DEXT) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Texture.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Texture.value(b, ω, s))
	} else {
		b.Push(ω.In.Texture.value(b, ω, s))
	}
	b.Push(value.U32(ω.In.Target))
	b.Push(value.S32(ω.In.Levels))
	b.Push(value.U32(ω.In.Format))
	b.Push(value.S32(ω.In.Width))
	b.Push(value.S32(ω.In.Height))
	b.Push(value.S32(ω.In.Depth))
	b.CallNoPush(funcInfoGlTextureStorage3DEXT)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlGenVertexArraysOES{}) // interface compliance check
func (ω *GlGenVertexArraysOES) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	arrays_cnt := uint64(ω.In.Count)
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{arrays_cnt * 4 /* arrays */})
	b.Push(value.S32(ω.In.Count))
	b.Push(outputs[0]) // arrays
	b.CallNoPush(funcInfoGlGenVertexArraysOES)
	StateMutator{State: s}.Write(id, ω)
	for i, e := range ω.Out.Arrays {
		ptr := outputs[0].Offset(uint64(i * 4))
		if key, remap := e.remap(ω, s); remap {
			storeRemap(b, key, ptr, protocol.TypeUint32)
		}
	}
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGenVertexArraysOES_Postback{}
			if err := postback.Decode(arrays_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlBindVertexArrayOES{}) // interface compliance check
func (ω *GlBindVertexArrayOES) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Array.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Array.value(b, ω, s))
	} else {
		b.Push(ω.In.Array.value(b, ω, s))
	}
	b.CallNoPush(funcInfoGlBindVertexArrayOES)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlDeleteVertexArraysOES{}) // interface compliance check
func (ω *GlDeleteVertexArraysOES) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.S32(ω.In.Count))
	b.Push(ω.In.Arrays.value(b, ω, s))
	b.CallNoPush(funcInfoGlDeleteVertexArraysOES)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlIsVertexArrayOES{}) // interface compliance check
func (ω *GlIsVertexArrayOES) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{1 /* result */})
	if key, remap := ω.In.Array.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Array.value(b, ω, s))
	} else {
		b.Push(ω.In.Array.value(b, ω, s))
	}
	b.CallPush(funcInfoGlIsVertexArrayOES)
	b.Store(outputs[0])
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlIsVertexArrayOES_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlEGLImageTargetTexture2DOES{}) // interface compliance check
func (ω *GlEGLImageTargetTexture2DOES) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Target))
	b.Push(ω.In.Image.value(b, ω, s))
	b.CallNoPush(funcInfoGlEGLImageTargetTexture2DOES)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlEGLImageTargetRenderbufferStorageOES{}) // interface compliance check
func (ω *GlEGLImageTargetRenderbufferStorageOES) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Target))
	b.Push(ω.In.Image.value(b, ω, s))
	b.CallNoPush(funcInfoGlEGLImageTargetRenderbufferStorageOES)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlGetGraphicsResetStatusEXT{}) // interface compliance check
func (ω *GlGetGraphicsResetStatusEXT) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{4 /* result */})
	b.CallPush(funcInfoGlGetGraphicsResetStatusEXT)
	b.Store(outputs[0])
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGetGraphicsResetStatusEXT_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlBindAttribLocation{}) // interface compliance check
func (ω *GlBindAttribLocation) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Program.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Program.value(b, ω, s))
	} else {
		b.Push(ω.In.Program.value(b, ω, s))
	}
	b.Push(ω.In.Location.value(b, ω, s))
	b.Push(b.String(ω.In.Name))
	b.CallNoPush(funcInfoGlBindAttribLocation)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlBlendFunc{}) // interface compliance check
func (ω *GlBlendFunc) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.SrcFactor))
	b.Push(value.U32(ω.In.DstFactor))
	b.CallNoPush(funcInfoGlBlendFunc)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlBlendFuncSeparate{}) // interface compliance check
func (ω *GlBlendFuncSeparate) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.SrcFactorRgb))
	b.Push(value.U32(ω.In.DstFactorRgb))
	b.Push(value.U32(ω.In.SrcFactorAlpha))
	b.Push(value.U32(ω.In.DstFactorAlpha))
	b.CallNoPush(funcInfoGlBlendFuncSeparate)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlBlendEquation{}) // interface compliance check
func (ω *GlBlendEquation) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Equation))
	b.CallNoPush(funcInfoGlBlendEquation)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlBlendEquationSeparate{}) // interface compliance check
func (ω *GlBlendEquationSeparate) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Rgb))
	b.Push(value.U32(ω.In.Alpha))
	b.CallNoPush(funcInfoGlBlendEquationSeparate)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlBlendColor{}) // interface compliance check
func (ω *GlBlendColor) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.F32(ω.In.Red))
	b.Push(value.F32(ω.In.Green))
	b.Push(value.F32(ω.In.Blue))
	b.Push(value.F32(ω.In.Alpha))
	b.CallNoPush(funcInfoGlBlendColor)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlEnableVertexAttribArray{}) // interface compliance check
func (ω *GlEnableVertexAttribArray) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(ω.In.Location.value(b, ω, s))
	b.CallNoPush(funcInfoGlEnableVertexAttribArray)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlDisableVertexAttribArray{}) // interface compliance check
func (ω *GlDisableVertexAttribArray) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(ω.In.Location.value(b, ω, s))
	b.CallNoPush(funcInfoGlDisableVertexAttribArray)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlVertexAttribPointer{}) // interface compliance check
func (ω *GlVertexAttribPointer) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(ω.In.Location.value(b, ω, s))
	b.Push(value.U32(ω.In.Size))
	b.Push(value.U32(ω.In.Type))
	b.Push(value.Bool(ω.In.Normalized))
	b.Push(value.S32(ω.In.Stride))
	b.Push(ω.In.Data.value(b, ω, s))
	b.CallNoPush(funcInfoGlVertexAttribPointer)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlGetActiveAttrib{}) // interface compliance check
func (ω *GlGetActiveAttrib) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	name_cnt := uint64(ω.In.BufferSize)
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{4 /* buffer_bytes_written */, 4 /* vector_count */, 4 /* type */, name_cnt /* name */})
	if key, remap := ω.In.Program.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Program.value(b, ω, s))
	} else {
		b.Push(ω.In.Program.value(b, ω, s))
	}
	b.Push(ω.In.Location.value(b, ω, s))
	b.Push(value.S32(ω.In.BufferSize))
	b.Push(outputs[0]) // buffer_bytes_written
	b.Push(outputs[1]) // vector_count
	b.Push(outputs[2]) // type
	b.Push(outputs[3]) // name
	b.CallNoPush(funcInfoGlGetActiveAttrib)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGetActiveAttrib_Postback{}
			if err := postback.Decode(name_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlGetActiveUniform{}) // interface compliance check
func (ω *GlGetActiveUniform) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	name_cnt := uint64(ω.In.BufferSize)
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{4 /* buffer_bytes_written */, 4 /* size */, 4 /* type */, name_cnt /* name */})
	if key, remap := ω.In.Program.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Program.value(b, ω, s))
	} else {
		b.Push(ω.In.Program.value(b, ω, s))
	}
	b.Push(value.S32(ω.In.Location))
	b.Push(value.S32(ω.In.BufferSize))
	b.Push(outputs[0]) // buffer_bytes_written
	b.Push(outputs[1]) // size
	b.Push(outputs[2]) // type
	b.Push(outputs[3]) // name
	b.CallNoPush(funcInfoGlGetActiveUniform)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGetActiveUniform_Postback{}
			if err := postback.Decode(name_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlGetError{}) // interface compliance check
func (ω *GlGetError) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{4 /* result */})
	b.CallPush(funcInfoGlGetError)
	b.Store(outputs[0])
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGetError_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlGetProgramiv{}) // interface compliance check
func (ω *GlGetProgramiv) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	value_cnt := uint64(1)
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{value_cnt * 4 /* value */})
	if key, remap := ω.In.Program.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Program.value(b, ω, s))
	} else {
		b.Push(ω.In.Program.value(b, ω, s))
	}
	b.Push(value.U32(ω.In.Parameter))
	b.Push(outputs[0]) // value
	b.CallNoPush(funcInfoGlGetProgramiv)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGetProgramiv_Postback{}
			if err := postback.Decode(value_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlGetShaderiv{}) // interface compliance check
func (ω *GlGetShaderiv) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	value_cnt := uint64(1)
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{value_cnt * 4 /* value */})
	if key, remap := ω.In.Shader.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Shader.value(b, ω, s))
	} else {
		b.Push(ω.In.Shader.value(b, ω, s))
	}
	b.Push(value.U32(ω.In.Parameter))
	b.Push(outputs[0]) // value
	b.CallNoPush(funcInfoGlGetShaderiv)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGetShaderiv_Postback{}
			if err := postback.Decode(value_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlGetUniformLocation{}) // interface compliance check
func (ω *GlGetUniformLocation) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{4 /* result */})
	if key, remap := ω.In.Program.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Program.value(b, ω, s))
	} else {
		b.Push(ω.In.Program.value(b, ω, s))
	}
	b.Push(b.String(ω.In.Name))
	b.CallPush(funcInfoGlGetUniformLocation)
	b.Store(outputs[0])
	StateMutator{State: s}.Write(id, ω)
	if key, remap := ω.Out.Result.remap(ω, s); remap {
		storeRemap(b, key, outputs[0], protocol.TypeInt32)
	}
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGetUniformLocation_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlGetAttribLocation{}) // interface compliance check
func (ω *GlGetAttribLocation) defaultReplay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{4 /* result */})
	if key, remap := ω.In.Program.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Program.value(b, ω, s))
	} else {
		b.Push(ω.In.Program.value(b, ω, s))
	}
	b.Push(b.String(ω.In.Name))
	b.CallPush(funcInfoGlGetAttribLocation)
	b.Store(outputs[0])
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGetAttribLocation_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlPixelStorei{}) // interface compliance check
func (ω *GlPixelStorei) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Parameter))
	b.Push(value.S32(ω.In.Value))
	b.CallNoPush(funcInfoGlPixelStorei)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlTexParameteri{}) // interface compliance check
func (ω *GlTexParameteri) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Target))
	b.Push(value.U32(ω.In.Parameter))
	b.Push(value.S32(ω.In.Value))
	b.CallNoPush(funcInfoGlTexParameteri)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlTexParameterf{}) // interface compliance check
func (ω *GlTexParameterf) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Target))
	b.Push(value.U32(ω.In.Parameter))
	b.Push(value.F32(ω.In.Value))
	b.CallNoPush(funcInfoGlTexParameterf)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlGetTexParameteriv{}) // interface compliance check
func (ω *GlGetTexParameteriv) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	values_cnt := uint64(1)
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{values_cnt * 4 /* values */})
	b.Push(value.U32(ω.In.Target))
	b.Push(value.U32(ω.In.Parameter))
	b.Push(outputs[0]) // values
	b.CallNoPush(funcInfoGlGetTexParameteriv)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGetTexParameteriv_Postback{}
			if err := postback.Decode(values_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlGetTexParameterfv{}) // interface compliance check
func (ω *GlGetTexParameterfv) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	values_cnt := uint64(1)
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{values_cnt * 4 /* values */})
	b.Push(value.U32(ω.In.Target))
	b.Push(value.U32(ω.In.Parameter))
	b.Push(outputs[0]) // values
	b.CallNoPush(funcInfoGlGetTexParameterfv)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGetTexParameterfv_Postback{}
			if err := postback.Decode(values_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlUniform1i{}) // interface compliance check
func (ω *GlUniform1i) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Location.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Location.value(b, ω, s))
	} else {
		b.Push(ω.In.Location.value(b, ω, s))
	}
	b.Push(value.S32(ω.In.Value))
	b.CallNoPush(funcInfoGlUniform1i)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlUniform2i{}) // interface compliance check
func (ω *GlUniform2i) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Location.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Location.value(b, ω, s))
	} else {
		b.Push(ω.In.Location.value(b, ω, s))
	}
	b.Push(value.S32(ω.In.Value0))
	b.Push(value.S32(ω.In.Value1))
	b.CallNoPush(funcInfoGlUniform2i)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlUniform3i{}) // interface compliance check
func (ω *GlUniform3i) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Location.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Location.value(b, ω, s))
	} else {
		b.Push(ω.In.Location.value(b, ω, s))
	}
	b.Push(value.S32(ω.In.Value0))
	b.Push(value.S32(ω.In.Value1))
	b.Push(value.S32(ω.In.Value2))
	b.CallNoPush(funcInfoGlUniform3i)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlUniform4i{}) // interface compliance check
func (ω *GlUniform4i) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Location.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Location.value(b, ω, s))
	} else {
		b.Push(ω.In.Location.value(b, ω, s))
	}
	b.Push(value.S32(ω.In.Value0))
	b.Push(value.S32(ω.In.Value1))
	b.Push(value.S32(ω.In.Value2))
	b.Push(value.S32(ω.In.Value3))
	b.CallNoPush(funcInfoGlUniform4i)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlUniform1iv{}) // interface compliance check
func (ω *GlUniform1iv) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Location.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Location.value(b, ω, s))
	} else {
		b.Push(ω.In.Location.value(b, ω, s))
	}
	b.Push(value.S32(ω.In.Count))
	b.Push(ω.In.Value.value(b, ω, s))
	b.CallNoPush(funcInfoGlUniform1iv)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlUniform2iv{}) // interface compliance check
func (ω *GlUniform2iv) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Location.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Location.value(b, ω, s))
	} else {
		b.Push(ω.In.Location.value(b, ω, s))
	}
	b.Push(value.S32(ω.In.Count))
	b.Push(ω.In.Value.value(b, ω, s))
	b.CallNoPush(funcInfoGlUniform2iv)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlUniform3iv{}) // interface compliance check
func (ω *GlUniform3iv) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Location.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Location.value(b, ω, s))
	} else {
		b.Push(ω.In.Location.value(b, ω, s))
	}
	b.Push(value.S32(ω.In.Count))
	b.Push(ω.In.Value.value(b, ω, s))
	b.CallNoPush(funcInfoGlUniform3iv)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlUniform4iv{}) // interface compliance check
func (ω *GlUniform4iv) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Location.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Location.value(b, ω, s))
	} else {
		b.Push(ω.In.Location.value(b, ω, s))
	}
	b.Push(value.S32(ω.In.Count))
	b.Push(ω.In.Value.value(b, ω, s))
	b.CallNoPush(funcInfoGlUniform4iv)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlUniform1f{}) // interface compliance check
func (ω *GlUniform1f) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Location.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Location.value(b, ω, s))
	} else {
		b.Push(ω.In.Location.value(b, ω, s))
	}
	b.Push(value.F32(ω.In.Value))
	b.CallNoPush(funcInfoGlUniform1f)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlUniform2f{}) // interface compliance check
func (ω *GlUniform2f) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Location.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Location.value(b, ω, s))
	} else {
		b.Push(ω.In.Location.value(b, ω, s))
	}
	b.Push(value.F32(ω.In.Value0))
	b.Push(value.F32(ω.In.Value1))
	b.CallNoPush(funcInfoGlUniform2f)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlUniform3f{}) // interface compliance check
func (ω *GlUniform3f) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Location.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Location.value(b, ω, s))
	} else {
		b.Push(ω.In.Location.value(b, ω, s))
	}
	b.Push(value.F32(ω.In.Value0))
	b.Push(value.F32(ω.In.Value1))
	b.Push(value.F32(ω.In.Value2))
	b.CallNoPush(funcInfoGlUniform3f)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlUniform4f{}) // interface compliance check
func (ω *GlUniform4f) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Location.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Location.value(b, ω, s))
	} else {
		b.Push(ω.In.Location.value(b, ω, s))
	}
	b.Push(value.F32(ω.In.Value0))
	b.Push(value.F32(ω.In.Value1))
	b.Push(value.F32(ω.In.Value2))
	b.Push(value.F32(ω.In.Value3))
	b.CallNoPush(funcInfoGlUniform4f)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlUniform1fv{}) // interface compliance check
func (ω *GlUniform1fv) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Location.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Location.value(b, ω, s))
	} else {
		b.Push(ω.In.Location.value(b, ω, s))
	}
	b.Push(value.S32(ω.In.Count))
	b.Push(ω.In.Value.value(b, ω, s))
	b.CallNoPush(funcInfoGlUniform1fv)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlUniform2fv{}) // interface compliance check
func (ω *GlUniform2fv) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Location.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Location.value(b, ω, s))
	} else {
		b.Push(ω.In.Location.value(b, ω, s))
	}
	b.Push(value.S32(ω.In.Count))
	b.Push(ω.In.Value.value(b, ω, s))
	b.CallNoPush(funcInfoGlUniform2fv)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlUniform3fv{}) // interface compliance check
func (ω *GlUniform3fv) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Location.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Location.value(b, ω, s))
	} else {
		b.Push(ω.In.Location.value(b, ω, s))
	}
	b.Push(value.S32(ω.In.Count))
	b.Push(ω.In.Value.value(b, ω, s))
	b.CallNoPush(funcInfoGlUniform3fv)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlUniform4fv{}) // interface compliance check
func (ω *GlUniform4fv) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Location.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Location.value(b, ω, s))
	} else {
		b.Push(ω.In.Location.value(b, ω, s))
	}
	b.Push(value.S32(ω.In.Count))
	b.Push(ω.In.Value.value(b, ω, s))
	b.CallNoPush(funcInfoGlUniform4fv)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlUniformMatrix2fv{}) // interface compliance check
func (ω *GlUniformMatrix2fv) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Location.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Location.value(b, ω, s))
	} else {
		b.Push(ω.In.Location.value(b, ω, s))
	}
	b.Push(value.S32(ω.In.Count))
	b.Push(value.Bool(ω.In.Transpose))
	b.Push(ω.In.Values.value(b, ω, s))
	b.CallNoPush(funcInfoGlUniformMatrix2fv)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlUniformMatrix3fv{}) // interface compliance check
func (ω *GlUniformMatrix3fv) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Location.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Location.value(b, ω, s))
	} else {
		b.Push(ω.In.Location.value(b, ω, s))
	}
	b.Push(value.S32(ω.In.Count))
	b.Push(value.Bool(ω.In.Transpose))
	b.Push(ω.In.Values.value(b, ω, s))
	b.CallNoPush(funcInfoGlUniformMatrix3fv)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlUniformMatrix4fv{}) // interface compliance check
func (ω *GlUniformMatrix4fv) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Location.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Location.value(b, ω, s))
	} else {
		b.Push(ω.In.Location.value(b, ω, s))
	}
	b.Push(value.S32(ω.In.Count))
	b.Push(value.Bool(ω.In.Transpose))
	b.Push(ω.In.Values.value(b, ω, s))
	b.CallNoPush(funcInfoGlUniformMatrix4fv)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlGetUniformfv{}) // interface compliance check
func (ω *GlGetUniformfv) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Program.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Program.value(b, ω, s))
	} else {
		b.Push(ω.In.Program.value(b, ω, s))
	}
	if key, remap := ω.In.Location.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Location.value(b, ω, s))
	} else {
		b.Push(ω.In.Location.value(b, ω, s))
	}
	b.Push(ω.In.Values.value(b, ω, s))
	b.CallNoPush(funcInfoGlGetUniformfv)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlGetUniformiv{}) // interface compliance check
func (ω *GlGetUniformiv) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Program.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Program.value(b, ω, s))
	} else {
		b.Push(ω.In.Program.value(b, ω, s))
	}
	if key, remap := ω.In.Location.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Location.value(b, ω, s))
	} else {
		b.Push(ω.In.Location.value(b, ω, s))
	}
	b.Push(ω.In.Values.value(b, ω, s))
	b.CallNoPush(funcInfoGlGetUniformiv)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlVertexAttrib1f{}) // interface compliance check
func (ω *GlVertexAttrib1f) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(ω.In.Location.value(b, ω, s))
	b.Push(value.F32(ω.In.Value0))
	b.CallNoPush(funcInfoGlVertexAttrib1f)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlVertexAttrib2f{}) // interface compliance check
func (ω *GlVertexAttrib2f) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(ω.In.Location.value(b, ω, s))
	b.Push(value.F32(ω.In.Value0))
	b.Push(value.F32(ω.In.Value1))
	b.CallNoPush(funcInfoGlVertexAttrib2f)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlVertexAttrib3f{}) // interface compliance check
func (ω *GlVertexAttrib3f) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(ω.In.Location.value(b, ω, s))
	b.Push(value.F32(ω.In.Value0))
	b.Push(value.F32(ω.In.Value1))
	b.Push(value.F32(ω.In.Value2))
	b.CallNoPush(funcInfoGlVertexAttrib3f)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlVertexAttrib4f{}) // interface compliance check
func (ω *GlVertexAttrib4f) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(ω.In.Location.value(b, ω, s))
	b.Push(value.F32(ω.In.Value0))
	b.Push(value.F32(ω.In.Value1))
	b.Push(value.F32(ω.In.Value2))
	b.Push(value.F32(ω.In.Value3))
	b.CallNoPush(funcInfoGlVertexAttrib4f)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlVertexAttrib1fv{}) // interface compliance check
func (ω *GlVertexAttrib1fv) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(ω.In.Location.value(b, ω, s))
	b.Push(ω.In.Value.value(b, ω, s))
	b.CallNoPush(funcInfoGlVertexAttrib1fv)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlVertexAttrib2fv{}) // interface compliance check
func (ω *GlVertexAttrib2fv) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(ω.In.Location.value(b, ω, s))
	b.Push(ω.In.Value.value(b, ω, s))
	b.CallNoPush(funcInfoGlVertexAttrib2fv)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlVertexAttrib3fv{}) // interface compliance check
func (ω *GlVertexAttrib3fv) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(ω.In.Location.value(b, ω, s))
	b.Push(ω.In.Value.value(b, ω, s))
	b.CallNoPush(funcInfoGlVertexAttrib3fv)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlVertexAttrib4fv{}) // interface compliance check
func (ω *GlVertexAttrib4fv) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(ω.In.Location.value(b, ω, s))
	b.Push(ω.In.Value.value(b, ω, s))
	b.CallNoPush(funcInfoGlVertexAttrib4fv)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlGetShaderPrecisionFormat{}) // interface compliance check
func (ω *GlGetShaderPrecisionFormat) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	range_cnt := uint64(2)
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{range_cnt * 4 /* range */, 4 /* precision */})
	b.Push(value.U32(ω.In.ShaderType))
	b.Push(value.U32(ω.In.PrecisionType))
	b.Push(outputs[0]) // range
	b.Push(outputs[1]) // precision
	b.CallNoPush(funcInfoGlGetShaderPrecisionFormat)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGetShaderPrecisionFormat_Postback{}
			if err := postback.Decode(range_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlDepthMask{}) // interface compliance check
func (ω *GlDepthMask) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.Bool(ω.In.Enabled))
	b.CallNoPush(funcInfoGlDepthMask)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlDepthFunc{}) // interface compliance check
func (ω *GlDepthFunc) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Function))
	b.CallNoPush(funcInfoGlDepthFunc)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlDepthRangef{}) // interface compliance check
func (ω *GlDepthRangef) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.F32(ω.In.Near))
	b.Push(value.F32(ω.In.Far))
	b.CallNoPush(funcInfoGlDepthRangef)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlColorMask{}) // interface compliance check
func (ω *GlColorMask) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.Bool(ω.In.Red))
	b.Push(value.Bool(ω.In.Green))
	b.Push(value.Bool(ω.In.Blue))
	b.Push(value.Bool(ω.In.Alpha))
	b.CallNoPush(funcInfoGlColorMask)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlStencilMask{}) // interface compliance check
func (ω *GlStencilMask) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Mask))
	b.CallNoPush(funcInfoGlStencilMask)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlStencilMaskSeparate{}) // interface compliance check
func (ω *GlStencilMaskSeparate) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Face))
	b.Push(value.U32(ω.In.Mask))
	b.CallNoPush(funcInfoGlStencilMaskSeparate)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlStencilFuncSeparate{}) // interface compliance check
func (ω *GlStencilFuncSeparate) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Face))
	b.Push(value.U32(ω.In.Function))
	b.Push(value.S32(ω.In.ReferenceValue))
	b.Push(value.S32(ω.In.Mask))
	b.CallNoPush(funcInfoGlStencilFuncSeparate)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlStencilOpSeparate{}) // interface compliance check
func (ω *GlStencilOpSeparate) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Face))
	b.Push(value.U32(ω.In.StencilFail))
	b.Push(value.U32(ω.In.StencilPassDepthFail))
	b.Push(value.U32(ω.In.StencilPassDepthPass))
	b.CallNoPush(funcInfoGlStencilOpSeparate)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlFrontFace{}) // interface compliance check
func (ω *GlFrontFace) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Orientation))
	b.CallNoPush(funcInfoGlFrontFace)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlViewport{}) // interface compliance check
func (ω *GlViewport) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.S32(ω.In.X))
	b.Push(value.S32(ω.In.Y))
	b.Push(value.S32(ω.In.Width))
	b.Push(value.S32(ω.In.Height))
	b.CallNoPush(funcInfoGlViewport)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlScissor{}) // interface compliance check
func (ω *GlScissor) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.S32(ω.In.X))
	b.Push(value.S32(ω.In.Y))
	b.Push(value.S32(ω.In.Width))
	b.Push(value.S32(ω.In.Height))
	b.CallNoPush(funcInfoGlScissor)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlActiveTexture{}) // interface compliance check
func (ω *GlActiveTexture) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Unit))
	b.CallNoPush(funcInfoGlActiveTexture)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlGenTextures{}) // interface compliance check
func (ω *GlGenTextures) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	textures_cnt := uint64(ω.In.Count)
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{textures_cnt * 4 /* textures */})
	b.Push(value.S32(ω.In.Count))
	b.Push(outputs[0]) // textures
	b.CallNoPush(funcInfoGlGenTextures)
	StateMutator{State: s}.Write(id, ω)
	for i, e := range ω.Out.Textures {
		ptr := outputs[0].Offset(uint64(i * 4))
		if key, remap := e.remap(ω, s); remap {
			storeRemap(b, key, ptr, protocol.TypeUint32)
		}
	}
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGenTextures_Postback{}
			if err := postback.Decode(textures_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlDeleteTextures{}) // interface compliance check
func (ω *GlDeleteTextures) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.S32(ω.In.Count))
	b.Push(ω.In.Textures.value(b, ω, s))
	b.CallNoPush(funcInfoGlDeleteTextures)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlIsTexture{}) // interface compliance check
func (ω *GlIsTexture) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{1 /* result */})
	if key, remap := ω.In.Texture.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Texture.value(b, ω, s))
	} else {
		b.Push(ω.In.Texture.value(b, ω, s))
	}
	b.CallPush(funcInfoGlIsTexture)
	b.Store(outputs[0])
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlIsTexture_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlBindTexture{}) // interface compliance check
func (ω *GlBindTexture) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Target))
	if key, remap := ω.In.Texture.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Texture.value(b, ω, s))
	} else {
		b.Push(ω.In.Texture.value(b, ω, s))
	}
	b.CallNoPush(funcInfoGlBindTexture)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlTexImage2D{}) // interface compliance check
func (ω *GlTexImage2D) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Target))
	b.Push(value.S32(ω.In.Level))
	b.Push(value.U32(ω.In.InternalFormat))
	b.Push(value.S32(ω.In.Width))
	b.Push(value.S32(ω.In.Height))
	b.Push(value.S32(ω.In.Border))
	b.Push(value.U32(ω.In.Format))
	b.Push(value.U32(ω.In.Type))
	b.Push(ω.In.Data.value(b, ω, s))
	b.CallNoPush(funcInfoGlTexImage2D)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlTexSubImage2D{}) // interface compliance check
func (ω *GlTexSubImage2D) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Target))
	b.Push(value.S32(ω.In.Level))
	b.Push(value.S32(ω.In.Xoffset))
	b.Push(value.S32(ω.In.Yoffset))
	b.Push(value.S32(ω.In.Width))
	b.Push(value.S32(ω.In.Height))
	b.Push(value.U32(ω.In.Format))
	b.Push(value.U32(ω.In.Type))
	b.Push(ω.In.Data.value(b, ω, s))
	b.CallNoPush(funcInfoGlTexSubImage2D)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlCopyTexImage2D{}) // interface compliance check
func (ω *GlCopyTexImage2D) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Target))
	b.Push(value.S32(ω.In.Level))
	b.Push(value.U32(ω.In.Format))
	b.Push(value.S32(ω.In.X))
	b.Push(value.S32(ω.In.Y))
	b.Push(value.S32(ω.In.Width))
	b.Push(value.S32(ω.In.Height))
	b.Push(value.S32(ω.In.Border))
	b.CallNoPush(funcInfoGlCopyTexImage2D)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlCopyTexSubImage2D{}) // interface compliance check
func (ω *GlCopyTexSubImage2D) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Target))
	b.Push(value.S32(ω.In.Level))
	b.Push(value.S32(ω.In.Xoffset))
	b.Push(value.S32(ω.In.Yoffset))
	b.Push(value.S32(ω.In.X))
	b.Push(value.S32(ω.In.Y))
	b.Push(value.S32(ω.In.Width))
	b.Push(value.S32(ω.In.Height))
	b.CallNoPush(funcInfoGlCopyTexSubImage2D)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlCompressedTexImage2D{}) // interface compliance check
func (ω *GlCompressedTexImage2D) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Target))
	b.Push(value.S32(ω.In.Level))
	b.Push(value.U32(ω.In.Format))
	b.Push(value.S32(ω.In.Width))
	b.Push(value.S32(ω.In.Height))
	b.Push(value.S32(ω.In.Border))
	b.Push(value.S32(ω.In.ImageSize))
	b.Push(ω.In.Data.value(b, ω, s))
	b.CallNoPush(funcInfoGlCompressedTexImage2D)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlCompressedTexSubImage2D{}) // interface compliance check
func (ω *GlCompressedTexSubImage2D) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Target))
	b.Push(value.S32(ω.In.Level))
	b.Push(value.S32(ω.In.Xoffset))
	b.Push(value.S32(ω.In.Yoffset))
	b.Push(value.S32(ω.In.Width))
	b.Push(value.S32(ω.In.Height))
	b.Push(value.U32(ω.In.Format))
	b.Push(value.S32(ω.In.ImageSize))
	b.Push(ω.In.Data.value(b, ω, s))
	b.CallNoPush(funcInfoGlCompressedTexSubImage2D)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlGenerateMipmap{}) // interface compliance check
func (ω *GlGenerateMipmap) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Target))
	b.CallNoPush(funcInfoGlGenerateMipmap)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlReadPixels{}) // interface compliance check
func (ω *GlReadPixels) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	data_cnt := uint64(imageSize(ω.In.Width, ω.In.Height, TexelFormat(ω.In.Format), ω.In.Type))
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{data_cnt /* data */})
	b.Push(value.S32(ω.In.X))
	b.Push(value.S32(ω.In.Y))
	b.Push(value.S32(ω.In.Width))
	b.Push(value.S32(ω.In.Height))
	b.Push(value.U32(ω.In.Format))
	b.Push(value.U32(ω.In.Type))
	b.Push(outputs[0]) // data
	b.CallNoPush(funcInfoGlReadPixels)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlReadPixels_Postback{}
			if err := postback.Decode(data_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlGenFramebuffers{}) // interface compliance check
func (ω *GlGenFramebuffers) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	framebuffers_cnt := uint64(ω.In.Count)
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{framebuffers_cnt * 4 /* framebuffers */})
	b.Push(value.S32(ω.In.Count))
	b.Push(outputs[0]) // framebuffers
	b.CallNoPush(funcInfoGlGenFramebuffers)
	StateMutator{State: s}.Write(id, ω)
	for i, e := range ω.Out.Framebuffers {
		ptr := outputs[0].Offset(uint64(i * 4))
		if key, remap := e.remap(ω, s); remap {
			storeRemap(b, key, ptr, protocol.TypeUint32)
		}
	}
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGenFramebuffers_Postback{}
			if err := postback.Decode(framebuffers_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlBindFramebuffer{}) // interface compliance check
func (ω *GlBindFramebuffer) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Target))
	if key, remap := ω.In.Framebuffer.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Framebuffer.value(b, ω, s))
	} else {
		b.Push(ω.In.Framebuffer.value(b, ω, s))
	}
	b.CallNoPush(funcInfoGlBindFramebuffer)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlCheckFramebufferStatus{}) // interface compliance check
func (ω *GlCheckFramebufferStatus) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{4 /* result */})
	b.Push(value.U32(ω.In.Target))
	b.CallPush(funcInfoGlCheckFramebufferStatus)
	b.Store(outputs[0])
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlCheckFramebufferStatus_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlDeleteFramebuffers{}) // interface compliance check
func (ω *GlDeleteFramebuffers) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.S32(ω.In.Count))
	b.Push(ω.In.Framebuffers.value(b, ω, s))
	b.CallNoPush(funcInfoGlDeleteFramebuffers)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlIsFramebuffer{}) // interface compliance check
func (ω *GlIsFramebuffer) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{1 /* result */})
	if key, remap := ω.In.Framebuffer.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Framebuffer.value(b, ω, s))
	} else {
		b.Push(ω.In.Framebuffer.value(b, ω, s))
	}
	b.CallPush(funcInfoGlIsFramebuffer)
	b.Store(outputs[0])
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlIsFramebuffer_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlGenRenderbuffers{}) // interface compliance check
func (ω *GlGenRenderbuffers) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	renderbuffers_cnt := uint64(ω.In.Count)
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{renderbuffers_cnt * 4 /* renderbuffers */})
	b.Push(value.S32(ω.In.Count))
	b.Push(outputs[0]) // renderbuffers
	b.CallNoPush(funcInfoGlGenRenderbuffers)
	StateMutator{State: s}.Write(id, ω)
	for i, e := range ω.Out.Renderbuffers {
		ptr := outputs[0].Offset(uint64(i * 4))
		if key, remap := e.remap(ω, s); remap {
			storeRemap(b, key, ptr, protocol.TypeUint32)
		}
	}
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGenRenderbuffers_Postback{}
			if err := postback.Decode(renderbuffers_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlBindRenderbuffer{}) // interface compliance check
func (ω *GlBindRenderbuffer) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Target))
	if key, remap := ω.In.Renderbuffer.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Renderbuffer.value(b, ω, s))
	} else {
		b.Push(ω.In.Renderbuffer.value(b, ω, s))
	}
	b.CallNoPush(funcInfoGlBindRenderbuffer)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlRenderbufferStorage{}) // interface compliance check
func (ω *GlRenderbufferStorage) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Target))
	b.Push(value.U32(ω.In.Format))
	b.Push(value.S32(ω.In.Width))
	b.Push(value.S32(ω.In.Height))
	b.CallNoPush(funcInfoGlRenderbufferStorage)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlDeleteRenderbuffers{}) // interface compliance check
func (ω *GlDeleteRenderbuffers) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.S32(ω.In.Count))
	b.Push(ω.In.Renderbuffers.value(b, ω, s))
	b.CallNoPush(funcInfoGlDeleteRenderbuffers)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlIsRenderbuffer{}) // interface compliance check
func (ω *GlIsRenderbuffer) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{1 /* result */})
	if key, remap := ω.In.Renderbuffer.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Renderbuffer.value(b, ω, s))
	} else {
		b.Push(ω.In.Renderbuffer.value(b, ω, s))
	}
	b.CallPush(funcInfoGlIsRenderbuffer)
	b.Store(outputs[0])
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlIsRenderbuffer_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlGetRenderbufferParameteriv{}) // interface compliance check
func (ω *GlGetRenderbufferParameteriv) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	values_cnt := uint64(1)
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{values_cnt * 4 /* values */})
	b.Push(value.U32(ω.In.Target))
	b.Push(value.U32(ω.In.Parameter))
	b.Push(outputs[0]) // values
	b.CallNoPush(funcInfoGlGetRenderbufferParameteriv)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGetRenderbufferParameteriv_Postback{}
			if err := postback.Decode(values_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlGenBuffers{}) // interface compliance check
func (ω *GlGenBuffers) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	buffers_cnt := uint64(ω.In.Count)
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{buffers_cnt * 4 /* buffers */})
	b.Push(value.S32(ω.In.Count))
	b.Push(outputs[0]) // buffers
	b.CallNoPush(funcInfoGlGenBuffers)
	StateMutator{State: s}.Write(id, ω)
	for i, e := range ω.Out.Buffers {
		ptr := outputs[0].Offset(uint64(i * 4))
		if key, remap := e.remap(ω, s); remap {
			storeRemap(b, key, ptr, protocol.TypeUint32)
		}
	}
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGenBuffers_Postback{}
			if err := postback.Decode(buffers_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlBindBuffer{}) // interface compliance check
func (ω *GlBindBuffer) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Target))
	if key, remap := ω.In.Buffer.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Buffer.value(b, ω, s))
	} else {
		b.Push(ω.In.Buffer.value(b, ω, s))
	}
	b.CallNoPush(funcInfoGlBindBuffer)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlBufferData{}) // interface compliance check
func (ω *GlBufferData) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Target))
	b.Push(value.S32(ω.In.Size))
	b.Push(ω.In.Data.value(b, ω, s))
	b.Push(value.U32(ω.In.Usage))
	b.CallNoPush(funcInfoGlBufferData)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlBufferSubData{}) // interface compliance check
func (ω *GlBufferSubData) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Target))
	b.Push(value.S32(ω.In.Offset))
	b.Push(value.S32(ω.In.Size))
	b.Push(value.VolatileCapturePointer(uint64(ω.In.Data)))
	b.CallNoPush(funcInfoGlBufferSubData)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlDeleteBuffers{}) // interface compliance check
func (ω *GlDeleteBuffers) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.S32(ω.In.Count))
	b.Push(ω.In.Buffers.value(b, ω, s))
	b.CallNoPush(funcInfoGlDeleteBuffers)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlIsBuffer{}) // interface compliance check
func (ω *GlIsBuffer) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{1 /* result */})
	if key, remap := ω.In.Buffer.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Buffer.value(b, ω, s))
	} else {
		b.Push(ω.In.Buffer.value(b, ω, s))
	}
	b.CallPush(funcInfoGlIsBuffer)
	b.Store(outputs[0])
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlIsBuffer_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlGetBufferParameteriv{}) // interface compliance check
func (ω *GlGetBufferParameteriv) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{4 /* value */})
	b.Push(value.U32(ω.In.Target))
	b.Push(value.U32(ω.In.Parameter))
	b.Push(outputs[0]) // value
	b.CallNoPush(funcInfoGlGetBufferParameteriv)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGetBufferParameteriv_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlCreateShader{}) // interface compliance check
func (ω *GlCreateShader) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{4 /* result */})
	b.Push(value.U32(ω.In.Type))
	b.CallPush(funcInfoGlCreateShader)
	b.Store(outputs[0])
	StateMutator{State: s}.Write(id, ω)
	if key, remap := ω.Out.Result.remap(ω, s); remap {
		storeRemap(b, key, outputs[0], protocol.TypeUint32)
	}
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlCreateShader_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlDeleteShader{}) // interface compliance check
func (ω *GlDeleteShader) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Shader.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Shader.value(b, ω, s))
	} else {
		b.Push(ω.In.Shader.value(b, ω, s))
	}
	b.CallNoPush(funcInfoGlDeleteShader)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlShaderSource{}) // interface compliance check
func (ω *GlShaderSource) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Shader.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Shader.value(b, ω, s))
	} else {
		b.Push(ω.In.Shader.value(b, ω, s))
	}
	b.Push(value.S32(ω.In.Count))
	b.Push(ω.In.Source.value(b, ω, s))
	b.Push(ω.In.Length.value(b, ω, s))
	b.CallNoPush(funcInfoGlShaderSource)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlShaderBinary{}) // interface compliance check
func (ω *GlShaderBinary) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.S32(ω.In.Count))
	b.Push(ω.In.Shaders.value(b, ω, s))
	b.Push(value.U32(ω.In.BinaryFormat))
	b.Push(value.VolatileCapturePointer(uint64(ω.In.Binary)))
	b.Push(value.S32(ω.In.BinarySize))
	b.CallNoPush(funcInfoGlShaderBinary)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlGetShaderInfoLog{}) // interface compliance check
func (ω *GlGetShaderInfoLog) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	info_cnt := uint64(ω.In.BufferLength)
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{4 /* string_length_written */, info_cnt /* info */})
	if key, remap := ω.In.Shader.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Shader.value(b, ω, s))
	} else {
		b.Push(ω.In.Shader.value(b, ω, s))
	}
	b.Push(value.S32(ω.In.BufferLength))
	b.Push(outputs[0]) // string_length_written
	b.Push(outputs[1]) // info
	b.CallNoPush(funcInfoGlGetShaderInfoLog)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGetShaderInfoLog_Postback{}
			if err := postback.Decode(info_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlGetShaderSource{}) // interface compliance check
func (ω *GlGetShaderSource) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	source_cnt := uint64(ω.In.BufferLength)
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{4 /* string_length_written */, source_cnt /* source */})
	if key, remap := ω.In.Shader.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Shader.value(b, ω, s))
	} else {
		b.Push(ω.In.Shader.value(b, ω, s))
	}
	b.Push(value.S32(ω.In.BufferLength))
	b.Push(outputs[0]) // string_length_written
	b.Push(outputs[1]) // source
	b.CallNoPush(funcInfoGlGetShaderSource)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGetShaderSource_Postback{}
			if err := postback.Decode(source_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlReleaseShaderCompiler{}) // interface compliance check
func (ω *GlReleaseShaderCompiler) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.CallNoPush(funcInfoGlReleaseShaderCompiler)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlCompileShader{}) // interface compliance check
func (ω *GlCompileShader) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Shader.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Shader.value(b, ω, s))
	} else {
		b.Push(ω.In.Shader.value(b, ω, s))
	}
	b.CallNoPush(funcInfoGlCompileShader)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlIsShader{}) // interface compliance check
func (ω *GlIsShader) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{1 /* result */})
	if key, remap := ω.In.Shader.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Shader.value(b, ω, s))
	} else {
		b.Push(ω.In.Shader.value(b, ω, s))
	}
	b.CallPush(funcInfoGlIsShader)
	b.Store(outputs[0])
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlIsShader_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlCreateProgram{}) // interface compliance check
func (ω *GlCreateProgram) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{4 /* result */})
	b.CallPush(funcInfoGlCreateProgram)
	b.Store(outputs[0])
	StateMutator{State: s}.Write(id, ω)
	if key, remap := ω.Out.Result.remap(ω, s); remap {
		storeRemap(b, key, outputs[0], protocol.TypeUint32)
	}
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlCreateProgram_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlDeleteProgram{}) // interface compliance check
func (ω *GlDeleteProgram) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Program.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Program.value(b, ω, s))
	} else {
		b.Push(ω.In.Program.value(b, ω, s))
	}
	b.CallNoPush(funcInfoGlDeleteProgram)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlAttachShader{}) // interface compliance check
func (ω *GlAttachShader) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Program.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Program.value(b, ω, s))
	} else {
		b.Push(ω.In.Program.value(b, ω, s))
	}
	if key, remap := ω.In.Shader.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Shader.value(b, ω, s))
	} else {
		b.Push(ω.In.Shader.value(b, ω, s))
	}
	b.CallNoPush(funcInfoGlAttachShader)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlDetachShader{}) // interface compliance check
func (ω *GlDetachShader) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Program.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Program.value(b, ω, s))
	} else {
		b.Push(ω.In.Program.value(b, ω, s))
	}
	if key, remap := ω.In.Shader.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Shader.value(b, ω, s))
	} else {
		b.Push(ω.In.Shader.value(b, ω, s))
	}
	b.CallNoPush(funcInfoGlDetachShader)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlGetAttachedShaders{}) // interface compliance check
func (ω *GlGetAttachedShaders) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	shaders_cnt := uint64(ω.In.BufferLength)
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{4 /* shaders_length_written */, shaders_cnt * 4 /* shaders */})
	if key, remap := ω.In.Program.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Program.value(b, ω, s))
	} else {
		b.Push(ω.In.Program.value(b, ω, s))
	}
	b.Push(value.S32(ω.In.BufferLength))
	b.Push(outputs[0]) // shaders_length_written
	b.Push(outputs[1]) // shaders
	b.CallNoPush(funcInfoGlGetAttachedShaders)
	StateMutator{State: s}.Write(id, ω)
	for i, e := range ω.Out.Shaders {
		ptr := outputs[1].Offset(uint64(i * 4))
		if key, remap := e.remap(ω, s); remap {
			storeRemap(b, key, ptr, protocol.TypeUint32)
		}
	}
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGetAttachedShaders_Postback{}
			if err := postback.Decode(shaders_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlLinkProgram{}) // interface compliance check
func (ω *GlLinkProgram) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Program.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Program.value(b, ω, s))
	} else {
		b.Push(ω.In.Program.value(b, ω, s))
	}
	b.CallNoPush(funcInfoGlLinkProgram)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlGetProgramInfoLog{}) // interface compliance check
func (ω *GlGetProgramInfoLog) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	info_cnt := uint64(ω.In.BufferLength)
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{4 /* string_length_written */, info_cnt /* info */})
	if key, remap := ω.In.Program.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Program.value(b, ω, s))
	} else {
		b.Push(ω.In.Program.value(b, ω, s))
	}
	b.Push(value.S32(ω.In.BufferLength))
	b.Push(outputs[0]) // string_length_written
	b.Push(outputs[1]) // info
	b.CallNoPush(funcInfoGlGetProgramInfoLog)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGetProgramInfoLog_Postback{}
			if err := postback.Decode(info_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlUseProgram{}) // interface compliance check
func (ω *GlUseProgram) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Program.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Program.value(b, ω, s))
	} else {
		b.Push(ω.In.Program.value(b, ω, s))
	}
	b.CallNoPush(funcInfoGlUseProgram)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlIsProgram{}) // interface compliance check
func (ω *GlIsProgram) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{1 /* result */})
	if key, remap := ω.In.Program.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Program.value(b, ω, s))
	} else {
		b.Push(ω.In.Program.value(b, ω, s))
	}
	b.CallPush(funcInfoGlIsProgram)
	b.Store(outputs[0])
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlIsProgram_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlValidateProgram{}) // interface compliance check
func (ω *GlValidateProgram) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Program.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Program.value(b, ω, s))
	} else {
		b.Push(ω.In.Program.value(b, ω, s))
	}
	b.CallNoPush(funcInfoGlValidateProgram)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlClearColor{}) // interface compliance check
func (ω *GlClearColor) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.F32(ω.In.R))
	b.Push(value.F32(ω.In.G))
	b.Push(value.F32(ω.In.B))
	b.Push(value.F32(ω.In.A))
	b.CallNoPush(funcInfoGlClearColor)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlClearDepthf{}) // interface compliance check
func (ω *GlClearDepthf) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.F32(ω.In.Depth))
	b.CallNoPush(funcInfoGlClearDepthf)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlClearStencil{}) // interface compliance check
func (ω *GlClearStencil) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.S32(ω.In.Stencil))
	b.CallNoPush(funcInfoGlClearStencil)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlClear{}) // interface compliance check
func (ω *GlClear) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Mask))
	b.CallNoPush(funcInfoGlClear)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlCullFace{}) // interface compliance check
func (ω *GlCullFace) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Mode))
	b.CallNoPush(funcInfoGlCullFace)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlPolygonOffset{}) // interface compliance check
func (ω *GlPolygonOffset) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.F32(ω.In.ScaleFactor))
	b.Push(value.F32(ω.In.Units))
	b.CallNoPush(funcInfoGlPolygonOffset)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlLineWidth{}) // interface compliance check
func (ω *GlLineWidth) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.F32(ω.In.Width))
	b.CallNoPush(funcInfoGlLineWidth)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlSampleCoverage{}) // interface compliance check
func (ω *GlSampleCoverage) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.F32(ω.In.Value))
	b.Push(value.Bool(ω.In.Invert))
	b.CallNoPush(funcInfoGlSampleCoverage)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlHint{}) // interface compliance check
func (ω *GlHint) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Target))
	b.Push(value.U32(ω.In.Mode))
	b.CallNoPush(funcInfoGlHint)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlFramebufferRenderbuffer{}) // interface compliance check
func (ω *GlFramebufferRenderbuffer) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.FramebufferTarget))
	b.Push(value.U32(ω.In.FramebufferAttachment))
	b.Push(value.U32(ω.In.RenderbufferTarget))
	if key, remap := ω.In.Renderbuffer.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Renderbuffer.value(b, ω, s))
	} else {
		b.Push(ω.In.Renderbuffer.value(b, ω, s))
	}
	b.CallNoPush(funcInfoGlFramebufferRenderbuffer)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlFramebufferTexture2D{}) // interface compliance check
func (ω *GlFramebufferTexture2D) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.FramebufferTarget))
	b.Push(value.U32(ω.In.FramebufferAttachment))
	b.Push(value.U32(ω.In.TextureTarget))
	if key, remap := ω.In.Texture.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Texture.value(b, ω, s))
	} else {
		b.Push(ω.In.Texture.value(b, ω, s))
	}
	b.Push(value.S32(ω.In.Level))
	b.CallNoPush(funcInfoGlFramebufferTexture2D)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlGetFramebufferAttachmentParameteriv{}) // interface compliance check
func (ω *GlGetFramebufferAttachmentParameteriv) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	value_cnt := uint64(1)
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{value_cnt * 4 /* value */})
	b.Push(value.U32(ω.In.Target))
	b.Push(value.U32(ω.In.Attachment))
	b.Push(value.U32(ω.In.Parameter))
	b.Push(outputs[0]) // value
	b.CallNoPush(funcInfoGlGetFramebufferAttachmentParameteriv)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGetFramebufferAttachmentParameteriv_Postback{}
			if err := postback.Decode(value_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlDrawElements{}) // interface compliance check
func (ω *GlDrawElements) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.DrawMode))
	b.Push(value.S32(ω.In.ElementCount))
	b.Push(value.U32(ω.In.IndicesType))
	b.Push(ω.In.Indices.value(b, ω, s))
	b.CallNoPush(funcInfoGlDrawElements)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlDrawArrays{}) // interface compliance check
func (ω *GlDrawArrays) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.DrawMode))
	b.Push(value.S32(ω.In.FirstIndex))
	b.Push(value.S32(ω.In.IndexCount))
	b.CallNoPush(funcInfoGlDrawArrays)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlFlush{}) // interface compliance check
func (ω *GlFlush) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.CallNoPush(funcInfoGlFlush)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlFinish{}) // interface compliance check
func (ω *GlFinish) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.CallNoPush(funcInfoGlFinish)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlGetBooleanv{}) // interface compliance check
func (ω *GlGetBooleanv) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	values_cnt := uint64(stateVariableSize(ω.In.Param))
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{values_cnt * 1 /* values */})
	b.Push(value.U32(ω.In.Param))
	b.Push(outputs[0]) // values
	b.CallNoPush(funcInfoGlGetBooleanv)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGetBooleanv_Postback{}
			if err := postback.Decode(values_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlGetFloatv{}) // interface compliance check
func (ω *GlGetFloatv) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	values_cnt := uint64(stateVariableSize(ω.In.Param))
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{values_cnt * 4 /* values */})
	b.Push(value.U32(ω.In.Param))
	b.Push(outputs[0]) // values
	b.CallNoPush(funcInfoGlGetFloatv)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGetFloatv_Postback{}
			if err := postback.Decode(values_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlGetIntegerv{}) // interface compliance check
func (ω *GlGetIntegerv) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	values_cnt := uint64(stateVariableSize(ω.In.Param))
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{values_cnt * 4 /* values */})
	b.Push(value.U32(ω.In.Param))
	b.Push(outputs[0]) // values
	b.CallNoPush(funcInfoGlGetIntegerv)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGetIntegerv_Postback{}
			if err := postback.Decode(values_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlGetString{}) // interface compliance check
func (ω *GlGetString) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	result_cnt := uint64(256)
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{result_cnt /* result */})
	b.Push(value.U32(ω.In.Param))
	b.CallPush(funcInfoGlGetString)
	b.Push(outputs[0])
	b.Strcpy(result_cnt)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGetString_Postback{}
			if err := postback.Decode(result_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlEnable{}) // interface compliance check
func (ω *GlEnable) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Capability))
	b.CallNoPush(funcInfoGlEnable)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlDisable{}) // interface compliance check
func (ω *GlDisable) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Capability))
	b.CallNoPush(funcInfoGlDisable)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlIsEnabled{}) // interface compliance check
func (ω *GlIsEnabled) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{1 /* result */})
	b.Push(value.U32(ω.In.Capability))
	b.CallPush(funcInfoGlIsEnabled)
	b.Store(outputs[0])
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlIsEnabled_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlMapBufferRange{}) // interface compliance check
func (ω *GlMapBufferRange) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	result_cnt := uint64(ω.In.Length)
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{result_cnt /* result */})
	b.Push(value.U32(ω.In.Target))
	b.Push(value.S32(ω.In.Offset))
	b.Push(value.S32(ω.In.Length))
	b.Push(value.U32(ω.In.Access))
	b.CallPush(funcInfoGlMapBufferRange)
	b.Push(outputs[0])
	b.Copy(result_cnt)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlMapBufferRange_Postback{}
			if err := postback.Decode(result_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlUnmapBuffer{}) // interface compliance check
func (ω *GlUnmapBuffer) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Target))
	b.CallNoPush(funcInfoGlUnmapBuffer)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlInvalidateFramebuffer{}) // interface compliance check
func (ω *GlInvalidateFramebuffer) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Target))
	b.Push(value.S32(ω.In.Count))
	b.Push(ω.In.Attachments.value(b, ω, s))
	b.CallNoPush(funcInfoGlInvalidateFramebuffer)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlRenderbufferStorageMultisample{}) // interface compliance check
func (ω *GlRenderbufferStorageMultisample) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Target))
	b.Push(value.S32(ω.In.Samples))
	b.Push(value.U32(ω.In.Format))
	b.Push(value.S32(ω.In.Width))
	b.Push(value.S32(ω.In.Height))
	b.CallNoPush(funcInfoGlRenderbufferStorageMultisample)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlBlitFramebuffer{}) // interface compliance check
func (ω *GlBlitFramebuffer) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.S32(ω.In.SrcX0))
	b.Push(value.S32(ω.In.SrcY0))
	b.Push(value.S32(ω.In.SrcX1))
	b.Push(value.S32(ω.In.SrcY1))
	b.Push(value.S32(ω.In.DstX0))
	b.Push(value.S32(ω.In.DstY0))
	b.Push(value.S32(ω.In.DstX1))
	b.Push(value.S32(ω.In.DstY1))
	b.Push(value.U32(ω.In.Mask))
	b.Push(value.U32(ω.In.Filter))
	b.CallNoPush(funcInfoGlBlitFramebuffer)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlGenQueries{}) // interface compliance check
func (ω *GlGenQueries) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	queries_cnt := uint64(ω.In.Count)
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{queries_cnt * 4 /* queries */})
	b.Push(value.S32(ω.In.Count))
	b.Push(outputs[0]) // queries
	b.CallNoPush(funcInfoGlGenQueries)
	StateMutator{State: s}.Write(id, ω)
	for i, e := range ω.Out.Queries {
		ptr := outputs[0].Offset(uint64(i * 4))
		if key, remap := e.remap(ω, s); remap {
			storeRemap(b, key, ptr, protocol.TypeUint32)
		}
	}
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGenQueries_Postback{}
			if err := postback.Decode(queries_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlBeginQuery{}) // interface compliance check
func (ω *GlBeginQuery) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Target))
	if key, remap := ω.In.Query.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Query.value(b, ω, s))
	} else {
		b.Push(ω.In.Query.value(b, ω, s))
	}
	b.CallNoPush(funcInfoGlBeginQuery)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlEndQuery{}) // interface compliance check
func (ω *GlEndQuery) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Target))
	b.CallNoPush(funcInfoGlEndQuery)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlDeleteQueries{}) // interface compliance check
func (ω *GlDeleteQueries) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.S32(ω.In.Count))
	b.Push(ω.In.Queries.value(b, ω, s))
	b.CallNoPush(funcInfoGlDeleteQueries)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlIsQuery{}) // interface compliance check
func (ω *GlIsQuery) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{1 /* result */})
	if key, remap := ω.In.Query.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Query.value(b, ω, s))
	} else {
		b.Push(ω.In.Query.value(b, ω, s))
	}
	b.CallPush(funcInfoGlIsQuery)
	b.Store(outputs[0])
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlIsQuery_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlGetQueryiv{}) // interface compliance check
func (ω *GlGetQueryiv) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{4 /* value */})
	b.Push(value.U32(ω.In.Target))
	b.Push(value.U32(ω.In.Parameter))
	b.Push(outputs[0]) // value
	b.CallNoPush(funcInfoGlGetQueryiv)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGetQueryiv_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlGetQueryObjectuiv{}) // interface compliance check
func (ω *GlGetQueryObjectuiv) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{4 /* value */})
	if key, remap := ω.In.Query.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Query.value(b, ω, s))
	} else {
		b.Push(ω.In.Query.value(b, ω, s))
	}
	b.Push(value.U32(ω.In.Parameter))
	b.Push(outputs[0]) // value
	b.CallNoPush(funcInfoGlGetQueryObjectuiv)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGetQueryObjectuiv_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlGenQueriesEXT{}) // interface compliance check
func (ω *GlGenQueriesEXT) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	queries_cnt := uint64(ω.In.Count)
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{queries_cnt * 4 /* queries */})
	b.Push(value.S32(ω.In.Count))
	b.Push(outputs[0]) // queries
	b.CallNoPush(funcInfoGlGenQueriesEXT)
	StateMutator{State: s}.Write(id, ω)
	for i, e := range ω.Out.Queries {
		ptr := outputs[0].Offset(uint64(i * 4))
		if key, remap := e.remap(ω, s); remap {
			storeRemap(b, key, ptr, protocol.TypeUint32)
		}
	}
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGenQueriesEXT_Postback{}
			if err := postback.Decode(queries_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlBeginQueryEXT{}) // interface compliance check
func (ω *GlBeginQueryEXT) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Target))
	if key, remap := ω.In.Query.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Query.value(b, ω, s))
	} else {
		b.Push(ω.In.Query.value(b, ω, s))
	}
	b.CallNoPush(funcInfoGlBeginQueryEXT)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlEndQueryEXT{}) // interface compliance check
func (ω *GlEndQueryEXT) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.Target))
	b.CallNoPush(funcInfoGlEndQueryEXT)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlDeleteQueriesEXT{}) // interface compliance check
func (ω *GlDeleteQueriesEXT) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.S32(ω.In.Count))
	b.Push(ω.In.Queries.value(b, ω, s))
	b.CallNoPush(funcInfoGlDeleteQueriesEXT)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlIsQueryEXT{}) // interface compliance check
func (ω *GlIsQueryEXT) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{1 /* result */})
	if key, remap := ω.In.Query.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Query.value(b, ω, s))
	} else {
		b.Push(ω.In.Query.value(b, ω, s))
	}
	b.CallPush(funcInfoGlIsQueryEXT)
	b.Store(outputs[0])
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlIsQueryEXT_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlQueryCounterEXT{}) // interface compliance check
func (ω *GlQueryCounterEXT) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.Query.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Query.value(b, ω, s))
	} else {
		b.Push(ω.In.Query.value(b, ω, s))
	}
	b.Push(value.U32(ω.In.Target))
	b.CallNoPush(funcInfoGlQueryCounterEXT)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&GlGetQueryivEXT{}) // interface compliance check
func (ω *GlGetQueryivEXT) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{4 /* value */})
	b.Push(value.U32(ω.In.Target))
	b.Push(value.U32(ω.In.Parameter))
	b.Push(outputs[0]) // value
	b.CallNoPush(funcInfoGlGetQueryivEXT)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGetQueryivEXT_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlGetQueryObjectivEXT{}) // interface compliance check
func (ω *GlGetQueryObjectivEXT) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{4 /* value */})
	if key, remap := ω.In.Query.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Query.value(b, ω, s))
	} else {
		b.Push(ω.In.Query.value(b, ω, s))
	}
	b.Push(value.U32(ω.In.Parameter))
	b.Push(outputs[0]) // value
	b.CallNoPush(funcInfoGlGetQueryObjectivEXT)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGetQueryObjectivEXT_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlGetQueryObjectuivEXT{}) // interface compliance check
func (ω *GlGetQueryObjectuivEXT) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{4 /* value */})
	if key, remap := ω.In.Query.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Query.value(b, ω, s))
	} else {
		b.Push(ω.In.Query.value(b, ω, s))
	}
	b.Push(value.U32(ω.In.Parameter))
	b.Push(outputs[0]) // value
	b.CallNoPush(funcInfoGlGetQueryObjectuivEXT)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGetQueryObjectuivEXT_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlGetQueryObjecti64vEXT{}) // interface compliance check
func (ω *GlGetQueryObjecti64vEXT) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{8 /* value */})
	if key, remap := ω.In.Query.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Query.value(b, ω, s))
	} else {
		b.Push(ω.In.Query.value(b, ω, s))
	}
	b.Push(value.U32(ω.In.Parameter))
	b.Push(outputs[0]) // value
	b.CallNoPush(funcInfoGlGetQueryObjecti64vEXT)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGetQueryObjecti64vEXT_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&GlGetQueryObjectui64vEXT{}) // interface compliance check
func (ω *GlGetQueryObjectui64vEXT) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{8 /* value */})
	if key, remap := ω.In.Query.remap(ω, s); remap {
		loadRemap(b, key, ω.In.Query.value(b, ω, s))
	} else {
		b.Push(ω.In.Query.value(b, ω, s))
	}
	b.Push(value.U32(ω.In.Parameter))
	b.Push(outputs[0]) // value
	b.CallNoPush(funcInfoGlGetQueryObjectui64vEXT)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := GlGetQueryObjectui64vEXT_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}
