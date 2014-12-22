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

#include "GlFunctions.h"
#include "GlInclude.h"
#include "Interpreter.h"
#include "Log.h"
#include "Stack.h"

namespace android {
namespace caze {
namespace {

typedef GLenum DrawMode;
typedef GLenum IndicesType;
typedef GLenum TextureTarget_GLES_1_1;
typedef GLenum TextureTarget_GLES_2_0;
typedef GLenum TextureTarget_OES_EGL_image_external;
typedef GLenum TextureTarget;
typedef GLenum CubeMapImageTarget;
typedef GLenum Texture2DImageTarget;
typedef GLenum TextureImageTarget;
typedef GLenum TexelFormat_GLES_1_1;
typedef GLenum TexelFormat_GLES_3_0;
typedef GLenum TexelFormat;
typedef GLenum ReadPixelsFormat;
typedef GLenum RenderbufferFormat;
typedef GLenum Type_OES_vertex_half_float;
typedef GLenum CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture;
typedef GLenum CompressedTexelFormat_AMD_compressed_ATC_texture;
typedef GLenum CompressedTexelFormat;
typedef GLenum ImageTexelFormat;
typedef GLenum TexelType;
typedef GLenum FramebufferAttachment;
typedef GLenum FramebufferAttachmentType;
typedef GLenum FramebufferTarget_GLES_2_0;
typedef GLenum FramebufferTarget_GLES_3_1;
typedef GLenum FramebufferTarget;
typedef GLenum FramebufferAttachmentParameter;
typedef GLenum FramebufferStatus;
typedef GLenum RenderbufferTarget;
typedef GLenum RenderbufferParameter;
typedef GLenum BufferTarget;
typedef GLenum BufferParameter;
typedef GLenum TextureUnit;
typedef GLenum BufferUsage;
typedef GLenum ShaderType;
typedef GLenum StateVariable_GLES_2_0;
typedef GLenum StateVariable_GLES_3_1;
typedef GLenum StateVariable_EXT_texture_filter_anisotropic;
typedef GLenum StateVariable;
typedef GLenum FaceMode;
typedef GLenum ArrayType_GLES_1_1;
typedef GLenum ArrayType_OES_point_size_array;
typedef GLenum ArrayType;
typedef GLenum Capability;
typedef GLenum StringConstant;
typedef GLenum VertexAttribType;
typedef GLenum ShaderAttribType;
typedef GLenum ShaderUniformType;
typedef GLenum Error;
typedef GLenum HintTarget;
typedef GLenum HintMode;
typedef GLenum DiscardFramebufferAttachment;
typedef GLenum ProgramParameter;
typedef GLenum ShaderParameter;
typedef GLenum PixelStoreParameter;
typedef GLenum TextureParameter_FilterMode;
typedef GLenum TextureParameter_WrapMode;
typedef GLenum TextureParameter_EXT_texture_filter_anisotropic;
typedef GLenum TextureParameter_SwizzleMode;
typedef GLenum TextureParameter;
typedef GLenum TextureFilterMode;
typedef GLenum TextureWrapMode;
typedef GLenum TexelComponent;
typedef GLenum BlendFactor;
typedef GLenum PrecisionType;
typedef GLenum TestFunction;
typedef GLenum StencilAction;
typedef GLenum FaceOrientation;
typedef GLenum BlendEquation;
typedef GLenum MapBufferTarget;
typedef GLenum ImageTargetTexture_OES_EGL_image;
typedef GLenum ImageTargetTexture_OES_EGL_image_external;
typedef GLenum ImageTargetTexture;
typedef GLenum ImageTargetRenderbufferStorage;
typedef GLenum ResetStatus;
typedef GLenum TextureKind;
typedef GLenum VertexAttribSize;
typedef GLuint TilePreserveMaskQCOM;
typedef GLuint ClearMask;
typedef GLuint MapBufferRangeAccess;
typedef GLenum SyncCondition;
typedef GLenum ClientWaitSyncSignal;
typedef GLuint SyncFlags;

typedef GLuint ColorId;
typedef GLuint RectId;
typedef GLuint ImageId;
typedef GLuint FramebufferAttachableId;
typedef GLuint RenderbufferId;
typedef GLuint TextureId;
typedef GLuint CubemapLevelId;
typedef GLuint FramebufferAttachmentInfoId;
typedef GLuint FramebufferId;
typedef GLuint BufferId;
typedef GLuint ShaderId;
typedef GLuint VertexAttributeId;
typedef GLuint Vec2iId;
typedef GLuint Vec3iId;
typedef GLuint Vec4iId;
typedef GLuint Vec2fId;
typedef GLuint Vec3fId;
typedef GLuint Vec4fId;
typedef GLuint Mat2fId;
typedef GLuint Mat3fId;
typedef GLuint Mat4fId;
typedef GLuint UniformId;
typedef GLuint ProgramId;
typedef GLuint VertexArrayId;
typedef GLuint VertexAttributeArrayId;
typedef GLuint BlendStateId;
typedef GLuint RasterizerStateId;
typedef GLuint ClearStateId;
typedef GLuint InternalStateId;
typedef GLuint ObjectsId;
typedef GLuint SyncObjectId;

typedef int32_t s32;
typedef uint32_t u32;

#ifdef EGL_VERSION_1_0
inline void glUnmapBuffer(GLenum target) { CAZE_DEBUG("WARNING: stub called: glUnmapBuffer\n"); }
#else
inline void glClearDepthf(GLfloat depth) { glClearDepth(depth); }
inline void glDepthRangef(GLfloat near, GLfloat far) { glDepthRange(near, far); }
inline void glShaderBinary(GLsizei n, const GLuint* shaders, GLenum binaryformat,
                           const void* binary, GLsizei length) {
    CAZE_DEBUG("WARNING: stub called: glShaderBinary\n");
}
inline void glReleaseShaderCompiler() {
    CAZE_DEBUG("WARNING: stub called: glReleaseShaderCompiler\n");
}
#endif
inline void glPopGroupMarkerEXT() { CAZE_DEBUG("WARNING: stub called: glPopGroupMarkerEXT\n"); }
inline void glInsertEventMarkerEXT(GLint length, const GLchar* marker) {
    CAZE_DEBUG("WARNING: stub called: glInsertEventMarkerEXT\n");
}
inline void glPushGroupMarkerEXT(GLint length, const GLchar* marker) {
    CAZE_DEBUG("WARNING: stub called: glPushGroupMarkerEXT\n");
}
// EGL APIs
inline void eglMakeCurrent(int context) { CAZE_DEBUG("WARNING: stub called: eglMakeCurrent\n"); }
inline void eglCreateContext(int* version, int* context) {
    CAZE_DEBUG("WARNING: stub called: eglCreateContext\n");
}
inline void eglSwapBuffers() { CAZE_DEBUG("WARNING: stub called: eglSwapBuffers\n"); }
inline void glGetShaderPrecisionFormat(GLenum shaderType, GLenum precisionType) {
    CAZE_DEBUG("WARNING: stub called: glGetShaderPrecisionFormat\n");
}
// GLES 1 extension APIs
inline void glStartTilingQCOM(GLuint x, GLuint y, GLuint width, GLuint height,
                              TilePreserveMaskQCOM preserveMask) {
    CAZE_DEBUG("WARNING: stub called: glStartTilingQCOM\n");
}
inline void glEndTilingQCOM(TilePreserveMaskQCOM preserveMask) {
    CAZE_DEBUG("WARNING: stub called: glEndTilingQCOM\n");
}
inline void glEGLImageTargetTexture2DOES(GLenum target, const GLvoid* image) {
    CAZE_DEBUG("WARNING: stub called: glEGLImageTargetTexture2DOES\n");
}
inline void glEGLImageTargetRenderbufferStorageOES(GLenum target, const GLvoid* image) {
    CAZE_DEBUG("WARNING: stub called: glEGLImageTargetRenderbufferStorageOES\n");
}
inline void glDiscardFramebufferEXT(GLenum target, int numAttachments,
                                    const DiscardFramebufferAttachment* attachments) {
    CAZE_DEBUG("WARNING: stub called: glDiscardFramebufferEXT\n");
}
inline void glGetBooleanv(GLenum pname, bool* params) {
    CAZE_DEBUG("WARNING: stub called: glGetBooleanv\n");
}
inline int glMapBufferRange(GLenum target, int, int, MapBufferRangeAccess) {
    CAZE_DEBUG("WARNING: stub called: glMapBufferRange\n");
    return -1;
}
inline void glGenVertexArraysOES(GLsizei n, GLuint* arrays) {
    CAZE_DEBUG("WARNING: stub called: glGenVertexArraysOES\n");
}
inline void glBindVertexArrayOES(GLuint array) {
    CAZE_DEBUG("WARNING: stub called: glBindVertexArrayOES\n");
}
inline void glDeleteVertexArraysOES(GLsizei n, const GLuint* arrays) {
    CAZE_DEBUG("WARNING: stub called: glDeleteVertexArraysOES\n");
}
inline GLboolean glIsVertexArrayOES(GLuint array) {
    CAZE_DEBUG("WARNING: stub called: glIsVertexArrayOES\n");
    return 0;
}
inline void glGetProgramBinaryOES(GLuint program, GLsizei bufSize, GLsizei* length,
                                  GLenum* binaryFormat, void* binary) {
    CAZE_DEBUG("WARNING: stub called: glGetProgramBinaryOES\n");
}
inline void glProgramBinaryOES(GLuint program, GLenum binaryFormat, const void* binary,
                               int length) {
    CAZE_DEBUG("WARNING: stub called: glProgramBinaryOES\n");
}
inline GLenum glGetGraphicsResetStatusEXT() {
    CAZE_DEBUG("WARNING: stub called: glGetGraphicsResetStatusEXT\n");
    return 0;
}
inline void glInvalidateFramebuffer(GLenum target, GLsizei numAttachments,
                                    const GLenum* attachments) {
    CAZE_DEBUG("WARNING: stub called: glInvalidateFramebuffer\n");
}
inline void glTexStorage1DEXT(GLenum target, GLsizei levels, GLenum internalformat, GLsizei width) {
    CAZE_DEBUG("WARNING: stub called: glTexStorage1DEXT\n");
}
inline void glTexStorage2DEXT(GLenum target, GLsizei levels, GLenum internalformat, GLsizei width,
                              GLsizei height) {
    CAZE_DEBUG("WARNING: stub called: glTexStorage2DEXT\n");
}
inline void glTexStorage3DEXT(GLenum target, GLsizei levels, GLenum internalformat, GLsizei width,
                              GLsizei height, GLsizei depth) {
    CAZE_DEBUG("WARNING: stub called: glTexStorage3DEXT\n");
}
inline void glTextureStorage1DEXT(GLuint texture, GLenum target, GLsizei levels,
                                  GLenum internalformat, GLsizei width) {
    CAZE_DEBUG("WARNING: stub called: glTextureStorage1DEXT\n");
}
inline void glTextureStorage2DEXT(GLuint texture, GLenum target, GLsizei levels,
                                  GLenum internalformat, GLsizei width, GLsizei height) {
    CAZE_DEBUG("WARNING: stub called: glTextureStorage2DEXT\n");
}
inline void glTextureStorage3DEXT(GLuint texture, GLenum target, GLsizei levels,
                                  GLenum internalformat, GLsizei width, GLsizei height,
                                  GLsizei depth) {
    CAZE_DEBUG("WARNING: stub called: glTextureStorage3DEXT\n");
}

bool handleEglCreateContext(Stack* stack, bool pushReturn) {
    s32* context = stack->pop<s32*>();
    s32* version = stack->pop<s32*>();
    if (stack->isValid()) {
        CAZE_DEBUG("eglCreateContext(%p, %p)\n", version, context);
        eglCreateContext(version, context);
        return true;
    } else {
        CAZE_WARNING("Error during calling function eglCreateContext\n");
        return false;
    }
}

bool handleEglMakeCurrent(Stack* stack, bool pushReturn) {
    s32 context = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("eglMakeCurrent(%d)\n", context);
        eglMakeCurrent(context);
        return true;
    } else {
        CAZE_WARNING("Error during calling function eglMakeCurrent\n");
        return false;
    }
}

bool handleEglSwapBuffers(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        CAZE_DEBUG("eglSwapBuffers()\n");
        eglSwapBuffers();
        return true;
    } else {
        CAZE_WARNING("Error during calling function eglSwapBuffers\n");
        return false;
    }
}

bool handleGlEnableClientState(Stack* stack, bool pushReturn) {
    ArrayType type = stack->pop<ArrayType>();
    if (stack->isValid()) {
        CAZE_DEBUG("glEnableClientState(%u)\n", type);
        glEnableClientState(type);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glEnableClientState\n");
        return false;
    }
}

bool handleGlDisableClientState(Stack* stack, bool pushReturn) {
    ArrayType type = stack->pop<ArrayType>();
    if (stack->isValid()) {
        CAZE_DEBUG("glDisableClientState(%u)\n", type);
        glDisableClientState(type);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glDisableClientState\n");
        return false;
    }
}

bool handleGlGetProgramBinaryOES(Stack* stack, bool pushReturn) {
    void* binary = stack->pop<void*>();
    u32* binary_format = stack->pop<u32*>();
    s32* bytes_written = stack->pop<s32*>();
    s32 buffer_size = stack->pop<s32>();
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glGetProgramBinaryOES(%u, %d, %p, %p, %p)\n", program, buffer_size,
                   bytes_written, binary_format, binary);
        glGetProgramBinaryOES(program, buffer_size, bytes_written, binary_format, binary);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glGetProgramBinaryOES\n");
        return false;
    }
}

