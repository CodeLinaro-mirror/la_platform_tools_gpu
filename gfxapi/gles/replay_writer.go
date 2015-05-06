////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

package gles

import (
	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/replay/builder"
	"android.googlesource.com/platform/tools/gpu/replay/protocol"
	"android.googlesource.com/platform/tools/gpu/replay/value"
)

func readBytes(r binary.Reader, c uint64) ([]byte, error) {
	b := make([]byte, c)
	err := r.Data(b)
	return b, err
}
func readString(r binary.Reader, c uint64) (string, error) {
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

var funcInfoReplayCreateRenderer = builder.FunctionInfo{ID: 0, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoReplayBindRenderer = builder.FunctionInfo{ID: 1, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoBackbufferInfo = builder.FunctionInfo{ID: 2, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoStartTimer = builder.FunctionInfo{ID: 3, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoStopTimer = builder.FunctionInfo{ID: 4, ReturnType: protocol.TypeUint64, Parameters: 1}
var funcInfoFlushPostBuffer = builder.FunctionInfo{ID: 5, ReturnType: protocol.TypeVoid, Parameters: 0}
var funcInfoEglInitialize = builder.FunctionInfo{ID: 6, ReturnType: protocol.TypeInt32, Parameters: 3}
var funcInfoEglCreateContext = builder.FunctionInfo{ID: 7, ReturnType: protocol.TypeAbsolutePointer, Parameters: 4}
var funcInfoEglMakeCurrent = builder.FunctionInfo{ID: 8, ReturnType: protocol.TypeInt32, Parameters: 4}
var funcInfoEglSwapBuffers = builder.FunctionInfo{ID: 9, ReturnType: protocol.TypeInt32, Parameters: 2}
var funcInfoEglQuerySurface = builder.FunctionInfo{ID: 10, ReturnType: protocol.TypeInt32, Parameters: 4}
var funcInfoGlXCreateContext = builder.FunctionInfo{ID: 11, ReturnType: protocol.TypeAbsolutePointer, Parameters: 4}
var funcInfoGlXCreateNewContext = builder.FunctionInfo{ID: 12, ReturnType: protocol.TypeAbsolutePointer, Parameters: 5}
var funcInfoGlXMakeContextCurrent = builder.FunctionInfo{ID: 13, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlXSwapBuffers = builder.FunctionInfo{ID: 14, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoWglCreateContext = builder.FunctionInfo{ID: 15, ReturnType: protocol.TypeAbsolutePointer, Parameters: 1}
var funcInfoWglCreateContextAttribsARB = builder.FunctionInfo{ID: 16, ReturnType: protocol.TypeAbsolutePointer, Parameters: 3}
var funcInfoWglMakeCurrent = builder.FunctionInfo{ID: 17, ReturnType: protocol.TypeInt32, Parameters: 2}
var funcInfoWglSwapBuffers = builder.FunctionInfo{ID: 18, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCGLCreateContext = builder.FunctionInfo{ID: 19, ReturnType: protocol.TypeInt32, Parameters: 3}
var funcInfoCGLSetCurrentContext = builder.FunctionInfo{ID: 20, ReturnType: protocol.TypeInt32, Parameters: 1}
var funcInfoGlEnableClientState = builder.FunctionInfo{ID: 21, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlDisableClientState = builder.FunctionInfo{ID: 22, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlGetProgramBinaryOES = builder.FunctionInfo{ID: 23, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlProgramBinaryOES = builder.FunctionInfo{ID: 24, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlStartTilingQCOM = builder.FunctionInfo{ID: 25, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlEndTilingQCOM = builder.FunctionInfo{ID: 26, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlDiscardFramebufferEXT = builder.FunctionInfo{ID: 27, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlInsertEventMarkerEXT = builder.FunctionInfo{ID: 28, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlPushGroupMarkerEXT = builder.FunctionInfo{ID: 29, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlPopGroupMarkerEXT = builder.FunctionInfo{ID: 30, ReturnType: protocol.TypeVoid, Parameters: 0}
var funcInfoGlTexStorage1DEXT = builder.FunctionInfo{ID: 31, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlTexStorage2DEXT = builder.FunctionInfo{ID: 32, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlTexStorage3DEXT = builder.FunctionInfo{ID: 33, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlTextureStorage1DEXT = builder.FunctionInfo{ID: 34, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlTextureStorage2DEXT = builder.FunctionInfo{ID: 35, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlTextureStorage3DEXT = builder.FunctionInfo{ID: 36, ReturnType: protocol.TypeVoid, Parameters: 7}
var funcInfoGlGenVertexArraysOES = builder.FunctionInfo{ID: 37, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBindVertexArrayOES = builder.FunctionInfo{ID: 38, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlDeleteVertexArraysOES = builder.FunctionInfo{ID: 39, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlIsVertexArrayOES = builder.FunctionInfo{ID: 40, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlEGLImageTargetTexture2DOES = builder.FunctionInfo{ID: 41, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlEGLImageTargetRenderbufferStorageOES = builder.FunctionInfo{ID: 42, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetGraphicsResetStatusEXT = builder.FunctionInfo{ID: 43, ReturnType: protocol.TypeUint32, Parameters: 0}
var funcInfoGlBindAttribLocation = builder.FunctionInfo{ID: 44, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlBlendFunc = builder.FunctionInfo{ID: 45, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBlendFuncSeparate = builder.FunctionInfo{ID: 46, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlBlendEquation = builder.FunctionInfo{ID: 47, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlBlendEquationSeparate = builder.FunctionInfo{ID: 48, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBlendColor = builder.FunctionInfo{ID: 49, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlEnableVertexAttribArray = builder.FunctionInfo{ID: 50, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlDisableVertexAttribArray = builder.FunctionInfo{ID: 51, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlVertexAttribPointer = builder.FunctionInfo{ID: 52, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlGetActiveAttrib = builder.FunctionInfo{ID: 53, ReturnType: protocol.TypeVoid, Parameters: 7}
var funcInfoGlGetActiveUniform = builder.FunctionInfo{ID: 54, ReturnType: protocol.TypeVoid, Parameters: 7}
var funcInfoGlGetError = builder.FunctionInfo{ID: 55, ReturnType: protocol.TypeUint32, Parameters: 0}
var funcInfoGlGetProgramiv = builder.FunctionInfo{ID: 56, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetShaderiv = builder.FunctionInfo{ID: 57, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetUniformLocation = builder.FunctionInfo{ID: 58, ReturnType: protocol.TypeInt32, Parameters: 2}
var funcInfoGlGetAttribLocation = builder.FunctionInfo{ID: 59, ReturnType: protocol.TypeUint32, Parameters: 2}
var funcInfoGlPixelStorei = builder.FunctionInfo{ID: 60, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlTexParameteri = builder.FunctionInfo{ID: 61, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlTexParameterf = builder.FunctionInfo{ID: 62, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetTexParameteriv = builder.FunctionInfo{ID: 63, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetTexParameterfv = builder.FunctionInfo{ID: 64, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform1i = builder.FunctionInfo{ID: 65, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlUniform2i = builder.FunctionInfo{ID: 66, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform3i = builder.FunctionInfo{ID: 67, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUniform4i = builder.FunctionInfo{ID: 68, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlUniform1iv = builder.FunctionInfo{ID: 69, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform2iv = builder.FunctionInfo{ID: 70, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform3iv = builder.FunctionInfo{ID: 71, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform4iv = builder.FunctionInfo{ID: 72, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform1f = builder.FunctionInfo{ID: 73, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlUniform2f = builder.FunctionInfo{ID: 74, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform3f = builder.FunctionInfo{ID: 75, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUniform4f = builder.FunctionInfo{ID: 76, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlUniform1fv = builder.FunctionInfo{ID: 77, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform2fv = builder.FunctionInfo{ID: 78, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform3fv = builder.FunctionInfo{ID: 79, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform4fv = builder.FunctionInfo{ID: 80, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniformMatrix2fv = builder.FunctionInfo{ID: 81, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUniformMatrix3fv = builder.FunctionInfo{ID: 82, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUniformMatrix4fv = builder.FunctionInfo{ID: 83, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlGetUniformfv = builder.FunctionInfo{ID: 84, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetUniformiv = builder.FunctionInfo{ID: 85, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlVertexAttrib1f = builder.FunctionInfo{ID: 86, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlVertexAttrib2f = builder.FunctionInfo{ID: 87, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlVertexAttrib3f = builder.FunctionInfo{ID: 88, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlVertexAttrib4f = builder.FunctionInfo{ID: 89, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlVertexAttrib1fv = builder.FunctionInfo{ID: 90, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlVertexAttrib2fv = builder.FunctionInfo{ID: 91, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlVertexAttrib3fv = builder.FunctionInfo{ID: 92, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlVertexAttrib4fv = builder.FunctionInfo{ID: 93, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetShaderPrecisionFormat = builder.FunctionInfo{ID: 94, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlDepthMask = builder.FunctionInfo{ID: 95, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlDepthFunc = builder.FunctionInfo{ID: 96, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlDepthRangef = builder.FunctionInfo{ID: 97, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlColorMask = builder.FunctionInfo{ID: 98, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlStencilMask = builder.FunctionInfo{ID: 99, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlStencilMaskSeparate = builder.FunctionInfo{ID: 100, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlStencilFuncSeparate = builder.FunctionInfo{ID: 101, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlStencilOpSeparate = builder.FunctionInfo{ID: 102, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlFrontFace = builder.FunctionInfo{ID: 103, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlViewport = builder.FunctionInfo{ID: 104, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlScissor = builder.FunctionInfo{ID: 105, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlActiveTexture = builder.FunctionInfo{ID: 106, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlGenTextures = builder.FunctionInfo{ID: 107, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlDeleteTextures = builder.FunctionInfo{ID: 108, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlIsTexture = builder.FunctionInfo{ID: 109, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlBindTexture = builder.FunctionInfo{ID: 110, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlTexImage2D = builder.FunctionInfo{ID: 111, ReturnType: protocol.TypeVoid, Parameters: 9}
var funcInfoGlTexSubImage2D = builder.FunctionInfo{ID: 112, ReturnType: protocol.TypeVoid, Parameters: 9}
var funcInfoGlCopyTexImage2D = builder.FunctionInfo{ID: 113, ReturnType: protocol.TypeVoid, Parameters: 8}
var funcInfoGlCopyTexSubImage2D = builder.FunctionInfo{ID: 114, ReturnType: protocol.TypeVoid, Parameters: 8}
var funcInfoGlCompressedTexImage2D = builder.FunctionInfo{ID: 115, ReturnType: protocol.TypeVoid, Parameters: 8}
var funcInfoGlCompressedTexSubImage2D = builder.FunctionInfo{ID: 116, ReturnType: protocol.TypeVoid, Parameters: 9}
var funcInfoGlGenerateMipmap = builder.FunctionInfo{ID: 117, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlReadPixels = builder.FunctionInfo{ID: 118, ReturnType: protocol.TypeVoid, Parameters: 7}
var funcInfoGlGenFramebuffers = builder.FunctionInfo{ID: 119, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBindFramebuffer = builder.FunctionInfo{ID: 120, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlCheckFramebufferStatus = builder.FunctionInfo{ID: 121, ReturnType: protocol.TypeUint32, Parameters: 1}
var funcInfoGlDeleteFramebuffers = builder.FunctionInfo{ID: 122, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlIsFramebuffer = builder.FunctionInfo{ID: 123, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlGenRenderbuffers = builder.FunctionInfo{ID: 124, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBindRenderbuffer = builder.FunctionInfo{ID: 125, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlRenderbufferStorage = builder.FunctionInfo{ID: 126, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlDeleteRenderbuffers = builder.FunctionInfo{ID: 127, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlIsRenderbuffer = builder.FunctionInfo{ID: 128, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlGetRenderbufferParameteriv = builder.FunctionInfo{ID: 129, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGenBuffers = builder.FunctionInfo{ID: 130, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBindBuffer = builder.FunctionInfo{ID: 131, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBufferData = builder.FunctionInfo{ID: 132, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlBufferSubData = builder.FunctionInfo{ID: 133, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlDeleteBuffers = builder.FunctionInfo{ID: 134, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlIsBuffer = builder.FunctionInfo{ID: 135, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlGetBufferParameteriv = builder.FunctionInfo{ID: 136, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlCreateShader = builder.FunctionInfo{ID: 137, ReturnType: protocol.TypeUint32, Parameters: 1}
var funcInfoGlDeleteShader = builder.FunctionInfo{ID: 138, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlShaderSource = builder.FunctionInfo{ID: 139, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlShaderBinary = builder.FunctionInfo{ID: 140, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlGetShaderInfoLog = builder.FunctionInfo{ID: 141, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlGetShaderSource = builder.FunctionInfo{ID: 142, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlReleaseShaderCompiler = builder.FunctionInfo{ID: 143, ReturnType: protocol.TypeVoid, Parameters: 0}
var funcInfoGlCompileShader = builder.FunctionInfo{ID: 144, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlIsShader = builder.FunctionInfo{ID: 145, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlCreateProgram = builder.FunctionInfo{ID: 146, ReturnType: protocol.TypeUint32, Parameters: 0}
var funcInfoGlDeleteProgram = builder.FunctionInfo{ID: 147, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlAttachShader = builder.FunctionInfo{ID: 148, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlDetachShader = builder.FunctionInfo{ID: 149, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetAttachedShaders = builder.FunctionInfo{ID: 150, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlLinkProgram = builder.FunctionInfo{ID: 151, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlGetProgramInfoLog = builder.FunctionInfo{ID: 152, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUseProgram = builder.FunctionInfo{ID: 153, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlIsProgram = builder.FunctionInfo{ID: 154, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlValidateProgram = builder.FunctionInfo{ID: 155, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlClearColor = builder.FunctionInfo{ID: 156, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlClearDepthf = builder.FunctionInfo{ID: 157, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlClearStencil = builder.FunctionInfo{ID: 158, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlClear = builder.FunctionInfo{ID: 159, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlCullFace = builder.FunctionInfo{ID: 160, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlPolygonOffset = builder.FunctionInfo{ID: 161, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlLineWidth = builder.FunctionInfo{ID: 162, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlSampleCoverage = builder.FunctionInfo{ID: 163, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlHint = builder.FunctionInfo{ID: 164, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlFramebufferRenderbuffer = builder.FunctionInfo{ID: 165, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlFramebufferTexture2D = builder.FunctionInfo{ID: 166, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlGetFramebufferAttachmentParameteriv = builder.FunctionInfo{ID: 167, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlDrawElements = builder.FunctionInfo{ID: 168, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlDrawArrays = builder.FunctionInfo{ID: 169, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlFlush = builder.FunctionInfo{ID: 170, ReturnType: protocol.TypeVoid, Parameters: 0}
var funcInfoGlFinish = builder.FunctionInfo{ID: 171, ReturnType: protocol.TypeVoid, Parameters: 0}
var funcInfoGlGetBooleanv = builder.FunctionInfo{ID: 172, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetFloatv = builder.FunctionInfo{ID: 173, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetIntegerv = builder.FunctionInfo{ID: 174, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetString = builder.FunctionInfo{ID: 175, ReturnType: protocol.TypeAbsolutePointer, Parameters: 1}
var funcInfoGlEnable = builder.FunctionInfo{ID: 176, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlDisable = builder.FunctionInfo{ID: 177, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlIsEnabled = builder.FunctionInfo{ID: 178, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlMapBufferRange = builder.FunctionInfo{ID: 179, ReturnType: protocol.TypeAbsolutePointer, Parameters: 4}
var funcInfoGlUnmapBuffer = builder.FunctionInfo{ID: 180, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlInvalidateFramebuffer = builder.FunctionInfo{ID: 181, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlRenderbufferStorageMultisample = builder.FunctionInfo{ID: 182, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlBlitFramebuffer = builder.FunctionInfo{ID: 183, ReturnType: protocol.TypeVoid, Parameters: 10}
var funcInfoGlGenQueries = builder.FunctionInfo{ID: 184, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBeginQuery = builder.FunctionInfo{ID: 185, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlEndQuery = builder.FunctionInfo{ID: 186, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlDeleteQueries = builder.FunctionInfo{ID: 187, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlIsQuery = builder.FunctionInfo{ID: 188, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlGetQueryiv = builder.FunctionInfo{ID: 189, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetQueryObjectuiv = builder.FunctionInfo{ID: 190, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGenQueriesEXT = builder.FunctionInfo{ID: 191, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBeginQueryEXT = builder.FunctionInfo{ID: 192, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlEndQueryEXT = builder.FunctionInfo{ID: 193, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlDeleteQueriesEXT = builder.FunctionInfo{ID: 194, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlIsQueryEXT = builder.FunctionInfo{ID: 195, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlQueryCounterEXT = builder.FunctionInfo{ID: 196, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetQueryivEXT = builder.FunctionInfo{ID: 197, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetQueryObjectivEXT = builder.FunctionInfo{ID: 198, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetQueryObjectuivEXT = builder.FunctionInfo{ID: 199, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetQueryObjecti64vEXT = builder.FunctionInfo{ID: 200, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetQueryObjectui64vEXT = builder.FunctionInfo{ID: 201, ReturnType: protocol.TypeVoid, Parameters: 3}

func (c RenderbufferId) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U32(uint32(c))
}
func (c TextureId) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U32(uint32(c))
}
func (c FramebufferId) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U32(uint32(c))
}
func (c BufferId) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U32(uint32(c))
}
func (c ShaderId) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U32(uint32(c))
}
func (c ProgramId) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U32(uint32(c))
}
func (c VertexArrayId) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U32(uint32(c))
}
func (c QueryId) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U32(uint32(c))
}
func (c UniformLocation) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.S32(int32(c))
}
func (c AttributeLocation) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U32(uint32(c))
}
func (c ContextID) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U32(uint32(c))
}
func (c ThreadID) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U32(uint32(c))
}
func (c EGLBoolean) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.S64(int64(c))
}
func (c EGLint) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.S64(int64(c))
}
func (c EGLConfig) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.VolatileCapturePointer(uint64(memory.Pointer(c)))
}
func (c EGLContext) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.VolatileCapturePointer(uint64(memory.Pointer(c)))
}
func (c EGLDisplay) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.VolatileCapturePointer(uint64(memory.Pointer(c)))
}
func (c EGLSurface) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.VolatileCapturePointer(uint64(memory.Pointer(c)))
}
func (c GLXContext) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.VolatileCapturePointer(uint64(memory.Pointer(c)))
}
func (c GLXDrawable) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.VolatileCapturePointer(uint64(memory.Pointer(c)))
}
func (c HGLRC) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.VolatileCapturePointer(uint64(memory.Pointer(c)))
}
func (c HDC) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.VolatileCapturePointer(uint64(memory.Pointer(c)))
}
func (c BOOL) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.S64(int64(c))
}
func (c CGLError) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.S64(int64(c))
}
func (c CGLPixelFormatObj) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.VolatileCapturePointer(uint64(memory.Pointer(c)))
}
func (c CGLContextObj) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.VolatileCapturePointer(uint64(memory.Pointer(c)))
}
func (arr BoolArray) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			ϟb.Push(value.Bool(e))
		}
		return ϟb.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr BufferIdArray) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			if key, remap := e.remap(ϟa, ϟs); remap {
				loadRemap(ϟb, key, e.value(ϟb, ϟa, ϟs))
			} else {
				ϟb.Push(e.value(ϟb, ϟa, ϟs))
			}
		}
		return ϟb.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr CharBufferArray) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			ϟb.Push(ϟb.String(e))
		}
		return ϟb.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr DiscardFramebufferAttachmentArray) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			ϟb.Push(value.U32(e))
		}
		return ϟb.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr EGLintArray) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			ϟb.Push(e.value(ϟb, ϟa, ϟs))
		}
		return ϟb.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr F32Array) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			ϟb.Push(value.F32(e))
		}
		return ϟb.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr FramebufferAttachmentArray) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			ϟb.Push(value.U32(e))
		}
		return ϟb.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr FramebufferIdArray) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			if key, remap := e.remap(ϟa, ϟs); remap {
				loadRemap(ϟb, key, e.value(ϟb, ϟa, ϟs))
			} else {
				ϟb.Push(e.value(ϟb, ϟa, ϟs))
			}
		}
		return ϟb.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr IntArray) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			ϟb.Push(value.S64(e))
		}
		return ϟb.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr QueryIdArray) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			if key, remap := e.remap(ϟa, ϟs); remap {
				loadRemap(ϟb, key, e.value(ϟb, ϟa, ϟs))
			} else {
				ϟb.Push(e.value(ϟb, ϟa, ϟs))
			}
		}
		return ϟb.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr RenderbufferIdArray) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			if key, remap := e.remap(ϟa, ϟs); remap {
				loadRemap(ϟb, key, e.value(ϟb, ϟa, ϟs))
			} else {
				ϟb.Push(e.value(ϟb, ϟa, ϟs))
			}
		}
		return ϟb.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr S32Array) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			ϟb.Push(value.S32(e))
		}
		return ϟb.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr ShaderIdArray) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			if key, remap := e.remap(ϟa, ϟs); remap {
				loadRemap(ϟb, key, e.value(ϟb, ϟa, ϟs))
			} else {
				ϟb.Push(e.value(ϟb, ϟa, ϟs))
			}
		}
		return ϟb.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr StringArray) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			ϟb.Push(ϟb.String(e))
		}
		return ϟb.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr TextureIdArray) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			if key, remap := e.remap(ϟa, ϟs); remap {
				loadRemap(ϟb, key, e.value(ϟb, ϟa, ϟs))
			} else {
				ϟb.Push(e.value(ϟb, ϟa, ϟs))
			}
		}
		return ϟb.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr VertexArrayIdArray) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			if key, remap := e.remap(ϟa, ϟs); remap {
				loadRemap(ϟb, key, e.value(ϟb, ϟa, ϟs))
			} else {
				ϟb.Push(e.value(ϟb, ϟa, ϟs))
			}
		}
		return ϟb.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}

type StopTimer_Postback struct {
	Result uint64
}

func (o *StopTimer_Postback) Decode(d binary.Decoder) error {
	if v, err := d.Uint64(); err == nil {
		o.Result = v
	} else {
		return err
	}
	return nil
}

type EglInitialize_Postback struct {
	Major  EGLint
	Minor  EGLint
	Result EGLBoolean
}

func (o *EglInitialize_Postback) Decode(d binary.Decoder) error {
	{
		var x int64
		if v, err := d.Int64(); err == nil {
			x = v
		} else {
			return err
		}
		o.Major = EGLint(x)
	}
	{
		var x int64
		if v, err := d.Int64(); err == nil {
			x = v
		} else {
			return err
		}
		o.Minor = EGLint(x)
	}
	{
		var x int64
		if v, err := d.Int64(); err == nil {
			x = v
		} else {
			return err
		}
		o.Result = EGLBoolean(x)
	}
	return nil
}

type EglCreateContext_Postback struct {
	Result []byte
}

func (o *EglCreateContext_Postback) Decode(result_cnt uint64, d binary.Decoder) error {
	{
		var x []byte
		if val, err := readBytes(d, result_cnt); err == nil {
			x = val
		} else {
			return err
		}
		o.Result = []byte(x)
	}
	return nil
}

type EglMakeCurrent_Postback struct {
	Result EGLBoolean
}

func (o *EglMakeCurrent_Postback) Decode(d binary.Decoder) error {
	{
		var x int64
		if v, err := d.Int64(); err == nil {
			x = v
		} else {
			return err
		}
		o.Result = EGLBoolean(x)
	}
	return nil
}

type EglSwapBuffers_Postback struct {
	Result EGLBoolean
}

func (o *EglSwapBuffers_Postback) Decode(d binary.Decoder) error {
	{
		var x int64
		if v, err := d.Int64(); err == nil {
			x = v
		} else {
			return err
		}
		o.Result = EGLBoolean(x)
	}
	return nil
}

type EglQuerySurface_Postback struct {
	Value  EGLint
	Result EGLBoolean
}

func (o *EglQuerySurface_Postback) Decode(d binary.Decoder) error {
	{
		var x int64
		if v, err := d.Int64(); err == nil {
			x = v
		} else {
			return err
		}
		o.Value = EGLint(x)
	}
	{
		var x int64
		if v, err := d.Int64(); err == nil {
			x = v
		} else {
			return err
		}
		o.Result = EGLBoolean(x)
	}
	return nil
}

type GlXCreateContext_Postback struct {
	Result []byte
}

func (o *GlXCreateContext_Postback) Decode(result_cnt uint64, d binary.Decoder) error {
	{
		var x []byte
		if val, err := readBytes(d, result_cnt); err == nil {
			x = val
		} else {
			return err
		}
		o.Result = []byte(x)
	}
	return nil
}

type GlXCreateNewContext_Postback struct {
	Result []byte
}

func (o *GlXCreateNewContext_Postback) Decode(result_cnt uint64, d binary.Decoder) error {
	{
		var x []byte
		if val, err := readBytes(d, result_cnt); err == nil {
			x = val
		} else {
			return err
		}
		o.Result = []byte(x)
	}
	return nil
}

type WglCreateContext_Postback struct {
	Result []byte
}

func (o *WglCreateContext_Postback) Decode(result_cnt uint64, d binary.Decoder) error {
	{
		var x []byte
		if val, err := readBytes(d, result_cnt); err == nil {
			x = val
		} else {
			return err
		}
		o.Result = []byte(x)
	}
	return nil
}

type WglCreateContextAttribsARB_Postback struct {
	Result []byte
}

func (o *WglCreateContextAttribsARB_Postback) Decode(result_cnt uint64, d binary.Decoder) error {
	{
		var x []byte
		if val, err := readBytes(d, result_cnt); err == nil {
			x = val
		} else {
			return err
		}
		o.Result = []byte(x)
	}
	return nil
}

type WglMakeCurrent_Postback struct {
	Result BOOL
}

func (o *WglMakeCurrent_Postback) Decode(result_cnt uint64, d binary.Decoder) error {
	{
		var x int64
		if v, err := d.Int64(); err == nil {
			x = v
		} else {
			return err
		}
		o.Result = BOOL(x)
	}
	return nil
}

type CGLCreateContext_Postback struct {
	Ctx    []byte
	Result CGLError
}

func (o *CGLCreateContext_Postback) Decode(ctx_cnt uint64, d binary.Decoder) error {
	{
		var x []byte
		if val, err := readBytes(d, ctx_cnt); err == nil {
			x = val
		} else {
			return err
		}
		o.Ctx = []byte(x)
	}
	{
		var x int64
		if v, err := d.Int64(); err == nil {
			x = v
		} else {
			return err
		}
		o.Result = CGLError(x)
	}
	return nil
}

type CGLSetCurrentContext_Postback struct {
	Result CGLError
}

func (o *CGLSetCurrentContext_Postback) Decode(d binary.Decoder) error {
	{
		var x int64
		if v, err := d.Int64(); err == nil {
			x = v
		} else {
			return err
		}
		o.Result = CGLError(x)
	}
	return nil
}

type GlGetProgramBinaryOES_Postback struct {
	BytesWritten int32
	BinaryFormat uint32
	Binary       []byte
}

func (o *GlGetProgramBinaryOES_Postback) Decode(binary_cnt uint64, d binary.Decoder) error {
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

func (o *GlGenVertexArraysOES_Postback) Decode(arrays_cnt uint64, d binary.Decoder) error {
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

func (o *GlIsVertexArrayOES_Postback) Decode(d binary.Decoder) error {
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

func (o *GlGetGraphicsResetStatusEXT_Postback) Decode(d binary.Decoder) error {
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

func (o *GlGetActiveAttrib_Postback) Decode(name_cnt uint64, d binary.Decoder) error {
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

func (o *GlGetActiveUniform_Postback) Decode(name_cnt uint64, d binary.Decoder) error {
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

func (o *GlGetError_Postback) Decode(d binary.Decoder) error {
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

func (o *GlGetProgramiv_Postback) Decode(value_cnt uint64, d binary.Decoder) error {
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

func (o *GlGetShaderiv_Postback) Decode(value_cnt uint64, d binary.Decoder) error {
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

func (o *GlGetUniformLocation_Postback) Decode(d binary.Decoder) error {
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

func (o *GlGetAttribLocation_Postback) Decode(d binary.Decoder) error {
	{
		var x uint32
		if v, err := d.Uint32(); err == nil {
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

func (o *GlGetTexParameteriv_Postback) Decode(values_cnt uint64, d binary.Decoder) error {
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

func (o *GlGetTexParameterfv_Postback) Decode(values_cnt uint64, d binary.Decoder) error {
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

func (o *GlGetShaderPrecisionFormat_Postback) Decode(range_cnt uint64, d binary.Decoder) error {
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

func (o *GlGenTextures_Postback) Decode(textures_cnt uint64, d binary.Decoder) error {
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

func (o *GlIsTexture_Postback) Decode(d binary.Decoder) error {
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

func (o *GlReadPixels_Postback) Decode(data_cnt uint64, d binary.Decoder) error {
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

func (o *GlGenFramebuffers_Postback) Decode(framebuffers_cnt uint64, d binary.Decoder) error {
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

func (o *GlCheckFramebufferStatus_Postback) Decode(d binary.Decoder) error {
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

func (o *GlIsFramebuffer_Postback) Decode(d binary.Decoder) error {
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

func (o *GlGenRenderbuffers_Postback) Decode(renderbuffers_cnt uint64, d binary.Decoder) error {
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

func (o *GlIsRenderbuffer_Postback) Decode(d binary.Decoder) error {
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

func (o *GlGetRenderbufferParameteriv_Postback) Decode(values_cnt uint64, d binary.Decoder) error {
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

func (o *GlGenBuffers_Postback) Decode(buffers_cnt uint64, d binary.Decoder) error {
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

func (o *GlIsBuffer_Postback) Decode(d binary.Decoder) error {
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

func (o *GlGetBufferParameteriv_Postback) Decode(d binary.Decoder) error {
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

func (o *GlCreateShader_Postback) Decode(d binary.Decoder) error {
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

func (o *GlGetShaderInfoLog_Postback) Decode(info_cnt uint64, d binary.Decoder) error {
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

func (o *GlGetShaderSource_Postback) Decode(source_cnt uint64, d binary.Decoder) error {
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

func (o *GlIsShader_Postback) Decode(d binary.Decoder) error {
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

func (o *GlCreateProgram_Postback) Decode(d binary.Decoder) error {
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

func (o *GlGetAttachedShaders_Postback) Decode(shaders_cnt uint64, d binary.Decoder) error {
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

func (o *GlGetProgramInfoLog_Postback) Decode(info_cnt uint64, d binary.Decoder) error {
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

func (o *GlIsProgram_Postback) Decode(d binary.Decoder) error {
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

func (o *GlGetFramebufferAttachmentParameteriv_Postback) Decode(value_cnt uint64, d binary.Decoder) error {
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

func (o *GlGetBooleanv_Postback) Decode(values_cnt uint64, d binary.Decoder) error {
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

func (o *GlGetFloatv_Postback) Decode(values_cnt uint64, d binary.Decoder) error {
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

func (o *GlGetIntegerv_Postback) Decode(values_cnt uint64, d binary.Decoder) error {
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

func (o *GlGetString_Postback) Decode(result_cnt uint64, d binary.Decoder) error {
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

func (o *GlIsEnabled_Postback) Decode(d binary.Decoder) error {
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

func (o *GlMapBufferRange_Postback) Decode(result_cnt uint64, d binary.Decoder) error {
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

func (o *GlGenQueries_Postback) Decode(queries_cnt uint64, d binary.Decoder) error {
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

func (o *GlIsQuery_Postback) Decode(d binary.Decoder) error {
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

func (o *GlGetQueryiv_Postback) Decode(d binary.Decoder) error {
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

func (o *GlGetQueryObjectuiv_Postback) Decode(d binary.Decoder) error {
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

func (o *GlGenQueriesEXT_Postback) Decode(queries_cnt uint64, d binary.Decoder) error {
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

func (o *GlIsQueryEXT_Postback) Decode(d binary.Decoder) error {
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

func (o *GlGetQueryivEXT_Postback) Decode(d binary.Decoder) error {
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

func (o *GlGetQueryObjectivEXT_Postback) Decode(d binary.Decoder) error {
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

func (o *GlGetQueryObjectuivEXT_Postback) Decode(d binary.Decoder) error {
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

func (o *GlGetQueryObjecti64vEXT_Postback) Decode(d binary.Decoder) error {
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

func (o *GlGetQueryObjectui64vEXT_Postback) Decode(d binary.Decoder) error {
	if v, err := d.Uint64(); err == nil {
		o.Value = v
	} else {
		return err
	}
	return nil
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

var _ = replay.Replayer(&ReplayCreateRenderer{}) // interface compliance check
func (ϟa *ReplayCreateRenderer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Id))
	ϟb.CallNoPush(funcInfoReplayCreateRenderer)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&ReplayBindRenderer{}) // interface compliance check
func (ϟa *ReplayBindRenderer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Id))
	ϟb.CallNoPush(funcInfoReplayBindRenderer)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&BackbufferInfo{}) // interface compliance check
func (ϟa *BackbufferInfo) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.Push(value.U32(ϟa.ColorFmt))
	ϟb.Push(value.U32(ϟa.DepthFmt))
	ϟb.Push(value.U32(ϟa.StencilFmt))
	ϟb.Push(value.Bool(ϟa.ResetViewportScissor))
	ϟb.CallNoPush(funcInfoBackbufferInfo)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&StartTimer{}) // interface compliance check
func (ϟa *StartTimer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U8(ϟa.Index))
	ϟb.CallNoPush(funcInfoStartTimer)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&StopTimer{}) // interface compliance check
func (ϟa *StopTimer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{8 /* result */})
	ϟb.Push(value.U8(ϟa.Index))
	ϟb.CallPush(funcInfoStopTimer)
	ϟb.Store(outputs[0])
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := StopTimer_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&FlushPostBuffer{}) // interface compliance check
func (ϟa *FlushPostBuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.CallNoPush(funcInfoFlushPostBuffer)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlEnableClientState{}) // interface compliance check
func (ϟa *GlEnableClientState) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.CallNoPush(funcInfoGlEnableClientState)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlDisableClientState{}) // interface compliance check
func (ϟa *GlDisableClientState) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.CallNoPush(funcInfoGlDisableClientState)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGetProgramBinaryOES{}) // interface compliance check
func (ϟa *GlGetProgramBinaryOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	binary_cnt := uint64(ϟa.BufferSize)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{4 /* bytes_written */, 4 /* binary_format */, binary_cnt /* binary */})
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.BufferSize))
	ϟb.Push(outputs[0]) // bytes_written
	ϟb.Push(outputs[1]) // binary_format
	ϟb.Push(outputs[2]) // binary
	ϟb.CallNoPush(funcInfoGlGetProgramBinaryOES)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGetProgramBinaryOES_Postback{}
			if err := postback.Decode(binary_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlProgramBinaryOES{}) // interface compliance check
func (ϟa *GlProgramBinaryOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.BinaryFormat))
	ϟb.Push(value.VolatileCapturePointer(uint64(ϟa.Binary)))
	ϟb.Push(value.S32(ϟa.BinarySize))
	ϟb.CallNoPush(funcInfoGlProgramBinaryOES)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlStartTilingQCOM{}) // interface compliance check
func (ϟa *GlStartTilingQCOM) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.S32(ϟa.X))
	ϟb.Push(value.S32(ϟa.Y))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.Push(value.U32(ϟa.PreserveMask))
	ϟb.CallNoPush(funcInfoGlStartTilingQCOM)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlEndTilingQCOM{}) // interface compliance check
func (ϟa *GlEndTilingQCOM) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.PreserveMask))
	ϟb.CallNoPush(funcInfoGlEndTilingQCOM)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlDiscardFramebufferEXT{}) // interface compliance check
func (ϟa *GlDiscardFramebufferEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.NumAttachments))
	ϟb.Push(ϟa.Attachments.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlDiscardFramebufferEXT)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlInsertEventMarkerEXT{}) // interface compliance check
func (ϟa *GlInsertEventMarkerEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.S32(ϟa.Length))
	ϟb.Push(ϟb.String(ϟa.Marker))
	ϟb.CallNoPush(funcInfoGlInsertEventMarkerEXT)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlPushGroupMarkerEXT{}) // interface compliance check
func (ϟa *GlPushGroupMarkerEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.S32(ϟa.Length))
	ϟb.Push(ϟb.String(ϟa.Marker))
	ϟb.CallNoPush(funcInfoGlPushGroupMarkerEXT)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlPopGroupMarkerEXT{}) // interface compliance check
func (ϟa *GlPopGroupMarkerEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.CallNoPush(funcInfoGlPopGroupMarkerEXT)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlTexStorage1DEXT{}) // interface compliance check
func (ϟa *GlTexStorage1DEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Levels))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.CallNoPush(funcInfoGlTexStorage1DEXT)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlTexStorage2DEXT{}) // interface compliance check
func (ϟa *GlTexStorage2DEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Levels))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.CallNoPush(funcInfoGlTexStorage2DEXT)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlTexStorage3DEXT{}) // interface compliance check
func (ϟa *GlTexStorage3DEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Levels))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.Push(value.S32(ϟa.Depth))
	ϟb.CallNoPush(funcInfoGlTexStorage3DEXT)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlTextureStorage1DEXT{}) // interface compliance check
func (ϟa *GlTextureStorage1DEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Texture.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Texture.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Texture.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Levels))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.CallNoPush(funcInfoGlTextureStorage1DEXT)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlTextureStorage2DEXT{}) // interface compliance check
func (ϟa *GlTextureStorage2DEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Texture.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Texture.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Texture.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Levels))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.CallNoPush(funcInfoGlTextureStorage2DEXT)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlTextureStorage3DEXT{}) // interface compliance check
func (ϟa *GlTextureStorage3DEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Texture.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Texture.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Texture.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Levels))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.Push(value.S32(ϟa.Depth))
	ϟb.CallNoPush(funcInfoGlTextureStorage3DEXT)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGenVertexArraysOES{}) // interface compliance check
func (ϟa *GlGenVertexArraysOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	arrays_cnt := uint64(ϟa.Count)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{arrays_cnt * 4 /* arrays */})
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(outputs[0]) // arrays
	ϟb.CallNoPush(funcInfoGlGenVertexArraysOES)
	ϟa.Mutate(ϟs)
	for i, e := range ϟa.Arrays {
		ptr := outputs[0].Offset(uint64(i * 4))
		if key, remap := e.remap(ϟa, ϟs); remap {
			storeRemap(ϟb, key, ptr, protocol.TypeUint32)
		}
	}
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGenVertexArraysOES_Postback{}
			if err := postback.Decode(arrays_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlBindVertexArrayOES{}) // interface compliance check
func (ϟa *GlBindVertexArrayOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Array.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Array.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Array.value(ϟb, ϟa, ϟs))
	}
	ϟb.CallNoPush(funcInfoGlBindVertexArrayOES)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlDeleteVertexArraysOES{}) // interface compliance check
func (ϟa *GlDeleteVertexArraysOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Arrays.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlDeleteVertexArraysOES)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlIsVertexArrayOES{}) // interface compliance check
func (ϟa *GlIsVertexArrayOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{1 /* result */})
	if key, remap := ϟa.Array.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Array.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Array.value(ϟb, ϟa, ϟs))
	}
	ϟb.CallPush(funcInfoGlIsVertexArrayOES)
	ϟb.Store(outputs[0])
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlIsVertexArrayOES_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlEGLImageTargetTexture2DOES{}) // interface compliance check
func (ϟa *GlEGLImageTargetTexture2DOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Image.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlEGLImageTargetTexture2DOES)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlEGLImageTargetRenderbufferStorageOES{}) // interface compliance check
func (ϟa *GlEGLImageTargetRenderbufferStorageOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Image.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlEGLImageTargetRenderbufferStorageOES)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGetGraphicsResetStatusEXT{}) // interface compliance check
func (ϟa *GlGetGraphicsResetStatusEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{4 /* result */})
	ϟb.CallPush(funcInfoGlGetGraphicsResetStatusEXT)
	ϟb.Store(outputs[0])
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGetGraphicsResetStatusEXT_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlBindAttribLocation{}) // interface compliance check
func (ϟa *GlBindAttribLocation) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟb.String(ϟa.Name))
	ϟb.CallNoPush(funcInfoGlBindAttribLocation)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlBlendFunc{}) // interface compliance check
func (ϟa *GlBlendFunc) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.SrcFactor))
	ϟb.Push(value.U32(ϟa.DstFactor))
	ϟb.CallNoPush(funcInfoGlBlendFunc)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlBlendFuncSeparate{}) // interface compliance check
func (ϟa *GlBlendFuncSeparate) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.SrcFactorRgb))
	ϟb.Push(value.U32(ϟa.DstFactorRgb))
	ϟb.Push(value.U32(ϟa.SrcFactorAlpha))
	ϟb.Push(value.U32(ϟa.DstFactorAlpha))
	ϟb.CallNoPush(funcInfoGlBlendFuncSeparate)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlBlendEquation{}) // interface compliance check
func (ϟa *GlBlendEquation) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Equation))
	ϟb.CallNoPush(funcInfoGlBlendEquation)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlBlendEquationSeparate{}) // interface compliance check
func (ϟa *GlBlendEquationSeparate) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Rgb))
	ϟb.Push(value.U32(ϟa.Alpha))
	ϟb.CallNoPush(funcInfoGlBlendEquationSeparate)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlBlendColor{}) // interface compliance check
func (ϟa *GlBlendColor) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.F32(ϟa.Red))
	ϟb.Push(value.F32(ϟa.Green))
	ϟb.Push(value.F32(ϟa.Blue))
	ϟb.Push(value.F32(ϟa.Alpha))
	ϟb.CallNoPush(funcInfoGlBlendColor)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlEnableVertexAttribArray{}) // interface compliance check
func (ϟa *GlEnableVertexAttribArray) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlEnableVertexAttribArray)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlDisableVertexAttribArray{}) // interface compliance check
func (ϟa *GlDisableVertexAttribArray) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlDisableVertexAttribArray)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlVertexAttribPointer{}) // interface compliance check
func (ϟa *GlVertexAttribPointer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.S32(ϟa.Size))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(value.Bool(ϟa.Normalized))
	ϟb.Push(value.S32(ϟa.Stride))
	ϟb.Push(ϟa.Data.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlVertexAttribPointer)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGetActiveAttrib{}) // interface compliance check
func (ϟa *GlGetActiveAttrib) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	name_cnt := uint64(ϟa.BufferSize)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{4 /* buffer_bytes_written */, 4 /* vector_count */, 4 /* type */, name_cnt /* name */})
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.S32(ϟa.BufferSize))
	ϟb.Push(outputs[0]) // buffer_bytes_written
	ϟb.Push(outputs[1]) // vector_count
	ϟb.Push(outputs[2]) // type
	ϟb.Push(outputs[3]) // name
	ϟb.CallNoPush(funcInfoGlGetActiveAttrib)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGetActiveAttrib_Postback{}
			if err := postback.Decode(name_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGetActiveUniform{}) // interface compliance check
func (ϟa *GlGetActiveUniform) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	name_cnt := uint64(ϟa.BufferSize)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{4 /* buffer_bytes_written */, 4 /* size */, 4 /* type */, name_cnt /* name */})
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Location))
	ϟb.Push(value.S32(ϟa.BufferSize))
	ϟb.Push(outputs[0]) // buffer_bytes_written
	ϟb.Push(outputs[1]) // size
	ϟb.Push(outputs[2]) // type
	ϟb.Push(outputs[3]) // name
	ϟb.CallNoPush(funcInfoGlGetActiveUniform)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGetActiveUniform_Postback{}
			if err := postback.Decode(name_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGetError{}) // interface compliance check
func (ϟa *GlGetError) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{4 /* result */})
	ϟb.CallPush(funcInfoGlGetError)
	ϟb.Store(outputs[0])
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGetError_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGetProgramiv{}) // interface compliance check
func (ϟa *GlGetProgramiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	value_cnt := uint64(int32(1))
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{value_cnt * 4 /* value */})
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(outputs[0]) // value
	ϟb.CallNoPush(funcInfoGlGetProgramiv)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGetProgramiv_Postback{}
			if err := postback.Decode(value_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGetShaderiv{}) // interface compliance check
func (ϟa *GlGetShaderiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	value_cnt := uint64(int32(1))
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{value_cnt * 4 /* value */})
	if key, remap := ϟa.Shader.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Shader.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Shader.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(outputs[0]) // value
	ϟb.CallNoPush(funcInfoGlGetShaderiv)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGetShaderiv_Postback{}
			if err := postback.Decode(value_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGetUniformLocation{}) // interface compliance check
func (ϟa *GlGetUniformLocation) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{4 /* result */})
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟb.String(ϟa.Name))
	ϟb.CallPush(funcInfoGlGetUniformLocation)
	ϟb.Store(outputs[0])
	ϟa.Mutate(ϟs)
	if key, remap := ϟa.Result.remap(ϟa, ϟs); remap {
		storeRemap(ϟb, key, outputs[0], protocol.TypeInt32)
	}
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGetUniformLocation_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGetAttribLocation{}) // interface compliance check
func (ϟa *GlGetAttribLocation) defaultReplay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{4 /* result */})
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟb.String(ϟa.Name))
	ϟb.CallPush(funcInfoGlGetAttribLocation)
	ϟb.Store(outputs[0])
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGetAttribLocation_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlPixelStorei{}) // interface compliance check
func (ϟa *GlPixelStorei) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(value.S32(ϟa.Value))
	ϟb.CallNoPush(funcInfoGlPixelStorei)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlTexParameteri{}) // interface compliance check
func (ϟa *GlTexParameteri) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(value.S32(ϟa.Value))
	ϟb.CallNoPush(funcInfoGlTexParameteri)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlTexParameterf{}) // interface compliance check
func (ϟa *GlTexParameterf) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(value.F32(ϟa.Value))
	ϟb.CallNoPush(funcInfoGlTexParameterf)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGetTexParameteriv{}) // interface compliance check
func (ϟa *GlGetTexParameteriv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	values_cnt := uint64(int32(1))
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{values_cnt * 4 /* values */})
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(outputs[0]) // values
	ϟb.CallNoPush(funcInfoGlGetTexParameteriv)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGetTexParameteriv_Postback{}
			if err := postback.Decode(values_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGetTexParameterfv{}) // interface compliance check
func (ϟa *GlGetTexParameterfv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	values_cnt := uint64(int32(1))
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{values_cnt * 4 /* values */})
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(outputs[0]) // values
	ϟb.CallNoPush(funcInfoGlGetTexParameterfv)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGetTexParameterfv_Postback{}
			if err := postback.Decode(values_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlUniform1i{}) // interface compliance check
func (ϟa *GlUniform1i) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Value))
	ϟb.CallNoPush(funcInfoGlUniform1i)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlUniform2i{}) // interface compliance check
func (ϟa *GlUniform2i) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Value0))
	ϟb.Push(value.S32(ϟa.Value1))
	ϟb.CallNoPush(funcInfoGlUniform2i)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlUniform3i{}) // interface compliance check
func (ϟa *GlUniform3i) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Value0))
	ϟb.Push(value.S32(ϟa.Value1))
	ϟb.Push(value.S32(ϟa.Value2))
	ϟb.CallNoPush(funcInfoGlUniform3i)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlUniform4i{}) // interface compliance check
func (ϟa *GlUniform4i) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Value0))
	ϟb.Push(value.S32(ϟa.Value1))
	ϟb.Push(value.S32(ϟa.Value2))
	ϟb.Push(value.S32(ϟa.Value3))
	ϟb.CallNoPush(funcInfoGlUniform4i)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlUniform1iv{}) // interface compliance check
func (ϟa *GlUniform1iv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Value.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlUniform1iv)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlUniform2iv{}) // interface compliance check
func (ϟa *GlUniform2iv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Value.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlUniform2iv)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlUniform3iv{}) // interface compliance check
func (ϟa *GlUniform3iv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Value.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlUniform3iv)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlUniform4iv{}) // interface compliance check
func (ϟa *GlUniform4iv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Value.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlUniform4iv)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlUniform1f{}) // interface compliance check
func (ϟa *GlUniform1f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.F32(ϟa.Value))
	ϟb.CallNoPush(funcInfoGlUniform1f)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlUniform2f{}) // interface compliance check
func (ϟa *GlUniform2f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.F32(ϟa.Value0))
	ϟb.Push(value.F32(ϟa.Value1))
	ϟb.CallNoPush(funcInfoGlUniform2f)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlUniform3f{}) // interface compliance check
func (ϟa *GlUniform3f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.F32(ϟa.Value0))
	ϟb.Push(value.F32(ϟa.Value1))
	ϟb.Push(value.F32(ϟa.Value2))
	ϟb.CallNoPush(funcInfoGlUniform3f)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlUniform4f{}) // interface compliance check
func (ϟa *GlUniform4f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.F32(ϟa.Value0))
	ϟb.Push(value.F32(ϟa.Value1))
	ϟb.Push(value.F32(ϟa.Value2))
	ϟb.Push(value.F32(ϟa.Value3))
	ϟb.CallNoPush(funcInfoGlUniform4f)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlUniform1fv{}) // interface compliance check
func (ϟa *GlUniform1fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Value.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlUniform1fv)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlUniform2fv{}) // interface compliance check
func (ϟa *GlUniform2fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Value.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlUniform2fv)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlUniform3fv{}) // interface compliance check
func (ϟa *GlUniform3fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Value.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlUniform3fv)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlUniform4fv{}) // interface compliance check
func (ϟa *GlUniform4fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Value.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlUniform4fv)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlUniformMatrix2fv{}) // interface compliance check
func (ϟa *GlUniformMatrix2fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(value.Bool(ϟa.Transpose))
	ϟb.Push(ϟa.Values.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlUniformMatrix2fv)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlUniformMatrix3fv{}) // interface compliance check
func (ϟa *GlUniformMatrix3fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(value.Bool(ϟa.Transpose))
	ϟb.Push(ϟa.Values.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlUniformMatrix3fv)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlUniformMatrix4fv{}) // interface compliance check
func (ϟa *GlUniformMatrix4fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(value.Bool(ϟa.Transpose))
	ϟb.Push(ϟa.Values.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlUniformMatrix4fv)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGetUniformfv{}) // interface compliance check
func (ϟa *GlGetUniformfv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Values.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlGetUniformfv)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGetUniformiv{}) // interface compliance check
func (ϟa *GlGetUniformiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Values.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlGetUniformiv)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlVertexAttrib1f{}) // interface compliance check
func (ϟa *GlVertexAttrib1f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.F32(ϟa.Value0))
	ϟb.CallNoPush(funcInfoGlVertexAttrib1f)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlVertexAttrib2f{}) // interface compliance check
func (ϟa *GlVertexAttrib2f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.F32(ϟa.Value0))
	ϟb.Push(value.F32(ϟa.Value1))
	ϟb.CallNoPush(funcInfoGlVertexAttrib2f)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlVertexAttrib3f{}) // interface compliance check
func (ϟa *GlVertexAttrib3f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.F32(ϟa.Value0))
	ϟb.Push(value.F32(ϟa.Value1))
	ϟb.Push(value.F32(ϟa.Value2))
	ϟb.CallNoPush(funcInfoGlVertexAttrib3f)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlVertexAttrib4f{}) // interface compliance check
func (ϟa *GlVertexAttrib4f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.F32(ϟa.Value0))
	ϟb.Push(value.F32(ϟa.Value1))
	ϟb.Push(value.F32(ϟa.Value2))
	ϟb.Push(value.F32(ϟa.Value3))
	ϟb.CallNoPush(funcInfoGlVertexAttrib4f)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlVertexAttrib1fv{}) // interface compliance check
func (ϟa *GlVertexAttrib1fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlVertexAttrib1fv)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlVertexAttrib2fv{}) // interface compliance check
func (ϟa *GlVertexAttrib2fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlVertexAttrib2fv)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlVertexAttrib3fv{}) // interface compliance check
func (ϟa *GlVertexAttrib3fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlVertexAttrib3fv)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlVertexAttrib4fv{}) // interface compliance check
func (ϟa *GlVertexAttrib4fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlVertexAttrib4fv)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGetShaderPrecisionFormat{}) // interface compliance check
func (ϟa *GlGetShaderPrecisionFormat) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	range_cnt := uint64(int32(2))
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{range_cnt * 4 /* range */, 4 /* precision */})
	ϟb.Push(value.U32(ϟa.ShaderType))
	ϟb.Push(value.U32(ϟa.PrecisionType))
	ϟb.Push(outputs[0]) // range
	ϟb.Push(outputs[1]) // precision
	ϟb.CallNoPush(funcInfoGlGetShaderPrecisionFormat)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGetShaderPrecisionFormat_Postback{}
			if err := postback.Decode(range_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlDepthMask{}) // interface compliance check
func (ϟa *GlDepthMask) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.Bool(ϟa.Enabled))
	ϟb.CallNoPush(funcInfoGlDepthMask)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlDepthFunc{}) // interface compliance check
func (ϟa *GlDepthFunc) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Function))
	ϟb.CallNoPush(funcInfoGlDepthFunc)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlDepthRangef{}) // interface compliance check
func (ϟa *GlDepthRangef) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.F32(ϟa.Near))
	ϟb.Push(value.F32(ϟa.Far))
	ϟb.CallNoPush(funcInfoGlDepthRangef)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlColorMask{}) // interface compliance check
func (ϟa *GlColorMask) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.Bool(ϟa.Red))
	ϟb.Push(value.Bool(ϟa.Green))
	ϟb.Push(value.Bool(ϟa.Blue))
	ϟb.Push(value.Bool(ϟa.Alpha))
	ϟb.CallNoPush(funcInfoGlColorMask)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlStencilMask{}) // interface compliance check
func (ϟa *GlStencilMask) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Mask))
	ϟb.CallNoPush(funcInfoGlStencilMask)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlStencilMaskSeparate{}) // interface compliance check
func (ϟa *GlStencilMaskSeparate) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Face))
	ϟb.Push(value.U32(ϟa.Mask))
	ϟb.CallNoPush(funcInfoGlStencilMaskSeparate)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlStencilFuncSeparate{}) // interface compliance check
func (ϟa *GlStencilFuncSeparate) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Face))
	ϟb.Push(value.U32(ϟa.Function))
	ϟb.Push(value.S32(ϟa.ReferenceValue))
	ϟb.Push(value.S32(ϟa.Mask))
	ϟb.CallNoPush(funcInfoGlStencilFuncSeparate)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlStencilOpSeparate{}) // interface compliance check
func (ϟa *GlStencilOpSeparate) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Face))
	ϟb.Push(value.U32(ϟa.StencilFail))
	ϟb.Push(value.U32(ϟa.StencilPassDepthFail))
	ϟb.Push(value.U32(ϟa.StencilPassDepthPass))
	ϟb.CallNoPush(funcInfoGlStencilOpSeparate)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlFrontFace{}) // interface compliance check
func (ϟa *GlFrontFace) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Orientation))
	ϟb.CallNoPush(funcInfoGlFrontFace)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlViewport{}) // interface compliance check
func (ϟa *GlViewport) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.S32(ϟa.X))
	ϟb.Push(value.S32(ϟa.Y))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.CallNoPush(funcInfoGlViewport)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlScissor{}) // interface compliance check
func (ϟa *GlScissor) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.S32(ϟa.X))
	ϟb.Push(value.S32(ϟa.Y))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.CallNoPush(funcInfoGlScissor)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlActiveTexture{}) // interface compliance check
func (ϟa *GlActiveTexture) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Unit))
	ϟb.CallNoPush(funcInfoGlActiveTexture)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGenTextures{}) // interface compliance check
func (ϟa *GlGenTextures) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	textures_cnt := uint64(ϟa.Count)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{textures_cnt * 4 /* textures */})
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(outputs[0]) // textures
	ϟb.CallNoPush(funcInfoGlGenTextures)
	ϟa.Mutate(ϟs)
	for i, e := range ϟa.Textures {
		ptr := outputs[0].Offset(uint64(i * 4))
		if key, remap := e.remap(ϟa, ϟs); remap {
			storeRemap(ϟb, key, ptr, protocol.TypeUint32)
		}
	}
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGenTextures_Postback{}
			if err := postback.Decode(textures_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlDeleteTextures{}) // interface compliance check
func (ϟa *GlDeleteTextures) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Textures.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlDeleteTextures)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlIsTexture{}) // interface compliance check
func (ϟa *GlIsTexture) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{1 /* result */})
	if key, remap := ϟa.Texture.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Texture.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Texture.value(ϟb, ϟa, ϟs))
	}
	ϟb.CallPush(funcInfoGlIsTexture)
	ϟb.Store(outputs[0])
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlIsTexture_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlBindTexture{}) // interface compliance check
func (ϟa *GlBindTexture) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Target))
	if key, remap := ϟa.Texture.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Texture.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Texture.value(ϟb, ϟa, ϟs))
	}
	ϟb.CallNoPush(funcInfoGlBindTexture)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlTexImage2D{}) // interface compliance check
func (ϟa *GlTexImage2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Level))
	ϟb.Push(value.U32(ϟa.InternalFormat))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.Push(value.S32(ϟa.Border))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Data.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlTexImage2D)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlTexSubImage2D{}) // interface compliance check
