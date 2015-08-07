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
        GAPID_INFO("glBlendBarrierKHR()\n");
        if (glBlendBarrierKHR != nullptr) {
            glBlendBarrierKHR();
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBlendBarrierKHR returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendBarrierKHR\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendBarrierKHR\n");
        return false;
    }
}

bool callGlBlendEquationSeparateiEXT(Stack* stack, bool pushReturn) {
    GLenum modeAlpha = stack->pop<GLenum>();
    GLenum modeRGB = stack->pop<GLenum>();
    uint32_t buf = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendEquationSeparateiEXT(%" PRIu32 ", %u, %u)\n", buf, modeRGB, modeAlpha);
        if (glBlendEquationSeparateiEXT != nullptr) {
            glBlendEquationSeparateiEXT(buf, modeRGB, modeAlpha);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBlendEquationSeparateiEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendEquationSeparateiEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendEquationSeparateiEXT\n");
        return false;
    }
}

bool callGlBlendEquationiEXT(Stack* stack, bool pushReturn) {
    GLenum mode = stack->pop<GLenum>();
    uint32_t buf = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendEquationiEXT(%" PRIu32 ", %u)\n", buf, mode);
        if (glBlendEquationiEXT != nullptr) {
            glBlendEquationiEXT(buf, mode);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBlendEquationiEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendEquationiEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendEquationiEXT\n");
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
        GAPID_INFO("glBlendFuncSeparateiEXT(%" PRIu32 ", %u, %u, %u, %u)\n", buf, srcRGB, dstRGB,
                   srcAlpha, dstAlpha);
        if (glBlendFuncSeparateiEXT != nullptr) {
            glBlendFuncSeparateiEXT(buf, srcRGB, dstRGB, srcAlpha, dstAlpha);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBlendFuncSeparateiEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendFuncSeparateiEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendFuncSeparateiEXT\n");
        return false;
    }
}

bool callGlBlendFunciEXT(Stack* stack, bool pushReturn) {
    GLenum dst = stack->pop<GLenum>();
    GLenum src = stack->pop<GLenum>();
    uint32_t buf = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendFunciEXT(%" PRIu32 ", %u, %u)\n", buf, src, dst);
        if (glBlendFunciEXT != nullptr) {
            glBlendFunciEXT(buf, src, dst);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBlendFunciEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendFunciEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendFunciEXT\n");
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
        GAPID_INFO("glColorMaskiEXT(%" PRIu32 ", %" PRIu8 ", %" PRIu8 ", %" PRIu8 ", %" PRIu8 ")\n",
                   index, r, g, b, a);
        if (glColorMaskiEXT != nullptr) {
            glColorMaskiEXT(index, r, g, b, a);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glColorMaskiEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glColorMaskiEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glColorMaskiEXT\n");
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
                   ", %" PRId32 ", %" PRId32 ", %" PRId32 ")\n",
                   srcName, srcTarget, srcLevel, srcX, srcY, srcZ, dstName, dstTarget, dstLevel,
                   dstX, dstY, dstZ, srcWidth, srcHeight, srcDepth);
        if (glCopyImageSubDataEXT != nullptr) {
            glCopyImageSubDataEXT(srcName, srcTarget, srcLevel, srcX, srcY, srcZ, dstName,
                                  dstTarget, dstLevel, dstX, dstY, dstZ, srcWidth, srcHeight,
                                  srcDepth);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glCopyImageSubDataEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCopyImageSubDataEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCopyImageSubDataEXT\n");
        return false;
    }
}

bool callGlDebugMessageCallbackKHR(Stack* stack, bool pushReturn) {
    void* userParam = stack->pop<void*>();
    void* callback = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("glDebugMessageCallbackKHR(%p, %p)\n", callback, userParam);
        if (glDebugMessageCallbackKHR != nullptr) {
            glDebugMessageCallbackKHR(callback, userParam);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDebugMessageCallbackKHR returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDebugMessageCallbackKHR\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDebugMessageCallbackKHR\n");
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
        GAPID_INFO("glDebugMessageControlKHR(%u, %u, %u, %" PRId32 ", %p, %" PRIu8 ")\n", source,
                   type, severity, count, ids, enabled);
        if (glDebugMessageControlKHR != nullptr) {
            glDebugMessageControlKHR(source, type, severity, count, ids, enabled);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDebugMessageControlKHR returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDebugMessageControlKHR\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDebugMessageControlKHR\n");
        return false;
    }
}

bool callGlDebugMessageInsertKHR(Stack* stack, bool pushReturn) {
    char* buf = stack->pop<char*>();
    int32_t length = stack->pop<int32_t>();
    GLenum severity = stack->pop<GLenum>();
    uint32_t id = stack->pop<uint32_t>();
    GLenum type = stack->pop<GLenum>();
    GLenum source = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDebugMessageInsertKHR(%u, %u, %" PRIu32 ", %u, %" PRId32 ", %p)\n", source,
                   type, id, severity, length, buf);
        if (glDebugMessageInsertKHR != nullptr) {
            glDebugMessageInsertKHR(source, type, id, severity, length, buf);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDebugMessageInsertKHR returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDebugMessageInsertKHR\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDebugMessageInsertKHR\n");
        return false;
    }
}

bool callGlDisableiEXT(Stack* stack, bool pushReturn) {
    uint32_t index = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDisableiEXT(%u, %" PRIu32 ")\n", target, index);
        if (glDisableiEXT != nullptr) {
            glDisableiEXT(target, index);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDisableiEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDisableiEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDisableiEXT\n");
        return false;
    }
}

bool callGlEnableiEXT(Stack* stack, bool pushReturn) {
    uint32_t index = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glEnableiEXT(%u, %" PRIu32 ")\n", target, index);
        if (glEnableiEXT != nullptr) {
            glEnableiEXT(target, index);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glEnableiEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEnableiEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEnableiEXT\n");
        return false;
    }
}

bool callGlFramebufferTextureEXT(Stack* stack, bool pushReturn) {
    int32_t level = stack->pop<int32_t>();
    uint32_t texture = stack->pop<uint32_t>();
    GLenum attachment = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glFramebufferTextureEXT(%u, %u, %" PRIu32 ", %" PRId32 ")\n", target,
                   attachment, texture, level);
        if (glFramebufferTextureEXT != nullptr) {
            glFramebufferTextureEXT(target, attachment, texture, level);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glFramebufferTextureEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFramebufferTextureEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFramebufferTextureEXT\n");
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
        GAPID_INFO("glGetDebugMessageLogKHR(%" PRIu32 ", %" PRId32 ", %p, %p, %p, %p, %p, %p)\n",
                   count, bufSize, sources, types, ids, severities, lengths, messageLog);
        if (glGetDebugMessageLogKHR != nullptr) {
            uint32_t return_value = glGetDebugMessageLogKHR(count, bufSize, sources, types, ids,
                                                            severities, lengths, messageLog);
            GAPID_INFO("Returned: %" PRIu32 "\n", return_value);
            if (pushReturn) {
                stack->push<uint32_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetDebugMessageLogKHR returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetDebugMessageLogKHR\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetDebugMessageLogKHR\n");
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
        GAPID_INFO("glGetObjectLabelKHR(%u, %" PRIu32 ", %" PRId32 ", %p, %p)\n", identifier, name,
                   bufSize, length, label);
        if (glGetObjectLabelKHR != nullptr) {
            glGetObjectLabelKHR(identifier, name, bufSize, length, label);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetObjectLabelKHR returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetObjectLabelKHR\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetObjectLabelKHR\n");
        return false;
    }
}

bool callGlGetObjectPtrLabelKHR(Stack* stack, bool pushReturn) {
    char* label = stack->pop<char*>();
    int32_t* length = stack->pop<int32_t*>();
    int32_t bufSize = stack->pop<int32_t>();
    void* ptr = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("glGetObjectPtrLabelKHR(%p, %" PRId32 ", %p, %p)\n", ptr, bufSize, length,
                   label);
        if (glGetObjectPtrLabelKHR != nullptr) {
            glGetObjectPtrLabelKHR(ptr, bufSize, length, label);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetObjectPtrLabelKHR returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetObjectPtrLabelKHR\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetObjectPtrLabelKHR\n");
        return false;
    }
}

bool callGlGetPointervKHR(Stack* stack, bool pushReturn) {
    void** params = stack->pop<void**>();
    GLenum pname = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetPointervKHR(%u, %p)\n", pname, params);
        if (glGetPointervKHR != nullptr) {
            glGetPointervKHR(pname, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetPointervKHR returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPointervKHR\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPointervKHR\n");
        return false;
    }
}

bool callGlGetSamplerParameterIivEXT(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetSamplerParameterIivEXT(%" PRIu32 ", %u, %p)\n", sampler, pname, params);
        if (glGetSamplerParameterIivEXT != nullptr) {
            glGetSamplerParameterIivEXT(sampler, pname, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetSamplerParameterIivEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetSamplerParameterIivEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetSamplerParameterIivEXT\n");
        return false;
    }
}

bool callGlGetSamplerParameterIuivEXT(Stack* stack, bool pushReturn) {
    uint32_t* params = stack->pop<uint32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetSamplerParameterIuivEXT(%" PRIu32 ", %u, %p)\n", sampler, pname, params);
        if (glGetSamplerParameterIuivEXT != nullptr) {
            glGetSamplerParameterIuivEXT(sampler, pname, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetSamplerParameterIuivEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetSamplerParameterIuivEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetSamplerParameterIuivEXT\n");
        return false;
    }
}

bool callGlGetTexParameterIivEXT(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetTexParameterIivEXT(%u, %u, %p)\n", target, pname, params);
        if (glGetTexParameterIivEXT != nullptr) {
            glGetTexParameterIivEXT(target, pname, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetTexParameterIivEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetTexParameterIivEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetTexParameterIivEXT\n");
        return false;
    }
}

bool callGlGetTexParameterIuivEXT(Stack* stack, bool pushReturn) {
    uint32_t* params = stack->pop<uint32_t*>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetTexParameterIuivEXT(%u, %u, %p)\n", target, pname, params);
        if (glGetTexParameterIuivEXT != nullptr) {
            glGetTexParameterIuivEXT(target, pname, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetTexParameterIuivEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetTexParameterIuivEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetTexParameterIuivEXT\n");
        return false;
    }
}

bool callGlIsEnablediEXT(Stack* stack, bool pushReturn) {
    uint32_t index = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glIsEnablediEXT(%u, %" PRIu32 ")\n", target, index);
        if (glIsEnablediEXT != nullptr) {
            uint8_t return_value = glIsEnablediEXT(target, index);
            GAPID_INFO("Returned: %" PRIu8 "\n", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glIsEnablediEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsEnablediEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsEnablediEXT\n");
        return false;
    }
}

bool callGlMinSampleShadingOES(Stack* stack, bool pushReturn) {
    float value = stack->pop<float>();
    if (stack->isValid()) {
        GAPID_INFO("glMinSampleShadingOES(%f)\n", value);
        if (glMinSampleShadingOES != nullptr) {
            glMinSampleShadingOES(value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glMinSampleShadingOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMinSampleShadingOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMinSampleShadingOES\n");
        return false;
    }
}

bool callGlObjectLabelKHR(Stack* stack, bool pushReturn) {
    char* label = stack->pop<char*>();
    int32_t length = stack->pop<int32_t>();
    uint32_t name = stack->pop<uint32_t>();
    GLenum identifier = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glObjectLabelKHR(%u, %" PRIu32 ", %" PRId32 ", %p)\n", identifier, name, length,
                   label);
        if (glObjectLabelKHR != nullptr) {
            glObjectLabelKHR(identifier, name, length, label);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glObjectLabelKHR returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glObjectLabelKHR\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glObjectLabelKHR\n");
        return false;
    }
}

bool callGlObjectPtrLabelKHR(Stack* stack, bool pushReturn) {
    char* label = stack->pop<char*>();
    int32_t length = stack->pop<int32_t>();
    void* ptr = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("glObjectPtrLabelKHR(%p, %" PRId32 ", %p)\n", ptr, length, label);
        if (glObjectPtrLabelKHR != nullptr) {
            glObjectPtrLabelKHR(ptr, length, label);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glObjectPtrLabelKHR returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glObjectPtrLabelKHR\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glObjectPtrLabelKHR\n");
        return false;
    }
}

bool callGlPatchParameteriEXT(Stack* stack, bool pushReturn) {
    int32_t value = stack->pop<int32_t>();
    GLenum pname = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glPatchParameteriEXT(%u, %" PRId32 ")\n", pname, value);
        if (glPatchParameteriEXT != nullptr) {
            glPatchParameteriEXT(pname, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glPatchParameteriEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPatchParameteriEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPatchParameteriEXT\n");
        return false;
    }
}

bool callGlPopDebugGroupKHR(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glPopDebugGroupKHR()\n");
        if (glPopDebugGroupKHR != nullptr) {
            glPopDebugGroupKHR();
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glPopDebugGroupKHR returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPopDebugGroupKHR\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPopDebugGroupKHR\n");
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
        GAPID_INFO("glPrimitiveBoundingBoxEXT(%f, %f, %f, %f, %f, %f, %f, %f)\n", minX, minY, minZ,
                   minW, maxX, maxY, maxZ, maxW);
        if (glPrimitiveBoundingBoxEXT != nullptr) {
            glPrimitiveBoundingBoxEXT(minX, minY, minZ, minW, maxX, maxY, maxZ, maxW);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glPrimitiveBoundingBoxEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPrimitiveBoundingBoxEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPrimitiveBoundingBoxEXT\n");
        return false;
    }
}

bool callGlPushDebugGroupKHR(Stack* stack, bool pushReturn) {
    char* message = stack->pop<char*>();
    int32_t length = stack->pop<int32_t>();
    uint32_t id = stack->pop<uint32_t>();
    GLenum source = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glPushDebugGroupKHR(%u, %" PRIu32 ", %" PRId32 ", %p)\n", source, id, length,
                   message);
        if (glPushDebugGroupKHR != nullptr) {
            glPushDebugGroupKHR(source, id, length, message);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glPushDebugGroupKHR returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPushDebugGroupKHR\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPushDebugGroupKHR\n");
        return false;
    }
}

bool callGlSamplerParameterIivEXT(Stack* stack, bool pushReturn) {
    int32_t* param = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glSamplerParameterIivEXT(%" PRIu32 ", %u, %p)\n", sampler, pname, param);
        if (glSamplerParameterIivEXT != nullptr) {
            glSamplerParameterIivEXT(sampler, pname, param);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glSamplerParameterIivEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glSamplerParameterIivEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glSamplerParameterIivEXT\n");
        return false;
    }
}

bool callGlSamplerParameterIuivEXT(Stack* stack, bool pushReturn) {
    uint32_t* param = stack->pop<uint32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glSamplerParameterIuivEXT(%" PRIu32 ", %u, %p)\n", sampler, pname, param);
        if (glSamplerParameterIuivEXT != nullptr) {
            glSamplerParameterIuivEXT(sampler, pname, param);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glSamplerParameterIuivEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glSamplerParameterIuivEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glSamplerParameterIuivEXT\n");
        return false;
    }
}

bool callGlTexBufferEXT(Stack* stack, bool pushReturn) {
    uint32_t buffer = stack->pop<uint32_t>();
    GLenum internalformat = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexBufferEXT(%u, %u, %" PRIu32 ")\n", target, internalformat, buffer);
        if (glTexBufferEXT != nullptr) {
            glTexBufferEXT(target, internalformat, buffer);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTexBufferEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexBufferEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexBufferEXT\n");
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
        GAPID_INFO("glTexBufferRangeEXT(%u, %u, %" PRIu32 ", %" PRId32 ", %" PRId32 ")\n", target,
                   internalformat, buffer, offset, size);
        if (glTexBufferRangeEXT != nullptr) {
            glTexBufferRangeEXT(target, internalformat, buffer, offset, size);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTexBufferRangeEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexBufferRangeEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexBufferRangeEXT\n");
        return false;
    }
}

bool callGlTexParameterIivEXT(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexParameterIivEXT(%u, %u, %p)\n", target, pname, params);
        if (glTexParameterIivEXT != nullptr) {
            glTexParameterIivEXT(target, pname, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTexParameterIivEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexParameterIivEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexParameterIivEXT\n");
        return false;
    }
}

bool callGlTexParameterIuivEXT(Stack* stack, bool pushReturn) {
    uint32_t* params = stack->pop<uint32_t*>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexParameterIuivEXT(%u, %u, %p)\n", target, pname, params);
        if (glTexParameterIuivEXT != nullptr) {
            glTexParameterIuivEXT(target, pname, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTexParameterIuivEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexParameterIuivEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexParameterIuivEXT\n");
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
                   ", %" PRId32 ", %" PRIu8 ")\n",
                   target, samples, internalformat, width, height, depth, fixedsamplelocations);
        if (glTexStorage3DMultisampleOES != nullptr) {
            glTexStorage3DMultisampleOES(target, samples, internalformat, width, height, depth,
                                         fixedsamplelocations);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTexStorage3DMultisampleOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexStorage3DMultisampleOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexStorage3DMultisampleOES\n");
        return false;
    }
}

bool callGlBeginQuery(Stack* stack, bool pushReturn) {
    uint32_t query = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBeginQuery(%u, %" PRIu32 ")\n", target, query);
        if (glBeginQuery != nullptr) {
            glBeginQuery(target, query);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBeginQuery returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBeginQuery\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBeginQuery\n");
        return false;
    }
}

bool callGlDeleteQueries(Stack* stack, bool pushReturn) {
    uint32_t* queries = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteQueries(%" PRId32 ", %p)\n", count, queries);
        if (glDeleteQueries != nullptr) {
            glDeleteQueries(count, queries);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDeleteQueries returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteQueries\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteQueries\n");
        return false;
    }
}

bool callGlEndQuery(Stack* stack, bool pushReturn) {
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glEndQuery(%u)\n", target);
        if (glEndQuery != nullptr) {
            glEndQuery(target);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glEndQuery returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEndQuery\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEndQuery\n");
        return false;
    }
}

bool callGlGenQueries(Stack* stack, bool pushReturn) {
    uint32_t* queries = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenQueries(%" PRId32 ", %p)\n", count, queries);
        if (glGenQueries != nullptr) {
            glGenQueries(count, queries);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGenQueries returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenQueries\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenQueries\n");
        return false;
    }
}

bool callGlGetQueryObjectuiv(Stack* stack, bool pushReturn) {
    uint32_t* value = stack->pop<uint32_t*>();
    GLenum parameter = stack->pop<GLenum>();
    uint32_t query = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetQueryObjectuiv(%" PRIu32 ", %u, %p)\n", query, parameter, value);
        if (glGetQueryObjectuiv != nullptr) {
            glGetQueryObjectuiv(query, parameter, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetQueryObjectuiv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetQueryObjectuiv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetQueryObjectuiv\n");
        return false;
    }
}

bool callGlGetQueryiv(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    GLenum parameter = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetQueryiv(%u, %u, %p)\n", target, parameter, value);
        if (glGetQueryiv != nullptr) {
            glGetQueryiv(target, parameter, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetQueryiv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetQueryiv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetQueryiv\n");
        return false;
    }
}

bool callGlIsQuery(Stack* stack, bool pushReturn) {
    uint32_t query = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsQuery(%" PRIu32 ")\n", query);
        if (glIsQuery != nullptr) {
            uint8_t return_value = glIsQuery(query);
            GAPID_INFO("Returned: %" PRIu8 "\n", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glIsQuery returned error: 0x%x\n", err);
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

bool callGlBindBuffer(Stack* stack, bool pushReturn) {
    uint32_t buffer = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBindBuffer(%u, %" PRIu32 ")\n", target, buffer);
        if (glBindBuffer != nullptr) {
            glBindBuffer(target, buffer);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBindBuffer returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindBuffer\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindBuffer\n");
        return false;
    }
}

bool callGlBindBufferBase(Stack* stack, bool pushReturn) {
    uint32_t buffer = stack->pop<uint32_t>();
    uint32_t index = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBindBufferBase(%u, %" PRIu32 ", %" PRIu32 ")\n", target, index, buffer);
        if (glBindBufferBase != nullptr) {
            glBindBufferBase(target, index, buffer);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBindBufferBase returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindBufferBase\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindBufferBase\n");
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
        GAPID_INFO("glBindBufferRange(%u, %" PRIu32 ", %" PRIu32 ", %" PRId32 ", %" PRId32 ")\n",
                   target, index, buffer, offset, size);
        if (glBindBufferRange != nullptr) {
            glBindBufferRange(target, index, buffer, offset, size);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBindBufferRange returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindBufferRange\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindBufferRange\n");
        return false;
    }
}

bool callGlBufferData(Stack* stack, bool pushReturn) {
    GLenum usage = stack->pop<GLenum>();
    void* data = stack->pop<void*>();
    int32_t size = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBufferData(%u, %" PRId32 ", %p, %u)\n", target, size, data, usage);
        if (glBufferData != nullptr) {
            glBufferData(target, size, data, usage);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBufferData returned error: 0x%x\n", err);
            }
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
    void* data = stack->pop<void*>();
    int32_t size = stack->pop<int32_t>();
    int32_t offset = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBufferSubData(%u, %" PRId32 ", %" PRId32 ", %p)\n", target, offset, size,
                   data);
        if (glBufferSubData != nullptr) {
            glBufferSubData(target, offset, size, data);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBufferSubData returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBufferSubData\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBufferSubData\n");
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
        GAPID_INFO("glCopyBufferSubData(%u, %u, %" PRId32 ", %" PRId32 ", %" PRId32 ")\n",
                   readTarget, writeTarget, readOffset, writeOffset, size);
        if (glCopyBufferSubData != nullptr) {
            glCopyBufferSubData(readTarget, writeTarget, readOffset, writeOffset, size);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glCopyBufferSubData returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCopyBufferSubData\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCopyBufferSubData\n");
        return false;
    }
}

bool callGlDeleteBuffers(Stack* stack, bool pushReturn) {
    uint32_t* buffers = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteBuffers(%" PRId32 ", %p)\n", count, buffers);
        if (glDeleteBuffers != nullptr) {
            glDeleteBuffers(count, buffers);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDeleteBuffers returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteBuffers\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteBuffers\n");
        return false;
    }
}

bool callGlGenBuffers(Stack* stack, bool pushReturn) {
    uint32_t* buffers = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenBuffers(%" PRId32 ", %p)\n", count, buffers);
        if (glGenBuffers != nullptr) {
            glGenBuffers(count, buffers);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGenBuffers returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenBuffers\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenBuffers\n");
        return false;
    }
}

bool callGlGetBufferParameteri64v(Stack* stack, bool pushReturn) {
    int64_t* params = stack->pop<int64_t*>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetBufferParameteri64v(%u, %u, %p)\n", target, pname, params);
        if (glGetBufferParameteri64v != nullptr) {
            glGetBufferParameteri64v(target, pname, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetBufferParameteri64v returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetBufferParameteri64v\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetBufferParameteri64v\n");
        return false;
    }
}

bool callGlGetBufferParameteriv(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    GLenum parameter = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetBufferParameteriv(%u, %u, %p)\n", target, parameter, value);
        if (glGetBufferParameteriv != nullptr) {
            glGetBufferParameteriv(target, parameter, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetBufferParameteriv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetBufferParameteriv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetBufferParameteriv\n");
        return false;
    }
}

bool callGlGetBufferPointerv(Stack* stack, bool pushReturn) {
    void** params = stack->pop<void**>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetBufferPointerv(%u, %u, %p)\n", target, pname, params);
        if (glGetBufferPointerv != nullptr) {
            glGetBufferPointerv(target, pname, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetBufferPointerv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetBufferPointerv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetBufferPointerv\n");
        return false;
    }
}

bool callGlIsBuffer(Stack* stack, bool pushReturn) {
    uint32_t buffer = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsBuffer(%" PRIu32 ")\n", buffer);
        if (glIsBuffer != nullptr) {
            uint8_t return_value = glIsBuffer(buffer);
            GAPID_INFO("Returned: %" PRIu8 "\n", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glIsBuffer returned error: 0x%x\n", err);
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

bool callGlMapBufferRange(Stack* stack, bool pushReturn) {
    GLbitfield access = stack->pop<GLbitfield>();
    int32_t length = stack->pop<int32_t>();
    int32_t offset = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glMapBufferRange(%u, %" PRId32 ", %" PRId32 ", %u)\n", target, offset, length,
                   access);
        if (glMapBufferRange != nullptr) {
            void* return_value = glMapBufferRange(target, offset, length, access);
            GAPID_INFO("Returned: %p\n", return_value);
            if (pushReturn) {
                stack->push<void*>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glMapBufferRange returned error: 0x%x\n", err);
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
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glUnmapBuffer(%u)\n", target);
        if (glUnmapBuffer != nullptr) {
            uint8_t return_value = glUnmapBuffer(target);
            GAPID_INFO("Returned: %" PRIu8 "\n", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUnmapBuffer returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUnmapBuffer\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUnmapBuffer\n");
        return false;
    }
}

bool callGlDrawArrays(Stack* stack, bool pushReturn) {
    int32_t index_count = stack->pop<int32_t>();
    int32_t first_index = stack->pop<int32_t>();
    GLenum draw_mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawArrays(%u, %" PRId32 ", %" PRId32 ")\n", draw_mode, first_index,
                   index_count);
        if (glDrawArrays != nullptr) {
            glDrawArrays(draw_mode, first_index, index_count);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDrawArrays returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawArrays\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawArrays\n");
        return false;
    }
}

bool callGlDrawArraysIndirect(Stack* stack, bool pushReturn) {
    void* indirect = stack->pop<void*>();
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawArraysIndirect(%u, %p)\n", mode, indirect);
        if (glDrawArraysIndirect != nullptr) {
            glDrawArraysIndirect(mode, indirect);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDrawArraysIndirect returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawArraysIndirect\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawArraysIndirect\n");
        return false;
    }
}

bool callGlDrawArraysInstanced(Stack* stack, bool pushReturn) {
    int32_t instancecount = stack->pop<int32_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t first = stack->pop<int32_t>();
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawArraysInstanced(%u, %" PRId32 ", %" PRId32 ", %" PRId32 ")\n", mode,
                   first, count, instancecount);
        if (glDrawArraysInstanced != nullptr) {
            glDrawArraysInstanced(mode, first, count, instancecount);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDrawArraysInstanced returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawArraysInstanced\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawArraysInstanced\n");
        return false;
    }
}

bool callGlDrawBuffers(Stack* stack, bool pushReturn) {
    GLenum* bufs = stack->pop<GLenum*>();
    int32_t n = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawBuffers(%" PRId32 ", %p)\n", n, bufs);
        if (glDrawBuffers != nullptr) {
            glDrawBuffers(n, bufs);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDrawBuffers returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawBuffers\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawBuffers\n");
        return false;
    }
}

bool callGlDrawElements(Stack* stack, bool pushReturn) {
    void* indices = stack->pop<void*>();
    GLenum indices_type = stack->pop<GLenum>();
    int32_t element_count = stack->pop<int32_t>();
    GLenum draw_mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawElements(%u, %" PRId32 ", %u, %p)\n", draw_mode, element_count,
                   indices_type, indices);
        if (glDrawElements != nullptr) {
            glDrawElements(draw_mode, element_count, indices_type, indices);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDrawElements returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawElements\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawElements\n");
        return false;
    }
}

bool callGlDrawElementsIndirect(Stack* stack, bool pushReturn) {
    void* indirect = stack->pop<void*>();
    GLenum type = stack->pop<GLenum>();
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawElementsIndirect(%u, %u, %p)\n", mode, type, indirect);
        if (glDrawElementsIndirect != nullptr) {
            glDrawElementsIndirect(mode, type, indirect);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDrawElementsIndirect returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawElementsIndirect\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawElementsIndirect\n");
        return false;
    }
}

bool callGlDrawElementsInstanced(Stack* stack, bool pushReturn) {
    int32_t instancecount = stack->pop<int32_t>();
    void* indices = stack->pop<void*>();
    GLenum type = stack->pop<GLenum>();
    int32_t count = stack->pop<int32_t>();
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawElementsInstanced(%u, %" PRId32 ", %u, %p, %" PRId32 ")\n", mode, count,
                   type, indices, instancecount);
        if (glDrawElementsInstanced != nullptr) {
            glDrawElementsInstanced(mode, count, type, indices, instancecount);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDrawElementsInstanced returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawElementsInstanced\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawElementsInstanced\n");
        return false;
    }
}

bool callGlDrawRangeElements(Stack* stack, bool pushReturn) {
    void* indices = stack->pop<void*>();
    GLenum type = stack->pop<GLenum>();
    int32_t count = stack->pop<int32_t>();
    uint32_t end = stack->pop<uint32_t>();
    uint32_t start = stack->pop<uint32_t>();
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawRangeElements(%u, %" PRIu32 ", %" PRIu32 ", %" PRId32 ", %u, %p)\n", mode,
                   start, end, count, type, indices);
        if (glDrawRangeElements != nullptr) {
            glDrawRangeElements(mode, start, end, count, type, indices);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDrawRangeElements returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawRangeElements\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawRangeElements\n");
        return false;
    }
}

bool callGlActiveShaderProgramEXT(Stack* stack, bool pushReturn) {
    uint32_t program = stack->pop<uint32_t>();
    uint32_t pipeline = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glActiveShaderProgramEXT(%" PRIu32 ", %" PRIu32 ")\n", pipeline, program);
        if (glActiveShaderProgramEXT != nullptr) {
            glActiveShaderProgramEXT(pipeline, program);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glActiveShaderProgramEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glActiveShaderProgramEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glActiveShaderProgramEXT\n");
        return false;
    }
}

bool callGlAlphaFuncQCOM(Stack* stack, bool pushReturn) {
    float ref = stack->pop<float>();
    GLenum func = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glAlphaFuncQCOM(%u, %f)\n", func, ref);
        if (glAlphaFuncQCOM != nullptr) {
            glAlphaFuncQCOM(func, ref);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glAlphaFuncQCOM returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glAlphaFuncQCOM\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glAlphaFuncQCOM\n");
        return false;
    }
}

bool callGlBeginConditionalRenderNV(Stack* stack, bool pushReturn) {
    GLenum mode = stack->pop<GLenum>();
    uint32_t id = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBeginConditionalRenderNV(%" PRIu32 ", %u)\n", id, mode);
        if (glBeginConditionalRenderNV != nullptr) {
            glBeginConditionalRenderNV(id, mode);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBeginConditionalRenderNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBeginConditionalRenderNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBeginConditionalRenderNV\n");
        return false;
    }
}

bool callGlBeginPerfMonitorAMD(Stack* stack, bool pushReturn) {
    uint32_t monitor = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBeginPerfMonitorAMD(%" PRIu32 ")\n", monitor);
        if (glBeginPerfMonitorAMD != nullptr) {
            glBeginPerfMonitorAMD(monitor);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBeginPerfMonitorAMD returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBeginPerfMonitorAMD\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBeginPerfMonitorAMD\n");
        return false;
    }
}

bool callGlBeginPerfQueryINTEL(Stack* stack, bool pushReturn) {
    uint32_t queryHandle = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBeginPerfQueryINTEL(%" PRIu32 ")\n", queryHandle);
        if (glBeginPerfQueryINTEL != nullptr) {
            glBeginPerfQueryINTEL(queryHandle);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBeginPerfQueryINTEL returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBeginPerfQueryINTEL\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBeginPerfQueryINTEL\n");
        return false;
    }
}

bool callGlBeginQueryEXT(Stack* stack, bool pushReturn) {
    uint32_t query = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBeginQueryEXT(%u, %" PRIu32 ")\n", target, query);
        if (glBeginQueryEXT != nullptr) {
            glBeginQueryEXT(target, query);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBeginQueryEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBeginQueryEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBeginQueryEXT\n");
        return false;
    }
}

bool callGlBindProgramPipelineEXT(Stack* stack, bool pushReturn) {
    uint32_t pipeline = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBindProgramPipelineEXT(%" PRIu32 ")\n", pipeline);
        if (glBindProgramPipelineEXT != nullptr) {
            glBindProgramPipelineEXT(pipeline);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBindProgramPipelineEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindProgramPipelineEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindProgramPipelineEXT\n");
        return false;
    }
}

bool callGlBindVertexArrayOES(Stack* stack, bool pushReturn) {
    uint32_t array = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBindVertexArrayOES(%" PRIu32 ")\n", array);
        if (glBindVertexArrayOES != nullptr) {
            glBindVertexArrayOES(array);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBindVertexArrayOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindVertexArrayOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindVertexArrayOES\n");
        return false;
    }
}

bool callGlBlendBarrierNV(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glBlendBarrierNV()\n");
        if (glBlendBarrierNV != nullptr) {
            glBlendBarrierNV();
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBlendBarrierNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendBarrierNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendBarrierNV\n");
        return false;
    }
}

bool callGlBlendEquationSeparateiOES(Stack* stack, bool pushReturn) {
    GLenum modeAlpha = stack->pop<GLenum>();
    GLenum modeRGB = stack->pop<GLenum>();
    uint32_t buf = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendEquationSeparateiOES(%" PRIu32 ", %u, %u)\n", buf, modeRGB, modeAlpha);
        if (glBlendEquationSeparateiOES != nullptr) {
            glBlendEquationSeparateiOES(buf, modeRGB, modeAlpha);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBlendEquationSeparateiOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendEquationSeparateiOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendEquationSeparateiOES\n");
        return false;
    }
}

bool callGlBlendEquationiOES(Stack* stack, bool pushReturn) {
    GLenum mode = stack->pop<GLenum>();
    uint32_t buf = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendEquationiOES(%" PRIu32 ", %u)\n", buf, mode);
        if (glBlendEquationiOES != nullptr) {
            glBlendEquationiOES(buf, mode);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBlendEquationiOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendEquationiOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendEquationiOES\n");
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
        GAPID_INFO("glBlendFuncSeparateiOES(%" PRIu32 ", %u, %u, %u, %u)\n", buf, srcRGB, dstRGB,
                   srcAlpha, dstAlpha);
        if (glBlendFuncSeparateiOES != nullptr) {
            glBlendFuncSeparateiOES(buf, srcRGB, dstRGB, srcAlpha, dstAlpha);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBlendFuncSeparateiOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendFuncSeparateiOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendFuncSeparateiOES\n");
        return false;
    }
}

bool callGlBlendFunciOES(Stack* stack, bool pushReturn) {
    GLenum dst = stack->pop<GLenum>();
    GLenum src = stack->pop<GLenum>();
    uint32_t buf = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendFunciOES(%" PRIu32 ", %u, %u)\n", buf, src, dst);
        if (glBlendFunciOES != nullptr) {
            glBlendFunciOES(buf, src, dst);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBlendFunciOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendFunciOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendFunciOES\n");
        return false;
    }
}

bool callGlBlendParameteriNV(Stack* stack, bool pushReturn) {
    int32_t value = stack->pop<int32_t>();
    GLenum pname = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendParameteriNV(%u, %" PRId32 ")\n", pname, value);
        if (glBlendParameteriNV != nullptr) {
            glBlendParameteriNV(pname, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBlendParameteriNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendParameteriNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendParameteriNV\n");
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
                   ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %u, %u)\n",
                   srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter);
        if (glBlitFramebufferANGLE != nullptr) {
            glBlitFramebufferANGLE(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask,
                                   filter);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBlitFramebufferANGLE returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlitFramebufferANGLE\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlitFramebufferANGLE\n");
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
                   ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %u, %u)\n",
                   srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter);
        if (glBlitFramebufferNV != nullptr) {
            glBlitFramebufferNV(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask,
                                filter);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBlitFramebufferNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlitFramebufferNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlitFramebufferNV\n");
        return false;
    }
}

bool callGlBufferStorageEXT(Stack* stack, bool pushReturn) {
    GLbitfield flag = stack->pop<GLbitfield>();
    void* data = stack->pop<void*>();
    int32_t size = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBufferStorageEXT(%u, %" PRId32 ", %p, %u)\n", target, size, data, flag);
        if (glBufferStorageEXT != nullptr) {
            glBufferStorageEXT(target, size, data, flag);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBufferStorageEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBufferStorageEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBufferStorageEXT\n");
        return false;
    }
}

bool callGlClientWaitSyncAPPLE(Stack* stack, bool pushReturn) {
    uint64_t timeout = stack->pop<uint64_t>();
    GLbitfield flag = stack->pop<GLbitfield>();
    uint64_t sync = stack->pop<uint64_t>();
    if (stack->isValid()) {
        GAPID_INFO("glClientWaitSyncAPPLE(%" PRIu64 ", %u, %" PRIu64 ")\n", sync, flag, timeout);
        if (glClientWaitSyncAPPLE != nullptr) {
            GLenum return_value = glClientWaitSyncAPPLE(sync, flag, timeout);
            GAPID_INFO("Returned: %u\n", return_value);
            if (pushReturn) {
                stack->push<GLenum>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glClientWaitSyncAPPLE returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glClientWaitSyncAPPLE\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glClientWaitSyncAPPLE\n");
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
        GAPID_INFO("glColorMaskiOES(%" PRIu32 ", %" PRIu8 ", %" PRIu8 ", %" PRIu8 ", %" PRIu8 ")\n",
                   index, r, g, b, a);
        if (glColorMaskiOES != nullptr) {
            glColorMaskiOES(index, r, g, b, a);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glColorMaskiOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glColorMaskiOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glColorMaskiOES\n");
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
                   ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %p)\n",
                   target, level, internalformat, width, height, depth, border, imageSize, data);
        if (glCompressedTexImage3DOES != nullptr) {
            glCompressedTexImage3DOES(target, level, internalformat, width, height, depth, border,
                                      imageSize, data);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glCompressedTexImage3DOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCompressedTexImage3DOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCompressedTexImage3DOES\n");
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
                   ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %u, %" PRId32 ", %p)\n",
                   target, level, xoffset, yoffset, zoffset, width, height, depth, format,
                   imageSize, data);
        if (glCompressedTexSubImage3DOES != nullptr) {
            glCompressedTexSubImage3DOES(target, level, xoffset, yoffset, zoffset, width, height,
                                         depth, format, imageSize, data);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glCompressedTexSubImage3DOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCompressedTexSubImage3DOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCompressedTexSubImage3DOES\n");
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
        GAPID_INFO("glCopyBufferSubDataNV(%u, %u, %" PRId32 ", %" PRId32 ", %" PRId32 ")\n",
                   readTarget, writeTarget, readOffset, writeOffset, size);
        if (glCopyBufferSubDataNV != nullptr) {
            glCopyBufferSubDataNV(readTarget, writeTarget, readOffset, writeOffset, size);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glCopyBufferSubDataNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCopyBufferSubDataNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCopyBufferSubDataNV\n");
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
                   ", %" PRId32 ", %" PRId32 ", %" PRId32 ")\n",
                   srcName, srcTarget, srcLevel, srcX, srcY, srcZ, dstName, dstTarget, dstLevel,
                   dstX, dstY, dstZ, srcWidth, srcHeight, srcDepth);
        if (glCopyImageSubDataOES != nullptr) {
            glCopyImageSubDataOES(srcName, srcTarget, srcLevel, srcX, srcY, srcZ, dstName,
                                  dstTarget, dstLevel, dstX, dstY, dstZ, srcWidth, srcHeight,
                                  srcDepth);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glCopyImageSubDataOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCopyImageSubDataOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCopyImageSubDataOES\n");
        return false;
    }
}

bool callGlCopyPathNV(Stack* stack, bool pushReturn) {
    uint32_t srcPath = stack->pop<uint32_t>();
    uint32_t resultPath = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glCopyPathNV(%" PRIu32 ", %" PRIu32 ")\n", resultPath, srcPath);
        if (glCopyPathNV != nullptr) {
            glCopyPathNV(resultPath, srcPath);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glCopyPathNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCopyPathNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCopyPathNV\n");
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
                   ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ")\n",
                   target, level, xoffset, yoffset, zoffset, x, y, width, height);
        if (glCopyTexSubImage3DOES != nullptr) {
            glCopyTexSubImage3DOES(target, level, xoffset, yoffset, zoffset, x, y, width, height);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glCopyTexSubImage3DOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCopyTexSubImage3DOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCopyTexSubImage3DOES\n");
        return false;
    }
}

bool callGlCopyTextureLevelsAPPLE(Stack* stack, bool pushReturn) {
    int32_t sourceLevelCount = stack->pop<int32_t>();
    int32_t sourceBaseLevel = stack->pop<int32_t>();
    uint32_t sourceTexture = stack->pop<uint32_t>();
    uint32_t destinationTexture = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glCopyTextureLevelsAPPLE(%" PRIu32 ", %" PRIu32 ", %" PRId32 ", %" PRId32 ")\n",
                   destinationTexture, sourceTexture, sourceBaseLevel, sourceLevelCount);
        if (glCopyTextureLevelsAPPLE != nullptr) {
            glCopyTextureLevelsAPPLE(destinationTexture, sourceTexture, sourceBaseLevel,
                                     sourceLevelCount);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glCopyTextureLevelsAPPLE returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCopyTextureLevelsAPPLE\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCopyTextureLevelsAPPLE\n");
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
        GAPID_INFO("glCoverFillPathInstancedNV(%" PRId32 ", %u, %p, %" PRIu32 ", %u, %u, %p)\n",
                   numPaths, pathNameType, paths, pathBase, coverMode, transformType,
                   transformValues);
        if (glCoverFillPathInstancedNV != nullptr) {
            glCoverFillPathInstancedNV(numPaths, pathNameType, paths, pathBase, coverMode,
                                       transformType, transformValues);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glCoverFillPathInstancedNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCoverFillPathInstancedNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCoverFillPathInstancedNV\n");
        return false;
    }
}

bool callGlCoverFillPathNV(Stack* stack, bool pushReturn) {
    GLenum coverMode = stack->pop<GLenum>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glCoverFillPathNV(%" PRIu32 ", %u)\n", path, coverMode);
        if (glCoverFillPathNV != nullptr) {
            glCoverFillPathNV(path, coverMode);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glCoverFillPathNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCoverFillPathNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCoverFillPathNV\n");
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
        GAPID_INFO("glCoverStrokePathInstancedNV(%" PRId32 ", %u, %p, %" PRIu32 ", %u, %u, %p)\n",
                   numPaths, pathNameType, paths, pathBase, coverMode, transformType,
                   transformValues);
        if (glCoverStrokePathInstancedNV != nullptr) {
            glCoverStrokePathInstancedNV(numPaths, pathNameType, paths, pathBase, coverMode,
                                         transformType, transformValues);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glCoverStrokePathInstancedNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCoverStrokePathInstancedNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCoverStrokePathInstancedNV\n");
        return false;
    }
}

bool callGlCoverStrokePathNV(Stack* stack, bool pushReturn) {
    GLenum coverMode = stack->pop<GLenum>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glCoverStrokePathNV(%" PRIu32 ", %u)\n", path, coverMode);
        if (glCoverStrokePathNV != nullptr) {
            glCoverStrokePathNV(path, coverMode);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glCoverStrokePathNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCoverStrokePathNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCoverStrokePathNV\n");
        return false;
    }
}

bool callGlCoverageMaskNV(Stack* stack, bool pushReturn) {
    uint8_t mask = stack->pop<uint8_t>();
    if (stack->isValid()) {
        GAPID_INFO("glCoverageMaskNV(%" PRIu8 ")\n", mask);
        if (glCoverageMaskNV != nullptr) {
            glCoverageMaskNV(mask);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glCoverageMaskNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCoverageMaskNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCoverageMaskNV\n");
        return false;
    }
}

bool callGlCoverageOperationNV(Stack* stack, bool pushReturn) {
    GLenum operation = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glCoverageOperationNV(%u)\n", operation);
        if (glCoverageOperationNV != nullptr) {
            glCoverageOperationNV(operation);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glCoverageOperationNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCoverageOperationNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCoverageOperationNV\n");
        return false;
    }
}

bool callGlCreatePerfQueryINTEL(Stack* stack, bool pushReturn) {
    uint32_t* queryHandle = stack->pop<uint32_t*>();
    uint32_t queryId = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glCreatePerfQueryINTEL(%" PRIu32 ", %p)\n", queryId, queryHandle);
        if (glCreatePerfQueryINTEL != nullptr) {
            glCreatePerfQueryINTEL(queryId, queryHandle);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glCreatePerfQueryINTEL returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCreatePerfQueryINTEL\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCreatePerfQueryINTEL\n");
        return false;
    }
}

bool callGlCreateShaderProgramvEXT(Stack* stack, bool pushReturn) {
    char** strings = stack->pop<char**>();
    int32_t count = stack->pop<int32_t>();
    GLenum type = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glCreateShaderProgramvEXT(%u, %" PRId32 ", %p)\n", type, count, strings);
        if (glCreateShaderProgramvEXT != nullptr) {
            uint32_t return_value = glCreateShaderProgramvEXT(type, count, strings);
            GAPID_INFO("Returned: %" PRIu32 "\n", return_value);
            if (pushReturn) {
                stack->push<uint32_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glCreateShaderProgramvEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCreateShaderProgramvEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCreateShaderProgramvEXT\n");
        return false;
    }
}

bool callGlDeleteFencesNV(Stack* stack, bool pushReturn) {
    uint32_t* fences = stack->pop<uint32_t*>();
    int32_t n = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteFencesNV(%" PRId32 ", %p)\n", n, fences);
        if (glDeleteFencesNV != nullptr) {
            glDeleteFencesNV(n, fences);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDeleteFencesNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteFencesNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteFencesNV\n");
        return false;
    }
}

bool callGlDeletePathsNV(Stack* stack, bool pushReturn) {
    int32_t range = stack->pop<int32_t>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeletePathsNV(%" PRIu32 ", %" PRId32 ")\n", path, range);
        if (glDeletePathsNV != nullptr) {
            glDeletePathsNV(path, range);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDeletePathsNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeletePathsNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeletePathsNV\n");
        return false;
    }
}

bool callGlDeletePerfMonitorsAMD(Stack* stack, bool pushReturn) {
    uint32_t* monitors = stack->pop<uint32_t*>();
    int32_t n = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeletePerfMonitorsAMD(%" PRId32 ", %p)\n", n, monitors);
        if (glDeletePerfMonitorsAMD != nullptr) {
            glDeletePerfMonitorsAMD(n, monitors);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDeletePerfMonitorsAMD returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeletePerfMonitorsAMD\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeletePerfMonitorsAMD\n");
        return false;
    }
}

bool callGlDeletePerfQueryINTEL(Stack* stack, bool pushReturn) {
    uint32_t queryHandle = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeletePerfQueryINTEL(%" PRIu32 ")\n", queryHandle);
        if (glDeletePerfQueryINTEL != nullptr) {
            glDeletePerfQueryINTEL(queryHandle);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDeletePerfQueryINTEL returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeletePerfQueryINTEL\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeletePerfQueryINTEL\n");
        return false;
    }
}

bool callGlDeleteProgramPipelinesEXT(Stack* stack, bool pushReturn) {
    uint32_t* pipelines = stack->pop<uint32_t*>();
    int32_t n = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteProgramPipelinesEXT(%" PRId32 ", %p)\n", n, pipelines);
        if (glDeleteProgramPipelinesEXT != nullptr) {
            glDeleteProgramPipelinesEXT(n, pipelines);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDeleteProgramPipelinesEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteProgramPipelinesEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteProgramPipelinesEXT\n");
        return false;
    }
}

bool callGlDeleteQueriesEXT(Stack* stack, bool pushReturn) {
    uint32_t* queries = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteQueriesEXT(%" PRId32 ", %p)\n", count, queries);
        if (glDeleteQueriesEXT != nullptr) {
            glDeleteQueriesEXT(count, queries);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDeleteQueriesEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteQueriesEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteQueriesEXT\n");
        return false;
    }
}

bool callGlDeleteSyncAPPLE(Stack* stack, bool pushReturn) {
    uint64_t sync = stack->pop<uint64_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteSyncAPPLE(%" PRIu64 ")\n", sync);
        if (glDeleteSyncAPPLE != nullptr) {
            glDeleteSyncAPPLE(sync);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDeleteSyncAPPLE returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteSyncAPPLE\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteSyncAPPLE\n");
        return false;
    }
}

bool callGlDeleteVertexArraysOES(Stack* stack, bool pushReturn) {
    uint32_t* arrays = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteVertexArraysOES(%" PRId32 ", %p)\n", count, arrays);
        if (glDeleteVertexArraysOES != nullptr) {
            glDeleteVertexArraysOES(count, arrays);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDeleteVertexArraysOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteVertexArraysOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteVertexArraysOES\n");
        return false;
    }
}

bool callGlDepthRangeArrayfvNV(Stack* stack, bool pushReturn) {
    float* v = stack->pop<float*>();
    int32_t count = stack->pop<int32_t>();
    uint32_t first = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDepthRangeArrayfvNV(%" PRIu32 ", %" PRId32 ", %p)\n", first, count, v);
        if (glDepthRangeArrayfvNV != nullptr) {
            glDepthRangeArrayfvNV(first, count, v);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDepthRangeArrayfvNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDepthRangeArrayfvNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDepthRangeArrayfvNV\n");
        return false;
    }
}

bool callGlDepthRangeIndexedfNV(Stack* stack, bool pushReturn) {
    float f = stack->pop<float>();
    float n = stack->pop<float>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDepthRangeIndexedfNV(%" PRIu32 ", %f, %f)\n", index, n, f);
        if (glDepthRangeIndexedfNV != nullptr) {
            glDepthRangeIndexedfNV(index, n, f);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDepthRangeIndexedfNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDepthRangeIndexedfNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDepthRangeIndexedfNV\n");
        return false;
    }
}

bool callGlDisableDriverControlQCOM(Stack* stack, bool pushReturn) {
    uint32_t driverControl = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDisableDriverControlQCOM(%" PRIu32 ")\n", driverControl);
        if (glDisableDriverControlQCOM != nullptr) {
            glDisableDriverControlQCOM(driverControl);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDisableDriverControlQCOM returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDisableDriverControlQCOM\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDisableDriverControlQCOM\n");
        return false;
    }
}

bool callGlDisableiNV(Stack* stack, bool pushReturn) {
    uint32_t index = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDisableiNV(%u, %" PRIu32 ")\n", target, index);
        if (glDisableiNV != nullptr) {
            glDisableiNV(target, index);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDisableiNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDisableiNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDisableiNV\n");
        return false;
    }
}

bool callGlDisableiOES(Stack* stack, bool pushReturn) {
    uint32_t index = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDisableiOES(%u, %" PRIu32 ")\n", target, index);
        if (glDisableiOES != nullptr) {
            glDisableiOES(target, index);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDisableiOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDisableiOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDisableiOES\n");
        return false;
    }
}

bool callGlDiscardFramebufferEXT(Stack* stack, bool pushReturn) {
    GLenum* attachments = stack->pop<GLenum*>();
    int32_t numAttachments = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDiscardFramebufferEXT(%u, %" PRId32 ", %p)\n", target, numAttachments,
                   attachments);
        if (glDiscardFramebufferEXT != nullptr) {
            glDiscardFramebufferEXT(target, numAttachments, attachments);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDiscardFramebufferEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDiscardFramebufferEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDiscardFramebufferEXT\n");
        return false;
    }
}

bool callGlDrawArraysInstancedANGLE(Stack* stack, bool pushReturn) {
    int32_t primcount = stack->pop<int32_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t first = stack->pop<int32_t>();
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawArraysInstancedANGLE(%u, %" PRId32 ", %" PRId32 ", %" PRId32 ")\n", mode,
                   first, count, primcount);
        if (glDrawArraysInstancedANGLE != nullptr) {
            glDrawArraysInstancedANGLE(mode, first, count, primcount);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDrawArraysInstancedANGLE returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawArraysInstancedANGLE\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawArraysInstancedANGLE\n");
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
                   ", %" PRIu32 ")\n",
                   mode, first, count, instancecount, baseinstance);
        if (glDrawArraysInstancedBaseInstanceEXT != nullptr) {
            glDrawArraysInstancedBaseInstanceEXT(mode, first, count, instancecount, baseinstance);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDrawArraysInstancedBaseInstanceEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function "
                    "glDrawArraysInstancedBaseInstanceEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawArraysInstancedBaseInstanceEXT\n");
        return false;
    }
}

bool callGlDrawArraysInstancedEXT(Stack* stack, bool pushReturn) {
    int32_t primcount = stack->pop<int32_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t start = stack->pop<int32_t>();
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawArraysInstancedEXT(%u, %" PRId32 ", %" PRId32 ", %" PRId32 ")\n", mode,
                   start, count, primcount);
        if (glDrawArraysInstancedEXT != nullptr) {
            glDrawArraysInstancedEXT(mode, start, count, primcount);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDrawArraysInstancedEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawArraysInstancedEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawArraysInstancedEXT\n");
        return false;
    }
}

bool callGlDrawArraysInstancedNV(Stack* stack, bool pushReturn) {
    int32_t primcount = stack->pop<int32_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t first = stack->pop<int32_t>();
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawArraysInstancedNV(%u, %" PRId32 ", %" PRId32 ", %" PRId32 ")\n", mode,
                   first, count, primcount);
        if (glDrawArraysInstancedNV != nullptr) {
            glDrawArraysInstancedNV(mode, first, count, primcount);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDrawArraysInstancedNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawArraysInstancedNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawArraysInstancedNV\n");
        return false;
    }
}

bool callGlDrawBuffersEXT(Stack* stack, bool pushReturn) {
    GLenum* bufs = stack->pop<GLenum*>();
    int32_t n = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawBuffersEXT(%" PRId32 ", %p)\n", n, bufs);
        if (glDrawBuffersEXT != nullptr) {
            glDrawBuffersEXT(n, bufs);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDrawBuffersEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawBuffersEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawBuffersEXT\n");
        return false;
    }
}

bool callGlDrawBuffersIndexedEXT(Stack* stack, bool pushReturn) {
    int32_t* indices = stack->pop<int32_t*>();
    GLenum* location = stack->pop<GLenum*>();
    int32_t n = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawBuffersIndexedEXT(%" PRId32 ", %p, %p)\n", n, location, indices);
        if (glDrawBuffersIndexedEXT != nullptr) {
            glDrawBuffersIndexedEXT(n, location, indices);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDrawBuffersIndexedEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawBuffersIndexedEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawBuffersIndexedEXT\n");
        return false;
    }
}

bool callGlDrawBuffersNV(Stack* stack, bool pushReturn) {
    GLenum* bufs = stack->pop<GLenum*>();
    int32_t n = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDrawBuffersNV(%" PRId32 ", %p)\n", n, bufs);
        if (glDrawBuffersNV != nullptr) {
            glDrawBuffersNV(n, bufs);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDrawBuffersNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawBuffersNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawBuffersNV\n");
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
        GAPID_INFO("glDrawElementsBaseVertexEXT(%u, %" PRId32 ", %u, %p, %" PRId32 ")\n", mode,
                   count, type, indices, basevertex);
        if (glDrawElementsBaseVertexEXT != nullptr) {
            glDrawElementsBaseVertexEXT(mode, count, type, indices, basevertex);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDrawElementsBaseVertexEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawElementsBaseVertexEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawElementsBaseVertexEXT\n");
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
        GAPID_INFO("glDrawElementsBaseVertexOES(%u, %" PRId32 ", %u, %p, %" PRId32 ")\n", mode,
                   count, type, indices, basevertex);
        if (glDrawElementsBaseVertexOES != nullptr) {
            glDrawElementsBaseVertexOES(mode, count, type, indices, basevertex);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDrawElementsBaseVertexOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawElementsBaseVertexOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawElementsBaseVertexOES\n");
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
        GAPID_INFO("glDrawElementsInstancedANGLE(%u, %" PRId32 ", %u, %p, %" PRId32 ")\n", mode,
                   count, type, indices, primcount);
        if (glDrawElementsInstancedANGLE != nullptr) {
            glDrawElementsInstancedANGLE(mode, count, type, indices, primcount);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDrawElementsInstancedANGLE returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawElementsInstancedANGLE\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawElementsInstancedANGLE\n");
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
                   ", %" PRIu32 ")\n",
                   mode, count, type, indices, instancecount, baseinstance);
        if (glDrawElementsInstancedBaseInstanceEXT != nullptr) {
            glDrawElementsInstancedBaseInstanceEXT(mode, count, type, indices, instancecount,
                                                   baseinstance);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDrawElementsInstancedBaseInstanceEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function "
                    "glDrawElementsInstancedBaseInstanceEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawElementsInstancedBaseInstanceEXT\n");
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
                   ", %" PRId32 ")\n",
                   mode, count, type, indices, instancecount, basevertex);
        if (glDrawElementsInstancedBaseVertexEXT != nullptr) {
            glDrawElementsInstancedBaseVertexEXT(mode, count, type, indices, instancecount,
                                                 basevertex);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDrawElementsInstancedBaseVertexEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function "
                    "glDrawElementsInstancedBaseVertexEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawElementsInstancedBaseVertexEXT\n");
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
                   ", %" PRId32 ")\n",
                   mode, count, type, indices, instancecount, basevertex);
        if (glDrawElementsInstancedBaseVertexOES != nullptr) {
            glDrawElementsInstancedBaseVertexOES(mode, count, type, indices, instancecount,
                                                 basevertex);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDrawElementsInstancedBaseVertexOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function "
                    "glDrawElementsInstancedBaseVertexOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawElementsInstancedBaseVertexOES\n");
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
                   ", %u, %p, %" PRId32 ", %" PRId32 ", %" PRIu32 ")\n",
                   mode, count, type, indices, instancecount, basevertex, baseinstance);
        if (glDrawElementsInstancedBaseVertexBaseInstanceEXT != nullptr) {
            glDrawElementsInstancedBaseVertexBaseInstanceEXT(
                    mode, count, type, indices, instancecount, basevertex, baseinstance);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING(
                        "glDrawElementsInstancedBaseVertexBaseInstanceEXT returned error: 0x%x\n",
                        err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function "
                    "glDrawElementsInstancedBaseVertexBaseInstanceEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING(
                "Error during calling function glDrawElementsInstancedBaseVertexBaseInstanceEXT\n");
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
        GAPID_INFO("glDrawElementsInstancedEXT(%u, %" PRId32 ", %u, %p, %" PRId32 ")\n", mode,
                   count, type, indices, primcount);
        if (glDrawElementsInstancedEXT != nullptr) {
            glDrawElementsInstancedEXT(mode, count, type, indices, primcount);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDrawElementsInstancedEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawElementsInstancedEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawElementsInstancedEXT\n");
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
        GAPID_INFO("glDrawElementsInstancedNV(%u, %" PRId32 ", %u, %p, %" PRId32 ")\n", mode, count,
                   type, indices, primcount);
        if (glDrawElementsInstancedNV != nullptr) {
            glDrawElementsInstancedNV(mode, count, type, indices, primcount);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDrawElementsInstancedNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDrawElementsInstancedNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawElementsInstancedNV\n");
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
                   ", %u, %p, %" PRId32 ")\n",
                   mode, start, end, count, type, indices, basevertex);
        if (glDrawRangeElementsBaseVertexEXT != nullptr) {
            glDrawRangeElementsBaseVertexEXT(mode, start, end, count, type, indices, basevertex);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDrawRangeElementsBaseVertexEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glDrawRangeElementsBaseVertexEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawRangeElementsBaseVertexEXT\n");
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
                   ", %u, %p, %" PRId32 ")\n",
                   mode, start, end, count, type, indices, basevertex);
        if (glDrawRangeElementsBaseVertexOES != nullptr) {
            glDrawRangeElementsBaseVertexOES(mode, start, end, count, type, indices, basevertex);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDrawRangeElementsBaseVertexOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glDrawRangeElementsBaseVertexOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDrawRangeElementsBaseVertexOES\n");
        return false;
    }
}

bool callGlEGLImageTargetRenderbufferStorageOES(Stack* stack, bool pushReturn) {
    void* image = stack->pop<void*>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glEGLImageTargetRenderbufferStorageOES(%u, %p)\n", target, image);
        if (glEGLImageTargetRenderbufferStorageOES != nullptr) {
            glEGLImageTargetRenderbufferStorageOES(target, image);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glEGLImageTargetRenderbufferStorageOES returned error: 0x%x\n", err);
            }
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

bool callGlEGLImageTargetTexture2DOES(Stack* stack, bool pushReturn) {
    void* image = stack->pop<void*>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glEGLImageTargetTexture2DOES(%u, %p)\n", target, image);
        if (glEGLImageTargetTexture2DOES != nullptr) {
            glEGLImageTargetTexture2DOES(target, image);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glEGLImageTargetTexture2DOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEGLImageTargetTexture2DOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEGLImageTargetTexture2DOES\n");
        return false;
    }
}

bool callGlEnableDriverControlQCOM(Stack* stack, bool pushReturn) {
    uint32_t driverControl = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glEnableDriverControlQCOM(%" PRIu32 ")\n", driverControl);
        if (glEnableDriverControlQCOM != nullptr) {
            glEnableDriverControlQCOM(driverControl);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glEnableDriverControlQCOM returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEnableDriverControlQCOM\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEnableDriverControlQCOM\n");
        return false;
    }
}

bool callGlEnableiNV(Stack* stack, bool pushReturn) {
    uint32_t index = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glEnableiNV(%u, %" PRIu32 ")\n", target, index);
        if (glEnableiNV != nullptr) {
            glEnableiNV(target, index);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glEnableiNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEnableiNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEnableiNV\n");
        return false;
    }
}

bool callGlEnableiOES(Stack* stack, bool pushReturn) {
    uint32_t index = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glEnableiOES(%u, %" PRIu32 ")\n", target, index);
        if (glEnableiOES != nullptr) {
            glEnableiOES(target, index);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glEnableiOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEnableiOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEnableiOES\n");
        return false;
    }
}

bool callGlEndConditionalRenderNV(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glEndConditionalRenderNV()\n");
        if (glEndConditionalRenderNV != nullptr) {
            glEndConditionalRenderNV();
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glEndConditionalRenderNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEndConditionalRenderNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEndConditionalRenderNV\n");
        return false;
    }
}

bool callGlEndPerfMonitorAMD(Stack* stack, bool pushReturn) {
    uint32_t monitor = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glEndPerfMonitorAMD(%" PRIu32 ")\n", monitor);
        if (glEndPerfMonitorAMD != nullptr) {
            glEndPerfMonitorAMD(monitor);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glEndPerfMonitorAMD returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEndPerfMonitorAMD\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEndPerfMonitorAMD\n");
        return false;
    }
}

bool callGlEndPerfQueryINTEL(Stack* stack, bool pushReturn) {
    uint32_t queryHandle = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glEndPerfQueryINTEL(%" PRIu32 ")\n", queryHandle);
        if (glEndPerfQueryINTEL != nullptr) {
            glEndPerfQueryINTEL(queryHandle);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glEndPerfQueryINTEL returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEndPerfQueryINTEL\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEndPerfQueryINTEL\n");
        return false;
    }
}

bool callGlEndQueryEXT(Stack* stack, bool pushReturn) {
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glEndQueryEXT(%u)\n", target);
        if (glEndQueryEXT != nullptr) {
            glEndQueryEXT(target);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glEndQueryEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEndQueryEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEndQueryEXT\n");
        return false;
    }
}

bool callGlEndTilingQCOM(Stack* stack, bool pushReturn) {
    GLbitfield preserve_mask = stack->pop<GLbitfield>();
    if (stack->isValid()) {
        GAPID_INFO("glEndTilingQCOM(%u)\n", preserve_mask);
        if (glEndTilingQCOM != nullptr) {
            glEndTilingQCOM(preserve_mask);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glEndTilingQCOM returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEndTilingQCOM\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEndTilingQCOM\n");
        return false;
    }
}

bool callGlExtGetBufferPointervQCOM(Stack* stack, bool pushReturn) {
    void** params = stack->pop<void**>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glExtGetBufferPointervQCOM(%u, %p)\n", target, params);
        if (glExtGetBufferPointervQCOM != nullptr) {
            glExtGetBufferPointervQCOM(target, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glExtGetBufferPointervQCOM returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glExtGetBufferPointervQCOM\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glExtGetBufferPointervQCOM\n");
        return false;
    }
}

bool callGlExtGetBuffersQCOM(Stack* stack, bool pushReturn) {
    int32_t* numBuffers = stack->pop<int32_t*>();
    int32_t maxBuffers = stack->pop<int32_t>();
    uint32_t* buffers = stack->pop<uint32_t*>();
    if (stack->isValid()) {
        GAPID_INFO("glExtGetBuffersQCOM(%p, %" PRId32 ", %p)\n", buffers, maxBuffers, numBuffers);
        if (glExtGetBuffersQCOM != nullptr) {
            glExtGetBuffersQCOM(buffers, maxBuffers, numBuffers);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glExtGetBuffersQCOM returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glExtGetBuffersQCOM\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glExtGetBuffersQCOM\n");
        return false;
    }
}

bool callGlExtGetFramebuffersQCOM(Stack* stack, bool pushReturn) {
    int32_t* numFramebuffers = stack->pop<int32_t*>();
    int32_t maxFramebuffers = stack->pop<int32_t>();
    uint32_t* framebuffers = stack->pop<uint32_t*>();
    if (stack->isValid()) {
        GAPID_INFO("glExtGetFramebuffersQCOM(%p, %" PRId32 ", %p)\n", framebuffers, maxFramebuffers,
                   numFramebuffers);
        if (glExtGetFramebuffersQCOM != nullptr) {
            glExtGetFramebuffersQCOM(framebuffers, maxFramebuffers, numFramebuffers);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glExtGetFramebuffersQCOM returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glExtGetFramebuffersQCOM\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glExtGetFramebuffersQCOM\n");
        return false;
    }
}

bool callGlExtGetProgramBinarySourceQCOM(Stack* stack, bool pushReturn) {
    int32_t* length = stack->pop<int32_t*>();
    char* source = stack->pop<char*>();
    GLenum shadertype = stack->pop<GLenum>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glExtGetProgramBinarySourceQCOM(%" PRIu32 ", %u, %p, %p)\n", program,
                   shadertype, source, length);
        if (glExtGetProgramBinarySourceQCOM != nullptr) {
            glExtGetProgramBinarySourceQCOM(program, shadertype, source, length);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glExtGetProgramBinarySourceQCOM returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glExtGetProgramBinarySourceQCOM\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glExtGetProgramBinarySourceQCOM\n");
        return false;
    }
}

bool callGlExtGetProgramsQCOM(Stack* stack, bool pushReturn) {
    int32_t* numPrograms = stack->pop<int32_t*>();
    int32_t maxPrograms = stack->pop<int32_t>();
    uint32_t* programs = stack->pop<uint32_t*>();
    if (stack->isValid()) {
        GAPID_INFO("glExtGetProgramsQCOM(%p, %" PRId32 ", %p)\n", programs, maxPrograms,
                   numPrograms);
        if (glExtGetProgramsQCOM != nullptr) {
            glExtGetProgramsQCOM(programs, maxPrograms, numPrograms);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glExtGetProgramsQCOM returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glExtGetProgramsQCOM\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glExtGetProgramsQCOM\n");
        return false;
    }
}

bool callGlExtGetRenderbuffersQCOM(Stack* stack, bool pushReturn) {
    int32_t* numRenderbuffers = stack->pop<int32_t*>();
    int32_t maxRenderbuffers = stack->pop<int32_t>();
    uint32_t* renderbuffers = stack->pop<uint32_t*>();
    if (stack->isValid()) {
        GAPID_INFO("glExtGetRenderbuffersQCOM(%p, %" PRId32 ", %p)\n", renderbuffers,
                   maxRenderbuffers, numRenderbuffers);
        if (glExtGetRenderbuffersQCOM != nullptr) {
            glExtGetRenderbuffersQCOM(renderbuffers, maxRenderbuffers, numRenderbuffers);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glExtGetRenderbuffersQCOM returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glExtGetRenderbuffersQCOM\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glExtGetRenderbuffersQCOM\n");
        return false;
    }
}

bool callGlExtGetShadersQCOM(Stack* stack, bool pushReturn) {
    int32_t* numShaders = stack->pop<int32_t*>();
    int32_t maxShaders = stack->pop<int32_t>();
    uint32_t* shaders = stack->pop<uint32_t*>();
    if (stack->isValid()) {
        GAPID_INFO("glExtGetShadersQCOM(%p, %" PRId32 ", %p)\n", shaders, maxShaders, numShaders);
        if (glExtGetShadersQCOM != nullptr) {
            glExtGetShadersQCOM(shaders, maxShaders, numShaders);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glExtGetShadersQCOM returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glExtGetShadersQCOM\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glExtGetShadersQCOM\n");
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
        GAPID_INFO("glExtGetTexLevelParameterivQCOM(%" PRIu32 ", %u, %" PRId32 ", %u, %p)\n",
                   texture, face, level, pname, params);
        if (glExtGetTexLevelParameterivQCOM != nullptr) {
            glExtGetTexLevelParameterivQCOM(texture, face, level, pname, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glExtGetTexLevelParameterivQCOM returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glExtGetTexLevelParameterivQCOM\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glExtGetTexLevelParameterivQCOM\n");
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
                   ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %u, %u, %p)\n",
                   target, level, xoffset, yoffset, zoffset, width, height, depth, format, type,
                   texels);
        if (glExtGetTexSubImageQCOM != nullptr) {
            glExtGetTexSubImageQCOM(target, level, xoffset, yoffset, zoffset, width, height, depth,
                                    format, type, texels);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glExtGetTexSubImageQCOM returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glExtGetTexSubImageQCOM\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glExtGetTexSubImageQCOM\n");
        return false;
    }
}

bool callGlExtGetTexturesQCOM(Stack* stack, bool pushReturn) {
    int32_t* numTextures = stack->pop<int32_t*>();
    int32_t maxTextures = stack->pop<int32_t>();
    uint32_t* textures = stack->pop<uint32_t*>();
    if (stack->isValid()) {
        GAPID_INFO("glExtGetTexturesQCOM(%p, %" PRId32 ", %p)\n", textures, maxTextures,
                   numTextures);
        if (glExtGetTexturesQCOM != nullptr) {
            glExtGetTexturesQCOM(textures, maxTextures, numTextures);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glExtGetTexturesQCOM returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glExtGetTexturesQCOM\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glExtGetTexturesQCOM\n");
        return false;
    }
}

bool callGlExtIsProgramBinaryQCOM(Stack* stack, bool pushReturn) {
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glExtIsProgramBinaryQCOM(%" PRIu32 ")\n", program);
        if (glExtIsProgramBinaryQCOM != nullptr) {
            uint8_t return_value = glExtIsProgramBinaryQCOM(program);
            GAPID_INFO("Returned: %" PRIu8 "\n", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glExtIsProgramBinaryQCOM returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glExtIsProgramBinaryQCOM\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glExtIsProgramBinaryQCOM\n");
        return false;
    }
}

bool callGlExtTexObjectStateOverrideiQCOM(Stack* stack, bool pushReturn) {
    int32_t param = stack->pop<int32_t>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glExtTexObjectStateOverrideiQCOM(%u, %u, %" PRId32 ")\n", target, pname, param);
        if (glExtTexObjectStateOverrideiQCOM != nullptr) {
            glExtTexObjectStateOverrideiQCOM(target, pname, param);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glExtTexObjectStateOverrideiQCOM returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glExtTexObjectStateOverrideiQCOM\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glExtTexObjectStateOverrideiQCOM\n");
        return false;
    }
}

bool callGlFenceSyncAPPLE(Stack* stack, bool pushReturn) {
    GLbitfield flag = stack->pop<GLbitfield>();
    GLenum condition = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glFenceSyncAPPLE(%u, %u)\n", condition, flag);
        if (glFenceSyncAPPLE != nullptr) {
            uint64_t return_value = glFenceSyncAPPLE(condition, flag);
            GAPID_INFO("Returned: %" PRIu64 "\n", return_value);
            if (pushReturn) {
                stack->push<uint64_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glFenceSyncAPPLE returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFenceSyncAPPLE\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFenceSyncAPPLE\n");
        return false;
    }
}

bool callGlFinishFenceNV(Stack* stack, bool pushReturn) {
    uint32_t fence = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glFinishFenceNV(%" PRIu32 ")\n", fence);
        if (glFinishFenceNV != nullptr) {
            glFinishFenceNV(fence);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glFinishFenceNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFinishFenceNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFinishFenceNV\n");
        return false;
    }
}

bool callGlFlushMappedBufferRangeEXT(Stack* stack, bool pushReturn) {
    int32_t length = stack->pop<int32_t>();
    int32_t offset = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glFlushMappedBufferRangeEXT(%u, %" PRId32 ", %" PRId32 ")\n", target, offset,
                   length);
        if (glFlushMappedBufferRangeEXT != nullptr) {
            glFlushMappedBufferRangeEXT(target, offset, length);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glFlushMappedBufferRangeEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFlushMappedBufferRangeEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFlushMappedBufferRangeEXT\n");
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
                   ", %" PRId32 ")\n",
                   target, attachment, textarget, texture, level, samples);
        if (glFramebufferTexture2DMultisampleEXT != nullptr) {
            glFramebufferTexture2DMultisampleEXT(target, attachment, textarget, texture, level,
                                                 samples);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glFramebufferTexture2DMultisampleEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function "
                    "glFramebufferTexture2DMultisampleEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFramebufferTexture2DMultisampleEXT\n");
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
                   ", %" PRId32 ")\n",
                   target, attachment, textarget, texture, level, samples);
        if (glFramebufferTexture2DMultisampleIMG != nullptr) {
            glFramebufferTexture2DMultisampleIMG(target, attachment, textarget, texture, level,
                                                 samples);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glFramebufferTexture2DMultisampleIMG returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function "
                    "glFramebufferTexture2DMultisampleIMG\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFramebufferTexture2DMultisampleIMG\n");
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
        GAPID_INFO("glFramebufferTexture3DOES(%u, %u, %u, %" PRIu32 ", %" PRId32 ", %" PRId32 ")\n",
                   target, attachment, textarget, texture, level, zoffset);
        if (glFramebufferTexture3DOES != nullptr) {
            glFramebufferTexture3DOES(target, attachment, textarget, texture, level, zoffset);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glFramebufferTexture3DOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFramebufferTexture3DOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFramebufferTexture3DOES\n");
        return false;
    }
}

bool callGlFramebufferTextureOES(Stack* stack, bool pushReturn) {
    int32_t level = stack->pop<int32_t>();
    uint32_t texture = stack->pop<uint32_t>();
    GLenum attachment = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glFramebufferTextureOES(%u, %u, %" PRIu32 ", %" PRId32 ")\n", target,
                   attachment, texture, level);
        if (glFramebufferTextureOES != nullptr) {
            glFramebufferTextureOES(target, attachment, texture, level);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glFramebufferTextureOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFramebufferTextureOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFramebufferTextureOES\n");
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
                   ", %" PRId32 ")\n",
                   target, attachment, texture, level, baseViewIndex, numViews);
        if (glFramebufferTextureMultiviewOVR != nullptr) {
            glFramebufferTextureMultiviewOVR(target, attachment, texture, level, baseViewIndex,
                                             numViews);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glFramebufferTextureMultiviewOVR returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glFramebufferTextureMultiviewOVR\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFramebufferTextureMultiviewOVR\n");
        return false;
    }
}

bool callGlGenFencesNV(Stack* stack, bool pushReturn) {
    uint32_t* fences = stack->pop<uint32_t*>();
    int32_t n = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenFencesNV(%" PRId32 ", %p)\n", n, fences);
        if (glGenFencesNV != nullptr) {
            glGenFencesNV(n, fences);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGenFencesNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenFencesNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenFencesNV\n");
        return false;
    }
}

bool callGlGenPathsNV(Stack* stack, bool pushReturn) {
    int32_t range = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenPathsNV(%" PRId32 ")\n", range);
        if (glGenPathsNV != nullptr) {
            uint32_t return_value = glGenPathsNV(range);
            GAPID_INFO("Returned: %" PRIu32 "\n", return_value);
            if (pushReturn) {
                stack->push<uint32_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGenPathsNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenPathsNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenPathsNV\n");
        return false;
    }
}

bool callGlGenPerfMonitorsAMD(Stack* stack, bool pushReturn) {
    uint32_t* monitors = stack->pop<uint32_t*>();
    int32_t n = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenPerfMonitorsAMD(%" PRId32 ", %p)\n", n, monitors);
        if (glGenPerfMonitorsAMD != nullptr) {
            glGenPerfMonitorsAMD(n, monitors);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGenPerfMonitorsAMD returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenPerfMonitorsAMD\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenPerfMonitorsAMD\n");
        return false;
    }
}

bool callGlGenProgramPipelinesEXT(Stack* stack, bool pushReturn) {
    uint32_t* pipelines = stack->pop<uint32_t*>();
    int32_t n = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenProgramPipelinesEXT(%" PRId32 ", %p)\n", n, pipelines);
        if (glGenProgramPipelinesEXT != nullptr) {
            glGenProgramPipelinesEXT(n, pipelines);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGenProgramPipelinesEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenProgramPipelinesEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenProgramPipelinesEXT\n");
        return false;
    }
}

bool callGlGenQueriesEXT(Stack* stack, bool pushReturn) {
    uint32_t* queries = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenQueriesEXT(%" PRId32 ", %p)\n", count, queries);
        if (glGenQueriesEXT != nullptr) {
            glGenQueriesEXT(count, queries);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGenQueriesEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenQueriesEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenQueriesEXT\n");
        return false;
    }
}

bool callGlGenVertexArraysOES(Stack* stack, bool pushReturn) {
    uint32_t* arrays = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenVertexArraysOES(%" PRId32 ", %p)\n", count, arrays);
        if (glGenVertexArraysOES != nullptr) {
            glGenVertexArraysOES(count, arrays);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGenVertexArraysOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenVertexArraysOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenVertexArraysOES\n");
        return false;
    }
}

bool callGlGetBufferPointervOES(Stack* stack, bool pushReturn) {
    void** params = stack->pop<void**>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetBufferPointervOES(%u, %u, %p)\n", target, pname, params);
        if (glGetBufferPointervOES != nullptr) {
            glGetBufferPointervOES(target, pname, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetBufferPointervOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetBufferPointervOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetBufferPointervOES\n");
        return false;
    }
}

bool callGlGetDriverControlStringQCOM(Stack* stack, bool pushReturn) {
    char* driverControlString = stack->pop<char*>();
    int32_t* length = stack->pop<int32_t*>();
    int32_t bufSize = stack->pop<int32_t>();
    uint32_t driverControl = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetDriverControlStringQCOM(%" PRIu32 ", %" PRId32 ", %p, %p)\n",
                   driverControl, bufSize, length, driverControlString);
        if (glGetDriverControlStringQCOM != nullptr) {
            glGetDriverControlStringQCOM(driverControl, bufSize, length, driverControlString);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetDriverControlStringQCOM returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetDriverControlStringQCOM\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetDriverControlStringQCOM\n");
        return false;
    }
}

bool callGlGetDriverControlsQCOM(Stack* stack, bool pushReturn) {
    uint32_t* driverControls = stack->pop<uint32_t*>();
    int32_t size = stack->pop<int32_t>();
    int32_t* num = stack->pop<int32_t*>();
    if (stack->isValid()) {
        GAPID_INFO("glGetDriverControlsQCOM(%p, %" PRId32 ", %p)\n", num, size, driverControls);
        if (glGetDriverControlsQCOM != nullptr) {
            glGetDriverControlsQCOM(num, size, driverControls);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetDriverControlsQCOM returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetDriverControlsQCOM\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetDriverControlsQCOM\n");
        return false;
    }
}

bool callGlGetFenceivNV(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t fence = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetFenceivNV(%" PRIu32 ", %u, %p)\n", fence, pname, params);
        if (glGetFenceivNV != nullptr) {
            glGetFenceivNV(fence, pname, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetFenceivNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetFenceivNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetFenceivNV\n");
        return false;
    }
}

bool callGlGetFirstPerfQueryIdINTEL(Stack* stack, bool pushReturn) {
    uint32_t* queryId = stack->pop<uint32_t*>();
    if (stack->isValid()) {
        GAPID_INFO("glGetFirstPerfQueryIdINTEL(%p)\n", queryId);
        if (glGetFirstPerfQueryIdINTEL != nullptr) {
            glGetFirstPerfQueryIdINTEL(queryId);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetFirstPerfQueryIdINTEL returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetFirstPerfQueryIdINTEL\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetFirstPerfQueryIdINTEL\n");
        return false;
    }
}

bool callGlGetFloatiVNV(Stack* stack, bool pushReturn) {
    float* data = stack->pop<float*>();
    uint32_t index = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetFloati_vNV(%u, %" PRIu32 ", %p)\n", target, index, data);
        if (glGetFloati_vNV != nullptr) {
            glGetFloati_vNV(target, index, data);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetFloati_vNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetFloati_vNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetFloati_vNV\n");
        return false;
    }
}

bool callGlGetGraphicsResetStatusEXT(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glGetGraphicsResetStatusEXT()\n");
        if (glGetGraphicsResetStatusEXT != nullptr) {
            GLenum return_value = glGetGraphicsResetStatusEXT();
            GAPID_INFO("Returned: %u\n", return_value);
            if (pushReturn) {
                stack->push<GLenum>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetGraphicsResetStatusEXT returned error: 0x%x\n", err);
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

bool callGlGetGraphicsResetStatusKHR(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glGetGraphicsResetStatusKHR()\n");
        if (glGetGraphicsResetStatusKHR != nullptr) {
            GLenum return_value = glGetGraphicsResetStatusKHR();
            GAPID_INFO("Returned: %u\n", return_value);
            if (pushReturn) {
                stack->push<GLenum>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetGraphicsResetStatusKHR returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetGraphicsResetStatusKHR\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetGraphicsResetStatusKHR\n");
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
        GAPID_INFO("glGetImageHandleNV(%" PRIu32 ", %" PRId32 ", %" PRIu8 ", %" PRId32 ", %u)\n",
                   texture, level, layered, layer, format);
        if (glGetImageHandleNV != nullptr) {
            uint64_t return_value = glGetImageHandleNV(texture, level, layered, layer, format);
            GAPID_INFO("Returned: %" PRIu64 "\n", return_value);
            if (pushReturn) {
                stack->push<uint64_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetImageHandleNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetImageHandleNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetImageHandleNV\n");
        return false;
    }
}

bool callGlGetInteger64vAPPLE(Stack* stack, bool pushReturn) {
    int64_t* params = stack->pop<int64_t*>();
    GLenum pname = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetInteger64vAPPLE(%u, %p)\n", pname, params);
        if (glGetInteger64vAPPLE != nullptr) {
            glGetInteger64vAPPLE(pname, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetInteger64vAPPLE returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetInteger64vAPPLE\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetInteger64vAPPLE\n");
        return false;
    }
}

bool callGlGetIntegeriVEXT(Stack* stack, bool pushReturn) {
    int32_t* data = stack->pop<int32_t*>();
    uint32_t index = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetIntegeri_vEXT(%u, %" PRIu32 ", %p)\n", target, index, data);
        if (glGetIntegeri_vEXT != nullptr) {
            glGetIntegeri_vEXT(target, index, data);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetIntegeri_vEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetIntegeri_vEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetIntegeri_vEXT\n");
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
        GAPID_INFO("glGetInternalformatSampleivNV(%u, %u, %" PRId32 ", %u, %" PRId32 ", %p)\n",
                   target, internalformat, samples, pname, bufSize, params);
        if (glGetInternalformatSampleivNV != nullptr) {
            glGetInternalformatSampleivNV(target, internalformat, samples, pname, bufSize, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetInternalformatSampleivNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetInternalformatSampleivNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetInternalformatSampleivNV\n");
        return false;
    }
}

bool callGlGetNextPerfQueryIdINTEL(Stack* stack, bool pushReturn) {
    uint32_t* nextQueryId = stack->pop<uint32_t*>();
    uint32_t queryId = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetNextPerfQueryIdINTEL(%" PRIu32 ", %p)\n", queryId, nextQueryId);
        if (glGetNextPerfQueryIdINTEL != nullptr) {
            glGetNextPerfQueryIdINTEL(queryId, nextQueryId);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetNextPerfQueryIdINTEL returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetNextPerfQueryIdINTEL\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetNextPerfQueryIdINTEL\n");
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
        GAPID_INFO("glGetObjectLabelEXT(%u, %" PRIu32 ", %" PRId32 ", %p, %p)\n", type, object,
                   bufSize, length, label);
        if (glGetObjectLabelEXT != nullptr) {
            glGetObjectLabelEXT(type, object, bufSize, length, label);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetObjectLabelEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetObjectLabelEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetObjectLabelEXT\n");
        return false;
    }
}

bool callGlGetPathCommandsNV(Stack* stack, bool pushReturn) {
    uint8_t* commands = stack->pop<uint8_t*>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetPathCommandsNV(%" PRIu32 ", %p)\n", path, commands);
        if (glGetPathCommandsNV != nullptr) {
            glGetPathCommandsNV(path, commands);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetPathCommandsNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPathCommandsNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPathCommandsNV\n");
        return false;
    }
}

bool callGlGetPathCoordsNV(Stack* stack, bool pushReturn) {
    float* coords = stack->pop<float*>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetPathCoordsNV(%" PRIu32 ", %p)\n", path, coords);
        if (glGetPathCoordsNV != nullptr) {
            glGetPathCoordsNV(path, coords);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetPathCoordsNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPathCoordsNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPathCoordsNV\n");
        return false;
    }
}

bool callGlGetPathDashArrayNV(Stack* stack, bool pushReturn) {
    float* dashArray = stack->pop<float*>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetPathDashArrayNV(%" PRIu32 ", %p)\n", path, dashArray);
        if (glGetPathDashArrayNV != nullptr) {
            glGetPathDashArrayNV(path, dashArray);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetPathDashArrayNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPathDashArrayNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPathDashArrayNV\n");
        return false;
    }
}

bool callGlGetPathLengthNV(Stack* stack, bool pushReturn) {
    int32_t numSegments = stack->pop<int32_t>();
    int32_t startSegment = stack->pop<int32_t>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetPathLengthNV(%" PRIu32 ", %" PRId32 ", %" PRId32 ")\n", path, startSegment,
                   numSegments);
        if (glGetPathLengthNV != nullptr) {
            float return_value = glGetPathLengthNV(path, startSegment, numSegments);
            GAPID_INFO("Returned: %f\n", return_value);
            if (pushReturn) {
                stack->push<float>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetPathLengthNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPathLengthNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPathLengthNV\n");
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
        GAPID_INFO("glGetPathMetricRangeNV(%u, %" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)\n",
                   metricQueryMask, firstPathName, numPaths, stride, metrics);
        if (glGetPathMetricRangeNV != nullptr) {
            glGetPathMetricRangeNV(metricQueryMask, firstPathName, numPaths, stride, metrics);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetPathMetricRangeNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPathMetricRangeNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPathMetricRangeNV\n");
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
        GAPID_INFO("glGetPathMetricsNV(%u, %" PRId32 ", %u, %p, %" PRIu32 ", %" PRId32 ", %p)\n",
                   metricQueryMask, numPaths, pathNameType, paths, pathBase, stride, metrics);
        if (glGetPathMetricsNV != nullptr) {
            glGetPathMetricsNV(metricQueryMask, numPaths, pathNameType, paths, pathBase, stride,
                               metrics);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetPathMetricsNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPathMetricsNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPathMetricsNV\n");
        return false;
    }
}

bool callGlGetPathParameterfvNV(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetPathParameterfvNV(%" PRIu32 ", %u, %p)\n", path, pname, value);
        if (glGetPathParameterfvNV != nullptr) {
            glGetPathParameterfvNV(path, pname, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetPathParameterfvNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPathParameterfvNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPathParameterfvNV\n");
        return false;
    }
}

bool callGlGetPathParameterivNV(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetPathParameterivNV(%" PRIu32 ", %u, %p)\n", path, pname, value);
        if (glGetPathParameterivNV != nullptr) {
            glGetPathParameterivNV(path, pname, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetPathParameterivNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPathParameterivNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPathParameterivNV\n");
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
        GAPID_INFO("glGetPathSpacingNV(%u, %" PRId32 ", %u, %p, %" PRIu32 ", %f, %f, %u, %p)\n",
                   pathListMode, numPaths, pathNameType, paths, pathBase, advanceScale,
                   kerningScale, transformType, returnedSpacing);
        if (glGetPathSpacingNV != nullptr) {
            glGetPathSpacingNV(pathListMode, numPaths, pathNameType, paths, pathBase, advanceScale,
                               kerningScale, transformType, returnedSpacing);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetPathSpacingNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPathSpacingNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPathSpacingNV\n");
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
                   ", %p, %p, %p, %p, %p, %p)\n",
                   queryId, counterId, counterNameLength, counterName, counterDescLength,
                   counterDesc, counterOffset, counterDataSize, counterTypeEnum,
                   counterDataTypeEnum, rawCounterMaxValue);
        if (glGetPerfCounterInfoINTEL != nullptr) {
            glGetPerfCounterInfoINTEL(queryId, counterId, counterNameLength, counterName,
                                      counterDescLength, counterDesc, counterOffset,
                                      counterDataSize, counterTypeEnum, counterDataTypeEnum,
                                      rawCounterMaxValue);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetPerfCounterInfoINTEL returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPerfCounterInfoINTEL\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPerfCounterInfoINTEL\n");
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
        GAPID_INFO("glGetPerfMonitorCounterDataAMD(%" PRIu32 ", %u, %" PRId32 ", %p, %p)\n",
                   monitor, pname, dataSize, data, bytesWritten);
        if (glGetPerfMonitorCounterDataAMD != nullptr) {
            glGetPerfMonitorCounterDataAMD(monitor, pname, dataSize, data, bytesWritten);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetPerfMonitorCounterDataAMD returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glGetPerfMonitorCounterDataAMD\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPerfMonitorCounterDataAMD\n");
        return false;
    }
}

bool callGlGetPerfMonitorCounterInfoAMD(Stack* stack, bool pushReturn) {
    void* data = stack->pop<void*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t counter = stack->pop<uint32_t>();
    uint32_t group = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetPerfMonitorCounterInfoAMD(%" PRIu32 ", %" PRIu32 ", %u, %p)\n", group,
                   counter, pname, data);
        if (glGetPerfMonitorCounterInfoAMD != nullptr) {
            glGetPerfMonitorCounterInfoAMD(group, counter, pname, data);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetPerfMonitorCounterInfoAMD returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glGetPerfMonitorCounterInfoAMD\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPerfMonitorCounterInfoAMD\n");
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
                   ", %p, %p)\n",
                   group, counter, bufSize, length, counterString);
        if (glGetPerfMonitorCounterStringAMD != nullptr) {
            glGetPerfMonitorCounterStringAMD(group, counter, bufSize, length, counterString);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetPerfMonitorCounterStringAMD returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glGetPerfMonitorCounterStringAMD\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPerfMonitorCounterStringAMD\n");
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
        GAPID_INFO("glGetPerfMonitorCountersAMD(%" PRIu32 ", %p, %p, %" PRId32 ", %p)\n", group,
                   numCounters, maxActiveCounters, counterSize, counters);
        if (glGetPerfMonitorCountersAMD != nullptr) {
            glGetPerfMonitorCountersAMD(group, numCounters, maxActiveCounters, counterSize,
                                        counters);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetPerfMonitorCountersAMD returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPerfMonitorCountersAMD\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPerfMonitorCountersAMD\n");
        return false;
    }
}

bool callGlGetPerfMonitorGroupStringAMD(Stack* stack, bool pushReturn) {
    char* groupString = stack->pop<char*>();
    int32_t* length = stack->pop<int32_t*>();
    int32_t bufSize = stack->pop<int32_t>();
    uint32_t group = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetPerfMonitorGroupStringAMD(%" PRIu32 ", %" PRId32 ", %p, %p)\n", group,
                   bufSize, length, groupString);
        if (glGetPerfMonitorGroupStringAMD != nullptr) {
            glGetPerfMonitorGroupStringAMD(group, bufSize, length, groupString);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetPerfMonitorGroupStringAMD returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glGetPerfMonitorGroupStringAMD\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPerfMonitorGroupStringAMD\n");
        return false;
    }
}

bool callGlGetPerfMonitorGroupsAMD(Stack* stack, bool pushReturn) {
    uint32_t* groups = stack->pop<uint32_t*>();
    int32_t groupsSize = stack->pop<int32_t>();
    int32_t* numGroups = stack->pop<int32_t*>();
    if (stack->isValid()) {
        GAPID_INFO("glGetPerfMonitorGroupsAMD(%p, %" PRId32 ", %p)\n", numGroups, groupsSize,
                   groups);
        if (glGetPerfMonitorGroupsAMD != nullptr) {
            glGetPerfMonitorGroupsAMD(numGroups, groupsSize, groups);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetPerfMonitorGroupsAMD returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPerfMonitorGroupsAMD\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPerfMonitorGroupsAMD\n");
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
        GAPID_INFO("glGetPerfQueryDataINTEL(%" PRIu32 ", %" PRIu32 ", %" PRId32 ", %p, %p)\n",
                   queryHandle, flag, dataSize, data, bytesWritten);
        if (glGetPerfQueryDataINTEL != nullptr) {
            glGetPerfQueryDataINTEL(queryHandle, flag, dataSize, data, bytesWritten);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetPerfQueryDataINTEL returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPerfQueryDataINTEL\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPerfQueryDataINTEL\n");
        return false;
    }
}

bool callGlGetPerfQueryIdByNameINTEL(Stack* stack, bool pushReturn) {
    uint32_t* queryId = stack->pop<uint32_t*>();
    char* queryName = stack->pop<char*>();
    if (stack->isValid()) {
        GAPID_INFO("glGetPerfQueryIdByNameINTEL(%p, %p)\n", queryName, queryId);
        if (glGetPerfQueryIdByNameINTEL != nullptr) {
            glGetPerfQueryIdByNameINTEL(queryName, queryId);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetPerfQueryIdByNameINTEL returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPerfQueryIdByNameINTEL\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPerfQueryIdByNameINTEL\n");
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
        GAPID_INFO("glGetPerfQueryInfoINTEL(%" PRIu32 ", %" PRIu32 ", %p, %p, %p, %p, %p)\n",
                   queryId, queryNameLength, queryName, dataSize, noCounters, noInstances,
                   capsMask);
        if (glGetPerfQueryInfoINTEL != nullptr) {
            glGetPerfQueryInfoINTEL(queryId, queryNameLength, queryName, dataSize, noCounters,
                                    noInstances, capsMask);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetPerfQueryInfoINTEL returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetPerfQueryInfoINTEL\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetPerfQueryInfoINTEL\n");
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
        GAPID_INFO("glGetProgramBinaryOES(%" PRIu32 ", %" PRId32 ", %p, %p, %p)\n", program,
                   buffer_size, bytes_written, binary_format, binary);
        if (glGetProgramBinaryOES != nullptr) {
            glGetProgramBinaryOES(program, buffer_size, bytes_written, binary_format, binary);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetProgramBinaryOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetProgramBinaryOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetProgramBinaryOES\n");
        return false;
    }
}

bool callGlGetProgramPipelineInfoLogEXT(Stack* stack, bool pushReturn) {
    char* infoLog = stack->pop<char*>();
    int32_t* length = stack->pop<int32_t*>();
    int32_t bufSize = stack->pop<int32_t>();
    uint32_t pipeline = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetProgramPipelineInfoLogEXT(%" PRIu32 ", %" PRId32 ", %p, %p)\n", pipeline,
                   bufSize, length, infoLog);
        if (glGetProgramPipelineInfoLogEXT != nullptr) {
            glGetProgramPipelineInfoLogEXT(pipeline, bufSize, length, infoLog);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetProgramPipelineInfoLogEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glGetProgramPipelineInfoLogEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetProgramPipelineInfoLogEXT\n");
        return false;
    }
}

bool callGlGetProgramPipelineivEXT(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t pipeline = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetProgramPipelineivEXT(%" PRIu32 ", %u, %p)\n", pipeline, pname, params);
        if (glGetProgramPipelineivEXT != nullptr) {
            glGetProgramPipelineivEXT(pipeline, pname, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetProgramPipelineivEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetProgramPipelineivEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetProgramPipelineivEXT\n");
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
                   ", %p, %" PRId32 ", %p, %p)\n",
                   program, programInterface, index, propCount, props, bufSize, length, params);
        if (glGetProgramResourcefvNV != nullptr) {
            glGetProgramResourcefvNV(program, programInterface, index, propCount, props, bufSize,
                                     length, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetProgramResourcefvNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetProgramResourcefvNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetProgramResourcefvNV\n");
        return false;
    }
}

bool callGlGetQueryObjecti64vEXT(Stack* stack, bool pushReturn) {
    int64_t* value = stack->pop<int64_t*>();
    GLenum parameter = stack->pop<GLenum>();
    uint32_t query = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetQueryObjecti64vEXT(%" PRIu32 ", %u, %p)\n", query, parameter, value);
        if (glGetQueryObjecti64vEXT != nullptr) {
            glGetQueryObjecti64vEXT(query, parameter, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetQueryObjecti64vEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetQueryObjecti64vEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetQueryObjecti64vEXT\n");
        return false;
    }
}

bool callGlGetQueryObjectivEXT(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    GLenum parameter = stack->pop<GLenum>();
    uint32_t query = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetQueryObjectivEXT(%" PRIu32 ", %u, %p)\n", query, parameter, value);
        if (glGetQueryObjectivEXT != nullptr) {
            glGetQueryObjectivEXT(query, parameter, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetQueryObjectivEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetQueryObjectivEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetQueryObjectivEXT\n");
        return false;
    }
}

bool callGlGetQueryObjectui64vEXT(Stack* stack, bool pushReturn) {
    uint64_t* value = stack->pop<uint64_t*>();
    GLenum parameter = stack->pop<GLenum>();
    uint32_t query = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetQueryObjectui64vEXT(%" PRIu32 ", %u, %p)\n", query, parameter, value);
        if (glGetQueryObjectui64vEXT != nullptr) {
            glGetQueryObjectui64vEXT(query, parameter, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetQueryObjectui64vEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetQueryObjectui64vEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetQueryObjectui64vEXT\n");
        return false;
    }
}

bool callGlGetQueryObjectuivEXT(Stack* stack, bool pushReturn) {
    uint32_t* value = stack->pop<uint32_t*>();
    GLenum parameter = stack->pop<GLenum>();
    uint32_t query = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetQueryObjectuivEXT(%" PRIu32 ", %u, %p)\n", query, parameter, value);
        if (glGetQueryObjectuivEXT != nullptr) {
            glGetQueryObjectuivEXT(query, parameter, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetQueryObjectuivEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetQueryObjectuivEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetQueryObjectuivEXT\n");
        return false;
    }
}

bool callGlGetQueryivEXT(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    GLenum parameter = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetQueryivEXT(%u, %u, %p)\n", target, parameter, value);
        if (glGetQueryivEXT != nullptr) {
            glGetQueryivEXT(target, parameter, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetQueryivEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetQueryivEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetQueryivEXT\n");
        return false;
    }
}

bool callGlGetSamplerParameterIivOES(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetSamplerParameterIivOES(%" PRIu32 ", %u, %p)\n", sampler, pname, params);
        if (glGetSamplerParameterIivOES != nullptr) {
            glGetSamplerParameterIivOES(sampler, pname, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetSamplerParameterIivOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetSamplerParameterIivOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetSamplerParameterIivOES\n");
        return false;
    }
}

bool callGlGetSamplerParameterIuivOES(Stack* stack, bool pushReturn) {
    uint32_t* params = stack->pop<uint32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetSamplerParameterIuivOES(%" PRIu32 ", %u, %p)\n", sampler, pname, params);
        if (glGetSamplerParameterIuivOES != nullptr) {
            glGetSamplerParameterIuivOES(sampler, pname, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetSamplerParameterIuivOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetSamplerParameterIuivOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetSamplerParameterIuivOES\n");
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
        GAPID_INFO("glGetSyncivAPPLE(%" PRIu64 ", %u, %" PRId32 ", %p, %p)\n", sync, pname, bufSize,
                   length, values);
        if (glGetSyncivAPPLE != nullptr) {
            glGetSyncivAPPLE(sync, pname, bufSize, length, values);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetSyncivAPPLE returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetSyncivAPPLE\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetSyncivAPPLE\n");
        return false;
    }
}

bool callGlGetTexParameterIivOES(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetTexParameterIivOES(%u, %u, %p)\n", target, pname, params);
        if (glGetTexParameterIivOES != nullptr) {
            glGetTexParameterIivOES(target, pname, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetTexParameterIivOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetTexParameterIivOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetTexParameterIivOES\n");
        return false;
    }
}

bool callGlGetTexParameterIuivOES(Stack* stack, bool pushReturn) {
    uint32_t* params = stack->pop<uint32_t*>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetTexParameterIuivOES(%u, %u, %p)\n", target, pname, params);
        if (glGetTexParameterIuivOES != nullptr) {
            glGetTexParameterIuivOES(target, pname, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetTexParameterIuivOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetTexParameterIuivOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetTexParameterIuivOES\n");
        return false;
    }
}

bool callGlGetTextureHandleNV(Stack* stack, bool pushReturn) {
    uint32_t texture = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetTextureHandleNV(%" PRIu32 ")\n", texture);
        if (glGetTextureHandleNV != nullptr) {
            uint64_t return_value = glGetTextureHandleNV(texture);
            GAPID_INFO("Returned: %" PRIu64 "\n", return_value);
            if (pushReturn) {
                stack->push<uint64_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetTextureHandleNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetTextureHandleNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetTextureHandleNV\n");
        return false;
    }
}

bool callGlGetTextureSamplerHandleNV(Stack* stack, bool pushReturn) {
    uint32_t sampler = stack->pop<uint32_t>();
    uint32_t texture = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetTextureSamplerHandleNV(%" PRIu32 ", %" PRIu32 ")\n", texture, sampler);
        if (glGetTextureSamplerHandleNV != nullptr) {
            uint64_t return_value = glGetTextureSamplerHandleNV(texture, sampler);
            GAPID_INFO("Returned: %" PRIu64 "\n", return_value);
            if (pushReturn) {
                stack->push<uint64_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetTextureSamplerHandleNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetTextureSamplerHandleNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetTextureSamplerHandleNV\n");
        return false;
    }
}

bool callGlGetTranslatedShaderSourceANGLE(Stack* stack, bool pushReturn) {
    char* source = stack->pop<char*>();
    int32_t* length = stack->pop<int32_t*>();
    int32_t bufsize = stack->pop<int32_t>();
    uint32_t shader = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetTranslatedShaderSourceANGLE(%" PRIu32 ", %" PRId32 ", %p, %p)\n", shader,
                   bufsize, length, source);
        if (glGetTranslatedShaderSourceANGLE != nullptr) {
            glGetTranslatedShaderSourceANGLE(shader, bufsize, length, source);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetTranslatedShaderSourceANGLE returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glGetTranslatedShaderSourceANGLE\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetTranslatedShaderSourceANGLE\n");
        return false;
    }
}

bool callGlGetnUniformfvEXT(Stack* stack, bool pushReturn) {
    float* params = stack->pop<float*>();
    int32_t bufSize = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetnUniformfvEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)\n", program,
                   location, bufSize, params);
        if (glGetnUniformfvEXT != nullptr) {
            glGetnUniformfvEXT(program, location, bufSize, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetnUniformfvEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetnUniformfvEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetnUniformfvEXT\n");
        return false;
    }
}

bool callGlGetnUniformfvKHR(Stack* stack, bool pushReturn) {
    float* params = stack->pop<float*>();
    int32_t bufSize = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetnUniformfvKHR(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)\n", program,
                   location, bufSize, params);
        if (glGetnUniformfvKHR != nullptr) {
            glGetnUniformfvKHR(program, location, bufSize, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetnUniformfvKHR returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetnUniformfvKHR\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetnUniformfvKHR\n");
        return false;
    }
}

bool callGlGetnUniformivEXT(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    int32_t bufSize = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetnUniformivEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)\n", program,
                   location, bufSize, params);
        if (glGetnUniformivEXT != nullptr) {
            glGetnUniformivEXT(program, location, bufSize, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetnUniformivEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetnUniformivEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetnUniformivEXT\n");
        return false;
    }
}

bool callGlGetnUniformivKHR(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    int32_t bufSize = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetnUniformivKHR(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)\n", program,
                   location, bufSize, params);
        if (glGetnUniformivKHR != nullptr) {
            glGetnUniformivKHR(program, location, bufSize, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetnUniformivKHR returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetnUniformivKHR\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetnUniformivKHR\n");
        return false;
    }
}

bool callGlGetnUniformuivKHR(Stack* stack, bool pushReturn) {
    uint32_t* params = stack->pop<uint32_t*>();
    int32_t bufSize = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetnUniformuivKHR(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)\n", program,
                   location, bufSize, params);
        if (glGetnUniformuivKHR != nullptr) {
            glGetnUniformuivKHR(program, location, bufSize, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetnUniformuivKHR returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetnUniformuivKHR\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetnUniformuivKHR\n");
        return false;
    }
}

bool callGlInsertEventMarkerEXT(Stack* stack, bool pushReturn) {
    char* marker = stack->pop<char*>();
    int32_t length = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glInsertEventMarkerEXT(%" PRId32 ", %p)\n", length, marker);
        if (glInsertEventMarkerEXT != nullptr) {
            glInsertEventMarkerEXT(length, marker);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glInsertEventMarkerEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glInsertEventMarkerEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glInsertEventMarkerEXT\n");
        return false;
    }
}

bool callGlInterpolatePathsNV(Stack* stack, bool pushReturn) {
    float weight = stack->pop<float>();
    uint32_t pathB = stack->pop<uint32_t>();
    uint32_t pathA = stack->pop<uint32_t>();
    uint32_t resultPath = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glInterpolatePathsNV(%" PRIu32 ", %" PRIu32 ", %" PRIu32 ", %f)\n", resultPath,
                   pathA, pathB, weight);
        if (glInterpolatePathsNV != nullptr) {
            glInterpolatePathsNV(resultPath, pathA, pathB, weight);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glInterpolatePathsNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glInterpolatePathsNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glInterpolatePathsNV\n");
        return false;
    }
}

bool callGlIsEnablediOES(Stack* stack, bool pushReturn) {
    uint32_t index = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glIsEnablediOES(%u, %" PRIu32 ")\n", target, index);
        if (glIsEnablediOES != nullptr) {
            uint8_t return_value = glIsEnablediOES(target, index);
            GAPID_INFO("Returned: %" PRIu8 "\n", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glIsEnablediOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsEnablediOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsEnablediOES\n");
        return false;
    }
}

bool callGlIsEnablediNV(Stack* stack, bool pushReturn) {
    uint32_t index = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glIsEnablediNV(%u, %" PRIu32 ")\n", target, index);
        if (glIsEnablediNV != nullptr) {
            uint8_t return_value = glIsEnablediNV(target, index);
            GAPID_INFO("Returned: %" PRIu8 "\n", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glIsEnablediNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsEnablediNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsEnablediNV\n");
        return false;
    }
}

bool callGlIsFenceNV(Stack* stack, bool pushReturn) {
    uint32_t fence = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsFenceNV(%" PRIu32 ")\n", fence);
        if (glIsFenceNV != nullptr) {
            uint8_t return_value = glIsFenceNV(fence);
            GAPID_INFO("Returned: %" PRIu8 "\n", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glIsFenceNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsFenceNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsFenceNV\n");
        return false;
    }
}

bool callGlIsImageHandleResidentNV(Stack* stack, bool pushReturn) {
    uint64_t handle = stack->pop<uint64_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsImageHandleResidentNV(%" PRIu64 ")\n", handle);
        if (glIsImageHandleResidentNV != nullptr) {
            uint8_t return_value = glIsImageHandleResidentNV(handle);
            GAPID_INFO("Returned: %" PRIu8 "\n", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glIsImageHandleResidentNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsImageHandleResidentNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsImageHandleResidentNV\n");
        return false;
    }
}

bool callGlIsPathNV(Stack* stack, bool pushReturn) {
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsPathNV(%" PRIu32 ")\n", path);
        if (glIsPathNV != nullptr) {
            uint8_t return_value = glIsPathNV(path);
            GAPID_INFO("Returned: %" PRIu8 "\n", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glIsPathNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsPathNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsPathNV\n");
        return false;
    }
}

bool callGlIsPointInFillPathNV(Stack* stack, bool pushReturn) {
    float y = stack->pop<float>();
    float x = stack->pop<float>();
    uint32_t mask = stack->pop<uint32_t>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsPointInFillPathNV(%" PRIu32 ", %" PRIu32 ", %f, %f)\n", path, mask, x, y);
        if (glIsPointInFillPathNV != nullptr) {
            uint8_t return_value = glIsPointInFillPathNV(path, mask, x, y);
            GAPID_INFO("Returned: %" PRIu8 "\n", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glIsPointInFillPathNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsPointInFillPathNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsPointInFillPathNV\n");
        return false;
    }
}

bool callGlIsPointInStrokePathNV(Stack* stack, bool pushReturn) {
    float y = stack->pop<float>();
    float x = stack->pop<float>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsPointInStrokePathNV(%" PRIu32 ", %f, %f)\n", path, x, y);
        if (glIsPointInStrokePathNV != nullptr) {
            uint8_t return_value = glIsPointInStrokePathNV(path, x, y);
            GAPID_INFO("Returned: %" PRIu8 "\n", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glIsPointInStrokePathNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsPointInStrokePathNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsPointInStrokePathNV\n");
        return false;
    }
}

bool callGlIsProgramPipelineEXT(Stack* stack, bool pushReturn) {
    uint32_t pipeline = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsProgramPipelineEXT(%" PRIu32 ")\n", pipeline);
        if (glIsProgramPipelineEXT != nullptr) {
            uint8_t return_value = glIsProgramPipelineEXT(pipeline);
            GAPID_INFO("Returned: %" PRIu8 "\n", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glIsProgramPipelineEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsProgramPipelineEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsProgramPipelineEXT\n");
        return false;
    }
}

bool callGlIsQueryEXT(Stack* stack, bool pushReturn) {
    uint32_t query = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsQueryEXT(%" PRIu32 ")\n", query);
        if (glIsQueryEXT != nullptr) {
            uint8_t return_value = glIsQueryEXT(query);
            GAPID_INFO("Returned: %" PRIu8 "\n", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glIsQueryEXT returned error: 0x%x\n", err);
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

bool callGlIsSyncAPPLE(Stack* stack, bool pushReturn) {
    uint64_t sync = stack->pop<uint64_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsSyncAPPLE(%" PRIu64 ")\n", sync);
        if (glIsSyncAPPLE != nullptr) {
            uint8_t return_value = glIsSyncAPPLE(sync);
            GAPID_INFO("Returned: %" PRIu8 "\n", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glIsSyncAPPLE returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsSyncAPPLE\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsSyncAPPLE\n");
        return false;
    }
}

bool callGlIsTextureHandleResidentNV(Stack* stack, bool pushReturn) {
    uint64_t handle = stack->pop<uint64_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsTextureHandleResidentNV(%" PRIu64 ")\n", handle);
        if (glIsTextureHandleResidentNV != nullptr) {
            uint8_t return_value = glIsTextureHandleResidentNV(handle);
            GAPID_INFO("Returned: %" PRIu8 "\n", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glIsTextureHandleResidentNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsTextureHandleResidentNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsTextureHandleResidentNV\n");
        return false;
    }
}

bool callGlIsVertexArrayOES(Stack* stack, bool pushReturn) {
    uint32_t array = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsVertexArrayOES(%" PRIu32 ")\n", array);
        if (glIsVertexArrayOES != nullptr) {
            uint8_t return_value = glIsVertexArrayOES(array);
            GAPID_INFO("Returned: %" PRIu8 "\n", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glIsVertexArrayOES returned error: 0x%x\n", err);
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

bool callGlLabelObjectEXT(Stack* stack, bool pushReturn) {
    char* label = stack->pop<char*>();
    int32_t length = stack->pop<int32_t>();
    uint32_t object = stack->pop<uint32_t>();
    GLenum type = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glLabelObjectEXT(%u, %" PRIu32 ", %" PRId32 ", %p)\n", type, object, length,
                   label);
        if (glLabelObjectEXT != nullptr) {
            glLabelObjectEXT(type, object, length, label);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glLabelObjectEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glLabelObjectEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glLabelObjectEXT\n");
        return false;
    }
}

bool callGlMakeImageHandleNonResidentNV(Stack* stack, bool pushReturn) {
    uint64_t handle = stack->pop<uint64_t>();
    if (stack->isValid()) {
        GAPID_INFO("glMakeImageHandleNonResidentNV(%" PRIu64 ")\n", handle);
        if (glMakeImageHandleNonResidentNV != nullptr) {
            glMakeImageHandleNonResidentNV(handle);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glMakeImageHandleNonResidentNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glMakeImageHandleNonResidentNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMakeImageHandleNonResidentNV\n");
        return false;
    }
}

bool callGlMakeImageHandleResidentNV(Stack* stack, bool pushReturn) {
    GLenum access = stack->pop<GLenum>();
    uint64_t handle = stack->pop<uint64_t>();
    if (stack->isValid()) {
        GAPID_INFO("glMakeImageHandleResidentNV(%" PRIu64 ", %u)\n", handle, access);
        if (glMakeImageHandleResidentNV != nullptr) {
            glMakeImageHandleResidentNV(handle, access);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glMakeImageHandleResidentNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMakeImageHandleResidentNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMakeImageHandleResidentNV\n");
        return false;
    }
}

bool callGlMakeTextureHandleNonResidentNV(Stack* stack, bool pushReturn) {
    uint64_t handle = stack->pop<uint64_t>();
    if (stack->isValid()) {
        GAPID_INFO("glMakeTextureHandleNonResidentNV(%" PRIu64 ")\n", handle);
        if (glMakeTextureHandleNonResidentNV != nullptr) {
            glMakeTextureHandleNonResidentNV(handle);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glMakeTextureHandleNonResidentNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glMakeTextureHandleNonResidentNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMakeTextureHandleNonResidentNV\n");
        return false;
    }
}

bool callGlMakeTextureHandleResidentNV(Stack* stack, bool pushReturn) {
    uint64_t handle = stack->pop<uint64_t>();
    if (stack->isValid()) {
        GAPID_INFO("glMakeTextureHandleResidentNV(%" PRIu64 ")\n", handle);
        if (glMakeTextureHandleResidentNV != nullptr) {
            glMakeTextureHandleResidentNV(handle);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glMakeTextureHandleResidentNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMakeTextureHandleResidentNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMakeTextureHandleResidentNV\n");
        return false;
    }
}

bool callGlMapBufferOES(Stack* stack, bool pushReturn) {
    GLenum access = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glMapBufferOES(%u, %u)\n", target, access);
        if (glMapBufferOES != nullptr) {
            void* return_value = glMapBufferOES(target, access);
            GAPID_INFO("Returned: %p\n", return_value);
            if (pushReturn) {
                stack->push<void*>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glMapBufferOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMapBufferOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMapBufferOES\n");
        return false;
    }
}

bool callGlMapBufferRangeEXT(Stack* stack, bool pushReturn) {
    GLbitfield access = stack->pop<GLbitfield>();
    int32_t length = stack->pop<int32_t>();
    int32_t offset = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glMapBufferRangeEXT(%u, %" PRId32 ", %" PRId32 ", %u)\n", target, offset,
                   length, access);
        if (glMapBufferRangeEXT != nullptr) {
            void* return_value = glMapBufferRangeEXT(target, offset, length, access);
            GAPID_INFO("Returned: %p\n", return_value);
            if (pushReturn) {
                stack->push<void*>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glMapBufferRangeEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMapBufferRangeEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMapBufferRangeEXT\n");
        return false;
    }
}

bool callGlMatrixLoad3x2fNV(Stack* stack, bool pushReturn) {
    float* m = stack->pop<float*>();
    GLenum matrixMode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glMatrixLoad3x2fNV(%u, %p)\n", matrixMode, m);
        if (glMatrixLoad3x2fNV != nullptr) {
            glMatrixLoad3x2fNV(matrixMode, m);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glMatrixLoad3x2fNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMatrixLoad3x2fNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMatrixLoad3x2fNV\n");
        return false;
    }
}

bool callGlMatrixLoad3x3fNV(Stack* stack, bool pushReturn) {
    float* m = stack->pop<float*>();
    GLenum matrixMode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glMatrixLoad3x3fNV(%u, %p)\n", matrixMode, m);
        if (glMatrixLoad3x3fNV != nullptr) {
            glMatrixLoad3x3fNV(matrixMode, m);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glMatrixLoad3x3fNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMatrixLoad3x3fNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMatrixLoad3x3fNV\n");
        return false;
    }
}

bool callGlMatrixLoadTranspose3x3fNV(Stack* stack, bool pushReturn) {
    float* m = stack->pop<float*>();
    GLenum matrixMode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glMatrixLoadTranspose3x3fNV(%u, %p)\n", matrixMode, m);
        if (glMatrixLoadTranspose3x3fNV != nullptr) {
            glMatrixLoadTranspose3x3fNV(matrixMode, m);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glMatrixLoadTranspose3x3fNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMatrixLoadTranspose3x3fNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMatrixLoadTranspose3x3fNV\n");
        return false;
    }
}

bool callGlMatrixMult3x2fNV(Stack* stack, bool pushReturn) {
    float* m = stack->pop<float*>();
    GLenum matrixMode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glMatrixMult3x2fNV(%u, %p)\n", matrixMode, m);
        if (glMatrixMult3x2fNV != nullptr) {
            glMatrixMult3x2fNV(matrixMode, m);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glMatrixMult3x2fNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMatrixMult3x2fNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMatrixMult3x2fNV\n");
        return false;
    }
}

bool callGlMatrixMult3x3fNV(Stack* stack, bool pushReturn) {
    float* m = stack->pop<float*>();
    GLenum matrixMode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glMatrixMult3x3fNV(%u, %p)\n", matrixMode, m);
        if (glMatrixMult3x3fNV != nullptr) {
            glMatrixMult3x3fNV(matrixMode, m);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glMatrixMult3x3fNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMatrixMult3x3fNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMatrixMult3x3fNV\n");
        return false;
    }
}

bool callGlMatrixMultTranspose3x3fNV(Stack* stack, bool pushReturn) {
    float* m = stack->pop<float*>();
    GLenum matrixMode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glMatrixMultTranspose3x3fNV(%u, %p)\n", matrixMode, m);
        if (glMatrixMultTranspose3x3fNV != nullptr) {
            glMatrixMultTranspose3x3fNV(matrixMode, m);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glMatrixMultTranspose3x3fNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMatrixMultTranspose3x3fNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMatrixMultTranspose3x3fNV\n");
        return false;
    }
}

bool callGlMultiDrawArraysEXT(Stack* stack, bool pushReturn) {
    int32_t primcount = stack->pop<int32_t>();
    int32_t* count = stack->pop<int32_t*>();
    int32_t* first = stack->pop<int32_t*>();
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glMultiDrawArraysEXT(%u, %p, %p, %" PRId32 ")\n", mode, first, count,
                   primcount);
        if (glMultiDrawArraysEXT != nullptr) {
            glMultiDrawArraysEXT(mode, first, count, primcount);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glMultiDrawArraysEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMultiDrawArraysEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMultiDrawArraysEXT\n");
        return false;
    }
}

bool callGlMultiDrawArraysIndirectEXT(Stack* stack, bool pushReturn) {
    int32_t stride = stack->pop<int32_t>();
    int32_t drawcount = stack->pop<int32_t>();
    void* indirect = stack->pop<void*>();
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glMultiDrawArraysIndirectEXT(%u, %p, %" PRId32 ", %" PRId32 ")\n", mode,
                   indirect, drawcount, stride);
        if (glMultiDrawArraysIndirectEXT != nullptr) {
            glMultiDrawArraysIndirectEXT(mode, indirect, drawcount, stride);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glMultiDrawArraysIndirectEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMultiDrawArraysIndirectEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMultiDrawArraysIndirectEXT\n");
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
        GAPID_INFO("glMultiDrawElementsBaseVertexEXT(%u, %p, %u, %p, %" PRId32 ", %p)\n", mode,
                   count, type, indices, primcount, basevertex);
        if (glMultiDrawElementsBaseVertexEXT != nullptr) {
            glMultiDrawElementsBaseVertexEXT(mode, count, type, indices, primcount, basevertex);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glMultiDrawElementsBaseVertexEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glMultiDrawElementsBaseVertexEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMultiDrawElementsBaseVertexEXT\n");
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
        GAPID_INFO("glMultiDrawElementsBaseVertexOES(%u, %p, %u, %p, %" PRId32 ", %p)\n", mode,
                   count, type, indices, primcount, basevertex);
        if (glMultiDrawElementsBaseVertexOES != nullptr) {
            glMultiDrawElementsBaseVertexOES(mode, count, type, indices, primcount, basevertex);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glMultiDrawElementsBaseVertexOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glMultiDrawElementsBaseVertexOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMultiDrawElementsBaseVertexOES\n");
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
        GAPID_INFO("glMultiDrawElementsEXT(%u, %p, %u, %p, %" PRId32 ")\n", mode, count, type,
                   indices, primcount);
        if (glMultiDrawElementsEXT != nullptr) {
            glMultiDrawElementsEXT(mode, count, type, indices, primcount);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glMultiDrawElementsEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMultiDrawElementsEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMultiDrawElementsEXT\n");
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
        GAPID_INFO("glMultiDrawElementsIndirectEXT(%u, %u, %p, %" PRId32 ", %" PRId32 ")\n", mode,
                   type, indirect, drawcount, stride);
        if (glMultiDrawElementsIndirectEXT != nullptr) {
            glMultiDrawElementsIndirectEXT(mode, type, indirect, drawcount, stride);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glMultiDrawElementsIndirectEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glMultiDrawElementsIndirectEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMultiDrawElementsIndirectEXT\n");
        return false;
    }
}

bool callGlPatchParameteriOES(Stack* stack, bool pushReturn) {
    int32_t value = stack->pop<int32_t>();
    GLenum pname = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glPatchParameteriOES(%u, %" PRId32 ")\n", pname, value);
        if (glPatchParameteriOES != nullptr) {
            glPatchParameteriOES(pname, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glPatchParameteriOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPatchParameteriOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPatchParameteriOES\n");
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
        GAPID_INFO("glPathCommandsNV(%" PRIu32 ", %" PRId32 ", %p, %" PRId32 ", %u, %p)\n", path,
                   numCommands, commands, numCoords, coordType, coords);
        if (glPathCommandsNV != nullptr) {
            glPathCommandsNV(path, numCommands, commands, numCoords, coordType, coords);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glPathCommandsNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathCommandsNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathCommandsNV\n");
        return false;
    }
}

bool callGlPathCoordsNV(Stack* stack, bool pushReturn) {
    void* coords = stack->pop<void*>();
    GLenum coordType = stack->pop<GLenum>();
    int32_t numCoords = stack->pop<int32_t>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glPathCoordsNV(%" PRIu32 ", %" PRId32 ", %u, %p)\n", path, numCoords, coordType,
                   coords);
        if (glPathCoordsNV != nullptr) {
            glPathCoordsNV(path, numCoords, coordType, coords);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glPathCoordsNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathCoordsNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathCoordsNV\n");
        return false;
    }
}

bool callGlPathCoverDepthFuncNV(Stack* stack, bool pushReturn) {
    GLenum func = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glPathCoverDepthFuncNV(%u)\n", func);
        if (glPathCoverDepthFuncNV != nullptr) {
            glPathCoverDepthFuncNV(func);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glPathCoverDepthFuncNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathCoverDepthFuncNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathCoverDepthFuncNV\n");
        return false;
    }
}

bool callGlPathDashArrayNV(Stack* stack, bool pushReturn) {
    float* dashArray = stack->pop<float*>();
    int32_t dashCount = stack->pop<int32_t>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glPathDashArrayNV(%" PRIu32 ", %" PRId32 ", %p)\n", path, dashCount, dashArray);
        if (glPathDashArrayNV != nullptr) {
            glPathDashArrayNV(path, dashCount, dashArray);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glPathDashArrayNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathDashArrayNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathDashArrayNV\n");
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
                   ", %" PRIu32 ", %f)\n",
                   firstPathName, fontTarget, fontName, fontStyle, firstGlyphIndex, numGlyphs,
                   pathParameterTemplate, emScale);
        if (glPathGlyphIndexArrayNV != nullptr) {
            GLenum return_value = glPathGlyphIndexArrayNV(firstPathName, fontTarget, fontName,
                                                          fontStyle, firstGlyphIndex, numGlyphs,
                                                          pathParameterTemplate, emScale);
            GAPID_INFO("Returned: %u\n", return_value);
            if (pushReturn) {
                stack->push<GLenum>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glPathGlyphIndexArrayNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathGlyphIndexArrayNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathGlyphIndexArrayNV\n");
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
        GAPID_INFO("glPathGlyphIndexRangeNV(%u, %p, %u, %" PRIu32 ", %f, %" PRIu32 ")\n",
                   fontTarget, fontName, fontStyle, pathParameterTemplate, emScale, baseAndCount);
        if (glPathGlyphIndexRangeNV != nullptr) {
            GLenum return_value = glPathGlyphIndexRangeNV(
                    fontTarget, fontName, fontStyle, pathParameterTemplate, emScale, baseAndCount);
            GAPID_INFO("Returned: %u\n", return_value);
            if (pushReturn) {
                stack->push<GLenum>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glPathGlyphIndexRangeNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathGlyphIndexRangeNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathGlyphIndexRangeNV\n");
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
                   ", %u, %" PRIu32 ", %f)\n",
                   firstPathName, fontTarget, fontName, fontStyle, firstGlyph, numGlyphs,
                   handleMissingGlyphs, pathParameterTemplate, emScale);
        if (glPathGlyphRangeNV != nullptr) {
            glPathGlyphRangeNV(firstPathName, fontTarget, fontName, fontStyle, firstGlyph,
                               numGlyphs, handleMissingGlyphs, pathParameterTemplate, emScale);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glPathGlyphRangeNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathGlyphRangeNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathGlyphRangeNV\n");
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
                   ", %f)\n",
                   firstPathName, fontTarget, fontName, fontStyle, numGlyphs, type, charcodes,
                   handleMissingGlyphs, pathParameterTemplate, emScale);
        if (glPathGlyphsNV != nullptr) {
            glPathGlyphsNV(firstPathName, fontTarget, fontName, fontStyle, numGlyphs, type,
                           charcodes, handleMissingGlyphs, pathParameterTemplate, emScale);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glPathGlyphsNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathGlyphsNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathGlyphsNV\n");
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
                   ", %" PRIu32 ", %" PRId32 ", %" PRIu32 ", %f)\n",
                   firstPathName, fontTarget, fontSize, fontData, faceIndex, firstGlyphIndex,
                   numGlyphs, pathParameterTemplate, emScale);
        if (glPathMemoryGlyphIndexArrayNV != nullptr) {
            GLenum return_value = glPathMemoryGlyphIndexArrayNV(
                    firstPathName, fontTarget, fontSize, fontData, faceIndex, firstGlyphIndex,
                    numGlyphs, pathParameterTemplate, emScale);
            GAPID_INFO("Returned: %u\n", return_value);
            if (pushReturn) {
                stack->push<GLenum>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glPathMemoryGlyphIndexArrayNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathMemoryGlyphIndexArrayNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathMemoryGlyphIndexArrayNV\n");
        return false;
    }
}

bool callGlPathParameterfNV(Stack* stack, bool pushReturn) {
    float value = stack->pop<float>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glPathParameterfNV(%" PRIu32 ", %u, %f)\n", path, pname, value);
        if (glPathParameterfNV != nullptr) {
            glPathParameterfNV(path, pname, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glPathParameterfNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathParameterfNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathParameterfNV\n");
        return false;
    }
}

bool callGlPathParameterfvNV(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glPathParameterfvNV(%" PRIu32 ", %u, %p)\n", path, pname, value);
        if (glPathParameterfvNV != nullptr) {
            glPathParameterfvNV(path, pname, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glPathParameterfvNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathParameterfvNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathParameterfvNV\n");
        return false;
    }
}

bool callGlPathParameteriNV(Stack* stack, bool pushReturn) {
    int32_t value = stack->pop<int32_t>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glPathParameteriNV(%" PRIu32 ", %u, %" PRId32 ")\n", path, pname, value);
        if (glPathParameteriNV != nullptr) {
            glPathParameteriNV(path, pname, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glPathParameteriNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathParameteriNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathParameteriNV\n");
        return false;
    }
}

bool callGlPathParameterivNV(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glPathParameterivNV(%" PRIu32 ", %u, %p)\n", path, pname, value);
        if (glPathParameterivNV != nullptr) {
            glPathParameterivNV(path, pname, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glPathParameterivNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathParameterivNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathParameterivNV\n");
        return false;
    }
}

bool callGlPathStencilDepthOffsetNV(Stack* stack, bool pushReturn) {
    float units = stack->pop<float>();
    float factor = stack->pop<float>();
    if (stack->isValid()) {
        GAPID_INFO("glPathStencilDepthOffsetNV(%f, %f)\n", factor, units);
        if (glPathStencilDepthOffsetNV != nullptr) {
            glPathStencilDepthOffsetNV(factor, units);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glPathStencilDepthOffsetNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathStencilDepthOffsetNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathStencilDepthOffsetNV\n");
        return false;
    }
}

bool callGlPathStencilFuncNV(Stack* stack, bool pushReturn) {
    uint32_t mask = stack->pop<uint32_t>();
    int32_t ref = stack->pop<int32_t>();
    GLenum func = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glPathStencilFuncNV(%u, %" PRId32 ", %" PRIu32 ")\n", func, ref, mask);
        if (glPathStencilFuncNV != nullptr) {
            glPathStencilFuncNV(func, ref, mask);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glPathStencilFuncNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathStencilFuncNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathStencilFuncNV\n");
        return false;
    }
}

bool callGlPathStringNV(Stack* stack, bool pushReturn) {
    void* pathString = stack->pop<void*>();
    int32_t length = stack->pop<int32_t>();
    GLenum format = stack->pop<GLenum>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glPathStringNV(%" PRIu32 ", %u, %" PRId32 ", %p)\n", path, format, length,
                   pathString);
        if (glPathStringNV != nullptr) {
            glPathStringNV(path, format, length, pathString);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glPathStringNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathStringNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathStringNV\n");
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
                   ", %p, %" PRId32 ", %u, %p)\n",
                   path, commandStart, commandsToDelete, numCommands, commands, numCoords,
                   coordType, coords);
        if (glPathSubCommandsNV != nullptr) {
            glPathSubCommandsNV(path, commandStart, commandsToDelete, numCommands, commands,
                                numCoords, coordType, coords);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glPathSubCommandsNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathSubCommandsNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathSubCommandsNV\n");
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
        GAPID_INFO("glPathSubCoordsNV(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %u, %p)\n", path,
                   coordStart, numCoords, coordType, coords);
        if (glPathSubCoordsNV != nullptr) {
            glPathSubCoordsNV(path, coordStart, numCoords, coordType, coords);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glPathSubCoordsNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPathSubCoordsNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPathSubCoordsNV\n");
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
        GAPID_INFO("glPointAlongPathNV(%" PRIu32 ", %" PRId32 ", %" PRId32
                   ", %f, %p, %p, %p, %p)\n",
                   path, startSegment, numSegments, distance, x, y, tangentX, tangentY);
        if (glPointAlongPathNV != nullptr) {
            uint8_t return_value = glPointAlongPathNV(path, startSegment, numSegments, distance, x,
                                                      y, tangentX, tangentY);
            GAPID_INFO("Returned: %" PRIu8 "\n", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glPointAlongPathNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPointAlongPathNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPointAlongPathNV\n");
        return false;
    }
}

bool callGlPolygonModeNV(Stack* stack, bool pushReturn) {
    GLenum mode = stack->pop<GLenum>();
    GLenum face = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glPolygonModeNV(%u, %u)\n", face, mode);
        if (glPolygonModeNV != nullptr) {
            glPolygonModeNV(face, mode);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glPolygonModeNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPolygonModeNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPolygonModeNV\n");
        return false;
    }
}

bool callGlPopGroupMarkerEXT(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glPopGroupMarkerEXT()\n");
        if (glPopGroupMarkerEXT != nullptr) {
            glPopGroupMarkerEXT();
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glPopGroupMarkerEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPopGroupMarkerEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPopGroupMarkerEXT\n");
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
        GAPID_INFO("glPrimitiveBoundingBoxOES(%f, %f, %f, %f, %f, %f, %f, %f)\n", minX, minY, minZ,
                   minW, maxX, maxY, maxZ, maxW);
        if (glPrimitiveBoundingBoxOES != nullptr) {
            glPrimitiveBoundingBoxOES(minX, minY, minZ, minW, maxX, maxY, maxZ, maxW);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glPrimitiveBoundingBoxOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPrimitiveBoundingBoxOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPrimitiveBoundingBoxOES\n");
        return false;
    }
}

bool callGlProgramBinaryOES(Stack* stack, bool pushReturn) {
    int32_t binary_size = stack->pop<int32_t>();
    void* binary = stack->pop<void*>();
    GLenum binary_format = stack->pop<GLenum>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramBinaryOES(%" PRIu32 ", %u, %p, %" PRId32 ")\n", program, binary_format,
                   binary, binary_size);
        if (glProgramBinaryOES != nullptr) {
            glProgramBinaryOES(program, binary_format, binary, binary_size);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramBinaryOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramBinaryOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramBinaryOES\n");
        return false;
    }
}

bool callGlProgramParameteriEXT(Stack* stack, bool pushReturn) {
    int32_t value = stack->pop<int32_t>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramParameteriEXT(%" PRIu32 ", %u, %" PRId32 ")\n", program, pname, value);
        if (glProgramParameteriEXT != nullptr) {
            glProgramParameteriEXT(program, pname, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramParameteriEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramParameteriEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramParameteriEXT\n");
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
        GAPID_INFO("glProgramPathFragmentInputGenNV(%" PRIu32 ", %" PRId32 ", %u, %" PRId32
                   ", %p)\n",
                   program, location, genMode, components, coeffs);
        if (glProgramPathFragmentInputGenNV != nullptr) {
            glProgramPathFragmentInputGenNV(program, location, genMode, components, coeffs);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramPathFragmentInputGenNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glProgramPathFragmentInputGenNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramPathFragmentInputGenNV\n");
        return false;
    }
}

bool callGlProgramUniform1fEXT(Stack* stack, bool pushReturn) {
    float v0 = stack->pop<float>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform1fEXT(%" PRIu32 ", %" PRId32 ", %f)\n", program, location, v0);
        if (glProgramUniform1fEXT != nullptr) {
            glProgramUniform1fEXT(program, location, v0);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform1fEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform1fEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform1fEXT\n");
        return false;
    }
}

bool callGlProgramUniform1fvEXT(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform1fvEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)\n", program,
                   location, count, value);
        if (glProgramUniform1fvEXT != nullptr) {
            glProgramUniform1fvEXT(program, location, count, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform1fvEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform1fvEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform1fvEXT\n");
        return false;
    }
}

bool callGlProgramUniform1iEXT(Stack* stack, bool pushReturn) {
    int32_t v0 = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform1iEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ")\n", program,
                   location, v0);
        if (glProgramUniform1iEXT != nullptr) {
            glProgramUniform1iEXT(program, location, v0);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform1iEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform1iEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform1iEXT\n");
        return false;
    }
}

bool callGlProgramUniform1ivEXT(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform1ivEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)\n", program,
                   location, count, value);
        if (glProgramUniform1ivEXT != nullptr) {
            glProgramUniform1ivEXT(program, location, count, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform1ivEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform1ivEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform1ivEXT\n");
        return false;
    }
}

bool callGlProgramUniform1uiEXT(Stack* stack, bool pushReturn) {
    uint32_t v0 = stack->pop<uint32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform1uiEXT(%" PRIu32 ", %" PRId32 ", %" PRIu32 ")\n", program,
                   location, v0);
        if (glProgramUniform1uiEXT != nullptr) {
            glProgramUniform1uiEXT(program, location, v0);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform1uiEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform1uiEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform1uiEXT\n");
        return false;
    }
}

bool callGlProgramUniform1uivEXT(Stack* stack, bool pushReturn) {
    uint32_t* value = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform1uivEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)\n", program,
                   location, count, value);
        if (glProgramUniform1uivEXT != nullptr) {
            glProgramUniform1uivEXT(program, location, count, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform1uivEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform1uivEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform1uivEXT\n");
        return false;
    }
}

bool callGlProgramUniform2fEXT(Stack* stack, bool pushReturn) {
    float v1 = stack->pop<float>();
    float v0 = stack->pop<float>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform2fEXT(%" PRIu32 ", %" PRId32 ", %f, %f)\n", program, location,
                   v0, v1);
        if (glProgramUniform2fEXT != nullptr) {
            glProgramUniform2fEXT(program, location, v0, v1);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform2fEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform2fEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform2fEXT\n");
        return false;
    }
}

bool callGlProgramUniform2fvEXT(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform2fvEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)\n", program,
                   location, count, value);
        if (glProgramUniform2fvEXT != nullptr) {
            glProgramUniform2fvEXT(program, location, count, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform2fvEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform2fvEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform2fvEXT\n");
        return false;
    }
}

bool callGlProgramUniform2iEXT(Stack* stack, bool pushReturn) {
    int32_t v1 = stack->pop<int32_t>();
    int32_t v0 = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform2iEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ")\n",
                   program, location, v0, v1);
        if (glProgramUniform2iEXT != nullptr) {
            glProgramUniform2iEXT(program, location, v0, v1);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform2iEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform2iEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform2iEXT\n");
        return false;
    }
}

bool callGlProgramUniform2ivEXT(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform2ivEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)\n", program,
                   location, count, value);
        if (glProgramUniform2ivEXT != nullptr) {
            glProgramUniform2ivEXT(program, location, count, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform2ivEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform2ivEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform2ivEXT\n");
        return false;
    }
}

bool callGlProgramUniform2uiEXT(Stack* stack, bool pushReturn) {
    uint32_t v1 = stack->pop<uint32_t>();
    uint32_t v0 = stack->pop<uint32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform2uiEXT(%" PRIu32 ", %" PRId32 ", %" PRIu32 ", %" PRIu32 ")\n",
                   program, location, v0, v1);
        if (glProgramUniform2uiEXT != nullptr) {
            glProgramUniform2uiEXT(program, location, v0, v1);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform2uiEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform2uiEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform2uiEXT\n");
        return false;
    }
}

bool callGlProgramUniform2uivEXT(Stack* stack, bool pushReturn) {
    uint32_t* value = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform2uivEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)\n", program,
                   location, count, value);
        if (glProgramUniform2uivEXT != nullptr) {
            glProgramUniform2uivEXT(program, location, count, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform2uivEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform2uivEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform2uivEXT\n");
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
        GAPID_INFO("glProgramUniform3fEXT(%" PRIu32 ", %" PRId32 ", %f, %f, %f)\n", program,
                   location, v0, v1, v2);
        if (glProgramUniform3fEXT != nullptr) {
            glProgramUniform3fEXT(program, location, v0, v1, v2);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform3fEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform3fEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform3fEXT\n");
        return false;
    }
}

bool callGlProgramUniform3fvEXT(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform3fvEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)\n", program,
                   location, count, value);
        if (glProgramUniform3fvEXT != nullptr) {
            glProgramUniform3fvEXT(program, location, count, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform3fvEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform3fvEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform3fvEXT\n");
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
                   ", %" PRId32 ")\n",
                   program, location, v0, v1, v2);
        if (glProgramUniform3iEXT != nullptr) {
            glProgramUniform3iEXT(program, location, v0, v1, v2);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform3iEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform3iEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform3iEXT\n");
        return false;
    }
}

bool callGlProgramUniform3ivEXT(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform3ivEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)\n", program,
                   location, count, value);
        if (glProgramUniform3ivEXT != nullptr) {
            glProgramUniform3ivEXT(program, location, count, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform3ivEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform3ivEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform3ivEXT\n");
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
                   ", %" PRIu32 ")\n",
                   program, location, v0, v1, v2);
        if (glProgramUniform3uiEXT != nullptr) {
            glProgramUniform3uiEXT(program, location, v0, v1, v2);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform3uiEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform3uiEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform3uiEXT\n");
        return false;
    }
}

bool callGlProgramUniform3uivEXT(Stack* stack, bool pushReturn) {
    uint32_t* value = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform3uivEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)\n", program,
                   location, count, value);
        if (glProgramUniform3uivEXT != nullptr) {
            glProgramUniform3uivEXT(program, location, count, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform3uivEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform3uivEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform3uivEXT\n");
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
        GAPID_INFO("glProgramUniform4fEXT(%" PRIu32 ", %" PRId32 ", %f, %f, %f, %f)\n", program,
                   location, v0, v1, v2, v3);
        if (glProgramUniform4fEXT != nullptr) {
            glProgramUniform4fEXT(program, location, v0, v1, v2, v3);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform4fEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform4fEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform4fEXT\n");
        return false;
    }
}

bool callGlProgramUniform4fvEXT(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform4fvEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)\n", program,
                   location, count, value);
        if (glProgramUniform4fvEXT != nullptr) {
            glProgramUniform4fvEXT(program, location, count, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform4fvEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform4fvEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform4fvEXT\n");
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
                   ", %" PRId32 ", %" PRId32 ")\n",
                   program, location, v0, v1, v2, v3);
        if (glProgramUniform4iEXT != nullptr) {
            glProgramUniform4iEXT(program, location, v0, v1, v2, v3);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform4iEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform4iEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform4iEXT\n");
        return false;
    }
}

bool callGlProgramUniform4ivEXT(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform4ivEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)\n", program,
                   location, count, value);
        if (glProgramUniform4ivEXT != nullptr) {
            glProgramUniform4ivEXT(program, location, count, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform4ivEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform4ivEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform4ivEXT\n");
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
                   ", %" PRIu32 ", %" PRIu32 ")\n",
                   program, location, v0, v1, v2, v3);
        if (glProgramUniform4uiEXT != nullptr) {
            glProgramUniform4uiEXT(program, location, v0, v1, v2, v3);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform4uiEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform4uiEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform4uiEXT\n");
        return false;
    }
}

bool callGlProgramUniform4uivEXT(Stack* stack, bool pushReturn) {
    uint32_t* value = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform4uivEXT(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)\n", program,
                   location, count, value);
        if (glProgramUniform4uivEXT != nullptr) {
            glProgramUniform4uivEXT(program, location, count, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform4uivEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform4uivEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform4uivEXT\n");
        return false;
    }
}

bool callGlProgramUniformHandleui64NV(Stack* stack, bool pushReturn) {
    uint64_t value = stack->pop<uint64_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniformHandleui64NV(%" PRIu32 ", %" PRId32 ", %" PRIu64 ")\n", program,
                   location, value);
        if (glProgramUniformHandleui64NV != nullptr) {
            glProgramUniformHandleui64NV(program, location, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniformHandleui64NV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformHandleui64NV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformHandleui64NV\n");
        return false;
    }
}

bool callGlProgramUniformHandleui64vNV(Stack* stack, bool pushReturn) {
    uint64_t* values = stack->pop<uint64_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniformHandleui64vNV(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)\n",
                   program, location, count, values);
        if (glProgramUniformHandleui64vNV != nullptr) {
            glProgramUniformHandleui64vNV(program, location, count, values);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniformHandleui64vNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformHandleui64vNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformHandleui64vNV\n");
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
                   ", %p)\n",
                   program, location, count, transpose, value);
        if (glProgramUniformMatrix2fvEXT != nullptr) {
            glProgramUniformMatrix2fvEXT(program, location, count, transpose, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniformMatrix2fvEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformMatrix2fvEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix2fvEXT\n");
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
                   ", %p)\n",
                   program, location, count, transpose, value);
        if (glProgramUniformMatrix2x3fvEXT != nullptr) {
            glProgramUniformMatrix2x3fvEXT(program, location, count, transpose, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniformMatrix2x3fvEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glProgramUniformMatrix2x3fvEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix2x3fvEXT\n");
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
                   ", %p)\n",
                   program, location, count, transpose, value);
        if (glProgramUniformMatrix2x4fvEXT != nullptr) {
            glProgramUniformMatrix2x4fvEXT(program, location, count, transpose, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniformMatrix2x4fvEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glProgramUniformMatrix2x4fvEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix2x4fvEXT\n");
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
                   ", %p)\n",
                   program, location, count, transpose, value);
        if (glProgramUniformMatrix3fvEXT != nullptr) {
            glProgramUniformMatrix3fvEXT(program, location, count, transpose, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniformMatrix3fvEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformMatrix3fvEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix3fvEXT\n");
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
                   ", %p)\n",
                   program, location, count, transpose, value);
        if (glProgramUniformMatrix3x2fvEXT != nullptr) {
            glProgramUniformMatrix3x2fvEXT(program, location, count, transpose, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniformMatrix3x2fvEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glProgramUniformMatrix3x2fvEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix3x2fvEXT\n");
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
                   ", %p)\n",
                   program, location, count, transpose, value);
        if (glProgramUniformMatrix3x4fvEXT != nullptr) {
            glProgramUniformMatrix3x4fvEXT(program, location, count, transpose, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniformMatrix3x4fvEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glProgramUniformMatrix3x4fvEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix3x4fvEXT\n");
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
                   ", %p)\n",
                   program, location, count, transpose, value);
        if (glProgramUniformMatrix4fvEXT != nullptr) {
            glProgramUniformMatrix4fvEXT(program, location, count, transpose, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniformMatrix4fvEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformMatrix4fvEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix4fvEXT\n");
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
                   ", %p)\n",
                   program, location, count, transpose, value);
        if (glProgramUniformMatrix4x2fvEXT != nullptr) {
            glProgramUniformMatrix4x2fvEXT(program, location, count, transpose, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniformMatrix4x2fvEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glProgramUniformMatrix4x2fvEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix4x2fvEXT\n");
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
                   ", %p)\n",
                   program, location, count, transpose, value);
        if (glProgramUniformMatrix4x3fvEXT != nullptr) {
            glProgramUniformMatrix4x3fvEXT(program, location, count, transpose, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniformMatrix4x3fvEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glProgramUniformMatrix4x3fvEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix4x3fvEXT\n");
        return false;
    }
}

bool callGlPushGroupMarkerEXT(Stack* stack, bool pushReturn) {
    char* marker = stack->pop<char*>();
    int32_t length = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glPushGroupMarkerEXT(%" PRId32 ", %p)\n", length, marker);
        if (glPushGroupMarkerEXT != nullptr) {
            glPushGroupMarkerEXT(length, marker);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glPushGroupMarkerEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPushGroupMarkerEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPushGroupMarkerEXT\n");
        return false;
    }
}

bool callGlQueryCounterEXT(Stack* stack, bool pushReturn) {
    GLenum target = stack->pop<GLenum>();
    uint32_t query = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glQueryCounterEXT(%" PRIu32 ", %u)\n", query, target);
        if (glQueryCounterEXT != nullptr) {
            glQueryCounterEXT(query, target);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glQueryCounterEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glQueryCounterEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glQueryCounterEXT\n");
        return false;
    }
}

bool callGlReadBufferIndexedEXT(Stack* stack, bool pushReturn) {
    int32_t index = stack->pop<int32_t>();
    GLenum src = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glReadBufferIndexedEXT(%u, %" PRId32 ")\n", src, index);
        if (glReadBufferIndexedEXT != nullptr) {
            glReadBufferIndexedEXT(src, index);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glReadBufferIndexedEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glReadBufferIndexedEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glReadBufferIndexedEXT\n");
        return false;
    }
}

bool callGlReadBufferNV(Stack* stack, bool pushReturn) {
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glReadBufferNV(%u)\n", mode);
        if (glReadBufferNV != nullptr) {
            glReadBufferNV(mode);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glReadBufferNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glReadBufferNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glReadBufferNV\n");
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
                   ", %u, %u, %" PRId32 ", %p)\n",
                   x, y, width, height, format, type, bufSize, data);
        if (glReadnPixelsEXT != nullptr) {
            glReadnPixelsEXT(x, y, width, height, format, type, bufSize, data);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glReadnPixelsEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glReadnPixelsEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glReadnPixelsEXT\n");
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
                   ", %u, %u, %" PRId32 ", %p)\n",
                   x, y, width, height, format, type, bufSize, data);
        if (glReadnPixelsKHR != nullptr) {
            glReadnPixelsKHR(x, y, width, height, format, type, bufSize, data);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glReadnPixelsKHR returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glReadnPixelsKHR\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glReadnPixelsKHR\n");
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
                   ", %" PRId32 ")\n",
                   target, samples, internalformat, width, height);
        if (glRenderbufferStorageMultisampleANGLE != nullptr) {
            glRenderbufferStorageMultisampleANGLE(target, samples, internalformat, width, height);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glRenderbufferStorageMultisampleANGLE returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function "
                    "glRenderbufferStorageMultisampleANGLE\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glRenderbufferStorageMultisampleANGLE\n");
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
                   ", %" PRId32 ")\n",
                   target, samples, internalformat, width, height);
        if (glRenderbufferStorageMultisampleAPPLE != nullptr) {
            glRenderbufferStorageMultisampleAPPLE(target, samples, internalformat, width, height);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glRenderbufferStorageMultisampleAPPLE returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function "
                    "glRenderbufferStorageMultisampleAPPLE\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glRenderbufferStorageMultisampleAPPLE\n");
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
                   ")\n",
                   target, samples, internalformat, width, height);
        if (glRenderbufferStorageMultisampleEXT != nullptr) {
            glRenderbufferStorageMultisampleEXT(target, samples, internalformat, width, height);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glRenderbufferStorageMultisampleEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glRenderbufferStorageMultisampleEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glRenderbufferStorageMultisampleEXT\n");
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
                   ")\n",
                   target, samples, internalformat, width, height);
        if (glRenderbufferStorageMultisampleIMG != nullptr) {
            glRenderbufferStorageMultisampleIMG(target, samples, internalformat, width, height);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glRenderbufferStorageMultisampleIMG returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glRenderbufferStorageMultisampleIMG\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glRenderbufferStorageMultisampleIMG\n");
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
                   ")\n",
                   target, samples, internalformat, width, height);
        if (glRenderbufferStorageMultisampleNV != nullptr) {
            glRenderbufferStorageMultisampleNV(target, samples, internalformat, width, height);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glRenderbufferStorageMultisampleNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glRenderbufferStorageMultisampleNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glRenderbufferStorageMultisampleNV\n");
        return false;
    }
}

bool callGlResolveMultisampleFramebufferAPPLE(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glResolveMultisampleFramebufferAPPLE()\n");
        if (glResolveMultisampleFramebufferAPPLE != nullptr) {
            glResolveMultisampleFramebufferAPPLE();
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glResolveMultisampleFramebufferAPPLE returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function "
                    "glResolveMultisampleFramebufferAPPLE\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glResolveMultisampleFramebufferAPPLE\n");
        return false;
    }
}

bool callGlSamplerParameterIivOES(Stack* stack, bool pushReturn) {
    int32_t* param = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glSamplerParameterIivOES(%" PRIu32 ", %u, %p)\n", sampler, pname, param);
        if (glSamplerParameterIivOES != nullptr) {
            glSamplerParameterIivOES(sampler, pname, param);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glSamplerParameterIivOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glSamplerParameterIivOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glSamplerParameterIivOES\n");
        return false;
    }
}

bool callGlSamplerParameterIuivOES(Stack* stack, bool pushReturn) {
    uint32_t* param = stack->pop<uint32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glSamplerParameterIuivOES(%" PRIu32 ", %u, %p)\n", sampler, pname, param);
        if (glSamplerParameterIuivOES != nullptr) {
            glSamplerParameterIuivOES(sampler, pname, param);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glSamplerParameterIuivOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glSamplerParameterIuivOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glSamplerParameterIuivOES\n");
        return false;
    }
}

bool callGlScissorArrayvNV(Stack* stack, bool pushReturn) {
    int32_t* v = stack->pop<int32_t*>();
    int32_t count = stack->pop<int32_t>();
    uint32_t first = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glScissorArrayvNV(%" PRIu32 ", %" PRId32 ", %p)\n", first, count, v);
        if (glScissorArrayvNV != nullptr) {
            glScissorArrayvNV(first, count, v);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glScissorArrayvNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glScissorArrayvNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glScissorArrayvNV\n");
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
                   ")\n",
                   index, left, bottom, width, height);
        if (glScissorIndexedNV != nullptr) {
            glScissorIndexedNV(index, left, bottom, width, height);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glScissorIndexedNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glScissorIndexedNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glScissorIndexedNV\n");
        return false;
    }
}

bool callGlScissorIndexedvNV(Stack* stack, bool pushReturn) {
    int32_t* v = stack->pop<int32_t*>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glScissorIndexedvNV(%" PRIu32 ", %p)\n", index, v);
        if (glScissorIndexedvNV != nullptr) {
            glScissorIndexedvNV(index, v);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glScissorIndexedvNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glScissorIndexedvNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glScissorIndexedvNV\n");
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
                   ", %p)\n",
                   monitor, enable, group, numCounters, counterList);
        if (glSelectPerfMonitorCountersAMD != nullptr) {
            glSelectPerfMonitorCountersAMD(monitor, enable, group, numCounters, counterList);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glSelectPerfMonitorCountersAMD returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glSelectPerfMonitorCountersAMD\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glSelectPerfMonitorCountersAMD\n");
        return false;
    }
}

bool callGlSetFenceNV(Stack* stack, bool pushReturn) {
    GLenum condition = stack->pop<GLenum>();
    uint32_t fence = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glSetFenceNV(%" PRIu32 ", %u)\n", fence, condition);
        if (glSetFenceNV != nullptr) {
            glSetFenceNV(fence, condition);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glSetFenceNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glSetFenceNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glSetFenceNV\n");
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
        GAPID_INFO("glStartTilingQCOM(%" PRIu32 ", %" PRIu32 ", %" PRIu32 ", %" PRIu32 ", %u)\n", x,
                   y, width, height, preserveMask);
        if (glStartTilingQCOM != nullptr) {
            glStartTilingQCOM(x, y, width, height, preserveMask);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glStartTilingQCOM returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glStartTilingQCOM\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStartTilingQCOM\n");
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
                   ", %u, %p)\n",
                   numPaths, pathNameType, paths, pathBase, fillMode, mask, transformType,
                   transformValues);
        if (glStencilFillPathInstancedNV != nullptr) {
            glStencilFillPathInstancedNV(numPaths, pathNameType, paths, pathBase, fillMode, mask,
                                         transformType, transformValues);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glStencilFillPathInstancedNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glStencilFillPathInstancedNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStencilFillPathInstancedNV\n");
        return false;
    }
}

bool callGlStencilFillPathNV(Stack* stack, bool pushReturn) {
    uint32_t mask = stack->pop<uint32_t>();
    GLenum fillMode = stack->pop<GLenum>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glStencilFillPathNV(%" PRIu32 ", %u, %" PRIu32 ")\n", path, fillMode, mask);
        if (glStencilFillPathNV != nullptr) {
            glStencilFillPathNV(path, fillMode, mask);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glStencilFillPathNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glStencilFillPathNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStencilFillPathNV\n");
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
                   ", %" PRIu32 ", %u, %p)\n",
                   numPaths, pathNameType, paths, pathBase, reference, mask, transformType,
                   transformValues);
        if (glStencilStrokePathInstancedNV != nullptr) {
            glStencilStrokePathInstancedNV(numPaths, pathNameType, paths, pathBase, reference, mask,
                                           transformType, transformValues);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glStencilStrokePathInstancedNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glStencilStrokePathInstancedNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStencilStrokePathInstancedNV\n");
        return false;
    }
}

bool callGlStencilStrokePathNV(Stack* stack, bool pushReturn) {
    uint32_t mask = stack->pop<uint32_t>();
    int32_t reference = stack->pop<int32_t>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glStencilStrokePathNV(%" PRIu32 ", %" PRId32 ", %" PRIu32 ")\n", path,
                   reference, mask);
        if (glStencilStrokePathNV != nullptr) {
            glStencilStrokePathNV(path, reference, mask);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glStencilStrokePathNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glStencilStrokePathNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStencilStrokePathNV\n");
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
                   ", %u, %" PRIu32 ", %u, %u, %p)\n",
                   numPaths, pathNameType, paths, pathBase, fillMode, mask, coverMode,
                   transformType, transformValues);
        if (glStencilThenCoverFillPathInstancedNV != nullptr) {
            glStencilThenCoverFillPathInstancedNV(numPaths, pathNameType, paths, pathBase, fillMode,
                                                  mask, coverMode, transformType, transformValues);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glStencilThenCoverFillPathInstancedNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function "
                    "glStencilThenCoverFillPathInstancedNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStencilThenCoverFillPathInstancedNV\n");
        return false;
    }
}

bool callGlStencilThenCoverFillPathNV(Stack* stack, bool pushReturn) {
    GLenum coverMode = stack->pop<GLenum>();
    uint32_t mask = stack->pop<uint32_t>();
    GLenum fillMode = stack->pop<GLenum>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glStencilThenCoverFillPathNV(%" PRIu32 ", %u, %" PRIu32 ", %u)\n", path,
                   fillMode, mask, coverMode);
        if (glStencilThenCoverFillPathNV != nullptr) {
            glStencilThenCoverFillPathNV(path, fillMode, mask, coverMode);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glStencilThenCoverFillPathNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glStencilThenCoverFillPathNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStencilThenCoverFillPathNV\n");
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
                   ", %" PRId32 ", %" PRIu32 ", %u, %u, %p)\n",
                   numPaths, pathNameType, paths, pathBase, reference, mask, coverMode,
                   transformType, transformValues);
        if (glStencilThenCoverStrokePathInstancedNV != nullptr) {
            glStencilThenCoverStrokePathInstancedNV(numPaths, pathNameType, paths, pathBase,
                                                    reference, mask, coverMode, transformType,
                                                    transformValues);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glStencilThenCoverStrokePathInstancedNV returned error: 0x%x\n",
                              err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function "
                    "glStencilThenCoverStrokePathInstancedNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStencilThenCoverStrokePathInstancedNV\n");
        return false;
    }
}

bool callGlStencilThenCoverStrokePathNV(Stack* stack, bool pushReturn) {
    GLenum coverMode = stack->pop<GLenum>();
    uint32_t mask = stack->pop<uint32_t>();
    int32_t reference = stack->pop<int32_t>();
    uint32_t path = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glStencilThenCoverStrokePathNV(%" PRIu32 ", %" PRId32 ", %" PRIu32 ", %u)\n",
                   path, reference, mask, coverMode);
        if (glStencilThenCoverStrokePathNV != nullptr) {
            glStencilThenCoverStrokePathNV(path, reference, mask, coverMode);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glStencilThenCoverStrokePathNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glStencilThenCoverStrokePathNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStencilThenCoverStrokePathNV\n");
        return false;
    }
}

bool callGlTestFenceNV(Stack* stack, bool pushReturn) {
    uint32_t fence = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glTestFenceNV(%" PRIu32 ")\n", fence);
        if (glTestFenceNV != nullptr) {
            uint8_t return_value = glTestFenceNV(fence);
            GAPID_INFO("Returned: %" PRIu8 "\n", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTestFenceNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTestFenceNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTestFenceNV\n");
        return false;
    }
}

bool callGlTexBufferOES(Stack* stack, bool pushReturn) {
    uint32_t buffer = stack->pop<uint32_t>();
    GLenum internalformat = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexBufferOES(%u, %u, %" PRIu32 ")\n", target, internalformat, buffer);
        if (glTexBufferOES != nullptr) {
            glTexBufferOES(target, internalformat, buffer);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTexBufferOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexBufferOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexBufferOES\n");
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
        GAPID_INFO("glTexBufferRangeOES(%u, %u, %" PRIu32 ", %" PRId32 ", %" PRId32 ")\n", target,
                   internalformat, buffer, offset, size);
        if (glTexBufferRangeOES != nullptr) {
            glTexBufferRangeOES(target, internalformat, buffer, offset, size);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTexBufferRangeOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexBufferRangeOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexBufferRangeOES\n");
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
                   ", %" PRId32 ", %u, %u, %p)\n",
                   target, level, internalformat, width, height, depth, border, format, type,
                   pixels);
        if (glTexImage3DOES != nullptr) {
            glTexImage3DOES(target, level, internalformat, width, height, depth, border, format,
                            type, pixels);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTexImage3DOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexImage3DOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexImage3DOES\n");
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
                   ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRIu8 ")\n",
                   target, level, xoffset, yoffset, zoffset, width, height, depth, commit);
        if (glTexPageCommitmentARB != nullptr) {
            glTexPageCommitmentARB(target, level, xoffset, yoffset, zoffset, width, height, depth,
                                   commit);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTexPageCommitmentARB returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexPageCommitmentARB\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexPageCommitmentARB\n");
        return false;
    }
}

bool callGlTexParameterIivOES(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexParameterIivOES(%u, %u, %p)\n", target, pname, params);
        if (glTexParameterIivOES != nullptr) {
            glTexParameterIivOES(target, pname, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTexParameterIivOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexParameterIivOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexParameterIivOES\n");
        return false;
    }
}

bool callGlTexParameterIuivOES(Stack* stack, bool pushReturn) {
    uint32_t* params = stack->pop<uint32_t*>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexParameterIuivOES(%u, %u, %p)\n", target, pname, params);
        if (glTexParameterIuivOES != nullptr) {
            glTexParameterIuivOES(target, pname, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTexParameterIuivOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexParameterIuivOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexParameterIuivOES\n");
        return false;
    }
}

bool callGlTexStorage1DEXT(Stack* stack, bool pushReturn) {
    int32_t width = stack->pop<int32_t>();
    GLenum format = stack->pop<GLenum>();
    int32_t levels = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexStorage1DEXT(%u, %" PRId32 ", %u, %" PRId32 ")\n", target, levels, format,
                   width);
        if (glTexStorage1DEXT != nullptr) {
            glTexStorage1DEXT(target, levels, format, width);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTexStorage1DEXT returned error: 0x%x\n", err);
            }
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
    GLenum format = stack->pop<GLenum>();
    int32_t levels = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexStorage2DEXT(%u, %" PRId32 ", %u, %" PRId32 ", %" PRId32 ")\n", target,
                   levels, format, width, height);
        if (glTexStorage2DEXT != nullptr) {
            glTexStorage2DEXT(target, levels, format, width, height);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTexStorage2DEXT returned error: 0x%x\n", err);
            }
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
    GLenum format = stack->pop<GLenum>();
    int32_t levels = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexStorage3DEXT(%u, %" PRId32 ", %u, %" PRId32 ", %" PRId32 ", %" PRId32
                   ")\n",
                   target, levels, format, width, height, depth);
        if (glTexStorage3DEXT != nullptr) {
            glTexStorage3DEXT(target, levels, format, width, height, depth);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTexStorage3DEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexStorage3DEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexStorage3DEXT\n");
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
                   ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %u, %u, %p)\n",
                   target, level, xoffset, yoffset, zoffset, width, height, depth, format, type,
                   pixels);
        if (glTexSubImage3DOES != nullptr) {
            glTexSubImage3DOES(target, level, xoffset, yoffset, zoffset, width, height, depth,
                               format, type, pixels);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTexSubImage3DOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexSubImage3DOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexSubImage3DOES\n");
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
        GAPID_INFO("glTextureStorage1DEXT(%" PRIu32 ", %u, %" PRId32 ", %u, %" PRId32 ")\n",
                   texture, target, levels, format, width);
        if (glTextureStorage1DEXT != nullptr) {
            glTextureStorage1DEXT(texture, target, levels, format, width);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTextureStorage1DEXT returned error: 0x%x\n", err);
            }
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
    GLenum format = stack->pop<GLenum>();
    int32_t levels = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    uint32_t texture = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glTextureStorage2DEXT(%" PRIu32 ", %u, %" PRId32 ", %u, %" PRId32 ", %" PRId32
                   ")\n",
                   texture, target, levels, format, width, height);
        if (glTextureStorage2DEXT != nullptr) {
            glTextureStorage2DEXT(texture, target, levels, format, width, height);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTextureStorage2DEXT returned error: 0x%x\n", err);
            }
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
    GLenum format = stack->pop<GLenum>();
    int32_t levels = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    uint32_t texture = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glTextureStorage3DEXT(%" PRIu32 ", %u, %" PRId32 ", %u, %" PRId32 ", %" PRId32
                   ", %" PRId32 ")\n",
                   texture, target, levels, format, width, height, depth);
        if (glTextureStorage3DEXT != nullptr) {
            glTextureStorage3DEXT(texture, target, levels, format, width, height, depth);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTextureStorage3DEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTextureStorage3DEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTextureStorage3DEXT\n");
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
                   ", %" PRIu32 ", %" PRIu32 ")\n",
                   texture, target, origtexture, internalformat, minlevel, numlevels, minlayer,
                   numlayers);
        if (glTextureViewEXT != nullptr) {
            glTextureViewEXT(texture, target, origtexture, internalformat, minlevel, numlevels,
                             minlayer, numlayers);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTextureViewEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTextureViewEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTextureViewEXT\n");
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
                   ", %" PRIu32 ", %" PRIu32 ")\n",
                   texture, target, origtexture, internalformat, minlevel, numlevels, minlayer,
                   numlayers);
        if (glTextureViewOES != nullptr) {
            glTextureViewOES(texture, target, origtexture, internalformat, minlevel, numlevels,
                             minlayer, numlayers);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTextureViewOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTextureViewOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTextureViewOES\n");
        return false;
    }
}

bool callGlTransformPathNV(Stack* stack, bool pushReturn) {
    float* transformValues = stack->pop<float*>();
    GLenum transformType = stack->pop<GLenum>();
    uint32_t srcPath = stack->pop<uint32_t>();
    uint32_t resultPath = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glTransformPathNV(%" PRIu32 ", %" PRIu32 ", %u, %p)\n", resultPath, srcPath,
                   transformType, transformValues);
        if (glTransformPathNV != nullptr) {
            glTransformPathNV(resultPath, srcPath, transformType, transformValues);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTransformPathNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTransformPathNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTransformPathNV\n");
        return false;
    }
}

bool callGlUniformHandleui64NV(Stack* stack, bool pushReturn) {
    uint64_t value = stack->pop<uint64_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformHandleui64NV(%" PRId32 ", %" PRIu64 ")\n", location, value);
        if (glUniformHandleui64NV != nullptr) {
            glUniformHandleui64NV(location, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniformHandleui64NV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformHandleui64NV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformHandleui64NV\n");
        return false;
    }
}

bool callGlUniformHandleui64vNV(Stack* stack, bool pushReturn) {
    uint64_t* value = stack->pop<uint64_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformHandleui64vNV(%" PRId32 ", %" PRId32 ", %p)\n", location, count,
                   value);
        if (glUniformHandleui64vNV != nullptr) {
            glUniformHandleui64vNV(location, count, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniformHandleui64vNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformHandleui64vNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformHandleui64vNV\n");
        return false;
    }
}

bool callGlUniformMatrix2x3fvNV(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformMatrix2x3fvNV(%" PRId32 ", %" PRId32 ", %" PRIu8 ", %p)\n", location,
                   count, transpose, value);
        if (glUniformMatrix2x3fvNV != nullptr) {
            glUniformMatrix2x3fvNV(location, count, transpose, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniformMatrix2x3fvNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformMatrix2x3fvNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformMatrix2x3fvNV\n");
        return false;
    }
}

bool callGlUniformMatrix2x4fvNV(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformMatrix2x4fvNV(%" PRId32 ", %" PRId32 ", %" PRIu8 ", %p)\n", location,
                   count, transpose, value);
        if (glUniformMatrix2x4fvNV != nullptr) {
            glUniformMatrix2x4fvNV(location, count, transpose, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniformMatrix2x4fvNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformMatrix2x4fvNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformMatrix2x4fvNV\n");
        return false;
    }
}

bool callGlUniformMatrix3x2fvNV(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformMatrix3x2fvNV(%" PRId32 ", %" PRId32 ", %" PRIu8 ", %p)\n", location,
                   count, transpose, value);
        if (glUniformMatrix3x2fvNV != nullptr) {
            glUniformMatrix3x2fvNV(location, count, transpose, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniformMatrix3x2fvNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformMatrix3x2fvNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformMatrix3x2fvNV\n");
        return false;
    }
}

bool callGlUniformMatrix3x4fvNV(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformMatrix3x4fvNV(%" PRId32 ", %" PRId32 ", %" PRIu8 ", %p)\n", location,
                   count, transpose, value);
        if (glUniformMatrix3x4fvNV != nullptr) {
            glUniformMatrix3x4fvNV(location, count, transpose, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniformMatrix3x4fvNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformMatrix3x4fvNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformMatrix3x4fvNV\n");
        return false;
    }
}

bool callGlUniformMatrix4x2fvNV(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformMatrix4x2fvNV(%" PRId32 ", %" PRId32 ", %" PRIu8 ", %p)\n", location,
                   count, transpose, value);
        if (glUniformMatrix4x2fvNV != nullptr) {
            glUniformMatrix4x2fvNV(location, count, transpose, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniformMatrix4x2fvNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformMatrix4x2fvNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformMatrix4x2fvNV\n");
        return false;
    }
}

bool callGlUniformMatrix4x3fvNV(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformMatrix4x3fvNV(%" PRId32 ", %" PRId32 ", %" PRIu8 ", %p)\n", location,
                   count, transpose, value);
        if (glUniformMatrix4x3fvNV != nullptr) {
            glUniformMatrix4x3fvNV(location, count, transpose, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniformMatrix4x3fvNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformMatrix4x3fvNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformMatrix4x3fvNV\n");
        return false;
    }
}

bool callGlUnmapBufferOES(Stack* stack, bool pushReturn) {
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glUnmapBufferOES(%u)\n", target);
        if (glUnmapBufferOES != nullptr) {
            uint8_t return_value = glUnmapBufferOES(target);
            GAPID_INFO("Returned: %" PRIu8 "\n", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUnmapBufferOES returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUnmapBufferOES\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUnmapBufferOES\n");
        return false;
    }
}

bool callGlUseProgramStagesEXT(Stack* stack, bool pushReturn) {
    uint32_t program = stack->pop<uint32_t>();
    GLbitfield stages = stack->pop<GLbitfield>();
    uint32_t pipeline = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUseProgramStagesEXT(%" PRIu32 ", %u, %" PRIu32 ")\n", pipeline, stages,
                   program);
        if (glUseProgramStagesEXT != nullptr) {
            glUseProgramStagesEXT(pipeline, stages, program);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUseProgramStagesEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUseProgramStagesEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUseProgramStagesEXT\n");
        return false;
    }
}

bool callGlValidateProgramPipelineEXT(Stack* stack, bool pushReturn) {
    uint32_t pipeline = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glValidateProgramPipelineEXT(%" PRIu32 ")\n", pipeline);
        if (glValidateProgramPipelineEXT != nullptr) {
            glValidateProgramPipelineEXT(pipeline);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glValidateProgramPipelineEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glValidateProgramPipelineEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glValidateProgramPipelineEXT\n");
        return false;
    }
}

bool callGlVertexAttribDivisorANGLE(Stack* stack, bool pushReturn) {
    uint32_t divisor = stack->pop<uint32_t>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttribDivisorANGLE(%" PRIu32 ", %" PRIu32 ")\n", index, divisor);
        if (glVertexAttribDivisorANGLE != nullptr) {
            glVertexAttribDivisorANGLE(index, divisor);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glVertexAttribDivisorANGLE returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttribDivisorANGLE\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttribDivisorANGLE\n");
        return false;
    }
}

bool callGlVertexAttribDivisorEXT(Stack* stack, bool pushReturn) {
    uint32_t divisor = stack->pop<uint32_t>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttribDivisorEXT(%" PRIu32 ", %" PRIu32 ")\n", index, divisor);
        if (glVertexAttribDivisorEXT != nullptr) {
            glVertexAttribDivisorEXT(index, divisor);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glVertexAttribDivisorEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttribDivisorEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttribDivisorEXT\n");
        return false;
    }
}

bool callGlVertexAttribDivisorNV(Stack* stack, bool pushReturn) {
    uint32_t divisor = stack->pop<uint32_t>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttribDivisorNV(%" PRIu32 ", %" PRIu32 ")\n", index, divisor);
        if (glVertexAttribDivisorNV != nullptr) {
            glVertexAttribDivisorNV(index, divisor);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glVertexAttribDivisorNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttribDivisorNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttribDivisorNV\n");
        return false;
    }
}

bool callGlViewportArrayvNV(Stack* stack, bool pushReturn) {
    float* v = stack->pop<float*>();
    int32_t count = stack->pop<int32_t>();
    uint32_t first = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glViewportArrayvNV(%" PRIu32 ", %" PRId32 ", %p)\n", first, count, v);
        if (glViewportArrayvNV != nullptr) {
            glViewportArrayvNV(first, count, v);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glViewportArrayvNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glViewportArrayvNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glViewportArrayvNV\n");
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
        GAPID_INFO("glViewportIndexedfNV(%" PRIu32 ", %f, %f, %f, %f)\n", index, x, y, w, h);
        if (glViewportIndexedfNV != nullptr) {
            glViewportIndexedfNV(index, x, y, w, h);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glViewportIndexedfNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glViewportIndexedfNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glViewportIndexedfNV\n");
        return false;
    }
}

bool callGlViewportIndexedfvNV(Stack* stack, bool pushReturn) {
    float* v = stack->pop<float*>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glViewportIndexedfvNV(%" PRIu32 ", %p)\n", index, v);
        if (glViewportIndexedfvNV != nullptr) {
            glViewportIndexedfvNV(index, v);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glViewportIndexedfvNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glViewportIndexedfvNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glViewportIndexedfvNV\n");
        return false;
    }
}

bool callGlWaitSyncAPPLE(Stack* stack, bool pushReturn) {
    uint64_t timeout = stack->pop<uint64_t>();
    GLbitfield flag = stack->pop<GLbitfield>();
    uint64_t sync = stack->pop<uint64_t>();
    if (stack->isValid()) {
        GAPID_INFO("glWaitSyncAPPLE(%" PRIu64 ", %u, %" PRIu64 ")\n", sync, flag, timeout);
        if (glWaitSyncAPPLE != nullptr) {
            glWaitSyncAPPLE(sync, flag, timeout);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glWaitSyncAPPLE returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glWaitSyncAPPLE\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glWaitSyncAPPLE\n");
        return false;
    }
}

bool callGlWeightPathsNV(Stack* stack, bool pushReturn) {
    float* weights = stack->pop<float*>();
    uint32_t* paths = stack->pop<uint32_t*>();
    int32_t numPaths = stack->pop<int32_t>();
    uint32_t resultPath = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glWeightPathsNV(%" PRIu32 ", %" PRId32 ", %p, %p)\n", resultPath, numPaths,
                   paths, weights);
        if (glWeightPathsNV != nullptr) {
            glWeightPathsNV(resultPath, numPaths, paths, weights);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glWeightPathsNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glWeightPathsNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glWeightPathsNV\n");
        return false;
    }
}

bool callGlCoverageModulationNV(Stack* stack, bool pushReturn) {
    GLenum components = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glCoverageModulationNV(%u)\n", components);
        if (glCoverageModulationNV != nullptr) {
            glCoverageModulationNV(components);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glCoverageModulationNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCoverageModulationNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCoverageModulationNV\n");
        return false;
    }
}

bool callGlCoverageModulationTableNV(Stack* stack, bool pushReturn) {
    float* v = stack->pop<float*>();
    int32_t n = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glCoverageModulationTableNV(%" PRId32 ", %p)\n", n, v);
        if (glCoverageModulationTableNV != nullptr) {
            glCoverageModulationTableNV(n, v);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glCoverageModulationTableNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCoverageModulationTableNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCoverageModulationTableNV\n");
        return false;
    }
}

bool callGlFragmentCoverageColorNV(Stack* stack, bool pushReturn) {
    uint32_t color = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glFragmentCoverageColorNV(%" PRIu32 ")\n", color);
        if (glFragmentCoverageColorNV != nullptr) {
            glFragmentCoverageColorNV(color);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glFragmentCoverageColorNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFragmentCoverageColorNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFragmentCoverageColorNV\n");
        return false;
    }
}

bool callGlFramebufferSampleLocationsfvNV(Stack* stack, bool pushReturn) {
    float* v = stack->pop<float*>();
    int32_t count = stack->pop<int32_t>();
    uint32_t start = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glFramebufferSampleLocationsfvNV(%u, %" PRIu32 ", %" PRId32 ", %p)\n", target,
                   start, count, v);
        if (glFramebufferSampleLocationsfvNV != nullptr) {
            glFramebufferSampleLocationsfvNV(target, start, count, v);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glFramebufferSampleLocationsfvNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glFramebufferSampleLocationsfvNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFramebufferSampleLocationsfvNV\n");
        return false;
    }
}

bool callGlGetCoverageModulationTableNV(Stack* stack, bool pushReturn) {
    float* v = stack->pop<float*>();
    int32_t bufsize = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetCoverageModulationTableNV(%" PRId32 ", %p)\n", bufsize, v);
        if (glGetCoverageModulationTableNV != nullptr) {
            glGetCoverageModulationTableNV(bufsize, v);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetCoverageModulationTableNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function glGetCoverageModulationTableNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetCoverageModulationTableNV\n");
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
                   ", %p)\n",
                   framebuffer, start, count, v);
        if (glNamedFramebufferSampleLocationsfvNV != nullptr) {
            glNamedFramebufferSampleLocationsfvNV(framebuffer, start, count, v);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glNamedFramebufferSampleLocationsfvNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING(
                    "Attempted to call unsupported function "
                    "glNamedFramebufferSampleLocationsfvNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glNamedFramebufferSampleLocationsfvNV\n");
        return false;
    }
}

bool callGlRasterSamplesEXT(Stack* stack, bool pushReturn) {
    uint8_t fixedsamplelocations = stack->pop<uint8_t>();
    uint32_t samples = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glRasterSamplesEXT(%" PRIu32 ", %" PRIu8 ")\n", samples, fixedsamplelocations);
        if (glRasterSamplesEXT != nullptr) {
            glRasterSamplesEXT(samples, fixedsamplelocations);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glRasterSamplesEXT returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glRasterSamplesEXT\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glRasterSamplesEXT\n");
        return false;
    }
}

bool callGlResolveDepthValuesNV(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glResolveDepthValuesNV()\n");
        if (glResolveDepthValuesNV != nullptr) {
            glResolveDepthValuesNV();
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glResolveDepthValuesNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glResolveDepthValuesNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glResolveDepthValuesNV\n");
        return false;
    }
}

bool callGlSubpixelPrecisionBiasNV(Stack* stack, bool pushReturn) {
    uint32_t ybits = stack->pop<uint32_t>();
    uint32_t xbits = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glSubpixelPrecisionBiasNV(%" PRIu32 ", %" PRIu32 ")\n", xbits, ybits);
        if (glSubpixelPrecisionBiasNV != nullptr) {
            glSubpixelPrecisionBiasNV(xbits, ybits);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glSubpixelPrecisionBiasNV returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glSubpixelPrecisionBiasNV\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glSubpixelPrecisionBiasNV\n");
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
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBlendColor returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendColor\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendColor\n");
        return false;
    }
}

bool callGlBlendEquation(Stack* stack, bool pushReturn) {
    GLenum equation = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendEquation(%u)\n", equation);
        if (glBlendEquation != nullptr) {
            glBlendEquation(equation);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBlendEquation returned error: 0x%x\n", err);
            }
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
    GLenum alpha = stack->pop<GLenum>();
    GLenum rgb = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendEquationSeparate(%u, %u)\n", rgb, alpha);
        if (glBlendEquationSeparate != nullptr) {
            glBlendEquationSeparate(rgb, alpha);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBlendEquationSeparate returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendEquationSeparate\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendEquationSeparate\n");
        return false;
    }
}

bool callGlBlendFunc(Stack* stack, bool pushReturn) {
    GLenum dst_factor = stack->pop<GLenum>();
    GLenum src_factor = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendFunc(%u, %u)\n", src_factor, dst_factor);
        if (glBlendFunc != nullptr) {
            glBlendFunc(src_factor, dst_factor);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBlendFunc returned error: 0x%x\n", err);
            }
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
    GLenum dst_factor_alpha = stack->pop<GLenum>();
    GLenum src_factor_alpha = stack->pop<GLenum>();
    GLenum dst_factor_rgb = stack->pop<GLenum>();
    GLenum src_factor_rgb = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBlendFuncSeparate(%u, %u, %u, %u)\n", src_factor_rgb, dst_factor_rgb,
                   src_factor_alpha, dst_factor_alpha);
        if (glBlendFuncSeparate != nullptr) {
            glBlendFuncSeparate(src_factor_rgb, dst_factor_rgb, src_factor_alpha, dst_factor_alpha);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBlendFuncSeparate returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlendFuncSeparate\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlendFuncSeparate\n");
        return false;
    }
}

bool callGlDepthFunc(Stack* stack, bool pushReturn) {
    GLenum function = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDepthFunc(%u)\n", function);
        if (glDepthFunc != nullptr) {
            glDepthFunc(function);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDepthFunc returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDepthFunc\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDepthFunc\n");
        return false;
    }
}

bool callGlSampleCoverage(Stack* stack, bool pushReturn) {
    uint8_t invert = stack->pop<uint8_t>();
    float value = stack->pop<float>();
    if (stack->isValid()) {
        GAPID_INFO("glSampleCoverage(%f, %" PRIu8 ")\n", value, invert);
        if (glSampleCoverage != nullptr) {
            glSampleCoverage(value, invert);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glSampleCoverage returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glSampleCoverage\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glSampleCoverage\n");
        return false;
    }
}

bool callGlSampleMaski(Stack* stack, bool pushReturn) {
    GLbitfield mask = stack->pop<GLbitfield>();
    uint32_t maskNumber = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glSampleMaski(%" PRIu32 ", %u)\n", maskNumber, mask);
        if (glSampleMaski != nullptr) {
            glSampleMaski(maskNumber, mask);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glSampleMaski returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glSampleMaski\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glSampleMaski\n");
        return false;
    }
}

bool callGlScissor(Stack* stack, bool pushReturn) {
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    int32_t y = stack->pop<int32_t>();
    int32_t x = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glScissor(%" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ")\n", x, y, width,
                   height);
        if (glScissor != nullptr) {
            glScissor(x, y, width, height);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glScissor returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glScissor\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glScissor\n");
        return false;
    }
}

bool callGlStencilFunc(Stack* stack, bool pushReturn) {
    uint32_t mask = stack->pop<uint32_t>();
    int32_t ref = stack->pop<int32_t>();
    GLenum func = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glStencilFunc(%u, %" PRId32 ", %" PRIu32 ")\n", func, ref, mask);
        if (glStencilFunc != nullptr) {
            glStencilFunc(func, ref, mask);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glStencilFunc returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glStencilFunc\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStencilFunc\n");
        return false;
    }
}

bool callGlStencilFuncSeparate(Stack* stack, bool pushReturn) {
    uint32_t mask = stack->pop<uint32_t>();
    int32_t reference_value = stack->pop<int32_t>();
    GLenum function = stack->pop<GLenum>();
    GLenum face = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glStencilFuncSeparate(%u, %u, %" PRId32 ", %" PRIu32 ")\n", face, function,
                   reference_value, mask);
        if (glStencilFuncSeparate != nullptr) {
            glStencilFuncSeparate(face, function, reference_value, mask);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glStencilFuncSeparate returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glStencilFuncSeparate\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStencilFuncSeparate\n");
        return false;
    }
}

bool callGlStencilOp(Stack* stack, bool pushReturn) {
    GLenum zpass = stack->pop<GLenum>();
    GLenum zfail = stack->pop<GLenum>();
    GLenum fail = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glStencilOp(%u, %u, %u)\n", fail, zfail, zpass);
        if (glStencilOp != nullptr) {
            glStencilOp(fail, zfail, zpass);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glStencilOp returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glStencilOp\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStencilOp\n");
        return false;
    }
}

bool callGlStencilOpSeparate(Stack* stack, bool pushReturn) {
    GLenum stencil_pass_depth_pass = stack->pop<GLenum>();
    GLenum stencil_pass_depth_fail = stack->pop<GLenum>();
    GLenum stencil_fail = stack->pop<GLenum>();
    GLenum face = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glStencilOpSeparate(%u, %u, %u, %u)\n", face, stencil_fail,
                   stencil_pass_depth_fail, stencil_pass_depth_pass);
        if (glStencilOpSeparate != nullptr) {
            glStencilOpSeparate(face, stencil_fail, stencil_pass_depth_fail,
                                stencil_pass_depth_pass);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glStencilOpSeparate returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glStencilOpSeparate\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStencilOpSeparate\n");
        return false;
    }
}

bool callGlBindFramebuffer(Stack* stack, bool pushReturn) {
    uint32_t framebuffer = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBindFramebuffer(%u, %" PRIu32 ")\n", target, framebuffer);
        if (glBindFramebuffer != nullptr) {
            glBindFramebuffer(target, framebuffer);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBindFramebuffer returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindFramebuffer\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindFramebuffer\n");
        return false;
    }
}

bool callGlBindRenderbuffer(Stack* stack, bool pushReturn) {
    uint32_t renderbuffer = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBindRenderbuffer(%u, %" PRIu32 ")\n", target, renderbuffer);
        if (glBindRenderbuffer != nullptr) {
            glBindRenderbuffer(target, renderbuffer);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBindRenderbuffer returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindRenderbuffer\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindRenderbuffer\n");
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
                   ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %u, %u)\n",
                   srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter);
        if (glBlitFramebuffer != nullptr) {
            glBlitFramebuffer(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBlitFramebuffer returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBlitFramebuffer\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBlitFramebuffer\n");
        return false;
    }
}

bool callGlCheckFramebufferStatus(Stack* stack, bool pushReturn) {
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glCheckFramebufferStatus(%u)\n", target);
        if (glCheckFramebufferStatus != nullptr) {
            GLenum return_value = glCheckFramebufferStatus(target);
            GAPID_INFO("Returned: %u\n", return_value);
            if (pushReturn) {
                stack->push<GLenum>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glCheckFramebufferStatus returned error: 0x%x\n", err);
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

bool callGlClear(Stack* stack, bool pushReturn) {
    GLbitfield mask = stack->pop<GLbitfield>();
    if (stack->isValid()) {
        GAPID_INFO("glClear(%u)\n", mask);
        if (glClear != nullptr) {
            glClear(mask);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glClear returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glClear\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glClear\n");
        return false;
    }
}

bool callGlClearBufferfi(Stack* stack, bool pushReturn) {
    int32_t stencil = stack->pop<int32_t>();
    float depth = stack->pop<float>();
    int32_t drawbuffer = stack->pop<int32_t>();
    GLenum buffer = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glClearBufferfi(%u, %" PRId32 ", %f, %" PRId32 ")\n", buffer, drawbuffer, depth,
                   stencil);
        if (glClearBufferfi != nullptr) {
            glClearBufferfi(buffer, drawbuffer, depth, stencil);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glClearBufferfi returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glClearBufferfi\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glClearBufferfi\n");
        return false;
    }
}

bool callGlClearBufferfv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    int32_t drawbuffer = stack->pop<int32_t>();
    GLenum buffer = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glClearBufferfv(%u, %" PRId32 ", %p)\n", buffer, drawbuffer, value);
        if (glClearBufferfv != nullptr) {
            glClearBufferfv(buffer, drawbuffer, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glClearBufferfv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glClearBufferfv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glClearBufferfv\n");
        return false;
    }
}

bool callGlClearBufferiv(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    int32_t drawbuffer = stack->pop<int32_t>();
    GLenum buffer = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glClearBufferiv(%u, %" PRId32 ", %p)\n", buffer, drawbuffer, value);
        if (glClearBufferiv != nullptr) {
            glClearBufferiv(buffer, drawbuffer, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glClearBufferiv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glClearBufferiv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glClearBufferiv\n");
        return false;
    }
}

bool callGlClearBufferuiv(Stack* stack, bool pushReturn) {
    uint32_t* value = stack->pop<uint32_t*>();
    int32_t drawbuffer = stack->pop<int32_t>();
    GLenum buffer = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glClearBufferuiv(%u, %" PRId32 ", %p)\n", buffer, drawbuffer, value);
        if (glClearBufferuiv != nullptr) {
            glClearBufferuiv(buffer, drawbuffer, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glClearBufferuiv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glClearBufferuiv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glClearBufferuiv\n");
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
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glClearColor returned error: 0x%x\n", err);
            }
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
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glClearDepthf returned error: 0x%x\n", err);
            }
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
        GAPID_INFO("glClearStencil(%" PRId32 ")\n", stencil);
        if (glClearStencil != nullptr) {
            glClearStencil(stencil);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glClearStencil returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glClearStencil\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glClearStencil\n");
        return false;
    }
}

bool callGlColorMask(Stack* stack, bool pushReturn) {
    uint8_t alpha = stack->pop<uint8_t>();
    uint8_t blue = stack->pop<uint8_t>();
    uint8_t green = stack->pop<uint8_t>();
    uint8_t red = stack->pop<uint8_t>();
    if (stack->isValid()) {
        GAPID_INFO("glColorMask(%" PRIu8 ", %" PRIu8 ", %" PRIu8 ", %" PRIu8 ")\n", red, green,
                   blue, alpha);
        if (glColorMask != nullptr) {
            glColorMask(red, green, blue, alpha);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glColorMask returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glColorMask\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glColorMask\n");
        return false;
    }
}

bool callGlDeleteFramebuffers(Stack* stack, bool pushReturn) {
    uint32_t* framebuffers = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteFramebuffers(%" PRId32 ", %p)\n", count, framebuffers);
        if (glDeleteFramebuffers != nullptr) {
            glDeleteFramebuffers(count, framebuffers);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDeleteFramebuffers returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteFramebuffers\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteFramebuffers\n");
        return false;
    }
}

bool callGlDeleteRenderbuffers(Stack* stack, bool pushReturn) {
    uint32_t* renderbuffers = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteRenderbuffers(%" PRId32 ", %p)\n", count, renderbuffers);
        if (glDeleteRenderbuffers != nullptr) {
            glDeleteRenderbuffers(count, renderbuffers);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDeleteRenderbuffers returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteRenderbuffers\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteRenderbuffers\n");
        return false;
    }
}

bool callGlDepthMask(Stack* stack, bool pushReturn) {
    uint8_t enabled = stack->pop<uint8_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDepthMask(%" PRIu8 ")\n", enabled);
        if (glDepthMask != nullptr) {
            glDepthMask(enabled);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDepthMask returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDepthMask\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDepthMask\n");
        return false;
    }
}

bool callGlFramebufferParameteri(Stack* stack, bool pushReturn) {
    int32_t param = stack->pop<int32_t>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glFramebufferParameteri(%u, %u, %" PRId32 ")\n", target, pname, param);
        if (glFramebufferParameteri != nullptr) {
            glFramebufferParameteri(target, pname, param);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glFramebufferParameteri returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFramebufferParameteri\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFramebufferParameteri\n");
        return false;
    }
}

bool callGlFramebufferRenderbuffer(Stack* stack, bool pushReturn) {
    uint32_t renderbuffer = stack->pop<uint32_t>();
    GLenum renderbuffer_target = stack->pop<GLenum>();
    GLenum framebuffer_attachment = stack->pop<GLenum>();
    GLenum framebuffer_target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glFramebufferRenderbuffer(%u, %u, %u, %" PRIu32 ")\n", framebuffer_target,
                   framebuffer_attachment, renderbuffer_target, renderbuffer);
        if (glFramebufferRenderbuffer != nullptr) {
            glFramebufferRenderbuffer(framebuffer_target, framebuffer_attachment,
                                      renderbuffer_target, renderbuffer);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glFramebufferRenderbuffer returned error: 0x%x\n", err);
            }
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
    uint32_t texture = stack->pop<uint32_t>();
    GLenum texture_target = stack->pop<GLenum>();
    GLenum framebuffer_attachment = stack->pop<GLenum>();
    GLenum framebuffer_target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glFramebufferTexture2D(%u, %u, %u, %" PRIu32 ", %" PRId32 ")\n",
                   framebuffer_target, framebuffer_attachment, texture_target, texture, level);
        if (glFramebufferTexture2D != nullptr) {
            glFramebufferTexture2D(framebuffer_target, framebuffer_attachment, texture_target,
                                   texture, level);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glFramebufferTexture2D returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFramebufferTexture2D\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFramebufferTexture2D\n");
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
        GAPID_INFO("glFramebufferTextureLayer(%u, %u, %" PRIu32 ", %" PRId32 ", %" PRId32 ")\n",
                   target, attachment, texture, level, layer);
        if (glFramebufferTextureLayer != nullptr) {
            glFramebufferTextureLayer(target, attachment, texture, level, layer);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glFramebufferTextureLayer returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFramebufferTextureLayer\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFramebufferTextureLayer\n");
        return false;
    }
}

bool callGlGenFramebuffers(Stack* stack, bool pushReturn) {
    uint32_t* framebuffers = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenFramebuffers(%" PRId32 ", %p)\n", count, framebuffers);
        if (glGenFramebuffers != nullptr) {
            glGenFramebuffers(count, framebuffers);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGenFramebuffers returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenFramebuffers\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenFramebuffers\n");
        return false;
    }
}

bool callGlGenRenderbuffers(Stack* stack, bool pushReturn) {
    uint32_t* renderbuffers = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenRenderbuffers(%" PRId32 ", %p)\n", count, renderbuffers);
        if (glGenRenderbuffers != nullptr) {
            glGenRenderbuffers(count, renderbuffers);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGenRenderbuffers returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenRenderbuffers\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenRenderbuffers\n");
        return false;
    }
}

bool callGlGetFramebufferAttachmentParameteriv(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    GLenum parameter = stack->pop<GLenum>();
    GLenum attachment = stack->pop<GLenum>();
    GLenum framebuffer_target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetFramebufferAttachmentParameteriv(%u, %u, %u, %p)\n", framebuffer_target,
                   attachment, parameter, value);
        if (glGetFramebufferAttachmentParameteriv != nullptr) {
            glGetFramebufferAttachmentParameteriv(framebuffer_target, attachment, parameter, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetFramebufferAttachmentParameteriv returned error: 0x%x\n", err);
            }
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

bool callGlGetFramebufferParameteriv(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetFramebufferParameteriv(%u, %u, %p)\n", target, pname, params);
        if (glGetFramebufferParameteriv != nullptr) {
            glGetFramebufferParameteriv(target, pname, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetFramebufferParameteriv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetFramebufferParameteriv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetFramebufferParameteriv\n");
        return false;
    }
}

bool callGlGetRenderbufferParameteriv(Stack* stack, bool pushReturn) {
    int32_t* values = stack->pop<int32_t*>();
    GLenum parameter = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetRenderbufferParameteriv(%u, %u, %p)\n", target, parameter, values);
        if (glGetRenderbufferParameteriv != nullptr) {
            glGetRenderbufferParameteriv(target, parameter, values);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetRenderbufferParameteriv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetRenderbufferParameteriv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetRenderbufferParameteriv\n");
        return false;
    }
}

bool callGlInvalidateFramebuffer(Stack* stack, bool pushReturn) {
    GLenum* attachments = stack->pop<GLenum*>();
    int32_t count = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glInvalidateFramebuffer(%u, %" PRId32 ", %p)\n", target, count, attachments);
        if (glInvalidateFramebuffer != nullptr) {
            glInvalidateFramebuffer(target, count, attachments);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glInvalidateFramebuffer returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glInvalidateFramebuffer\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glInvalidateFramebuffer\n");
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
                   ", %" PRId32 ", %" PRId32 ")\n",
                   target, numAttachments, attachments, x, y, width, height);
        if (glInvalidateSubFramebuffer != nullptr) {
            glInvalidateSubFramebuffer(target, numAttachments, attachments, x, y, width, height);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glInvalidateSubFramebuffer returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glInvalidateSubFramebuffer\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glInvalidateSubFramebuffer\n");
        return false;
    }
}

bool callGlIsFramebuffer(Stack* stack, bool pushReturn) {
    uint32_t framebuffer = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsFramebuffer(%" PRIu32 ")\n", framebuffer);
        if (glIsFramebuffer != nullptr) {
            uint8_t return_value = glIsFramebuffer(framebuffer);
            GAPID_INFO("Returned: %" PRIu8 "\n", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glIsFramebuffer returned error: 0x%x\n", err);
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

bool callGlIsRenderbuffer(Stack* stack, bool pushReturn) {
    uint32_t renderbuffer = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsRenderbuffer(%" PRIu32 ")\n", renderbuffer);
        if (glIsRenderbuffer != nullptr) {
            uint8_t return_value = glIsRenderbuffer(renderbuffer);
            GAPID_INFO("Returned: %" PRIu8 "\n", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glIsRenderbuffer returned error: 0x%x\n", err);
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

bool callGlReadBuffer(Stack* stack, bool pushReturn) {
    GLenum src = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glReadBuffer(%u)\n", src);
        if (glReadBuffer != nullptr) {
            glReadBuffer(src);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glReadBuffer returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glReadBuffer\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glReadBuffer\n");
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
        GAPID_INFO("glReadPixels(%" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %u, %u, %p)\n",
                   x, y, width, height, format, type, data);
        if (glReadPixels != nullptr) {
            glReadPixels(x, y, width, height, format, type, data);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glReadPixels returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glReadPixels\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glReadPixels\n");
        return false;
    }
}

bool callGlRenderbufferStorage(Stack* stack, bool pushReturn) {
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    GLenum format = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glRenderbufferStorage(%u, %u, %" PRId32 ", %" PRId32 ")\n", target, format,
                   width, height);
        if (glRenderbufferStorage != nullptr) {
            glRenderbufferStorage(target, format, width, height);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glRenderbufferStorage returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glRenderbufferStorage\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glRenderbufferStorage\n");
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
                   ")\n",
                   target, samples, format, width, height);
        if (glRenderbufferStorageMultisample != nullptr) {
            glRenderbufferStorageMultisample(target, samples, format, width, height);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glRenderbufferStorageMultisample returned error: 0x%x\n", err);
            }
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

bool callGlStencilMask(Stack* stack, bool pushReturn) {
    uint32_t mask = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glStencilMask(%" PRIu32 ")\n", mask);
        if (glStencilMask != nullptr) {
            glStencilMask(mask);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glStencilMask returned error: 0x%x\n", err);
            }
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
    GLenum face = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glStencilMaskSeparate(%u, %" PRIu32 ")\n", face, mask);
        if (glStencilMaskSeparate != nullptr) {
            glStencilMaskSeparate(face, mask);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glStencilMaskSeparate returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glStencilMaskSeparate\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glStencilMaskSeparate\n");
        return false;
    }
}

bool callGlDisable(Stack* stack, bool pushReturn) {
    GLenum capability = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glDisable(%u)\n", capability);
        if (glDisable != nullptr) {
            glDisable(capability);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDisable returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDisable\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDisable\n");
        return false;
    }
}

bool callGlEnable(Stack* stack, bool pushReturn) {
    GLenum capability = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glEnable(%u)\n", capability);
        if (glEnable != nullptr) {
            glEnable(capability);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glEnable returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEnable\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEnable\n");
        return false;
    }
}

bool callGlFinish(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glFinish()\n");
        if (glFinish != nullptr) {
            glFinish();
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glFinish returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFinish\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFinish\n");
        return false;
    }
}

bool callGlFlush(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glFlush()\n");
        if (glFlush != nullptr) {
            glFlush();
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glFlush returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFlush\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFlush\n");
        return false;
    }
}

bool callGlFlushMappedBufferRange(Stack* stack, bool pushReturn) {
    int32_t length = stack->pop<int32_t>();
    int32_t offset = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glFlushMappedBufferRange(%u, %" PRId32 ", %" PRId32 ")\n", target, offset,
                   length);
        if (glFlushMappedBufferRange != nullptr) {
            glFlushMappedBufferRange(target, offset, length);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glFlushMappedBufferRange returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFlushMappedBufferRange\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFlushMappedBufferRange\n");
        return false;
    }
}

bool callGlGetError(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glGetError()\n");
        if (glGetError != nullptr) {
            GLenum return_value = glGetError();
            GAPID_INFO("Returned: %u\n", return_value);
            if (pushReturn) {
                stack->push<GLenum>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetError returned error: 0x%x\n", err);
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

bool callGlHint(Stack* stack, bool pushReturn) {
    GLenum mode = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glHint(%u, %u)\n", target, mode);
        if (glHint != nullptr) {
            glHint(target, mode);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glHint returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glHint\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glHint\n");
        return false;
    }
}

bool callGlActiveShaderProgram(Stack* stack, bool pushReturn) {
    uint32_t program = stack->pop<uint32_t>();
    uint32_t pipeline = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glActiveShaderProgram(%" PRIu32 ", %" PRIu32 ")\n", pipeline, program);
        if (glActiveShaderProgram != nullptr) {
            glActiveShaderProgram(pipeline, program);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glActiveShaderProgram returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glActiveShaderProgram\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glActiveShaderProgram\n");
        return false;
    }
}

bool callGlAttachShader(Stack* stack, bool pushReturn) {
    uint32_t shader = stack->pop<uint32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glAttachShader(%" PRIu32 ", %" PRIu32 ")\n", program, shader);
        if (glAttachShader != nullptr) {
            glAttachShader(program, shader);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glAttachShader returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glAttachShader\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glAttachShader\n");
        return false;
    }
}

bool callGlBindAttribLocation(Stack* stack, bool pushReturn) {
    char* name = stack->pop<char*>();
    uint32_t location = stack->pop<uint32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBindAttribLocation(%" PRIu32 ", %" PRIu32 ", %s)\n", program, location, name);
        if (glBindAttribLocation != nullptr) {
            glBindAttribLocation(program, location, name);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBindAttribLocation returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindAttribLocation\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindAttribLocation\n");
        return false;
    }
}

bool callGlBindProgramPipeline(Stack* stack, bool pushReturn) {
    uint32_t pipeline = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBindProgramPipeline(%" PRIu32 ")\n", pipeline);
        if (glBindProgramPipeline != nullptr) {
            glBindProgramPipeline(pipeline);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBindProgramPipeline returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindProgramPipeline\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindProgramPipeline\n");
        return false;
    }
}

bool callGlCompileShader(Stack* stack, bool pushReturn) {
    uint32_t shader = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glCompileShader(%" PRIu32 ")\n", shader);
        if (glCompileShader != nullptr) {
            glCompileShader(shader);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glCompileShader returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCompileShader\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCompileShader\n");
        return false;
    }
}

bool callGlCreateProgram(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glCreateProgram()\n");
        if (glCreateProgram != nullptr) {
            uint32_t return_value = glCreateProgram();
            GAPID_INFO("Returned: %" PRIu32 "\n", return_value);
            if (pushReturn) {
                stack->push<uint32_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glCreateProgram returned error: 0x%x\n", err);
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

bool callGlCreateShader(Stack* stack, bool pushReturn) {
    GLenum type = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glCreateShader(%u)\n", type);
        if (glCreateShader != nullptr) {
            uint32_t return_value = glCreateShader(type);
            GAPID_INFO("Returned: %" PRIu32 "\n", return_value);
            if (pushReturn) {
                stack->push<uint32_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glCreateShader returned error: 0x%x\n", err);
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

bool callGlCreateShaderProgramv(Stack* stack, bool pushReturn) {
    char** strings = stack->pop<char**>();
    int32_t count = stack->pop<int32_t>();
    GLenum type = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glCreateShaderProgramv(%u, %" PRId32 ", %p)\n", type, count, strings);
        if (glCreateShaderProgramv != nullptr) {
            uint32_t return_value = glCreateShaderProgramv(type, count, strings);
            GAPID_INFO("Returned: %" PRIu32 "\n", return_value);
            if (pushReturn) {
                stack->push<uint32_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glCreateShaderProgramv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCreateShaderProgramv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCreateShaderProgramv\n");
        return false;
    }
}

bool callGlDeleteProgram(Stack* stack, bool pushReturn) {
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteProgram(%" PRIu32 ")\n", program);
        if (glDeleteProgram != nullptr) {
            glDeleteProgram(program);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDeleteProgram returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteProgram\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteProgram\n");
        return false;
    }
}

bool callGlDeleteProgramPipelines(Stack* stack, bool pushReturn) {
    uint32_t* pipelines = stack->pop<uint32_t*>();
    int32_t n = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteProgramPipelines(%" PRId32 ", %p)\n", n, pipelines);
        if (glDeleteProgramPipelines != nullptr) {
            glDeleteProgramPipelines(n, pipelines);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDeleteProgramPipelines returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteProgramPipelines\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteProgramPipelines\n");
        return false;
    }
}

bool callGlDeleteShader(Stack* stack, bool pushReturn) {
    uint32_t shader = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteShader(%" PRIu32 ")\n", shader);
        if (glDeleteShader != nullptr) {
            glDeleteShader(shader);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDeleteShader returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteShader\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteShader\n");
        return false;
    }
}

bool callGlDetachShader(Stack* stack, bool pushReturn) {
    uint32_t shader = stack->pop<uint32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDetachShader(%" PRIu32 ", %" PRIu32 ")\n", program, shader);
        if (glDetachShader != nullptr) {
            glDetachShader(program, shader);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDetachShader returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDetachShader\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDetachShader\n");
        return false;
    }
}

bool callGlDispatchCompute(Stack* stack, bool pushReturn) {
    uint32_t num_groups_z = stack->pop<uint32_t>();
    uint32_t num_groups_y = stack->pop<uint32_t>();
    uint32_t num_groups_x = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDispatchCompute(%" PRIu32 ", %" PRIu32 ", %" PRIu32 ")\n", num_groups_x,
                   num_groups_y, num_groups_z);
        if (glDispatchCompute != nullptr) {
            glDispatchCompute(num_groups_x, num_groups_y, num_groups_z);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDispatchCompute returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDispatchCompute\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDispatchCompute\n");
        return false;
    }
}

bool callGlDispatchComputeIndirect(Stack* stack, bool pushReturn) {
    int32_t indirect = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDispatchComputeIndirect(%" PRId32 ")\n", indirect);
        if (glDispatchComputeIndirect != nullptr) {
            glDispatchComputeIndirect(indirect);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDispatchComputeIndirect returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDispatchComputeIndirect\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDispatchComputeIndirect\n");
        return false;
    }
}

bool callGlGenProgramPipelines(Stack* stack, bool pushReturn) {
    uint32_t* pipelines = stack->pop<uint32_t*>();
    int32_t n = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenProgramPipelines(%" PRId32 ", %p)\n", n, pipelines);
        if (glGenProgramPipelines != nullptr) {
            glGenProgramPipelines(n, pipelines);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGenProgramPipelines returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenProgramPipelines\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenProgramPipelines\n");
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
        GAPID_INFO("glGetActiveAttrib(%" PRIu32 ", %" PRIu32 ", %" PRId32 ", %p, %p, %p, %p)\n",
                   program, location, buffer_size, buffer_bytes_written, vector_count, type, name);
        if (glGetActiveAttrib != nullptr) {
            glGetActiveAttrib(program, location, buffer_size, buffer_bytes_written, vector_count,
                              type, name);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetActiveAttrib returned error: 0x%x\n", err);
            }
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
    GLenum* type = stack->pop<GLenum*>();
    int32_t* vector_count = stack->pop<int32_t*>();
    int32_t* buffer_bytes_written = stack->pop<int32_t*>();
    int32_t buffer_size = stack->pop<int32_t>();
    uint32_t location = stack->pop<uint32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetActiveUniform(%" PRIu32 ", %" PRIu32 ", %" PRId32 ", %p, %p, %p, %p)\n",
                   program, location, buffer_size, buffer_bytes_written, vector_count, type, name);
        if (glGetActiveUniform != nullptr) {
            glGetActiveUniform(program, location, buffer_size, buffer_bytes_written, vector_count,
                               type, name);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetActiveUniform returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetActiveUniform\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetActiveUniform\n");
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
        GAPID_INFO("glGetActiveUniformBlockName(%" PRIu32 ", %" PRIu32 ", %" PRId32 ", %p, %p)\n",
                   program, uniform_block_index, buffer_size, buffer_bytes_written, name);
        if (glGetActiveUniformBlockName != nullptr) {
            glGetActiveUniformBlockName(program, uniform_block_index, buffer_size,
                                        buffer_bytes_written, name);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetActiveUniformBlockName returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetActiveUniformBlockName\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetActiveUniformBlockName\n");
        return false;
    }
}

bool callGlGetActiveUniformBlockiv(Stack* stack, bool pushReturn) {
    int32_t* parameters = stack->pop<int32_t*>();
    GLenum parameter_name = stack->pop<GLenum>();
    uint32_t uniform_block_index = stack->pop<uint32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetActiveUniformBlockiv(%" PRIu32 ", %" PRIu32 ", %u, %p)\n", program,
                   uniform_block_index, parameter_name, parameters);
        if (glGetActiveUniformBlockiv != nullptr) {
            glGetActiveUniformBlockiv(program, uniform_block_index, parameter_name, parameters);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetActiveUniformBlockiv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetActiveUniformBlockiv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetActiveUniformBlockiv\n");
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
        GAPID_INFO("glGetActiveUniformsiv(%" PRIu32 ", %" PRId32 ", %p, %u, %p)\n", program,
                   uniform_count, uniform_indices, parameter_name, parameters);
        if (glGetActiveUniformsiv != nullptr) {
            glGetActiveUniformsiv(program, uniform_count, uniform_indices, parameter_name,
                                  parameters);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetActiveUniformsiv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetActiveUniformsiv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetActiveUniformsiv\n");
        return false;
    }
}

bool callGlGetAttachedShaders(Stack* stack, bool pushReturn) {
    uint32_t* shaders = stack->pop<uint32_t*>();
    int32_t* shaders_length_written = stack->pop<int32_t*>();
    int32_t buffer_length = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetAttachedShaders(%" PRIu32 ", %" PRId32 ", %p, %p)\n", program,
                   buffer_length, shaders_length_written, shaders);
        if (glGetAttachedShaders != nullptr) {
            glGetAttachedShaders(program, buffer_length, shaders_length_written, shaders);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetAttachedShaders returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetAttachedShaders\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetAttachedShaders\n");
        return false;
    }
}

bool callGlGetAttribLocation(Stack* stack, bool pushReturn) {
    char* name = stack->pop<char*>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetAttribLocation(%" PRIu32 ", %s)\n", program, name);
        if (glGetAttribLocation != nullptr) {
            int32_t return_value = glGetAttribLocation(program, name);
            GAPID_INFO("Returned: %" PRId32 "\n", return_value);
            if (pushReturn) {
                stack->push<int32_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetAttribLocation returned error: 0x%x\n", err);
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

bool callGlGetFragDataLocation(Stack* stack, bool pushReturn) {
    char* name = stack->pop<char*>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetFragDataLocation(%" PRIu32 ", %p)\n", program, name);
        if (glGetFragDataLocation != nullptr) {
            int32_t return_value = glGetFragDataLocation(program, name);
            GAPID_INFO("Returned: %" PRId32 "\n", return_value);
            if (pushReturn) {
                stack->push<int32_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetFragDataLocation returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetFragDataLocation\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetFragDataLocation\n");
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
        GAPID_INFO("glGetProgramBinary(%" PRIu32 ", %" PRId32 ", %p, %p, %p)\n", program, bufSize,
                   length, binaryFormat, binary);
        if (glGetProgramBinary != nullptr) {
            glGetProgramBinary(program, bufSize, length, binaryFormat, binary);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetProgramBinary returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetProgramBinary\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetProgramBinary\n");
        return false;
    }
}

bool callGlGetProgramInfoLog(Stack* stack, bool pushReturn) {
    char* info = stack->pop<char*>();
    int32_t* string_length_written = stack->pop<int32_t*>();
    int32_t buffer_length = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetProgramInfoLog(%" PRIu32 ", %" PRId32 ", %p, %p)\n", program,
                   buffer_length, string_length_written, info);
        if (glGetProgramInfoLog != nullptr) {
            glGetProgramInfoLog(program, buffer_length, string_length_written, info);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetProgramInfoLog returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetProgramInfoLog\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetProgramInfoLog\n");
        return false;
    }
}

bool callGlGetProgramInterfaceiv(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    GLenum programInterface = stack->pop<GLenum>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetProgramInterfaceiv(%" PRIu32 ", %u, %u, %p)\n", program, programInterface,
                   pname, params);
        if (glGetProgramInterfaceiv != nullptr) {
            glGetProgramInterfaceiv(program, programInterface, pname, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetProgramInterfaceiv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetProgramInterfaceiv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetProgramInterfaceiv\n");
        return false;
    }
}

bool callGlGetProgramPipelineInfoLog(Stack* stack, bool pushReturn) {
    char* infoLog = stack->pop<char*>();
    int32_t* length = stack->pop<int32_t*>();
    int32_t bufSize = stack->pop<int32_t>();
    uint32_t pipeline = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetProgramPipelineInfoLog(%" PRIu32 ", %" PRId32 ", %p, %p)\n", pipeline,
                   bufSize, length, infoLog);
        if (glGetProgramPipelineInfoLog != nullptr) {
            glGetProgramPipelineInfoLog(pipeline, bufSize, length, infoLog);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetProgramPipelineInfoLog returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetProgramPipelineInfoLog\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetProgramPipelineInfoLog\n");
        return false;
    }
}

bool callGlGetProgramPipelineiv(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t pipeline = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetProgramPipelineiv(%" PRIu32 ", %u, %p)\n", pipeline, pname, params);
        if (glGetProgramPipelineiv != nullptr) {
            glGetProgramPipelineiv(pipeline, pname, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetProgramPipelineiv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetProgramPipelineiv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetProgramPipelineiv\n");
        return false;
    }
}

bool callGlGetProgramResourceIndex(Stack* stack, bool pushReturn) {
    char* name = stack->pop<char*>();
    GLenum programInterface = stack->pop<GLenum>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetProgramResourceIndex(%" PRIu32 ", %u, %p)\n", program, programInterface,
                   name);
        if (glGetProgramResourceIndex != nullptr) {
            uint32_t return_value = glGetProgramResourceIndex(program, programInterface, name);
            GAPID_INFO("Returned: %" PRIu32 "\n", return_value);
            if (pushReturn) {
                stack->push<uint32_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetProgramResourceIndex returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetProgramResourceIndex\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetProgramResourceIndex\n");
        return false;
    }
}

bool callGlGetProgramResourceLocation(Stack* stack, bool pushReturn) {
    char* name = stack->pop<char*>();
    GLenum programInterface = stack->pop<GLenum>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetProgramResourceLocation(%" PRIu32 ", %u, %p)\n", program, programInterface,
                   name);
        if (glGetProgramResourceLocation != nullptr) {
            int32_t return_value = glGetProgramResourceLocation(program, programInterface, name);
            GAPID_INFO("Returned: %" PRId32 "\n", return_value);
            if (pushReturn) {
                stack->push<int32_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetProgramResourceLocation returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetProgramResourceLocation\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetProgramResourceLocation\n");
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
        GAPID_INFO("glGetProgramResourceName(%" PRIu32 ", %u, %" PRIu32 ", %" PRId32 ", %p, %p)\n",
                   program, programInterface, index, bufSize, length, name);
        if (glGetProgramResourceName != nullptr) {
            glGetProgramResourceName(program, programInterface, index, bufSize, length, name);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetProgramResourceName returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetProgramResourceName\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetProgramResourceName\n");
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
                   ", %p, %p)\n",
                   program, programInterface, index, propCount, props, bufSize, length, params);
        if (glGetProgramResourceiv != nullptr) {
            glGetProgramResourceiv(program, programInterface, index, propCount, props, bufSize,
                                   length, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetProgramResourceiv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetProgramResourceiv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetProgramResourceiv\n");
        return false;
    }
}

bool callGlGetProgramiv(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    GLenum parameter = stack->pop<GLenum>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetProgramiv(%" PRIu32 ", %u, %p)\n", program, parameter, value);
        if (glGetProgramiv != nullptr) {
            glGetProgramiv(program, parameter, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetProgramiv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetProgramiv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetProgramiv\n");
        return false;
    }
}

bool callGlGetShaderInfoLog(Stack* stack, bool pushReturn) {
    char* info = stack->pop<char*>();
    int32_t* string_length_written = stack->pop<int32_t*>();
    int32_t buffer_length = stack->pop<int32_t>();
    uint32_t shader = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetShaderInfoLog(%" PRIu32 ", %" PRId32 ", %p, %p)\n", shader, buffer_length,
                   string_length_written, info);
        if (glGetShaderInfoLog != nullptr) {
            glGetShaderInfoLog(shader, buffer_length, string_length_written, info);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetShaderInfoLog returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetShaderInfoLog\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetShaderInfoLog\n");
        return false;
    }
}

bool callGlGetShaderPrecisionFormat(Stack* stack, bool pushReturn) {
    int32_t* precision = stack->pop<int32_t*>();
    int32_t* range = stack->pop<int32_t*>();
    GLenum precision_type = stack->pop<GLenum>();
    GLenum shader_type = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetShaderPrecisionFormat(%u, %u, %p, %p)\n", shader_type, precision_type,
                   range, precision);
        if (glGetShaderPrecisionFormat != nullptr) {
            glGetShaderPrecisionFormat(shader_type, precision_type, range, precision);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetShaderPrecisionFormat returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetShaderPrecisionFormat\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetShaderPrecisionFormat\n");
        return false;
    }
}

bool callGlGetShaderSource(Stack* stack, bool pushReturn) {
    char* source = stack->pop<char*>();
    int32_t* string_length_written = stack->pop<int32_t*>();
    int32_t buffer_length = stack->pop<int32_t>();
    uint32_t shader = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetShaderSource(%" PRIu32 ", %" PRId32 ", %p, %p)\n", shader, buffer_length,
                   string_length_written, source);
        if (glGetShaderSource != nullptr) {
            glGetShaderSource(shader, buffer_length, string_length_written, source);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetShaderSource returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetShaderSource\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetShaderSource\n");
        return false;
    }
}

bool callGlGetShaderiv(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    GLenum parameter = stack->pop<GLenum>();
    uint32_t shader = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetShaderiv(%" PRIu32 ", %u, %p)\n", shader, parameter, value);
        if (glGetShaderiv != nullptr) {
            glGetShaderiv(shader, parameter, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetShaderiv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetShaderiv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetShaderiv\n");
        return false;
    }
}

bool callGlGetUniformBlockIndex(Stack* stack, bool pushReturn) {
    char* uniformBlockName = stack->pop<char*>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetUniformBlockIndex(%" PRIu32 ", %p)\n", program, uniformBlockName);
        if (glGetUniformBlockIndex != nullptr) {
            uint32_t return_value = glGetUniformBlockIndex(program, uniformBlockName);
            GAPID_INFO("Returned: %" PRIu32 "\n", return_value);
            if (pushReturn) {
                stack->push<uint32_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetUniformBlockIndex returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetUniformBlockIndex\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetUniformBlockIndex\n");
        return false;
    }
}

bool callGlGetUniformIndices(Stack* stack, bool pushReturn) {
    uint32_t* uniformIndices = stack->pop<uint32_t*>();
    char** uniformNames = stack->pop<char**>();
    int32_t uniformCount = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetUniformIndices(%" PRIu32 ", %" PRId32 ", %p, %p)\n", program, uniformCount,
                   uniformNames, uniformIndices);
        if (glGetUniformIndices != nullptr) {
            glGetUniformIndices(program, uniformCount, uniformNames, uniformIndices);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetUniformIndices returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetUniformIndices\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetUniformIndices\n");
        return false;
    }
}

bool callGlGetUniformLocation(Stack* stack, bool pushReturn) {
    char* name = stack->pop<char*>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetUniformLocation(%" PRIu32 ", %s)\n", program, name);
        if (glGetUniformLocation != nullptr) {
            int32_t return_value = glGetUniformLocation(program, name);
            GAPID_INFO("Returned: %" PRId32 "\n", return_value);
            if (pushReturn) {
                stack->push<int32_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetUniformLocation returned error: 0x%x\n", err);
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

bool callGlGetUniformfv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetUniformfv(%" PRIu32 ", %" PRId32 ", %p)\n", program, location, values);
        if (glGetUniformfv != nullptr) {
            glGetUniformfv(program, location, values);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetUniformfv returned error: 0x%x\n", err);
            }
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
    int32_t* values = stack->pop<int32_t*>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetUniformiv(%" PRIu32 ", %" PRId32 ", %p)\n", program, location, values);
        if (glGetUniformiv != nullptr) {
            glGetUniformiv(program, location, values);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetUniformiv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetUniformiv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetUniformiv\n");
        return false;
    }
}

bool callGlGetUniformuiv(Stack* stack, bool pushReturn) {
    uint32_t* params = stack->pop<uint32_t*>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetUniformuiv(%" PRIu32 ", %" PRId32 ", %p)\n", program, location, params);
        if (glGetUniformuiv != nullptr) {
            glGetUniformuiv(program, location, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetUniformuiv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetUniformuiv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetUniformuiv\n");
        return false;
    }
}

bool callGlIsProgram(Stack* stack, bool pushReturn) {
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsProgram(%" PRIu32 ")\n", program);
        if (glIsProgram != nullptr) {
            uint8_t return_value = glIsProgram(program);
            GAPID_INFO("Returned: %" PRIu8 "\n", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glIsProgram returned error: 0x%x\n", err);
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

bool callGlIsProgramPipeline(Stack* stack, bool pushReturn) {
    uint32_t pipeline = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsProgramPipeline(%" PRIu32 ")\n", pipeline);
        if (glIsProgramPipeline != nullptr) {
            uint8_t return_value = glIsProgramPipeline(pipeline);
            GAPID_INFO("Returned: %" PRIu8 "\n", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glIsProgramPipeline returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsProgramPipeline\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsProgramPipeline\n");
        return false;
    }
}

bool callGlIsShader(Stack* stack, bool pushReturn) {
    uint32_t shader = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsShader(%" PRIu32 ")\n", shader);
        if (glIsShader != nullptr) {
            uint8_t return_value = glIsShader(shader);
            GAPID_INFO("Returned: %" PRIu8 "\n", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glIsShader returned error: 0x%x\n", err);
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

bool callGlLinkProgram(Stack* stack, bool pushReturn) {
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glLinkProgram(%" PRIu32 ")\n", program);
        if (glLinkProgram != nullptr) {
            glLinkProgram(program);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glLinkProgram returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glLinkProgram\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glLinkProgram\n");
        return false;
    }
}

bool callGlMemoryBarrier(Stack* stack, bool pushReturn) {
    GLbitfield barriers = stack->pop<GLbitfield>();
    if (stack->isValid()) {
        GAPID_INFO("glMemoryBarrier(%u)\n", barriers);
        if (glMemoryBarrier != nullptr) {
            glMemoryBarrier(barriers);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glMemoryBarrier returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMemoryBarrier\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMemoryBarrier\n");
        return false;
    }
}

bool callGlMemoryBarrierByRegion(Stack* stack, bool pushReturn) {
    GLbitfield barriers = stack->pop<GLbitfield>();
    if (stack->isValid()) {
        GAPID_INFO("glMemoryBarrierByRegion(%u)\n", barriers);
        if (glMemoryBarrierByRegion != nullptr) {
            glMemoryBarrierByRegion(barriers);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glMemoryBarrierByRegion returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glMemoryBarrierByRegion\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glMemoryBarrierByRegion\n");
        return false;
    }
}

bool callGlProgramBinary(Stack* stack, bool pushReturn) {
    int32_t length = stack->pop<int32_t>();
    void* binary = stack->pop<void*>();
    GLenum binaryFormat = stack->pop<GLenum>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramBinary(%" PRIu32 ", %u, %p, %" PRId32 ")\n", program, binaryFormat,
                   binary, length);
        if (glProgramBinary != nullptr) {
            glProgramBinary(program, binaryFormat, binary, length);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramBinary returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramBinary\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramBinary\n");
        return false;
    }
}

bool callGlProgramParameteri(Stack* stack, bool pushReturn) {
    int32_t value = stack->pop<int32_t>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramParameteri(%" PRIu32 ", %u, %" PRId32 ")\n", program, pname, value);
        if (glProgramParameteri != nullptr) {
            glProgramParameteri(program, pname, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramParameteri returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramParameteri\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramParameteri\n");
        return false;
    }
}

bool callGlProgramUniform1f(Stack* stack, bool pushReturn) {
    float v0 = stack->pop<float>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform1f(%" PRIu32 ", %" PRId32 ", %f)\n", program, location, v0);
        if (glProgramUniform1f != nullptr) {
            glProgramUniform1f(program, location, v0);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform1f returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform1f\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform1f\n");
        return false;
    }
}

bool callGlProgramUniform1fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform1fv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)\n", program,
                   location, count, value);
        if (glProgramUniform1fv != nullptr) {
            glProgramUniform1fv(program, location, count, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform1fv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform1fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform1fv\n");
        return false;
    }
}

bool callGlProgramUniform1i(Stack* stack, bool pushReturn) {
    int32_t v0 = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform1i(%" PRIu32 ", %" PRId32 ", %" PRId32 ")\n", program, location,
                   v0);
        if (glProgramUniform1i != nullptr) {
            glProgramUniform1i(program, location, v0);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform1i returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform1i\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform1i\n");
        return false;
    }
}

bool callGlProgramUniform1iv(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform1iv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)\n", program,
                   location, count, value);
        if (glProgramUniform1iv != nullptr) {
            glProgramUniform1iv(program, location, count, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform1iv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform1iv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform1iv\n");
        return false;
    }
}

bool callGlProgramUniform1ui(Stack* stack, bool pushReturn) {
    uint32_t v0 = stack->pop<uint32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform1ui(%" PRIu32 ", %" PRId32 ", %" PRIu32 ")\n", program,
                   location, v0);
        if (glProgramUniform1ui != nullptr) {
            glProgramUniform1ui(program, location, v0);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform1ui returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform1ui\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform1ui\n");
        return false;
    }
}

bool callGlProgramUniform1uiv(Stack* stack, bool pushReturn) {
    uint32_t* value = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform1uiv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)\n", program,
                   location, count, value);
        if (glProgramUniform1uiv != nullptr) {
            glProgramUniform1uiv(program, location, count, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform1uiv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform1uiv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform1uiv\n");
        return false;
    }
}

bool callGlProgramUniform2f(Stack* stack, bool pushReturn) {
    float v1 = stack->pop<float>();
    float v0 = stack->pop<float>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform2f(%" PRIu32 ", %" PRId32 ", %f, %f)\n", program, location, v0,
                   v1);
        if (glProgramUniform2f != nullptr) {
            glProgramUniform2f(program, location, v0, v1);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform2f returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform2f\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform2f\n");
        return false;
    }
}

bool callGlProgramUniform2fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform2fv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)\n", program,
                   location, count, value);
        if (glProgramUniform2fv != nullptr) {
            glProgramUniform2fv(program, location, count, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform2fv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform2fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform2fv\n");
        return false;
    }
}

bool callGlProgramUniform2i(Stack* stack, bool pushReturn) {
    int32_t v1 = stack->pop<int32_t>();
    int32_t v0 = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform2i(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ")\n",
                   program, location, v0, v1);
        if (glProgramUniform2i != nullptr) {
            glProgramUniform2i(program, location, v0, v1);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform2i returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform2i\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform2i\n");
        return false;
    }
}

bool callGlProgramUniform2iv(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform2iv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)\n", program,
                   location, count, value);
        if (glProgramUniform2iv != nullptr) {
            glProgramUniform2iv(program, location, count, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform2iv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform2iv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform2iv\n");
        return false;
    }
}

bool callGlProgramUniform2ui(Stack* stack, bool pushReturn) {
    uint32_t v1 = stack->pop<uint32_t>();
    uint32_t v0 = stack->pop<uint32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform2ui(%" PRIu32 ", %" PRId32 ", %" PRIu32 ", %" PRIu32 ")\n",
                   program, location, v0, v1);
        if (glProgramUniform2ui != nullptr) {
            glProgramUniform2ui(program, location, v0, v1);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform2ui returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform2ui\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform2ui\n");
        return false;
    }
}

bool callGlProgramUniform2uiv(Stack* stack, bool pushReturn) {
    uint32_t* value = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform2uiv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)\n", program,
                   location, count, value);
        if (glProgramUniform2uiv != nullptr) {
            glProgramUniform2uiv(program, location, count, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform2uiv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform2uiv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform2uiv\n");
        return false;
    }
}

bool callGlProgramUniform3f(Stack* stack, bool pushReturn) {
    float v2 = stack->pop<float>();
    float v1 = stack->pop<float>();
    float v0 = stack->pop<float>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform3f(%" PRIu32 ", %" PRId32 ", %f, %f, %f)\n", program, location,
                   v0, v1, v2);
        if (glProgramUniform3f != nullptr) {
            glProgramUniform3f(program, location, v0, v1, v2);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform3f returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform3f\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform3f\n");
        return false;
    }
}

bool callGlProgramUniform3fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform3fv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)\n", program,
                   location, count, value);
        if (glProgramUniform3fv != nullptr) {
            glProgramUniform3fv(program, location, count, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform3fv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform3fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform3fv\n");
        return false;
    }
}

bool callGlProgramUniform3i(Stack* stack, bool pushReturn) {
    int32_t v2 = stack->pop<int32_t>();
    int32_t v1 = stack->pop<int32_t>();
    int32_t v0 = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform3i(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32
                   ")\n",
                   program, location, v0, v1, v2);
        if (glProgramUniform3i != nullptr) {
            glProgramUniform3i(program, location, v0, v1, v2);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform3i returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform3i\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform3i\n");
        return false;
    }
}

bool callGlProgramUniform3iv(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform3iv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)\n", program,
                   location, count, value);
        if (glProgramUniform3iv != nullptr) {
            glProgramUniform3iv(program, location, count, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform3iv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform3iv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform3iv\n");
        return false;
    }
}

bool callGlProgramUniform3ui(Stack* stack, bool pushReturn) {
    uint32_t v2 = stack->pop<uint32_t>();
    uint32_t v1 = stack->pop<uint32_t>();
    uint32_t v0 = stack->pop<uint32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform3ui(%" PRIu32 ", %" PRId32 ", %" PRIu32 ", %" PRIu32
                   ", %" PRIu32 ")\n",
                   program, location, v0, v1, v2);
        if (glProgramUniform3ui != nullptr) {
            glProgramUniform3ui(program, location, v0, v1, v2);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform3ui returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform3ui\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform3ui\n");
        return false;
    }
}

bool callGlProgramUniform3uiv(Stack* stack, bool pushReturn) {
    uint32_t* value = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform3uiv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)\n", program,
                   location, count, value);
        if (glProgramUniform3uiv != nullptr) {
            glProgramUniform3uiv(program, location, count, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform3uiv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform3uiv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform3uiv\n");
        return false;
    }
}

bool callGlProgramUniform4f(Stack* stack, bool pushReturn) {
    float v3 = stack->pop<float>();
    float v2 = stack->pop<float>();
    float v1 = stack->pop<float>();
    float v0 = stack->pop<float>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform4f(%" PRIu32 ", %" PRId32 ", %f, %f, %f, %f)\n", program,
                   location, v0, v1, v2, v3);
        if (glProgramUniform4f != nullptr) {
            glProgramUniform4f(program, location, v0, v1, v2, v3);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform4f returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform4f\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform4f\n");
        return false;
    }
}

bool callGlProgramUniform4fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform4fv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)\n", program,
                   location, count, value);
        if (glProgramUniform4fv != nullptr) {
            glProgramUniform4fv(program, location, count, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform4fv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform4fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform4fv\n");
        return false;
    }
}

bool callGlProgramUniform4i(Stack* stack, bool pushReturn) {
    int32_t v3 = stack->pop<int32_t>();
    int32_t v2 = stack->pop<int32_t>();
    int32_t v1 = stack->pop<int32_t>();
    int32_t v0 = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform4i(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ")\n",
                   program, location, v0, v1, v2, v3);
        if (glProgramUniform4i != nullptr) {
            glProgramUniform4i(program, location, v0, v1, v2, v3);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform4i returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform4i\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform4i\n");
        return false;
    }
}

bool callGlProgramUniform4iv(Stack* stack, bool pushReturn) {
    int32_t* value = stack->pop<int32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform4iv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)\n", program,
                   location, count, value);
        if (glProgramUniform4iv != nullptr) {
            glProgramUniform4iv(program, location, count, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform4iv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform4iv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform4iv\n");
        return false;
    }
}

bool callGlProgramUniform4ui(Stack* stack, bool pushReturn) {
    uint32_t v3 = stack->pop<uint32_t>();
    uint32_t v2 = stack->pop<uint32_t>();
    uint32_t v1 = stack->pop<uint32_t>();
    uint32_t v0 = stack->pop<uint32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform4ui(%" PRIu32 ", %" PRId32 ", %" PRIu32 ", %" PRIu32
                   ", %" PRIu32 ", %" PRIu32 ")\n",
                   program, location, v0, v1, v2, v3);
        if (glProgramUniform4ui != nullptr) {
            glProgramUniform4ui(program, location, v0, v1, v2, v3);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform4ui returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform4ui\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform4ui\n");
        return false;
    }
}

bool callGlProgramUniform4uiv(Stack* stack, bool pushReturn) {
    uint32_t* value = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniform4uiv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %p)\n", program,
                   location, count, value);
        if (glProgramUniform4uiv != nullptr) {
            glProgramUniform4uiv(program, location, count, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniform4uiv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniform4uiv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniform4uiv\n");
        return false;
    }
}

bool callGlProgramUniformMatrix2fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniformMatrix2fv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRIu8
                   ", %p)\n",
                   program, location, count, transpose, value);
        if (glProgramUniformMatrix2fv != nullptr) {
            glProgramUniformMatrix2fv(program, location, count, transpose, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniformMatrix2fv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformMatrix2fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix2fv\n");
        return false;
    }
}

bool callGlProgramUniformMatrix2x3fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniformMatrix2x3fv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRIu8
                   ", %p)\n",
                   program, location, count, transpose, value);
        if (glProgramUniformMatrix2x3fv != nullptr) {
            glProgramUniformMatrix2x3fv(program, location, count, transpose, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniformMatrix2x3fv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformMatrix2x3fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix2x3fv\n");
        return false;
    }
}

bool callGlProgramUniformMatrix2x4fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniformMatrix2x4fv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRIu8
                   ", %p)\n",
                   program, location, count, transpose, value);
        if (glProgramUniformMatrix2x4fv != nullptr) {
            glProgramUniformMatrix2x4fv(program, location, count, transpose, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniformMatrix2x4fv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformMatrix2x4fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix2x4fv\n");
        return false;
    }
}

bool callGlProgramUniformMatrix3fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniformMatrix3fv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRIu8
                   ", %p)\n",
                   program, location, count, transpose, value);
        if (glProgramUniformMatrix3fv != nullptr) {
            glProgramUniformMatrix3fv(program, location, count, transpose, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniformMatrix3fv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformMatrix3fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix3fv\n");
        return false;
    }
}

bool callGlProgramUniformMatrix3x2fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniformMatrix3x2fv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRIu8
                   ", %p)\n",
                   program, location, count, transpose, value);
        if (glProgramUniformMatrix3x2fv != nullptr) {
            glProgramUniformMatrix3x2fv(program, location, count, transpose, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniformMatrix3x2fv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformMatrix3x2fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix3x2fv\n");
        return false;
    }
}

bool callGlProgramUniformMatrix3x4fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniformMatrix3x4fv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRIu8
                   ", %p)\n",
                   program, location, count, transpose, value);
        if (glProgramUniformMatrix3x4fv != nullptr) {
            glProgramUniformMatrix3x4fv(program, location, count, transpose, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniformMatrix3x4fv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformMatrix3x4fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix3x4fv\n");
        return false;
    }
}

bool callGlProgramUniformMatrix4fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniformMatrix4fv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRIu8
                   ", %p)\n",
                   program, location, count, transpose, value);
        if (glProgramUniformMatrix4fv != nullptr) {
            glProgramUniformMatrix4fv(program, location, count, transpose, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniformMatrix4fv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformMatrix4fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix4fv\n");
        return false;
    }
}

bool callGlProgramUniformMatrix4x2fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniformMatrix4x2fv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRIu8
                   ", %p)\n",
                   program, location, count, transpose, value);
        if (glProgramUniformMatrix4x2fv != nullptr) {
            glProgramUniformMatrix4x2fv(program, location, count, transpose, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniformMatrix4x2fv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformMatrix4x2fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix4x2fv\n");
        return false;
    }
}

bool callGlProgramUniformMatrix4x3fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glProgramUniformMatrix4x3fv(%" PRIu32 ", %" PRId32 ", %" PRId32 ", %" PRIu8
                   ", %p)\n",
                   program, location, count, transpose, value);
        if (glProgramUniformMatrix4x3fv != nullptr) {
            glProgramUniformMatrix4x3fv(program, location, count, transpose, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glProgramUniformMatrix4x3fv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glProgramUniformMatrix4x3fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glProgramUniformMatrix4x3fv\n");
        return false;
    }
}

bool callGlReleaseShaderCompiler(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glReleaseShaderCompiler()\n");
        if (glReleaseShaderCompiler != nullptr) {
            glReleaseShaderCompiler();
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glReleaseShaderCompiler returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glReleaseShaderCompiler\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glReleaseShaderCompiler\n");
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
        GAPID_INFO("glShaderBinary(%" PRId32 ", %p, %u, %p, %" PRId32 ")\n", count, shaders,
                   binary_format, binary, binary_size);
        if (glShaderBinary != nullptr) {
            glShaderBinary(count, shaders, binary_format, binary, binary_size);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glShaderBinary returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glShaderBinary\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glShaderBinary\n");
        return false;
    }
}

bool callGlShaderSource(Stack* stack, bool pushReturn) {
    int32_t* length = stack->pop<int32_t*>();
    char** source = stack->pop<char**>();
    int32_t count = stack->pop<int32_t>();
    uint32_t shader = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glShaderSource(%" PRIu32 ", %" PRId32 ", %p, %p)\n", shader, count, source,
                   length);
        if (glShaderSource != nullptr) {
            glShaderSource(shader, count, source, length);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glShaderSource returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glShaderSource\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glShaderSource\n");
        return false;
    }
}

bool callGlUniform1f(Stack* stack, bool pushReturn) {
    float value = stack->pop<float>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform1f(%" PRId32 ", %f)\n", location, value);
        if (glUniform1f != nullptr) {
            glUniform1f(location, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniform1f returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform1f\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform1f\n");
        return false;
    }
}

bool callGlUniform1fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform1fv(%" PRId32 ", %" PRId32 ", %p)\n", location, count, values);
        if (glUniform1fv != nullptr) {
            glUniform1fv(location, count, values);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniform1fv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform1fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform1fv\n");
        return false;
    }
}

bool callGlUniform1i(Stack* stack, bool pushReturn) {
    int32_t value = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform1i(%" PRId32 ", %" PRId32 ")\n", location, value);
        if (glUniform1i != nullptr) {
            glUniform1i(location, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniform1i returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform1i\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform1i\n");
        return false;
    }
}

bool callGlUniform1iv(Stack* stack, bool pushReturn) {
    int32_t* values = stack->pop<int32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform1iv(%" PRId32 ", %" PRId32 ", %p)\n", location, count, values);
        if (glUniform1iv != nullptr) {
            glUniform1iv(location, count, values);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniform1iv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform1iv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform1iv\n");
        return false;
    }
}

bool callGlUniform1ui(Stack* stack, bool pushReturn) {
    uint32_t v0 = stack->pop<uint32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform1ui(%" PRId32 ", %" PRIu32 ")\n", location, v0);
        if (glUniform1ui != nullptr) {
            glUniform1ui(location, v0);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniform1ui returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform1ui\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform1ui\n");
        return false;
    }
}

bool callGlUniform1uiv(Stack* stack, bool pushReturn) {
    uint32_t* value = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform1uiv(%" PRId32 ", %" PRId32 ", %p)\n", location, count, value);
        if (glUniform1uiv != nullptr) {
            glUniform1uiv(location, count, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniform1uiv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform1uiv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform1uiv\n");
        return false;
    }
}

bool callGlUniform2f(Stack* stack, bool pushReturn) {
    float value1 = stack->pop<float>();
    float value0 = stack->pop<float>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform2f(%" PRId32 ", %f, %f)\n", location, value0, value1);
        if (glUniform2f != nullptr) {
            glUniform2f(location, value0, value1);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniform2f returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform2f\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform2f\n");
        return false;
    }
}

bool callGlUniform2fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform2fv(%" PRId32 ", %" PRId32 ", %p)\n", location, count, values);
        if (glUniform2fv != nullptr) {
            glUniform2fv(location, count, values);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniform2fv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform2fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform2fv\n");
        return false;
    }
}

bool callGlUniform2i(Stack* stack, bool pushReturn) {
    int32_t value1 = stack->pop<int32_t>();
    int32_t value0 = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform2i(%" PRId32 ", %" PRId32 ", %" PRId32 ")\n", location, value0,
                   value1);
        if (glUniform2i != nullptr) {
            glUniform2i(location, value0, value1);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniform2i returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform2i\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform2i\n");
        return false;
    }
}

bool callGlUniform2iv(Stack* stack, bool pushReturn) {
    int32_t* values = stack->pop<int32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform2iv(%" PRId32 ", %" PRId32 ", %p)\n", location, count, values);
        if (glUniform2iv != nullptr) {
            glUniform2iv(location, count, values);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniform2iv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform2iv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform2iv\n");
        return false;
    }
}

bool callGlUniform2ui(Stack* stack, bool pushReturn) {
    uint32_t v1 = stack->pop<uint32_t>();
    uint32_t v0 = stack->pop<uint32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform2ui(%" PRId32 ", %" PRIu32 ", %" PRIu32 ")\n", location, v0, v1);
        if (glUniform2ui != nullptr) {
            glUniform2ui(location, v0, v1);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniform2ui returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform2ui\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform2ui\n");
        return false;
    }
}

bool callGlUniform2uiv(Stack* stack, bool pushReturn) {
    uint32_t* value = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform2uiv(%" PRId32 ", %" PRId32 ", %p)\n", location, count, value);
        if (glUniform2uiv != nullptr) {
            glUniform2uiv(location, count, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniform2uiv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform2uiv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform2uiv\n");
        return false;
    }
}

bool callGlUniform3f(Stack* stack, bool pushReturn) {
    float value2 = stack->pop<float>();
    float value1 = stack->pop<float>();
    float value0 = stack->pop<float>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform3f(%" PRId32 ", %f, %f, %f)\n", location, value0, value1, value2);
        if (glUniform3f != nullptr) {
            glUniform3f(location, value0, value1, value2);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniform3f returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform3f\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform3f\n");
        return false;
    }
}

bool callGlUniform3fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform3fv(%" PRId32 ", %" PRId32 ", %p)\n", location, count, values);
        if (glUniform3fv != nullptr) {
            glUniform3fv(location, count, values);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniform3fv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform3fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform3fv\n");
        return false;
    }
}

bool callGlUniform3i(Stack* stack, bool pushReturn) {
    int32_t value2 = stack->pop<int32_t>();
    int32_t value1 = stack->pop<int32_t>();
    int32_t value0 = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform3i(%" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ")\n", location,
                   value0, value1, value2);
        if (glUniform3i != nullptr) {
            glUniform3i(location, value0, value1, value2);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniform3i returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform3i\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform3i\n");
        return false;
    }
}

bool callGlUniform3iv(Stack* stack, bool pushReturn) {
    int32_t* values = stack->pop<int32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform3iv(%" PRId32 ", %" PRId32 ", %p)\n", location, count, values);
        if (glUniform3iv != nullptr) {
            glUniform3iv(location, count, values);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniform3iv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform3iv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform3iv\n");
        return false;
    }
}

bool callGlUniform3ui(Stack* stack, bool pushReturn) {
    uint32_t v2 = stack->pop<uint32_t>();
    uint32_t v1 = stack->pop<uint32_t>();
    uint32_t v0 = stack->pop<uint32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform3ui(%" PRId32 ", %" PRIu32 ", %" PRIu32 ", %" PRIu32 ")\n", location,
                   v0, v1, v2);
        if (glUniform3ui != nullptr) {
            glUniform3ui(location, v0, v1, v2);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniform3ui returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform3ui\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform3ui\n");
        return false;
    }
}

bool callGlUniform3uiv(Stack* stack, bool pushReturn) {
    uint32_t* value = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform3uiv(%" PRId32 ", %" PRId32 ", %p)\n", location, count, value);
        if (glUniform3uiv != nullptr) {
            glUniform3uiv(location, count, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniform3uiv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform3uiv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform3uiv\n");
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
        GAPID_INFO("glUniform4f(%" PRId32 ", %f, %f, %f, %f)\n", location, value0, value1, value2,
                   value3);
        if (glUniform4f != nullptr) {
            glUniform4f(location, value0, value1, value2, value3);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniform4f returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform4f\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform4f\n");
        return false;
    }
}

bool callGlUniform4fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform4fv(%" PRId32 ", %" PRId32 ", %p)\n", location, count, values);
        if (glUniform4fv != nullptr) {
            glUniform4fv(location, count, values);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniform4fv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform4fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform4fv\n");
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
        GAPID_INFO("glUniform4i(%" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ")\n",
                   location, value0, value1, value2, value3);
        if (glUniform4i != nullptr) {
            glUniform4i(location, value0, value1, value2, value3);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniform4i returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform4i\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform4i\n");
        return false;
    }
}

bool callGlUniform4iv(Stack* stack, bool pushReturn) {
    int32_t* values = stack->pop<int32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform4iv(%" PRId32 ", %" PRId32 ", %p)\n", location, count, values);
        if (glUniform4iv != nullptr) {
            glUniform4iv(location, count, values);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniform4iv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform4iv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform4iv\n");
        return false;
    }
}

bool callGlUniform4ui(Stack* stack, bool pushReturn) {
    uint32_t v3 = stack->pop<uint32_t>();
    uint32_t v2 = stack->pop<uint32_t>();
    uint32_t v1 = stack->pop<uint32_t>();
    uint32_t v0 = stack->pop<uint32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform4ui(%" PRId32 ", %" PRIu32 ", %" PRIu32 ", %" PRIu32 ", %" PRIu32
                   ")\n",
                   location, v0, v1, v2, v3);
        if (glUniform4ui != nullptr) {
            glUniform4ui(location, v0, v1, v2, v3);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniform4ui returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform4ui\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform4ui\n");
        return false;
    }
}

bool callGlUniform4uiv(Stack* stack, bool pushReturn) {
    uint32_t* value = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniform4uiv(%" PRId32 ", %" PRId32 ", %p)\n", location, count, value);
        if (glUniform4uiv != nullptr) {
            glUniform4uiv(location, count, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniform4uiv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniform4uiv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniform4uiv\n");
        return false;
    }
}

bool callGlUniformBlockBinding(Stack* stack, bool pushReturn) {
    uint32_t uniform_block_binding = stack->pop<uint32_t>();
    uint32_t uniform_block_index = stack->pop<uint32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformBlockBinding(%" PRIu32 ", %" PRIu32 ", %" PRIu32 ")\n", program,
                   uniform_block_index, uniform_block_binding);
        if (glUniformBlockBinding != nullptr) {
            glUniformBlockBinding(program, uniform_block_index, uniform_block_binding);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniformBlockBinding returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformBlockBinding\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformBlockBinding\n");
        return false;
    }
}

bool callGlUniformMatrix2fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformMatrix2fv(%" PRId32 ", %" PRId32 ", %" PRIu8 ", %p)\n", location,
                   count, transpose, values);
        if (glUniformMatrix2fv != nullptr) {
            glUniformMatrix2fv(location, count, transpose, values);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniformMatrix2fv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformMatrix2fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformMatrix2fv\n");
        return false;
    }
}

bool callGlUniformMatrix2x3fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformMatrix2x3fv(%" PRId32 ", %" PRId32 ", %" PRIu8 ", %p)\n", location,
                   count, transpose, value);
        if (glUniformMatrix2x3fv != nullptr) {
            glUniformMatrix2x3fv(location, count, transpose, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniformMatrix2x3fv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformMatrix2x3fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformMatrix2x3fv\n");
        return false;
    }
}

bool callGlUniformMatrix2x4fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformMatrix2x4fv(%" PRId32 ", %" PRId32 ", %" PRIu8 ", %p)\n", location,
                   count, transpose, value);
        if (glUniformMatrix2x4fv != nullptr) {
            glUniformMatrix2x4fv(location, count, transpose, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniformMatrix2x4fv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformMatrix2x4fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformMatrix2x4fv\n");
        return false;
    }
}

bool callGlUniformMatrix3fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformMatrix3fv(%" PRId32 ", %" PRId32 ", %" PRIu8 ", %p)\n", location,
                   count, transpose, values);
        if (glUniformMatrix3fv != nullptr) {
            glUniformMatrix3fv(location, count, transpose, values);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniformMatrix3fv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformMatrix3fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformMatrix3fv\n");
        return false;
    }
}

bool callGlUniformMatrix3x2fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformMatrix3x2fv(%" PRId32 ", %" PRId32 ", %" PRIu8 ", %p)\n", location,
                   count, transpose, value);
        if (glUniformMatrix3x2fv != nullptr) {
            glUniformMatrix3x2fv(location, count, transpose, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniformMatrix3x2fv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformMatrix3x2fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformMatrix3x2fv\n");
        return false;
    }
}

bool callGlUniformMatrix3x4fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformMatrix3x4fv(%" PRId32 ", %" PRId32 ", %" PRIu8 ", %p)\n", location,
                   count, transpose, value);
        if (glUniformMatrix3x4fv != nullptr) {
            glUniformMatrix3x4fv(location, count, transpose, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniformMatrix3x4fv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformMatrix3x4fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformMatrix3x4fv\n");
        return false;
    }
}

bool callGlUniformMatrix4fv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformMatrix4fv(%" PRId32 ", %" PRId32 ", %" PRIu8 ", %p)\n", location,
                   count, transpose, values);
        if (glUniformMatrix4fv != nullptr) {
            glUniformMatrix4fv(location, count, transpose, values);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniformMatrix4fv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformMatrix4fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformMatrix4fv\n");
        return false;
    }
}

bool callGlUniformMatrix4x2fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformMatrix4x2fv(%" PRId32 ", %" PRId32 ", %" PRIu8 ", %p)\n", location,
                   count, transpose, value);
        if (glUniformMatrix4x2fv != nullptr) {
            glUniformMatrix4x2fv(location, count, transpose, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniformMatrix4x2fv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformMatrix4x2fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformMatrix4x2fv\n");
        return false;
    }
}

bool callGlUniformMatrix4x3fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint8_t transpose = stack->pop<uint8_t>();
    int32_t count = stack->pop<int32_t>();
    int32_t location = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUniformMatrix4x3fv(%" PRId32 ", %" PRId32 ", %" PRIu8 ", %p)\n", location,
                   count, transpose, value);
        if (glUniformMatrix4x3fv != nullptr) {
            glUniformMatrix4x3fv(location, count, transpose, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUniformMatrix4x3fv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUniformMatrix4x3fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUniformMatrix4x3fv\n");
        return false;
    }
}

bool callGlUseProgram(Stack* stack, bool pushReturn) {
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUseProgram(%" PRIu32 ")\n", program);
        if (glUseProgram != nullptr) {
            glUseProgram(program);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUseProgram returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUseProgram\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUseProgram\n");
        return false;
    }
}

bool callGlUseProgramStages(Stack* stack, bool pushReturn) {
    uint32_t program = stack->pop<uint32_t>();
    GLbitfield stages = stack->pop<GLbitfield>();
    uint32_t pipeline = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glUseProgramStages(%" PRIu32 ", %u, %" PRIu32 ")\n", pipeline, stages, program);
        if (glUseProgramStages != nullptr) {
            glUseProgramStages(pipeline, stages, program);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glUseProgramStages returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glUseProgramStages\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glUseProgramStages\n");
        return false;
    }
}

bool callGlValidateProgram(Stack* stack, bool pushReturn) {
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glValidateProgram(%" PRIu32 ")\n", program);
        if (glValidateProgram != nullptr) {
            glValidateProgram(program);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glValidateProgram returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glValidateProgram\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glValidateProgram\n");
        return false;
    }
}

bool callGlValidateProgramPipeline(Stack* stack, bool pushReturn) {
    uint32_t pipeline = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glValidateProgramPipeline(%" PRIu32 ")\n", pipeline);
        if (glValidateProgramPipeline != nullptr) {
            glValidateProgramPipeline(pipeline);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glValidateProgramPipeline returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glValidateProgramPipeline\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glValidateProgramPipeline\n");
        return false;
    }
}

bool callGlCullFace(Stack* stack, bool pushReturn) {
    GLenum mode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glCullFace(%u)\n", mode);
        if (glCullFace != nullptr) {
            glCullFace(mode);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glCullFace returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCullFace\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCullFace\n");
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
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDepthRangef returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDepthRangef\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDepthRangef\n");
        return false;
    }
}

bool callGlFrontFace(Stack* stack, bool pushReturn) {
    GLenum orientation = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glFrontFace(%u)\n", orientation);
        if (glFrontFace != nullptr) {
            glFrontFace(orientation);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glFrontFace returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFrontFace\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFrontFace\n");
        return false;
    }
}

bool callGlGetMultisamplefv(Stack* stack, bool pushReturn) {
    float* val = stack->pop<float*>();
    uint32_t index = stack->pop<uint32_t>();
    GLenum pname = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetMultisamplefv(%u, %" PRIu32 ", %p)\n", pname, index, val);
        if (glGetMultisamplefv != nullptr) {
            glGetMultisamplefv(pname, index, val);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetMultisamplefv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetMultisamplefv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetMultisamplefv\n");
        return false;
    }
}

bool callGlLineWidth(Stack* stack, bool pushReturn) {
    float width = stack->pop<float>();
    if (stack->isValid()) {
        GAPID_INFO("glLineWidth(%f)\n", width);
        if (glLineWidth != nullptr) {
            glLineWidth(width);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glLineWidth returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glLineWidth\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glLineWidth\n");
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
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glPolygonOffset returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPolygonOffset\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPolygonOffset\n");
        return false;
    }
}

bool callGlViewport(Stack* stack, bool pushReturn) {
    int32_t height = stack->pop<int32_t>();
    int32_t width = stack->pop<int32_t>();
    int32_t y = stack->pop<int32_t>();
    int32_t x = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glViewport(%" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ")\n", x, y, width,
                   height);
        if (glViewport != nullptr) {
            glViewport(x, y, width, height);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glViewport returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glViewport\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glViewport\n");
        return false;
    }
}

bool callGlGetBooleaniV(Stack* stack, bool pushReturn) {
    uint8_t* data = stack->pop<uint8_t*>();
    uint32_t index = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetBooleani_v(%u, %" PRIu32 ", %p)\n", target, index, data);
        if (glGetBooleani_v != nullptr) {
            glGetBooleani_v(target, index, data);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetBooleani_v returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetBooleani_v\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetBooleani_v\n");
        return false;
    }
}

bool callGlGetBooleanv(Stack* stack, bool pushReturn) {
    uint8_t* values = stack->pop<uint8_t*>();
    GLenum param = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetBooleanv(%u, %p)\n", param, values);
        if (glGetBooleanv != nullptr) {
            glGetBooleanv(param, values);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetBooleanv returned error: 0x%x\n", err);
            }
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
    GLenum param = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetFloatv(%u, %p)\n", param, values);
        if (glGetFloatv != nullptr) {
            glGetFloatv(param, values);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetFloatv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetFloatv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetFloatv\n");
        return false;
    }
}

bool callGlGetInteger64iV(Stack* stack, bool pushReturn) {
    int64_t* data = stack->pop<int64_t*>();
    uint32_t index = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetInteger64i_v(%u, %" PRIu32 ", %p)\n", target, index, data);
        if (glGetInteger64i_v != nullptr) {
            glGetInteger64i_v(target, index, data);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetInteger64i_v returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetInteger64i_v\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetInteger64i_v\n");
        return false;
    }
}

bool callGlGetInteger64v(Stack* stack, bool pushReturn) {
    int64_t* data = stack->pop<int64_t*>();
    GLenum pname = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetInteger64v(%u, %p)\n", pname, data);
        if (glGetInteger64v != nullptr) {
            glGetInteger64v(pname, data);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetInteger64v returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetInteger64v\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetInteger64v\n");
        return false;
    }
}

bool callGlGetIntegeriV(Stack* stack, bool pushReturn) {
    int32_t* data = stack->pop<int32_t*>();
    uint32_t index = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetIntegeri_v(%u, %" PRIu32 ", %p)\n", target, index, data);
        if (glGetIntegeri_v != nullptr) {
            glGetIntegeri_v(target, index, data);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetIntegeri_v returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetIntegeri_v\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetIntegeri_v\n");
        return false;
    }
}

bool callGlGetIntegerv(Stack* stack, bool pushReturn) {
    int32_t* values = stack->pop<int32_t*>();
    GLenum param = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetIntegerv(%u, %p)\n", param, values);
        if (glGetIntegerv != nullptr) {
            glGetIntegerv(param, values);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetIntegerv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetIntegerv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetIntegerv\n");
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
        GAPID_INFO("glGetInternalformativ(%u, %u, %u, %" PRId32 ", %p)\n", target, internalformat,
                   pname, bufSize, params);
        if (glGetInternalformativ != nullptr) {
            glGetInternalformativ(target, internalformat, pname, bufSize, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetInternalformativ returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetInternalformativ\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetInternalformativ\n");
        return false;
    }
}

bool callGlGetString(Stack* stack, bool pushReturn) {
    GLenum param = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetString(%u)\n", param);
        if (glGetString != nullptr) {
            uint8_t* return_value = glGetString(param);
            GAPID_INFO("Returned: %p\n", return_value);
            if (pushReturn) {
                stack->push<uint8_t*>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetString returned error: 0x%x\n", err);
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

bool callGlGetStringi(Stack* stack, bool pushReturn) {
    uint32_t index = stack->pop<uint32_t>();
    GLenum name = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetStringi(%u, %" PRIu32 ")\n", name, index);
        if (glGetStringi != nullptr) {
            uint8_t* return_value = glGetStringi(name, index);
            GAPID_INFO("Returned: %p\n", return_value);
            if (pushReturn) {
                stack->push<uint8_t*>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetStringi returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetStringi\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetStringi\n");
        return false;
    }
}

bool callGlIsEnabled(Stack* stack, bool pushReturn) {
    GLenum capability = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glIsEnabled(%u)\n", capability);
        if (glIsEnabled != nullptr) {
            uint8_t return_value = glIsEnabled(capability);
            GAPID_INFO("Returned: %" PRIu8 "\n", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glIsEnabled returned error: 0x%x\n", err);
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

bool callGlClientWaitSync(Stack* stack, bool pushReturn) {
    uint64_t timeout = stack->pop<uint64_t>();
    GLbitfield syncFlags = stack->pop<GLbitfield>();
    uint64_t sync = stack->pop<uint64_t>();
    if (stack->isValid()) {
        GAPID_INFO("glClientWaitSync(%" PRIu64 ", %u, %" PRIu64 ")\n", sync, syncFlags, timeout);
        if (glClientWaitSync != nullptr) {
            GLenum return_value = glClientWaitSync(sync, syncFlags, timeout);
            GAPID_INFO("Returned: %u\n", return_value);
            if (pushReturn) {
                stack->push<GLenum>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glClientWaitSync returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glClientWaitSync\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glClientWaitSync\n");
        return false;
    }
}

bool callGlDeleteSync(Stack* stack, bool pushReturn) {
    uint64_t sync = stack->pop<uint64_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteSync(%" PRIu64 ")\n", sync);
        if (glDeleteSync != nullptr) {
            glDeleteSync(sync);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDeleteSync returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteSync\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteSync\n");
        return false;
    }
}

bool callGlFenceSync(Stack* stack, bool pushReturn) {
    GLbitfield syncFlags = stack->pop<GLbitfield>();
    GLenum condition = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glFenceSync(%u, %u)\n", condition, syncFlags);
        if (glFenceSync != nullptr) {
            uint64_t return_value = glFenceSync(condition, syncFlags);
            GAPID_INFO("Returned: %" PRIu64 "\n", return_value);
            if (pushReturn) {
                stack->push<uint64_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glFenceSync returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glFenceSync\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glFenceSync\n");
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
        GAPID_INFO("glGetSynciv(%" PRIu64 ", %u, %" PRId32 ", %p, %p)\n", sync, pname, bufSize,
                   length, values);
        if (glGetSynciv != nullptr) {
            glGetSynciv(sync, pname, bufSize, length, values);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetSynciv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetSynciv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetSynciv\n");
        return false;
    }
}

bool callGlIsSync(Stack* stack, bool pushReturn) {
    uint64_t sync = stack->pop<uint64_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsSync(%" PRIu64 ")\n", sync);
        if (glIsSync != nullptr) {
            uint8_t return_value = glIsSync(sync);
            GAPID_INFO("Returned: %" PRIu8 "\n", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glIsSync returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsSync\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsSync\n");
        return false;
    }
}

bool callGlWaitSync(Stack* stack, bool pushReturn) {
    uint64_t timeout = stack->pop<uint64_t>();
    GLbitfield syncFlags = stack->pop<GLbitfield>();
    uint64_t sync = stack->pop<uint64_t>();
    if (stack->isValid()) {
        GAPID_INFO("glWaitSync(%" PRIu64 ", %u, %" PRIu64 ")\n", sync, syncFlags, timeout);
        if (glWaitSync != nullptr) {
            glWaitSync(sync, syncFlags, timeout);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glWaitSync returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glWaitSync\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glWaitSync\n");
        return false;
    }
}

bool callGlActiveTexture(Stack* stack, bool pushReturn) {
    GLenum unit = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glActiveTexture(%u)\n", unit);
        if (glActiveTexture != nullptr) {
            glActiveTexture(unit);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glActiveTexture returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glActiveTexture\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glActiveTexture\n");
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
                   ", %u, %u)\n",
                   unit, texture, level, layered, layer, access, format);
        if (glBindImageTexture != nullptr) {
            glBindImageTexture(unit, texture, level, layered, layer, access, format);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBindImageTexture returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindImageTexture\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindImageTexture\n");
        return false;
    }
}

bool callGlBindSampler(Stack* stack, bool pushReturn) {
    uint32_t sampler = stack->pop<uint32_t>();
    uint32_t unit = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBindSampler(%" PRIu32 ", %" PRIu32 ")\n", unit, sampler);
        if (glBindSampler != nullptr) {
            glBindSampler(unit, sampler);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBindSampler returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindSampler\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindSampler\n");
        return false;
    }
}

bool callGlBindTexture(Stack* stack, bool pushReturn) {
    uint32_t texture = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBindTexture(%u, %" PRIu32 ")\n", target, texture);
        if (glBindTexture != nullptr) {
            glBindTexture(target, texture);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBindTexture returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindTexture\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindTexture\n");
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
                   ", %" PRId32 ", %p)\n",
                   target, level, format, width, height, border, image_size, data);
        if (glCompressedTexImage2D != nullptr) {
            glCompressedTexImage2D(target, level, format, width, height, border, image_size, data);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glCompressedTexImage2D returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCompressedTexImage2D\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCompressedTexImage2D\n");
        return false;
    }
}

bool callGlCompressedTexImage3D(Stack* stack, bool pushReturn) {
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
        GAPID_INFO("glCompressedTexImage3D(%u, %" PRId32 ", %u, %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %" PRId32 ", %p)\n",
                   target, level, internalformat, width, height, depth, border, imageSize, data);
        if (glCompressedTexImage3D != nullptr) {
            glCompressedTexImage3D(target, level, internalformat, width, height, depth, border,
                                   imageSize, data);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glCompressedTexImage3D returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCompressedTexImage3D\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCompressedTexImage3D\n");
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
                   ", %" PRId32 ", %u, %" PRId32 ", %p)\n",
                   target, level, xoffset, yoffset, width, height, format, image_size, data);
        if (glCompressedTexSubImage2D != nullptr) {
            glCompressedTexSubImage2D(target, level, xoffset, yoffset, width, height, format,
                                      image_size, data);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glCompressedTexSubImage2D returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCompressedTexSubImage2D\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCompressedTexSubImage2D\n");
        return false;
    }
}

bool callGlCompressedTexSubImage3D(Stack* stack, bool pushReturn) {
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
        GAPID_INFO("glCompressedTexSubImage3D(%u, %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %u, %" PRId32 ", %p)\n",
                   target, level, xoffset, yoffset, zoffset, width, height, depth, format,
                   imageSize, data);
        if (glCompressedTexSubImage3D != nullptr) {
            glCompressedTexSubImage3D(target, level, xoffset, yoffset, zoffset, width, height,
                                      depth, format, imageSize, data);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glCompressedTexSubImage3D returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCompressedTexSubImage3D\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCompressedTexSubImage3D\n");
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
                   ", %" PRId32 ", %" PRId32 ")\n",
                   target, level, format, x, y, width, height, border);
        if (glCopyTexImage2D != nullptr) {
            glCopyTexImage2D(target, level, format, x, y, width, height, border);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glCopyTexImage2D returned error: 0x%x\n", err);
            }
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
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glCopyTexSubImage2D(%u, %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %" PRId32 ", %" PRId32 ")\n",
                   target, level, xoffset, yoffset, x, y, width, height);
        if (glCopyTexSubImage2D != nullptr) {
            glCopyTexSubImage2D(target, level, xoffset, yoffset, x, y, width, height);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glCopyTexSubImage2D returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCopyTexSubImage2D\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCopyTexSubImage2D\n");
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
                   ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32 ")\n",
                   target, level, xoffset, yoffset, zoffset, x, y, width, height);
        if (glCopyTexSubImage3D != nullptr) {
            glCopyTexSubImage3D(target, level, xoffset, yoffset, zoffset, x, y, width, height);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glCopyTexSubImage3D returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glCopyTexSubImage3D\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glCopyTexSubImage3D\n");
        return false;
    }
}

bool callGlDeleteSamplers(Stack* stack, bool pushReturn) {
    uint32_t* samplers = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteSamplers(%" PRId32 ", %p)\n", count, samplers);
        if (glDeleteSamplers != nullptr) {
            glDeleteSamplers(count, samplers);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDeleteSamplers returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteSamplers\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteSamplers\n");
        return false;
    }
}

bool callGlDeleteTextures(Stack* stack, bool pushReturn) {
    uint32_t* textures = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteTextures(%" PRId32 ", %p)\n", count, textures);
        if (glDeleteTextures != nullptr) {
            glDeleteTextures(count, textures);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDeleteTextures returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteTextures\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteTextures\n");
        return false;
    }
}

bool callGlGenSamplers(Stack* stack, bool pushReturn) {
    uint32_t* samplers = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenSamplers(%" PRId32 ", %p)\n", count, samplers);
        if (glGenSamplers != nullptr) {
            glGenSamplers(count, samplers);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGenSamplers returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenSamplers\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenSamplers\n");
        return false;
    }
}

bool callGlGenTextures(Stack* stack, bool pushReturn) {
    uint32_t* textures = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenTextures(%" PRId32 ", %p)\n", count, textures);
        if (glGenTextures != nullptr) {
            glGenTextures(count, textures);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGenTextures returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenTextures\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenTextures\n");
        return false;
    }
}

bool callGlGenerateMipmap(Stack* stack, bool pushReturn) {
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGenerateMipmap(%u)\n", target);
        if (glGenerateMipmap != nullptr) {
            glGenerateMipmap(target);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGenerateMipmap returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenerateMipmap\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenerateMipmap\n");
        return false;
    }
}

bool callGlGetSamplerParameterfv(Stack* stack, bool pushReturn) {
    float* params = stack->pop<float*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetSamplerParameterfv(%" PRIu32 ", %u, %p)\n", sampler, pname, params);
        if (glGetSamplerParameterfv != nullptr) {
            glGetSamplerParameterfv(sampler, pname, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetSamplerParameterfv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetSamplerParameterfv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetSamplerParameterfv\n");
        return false;
    }
}

bool callGlGetSamplerParameteriv(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetSamplerParameteriv(%" PRIu32 ", %u, %p)\n", sampler, pname, params);
        if (glGetSamplerParameteriv != nullptr) {
            glGetSamplerParameteriv(sampler, pname, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetSamplerParameteriv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetSamplerParameteriv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetSamplerParameteriv\n");
        return false;
    }
}

bool callGlGetTexLevelParameterfv(Stack* stack, bool pushReturn) {
    float* params = stack->pop<float*>();
    GLenum pname = stack->pop<GLenum>();
    int32_t level = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetTexLevelParameterfv(%u, %" PRId32 ", %u, %p)\n", target, level, pname,
                   params);
        if (glGetTexLevelParameterfv != nullptr) {
            glGetTexLevelParameterfv(target, level, pname, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetTexLevelParameterfv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetTexLevelParameterfv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetTexLevelParameterfv\n");
        return false;
    }
}

bool callGlGetTexLevelParameteriv(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    int32_t level = stack->pop<int32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetTexLevelParameteriv(%u, %" PRId32 ", %u, %p)\n", target, level, pname,
                   params);
        if (glGetTexLevelParameteriv != nullptr) {
            glGetTexLevelParameteriv(target, level, pname, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetTexLevelParameteriv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetTexLevelParameteriv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetTexLevelParameteriv\n");
        return false;
    }
}

bool callGlGetTexParameterfv(Stack* stack, bool pushReturn) {
    float* values = stack->pop<float*>();
    GLenum parameter = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetTexParameterfv(%u, %u, %p)\n", target, parameter, values);
        if (glGetTexParameterfv != nullptr) {
            glGetTexParameterfv(target, parameter, values);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetTexParameterfv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetTexParameterfv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetTexParameterfv\n");
        return false;
    }
}

bool callGlGetTexParameteriv(Stack* stack, bool pushReturn) {
    int32_t* values = stack->pop<int32_t*>();
    GLenum parameter = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glGetTexParameteriv(%u, %u, %p)\n", target, parameter, values);
        if (glGetTexParameteriv != nullptr) {
            glGetTexParameteriv(target, parameter, values);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetTexParameteriv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetTexParameteriv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetTexParameteriv\n");
        return false;
    }
}

bool callGlIsSampler(Stack* stack, bool pushReturn) {
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsSampler(%" PRIu32 ")\n", sampler);
        if (glIsSampler != nullptr) {
            uint8_t return_value = glIsSampler(sampler);
            GAPID_INFO("Returned: %" PRIu8 "\n", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glIsSampler returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsSampler\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsSampler\n");
        return false;
    }
}

bool callGlIsTexture(Stack* stack, bool pushReturn) {
    uint32_t texture = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsTexture(%" PRIu32 ")\n", texture);
        if (glIsTexture != nullptr) {
            uint8_t return_value = glIsTexture(texture);
            GAPID_INFO("Returned: %" PRIu8 "\n", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glIsTexture returned error: 0x%x\n", err);
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

bool callGlPixelStorei(Stack* stack, bool pushReturn) {
    int32_t value = stack->pop<int32_t>();
    GLenum parameter = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glPixelStorei(%u, %" PRId32 ")\n", parameter, value);
        if (glPixelStorei != nullptr) {
            glPixelStorei(parameter, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glPixelStorei returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPixelStorei\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPixelStorei\n");
        return false;
    }
}

bool callGlSamplerParameterf(Stack* stack, bool pushReturn) {
    float param = stack->pop<float>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glSamplerParameterf(%" PRIu32 ", %u, %f)\n", sampler, pname, param);
        if (glSamplerParameterf != nullptr) {
            glSamplerParameterf(sampler, pname, param);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glSamplerParameterf returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glSamplerParameterf\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glSamplerParameterf\n");
        return false;
    }
}

bool callGlSamplerParameterfv(Stack* stack, bool pushReturn) {
    float* param = stack->pop<float*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glSamplerParameterfv(%" PRIu32 ", %u, %p)\n", sampler, pname, param);
        if (glSamplerParameterfv != nullptr) {
            glSamplerParameterfv(sampler, pname, param);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glSamplerParameterfv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glSamplerParameterfv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glSamplerParameterfv\n");
        return false;
    }
}

bool callGlSamplerParameteri(Stack* stack, bool pushReturn) {
    int32_t param = stack->pop<int32_t>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glSamplerParameteri(%" PRIu32 ", %u, %" PRId32 ")\n", sampler, pname, param);
        if (glSamplerParameteri != nullptr) {
            glSamplerParameteri(sampler, pname, param);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glSamplerParameteri returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glSamplerParameteri\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glSamplerParameteri\n");
        return false;
    }
}

bool callGlSamplerParameteriv(Stack* stack, bool pushReturn) {
    int32_t* param = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t sampler = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glSamplerParameteriv(%" PRIu32 ", %u, %p)\n", sampler, pname, param);
        if (glSamplerParameteriv != nullptr) {
            glSamplerParameteriv(sampler, pname, param);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glSamplerParameteriv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glSamplerParameteriv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glSamplerParameteriv\n");
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
                   ", %u, %u, %p)\n",
                   target, level, internal_format, width, height, border, format, type, data);
        if (glTexImage2D != nullptr) {
            glTexImage2D(target, level, internal_format, width, height, border, format, type, data);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTexImage2D returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexImage2D\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexImage2D\n");
        return false;
    }
}

bool callGlTexImage3D(Stack* stack, bool pushReturn) {
    void* pixels = stack->pop<void*>();
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
                   ", %" PRId32 ", %u, %u, %p)\n",
                   target, level, internalformat, width, height, depth, border, format, type,
                   pixels);
        if (glTexImage3D != nullptr) {
            glTexImage3D(target, level, internalformat, width, height, depth, border, format, type,
                         pixels);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTexImage3D returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexImage3D\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexImage3D\n");
        return false;
    }
}

bool callGlTexParameterf(Stack* stack, bool pushReturn) {
    float value = stack->pop<float>();
    GLenum parameter = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexParameterf(%u, %u, %f)\n", target, parameter, value);
        if (glTexParameterf != nullptr) {
            glTexParameterf(target, parameter, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTexParameterf returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexParameterf\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexParameterf\n");
        return false;
    }
}

bool callGlTexParameterfv(Stack* stack, bool pushReturn) {
    float* params = stack->pop<float*>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexParameterfv(%u, %u, %p)\n", target, pname, params);
        if (glTexParameterfv != nullptr) {
            glTexParameterfv(target, pname, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTexParameterfv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexParameterfv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexParameterfv\n");
        return false;
    }
}

bool callGlTexParameteri(Stack* stack, bool pushReturn) {
    int32_t value = stack->pop<int32_t>();
    GLenum parameter = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexParameteri(%u, %u, %" PRId32 ")\n", target, parameter, value);
        if (glTexParameteri != nullptr) {
            glTexParameteri(target, parameter, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTexParameteri returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexParameteri\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexParameteri\n");
        return false;
    }
}

bool callGlTexParameteriv(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glTexParameteriv(%u, %u, %p)\n", target, pname, params);
        if (glTexParameteriv != nullptr) {
            glTexParameteriv(target, pname, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTexParameteriv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexParameteriv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexParameteriv\n");
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
        GAPID_INFO("glTexStorage2D(%u, %" PRId32 ", %u, %" PRId32 ", %" PRId32 ")\n", target,
                   levels, internalformat, width, height);
        if (glTexStorage2D != nullptr) {
            glTexStorage2D(target, levels, internalformat, width, height);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTexStorage2D returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexStorage2D\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexStorage2D\n");
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
                   ", %" PRIu8 ")\n",
                   target, samples, internalformat, width, height, fixedsamplelocations);
        if (glTexStorage2DMultisample != nullptr) {
            glTexStorage2DMultisample(target, samples, internalformat, width, height,
                                      fixedsamplelocations);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTexStorage2DMultisample returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexStorage2DMultisample\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexStorage2DMultisample\n");
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
        GAPID_INFO("glTexStorage3D(%u, %" PRId32 ", %u, %" PRId32 ", %" PRId32 ", %" PRId32 ")\n",
                   target, levels, internalformat, width, height, depth);
        if (glTexStorage3D != nullptr) {
            glTexStorage3D(target, levels, internalformat, width, height, depth);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTexStorage3D returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexStorage3D\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexStorage3D\n");
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
                   ", %" PRId32 ", %u, %u, %p)\n",
                   target, level, xoffset, yoffset, width, height, format, type, data);
        if (glTexSubImage2D != nullptr) {
            glTexSubImage2D(target, level, xoffset, yoffset, width, height, format, type, data);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTexSubImage2D returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexSubImage2D\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexSubImage2D\n");
        return false;
    }
}

bool callGlTexSubImage3D(Stack* stack, bool pushReturn) {
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
        GAPID_INFO("glTexSubImage3D(%u, %" PRId32 ", %" PRId32 ", %" PRId32 ", %" PRId32
                   ", %" PRId32 ", %" PRId32 ", %" PRId32 ", %u, %u, %p)\n",
                   target, level, xoffset, yoffset, zoffset, width, height, depth, format, type,
                   pixels);
        if (glTexSubImage3D != nullptr) {
            glTexSubImage3D(target, level, xoffset, yoffset, zoffset, width, height, depth, format,
                            type, pixels);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTexSubImage3D returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTexSubImage3D\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTexSubImage3D\n");
        return false;
    }
}

bool callGlBeginTransformFeedback(Stack* stack, bool pushReturn) {
    GLenum primitiveMode = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBeginTransformFeedback(%u)\n", primitiveMode);
        if (glBeginTransformFeedback != nullptr) {
            glBeginTransformFeedback(primitiveMode);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBeginTransformFeedback returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBeginTransformFeedback\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBeginTransformFeedback\n");
        return false;
    }
}

bool callGlBindTransformFeedback(Stack* stack, bool pushReturn) {
    uint32_t id = stack->pop<uint32_t>();
    GLenum target = stack->pop<GLenum>();
    if (stack->isValid()) {
        GAPID_INFO("glBindTransformFeedback(%u, %" PRIu32 ")\n", target, id);
        if (glBindTransformFeedback != nullptr) {
            glBindTransformFeedback(target, id);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBindTransformFeedback returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindTransformFeedback\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindTransformFeedback\n");
        return false;
    }
}

bool callGlDeleteTransformFeedbacks(Stack* stack, bool pushReturn) {
    uint32_t* ids = stack->pop<uint32_t*>();
    int32_t n = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteTransformFeedbacks(%" PRId32 ", %p)\n", n, ids);
        if (glDeleteTransformFeedbacks != nullptr) {
            glDeleteTransformFeedbacks(n, ids);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDeleteTransformFeedbacks returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteTransformFeedbacks\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteTransformFeedbacks\n");
        return false;
    }
}

bool callGlEndTransformFeedback(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glEndTransformFeedback()\n");
        if (glEndTransformFeedback != nullptr) {
            glEndTransformFeedback();
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glEndTransformFeedback returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEndTransformFeedback\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEndTransformFeedback\n");
        return false;
    }
}

bool callGlGenTransformFeedbacks(Stack* stack, bool pushReturn) {
    uint32_t* ids = stack->pop<uint32_t*>();
    int32_t n = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenTransformFeedbacks(%" PRId32 ", %p)\n", n, ids);
        if (glGenTransformFeedbacks != nullptr) {
            glGenTransformFeedbacks(n, ids);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGenTransformFeedbacks returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenTransformFeedbacks\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenTransformFeedbacks\n");
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
                   ", %p, %p, %p, %p)\n",
                   program, index, bufSize, length, size, type, name);
        if (glGetTransformFeedbackVarying != nullptr) {
            glGetTransformFeedbackVarying(program, index, bufSize, length, size, type, name);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetTransformFeedbackVarying returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetTransformFeedbackVarying\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetTransformFeedbackVarying\n");
        return false;
    }
}

bool callGlIsTransformFeedback(Stack* stack, bool pushReturn) {
    uint32_t id = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsTransformFeedback(%" PRIu32 ")\n", id);
        if (glIsTransformFeedback != nullptr) {
            uint8_t return_value = glIsTransformFeedback(id);
            GAPID_INFO("Returned: %" PRIu8 "\n", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glIsTransformFeedback returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsTransformFeedback\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsTransformFeedback\n");
        return false;
    }
}

bool callGlPauseTransformFeedback(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glPauseTransformFeedback()\n");
        if (glPauseTransformFeedback != nullptr) {
            glPauseTransformFeedback();
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glPauseTransformFeedback returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glPauseTransformFeedback\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glPauseTransformFeedback\n");
        return false;
    }
}

bool callGlResumeTransformFeedback(Stack* stack, bool pushReturn) {
    if (stack->isValid()) {
        GAPID_INFO("glResumeTransformFeedback()\n");
        if (glResumeTransformFeedback != nullptr) {
            glResumeTransformFeedback();
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glResumeTransformFeedback returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glResumeTransformFeedback\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glResumeTransformFeedback\n");
        return false;
    }
}

bool callGlTransformFeedbackVaryings(Stack* stack, bool pushReturn) {
    GLenum bufferMode = stack->pop<GLenum>();
    char** varyings = stack->pop<char**>();
    int32_t count = stack->pop<int32_t>();
    uint32_t program = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glTransformFeedbackVaryings(%" PRIu32 ", %" PRId32 ", %p, %u)\n", program,
                   count, varyings, bufferMode);
        if (glTransformFeedbackVaryings != nullptr) {
            glTransformFeedbackVaryings(program, count, varyings, bufferMode);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glTransformFeedbackVaryings returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glTransformFeedbackVaryings\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glTransformFeedbackVaryings\n");
        return false;
    }
}

bool callGlBindVertexArray(Stack* stack, bool pushReturn) {
    uint32_t array = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBindVertexArray(%" PRIu32 ")\n", array);
        if (glBindVertexArray != nullptr) {
            glBindVertexArray(array);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBindVertexArray returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindVertexArray\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindVertexArray\n");
        return false;
    }
}

bool callGlBindVertexBuffer(Stack* stack, bool pushReturn) {
    int32_t stride = stack->pop<int32_t>();
    int32_t offset = stack->pop<int32_t>();
    uint32_t buffer = stack->pop<uint32_t>();
    uint32_t bindingindex = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glBindVertexBuffer(%" PRIu32 ", %" PRIu32 ", %" PRId32 ", %" PRId32 ")\n",
                   bindingindex, buffer, offset, stride);
        if (glBindVertexBuffer != nullptr) {
            glBindVertexBuffer(bindingindex, buffer, offset, stride);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glBindVertexBuffer returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glBindVertexBuffer\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glBindVertexBuffer\n");
        return false;
    }
}

bool callGlDeleteVertexArrays(Stack* stack, bool pushReturn) {
    uint32_t* arrays = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDeleteVertexArrays(%" PRId32 ", %p)\n", count, arrays);
        if (glDeleteVertexArrays != nullptr) {
            glDeleteVertexArrays(count, arrays);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDeleteVertexArrays returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDeleteVertexArrays\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDeleteVertexArrays\n");
        return false;
    }
}

bool callGlDisableVertexAttribArray(Stack* stack, bool pushReturn) {
    uint32_t location = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glDisableVertexAttribArray(%" PRIu32 ")\n", location);
        if (glDisableVertexAttribArray != nullptr) {
            glDisableVertexAttribArray(location);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glDisableVertexAttribArray returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glDisableVertexAttribArray\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glDisableVertexAttribArray\n");
        return false;
    }
}

bool callGlEnableVertexAttribArray(Stack* stack, bool pushReturn) {
    uint32_t location = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glEnableVertexAttribArray(%" PRIu32 ")\n", location);
        if (glEnableVertexAttribArray != nullptr) {
            glEnableVertexAttribArray(location);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glEnableVertexAttribArray returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glEnableVertexAttribArray\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glEnableVertexAttribArray\n");
        return false;
    }
}

bool callGlGenVertexArrays(Stack* stack, bool pushReturn) {
    uint32_t* arrays = stack->pop<uint32_t*>();
    int32_t count = stack->pop<int32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGenVertexArrays(%" PRId32 ", %p)\n", count, arrays);
        if (glGenVertexArrays != nullptr) {
            glGenVertexArrays(count, arrays);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGenVertexArrays returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGenVertexArrays\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGenVertexArrays\n");
        return false;
    }
}

bool callGlGetVertexAttribIiv(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetVertexAttribIiv(%" PRIu32 ", %u, %p)\n", index, pname, params);
        if (glGetVertexAttribIiv != nullptr) {
            glGetVertexAttribIiv(index, pname, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetVertexAttribIiv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetVertexAttribIiv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetVertexAttribIiv\n");
        return false;
    }
}

bool callGlGetVertexAttribIuiv(Stack* stack, bool pushReturn) {
    uint32_t* params = stack->pop<uint32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetVertexAttribIuiv(%" PRIu32 ", %u, %p)\n", index, pname, params);
        if (glGetVertexAttribIuiv != nullptr) {
            glGetVertexAttribIuiv(index, pname, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetVertexAttribIuiv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetVertexAttribIuiv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetVertexAttribIuiv\n");
        return false;
    }
}

bool callGlGetVertexAttribPointerv(Stack* stack, bool pushReturn) {
    void** pointer = stack->pop<void**>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetVertexAttribPointerv(%" PRIu32 ", %u, %p)\n", index, pname, pointer);
        if (glGetVertexAttribPointerv != nullptr) {
            glGetVertexAttribPointerv(index, pname, pointer);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetVertexAttribPointerv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetVertexAttribPointerv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetVertexAttribPointerv\n");
        return false;
    }
}

bool callGlGetVertexAttribfv(Stack* stack, bool pushReturn) {
    float* params = stack->pop<float*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetVertexAttribfv(%" PRIu32 ", %u, %p)\n", index, pname, params);
        if (glGetVertexAttribfv != nullptr) {
            glGetVertexAttribfv(index, pname, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetVertexAttribfv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetVertexAttribfv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetVertexAttribfv\n");
        return false;
    }
}

bool callGlGetVertexAttribiv(Stack* stack, bool pushReturn) {
    int32_t* params = stack->pop<int32_t*>();
    GLenum pname = stack->pop<GLenum>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetVertexAttribiv(%" PRIu32 ", %u, %p)\n", index, pname, params);
        if (glGetVertexAttribiv != nullptr) {
            glGetVertexAttribiv(index, pname, params);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetVertexAttribiv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetVertexAttribiv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetVertexAttribiv\n");
        return false;
    }
}

bool callGlIsVertexArray(Stack* stack, bool pushReturn) {
    uint32_t array = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glIsVertexArray(%" PRIu32 ")\n", array);
        if (glIsVertexArray != nullptr) {
            uint8_t return_value = glIsVertexArray(array);
            GAPID_INFO("Returned: %" PRIu8 "\n", return_value);
            if (pushReturn) {
                stack->push<uint8_t>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glIsVertexArray returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glIsVertexArray\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glIsVertexArray\n");
        return false;
    }
}

bool callGlVertexAttrib1f(Stack* stack, bool pushReturn) {
    float value0 = stack->pop<float>();
    uint32_t location = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttrib1f(%" PRIu32 ", %f)\n", location, value0);
        if (glVertexAttrib1f != nullptr) {
            glVertexAttrib1f(location, value0);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glVertexAttrib1f returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttrib1f\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttrib1f\n");
        return false;
    }
}

bool callGlVertexAttrib1fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint32_t location = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttrib1fv(%" PRIu32 ", %p)\n", location, value);
        if (glVertexAttrib1fv != nullptr) {
            glVertexAttrib1fv(location, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glVertexAttrib1fv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttrib1fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttrib1fv\n");
        return false;
    }
}

bool callGlVertexAttrib2f(Stack* stack, bool pushReturn) {
    float value1 = stack->pop<float>();
    float value0 = stack->pop<float>();
    uint32_t location = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttrib2f(%" PRIu32 ", %f, %f)\n", location, value0, value1);
        if (glVertexAttrib2f != nullptr) {
            glVertexAttrib2f(location, value0, value1);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glVertexAttrib2f returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttrib2f\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttrib2f\n");
        return false;
    }
}

bool callGlVertexAttrib2fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint32_t location = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttrib2fv(%" PRIu32 ", %p)\n", location, value);
        if (glVertexAttrib2fv != nullptr) {
            glVertexAttrib2fv(location, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glVertexAttrib2fv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttrib2fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttrib2fv\n");
        return false;
    }
}

bool callGlVertexAttrib3f(Stack* stack, bool pushReturn) {
    float value2 = stack->pop<float>();
    float value1 = stack->pop<float>();
    float value0 = stack->pop<float>();
    uint32_t location = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttrib3f(%" PRIu32 ", %f, %f, %f)\n", location, value0, value1, value2);
        if (glVertexAttrib3f != nullptr) {
            glVertexAttrib3f(location, value0, value1, value2);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glVertexAttrib3f returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttrib3f\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttrib3f\n");
        return false;
    }
}

bool callGlVertexAttrib3fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint32_t location = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttrib3fv(%" PRIu32 ", %p)\n", location, value);
        if (glVertexAttrib3fv != nullptr) {
            glVertexAttrib3fv(location, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glVertexAttrib3fv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttrib3fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttrib3fv\n");
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
        GAPID_INFO("glVertexAttrib4f(%" PRIu32 ", %f, %f, %f, %f)\n", location, value0, value1,
                   value2, value3);
        if (glVertexAttrib4f != nullptr) {
            glVertexAttrib4f(location, value0, value1, value2, value3);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glVertexAttrib4f returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttrib4f\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttrib4f\n");
        return false;
    }
}

bool callGlVertexAttrib4fv(Stack* stack, bool pushReturn) {
    float* value = stack->pop<float*>();
    uint32_t location = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttrib4fv(%" PRIu32 ", %p)\n", location, value);
        if (glVertexAttrib4fv != nullptr) {
            glVertexAttrib4fv(location, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glVertexAttrib4fv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttrib4fv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttrib4fv\n");
        return false;
    }
}

bool callGlVertexAttribBinding(Stack* stack, bool pushReturn) {
    uint32_t bindingindex = stack->pop<uint32_t>();
    uint32_t attribindex = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttribBinding(%" PRIu32 ", %" PRIu32 ")\n", attribindex, bindingindex);
        if (glVertexAttribBinding != nullptr) {
            glVertexAttribBinding(attribindex, bindingindex);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glVertexAttribBinding returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttribBinding\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttribBinding\n");
        return false;
    }
}

bool callGlVertexAttribDivisor(Stack* stack, bool pushReturn) {
    uint32_t divisor = stack->pop<uint32_t>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttribDivisor(%" PRIu32 ", %" PRIu32 ")\n", index, divisor);
        if (glVertexAttribDivisor != nullptr) {
            glVertexAttribDivisor(index, divisor);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glVertexAttribDivisor returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttribDivisor\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttribDivisor\n");
        return false;
    }
}

bool callGlVertexAttribFormat(Stack* stack, bool pushReturn) {
    uint32_t relativeoffset = stack->pop<uint32_t>();
    uint8_t normalized = stack->pop<uint8_t>();
    GLenum type = stack->pop<GLenum>();
    int32_t size = stack->pop<int32_t>();
    uint32_t attribindex = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttribFormat(%" PRIu32 ", %" PRId32 ", %u, %" PRIu8 ", %" PRIu32 ")\n",
                   attribindex, size, type, normalized, relativeoffset);
        if (glVertexAttribFormat != nullptr) {
            glVertexAttribFormat(attribindex, size, type, normalized, relativeoffset);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glVertexAttribFormat returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttribFormat\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttribFormat\n");
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
                   ")\n",
                   index, x, y, z, w);
        if (glVertexAttribI4i != nullptr) {
            glVertexAttribI4i(index, x, y, z, w);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glVertexAttribI4i returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttribI4i\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttribI4i\n");
        return false;
    }
}

bool callGlVertexAttribI4iv(Stack* stack, bool pushReturn) {
    int32_t* v = stack->pop<int32_t*>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttribI4iv(%" PRIu32 ", %p)\n", index, v);
        if (glVertexAttribI4iv != nullptr) {
            glVertexAttribI4iv(index, v);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glVertexAttribI4iv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttribI4iv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttribI4iv\n");
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
                   ")\n",
                   index, x, y, z, w);
        if (glVertexAttribI4ui != nullptr) {
            glVertexAttribI4ui(index, x, y, z, w);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glVertexAttribI4ui returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttribI4ui\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttribI4ui\n");
        return false;
    }
}

bool callGlVertexAttribI4uiv(Stack* stack, bool pushReturn) {
    uint32_t* v = stack->pop<uint32_t*>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttribI4uiv(%" PRIu32 ", %p)\n", index, v);
        if (glVertexAttribI4uiv != nullptr) {
            glVertexAttribI4uiv(index, v);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glVertexAttribI4uiv returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttribI4uiv\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttribI4uiv\n");
        return false;
    }
}

bool callGlVertexAttribIFormat(Stack* stack, bool pushReturn) {
    uint32_t relativeoffset = stack->pop<uint32_t>();
    GLenum type = stack->pop<GLenum>();
    int32_t size = stack->pop<int32_t>();
    uint32_t attribindex = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttribIFormat(%" PRIu32 ", %" PRId32 ", %u, %" PRIu32 ")\n",
                   attribindex, size, type, relativeoffset);
        if (glVertexAttribIFormat != nullptr) {
            glVertexAttribIFormat(attribindex, size, type, relativeoffset);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glVertexAttribIFormat returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttribIFormat\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttribIFormat\n");
        return false;
    }
}

bool callGlVertexAttribIPointer(Stack* stack, bool pushReturn) {
    void* pointer = stack->pop<void*>();
    int32_t stride = stack->pop<int32_t>();
    GLenum type = stack->pop<GLenum>();
    int32_t size = stack->pop<int32_t>();
    uint32_t index = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexAttribIPointer(%" PRIu32 ", %" PRId32 ", %u, %" PRId32 ", %p)\n", index,
                   size, type, stride, pointer);
        if (glVertexAttribIPointer != nullptr) {
            glVertexAttribIPointer(index, size, type, stride, pointer);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glVertexAttribIPointer returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttribIPointer\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttribIPointer\n");
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
                   ", %p)\n",
                   location, size, type, normalized, stride, data);
        if (glVertexAttribPointer != nullptr) {
            glVertexAttribPointer(location, size, type, normalized, stride, data);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glVertexAttribPointer returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexAttribPointer\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexAttribPointer\n");
        return false;
    }
}

bool callGlVertexBindingDivisor(Stack* stack, bool pushReturn) {
    uint32_t divisor = stack->pop<uint32_t>();
    uint32_t bindingindex = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glVertexBindingDivisor(%" PRIu32 ", %" PRIu32 ")\n", bindingindex, divisor);
        if (glVertexBindingDivisor != nullptr) {
            glVertexBindingDivisor(bindingindex, divisor);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glVertexBindingDivisor returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glVertexBindingDivisor\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glVertexBindingDivisor\n");
        return false;
    }
}

bool callEglInitialize(Stack* stack, bool pushReturn) {
    int* minor = stack->pop<int*>();
    int* major = stack->pop<int*>();
    void* dpy = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("eglInitialize(%p, %p, %p)\n", dpy, major, minor);
        if (eglInitialize != nullptr) {
            int return_value = eglInitialize(dpy, major, minor);
            GAPID_INFO("Returned: %d\n", return_value);
            if (pushReturn) {
                stack->push<int>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("eglInitialize returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function eglInitialize\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function eglInitialize\n");
        return false;
    }
}

bool callEglCreateContext(Stack* stack, bool pushReturn) {
    int* attrib_list = stack->pop<int*>();
    void* share_context = stack->pop<void*>();
    void* config = stack->pop<void*>();
    void* display = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("eglCreateContext(%p, %p, %p, %p)\n", display, config, share_context,
                   attrib_list);
        if (eglCreateContext != nullptr) {
            void* return_value = eglCreateContext(display, config, share_context, attrib_list);
            GAPID_INFO("Returned: %p\n", return_value);
            if (pushReturn) {
                stack->push<void*>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("eglCreateContext returned error: 0x%x\n", err);
            }
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
    void* context = stack->pop<void*>();
    void* read = stack->pop<void*>();
    void* draw = stack->pop<void*>();
    void* display = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("eglMakeCurrent(%p, %p, %p, %p)\n", display, draw, read, context);
        if (eglMakeCurrent != nullptr) {
            int return_value = eglMakeCurrent(display, draw, read, context);
            GAPID_INFO("Returned: %d\n", return_value);
            if (pushReturn) {
                stack->push<int>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("eglMakeCurrent returned error: 0x%x\n", err);
            }
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
    void* surface = stack->pop<void*>();
    void* display = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("eglSwapBuffers(%p, %p)\n", display, surface);
        if (eglSwapBuffers != nullptr) {
            int return_value = eglSwapBuffers(display, surface);
            GAPID_INFO("Returned: %d\n", return_value);
            if (pushReturn) {
                stack->push<int>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("eglSwapBuffers returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function eglSwapBuffers\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function eglSwapBuffers\n");
        return false;
    }
}

bool callEglQuerySurface(Stack* stack, bool pushReturn) {
    int* value = stack->pop<int*>();
    int attribute = stack->pop<int>();
    void* surface = stack->pop<void*>();
    void* display = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("eglQuerySurface(%p, %p, %d, %p)\n", display, surface, attribute, value);
        if (eglQuerySurface != nullptr) {
            int return_value = eglQuerySurface(display, surface, attribute, value);
            GAPID_INFO("Returned: %d\n", return_value);
            if (pushReturn) {
                stack->push<int>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("eglQuerySurface returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function eglQuerySurface\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function eglQuerySurface\n");
        return false;
    }
}

bool callGlXCreateContext(Stack* stack, bool pushReturn) {
    bool direct = stack->pop<bool>();
    void* shareList = stack->pop<void*>();
    void* vis = stack->pop<void*>();
    void* dpy = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("glXCreateContext(%p, %p, %p, %d)\n", dpy, vis, shareList, direct);
        if (glXCreateContext != nullptr) {
            void* return_value = glXCreateContext(dpy, vis, shareList, direct);
            GAPID_INFO("Returned: %p\n", return_value);
            if (pushReturn) {
                stack->push<void*>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glXCreateContext returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glXCreateContext\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glXCreateContext\n");
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
        GAPID_INFO("glXCreateNewContext(%p, %p, %" PRIu32 ", %p, %d)\n", display, fbconfig, type,
                   shared, direct);
        if (glXCreateNewContext != nullptr) {
            void* return_value = glXCreateNewContext(display, fbconfig, type, shared, direct);
            GAPID_INFO("Returned: %p\n", return_value);
            if (pushReturn) {
                stack->push<void*>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glXCreateNewContext returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glXCreateNewContext\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glXCreateNewContext\n");
        return false;
    }
}

bool callGlXMakeContextCurrent(Stack* stack, bool pushReturn) {
    void* ctx = stack->pop<void*>();
    void* read = stack->pop<void*>();
    void* draw = stack->pop<void*>();
    void* display = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("glXMakeContextCurrent(%p, %p, %p, %p)\n", display, draw, read, ctx);
        if (glXMakeContextCurrent != nullptr) {
            int return_value = glXMakeContextCurrent(display, draw, read, ctx);
            GAPID_INFO("Returned: %d\n", return_value);
            if (pushReturn) {
                stack->push<int>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glXMakeContextCurrent returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glXMakeContextCurrent\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glXMakeContextCurrent\n");
        return false;
    }
}

bool callGlXMakeCurrent(Stack* stack, bool pushReturn) {
    void* ctx = stack->pop<void*>();
    void* drawable = stack->pop<void*>();
    void* display = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("glXMakeCurrent(%p, %p, %p)\n", display, drawable, ctx);
        if (glXMakeCurrent != nullptr) {
            int return_value = glXMakeCurrent(display, drawable, ctx);
            GAPID_INFO("Returned: %d\n", return_value);
            if (pushReturn) {
                stack->push<int>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glXMakeCurrent returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glXMakeCurrent\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glXMakeCurrent\n");
        return false;
    }
}

bool callGlXSwapBuffers(Stack* stack, bool pushReturn) {
    void* drawable = stack->pop<void*>();
    void* display = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("glXSwapBuffers(%p, %p)\n", display, drawable);
        if (glXSwapBuffers != nullptr) {
            glXSwapBuffers(display, drawable);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glXSwapBuffers returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glXSwapBuffers\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glXSwapBuffers\n");
        return false;
    }
}

bool callGlXQueryDrawable(Stack* stack, bool pushReturn) {
    int* value = stack->pop<int*>();
    int attribute = stack->pop<int>();
    void* draw = stack->pop<void*>();
    void* display = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("glXQueryDrawable(%p, %p, %d, %p)\n", display, draw, attribute, value);
        if (glXQueryDrawable != nullptr) {
            int return_value = glXQueryDrawable(display, draw, attribute, value);
            GAPID_INFO("Returned: %d\n", return_value);
            if (pushReturn) {
                stack->push<int>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glXQueryDrawable returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glXQueryDrawable\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glXQueryDrawable\n");
        return false;
    }
}

bool callWglCreateContext(Stack* stack, bool pushReturn) {
    void* hdc = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("wglCreateContext(%p)\n", hdc);
        if (wglCreateContext != nullptr) {
            void* return_value = wglCreateContext(hdc);
            GAPID_INFO("Returned: %p\n", return_value);
            if (pushReturn) {
                stack->push<void*>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("wglCreateContext returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function wglCreateContext\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function wglCreateContext\n");
        return false;
    }
}

bool callWglCreateContextAttribsARB(Stack* stack, bool pushReturn) {
    int* attribList = stack->pop<int*>();
    void* hShareContext = stack->pop<void*>();
    void* hdc = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("wglCreateContextAttribsARB(%p, %p, %p)\n", hdc, hShareContext, attribList);
        if (wglCreateContextAttribsARB != nullptr) {
            void* return_value = wglCreateContextAttribsARB(hdc, hShareContext, attribList);
            GAPID_INFO("Returned: %p\n", return_value);
            if (pushReturn) {
                stack->push<void*>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("wglCreateContextAttribsARB returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function wglCreateContextAttribsARB\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function wglCreateContextAttribsARB\n");
        return false;
    }
}

bool callWglMakeCurrent(Stack* stack, bool pushReturn) {
    void* hglrc = stack->pop<void*>();
    void* hdc = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("wglMakeCurrent(%p, %p)\n", hdc, hglrc);
        if (wglMakeCurrent != nullptr) {
            int return_value = wglMakeCurrent(hdc, hglrc);
            GAPID_INFO("Returned: %d\n", return_value);
            if (pushReturn) {
                stack->push<int>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("wglMakeCurrent returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function wglMakeCurrent\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function wglMakeCurrent\n");
        return false;
    }
}

bool callWglSwapBuffers(Stack* stack, bool pushReturn) {
    void* hdc = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("wglSwapBuffers(%p)\n", hdc);
        if (wglSwapBuffers != nullptr) {
            wglSwapBuffers(hdc);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("wglSwapBuffers returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function wglSwapBuffers\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function wglSwapBuffers\n");
        return false;
    }
}

bool callCGLCreateContext(Stack* stack, bool pushReturn) {
    void** ctx = stack->pop<void**>();
    void* share = stack->pop<void*>();
    void* pix = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("CGLCreateContext(%p, %p, %p)\n", pix, share, ctx);
        if (CGLCreateContext != nullptr) {
            int return_value = CGLCreateContext(pix, share, ctx);
            GAPID_INFO("Returned: %d\n", return_value);
            if (pushReturn) {
                stack->push<int>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("CGLCreateContext returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function CGLCreateContext\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function CGLCreateContext\n");
        return false;
    }
}

bool callCGLSetCurrentContext(Stack* stack, bool pushReturn) {
    void* ctx = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("CGLSetCurrentContext(%p)\n", ctx);
        if (CGLSetCurrentContext != nullptr) {
            int return_value = CGLSetCurrentContext(ctx);
            GAPID_INFO("Returned: %d\n", return_value);
            if (pushReturn) {
                stack->push<int>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("CGLSetCurrentContext returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function CGLSetCurrentContext\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function CGLSetCurrentContext\n");
        return false;
    }
}

bool callCGLGetSurface(Stack* stack, bool pushReturn) {
    int32_t* sid = stack->pop<int32_t*>();
    int32_t* wid = stack->pop<int32_t*>();
    void** cid = stack->pop<void**>();
    void* ctx = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("CGLGetSurface(%p, %p, %p, %p)\n", ctx, cid, wid, sid);
        if (CGLGetSurface != nullptr) {
            int return_value = CGLGetSurface(ctx, cid, wid, sid);
            GAPID_INFO("Returned: %d\n", return_value);
            if (pushReturn) {
                stack->push<int>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("CGLGetSurface returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function CGLGetSurface\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function CGLGetSurface\n");
        return false;
    }
}

bool callCGSGetSurfaceBounds(Stack* stack, bool pushReturn) {
    double* bounds = stack->pop<double*>();
    int32_t sid = stack->pop<int32_t>();
    int32_t wid = stack->pop<int32_t>();
    void* cid = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("CGSGetSurfaceBounds(%p, %" PRId32 ", %" PRId32 ", %p)\n", cid, wid, sid,
                   bounds);
        if (CGSGetSurfaceBounds != nullptr) {
            int return_value = CGSGetSurfaceBounds(cid, wid, sid, bounds);
            GAPID_INFO("Returned: %d\n", return_value);
            if (pushReturn) {
                stack->push<int>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("CGSGetSurfaceBounds returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function CGSGetSurfaceBounds\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function CGSGetSurfaceBounds\n");
        return false;
    }
}

bool callCGLFlushDrawable(Stack* stack, bool pushReturn) {
    void* ctx = stack->pop<void*>();
    if (stack->isValid()) {
        GAPID_INFO("CGLFlushDrawable(%p)\n", ctx);
        if (CGLFlushDrawable != nullptr) {
            int return_value = CGLFlushDrawable(ctx);
            GAPID_INFO("Returned: %d\n", return_value);
            if (pushReturn) {
                stack->push<int>(return_value);
            }
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("CGLFlushDrawable returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function CGLFlushDrawable\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function CGLFlushDrawable\n");
        return false;
    }
}

bool callGlGetQueryObjecti64v(Stack* stack, bool pushReturn) {
    int64_t* value = stack->pop<int64_t*>();
    GLenum parameter = stack->pop<GLenum>();
    uint32_t query = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetQueryObjecti64v(%" PRIu32 ", %u, %p)\n", query, parameter, value);
        if (glGetQueryObjecti64v != nullptr) {
            glGetQueryObjecti64v(query, parameter, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetQueryObjecti64v returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetQueryObjecti64v\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetQueryObjecti64v\n");
        return false;
    }
}

bool callGlGetQueryObjectui64v(Stack* stack, bool pushReturn) {
    uint64_t* value = stack->pop<uint64_t*>();
    GLenum parameter = stack->pop<GLenum>();
    uint32_t query = stack->pop<uint32_t>();
    if (stack->isValid()) {
        GAPID_INFO("glGetQueryObjectui64v(%" PRIu32 ", %u, %p)\n", query, parameter, value);
        if (glGetQueryObjectui64v != nullptr) {
            glGetQueryObjectui64v(query, parameter, value);
            const GLenum err = glGetError();
            if (err != GLenum::GL_NO_ERROR) {
                GAPID_WARNING("glGetQueryObjectui64v returned error: 0x%x\n", err);
            }
        } else {
            GAPID_WARNING("Attempted to call unsupported function glGetQueryObjectui64v\n");
        }
        return true;
    } else {
        GAPID_WARNING("Error during calling function glGetQueryObjectui64v\n");
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
PFNGLDRAWARRAYS glDrawArrays = nullptr;
PFNGLDRAWARRAYSINDIRECT glDrawArraysIndirect = nullptr;
PFNGLDRAWARRAYSINSTANCED glDrawArraysInstanced = nullptr;
PFNGLDRAWBUFFERS glDrawBuffers = nullptr;
PFNGLDRAWELEMENTS glDrawElements = nullptr;
PFNGLDRAWELEMENTSINDIRECT glDrawElementsIndirect = nullptr;
PFNGLDRAWELEMENTSINSTANCED glDrawElementsInstanced = nullptr;
PFNGLDRAWRANGEELEMENTS glDrawRangeElements = nullptr;
PFNGLACTIVESHADERPROGRAMEXT glActiveShaderProgramEXT = nullptr;
PFNGLALPHAFUNCQCOM glAlphaFuncQCOM = nullptr;
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
PFNGLDRAWELEMENTSINSTANCEDBASEVERTEXEXT glDrawElementsInstancedBaseVertexEXT = nullptr;
PFNGLDRAWELEMENTSINSTANCEDBASEVERTEXOES glDrawElementsInstancedBaseVertexOES = nullptr;
PFNGLDRAWELEMENTSINSTANCEDBASEVERTEXBASEINSTANCEEXT
        glDrawElementsInstancedBaseVertexBaseInstanceEXT = nullptr;
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
PFNGLFRAMEBUFFERTEXTURE2DMULTISAMPLEEXT glFramebufferTexture2DMultisampleEXT = nullptr;
PFNGLFRAMEBUFFERTEXTURE2DMULTISAMPLEIMG glFramebufferTexture2DMultisampleIMG = nullptr;
PFNGLFRAMEBUFFERTEXTURE3DOES glFramebufferTexture3DOES = nullptr;
PFNGLFRAMEBUFFERTEXTUREOES glFramebufferTextureOES = nullptr;
PFNGLFRAMEBUFFERTEXTUREMULTIVIEWOVR glFramebufferTextureMultiviewOVR = nullptr;
PFNGLGENFENCESNV glGenFencesNV = nullptr;
PFNGLGENPATHSNV glGenPathsNV = nullptr;
PFNGLGENPERFMONITORSAMD glGenPerfMonitorsAMD = nullptr;
PFNGLGENPROGRAMPIPELINESEXT glGenProgramPipelinesEXT = nullptr;
PFNGLGENQUERIESEXT glGenQueriesEXT = nullptr;
PFNGLGENVERTEXARRAYSOES glGenVertexArraysOES = nullptr;
PFNGLGETBUFFERPOINTERVOES glGetBufferPointervOES = nullptr;
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
PFNGLISENABLEDIOES glIsEnablediOES = nullptr;
PFNGLISENABLEDINV glIsEnablediNV = nullptr;
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
PFNGLREADBUFFERINDEXEDEXT glReadBufferIndexedEXT = nullptr;
PFNGLREADBUFFERNV glReadBufferNV = nullptr;
PFNGLREADNPIXELSEXT glReadnPixelsEXT = nullptr;
PFNGLREADNPIXELSKHR glReadnPixelsKHR = nullptr;
PFNGLRENDERBUFFERSTORAGEMULTISAMPLEANGLE glRenderbufferStorageMultisampleANGLE = nullptr;
PFNGLRENDERBUFFERSTORAGEMULTISAMPLEAPPLE glRenderbufferStorageMultisampleAPPLE = nullptr;
PFNGLRENDERBUFFERSTORAGEMULTISAMPLEEXT glRenderbufferStorageMultisampleEXT = nullptr;
PFNGLRENDERBUFFERSTORAGEMULTISAMPLEIMG glRenderbufferStorageMultisampleIMG = nullptr;
PFNGLRENDERBUFFERSTORAGEMULTISAMPLENV glRenderbufferStorageMultisampleNV = nullptr;
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
PFNGLCOVERAGEMODULATIONNV glCoverageModulationNV = nullptr;
PFNGLCOVERAGEMODULATIONTABLENV glCoverageModulationTableNV = nullptr;
PFNGLFRAGMENTCOVERAGECOLORNV glFragmentCoverageColorNV = nullptr;
PFNGLFRAMEBUFFERSAMPLELOCATIONSFVNV glFramebufferSampleLocationsfvNV = nullptr;
PFNGLGETCOVERAGEMODULATIONTABLENV glGetCoverageModulationTableNV = nullptr;
PFNGLNAMEDFRAMEBUFFERSAMPLELOCATIONSFVNV glNamedFramebufferSampleLocationsfvNV = nullptr;
PFNGLRASTERSAMPLESEXT glRasterSamplesEXT = nullptr;
PFNGLRESOLVEDEPTHVALUESNV glResolveDepthValuesNV = nullptr;
PFNGLSUBPIXELPRECISIONBIASNV glSubpixelPrecisionBiasNV = nullptr;
PFNGLBLENDCOLOR glBlendColor = nullptr;
PFNGLBLENDEQUATION glBlendEquation = nullptr;
PFNGLBLENDEQUATIONSEPARATE glBlendEquationSeparate = nullptr;
PFNGLBLENDFUNC glBlendFunc = nullptr;
PFNGLBLENDFUNCSEPARATE glBlendFuncSeparate = nullptr;
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
PFNGLDELETEFRAMEBUFFERS glDeleteFramebuffers = nullptr;
PFNGLDELETERENDERBUFFERS glDeleteRenderbuffers = nullptr;
PFNGLDEPTHMASK glDepthMask = nullptr;
PFNGLFRAMEBUFFERPARAMETERI glFramebufferParameteri = nullptr;
PFNGLFRAMEBUFFERRENDERBUFFER glFramebufferRenderbuffer = nullptr;
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
PFNGLRENDERBUFFERSTORAGE glRenderbufferStorage = nullptr;
PFNGLRENDERBUFFERSTORAGEMULTISAMPLE glRenderbufferStorageMultisample = nullptr;
PFNGLSTENCILMASK glStencilMask = nullptr;
PFNGLSTENCILMASKSEPARATE glStencilMaskSeparate = nullptr;
PFNGLDISABLE glDisable = nullptr;
PFNGLENABLE glEnable = nullptr;
PFNGLFINISH glFinish = nullptr;
PFNGLFLUSH glFlush = nullptr;
PFNGLFLUSHMAPPEDBUFFERRANGE glFlushMappedBufferRange = nullptr;
PFNGLGETERROR glGetError = nullptr;
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
PFNGLCOPYTEXIMAGE2D glCopyTexImage2D = nullptr;
PFNGLCOPYTEXSUBIMAGE2D glCopyTexSubImage2D = nullptr;
PFNGLCOPYTEXSUBIMAGE3D glCopyTexSubImage3D = nullptr;
PFNGLDELETESAMPLERS glDeleteSamplers = nullptr;
PFNGLDELETETEXTURES glDeleteTextures = nullptr;
PFNGLGENSAMPLERS glGenSamplers = nullptr;
PFNGLGENTEXTURES glGenTextures = nullptr;
PFNGLGENERATEMIPMAP glGenerateMipmap = nullptr;
PFNGLGETSAMPLERPARAMETERFV glGetSamplerParameterfv = nullptr;
PFNGLGETSAMPLERPARAMETERIV glGetSamplerParameteriv = nullptr;
PFNGLGETTEXLEVELPARAMETERFV glGetTexLevelParameterfv = nullptr;
PFNGLGETTEXLEVELPARAMETERIV glGetTexLevelParameteriv = nullptr;
PFNGLGETTEXPARAMETERFV glGetTexParameterfv = nullptr;
PFNGLGETTEXPARAMETERIV glGetTexParameteriv = nullptr;
PFNGLISSAMPLER glIsSampler = nullptr;
PFNGLISTEXTURE glIsTexture = nullptr;
PFNGLPIXELSTOREI glPixelStorei = nullptr;
PFNGLSAMPLERPARAMETERF glSamplerParameterf = nullptr;
PFNGLSAMPLERPARAMETERFV glSamplerParameterfv = nullptr;
PFNGLSAMPLERPARAMETERI glSamplerParameteri = nullptr;
PFNGLSAMPLERPARAMETERIV glSamplerParameteriv = nullptr;
PFNGLTEXIMAGE2D glTexImage2D = nullptr;
PFNGLTEXIMAGE3D glTexImage3D = nullptr;
PFNGLTEXPARAMETERF glTexParameterf = nullptr;
PFNGLTEXPARAMETERFV glTexParameterfv = nullptr;
PFNGLTEXPARAMETERI glTexParameteri = nullptr;
PFNGLTEXPARAMETERIV glTexParameteriv = nullptr;
PFNGLTEXSTORAGE2D glTexStorage2D = nullptr;
PFNGLTEXSTORAGE2DMULTISAMPLE glTexStorage2DMultisample = nullptr;
PFNGLTEXSTORAGE3D glTexStorage3D = nullptr;
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
    interpreter->registerFunction(Ids::GlDrawArrays, callGlDrawArrays);
    interpreter->registerFunction(Ids::GlDrawArraysIndirect, callGlDrawArraysIndirect);
    interpreter->registerFunction(Ids::GlDrawArraysInstanced, callGlDrawArraysInstanced);
    interpreter->registerFunction(Ids::GlDrawBuffers, callGlDrawBuffers);
    interpreter->registerFunction(Ids::GlDrawElements, callGlDrawElements);
    interpreter->registerFunction(Ids::GlDrawElementsIndirect, callGlDrawElementsIndirect);
    interpreter->registerFunction(Ids::GlDrawElementsInstanced, callGlDrawElementsInstanced);
    interpreter->registerFunction(Ids::GlDrawRangeElements, callGlDrawRangeElements);
    interpreter->registerFunction(Ids::GlActiveShaderProgramEXT, callGlActiveShaderProgramEXT);
    interpreter->registerFunction(Ids::GlAlphaFuncQCOM, callGlAlphaFuncQCOM);
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
    interpreter->registerFunction(Ids::GlDrawElementsInstancedBaseVertexEXT,
                                  callGlDrawElementsInstancedBaseVertexEXT);
    interpreter->registerFunction(Ids::GlDrawElementsInstancedBaseVertexOES,
                                  callGlDrawElementsInstancedBaseVertexOES);
    interpreter->registerFunction(Ids::GlDrawElementsInstancedBaseVertexBaseInstanceEXT,
                                  callGlDrawElementsInstancedBaseVertexBaseInstanceEXT);
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
    interpreter->registerFunction(Ids::GlFramebufferTexture2DMultisampleEXT,
                                  callGlFramebufferTexture2DMultisampleEXT);
    interpreter->registerFunction(Ids::GlFramebufferTexture2DMultisampleIMG,
                                  callGlFramebufferTexture2DMultisampleIMG);
    interpreter->registerFunction(Ids::GlFramebufferTexture3DOES, callGlFramebufferTexture3DOES);
    interpreter->registerFunction(Ids::GlFramebufferTextureOES, callGlFramebufferTextureOES);
    interpreter->registerFunction(Ids::GlFramebufferTextureMultiviewOVR,
                                  callGlFramebufferTextureMultiviewOVR);
    interpreter->registerFunction(Ids::GlGenFencesNV, callGlGenFencesNV);
    interpreter->registerFunction(Ids::GlGenPathsNV, callGlGenPathsNV);
    interpreter->registerFunction(Ids::GlGenPerfMonitorsAMD, callGlGenPerfMonitorsAMD);
    interpreter->registerFunction(Ids::GlGenProgramPipelinesEXT, callGlGenProgramPipelinesEXT);
    interpreter->registerFunction(Ids::GlGenQueriesEXT, callGlGenQueriesEXT);
    interpreter->registerFunction(Ids::GlGenVertexArraysOES, callGlGenVertexArraysOES);
    interpreter->registerFunction(Ids::GlGetBufferPointervOES, callGlGetBufferPointervOES);
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
    interpreter->registerFunction(Ids::GlIsEnablediOES, callGlIsEnablediOES);
    interpreter->registerFunction(Ids::GlIsEnablediNV, callGlIsEnablediNV);
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
    interpreter->registerFunction(Ids::GlCoverageModulationNV, callGlCoverageModulationNV);
    interpreter->registerFunction(Ids::GlCoverageModulationTableNV,
                                  callGlCoverageModulationTableNV);
    interpreter->registerFunction(Ids::GlFragmentCoverageColorNV, callGlFragmentCoverageColorNV);
    interpreter->registerFunction(Ids::GlFramebufferSampleLocationsfvNV,
                                  callGlFramebufferSampleLocationsfvNV);
    interpreter->registerFunction(Ids::GlGetCoverageModulationTableNV,
                                  callGlGetCoverageModulationTableNV);
    interpreter->registerFunction(Ids::GlNamedFramebufferSampleLocationsfvNV,
                                  callGlNamedFramebufferSampleLocationsfvNV);
    interpreter->registerFunction(Ids::GlRasterSamplesEXT, callGlRasterSamplesEXT);
    interpreter->registerFunction(Ids::GlResolveDepthValuesNV, callGlResolveDepthValuesNV);
    interpreter->registerFunction(Ids::GlSubpixelPrecisionBiasNV, callGlSubpixelPrecisionBiasNV);
    interpreter->registerFunction(Ids::GlBlendColor, callGlBlendColor);
    interpreter->registerFunction(Ids::GlBlendEquation, callGlBlendEquation);
    interpreter->registerFunction(Ids::GlBlendEquationSeparate, callGlBlendEquationSeparate);
    interpreter->registerFunction(Ids::GlBlendFunc, callGlBlendFunc);
    interpreter->registerFunction(Ids::GlBlendFuncSeparate, callGlBlendFuncSeparate);
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
    interpreter->registerFunction(Ids::GlDeleteFramebuffers, callGlDeleteFramebuffers);
    interpreter->registerFunction(Ids::GlDeleteRenderbuffers, callGlDeleteRenderbuffers);
    interpreter->registerFunction(Ids::GlDepthMask, callGlDepthMask);
    interpreter->registerFunction(Ids::GlFramebufferParameteri, callGlFramebufferParameteri);
    interpreter->registerFunction(Ids::GlFramebufferRenderbuffer, callGlFramebufferRenderbuffer);
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
    interpreter->registerFunction(Ids::GlRenderbufferStorage, callGlRenderbufferStorage);
    interpreter->registerFunction(Ids::GlRenderbufferStorageMultisample,
                                  callGlRenderbufferStorageMultisample);
    interpreter->registerFunction(Ids::GlStencilMask, callGlStencilMask);
    interpreter->registerFunction(Ids::GlStencilMaskSeparate, callGlStencilMaskSeparate);
    interpreter->registerFunction(Ids::GlDisable, callGlDisable);
    interpreter->registerFunction(Ids::GlEnable, callGlEnable);
    interpreter->registerFunction(Ids::GlFinish, callGlFinish);
    interpreter->registerFunction(Ids::GlFlush, callGlFlush);
    interpreter->registerFunction(Ids::GlFlushMappedBufferRange, callGlFlushMappedBufferRange);
    interpreter->registerFunction(Ids::GlGetError, callGlGetError);
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
    interpreter->registerFunction(Ids::GlCopyTexImage2D, callGlCopyTexImage2D);
    interpreter->registerFunction(Ids::GlCopyTexSubImage2D, callGlCopyTexSubImage2D);
    interpreter->registerFunction(Ids::GlCopyTexSubImage3D, callGlCopyTexSubImage3D);
    interpreter->registerFunction(Ids::GlDeleteSamplers, callGlDeleteSamplers);
    interpreter->registerFunction(Ids::GlDeleteTextures, callGlDeleteTextures);
    interpreter->registerFunction(Ids::GlGenSamplers, callGlGenSamplers);
    interpreter->registerFunction(Ids::GlGenTextures, callGlGenTextures);
    interpreter->registerFunction(Ids::GlGenerateMipmap, callGlGenerateMipmap);
    interpreter->registerFunction(Ids::GlGetSamplerParameterfv, callGlGetSamplerParameterfv);
    interpreter->registerFunction(Ids::GlGetSamplerParameteriv, callGlGetSamplerParameteriv);
    interpreter->registerFunction(Ids::GlGetTexLevelParameterfv, callGlGetTexLevelParameterfv);
    interpreter->registerFunction(Ids::GlGetTexLevelParameteriv, callGlGetTexLevelParameteriv);
    interpreter->registerFunction(Ids::GlGetTexParameterfv, callGlGetTexParameterfv);
    interpreter->registerFunction(Ids::GlGetTexParameteriv, callGlGetTexParameteriv);
    interpreter->registerFunction(Ids::GlIsSampler, callGlIsSampler);
    interpreter->registerFunction(Ids::GlIsTexture, callGlIsTexture);
    interpreter->registerFunction(Ids::GlPixelStorei, callGlPixelStorei);
    interpreter->registerFunction(Ids::GlSamplerParameterf, callGlSamplerParameterf);
    interpreter->registerFunction(Ids::GlSamplerParameterfv, callGlSamplerParameterfv);
    interpreter->registerFunction(Ids::GlSamplerParameteri, callGlSamplerParameteri);
    interpreter->registerFunction(Ids::GlSamplerParameteriv, callGlSamplerParameteriv);
    interpreter->registerFunction(Ids::GlTexImage2D, callGlTexImage2D);
    interpreter->registerFunction(Ids::GlTexImage3D, callGlTexImage3D);
    interpreter->registerFunction(Ids::GlTexParameterf, callGlTexParameterf);
    interpreter->registerFunction(Ids::GlTexParameterfv, callGlTexParameterfv);
    interpreter->registerFunction(Ids::GlTexParameteri, callGlTexParameteri);
    interpreter->registerFunction(Ids::GlTexParameteriv, callGlTexParameteriv);
    interpreter->registerFunction(Ids::GlTexStorage2D, callGlTexStorage2D);
    interpreter->registerFunction(Ids::GlTexStorage2DMultisample, callGlTexStorage2DMultisample);
    interpreter->registerFunction(Ids::GlTexStorage3D, callGlTexStorage3D);
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
    glDrawArrays =
            reinterpret_cast<PFNGLDRAWARRAYS>(gapic::GetGfxProcAddress("glDrawArrays", false));
    glDrawArraysIndirect = reinterpret_cast<PFNGLDRAWARRAYSINDIRECT>(
            gapic::GetGfxProcAddress("glDrawArraysIndirect", false));
    glDrawArraysInstanced = reinterpret_cast<PFNGLDRAWARRAYSINSTANCED>(
            gapic::GetGfxProcAddress("glDrawArraysInstanced", false));
    glDrawBuffers =
            reinterpret_cast<PFNGLDRAWBUFFERS>(gapic::GetGfxProcAddress("glDrawBuffers", false));
    glDrawElements =
            reinterpret_cast<PFNGLDRAWELEMENTS>(gapic::GetGfxProcAddress("glDrawElements", false));
    glDrawElementsIndirect = reinterpret_cast<PFNGLDRAWELEMENTSINDIRECT>(
            gapic::GetGfxProcAddress("glDrawElementsIndirect", false));
    glDrawElementsInstanced = reinterpret_cast<PFNGLDRAWELEMENTSINSTANCED>(
            gapic::GetGfxProcAddress("glDrawElementsInstanced", false));
    glDrawRangeElements = reinterpret_cast<PFNGLDRAWRANGEELEMENTS>(
            gapic::GetGfxProcAddress("glDrawRangeElements", false));
    glActiveShaderProgramEXT = reinterpret_cast<PFNGLACTIVESHADERPROGRAMEXT>(
            gapic::GetGfxProcAddress("glActiveShaderProgramEXT", false));
    glAlphaFuncQCOM = reinterpret_cast<PFNGLALPHAFUNCQCOM>(
            gapic::GetGfxProcAddress("glAlphaFuncQCOM", false));
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
    glDrawElementsInstancedBaseVertexEXT =
            reinterpret_cast<PFNGLDRAWELEMENTSINSTANCEDBASEVERTEXEXT>(
                    gapic::GetGfxProcAddress("glDrawElementsInstancedBaseVertexEXT", false));
    glDrawElementsInstancedBaseVertexOES =
            reinterpret_cast<PFNGLDRAWELEMENTSINSTANCEDBASEVERTEXOES>(
                    gapic::GetGfxProcAddress("glDrawElementsInstancedBaseVertexOES", false));
    glDrawElementsInstancedBaseVertexBaseInstanceEXT =
            reinterpret_cast<PFNGLDRAWELEMENTSINSTANCEDBASEVERTEXBASEINSTANCEEXT>(
                    gapic::GetGfxProcAddress("glDrawElementsInstancedBaseVertexBaseInstanceEXT",
                                             false));
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
    glFramebufferTexture2DMultisampleEXT =
            reinterpret_cast<PFNGLFRAMEBUFFERTEXTURE2DMULTISAMPLEEXT>(
                    gapic::GetGfxProcAddress("glFramebufferTexture2DMultisampleEXT", false));
    glFramebufferTexture2DMultisampleIMG =
            reinterpret_cast<PFNGLFRAMEBUFFERTEXTURE2DMULTISAMPLEIMG>(
                    gapic::GetGfxProcAddress("glFramebufferTexture2DMultisampleIMG", false));
    glFramebufferTexture3DOES = reinterpret_cast<PFNGLFRAMEBUFFERTEXTURE3DOES>(
            gapic::GetGfxProcAddress("glFramebufferTexture3DOES", false));
    glFramebufferTextureOES = reinterpret_cast<PFNGLFRAMEBUFFERTEXTUREOES>(
            gapic::GetGfxProcAddress("glFramebufferTextureOES", false));
    glFramebufferTextureMultiviewOVR = reinterpret_cast<PFNGLFRAMEBUFFERTEXTUREMULTIVIEWOVR>(
            gapic::GetGfxProcAddress("glFramebufferTextureMultiviewOVR", false));
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
    glIsEnablediOES = reinterpret_cast<PFNGLISENABLEDIOES>(
            gapic::GetGfxProcAddress("glIsEnablediOES", false));
    glIsEnablediNV =
            reinterpret_cast<PFNGLISENABLEDINV>(gapic::GetGfxProcAddress("glIsEnablediNV", false));
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
    glCoverageModulationNV = reinterpret_cast<PFNGLCOVERAGEMODULATIONNV>(
            gapic::GetGfxProcAddress("glCoverageModulationNV", false));
    glCoverageModulationTableNV = reinterpret_cast<PFNGLCOVERAGEMODULATIONTABLENV>(
            gapic::GetGfxProcAddress("glCoverageModulationTableNV", false));
    glFragmentCoverageColorNV = reinterpret_cast<PFNGLFRAGMENTCOVERAGECOLORNV>(
            gapic::GetGfxProcAddress("glFragmentCoverageColorNV", false));
    glFramebufferSampleLocationsfvNV = reinterpret_cast<PFNGLFRAMEBUFFERSAMPLELOCATIONSFVNV>(
            gapic::GetGfxProcAddress("glFramebufferSampleLocationsfvNV", false));
    glGetCoverageModulationTableNV = reinterpret_cast<PFNGLGETCOVERAGEMODULATIONTABLENV>(
            gapic::GetGfxProcAddress("glGetCoverageModulationTableNV", false));
    glNamedFramebufferSampleLocationsfvNV =
            reinterpret_cast<PFNGLNAMEDFRAMEBUFFERSAMPLELOCATIONSFVNV>(
                    gapic::GetGfxProcAddress("glNamedFramebufferSampleLocationsfvNV", false));
    glRasterSamplesEXT = reinterpret_cast<PFNGLRASTERSAMPLESEXT>(
            gapic::GetGfxProcAddress("glRasterSamplesEXT", false));
    glResolveDepthValuesNV = reinterpret_cast<PFNGLRESOLVEDEPTHVALUESNV>(
            gapic::GetGfxProcAddress("glResolveDepthValuesNV", false));
    glSubpixelPrecisionBiasNV = reinterpret_cast<PFNGLSUBPIXELPRECISIONBIASNV>(
            gapic::GetGfxProcAddress("glSubpixelPrecisionBiasNV", false));
    glBlendColor =
            reinterpret_cast<PFNGLBLENDCOLOR>(gapic::GetGfxProcAddress("glBlendColor", false));
    glBlendEquation = reinterpret_cast<PFNGLBLENDEQUATION>(
            gapic::GetGfxProcAddress("glBlendEquation", false));
    glBlendEquationSeparate = reinterpret_cast<PFNGLBLENDEQUATIONSEPARATE>(
            gapic::GetGfxProcAddress("glBlendEquationSeparate", false));
    glBlendFunc = reinterpret_cast<PFNGLBLENDFUNC>(gapic::GetGfxProcAddress("glBlendFunc", false));
    glBlendFuncSeparate = reinterpret_cast<PFNGLBLENDFUNCSEPARATE>(
            gapic::GetGfxProcAddress("glBlendFuncSeparate", false));
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
    glDeleteFramebuffers = reinterpret_cast<PFNGLDELETEFRAMEBUFFERS>(
            gapic::GetGfxProcAddress("glDeleteFramebuffers", false));
    glDeleteRenderbuffers = reinterpret_cast<PFNGLDELETERENDERBUFFERS>(
            gapic::GetGfxProcAddress("glDeleteRenderbuffers", false));
    glDepthMask = reinterpret_cast<PFNGLDEPTHMASK>(gapic::GetGfxProcAddress("glDepthMask", false));
    glFramebufferParameteri = reinterpret_cast<PFNGLFRAMEBUFFERPARAMETERI>(
            gapic::GetGfxProcAddress("glFramebufferParameteri", false));
    glFramebufferRenderbuffer = reinterpret_cast<PFNGLFRAMEBUFFERRENDERBUFFER>(
            gapic::GetGfxProcAddress("glFramebufferRenderbuffer", false));
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
    glRenderbufferStorage = reinterpret_cast<PFNGLRENDERBUFFERSTORAGE>(
            gapic::GetGfxProcAddress("glRenderbufferStorage", false));
    glRenderbufferStorageMultisample = reinterpret_cast<PFNGLRENDERBUFFERSTORAGEMULTISAMPLE>(
            gapic::GetGfxProcAddress("glRenderbufferStorageMultisample", false));
    glStencilMask =
            reinterpret_cast<PFNGLSTENCILMASK>(gapic::GetGfxProcAddress("glStencilMask", false));
    glStencilMaskSeparate = reinterpret_cast<PFNGLSTENCILMASKSEPARATE>(
            gapic::GetGfxProcAddress("glStencilMaskSeparate", false));
    glDisable = reinterpret_cast<PFNGLDISABLE>(gapic::GetGfxProcAddress("glDisable", false));
    glEnable = reinterpret_cast<PFNGLENABLE>(gapic::GetGfxProcAddress("glEnable", false));
    glFinish = reinterpret_cast<PFNGLFINISH>(gapic::GetGfxProcAddress("glFinish", false));
    glFlush = reinterpret_cast<PFNGLFLUSH>(gapic::GetGfxProcAddress("glFlush", false));
    glFlushMappedBufferRange = reinterpret_cast<PFNGLFLUSHMAPPEDBUFFERRANGE>(
            gapic::GetGfxProcAddress("glFlushMappedBufferRange", false));
    glGetError = reinterpret_cast<PFNGLGETERROR>(gapic::GetGfxProcAddress("glGetError", false));
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
    glGetSamplerParameterfv = reinterpret_cast<PFNGLGETSAMPLERPARAMETERFV>(
            gapic::GetGfxProcAddress("glGetSamplerParameterfv", false));
    glGetSamplerParameteriv = reinterpret_cast<PFNGLGETSAMPLERPARAMETERIV>(
            gapic::GetGfxProcAddress("glGetSamplerParameteriv", false));
    glGetTexLevelParameterfv = reinterpret_cast<PFNGLGETTEXLEVELPARAMETERFV>(
            gapic::GetGfxProcAddress("glGetTexLevelParameterfv", false));
    glGetTexLevelParameteriv = reinterpret_cast<PFNGLGETTEXLEVELPARAMETERIV>(
            gapic::GetGfxProcAddress("glGetTexLevelParameteriv", false));
    glGetTexParameterfv = reinterpret_cast<PFNGLGETTEXPARAMETERFV>(
            gapic::GetGfxProcAddress("glGetTexParameterfv", false));
    glGetTexParameteriv = reinterpret_cast<PFNGLGETTEXPARAMETERIV>(
            gapic::GetGfxProcAddress("glGetTexParameteriv", false));
    glIsSampler = reinterpret_cast<PFNGLISSAMPLER>(gapic::GetGfxProcAddress("glIsSampler", false));
    glIsTexture = reinterpret_cast<PFNGLISTEXTURE>(gapic::GetGfxProcAddress("glIsTexture", false));
    glPixelStorei =
            reinterpret_cast<PFNGLPIXELSTOREI>(gapic::GetGfxProcAddress("glPixelStorei", false));
    glSamplerParameterf = reinterpret_cast<PFNGLSAMPLERPARAMETERF>(
            gapic::GetGfxProcAddress("glSamplerParameterf", false));
    glSamplerParameterfv = reinterpret_cast<PFNGLSAMPLERPARAMETERFV>(
            gapic::GetGfxProcAddress("glSamplerParameterfv", false));
    glSamplerParameteri = reinterpret_cast<PFNGLSAMPLERPARAMETERI>(
            gapic::GetGfxProcAddress("glSamplerParameteri", false));
    glSamplerParameteriv = reinterpret_cast<PFNGLSAMPLERPARAMETERIV>(
            gapic::GetGfxProcAddress("glSamplerParameteriv", false));
    glTexImage2D =
            reinterpret_cast<PFNGLTEXIMAGE2D>(gapic::GetGfxProcAddress("glTexImage2D", false));
    glTexImage3D =
            reinterpret_cast<PFNGLTEXIMAGE3D>(gapic::GetGfxProcAddress("glTexImage3D", false));
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