bool handleGlProgramBinaryOES(Stack* stack, bool pushReturn) {
    s32 binary_size = stack->pop<s32>();
    void* binary = stack->pop<void*>();
    u32 binary_format = stack->pop<u32>();
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glProgramBinaryOES(%u, %u, %p, %d)\n", program, binary_format, binary,
                   binary_size);
        glProgramBinaryOES(program, binary_format, binary, binary_size);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glProgramBinaryOES\n");
        return false;
    }
}

bool handleGlStartTilingQCOM(Stack* stack, bool pushReturn) {
    TilePreserveMaskQCOM preserveMask = stack->pop<TilePreserveMaskQCOM>();
    s32 height = stack->pop<s32>();
    s32 width = stack->pop<s32>();
    s32 y = stack->pop<s32>();
    s32 x = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glStartTilingQCOM(%d, %d, %d, %d, %u)\n", x, y, width, height, preserveMask);
        glStartTilingQCOM(x, y, width, height, preserveMask);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glStartTilingQCOM\n");
        return false;
    }
}

bool handleGlEndTilingQCOM(Stack* stack, bool pushReturn) {
    TilePreserveMaskQCOM preserve_mask = stack->pop<TilePreserveMaskQCOM>();
    if (stack->isValid()) {
        CAZE_DEBUG("glEndTilingQCOM(%u)\n", preserve_mask);
        glEndTilingQCOM(preserve_mask);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glEndTilingQCOM\n");
        return false;
    }
}

bool handleGlDiscardFramebufferEXT(Stack* stack, bool pushReturn) {
    DiscardFramebufferAttachment* attachments = stack->pop<DiscardFramebufferAttachment*>();
    s32 numAttachments = stack->pop<s32>();
    FramebufferTarget target = stack->pop<FramebufferTarget>();
    if (stack->isValid()) {
        CAZE_DEBUG("glDiscardFramebufferEXT(%u, %d, %p)\n", target, numAttachments, attachments);
        glDiscardFramebufferEXT(target, numAttachments, attachments);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glDiscardFramebufferEXT\n");
        return false;
    }
}

bool handleGlInsertEventMarkerEXT(Stack* stack, bool pushReturn) {
    char* marker = stack->pop<char*>();
    s32 length = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glInsertEventMarkerEXT(%d, %s)\n", length, marker);
        glInsertEventMarkerEXT(length, marker);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glInsertEventMarkerEXT\n");
        return false;
    }
}

bool handleGlPushGroupMarkerEXT(Stack* stack, bool pushReturn) {
    char* marker = stack->pop<char*>();
    s32 length = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glPushGroupMarkerEXT(%d, %s)\n", length, marker);
        glPushGroupMarkerEXT(length, marker);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glPushGroupMarkerEXT\n");
        return false;
    }
}

bool handleGlPopGroupMarkerEXT(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        CAZE_DEBUG("glPopGroupMarkerEXT()\n");
        glPopGroupMarkerEXT();
        return true;
    } else {
        CAZE_WARNING("Error during calling function glPopGroupMarkerEXT\n");
        return false;
    }
}

bool handleGlTexStorage1DEXT(Stack* stack, bool pushReturn) {
    s32 width = stack->pop<s32>();
    TexelFormat format = stack->pop<TexelFormat>();
    s32 levels = stack->pop<s32>();
    TextureTarget target = stack->pop<TextureTarget>();
    if (stack->isValid()) {
        CAZE_DEBUG("glTexStorage1DEXT(%u, %d, %u, %d)\n", target, levels, format, width);
        glTexStorage1DEXT(target, levels, format, width);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glTexStorage1DEXT\n");
        return false;
    }
}

bool handleGlTexStorage2DEXT(Stack* stack, bool pushReturn) {
    s32 height = stack->pop<s32>();
    s32 width = stack->pop<s32>();
    TexelFormat format = stack->pop<TexelFormat>();
    s32 levels = stack->pop<s32>();
    TextureTarget target = stack->pop<TextureTarget>();
    if (stack->isValid()) {
        CAZE_DEBUG("glTexStorage2DEXT(%u, %d, %u, %d, %d)\n", target, levels, format, width,
                   height);
        glTexStorage2DEXT(target, levels, format, width, height);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glTexStorage2DEXT\n");
        return false;
    }
}

bool handleGlTexStorage3DEXT(Stack* stack, bool pushReturn) {
    s32 depth = stack->pop<s32>();
    s32 height = stack->pop<s32>();
    s32 width = stack->pop<s32>();
    TexelFormat format = stack->pop<TexelFormat>();
    s32 levels = stack->pop<s32>();
    TextureTarget target = stack->pop<TextureTarget>();
    if (stack->isValid()) {
        CAZE_DEBUG("glTexStorage3DEXT(%u, %d, %u, %d, %d, %d)\n", target, levels, format, width,
                   height, depth);
        glTexStorage3DEXT(target, levels, format, width, height, depth);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glTexStorage3DEXT\n");
        return false;
    }
}

bool handleGlTextureStorage1DEXT(Stack* stack, bool pushReturn) {
    s32 width = stack->pop<s32>();
    TexelFormat format = stack->pop<TexelFormat>();
    s32 levels = stack->pop<s32>();
    TextureTarget target = stack->pop<TextureTarget>();
    TextureId texture = stack->pop<TextureId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glTextureStorage1DEXT(%u, %u, %d, %u, %d)\n", texture, target, levels, format,
                   width);
        glTextureStorage1DEXT(texture, target, levels, format, width);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glTextureStorage1DEXT\n");
        return false;
    }
}

bool handleGlTextureStorage2DEXT(Stack* stack, bool pushReturn) {
    s32 height = stack->pop<s32>();
    s32 width = stack->pop<s32>();
    TexelFormat format = stack->pop<TexelFormat>();
    s32 levels = stack->pop<s32>();
    TextureTarget target = stack->pop<TextureTarget>();
    TextureId texture = stack->pop<TextureId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glTextureStorage2DEXT(%u, %u, %d, %u, %d, %d)\n", texture, target, levels,
                   format, width, height);
        glTextureStorage2DEXT(texture, target, levels, format, width, height);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glTextureStorage2DEXT\n");
        return false;
    }
}

bool handleGlTextureStorage3DEXT(Stack* stack, bool pushReturn) {
    s32 depth = stack->pop<s32>();
    s32 height = stack->pop<s32>();
    s32 width = stack->pop<s32>();
    TexelFormat format = stack->pop<TexelFormat>();
    s32 levels = stack->pop<s32>();
    TextureTarget target = stack->pop<TextureTarget>();
    TextureId texture = stack->pop<TextureId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glTextureStorage3DEXT(%u, %u, %d, %u, %d, %d, %d)\n", texture, target, levels,
                   format, width, height, depth);
        glTextureStorage3DEXT(texture, target, levels, format, width, height, depth);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glTextureStorage3DEXT\n");
        return false;
    }
}

bool handleGlGenVertexArraysOES(Stack* stack, bool pushReturn) {
    VertexArrayId* arrays = stack->pop<VertexArrayId*>();
    s32 count = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glGenVertexArraysOES(%d, %p)\n", count, arrays);
        glGenVertexArraysOES(count, arrays);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glGenVertexArraysOES\n");
        return false;
    }
}

bool handleGlBindVertexArrayOES(Stack* stack, bool pushReturn) {
    VertexArrayId array = stack->pop<VertexArrayId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glBindVertexArrayOES(%u)\n", array);
        glBindVertexArrayOES(array);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glBindVertexArrayOES\n");
        return false;
    }
}

bool handleGlDeleteVertexArraysOES(Stack* stack, bool pushReturn) {
    VertexArrayId* arrays = stack->pop<VertexArrayId*>();
    s32 count = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glDeleteVertexArraysOES(%d, %p)\n", count, arrays);
        glDeleteVertexArraysOES(count, arrays);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glDeleteVertexArraysOES\n");
        return false;
    }
}

bool handleGlIsVertexArrayOES(Stack* stack, bool pushReturn) {
    VertexArrayId array = stack->pop<VertexArrayId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glIsVertexArrayOES(%u)\n", array);
        const bool return_value = glIsVertexArrayOES(array);
        if (pushReturn) {
            stack->push<const bool>(return_value);
        }
        return true;
    } else {
        CAZE_WARNING("Error during calling function glIsVertexArrayOES\n");
        return false;
    }
}

bool handleGlEGLImageTargetTexture2DOES(Stack* stack, bool pushReturn) {
    void* image = stack->pop<void*>();
    ImageTargetTexture target = stack->pop<ImageTargetTexture>();
    if (stack->isValid()) {
        CAZE_DEBUG("glEGLImageTargetTexture2DOES(%u, %p)\n", target, image);
        glEGLImageTargetTexture2DOES(target, image);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glEGLImageTargetTexture2DOES\n");
        return false;
    }
}

bool handleGlEGLImageTargetRenderbufferStorageOES(Stack* stack, bool pushReturn) {
    void* image = stack->pop<void*>();
    ImageTargetRenderbufferStorage target = stack->pop<ImageTargetRenderbufferStorage>();
    if (stack->isValid()) {
        CAZE_DEBUG("glEGLImageTargetRenderbufferStorageOES(%u, %p)\n", target, image);
        glEGLImageTargetRenderbufferStorageOES(target, image);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glEGLImageTargetRenderbufferStorageOES\n");
        return false;
    }
}

bool handleGlGetGraphicsResetStatusEXT(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        CAZE_DEBUG("glGetGraphicsResetStatusEXT()\n");
        const ResetStatus return_value = glGetGraphicsResetStatusEXT();
        if (pushReturn) {
            stack->push<const ResetStatus>(return_value);
        }
        return true;
    } else {
        CAZE_WARNING("Error during calling function glGetGraphicsResetStatusEXT\n");
        return false;
    }
}

bool handleGlBindAttribLocation(Stack* stack, bool pushReturn) {
    char* name = stack->pop<char*>();
    s32 index = stack->pop<s32>();
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glBindAttribLocation(%u, %d, %s)\n", program, index, name);
        glBindAttribLocation(program, index, name);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glBindAttribLocation\n");
        return false;
    }
}

bool handleGlBlendFunc(Stack* stack, bool pushReturn) {
    BlendFactor dst_factor = stack->pop<BlendFactor>();
    BlendFactor src_factor = stack->pop<BlendFactor>();
    if (stack->isValid()) {
        CAZE_DEBUG("glBlendFunc(%u, %u)\n", src_factor, dst_factor);
        glBlendFunc(src_factor, dst_factor);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glBlendFunc\n");
        return false;
    }
}

bool handleGlBlendFuncSeparate(Stack* stack, bool pushReturn) {
    BlendFactor dst_factor_alpha = stack->pop<BlendFactor>();
    BlendFactor src_factor_alpha = stack->pop<BlendFactor>();
    BlendFactor dst_factor_rgb = stack->pop<BlendFactor>();
    BlendFactor src_factor_rgb = stack->pop<BlendFactor>();
    if (stack->isValid()) {
        CAZE_DEBUG("glBlendFuncSeparate(%u, %u, %u, %u)\n", src_factor_rgb, dst_factor_rgb,
                   src_factor_alpha, dst_factor_alpha);
        glBlendFuncSeparate(src_factor_rgb, dst_factor_rgb, src_factor_alpha, dst_factor_alpha);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glBlendFuncSeparate\n");
        return false;
    }
}

bool handleGlBlendEquation(Stack* stack, bool pushReturn) {
    BlendEquation equation = stack->pop<BlendEquation>();
    if (stack->isValid()) {
        CAZE_DEBUG("glBlendEquation(%u)\n", equation);
        glBlendEquation(equation);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glBlendEquation\n");
        return false;
    }
}