func (ϟa *GlTexSubImage2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Level))
	ϟb.Push(value.S32(ϟa.Xoffset))
	ϟb.Push(value.S32(ϟa.Yoffset))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Data.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlTexSubImage2D)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlCopyTexImage2D{}) // interface compliance check
func (ϟa *GlCopyTexImage2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Level))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.S32(ϟa.X))
	ϟb.Push(value.S32(ϟa.Y))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.Push(value.S32(ϟa.Border))
	ϟb.CallNoPush(funcInfoGlCopyTexImage2D)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlCopyTexSubImage2D{}) // interface compliance check
func (ϟa *GlCopyTexSubImage2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Level))
	ϟb.Push(value.S32(ϟa.Xoffset))
	ϟb.Push(value.S32(ϟa.Yoffset))
	ϟb.Push(value.S32(ϟa.X))
	ϟb.Push(value.S32(ϟa.Y))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.CallNoPush(funcInfoGlCopyTexSubImage2D)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlCompressedTexImage2D{}) // interface compliance check
func (ϟa *GlCompressedTexImage2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Level))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.Push(value.S32(ϟa.Border))
	ϟb.Push(value.S32(ϟa.ImageSize))
	ϟb.Push(ϟa.Data.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlCompressedTexImage2D)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlCompressedTexSubImage2D{}) // interface compliance check
