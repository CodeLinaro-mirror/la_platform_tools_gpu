////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

package gles

import (
	"fmt"
	"strings"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/replay/builder"
	"android.googlesource.com/platform/tools/gpu/replay/protocol"
	"android.googlesource.com/platform/tools/gpu/replay/value"
)

func loadRemap(ϟb *builder.Builder, key interface{}, ty protocol.Type, val value.Value) {
	if ptr, found := ϟb.Remappings[key]; found {
		ϟb.Load(ty, ptr)
	} else {
		ptr = ϟb.AllocateMemory(uint64(ty.Size(ϟb.Architecture().PointerSize)))
		ϟb.Push(val) // We have an input to an unknown id, use the unmapped value.
		ϟb.Clone(0)
		ϟb.Store(ptr)
		ϟb.Remappings[key] = ptr
	}
}

var funcInfoEglInitialize = builder.FunctionInfo{ID: 0, ReturnType: protocol.TypeInt64, Parameters: 3}
var funcInfoEglCreateContext = builder.FunctionInfo{ID: 1, ReturnType: protocol.TypeVolatilePointer, Parameters: 4}
var funcInfoEglMakeCurrent = builder.FunctionInfo{ID: 2, ReturnType: protocol.TypeInt64, Parameters: 4}
var funcInfoEglSwapBuffers = builder.FunctionInfo{ID: 3, ReturnType: protocol.TypeInt64, Parameters: 2}
var funcInfoEglQuerySurface = builder.FunctionInfo{ID: 4, ReturnType: protocol.TypeInt64, Parameters: 4}
var funcInfoGlXCreateContext = builder.FunctionInfo{ID: 5, ReturnType: protocol.TypeVolatilePointer, Parameters: 4}
var funcInfoGlXCreateNewContext = builder.FunctionInfo{ID: 6, ReturnType: protocol.TypeVolatilePointer, Parameters: 5}
var funcInfoGlXMakeContextCurrent = builder.FunctionInfo{ID: 7, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlXSwapBuffers = builder.FunctionInfo{ID: 8, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoWglCreateContext = builder.FunctionInfo{ID: 9, ReturnType: protocol.TypeVolatilePointer, Parameters: 1}
var funcInfoWglCreateContextAttribsARB = builder.FunctionInfo{ID: 10, ReturnType: protocol.TypeVolatilePointer, Parameters: 3}
var funcInfoWglMakeCurrent = builder.FunctionInfo{ID: 11, ReturnType: protocol.TypeInt64, Parameters: 2}
var funcInfoWglSwapBuffers = builder.FunctionInfo{ID: 12, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCGLCreateContext = builder.FunctionInfo{ID: 13, ReturnType: protocol.TypeInt64, Parameters: 3}
var funcInfoCGLSetCurrentContext = builder.FunctionInfo{ID: 14, ReturnType: protocol.TypeInt64, Parameters: 1}
var funcInfoGlEnableClientState = builder.FunctionInfo{ID: 15, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlDisableClientState = builder.FunctionInfo{ID: 16, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlGetProgramBinaryOES = builder.FunctionInfo{ID: 17, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlProgramBinaryOES = builder.FunctionInfo{ID: 18, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlStartTilingQCOM = builder.FunctionInfo{ID: 19, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlEndTilingQCOM = builder.FunctionInfo{ID: 20, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlDiscardFramebufferEXT = builder.FunctionInfo{ID: 21, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlInsertEventMarkerEXT = builder.FunctionInfo{ID: 22, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlPushGroupMarkerEXT = builder.FunctionInfo{ID: 23, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlPopGroupMarkerEXT = builder.FunctionInfo{ID: 24, ReturnType: protocol.TypeVoid, Parameters: 0}
var funcInfoGlTexStorage1DEXT = builder.FunctionInfo{ID: 25, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlTexStorage2DEXT = builder.FunctionInfo{ID: 26, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlTexStorage3DEXT = builder.FunctionInfo{ID: 27, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlTextureStorage1DEXT = builder.FunctionInfo{ID: 28, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlTextureStorage2DEXT = builder.FunctionInfo{ID: 29, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlTextureStorage3DEXT = builder.FunctionInfo{ID: 30, ReturnType: protocol.TypeVoid, Parameters: 7}
var funcInfoGlGenVertexArraysOES = builder.FunctionInfo{ID: 31, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBindVertexArrayOES = builder.FunctionInfo{ID: 32, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlDeleteVertexArraysOES = builder.FunctionInfo{ID: 33, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlIsVertexArrayOES = builder.FunctionInfo{ID: 34, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlEGLImageTargetTexture2DOES = builder.FunctionInfo{ID: 35, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlEGLImageTargetRenderbufferStorageOES = builder.FunctionInfo{ID: 36, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetGraphicsResetStatusEXT = builder.FunctionInfo{ID: 37, ReturnType: protocol.TypeUint32, Parameters: 0}
var funcInfoGlBindAttribLocation = builder.FunctionInfo{ID: 38, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlBlendFunc = builder.FunctionInfo{ID: 39, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBlendFuncSeparate = builder.FunctionInfo{ID: 40, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlBlendEquation = builder.FunctionInfo{ID: 41, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlBlendEquationSeparate = builder.FunctionInfo{ID: 42, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBlendColor = builder.FunctionInfo{ID: 43, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlEnableVertexAttribArray = builder.FunctionInfo{ID: 44, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlDisableVertexAttribArray = builder.FunctionInfo{ID: 45, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlVertexAttribPointer = builder.FunctionInfo{ID: 46, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlGetActiveAttrib = builder.FunctionInfo{ID: 47, ReturnType: protocol.TypeVoid, Parameters: 7}
var funcInfoGlGetActiveUniform = builder.FunctionInfo{ID: 48, ReturnType: protocol.TypeVoid, Parameters: 7}
var funcInfoGlGetError = builder.FunctionInfo{ID: 49, ReturnType: protocol.TypeUint32, Parameters: 0}
var funcInfoGlGetProgramiv = builder.FunctionInfo{ID: 50, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetShaderiv = builder.FunctionInfo{ID: 51, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetUniformLocation = builder.FunctionInfo{ID: 52, ReturnType: protocol.TypeInt32, Parameters: 2}
var funcInfoGlGetAttribLocation = builder.FunctionInfo{ID: 53, ReturnType: protocol.TypeInt32, Parameters: 2}
var funcInfoGlPixelStorei = builder.FunctionInfo{ID: 54, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlTexParameteri = builder.FunctionInfo{ID: 55, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlTexParameterf = builder.FunctionInfo{ID: 56, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetTexParameteriv = builder.FunctionInfo{ID: 57, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetTexParameterfv = builder.FunctionInfo{ID: 58, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform1i = builder.FunctionInfo{ID: 59, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlUniform2i = builder.FunctionInfo{ID: 60, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform3i = builder.FunctionInfo{ID: 61, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUniform4i = builder.FunctionInfo{ID: 62, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlUniform1iv = builder.FunctionInfo{ID: 63, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform2iv = builder.FunctionInfo{ID: 64, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform3iv = builder.FunctionInfo{ID: 65, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform4iv = builder.FunctionInfo{ID: 66, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform1f = builder.FunctionInfo{ID: 67, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlUniform2f = builder.FunctionInfo{ID: 68, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform3f = builder.FunctionInfo{ID: 69, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUniform4f = builder.FunctionInfo{ID: 70, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlUniform1fv = builder.FunctionInfo{ID: 71, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform2fv = builder.FunctionInfo{ID: 72, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform3fv = builder.FunctionInfo{ID: 73, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform4fv = builder.FunctionInfo{ID: 74, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniformMatrix2fv = builder.FunctionInfo{ID: 75, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUniformMatrix3fv = builder.FunctionInfo{ID: 76, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUniformMatrix4fv = builder.FunctionInfo{ID: 77, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlGetUniformfv = builder.FunctionInfo{ID: 78, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetUniformiv = builder.FunctionInfo{ID: 79, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlVertexAttrib1f = builder.FunctionInfo{ID: 80, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlVertexAttrib2f = builder.FunctionInfo{ID: 81, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlVertexAttrib3f = builder.FunctionInfo{ID: 82, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlVertexAttrib4f = builder.FunctionInfo{ID: 83, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlVertexAttrib1fv = builder.FunctionInfo{ID: 84, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlVertexAttrib2fv = builder.FunctionInfo{ID: 85, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlVertexAttrib3fv = builder.FunctionInfo{ID: 86, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlVertexAttrib4fv = builder.FunctionInfo{ID: 87, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetShaderPrecisionFormat = builder.FunctionInfo{ID: 88, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlDepthMask = builder.FunctionInfo{ID: 89, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlDepthFunc = builder.FunctionInfo{ID: 90, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlDepthRangef = builder.FunctionInfo{ID: 91, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlColorMask = builder.FunctionInfo{ID: 92, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlStencilMask = builder.FunctionInfo{ID: 93, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlStencilMaskSeparate = builder.FunctionInfo{ID: 94, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlStencilFuncSeparate = builder.FunctionInfo{ID: 95, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlStencilOpSeparate = builder.FunctionInfo{ID: 96, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlFrontFace = builder.FunctionInfo{ID: 97, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlViewport = builder.FunctionInfo{ID: 98, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlScissor = builder.FunctionInfo{ID: 99, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlActiveTexture = builder.FunctionInfo{ID: 100, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlGenTextures = builder.FunctionInfo{ID: 101, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlDeleteTextures = builder.FunctionInfo{ID: 102, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlIsTexture = builder.FunctionInfo{ID: 103, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlBindTexture = builder.FunctionInfo{ID: 104, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlTexImage2D = builder.FunctionInfo{ID: 105, ReturnType: protocol.TypeVoid, Parameters: 9}
var funcInfoGlTexSubImage2D = builder.FunctionInfo{ID: 106, ReturnType: protocol.TypeVoid, Parameters: 9}
var funcInfoGlCopyTexImage2D = builder.FunctionInfo{ID: 107, ReturnType: protocol.TypeVoid, Parameters: 8}
var funcInfoGlCopyTexSubImage2D = builder.FunctionInfo{ID: 108, ReturnType: protocol.TypeVoid, Parameters: 8}
var funcInfoGlCompressedTexImage2D = builder.FunctionInfo{ID: 109, ReturnType: protocol.TypeVoid, Parameters: 8}
var funcInfoGlCompressedTexSubImage2D = builder.FunctionInfo{ID: 110, ReturnType: protocol.TypeVoid, Parameters: 9}
var funcInfoGlGenerateMipmap = builder.FunctionInfo{ID: 111, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlReadPixels = builder.FunctionInfo{ID: 112, ReturnType: protocol.TypeVoid, Parameters: 7}
var funcInfoGlGenFramebuffers = builder.FunctionInfo{ID: 113, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBindFramebuffer = builder.FunctionInfo{ID: 114, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlCheckFramebufferStatus = builder.FunctionInfo{ID: 115, ReturnType: protocol.TypeUint32, Parameters: 1}
var funcInfoGlDeleteFramebuffers = builder.FunctionInfo{ID: 116, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlIsFramebuffer = builder.FunctionInfo{ID: 117, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlGenRenderbuffers = builder.FunctionInfo{ID: 118, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBindRenderbuffer = builder.FunctionInfo{ID: 119, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlRenderbufferStorage = builder.FunctionInfo{ID: 120, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlDeleteRenderbuffers = builder.FunctionInfo{ID: 121, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlIsRenderbuffer = builder.FunctionInfo{ID: 122, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlGetRenderbufferParameteriv = builder.FunctionInfo{ID: 123, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGenBuffers = builder.FunctionInfo{ID: 124, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBindBuffer = builder.FunctionInfo{ID: 125, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBufferData = builder.FunctionInfo{ID: 126, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlBufferSubData = builder.FunctionInfo{ID: 127, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlDeleteBuffers = builder.FunctionInfo{ID: 128, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlIsBuffer = builder.FunctionInfo{ID: 129, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlGetBufferParameteriv = builder.FunctionInfo{ID: 130, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlCreateShader = builder.FunctionInfo{ID: 131, ReturnType: protocol.TypeUint32, Parameters: 1}
var funcInfoGlDeleteShader = builder.FunctionInfo{ID: 132, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlShaderSource = builder.FunctionInfo{ID: 133, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlShaderBinary = builder.FunctionInfo{ID: 134, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlGetShaderInfoLog = builder.FunctionInfo{ID: 135, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlGetShaderSource = builder.FunctionInfo{ID: 136, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlReleaseShaderCompiler = builder.FunctionInfo{ID: 137, ReturnType: protocol.TypeVoid, Parameters: 0}
var funcInfoGlCompileShader = builder.FunctionInfo{ID: 138, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlIsShader = builder.FunctionInfo{ID: 139, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlCreateProgram = builder.FunctionInfo{ID: 140, ReturnType: protocol.TypeUint32, Parameters: 0}
var funcInfoGlDeleteProgram = builder.FunctionInfo{ID: 141, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlAttachShader = builder.FunctionInfo{ID: 142, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlDetachShader = builder.FunctionInfo{ID: 143, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetAttachedShaders = builder.FunctionInfo{ID: 144, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlLinkProgram = builder.FunctionInfo{ID: 145, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlGetProgramInfoLog = builder.FunctionInfo{ID: 146, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUseProgram = builder.FunctionInfo{ID: 147, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlIsProgram = builder.FunctionInfo{ID: 148, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlValidateProgram = builder.FunctionInfo{ID: 149, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlClearColor = builder.FunctionInfo{ID: 150, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlClearDepthf = builder.FunctionInfo{ID: 151, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlClearStencil = builder.FunctionInfo{ID: 152, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlClear = builder.FunctionInfo{ID: 153, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlCullFace = builder.FunctionInfo{ID: 154, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlPolygonOffset = builder.FunctionInfo{ID: 155, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlLineWidth = builder.FunctionInfo{ID: 156, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlSampleCoverage = builder.FunctionInfo{ID: 157, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlHint = builder.FunctionInfo{ID: 158, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlFramebufferRenderbuffer = builder.FunctionInfo{ID: 159, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlFramebufferTexture2D = builder.FunctionInfo{ID: 160, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlGetFramebufferAttachmentParameteriv = builder.FunctionInfo{ID: 161, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlDrawElements = builder.FunctionInfo{ID: 162, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlDrawArrays = builder.FunctionInfo{ID: 163, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlFlush = builder.FunctionInfo{ID: 164, ReturnType: protocol.TypeVoid, Parameters: 0}
var funcInfoGlFinish = builder.FunctionInfo{ID: 165, ReturnType: protocol.TypeVoid, Parameters: 0}
var funcInfoGlGetBooleanv = builder.FunctionInfo{ID: 166, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetFloatv = builder.FunctionInfo{ID: 167, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetIntegerv = builder.FunctionInfo{ID: 168, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetString = builder.FunctionInfo{ID: 169, ReturnType: protocol.TypeVolatilePointer, Parameters: 1}
var funcInfoGlEnable = builder.FunctionInfo{ID: 170, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlDisable = builder.FunctionInfo{ID: 171, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlIsEnabled = builder.FunctionInfo{ID: 172, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlMapBufferRange = builder.FunctionInfo{ID: 173, ReturnType: protocol.TypeVolatilePointer, Parameters: 4}
var funcInfoGlUnmapBuffer = builder.FunctionInfo{ID: 174, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlInvalidateFramebuffer = builder.FunctionInfo{ID: 175, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlRenderbufferStorageMultisample = builder.FunctionInfo{ID: 176, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlBlitFramebuffer = builder.FunctionInfo{ID: 177, ReturnType: protocol.TypeVoid, Parameters: 10}
var funcInfoGlGenQueries = builder.FunctionInfo{ID: 178, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBeginQuery = builder.FunctionInfo{ID: 179, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlEndQuery = builder.FunctionInfo{ID: 180, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlDeleteQueries = builder.FunctionInfo{ID: 181, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlIsQuery = builder.FunctionInfo{ID: 182, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlGetQueryiv = builder.FunctionInfo{ID: 183, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetQueryObjectuiv = builder.FunctionInfo{ID: 184, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGenQueriesEXT = builder.FunctionInfo{ID: 185, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBeginQueryEXT = builder.FunctionInfo{ID: 186, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlEndQueryEXT = builder.FunctionInfo{ID: 187, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlDeleteQueriesEXT = builder.FunctionInfo{ID: 188, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlIsQueryEXT = builder.FunctionInfo{ID: 189, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlQueryCounterEXT = builder.FunctionInfo{ID: 190, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetQueryivEXT = builder.FunctionInfo{ID: 191, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetQueryObjectivEXT = builder.FunctionInfo{ID: 192, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetQueryObjectuivEXT = builder.FunctionInfo{ID: 193, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetQueryObjecti64vEXT = builder.FunctionInfo{ID: 194, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetQueryObjectui64vEXT = builder.FunctionInfo{ID: 195, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoArchitecture = builder.FunctionInfo{ID: 196, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoReplayCreateRenderer = builder.FunctionInfo{ID: 197, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoReplayBindRenderer = builder.FunctionInfo{ID: 198, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoBackbufferInfo = builder.FunctionInfo{ID: 199, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoStartTimer = builder.FunctionInfo{ID: 200, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoStopTimer = builder.FunctionInfo{ID: 201, ReturnType: protocol.TypeUint64, Parameters: 1}
var funcInfoFlushPostBuffer = builder.FunctionInfo{ID: 202, ReturnType: protocol.TypeVoid, Parameters: 0}

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
	return value.S32(int32(c))
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
	return Voidᵖ(c).value()
}
func (c EGLContext) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return Voidᵖ(c).value()
}
func (c EGLDisplay) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return Voidᵖ(c).value()
}
func (c EGLSurface) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return Voidᵖ(c).value()
}
func (c GLXContext) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return Voidᵖ(c).value()
}
func (c GLXDrawable) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return Voidᵖ(c).value()
}
func (c HGLRC) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return Voidᵖ(c).value()
}
func (c HDC) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return Voidᵖ(c).value()
}
func (c BOOL) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.S64(int64(c))
}
func (c CGLError) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.S64(int64(c))
}
func (c CGLPixelFormatObj) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return Voidᵖ(c).value()
}
func (c CGLContextObj) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return Voidᵖ(c).value()
}

var _ = replay.Replayer(&GlEnableClientState{}) // interface compliance check
func (ϟa *GlEnableClientState) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_10_result := context              // Contextʳ
	ctx := GetContext_10_result                  // Contextʳ
	ctx.Capabilities[Capability(ϟa.Type)] = true
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Call(funcInfoGlEnableClientState)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_10_result, ctx
	return nil
}

var _ = replay.Replayer(&GlDisableClientState{}) // interface compliance check
func (ϟa *GlDisableClientState) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_11_result := context              // Contextʳ
	ctx := GetContext_11_result                  // Contextʳ
	ctx.Capabilities[Capability(ϟa.Type)] = false
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Call(funcInfoGlDisableClientState)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_11_result, ctx
	return nil
}

var _ = replay.Replayer(&GlGetProgramBinaryOES{}) // interface compliance check
func (ϟa *GlGetProgramBinaryOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.BufferSize))
	ϟb.Push(ϟa.BytesWritten.value())
	ϟb.Push(ϟa.BinaryFormat.value())
	ϟb.Push(ϟa.Binary.value())
	ϟb.Call(funcInfoGlGetProgramBinaryOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	l := int32(ϟa.BytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)) // s32
	ϟa.BytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(l, ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.BinaryFormat.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(ϟa.BinaryFormat.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Binary.Slice(uint64(int32(0)), uint64(l), ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	_ = l
	return nil
}

var _ = replay.Replayer(&GlProgramBinaryOES{}) // interface compliance check
func (ϟa *GlProgramBinaryOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.BinaryFormat))
	ϟb.Push(ϟa.Binary.value())
	ϟb.Push(value.S32(ϟa.BinarySize))
	ϟb.Call(funcInfoGlProgramBinaryOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlStartTilingQCOM{}) // interface compliance check
func (ϟa *GlStartTilingQCOM) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.S32(ϟa.X))
	ϟb.Push(value.S32(ϟa.Y))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.Push(value.U32(ϟa.PreserveMask))
	ϟb.Call(funcInfoGlStartTilingQCOM)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlEndTilingQCOM{}) // interface compliance check
func (ϟa *GlEndTilingQCOM) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U32(ϟa.PreserveMask))
	ϟb.Call(funcInfoGlEndTilingQCOM)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlDiscardFramebufferEXT{}) // interface compliance check
func (ϟa *GlDiscardFramebufferEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.NumAttachments))
	ϟb.Push(ϟa.Attachments.value())
	ϟb.Call(funcInfoGlDiscardFramebufferEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlInsertEventMarkerEXT{}) // interface compliance check
func (ϟa *GlInsertEventMarkerEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	if (ϟa.Length) > (int32(0)) {
		_ = strings.TrimRight(string(ϟa.Marker.StringSlice(ϟs, ϟd, ϟl, true).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)), "\x00")
	} else {
		ϟa.Marker.Slice(uint64(int32(0)), uint64(ϟa.Length), ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	}
	ϟb.Push(value.S32(ϟa.Length))
	ϟb.Push(ϟa.Marker.value())
	ϟb.Call(funcInfoGlInsertEventMarkerEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlPushGroupMarkerEXT{}) // interface compliance check
func (ϟa *GlPushGroupMarkerEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	if (ϟa.Length) > (int32(0)) {
		_ = strings.TrimRight(string(ϟa.Marker.StringSlice(ϟs, ϟd, ϟl, true).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)), "\x00")
	} else {
		ϟa.Marker.Slice(uint64(int32(0)), uint64(ϟa.Length), ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	}
	ϟb.Push(value.S32(ϟa.Length))
	ϟb.Push(ϟa.Marker.value())
	ϟb.Call(funcInfoGlPushGroupMarkerEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlPopGroupMarkerEXT{}) // interface compliance check
func (ϟa *GlPopGroupMarkerEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoGlPopGroupMarkerEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlTexStorage1DEXT{}) // interface compliance check
func (ϟa *GlTexStorage1DEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Levels))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Call(funcInfoGlTexStorage1DEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlTexStorage2DEXT{}) // interface compliance check
func (ϟa *GlTexStorage2DEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Levels))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.Call(funcInfoGlTexStorage2DEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlTexStorage3DEXT{}) // interface compliance check
func (ϟa *GlTexStorage3DEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Levels))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.Push(value.S32(ϟa.Depth))
	ϟb.Call(funcInfoGlTexStorage3DEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlTextureStorage1DEXT{}) // interface compliance check
func (ϟa *GlTextureStorage1DEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	if key, remap := ϟa.Texture.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Texture.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Texture.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Levels))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Call(funcInfoGlTextureStorage1DEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlTextureStorage2DEXT{}) // interface compliance check
func (ϟa *GlTextureStorage2DEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	if key, remap := ϟa.Texture.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Texture.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Texture.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Levels))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.Call(funcInfoGlTextureStorage2DEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlTextureStorage3DEXT{}) // interface compliance check
func (ϟa *GlTextureStorage3DEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	if key, remap := ϟa.Texture.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Texture.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Texture.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Levels))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.Push(value.S32(ϟa.Depth))
	ϟb.Call(funcInfoGlTextureStorage3DEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlGenVertexArraysOES{}) // interface compliance check
func (ϟa *GlGenVertexArraysOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	a := ϟa.Arrays.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // VertexArrayIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                 // Contextʳ
	GetContext_12_result := context                              // Contextʳ
	ctx := GetContext_12_result                                  // Contextʳ
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Arrays.value())
	ϟb.Call(funcInfoGlGenVertexArraysOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		id := VertexArrayId(ϟa.Arrays.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)) // VertexArrayId
		ctx.Instances.VertexArrays[id] = func() *VertexArray {
			s := &VertexArray{}
			s.Init()
			return s
		}()
		a.Index(uint64(i), ϟs).replayWrite(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _ = a, context, GetContext_12_result, ctx
	return nil
}

var _ = replay.Replayer(&GlBindVertexArrayOES{}) // interface compliance check
func (ϟa *GlBindVertexArrayOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	if key, remap := ϟa.Array.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Array.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Array.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlBindVertexArrayOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_13_result, ctx
	return nil
}

var _ = replay.Replayer(&GlDeleteVertexArraysOES{}) // interface compliance check
func (ϟa *GlDeleteVertexArraysOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                 // Contextʳ
	GetContext_14_result := context                              // Contextʳ
	ctx := GetContext_14_result                                  // Contextʳ
	a := ϟa.Arrays.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // VertexArrayIdˢ
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		ctx.Instances.VertexArrays[a.Index(uint64(i), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)] = (*VertexArray)(nil)
	}
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Arrays.value())
	ϟb.Call(funcInfoGlDeleteVertexArraysOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = context, GetContext_14_result, ctx, a
	return nil
}

var _ = replay.Replayer(&GlIsVertexArrayOES{}) // interface compliance check
func (ϟa *GlIsVertexArrayOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_15_result := context              // Contextʳ
	ctx := GetContext_15_result                  // Contextʳ
	if key, remap := ϟa.Array.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Array.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Array.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlIsVertexArrayOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_15_result, ctx
	return nil
}

var _ = replay.Replayer(&GlEGLImageTargetTexture2DOES{}) // interface compliance check
func (ϟa *GlEGLImageTargetTexture2DOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Image.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlEGLImageTargetTexture2DOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlEGLImageTargetRenderbufferStorageOES{}) // interface compliance check
func (ϟa *GlEGLImageTargetRenderbufferStorageOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Image.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlEGLImageTargetRenderbufferStorageOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlGetGraphicsResetStatusEXT{}) // interface compliance check
func (ϟa *GlGetGraphicsResetStatusEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoGlGetGraphicsResetStatusEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlBindAttribLocation{}) // interface compliance check
func (ϟa *GlBindAttribLocation) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_16_result := context              // Contextʳ
	ctx := GetContext_16_result                  // Contextʳ
	p := ctx.Instances.Programs.Get(ϟa.Program)  // Programʳ
	p.AttributeBindings[ϟa.Name] = ϟa.Location
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟb.String(ϟa.Name))
	ϟb.Call(funcInfoGlBindAttribLocation)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = context, GetContext_16_result, ctx, p
	return nil
}

var _ = replay.Replayer(&GlBlendFunc{}) // interface compliance check
func (ϟa *GlBlendFunc) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	ϟb.Push(value.U32(ϟa.SrcFactor))
	ϟb.Push(value.U32(ϟa.DstFactor))
	ϟb.Call(funcInfoGlBlendFunc)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_17_result, ctx
	return nil
}

var _ = replay.Replayer(&GlBlendFuncSeparate{}) // interface compliance check
func (ϟa *GlBlendFuncSeparate) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	ϟb.Push(value.U32(ϟa.SrcFactorRgb))
	ϟb.Push(value.U32(ϟa.DstFactorRgb))
	ϟb.Push(value.U32(ϟa.SrcFactorAlpha))
	ϟb.Push(value.U32(ϟa.DstFactorAlpha))
	ϟb.Call(funcInfoGlBlendFuncSeparate)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_18_result, ctx
	return nil
}

var _ = replay.Replayer(&GlBlendEquation{}) // interface compliance check
func (ϟa *GlBlendEquation) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_19_result := context              // Contextʳ
	ctx := GetContext_19_result                  // Contextʳ
	ctx.Blending.BlendEquationRgb = ϟa.Equation
	ctx.Blending.BlendEquationAlpha = ϟa.Equation
	ϟb.Push(value.U32(ϟa.Equation))
	ϟb.Call(funcInfoGlBlendEquation)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_19_result, ctx
	return nil
}

var _ = replay.Replayer(&GlBlendEquationSeparate{}) // interface compliance check
func (ϟa *GlBlendEquationSeparate) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_20_result := context              // Contextʳ
	ctx := GetContext_20_result                  // Contextʳ
	ctx.Blending.BlendEquationRgb = ϟa.Rgb
	ctx.Blending.BlendEquationAlpha = ϟa.Alpha
	ϟb.Push(value.U32(ϟa.Rgb))
	ϟb.Push(value.U32(ϟa.Alpha))
	ϟb.Call(funcInfoGlBlendEquationSeparate)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_20_result, ctx
	return nil
}

var _ = replay.Replayer(&GlBlendColor{}) // interface compliance check
func (ϟa *GlBlendColor) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	ϟb.Push(value.F32(ϟa.Red))
	ϟb.Push(value.F32(ϟa.Green))
	ϟb.Push(value.F32(ϟa.Blue))
	ϟb.Push(value.F32(ϟa.Alpha))
	ϟb.Call(funcInfoGlBlendColor)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_21_result, ctx
	return nil
}

var _ = replay.Replayer(&GlEnableVertexAttribArray{}) // interface compliance check
func (ϟa *GlEnableVertexAttribArray) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_22_result := context              // Contextʳ
	ctx := GetContext_22_result                  // Contextʳ
	ctx.VertexAttributeArrays.Get(ϟa.Location).Enabled = true
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlEnableVertexAttribArray)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_22_result, ctx
	return nil
}

var _ = replay.Replayer(&GlDisableVertexAttribArray{}) // interface compliance check
func (ϟa *GlDisableVertexAttribArray) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_23_result := context              // Contextʳ
	ctx := GetContext_23_result                  // Contextʳ
	ctx.VertexAttributeArrays.Get(ϟa.Location).Enabled = false
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlDisableVertexAttribArray)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_23_result, ctx
	return nil
}

var _ = replay.Replayer(&GlVertexAttribPointer{}) // interface compliance check
func (ϟa *GlVertexAttribPointer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.S32(ϟa.Size))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(value.Bool(ϟa.Normalized))
	ϟb.Push(value.S32(ϟa.Stride))
	ϟb.Push(ϟa.Data.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlVertexAttribPointer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = context, GetContext_24_result, ctx, a
	return nil
}

var _ = replay.Replayer(&GlGetActiveAttrib{}) // interface compliance check
func (ϟa *GlGetActiveAttrib) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.S32(ϟa.BufferSize))
	ϟb.Push(ϟa.BufferBytesWritten.value())
	ϟb.Push(ϟa.VectorCount.value())
	ϟb.Push(ϟa.Type.value())
	ϟb.Push(ϟa.Name.value())
	ϟb.Call(funcInfoGlGetActiveAttrib)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	l := int32(ϟa.BufferBytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)) // s32
	ϟa.BufferBytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(l, ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.VectorCount.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(int32(ϟa.VectorCount.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Type.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(ShaderAttribType(ϟa.Type.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Name.Slice(uint64(int32(0)), uint64(l), ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	_ = l
	return nil
}

var _ = replay.Replayer(&GlGetActiveUniform{}) // interface compliance check
func (ϟa *GlGetActiveUniform) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Location))
	ϟb.Push(value.S32(ϟa.BufferSize))
	ϟb.Push(ϟa.BufferBytesWritten.value())
	ϟb.Push(ϟa.VectorCount.value())
	ϟb.Push(ϟa.Type.value())
	ϟb.Push(ϟa.Name.value())
	ϟb.Call(funcInfoGlGetActiveUniform)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	l := int32(ϟa.BufferBytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)) // s32
	ϟa.BufferBytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(l, ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.VectorCount.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(int32(ϟa.VectorCount.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Type.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(ShaderUniformType(ϟa.Type.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Name.Slice(uint64(int32(0)), uint64(l), ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	_ = l
	return nil
}

var _ = replay.Replayer(&GlGetError{}) // interface compliance check
func (ϟa *GlGetError) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoGlGetError)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlGetProgramiv{}) // interface compliance check
func (ϟa *GlGetProgramiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlGetProgramiv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

var _ = replay.Replayer(&GlGetShaderiv{}) // interface compliance check
func (ϟa *GlGetShaderiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_25_result := context              // Contextʳ
	ctx := GetContext_25_result                  // Contextʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // Shaderʳ
	if key, remap := ϟa.Shader.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Shader.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Shader.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlGetShaderiv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(func() (result int32) {
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
	}(), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _ = context, GetContext_25_result, ctx, s
	return nil
}

var _ = replay.Replayer(&GlGetUniformLocation{}) // interface compliance check
func (ϟa *GlGetUniformLocation) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟb.String(ϟa.Name))
	ϟb.Call(funcInfoGlGetUniformLocation)
	if key, remap := ϟa.Result.remap(ϟa, ϟs); remap {
		if ptr, found := ϟb.Remappings[key]; !found {
			ptr = ϟb.AllocateMemory(uint64(4))
			ϟb.Clone(0)
			ϟb.Store(ptr)
			ϟb.Remappings[key] = ptr
		}
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlGetAttribLocation{}) // interface compliance check
func (ϟa *GlGetAttribLocation) defaultReplay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟb.String(ϟa.Name))
	ϟb.Call(funcInfoGlGetAttribLocation)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlPixelStorei{}) // interface compliance check
func (ϟa *GlPixelStorei) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_26_result := context              // Contextʳ
	ctx := GetContext_26_result                  // Contextʳ
	ctx.PixelStorage[ϟa.Parameter] = ϟa.Value
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(value.S32(ϟa.Value))
	ϟb.Call(funcInfoGlPixelStorei)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_26_result, ctx
	return nil
}

var _ = replay.Replayer(&GlTexParameteri{}) // interface compliance check
func (ϟa *GlTexParameteri) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(value.S32(ϟa.Value))
	ϟb.Call(funcInfoGlTexParameteri)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = context, GetContext_27_result, ctx, id, t
	return nil
}

var _ = replay.Replayer(&GlTexParameterf{}) // interface compliance check
func (ϟa *GlTexParameterf) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(value.F32(ϟa.Value))
	ϟb.Call(funcInfoGlTexParameterf)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = context, GetContext_28_result, ctx, id, t
	return nil
}

var _ = replay.Replayer(&GlGetTexParameteriv{}) // interface compliance check
func (ϟa *GlGetTexParameteriv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                     // Contextʳ
	GetContext_29_result := context                                  // Contextʳ
	ctx := GetContext_29_result                                      // Contextʳ
	id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(ϟa.Target) // TextureId
	t := ctx.Instances.Textures.Get(id)                              // Textureʳ
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(ϟa.Values.value())
	ϟb.Call(funcInfoGlGetTexParameteriv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Values.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(func() (result int32) {
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
	}(), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _ = context, GetContext_29_result, ctx, id, t
	return nil
}

var _ = replay.Replayer(&GlGetTexParameterfv{}) // interface compliance check
func (ϟa *GlGetTexParameterfv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                     // Contextʳ
	GetContext_30_result := context                                  // Contextʳ
	ctx := GetContext_30_result                                      // Contextʳ
	id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(ϟa.Target) // TextureId
	t := ctx.Instances.Textures.Get(id)                              // Textureʳ
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(ϟa.Values.value())
	ϟb.Call(funcInfoGlGetTexParameterfv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Values.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(func() (result float32) {
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
	}(), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _ = context, GetContext_30_result, ctx, id, t
	return nil
}

var _ = replay.Replayer(&GlUniform1i{}) // interface compliance check
func (ϟa *GlUniform1i) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Value))
	ϟb.Call(funcInfoGlUniform1i)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = context, GetContext_31_result, ctx, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniform2i{}) // interface compliance check
func (ϟa *GlUniform2i) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Value0))
	ϟb.Push(value.S32(ϟa.Value1))
	ϟb.Call(funcInfoGlUniform2i)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = context, GetContext_32_result, ctx, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniform3i{}) // interface compliance check
func (ϟa *GlUniform3i) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Value0))
	ϟb.Push(value.S32(ϟa.Value1))
	ϟb.Push(value.S32(ϟa.Value2))
	ϟb.Call(funcInfoGlUniform3i)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = context, GetContext_33_result, ctx, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniform4i{}) // interface compliance check
func (ϟa *GlUniform4i) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Value0))
	ϟb.Push(value.S32(ϟa.Value1))
	ϟb.Push(value.S32(ϟa.Value2))
	ϟb.Push(value.S32(ϟa.Value3))
	ϟb.Call(funcInfoGlUniform4i)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = context, GetContext_34_result, ctx, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniform1iv{}) // interface compliance check
func (ϟa *GlUniform1iv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)            // Contextʳ
	GetContext_35_result := context                         // Contextʳ
	ctx := GetContext_35_result                             // Contextʳ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = ShaderUniformType_GL_INT
	uniform.Value.S32 = ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlUniform1iv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = context, GetContext_35_result, ctx, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniform2iv{}) // interface compliance check
func (ϟa *GlUniform2iv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
		s.X = v.Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		s.Y = v.Index(uint64(1), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlUniform2iv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = context, GetContext_36_result, ctx, v, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniform3iv{}) // interface compliance check
func (ϟa *GlUniform3iv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
		s.X = v.Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		s.Y = v.Index(uint64(1), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		s.Z = v.Index(uint64(2), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlUniform3iv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = context, GetContext_37_result, ctx, v, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniform4iv{}) // interface compliance check
func (ϟa *GlUniform4iv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
		s.X = v.Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		s.Y = v.Index(uint64(1), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		s.Z = v.Index(uint64(2), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		s.W = v.Index(uint64(3), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlUniform4iv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = context, GetContext_38_result, ctx, v, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniform1f{}) // interface compliance check
func (ϟa *GlUniform1f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.F32(ϟa.Value))
	ϟb.Call(funcInfoGlUniform1f)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = context, GetContext_39_result, ctx, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniform2f{}) // interface compliance check
func (ϟa *GlUniform2f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.F32(ϟa.Value0))
	ϟb.Push(value.F32(ϟa.Value1))
	ϟb.Call(funcInfoGlUniform2f)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = context, GetContext_40_result, ctx, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniform3f{}) // interface compliance check
func (ϟa *GlUniform3f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.F32(ϟa.Value0))
	ϟb.Push(value.F32(ϟa.Value1))
	ϟb.Push(value.F32(ϟa.Value2))
	ϟb.Call(funcInfoGlUniform3f)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = context, GetContext_41_result, ctx, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniform4f{}) // interface compliance check
func (ϟa *GlUniform4f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.F32(ϟa.Value0))
	ϟb.Push(value.F32(ϟa.Value1))
	ϟb.Push(value.F32(ϟa.Value2))
	ϟb.Push(value.F32(ϟa.Value3))
	ϟb.Call(funcInfoGlUniform4f)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = context, GetContext_42_result, ctx, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniform1fv{}) // interface compliance check
func (ϟa *GlUniform1fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                // Contextʳ
	GetContext_43_result := context                             // Contextʳ
	ctx := GetContext_43_result                                 // Contextʳ
	v := ϟa.Value.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // F32ˢ
	v.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT
	uniform.Value.F32 = v.Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlUniform1fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = context, GetContext_43_result, ctx, v, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniform2fv{}) // interface compliance check
func (ϟa *GlUniform2fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                             // Contextʳ
	GetContext_44_result := context                                          // Contextʳ
	ctx := GetContext_44_result                                              // Contextʳ
	v := ϟa.Value.Slice(uint64(int32(0)), uint64((ϟa.Count)*(int32(2))), ϟs) // F32ˢ
	v.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT_VEC2
	uniform.Value.Vec2f = func() Vec2f {
		s := Vec2f{}
		s.Init()
		s.X = v.Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		s.Y = v.Index(uint64(1), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlUniform2fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = context, GetContext_44_result, ctx, v, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniform3fv{}) // interface compliance check
func (ϟa *GlUniform3fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                             // Contextʳ
	GetContext_45_result := context                                          // Contextʳ
	ctx := GetContext_45_result                                              // Contextʳ
	v := ϟa.Value.Slice(uint64(int32(0)), uint64((ϟa.Count)*(int32(3))), ϟs) // F32ˢ
	v.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT_VEC3
	uniform.Value.Vec3f = func() Vec3f {
		s := Vec3f{}
		s.Init()
		s.X = v.Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		s.Y = v.Index(uint64(1), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		s.Z = v.Index(uint64(2), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlUniform3fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = context, GetContext_45_result, ctx, v, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniform4fv{}) // interface compliance check
func (ϟa *GlUniform4fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                             // Contextʳ
	GetContext_46_result := context                                          // Contextʳ
	ctx := GetContext_46_result                                              // Contextʳ
	v := ϟa.Value.Slice(uint64(int32(0)), uint64((ϟa.Count)*(int32(4))), ϟs) // F32ˢ
	v.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT_VEC4
	uniform.Value.Vec4f = func() Vec4f {
		s := Vec4f{}
		s.Init()
		s.X = v.Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		s.Y = v.Index(uint64(1), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		s.Z = v.Index(uint64(2), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		s.W = v.Index(uint64(3), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlUniform4fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = context, GetContext_46_result, ctx, v, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniformMatrix2fv{}) // interface compliance check
func (ϟa *GlUniformMatrix2fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                              // Contextʳ
	GetContext_47_result := context                                           // Contextʳ
	ctx := GetContext_47_result                                               // Contextʳ
	v := ϟa.Values.Slice(uint64(int32(0)), uint64((ϟa.Count)*(int32(4))), ϟs) // F32ˢ
	v.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT_MAT2
	uniform.Value.Mat2f = func() Mat2f {
		s := Mat2f{}
		s.Init()
		s.Col0 = func() Vec2f {
			s := Vec2f{}
			s.Init()
			s.X = v.Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
			s.Y = v.Index(uint64(1), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
			return s
		}()
		s.Col1 = func() Vec2f {
			s := Vec2f{}
			s.Init()
			s.X = v.Index(uint64(3), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
			s.Y = v.Index(uint64(4), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
			return s
		}()
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(value.Bool(ϟa.Transpose))
	ϟb.Push(ϟa.Values.value())
	ϟb.Call(funcInfoGlUniformMatrix2fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = context, GetContext_47_result, ctx, v, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniformMatrix3fv{}) // interface compliance check
func (ϟa *GlUniformMatrix3fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                              // Contextʳ
	GetContext_48_result := context                                           // Contextʳ
	ctx := GetContext_48_result                                               // Contextʳ
	v := ϟa.Values.Slice(uint64(int32(0)), uint64((ϟa.Count)*(int32(9))), ϟs) // F32ˢ
	v.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = ShaderUniformType_GL_FLOAT_MAT3
	uniform.Value.Mat3f = func() Mat3f {
		s := Mat3f{}
		s.Init()
		s.Col0 = func() Vec3f {
			s := Vec3f{}
			s.Init()
			s.X = v.Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
			s.Y = v.Index(uint64(1), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
			s.Z = v.Index(uint64(2), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
			return s
		}()
		s.Col1 = func() Vec3f {
			s := Vec3f{}
			s.Init()
			s.X = v.Index(uint64(3), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
			s.Y = v.Index(uint64(4), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
			s.Z = v.Index(uint64(5), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
			return s
		}()
		s.Col2 = func() Vec3f {
			s := Vec3f{}
			s.Init()
			s.X = v.Index(uint64(6), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
			s.Y = v.Index(uint64(7), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
			s.Z = v.Index(uint64(8), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
			return s
		}()
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(value.Bool(ϟa.Transpose))
	ϟb.Push(ϟa.Values.value())
	ϟb.Call(funcInfoGlUniformMatrix3fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = context, GetContext_48_result, ctx, v, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniformMatrix4fv{}) // interface compliance check
func (ϟa *GlUniformMatrix4fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                               // Contextʳ
	GetContext_49_result := context                                            // Contextʳ
	ctx := GetContext_49_result                                                // Contextʳ
	v := ϟa.Values.Slice(uint64(int32(0)), uint64((ϟa.Count)*(int32(16))), ϟs) // F32ˢ
	v.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Value.Mat4f = func() Mat4f {
		s := Mat4f{}
		s.Init()
		s.Col0 = func() Vec4f {
			s := Vec4f{}
			s.Init()
			s.X = v.Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
			s.Y = v.Index(uint64(1), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
			s.Z = v.Index(uint64(2), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
			s.W = v.Index(uint64(3), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
			return s
		}()
		s.Col1 = func() Vec4f {
			s := Vec4f{}
			s.Init()
			s.X = v.Index(uint64(4), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
			s.Y = v.Index(uint64(5), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
			s.Z = v.Index(uint64(6), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
			s.W = v.Index(uint64(7), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
			return s
		}()
		s.Col2 = func() Vec4f {
			s := Vec4f{}
			s.Init()
			s.X = v.Index(uint64(8), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
			s.Y = v.Index(uint64(9), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
			s.Z = v.Index(uint64(10), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
			s.W = v.Index(uint64(11), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
			return s
		}()
		s.Col3 = func() Vec4f {
			s := Vec4f{}
			s.Init()
			s.X = v.Index(uint64(12), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
			s.Y = v.Index(uint64(13), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
			s.Z = v.Index(uint64(14), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
			s.W = v.Index(uint64(15), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
			return s
		}()
		return s
	}()
	program.Uniforms[ϟa.Location] = uniform
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(value.Bool(ϟa.Transpose))
	ϟb.Push(ϟa.Values.value())
	ϟb.Call(funcInfoGlUniformMatrix4fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = context, GetContext_49_result, ctx, v, program, uniform
	return nil
}

var _ = replay.Replayer(&GlGetUniformfv{}) // interface compliance check
func (ϟa *GlGetUniformfv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Values.value())
	ϟb.Call(funcInfoGlGetUniformfv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlGetUniformiv{}) // interface compliance check
func (ϟa *GlGetUniformiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Values.value())
	ϟb.Call(funcInfoGlGetUniformiv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlVertexAttrib1f{}) // interface compliance check
func (ϟa *GlVertexAttrib1f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.F32(ϟa.Value0))
	ϟb.Call(funcInfoGlVertexAttrib1f)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlVertexAttrib2f{}) // interface compliance check
func (ϟa *GlVertexAttrib2f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.F32(ϟa.Value0))
	ϟb.Push(value.F32(ϟa.Value1))
	ϟb.Call(funcInfoGlVertexAttrib2f)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlVertexAttrib3f{}) // interface compliance check
func (ϟa *GlVertexAttrib3f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.F32(ϟa.Value0))
	ϟb.Push(value.F32(ϟa.Value1))
	ϟb.Push(value.F32(ϟa.Value2))
	ϟb.Call(funcInfoGlVertexAttrib3f)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlVertexAttrib4f{}) // interface compliance check
func (ϟa *GlVertexAttrib4f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.F32(ϟa.Value0))
	ϟb.Push(value.F32(ϟa.Value1))
	ϟb.Push(value.F32(ϟa.Value2))
	ϟb.Push(value.F32(ϟa.Value3))
	ϟb.Call(funcInfoGlVertexAttrib4f)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlVertexAttrib1fv{}) // interface compliance check
func (ϟa *GlVertexAttrib1fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlVertexAttrib1fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlVertexAttrib2fv{}) // interface compliance check
func (ϟa *GlVertexAttrib2fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(2), ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlVertexAttrib2fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlVertexAttrib3fv{}) // interface compliance check
func (ϟa *GlVertexAttrib3fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(3), ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlVertexAttrib3fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlVertexAttrib4fv{}) // interface compliance check
func (ϟa *GlVertexAttrib4fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(4), ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlVertexAttrib4fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlGetShaderPrecisionFormat{}) // interface compliance check
func (ϟa *GlGetShaderPrecisionFormat) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U32(ϟa.ShaderType))
	ϟb.Push(value.U32(ϟa.PrecisionType))
	ϟb.Push(ϟa.Range.value())
	ϟb.Push(ϟa.Precision.value())
	ϟb.Call(funcInfoGlGetShaderPrecisionFormat)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Range.Slice(uint64(0), uint64(2), ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Precision.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(ϟa.Precision.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

var _ = replay.Replayer(&GlDepthMask{}) // interface compliance check
func (ϟa *GlDepthMask) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_50_result := context              // Contextʳ
	ctx := GetContext_50_result                  // Contextʳ
	ctx.Rasterizing.DepthMask = ϟa.Enabled
	ϟb.Push(value.Bool(ϟa.Enabled))
	ϟb.Call(funcInfoGlDepthMask)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_50_result, ctx
	return nil
}

var _ = replay.Replayer(&GlDepthFunc{}) // interface compliance check
func (ϟa *GlDepthFunc) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_51_result := context              // Contextʳ
	ctx := GetContext_51_result                  // Contextʳ
	ctx.Rasterizing.DepthTestFunction = ϟa.Function
	ϟb.Push(value.U32(ϟa.Function))
	ϟb.Call(funcInfoGlDepthFunc)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_51_result, ctx
	return nil
}

var _ = replay.Replayer(&GlDepthRangef{}) // interface compliance check
func (ϟa *GlDepthRangef) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_52_result := context              // Contextʳ
	ctx := GetContext_52_result                  // Contextʳ
	ctx.Rasterizing.DepthNear = ϟa.Near
	ctx.Rasterizing.DepthFar = ϟa.Far
	ϟb.Push(value.F32(ϟa.Near))
	ϟb.Push(value.F32(ϟa.Far))
	ϟb.Call(funcInfoGlDepthRangef)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_52_result, ctx
	return nil
}

var _ = replay.Replayer(&GlColorMask{}) // interface compliance check
func (ϟa *GlColorMask) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	ϟb.Push(value.Bool(ϟa.Red))
	ϟb.Push(value.Bool(ϟa.Green))
	ϟb.Push(value.Bool(ϟa.Blue))
	ϟb.Push(value.Bool(ϟa.Alpha))
	ϟb.Call(funcInfoGlColorMask)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_53_result, ctx
	return nil
}

var _ = replay.Replayer(&GlStencilMask{}) // interface compliance check
func (ϟa *GlStencilMask) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_54_result := context              // Contextʳ
	ctx := GetContext_54_result                  // Contextʳ
	ctx.Rasterizing.StencilMask[FaceMode_GL_FRONT] = ϟa.Mask
	ctx.Rasterizing.StencilMask[FaceMode_GL_BACK] = ϟa.Mask
	ϟb.Push(value.U32(ϟa.Mask))
	ϟb.Call(funcInfoGlStencilMask)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_54_result, ctx
	return nil
}

var _ = replay.Replayer(&GlStencilMaskSeparate{}) // interface compliance check
func (ϟa *GlStencilMaskSeparate) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	ϟb.Push(value.U32(ϟa.Face))
	ϟb.Push(value.U32(ϟa.Mask))
	ϟb.Call(funcInfoGlStencilMaskSeparate)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_55_result, ctx
	return nil
}

var _ = replay.Replayer(&GlStencilFuncSeparate{}) // interface compliance check
func (ϟa *GlStencilFuncSeparate) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U32(ϟa.Face))
	ϟb.Push(value.U32(ϟa.Function))
	ϟb.Push(value.S32(ϟa.ReferenceValue))
	ϟb.Push(value.S32(ϟa.Mask))
	ϟb.Call(funcInfoGlStencilFuncSeparate)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlStencilOpSeparate{}) // interface compliance check
func (ϟa *GlStencilOpSeparate) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U32(ϟa.Face))
	ϟb.Push(value.U32(ϟa.StencilFail))
	ϟb.Push(value.U32(ϟa.StencilPassDepthFail))
	ϟb.Push(value.U32(ϟa.StencilPassDepthPass))
	ϟb.Call(funcInfoGlStencilOpSeparate)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlFrontFace{}) // interface compliance check
func (ϟa *GlFrontFace) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_56_result := context              // Contextʳ
	ctx := GetContext_56_result                  // Contextʳ
	ctx.Rasterizing.FrontFace = ϟa.Orientation
	ϟb.Push(value.U32(ϟa.Orientation))
	ϟb.Call(funcInfoGlFrontFace)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_56_result, ctx
	return nil
}

var _ = replay.Replayer(&GlViewport{}) // interface compliance check
func (ϟa *GlViewport) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	ϟb.Push(value.S32(ϟa.X))
	ϟb.Push(value.S32(ϟa.Y))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.Call(funcInfoGlViewport)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_57_result, ctx
	return nil
}

var _ = replay.Replayer(&GlScissor{}) // interface compliance check
func (ϟa *GlScissor) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	ϟb.Push(value.S32(ϟa.X))
	ϟb.Push(value.S32(ϟa.Y))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.Call(funcInfoGlScissor)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_58_result, ctx
	return nil
}

var _ = replay.Replayer(&GlActiveTexture{}) // interface compliance check
func (ϟa *GlActiveTexture) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	ϟb.Push(value.U32(ϟa.Unit))
	ϟb.Call(funcInfoGlActiveTexture)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_59_result, ctx
	return nil
}

var _ = replay.Replayer(&GlGenTextures{}) // interface compliance check
func (ϟa *GlGenTextures) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	t := ϟa.Textures.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // TextureIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                   // Contextʳ
	GetContext_60_result := context                                // Contextʳ
	ctx := GetContext_60_result                                    // Contextʳ
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Textures.value())
	ϟb.Call(funcInfoGlGenTextures)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		id := TextureId(ϟa.Textures.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)) // TextureId
		ctx.Instances.Textures[id] = func() *Texture {
			s := &Texture{}
			s.Init()
			return s
		}()
		t.Index(uint64(i), ϟs).replayWrite(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _ = t, context, GetContext_60_result, ctx
	return nil
}

var _ = replay.Replayer(&GlDeleteTextures{}) // interface compliance check
func (ϟa *GlDeleteTextures) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	t := ϟa.Textures.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // TextureIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                   // Contextʳ
	GetContext_61_result := context                                // Contextʳ
	ctx := GetContext_61_result                                    // Contextʳ
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		ctx.Instances.Textures[t.Index(uint64(i), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)] = (*Texture)(nil)
	}
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Textures.value())
	ϟb.Call(funcInfoGlDeleteTextures)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = t, context, GetContext_61_result, ctx
	return nil
}

var _ = replay.Replayer(&GlIsTexture{}) // interface compliance check
func (ϟa *GlIsTexture) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_62_result := context              // Contextʳ
	ctx := GetContext_62_result                  // Contextʳ
	if key, remap := ϟa.Texture.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Texture.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Texture.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlIsTexture)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_62_result, ctx
	return nil
}

var _ = replay.Replayer(&GlBindTexture{}) // interface compliance check
func (ϟa *GlBindTexture) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	ϟb.Push(value.U32(ϟa.Target))
	if key, remap := ϟa.Texture.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Texture.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Texture.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlBindTexture)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_63_result, ctx
	return nil
}

var _ = replay.Replayer(&GlTexImage2D{}) // interface compliance check
func (ϟa *GlTexImage2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
				l.Data = U8ᵖ(ϟa.Data).Slice(uint64(uint32(0)), uint64(l.Size), ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb).Clone(ϟs)
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
				l.Data = U8ᵖ(ϟa.Data).Slice(uint64(uint32(0)), uint64(l.Size), ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb).Clone(ϟs)
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
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Level))
	ϟb.Push(value.U32(ϟa.InternalFormat))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.Push(value.S32(ϟa.Border))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Data.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlTexImage2D)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_64_result, ctx
	return nil
}

var _ = replay.Replayer(&GlTexSubImage2D{}) // interface compliance check
func (ϟa *GlTexSubImage2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
			l.Data = U8ᵖ(ϟa.Data).Slice(uint64(uint32(0)), uint64(l.Size), ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb).Clone(ϟs)
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
			l.Data = U8ᵖ(ϟa.Data).Slice(uint64(uint32(0)), uint64(l.Size), ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb).Clone(ϟs)
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
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Level))
	ϟb.Push(value.S32(ϟa.Xoffset))
	ϟb.Push(value.S32(ϟa.Yoffset))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Data.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlTexSubImage2D)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_65_result, ctx
	return nil
}

var _ = replay.Replayer(&GlCopyTexImage2D{}) // interface compliance check
func (ϟa *GlCopyTexImage2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Level))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.S32(ϟa.X))
	ϟb.Push(value.S32(ϟa.Y))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.Push(value.S32(ϟa.Border))
	ϟb.Call(funcInfoGlCopyTexImage2D)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlCopyTexSubImage2D{}) // interface compliance check
func (ϟa *GlCopyTexSubImage2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Level))
	ϟb.Push(value.S32(ϟa.Xoffset))
	ϟb.Push(value.S32(ϟa.Yoffset))
	ϟb.Push(value.S32(ϟa.X))
	ϟb.Push(value.S32(ϟa.Y))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.Call(funcInfoGlCopyTexSubImage2D)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlCompressedTexImage2D{}) // interface compliance check
func (ϟa *GlCompressedTexImage2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
			l.Data = U8ᵖ(ϟa.Data).Slice(uint64(uint32(0)), uint64(l.Size), ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb).Clone(ϟs)
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
			l.Data = U8ᵖ(ϟa.Data).Slice(uint64(uint32(0)), uint64(l.Size), ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb).Clone(ϟs)
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
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Level))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.Push(value.S32(ϟa.Border))
	ϟb.Push(value.S32(ϟa.ImageSize))
	ϟb.Push(ϟa.Data.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlCompressedTexImage2D)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_66_result, ctx
	return nil
}

var _ = replay.Replayer(&GlCompressedTexSubImage2D{}) // interface compliance check
func (ϟa *GlCompressedTexSubImage2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Level))
	ϟb.Push(value.S32(ϟa.Xoffset))
	ϟb.Push(value.S32(ϟa.Yoffset))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.S32(ϟa.ImageSize))
	ϟb.Push(ϟa.Data.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlCompressedTexSubImage2D)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlGenerateMipmap{}) // interface compliance check
func (ϟa *GlGenerateMipmap) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Call(funcInfoGlGenerateMipmap)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlReadPixels{}) // interface compliance check
func (ϟa *GlReadPixels) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.S32(ϟa.X))
	ϟb.Push(value.S32(ϟa.Y))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Data.value())
	ϟb.Call(funcInfoGlReadPixels)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Data.Slice(uint64(uint32(0)), uint64(externs{ϟs, ϟd, ϟl}.imageSize(uint32(ϟa.Width), uint32(ϟa.Height), TexelFormat(ϟa.Format), ϟa.Type)), ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

var _ = replay.Replayer(&GlGenFramebuffers{}) // interface compliance check
func (ϟa *GlGenFramebuffers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	f := ϟa.Framebuffers.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // FramebufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                       // Contextʳ
	GetContext_67_result := context                                    // Contextʳ
	ctx := GetContext_67_result                                        // Contextʳ
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Framebuffers.value())
	ϟb.Call(funcInfoGlGenFramebuffers)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		id := FramebufferId(ϟa.Framebuffers.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)) // FramebufferId
		ctx.Instances.Framebuffers[id] = func() *Framebuffer {
			s := &Framebuffer{}
			s.Init()
			return s
		}()
		f.Index(uint64(i), ϟs).replayWrite(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _ = f, context, GetContext_67_result, ctx
	return nil
}

var _ = replay.Replayer(&GlBindFramebuffer{}) // interface compliance check
func (ϟa *GlBindFramebuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	ϟb.Push(value.U32(ϟa.Target))
	if key, remap := ϟa.Framebuffer.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Framebuffer.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Framebuffer.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlBindFramebuffer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_68_result, ctx
	return nil
}

var _ = replay.Replayer(&GlCheckFramebufferStatus{}) // interface compliance check
func (ϟa *GlCheckFramebufferStatus) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Call(funcInfoGlCheckFramebufferStatus)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlDeleteFramebuffers{}) // interface compliance check
func (ϟa *GlDeleteFramebuffers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	f := ϟa.Framebuffers.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // FramebufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                       // Contextʳ
	GetContext_69_result := context                                    // Contextʳ
	ctx := GetContext_69_result                                        // Contextʳ
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		ctx.Instances.Framebuffers[f.Index(uint64(i), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)] = (*Framebuffer)(nil)
	}
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Framebuffers.value())
	ϟb.Call(funcInfoGlDeleteFramebuffers)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = f, context, GetContext_69_result, ctx
	return nil
}

var _ = replay.Replayer(&GlIsFramebuffer{}) // interface compliance check
func (ϟa *GlIsFramebuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_70_result := context              // Contextʳ
	ctx := GetContext_70_result                  // Contextʳ
	if key, remap := ϟa.Framebuffer.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Framebuffer.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Framebuffer.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlIsFramebuffer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_70_result, ctx
	return nil
}

var _ = replay.Replayer(&GlGenRenderbuffers{}) // interface compliance check
func (ϟa *GlGenRenderbuffers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	r := ϟa.Renderbuffers.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // RenderbufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                        // Contextʳ
	GetContext_71_result := context                                     // Contextʳ
	ctx := GetContext_71_result                                         // Contextʳ
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Renderbuffers.value())
	ϟb.Call(funcInfoGlGenRenderbuffers)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		id := RenderbufferId(ϟa.Renderbuffers.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)) // RenderbufferId
		ctx.Instances.Renderbuffers[id] = func() *Renderbuffer {
			s := &Renderbuffer{}
			s.Init()
			return s
		}()
		r.Index(uint64(i), ϟs).replayWrite(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _ = r, context, GetContext_71_result, ctx
	return nil
}

var _ = replay.Replayer(&GlBindRenderbuffer{}) // interface compliance check
func (ϟa *GlBindRenderbuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	ϟb.Push(value.U32(ϟa.Target))
	if key, remap := ϟa.Renderbuffer.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Renderbuffer.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Renderbuffer.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlBindRenderbuffer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_72_result, ctx
	return nil
}

var _ = replay.Replayer(&GlRenderbufferStorage{}) // interface compliance check
func (ϟa *GlRenderbufferStorage) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.Call(funcInfoGlRenderbufferStorage)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = context, GetContext_73_result, ctx, id, rb
	return nil
}

var _ = replay.Replayer(&GlDeleteRenderbuffers{}) // interface compliance check
func (ϟa *GlDeleteRenderbuffers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	r := ϟa.Renderbuffers.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // RenderbufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                        // Contextʳ
	GetContext_74_result := context                                     // Contextʳ
	ctx := GetContext_74_result                                         // Contextʳ
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		ctx.Instances.Renderbuffers[r.Index(uint64(i), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)] = (*Renderbuffer)(nil)
	}
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Renderbuffers.value())
	ϟb.Call(funcInfoGlDeleteRenderbuffers)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = r, context, GetContext_74_result, ctx
	return nil
}

var _ = replay.Replayer(&GlIsRenderbuffer{}) // interface compliance check
func (ϟa *GlIsRenderbuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_75_result := context              // Contextʳ
	ctx := GetContext_75_result                  // Contextʳ
	if key, remap := ϟa.Renderbuffer.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Renderbuffer.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Renderbuffer.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlIsRenderbuffer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_75_result, ctx
	return nil
}

var _ = replay.Replayer(&GlGetRenderbufferParameteriv{}) // interface compliance check
func (ϟa *GlGetRenderbufferParameteriv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_76_result := context              // Contextʳ
	ctx := GetContext_76_result                  // Contextʳ
	id := ctx.BoundRenderbuffers.Get(ϟa.Target)  // RenderbufferId
	rb := ctx.Instances.Renderbuffers.Get(id)    // Renderbufferʳ
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(ϟa.Values.value())
	ϟb.Call(funcInfoGlGetRenderbufferParameteriv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Values.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(func() (result int32) {
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
	}(), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _ = context, GetContext_76_result, ctx, id, rb
	return nil
}

var _ = replay.Replayer(&GlGenBuffers{}) // interface compliance check
func (ϟa *GlGenBuffers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	b := ϟa.Buffers.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // BufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                  // Contextʳ
	GetContext_77_result := context                               // Contextʳ
	ctx := GetContext_77_result                                   // Contextʳ
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Buffers.value())
	ϟb.Call(funcInfoGlGenBuffers)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		id := BufferId(ϟa.Buffers.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)) // BufferId
		ctx.Instances.Buffers[id] = func() *Buffer {
			s := &Buffer{}
			s.Init()
			return s
		}()
		b.Index(uint64(i), ϟs).replayWrite(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _ = b, context, GetContext_77_result, ctx
	return nil
}

var _ = replay.Replayer(&GlBindBuffer{}) // interface compliance check
func (ϟa *GlBindBuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	ϟb.Push(value.U32(ϟa.Target))
	if key, remap := ϟa.Buffer.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Buffer.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Buffer.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlBindBuffer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_78_result, ctx
	return nil
}

var _ = replay.Replayer(&GlBufferData{}) // interface compliance check
func (ϟa *GlBufferData) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
			return U8ᵖ(ϟa.Data).Slice(uint64(int32(0)), uint64(ϟa.Size), ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb).Clone(ϟs)
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
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Size))
	ϟb.Push(ϟa.Data.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Usage))
	ϟb.Call(funcInfoGlBufferData)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = context, GetContext_79_result, ctx, id, b
	return nil
}

var _ = replay.Replayer(&GlBufferSubData{}) // interface compliance check
func (ϟa *GlBufferSubData) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Data.Slice(uint64(int32(0)), uint64(ϟa.Size), ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Offset))
	ϟb.Push(value.S32(ϟa.Size))
	ϟb.Push(ϟa.Data.value())
	ϟb.Call(funcInfoGlBufferSubData)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlDeleteBuffers{}) // interface compliance check
func (ϟa *GlDeleteBuffers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	b := ϟa.Buffers.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // BufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                  // Contextʳ
	GetContext_80_result := context                               // Contextʳ
	ctx := GetContext_80_result                                   // Contextʳ
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		ctx.Instances.Buffers[b.Index(uint64(i), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)] = (*Buffer)(nil)
	}
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Buffers.value())
	ϟb.Call(funcInfoGlDeleteBuffers)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = b, context, GetContext_80_result, ctx
	return nil
}

var _ = replay.Replayer(&GlIsBuffer{}) // interface compliance check
func (ϟa *GlIsBuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_81_result := context              // Contextʳ
	ctx := GetContext_81_result                  // Contextʳ
	if key, remap := ϟa.Buffer.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Buffer.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Buffer.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlIsBuffer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_81_result, ctx
	return nil
}

var _ = replay.Replayer(&GlGetBufferParameteriv{}) // interface compliance check
func (ϟa *GlGetBufferParameteriv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_82_result := context              // Contextʳ
	ctx := GetContext_82_result                  // Contextʳ
	id := ctx.BoundBuffers.Get(ϟa.Target)        // BufferId
	b := ctx.Instances.Buffers.Get(id)           // Bufferʳ
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlGetBufferParameteriv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(func() (result int32) {
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
	}(), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _ = context, GetContext_82_result, ctx, id, b
	return nil
}

var _ = replay.Replayer(&GlCreateShader{}) // interface compliance check
func (ϟa *GlCreateShader) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_83_result := context              // Contextʳ
	ctx := GetContext_83_result                  // Contextʳ
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Call(funcInfoGlCreateShader)
	if key, remap := ϟa.Result.remap(ϟa, ϟs); remap {
		if ptr, found := ϟb.Remappings[key]; !found {
			ptr = ϟb.AllocateMemory(uint64(4))
			ϟb.Clone(0)
			ϟb.Store(ptr)
			ϟb.Remappings[key] = ptr
		}
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	id := ShaderId(ϟa.Result) // ShaderId
	ctx.Instances.Shaders[id] = func() *Shader {
		s := &Shader{}
		s.Init()
		return s
	}()
	s := ctx.Instances.Shaders.Get(id) // Shaderʳ
	s.Type = ϟa.Type
	_, _, _, _, _ = context, GetContext_83_result, ctx, id, s
	return nil
}

var _ = replay.Replayer(&GlDeleteShader{}) // interface compliance check
func (ϟa *GlDeleteShader) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_84_result := context              // Contextʳ
	ctx := GetContext_84_result                  // Contextʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // Shaderʳ
	s.Deletable = true
	ctx.Instances.Shaders[ϟa.Shader] = (*Shader)(nil)
	if key, remap := ϟa.Shader.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Shader.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Shader.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlDeleteShader)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = context, GetContext_84_result, ctx, s
	return nil
}

var _ = replay.Replayer(&GlShaderSource{}) // interface compliance check
func (ϟa *GlShaderSource) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
			switch ((ϟa.Length) == (S32ᵖ{})) || ((lengths.Index(uint64(i), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)) < (int32(0))) {
			case true:
				return strings.TrimRight(string(sources.Index(uint64(i), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb).StringSlice(ϟs, ϟd, ϟl, true).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)), "\x00")
			case false:
				return string(sources.Index(uint64(i), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb).Slice(uint64(int32(0)), uint64(lengths.Index(uint64(i), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb))
			default:
				// TODO: better unmatched handling
				panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ((ϟa.Length) == (S32ᵖ{})) || ((lengths.Index(uint64(i), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)) < (int32(0))), ϟa))
				return result
			}
		}() // string
		s.Source += str
		_ = str
	}
	if key, remap := ϟa.Shader.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Shader.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Shader.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Source.value())
	ϟb.Push(ϟa.Length.value())
	ϟb.Call(funcInfoGlShaderSource)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = sources, lengths, context, GetContext_85_result, ctx, s
	return nil
}

var _ = replay.Replayer(&GlShaderBinary{}) // interface compliance check
func (ϟa *GlShaderBinary) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	for i := int32(int32(0)); i < ϟa.Count; i++ {
	}
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Shaders.value())
	ϟb.Push(value.U32(ϟa.BinaryFormat))
	ϟb.Push(ϟa.Binary.value())
	ϟb.Push(value.S32(ϟa.BinarySize))
	ϟb.Call(funcInfoGlShaderBinary)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlGetShaderInfoLog{}) // interface compliance check
func (ϟa *GlGetShaderInfoLog) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	ϟdst, ϟsrc := ϟa.Info.Slice(uint64(int32(0)), uint64(l), ϟs).Copy(s.InfoLog.Slice(uint64(int32(0)), uint64(l), ϟs), ϟs, ϟd, ϟl)
	ϟsrc.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	if key, remap := ϟa.Shader.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Shader.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Shader.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.BufferLength))
	ϟb.Push(ϟa.StringLengthWritten.value())
	ϟb.Push(ϟa.Info.value())
	ϟb.Call(funcInfoGlGetShaderInfoLog)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟdst.onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.StringLengthWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(l, ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _, _, _ = context, GetContext_86_result, ctx, s, min_87_a, min_87_b, min_87_result, l
	return nil
}

var _ = replay.Replayer(&GlGetShaderSource{}) // interface compliance check
func (ϟa *GlGetShaderSource) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	ϟdst, ϟsrc := ϟa.Source.Slice(uint64(int32(0)), uint64(l), ϟs).Copy(MakeCharˢFromString(s.Source, ϟs).Slice(uint64(int32(0)), uint64(l), ϟs), ϟs, ϟd, ϟl)
	ϟsrc.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	if key, remap := ϟa.Shader.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Shader.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Shader.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.BufferLength))
	ϟb.Push(ϟa.StringLengthWritten.value())
	ϟb.Push(ϟa.Source.value())
	ϟb.Call(funcInfoGlGetShaderSource)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟdst.onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.StringLengthWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(l, ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _, _, _ = context, GetContext_88_result, ctx, s, min_89_a, min_89_b, min_89_result, l
	return nil
}

var _ = replay.Replayer(&GlReleaseShaderCompiler{}) // interface compliance check
func (ϟa *GlReleaseShaderCompiler) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoGlReleaseShaderCompiler)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlCompileShader{}) // interface compliance check
func (ϟa *GlCompileShader) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	if key, remap := ϟa.Shader.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Shader.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Shader.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlCompileShader)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlIsShader{}) // interface compliance check
func (ϟa *GlIsShader) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_90_result := context              // Contextʳ
	ctx := GetContext_90_result                  // Contextʳ
	if key, remap := ϟa.Shader.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Shader.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Shader.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlIsShader)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_90_result, ctx
	return nil
}

var _ = replay.Replayer(&GlCreateProgram{}) // interface compliance check
func (ϟa *GlCreateProgram) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_91_result := context              // Contextʳ
	ctx := GetContext_91_result                  // Contextʳ
	ϟb.Call(funcInfoGlCreateProgram)
	if key, remap := ϟa.Result.remap(ϟa, ϟs); remap {
		if ptr, found := ϟb.Remappings[key]; !found {
			ptr = ϟb.AllocateMemory(uint64(4))
			ϟb.Clone(0)
			ϟb.Store(ptr)
			ϟb.Remappings[key] = ptr
		}
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	id := ProgramId(ϟa.Result) // ProgramId
	ctx.Instances.Programs[id] = func() *Program {
		s := &Program{}
		s.Init()
		return s
	}()
	_, _, _, _ = context, GetContext_91_result, ctx, id
	return nil
}

var _ = replay.Replayer(&GlDeleteProgram{}) // interface compliance check
func (ϟa *GlDeleteProgram) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_92_result := context              // Contextʳ
	ctx := GetContext_92_result                  // Contextʳ
	ctx.Instances.Programs[ϟa.Program] = (*Program)(nil)
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlDeleteProgram)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_92_result, ctx
	return nil
}

var _ = replay.Replayer(&GlAttachShader{}) // interface compliance check
func (ϟa *GlAttachShader) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_93_result := context              // Contextʳ
	ctx := GetContext_93_result                  // Contextʳ
	p := ctx.Instances.Programs.Get(ϟa.Program)  // Programʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // Shaderʳ
	p.Shaders[s.Type] = ϟa.Shader
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Shader.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Shader.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Shader.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlAttachShader)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = context, GetContext_93_result, ctx, p, s
	return nil
}

var _ = replay.Replayer(&GlDetachShader{}) // interface compliance check
func (ϟa *GlDetachShader) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_94_result := context              // Contextʳ
	ctx := GetContext_94_result                  // Contextʳ
	p := ctx.Instances.Programs.Get(ϟa.Program)  // Programʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // Shaderʳ
	p.Shaders[s.Type] = ShaderId(uint32(0))
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Shader.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Shader.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Shader.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlDetachShader)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = context, GetContext_94_result, ctx, p, s
	return nil
}

var _ = replay.Replayer(&GlGetAttachedShaders{}) // interface compliance check
func (ϟa *GlGetAttachedShaders) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.BufferLength))
	ϟb.Push(ϟa.ShadersLengthWritten.value())
	ϟb.Push(ϟa.Shaders.value())
	ϟb.Call(funcInfoGlGetAttachedShaders)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.ShadersLengthWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(l, ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _, _, _ = context, GetContext_95_result, ctx, p, min_96_a, min_96_b, min_96_result, l
	return nil
}

var _ = replay.Replayer(&GlLinkProgram{}) // interface compliance check
func (ϟa *GlLinkProgram) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlLinkProgram)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlGetProgramInfoLog{}) // interface compliance check
func (ϟa *GlGetProgramInfoLog) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	ϟdst, ϟsrc := ϟa.Info.Slice(uint64(int32(0)), uint64(l), ϟs).Copy(p.InfoLog.Slice(uint64(int32(0)), uint64(l), ϟs), ϟs, ϟd, ϟl)
	ϟsrc.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.BufferLength))
	ϟb.Push(ϟa.StringLengthWritten.value())
	ϟb.Push(ϟa.Info.value())
	ϟb.Call(funcInfoGlGetProgramInfoLog)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟdst.onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.StringLengthWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(l, ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _, _, _ = context, GetContext_97_result, ctx, p, min_98_a, min_98_b, min_98_result, l
	return nil
}

var _ = replay.Replayer(&GlUseProgram{}) // interface compliance check
func (ϟa *GlUseProgram) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_99_result := context              // Contextʳ
	ctx := GetContext_99_result                  // Contextʳ
	ctx.BoundProgram = ϟa.Program
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlUseProgram)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_99_result, ctx
	return nil
}

var _ = replay.Replayer(&GlIsProgram{}) // interface compliance check
func (ϟa *GlIsProgram) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_100_result := context             // Contextʳ
	ctx := GetContext_100_result                 // Contextʳ
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlIsProgram)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_100_result, ctx
	return nil
}

var _ = replay.Replayer(&GlValidateProgram{}) // interface compliance check
func (ϟa *GlValidateProgram) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlValidateProgram)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlClearColor{}) // interface compliance check
func (ϟa *GlClearColor) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	ϟb.Push(value.F32(ϟa.R))
	ϟb.Push(value.F32(ϟa.G))
	ϟb.Push(value.F32(ϟa.B))
	ϟb.Push(value.F32(ϟa.A))
	ϟb.Call(funcInfoGlClearColor)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_101_result, ctx
	return nil
}

var _ = replay.Replayer(&GlClearDepthf{}) // interface compliance check
func (ϟa *GlClearDepthf) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_102_result := context             // Contextʳ
	ctx := GetContext_102_result                 // Contextʳ
	ctx.Clearing.ClearDepth = ϟa.Depth
	ϟb.Push(value.F32(ϟa.Depth))
	ϟb.Call(funcInfoGlClearDepthf)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_102_result, ctx
	return nil
}

var _ = replay.Replayer(&GlClearStencil{}) // interface compliance check
func (ϟa *GlClearStencil) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_103_result := context             // Contextʳ
	ctx := GetContext_103_result                 // Contextʳ
	ctx.Clearing.ClearStencil = ϟa.Stencil
	ϟb.Push(value.S32(ϟa.Stencil))
	ϟb.Call(funcInfoGlClearStencil)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_103_result, ctx
	return nil
}

var _ = replay.Replayer(&GlClear{}) // interface compliance check
func (ϟa *GlClear) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	if (ClearMask_GL_COLOR_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	ϟb.Push(value.U32(ϟa.Mask))
	ϟb.Call(funcInfoGlClear)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlCullFace{}) // interface compliance check
func (ϟa *GlCullFace) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_104_result := context             // Contextʳ
	ctx := GetContext_104_result                 // Contextʳ
	ctx.Rasterizing.CullFace = ϟa.Mode
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Call(funcInfoGlCullFace)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_104_result, ctx
	return nil
}

var _ = replay.Replayer(&GlPolygonOffset{}) // interface compliance check
func (ϟa *GlPolygonOffset) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_105_result := context             // Contextʳ
	ctx := GetContext_105_result                 // Contextʳ
	ctx.Rasterizing.PolygonOffsetUnits = ϟa.Units
	ctx.Rasterizing.PolygonOffsetFactor = ϟa.ScaleFactor
	ϟb.Push(value.F32(ϟa.ScaleFactor))
	ϟb.Push(value.F32(ϟa.Units))
	ϟb.Call(funcInfoGlPolygonOffset)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_105_result, ctx
	return nil
}

var _ = replay.Replayer(&GlLineWidth{}) // interface compliance check
func (ϟa *GlLineWidth) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_106_result := context             // Contextʳ
	ctx := GetContext_106_result                 // Contextʳ
	ctx.Rasterizing.LineWidth = ϟa.Width
	ϟb.Push(value.F32(ϟa.Width))
	ϟb.Call(funcInfoGlLineWidth)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_106_result, ctx
	return nil
}

var _ = replay.Replayer(&GlSampleCoverage{}) // interface compliance check
func (ϟa *GlSampleCoverage) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_107_result := context             // Contextʳ
	ctx := GetContext_107_result                 // Contextʳ
	ctx.Rasterizing.SampleCoverageValue = ϟa.Value
	ctx.Rasterizing.SampleCoverageInvert = ϟa.Invert
	ϟb.Push(value.F32(ϟa.Value))
	ϟb.Push(value.Bool(ϟa.Invert))
	ϟb.Call(funcInfoGlSampleCoverage)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_107_result, ctx
	return nil
}

var _ = replay.Replayer(&GlHint{}) // interface compliance check
func (ϟa *GlHint) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_108_result := context             // Contextʳ
	ctx := GetContext_108_result                 // Contextʳ
	ctx.GenerateMipmapHint = ϟa.Mode
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Call(funcInfoGlHint)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_108_result, ctx
	return nil
}

var _ = replay.Replayer(&GlFramebufferRenderbuffer{}) // interface compliance check
func (ϟa *GlFramebufferRenderbuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	ϟb.Push(value.U32(ϟa.FramebufferTarget))
	ϟb.Push(value.U32(ϟa.FramebufferAttachment))
	ϟb.Push(value.U32(ϟa.RenderbufferTarget))
	if key, remap := ϟa.Renderbuffer.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Renderbuffer.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Renderbuffer.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlFramebufferRenderbuffer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = context, GetContext_109_result, ctx, target, framebufferId, framebuffer, attachment
	return nil
}

var _ = replay.Replayer(&GlFramebufferTexture2D{}) // interface compliance check
func (ϟa *GlFramebufferTexture2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	ϟb.Push(value.U32(ϟa.FramebufferTarget))
	ϟb.Push(value.U32(ϟa.FramebufferAttachment))
	ϟb.Push(value.U32(ϟa.TextureTarget))
	if key, remap := ϟa.Texture.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Texture.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Texture.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.S32(ϟa.Level))
	ϟb.Call(funcInfoGlFramebufferTexture2D)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = context, GetContext_110_result, ctx, target, framebufferId, framebuffer, attachment
	return nil
}

var _ = replay.Replayer(&GlGetFramebufferAttachmentParameteriv{}) // interface compliance check
func (ϟa *GlGetFramebufferAttachmentParameteriv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	ϟb.Push(value.U32(ϟa.FramebufferTarget))
	ϟb.Push(value.U32(ϟa.Attachment))
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlGetFramebufferAttachmentParameteriv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(func() (result int32) {
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
	}(), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _, _ = context, GetContext_111_result, ctx, target, framebufferId, framebuffer, a
	return nil
}

var _ = replay.Replayer(&GlDrawElements{}) // interface compliance check
func (ϟa *GlDrawElements) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
					arr.Pointer.Slice(uint64(offset), uint64((offset)+(elsize)), ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
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
					arr.Pointer.Slice(uint64(offset), uint64((offset)+(elsize)), ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
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
		index_data.Slice(uint64(uint32(0)), uint64((uint32(ϟa.ElementCount))*(IndexSize_117_result)), ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _, _, _, _, _, _, _ = index_data, first, last, ReadVertexArrays_115_ctx, ReadVertexArrays_115_first_index, ReadVertexArrays_115_last_index, IndexSize_117_indices_type, IndexSize_117_result
	}
	ϟb.Push(value.U32(ϟa.DrawMode))
	ϟb.Push(value.S32(ϟa.ElementCount))
	ϟb.Push(value.U32(ϟa.IndicesType))
	ϟb.Push(ϟa.Indices.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlDrawElements)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = context, GetContext_112_result, ctx, count, id
	return nil
}

var _ = replay.Replayer(&GlDrawArrays{}) // interface compliance check
func (ϟa *GlDrawArrays) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
				arr.Pointer.Slice(uint64(offset), uint64((offset)+(elsize)), ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
				_ = offset
			}
			_, _, _, _ = vertexAttribTypeSize_120_t, vertexAttribTypeSize_120_result, elsize, elstride
		}
		_ = arr
	}
	ϟb.Push(value.U32(ϟa.DrawMode))
	ϟb.Push(value.S32(ϟa.FirstIndex))
	ϟb.Push(value.S32(ϟa.IndexCount))
	ϟb.Call(funcInfoGlDrawArrays)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = context, GetContext_118_result, ctx, last_index, ReadVertexArrays_119_ctx, ReadVertexArrays_119_first_index, ReadVertexArrays_119_last_index
	return nil
}

var _ = replay.Replayer(&GlFlush{}) // interface compliance check
func (ϟa *GlFlush) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoGlFlush)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlFinish{}) // interface compliance check
func (ϟa *GlFinish) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoGlFinish)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlGetBooleanv{}) // interface compliance check
func (ϟa *GlGetBooleanv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	v := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs) // Boolˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                                        // Contextʳ
	GetContext_121_result := context                                                                    // Contextʳ
	ctx := GetContext_121_result                                                                        // Contextʳ
	ϟb.Push(value.U32(ϟa.Param))
	ϟb.Push(ϟa.Values.value())
	ϟb.Call(funcInfoGlGetBooleanv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	switch ϟa.Param {
	case StateVariable_GL_BLEND:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Capabilities.Get(Capability_GL_BLEND), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_CULL_FACE:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Capabilities.Get(Capability_GL_CULL_FACE), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_DEPTH_TEST:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Capabilities.Get(Capability_GL_DEPTH_TEST), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_DITHER:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Capabilities.Get(Capability_GL_DITHER), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_POLYGON_OFFSET_FILL:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Capabilities.Get(Capability_GL_POLYGON_OFFSET_FILL), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_SAMPLE_ALPHA_TO_COVERAGE:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Capabilities.Get(Capability_GL_SAMPLE_ALPHA_TO_COVERAGE), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_SAMPLE_COVERAGE:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Capabilities.Get(Capability_GL_SAMPLE_COVERAGE), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_SCISSOR_TEST:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Capabilities.Get(Capability_GL_SCISSOR_TEST), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_STENCIL_TEST:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Capabilities.Get(Capability_GL_STENCIL_TEST), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_DEPTH_WRITEMASK:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Rasterizing.DepthMask, ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_COLOR_WRITEMASK:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Rasterizing.ColorMaskRed, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(1), ϟs).replayWrite(ctx.Rasterizing.ColorMaskGreen, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(2), ϟs).replayWrite(ctx.Rasterizing.ColorMaskBlue, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(3), ϟs).replayWrite(ctx.Rasterizing.ColorMaskAlpha, ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_SAMPLE_COVERAGE_INVERT:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Rasterizing.SampleCoverageInvert, ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_SHADER_COMPILER:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	default:
		v := ϟa.Param
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	_, _, _, _ = v, context, GetContext_121_result, ctx
	return nil
}

var _ = replay.Replayer(&GlGetFloatv{}) // interface compliance check
func (ϟa *GlGetFloatv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	v := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs) // F32ˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                                        // Contextʳ
	GetContext_122_result := context                                                                    // Contextʳ
	ctx := GetContext_122_result                                                                        // Contextʳ
	ϟb.Push(value.U32(ϟa.Param))
	ϟb.Push(ϟa.Values.value())
	ϟb.Call(funcInfoGlGetFloatv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	switch ϟa.Param {
	case StateVariable_GL_DEPTH_RANGE:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Rasterizing.DepthNear, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(1), ϟs).replayWrite(ctx.Rasterizing.DepthFar, ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_LINE_WIDTH:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Rasterizing.LineWidth, ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_POLYGON_OFFSET_FACTOR:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Rasterizing.PolygonOffsetFactor, ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_POLYGON_OFFSET_UNITS:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Rasterizing.PolygonOffsetUnits, ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_SAMPLE_COVERAGE_VALUE:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Rasterizing.SampleCoverageValue, ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_COLOR_CLEAR_VALUE:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Clearing.ClearColor.Red, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(1), ϟs).replayWrite(ctx.Clearing.ClearColor.Green, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(2), ϟs).replayWrite(ctx.Clearing.ClearColor.Blue, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(3), ϟs).replayWrite(ctx.Clearing.ClearColor.Alpha, ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_DEPTH_CLEAR_VALUE:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Clearing.ClearDepth, ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_ALIASED_LINE_WIDTH_RANGE:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(1), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(1), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_ALIASED_POINT_SIZE_RANGE:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(1), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(1), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	default:
		v := ϟa.Param
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	_, _, _, _ = v, context, GetContext_122_result, ctx
	return nil
}

var _ = replay.Replayer(&GlGetIntegerv{}) // interface compliance check
func (ϟa *GlGetIntegerv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	v := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs) // S32ˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                                        // Contextʳ
	GetContext_123_result := context                                                                    // Contextʳ
	ctx := GetContext_123_result                                                                        // Contextʳ
	ϟb.Push(value.U32(ϟa.Param))
	ϟb.Push(ϟa.Values.value())
	ϟb.Call(funcInfoGlGetIntegerv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	switch ϟa.Param {
	case StateVariable_GL_ACTIVE_TEXTURE:
		v.Index(uint64(0), ϟs).replayWrite(int32(ctx.ActiveTextureUnit), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_ARRAY_BUFFER_BINDING:
		v.Index(uint64(0), ϟs).replayWrite(int32(ctx.BoundBuffers.Get(BufferTarget_GL_ARRAY_BUFFER)), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_ELEMENT_ARRAY_BUFFER_BINDING:
		v.Index(uint64(0), ϟs).replayWrite(int32(ctx.BoundBuffers.Get(BufferTarget_GL_ELEMENT_ARRAY_BUFFER)), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_BLEND_SRC_ALPHA:
		v.Index(uint64(0), ϟs).replayWrite(int32(ctx.Blending.SrcAlphaBlendFactor), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_BLEND_SRC_RGB:
		v.Index(uint64(0), ϟs).replayWrite(int32(ctx.Blending.SrcRgbBlendFactor), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_BLEND_DST_ALPHA:
		v.Index(uint64(0), ϟs).replayWrite(int32(ctx.Blending.DstAlphaBlendFactor), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_BLEND_DST_RGB:
		v.Index(uint64(0), ϟs).replayWrite(int32(ctx.Blending.DstRgbBlendFactor), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_BLEND_EQUATION_RGB:
		v.Index(uint64(0), ϟs).replayWrite(int32(ctx.Blending.BlendEquationRgb), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_BLEND_EQUATION_ALPHA:
		v.Index(uint64(0), ϟs).replayWrite(int32(ctx.Blending.BlendEquationAlpha), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_BLEND_COLOR:
		v.Index(uint64(0), ϟs).replayWrite(int32(ctx.Blending.BlendColor.Red), ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(1), ϟs).replayWrite(int32(ctx.Blending.BlendColor.Green), ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(2), ϟs).replayWrite(int32(ctx.Blending.BlendColor.Blue), ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(3), ϟs).replayWrite(int32(ctx.Blending.BlendColor.Alpha), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_DEPTH_FUNC:
		v.Index(uint64(0), ϟs).replayWrite(int32(ctx.Rasterizing.DepthTestFunction), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_DEPTH_CLEAR_VALUE:
		v.Index(uint64(0), ϟs).replayWrite(int32(ctx.Clearing.ClearDepth), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_STENCIL_WRITEMASK:
		v.Index(uint64(0), ϟs).replayWrite(int32(ctx.Rasterizing.StencilMask.Get(FaceMode_GL_FRONT)), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_STENCIL_BACK_WRITEMASK:
		v.Index(uint64(0), ϟs).replayWrite(int32(ctx.Rasterizing.StencilMask.Get(FaceMode_GL_BACK)), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_VIEWPORT:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Rasterizing.Viewport.X, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(1), ϟs).replayWrite(ctx.Rasterizing.Viewport.Y, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(2), ϟs).replayWrite(ctx.Rasterizing.Viewport.Width, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(3), ϟs).replayWrite(ctx.Rasterizing.Viewport.Height, ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_SCISSOR_BOX:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Rasterizing.Scissor.X, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(1), ϟs).replayWrite(ctx.Rasterizing.Scissor.Y, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(2), ϟs).replayWrite(ctx.Rasterizing.Scissor.Width, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(3), ϟs).replayWrite(ctx.Rasterizing.Scissor.Height, ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_FRONT_FACE:
		v.Index(uint64(0), ϟs).replayWrite(int32(ctx.Rasterizing.FrontFace), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_CULL_FACE_MODE:
		v.Index(uint64(0), ϟs).replayWrite(int32(ctx.Rasterizing.CullFace), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_STENCIL_CLEAR_VALUE:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Clearing.ClearStencil, ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_FRAMEBUFFER_BINDING:
		v.Index(uint64(0), ϟs).replayWrite(int32(ctx.BoundFramebuffers.Get(FramebufferTarget_GL_FRAMEBUFFER)), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_READ_FRAMEBUFFER_BINDING:
		v.Index(uint64(0), ϟs).replayWrite(int32(ctx.BoundFramebuffers.Get(FramebufferTarget_GL_READ_FRAMEBUFFER)), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_RENDERBUFFER_BINDING:
		v.Index(uint64(0), ϟs).replayWrite(int32(ctx.BoundRenderbuffers.Get(RenderbufferTarget_GL_RENDERBUFFER)), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_CURRENT_PROGRAM:
		v.Index(uint64(0), ϟs).replayWrite(int32(ctx.BoundProgram), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_TEXTURE_BINDING_2D:
		v.Index(uint64(0), ϟs).replayWrite(int32(ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_2D)), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_TEXTURE_BINDING_CUBE_MAP:
		v.Index(uint64(0), ϟs).replayWrite(int32(ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(TextureTarget_GL_TEXTURE_CUBE_MAP)), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_GENERATE_MIPMAP_HINT:
		v.Index(uint64(0), ϟs).replayWrite(int32(ctx.GenerateMipmapHint), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb) // any
		v.Index(uint64(0), ϟs).replayWrite(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case StateVariable_GL_MAX_CUBE_MAP_TEXTURE_SIZE:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb) // any
		v.Index(uint64(0), ϟs).replayWrite(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case StateVariable_GL_MAX_FRAGMENT_UNIFORM_VECTORS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb) // any
		v.Index(uint64(0), ϟs).replayWrite(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case StateVariable_GL_MAX_RENDERBUFFER_SIZE:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb) // any
		v.Index(uint64(0), ϟs).replayWrite(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case StateVariable_GL_MAX_TEXTURE_IMAGE_UNITS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb) // any
		v.Index(uint64(0), ϟs).replayWrite(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case StateVariable_GL_MAX_TEXTURE_SIZE:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb) // any
		v.Index(uint64(0), ϟs).replayWrite(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case StateVariable_GL_MAX_VARYING_VECTORS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb) // any
		v.Index(uint64(0), ϟs).replayWrite(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case StateVariable_GL_MAX_VERTEX_ATTRIBS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb) // any
		v.Index(uint64(0), ϟs).replayWrite(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case StateVariable_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_MAX_VERTEX_UNIFORM_VECTORS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb) // any
		v.Index(uint64(0), ϟs).replayWrite(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case StateVariable_GL_MAX_VIEWPORT_DIMS:
		max_width := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)  // any
		max_height := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(1), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb) // any
		v.Index(uint64(0), ϟs).replayWrite(max_width, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(1), ϟs).replayWrite(max_height, ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = max_width, max_height
	case StateVariable_GL_NUM_COMPRESSED_TEXTURE_FORMATS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb) // any
		v.Index(uint64(0), ϟs).replayWrite(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case StateVariable_GL_NUM_SHADER_BINARY_FORMATS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb) // any
		v.Index(uint64(0), ϟs).replayWrite(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case StateVariable_GL_PACK_ALIGNMENT:
		v.Index(uint64(0), ϟs).replayWrite(ctx.PixelStorage.Get(PixelStoreParameter_GL_PACK_ALIGNMENT), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_UNPACK_ALIGNMENT:
		v.Index(uint64(0), ϟs).replayWrite(ctx.PixelStorage.Get(PixelStoreParameter_GL_UNPACK_ALIGNMENT), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_ALPHA_BITS:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_BLUE_BITS:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_GREEN_BITS:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_RED_BITS:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_DEPTH_BITS:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_SAMPLE_BUFFERS:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_SAMPLES:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_STENCIL_BITS:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_SUBPIXEL_BITS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb) // any
		v.Index(uint64(0), ϟs).replayWrite(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case StateVariable_GL_IMPLEMENTATION_COLOR_READ_FORMAT:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_IMPLEMENTATION_COLOR_READ_TYPE:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	case StateVariable_GL_GPU_DISJOINT_EXT:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	default:
		v := ϟa.Param
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	_, _, _, _ = v, context, GetContext_123_result, ctx
	return nil
}

var _ = replay.Replayer(&GlGetString{}) // interface compliance check
func (ϟa *GlGetString) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U32(ϟa.Param))
	ϟb.Call(funcInfoGlGetString)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlEnable{}) // interface compliance check
func (ϟa *GlEnable) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_124_result := context             // Contextʳ
	ctx := GetContext_124_result                 // Contextʳ
	ctx.Capabilities[ϟa.Capability] = true
	ϟb.Push(value.U32(ϟa.Capability))
	ϟb.Call(funcInfoGlEnable)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_124_result, ctx
	return nil
}

var _ = replay.Replayer(&GlDisable{}) // interface compliance check
func (ϟa *GlDisable) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_125_result := context             // Contextʳ
	ctx := GetContext_125_result                 // Contextʳ
	ctx.Capabilities[ϟa.Capability] = false
	ϟb.Push(value.U32(ϟa.Capability))
	ϟb.Call(funcInfoGlDisable)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_125_result, ctx
	return nil
}

var _ = replay.Replayer(&GlIsEnabled{}) // interface compliance check
func (ϟa *GlIsEnabled) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_126_result := context             // Contextʳ
	ctx := GetContext_126_result                 // Contextʳ
	ϟb.Push(value.U32(ϟa.Capability))
	ϟb.Call(funcInfoGlIsEnabled)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_126_result, ctx
	return nil
}

var _ = replay.Replayer(&GlMapBufferRange{}) // interface compliance check
func (ϟa *GlMapBufferRange) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Offset))
	ϟb.Push(value.S32(ϟa.Length))
	ϟb.Push(value.U32(ϟa.Access))
	ϟb.Call(funcInfoGlMapBufferRange)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlUnmapBuffer{}) // interface compliance check
func (ϟa *GlUnmapBuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Call(funcInfoGlUnmapBuffer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlInvalidateFramebuffer{}) // interface compliance check
func (ϟa *GlInvalidateFramebuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Attachments.value())
	ϟb.Call(funcInfoGlInvalidateFramebuffer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlRenderbufferStorageMultisample{}) // interface compliance check
func (ϟa *GlRenderbufferStorageMultisample) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.S32(ϟa.Samples))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.Call(funcInfoGlRenderbufferStorageMultisample)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlBlitFramebuffer{}) // interface compliance check
func (ϟa *GlBlitFramebuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
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
	ϟb.Call(funcInfoGlBlitFramebuffer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlGenQueries{}) // interface compliance check
func (ϟa *GlGenQueries) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	q := ϟa.Queries.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // QueryIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                  // Contextʳ
	GetContext_127_result := context                              // Contextʳ
	ctx := GetContext_127_result                                  // Contextʳ
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Queries.value())
	ϟb.Call(funcInfoGlGenQueries)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		id := QueryId(ϟa.Queries.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)) // QueryId
		ctx.Instances.Queries[id] = func() *Query {
			s := &Query{}
			s.Init()
			return s
		}()
		q.Index(uint64(i), ϟs).replayWrite(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _ = q, context, GetContext_127_result, ctx
	return nil
}

var _ = replay.Replayer(&GlBeginQuery{}) // interface compliance check
func (ϟa *GlBeginQuery) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U32(ϟa.Target))
	if key, remap := ϟa.Query.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Query.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Query.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlBeginQuery)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlEndQuery{}) // interface compliance check
func (ϟa *GlEndQuery) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Call(funcInfoGlEndQuery)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlDeleteQueries{}) // interface compliance check
func (ϟa *GlDeleteQueries) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	q := ϟa.Queries.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // QueryIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                  // Contextʳ
	GetContext_128_result := context                              // Contextʳ
	ctx := GetContext_128_result                                  // Contextʳ
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		ctx.Instances.Queries[q.Index(uint64(i), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)] = (*Query)(nil)
	}
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Queries.value())
	ϟb.Call(funcInfoGlDeleteQueries)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = q, context, GetContext_128_result, ctx
	return nil
}

var _ = replay.Replayer(&GlIsQuery{}) // interface compliance check
func (ϟa *GlIsQuery) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_129_result := context             // Contextʳ
	ctx := GetContext_129_result                 // Contextʳ
	if key, remap := ϟa.Query.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Query.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Query.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlIsQuery)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_129_result, ctx
	return nil
}

var _ = replay.Replayer(&GlGetQueryiv{}) // interface compliance check
func (ϟa *GlGetQueryiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlGetQueryiv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

var _ = replay.Replayer(&GlGetQueryObjectuiv{}) // interface compliance check
func (ϟa *GlGetQueryObjectuiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	if key, remap := ϟa.Query.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Query.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Query.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlGetQueryObjectuiv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

var _ = replay.Replayer(&GlGenQueriesEXT{}) // interface compliance check
func (ϟa *GlGenQueriesEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	q := ϟa.Queries.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // QueryIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                  // Contextʳ
	GetContext_130_result := context                              // Contextʳ
	ctx := GetContext_130_result                                  // Contextʳ
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Queries.value())
	ϟb.Call(funcInfoGlGenQueriesEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		id := QueryId(ϟa.Queries.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)) // QueryId
		ctx.Instances.Queries[id] = func() *Query {
			s := &Query{}
			s.Init()
			return s
		}()
		q.Index(uint64(i), ϟs).replayWrite(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _ = q, context, GetContext_130_result, ctx
	return nil
}

var _ = replay.Replayer(&GlBeginQueryEXT{}) // interface compliance check
func (ϟa *GlBeginQueryEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U32(ϟa.Target))
	if key, remap := ϟa.Query.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Query.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Query.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlBeginQueryEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlEndQueryEXT{}) // interface compliance check
func (ϟa *GlEndQueryEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Call(funcInfoGlEndQueryEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlDeleteQueriesEXT{}) // interface compliance check
func (ϟa *GlDeleteQueriesEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	q := ϟa.Queries.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // QueryIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                  // Contextʳ
	GetContext_131_result := context                              // Contextʳ
	ctx := GetContext_131_result                                  // Contextʳ
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		ctx.Instances.Queries[q.Index(uint64(i), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)] = (*Query)(nil)
	}
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Push(ϟa.Queries.value())
	ϟb.Call(funcInfoGlDeleteQueriesEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = q, context, GetContext_131_result, ctx
	return nil
}

var _ = replay.Replayer(&GlIsQueryEXT{}) // interface compliance check
func (ϟa *GlIsQueryEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_132_result := context             // Contextʳ
	ctx := GetContext_132_result                 // Contextʳ
	if key, remap := ϟa.Query.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Query.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Query.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlIsQueryEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = context, GetContext_132_result, ctx
	return nil
}

var _ = replay.Replayer(&GlQueryCounterEXT{}) // interface compliance check
func (ϟa *GlQueryCounterEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	if key, remap := ϟa.Query.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Query.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Query.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Call(funcInfoGlQueryCounterEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&GlGetQueryivEXT{}) // interface compliance check
func (ϟa *GlGetQueryivEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlGetQueryivEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

var _ = replay.Replayer(&GlGetQueryObjectivEXT{}) // interface compliance check
func (ϟa *GlGetQueryObjectivEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	if key, remap := ϟa.Query.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Query.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Query.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlGetQueryObjectivEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

var _ = replay.Replayer(&GlGetQueryObjectuivEXT{}) // interface compliance check
func (ϟa *GlGetQueryObjectuivEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	if key, remap := ϟa.Query.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Query.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Query.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlGetQueryObjectuivEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

var _ = replay.Replayer(&GlGetQueryObjecti64vEXT{}) // interface compliance check
func (ϟa *GlGetQueryObjecti64vEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	if key, remap := ϟa.Query.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Query.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Query.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlGetQueryObjecti64vEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

var _ = replay.Replayer(&GlGetQueryObjectui64vEXT{}) // interface compliance check
func (ϟa *GlGetQueryObjectui64vEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	if key, remap := ϟa.Query.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Query.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Query.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlGetQueryObjectui64vEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

var _ = replay.Replayer(&ReplayCreateRenderer{}) // interface compliance check
func (ϟa *ReplayCreateRenderer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U32(ϟa.Id))
	ϟb.Call(funcInfoReplayCreateRenderer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&ReplayBindRenderer{}) // interface compliance check
func (ϟa *ReplayBindRenderer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U32(ϟa.Id))
	ϟb.Call(funcInfoReplayBindRenderer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&BackbufferInfo{}) // interface compliance check
func (ϟa *BackbufferInfo) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
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
	ϟb.Push(value.S32(ϟa.Width))
	ϟb.Push(value.S32(ϟa.Height))
	ϟb.Push(value.U32(ϟa.ColorFmt))
	ϟb.Push(value.U32(ϟa.DepthFmt))
	ϟb.Push(value.U32(ϟa.StencilFmt))
	ϟb.Push(value.Bool(ϟa.ResetViewportScissor))
	ϟb.Call(funcInfoBackbufferInfo)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _, _, _ = context, GetContext_133_result, ctx, backbuffer, color_id, color_buffer, depth_id, depth_buffer, stencil_id, stencil_buffer
	return nil
}

var _ = replay.Replayer(&StartTimer{}) // interface compliance check
func (ϟa *StartTimer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U8(ϟa.Index))
	ϟb.Call(funcInfoStartTimer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&StopTimer{}) // interface compliance check
func (ϟa *StopTimer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U8(ϟa.Index))
	ϟb.Call(funcInfoStopTimer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&FlushPostBuffer{}) // interface compliance check
func (ϟa *FlushPostBuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoFlushPostBuffer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (p Voidᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p U8ᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint8 {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p U8ᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint8 {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p U8ᵖ) replayWrite(value uint8, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p U8ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p Charᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) byte {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Charᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) byte {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Charᵖ) replayWrite(value byte, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p Charᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p EGLintᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) EGLint {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p EGLintᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) EGLint {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p EGLintᵖ) replayWrite(value EGLint, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p EGLintᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p Intᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) int64 {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Intᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) int64 {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Intᵖ) replayWrite(value int64, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p Intᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p CGLContextObjᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) CGLContextObj {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p CGLContextObjᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) CGLContextObj {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p CGLContextObjᵖ) replayWrite(value CGLContextObj, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p CGLContextObjᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p S32ᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) int32 {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p S32ᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) int32 {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p S32ᵖ) replayWrite(value int32, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p S32ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p U32ᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint32 {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p U32ᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint32 {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p U32ᵖ) replayWrite(value uint32, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p U32ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p DiscardFramebufferAttachmentᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) DiscardFramebufferAttachment {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p DiscardFramebufferAttachmentᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) DiscardFramebufferAttachment {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p DiscardFramebufferAttachmentᵖ) replayWrite(value DiscardFramebufferAttachment, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p DiscardFramebufferAttachmentᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p VertexArrayIdᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) VertexArrayId {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p VertexArrayIdᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) VertexArrayId {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p VertexArrayIdᵖ) replayWrite(value VertexArrayId, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p VertexArrayIdᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p ShaderAttribTypeᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) ShaderAttribType {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p ShaderAttribTypeᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) ShaderAttribType {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p ShaderAttribTypeᵖ) replayWrite(value ShaderAttribType, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p ShaderAttribTypeᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p ShaderUniformTypeᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) ShaderUniformType {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p ShaderUniformTypeᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) ShaderUniformType {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p ShaderUniformTypeᵖ) replayWrite(value ShaderUniformType, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p ShaderUniformTypeᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p F32ᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) float32 {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p F32ᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) float32 {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p F32ᵖ) replayWrite(value float32, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p F32ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p TextureIdᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) TextureId {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p TextureIdᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) TextureId {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p TextureIdᵖ) replayWrite(value TextureId, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p TextureIdᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p FramebufferIdᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) FramebufferId {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p FramebufferIdᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) FramebufferId {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p FramebufferIdᵖ) replayWrite(value FramebufferId, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p FramebufferIdᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p RenderbufferIdᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) RenderbufferId {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p RenderbufferIdᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) RenderbufferId {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p RenderbufferIdᵖ) replayWrite(value RenderbufferId, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p RenderbufferIdᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p BufferIdᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) BufferId {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p BufferIdᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) BufferId {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p BufferIdᵖ) replayWrite(value BufferId, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p BufferIdᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p Charᵖᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Charᵖ {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Charᵖᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Charᵖ {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Charᵖᵖ) replayWrite(value Charᵖ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p Charᵖᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p ShaderIdᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) ShaderId {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p ShaderIdᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) ShaderId {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p ShaderIdᵖ) replayWrite(value ShaderId, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p ShaderIdᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p Boolᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) bool {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Boolᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) bool {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Boolᵖ) replayWrite(value bool, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p Boolᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p FramebufferAttachmentᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) FramebufferAttachment {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p FramebufferAttachmentᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) FramebufferAttachment {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p FramebufferAttachmentᵖ) replayWrite(value FramebufferAttachment, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p FramebufferAttachmentᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p QueryIdᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) QueryId {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p QueryIdᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) QueryId {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p QueryIdᵖ) replayWrite(value QueryId, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p QueryIdᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p S64ᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) int64 {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p S64ᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) int64 {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p S64ᵖ) replayWrite(value int64, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p S64ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p U64ᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint64 {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p U64ᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint64 {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p U64ᵖ) replayWrite(value uint64, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p U64ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (s Boolˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Boolˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s Boolˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Boolˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
	return s
}
func (s Boolˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s Boolˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []bool {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s BufferIdˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) BufferIdˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(s.ElementSize(ϟs))
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
			if key, remap := v.remap(ϟa, ϟs); remap {
				loadRemap(ϟb, key, protocol.TypeUint32, v.value(ϟb, ϟa, ϟs))
			} else {
				ϟb.Push(v.value(ϟb, ϟa, ϟs))
			}
			ϟb.Store(ptr)
			ptr += step
		}
	}
	return s
}
func (s BufferIdˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) BufferIdˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
		size := s.ElementSize(ϟs)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(size)
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
			if key, remap := v.remap(ϟa, ϟs); remap {
				if _, found := ϟb.Remappings[key]; !found {
					dst := ϟb.AllocateMemory(size)
					ϟb.Load(protocol.TypeUint32, ptr)
					ϟb.Store(dst)
					ϟb.Remappings[key] = dst
				}
			}
			ptr += step
		}
	}
	return s
}
func (s BufferIdˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s BufferIdˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []BufferId {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s CGLContextObjˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) CGLContextObjˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(s.ElementSize(ϟs))
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
			ϟb.Push(v.value(ϟb, ϟa, ϟs))
			ϟb.Store(ptr)
			ptr += step
		}
	}
	return s
}
func (s CGLContextObjˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) CGLContextObjˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
	return s
}
func (s CGLContextObjˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s CGLContextObjˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []CGLContextObj {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s Charˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Charˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s Charˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Charˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
	return s
}
func (s Charˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s Charˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []byte {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s Charᵖˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Charᵖˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(s.ElementSize(ϟs))
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
			ϟb.Push(v.value())
			ϟb.Store(ptr)
			ptr += step
		}
	}
	return s
}
func (s Charᵖˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Charᵖˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
	return s
}
func (s Charᵖˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s Charᵖˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []Charᵖ {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s DiscardFramebufferAttachmentˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) DiscardFramebufferAttachmentˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s DiscardFramebufferAttachmentˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) DiscardFramebufferAttachmentˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
	return s
}
func (s DiscardFramebufferAttachmentˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s DiscardFramebufferAttachmentˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []DiscardFramebufferAttachment {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s EGLintˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) EGLintˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s EGLintˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) EGLintˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
	return s
}
func (s EGLintˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s EGLintˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []EGLint {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s F32ˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) F32ˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s F32ˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) F32ˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
	return s
}
func (s F32ˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s F32ˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []float32 {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s FramebufferAttachmentˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) FramebufferAttachmentˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s FramebufferAttachmentˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) FramebufferAttachmentˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
	return s
}
func (s FramebufferAttachmentˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s FramebufferAttachmentˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []FramebufferAttachment {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s FramebufferIdˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) FramebufferIdˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(s.ElementSize(ϟs))
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
			if key, remap := v.remap(ϟa, ϟs); remap {
				loadRemap(ϟb, key, protocol.TypeUint32, v.value(ϟb, ϟa, ϟs))
			} else {
				ϟb.Push(v.value(ϟb, ϟa, ϟs))
			}
			ϟb.Store(ptr)
			ptr += step
		}
	}
	return s
}
func (s FramebufferIdˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) FramebufferIdˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
		size := s.ElementSize(ϟs)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(size)
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
			if key, remap := v.remap(ϟa, ϟs); remap {
				if _, found := ϟb.Remappings[key]; !found {
					dst := ϟb.AllocateMemory(size)
					ϟb.Load(protocol.TypeUint32, ptr)
					ϟb.Store(dst)
					ϟb.Remappings[key] = dst
				}
			}
			ptr += step
		}
	}
	return s
}
func (s FramebufferIdˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s FramebufferIdˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []FramebufferId {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s Intˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Intˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s Intˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Intˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
	return s
}
func (s Intˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s Intˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []int64 {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s QueryIdˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) QueryIdˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(s.ElementSize(ϟs))
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
			if key, remap := v.remap(ϟa, ϟs); remap {
				loadRemap(ϟb, key, protocol.TypeUint32, v.value(ϟb, ϟa, ϟs))
			} else {
				ϟb.Push(v.value(ϟb, ϟa, ϟs))
			}
			ϟb.Store(ptr)
			ptr += step
		}
	}
	return s
}
func (s QueryIdˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) QueryIdˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
		size := s.ElementSize(ϟs)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(size)
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
			if key, remap := v.remap(ϟa, ϟs); remap {
				if _, found := ϟb.Remappings[key]; !found {
					dst := ϟb.AllocateMemory(size)
					ϟb.Load(protocol.TypeUint32, ptr)
					ϟb.Store(dst)
					ϟb.Remappings[key] = dst
				}
			}
			ptr += step
		}
	}
	return s
}
func (s QueryIdˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s QueryIdˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []QueryId {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s RenderbufferIdˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) RenderbufferIdˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(s.ElementSize(ϟs))
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
			if key, remap := v.remap(ϟa, ϟs); remap {
				loadRemap(ϟb, key, protocol.TypeUint32, v.value(ϟb, ϟa, ϟs))
			} else {
				ϟb.Push(v.value(ϟb, ϟa, ϟs))
			}
			ϟb.Store(ptr)
			ptr += step
		}
	}
	return s
}
func (s RenderbufferIdˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) RenderbufferIdˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
		size := s.ElementSize(ϟs)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(size)
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
			if key, remap := v.remap(ϟa, ϟs); remap {
				if _, found := ϟb.Remappings[key]; !found {
					dst := ϟb.AllocateMemory(size)
					ϟb.Load(protocol.TypeUint32, ptr)
					ϟb.Store(dst)
					ϟb.Remappings[key] = dst
				}
			}
			ptr += step
		}
	}
	return s
}
func (s RenderbufferIdˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s RenderbufferIdˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []RenderbufferId {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s S32ˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S32ˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s S32ˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S32ˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
	return s
}
func (s S32ˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s S32ˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []int32 {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s S64ˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S64ˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s S64ˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S64ˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
	return s
}
func (s S64ˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s S64ˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []int64 {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s ShaderAttribTypeˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) ShaderAttribTypeˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s ShaderAttribTypeˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) ShaderAttribTypeˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
	return s
}
func (s ShaderAttribTypeˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s ShaderAttribTypeˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []ShaderAttribType {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s ShaderIdˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) ShaderIdˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(s.ElementSize(ϟs))
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
			if key, remap := v.remap(ϟa, ϟs); remap {
				loadRemap(ϟb, key, protocol.TypeUint32, v.value(ϟb, ϟa, ϟs))
			} else {
				ϟb.Push(v.value(ϟb, ϟa, ϟs))
			}
			ϟb.Store(ptr)
			ptr += step
		}
	}
	return s
}
func (s ShaderIdˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) ShaderIdˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
		size := s.ElementSize(ϟs)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(size)
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
			if key, remap := v.remap(ϟa, ϟs); remap {
				if _, found := ϟb.Remappings[key]; !found {
					dst := ϟb.AllocateMemory(size)
					ϟb.Load(protocol.TypeUint32, ptr)
					ϟb.Store(dst)
					ϟb.Remappings[key] = dst
				}
			}
			ptr += step
		}
	}
	return s
}
func (s ShaderIdˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s ShaderIdˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []ShaderId {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s ShaderUniformTypeˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) ShaderUniformTypeˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s ShaderUniformTypeˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) ShaderUniformTypeˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
	return s
}
func (s ShaderUniformTypeˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s ShaderUniformTypeˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []ShaderUniformType {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s TextureIdˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) TextureIdˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(s.ElementSize(ϟs))
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
			if key, remap := v.remap(ϟa, ϟs); remap {
				loadRemap(ϟb, key, protocol.TypeUint32, v.value(ϟb, ϟa, ϟs))
			} else {
				ϟb.Push(v.value(ϟb, ϟa, ϟs))
			}
			ϟb.Store(ptr)
			ptr += step
		}
	}
	return s
}
func (s TextureIdˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) TextureIdˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
		size := s.ElementSize(ϟs)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(size)
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
			if key, remap := v.remap(ϟa, ϟs); remap {
				if _, found := ϟb.Remappings[key]; !found {
					dst := ϟb.AllocateMemory(size)
					ϟb.Load(protocol.TypeUint32, ptr)
					ϟb.Store(dst)
					ϟb.Remappings[key] = dst
				}
			}
			ptr += step
		}
	}
	return s
}
func (s TextureIdˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s TextureIdˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []TextureId {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s U32ˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U32ˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s U32ˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U32ˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
	return s
}
func (s U32ˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s U32ˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []uint32 {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s U64ˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U64ˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s U64ˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U64ˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
	return s
}
func (s U64ˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s U64ˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []uint64 {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s U8ˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U8ˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s U8ˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U8ˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
	return s
}
func (s U8ˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s U8ˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []uint8 {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s VertexArrayIdˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) VertexArrayIdˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(s.ElementSize(ϟs))
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
			if key, remap := v.remap(ϟa, ϟs); remap {
				loadRemap(ϟb, key, protocol.TypeUint32, v.value(ϟb, ϟa, ϟs))
			} else {
				ϟb.Push(v.value(ϟb, ϟa, ϟs))
			}
			ϟb.Store(ptr)
			ptr += step
		}
	}
	return s
}
func (s VertexArrayIdˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) VertexArrayIdˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
		size := s.ElementSize(ϟs)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(size)
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
			if key, remap := v.remap(ϟa, ϟs); remap {
				if _, found := ϟb.Remappings[key]; !found {
					dst := ϟb.AllocateMemory(size)
					ϟb.Load(protocol.TypeUint32, ptr)
					ϟb.Store(dst)
					ϟb.Remappings[key] = dst
				}
			}
			ptr += step
		}
	}
	return s
}
func (s VertexArrayIdˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s VertexArrayIdˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []VertexArrayId {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s Voidˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Voidˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s Voidˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Voidˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
	return s
}
func (s Voidˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