bool handleGlBlendEquationSeparate(Stack* stack, bool pushReturn) {
    BlendEquation alpha = stack->pop<BlendEquation>();
    BlendEquation rgb = stack->pop<BlendEquation>();
    if (stack->isValid()) {
        CAZE_DEBUG("glBlendEquationSeparate(%u, %u)\n", rgb, alpha);
        glBlendEquationSeparate(rgb, alpha);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glBlendEquationSeparate\n");
        return false;
    }
}

bool handleGlBlendColor(Stack* stack, bool pushReturn) {
    float alpha = stack->pop<float>();
    float blue = stack->pop<float>();
    float green = stack->pop<float>();
    float red = stack->pop<float>();
    if (stack->isValid()) {
        CAZE_DEBUG("glBlendColor(%f, %f, %f, %f)\n", red, green, blue, alpha);
        glBlendColor(red, green, blue, alpha);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glBlendColor\n");
        return false;
    }
}

bool handleGlEnableVertexAttribArray(Stack* stack, bool pushReturn) {
    s32 index = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glEnableVertexAttribArray(%d)\n", index);
        glEnableVertexAttribArray(index);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glEnableVertexAttribArray\n");
        return false;
    }
}

bool handleGlDisableVertexAttribArray(Stack* stack, bool pushReturn) {
    s32 index = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glDisableVertexAttribArray(%d)\n", index);
        glDisableVertexAttribArray(index);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glDisableVertexAttribArray\n");
        return false;
    }
}

bool handleGlVertexAttribPointer(Stack* stack, bool pushReturn) {
    void* data = stack->pop<void*>();
    s32 stride = stack->pop<s32>();
    bool normalized = stack->pop<bool>();
    VertexAttribType type = stack->pop<VertexAttribType>();
    VertexAttribSize size = stack->pop<VertexAttribSize>();
    s32 index = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glVertexAttribPointer(%d, %u, %u, %d, %d, %p)\n", index, size, type, normalized,
                   stride, data);
        glVertexAttribPointer(index, size, type, normalized, stride, data);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glVertexAttribPointer\n");
        return false;
    }
}

bool handleGlGetActiveAttrib(Stack* stack, bool pushReturn) {
    char* name = stack->pop<char*>();
    ShaderAttribType* type = stack->pop<ShaderAttribType*>();
    s32* vector_count = stack->pop<s32*>();
    s32* buffer_bytes_written = stack->pop<s32*>();
    s32 buffer_size = stack->pop<s32>();
    s32 location = stack->pop<s32>();
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glGetActiveAttrib(%u, %d, %d, %p, %p, %p, %p)\n", program, location,
                   buffer_size, buffer_bytes_written, vector_count, type, name);
        glGetActiveAttrib(program, location, buffer_size, buffer_bytes_written, vector_count, type,
                          name);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glGetActiveAttrib\n");
        return false;
    }
}

bool handleGlGetActiveUniform(Stack* stack, bool pushReturn) {
    char* name = stack->pop<char*>();
    ShaderUniformType* type = stack->pop<ShaderUniformType*>();
    s32* size = stack->pop<s32*>();
    s32* buffer_bytes_written = stack->pop<s32*>();
    s32 buffer_size = stack->pop<s32>();
    s32 location = stack->pop<s32>();
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glGetActiveUniform(%u, %d, %d, %p, %p, %p, %p)\n", program, location,
                   buffer_size, buffer_bytes_written, size, type, name);
        glGetActiveUniform(program, location, buffer_size, buffer_bytes_written, size, type, name);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glGetActiveUniform\n");
        return false;
    }
}

bool handleGlGetError(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        CAZE_DEBUG("glGetError()\n");
        const Error return_value = glGetError();
        if (pushReturn) {
            stack->push<const Error>(return_value);
        }
        return true;
    } else {
        CAZE_WARNING("Error during calling function glGetError\n");
        return false;
    }
}

bool handleGlGetProgramiv(Stack* stack, bool pushReturn) {
    s32* value = stack->pop<s32*>();
    ProgramParameter parameter = stack->pop<ProgramParameter>();
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glGetProgramiv(%u, %u, %p)\n", program, parameter, value);
        glGetProgramiv(program, parameter, value);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glGetProgramiv\n");
        return false;
    }
}

bool handleGlGetShaderiv(Stack* stack, bool pushReturn) {
    s32* value = stack->pop<s32*>();
    ShaderParameter parameter = stack->pop<ShaderParameter>();
    ShaderId shader = stack->pop<ShaderId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glGetShaderiv(%u, %u, %p)\n", shader, parameter, value);
        glGetShaderiv(shader, parameter, value);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glGetShaderiv\n");
        return false;
    }
}

bool handleGlGetUniformLocation(Stack* stack, bool pushReturn) {
    char* name = stack->pop<char*>();
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glGetUniformLocation(%u, %s)\n", program, name);
        const s32 return_value = glGetUniformLocation(program, name);
        if (pushReturn) {
            stack->push<const s32>(return_value);
        }
        return true;
    } else {
        CAZE_WARNING("Error during calling function glGetUniformLocation\n");
        return false;
    }
}

bool handleGlGetAttribLocation(Stack* stack, bool pushReturn) {
    char* name = stack->pop<char*>();
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glGetAttribLocation(%u, %s)\n", program, name);
        const s32 return_value = glGetAttribLocation(program, name);
        if (pushReturn) {
            stack->push<const s32>(return_value);
        }
        return true;
    } else {
        CAZE_WARNING("Error during calling function glGetAttribLocation\n");
        return false;
    }
}

bool handleGlPixelStorei(Stack* stack, bool pushReturn) {
    s32 value = stack->pop<s32>();
    PixelStoreParameter parameter = stack->pop<PixelStoreParameter>();
    if (stack->isValid()) {
        CAZE_DEBUG("glPixelStorei(%u, %d)\n", parameter, value);
        glPixelStorei(parameter, value);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glPixelStorei\n");
        return false;
    }
}

bool handleGlTexParameteri(Stack* stack, bool pushReturn) {
    s32 value = stack->pop<s32>();
    TextureParameter parameter = stack->pop<TextureParameter>();
    TextureTarget target = stack->pop<TextureTarget>();
    if (stack->isValid()) {
        CAZE_DEBUG("glTexParameteri(%u, %u, %d)\n", target, parameter, value);
        glTexParameteri(target, parameter, value);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glTexParameteri\n");
        return false;
    }
}

bool handleGlTexParameterf(Stack* stack, bool pushReturn) {
    float value = stack->pop<float>();
    TextureParameter parameter = stack->pop<TextureParameter>();
    TextureTarget target = stack->pop<TextureTarget>();
    if (stack->isValid()) {
        CAZE_DEBUG("glTexParameterf(%u, %u, %f)\n", target, parameter, value);
        glTexParameterf(target, parameter, value);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glTexParameterf\n");
        return false;
    }
}

bool handleGlGetTexParameteriv(Stack* stack, bool pushReturn) {
    s32* values = stack->pop<s32*>();
    TextureParameter parameter = stack->pop<TextureParameter>();
    TextureTarget target = stack->pop<TextureTarget>();
    if (stack->isValid()) {
        CAZE_DEBUG("glGetTexParameteriv(%u, %u, %p)\n", target, parameter, values);
        glGetTexParameteriv(target, parameter, values);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glGetTexParameteriv\n");
        return false;
    }
}

bool handleGlGetTexParameterfv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    TextureParameter parameter = stack->pop<TextureParameter>();
    TextureTarget target = stack->pop<TextureTarget>();
    if (stack->isValid()) {
        CAZE_DEBUG("glGetTexParameterfv(%u, %u, %p)\n", target, parameter, values);
        glGetTexParameterfv(target, parameter, values);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glGetTexParameterfv\n");
        return false;
    }
}

bool handleGlUniform1i(Stack* stack, bool pushReturn) {
    s32 value = stack->pop<s32>();
    s32 location = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glUniform1i(%d, %d)\n", location, value);
        glUniform1i(location, value);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glUniform1i\n");
        return false;
    }
}

bool handleGlUniform2i(Stack* stack, bool pushReturn) {
    s32 value1 = stack->pop<s32>();
    s32 value0 = stack->pop<s32>();
    s32 location = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glUniform2i(%d, %d, %d)\n", location, value0, value1);
        glUniform2i(location, value0, value1);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glUniform2i\n");
        return false;
    }
}

bool handleGlUniform3i(Stack* stack, bool pushReturn) {
    s32 value2 = stack->pop<s32>();
    s32 value1 = stack->pop<s32>();
    s32 value0 = stack->pop<s32>();
    s32 location = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glUniform3i(%d, %d, %d, %d)\n", location, value0, value1, value2);
        glUniform3i(location, value0, value1, value2);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glUniform3i\n");
        return false;
    }
}

bool handleGlUniform4i(Stack* stack, bool pushReturn) {
    s32 value3 = stack->pop<s32>();
    s32 value2 = stack->pop<s32>();
    s32 value1 = stack->pop<s32>();
    s32 value0 = stack->pop<s32>();
    s32 location = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glUniform4i(%d, %d, %d, %d, %d)\n", location, value0, value1, value2, value3);
        glUniform4i(location, value0, value1, value2, value3);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glUniform4i\n");
        return false;
    }
}

bool handleGlUniform1iv(Stack* stack, bool pushReturn) {
    s32* value = stack->pop<s32*>();
    s32 count = stack->pop<s32>();
    s32 location = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glUniform1iv(%d, %d, %p)\n", location, count, value);
        glUniform1iv(location, count, value);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glUniform1iv\n");
        return false;
    }
}

bool handleGlUniform2iv(Stack* stack, bool pushReturn) {
    s32* value = stack->pop<s32*>();
    s32 count = stack->pop<s32>();
    s32 location = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glUniform2iv(%d, %d, %p)\n", location, count, value);
        glUniform2iv(location, count, value);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glUniform2iv\n");
        return false;
    }
}

bool handleGlUniform3iv(Stack* stack, bool pushReturn) {
    s32* value = stack->pop<s32*>();
    s32 count = stack->pop<s32>();
    s32 location = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glUniform3iv(%d, %d, %p)\n", location, count, value);
        glUniform3iv(location, count, value);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glUniform3iv\n");
        return false;
    }
}

bool handleGlUniform4iv(Stack* stack, bool pushReturn) {
    s32* value = stack->pop<s32*>();
    s32 count = stack->pop<s32>();
    s32 location = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glUniform4iv(%d, %d, %p)\n", location, count, value);
        glUniform4iv(location, count, value);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glUniform4iv\n");
        return false;
    }
}

bool handleGlUniform1f(Stack* stack, bool pushReturn) {
    float value = stack->pop<float>();
    s32 location = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glUniform1f(%d, %f)\n", location, value);
        glUniform1f(location, value);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glUniform1f\n");
        return false;
    }
}

bool handleGlUniform2f(Stack* stack, bool pushReturn) {
    float value1 = stack->pop<float>();
    float value0 = stack->pop<float>();
    s32 location = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glUniform2f(%d, %f, %f)\n", location, value0, value1);
        glUniform2f(location, value0, value1);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glUniform2f\n");
        return false;
    }
}

bool handleGlUniform3f(Stack* stack, bool pushReturn) {
    float value2 = stack->pop<float>();
    float value1 = stack->pop<float>();
    float value0 = stack->pop<float>();
    s32 location = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glUniform3f(%d, %f, %f, %f)\n", location, value0, value1, value2);
        glUniform3f(location, value0, value1, value2);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glUniform3f\n");
        return false;
    }
}

bool handleGlUniform4f(Stack* stack, bool pushReturn) {
    float value3 = stack->pop<float>();
    float value2 = stack->pop<float>();
    float value1 = stack->pop<float>();
    float value0 = stack->pop<float>();
    s32 location = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glUniform4f(%d, %f, %f, %f, %f)\n", location, value0, value1, value2, value3);
        glUniform4f(location, value0, value1, value2, value3);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glUniform4f\n");
        return false;
    }
}

bool handleGlUniform1fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    s32 count = stack->pop<s32>();
    s32 location = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glUniform1fv(%d, %d, %p)\n", location, count, value);
        glUniform1fv(location, count, value);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glUniform1fv\n");
        return false;
    }
}

bool handleGlUniform2fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    s32 count = stack->pop<s32>();
    s32 location = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glUniform2fv(%d, %d, %p)\n", location, count, value);
        glUniform2fv(location, count, value);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glUniform2fv\n");
        return false;
    }
}

