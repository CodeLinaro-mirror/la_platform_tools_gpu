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

#ifndef ANDROID_GLTRACE_HOOKS_H
#define ANDROID_GLTRACE_HOOKS_H

#include <EGL/egl.h>
#include <EGL/eglext.h>
#include <GLES/gl.h>
#include <GLES/glext.h>
#include <GLES2/gl2.h>
#include <GLES2/gl2ext.h>
#include <GLES3/gl3.h>
#include <GLES3/gl31.h>

// maximum number of GL extensions that can be used simultaneously in
// a given process. this limitation exists because we need to have
// a static function for each extension and currently these static functions
// are generated at compile time.
#define MAX_NUMBER_OF_GL_EXTENSIONS 256

namespace android {
namespace gltrace {

#undef EGL_ENTRY
#define EGL_ENTRY(_r, _api, ...) _r (*_api)(__VA_ARGS__);
struct egl_t {
    #include "egl_entries.in"
};
#undef EGL_ENTRY

#undef GL_ENTRY
#define GL_ENTRY(_r, _api, ...) _r (*_api)(__VA_ARGS__);
struct gl_hooks_t {
    struct gl_t {
        #include "entries.in"
    } gl;
    struct gl_ext_t {
        __eglMustCastToProperFunctionPointerType extensions[MAX_NUMBER_OF_GL_EXTENSIONS];
    } ext;
};
#undef GL_ENTRY

} // end of namespace gltrace
} // end of namespace android

#endif /* ANDROID_GLTRACE_HOOKS_H */
