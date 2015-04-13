////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

#include "gfx_api.h"
#include "interpreter.h"
#include "stack.h"

#include <gapic/get_gfx_proc_address.h>
#include <gapic/log.h>

namespace gapir {
namespace gfxapi {
namespace {

bool callEglCreateContext(Stack* stack, bool pushReturn) {
    int32_t* context = stack->pop<int32_t*>();
    int32_t* version = stack->pop<int32_t*>();
    if (stack->isValid()) {
        GAPID_INFO("eglCreateContext(%p, %p)\n", version, context);
        if (eglCreateContext != nullptr) {
            eglCreateContext(version, context);
        } else {
            GAPID_WARNING("Attempted to call unsupported function eglCreateContext\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function eglCreateContext\n");
        return false;
    }
}

bool callEglMakeCurrent(Stack* stack, bool pushReturn) {
    int32_t context = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("eglMakeCurrent(%d)\n", context);
        if (eglMakeCurrent != nullptr) {
            eglMakeCurrent(context);
        } else {
            GAPID_WARNING("Attempted to call unsupported function eglMakeCurrent\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function eglMakeCurrent\n");
        return false;
    }
}

bool callEglSwapBuffers(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("eglSwapBuffers()\n");
        if (eglSwapBuffers != nullptr) {
            eglSwapBuffers();
        } else {
            GAPID_WARNING("Attempted to call unsupported function eglSwapBuffers\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function eglSwapBuffers\n");
        return false;
    }
}

bool callWglSwapBuffers(Stack* stack, bool pushReturn) {
    HDC hdc = stack->pop<HDC>();
    if (stack->isValid()) {
        GAPID_INFO("wglSwapBuffers(%p)\n", hdc);
        if (wglSwapBuffers != nullptr) {
            wglSwapBuffers(hdc);
        } else {
            GAPID_WARNING("Attempted to call unsupported function wglSwapBuffers\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function wglSwapBuffers\n");
        return false;
    }
}

bool callGlEnableClientState(Stack* stack, bool pushReturn) {
    ArrayType type = stack->pop<ArrayType>();
    if (stack->isValid()) {
        GAPID_INFO("glEnableClientState(%u)\n", type);
        if (glEnableClientState != nullptr) {
            glEnableClientState(type);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEnableClientState\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEnableClientState\n");
        return false;
    }
}

bool callGlDisableClientState(Stack* stack, bool pushReturn) {
    ArrayType type = stack->pop<ArrayType>();
    if (stack->isValid()) {
        GAPID_INFO("glDisableClientState(%u)\n", type);
        if (glDisableClientState != nullptr) {
            glDisableClientState(type);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDisableClientState\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDisableClientState\n");
        return false;
    }
}

bool callGlGetProgramBinaryOES(Stack* stack, bool pushReturn) {
    void* binary = stack->pop<void*>();
    uint32_t* binary_format = stack->pop<uint32_t*>();
    int32_t* bytes_written = stack->pop<int32_t*>();
    int32_t buffer_size = stack->pop<int32_t>();
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        GAPID_INFO("glGetProgramBinaryOES(%u, %d, %p, %p, %p)\n", program, buffer_size,
                   bytes_written, binary_format, binary);
        if (glGetProgramBinaryOES != nullptr) {
            glGetProgramBinaryOES(program, buffer_size, bytes_written, binary_format, binary);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetProgramBinaryOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetProgramBinaryOES\n");
        return false;
    }
}

bool callGlProgramBinaryOES(Stack* stack, bool pushReturn) {
    int32_t binary_size = stack->pop<int32_t>();
    const void* binary = stack->pop<const void*>();
    uint32_t binary_format = stack->pop<uint32_t>();
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramBinaryOES(%u, %u, %p, %d)\n", program, binary_format, binary,
                   binary_size);
        if (glProgramBinaryOES != nullptr) {
            glProgramBinaryOES(program, binary_format, binary, binary_size);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramBinaryOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramBinaryOES\n");
        return false;
    }
}

bool callGlStartTilingQCOM(Stack* stack, bool pushReturn) {
    TilePreserveMaskQCOM preserveMask = stack->pop<TilePreserveMaskQCOM>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    int32_t y = stack->pop<int32_t>();
    int32_t x = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glStartTilingQCOM(%d, %d, %d, %d, %u)\n", x, y, width, height, preserveMask);
        if (glStartTilingQCOM != nullptr) {
            glStartTilingQCOM(x, y, width, height, preserveMask);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glStartTilingQCOM\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStartTilingQCOM\n");
        return false;
    }
}

bool callGlEndTilingQCOM(Stack* stack, bool pushReturn) {
    TilePreserveMaskQCOM preserve_mask = stack->pop<TilePreserveMaskQCOM>();
    if (stack->isValid()) {
        GAPID_INFO("glEndTilingQCOM(%u)\n", preserve_mask);
        if (glEndTilingQCOM != nullptr) {
            glEndTilingQCOM(preserve_mask);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEndTilingQCOM\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEndTilingQCOM\n");
        return false;
    }
}

bool callGlDiscardFramebufferEXT(Stack* stack, bool pushReturn) {
    const DiscardFramebufferAttachment* attachments =
            stack->pop<const DiscardFramebufferAttachment*>();
    int32_t numAttachments = stack->pop<int32_t>();
    FramebufferTarget target = stack->pop<FramebufferTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glDiscardFramebufferEXT(%u, %d, %p)\n", target, numAttachments, attachments);
        if (glDiscardFramebufferEXT != nullptr) {
            glDiscardFramebufferEXT(target, numAttachments, attachments);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDiscardFramebufferEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDiscardFramebufferEXT\n");
        return false;
    }
}

bool callGlInsertEventMarkerEXT(Stack* stack, bool pushReturn) {
    const char* marker = stack->pop<const char*>();
    int32_t length = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glInsertEventMarkerEXT(%d, %s)\n", length, marker);
        if (glInsertEventMarkerEXT != nullptr) {
            glInsertEventMarkerEXT(length, marker);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glInsertEventMarkerEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glInsertEventMarkerEXT\n");
        return false;
    }
}

bool callGlPushGroupMarkerEXT(Stack* stack, bool pushReturn) {
    const char* marker = stack->pop<const char*>();
    int32_t length = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glPushGroupMarkerEXT(%d, %s)\n", length, marker);
        if (glPushGroupMarkerEXT != nullptr) {
            glPushGroupMarkerEXT(length, marker);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPushGroupMarkerEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPushGroupMarkerEXT\n");
        return false;
    }
}

bool callGlPopGroupMarkerEXT(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glPopGroupMarkerEXT()\n");
        if (glPopGroupMarkerEXT != nullptr) {
            glPopGroupMarkerEXT();
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPopGroupMarkerEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPopGroupMarkerEXT\n");
        return false;
    }
}

bool callGlTexStorage1DEXT(Stack* stack, bool pushReturn) {
    int32_t width = stack->pop<int32_t>();
    TexelFormat format = stack->pop<TexelFormat>();
    int32_t levels = stack->pop<int32_t>();
    TextureTarget target = stack->pop<TextureTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glTexStorage1DEXT(%u, %d, %u, %d)\n", target, levels, format, width);
        if (glTexStorage1DEXT != nullptr) {
            glTexStorage1DEXT(target, levels, format, width);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexStorage1DEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexStorage1DEXT\n");
        return false;
    }
}

bool callGlTexStorage2DEXT(Stack* stack, bool pushReturn) {
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    TexelFormat format = stack->pop<TexelFormat>();
    int32_t levels = stack->pop<int32_t>();
    TextureTarget target = stack->pop<TextureTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glTexStorage2DEXT(%u, %d, %u, %d, %d)\n", target, levels, format, width,
                   height);
        if (glTexStorage2DEXT != nullptr) {
            glTexStorage2DEXT(target, levels, format, width, height);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexStorage2DEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexStorage2DEXT\n");
        return false;
    }
}

bool callGlTexStorage3DEXT(Stack* stack, bool pushReturn) {
    int32_t depth = stack->pop<int32_t>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    TexelFormat format = stack->pop<TexelFormat>();
    int32_t levels = stack->pop<int32_t>();
    TextureTarget target = stack->pop<TextureTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glTexStorage3DEXT(%u, %d, %u, %d, %d, %d)\n", target, levels, format, width,
                   height, depth);
        if (glTexStorage3DEXT != nullptr) {
            glTexStorage3DEXT(target, levels, format, width, height, depth);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexStorage3DEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexStorage3DEXT\n");
        return false;
    }
}

bool callGlTextureStorage1DEXT(Stack* stack, bool pushReturn) {
    int32_t width = stack->pop<int32_t>();
    TexelFormat format = stack->pop<TexelFormat>();
    int32_t levels = stack->pop<int32_t>();
    TextureTarget target = stack->pop<TextureTarget>();
    TextureId texture = stack->pop<TextureId>();
    if (stack->isValid()) {
        GAPID_INFO("glTextureStorage1DEXT(%u, %u, %d, %u, %d)\n", texture, target, levels, format,
                   width);
        if (glTextureStorage1DEXT != nullptr) {
            glTextureStorage1DEXT(texture, target, levels, format, width);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTextureStorage1DEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTextureStorage1DEXT\n");
        return false;
    }
}

bool callGlTextureStorage2DEXT(Stack* stack, bool pushReturn) {
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    TexelFormat format = stack->pop<TexelFormat>();
    int32_t levels = stack->pop<int32_t>();
    TextureTarget target = stack->pop<TextureTarget>();
    TextureId texture = stack->pop<TextureId>();
    if (stack->isValid()) {
        GAPID_INFO("glTextureStorage2DEXT(%u, %u, %d, %u, %d, %d)\n", texture, target, levels,
                   format, width, height);
        if (glTextureStorage2DEXT != nullptr) {
            glTextureStorage2DEXT(texture, target, levels, format, width, height);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTextureStorage2DEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTextureStorage2DEXT\n");
        return false;
    }
}

bool callGlTextureStorage3DEXT(Stack* stack, bool pushReturn) {
    int32_t depth = stack->pop<int32_t>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    TexelFormat format = stack->pop<TexelFormat>();
    int32_t levels = stack->pop<int32_t>();
    TextureTarget target = stack->pop<TextureTarget>();
    TextureId texture = stack->pop<TextureId>();
    if (stack->isValid()) {
        GAPID_INFO("glTextureStorage3DEXT(%u, %u, %d, %u, %d, %d, %d)\n", texture, target, levels,
                   format, width, height, depth);
        if (glTextureStorage3DEXT != nullptr) {
            glTextureStorage3DEXT(texture, target, levels, format, width, height, depth);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTextureStorage3DEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTextureStorage3DEXT\n");
        return false;
    }
}

bool callGlGenVertexArraysOES(Stack* stack, bool pushReturn) {
    VertexArrayId* arrays = stack->pop<VertexArrayId*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenVertexArraysOES(%d, %p)\n", count, arrays);
        if (glGenVertexArraysOES != nullptr) {
            glGenVertexArraysOES(count, arrays);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenVertexArraysOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenVertexArraysOES\n");
        return false;
    }
}

bool callGlBindVertexArrayOES(Stack* stack, bool pushReturn) {
    VertexArrayId array = stack->pop<VertexArrayId>();
    if (stack->isValid()) {
        GAPID_INFO("glBindVertexArrayOES(%u)\n", array);
        if (glBindVertexArrayOES != nullptr) {
            glBindVertexArrayOES(array);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindVertexArrayOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindVertexArrayOES\n");
        return false;
    }
}

bool callGlDeleteVertexArraysOES(Stack* stack, bool pushReturn) {
    const VertexArrayId* arrays = stack->pop<const VertexArrayId*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteVertexArraysOES(%d, %p)\n", count, arrays);
        if (glDeleteVertexArraysOES != nullptr) {
            glDeleteVertexArraysOES(count, arrays);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteVertexArraysOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteVertexArraysOES\n");
        return false;
    }
}

bool callGlIsVertexArrayOES(Stack* stack, bool pushReturn) {
    VertexArrayId array = stack->pop<VertexArrayId>();
    if (stack->isValid()) {
        GAPID_INFO("glIsVertexArrayOES(%u)\n", array);
        if (glIsVertexArrayOES != nullptr) {
            bool return_value = glIsVertexArrayOES(array);
            GAPID_INFO("Returned: %d\n", return_value);
            if (pushReturn) {
                stack->push<bool>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsVertexArrayOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsVertexArrayOES\n");
        return false;
    }
}

bool callGlEGLImageTargetTexture2DOES(Stack* stack, bool pushReturn) {
    ImageOES image = stack->pop<ImageOES>();
    ImageTargetTexture target = stack->pop<ImageTargetTexture>();
    if (stack->isValid()) {
        GAPID_INFO("glEGLImageTargetTexture2DOES(%u, %p)\n", target, image);
        if (glEGLImageTargetTexture2DOES != nullptr) {
            glEGLImageTargetTexture2DOES(target, image);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEGLImageTargetTexture2DOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEGLImageTargetTexture2DOES\n");
        return false;
    }
}

bool callGlEGLImageTargetRenderbufferStorageOES(Stack* stack, bool pushReturn) {
    TexturePointer image = stack->pop<TexturePointer>();
    ImageTargetRenderbufferStorage target = stack->pop<ImageTargetRenderbufferStorage>();
    if (stack->isValid()) {
        GAPID_INFO("glEGLImageTargetRenderbufferStorageOES(%u, %p)\n", target, image);
        if (glEGLImageTargetRenderbufferStorageOES != nullptr) {
            glEGLImageTargetRenderbufferStorageOES(target, image);
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function "
                    "glEGLImageTargetRenderbufferStorageOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEGLImageTargetRenderbufferStorageOES\n");
        return false;
    }
}

bool callGlGetGraphicsResetStatusEXT(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glGetGraphicsResetStatusEXT()\n");
        if (glGetGraphicsResetStatusEXT != nullptr) {
            ResetStatus return_value = glGetGraphicsResetStatusEXT();
            GAPID_INFO("Returned: %u\n", return_value);
            if (pushReturn) {
                stack->push<ResetStatus>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetGraphicsResetStatusEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetGraphicsResetStatusEXT\n");
        return false;
    }
}

bool callGlBindAttribLocation(Stack* stack, bool pushReturn) {
    const char* name = stack->pop<const char*>();
    AttributeLocation location = stack->pop<AttributeLocation>();
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        GAPID_INFO("glBindAttribLocation(%u, %u, %s)\n", program, location, name);
        if (glBindAttribLocation != nullptr) {
            glBindAttribLocation(program, location, name);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindAttribLocation\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindAttribLocation\n");
        return false;
    }
}

bool callGlBlendFunc(Stack* stack, bool pushReturn) {
    BlendFactor dst_factor = stack->pop<BlendFactor>();
    BlendFactor src_factor = stack->pop<BlendFactor>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendFunc(%u, %u)\n", src_factor, dst_factor);
        if (glBlendFunc != nullptr) {
            glBlendFunc(src_factor, dst_factor);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendFunc\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendFunc\n");
        return false;
    }
}

bool callGlBlendFuncSeparate(Stack* stack, bool pushReturn) {
    BlendFactor dst_factor_alpha = stack->pop<BlendFactor>();
    BlendFactor src_factor_alpha = stack->pop<BlendFactor>();
    BlendFactor dst_factor_rgb = stack->pop<BlendFactor>();
    BlendFactor src_factor_rgb = stack->pop<BlendFactor>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendFuncSeparate(%u, %u, %u, %u)\n", src_factor_rgb, dst_factor_rgb,
                   src_factor_alpha, dst_factor_alpha);
        if (glBlendFuncSeparate != nullptr) {
            glBlendFuncSeparate(src_factor_rgb, dst_factor_rgb, src_factor_alpha, dst_factor_alpha);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendFuncSeparate\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendFuncSeparate\n");
        return false;
    }
}

bool callGlBlendEquation(Stack* stack, bool pushReturn) {
    BlendEquation equation = stack->pop<BlendEquation>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendEquation(%u)\n", equation);
        if (glBlendEquation != nullptr) {
            glBlendEquation(equation);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendEquation\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendEquation\n");
        return false;
    }
}

bool callGlBlendEquationSeparate(Stack* stack, bool pushReturn) {
    BlendEquation alpha = stack->pop<BlendEquation>();
    BlendEquation rgb = stack->pop<BlendEquation>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendEquationSeparate(%u, %u)\n", rgb, alpha);
        if (glBlendEquationSeparate != nullptr) {
            glBlendEquationSeparate(rgb, alpha);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendEquationSeparate\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendEquationSeparate\n");
        return false;
    }
}

bool callGlBlendColor(Stack* stack, bool pushReturn) {
    float alpha = stack->pop<float>();
    float blue = stack->pop<float>();
    float green = stack->pop<float>();
    float red = stack->pop<float>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendColor(%f, %f, %f, %f)\n", red, green, blue, alpha);
        if (glBlendColor != nullptr) {
            glBlendColor(red, green, blue, alpha);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendColor\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendColor\n");
        return false;
    }
}

bool callGlEnableVertexAttribArray(Stack* stack, bool pushReturn) {
    AttributeLocation location = stack->pop<AttributeLocation>();
    if (stack->isValid()) {
        GAPID_INFO("glEnableVertexAttribArray(%u)\n", location);
        if (glEnableVertexAttribArray != nullptr) {
            glEnableVertexAttribArray(location);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEnableVertexAttribArray\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEnableVertexAttribArray\n");
        return false;
    }
}

bool callGlDisableVertexAttribArray(Stack* stack, bool pushReturn) {
    AttributeLocation location = stack->pop<AttributeLocation>();
    if (stack->isValid()) {
        GAPID_INFO("glDisableVertexAttribArray(%u)\n", location);
        if (glDisableVertexAttribArray != nullptr) {
            glDisableVertexAttribArray(location);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDisableVertexAttribArray\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDisableVertexAttribArray\n");
        return false;
    }
}

bool callGlVertexAttribPointer(Stack* stack, bool pushReturn) {
    VertexPointer data = stack->pop<VertexPointer>();
    int32_t stride = stack->pop<int32_t>();
    bool normalized = stack->pop<bool>();
    VertexAttribType type = stack->pop<VertexAttribType>();
    int32_t size = stack->pop<int32_t>();
    AttributeLocation location = stack->pop<AttributeLocation>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttribPointer(%u, %d, %u, %d, %d, %p)\n", location, size, type,
                   normalized, stride, data);
        if (glVertexAttribPointer != nullptr) {
            glVertexAttribPointer(location, size, type, normalized, stride, data);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttribPointer\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttribPointer\n");
        return false;
    }
}

bool callGlGetActiveAttrib(Stack* stack, bool pushReturn) {
    char* name = stack->pop<char*>();
    ShaderAttribType* type = stack->pop<ShaderAttribType*>();
    int32_t* vector_count = stack->pop<int32_t*>();
    int32_t* buffer_bytes_written = stack->pop<int32_t*>();
    int32_t buffer_size = stack->pop<int32_t>();
    AttributeLocation location = stack->pop<AttributeLocation>();
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        GAPID_INFO("glGetActiveAttrib(%u, %u, %d, %p, %p, %p, %p)\n", program, location,
                   buffer_size, buffer_bytes_written, vector_count, type, name);
        if (glGetActiveAttrib != nullptr) {
            glGetActiveAttrib(program, location, buffer_size, buffer_bytes_written, vector_count,
                              type, name);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetActiveAttrib\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetActiveAttrib\n");
        return false;
    }
}

bool callGlGetActiveUniform(Stack* stack, bool pushReturn) {
    char* name = stack->pop<char*>();
    ShaderUniformType* type = stack->pop<ShaderUniformType*>();
    int32_t* size = stack->pop<int32_t*>();
    int32_t* buffer_bytes_written = stack->pop<int32_t*>();
    int32_t buffer_size = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        GAPID_INFO("glGetActiveUniform(%u, %d, %d, %p, %p, %p, %p)\n", program, location,
                   buffer_size, buffer_bytes_written, size, type, name);
        if (glGetActiveUniform != nullptr) {
            glGetActiveUniform(program, location, buffer_size, buffer_bytes_written, size, type,
                               name);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetActiveUniform\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetActiveUniform\n");
        return false;
    }
}

bool callGlGetError(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glGetError()\n");
        if (glGetError != nullptr) {
            Error return_value = glGetError();
            GAPID_INFO("Returned: %u\n", return_value);
            if (pushReturn) {
                stack->push<Error>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetError\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetError\n");
        return false;
    }
}

bool callGlGetProgramiv(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    ProgramParameter parameter = stack->pop<ProgramParameter>();
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        GAPID_INFO("glGetProgramiv(%u, %u, %p)\n", program, parameter, value);
        if (glGetProgramiv != nullptr) {
            glGetProgramiv(program, parameter, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetProgramiv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetProgramiv\n");
        return false;
    }
}

bool callGlGetShaderiv(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    ShaderParameter parameter = stack->pop<ShaderParameter>();
    ShaderId shader = stack->pop<ShaderId>();
    if (stack->isValid()) {
        GAPID_INFO("glGetShaderiv(%u, %u, %p)\n", shader, parameter, value);
        if (glGetShaderiv != nullptr) {
            glGetShaderiv(shader, parameter, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetShaderiv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetShaderiv\n");
        return false;
    }
}

bool callGlGetUniformLocation(Stack* stack, bool pushReturn) {
    const char* name = stack->pop<const char*>();
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        GAPID_INFO("glGetUniformLocation(%u, %s)\n", program, name);
        if (glGetUniformLocation != nullptr) {
            UniformLocation return_value = glGetUniformLocation(program, name);
            GAPID_INFO("Returned: %d\n", return_value);
            if (pushReturn) {
                stack->push<UniformLocation>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetUniformLocation\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetUniformLocation\n");
        return false;
    }
}

bool callGlGetAttribLocation(Stack* stack, bool pushReturn) {
    const char* name = stack->pop<const char*>();
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        GAPID_INFO("glGetAttribLocation(%u, %s)\n", program, name);
        if (glGetAttribLocation != nullptr) {
            AttributeLocation return_value = glGetAttribLocation(program, name);
            GAPID_INFO("Returned: %u\n", return_value);
            if (pushReturn) {
                stack->push<AttributeLocation>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetAttribLocation\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetAttribLocation\n");
        return false;
    }
}

bool callGlPixelStorei(Stack* stack, bool pushReturn) {
    int32_t value = stack->pop<int32_t>();
    PixelStoreParameter parameter = stack->pop<PixelStoreParameter>();
    if (stack->isValid()) {
        GAPID_INFO("glPixelStorei(%u, %d)\n", parameter, value);
        if (glPixelStorei != nullptr) {
            glPixelStorei(parameter, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPixelStorei\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPixelStorei\n");
        return false;
    }
}

bool callGlTexParameteri(Stack* stack, bool pushReturn) {
    int32_t value = stack->pop<int32_t>();
    TextureParameter parameter = stack->pop<TextureParameter>();
    TextureTarget target = stack->pop<TextureTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glTexParameteri(%u, %u, %d)\n", target, parameter, value);
        if (glTexParameteri != nullptr) {
            glTexParameteri(target, parameter, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexParameteri\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexParameteri\n");
        return false;
    }
}

bool callGlTexParameterf(Stack* stack, bool pushReturn) {
    float value = stack->pop<float>();
    TextureParameter parameter = stack->pop<TextureParameter>();
    TextureTarget target = stack->pop<TextureTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glTexParameterf(%u, %u, %f)\n", target, parameter, value);
        if (glTexParameterf != nullptr) {
            glTexParameterf(target, parameter, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexParameterf\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexParameterf\n");
        return false;
    }
}

bool callGlGetTexParameteriv(Stack* stack, bool pushReturn) {
    int32_t* values = stack->pop<int32_t*>();
    TextureParameter parameter = stack->pop<TextureParameter>();
    TextureTarget target = stack->pop<TextureTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glGetTexParameteriv(%u, %u, %p)\n", target, parameter, values);
        if (glGetTexParameteriv != nullptr) {
            glGetTexParameteriv(target, parameter, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetTexParameteriv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetTexParameteriv\n");
        return false;
    }
}

bool callGlGetTexParameterfv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    TextureParameter parameter = stack->pop<TextureParameter>();
    TextureTarget target = stack->pop<TextureTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glGetTexParameterfv(%u, %u, %p)\n", target, parameter, values);
        if (glGetTexParameterfv != nullptr) {
            glGetTexParameterfv(target, parameter, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetTexParameterfv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetTexParameterfv\n");
        return false;
    }
}

bool callGlUniform1i(Stack* stack, bool pushReturn) {
    int32_t value = stack->pop<int32_t>();
    UniformLocation location = stack->pop<UniformLocation>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform1i(%d, %d)\n", location, value);
        if (glUniform1i != nullptr) {
            glUniform1i(location, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform1i\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform1i\n");
        return false;
    }
}

bool callGlUniform2i(Stack* stack, bool pushReturn) {
    int32_t value1 = stack->pop<int32_t>();
    int32_t value0 = stack->pop<int32_t>();
    UniformLocation location = stack->pop<UniformLocation>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform2i(%d, %d, %d)\n", location, value0, value1);
        if (glUniform2i != nullptr) {
            glUniform2i(location, value0, value1);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform2i\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform2i\n");
        return false;
    }
}

bool callGlUniform3i(Stack* stack, bool pushReturn) {
    int32_t value2 = stack->pop<int32_t>();
    int32_t value1 = stack->pop<int32_t>();
    int32_t value0 = stack->pop<int32_t>();
    UniformLocation location = stack->pop<UniformLocation>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform3i(%d, %d, %d, %d)\n", location, value0, value1, value2);
        if (glUniform3i != nullptr) {
            glUniform3i(location, value0, value1, value2);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform3i\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform3i\n");
        return false;
    }
}

bool callGlUniform4i(Stack* stack, bool pushReturn) {
    int32_t value3 = stack->pop<int32_t>();
    int32_t value2 = stack->pop<int32_t>();
    int32_t value1 = stack->pop<int32_t>();
    int32_t value0 = stack->pop<int32_t>();
    UniformLocation location = stack->pop<UniformLocation>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform4i(%d, %d, %d, %d, %d)\n", location, value0, value1, value2, value3);
        if (glUniform4i != nullptr) {
            glUniform4i(location, value0, value1, value2, value3);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform4i\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform4i\n");
        return false;
    }
}

bool callGlUniform1iv(Stack* stack, bool pushReturn) {
    const int32_t* value = stack->pop<const int32_t*>();
    int32_t count = stack->pop<int32_t>();
    UniformLocation location = stack->pop<UniformLocation>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform1iv(%d, %d, %p)\n", location, count, value);
        if (glUniform1iv != nullptr) {
            glUniform1iv(location, count, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform1iv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform1iv\n");
        return false;
    }
}

bool callGlUniform2iv(Stack* stack, bool pushReturn) {
    const int32_t* value = stack->pop<const int32_t*>();
    int32_t count = stack->pop<int32_t>();
    UniformLocation location = stack->pop<UniformLocation>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform2iv(%d, %d, %p)\n", location, count, value);
        if (glUniform2iv != nullptr) {
            glUniform2iv(location, count, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform2iv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform2iv\n");
        return false;
    }
}

bool callGlUniform3iv(Stack* stack, bool pushReturn) {
    const int32_t* value = stack->pop<const int32_t*>();
    int32_t count = stack->pop<int32_t>();
    UniformLocation location = stack->pop<UniformLocation>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform3iv(%d, %d, %p)\n", location, count, value);
        if (glUniform3iv != nullptr) {
            glUniform3iv(location, count, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform3iv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform3iv\n");
        return false;
    }
}

bool callGlUniform4iv(Stack* stack, bool pushReturn) {
    const int32_t* value = stack->pop<const int32_t*>();
    int32_t count = stack->pop<int32_t>();
    UniformLocation location = stack->pop<UniformLocation>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform4iv(%d, %d, %p)\n", location, count, value);
        if (glUniform4iv != nullptr) {
            glUniform4iv(location, count, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform4iv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform4iv\n");
        return false;
    }
}

bool callGlUniform1f(Stack* stack, bool pushReturn) {
    float value = stack->pop<float>();
    UniformLocation location = stack->pop<UniformLocation>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform1f(%d, %f)\n", location, value);
        if (glUniform1f != nullptr) {
            glUniform1f(location, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform1f\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform1f\n");
        return false;
    }
}

bool callGlUniform2f(Stack* stack, bool pushReturn) {
    float value1 = stack->pop<float>();
    float value0 = stack->pop<float>();
    UniformLocation location = stack->pop<UniformLocation>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform2f(%d, %f, %f)\n", location, value0, value1);
        if (glUniform2f != nullptr) {
            glUniform2f(location, value0, value1);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform2f\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform2f\n");
        return false;
    }
}

bool callGlUniform3f(Stack* stack, bool pushReturn) {
    float value2 = stack->pop<float>();
    float value1 = stack->pop<float>();
    float value0 = stack->pop<float>();
    UniformLocation location = stack->pop<UniformLocation>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform3f(%d, %f, %f, %f)\n", location, value0, value1, value2);
        if (glUniform3f != nullptr) {
            glUniform3f(location, value0, value1, value2);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform3f\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform3f\n");
        return false;
    }
}

bool callGlUniform4f(Stack* stack, bool pushReturn) {
    float value3 = stack->pop<float>();
    float value2 = stack->pop<float>();
    float value1 = stack->pop<float>();
    float value0 = stack->pop<float>();
    UniformLocation location = stack->pop<UniformLocation>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform4f(%d, %f, %f, %f, %f)\n", location, value0, value1, value2, value3);
        if (glUniform4f != nullptr) {
            glUniform4f(location, value0, value1, value2, value3);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform4f\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform4f\n");
        return false;
    }
}

bool callGlUniform1fv(Stack* stack, bool pushReturn) {
    const float* value = stack->pop<const float*>();
    int32_t count = stack->pop<int32_t>();
    UniformLocation location = stack->pop<UniformLocation>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform1fv(%d, %d, %p)\n", location, count, value);
        if (glUniform1fv != nullptr) {
            glUniform1fv(location, count, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform1fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform1fv\n");
        return false;
    }
}

bool callGlUniform2fv(Stack* stack, bool pushReturn) {
    const float* value = stack->pop<const float*>();
    int32_t count = stack->pop<int32_t>();
    UniformLocation location = stack->pop<UniformLocation>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform2fv(%d, %d, %p)\n", location, count, value);
        if (glUniform2fv != nullptr) {
            glUniform2fv(location, count, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform2fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform2fv\n");
        return false;
    }
}

bool callGlUniform3fv(Stack* stack, bool pushReturn) {
    const float* value = stack->pop<const float*>();
    int32_t count = stack->pop<int32_t>();
    UniformLocation location = stack->pop<UniformLocation>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform3fv(%d, %d, %p)\n", location, count, value);
        if (glUniform3fv != nullptr) {
            glUniform3fv(location, count, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform3fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform3fv\n");
        return false;
    }
}

bool callGlUniform4fv(Stack* stack, bool pushReturn) {
    const float* value = stack->pop<const float*>();
    int32_t count = stack->pop<int32_t>();
    UniformLocation location = stack->pop<UniformLocation>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform4fv(%d, %d, %p)\n", location, count, value);
        if (glUniform4fv != nullptr) {
            glUniform4fv(location, count, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform4fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform4fv\n");
        return false;
    }
}

bool callGlUniformMatrix2fv(Stack* stack, bool pushReturn) {
    const float* values = stack->pop<const float*>();
    bool transpose = stack->pop<bool>();
    int32_t count = stack->pop<int32_t>();
    UniformLocation location = stack->pop<UniformLocation>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformMatrix2fv(%d, %d, %d, %p)\n", location, count, transpose, values);
        if (glUniformMatrix2fv != nullptr) {
            glUniformMatrix2fv(location, count, transpose, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformMatrix2fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformMatrix2fv\n");
        return false;
    }
}

bool callGlUniformMatrix3fv(Stack* stack, bool pushReturn) {
    const float* values = stack->pop<const float*>();
    bool transpose = stack->pop<bool>();
    int32_t count = stack->pop<int32_t>();
    UniformLocation location = stack->pop<UniformLocation>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformMatrix3fv(%d, %d, %d, %p)\n", location, count, transpose, values);
        if (glUniformMatrix3fv != nullptr) {
            glUniformMatrix3fv(location, count, transpose, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformMatrix3fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformMatrix3fv\n");
        return false;
    }
}

bool callGlUniformMatrix4fv(Stack* stack, bool pushReturn) {
    const float* values = stack->pop<const float*>();
    bool transpose = stack->pop<bool>();
    int32_t count = stack->pop<int32_t>();
    UniformLocation location = stack->pop<UniformLocation>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformMatrix4fv(%d, %d, %d, %p)\n", location, count, transpose, values);
        if (glUniformMatrix4fv != nullptr) {
            glUniformMatrix4fv(location, count, transpose, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformMatrix4fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformMatrix4fv\n");
        return false;
    }
}

bool callGlGetUniformfv(Stack* stack, bool pushReturn) {
    const float* values = stack->pop<const float*>();
    UniformLocation location = stack->pop<UniformLocation>();
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        GAPID_INFO("glGetUniformfv(%u, %d, %p)\n", program, location, values);
        if (glGetUniformfv != nullptr) {
            glGetUniformfv(program, location, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetUniformfv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetUniformfv\n");
        return false;
    }
}

bool callGlGetUniformiv(Stack* stack, bool pushReturn) {
    const int32_t* values = stack->pop<const int32_t*>();
    UniformLocation location = stack->pop<UniformLocation>();
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        GAPID_INFO("glGetUniformiv(%u, %d, %p)\n", program, location, values);
        if (glGetUniformiv != nullptr) {
            glGetUniformiv(program, location, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetUniformiv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetUniformiv\n");
        return false;
    }
}

bool callGlVertexAttrib1f(Stack* stack, bool pushReturn) {
    float value0 = stack->pop<float>();
    AttributeLocation location = stack->pop<AttributeLocation>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttrib1f(%u, %f)\n", location, value0);
        if (glVertexAttrib1f != nullptr) {
            glVertexAttrib1f(location, value0);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttrib1f\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttrib1f\n");
        return false;
    }
}

bool callGlVertexAttrib2f(Stack* stack, bool pushReturn) {
    float value1 = stack->pop<float>();
    float value0 = stack->pop<float>();
    AttributeLocation location = stack->pop<AttributeLocation>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttrib2f(%u, %f, %f)\n", location, value0, value1);
        if (glVertexAttrib2f != nullptr) {
            glVertexAttrib2f(location, value0, value1);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttrib2f\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttrib2f\n");
        return false;
    }
}

bool callGlVertexAttrib3f(Stack* stack, bool pushReturn) {
    float value2 = stack->pop<float>();
    float value1 = stack->pop<float>();
    float value0 = stack->pop<float>();
    AttributeLocation location = stack->pop<AttributeLocation>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttrib3f(%u, %f, %f, %f)\n", location, value0, value1, value2);
        if (glVertexAttrib3f != nullptr) {
            glVertexAttrib3f(location, value0, value1, value2);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttrib3f\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttrib3f\n");
        return false;
    }
}

bool callGlVertexAttrib4f(Stack* stack, bool pushReturn) {
    float value3 = stack->pop<float>();
    float value2 = stack->pop<float>();
    float value1 = stack->pop<float>();
    float value0 = stack->pop<float>();
    AttributeLocation location = stack->pop<AttributeLocation>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttrib4f(%u, %f, %f, %f, %f)\n", location, value0, value1, value2,
                   value3);
        if (glVertexAttrib4f != nullptr) {
            glVertexAttrib4f(location, value0, value1, value2, value3);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttrib4f\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttrib4f\n");
        return false;
    }
}

bool callGlVertexAttrib1fv(Stack* stack, bool pushReturn) {
    const float* value = stack->pop<const float*>();
    AttributeLocation location = stack->pop<AttributeLocation>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttrib1fv(%u, %p)\n", location, value);
        if (glVertexAttrib1fv != nullptr) {
            glVertexAttrib1fv(location, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttrib1fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttrib1fv\n");
        return false;
    }
}

bool callGlVertexAttrib2fv(Stack* stack, bool pushReturn) {
    const float* value = stack->pop<const float*>();
    AttributeLocation location = stack->pop<AttributeLocation>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttrib2fv(%u, %p)\n", location, value);
        if (glVertexAttrib2fv != nullptr) {
            glVertexAttrib2fv(location, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttrib2fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttrib2fv\n");
        return false;
    }
}

bool callGlVertexAttrib3fv(Stack* stack, bool pushReturn) {
    const float* value = stack->pop<const float*>();
    AttributeLocation location = stack->pop<AttributeLocation>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttrib3fv(%u, %p)\n", location, value);
        if (glVertexAttrib3fv != nullptr) {
            glVertexAttrib3fv(location, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttrib3fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttrib3fv\n");
        return false;
    }
}

bool callGlVertexAttrib4fv(Stack* stack, bool pushReturn) {
    const float* value = stack->pop<const float*>();
    AttributeLocation location = stack->pop<AttributeLocation>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttrib4fv(%u, %p)\n", location, value);
        if (glVertexAttrib4fv != nullptr) {
            glVertexAttrib4fv(location, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttrib4fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttrib4fv\n");
        return false;
    }
}

bool callGlGetShaderPrecisionFormat(Stack* stack, bool pushReturn) {
    int32_t* precision = stack->pop<int32_t*>();
    int32_t* range = stack->pop<int32_t*>();
    PrecisionType precision_type = stack->pop<PrecisionType>();
    ShaderType shader_type = stack->pop<ShaderType>();
    if (stack->isValid()) {
        GAPID_INFO("glGetShaderPrecisionFormat(%u, %u, %p, %p)\n", shader_type, precision_type,
                   range, precision);
        if (glGetShaderPrecisionFormat != nullptr) {
            glGetShaderPrecisionFormat(shader_type, precision_type, range, precision);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetShaderPrecisionFormat\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetShaderPrecisionFormat\n");
        return false;
    }
}

bool callGlDepthMask(Stack* stack, bool pushReturn) {
    bool enabled = stack->pop<bool>();
    if (stack->isValid()) {
        GAPID_INFO("glDepthMask(%d)\n", enabled);
        if (glDepthMask != nullptr) {
            glDepthMask(enabled);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDepthMask\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDepthMask\n");
        return false;
    }
}

bool callGlDepthFunc(Stack* stack, bool pushReturn) {
    TestFunction function = stack->pop<TestFunction>();
    if (stack->isValid()) {
        GAPID_INFO("glDepthFunc(%u)\n", function);
        if (glDepthFunc != nullptr) {
            glDepthFunc(function);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDepthFunc\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDepthFunc\n");
        return false;
    }
}

bool callGlDepthRangef(Stack* stack, bool pushReturn) {
    float far = stack->pop<float>();
    float near = stack->pop<float>();
    if (stack->isValid()) {
        GAPID_INFO("glDepthRangef(%f, %f)\n", near, far);
        if (glDepthRangef != nullptr) {
            glDepthRangef(near, far);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDepthRangef\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDepthRangef\n");
        return false;
    }
}

bool callGlColorMask(Stack* stack, bool pushReturn) {
    bool alpha = stack->pop<bool>();
    bool blue = stack->pop<bool>();
    bool green = stack->pop<bool>();
    bool red = stack->pop<bool>();
    if (stack->isValid()) {
        GAPID_INFO("glColorMask(%d, %d, %d, %d)\n", red, green, blue, alpha);
        if (glColorMask != nullptr) {
            glColorMask(red, green, blue, alpha);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glColorMask\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glColorMask\n");
        return false;
    }
}

bool callGlStencilMask(Stack* stack, bool pushReturn) {
    uint32_t mask = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glStencilMask(%u)\n", mask);
        if (glStencilMask != nullptr) {
            glStencilMask(mask);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glStencilMask\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStencilMask\n");
        return false;
    }
}

bool callGlStencilMaskSeparate(Stack* stack, bool pushReturn) {
    uint32_t mask = stack->pop<uint32_t>();
    FaceMode face = stack->pop<FaceMode>();
    if (stack->isValid()) {
        GAPID_INFO("glStencilMaskSeparate(%u, %u)\n", face, mask);
        if (glStencilMaskSeparate != nullptr) {
            glStencilMaskSeparate(face, mask);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glStencilMaskSeparate\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStencilMaskSeparate\n");
        return false;
    }
}

bool callGlStencilFuncSeparate(Stack* stack, bool pushReturn) {
    int32_t mask = stack->pop<int32_t>();
    int32_t reference_value = stack->pop<int32_t>();
    TestFunction function = stack->pop<TestFunction>();
    FaceMode face = stack->pop<FaceMode>();
    if (stack->isValid()) {
        GAPID_INFO("glStencilFuncSeparate(%u, %u, %d, %d)\n", face, function, reference_value,
                   mask);
        if (glStencilFuncSeparate != nullptr) {
            glStencilFuncSeparate(face, function, reference_value, mask);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glStencilFuncSeparate\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStencilFuncSeparate\n");
        return false;
    }
}

bool callGlStencilOpSeparate(Stack* stack, bool pushReturn) {
    StencilAction stencil_pass_depth_pass = stack->pop<StencilAction>();
    StencilAction stencil_pass_depth_fail = stack->pop<StencilAction>();
    StencilAction stencil_fail = stack->pop<StencilAction>();
    FaceMode face = stack->pop<FaceMode>();
    if (stack->isValid()) {
        GAPID_INFO("glStencilOpSeparate(%u, %u, %u, %u)\n", face, stencil_fail,
                   stencil_pass_depth_fail, stencil_pass_depth_pass);
        if (glStencilOpSeparate != nullptr) {
            glStencilOpSeparate(face, stencil_fail, stencil_pass_depth_fail,
                                stencil_pass_depth_pass);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glStencilOpSeparate\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStencilOpSeparate\n");
        return false;
    }
}

bool callGlFrontFace(Stack* stack, bool pushReturn) {
    FaceOrientation orientation = stack->pop<FaceOrientation>();
    if (stack->isValid()) {
        GAPID_INFO("glFrontFace(%u)\n", orientation);
        if (glFrontFace != nullptr) {
            glFrontFace(orientation);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFrontFace\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFrontFace\n");
        return false;
    }
}

bool callGlViewport(Stack* stack, bool pushReturn) {
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    int32_t y = stack->pop<int32_t>();
    int32_t x = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glViewport(%d, %d, %d, %d)\n", x, y, width, height);
        if (glViewport != nullptr) {
            glViewport(x, y, width, height);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glViewport\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glViewport\n");
        return false;
    }
}

bool callGlScissor(Stack* stack, bool pushReturn) {
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    int32_t y = stack->pop<int32_t>();
    int32_t x = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glScissor(%d, %d, %d, %d)\n", x, y, width, height);
        if (glScissor != nullptr) {
            glScissor(x, y, width, height);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glScissor\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glScissor\n");
        return false;
    }
}

bool callGlActiveTexture(Stack* stack, bool pushReturn) {
    TextureUnit unit = stack->pop<TextureUnit>();
    if (stack->isValid()) {
        GAPID_INFO("glActiveTexture(%u)\n", unit);
        if (glActiveTexture != nullptr) {
            glActiveTexture(unit);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glActiveTexture\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glActiveTexture\n");
        return false;
    }
}

bool callGlGenTextures(Stack* stack, bool pushReturn) {
    TextureId* textures = stack->pop<TextureId*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenTextures(%d, %p)\n", count, textures);
        if (glGenTextures != nullptr) {
            glGenTextures(count, textures);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenTextures\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenTextures\n");
        return false;
    }
}

bool callGlDeleteTextures(Stack* stack, bool pushReturn) {
    const TextureId* textures = stack->pop<const TextureId*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteTextures(%d, %p)\n", count, textures);
        if (glDeleteTextures != nullptr) {
            glDeleteTextures(count, textures);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteTextures\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteTextures\n");
        return false;
    }
}

bool callGlIsTexture(Stack* stack, bool pushReturn) {
    TextureId texture = stack->pop<TextureId>();
    if (stack->isValid()) {
        GAPID_INFO("glIsTexture(%u)\n", texture);
        if (glIsTexture != nullptr) {
            bool return_value = glIsTexture(texture);
            GAPID_INFO("Returned: %d\n", return_value);
            if (pushReturn) {
                stack->push<bool>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsTexture\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsTexture\n");
        return false;
    }
}

bool callGlBindTexture(Stack* stack, bool pushReturn) {
    TextureId texture = stack->pop<TextureId>();
    TextureTarget target = stack->pop<TextureTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glBindTexture(%u, %u)\n", target, texture);
        if (glBindTexture != nullptr) {
            glBindTexture(target, texture);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindTexture\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindTexture\n");
        return false;
    }
}

bool callGlTexImage2D(Stack* stack, bool pushReturn) {
    TexturePointer data = stack->pop<TexturePointer>();
    TexelType type = stack->pop<TexelType>();
    TexelFormat format = stack->pop<TexelFormat>();
    int32_t border = stack->pop<int32_t>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    TexelFormat internal_format = stack->pop<TexelFormat>();
    int32_t level = stack->pop<int32_t>();
    TextureImageTarget target = stack->pop<TextureImageTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glTexImage2D(%u, %d, %u, %d, %d, %d, %u, %u, %p)\n", target, level,
                   internal_format, width, height, border, format, type, data);
        if (glTexImage2D != nullptr) {
            glTexImage2D(target, level, internal_format, width, height, border, format, type, data);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexImage2D\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexImage2D\n");
        return false;
    }
}

bool callGlTexSubImage2D(Stack* stack, bool pushReturn) {
    TexturePointer data = stack->pop<TexturePointer>();
    TexelType type = stack->pop<TexelType>();
    TexelFormat format = stack->pop<TexelFormat>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    int32_t yoffset = stack->pop<int32_t>();
    int32_t xoffset = stack->pop<int32_t>();
    int32_t level = stack->pop<int32_t>();
    TextureImageTarget target = stack->pop<TextureImageTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glTexSubImage2D(%u, %d, %d, %d, %d, %d, %u, %u, %p)\n", target, level, xoffset,
                   yoffset, width, height, format, type, data);
        if (glTexSubImage2D != nullptr) {
            glTexSubImage2D(target, level, xoffset, yoffset, width, height, format, type, data);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexSubImage2D\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexSubImage2D\n");
        return false;
    }
}

bool callGlCopyTexImage2D(Stack* stack, bool pushReturn) {
    int32_t border = stack->pop<int32_t>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    int32_t y = stack->pop<int32_t>();
    int32_t x = stack->pop<int32_t>();
    TexelFormat format = stack->pop<TexelFormat>();
    int32_t level = stack->pop<int32_t>();
    TextureImageTarget target = stack->pop<TextureImageTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glCopyTexImage2D(%u, %d, %u, %d, %d, %d, %d, %d)\n", target, level, format, x,
                   y, width, height, border);
        if (glCopyTexImage2D != nullptr) {
            glCopyTexImage2D(target, level, format, x, y, width, height, border);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCopyTexImage2D\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCopyTexImage2D\n");
        return false;
    }
}

bool callGlCopyTexSubImage2D(Stack* stack, bool pushReturn) {
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    int32_t y = stack->pop<int32_t>();
    int32_t x = stack->pop<int32_t>();
    int32_t yoffset = stack->pop<int32_t>();
    int32_t xoffset = stack->pop<int32_t>();
    int32_t level = stack->pop<int32_t>();
    TextureImageTarget target = stack->pop<TextureImageTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glCopyTexSubImage2D(%u, %d, %d, %d, %d, %d, %d, %d)\n", target, level, xoffset,
                   yoffset, x, y, width, height);
        if (glCopyTexSubImage2D != nullptr) {
            glCopyTexSubImage2D(target, level, xoffset, yoffset, x, y, width, height);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCopyTexSubImage2D\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCopyTexSubImage2D\n");
        return false;
    }
}

bool callGlCompressedTexImage2D(Stack* stack, bool pushReturn) {
    TexturePointer data = stack->pop<TexturePointer>();
    int32_t image_size = stack->pop<int32_t>();
    int32_t border = stack->pop<int32_t>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    CompressedTexelFormat format = stack->pop<CompressedTexelFormat>();
    int32_t level = stack->pop<int32_t>();
    TextureImageTarget target = stack->pop<TextureImageTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glCompressedTexImage2D(%u, %d, %u, %d, %d, %d, %d, %p)\n", target, level,
                   format, width, height, border, image_size, data);
        if (glCompressedTexImage2D != nullptr) {
            glCompressedTexImage2D(target, level, format, width, height, border, image_size, data);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCompressedTexImage2D\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCompressedTexImage2D\n");
        return false;
    }
}

bool callGlCompressedTexSubImage2D(Stack* stack, bool pushReturn) {
    TexturePointer data = stack->pop<TexturePointer>();
    int32_t image_size = stack->pop<int32_t>();
    CompressedTexelFormat format = stack->pop<CompressedTexelFormat>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    int32_t yoffset = stack->pop<int32_t>();
    int32_t xoffset = stack->pop<int32_t>();
    int32_t level = stack->pop<int32_t>();
    TextureImageTarget target = stack->pop<TextureImageTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glCompressedTexSubImage2D(%u, %d, %d, %d, %d, %d, %u, %d, %p)\n", target, level,
                   xoffset, yoffset, width, height, format, image_size, data);
        if (glCompressedTexSubImage2D != nullptr) {
            glCompressedTexSubImage2D(target, level, xoffset, yoffset, width, height, format,
                                      image_size, data);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCompressedTexSubImage2D\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCompressedTexSubImage2D\n");
        return false;
    }
}

bool callGlGenerateMipmap(Stack* stack, bool pushReturn) {
    TextureImageTarget target = stack->pop<TextureImageTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glGenerateMipmap(%u)\n", target);
        if (glGenerateMipmap != nullptr) {
            glGenerateMipmap(target);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenerateMipmap\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenerateMipmap\n");
        return false;
    }
}

bool callGlReadPixels(Stack* stack, bool pushReturn) {
    void* data = stack->pop<void*>();
    TexelType type = stack->pop<TexelType>();
    BaseTexelFormat format = stack->pop<BaseTexelFormat>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    int32_t y = stack->pop<int32_t>();
    int32_t x = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glReadPixels(%d, %d, %d, %d, %u, %u, %p)\n", x, y, width, height, format, type,
                   data);
        if (glReadPixels != nullptr) {
            glReadPixels(x, y, width, height, format, type, data);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glReadPixels\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glReadPixels\n");
        return false;
    }
}

bool callGlGenFramebuffers(Stack* stack, bool pushReturn) {
    FramebufferId* framebuffers = stack->pop<FramebufferId*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenFramebuffers(%d, %p)\n", count, framebuffers);
        if (glGenFramebuffers != nullptr) {
            glGenFramebuffers(count, framebuffers);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenFramebuffers\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenFramebuffers\n");
        return false;
    }
}

bool callGlBindFramebuffer(Stack* stack, bool pushReturn) {
    FramebufferId framebuffer = stack->pop<FramebufferId>();
    FramebufferTarget target = stack->pop<FramebufferTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glBindFramebuffer(%u, %u)\n", target, framebuffer);
        if (glBindFramebuffer != nullptr) {
            glBindFramebuffer(target, framebuffer);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindFramebuffer\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindFramebuffer\n");
        return false;
    }
}

bool callGlCheckFramebufferStatus(Stack* stack, bool pushReturn) {
    FramebufferTarget target = stack->pop<FramebufferTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glCheckFramebufferStatus(%u)\n", target);
        if (glCheckFramebufferStatus != nullptr) {
            FramebufferStatus return_value = glCheckFramebufferStatus(target);
            GAPID_INFO("Returned: %u\n", return_value);
            if (pushReturn) {
                stack->push<FramebufferStatus>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCheckFramebufferStatus\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCheckFramebufferStatus\n");
        return false;
    }
}

bool callGlDeleteFramebuffers(Stack* stack, bool pushReturn) {
    const FramebufferId* framebuffers = stack->pop<const FramebufferId*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteFramebuffers(%d, %p)\n", count, framebuffers);
        if (glDeleteFramebuffers != nullptr) {
            glDeleteFramebuffers(count, framebuffers);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteFramebuffers\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteFramebuffers\n");
        return false;
    }
}

bool callGlIsFramebuffer(Stack* stack, bool pushReturn) {
    FramebufferId framebuffer = stack->pop<FramebufferId>();
    if (stack->isValid()) {
        GAPID_INFO("glIsFramebuffer(%u)\n", framebuffer);
        if (glIsFramebuffer != nullptr) {
            bool return_value = glIsFramebuffer(framebuffer);
            GAPID_INFO("Returned: %d\n", return_value);
            if (pushReturn) {
                stack->push<bool>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsFramebuffer\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsFramebuffer\n");
        return false;
    }
}

bool callGlGenRenderbuffers(Stack* stack, bool pushReturn) {
    RenderbufferId* renderbuffers = stack->pop<RenderbufferId*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenRenderbuffers(%d, %p)\n", count, renderbuffers);
        if (glGenRenderbuffers != nullptr) {
            glGenRenderbuffers(count, renderbuffers);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenRenderbuffers\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenRenderbuffers\n");
        return false;
    }
}

bool callGlBindRenderbuffer(Stack* stack, bool pushReturn) {
    RenderbufferId renderbuffer = stack->pop<RenderbufferId>();
    RenderbufferTarget target = stack->pop<RenderbufferTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glBindRenderbuffer(%u, %u)\n", target, renderbuffer);
        if (glBindRenderbuffer != nullptr) {
            glBindRenderbuffer(target, renderbuffer);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindRenderbuffer\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindRenderbuffer\n");
        return false;
    }
}

bool callGlRenderbufferStorage(Stack* stack, bool pushReturn) {
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    RenderbufferFormat format = stack->pop<RenderbufferFormat>();
    RenderbufferTarget target = stack->pop<RenderbufferTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glRenderbufferStorage(%u, %u, %d, %d)\n", target, format, width, height);
        if (glRenderbufferStorage != nullptr) {
            glRenderbufferStorage(target, format, width, height);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glRenderbufferStorage\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glRenderbufferStorage\n");
        return false;
    }
}

bool callGlDeleteRenderbuffers(Stack* stack, bool pushReturn) {
    const RenderbufferId* renderbuffers = stack->pop<const RenderbufferId*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteRenderbuffers(%d, %p)\n", count, renderbuffers);
        if (glDeleteRenderbuffers != nullptr) {
            glDeleteRenderbuffers(count, renderbuffers);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteRenderbuffers\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteRenderbuffers\n");
        return false;
    }
}

bool callGlIsRenderbuffer(Stack* stack, bool pushReturn) {
    RenderbufferId renderbuffer = stack->pop<RenderbufferId>();
    if (stack->isValid()) {
        GAPID_INFO("glIsRenderbuffer(%u)\n", renderbuffer);
        if (glIsRenderbuffer != nullptr) {
            bool return_value = glIsRenderbuffer(renderbuffer);
            GAPID_INFO("Returned: %d\n", return_value);
            if (pushReturn) {
                stack->push<bool>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsRenderbuffer\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsRenderbuffer\n");
        return false;
    }
}

bool callGlGetRenderbufferParameteriv(Stack* stack, bool pushReturn) {
    int32_t* values = stack->pop<int32_t*>();
    RenderbufferParameter parameter = stack->pop<RenderbufferParameter>();
    RenderbufferTarget target = stack->pop<RenderbufferTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glGetRenderbufferParameteriv(%u, %u, %p)\n", target, parameter, values);
        if (glGetRenderbufferParameteriv != nullptr) {
            glGetRenderbufferParameteriv(target, parameter, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetRenderbufferParameteriv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetRenderbufferParameteriv\n");
        return false;
    }
}

bool callGlGenBuffers(Stack* stack, bool pushReturn) {
    BufferId* buffers = stack->pop<BufferId*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenBuffers(%d, %p)\n", count, buffers);
        if (glGenBuffers != nullptr) {
            glGenBuffers(count, buffers);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenBuffers\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenBuffers\n");
        return false;
    }
}

bool callGlBindBuffer(Stack* stack, bool pushReturn) {
    BufferId buffer = stack->pop<BufferId>();
    BufferTarget target = stack->pop<BufferTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glBindBuffer(%u, %u)\n", target, buffer);
        if (glBindBuffer != nullptr) {
            glBindBuffer(target, buffer);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindBuffer\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindBuffer\n");
        return false;
    }
}

bool callGlBufferData(Stack* stack, bool pushReturn) {
    BufferUsage usage = stack->pop<BufferUsage>();
    BufferDataPointer data = stack->pop<BufferDataPointer>();
    int32_t size = stack->pop<int32_t>();
    BufferTarget target = stack->pop<BufferTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glBufferData(%u, %d, %p, %u)\n", target, size, data, usage);
        if (glBufferData != nullptr) {
            glBufferData(target, size, data, usage);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBufferData\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBufferData\n");
        return false;
    }
}

bool callGlBufferSubData(Stack* stack, bool pushReturn) {
    const void* data = stack->pop<const void*>();
    int32_t size = stack->pop<int32_t>();
    int32_t offset = stack->pop<int32_t>();
    BufferTarget target = stack->pop<BufferTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glBufferSubData(%u, %d, %d, %p)\n", target, offset, size, data);
        if (glBufferSubData != nullptr) {
            glBufferSubData(target, offset, size, data);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBufferSubData\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBufferSubData\n");
        return false;
    }
}

bool callGlDeleteBuffers(Stack* stack, bool pushReturn) {
    const BufferId* buffers = stack->pop<const BufferId*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteBuffers(%d, %p)\n", count, buffers);
        if (glDeleteBuffers != nullptr) {
            glDeleteBuffers(count, buffers);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteBuffers\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteBuffers\n");
        return false;
    }
}

bool callGlIsBuffer(Stack* stack, bool pushReturn) {
    BufferId buffer = stack->pop<BufferId>();
    if (stack->isValid()) {
        GAPID_INFO("glIsBuffer(%u)\n", buffer);
        if (glIsBuffer != nullptr) {
            bool return_value = glIsBuffer(buffer);
            GAPID_INFO("Returned: %d\n", return_value);
            if (pushReturn) {
                stack->push<bool>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsBuffer\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsBuffer\n");
        return false;
    }
}

bool callGlGetBufferParameteriv(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    BufferParameter parameter = stack->pop<BufferParameter>();
    BufferTarget target = stack->pop<BufferTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glGetBufferParameteriv(%u, %u, %p)\n", target, parameter, value);
        if (glGetBufferParameteriv != nullptr) {
            glGetBufferParameteriv(target, parameter, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetBufferParameteriv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetBufferParameteriv\n");
        return false;
    }
}

bool callGlCreateShader(Stack* stack, bool pushReturn) {
    ShaderType type = stack->pop<ShaderType>();
    if (stack->isValid()) {
        GAPID_INFO("glCreateShader(%u)\n", type);
        if (glCreateShader != nullptr) {
            ShaderId return_value = glCreateShader(type);
            GAPID_INFO("Returned: %u\n", return_value);
            if (pushReturn) {
                stack->push<ShaderId>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCreateShader\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCreateShader\n");
        return false;
    }
}

bool callGlDeleteShader(Stack* stack, bool pushReturn) {
    ShaderId shader = stack->pop<ShaderId>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteShader(%u)\n", shader);
        if (glDeleteShader != nullptr) {
            glDeleteShader(shader);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteShader\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteShader\n");
        return false;
    }
}

bool callGlShaderSource(Stack* stack, bool pushReturn) {
    const int32_t* length = stack->pop<const int32_t*>();
    const char** source = stack->pop<const char**>();
    int32_t count = stack->pop<int32_t>();
    ShaderId shader = stack->pop<ShaderId>();
    if (stack->isValid()) {
        GAPID_INFO("glShaderSource(%u, %d, %p, %p)\n", shader, count, source, length);
        if (glShaderSource != nullptr) {
            glShaderSource(shader, count, source, length);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glShaderSource\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glShaderSource\n");
        return false;
    }
}

bool callGlShaderBinary(Stack* stack, bool pushReturn) {
    int32_t binary_size = stack->pop<int32_t>();
    const void* binary = stack->pop<const void*>();
    uint32_t binary_format = stack->pop<uint32_t>();
    const ShaderId* shaders = stack->pop<const ShaderId*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glShaderBinary(%d, %p, %u, %p, %d)\n", count, shaders, binary_format, binary,
                   binary_size);
        if (glShaderBinary != nullptr) {
            glShaderBinary(count, shaders, binary_format, binary, binary_size);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glShaderBinary\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glShaderBinary\n");
        return false;
    }
}

bool callGlGetShaderInfoLog(Stack* stack, bool pushReturn) {
    char* info = stack->pop<char*>();
    int32_t* string_length_written = stack->pop<int32_t*>();
    int32_t buffer_length = stack->pop<int32_t>();
    ShaderId shader = stack->pop<ShaderId>();
    if (stack->isValid()) {
        GAPID_INFO("glGetShaderInfoLog(%u, %d, %p, %p)\n", shader, buffer_length,
                   string_length_written, info);
        if (glGetShaderInfoLog != nullptr) {
            glGetShaderInfoLog(shader, buffer_length, string_length_written, info);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetShaderInfoLog\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetShaderInfoLog\n");
        return false;
    }
}

bool callGlGetShaderSource(Stack* stack, bool pushReturn) {
    char* source = stack->pop<char*>();
    int32_t* string_length_written = stack->pop<int32_t*>();
    int32_t buffer_length = stack->pop<int32_t>();
    ShaderId shader = stack->pop<ShaderId>();
    if (stack->isValid()) {
        GAPID_INFO("glGetShaderSource(%u, %d, %p, %p)\n", shader, buffer_length,
                   string_length_written, source);
        if (glGetShaderSource != nullptr) {
            glGetShaderSource(shader, buffer_length, string_length_written, source);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetShaderSource\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetShaderSource\n");
        return false;
    }
}

bool callGlReleaseShaderCompiler(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glReleaseShaderCompiler()\n");
        if (glReleaseShaderCompiler != nullptr) {
            glReleaseShaderCompiler();
        } else {
            GAPID_WARNING("Attempted to call unsupported function glReleaseShaderCompiler\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glReleaseShaderCompiler\n");
        return false;
    }
}

bool callGlCompileShader(Stack* stack, bool pushReturn) {
    ShaderId shader = stack->pop<ShaderId>();
    if (stack->isValid()) {
        GAPID_INFO("glCompileShader(%u)\n", shader);
        if (glCompileShader != nullptr) {
            glCompileShader(shader);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCompileShader\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCompileShader\n");
        return false;
    }
}

bool callGlIsShader(Stack* stack, bool pushReturn) {
    ShaderId shader = stack->pop<ShaderId>();
    if (stack->isValid()) {
        GAPID_INFO("glIsShader(%u)\n", shader);
        if (glIsShader != nullptr) {
            bool return_value = glIsShader(shader);
            GAPID_INFO("Returned: %d\n", return_value);
            if (pushReturn) {
                stack->push<bool>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsShader\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsShader\n");
        return false;
    }
}

bool callGlCreateProgram(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glCreateProgram()\n");
        if (glCreateProgram != nullptr) {
            ProgramId return_value = glCreateProgram();
            GAPID_INFO("Returned: %u\n", return_value);
            if (pushReturn) {
                stack->push<ProgramId>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCreateProgram\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCreateProgram\n");
        return false;
    }
}

bool callGlDeleteProgram(Stack* stack, bool pushReturn) {
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteProgram(%u)\n", program);
        if (glDeleteProgram != nullptr) {
            glDeleteProgram(program);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteProgram\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteProgram\n");
        return false;
    }
}

bool callGlAttachShader(Stack* stack, bool pushReturn) {
    ShaderId shader = stack->pop<ShaderId>();
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        GAPID_INFO("glAttachShader(%u, %u)\n", program, shader);
        if (glAttachShader != nullptr) {
            glAttachShader(program, shader);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glAttachShader\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glAttachShader\n");
        return false;
    }
}

bool callGlDetachShader(Stack* stack, bool pushReturn) {
    ShaderId shader = stack->pop<ShaderId>();
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        GAPID_INFO("glDetachShader(%u, %u)\n", program, shader);
        if (glDetachShader != nullptr) {
            glDetachShader(program, shader);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDetachShader\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDetachShader\n");
        return false;
    }
}

bool callGlGetAttachedShaders(Stack* stack, bool pushReturn) {
    ShaderId* shaders = stack->pop<ShaderId*>();
    int32_t* shaders_length_written = stack->pop<int32_t*>();
    int32_t buffer_length = stack->pop<int32_t>();
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        GAPID_INFO("glGetAttachedShaders(%u, %d, %p, %p)\n", program, buffer_length,
                   shaders_length_written, shaders);
        if (glGetAttachedShaders != nullptr) {
            glGetAttachedShaders(program, buffer_length, shaders_length_written, shaders);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetAttachedShaders\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetAttachedShaders\n");
        return false;
    }
}

bool callGlLinkProgram(Stack* stack, bool pushReturn) {
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        GAPID_INFO("glLinkProgram(%u)\n", program);
        if (glLinkProgram != nullptr) {
            glLinkProgram(program);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glLinkProgram\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glLinkProgram\n");
        return false;
    }
}

bool callGlGetProgramInfoLog(Stack* stack, bool pushReturn) {
    char* info = stack->pop<char*>();
    int32_t* string_length_written = stack->pop<int32_t*>();
    int32_t buffer_length = stack->pop<int32_t>();
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        GAPID_INFO("glGetProgramInfoLog(%u, %d, %p, %p)\n", program, buffer_length,
                   string_length_written, info);
        if (glGetProgramInfoLog != nullptr) {
            glGetProgramInfoLog(program, buffer_length, string_length_written, info);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetProgramInfoLog\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetProgramInfoLog\n");
        return false;
    }
}

bool callGlUseProgram(Stack* stack, bool pushReturn) {
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        GAPID_INFO("glUseProgram(%u)\n", program);
        if (glUseProgram != nullptr) {
            glUseProgram(program);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUseProgram\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUseProgram\n");
        return false;
    }
}

bool callGlIsProgram(Stack* stack, bool pushReturn) {
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        GAPID_INFO("glIsProgram(%u)\n", program);
        if (glIsProgram != nullptr) {
            bool return_value = glIsProgram(program);
            GAPID_INFO("Returned: %d\n", return_value);
            if (pushReturn) {
                stack->push<bool>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsProgram\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsProgram\n");
        return false;
    }
}

bool callGlValidateProgram(Stack* stack, bool pushReturn) {
    ProgramId program = stack->pop<ProgramId>();
    if (stack->isValid()) {
        GAPID_INFO("glValidateProgram(%u)\n", program);
        if (glValidateProgram != nullptr) {
            glValidateProgram(program);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glValidateProgram\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glValidateProgram\n");
        return false;
    }
}

bool callGlClearColor(Stack* stack, bool pushReturn) {
    float a = stack->pop<float>();
    float b = stack->pop<float>();
    float g = stack->pop<float>();
    float r = stack->pop<float>();
    if (stack->isValid()) {
        GAPID_INFO("glClearColor(%f, %f, %f, %f)\n", r, g, b, a);
        if (glClearColor != nullptr) {
            glClearColor(r, g, b, a);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glClearColor\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glClearColor\n");
        return false;
    }
}

bool callGlClearDepthf(Stack* stack, bool pushReturn) {
    float depth = stack->pop<float>();
    if (stack->isValid()) {
        GAPID_INFO("glClearDepthf(%f)\n", depth);
        if (glClearDepthf != nullptr) {
            glClearDepthf(depth);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glClearDepthf\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glClearDepthf\n");
        return false;
    }
}

bool callGlClearStencil(Stack* stack, bool pushReturn) {
    int32_t stencil = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glClearStencil(%d)\n", stencil);
        if (glClearStencil != nullptr) {
            glClearStencil(stencil);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glClearStencil\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glClearStencil\n");
        return false;
    }
}

bool callGlClear(Stack* stack, bool pushReturn) {
    ClearMask mask = stack->pop<ClearMask>();
    if (stack->isValid()) {
        GAPID_INFO("glClear(%u)\n", mask);
        if (glClear != nullptr) {
            glClear(mask);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glClear\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glClear\n");
        return false;
    }
}

bool callGlCullFace(Stack* stack, bool pushReturn) {
    FaceMode mode = stack->pop<FaceMode>();
    if (stack->isValid()) {
        GAPID_INFO("glCullFace(%u)\n", mode);
        if (glCullFace != nullptr) {
            glCullFace(mode);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCullFace\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCullFace\n");
        return false;
    }
}

bool callGlPolygonOffset(Stack* stack, bool pushReturn) {
    float units = stack->pop<float>();
    float scale_factor = stack->pop<float>();
    if (stack->isValid()) {
        GAPID_INFO("glPolygonOffset(%f, %f)\n", scale_factor, units);
        if (glPolygonOffset != nullptr) {
            glPolygonOffset(scale_factor, units);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPolygonOffset\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPolygonOffset\n");
        return false;
    }
}

bool callGlLineWidth(Stack* stack, bool pushReturn) {
    float width = stack->pop<float>();
    if (stack->isValid()) {
        GAPID_INFO("glLineWidth(%f)\n", width);
        if (glLineWidth != nullptr) {
            glLineWidth(width);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glLineWidth\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glLineWidth\n");
        return false;
    }
}

bool callGlSampleCoverage(Stack* stack, bool pushReturn) {
    bool invert = stack->pop<bool>();
    float value = stack->pop<float>();
    if (stack->isValid()) {
        GAPID_INFO("glSampleCoverage(%f, %d)\n", value, invert);
        if (glSampleCoverage != nullptr) {
            glSampleCoverage(value, invert);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glSampleCoverage\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glSampleCoverage\n");
        return false;
    }
}

bool callGlHint(Stack* stack, bool pushReturn) {
    HintMode mode = stack->pop<HintMode>();
    HintTarget target = stack->pop<HintTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glHint(%u, %u)\n", target, mode);
        if (glHint != nullptr) {
            glHint(target, mode);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glHint\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glHint\n");
        return false;
    }
}

bool callGlFramebufferRenderbuffer(Stack* stack, bool pushReturn) {
    RenderbufferId renderbuffer = stack->pop<RenderbufferId>();
    RenderbufferTarget renderbuffer_target = stack->pop<RenderbufferTarget>();
    FramebufferAttachment framebuffer_attachment = stack->pop<FramebufferAttachment>();
    FramebufferTarget framebuffer_target = stack->pop<FramebufferTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glFramebufferRenderbuffer(%u, %u, %u, %u)\n", framebuffer_target,
                   framebuffer_attachment, renderbuffer_target, renderbuffer);
        if (glFramebufferRenderbuffer != nullptr) {
            glFramebufferRenderbuffer(framebuffer_target, framebuffer_attachment,
                                      renderbuffer_target, renderbuffer);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFramebufferRenderbuffer\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFramebufferRenderbuffer\n");
        return false;
    }
}

bool callGlFramebufferTexture2D(Stack* stack, bool pushReturn) {
    int32_t level = stack->pop<int32_t>();
    TextureId texture = stack->pop<TextureId>();
    TextureImageTarget texture_target = stack->pop<TextureImageTarget>();
    FramebufferAttachment framebuffer_attachment = stack->pop<FramebufferAttachment>();
    FramebufferTarget framebuffer_target = stack->pop<FramebufferTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glFramebufferTexture2D(%u, %u, %u, %u, %d)\n", framebuffer_target,
                   framebuffer_attachment, texture_target, texture, level);
        if (glFramebufferTexture2D != nullptr) {
            glFramebufferTexture2D(framebuffer_target, framebuffer_attachment, texture_target,
                                   texture, level);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFramebufferTexture2D\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFramebufferTexture2D\n");
        return false;
    }
}

bool callGlGetFramebufferAttachmentParameteriv(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    FramebufferAttachmentParameter parameter = stack->pop<FramebufferAttachmentParameter>();
    FramebufferAttachment attachment = stack->pop<FramebufferAttachment>();
    FramebufferTarget framebuffer_target = stack->pop<FramebufferTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glGetFramebufferAttachmentParameteriv(%u, %u, %u, %p)\n", framebuffer_target,
                   attachment, parameter, value);
        if (glGetFramebufferAttachmentParameteriv != nullptr) {
            glGetFramebufferAttachmentParameteriv(framebuffer_target, attachment, parameter, value);
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function "
                    "glGetFramebufferAttachmentParameteriv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetFramebufferAttachmentParameteriv\n");
        return false;
    }
}

bool callGlDrawElements(Stack* stack, bool pushReturn) {
    IndicesPointer indices = stack->pop<IndicesPointer>();
    IndicesType indices_type = stack->pop<IndicesType>();
    int32_t element_count = stack->pop<int32_t>();
    DrawMode draw_mode = stack->pop<DrawMode>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawElements(%u, %d, %u, %p)\n", draw_mode, element_count, indices_type,
                   indices);
        if (glDrawElements != nullptr) {
            glDrawElements(draw_mode, element_count, indices_type, indices);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawElements\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawElements\n");
        return false;
    }
}

bool callGlDrawArrays(Stack* stack, bool pushReturn) {
    int32_t index_count = stack->pop<int32_t>();
    int32_t first_index = stack->pop<int32_t>();
    DrawMode draw_mode = stack->pop<DrawMode>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawArrays(%u, %d, %d)\n", draw_mode, first_index, index_count);
        if (glDrawArrays != nullptr) {
            glDrawArrays(draw_mode, first_index, index_count);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawArrays\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawArrays\n");
        return false;
    }
}

bool callGlFlush(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glFlush()\n");
        if (glFlush != nullptr) {
            glFlush();
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFlush\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFlush\n");
        return false;
    }
}

bool callGlFinish(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glFinish()\n");
        if (glFinish != nullptr) {
            glFinish();
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFinish\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFinish\n");
        return false;
    }
}

bool callGlGetBooleanv(Stack* stack, bool pushReturn) {
    bool* values = stack->pop<bool*>();
    StateVariable param = stack->pop<StateVariable>();
    if (stack->isValid()) {
        GAPID_INFO("glGetBooleanv(%u, %p)\n", param, values);
        if (glGetBooleanv != nullptr) {
            glGetBooleanv(param, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetBooleanv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetBooleanv\n");
        return false;
    }
}

bool callGlGetFloatv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    StateVariable param = stack->pop<StateVariable>();
    if (stack->isValid()) {
        GAPID_INFO("glGetFloatv(%u, %p)\n", param, values);
        if (glGetFloatv != nullptr) {
            glGetFloatv(param, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetFloatv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetFloatv\n");
        return false;
    }
}

bool callGlGetIntegerv(Stack* stack, bool pushReturn) {
    int32_t* values = stack->pop<int32_t*>();
    StateVariable param = stack->pop<StateVariable>();
    if (stack->isValid()) {
        GAPID_INFO("glGetIntegerv(%u, %p)\n", param, values);
        if (glGetIntegerv != nullptr) {
            glGetIntegerv(param, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetIntegerv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetIntegerv\n");
        return false;
    }
}

bool callGlGetString(Stack* stack, bool pushReturn) {
    StringConstant param = stack->pop<StringConstant>();
    if (stack->isValid()) {
        GAPID_INFO("glGetString(%u)\n", param);
        if (glGetString != nullptr) {
            char* return_value = glGetString(param);
            GAPID_INFO("Returned: %s\n", return_value);
            if (pushReturn) {
                stack->push<char*>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetString\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetString\n");
        return false;
    }
}

bool callGlEnable(Stack* stack, bool pushReturn) {
    Capability capability = stack->pop<Capability>();
    if (stack->isValid()) {
        GAPID_INFO("glEnable(%u)\n", capability);
        if (glEnable != nullptr) {
            glEnable(capability);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEnable\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEnable\n");
        return false;
    }
}

bool callGlDisable(Stack* stack, bool pushReturn) {
    Capability capability = stack->pop<Capability>();
    if (stack->isValid()) {
        GAPID_INFO("glDisable(%u)\n", capability);
        if (glDisable != nullptr) {
            glDisable(capability);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDisable\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDisable\n");
        return false;
    }
}

bool callGlIsEnabled(Stack* stack, bool pushReturn) {
    Capability capability = stack->pop<Capability>();
    if (stack->isValid()) {
        GAPID_INFO("glIsEnabled(%u)\n", capability);
        if (glIsEnabled != nullptr) {
            bool return_value = glIsEnabled(capability);
            GAPID_INFO("Returned: %d\n", return_value);
            if (pushReturn) {
                stack->push<bool>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsEnabled\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsEnabled\n");
        return false;
    }
}

bool callGlMapBufferRange(Stack* stack, bool pushReturn) {
    MapBufferRangeAccess access = stack->pop<MapBufferRangeAccess>();
    int32_t length = stack->pop<int32_t>();
    int32_t offset = stack->pop<int32_t>();
    MapBufferTarget target = stack->pop<MapBufferTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glMapBufferRange(%u, %d, %d, %u)\n", target, offset, length, access);
        if (glMapBufferRange != nullptr) {
            void* return_value = glMapBufferRange(target, offset, length, access);
            GAPID_INFO("Returned: %p\n", return_value);
            if (pushReturn) {
                stack->push<void*>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMapBufferRange\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMapBufferRange\n");
        return false;
    }
}

bool callGlUnmapBuffer(Stack* stack, bool pushReturn) {
    MapBufferTarget target = stack->pop<MapBufferTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glUnmapBuffer(%u)\n", target);
        if (glUnmapBuffer != nullptr) {
            glUnmapBuffer(target);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUnmapBuffer\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUnmapBuffer\n");
        return false;
    }
}

bool callGlInvalidateFramebuffer(Stack* stack, bool pushReturn) {
    const FramebufferAttachment* attachments = stack->pop<const FramebufferAttachment*>();
    int32_t count = stack->pop<int32_t>();
    FramebufferTarget target = stack->pop<FramebufferTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glInvalidateFramebuffer(%u, %d, %p)\n", target, count, attachments);
        if (glInvalidateFramebuffer != nullptr) {
            glInvalidateFramebuffer(target, count, attachments);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glInvalidateFramebuffer\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glInvalidateFramebuffer\n");
        return false;
    }
}

bool callGlRenderbufferStorageMultisample(Stack* stack, bool pushReturn) {
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    RenderbufferFormat format = stack->pop<RenderbufferFormat>();
    int32_t samples = stack->pop<int32_t>();
    RenderbufferTarget target = stack->pop<RenderbufferTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glRenderbufferStorageMultisample(%u, %d, %u, %d, %d)\n", target, samples,
                   format, width, height);
        if (glRenderbufferStorageMultisample != nullptr) {
            glRenderbufferStorageMultisample(target, samples, format, width, height);
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glRenderbufferStorageMultisample\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glRenderbufferStorageMultisample\n");
        return false;
    }
}

bool callGlBlitFramebuffer(Stack* stack, bool pushReturn) {
    TextureFilterMode filter = stack->pop<TextureFilterMode>();
    ClearMask mask = stack->pop<ClearMask>();
    int32_t dstY1 = stack->pop<int32_t>();
    int32_t dstX1 = stack->pop<int32_t>();
    int32_t dstY0 = stack->pop<int32_t>();
    int32_t dstX0 = stack->pop<int32_t>();
    int32_t srcY1 = stack->pop<int32_t>();
    int32_t srcX1 = stack->pop<int32_t>();
    int32_t srcY0 = stack->pop<int32_t>();
    int32_t srcX0 = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBlitFramebuffer(%d, %d, %d, %d, %d, %d, %d, %d, %u, %u)\n", srcX0, srcY0,
                   srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter);
        if (glBlitFramebuffer != nullptr) {
            glBlitFramebuffer(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlitFramebuffer\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlitFramebuffer\n");
        return false;
    }
}

bool callGlGenQueries(Stack* stack, bool pushReturn) {
    QueryId* queries = stack->pop<QueryId*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenQueries(%d, %p)\n", count, queries);
        if (glGenQueries != nullptr) {
            glGenQueries(count, queries);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenQueries\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenQueries\n");
        return false;
    }
}

bool callGlBeginQuery(Stack* stack, bool pushReturn) {
    QueryId query = stack->pop<QueryId>();
    QueryTarget target = stack->pop<QueryTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glBeginQuery(%u, %u)\n", target, query);
        if (glBeginQuery != nullptr) {
            glBeginQuery(target, query);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBeginQuery\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBeginQuery\n");
        return false;
    }
}

bool callGlEndQuery(Stack* stack, bool pushReturn) {
    QueryTarget target = stack->pop<QueryTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glEndQuery(%u)\n", target);
        if (glEndQuery != nullptr) {
            glEndQuery(target);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEndQuery\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEndQuery\n");
        return false;
    }
}

bool callGlDeleteQueries(Stack* stack, bool pushReturn) {
    const QueryId* queries = stack->pop<const QueryId*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteQueries(%d, %p)\n", count, queries);
        if (glDeleteQueries != nullptr) {
            glDeleteQueries(count, queries);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteQueries\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteQueries\n");
        return false;
    }
}

bool callGlIsQuery(Stack* stack, bool pushReturn) {
    QueryId query = stack->pop<QueryId>();
    if (stack->isValid()) {
        GAPID_INFO("glIsQuery(%u)\n", query);
        if (glIsQuery != nullptr) {
            bool return_value = glIsQuery(query);
            GAPID_INFO("Returned: %d\n", return_value);
            if (pushReturn) {
                stack->push<bool>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsQuery\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsQuery\n");
        return false;
    }
}

bool callGlGetQueryiv(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    QueryParameter parameter = stack->pop<QueryParameter>();
    QueryTarget target = stack->pop<QueryTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glGetQueryiv(%u, %u, %p)\n", target, parameter, value);
        if (glGetQueryiv != nullptr) {
            glGetQueryiv(target, parameter, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetQueryiv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetQueryiv\n");
        return false;
    }
}

bool callGlGetQueryObjectuiv(Stack* stack, bool pushReturn) {
    uint32_t* value = stack->pop<uint32_t*>();
    QueryObjectParameter parameter = stack->pop<QueryObjectParameter>();
    QueryId query = stack->pop<QueryId>();
    if (stack->isValid()) {
        GAPID_INFO("glGetQueryObjectuiv(%u, %u, %p)\n", query, parameter, value);
        if (glGetQueryObjectuiv != nullptr) {
            glGetQueryObjectuiv(query, parameter, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetQueryObjectuiv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetQueryObjectuiv\n");
        return false;
    }
}

bool callGlGenQueriesEXT(Stack* stack, bool pushReturn) {
    QueryId* queries = stack->pop<QueryId*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenQueriesEXT(%d, %p)\n", count, queries);
        if (glGenQueriesEXT != nullptr) {
            glGenQueriesEXT(count, queries);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenQueriesEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenQueriesEXT\n");
        return false;
    }
}

bool callGlBeginQueryEXT(Stack* stack, bool pushReturn) {
    QueryId query = stack->pop<QueryId>();
    QueryTarget target = stack->pop<QueryTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glBeginQueryEXT(%u, %u)\n", target, query);
        if (glBeginQueryEXT != nullptr) {
            glBeginQueryEXT(target, query);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBeginQueryEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBeginQueryEXT\n");
        return false;
    }
}

bool callGlEndQueryEXT(Stack* stack, bool pushReturn) {
    QueryTarget target = stack->pop<QueryTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glEndQueryEXT(%u)\n", target);
        if (glEndQueryEXT != nullptr) {
            glEndQueryEXT(target);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEndQueryEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEndQueryEXT\n");
        return false;
    }
}

bool callGlDeleteQueriesEXT(Stack* stack, bool pushReturn) {
    const QueryId* queries = stack->pop<const QueryId*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteQueriesEXT(%d, %p)\n", count, queries);
        if (glDeleteQueriesEXT != nullptr) {
            glDeleteQueriesEXT(count, queries);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteQueriesEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteQueriesEXT\n");
        return false;
    }
}

bool callGlIsQueryEXT(Stack* stack, bool pushReturn) {
    QueryId query = stack->pop<QueryId>();
    if (stack->isValid()) {
        GAPID_INFO("glIsQueryEXT(%u)\n", query);
        if (glIsQueryEXT != nullptr) {
            bool return_value = glIsQueryEXT(query);
            GAPID_INFO("Returned: %d\n", return_value);
            if (pushReturn) {
                stack->push<bool>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsQueryEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsQueryEXT\n");
        return false;
    }
}

bool callGlQueryCounterEXT(Stack* stack, bool pushReturn) {
    QueryTarget target = stack->pop<QueryTarget>();
    QueryId query = stack->pop<QueryId>();
    if (stack->isValid()) {
        GAPID_INFO("glQueryCounterEXT(%u, %u)\n", query, target);
        if (glQueryCounterEXT != nullptr) {
            glQueryCounterEXT(query, target);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glQueryCounterEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glQueryCounterEXT\n");
        return false;
    }
}

bool callGlGetQueryivEXT(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    QueryParameter parameter = stack->pop<QueryParameter>();
    QueryTarget target = stack->pop<QueryTarget>();
    if (stack->isValid()) {
        GAPID_INFO("glGetQueryivEXT(%u, %u, %p)\n", target, parameter, value);
        if (glGetQueryivEXT != nullptr) {
            glGetQueryivEXT(target, parameter, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetQueryivEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetQueryivEXT\n");
        return false;
    }
}

bool callGlGetQueryObjectivEXT(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    QueryObjectParameter parameter = stack->pop<QueryObjectParameter>();
    QueryId query = stack->pop<QueryId>();
    if (stack->isValid()) {
        GAPID_INFO("glGetQueryObjectivEXT(%u, %u, %p)\n", query, parameter, value);
        if (glGetQueryObjectivEXT != nullptr) {
            glGetQueryObjectivEXT(query, parameter, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetQueryObjectivEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetQueryObjectivEXT\n");
        return false;
    }
}

bool callGlGetQueryObjectuivEXT(Stack* stack, bool pushReturn) {
    uint32_t* value = stack->pop<uint32_t*>();
    QueryObjectParameter parameter = stack->pop<QueryObjectParameter>();
    QueryId query = stack->pop<QueryId>();
    if (stack->isValid()) {
        GAPID_INFO("glGetQueryObjectuivEXT(%u, %u, %p)\n", query, parameter, value);
        if (glGetQueryObjectuivEXT != nullptr) {
            glGetQueryObjectuivEXT(query, parameter, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetQueryObjectuivEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetQueryObjectuivEXT\n");
        return false;
    }
}

bool callGlGetQueryObjecti64vEXT(Stack* stack, bool pushReturn) {
    int64_t* value = stack->pop<int64_t*>();
    QueryObjectParameter parameter = stack->pop<QueryObjectParameter>();
    QueryId query = stack->pop<QueryId>();
    if (stack->isValid()) {
        GAPID_INFO("glGetQueryObjecti64vEXT(%u, %u, %p)\n", query, parameter, value);
        if (glGetQueryObjecti64vEXT != nullptr) {
            glGetQueryObjecti64vEXT(query, parameter, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetQueryObjecti64vEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetQueryObjecti64vEXT\n");
        return false;
    }
}

bool callGlGetQueryObjectui64vEXT(Stack* stack, bool pushReturn) {
    uint64_t* value = stack->pop<uint64_t*>();
    QueryObjectParameter parameter = stack->pop<QueryObjectParameter>();
    QueryId query = stack->pop<QueryId>();
    if (stack->isValid()) {
        GAPID_INFO("glGetQueryObjectui64vEXT(%u, %u, %p)\n", query, parameter, value);
        if (glGetQueryObjectui64vEXT != nullptr) {
            glGetQueryObjectui64vEXT(query, parameter, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetQueryObjectui64vEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetQueryObjectui64vEXT\n");
        return false;
    }
}

}  // end of anonymous namespace

PFNEGLCREATECONTEXT eglCreateContext = nullptr;
PFNEGLMAKECURRENT eglMakeCurrent = nullptr;
PFNEGLSWAPBUFFERS eglSwapBuffers = nullptr;
PFNWGLSWAPBUFFERS wglSwapBuffers = nullptr;
PFNGLENABLECLIENTSTATE glEnableClientState = nullptr;
PFNGLDISABLECLIENTSTATE glDisableClientState = nullptr;
PFNGLGETPROGRAMBINARYOES glGetProgramBinaryOES = nullptr;
PFNGLPROGRAMBINARYOES glProgramBinaryOES = nullptr;
PFNGLSTARTTILINGQCOM glStartTilingQCOM = nullptr;
PFNGLENDTILINGQCOM glEndTilingQCOM = nullptr;
PFNGLDISCARDFRAMEBUFFEREXT glDiscardFramebufferEXT = nullptr;
PFNGLINSERTEVENTMARKEREXT glInsertEventMarkerEXT = nullptr;
PFNGLPUSHGROUPMARKEREXT glPushGroupMarkerEXT = nullptr;
PFNGLPOPGROUPMARKEREXT glPopGroupMarkerEXT = nullptr;
PFNGLTEXSTORAGE1DEXT glTexStorage1DEXT = nullptr;
PFNGLTEXSTORAGE2DEXT glTexStorage2DEXT = nullptr;
PFNGLTEXSTORAGE3DEXT glTexStorage3DEXT = nullptr;
PFNGLTEXTURESTORAGE1DEXT glTextureStorage1DEXT = nullptr;
PFNGLTEXTURESTORAGE2DEXT glTextureStorage2DEXT = nullptr;
PFNGLTEXTURESTORAGE3DEXT glTextureStorage3DEXT = nullptr;
PFNGLGENVERTEXARRAYSOES glGenVertexArraysOES = nullptr;
PFNGLBINDVERTEXARRAYOES glBindVertexArrayOES = nullptr;
PFNGLDELETEVERTEXARRAYSOES glDeleteVertexArraysOES = nullptr;
PFNGLISVERTEXARRAYOES glIsVertexArrayOES = nullptr;
PFNGLEGLIMAGETARGETTEXTURE2DOES glEGLImageTargetTexture2DOES = nullptr;
PFNGLEGLIMAGETARGETRENDERBUFFERSTORAGEOES glEGLImageTargetRenderbufferStorageOES = nullptr;
PFNGLGETGRAPHICSRESETSTATUSEXT glGetGraphicsResetStatusEXT = nullptr;
PFNGLBINDATTRIBLOCATION glBindAttribLocation = nullptr;
PFNGLBLENDFUNC glBlendFunc = nullptr;
PFNGLBLENDFUNCSEPARATE glBlendFuncSeparate = nullptr;
PFNGLBLENDEQUATION glBlendEquation = nullptr;
PFNGLBLENDEQUATIONSEPARATE glBlendEquationSeparate = nullptr;
PFNGLBLENDCOLOR glBlendColor = nullptr;
PFNGLENABLEVERTEXATTRIBARRAY glEnableVertexAttribArray = nullptr;
PFNGLDISABLEVERTEXATTRIBARRAY glDisableVertexAttribArray = nullptr;
PFNGLVERTEXATTRIBPOINTER glVertexAttribPointer = nullptr;
PFNGLGETACTIVEATTRIB glGetActiveAttrib = nullptr;
PFNGLGETACTIVEUNIFORM glGetActiveUniform = nullptr;
PFNGLGETERROR glGetError = nullptr;
PFNGLGETPROGRAMIV glGetProgramiv = nullptr;
PFNGLGETSHADERIV glGetShaderiv = nullptr;
PFNGLGETUNIFORMLOCATION glGetUniformLocation = nullptr;
PFNGLGETATTRIBLOCATION glGetAttribLocation = nullptr;
PFNGLPIXELSTOREI glPixelStorei = nullptr;
PFNGLTEXPARAMETERI glTexParameteri = nullptr;
PFNGLTEXPARAMETERF glTexParameterf = nullptr;
PFNGLGETTEXPARAMETERIV glGetTexParameteriv = nullptr;
PFNGLGETTEXPARAMETERFV glGetTexParameterfv = nullptr;
PFNGLUNIFORM1I glUniform1i = nullptr;
PFNGLUNIFORM2I glUniform2i = nullptr;
PFNGLUNIFORM3I glUniform3i = nullptr;
PFNGLUNIFORM4I glUniform4i = nullptr;
PFNGLUNIFORM1IV glUniform1iv = nullptr;
PFNGLUNIFORM2IV glUniform2iv = nullptr;
PFNGLUNIFORM3IV glUniform3iv = nullptr;
PFNGLUNIFORM4IV glUniform4iv = nullptr;
PFNGLUNIFORM1F glUniform1f = nullptr;
PFNGLUNIFORM2F glUniform2f = nullptr;
PFNGLUNIFORM3F glUniform3f = nullptr;
PFNGLUNIFORM4F glUniform4f = nullptr;
PFNGLUNIFORM1FV glUniform1fv = nullptr;
PFNGLUNIFORM2FV glUniform2fv = nullptr;
PFNGLUNIFORM3FV glUniform3fv = nullptr;
PFNGLUNIFORM4FV glUniform4fv = nullptr;
PFNGLUNIFORMMATRIX2FV glUniformMatrix2fv = nullptr;
PFNGLUNIFORMMATRIX3FV glUniformMatrix3fv = nullptr;
PFNGLUNIFORMMATRIX4FV glUniformMatrix4fv = nullptr;
PFNGLGETUNIFORMFV glGetUniformfv = nullptr;
PFNGLGETUNIFORMIV glGetUniformiv = nullptr;
PFNGLVERTEXATTRIB1F glVertexAttrib1f = nullptr;
PFNGLVERTEXATTRIB2F glVertexAttrib2f = nullptr;
PFNGLVERTEXATTRIB3F glVertexAttrib3f = nullptr;
PFNGLVERTEXATTRIB4F glVertexAttrib4f = nullptr;
PFNGLVERTEXATTRIB1FV glVertexAttrib1fv = nullptr;
PFNGLVERTEXATTRIB2FV glVertexAttrib2fv = nullptr;
PFNGLVERTEXATTRIB3FV glVertexAttrib3fv = nullptr;
PFNGLVERTEXATTRIB4FV glVertexAttrib4fv = nullptr;
PFNGLGETSHADERPRECISIONFORMAT glGetShaderPrecisionFormat = nullptr;
PFNGLDEPTHMASK glDepthMask = nullptr;
PFNGLDEPTHFUNC glDepthFunc = nullptr;
PFNGLDEPTHRANGEF glDepthRangef = nullptr;
PFNGLCOLORMASK glColorMask = nullptr;
PFNGLSTENCILMASK glStencilMask = nullptr;
PFNGLSTENCILMASKSEPARATE glStencilMaskSeparate = nullptr;
PFNGLSTENCILFUNCSEPARATE glStencilFuncSeparate = nullptr;
PFNGLSTENCILOPSEPARATE glStencilOpSeparate = nullptr;
PFNGLFRONTFACE glFrontFace = nullptr;
PFNGLVIEWPORT glViewport = nullptr;
PFNGLSCISSOR glScissor = nullptr;
PFNGLACTIVETEXTURE glActiveTexture = nullptr;
PFNGLGENTEXTURES glGenTextures = nullptr;
PFNGLDELETETEXTURES glDeleteTextures = nullptr;
PFNGLISTEXTURE glIsTexture = nullptr;
PFNGLBINDTEXTURE glBindTexture = nullptr;
PFNGLTEXIMAGE2D glTexImage2D = nullptr;
PFNGLTEXSUBIMAGE2D glTexSubImage2D = nullptr;
PFNGLCOPYTEXIMAGE2D glCopyTexImage2D = nullptr;
PFNGLCOPYTEXSUBIMAGE2D glCopyTexSubImage2D = nullptr;
PFNGLCOMPRESSEDTEXIMAGE2D glCompressedTexImage2D = nullptr;
PFNGLCOMPRESSEDTEXSUBIMAGE2D glCompressedTexSubImage2D = nullptr;
PFNGLGENERATEMIPMAP glGenerateMipmap = nullptr;
PFNGLREADPIXELS glReadPixels = nullptr;
PFNGLGENFRAMEBUFFERS glGenFramebuffers = nullptr;
PFNGLBINDFRAMEBUFFER glBindFramebuffer = nullptr;
PFNGLCHECKFRAMEBUFFERSTATUS glCheckFramebufferStatus = nullptr;
PFNGLDELETEFRAMEBUFFERS glDeleteFramebuffers = nullptr;
PFNGLISFRAMEBUFFER glIsFramebuffer = nullptr;
PFNGLGENRENDERBUFFERS glGenRenderbuffers = nullptr;
PFNGLBINDRENDERBUFFER glBindRenderbuffer = nullptr;
PFNGLRENDERBUFFERSTORAGE glRenderbufferStorage = nullptr;
PFNGLDELETERENDERBUFFERS glDeleteRenderbuffers = nullptr;
PFNGLISRENDERBUFFER glIsRenderbuffer = nullptr;
PFNGLGETRENDERBUFFERPARAMETERIV glGetRenderbufferParameteriv = nullptr;
PFNGLGENBUFFERS glGenBuffers = nullptr;
PFNGLBINDBUFFER glBindBuffer = nullptr;
PFNGLBUFFERDATA glBufferData = nullptr;
PFNGLBUFFERSUBDATA glBufferSubData = nullptr;
PFNGLDELETEBUFFERS glDeleteBuffers = nullptr;
PFNGLISBUFFER glIsBuffer = nullptr;
PFNGLGETBUFFERPARAMETERIV glGetBufferParameteriv = nullptr;
PFNGLCREATESHADER glCreateShader = nullptr;
PFNGLDELETESHADER glDeleteShader = nullptr;
PFNGLSHADERSOURCE glShaderSource = nullptr;
PFNGLSHADERBINARY glShaderBinary = nullptr;
PFNGLGETSHADERINFOLOG glGetShaderInfoLog = nullptr;
PFNGLGETSHADERSOURCE glGetShaderSource = nullptr;
PFNGLRELEASESHADERCOMPILER glReleaseShaderCompiler = nullptr;
PFNGLCOMPILESHADER glCompileShader = nullptr;
PFNGLISSHADER glIsShader = nullptr;
PFNGLCREATEPROGRAM glCreateProgram = nullptr;
PFNGLDELETEPROGRAM glDeleteProgram = nullptr;
PFNGLATTACHSHADER glAttachShader = nullptr;
PFNGLDETACHSHADER glDetachShader = nullptr;
PFNGLGETATTACHEDSHADERS glGetAttachedShaders = nullptr;
PFNGLLINKPROGRAM glLinkProgram = nullptr;
PFNGLGETPROGRAMINFOLOG glGetProgramInfoLog = nullptr;
PFNGLUSEPROGRAM glUseProgram = nullptr;
PFNGLISPROGRAM glIsProgram = nullptr;
PFNGLVALIDATEPROGRAM glValidateProgram = nullptr;
PFNGLCLEARCOLOR glClearColor = nullptr;
PFNGLCLEARDEPTHF glClearDepthf = nullptr;
PFNGLCLEARSTENCIL glClearStencil = nullptr;
PFNGLCLEAR glClear = nullptr;
PFNGLCULLFACE glCullFace = nullptr;
PFNGLPOLYGONOFFSET glPolygonOffset = nullptr;
PFNGLLINEWIDTH glLineWidth = nullptr;
PFNGLSAMPLECOVERAGE glSampleCoverage = nullptr;
PFNGLHINT glHint = nullptr;
PFNGLFRAMEBUFFERRENDERBUFFER glFramebufferRenderbuffer = nullptr;
PFNGLFRAMEBUFFERTEXTURE2D glFramebufferTexture2D = nullptr;
PFNGLGETFRAMEBUFFERATTACHMENTPARAMETERIV glGetFramebufferAttachmentParameteriv = nullptr;
PFNGLDRAWELEMENTS glDrawElements = nullptr;
PFNGLDRAWARRAYS glDrawArrays = nullptr;
PFNGLFLUSH glFlush = nullptr;
PFNGLFINISH glFinish = nullptr;
PFNGLGETBOOLEANV glGetBooleanv = nullptr;
PFNGLGETFLOATV glGetFloatv = nullptr;
PFNGLGETINTEGERV glGetIntegerv = nullptr;
PFNGLGETSTRING glGetString = nullptr;
PFNGLENABLE glEnable = nullptr;
PFNGLDISABLE glDisable = nullptr;
PFNGLISENABLED glIsEnabled = nullptr;
PFNGLMAPBUFFERRANGE glMapBufferRange = nullptr;
PFNGLUNMAPBUFFER glUnmapBuffer = nullptr;
PFNGLINVALIDATEFRAMEBUFFER glInvalidateFramebuffer = nullptr;
PFNGLRENDERBUFFERSTORAGEMULTISAMPLE glRenderbufferStorageMultisample = nullptr;
PFNGLBLITFRAMEBUFFER glBlitFramebuffer = nullptr;
PFNGLGENQUERIES glGenQueries = nullptr;
PFNGLBEGINQUERY glBeginQuery = nullptr;
PFNGLENDQUERY glEndQuery = nullptr;
PFNGLDELETEQUERIES glDeleteQueries = nullptr;
PFNGLISQUERY glIsQuery = nullptr;
PFNGLGETQUERYIV glGetQueryiv = nullptr;
PFNGLGETQUERYOBJECTUIV glGetQueryObjectuiv = nullptr;
PFNGLGENQUERIESEXT glGenQueriesEXT = nullptr;
PFNGLBEGINQUERYEXT glBeginQueryEXT = nullptr;
PFNGLENDQUERYEXT glEndQueryEXT = nullptr;
PFNGLDELETEQUERIESEXT glDeleteQueriesEXT = nullptr;
PFNGLISQUERYEXT glIsQueryEXT = nullptr;
PFNGLQUERYCOUNTEREXT glQueryCounterEXT = nullptr;
PFNGLGETQUERYIVEXT glGetQueryivEXT = nullptr;
PFNGLGETQUERYOBJECTIVEXT glGetQueryObjectivEXT = nullptr;
PFNGLGETQUERYOBJECTUIVEXT glGetQueryObjectuivEXT = nullptr;
PFNGLGETQUERYOBJECTI64VEXT glGetQueryObjecti64vEXT = nullptr;
PFNGLGETQUERYOBJECTUI64VEXT glGetQueryObjectui64vEXT = nullptr;

void Register(Interpreter* interpreter) {
    interpreter->registerFunction(Ids::EglCreateContext, callEglCreateContext);
    interpreter->registerFunction(Ids::EglMakeCurrent, callEglMakeCurrent);
    interpreter->registerFunction(Ids::EglSwapBuffers, callEglSwapBuffers);
    interpreter->registerFunction(Ids::WglSwapBuffers, callWglSwapBuffers);
    interpreter->registerFunction(Ids::GlEnableClientState, callGlEnableClientState);
    interpreter->registerFunction(Ids::GlDisableClientState, callGlDisableClientState);
    interpreter->registerFunction(Ids::GlGetProgramBinaryOES, callGlGetProgramBinaryOES);
    interpreter->registerFunction(Ids::GlProgramBinaryOES, callGlProgramBinaryOES);
    interpreter->registerFunction(Ids::GlStartTilingQCOM, callGlStartTilingQCOM);
    interpreter->registerFunction(Ids::GlEndTilingQCOM, callGlEndTilingQCOM);
    interpreter->registerFunction(Ids::GlDiscardFramebufferEXT, callGlDiscardFramebufferEXT);
    interpreter->registerFunction(Ids::GlInsertEventMarkerEXT, callGlInsertEventMarkerEXT);
    interpreter->registerFunction(Ids::GlPushGroupMarkerEXT, callGlPushGroupMarkerEXT);
    interpreter->registerFunction(Ids::GlPopGroupMarkerEXT, callGlPopGroupMarkerEXT);
    interpreter->registerFunction(Ids::GlTexStorage1DEXT, callGlTexStorage1DEXT);
    interpreter->registerFunction(Ids::GlTexStorage2DEXT, callGlTexStorage2DEXT);
    interpreter->registerFunction(Ids::GlTexStorage3DEXT, callGlTexStorage3DEXT);
    interpreter->registerFunction(Ids::GlTextureStorage1DEXT, callGlTextureStorage1DEXT);
    interpreter->registerFunction(Ids::GlTextureStorage2DEXT, callGlTextureStorage2DEXT);
    interpreter->registerFunction(Ids::GlTextureStorage3DEXT, callGlTextureStorage3DEXT);
    interpreter->registerFunction(Ids::GlGenVertexArraysOES, callGlGenVertexArraysOES);
    interpreter->registerFunction(Ids::GlBindVertexArrayOES, callGlBindVertexArrayOES);
    interpreter->registerFunction(Ids::GlDeleteVertexArraysOES, callGlDeleteVertexArraysOES);
    interpreter->registerFunction(Ids::GlIsVertexArrayOES, callGlIsVertexArrayOES);
    interpreter->registerFunction(Ids::GlEGLImageTargetTexture2DOES,
                                  callGlEGLImageTargetTexture2DOES);
    interpreter->registerFunction(Ids::GlEGLImageTargetRenderbufferStorageOES,
                                  callGlEGLImageTargetRenderbufferStorageOES);
    interpreter->registerFunction(Ids::GlGetGraphicsResetStatusEXT,
                                  callGlGetGraphicsResetStatusEXT);
    interpreter->registerFunction(Ids::GlBindAttribLocation, callGlBindAttribLocation);
    interpreter->registerFunction(Ids::GlBlendFunc, callGlBlendFunc);
    interpreter->registerFunction(Ids::GlBlendFuncSeparate, callGlBlendFuncSeparate);
    interpreter->registerFunction(Ids::GlBlendEquation, callGlBlendEquation);
    interpreter->registerFunction(Ids::GlBlendEquationSeparate, callGlBlendEquationSeparate);
    interpreter->registerFunction(Ids::GlBlendColor, callGlBlendColor);
    interpreter->registerFunction(Ids::GlEnableVertexAttribArray, callGlEnableVertexAttribArray);
    interpreter->registerFunction(Ids::GlDisableVertexAttribArray, callGlDisableVertexAttribArray);
    interpreter->registerFunction(Ids::GlVertexAttribPointer, callGlVertexAttribPointer);
    interpreter->registerFunction(Ids::GlGetActiveAttrib, callGlGetActiveAttrib);
    interpreter->registerFunction(Ids::GlGetActiveUniform, callGlGetActiveUniform);
    interpreter->registerFunction(Ids::GlGetError, callGlGetError);
    interpreter->registerFunction(Ids::GlGetProgramiv, callGlGetProgramiv);
    interpreter->registerFunction(Ids::GlGetShaderiv, callGlGetShaderiv);
    interpreter->registerFunction(Ids::GlGetUniformLocation, callGlGetUniformLocation);
    interpreter->registerFunction(Ids::GlGetAttribLocation, callGlGetAttribLocation);
    interpreter->registerFunction(Ids::GlPixelStorei, callGlPixelStorei);
    interpreter->registerFunction(Ids::GlTexParameteri, callGlTexParameteri);
    interpreter->registerFunction(Ids::GlTexParameterf, callGlTexParameterf);
    interpreter->registerFunction(Ids::GlGetTexParameteriv, callGlGetTexParameteriv);
    interpreter->registerFunction(Ids::GlGetTexParameterfv, callGlGetTexParameterfv);
    interpreter->registerFunction(Ids::GlUniform1i, callGlUniform1i);
    interpreter->registerFunction(Ids::GlUniform2i, callGlUniform2i);
    interpreter->registerFunction(Ids::GlUniform3i, callGlUniform3i);
    interpreter->registerFunction(Ids::GlUniform4i, callGlUniform4i);
    interpreter->registerFunction(Ids::GlUniform1iv, callGlUniform1iv);
    interpreter->registerFunction(Ids::GlUniform2iv, callGlUniform2iv);
    interpreter->registerFunction(Ids::GlUniform3iv, callGlUniform3iv);
    interpreter->registerFunction(Ids::GlUniform4iv, callGlUniform4iv);
    interpreter->registerFunction(Ids::GlUniform1f, callGlUniform1f);
    interpreter->registerFunction(Ids::GlUniform2f, callGlUniform2f);
    interpreter->registerFunction(Ids::GlUniform3f, callGlUniform3f);
    interpreter->registerFunction(Ids::GlUniform4f, callGlUniform4f);
    interpreter->registerFunction(Ids::GlUniform1fv, callGlUniform1fv);
    interpreter->registerFunction(Ids::GlUniform2fv, callGlUniform2fv);
    interpreter->registerFunction(Ids::GlUniform3fv, callGlUniform3fv);
    interpreter->registerFunction(Ids::GlUniform4fv, callGlUniform4fv);
    interpreter->registerFunction(Ids::GlUniformMatrix2fv, callGlUniformMatrix2fv);
    interpreter->registerFunction(Ids::GlUniformMatrix3fv, callGlUniformMatrix3fv);
    interpreter->registerFunction(Ids::GlUniformMatrix4fv, callGlUniformMatrix4fv);
    interpreter->registerFunction(Ids::GlGetUniformfv, callGlGetUniformfv);
    interpreter->registerFunction(Ids::GlGetUniformiv, callGlGetUniformiv);
    interpreter->registerFunction(Ids::GlVertexAttrib1f, callGlVertexAttrib1f);
    interpreter->registerFunction(Ids::GlVertexAttrib2f, callGlVertexAttrib2f);
    interpreter->registerFunction(Ids::GlVertexAttrib3f, callGlVertexAttrib3f);
    interpreter->registerFunction(Ids::GlVertexAttrib4f, callGlVertexAttrib4f);
    interpreter->registerFunction(Ids::GlVertexAttrib1fv, callGlVertexAttrib1fv);
    interpreter->registerFunction(Ids::GlVertexAttrib2fv, callGlVertexAttrib2fv);
    interpreter->registerFunction(Ids::GlVertexAttrib3fv, callGlVertexAttrib3fv);
    interpreter->registerFunction(Ids::GlVertexAttrib4fv, callGlVertexAttrib4fv);
    interpreter->registerFunction(Ids::GlGetShaderPrecisionFormat, callGlGetShaderPrecisionFormat);
    interpreter->registerFunction(Ids::GlDepthMask, callGlDepthMask);
    interpreter->registerFunction(Ids::GlDepthFunc, callGlDepthFunc);
    interpreter->registerFunction(Ids::GlDepthRangef, callGlDepthRangef);
    interpreter->registerFunction(Ids::GlColorMask, callGlColorMask);
    interpreter->registerFunction(Ids::GlStencilMask, callGlStencilMask);
    interpreter->registerFunction(Ids::GlStencilMaskSeparate, callGlStencilMaskSeparate);
    interpreter->registerFunction(Ids::GlStencilFuncSeparate, callGlStencilFuncSeparate);
    interpreter->registerFunction(Ids::GlStencilOpSeparate, callGlStencilOpSeparate);
    interpreter->registerFunction(Ids::GlFrontFace, callGlFrontFace);
    interpreter->registerFunction(Ids::GlViewport, callGlViewport);
    interpreter->registerFunction(Ids::GlScissor, callGlScissor);
    interpreter->registerFunction(Ids::GlActiveTexture, callGlActiveTexture);
    interpreter->registerFunction(Ids::GlGenTextures, callGlGenTextures);
    interpreter->registerFunction(Ids::GlDeleteTextures, callGlDeleteTextures);
    interpreter->registerFunction(Ids::GlIsTexture, callGlIsTexture);
    interpreter->registerFunction(Ids::GlBindTexture, callGlBindTexture);
    interpreter->registerFunction(Ids::GlTexImage2D, callGlTexImage2D);
    interpreter->registerFunction(Ids::GlTexSubImage2D, callGlTexSubImage2D);
    interpreter->registerFunction(Ids::GlCopyTexImage2D, callGlCopyTexImage2D);
    interpreter->registerFunction(Ids::GlCopyTexSubImage2D, callGlCopyTexSubImage2D);
    interpreter->registerFunction(Ids::GlCompressedTexImage2D, callGlCompressedTexImage2D);
    interpreter->registerFunction(Ids::GlCompressedTexSubImage2D, callGlCompressedTexSubImage2D);
    interpreter->registerFunction(Ids::GlGenerateMipmap, callGlGenerateMipmap);
    interpreter->registerFunction(Ids::GlReadPixels, callGlReadPixels);
    interpreter->registerFunction(Ids::GlGenFramebuffers, callGlGenFramebuffers);
    interpreter->registerFunction(Ids::GlBindFramebuffer, callGlBindFramebuffer);
    interpreter->registerFunction(Ids::GlCheckFramebufferStatus, callGlCheckFramebufferStatus);
    interpreter->registerFunction(Ids::GlDeleteFramebuffers, callGlDeleteFramebuffers);
    interpreter->registerFunction(Ids::GlIsFramebuffer, callGlIsFramebuffer);
    interpreter->registerFunction(Ids::GlGenRenderbuffers, callGlGenRenderbuffers);
    interpreter->registerFunction(Ids::GlBindRenderbuffer, callGlBindRenderbuffer);
    interpreter->registerFunction(Ids::GlRenderbufferStorage, callGlRenderbufferStorage);
    interpreter->registerFunction(Ids::GlDeleteRenderbuffers, callGlDeleteRenderbuffers);
    interpreter->registerFunction(Ids::GlIsRenderbuffer, callGlIsRenderbuffer);
    interpreter->registerFunction(Ids::GlGetRenderbufferParameteriv,
                                  callGlGetRenderbufferParameteriv);
    interpreter->registerFunction(Ids::GlGenBuffers, callGlGenBuffers);
    interpreter->registerFunction(Ids::GlBindBuffer, callGlBindBuffer);
    interpreter->registerFunction(Ids::GlBufferData, callGlBufferData);
    interpreter->registerFunction(Ids::GlBufferSubData, callGlBufferSubData);
    interpreter->registerFunction(Ids::GlDeleteBuffers, callGlDeleteBuffers);
    interpreter->registerFunction(Ids::GlIsBuffer, callGlIsBuffer);
    interpreter->registerFunction(Ids::GlGetBufferParameteriv, callGlGetBufferParameteriv);
    interpreter->registerFunction(Ids::GlCreateShader, callGlCreateShader);
    interpreter->registerFunction(Ids::GlDeleteShader, callGlDeleteShader);
    interpreter->registerFunction(Ids::GlShaderSource, callGlShaderSource);
    interpreter->registerFunction(Ids::GlShaderBinary, callGlShaderBinary);
    interpreter->registerFunction(Ids::GlGetShaderInfoLog, callGlGetShaderInfoLog);
    interpreter->registerFunction(Ids::GlGetShaderSource, callGlGetShaderSource);
    interpreter->registerFunction(Ids::GlReleaseShaderCompiler, callGlReleaseShaderCompiler);
    interpreter->registerFunction(Ids::GlCompileShader, callGlCompileShader);
    interpreter->registerFunction(Ids::GlIsShader, callGlIsShader);
    interpreter->registerFunction(Ids::GlCreateProgram, callGlCreateProgram);
    interpreter->registerFunction(Ids::GlDeleteProgram, callGlDeleteProgram);
    interpreter->registerFunction(Ids::GlAttachShader, callGlAttachShader);
    interpreter->registerFunction(Ids::GlDetachShader, callGlDetachShader);
    interpreter->registerFunction(Ids::GlGetAttachedShaders, callGlGetAttachedShaders);
    interpreter->registerFunction(Ids::GlLinkProgram, callGlLinkProgram);
    interpreter->registerFunction(Ids::GlGetProgramInfoLog, callGlGetProgramInfoLog);
    interpreter->registerFunction(Ids::GlUseProgram, callGlUseProgram);
    interpreter->registerFunction(Ids::GlIsProgram, callGlIsProgram);
    interpreter->registerFunction(Ids::GlValidateProgram, callGlValidateProgram);
    interpreter->registerFunction(Ids::GlClearColor, callGlClearColor);
    interpreter->registerFunction(Ids::GlClearDepthf, callGlClearDepthf);
    interpreter->registerFunction(Ids::GlClearStencil, callGlClearStencil);
    interpreter->registerFunction(Ids::GlClear, callGlClear);
    interpreter->registerFunction(Ids::GlCullFace, callGlCullFace);
    interpreter->registerFunction(Ids::GlPolygonOffset, callGlPolygonOffset);
    interpreter->registerFunction(Ids::GlLineWidth, callGlLineWidth);
    interpreter->registerFunction(Ids::GlSampleCoverage, callGlSampleCoverage);
    interpreter->registerFunction(Ids::GlHint, callGlHint);
    interpreter->registerFunction(Ids::GlFramebufferRenderbuffer, callGlFramebufferRenderbuffer);
    interpreter->registerFunction(Ids::GlFramebufferTexture2D, callGlFramebufferTexture2D);
    interpreter->registerFunction(Ids::GlGetFramebufferAttachmentParameteriv,
                                  callGlGetFramebufferAttachmentParameteriv);
    interpreter->registerFunction(Ids::GlDrawElements, callGlDrawElements);
    interpreter->registerFunction(Ids::GlDrawArrays, callGlDrawArrays);
    interpreter->registerFunction(Ids::GlFlush, callGlFlush);
    interpreter->registerFunction(Ids::GlFinish, callGlFinish);
    interpreter->registerFunction(Ids::GlGetBooleanv, callGlGetBooleanv);
    interpreter->registerFunction(Ids::GlGetFloatv, callGlGetFloatv);
    interpreter->registerFunction(Ids::GlGetIntegerv, callGlGetIntegerv);
    interpreter->registerFunction(Ids::GlGetString, callGlGetString);
    interpreter->registerFunction(Ids::GlEnable, callGlEnable);
    interpreter->registerFunction(Ids::GlDisable, callGlDisable);
    interpreter->registerFunction(Ids::GlIsEnabled, callGlIsEnabled);
    interpreter->registerFunction(Ids::GlMapBufferRange, callGlMapBufferRange);
    interpreter->registerFunction(Ids::GlUnmapBuffer, callGlUnmapBuffer);
    interpreter->registerFunction(Ids::GlInvalidateFramebuffer, callGlInvalidateFramebuffer);
    interpreter->registerFunction(Ids::GlRenderbufferStorageMultisample,
                                  callGlRenderbufferStorageMultisample);
    interpreter->registerFunction(Ids::GlBlitFramebuffer, callGlBlitFramebuffer);
    interpreter->registerFunction(Ids::GlGenQueries, callGlGenQueries);
    interpreter->registerFunction(Ids::GlBeginQuery, callGlBeginQuery);
    interpreter->registerFunction(Ids::GlEndQuery, callGlEndQuery);
    interpreter->registerFunction(Ids::GlDeleteQueries, callGlDeleteQueries);
    interpreter->registerFunction(Ids::GlIsQuery, callGlIsQuery);
    interpreter->registerFunction(Ids::GlGetQueryiv, callGlGetQueryiv);
    interpreter->registerFunction(Ids::GlGetQueryObjectuiv, callGlGetQueryObjectuiv);
    interpreter->registerFunction(Ids::GlGenQueriesEXT, callGlGenQueriesEXT);
    interpreter->registerFunction(Ids::GlBeginQueryEXT, callGlBeginQueryEXT);
    interpreter->registerFunction(Ids::GlEndQueryEXT, callGlEndQueryEXT);
    interpreter->registerFunction(Ids::GlDeleteQueriesEXT, callGlDeleteQueriesEXT);
    interpreter->registerFunction(Ids::GlIsQueryEXT, callGlIsQueryEXT);
    interpreter->registerFunction(Ids::GlQueryCounterEXT, callGlQueryCounterEXT);
    interpreter->registerFunction(Ids::GlGetQueryivEXT, callGlGetQueryivEXT);
    interpreter->registerFunction(Ids::GlGetQueryObjectivEXT, callGlGetQueryObjectivEXT);
    interpreter->registerFunction(Ids::GlGetQueryObjectuivEXT, callGlGetQueryObjectuivEXT);
    interpreter->registerFunction(Ids::GlGetQueryObjecti64vEXT, callGlGetQueryObjecti64vEXT);
    interpreter->registerFunction(Ids::GlGetQueryObjectui64vEXT, callGlGetQueryObjectui64vEXT);
}
void Initialize() {
    eglCreateContext =
            reinterpret_cast<PFNEGLCREATECONTEXT>(gapic::GetGfxProcAddress("eglCreateContext"));
    eglMakeCurrent =
            reinterpret_cast<PFNEGLMAKECURRENT>(gapic::GetGfxProcAddress("eglMakeCurrent"));
    eglSwapBuffers =
            reinterpret_cast<PFNEGLSWAPBUFFERS>(gapic::GetGfxProcAddress("eglSwapBuffers"));
    wglSwapBuffers =
            reinterpret_cast<PFNWGLSWAPBUFFERS>(gapic::GetGfxProcAddress("wglSwapBuffers"));
    glEnableClientState = reinterpret_cast<PFNGLENABLECLIENTSTATE>(
            gapic::GetGfxProcAddress("glEnableClientState"));
    glDisableClientState = reinterpret_cast<PFNGLDISABLECLIENTSTATE>(
            gapic::GetGfxProcAddress("glDisableClientState"));
    glGetProgramBinaryOES = reinterpret_cast<PFNGLGETPROGRAMBINARYOES>(
            gapic::GetGfxProcAddress("glGetProgramBinaryOES"));
    glProgramBinaryOES =
            reinterpret_cast<PFNGLPROGRAMBINARYOES>(gapic::GetGfxProcAddress("glProgramBinaryOES"));
    glStartTilingQCOM =
            reinterpret_cast<PFNGLSTARTTILINGQCOM>(gapic::GetGfxProcAddress("glStartTilingQCOM"));
    glEndTilingQCOM =
            reinterpret_cast<PFNGLENDTILINGQCOM>(gapic::GetGfxProcAddress("glEndTilingQCOM"));
    glDiscardFramebufferEXT = reinterpret_cast<PFNGLDISCARDFRAMEBUFFEREXT>(
            gapic::GetGfxProcAddress("glDiscardFramebufferEXT"));
    glInsertEventMarkerEXT = reinterpret_cast<PFNGLINSERTEVENTMARKEREXT>(
            gapic::GetGfxProcAddress("glInsertEventMarkerEXT"));
    glPushGroupMarkerEXT = reinterpret_cast<PFNGLPUSHGROUPMARKEREXT>(
            gapic::GetGfxProcAddress("glPushGroupMarkerEXT"));
    glPopGroupMarkerEXT = reinterpret_cast<PFNGLPOPGROUPMARKEREXT>(
            gapic::GetGfxProcAddress("glPopGroupMarkerEXT"));
    glTexStorage1DEXT =
            reinterpret_cast<PFNGLTEXSTORAGE1DEXT>(gapic::GetGfxProcAddress("glTexStorage1DEXT"));
    glTexStorage2DEXT =
            reinterpret_cast<PFNGLTEXSTORAGE2DEXT>(gapic::GetGfxProcAddress("glTexStorage2DEXT"));
    glTexStorage3DEXT =
            reinterpret_cast<PFNGLTEXSTORAGE3DEXT>(gapic::GetGfxProcAddress("glTexStorage3DEXT"));
    glTextureStorage1DEXT = reinterpret_cast<PFNGLTEXTURESTORAGE1DEXT>(
            gapic::GetGfxProcAddress("glTextureStorage1DEXT"));
    glTextureStorage2DEXT = reinterpret_cast<PFNGLTEXTURESTORAGE2DEXT>(
            gapic::GetGfxProcAddress("glTextureStorage2DEXT"));
    glTextureStorage3DEXT = reinterpret_cast<PFNGLTEXTURESTORAGE3DEXT>(
            gapic::GetGfxProcAddress("glTextureStorage3DEXT"));
    glGenVertexArraysOES = reinterpret_cast<PFNGLGENVERTEXARRAYSOES>(
            gapic::GetGfxProcAddress("glGenVertexArraysOES"));
    glBindVertexArrayOES = reinterpret_cast<PFNGLBINDVERTEXARRAYOES>(
            gapic::GetGfxProcAddress("glBindVertexArrayOES"));
    glDeleteVertexArraysOES = reinterpret_cast<PFNGLDELETEVERTEXARRAYSOES>(
            gapic::GetGfxProcAddress("glDeleteVertexArraysOES"));
    glIsVertexArrayOES =
            reinterpret_cast<PFNGLISVERTEXARRAYOES>(gapic::GetGfxProcAddress("glIsVertexArrayOES"));
    glEGLImageTargetTexture2DOES = reinterpret_cast<PFNGLEGLIMAGETARGETTEXTURE2DOES>(
            gapic::GetGfxProcAddress("glEGLImageTargetTexture2DOES"));
    glEGLImageTargetRenderbufferStorageOES =
            reinterpret_cast<PFNGLEGLIMAGETARGETRENDERBUFFERSTORAGEOES>(
                    gapic::GetGfxProcAddress("glEGLImageTargetRenderbufferStorageOES"));
    glGetGraphicsResetStatusEXT = reinterpret_cast<PFNGLGETGRAPHICSRESETSTATUSEXT>(
            gapic::GetGfxProcAddress("glGetGraphicsResetStatusEXT"));
    glBindAttribLocation = reinterpret_cast<PFNGLBINDATTRIBLOCATION>(
            gapic::GetGfxProcAddress("glBindAttribLocation"));
    glBlendFunc = reinterpret_cast<PFNGLBLENDFUNC>(gapic::GetGfxProcAddress("glBlendFunc"));
    glBlendFuncSeparate = reinterpret_cast<PFNGLBLENDFUNCSEPARATE>(
            gapic::GetGfxProcAddress("glBlendFuncSeparate"));
    glBlendEquation =
            reinterpret_cast<PFNGLBLENDEQUATION>(gapic::GetGfxProcAddress("glBlendEquation"));
    glBlendEquationSeparate = reinterpret_cast<PFNGLBLENDEQUATIONSEPARATE>(
            gapic::GetGfxProcAddress("glBlendEquationSeparate"));
    glBlendColor = reinterpret_cast<PFNGLBLENDCOLOR>(gapic::GetGfxProcAddress("glBlendColor"));
    glEnableVertexAttribArray = reinterpret_cast<PFNGLENABLEVERTEXATTRIBARRAY>(
            gapic::GetGfxProcAddress("glEnableVertexAttribArray"));
    glDisableVertexAttribArray = reinterpret_cast<PFNGLDISABLEVERTEXATTRIBARRAY>(
            gapic::GetGfxProcAddress("glDisableVertexAttribArray"));
    glVertexAttribPointer = reinterpret_cast<PFNGLVERTEXATTRIBPOINTER>(
            gapic::GetGfxProcAddress("glVertexAttribPointer"));
    glGetActiveAttrib =
            reinterpret_cast<PFNGLGETACTIVEATTRIB>(gapic::GetGfxProcAddress("glGetActiveAttrib"));
    glGetActiveUniform =
            reinterpret_cast<PFNGLGETACTIVEUNIFORM>(gapic::GetGfxProcAddress("glGetActiveUniform"));
    glGetError = reinterpret_cast<PFNGLGETERROR>(gapic::GetGfxProcAddress("glGetError"));
    glGetProgramiv =
            reinterpret_cast<PFNGLGETPROGRAMIV>(gapic::GetGfxProcAddress("glGetProgramiv"));
    glGetShaderiv = reinterpret_cast<PFNGLGETSHADERIV>(gapic::GetGfxProcAddress("glGetShaderiv"));
    glGetUniformLocation = reinterpret_cast<PFNGLGETUNIFORMLOCATION>(
            gapic::GetGfxProcAddress("glGetUniformLocation"));
    glGetAttribLocation = reinterpret_cast<PFNGLGETATTRIBLOCATION>(
            gapic::GetGfxProcAddress("glGetAttribLocation"));
    glPixelStorei = reinterpret_cast<PFNGLPIXELSTOREI>(gapic::GetGfxProcAddress("glPixelStorei"));
    glTexParameteri =
            reinterpret_cast<PFNGLTEXPARAMETERI>(gapic::GetGfxProcAddress("glTexParameteri"));
    glTexParameterf =
            reinterpret_cast<PFNGLTEXPARAMETERF>(gapic::GetGfxProcAddress("glTexParameterf"));
    glGetTexParameteriv = reinterpret_cast<PFNGLGETTEXPARAMETERIV>(
            gapic::GetGfxProcAddress("glGetTexParameteriv"));
    glGetTexParameterfv = reinterpret_cast<PFNGLGETTEXPARAMETERFV>(
            gapic::GetGfxProcAddress("glGetTexParameterfv"));
    glUniform1i = reinterpret_cast<PFNGLUNIFORM1I>(gapic::GetGfxProcAddress("glUniform1i"));
    glUniform2i = reinterpret_cast<PFNGLUNIFORM2I>(gapic::GetGfxProcAddress("glUniform2i"));
    glUniform3i = reinterpret_cast<PFNGLUNIFORM3I>(gapic::GetGfxProcAddress("glUniform3i"));
    glUniform4i = reinterpret_cast<PFNGLUNIFORM4I>(gapic::GetGfxProcAddress("glUniform4i"));
    glUniform1iv = reinterpret_cast<PFNGLUNIFORM1IV>(gapic::GetGfxProcAddress("glUniform1iv"));
    glUniform2iv = reinterpret_cast<PFNGLUNIFORM2IV>(gapic::GetGfxProcAddress("glUniform2iv"));
    glUniform3iv = reinterpret_cast<PFNGLUNIFORM3IV>(gapic::GetGfxProcAddress("glUniform3iv"));
    glUniform4iv = reinterpret_cast<PFNGLUNIFORM4IV>(gapic::GetGfxProcAddress("glUniform4iv"));
    glUniform1f = reinterpret_cast<PFNGLUNIFORM1F>(gapic::GetGfxProcAddress("glUniform1f"));
    glUniform2f = reinterpret_cast<PFNGLUNIFORM2F>(gapic::GetGfxProcAddress("glUniform2f"));
    glUniform3f = reinterpret_cast<PFNGLUNIFORM3F>(gapic::GetGfxProcAddress("glUniform3f"));
    glUniform4f = reinterpret_cast<PFNGLUNIFORM4F>(gapic::GetGfxProcAddress("glUniform4f"));
    glUniform1fv = reinterpret_cast<PFNGLUNIFORM1FV>(gapic::GetGfxProcAddress("glUniform1fv"));
    glUniform2fv = reinterpret_cast<PFNGLUNIFORM2FV>(gapic::GetGfxProcAddress("glUniform2fv"));
    glUniform3fv = reinterpret_cast<PFNGLUNIFORM3FV>(gapic::GetGfxProcAddress("glUniform3fv"));
    glUniform4fv = reinterpret_cast<PFNGLUNIFORM4FV>(gapic::GetGfxProcAddress("glUniform4fv"));
    glUniformMatrix2fv =
            reinterpret_cast<PFNGLUNIFORMMATRIX2FV>(gapic::GetGfxProcAddress("glUniformMatrix2fv"));
    glUniformMatrix3fv =
            reinterpret_cast<PFNGLUNIFORMMATRIX3FV>(gapic::GetGfxProcAddress("glUniformMatrix3fv"));
    glUniformMatrix4fv =
            reinterpret_cast<PFNGLUNIFORMMATRIX4FV>(gapic::GetGfxProcAddress("glUniformMatrix4fv"));
    glGetUniformfv =
            reinterpret_cast<PFNGLGETUNIFORMFV>(gapic::GetGfxProcAddress("glGetUniformfv"));
    glGetUniformiv =
            reinterpret_cast<PFNGLGETUNIFORMIV>(gapic::GetGfxProcAddress("glGetUniformiv"));
    glVertexAttrib1f =
            reinterpret_cast<PFNGLVERTEXATTRIB1F>(gapic::GetGfxProcAddress("glVertexAttrib1f"));
    glVertexAttrib2f =
            reinterpret_cast<PFNGLVERTEXATTRIB2F>(gapic::GetGfxProcAddress("glVertexAttrib2f"));
    glVertexAttrib3f =
            reinterpret_cast<PFNGLVERTEXATTRIB3F>(gapic::GetGfxProcAddress("glVertexAttrib3f"));
    glVertexAttrib4f =
            reinterpret_cast<PFNGLVERTEXATTRIB4F>(gapic::GetGfxProcAddress("glVertexAttrib4f"));
    glVertexAttrib1fv =
            reinterpret_cast<PFNGLVERTEXATTRIB1FV>(gapic::GetGfxProcAddress("glVertexAttrib1fv"));
    glVertexAttrib2fv =
            reinterpret_cast<PFNGLVERTEXATTRIB2FV>(gapic::GetGfxProcAddress("glVertexAttrib2fv"));
    glVertexAttrib3fv =
            reinterpret_cast<PFNGLVERTEXATTRIB3FV>(gapic::GetGfxProcAddress("glVertexAttrib3fv"));
    glVertexAttrib4fv =
            reinterpret_cast<PFNGLVERTEXATTRIB4FV>(gapic::GetGfxProcAddress("glVertexAttrib4fv"));
    glGetShaderPrecisionFormat = reinterpret_cast<PFNGLGETSHADERPRECISIONFORMAT>(
            gapic::GetGfxProcAddress("glGetShaderPrecisionFormat"));
    glDepthMask = reinterpret_cast<PFNGLDEPTHMASK>(gapic::GetGfxProcAddress("glDepthMask"));
    glDepthFunc = reinterpret_cast<PFNGLDEPTHFUNC>(gapic::GetGfxProcAddress("glDepthFunc"));
    glDepthRangef = reinterpret_cast<PFNGLDEPTHRANGEF>(gapic::GetGfxProcAddress("glDepthRangef"));
    glColorMask = reinterpret_cast<PFNGLCOLORMASK>(gapic::GetGfxProcAddress("glColorMask"));
    glStencilMask = reinterpret_cast<PFNGLSTENCILMASK>(gapic::GetGfxProcAddress("glStencilMask"));
    glStencilMaskSeparate = reinterpret_cast<PFNGLSTENCILMASKSEPARATE>(
            gapic::GetGfxProcAddress("glStencilMaskSeparate"));
    glStencilFuncSeparate = reinterpret_cast<PFNGLSTENCILFUNCSEPARATE>(
            gapic::GetGfxProcAddress("glStencilFuncSeparate"));
    glStencilOpSeparate = reinterpret_cast<PFNGLSTENCILOPSEPARATE>(
            gapic::GetGfxProcAddress("glStencilOpSeparate"));
    glFrontFace = reinterpret_cast<PFNGLFRONTFACE>(gapic::GetGfxProcAddress("glFrontFace"));
    glViewport = reinterpret_cast<PFNGLVIEWPORT>(gapic::GetGfxProcAddress("glViewport"));
    glScissor = reinterpret_cast<PFNGLSCISSOR>(gapic::GetGfxProcAddress("glScissor"));
    glActiveTexture =
            reinterpret_cast<PFNGLACTIVETEXTURE>(gapic::GetGfxProcAddress("glActiveTexture"));
    glGenTextures = reinterpret_cast<PFNGLGENTEXTURES>(gapic::GetGfxProcAddress("glGenTextures"));
    glDeleteTextures =
            reinterpret_cast<PFNGLDELETETEXTURES>(gapic::GetGfxProcAddress("glDeleteTextures"));
    glIsTexture = reinterpret_cast<PFNGLISTEXTURE>(gapic::GetGfxProcAddress("glIsTexture"));
    glBindTexture = reinterpret_cast<PFNGLBINDTEXTURE>(gapic::GetGfxProcAddress("glBindTexture"));
    glTexImage2D = reinterpret_cast<PFNGLTEXIMAGE2D>(gapic::GetGfxProcAddress("glTexImage2D"));
    glTexSubImage2D =
            reinterpret_cast<PFNGLTEXSUBIMAGE2D>(gapic::GetGfxProcAddress("glTexSubImage2D"));
    glCopyTexImage2D =
            reinterpret_cast<PFNGLCOPYTEXIMAGE2D>(gapic::GetGfxProcAddress("glCopyTexImage2D"));
    glCopyTexSubImage2D = reinterpret_cast<PFNGLCOPYTEXSUBIMAGE2D>(
            gapic::GetGfxProcAddress("glCopyTexSubImage2D"));
    glCompressedTexImage2D = reinterpret_cast<PFNGLCOMPRESSEDTEXIMAGE2D>(
            gapic::GetGfxProcAddress("glCompressedTexImage2D"));
    glCompressedTexSubImage2D = reinterpret_cast<PFNGLCOMPRESSEDTEXSUBIMAGE2D>(
            gapic::GetGfxProcAddress("glCompressedTexSubImage2D"));
    glGenerateMipmap =
            reinterpret_cast<PFNGLGENERATEMIPMAP>(gapic::GetGfxProcAddress("glGenerateMipmap"));
    glReadPixels = reinterpret_cast<PFNGLREADPIXELS>(gapic::GetGfxProcAddress("glReadPixels"));
    glGenFramebuffers =
            reinterpret_cast<PFNGLGENFRAMEBUFFERS>(gapic::GetGfxProcAddress("glGenFramebuffers"));
    glBindFramebuffer =
            reinterpret_cast<PFNGLBINDFRAMEBUFFER>(gapic::GetGfxProcAddress("glBindFramebuffer"));
    glCheckFramebufferStatus = reinterpret_cast<PFNGLCHECKFRAMEBUFFERSTATUS>(
            gapic::GetGfxProcAddress("glCheckFramebufferStatus"));
    glDeleteFramebuffers = reinterpret_cast<PFNGLDELETEFRAMEBUFFERS>(
            gapic::GetGfxProcAddress("glDeleteFramebuffers"));
    glIsFramebuffer =
            reinterpret_cast<PFNGLISFRAMEBUFFER>(gapic::GetGfxProcAddress("glIsFramebuffer"));
    glGenRenderbuffers =
            reinterpret_cast<PFNGLGENRENDERBUFFERS>(gapic::GetGfxProcAddress("glGenRenderbuffers"));
    glBindRenderbuffer =
            reinterpret_cast<PFNGLBINDRENDERBUFFER>(gapic::GetGfxProcAddress("glBindRenderbuffer"));
    glRenderbufferStorage = reinterpret_cast<PFNGLRENDERBUFFERSTORAGE>(
            gapic::GetGfxProcAddress("glRenderbufferStorage"));
    glDeleteRenderbuffers = reinterpret_cast<PFNGLDELETERENDERBUFFERS>(
            gapic::GetGfxProcAddress("glDeleteRenderbuffers"));
    glIsRenderbuffer =
            reinterpret_cast<PFNGLISRENDERBUFFER>(gapic::GetGfxProcAddress("glIsRenderbuffer"));
    glGetRenderbufferParameteriv = reinterpret_cast<PFNGLGETRENDERBUFFERPARAMETERIV>(
            gapic::GetGfxProcAddress("glGetRenderbufferParameteriv"));
    glGenBuffers = reinterpret_cast<PFNGLGENBUFFERS>(gapic::GetGfxProcAddress("glGenBuffers"));
    glBindBuffer = reinterpret_cast<PFNGLBINDBUFFER>(gapic::GetGfxProcAddress("glBindBuffer"));
    glBufferData = reinterpret_cast<PFNGLBUFFERDATA>(gapic::GetGfxProcAddress("glBufferData"));
    glBufferSubData =
            reinterpret_cast<PFNGLBUFFERSUBDATA>(gapic::GetGfxProcAddress("glBufferSubData"));
    glDeleteBuffers =
            reinterpret_cast<PFNGLDELETEBUFFERS>(gapic::GetGfxProcAddress("glDeleteBuffers"));
    glIsBuffer = reinterpret_cast<PFNGLISBUFFER>(gapic::GetGfxProcAddress("glIsBuffer"));
    glGetBufferParameteriv = reinterpret_cast<PFNGLGETBUFFERPARAMETERIV>(
            gapic::GetGfxProcAddress("glGetBufferParameteriv"));
    glCreateShader =
            reinterpret_cast<PFNGLCREATESHADER>(gapic::GetGfxProcAddress("glCreateShader"));
    glDeleteShader =
            reinterpret_cast<PFNGLDELETESHADER>(gapic::GetGfxProcAddress("glDeleteShader"));
    glShaderSource =
            reinterpret_cast<PFNGLSHADERSOURCE>(gapic::GetGfxProcAddress("glShaderSource"));
    glShaderBinary =
            reinterpret_cast<PFNGLSHADERBINARY>(gapic::GetGfxProcAddress("glShaderBinary"));
    glGetShaderInfoLog =
            reinterpret_cast<PFNGLGETSHADERINFOLOG>(gapic::GetGfxProcAddress("glGetShaderInfoLog"));
    glGetShaderSource =
            reinterpret_cast<PFNGLGETSHADERSOURCE>(gapic::GetGfxProcAddress("glGetShaderSource"));
    glReleaseShaderCompiler = reinterpret_cast<PFNGLRELEASESHADERCOMPILER>(
            gapic::GetGfxProcAddress("glReleaseShaderCompiler"));
    glCompileShader =
            reinterpret_cast<PFNGLCOMPILESHADER>(gapic::GetGfxProcAddress("glCompileShader"));
    glIsShader = reinterpret_cast<PFNGLISSHADER>(gapic::GetGfxProcAddress("glIsShader"));
    glCreateProgram =
            reinterpret_cast<PFNGLCREATEPROGRAM>(gapic::GetGfxProcAddress("glCreateProgram"));
    glDeleteProgram =
            reinterpret_cast<PFNGLDELETEPROGRAM>(gapic::GetGfxProcAddress("glDeleteProgram"));
    glAttachShader =
            reinterpret_cast<PFNGLATTACHSHADER>(gapic::GetGfxProcAddress("glAttachShader"));
    glDetachShader =
            reinterpret_cast<PFNGLDETACHSHADER>(gapic::GetGfxProcAddress("glDetachShader"));
    glGetAttachedShaders = reinterpret_cast<PFNGLGETATTACHEDSHADERS>(
            gapic::GetGfxProcAddress("glGetAttachedShaders"));
    glLinkProgram = reinterpret_cast<PFNGLLINKPROGRAM>(gapic::GetGfxProcAddress("glLinkProgram"));
    glGetProgramInfoLog = reinterpret_cast<PFNGLGETPROGRAMINFOLOG>(
            gapic::GetGfxProcAddress("glGetProgramInfoLog"));
    glUseProgram = reinterpret_cast<PFNGLUSEPROGRAM>(gapic::GetGfxProcAddress("glUseProgram"));
    glIsProgram = reinterpret_cast<PFNGLISPROGRAM>(gapic::GetGfxProcAddress("glIsProgram"));
    glValidateProgram =
            reinterpret_cast<PFNGLVALIDATEPROGRAM>(gapic::GetGfxProcAddress("glValidateProgram"));
    glClearColor = reinterpret_cast<PFNGLCLEARCOLOR>(gapic::GetGfxProcAddress("glClearColor"));
    glClearDepthf = reinterpret_cast<PFNGLCLEARDEPTHF>(gapic::GetGfxProcAddress("glClearDepthf"));
    glClearStencil =
            reinterpret_cast<PFNGLCLEARSTENCIL>(gapic::GetGfxProcAddress("glClearStencil"));
    glClear = reinterpret_cast<PFNGLCLEAR>(gapic::GetGfxProcAddress("glClear"));
    glCullFace = reinterpret_cast<PFNGLCULLFACE>(gapic::GetGfxProcAddress("glCullFace"));
    glPolygonOffset =
            reinterpret_cast<PFNGLPOLYGONOFFSET>(gapic::GetGfxProcAddress("glPolygonOffset"));
    glLineWidth = reinterpret_cast<PFNGLLINEWIDTH>(gapic::GetGfxProcAddress("glLineWidth"));
    glSampleCoverage =
            reinterpret_cast<PFNGLSAMPLECOVERAGE>(gapic::GetGfxProcAddress("glSampleCoverage"));
    glHint = reinterpret_cast<PFNGLHINT>(gapic::GetGfxProcAddress("glHint"));
    glFramebufferRenderbuffer = reinterpret_cast<PFNGLFRAMEBUFFERRENDERBUFFER>(
            gapic::GetGfxProcAddress("glFramebufferRenderbuffer"));
    glFramebufferTexture2D = reinterpret_cast<PFNGLFRAMEBUFFERTEXTURE2D>(
            gapic::GetGfxProcAddress("glFramebufferTexture2D"));
    glGetFramebufferAttachmentParameteriv =
            reinterpret_cast<PFNGLGETFRAMEBUFFERATTACHMENTPARAMETERIV>(
                    gapic::GetGfxProcAddress("glGetFramebufferAttachmentParameteriv"));
    glDrawElements =
            reinterpret_cast<PFNGLDRAWELEMENTS>(gapic::GetGfxProcAddress("glDrawElements"));
    glDrawArrays = reinterpret_cast<PFNGLDRAWARRAYS>(gapic::GetGfxProcAddress("glDrawArrays"));
    glFlush = reinterpret_cast<PFNGLFLUSH>(gapic::GetGfxProcAddress("glFlush"));
    glFinish = reinterpret_cast<PFNGLFINISH>(gapic::GetGfxProcAddress("glFinish"));
    glGetBooleanv = reinterpret_cast<PFNGLGETBOOLEANV>(gapic::GetGfxProcAddress("glGetBooleanv"));
    glGetFloatv = reinterpret_cast<PFNGLGETFLOATV>(gapic::GetGfxProcAddress("glGetFloatv"));
    glGetIntegerv = reinterpret_cast<PFNGLGETINTEGERV>(gapic::GetGfxProcAddress("glGetIntegerv"));
    glGetString = reinterpret_cast<PFNGLGETSTRING>(gapic::GetGfxProcAddress("glGetString"));
    glEnable = reinterpret_cast<PFNGLENABLE>(gapic::GetGfxProcAddress("glEnable"));
    glDisable = reinterpret_cast<PFNGLDISABLE>(gapic::GetGfxProcAddress("glDisable"));
    glIsEnabled = reinterpret_cast<PFNGLISENABLED>(gapic::GetGfxProcAddress("glIsEnabled"));
    glMapBufferRange =
            reinterpret_cast<PFNGLMAPBUFFERRANGE>(gapic::GetGfxProcAddress("glMapBufferRange"));
    glUnmapBuffer = reinterpret_cast<PFNGLUNMAPBUFFER>(gapic::GetGfxProcAddress("glUnmapBuffer"));
    glInvalidateFramebuffer = reinterpret_cast<PFNGLINVALIDATEFRAMEBUFFER>(
            gapic::GetGfxProcAddress("glInvalidateFramebuffer"));
    glRenderbufferStorageMultisample = reinterpret_cast<PFNGLRENDERBUFFERSTORAGEMULTISAMPLE>(
            gapic::GetGfxProcAddress("glRenderbufferStorageMultisample"));
    glBlitFramebuffer =
            reinterpret_cast<PFNGLBLITFRAMEBUFFER>(gapic::GetGfxProcAddress("glBlitFramebuffer"));
    glGenQueries = reinterpret_cast<PFNGLGENQUERIES>(gapic::GetGfxProcAddress("glGenQueries"));
    glBeginQuery = reinterpret_cast<PFNGLBEGINQUERY>(gapic::GetGfxProcAddress("glBeginQuery"));
    glEndQuery = reinterpret_cast<PFNGLENDQUERY>(gapic::GetGfxProcAddress("glEndQuery"));
    glDeleteQueries =
            reinterpret_cast<PFNGLDELETEQUERIES>(gapic::GetGfxProcAddress("glDeleteQueries"));
    glIsQuery = reinterpret_cast<PFNGLISQUERY>(gapic::GetGfxProcAddress("glIsQuery"));
    glGetQueryiv = reinterpret_cast<PFNGLGETQUERYIV>(gapic::GetGfxProcAddress("glGetQueryiv"));
    glGetQueryObjectuiv = reinterpret_cast<PFNGLGETQUERYOBJECTUIV>(
            gapic::GetGfxProcAddress("glGetQueryObjectuiv"));
    glGenQueriesEXT =
            reinterpret_cast<PFNGLGENQUERIESEXT>(gapic::GetGfxProcAddress("glGenQueriesEXT"));
    glBeginQueryEXT =
            reinterpret_cast<PFNGLBEGINQUERYEXT>(gapic::GetGfxProcAddress("glBeginQueryEXT"));
    glEndQueryEXT = reinterpret_cast<PFNGLENDQUERYEXT>(gapic::GetGfxProcAddress("glEndQueryEXT"));
    glDeleteQueriesEXT =
            reinterpret_cast<PFNGLDELETEQUERIESEXT>(gapic::GetGfxProcAddress("glDeleteQueriesEXT"));
    glIsQueryEXT = reinterpret_cast<PFNGLISQUERYEXT>(gapic::GetGfxProcAddress("glIsQueryEXT"));
    glQueryCounterEXT =
            reinterpret_cast<PFNGLQUERYCOUNTEREXT>(gapic::GetGfxProcAddress("glQueryCounterEXT"));
    glGetQueryivEXT =
            reinterpret_cast<PFNGLGETQUERYIVEXT>(gapic::GetGfxProcAddress("glGetQueryivEXT"));
    glGetQueryObjectivEXT = reinterpret_cast<PFNGLGETQUERYOBJECTIVEXT>(
            gapic::GetGfxProcAddress("glGetQueryObjectivEXT"));
    glGetQueryObjectuivEXT = reinterpret_cast<PFNGLGETQUERYOBJECTUIVEXT>(
            gapic::GetGfxProcAddress("glGetQueryObjectuivEXT"));
    glGetQueryObjecti64vEXT = reinterpret_cast<PFNGLGETQUERYOBJECTI64VEXT>(
            gapic::GetGfxProcAddress("glGetQueryObjecti64vEXT"));
    glGetQueryObjectui64vEXT = reinterpret_cast<PFNGLGETQUERYOBJECTUI64VEXT>(
            gapic::GetGfxProcAddress("glGetQueryObjectui64vEXT"));
}

}  // namespace gfxapi
}  // namespace gapir