bool handleGlUniform3fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    s32 count = stack->pop<s32>();
    s32 location = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glUniform3fv(%d, %d, %p)\n", location, count, value);
        glUniform3fv(location, count, value);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glUniform3fv\n");
        return false;
    }
}

bool handleGlUniform4fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    s32 count = stack->pop<s32>();
    s32 location = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glUniform4fv(%d, %d, %p)\n", location, count, value);
        glUniform4fv(location, count, value);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glUniform4fv\n");
        return false;
    }
}

bool handleGlUniformMatrix2fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    bool transpose = stack->pop<bool>();
    s32 count = stack->pop<s32>();
    s32 location = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glUniformMatrix2fv(%d, %d, %d, %p)\n", location, count, transpose, values);
        glUniformMatrix2fv(location, count, transpose, values);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glUniformMatrix2fv\n");
        return false;
    }
}

bool handleGlUniformMatrix3fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    bool transpose = stack->pop<bool>();
    s32 count = stack->pop<s32>();
    s32 location = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glUniformMatrix3fv(%d, %d, %d, %p)\n", location, count, transpose, values);
        glUniformMatrix3fv(location, count, transpose, values);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glUniformMatrix3fv\n");
        return false;
    }
}

bool handleGlUniformMatrix4fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    bool transpose = stack->pop<bool>();
    s32 count = stack->pop<s32>();
    s32 location = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glUniformMatrix4fv(%d, %d, %d, %p)\n", location, count, transpose, values);
        glUniformMatrix4fv(location, count, transpose, values);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glUniformMatrix4fv\n");
        return false;
    }
}

bool handleGlGetUniformfv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    s32 location = stack->pop<s32>();
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glGetUniformfv(%u, %d, %p)\n", program, location, values);
        glGetUniformfv(program, location, values);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glGetUniformfv\n");
        return false;
    }
}

bool handleGlGetUniformiv(Stack* stack, bool pushReturn) {
    s32* values = stack->pop<s32*>();
    s32 location = stack->pop<s32>();
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glGetUniformiv(%u, %d, %p)\n", program, location, values);
        glGetUniformiv(program, location, values);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glGetUniformiv\n");
        return false;
    }
}

bool handleGlVertexAttrib1f(Stack* stack, bool pushReturn) {
    float value0 = stack->pop<float>();
    s32 index = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glVertexAttrib1f(%d, %f)\n", index, value0);
        glVertexAttrib1f(index, value0);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glVertexAttrib1f\n");
        return false;
    }
}

bool handleGlVertexAttrib2f(Stack* stack, bool pushReturn) {
    float value1 = stack->pop<float>();
    float value0 = stack->pop<float>();
    s32 index = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glVertexAttrib2f(%d, %f, %f)\n", index, value0, value1);
        glVertexAttrib2f(index, value0, value1);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glVertexAttrib2f\n");
        return false;
    }
}

bool handleGlVertexAttrib3f(Stack* stack, bool pushReturn) {
    float value2 = stack->pop<float>();
    float value1 = stack->pop<float>();
    float value0 = stack->pop<float>();
    s32 index = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glVertexAttrib3f(%d, %f, %f, %f)\n", index, value0, value1, value2);
        glVertexAttrib3f(index, value0, value1, value2);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glVertexAttrib3f\n");
        return false;
    }
}

bool handleGlVertexAttrib4f(Stack* stack, bool pushReturn) {
    float value3 = stack->pop<float>();
    float value2 = stack->pop<float>();
    float value1 = stack->pop<float>();
    float value0 = stack->pop<float>();
    s32 index = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glVertexAttrib4f(%d, %f, %f, %f, %f)\n", index, value0, value1, value2, value3);
        glVertexAttrib4f(index, value0, value1, value2, value3);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glVertexAttrib4f\n");
        return false;
    }
}

bool handleGlVertexAttrib1fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    s32 index = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glVertexAttrib1fv(%d, %p)\n", index, value);
        glVertexAttrib1fv(index, value);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glVertexAttrib1fv\n");
        return false;
    }
}

bool handleGlVertexAttrib2fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    s32 index = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glVertexAttrib2fv(%d, %p)\n", index, value);
        glVertexAttrib2fv(index, value);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glVertexAttrib2fv\n");
        return false;
    }
}

bool handleGlVertexAttrib3fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    s32 index = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glVertexAttrib3fv(%d, %p)\n", index, value);
        glVertexAttrib3fv(index, value);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glVertexAttrib3fv\n");
        return false;
    }
}

bool handleGlVertexAttrib4fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    s32 index = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glVertexAttrib4fv(%d, %p)\n", index, value);
        glVertexAttrib4fv(index, value);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glVertexAttrib4fv\n");
        return false;
    }
}

bool handleGlGetShaderPrecisionFormat(Stack* stack, bool pushReturn) {
    PrecisionType precision_type = stack->pop<PrecisionType>();
    ShaderType shader_type = stack->pop<ShaderType>();
    if (stack->isValid()) {
        CAZE_DEBUG("glGetShaderPrecisionFormat(%u, %u)\n", shader_type, precision_type);
        glGetShaderPrecisionFormat(shader_type, precision_type);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glGetShaderPrecisionFormat\n");
        return false;
    }
}

bool handleGlDepthMask(Stack* stack, bool pushReturn) {
    bool enabled = stack->pop<bool>();
    if (stack->isValid()) {
        CAZE_DEBUG("glDepthMask(%d)\n", enabled);
        glDepthMask(enabled);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glDepthMask\n");
        return false;
    }
}

bool handleGlDepthFunc(Stack* stack, bool pushReturn) {
    TestFunction function = stack->pop<TestFunction>();
    if (stack->isValid()) {
        CAZE_DEBUG("glDepthFunc(%u)\n", function);
        glDepthFunc(function);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glDepthFunc\n");
        return false;
    }
}

bool handleGlDepthRangef(Stack* stack, bool pushReturn) {
    float far = stack->pop<float>();
    float near = stack->pop<float>();
    if (stack->isValid()) {
        CAZE_DEBUG("glDepthRangef(%f, %f)\n", near, far);
        glDepthRangef(near, far);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glDepthRangef\n");
        return false;
    }
}

bool handleGlColorMask(Stack* stack, bool pushReturn) {
    bool alpha = stack->pop<bool>();
    bool blue = stack->pop<bool>();
    bool green = stack->pop<bool>();
    bool red = stack->pop<bool>();
    if (stack->isValid()) {
        CAZE_DEBUG("glColorMask(%d, %d, %d, %d)\n", red, green, blue, alpha);
        glColorMask(red, green, blue, alpha);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glColorMask\n");
        return false;
    }
}

bool handleGlStencilMask(Stack* stack, bool pushReturn) {
    u32 mask = stack->pop<u32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glStencilMask(%u)\n", mask);
        glStencilMask(mask);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glStencilMask\n");
        return false;
    }
}

bool handleGlStencilMaskSeparate(Stack* stack, bool pushReturn) {
    u32 mask = stack->pop<u32>();
    FaceMode face = stack->pop<FaceMode>();
    if (stack->isValid()) {
        CAZE_DEBUG("glStencilMaskSeparate(%u, %u)\n", face, mask);
        glStencilMaskSeparate(face, mask);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glStencilMaskSeparate\n");
        return false;
    }
}

bool handleGlStencilFuncSeparate(Stack* stack, bool pushReturn) {
    s32 mask = stack->pop<s32>();
    s32 reference_value = stack->pop<s32>();
    TestFunction function = stack->pop<TestFunction>();
    FaceMode face = stack->pop<FaceMode>();
    if (stack->isValid()) {
        CAZE_DEBUG("glStencilFuncSeparate(%u, %u, %d, %d)\n", face, function, reference_value,
                   mask);
        glStencilFuncSeparate(face, function, reference_value, mask);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glStencilFuncSeparate\n");
        return false;
    }
}

bool handleGlStencilOpSeparate(Stack* stack, bool pushReturn) {
    StencilAction stencil_pass_depth_pass = stack->pop<StencilAction>();
    StencilAction stencil_pass_depth_fail = stack->pop<StencilAction>();
    StencilAction stencil_fail = stack->pop<StencilAction>();
    FaceMode face = stack->pop<FaceMode>();
    if (stack->isValid()) {
        CAZE_DEBUG("glStencilOpSeparate(%u, %u, %u, %u)\n", face, stencil_fail,
                   stencil_pass_depth_fail, stencil_pass_depth_pass);
        glStencilOpSeparate(face, stencil_fail, stencil_pass_depth_fail, stencil_pass_depth_pass);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glStencilOpSeparate\n");
        return false;
    }
}

bool handleGlFrontFace(Stack* stack, bool pushReturn) {
    FaceOrientation orientation = stack->pop<FaceOrientation>();
    if (stack->isValid()) {
        CAZE_DEBUG("glFrontFace(%u)\n", orientation);
        glFrontFace(orientation);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glFrontFace\n");
        return false;
    }
}

bool handleGlViewport(Stack* stack, bool pushReturn) {
    s32 height = stack->pop<s32>();
    s32 width = stack->pop<s32>();
    s32 y = stack->pop<s32>();
    s32 x = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glViewport(%d, %d, %d, %d)\n", x, y, width, height);
        glViewport(x, y, width, height);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glViewport\n");
        return false;
    }
}

bool handleGlScissor(Stack* stack, bool pushReturn) {
    s32 height = stack->pop<s32>();
    s32 width = stack->pop<s32>();
    s32 y = stack->pop<s32>();
    s32 x = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glScissor(%d, %d, %d, %d)\n", x, y, width, height);
        glScissor(x, y, width, height);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glScissor\n");
        return false;
    }
}

bool handleGlActiveTexture(Stack* stack, bool pushReturn) {
    TextureUnit unit = stack->pop<TextureUnit>();
    if (stack->isValid()) {
        CAZE_DEBUG("glActiveTexture(%u)\n", unit);
        glActiveTexture(unit);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glActiveTexture\n");
        return false;
    }
}

bool handleGlGenTextures(Stack* stack, bool pushReturn) {
    TextureId* textures = stack->pop<TextureId*>();
    s32 count = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glGenTextures(%d, %p)\n", count, textures);
        glGenTextures(count, textures);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glGenTextures\n");
        return false;
    }
}

bool handleGlDeleteTextures(Stack* stack, bool pushReturn) {
    TextureId* textures = stack->pop<TextureId*>();
    s32 count = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glDeleteTextures(%d, %p)\n", count, textures);
        glDeleteTextures(count, textures);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glDeleteTextures\n");
        return false;
    }
}

bool handleGlIsTexture(Stack* stack, bool pushReturn) {
    TextureId texture = stack->pop<TextureId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glIsTexture(%u)\n", texture);
        const bool return_value = glIsTexture(texture);
        if (pushReturn) {
            stack->push<const bool>(return_value);
        }
        return true;
    } else {
        CAZE_WARNING("Error during calling function glIsTexture\n");
        return false;
    }
}

bool handleGlBindTexture(Stack* stack, bool pushReturn) {
    TextureId texture = stack->pop<TextureId>();
    TextureTarget target = stack->pop<TextureTarget>();
    if (stack->isValid()) {
        CAZE_DEBUG("glBindTexture(%u, %u)\n", target, texture);
        glBindTexture(target, texture);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glBindTexture\n");
        return false;
    }
}

bool handleGlTexImage2D(Stack* stack, bool pushReturn) {
    void* data = stack->pop<void*>();
    TexelType type = stack->pop<TexelType>();
    TexelFormat format = stack->pop<TexelFormat>();
    s32 border = stack->pop<s32>();
    s32 height = stack->pop<s32>();
    s32 width = stack->pop<s32>();
    TexelFormat internal_format = stack->pop<TexelFormat>();
    s32 level = stack->pop<s32>();
    TextureImageTarget target = stack->pop<TextureImageTarget>();
    if (stack->isValid()) {
        CAZE_DEBUG("glTexImage2D(%u, %d, %u, %d, %d, %d, %u, %u, %p)\n", target, level,
                   internal_format, width, height, border, format, type, data);
        glTexImage2D(target, level, internal_format, width, height, border, format, type, data);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glTexImage2D\n");
        return false;
    }
}

