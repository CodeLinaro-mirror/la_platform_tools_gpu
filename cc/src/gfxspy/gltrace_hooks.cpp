/*
 * Copyright 2011, The Android Open Source Project
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

#include "glestrace.h"
#include "gltrace_api.h"
#include "gltrace_context.h"
#include "gltrace_egl.h"
#include "hooks.h"
#include "log/log.h"

#undef TRACE_GL_VOID
#undef TRACE_GL
#undef TRACE_EGL

// These macros generate wrapper functions with the same signatures as the
// OpenGL and EGL API.  These wrappers simply call the tracing version
// of the corresponding function.
#define TRACE_GL_VOID(_api, _args, _argList, ...)                              \
    EGLAPI void _api _args {                                                   \
        /* ALOGD(#_api "() stub..."); */                                       \
        android::gltrace::GLTrace_ ## _api _argList;                           \
    }

#define TRACE_GL(_type, _api, _args, _argList, ...)                            \
    EGLAPI _type _api _args {                                                  \
        /* ALOGD(#_api "() stub..."); */                                       \
        return android::gltrace::GLTrace_ ## _api _argList;                    \
    }

#define TRACE_EGL(_type, _api, _args, _argList, ...)                           \
    EGLAPI _type _api _args {                                                  \
        /* ALOGD(#_api "() stub..."); */                                       \
        return android::gltrace::EGLTrace_wrapper_ ## _api _argList;           \
    }

extern "C" {
#ifdef GLTRACE_GENERATE_API_WRAPPERS
#include "trace.in"
#include "egl_trace.in"
#endif
}

#undef TRACE_GL_VOID
#undef TRACE_GL
#undef TRACE_EGL
