/*
 * Copyright (C) 2015 The Android Open Source Project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#ifndef ANDROID_GLESTRACE_H
#define ANDROID_GLESTRACE_H

#include "hooks.h"

namespace android {
namespace gltrace {

/* Hooks to be called by "interesting" EGL functions. */
void GLTrace_eglCreateContext(int version, EGLContext c);
void GLTrace_eglMakeCurrent(EGLContext c);
void GLTrace_eglReleaseThread();
void GLTrace_eglSwapBuffers_internal(void*, void*);

/* Start and stop GL Tracing. */
int GLTrace_start();
void GLTrace_stop();

/* Obtain the gl_hooks structure filled with the trace implementation for all GL functions. */
gl_hooks_t *GLTrace_getGLHooks();

}  // end of namespace gltrace
}  // end of namespace android

#endif // ANDROID_GLESTRACE_H