func (ϟa *GlCompressedTexSubImage2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Level))
	ϟb.Push(value.S32(ϟa.Xoffset))
	ϟb.Push(value.S32(ϟa.Yoffset))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.S32(ϟa.ImageSize))
	ϟb.Push(ϟa.Data.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlCompressedTexSubImage2D)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGenerateMipmap{}) // interface compliance check
func (ϟa *GlGenerateMipmap) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.CallNoPush(funcInfoGlGenerateMipmap)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlReadPixels{}) // interface compliance check
func (ϟa *GlReadPixels) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	data_cnt := uint64(imageSize(uint32(ϟa.Width), uint32(ϟa.Height), TexelFormat(ϟa.Format), ϟa.Type))
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{data_cnt /* data */})
	ϟb.Push(value.S32(ϟa.X))
	ϟb.Push(value.S32(ϟa.Y))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(outputs[0]) // data
	ϟb.CallNoPush(funcInfoGlReadPixels)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlReadPixels_Postback{}
			if err := postback.Decode(data_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGenFramebuffers{}) // interface compliance check
func (ϟa *GlGenFramebuffers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	framebuffers_cnt := uint64(ϟa.Count)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{framebuffers_cnt * 4 /* framebuffers */})
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(outputs[0]) // framebuffers
	ϟb.CallNoPush(funcInfoGlGenFramebuffers)
	ϟa.Mutate(ϟs)
	for i, e := range ϟa.Framebuffers {
		ptr := outputs[0].Offset(uint64(i * 4))
		if key, remap := e.remap(ϟa, ϟs); remap {
			storeRemap(ϟb, key, ptr, protocol.TypeUint32)
		}
	}
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGenFramebuffers_Postback{}
			if err := postback.Decode(framebuffers_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlBindFramebuffer{}) // interface compliance check