bool handleGlTexSubImage2D(Stack* stack, bool pushReturn) {
    void* data = stack->pop<void*>();
    TexelType type = stack->pop<TexelType>();
    TexelFormat format = stack->pop<TexelFormat>();
    s32 height = stack->pop<s32>();
    s32 width = stack->pop<s32>();
    s32 yoffset = stack->pop<s32>();
    s32 xoffset = stack->pop<s32>();
    s32 level = stack->pop<s32>();
    TextureImageTarget target = stack->pop<TextureImageTarget>();
    if (stack->isValid()) {
        CAZE_DEBUG("glTexSubImage2D(%u, %d, %d, %d, %d, %d, %u, %u, %p)\n", target, level, xoffset,
                   yoffset, width, height, format, type, data);
        glTexSubImage2D(target, level, xoffset, yoffset, width, height, format, type, data);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glTexSubImage2D\n");
        return false;
    }
}

bool handleGlCopyTexImage2D(Stack* stack, bool pushReturn) {
    s32 border = stack->pop<s32>();
    s32 height = stack->pop<s32>();
    s32 width = stack->pop<s32>();
    s32 y = stack->pop<s32>();
    s32 x = stack->pop<s32>();
    TexelFormat format = stack->pop<TexelFormat>();
    s32 level = stack->pop<s32>();
    TextureImageTarget target = stack->pop<TextureImageTarget>();
    if (stack->isValid()) {
        CAZE_DEBUG("glCopyTexImage2D(%u, %d, %u, %d, %d, %d, %d, %d)\n", target, level, format, x,
                   y, width, height, border);
        glCopyTexImage2D(target, level, format, x, y, width, height, border);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glCopyTexImage2D\n");
        return false;
    }
}

bool handleGlCopyTexSubImage2D(Stack* stack, bool pushReturn) {
    s32 height = stack->pop<s32>();
    s32 width = stack->pop<s32>();
    s32 y = stack->pop<s32>();
    s32 x = stack->pop<s32>();
    s32 yoffset = stack->pop<s32>();
    s32 xoffset = stack->pop<s32>();
    s32 level = stack->pop<s32>();
    TextureImageTarget target = stack->pop<TextureImageTarget>();
    if (stack->isValid()) {
        CAZE_DEBUG("glCopyTexSubImage2D(%u, %d, %d, %d, %d, %d, %d, %d)\n", target, level, xoffset,
                   yoffset, x, y, width, height);
        glCopyTexSubImage2D(target, level, xoffset, yoffset, x, y, width, height);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glCopyTexSubImage2D\n");
        return false;
    }
}

bool handleGlCompressedTexImage2D(Stack* stack, bool pushReturn) {
    void* data = stack->pop<void*>();
    s32 image_size = stack->pop<s32>();
    s32 border = stack->pop<s32>();
    s32 height = stack->pop<s32>();
    s32 width = stack->pop<s32>();
    CompressedTexelFormat format = stack->pop<CompressedTexelFormat>();
    s32 level = stack->pop<s32>();
    TextureImageTarget target = stack->pop<TextureImageTarget>();
    if (stack->isValid()) {
        CAZE_DEBUG("glCompressedTexImage2D(%u, %d, %u, %d, %d, %d, %d, %p)\n", target, level,
                   format, width, height, border, image_size, data);
        glCompressedTexImage2D(target, level, format, width, height, border, image_size, data);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glCompressedTexImage2D\n");
        return false;
    }
}

bool handleGlCompressedTexSubImage2D(Stack* stack, bool pushReturn) {
    void* data = stack->pop<void*>();
    s32 image_size = stack->pop<s32>();
    CompressedTexelFormat format = stack->pop<CompressedTexelFormat>();
    s32 height = stack->pop<s32>();
    s32 width = stack->pop<s32>();
    s32 yoffset = stack->pop<s32>();
    s32 xoffset = stack->pop<s32>();
    s32 level = stack->pop<s32>();
    TextureImageTarget target = stack->pop<TextureImageTarget>();
    if (stack->isValid()) {
        CAZE_DEBUG("glCompressedTexSubImage2D(%u, %d, %d, %d, %d, %d, %u, %d, %p)\n", target, level,
                   xoffset, yoffset, width, height, format, image_size, data);
        glCompressedTexSubImage2D(target, level, xoffset, yoffset, width, height, format,
                                  image_size, data);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glCompressedTexSubImage2D\n");
        return false;
    }
}

bool handleGlGenerateMipmap(Stack* stack, bool pushReturn) {
    TextureImageTarget target = stack->pop<TextureImageTarget>();
    if (stack->isValid()) {
        CAZE_DEBUG("glGenerateMipmap(%u)\n", target);
        glGenerateMipmap(target);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glGenerateMipmap\n");
        return false;
    }
}

bool handleGlReadPixels(Stack* stack, bool pushReturn) {
    void* data = stack->pop<void*>();
    TexelType type = stack->pop<TexelType>();
    ReadPixelsFormat format = stack->pop<ReadPixelsFormat>();
    s32 height = stack->pop<s32>();
    s32 width = stack->pop<s32>();
    s32 y = stack->pop<s32>();
    s32 x = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glReadPixels(%d, %d, %d, %d, %u, %u, %p)\n", x, y, width, height, format, type,
                   data);
        glReadPixels(x, y, width, height, format, type, data);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glReadPixels\n");
        return false;
    }
}

bool handleGlGenFramebuffers(Stack* stack, bool pushReturn) {
    FramebufferId* framebuffers = stack->pop<FramebufferId*>();
    s32 count = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glGenFramebuffers(%d, %p)\n", count, framebuffers);
        glGenFramebuffers(count, framebuffers);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glGenFramebuffers\n");
        return false;
    }
}

bool handleGlBindFramebuffer(Stack* stack, bool pushReturn) {
    FramebufferId framebuffer = stack->pop<FramebufferId>();
    FramebufferTarget target = stack->pop<FramebufferTarget>();
    if (stack->isValid()) {
        CAZE_DEBUG("glBindFramebuffer(%u, %u)\n", target, framebuffer);
        glBindFramebuffer(target, framebuffer);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glBindFramebuffer\n");
        return false;
    }
}

bool handleGlCheckFramebufferStatus(Stack* stack, bool pushReturn) {
    FramebufferTarget target = stack->pop<FramebufferTarget>();
    if (stack->isValid()) {
        CAZE_DEBUG("glCheckFramebufferStatus(%u)\n", target);
        const FramebufferStatus return_value = glCheckFramebufferStatus(target);
        if (pushReturn) {
            stack->push<const FramebufferStatus>(return_value);
        }
        return true;
    } else {
        CAZE_WARNING("Error during calling function glCheckFramebufferStatus\n");
        return false;
    }
}

bool handleGlDeleteFramebuffers(Stack* stack, bool pushReturn) {
    FramebufferId* framebuffers = stack->pop<FramebufferId*>();
    s32 count = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glDeleteFramebuffers(%d, %p)\n", count, framebuffers);
        glDeleteFramebuffers(count, framebuffers);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glDeleteFramebuffers\n");
        return false;
    }
}

bool handleGlIsFramebuffer(Stack* stack, bool pushReturn) {
    FramebufferId framebuffer = stack->pop<FramebufferId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glIsFramebuffer(%u)\n", framebuffer);
        const bool return_value = glIsFramebuffer(framebuffer);
        if (pushReturn) {
            stack->push<const bool>(return_value);
        }
        return true;
    } else {
        CAZE_WARNING("Error during calling function glIsFramebuffer\n");
        return false;
    }
}

bool handleGlGenRenderbuffers(Stack* stack, bool pushReturn) {
    RenderbufferId* renderbuffers = stack->pop<RenderbufferId*>();
    s32 count = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glGenRenderbuffers(%d, %p)\n", count, renderbuffers);
        glGenRenderbuffers(count, renderbuffers);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glGenRenderbuffers\n");
        return false;
    }
}

bool handleGlBindRenderbuffer(Stack* stack, bool pushReturn) {
    RenderbufferId renderbuffer = stack->pop<RenderbufferId>();
    RenderbufferTarget target = stack->pop<RenderbufferTarget>();
    if (stack->isValid()) {
        CAZE_DEBUG("glBindRenderbuffer(%u, %u)\n", target, renderbuffer);
        glBindRenderbuffer(target, renderbuffer);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glBindRenderbuffer\n");
        return false;
    }
}

bool handleGlRenderbufferStorage(Stack* stack, bool pushReturn) {
    s32 height = stack->pop<s32>();
    s32 width = stack->pop<s32>();
    RenderbufferFormat format = stack->pop<RenderbufferFormat>();
    RenderbufferTarget target = stack->pop<RenderbufferTarget>();
    if (stack->isValid()) {
        CAZE_DEBUG("glRenderbufferStorage(%u, %u, %d, %d)\n", target, format, width, height);
        glRenderbufferStorage(target, format, width, height);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glRenderbufferStorage\n");
        return false;
    }
}

bool handleGlDeleteRenderbuffers(Stack* stack, bool pushReturn) {
    RenderbufferId* renderbuffers = stack->pop<RenderbufferId*>();
    s32 count = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glDeleteRenderbuffers(%d, %p)\n", count, renderbuffers);
        glDeleteRenderbuffers(count, renderbuffers);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glDeleteRenderbuffers\n");
        return false;
    }
}

bool handleGlIsRenderbuffer(Stack* stack, bool pushReturn) {
    RenderbufferId renderbuffer = stack->pop<RenderbufferId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glIsRenderbuffer(%u)\n", renderbuffer);
        const bool return_value = glIsRenderbuffer(renderbuffer);
        if (pushReturn) {
            stack->push<const bool>(return_value);
        }
        return true;
    } else {
        CAZE_WARNING("Error during calling function glIsRenderbuffer\n");
        return false;
    }
}

bool handleGlGetRenderbufferParameteriv(Stack* stack, bool pushReturn) {
    s32* values = stack->pop<s32*>();
    RenderbufferParameter parameter = stack->pop<RenderbufferParameter>();
    RenderbufferTarget target = stack->pop<RenderbufferTarget>();
    if (stack->isValid()) {
        CAZE_DEBUG("glGetRenderbufferParameteriv(%u, %u, %p)\n", target, parameter, values);
        glGetRenderbufferParameteriv(target, parameter, values);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glGetRenderbufferParameteriv\n");
        return false;
    }
}

bool handleGlGenBuffers(Stack* stack, bool pushReturn) {
    BufferId* buffers = stack->pop<BufferId*>();
    s32 count = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glGenBuffers(%d, %p)\n", count, buffers);
        glGenBuffers(count, buffers);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glGenBuffers\n");
        return false;
    }
}

bool handleGlBindBuffer(Stack* stack, bool pushReturn) {
    BufferId buffer = stack->pop<BufferId>();
    BufferTarget target = stack->pop<BufferTarget>();
    if (stack->isValid()) {
        CAZE_DEBUG("glBindBuffer(%u, %u)\n", target, buffer);
        glBindBuffer(target, buffer);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glBindBuffer\n");
        return false;
    }
}

bool handleGlBufferData(Stack* stack, bool pushReturn) {
    BufferUsage usage = stack->pop<BufferUsage>();
    void* data = stack->pop<void*>();
    s32 size = stack->pop<s32>();
    BufferTarget target = stack->pop<BufferTarget>();
    if (stack->isValid()) {
        CAZE_DEBUG("glBufferData(%u, %d, %p, %u)\n", target, size, data, usage);
        glBufferData(target, size, data, usage);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glBufferData\n");
        return false;
    }
}

bool handleGlBufferSubData(Stack* stack, bool pushReturn) {
    void* data = stack->pop<void*>();
    s32 size = stack->pop<s32>();
    s32 offset = stack->pop<s32>();
    BufferTarget target = stack->pop<BufferTarget>();
    if (stack->isValid()) {
        CAZE_DEBUG("glBufferSubData(%u, %d, %d, %p)\n", target, offset, size, data);
        glBufferSubData(target, offset, size, data);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glBufferSubData\n");
        return false;
    }
}

bool handleGlDeleteBuffers(Stack* stack, bool pushReturn) {
    BufferId* buffers = stack->pop<BufferId*>();
    s32 count = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glDeleteBuffers(%d, %p)\n", count, buffers);
        glDeleteBuffers(count, buffers);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glDeleteBuffers\n");
        return false;
    }
}

bool handleGlIsBuffer(Stack* stack, bool pushReturn) {
    BufferId buffer = stack->pop<BufferId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glIsBuffer(%u)\n", buffer);
        const bool return_value = glIsBuffer(buffer);
        if (pushReturn) {
            stack->push<const bool>(return_value);
        }
        return true;
    } else {
        CAZE_WARNING("Error during calling function glIsBuffer\n");
        return false;
    }
}

