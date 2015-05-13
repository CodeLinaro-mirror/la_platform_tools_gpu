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

#ifndef ANDROID_GLTRACE_EGLDEFS_H
#define ANDROID_GLTRACE_EGLDEFS_H

#include "hooks.h"

namespace android {
namespace gltrace {

struct egl_connection_t {
    enum {
        GLESv1_INDEX = 0,
        GLESv2_INDEX = 1
    };

    inline egl_connection_t() : dso(0) { }
    void *              dso;
    gl_hooks_t *        hooks[2];
    EGLint              major;
    EGLint              minor;
    egl_t               egl;

    void*               libEgl;
    void*               libGles1;
    void*               libGles2;
};

extern egl_connection_t gEGLImpl;

} // end of namespace gltrace
} // end of namespace android

#endif /* ANDROID_GLTRACE_EGLDEFS_H */