func (ϟa *GlBindFramebuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Target))
	if key, remap := ϟa.Framebuffer.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Framebuffer.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Framebuffer.value(ϟb, ϟa, ϟs))
	}
	ϟb.CallNoPush(funcInfoGlBindFramebuffer)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlCheckFramebufferStatus{}) // interface compliance check
func (ϟa *GlCheckFramebufferStatus) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{4 /* result */})
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.CallPush(funcInfoGlCheckFramebufferStatus)
	ϟb.Store(outputs[0])
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlCheckFramebufferStatus_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlDeleteFramebuffers{}) // interface compliance check
func (ϟa *GlDeleteFramebuffers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Framebuffers.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlDeleteFramebuffers)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlIsFramebuffer{}) // interface compliance check
func (ϟa *GlIsFramebuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{1 /* result */})
	if key, remap := ϟa.Framebuffer.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Framebuffer.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Framebuffer.value(ϟb, ϟa, ϟs))
	}
	ϟb.CallPush(funcInfoGlIsFramebuffer)
	ϟb.Store(outputs[0])
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlIsFramebuffer_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGenRenderbuffers{}) // interface compliance check
func (ϟa *GlGenRenderbuffers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	renderbuffers_cnt := uint64(ϟa.Count)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{renderbuffers_cnt * 4 /* renderbuffers */})
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(outputs[0]) // renderbuffers
	ϟb.CallNoPush(funcInfoGlGenRenderbuffers)
	ϟa.Mutate(ϟs)
	for i, e := range ϟa.Renderbuffers {
		ptr := outputs[0].Offset(uint64(i * 4))
		if key, remap := e.remap(ϟa, ϟs); remap {
			storeRemap(ϟb, key, ptr, protocol.TypeUint32)
		}
	}
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGenRenderbuffers_Postback{}
			if err := postback.Decode(renderbuffers_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlBindRenderbuffer{}) // interface compliance check