bool handleGlGetBufferParameteriv(Stack* stack, bool pushReturn) {
    s32* value = stack->pop<s32*>();
    BufferParameter parameter = stack->pop<BufferParameter>();
    BufferTarget target = stack->pop<BufferTarget>();
    if (stack->isValid()) {
        CAZE_DEBUG("glGetBufferParameteriv(%u, %u, %p)\n", target, parameter, value);
        glGetBufferParameteriv(target, parameter, value);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glGetBufferParameteriv\n");
        return false;
    }
}

bool handleGlCreateShader(Stack* stack, bool pushReturn) {
    ShaderType type = stack->pop<ShaderType>();
    if (stack->isValid()) {
        CAZE_DEBUG("glCreateShader(%u)\n", type);
        const ShaderId return_value = glCreateShader(type);
        if (pushReturn) {
            stack->push<const ShaderId>(return_value);
        }
        return true;
    } else {
        CAZE_WARNING("Error during calling function glCreateShader\n");
        return false;
    }
}

bool handleGlDeleteShader(Stack* stack, bool pushReturn) {
    ShaderId shader = stack->pop<ShaderId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glDeleteShader(%u)\n", shader);
        glDeleteShader(shader);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glDeleteShader\n");
        return false;
    }
}

bool handleGlShaderSource(Stack* stack, bool pushReturn) {
    s32* length = stack->pop<s32*>();
    const char** source = stack->pop<const char**>();
    s32 count = stack->pop<s32>();
    ShaderId shader = stack->pop<ShaderId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glShaderSource(%u, %d, %p, %p)\n", shader, count, source, length);
        glShaderSource(shader, count, source, length);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glShaderSource\n");
        return false;
    }
}

bool handleGlShaderBinary(Stack* stack, bool pushReturn) {
    s32 binary_size = stack->pop<s32>();
    void* binary = stack->pop<void*>();
    u32 binary_format = stack->pop<u32>();
    ShaderId* shaders = stack->pop<ShaderId*>();
    s32 count = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glShaderBinary(%d, %p, %u, %p, %d)\n", count, shaders, binary_format, binary,
                   binary_size);
        glShaderBinary(count, shaders, binary_format, binary, binary_size);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glShaderBinary\n");
        return false;
    }
}

bool handleGlGetShaderInfoLog(Stack* stack, bool pushReturn) {
    char* info = stack->pop<char*>();
    s32* string_length_written = stack->pop<s32*>();
    s32 buffer_length = stack->pop<s32>();
    ShaderId shader = stack->pop<ShaderId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glGetShaderInfoLog(%u, %d, %p, %p)\n", shader, buffer_length,
                   string_length_written, info);
        glGetShaderInfoLog(shader, buffer_length, string_length_written, info);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glGetShaderInfoLog\n");
        return false;
    }
}

bool handleGlGetShaderSource(Stack* stack, bool pushReturn) {
    char* source = stack->pop<char*>();
    s32* string_length_written = stack->pop<s32*>();
    s32 buffer_length = stack->pop<s32>();
    ShaderId shader = stack->pop<ShaderId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glGetShaderSource(%u, %d, %p, %p)\n", shader, buffer_length,
                   string_length_written, source);
        glGetShaderSource(shader, buffer_length, string_length_written, source);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glGetShaderSource\n");
        return false;
    }
}

bool handleGlReleaseShaderCompiler(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        CAZE_DEBUG("glReleaseShaderCompiler()\n");
        glReleaseShaderCompiler();
        return true;
    } else {
        CAZE_WARNING("Error during calling function glReleaseShaderCompiler\n");
        return false;
    }
}

bool handleGlCompileShader(Stack* stack, bool pushReturn) {
    ShaderId shader = stack->pop<ShaderId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glCompileShader(%u)\n", shader);
        glCompileShader(shader);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glCompileShader\n");
        return false;
    }
}

bool handleGlIsShader(Stack* stack, bool pushReturn) {
    ShaderId shader = stack->pop<ShaderId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glIsShader(%u)\n", shader);
        const bool return_value = glIsShader(shader);
        if (pushReturn) {
            stack->push<const bool>(return_value);
        }
        return true;
    } else {
        CAZE_WARNING("Error during calling function glIsShader\n");
        return false;
    }
}

bool handleGlCreateProgram(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        CAZE_DEBUG("glCreateProgram()\n");
        const ProgramId return_value = glCreateProgram();
        if (pushReturn) {
            stack->push<const ProgramId>(return_value);
        }
        return true;
    } else {
        CAZE_WARNING("Error during calling function glCreateProgram\n");
        return false;
    }
}

bool handleGlDeleteProgram(Stack* stack, bool pushReturn) {
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glDeleteProgram(%u)\n", program);
        glDeleteProgram(program);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glDeleteProgram\n");
        return false;
    }
}

bool handleGlAttachShader(Stack* stack, bool pushReturn) {
    ShaderId shader = stack->pop<ShaderId>();
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glAttachShader(%u, %u)\n", program, shader);
        glAttachShader(program, shader);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glAttachShader\n");
        return false;
    }
}

bool handleGlDetachShader(Stack* stack, bool pushReturn) {
    ShaderId shader = stack->pop<ShaderId>();
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glDetachShader(%u, %u)\n", program, shader);
        glDetachShader(program, shader);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glDetachShader\n");
        return false;
    }
}

bool handleGlGetAttachedShaders(Stack* stack, bool pushReturn) {
    ShaderId* shaders = stack->pop<ShaderId*>();
    s32* shaders_length_written = stack->pop<s32*>();
    s32 buffer_length = stack->pop<s32>();
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glGetAttachedShaders(%u, %d, %p, %p)\n", program, buffer_length,
                   shaders_length_written, shaders);
        glGetAttachedShaders(program, buffer_length, shaders_length_written, shaders);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glGetAttachedShaders\n");
        return false;
    }
}

bool handleGlLinkProgram(Stack* stack, bool pushReturn) {
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glLinkProgram(%u)\n", program);
        glLinkProgram(program);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glLinkProgram\n");
        return false;
    }
}

bool handleGlGetProgramInfoLog(Stack* stack, bool pushReturn) {
    char* info = stack->pop<char*>();
    s32* string_length_written = stack->pop<s32*>();
    s32 buffer_length = stack->pop<s32>();
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glGetProgramInfoLog(%u, %d, %p, %p)\n", program, buffer_length,
                   string_length_written, info);
        glGetProgramInfoLog(program, buffer_length, string_length_written, info);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glGetProgramInfoLog\n");
        return false;
    }
}

bool handleGlUseProgram(Stack* stack, bool pushReturn) {
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glUseProgram(%u)\n", program);
        glUseProgram(program);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glUseProgram\n");
        return false;
    }
}

bool handleGlIsProgram(Stack* stack, bool pushReturn) {
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glIsProgram(%u)\n", program);
        const bool return_value = glIsProgram(program);
        if (pushReturn) {
            stack->push<const bool>(return_value);
        }
        return true;
    } else {
        CAZE_WARNING("Error during calling function glIsProgram\n");
        return false;
    }
}

bool handleGlValidateProgram(Stack* stack, bool pushReturn) {
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        CAZE_DEBUG("glValidateProgram(%u)\n", program);
        glValidateProgram(program);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glValidateProgram\n");
        return false;
    }
}

bool handleGlClearColor(Stack* stack, bool pushReturn) {
    float a = stack->pop<float>();
    float b = stack->pop<float>();
    float g = stack->pop<float>();
    float r = stack->pop<float>();
    if (stack->isValid()) {
        CAZE_DEBUG("glClearColor(%f, %f, %f, %f)\n", r, g, b, a);
        glClearColor(r, g, b, a);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glClearColor\n");
        return false;
    }
}

bool handleGlClearDepthf(Stack* stack, bool pushReturn) {
    float depth = stack->pop<float>();
    if (stack->isValid()) {
        CAZE_DEBUG("glClearDepthf(%f)\n", depth);
        glClearDepthf(depth);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glClearDepthf\n");
        return false;
    }
}

bool handleGlClearStencil(Stack* stack, bool pushReturn) {
    s32 stencil = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glClearStencil(%d)\n", stencil);
        glClearStencil(stencil);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glClearStencil\n");
        return false;
    }
}

bool handleGlClear(Stack* stack, bool pushReturn) {
    ClearMask mask = stack->pop<ClearMask>();
    if (stack->isValid()) {
        CAZE_DEBUG("glClear(%u)\n", mask);
        glClear(mask);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glClear\n");
        return false;
    }
}

bool handleGlCullFace(Stack* stack, bool pushReturn) {
    FaceMode mode = stack->pop<FaceMode>();
    if (stack->isValid()) {
        CAZE_DEBUG("glCullFace(%u)\n", mode);
        glCullFace(mode);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glCullFace\n");
        return false;
    }
}

bool handleGlPolygonOffset(Stack* stack, bool pushReturn) {
    float units = stack->pop<float>();
    float scale_factor = stack->pop<float>();
    if (stack->isValid()) {
        CAZE_DEBUG("glPolygonOffset(%f, %f)\n", scale_factor, units);
        glPolygonOffset(scale_factor, units);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glPolygonOffset\n");
        return false;
    }
}

bool handleGlLineWidth(Stack* stack, bool pushReturn) {
    float width = stack->pop<float>();
    if (stack->isValid()) {
        CAZE_DEBUG("glLineWidth(%f)\n", width);
        glLineWidth(width);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glLineWidth\n");
        return false;
    }
}

bool handleGlSampleCoverage(Stack* stack, bool pushReturn) {
    bool invert = stack->pop<bool>();
    float value = stack->pop<float>();
    if (stack->isValid()) {
        CAZE_DEBUG("glSampleCoverage(%f, %d)\n", value, invert);
        glSampleCoverage(value, invert);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glSampleCoverage\n");
        return false;
    }
}

bool handleGlHint(Stack* stack, bool pushReturn) {
    HintMode mode = stack->pop<HintMode>();
    HintTarget target = stack->pop<HintTarget>();
    if (stack->isValid()) {
        CAZE_DEBUG("glHint(%u, %u)\n", target, mode);
        glHint(target, mode);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glHint\n");
        return false;
    }
}

bool handleGlFramebufferRenderbuffer(Stack* stack, bool pushReturn) {
    RenderbufferId renderbuffer = stack->pop<RenderbufferId>();
    RenderbufferTarget renderbuffer_target = stack->pop<RenderbufferTarget>();
    FramebufferAttachment framebuffer_attachment = stack->pop<FramebufferAttachment>();
    FramebufferTarget framebuffer_target = stack->pop<FramebufferTarget>();
    if (stack->isValid()) {
        CAZE_DEBUG("glFramebufferRenderbuffer(%u, %u, %u, %u)\n", framebuffer_target,
                   framebuffer_attachment, renderbuffer_target, renderbuffer);
        glFramebufferRenderbuffer(framebuffer_target, framebuffer_attachment, renderbuffer_target,
                                  renderbuffer);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glFramebufferRenderbuffer\n");
        return false;
    }
}

bool handleGlFramebufferTexture2D(Stack* stack, bool pushReturn) {
    s32 level = stack->pop<s32>();
    TextureId texture = stack->pop<TextureId>();
    TextureImageTarget texture_target = stack->pop<TextureImageTarget>();
    FramebufferAttachment framebuffer_attachment = stack->pop<FramebufferAttachment>();
    FramebufferTarget framebuffer_target = stack->pop<FramebufferTarget>();
    if (stack->isValid()) {
        CAZE_DEBUG("glFramebufferTexture2D(%u, %u, %u, %u, %d)\n", framebuffer_target,
                   framebuffer_attachment, texture_target, texture, level);
        glFramebufferTexture2D(framebuffer_target, framebuffer_attachment, texture_target, texture,
                               level);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glFramebufferTexture2D\n");
        return false;
    }
}

bool handleGlGetFramebufferAttachmentParameteriv(Stack* stack, bool pushReturn) {
    s32* value = stack->pop<s32*>();
    FramebufferAttachmentParameter parameter = stack->pop<FramebufferAttachmentParameter>();
    FramebufferAttachment attachment = stack->pop<FramebufferAttachment>();
    FramebufferTarget target = stack->pop<FramebufferTarget>();
    if (stack->isValid()) {
        CAZE_DEBUG("glGetFramebufferAttachmentParameteriv(%u, %u, %u, %p)\n", target, attachment,
                   parameter, value);
        glGetFramebufferAttachmentParameteriv(target, attachment, parameter, value);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glGetFramebufferAttachmentParameteriv\n");
        return false;
    }
}

