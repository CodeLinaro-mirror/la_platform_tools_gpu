////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

#include "gfx_api.h"
#include "interpreter.h"
#include "stack.h"

#include <gapic/get_gfx_proc_address.h>
#include <gapic/log.h>

#define __STDC_FORMAT_MACROS
#include <inttypes.h>

namespace gapir {
namespace gfxapi {
namespace {

bool callGlBlendBarrierKHR(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glBlendBarrierKHR()");
        if (glBlendBarrierKHR != nullptr) {
            glBlendBarrierKHR();
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendBarrierKHR");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendBarrierKHR");
        return false;
    }
}

bool callGlBlendEquationSeparateiEXT(Stack* stack, bool pushReturn) {
    GLenum modeAlpha = stack->pop<GLenum>();
    GLenum modeRGB = stack->pop<GLenum>();
    uint32_t buf = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendEquationSeparateiEXT(%" PRIu32 ", %u, %u)", buf, modeRGB, modeAlpha);
        if (glBlendEquationSeparateiEXT != nullptr) {
            glBlendEquationSeparateiEXT(buf, modeRGB, modeAlpha);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendEquationSeparateiEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendEquationSeparateiEXT");
        return false;
    }
}

bool callGlBlendEquationiEXT(Stack* stack, bool pushReturn) {
    GLenum mode = stack->pop<GLenum>();
    uint32_t buf = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendEquationiEXT(%" PRIu32 ", %u)", buf, mode);
        if (glBlendEquationiEXT != nullptr) {
            glBlendEquationiEXT(buf, mode);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendEquationiEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendEquationiEXT");
        return false;
    }
}

bool callGlBlendFuncSeparateiEXT(Stack* stack, bool pushReturn) {
    GLenum dstAlpha = stack->pop<GLenum>();
    GLenum srcAlpha = stack->pop<GLenum>();
    GLenum dstRGB = stack->pop<GLenum>();
    GLenum srcRGB = stack->pop<GLenum>();
    uint32_t buf = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendFuncSeparateiEXT(%" PRIu32 ", %u, %u, %u, %u)", buf, srcRGB, dstRGB,
                   srcAlpha, dstAlpha);
        if (glBlendFuncSeparateiEXT != nullptr) {
            glBlendFuncSeparateiEXT(buf, srcRGB, dstRGB, srcAlpha, dstAlpha);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendFuncSeparateiEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendFuncSeparateiEXT");
        return false;
    }
}

bool callGlBlendFunciEXT(Stack* stack, bool pushReturn) {
    GLenum dst = stack->pop<GLenum>();
    GLenum src = stack->pop<GLenum>();
    uint32_t buf = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendFunciEXT(%" PRIu32 ", %u, %u)", buf, src, dst);
        if (glBlendFunciEXT != nullptr) {
            glBlendFunciEXT(buf, src, dst);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendFunciEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendFunciEXT");
        return false;
    }
}

bool callGlColorMaskiEXT(Stack* stack, bool pushReturn) {
    uint8_t a = stack->pop<uint8_t>();
    uint8_t b = stack->pop<uint8_t>();
    uint8_t g = stack->pop<uint8_t>();
    uint8_t r = stack->pop<uint8_t>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glColorMaskiEXT(%" PRIu32 ", %" PRIu8 ", %" PRIu8 ", %" PRIu8 ", %" PRIu8 ")",
                   index, r, g, b, a);
        if (glColorMaskiEXT != nullptr) {
            glColorMaskiEXT(index, r, g, b, a);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glColorMaskiEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glColorMaskiEXT");
        return false;
    }
}

bool callGlCopyImageSubDataEXT(Stack* stack, bool pushReturn) {
    int32_t srcDepth = stack->pop<int32_t>();
    int32_t srcHeight = stack->pop<int32_t>();
    int32_t srcWidth = stack->pop<int32_t>();
    int32_t dstZ = stack->pop<int32_t>();
    int32_t dstY = stack->pop<int32_t>();
    int32_t dstX = stack->pop<int32_t>();
    int32_t dstLevel = stack->pop<int32_t>();
    GLenum dstTarget = stack->pop<GLenum>();
    uint32_t dstName = stack->pop<uint32_t>();
    int32_t srcZ = stack->pop<int32_t>();
    int32_t srcY = stack->pop<int32_t>();
    int32_t srcX = stack->pop<int32_t>();
    int32_t srcLevel = stack->pop<int32_t>();
    GLenum srcTarget = stack->pop<GLenum>();
    uint32_t srcName = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glCopyImageSubDataEXT(%" PRIu32 ", %u, %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %" PRIu32 ", %u, %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %" PRId32 ", %" PRId32 ")",
                   srcName, srcTarget, srcLevel, srcX, srcY, srcZ, dstName, dstTarget, dstLevel,
                   dstX, dstY, dstZ, srcWidth, srcHeight, srcDepth);
        if (glCopyImageSubDataEXT != nullptr) {
            glCopyImageSubDataEXT(srcName, srcTarget, srcLevel, srcX, srcY, srcZ, dstName,
                                  dstTarget, dstLevel, dstX, dstY, dstZ, srcWidth, srcHeight,
                                  srcDepth);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCopyImageSubDataEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCopyImageSubDataEXT");
        return false;
    }
}

bool callGlDebugMessageCallbackKHR(Stack* stack, bool pushReturn) {
    void* userParam = stack->pop<void*>();
    void* callback = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("glDebugMessageCallbackKHR(%p, %p)", callback, userParam);
        if (glDebugMessageCallbackKHR != nullptr) {
            glDebugMessageCallbackKHR(callback, userParam);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDebugMessageCallbackKHR");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDebugMessageCallbackKHR");
        return false;
    }
}

bool callGlDebugMessageControlKHR(Stack* stack, bool pushReturn) {
    uint8_t enabled = stack->pop<uint8_t>();
    uint32_t* ids = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    GLenum severity = stack->pop<GLenum>();
    GLenum type = stack->pop<GLenum>();
    GLenum source = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDebugMessageControlKHR(%u, %u, %u, %" PRId32 ", %p, %" PRIu8 ")", source,
                   type, severity, count, ids, enabled);
        if (glDebugMessageControlKHR != nullptr) {
            glDebugMessageControlKHR(source, type, severity, count, ids, enabled);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDebugMessageControlKHR");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDebugMessageControlKHR");
        return false;
    }
}

bool callGlDebugMessageInsertKHR(Stack* stack, bool pushReturn) {
    char* message = stack->pop<char*>();
    int32_t length = stack->pop<int32_t>();
    GLenum severity = stack->pop<GLenum>();
    uint32_t id = stack->pop<uint32_t>();
    GLenum type = stack->pop<GLenum>();
    GLenum source = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDebugMessageInsertKHR(%u, %u, %" PRIu32 ", %u, %" PRId32 ", %p)", source,
                   type, id, severity, length, message);
        if (glDebugMessageInsertKHR != nullptr) {
            glDebugMessageInsertKHR(source, type, id, severity, length, message);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDebugMessageInsertKHR");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDebugMessageInsertKHR");
        return false;
    }
}

bool callGlDisableiEXT(Stack* stack, bool pushReturn) {
    uint32_t index = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDisableiEXT(%u, %" PRIu32 ")", target, index);
        if (glDisableiEXT != nullptr) {
            glDisableiEXT(target, index);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDisableiEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDisableiEXT");
        return false;
    }
}

bool callGlEnableiEXT(Stack* stack, bool pushReturn) {
    uint32_t index = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glEnableiEXT(%u, %" PRIu32 ")", target, index);
        if (glEnableiEXT != nullptr) {
            glEnableiEXT(target, index);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEnableiEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEnableiEXT");
        return false;
    }
}

bool callGlFramebufferTextureEXT(Stack* stack, bool pushReturn) {
    int32_t level = stack->pop<int32_t>();
    uint32_t texture = stack->pop<uint32_t>();
    GLenum attachment = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glFramebufferTextureEXT(%u, %u, %" PRIu32 ", %" PRId32 ")", target, attachment,
                   texture, level);
        if (glFramebufferTextureEXT != nullptr) {
            glFramebufferTextureEXT(target, attachment, texture, level);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFramebufferTextureEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFramebufferTextureEXT");
        return false;
    }
}

bool callGlGetDebugMessageLogKHR(Stack* stack, bool pushReturn) {
    char* messageLog = stack->pop<char*>();
    int32_t* lengths = stack->pop<int32_t*>();
    GLenum* severities = stack->pop<GLenum*>();
    uint32_t* ids = stack->pop<uint32_t*>();
    GLenum* types = stack->pop<GLenum*>();
    GLenum* sources = stack->pop<GLenum*>();
    int32_t bufSize = stack->pop<int32_t>();
    uint32_t count = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetDebugMessageLogKHR(%" PRIu32 ", %" PRId32 ", %p, %p, %p, %p, %p, %p)",
                   count, bufSize, sources, types, ids, severities, lengths, messageLog);
        if (glGetDebugMessageLogKHR != nullptr) {
            uint32_t return_value = glGetDebugMessageLogKHR(count, bufSize, sources, types, ids,
                                                            severities, lengths, messageLog);
            GAPID_INFO("Returned: %" PRIu32 "", return_value);
            if (pushReturn) {
                stack->push<uint32_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetDebugMessageLogKHR");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetDebugMessageLogKHR");
        return false;
    }
}

bool callGlGetObjectLabelKHR(Stack* stack, bool pushReturn) {
    char* label = stack->pop<char*>();
    int32_t* length = stack->pop<int32_t*>();
    int32_t bufSize = stack->pop<int32_t>();
    uint32_t name = stack->pop<uint32_t>();
    GLenum identifier = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetObjectLabelKHR(%u, %" PRIu32 ", %" PRId32 ", %p, %p)", identifier, name,
                   bufSize, length, label);
        if (glGetObjectLabelKHR != nullptr) {
            glGetObjectLabelKHR(identifier, name, bufSize, length, label);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetObjectLabelKHR");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetObjectLabelKHR");
        return false;
    }
}

bool callGlGetObjectPtrLabelKHR(Stack* stack, bool pushReturn) {
    char* label = stack->pop<char*>();
    int32_t* length = stack->pop<int32_t*>();
    int32_t bufSize = stack->pop<int32_t>();
    void* ptr = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("glGetObjectPtrLabelKHR(%p, %" PRId32 ", %p, %p)", ptr, bufSize, length, label);
        if (glGetObjectPtrLabelKHR != nullptr) {
            glGetObjectPtrLabelKHR(ptr, bufSize, length, label);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetObjectPtrLabelKHR");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetObjectPtrLabelKHR");
        return false;
    }
}

bool callGlGetPointervKHR(Stack* stack, bool pushReturn) {
    void** params = stack->pop<void**>();
    GLenum pname = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetPointervKHR(%u, %p)", pname, params);
        if (glGetPointervKHR != nullptr) {
            glGetPointervKHR(pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPointervKHR");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPointervKHR");
        return false;
    }
}

bool callGlGetSamplerParameterIivEXT(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetSamplerParameterIivEXT(%" PRIu32 ", %u, %p)", sampler, pname, params);
        if (glGetSamplerParameterIivEXT != nullptr) {
            glGetSamplerParameterIivEXT(sampler, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetSamplerParameterIivEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetSamplerParameterIivEXT");
        return false;
    }
}

bool callGlGetSamplerParameterIuivEXT(Stack* stack, bool pushReturn) {
    uint32_t* params = stack->pop<uint32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetSamplerParameterIuivEXT(%" PRIu32 ", %u, %p)", sampler, pname, params);
        if (glGetSamplerParameterIuivEXT != nullptr) {
            glGetSamplerParameterIuivEXT(sampler, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetSamplerParameterIuivEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetSamplerParameterIuivEXT");
        return false;
    }
}

bool callGlGetTexParameterIivEXT(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetTexParameterIivEXT(%u, %u, %p)", target, pname, params);
        if (glGetTexParameterIivEXT != nullptr) {
            glGetTexParameterIivEXT(target, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetTexParameterIivEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetTexParameterIivEXT");
        return false;
    }
}

bool callGlGetTexParameterIuivEXT(Stack* stack, bool pushReturn) {
    uint32_t* params = stack->pop<uint32_t*>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetTexParameterIuivEXT(%u, %u, %p)", target, pname, params);
        if (glGetTexParameterIuivEXT != nullptr) {
            glGetTexParameterIuivEXT(target, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetTexParameterIuivEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetTexParameterIuivEXT");
        return false;
    }
}

bool callGlIsEnablediEXT(Stack* stack, bool pushReturn) {
    uint32_t index = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glIsEnablediEXT(%u, %" PRIu32 ")", target, index);
        if (glIsEnablediEXT != nullptr) {
            uint8_t return_value = glIsEnablediEXT(target, index);
            GAPID_INFO("Returned: %" PRIu8 "", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsEnablediEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsEnablediEXT");
        return false;
    }
}

bool callGlMinSampleShadingOES(Stack* stack, bool pushReturn) {
    float value = stack->pop<float>();
    if (stack->isValid()) {
        GAPID_INFO("glMinSampleShadingOES(%f)", value);
        if (glMinSampleShadingOES != nullptr) {
            glMinSampleShadingOES(value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMinSampleShadingOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMinSampleShadingOES");
        return false;
    }
}

bool callGlObjectLabelKHR(Stack* stack, bool pushReturn) {
    char* label = stack->pop<char*>();
    int32_t length = stack->pop<int32_t>();
    uint32_t name = stack->pop<uint32_t>();
    GLenum identifier = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glObjectLabelKHR(%u, %" PRIu32 ", %" PRId32 ", %p)", identifier, name, length,
                   label);
        if (glObjectLabelKHR != nullptr) {
            glObjectLabelKHR(identifier, name, length, label);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glObjectLabelKHR");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glObjectLabelKHR");
        return false;
    }
}

bool callGlObjectPtrLabelKHR(Stack* stack, bool pushReturn) {
    char* label = stack->pop<char*>();
    int32_t length = stack->pop<int32_t>();
    void* ptr = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("glObjectPtrLabelKHR(%p, %" PRId32 ", %p)", ptr, length, label);
        if (glObjectPtrLabelKHR != nullptr) {
            glObjectPtrLabelKHR(ptr, length, label);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glObjectPtrLabelKHR");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glObjectPtrLabelKHR");
        return false;
    }
}

bool callGlPatchParameteriEXT(Stack* stack, bool pushReturn) {
    int32_t value = stack->pop<int32_t>();
    GLenum pname = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glPatchParameteriEXT(%u, %" PRId32 ")", pname, value);
        if (glPatchParameteriEXT != nullptr) {
            glPatchParameteriEXT(pname, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPatchParameteriEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPatchParameteriEXT");
        return false;
    }
}

bool callGlPopDebugGroupKHR(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glPopDebugGroupKHR()");
        if (glPopDebugGroupKHR != nullptr) {
            glPopDebugGroupKHR();
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPopDebugGroupKHR");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPopDebugGroupKHR");
        return false;
    }
}

bool callGlPrimitiveBoundingBoxEXT(Stack* stack, bool pushReturn) {
    float maxW = stack->pop<float>();
    float maxZ = stack->pop<float>();
    float maxY = stack->pop<float>();
    float maxX = stack->pop<float>();
    float minW = stack->pop<float>();
    float minZ = stack->pop<float>();
    float minY = stack->pop<float>();
    float minX = stack->pop<float>();
    if (stack->isValid()) {
        GAPID_INFO("glPrimitiveBoundingBoxEXT(%f, %f, %f, %f, %f, %f, %f, %f)", minX, minY, minZ,
                   minW, maxX, maxY, maxZ, maxW);
        if (glPrimitiveBoundingBoxEXT != nullptr) {
            glPrimitiveBoundingBoxEXT(minX, minY, minZ, minW, maxX, maxY, maxZ, maxW);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPrimitiveBoundingBoxEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPrimitiveBoundingBoxEXT");
        return false;
    }
}

bool callGlPushDebugGroupKHR(Stack* stack, bool pushReturn) {
    char* message = stack->pop<char*>();
    int32_t length = stack->pop<int32_t>();
    uint32_t id = stack->pop<uint32_t>();
    GLenum source = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glPushDebugGroupKHR(%u, %" PRIu32 ", %" PRId32 ", %p)", source, id, length,
                   message);
        if (glPushDebugGroupKHR != nullptr) {
            glPushDebugGroupKHR(source, id, length, message);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPushDebugGroupKHR");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPushDebugGroupKHR");
        return false;
    }
}

bool callGlSamplerParameterIivEXT(Stack* stack, bool pushReturn) {
    int32_t* param = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glSamplerParameterIivEXT(%" PRIu32 ", %u, %p)", sampler, pname, param);
        if (glSamplerParameterIivEXT != nullptr) {
            glSamplerParameterIivEXT(sampler, pname, param);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glSamplerParameterIivEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glSamplerParameterIivEXT");
        return false;
    }
}

bool callGlSamplerParameterIuivEXT(Stack* stack, bool pushReturn) {
    uint32_t* param = stack->pop<uint32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glSamplerParameterIuivEXT(%" PRIu32 ", %u, %p)", sampler, pname, param);
        if (glSamplerParameterIuivEXT != nullptr) {
            glSamplerParameterIuivEXT(sampler, pname, param);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glSamplerParameterIuivEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glSamplerParameterIuivEXT");
        return false;
    }
}

bool callGlTexBufferEXT(Stack* stack, bool pushReturn) {
    uint32_t buffer = stack->pop<uint32_t>();
    GLenum internalformat = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexBufferEXT(%u, %u, %" PRIu32 ")", target, internalformat, buffer);
        if (glTexBufferEXT != nullptr) {
            glTexBufferEXT(target, internalformat, buffer);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexBufferEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexBufferEXT");
        return false;
    }
}

bool callGlTexBufferRangeEXT(Stack* stack, bool pushReturn) {
    int32_t size = stack->pop<int32_t>();
    int32_t offset = stack->pop<int32_t>();
    uint32_t buffer = stack->pop<uint32_t>();
    GLenum internalformat = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexBufferRangeEXT(%u, %u, %" PRIu32 ", %" PRId32 ", %" PRId32 ")", target,
                   internalformat, buffer, offset, size);
        if (glTexBufferRangeEXT != nullptr) {
            glTexBufferRangeEXT(target, internalformat, buffer, offset, size);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexBufferRangeEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexBufferRangeEXT");
        return false;
    }
}

bool callGlTexParameterIivEXT(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexParameterIivEXT(%u, %u, %p)", target, pname, params);
        if (glTexParameterIivEXT != nullptr) {
            glTexParameterIivEXT(target, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexParameterIivEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexParameterIivEXT");
        return false;
    }
}

bool callGlTexParameterIuivEXT(Stack* stack, bool pushReturn) {
    uint32_t* params = stack->pop<uint32_t*>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexParameterIuivEXT(%u, %u, %p)", target, pname, params);
        if (glTexParameterIuivEXT != nullptr) {
            glTexParameterIuivEXT(target, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexParameterIuivEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexParameterIuivEXT");
        return false;
    }
}

bool callGlTexStorage3DMultisampleOES(Stack* stack, bool pushReturn) {
    uint8_t fixedsamplelocations = stack->pop<uint8_t>();
    int32_t depth = stack->pop<int32_t>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    GLenum internalformat = stack->pop<GLenum>();
    int32_t samples = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexStorage3DMultisampleOES(%u, %" PRId32 ", %u, %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %" PRIu8 ")",
                   target, samples, internalformat, width, height, depth, fixedsamplelocations);
        if (glTexStorage3DMultisampleOES != nullptr) {
            glTexStorage3DMultisampleOES(target, samples, internalformat, width, height, depth,
                                         fixedsamplelocations);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexStorage3DMultisampleOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexStorage3DMultisampleOES");
        return false;
    }
}

bool callGlBeginQuery(Stack* stack, bool pushReturn) {
    uint32_t query = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBeginQuery(%u, %" PRIu32 ")", target, query);
        if (glBeginQuery != nullptr) {
            glBeginQuery(target, query);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBeginQuery");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBeginQuery");
        return false;
    }
}

bool callGlDeleteQueries(Stack* stack, bool pushReturn) {
    uint32_t* queries = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteQueries(%" PRId32 ", %p)", count, queries);
        if (glDeleteQueries != nullptr) {
            glDeleteQueries(count, queries);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteQueries");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteQueries");
        return false;
    }
}

bool callGlEndQuery(Stack* stack, bool pushReturn) {
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glEndQuery(%u)", target);
        if (glEndQuery != nullptr) {
            glEndQuery(target);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEndQuery");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEndQuery");
        return false;
    }
}

bool callGlGenQueries(Stack* stack, bool pushReturn) {
    uint32_t* queries = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenQueries(%" PRId32 ", %p)", count, queries);
        if (glGenQueries != nullptr) {
            glGenQueries(count, queries);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenQueries");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenQueries");
        return false;
    }
}

bool callGlGetQueryObjectuiv(Stack* stack, bool pushReturn) {
    uint32_t* value = stack->pop<uint32_t*>();
    GLenum parameter = stack->pop<GLenum>();
    uint32_t query = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetQueryObjectuiv(%" PRIu32 ", %u, %p)", query, parameter, value);
        if (glGetQueryObjectuiv != nullptr) {
            glGetQueryObjectuiv(query, parameter, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetQueryObjectuiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetQueryObjectuiv");
        return false;
    }
}

bool callGlGetQueryiv(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    GLenum parameter = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetQueryiv(%u, %u, %p)", target, parameter, value);
        if (glGetQueryiv != nullptr) {
            glGetQueryiv(target, parameter, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetQueryiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetQueryiv");
        return false;
    }
}

bool callGlIsQuery(Stack* stack, bool pushReturn) {
    uint32_t query = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsQuery(%" PRIu32 ")", query);
        if (glIsQuery != nullptr) {
            uint8_t return_value = glIsQuery(query);
            GAPID_INFO("Returned: %" PRIu8 "", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsQuery");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsQuery");
        return false;
    }
}

bool callGlBindBuffer(Stack* stack, bool pushReturn) {
    uint32_t buffer = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBindBuffer(%u, %" PRIu32 ")", target, buffer);
        if (glBindBuffer != nullptr) {
            glBindBuffer(target, buffer);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindBuffer");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindBuffer");
        return false;
    }
}

bool callGlBindBufferBase(Stack* stack, bool pushReturn) {
    uint32_t buffer = stack->pop<uint32_t>();
    uint32_t index = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBindBufferBase(%u, %" PRIu32 ", %" PRIu32 ")", target, index, buffer);
        if (glBindBufferBase != nullptr) {
            glBindBufferBase(target, index, buffer);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindBufferBase");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindBufferBase");
        return false;
    }
}

bool callGlBindBufferRange(Stack* stack, bool pushReturn) {
    int32_t size = stack->pop<int32_t>();
    int32_t offset = stack->pop<int32_t>();
    uint32_t buffer = stack->pop<uint32_t>();
    uint32_t index = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBindBufferRange(%u, %" PRIu32 ", %" PRIu32 ", %" PRId32 ", %" PRId32 ")",
                   target, index, buffer, offset, size);
        if (glBindBufferRange != nullptr) {
            glBindBufferRange(target, index, buffer, offset, size);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindBufferRange");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindBufferRange");
        return false;
    }
}

bool callGlBufferData(Stack* stack, bool pushReturn) {
    GLenum usage = stack->pop<GLenum>();
    void* data = stack->pop<void*>();
    int32_t size = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBufferData(%u, %" PRId32 ", %p, %u)", target, size, data, usage);
        if (glBufferData != nullptr) {
            glBufferData(target, size, data, usage);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBufferData");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBufferData");
        return false;
    }
}

bool callGlBufferSubData(Stack* stack, bool pushReturn) {
    void* data = stack->pop<void*>();
    int32_t size = stack->pop<int32_t>();
    int32_t offset = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBufferSubData(%u, %" PRId32 ", %" PRId32 ", %p)", target, offset, size, data);
        if (glBufferSubData != nullptr) {
            glBufferSubData(target, offset, size, data);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBufferSubData");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBufferSubData");
        return false;
    }
}

bool callGlCopyBufferSubData(Stack* stack, bool pushReturn) {
    int32_t size = stack->pop<int32_t>();
    int32_t writeOffset = stack->pop<int32_t>();
    int32_t readOffset = stack->pop<int32_t>();
    GLenum writeTarget = stack->pop<GLenum>();
    GLenum readTarget = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glCopyBufferSubData(%u, %u, %" PRId32 ", %" PRId32 ", %" PRId32 ")", readTarget,
                   writeTarget, readOffset, writeOffset, size);
        if (glCopyBufferSubData != nullptr) {
            glCopyBufferSubData(readTarget, writeTarget, readOffset, writeOffset, size);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCopyBufferSubData");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCopyBufferSubData");
        return false;
    }
}

bool callGlDeleteBuffers(Stack* stack, bool pushReturn) {
    uint32_t* buffers = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteBuffers(%" PRId32 ", %p)", count, buffers);
        if (glDeleteBuffers != nullptr) {
            glDeleteBuffers(count, buffers);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteBuffers");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteBuffers");
        return false;
    }
}

bool callGlGenBuffers(Stack* stack, bool pushReturn) {
    uint32_t* buffers = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenBuffers(%" PRId32 ", %p)", count, buffers);
        if (glGenBuffers != nullptr) {
            glGenBuffers(count, buffers);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenBuffers");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenBuffers");
        return false;
    }
}

bool callGlGetBufferParameteri64v(Stack* stack, bool pushReturn) {
    int64_t* params = stack->pop<int64_t*>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetBufferParameteri64v(%u, %u, %p)", target, pname, params);
        if (glGetBufferParameteri64v != nullptr) {
            glGetBufferParameteri64v(target, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetBufferParameteri64v");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetBufferParameteri64v");
        return false;
    }
}

bool callGlGetBufferParameteriv(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    GLenum parameter = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetBufferParameteriv(%u, %u, %p)", target, parameter, value);
        if (glGetBufferParameteriv != nullptr) {
            glGetBufferParameteriv(target, parameter, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetBufferParameteriv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetBufferParameteriv");
        return false;
    }
}

bool callGlGetBufferPointerv(Stack* stack, bool pushReturn) {
    void** params = stack->pop<void**>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetBufferPointerv(%u, %u, %p)", target, pname, params);
        if (glGetBufferPointerv != nullptr) {
            glGetBufferPointerv(target, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetBufferPointerv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetBufferPointerv");
        return false;
    }
}

bool callGlIsBuffer(Stack* stack, bool pushReturn) {
    uint32_t buffer = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsBuffer(%" PRIu32 ")", buffer);
        if (glIsBuffer != nullptr) {
            uint8_t return_value = glIsBuffer(buffer);
            GAPID_INFO("Returned: %" PRIu8 "", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsBuffer");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsBuffer");
        return false;
    }
}

bool callGlMapBufferRange(Stack* stack, bool pushReturn) {
    GLbitfield access = stack->pop<GLbitfield>();
    int32_t length = stack->pop<int32_t>();
    int32_t offset = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glMapBufferRange(%u, %" PRId32 ", %" PRId32 ", %u)", target, offset, length,
                   access);
        if (glMapBufferRange != nullptr) {
            void* return_value = glMapBufferRange(target, offset, length, access);
            GAPID_INFO("Returned: %p", return_value);
            if (pushReturn) {
                stack->push<void*>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMapBufferRange");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMapBufferRange");
        return false;
    }
}

bool callGlUnmapBuffer(Stack* stack, bool pushReturn) {
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glUnmapBuffer(%u)", target);
        if (glUnmapBuffer != nullptr) {
            uint8_t return_value = glUnmapBuffer(target);
            GAPID_INFO("Returned: %" PRIu8 "", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUnmapBuffer");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUnmapBuffer");
        return false;
    }
}

bool callGlDebugMessageCallback(Stack* stack, bool pushReturn) {
    void* userParam = stack->pop<void*>();
    void* callback = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("glDebugMessageCallback(%p, %p)", callback, userParam);
        if (glDebugMessageCallback != nullptr) {
            glDebugMessageCallback(callback, userParam);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDebugMessageCallback");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDebugMessageCallback");
        return false;
    }
}

bool callGlDebugMessageControl(Stack* stack, bool pushReturn) {
    uint8_t enabled = stack->pop<uint8_t>();
    uint32_t* ids = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    GLenum severity = stack->pop<GLenum>();
    GLenum type = stack->pop<GLenum>();
    GLenum source = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDebugMessageControl(%u, %u, %u, %" PRId32 ", %p, %" PRIu8 ")", source, type,
                   severity, count, ids, enabled);
        if (glDebugMessageControl != nullptr) {
            glDebugMessageControl(source, type, severity, count, ids, enabled);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDebugMessageControl");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDebugMessageControl");
        return false;
    }
}

bool callGlDebugMessageInsert(Stack* stack, bool pushReturn) {
    char* message = stack->pop<char*>();
    int32_t length = stack->pop<int32_t>();
    GLenum severity = stack->pop<GLenum>();
    uint32_t id = stack->pop<uint32_t>();
    GLenum type = stack->pop<GLenum>();
    GLenum source = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDebugMessageInsert(%u, %u, %" PRIu32 ", %u, %" PRId32 ", %p)", source, type,
                   id, severity, length, message);
        if (glDebugMessageInsert != nullptr) {
            glDebugMessageInsert(source, type, id, severity, length, message);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDebugMessageInsert");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDebugMessageInsert");
        return false;
    }
}

bool callGlGetDebugMessageLog(Stack* stack, bool pushReturn) {
    char* messageLog = stack->pop<char*>();
    int32_t* lengths = stack->pop<int32_t*>();
    GLenum* severities = stack->pop<GLenum*>();
    uint32_t* ids = stack->pop<uint32_t*>();
    GLenum* types = stack->pop<GLenum*>();
    GLenum* sources = stack->pop<GLenum*>();
    int32_t bufSize = stack->pop<int32_t>();
    uint32_t count = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetDebugMessageLog(%" PRIu32 ", %" PRId32 ", %p, %p, %p, %p, %p, %p)", count,
                   bufSize, sources, types, ids, severities, lengths, messageLog);
        if (glGetDebugMessageLog != nullptr) {
            uint32_t return_value = glGetDebugMessageLog(count, bufSize, sources, types, ids,
                                                         severities, lengths, messageLog);
            GAPID_INFO("Returned: %" PRIu32 "", return_value);
            if (pushReturn) {
                stack->push<uint32_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetDebugMessageLog");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetDebugMessageLog");
        return false;
    }
}

bool callGlGetObjectLabel(Stack* stack, bool pushReturn) {
    char* label = stack->pop<char*>();
    int32_t* length = stack->pop<int32_t*>();
    int32_t bufSize = stack->pop<int32_t>();
    uint32_t name = stack->pop<uint32_t>();
    GLenum identifier = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetObjectLabel(%u, %" PRIu32 ", %" PRId32 ", %p, %p)", identifier, name,
                   bufSize, length, label);
        if (glGetObjectLabel != nullptr) {
            glGetObjectLabel(identifier, name, bufSize, length, label);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetObjectLabel");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetObjectLabel");
        return false;
    }
}

bool callGlGetObjectPtrLabel(Stack* stack, bool pushReturn) {
    char* label = stack->pop<char*>();
    int32_t* length = stack->pop<int32_t*>();
    int32_t bufSize = stack->pop<int32_t>();
    void* ptr = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("glGetObjectPtrLabel(%p, %" PRId32 ", %p, %p)", ptr, bufSize, length, label);
        if (glGetObjectPtrLabel != nullptr) {
            glGetObjectPtrLabel(ptr, bufSize, length, label);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetObjectPtrLabel");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetObjectPtrLabel");
        return false;
    }
}

bool callGlGetPointerv(Stack* stack, bool pushReturn) {
    void** params = stack->pop<void**>();
    GLenum pname = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetPointerv(%u, %p)", pname, params);
        if (glGetPointerv != nullptr) {
            glGetPointerv(pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPointerv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPointerv");
        return false;
    }
}

bool callGlObjectLabel(Stack* stack, bool pushReturn) {
    char* label = stack->pop<char*>();
    int32_t length = stack->pop<int32_t>();
    uint32_t name = stack->pop<uint32_t>();
    GLenum identifier = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glObjectLabel(%u, %" PRIu32 ", %" PRId32 ", %p)", identifier, name, length,
                   label);
        if (glObjectLabel != nullptr) {
            glObjectLabel(identifier, name, length, label);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glObjectLabel");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glObjectLabel");
        return false;
    }
}

bool callGlObjectPtrLabel(Stack* stack, bool pushReturn) {
    char* label = stack->pop<char*>();
    int32_t length = stack->pop<int32_t>();
    void* ptr = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("glObjectPtrLabel(%p, %" PRId32 ", %p)", ptr, length, label);
        if (glObjectPtrLabel != nullptr) {
            glObjectPtrLabel(ptr, length, label);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glObjectPtrLabel");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glObjectPtrLabel");
        return false;
    }
}

bool callGlPopDebugGroup(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glPopDebugGroup()");
        if (glPopDebugGroup != nullptr) {
            glPopDebugGroup();
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPopDebugGroup");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPopDebugGroup");
        return false;
    }
}

bool callGlPushDebugGroup(Stack* stack, bool pushReturn) {
    char* message = stack->pop<char*>();
    int32_t length = stack->pop<int32_t>();
    uint32_t id = stack->pop<uint32_t>();
    GLenum source = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glPushDebugGroup(%u, %" PRIu32 ", %" PRId32 ", %p)", source, id, length,
                   message);
        if (glPushDebugGroup != nullptr) {
            glPushDebugGroup(source, id, length, message);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPushDebugGroup");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPushDebugGroup");
        return false;
    }
}

bool callGlDrawArrays(Stack* stack, bool pushReturn) {
    int32_t indices_count = stack->pop<int32_t>();
    int32_t first_index = stack->pop<int32_t>();
    GLenum draw_mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawArrays(%u, %" PRId32 ", %" PRId32 ")", draw_mode, first_index,
                   indices_count);
        if (glDrawArrays != nullptr) {
            glDrawArrays(draw_mode, first_index, indices_count);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawArrays");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawArrays");
        return false;
    }
}

bool callGlDrawArraysIndirect(Stack* stack, bool pushReturn) {
    void* indirect = stack->pop<void*>();
    GLenum draw_mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawArraysIndirect(%u, %p)", draw_mode, indirect);
        if (glDrawArraysIndirect != nullptr) {
            glDrawArraysIndirect(draw_mode, indirect);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawArraysIndirect");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawArraysIndirect");
        return false;
    }
}

bool callGlDrawArraysInstanced(Stack* stack, bool pushReturn) {
    int32_t instance_count = stack->pop<int32_t>();
    int32_t indices_count = stack->pop<int32_t>();
    int32_t first_index = stack->pop<int32_t>();
    GLenum draw_mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawArraysInstanced(%u, %" PRId32 ", %" PRId32 ", %" PRId32 ")", draw_mode,
                   first_index, indices_count, instance_count);
        if (glDrawArraysInstanced != nullptr) {
            glDrawArraysInstanced(draw_mode, first_index, indices_count, instance_count);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawArraysInstanced");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawArraysInstanced");
        return false;
    }
}

bool callGlDrawElements(Stack* stack, bool pushReturn) {
    void* indices = stack->pop<void*>();
    GLenum indices_type = stack->pop<GLenum>();
    int32_t indices_count = stack->pop<int32_t>();
    GLenum draw_mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawElements(%u, %" PRId32 ", %u, %p)", draw_mode, indices_count,
                   indices_type, indices);
        if (glDrawElements != nullptr) {
            glDrawElements(draw_mode, indices_count, indices_type, indices);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawElements");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawElements");
        return false;
    }
}

bool callGlDrawElementsBaseVertex(Stack* stack, bool pushReturn) {
    int32_t base_vertex = stack->pop<int32_t>();
    void* indices = stack->pop<void*>();
    GLenum indices_type = stack->pop<GLenum>();
    int32_t indices_count = stack->pop<int32_t>();
    GLenum draw_mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawElementsBaseVertex(%u, %" PRId32 ", %u, %p, %" PRId32 ")", draw_mode,
                   indices_count, indices_type, indices, base_vertex);
        if (glDrawElementsBaseVertex != nullptr) {
            glDrawElementsBaseVertex(draw_mode, indices_count, indices_type, indices, base_vertex);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawElementsBaseVertex");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawElementsBaseVertex");
        return false;
    }
}

bool callGlDrawElementsIndirect(Stack* stack, bool pushReturn) {
    void* indirect = stack->pop<void*>();
    GLenum indices_type = stack->pop<GLenum>();
    GLenum draw_mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawElementsIndirect(%u, %u, %p)", draw_mode, indices_type, indirect);
        if (glDrawElementsIndirect != nullptr) {
            glDrawElementsIndirect(draw_mode, indices_type, indirect);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawElementsIndirect");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawElementsIndirect");
        return false;
    }
}

bool callGlDrawElementsInstanced(Stack* stack, bool pushReturn) {
    int32_t instance_count = stack->pop<int32_t>();
    void* indices = stack->pop<void*>();
    GLenum indices_type = stack->pop<GLenum>();
    int32_t indices_count = stack->pop<int32_t>();
    GLenum draw_mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawElementsInstanced(%u, %" PRId32 ", %u, %p, %" PRId32 ")", draw_mode,
                   indices_count, indices_type, indices, instance_count);
        if (glDrawElementsInstanced != nullptr) {
            glDrawElementsInstanced(draw_mode, indices_count, indices_type, indices,
                                    instance_count);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawElementsInstanced");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawElementsInstanced");
        return false;
    }
}

bool callGlDrawElementsInstancedBaseVertex(Stack* stack, bool pushReturn) {
    int32_t base_vertex = stack->pop<int32_t>();
    int32_t instance_count = stack->pop<int32_t>();
    void* indices = stack->pop<void*>();
    GLenum indices_type = stack->pop<GLenum>();
    int32_t indices_count = stack->pop<int32_t>();
    GLenum draw_mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawElementsInstancedBaseVertex(%u, %" PRId32 ", %u, %p, %" PRId32
                   ", %" PRId32 ")",
                   draw_mode, indices_count, indices_type, indices, instance_count, base_vertex);
        if (glDrawElementsInstancedBaseVertex != nullptr) {
            glDrawElementsInstancedBaseVertex(draw_mode, indices_count, indices_type, indices,
                                              instance_count, base_vertex);
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glDrawElementsInstancedBaseVertex");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawElementsInstancedBaseVertex");
        return false;
    }
}

bool callGlDrawRangeElements(Stack* stack, bool pushReturn) {
    void* indices = stack->pop<void*>();
    GLenum indices_type = stack->pop<GLenum>();
    int32_t indices_count = stack->pop<int32_t>();
    uint32_t end = stack->pop<uint32_t>();
    uint32_t start = stack->pop<uint32_t>();
    GLenum draw_mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawRangeElements(%u, %" PRIu32 ", %" PRIu32 ", %" PRId32 ", %u, %p)",
                   draw_mode, start, end, indices_count, indices_type, indices);
        if (glDrawRangeElements != nullptr) {
            glDrawRangeElements(draw_mode, start, end, indices_count, indices_type, indices);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawRangeElements");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawRangeElements");
        return false;
    }
}

bool callGlDrawRangeElementsBaseVertex(Stack* stack, bool pushReturn) {
    int32_t base_vertex = stack->pop<int32_t>();
    void* indices = stack->pop<void*>();
    GLenum indices_type = stack->pop<GLenum>();
    int32_t indices_count = stack->pop<int32_t>();
    uint32_t end = stack->pop<uint32_t>();
    uint32_t start = stack->pop<uint32_t>();
    GLenum draw_mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawRangeElementsBaseVertex(%u, %" PRIu32 ", %" PRIu32 ", %" PRId32
                   ", %u, %p, %" PRId32 ")",
                   draw_mode, start, end, indices_count, indices_type, indices, base_vertex);
        if (glDrawRangeElementsBaseVertex != nullptr) {
            glDrawRangeElementsBaseVertex(draw_mode, start, end, indices_count, indices_type,
                                          indices, base_vertex);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawRangeElementsBaseVertex");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawRangeElementsBaseVertex");
        return false;
    }
}

bool callGlPatchParameteri(Stack* stack, bool pushReturn) {
    int32_t value = stack->pop<int32_t>();
    GLenum pname = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glPatchParameteri(%u, %" PRId32 ")", pname, value);
        if (glPatchParameteri != nullptr) {
            glPatchParameteri(pname, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPatchParameteri");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPatchParameteri");
        return false;
    }
}

bool callGlPrimitiveBoundingBox(Stack* stack, bool pushReturn) {
    float maxW = stack->pop<float>();
    float maxZ = stack->pop<float>();
    float maxY = stack->pop<float>();
    float maxX = stack->pop<float>();
    float minW = stack->pop<float>();
    float minZ = stack->pop<float>();
    float minY = stack->pop<float>();
    float minX = stack->pop<float>();
    if (stack->isValid()) {
        GAPID_INFO("glPrimitiveBoundingBox(%f, %f, %f, %f, %f, %f, %f, %f)", minX, minY, minZ, minW,
                   maxX, maxY, maxZ, maxW);
        if (glPrimitiveBoundingBox != nullptr) {
            glPrimitiveBoundingBox(minX, minY, minZ, minW, maxX, maxY, maxZ, maxW);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPrimitiveBoundingBox");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPrimitiveBoundingBox");
        return false;
    }
}

bool callGlActiveShaderProgramEXT(Stack* stack, bool pushReturn) {
    uint32_t program = stack->pop<uint32_t>();
    uint32_t pipeline = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glActiveShaderProgramEXT(%" PRIu32 ", %" PRIu32 ")", pipeline, program);
        if (glActiveShaderProgramEXT != nullptr) {
            glActiveShaderProgramEXT(pipeline, program);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glActiveShaderProgramEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glActiveShaderProgramEXT");
        return false;
    }
}

bool callGlAlphaFuncQCOM(Stack* stack, bool pushReturn) {
    float ref = stack->pop<float>();
    GLenum func = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glAlphaFuncQCOM(%u, %f)", func, ref);
        if (glAlphaFuncQCOM != nullptr) {
            glAlphaFuncQCOM(func, ref);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glAlphaFuncQCOM");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glAlphaFuncQCOM");
        return false;
    }
}

bool callGlApplyFramebufferAttachmentCMAAINTEL(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glApplyFramebufferAttachmentCMAAINTEL()");
        if (glApplyFramebufferAttachmentCMAAINTEL != nullptr) {
            glApplyFramebufferAttachmentCMAAINTEL();
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glApplyFramebufferAttachmentCMAAINTEL");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glApplyFramebufferAttachmentCMAAINTEL");
        return false;
    }
}

bool callGlBeginConditionalRenderNV(Stack* stack, bool pushReturn) {
    GLenum mode = stack->pop<GLenum>();
    uint32_t id = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBeginConditionalRenderNV(%" PRIu32 ", %u)", id, mode);
        if (glBeginConditionalRenderNV != nullptr) {
            glBeginConditionalRenderNV(id, mode);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBeginConditionalRenderNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBeginConditionalRenderNV");
        return false;
    }
}

bool callGlBeginPerfMonitorAMD(Stack* stack, bool pushReturn) {
    uint32_t monitor = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBeginPerfMonitorAMD(%" PRIu32 ")", monitor);
        if (glBeginPerfMonitorAMD != nullptr) {
            glBeginPerfMonitorAMD(monitor);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBeginPerfMonitorAMD");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBeginPerfMonitorAMD");
        return false;
    }
}

bool callGlBeginPerfQueryINTEL(Stack* stack, bool pushReturn) {
    uint32_t queryHandle = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBeginPerfQueryINTEL(%" PRIu32 ")", queryHandle);
        if (glBeginPerfQueryINTEL != nullptr) {
            glBeginPerfQueryINTEL(queryHandle);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBeginPerfQueryINTEL");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBeginPerfQueryINTEL");
        return false;
    }
}

bool callGlBeginQueryEXT(Stack* stack, bool pushReturn) {
    uint32_t query = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBeginQueryEXT(%u, %" PRIu32 ")", target, query);
        if (glBeginQueryEXT != nullptr) {
            glBeginQueryEXT(target, query);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBeginQueryEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBeginQueryEXT");
        return false;
    }
}

bool callGlBindProgramPipelineEXT(Stack* stack, bool pushReturn) {
    uint32_t pipeline = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBindProgramPipelineEXT(%" PRIu32 ")", pipeline);
        if (glBindProgramPipelineEXT != nullptr) {
            glBindProgramPipelineEXT(pipeline);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindProgramPipelineEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindProgramPipelineEXT");
        return false;
    }
}

bool callGlBindVertexArrayOES(Stack* stack, bool pushReturn) {
    uint32_t array = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBindVertexArrayOES(%" PRIu32 ")", array);
        if (glBindVertexArrayOES != nullptr) {
            glBindVertexArrayOES(array);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindVertexArrayOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindVertexArrayOES");
        return false;
    }
}

bool callGlBlendBarrierNV(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glBlendBarrierNV()");
        if (glBlendBarrierNV != nullptr) {
            glBlendBarrierNV();
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendBarrierNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendBarrierNV");
        return false;
    }
}

bool callGlBlendEquationSeparateiOES(Stack* stack, bool pushReturn) {
    GLenum modeAlpha = stack->pop<GLenum>();
    GLenum modeRGB = stack->pop<GLenum>();
    uint32_t buf = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendEquationSeparateiOES(%" PRIu32 ", %u, %u)", buf, modeRGB, modeAlpha);
        if (glBlendEquationSeparateiOES != nullptr) {
            glBlendEquationSeparateiOES(buf, modeRGB, modeAlpha);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendEquationSeparateiOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendEquationSeparateiOES");
        return false;
    }
}

bool callGlBlendEquationiOES(Stack* stack, bool pushReturn) {
    GLenum mode = stack->pop<GLenum>();
    uint32_t buf = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendEquationiOES(%" PRIu32 ", %u)", buf, mode);
        if (glBlendEquationiOES != nullptr) {
            glBlendEquationiOES(buf, mode);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendEquationiOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendEquationiOES");
        return false;
    }
}

bool callGlBlendFuncSeparateiOES(Stack* stack, bool pushReturn) {
    GLenum dstAlpha = stack->pop<GLenum>();
    GLenum srcAlpha = stack->pop<GLenum>();
    GLenum dstRGB = stack->pop<GLenum>();
    GLenum srcRGB = stack->pop<GLenum>();
    uint32_t buf = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendFuncSeparateiOES(%" PRIu32 ", %u, %u, %u, %u)", buf, srcRGB, dstRGB,
                   srcAlpha, dstAlpha);
        if (glBlendFuncSeparateiOES != nullptr) {
            glBlendFuncSeparateiOES(buf, srcRGB, dstRGB, srcAlpha, dstAlpha);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendFuncSeparateiOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendFuncSeparateiOES");
        return false;
    }
}

bool callGlBlendFunciOES(Stack* stack, bool pushReturn) {
    GLenum dst = stack->pop<GLenum>();
    GLenum src = stack->pop<GLenum>();
    uint32_t buf = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendFunciOES(%" PRIu32 ", %u, %u)", buf, src, dst);
        if (glBlendFunciOES != nullptr) {
            glBlendFunciOES(buf, src, dst);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendFunciOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendFunciOES");
        return false;
    }
}

bool callGlBlendParameteriNV(Stack* stack, bool pushReturn) {
    int32_t value = stack->pop<int32_t>();
    GLenum pname = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendParameteriNV(%u, %" PRId32 ")", pname, value);
        if (glBlendParameteriNV != nullptr) {
            glBlendParameteriNV(pname, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendParameteriNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendParameteriNV");
        return false;
    }
}

bool callGlBlitFramebufferANGLE(Stack* stack, bool pushReturn) {
    GLenum filter = stack->pop<GLenum>();
    GLbitfield mask = stack->pop<GLbitfield>();
    int32_t dstY1 = stack->pop<int32_t>();
    int32_t dstX1 = stack->pop<int32_t>();
    int32_t dstY0 = stack->pop<int32_t>();
    int32_t dstX0 = stack->pop<int32_t>();
    int32_t srcY1 = stack->pop<int32_t>();
    int32_t srcX1 = stack->pop<int32_t>();
    int32_t srcY0 = stack->pop<int32_t>();
    int32_t srcX0 = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBlitFramebufferANGLE(%" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %u, %u)",
                   srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter);
        if (glBlitFramebufferANGLE != nullptr) {
            glBlitFramebufferANGLE(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask,
                                   filter);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlitFramebufferANGLE");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlitFramebufferANGLE");
        return false;
    }
}

bool callGlBlitFramebufferNV(Stack* stack, bool pushReturn) {
    GLenum filter = stack->pop<GLenum>();
    GLbitfield mask = stack->pop<GLbitfield>();
    int32_t dstY1 = stack->pop<int32_t>();
    int32_t dstX1 = stack->pop<int32_t>();
    int32_t dstY0 = stack->pop<int32_t>();
    int32_t dstX0 = stack->pop<int32_t>();
    int32_t srcY1 = stack->pop<int32_t>();
    int32_t srcX1 = stack->pop<int32_t>();
    int32_t srcY0 = stack->pop<int32_t>();
    int32_t srcX0 = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBlitFramebufferNV(%" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %u, %u)",
                   srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter);
        if (glBlitFramebufferNV != nullptr) {
            glBlitFramebufferNV(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask,
                                filter);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlitFramebufferNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlitFramebufferNV");
        return false;
    }
}

bool callGlBufferStorageEXT(Stack* stack, bool pushReturn) {
    GLbitfield flag = stack->pop<GLbitfield>();
    void* data = stack->pop<void*>();
    int32_t size = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBufferStorageEXT(%u, %" PRId32 ", %p, %u)", target, size, data, flag);
        if (glBufferStorageEXT != nullptr) {
            glBufferStorageEXT(target, size, data, flag);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBufferStorageEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBufferStorageEXT");
        return false;
    }
}

bool callGlClientWaitSyncAPPLE(Stack* stack, bool pushReturn) {
    uint64_t timeout = stack->pop<uint64_t>();
    GLbitfield flag = stack->pop<GLbitfield>();
    uint64_t sync = stack->pop<uint64_t>();
    if (stack->isValid()) {
        GAPID_INFO("glClientWaitSyncAPPLE(%" PRIu64 ", %u, %" PRIu64 ")", sync, flag, timeout);
        if (glClientWaitSyncAPPLE != nullptr) {
            GLenum return_value = glClientWaitSyncAPPLE(sync, flag, timeout);
            GAPID_INFO("Returned: %u", return_value);
            if (pushReturn) {
                stack->push<GLenum>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glClientWaitSyncAPPLE");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glClientWaitSyncAPPLE");
        return false;
    }
}

bool callGlColorMaskiOES(Stack* stack, bool pushReturn) {
    uint8_t a = stack->pop<uint8_t>();
    uint8_t b = stack->pop<uint8_t>();
    uint8_t g = stack->pop<uint8_t>();
    uint8_t r = stack->pop<uint8_t>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glColorMaskiOES(%" PRIu32 ", %" PRIu8 ", %" PRIu8 ", %" PRIu8 ", %" PRIu8 ")",
                   index, r, g, b, a);
        if (glColorMaskiOES != nullptr) {
            glColorMaskiOES(index, r, g, b, a);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glColorMaskiOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glColorMaskiOES");
        return false;
    }
}

bool callGlCompressedTexImage3DOES(Stack* stack, bool pushReturn) {
    void* data = stack->pop<void*>();
    int32_t imageSize = stack->pop<int32_t>();
    int32_t border = stack->pop<int32_t>();
    int32_t depth = stack->pop<int32_t>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    GLenum internalformat = stack->pop<GLenum>();
    int32_t level = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glCompressedTexImage3DOES(%u, %" PRId32 ", %u, %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %p)",
                   target, level, internalformat, width, height, depth, border, imageSize, data);
        if (glCompressedTexImage3DOES != nullptr) {
            glCompressedTexImage3DOES(target, level, internalformat, width, height, depth, border,
                                      imageSize, data);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCompressedTexImage3DOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCompressedTexImage3DOES");
        return false;
    }
}

bool callGlCompressedTexSubImage3DOES(Stack* stack, bool pushReturn) {
    void* data = stack->pop<void*>();
    int32_t imageSize = stack->pop<int32_t>();
    GLenum format = stack->pop<GLenum>();
    int32_t depth = stack->pop<int32_t>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    int32_t zoffset = stack->pop<int32_t>();
    int32_t yoffset = stack->pop<int32_t>();
    int32_t xoffset = stack->pop<int32_t>();
    int32_t level = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glCompressedTexSubImage3DOES(%u, %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %u, %" PRId32 ", %p)",
                   target, level, xoffset, yoffset, zoffset, width, height, depth, format,
                   imageSize, data);
        if (glCompressedTexSubImage3DOES != nullptr) {
            glCompressedTexSubImage3DOES(target, level, xoffset, yoffset, zoffset, width, height,
                                         depth, format, imageSize, data);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCompressedTexSubImage3DOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCompressedTexSubImage3DOES");
        return false;
    }
}

bool callGlCopyBufferSubDataNV(Stack* stack, bool pushReturn) {
    int32_t size = stack->pop<int32_t>();
    int32_t writeOffset = stack->pop<int32_t>();
    int32_t readOffset = stack->pop<int32_t>();
    GLenum writeTarget = stack->pop<GLenum>();
    GLenum readTarget = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glCopyBufferSubDataNV(%u, %u, %" PRId32 ", %" PRId32 ", %" PRId32 ")",
                   readTarget, writeTarget, readOffset, writeOffset, size);
        if (glCopyBufferSubDataNV != nullptr) {
            glCopyBufferSubDataNV(readTarget, writeTarget, readOffset, writeOffset, size);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCopyBufferSubDataNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCopyBufferSubDataNV");
        return false;
    }
}

bool callGlCopyImageSubDataOES(Stack* stack, bool pushReturn) {
    int32_t srcDepth = stack->pop<int32_t>();
    int32_t srcHeight = stack->pop<int32_t>();
    int32_t srcWidth = stack->pop<int32_t>();
    int32_t dstZ = stack->pop<int32_t>();
    int32_t dstY = stack->pop<int32_t>();
    int32_t dstX = stack->pop<int32_t>();
    int32_t dstLevel = stack->pop<int32_t>();
    GLenum dstTarget = stack->pop<GLenum>();
    uint32_t dstName = stack->pop<uint32_t>();
    int32_t srcZ = stack->pop<int32_t>();
    int32_t srcY = stack->pop<int32_t>();
    int32_t srcX = stack->pop<int32_t>();
    int32_t srcLevel = stack->pop<int32_t>();
    GLenum srcTarget = stack->pop<GLenum>();
    uint32_t srcName = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glCopyImageSubDataOES(%" PRIu32 ", %u, %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %" PRIu32 ", %u, %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %" PRId32 ", %" PRId32 ")",
                   srcName, srcTarget, srcLevel, srcX, srcY, srcZ, dstName, dstTarget, dstLevel,
                   dstX, dstY, dstZ, srcWidth, srcHeight, srcDepth);
        if (glCopyImageSubDataOES != nullptr) {
            glCopyImageSubDataOES(srcName, srcTarget, srcLevel, srcX, srcY, srcZ, dstName,
                                  dstTarget, dstLevel, dstX, dstY, dstZ, srcWidth, srcHeight,
                                  srcDepth);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCopyImageSubDataOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCopyImageSubDataOES");
        return false;
    }
}

bool callGlCopyPathNV(Stack* stack, bool pushReturn) {
    uint32_t srcPath = stack->pop<uint32_t>();
    uint32_t resultPath = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glCopyPathNV(%" PRIu32 ", %" PRIu32 ")", resultPath, srcPath);
        if (glCopyPathNV != nullptr) {
            glCopyPathNV(resultPath, srcPath);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCopyPathNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCopyPathNV");
        return false;
    }
}

bool callGlCopyTexSubImage3DOES(Stack* stack, bool pushReturn) {
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    int32_t y = stack->pop<int32_t>();
    int32_t x = stack->pop<int32_t>();
    int32_t zoffset = stack->pop<int32_t>();
    int32_t yoffset = stack->pop<int32_t>();
    int32_t xoffset = stack->pop<int32_t>();
    int32_t level = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glCopyTexSubImage3DOES(%u, %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ")",
                   target, level, xoffset, yoffset, zoffset, x, y, width, height);
        if (glCopyTexSubImage3DOES != nullptr) {
            glCopyTexSubImage3DOES(target, level, xoffset, yoffset, zoffset, x, y, width, height);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCopyTexSubImage3DOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCopyTexSubImage3DOES");
        return false;
    }
}

bool callGlCopyTextureLevelsAPPLE(Stack* stack, bool pushReturn) {
    int32_t sourceLevelCount = stack->pop<int32_t>();
    int32_t sourceBaseLevel = stack->pop<int32_t>();
    uint32_t sourceTexture = stack->pop<uint32_t>();
    uint32_t destinationTexture = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glCopyTextureLevelsAPPLE(%" PRIu32 ", %" PRIu32 ", %" PRId32 ", %" PRId32 ")",
                   destinationTexture, sourceTexture, sourceBaseLevel, sourceLevelCount);
        if (glCopyTextureLevelsAPPLE != nullptr) {
            glCopyTextureLevelsAPPLE(destinationTexture, sourceTexture, sourceBaseLevel,
                                     sourceLevelCount);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCopyTextureLevelsAPPLE");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCopyTextureLevelsAPPLE");
        return false;
    }
}

bool callGlCoverFillPathInstancedNV(Stack* stack, bool pushReturn) {
    float* transformValues = stack->pop<float*>();
    GLenum transformType = stack->pop<GLenum>();
    GLenum coverMode = stack->pop<GLenum>();
    uint32_t pathBase = stack->pop<uint32_t>();
    void* paths = stack->pop<void*>();
    GLenum pathNameType = stack->pop<GLenum>();
    int32_t numPaths = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glCoverFillPathInstancedNV(%" PRId32 ", %u, %p, %" PRIu32 ", %u, %u, %p)",
                   numPaths, pathNameType, paths, pathBase, coverMode, transformType,
                   transformValues);
        if (glCoverFillPathInstancedNV != nullptr) {
            glCoverFillPathInstancedNV(numPaths, pathNameType, paths, pathBase, coverMode,
                                       transformType, transformValues);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCoverFillPathInstancedNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCoverFillPathInstancedNV");
        return false;
    }
}

bool callGlCoverFillPathNV(Stack* stack, bool pushReturn) {
    GLenum coverMode = stack->pop<GLenum>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glCoverFillPathNV(%" PRIu32 ", %u)", path, coverMode);
        if (glCoverFillPathNV != nullptr) {
            glCoverFillPathNV(path, coverMode);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCoverFillPathNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCoverFillPathNV");
        return false;
    }
}

bool callGlCoverStrokePathInstancedNV(Stack* stack, bool pushReturn) {
    float* transformValues = stack->pop<float*>();
    GLenum transformType = stack->pop<GLenum>();
    GLenum coverMode = stack->pop<GLenum>();
    uint32_t pathBase = stack->pop<uint32_t>();
    void* paths = stack->pop<void*>();
    GLenum pathNameType = stack->pop<GLenum>();
    int32_t numPaths = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glCoverStrokePathInstancedNV(%" PRId32 ", %u, %p, %" PRIu32 ", %u, %u, %p)",
                   numPaths, pathNameType, paths, pathBase, coverMode, transformType,
                   transformValues);
        if (glCoverStrokePathInstancedNV != nullptr) {
            glCoverStrokePathInstancedNV(numPaths, pathNameType, paths, pathBase, coverMode,
                                         transformType, transformValues);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCoverStrokePathInstancedNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCoverStrokePathInstancedNV");
        return false;
    }
}

bool callGlCoverStrokePathNV(Stack* stack, bool pushReturn) {
    GLenum coverMode = stack->pop<GLenum>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glCoverStrokePathNV(%" PRIu32 ", %u)", path, coverMode);
        if (glCoverStrokePathNV != nullptr) {
            glCoverStrokePathNV(path, coverMode);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCoverStrokePathNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCoverStrokePathNV");
        return false;
    }
}

bool callGlCoverageMaskNV(Stack* stack, bool pushReturn) {
    uint8_t mask = stack->pop<uint8_t>();
    if (stack->isValid()) {
        GAPID_INFO("glCoverageMaskNV(%" PRIu8 ")", mask);
        if (glCoverageMaskNV != nullptr) {
            glCoverageMaskNV(mask);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCoverageMaskNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCoverageMaskNV");
        return false;
    }
}

bool callGlCoverageModulationNV(Stack* stack, bool pushReturn) {
    GLenum components = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glCoverageModulationNV(%u)", components);
        if (glCoverageModulationNV != nullptr) {
            glCoverageModulationNV(components);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCoverageModulationNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCoverageModulationNV");
        return false;
    }
}

bool callGlCoverageModulationTableNV(Stack* stack, bool pushReturn) {
    float* v = stack->pop<float*>();
    int32_t n = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glCoverageModulationTableNV(%" PRId32 ", %p)", n, v);
        if (glCoverageModulationTableNV != nullptr) {
            glCoverageModulationTableNV(n, v);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCoverageModulationTableNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCoverageModulationTableNV");
        return false;
    }
}

bool callGlCoverageOperationNV(Stack* stack, bool pushReturn) {
    GLenum operation = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glCoverageOperationNV(%u)", operation);
        if (glCoverageOperationNV != nullptr) {
            glCoverageOperationNV(operation);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCoverageOperationNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCoverageOperationNV");
        return false;
    }
}

bool callGlCreatePerfQueryINTEL(Stack* stack, bool pushReturn) {
    uint32_t* queryHandle = stack->pop<uint32_t*>();
    uint32_t queryId = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glCreatePerfQueryINTEL(%" PRIu32 ", %p)", queryId, queryHandle);
        if (glCreatePerfQueryINTEL != nullptr) {
            glCreatePerfQueryINTEL(queryId, queryHandle);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCreatePerfQueryINTEL");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCreatePerfQueryINTEL");
        return false;
    }
}

bool callGlCreateShaderProgramvEXT(Stack* stack, bool pushReturn) {
    char** strings = stack->pop<char**>();
    int32_t count = stack->pop<int32_t>();
    GLenum type = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glCreateShaderProgramvEXT(%u, %" PRId32 ", %p)", type, count, strings);
        if (glCreateShaderProgramvEXT != nullptr) {
            uint32_t return_value = glCreateShaderProgramvEXT(type, count, strings);
            GAPID_INFO("Returned: %" PRIu32 "", return_value);
            if (pushReturn) {
                stack->push<uint32_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCreateShaderProgramvEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCreateShaderProgramvEXT");
        return false;
    }
}

bool callGlDeleteFencesNV(Stack* stack, bool pushReturn) {
    uint32_t* fences = stack->pop<uint32_t*>();
    int32_t n = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteFencesNV(%" PRId32 ", %p)", n, fences);
        if (glDeleteFencesNV != nullptr) {
            glDeleteFencesNV(n, fences);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteFencesNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteFencesNV");
        return false;
    }
}

bool callGlDeletePathsNV(Stack* stack, bool pushReturn) {
    int32_t range = stack->pop<int32_t>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeletePathsNV(%" PRIu32 ", %" PRId32 ")", path, range);
        if (glDeletePathsNV != nullptr) {
            glDeletePathsNV(path, range);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeletePathsNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeletePathsNV");
        return false;
    }
}

bool callGlDeletePerfMonitorsAMD(Stack* stack, bool pushReturn) {
    uint32_t* monitors = stack->pop<uint32_t*>();
    int32_t n = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeletePerfMonitorsAMD(%" PRId32 ", %p)", n, monitors);
        if (glDeletePerfMonitorsAMD != nullptr) {
            glDeletePerfMonitorsAMD(n, monitors);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeletePerfMonitorsAMD");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeletePerfMonitorsAMD");
        return false;
    }
}

bool callGlDeletePerfQueryINTEL(Stack* stack, bool pushReturn) {
    uint32_t queryHandle = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeletePerfQueryINTEL(%" PRIu32 ")", queryHandle);
        if (glDeletePerfQueryINTEL != nullptr) {
            glDeletePerfQueryINTEL(queryHandle);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeletePerfQueryINTEL");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeletePerfQueryINTEL");
        return false;
    }
}

bool callGlDeleteProgramPipelinesEXT(Stack* stack, bool pushReturn) {
    uint32_t* pipelines = stack->pop<uint32_t*>();
    int32_t n = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteProgramPipelinesEXT(%" PRId32 ", %p)", n, pipelines);
        if (glDeleteProgramPipelinesEXT != nullptr) {
            glDeleteProgramPipelinesEXT(n, pipelines);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteProgramPipelinesEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteProgramPipelinesEXT");
        return false;
    }
}

bool callGlDeleteQueriesEXT(Stack* stack, bool pushReturn) {
    uint32_t* queries = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteQueriesEXT(%" PRId32 ", %p)", count, queries);
        if (glDeleteQueriesEXT != nullptr) {
            glDeleteQueriesEXT(count, queries);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteQueriesEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteQueriesEXT");
        return false;
    }
}

bool callGlDeleteSyncAPPLE(Stack* stack, bool pushReturn) {
    uint64_t sync = stack->pop<uint64_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteSyncAPPLE(%" PRIu64 ")", sync);
        if (glDeleteSyncAPPLE != nullptr) {
            glDeleteSyncAPPLE(sync);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteSyncAPPLE");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteSyncAPPLE");
        return false;
    }
}

bool callGlDeleteVertexArraysOES(Stack* stack, bool pushReturn) {
    uint32_t* arrays = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteVertexArraysOES(%" PRId32 ", %p)", count, arrays);
        if (glDeleteVertexArraysOES != nullptr) {
            glDeleteVertexArraysOES(count, arrays);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteVertexArraysOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteVertexArraysOES");
        return false;
    }
}

bool callGlDepthRangeArrayfvNV(Stack* stack, bool pushReturn) {
    float* v = stack->pop<float*>();
    int32_t count = stack->pop<int32_t>();
    uint32_t first = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDepthRangeArrayfvNV(%" PRIu32 ", %" PRId32 ", %p)", first, count, v);
        if (glDepthRangeArrayfvNV != nullptr) {
            glDepthRangeArrayfvNV(first, count, v);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDepthRangeArrayfvNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDepthRangeArrayfvNV");
        return false;
    }
}

bool callGlDepthRangeIndexedfNV(Stack* stack, bool pushReturn) {
    float f = stack->pop<float>();
    float n = stack->pop<float>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDepthRangeIndexedfNV(%" PRIu32 ", %f, %f)", index, n, f);
        if (glDepthRangeIndexedfNV != nullptr) {
            glDepthRangeIndexedfNV(index, n, f);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDepthRangeIndexedfNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDepthRangeIndexedfNV");
        return false;
    }
}

bool callGlDisableDriverControlQCOM(Stack* stack, bool pushReturn) {
    uint32_t driverControl = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDisableDriverControlQCOM(%" PRIu32 ")", driverControl);
        if (glDisableDriverControlQCOM != nullptr) {
            glDisableDriverControlQCOM(driverControl);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDisableDriverControlQCOM");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDisableDriverControlQCOM");
        return false;
    }
}

bool callGlDisableiNV(Stack* stack, bool pushReturn) {
    uint32_t index = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDisableiNV(%u, %" PRIu32 ")", target, index);
        if (glDisableiNV != nullptr) {
            glDisableiNV(target, index);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDisableiNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDisableiNV");
        return false;
    }
}

bool callGlDisableiOES(Stack* stack, bool pushReturn) {
    uint32_t index = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDisableiOES(%u, %" PRIu32 ")", target, index);
        if (glDisableiOES != nullptr) {
            glDisableiOES(target, index);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDisableiOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDisableiOES");
        return false;
    }
}

bool callGlDiscardFramebufferEXT(Stack* stack, bool pushReturn) {
    GLenum* attachments = stack->pop<GLenum*>();
    int32_t numAttachments = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDiscardFramebufferEXT(%u, %" PRId32 ", %p)", target, numAttachments,
                   attachments);
        if (glDiscardFramebufferEXT != nullptr) {
            glDiscardFramebufferEXT(target, numAttachments, attachments);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDiscardFramebufferEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDiscardFramebufferEXT");
        return false;
    }
}

bool callGlDrawArraysInstancedANGLE(Stack* stack, bool pushReturn) {
    int32_t primcount = stack->pop<int32_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t first = stack->pop<int32_t>();
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawArraysInstancedANGLE(%u, %" PRId32 ", %" PRId32 ", %" PRId32 ")", mode,
                   first, count, primcount);
        if (glDrawArraysInstancedANGLE != nullptr) {
            glDrawArraysInstancedANGLE(mode, first, count, primcount);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawArraysInstancedANGLE");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawArraysInstancedANGLE");
        return false;
    }
}

bool callGlDrawArraysInstancedBaseInstanceEXT(Stack* stack, bool pushReturn) {
    uint32_t baseinstance = stack->pop<uint32_t>();
    int32_t instancecount = stack->pop<int32_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t first = stack->pop<int32_t>();
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawArraysInstancedBaseInstanceEXT(%u, %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRIu32 ")",
                   mode, first, count, instancecount, baseinstance);
        if (glDrawArraysInstancedBaseInstanceEXT != nullptr) {
            glDrawArraysInstancedBaseInstanceEXT(mode, first, count, instancecount, baseinstance);
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glDrawArraysInstancedBaseInstanceEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawArraysInstancedBaseInstanceEXT");
        return false;
    }
}

bool callGlDrawArraysInstancedEXT(Stack* stack, bool pushReturn) {
    int32_t primcount = stack->pop<int32_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t start = stack->pop<int32_t>();
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawArraysInstancedEXT(%u, %" PRId32 ", %" PRId32 ", %" PRId32 ")", mode,
                   start, count, primcount);
        if (glDrawArraysInstancedEXT != nullptr) {
            glDrawArraysInstancedEXT(mode, start, count, primcount);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawArraysInstancedEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawArraysInstancedEXT");
        return false;
    }
}

bool callGlDrawArraysInstancedNV(Stack* stack, bool pushReturn) {
    int32_t primcount = stack->pop<int32_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t first = stack->pop<int32_t>();
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawArraysInstancedNV(%u, %" PRId32 ", %" PRId32 ", %" PRId32 ")", mode,
                   first, count, primcount);
        if (glDrawArraysInstancedNV != nullptr) {
            glDrawArraysInstancedNV(mode, first, count, primcount);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawArraysInstancedNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawArraysInstancedNV");
        return false;
    }
}

bool callGlDrawBuffersEXT(Stack* stack, bool pushReturn) {
    GLenum* bufs = stack->pop<GLenum*>();
    int32_t n = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawBuffersEXT(%" PRId32 ", %p)", n, bufs);
        if (glDrawBuffersEXT != nullptr) {
            glDrawBuffersEXT(n, bufs);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawBuffersEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawBuffersEXT");
        return false;
    }
}

bool callGlDrawBuffersIndexedEXT(Stack* stack, bool pushReturn) {
    int32_t* indices = stack->pop<int32_t*>();
    GLenum* location = stack->pop<GLenum*>();
    int32_t n = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawBuffersIndexedEXT(%" PRId32 ", %p, %p)", n, location, indices);
        if (glDrawBuffersIndexedEXT != nullptr) {
            glDrawBuffersIndexedEXT(n, location, indices);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawBuffersIndexedEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawBuffersIndexedEXT");
        return false;
    }
}

bool callGlDrawBuffersNV(Stack* stack, bool pushReturn) {
    GLenum* bufs = stack->pop<GLenum*>();
    int32_t n = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawBuffersNV(%" PRId32 ", %p)", n, bufs);
        if (glDrawBuffersNV != nullptr) {
            glDrawBuffersNV(n, bufs);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawBuffersNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawBuffersNV");
        return false;
    }
}

bool callGlDrawElementsBaseVertexEXT(Stack* stack, bool pushReturn) {
    int32_t basevertex = stack->pop<int32_t>();
    void* indices = stack->pop<void*>();
    GLenum type = stack->pop<GLenum>();
    int32_t count = stack->pop<int32_t>();
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawElementsBaseVertexEXT(%u, %" PRId32 ", %u, %p, %" PRId32 ")", mode, count,
                   type, indices, basevertex);
        if (glDrawElementsBaseVertexEXT != nullptr) {
            glDrawElementsBaseVertexEXT(mode, count, type, indices, basevertex);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawElementsBaseVertexEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawElementsBaseVertexEXT");
        return false;
    }
}

bool callGlDrawElementsBaseVertexOES(Stack* stack, bool pushReturn) {
    int32_t basevertex = stack->pop<int32_t>();
    void* indices = stack->pop<void*>();
    GLenum type = stack->pop<GLenum>();
    int32_t count = stack->pop<int32_t>();
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawElementsBaseVertexOES(%u, %" PRId32 ", %u, %p, %" PRId32 ")", mode, count,
                   type, indices, basevertex);
        if (glDrawElementsBaseVertexOES != nullptr) {
            glDrawElementsBaseVertexOES(mode, count, type, indices, basevertex);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawElementsBaseVertexOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawElementsBaseVertexOES");
        return false;
    }
}

bool callGlDrawElementsInstancedANGLE(Stack* stack, bool pushReturn) {
    int32_t primcount = stack->pop<int32_t>();
    void* indices = stack->pop<void*>();
    GLenum type = stack->pop<GLenum>();
    int32_t count = stack->pop<int32_t>();
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawElementsInstancedANGLE(%u, %" PRId32 ", %u, %p, %" PRId32 ")", mode,
                   count, type, indices, primcount);
        if (glDrawElementsInstancedANGLE != nullptr) {
            glDrawElementsInstancedANGLE(mode, count, type, indices, primcount);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawElementsInstancedANGLE");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawElementsInstancedANGLE");
        return false;
    }
}

bool callGlDrawElementsInstancedBaseInstanceEXT(Stack* stack, bool pushReturn) {
    uint32_t baseinstance = stack->pop<uint32_t>();
    int32_t instancecount = stack->pop<int32_t>();
    void* indices = stack->pop<void*>();
    GLenum type = stack->pop<GLenum>();
    int32_t count = stack->pop<int32_t>();
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawElementsInstancedBaseInstanceEXT(%u, %" PRId32 ", %u, %p, %" PRId32
                   ", %" PRIu32 ")",
                   mode, count, type, indices, instancecount, baseinstance);
        if (glDrawElementsInstancedBaseInstanceEXT != nullptr) {
            glDrawElementsInstancedBaseInstanceEXT(mode, count, type, indices, instancecount,
                                                   baseinstance);
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function "
                    "glDrawElementsInstancedBaseInstanceEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawElementsInstancedBaseInstanceEXT");
        return false;
    }
}

bool callGlDrawElementsInstancedBaseVertexBaseInstanceEXT(Stack* stack, bool pushReturn) {
    uint32_t baseinstance = stack->pop<uint32_t>();
    int32_t basevertex = stack->pop<int32_t>();
    int32_t instancecount = stack->pop<int32_t>();
    void* indices = stack->pop<void*>();
    GLenum type = stack->pop<GLenum>();
    int32_t count = stack->pop<int32_t>();
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawElementsInstancedBaseVertexBaseInstanceEXT(%u, %" PRId32
                   ", %u, %p, %" PRId32 ", %" PRId32 ", %" PRIu32 ")",
                   mode, count, type, indices, instancecount, basevertex, baseinstance);
        if (glDrawElementsInstancedBaseVertexBaseInstanceEXT != nullptr) {
            glDrawElementsInstancedBaseVertexBaseInstanceEXT(
                    mode, count, type, indices, instancecount, basevertex, baseinstance);
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function "
                    "glDrawElementsInstancedBaseVertexBaseInstanceEXT");
        }
        return true;
    } else {
        GAPID_WARNING(
                "Error during calling function glDrawElementsInstancedBaseVertexBaseInstanceEXT");
        return false;
    }
}

bool callGlDrawElementsInstancedBaseVertexEXT(Stack* stack, bool pushReturn) {
    int32_t basevertex = stack->pop<int32_t>();
    int32_t instancecount = stack->pop<int32_t>();
    void* indices = stack->pop<void*>();
    GLenum type = stack->pop<GLenum>();
    int32_t count = stack->pop<int32_t>();
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawElementsInstancedBaseVertexEXT(%u, %" PRId32 ", %u, %p, %" PRId32
                   ", %" PRId32 ")",
                   mode, count, type, indices, instancecount, basevertex);
        if (glDrawElementsInstancedBaseVertexEXT != nullptr) {
            glDrawElementsInstancedBaseVertexEXT(mode, count, type, indices, instancecount,
                                                 basevertex);
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glDrawElementsInstancedBaseVertexEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawElementsInstancedBaseVertexEXT");
        return false;
    }
}

bool callGlDrawElementsInstancedBaseVertexOES(Stack* stack, bool pushReturn) {
    int32_t basevertex = stack->pop<int32_t>();
    int32_t instancecount = stack->pop<int32_t>();
    void* indices = stack->pop<void*>();
    GLenum type = stack->pop<GLenum>();
    int32_t count = stack->pop<int32_t>();
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawElementsInstancedBaseVertexOES(%u, %" PRId32 ", %u, %p, %" PRId32
                   ", %" PRId32 ")",
                   mode, count, type, indices, instancecount, basevertex);
        if (glDrawElementsInstancedBaseVertexOES != nullptr) {
            glDrawElementsInstancedBaseVertexOES(mode, count, type, indices, instancecount,
                                                 basevertex);
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glDrawElementsInstancedBaseVertexOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawElementsInstancedBaseVertexOES");
        return false;
    }
}

bool callGlDrawElementsInstancedEXT(Stack* stack, bool pushReturn) {
    int32_t primcount = stack->pop<int32_t>();
    void* indices = stack->pop<void*>();
    GLenum type = stack->pop<GLenum>();
    int32_t count = stack->pop<int32_t>();
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawElementsInstancedEXT(%u, %" PRId32 ", %u, %p, %" PRId32 ")", mode, count,
                   type, indices, primcount);
        if (glDrawElementsInstancedEXT != nullptr) {
            glDrawElementsInstancedEXT(mode, count, type, indices, primcount);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawElementsInstancedEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawElementsInstancedEXT");
        return false;
    }
}

bool callGlDrawElementsInstancedNV(Stack* stack, bool pushReturn) {
    int32_t primcount = stack->pop<int32_t>();
    void* indices = stack->pop<void*>();
    GLenum type = stack->pop<GLenum>();
    int32_t count = stack->pop<int32_t>();
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawElementsInstancedNV(%u, %" PRId32 ", %u, %p, %" PRId32 ")", mode, count,
                   type, indices, primcount);
        if (glDrawElementsInstancedNV != nullptr) {
            glDrawElementsInstancedNV(mode, count, type, indices, primcount);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawElementsInstancedNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawElementsInstancedNV");
        return false;
    }
}

bool callGlDrawRangeElementsBaseVertexEXT(Stack* stack, bool pushReturn) {
    int32_t basevertex = stack->pop<int32_t>();
    void* indices = stack->pop<void*>();
    GLenum type = stack->pop<GLenum>();
    int32_t count = stack->pop<int32_t>();
    uint32_t end = stack->pop<uint32_t>();
    uint32_t start = stack->pop<uint32_t>();
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawRangeElementsBaseVertexEXT(%u, %" PRIu32 ", %" PRIu32 ", %" PRId32
                   ", %u, %p, %" PRId32 ")",
                   mode, start, end, count, type, indices, basevertex);
        if (glDrawRangeElementsBaseVertexEXT != nullptr) {
            glDrawRangeElementsBaseVertexEXT(mode, start, end, count, type, indices, basevertex);
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glDrawRangeElementsBaseVertexEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawRangeElementsBaseVertexEXT");
        return false;
    }
}

bool callGlDrawRangeElementsBaseVertexOES(Stack* stack, bool pushReturn) {
    int32_t basevertex = stack->pop<int32_t>();
    void* indices = stack->pop<void*>();
    GLenum type = stack->pop<GLenum>();
    int32_t count = stack->pop<int32_t>();
    uint32_t end = stack->pop<uint32_t>();
    uint32_t start = stack->pop<uint32_t>();
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawRangeElementsBaseVertexOES(%u, %" PRIu32 ", %" PRIu32 ", %" PRId32
                   ", %u, %p, %" PRId32 ")",
                   mode, start, end, count, type, indices, basevertex);
        if (glDrawRangeElementsBaseVertexOES != nullptr) {
            glDrawRangeElementsBaseVertexOES(mode, start, end, count, type, indices, basevertex);
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glDrawRangeElementsBaseVertexOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawRangeElementsBaseVertexOES");
        return false;
    }
}

bool callGlEGLImageTargetRenderbufferStorageOES(Stack* stack, bool pushReturn) {
    void* image = stack->pop<void*>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glEGLImageTargetRenderbufferStorageOES(%u, %p)", target, image);
        if (glEGLImageTargetRenderbufferStorageOES != nullptr) {
            glEGLImageTargetRenderbufferStorageOES(target, image);
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function "
                    "glEGLImageTargetRenderbufferStorageOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEGLImageTargetRenderbufferStorageOES");
        return false;
    }
}

bool callGlEGLImageTargetTexture2DOES(Stack* stack, bool pushReturn) {
    void* image = stack->pop<void*>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glEGLImageTargetTexture2DOES(%u, %p)", target, image);
        if (glEGLImageTargetTexture2DOES != nullptr) {
            glEGLImageTargetTexture2DOES(target, image);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEGLImageTargetTexture2DOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEGLImageTargetTexture2DOES");
        return false;
    }
}

bool callGlEnableDriverControlQCOM(Stack* stack, bool pushReturn) {
    uint32_t driverControl = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glEnableDriverControlQCOM(%" PRIu32 ")", driverControl);
        if (glEnableDriverControlQCOM != nullptr) {
            glEnableDriverControlQCOM(driverControl);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEnableDriverControlQCOM");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEnableDriverControlQCOM");
        return false;
    }
}

bool callGlEnableiNV(Stack* stack, bool pushReturn) {
    uint32_t index = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glEnableiNV(%u, %" PRIu32 ")", target, index);
        if (glEnableiNV != nullptr) {
            glEnableiNV(target, index);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEnableiNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEnableiNV");
        return false;
    }
}

bool callGlEnableiOES(Stack* stack, bool pushReturn) {
    uint32_t index = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glEnableiOES(%u, %" PRIu32 ")", target, index);
        if (glEnableiOES != nullptr) {
            glEnableiOES(target, index);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEnableiOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEnableiOES");
        return false;
    }
}

bool callGlEndConditionalRenderNV(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glEndConditionalRenderNV()");
        if (glEndConditionalRenderNV != nullptr) {
            glEndConditionalRenderNV();
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEndConditionalRenderNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEndConditionalRenderNV");
        return false;
    }
}

bool callGlEndPerfMonitorAMD(Stack* stack, bool pushReturn) {
    uint32_t monitor = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glEndPerfMonitorAMD(%" PRIu32 ")", monitor);
        if (glEndPerfMonitorAMD != nullptr) {
            glEndPerfMonitorAMD(monitor);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEndPerfMonitorAMD");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEndPerfMonitorAMD");
        return false;
    }
}

bool callGlEndPerfQueryINTEL(Stack* stack, bool pushReturn) {
    uint32_t queryHandle = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glEndPerfQueryINTEL(%" PRIu32 ")", queryHandle);
        if (glEndPerfQueryINTEL != nullptr) {
            glEndPerfQueryINTEL(queryHandle);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEndPerfQueryINTEL");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEndPerfQueryINTEL");
        return false;
    }
}

bool callGlEndQueryEXT(Stack* stack, bool pushReturn) {
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glEndQueryEXT(%u)", target);
        if (glEndQueryEXT != nullptr) {
            glEndQueryEXT(target);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEndQueryEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEndQueryEXT");
        return false;
    }
}

bool callGlEndTilingQCOM(Stack* stack, bool pushReturn) {
    GLbitfield preserve_mask = stack->pop<GLbitfield>();
    if (stack->isValid()) {
        GAPID_INFO("glEndTilingQCOM(%u)", preserve_mask);
        if (glEndTilingQCOM != nullptr) {
            glEndTilingQCOM(preserve_mask);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEndTilingQCOM");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEndTilingQCOM");
        return false;
    }
}

bool callGlExtGetBufferPointervQCOM(Stack* stack, bool pushReturn) {
    void** params = stack->pop<void**>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glExtGetBufferPointervQCOM(%u, %p)", target, params);
        if (glExtGetBufferPointervQCOM != nullptr) {
            glExtGetBufferPointervQCOM(target, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glExtGetBufferPointervQCOM");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glExtGetBufferPointervQCOM");
        return false;
    }
}

bool callGlExtGetBuffersQCOM(Stack* stack, bool pushReturn) {
    int32_t* numBuffers = stack->pop<int32_t*>();
    int32_t maxBuffers = stack->pop<int32_t>();
    uint32_t* buffers = stack->pop<uint32_t*>();
    if (stack->isValid()) {
        GAPID_INFO("glExtGetBuffersQCOM(%p, %" PRId32 ", %p)", buffers, maxBuffers, numBuffers);
        if (glExtGetBuffersQCOM != nullptr) {
            glExtGetBuffersQCOM(buffers, maxBuffers, numBuffers);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glExtGetBuffersQCOM");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glExtGetBuffersQCOM");
        return false;
    }
}

bool callGlExtGetFramebuffersQCOM(Stack* stack, bool pushReturn) {
    int32_t* numFramebuffers = stack->pop<int32_t*>();
    int32_t maxFramebuffers = stack->pop<int32_t>();
    uint32_t* framebuffers = stack->pop<uint32_t*>();
    if (stack->isValid()) {
        GAPID_INFO("glExtGetFramebuffersQCOM(%p, %" PRId32 ", %p)", framebuffers, maxFramebuffers,
                   numFramebuffers);
        if (glExtGetFramebuffersQCOM != nullptr) {
            glExtGetFramebuffersQCOM(framebuffers, maxFramebuffers, numFramebuffers);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glExtGetFramebuffersQCOM");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glExtGetFramebuffersQCOM");
        return false;
    }
}

bool callGlExtGetProgramBinarySourceQCOM(Stack* stack, bool pushReturn) {
    int32_t* length = stack->pop<int32_t*>();
    char* source = stack->pop<char*>();
    GLenum shadertype = stack->pop<GLenum>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glExtGetProgramBinarySourceQCOM(%" PRIu32 ", %u, %p, %p)", program, shadertype,
                   source, length);
        if (glExtGetProgramBinarySourceQCOM != nullptr) {
            glExtGetProgramBinarySourceQCOM(program, shadertype, source, length);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glExtGetProgramBinarySourceQCOM");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glExtGetProgramBinarySourceQCOM");
        return false;
    }
}

bool callGlExtGetProgramsQCOM(Stack* stack, bool pushReturn) {
    int32_t* numPrograms = stack->pop<int32_t*>();
    int32_t maxPrograms = stack->pop<int32_t>();
    uint32_t* programs = stack->pop<uint32_t*>();
    if (stack->isValid()) {
        GAPID_INFO("glExtGetProgramsQCOM(%p, %" PRId32 ", %p)", programs, maxPrograms, numPrograms);
        if (glExtGetProgramsQCOM != nullptr) {
            glExtGetProgramsQCOM(programs, maxPrograms, numPrograms);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glExtGetProgramsQCOM");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glExtGetProgramsQCOM");
        return false;
    }
}

bool callGlExtGetRenderbuffersQCOM(Stack* stack, bool pushReturn) {
    int32_t* numRenderbuffers = stack->pop<int32_t*>();
    int32_t maxRenderbuffers = stack->pop<int32_t>();
    uint32_t* renderbuffers = stack->pop<uint32_t*>();
    if (stack->isValid()) {
        GAPID_INFO("glExtGetRenderbuffersQCOM(%p, %" PRId32 ", %p)", renderbuffers,
                   maxRenderbuffers, numRenderbuffers);
        if (glExtGetRenderbuffersQCOM != nullptr) {
            glExtGetRenderbuffersQCOM(renderbuffers, maxRenderbuffers, numRenderbuffers);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glExtGetRenderbuffersQCOM");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glExtGetRenderbuffersQCOM");
        return false;
    }
}

bool callGlExtGetShadersQCOM(Stack* stack, bool pushReturn) {
    int32_t* numShaders = stack->pop<int32_t*>();
    int32_t maxShaders = stack->pop<int32_t>();
    uint32_t* shaders = stack->pop<uint32_t*>();
    if (stack->isValid()) {
        GAPID_INFO("glExtGetShadersQCOM(%p, %" PRId32 ", %p)", shaders, maxShaders, numShaders);
        if (glExtGetShadersQCOM != nullptr) {
            glExtGetShadersQCOM(shaders, maxShaders, numShaders);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glExtGetShadersQCOM");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glExtGetShadersQCOM");
        return false;
    }
}

bool callGlExtGetTexLevelParameterivQCOM(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    int32_t level = stack->pop<int32_t>();
    GLenum face = stack->pop<GLenum>();
    uint32_t texture = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glExtGetTexLevelParameterivQCOM(%" PRIu32 ", %u, %" PRId32 ", %u, %p)", texture,
                   face, level, pname, params);
        if (glExtGetTexLevelParameterivQCOM != nullptr) {
            glExtGetTexLevelParameterivQCOM(texture, face, level, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glExtGetTexLevelParameterivQCOM");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glExtGetTexLevelParameterivQCOM");
        return false;
    }
}

bool callGlExtGetTexSubImageQCOM(Stack* stack, bool pushReturn) {
    void* texels = stack->pop<void*>();
    GLenum type = stack->pop<GLenum>();
    GLenum format = stack->pop<GLenum>();
    int32_t depth = stack->pop<int32_t>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    int32_t zoffset = stack->pop<int32_t>();
    int32_t yoffset = stack->pop<int32_t>();
    int32_t xoffset = stack->pop<int32_t>();
    int32_t level = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glExtGetTexSubImageQCOM(%u, %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %u, %u, %p)",
                   target, level, xoffset, yoffset, zoffset, width, height, depth, format, type,
                   texels);
        if (glExtGetTexSubImageQCOM != nullptr) {
            glExtGetTexSubImageQCOM(target, level, xoffset, yoffset, zoffset, width, height, depth,
                                    format, type, texels);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glExtGetTexSubImageQCOM");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glExtGetTexSubImageQCOM");
        return false;
    }
}

bool callGlExtGetTexturesQCOM(Stack* stack, bool pushReturn) {
    int32_t* numTextures = stack->pop<int32_t*>();
    int32_t maxTextures = stack->pop<int32_t>();
    uint32_t* textures = stack->pop<uint32_t*>();
    if (stack->isValid()) {
        GAPID_INFO("glExtGetTexturesQCOM(%p, %" PRId32 ", %p)", textures, maxTextures, numTextures);
        if (glExtGetTexturesQCOM != nullptr) {
            glExtGetTexturesQCOM(textures, maxTextures, numTextures);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glExtGetTexturesQCOM");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glExtGetTexturesQCOM");
        return false;
    }
}

bool callGlExtIsProgramBinaryQCOM(Stack* stack, bool pushReturn) {
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glExtIsProgramBinaryQCOM(%" PRIu32 ")", program);
        if (glExtIsProgramBinaryQCOM != nullptr) {
            uint8_t return_value = glExtIsProgramBinaryQCOM(program);
            GAPID_INFO("Returned: %" PRIu8 "", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glExtIsProgramBinaryQCOM");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glExtIsProgramBinaryQCOM");
        return false;
    }
}

bool callGlExtTexObjectStateOverrideiQCOM(Stack* stack, bool pushReturn) {
    int32_t param = stack->pop<int32_t>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glExtTexObjectStateOverrideiQCOM(%u, %u, %" PRId32 ")", target, pname, param);
        if (glExtTexObjectStateOverrideiQCOM != nullptr) {
            glExtTexObjectStateOverrideiQCOM(target, pname, param);
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glExtTexObjectStateOverrideiQCOM");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glExtTexObjectStateOverrideiQCOM");
        return false;
    }
}

bool callGlFenceSyncAPPLE(Stack* stack, bool pushReturn) {
    GLbitfield flag = stack->pop<GLbitfield>();
    GLenum condition = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glFenceSyncAPPLE(%u, %u)", condition, flag);
        if (glFenceSyncAPPLE != nullptr) {
            uint64_t return_value = glFenceSyncAPPLE(condition, flag);
            GAPID_INFO("Returned: %" PRIu64 "", return_value);
            if (pushReturn) {
                stack->push<uint64_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFenceSyncAPPLE");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFenceSyncAPPLE");
        return false;
    }
}

bool callGlFinishFenceNV(Stack* stack, bool pushReturn) {
    uint32_t fence = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glFinishFenceNV(%" PRIu32 ")", fence);
        if (glFinishFenceNV != nullptr) {
            glFinishFenceNV(fence);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFinishFenceNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFinishFenceNV");
        return false;
    }
}

bool callGlFlushMappedBufferRangeEXT(Stack* stack, bool pushReturn) {
    int32_t length = stack->pop<int32_t>();
    int32_t offset = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glFlushMappedBufferRangeEXT(%u, %" PRId32 ", %" PRId32 ")", target, offset,
                   length);
        if (glFlushMappedBufferRangeEXT != nullptr) {
            glFlushMappedBufferRangeEXT(target, offset, length);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFlushMappedBufferRangeEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFlushMappedBufferRangeEXT");
        return false;
    }
}

bool callGlFragmentCoverageColorNV(Stack* stack, bool pushReturn) {
    uint32_t color = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glFragmentCoverageColorNV(%" PRIu32 ")", color);
        if (glFragmentCoverageColorNV != nullptr) {
            glFragmentCoverageColorNV(color);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFragmentCoverageColorNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFragmentCoverageColorNV");
        return false;
    }
}

bool callGlFramebufferSampleLocationsfvNV(Stack* stack, bool pushReturn) {
    float* v = stack->pop<float*>();
    int32_t count = stack->pop<int32_t>();
    uint32_t start = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glFramebufferSampleLocationsfvNV(%u, %" PRIu32 ", %" PRId32 ", %p)", target,
                   start, count, v);
        if (glFramebufferSampleLocationsfvNV != nullptr) {
            glFramebufferSampleLocationsfvNV(target, start, count, v);
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glFramebufferSampleLocationsfvNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFramebufferSampleLocationsfvNV");
        return false;
    }
}

bool callGlFramebufferTexture2DMultisampleEXT(Stack* stack, bool pushReturn) {
    int32_t samples = stack->pop<int32_t>();
    int32_t level = stack->pop<int32_t>();
    uint32_t texture = stack->pop<uint32_t>();
    GLenum textarget = stack->pop<GLenum>();
    GLenum attachment = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glFramebufferTexture2DMultisampleEXT(%u, %u, %u, %" PRIu32 ", %" PRId32
                   ", %" PRId32 ")",
                   target, attachment, textarget, texture, level, samples);
        if (glFramebufferTexture2DMultisampleEXT != nullptr) {
            glFramebufferTexture2DMultisampleEXT(target, attachment, textarget, texture, level,
                                                 samples);
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glFramebufferTexture2DMultisampleEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFramebufferTexture2DMultisampleEXT");
        return false;
    }
}

bool callGlFramebufferTexture2DMultisampleIMG(Stack* stack, bool pushReturn) {
    int32_t samples = stack->pop<int32_t>();
    int32_t level = stack->pop<int32_t>();
    uint32_t texture = stack->pop<uint32_t>();
    GLenum textarget = stack->pop<GLenum>();
    GLenum attachment = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glFramebufferTexture2DMultisampleIMG(%u, %u, %u, %" PRIu32 ", %" PRId32
                   ", %" PRId32 ")",
                   target, attachment, textarget, texture, level, samples);
        if (glFramebufferTexture2DMultisampleIMG != nullptr) {
            glFramebufferTexture2DMultisampleIMG(target, attachment, textarget, texture, level,
                                                 samples);
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glFramebufferTexture2DMultisampleIMG");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFramebufferTexture2DMultisampleIMG");
        return false;
    }
}

bool callGlFramebufferTexture3DOES(Stack* stack, bool pushReturn) {
    int32_t zoffset = stack->pop<int32_t>();
    int32_t level = stack->pop<int32_t>();
    uint32_t texture = stack->pop<uint32_t>();
    GLenum textarget = stack->pop<GLenum>();
    GLenum attachment = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glFramebufferTexture3DOES(%u, %u, %u, %" PRIu32 ", %" PRId32 ", %" PRId32 ")",
                   target, attachment, textarget, texture, level, zoffset);
        if (glFramebufferTexture3DOES != nullptr) {
            glFramebufferTexture3DOES(target, attachment, textarget, texture, level, zoffset);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFramebufferTexture3DOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFramebufferTexture3DOES");
        return false;
    }
}

bool callGlFramebufferTextureMultiviewOVR(Stack* stack, bool pushReturn) {
    int32_t numViews = stack->pop<int32_t>();
    int32_t baseViewIndex = stack->pop<int32_t>();
    int32_t level = stack->pop<int32_t>();
    uint32_t texture = stack->pop<uint32_t>();
    GLenum attachment = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glFramebufferTextureMultiviewOVR(%u, %u, %" PRIu32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ")",
                   target, attachment, texture, level, baseViewIndex, numViews);
        if (glFramebufferTextureMultiviewOVR != nullptr) {
            glFramebufferTextureMultiviewOVR(target, attachment, texture, level, baseViewIndex,
                                             numViews);
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glFramebufferTextureMultiviewOVR");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFramebufferTextureMultiviewOVR");
        return false;
    }
}

bool callGlFramebufferTextureOES(Stack* stack, bool pushReturn) {
    int32_t level = stack->pop<int32_t>();
    uint32_t texture = stack->pop<uint32_t>();
    GLenum attachment = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glFramebufferTextureOES(%u, %u, %" PRIu32 ", %" PRId32 ")", target, attachment,
                   texture, level);
        if (glFramebufferTextureOES != nullptr) {
            glFramebufferTextureOES(target, attachment, texture, level);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFramebufferTextureOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFramebufferTextureOES");
        return false;
    }
}

bool callGlGenFencesNV(Stack* stack, bool pushReturn) {
    uint32_t* fences = stack->pop<uint32_t*>();
    int32_t n = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenFencesNV(%" PRId32 ", %p)", n, fences);
        if (glGenFencesNV != nullptr) {
            glGenFencesNV(n, fences);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenFencesNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenFencesNV");
        return false;
    }
}

bool callGlGenPathsNV(Stack* stack, bool pushReturn) {
    int32_t range = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenPathsNV(%" PRId32 ")", range);
        if (glGenPathsNV != nullptr) {
            uint32_t return_value = glGenPathsNV(range);
            GAPID_INFO("Returned: %" PRIu32 "", return_value);
            if (pushReturn) {
                stack->push<uint32_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenPathsNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenPathsNV");
        return false;
    }
}

bool callGlGenPerfMonitorsAMD(Stack* stack, bool pushReturn) {
    uint32_t* monitors = stack->pop<uint32_t*>();
    int32_t n = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenPerfMonitorsAMD(%" PRId32 ", %p)", n, monitors);
        if (glGenPerfMonitorsAMD != nullptr) {
            glGenPerfMonitorsAMD(n, monitors);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenPerfMonitorsAMD");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenPerfMonitorsAMD");
        return false;
    }
}

bool callGlGenProgramPipelinesEXT(Stack* stack, bool pushReturn) {
    uint32_t* pipelines = stack->pop<uint32_t*>();
    int32_t n = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenProgramPipelinesEXT(%" PRId32 ", %p)", n, pipelines);
        if (glGenProgramPipelinesEXT != nullptr) {
            glGenProgramPipelinesEXT(n, pipelines);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenProgramPipelinesEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenProgramPipelinesEXT");
        return false;
    }
}

bool callGlGenQueriesEXT(Stack* stack, bool pushReturn) {
    uint32_t* queries = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenQueriesEXT(%" PRId32 ", %p)", count, queries);
        if (glGenQueriesEXT != nullptr) {
            glGenQueriesEXT(count, queries);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenQueriesEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenQueriesEXT");
        return false;
    }
}

bool callGlGenVertexArraysOES(Stack* stack, bool pushReturn) {
    uint32_t* arrays = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenVertexArraysOES(%" PRId32 ", %p)", count, arrays);
        if (glGenVertexArraysOES != nullptr) {
            glGenVertexArraysOES(count, arrays);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenVertexArraysOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenVertexArraysOES");
        return false;
    }
}

bool callGlGetBufferPointervOES(Stack* stack, bool pushReturn) {
    void** params = stack->pop<void**>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetBufferPointervOES(%u, %u, %p)", target, pname, params);
        if (glGetBufferPointervOES != nullptr) {
            glGetBufferPointervOES(target, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetBufferPointervOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetBufferPointervOES");
        return false;
    }
}

bool callGlGetCoverageModulationTableNV(Stack* stack, bool pushReturn) {
    float* v = stack->pop<float*>();
    int32_t bufsize = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetCoverageModulationTableNV(%" PRId32 ", %p)", bufsize, v);
        if (glGetCoverageModulationTableNV != nullptr) {
            glGetCoverageModulationTableNV(bufsize, v);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetCoverageModulationTableNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetCoverageModulationTableNV");
        return false;
    }
}

bool callGlGetDriverControlStringQCOM(Stack* stack, bool pushReturn) {
    char* driverControlString = stack->pop<char*>();
    int32_t* length = stack->pop<int32_t*>();
    int32_t bufSize = stack->pop<int32_t>();
    uint32_t driverControl = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetDriverControlStringQCOM(%" PRIu32 ", %" PRId32 ", %p, %p)", driverControl,
                   bufSize, length, driverControlString);
        if (glGetDriverControlStringQCOM != nullptr) {
            glGetDriverControlStringQCOM(driverControl, bufSize, length, driverControlString);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetDriverControlStringQCOM");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetDriverControlStringQCOM");
        return false;
    }
}

bool callGlGetDriverControlsQCOM(Stack* stack, bool pushReturn) {
    uint32_t* driverControls = stack->pop<uint32_t*>();
    int32_t size = stack->pop<int32_t>();
    int32_t* num = stack->pop<int32_t*>();
    if (stack->isValid()) {
        GAPID_INFO("glGetDriverControlsQCOM(%p, %" PRId32 ", %p)", num, size, driverControls);
        if (glGetDriverControlsQCOM != nullptr) {
            glGetDriverControlsQCOM(num, size, driverControls);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetDriverControlsQCOM");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetDriverControlsQCOM");
        return false;
    }
}

bool callGlGetFenceivNV(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t fence = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetFenceivNV(%" PRIu32 ", %u, %p)", fence, pname, params);
        if (glGetFenceivNV != nullptr) {
            glGetFenceivNV(fence, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetFenceivNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetFenceivNV");
        return false;
    }
}

bool callGlGetFirstPerfQueryIdINTEL(Stack* stack, bool pushReturn) {
    uint32_t* queryId = stack->pop<uint32_t*>();
    if (stack->isValid()) {
        GAPID_INFO("glGetFirstPerfQueryIdINTEL(%p)", queryId);
        if (glGetFirstPerfQueryIdINTEL != nullptr) {
            glGetFirstPerfQueryIdINTEL(queryId);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetFirstPerfQueryIdINTEL");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetFirstPerfQueryIdINTEL");
        return false;
    }
}

bool callGlGetFloatiVNV(Stack* stack, bool pushReturn) {
    float* data = stack->pop<float*>();
    uint32_t index = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetFloati_vNV(%u, %" PRIu32 ", %p)", target, index, data);
        if (glGetFloati_vNV != nullptr) {
            glGetFloati_vNV(target, index, data);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetFloati_vNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetFloati_vNV");
        return false;
    }
}

bool callGlGetGraphicsResetStatusEXT(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glGetGraphicsResetStatusEXT()");
        if (glGetGraphicsResetStatusEXT != nullptr) {
            GLenum return_value = glGetGraphicsResetStatusEXT();
            GAPID_INFO("Returned: %u", return_value);
            if (pushReturn) {
                stack->push<GLenum>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetGraphicsResetStatusEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetGraphicsResetStatusEXT");
        return false;
    }
}

bool callGlGetGraphicsResetStatusKHR(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glGetGraphicsResetStatusKHR()");
        if (glGetGraphicsResetStatusKHR != nullptr) {
            GLenum return_value = glGetGraphicsResetStatusKHR();
            GAPID_INFO("Returned: %u", return_value);
            if (pushReturn) {
                stack->push<GLenum>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetGraphicsResetStatusKHR");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetGraphicsResetStatusKHR");
        return false;
    }
}

bool callGlGetImageHandleNV(Stack* stack, bool pushReturn) {
    GLenum format = stack->pop<GLenum>();
    int32_t layer = stack->pop<int32_t>();
    uint8_t layered = stack->pop<uint8_t>();
    int32_t level = stack->pop<int32_t>();
    uint32_t texture = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetImageHandleNV(%" PRIu32 ", %" PRId32 ", %" PRIu8 ", %" PRId32 ", %u)",
                   texture, level, layered, layer, format);
        if (glGetImageHandleNV != nullptr) {
            uint64_t return_value = glGetImageHandleNV(texture, level, layered, layer, format);
            GAPID_INFO("Returned: %" PRIu64 "", return_value);
            if (pushReturn) {
                stack->push<uint64_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetImageHandleNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetImageHandleNV");
        return false;
    }
}

bool callGlGetInteger64vAPPLE(Stack* stack, bool pushReturn) {
    int64_t* params = stack->pop<int64_t*>();
    GLenum pname = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetInteger64vAPPLE(%u, %p)", pname, params);
        if (glGetInteger64vAPPLE != nullptr) {
            glGetInteger64vAPPLE(pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetInteger64vAPPLE");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetInteger64vAPPLE");
        return false;
    }
}

bool callGlGetIntegeriVEXT(Stack* stack, bool pushReturn) {
    int32_t* data = stack->pop<int32_t*>();
    uint32_t index = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetIntegeri_vEXT(%u, %" PRIu32 ", %p)", target, index, data);
        if (glGetIntegeri_vEXT != nullptr) {
            glGetIntegeri_vEXT(target, index, data);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetIntegeri_vEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetIntegeri_vEXT");
        return false;
    }
}

bool callGlGetInternalformatSampleivNV(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    int32_t bufSize = stack->pop<int32_t>();
    GLenum pname = stack->pop<GLenum>();
    int32_t samples = stack->pop<int32_t>();
    GLenum internalformat = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetInternalformatSampleivNV(%u, %u, %" PRId32 ", %u, %" PRId32 ", %p)",
                   target, internalformat, samples, pname, bufSize, params);
        if (glGetInternalformatSampleivNV != nullptr) {
            glGetInternalformatSampleivNV(target, internalformat, samples, pname, bufSize, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetInternalformatSampleivNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetInternalformatSampleivNV");
        return false;
    }
}

bool callGlGetNextPerfQueryIdINTEL(Stack* stack, bool pushReturn) {
    uint32_t* nextQueryId = stack->pop<uint32_t*>();
    uint32_t queryId = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetNextPerfQueryIdINTEL(%" PRIu32 ", %p)", queryId, nextQueryId);
        if (glGetNextPerfQueryIdINTEL != nullptr) {
            glGetNextPerfQueryIdINTEL(queryId, nextQueryId);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetNextPerfQueryIdINTEL");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetNextPerfQueryIdINTEL");
        return false;
    }
}

bool callGlGetObjectLabelEXT(Stack* stack, bool pushReturn) {
    char* label = stack->pop<char*>();
    int32_t* length = stack->pop<int32_t*>();
    int32_t bufSize = stack->pop<int32_t>();
    uint32_t object = stack->pop<uint32_t>();
    GLenum type = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetObjectLabelEXT(%u, %" PRIu32 ", %" PRId32 ", %p, %p)", type, object,
                   bufSize, length, label);
        if (glGetObjectLabelEXT != nullptr) {
            glGetObjectLabelEXT(type, object, bufSize, length, label);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetObjectLabelEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetObjectLabelEXT");
        return false;
    }
}

bool callGlGetPathCommandsNV(Stack* stack, bool pushReturn) {
    uint8_t* commands = stack->pop<uint8_t*>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetPathCommandsNV(%" PRIu32 ", %p)", path, commands);
        if (glGetPathCommandsNV != nullptr) {
            glGetPathCommandsNV(path, commands);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPathCommandsNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPathCommandsNV");
        return false;
    }
}

bool callGlGetPathCoordsNV(Stack* stack, bool pushReturn) {
    float* coords = stack->pop<float*>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetPathCoordsNV(%" PRIu32 ", %p)", path, coords);
        if (glGetPathCoordsNV != nullptr) {
            glGetPathCoordsNV(path, coords);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPathCoordsNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPathCoordsNV");
        return false;
    }
}

bool callGlGetPathDashArrayNV(Stack* stack, bool pushReturn) {
    float* dashArray = stack->pop<float*>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetPathDashArrayNV(%" PRIu32 ", %p)", path, dashArray);
        if (glGetPathDashArrayNV != nullptr) {
            glGetPathDashArrayNV(path, dashArray);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPathDashArrayNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPathDashArrayNV");
        return false;
    }
}

bool callGlGetPathLengthNV(Stack* stack, bool pushReturn) {
    int32_t numSegments = stack->pop<int32_t>();
    int32_t startSegment = stack->pop<int32_t>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetPathLengthNV(%" PRIu32 ", %" PRId32 ", %" PRId32 ")", path, startSegment,
                   numSegments);
        if (glGetPathLengthNV != nullptr) {
            float return_value = glGetPathLengthNV(path, startSegment, numSegments);
            GAPID_INFO("Returned: %f", return_value);
            if (pushReturn) {
                stack->push<float>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPathLengthNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPathLengthNV");
        return false;
    }
}

bool callGlGetPathMetricRangeNV(Stack* stack, bool pushReturn) {
    float* metrics = stack->pop<float*>();
    int32_t stride = stack->pop<int32_t>();
    int32_t numPaths = stack->pop<int32_t>();
    uint32_t firstPathName = stack->pop<uint32_t>();
    GLbitfield metricQueryMask = stack->pop<GLbitfield>();
    if (stack->isValid()) {
        GAPID_INFO("glGetPathMetricRangeNV(%u, %" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)",
                   metricQueryMask, firstPathName, numPaths, stride, metrics);
        if (glGetPathMetricRangeNV != nullptr) {
            glGetPathMetricRangeNV(metricQueryMask, firstPathName, numPaths, stride, metrics);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPathMetricRangeNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPathMetricRangeNV");
        return false;
    }
}

bool callGlGetPathMetricsNV(Stack* stack, bool pushReturn) {
    float* metrics = stack->pop<float*>();
    int32_t stride = stack->pop<int32_t>();
    uint32_t pathBase = stack->pop<uint32_t>();
    void* paths = stack->pop<void*>();
    GLenum pathNameType = stack->pop<GLenum>();
    int32_t numPaths = stack->pop<int32_t>();
    GLbitfield metricQueryMask = stack->pop<GLbitfield>();
    if (stack->isValid()) {
        GAPID_INFO("glGetPathMetricsNV(%u, %" PRId32 ", %u, %p, %" PRIu32 ", %" PRId32 ", %p)",
                   metricQueryMask, numPaths, pathNameType, paths, pathBase, stride, metrics);
        if (glGetPathMetricsNV != nullptr) {
            glGetPathMetricsNV(metricQueryMask, numPaths, pathNameType, paths, pathBase, stride,
                               metrics);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPathMetricsNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPathMetricsNV");
        return false;
    }
}

bool callGlGetPathParameterfvNV(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetPathParameterfvNV(%" PRIu32 ", %u, %p)", path, pname, value);
        if (glGetPathParameterfvNV != nullptr) {
            glGetPathParameterfvNV(path, pname, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPathParameterfvNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPathParameterfvNV");
        return false;
    }
}

bool callGlGetPathParameterivNV(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetPathParameterivNV(%" PRIu32 ", %u, %p)", path, pname, value);
        if (glGetPathParameterivNV != nullptr) {
            glGetPathParameterivNV(path, pname, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPathParameterivNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPathParameterivNV");
        return false;
    }
}

bool callGlGetPathSpacingNV(Stack* stack, bool pushReturn) {
    float* returnedSpacing = stack->pop<float*>();
    GLenum transformType = stack->pop<GLenum>();
    float kerningScale = stack->pop<float>();
    float advanceScale = stack->pop<float>();
    uint32_t pathBase = stack->pop<uint32_t>();
    void* paths = stack->pop<void*>();
    GLenum pathNameType = stack->pop<GLenum>();
    int32_t numPaths = stack->pop<int32_t>();
    GLenum pathListMode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetPathSpacingNV(%u, %" PRId32 ", %u, %p, %" PRIu32 ", %f, %f, %u, %p)",
                   pathListMode, numPaths, pathNameType, paths, pathBase, advanceScale,
                   kerningScale, transformType, returnedSpacing);
        if (glGetPathSpacingNV != nullptr) {
            glGetPathSpacingNV(pathListMode, numPaths, pathNameType, paths, pathBase, advanceScale,
                               kerningScale, transformType, returnedSpacing);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPathSpacingNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPathSpacingNV");
        return false;
    }
}

bool callGlGetPerfCounterInfoINTEL(Stack* stack, bool pushReturn) {
    uint64_t* rawCounterMaxValue = stack->pop<uint64_t*>();
    uint32_t* counterDataTypeEnum = stack->pop<uint32_t*>();
    uint32_t* counterTypeEnum = stack->pop<uint32_t*>();
    uint32_t* counterDataSize = stack->pop<uint32_t*>();
    uint32_t* counterOffset = stack->pop<uint32_t*>();
    char* counterDesc = stack->pop<char*>();
    uint32_t counterDescLength = stack->pop<uint32_t>();
    char* counterName = stack->pop<char*>();
    uint32_t counterNameLength = stack->pop<uint32_t>();
    uint32_t counterId = stack->pop<uint32_t>();
    uint32_t queryId = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetPerfCounterInfoINTEL(%" PRIu32 ", %" PRIu32 ", %" PRIu32 ", %p, %" PRIu32
                   ", %p, %p, %p, %p, %p, %p)",
                   queryId, counterId, counterNameLength, counterName, counterDescLength,
                   counterDesc, counterOffset, counterDataSize, counterTypeEnum,
                   counterDataTypeEnum, rawCounterMaxValue);
        if (glGetPerfCounterInfoINTEL != nullptr) {
            glGetPerfCounterInfoINTEL(queryId, counterId, counterNameLength, counterName,
                                      counterDescLength, counterDesc, counterOffset,
                                      counterDataSize, counterTypeEnum, counterDataTypeEnum,
                                      rawCounterMaxValue);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPerfCounterInfoINTEL");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPerfCounterInfoINTEL");
        return false;
    }
}

bool callGlGetPerfMonitorCounterDataAMD(Stack* stack, bool pushReturn) {
    int32_t* bytesWritten = stack->pop<int32_t*>();
    uint32_t* data = stack->pop<uint32_t*>();
    int32_t dataSize = stack->pop<int32_t>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t monitor = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetPerfMonitorCounterDataAMD(%" PRIu32 ", %u, %" PRId32 ", %p, %p)", monitor,
                   pname, dataSize, data, bytesWritten);
        if (glGetPerfMonitorCounterDataAMD != nullptr) {
            glGetPerfMonitorCounterDataAMD(monitor, pname, dataSize, data, bytesWritten);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPerfMonitorCounterDataAMD");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPerfMonitorCounterDataAMD");
        return false;
    }
}

bool callGlGetPerfMonitorCounterInfoAMD(Stack* stack, bool pushReturn) {
    void* data = stack->pop<void*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t counter = stack->pop<uint32_t>();
    uint32_t group = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetPerfMonitorCounterInfoAMD(%" PRIu32 ", %" PRIu32 ", %u, %p)", group,
                   counter, pname, data);
        if (glGetPerfMonitorCounterInfoAMD != nullptr) {
            glGetPerfMonitorCounterInfoAMD(group, counter, pname, data);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPerfMonitorCounterInfoAMD");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPerfMonitorCounterInfoAMD");
        return false;
    }
}

bool callGlGetPerfMonitorCounterStringAMD(Stack* stack, bool pushReturn) {
    char* counterString = stack->pop<char*>();
    int32_t* length = stack->pop<int32_t*>();
    int32_t bufSize = stack->pop<int32_t>();
    uint32_t counter = stack->pop<uint32_t>();
    uint32_t group = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetPerfMonitorCounterStringAMD(%" PRIu32 ", %" PRIu32 ", %" PRId32
                   ", %p, %p)",
                   group, counter, bufSize, length, counterString);
        if (glGetPerfMonitorCounterStringAMD != nullptr) {
            glGetPerfMonitorCounterStringAMD(group, counter, bufSize, length, counterString);
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glGetPerfMonitorCounterStringAMD");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPerfMonitorCounterStringAMD");
        return false;
    }
}

bool callGlGetPerfMonitorCountersAMD(Stack* stack, bool pushReturn) {
    uint32_t* counters = stack->pop<uint32_t*>();
    int32_t counterSize = stack->pop<int32_t>();
    int32_t* maxActiveCounters = stack->pop<int32_t*>();
    int32_t* numCounters = stack->pop<int32_t*>();
    uint32_t group = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetPerfMonitorCountersAMD(%" PRIu32 ", %p, %p, %" PRId32 ", %p)", group,
                   numCounters, maxActiveCounters, counterSize, counters);
        if (glGetPerfMonitorCountersAMD != nullptr) {
            glGetPerfMonitorCountersAMD(group, numCounters, maxActiveCounters, counterSize,
                                        counters);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPerfMonitorCountersAMD");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPerfMonitorCountersAMD");
        return false;
    }
}

bool callGlGetPerfMonitorGroupStringAMD(Stack* stack, bool pushReturn) {
    char* groupString = stack->pop<char*>();
    int32_t* length = stack->pop<int32_t*>();
    int32_t bufSize = stack->pop<int32_t>();
    uint32_t group = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetPerfMonitorGroupStringAMD(%" PRIu32 ", %" PRId32 ", %p, %p)", group,
                   bufSize, length, groupString);
        if (glGetPerfMonitorGroupStringAMD != nullptr) {
            glGetPerfMonitorGroupStringAMD(group, bufSize, length, groupString);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPerfMonitorGroupStringAMD");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPerfMonitorGroupStringAMD");
        return false;
    }
}

bool callGlGetPerfMonitorGroupsAMD(Stack* stack, bool pushReturn) {
    uint32_t* groups = stack->pop<uint32_t*>();
    int32_t groupsSize = stack->pop<int32_t>();
    int32_t* numGroups = stack->pop<int32_t*>();
    if (stack->isValid()) {
        GAPID_INFO("glGetPerfMonitorGroupsAMD(%p, %" PRId32 ", %p)", numGroups, groupsSize, groups);
        if (glGetPerfMonitorGroupsAMD != nullptr) {
            glGetPerfMonitorGroupsAMD(numGroups, groupsSize, groups);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPerfMonitorGroupsAMD");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPerfMonitorGroupsAMD");
        return false;
    }
}

bool callGlGetPerfQueryDataINTEL(Stack* stack, bool pushReturn) {
    uint32_t* bytesWritten = stack->pop<uint32_t*>();
    void* data = stack->pop<void*>();
    int32_t dataSize = stack->pop<int32_t>();
    uint32_t flag = stack->pop<uint32_t>();
    uint32_t queryHandle = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetPerfQueryDataINTEL(%" PRIu32 ", %" PRIu32 ", %" PRId32 ", %p, %p)",
                   queryHandle, flag, dataSize, data, bytesWritten);
        if (glGetPerfQueryDataINTEL != nullptr) {
            glGetPerfQueryDataINTEL(queryHandle, flag, dataSize, data, bytesWritten);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPerfQueryDataINTEL");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPerfQueryDataINTEL");
        return false;
    }
}

bool callGlGetPerfQueryIdByNameINTEL(Stack* stack, bool pushReturn) {
    uint32_t* queryId = stack->pop<uint32_t*>();
    char* queryName = stack->pop<char*>();
    if (stack->isValid()) {
        GAPID_INFO("glGetPerfQueryIdByNameINTEL(%p, %p)", queryName, queryId);
        if (glGetPerfQueryIdByNameINTEL != nullptr) {
            glGetPerfQueryIdByNameINTEL(queryName, queryId);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPerfQueryIdByNameINTEL");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPerfQueryIdByNameINTEL");
        return false;
    }
}

bool callGlGetPerfQueryInfoINTEL(Stack* stack, bool pushReturn) {
    uint32_t* capsMask = stack->pop<uint32_t*>();
    uint32_t* noInstances = stack->pop<uint32_t*>();
    uint32_t* noCounters = stack->pop<uint32_t*>();
    uint32_t* dataSize = stack->pop<uint32_t*>();
    char* queryName = stack->pop<char*>();
    uint32_t queryNameLength = stack->pop<uint32_t>();
    uint32_t queryId = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetPerfQueryInfoINTEL(%" PRIu32 ", %" PRIu32 ", %p, %p, %p, %p, %p)", queryId,
                   queryNameLength, queryName, dataSize, noCounters, noInstances, capsMask);
        if (glGetPerfQueryInfoINTEL != nullptr) {
            glGetPerfQueryInfoINTEL(queryId, queryNameLength, queryName, dataSize, noCounters,
                                    noInstances, capsMask);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPerfQueryInfoINTEL");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPerfQueryInfoINTEL");
        return false;
    }
}

bool callGlGetProgramBinaryOES(Stack* stack, bool pushReturn) {
    void* binary = stack->pop<void*>();
    GLenum* binary_format = stack->pop<GLenum*>();
    int32_t* bytes_written = stack->pop<int32_t*>();
    int32_t buffer_size = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetProgramBinaryOES(%" PRIu32 ", %" PRId32 ", %p, %p, %p)", program,
                   buffer_size, bytes_written, binary_format, binary);
        if (glGetProgramBinaryOES != nullptr) {
            glGetProgramBinaryOES(program, buffer_size, bytes_written, binary_format, binary);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetProgramBinaryOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetProgramBinaryOES");
        return false;
    }
}

bool callGlGetProgramPipelineInfoLogEXT(Stack* stack, bool pushReturn) {
    char* infoLog = stack->pop<char*>();
    int32_t* length = stack->pop<int32_t*>();
    int32_t bufSize = stack->pop<int32_t>();
    uint32_t pipeline = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetProgramPipelineInfoLogEXT(%" PRIu32 ", %" PRId32 ", %p, %p)", pipeline,
                   bufSize, length, infoLog);
        if (glGetProgramPipelineInfoLogEXT != nullptr) {
            glGetProgramPipelineInfoLogEXT(pipeline, bufSize, length, infoLog);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetProgramPipelineInfoLogEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetProgramPipelineInfoLogEXT");
        return false;
    }
}

bool callGlGetProgramPipelineivEXT(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t pipeline = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetProgramPipelineivEXT(%" PRIu32 ", %u, %p)", pipeline, pname, params);
        if (glGetProgramPipelineivEXT != nullptr) {
            glGetProgramPipelineivEXT(pipeline, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetProgramPipelineivEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetProgramPipelineivEXT");
        return false;
    }
}

bool callGlGetProgramResourcefvNV(Stack* stack, bool pushReturn) {
    float* params = stack->pop<float*>();
    int32_t* length = stack->pop<int32_t*>();
    int32_t bufSize = stack->pop<int32_t>();
    GLenum* props = stack->pop<GLenum*>();
    int32_t propCount = stack->pop<int32_t>();
    uint32_t index = stack->pop<uint32_t>();
    GLenum programInterface = stack->pop<GLenum>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetProgramResourcefvNV(%" PRIu32 ", %u, %" PRIu32 ", %" PRId32
                   ", %p, %" PRId32 ", %p, %p)",
                   program, programInterface, index, propCount, props, bufSize, length, params);
        if (glGetProgramResourcefvNV != nullptr) {
            glGetProgramResourcefvNV(program, programInterface, index, propCount, props, bufSize,
                                     length, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetProgramResourcefvNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetProgramResourcefvNV");
        return false;
    }
}

bool callGlGetQueryObjecti64vEXT(Stack* stack, bool pushReturn) {
    int64_t* value = stack->pop<int64_t*>();
    GLenum parameter = stack->pop<GLenum>();
    uint32_t query = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetQueryObjecti64vEXT(%" PRIu32 ", %u, %p)", query, parameter, value);
        if (glGetQueryObjecti64vEXT != nullptr) {
            glGetQueryObjecti64vEXT(query, parameter, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetQueryObjecti64vEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetQueryObjecti64vEXT");
        return false;
    }
}

bool callGlGetQueryObjectivEXT(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    GLenum parameter = stack->pop<GLenum>();
    uint32_t query = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetQueryObjectivEXT(%" PRIu32 ", %u, %p)", query, parameter, value);
        if (glGetQueryObjectivEXT != nullptr) {
            glGetQueryObjectivEXT(query, parameter, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetQueryObjectivEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetQueryObjectivEXT");
        return false;
    }
}

bool callGlGetQueryObjectui64vEXT(Stack* stack, bool pushReturn) {
    uint64_t* value = stack->pop<uint64_t*>();
    GLenum parameter = stack->pop<GLenum>();
    uint32_t query = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetQueryObjectui64vEXT(%" PRIu32 ", %u, %p)", query, parameter, value);
        if (glGetQueryObjectui64vEXT != nullptr) {
            glGetQueryObjectui64vEXT(query, parameter, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetQueryObjectui64vEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetQueryObjectui64vEXT");
        return false;
    }
}

bool callGlGetQueryObjectuivEXT(Stack* stack, bool pushReturn) {
    uint32_t* value = stack->pop<uint32_t*>();
    GLenum parameter = stack->pop<GLenum>();
    uint32_t query = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetQueryObjectuivEXT(%" PRIu32 ", %u, %p)", query, parameter, value);
        if (glGetQueryObjectuivEXT != nullptr) {
            glGetQueryObjectuivEXT(query, parameter, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetQueryObjectuivEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetQueryObjectuivEXT");
        return false;
    }
}

bool callGlGetQueryivEXT(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    GLenum parameter = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetQueryivEXT(%u, %u, %p)", target, parameter, value);
        if (glGetQueryivEXT != nullptr) {
            glGetQueryivEXT(target, parameter, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetQueryivEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetQueryivEXT");
        return false;
    }
}

bool callGlGetSamplerParameterIivOES(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetSamplerParameterIivOES(%" PRIu32 ", %u, %p)", sampler, pname, params);
        if (glGetSamplerParameterIivOES != nullptr) {
            glGetSamplerParameterIivOES(sampler, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetSamplerParameterIivOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetSamplerParameterIivOES");
        return false;
    }
}

bool callGlGetSamplerParameterIuivOES(Stack* stack, bool pushReturn) {
    uint32_t* params = stack->pop<uint32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetSamplerParameterIuivOES(%" PRIu32 ", %u, %p)", sampler, pname, params);
        if (glGetSamplerParameterIuivOES != nullptr) {
            glGetSamplerParameterIuivOES(sampler, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetSamplerParameterIuivOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetSamplerParameterIuivOES");
        return false;
    }
}

bool callGlGetSyncivAPPLE(Stack* stack, bool pushReturn) {
    int32_t* values = stack->pop<int32_t*>();
    int32_t* length = stack->pop<int32_t*>();
    int32_t bufSize = stack->pop<int32_t>();
    GLenum pname = stack->pop<GLenum>();
    uint64_t sync = stack->pop<uint64_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetSyncivAPPLE(%" PRIu64 ", %u, %" PRId32 ", %p, %p)", sync, pname, bufSize,
                   length, values);
        if (glGetSyncivAPPLE != nullptr) {
            glGetSyncivAPPLE(sync, pname, bufSize, length, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetSyncivAPPLE");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetSyncivAPPLE");
        return false;
    }
}

bool callGlGetTexParameterIivOES(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetTexParameterIivOES(%u, %u, %p)", target, pname, params);
        if (glGetTexParameterIivOES != nullptr) {
            glGetTexParameterIivOES(target, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetTexParameterIivOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetTexParameterIivOES");
        return false;
    }
}

bool callGlGetTexParameterIuivOES(Stack* stack, bool pushReturn) {
    uint32_t* params = stack->pop<uint32_t*>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetTexParameterIuivOES(%u, %u, %p)", target, pname, params);
        if (glGetTexParameterIuivOES != nullptr) {
            glGetTexParameterIuivOES(target, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetTexParameterIuivOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetTexParameterIuivOES");
        return false;
    }
}

bool callGlGetTextureHandleNV(Stack* stack, bool pushReturn) {
    uint32_t texture = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetTextureHandleNV(%" PRIu32 ")", texture);
        if (glGetTextureHandleNV != nullptr) {
            uint64_t return_value = glGetTextureHandleNV(texture);
            GAPID_INFO("Returned: %" PRIu64 "", return_value);
            if (pushReturn) {
                stack->push<uint64_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetTextureHandleNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetTextureHandleNV");
        return false;
    }
}

bool callGlGetTextureSamplerHandleNV(Stack* stack, bool pushReturn) {
    uint32_t sampler = stack->pop<uint32_t>();
    uint32_t texture = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetTextureSamplerHandleNV(%" PRIu32 ", %" PRIu32 ")", texture, sampler);
        if (glGetTextureSamplerHandleNV != nullptr) {
            uint64_t return_value = glGetTextureSamplerHandleNV(texture, sampler);
            GAPID_INFO("Returned: %" PRIu64 "", return_value);
            if (pushReturn) {
                stack->push<uint64_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetTextureSamplerHandleNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetTextureSamplerHandleNV");
        return false;
    }
}

bool callGlGetTranslatedShaderSourceANGLE(Stack* stack, bool pushReturn) {
    char* source = stack->pop<char*>();
    int32_t* length = stack->pop<int32_t*>();
    int32_t bufsize = stack->pop<int32_t>();
    uint32_t shader = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetTranslatedShaderSourceANGLE(%" PRIu32 ", %" PRId32 ", %p, %p)", shader,
                   bufsize, length, source);
        if (glGetTranslatedShaderSourceANGLE != nullptr) {
            glGetTranslatedShaderSourceANGLE(shader, bufsize, length, source);
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glGetTranslatedShaderSourceANGLE");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetTranslatedShaderSourceANGLE");
        return false;
    }
}

bool callGlGetnUniformfvEXT(Stack* stack, bool pushReturn) {
    float* params = stack->pop<float*>();
    int32_t bufSize = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetnUniformfvEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)", program,
                   location, bufSize, params);
        if (glGetnUniformfvEXT != nullptr) {
            glGetnUniformfvEXT(program, location, bufSize, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetnUniformfvEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetnUniformfvEXT");
        return false;
    }
}

bool callGlGetnUniformfvKHR(Stack* stack, bool pushReturn) {
    float* params = stack->pop<float*>();
    int32_t bufSize = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetnUniformfvKHR(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)", program,
                   location, bufSize, params);
        if (glGetnUniformfvKHR != nullptr) {
            glGetnUniformfvKHR(program, location, bufSize, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetnUniformfvKHR");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetnUniformfvKHR");
        return false;
    }
}

bool callGlGetnUniformivEXT(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    int32_t bufSize = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetnUniformivEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)", program,
                   location, bufSize, params);
        if (glGetnUniformivEXT != nullptr) {
            glGetnUniformivEXT(program, location, bufSize, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetnUniformivEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetnUniformivEXT");
        return false;
    }
}

bool callGlGetnUniformivKHR(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    int32_t bufSize = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetnUniformivKHR(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)", program,
                   location, bufSize, params);
        if (glGetnUniformivKHR != nullptr) {
            glGetnUniformivKHR(program, location, bufSize, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetnUniformivKHR");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetnUniformivKHR");
        return false;
    }
}

bool callGlGetnUniformuivKHR(Stack* stack, bool pushReturn) {
    uint32_t* params = stack->pop<uint32_t*>();
    int32_t bufSize = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetnUniformuivKHR(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)", program,
                   location, bufSize, params);
        if (glGetnUniformuivKHR != nullptr) {
            glGetnUniformuivKHR(program, location, bufSize, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetnUniformuivKHR");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetnUniformuivKHR");
        return false;
    }
}

bool callGlInsertEventMarkerEXT(Stack* stack, bool pushReturn) {
    char* marker = stack->pop<char*>();
    int32_t length = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glInsertEventMarkerEXT(%" PRId32 ", %p)", length, marker);
        if (glInsertEventMarkerEXT != nullptr) {
            glInsertEventMarkerEXT(length, marker);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glInsertEventMarkerEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glInsertEventMarkerEXT");
        return false;
    }
}

bool callGlInterpolatePathsNV(Stack* stack, bool pushReturn) {
    float weight = stack->pop<float>();
    uint32_t pathB = stack->pop<uint32_t>();
    uint32_t pathA = stack->pop<uint32_t>();
    uint32_t resultPath = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glInterpolatePathsNV(%" PRIu32 ", %" PRIu32 ", %" PRIu32 ", %f)", resultPath,
                   pathA, pathB, weight);
        if (glInterpolatePathsNV != nullptr) {
            glInterpolatePathsNV(resultPath, pathA, pathB, weight);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glInterpolatePathsNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glInterpolatePathsNV");
        return false;
    }
}

bool callGlIsEnablediNV(Stack* stack, bool pushReturn) {
    uint32_t index = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glIsEnablediNV(%u, %" PRIu32 ")", target, index);
        if (glIsEnablediNV != nullptr) {
            uint8_t return_value = glIsEnablediNV(target, index);
            GAPID_INFO("Returned: %" PRIu8 "", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsEnablediNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsEnablediNV");
        return false;
    }
}

bool callGlIsEnablediOES(Stack* stack, bool pushReturn) {
    uint32_t index = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glIsEnablediOES(%u, %" PRIu32 ")", target, index);
        if (glIsEnablediOES != nullptr) {
            uint8_t return_value = glIsEnablediOES(target, index);
            GAPID_INFO("Returned: %" PRIu8 "", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsEnablediOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsEnablediOES");
        return false;
    }
}

bool callGlIsFenceNV(Stack* stack, bool pushReturn) {
    uint32_t fence = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsFenceNV(%" PRIu32 ")", fence);
        if (glIsFenceNV != nullptr) {
            uint8_t return_value = glIsFenceNV(fence);
            GAPID_INFO("Returned: %" PRIu8 "", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsFenceNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsFenceNV");
        return false;
    }
}

bool callGlIsImageHandleResidentNV(Stack* stack, bool pushReturn) {
    uint64_t handle = stack->pop<uint64_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsImageHandleResidentNV(%" PRIu64 ")", handle);
        if (glIsImageHandleResidentNV != nullptr) {
            uint8_t return_value = glIsImageHandleResidentNV(handle);
            GAPID_INFO("Returned: %" PRIu8 "", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsImageHandleResidentNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsImageHandleResidentNV");
        return false;
    }
}

bool callGlIsPathNV(Stack* stack, bool pushReturn) {
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsPathNV(%" PRIu32 ")", path);
        if (glIsPathNV != nullptr) {
            uint8_t return_value = glIsPathNV(path);
            GAPID_INFO("Returned: %" PRIu8 "", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsPathNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsPathNV");
        return false;
    }
}

bool callGlIsPointInFillPathNV(Stack* stack, bool pushReturn) {
    float y = stack->pop<float>();
    float x = stack->pop<float>();
    uint32_t mask = stack->pop<uint32_t>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsPointInFillPathNV(%" PRIu32 ", %" PRIu32 ", %f, %f)", path, mask, x, y);
        if (glIsPointInFillPathNV != nullptr) {
            uint8_t return_value = glIsPointInFillPathNV(path, mask, x, y);
            GAPID_INFO("Returned: %" PRIu8 "", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsPointInFillPathNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsPointInFillPathNV");
        return false;
    }
}

bool callGlIsPointInStrokePathNV(Stack* stack, bool pushReturn) {
    float y = stack->pop<float>();
    float x = stack->pop<float>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsPointInStrokePathNV(%" PRIu32 ", %f, %f)", path, x, y);
        if (glIsPointInStrokePathNV != nullptr) {
            uint8_t return_value = glIsPointInStrokePathNV(path, x, y);
            GAPID_INFO("Returned: %" PRIu8 "", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsPointInStrokePathNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsPointInStrokePathNV");
        return false;
    }
}

bool callGlIsProgramPipelineEXT(Stack* stack, bool pushReturn) {
    uint32_t pipeline = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsProgramPipelineEXT(%" PRIu32 ")", pipeline);
        if (glIsProgramPipelineEXT != nullptr) {
            uint8_t return_value = glIsProgramPipelineEXT(pipeline);
            GAPID_INFO("Returned: %" PRIu8 "", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsProgramPipelineEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsProgramPipelineEXT");
        return false;
    }
}

bool callGlIsQueryEXT(Stack* stack, bool pushReturn) {
    uint32_t query = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsQueryEXT(%" PRIu32 ")", query);
        if (glIsQueryEXT != nullptr) {
            uint8_t return_value = glIsQueryEXT(query);
            GAPID_INFO("Returned: %" PRIu8 "", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsQueryEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsQueryEXT");
        return false;
    }
}

bool callGlIsSyncAPPLE(Stack* stack, bool pushReturn) {
    uint64_t sync = stack->pop<uint64_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsSyncAPPLE(%" PRIu64 ")", sync);
        if (glIsSyncAPPLE != nullptr) {
            uint8_t return_value = glIsSyncAPPLE(sync);
            GAPID_INFO("Returned: %" PRIu8 "", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsSyncAPPLE");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsSyncAPPLE");
        return false;
    }
}

bool callGlIsTextureHandleResidentNV(Stack* stack, bool pushReturn) {
    uint64_t handle = stack->pop<uint64_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsTextureHandleResidentNV(%" PRIu64 ")", handle);
        if (glIsTextureHandleResidentNV != nullptr) {
            uint8_t return_value = glIsTextureHandleResidentNV(handle);
            GAPID_INFO("Returned: %" PRIu8 "", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsTextureHandleResidentNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsTextureHandleResidentNV");
        return false;
    }
}

bool callGlIsVertexArrayOES(Stack* stack, bool pushReturn) {
    uint32_t array = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsVertexArrayOES(%" PRIu32 ")", array);
        if (glIsVertexArrayOES != nullptr) {
            uint8_t return_value = glIsVertexArrayOES(array);
            GAPID_INFO("Returned: %" PRIu8 "", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsVertexArrayOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsVertexArrayOES");
        return false;
    }
}

bool callGlLabelObjectEXT(Stack* stack, bool pushReturn) {
    char* label = stack->pop<char*>();
    int32_t length = stack->pop<int32_t>();
    uint32_t object = stack->pop<uint32_t>();
    GLenum type = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glLabelObjectEXT(%u, %" PRIu32 ", %" PRId32 ", %p)", type, object, length,
                   label);
        if (glLabelObjectEXT != nullptr) {
            glLabelObjectEXT(type, object, length, label);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glLabelObjectEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glLabelObjectEXT");
        return false;
    }
}

bool callGlMakeImageHandleNonResidentNV(Stack* stack, bool pushReturn) {
    uint64_t handle = stack->pop<uint64_t>();
    if (stack->isValid()) {
        GAPID_INFO("glMakeImageHandleNonResidentNV(%" PRIu64 ")", handle);
        if (glMakeImageHandleNonResidentNV != nullptr) {
            glMakeImageHandleNonResidentNV(handle);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMakeImageHandleNonResidentNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMakeImageHandleNonResidentNV");
        return false;
    }
}

bool callGlMakeImageHandleResidentNV(Stack* stack, bool pushReturn) {
    GLenum access = stack->pop<GLenum>();
    uint64_t handle = stack->pop<uint64_t>();
    if (stack->isValid()) {
        GAPID_INFO("glMakeImageHandleResidentNV(%" PRIu64 ", %u)", handle, access);
        if (glMakeImageHandleResidentNV != nullptr) {
            glMakeImageHandleResidentNV(handle, access);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMakeImageHandleResidentNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMakeImageHandleResidentNV");
        return false;
    }
}

bool callGlMakeTextureHandleNonResidentNV(Stack* stack, bool pushReturn) {
    uint64_t handle = stack->pop<uint64_t>();
    if (stack->isValid()) {
        GAPID_INFO("glMakeTextureHandleNonResidentNV(%" PRIu64 ")", handle);
        if (glMakeTextureHandleNonResidentNV != nullptr) {
            glMakeTextureHandleNonResidentNV(handle);
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glMakeTextureHandleNonResidentNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMakeTextureHandleNonResidentNV");
        return false;
    }
}

bool callGlMakeTextureHandleResidentNV(Stack* stack, bool pushReturn) {
    uint64_t handle = stack->pop<uint64_t>();
    if (stack->isValid()) {
        GAPID_INFO("glMakeTextureHandleResidentNV(%" PRIu64 ")", handle);
        if (glMakeTextureHandleResidentNV != nullptr) {
            glMakeTextureHandleResidentNV(handle);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMakeTextureHandleResidentNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMakeTextureHandleResidentNV");
        return false;
    }
}

bool callGlMapBufferOES(Stack* stack, bool pushReturn) {
    GLenum access = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glMapBufferOES(%u, %u)", target, access);
        if (glMapBufferOES != nullptr) {
            void* return_value = glMapBufferOES(target, access);
            GAPID_INFO("Returned: %p", return_value);
            if (pushReturn) {
                stack->push<void*>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMapBufferOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMapBufferOES");
        return false;
    }
}

bool callGlMapBufferRangeEXT(Stack* stack, bool pushReturn) {
    GLbitfield access = stack->pop<GLbitfield>();
    int32_t length = stack->pop<int32_t>();
    int32_t offset = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glMapBufferRangeEXT(%u, %" PRId32 ", %" PRId32 ", %u)", target, offset, length,
                   access);
        if (glMapBufferRangeEXT != nullptr) {
            void* return_value = glMapBufferRangeEXT(target, offset, length, access);
            GAPID_INFO("Returned: %p", return_value);
            if (pushReturn) {
                stack->push<void*>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMapBufferRangeEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMapBufferRangeEXT");
        return false;
    }
}

bool callGlMatrixLoad3x2fNV(Stack* stack, bool pushReturn) {
    float* m = stack->pop<float*>();
    GLenum matrixMode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glMatrixLoad3x2fNV(%u, %p)", matrixMode, m);
        if (glMatrixLoad3x2fNV != nullptr) {
            glMatrixLoad3x2fNV(matrixMode, m);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMatrixLoad3x2fNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMatrixLoad3x2fNV");
        return false;
    }
}

bool callGlMatrixLoad3x3fNV(Stack* stack, bool pushReturn) {
    float* m = stack->pop<float*>();
    GLenum matrixMode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glMatrixLoad3x3fNV(%u, %p)", matrixMode, m);
        if (glMatrixLoad3x3fNV != nullptr) {
            glMatrixLoad3x3fNV(matrixMode, m);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMatrixLoad3x3fNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMatrixLoad3x3fNV");
        return false;
    }
}

bool callGlMatrixLoadTranspose3x3fNV(Stack* stack, bool pushReturn) {
    float* m = stack->pop<float*>();
    GLenum matrixMode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glMatrixLoadTranspose3x3fNV(%u, %p)", matrixMode, m);
        if (glMatrixLoadTranspose3x3fNV != nullptr) {
            glMatrixLoadTranspose3x3fNV(matrixMode, m);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMatrixLoadTranspose3x3fNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMatrixLoadTranspose3x3fNV");
        return false;
    }
}

bool callGlMatrixMult3x2fNV(Stack* stack, bool pushReturn) {
    float* m = stack->pop<float*>();
    GLenum matrixMode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glMatrixMult3x2fNV(%u, %p)", matrixMode, m);
        if (glMatrixMult3x2fNV != nullptr) {
            glMatrixMult3x2fNV(matrixMode, m);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMatrixMult3x2fNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMatrixMult3x2fNV");
        return false;
    }
}

bool callGlMatrixMult3x3fNV(Stack* stack, bool pushReturn) {
    float* m = stack->pop<float*>();
    GLenum matrixMode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glMatrixMult3x3fNV(%u, %p)", matrixMode, m);
        if (glMatrixMult3x3fNV != nullptr) {
            glMatrixMult3x3fNV(matrixMode, m);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMatrixMult3x3fNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMatrixMult3x3fNV");
        return false;
    }
}

bool callGlMatrixMultTranspose3x3fNV(Stack* stack, bool pushReturn) {
    float* m = stack->pop<float*>();
    GLenum matrixMode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glMatrixMultTranspose3x3fNV(%u, %p)", matrixMode, m);
        if (glMatrixMultTranspose3x3fNV != nullptr) {
            glMatrixMultTranspose3x3fNV(matrixMode, m);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMatrixMultTranspose3x3fNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMatrixMultTranspose3x3fNV");
        return false;
    }
}

bool callGlMultiDrawArraysEXT(Stack* stack, bool pushReturn) {
    int32_t primcount = stack->pop<int32_t>();
    int32_t* count = stack->pop<int32_t*>();
    int32_t* first = stack->pop<int32_t*>();
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glMultiDrawArraysEXT(%u, %p, %p, %" PRId32 ")", mode, first, count, primcount);
        if (glMultiDrawArraysEXT != nullptr) {
            glMultiDrawArraysEXT(mode, first, count, primcount);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMultiDrawArraysEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMultiDrawArraysEXT");
        return false;
    }
}

bool callGlMultiDrawArraysIndirectEXT(Stack* stack, bool pushReturn) {
    int32_t stride = stack->pop<int32_t>();
    int32_t drawcount = stack->pop<int32_t>();
    void* indirect = stack->pop<void*>();
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glMultiDrawArraysIndirectEXT(%u, %p, %" PRId32 ", %" PRId32 ")", mode, indirect,
                   drawcount, stride);
        if (glMultiDrawArraysIndirectEXT != nullptr) {
            glMultiDrawArraysIndirectEXT(mode, indirect, drawcount, stride);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMultiDrawArraysIndirectEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMultiDrawArraysIndirectEXT");
        return false;
    }
}

bool callGlMultiDrawElementsBaseVertexEXT(Stack* stack, bool pushReturn) {
    int32_t* basevertex = stack->pop<int32_t*>();
    int32_t primcount = stack->pop<int32_t>();
    void** indices = stack->pop<void**>();
    GLenum type = stack->pop<GLenum>();
    int32_t* count = stack->pop<int32_t*>();
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glMultiDrawElementsBaseVertexEXT(%u, %p, %u, %p, %" PRId32 ", %p)", mode, count,
                   type, indices, primcount, basevertex);
        if (glMultiDrawElementsBaseVertexEXT != nullptr) {
            glMultiDrawElementsBaseVertexEXT(mode, count, type, indices, primcount, basevertex);
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glMultiDrawElementsBaseVertexEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMultiDrawElementsBaseVertexEXT");
        return false;
    }
}

bool callGlMultiDrawElementsBaseVertexOES(Stack* stack, bool pushReturn) {
    int32_t* basevertex = stack->pop<int32_t*>();
    int32_t primcount = stack->pop<int32_t>();
    void** indices = stack->pop<void**>();
    GLenum type = stack->pop<GLenum>();
    int32_t* count = stack->pop<int32_t*>();
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glMultiDrawElementsBaseVertexOES(%u, %p, %u, %p, %" PRId32 ", %p)", mode, count,
                   type, indices, primcount, basevertex);
        if (glMultiDrawElementsBaseVertexOES != nullptr) {
            glMultiDrawElementsBaseVertexOES(mode, count, type, indices, primcount, basevertex);
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glMultiDrawElementsBaseVertexOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMultiDrawElementsBaseVertexOES");
        return false;
    }
}

bool callGlMultiDrawElementsEXT(Stack* stack, bool pushReturn) {
    int32_t primcount = stack->pop<int32_t>();
    void** indices = stack->pop<void**>();
    GLenum type = stack->pop<GLenum>();
    int32_t* count = stack->pop<int32_t*>();
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glMultiDrawElementsEXT(%u, %p, %u, %p, %" PRId32 ")", mode, count, type,
                   indices, primcount);
        if (glMultiDrawElementsEXT != nullptr) {
            glMultiDrawElementsEXT(mode, count, type, indices, primcount);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMultiDrawElementsEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMultiDrawElementsEXT");
        return false;
    }
}

bool callGlMultiDrawElementsIndirectEXT(Stack* stack, bool pushReturn) {
    int32_t stride = stack->pop<int32_t>();
    int32_t drawcount = stack->pop<int32_t>();
    void* indirect = stack->pop<void*>();
    GLenum type = stack->pop<GLenum>();
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glMultiDrawElementsIndirectEXT(%u, %u, %p, %" PRId32 ", %" PRId32 ")", mode,
                   type, indirect, drawcount, stride);
        if (glMultiDrawElementsIndirectEXT != nullptr) {
            glMultiDrawElementsIndirectEXT(mode, type, indirect, drawcount, stride);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMultiDrawElementsIndirectEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMultiDrawElementsIndirectEXT");
        return false;
    }
}

bool callGlNamedFramebufferSampleLocationsfvNV(Stack* stack, bool pushReturn) {
    float* v = stack->pop<float*>();
    int32_t count = stack->pop<int32_t>();
    uint32_t start = stack->pop<uint32_t>();
    uint32_t framebuffer = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glNamedFramebufferSampleLocationsfvNV(%" PRIu32 ", %" PRIu32 ", %" PRId32
                   ", %p)",
                   framebuffer, start, count, v);
        if (glNamedFramebufferSampleLocationsfvNV != nullptr) {
            glNamedFramebufferSampleLocationsfvNV(framebuffer, start, count, v);
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glNamedFramebufferSampleLocationsfvNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glNamedFramebufferSampleLocationsfvNV");
        return false;
    }
}

bool callGlPatchParameteriOES(Stack* stack, bool pushReturn) {
    int32_t value = stack->pop<int32_t>();
    GLenum pname = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glPatchParameteriOES(%u, %" PRId32 ")", pname, value);
        if (glPatchParameteriOES != nullptr) {
            glPatchParameteriOES(pname, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPatchParameteriOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPatchParameteriOES");
        return false;
    }
}

bool callGlPathCommandsNV(Stack* stack, bool pushReturn) {
    void* coords = stack->pop<void*>();
    GLenum coordType = stack->pop<GLenum>();
    int32_t numCoords = stack->pop<int32_t>();
    uint8_t* commands = stack->pop<uint8_t*>();
    int32_t numCommands = stack->pop<int32_t>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glPathCommandsNV(%" PRIu32 ", %" PRId32 ", %p, %" PRId32 ", %u, %p)", path,
                   numCommands, commands, numCoords, coordType, coords);
        if (glPathCommandsNV != nullptr) {
            glPathCommandsNV(path, numCommands, commands, numCoords, coordType, coords);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathCommandsNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathCommandsNV");
        return false;
    }
}

bool callGlPathCoordsNV(Stack* stack, bool pushReturn) {
    void* coords = stack->pop<void*>();
    GLenum coordType = stack->pop<GLenum>();
    int32_t numCoords = stack->pop<int32_t>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glPathCoordsNV(%" PRIu32 ", %" PRId32 ", %u, %p)", path, numCoords, coordType,
                   coords);
        if (glPathCoordsNV != nullptr) {
            glPathCoordsNV(path, numCoords, coordType, coords);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathCoordsNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathCoordsNV");
        return false;
    }
}

bool callGlPathCoverDepthFuncNV(Stack* stack, bool pushReturn) {
    GLenum func = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glPathCoverDepthFuncNV(%u)", func);
        if (glPathCoverDepthFuncNV != nullptr) {
            glPathCoverDepthFuncNV(func);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathCoverDepthFuncNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathCoverDepthFuncNV");
        return false;
    }
}

bool callGlPathDashArrayNV(Stack* stack, bool pushReturn) {
    float* dashArray = stack->pop<float*>();
    int32_t dashCount = stack->pop<int32_t>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glPathDashArrayNV(%" PRIu32 ", %" PRId32 ", %p)", path, dashCount, dashArray);
        if (glPathDashArrayNV != nullptr) {
            glPathDashArrayNV(path, dashCount, dashArray);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathDashArrayNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathDashArrayNV");
        return false;
    }
}

bool callGlPathGlyphIndexArrayNV(Stack* stack, bool pushReturn) {
    float emScale = stack->pop<float>();
    uint32_t pathParameterTemplate = stack->pop<uint32_t>();
    int32_t numGlyphs = stack->pop<int32_t>();
    uint32_t firstGlyphIndex = stack->pop<uint32_t>();
    GLbitfield fontStyle = stack->pop<GLbitfield>();
    void* fontName = stack->pop<void*>();
    GLenum fontTarget = stack->pop<GLenum>();
    uint32_t firstPathName = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glPathGlyphIndexArrayNV(%" PRIu32 ", %u, %p, %u, %" PRIu32 ", %" PRId32
                   ", %" PRIu32 ", %f)",
                   firstPathName, fontTarget, fontName, fontStyle, firstGlyphIndex, numGlyphs,
                   pathParameterTemplate, emScale);
        if (glPathGlyphIndexArrayNV != nullptr) {
            GLenum return_value = glPathGlyphIndexArrayNV(firstPathName, fontTarget, fontName,
                                                          fontStyle, firstGlyphIndex, numGlyphs,
                                                          pathParameterTemplate, emScale);
            GAPID_INFO("Returned: %u", return_value);
            if (pushReturn) {
                stack->push<GLenum>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathGlyphIndexArrayNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathGlyphIndexArrayNV");
        return false;
    }
}

bool callGlPathGlyphIndexRangeNV(Stack* stack, bool pushReturn) {
    uint32_t baseAndCount = stack->pop<uint32_t>();
    float emScale = stack->pop<float>();
    uint32_t pathParameterTemplate = stack->pop<uint32_t>();
    GLbitfield fontStyle = stack->pop<GLbitfield>();
    void* fontName = stack->pop<void*>();
    GLenum fontTarget = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glPathGlyphIndexRangeNV(%u, %p, %u, %" PRIu32 ", %f, %" PRIu32 ")", fontTarget,
                   fontName, fontStyle, pathParameterTemplate, emScale, baseAndCount);
        if (glPathGlyphIndexRangeNV != nullptr) {
            GLenum return_value = glPathGlyphIndexRangeNV(
                    fontTarget, fontName, fontStyle, pathParameterTemplate, emScale, baseAndCount);
            GAPID_INFO("Returned: %u", return_value);
            if (pushReturn) {
                stack->push<GLenum>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathGlyphIndexRangeNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathGlyphIndexRangeNV");
        return false;
    }
}

bool callGlPathGlyphRangeNV(Stack* stack, bool pushReturn) {
    float emScale = stack->pop<float>();
    uint32_t pathParameterTemplate = stack->pop<uint32_t>();
    GLenum handleMissingGlyphs = stack->pop<GLenum>();
    int32_t numGlyphs = stack->pop<int32_t>();
    uint32_t firstGlyph = stack->pop<uint32_t>();
    GLbitfield fontStyle = stack->pop<GLbitfield>();
    void* fontName = stack->pop<void*>();
    GLenum fontTarget = stack->pop<GLenum>();
    uint32_t firstPathName = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glPathGlyphRangeNV(%" PRIu32 ", %u, %p, %u, %" PRIu32 ", %" PRId32
                   ", %u, %" PRIu32 ", %f)",
                   firstPathName, fontTarget, fontName, fontStyle, firstGlyph, numGlyphs,
                   handleMissingGlyphs, pathParameterTemplate, emScale);
        if (glPathGlyphRangeNV != nullptr) {
            glPathGlyphRangeNV(firstPathName, fontTarget, fontName, fontStyle, firstGlyph,
                               numGlyphs, handleMissingGlyphs, pathParameterTemplate, emScale);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathGlyphRangeNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathGlyphRangeNV");
        return false;
    }
}

bool callGlPathGlyphsNV(Stack* stack, bool pushReturn) {
    float emScale = stack->pop<float>();
    uint32_t pathParameterTemplate = stack->pop<uint32_t>();
    GLenum handleMissingGlyphs = stack->pop<GLenum>();
    void* charcodes = stack->pop<void*>();
    GLenum type = stack->pop<GLenum>();
    int32_t numGlyphs = stack->pop<int32_t>();
    GLbitfield fontStyle = stack->pop<GLbitfield>();
    void* fontName = stack->pop<void*>();
    GLenum fontTarget = stack->pop<GLenum>();
    uint32_t firstPathName = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glPathGlyphsNV(%" PRIu32 ", %u, %p, %u, %" PRId32 ", %u, %p, %u, %" PRIu32
                   ", %f)",
                   firstPathName, fontTarget, fontName, fontStyle, numGlyphs, type, charcodes,
                   handleMissingGlyphs, pathParameterTemplate, emScale);
        if (glPathGlyphsNV != nullptr) {
            glPathGlyphsNV(firstPathName, fontTarget, fontName, fontStyle, numGlyphs, type,
                           charcodes, handleMissingGlyphs, pathParameterTemplate, emScale);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathGlyphsNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathGlyphsNV");
        return false;
    }
}

bool callGlPathMemoryGlyphIndexArrayNV(Stack* stack, bool pushReturn) {
    float emScale = stack->pop<float>();
    uint32_t pathParameterTemplate = stack->pop<uint32_t>();
    int32_t numGlyphs = stack->pop<int32_t>();
    uint32_t firstGlyphIndex = stack->pop<uint32_t>();
    int32_t faceIndex = stack->pop<int32_t>();
    void* fontData = stack->pop<void*>();
    int32_t fontSize = stack->pop<int32_t>();
    GLenum fontTarget = stack->pop<GLenum>();
    uint32_t firstPathName = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glPathMemoryGlyphIndexArrayNV(%" PRIu32 ", %u, %" PRId32 ", %p, %" PRId32
                   ", %" PRIu32 ", %" PRId32 ", %" PRIu32 ", %f)",
                   firstPathName, fontTarget, fontSize, fontData, faceIndex, firstGlyphIndex,
                   numGlyphs, pathParameterTemplate, emScale);
        if (glPathMemoryGlyphIndexArrayNV != nullptr) {
            GLenum return_value = glPathMemoryGlyphIndexArrayNV(
                    firstPathName, fontTarget, fontSize, fontData, faceIndex, firstGlyphIndex,
                    numGlyphs, pathParameterTemplate, emScale);
            GAPID_INFO("Returned: %u", return_value);
            if (pushReturn) {
                stack->push<GLenum>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathMemoryGlyphIndexArrayNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathMemoryGlyphIndexArrayNV");
        return false;
    }
}

bool callGlPathParameterfNV(Stack* stack, bool pushReturn) {
    float value = stack->pop<float>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glPathParameterfNV(%" PRIu32 ", %u, %f)", path, pname, value);
        if (glPathParameterfNV != nullptr) {
            glPathParameterfNV(path, pname, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathParameterfNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathParameterfNV");
        return false;
    }
}

bool callGlPathParameterfvNV(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glPathParameterfvNV(%" PRIu32 ", %u, %p)", path, pname, value);
        if (glPathParameterfvNV != nullptr) {
            glPathParameterfvNV(path, pname, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathParameterfvNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathParameterfvNV");
        return false;
    }
}

bool callGlPathParameteriNV(Stack* stack, bool pushReturn) {
    int32_t value = stack->pop<int32_t>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glPathParameteriNV(%" PRIu32 ", %u, %" PRId32 ")", path, pname, value);
        if (glPathParameteriNV != nullptr) {
            glPathParameteriNV(path, pname, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathParameteriNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathParameteriNV");
        return false;
    }
}

bool callGlPathParameterivNV(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glPathParameterivNV(%" PRIu32 ", %u, %p)", path, pname, value);
        if (glPathParameterivNV != nullptr) {
            glPathParameterivNV(path, pname, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathParameterivNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathParameterivNV");
        return false;
    }
}

bool callGlPathStencilDepthOffsetNV(Stack* stack, bool pushReturn) {
    float units = stack->pop<float>();
    float factor = stack->pop<float>();
    if (stack->isValid()) {
        GAPID_INFO("glPathStencilDepthOffsetNV(%f, %f)", factor, units);
        if (glPathStencilDepthOffsetNV != nullptr) {
            glPathStencilDepthOffsetNV(factor, units);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathStencilDepthOffsetNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathStencilDepthOffsetNV");
        return false;
    }
}

bool callGlPathStencilFuncNV(Stack* stack, bool pushReturn) {
    uint32_t mask = stack->pop<uint32_t>();
    int32_t ref = stack->pop<int32_t>();
    GLenum func = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glPathStencilFuncNV(%u, %" PRId32 ", %" PRIu32 ")", func, ref, mask);
        if (glPathStencilFuncNV != nullptr) {
            glPathStencilFuncNV(func, ref, mask);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathStencilFuncNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathStencilFuncNV");
        return false;
    }
}

bool callGlPathStringNV(Stack* stack, bool pushReturn) {
    void* pathString = stack->pop<void*>();
    int32_t length = stack->pop<int32_t>();
    GLenum format = stack->pop<GLenum>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glPathStringNV(%" PRIu32 ", %u, %" PRId32 ", %p)", path, format, length,
                   pathString);
        if (glPathStringNV != nullptr) {
            glPathStringNV(path, format, length, pathString);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathStringNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathStringNV");
        return false;
    }
}

bool callGlPathSubCommandsNV(Stack* stack, bool pushReturn) {
    void* coords = stack->pop<void*>();
    GLenum coordType = stack->pop<GLenum>();
    int32_t numCoords = stack->pop<int32_t>();
    uint8_t* commands = stack->pop<uint8_t*>();
    int32_t numCommands = stack->pop<int32_t>();
    int32_t commandsToDelete = stack->pop<int32_t>();
    int32_t commandStart = stack->pop<int32_t>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glPathSubCommandsNV(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %p, %" PRId32 ", %u, %p)",
                   path, commandStart, commandsToDelete, numCommands, commands, numCoords,
                   coordType, coords);
        if (glPathSubCommandsNV != nullptr) {
            glPathSubCommandsNV(path, commandStart, commandsToDelete, numCommands, commands,
                                numCoords, coordType, coords);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathSubCommandsNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathSubCommandsNV");
        return false;
    }
}

bool callGlPathSubCoordsNV(Stack* stack, bool pushReturn) {
    void* coords = stack->pop<void*>();
    GLenum coordType = stack->pop<GLenum>();
    int32_t numCoords = stack->pop<int32_t>();
    int32_t coordStart = stack->pop<int32_t>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glPathSubCoordsNV(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %u, %p)", path,
                   coordStart, numCoords, coordType, coords);
        if (glPathSubCoordsNV != nullptr) {
            glPathSubCoordsNV(path, coordStart, numCoords, coordType, coords);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathSubCoordsNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathSubCoordsNV");
        return false;
    }
}

bool callGlPointAlongPathNV(Stack* stack, bool pushReturn) {
    float* tangentY = stack->pop<float*>();
    float* tangentX = stack->pop<float*>();
    float* y = stack->pop<float*>();
    float* x = stack->pop<float*>();
    float distance = stack->pop<float>();
    int32_t numSegments = stack->pop<int32_t>();
    int32_t startSegment = stack->pop<int32_t>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glPointAlongPathNV(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %f, %p, %p, %p, %p)",
                   path, startSegment, numSegments, distance, x, y, tangentX, tangentY);
        if (glPointAlongPathNV != nullptr) {
            uint8_t return_value = glPointAlongPathNV(path, startSegment, numSegments, distance, x,
                                                      y, tangentX, tangentY);
            GAPID_INFO("Returned: %" PRIu8 "", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPointAlongPathNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPointAlongPathNV");
        return false;
    }
}

bool callGlPolygonModeNV(Stack* stack, bool pushReturn) {
    GLenum mode = stack->pop<GLenum>();
    GLenum face = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glPolygonModeNV(%u, %u)", face, mode);
        if (glPolygonModeNV != nullptr) {
            glPolygonModeNV(face, mode);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPolygonModeNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPolygonModeNV");
        return false;
    }
}

bool callGlPopGroupMarkerEXT(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glPopGroupMarkerEXT()");
        if (glPopGroupMarkerEXT != nullptr) {
            glPopGroupMarkerEXT();
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPopGroupMarkerEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPopGroupMarkerEXT");
        return false;
    }
}

bool callGlPrimitiveBoundingBoxOES(Stack* stack, bool pushReturn) {
    float maxW = stack->pop<float>();
    float maxZ = stack->pop<float>();
    float maxY = stack->pop<float>();
    float maxX = stack->pop<float>();
    float minW = stack->pop<float>();
    float minZ = stack->pop<float>();
    float minY = stack->pop<float>();
    float minX = stack->pop<float>();
    if (stack->isValid()) {
        GAPID_INFO("glPrimitiveBoundingBoxOES(%f, %f, %f, %f, %f, %f, %f, %f)", minX, minY, minZ,
                   minW, maxX, maxY, maxZ, maxW);
        if (glPrimitiveBoundingBoxOES != nullptr) {
            glPrimitiveBoundingBoxOES(minX, minY, minZ, minW, maxX, maxY, maxZ, maxW);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPrimitiveBoundingBoxOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPrimitiveBoundingBoxOES");
        return false;
    }
}

bool callGlProgramBinaryOES(Stack* stack, bool pushReturn) {
    int32_t binary_size = stack->pop<int32_t>();
    void* binary = stack->pop<void*>();
    GLenum binary_format = stack->pop<GLenum>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramBinaryOES(%" PRIu32 ", %u, %p, %" PRId32 ")", program, binary_format,
                   binary, binary_size);
        if (glProgramBinaryOES != nullptr) {
            glProgramBinaryOES(program, binary_format, binary, binary_size);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramBinaryOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramBinaryOES");
        return false;
    }
}

bool callGlProgramParameteriEXT(Stack* stack, bool pushReturn) {
    int32_t value = stack->pop<int32_t>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramParameteriEXT(%" PRIu32 ", %u, %" PRId32 ")", program, pname, value);
        if (glProgramParameteriEXT != nullptr) {
            glProgramParameteriEXT(program, pname, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramParameteriEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramParameteriEXT");
        return false;
    }
}

bool callGlProgramPathFragmentInputGenNV(Stack* stack, bool pushReturn) {
    float* coeffs = stack->pop<float*>();
    int32_t components = stack->pop<int32_t>();
    GLenum genMode = stack->pop<GLenum>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramPathFragmentInputGenNV(%" PRIu32 ", %" PRId32 ", %u, %" PRId32 ", %p)",
                   program, location, genMode, components, coeffs);
        if (glProgramPathFragmentInputGenNV != nullptr) {
            glProgramPathFragmentInputGenNV(program, location, genMode, components, coeffs);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramPathFragmentInputGenNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramPathFragmentInputGenNV");
        return false;
    }
}

bool callGlProgramUniform1fEXT(Stack* stack, bool pushReturn) {
    float v0 = stack->pop<float>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform1fEXT(%" PRIu32 ", %" PRId32 ", %f)", program, location, v0);
        if (glProgramUniform1fEXT != nullptr) {
            glProgramUniform1fEXT(program, location, v0);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform1fEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform1fEXT");
        return false;
    }
}

bool callGlProgramUniform1fvEXT(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform1fvEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)", program,
                   location, count, value);
        if (glProgramUniform1fvEXT != nullptr) {
            glProgramUniform1fvEXT(program, location, count, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform1fvEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform1fvEXT");
        return false;
    }
}

bool callGlProgramUniform1iEXT(Stack* stack, bool pushReturn) {
    int32_t v0 = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform1iEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ")", program,
                   location, v0);
        if (glProgramUniform1iEXT != nullptr) {
            glProgramUniform1iEXT(program, location, v0);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform1iEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform1iEXT");
        return false;
    }
}

bool callGlProgramUniform1ivEXT(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform1ivEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)", program,
                   location, count, value);
        if (glProgramUniform1ivEXT != nullptr) {
            glProgramUniform1ivEXT(program, location, count, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform1ivEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform1ivEXT");
        return false;
    }
}

bool callGlProgramUniform1uiEXT(Stack* stack, bool pushReturn) {
    uint32_t v0 = stack->pop<uint32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform1uiEXT(%" PRIu32 ", %" PRId32 ", %" PRIu32 ")", program,
                   location, v0);
        if (glProgramUniform1uiEXT != nullptr) {
            glProgramUniform1uiEXT(program, location, v0);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform1uiEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform1uiEXT");
        return false;
    }
}

bool callGlProgramUniform1uivEXT(Stack* stack, bool pushReturn) {
    uint32_t* value = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform1uivEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)", program,
                   location, count, value);
        if (glProgramUniform1uivEXT != nullptr) {
            glProgramUniform1uivEXT(program, location, count, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform1uivEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform1uivEXT");
        return false;
    }
}

bool callGlProgramUniform2fEXT(Stack* stack, bool pushReturn) {
    float v1 = stack->pop<float>();
    float v0 = stack->pop<float>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform2fEXT(%" PRIu32 ", %" PRId32 ", %f, %f)", program, location, v0,
                   v1);
        if (glProgramUniform2fEXT != nullptr) {
            glProgramUniform2fEXT(program, location, v0, v1);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform2fEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform2fEXT");
        return false;
    }
}

bool callGlProgramUniform2fvEXT(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform2fvEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)", program,
                   location, count, value);
        if (glProgramUniform2fvEXT != nullptr) {
            glProgramUniform2fvEXT(program, location, count, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform2fvEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform2fvEXT");
        return false;
    }
}

bool callGlProgramUniform2iEXT(Stack* stack, bool pushReturn) {
    int32_t v1 = stack->pop<int32_t>();
    int32_t v0 = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform2iEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ")",
                   program, location, v0, v1);
        if (glProgramUniform2iEXT != nullptr) {
            glProgramUniform2iEXT(program, location, v0, v1);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform2iEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform2iEXT");
        return false;
    }
}

bool callGlProgramUniform2ivEXT(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform2ivEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)", program,
                   location, count, value);
        if (glProgramUniform2ivEXT != nullptr) {
            glProgramUniform2ivEXT(program, location, count, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform2ivEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform2ivEXT");
        return false;
    }
}

bool callGlProgramUniform2uiEXT(Stack* stack, bool pushReturn) {
    uint32_t v1 = stack->pop<uint32_t>();
    uint32_t v0 = stack->pop<uint32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform2uiEXT(%" PRIu32 ", %" PRId32 ", %" PRIu32 ", %" PRIu32 ")",
                   program, location, v0, v1);
        if (glProgramUniform2uiEXT != nullptr) {
            glProgramUniform2uiEXT(program, location, v0, v1);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform2uiEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform2uiEXT");
        return false;
    }
}

bool callGlProgramUniform2uivEXT(Stack* stack, bool pushReturn) {
    uint32_t* value = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform2uivEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)", program,
                   location, count, value);
        if (glProgramUniform2uivEXT != nullptr) {
            glProgramUniform2uivEXT(program, location, count, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform2uivEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform2uivEXT");
        return false;
    }
}

bool callGlProgramUniform3fEXT(Stack* stack, bool pushReturn) {
    float v2 = stack->pop<float>();
    float v1 = stack->pop<float>();
    float v0 = stack->pop<float>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform3fEXT(%" PRIu32 ", %" PRId32 ", %f, %f, %f)", program, location,
                   v0, v1, v2);
        if (glProgramUniform3fEXT != nullptr) {
            glProgramUniform3fEXT(program, location, v0, v1, v2);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform3fEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform3fEXT");
        return false;
    }
}

bool callGlProgramUniform3fvEXT(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform3fvEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)", program,
                   location, count, value);
        if (glProgramUniform3fvEXT != nullptr) {
            glProgramUniform3fvEXT(program, location, count, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform3fvEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform3fvEXT");
        return false;
    }
}

bool callGlProgramUniform3iEXT(Stack* stack, bool pushReturn) {
    int32_t v2 = stack->pop<int32_t>();
    int32_t v1 = stack->pop<int32_t>();
    int32_t v0 = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform3iEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ")",
                   program, location, v0, v1, v2);
        if (glProgramUniform3iEXT != nullptr) {
            glProgramUniform3iEXT(program, location, v0, v1, v2);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform3iEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform3iEXT");
        return false;
    }
}

bool callGlProgramUniform3ivEXT(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform3ivEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)", program,
                   location, count, value);
        if (glProgramUniform3ivEXT != nullptr) {
            glProgramUniform3ivEXT(program, location, count, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform3ivEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform3ivEXT");
        return false;
    }
}

bool callGlProgramUniform3uiEXT(Stack* stack, bool pushReturn) {
    uint32_t v2 = stack->pop<uint32_t>();
    uint32_t v1 = stack->pop<uint32_t>();
    uint32_t v0 = stack->pop<uint32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform3uiEXT(%" PRIu32 ", %" PRId32 ", %" PRIu32 ", %" PRIu32
                   ", %" PRIu32 ")",
                   program, location, v0, v1, v2);
        if (glProgramUniform3uiEXT != nullptr) {
            glProgramUniform3uiEXT(program, location, v0, v1, v2);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform3uiEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform3uiEXT");
        return false;
    }
}

bool callGlProgramUniform3uivEXT(Stack* stack, bool pushReturn) {
    uint32_t* value = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform3uivEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)", program,
                   location, count, value);
        if (glProgramUniform3uivEXT != nullptr) {
            glProgramUniform3uivEXT(program, location, count, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform3uivEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform3uivEXT");
        return false;
    }
}

bool callGlProgramUniform4fEXT(Stack* stack, bool pushReturn) {
    float v3 = stack->pop<float>();
    float v2 = stack->pop<float>();
    float v1 = stack->pop<float>();
    float v0 = stack->pop<float>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform4fEXT(%" PRIu32 ", %" PRId32 ", %f, %f, %f, %f)", program,
                   location, v0, v1, v2, v3);
        if (glProgramUniform4fEXT != nullptr) {
            glProgramUniform4fEXT(program, location, v0, v1, v2, v3);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform4fEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform4fEXT");
        return false;
    }
}

bool callGlProgramUniform4fvEXT(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform4fvEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)", program,
                   location, count, value);
        if (glProgramUniform4fvEXT != nullptr) {
            glProgramUniform4fvEXT(program, location, count, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform4fvEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform4fvEXT");
        return false;
    }
}

bool callGlProgramUniform4iEXT(Stack* stack, bool pushReturn) {
    int32_t v3 = stack->pop<int32_t>();
    int32_t v2 = stack->pop<int32_t>();
    int32_t v1 = stack->pop<int32_t>();
    int32_t v0 = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform4iEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %" PRId32 ")",
                   program, location, v0, v1, v2, v3);
        if (glProgramUniform4iEXT != nullptr) {
            glProgramUniform4iEXT(program, location, v0, v1, v2, v3);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform4iEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform4iEXT");
        return false;
    }
}

bool callGlProgramUniform4ivEXT(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform4ivEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)", program,
                   location, count, value);
        if (glProgramUniform4ivEXT != nullptr) {
            glProgramUniform4ivEXT(program, location, count, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform4ivEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform4ivEXT");
        return false;
    }
}

bool callGlProgramUniform4uiEXT(Stack* stack, bool pushReturn) {
    uint32_t v3 = stack->pop<uint32_t>();
    uint32_t v2 = stack->pop<uint32_t>();
    uint32_t v1 = stack->pop<uint32_t>();
    uint32_t v0 = stack->pop<uint32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform4uiEXT(%" PRIu32 ", %" PRId32 ", %" PRIu32 ", %" PRIu32
                   ", %" PRIu32 ", %" PRIu32 ")",
                   program, location, v0, v1, v2, v3);
        if (glProgramUniform4uiEXT != nullptr) {
            glProgramUniform4uiEXT(program, location, v0, v1, v2, v3);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform4uiEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform4uiEXT");
        return false;
    }
}

bool callGlProgramUniform4uivEXT(Stack* stack, bool pushReturn) {
    uint32_t* value = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform4uivEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)", program,
                   location, count, value);
        if (glProgramUniform4uivEXT != nullptr) {
            glProgramUniform4uivEXT(program, location, count, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform4uivEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform4uivEXT");
        return false;
    }
}

bool callGlProgramUniformHandleui64NV(Stack* stack, bool pushReturn) {
    uint64_t value = stack->pop<uint64_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniformHandleui64NV(%" PRIu32 ", %" PRId32 ", %" PRIu64 ")", program,
                   location, value);
        if (glProgramUniformHandleui64NV != nullptr) {
            glProgramUniformHandleui64NV(program, location, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformHandleui64NV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformHandleui64NV");
        return false;
    }
}

bool callGlProgramUniformHandleui64vNV(Stack* stack, bool pushReturn) {
    uint64_t* values = stack->pop<uint64_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniformHandleui64vNV(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)",
                   program, location, count, values);
        if (glProgramUniformHandleui64vNV != nullptr) {
            glProgramUniformHandleui64vNV(program, location, count, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformHandleui64vNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformHandleui64vNV");
        return false;
    }
}

bool callGlProgramUniformMatrix2fvEXT(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniformMatrix2fvEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRIu8
                   ", %p)",
                   program, location, count, transpose, value);
        if (glProgramUniformMatrix2fvEXT != nullptr) {
            glProgramUniformMatrix2fvEXT(program, location, count, transpose, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformMatrix2fvEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix2fvEXT");
        return false;
    }
}

bool callGlProgramUniformMatrix2x3fvEXT(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniformMatrix2x3fvEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRIu8
                   ", %p)",
                   program, location, count, transpose, value);
        if (glProgramUniformMatrix2x3fvEXT != nullptr) {
            glProgramUniformMatrix2x3fvEXT(program, location, count, transpose, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformMatrix2x3fvEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix2x3fvEXT");
        return false;
    }
}

bool callGlProgramUniformMatrix2x4fvEXT(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniformMatrix2x4fvEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRIu8
                   ", %p)",
                   program, location, count, transpose, value);
        if (glProgramUniformMatrix2x4fvEXT != nullptr) {
            glProgramUniformMatrix2x4fvEXT(program, location, count, transpose, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformMatrix2x4fvEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix2x4fvEXT");
        return false;
    }
}

bool callGlProgramUniformMatrix3fvEXT(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniformMatrix3fvEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRIu8
                   ", %p)",
                   program, location, count, transpose, value);
        if (glProgramUniformMatrix3fvEXT != nullptr) {
            glProgramUniformMatrix3fvEXT(program, location, count, transpose, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformMatrix3fvEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix3fvEXT");
        return false;
    }
}

bool callGlProgramUniformMatrix3x2fvEXT(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniformMatrix3x2fvEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRIu8
                   ", %p)",
                   program, location, count, transpose, value);
        if (glProgramUniformMatrix3x2fvEXT != nullptr) {
            glProgramUniformMatrix3x2fvEXT(program, location, count, transpose, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformMatrix3x2fvEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix3x2fvEXT");
        return false;
    }
}

bool callGlProgramUniformMatrix3x4fvEXT(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniformMatrix3x4fvEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRIu8
                   ", %p)",
                   program, location, count, transpose, value);
        if (glProgramUniformMatrix3x4fvEXT != nullptr) {
            glProgramUniformMatrix3x4fvEXT(program, location, count, transpose, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformMatrix3x4fvEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix3x4fvEXT");
        return false;
    }
}

bool callGlProgramUniformMatrix4fvEXT(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniformMatrix4fvEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRIu8
                   ", %p)",
                   program, location, count, transpose, value);
        if (glProgramUniformMatrix4fvEXT != nullptr) {
            glProgramUniformMatrix4fvEXT(program, location, count, transpose, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformMatrix4fvEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix4fvEXT");
        return false;
    }
}

bool callGlProgramUniformMatrix4x2fvEXT(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniformMatrix4x2fvEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRIu8
                   ", %p)",
                   program, location, count, transpose, value);
        if (glProgramUniformMatrix4x2fvEXT != nullptr) {
            glProgramUniformMatrix4x2fvEXT(program, location, count, transpose, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformMatrix4x2fvEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix4x2fvEXT");
        return false;
    }
}

bool callGlProgramUniformMatrix4x3fvEXT(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniformMatrix4x3fvEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRIu8
                   ", %p)",
                   program, location, count, transpose, value);
        if (glProgramUniformMatrix4x3fvEXT != nullptr) {
            glProgramUniformMatrix4x3fvEXT(program, location, count, transpose, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformMatrix4x3fvEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix4x3fvEXT");
        return false;
    }
}

bool callGlPushGroupMarkerEXT(Stack* stack, bool pushReturn) {
    char* marker = stack->pop<char*>();
    int32_t length = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glPushGroupMarkerEXT(%" PRId32 ", %p)", length, marker);
        if (glPushGroupMarkerEXT != nullptr) {
            glPushGroupMarkerEXT(length, marker);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPushGroupMarkerEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPushGroupMarkerEXT");
        return false;
    }
}

bool callGlQueryCounterEXT(Stack* stack, bool pushReturn) {
    GLenum target = stack->pop<GLenum>();
    uint32_t query = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glQueryCounterEXT(%" PRIu32 ", %u)", query, target);
        if (glQueryCounterEXT != nullptr) {
            glQueryCounterEXT(query, target);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glQueryCounterEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glQueryCounterEXT");
        return false;
    }
}

bool callGlRasterSamplesEXT(Stack* stack, bool pushReturn) {
    uint8_t fixedsamplelocations = stack->pop<uint8_t>();
    uint32_t samples = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glRasterSamplesEXT(%" PRIu32 ", %" PRIu8 ")", samples, fixedsamplelocations);
        if (glRasterSamplesEXT != nullptr) {
            glRasterSamplesEXT(samples, fixedsamplelocations);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glRasterSamplesEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glRasterSamplesEXT");
        return false;
    }
}

bool callGlReadBufferIndexedEXT(Stack* stack, bool pushReturn) {
    int32_t index = stack->pop<int32_t>();
    GLenum src = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glReadBufferIndexedEXT(%u, %" PRId32 ")", src, index);
        if (glReadBufferIndexedEXT != nullptr) {
            glReadBufferIndexedEXT(src, index);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glReadBufferIndexedEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glReadBufferIndexedEXT");
        return false;
    }
}

bool callGlReadBufferNV(Stack* stack, bool pushReturn) {
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glReadBufferNV(%u)", mode);
        if (glReadBufferNV != nullptr) {
            glReadBufferNV(mode);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glReadBufferNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glReadBufferNV");
        return false;
    }
}

bool callGlReadnPixelsEXT(Stack* stack, bool pushReturn) {
    void* data = stack->pop<void*>();
    int32_t bufSize = stack->pop<int32_t>();
    GLenum type = stack->pop<GLenum>();
    GLenum format = stack->pop<GLenum>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    int32_t y = stack->pop<int32_t>();
    int32_t x = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glReadnPixelsEXT(%" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %u, %u, %" PRId32 ", %p)",
                   x, y, width, height, format, type, bufSize, data);
        if (glReadnPixelsEXT != nullptr) {
            glReadnPixelsEXT(x, y, width, height, format, type, bufSize, data);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glReadnPixelsEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glReadnPixelsEXT");
        return false;
    }
}

bool callGlReadnPixelsKHR(Stack* stack, bool pushReturn) {
    void* data = stack->pop<void*>();
    int32_t bufSize = stack->pop<int32_t>();
    GLenum type = stack->pop<GLenum>();
    GLenum format = stack->pop<GLenum>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    int32_t y = stack->pop<int32_t>();
    int32_t x = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glReadnPixelsKHR(%" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %u, %u, %" PRId32 ", %p)",
                   x, y, width, height, format, type, bufSize, data);
        if (glReadnPixelsKHR != nullptr) {
            glReadnPixelsKHR(x, y, width, height, format, type, bufSize, data);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glReadnPixelsKHR");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glReadnPixelsKHR");
        return false;
    }
}

bool callGlRenderbufferStorageMultisampleANGLE(Stack* stack, bool pushReturn) {
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    GLenum internalformat = stack->pop<GLenum>();
    int32_t samples = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glRenderbufferStorageMultisampleANGLE(%u, %" PRId32 ", %u, %" PRId32
                   ", %" PRId32 ")",
                   target, samples, internalformat, width, height);
        if (glRenderbufferStorageMultisampleANGLE != nullptr) {
            glRenderbufferStorageMultisampleANGLE(target, samples, internalformat, width, height);
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glRenderbufferStorageMultisampleANGLE");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glRenderbufferStorageMultisampleANGLE");
        return false;
    }
}

bool callGlRenderbufferStorageMultisampleAPPLE(Stack* stack, bool pushReturn) {
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    GLenum internalformat = stack->pop<GLenum>();
    int32_t samples = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glRenderbufferStorageMultisampleAPPLE(%u, %" PRId32 ", %u, %" PRId32
                   ", %" PRId32 ")",
                   target, samples, internalformat, width, height);
        if (glRenderbufferStorageMultisampleAPPLE != nullptr) {
            glRenderbufferStorageMultisampleAPPLE(target, samples, internalformat, width, height);
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glRenderbufferStorageMultisampleAPPLE");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glRenderbufferStorageMultisampleAPPLE");
        return false;
    }
}

bool callGlRenderbufferStorageMultisampleEXT(Stack* stack, bool pushReturn) {
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    GLenum internalformat = stack->pop<GLenum>();
    int32_t samples = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glRenderbufferStorageMultisampleEXT(%u, %" PRId32 ", %u, %" PRId32 ", %" PRId32
                   ")",
                   target, samples, internalformat, width, height);
        if (glRenderbufferStorageMultisampleEXT != nullptr) {
            glRenderbufferStorageMultisampleEXT(target, samples, internalformat, width, height);
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glRenderbufferStorageMultisampleEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glRenderbufferStorageMultisampleEXT");
        return false;
    }
}

bool callGlRenderbufferStorageMultisampleIMG(Stack* stack, bool pushReturn) {
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    GLenum internalformat = stack->pop<GLenum>();
    int32_t samples = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glRenderbufferStorageMultisampleIMG(%u, %" PRId32 ", %u, %" PRId32 ", %" PRId32
                   ")",
                   target, samples, internalformat, width, height);
        if (glRenderbufferStorageMultisampleIMG != nullptr) {
            glRenderbufferStorageMultisampleIMG(target, samples, internalformat, width, height);
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glRenderbufferStorageMultisampleIMG");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glRenderbufferStorageMultisampleIMG");
        return false;
    }
}

bool callGlRenderbufferStorageMultisampleNV(Stack* stack, bool pushReturn) {
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    GLenum internalformat = stack->pop<GLenum>();
    int32_t samples = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glRenderbufferStorageMultisampleNV(%u, %" PRId32 ", %u, %" PRId32 ", %" PRId32
                   ")",
                   target, samples, internalformat, width, height);
        if (glRenderbufferStorageMultisampleNV != nullptr) {
            glRenderbufferStorageMultisampleNV(target, samples, internalformat, width, height);
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glRenderbufferStorageMultisampleNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glRenderbufferStorageMultisampleNV");
        return false;
    }
}

bool callGlResolveDepthValuesNV(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glResolveDepthValuesNV()");
        if (glResolveDepthValuesNV != nullptr) {
            glResolveDepthValuesNV();
        } else {
            GAPID_WARNING("Attempted to call unsupported function glResolveDepthValuesNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glResolveDepthValuesNV");
        return false;
    }
}

bool callGlResolveMultisampleFramebufferAPPLE(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glResolveMultisampleFramebufferAPPLE()");
        if (glResolveMultisampleFramebufferAPPLE != nullptr) {
            glResolveMultisampleFramebufferAPPLE();
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glResolveMultisampleFramebufferAPPLE");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glResolveMultisampleFramebufferAPPLE");
        return false;
    }
}

bool callGlSamplerParameterIivOES(Stack* stack, bool pushReturn) {
    int32_t* param = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glSamplerParameterIivOES(%" PRIu32 ", %u, %p)", sampler, pname, param);
        if (glSamplerParameterIivOES != nullptr) {
            glSamplerParameterIivOES(sampler, pname, param);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glSamplerParameterIivOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glSamplerParameterIivOES");
        return false;
    }
}

bool callGlSamplerParameterIuivOES(Stack* stack, bool pushReturn) {
    uint32_t* param = stack->pop<uint32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glSamplerParameterIuivOES(%" PRIu32 ", %u, %p)", sampler, pname, param);
        if (glSamplerParameterIuivOES != nullptr) {
            glSamplerParameterIuivOES(sampler, pname, param);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glSamplerParameterIuivOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glSamplerParameterIuivOES");
        return false;
    }
}

bool callGlScissorArrayvNV(Stack* stack, bool pushReturn) {
    int32_t* v = stack->pop<int32_t*>();
    int32_t count = stack->pop<int32_t>();
    uint32_t first = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glScissorArrayvNV(%" PRIu32 ", %" PRId32 ", %p)", first, count, v);
        if (glScissorArrayvNV != nullptr) {
            glScissorArrayvNV(first, count, v);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glScissorArrayvNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glScissorArrayvNV");
        return false;
    }
}

bool callGlScissorIndexedNV(Stack* stack, bool pushReturn) {
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    int32_t bottom = stack->pop<int32_t>();
    int32_t left = stack->pop<int32_t>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glScissorIndexedNV(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32
                   ")",
                   index, left, bottom, width, height);
        if (glScissorIndexedNV != nullptr) {
            glScissorIndexedNV(index, left, bottom, width, height);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glScissorIndexedNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glScissorIndexedNV");
        return false;
    }
}

bool callGlScissorIndexedvNV(Stack* stack, bool pushReturn) {
    int32_t* v = stack->pop<int32_t*>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glScissorIndexedvNV(%" PRIu32 ", %p)", index, v);
        if (glScissorIndexedvNV != nullptr) {
            glScissorIndexedvNV(index, v);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glScissorIndexedvNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glScissorIndexedvNV");
        return false;
    }
}

bool callGlSelectPerfMonitorCountersAMD(Stack* stack, bool pushReturn) {
    uint32_t* counterList = stack->pop<uint32_t*>();
    int32_t numCounters = stack->pop<int32_t>();
    uint32_t group = stack->pop<uint32_t>();
    uint8_t enable = stack->pop<uint8_t>();
    uint32_t monitor = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glSelectPerfMonitorCountersAMD(%" PRIu32 ", %" PRIu8 ", %" PRIu32 ", %" PRId32
                   ", %p)",
                   monitor, enable, group, numCounters, counterList);
        if (glSelectPerfMonitorCountersAMD != nullptr) {
            glSelectPerfMonitorCountersAMD(monitor, enable, group, numCounters, counterList);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glSelectPerfMonitorCountersAMD");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glSelectPerfMonitorCountersAMD");
        return false;
    }
}

bool callGlSetFenceNV(Stack* stack, bool pushReturn) {
    GLenum condition = stack->pop<GLenum>();
    uint32_t fence = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glSetFenceNV(%" PRIu32 ", %u)", fence, condition);
        if (glSetFenceNV != nullptr) {
            glSetFenceNV(fence, condition);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glSetFenceNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glSetFenceNV");
        return false;
    }
}

bool callGlStartTilingQCOM(Stack* stack, bool pushReturn) {
    GLbitfield preserveMask = stack->pop<GLbitfield>();
    uint32_t height = stack->pop<uint32_t>();
    uint32_t width = stack->pop<uint32_t>();
    uint32_t y = stack->pop<uint32_t>();
    uint32_t x = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glStartTilingQCOM(%" PRIu32 ", %" PRIu32 ", %" PRIu32 ", %" PRIu32 ", %u)", x,
                   y, width, height, preserveMask);
        if (glStartTilingQCOM != nullptr) {
            glStartTilingQCOM(x, y, width, height, preserveMask);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glStartTilingQCOM");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStartTilingQCOM");
        return false;
    }
}

bool callGlStencilFillPathInstancedNV(Stack* stack, bool pushReturn) {
    float* transformValues = stack->pop<float*>();
    GLenum transformType = stack->pop<GLenum>();
    uint32_t mask = stack->pop<uint32_t>();
    GLenum fillMode = stack->pop<GLenum>();
    uint32_t pathBase = stack->pop<uint32_t>();
    void* paths = stack->pop<void*>();
    GLenum pathNameType = stack->pop<GLenum>();
    int32_t numPaths = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glStencilFillPathInstancedNV(%" PRId32 ", %u, %p, %" PRIu32 ", %u, %" PRIu32
                   ", %u, %p)",
                   numPaths, pathNameType, paths, pathBase, fillMode, mask, transformType,
                   transformValues);
        if (glStencilFillPathInstancedNV != nullptr) {
            glStencilFillPathInstancedNV(numPaths, pathNameType, paths, pathBase, fillMode, mask,
                                         transformType, transformValues);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glStencilFillPathInstancedNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStencilFillPathInstancedNV");
        return false;
    }
}

bool callGlStencilFillPathNV(Stack* stack, bool pushReturn) {
    uint32_t mask = stack->pop<uint32_t>();
    GLenum fillMode = stack->pop<GLenum>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glStencilFillPathNV(%" PRIu32 ", %u, %" PRIu32 ")", path, fillMode, mask);
        if (glStencilFillPathNV != nullptr) {
            glStencilFillPathNV(path, fillMode, mask);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glStencilFillPathNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStencilFillPathNV");
        return false;
    }
}

bool callGlStencilStrokePathInstancedNV(Stack* stack, bool pushReturn) {
    float* transformValues = stack->pop<float*>();
    GLenum transformType = stack->pop<GLenum>();
    uint32_t mask = stack->pop<uint32_t>();
    int32_t reference = stack->pop<int32_t>();
    uint32_t pathBase = stack->pop<uint32_t>();
    void* paths = stack->pop<void*>();
    GLenum pathNameType = stack->pop<GLenum>();
    int32_t numPaths = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glStencilStrokePathInstancedNV(%" PRId32 ", %u, %p, %" PRIu32 ", %" PRId32
                   ", %" PRIu32 ", %u, %p)",
                   numPaths, pathNameType, paths, pathBase, reference, mask, transformType,
                   transformValues);
        if (glStencilStrokePathInstancedNV != nullptr) {
            glStencilStrokePathInstancedNV(numPaths, pathNameType, paths, pathBase, reference, mask,
                                           transformType, transformValues);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glStencilStrokePathInstancedNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStencilStrokePathInstancedNV");
        return false;
    }
}

bool callGlStencilStrokePathNV(Stack* stack, bool pushReturn) {
    uint32_t mask = stack->pop<uint32_t>();
    int32_t reference = stack->pop<int32_t>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glStencilStrokePathNV(%" PRIu32 ", %" PRId32 ", %" PRIu32 ")", path, reference,
                   mask);
        if (glStencilStrokePathNV != nullptr) {
            glStencilStrokePathNV(path, reference, mask);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glStencilStrokePathNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStencilStrokePathNV");
        return false;
    }
}

bool callGlStencilThenCoverFillPathInstancedNV(Stack* stack, bool pushReturn) {
    float* transformValues = stack->pop<float*>();
    GLenum transformType = stack->pop<GLenum>();
    GLenum coverMode = stack->pop<GLenum>();
    uint32_t mask = stack->pop<uint32_t>();
    GLenum fillMode = stack->pop<GLenum>();
    uint32_t pathBase = stack->pop<uint32_t>();
    void* paths = stack->pop<void*>();
    GLenum pathNameType = stack->pop<GLenum>();
    int32_t numPaths = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glStencilThenCoverFillPathInstancedNV(%" PRId32 ", %u, %p, %" PRIu32
                   ", %u, %" PRIu32 ", %u, %u, %p)",
                   numPaths, pathNameType, paths, pathBase, fillMode, mask, coverMode,
                   transformType, transformValues);
        if (glStencilThenCoverFillPathInstancedNV != nullptr) {
            glStencilThenCoverFillPathInstancedNV(numPaths, pathNameType, paths, pathBase, fillMode,
                                                  mask, coverMode, transformType, transformValues);
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glStencilThenCoverFillPathInstancedNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStencilThenCoverFillPathInstancedNV");
        return false;
    }
}

bool callGlStencilThenCoverFillPathNV(Stack* stack, bool pushReturn) {
    GLenum coverMode = stack->pop<GLenum>();
    uint32_t mask = stack->pop<uint32_t>();
    GLenum fillMode = stack->pop<GLenum>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glStencilThenCoverFillPathNV(%" PRIu32 ", %u, %" PRIu32 ", %u)", path, fillMode,
                   mask, coverMode);
        if (glStencilThenCoverFillPathNV != nullptr) {
            glStencilThenCoverFillPathNV(path, fillMode, mask, coverMode);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glStencilThenCoverFillPathNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStencilThenCoverFillPathNV");
        return false;
    }
}

bool callGlStencilThenCoverStrokePathInstancedNV(Stack* stack, bool pushReturn) {
    float* transformValues = stack->pop<float*>();
    GLenum transformType = stack->pop<GLenum>();
    GLenum coverMode = stack->pop<GLenum>();
    uint32_t mask = stack->pop<uint32_t>();
    int32_t reference = stack->pop<int32_t>();
    uint32_t pathBase = stack->pop<uint32_t>();
    void* paths = stack->pop<void*>();
    GLenum pathNameType = stack->pop<GLenum>();
    int32_t numPaths = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glStencilThenCoverStrokePathInstancedNV(%" PRId32 ", %u, %p, %" PRIu32
                   ", %" PRId32 ", %" PRIu32 ", %u, %u, %p)",
                   numPaths, pathNameType, paths, pathBase, reference, mask, coverMode,
                   transformType, transformValues);
        if (glStencilThenCoverStrokePathInstancedNV != nullptr) {
            glStencilThenCoverStrokePathInstancedNV(numPaths, pathNameType, paths, pathBase,
                                                    reference, mask, coverMode, transformType,
                                                    transformValues);
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function "
                    "glStencilThenCoverStrokePathInstancedNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStencilThenCoverStrokePathInstancedNV");
        return false;
    }
}

bool callGlStencilThenCoverStrokePathNV(Stack* stack, bool pushReturn) {
    GLenum coverMode = stack->pop<GLenum>();
    uint32_t mask = stack->pop<uint32_t>();
    int32_t reference = stack->pop<int32_t>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glStencilThenCoverStrokePathNV(%" PRIu32 ", %" PRId32 ", %" PRIu32 ", %u)",
                   path, reference, mask, coverMode);
        if (glStencilThenCoverStrokePathNV != nullptr) {
            glStencilThenCoverStrokePathNV(path, reference, mask, coverMode);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glStencilThenCoverStrokePathNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStencilThenCoverStrokePathNV");
        return false;
    }
}

bool callGlSubpixelPrecisionBiasNV(Stack* stack, bool pushReturn) {
    uint32_t ybits = stack->pop<uint32_t>();
    uint32_t xbits = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glSubpixelPrecisionBiasNV(%" PRIu32 ", %" PRIu32 ")", xbits, ybits);
        if (glSubpixelPrecisionBiasNV != nullptr) {
            glSubpixelPrecisionBiasNV(xbits, ybits);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glSubpixelPrecisionBiasNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glSubpixelPrecisionBiasNV");
        return false;
    }
}

bool callGlTestFenceNV(Stack* stack, bool pushReturn) {
    uint32_t fence = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glTestFenceNV(%" PRIu32 ")", fence);
        if (glTestFenceNV != nullptr) {
            uint8_t return_value = glTestFenceNV(fence);
            GAPID_INFO("Returned: %" PRIu8 "", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTestFenceNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTestFenceNV");
        return false;
    }
}

bool callGlTexBufferOES(Stack* stack, bool pushReturn) {
    uint32_t buffer = stack->pop<uint32_t>();
    GLenum internalformat = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexBufferOES(%u, %u, %" PRIu32 ")", target, internalformat, buffer);
        if (glTexBufferOES != nullptr) {
            glTexBufferOES(target, internalformat, buffer);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexBufferOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexBufferOES");
        return false;
    }
}

bool callGlTexBufferRangeOES(Stack* stack, bool pushReturn) {
    int32_t size = stack->pop<int32_t>();
    int32_t offset = stack->pop<int32_t>();
    uint32_t buffer = stack->pop<uint32_t>();
    GLenum internalformat = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexBufferRangeOES(%u, %u, %" PRIu32 ", %" PRId32 ", %" PRId32 ")", target,
                   internalformat, buffer, offset, size);
        if (glTexBufferRangeOES != nullptr) {
            glTexBufferRangeOES(target, internalformat, buffer, offset, size);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexBufferRangeOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexBufferRangeOES");
        return false;
    }
}

bool callGlTexImage3DOES(Stack* stack, bool pushReturn) {
    void* pixels = stack->pop<void*>();
    GLenum type = stack->pop<GLenum>();
    GLenum format = stack->pop<GLenum>();
    int32_t border = stack->pop<int32_t>();
    int32_t depth = stack->pop<int32_t>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    GLenum internalformat = stack->pop<GLenum>();
    int32_t level = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexImage3DOES(%u, %" PRId32 ", %u, %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %u, %u, %p)",
                   target, level, internalformat, width, height, depth, border, format, type,
                   pixels);
        if (glTexImage3DOES != nullptr) {
            glTexImage3DOES(target, level, internalformat, width, height, depth, border, format,
                            type, pixels);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexImage3DOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexImage3DOES");
        return false;
    }
}

bool callGlTexPageCommitmentARB(Stack* stack, bool pushReturn) {
    uint8_t commit = stack->pop<uint8_t>();
    int32_t depth = stack->pop<int32_t>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    int32_t zoffset = stack->pop<int32_t>();
    int32_t yoffset = stack->pop<int32_t>();
    int32_t xoffset = stack->pop<int32_t>();
    int32_t level = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexPageCommitmentARB(%u, %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRIu8 ")",
                   target, level, xoffset, yoffset, zoffset, width, height, depth, commit);
        if (glTexPageCommitmentARB != nullptr) {
            glTexPageCommitmentARB(target, level, xoffset, yoffset, zoffset, width, height, depth,
                                   commit);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexPageCommitmentARB");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexPageCommitmentARB");
        return false;
    }
}

bool callGlTexParameterIivOES(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexParameterIivOES(%u, %u, %p)", target, pname, params);
        if (glTexParameterIivOES != nullptr) {
            glTexParameterIivOES(target, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexParameterIivOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexParameterIivOES");
        return false;
    }
}

bool callGlTexParameterIuivOES(Stack* stack, bool pushReturn) {
    uint32_t* params = stack->pop<uint32_t*>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexParameterIuivOES(%u, %u, %p)", target, pname, params);
        if (glTexParameterIuivOES != nullptr) {
            glTexParameterIuivOES(target, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexParameterIuivOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexParameterIuivOES");
        return false;
    }
}

bool callGlTexStorage1DEXT(Stack* stack, bool pushReturn) {
    int32_t width = stack->pop<int32_t>();
    GLenum format = stack->pop<GLenum>();
    int32_t levels = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexStorage1DEXT(%u, %" PRId32 ", %u, %" PRId32 ")", target, levels, format,
                   width);
        if (glTexStorage1DEXT != nullptr) {
            glTexStorage1DEXT(target, levels, format, width);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexStorage1DEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexStorage1DEXT");
        return false;
    }
}

bool callGlTexStorage2DEXT(Stack* stack, bool pushReturn) {
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    GLenum format = stack->pop<GLenum>();
    int32_t levels = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexStorage2DEXT(%u, %" PRId32 ", %u, %" PRId32 ", %" PRId32 ")", target,
                   levels, format, width, height);
        if (glTexStorage2DEXT != nullptr) {
            glTexStorage2DEXT(target, levels, format, width, height);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexStorage2DEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexStorage2DEXT");
        return false;
    }
}

bool callGlTexStorage3DEXT(Stack* stack, bool pushReturn) {
    int32_t depth = stack->pop<int32_t>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    GLenum format = stack->pop<GLenum>();
    int32_t levels = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexStorage3DEXT(%u, %" PRId32 ", %u, %" PRId32 ", %" PRId32 ", %" PRId32 ")",
                   target, levels, format, width, height, depth);
        if (glTexStorage3DEXT != nullptr) {
            glTexStorage3DEXT(target, levels, format, width, height, depth);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexStorage3DEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexStorage3DEXT");
        return false;
    }
}

bool callGlTexSubImage3DOES(Stack* stack, bool pushReturn) {
    void* pixels = stack->pop<void*>();
    GLenum type = stack->pop<GLenum>();
    GLenum format = stack->pop<GLenum>();
    int32_t depth = stack->pop<int32_t>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    int32_t zoffset = stack->pop<int32_t>();
    int32_t yoffset = stack->pop<int32_t>();
    int32_t xoffset = stack->pop<int32_t>();
    int32_t level = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexSubImage3DOES(%u, %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %u, %u, %p)",
                   target, level, xoffset, yoffset, zoffset, width, height, depth, format, type,
                   pixels);
        if (glTexSubImage3DOES != nullptr) {
            glTexSubImage3DOES(target, level, xoffset, yoffset, zoffset, width, height, depth,
                               format, type, pixels);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexSubImage3DOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexSubImage3DOES");
        return false;
    }
}

bool callGlTextureStorage1DEXT(Stack* stack, bool pushReturn) {
    int32_t width = stack->pop<int32_t>();
    GLenum format = stack->pop<GLenum>();
    int32_t levels = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    uint32_t texture = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glTextureStorage1DEXT(%" PRIu32 ", %u, %" PRId32 ", %u, %" PRId32 ")", texture,
                   target, levels, format, width);
        if (glTextureStorage1DEXT != nullptr) {
            glTextureStorage1DEXT(texture, target, levels, format, width);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTextureStorage1DEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTextureStorage1DEXT");
        return false;
    }
}

bool callGlTextureStorage2DEXT(Stack* stack, bool pushReturn) {
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    GLenum format = stack->pop<GLenum>();
    int32_t levels = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    uint32_t texture = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glTextureStorage2DEXT(%" PRIu32 ", %u, %" PRId32 ", %u, %" PRId32 ", %" PRId32
                   ")",
                   texture, target, levels, format, width, height);
        if (glTextureStorage2DEXT != nullptr) {
            glTextureStorage2DEXT(texture, target, levels, format, width, height);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTextureStorage2DEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTextureStorage2DEXT");
        return false;
    }
}

bool callGlTextureStorage3DEXT(Stack* stack, bool pushReturn) {
    int32_t depth = stack->pop<int32_t>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    GLenum format = stack->pop<GLenum>();
    int32_t levels = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    uint32_t texture = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glTextureStorage3DEXT(%" PRIu32 ", %u, %" PRId32 ", %u, %" PRId32 ", %" PRId32
                   ", %" PRId32 ")",
                   texture, target, levels, format, width, height, depth);
        if (glTextureStorage3DEXT != nullptr) {
            glTextureStorage3DEXT(texture, target, levels, format, width, height, depth);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTextureStorage3DEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTextureStorage3DEXT");
        return false;
    }
}

bool callGlTextureViewEXT(Stack* stack, bool pushReturn) {
    uint32_t numlayers = stack->pop<uint32_t>();
    uint32_t minlayer = stack->pop<uint32_t>();
    uint32_t numlevels = stack->pop<uint32_t>();
    uint32_t minlevel = stack->pop<uint32_t>();
    GLenum internalformat = stack->pop<GLenum>();
    uint32_t origtexture = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    uint32_t texture = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glTextureViewEXT(%" PRIu32 ", %u, %" PRIu32 ", %u, %" PRIu32 ", %" PRIu32
                   ", %" PRIu32 ", %" PRIu32 ")",
                   texture, target, origtexture, internalformat, minlevel, numlevels, minlayer,
                   numlayers);
        if (glTextureViewEXT != nullptr) {
            glTextureViewEXT(texture, target, origtexture, internalformat, minlevel, numlevels,
                             minlayer, numlayers);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTextureViewEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTextureViewEXT");
        return false;
    }
}

bool callGlTextureViewOES(Stack* stack, bool pushReturn) {
    uint32_t numlayers = stack->pop<uint32_t>();
    uint32_t minlayer = stack->pop<uint32_t>();
    uint32_t numlevels = stack->pop<uint32_t>();
    uint32_t minlevel = stack->pop<uint32_t>();
    GLenum internalformat = stack->pop<GLenum>();
    uint32_t origtexture = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    uint32_t texture = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glTextureViewOES(%" PRIu32 ", %u, %" PRIu32 ", %u, %" PRIu32 ", %" PRIu32
                   ", %" PRIu32 ", %" PRIu32 ")",
                   texture, target, origtexture, internalformat, minlevel, numlevels, minlayer,
                   numlayers);
        if (glTextureViewOES != nullptr) {
            glTextureViewOES(texture, target, origtexture, internalformat, minlevel, numlevels,
                             minlayer, numlayers);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTextureViewOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTextureViewOES");
        return false;
    }
}

bool callGlTransformPathNV(Stack* stack, bool pushReturn) {
    float* transformValues = stack->pop<float*>();
    GLenum transformType = stack->pop<GLenum>();
    uint32_t srcPath = stack->pop<uint32_t>();
    uint32_t resultPath = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glTransformPathNV(%" PRIu32 ", %" PRIu32 ", %u, %p)", resultPath, srcPath,
                   transformType, transformValues);
        if (glTransformPathNV != nullptr) {
            glTransformPathNV(resultPath, srcPath, transformType, transformValues);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTransformPathNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTransformPathNV");
        return false;
    }
}

bool callGlUniformHandleui64NV(Stack* stack, bool pushReturn) {
    uint64_t value = stack->pop<uint64_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformHandleui64NV(%" PRId32 ", %" PRIu64 ")", location, value);
        if (glUniformHandleui64NV != nullptr) {
            glUniformHandleui64NV(location, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformHandleui64NV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformHandleui64NV");
        return false;
    }
}

bool callGlUniformHandleui64vNV(Stack* stack, bool pushReturn) {
    uint64_t* value = stack->pop<uint64_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformHandleui64vNV(%" PRId32 ", %" PRId32 ", %p)", location, count, value);
        if (glUniformHandleui64vNV != nullptr) {
            glUniformHandleui64vNV(location, count, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformHandleui64vNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformHandleui64vNV");
        return false;
    }
}

bool callGlUniformMatrix2x3fvNV(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformMatrix2x3fvNV(%" PRId32 ", %" PRId32 ", %" PRIu8 ", %p)", location,
                   count, transpose, value);
        if (glUniformMatrix2x3fvNV != nullptr) {
            glUniformMatrix2x3fvNV(location, count, transpose, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformMatrix2x3fvNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformMatrix2x3fvNV");
        return false;
    }
}

bool callGlUniformMatrix2x4fvNV(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformMatrix2x4fvNV(%" PRId32 ", %" PRId32 ", %" PRIu8 ", %p)", location,
                   count, transpose, value);
        if (glUniformMatrix2x4fvNV != nullptr) {
            glUniformMatrix2x4fvNV(location, count, transpose, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformMatrix2x4fvNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformMatrix2x4fvNV");
        return false;
    }
}

bool callGlUniformMatrix3x2fvNV(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformMatrix3x2fvNV(%" PRId32 ", %" PRId32 ", %" PRIu8 ", %p)", location,
                   count, transpose, value);
        if (glUniformMatrix3x2fvNV != nullptr) {
            glUniformMatrix3x2fvNV(location, count, transpose, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformMatrix3x2fvNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformMatrix3x2fvNV");
        return false;
    }
}

bool callGlUniformMatrix3x4fvNV(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformMatrix3x4fvNV(%" PRId32 ", %" PRId32 ", %" PRIu8 ", %p)", location,
                   count, transpose, value);
        if (glUniformMatrix3x4fvNV != nullptr) {
            glUniformMatrix3x4fvNV(location, count, transpose, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformMatrix3x4fvNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformMatrix3x4fvNV");
        return false;
    }
}

bool callGlUniformMatrix4x2fvNV(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformMatrix4x2fvNV(%" PRId32 ", %" PRId32 ", %" PRIu8 ", %p)", location,
                   count, transpose, value);
        if (glUniformMatrix4x2fvNV != nullptr) {
            glUniformMatrix4x2fvNV(location, count, transpose, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformMatrix4x2fvNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformMatrix4x2fvNV");
        return false;
    }
}

bool callGlUniformMatrix4x3fvNV(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformMatrix4x3fvNV(%" PRId32 ", %" PRId32 ", %" PRIu8 ", %p)", location,
                   count, transpose, value);
        if (glUniformMatrix4x3fvNV != nullptr) {
            glUniformMatrix4x3fvNV(location, count, transpose, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformMatrix4x3fvNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformMatrix4x3fvNV");
        return false;
    }
}

bool callGlUnmapBufferOES(Stack* stack, bool pushReturn) {
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glUnmapBufferOES(%u)", target);
        if (glUnmapBufferOES != nullptr) {
            uint8_t return_value = glUnmapBufferOES(target);
            GAPID_INFO("Returned: %" PRIu8 "", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUnmapBufferOES");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUnmapBufferOES");
        return false;
    }
}

bool callGlUseProgramStagesEXT(Stack* stack, bool pushReturn) {
    uint32_t program = stack->pop<uint32_t>();
    GLbitfield stages = stack->pop<GLbitfield>();
    uint32_t pipeline = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUseProgramStagesEXT(%" PRIu32 ", %u, %" PRIu32 ")", pipeline, stages,
                   program);
        if (glUseProgramStagesEXT != nullptr) {
            glUseProgramStagesEXT(pipeline, stages, program);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUseProgramStagesEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUseProgramStagesEXT");
        return false;
    }
}

bool callGlValidateProgramPipelineEXT(Stack* stack, bool pushReturn) {
    uint32_t pipeline = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glValidateProgramPipelineEXT(%" PRIu32 ")", pipeline);
        if (glValidateProgramPipelineEXT != nullptr) {
            glValidateProgramPipelineEXT(pipeline);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glValidateProgramPipelineEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glValidateProgramPipelineEXT");
        return false;
    }
}

bool callGlVertexAttribDivisorANGLE(Stack* stack, bool pushReturn) {
    uint32_t divisor = stack->pop<uint32_t>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttribDivisorANGLE(%" PRIu32 ", %" PRIu32 ")", index, divisor);
        if (glVertexAttribDivisorANGLE != nullptr) {
            glVertexAttribDivisorANGLE(index, divisor);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttribDivisorANGLE");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttribDivisorANGLE");
        return false;
    }
}

bool callGlVertexAttribDivisorEXT(Stack* stack, bool pushReturn) {
    uint32_t divisor = stack->pop<uint32_t>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttribDivisorEXT(%" PRIu32 ", %" PRIu32 ")", index, divisor);
        if (glVertexAttribDivisorEXT != nullptr) {
            glVertexAttribDivisorEXT(index, divisor);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttribDivisorEXT");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttribDivisorEXT");
        return false;
    }
}

bool callGlVertexAttribDivisorNV(Stack* stack, bool pushReturn) {
    uint32_t divisor = stack->pop<uint32_t>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttribDivisorNV(%" PRIu32 ", %" PRIu32 ")", index, divisor);
        if (glVertexAttribDivisorNV != nullptr) {
            glVertexAttribDivisorNV(index, divisor);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttribDivisorNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttribDivisorNV");
        return false;
    }
}

bool callGlViewportArrayvNV(Stack* stack, bool pushReturn) {
    float* v = stack->pop<float*>();
    int32_t count = stack->pop<int32_t>();
    uint32_t first = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glViewportArrayvNV(%" PRIu32 ", %" PRId32 ", %p)", first, count, v);
        if (glViewportArrayvNV != nullptr) {
            glViewportArrayvNV(first, count, v);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glViewportArrayvNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glViewportArrayvNV");
        return false;
    }
}

bool callGlViewportIndexedfNV(Stack* stack, bool pushReturn) {
    float h = stack->pop<float>();
    float w = stack->pop<float>();
    float y = stack->pop<float>();
    float x = stack->pop<float>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glViewportIndexedfNV(%" PRIu32 ", %f, %f, %f, %f)", index, x, y, w, h);
        if (glViewportIndexedfNV != nullptr) {
            glViewportIndexedfNV(index, x, y, w, h);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glViewportIndexedfNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glViewportIndexedfNV");
        return false;
    }
}

bool callGlViewportIndexedfvNV(Stack* stack, bool pushReturn) {
    float* v = stack->pop<float*>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glViewportIndexedfvNV(%" PRIu32 ", %p)", index, v);
        if (glViewportIndexedfvNV != nullptr) {
            glViewportIndexedfvNV(index, v);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glViewportIndexedfvNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glViewportIndexedfvNV");
        return false;
    }
}

bool callGlWaitSyncAPPLE(Stack* stack, bool pushReturn) {
    uint64_t timeout = stack->pop<uint64_t>();
    GLbitfield flag = stack->pop<GLbitfield>();
    uint64_t sync = stack->pop<uint64_t>();
    if (stack->isValid()) {
        GAPID_INFO("glWaitSyncAPPLE(%" PRIu64 ", %u, %" PRIu64 ")", sync, flag, timeout);
        if (glWaitSyncAPPLE != nullptr) {
            glWaitSyncAPPLE(sync, flag, timeout);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glWaitSyncAPPLE");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glWaitSyncAPPLE");
        return false;
    }
}

bool callGlWeightPathsNV(Stack* stack, bool pushReturn) {
    float* weights = stack->pop<float*>();
    uint32_t* paths = stack->pop<uint32_t*>();
    int32_t numPaths = stack->pop<int32_t>();
    uint32_t resultPath = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glWeightPathsNV(%" PRIu32 ", %" PRId32 ", %p, %p)", resultPath, numPaths, paths,
                   weights);
        if (glWeightPathsNV != nullptr) {
            glWeightPathsNV(resultPath, numPaths, paths, weights);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glWeightPathsNV");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glWeightPathsNV");
        return false;
    }
}

bool callGlBlendBarrier(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glBlendBarrier()");
        if (glBlendBarrier != nullptr) {
            glBlendBarrier();
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendBarrier");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendBarrier");
        return false;
    }
}

bool callGlBlendColor(Stack* stack, bool pushReturn) {
    float alpha = stack->pop<float>();
    float blue = stack->pop<float>();
    float green = stack->pop<float>();
    float red = stack->pop<float>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendColor(%f, %f, %f, %f)", red, green, blue, alpha);
        if (glBlendColor != nullptr) {
            glBlendColor(red, green, blue, alpha);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendColor");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendColor");
        return false;
    }
}

bool callGlBlendEquation(Stack* stack, bool pushReturn) {
    GLenum equation = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendEquation(%u)", equation);
        if (glBlendEquation != nullptr) {
            glBlendEquation(equation);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendEquation");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendEquation");
        return false;
    }
}

bool callGlBlendEquationSeparate(Stack* stack, bool pushReturn) {
    GLenum alpha = stack->pop<GLenum>();
    GLenum rgb = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendEquationSeparate(%u, %u)", rgb, alpha);
        if (glBlendEquationSeparate != nullptr) {
            glBlendEquationSeparate(rgb, alpha);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendEquationSeparate");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendEquationSeparate");
        return false;
    }
}

bool callGlBlendEquationSeparatei(Stack* stack, bool pushReturn) {
    GLenum modeAlpha = stack->pop<GLenum>();
    GLenum modeRGB = stack->pop<GLenum>();
    uint32_t buf = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendEquationSeparatei(%" PRIu32 ", %u, %u)", buf, modeRGB, modeAlpha);
        if (glBlendEquationSeparatei != nullptr) {
            glBlendEquationSeparatei(buf, modeRGB, modeAlpha);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendEquationSeparatei");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendEquationSeparatei");
        return false;
    }
}

bool callGlBlendEquationi(Stack* stack, bool pushReturn) {
    GLenum mode = stack->pop<GLenum>();
    uint32_t buf = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendEquationi(%" PRIu32 ", %u)", buf, mode);
        if (glBlendEquationi != nullptr) {
            glBlendEquationi(buf, mode);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendEquationi");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendEquationi");
        return false;
    }
}

bool callGlBlendFunc(Stack* stack, bool pushReturn) {
    GLenum dst_factor = stack->pop<GLenum>();
    GLenum src_factor = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendFunc(%u, %u)", src_factor, dst_factor);
        if (glBlendFunc != nullptr) {
            glBlendFunc(src_factor, dst_factor);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendFunc");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendFunc");
        return false;
    }
}

bool callGlBlendFuncSeparate(Stack* stack, bool pushReturn) {
    GLenum dst_factor_alpha = stack->pop<GLenum>();
    GLenum src_factor_alpha = stack->pop<GLenum>();
    GLenum dst_factor_rgb = stack->pop<GLenum>();
    GLenum src_factor_rgb = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendFuncSeparate(%u, %u, %u, %u)", src_factor_rgb, dst_factor_rgb,
                   src_factor_alpha, dst_factor_alpha);
        if (glBlendFuncSeparate != nullptr) {
            glBlendFuncSeparate(src_factor_rgb, dst_factor_rgb, src_factor_alpha, dst_factor_alpha);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendFuncSeparate");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendFuncSeparate");
        return false;
    }
}

bool callGlBlendFuncSeparatei(Stack* stack, bool pushReturn) {
    GLenum dstAlpha = stack->pop<GLenum>();
    GLenum srcAlpha = stack->pop<GLenum>();
    GLenum dstRGB = stack->pop<GLenum>();
    GLenum srcRGB = stack->pop<GLenum>();
    uint32_t buf = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendFuncSeparatei(%" PRIu32 ", %u, %u, %u, %u)", buf, srcRGB, dstRGB,
                   srcAlpha, dstAlpha);
        if (glBlendFuncSeparatei != nullptr) {
            glBlendFuncSeparatei(buf, srcRGB, dstRGB, srcAlpha, dstAlpha);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendFuncSeparatei");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendFuncSeparatei");
        return false;
    }
}

bool callGlBlendFunci(Stack* stack, bool pushReturn) {
    GLenum dst = stack->pop<GLenum>();
    GLenum src = stack->pop<GLenum>();
    uint32_t buf = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendFunci(%" PRIu32 ", %u, %u)", buf, src, dst);
        if (glBlendFunci != nullptr) {
            glBlendFunci(buf, src, dst);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendFunci");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendFunci");
        return false;
    }
}

bool callGlDepthFunc(Stack* stack, bool pushReturn) {
    GLenum function = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDepthFunc(%u)", function);
        if (glDepthFunc != nullptr) {
            glDepthFunc(function);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDepthFunc");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDepthFunc");
        return false;
    }
}

bool callGlSampleCoverage(Stack* stack, bool pushReturn) {
    uint8_t invert = stack->pop<uint8_t>();
    float value = stack->pop<float>();
    if (stack->isValid()) {
        GAPID_INFO("glSampleCoverage(%f, %" PRIu8 ")", value, invert);
        if (glSampleCoverage != nullptr) {
            glSampleCoverage(value, invert);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glSampleCoverage");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glSampleCoverage");
        return false;
    }
}

bool callGlSampleMaski(Stack* stack, bool pushReturn) {
    GLbitfield mask = stack->pop<GLbitfield>();
    uint32_t maskNumber = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glSampleMaski(%" PRIu32 ", %u)", maskNumber, mask);
        if (glSampleMaski != nullptr) {
            glSampleMaski(maskNumber, mask);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glSampleMaski");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glSampleMaski");
        return false;
    }
}

bool callGlScissor(Stack* stack, bool pushReturn) {
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    int32_t y = stack->pop<int32_t>();
    int32_t x = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glScissor(%" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ")", x, y, width,
                   height);
        if (glScissor != nullptr) {
            glScissor(x, y, width, height);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glScissor");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glScissor");
        return false;
    }
}

bool callGlStencilFunc(Stack* stack, bool pushReturn) {
    uint32_t mask = stack->pop<uint32_t>();
    int32_t ref = stack->pop<int32_t>();
    GLenum func = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glStencilFunc(%u, %" PRId32 ", %" PRIu32 ")", func, ref, mask);
        if (glStencilFunc != nullptr) {
            glStencilFunc(func, ref, mask);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glStencilFunc");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStencilFunc");
        return false;
    }
}

bool callGlStencilFuncSeparate(Stack* stack, bool pushReturn) {
    uint32_t mask = stack->pop<uint32_t>();
    int32_t reference_value = stack->pop<int32_t>();
    GLenum function = stack->pop<GLenum>();
    GLenum face = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glStencilFuncSeparate(%u, %u, %" PRId32 ", %" PRIu32 ")", face, function,
                   reference_value, mask);
        if (glStencilFuncSeparate != nullptr) {
            glStencilFuncSeparate(face, function, reference_value, mask);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glStencilFuncSeparate");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStencilFuncSeparate");
        return false;
    }
}

bool callGlStencilOp(Stack* stack, bool pushReturn) {
    GLenum zpass = stack->pop<GLenum>();
    GLenum zfail = stack->pop<GLenum>();
    GLenum fail = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glStencilOp(%u, %u, %u)", fail, zfail, zpass);
        if (glStencilOp != nullptr) {
            glStencilOp(fail, zfail, zpass);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glStencilOp");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStencilOp");
        return false;
    }
}

bool callGlStencilOpSeparate(Stack* stack, bool pushReturn) {
    GLenum stencil_pass_depth_pass = stack->pop<GLenum>();
    GLenum stencil_pass_depth_fail = stack->pop<GLenum>();
    GLenum stencil_fail = stack->pop<GLenum>();
    GLenum face = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glStencilOpSeparate(%u, %u, %u, %u)", face, stencil_fail,
                   stencil_pass_depth_fail, stencil_pass_depth_pass);
        if (glStencilOpSeparate != nullptr) {
            glStencilOpSeparate(face, stencil_fail, stencil_pass_depth_fail,
                                stencil_pass_depth_pass);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glStencilOpSeparate");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStencilOpSeparate");
        return false;
    }
}

bool callGlBindFramebuffer(Stack* stack, bool pushReturn) {
    uint32_t framebuffer = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBindFramebuffer(%u, %" PRIu32 ")", target, framebuffer);
        if (glBindFramebuffer != nullptr) {
            glBindFramebuffer(target, framebuffer);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindFramebuffer");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindFramebuffer");
        return false;
    }
}

bool callGlBindRenderbuffer(Stack* stack, bool pushReturn) {
    uint32_t renderbuffer = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBindRenderbuffer(%u, %" PRIu32 ")", target, renderbuffer);
        if (glBindRenderbuffer != nullptr) {
            glBindRenderbuffer(target, renderbuffer);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindRenderbuffer");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindRenderbuffer");
        return false;
    }
}

bool callGlBlitFramebuffer(Stack* stack, bool pushReturn) {
    GLenum filter = stack->pop<GLenum>();
    GLbitfield mask = stack->pop<GLbitfield>();
    int32_t dstY1 = stack->pop<int32_t>();
    int32_t dstX1 = stack->pop<int32_t>();
    int32_t dstY0 = stack->pop<int32_t>();
    int32_t dstX0 = stack->pop<int32_t>();
    int32_t srcY1 = stack->pop<int32_t>();
    int32_t srcX1 = stack->pop<int32_t>();
    int32_t srcY0 = stack->pop<int32_t>();
    int32_t srcX0 = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBlitFramebuffer(%" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %u, %u)",
                   srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter);
        if (glBlitFramebuffer != nullptr) {
            glBlitFramebuffer(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlitFramebuffer");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlitFramebuffer");
        return false;
    }
}

bool callGlCheckFramebufferStatus(Stack* stack, bool pushReturn) {
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glCheckFramebufferStatus(%u)", target);
        if (glCheckFramebufferStatus != nullptr) {
            GLenum return_value = glCheckFramebufferStatus(target);
            GAPID_INFO("Returned: %u", return_value);
            if (pushReturn) {
                stack->push<GLenum>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCheckFramebufferStatus");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCheckFramebufferStatus");
        return false;
    }
}

bool callGlClear(Stack* stack, bool pushReturn) {
    GLbitfield mask = stack->pop<GLbitfield>();
    if (stack->isValid()) {
        GAPID_INFO("glClear(%u)", mask);
        if (glClear != nullptr) {
            glClear(mask);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glClear");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glClear");
        return false;
    }
}

bool callGlClearBufferfi(Stack* stack, bool pushReturn) {
    int32_t stencil = stack->pop<int32_t>();
    float depth = stack->pop<float>();
    int32_t drawbuffer = stack->pop<int32_t>();
    GLenum buffer = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glClearBufferfi(%u, %" PRId32 ", %f, %" PRId32 ")", buffer, drawbuffer, depth,
                   stencil);
        if (glClearBufferfi != nullptr) {
            glClearBufferfi(buffer, drawbuffer, depth, stencil);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glClearBufferfi");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glClearBufferfi");
        return false;
    }
}

bool callGlClearBufferfv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    int32_t drawbuffer = stack->pop<int32_t>();
    GLenum buffer = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glClearBufferfv(%u, %" PRId32 ", %p)", buffer, drawbuffer, value);
        if (glClearBufferfv != nullptr) {
            glClearBufferfv(buffer, drawbuffer, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glClearBufferfv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glClearBufferfv");
        return false;
    }
}

bool callGlClearBufferiv(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    int32_t drawbuffer = stack->pop<int32_t>();
    GLenum buffer = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glClearBufferiv(%u, %" PRId32 ", %p)", buffer, drawbuffer, value);
        if (glClearBufferiv != nullptr) {
            glClearBufferiv(buffer, drawbuffer, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glClearBufferiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glClearBufferiv");
        return false;
    }
}

bool callGlClearBufferuiv(Stack* stack, bool pushReturn) {
    uint32_t* value = stack->pop<uint32_t*>();
    int32_t drawbuffer = stack->pop<int32_t>();
    GLenum buffer = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glClearBufferuiv(%u, %" PRId32 ", %p)", buffer, drawbuffer, value);
        if (glClearBufferuiv != nullptr) {
            glClearBufferuiv(buffer, drawbuffer, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glClearBufferuiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glClearBufferuiv");
        return false;
    }
}

bool callGlClearColor(Stack* stack, bool pushReturn) {
    float a = stack->pop<float>();
    float b = stack->pop<float>();
    float g = stack->pop<float>();
    float r = stack->pop<float>();
    if (stack->isValid()) {
        GAPID_INFO("glClearColor(%f, %f, %f, %f)", r, g, b, a);
        if (glClearColor != nullptr) {
            glClearColor(r, g, b, a);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glClearColor");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glClearColor");
        return false;
    }
}

bool callGlClearDepthf(Stack* stack, bool pushReturn) {
    float depth = stack->pop<float>();
    if (stack->isValid()) {
        GAPID_INFO("glClearDepthf(%f)", depth);
        if (glClearDepthf != nullptr) {
            glClearDepthf(depth);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glClearDepthf");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glClearDepthf");
        return false;
    }
}

bool callGlClearStencil(Stack* stack, bool pushReturn) {
    int32_t stencil = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glClearStencil(%" PRId32 ")", stencil);
        if (glClearStencil != nullptr) {
            glClearStencil(stencil);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glClearStencil");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glClearStencil");
        return false;
    }
}

bool callGlColorMask(Stack* stack, bool pushReturn) {
    uint8_t alpha = stack->pop<uint8_t>();
    uint8_t blue = stack->pop<uint8_t>();
    uint8_t green = stack->pop<uint8_t>();
    uint8_t red = stack->pop<uint8_t>();
    if (stack->isValid()) {
        GAPID_INFO("glColorMask(%" PRIu8 ", %" PRIu8 ", %" PRIu8 ", %" PRIu8 ")", red, green, blue,
                   alpha);
        if (glColorMask != nullptr) {
            glColorMask(red, green, blue, alpha);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glColorMask");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glColorMask");
        return false;
    }
}

bool callGlColorMaski(Stack* stack, bool pushReturn) {
    uint8_t a = stack->pop<uint8_t>();
    uint8_t b = stack->pop<uint8_t>();
    uint8_t g = stack->pop<uint8_t>();
    uint8_t r = stack->pop<uint8_t>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glColorMaski(%" PRIu32 ", %" PRIu8 ", %" PRIu8 ", %" PRIu8 ", %" PRIu8 ")",
                   index, r, g, b, a);
        if (glColorMaski != nullptr) {
            glColorMaski(index, r, g, b, a);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glColorMaski");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glColorMaski");
        return false;
    }
}

bool callGlDeleteFramebuffers(Stack* stack, bool pushReturn) {
    uint32_t* framebuffers = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteFramebuffers(%" PRId32 ", %p)", count, framebuffers);
        if (glDeleteFramebuffers != nullptr) {
            glDeleteFramebuffers(count, framebuffers);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteFramebuffers");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteFramebuffers");
        return false;
    }
}

bool callGlDeleteRenderbuffers(Stack* stack, bool pushReturn) {
    uint32_t* renderbuffers = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteRenderbuffers(%" PRId32 ", %p)", count, renderbuffers);
        if (glDeleteRenderbuffers != nullptr) {
            glDeleteRenderbuffers(count, renderbuffers);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteRenderbuffers");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteRenderbuffers");
        return false;
    }
}

bool callGlDepthMask(Stack* stack, bool pushReturn) {
    uint8_t enabled = stack->pop<uint8_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDepthMask(%" PRIu8 ")", enabled);
        if (glDepthMask != nullptr) {
            glDepthMask(enabled);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDepthMask");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDepthMask");
        return false;
    }
}

bool callGlDrawBuffers(Stack* stack, bool pushReturn) {
    GLenum* bufs = stack->pop<GLenum*>();
    int32_t n = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawBuffers(%" PRId32 ", %p)", n, bufs);
        if (glDrawBuffers != nullptr) {
            glDrawBuffers(n, bufs);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawBuffers");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawBuffers");
        return false;
    }
}

bool callGlFramebufferParameteri(Stack* stack, bool pushReturn) {
    int32_t param = stack->pop<int32_t>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glFramebufferParameteri(%u, %u, %" PRId32 ")", target, pname, param);
        if (glFramebufferParameteri != nullptr) {
            glFramebufferParameteri(target, pname, param);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFramebufferParameteri");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFramebufferParameteri");
        return false;
    }
}

bool callGlFramebufferRenderbuffer(Stack* stack, bool pushReturn) {
    uint32_t renderbuffer = stack->pop<uint32_t>();
    GLenum renderbuffer_target = stack->pop<GLenum>();
    GLenum framebuffer_attachment = stack->pop<GLenum>();
    GLenum framebuffer_target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glFramebufferRenderbuffer(%u, %u, %u, %" PRIu32 ")", framebuffer_target,
                   framebuffer_attachment, renderbuffer_target, renderbuffer);
        if (glFramebufferRenderbuffer != nullptr) {
            glFramebufferRenderbuffer(framebuffer_target, framebuffer_attachment,
                                      renderbuffer_target, renderbuffer);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFramebufferRenderbuffer");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFramebufferRenderbuffer");
        return false;
    }
}

bool callGlFramebufferTexture(Stack* stack, bool pushReturn) {
    int32_t level = stack->pop<int32_t>();
    uint32_t texture = stack->pop<uint32_t>();
    GLenum attachment = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glFramebufferTexture(%u, %u, %" PRIu32 ", %" PRId32 ")", target, attachment,
                   texture, level);
        if (glFramebufferTexture != nullptr) {
            glFramebufferTexture(target, attachment, texture, level);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFramebufferTexture");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFramebufferTexture");
        return false;
    }
}

bool callGlFramebufferTexture2D(Stack* stack, bool pushReturn) {
    int32_t level = stack->pop<int32_t>();
    uint32_t texture = stack->pop<uint32_t>();
    GLenum texture_target = stack->pop<GLenum>();
    GLenum framebuffer_attachment = stack->pop<GLenum>();
    GLenum framebuffer_target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glFramebufferTexture2D(%u, %u, %u, %" PRIu32 ", %" PRId32 ")",
                   framebuffer_target, framebuffer_attachment, texture_target, texture, level);
        if (glFramebufferTexture2D != nullptr) {
            glFramebufferTexture2D(framebuffer_target, framebuffer_attachment, texture_target,
                                   texture, level);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFramebufferTexture2D");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFramebufferTexture2D");
        return false;
    }
}

bool callGlFramebufferTextureLayer(Stack* stack, bool pushReturn) {
    int32_t layer = stack->pop<int32_t>();
    int32_t level = stack->pop<int32_t>();
    uint32_t texture = stack->pop<uint32_t>();
    GLenum attachment = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glFramebufferTextureLayer(%u, %u, %" PRIu32 ", %" PRId32 ", %" PRId32 ")",
                   target, attachment, texture, level, layer);
        if (glFramebufferTextureLayer != nullptr) {
            glFramebufferTextureLayer(target, attachment, texture, level, layer);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFramebufferTextureLayer");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFramebufferTextureLayer");
        return false;
    }
}

bool callGlGenFramebuffers(Stack* stack, bool pushReturn) {
    uint32_t* framebuffers = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenFramebuffers(%" PRId32 ", %p)", count, framebuffers);
        if (glGenFramebuffers != nullptr) {
            glGenFramebuffers(count, framebuffers);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenFramebuffers");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenFramebuffers");
        return false;
    }
}

bool callGlGenRenderbuffers(Stack* stack, bool pushReturn) {
    uint32_t* renderbuffers = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenRenderbuffers(%" PRId32 ", %p)", count, renderbuffers);
        if (glGenRenderbuffers != nullptr) {
            glGenRenderbuffers(count, renderbuffers);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenRenderbuffers");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenRenderbuffers");
        return false;
    }
}

bool callGlGetFramebufferAttachmentParameteriv(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    GLenum parameter = stack->pop<GLenum>();
    GLenum attachment = stack->pop<GLenum>();
    GLenum framebuffer_target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetFramebufferAttachmentParameteriv(%u, %u, %u, %p)", framebuffer_target,
                   attachment, parameter, value);
        if (glGetFramebufferAttachmentParameteriv != nullptr) {
            glGetFramebufferAttachmentParameteriv(framebuffer_target, attachment, parameter, value);
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glGetFramebufferAttachmentParameteriv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetFramebufferAttachmentParameteriv");
        return false;
    }
}

bool callGlGetFramebufferParameteriv(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetFramebufferParameteriv(%u, %u, %p)", target, pname, params);
        if (glGetFramebufferParameteriv != nullptr) {
            glGetFramebufferParameteriv(target, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetFramebufferParameteriv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetFramebufferParameteriv");
        return false;
    }
}

bool callGlGetRenderbufferParameteriv(Stack* stack, bool pushReturn) {
    int32_t* values = stack->pop<int32_t*>();
    GLenum parameter = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetRenderbufferParameteriv(%u, %u, %p)", target, parameter, values);
        if (glGetRenderbufferParameteriv != nullptr) {
            glGetRenderbufferParameteriv(target, parameter, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetRenderbufferParameteriv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetRenderbufferParameteriv");
        return false;
    }
}

bool callGlInvalidateFramebuffer(Stack* stack, bool pushReturn) {
    GLenum* attachments = stack->pop<GLenum*>();
    int32_t count = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glInvalidateFramebuffer(%u, %" PRId32 ", %p)", target, count, attachments);
        if (glInvalidateFramebuffer != nullptr) {
            glInvalidateFramebuffer(target, count, attachments);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glInvalidateFramebuffer");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glInvalidateFramebuffer");
        return false;
    }
}

bool callGlInvalidateSubFramebuffer(Stack* stack, bool pushReturn) {
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    int32_t y = stack->pop<int32_t>();
    int32_t x = stack->pop<int32_t>();
    GLenum* attachments = stack->pop<GLenum*>();
    int32_t numAttachments = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glInvalidateSubFramebuffer(%u, %" PRId32 ", %p, %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %" PRId32 ")",
                   target, numAttachments, attachments, x, y, width, height);
        if (glInvalidateSubFramebuffer != nullptr) {
            glInvalidateSubFramebuffer(target, numAttachments, attachments, x, y, width, height);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glInvalidateSubFramebuffer");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glInvalidateSubFramebuffer");
        return false;
    }
}

bool callGlIsFramebuffer(Stack* stack, bool pushReturn) {
    uint32_t framebuffer = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsFramebuffer(%" PRIu32 ")", framebuffer);
        if (glIsFramebuffer != nullptr) {
            uint8_t return_value = glIsFramebuffer(framebuffer);
            GAPID_INFO("Returned: %" PRIu8 "", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsFramebuffer");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsFramebuffer");
        return false;
    }
}

bool callGlIsRenderbuffer(Stack* stack, bool pushReturn) {
    uint32_t renderbuffer = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsRenderbuffer(%" PRIu32 ")", renderbuffer);
        if (glIsRenderbuffer != nullptr) {
            uint8_t return_value = glIsRenderbuffer(renderbuffer);
            GAPID_INFO("Returned: %" PRIu8 "", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsRenderbuffer");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsRenderbuffer");
        return false;
    }
}

bool callGlReadBuffer(Stack* stack, bool pushReturn) {
    GLenum src = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glReadBuffer(%u)", src);
        if (glReadBuffer != nullptr) {
            glReadBuffer(src);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glReadBuffer");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glReadBuffer");
        return false;
    }
}

bool callGlReadPixels(Stack* stack, bool pushReturn) {
    void* data = stack->pop<void*>();
    GLenum type = stack->pop<GLenum>();
    GLenum format = stack->pop<GLenum>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    int32_t y = stack->pop<int32_t>();
    int32_t x = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glReadPixels(%" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %u, %u, %p)",
                   x, y, width, height, format, type, data);
        if (glReadPixels != nullptr) {
            glReadPixels(x, y, width, height, format, type, data);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glReadPixels");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glReadPixels");
        return false;
    }
}

bool callGlReadnPixels(Stack* stack, bool pushReturn) {
    void* data = stack->pop<void*>();
    int32_t bufSize = stack->pop<int32_t>();
    GLenum type = stack->pop<GLenum>();
    GLenum format = stack->pop<GLenum>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    int32_t y = stack->pop<int32_t>();
    int32_t x = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glReadnPixels(%" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %u, %u, %" PRId32 ", %p)",
                   x, y, width, height, format, type, bufSize, data);
        if (glReadnPixels != nullptr) {
            glReadnPixels(x, y, width, height, format, type, bufSize, data);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glReadnPixels");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glReadnPixels");
        return false;
    }
}

bool callGlRenderbufferStorage(Stack* stack, bool pushReturn) {
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    GLenum format = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glRenderbufferStorage(%u, %u, %" PRId32 ", %" PRId32 ")", target, format, width,
                   height);
        if (glRenderbufferStorage != nullptr) {
            glRenderbufferStorage(target, format, width, height);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glRenderbufferStorage");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glRenderbufferStorage");
        return false;
    }
}

bool callGlRenderbufferStorageMultisample(Stack* stack, bool pushReturn) {
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    GLenum format = stack->pop<GLenum>();
    int32_t samples = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glRenderbufferStorageMultisample(%u, %" PRId32 ", %u, %" PRId32 ", %" PRId32
                   ")",
                   target, samples, format, width, height);
        if (glRenderbufferStorageMultisample != nullptr) {
            glRenderbufferStorageMultisample(target, samples, format, width, height);
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glRenderbufferStorageMultisample");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glRenderbufferStorageMultisample");
        return false;
    }
}

bool callGlStencilMask(Stack* stack, bool pushReturn) {
    uint32_t mask = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glStencilMask(%" PRIu32 ")", mask);
        if (glStencilMask != nullptr) {
            glStencilMask(mask);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glStencilMask");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStencilMask");
        return false;
    }
}

bool callGlStencilMaskSeparate(Stack* stack, bool pushReturn) {
    uint32_t mask = stack->pop<uint32_t>();
    GLenum face = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glStencilMaskSeparate(%u, %" PRIu32 ")", face, mask);
        if (glStencilMaskSeparate != nullptr) {
            glStencilMaskSeparate(face, mask);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glStencilMaskSeparate");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStencilMaskSeparate");
        return false;
    }
}

bool callGlDisable(Stack* stack, bool pushReturn) {
    GLenum capability = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDisable(%u)", capability);
        if (glDisable != nullptr) {
            glDisable(capability);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDisable");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDisable");
        return false;
    }
}

bool callGlDisablei(Stack* stack, bool pushReturn) {
    uint32_t index = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDisablei(%u, %" PRIu32 ")", target, index);
        if (glDisablei != nullptr) {
            glDisablei(target, index);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDisablei");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDisablei");
        return false;
    }
}

bool callGlEnable(Stack* stack, bool pushReturn) {
    GLenum capability = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glEnable(%u)", capability);
        if (glEnable != nullptr) {
            glEnable(capability);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEnable");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEnable");
        return false;
    }
}

bool callGlEnablei(Stack* stack, bool pushReturn) {
    uint32_t index = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glEnablei(%u, %" PRIu32 ")", target, index);
        if (glEnablei != nullptr) {
            glEnablei(target, index);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEnablei");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEnablei");
        return false;
    }
}

bool callGlFinish(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glFinish()");
        if (glFinish != nullptr) {
            glFinish();
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFinish");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFinish");
        return false;
    }
}

bool callGlFlush(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glFlush()");
        if (glFlush != nullptr) {
            glFlush();
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFlush");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFlush");
        return false;
    }
}

bool callGlFlushMappedBufferRange(Stack* stack, bool pushReturn) {
    int32_t length = stack->pop<int32_t>();
    int32_t offset = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glFlushMappedBufferRange(%u, %" PRId32 ", %" PRId32 ")", target, offset,
                   length);
        if (glFlushMappedBufferRange != nullptr) {
            glFlushMappedBufferRange(target, offset, length);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFlushMappedBufferRange");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFlushMappedBufferRange");
        return false;
    }
}

bool callGlGetError(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glGetError()");
        if (glGetError != nullptr) {
            GLenum return_value = glGetError();
            GAPID_INFO("Returned: %u", return_value);
            if (pushReturn) {
                stack->push<GLenum>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetError");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetError");
        return false;
    }
}

bool callGlGetGraphicsResetStatus(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glGetGraphicsResetStatus()");
        if (glGetGraphicsResetStatus != nullptr) {
            GLenum return_value = glGetGraphicsResetStatus();
            GAPID_INFO("Returned: %u", return_value);
            if (pushReturn) {
                stack->push<GLenum>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetGraphicsResetStatus");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetGraphicsResetStatus");
        return false;
    }
}

bool callGlHint(Stack* stack, bool pushReturn) {
    GLenum mode = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glHint(%u, %u)", target, mode);
        if (glHint != nullptr) {
            glHint(target, mode);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glHint");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glHint");
        return false;
    }
}

bool callGlActiveShaderProgram(Stack* stack, bool pushReturn) {
    uint32_t program = stack->pop<uint32_t>();
    uint32_t pipeline = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glActiveShaderProgram(%" PRIu32 ", %" PRIu32 ")", pipeline, program);
        if (glActiveShaderProgram != nullptr) {
            glActiveShaderProgram(pipeline, program);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glActiveShaderProgram");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glActiveShaderProgram");
        return false;
    }
}

bool callGlAttachShader(Stack* stack, bool pushReturn) {
    uint32_t shader = stack->pop<uint32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glAttachShader(%" PRIu32 ", %" PRIu32 ")", program, shader);
        if (glAttachShader != nullptr) {
            glAttachShader(program, shader);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glAttachShader");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glAttachShader");
        return false;
    }
}

bool callGlBindAttribLocation(Stack* stack, bool pushReturn) {
    char* name = stack->pop<char*>();
    uint32_t location = stack->pop<uint32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBindAttribLocation(%" PRIu32 ", %" PRIu32 ", %s)", program, location, name);
        if (glBindAttribLocation != nullptr) {
            glBindAttribLocation(program, location, name);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindAttribLocation");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindAttribLocation");
        return false;
    }
}

bool callGlBindProgramPipeline(Stack* stack, bool pushReturn) {
    uint32_t pipeline = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBindProgramPipeline(%" PRIu32 ")", pipeline);
        if (glBindProgramPipeline != nullptr) {
            glBindProgramPipeline(pipeline);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindProgramPipeline");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindProgramPipeline");
        return false;
    }
}

bool callGlCompileShader(Stack* stack, bool pushReturn) {
    uint32_t shader = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glCompileShader(%" PRIu32 ")", shader);
        if (glCompileShader != nullptr) {
            glCompileShader(shader);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCompileShader");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCompileShader");
        return false;
    }
}

bool callGlCreateProgram(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glCreateProgram()");
        if (glCreateProgram != nullptr) {
            uint32_t return_value = glCreateProgram();
            GAPID_INFO("Returned: %" PRIu32 "", return_value);
            if (pushReturn) {
                stack->push<uint32_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCreateProgram");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCreateProgram");
        return false;
    }
}

bool callGlCreateShader(Stack* stack, bool pushReturn) {
    GLenum type = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glCreateShader(%u)", type);
        if (glCreateShader != nullptr) {
            uint32_t return_value = glCreateShader(type);
            GAPID_INFO("Returned: %" PRIu32 "", return_value);
            if (pushReturn) {
                stack->push<uint32_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCreateShader");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCreateShader");
        return false;
    }
}

bool callGlCreateShaderProgramv(Stack* stack, bool pushReturn) {
    char** strings = stack->pop<char**>();
    int32_t count = stack->pop<int32_t>();
    GLenum type = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glCreateShaderProgramv(%u, %" PRId32 ", %p)", type, count, strings);
        if (glCreateShaderProgramv != nullptr) {
            uint32_t return_value = glCreateShaderProgramv(type, count, strings);
            GAPID_INFO("Returned: %" PRIu32 "", return_value);
            if (pushReturn) {
                stack->push<uint32_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCreateShaderProgramv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCreateShaderProgramv");
        return false;
    }
}

bool callGlDeleteProgram(Stack* stack, bool pushReturn) {
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteProgram(%" PRIu32 ")", program);
        if (glDeleteProgram != nullptr) {
            glDeleteProgram(program);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteProgram");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteProgram");
        return false;
    }
}

bool callGlDeleteProgramPipelines(Stack* stack, bool pushReturn) {
    uint32_t* pipelines = stack->pop<uint32_t*>();
    int32_t n = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteProgramPipelines(%" PRId32 ", %p)", n, pipelines);
        if (glDeleteProgramPipelines != nullptr) {
            glDeleteProgramPipelines(n, pipelines);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteProgramPipelines");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteProgramPipelines");
        return false;
    }
}

bool callGlDeleteShader(Stack* stack, bool pushReturn) {
    uint32_t shader = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteShader(%" PRIu32 ")", shader);
        if (glDeleteShader != nullptr) {
            glDeleteShader(shader);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteShader");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteShader");
        return false;
    }
}

bool callGlDetachShader(Stack* stack, bool pushReturn) {
    uint32_t shader = stack->pop<uint32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDetachShader(%" PRIu32 ", %" PRIu32 ")", program, shader);
        if (glDetachShader != nullptr) {
            glDetachShader(program, shader);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDetachShader");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDetachShader");
        return false;
    }
}

bool callGlDispatchCompute(Stack* stack, bool pushReturn) {
    uint32_t num_groups_z = stack->pop<uint32_t>();
    uint32_t num_groups_y = stack->pop<uint32_t>();
    uint32_t num_groups_x = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDispatchCompute(%" PRIu32 ", %" PRIu32 ", %" PRIu32 ")", num_groups_x,
                   num_groups_y, num_groups_z);
        if (glDispatchCompute != nullptr) {
            glDispatchCompute(num_groups_x, num_groups_y, num_groups_z);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDispatchCompute");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDispatchCompute");
        return false;
    }
}

bool callGlDispatchComputeIndirect(Stack* stack, bool pushReturn) {
    int32_t indirect = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDispatchComputeIndirect(%" PRId32 ")", indirect);
        if (glDispatchComputeIndirect != nullptr) {
            glDispatchComputeIndirect(indirect);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDispatchComputeIndirect");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDispatchComputeIndirect");
        return false;
    }
}

bool callGlGenProgramPipelines(Stack* stack, bool pushReturn) {
    uint32_t* pipelines = stack->pop<uint32_t*>();
    int32_t n = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenProgramPipelines(%" PRId32 ", %p)", n, pipelines);
        if (glGenProgramPipelines != nullptr) {
            glGenProgramPipelines(n, pipelines);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenProgramPipelines");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenProgramPipelines");
        return false;
    }
}

bool callGlGetActiveAttrib(Stack* stack, bool pushReturn) {
    char* name = stack->pop<char*>();
    GLenum* type = stack->pop<GLenum*>();
    int32_t* vector_count = stack->pop<int32_t*>();
    int32_t* buffer_bytes_written = stack->pop<int32_t*>();
    int32_t buffer_size = stack->pop<int32_t>();
    uint32_t location = stack->pop<uint32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetActiveAttrib(%" PRIu32 ", %" PRIu32 ", %" PRId32 ", %p, %p, %p, %p)",
                   program, location, buffer_size, buffer_bytes_written, vector_count, type, name);
        if (glGetActiveAttrib != nullptr) {
            glGetActiveAttrib(program, location, buffer_size, buffer_bytes_written, vector_count,
                              type, name);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetActiveAttrib");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetActiveAttrib");
        return false;
    }
}

bool callGlGetActiveUniform(Stack* stack, bool pushReturn) {
    char* name = stack->pop<char*>();
    GLenum* type = stack->pop<GLenum*>();
    int32_t* vector_count = stack->pop<int32_t*>();
    int32_t* buffer_bytes_written = stack->pop<int32_t*>();
    int32_t buffer_size = stack->pop<int32_t>();
    uint32_t location = stack->pop<uint32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetActiveUniform(%" PRIu32 ", %" PRIu32 ", %" PRId32 ", %p, %p, %p, %p)",
                   program, location, buffer_size, buffer_bytes_written, vector_count, type, name);
        if (glGetActiveUniform != nullptr) {
            glGetActiveUniform(program, location, buffer_size, buffer_bytes_written, vector_count,
                               type, name);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetActiveUniform");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetActiveUniform");
        return false;
    }
}

bool callGlGetActiveUniformBlockName(Stack* stack, bool pushReturn) {
    char* name = stack->pop<char*>();
    int32_t* buffer_bytes_written = stack->pop<int32_t*>();
    int32_t buffer_size = stack->pop<int32_t>();
    uint32_t uniform_block_index = stack->pop<uint32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetActiveUniformBlockName(%" PRIu32 ", %" PRIu32 ", %" PRId32 ", %p, %p)",
                   program, uniform_block_index, buffer_size, buffer_bytes_written, name);
        if (glGetActiveUniformBlockName != nullptr) {
            glGetActiveUniformBlockName(program, uniform_block_index, buffer_size,
                                        buffer_bytes_written, name);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetActiveUniformBlockName");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetActiveUniformBlockName");
        return false;
    }
}

bool callGlGetActiveUniformBlockiv(Stack* stack, bool pushReturn) {
    int32_t* parameters = stack->pop<int32_t*>();
    GLenum parameter_name = stack->pop<GLenum>();
    uint32_t uniform_block_index = stack->pop<uint32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetActiveUniformBlockiv(%" PRIu32 ", %" PRIu32 ", %u, %p)", program,
                   uniform_block_index, parameter_name, parameters);
        if (glGetActiveUniformBlockiv != nullptr) {
            glGetActiveUniformBlockiv(program, uniform_block_index, parameter_name, parameters);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetActiveUniformBlockiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetActiveUniformBlockiv");
        return false;
    }
}

bool callGlGetActiveUniformsiv(Stack* stack, bool pushReturn) {
    int32_t* parameters = stack->pop<int32_t*>();
    GLenum parameter_name = stack->pop<GLenum>();
    uint32_t* uniform_indices = stack->pop<uint32_t*>();
    int32_t uniform_count = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetActiveUniformsiv(%" PRIu32 ", %" PRId32 ", %p, %u, %p)", program,
                   uniform_count, uniform_indices, parameter_name, parameters);
        if (glGetActiveUniformsiv != nullptr) {
            glGetActiveUniformsiv(program, uniform_count, uniform_indices, parameter_name,
                                  parameters);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetActiveUniformsiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetActiveUniformsiv");
        return false;
    }
}

bool callGlGetAttachedShaders(Stack* stack, bool pushReturn) {
    uint32_t* shaders = stack->pop<uint32_t*>();
    int32_t* shaders_length_written = stack->pop<int32_t*>();
    int32_t buffer_length = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetAttachedShaders(%" PRIu32 ", %" PRId32 ", %p, %p)", program, buffer_length,
                   shaders_length_written, shaders);
        if (glGetAttachedShaders != nullptr) {
            glGetAttachedShaders(program, buffer_length, shaders_length_written, shaders);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetAttachedShaders");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetAttachedShaders");
        return false;
    }
}

bool callGlGetAttribLocation(Stack* stack, bool pushReturn) {
    char* name = stack->pop<char*>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetAttribLocation(%" PRIu32 ", %s)", program, name);
        if (glGetAttribLocation != nullptr) {
            int32_t return_value = glGetAttribLocation(program, name);
            GAPID_INFO("Returned: %" PRId32 "", return_value);
            if (pushReturn) {
                stack->push<int32_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetAttribLocation");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetAttribLocation");
        return false;
    }
}

bool callGlGetFragDataLocation(Stack* stack, bool pushReturn) {
    char* name = stack->pop<char*>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetFragDataLocation(%" PRIu32 ", %s)", program, name);
        if (glGetFragDataLocation != nullptr) {
            int32_t return_value = glGetFragDataLocation(program, name);
            GAPID_INFO("Returned: %" PRId32 "", return_value);
            if (pushReturn) {
                stack->push<int32_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetFragDataLocation");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetFragDataLocation");
        return false;
    }
}

bool callGlGetProgramBinary(Stack* stack, bool pushReturn) {
    void* binary = stack->pop<void*>();
    GLenum* binaryFormat = stack->pop<GLenum*>();
    int32_t* length = stack->pop<int32_t*>();
    int32_t bufSize = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetProgramBinary(%" PRIu32 ", %" PRId32 ", %p, %p, %p)", program, bufSize,
                   length, binaryFormat, binary);
        if (glGetProgramBinary != nullptr) {
            glGetProgramBinary(program, bufSize, length, binaryFormat, binary);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetProgramBinary");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetProgramBinary");
        return false;
    }
}

bool callGlGetProgramInfoLog(Stack* stack, bool pushReturn) {
    char* info = stack->pop<char*>();
    int32_t* string_length_written = stack->pop<int32_t*>();
    int32_t buffer_length = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetProgramInfoLog(%" PRIu32 ", %" PRId32 ", %p, %p)", program, buffer_length,
                   string_length_written, info);
        if (glGetProgramInfoLog != nullptr) {
            glGetProgramInfoLog(program, buffer_length, string_length_written, info);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetProgramInfoLog");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetProgramInfoLog");
        return false;
    }
}

bool callGlGetProgramInterfaceiv(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    GLenum programInterface = stack->pop<GLenum>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetProgramInterfaceiv(%" PRIu32 ", %u, %u, %p)", program, programInterface,
                   pname, params);
        if (glGetProgramInterfaceiv != nullptr) {
            glGetProgramInterfaceiv(program, programInterface, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetProgramInterfaceiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetProgramInterfaceiv");
        return false;
    }
}

bool callGlGetProgramPipelineInfoLog(Stack* stack, bool pushReturn) {
    char* infoLog = stack->pop<char*>();
    int32_t* length = stack->pop<int32_t*>();
    int32_t bufSize = stack->pop<int32_t>();
    uint32_t pipeline = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetProgramPipelineInfoLog(%" PRIu32 ", %" PRId32 ", %p, %p)", pipeline,
                   bufSize, length, infoLog);
        if (glGetProgramPipelineInfoLog != nullptr) {
            glGetProgramPipelineInfoLog(pipeline, bufSize, length, infoLog);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetProgramPipelineInfoLog");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetProgramPipelineInfoLog");
        return false;
    }
}

bool callGlGetProgramPipelineiv(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t pipeline = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetProgramPipelineiv(%" PRIu32 ", %u, %p)", pipeline, pname, params);
        if (glGetProgramPipelineiv != nullptr) {
            glGetProgramPipelineiv(pipeline, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetProgramPipelineiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetProgramPipelineiv");
        return false;
    }
}

bool callGlGetProgramResourceIndex(Stack* stack, bool pushReturn) {
    char* name = stack->pop<char*>();
    GLenum programInterface = stack->pop<GLenum>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetProgramResourceIndex(%" PRIu32 ", %u, %s)", program, programInterface,
                   name);
        if (glGetProgramResourceIndex != nullptr) {
            uint32_t return_value = glGetProgramResourceIndex(program, programInterface, name);
            GAPID_INFO("Returned: %" PRIu32 "", return_value);
            if (pushReturn) {
                stack->push<uint32_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetProgramResourceIndex");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetProgramResourceIndex");
        return false;
    }
}

bool callGlGetProgramResourceLocation(Stack* stack, bool pushReturn) {
    char* name = stack->pop<char*>();
    GLenum programInterface = stack->pop<GLenum>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetProgramResourceLocation(%" PRIu32 ", %u, %s)", program, programInterface,
                   name);
        if (glGetProgramResourceLocation != nullptr) {
            int32_t return_value = glGetProgramResourceLocation(program, programInterface, name);
            GAPID_INFO("Returned: %" PRId32 "", return_value);
            if (pushReturn) {
                stack->push<int32_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetProgramResourceLocation");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetProgramResourceLocation");
        return false;
    }
}

bool callGlGetProgramResourceName(Stack* stack, bool pushReturn) {
    char* name = stack->pop<char*>();
    int32_t* length = stack->pop<int32_t*>();
    int32_t bufSize = stack->pop<int32_t>();
    uint32_t index = stack->pop<uint32_t>();
    GLenum programInterface = stack->pop<GLenum>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetProgramResourceName(%" PRIu32 ", %u, %" PRIu32 ", %" PRId32 ", %p, %p)",
                   program, programInterface, index, bufSize, length, name);
        if (glGetProgramResourceName != nullptr) {
            glGetProgramResourceName(program, programInterface, index, bufSize, length, name);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetProgramResourceName");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetProgramResourceName");
        return false;
    }
}

bool callGlGetProgramResourceiv(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    int32_t* length = stack->pop<int32_t*>();
    int32_t bufSize = stack->pop<int32_t>();
    GLenum* props = stack->pop<GLenum*>();
    int32_t propCount = stack->pop<int32_t>();
    uint32_t index = stack->pop<uint32_t>();
    GLenum programInterface = stack->pop<GLenum>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetProgramResourceiv(%" PRIu32 ", %u, %" PRIu32 ", %" PRId32 ", %p, %" PRId32
                   ", %p, %p)",
                   program, programInterface, index, propCount, props, bufSize, length, params);
        if (glGetProgramResourceiv != nullptr) {
            glGetProgramResourceiv(program, programInterface, index, propCount, props, bufSize,
                                   length, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetProgramResourceiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetProgramResourceiv");
        return false;
    }
}

bool callGlGetProgramiv(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    GLenum parameter = stack->pop<GLenum>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetProgramiv(%" PRIu32 ", %u, %p)", program, parameter, value);
        if (glGetProgramiv != nullptr) {
            glGetProgramiv(program, parameter, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetProgramiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetProgramiv");
        return false;
    }
}

bool callGlGetShaderInfoLog(Stack* stack, bool pushReturn) {
    char* info = stack->pop<char*>();
    int32_t* string_length_written = stack->pop<int32_t*>();
    int32_t buffer_length = stack->pop<int32_t>();
    uint32_t shader = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetShaderInfoLog(%" PRIu32 ", %" PRId32 ", %p, %p)", shader, buffer_length,
                   string_length_written, info);
        if (glGetShaderInfoLog != nullptr) {
            glGetShaderInfoLog(shader, buffer_length, string_length_written, info);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetShaderInfoLog");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetShaderInfoLog");
        return false;
    }
}

bool callGlGetShaderPrecisionFormat(Stack* stack, bool pushReturn) {
    int32_t* precision = stack->pop<int32_t*>();
    int32_t* range = stack->pop<int32_t*>();
    GLenum precision_type = stack->pop<GLenum>();
    GLenum shader_type = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetShaderPrecisionFormat(%u, %u, %p, %p)", shader_type, precision_type, range,
                   precision);
        if (glGetShaderPrecisionFormat != nullptr) {
            glGetShaderPrecisionFormat(shader_type, precision_type, range, precision);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetShaderPrecisionFormat");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetShaderPrecisionFormat");
        return false;
    }
}

bool callGlGetShaderSource(Stack* stack, bool pushReturn) {
    char* source = stack->pop<char*>();
    int32_t* string_length_written = stack->pop<int32_t*>();
    int32_t buffer_length = stack->pop<int32_t>();
    uint32_t shader = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetShaderSource(%" PRIu32 ", %" PRId32 ", %p, %p)", shader, buffer_length,
                   string_length_written, source);
        if (glGetShaderSource != nullptr) {
            glGetShaderSource(shader, buffer_length, string_length_written, source);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetShaderSource");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetShaderSource");
        return false;
    }
}

bool callGlGetShaderiv(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    GLenum parameter = stack->pop<GLenum>();
    uint32_t shader = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetShaderiv(%" PRIu32 ", %u, %p)", shader, parameter, value);
        if (glGetShaderiv != nullptr) {
            glGetShaderiv(shader, parameter, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetShaderiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetShaderiv");
        return false;
    }
}

bool callGlGetUniformBlockIndex(Stack* stack, bool pushReturn) {
    char* uniformBlockName = stack->pop<char*>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetUniformBlockIndex(%" PRIu32 ", %s)", program, uniformBlockName);
        if (glGetUniformBlockIndex != nullptr) {
            uint32_t return_value = glGetUniformBlockIndex(program, uniformBlockName);
            GAPID_INFO("Returned: %" PRIu32 "", return_value);
            if (pushReturn) {
                stack->push<uint32_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetUniformBlockIndex");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetUniformBlockIndex");
        return false;
    }
}

bool callGlGetUniformIndices(Stack* stack, bool pushReturn) {
    uint32_t* uniformIndices = stack->pop<uint32_t*>();
    char** uniformNames = stack->pop<char**>();
    int32_t uniformCount = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetUniformIndices(%" PRIu32 ", %" PRId32 ", %p, %p)", program, uniformCount,
                   uniformNames, uniformIndices);
        if (glGetUniformIndices != nullptr) {
            glGetUniformIndices(program, uniformCount, uniformNames, uniformIndices);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetUniformIndices");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetUniformIndices");
        return false;
    }
}

bool callGlGetUniformLocation(Stack* stack, bool pushReturn) {
    char* name = stack->pop<char*>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetUniformLocation(%" PRIu32 ", %s)", program, name);
        if (glGetUniformLocation != nullptr) {
            int32_t return_value = glGetUniformLocation(program, name);
            GAPID_INFO("Returned: %" PRId32 "", return_value);
            if (pushReturn) {
                stack->push<int32_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetUniformLocation");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetUniformLocation");
        return false;
    }
}

bool callGlGetUniformfv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetUniformfv(%" PRIu32 ", %" PRId32 ", %p)", program, location, values);
        if (glGetUniformfv != nullptr) {
            glGetUniformfv(program, location, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetUniformfv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetUniformfv");
        return false;
    }
}

bool callGlGetUniformiv(Stack* stack, bool pushReturn) {
    int32_t* values = stack->pop<int32_t*>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetUniformiv(%" PRIu32 ", %" PRId32 ", %p)", program, location, values);
        if (glGetUniformiv != nullptr) {
            glGetUniformiv(program, location, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetUniformiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetUniformiv");
        return false;
    }
}

bool callGlGetUniformuiv(Stack* stack, bool pushReturn) {
    uint32_t* values = stack->pop<uint32_t*>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetUniformuiv(%" PRIu32 ", %" PRId32 ", %p)", program, location, values);
        if (glGetUniformuiv != nullptr) {
            glGetUniformuiv(program, location, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetUniformuiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetUniformuiv");
        return false;
    }
}

bool callGlGetnUniformfv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    int32_t bufSize = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetnUniformfv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)", program, location,
                   bufSize, values);
        if (glGetnUniformfv != nullptr) {
            glGetnUniformfv(program, location, bufSize, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetnUniformfv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetnUniformfv");
        return false;
    }
}

bool callGlGetnUniformiv(Stack* stack, bool pushReturn) {
    int32_t* values = stack->pop<int32_t*>();
    int32_t bufSize = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetnUniformiv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)", program, location,
                   bufSize, values);
        if (glGetnUniformiv != nullptr) {
            glGetnUniformiv(program, location, bufSize, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetnUniformiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetnUniformiv");
        return false;
    }
}

bool callGlGetnUniformuiv(Stack* stack, bool pushReturn) {
    uint32_t* values = stack->pop<uint32_t*>();
    int32_t bufSize = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetnUniformuiv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)", program, location,
                   bufSize, values);
        if (glGetnUniformuiv != nullptr) {
            glGetnUniformuiv(program, location, bufSize, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetnUniformuiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetnUniformuiv");
        return false;
    }
}

bool callGlIsProgram(Stack* stack, bool pushReturn) {
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsProgram(%" PRIu32 ")", program);
        if (glIsProgram != nullptr) {
            uint8_t return_value = glIsProgram(program);
            GAPID_INFO("Returned: %" PRIu8 "", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsProgram");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsProgram");
        return false;
    }
}

bool callGlIsProgramPipeline(Stack* stack, bool pushReturn) {
    uint32_t pipeline = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsProgramPipeline(%" PRIu32 ")", pipeline);
        if (glIsProgramPipeline != nullptr) {
            uint8_t return_value = glIsProgramPipeline(pipeline);
            GAPID_INFO("Returned: %" PRIu8 "", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsProgramPipeline");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsProgramPipeline");
        return false;
    }
}

bool callGlIsShader(Stack* stack, bool pushReturn) {
    uint32_t shader = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsShader(%" PRIu32 ")", shader);
        if (glIsShader != nullptr) {
            uint8_t return_value = glIsShader(shader);
            GAPID_INFO("Returned: %" PRIu8 "", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsShader");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsShader");
        return false;
    }
}

bool callGlLinkProgram(Stack* stack, bool pushReturn) {
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glLinkProgram(%" PRIu32 ")", program);
        if (glLinkProgram != nullptr) {
            glLinkProgram(program);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glLinkProgram");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glLinkProgram");
        return false;
    }
}

bool callGlMemoryBarrier(Stack* stack, bool pushReturn) {
    GLbitfield barriers = stack->pop<GLbitfield>();
    if (stack->isValid()) {
        GAPID_INFO("glMemoryBarrier(%u)", barriers);
        if (glMemoryBarrier != nullptr) {
            glMemoryBarrier(barriers);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMemoryBarrier");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMemoryBarrier");
        return false;
    }
}

bool callGlMemoryBarrierByRegion(Stack* stack, bool pushReturn) {
    GLbitfield barriers = stack->pop<GLbitfield>();
    if (stack->isValid()) {
        GAPID_INFO("glMemoryBarrierByRegion(%u)", barriers);
        if (glMemoryBarrierByRegion != nullptr) {
            glMemoryBarrierByRegion(barriers);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMemoryBarrierByRegion");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMemoryBarrierByRegion");
        return false;
    }
}

bool callGlProgramBinary(Stack* stack, bool pushReturn) {
    int32_t length = stack->pop<int32_t>();
    void* binary = stack->pop<void*>();
    GLenum binaryFormat = stack->pop<GLenum>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramBinary(%" PRIu32 ", %u, %p, %" PRId32 ")", program, binaryFormat,
                   binary, length);
        if (glProgramBinary != nullptr) {
            glProgramBinary(program, binaryFormat, binary, length);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramBinary");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramBinary");
        return false;
    }
}

bool callGlProgramParameteri(Stack* stack, bool pushReturn) {
    int32_t value = stack->pop<int32_t>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramParameteri(%" PRIu32 ", %u, %" PRId32 ")", program, pname, value);
        if (glProgramParameteri != nullptr) {
            glProgramParameteri(program, pname, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramParameteri");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramParameteri");
        return false;
    }
}

bool callGlProgramUniform1f(Stack* stack, bool pushReturn) {
    float value0 = stack->pop<float>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform1f(%" PRIu32 ", %" PRId32 ", %f)", program, location, value0);
        if (glProgramUniform1f != nullptr) {
            glProgramUniform1f(program, location, value0);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform1f");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform1f");
        return false;
    }
}

bool callGlProgramUniform1fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform1fv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)", program,
                   location, count, values);
        if (glProgramUniform1fv != nullptr) {
            glProgramUniform1fv(program, location, count, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform1fv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform1fv");
        return false;
    }
}

bool callGlProgramUniform1i(Stack* stack, bool pushReturn) {
    int32_t value0 = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform1i(%" PRIu32 ", %" PRId32 ", %" PRId32 ")", program, location,
                   value0);
        if (glProgramUniform1i != nullptr) {
            glProgramUniform1i(program, location, value0);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform1i");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform1i");
        return false;
    }
}

bool callGlProgramUniform1iv(Stack* stack, bool pushReturn) {
    int32_t* values = stack->pop<int32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform1iv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)", program,
                   location, count, values);
        if (glProgramUniform1iv != nullptr) {
            glProgramUniform1iv(program, location, count, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform1iv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform1iv");
        return false;
    }
}

bool callGlProgramUniform1ui(Stack* stack, bool pushReturn) {
    uint32_t value0 = stack->pop<uint32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform1ui(%" PRIu32 ", %" PRId32 ", %" PRIu32 ")", program, location,
                   value0);
        if (glProgramUniform1ui != nullptr) {
            glProgramUniform1ui(program, location, value0);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform1ui");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform1ui");
        return false;
    }
}

bool callGlProgramUniform1uiv(Stack* stack, bool pushReturn) {
    uint32_t* values = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform1uiv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)", program,
                   location, count, values);
        if (glProgramUniform1uiv != nullptr) {
            glProgramUniform1uiv(program, location, count, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform1uiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform1uiv");
        return false;
    }
}

bool callGlProgramUniform2f(Stack* stack, bool pushReturn) {
    float value1 = stack->pop<float>();
    float value0 = stack->pop<float>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform2f(%" PRIu32 ", %" PRId32 ", %f, %f)", program, location,
                   value0, value1);
        if (glProgramUniform2f != nullptr) {
            glProgramUniform2f(program, location, value0, value1);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform2f");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform2f");
        return false;
    }
}

bool callGlProgramUniform2fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform2fv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)", program,
                   location, count, values);
        if (glProgramUniform2fv != nullptr) {
            glProgramUniform2fv(program, location, count, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform2fv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform2fv");
        return false;
    }
}

bool callGlProgramUniform2i(Stack* stack, bool pushReturn) {
    int32_t value1 = stack->pop<int32_t>();
    int32_t value0 = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform2i(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ")",
                   program, location, value0, value1);
        if (glProgramUniform2i != nullptr) {
            glProgramUniform2i(program, location, value0, value1);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform2i");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform2i");
        return false;
    }
}

bool callGlProgramUniform2iv(Stack* stack, bool pushReturn) {
    int32_t* values = stack->pop<int32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform2iv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)", program,
                   location, count, values);
        if (glProgramUniform2iv != nullptr) {
            glProgramUniform2iv(program, location, count, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform2iv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform2iv");
        return false;
    }
}

bool callGlProgramUniform2ui(Stack* stack, bool pushReturn) {
    uint32_t value1 = stack->pop<uint32_t>();
    uint32_t value0 = stack->pop<uint32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform2ui(%" PRIu32 ", %" PRId32 ", %" PRIu32 ", %" PRIu32 ")",
                   program, location, value0, value1);
        if (glProgramUniform2ui != nullptr) {
            glProgramUniform2ui(program, location, value0, value1);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform2ui");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform2ui");
        return false;
    }
}

bool callGlProgramUniform2uiv(Stack* stack, bool pushReturn) {
    uint32_t* values = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform2uiv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)", program,
                   location, count, values);
        if (glProgramUniform2uiv != nullptr) {
            glProgramUniform2uiv(program, location, count, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform2uiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform2uiv");
        return false;
    }
}

bool callGlProgramUniform3f(Stack* stack, bool pushReturn) {
    float value2 = stack->pop<float>();
    float value1 = stack->pop<float>();
    float value0 = stack->pop<float>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform3f(%" PRIu32 ", %" PRId32 ", %f, %f, %f)", program, location,
                   value0, value1, value2);
        if (glProgramUniform3f != nullptr) {
            glProgramUniform3f(program, location, value0, value1, value2);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform3f");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform3f");
        return false;
    }
}

bool callGlProgramUniform3fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform3fv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)", program,
                   location, count, values);
        if (glProgramUniform3fv != nullptr) {
            glProgramUniform3fv(program, location, count, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform3fv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform3fv");
        return false;
    }
}

bool callGlProgramUniform3i(Stack* stack, bool pushReturn) {
    int32_t value2 = stack->pop<int32_t>();
    int32_t value1 = stack->pop<int32_t>();
    int32_t value0 = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform3i(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32
                   ")",
                   program, location, value0, value1, value2);
        if (glProgramUniform3i != nullptr) {
            glProgramUniform3i(program, location, value0, value1, value2);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform3i");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform3i");
        return false;
    }
}

bool callGlProgramUniform3iv(Stack* stack, bool pushReturn) {
    int32_t* values = stack->pop<int32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform3iv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)", program,
                   location, count, values);
        if (glProgramUniform3iv != nullptr) {
            glProgramUniform3iv(program, location, count, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform3iv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform3iv");
        return false;
    }
}

bool callGlProgramUniform3ui(Stack* stack, bool pushReturn) {
    uint32_t value2 = stack->pop<uint32_t>();
    uint32_t value1 = stack->pop<uint32_t>();
    uint32_t value0 = stack->pop<uint32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform3ui(%" PRIu32 ", %" PRId32 ", %" PRIu32 ", %" PRIu32
                   ", %" PRIu32 ")",
                   program, location, value0, value1, value2);
        if (glProgramUniform3ui != nullptr) {
            glProgramUniform3ui(program, location, value0, value1, value2);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform3ui");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform3ui");
        return false;
    }
}

bool callGlProgramUniform3uiv(Stack* stack, bool pushReturn) {
    uint32_t* values = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform3uiv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)", program,
                   location, count, values);
        if (glProgramUniform3uiv != nullptr) {
            glProgramUniform3uiv(program, location, count, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform3uiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform3uiv");
        return false;
    }
}

bool callGlProgramUniform4f(Stack* stack, bool pushReturn) {
    float value3 = stack->pop<float>();
    float value2 = stack->pop<float>();
    float value1 = stack->pop<float>();
    float value0 = stack->pop<float>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform4f(%" PRIu32 ", %" PRId32 ", %f, %f, %f, %f)", program,
                   location, value0, value1, value2, value3);
        if (glProgramUniform4f != nullptr) {
            glProgramUniform4f(program, location, value0, value1, value2, value3);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform4f");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform4f");
        return false;
    }
}

bool callGlProgramUniform4fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform4fv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)", program,
                   location, count, values);
        if (glProgramUniform4fv != nullptr) {
            glProgramUniform4fv(program, location, count, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform4fv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform4fv");
        return false;
    }
}

bool callGlProgramUniform4i(Stack* stack, bool pushReturn) {
    int32_t value3 = stack->pop<int32_t>();
    int32_t value2 = stack->pop<int32_t>();
    int32_t value1 = stack->pop<int32_t>();
    int32_t value0 = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform4i(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ")",
                   program, location, value0, value1, value2, value3);
        if (glProgramUniform4i != nullptr) {
            glProgramUniform4i(program, location, value0, value1, value2, value3);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform4i");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform4i");
        return false;
    }
}

bool callGlProgramUniform4iv(Stack* stack, bool pushReturn) {
    int32_t* values = stack->pop<int32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform4iv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)", program,
                   location, count, values);
        if (glProgramUniform4iv != nullptr) {
            glProgramUniform4iv(program, location, count, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform4iv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform4iv");
        return false;
    }
}

bool callGlProgramUniform4ui(Stack* stack, bool pushReturn) {
    uint32_t value3 = stack->pop<uint32_t>();
    uint32_t value2 = stack->pop<uint32_t>();
    uint32_t value1 = stack->pop<uint32_t>();
    uint32_t value0 = stack->pop<uint32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform4ui(%" PRIu32 ", %" PRId32 ", %" PRIu32 ", %" PRIu32
                   ", %" PRIu32 ", %" PRIu32 ")",
                   program, location, value0, value1, value2, value3);
        if (glProgramUniform4ui != nullptr) {
            glProgramUniform4ui(program, location, value0, value1, value2, value3);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform4ui");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform4ui");
        return false;
    }
}

bool callGlProgramUniform4uiv(Stack* stack, bool pushReturn) {
    uint32_t* values = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform4uiv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)", program,
                   location, count, values);
        if (glProgramUniform4uiv != nullptr) {
            glProgramUniform4uiv(program, location, count, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform4uiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform4uiv");
        return false;
    }
}

bool callGlProgramUniformMatrix2fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniformMatrix2fv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRIu8
                   ", %p)",
                   program, location, count, transpose, values);
        if (glProgramUniformMatrix2fv != nullptr) {
            glProgramUniformMatrix2fv(program, location, count, transpose, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformMatrix2fv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix2fv");
        return false;
    }
}

bool callGlProgramUniformMatrix2x3fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniformMatrix2x3fv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRIu8
                   ", %p)",
                   program, location, count, transpose, values);
        if (glProgramUniformMatrix2x3fv != nullptr) {
            glProgramUniformMatrix2x3fv(program, location, count, transpose, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformMatrix2x3fv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix2x3fv");
        return false;
    }
}

bool callGlProgramUniformMatrix2x4fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniformMatrix2x4fv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRIu8
                   ", %p)",
                   program, location, count, transpose, values);
        if (glProgramUniformMatrix2x4fv != nullptr) {
            glProgramUniformMatrix2x4fv(program, location, count, transpose, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformMatrix2x4fv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix2x4fv");
        return false;
    }
}

bool callGlProgramUniformMatrix3fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniformMatrix3fv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRIu8
                   ", %p)",
                   program, location, count, transpose, values);
        if (glProgramUniformMatrix3fv != nullptr) {
            glProgramUniformMatrix3fv(program, location, count, transpose, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformMatrix3fv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix3fv");
        return false;
    }
}

bool callGlProgramUniformMatrix3x2fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniformMatrix3x2fv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRIu8
                   ", %p)",
                   program, location, count, transpose, values);
        if (glProgramUniformMatrix3x2fv != nullptr) {
            glProgramUniformMatrix3x2fv(program, location, count, transpose, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformMatrix3x2fv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix3x2fv");
        return false;
    }
}

bool callGlProgramUniformMatrix3x4fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniformMatrix3x4fv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRIu8
                   ", %p)",
                   program, location, count, transpose, values);
        if (glProgramUniformMatrix3x4fv != nullptr) {
            glProgramUniformMatrix3x4fv(program, location, count, transpose, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformMatrix3x4fv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix3x4fv");
        return false;
    }
}

bool callGlProgramUniformMatrix4fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniformMatrix4fv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRIu8
                   ", %p)",
                   program, location, count, transpose, values);
        if (glProgramUniformMatrix4fv != nullptr) {
            glProgramUniformMatrix4fv(program, location, count, transpose, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformMatrix4fv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix4fv");
        return false;
    }
}

bool callGlProgramUniformMatrix4x2fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniformMatrix4x2fv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRIu8
                   ", %p)",
                   program, location, count, transpose, values);
        if (glProgramUniformMatrix4x2fv != nullptr) {
            glProgramUniformMatrix4x2fv(program, location, count, transpose, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformMatrix4x2fv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix4x2fv");
        return false;
    }
}

bool callGlProgramUniformMatrix4x3fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniformMatrix4x3fv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRIu8
                   ", %p)",
                   program, location, count, transpose, values);
        if (glProgramUniformMatrix4x3fv != nullptr) {
            glProgramUniformMatrix4x3fv(program, location, count, transpose, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformMatrix4x3fv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix4x3fv");
        return false;
    }
}

bool callGlReleaseShaderCompiler(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glReleaseShaderCompiler()");
        if (glReleaseShaderCompiler != nullptr) {
            glReleaseShaderCompiler();
        } else {
            GAPID_WARNING("Attempted to call unsupported function glReleaseShaderCompiler");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glReleaseShaderCompiler");
        return false;
    }
}

bool callGlShaderBinary(Stack* stack, bool pushReturn) {
    int32_t binary_size = stack->pop<int32_t>();
    void* binary = stack->pop<void*>();
    GLenum binary_format = stack->pop<GLenum>();
    uint32_t* shaders = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glShaderBinary(%" PRId32 ", %p, %u, %p, %" PRId32 ")", count, shaders,
                   binary_format, binary, binary_size);
        if (glShaderBinary != nullptr) {
            glShaderBinary(count, shaders, binary_format, binary, binary_size);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glShaderBinary");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glShaderBinary");
        return false;
    }
}

bool callGlShaderSource(Stack* stack, bool pushReturn) {
    int32_t* length = stack->pop<int32_t*>();
    char** source = stack->pop<char**>();
    int32_t count = stack->pop<int32_t>();
    uint32_t shader = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glShaderSource(%" PRIu32 ", %" PRId32 ", %p, %p)", shader, count, source,
                   length);
        if (glShaderSource != nullptr) {
            glShaderSource(shader, count, source, length);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glShaderSource");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glShaderSource");
        return false;
    }
}

bool callGlUniform1f(Stack* stack, bool pushReturn) {
    float value = stack->pop<float>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform1f(%" PRId32 ", %f)", location, value);
        if (glUniform1f != nullptr) {
            glUniform1f(location, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform1f");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform1f");
        return false;
    }
}

bool callGlUniform1fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform1fv(%" PRId32 ", %" PRId32 ", %p)", location, count, values);
        if (glUniform1fv != nullptr) {
            glUniform1fv(location, count, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform1fv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform1fv");
        return false;
    }
}

bool callGlUniform1i(Stack* stack, bool pushReturn) {
    int32_t value = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform1i(%" PRId32 ", %" PRId32 ")", location, value);
        if (glUniform1i != nullptr) {
            glUniform1i(location, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform1i");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform1i");
        return false;
    }
}

bool callGlUniform1iv(Stack* stack, bool pushReturn) {
    int32_t* values = stack->pop<int32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform1iv(%" PRId32 ", %" PRId32 ", %p)", location, count, values);
        if (glUniform1iv != nullptr) {
            glUniform1iv(location, count, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform1iv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform1iv");
        return false;
    }
}

bool callGlUniform1ui(Stack* stack, bool pushReturn) {
    uint32_t value0 = stack->pop<uint32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform1ui(%" PRId32 ", %" PRIu32 ")", location, value0);
        if (glUniform1ui != nullptr) {
            glUniform1ui(location, value0);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform1ui");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform1ui");
        return false;
    }
}

bool callGlUniform1uiv(Stack* stack, bool pushReturn) {
    uint32_t* values = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform1uiv(%" PRId32 ", %" PRId32 ", %p)", location, count, values);
        if (glUniform1uiv != nullptr) {
            glUniform1uiv(location, count, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform1uiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform1uiv");
        return false;
    }
}

bool callGlUniform2f(Stack* stack, bool pushReturn) {
    float value1 = stack->pop<float>();
    float value0 = stack->pop<float>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform2f(%" PRId32 ", %f, %f)", location, value0, value1);
        if (glUniform2f != nullptr) {
            glUniform2f(location, value0, value1);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform2f");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform2f");
        return false;
    }
}

bool callGlUniform2fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform2fv(%" PRId32 ", %" PRId32 ", %p)", location, count, values);
        if (glUniform2fv != nullptr) {
            glUniform2fv(location, count, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform2fv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform2fv");
        return false;
    }
}

bool callGlUniform2i(Stack* stack, bool pushReturn) {
    int32_t value1 = stack->pop<int32_t>();
    int32_t value0 = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform2i(%" PRId32 ", %" PRId32 ", %" PRId32 ")", location, value0, value1);
        if (glUniform2i != nullptr) {
            glUniform2i(location, value0, value1);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform2i");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform2i");
        return false;
    }
}

bool callGlUniform2iv(Stack* stack, bool pushReturn) {
    int32_t* values = stack->pop<int32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform2iv(%" PRId32 ", %" PRId32 ", %p)", location, count, values);
        if (glUniform2iv != nullptr) {
            glUniform2iv(location, count, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform2iv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform2iv");
        return false;
    }
}

bool callGlUniform2ui(Stack* stack, bool pushReturn) {
    uint32_t value1 = stack->pop<uint32_t>();
    uint32_t value0 = stack->pop<uint32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform2ui(%" PRId32 ", %" PRIu32 ", %" PRIu32 ")", location, value0, value1);
        if (glUniform2ui != nullptr) {
            glUniform2ui(location, value0, value1);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform2ui");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform2ui");
        return false;
    }
}

bool callGlUniform2uiv(Stack* stack, bool pushReturn) {
    uint32_t* values = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform2uiv(%" PRId32 ", %" PRId32 ", %p)", location, count, values);
        if (glUniform2uiv != nullptr) {
            glUniform2uiv(location, count, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform2uiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform2uiv");
        return false;
    }
}

bool callGlUniform3f(Stack* stack, bool pushReturn) {
    float value2 = stack->pop<float>();
    float value1 = stack->pop<float>();
    float value0 = stack->pop<float>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform3f(%" PRId32 ", %f, %f, %f)", location, value0, value1, value2);
        if (glUniform3f != nullptr) {
            glUniform3f(location, value0, value1, value2);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform3f");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform3f");
        return false;
    }
}

bool callGlUniform3fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform3fv(%" PRId32 ", %" PRId32 ", %p)", location, count, values);
        if (glUniform3fv != nullptr) {
            glUniform3fv(location, count, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform3fv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform3fv");
        return false;
    }
}

bool callGlUniform3i(Stack* stack, bool pushReturn) {
    int32_t value2 = stack->pop<int32_t>();
    int32_t value1 = stack->pop<int32_t>();
    int32_t value0 = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform3i(%" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ")", location,
                   value0, value1, value2);
        if (glUniform3i != nullptr) {
            glUniform3i(location, value0, value1, value2);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform3i");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform3i");
        return false;
    }
}

bool callGlUniform3iv(Stack* stack, bool pushReturn) {
    int32_t* values = stack->pop<int32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform3iv(%" PRId32 ", %" PRId32 ", %p)", location, count, values);
        if (glUniform3iv != nullptr) {
            glUniform3iv(location, count, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform3iv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform3iv");
        return false;
    }
}

bool callGlUniform3ui(Stack* stack, bool pushReturn) {
    uint32_t value2 = stack->pop<uint32_t>();
    uint32_t value1 = stack->pop<uint32_t>();
    uint32_t value0 = stack->pop<uint32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform3ui(%" PRId32 ", %" PRIu32 ", %" PRIu32 ", %" PRIu32 ")", location,
                   value0, value1, value2);
        if (glUniform3ui != nullptr) {
            glUniform3ui(location, value0, value1, value2);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform3ui");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform3ui");
        return false;
    }
}

bool callGlUniform3uiv(Stack* stack, bool pushReturn) {
    uint32_t* values = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform3uiv(%" PRId32 ", %" PRId32 ", %p)", location, count, values);
        if (glUniform3uiv != nullptr) {
            glUniform3uiv(location, count, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform3uiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform3uiv");
        return false;
    }
}

bool callGlUniform4f(Stack* stack, bool pushReturn) {
    float value3 = stack->pop<float>();
    float value2 = stack->pop<float>();
    float value1 = stack->pop<float>();
    float value0 = stack->pop<float>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform4f(%" PRId32 ", %f, %f, %f, %f)", location, value0, value1, value2,
                   value3);
        if (glUniform4f != nullptr) {
            glUniform4f(location, value0, value1, value2, value3);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform4f");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform4f");
        return false;
    }
}

bool callGlUniform4fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform4fv(%" PRId32 ", %" PRId32 ", %p)", location, count, values);
        if (glUniform4fv != nullptr) {
            glUniform4fv(location, count, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform4fv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform4fv");
        return false;
    }
}

bool callGlUniform4i(Stack* stack, bool pushReturn) {
    int32_t value3 = stack->pop<int32_t>();
    int32_t value2 = stack->pop<int32_t>();
    int32_t value1 = stack->pop<int32_t>();
    int32_t value0 = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform4i(%" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ")",
                   location, value0, value1, value2, value3);
        if (glUniform4i != nullptr) {
            glUniform4i(location, value0, value1, value2, value3);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform4i");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform4i");
        return false;
    }
}

bool callGlUniform4iv(Stack* stack, bool pushReturn) {
    int32_t* values = stack->pop<int32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform4iv(%" PRId32 ", %" PRId32 ", %p)", location, count, values);
        if (glUniform4iv != nullptr) {
            glUniform4iv(location, count, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform4iv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform4iv");
        return false;
    }
}

bool callGlUniform4ui(Stack* stack, bool pushReturn) {
    uint32_t value3 = stack->pop<uint32_t>();
    uint32_t value2 = stack->pop<uint32_t>();
    uint32_t value1 = stack->pop<uint32_t>();
    uint32_t value0 = stack->pop<uint32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform4ui(%" PRId32 ", %" PRIu32 ", %" PRIu32 ", %" PRIu32 ", %" PRIu32 ")",
                   location, value0, value1, value2, value3);
        if (glUniform4ui != nullptr) {
            glUniform4ui(location, value0, value1, value2, value3);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform4ui");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform4ui");
        return false;
    }
}

bool callGlUniform4uiv(Stack* stack, bool pushReturn) {
    uint32_t* values = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform4uiv(%" PRId32 ", %" PRId32 ", %p)", location, count, values);
        if (glUniform4uiv != nullptr) {
            glUniform4uiv(location, count, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform4uiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform4uiv");
        return false;
    }
}

bool callGlUniformBlockBinding(Stack* stack, bool pushReturn) {
    uint32_t uniform_block_binding = stack->pop<uint32_t>();
    uint32_t uniform_block_index = stack->pop<uint32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformBlockBinding(%" PRIu32 ", %" PRIu32 ", %" PRIu32 ")", program,
                   uniform_block_index, uniform_block_binding);
        if (glUniformBlockBinding != nullptr) {
            glUniformBlockBinding(program, uniform_block_index, uniform_block_binding);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformBlockBinding");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformBlockBinding");
        return false;
    }
}

bool callGlUniformMatrix2fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformMatrix2fv(%" PRId32 ", %" PRId32 ", %" PRIu8 ", %p)", location, count,
                   transpose, values);
        if (glUniformMatrix2fv != nullptr) {
            glUniformMatrix2fv(location, count, transpose, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformMatrix2fv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformMatrix2fv");
        return false;
    }
}

bool callGlUniformMatrix2x3fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformMatrix2x3fv(%" PRId32 ", %" PRId32 ", %" PRIu8 ", %p)", location,
                   count, transpose, values);
        if (glUniformMatrix2x3fv != nullptr) {
            glUniformMatrix2x3fv(location, count, transpose, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformMatrix2x3fv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformMatrix2x3fv");
        return false;
    }
}

bool callGlUniformMatrix2x4fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformMatrix2x4fv(%" PRId32 ", %" PRId32 ", %" PRIu8 ", %p)", location,
                   count, transpose, values);
        if (glUniformMatrix2x4fv != nullptr) {
            glUniformMatrix2x4fv(location, count, transpose, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformMatrix2x4fv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformMatrix2x4fv");
        return false;
    }
}

bool callGlUniformMatrix3fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformMatrix3fv(%" PRId32 ", %" PRId32 ", %" PRIu8 ", %p)", location, count,
                   transpose, values);
        if (glUniformMatrix3fv != nullptr) {
            glUniformMatrix3fv(location, count, transpose, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformMatrix3fv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformMatrix3fv");
        return false;
    }
}

bool callGlUniformMatrix3x2fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformMatrix3x2fv(%" PRId32 ", %" PRId32 ", %" PRIu8 ", %p)", location,
                   count, transpose, values);
        if (glUniformMatrix3x2fv != nullptr) {
            glUniformMatrix3x2fv(location, count, transpose, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformMatrix3x2fv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformMatrix3x2fv");
        return false;
    }
}

bool callGlUniformMatrix3x4fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformMatrix3x4fv(%" PRId32 ", %" PRId32 ", %" PRIu8 ", %p)", location,
                   count, transpose, values);
        if (glUniformMatrix3x4fv != nullptr) {
            glUniformMatrix3x4fv(location, count, transpose, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformMatrix3x4fv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformMatrix3x4fv");
        return false;
    }
}

bool callGlUniformMatrix4fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformMatrix4fv(%" PRId32 ", %" PRId32 ", %" PRIu8 ", %p)", location, count,
                   transpose, values);
        if (glUniformMatrix4fv != nullptr) {
            glUniformMatrix4fv(location, count, transpose, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformMatrix4fv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformMatrix4fv");
        return false;
    }
}

bool callGlUniformMatrix4x2fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformMatrix4x2fv(%" PRId32 ", %" PRId32 ", %" PRIu8 ", %p)", location,
                   count, transpose, values);
        if (glUniformMatrix4x2fv != nullptr) {
            glUniformMatrix4x2fv(location, count, transpose, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformMatrix4x2fv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformMatrix4x2fv");
        return false;
    }
}

bool callGlUniformMatrix4x3fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformMatrix4x3fv(%" PRId32 ", %" PRId32 ", %" PRIu8 ", %p)", location,
                   count, transpose, values);
        if (glUniformMatrix4x3fv != nullptr) {
            glUniformMatrix4x3fv(location, count, transpose, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformMatrix4x3fv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformMatrix4x3fv");
        return false;
    }
}

bool callGlUseProgram(Stack* stack, bool pushReturn) {
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUseProgram(%" PRIu32 ")", program);
        if (glUseProgram != nullptr) {
            glUseProgram(program);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUseProgram");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUseProgram");
        return false;
    }
}

bool callGlUseProgramStages(Stack* stack, bool pushReturn) {
    uint32_t program = stack->pop<uint32_t>();
    GLbitfield stages = stack->pop<GLbitfield>();
    uint32_t pipeline = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUseProgramStages(%" PRIu32 ", %u, %" PRIu32 ")", pipeline, stages, program);
        if (glUseProgramStages != nullptr) {
            glUseProgramStages(pipeline, stages, program);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUseProgramStages");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUseProgramStages");
        return false;
    }
}

bool callGlValidateProgram(Stack* stack, bool pushReturn) {
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glValidateProgram(%" PRIu32 ")", program);
        if (glValidateProgram != nullptr) {
            glValidateProgram(program);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glValidateProgram");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glValidateProgram");
        return false;
    }
}

bool callGlValidateProgramPipeline(Stack* stack, bool pushReturn) {
    uint32_t pipeline = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glValidateProgramPipeline(%" PRIu32 ")", pipeline);
        if (glValidateProgramPipeline != nullptr) {
            glValidateProgramPipeline(pipeline);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glValidateProgramPipeline");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glValidateProgramPipeline");
        return false;
    }
}

bool callGlCullFace(Stack* stack, bool pushReturn) {
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glCullFace(%u)", mode);
        if (glCullFace != nullptr) {
            glCullFace(mode);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCullFace");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCullFace");
        return false;
    }
}

bool callGlDepthRangef(Stack* stack, bool pushReturn) {
    float far = stack->pop<float>();
    float near = stack->pop<float>();
    if (stack->isValid()) {
        GAPID_INFO("glDepthRangef(%f, %f)", near, far);
        if (glDepthRangef != nullptr) {
            glDepthRangef(near, far);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDepthRangef");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDepthRangef");
        return false;
    }
}

bool callGlFrontFace(Stack* stack, bool pushReturn) {
    GLenum orientation = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glFrontFace(%u)", orientation);
        if (glFrontFace != nullptr) {
            glFrontFace(orientation);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFrontFace");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFrontFace");
        return false;
    }
}

bool callGlGetMultisamplefv(Stack* stack, bool pushReturn) {
    float* val = stack->pop<float*>();
    uint32_t index = stack->pop<uint32_t>();
    GLenum pname = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetMultisamplefv(%u, %" PRIu32 ", %p)", pname, index, val);
        if (glGetMultisamplefv != nullptr) {
            glGetMultisamplefv(pname, index, val);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetMultisamplefv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetMultisamplefv");
        return false;
    }
}

bool callGlLineWidth(Stack* stack, bool pushReturn) {
    float width = stack->pop<float>();
    if (stack->isValid()) {
        GAPID_INFO("glLineWidth(%f)", width);
        if (glLineWidth != nullptr) {
            glLineWidth(width);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glLineWidth");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glLineWidth");
        return false;
    }
}

bool callGlMinSampleShading(Stack* stack, bool pushReturn) {
    float value = stack->pop<float>();
    if (stack->isValid()) {
        GAPID_INFO("glMinSampleShading(%f)", value);
        if (glMinSampleShading != nullptr) {
            glMinSampleShading(value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMinSampleShading");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMinSampleShading");
        return false;
    }
}

bool callGlPolygonOffset(Stack* stack, bool pushReturn) {
    float units = stack->pop<float>();
    float scale_factor = stack->pop<float>();
    if (stack->isValid()) {
        GAPID_INFO("glPolygonOffset(%f, %f)", scale_factor, units);
        if (glPolygonOffset != nullptr) {
            glPolygonOffset(scale_factor, units);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPolygonOffset");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPolygonOffset");
        return false;
    }
}

bool callGlViewport(Stack* stack, bool pushReturn) {
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    int32_t y = stack->pop<int32_t>();
    int32_t x = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glViewport(%" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ")", x, y, width,
                   height);
        if (glViewport != nullptr) {
            glViewport(x, y, width, height);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glViewport");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glViewport");
        return false;
    }
}

bool callGlGetBooleaniV(Stack* stack, bool pushReturn) {
    uint8_t* values = stack->pop<uint8_t*>();
    uint32_t index = stack->pop<uint32_t>();
    GLenum param = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetBooleani_v(%u, %" PRIu32 ", %p)", param, index, values);
        if (glGetBooleani_v != nullptr) {
            glGetBooleani_v(param, index, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetBooleani_v");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetBooleani_v");
        return false;
    }
}

bool callGlGetBooleanv(Stack* stack, bool pushReturn) {
    uint8_t* values = stack->pop<uint8_t*>();
    GLenum param = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetBooleanv(%u, %p)", param, values);
        if (glGetBooleanv != nullptr) {
            glGetBooleanv(param, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetBooleanv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetBooleanv");
        return false;
    }
}

bool callGlGetFloatv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    GLenum param = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetFloatv(%u, %p)", param, values);
        if (glGetFloatv != nullptr) {
            glGetFloatv(param, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetFloatv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetFloatv");
        return false;
    }
}

bool callGlGetInteger64iV(Stack* stack, bool pushReturn) {
    int64_t* values = stack->pop<int64_t*>();
    uint32_t index = stack->pop<uint32_t>();
    GLenum param = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetInteger64i_v(%u, %" PRIu32 ", %p)", param, index, values);
        if (glGetInteger64i_v != nullptr) {
            glGetInteger64i_v(param, index, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetInteger64i_v");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetInteger64i_v");
        return false;
    }
}

bool callGlGetInteger64v(Stack* stack, bool pushReturn) {
    int64_t* values = stack->pop<int64_t*>();
    GLenum param = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetInteger64v(%u, %p)", param, values);
        if (glGetInteger64v != nullptr) {
            glGetInteger64v(param, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetInteger64v");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetInteger64v");
        return false;
    }
}

bool callGlGetIntegeriV(Stack* stack, bool pushReturn) {
    int32_t* values = stack->pop<int32_t*>();
    uint32_t index = stack->pop<uint32_t>();
    GLenum param = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetIntegeri_v(%u, %" PRIu32 ", %p)", param, index, values);
        if (glGetIntegeri_v != nullptr) {
            glGetIntegeri_v(param, index, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetIntegeri_v");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetIntegeri_v");
        return false;
    }
}

bool callGlGetIntegerv(Stack* stack, bool pushReturn) {
    int32_t* values = stack->pop<int32_t*>();
    GLenum param = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetIntegerv(%u, %p)", param, values);
        if (glGetIntegerv != nullptr) {
            glGetIntegerv(param, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetIntegerv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetIntegerv");
        return false;
    }
}

bool callGlGetInternalformativ(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    int32_t bufSize = stack->pop<int32_t>();
    GLenum pname = stack->pop<GLenum>();
    GLenum internalformat = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetInternalformativ(%u, %u, %u, %" PRId32 ", %p)", target, internalformat,
                   pname, bufSize, params);
        if (glGetInternalformativ != nullptr) {
            glGetInternalformativ(target, internalformat, pname, bufSize, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetInternalformativ");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetInternalformativ");
        return false;
    }
}

bool callGlGetString(Stack* stack, bool pushReturn) {
    GLenum param = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetString(%u)", param);
        if (glGetString != nullptr) {
            uint8_t* return_value = glGetString(param);
            GAPID_INFO("Returned: %p", return_value);
            if (pushReturn) {
                stack->push<uint8_t*>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetString");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetString");
        return false;
    }
}

bool callGlGetStringi(Stack* stack, bool pushReturn) {
    uint32_t index = stack->pop<uint32_t>();
    GLenum name = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetStringi(%u, %" PRIu32 ")", name, index);
        if (glGetStringi != nullptr) {
            uint8_t* return_value = glGetStringi(name, index);
            GAPID_INFO("Returned: %p", return_value);
            if (pushReturn) {
                stack->push<uint8_t*>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetStringi");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetStringi");
        return false;
    }
}

bool callGlIsEnabled(Stack* stack, bool pushReturn) {
    GLenum capability = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glIsEnabled(%u)", capability);
        if (glIsEnabled != nullptr) {
            uint8_t return_value = glIsEnabled(capability);
            GAPID_INFO("Returned: %" PRIu8 "", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsEnabled");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsEnabled");
        return false;
    }
}

bool callGlIsEnabledi(Stack* stack, bool pushReturn) {
    uint32_t index = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glIsEnabledi(%u, %" PRIu32 ")", target, index);
        if (glIsEnabledi != nullptr) {
            uint8_t return_value = glIsEnabledi(target, index);
            GAPID_INFO("Returned: %" PRIu8 "", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsEnabledi");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsEnabledi");
        return false;
    }
}

bool callGlClientWaitSync(Stack* stack, bool pushReturn) {
    uint64_t timeout = stack->pop<uint64_t>();
    GLbitfield syncFlags = stack->pop<GLbitfield>();
    uint64_t sync = stack->pop<uint64_t>();
    if (stack->isValid()) {
        GAPID_INFO("glClientWaitSync(%" PRIu64 ", %u, %" PRIu64 ")", sync, syncFlags, timeout);
        if (glClientWaitSync != nullptr) {
            GLenum return_value = glClientWaitSync(sync, syncFlags, timeout);
            GAPID_INFO("Returned: %u", return_value);
            if (pushReturn) {
                stack->push<GLenum>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glClientWaitSync");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glClientWaitSync");
        return false;
    }
}

bool callGlDeleteSync(Stack* stack, bool pushReturn) {
    uint64_t sync = stack->pop<uint64_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteSync(%" PRIu64 ")", sync);
        if (glDeleteSync != nullptr) {
            glDeleteSync(sync);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteSync");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteSync");
        return false;
    }
}

bool callGlFenceSync(Stack* stack, bool pushReturn) {
    GLbitfield syncFlags = stack->pop<GLbitfield>();
    GLenum condition = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glFenceSync(%u, %u)", condition, syncFlags);
        if (glFenceSync != nullptr) {
            uint64_t return_value = glFenceSync(condition, syncFlags);
            GAPID_INFO("Returned: %" PRIu64 "", return_value);
            if (pushReturn) {
                stack->push<uint64_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFenceSync");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFenceSync");
        return false;
    }
}

bool callGlGetSynciv(Stack* stack, bool pushReturn) {
    int32_t* values = stack->pop<int32_t*>();
    int32_t* length = stack->pop<int32_t*>();
    int32_t bufSize = stack->pop<int32_t>();
    GLenum pname = stack->pop<GLenum>();
    uint64_t sync = stack->pop<uint64_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetSynciv(%" PRIu64 ", %u, %" PRId32 ", %p, %p)", sync, pname, bufSize,
                   length, values);
        if (glGetSynciv != nullptr) {
            glGetSynciv(sync, pname, bufSize, length, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetSynciv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetSynciv");
        return false;
    }
}

bool callGlIsSync(Stack* stack, bool pushReturn) {
    uint64_t sync = stack->pop<uint64_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsSync(%" PRIu64 ")", sync);
        if (glIsSync != nullptr) {
            uint8_t return_value = glIsSync(sync);
            GAPID_INFO("Returned: %" PRIu8 "", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsSync");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsSync");
        return false;
    }
}

bool callGlWaitSync(Stack* stack, bool pushReturn) {
    uint64_t timeout = stack->pop<uint64_t>();
    GLbitfield syncFlags = stack->pop<GLbitfield>();
    uint64_t sync = stack->pop<uint64_t>();
    if (stack->isValid()) {
        GAPID_INFO("glWaitSync(%" PRIu64 ", %u, %" PRIu64 ")", sync, syncFlags, timeout);
        if (glWaitSync != nullptr) {
            glWaitSync(sync, syncFlags, timeout);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glWaitSync");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glWaitSync");
        return false;
    }
}

bool callGlActiveTexture(Stack* stack, bool pushReturn) {
    GLenum unit = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glActiveTexture(%u)", unit);
        if (glActiveTexture != nullptr) {
            glActiveTexture(unit);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glActiveTexture");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glActiveTexture");
        return false;
    }
}

bool callGlBindImageTexture(Stack* stack, bool pushReturn) {
    GLenum format = stack->pop<GLenum>();
    GLenum access = stack->pop<GLenum>();
    int32_t layer = stack->pop<int32_t>();
    uint8_t layered = stack->pop<uint8_t>();
    int32_t level = stack->pop<int32_t>();
    uint32_t texture = stack->pop<uint32_t>();
    uint32_t unit = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBindImageTexture(%" PRIu32 ", %" PRIu32 ", %" PRId32 ", %" PRIu8 ", %" PRId32
                   ", %u, %u)",
                   unit, texture, level, layered, layer, access, format);
        if (glBindImageTexture != nullptr) {
            glBindImageTexture(unit, texture, level, layered, layer, access, format);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindImageTexture");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindImageTexture");
        return false;
    }
}

bool callGlBindSampler(Stack* stack, bool pushReturn) {
    uint32_t sampler = stack->pop<uint32_t>();
    uint32_t unit = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBindSampler(%" PRIu32 ", %" PRIu32 ")", unit, sampler);
        if (glBindSampler != nullptr) {
            glBindSampler(unit, sampler);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindSampler");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindSampler");
        return false;
    }
}

bool callGlBindTexture(Stack* stack, bool pushReturn) {
    uint32_t texture = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBindTexture(%u, %" PRIu32 ")", target, texture);
        if (glBindTexture != nullptr) {
            glBindTexture(target, texture);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindTexture");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindTexture");
        return false;
    }
}

bool callGlCompressedTexImage2D(Stack* stack, bool pushReturn) {
    void* data = stack->pop<void*>();
    int32_t image_size = stack->pop<int32_t>();
    int32_t border = stack->pop<int32_t>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    GLenum format = stack->pop<GLenum>();
    int32_t level = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glCompressedTexImage2D(%u, %" PRId32 ", %u, %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %p)",
                   target, level, format, width, height, border, image_size, data);
        if (glCompressedTexImage2D != nullptr) {
            glCompressedTexImage2D(target, level, format, width, height, border, image_size, data);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCompressedTexImage2D");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCompressedTexImage2D");
        return false;
    }
}

bool callGlCompressedTexImage3D(Stack* stack, bool pushReturn) {
    void* data = stack->pop<void*>();
    int32_t image_size = stack->pop<int32_t>();
    int32_t border = stack->pop<int32_t>();
    int32_t depth = stack->pop<int32_t>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    GLenum internalformat = stack->pop<GLenum>();
    int32_t level = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glCompressedTexImage3D(%u, %" PRId32 ", %u, %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %" PRId32 ", %p)",
                   target, level, internalformat, width, height, depth, border, image_size, data);
        if (glCompressedTexImage3D != nullptr) {
            glCompressedTexImage3D(target, level, internalformat, width, height, depth, border,
                                   image_size, data);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCompressedTexImage3D");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCompressedTexImage3D");
        return false;
    }
}

bool callGlCompressedTexSubImage2D(Stack* stack, bool pushReturn) {
    void* data = stack->pop<void*>();
    int32_t image_size = stack->pop<int32_t>();
    GLenum format = stack->pop<GLenum>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    int32_t yoffset = stack->pop<int32_t>();
    int32_t xoffset = stack->pop<int32_t>();
    int32_t level = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glCompressedTexSubImage2D(%u, %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %u, %" PRId32 ", %p)",
                   target, level, xoffset, yoffset, width, height, format, image_size, data);
        if (glCompressedTexSubImage2D != nullptr) {
            glCompressedTexSubImage2D(target, level, xoffset, yoffset, width, height, format,
                                      image_size, data);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCompressedTexSubImage2D");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCompressedTexSubImage2D");
        return false;
    }
}

bool callGlCompressedTexSubImage3D(Stack* stack, bool pushReturn) {
    void* data = stack->pop<void*>();
    int32_t image_size = stack->pop<int32_t>();
    GLenum format = stack->pop<GLenum>();
    int32_t depth = stack->pop<int32_t>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    int32_t zoffset = stack->pop<int32_t>();
    int32_t yoffset = stack->pop<int32_t>();
    int32_t xoffset = stack->pop<int32_t>();
    int32_t level = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glCompressedTexSubImage3D(%u, %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %u, %" PRId32 ", %p)",
                   target, level, xoffset, yoffset, zoffset, width, height, depth, format,
                   image_size, data);
        if (glCompressedTexSubImage3D != nullptr) {
            glCompressedTexSubImage3D(target, level, xoffset, yoffset, zoffset, width, height,
                                      depth, format, image_size, data);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCompressedTexSubImage3D");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCompressedTexSubImage3D");
        return false;
    }
}

bool callGlCopyImageSubData(Stack* stack, bool pushReturn) {
    int32_t srcDepth = stack->pop<int32_t>();
    int32_t srcHeight = stack->pop<int32_t>();
    int32_t srcWidth = stack->pop<int32_t>();
    int32_t dstZ = stack->pop<int32_t>();
    int32_t dstY = stack->pop<int32_t>();
    int32_t dstX = stack->pop<int32_t>();
    int32_t dstLevel = stack->pop<int32_t>();
    GLenum dstTarget = stack->pop<GLenum>();
    uint32_t dstName = stack->pop<uint32_t>();
    int32_t srcZ = stack->pop<int32_t>();
    int32_t srcY = stack->pop<int32_t>();
    int32_t srcX = stack->pop<int32_t>();
    int32_t srcLevel = stack->pop<int32_t>();
    GLenum srcTarget = stack->pop<GLenum>();
    uint32_t srcName = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glCopyImageSubData(%" PRIu32 ", %u, %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %" PRIu32 ", %u, %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %" PRId32 ", %" PRId32 ")",
                   srcName, srcTarget, srcLevel, srcX, srcY, srcZ, dstName, dstTarget, dstLevel,
                   dstX, dstY, dstZ, srcWidth, srcHeight, srcDepth);
        if (glCopyImageSubData != nullptr) {
            glCopyImageSubData(srcName, srcTarget, srcLevel, srcX, srcY, srcZ, dstName, dstTarget,
                               dstLevel, dstX, dstY, dstZ, srcWidth, srcHeight, srcDepth);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCopyImageSubData");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCopyImageSubData");
        return false;
    }
}

bool callGlCopyTexImage2D(Stack* stack, bool pushReturn) {
    int32_t border = stack->pop<int32_t>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    int32_t y = stack->pop<int32_t>();
    int32_t x = stack->pop<int32_t>();
    GLenum format = stack->pop<GLenum>();
    int32_t level = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glCopyTexImage2D(%u, %" PRId32 ", %u, %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %" PRId32 ")",
                   target, level, format, x, y, width, height, border);
        if (glCopyTexImage2D != nullptr) {
            glCopyTexImage2D(target, level, format, x, y, width, height, border);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCopyTexImage2D");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCopyTexImage2D");
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
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glCopyTexSubImage2D(%u, %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %" PRId32 ", %" PRId32 ")",
                   target, level, xoffset, yoffset, x, y, width, height);
        if (glCopyTexSubImage2D != nullptr) {
            glCopyTexSubImage2D(target, level, xoffset, yoffset, x, y, width, height);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCopyTexSubImage2D");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCopyTexSubImage2D");
        return false;
    }
}

bool callGlCopyTexSubImage3D(Stack* stack, bool pushReturn) {
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    int32_t y = stack->pop<int32_t>();
    int32_t x = stack->pop<int32_t>();
    int32_t zoffset = stack->pop<int32_t>();
    int32_t yoffset = stack->pop<int32_t>();
    int32_t xoffset = stack->pop<int32_t>();
    int32_t level = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glCopyTexSubImage3D(%u, %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ")",
                   target, level, xoffset, yoffset, zoffset, x, y, width, height);
        if (glCopyTexSubImage3D != nullptr) {
            glCopyTexSubImage3D(target, level, xoffset, yoffset, zoffset, x, y, width, height);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCopyTexSubImage3D");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCopyTexSubImage3D");
        return false;
    }
}

bool callGlDeleteSamplers(Stack* stack, bool pushReturn) {
    uint32_t* samplers = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteSamplers(%" PRId32 ", %p)", count, samplers);
        if (glDeleteSamplers != nullptr) {
            glDeleteSamplers(count, samplers);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteSamplers");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteSamplers");
        return false;
    }
}

bool callGlDeleteTextures(Stack* stack, bool pushReturn) {
    uint32_t* textures = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteTextures(%" PRId32 ", %p)", count, textures);
        if (glDeleteTextures != nullptr) {
            glDeleteTextures(count, textures);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteTextures");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteTextures");
        return false;
    }
}

bool callGlGenSamplers(Stack* stack, bool pushReturn) {
    uint32_t* samplers = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenSamplers(%" PRId32 ", %p)", count, samplers);
        if (glGenSamplers != nullptr) {
            glGenSamplers(count, samplers);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenSamplers");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenSamplers");
        return false;
    }
}

bool callGlGenTextures(Stack* stack, bool pushReturn) {
    uint32_t* textures = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenTextures(%" PRId32 ", %p)", count, textures);
        if (glGenTextures != nullptr) {
            glGenTextures(count, textures);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenTextures");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenTextures");
        return false;
    }
}

bool callGlGenerateMipmap(Stack* stack, bool pushReturn) {
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGenerateMipmap(%u)", target);
        if (glGenerateMipmap != nullptr) {
            glGenerateMipmap(target);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenerateMipmap");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenerateMipmap");
        return false;
    }
}

bool callGlGetSamplerParameterIiv(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetSamplerParameterIiv(%" PRIu32 ", %u, %p)", sampler, pname, params);
        if (glGetSamplerParameterIiv != nullptr) {
            glGetSamplerParameterIiv(sampler, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetSamplerParameterIiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetSamplerParameterIiv");
        return false;
    }
}

bool callGlGetSamplerParameterIuiv(Stack* stack, bool pushReturn) {
    uint32_t* params = stack->pop<uint32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetSamplerParameterIuiv(%" PRIu32 ", %u, %p)", sampler, pname, params);
        if (glGetSamplerParameterIuiv != nullptr) {
            glGetSamplerParameterIuiv(sampler, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetSamplerParameterIuiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetSamplerParameterIuiv");
        return false;
    }
}

bool callGlGetSamplerParameterfv(Stack* stack, bool pushReturn) {
    float* params = stack->pop<float*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetSamplerParameterfv(%" PRIu32 ", %u, %p)", sampler, pname, params);
        if (glGetSamplerParameterfv != nullptr) {
            glGetSamplerParameterfv(sampler, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetSamplerParameterfv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetSamplerParameterfv");
        return false;
    }
}

bool callGlGetSamplerParameteriv(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetSamplerParameteriv(%" PRIu32 ", %u, %p)", sampler, pname, params);
        if (glGetSamplerParameteriv != nullptr) {
            glGetSamplerParameteriv(sampler, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetSamplerParameteriv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetSamplerParameteriv");
        return false;
    }
}

bool callGlGetTexLevelParameterfv(Stack* stack, bool pushReturn) {
    float* params = stack->pop<float*>();
    GLenum pname = stack->pop<GLenum>();
    int32_t level = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetTexLevelParameterfv(%u, %" PRId32 ", %u, %p)", target, level, pname,
                   params);
        if (glGetTexLevelParameterfv != nullptr) {
            glGetTexLevelParameterfv(target, level, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetTexLevelParameterfv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetTexLevelParameterfv");
        return false;
    }
}

bool callGlGetTexLevelParameteriv(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    int32_t level = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetTexLevelParameteriv(%u, %" PRId32 ", %u, %p)", target, level, pname,
                   params);
        if (glGetTexLevelParameteriv != nullptr) {
            glGetTexLevelParameteriv(target, level, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetTexLevelParameteriv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetTexLevelParameteriv");
        return false;
    }
}

bool callGlGetTexParameterIiv(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetTexParameterIiv(%u, %u, %p)", target, pname, params);
        if (glGetTexParameterIiv != nullptr) {
            glGetTexParameterIiv(target, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetTexParameterIiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetTexParameterIiv");
        return false;
    }
}

bool callGlGetTexParameterIuiv(Stack* stack, bool pushReturn) {
    uint32_t* params = stack->pop<uint32_t*>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetTexParameterIuiv(%u, %u, %p)", target, pname, params);
        if (glGetTexParameterIuiv != nullptr) {
            glGetTexParameterIuiv(target, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetTexParameterIuiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetTexParameterIuiv");
        return false;
    }
}

bool callGlGetTexParameterfv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    GLenum parameter = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetTexParameterfv(%u, %u, %p)", target, parameter, values);
        if (glGetTexParameterfv != nullptr) {
            glGetTexParameterfv(target, parameter, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetTexParameterfv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetTexParameterfv");
        return false;
    }
}

bool callGlGetTexParameteriv(Stack* stack, bool pushReturn) {
    int32_t* values = stack->pop<int32_t*>();
    GLenum parameter = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetTexParameteriv(%u, %u, %p)", target, parameter, values);
        if (glGetTexParameteriv != nullptr) {
            glGetTexParameteriv(target, parameter, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetTexParameteriv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetTexParameteriv");
        return false;
    }
}

bool callGlIsSampler(Stack* stack, bool pushReturn) {
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsSampler(%" PRIu32 ")", sampler);
        if (glIsSampler != nullptr) {
            uint8_t return_value = glIsSampler(sampler);
            GAPID_INFO("Returned: %" PRIu8 "", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsSampler");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsSampler");
        return false;
    }
}

bool callGlIsTexture(Stack* stack, bool pushReturn) {
    uint32_t texture = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsTexture(%" PRIu32 ")", texture);
        if (glIsTexture != nullptr) {
            uint8_t return_value = glIsTexture(texture);
            GAPID_INFO("Returned: %" PRIu8 "", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsTexture");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsTexture");
        return false;
    }
}

bool callGlPixelStorei(Stack* stack, bool pushReturn) {
    int32_t value = stack->pop<int32_t>();
    GLenum parameter = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glPixelStorei(%u, %" PRId32 ")", parameter, value);
        if (glPixelStorei != nullptr) {
            glPixelStorei(parameter, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPixelStorei");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPixelStorei");
        return false;
    }
}

bool callGlSamplerParameterIiv(Stack* stack, bool pushReturn) {
    int32_t* param = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glSamplerParameterIiv(%" PRIu32 ", %u, %p)", sampler, pname, param);
        if (glSamplerParameterIiv != nullptr) {
            glSamplerParameterIiv(sampler, pname, param);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glSamplerParameterIiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glSamplerParameterIiv");
        return false;
    }
}

bool callGlSamplerParameterIuiv(Stack* stack, bool pushReturn) {
    uint32_t* param = stack->pop<uint32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glSamplerParameterIuiv(%" PRIu32 ", %u, %p)", sampler, pname, param);
        if (glSamplerParameterIuiv != nullptr) {
            glSamplerParameterIuiv(sampler, pname, param);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glSamplerParameterIuiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glSamplerParameterIuiv");
        return false;
    }
}

bool callGlSamplerParameterf(Stack* stack, bool pushReturn) {
    float param = stack->pop<float>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glSamplerParameterf(%" PRIu32 ", %u, %f)", sampler, pname, param);
        if (glSamplerParameterf != nullptr) {
            glSamplerParameterf(sampler, pname, param);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glSamplerParameterf");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glSamplerParameterf");
        return false;
    }
}

bool callGlSamplerParameterfv(Stack* stack, bool pushReturn) {
    float* param = stack->pop<float*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glSamplerParameterfv(%" PRIu32 ", %u, %p)", sampler, pname, param);
        if (glSamplerParameterfv != nullptr) {
            glSamplerParameterfv(sampler, pname, param);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glSamplerParameterfv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glSamplerParameterfv");
        return false;
    }
}

bool callGlSamplerParameteri(Stack* stack, bool pushReturn) {
    int32_t param = stack->pop<int32_t>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glSamplerParameteri(%" PRIu32 ", %u, %" PRId32 ")", sampler, pname, param);
        if (glSamplerParameteri != nullptr) {
            glSamplerParameteri(sampler, pname, param);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glSamplerParameteri");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glSamplerParameteri");
        return false;
    }
}

bool callGlSamplerParameteriv(Stack* stack, bool pushReturn) {
    int32_t* param = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glSamplerParameteriv(%" PRIu32 ", %u, %p)", sampler, pname, param);
        if (glSamplerParameteriv != nullptr) {
            glSamplerParameteriv(sampler, pname, param);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glSamplerParameteriv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glSamplerParameteriv");
        return false;
    }
}

bool callGlTexBuffer(Stack* stack, bool pushReturn) {
    uint32_t buffer = stack->pop<uint32_t>();
    GLenum internalformat = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexBuffer(%u, %u, %" PRIu32 ")", target, internalformat, buffer);
        if (glTexBuffer != nullptr) {
            glTexBuffer(target, internalformat, buffer);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexBuffer");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexBuffer");
        return false;
    }
}

bool callGlTexBufferRange(Stack* stack, bool pushReturn) {
    int32_t size = stack->pop<int32_t>();
    int32_t offset = stack->pop<int32_t>();
    uint32_t buffer = stack->pop<uint32_t>();
    GLenum internalformat = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexBufferRange(%u, %u, %" PRIu32 ", %" PRId32 ", %" PRId32 ")", target,
                   internalformat, buffer, offset, size);
        if (glTexBufferRange != nullptr) {
            glTexBufferRange(target, internalformat, buffer, offset, size);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexBufferRange");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexBufferRange");
        return false;
    }
}

bool callGlTexImage2D(Stack* stack, bool pushReturn) {
    void* data = stack->pop<void*>();
    GLenum type = stack->pop<GLenum>();
    GLenum format = stack->pop<GLenum>();
    int32_t border = stack->pop<int32_t>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    int32_t internal_format = stack->pop<int32_t>();
    int32_t level = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexImage2D(%u, %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %u, %u, %p)",
                   target, level, internal_format, width, height, border, format, type, data);
        if (glTexImage2D != nullptr) {
            glTexImage2D(target, level, internal_format, width, height, border, format, type, data);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexImage2D");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexImage2D");
        return false;
    }
}

bool callGlTexImage3D(Stack* stack, bool pushReturn) {
    void* data = stack->pop<void*>();
    GLenum type = stack->pop<GLenum>();
    GLenum format = stack->pop<GLenum>();
    int32_t border = stack->pop<int32_t>();
    int32_t depth = stack->pop<int32_t>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    int32_t internalformat = stack->pop<int32_t>();
    int32_t level = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexImage3D(%u, %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %u, %u, %p)",
                   target, level, internalformat, width, height, depth, border, format, type, data);
        if (glTexImage3D != nullptr) {
            glTexImage3D(target, level, internalformat, width, height, depth, border, format, type,
                         data);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexImage3D");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexImage3D");
        return false;
    }
}

bool callGlTexParameterIiv(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexParameterIiv(%u, %u, %p)", target, pname, params);
        if (glTexParameterIiv != nullptr) {
            glTexParameterIiv(target, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexParameterIiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexParameterIiv");
        return false;
    }
}

bool callGlTexParameterIuiv(Stack* stack, bool pushReturn) {
    uint32_t* params = stack->pop<uint32_t*>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexParameterIuiv(%u, %u, %p)", target, pname, params);
        if (glTexParameterIuiv != nullptr) {
            glTexParameterIuiv(target, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexParameterIuiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexParameterIuiv");
        return false;
    }
}

bool callGlTexParameterf(Stack* stack, bool pushReturn) {
    float value = stack->pop<float>();
    GLenum parameter = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexParameterf(%u, %u, %f)", target, parameter, value);
        if (glTexParameterf != nullptr) {
            glTexParameterf(target, parameter, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexParameterf");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexParameterf");
        return false;
    }
}

bool callGlTexParameterfv(Stack* stack, bool pushReturn) {
    float* params = stack->pop<float*>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexParameterfv(%u, %u, %p)", target, pname, params);
        if (glTexParameterfv != nullptr) {
            glTexParameterfv(target, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexParameterfv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexParameterfv");
        return false;
    }
}

bool callGlTexParameteri(Stack* stack, bool pushReturn) {
    int32_t value = stack->pop<int32_t>();
    GLenum parameter = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexParameteri(%u, %u, %" PRId32 ")", target, parameter, value);
        if (glTexParameteri != nullptr) {
            glTexParameteri(target, parameter, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexParameteri");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexParameteri");
        return false;
    }
}

bool callGlTexParameteriv(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexParameteriv(%u, %u, %p)", target, pname, params);
        if (glTexParameteriv != nullptr) {
            glTexParameteriv(target, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexParameteriv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexParameteriv");
        return false;
    }
}

bool callGlTexStorage2D(Stack* stack, bool pushReturn) {
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    GLenum internalformat = stack->pop<GLenum>();
    int32_t levels = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexStorage2D(%u, %" PRId32 ", %u, %" PRId32 ", %" PRId32 ")", target, levels,
                   internalformat, width, height);
        if (glTexStorage2D != nullptr) {
            glTexStorage2D(target, levels, internalformat, width, height);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexStorage2D");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexStorage2D");
        return false;
    }
}

bool callGlTexStorage2DMultisample(Stack* stack, bool pushReturn) {
    uint8_t fixedsamplelocations = stack->pop<uint8_t>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    GLenum internalformat = stack->pop<GLenum>();
    int32_t samples = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexStorage2DMultisample(%u, %" PRId32 ", %u, %" PRId32 ", %" PRId32
                   ", %" PRIu8 ")",
                   target, samples, internalformat, width, height, fixedsamplelocations);
        if (glTexStorage2DMultisample != nullptr) {
            glTexStorage2DMultisample(target, samples, internalformat, width, height,
                                      fixedsamplelocations);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexStorage2DMultisample");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexStorage2DMultisample");
        return false;
    }
}

bool callGlTexStorage3D(Stack* stack, bool pushReturn) {
    int32_t depth = stack->pop<int32_t>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    GLenum internalformat = stack->pop<GLenum>();
    int32_t levels = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexStorage3D(%u, %" PRId32 ", %u, %" PRId32 ", %" PRId32 ", %" PRId32 ")",
                   target, levels, internalformat, width, height, depth);
        if (glTexStorage3D != nullptr) {
            glTexStorage3D(target, levels, internalformat, width, height, depth);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexStorage3D");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexStorage3D");
        return false;
    }
}

bool callGlTexStorage3DMultisample(Stack* stack, bool pushReturn) {
    uint8_t fixedsamplelocations = stack->pop<uint8_t>();
    int32_t depth = stack->pop<int32_t>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    GLenum internalformat = stack->pop<GLenum>();
    int32_t samples = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexStorage3DMultisample(%u, %" PRId32 ", %u, %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %" PRIu8 ")",
                   target, samples, internalformat, width, height, depth, fixedsamplelocations);
        if (glTexStorage3DMultisample != nullptr) {
            glTexStorage3DMultisample(target, samples, internalformat, width, height, depth,
                                      fixedsamplelocations);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexStorage3DMultisample");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexStorage3DMultisample");
        return false;
    }
}

bool callGlTexSubImage2D(Stack* stack, bool pushReturn) {
    void* data = stack->pop<void*>();
    GLenum type = stack->pop<GLenum>();
    GLenum format = stack->pop<GLenum>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    int32_t yoffset = stack->pop<int32_t>();
    int32_t xoffset = stack->pop<int32_t>();
    int32_t level = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexSubImage2D(%u, %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %u, %u, %p)",
                   target, level, xoffset, yoffset, width, height, format, type, data);
        if (glTexSubImage2D != nullptr) {
            glTexSubImage2D(target, level, xoffset, yoffset, width, height, format, type, data);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexSubImage2D");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexSubImage2D");
        return false;
    }
}

bool callGlTexSubImage3D(Stack* stack, bool pushReturn) {
    void* data = stack->pop<void*>();
    GLenum type = stack->pop<GLenum>();
    GLenum format = stack->pop<GLenum>();
    int32_t depth = stack->pop<int32_t>();
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    int32_t zoffset = stack->pop<int32_t>();
    int32_t yoffset = stack->pop<int32_t>();
    int32_t xoffset = stack->pop<int32_t>();
    int32_t level = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexSubImage3D(%u, %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %u, %u, %p)",
                   target, level, xoffset, yoffset, zoffset, width, height, depth, format, type,
                   data);
        if (glTexSubImage3D != nullptr) {
            glTexSubImage3D(target, level, xoffset, yoffset, zoffset, width, height, depth, format,
                            type, data);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexSubImage3D");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexSubImage3D");
        return false;
    }
}

bool callGlBeginTransformFeedback(Stack* stack, bool pushReturn) {
    GLenum primitiveMode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBeginTransformFeedback(%u)", primitiveMode);
        if (glBeginTransformFeedback != nullptr) {
            glBeginTransformFeedback(primitiveMode);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBeginTransformFeedback");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBeginTransformFeedback");
        return false;
    }
}

bool callGlBindTransformFeedback(Stack* stack, bool pushReturn) {
    uint32_t id = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBindTransformFeedback(%u, %" PRIu32 ")", target, id);
        if (glBindTransformFeedback != nullptr) {
            glBindTransformFeedback(target, id);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindTransformFeedback");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindTransformFeedback");
        return false;
    }
}

bool callGlDeleteTransformFeedbacks(Stack* stack, bool pushReturn) {
    uint32_t* ids = stack->pop<uint32_t*>();
    int32_t n = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteTransformFeedbacks(%" PRId32 ", %p)", n, ids);
        if (glDeleteTransformFeedbacks != nullptr) {
            glDeleteTransformFeedbacks(n, ids);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteTransformFeedbacks");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteTransformFeedbacks");
        return false;
    }
}

bool callGlEndTransformFeedback(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glEndTransformFeedback()");
        if (glEndTransformFeedback != nullptr) {
            glEndTransformFeedback();
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEndTransformFeedback");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEndTransformFeedback");
        return false;
    }
}

bool callGlGenTransformFeedbacks(Stack* stack, bool pushReturn) {
    uint32_t* ids = stack->pop<uint32_t*>();
    int32_t n = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenTransformFeedbacks(%" PRId32 ", %p)", n, ids);
        if (glGenTransformFeedbacks != nullptr) {
            glGenTransformFeedbacks(n, ids);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenTransformFeedbacks");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenTransformFeedbacks");
        return false;
    }
}

bool callGlGetTransformFeedbackVarying(Stack* stack, bool pushReturn) {
    char* name = stack->pop<char*>();
    GLenum* type = stack->pop<GLenum*>();
    int32_t* size = stack->pop<int32_t*>();
    int32_t* length = stack->pop<int32_t*>();
    int32_t bufSize = stack->pop<int32_t>();
    uint32_t index = stack->pop<uint32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetTransformFeedbackVarying(%" PRIu32 ", %" PRIu32 ", %" PRId32
                   ", %p, %p, %p, %p)",
                   program, index, bufSize, length, size, type, name);
        if (glGetTransformFeedbackVarying != nullptr) {
            glGetTransformFeedbackVarying(program, index, bufSize, length, size, type, name);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetTransformFeedbackVarying");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetTransformFeedbackVarying");
        return false;
    }
}

bool callGlIsTransformFeedback(Stack* stack, bool pushReturn) {
    uint32_t id = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsTransformFeedback(%" PRIu32 ")", id);
        if (glIsTransformFeedback != nullptr) {
            uint8_t return_value = glIsTransformFeedback(id);
            GAPID_INFO("Returned: %" PRIu8 "", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsTransformFeedback");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsTransformFeedback");
        return false;
    }
}

bool callGlPauseTransformFeedback(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glPauseTransformFeedback()");
        if (glPauseTransformFeedback != nullptr) {
            glPauseTransformFeedback();
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPauseTransformFeedback");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPauseTransformFeedback");
        return false;
    }
}

bool callGlResumeTransformFeedback(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glResumeTransformFeedback()");
        if (glResumeTransformFeedback != nullptr) {
            glResumeTransformFeedback();
        } else {
            GAPID_WARNING("Attempted to call unsupported function glResumeTransformFeedback");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glResumeTransformFeedback");
        return false;
    }
}

bool callGlTransformFeedbackVaryings(Stack* stack, bool pushReturn) {
    GLenum bufferMode = stack->pop<GLenum>();
    char** varyings = stack->pop<char**>();
    int32_t count = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glTransformFeedbackVaryings(%" PRIu32 ", %" PRId32 ", %p, %u)", program, count,
                   varyings, bufferMode);
        if (glTransformFeedbackVaryings != nullptr) {
            glTransformFeedbackVaryings(program, count, varyings, bufferMode);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTransformFeedbackVaryings");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTransformFeedbackVaryings");
        return false;
    }
}

bool callGlBindVertexArray(Stack* stack, bool pushReturn) {
    uint32_t array = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBindVertexArray(%" PRIu32 ")", array);
        if (glBindVertexArray != nullptr) {
            glBindVertexArray(array);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindVertexArray");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindVertexArray");
        return false;
    }
}

bool callGlBindVertexBuffer(Stack* stack, bool pushReturn) {
    int32_t stride = stack->pop<int32_t>();
    int32_t offset = stack->pop<int32_t>();
    uint32_t buffer = stack->pop<uint32_t>();
    uint32_t binding_index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBindVertexBuffer(%" PRIu32 ", %" PRIu32 ", %" PRId32 ", %" PRId32 ")",
                   binding_index, buffer, offset, stride);
        if (glBindVertexBuffer != nullptr) {
            glBindVertexBuffer(binding_index, buffer, offset, stride);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindVertexBuffer");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindVertexBuffer");
        return false;
    }
}

bool callGlDeleteVertexArrays(Stack* stack, bool pushReturn) {
    uint32_t* arrays = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteVertexArrays(%" PRId32 ", %p)", count, arrays);
        if (glDeleteVertexArrays != nullptr) {
            glDeleteVertexArrays(count, arrays);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteVertexArrays");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteVertexArrays");
        return false;
    }
}

bool callGlDisableVertexAttribArray(Stack* stack, bool pushReturn) {
    uint32_t location = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDisableVertexAttribArray(%" PRIu32 ")", location);
        if (glDisableVertexAttribArray != nullptr) {
            glDisableVertexAttribArray(location);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDisableVertexAttribArray");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDisableVertexAttribArray");
        return false;
    }
}

bool callGlEnableVertexAttribArray(Stack* stack, bool pushReturn) {
    uint32_t location = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glEnableVertexAttribArray(%" PRIu32 ")", location);
        if (glEnableVertexAttribArray != nullptr) {
            glEnableVertexAttribArray(location);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEnableVertexAttribArray");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEnableVertexAttribArray");
        return false;
    }
}

bool callGlGenVertexArrays(Stack* stack, bool pushReturn) {
    uint32_t* arrays = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenVertexArrays(%" PRId32 ", %p)", count, arrays);
        if (glGenVertexArrays != nullptr) {
            glGenVertexArrays(count, arrays);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenVertexArrays");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenVertexArrays");
        return false;
    }
}

bool callGlGetVertexAttribIiv(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetVertexAttribIiv(%" PRIu32 ", %u, %p)", index, pname, params);
        if (glGetVertexAttribIiv != nullptr) {
            glGetVertexAttribIiv(index, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetVertexAttribIiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetVertexAttribIiv");
        return false;
    }
}

bool callGlGetVertexAttribIuiv(Stack* stack, bool pushReturn) {
    uint32_t* params = stack->pop<uint32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetVertexAttribIuiv(%" PRIu32 ", %u, %p)", index, pname, params);
        if (glGetVertexAttribIuiv != nullptr) {
            glGetVertexAttribIuiv(index, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetVertexAttribIuiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetVertexAttribIuiv");
        return false;
    }
}

bool callGlGetVertexAttribPointerv(Stack* stack, bool pushReturn) {
    void** pointer = stack->pop<void**>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetVertexAttribPointerv(%" PRIu32 ", %u, %p)", index, pname, pointer);
        if (glGetVertexAttribPointerv != nullptr) {
            glGetVertexAttribPointerv(index, pname, pointer);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetVertexAttribPointerv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetVertexAttribPointerv");
        return false;
    }
}

bool callGlGetVertexAttribfv(Stack* stack, bool pushReturn) {
    float* params = stack->pop<float*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetVertexAttribfv(%" PRIu32 ", %u, %p)", index, pname, params);
        if (glGetVertexAttribfv != nullptr) {
            glGetVertexAttribfv(index, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetVertexAttribfv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetVertexAttribfv");
        return false;
    }
}

bool callGlGetVertexAttribiv(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetVertexAttribiv(%" PRIu32 ", %u, %p)", index, pname, params);
        if (glGetVertexAttribiv != nullptr) {
            glGetVertexAttribiv(index, pname, params);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetVertexAttribiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetVertexAttribiv");
        return false;
    }
}

bool callGlIsVertexArray(Stack* stack, bool pushReturn) {
    uint32_t array = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsVertexArray(%" PRIu32 ")", array);
        if (glIsVertexArray != nullptr) {
            uint8_t return_value = glIsVertexArray(array);
            GAPID_INFO("Returned: %" PRIu8 "", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsVertexArray");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsVertexArray");
        return false;
    }
}

bool callGlVertexAttrib1f(Stack* stack, bool pushReturn) {
    float value0 = stack->pop<float>();
    uint32_t location = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttrib1f(%" PRIu32 ", %f)", location, value0);
        if (glVertexAttrib1f != nullptr) {
            glVertexAttrib1f(location, value0);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttrib1f");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttrib1f");
        return false;
    }
}

bool callGlVertexAttrib1fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint32_t location = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttrib1fv(%" PRIu32 ", %p)", location, value);
        if (glVertexAttrib1fv != nullptr) {
            glVertexAttrib1fv(location, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttrib1fv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttrib1fv");
        return false;
    }
}

bool callGlVertexAttrib2f(Stack* stack, bool pushReturn) {
    float value1 = stack->pop<float>();
    float value0 = stack->pop<float>();
    uint32_t location = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttrib2f(%" PRIu32 ", %f, %f)", location, value0, value1);
        if (glVertexAttrib2f != nullptr) {
            glVertexAttrib2f(location, value0, value1);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttrib2f");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttrib2f");
        return false;
    }
}

bool callGlVertexAttrib2fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint32_t location = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttrib2fv(%" PRIu32 ", %p)", location, value);
        if (glVertexAttrib2fv != nullptr) {
            glVertexAttrib2fv(location, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttrib2fv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttrib2fv");
        return false;
    }
}

bool callGlVertexAttrib3f(Stack* stack, bool pushReturn) {
    float value2 = stack->pop<float>();
    float value1 = stack->pop<float>();
    float value0 = stack->pop<float>();
    uint32_t location = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttrib3f(%" PRIu32 ", %f, %f, %f)", location, value0, value1, value2);
        if (glVertexAttrib3f != nullptr) {
            glVertexAttrib3f(location, value0, value1, value2);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttrib3f");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttrib3f");
        return false;
    }
}

bool callGlVertexAttrib3fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint32_t location = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttrib3fv(%" PRIu32 ", %p)", location, value);
        if (glVertexAttrib3fv != nullptr) {
            glVertexAttrib3fv(location, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttrib3fv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttrib3fv");
        return false;
    }
}

bool callGlVertexAttrib4f(Stack* stack, bool pushReturn) {
    float value3 = stack->pop<float>();
    float value2 = stack->pop<float>();
    float value1 = stack->pop<float>();
    float value0 = stack->pop<float>();
    uint32_t location = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttrib4f(%" PRIu32 ", %f, %f, %f, %f)", location, value0, value1,
                   value2, value3);
        if (glVertexAttrib4f != nullptr) {
            glVertexAttrib4f(location, value0, value1, value2, value3);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttrib4f");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttrib4f");
        return false;
    }
}

bool callGlVertexAttrib4fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint32_t location = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttrib4fv(%" PRIu32 ", %p)", location, value);
        if (glVertexAttrib4fv != nullptr) {
            glVertexAttrib4fv(location, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttrib4fv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttrib4fv");
        return false;
    }
}

bool callGlVertexAttribBinding(Stack* stack, bool pushReturn) {
    uint32_t binding_index = stack->pop<uint32_t>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttribBinding(%" PRIu32 ", %" PRIu32 ")", index, binding_index);
        if (glVertexAttribBinding != nullptr) {
            glVertexAttribBinding(index, binding_index);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttribBinding");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttribBinding");
        return false;
    }
}

bool callGlVertexAttribDivisor(Stack* stack, bool pushReturn) {
    uint32_t divisor = stack->pop<uint32_t>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttribDivisor(%" PRIu32 ", %" PRIu32 ")", index, divisor);
        if (glVertexAttribDivisor != nullptr) {
            glVertexAttribDivisor(index, divisor);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttribDivisor");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttribDivisor");
        return false;
    }
}

bool callGlVertexAttribFormat(Stack* stack, bool pushReturn) {
    uint32_t relativeoffset = stack->pop<uint32_t>();
    uint8_t normalized = stack->pop<uint8_t>();
    GLenum type = stack->pop<GLenum>();
    int32_t size = stack->pop<int32_t>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttribFormat(%" PRIu32 ", %" PRId32 ", %u, %" PRIu8 ", %" PRIu32 ")",
                   index, size, type, normalized, relativeoffset);
        if (glVertexAttribFormat != nullptr) {
            glVertexAttribFormat(index, size, type, normalized, relativeoffset);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttribFormat");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttribFormat");
        return false;
    }
}

bool callGlVertexAttribI4i(Stack* stack, bool pushReturn) {
    int32_t w = stack->pop<int32_t>();
    int32_t z = stack->pop<int32_t>();
    int32_t y = stack->pop<int32_t>();
    int32_t x = stack->pop<int32_t>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttribI4i(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32
                   ")",
                   index, x, y, z, w);
        if (glVertexAttribI4i != nullptr) {
            glVertexAttribI4i(index, x, y, z, w);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttribI4i");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttribI4i");
        return false;
    }
}

bool callGlVertexAttribI4iv(Stack* stack, bool pushReturn) {
    int32_t* values = stack->pop<int32_t*>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttribI4iv(%" PRIu32 ", %p)", index, values);
        if (glVertexAttribI4iv != nullptr) {
            glVertexAttribI4iv(index, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttribI4iv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttribI4iv");
        return false;
    }
}

bool callGlVertexAttribI4ui(Stack* stack, bool pushReturn) {
    uint32_t w = stack->pop<uint32_t>();
    uint32_t z = stack->pop<uint32_t>();
    uint32_t y = stack->pop<uint32_t>();
    uint32_t x = stack->pop<uint32_t>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttribI4ui(%" PRIu32 ", %" PRIu32 ", %" PRIu32 ", %" PRIu32 ", %" PRIu32
                   ")",
                   index, x, y, z, w);
        if (glVertexAttribI4ui != nullptr) {
            glVertexAttribI4ui(index, x, y, z, w);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttribI4ui");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttribI4ui");
        return false;
    }
}

bool callGlVertexAttribI4uiv(Stack* stack, bool pushReturn) {
    uint32_t* values = stack->pop<uint32_t*>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttribI4uiv(%" PRIu32 ", %p)", index, values);
        if (glVertexAttribI4uiv != nullptr) {
            glVertexAttribI4uiv(index, values);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttribI4uiv");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttribI4uiv");
        return false;
    }
}

bool callGlVertexAttribIFormat(Stack* stack, bool pushReturn) {
    uint32_t relativeoffset = stack->pop<uint32_t>();
    GLenum type = stack->pop<GLenum>();
    int32_t size = stack->pop<int32_t>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttribIFormat(%" PRIu32 ", %" PRId32 ", %u, %" PRIu32 ")", index, size,
                   type, relativeoffset);
        if (glVertexAttribIFormat != nullptr) {
            glVertexAttribIFormat(index, size, type, relativeoffset);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttribIFormat");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttribIFormat");
        return false;
    }
}

bool callGlVertexAttribIPointer(Stack* stack, bool pushReturn) {
    void* data = stack->pop<void*>();
    int32_t stride = stack->pop<int32_t>();
    GLenum type = stack->pop<GLenum>();
    int32_t size = stack->pop<int32_t>();
    uint32_t location = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttribIPointer(%" PRIu32 ", %" PRId32 ", %u, %" PRId32 ", %p)",
                   location, size, type, stride, data);
        if (glVertexAttribIPointer != nullptr) {
            glVertexAttribIPointer(location, size, type, stride, data);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttribIPointer");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttribIPointer");
        return false;
    }
}

bool callGlVertexAttribPointer(Stack* stack, bool pushReturn) {
    void* data = stack->pop<void*>();
    int32_t stride = stack->pop<int32_t>();
    uint8_t normalized = stack->pop<uint8_t>();
    GLenum type = stack->pop<GLenum>();
    int32_t size = stack->pop<int32_t>();
    uint32_t location = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttribPointer(%" PRIu32 ", %" PRId32 ", %u, %" PRIu8 ", %" PRId32
                   ", %p)",
                   location, size, type, normalized, stride, data);
        if (glVertexAttribPointer != nullptr) {
            glVertexAttribPointer(location, size, type, normalized, stride, data);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttribPointer");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttribPointer");
        return false;
    }
}

bool callGlVertexBindingDivisor(Stack* stack, bool pushReturn) {
    uint32_t divisor = stack->pop<uint32_t>();
    uint32_t binding_index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexBindingDivisor(%" PRIu32 ", %" PRIu32 ")", binding_index, divisor);
        if (glVertexBindingDivisor != nullptr) {
            glVertexBindingDivisor(binding_index, divisor);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexBindingDivisor");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexBindingDivisor");
        return false;
    }
}

bool callEglInitialize(Stack* stack, bool pushReturn) {
    int* minor = stack->pop<int*>();
    int* major = stack->pop<int*>();
    void* dpy = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("eglInitialize(%p, %p, %p)", dpy, major, minor);
        if (eglInitialize != nullptr) {
            int return_value = eglInitialize(dpy, major, minor);
            GAPID_INFO("Returned: %d", return_value);
            if (pushReturn) {
                stack->push<int>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function eglInitialize");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function eglInitialize");
        return false;
    }
}

bool callEglCreateContext(Stack* stack, bool pushReturn) {
    int* attrib_list = stack->pop<int*>();
    void* share_context = stack->pop<void*>();
    void* config = stack->pop<void*>();
    void* display = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("eglCreateContext(%p, %p, %p, %p)", display, config, share_context, attrib_list);
        if (eglCreateContext != nullptr) {
            void* return_value = eglCreateContext(display, config, share_context, attrib_list);
            GAPID_INFO("Returned: %p", return_value);
            if (pushReturn) {
                stack->push<void*>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function eglCreateContext");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function eglCreateContext");
        return false;
    }
}

bool callEglMakeCurrent(Stack* stack, bool pushReturn) {
    void* context = stack->pop<void*>();
    void* read = stack->pop<void*>();
    void* draw = stack->pop<void*>();
    void* display = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("eglMakeCurrent(%p, %p, %p, %p)", display, draw, read, context);
        if (eglMakeCurrent != nullptr) {
            int return_value = eglMakeCurrent(display, draw, read, context);
            GAPID_INFO("Returned: %d", return_value);
            if (pushReturn) {
                stack->push<int>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function eglMakeCurrent");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function eglMakeCurrent");
        return false;
    }
}

bool callEglSwapBuffers(Stack* stack, bool pushReturn) {
    void* surface = stack->pop<void*>();
    void* display = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("eglSwapBuffers(%p, %p)", display, surface);
        if (eglSwapBuffers != nullptr) {
            int return_value = eglSwapBuffers(display, surface);
            GAPID_INFO("Returned: %d", return_value);
            if (pushReturn) {
                stack->push<int>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function eglSwapBuffers");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function eglSwapBuffers");
        return false;
    }
}

bool callEglQuerySurface(Stack* stack, bool pushReturn) {
    int* value = stack->pop<int*>();
    int attribute = stack->pop<int>();
    void* surface = stack->pop<void*>();
    void* display = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("eglQuerySurface(%p, %p, %d, %p)", display, surface, attribute, value);
        if (eglQuerySurface != nullptr) {
            int return_value = eglQuerySurface(display, surface, attribute, value);
            GAPID_INFO("Returned: %d", return_value);
            if (pushReturn) {
                stack->push<int>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function eglQuerySurface");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function eglQuerySurface");
        return false;
    }
}

bool callGlXCreateContext(Stack* stack, bool pushReturn) {
    bool direct = stack->pop<bool>();
    void* shareList = stack->pop<void*>();
    void* vis = stack->pop<void*>();
    void* dpy = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("glXCreateContext(%p, %p, %p, %d)", dpy, vis, shareList, direct);
        if (glXCreateContext != nullptr) {
            void* return_value = glXCreateContext(dpy, vis, shareList, direct);
            GAPID_INFO("Returned: %p", return_value);
            if (pushReturn) {
                stack->push<void*>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glXCreateContext");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glXCreateContext");
        return false;
    }
}

bool callGlXCreateNewContext(Stack* stack, bool pushReturn) {
    bool direct = stack->pop<bool>();
    void* shared = stack->pop<void*>();
    uint32_t type = stack->pop<uint32_t>();
    void* fbconfig = stack->pop<void*>();
    void* display = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("glXCreateNewContext(%p, %p, %" PRIu32 ", %p, %d)", display, fbconfig, type,
                   shared, direct);
        if (glXCreateNewContext != nullptr) {
            void* return_value = glXCreateNewContext(display, fbconfig, type, shared, direct);
            GAPID_INFO("Returned: %p", return_value);
            if (pushReturn) {
                stack->push<void*>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glXCreateNewContext");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glXCreateNewContext");
        return false;
    }
}

bool callGlXMakeContextCurrent(Stack* stack, bool pushReturn) {
    void* ctx = stack->pop<void*>();
    void* read = stack->pop<void*>();
    void* draw = stack->pop<void*>();
    void* display = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("glXMakeContextCurrent(%p, %p, %p, %p)", display, draw, read, ctx);
        if (glXMakeContextCurrent != nullptr) {
            int return_value = glXMakeContextCurrent(display, draw, read, ctx);
            GAPID_INFO("Returned: %d", return_value);
            if (pushReturn) {
                stack->push<int>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glXMakeContextCurrent");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glXMakeContextCurrent");
        return false;
    }
}

bool callGlXMakeCurrent(Stack* stack, bool pushReturn) {
    void* ctx = stack->pop<void*>();
    void* drawable = stack->pop<void*>();
    void* display = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("glXMakeCurrent(%p, %p, %p)", display, drawable, ctx);
        if (glXMakeCurrent != nullptr) {
            int return_value = glXMakeCurrent(display, drawable, ctx);
            GAPID_INFO("Returned: %d", return_value);
            if (pushReturn) {
                stack->push<int>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glXMakeCurrent");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glXMakeCurrent");
        return false;
    }
}

bool callGlXSwapBuffers(Stack* stack, bool pushReturn) {
    void* drawable = stack->pop<void*>();
    void* display = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("glXSwapBuffers(%p, %p)", display, drawable);
        if (glXSwapBuffers != nullptr) {
            glXSwapBuffers(display, drawable);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glXSwapBuffers");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glXSwapBuffers");
        return false;
    }
}

bool callGlXQueryDrawable(Stack* stack, bool pushReturn) {
    int* value = stack->pop<int*>();
    int attribute = stack->pop<int>();
    void* draw = stack->pop<void*>();
    void* display = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("glXQueryDrawable(%p, %p, %d, %p)", display, draw, attribute, value);
        if (glXQueryDrawable != nullptr) {
            int return_value = glXQueryDrawable(display, draw, attribute, value);
            GAPID_INFO("Returned: %d", return_value);
            if (pushReturn) {
                stack->push<int>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glXQueryDrawable");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glXQueryDrawable");
        return false;
    }
}

bool callWglCreateContext(Stack* stack, bool pushReturn) {
    void* hdc = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("wglCreateContext(%p)", hdc);
        if (wglCreateContext != nullptr) {
            void* return_value = wglCreateContext(hdc);
            GAPID_INFO("Returned: %p", return_value);
            if (pushReturn) {
                stack->push<void*>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function wglCreateContext");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function wglCreateContext");
        return false;
    }
}

bool callWglCreateContextAttribsARB(Stack* stack, bool pushReturn) {
    int* attribList = stack->pop<int*>();
    void* hShareContext = stack->pop<void*>();
    void* hdc = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("wglCreateContextAttribsARB(%p, %p, %p)", hdc, hShareContext, attribList);
        if (wglCreateContextAttribsARB != nullptr) {
            void* return_value = wglCreateContextAttribsARB(hdc, hShareContext, attribList);
            GAPID_INFO("Returned: %p", return_value);
            if (pushReturn) {
                stack->push<void*>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function wglCreateContextAttribsARB");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function wglCreateContextAttribsARB");
        return false;
    }
}

bool callWglMakeCurrent(Stack* stack, bool pushReturn) {
    void* hglrc = stack->pop<void*>();
    void* hdc = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("wglMakeCurrent(%p, %p)", hdc, hglrc);
        if (wglMakeCurrent != nullptr) {
            int return_value = wglMakeCurrent(hdc, hglrc);
            GAPID_INFO("Returned: %d", return_value);
            if (pushReturn) {
                stack->push<int>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function wglMakeCurrent");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function wglMakeCurrent");
        return false;
    }
}

bool callWglSwapBuffers(Stack* stack, bool pushReturn) {
    void* hdc = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("wglSwapBuffers(%p)", hdc);
        if (wglSwapBuffers != nullptr) {
            wglSwapBuffers(hdc);
        } else {
            GAPID_WARNING("Attempted to call unsupported function wglSwapBuffers");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function wglSwapBuffers");
        return false;
    }
}

bool callCGLCreateContext(Stack* stack, bool pushReturn) {
    void** ctx = stack->pop<void**>();
    void* share = stack->pop<void*>();
    void* pix = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("CGLCreateContext(%p, %p, %p)", pix, share, ctx);
        if (CGLCreateContext != nullptr) {
            int return_value = CGLCreateContext(pix, share, ctx);
            GAPID_INFO("Returned: %d", return_value);
            if (pushReturn) {
                stack->push<int>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function CGLCreateContext");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function CGLCreateContext");
        return false;
    }
}

bool callCGLSetCurrentContext(Stack* stack, bool pushReturn) {
    void* ctx = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("CGLSetCurrentContext(%p)", ctx);
        if (CGLSetCurrentContext != nullptr) {
            int return_value = CGLSetCurrentContext(ctx);
            GAPID_INFO("Returned: %d", return_value);
            if (pushReturn) {
                stack->push<int>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function CGLSetCurrentContext");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function CGLSetCurrentContext");
        return false;
    }
}

bool callCGLGetSurface(Stack* stack, bool pushReturn) {
    int32_t* sid = stack->pop<int32_t*>();
    int32_t* wid = stack->pop<int32_t*>();
    void** cid = stack->pop<void**>();
    void* ctx = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("CGLGetSurface(%p, %p, %p, %p)", ctx, cid, wid, sid);
        if (CGLGetSurface != nullptr) {
            int return_value = CGLGetSurface(ctx, cid, wid, sid);
            GAPID_INFO("Returned: %d", return_value);
            if (pushReturn) {
                stack->push<int>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function CGLGetSurface");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function CGLGetSurface");
        return false;
    }
}

bool callCGSGetSurfaceBounds(Stack* stack, bool pushReturn) {
    double* bounds = stack->pop<double*>();
    int32_t sid = stack->pop<int32_t>();
    int32_t wid = stack->pop<int32_t>();
    void* cid = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("CGSGetSurfaceBounds(%p, %" PRId32 ", %" PRId32 ", %p)", cid, wid, sid, bounds);
        if (CGSGetSurfaceBounds != nullptr) {
            int return_value = CGSGetSurfaceBounds(cid, wid, sid, bounds);
            GAPID_INFO("Returned: %d", return_value);
            if (pushReturn) {
                stack->push<int>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function CGSGetSurfaceBounds");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function CGSGetSurfaceBounds");
        return false;
    }
}

bool callCGLFlushDrawable(Stack* stack, bool pushReturn) {
    void* ctx = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("CGLFlushDrawable(%p)", ctx);
        if (CGLFlushDrawable != nullptr) {
            int return_value = CGLFlushDrawable(ctx);
            GAPID_INFO("Returned: %d", return_value);
            if (pushReturn) {
                stack->push<int>(return_value);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function CGLFlushDrawable");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function CGLFlushDrawable");
        return false;
    }
}

bool callGlGetQueryObjecti64v(Stack* stack, bool pushReturn) {
    int64_t* value = stack->pop<int64_t*>();
    GLenum parameter = stack->pop<GLenum>();
    uint32_t query = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetQueryObjecti64v(%" PRIu32 ", %u, %p)", query, parameter, value);
        if (glGetQueryObjecti64v != nullptr) {
            glGetQueryObjecti64v(query, parameter, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetQueryObjecti64v");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetQueryObjecti64v");
        return false;
    }
}

bool callGlGetQueryObjectui64v(Stack* stack, bool pushReturn) {
    uint64_t* value = stack->pop<uint64_t*>();
    GLenum parameter = stack->pop<GLenum>();
    uint32_t query = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetQueryObjectui64v(%" PRIu32 ", %u, %p)", query, parameter, value);
        if (glGetQueryObjectui64v != nullptr) {
            glGetQueryObjectui64v(query, parameter, value);
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetQueryObjectui64v");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetQueryObjectui64v");
        return false;
    }
}

}  // end of anonymous namespace

PFNGLBLENDBARRIERKHR glBlendBarrierKHR = nullptr;
PFNGLBLENDEQUATIONSEPARATEIEXT glBlendEquationSeparateiEXT = nullptr;
PFNGLBLENDEQUATIONIEXT glBlendEquationiEXT = nullptr;
PFNGLBLENDFUNCSEPARATEIEXT glBlendFuncSeparateiEXT = nullptr;
PFNGLBLENDFUNCIEXT glBlendFunciEXT = nullptr;
PFNGLCOLORMASKIEXT glColorMaskiEXT = nullptr;
PFNGLCOPYIMAGESUBDATAEXT glCopyImageSubDataEXT = nullptr;
PFNGLDEBUGMESSAGECALLBACKKHR glDebugMessageCallbackKHR = nullptr;
PFNGLDEBUGMESSAGECONTROLKHR glDebugMessageControlKHR = nullptr;
PFNGLDEBUGMESSAGEINSERTKHR glDebugMessageInsertKHR = nullptr;
PFNGLDISABLEIEXT glDisableiEXT = nullptr;
PFNGLENABLEIEXT glEnableiEXT = nullptr;
PFNGLFRAMEBUFFERTEXTUREEXT glFramebufferTextureEXT = nullptr;
PFNGLGETDEBUGMESSAGELOGKHR glGetDebugMessageLogKHR = nullptr;
PFNGLGETOBJECTLABELKHR glGetObjectLabelKHR = nullptr;
PFNGLGETOBJECTPTRLABELKHR glGetObjectPtrLabelKHR = nullptr;
PFNGLGETPOINTERVKHR glGetPointervKHR = nullptr;
PFNGLGETSAMPLERPARAMETERIIVEXT glGetSamplerParameterIivEXT = nullptr;
PFNGLGETSAMPLERPARAMETERIUIVEXT glGetSamplerParameterIuivEXT = nullptr;
PFNGLGETTEXPARAMETERIIVEXT glGetTexParameterIivEXT = nullptr;
PFNGLGETTEXPARAMETERIUIVEXT glGetTexParameterIuivEXT = nullptr;
PFNGLISENABLEDIEXT glIsEnablediEXT = nullptr;
PFNGLMINSAMPLESHADINGOES glMinSampleShadingOES = nullptr;
PFNGLOBJECTLABELKHR glObjectLabelKHR = nullptr;
PFNGLOBJECTPTRLABELKHR glObjectPtrLabelKHR = nullptr;
PFNGLPATCHPARAMETERIEXT glPatchParameteriEXT = nullptr;
PFNGLPOPDEBUGGROUPKHR glPopDebugGroupKHR = nullptr;
PFNGLPRIMITIVEBOUNDINGBOXEXT glPrimitiveBoundingBoxEXT = nullptr;
PFNGLPUSHDEBUGGROUPKHR glPushDebugGroupKHR = nullptr;
PFNGLSAMPLERPARAMETERIIVEXT glSamplerParameterIivEXT = nullptr;
PFNGLSAMPLERPARAMETERIUIVEXT glSamplerParameterIuivEXT = nullptr;
PFNGLTEXBUFFEREXT glTexBufferEXT = nullptr;
PFNGLTEXBUFFERRANGEEXT glTexBufferRangeEXT = nullptr;
PFNGLTEXPARAMETERIIVEXT glTexParameterIivEXT = nullptr;
PFNGLTEXPARAMETERIUIVEXT glTexParameterIuivEXT = nullptr;
PFNGLTEXSTORAGE3DMULTISAMPLEOES glTexStorage3DMultisampleOES = nullptr;
PFNGLBEGINQUERY glBeginQuery = nullptr;
PFNGLDELETEQUERIES glDeleteQueries = nullptr;
PFNGLENDQUERY glEndQuery = nullptr;
PFNGLGENQUERIES glGenQueries = nullptr;
PFNGLGETQUERYOBJECTUIV glGetQueryObjectuiv = nullptr;
PFNGLGETQUERYIV glGetQueryiv = nullptr;
PFNGLISQUERY glIsQuery = nullptr;
PFNGLBINDBUFFER glBindBuffer = nullptr;
PFNGLBINDBUFFERBASE glBindBufferBase = nullptr;
PFNGLBINDBUFFERRANGE glBindBufferRange = nullptr;
PFNGLBUFFERDATA glBufferData = nullptr;
PFNGLBUFFERSUBDATA glBufferSubData = nullptr;
PFNGLCOPYBUFFERSUBDATA glCopyBufferSubData = nullptr;
PFNGLDELETEBUFFERS glDeleteBuffers = nullptr;
PFNGLGENBUFFERS glGenBuffers = nullptr;
PFNGLGETBUFFERPARAMETERI64V glGetBufferParameteri64v = nullptr;
PFNGLGETBUFFERPARAMETERIV glGetBufferParameteriv = nullptr;
PFNGLGETBUFFERPOINTERV glGetBufferPointerv = nullptr;
PFNGLISBUFFER glIsBuffer = nullptr;
PFNGLMAPBUFFERRANGE glMapBufferRange = nullptr;
PFNGLUNMAPBUFFER glUnmapBuffer = nullptr;
PFNGLDEBUGMESSAGECALLBACK glDebugMessageCallback = nullptr;
PFNGLDEBUGMESSAGECONTROL glDebugMessageControl = nullptr;
PFNGLDEBUGMESSAGEINSERT glDebugMessageInsert = nullptr;
PFNGLGETDEBUGMESSAGELOG glGetDebugMessageLog = nullptr;
PFNGLGETOBJECTLABEL glGetObjectLabel = nullptr;
PFNGLGETOBJECTPTRLABEL glGetObjectPtrLabel = nullptr;
PFNGLGETPOINTERV glGetPointerv = nullptr;
PFNGLOBJECTLABEL glObjectLabel = nullptr;
PFNGLOBJECTPTRLABEL glObjectPtrLabel = nullptr;
PFNGLPOPDEBUGGROUP glPopDebugGroup = nullptr;
PFNGLPUSHDEBUGGROUP glPushDebugGroup = nullptr;
PFNGLDRAWARRAYS glDrawArrays = nullptr;
PFNGLDRAWARRAYSINDIRECT glDrawArraysIndirect = nullptr;
PFNGLDRAWARRAYSINSTANCED glDrawArraysInstanced = nullptr;
PFNGLDRAWELEMENTS glDrawElements = nullptr;
PFNGLDRAWELEMENTSBASEVERTEX glDrawElementsBaseVertex = nullptr;
PFNGLDRAWELEMENTSINDIRECT glDrawElementsIndirect = nullptr;
PFNGLDRAWELEMENTSINSTANCED glDrawElementsInstanced = nullptr;
PFNGLDRAWELEMENTSINSTANCEDBASEVERTEX glDrawElementsInstancedBaseVertex = nullptr;
PFNGLDRAWRANGEELEMENTS glDrawRangeElements = nullptr;
PFNGLDRAWRANGEELEMENTSBASEVERTEX glDrawRangeElementsBaseVertex = nullptr;
PFNGLPATCHPARAMETERI glPatchParameteri = nullptr;
PFNGLPRIMITIVEBOUNDINGBOX glPrimitiveBoundingBox = nullptr;
PFNGLACTIVESHADERPROGRAMEXT glActiveShaderProgramEXT = nullptr;
PFNGLALPHAFUNCQCOM glAlphaFuncQCOM = nullptr;
PFNGLAPPLYFRAMEBUFFERATTACHMENTCMAAINTEL glApplyFramebufferAttachmentCMAAINTEL = nullptr;
PFNGLBEGINCONDITIONALRENDERNV glBeginConditionalRenderNV = nullptr;
PFNGLBEGINPERFMONITORAMD glBeginPerfMonitorAMD = nullptr;
PFNGLBEGINPERFQUERYINTEL glBeginPerfQueryINTEL = nullptr;
PFNGLBEGINQUERYEXT glBeginQueryEXT = nullptr;
PFNGLBINDPROGRAMPIPELINEEXT glBindProgramPipelineEXT = nullptr;
PFNGLBINDVERTEXARRAYOES glBindVertexArrayOES = nullptr;
PFNGLBLENDBARRIERNV glBlendBarrierNV = nullptr;
PFNGLBLENDEQUATIONSEPARATEIOES glBlendEquationSeparateiOES = nullptr;
PFNGLBLENDEQUATIONIOES glBlendEquationiOES = nullptr;
PFNGLBLENDFUNCSEPARATEIOES glBlendFuncSeparateiOES = nullptr;
PFNGLBLENDFUNCIOES glBlendFunciOES = nullptr;
PFNGLBLENDPARAMETERINV glBlendParameteriNV = nullptr;
PFNGLBLITFRAMEBUFFERANGLE glBlitFramebufferANGLE = nullptr;
PFNGLBLITFRAMEBUFFERNV glBlitFramebufferNV = nullptr;
PFNGLBUFFERSTORAGEEXT glBufferStorageEXT = nullptr;
PFNGLCLIENTWAITSYNCAPPLE glClientWaitSyncAPPLE = nullptr;
PFNGLCOLORMASKIOES glColorMaskiOES = nullptr;
PFNGLCOMPRESSEDTEXIMAGE3DOES glCompressedTexImage3DOES = nullptr;
PFNGLCOMPRESSEDTEXSUBIMAGE3DOES glCompressedTexSubImage3DOES = nullptr;
PFNGLCOPYBUFFERSUBDATANV glCopyBufferSubDataNV = nullptr;
PFNGLCOPYIMAGESUBDATAOES glCopyImageSubDataOES = nullptr;
PFNGLCOPYPATHNV glCopyPathNV = nullptr;
PFNGLCOPYTEXSUBIMAGE3DOES glCopyTexSubImage3DOES = nullptr;
PFNGLCOPYTEXTURELEVELSAPPLE glCopyTextureLevelsAPPLE = nullptr;
PFNGLCOVERFILLPATHINSTANCEDNV glCoverFillPathInstancedNV = nullptr;
PFNGLCOVERFILLPATHNV glCoverFillPathNV = nullptr;
PFNGLCOVERSTROKEPATHINSTANCEDNV glCoverStrokePathInstancedNV = nullptr;
PFNGLCOVERSTROKEPATHNV glCoverStrokePathNV = nullptr;
PFNGLCOVERAGEMASKNV glCoverageMaskNV = nullptr;
PFNGLCOVERAGEMODULATIONNV glCoverageModulationNV = nullptr;
PFNGLCOVERAGEMODULATIONTABLENV glCoverageModulationTableNV = nullptr;
PFNGLCOVERAGEOPERATIONNV glCoverageOperationNV = nullptr;
PFNGLCREATEPERFQUERYINTEL glCreatePerfQueryINTEL = nullptr;
PFNGLCREATESHADERPROGRAMVEXT glCreateShaderProgramvEXT = nullptr;
PFNGLDELETEFENCESNV glDeleteFencesNV = nullptr;
PFNGLDELETEPATHSNV glDeletePathsNV = nullptr;
PFNGLDELETEPERFMONITORSAMD glDeletePerfMonitorsAMD = nullptr;
PFNGLDELETEPERFQUERYINTEL glDeletePerfQueryINTEL = nullptr;
PFNGLDELETEPROGRAMPIPELINESEXT glDeleteProgramPipelinesEXT = nullptr;
PFNGLDELETEQUERIESEXT glDeleteQueriesEXT = nullptr;
PFNGLDELETESYNCAPPLE glDeleteSyncAPPLE = nullptr;
PFNGLDELETEVERTEXARRAYSOES glDeleteVertexArraysOES = nullptr;
PFNGLDEPTHRANGEARRAYFVNV glDepthRangeArrayfvNV = nullptr;
PFNGLDEPTHRANGEINDEXEDFNV glDepthRangeIndexedfNV = nullptr;
PFNGLDISABLEDRIVERCONTROLQCOM glDisableDriverControlQCOM = nullptr;
PFNGLDISABLEINV glDisableiNV = nullptr;
PFNGLDISABLEIOES glDisableiOES = nullptr;
PFNGLDISCARDFRAMEBUFFEREXT glDiscardFramebufferEXT = nullptr;
PFNGLDRAWARRAYSINSTANCEDANGLE glDrawArraysInstancedANGLE = nullptr;
PFNGLDRAWARRAYSINSTANCEDBASEINSTANCEEXT glDrawArraysInstancedBaseInstanceEXT = nullptr;
PFNGLDRAWARRAYSINSTANCEDEXT glDrawArraysInstancedEXT = nullptr;
PFNGLDRAWARRAYSINSTANCEDNV glDrawArraysInstancedNV = nullptr;
PFNGLDRAWBUFFERSEXT glDrawBuffersEXT = nullptr;
PFNGLDRAWBUFFERSINDEXEDEXT glDrawBuffersIndexedEXT = nullptr;
PFNGLDRAWBUFFERSNV glDrawBuffersNV = nullptr;
PFNGLDRAWELEMENTSBASEVERTEXEXT glDrawElementsBaseVertexEXT = nullptr;
PFNGLDRAWELEMENTSBASEVERTEXOES glDrawElementsBaseVertexOES = nullptr;
PFNGLDRAWELEMENTSINSTANCEDANGLE glDrawElementsInstancedANGLE = nullptr;
PFNGLDRAWELEMENTSINSTANCEDBASEINSTANCEEXT glDrawElementsInstancedBaseInstanceEXT = nullptr;
PFNGLDRAWELEMENTSINSTANCEDBASEVERTEXBASEINSTANCEEXT
        glDrawElementsInstancedBaseVertexBaseInstanceEXT = nullptr;
PFNGLDRAWELEMENTSINSTANCEDBASEVERTEXEXT glDrawElementsInstancedBaseVertexEXT = nullptr;
PFNGLDRAWELEMENTSINSTANCEDBASEVERTEXOES glDrawElementsInstancedBaseVertexOES = nullptr;
PFNGLDRAWELEMENTSINSTANCEDEXT glDrawElementsInstancedEXT = nullptr;
PFNGLDRAWELEMENTSINSTANCEDNV glDrawElementsInstancedNV = nullptr;
PFNGLDRAWRANGEELEMENTSBASEVERTEXEXT glDrawRangeElementsBaseVertexEXT = nullptr;
PFNGLDRAWRANGEELEMENTSBASEVERTEXOES glDrawRangeElementsBaseVertexOES = nullptr;
PFNGLEGLIMAGETARGETRENDERBUFFERSTORAGEOES glEGLImageTargetRenderbufferStorageOES = nullptr;
PFNGLEGLIMAGETARGETTEXTURE2DOES glEGLImageTargetTexture2DOES = nullptr;
PFNGLENABLEDRIVERCONTROLQCOM glEnableDriverControlQCOM = nullptr;
PFNGLENABLEINV glEnableiNV = nullptr;
PFNGLENABLEIOES glEnableiOES = nullptr;
PFNGLENDCONDITIONALRENDERNV glEndConditionalRenderNV = nullptr;
PFNGLENDPERFMONITORAMD glEndPerfMonitorAMD = nullptr;
PFNGLENDPERFQUERYINTEL glEndPerfQueryINTEL = nullptr;
PFNGLENDQUERYEXT glEndQueryEXT = nullptr;
PFNGLENDTILINGQCOM glEndTilingQCOM = nullptr;
PFNGLEXTGETBUFFERPOINTERVQCOM glExtGetBufferPointervQCOM = nullptr;
PFNGLEXTGETBUFFERSQCOM glExtGetBuffersQCOM = nullptr;
PFNGLEXTGETFRAMEBUFFERSQCOM glExtGetFramebuffersQCOM = nullptr;
PFNGLEXTGETPROGRAMBINARYSOURCEQCOM glExtGetProgramBinarySourceQCOM = nullptr;
PFNGLEXTGETPROGRAMSQCOM glExtGetProgramsQCOM = nullptr;
PFNGLEXTGETRENDERBUFFERSQCOM glExtGetRenderbuffersQCOM = nullptr;
PFNGLEXTGETSHADERSQCOM glExtGetShadersQCOM = nullptr;
PFNGLEXTGETTEXLEVELPARAMETERIVQCOM glExtGetTexLevelParameterivQCOM = nullptr;
PFNGLEXTGETTEXSUBIMAGEQCOM glExtGetTexSubImageQCOM = nullptr;
PFNGLEXTGETTEXTURESQCOM glExtGetTexturesQCOM = nullptr;
PFNGLEXTISPROGRAMBINARYQCOM glExtIsProgramBinaryQCOM = nullptr;
PFNGLEXTTEXOBJECTSTATEOVERRIDEIQCOM glExtTexObjectStateOverrideiQCOM = nullptr;
PFNGLFENCESYNCAPPLE glFenceSyncAPPLE = nullptr;
PFNGLFINISHFENCENV glFinishFenceNV = nullptr;
PFNGLFLUSHMAPPEDBUFFERRANGEEXT glFlushMappedBufferRangeEXT = nullptr;
PFNGLFRAGMENTCOVERAGECOLORNV glFragmentCoverageColorNV = nullptr;
PFNGLFRAMEBUFFERSAMPLELOCATIONSFVNV glFramebufferSampleLocationsfvNV = nullptr;
PFNGLFRAMEBUFFERTEXTURE2DMULTISAMPLEEXT glFramebufferTexture2DMultisampleEXT = nullptr;
PFNGLFRAMEBUFFERTEXTURE2DMULTISAMPLEIMG glFramebufferTexture2DMultisampleIMG = nullptr;
PFNGLFRAMEBUFFERTEXTURE3DOES glFramebufferTexture3DOES = nullptr;
PFNGLFRAMEBUFFERTEXTUREMULTIVIEWOVR glFramebufferTextureMultiviewOVR = nullptr;
PFNGLFRAMEBUFFERTEXTUREOES glFramebufferTextureOES = nullptr;
PFNGLGENFENCESNV glGenFencesNV = nullptr;
PFNGLGENPATHSNV glGenPathsNV = nullptr;
PFNGLGENPERFMONITORSAMD glGenPerfMonitorsAMD = nullptr;
PFNGLGENPROGRAMPIPELINESEXT glGenProgramPipelinesEXT = nullptr;
PFNGLGENQUERIESEXT glGenQueriesEXT = nullptr;
PFNGLGENVERTEXARRAYSOES glGenVertexArraysOES = nullptr;
PFNGLGETBUFFERPOINTERVOES glGetBufferPointervOES = nullptr;
PFNGLGETCOVERAGEMODULATIONTABLENV glGetCoverageModulationTableNV = nullptr;
PFNGLGETDRIVERCONTROLSTRINGQCOM glGetDriverControlStringQCOM = nullptr;
PFNGLGETDRIVERCONTROLSQCOM glGetDriverControlsQCOM = nullptr;
PFNGLGETFENCEIVNV glGetFenceivNV = nullptr;
PFNGLGETFIRSTPERFQUERYIDINTEL glGetFirstPerfQueryIdINTEL = nullptr;
PFNGLGETFLOATI_VNV glGetFloati_vNV = nullptr;
PFNGLGETGRAPHICSRESETSTATUSEXT glGetGraphicsResetStatusEXT = nullptr;
PFNGLGETGRAPHICSRESETSTATUSKHR glGetGraphicsResetStatusKHR = nullptr;
PFNGLGETIMAGEHANDLENV glGetImageHandleNV = nullptr;
PFNGLGETINTEGER64VAPPLE glGetInteger64vAPPLE = nullptr;
PFNGLGETINTEGERI_VEXT glGetIntegeri_vEXT = nullptr;
PFNGLGETINTERNALFORMATSAMPLEIVNV glGetInternalformatSampleivNV = nullptr;
PFNGLGETNEXTPERFQUERYIDINTEL glGetNextPerfQueryIdINTEL = nullptr;
PFNGLGETOBJECTLABELEXT glGetObjectLabelEXT = nullptr;
PFNGLGETPATHCOMMANDSNV glGetPathCommandsNV = nullptr;
PFNGLGETPATHCOORDSNV glGetPathCoordsNV = nullptr;
PFNGLGETPATHDASHARRAYNV glGetPathDashArrayNV = nullptr;
PFNGLGETPATHLENGTHNV glGetPathLengthNV = nullptr;
PFNGLGETPATHMETRICRANGENV glGetPathMetricRangeNV = nullptr;
PFNGLGETPATHMETRICSNV glGetPathMetricsNV = nullptr;
PFNGLGETPATHPARAMETERFVNV glGetPathParameterfvNV = nullptr;
PFNGLGETPATHPARAMETERIVNV glGetPathParameterivNV = nullptr;
PFNGLGETPATHSPACINGNV glGetPathSpacingNV = nullptr;
PFNGLGETPERFCOUNTERINFOINTEL glGetPerfCounterInfoINTEL = nullptr;
PFNGLGETPERFMONITORCOUNTERDATAAMD glGetPerfMonitorCounterDataAMD = nullptr;
PFNGLGETPERFMONITORCOUNTERINFOAMD glGetPerfMonitorCounterInfoAMD = nullptr;
PFNGLGETPERFMONITORCOUNTERSTRINGAMD glGetPerfMonitorCounterStringAMD = nullptr;
PFNGLGETPERFMONITORCOUNTERSAMD glGetPerfMonitorCountersAMD = nullptr;
PFNGLGETPERFMONITORGROUPSTRINGAMD glGetPerfMonitorGroupStringAMD = nullptr;
PFNGLGETPERFMONITORGROUPSAMD glGetPerfMonitorGroupsAMD = nullptr;
PFNGLGETPERFQUERYDATAINTEL glGetPerfQueryDataINTEL = nullptr;
PFNGLGETPERFQUERYIDBYNAMEINTEL glGetPerfQueryIdByNameINTEL = nullptr;
PFNGLGETPERFQUERYINFOINTEL glGetPerfQueryInfoINTEL = nullptr;
PFNGLGETPROGRAMBINARYOES glGetProgramBinaryOES = nullptr;
PFNGLGETPROGRAMPIPELINEINFOLOGEXT glGetProgramPipelineInfoLogEXT = nullptr;
PFNGLGETPROGRAMPIPELINEIVEXT glGetProgramPipelineivEXT = nullptr;
PFNGLGETPROGRAMRESOURCEFVNV glGetProgramResourcefvNV = nullptr;
PFNGLGETQUERYOBJECTI64VEXT glGetQueryObjecti64vEXT = nullptr;
PFNGLGETQUERYOBJECTIVEXT glGetQueryObjectivEXT = nullptr;
PFNGLGETQUERYOBJECTUI64VEXT glGetQueryObjectui64vEXT = nullptr;
PFNGLGETQUERYOBJECTUIVEXT glGetQueryObjectuivEXT = nullptr;
PFNGLGETQUERYIVEXT glGetQueryivEXT = nullptr;
PFNGLGETSAMPLERPARAMETERIIVOES glGetSamplerParameterIivOES = nullptr;
PFNGLGETSAMPLERPARAMETERIUIVOES glGetSamplerParameterIuivOES = nullptr;
PFNGLGETSYNCIVAPPLE glGetSyncivAPPLE = nullptr;
PFNGLGETTEXPARAMETERIIVOES glGetTexParameterIivOES = nullptr;
PFNGLGETTEXPARAMETERIUIVOES glGetTexParameterIuivOES = nullptr;
PFNGLGETTEXTUREHANDLENV glGetTextureHandleNV = nullptr;
PFNGLGETTEXTURESAMPLERHANDLENV glGetTextureSamplerHandleNV = nullptr;
PFNGLGETTRANSLATEDSHADERSOURCEANGLE glGetTranslatedShaderSourceANGLE = nullptr;
PFNGLGETNUNIFORMFVEXT glGetnUniformfvEXT = nullptr;
PFNGLGETNUNIFORMFVKHR glGetnUniformfvKHR = nullptr;
PFNGLGETNUNIFORMIVEXT glGetnUniformivEXT = nullptr;
PFNGLGETNUNIFORMIVKHR glGetnUniformivKHR = nullptr;
PFNGLGETNUNIFORMUIVKHR glGetnUniformuivKHR = nullptr;
PFNGLINSERTEVENTMARKEREXT glInsertEventMarkerEXT = nullptr;
PFNGLINTERPOLATEPATHSNV glInterpolatePathsNV = nullptr;
PFNGLISENABLEDINV glIsEnablediNV = nullptr;
PFNGLISENABLEDIOES glIsEnablediOES = nullptr;
PFNGLISFENCENV glIsFenceNV = nullptr;
PFNGLISIMAGEHANDLERESIDENTNV glIsImageHandleResidentNV = nullptr;
PFNGLISPATHNV glIsPathNV = nullptr;
PFNGLISPOINTINFILLPATHNV glIsPointInFillPathNV = nullptr;
PFNGLISPOINTINSTROKEPATHNV glIsPointInStrokePathNV = nullptr;
PFNGLISPROGRAMPIPELINEEXT glIsProgramPipelineEXT = nullptr;
PFNGLISQUERYEXT glIsQueryEXT = nullptr;
PFNGLISSYNCAPPLE glIsSyncAPPLE = nullptr;
PFNGLISTEXTUREHANDLERESIDENTNV glIsTextureHandleResidentNV = nullptr;
PFNGLISVERTEXARRAYOES glIsVertexArrayOES = nullptr;
PFNGLLABELOBJECTEXT glLabelObjectEXT = nullptr;
PFNGLMAKEIMAGEHANDLENONRESIDENTNV glMakeImageHandleNonResidentNV = nullptr;
PFNGLMAKEIMAGEHANDLERESIDENTNV glMakeImageHandleResidentNV = nullptr;
PFNGLMAKETEXTUREHANDLENONRESIDENTNV glMakeTextureHandleNonResidentNV = nullptr;
PFNGLMAKETEXTUREHANDLERESIDENTNV glMakeTextureHandleResidentNV = nullptr;
PFNGLMAPBUFFEROES glMapBufferOES = nullptr;
PFNGLMAPBUFFERRANGEEXT glMapBufferRangeEXT = nullptr;
PFNGLMATRIXLOAD3X2FNV glMatrixLoad3x2fNV = nullptr;
PFNGLMATRIXLOAD3X3FNV glMatrixLoad3x3fNV = nullptr;
PFNGLMATRIXLOADTRANSPOSE3X3FNV glMatrixLoadTranspose3x3fNV = nullptr;
PFNGLMATRIXMULT3X2FNV glMatrixMult3x2fNV = nullptr;
PFNGLMATRIXMULT3X3FNV glMatrixMult3x3fNV = nullptr;
PFNGLMATRIXMULTTRANSPOSE3X3FNV glMatrixMultTranspose3x3fNV = nullptr;
PFNGLMULTIDRAWARRAYSEXT glMultiDrawArraysEXT = nullptr;
PFNGLMULTIDRAWARRAYSINDIRECTEXT glMultiDrawArraysIndirectEXT = nullptr;
PFNGLMULTIDRAWELEMENTSBASEVERTEXEXT glMultiDrawElementsBaseVertexEXT = nullptr;
PFNGLMULTIDRAWELEMENTSBASEVERTEXOES glMultiDrawElementsBaseVertexOES = nullptr;
PFNGLMULTIDRAWELEMENTSEXT glMultiDrawElementsEXT = nullptr;
PFNGLMULTIDRAWELEMENTSINDIRECTEXT glMultiDrawElementsIndirectEXT = nullptr;
PFNGLNAMEDFRAMEBUFFERSAMPLELOCATIONSFVNV glNamedFramebufferSampleLocationsfvNV = nullptr;
PFNGLPATCHPARAMETERIOES glPatchParameteriOES = nullptr;
PFNGLPATHCOMMANDSNV glPathCommandsNV = nullptr;
PFNGLPATHCOORDSNV glPathCoordsNV = nullptr;
PFNGLPATHCOVERDEPTHFUNCNV glPathCoverDepthFuncNV = nullptr;
PFNGLPATHDASHARRAYNV glPathDashArrayNV = nullptr;
PFNGLPATHGLYPHINDEXARRAYNV glPathGlyphIndexArrayNV = nullptr;
PFNGLPATHGLYPHINDEXRANGENV glPathGlyphIndexRangeNV = nullptr;
PFNGLPATHGLYPHRANGENV glPathGlyphRangeNV = nullptr;
PFNGLPATHGLYPHSNV glPathGlyphsNV = nullptr;
PFNGLPATHMEMORYGLYPHINDEXARRAYNV glPathMemoryGlyphIndexArrayNV = nullptr;
PFNGLPATHPARAMETERFNV glPathParameterfNV = nullptr;
PFNGLPATHPARAMETERFVNV glPathParameterfvNV = nullptr;
PFNGLPATHPARAMETERINV glPathParameteriNV = nullptr;
PFNGLPATHPARAMETERIVNV glPathParameterivNV = nullptr;
PFNGLPATHSTENCILDEPTHOFFSETNV glPathStencilDepthOffsetNV = nullptr;
PFNGLPATHSTENCILFUNCNV glPathStencilFuncNV = nullptr;
PFNGLPATHSTRINGNV glPathStringNV = nullptr;
PFNGLPATHSUBCOMMANDSNV glPathSubCommandsNV = nullptr;
PFNGLPATHSUBCOORDSNV glPathSubCoordsNV = nullptr;
PFNGLPOINTALONGPATHNV glPointAlongPathNV = nullptr;
PFNGLPOLYGONMODENV glPolygonModeNV = nullptr;
PFNGLPOPGROUPMARKEREXT glPopGroupMarkerEXT = nullptr;
PFNGLPRIMITIVEBOUNDINGBOXOES glPrimitiveBoundingBoxOES = nullptr;
PFNGLPROGRAMBINARYOES glProgramBinaryOES = nullptr;
PFNGLPROGRAMPARAMETERIEXT glProgramParameteriEXT = nullptr;
PFNGLPROGRAMPATHFRAGMENTINPUTGENNV glProgramPathFragmentInputGenNV = nullptr;
PFNGLPROGRAMUNIFORM1FEXT glProgramUniform1fEXT = nullptr;
PFNGLPROGRAMUNIFORM1FVEXT glProgramUniform1fvEXT = nullptr;
PFNGLPROGRAMUNIFORM1IEXT glProgramUniform1iEXT = nullptr;
PFNGLPROGRAMUNIFORM1IVEXT glProgramUniform1ivEXT = nullptr;
PFNGLPROGRAMUNIFORM1UIEXT glProgramUniform1uiEXT = nullptr;
PFNGLPROGRAMUNIFORM1UIVEXT glProgramUniform1uivEXT = nullptr;
PFNGLPROGRAMUNIFORM2FEXT glProgramUniform2fEXT = nullptr;
PFNGLPROGRAMUNIFORM2FVEXT glProgramUniform2fvEXT = nullptr;
PFNGLPROGRAMUNIFORM2IEXT glProgramUniform2iEXT = nullptr;
PFNGLPROGRAMUNIFORM2IVEXT glProgramUniform2ivEXT = nullptr;
PFNGLPROGRAMUNIFORM2UIEXT glProgramUniform2uiEXT = nullptr;
PFNGLPROGRAMUNIFORM2UIVEXT glProgramUniform2uivEXT = nullptr;
PFNGLPROGRAMUNIFORM3FEXT glProgramUniform3fEXT = nullptr;
PFNGLPROGRAMUNIFORM3FVEXT glProgramUniform3fvEXT = nullptr;
PFNGLPROGRAMUNIFORM3IEXT glProgramUniform3iEXT = nullptr;
PFNGLPROGRAMUNIFORM3IVEXT glProgramUniform3ivEXT = nullptr;
PFNGLPROGRAMUNIFORM3UIEXT glProgramUniform3uiEXT = nullptr;
PFNGLPROGRAMUNIFORM3UIVEXT glProgramUniform3uivEXT = nullptr;
PFNGLPROGRAMUNIFORM4FEXT glProgramUniform4fEXT = nullptr;
PFNGLPROGRAMUNIFORM4FVEXT glProgramUniform4fvEXT = nullptr;
PFNGLPROGRAMUNIFORM4IEXT glProgramUniform4iEXT = nullptr;
PFNGLPROGRAMUNIFORM4IVEXT glProgramUniform4ivEXT = nullptr;
PFNGLPROGRAMUNIFORM4UIEXT glProgramUniform4uiEXT = nullptr;
PFNGLPROGRAMUNIFORM4UIVEXT glProgramUniform4uivEXT = nullptr;
PFNGLPROGRAMUNIFORMHANDLEUI64NV glProgramUniformHandleui64NV = nullptr;
PFNGLPROGRAMUNIFORMHANDLEUI64VNV glProgramUniformHandleui64vNV = nullptr;
PFNGLPROGRAMUNIFORMMATRIX2FVEXT glProgramUniformMatrix2fvEXT = nullptr;
PFNGLPROGRAMUNIFORMMATRIX2X3FVEXT glProgramUniformMatrix2x3fvEXT = nullptr;
PFNGLPROGRAMUNIFORMMATRIX2X4FVEXT glProgramUniformMatrix2x4fvEXT = nullptr;
PFNGLPROGRAMUNIFORMMATRIX3FVEXT glProgramUniformMatrix3fvEXT = nullptr;
PFNGLPROGRAMUNIFORMMATRIX3X2FVEXT glProgramUniformMatrix3x2fvEXT = nullptr;
PFNGLPROGRAMUNIFORMMATRIX3X4FVEXT glProgramUniformMatrix3x4fvEXT = nullptr;
PFNGLPROGRAMUNIFORMMATRIX4FVEXT glProgramUniformMatrix4fvEXT = nullptr;
PFNGLPROGRAMUNIFORMMATRIX4X2FVEXT glProgramUniformMatrix4x2fvEXT = nullptr;
PFNGLPROGRAMUNIFORMMATRIX4X3FVEXT glProgramUniformMatrix4x3fvEXT = nullptr;
PFNGLPUSHGROUPMARKEREXT glPushGroupMarkerEXT = nullptr;
PFNGLQUERYCOUNTEREXT glQueryCounterEXT = nullptr;
PFNGLRASTERSAMPLESEXT glRasterSamplesEXT = nullptr;
PFNGLREADBUFFERINDEXEDEXT glReadBufferIndexedEXT = nullptr;
PFNGLREADBUFFERNV glReadBufferNV = nullptr;
PFNGLREADNPIXELSEXT glReadnPixelsEXT = nullptr;
PFNGLREADNPIXELSKHR glReadnPixelsKHR = nullptr;
PFNGLRENDERBUFFERSTORAGEMULTISAMPLEANGLE glRenderbufferStorageMultisampleANGLE = nullptr;
PFNGLRENDERBUFFERSTORAGEMULTISAMPLEAPPLE glRenderbufferStorageMultisampleAPPLE = nullptr;
PFNGLRENDERBUFFERSTORAGEMULTISAMPLEEXT glRenderbufferStorageMultisampleEXT = nullptr;
PFNGLRENDERBUFFERSTORAGEMULTISAMPLEIMG glRenderbufferStorageMultisampleIMG = nullptr;
PFNGLRENDERBUFFERSTORAGEMULTISAMPLENV glRenderbufferStorageMultisampleNV = nullptr;
PFNGLRESOLVEDEPTHVALUESNV glResolveDepthValuesNV = nullptr;
PFNGLRESOLVEMULTISAMPLEFRAMEBUFFERAPPLE glResolveMultisampleFramebufferAPPLE = nullptr;
PFNGLSAMPLERPARAMETERIIVOES glSamplerParameterIivOES = nullptr;
PFNGLSAMPLERPARAMETERIUIVOES glSamplerParameterIuivOES = nullptr;
PFNGLSCISSORARRAYVNV glScissorArrayvNV = nullptr;
PFNGLSCISSORINDEXEDNV glScissorIndexedNV = nullptr;
PFNGLSCISSORINDEXEDVNV glScissorIndexedvNV = nullptr;
PFNGLSELECTPERFMONITORCOUNTERSAMD glSelectPerfMonitorCountersAMD = nullptr;
PFNGLSETFENCENV glSetFenceNV = nullptr;
PFNGLSTARTTILINGQCOM glStartTilingQCOM = nullptr;
PFNGLSTENCILFILLPATHINSTANCEDNV glStencilFillPathInstancedNV = nullptr;
PFNGLSTENCILFILLPATHNV glStencilFillPathNV = nullptr;
PFNGLSTENCILSTROKEPATHINSTANCEDNV glStencilStrokePathInstancedNV = nullptr;
PFNGLSTENCILSTROKEPATHNV glStencilStrokePathNV = nullptr;
PFNGLSTENCILTHENCOVERFILLPATHINSTANCEDNV glStencilThenCoverFillPathInstancedNV = nullptr;
PFNGLSTENCILTHENCOVERFILLPATHNV glStencilThenCoverFillPathNV = nullptr;
PFNGLSTENCILTHENCOVERSTROKEPATHINSTANCEDNV glStencilThenCoverStrokePathInstancedNV = nullptr;
PFNGLSTENCILTHENCOVERSTROKEPATHNV glStencilThenCoverStrokePathNV = nullptr;
PFNGLSUBPIXELPRECISIONBIASNV glSubpixelPrecisionBiasNV = nullptr;
PFNGLTESTFENCENV glTestFenceNV = nullptr;
PFNGLTEXBUFFEROES glTexBufferOES = nullptr;
PFNGLTEXBUFFERRANGEOES glTexBufferRangeOES = nullptr;
PFNGLTEXIMAGE3DOES glTexImage3DOES = nullptr;
PFNGLTEXPAGECOMMITMENTARB glTexPageCommitmentARB = nullptr;
PFNGLTEXPARAMETERIIVOES glTexParameterIivOES = nullptr;
PFNGLTEXPARAMETERIUIVOES glTexParameterIuivOES = nullptr;
PFNGLTEXSTORAGE1DEXT glTexStorage1DEXT = nullptr;
PFNGLTEXSTORAGE2DEXT glTexStorage2DEXT = nullptr;
PFNGLTEXSTORAGE3DEXT glTexStorage3DEXT = nullptr;
PFNGLTEXSUBIMAGE3DOES glTexSubImage3DOES = nullptr;
PFNGLTEXTURESTORAGE1DEXT glTextureStorage1DEXT = nullptr;
PFNGLTEXTURESTORAGE2DEXT glTextureStorage2DEXT = nullptr;
PFNGLTEXTURESTORAGE3DEXT glTextureStorage3DEXT = nullptr;
PFNGLTEXTUREVIEWEXT glTextureViewEXT = nullptr;
PFNGLTEXTUREVIEWOES glTextureViewOES = nullptr;
PFNGLTRANSFORMPATHNV glTransformPathNV = nullptr;
PFNGLUNIFORMHANDLEUI64NV glUniformHandleui64NV = nullptr;
PFNGLUNIFORMHANDLEUI64VNV glUniformHandleui64vNV = nullptr;
PFNGLUNIFORMMATRIX2X3FVNV glUniformMatrix2x3fvNV = nullptr;
PFNGLUNIFORMMATRIX2X4FVNV glUniformMatrix2x4fvNV = nullptr;
PFNGLUNIFORMMATRIX3X2FVNV glUniformMatrix3x2fvNV = nullptr;
PFNGLUNIFORMMATRIX3X4FVNV glUniformMatrix3x4fvNV = nullptr;
PFNGLUNIFORMMATRIX4X2FVNV glUniformMatrix4x2fvNV = nullptr;
PFNGLUNIFORMMATRIX4X3FVNV glUniformMatrix4x3fvNV = nullptr;
PFNGLUNMAPBUFFEROES glUnmapBufferOES = nullptr;
PFNGLUSEPROGRAMSTAGESEXT glUseProgramStagesEXT = nullptr;
PFNGLVALIDATEPROGRAMPIPELINEEXT glValidateProgramPipelineEXT = nullptr;
PFNGLVERTEXATTRIBDIVISORANGLE glVertexAttribDivisorANGLE = nullptr;
PFNGLVERTEXATTRIBDIVISOREXT glVertexAttribDivisorEXT = nullptr;
PFNGLVERTEXATTRIBDIVISORNV glVertexAttribDivisorNV = nullptr;
PFNGLVIEWPORTARRAYVNV glViewportArrayvNV = nullptr;
PFNGLVIEWPORTINDEXEDFNV glViewportIndexedfNV = nullptr;
PFNGLVIEWPORTINDEXEDFVNV glViewportIndexedfvNV = nullptr;
PFNGLWAITSYNCAPPLE glWaitSyncAPPLE = nullptr;
PFNGLWEIGHTPATHSNV glWeightPathsNV = nullptr;
PFNGLBLENDBARRIER glBlendBarrier = nullptr;
PFNGLBLENDCOLOR glBlendColor = nullptr;
PFNGLBLENDEQUATION glBlendEquation = nullptr;
PFNGLBLENDEQUATIONSEPARATE glBlendEquationSeparate = nullptr;
PFNGLBLENDEQUATIONSEPARATEI glBlendEquationSeparatei = nullptr;
PFNGLBLENDEQUATIONI glBlendEquationi = nullptr;
PFNGLBLENDFUNC glBlendFunc = nullptr;
PFNGLBLENDFUNCSEPARATE glBlendFuncSeparate = nullptr;
PFNGLBLENDFUNCSEPARATEI glBlendFuncSeparatei = nullptr;
PFNGLBLENDFUNCI glBlendFunci = nullptr;
PFNGLDEPTHFUNC glDepthFunc = nullptr;
PFNGLSAMPLECOVERAGE glSampleCoverage = nullptr;
PFNGLSAMPLEMASKI glSampleMaski = nullptr;
PFNGLSCISSOR glScissor = nullptr;
PFNGLSTENCILFUNC glStencilFunc = nullptr;
PFNGLSTENCILFUNCSEPARATE glStencilFuncSeparate = nullptr;
PFNGLSTENCILOP glStencilOp = nullptr;
PFNGLSTENCILOPSEPARATE glStencilOpSeparate = nullptr;
PFNGLBINDFRAMEBUFFER glBindFramebuffer = nullptr;
PFNGLBINDRENDERBUFFER glBindRenderbuffer = nullptr;
PFNGLBLITFRAMEBUFFER glBlitFramebuffer = nullptr;
PFNGLCHECKFRAMEBUFFERSTATUS glCheckFramebufferStatus = nullptr;
PFNGLCLEAR glClear = nullptr;
PFNGLCLEARBUFFERFI glClearBufferfi = nullptr;
PFNGLCLEARBUFFERFV glClearBufferfv = nullptr;
PFNGLCLEARBUFFERIV glClearBufferiv = nullptr;
PFNGLCLEARBUFFERUIV glClearBufferuiv = nullptr;
PFNGLCLEARCOLOR glClearColor = nullptr;
PFNGLCLEARDEPTHF glClearDepthf = nullptr;
PFNGLCLEARSTENCIL glClearStencil = nullptr;
PFNGLCOLORMASK glColorMask = nullptr;
PFNGLCOLORMASKI glColorMaski = nullptr;
PFNGLDELETEFRAMEBUFFERS glDeleteFramebuffers = nullptr;
PFNGLDELETERENDERBUFFERS glDeleteRenderbuffers = nullptr;
PFNGLDEPTHMASK glDepthMask = nullptr;
PFNGLDRAWBUFFERS glDrawBuffers = nullptr;
PFNGLFRAMEBUFFERPARAMETERI glFramebufferParameteri = nullptr;
PFNGLFRAMEBUFFERRENDERBUFFER glFramebufferRenderbuffer = nullptr;
PFNGLFRAMEBUFFERTEXTURE glFramebufferTexture = nullptr;
PFNGLFRAMEBUFFERTEXTURE2D glFramebufferTexture2D = nullptr;
PFNGLFRAMEBUFFERTEXTURELAYER glFramebufferTextureLayer = nullptr;
PFNGLGENFRAMEBUFFERS glGenFramebuffers = nullptr;
PFNGLGENRENDERBUFFERS glGenRenderbuffers = nullptr;
PFNGLGETFRAMEBUFFERATTACHMENTPARAMETERIV glGetFramebufferAttachmentParameteriv = nullptr;
PFNGLGETFRAMEBUFFERPARAMETERIV glGetFramebufferParameteriv = nullptr;
PFNGLGETRENDERBUFFERPARAMETERIV glGetRenderbufferParameteriv = nullptr;
PFNGLINVALIDATEFRAMEBUFFER glInvalidateFramebuffer = nullptr;
PFNGLINVALIDATESUBFRAMEBUFFER glInvalidateSubFramebuffer = nullptr;
PFNGLISFRAMEBUFFER glIsFramebuffer = nullptr;
PFNGLISRENDERBUFFER glIsRenderbuffer = nullptr;
PFNGLREADBUFFER glReadBuffer = nullptr;
PFNGLREADPIXELS glReadPixels = nullptr;
PFNGLREADNPIXELS glReadnPixels = nullptr;
PFNGLRENDERBUFFERSTORAGE glRenderbufferStorage = nullptr;
PFNGLRENDERBUFFERSTORAGEMULTISAMPLE glRenderbufferStorageMultisample = nullptr;
PFNGLSTENCILMASK glStencilMask = nullptr;
PFNGLSTENCILMASKSEPARATE glStencilMaskSeparate = nullptr;
PFNGLDISABLE glDisable = nullptr;
PFNGLDISABLEI glDisablei = nullptr;
PFNGLENABLE glEnable = nullptr;
PFNGLENABLEI glEnablei = nullptr;
PFNGLFINISH glFinish = nullptr;
PFNGLFLUSH glFlush = nullptr;
PFNGLFLUSHMAPPEDBUFFERRANGE glFlushMappedBufferRange = nullptr;
PFNGLGETERROR glGetError = nullptr;
PFNGLGETGRAPHICSRESETSTATUS glGetGraphicsResetStatus = nullptr;
PFNGLHINT glHint = nullptr;
PFNGLACTIVESHADERPROGRAM glActiveShaderProgram = nullptr;
PFNGLATTACHSHADER glAttachShader = nullptr;
PFNGLBINDATTRIBLOCATION glBindAttribLocation = nullptr;
PFNGLBINDPROGRAMPIPELINE glBindProgramPipeline = nullptr;
PFNGLCOMPILESHADER glCompileShader = nullptr;
PFNGLCREATEPROGRAM glCreateProgram = nullptr;
PFNGLCREATESHADER glCreateShader = nullptr;
PFNGLCREATESHADERPROGRAMV glCreateShaderProgramv = nullptr;
PFNGLDELETEPROGRAM glDeleteProgram = nullptr;
PFNGLDELETEPROGRAMPIPELINES glDeleteProgramPipelines = nullptr;
PFNGLDELETESHADER glDeleteShader = nullptr;
PFNGLDETACHSHADER glDetachShader = nullptr;
PFNGLDISPATCHCOMPUTE glDispatchCompute = nullptr;
PFNGLDISPATCHCOMPUTEINDIRECT glDispatchComputeIndirect = nullptr;
PFNGLGENPROGRAMPIPELINES glGenProgramPipelines = nullptr;
PFNGLGETACTIVEATTRIB glGetActiveAttrib = nullptr;
PFNGLGETACTIVEUNIFORM glGetActiveUniform = nullptr;
PFNGLGETACTIVEUNIFORMBLOCKNAME glGetActiveUniformBlockName = nullptr;
PFNGLGETACTIVEUNIFORMBLOCKIV glGetActiveUniformBlockiv = nullptr;
PFNGLGETACTIVEUNIFORMSIV glGetActiveUniformsiv = nullptr;
PFNGLGETATTACHEDSHADERS glGetAttachedShaders = nullptr;
PFNGLGETATTRIBLOCATION glGetAttribLocation = nullptr;
PFNGLGETFRAGDATALOCATION glGetFragDataLocation = nullptr;
PFNGLGETPROGRAMBINARY glGetProgramBinary = nullptr;
PFNGLGETPROGRAMINFOLOG glGetProgramInfoLog = nullptr;
PFNGLGETPROGRAMINTERFACEIV glGetProgramInterfaceiv = nullptr;
PFNGLGETPROGRAMPIPELINEINFOLOG glGetProgramPipelineInfoLog = nullptr;
PFNGLGETPROGRAMPIPELINEIV glGetProgramPipelineiv = nullptr;
PFNGLGETPROGRAMRESOURCEINDEX glGetProgramResourceIndex = nullptr;
PFNGLGETPROGRAMRESOURCELOCATION glGetProgramResourceLocation = nullptr;
PFNGLGETPROGRAMRESOURCENAME glGetProgramResourceName = nullptr;
PFNGLGETPROGRAMRESOURCEIV glGetProgramResourceiv = nullptr;
PFNGLGETPROGRAMIV glGetProgramiv = nullptr;
PFNGLGETSHADERINFOLOG glGetShaderInfoLog = nullptr;
PFNGLGETSHADERPRECISIONFORMAT glGetShaderPrecisionFormat = nullptr;
PFNGLGETSHADERSOURCE glGetShaderSource = nullptr;
PFNGLGETSHADERIV glGetShaderiv = nullptr;
PFNGLGETUNIFORMBLOCKINDEX glGetUniformBlockIndex = nullptr;
PFNGLGETUNIFORMINDICES glGetUniformIndices = nullptr;
PFNGLGETUNIFORMLOCATION glGetUniformLocation = nullptr;
PFNGLGETUNIFORMFV glGetUniformfv = nullptr;
PFNGLGETUNIFORMIV glGetUniformiv = nullptr;
PFNGLGETUNIFORMUIV glGetUniformuiv = nullptr;
PFNGLGETNUNIFORMFV glGetnUniformfv = nullptr;
PFNGLGETNUNIFORMIV glGetnUniformiv = nullptr;
PFNGLGETNUNIFORMUIV glGetnUniformuiv = nullptr;
PFNGLISPROGRAM glIsProgram = nullptr;
PFNGLISPROGRAMPIPELINE glIsProgramPipeline = nullptr;
PFNGLISSHADER glIsShader = nullptr;
PFNGLLINKPROGRAM glLinkProgram = nullptr;
PFNGLMEMORYBARRIER glMemoryBarrier = nullptr;
PFNGLMEMORYBARRIERBYREGION glMemoryBarrierByRegion = nullptr;
PFNGLPROGRAMBINARY glProgramBinary = nullptr;
PFNGLPROGRAMPARAMETERI glProgramParameteri = nullptr;
PFNGLPROGRAMUNIFORM1F glProgramUniform1f = nullptr;
PFNGLPROGRAMUNIFORM1FV glProgramUniform1fv = nullptr;
PFNGLPROGRAMUNIFORM1I glProgramUniform1i = nullptr;
PFNGLPROGRAMUNIFORM1IV glProgramUniform1iv = nullptr;
PFNGLPROGRAMUNIFORM1UI glProgramUniform1ui = nullptr;
PFNGLPROGRAMUNIFORM1UIV glProgramUniform1uiv = nullptr;
PFNGLPROGRAMUNIFORM2F glProgramUniform2f = nullptr;
PFNGLPROGRAMUNIFORM2FV glProgramUniform2fv = nullptr;
PFNGLPROGRAMUNIFORM2I glProgramUniform2i = nullptr;
PFNGLPROGRAMUNIFORM2IV glProgramUniform2iv = nullptr;
PFNGLPROGRAMUNIFORM2UI glProgramUniform2ui = nullptr;
PFNGLPROGRAMUNIFORM2UIV glProgramUniform2uiv = nullptr;
PFNGLPROGRAMUNIFORM3F glProgramUniform3f = nullptr;
PFNGLPROGRAMUNIFORM3FV glProgramUniform3fv = nullptr;
PFNGLPROGRAMUNIFORM3I glProgramUniform3i = nullptr;
PFNGLPROGRAMUNIFORM3IV glProgramUniform3iv = nullptr;
PFNGLPROGRAMUNIFORM3UI glProgramUniform3ui = nullptr;
PFNGLPROGRAMUNIFORM3UIV glProgramUniform3uiv = nullptr;
PFNGLPROGRAMUNIFORM4F glProgramUniform4f = nullptr;
PFNGLPROGRAMUNIFORM4FV glProgramUniform4fv = nullptr;
PFNGLPROGRAMUNIFORM4I glProgramUniform4i = nullptr;
PFNGLPROGRAMUNIFORM4IV glProgramUniform4iv = nullptr;
PFNGLPROGRAMUNIFORM4UI glProgramUniform4ui = nullptr;
PFNGLPROGRAMUNIFORM4UIV glProgramUniform4uiv = nullptr;
PFNGLPROGRAMUNIFORMMATRIX2FV glProgramUniformMatrix2fv = nullptr;
PFNGLPROGRAMUNIFORMMATRIX2X3FV glProgramUniformMatrix2x3fv = nullptr;
PFNGLPROGRAMUNIFORMMATRIX2X4FV glProgramUniformMatrix2x4fv = nullptr;
PFNGLPROGRAMUNIFORMMATRIX3FV glProgramUniformMatrix3fv = nullptr;
PFNGLPROGRAMUNIFORMMATRIX3X2FV glProgramUniformMatrix3x2fv = nullptr;
PFNGLPROGRAMUNIFORMMATRIX3X4FV glProgramUniformMatrix3x4fv = nullptr;
PFNGLPROGRAMUNIFORMMATRIX4FV glProgramUniformMatrix4fv = nullptr;
PFNGLPROGRAMUNIFORMMATRIX4X2FV glProgramUniformMatrix4x2fv = nullptr;
PFNGLPROGRAMUNIFORMMATRIX4X3FV glProgramUniformMatrix4x3fv = nullptr;
PFNGLRELEASESHADERCOMPILER glReleaseShaderCompiler = nullptr;
PFNGLSHADERBINARY glShaderBinary = nullptr;
PFNGLSHADERSOURCE glShaderSource = nullptr;
PFNGLUNIFORM1F glUniform1f = nullptr;
PFNGLUNIFORM1FV glUniform1fv = nullptr;
PFNGLUNIFORM1I glUniform1i = nullptr;
PFNGLUNIFORM1IV glUniform1iv = nullptr;
PFNGLUNIFORM1UI glUniform1ui = nullptr;
PFNGLUNIFORM1UIV glUniform1uiv = nullptr;
PFNGLUNIFORM2F glUniform2f = nullptr;
PFNGLUNIFORM2FV glUniform2fv = nullptr;
PFNGLUNIFORM2I glUniform2i = nullptr;
PFNGLUNIFORM2IV glUniform2iv = nullptr;
PFNGLUNIFORM2UI glUniform2ui = nullptr;
PFNGLUNIFORM2UIV glUniform2uiv = nullptr;
PFNGLUNIFORM3F glUniform3f = nullptr;
PFNGLUNIFORM3FV glUniform3fv = nullptr;
PFNGLUNIFORM3I glUniform3i = nullptr;
PFNGLUNIFORM3IV glUniform3iv = nullptr;
PFNGLUNIFORM3UI glUniform3ui = nullptr;
PFNGLUNIFORM3UIV glUniform3uiv = nullptr;
PFNGLUNIFORM4F glUniform4f = nullptr;
PFNGLUNIFORM4FV glUniform4fv = nullptr;
PFNGLUNIFORM4I glUniform4i = nullptr;
PFNGLUNIFORM4IV glUniform4iv = nullptr;
PFNGLUNIFORM4UI glUniform4ui = nullptr;
PFNGLUNIFORM4UIV glUniform4uiv = nullptr;
PFNGLUNIFORMBLOCKBINDING glUniformBlockBinding = nullptr;
PFNGLUNIFORMMATRIX2FV glUniformMatrix2fv = nullptr;
PFNGLUNIFORMMATRIX2X3FV glUniformMatrix2x3fv = nullptr;
PFNGLUNIFORMMATRIX2X4FV glUniformMatrix2x4fv = nullptr;
PFNGLUNIFORMMATRIX3FV glUniformMatrix3fv = nullptr;
PFNGLUNIFORMMATRIX3X2FV glUniformMatrix3x2fv = nullptr;
PFNGLUNIFORMMATRIX3X4FV glUniformMatrix3x4fv = nullptr;
PFNGLUNIFORMMATRIX4FV glUniformMatrix4fv = nullptr;
PFNGLUNIFORMMATRIX4X2FV glUniformMatrix4x2fv = nullptr;
PFNGLUNIFORMMATRIX4X3FV glUniformMatrix4x3fv = nullptr;
PFNGLUSEPROGRAM glUseProgram = nullptr;
PFNGLUSEPROGRAMSTAGES glUseProgramStages = nullptr;
PFNGLVALIDATEPROGRAM glValidateProgram = nullptr;
PFNGLVALIDATEPROGRAMPIPELINE glValidateProgramPipeline = nullptr;
PFNGLCULLFACE glCullFace = nullptr;
PFNGLDEPTHRANGEF glDepthRangef = nullptr;
PFNGLFRONTFACE glFrontFace = nullptr;
PFNGLGETMULTISAMPLEFV glGetMultisamplefv = nullptr;
PFNGLLINEWIDTH glLineWidth = nullptr;
PFNGLMINSAMPLESHADING glMinSampleShading = nullptr;
PFNGLPOLYGONOFFSET glPolygonOffset = nullptr;
PFNGLVIEWPORT glViewport = nullptr;
PFNGLGETBOOLEANI_V glGetBooleani_v = nullptr;
PFNGLGETBOOLEANV glGetBooleanv = nullptr;
PFNGLGETFLOATV glGetFloatv = nullptr;
PFNGLGETINTEGER64I_V glGetInteger64i_v = nullptr;
PFNGLGETINTEGER64V glGetInteger64v = nullptr;
PFNGLGETINTEGERI_V glGetIntegeri_v = nullptr;
PFNGLGETINTEGERV glGetIntegerv = nullptr;
PFNGLGETINTERNALFORMATIV glGetInternalformativ = nullptr;
PFNGLGETSTRING glGetString = nullptr;
PFNGLGETSTRINGI glGetStringi = nullptr;
PFNGLISENABLED glIsEnabled = nullptr;
PFNGLISENABLEDI glIsEnabledi = nullptr;
PFNGLCLIENTWAITSYNC glClientWaitSync = nullptr;
PFNGLDELETESYNC glDeleteSync = nullptr;
PFNGLFENCESYNC glFenceSync = nullptr;
PFNGLGETSYNCIV glGetSynciv = nullptr;
PFNGLISSYNC glIsSync = nullptr;
PFNGLWAITSYNC glWaitSync = nullptr;
PFNGLACTIVETEXTURE glActiveTexture = nullptr;
PFNGLBINDIMAGETEXTURE glBindImageTexture = nullptr;
PFNGLBINDSAMPLER glBindSampler = nullptr;
PFNGLBINDTEXTURE glBindTexture = nullptr;
PFNGLCOMPRESSEDTEXIMAGE2D glCompressedTexImage2D = nullptr;
PFNGLCOMPRESSEDTEXIMAGE3D glCompressedTexImage3D = nullptr;
PFNGLCOMPRESSEDTEXSUBIMAGE2D glCompressedTexSubImage2D = nullptr;
PFNGLCOMPRESSEDTEXSUBIMAGE3D glCompressedTexSubImage3D = nullptr;
PFNGLCOPYIMAGESUBDATA glCopyImageSubData = nullptr;
PFNGLCOPYTEXIMAGE2D glCopyTexImage2D = nullptr;
PFNGLCOPYTEXSUBIMAGE2D glCopyTexSubImage2D = nullptr;
PFNGLCOPYTEXSUBIMAGE3D glCopyTexSubImage3D = nullptr;
PFNGLDELETESAMPLERS glDeleteSamplers = nullptr;
PFNGLDELETETEXTURES glDeleteTextures = nullptr;
PFNGLGENSAMPLERS glGenSamplers = nullptr;
PFNGLGENTEXTURES glGenTextures = nullptr;
PFNGLGENERATEMIPMAP glGenerateMipmap = nullptr;
PFNGLGETSAMPLERPARAMETERIIV glGetSamplerParameterIiv = nullptr;
PFNGLGETSAMPLERPARAMETERIUIV glGetSamplerParameterIuiv = nullptr;
PFNGLGETSAMPLERPARAMETERFV glGetSamplerParameterfv = nullptr;
PFNGLGETSAMPLERPARAMETERIV glGetSamplerParameteriv = nullptr;
PFNGLGETTEXLEVELPARAMETERFV glGetTexLevelParameterfv = nullptr;
PFNGLGETTEXLEVELPARAMETERIV glGetTexLevelParameteriv = nullptr;
PFNGLGETTEXPARAMETERIIV glGetTexParameterIiv = nullptr;
PFNGLGETTEXPARAMETERIUIV glGetTexParameterIuiv = nullptr;
PFNGLGETTEXPARAMETERFV glGetTexParameterfv = nullptr;
PFNGLGETTEXPARAMETERIV glGetTexParameteriv = nullptr;
PFNGLISSAMPLER glIsSampler = nullptr;
PFNGLISTEXTURE glIsTexture = nullptr;
PFNGLPIXELSTOREI glPixelStorei = nullptr;
PFNGLSAMPLERPARAMETERIIV glSamplerParameterIiv = nullptr;
PFNGLSAMPLERPARAMETERIUIV glSamplerParameterIuiv = nullptr;
PFNGLSAMPLERPARAMETERF glSamplerParameterf = nullptr;
PFNGLSAMPLERPARAMETERFV glSamplerParameterfv = nullptr;
PFNGLSAMPLERPARAMETERI glSamplerParameteri = nullptr;
PFNGLSAMPLERPARAMETERIV glSamplerParameteriv = nullptr;
PFNGLTEXBUFFER glTexBuffer = nullptr;
PFNGLTEXBUFFERRANGE glTexBufferRange = nullptr;
PFNGLTEXIMAGE2D glTexImage2D = nullptr;
PFNGLTEXIMAGE3D glTexImage3D = nullptr;
PFNGLTEXPARAMETERIIV glTexParameterIiv = nullptr;
PFNGLTEXPARAMETERIUIV glTexParameterIuiv = nullptr;
PFNGLTEXPARAMETERF glTexParameterf = nullptr;
PFNGLTEXPARAMETERFV glTexParameterfv = nullptr;
PFNGLTEXPARAMETERI glTexParameteri = nullptr;
PFNGLTEXPARAMETERIV glTexParameteriv = nullptr;
PFNGLTEXSTORAGE2D glTexStorage2D = nullptr;
PFNGLTEXSTORAGE2DMULTISAMPLE glTexStorage2DMultisample = nullptr;
PFNGLTEXSTORAGE3D glTexStorage3D = nullptr;
PFNGLTEXSTORAGE3DMULTISAMPLE glTexStorage3DMultisample = nullptr;
PFNGLTEXSUBIMAGE2D glTexSubImage2D = nullptr;
PFNGLTEXSUBIMAGE3D glTexSubImage3D = nullptr;
PFNGLBEGINTRANSFORMFEEDBACK glBeginTransformFeedback = nullptr;
PFNGLBINDTRANSFORMFEEDBACK glBindTransformFeedback = nullptr;
PFNGLDELETETRANSFORMFEEDBACKS glDeleteTransformFeedbacks = nullptr;
PFNGLENDTRANSFORMFEEDBACK glEndTransformFeedback = nullptr;
PFNGLGENTRANSFORMFEEDBACKS glGenTransformFeedbacks = nullptr;
PFNGLGETTRANSFORMFEEDBACKVARYING glGetTransformFeedbackVarying = nullptr;
PFNGLISTRANSFORMFEEDBACK glIsTransformFeedback = nullptr;
PFNGLPAUSETRANSFORMFEEDBACK glPauseTransformFeedback = nullptr;
PFNGLRESUMETRANSFORMFEEDBACK glResumeTransformFeedback = nullptr;
PFNGLTRANSFORMFEEDBACKVARYINGS glTransformFeedbackVaryings = nullptr;
PFNGLBINDVERTEXARRAY glBindVertexArray = nullptr;
PFNGLBINDVERTEXBUFFER glBindVertexBuffer = nullptr;
PFNGLDELETEVERTEXARRAYS glDeleteVertexArrays = nullptr;
PFNGLDISABLEVERTEXATTRIBARRAY glDisableVertexAttribArray = nullptr;
PFNGLENABLEVERTEXATTRIBARRAY glEnableVertexAttribArray = nullptr;
PFNGLGENVERTEXARRAYS glGenVertexArrays = nullptr;
PFNGLGETVERTEXATTRIBIIV glGetVertexAttribIiv = nullptr;
PFNGLGETVERTEXATTRIBIUIV glGetVertexAttribIuiv = nullptr;
PFNGLGETVERTEXATTRIBPOINTERV glGetVertexAttribPointerv = nullptr;
PFNGLGETVERTEXATTRIBFV glGetVertexAttribfv = nullptr;
PFNGLGETVERTEXATTRIBIV glGetVertexAttribiv = nullptr;
PFNGLISVERTEXARRAY glIsVertexArray = nullptr;
PFNGLVERTEXATTRIB1F glVertexAttrib1f = nullptr;
PFNGLVERTEXATTRIB1FV glVertexAttrib1fv = nullptr;
PFNGLVERTEXATTRIB2F glVertexAttrib2f = nullptr;
PFNGLVERTEXATTRIB2FV glVertexAttrib2fv = nullptr;
PFNGLVERTEXATTRIB3F glVertexAttrib3f = nullptr;
PFNGLVERTEXATTRIB3FV glVertexAttrib3fv = nullptr;
PFNGLVERTEXATTRIB4F glVertexAttrib4f = nullptr;
PFNGLVERTEXATTRIB4FV glVertexAttrib4fv = nullptr;
PFNGLVERTEXATTRIBBINDING glVertexAttribBinding = nullptr;
PFNGLVERTEXATTRIBDIVISOR glVertexAttribDivisor = nullptr;
PFNGLVERTEXATTRIBFORMAT glVertexAttribFormat = nullptr;
PFNGLVERTEXATTRIBI4I glVertexAttribI4i = nullptr;
PFNGLVERTEXATTRIBI4IV glVertexAttribI4iv = nullptr;
PFNGLVERTEXATTRIBI4UI glVertexAttribI4ui = nullptr;
PFNGLVERTEXATTRIBI4UIV glVertexAttribI4uiv = nullptr;
PFNGLVERTEXATTRIBIFORMAT glVertexAttribIFormat = nullptr;
PFNGLVERTEXATTRIBIPOINTER glVertexAttribIPointer = nullptr;
PFNGLVERTEXATTRIBPOINTER glVertexAttribPointer = nullptr;
PFNGLVERTEXBINDINGDIVISOR glVertexBindingDivisor = nullptr;
PFNEGLINITIALIZE eglInitialize = nullptr;
PFNEGLCREATECONTEXT eglCreateContext = nullptr;
PFNEGLMAKECURRENT eglMakeCurrent = nullptr;
PFNEGLSWAPBUFFERS eglSwapBuffers = nullptr;
PFNEGLQUERYSURFACE eglQuerySurface = nullptr;
PFNGLXCREATECONTEXT glXCreateContext = nullptr;
PFNGLXCREATENEWCONTEXT glXCreateNewContext = nullptr;
PFNGLXMAKECONTEXTCURRENT glXMakeContextCurrent = nullptr;
PFNGLXMAKECURRENT glXMakeCurrent = nullptr;
PFNGLXSWAPBUFFERS glXSwapBuffers = nullptr;
PFNGLXQUERYDRAWABLE glXQueryDrawable = nullptr;
PFNWGLCREATECONTEXT wglCreateContext = nullptr;
PFNWGLCREATECONTEXTATTRIBSARB wglCreateContextAttribsARB = nullptr;
PFNWGLMAKECURRENT wglMakeCurrent = nullptr;
PFNWGLSWAPBUFFERS wglSwapBuffers = nullptr;
PFNCGLCREATECONTEXT CGLCreateContext = nullptr;
PFNCGLSETCURRENTCONTEXT CGLSetCurrentContext = nullptr;
PFNCGLGETSURFACE CGLGetSurface = nullptr;
PFNCGSGETSURFACEBOUNDS CGSGetSurfaceBounds = nullptr;
PFNCGLFLUSHDRAWABLE CGLFlushDrawable = nullptr;
PFNGLGETQUERYOBJECTI64V glGetQueryObjecti64v = nullptr;
PFNGLGETQUERYOBJECTUI64V glGetQueryObjectui64v = nullptr;

void Register(Interpreter* interpreter) {
    interpreter->registerFunction(Ids::GlBlendBarrierKHR, callGlBlendBarrierKHR);
    interpreter->registerFunction(Ids::GlBlendEquationSeparateiEXT,
                                  callGlBlendEquationSeparateiEXT);
    interpreter->registerFunction(Ids::GlBlendEquationiEXT, callGlBlendEquationiEXT);
    interpreter->registerFunction(Ids::GlBlendFuncSeparateiEXT, callGlBlendFuncSeparateiEXT);
    interpreter->registerFunction(Ids::GlBlendFunciEXT, callGlBlendFunciEXT);
    interpreter->registerFunction(Ids::GlColorMaskiEXT, callGlColorMaskiEXT);
    interpreter->registerFunction(Ids::GlCopyImageSubDataEXT, callGlCopyImageSubDataEXT);
    interpreter->registerFunction(Ids::GlDebugMessageCallbackKHR, callGlDebugMessageCallbackKHR);
    interpreter->registerFunction(Ids::GlDebugMessageControlKHR, callGlDebugMessageControlKHR);
    interpreter->registerFunction(Ids::GlDebugMessageInsertKHR, callGlDebugMessageInsertKHR);
    interpreter->registerFunction(Ids::GlDisableiEXT, callGlDisableiEXT);
    interpreter->registerFunction(Ids::GlEnableiEXT, callGlEnableiEXT);
    interpreter->registerFunction(Ids::GlFramebufferTextureEXT, callGlFramebufferTextureEXT);
    interpreter->registerFunction(Ids::GlGetDebugMessageLogKHR, callGlGetDebugMessageLogKHR);
    interpreter->registerFunction(Ids::GlGetObjectLabelKHR, callGlGetObjectLabelKHR);
    interpreter->registerFunction(Ids::GlGetObjectPtrLabelKHR, callGlGetObjectPtrLabelKHR);
    interpreter->registerFunction(Ids::GlGetPointervKHR, callGlGetPointervKHR);
    interpreter->registerFunction(Ids::GlGetSamplerParameterIivEXT,
                                  callGlGetSamplerParameterIivEXT);
    interpreter->registerFunction(Ids::GlGetSamplerParameterIuivEXT,
                                  callGlGetSamplerParameterIuivEXT);
    interpreter->registerFunction(Ids::GlGetTexParameterIivEXT, callGlGetTexParameterIivEXT);
    interpreter->registerFunction(Ids::GlGetTexParameterIuivEXT, callGlGetTexParameterIuivEXT);
    interpreter->registerFunction(Ids::GlIsEnablediEXT, callGlIsEnablediEXT);
    interpreter->registerFunction(Ids::GlMinSampleShadingOES, callGlMinSampleShadingOES);
    interpreter->registerFunction(Ids::GlObjectLabelKHR, callGlObjectLabelKHR);
    interpreter->registerFunction(Ids::GlObjectPtrLabelKHR, callGlObjectPtrLabelKHR);
    interpreter->registerFunction(Ids::GlPatchParameteriEXT, callGlPatchParameteriEXT);
    interpreter->registerFunction(Ids::GlPopDebugGroupKHR, callGlPopDebugGroupKHR);
    interpreter->registerFunction(Ids::GlPrimitiveBoundingBoxEXT, callGlPrimitiveBoundingBoxEXT);
    interpreter->registerFunction(Ids::GlPushDebugGroupKHR, callGlPushDebugGroupKHR);
    interpreter->registerFunction(Ids::GlSamplerParameterIivEXT, callGlSamplerParameterIivEXT);
    interpreter->registerFunction(Ids::GlSamplerParameterIuivEXT, callGlSamplerParameterIuivEXT);
    interpreter->registerFunction(Ids::GlTexBufferEXT, callGlTexBufferEXT);
    interpreter->registerFunction(Ids::GlTexBufferRangeEXT, callGlTexBufferRangeEXT);
    interpreter->registerFunction(Ids::GlTexParameterIivEXT, callGlTexParameterIivEXT);
    interpreter->registerFunction(Ids::GlTexParameterIuivEXT, callGlTexParameterIuivEXT);
    interpreter->registerFunction(Ids::GlTexStorage3DMultisampleOES,
                                  callGlTexStorage3DMultisampleOES);
    interpreter->registerFunction(Ids::GlBeginQuery, callGlBeginQuery);
    interpreter->registerFunction(Ids::GlDeleteQueries, callGlDeleteQueries);
    interpreter->registerFunction(Ids::GlEndQuery, callGlEndQuery);
    interpreter->registerFunction(Ids::GlGenQueries, callGlGenQueries);
    interpreter->registerFunction(Ids::GlGetQueryObjectuiv, callGlGetQueryObjectuiv);
    interpreter->registerFunction(Ids::GlGetQueryiv, callGlGetQueryiv);
    interpreter->registerFunction(Ids::GlIsQuery, callGlIsQuery);
    interpreter->registerFunction(Ids::GlBindBuffer, callGlBindBuffer);
    interpreter->registerFunction(Ids::GlBindBufferBase, callGlBindBufferBase);
    interpreter->registerFunction(Ids::GlBindBufferRange, callGlBindBufferRange);
    interpreter->registerFunction(Ids::GlBufferData, callGlBufferData);
    interpreter->registerFunction(Ids::GlBufferSubData, callGlBufferSubData);
    interpreter->registerFunction(Ids::GlCopyBufferSubData, callGlCopyBufferSubData);
    interpreter->registerFunction(Ids::GlDeleteBuffers, callGlDeleteBuffers);
    interpreter->registerFunction(Ids::GlGenBuffers, callGlGenBuffers);
    interpreter->registerFunction(Ids::GlGetBufferParameteri64v, callGlGetBufferParameteri64v);
    interpreter->registerFunction(Ids::GlGetBufferParameteriv, callGlGetBufferParameteriv);
    interpreter->registerFunction(Ids::GlGetBufferPointerv, callGlGetBufferPointerv);
    interpreter->registerFunction(Ids::GlIsBuffer, callGlIsBuffer);
    interpreter->registerFunction(Ids::GlMapBufferRange, callGlMapBufferRange);
    interpreter->registerFunction(Ids::GlUnmapBuffer, callGlUnmapBuffer);
    interpreter->registerFunction(Ids::GlDebugMessageCallback, callGlDebugMessageCallback);
    interpreter->registerFunction(Ids::GlDebugMessageControl, callGlDebugMessageControl);
    interpreter->registerFunction(Ids::GlDebugMessageInsert, callGlDebugMessageInsert);
    interpreter->registerFunction(Ids::GlGetDebugMessageLog, callGlGetDebugMessageLog);
    interpreter->registerFunction(Ids::GlGetObjectLabel, callGlGetObjectLabel);
    interpreter->registerFunction(Ids::GlGetObjectPtrLabel, callGlGetObjectPtrLabel);
    interpreter->registerFunction(Ids::GlGetPointerv, callGlGetPointerv);
    interpreter->registerFunction(Ids::GlObjectLabel, callGlObjectLabel);
    interpreter->registerFunction(Ids::GlObjectPtrLabel, callGlObjectPtrLabel);
    interpreter->registerFunction(Ids::GlPopDebugGroup, callGlPopDebugGroup);
    interpreter->registerFunction(Ids::GlPushDebugGroup, callGlPushDebugGroup);
    interpreter->registerFunction(Ids::GlDrawArrays, callGlDrawArrays);
    interpreter->registerFunction(Ids::GlDrawArraysIndirect, callGlDrawArraysIndirect);
    interpreter->registerFunction(Ids::GlDrawArraysInstanced, callGlDrawArraysInstanced);
    interpreter->registerFunction(Ids::GlDrawElements, callGlDrawElements);
    interpreter->registerFunction(Ids::GlDrawElementsBaseVertex, callGlDrawElementsBaseVertex);
    interpreter->registerFunction(Ids::GlDrawElementsIndirect, callGlDrawElementsIndirect);
    interpreter->registerFunction(Ids::GlDrawElementsInstanced, callGlDrawElementsInstanced);
    interpreter->registerFunction(Ids::GlDrawElementsInstancedBaseVertex,
                                  callGlDrawElementsInstancedBaseVertex);
    interpreter->registerFunction(Ids::GlDrawRangeElements, callGlDrawRangeElements);
    interpreter->registerFunction(Ids::GlDrawRangeElementsBaseVertex,
                                  callGlDrawRangeElementsBaseVertex);
    interpreter->registerFunction(Ids::GlPatchParameteri, callGlPatchParameteri);
    interpreter->registerFunction(Ids::GlPrimitiveBoundingBox, callGlPrimitiveBoundingBox);
    interpreter->registerFunction(Ids::GlActiveShaderProgramEXT, callGlActiveShaderProgramEXT);
    interpreter->registerFunction(Ids::GlAlphaFuncQCOM, callGlAlphaFuncQCOM);
    interpreter->registerFunction(Ids::GlApplyFramebufferAttachmentCMAAINTEL,
                                  callGlApplyFramebufferAttachmentCMAAINTEL);
    interpreter->registerFunction(Ids::GlBeginConditionalRenderNV, callGlBeginConditionalRenderNV);
    interpreter->registerFunction(Ids::GlBeginPerfMonitorAMD, callGlBeginPerfMonitorAMD);
    interpreter->registerFunction(Ids::GlBeginPerfQueryINTEL, callGlBeginPerfQueryINTEL);
    interpreter->registerFunction(Ids::GlBeginQueryEXT, callGlBeginQueryEXT);
    interpreter->registerFunction(Ids::GlBindProgramPipelineEXT, callGlBindProgramPipelineEXT);
    interpreter->registerFunction(Ids::GlBindVertexArrayOES, callGlBindVertexArrayOES);
    interpreter->registerFunction(Ids::GlBlendBarrierNV, callGlBlendBarrierNV);
    interpreter->registerFunction(Ids::GlBlendEquationSeparateiOES,
                                  callGlBlendEquationSeparateiOES);
    interpreter->registerFunction(Ids::GlBlendEquationiOES, callGlBlendEquationiOES);
    interpreter->registerFunction(Ids::GlBlendFuncSeparateiOES, callGlBlendFuncSeparateiOES);
    interpreter->registerFunction(Ids::GlBlendFunciOES, callGlBlendFunciOES);
    interpreter->registerFunction(Ids::GlBlendParameteriNV, callGlBlendParameteriNV);
    interpreter->registerFunction(Ids::GlBlitFramebufferANGLE, callGlBlitFramebufferANGLE);
    interpreter->registerFunction(Ids::GlBlitFramebufferNV, callGlBlitFramebufferNV);
    interpreter->registerFunction(Ids::GlBufferStorageEXT, callGlBufferStorageEXT);
    interpreter->registerFunction(Ids::GlClientWaitSyncAPPLE, callGlClientWaitSyncAPPLE);
    interpreter->registerFunction(Ids::GlColorMaskiOES, callGlColorMaskiOES);
    interpreter->registerFunction(Ids::GlCompressedTexImage3DOES, callGlCompressedTexImage3DOES);
    interpreter->registerFunction(Ids::GlCompressedTexSubImage3DOES,
                                  callGlCompressedTexSubImage3DOES);
    interpreter->registerFunction(Ids::GlCopyBufferSubDataNV, callGlCopyBufferSubDataNV);
    interpreter->registerFunction(Ids::GlCopyImageSubDataOES, callGlCopyImageSubDataOES);
    interpreter->registerFunction(Ids::GlCopyPathNV, callGlCopyPathNV);
    interpreter->registerFunction(Ids::GlCopyTexSubImage3DOES, callGlCopyTexSubImage3DOES);
    interpreter->registerFunction(Ids::GlCopyTextureLevelsAPPLE, callGlCopyTextureLevelsAPPLE);
    interpreter->registerFunction(Ids::GlCoverFillPathInstancedNV, callGlCoverFillPathInstancedNV);
    interpreter->registerFunction(Ids::GlCoverFillPathNV, callGlCoverFillPathNV);
    interpreter->registerFunction(Ids::GlCoverStrokePathInstancedNV,
                                  callGlCoverStrokePathInstancedNV);
    interpreter->registerFunction(Ids::GlCoverStrokePathNV, callGlCoverStrokePathNV);
    interpreter->registerFunction(Ids::GlCoverageMaskNV, callGlCoverageMaskNV);
    interpreter->registerFunction(Ids::GlCoverageModulationNV, callGlCoverageModulationNV);
    interpreter->registerFunction(Ids::GlCoverageModulationTableNV,
                                  callGlCoverageModulationTableNV);
    interpreter->registerFunction(Ids::GlCoverageOperationNV, callGlCoverageOperationNV);
    interpreter->registerFunction(Ids::GlCreatePerfQueryINTEL, callGlCreatePerfQueryINTEL);
    interpreter->registerFunction(Ids::GlCreateShaderProgramvEXT, callGlCreateShaderProgramvEXT);
    interpreter->registerFunction(Ids::GlDeleteFencesNV, callGlDeleteFencesNV);
    interpreter->registerFunction(Ids::GlDeletePathsNV, callGlDeletePathsNV);
    interpreter->registerFunction(Ids::GlDeletePerfMonitorsAMD, callGlDeletePerfMonitorsAMD);
    interpreter->registerFunction(Ids::GlDeletePerfQueryINTEL, callGlDeletePerfQueryINTEL);
    interpreter->registerFunction(Ids::GlDeleteProgramPipelinesEXT,
                                  callGlDeleteProgramPipelinesEXT);
    interpreter->registerFunction(Ids::GlDeleteQueriesEXT, callGlDeleteQueriesEXT);
    interpreter->registerFunction(Ids::GlDeleteSyncAPPLE, callGlDeleteSyncAPPLE);
    interpreter->registerFunction(Ids::GlDeleteVertexArraysOES, callGlDeleteVertexArraysOES);
    interpreter->registerFunction(Ids::GlDepthRangeArrayfvNV, callGlDepthRangeArrayfvNV);
    interpreter->registerFunction(Ids::GlDepthRangeIndexedfNV, callGlDepthRangeIndexedfNV);
    interpreter->registerFunction(Ids::GlDisableDriverControlQCOM, callGlDisableDriverControlQCOM);
    interpreter->registerFunction(Ids::GlDisableiNV, callGlDisableiNV);
    interpreter->registerFunction(Ids::GlDisableiOES, callGlDisableiOES);
    interpreter->registerFunction(Ids::GlDiscardFramebufferEXT, callGlDiscardFramebufferEXT);
    interpreter->registerFunction(Ids::GlDrawArraysInstancedANGLE, callGlDrawArraysInstancedANGLE);
    interpreter->registerFunction(Ids::GlDrawArraysInstancedBaseInstanceEXT,
                                  callGlDrawArraysInstancedBaseInstanceEXT);
    interpreter->registerFunction(Ids::GlDrawArraysInstancedEXT, callGlDrawArraysInstancedEXT);
    interpreter->registerFunction(Ids::GlDrawArraysInstancedNV, callGlDrawArraysInstancedNV);
    interpreter->registerFunction(Ids::GlDrawBuffersEXT, callGlDrawBuffersEXT);
    interpreter->registerFunction(Ids::GlDrawBuffersIndexedEXT, callGlDrawBuffersIndexedEXT);
    interpreter->registerFunction(Ids::GlDrawBuffersNV, callGlDrawBuffersNV);
    interpreter->registerFunction(Ids::GlDrawElementsBaseVertexEXT,
                                  callGlDrawElementsBaseVertexEXT);
    interpreter->registerFunction(Ids::GlDrawElementsBaseVertexOES,
                                  callGlDrawElementsBaseVertexOES);
    interpreter->registerFunction(Ids::GlDrawElementsInstancedANGLE,
                                  callGlDrawElementsInstancedANGLE);
    interpreter->registerFunction(Ids::GlDrawElementsInstancedBaseInstanceEXT,
                                  callGlDrawElementsInstancedBaseInstanceEXT);
    interpreter->registerFunction(Ids::GlDrawElementsInstancedBaseVertexBaseInstanceEXT,
                                  callGlDrawElementsInstancedBaseVertexBaseInstanceEXT);
    interpreter->registerFunction(Ids::GlDrawElementsInstancedBaseVertexEXT,
                                  callGlDrawElementsInstancedBaseVertexEXT);
    interpreter->registerFunction(Ids::GlDrawElementsInstancedBaseVertexOES,
                                  callGlDrawElementsInstancedBaseVertexOES);
    interpreter->registerFunction(Ids::GlDrawElementsInstancedEXT, callGlDrawElementsInstancedEXT);
    interpreter->registerFunction(Ids::GlDrawElementsInstancedNV, callGlDrawElementsInstancedNV);
    interpreter->registerFunction(Ids::GlDrawRangeElementsBaseVertexEXT,
                                  callGlDrawRangeElementsBaseVertexEXT);
    interpreter->registerFunction(Ids::GlDrawRangeElementsBaseVertexOES,
                                  callGlDrawRangeElementsBaseVertexOES);
    interpreter->registerFunction(Ids::GlEGLImageTargetRenderbufferStorageOES,
                                  callGlEGLImageTargetRenderbufferStorageOES);
    interpreter->registerFunction(Ids::GlEGLImageTargetTexture2DOES,
                                  callGlEGLImageTargetTexture2DOES);
    interpreter->registerFunction(Ids::GlEnableDriverControlQCOM, callGlEnableDriverControlQCOM);
    interpreter->registerFunction(Ids::GlEnableiNV, callGlEnableiNV);
    interpreter->registerFunction(Ids::GlEnableiOES, callGlEnableiOES);
    interpreter->registerFunction(Ids::GlEndConditionalRenderNV, callGlEndConditionalRenderNV);
    interpreter->registerFunction(Ids::GlEndPerfMonitorAMD, callGlEndPerfMonitorAMD);
    interpreter->registerFunction(Ids::GlEndPerfQueryINTEL, callGlEndPerfQueryINTEL);
    interpreter->registerFunction(Ids::GlEndQueryEXT, callGlEndQueryEXT);
    interpreter->registerFunction(Ids::GlEndTilingQCOM, callGlEndTilingQCOM);
    interpreter->registerFunction(Ids::GlExtGetBufferPointervQCOM, callGlExtGetBufferPointervQCOM);
    interpreter->registerFunction(Ids::GlExtGetBuffersQCOM, callGlExtGetBuffersQCOM);
    interpreter->registerFunction(Ids::GlExtGetFramebuffersQCOM, callGlExtGetFramebuffersQCOM);
    interpreter->registerFunction(Ids::GlExtGetProgramBinarySourceQCOM,
                                  callGlExtGetProgramBinarySourceQCOM);
    interpreter->registerFunction(Ids::GlExtGetProgramsQCOM, callGlExtGetProgramsQCOM);
    interpreter->registerFunction(Ids::GlExtGetRenderbuffersQCOM, callGlExtGetRenderbuffersQCOM);
    interpreter->registerFunction(Ids::GlExtGetShadersQCOM, callGlExtGetShadersQCOM);
    interpreter->registerFunction(Ids::GlExtGetTexLevelParameterivQCOM,
                                  callGlExtGetTexLevelParameterivQCOM);
    interpreter->registerFunction(Ids::GlExtGetTexSubImageQCOM, callGlExtGetTexSubImageQCOM);
    interpreter->registerFunction(Ids::GlExtGetTexturesQCOM, callGlExtGetTexturesQCOM);
    interpreter->registerFunction(Ids::GlExtIsProgramBinaryQCOM, callGlExtIsProgramBinaryQCOM);
    interpreter->registerFunction(Ids::GlExtTexObjectStateOverrideiQCOM,
                                  callGlExtTexObjectStateOverrideiQCOM);
    interpreter->registerFunction(Ids::GlFenceSyncAPPLE, callGlFenceSyncAPPLE);
    interpreter->registerFunction(Ids::GlFinishFenceNV, callGlFinishFenceNV);
    interpreter->registerFunction(Ids::GlFlushMappedBufferRangeEXT,
                                  callGlFlushMappedBufferRangeEXT);
    interpreter->registerFunction(Ids::GlFragmentCoverageColorNV, callGlFragmentCoverageColorNV);
    interpreter->registerFunction(Ids::GlFramebufferSampleLocationsfvNV,
                                  callGlFramebufferSampleLocationsfvNV);
    interpreter->registerFunction(Ids::GlFramebufferTexture2DMultisampleEXT,
                                  callGlFramebufferTexture2DMultisampleEXT);
    interpreter->registerFunction(Ids::GlFramebufferTexture2DMultisampleIMG,
                                  callGlFramebufferTexture2DMultisampleIMG);
    interpreter->registerFunction(Ids::GlFramebufferTexture3DOES, callGlFramebufferTexture3DOES);
    interpreter->registerFunction(Ids::GlFramebufferTextureMultiviewOVR,
                                  callGlFramebufferTextureMultiviewOVR);
    interpreter->registerFunction(Ids::GlFramebufferTextureOES, callGlFramebufferTextureOES);
    interpreter->registerFunction(Ids::GlGenFencesNV, callGlGenFencesNV);
    interpreter->registerFunction(Ids::GlGenPathsNV, callGlGenPathsNV);
    interpreter->registerFunction(Ids::GlGenPerfMonitorsAMD, callGlGenPerfMonitorsAMD);
    interpreter->registerFunction(Ids::GlGenProgramPipelinesEXT, callGlGenProgramPipelinesEXT);
    interpreter->registerFunction(Ids::GlGenQueriesEXT, callGlGenQueriesEXT);
    interpreter->registerFunction(Ids::GlGenVertexArraysOES, callGlGenVertexArraysOES);
    interpreter->registerFunction(Ids::GlGetBufferPointervOES, callGlGetBufferPointervOES);
    interpreter->registerFunction(Ids::GlGetCoverageModulationTableNV,
                                  callGlGetCoverageModulationTableNV);
    interpreter->registerFunction(Ids::GlGetDriverControlStringQCOM,
                                  callGlGetDriverControlStringQCOM);
    interpreter->registerFunction(Ids::GlGetDriverControlsQCOM, callGlGetDriverControlsQCOM);
    interpreter->registerFunction(Ids::GlGetFenceivNV, callGlGetFenceivNV);
    interpreter->registerFunction(Ids::GlGetFirstPerfQueryIdINTEL, callGlGetFirstPerfQueryIdINTEL);
    interpreter->registerFunction(Ids::GlGetFloatiVNV, callGlGetFloatiVNV);
    interpreter->registerFunction(Ids::GlGetGraphicsResetStatusEXT,
                                  callGlGetGraphicsResetStatusEXT);
    interpreter->registerFunction(Ids::GlGetGraphicsResetStatusKHR,
                                  callGlGetGraphicsResetStatusKHR);
    interpreter->registerFunction(Ids::GlGetImageHandleNV, callGlGetImageHandleNV);
    interpreter->registerFunction(Ids::GlGetInteger64vAPPLE, callGlGetInteger64vAPPLE);
    interpreter->registerFunction(Ids::GlGetIntegeriVEXT, callGlGetIntegeriVEXT);
    interpreter->registerFunction(Ids::GlGetInternalformatSampleivNV,
                                  callGlGetInternalformatSampleivNV);
    interpreter->registerFunction(Ids::GlGetNextPerfQueryIdINTEL, callGlGetNextPerfQueryIdINTEL);
    interpreter->registerFunction(Ids::GlGetObjectLabelEXT, callGlGetObjectLabelEXT);
    interpreter->registerFunction(Ids::GlGetPathCommandsNV, callGlGetPathCommandsNV);
    interpreter->registerFunction(Ids::GlGetPathCoordsNV, callGlGetPathCoordsNV);
    interpreter->registerFunction(Ids::GlGetPathDashArrayNV, callGlGetPathDashArrayNV);
    interpreter->registerFunction(Ids::GlGetPathLengthNV, callGlGetPathLengthNV);
    interpreter->registerFunction(Ids::GlGetPathMetricRangeNV, callGlGetPathMetricRangeNV);
    interpreter->registerFunction(Ids::GlGetPathMetricsNV, callGlGetPathMetricsNV);
    interpreter->registerFunction(Ids::GlGetPathParameterfvNV, callGlGetPathParameterfvNV);
    interpreter->registerFunction(Ids::GlGetPathParameterivNV, callGlGetPathParameterivNV);
    interpreter->registerFunction(Ids::GlGetPathSpacingNV, callGlGetPathSpacingNV);
    interpreter->registerFunction(Ids::GlGetPerfCounterInfoINTEL, callGlGetPerfCounterInfoINTEL);
    interpreter->registerFunction(Ids::GlGetPerfMonitorCounterDataAMD,
                                  callGlGetPerfMonitorCounterDataAMD);
    interpreter->registerFunction(Ids::GlGetPerfMonitorCounterInfoAMD,
                                  callGlGetPerfMonitorCounterInfoAMD);
    interpreter->registerFunction(Ids::GlGetPerfMonitorCounterStringAMD,
                                  callGlGetPerfMonitorCounterStringAMD);
    interpreter->registerFunction(Ids::GlGetPerfMonitorCountersAMD,
                                  callGlGetPerfMonitorCountersAMD);
    interpreter->registerFunction(Ids::GlGetPerfMonitorGroupStringAMD,
                                  callGlGetPerfMonitorGroupStringAMD);
    interpreter->registerFunction(Ids::GlGetPerfMonitorGroupsAMD, callGlGetPerfMonitorGroupsAMD);
    interpreter->registerFunction(Ids::GlGetPerfQueryDataINTEL, callGlGetPerfQueryDataINTEL);
    interpreter->registerFunction(Ids::GlGetPerfQueryIdByNameINTEL,
                                  callGlGetPerfQueryIdByNameINTEL);
    interpreter->registerFunction(Ids::GlGetPerfQueryInfoINTEL, callGlGetPerfQueryInfoINTEL);
    interpreter->registerFunction(Ids::GlGetProgramBinaryOES, callGlGetProgramBinaryOES);
    interpreter->registerFunction(Ids::GlGetProgramPipelineInfoLogEXT,
                                  callGlGetProgramPipelineInfoLogEXT);
    interpreter->registerFunction(Ids::GlGetProgramPipelineivEXT, callGlGetProgramPipelineivEXT);
    interpreter->registerFunction(Ids::GlGetProgramResourcefvNV, callGlGetProgramResourcefvNV);
    interpreter->registerFunction(Ids::GlGetQueryObjecti64vEXT, callGlGetQueryObjecti64vEXT);
    interpreter->registerFunction(Ids::GlGetQueryObjectivEXT, callGlGetQueryObjectivEXT);
    interpreter->registerFunction(Ids::GlGetQueryObjectui64vEXT, callGlGetQueryObjectui64vEXT);
    interpreter->registerFunction(Ids::GlGetQueryObjectuivEXT, callGlGetQueryObjectuivEXT);
    interpreter->registerFunction(Ids::GlGetQueryivEXT, callGlGetQueryivEXT);
    interpreter->registerFunction(Ids::GlGetSamplerParameterIivOES,
                                  callGlGetSamplerParameterIivOES);
    interpreter->registerFunction(Ids::GlGetSamplerParameterIuivOES,
                                  callGlGetSamplerParameterIuivOES);
    interpreter->registerFunction(Ids::GlGetSyncivAPPLE, callGlGetSyncivAPPLE);
    interpreter->registerFunction(Ids::GlGetTexParameterIivOES, callGlGetTexParameterIivOES);
    interpreter->registerFunction(Ids::GlGetTexParameterIuivOES, callGlGetTexParameterIuivOES);
    interpreter->registerFunction(Ids::GlGetTextureHandleNV, callGlGetTextureHandleNV);
    interpreter->registerFunction(Ids::GlGetTextureSamplerHandleNV,
                                  callGlGetTextureSamplerHandleNV);
    interpreter->registerFunction(Ids::GlGetTranslatedShaderSourceANGLE,
                                  callGlGetTranslatedShaderSourceANGLE);
    interpreter->registerFunction(Ids::GlGetnUniformfvEXT, callGlGetnUniformfvEXT);
    interpreter->registerFunction(Ids::GlGetnUniformfvKHR, callGlGetnUniformfvKHR);
    interpreter->registerFunction(Ids::GlGetnUniformivEXT, callGlGetnUniformivEXT);
    interpreter->registerFunction(Ids::GlGetnUniformivKHR, callGlGetnUniformivKHR);
    interpreter->registerFunction(Ids::GlGetnUniformuivKHR, callGlGetnUniformuivKHR);
    interpreter->registerFunction(Ids::GlInsertEventMarkerEXT, callGlInsertEventMarkerEXT);
    interpreter->registerFunction(Ids::GlInterpolatePathsNV, callGlInterpolatePathsNV);
    interpreter->registerFunction(Ids::GlIsEnablediNV, callGlIsEnablediNV);
    interpreter->registerFunction(Ids::GlIsEnablediOES, callGlIsEnablediOES);
    interpreter->registerFunction(Ids::GlIsFenceNV, callGlIsFenceNV);
    interpreter->registerFunction(Ids::GlIsImageHandleResidentNV, callGlIsImageHandleResidentNV);
    interpreter->registerFunction(Ids::GlIsPathNV, callGlIsPathNV);
    interpreter->registerFunction(Ids::GlIsPointInFillPathNV, callGlIsPointInFillPathNV);
    interpreter->registerFunction(Ids::GlIsPointInStrokePathNV, callGlIsPointInStrokePathNV);
    interpreter->registerFunction(Ids::GlIsProgramPipelineEXT, callGlIsProgramPipelineEXT);
    interpreter->registerFunction(Ids::GlIsQueryEXT, callGlIsQueryEXT);
    interpreter->registerFunction(Ids::GlIsSyncAPPLE, callGlIsSyncAPPLE);
    interpreter->registerFunction(Ids::GlIsTextureHandleResidentNV,
                                  callGlIsTextureHandleResidentNV);
    interpreter->registerFunction(Ids::GlIsVertexArrayOES, callGlIsVertexArrayOES);
    interpreter->registerFunction(Ids::GlLabelObjectEXT, callGlLabelObjectEXT);
    interpreter->registerFunction(Ids::GlMakeImageHandleNonResidentNV,
                                  callGlMakeImageHandleNonResidentNV);
    interpreter->registerFunction(Ids::GlMakeImageHandleResidentNV,
                                  callGlMakeImageHandleResidentNV);
    interpreter->registerFunction(Ids::GlMakeTextureHandleNonResidentNV,
                                  callGlMakeTextureHandleNonResidentNV);
    interpreter->registerFunction(Ids::GlMakeTextureHandleResidentNV,
                                  callGlMakeTextureHandleResidentNV);
    interpreter->registerFunction(Ids::GlMapBufferOES, callGlMapBufferOES);
    interpreter->registerFunction(Ids::GlMapBufferRangeEXT, callGlMapBufferRangeEXT);
    interpreter->registerFunction(Ids::GlMatrixLoad3x2fNV, callGlMatrixLoad3x2fNV);
    interpreter->registerFunction(Ids::GlMatrixLoad3x3fNV, callGlMatrixLoad3x3fNV);
    interpreter->registerFunction(Ids::GlMatrixLoadTranspose3x3fNV,
                                  callGlMatrixLoadTranspose3x3fNV);
    interpreter->registerFunction(Ids::GlMatrixMult3x2fNV, callGlMatrixMult3x2fNV);
    interpreter->registerFunction(Ids::GlMatrixMult3x3fNV, callGlMatrixMult3x3fNV);
    interpreter->registerFunction(Ids::GlMatrixMultTranspose3x3fNV,
                                  callGlMatrixMultTranspose3x3fNV);
    interpreter->registerFunction(Ids::GlMultiDrawArraysEXT, callGlMultiDrawArraysEXT);
    interpreter->registerFunction(Ids::GlMultiDrawArraysIndirectEXT,
                                  callGlMultiDrawArraysIndirectEXT);
    interpreter->registerFunction(Ids::GlMultiDrawElementsBaseVertexEXT,
                                  callGlMultiDrawElementsBaseVertexEXT);
    interpreter->registerFunction(Ids::GlMultiDrawElementsBaseVertexOES,
                                  callGlMultiDrawElementsBaseVertexOES);
    interpreter->registerFunction(Ids::GlMultiDrawElementsEXT, callGlMultiDrawElementsEXT);
    interpreter->registerFunction(Ids::GlMultiDrawElementsIndirectEXT,
                                  callGlMultiDrawElementsIndirectEXT);
    interpreter->registerFunction(Ids::GlNamedFramebufferSampleLocationsfvNV,
                                  callGlNamedFramebufferSampleLocationsfvNV);
    interpreter->registerFunction(Ids::GlPatchParameteriOES, callGlPatchParameteriOES);
    interpreter->registerFunction(Ids::GlPathCommandsNV, callGlPathCommandsNV);
    interpreter->registerFunction(Ids::GlPathCoordsNV, callGlPathCoordsNV);
    interpreter->registerFunction(Ids::GlPathCoverDepthFuncNV, callGlPathCoverDepthFuncNV);
    interpreter->registerFunction(Ids::GlPathDashArrayNV, callGlPathDashArrayNV);
    interpreter->registerFunction(Ids::GlPathGlyphIndexArrayNV, callGlPathGlyphIndexArrayNV);
    interpreter->registerFunction(Ids::GlPathGlyphIndexRangeNV, callGlPathGlyphIndexRangeNV);
    interpreter->registerFunction(Ids::GlPathGlyphRangeNV, callGlPathGlyphRangeNV);
    interpreter->registerFunction(Ids::GlPathGlyphsNV, callGlPathGlyphsNV);
    interpreter->registerFunction(Ids::GlPathMemoryGlyphIndexArrayNV,
                                  callGlPathMemoryGlyphIndexArrayNV);
    interpreter->registerFunction(Ids::GlPathParameterfNV, callGlPathParameterfNV);
    interpreter->registerFunction(Ids::GlPathParameterfvNV, callGlPathParameterfvNV);
    interpreter->registerFunction(Ids::GlPathParameteriNV, callGlPathParameteriNV);
    interpreter->registerFunction(Ids::GlPathParameterivNV, callGlPathParameterivNV);
    interpreter->registerFunction(Ids::GlPathStencilDepthOffsetNV, callGlPathStencilDepthOffsetNV);
    interpreter->registerFunction(Ids::GlPathStencilFuncNV, callGlPathStencilFuncNV);
    interpreter->registerFunction(Ids::GlPathStringNV, callGlPathStringNV);
    interpreter->registerFunction(Ids::GlPathSubCommandsNV, callGlPathSubCommandsNV);
    interpreter->registerFunction(Ids::GlPathSubCoordsNV, callGlPathSubCoordsNV);
    interpreter->registerFunction(Ids::GlPointAlongPathNV, callGlPointAlongPathNV);
    interpreter->registerFunction(Ids::GlPolygonModeNV, callGlPolygonModeNV);
    interpreter->registerFunction(Ids::GlPopGroupMarkerEXT, callGlPopGroupMarkerEXT);
    interpreter->registerFunction(Ids::GlPrimitiveBoundingBoxOES, callGlPrimitiveBoundingBoxOES);
    interpreter->registerFunction(Ids::GlProgramBinaryOES, callGlProgramBinaryOES);
    interpreter->registerFunction(Ids::GlProgramParameteriEXT, callGlProgramParameteriEXT);
    interpreter->registerFunction(Ids::GlProgramPathFragmentInputGenNV,
                                  callGlProgramPathFragmentInputGenNV);
    interpreter->registerFunction(Ids::GlProgramUniform1fEXT, callGlProgramUniform1fEXT);
    interpreter->registerFunction(Ids::GlProgramUniform1fvEXT, callGlProgramUniform1fvEXT);
    interpreter->registerFunction(Ids::GlProgramUniform1iEXT, callGlProgramUniform1iEXT);
    interpreter->registerFunction(Ids::GlProgramUniform1ivEXT, callGlProgramUniform1ivEXT);
    interpreter->registerFunction(Ids::GlProgramUniform1uiEXT, callGlProgramUniform1uiEXT);
    interpreter->registerFunction(Ids::GlProgramUniform1uivEXT, callGlProgramUniform1uivEXT);
    interpreter->registerFunction(Ids::GlProgramUniform2fEXT, callGlProgramUniform2fEXT);
    interpreter->registerFunction(Ids::GlProgramUniform2fvEXT, callGlProgramUniform2fvEXT);
    interpreter->registerFunction(Ids::GlProgramUniform2iEXT, callGlProgramUniform2iEXT);
    interpreter->registerFunction(Ids::GlProgramUniform2ivEXT, callGlProgramUniform2ivEXT);
    interpreter->registerFunction(Ids::GlProgramUniform2uiEXT, callGlProgramUniform2uiEXT);
    interpreter->registerFunction(Ids::GlProgramUniform2uivEXT, callGlProgramUniform2uivEXT);
    interpreter->registerFunction(Ids::GlProgramUniform3fEXT, callGlProgramUniform3fEXT);
    interpreter->registerFunction(Ids::GlProgramUniform3fvEXT, callGlProgramUniform3fvEXT);
    interpreter->registerFunction(Ids::GlProgramUniform3iEXT, callGlProgramUniform3iEXT);
    interpreter->registerFunction(Ids::GlProgramUniform3ivEXT, callGlProgramUniform3ivEXT);
    interpreter->registerFunction(Ids::GlProgramUniform3uiEXT, callGlProgramUniform3uiEXT);
    interpreter->registerFunction(Ids::GlProgramUniform3uivEXT, callGlProgramUniform3uivEXT);
    interpreter->registerFunction(Ids::GlProgramUniform4fEXT, callGlProgramUniform4fEXT);
    interpreter->registerFunction(Ids::GlProgramUniform4fvEXT, callGlProgramUniform4fvEXT);
    interpreter->registerFunction(Ids::GlProgramUniform4iEXT, callGlProgramUniform4iEXT);
    interpreter->registerFunction(Ids::GlProgramUniform4ivEXT, callGlProgramUniform4ivEXT);
    interpreter->registerFunction(Ids::GlProgramUniform4uiEXT, callGlProgramUniform4uiEXT);
    interpreter->registerFunction(Ids::GlProgramUniform4uivEXT, callGlProgramUniform4uivEXT);
    interpreter->registerFunction(Ids::GlProgramUniformHandleui64NV,
                                  callGlProgramUniformHandleui64NV);
    interpreter->registerFunction(Ids::GlProgramUniformHandleui64vNV,
                                  callGlProgramUniformHandleui64vNV);
    interpreter->registerFunction(Ids::GlProgramUniformMatrix2fvEXT,
                                  callGlProgramUniformMatrix2fvEXT);
    interpreter->registerFunction(Ids::GlProgramUniformMatrix2x3fvEXT,
                                  callGlProgramUniformMatrix2x3fvEXT);
    interpreter->registerFunction(Ids::GlProgramUniformMatrix2x4fvEXT,
                                  callGlProgramUniformMatrix2x4fvEXT);
    interpreter->registerFunction(Ids::GlProgramUniformMatrix3fvEXT,
                                  callGlProgramUniformMatrix3fvEXT);
    interpreter->registerFunction(Ids::GlProgramUniformMatrix3x2fvEXT,
                                  callGlProgramUniformMatrix3x2fvEXT);
    interpreter->registerFunction(Ids::GlProgramUniformMatrix3x4fvEXT,
                                  callGlProgramUniformMatrix3x4fvEXT);
    interpreter->registerFunction(Ids::GlProgramUniformMatrix4fvEXT,
                                  callGlProgramUniformMatrix4fvEXT);
    interpreter->registerFunction(Ids::GlProgramUniformMatrix4x2fvEXT,
                                  callGlProgramUniformMatrix4x2fvEXT);
    interpreter->registerFunction(Ids::GlProgramUniformMatrix4x3fvEXT,
                                  callGlProgramUniformMatrix4x3fvEXT);
    interpreter->registerFunction(Ids::GlPushGroupMarkerEXT, callGlPushGroupMarkerEXT);
    interpreter->registerFunction(Ids::GlQueryCounterEXT, callGlQueryCounterEXT);
    interpreter->registerFunction(Ids::GlRasterSamplesEXT, callGlRasterSamplesEXT);
    interpreter->registerFunction(Ids::GlReadBufferIndexedEXT, callGlReadBufferIndexedEXT);
    interpreter->registerFunction(Ids::GlReadBufferNV, callGlReadBufferNV);
    interpreter->registerFunction(Ids::GlReadnPixelsEXT, callGlReadnPixelsEXT);
    interpreter->registerFunction(Ids::GlReadnPixelsKHR, callGlReadnPixelsKHR);
    interpreter->registerFunction(Ids::GlRenderbufferStorageMultisampleANGLE,
                                  callGlRenderbufferStorageMultisampleANGLE);
    interpreter->registerFunction(Ids::GlRenderbufferStorageMultisampleAPPLE,
                                  callGlRenderbufferStorageMultisampleAPPLE);
    interpreter->registerFunction(Ids::GlRenderbufferStorageMultisampleEXT,
                                  callGlRenderbufferStorageMultisampleEXT);
    interpreter->registerFunction(Ids::GlRenderbufferStorageMultisampleIMG,
                                  callGlRenderbufferStorageMultisampleIMG);
    interpreter->registerFunction(Ids::GlRenderbufferStorageMultisampleNV,
                                  callGlRenderbufferStorageMultisampleNV);
    interpreter->registerFunction(Ids::GlResolveDepthValuesNV, callGlResolveDepthValuesNV);
    interpreter->registerFunction(Ids::GlResolveMultisampleFramebufferAPPLE,
                                  callGlResolveMultisampleFramebufferAPPLE);
    interpreter->registerFunction(Ids::GlSamplerParameterIivOES, callGlSamplerParameterIivOES);
    interpreter->registerFunction(Ids::GlSamplerParameterIuivOES, callGlSamplerParameterIuivOES);
    interpreter->registerFunction(Ids::GlScissorArrayvNV, callGlScissorArrayvNV);
    interpreter->registerFunction(Ids::GlScissorIndexedNV, callGlScissorIndexedNV);
    interpreter->registerFunction(Ids::GlScissorIndexedvNV, callGlScissorIndexedvNV);
    interpreter->registerFunction(Ids::GlSelectPerfMonitorCountersAMD,
                                  callGlSelectPerfMonitorCountersAMD);
    interpreter->registerFunction(Ids::GlSetFenceNV, callGlSetFenceNV);
    interpreter->registerFunction(Ids::GlStartTilingQCOM, callGlStartTilingQCOM);
    interpreter->registerFunction(Ids::GlStencilFillPathInstancedNV,
                                  callGlStencilFillPathInstancedNV);
    interpreter->registerFunction(Ids::GlStencilFillPathNV, callGlStencilFillPathNV);
    interpreter->registerFunction(Ids::GlStencilStrokePathInstancedNV,
                                  callGlStencilStrokePathInstancedNV);
    interpreter->registerFunction(Ids::GlStencilStrokePathNV, callGlStencilStrokePathNV);
    interpreter->registerFunction(Ids::GlStencilThenCoverFillPathInstancedNV,
                                  callGlStencilThenCoverFillPathInstancedNV);
    interpreter->registerFunction(Ids::GlStencilThenCoverFillPathNV,
                                  callGlStencilThenCoverFillPathNV);
    interpreter->registerFunction(Ids::GlStencilThenCoverStrokePathInstancedNV,
                                  callGlStencilThenCoverStrokePathInstancedNV);
    interpreter->registerFunction(Ids::GlStencilThenCoverStrokePathNV,
                                  callGlStencilThenCoverStrokePathNV);
    interpreter->registerFunction(Ids::GlSubpixelPrecisionBiasNV, callGlSubpixelPrecisionBiasNV);
    interpreter->registerFunction(Ids::GlTestFenceNV, callGlTestFenceNV);
    interpreter->registerFunction(Ids::GlTexBufferOES, callGlTexBufferOES);
    interpreter->registerFunction(Ids::GlTexBufferRangeOES, callGlTexBufferRangeOES);
    interpreter->registerFunction(Ids::GlTexImage3DOES, callGlTexImage3DOES);
    interpreter->registerFunction(Ids::GlTexPageCommitmentARB, callGlTexPageCommitmentARB);
    interpreter->registerFunction(Ids::GlTexParameterIivOES, callGlTexParameterIivOES);
    interpreter->registerFunction(Ids::GlTexParameterIuivOES, callGlTexParameterIuivOES);
    interpreter->registerFunction(Ids::GlTexStorage1DEXT, callGlTexStorage1DEXT);
    interpreter->registerFunction(Ids::GlTexStorage2DEXT, callGlTexStorage2DEXT);
    interpreter->registerFunction(Ids::GlTexStorage3DEXT, callGlTexStorage3DEXT);
    interpreter->registerFunction(Ids::GlTexSubImage3DOES, callGlTexSubImage3DOES);
    interpreter->registerFunction(Ids::GlTextureStorage1DEXT, callGlTextureStorage1DEXT);
    interpreter->registerFunction(Ids::GlTextureStorage2DEXT, callGlTextureStorage2DEXT);
    interpreter->registerFunction(Ids::GlTextureStorage3DEXT, callGlTextureStorage3DEXT);
    interpreter->registerFunction(Ids::GlTextureViewEXT, callGlTextureViewEXT);
    interpreter->registerFunction(Ids::GlTextureViewOES, callGlTextureViewOES);
    interpreter->registerFunction(Ids::GlTransformPathNV, callGlTransformPathNV);
    interpreter->registerFunction(Ids::GlUniformHandleui64NV, callGlUniformHandleui64NV);
    interpreter->registerFunction(Ids::GlUniformHandleui64vNV, callGlUniformHandleui64vNV);
    interpreter->registerFunction(Ids::GlUniformMatrix2x3fvNV, callGlUniformMatrix2x3fvNV);
    interpreter->registerFunction(Ids::GlUniformMatrix2x4fvNV, callGlUniformMatrix2x4fvNV);
    interpreter->registerFunction(Ids::GlUniformMatrix3x2fvNV, callGlUniformMatrix3x2fvNV);
    interpreter->registerFunction(Ids::GlUniformMatrix3x4fvNV, callGlUniformMatrix3x4fvNV);
    interpreter->registerFunction(Ids::GlUniformMatrix4x2fvNV, callGlUniformMatrix4x2fvNV);
    interpreter->registerFunction(Ids::GlUniformMatrix4x3fvNV, callGlUniformMatrix4x3fvNV);
    interpreter->registerFunction(Ids::GlUnmapBufferOES, callGlUnmapBufferOES);
    interpreter->registerFunction(Ids::GlUseProgramStagesEXT, callGlUseProgramStagesEXT);
    interpreter->registerFunction(Ids::GlValidateProgramPipelineEXT,
                                  callGlValidateProgramPipelineEXT);
    interpreter->registerFunction(Ids::GlVertexAttribDivisorANGLE, callGlVertexAttribDivisorANGLE);
    interpreter->registerFunction(Ids::GlVertexAttribDivisorEXT, callGlVertexAttribDivisorEXT);
    interpreter->registerFunction(Ids::GlVertexAttribDivisorNV, callGlVertexAttribDivisorNV);
    interpreter->registerFunction(Ids::GlViewportArrayvNV, callGlViewportArrayvNV);
    interpreter->registerFunction(Ids::GlViewportIndexedfNV, callGlViewportIndexedfNV);
    interpreter->registerFunction(Ids::GlViewportIndexedfvNV, callGlViewportIndexedfvNV);
    interpreter->registerFunction(Ids::GlWaitSyncAPPLE, callGlWaitSyncAPPLE);
    interpreter->registerFunction(Ids::GlWeightPathsNV, callGlWeightPathsNV);
    interpreter->registerFunction(Ids::GlBlendBarrier, callGlBlendBarrier);
    interpreter->registerFunction(Ids::GlBlendColor, callGlBlendColor);
    interpreter->registerFunction(Ids::GlBlendEquation, callGlBlendEquation);
    interpreter->registerFunction(Ids::GlBlendEquationSeparate, callGlBlendEquationSeparate);
    interpreter->registerFunction(Ids::GlBlendEquationSeparatei, callGlBlendEquationSeparatei);
    interpreter->registerFunction(Ids::GlBlendEquationi, callGlBlendEquationi);
    interpreter->registerFunction(Ids::GlBlendFunc, callGlBlendFunc);
    interpreter->registerFunction(Ids::GlBlendFuncSeparate, callGlBlendFuncSeparate);
    interpreter->registerFunction(Ids::GlBlendFuncSeparatei, callGlBlendFuncSeparatei);
    interpreter->registerFunction(Ids::GlBlendFunci, callGlBlendFunci);
    interpreter->registerFunction(Ids::GlDepthFunc, callGlDepthFunc);
    interpreter->registerFunction(Ids::GlSampleCoverage, callGlSampleCoverage);
    interpreter->registerFunction(Ids::GlSampleMaski, callGlSampleMaski);
    interpreter->registerFunction(Ids::GlScissor, callGlScissor);
    interpreter->registerFunction(Ids::GlStencilFunc, callGlStencilFunc);
    interpreter->registerFunction(Ids::GlStencilFuncSeparate, callGlStencilFuncSeparate);
    interpreter->registerFunction(Ids::GlStencilOp, callGlStencilOp);
    interpreter->registerFunction(Ids::GlStencilOpSeparate, callGlStencilOpSeparate);
    interpreter->registerFunction(Ids::GlBindFramebuffer, callGlBindFramebuffer);
    interpreter->registerFunction(Ids::GlBindRenderbuffer, callGlBindRenderbuffer);
    interpreter->registerFunction(Ids::GlBlitFramebuffer, callGlBlitFramebuffer);
    interpreter->registerFunction(Ids::GlCheckFramebufferStatus, callGlCheckFramebufferStatus);
    interpreter->registerFunction(Ids::GlClear, callGlClear);
    interpreter->registerFunction(Ids::GlClearBufferfi, callGlClearBufferfi);
    interpreter->registerFunction(Ids::GlClearBufferfv, callGlClearBufferfv);
    interpreter->registerFunction(Ids::GlClearBufferiv, callGlClearBufferiv);
    interpreter->registerFunction(Ids::GlClearBufferuiv, callGlClearBufferuiv);
    interpreter->registerFunction(Ids::GlClearColor, callGlClearColor);
    interpreter->registerFunction(Ids::GlClearDepthf, callGlClearDepthf);
    interpreter->registerFunction(Ids::GlClearStencil, callGlClearStencil);
    interpreter->registerFunction(Ids::GlColorMask, callGlColorMask);
    interpreter->registerFunction(Ids::GlColorMaski, callGlColorMaski);
    interpreter->registerFunction(Ids::GlDeleteFramebuffers, callGlDeleteFramebuffers);
    interpreter->registerFunction(Ids::GlDeleteRenderbuffers, callGlDeleteRenderbuffers);
    interpreter->registerFunction(Ids::GlDepthMask, callGlDepthMask);
    interpreter->registerFunction(Ids::GlDrawBuffers, callGlDrawBuffers);
    interpreter->registerFunction(Ids::GlFramebufferParameteri, callGlFramebufferParameteri);
    interpreter->registerFunction(Ids::GlFramebufferRenderbuffer, callGlFramebufferRenderbuffer);
    interpreter->registerFunction(Ids::GlFramebufferTexture, callGlFramebufferTexture);
    interpreter->registerFunction(Ids::GlFramebufferTexture2D, callGlFramebufferTexture2D);
    interpreter->registerFunction(Ids::GlFramebufferTextureLayer, callGlFramebufferTextureLayer);
    interpreter->registerFunction(Ids::GlGenFramebuffers, callGlGenFramebuffers);
    interpreter->registerFunction(Ids::GlGenRenderbuffers, callGlGenRenderbuffers);
    interpreter->registerFunction(Ids::GlGetFramebufferAttachmentParameteriv,
                                  callGlGetFramebufferAttachmentParameteriv);
    interpreter->registerFunction(Ids::GlGetFramebufferParameteriv,
                                  callGlGetFramebufferParameteriv);
    interpreter->registerFunction(Ids::GlGetRenderbufferParameteriv,
                                  callGlGetRenderbufferParameteriv);
    interpreter->registerFunction(Ids::GlInvalidateFramebuffer, callGlInvalidateFramebuffer);
    interpreter->registerFunction(Ids::GlInvalidateSubFramebuffer, callGlInvalidateSubFramebuffer);
    interpreter->registerFunction(Ids::GlIsFramebuffer, callGlIsFramebuffer);
    interpreter->registerFunction(Ids::GlIsRenderbuffer, callGlIsRenderbuffer);
    interpreter->registerFunction(Ids::GlReadBuffer, callGlReadBuffer);
    interpreter->registerFunction(Ids::GlReadPixels, callGlReadPixels);
    interpreter->registerFunction(Ids::GlReadnPixels, callGlReadnPixels);
    interpreter->registerFunction(Ids::GlRenderbufferStorage, callGlRenderbufferStorage);
    interpreter->registerFunction(Ids::GlRenderbufferStorageMultisample,
                                  callGlRenderbufferStorageMultisample);
    interpreter->registerFunction(Ids::GlStencilMask, callGlStencilMask);
    interpreter->registerFunction(Ids::GlStencilMaskSeparate, callGlStencilMaskSeparate);
    interpreter->registerFunction(Ids::GlDisable, callGlDisable);
    interpreter->registerFunction(Ids::GlDisablei, callGlDisablei);
    interpreter->registerFunction(Ids::GlEnable, callGlEnable);
    interpreter->registerFunction(Ids::GlEnablei, callGlEnablei);
    interpreter->registerFunction(Ids::GlFinish, callGlFinish);
    interpreter->registerFunction(Ids::GlFlush, callGlFlush);
    interpreter->registerFunction(Ids::GlFlushMappedBufferRange, callGlFlushMappedBufferRange);
    interpreter->registerFunction(Ids::GlGetError, callGlGetError);
    interpreter->registerFunction(Ids::GlGetGraphicsResetStatus, callGlGetGraphicsResetStatus);
    interpreter->registerFunction(Ids::GlHint, callGlHint);
    interpreter->registerFunction(Ids::GlActiveShaderProgram, callGlActiveShaderProgram);
    interpreter->registerFunction(Ids::GlAttachShader, callGlAttachShader);
    interpreter->registerFunction(Ids::GlBindAttribLocation, callGlBindAttribLocation);
    interpreter->registerFunction(Ids::GlBindProgramPipeline, callGlBindProgramPipeline);
    interpreter->registerFunction(Ids::GlCompileShader, callGlCompileShader);
    interpreter->registerFunction(Ids::GlCreateProgram, callGlCreateProgram);
    interpreter->registerFunction(Ids::GlCreateShader, callGlCreateShader);
    interpreter->registerFunction(Ids::GlCreateShaderProgramv, callGlCreateShaderProgramv);
    interpreter->registerFunction(Ids::GlDeleteProgram, callGlDeleteProgram);
    interpreter->registerFunction(Ids::GlDeleteProgramPipelines, callGlDeleteProgramPipelines);
    interpreter->registerFunction(Ids::GlDeleteShader, callGlDeleteShader);
    interpreter->registerFunction(Ids::GlDetachShader, callGlDetachShader);
    interpreter->registerFunction(Ids::GlDispatchCompute, callGlDispatchCompute);
    interpreter->registerFunction(Ids::GlDispatchComputeIndirect, callGlDispatchComputeIndirect);
    interpreter->registerFunction(Ids::GlGenProgramPipelines, callGlGenProgramPipelines);
    interpreter->registerFunction(Ids::GlGetActiveAttrib, callGlGetActiveAttrib);
    interpreter->registerFunction(Ids::GlGetActiveUniform, callGlGetActiveUniform);
    interpreter->registerFunction(Ids::GlGetActiveUniformBlockName,
                                  callGlGetActiveUniformBlockName);
    interpreter->registerFunction(Ids::GlGetActiveUniformBlockiv, callGlGetActiveUniformBlockiv);
    interpreter->registerFunction(Ids::GlGetActiveUniformsiv, callGlGetActiveUniformsiv);
    interpreter->registerFunction(Ids::GlGetAttachedShaders, callGlGetAttachedShaders);
    interpreter->registerFunction(Ids::GlGetAttribLocation, callGlGetAttribLocation);
    interpreter->registerFunction(Ids::GlGetFragDataLocation, callGlGetFragDataLocation);
    interpreter->registerFunction(Ids::GlGetProgramBinary, callGlGetProgramBinary);
    interpreter->registerFunction(Ids::GlGetProgramInfoLog, callGlGetProgramInfoLog);
    interpreter->registerFunction(Ids::GlGetProgramInterfaceiv, callGlGetProgramInterfaceiv);
    interpreter->registerFunction(Ids::GlGetProgramPipelineInfoLog,
                                  callGlGetProgramPipelineInfoLog);
    interpreter->registerFunction(Ids::GlGetProgramPipelineiv, callGlGetProgramPipelineiv);
    interpreter->registerFunction(Ids::GlGetProgramResourceIndex, callGlGetProgramResourceIndex);
    interpreter->registerFunction(Ids::GlGetProgramResourceLocation,
                                  callGlGetProgramResourceLocation);
    interpreter->registerFunction(Ids::GlGetProgramResourceName, callGlGetProgramResourceName);
    interpreter->registerFunction(Ids::GlGetProgramResourceiv, callGlGetProgramResourceiv);
    interpreter->registerFunction(Ids::GlGetProgramiv, callGlGetProgramiv);
    interpreter->registerFunction(Ids::GlGetShaderInfoLog, callGlGetShaderInfoLog);
    interpreter->registerFunction(Ids::GlGetShaderPrecisionFormat, callGlGetShaderPrecisionFormat);
    interpreter->registerFunction(Ids::GlGetShaderSource, callGlGetShaderSource);
    interpreter->registerFunction(Ids::GlGetShaderiv, callGlGetShaderiv);
    interpreter->registerFunction(Ids::GlGetUniformBlockIndex, callGlGetUniformBlockIndex);
    interpreter->registerFunction(Ids::GlGetUniformIndices, callGlGetUniformIndices);
    interpreter->registerFunction(Ids::GlGetUniformLocation, callGlGetUniformLocation);
    interpreter->registerFunction(Ids::GlGetUniformfv, callGlGetUniformfv);
    interpreter->registerFunction(Ids::GlGetUniformiv, callGlGetUniformiv);
    interpreter->registerFunction(Ids::GlGetUniformuiv, callGlGetUniformuiv);
    interpreter->registerFunction(Ids::GlGetnUniformfv, callGlGetnUniformfv);
    interpreter->registerFunction(Ids::GlGetnUniformiv, callGlGetnUniformiv);
    interpreter->registerFunction(Ids::GlGetnUniformuiv, callGlGetnUniformuiv);
    interpreter->registerFunction(Ids::GlIsProgram, callGlIsProgram);
    interpreter->registerFunction(Ids::GlIsProgramPipeline, callGlIsProgramPipeline);
    interpreter->registerFunction(Ids::GlIsShader, callGlIsShader);
    interpreter->registerFunction(Ids::GlLinkProgram, callGlLinkProgram);
    interpreter->registerFunction(Ids::GlMemoryBarrier, callGlMemoryBarrier);
    interpreter->registerFunction(Ids::GlMemoryBarrierByRegion, callGlMemoryBarrierByRegion);
    interpreter->registerFunction(Ids::GlProgramBinary, callGlProgramBinary);
    interpreter->registerFunction(Ids::GlProgramParameteri, callGlProgramParameteri);
    interpreter->registerFunction(Ids::GlProgramUniform1f, callGlProgramUniform1f);
    interpreter->registerFunction(Ids::GlProgramUniform1fv, callGlProgramUniform1fv);
    interpreter->registerFunction(Ids::GlProgramUniform1i, callGlProgramUniform1i);
    interpreter->registerFunction(Ids::GlProgramUniform1iv, callGlProgramUniform1iv);
    interpreter->registerFunction(Ids::GlProgramUniform1ui, callGlProgramUniform1ui);
    interpreter->registerFunction(Ids::GlProgramUniform1uiv, callGlProgramUniform1uiv);
    interpreter->registerFunction(Ids::GlProgramUniform2f, callGlProgramUniform2f);
    interpreter->registerFunction(Ids::GlProgramUniform2fv, callGlProgramUniform2fv);
    interpreter->registerFunction(Ids::GlProgramUniform2i, callGlProgramUniform2i);
    interpreter->registerFunction(Ids::GlProgramUniform2iv, callGlProgramUniform2iv);
    interpreter->registerFunction(Ids::GlProgramUniform2ui, callGlProgramUniform2ui);
    interpreter->registerFunction(Ids::GlProgramUniform2uiv, callGlProgramUniform2uiv);
    interpreter->registerFunction(Ids::GlProgramUniform3f, callGlProgramUniform3f);
    interpreter->registerFunction(Ids::GlProgramUniform3fv, callGlProgramUniform3fv);
    interpreter->registerFunction(Ids::GlProgramUniform3i, callGlProgramUniform3i);
    interpreter->registerFunction(Ids::GlProgramUniform3iv, callGlProgramUniform3iv);
    interpreter->registerFunction(Ids::GlProgramUniform3ui, callGlProgramUniform3ui);
    interpreter->registerFunction(Ids::GlProgramUniform3uiv, callGlProgramUniform3uiv);
    interpreter->registerFunction(Ids::GlProgramUniform4f, callGlProgramUniform4f);
    interpreter->registerFunction(Ids::GlProgramUniform4fv, callGlProgramUniform4fv);
    interpreter->registerFunction(Ids::GlProgramUniform4i, callGlProgramUniform4i);
    interpreter->registerFunction(Ids::GlProgramUniform4iv, callGlProgramUniform4iv);
    interpreter->registerFunction(Ids::GlProgramUniform4ui, callGlProgramUniform4ui);
    interpreter->registerFunction(Ids::GlProgramUniform4uiv, callGlProgramUniform4uiv);
    interpreter->registerFunction(Ids::GlProgramUniformMatrix2fv, callGlProgramUniformMatrix2fv);
    interpreter->registerFunction(Ids::GlProgramUniformMatrix2x3fv,
                                  callGlProgramUniformMatrix2x3fv);
    interpreter->registerFunction(Ids::GlProgramUniformMatrix2x4fv,
                                  callGlProgramUniformMatrix2x4fv);
    interpreter->registerFunction(Ids::GlProgramUniformMatrix3fv, callGlProgramUniformMatrix3fv);
    interpreter->registerFunction(Ids::GlProgramUniformMatrix3x2fv,
                                  callGlProgramUniformMatrix3x2fv);
    interpreter->registerFunction(Ids::GlProgramUniformMatrix3x4fv,
                                  callGlProgramUniformMatrix3x4fv);
    interpreter->registerFunction(Ids::GlProgramUniformMatrix4fv, callGlProgramUniformMatrix4fv);
    interpreter->registerFunction(Ids::GlProgramUniformMatrix4x2fv,
                                  callGlProgramUniformMatrix4x2fv);
    interpreter->registerFunction(Ids::GlProgramUniformMatrix4x3fv,
                                  callGlProgramUniformMatrix4x3fv);
    interpreter->registerFunction(Ids::GlReleaseShaderCompiler, callGlReleaseShaderCompiler);
    interpreter->registerFunction(Ids::GlShaderBinary, callGlShaderBinary);
    interpreter->registerFunction(Ids::GlShaderSource, callGlShaderSource);
    interpreter->registerFunction(Ids::GlUniform1f, callGlUniform1f);
    interpreter->registerFunction(Ids::GlUniform1fv, callGlUniform1fv);
    interpreter->registerFunction(Ids::GlUniform1i, callGlUniform1i);
    interpreter->registerFunction(Ids::GlUniform1iv, callGlUniform1iv);
    interpreter->registerFunction(Ids::GlUniform1ui, callGlUniform1ui);
    interpreter->registerFunction(Ids::GlUniform1uiv, callGlUniform1uiv);
    interpreter->registerFunction(Ids::GlUniform2f, callGlUniform2f);
    interpreter->registerFunction(Ids::GlUniform2fv, callGlUniform2fv);
    interpreter->registerFunction(Ids::GlUniform2i, callGlUniform2i);
    interpreter->registerFunction(Ids::GlUniform2iv, callGlUniform2iv);
    interpreter->registerFunction(Ids::GlUniform2ui, callGlUniform2ui);
    interpreter->registerFunction(Ids::GlUniform2uiv, callGlUniform2uiv);
    interpreter->registerFunction(Ids::GlUniform3f, callGlUniform3f);
    interpreter->registerFunction(Ids::GlUniform3fv, callGlUniform3fv);
    interpreter->registerFunction(Ids::GlUniform3i, callGlUniform3i);
    interpreter->registerFunction(Ids::GlUniform3iv, callGlUniform3iv);
    interpreter->registerFunction(Ids::GlUniform3ui, callGlUniform3ui);
    interpreter->registerFunction(Ids::GlUniform3uiv, callGlUniform3uiv);
    interpreter->registerFunction(Ids::GlUniform4f, callGlUniform4f);
    interpreter->registerFunction(Ids::GlUniform4fv, callGlUniform4fv);
    interpreter->registerFunction(Ids::GlUniform4i, callGlUniform4i);
    interpreter->registerFunction(Ids::GlUniform4iv, callGlUniform4iv);
    interpreter->registerFunction(Ids::GlUniform4ui, callGlUniform4ui);
    interpreter->registerFunction(Ids::GlUniform4uiv, callGlUniform4uiv);
    interpreter->registerFunction(Ids::GlUniformBlockBinding, callGlUniformBlockBinding);
    interpreter->registerFunction(Ids::GlUniformMatrix2fv, callGlUniformMatrix2fv);
    interpreter->registerFunction(Ids::GlUniformMatrix2x3fv, callGlUniformMatrix2x3fv);
    interpreter->registerFunction(Ids::GlUniformMatrix2x4fv, callGlUniformMatrix2x4fv);
    interpreter->registerFunction(Ids::GlUniformMatrix3fv, callGlUniformMatrix3fv);
    interpreter->registerFunction(Ids::GlUniformMatrix3x2fv, callGlUniformMatrix3x2fv);
    interpreter->registerFunction(Ids::GlUniformMatrix3x4fv, callGlUniformMatrix3x4fv);
    interpreter->registerFunction(Ids::GlUniformMatrix4fv, callGlUniformMatrix4fv);
    interpreter->registerFunction(Ids::GlUniformMatrix4x2fv, callGlUniformMatrix4x2fv);
    interpreter->registerFunction(Ids::GlUniformMatrix4x3fv, callGlUniformMatrix4x3fv);
    interpreter->registerFunction(Ids::GlUseProgram, callGlUseProgram);
    interpreter->registerFunction(Ids::GlUseProgramStages, callGlUseProgramStages);
    interpreter->registerFunction(Ids::GlValidateProgram, callGlValidateProgram);
    interpreter->registerFunction(Ids::GlValidateProgramPipeline, callGlValidateProgramPipeline);
    interpreter->registerFunction(Ids::GlCullFace, callGlCullFace);
    interpreter->registerFunction(Ids::GlDepthRangef, callGlDepthRangef);
    interpreter->registerFunction(Ids::GlFrontFace, callGlFrontFace);
    interpreter->registerFunction(Ids::GlGetMultisamplefv, callGlGetMultisamplefv);
    interpreter->registerFunction(Ids::GlLineWidth, callGlLineWidth);
    interpreter->registerFunction(Ids::GlMinSampleShading, callGlMinSampleShading);
    interpreter->registerFunction(Ids::GlPolygonOffset, callGlPolygonOffset);
    interpreter->registerFunction(Ids::GlViewport, callGlViewport);
    interpreter->registerFunction(Ids::GlGetBooleaniV, callGlGetBooleaniV);
    interpreter->registerFunction(Ids::GlGetBooleanv, callGlGetBooleanv);
    interpreter->registerFunction(Ids::GlGetFloatv, callGlGetFloatv);
    interpreter->registerFunction(Ids::GlGetInteger64iV, callGlGetInteger64iV);
    interpreter->registerFunction(Ids::GlGetInteger64v, callGlGetInteger64v);
    interpreter->registerFunction(Ids::GlGetIntegeriV, callGlGetIntegeriV);
    interpreter->registerFunction(Ids::GlGetIntegerv, callGlGetIntegerv);
    interpreter->registerFunction(Ids::GlGetInternalformativ, callGlGetInternalformativ);
    interpreter->registerFunction(Ids::GlGetString, callGlGetString);
    interpreter->registerFunction(Ids::GlGetStringi, callGlGetStringi);
    interpreter->registerFunction(Ids::GlIsEnabled, callGlIsEnabled);
    interpreter->registerFunction(Ids::GlIsEnabledi, callGlIsEnabledi);
    interpreter->registerFunction(Ids::GlClientWaitSync, callGlClientWaitSync);
    interpreter->registerFunction(Ids::GlDeleteSync, callGlDeleteSync);
    interpreter->registerFunction(Ids::GlFenceSync, callGlFenceSync);
    interpreter->registerFunction(Ids::GlGetSynciv, callGlGetSynciv);
    interpreter->registerFunction(Ids::GlIsSync, callGlIsSync);
    interpreter->registerFunction(Ids::GlWaitSync, callGlWaitSync);
    interpreter->registerFunction(Ids::GlActiveTexture, callGlActiveTexture);
    interpreter->registerFunction(Ids::GlBindImageTexture, callGlBindImageTexture);
    interpreter->registerFunction(Ids::GlBindSampler, callGlBindSampler);
    interpreter->registerFunction(Ids::GlBindTexture, callGlBindTexture);
    interpreter->registerFunction(Ids::GlCompressedTexImage2D, callGlCompressedTexImage2D);
    interpreter->registerFunction(Ids::GlCompressedTexImage3D, callGlCompressedTexImage3D);
    interpreter->registerFunction(Ids::GlCompressedTexSubImage2D, callGlCompressedTexSubImage2D);
    interpreter->registerFunction(Ids::GlCompressedTexSubImage3D, callGlCompressedTexSubImage3D);
    interpreter->registerFunction(Ids::GlCopyImageSubData, callGlCopyImageSubData);
    interpreter->registerFunction(Ids::GlCopyTexImage2D, callGlCopyTexImage2D);
    interpreter->registerFunction(Ids::GlCopyTexSubImage2D, callGlCopyTexSubImage2D);
    interpreter->registerFunction(Ids::GlCopyTexSubImage3D, callGlCopyTexSubImage3D);
    interpreter->registerFunction(Ids::GlDeleteSamplers, callGlDeleteSamplers);
    interpreter->registerFunction(Ids::GlDeleteTextures, callGlDeleteTextures);
    interpreter->registerFunction(Ids::GlGenSamplers, callGlGenSamplers);
    interpreter->registerFunction(Ids::GlGenTextures, callGlGenTextures);
    interpreter->registerFunction(Ids::GlGenerateMipmap, callGlGenerateMipmap);
    interpreter->registerFunction(Ids::GlGetSamplerParameterIiv, callGlGetSamplerParameterIiv);
    interpreter->registerFunction(Ids::GlGetSamplerParameterIuiv, callGlGetSamplerParameterIuiv);
    interpreter->registerFunction(Ids::GlGetSamplerParameterfv, callGlGetSamplerParameterfv);
    interpreter->registerFunction(Ids::GlGetSamplerParameteriv, callGlGetSamplerParameteriv);
    interpreter->registerFunction(Ids::GlGetTexLevelParameterfv, callGlGetTexLevelParameterfv);
    interpreter->registerFunction(Ids::GlGetTexLevelParameteriv, callGlGetTexLevelParameteriv);
    interpreter->registerFunction(Ids::GlGetTexParameterIiv, callGlGetTexParameterIiv);
    interpreter->registerFunction(Ids::GlGetTexParameterIuiv, callGlGetTexParameterIuiv);
    interpreter->registerFunction(Ids::GlGetTexParameterfv, callGlGetTexParameterfv);
    interpreter->registerFunction(Ids::GlGetTexParameteriv, callGlGetTexParameteriv);
    interpreter->registerFunction(Ids::GlIsSampler, callGlIsSampler);
    interpreter->registerFunction(Ids::GlIsTexture, callGlIsTexture);
    interpreter->registerFunction(Ids::GlPixelStorei, callGlPixelStorei);
    interpreter->registerFunction(Ids::GlSamplerParameterIiv, callGlSamplerParameterIiv);
    interpreter->registerFunction(Ids::GlSamplerParameterIuiv, callGlSamplerParameterIuiv);
    interpreter->registerFunction(Ids::GlSamplerParameterf, callGlSamplerParameterf);
    interpreter->registerFunction(Ids::GlSamplerParameterfv, callGlSamplerParameterfv);
    interpreter->registerFunction(Ids::GlSamplerParameteri, callGlSamplerParameteri);
    interpreter->registerFunction(Ids::GlSamplerParameteriv, callGlSamplerParameteriv);
    interpreter->registerFunction(Ids::GlTexBuffer, callGlTexBuffer);
    interpreter->registerFunction(Ids::GlTexBufferRange, callGlTexBufferRange);
    interpreter->registerFunction(Ids::GlTexImage2D, callGlTexImage2D);
    interpreter->registerFunction(Ids::GlTexImage3D, callGlTexImage3D);
    interpreter->registerFunction(Ids::GlTexParameterIiv, callGlTexParameterIiv);
    interpreter->registerFunction(Ids::GlTexParameterIuiv, callGlTexParameterIuiv);
    interpreter->registerFunction(Ids::GlTexParameterf, callGlTexParameterf);
    interpreter->registerFunction(Ids::GlTexParameterfv, callGlTexParameterfv);
    interpreter->registerFunction(Ids::GlTexParameteri, callGlTexParameteri);
    interpreter->registerFunction(Ids::GlTexParameteriv, callGlTexParameteriv);
    interpreter->registerFunction(Ids::GlTexStorage2D, callGlTexStorage2D);
    interpreter->registerFunction(Ids::GlTexStorage2DMultisample, callGlTexStorage2DMultisample);
    interpreter->registerFunction(Ids::GlTexStorage3D, callGlTexStorage3D);
    interpreter->registerFunction(Ids::GlTexStorage3DMultisample, callGlTexStorage3DMultisample);
    interpreter->registerFunction(Ids::GlTexSubImage2D, callGlTexSubImage2D);
    interpreter->registerFunction(Ids::GlTexSubImage3D, callGlTexSubImage3D);
    interpreter->registerFunction(Ids::GlBeginTransformFeedback, callGlBeginTransformFeedback);
    interpreter->registerFunction(Ids::GlBindTransformFeedback, callGlBindTransformFeedback);
    interpreter->registerFunction(Ids::GlDeleteTransformFeedbacks, callGlDeleteTransformFeedbacks);
    interpreter->registerFunction(Ids::GlEndTransformFeedback, callGlEndTransformFeedback);
    interpreter->registerFunction(Ids::GlGenTransformFeedbacks, callGlGenTransformFeedbacks);
    interpreter->registerFunction(Ids::GlGetTransformFeedbackVarying,
                                  callGlGetTransformFeedbackVarying);
    interpreter->registerFunction(Ids::GlIsTransformFeedback, callGlIsTransformFeedback);
    interpreter->registerFunction(Ids::GlPauseTransformFeedback, callGlPauseTransformFeedback);
    interpreter->registerFunction(Ids::GlResumeTransformFeedback, callGlResumeTransformFeedback);
    interpreter->registerFunction(Ids::GlTransformFeedbackVaryings,
                                  callGlTransformFeedbackVaryings);
    interpreter->registerFunction(Ids::GlBindVertexArray, callGlBindVertexArray);
    interpreter->registerFunction(Ids::GlBindVertexBuffer, callGlBindVertexBuffer);
    interpreter->registerFunction(Ids::GlDeleteVertexArrays, callGlDeleteVertexArrays);
    interpreter->registerFunction(Ids::GlDisableVertexAttribArray, callGlDisableVertexAttribArray);
    interpreter->registerFunction(Ids::GlEnableVertexAttribArray, callGlEnableVertexAttribArray);
    interpreter->registerFunction(Ids::GlGenVertexArrays, callGlGenVertexArrays);
    interpreter->registerFunction(Ids::GlGetVertexAttribIiv, callGlGetVertexAttribIiv);
    interpreter->registerFunction(Ids::GlGetVertexAttribIuiv, callGlGetVertexAttribIuiv);
    interpreter->registerFunction(Ids::GlGetVertexAttribPointerv, callGlGetVertexAttribPointerv);
    interpreter->registerFunction(Ids::GlGetVertexAttribfv, callGlGetVertexAttribfv);
    interpreter->registerFunction(Ids::GlGetVertexAttribiv, callGlGetVertexAttribiv);
    interpreter->registerFunction(Ids::GlIsVertexArray, callGlIsVertexArray);
    interpreter->registerFunction(Ids::GlVertexAttrib1f, callGlVertexAttrib1f);
    interpreter->registerFunction(Ids::GlVertexAttrib1fv, callGlVertexAttrib1fv);
    interpreter->registerFunction(Ids::GlVertexAttrib2f, callGlVertexAttrib2f);
    interpreter->registerFunction(Ids::GlVertexAttrib2fv, callGlVertexAttrib2fv);
    interpreter->registerFunction(Ids::GlVertexAttrib3f, callGlVertexAttrib3f);
    interpreter->registerFunction(Ids::GlVertexAttrib3fv, callGlVertexAttrib3fv);
    interpreter->registerFunction(Ids::GlVertexAttrib4f, callGlVertexAttrib4f);
    interpreter->registerFunction(Ids::GlVertexAttrib4fv, callGlVertexAttrib4fv);
    interpreter->registerFunction(Ids::GlVertexAttribBinding, callGlVertexAttribBinding);
    interpreter->registerFunction(Ids::GlVertexAttribDivisor, callGlVertexAttribDivisor);
    interpreter->registerFunction(Ids::GlVertexAttribFormat, callGlVertexAttribFormat);
    interpreter->registerFunction(Ids::GlVertexAttribI4i, callGlVertexAttribI4i);
    interpreter->registerFunction(Ids::GlVertexAttribI4iv, callGlVertexAttribI4iv);
    interpreter->registerFunction(Ids::GlVertexAttribI4ui, callGlVertexAttribI4ui);
    interpreter->registerFunction(Ids::GlVertexAttribI4uiv, callGlVertexAttribI4uiv);
    interpreter->registerFunction(Ids::GlVertexAttribIFormat, callGlVertexAttribIFormat);
    interpreter->registerFunction(Ids::GlVertexAttribIPointer, callGlVertexAttribIPointer);
    interpreter->registerFunction(Ids::GlVertexAttribPointer, callGlVertexAttribPointer);
    interpreter->registerFunction(Ids::GlVertexBindingDivisor, callGlVertexBindingDivisor);
    interpreter->registerFunction(Ids::EglInitialize, callEglInitialize);
    interpreter->registerFunction(Ids::EglCreateContext, callEglCreateContext);
    interpreter->registerFunction(Ids::EglMakeCurrent, callEglMakeCurrent);
    interpreter->registerFunction(Ids::EglSwapBuffers, callEglSwapBuffers);
    interpreter->registerFunction(Ids::EglQuerySurface, callEglQuerySurface);
    interpreter->registerFunction(Ids::GlXCreateContext, callGlXCreateContext);
    interpreter->registerFunction(Ids::GlXCreateNewContext, callGlXCreateNewContext);
    interpreter->registerFunction(Ids::GlXMakeContextCurrent, callGlXMakeContextCurrent);
    interpreter->registerFunction(Ids::GlXMakeCurrent, callGlXMakeCurrent);
    interpreter->registerFunction(Ids::GlXSwapBuffers, callGlXSwapBuffers);
    interpreter->registerFunction(Ids::GlXQueryDrawable, callGlXQueryDrawable);
    interpreter->registerFunction(Ids::WglCreateContext, callWglCreateContext);
    interpreter->registerFunction(Ids::WglCreateContextAttribsARB, callWglCreateContextAttribsARB);
    interpreter->registerFunction(Ids::WglMakeCurrent, callWglMakeCurrent);
    interpreter->registerFunction(Ids::WglSwapBuffers, callWglSwapBuffers);
    interpreter->registerFunction(Ids::CGLCreateContext, callCGLCreateContext);
    interpreter->registerFunction(Ids::CGLSetCurrentContext, callCGLSetCurrentContext);
    interpreter->registerFunction(Ids::CGLGetSurface, callCGLGetSurface);
    interpreter->registerFunction(Ids::CGSGetSurfaceBounds, callCGSGetSurfaceBounds);
    interpreter->registerFunction(Ids::CGLFlushDrawable, callCGLFlushDrawable);
    interpreter->registerFunction(Ids::GlGetQueryObjecti64v, callGlGetQueryObjecti64v);
    interpreter->registerFunction(Ids::GlGetQueryObjectui64v, callGlGetQueryObjectui64v);
}
void Initialize() {
    glBlendBarrierKHR = reinterpret_cast<PFNGLBLENDBARRIERKHR>(
            gapic::GetGfxProcAddress("glBlendBarrierKHR", false));
    glBlendEquationSeparateiEXT = reinterpret_cast<PFNGLBLENDEQUATIONSEPARATEIEXT>(
            gapic::GetGfxProcAddress("glBlendEquationSeparateiEXT", false));
    glBlendEquationiEXT = reinterpret_cast<PFNGLBLENDEQUATIONIEXT>(
            gapic::GetGfxProcAddress("glBlendEquationiEXT", false));
    glBlendFuncSeparateiEXT = reinterpret_cast<PFNGLBLENDFUNCSEPARATEIEXT>(
            gapic::GetGfxProcAddress("glBlendFuncSeparateiEXT", false));
    glBlendFunciEXT = reinterpret_cast<PFNGLBLENDFUNCIEXT>(
            gapic::GetGfxProcAddress("glBlendFunciEXT", false));
    glColorMaskiEXT = reinterpret_cast<PFNGLCOLORMASKIEXT>(
            gapic::GetGfxProcAddress("glColorMaskiEXT", false));
    glCopyImageSubDataEXT = reinterpret_cast<PFNGLCOPYIMAGESUBDATAEXT>(
            gapic::GetGfxProcAddress("glCopyImageSubDataEXT", false));
    glDebugMessageCallbackKHR = reinterpret_cast<PFNGLDEBUGMESSAGECALLBACKKHR>(
            gapic::GetGfxProcAddress("glDebugMessageCallbackKHR", false));
    glDebugMessageControlKHR = reinterpret_cast<PFNGLDEBUGMESSAGECONTROLKHR>(
            gapic::GetGfxProcAddress("glDebugMessageControlKHR", false));
    glDebugMessageInsertKHR = reinterpret_cast<PFNGLDEBUGMESSAGEINSERTKHR>(
            gapic::GetGfxProcAddress("glDebugMessageInsertKHR", false));
    glDisableiEXT =
            reinterpret_cast<PFNGLDISABLEIEXT>(gapic::GetGfxProcAddress("glDisableiEXT", false));
    glEnableiEXT =
            reinterpret_cast<PFNGLENABLEIEXT>(gapic::GetGfxProcAddress("glEnableiEXT", false));
    glFramebufferTextureEXT = reinterpret_cast<PFNGLFRAMEBUFFERTEXTUREEXT>(
            gapic::GetGfxProcAddress("glFramebufferTextureEXT", false));
    glGetDebugMessageLogKHR = reinterpret_cast<PFNGLGETDEBUGMESSAGELOGKHR>(
            gapic::GetGfxProcAddress("glGetDebugMessageLogKHR", false));
    glGetObjectLabelKHR = reinterpret_cast<PFNGLGETOBJECTLABELKHR>(
            gapic::GetGfxProcAddress("glGetObjectLabelKHR", false));
    glGetObjectPtrLabelKHR = reinterpret_cast<PFNGLGETOBJECTPTRLABELKHR>(
            gapic::GetGfxProcAddress("glGetObjectPtrLabelKHR", false));
    glGetPointervKHR = reinterpret_cast<PFNGLGETPOINTERVKHR>(
            gapic::GetGfxProcAddress("glGetPointervKHR", false));
    glGetSamplerParameterIivEXT = reinterpret_cast<PFNGLGETSAMPLERPARAMETERIIVEXT>(
            gapic::GetGfxProcAddress("glGetSamplerParameterIivEXT", false));
    glGetSamplerParameterIuivEXT = reinterpret_cast<PFNGLGETSAMPLERPARAMETERIUIVEXT>(
            gapic::GetGfxProcAddress("glGetSamplerParameterIuivEXT", false));
    glGetTexParameterIivEXT = reinterpret_cast<PFNGLGETTEXPARAMETERIIVEXT>(
            gapic::GetGfxProcAddress("glGetTexParameterIivEXT", false));
    glGetTexParameterIuivEXT = reinterpret_cast<PFNGLGETTEXPARAMETERIUIVEXT>(
            gapic::GetGfxProcAddress("glGetTexParameterIuivEXT", false));
    glIsEnablediEXT = reinterpret_cast<PFNGLISENABLEDIEXT>(
            gapic::GetGfxProcAddress("glIsEnablediEXT", false));
    glMinSampleShadingOES = reinterpret_cast<PFNGLMINSAMPLESHADINGOES>(
            gapic::GetGfxProcAddress("glMinSampleShadingOES", false));
    glObjectLabelKHR = reinterpret_cast<PFNGLOBJECTLABELKHR>(
            gapic::GetGfxProcAddress("glObjectLabelKHR", false));
    glObjectPtrLabelKHR = reinterpret_cast<PFNGLOBJECTPTRLABELKHR>(
            gapic::GetGfxProcAddress("glObjectPtrLabelKHR", false));
    glPatchParameteriEXT = reinterpret_cast<PFNGLPATCHPARAMETERIEXT>(
            gapic::GetGfxProcAddress("glPatchParameteriEXT", false));
    glPopDebugGroupKHR = reinterpret_cast<PFNGLPOPDEBUGGROUPKHR>(
            gapic::GetGfxProcAddress("glPopDebugGroupKHR", false));
    glPrimitiveBoundingBoxEXT = reinterpret_cast<PFNGLPRIMITIVEBOUNDINGBOXEXT>(
            gapic::GetGfxProcAddress("glPrimitiveBoundingBoxEXT", false));
    glPushDebugGroupKHR = reinterpret_cast<PFNGLPUSHDEBUGGROUPKHR>(
            gapic::GetGfxProcAddress("glPushDebugGroupKHR", false));
    glSamplerParameterIivEXT = reinterpret_cast<PFNGLSAMPLERPARAMETERIIVEXT>(
            gapic::GetGfxProcAddress("glSamplerParameterIivEXT", false));
    glSamplerParameterIuivEXT = reinterpret_cast<PFNGLSAMPLERPARAMETERIUIVEXT>(
            gapic::GetGfxProcAddress("glSamplerParameterIuivEXT", false));
    glTexBufferEXT =
            reinterpret_cast<PFNGLTEXBUFFEREXT>(gapic::GetGfxProcAddress("glTexBufferEXT", false));
    glTexBufferRangeEXT = reinterpret_cast<PFNGLTEXBUFFERRANGEEXT>(
            gapic::GetGfxProcAddress("glTexBufferRangeEXT", false));
    glTexParameterIivEXT = reinterpret_cast<PFNGLTEXPARAMETERIIVEXT>(
            gapic::GetGfxProcAddress("glTexParameterIivEXT", false));
    glTexParameterIuivEXT = reinterpret_cast<PFNGLTEXPARAMETERIUIVEXT>(
            gapic::GetGfxProcAddress("glTexParameterIuivEXT", false));
    glTexStorage3DMultisampleOES = reinterpret_cast<PFNGLTEXSTORAGE3DMULTISAMPLEOES>(
            gapic::GetGfxProcAddress("glTexStorage3DMultisampleOES", false));
    glBeginQuery =
            reinterpret_cast<PFNGLBEGINQUERY>(gapic::GetGfxProcAddress("glBeginQuery", false));
    glDeleteQueries = reinterpret_cast<PFNGLDELETEQUERIES>(
            gapic::GetGfxProcAddress("glDeleteQueries", false));
    glEndQuery = reinterpret_cast<PFNGLENDQUERY>(gapic::GetGfxProcAddress("glEndQuery", false));
    glGenQueries =
            reinterpret_cast<PFNGLGENQUERIES>(gapic::GetGfxProcAddress("glGenQueries", false));
    glGetQueryObjectuiv = reinterpret_cast<PFNGLGETQUERYOBJECTUIV>(
            gapic::GetGfxProcAddress("glGetQueryObjectuiv", false));
    glGetQueryiv =
            reinterpret_cast<PFNGLGETQUERYIV>(gapic::GetGfxProcAddress("glGetQueryiv", false));
    glIsQuery = reinterpret_cast<PFNGLISQUERY>(gapic::GetGfxProcAddress("glIsQuery", false));
    glBindBuffer =
            reinterpret_cast<PFNGLBINDBUFFER>(gapic::GetGfxProcAddress("glBindBuffer", false));
    glBindBufferBase = reinterpret_cast<PFNGLBINDBUFFERBASE>(
            gapic::GetGfxProcAddress("glBindBufferBase", false));
    glBindBufferRange = reinterpret_cast<PFNGLBINDBUFFERRANGE>(
            gapic::GetGfxProcAddress("glBindBufferRange", false));
    glBufferData =
            reinterpret_cast<PFNGLBUFFERDATA>(gapic::GetGfxProcAddress("glBufferData", false));
    glBufferSubData = reinterpret_cast<PFNGLBUFFERSUBDATA>(
            gapic::GetGfxProcAddress("glBufferSubData", false));
    glCopyBufferSubData = reinterpret_cast<PFNGLCOPYBUFFERSUBDATA>(
            gapic::GetGfxProcAddress("glCopyBufferSubData", false));
    glDeleteBuffers = reinterpret_cast<PFNGLDELETEBUFFERS>(
            gapic::GetGfxProcAddress("glDeleteBuffers", false));
    glGenBuffers =
            reinterpret_cast<PFNGLGENBUFFERS>(gapic::GetGfxProcAddress("glGenBuffers", false));
    glGetBufferParameteri64v = reinterpret_cast<PFNGLGETBUFFERPARAMETERI64V>(
            gapic::GetGfxProcAddress("glGetBufferParameteri64v", false));
    glGetBufferParameteriv = reinterpret_cast<PFNGLGETBUFFERPARAMETERIV>(
            gapic::GetGfxProcAddress("glGetBufferParameteriv", false));
    glGetBufferPointerv = reinterpret_cast<PFNGLGETBUFFERPOINTERV>(
            gapic::GetGfxProcAddress("glGetBufferPointerv", false));
    glIsBuffer = reinterpret_cast<PFNGLISBUFFER>(gapic::GetGfxProcAddress("glIsBuffer", false));
    glMapBufferRange = reinterpret_cast<PFNGLMAPBUFFERRANGE>(
            gapic::GetGfxProcAddress("glMapBufferRange", false));
    glUnmapBuffer =
            reinterpret_cast<PFNGLUNMAPBUFFER>(gapic::GetGfxProcAddress("glUnmapBuffer", false));
    glDebugMessageCallback = reinterpret_cast<PFNGLDEBUGMESSAGECALLBACK>(
            gapic::GetGfxProcAddress("glDebugMessageCallback", false));
    glDebugMessageControl = reinterpret_cast<PFNGLDEBUGMESSAGECONTROL>(
            gapic::GetGfxProcAddress("glDebugMessageControl", false));
    glDebugMessageInsert = reinterpret_cast<PFNGLDEBUGMESSAGEINSERT>(
            gapic::GetGfxProcAddress("glDebugMessageInsert", false));
    glGetDebugMessageLog = reinterpret_cast<PFNGLGETDEBUGMESSAGELOG>(
            gapic::GetGfxProcAddress("glGetDebugMessageLog", false));
    glGetObjectLabel = reinterpret_cast<PFNGLGETOBJECTLABEL>(
            gapic::GetGfxProcAddress("glGetObjectLabel", false));
    glGetObjectPtrLabel = reinterpret_cast<PFNGLGETOBJECTPTRLABEL>(
            gapic::GetGfxProcAddress("glGetObjectPtrLabel", false));
    glGetPointerv =
            reinterpret_cast<PFNGLGETPOINTERV>(gapic::GetGfxProcAddress("glGetPointerv", false));
    glObjectLabel =
            reinterpret_cast<PFNGLOBJECTLABEL>(gapic::GetGfxProcAddress("glObjectLabel", false));
    glObjectPtrLabel = reinterpret_cast<PFNGLOBJECTPTRLABEL>(
            gapic::GetGfxProcAddress("glObjectPtrLabel", false));
    glPopDebugGroup = reinterpret_cast<PFNGLPOPDEBUGGROUP>(
            gapic::GetGfxProcAddress("glPopDebugGroup", false));
    glPushDebugGroup = reinterpret_cast<PFNGLPUSHDEBUGGROUP>(
            gapic::GetGfxProcAddress("glPushDebugGroup", false));
    glDrawArrays =
            reinterpret_cast<PFNGLDRAWARRAYS>(gapic::GetGfxProcAddress("glDrawArrays", false));
    glDrawArraysIndirect = reinterpret_cast<PFNGLDRAWARRAYSINDIRECT>(
            gapic::GetGfxProcAddress("glDrawArraysIndirect", false));
    glDrawArraysInstanced = reinterpret_cast<PFNGLDRAWARRAYSINSTANCED>(
            gapic::GetGfxProcAddress("glDrawArraysInstanced", false));
    glDrawElements =
            reinterpret_cast<PFNGLDRAWELEMENTS>(gapic::GetGfxProcAddress("glDrawElements", false));
    glDrawElementsBaseVertex = reinterpret_cast<PFNGLDRAWELEMENTSBASEVERTEX>(
            gapic::GetGfxProcAddress("glDrawElementsBaseVertex", false));
    glDrawElementsIndirect = reinterpret_cast<PFNGLDRAWELEMENTSINDIRECT>(
            gapic::GetGfxProcAddress("glDrawElementsIndirect", false));
    glDrawElementsInstanced = reinterpret_cast<PFNGLDRAWELEMENTSINSTANCED>(
            gapic::GetGfxProcAddress("glDrawElementsInstanced", false));
    glDrawElementsInstancedBaseVertex = reinterpret_cast<PFNGLDRAWELEMENTSINSTANCEDBASEVERTEX>(
            gapic::GetGfxProcAddress("glDrawElementsInstancedBaseVertex", false));
    glDrawRangeElements = reinterpret_cast<PFNGLDRAWRANGEELEMENTS>(
            gapic::GetGfxProcAddress("glDrawRangeElements", false));
    glDrawRangeElementsBaseVertex = reinterpret_cast<PFNGLDRAWRANGEELEMENTSBASEVERTEX>(
            gapic::GetGfxProcAddress("glDrawRangeElementsBaseVertex", false));
    glPatchParameteri = reinterpret_cast<PFNGLPATCHPARAMETERI>(
            gapic::GetGfxProcAddress("glPatchParameteri", false));
    glPrimitiveBoundingBox = reinterpret_cast<PFNGLPRIMITIVEBOUNDINGBOX>(
            gapic::GetGfxProcAddress("glPrimitiveBoundingBox", false));
    glActiveShaderProgramEXT = reinterpret_cast<PFNGLACTIVESHADERPROGRAMEXT>(
            gapic::GetGfxProcAddress("glActiveShaderProgramEXT", false));
    glAlphaFuncQCOM = reinterpret_cast<PFNGLALPHAFUNCQCOM>(
            gapic::GetGfxProcAddress("glAlphaFuncQCOM", false));
    glApplyFramebufferAttachmentCMAAINTEL =
            reinterpret_cast<PFNGLAPPLYFRAMEBUFFERATTACHMENTCMAAINTEL>(
                    gapic::GetGfxProcAddress("glApplyFramebufferAttachmentCMAAINTEL", false));
    glBeginConditionalRenderNV = reinterpret_cast<PFNGLBEGINCONDITIONALRENDERNV>(
            gapic::GetGfxProcAddress("glBeginConditionalRenderNV", false));
    glBeginPerfMonitorAMD = reinterpret_cast<PFNGLBEGINPERFMONITORAMD>(
            gapic::GetGfxProcAddress("glBeginPerfMonitorAMD", false));
    glBeginPerfQueryINTEL = reinterpret_cast<PFNGLBEGINPERFQUERYINTEL>(
            gapic::GetGfxProcAddress("glBeginPerfQueryINTEL", false));
    glBeginQueryEXT = reinterpret_cast<PFNGLBEGINQUERYEXT>(
            gapic::GetGfxProcAddress("glBeginQueryEXT", false));
    glBindProgramPipelineEXT = reinterpret_cast<PFNGLBINDPROGRAMPIPELINEEXT>(
            gapic::GetGfxProcAddress("glBindProgramPipelineEXT", false));
    glBindVertexArrayOES = reinterpret_cast<PFNGLBINDVERTEXARRAYOES>(
            gapic::GetGfxProcAddress("glBindVertexArrayOES", false));
    glBlendBarrierNV = reinterpret_cast<PFNGLBLENDBARRIERNV>(
            gapic::GetGfxProcAddress("glBlendBarrierNV", false));
    glBlendEquationSeparateiOES = reinterpret_cast<PFNGLBLENDEQUATIONSEPARATEIOES>(
            gapic::GetGfxProcAddress("glBlendEquationSeparateiOES", false));
    glBlendEquationiOES = reinterpret_cast<PFNGLBLENDEQUATIONIOES>(
            gapic::GetGfxProcAddress("glBlendEquationiOES", false));
    glBlendFuncSeparateiOES = reinterpret_cast<PFNGLBLENDFUNCSEPARATEIOES>(
            gapic::GetGfxProcAddress("glBlendFuncSeparateiOES", false));
    glBlendFunciOES = reinterpret_cast<PFNGLBLENDFUNCIOES>(
            gapic::GetGfxProcAddress("glBlendFunciOES", false));
    glBlendParameteriNV = reinterpret_cast<PFNGLBLENDPARAMETERINV>(
            gapic::GetGfxProcAddress("glBlendParameteriNV", false));
    glBlitFramebufferANGLE = reinterpret_cast<PFNGLBLITFRAMEBUFFERANGLE>(
            gapic::GetGfxProcAddress("glBlitFramebufferANGLE", false));
    glBlitFramebufferNV = reinterpret_cast<PFNGLBLITFRAMEBUFFERNV>(
            gapic::GetGfxProcAddress("glBlitFramebufferNV", false));
    glBufferStorageEXT = reinterpret_cast<PFNGLBUFFERSTORAGEEXT>(
            gapic::GetGfxProcAddress("glBufferStorageEXT", false));
    glClientWaitSyncAPPLE = reinterpret_cast<PFNGLCLIENTWAITSYNCAPPLE>(
            gapic::GetGfxProcAddress("glClientWaitSyncAPPLE", false));
    glColorMaskiOES = reinterpret_cast<PFNGLCOLORMASKIOES>(
            gapic::GetGfxProcAddress("glColorMaskiOES", false));
    glCompressedTexImage3DOES = reinterpret_cast<PFNGLCOMPRESSEDTEXIMAGE3DOES>(
            gapic::GetGfxProcAddress("glCompressedTexImage3DOES", false));
    glCompressedTexSubImage3DOES = reinterpret_cast<PFNGLCOMPRESSEDTEXSUBIMAGE3DOES>(
            gapic::GetGfxProcAddress("glCompressedTexSubImage3DOES", false));
    glCopyBufferSubDataNV = reinterpret_cast<PFNGLCOPYBUFFERSUBDATANV>(
            gapic::GetGfxProcAddress("glCopyBufferSubDataNV", false));
    glCopyImageSubDataOES = reinterpret_cast<PFNGLCOPYIMAGESUBDATAOES>(
            gapic::GetGfxProcAddress("glCopyImageSubDataOES", false));
    glCopyPathNV =
            reinterpret_cast<PFNGLCOPYPATHNV>(gapic::GetGfxProcAddress("glCopyPathNV", false));
    glCopyTexSubImage3DOES = reinterpret_cast<PFNGLCOPYTEXSUBIMAGE3DOES>(
            gapic::GetGfxProcAddress("glCopyTexSubImage3DOES", false));
    glCopyTextureLevelsAPPLE = reinterpret_cast<PFNGLCOPYTEXTURELEVELSAPPLE>(
            gapic::GetGfxProcAddress("glCopyTextureLevelsAPPLE", false));
    glCoverFillPathInstancedNV = reinterpret_cast<PFNGLCOVERFILLPATHINSTANCEDNV>(
            gapic::GetGfxProcAddress("glCoverFillPathInstancedNV", false));
    glCoverFillPathNV = reinterpret_cast<PFNGLCOVERFILLPATHNV>(
            gapic::GetGfxProcAddress("glCoverFillPathNV", false));
    glCoverStrokePathInstancedNV = reinterpret_cast<PFNGLCOVERSTROKEPATHINSTANCEDNV>(
            gapic::GetGfxProcAddress("glCoverStrokePathInstancedNV", false));
    glCoverStrokePathNV = reinterpret_cast<PFNGLCOVERSTROKEPATHNV>(
            gapic::GetGfxProcAddress("glCoverStrokePathNV", false));
    glCoverageMaskNV = reinterpret_cast<PFNGLCOVERAGEMASKNV>(
            gapic::GetGfxProcAddress("glCoverageMaskNV", false));
    glCoverageModulationNV = reinterpret_cast<PFNGLCOVERAGEMODULATIONNV>(
            gapic::GetGfxProcAddress("glCoverageModulationNV", false));
    glCoverageModulationTableNV = reinterpret_cast<PFNGLCOVERAGEMODULATIONTABLENV>(
            gapic::GetGfxProcAddress("glCoverageModulationTableNV", false));
    glCoverageOperationNV = reinterpret_cast<PFNGLCOVERAGEOPERATIONNV>(
            gapic::GetGfxProcAddress("glCoverageOperationNV", false));
    glCreatePerfQueryINTEL = reinterpret_cast<PFNGLCREATEPERFQUERYINTEL>(
            gapic::GetGfxProcAddress("glCreatePerfQueryINTEL", false));
    glCreateShaderProgramvEXT = reinterpret_cast<PFNGLCREATESHADERPROGRAMVEXT>(
            gapic::GetGfxProcAddress("glCreateShaderProgramvEXT", false));
    glDeleteFencesNV = reinterpret_cast<PFNGLDELETEFENCESNV>(
            gapic::GetGfxProcAddress("glDeleteFencesNV", false));
    glDeletePathsNV = reinterpret_cast<PFNGLDELETEPATHSNV>(
            gapic::GetGfxProcAddress("glDeletePathsNV", false));
    glDeletePerfMonitorsAMD = reinterpret_cast<PFNGLDELETEPERFMONITORSAMD>(
            gapic::GetGfxProcAddress("glDeletePerfMonitorsAMD", false));
    glDeletePerfQueryINTEL = reinterpret_cast<PFNGLDELETEPERFQUERYINTEL>(
            gapic::GetGfxProcAddress("glDeletePerfQueryINTEL", false));
    glDeleteProgramPipelinesEXT = reinterpret_cast<PFNGLDELETEPROGRAMPIPELINESEXT>(
            gapic::GetGfxProcAddress("glDeleteProgramPipelinesEXT", false));
    glDeleteQueriesEXT = reinterpret_cast<PFNGLDELETEQUERIESEXT>(
            gapic::GetGfxProcAddress("glDeleteQueriesEXT", false));
    glDeleteSyncAPPLE = reinterpret_cast<PFNGLDELETESYNCAPPLE>(
            gapic::GetGfxProcAddress("glDeleteSyncAPPLE", false));
    glDeleteVertexArraysOES = reinterpret_cast<PFNGLDELETEVERTEXARRAYSOES>(
            gapic::GetGfxProcAddress("glDeleteVertexArraysOES", false));
    glDepthRangeArrayfvNV = reinterpret_cast<PFNGLDEPTHRANGEARRAYFVNV>(
            gapic::GetGfxProcAddress("glDepthRangeArrayfvNV", false));
    glDepthRangeIndexedfNV = reinterpret_cast<PFNGLDEPTHRANGEINDEXEDFNV>(
            gapic::GetGfxProcAddress("glDepthRangeIndexedfNV", false));
    glDisableDriverControlQCOM = reinterpret_cast<PFNGLDISABLEDRIVERCONTROLQCOM>(
            gapic::GetGfxProcAddress("glDisableDriverControlQCOM", false));
    glDisableiNV =
            reinterpret_cast<PFNGLDISABLEINV>(gapic::GetGfxProcAddress("glDisableiNV", false));
    glDisableiOES =
            reinterpret_cast<PFNGLDISABLEIOES>(gapic::GetGfxProcAddress("glDisableiOES", false));
    glDiscardFramebufferEXT = reinterpret_cast<PFNGLDISCARDFRAMEBUFFEREXT>(
            gapic::GetGfxProcAddress("glDiscardFramebufferEXT", false));
    glDrawArraysInstancedANGLE = reinterpret_cast<PFNGLDRAWARRAYSINSTANCEDANGLE>(
            gapic::GetGfxProcAddress("glDrawArraysInstancedANGLE", false));
    glDrawArraysInstancedBaseInstanceEXT =
            reinterpret_cast<PFNGLDRAWARRAYSINSTANCEDBASEINSTANCEEXT>(
                    gapic::GetGfxProcAddress("glDrawArraysInstancedBaseInstanceEXT", false));
    glDrawArraysInstancedEXT = reinterpret_cast<PFNGLDRAWARRAYSINSTANCEDEXT>(
            gapic::GetGfxProcAddress("glDrawArraysInstancedEXT", false));
    glDrawArraysInstancedNV = reinterpret_cast<PFNGLDRAWARRAYSINSTANCEDNV>(
            gapic::GetGfxProcAddress("glDrawArraysInstancedNV", false));
    glDrawBuffersEXT = reinterpret_cast<PFNGLDRAWBUFFERSEXT>(
            gapic::GetGfxProcAddress("glDrawBuffersEXT", false));
    glDrawBuffersIndexedEXT = reinterpret_cast<PFNGLDRAWBUFFERSINDEXEDEXT>(
            gapic::GetGfxProcAddress("glDrawBuffersIndexedEXT", false));
    glDrawBuffersNV = reinterpret_cast<PFNGLDRAWBUFFERSNV>(
            gapic::GetGfxProcAddress("glDrawBuffersNV", false));
    glDrawElementsBaseVertexEXT = reinterpret_cast<PFNGLDRAWELEMENTSBASEVERTEXEXT>(
            gapic::GetGfxProcAddress("glDrawElementsBaseVertexEXT", false));
    glDrawElementsBaseVertexOES = reinterpret_cast<PFNGLDRAWELEMENTSBASEVERTEXOES>(
            gapic::GetGfxProcAddress("glDrawElementsBaseVertexOES", false));
    glDrawElementsInstancedANGLE = reinterpret_cast<PFNGLDRAWELEMENTSINSTANCEDANGLE>(
            gapic::GetGfxProcAddress("glDrawElementsInstancedANGLE", false));
    glDrawElementsInstancedBaseInstanceEXT =
            reinterpret_cast<PFNGLDRAWELEMENTSINSTANCEDBASEINSTANCEEXT>(
                    gapic::GetGfxProcAddress("glDrawElementsInstancedBaseInstanceEXT", false));
    glDrawElementsInstancedBaseVertexBaseInstanceEXT =
            reinterpret_cast<PFNGLDRAWELEMENTSINSTANCEDBASEVERTEXBASEINSTANCEEXT>(
                    gapic::GetGfxProcAddress("glDrawElementsInstancedBaseVertexBaseInstanceEXT",
                                             false));
    glDrawElementsInstancedBaseVertexEXT =
            reinterpret_cast<PFNGLDRAWELEMENTSINSTANCEDBASEVERTEXEXT>(
                    gapic::GetGfxProcAddress("glDrawElementsInstancedBaseVertexEXT", false));
    glDrawElementsInstancedBaseVertexOES =
            reinterpret_cast<PFNGLDRAWELEMENTSINSTANCEDBASEVERTEXOES>(
                    gapic::GetGfxProcAddress("glDrawElementsInstancedBaseVertexOES", false));
    glDrawElementsInstancedEXT = reinterpret_cast<PFNGLDRAWELEMENTSINSTANCEDEXT>(
            gapic::GetGfxProcAddress("glDrawElementsInstancedEXT", false));
    glDrawElementsInstancedNV = reinterpret_cast<PFNGLDRAWELEMENTSINSTANCEDNV>(
            gapic::GetGfxProcAddress("glDrawElementsInstancedNV", false));
    glDrawRangeElementsBaseVertexEXT = reinterpret_cast<PFNGLDRAWRANGEELEMENTSBASEVERTEXEXT>(
            gapic::GetGfxProcAddress("glDrawRangeElementsBaseVertexEXT", false));
    glDrawRangeElementsBaseVertexOES = reinterpret_cast<PFNGLDRAWRANGEELEMENTSBASEVERTEXOES>(
            gapic::GetGfxProcAddress("glDrawRangeElementsBaseVertexOES", false));
    glEGLImageTargetRenderbufferStorageOES =
            reinterpret_cast<PFNGLEGLIMAGETARGETRENDERBUFFERSTORAGEOES>(
                    gapic::GetGfxProcAddress("glEGLImageTargetRenderbufferStorageOES", false));
    glEGLImageTargetTexture2DOES = reinterpret_cast<PFNGLEGLIMAGETARGETTEXTURE2DOES>(
            gapic::GetGfxProcAddress("glEGLImageTargetTexture2DOES", false));
    glEnableDriverControlQCOM = reinterpret_cast<PFNGLENABLEDRIVERCONTROLQCOM>(
            gapic::GetGfxProcAddress("glEnableDriverControlQCOM", false));
    glEnableiNV = reinterpret_cast<PFNGLENABLEINV>(gapic::GetGfxProcAddress("glEnableiNV", false));
    glEnableiOES =
            reinterpret_cast<PFNGLENABLEIOES>(gapic::GetGfxProcAddress("glEnableiOES", false));
    glEndConditionalRenderNV = reinterpret_cast<PFNGLENDCONDITIONALRENDERNV>(
            gapic::GetGfxProcAddress("glEndConditionalRenderNV", false));
    glEndPerfMonitorAMD = reinterpret_cast<PFNGLENDPERFMONITORAMD>(
            gapic::GetGfxProcAddress("glEndPerfMonitorAMD", false));
    glEndPerfQueryINTEL = reinterpret_cast<PFNGLENDPERFQUERYINTEL>(
            gapic::GetGfxProcAddress("glEndPerfQueryINTEL", false));
    glEndQueryEXT =
            reinterpret_cast<PFNGLENDQUERYEXT>(gapic::GetGfxProcAddress("glEndQueryEXT", false));
    glEndTilingQCOM = reinterpret_cast<PFNGLENDTILINGQCOM>(
            gapic::GetGfxProcAddress("glEndTilingQCOM", false));
    glExtGetBufferPointervQCOM = reinterpret_cast<PFNGLEXTGETBUFFERPOINTERVQCOM>(
            gapic::GetGfxProcAddress("glExtGetBufferPointervQCOM", false));
    glExtGetBuffersQCOM = reinterpret_cast<PFNGLEXTGETBUFFERSQCOM>(
            gapic::GetGfxProcAddress("glExtGetBuffersQCOM", false));
    glExtGetFramebuffersQCOM = reinterpret_cast<PFNGLEXTGETFRAMEBUFFERSQCOM>(
            gapic::GetGfxProcAddress("glExtGetFramebuffersQCOM", false));
    glExtGetProgramBinarySourceQCOM = reinterpret_cast<PFNGLEXTGETPROGRAMBINARYSOURCEQCOM>(
            gapic::GetGfxProcAddress("glExtGetProgramBinarySourceQCOM", false));
    glExtGetProgramsQCOM = reinterpret_cast<PFNGLEXTGETPROGRAMSQCOM>(
            gapic::GetGfxProcAddress("glExtGetProgramsQCOM", false));
    glExtGetRenderbuffersQCOM = reinterpret_cast<PFNGLEXTGETRENDERBUFFERSQCOM>(
            gapic::GetGfxProcAddress("glExtGetRenderbuffersQCOM", false));
    glExtGetShadersQCOM = reinterpret_cast<PFNGLEXTGETSHADERSQCOM>(
            gapic::GetGfxProcAddress("glExtGetShadersQCOM", false));
    glExtGetTexLevelParameterivQCOM = reinterpret_cast<PFNGLEXTGETTEXLEVELPARAMETERIVQCOM>(
            gapic::GetGfxProcAddress("glExtGetTexLevelParameterivQCOM", false));
    glExtGetTexSubImageQCOM = reinterpret_cast<PFNGLEXTGETTEXSUBIMAGEQCOM>(
            gapic::GetGfxProcAddress("glExtGetTexSubImageQCOM", false));
    glExtGetTexturesQCOM = reinterpret_cast<PFNGLEXTGETTEXTURESQCOM>(
            gapic::GetGfxProcAddress("glExtGetTexturesQCOM", false));
    glExtIsProgramBinaryQCOM = reinterpret_cast<PFNGLEXTISPROGRAMBINARYQCOM>(
            gapic::GetGfxProcAddress("glExtIsProgramBinaryQCOM", false));
    glExtTexObjectStateOverrideiQCOM = reinterpret_cast<PFNGLEXTTEXOBJECTSTATEOVERRIDEIQCOM>(
            gapic::GetGfxProcAddress("glExtTexObjectStateOverrideiQCOM", false));
    glFenceSyncAPPLE = reinterpret_cast<PFNGLFENCESYNCAPPLE>(
            gapic::GetGfxProcAddress("glFenceSyncAPPLE", false));
    glFinishFenceNV = reinterpret_cast<PFNGLFINISHFENCENV>(
            gapic::GetGfxProcAddress("glFinishFenceNV", false));
    glFlushMappedBufferRangeEXT = reinterpret_cast<PFNGLFLUSHMAPPEDBUFFERRANGEEXT>(
            gapic::GetGfxProcAddress("glFlushMappedBufferRangeEXT", false));
    glFragmentCoverageColorNV = reinterpret_cast<PFNGLFRAGMENTCOVERAGECOLORNV>(
            gapic::GetGfxProcAddress("glFragmentCoverageColorNV", false));
    glFramebufferSampleLocationsfvNV = reinterpret_cast<PFNGLFRAMEBUFFERSAMPLELOCATIONSFVNV>(
            gapic::GetGfxProcAddress("glFramebufferSampleLocationsfvNV", false));
    glFramebufferTexture2DMultisampleEXT =
            reinterpret_cast<PFNGLFRAMEBUFFERTEXTURE2DMULTISAMPLEEXT>(
                    gapic::GetGfxProcAddress("glFramebufferTexture2DMultisampleEXT", false));
    glFramebufferTexture2DMultisampleIMG =
            reinterpret_cast<PFNGLFRAMEBUFFERTEXTURE2DMULTISAMPLEIMG>(
                    gapic::GetGfxProcAddress("glFramebufferTexture2DMultisampleIMG", false));
    glFramebufferTexture3DOES = reinterpret_cast<PFNGLFRAMEBUFFERTEXTURE3DOES>(
            gapic::GetGfxProcAddress("glFramebufferTexture3DOES", false));
    glFramebufferTextureMultiviewOVR = reinterpret_cast<PFNGLFRAMEBUFFERTEXTUREMULTIVIEWOVR>(
            gapic::GetGfxProcAddress("glFramebufferTextureMultiviewOVR", false));
    glFramebufferTextureOES = reinterpret_cast<PFNGLFRAMEBUFFERTEXTUREOES>(
            gapic::GetGfxProcAddress("glFramebufferTextureOES", false));
    glGenFencesNV =
            reinterpret_cast<PFNGLGENFENCESNV>(gapic::GetGfxProcAddress("glGenFencesNV", false));
    glGenPathsNV =
            reinterpret_cast<PFNGLGENPATHSNV>(gapic::GetGfxProcAddress("glGenPathsNV", false));
    glGenPerfMonitorsAMD = reinterpret_cast<PFNGLGENPERFMONITORSAMD>(
            gapic::GetGfxProcAddress("glGenPerfMonitorsAMD", false));
    glGenProgramPipelinesEXT = reinterpret_cast<PFNGLGENPROGRAMPIPELINESEXT>(
            gapic::GetGfxProcAddress("glGenProgramPipelinesEXT", false));
    glGenQueriesEXT = reinterpret_cast<PFNGLGENQUERIESEXT>(
            gapic::GetGfxProcAddress("glGenQueriesEXT", false));
    glGenVertexArraysOES = reinterpret_cast<PFNGLGENVERTEXARRAYSOES>(
            gapic::GetGfxProcAddress("glGenVertexArraysOES", false));
    glGetBufferPointervOES = reinterpret_cast<PFNGLGETBUFFERPOINTERVOES>(
            gapic::GetGfxProcAddress("glGetBufferPointervOES", false));
    glGetCoverageModulationTableNV = reinterpret_cast<PFNGLGETCOVERAGEMODULATIONTABLENV>(
            gapic::GetGfxProcAddress("glGetCoverageModulationTableNV", false));
    glGetDriverControlStringQCOM = reinterpret_cast<PFNGLGETDRIVERCONTROLSTRINGQCOM>(
            gapic::GetGfxProcAddress("glGetDriverControlStringQCOM", false));
    glGetDriverControlsQCOM = reinterpret_cast<PFNGLGETDRIVERCONTROLSQCOM>(
            gapic::GetGfxProcAddress("glGetDriverControlsQCOM", false));
    glGetFenceivNV =
            reinterpret_cast<PFNGLGETFENCEIVNV>(gapic::GetGfxProcAddress("glGetFenceivNV", false));
    glGetFirstPerfQueryIdINTEL = reinterpret_cast<PFNGLGETFIRSTPERFQUERYIDINTEL>(
            gapic::GetGfxProcAddress("glGetFirstPerfQueryIdINTEL", false));
    glGetFloati_vNV = reinterpret_cast<PFNGLGETFLOATI_VNV>(
            gapic::GetGfxProcAddress("glGetFloati_vNV", false));
    glGetGraphicsResetStatusEXT = reinterpret_cast<PFNGLGETGRAPHICSRESETSTATUSEXT>(
            gapic::GetGfxProcAddress("glGetGraphicsResetStatusEXT", false));
    glGetGraphicsResetStatusKHR = reinterpret_cast<PFNGLGETGRAPHICSRESETSTATUSKHR>(
            gapic::GetGfxProcAddress("glGetGraphicsResetStatusKHR", false));
    glGetImageHandleNV = reinterpret_cast<PFNGLGETIMAGEHANDLENV>(
            gapic::GetGfxProcAddress("glGetImageHandleNV", false));
    glGetInteger64vAPPLE = reinterpret_cast<PFNGLGETINTEGER64VAPPLE>(
            gapic::GetGfxProcAddress("glGetInteger64vAPPLE", false));
    glGetIntegeri_vEXT = reinterpret_cast<PFNGLGETINTEGERI_VEXT>(
            gapic::GetGfxProcAddress("glGetIntegeri_vEXT", false));
    glGetInternalformatSampleivNV = reinterpret_cast<PFNGLGETINTERNALFORMATSAMPLEIVNV>(
            gapic::GetGfxProcAddress("glGetInternalformatSampleivNV", false));
    glGetNextPerfQueryIdINTEL = reinterpret_cast<PFNGLGETNEXTPERFQUERYIDINTEL>(
            gapic::GetGfxProcAddress("glGetNextPerfQueryIdINTEL", false));
    glGetObjectLabelEXT = reinterpret_cast<PFNGLGETOBJECTLABELEXT>(
            gapic::GetGfxProcAddress("glGetObjectLabelEXT", false));
    glGetPathCommandsNV = reinterpret_cast<PFNGLGETPATHCOMMANDSNV>(
            gapic::GetGfxProcAddress("glGetPathCommandsNV", false));
    glGetPathCoordsNV = reinterpret_cast<PFNGLGETPATHCOORDSNV>(
            gapic::GetGfxProcAddress("glGetPathCoordsNV", false));
    glGetPathDashArrayNV = reinterpret_cast<PFNGLGETPATHDASHARRAYNV>(
            gapic::GetGfxProcAddress("glGetPathDashArrayNV", false));
    glGetPathLengthNV = reinterpret_cast<PFNGLGETPATHLENGTHNV>(
            gapic::GetGfxProcAddress("glGetPathLengthNV", false));
    glGetPathMetricRangeNV = reinterpret_cast<PFNGLGETPATHMETRICRANGENV>(
            gapic::GetGfxProcAddress("glGetPathMetricRangeNV", false));
    glGetPathMetricsNV = reinterpret_cast<PFNGLGETPATHMETRICSNV>(
            gapic::GetGfxProcAddress("glGetPathMetricsNV", false));
    glGetPathParameterfvNV = reinterpret_cast<PFNGLGETPATHPARAMETERFVNV>(
            gapic::GetGfxProcAddress("glGetPathParameterfvNV", false));
    glGetPathParameterivNV = reinterpret_cast<PFNGLGETPATHPARAMETERIVNV>(
            gapic::GetGfxProcAddress("glGetPathParameterivNV", false));
    glGetPathSpacingNV = reinterpret_cast<PFNGLGETPATHSPACINGNV>(
            gapic::GetGfxProcAddress("glGetPathSpacingNV", false));
    glGetPerfCounterInfoINTEL = reinterpret_cast<PFNGLGETPERFCOUNTERINFOINTEL>(
            gapic::GetGfxProcAddress("glGetPerfCounterInfoINTEL", false));
    glGetPerfMonitorCounterDataAMD = reinterpret_cast<PFNGLGETPERFMONITORCOUNTERDATAAMD>(
            gapic::GetGfxProcAddress("glGetPerfMonitorCounterDataAMD", false));
    glGetPerfMonitorCounterInfoAMD = reinterpret_cast<PFNGLGETPERFMONITORCOUNTERINFOAMD>(
            gapic::GetGfxProcAddress("glGetPerfMonitorCounterInfoAMD", false));
    glGetPerfMonitorCounterStringAMD = reinterpret_cast<PFNGLGETPERFMONITORCOUNTERSTRINGAMD>(
            gapic::GetGfxProcAddress("glGetPerfMonitorCounterStringAMD", false));
    glGetPerfMonitorCountersAMD = reinterpret_cast<PFNGLGETPERFMONITORCOUNTERSAMD>(
            gapic::GetGfxProcAddress("glGetPerfMonitorCountersAMD", false));
    glGetPerfMonitorGroupStringAMD = reinterpret_cast<PFNGLGETPERFMONITORGROUPSTRINGAMD>(
            gapic::GetGfxProcAddress("glGetPerfMonitorGroupStringAMD", false));
    glGetPerfMonitorGroupsAMD = reinterpret_cast<PFNGLGETPERFMONITORGROUPSAMD>(
            gapic::GetGfxProcAddress("glGetPerfMonitorGroupsAMD", false));
    glGetPerfQueryDataINTEL = reinterpret_cast<PFNGLGETPERFQUERYDATAINTEL>(
            gapic::GetGfxProcAddress("glGetPerfQueryDataINTEL", false));
    glGetPerfQueryIdByNameINTEL = reinterpret_cast<PFNGLGETPERFQUERYIDBYNAMEINTEL>(
            gapic::GetGfxProcAddress("glGetPerfQueryIdByNameINTEL", false));
    glGetPerfQueryInfoINTEL = reinterpret_cast<PFNGLGETPERFQUERYINFOINTEL>(
            gapic::GetGfxProcAddress("glGetPerfQueryInfoINTEL", false));
    glGetProgramBinaryOES = reinterpret_cast<PFNGLGETPROGRAMBINARYOES>(
            gapic::GetGfxProcAddress("glGetProgramBinaryOES", false));
    glGetProgramPipelineInfoLogEXT = reinterpret_cast<PFNGLGETPROGRAMPIPELINEINFOLOGEXT>(
            gapic::GetGfxProcAddress("glGetProgramPipelineInfoLogEXT", false));
    glGetProgramPipelineivEXT = reinterpret_cast<PFNGLGETPROGRAMPIPELINEIVEXT>(
            gapic::GetGfxProcAddress("glGetProgramPipelineivEXT", false));
    glGetProgramResourcefvNV = reinterpret_cast<PFNGLGETPROGRAMRESOURCEFVNV>(
            gapic::GetGfxProcAddress("glGetProgramResourcefvNV", false));
    glGetQueryObjecti64vEXT = reinterpret_cast<PFNGLGETQUERYOBJECTI64VEXT>(
            gapic::GetGfxProcAddress("glGetQueryObjecti64vEXT", false));
    glGetQueryObjectivEXT = reinterpret_cast<PFNGLGETQUERYOBJECTIVEXT>(
            gapic::GetGfxProcAddress("glGetQueryObjectivEXT", false));
    glGetQueryObjectui64vEXT = reinterpret_cast<PFNGLGETQUERYOBJECTUI64VEXT>(
            gapic::GetGfxProcAddress("glGetQueryObjectui64vEXT", false));
    glGetQueryObjectuivEXT = reinterpret_cast<PFNGLGETQUERYOBJECTUIVEXT>(
            gapic::GetGfxProcAddress("glGetQueryObjectuivEXT", false));
    glGetQueryivEXT = reinterpret_cast<PFNGLGETQUERYIVEXT>(
            gapic::GetGfxProcAddress("glGetQueryivEXT", false));
    glGetSamplerParameterIivOES = reinterpret_cast<PFNGLGETSAMPLERPARAMETERIIVOES>(
            gapic::GetGfxProcAddress("glGetSamplerParameterIivOES", false));
    glGetSamplerParameterIuivOES = reinterpret_cast<PFNGLGETSAMPLERPARAMETERIUIVOES>(
            gapic::GetGfxProcAddress("glGetSamplerParameterIuivOES", false));
    glGetSyncivAPPLE = reinterpret_cast<PFNGLGETSYNCIVAPPLE>(
            gapic::GetGfxProcAddress("glGetSyncivAPPLE", false));
    glGetTexParameterIivOES = reinterpret_cast<PFNGLGETTEXPARAMETERIIVOES>(
            gapic::GetGfxProcAddress("glGetTexParameterIivOES", false));
    glGetTexParameterIuivOES = reinterpret_cast<PFNGLGETTEXPARAMETERIUIVOES>(
            gapic::GetGfxProcAddress("glGetTexParameterIuivOES", false));
    glGetTextureHandleNV = reinterpret_cast<PFNGLGETTEXTUREHANDLENV>(
            gapic::GetGfxProcAddress("glGetTextureHandleNV", false));
    glGetTextureSamplerHandleNV = reinterpret_cast<PFNGLGETTEXTURESAMPLERHANDLENV>(
            gapic::GetGfxProcAddress("glGetTextureSamplerHandleNV", false));
    glGetTranslatedShaderSourceANGLE = reinterpret_cast<PFNGLGETTRANSLATEDSHADERSOURCEANGLE>(
            gapic::GetGfxProcAddress("glGetTranslatedShaderSourceANGLE", false));
    glGetnUniformfvEXT = reinterpret_cast<PFNGLGETNUNIFORMFVEXT>(
            gapic::GetGfxProcAddress("glGetnUniformfvEXT", false));
    glGetnUniformfvKHR = reinterpret_cast<PFNGLGETNUNIFORMFVKHR>(
            gapic::GetGfxProcAddress("glGetnUniformfvKHR", false));
    glGetnUniformivEXT = reinterpret_cast<PFNGLGETNUNIFORMIVEXT>(
            gapic::GetGfxProcAddress("glGetnUniformivEXT", false));
    glGetnUniformivKHR = reinterpret_cast<PFNGLGETNUNIFORMIVKHR>(
            gapic::GetGfxProcAddress("glGetnUniformivKHR", false));
    glGetnUniformuivKHR = reinterpret_cast<PFNGLGETNUNIFORMUIVKHR>(
            gapic::GetGfxProcAddress("glGetnUniformuivKHR", false));
    glInsertEventMarkerEXT = reinterpret_cast<PFNGLINSERTEVENTMARKEREXT>(
            gapic::GetGfxProcAddress("glInsertEventMarkerEXT", false));
    glInterpolatePathsNV = reinterpret_cast<PFNGLINTERPOLATEPATHSNV>(
            gapic::GetGfxProcAddress("glInterpolatePathsNV", false));
    glIsEnablediNV =
            reinterpret_cast<PFNGLISENABLEDINV>(gapic::GetGfxProcAddress("glIsEnablediNV", false));
    glIsEnablediOES = reinterpret_cast<PFNGLISENABLEDIOES>(
            gapic::GetGfxProcAddress("glIsEnablediOES", false));
    glIsFenceNV = reinterpret_cast<PFNGLISFENCENV>(gapic::GetGfxProcAddress("glIsFenceNV", false));
    glIsImageHandleResidentNV = reinterpret_cast<PFNGLISIMAGEHANDLERESIDENTNV>(
            gapic::GetGfxProcAddress("glIsImageHandleResidentNV", false));
    glIsPathNV = reinterpret_cast<PFNGLISPATHNV>(gapic::GetGfxProcAddress("glIsPathNV", false));
    glIsPointInFillPathNV = reinterpret_cast<PFNGLISPOINTINFILLPATHNV>(
            gapic::GetGfxProcAddress("glIsPointInFillPathNV", false));
    glIsPointInStrokePathNV = reinterpret_cast<PFNGLISPOINTINSTROKEPATHNV>(
            gapic::GetGfxProcAddress("glIsPointInStrokePathNV", false));
    glIsProgramPipelineEXT = reinterpret_cast<PFNGLISPROGRAMPIPELINEEXT>(
            gapic::GetGfxProcAddress("glIsProgramPipelineEXT", false));
    glIsQueryEXT =
            reinterpret_cast<PFNGLISQUERYEXT>(gapic::GetGfxProcAddress("glIsQueryEXT", false));
    glIsSyncAPPLE =
            reinterpret_cast<PFNGLISSYNCAPPLE>(gapic::GetGfxProcAddress("glIsSyncAPPLE", false));
    glIsTextureHandleResidentNV = reinterpret_cast<PFNGLISTEXTUREHANDLERESIDENTNV>(
            gapic::GetGfxProcAddress("glIsTextureHandleResidentNV", false));
    glIsVertexArrayOES = reinterpret_cast<PFNGLISVERTEXARRAYOES>(
            gapic::GetGfxProcAddress("glIsVertexArrayOES", false));
    glLabelObjectEXT = reinterpret_cast<PFNGLLABELOBJECTEXT>(
            gapic::GetGfxProcAddress("glLabelObjectEXT", false));
    glMakeImageHandleNonResidentNV = reinterpret_cast<PFNGLMAKEIMAGEHANDLENONRESIDENTNV>(
            gapic::GetGfxProcAddress("glMakeImageHandleNonResidentNV", false));
    glMakeImageHandleResidentNV = reinterpret_cast<PFNGLMAKEIMAGEHANDLERESIDENTNV>(
            gapic::GetGfxProcAddress("glMakeImageHandleResidentNV", false));
    glMakeTextureHandleNonResidentNV = reinterpret_cast<PFNGLMAKETEXTUREHANDLENONRESIDENTNV>(
            gapic::GetGfxProcAddress("glMakeTextureHandleNonResidentNV", false));
    glMakeTextureHandleResidentNV = reinterpret_cast<PFNGLMAKETEXTUREHANDLERESIDENTNV>(
            gapic::GetGfxProcAddress("glMakeTextureHandleResidentNV", false));
    glMapBufferOES =
            reinterpret_cast<PFNGLMAPBUFFEROES>(gapic::GetGfxProcAddress("glMapBufferOES", false));
    glMapBufferRangeEXT = reinterpret_cast<PFNGLMAPBUFFERRANGEEXT>(
            gapic::GetGfxProcAddress("glMapBufferRangeEXT", false));
    glMatrixLoad3x2fNV = reinterpret_cast<PFNGLMATRIXLOAD3X2FNV>(
            gapic::GetGfxProcAddress("glMatrixLoad3x2fNV", false));
    glMatrixLoad3x3fNV = reinterpret_cast<PFNGLMATRIXLOAD3X3FNV>(
            gapic::GetGfxProcAddress("glMatrixLoad3x3fNV", false));
    glMatrixLoadTranspose3x3fNV = reinterpret_cast<PFNGLMATRIXLOADTRANSPOSE3X3FNV>(
            gapic::GetGfxProcAddress("glMatrixLoadTranspose3x3fNV", false));
    glMatrixMult3x2fNV = reinterpret_cast<PFNGLMATRIXMULT3X2FNV>(
            gapic::GetGfxProcAddress("glMatrixMult3x2fNV", false));
    glMatrixMult3x3fNV = reinterpret_cast<PFNGLMATRIXMULT3X3FNV>(
            gapic::GetGfxProcAddress("glMatrixMult3x3fNV", false));
    glMatrixMultTranspose3x3fNV = reinterpret_cast<PFNGLMATRIXMULTTRANSPOSE3X3FNV>(
            gapic::GetGfxProcAddress("glMatrixMultTranspose3x3fNV", false));
    glMultiDrawArraysEXT = reinterpret_cast<PFNGLMULTIDRAWARRAYSEXT>(
            gapic::GetGfxProcAddress("glMultiDrawArraysEXT", false));
    glMultiDrawArraysIndirectEXT = reinterpret_cast<PFNGLMULTIDRAWARRAYSINDIRECTEXT>(
            gapic::GetGfxProcAddress("glMultiDrawArraysIndirectEXT", false));
    glMultiDrawElementsBaseVertexEXT = reinterpret_cast<PFNGLMULTIDRAWELEMENTSBASEVERTEXEXT>(
            gapic::GetGfxProcAddress("glMultiDrawElementsBaseVertexEXT", false));
    glMultiDrawElementsBaseVertexOES = reinterpret_cast<PFNGLMULTIDRAWELEMENTSBASEVERTEXOES>(
            gapic::GetGfxProcAddress("glMultiDrawElementsBaseVertexOES", false));
    glMultiDrawElementsEXT = reinterpret_cast<PFNGLMULTIDRAWELEMENTSEXT>(
            gapic::GetGfxProcAddress("glMultiDrawElementsEXT", false));
    glMultiDrawElementsIndirectEXT = reinterpret_cast<PFNGLMULTIDRAWELEMENTSINDIRECTEXT>(
            gapic::GetGfxProcAddress("glMultiDrawElementsIndirectEXT", false));
    glNamedFramebufferSampleLocationsfvNV =
            reinterpret_cast<PFNGLNAMEDFRAMEBUFFERSAMPLELOCATIONSFVNV>(
                    gapic::GetGfxProcAddress("glNamedFramebufferSampleLocationsfvNV", false));
    glPatchParameteriOES = reinterpret_cast<PFNGLPATCHPARAMETERIOES>(
            gapic::GetGfxProcAddress("glPatchParameteriOES", false));
    glPathCommandsNV = reinterpret_cast<PFNGLPATHCOMMANDSNV>(
            gapic::GetGfxProcAddress("glPathCommandsNV", false));
    glPathCoordsNV =
            reinterpret_cast<PFNGLPATHCOORDSNV>(gapic::GetGfxProcAddress("glPathCoordsNV", false));
    glPathCoverDepthFuncNV = reinterpret_cast<PFNGLPATHCOVERDEPTHFUNCNV>(
            gapic::GetGfxProcAddress("glPathCoverDepthFuncNV", false));
    glPathDashArrayNV = reinterpret_cast<PFNGLPATHDASHARRAYNV>(
            gapic::GetGfxProcAddress("glPathDashArrayNV", false));
    glPathGlyphIndexArrayNV = reinterpret_cast<PFNGLPATHGLYPHINDEXARRAYNV>(
            gapic::GetGfxProcAddress("glPathGlyphIndexArrayNV", false));
    glPathGlyphIndexRangeNV = reinterpret_cast<PFNGLPATHGLYPHINDEXRANGENV>(
            gapic::GetGfxProcAddress("glPathGlyphIndexRangeNV", false));
    glPathGlyphRangeNV = reinterpret_cast<PFNGLPATHGLYPHRANGENV>(
            gapic::GetGfxProcAddress("glPathGlyphRangeNV", false));
    glPathGlyphsNV =
            reinterpret_cast<PFNGLPATHGLYPHSNV>(gapic::GetGfxProcAddress("glPathGlyphsNV", false));
    glPathMemoryGlyphIndexArrayNV = reinterpret_cast<PFNGLPATHMEMORYGLYPHINDEXARRAYNV>(
            gapic::GetGfxProcAddress("glPathMemoryGlyphIndexArrayNV", false));
    glPathParameterfNV = reinterpret_cast<PFNGLPATHPARAMETERFNV>(
            gapic::GetGfxProcAddress("glPathParameterfNV", false));
    glPathParameterfvNV = reinterpret_cast<PFNGLPATHPARAMETERFVNV>(
            gapic::GetGfxProcAddress("glPathParameterfvNV", false));
    glPathParameteriNV = reinterpret_cast<PFNGLPATHPARAMETERINV>(
            gapic::GetGfxProcAddress("glPathParameteriNV", false));
    glPathParameterivNV = reinterpret_cast<PFNGLPATHPARAMETERIVNV>(
            gapic::GetGfxProcAddress("glPathParameterivNV", false));
    glPathStencilDepthOffsetNV = reinterpret_cast<PFNGLPATHSTENCILDEPTHOFFSETNV>(
            gapic::GetGfxProcAddress("glPathStencilDepthOffsetNV", false));
    glPathStencilFuncNV = reinterpret_cast<PFNGLPATHSTENCILFUNCNV>(
            gapic::GetGfxProcAddress("glPathStencilFuncNV", false));
    glPathStringNV =
            reinterpret_cast<PFNGLPATHSTRINGNV>(gapic::GetGfxProcAddress("glPathStringNV", false));
    glPathSubCommandsNV = reinterpret_cast<PFNGLPATHSUBCOMMANDSNV>(
            gapic::GetGfxProcAddress("glPathSubCommandsNV", false));
    glPathSubCoordsNV = reinterpret_cast<PFNGLPATHSUBCOORDSNV>(
            gapic::GetGfxProcAddress("glPathSubCoordsNV", false));
    glPointAlongPathNV = reinterpret_cast<PFNGLPOINTALONGPATHNV>(
            gapic::GetGfxProcAddress("glPointAlongPathNV", false));
    glPolygonModeNV = reinterpret_cast<PFNGLPOLYGONMODENV>(
            gapic::GetGfxProcAddress("glPolygonModeNV", false));
    glPopGroupMarkerEXT = reinterpret_cast<PFNGLPOPGROUPMARKEREXT>(
            gapic::GetGfxProcAddress("glPopGroupMarkerEXT", false));
    glPrimitiveBoundingBoxOES = reinterpret_cast<PFNGLPRIMITIVEBOUNDINGBOXOES>(
            gapic::GetGfxProcAddress("glPrimitiveBoundingBoxOES", false));
    glProgramBinaryOES = reinterpret_cast<PFNGLPROGRAMBINARYOES>(
            gapic::GetGfxProcAddress("glProgramBinaryOES", false));
    glProgramParameteriEXT = reinterpret_cast<PFNGLPROGRAMPARAMETERIEXT>(
            gapic::GetGfxProcAddress("glProgramParameteriEXT", false));
    glProgramPathFragmentInputGenNV = reinterpret_cast<PFNGLPROGRAMPATHFRAGMENTINPUTGENNV>(
            gapic::GetGfxProcAddress("glProgramPathFragmentInputGenNV", false));
    glProgramUniform1fEXT = reinterpret_cast<PFNGLPROGRAMUNIFORM1FEXT>(
            gapic::GetGfxProcAddress("glProgramUniform1fEXT", false));
    glProgramUniform1fvEXT = reinterpret_cast<PFNGLPROGRAMUNIFORM1FVEXT>(
            gapic::GetGfxProcAddress("glProgramUniform1fvEXT", false));
    glProgramUniform1iEXT = reinterpret_cast<PFNGLPROGRAMUNIFORM1IEXT>(
            gapic::GetGfxProcAddress("glProgramUniform1iEXT", false));
    glProgramUniform1ivEXT = reinterpret_cast<PFNGLPROGRAMUNIFORM1IVEXT>(
            gapic::GetGfxProcAddress("glProgramUniform1ivEXT", false));
    glProgramUniform1uiEXT = reinterpret_cast<PFNGLPROGRAMUNIFORM1UIEXT>(
            gapic::GetGfxProcAddress("glProgramUniform1uiEXT", false));
    glProgramUniform1uivEXT = reinterpret_cast<PFNGLPROGRAMUNIFORM1UIVEXT>(
            gapic::GetGfxProcAddress("glProgramUniform1uivEXT", false));
    glProgramUniform2fEXT = reinterpret_cast<PFNGLPROGRAMUNIFORM2FEXT>(
            gapic::GetGfxProcAddress("glProgramUniform2fEXT", false));
    glProgramUniform2fvEXT = reinterpret_cast<PFNGLPROGRAMUNIFORM2FVEXT>(
            gapic::GetGfxProcAddress("glProgramUniform2fvEXT", false));
    glProgramUniform2iEXT = reinterpret_cast<PFNGLPROGRAMUNIFORM2IEXT>(
            gapic::GetGfxProcAddress("glProgramUniform2iEXT", false));
    glProgramUniform2ivEXT = reinterpret_cast<PFNGLPROGRAMUNIFORM2IVEXT>(
            gapic::GetGfxProcAddress("glProgramUniform2ivEXT", false));
    glProgramUniform2uiEXT = reinterpret_cast<PFNGLPROGRAMUNIFORM2UIEXT>(
            gapic::GetGfxProcAddress("glProgramUniform2uiEXT", false));
    glProgramUniform2uivEXT = reinterpret_cast<PFNGLPROGRAMUNIFORM2UIVEXT>(
            gapic::GetGfxProcAddress("glProgramUniform2uivEXT", false));
    glProgramUniform3fEXT = reinterpret_cast<PFNGLPROGRAMUNIFORM3FEXT>(
            gapic::GetGfxProcAddress("glProgramUniform3fEXT", false));
    glProgramUniform3fvEXT = reinterpret_cast<PFNGLPROGRAMUNIFORM3FVEXT>(
            gapic::GetGfxProcAddress("glProgramUniform3fvEXT", false));
    glProgramUniform3iEXT = reinterpret_cast<PFNGLPROGRAMUNIFORM3IEXT>(
            gapic::GetGfxProcAddress("glProgramUniform3iEXT", false));
    glProgramUniform3ivEXT = reinterpret_cast<PFNGLPROGRAMUNIFORM3IVEXT>(
            gapic::GetGfxProcAddress("glProgramUniform3ivEXT", false));
    glProgramUniform3uiEXT = reinterpret_cast<PFNGLPROGRAMUNIFORM3UIEXT>(
            gapic::GetGfxProcAddress("glProgramUniform3uiEXT", false));
    glProgramUniform3uivEXT = reinterpret_cast<PFNGLPROGRAMUNIFORM3UIVEXT>(
            gapic::GetGfxProcAddress("glProgramUniform3uivEXT", false));
    glProgramUniform4fEXT = reinterpret_cast<PFNGLPROGRAMUNIFORM4FEXT>(
            gapic::GetGfxProcAddress("glProgramUniform4fEXT", false));
    glProgramUniform4fvEXT = reinterpret_cast<PFNGLPROGRAMUNIFORM4FVEXT>(
            gapic::GetGfxProcAddress("glProgramUniform4fvEXT", false));
    glProgramUniform4iEXT = reinterpret_cast<PFNGLPROGRAMUNIFORM4IEXT>(
            gapic::GetGfxProcAddress("glProgramUniform4iEXT", false));
    glProgramUniform4ivEXT = reinterpret_cast<PFNGLPROGRAMUNIFORM4IVEXT>(
            gapic::GetGfxProcAddress("glProgramUniform4ivEXT", false));
    glProgramUniform4uiEXT = reinterpret_cast<PFNGLPROGRAMUNIFORM4UIEXT>(
            gapic::GetGfxProcAddress("glProgramUniform4uiEXT", false));
    glProgramUniform4uivEXT = reinterpret_cast<PFNGLPROGRAMUNIFORM4UIVEXT>(
            gapic::GetGfxProcAddress("glProgramUniform4uivEXT", false));
    glProgramUniformHandleui64NV = reinterpret_cast<PFNGLPROGRAMUNIFORMHANDLEUI64NV>(
            gapic::GetGfxProcAddress("glProgramUniformHandleui64NV", false));
    glProgramUniformHandleui64vNV = reinterpret_cast<PFNGLPROGRAMUNIFORMHANDLEUI64VNV>(
            gapic::GetGfxProcAddress("glProgramUniformHandleui64vNV", false));
    glProgramUniformMatrix2fvEXT = reinterpret_cast<PFNGLPROGRAMUNIFORMMATRIX2FVEXT>(
            gapic::GetGfxProcAddress("glProgramUniformMatrix2fvEXT", false));
    glProgramUniformMatrix2x3fvEXT = reinterpret_cast<PFNGLPROGRAMUNIFORMMATRIX2X3FVEXT>(
            gapic::GetGfxProcAddress("glProgramUniformMatrix2x3fvEXT", false));
    glProgramUniformMatrix2x4fvEXT = reinterpret_cast<PFNGLPROGRAMUNIFORMMATRIX2X4FVEXT>(
            gapic::GetGfxProcAddress("glProgramUniformMatrix2x4fvEXT", false));
    glProgramUniformMatrix3fvEXT = reinterpret_cast<PFNGLPROGRAMUNIFORMMATRIX3FVEXT>(
            gapic::GetGfxProcAddress("glProgramUniformMatrix3fvEXT", false));
    glProgramUniformMatrix3x2fvEXT = reinterpret_cast<PFNGLPROGRAMUNIFORMMATRIX3X2FVEXT>(
            gapic::GetGfxProcAddress("glProgramUniformMatrix3x2fvEXT", false));
    glProgramUniformMatrix3x4fvEXT = reinterpret_cast<PFNGLPROGRAMUNIFORMMATRIX3X4FVEXT>(
            gapic::GetGfxProcAddress("glProgramUniformMatrix3x4fvEXT", false));
    glProgramUniformMatrix4fvEXT = reinterpret_cast<PFNGLPROGRAMUNIFORMMATRIX4FVEXT>(
            gapic::GetGfxProcAddress("glProgramUniformMatrix4fvEXT", false));
    glProgramUniformMatrix4x2fvEXT = reinterpret_cast<PFNGLPROGRAMUNIFORMMATRIX4X2FVEXT>(
            gapic::GetGfxProcAddress("glProgramUniformMatrix4x2fvEXT", false));
    glProgramUniformMatrix4x3fvEXT = reinterpret_cast<PFNGLPROGRAMUNIFORMMATRIX4X3FVEXT>(
            gapic::GetGfxProcAddress("glProgramUniformMatrix4x3fvEXT", false));
    glPushGroupMarkerEXT = reinterpret_cast<PFNGLPUSHGROUPMARKEREXT>(
            gapic::GetGfxProcAddress("glPushGroupMarkerEXT", false));
    glQueryCounterEXT = reinterpret_cast<PFNGLQUERYCOUNTEREXT>(
            gapic::GetGfxProcAddress("glQueryCounterEXT", false));
    glRasterSamplesEXT = reinterpret_cast<PFNGLRASTERSAMPLESEXT>(
            gapic::GetGfxProcAddress("glRasterSamplesEXT", false));
    glReadBufferIndexedEXT = reinterpret_cast<PFNGLREADBUFFERINDEXEDEXT>(
            gapic::GetGfxProcAddress("glReadBufferIndexedEXT", false));
    glReadBufferNV =
            reinterpret_cast<PFNGLREADBUFFERNV>(gapic::GetGfxProcAddress("glReadBufferNV", false));
    glReadnPixelsEXT = reinterpret_cast<PFNGLREADNPIXELSEXT>(
            gapic::GetGfxProcAddress("glReadnPixelsEXT", false));
    glReadnPixelsKHR = reinterpret_cast<PFNGLREADNPIXELSKHR>(
            gapic::GetGfxProcAddress("glReadnPixelsKHR", false));
    glRenderbufferStorageMultisampleANGLE =
            reinterpret_cast<PFNGLRENDERBUFFERSTORAGEMULTISAMPLEANGLE>(
                    gapic::GetGfxProcAddress("glRenderbufferStorageMultisampleANGLE", false));
    glRenderbufferStorageMultisampleAPPLE =
            reinterpret_cast<PFNGLRENDERBUFFERSTORAGEMULTISAMPLEAPPLE>(
                    gapic::GetGfxProcAddress("glRenderbufferStorageMultisampleAPPLE", false));
    glRenderbufferStorageMultisampleEXT = reinterpret_cast<PFNGLRENDERBUFFERSTORAGEMULTISAMPLEEXT>(
            gapic::GetGfxProcAddress("glRenderbufferStorageMultisampleEXT", false));
    glRenderbufferStorageMultisampleIMG = reinterpret_cast<PFNGLRENDERBUFFERSTORAGEMULTISAMPLEIMG>(
            gapic::GetGfxProcAddress("glRenderbufferStorageMultisampleIMG", false));
    glRenderbufferStorageMultisampleNV = reinterpret_cast<PFNGLRENDERBUFFERSTORAGEMULTISAMPLENV>(
            gapic::GetGfxProcAddress("glRenderbufferStorageMultisampleNV", false));
    glResolveDepthValuesNV = reinterpret_cast<PFNGLRESOLVEDEPTHVALUESNV>(
            gapic::GetGfxProcAddress("glResolveDepthValuesNV", false));
    glResolveMultisampleFramebufferAPPLE =
            reinterpret_cast<PFNGLRESOLVEMULTISAMPLEFRAMEBUFFERAPPLE>(
                    gapic::GetGfxProcAddress("glResolveMultisampleFramebufferAPPLE", false));
    glSamplerParameterIivOES = reinterpret_cast<PFNGLSAMPLERPARAMETERIIVOES>(
            gapic::GetGfxProcAddress("glSamplerParameterIivOES", false));
    glSamplerParameterIuivOES = reinterpret_cast<PFNGLSAMPLERPARAMETERIUIVOES>(
            gapic::GetGfxProcAddress("glSamplerParameterIuivOES", false));
    glScissorArrayvNV = reinterpret_cast<PFNGLSCISSORARRAYVNV>(
            gapic::GetGfxProcAddress("glScissorArrayvNV", false));
    glScissorIndexedNV = reinterpret_cast<PFNGLSCISSORINDEXEDNV>(
            gapic::GetGfxProcAddress("glScissorIndexedNV", false));
    glScissorIndexedvNV = reinterpret_cast<PFNGLSCISSORINDEXEDVNV>(
            gapic::GetGfxProcAddress("glScissorIndexedvNV", false));
    glSelectPerfMonitorCountersAMD = reinterpret_cast<PFNGLSELECTPERFMONITORCOUNTERSAMD>(
            gapic::GetGfxProcAddress("glSelectPerfMonitorCountersAMD", false));
    glSetFenceNV =
            reinterpret_cast<PFNGLSETFENCENV>(gapic::GetGfxProcAddress("glSetFenceNV", false));
    glStartTilingQCOM = reinterpret_cast<PFNGLSTARTTILINGQCOM>(
            gapic::GetGfxProcAddress("glStartTilingQCOM", false));
    glStencilFillPathInstancedNV = reinterpret_cast<PFNGLSTENCILFILLPATHINSTANCEDNV>(
            gapic::GetGfxProcAddress("glStencilFillPathInstancedNV", false));
    glStencilFillPathNV = reinterpret_cast<PFNGLSTENCILFILLPATHNV>(
            gapic::GetGfxProcAddress("glStencilFillPathNV", false));
    glStencilStrokePathInstancedNV = reinterpret_cast<PFNGLSTENCILSTROKEPATHINSTANCEDNV>(
            gapic::GetGfxProcAddress("glStencilStrokePathInstancedNV", false));
    glStencilStrokePathNV = reinterpret_cast<PFNGLSTENCILSTROKEPATHNV>(
            gapic::GetGfxProcAddress("glStencilStrokePathNV", false));
    glStencilThenCoverFillPathInstancedNV =
            reinterpret_cast<PFNGLSTENCILTHENCOVERFILLPATHINSTANCEDNV>(
                    gapic::GetGfxProcAddress("glStencilThenCoverFillPathInstancedNV", false));
    glStencilThenCoverFillPathNV = reinterpret_cast<PFNGLSTENCILTHENCOVERFILLPATHNV>(
            gapic::GetGfxProcAddress("glStencilThenCoverFillPathNV", false));
    glStencilThenCoverStrokePathInstancedNV =
            reinterpret_cast<PFNGLSTENCILTHENCOVERSTROKEPATHINSTANCEDNV>(
                    gapic::GetGfxProcAddress("glStencilThenCoverStrokePathInstancedNV", false));
    glStencilThenCoverStrokePathNV = reinterpret_cast<PFNGLSTENCILTHENCOVERSTROKEPATHNV>(
            gapic::GetGfxProcAddress("glStencilThenCoverStrokePathNV", false));
    glSubpixelPrecisionBiasNV = reinterpret_cast<PFNGLSUBPIXELPRECISIONBIASNV>(
            gapic::GetGfxProcAddress("glSubpixelPrecisionBiasNV", false));
    glTestFenceNV =
            reinterpret_cast<PFNGLTESTFENCENV>(gapic::GetGfxProcAddress("glTestFenceNV", false));
    glTexBufferOES =
            reinterpret_cast<PFNGLTEXBUFFEROES>(gapic::GetGfxProcAddress("glTexBufferOES", false));
    glTexBufferRangeOES = reinterpret_cast<PFNGLTEXBUFFERRANGEOES>(
            gapic::GetGfxProcAddress("glTexBufferRangeOES", false));
    glTexImage3DOES = reinterpret_cast<PFNGLTEXIMAGE3DOES>(
            gapic::GetGfxProcAddress("glTexImage3DOES", false));
    glTexPageCommitmentARB = reinterpret_cast<PFNGLTEXPAGECOMMITMENTARB>(
            gapic::GetGfxProcAddress("glTexPageCommitmentARB", false));
    glTexParameterIivOES = reinterpret_cast<PFNGLTEXPARAMETERIIVOES>(
            gapic::GetGfxProcAddress("glTexParameterIivOES", false));
    glTexParameterIuivOES = reinterpret_cast<PFNGLTEXPARAMETERIUIVOES>(
            gapic::GetGfxProcAddress("glTexParameterIuivOES", false));
    glTexStorage1DEXT = reinterpret_cast<PFNGLTEXSTORAGE1DEXT>(
            gapic::GetGfxProcAddress("glTexStorage1DEXT", false));
    glTexStorage2DEXT = reinterpret_cast<PFNGLTEXSTORAGE2DEXT>(
            gapic::GetGfxProcAddress("glTexStorage2DEXT", false));
    glTexStorage3DEXT = reinterpret_cast<PFNGLTEXSTORAGE3DEXT>(
            gapic::GetGfxProcAddress("glTexStorage3DEXT", false));
    glTexSubImage3DOES = reinterpret_cast<PFNGLTEXSUBIMAGE3DOES>(
            gapic::GetGfxProcAddress("glTexSubImage3DOES", false));
    glTextureStorage1DEXT = reinterpret_cast<PFNGLTEXTURESTORAGE1DEXT>(
            gapic::GetGfxProcAddress("glTextureStorage1DEXT", false));
    glTextureStorage2DEXT = reinterpret_cast<PFNGLTEXTURESTORAGE2DEXT>(
            gapic::GetGfxProcAddress("glTextureStorage2DEXT", false));
    glTextureStorage3DEXT = reinterpret_cast<PFNGLTEXTURESTORAGE3DEXT>(
            gapic::GetGfxProcAddress("glTextureStorage3DEXT", false));
    glTextureViewEXT = reinterpret_cast<PFNGLTEXTUREVIEWEXT>(
            gapic::GetGfxProcAddress("glTextureViewEXT", false));
    glTextureViewOES = reinterpret_cast<PFNGLTEXTUREVIEWOES>(
            gapic::GetGfxProcAddress("glTextureViewOES", false));
    glTransformPathNV = reinterpret_cast<PFNGLTRANSFORMPATHNV>(
            gapic::GetGfxProcAddress("glTransformPathNV", false));
    glUniformHandleui64NV = reinterpret_cast<PFNGLUNIFORMHANDLEUI64NV>(
            gapic::GetGfxProcAddress("glUniformHandleui64NV", false));
    glUniformHandleui64vNV = reinterpret_cast<PFNGLUNIFORMHANDLEUI64VNV>(
            gapic::GetGfxProcAddress("glUniformHandleui64vNV", false));
    glUniformMatrix2x3fvNV = reinterpret_cast<PFNGLUNIFORMMATRIX2X3FVNV>(
            gapic::GetGfxProcAddress("glUniformMatrix2x3fvNV", false));
    glUniformMatrix2x4fvNV = reinterpret_cast<PFNGLUNIFORMMATRIX2X4FVNV>(
            gapic::GetGfxProcAddress("glUniformMatrix2x4fvNV", false));
    glUniformMatrix3x2fvNV = reinterpret_cast<PFNGLUNIFORMMATRIX3X2FVNV>(
            gapic::GetGfxProcAddress("glUniformMatrix3x2fvNV", false));
    glUniformMatrix3x4fvNV = reinterpret_cast<PFNGLUNIFORMMATRIX3X4FVNV>(
            gapic::GetGfxProcAddress("glUniformMatrix3x4fvNV", false));
    glUniformMatrix4x2fvNV = reinterpret_cast<PFNGLUNIFORMMATRIX4X2FVNV>(
            gapic::GetGfxProcAddress("glUniformMatrix4x2fvNV", false));
    glUniformMatrix4x3fvNV = reinterpret_cast<PFNGLUNIFORMMATRIX4X3FVNV>(
            gapic::GetGfxProcAddress("glUniformMatrix4x3fvNV", false));
    glUnmapBufferOES = reinterpret_cast<PFNGLUNMAPBUFFEROES>(
            gapic::GetGfxProcAddress("glUnmapBufferOES", false));
    glUseProgramStagesEXT = reinterpret_cast<PFNGLUSEPROGRAMSTAGESEXT>(
            gapic::GetGfxProcAddress("glUseProgramStagesEXT", false));
    glValidateProgramPipelineEXT = reinterpret_cast<PFNGLVALIDATEPROGRAMPIPELINEEXT>(
            gapic::GetGfxProcAddress("glValidateProgramPipelineEXT", false));
    glVertexAttribDivisorANGLE = reinterpret_cast<PFNGLVERTEXATTRIBDIVISORANGLE>(
            gapic::GetGfxProcAddress("glVertexAttribDivisorANGLE", false));
    glVertexAttribDivisorEXT = reinterpret_cast<PFNGLVERTEXATTRIBDIVISOREXT>(
            gapic::GetGfxProcAddress("glVertexAttribDivisorEXT", false));
    glVertexAttribDivisorNV = reinterpret_cast<PFNGLVERTEXATTRIBDIVISORNV>(
            gapic::GetGfxProcAddress("glVertexAttribDivisorNV", false));
    glViewportArrayvNV = reinterpret_cast<PFNGLVIEWPORTARRAYVNV>(
            gapic::GetGfxProcAddress("glViewportArrayvNV", false));
    glViewportIndexedfNV = reinterpret_cast<PFNGLVIEWPORTINDEXEDFNV>(
            gapic::GetGfxProcAddress("glViewportIndexedfNV", false));
    glViewportIndexedfvNV = reinterpret_cast<PFNGLVIEWPORTINDEXEDFVNV>(
            gapic::GetGfxProcAddress("glViewportIndexedfvNV", false));
    glWaitSyncAPPLE = reinterpret_cast<PFNGLWAITSYNCAPPLE>(
            gapic::GetGfxProcAddress("glWaitSyncAPPLE", false));
    glWeightPathsNV = reinterpret_cast<PFNGLWEIGHTPATHSNV>(
            gapic::GetGfxProcAddress("glWeightPathsNV", false));
    glBlendBarrier =
            reinterpret_cast<PFNGLBLENDBARRIER>(gapic::GetGfxProcAddress("glBlendBarrier", false));
    glBlendColor =
            reinterpret_cast<PFNGLBLENDCOLOR>(gapic::GetGfxProcAddress("glBlendColor", false));
    glBlendEquation = reinterpret_cast<PFNGLBLENDEQUATION>(
            gapic::GetGfxProcAddress("glBlendEquation", false));
    glBlendEquationSeparate = reinterpret_cast<PFNGLBLENDEQUATIONSEPARATE>(
            gapic::GetGfxProcAddress("glBlendEquationSeparate", false));
    glBlendEquationSeparatei = reinterpret_cast<PFNGLBLENDEQUATIONSEPARATEI>(
            gapic::GetGfxProcAddress("glBlendEquationSeparatei", false));
    glBlendEquationi = reinterpret_cast<PFNGLBLENDEQUATIONI>(
            gapic::GetGfxProcAddress("glBlendEquationi", false));
    glBlendFunc = reinterpret_cast<PFNGLBLENDFUNC>(gapic::GetGfxProcAddress("glBlendFunc", false));
    glBlendFuncSeparate = reinterpret_cast<PFNGLBLENDFUNCSEPARATE>(
            gapic::GetGfxProcAddress("glBlendFuncSeparate", false));
    glBlendFuncSeparatei = reinterpret_cast<PFNGLBLENDFUNCSEPARATEI>(
            gapic::GetGfxProcAddress("glBlendFuncSeparatei", false));
    glBlendFunci =
            reinterpret_cast<PFNGLBLENDFUNCI>(gapic::GetGfxProcAddress("glBlendFunci", false));
    glDepthFunc = reinterpret_cast<PFNGLDEPTHFUNC>(gapic::GetGfxProcAddress("glDepthFunc", false));
    glSampleCoverage = reinterpret_cast<PFNGLSAMPLECOVERAGE>(
            gapic::GetGfxProcAddress("glSampleCoverage", false));
    glSampleMaski =
            reinterpret_cast<PFNGLSAMPLEMASKI>(gapic::GetGfxProcAddress("glSampleMaski", false));
    glScissor = reinterpret_cast<PFNGLSCISSOR>(gapic::GetGfxProcAddress("glScissor", false));
    glStencilFunc =
            reinterpret_cast<PFNGLSTENCILFUNC>(gapic::GetGfxProcAddress("glStencilFunc", false));
    glStencilFuncSeparate = reinterpret_cast<PFNGLSTENCILFUNCSEPARATE>(
            gapic::GetGfxProcAddress("glStencilFuncSeparate", false));
    glStencilOp = reinterpret_cast<PFNGLSTENCILOP>(gapic::GetGfxProcAddress("glStencilOp", false));
    glStencilOpSeparate = reinterpret_cast<PFNGLSTENCILOPSEPARATE>(
            gapic::GetGfxProcAddress("glStencilOpSeparate", false));
    glBindFramebuffer = reinterpret_cast<PFNGLBINDFRAMEBUFFER>(
            gapic::GetGfxProcAddress("glBindFramebuffer", false));
    glBindRenderbuffer = reinterpret_cast<PFNGLBINDRENDERBUFFER>(
            gapic::GetGfxProcAddress("glBindRenderbuffer", false));
    glBlitFramebuffer = reinterpret_cast<PFNGLBLITFRAMEBUFFER>(
            gapic::GetGfxProcAddress("glBlitFramebuffer", false));
    glCheckFramebufferStatus = reinterpret_cast<PFNGLCHECKFRAMEBUFFERSTATUS>(
            gapic::GetGfxProcAddress("glCheckFramebufferStatus", false));
    glClear = reinterpret_cast<PFNGLCLEAR>(gapic::GetGfxProcAddress("glClear", false));
    glClearBufferfi = reinterpret_cast<PFNGLCLEARBUFFERFI>(
            gapic::GetGfxProcAddress("glClearBufferfi", false));
    glClearBufferfv = reinterpret_cast<PFNGLCLEARBUFFERFV>(
            gapic::GetGfxProcAddress("glClearBufferfv", false));
    glClearBufferiv = reinterpret_cast<PFNGLCLEARBUFFERIV>(
            gapic::GetGfxProcAddress("glClearBufferiv", false));
    glClearBufferuiv = reinterpret_cast<PFNGLCLEARBUFFERUIV>(
            gapic::GetGfxProcAddress("glClearBufferuiv", false));
    glClearColor =
            reinterpret_cast<PFNGLCLEARCOLOR>(gapic::GetGfxProcAddress("glClearColor", false));
    glClearDepthf =
            reinterpret_cast<PFNGLCLEARDEPTHF>(gapic::GetGfxProcAddress("glClearDepthf", false));
    glClearStencil =
            reinterpret_cast<PFNGLCLEARSTENCIL>(gapic::GetGfxProcAddress("glClearStencil", false));
    glColorMask = reinterpret_cast<PFNGLCOLORMASK>(gapic::GetGfxProcAddress("glColorMask", false));
    glColorMaski =
            reinterpret_cast<PFNGLCOLORMASKI>(gapic::GetGfxProcAddress("glColorMaski", false));
    glDeleteFramebuffers = reinterpret_cast<PFNGLDELETEFRAMEBUFFERS>(
            gapic::GetGfxProcAddress("glDeleteFramebuffers", false));
    glDeleteRenderbuffers = reinterpret_cast<PFNGLDELETERENDERBUFFERS>(
            gapic::GetGfxProcAddress("glDeleteRenderbuffers", false));
    glDepthMask = reinterpret_cast<PFNGLDEPTHMASK>(gapic::GetGfxProcAddress("glDepthMask", false));
    glDrawBuffers =
            reinterpret_cast<PFNGLDRAWBUFFERS>(gapic::GetGfxProcAddress("glDrawBuffers", false));
    glFramebufferParameteri = reinterpret_cast<PFNGLFRAMEBUFFERPARAMETERI>(
            gapic::GetGfxProcAddress("glFramebufferParameteri", false));
    glFramebufferRenderbuffer = reinterpret_cast<PFNGLFRAMEBUFFERRENDERBUFFER>(
            gapic::GetGfxProcAddress("glFramebufferRenderbuffer", false));
    glFramebufferTexture = reinterpret_cast<PFNGLFRAMEBUFFERTEXTURE>(
            gapic::GetGfxProcAddress("glFramebufferTexture", false));
    glFramebufferTexture2D = reinterpret_cast<PFNGLFRAMEBUFFERTEXTURE2D>(
            gapic::GetGfxProcAddress("glFramebufferTexture2D", false));
    glFramebufferTextureLayer = reinterpret_cast<PFNGLFRAMEBUFFERTEXTURELAYER>(
            gapic::GetGfxProcAddress("glFramebufferTextureLayer", false));
    glGenFramebuffers = reinterpret_cast<PFNGLGENFRAMEBUFFERS>(
            gapic::GetGfxProcAddress("glGenFramebuffers", false));
    glGenRenderbuffers = reinterpret_cast<PFNGLGENRENDERBUFFERS>(
            gapic::GetGfxProcAddress("glGenRenderbuffers", false));
    glGetFramebufferAttachmentParameteriv =
            reinterpret_cast<PFNGLGETFRAMEBUFFERATTACHMENTPARAMETERIV>(
                    gapic::GetGfxProcAddress("glGetFramebufferAttachmentParameteriv", false));
    glGetFramebufferParameteriv = reinterpret_cast<PFNGLGETFRAMEBUFFERPARAMETERIV>(
            gapic::GetGfxProcAddress("glGetFramebufferParameteriv", false));
    glGetRenderbufferParameteriv = reinterpret_cast<PFNGLGETRENDERBUFFERPARAMETERIV>(
            gapic::GetGfxProcAddress("glGetRenderbufferParameteriv", false));
    glInvalidateFramebuffer = reinterpret_cast<PFNGLINVALIDATEFRAMEBUFFER>(
            gapic::GetGfxProcAddress("glInvalidateFramebuffer", false));
    glInvalidateSubFramebuffer = reinterpret_cast<PFNGLINVALIDATESUBFRAMEBUFFER>(
            gapic::GetGfxProcAddress("glInvalidateSubFramebuffer", false));
    glIsFramebuffer = reinterpret_cast<PFNGLISFRAMEBUFFER>(
            gapic::GetGfxProcAddress("glIsFramebuffer", false));
    glIsRenderbuffer = reinterpret_cast<PFNGLISRENDERBUFFER>(
            gapic::GetGfxProcAddress("glIsRenderbuffer", false));
    glReadBuffer =
            reinterpret_cast<PFNGLREADBUFFER>(gapic::GetGfxProcAddress("glReadBuffer", false));
    glReadPixels =
            reinterpret_cast<PFNGLREADPIXELS>(gapic::GetGfxProcAddress("glReadPixels", false));
    glReadnPixels =
            reinterpret_cast<PFNGLREADNPIXELS>(gapic::GetGfxProcAddress("glReadnPixels", false));
    glRenderbufferStorage = reinterpret_cast<PFNGLRENDERBUFFERSTORAGE>(
            gapic::GetGfxProcAddress("glRenderbufferStorage", false));
    glRenderbufferStorageMultisample = reinterpret_cast<PFNGLRENDERBUFFERSTORAGEMULTISAMPLE>(
            gapic::GetGfxProcAddress("glRenderbufferStorageMultisample", false));
    glStencilMask =
            reinterpret_cast<PFNGLSTENCILMASK>(gapic::GetGfxProcAddress("glStencilMask", false));
    glStencilMaskSeparate = reinterpret_cast<PFNGLSTENCILMASKSEPARATE>(
            gapic::GetGfxProcAddress("glStencilMaskSeparate", false));
    glDisable = reinterpret_cast<PFNGLDISABLE>(gapic::GetGfxProcAddress("glDisable", false));
    glDisablei = reinterpret_cast<PFNGLDISABLEI>(gapic::GetGfxProcAddress("glDisablei", false));
    glEnable = reinterpret_cast<PFNGLENABLE>(gapic::GetGfxProcAddress("glEnable", false));
    glEnablei = reinterpret_cast<PFNGLENABLEI>(gapic::GetGfxProcAddress("glEnablei", false));
    glFinish = reinterpret_cast<PFNGLFINISH>(gapic::GetGfxProcAddress("glFinish", false));
    glFlush = reinterpret_cast<PFNGLFLUSH>(gapic::GetGfxProcAddress("glFlush", false));
    glFlushMappedBufferRange = reinterpret_cast<PFNGLFLUSHMAPPEDBUFFERRANGE>(
            gapic::GetGfxProcAddress("glFlushMappedBufferRange", false));
    glGetError = reinterpret_cast<PFNGLGETERROR>(gapic::GetGfxProcAddress("glGetError", false));
    glGetGraphicsResetStatus = reinterpret_cast<PFNGLGETGRAPHICSRESETSTATUS>(
            gapic::GetGfxProcAddress("glGetGraphicsResetStatus", false));
    glHint = reinterpret_cast<PFNGLHINT>(gapic::GetGfxProcAddress("glHint", false));
    glActiveShaderProgram = reinterpret_cast<PFNGLACTIVESHADERPROGRAM>(
            gapic::GetGfxProcAddress("glActiveShaderProgram", false));
    glAttachShader =
            reinterpret_cast<PFNGLATTACHSHADER>(gapic::GetGfxProcAddress("glAttachShader", false));
    glBindAttribLocation = reinterpret_cast<PFNGLBINDATTRIBLOCATION>(
            gapic::GetGfxProcAddress("glBindAttribLocation", false));
    glBindProgramPipeline = reinterpret_cast<PFNGLBINDPROGRAMPIPELINE>(
            gapic::GetGfxProcAddress("glBindProgramPipeline", false));
    glCompileShader = reinterpret_cast<PFNGLCOMPILESHADER>(
            gapic::GetGfxProcAddress("glCompileShader", false));
    glCreateProgram = reinterpret_cast<PFNGLCREATEPROGRAM>(
            gapic::GetGfxProcAddress("glCreateProgram", false));
    glCreateShader =
            reinterpret_cast<PFNGLCREATESHADER>(gapic::GetGfxProcAddress("glCreateShader", false));
    glCreateShaderProgramv = reinterpret_cast<PFNGLCREATESHADERPROGRAMV>(
            gapic::GetGfxProcAddress("glCreateShaderProgramv", false));
    glDeleteProgram = reinterpret_cast<PFNGLDELETEPROGRAM>(
            gapic::GetGfxProcAddress("glDeleteProgram", false));
    glDeleteProgramPipelines = reinterpret_cast<PFNGLDELETEPROGRAMPIPELINES>(
            gapic::GetGfxProcAddress("glDeleteProgramPipelines", false));
    glDeleteShader =
            reinterpret_cast<PFNGLDELETESHADER>(gapic::GetGfxProcAddress("glDeleteShader", false));
    glDetachShader =
            reinterpret_cast<PFNGLDETACHSHADER>(gapic::GetGfxProcAddress("glDetachShader", false));
    glDispatchCompute = reinterpret_cast<PFNGLDISPATCHCOMPUTE>(
            gapic::GetGfxProcAddress("glDispatchCompute", false));
    glDispatchComputeIndirect = reinterpret_cast<PFNGLDISPATCHCOMPUTEINDIRECT>(
            gapic::GetGfxProcAddress("glDispatchComputeIndirect", false));
    glGenProgramPipelines = reinterpret_cast<PFNGLGENPROGRAMPIPELINES>(
            gapic::GetGfxProcAddress("glGenProgramPipelines", false));
    glGetActiveAttrib = reinterpret_cast<PFNGLGETACTIVEATTRIB>(
            gapic::GetGfxProcAddress("glGetActiveAttrib", false));
    glGetActiveUniform = reinterpret_cast<PFNGLGETACTIVEUNIFORM>(
            gapic::GetGfxProcAddress("glGetActiveUniform", false));
    glGetActiveUniformBlockName = reinterpret_cast<PFNGLGETACTIVEUNIFORMBLOCKNAME>(
            gapic::GetGfxProcAddress("glGetActiveUniformBlockName", false));
    glGetActiveUniformBlockiv = reinterpret_cast<PFNGLGETACTIVEUNIFORMBLOCKIV>(
            gapic::GetGfxProcAddress("glGetActiveUniformBlockiv", false));
    glGetActiveUniformsiv = reinterpret_cast<PFNGLGETACTIVEUNIFORMSIV>(
            gapic::GetGfxProcAddress("glGetActiveUniformsiv", false));
    glGetAttachedShaders = reinterpret_cast<PFNGLGETATTACHEDSHADERS>(
            gapic::GetGfxProcAddress("glGetAttachedShaders", false));
    glGetAttribLocation = reinterpret_cast<PFNGLGETATTRIBLOCATION>(
            gapic::GetGfxProcAddress("glGetAttribLocation", false));
    glGetFragDataLocation = reinterpret_cast<PFNGLGETFRAGDATALOCATION>(
            gapic::GetGfxProcAddress("glGetFragDataLocation", false));
    glGetProgramBinary = reinterpret_cast<PFNGLGETPROGRAMBINARY>(
            gapic::GetGfxProcAddress("glGetProgramBinary", false));
    glGetProgramInfoLog = reinterpret_cast<PFNGLGETPROGRAMINFOLOG>(
            gapic::GetGfxProcAddress("glGetProgramInfoLog", false));
    glGetProgramInterfaceiv = reinterpret_cast<PFNGLGETPROGRAMINTERFACEIV>(
            gapic::GetGfxProcAddress("glGetProgramInterfaceiv", false));
    glGetProgramPipelineInfoLog = reinterpret_cast<PFNGLGETPROGRAMPIPELINEINFOLOG>(
            gapic::GetGfxProcAddress("glGetProgramPipelineInfoLog", false));
    glGetProgramPipelineiv = reinterpret_cast<PFNGLGETPROGRAMPIPELINEIV>(
            gapic::GetGfxProcAddress("glGetProgramPipelineiv", false));
    glGetProgramResourceIndex = reinterpret_cast<PFNGLGETPROGRAMRESOURCEINDEX>(
            gapic::GetGfxProcAddress("glGetProgramResourceIndex", false));
    glGetProgramResourceLocation = reinterpret_cast<PFNGLGETPROGRAMRESOURCELOCATION>(
            gapic::GetGfxProcAddress("glGetProgramResourceLocation", false));
    glGetProgramResourceName = reinterpret_cast<PFNGLGETPROGRAMRESOURCENAME>(
            gapic::GetGfxProcAddress("glGetProgramResourceName", false));
    glGetProgramResourceiv = reinterpret_cast<PFNGLGETPROGRAMRESOURCEIV>(
            gapic::GetGfxProcAddress("glGetProgramResourceiv", false));
    glGetProgramiv =
            reinterpret_cast<PFNGLGETPROGRAMIV>(gapic::GetGfxProcAddress("glGetProgramiv", false));
    glGetShaderInfoLog = reinterpret_cast<PFNGLGETSHADERINFOLOG>(
            gapic::GetGfxProcAddress("glGetShaderInfoLog", false));
    glGetShaderPrecisionFormat = reinterpret_cast<PFNGLGETSHADERPRECISIONFORMAT>(
            gapic::GetGfxProcAddress("glGetShaderPrecisionFormat", false));
    glGetShaderSource = reinterpret_cast<PFNGLGETSHADERSOURCE>(
            gapic::GetGfxProcAddress("glGetShaderSource", false));
    glGetShaderiv =
            reinterpret_cast<PFNGLGETSHADERIV>(gapic::GetGfxProcAddress("glGetShaderiv", false));
    glGetUniformBlockIndex = reinterpret_cast<PFNGLGETUNIFORMBLOCKINDEX>(
            gapic::GetGfxProcAddress("glGetUniformBlockIndex", false));
    glGetUniformIndices = reinterpret_cast<PFNGLGETUNIFORMINDICES>(
            gapic::GetGfxProcAddress("glGetUniformIndices", false));
    glGetUniformLocation = reinterpret_cast<PFNGLGETUNIFORMLOCATION>(
            gapic::GetGfxProcAddress("glGetUniformLocation", false));
    glGetUniformfv =
            reinterpret_cast<PFNGLGETUNIFORMFV>(gapic::GetGfxProcAddress("glGetUniformfv", false));
    glGetUniformiv =
            reinterpret_cast<PFNGLGETUNIFORMIV>(gapic::GetGfxProcAddress("glGetUniformiv", false));
    glGetUniformuiv = reinterpret_cast<PFNGLGETUNIFORMUIV>(
            gapic::GetGfxProcAddress("glGetUniformuiv", false));
    glGetnUniformfv = reinterpret_cast<PFNGLGETNUNIFORMFV>(
            gapic::GetGfxProcAddress("glGetnUniformfv", false));
    glGetnUniformiv = reinterpret_cast<PFNGLGETNUNIFORMIV>(
            gapic::GetGfxProcAddress("glGetnUniformiv", false));
    glGetnUniformuiv = reinterpret_cast<PFNGLGETNUNIFORMUIV>(
            gapic::GetGfxProcAddress("glGetnUniformuiv", false));
    glIsProgram = reinterpret_cast<PFNGLISPROGRAM>(gapic::GetGfxProcAddress("glIsProgram", false));
    glIsProgramPipeline = reinterpret_cast<PFNGLISPROGRAMPIPELINE>(
            gapic::GetGfxProcAddress("glIsProgramPipeline", false));
    glIsShader = reinterpret_cast<PFNGLISSHADER>(gapic::GetGfxProcAddress("glIsShader", false));
    glLinkProgram =
            reinterpret_cast<PFNGLLINKPROGRAM>(gapic::GetGfxProcAddress("glLinkProgram", false));
    glMemoryBarrier = reinterpret_cast<PFNGLMEMORYBARRIER>(
            gapic::GetGfxProcAddress("glMemoryBarrier", false));
    glMemoryBarrierByRegion = reinterpret_cast<PFNGLMEMORYBARRIERBYREGION>(
            gapic::GetGfxProcAddress("glMemoryBarrierByRegion", false));
    glProgramBinary = reinterpret_cast<PFNGLPROGRAMBINARY>(
            gapic::GetGfxProcAddress("glProgramBinary", false));
    glProgramParameteri = reinterpret_cast<PFNGLPROGRAMPARAMETERI>(
            gapic::GetGfxProcAddress("glProgramParameteri", false));
    glProgramUniform1f = reinterpret_cast<PFNGLPROGRAMUNIFORM1F>(
            gapic::GetGfxProcAddress("glProgramUniform1f", false));
    glProgramUniform1fv = reinterpret_cast<PFNGLPROGRAMUNIFORM1FV>(
            gapic::GetGfxProcAddress("glProgramUniform1fv", false));
    glProgramUniform1i = reinterpret_cast<PFNGLPROGRAMUNIFORM1I>(
            gapic::GetGfxProcAddress("glProgramUniform1i", false));
    glProgramUniform1iv = reinterpret_cast<PFNGLPROGRAMUNIFORM1IV>(
            gapic::GetGfxProcAddress("glProgramUniform1iv", false));
    glProgramUniform1ui = reinterpret_cast<PFNGLPROGRAMUNIFORM1UI>(
            gapic::GetGfxProcAddress("glProgramUniform1ui", false));
    glProgramUniform1uiv = reinterpret_cast<PFNGLPROGRAMUNIFORM1UIV>(
            gapic::GetGfxProcAddress("glProgramUniform1uiv", false));
    glProgramUniform2f = reinterpret_cast<PFNGLPROGRAMUNIFORM2F>(
            gapic::GetGfxProcAddress("glProgramUniform2f", false));
    glProgramUniform2fv = reinterpret_cast<PFNGLPROGRAMUNIFORM2FV>(
            gapic::GetGfxProcAddress("glProgramUniform2fv", false));
    glProgramUniform2i = reinterpret_cast<PFNGLPROGRAMUNIFORM2I>(
            gapic::GetGfxProcAddress("glProgramUniform2i", false));
    glProgramUniform2iv = reinterpret_cast<PFNGLPROGRAMUNIFORM2IV>(
            gapic::GetGfxProcAddress("glProgramUniform2iv", false));
    glProgramUniform2ui = reinterpret_cast<PFNGLPROGRAMUNIFORM2UI>(
            gapic::GetGfxProcAddress("glProgramUniform2ui", false));
    glProgramUniform2uiv = reinterpret_cast<PFNGLPROGRAMUNIFORM2UIV>(
            gapic::GetGfxProcAddress("glProgramUniform2uiv", false));
    glProgramUniform3f = reinterpret_cast<PFNGLPROGRAMUNIFORM3F>(
            gapic::GetGfxProcAddress("glProgramUniform3f", false));
    glProgramUniform3fv = reinterpret_cast<PFNGLPROGRAMUNIFORM3FV>(
            gapic::GetGfxProcAddress("glProgramUniform3fv", false));
    glProgramUniform3i = reinterpret_cast<PFNGLPROGRAMUNIFORM3I>(
            gapic::GetGfxProcAddress("glProgramUniform3i", false));
    glProgramUniform3iv = reinterpret_cast<PFNGLPROGRAMUNIFORM3IV>(
            gapic::GetGfxProcAddress("glProgramUniform3iv", false));
    glProgramUniform3ui = reinterpret_cast<PFNGLPROGRAMUNIFORM3UI>(
            gapic::GetGfxProcAddress("glProgramUniform3ui", false));
    glProgramUniform3uiv = reinterpret_cast<PFNGLPROGRAMUNIFORM3UIV>(
            gapic::GetGfxProcAddress("glProgramUniform3uiv", false));
    glProgramUniform4f = reinterpret_cast<PFNGLPROGRAMUNIFORM4F>(
            gapic::GetGfxProcAddress("glProgramUniform4f", false));
    glProgramUniform4fv = reinterpret_cast<PFNGLPROGRAMUNIFORM4FV>(
            gapic::GetGfxProcAddress("glProgramUniform4fv", false));
    glProgramUniform4i = reinterpret_cast<PFNGLPROGRAMUNIFORM4I>(
            gapic::GetGfxProcAddress("glProgramUniform4i", false));
    glProgramUniform4iv = reinterpret_cast<PFNGLPROGRAMUNIFORM4IV>(
            gapic::GetGfxProcAddress("glProgramUniform4iv", false));
    glProgramUniform4ui = reinterpret_cast<PFNGLPROGRAMUNIFORM4UI>(
            gapic::GetGfxProcAddress("glProgramUniform4ui", false));
    glProgramUniform4uiv = reinterpret_cast<PFNGLPROGRAMUNIFORM4UIV>(
            gapic::GetGfxProcAddress("glProgramUniform4uiv", false));
    glProgramUniformMatrix2fv = reinterpret_cast<PFNGLPROGRAMUNIFORMMATRIX2FV>(
            gapic::GetGfxProcAddress("glProgramUniformMatrix2fv", false));
    glProgramUniformMatrix2x3fv = reinterpret_cast<PFNGLPROGRAMUNIFORMMATRIX2X3FV>(
            gapic::GetGfxProcAddress("glProgramUniformMatrix2x3fv", false));
    glProgramUniformMatrix2x4fv = reinterpret_cast<PFNGLPROGRAMUNIFORMMATRIX2X4FV>(
            gapic::GetGfxProcAddress("glProgramUniformMatrix2x4fv", false));
    glProgramUniformMatrix3fv = reinterpret_cast<PFNGLPROGRAMUNIFORMMATRIX3FV>(
            gapic::GetGfxProcAddress("glProgramUniformMatrix3fv", false));
    glProgramUniformMatrix3x2fv = reinterpret_cast<PFNGLPROGRAMUNIFORMMATRIX3X2FV>(
            gapic::GetGfxProcAddress("glProgramUniformMatrix3x2fv", false));
    glProgramUniformMatrix3x4fv = reinterpret_cast<PFNGLPROGRAMUNIFORMMATRIX3X4FV>(
            gapic::GetGfxProcAddress("glProgramUniformMatrix3x4fv", false));
    glProgramUniformMatrix4fv = reinterpret_cast<PFNGLPROGRAMUNIFORMMATRIX4FV>(
            gapic::GetGfxProcAddress("glProgramUniformMatrix4fv", false));
    glProgramUniformMatrix4x2fv = reinterpret_cast<PFNGLPROGRAMUNIFORMMATRIX4X2FV>(
            gapic::GetGfxProcAddress("glProgramUniformMatrix4x2fv", false));
    glProgramUniformMatrix4x3fv = reinterpret_cast<PFNGLPROGRAMUNIFORMMATRIX4X3FV>(
            gapic::GetGfxProcAddress("glProgramUniformMatrix4x3fv", false));
    glReleaseShaderCompiler = reinterpret_cast<PFNGLRELEASESHADERCOMPILER>(
            gapic::GetGfxProcAddress("glReleaseShaderCompiler", false));
    glShaderBinary =
            reinterpret_cast<PFNGLSHADERBINARY>(gapic::GetGfxProcAddress("glShaderBinary", false));
    glShaderSource =
            reinterpret_cast<PFNGLSHADERSOURCE>(gapic::GetGfxProcAddress("glShaderSource", false));
    glUniform1f = reinterpret_cast<PFNGLUNIFORM1F>(gapic::GetGfxProcAddress("glUniform1f", false));
    glUniform1fv =
            reinterpret_cast<PFNGLUNIFORM1FV>(gapic::GetGfxProcAddress("glUniform1fv", false));
    glUniform1i = reinterpret_cast<PFNGLUNIFORM1I>(gapic::GetGfxProcAddress("glUniform1i", false));
    glUniform1iv =
            reinterpret_cast<PFNGLUNIFORM1IV>(gapic::GetGfxProcAddress("glUniform1iv", false));
    glUniform1ui =
            reinterpret_cast<PFNGLUNIFORM1UI>(gapic::GetGfxProcAddress("glUniform1ui", false));
    glUniform1uiv =
            reinterpret_cast<PFNGLUNIFORM1UIV>(gapic::GetGfxProcAddress("glUniform1uiv", false));
    glUniform2f = reinterpret_cast<PFNGLUNIFORM2F>(gapic::GetGfxProcAddress("glUniform2f", false));
    glUniform2fv =
            reinterpret_cast<PFNGLUNIFORM2FV>(gapic::GetGfxProcAddress("glUniform2fv", false));
    glUniform2i = reinterpret_cast<PFNGLUNIFORM2I>(gapic::GetGfxProcAddress("glUniform2i", false));
    glUniform2iv =
            reinterpret_cast<PFNGLUNIFORM2IV>(gapic::GetGfxProcAddress("glUniform2iv", false));
    glUniform2ui =
            reinterpret_cast<PFNGLUNIFORM2UI>(gapic::GetGfxProcAddress("glUniform2ui", false));
    glUniform2uiv =
            reinterpret_cast<PFNGLUNIFORM2UIV>(gapic::GetGfxProcAddress("glUniform2uiv", false));
    glUniform3f = reinterpret_cast<PFNGLUNIFORM3F>(gapic::GetGfxProcAddress("glUniform3f", false));
    glUniform3fv =
            reinterpret_cast<PFNGLUNIFORM3FV>(gapic::GetGfxProcAddress("glUniform3fv", false));
    glUniform3i = reinterpret_cast<PFNGLUNIFORM3I>(gapic::GetGfxProcAddress("glUniform3i", false));
    glUniform3iv =
            reinterpret_cast<PFNGLUNIFORM3IV>(gapic::GetGfxProcAddress("glUniform3iv", false));
    glUniform3ui =
            reinterpret_cast<PFNGLUNIFORM3UI>(gapic::GetGfxProcAddress("glUniform3ui", false));
    glUniform3uiv =
            reinterpret_cast<PFNGLUNIFORM3UIV>(gapic::GetGfxProcAddress("glUniform3uiv", false));
    glUniform4f = reinterpret_cast<PFNGLUNIFORM4F>(gapic::GetGfxProcAddress("glUniform4f", false));
    glUniform4fv =
            reinterpret_cast<PFNGLUNIFORM4FV>(gapic::GetGfxProcAddress("glUniform4fv", false));
    glUniform4i = reinterpret_cast<PFNGLUNIFORM4I>(gapic::GetGfxProcAddress("glUniform4i", false));
    glUniform4iv =
            reinterpret_cast<PFNGLUNIFORM4IV>(gapic::GetGfxProcAddress("glUniform4iv", false));
    glUniform4ui =
            reinterpret_cast<PFNGLUNIFORM4UI>(gapic::GetGfxProcAddress("glUniform4ui", false));
    glUniform4uiv =
            reinterpret_cast<PFNGLUNIFORM4UIV>(gapic::GetGfxProcAddress("glUniform4uiv", false));
    glUniformBlockBinding = reinterpret_cast<PFNGLUNIFORMBLOCKBINDING>(
            gapic::GetGfxProcAddress("glUniformBlockBinding", false));
    glUniformMatrix2fv = reinterpret_cast<PFNGLUNIFORMMATRIX2FV>(
            gapic::GetGfxProcAddress("glUniformMatrix2fv", false));
    glUniformMatrix2x3fv = reinterpret_cast<PFNGLUNIFORMMATRIX2X3FV>(
            gapic::GetGfxProcAddress("glUniformMatrix2x3fv", false));
    glUniformMatrix2x4fv = reinterpret_cast<PFNGLUNIFORMMATRIX2X4FV>(
            gapic::GetGfxProcAddress("glUniformMatrix2x4fv", false));
    glUniformMatrix3fv = reinterpret_cast<PFNGLUNIFORMMATRIX3FV>(
            gapic::GetGfxProcAddress("glUniformMatrix3fv", false));
    glUniformMatrix3x2fv = reinterpret_cast<PFNGLUNIFORMMATRIX3X2FV>(
            gapic::GetGfxProcAddress("glUniformMatrix3x2fv", false));
    glUniformMatrix3x4fv = reinterpret_cast<PFNGLUNIFORMMATRIX3X4FV>(
            gapic::GetGfxProcAddress("glUniformMatrix3x4fv", false));
    glUniformMatrix4fv = reinterpret_cast<PFNGLUNIFORMMATRIX4FV>(
            gapic::GetGfxProcAddress("glUniformMatrix4fv", false));
    glUniformMatrix4x2fv = reinterpret_cast<PFNGLUNIFORMMATRIX4X2FV>(
            gapic::GetGfxProcAddress("glUniformMatrix4x2fv", false));
    glUniformMatrix4x3fv = reinterpret_cast<PFNGLUNIFORMMATRIX4X3FV>(
            gapic::GetGfxProcAddress("glUniformMatrix4x3fv", false));
    glUseProgram =
            reinterpret_cast<PFNGLUSEPROGRAM>(gapic::GetGfxProcAddress("glUseProgram", false));
    glUseProgramStages = reinterpret_cast<PFNGLUSEPROGRAMSTAGES>(
            gapic::GetGfxProcAddress("glUseProgramStages", false));
    glValidateProgram = reinterpret_cast<PFNGLVALIDATEPROGRAM>(
            gapic::GetGfxProcAddress("glValidateProgram", false));
    glValidateProgramPipeline = reinterpret_cast<PFNGLVALIDATEPROGRAMPIPELINE>(
            gapic::GetGfxProcAddress("glValidateProgramPipeline", false));
    glCullFace = reinterpret_cast<PFNGLCULLFACE>(gapic::GetGfxProcAddress("glCullFace", false));
    glDepthRangef =
            reinterpret_cast<PFNGLDEPTHRANGEF>(gapic::GetGfxProcAddress("glDepthRangef", false));
    glFrontFace = reinterpret_cast<PFNGLFRONTFACE>(gapic::GetGfxProcAddress("glFrontFace", false));
    glGetMultisamplefv = reinterpret_cast<PFNGLGETMULTISAMPLEFV>(
            gapic::GetGfxProcAddress("glGetMultisamplefv", false));
    glLineWidth = reinterpret_cast<PFNGLLINEWIDTH>(gapic::GetGfxProcAddress("glLineWidth", false));
    glMinSampleShading = reinterpret_cast<PFNGLMINSAMPLESHADING>(
            gapic::GetGfxProcAddress("glMinSampleShading", false));
    glPolygonOffset = reinterpret_cast<PFNGLPOLYGONOFFSET>(
            gapic::GetGfxProcAddress("glPolygonOffset", false));
    glViewport = reinterpret_cast<PFNGLVIEWPORT>(gapic::GetGfxProcAddress("glViewport", false));
    glGetBooleani_v = reinterpret_cast<PFNGLGETBOOLEANI_V>(
            gapic::GetGfxProcAddress("glGetBooleani_v", false));
    glGetBooleanv =
            reinterpret_cast<PFNGLGETBOOLEANV>(gapic::GetGfxProcAddress("glGetBooleanv", false));
    glGetFloatv = reinterpret_cast<PFNGLGETFLOATV>(gapic::GetGfxProcAddress("glGetFloatv", false));
    glGetInteger64i_v = reinterpret_cast<PFNGLGETINTEGER64I_V>(
            gapic::GetGfxProcAddress("glGetInteger64i_v", false));
    glGetInteger64v = reinterpret_cast<PFNGLGETINTEGER64V>(
            gapic::GetGfxProcAddress("glGetInteger64v", false));
    glGetIntegeri_v = reinterpret_cast<PFNGLGETINTEGERI_V>(
            gapic::GetGfxProcAddress("glGetIntegeri_v", false));
    glGetIntegerv =
            reinterpret_cast<PFNGLGETINTEGERV>(gapic::GetGfxProcAddress("glGetIntegerv", false));
    glGetInternalformativ = reinterpret_cast<PFNGLGETINTERNALFORMATIV>(
            gapic::GetGfxProcAddress("glGetInternalformativ", false));
    glGetString = reinterpret_cast<PFNGLGETSTRING>(gapic::GetGfxProcAddress("glGetString", false));
    glGetStringi =
            reinterpret_cast<PFNGLGETSTRINGI>(gapic::GetGfxProcAddress("glGetStringi", false));
    glIsEnabled = reinterpret_cast<PFNGLISENABLED>(gapic::GetGfxProcAddress("glIsEnabled", false));
    glIsEnabledi =
            reinterpret_cast<PFNGLISENABLEDI>(gapic::GetGfxProcAddress("glIsEnabledi", false));
    glClientWaitSync = reinterpret_cast<PFNGLCLIENTWAITSYNC>(
            gapic::GetGfxProcAddress("glClientWaitSync", false));
    glDeleteSync =
            reinterpret_cast<PFNGLDELETESYNC>(gapic::GetGfxProcAddress("glDeleteSync", false));
    glFenceSync = reinterpret_cast<PFNGLFENCESYNC>(gapic::GetGfxProcAddress("glFenceSync", false));
    glGetSynciv = reinterpret_cast<PFNGLGETSYNCIV>(gapic::GetGfxProcAddress("glGetSynciv", false));
    glIsSync = reinterpret_cast<PFNGLISSYNC>(gapic::GetGfxProcAddress("glIsSync", false));
    glWaitSync = reinterpret_cast<PFNGLWAITSYNC>(gapic::GetGfxProcAddress("glWaitSync", false));
    glActiveTexture = reinterpret_cast<PFNGLACTIVETEXTURE>(
            gapic::GetGfxProcAddress("glActiveTexture", false));
    glBindImageTexture = reinterpret_cast<PFNGLBINDIMAGETEXTURE>(
            gapic::GetGfxProcAddress("glBindImageTexture", false));
    glBindSampler =
            reinterpret_cast<PFNGLBINDSAMPLER>(gapic::GetGfxProcAddress("glBindSampler", false));
    glBindTexture =
            reinterpret_cast<PFNGLBINDTEXTURE>(gapic::GetGfxProcAddress("glBindTexture", false));
    glCompressedTexImage2D = reinterpret_cast<PFNGLCOMPRESSEDTEXIMAGE2D>(
            gapic::GetGfxProcAddress("glCompressedTexImage2D", false));
    glCompressedTexImage3D = reinterpret_cast<PFNGLCOMPRESSEDTEXIMAGE3D>(
            gapic::GetGfxProcAddress("glCompressedTexImage3D", false));
    glCompressedTexSubImage2D = reinterpret_cast<PFNGLCOMPRESSEDTEXSUBIMAGE2D>(
            gapic::GetGfxProcAddress("glCompressedTexSubImage2D", false));
    glCompressedTexSubImage3D = reinterpret_cast<PFNGLCOMPRESSEDTEXSUBIMAGE3D>(
            gapic::GetGfxProcAddress("glCompressedTexSubImage3D", false));
    glCopyImageSubData = reinterpret_cast<PFNGLCOPYIMAGESUBDATA>(
            gapic::GetGfxProcAddress("glCopyImageSubData", false));
    glCopyTexImage2D = reinterpret_cast<PFNGLCOPYTEXIMAGE2D>(
            gapic::GetGfxProcAddress("glCopyTexImage2D", false));
    glCopyTexSubImage2D = reinterpret_cast<PFNGLCOPYTEXSUBIMAGE2D>(
            gapic::GetGfxProcAddress("glCopyTexSubImage2D", false));
    glCopyTexSubImage3D = reinterpret_cast<PFNGLCOPYTEXSUBIMAGE3D>(
            gapic::GetGfxProcAddress("glCopyTexSubImage3D", false));
    glDeleteSamplers = reinterpret_cast<PFNGLDELETESAMPLERS>(
            gapic::GetGfxProcAddress("glDeleteSamplers", false));
    glDeleteTextures = reinterpret_cast<PFNGLDELETETEXTURES>(
            gapic::GetGfxProcAddress("glDeleteTextures", false));
    glGenSamplers =
            reinterpret_cast<PFNGLGENSAMPLERS>(gapic::GetGfxProcAddress("glGenSamplers", false));
    glGenTextures =
            reinterpret_cast<PFNGLGENTEXTURES>(gapic::GetGfxProcAddress("glGenTextures", false));
    glGenerateMipmap = reinterpret_cast<PFNGLGENERATEMIPMAP>(
            gapic::GetGfxProcAddress("glGenerateMipmap", false));
    glGetSamplerParameterIiv = reinterpret_cast<PFNGLGETSAMPLERPARAMETERIIV>(
            gapic::GetGfxProcAddress("glGetSamplerParameterIiv", false));
    glGetSamplerParameterIuiv = reinterpret_cast<PFNGLGETSAMPLERPARAMETERIUIV>(
            gapic::GetGfxProcAddress("glGetSamplerParameterIuiv", false));
    glGetSamplerParameterfv = reinterpret_cast<PFNGLGETSAMPLERPARAMETERFV>(
            gapic::GetGfxProcAddress("glGetSamplerParameterfv", false));
    glGetSamplerParameteriv = reinterpret_cast<PFNGLGETSAMPLERPARAMETERIV>(
            gapic::GetGfxProcAddress("glGetSamplerParameteriv", false));
    glGetTexLevelParameterfv = reinterpret_cast<PFNGLGETTEXLEVELPARAMETERFV>(
            gapic::GetGfxProcAddress("glGetTexLevelParameterfv", false));
    glGetTexLevelParameteriv = reinterpret_cast<PFNGLGETTEXLEVELPARAMETERIV>(
            gapic::GetGfxProcAddress("glGetTexLevelParameteriv", false));
    glGetTexParameterIiv = reinterpret_cast<PFNGLGETTEXPARAMETERIIV>(
            gapic::GetGfxProcAddress("glGetTexParameterIiv", false));
    glGetTexParameterIuiv = reinterpret_cast<PFNGLGETTEXPARAMETERIUIV>(
            gapic::GetGfxProcAddress("glGetTexParameterIuiv", false));
    glGetTexParameterfv = reinterpret_cast<PFNGLGETTEXPARAMETERFV>(
            gapic::GetGfxProcAddress("glGetTexParameterfv", false));
    glGetTexParameteriv = reinterpret_cast<PFNGLGETTEXPARAMETERIV>(
            gapic::GetGfxProcAddress("glGetTexParameteriv", false));
    glIsSampler = reinterpret_cast<PFNGLISSAMPLER>(gapic::GetGfxProcAddress("glIsSampler", false));
    glIsTexture = reinterpret_cast<PFNGLISTEXTURE>(gapic::GetGfxProcAddress("glIsTexture", false));
    glPixelStorei =
            reinterpret_cast<PFNGLPIXELSTOREI>(gapic::GetGfxProcAddress("glPixelStorei", false));
    glSamplerParameterIiv = reinterpret_cast<PFNGLSAMPLERPARAMETERIIV>(
            gapic::GetGfxProcAddress("glSamplerParameterIiv", false));
    glSamplerParameterIuiv = reinterpret_cast<PFNGLSAMPLERPARAMETERIUIV>(
            gapic::GetGfxProcAddress("glSamplerParameterIuiv", false));
    glSamplerParameterf = reinterpret_cast<PFNGLSAMPLERPARAMETERF>(
            gapic::GetGfxProcAddress("glSamplerParameterf", false));
    glSamplerParameterfv = reinterpret_cast<PFNGLSAMPLERPARAMETERFV>(
            gapic::GetGfxProcAddress("glSamplerParameterfv", false));
    glSamplerParameteri = reinterpret_cast<PFNGLSAMPLERPARAMETERI>(
            gapic::GetGfxProcAddress("glSamplerParameteri", false));
    glSamplerParameteriv = reinterpret_cast<PFNGLSAMPLERPARAMETERIV>(
            gapic::GetGfxProcAddress("glSamplerParameteriv", false));
    glTexBuffer = reinterpret_cast<PFNGLTEXBUFFER>(gapic::GetGfxProcAddress("glTexBuffer", false));
    glTexBufferRange = reinterpret_cast<PFNGLTEXBUFFERRANGE>(
            gapic::GetGfxProcAddress("glTexBufferRange", false));
    glTexImage2D =
            reinterpret_cast<PFNGLTEXIMAGE2D>(gapic::GetGfxProcAddress("glTexImage2D", false));
    glTexImage3D =
            reinterpret_cast<PFNGLTEXIMAGE3D>(gapic::GetGfxProcAddress("glTexImage3D", false));
    glTexParameterIiv = reinterpret_cast<PFNGLTEXPARAMETERIIV>(
            gapic::GetGfxProcAddress("glTexParameterIiv", false));
    glTexParameterIuiv = reinterpret_cast<PFNGLTEXPARAMETERIUIV>(
            gapic::GetGfxProcAddress("glTexParameterIuiv", false));
    glTexParameterf = reinterpret_cast<PFNGLTEXPARAMETERF>(
            gapic::GetGfxProcAddress("glTexParameterf", false));
    glTexParameterfv = reinterpret_cast<PFNGLTEXPARAMETERFV>(
            gapic::GetGfxProcAddress("glTexParameterfv", false));
    glTexParameteri = reinterpret_cast<PFNGLTEXPARAMETERI>(
            gapic::GetGfxProcAddress("glTexParameteri", false));
    glTexParameteriv = reinterpret_cast<PFNGLTEXPARAMETERIV>(
            gapic::GetGfxProcAddress("glTexParameteriv", false));
    glTexStorage2D =
            reinterpret_cast<PFNGLTEXSTORAGE2D>(gapic::GetGfxProcAddress("glTexStorage2D", false));
    glTexStorage2DMultisample = reinterpret_cast<PFNGLTEXSTORAGE2DMULTISAMPLE>(
            gapic::GetGfxProcAddress("glTexStorage2DMultisample", false));
    glTexStorage3D =
            reinterpret_cast<PFNGLTEXSTORAGE3D>(gapic::GetGfxProcAddress("glTexStorage3D", false));
    glTexStorage3DMultisample = reinterpret_cast<PFNGLTEXSTORAGE3DMULTISAMPLE>(
            gapic::GetGfxProcAddress("glTexStorage3DMultisample", false));
    glTexSubImage2D = reinterpret_cast<PFNGLTEXSUBIMAGE2D>(
            gapic::GetGfxProcAddress("glTexSubImage2D", false));
    glTexSubImage3D = reinterpret_cast<PFNGLTEXSUBIMAGE3D>(
            gapic::GetGfxProcAddress("glTexSubImage3D", false));
    glBeginTransformFeedback = reinterpret_cast<PFNGLBEGINTRANSFORMFEEDBACK>(
            gapic::GetGfxProcAddress("glBeginTransformFeedback", false));
    glBindTransformFeedback = reinterpret_cast<PFNGLBINDTRANSFORMFEEDBACK>(
            gapic::GetGfxProcAddress("glBindTransformFeedback", false));
    glDeleteTransformFeedbacks = reinterpret_cast<PFNGLDELETETRANSFORMFEEDBACKS>(
            gapic::GetGfxProcAddress("glDeleteTransformFeedbacks", false));
    glEndTransformFeedback = reinterpret_cast<PFNGLENDTRANSFORMFEEDBACK>(
            gapic::GetGfxProcAddress("glEndTransformFeedback", false));
    glGenTransformFeedbacks = reinterpret_cast<PFNGLGENTRANSFORMFEEDBACKS>(
            gapic::GetGfxProcAddress("glGenTransformFeedbacks", false));
    glGetTransformFeedbackVarying = reinterpret_cast<PFNGLGETTRANSFORMFEEDBACKVARYING>(
            gapic::GetGfxProcAddress("glGetTransformFeedbackVarying", false));
    glIsTransformFeedback = reinterpret_cast<PFNGLISTRANSFORMFEEDBACK>(
            gapic::GetGfxProcAddress("glIsTransformFeedback", false));
    glPauseTransformFeedback = reinterpret_cast<PFNGLPAUSETRANSFORMFEEDBACK>(
            gapic::GetGfxProcAddress("glPauseTransformFeedback", false));
    glResumeTransformFeedback = reinterpret_cast<PFNGLRESUMETRANSFORMFEEDBACK>(
            gapic::GetGfxProcAddress("glResumeTransformFeedback", false));
    glTransformFeedbackVaryings = reinterpret_cast<PFNGLTRANSFORMFEEDBACKVARYINGS>(
            gapic::GetGfxProcAddress("glTransformFeedbackVaryings", false));
    glBindVertexArray = reinterpret_cast<PFNGLBINDVERTEXARRAY>(
            gapic::GetGfxProcAddress("glBindVertexArray", false));
    glBindVertexBuffer = reinterpret_cast<PFNGLBINDVERTEXBUFFER>(
            gapic::GetGfxProcAddress("glBindVertexBuffer", false));
    glDeleteVertexArrays = reinterpret_cast<PFNGLDELETEVERTEXARRAYS>(
            gapic::GetGfxProcAddress("glDeleteVertexArrays", false));
    glDisableVertexAttribArray = reinterpret_cast<PFNGLDISABLEVERTEXATTRIBARRAY>(
            gapic::GetGfxProcAddress("glDisableVertexAttribArray", false));
    glEnableVertexAttribArray = reinterpret_cast<PFNGLENABLEVERTEXATTRIBARRAY>(
            gapic::GetGfxProcAddress("glEnableVertexAttribArray", false));
    glGenVertexArrays = reinterpret_cast<PFNGLGENVERTEXARRAYS>(
            gapic::GetGfxProcAddress("glGenVertexArrays", false));
    glGetVertexAttribIiv = reinterpret_cast<PFNGLGETVERTEXATTRIBIIV>(
            gapic::GetGfxProcAddress("glGetVertexAttribIiv", false));
    glGetVertexAttribIuiv = reinterpret_cast<PFNGLGETVERTEXATTRIBIUIV>(
            gapic::GetGfxProcAddress("glGetVertexAttribIuiv", false));
    glGetVertexAttribPointerv = reinterpret_cast<PFNGLGETVERTEXATTRIBPOINTERV>(
            gapic::GetGfxProcAddress("glGetVertexAttribPointerv", false));
    glGetVertexAttribfv = reinterpret_cast<PFNGLGETVERTEXATTRIBFV>(
            gapic::GetGfxProcAddress("glGetVertexAttribfv", false));
    glGetVertexAttribiv = reinterpret_cast<PFNGLGETVERTEXATTRIBIV>(
            gapic::GetGfxProcAddress("glGetVertexAttribiv", false));
    glIsVertexArray = reinterpret_cast<PFNGLISVERTEXARRAY>(
            gapic::GetGfxProcAddress("glIsVertexArray", false));
    glVertexAttrib1f = reinterpret_cast<PFNGLVERTEXATTRIB1F>(
            gapic::GetGfxProcAddress("glVertexAttrib1f", false));
    glVertexAttrib1fv = reinterpret_cast<PFNGLVERTEXATTRIB1FV>(
            gapic::GetGfxProcAddress("glVertexAttrib1fv", false));
    glVertexAttrib2f = reinterpret_cast<PFNGLVERTEXATTRIB2F>(
            gapic::GetGfxProcAddress("glVertexAttrib2f", false));
    glVertexAttrib2fv = reinterpret_cast<PFNGLVERTEXATTRIB2FV>(
            gapic::GetGfxProcAddress("glVertexAttrib2fv", false));
    glVertexAttrib3f = reinterpret_cast<PFNGLVERTEXATTRIB3F>(
            gapic::GetGfxProcAddress("glVertexAttrib3f", false));
    glVertexAttrib3fv = reinterpret_cast<PFNGLVERTEXATTRIB3FV>(
            gapic::GetGfxProcAddress("glVertexAttrib3fv", false));
    glVertexAttrib4f = reinterpret_cast<PFNGLVERTEXATTRIB4F>(
            gapic::GetGfxProcAddress("glVertexAttrib4f", false));
    glVertexAttrib4fv = reinterpret_cast<PFNGLVERTEXATTRIB4FV>(
            gapic::GetGfxProcAddress("glVertexAttrib4fv", false));
    glVertexAttribBinding = reinterpret_cast<PFNGLVERTEXATTRIBBINDING>(
            gapic::GetGfxProcAddress("glVertexAttribBinding", false));
    glVertexAttribDivisor = reinterpret_cast<PFNGLVERTEXATTRIBDIVISOR>(
            gapic::GetGfxProcAddress("glVertexAttribDivisor", false));
    glVertexAttribFormat = reinterpret_cast<PFNGLVERTEXATTRIBFORMAT>(
            gapic::GetGfxProcAddress("glVertexAttribFormat", false));
    glVertexAttribI4i = reinterpret_cast<PFNGLVERTEXATTRIBI4I>(
            gapic::GetGfxProcAddress("glVertexAttribI4i", false));
    glVertexAttribI4iv = reinterpret_cast<PFNGLVERTEXATTRIBI4IV>(
            gapic::GetGfxProcAddress("glVertexAttribI4iv", false));
    glVertexAttribI4ui = reinterpret_cast<PFNGLVERTEXATTRIBI4UI>(
            gapic::GetGfxProcAddress("glVertexAttribI4ui", false));
    glVertexAttribI4uiv = reinterpret_cast<PFNGLVERTEXATTRIBI4UIV>(
            gapic::GetGfxProcAddress("glVertexAttribI4uiv", false));
    glVertexAttribIFormat = reinterpret_cast<PFNGLVERTEXATTRIBIFORMAT>(
            gapic::GetGfxProcAddress("glVertexAttribIFormat", false));
    glVertexAttribIPointer = reinterpret_cast<PFNGLVERTEXATTRIBIPOINTER>(
            gapic::GetGfxProcAddress("glVertexAttribIPointer", false));
    glVertexAttribPointer = reinterpret_cast<PFNGLVERTEXATTRIBPOINTER>(
            gapic::GetGfxProcAddress("glVertexAttribPointer", false));
    glVertexBindingDivisor = reinterpret_cast<PFNGLVERTEXBINDINGDIVISOR>(
            gapic::GetGfxProcAddress("glVertexBindingDivisor", false));
    eglInitialize =
            reinterpret_cast<PFNEGLINITIALIZE>(gapic::GetGfxProcAddress("eglInitialize", false));
    eglCreateContext = reinterpret_cast<PFNEGLCREATECONTEXT>(
            gapic::GetGfxProcAddress("eglCreateContext", false));
    eglMakeCurrent =
            reinterpret_cast<PFNEGLMAKECURRENT>(gapic::GetGfxProcAddress("eglMakeCurrent", false));
    eglSwapBuffers =
            reinterpret_cast<PFNEGLSWAPBUFFERS>(gapic::GetGfxProcAddress("eglSwapBuffers", false));
    eglQuerySurface = reinterpret_cast<PFNEGLQUERYSURFACE>(
            gapic::GetGfxProcAddress("eglQuerySurface", false));
    glXCreateContext = reinterpret_cast<PFNGLXCREATECONTEXT>(
            gapic::GetGfxProcAddress("glXCreateContext", false));
    glXCreateNewContext = reinterpret_cast<PFNGLXCREATENEWCONTEXT>(
            gapic::GetGfxProcAddress("glXCreateNewContext", false));
    glXMakeContextCurrent = reinterpret_cast<PFNGLXMAKECONTEXTCURRENT>(
            gapic::GetGfxProcAddress("glXMakeContextCurrent", false));
    glXMakeCurrent =
            reinterpret_cast<PFNGLXMAKECURRENT>(gapic::GetGfxProcAddress("glXMakeCurrent", false));
    glXSwapBuffers =
            reinterpret_cast<PFNGLXSWAPBUFFERS>(gapic::GetGfxProcAddress("glXSwapBuffers", false));
    glXQueryDrawable = reinterpret_cast<PFNGLXQUERYDRAWABLE>(
            gapic::GetGfxProcAddress("glXQueryDrawable", false));
    wglCreateContext = reinterpret_cast<PFNWGLCREATECONTEXT>(
            gapic::GetGfxProcAddress("wglCreateContext", false));
    wglCreateContextAttribsARB = reinterpret_cast<PFNWGLCREATECONTEXTATTRIBSARB>(
            gapic::GetGfxProcAddress("wglCreateContextAttribsARB", false));
    wglMakeCurrent =
            reinterpret_cast<PFNWGLMAKECURRENT>(gapic::GetGfxProcAddress("wglMakeCurrent", false));
    wglSwapBuffers =
            reinterpret_cast<PFNWGLSWAPBUFFERS>(gapic::GetGfxProcAddress("wglSwapBuffers", false));
    CGLCreateContext = reinterpret_cast<PFNCGLCREATECONTEXT>(
            gapic::GetGfxProcAddress("CGLCreateContext", false));
    CGLSetCurrentContext = reinterpret_cast<PFNCGLSETCURRENTCONTEXT>(
            gapic::GetGfxProcAddress("CGLSetCurrentContext", false));
    CGLGetSurface =
            reinterpret_cast<PFNCGLGETSURFACE>(gapic::GetGfxProcAddress("CGLGetSurface", false));
    CGSGetSurfaceBounds = reinterpret_cast<PFNCGSGETSURFACEBOUNDS>(
            gapic::GetGfxProcAddress("CGSGetSurfaceBounds", false));
    CGLFlushDrawable = reinterpret_cast<PFNCGLFLUSHDRAWABLE>(
            gapic::GetGfxProcAddress("CGLFlushDrawable", false));
    glGetQueryObjecti64v = reinterpret_cast<PFNGLGETQUERYOBJECTI64V>(
            gapic::GetGfxProcAddress("glGetQueryObjecti64v", false));
    glGetQueryObjectui64v = reinterpret_cast<PFNGLGETQUERYOBJECTUI64V>(
            gapic::GetGfxProcAddress("glGetQueryObjectui64v", false));
}

}  // namespace gfxapi
}  // namespace gapir