func (ϟa *GlBindRenderbuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Target))
	if key, remap := ϟa.Renderbuffer.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Renderbuffer.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Renderbuffer.value(ϟb, ϟa, ϟs))
	}
	ϟb.CallNoPush(funcInfoGlBindRenderbuffer)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlRenderbufferStorage{}) // interface compliance check
func (ϟa *GlRenderbufferStorage) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.CallNoPush(funcInfoGlRenderbufferStorage)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlDeleteRenderbuffers{}) // interface compliance check
func (ϟa *GlDeleteRenderbuffers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Renderbuffers.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlDeleteRenderbuffers)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlIsRenderbuffer{}) // interface compliance check
func (ϟa *GlIsRenderbuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{1 /* result */})
	if key, remap := ϟa.Renderbuffer.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Renderbuffer.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Renderbuffer.value(ϟb, ϟa, ϟs))
	}
	ϟb.CallPush(funcInfoGlIsRenderbuffer)
	ϟb.Store(outputs[0])
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlIsRenderbuffer_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGetRenderbufferParameteriv{}) // interface compliance check
func (ϟa *GlGetRenderbufferParameteriv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	values_cnt := uint64(int32(1))
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{values_cnt * 4 /* values */})
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(outputs[0]) // values
	ϟb.CallNoPush(funcInfoGlGetRenderbufferParameteriv)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGetRenderbufferParameteriv_Postback{}
			if err := postback.Decode(values_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGenBuffers{}) // interface compliance check