bool handleGlDrawElements(Stack* stack, bool pushReturn) {
    void* indices = stack->pop<void*>();
    IndicesType indices_type = stack->pop<IndicesType>();
    s32 element_count = stack->pop<s32>();
    DrawMode draw_mode = stack->pop<DrawMode>();
    if (stack->isValid()) {
        CAZE_DEBUG("glDrawElements(%u, %d, %u, %p)\n", draw_mode, element_count, indices_type,
                   indices);
        glDrawElements(draw_mode, element_count, indices_type, indices);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glDrawElements\n");
        return false;
    }
}

bool handleGlDrawArrays(Stack* stack, bool pushReturn) {
    s32 index_count = stack->pop<s32>();
    s32 first_index = stack->pop<s32>();
    DrawMode draw_mode = stack->pop<DrawMode>();
    if (stack->isValid()) {
        CAZE_DEBUG("glDrawArrays(%u, %d, %d)\n", draw_mode, first_index, index_count);
        glDrawArrays(draw_mode, first_index, index_count);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glDrawArrays\n");
        return false;
    }
}

bool handleGlFlush(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        CAZE_DEBUG("glFlush()\n");
        glFlush();
        return true;
    } else {
        CAZE_WARNING("Error during calling function glFlush\n");
        return false;
    }
}

bool handleGlFinish(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        CAZE_DEBUG("glFinish()\n");
        glFinish();
        return true;
    } else {
        CAZE_WARNING("Error during calling function glFinish\n");
        return false;
    }
}

bool handleGlGetBooleanv(Stack* stack, bool pushReturn) {
    bool* values = stack->pop<bool*>();
    StateVariable param = stack->pop<StateVariable>();
    if (stack->isValid()) {
        CAZE_DEBUG("glGetBooleanv(%u, %p)\n", param, values);
        glGetBooleanv(param, values);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glGetBooleanv\n");
        return false;
    }
}

bool handleGlGetFloatv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    StateVariable param = stack->pop<StateVariable>();
    if (stack->isValid()) {
        CAZE_DEBUG("glGetFloatv(%u, %p)\n", param, values);
        glGetFloatv(param, values);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glGetFloatv\n");
        return false;
    }
}

bool handleGlGetIntegerv(Stack* stack, bool pushReturn) {
    s32* values = stack->pop<s32*>();
    StateVariable param = stack->pop<StateVariable>();
    if (stack->isValid()) {
        CAZE_DEBUG("glGetIntegerv(%u, %p)\n", param, values);
        glGetIntegerv(param, values);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glGetIntegerv\n");
        return false;
    }
}

bool handleGlGetString(Stack* stack, bool pushReturn) {
    StringConstant param = stack->pop<StringConstant>();
    if (stack->isValid()) {
        CAZE_DEBUG("glGetString(%u)\n", param);
        const unsigned char* return_value = glGetString(param);
        if (pushReturn) {
            stack->push<const unsigned char*>(return_value);
        }
        return true;
    } else {
        CAZE_WARNING("Error during calling function glGetString\n");
        return false;
    }
}

bool handleGlEnable(Stack* stack, bool pushReturn) {
    Capability capability = stack->pop<Capability>();
    if (stack->isValid()) {
        CAZE_DEBUG("glEnable(%u)\n", capability);
        glEnable(capability);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glEnable\n");
        return false;
    }
}

bool handleGlDisable(Stack* stack, bool pushReturn) {
    Capability capability = stack->pop<Capability>();
    if (stack->isValid()) {
        CAZE_DEBUG("glDisable(%u)\n", capability);
        glDisable(capability);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glDisable\n");
        return false;
    }
}

bool handleGlIsEnabled(Stack* stack, bool pushReturn) {
    Capability capability = stack->pop<Capability>();
    if (stack->isValid()) {
        CAZE_DEBUG("glIsEnabled(%u)\n", capability);
        const bool return_value = glIsEnabled(capability);
        if (pushReturn) {
            stack->push<const bool>(return_value);
        }
        return true;
    } else {
        CAZE_WARNING("Error during calling function glIsEnabled\n");
        return false;
    }
}

bool handleGlMapBufferRange(Stack* stack, bool pushReturn) {
    MapBufferRangeAccess access = stack->pop<MapBufferRangeAccess>();
    s32 length = stack->pop<s32>();
    s32 offset = stack->pop<s32>();
    MapBufferTarget target = stack->pop<MapBufferTarget>();
    if (stack->isValid()) {
        CAZE_DEBUG("glMapBufferRange(%u, %d, %d, %u)\n", target, offset, length, access);
        const s32 return_value = glMapBufferRange(target, offset, length, access);
        if (pushReturn) {
            stack->push<const s32>(return_value);
        }
        return true;
    } else {
        CAZE_WARNING("Error during calling function glMapBufferRange\n");
        return false;
    }
}

bool handleGlUnmapBuffer(Stack* stack, bool pushReturn) {
    MapBufferTarget target = stack->pop<MapBufferTarget>();
    if (stack->isValid()) {
        CAZE_DEBUG("glUnmapBuffer(%u)\n", target);
        glUnmapBuffer(target);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glUnmapBuffer\n");
        return false;
    }
}

bool handleGlInvalidateFramebuffer(Stack* stack, bool pushReturn) {
    FramebufferAttachment* attachments = stack->pop<FramebufferAttachment*>();
    s32 count = stack->pop<s32>();
    FramebufferTarget target = stack->pop<FramebufferTarget>();
    if (stack->isValid()) {
        CAZE_DEBUG("glInvalidateFramebuffer(%u, %d, %p)\n", target, count, attachments);
        glInvalidateFramebuffer(target, count, attachments);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glInvalidateFramebuffer\n");
        return false;
    }
}

bool handleGlRenderbufferStorageMultisample(Stack* stack, bool pushReturn) {
    s32 height = stack->pop<s32>();
    s32 width = stack->pop<s32>();
    RenderbufferFormat format = stack->pop<RenderbufferFormat>();
    s32 samples = stack->pop<s32>();
    RenderbufferTarget target = stack->pop<RenderbufferTarget>();
    if (stack->isValid()) {
        CAZE_DEBUG("glRenderbufferStorageMultisample(%u, %d, %u, %d, %d)\n", target, samples,
                   format, width, height);
        glRenderbufferStorageMultisample(target, samples, format, width, height);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glRenderbufferStorageMultisample\n");
        return false;
    }
}

bool handleGlBlitFramebuffer(Stack* stack, bool pushReturn) {
    TextureFilterMode filter = stack->pop<TextureFilterMode>();
    ClearMask mask = stack->pop<ClearMask>();
    s32 dstY1 = stack->pop<s32>();
    s32 dstX1 = stack->pop<s32>();
    s32 dstY0 = stack->pop<s32>();
    s32 dstX0 = stack->pop<s32>();
    s32 srcY1 = stack->pop<s32>();
    s32 srcX1 = stack->pop<s32>();
    s32 srcY0 = stack->pop<s32>();
    s32 srcX0 = stack->pop<s32>();
    if (stack->isValid()) {
        CAZE_DEBUG("glBlitFramebuffer(%d, %d, %d, %d, %d, %d, %d, %d, %u, %u)\n", srcX0, srcY0,
                   srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter);
        glBlitFramebuffer(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter);
        return true;
    } else {
        CAZE_WARNING("Error during calling function glBlitFramebuffer\n");
        return false;
    }
}

}  // end of anonymous namespace

