/*
* Copyright 2014, The Android Open Source Project
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/

////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

#ifndef ANDROID_CAZE_GL_FUNCTIONS_H
#define ANDROID_CAZE_GL_FUNCTIONS_H

#include <stdint.h>

namespace android {
namespace caze {

class Interpreter;

class GlFunctions {
public:
    // Register the GL functions to the interpreter
    static void Register(Interpreter* interpreter);

    // List of the function ids for the GL functions. This list has to be consistent with the
    // function ids on the server side because they are part of the communication protocol
    enum FunctionIds : uint16_t {
        Init = 0,
        EglCreateContext = 1,
        EglMakeCurrent = 2,
        EglSwapBuffers = 3,
        GlEnableClientState = 4,
        GlDisableClientState = 5,
        GlGetProgramBinaryOES = 6,
        GlProgramBinaryOES = 7,
        GlStartTilingQCOM = 8,
        GlEndTilingQCOM = 9,
        GlDiscardFramebufferEXT = 10,
        GlInsertEventMarkerEXT = 11,
        GlPushGroupMarkerEXT = 12,
        GlPopGroupMarkerEXT = 13,
        GlTexStorage1DEXT = 14,
        GlTexStorage2DEXT = 15,
        GlTexStorage3DEXT = 16,
        GlTextureStorage1DEXT = 17,
        GlTextureStorage2DEXT = 18,
        GlTextureStorage3DEXT = 19,
        GlGenVertexArraysOES = 20,
        GlBindVertexArrayOES = 21,
        GlDeleteVertexArraysOES = 22,
        GlIsVertexArrayOES = 23,
        GlEGLImageTargetTexture2DOES = 24,
        GlEGLImageTargetRenderbufferStorageOES = 25,
        GlGetGraphicsResetStatusEXT = 26,
        GlBindAttribLocation = 27,
        GlBlendFunc = 28,
        GlBlendFuncSeparate = 29,
        GlBlendEquation = 30,
        GlBlendEquationSeparate = 31,
        GlBlendColor = 32,
        GlEnableVertexAttribArray = 33,
        GlDisableVertexAttribArray = 34,
        GlVertexAttribPointer = 35,
        GlGetActiveAttrib = 36,
        GlGetActiveUniform = 37,
        GlGetError = 38,
        GlGetProgramiv = 39,
        GlGetShaderiv = 40,
        GlGetUniformLocation = 41,
        GlGetAttribLocation = 42,
        GlPixelStorei = 43,
        GlTexParameteri = 44,
        GlTexParameterf = 45,
        GlGetTexParameteriv = 46,
        GlGetTexParameterfv = 47,
        GlUniform1i = 48,
        GlUniform2i = 49,
        GlUniform3i = 50,
        GlUniform4i = 51,
        GlUniform1iv = 52,
        GlUniform2iv = 53,
        GlUniform3iv = 54,
        GlUniform4iv = 55,
        GlUniform1f = 56,
        GlUniform2f = 57,
        GlUniform3f = 58,
        GlUniform4f = 59,
        GlUniform1fv = 60,
        GlUniform2fv = 61,
        GlUniform3fv = 62,
        GlUniform4fv = 63,
        GlUniformMatrix2fv = 64,
        GlUniformMatrix3fv = 65,
        GlUniformMatrix4fv = 66,
        GlGetUniformfv = 67,
        GlGetUniformiv = 68,
        GlVertexAttrib1f = 69,
        GlVertexAttrib2f = 70,
        GlVertexAttrib3f = 71,
        GlVertexAttrib4f = 72,
        GlVertexAttrib1fv = 73,
        GlVertexAttrib2fv = 74,
        GlVertexAttrib3fv = 75,
        GlVertexAttrib4fv = 76,
        GlGetShaderPrecisionFormat = 77,
        GlDepthMask = 78,
        GlDepthFunc = 79,
        GlDepthRangef = 80,
        GlColorMask = 81,
        GlStencilMask = 82,
        GlStencilMaskSeparate = 83,
        GlStencilFuncSeparate = 84,
        GlStencilOpSeparate = 85,
        GlFrontFace = 86,
        GlViewport = 87,
        GlScissor = 88,
        GlActiveTexture = 89,
        GlGenTextures = 90,
        GlDeleteTextures = 91,
        GlIsTexture = 92,
        GlBindTexture = 93,
        GlTexImage2D = 94,
        GlTexSubImage2D = 95,
        GlCopyTexImage2D = 96,
        GlCopyTexSubImage2D = 97,
        GlCompressedTexImage2D = 98,
        GlCompressedTexSubImage2D = 99,
        GlGenerateMipmap = 100,
        GlReadPixels = 101,
        GlGenFramebuffers = 102,
        GlBindFramebuffer = 103,
        GlCheckFramebufferStatus = 104,
        GlDeleteFramebuffers = 105,
        GlIsFramebuffer = 106,
        GlGenRenderbuffers = 107,
        GlBindRenderbuffer = 108,
        GlRenderbufferStorage = 109,
        GlDeleteRenderbuffers = 110,
        GlIsRenderbuffer = 111,
        GlGetRenderbufferParameteriv = 112,
        GlGenBuffers = 113,
        GlBindBuffer = 114,
        GlBufferData = 115,
        GlBufferSubData = 116,
        GlDeleteBuffers = 117,
        GlIsBuffer = 118,
        GlGetBufferParameteriv = 119,
        GlCreateShader = 120,
        GlDeleteShader = 121,
        GlShaderSource = 122,
        GlShaderBinary = 123,
        GlGetShaderInfoLog = 124,
        GlGetShaderSource = 125,
        GlReleaseShaderCompiler = 126,
        GlCompileShader = 127,
        GlIsShader = 128,
        GlCreateProgram = 129,
        GlDeleteProgram = 130,
        GlAttachShader = 131,
        GlDetachShader = 132,
        GlGetAttachedShaders = 133,
        GlLinkProgram = 134,
        GlGetProgramInfoLog = 135,
        GlUseProgram = 136,
        GlIsProgram = 137,
        GlValidateProgram = 138,
        GlClearColor = 139,
        GlClearDepthf = 140,
        GlClearStencil = 141,
        GlClear = 142,
        GlCullFace = 143,
        GlPolygonOffset = 144,
        GlLineWidth = 145,
        GlSampleCoverage = 146,
        GlHint = 147,
        GlFramebufferRenderbuffer = 148,
        GlFramebufferTexture2D = 149,
        GlGetFramebufferAttachmentParameteriv = 150,
        GlDrawElements = 151,
        GlDrawArrays = 152,
        GlFlush = 153,
        GlFinish = 154,
        GlGetBooleanv = 155,
        GlGetFloatv = 156,
        GlGetIntegerv = 157,
        GlGetString = 158,
        GlEnable = 159,
        GlDisable = 160,
        GlIsEnabled = 161,
        GlMapBufferRange = 162,
        GlUnmapBuffer = 163,
        GlInvalidateFramebuffer = 164,
        GlRenderbufferStorageMultisample = 165,
        GlBlitFramebuffer = 166,
    };
};

}  // end of namespace caze
}  // end of namespace android

#endif  // ANDROID_CAZE_GL_FUNCTIONS_H