func (ϟa *GlGenBuffers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	buffers_cnt := uint64(ϟa.Count)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{buffers_cnt * 4 /* buffers */})
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(outputs[0]) // buffers
	ϟb.CallNoPush(funcInfoGlGenBuffers)
	ϟa.Mutate(ϟs)
	for i, e := range ϟa.Buffers {
		ptr := outputs[0].Offset(uint64(i * 4))
		if key, remap := e.remap(ϟa, ϟs); remap {
			storeRemap(ϟb, key, ptr, protocol.TypeUint32)
		}
	}
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGenBuffers_Postback{}
			if err := postback.Decode(buffers_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlBindBuffer{}) // interface compliance check
func (ϟa *GlBindBuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Target))
	if key, remap := ϟa.Buffer.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Buffer.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Buffer.value(ϟb, ϟa, ϟs))
	}
	ϟb.CallNoPush(funcInfoGlBindBuffer)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlBufferData{}) // interface compliance check
func (ϟa *GlBufferData) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Size))
	ϟb.Push(ϟa.Data.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Usage))
	ϟb.CallNoPush(funcInfoGlBufferData)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlBufferSubData{}) // interface compliance check
func (ϟa *GlBufferSubData) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Offset))
	ϟb.Push(value.S32(ϟa.Size))
	ϟb.Push(value.VolatileCapturePointer(uint64(ϟa.Data)))
	ϟb.CallNoPush(funcInfoGlBufferSubData)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlDeleteBuffers{}) // interface compliance check