void GlFunctions::Register(Interpreter* interpreter) {
    interpreter->registerFunction(GlFunctions::EglCreateContext, handleEglCreateContext);
    interpreter->registerFunction(GlFunctions::EglMakeCurrent, handleEglMakeCurrent);
    interpreter->registerFunction(GlFunctions::EglSwapBuffers, handleEglSwapBuffers);
    interpreter->registerFunction(GlFunctions::GlEnableClientState, handleGlEnableClientState);
    interpreter->registerFunction(GlFunctions::GlDisableClientState, handleGlDisableClientState);
    interpreter->registerFunction(GlFunctions::GlGetProgramBinaryOES, handleGlGetProgramBinaryOES);
    interpreter->registerFunction(GlFunctions::GlProgramBinaryOES, handleGlProgramBinaryOES);
    interpreter->registerFunction(GlFunctions::GlStartTilingQCOM, handleGlStartTilingQCOM);
    interpreter->registerFunction(GlFunctions::GlEndTilingQCOM, handleGlEndTilingQCOM);
    interpreter->registerFunction(GlFunctions::GlDiscardFramebufferEXT,
                                  handleGlDiscardFramebufferEXT);
    interpreter->registerFunction(GlFunctions::GlInsertEventMarkerEXT,
                                  handleGlInsertEventMarkerEXT);
    interpreter->registerFunction(GlFunctions::GlPushGroupMarkerEXT, handleGlPushGroupMarkerEXT);
    interpreter->registerFunction(GlFunctions::GlPopGroupMarkerEXT, handleGlPopGroupMarkerEXT);
    interpreter->registerFunction(GlFunctions::GlTexStorage1DEXT, handleGlTexStorage1DEXT);
    interpreter->registerFunction(GlFunctions::GlTexStorage2DEXT, handleGlTexStorage2DEXT);
    interpreter->registerFunction(GlFunctions::GlTexStorage3DEXT, handleGlTexStorage3DEXT);
    interpreter->registerFunction(GlFunctions::GlTextureStorage1DEXT, handleGlTextureStorage1DEXT);
    interpreter->registerFunction(GlFunctions::GlTextureStorage2DEXT, handleGlTextureStorage2DEXT);
    interpreter->registerFunction(GlFunctions::GlTextureStorage3DEXT, handleGlTextureStorage3DEXT);
    interpreter->registerFunction(GlFunctions::GlGenVertexArraysOES, handleGlGenVertexArraysOES);
    interpreter->registerFunction(GlFunctions::GlBindVertexArrayOES, handleGlBindVertexArrayOES);
    interpreter->registerFunction(GlFunctions::GlDeleteVertexArraysOES,
                                  handleGlDeleteVertexArraysOES);
    interpreter->registerFunction(GlFunctions::GlIsVertexArrayOES, handleGlIsVertexArrayOES);
    interpreter->registerFunction(GlFunctions::GlEGLImageTargetTexture2DOES,
                                  handleGlEGLImageTargetTexture2DOES);
    interpreter->registerFunction(GlFunctions::GlEGLImageTargetRenderbufferStorageOES,
                                  handleGlEGLImageTargetRenderbufferStorageOES);
    interpreter->registerFunction(GlFunctions::GlGetGraphicsResetStatusEXT,
                                  handleGlGetGraphicsResetStatusEXT);
    interpreter->registerFunction(GlFunctions::GlBindAttribLocation, handleGlBindAttribLocation);
    interpreter->registerFunction(GlFunctions::GlBlendFunc, handleGlBlendFunc);
    interpreter->registerFunction(GlFunctions::GlBlendFuncSeparate, handleGlBlendFuncSeparate);
    interpreter->registerFunction(GlFunctions::GlBlendEquation, handleGlBlendEquation);
    interpreter->registerFunction(GlFunctions::GlBlendEquationSeparate,
                                  handleGlBlendEquationSeparate);
    interpreter->registerFunction(GlFunctions::GlBlendColor, handleGlBlendColor);
    interpreter->registerFunction(GlFunctions::GlEnableVertexAttribArray,
                                  handleGlEnableVertexAttribArray);
    interpreter->registerFunction(GlFunctions::GlDisableVertexAttribArray,
                                  handleGlDisableVertexAttribArray);
    interpreter->registerFunction(GlFunctions::GlVertexAttribPointer, handleGlVertexAttribPointer);
    interpreter->registerFunction(GlFunctions::GlGetActiveAttrib, handleGlGetActiveAttrib);
    interpreter->registerFunction(GlFunctions::GlGetActiveUniform, handleGlGetActiveUniform);
    interpreter->registerFunction(GlFunctions::GlGetError, handleGlGetError);
    interpreter->registerFunction(GlFunctions::GlGetProgramiv, handleGlGetProgramiv);
    interpreter->registerFunction(GlFunctions::GlGetShaderiv, handleGlGetShaderiv);
    interpreter->registerFunction(GlFunctions::GlGetUniformLocation, handleGlGetUniformLocation);
    interpreter->registerFunction(GlFunctions::GlGetAttribLocation, handleGlGetAttribLocation);
    interpreter->registerFunction(GlFunctions::GlPixelStorei, handleGlPixelStorei);
    interpreter->registerFunction(GlFunctions::GlTexParameteri, handleGlTexParameteri);
    interpreter->registerFunction(GlFunctions::GlTexParameterf, handleGlTexParameterf);
    interpreter->registerFunction(GlFunctions::GlGetTexParameteriv, handleGlGetTexParameteriv);
    interpreter->registerFunction(GlFunctions::GlGetTexParameterfv, handleGlGetTexParameterfv);
    interpreter->registerFunction(GlFunctions::GlUniform1i, handleGlUniform1i);
    interpreter->registerFunction(GlFunctions::GlUniform2i, handleGlUniform2i);
    interpreter->registerFunction(GlFunctions::GlUniform3i, handleGlUniform3i);
    interpreter->registerFunction(GlFunctions::GlUniform4i, handleGlUniform4i);
    interpreter->registerFunction(GlFunctions::GlUniform1iv, handleGlUniform1iv);
    interpreter->registerFunction(GlFunctions::GlUniform2iv, handleGlUniform2iv);
    interpreter->registerFunction(GlFunctions::GlUniform3iv, handleGlUniform3iv);
    interpreter->registerFunction(GlFunctions::GlUniform4iv, handleGlUniform4iv);
    interpreter->registerFunction(GlFunctions::GlUniform1f, handleGlUniform1f);
    interpreter->registerFunction(GlFunctions::GlUniform2f, handleGlUniform2f);
    interpreter->registerFunction(GlFunctions::GlUniform3f, handleGlUniform3f);
    interpreter->registerFunction(GlFunctions::GlUniform4f, handleGlUniform4f);
    interpreter->registerFunction(GlFunctions::GlUniform1fv, handleGlUniform1fv);
    interpreter->registerFunction(GlFunctions::GlUniform2fv, handleGlUniform2fv);
    interpreter->registerFunction(GlFunctions::GlUniform3fv, handleGlUniform3fv);
    interpreter->registerFunction(GlFunctions::GlUniform4fv, handleGlUniform4fv);
    interpreter->registerFunction(GlFunctions::GlUniformMatrix2fv, handleGlUniformMatrix2fv);
    interpreter->registerFunction(GlFunctions::GlUniformMatrix3fv, handleGlUniformMatrix3fv);
    interpreter->registerFunction(GlFunctions::GlUniformMatrix4fv, handleGlUniformMatrix4fv);
    interpreter->registerFunction(GlFunctions::GlGetUniformfv, handleGlGetUniformfv);
    interpreter->registerFunction(GlFunctions::GlGetUniformiv, handleGlGetUniformiv);
    interpreter->registerFunction(GlFunctions::GlVertexAttrib1f, handleGlVertexAttrib1f);
    interpreter->registerFunction(GlFunctions::GlVertexAttrib2f, handleGlVertexAttrib2f);
    interpreter->registerFunction(GlFunctions::GlVertexAttrib3f, handleGlVertexAttrib3f);
    interpreter->registerFunction(GlFunctions::GlVertexAttrib4f, handleGlVertexAttrib4f);
    interpreter->registerFunction(GlFunctions::GlVertexAttrib1fv, handleGlVertexAttrib1fv);
    interpreter->registerFunction(GlFunctions::GlVertexAttrib2fv, handleGlVertexAttrib2fv);
    interpreter->registerFunction(GlFunctions::GlVertexAttrib3fv, handleGlVertexAttrib3fv);
    interpreter->registerFunction(GlFunctions::GlVertexAttrib4fv, handleGlVertexAttrib4fv);
    interpreter->registerFunction(GlFunctions::GlGetShaderPrecisionFormat,
                                  handleGlGetShaderPrecisionFormat);
    interpreter->registerFunction(GlFunctions::GlDepthMask, handleGlDepthMask);
    interpreter->registerFunction(GlFunctions::GlDepthFunc, handleGlDepthFunc);
    interpreter->registerFunction(GlFunctions::GlDepthRangef, handleGlDepthRangef);
    interpreter->registerFunction(GlFunctions::GlColorMask, handleGlColorMask);
    interpreter->registerFunction(GlFunctions::GlStencilMask, handleGlStencilMask);
    interpreter->registerFunction(GlFunctions::GlStencilMaskSeparate, handleGlStencilMaskSeparate);
    interpreter->registerFunction(GlFunctions::GlStencilFuncSeparate, handleGlStencilFuncSeparate);
    interpreter->registerFunction(GlFunctions::GlStencilOpSeparate, handleGlStencilOpSeparate);
    interpreter->registerFunction(GlFunctions::GlFrontFace, handleGlFrontFace);
    interpreter->registerFunction(GlFunctions::GlViewport, handleGlViewport);
    interpreter->registerFunction(GlFunctions::GlScissor, handleGlScissor);
    interpreter->registerFunction(GlFunctions::GlActiveTexture, handleGlActiveTexture);
    interpreter->registerFunction(GlFunctions::GlGenTextures, handleGlGenTextures);
    interpreter->registerFunction(GlFunctions::GlDeleteTextures, handleGlDeleteTextures);
    interpreter->registerFunction(GlFunctions::GlIsTexture, handleGlIsTexture);
    interpreter->registerFunction(GlFunctions::GlBindTexture, handleGlBindTexture);
    interpreter->registerFunction(GlFunctions::GlTexImage2D, handleGlTexImage2D);
    interpreter->registerFunction(GlFunctions::GlTexSubImage2D, handleGlTexSubImage2D);
    interpreter->registerFunction(GlFunctions::GlCopyTexImage2D, handleGlCopyTexImage2D);
    interpreter->registerFunction(GlFunctions::GlCopyTexSubImage2D, handleGlCopyTexSubImage2D);
    interpreter->registerFunction(GlFunctions::GlCompressedTexImage2D,
                                  handleGlCompressedTexImage2D);
    interpreter->registerFunction(GlFunctions::GlCompressedTexSubImage2D,
                                  handleGlCompressedTexSubImage2D);
    interpreter->registerFunction(GlFunctions::GlGenerateMipmap, handleGlGenerateMipmap);
    interpreter->registerFunction(GlFunctions::GlReadPixels, handleGlReadPixels);
    interpreter->registerFunction(GlFunctions::GlGenFramebuffers, handleGlGenFramebuffers);
    interpreter->registerFunction(GlFunctions::GlBindFramebuffer, handleGlBindFramebuffer);
    interpreter->registerFunction(GlFunctions::GlCheckFramebufferStatus,
                                  handleGlCheckFramebufferStatus);
    interpreter->registerFunction(GlFunctions::GlDeleteFramebuffers, handleGlDeleteFramebuffers);
    interpreter->registerFunction(GlFunctions::GlIsFramebuffer, handleGlIsFramebuffer);
    interpreter->registerFunction(GlFunctions::GlGenRenderbuffers, handleGlGenRenderbuffers);
    interpreter->registerFunction(GlFunctions::GlBindRenderbuffer, handleGlBindRenderbuffer);
    interpreter->registerFunction(GlFunctions::GlRenderbufferStorage, handleGlRenderbufferStorage);
    interpreter->registerFunction(GlFunctions::GlDeleteRenderbuffers, handleGlDeleteRenderbuffers);
    interpreter->registerFunction(GlFunctions::GlIsRenderbuffer, handleGlIsRenderbuffer);
    interpreter->registerFunction(GlFunctions::GlGetRenderbufferParameteriv,
                                  handleGlGetRenderbufferParameteriv);
    interpreter->registerFunction(GlFunctions::GlGenBuffers, handleGlGenBuffers);
    interpreter->registerFunction(GlFunctions::GlBindBuffer, handleGlBindBuffer);
    interpreter->registerFunction(GlFunctions::GlBufferData, handleGlBufferData);
    interpreter->registerFunction(GlFunctions::GlBufferSubData, handleGlBufferSubData);
    interpreter->registerFunction(GlFunctions::GlDeleteBuffers, handleGlDeleteBuffers);
    interpreter->registerFunction(GlFunctions::GlIsBuffer, handleGlIsBuffer);
    interpreter->registerFunction(GlFunctions::GlGetBufferParameteriv,
                                  handleGlGetBufferParameteriv);
    interpreter->registerFunction(GlFunctions::GlCreateShader, handleGlCreateShader);
    interpreter->registerFunction(GlFunctions::GlDeleteShader, handleGlDeleteShader);
    interpreter->registerFunction(GlFunctions::GlShaderSource, handleGlShaderSource);
    interpreter->registerFunction(GlFunctions::GlShaderBinary, handleGlShaderBinary);
    interpreter->registerFunction(GlFunctions::GlGetShaderInfoLog, handleGlGetShaderInfoLog);
    interpreter->registerFunction(GlFunctions::GlGetShaderSource, handleGlGetShaderSource);
    interpreter->registerFunction(GlFunctions::GlReleaseShaderCompiler,
                                  handleGlReleaseShaderCompiler);
    interpreter->registerFunction(GlFunctions::GlCompileShader, handleGlCompileShader);
    interpreter->registerFunction(GlFunctions::GlIsShader, handleGlIsShader);
    interpreter->registerFunction(GlFunctions::GlCreateProgram, handleGlCreateProgram);
    interpreter->registerFunction(GlFunctions::GlDeleteProgram, handleGlDeleteProgram);
    interpreter->registerFunction(GlFunctions::GlAttachShader, handleGlAttachShader);
    interpreter->registerFunction(GlFunctions::GlDetachShader, handleGlDetachShader);
    interpreter->registerFunction(GlFunctions::GlGetAttachedShaders, handleGlGetAttachedShaders);
    interpreter->registerFunction(GlFunctions::GlLinkProgram, handleGlLinkProgram);
    interpreter->registerFunction(GlFunctions::GlGetProgramInfoLog, handleGlGetProgramInfoLog);
    interpreter->registerFunction(GlFunctions::GlUseProgram, handleGlUseProgram);
    interpreter->registerFunction(GlFunctions::GlIsProgram, handleGlIsProgram);
    interpreter->registerFunction(GlFunctions::GlValidateProgram, handleGlValidateProgram);
    interpreter->registerFunction(GlFunctions::GlClearColor, handleGlClearColor);
    interpreter->registerFunction(GlFunctions::GlClearDepthf, handleGlClearDepthf);
    interpreter->registerFunction(GlFunctions::GlClearStencil, handleGlClearStencil);
    interpreter->registerFunction(GlFunctions::GlClear, handleGlClear);
    interpreter->registerFunction(GlFunctions::GlCullFace, handleGlCullFace);
    interpreter->registerFunction(GlFunctions::GlPolygonOffset, handleGlPolygonOffset);
    interpreter->registerFunction(GlFunctions::GlLineWidth, handleGlLineWidth);
    interpreter->registerFunction(GlFunctions::GlSampleCoverage, handleGlSampleCoverage);
    interpreter->registerFunction(GlFunctions::GlHint, handleGlHint);
    interpreter->registerFunction(GlFunctions::GlFramebufferRenderbuffer,
                                  handleGlFramebufferRenderbuffer);
    interpreter->registerFunction(GlFunctions::GlFramebufferTexture2D,
                                  handleGlFramebufferTexture2D);
    interpreter->registerFunction(GlFunctions::GlGetFramebufferAttachmentParameteriv,
                                  handleGlGetFramebufferAttachmentParameteriv);
    interpreter->registerFunction(GlFunctions::GlDrawElements, handleGlDrawElements);
    interpreter->registerFunction(GlFunctions::GlDrawArrays, handleGlDrawArrays);
    interpreter->registerFunction(GlFunctions::GlFlush, handleGlFlush);
    interpreter->registerFunction(GlFunctions::GlFinish, handleGlFinish);
    interpreter->registerFunction(GlFunctions::GlGetBooleanv, handleGlGetBooleanv);
    interpreter->registerFunction(GlFunctions::GlGetFloatv, handleGlGetFloatv);
    interpreter->registerFunction(GlFunctions::GlGetIntegerv, handleGlGetIntegerv);
    interpreter->registerFunction(GlFunctions::GlGetString, handleGlGetString);
    interpreter->registerFunction(GlFunctions::GlEnable, handleGlEnable);
    interpreter->registerFunction(GlFunctions::GlDisable, handleGlDisable);
    interpreter->registerFunction(GlFunctions::GlIsEnabled, handleGlIsEnabled);
    interpreter->registerFunction(GlFunctions::GlMapBufferRange, handleGlMapBufferRange);
    interpreter->registerFunction(GlFunctions::GlUnmapBuffer, handleGlUnmapBuffer);
    interpreter->registerFunction(GlFunctions::GlInvalidateFramebuffer,
                                  handleGlInvalidateFramebuffer);
    interpreter->registerFunction(GlFunctions::GlRenderbufferStorageMultisample,
                                  handleGlRenderbufferStorageMultisample);
    interpreter->registerFunction(GlFunctions::GlBlitFramebuffer, handleGlBlitFramebuffer);
}

}  // end of namespace caze
}  // end of namespace android
