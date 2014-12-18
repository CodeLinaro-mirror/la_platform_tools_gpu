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

#ifndef ANDROID_GLTRACE_EGL_H
#define ANDROID_GLTRACE_EGL_H

namespace android {
namespace gltrace {

void GLTrace_eglCreateContext(int version, int contextId);
void GLTrace_eglMakeCurrent(int contextId);
void GLTrace_eglSwapBuffers(void *dpy, void *draw);

// Forward declare the EGL tracing functions.
#undef TRACE_EGL
#define TRACE_EGL(_type, _api, _args, _argList, ...) EGLAPI _type EGLTrace_wrapper_ ## _api _args;
#include "egl_trace.in"
#undef TRACE_EGL

} // end of namespace gltrace
} // end of namespace android

#endif // ANDROID_GLTRACE_EGL_H