func (ϟa *GlDeleteBuffers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Buffers.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlDeleteBuffers)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlIsBuffer{}) // interface compliance check
func (ϟa *GlIsBuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{1 /* result */})
	if key, remap := ϟa.Buffer.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Buffer.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Buffer.value(ϟb, ϟa, ϟs))
	}
	ϟb.CallPush(funcInfoGlIsBuffer)
	ϟb.Store(outputs[0])
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlIsBuffer_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGetBufferParameteriv{}) // interface compliance check
func (ϟa *GlGetBufferParameteriv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{4 /* value */})
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(outputs[0]) // value
	ϟb.CallNoPush(funcInfoGlGetBufferParameteriv)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGetBufferParameteriv_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlCreateShader{}) // interface compliance check
func (ϟa *GlCreateShader) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{4 /* result */})
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.CallPush(funcInfoGlCreateShader)
	ϟb.Store(outputs[0])
	ϟa.Mutate(ϟs)
	if key, remap := ϟa.Result.remap(ϟa, ϟs); remap {
		storeRemap(ϟb, key, outputs[0], protocol.TypeUint32)
	}
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlCreateShader_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlDeleteShader{}) // interface compliance check
func (ϟa *GlDeleteShader) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Shader.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Shader.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Shader.value(ϟb, ϟa, ϟs))
	}
	ϟb.CallNoPush(funcInfoGlDeleteShader)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlShaderSource{}) // interface compliance check
func (ϟa *GlShaderSource) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Shader.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Shader.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Shader.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Source.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Length.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlShaderSource)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlShaderBinary{}) // interface compliance check
func (ϟa *GlShaderBinary) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Shaders.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.BinaryFormat))
	ϟb.Push(value.VolatileCapturePointer(uint64(ϟa.Binary)))
	ϟb.Push(value.S32(ϟa.BinarySize))
	ϟb.CallNoPush(funcInfoGlShaderBinary)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGetShaderInfoLog{}) // interface compliance check
func (ϟa *GlGetShaderInfoLog) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	info_cnt := uint64(ϟa.BufferLength)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{4 /* string_length_written */, info_cnt /* info */})
	if key, remap := ϟa.Shader.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Shader.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Shader.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.BufferLength))
	ϟb.Push(outputs[0]) // string_length_written
	ϟb.Push(outputs[1]) // info
	ϟb.CallNoPush(funcInfoGlGetShaderInfoLog)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGetShaderInfoLog_Postback{}
			if err := postback.Decode(info_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGetShaderSource{}) // interface compliance check
func (ϟa *GlGetShaderSource) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	source_cnt := uint64(ϟa.BufferLength)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{4 /* string_length_written */, source_cnt /* source */})
	if key, remap := ϟa.Shader.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Shader.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Shader.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.BufferLength))
	ϟb.Push(outputs[0]) // string_length_written
	ϟb.Push(outputs[1]) // source
	ϟb.CallNoPush(funcInfoGlGetShaderSource)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGetShaderSource_Postback{}
			if err := postback.Decode(source_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlReleaseShaderCompiler{}) // interface compliance check
func (ϟa *GlReleaseShaderCompiler) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.CallNoPush(funcInfoGlReleaseShaderCompiler)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlCompileShader{}) // interface compliance check
func (ϟa *GlCompileShader) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Shader.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Shader.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Shader.value(ϟb, ϟa, ϟs))
	}
	ϟb.CallNoPush(funcInfoGlCompileShader)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlIsShader{}) // interface compliance check
func (ϟa *GlIsShader) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{1 /* result */})
	if key, remap := ϟa.Shader.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Shader.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Shader.value(ϟb, ϟa, ϟs))
	}
	ϟb.CallPush(funcInfoGlIsShader)
	ϟb.Store(outputs[0])
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlIsShader_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlCreateProgram{}) // interface compliance check
func (ϟa *GlCreateProgram) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{4 /* result */})
	ϟb.CallPush(funcInfoGlCreateProgram)
	ϟb.Store(outputs[0])
	ϟa.Mutate(ϟs)
	if key, remap := ϟa.Result.remap(ϟa, ϟs); remap {
		storeRemap(ϟb, key, outputs[0], protocol.TypeUint32)
	}
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlCreateProgram_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlDeleteProgram{}) // interface compliance check
func (ϟa *GlDeleteProgram) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.CallNoPush(funcInfoGlDeleteProgram)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlAttachShader{}) // interface compliance check
func (ϟa *GlAttachShader) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Shader.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Shader.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Shader.value(ϟb, ϟa, ϟs))
	}
	ϟb.CallNoPush(funcInfoGlAttachShader)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlDetachShader{}) // interface compliance check
func (ϟa *GlDetachShader) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Shader.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Shader.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Shader.value(ϟb, ϟa, ϟs))
	}
	ϟb.CallNoPush(funcInfoGlDetachShader)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGetAttachedShaders{}) // interface compliance check
func (ϟa *GlGetAttachedShaders) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	shaders_cnt := uint64(ϟa.BufferLength)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{4 /* shaders_length_written */, shaders_cnt * 4 /* shaders */})
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.BufferLength))
	ϟb.Push(outputs[0]) // shaders_length_written
	ϟb.Push(outputs[1]) // shaders
	ϟb.CallNoPush(funcInfoGlGetAttachedShaders)
	ϟa.Mutate(ϟs)
	for i, e := range ϟa.Shaders {
		ptr := outputs[1].Offset(uint64(i * 4))
		if key, remap := e.remap(ϟa, ϟs); remap {
			storeRemap(ϟb, key, ptr, protocol.TypeUint32)
		}
	}
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGetAttachedShaders_Postback{}
			if err := postback.Decode(shaders_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlLinkProgram{}) // interface compliance check
func (ϟa *GlLinkProgram) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.CallNoPush(funcInfoGlLinkProgram)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGetProgramInfoLog{}) // interface compliance check
func (ϟa *GlGetProgramInfoLog) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	info_cnt := uint64(ϟa.BufferLength)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{4 /* string_length_written */, info_cnt /* info */})
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.BufferLength))
	ϟb.Push(outputs[0]) // string_length_written
	ϟb.Push(outputs[1]) // info
	ϟb.CallNoPush(funcInfoGlGetProgramInfoLog)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGetProgramInfoLog_Postback{}
			if err := postback.Decode(info_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlUseProgram{}) // interface compliance check
func (ϟa *GlUseProgram) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.CallNoPush(funcInfoGlUseProgram)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlIsProgram{}) // interface compliance check
func (ϟa *GlIsProgram) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{1 /* result */})
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.CallPush(funcInfoGlIsProgram)
	ϟb.Store(outputs[0])
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlIsProgram_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlValidateProgram{}) // interface compliance check
func (ϟa *GlValidateProgram) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.CallNoPush(funcInfoGlValidateProgram)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlClearColor{}) // interface compliance check
func (ϟa *GlClearColor) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.F32(ϟa.R))
	ϟb.Push(value.F32(ϟa.G))
	ϟb.Push(value.F32(ϟa.B))
	ϟb.Push(value.F32(ϟa.A))
	ϟb.CallNoPush(funcInfoGlClearColor)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlClearDepthf{}) // interface compliance check
func (ϟa *GlClearDepthf) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.F32(ϟa.Depth))
	ϟb.CallNoPush(funcInfoGlClearDepthf)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlClearStencil{}) // interface compliance check
func (ϟa *GlClearStencil) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.S32(ϟa.Stencil))
	ϟb.CallNoPush(funcInfoGlClearStencil)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlClear{}) // interface compliance check
func (ϟa *GlClear) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Mask))
	ϟb.CallNoPush(funcInfoGlClear)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlCullFace{}) // interface compliance check
func (ϟa *GlCullFace) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.CallNoPush(funcInfoGlCullFace)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlPolygonOffset{}) // interface compliance check
func (ϟa *GlPolygonOffset) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.F32(ϟa.ScaleFactor))
	ϟb.Push(value.F32(ϟa.Units))
	ϟb.CallNoPush(funcInfoGlPolygonOffset)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlLineWidth{}) // interface compliance check
func (ϟa *GlLineWidth) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.F32(ϟa.Width))
	ϟb.CallNoPush(funcInfoGlLineWidth)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlSampleCoverage{}) // interface compliance check
func (ϟa *GlSampleCoverage) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.F32(ϟa.Value))
	ϟb.Push(value.Bool(ϟa.Invert))
	ϟb.CallNoPush(funcInfoGlSampleCoverage)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlHint{}) // interface compliance check
func (ϟa *GlHint) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.CallNoPush(funcInfoGlHint)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlFramebufferRenderbuffer{}) // interface compliance check
func (ϟa *GlFramebufferRenderbuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.FramebufferTarget))
	ϟb.Push(value.U32(ϟa.FramebufferAttachment))
	ϟb.Push(value.U32(ϟa.RenderbufferTarget))
	if key, remap := ϟa.Renderbuffer.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Renderbuffer.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Renderbuffer.value(ϟb, ϟa, ϟs))
	}
	ϟb.CallNoPush(funcInfoGlFramebufferRenderbuffer)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlFramebufferTexture2D{}) // interface compliance check
func (ϟa *GlFramebufferTexture2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.FramebufferTarget))
	ϟb.Push(value.U32(ϟa.FramebufferAttachment))
	ϟb.Push(value.U32(ϟa.TextureTarget))
	if key, remap := ϟa.Texture.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Texture.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Texture.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Level))
	ϟb.CallNoPush(funcInfoGlFramebufferTexture2D)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGetFramebufferAttachmentParameteriv{}) // interface compliance check
func (ϟa *GlGetFramebufferAttachmentParameteriv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	value_cnt := uint64(int32(1))
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{value_cnt * 4 /* value */})
	ϟb.Push(value.U32(ϟa.FramebufferTarget))
	ϟb.Push(value.U32(ϟa.Attachment))
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(outputs[0]) // value
	ϟb.CallNoPush(funcInfoGlGetFramebufferAttachmentParameteriv)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGetFramebufferAttachmentParameteriv_Postback{}
			if err := postback.Decode(value_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlDrawElements{}) // interface compliance check
func (ϟa *GlDrawElements) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.DrawMode))
	ϟb.Push(value.S32(ϟa.ElementCount))
	ϟb.Push(value.U32(ϟa.IndicesType))
	ϟb.Push(ϟa.Indices.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlDrawElements)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlDrawArrays{}) // interface compliance check
func (ϟa *GlDrawArrays) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.DrawMode))
	ϟb.Push(value.S32(ϟa.FirstIndex))
	ϟb.Push(value.S32(ϟa.IndexCount))
	ϟb.CallNoPush(funcInfoGlDrawArrays)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlFlush{}) // interface compliance check
func (ϟa *GlFlush) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.CallNoPush(funcInfoGlFlush)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlFinish{}) // interface compliance check
func (ϟa *GlFinish) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.CallNoPush(funcInfoGlFinish)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGetBooleanv{}) // interface compliance check
func (ϟa *GlGetBooleanv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	values_cnt := uint64(stateVariableSize(ϟa.Param))
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{values_cnt * 1 /* values */})
	ϟb.Push(value.U32(ϟa.Param))
	ϟb.Push(outputs[0]) // values
	ϟb.CallNoPush(funcInfoGlGetBooleanv)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGetBooleanv_Postback{}
			if err := postback.Decode(values_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGetFloatv{}) // interface compliance check
func (ϟa *GlGetFloatv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	values_cnt := uint64(stateVariableSize(ϟa.Param))
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{values_cnt * 4 /* values */})
	ϟb.Push(value.U32(ϟa.Param))
	ϟb.Push(outputs[0]) // values
	ϟb.CallNoPush(funcInfoGlGetFloatv)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGetFloatv_Postback{}
			if err := postback.Decode(values_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGetIntegerv{}) // interface compliance check
func (ϟa *GlGetIntegerv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	values_cnt := uint64(stateVariableSize(ϟa.Param))
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{values_cnt * 4 /* values */})
	ϟb.Push(value.U32(ϟa.Param))
	ϟb.Push(outputs[0]) // values
	ϟb.CallNoPush(funcInfoGlGetIntegerv)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGetIntegerv_Postback{}
			if err := postback.Decode(values_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGetString{}) // interface compliance check
func (ϟa *GlGetString) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	result_cnt := uint64(int32(256))
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{result_cnt /* result */})
	ϟb.Push(value.U32(ϟa.Param))
	ϟb.CallPush(funcInfoGlGetString)
	ϟb.Push(outputs[0])
	ϟb.Strcpy(result_cnt)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGetString_Postback{}
			if err := postback.Decode(result_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlEnable{}) // interface compliance check
func (ϟa *GlEnable) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Capability))
	ϟb.CallNoPush(funcInfoGlEnable)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlDisable{}) // interface compliance check
func (ϟa *GlDisable) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Capability))
	ϟb.CallNoPush(funcInfoGlDisable)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlIsEnabled{}) // interface compliance check
func (ϟa *GlIsEnabled) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{1 /* result */})
	ϟb.Push(value.U32(ϟa.Capability))
	ϟb.CallPush(funcInfoGlIsEnabled)
	ϟb.Store(outputs[0])
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlIsEnabled_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlMapBufferRange{}) // interface compliance check
func (ϟa *GlMapBufferRange) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	result_cnt := uint64(ϟa.Length)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{result_cnt /* result */})
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Offset))
	ϟb.Push(value.S32(ϟa.Length))
	ϟb.Push(value.U32(ϟa.Access))
	ϟb.CallPush(funcInfoGlMapBufferRange)
	ϟb.Push(outputs[0])
	ϟb.Copy(result_cnt)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlMapBufferRange_Postback{}
			if err := postback.Decode(result_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlUnmapBuffer{}) // interface compliance check
func (ϟa *GlUnmapBuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.CallNoPush(funcInfoGlUnmapBuffer)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlInvalidateFramebuffer{}) // interface compliance check
func (ϟa *GlInvalidateFramebuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Attachments.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlInvalidateFramebuffer)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlRenderbufferStorageMultisample{}) // interface compliance check
func (ϟa *GlRenderbufferStorageMultisample) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Samples))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.CallNoPush(funcInfoGlRenderbufferStorageMultisample)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlBlitFramebuffer{}) // interface compliance check
func (ϟa *GlBlitFramebuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.S32(ϟa.SrcX0))
	ϟb.Push(value.S32(ϟa.SrcY0))
	ϟb.Push(value.S32(ϟa.SrcX1))
	ϟb.Push(value.S32(ϟa.SrcY1))
	ϟb.Push(value.S32(ϟa.DstX0))
	ϟb.Push(value.S32(ϟa.DstY0))
	ϟb.Push(value.S32(ϟa.DstX1))
	ϟb.Push(value.S32(ϟa.DstY1))
	ϟb.Push(value.U32(ϟa.Mask))
	ϟb.Push(value.U32(ϟa.Filter))
	ϟb.CallNoPush(funcInfoGlBlitFramebuffer)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGenQueries{}) // interface compliance check
func (ϟa *GlGenQueries) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	queries_cnt := uint64(ϟa.Count)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{queries_cnt * 4 /* queries */})
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(outputs[0]) // queries
	ϟb.CallNoPush(funcInfoGlGenQueries)
	ϟa.Mutate(ϟs)
	for i, e := range ϟa.Queries {
		ptr := outputs[0].Offset(uint64(i * 4))
		if key, remap := e.remap(ϟa, ϟs); remap {
			storeRemap(ϟb, key, ptr, protocol.TypeUint32)
		}
	}
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGenQueries_Postback{}
			if err := postback.Decode(queries_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlBeginQuery{}) // interface compliance check
func (ϟa *GlBeginQuery) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Target))
	if key, remap := ϟa.Query.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Query.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Query.value(ϟb, ϟa, ϟs))
	}
	ϟb.CallNoPush(funcInfoGlBeginQuery)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlEndQuery{}) // interface compliance check
func (ϟa *GlEndQuery) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.CallNoPush(funcInfoGlEndQuery)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlDeleteQueries{}) // interface compliance check
func (ϟa *GlDeleteQueries) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Queries.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlDeleteQueries)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlIsQuery{}) // interface compliance check
func (ϟa *GlIsQuery) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{1 /* result */})
	if key, remap := ϟa.Query.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Query.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Query.value(ϟb, ϟa, ϟs))
	}
	ϟb.CallPush(funcInfoGlIsQuery)
	ϟb.Store(outputs[0])
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlIsQuery_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGetQueryiv{}) // interface compliance check
func (ϟa *GlGetQueryiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{4 /* value */})
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(outputs[0]) // value
	ϟb.CallNoPush(funcInfoGlGetQueryiv)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGetQueryiv_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGetQueryObjectuiv{}) // interface compliance check
func (ϟa *GlGetQueryObjectuiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{4 /* value */})
	if key, remap := ϟa.Query.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Query.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Query.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(outputs[0]) // value
	ϟb.CallNoPush(funcInfoGlGetQueryObjectuiv)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGetQueryObjectuiv_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGenQueriesEXT{}) // interface compliance check
func (ϟa *GlGenQueriesEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	queries_cnt := uint64(ϟa.Count)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{queries_cnt * 4 /* queries */})
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(outputs[0]) // queries
	ϟb.CallNoPush(funcInfoGlGenQueriesEXT)
	ϟa.Mutate(ϟs)
	for i, e := range ϟa.Queries {
		ptr := outputs[0].Offset(uint64(i * 4))
		if key, remap := e.remap(ϟa, ϟs); remap {
			storeRemap(ϟb, key, ptr, protocol.TypeUint32)
		}
	}
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGenQueriesEXT_Postback{}
			if err := postback.Decode(queries_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlBeginQueryEXT{}) // interface compliance check
func (ϟa *GlBeginQueryEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Target))
	if key, remap := ϟa.Query.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Query.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Query.value(ϟb, ϟa, ϟs))
	}
	ϟb.CallNoPush(funcInfoGlBeginQueryEXT)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlEndQueryEXT{}) // interface compliance check
func (ϟa *GlEndQueryEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.CallNoPush(funcInfoGlEndQueryEXT)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlDeleteQueriesEXT{}) // interface compliance check
func (ϟa *GlDeleteQueriesEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Queries.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoGlDeleteQueriesEXT)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlIsQueryEXT{}) // interface compliance check
func (ϟa *GlIsQueryEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{1 /* result */})
	if key, remap := ϟa.Query.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Query.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Query.value(ϟb, ϟa, ϟs))
	}
	ϟb.CallPush(funcInfoGlIsQueryEXT)
	ϟb.Store(outputs[0])
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlIsQueryEXT_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlQueryCounterEXT{}) // interface compliance check
func (ϟa *GlQueryCounterEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.Query.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Query.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Query.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.CallNoPush(funcInfoGlQueryCounterEXT)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGetQueryivEXT{}) // interface compliance check
func (ϟa *GlGetQueryivEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{4 /* value */})
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(outputs[0]) // value
	ϟb.CallNoPush(funcInfoGlGetQueryivEXT)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGetQueryivEXT_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGetQueryObjectivEXT{}) // interface compliance check
func (ϟa *GlGetQueryObjectivEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{4 /* value */})
	if key, remap := ϟa.Query.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Query.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Query.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(outputs[0]) // value
	ϟb.CallNoPush(funcInfoGlGetQueryObjectivEXT)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGetQueryObjectivEXT_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGetQueryObjectuivEXT{}) // interface compliance check
func (ϟa *GlGetQueryObjectuivEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{4 /* value */})
	if key, remap := ϟa.Query.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Query.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Query.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(outputs[0]) // value
	ϟb.CallNoPush(funcInfoGlGetQueryObjectuivEXT)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGetQueryObjectuivEXT_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGetQueryObjecti64vEXT{}) // interface compliance check
func (ϟa *GlGetQueryObjecti64vEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{8 /* value */})
	if key, remap := ϟa.Query.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Query.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Query.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(outputs[0]) // value
	ϟb.CallNoPush(funcInfoGlGetQueryObjecti64vEXT)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGetQueryObjecti64vEXT_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&GlGetQueryObjectui64vEXT{}) // interface compliance check
func (ϟa *GlGetQueryObjectui64vEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{8 /* value */})
	if key, remap := ϟa.Query.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.Query.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Query.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(outputs[0]) // value
	ϟb.CallNoPush(funcInfoGlGetQueryObjectui64vEXT)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := GlGetQueryObjectui64vEXT_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}
