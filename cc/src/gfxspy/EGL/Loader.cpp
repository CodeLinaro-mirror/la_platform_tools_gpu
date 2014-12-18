/*
 ** Copyright 2007, The Android Open Source Project
 **
 ** Licensed under the Apache License, Version 2.0 (the "License");
 ** you may not use this file except in compliance with the License.
 ** You may obtain a copy of the License at
 **
 **     http://www.apache.org/licenses/LICENSE-2.0
 **
 ** Unless required by applicable law or agreed to in writing, software
 ** distributed under the License is distributed on an "AS IS" BASIS,
 ** WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 ** See the License for the specific language governing permissions and
 ** limitations under the License.
 */

#include "EGL/Loader.h"
#include "glestrace.h"

#include <EGL/egl.h>
#include <EGL/egldefs.h>
#include <cutils/log.h>
#include <dlfcn.h>

#if defined(__LP64__)
#define SYSTEM_LIB_PATH "/system/lib64/"
#else
#define SYSTEM_LIB_PATH "/system/lib/"
#endif

namespace android {

void gl_unimplemented() {
    ALOGE("called unimplemented OpenGL ES API");
}

void gl_noop() {
}

// OpenGL ES hooks

#undef GL_ENTRY
#define GL_ENTRY(_r, _api, ...) #_api,
char const * const gl_names[] = {
#include "entries.in"
    NULL
};
#undef GL_ENTRY

// EGL hooks

#undef EGL_ENTRY
#define EGL_ENTRY(_r, _api, ...) #_api,
char const * const egl_names[] = {
#include "egl_entries.in"
    NULL
};
#undef EGL_ENTRY


/*
 * Note: the loader will look for the following system drivers:
 *   /system/lib{,64}/egl/lib{EGL|GLESv1_CM|GLESv2}.so
 *
 * It does not support the deprecated libGLES.so all-in-one driver model.
 *
 */

ANDROID_SINGLETON_STATIC_INSTANCE(Loader);

static const size_t G_HOOKS_SIZE = 2;
::android::gl_hooks_t gHooks[G_HOOKS_SIZE];  // Declared in EGL/egldefs.h

Loader::Loader() {
}

Loader::~Loader() {
    gltrace::GLTrace_stop();
}

void Loader::open(egl_connection_t* cnx) {
    if (cnx->libEgl) return; // Already loaded.

    cnx->libEgl = dlopen(SYSTEM_LIB_PATH "libEGL.so", RTLD_NOW | RTLD_LOCAL);
    LOG_ALWAYS_FATAL_IF(!cnx->libEgl, "can't load system EGL library: %s", dlerror());

    cnx->libGles1 = dlopen(SYSTEM_LIB_PATH "libGLESv1_CM.so", RTLD_NOW | RTLD_LOCAL);
    cnx->libGles2 = dlopen(SYSTEM_LIB_PATH "libGLESv2.so", RTLD_NOW | RTLD_LOCAL);
    LOG_ALWAYS_FATAL_IF(!cnx->libGles2 || !cnx->libGles1,
                        "can't load system GLES library: %s", dlerror());

    getProcAddressType getProcAddress = (getProcAddressType)dlsym(cnx->libEgl, "eglGetProcAddress");
    LOG_ALWAYS_FATAL_IF(!getProcAddress, "can't find eglGetProcAddress(): %s", dlerror());

    egl_t *egl = &cnx->egl;
    __eglMustCastToProperFunctionPointerType *curr = (__eglMustCastToProperFunctionPointerType*)egl;

    char const *const *api = egl_names;
    while (*api) {
        char const *name = *api;
        __eglMustCastToProperFunctionPointerType fptr =
            (__eglMustCastToProperFunctionPointerType)dlsym(cnx->libEgl, name);
        if (fptr == NULL) fptr = getProcAddress(name);
        *curr++ = reinterpret_cast<__eglMustCastToProperFunctionPointerType>(fptr);
        api++;
    }

    memset(gHooks, sizeof(gHooks), 0);

    if (cnx->libGles1) {
        init_api(cnx->libGles1, gl_names,
            (__eglMustCastToProperFunctionPointerType*)&gHooks[egl_connection_t::GLESv1_INDEX].gl,
            getProcAddress);
        cnx->hooks[egl_connection_t::GLESv1_INDEX] = &gHooks[egl_connection_t::GLESv1_INDEX];
    }

    if (cnx->libGles2) {
        init_api(cnx->libGles2, gl_names,
            (__eglMustCastToProperFunctionPointerType*)&gHooks[egl_connection_t::GLESv2_INDEX].gl,
            getProcAddress);
        cnx->hooks[egl_connection_t::GLESv2_INDEX] = &gHooks[egl_connection_t::GLESv2_INDEX];
    }
}

void Loader::close(egl_connection_t* cnx) {
    if (cnx->libEgl) {
        dlclose(cnx->libEgl);
        cnx->libEgl = NULL;
    }

    if (cnx->libGles1) {
        dlclose(cnx->libGles1);
        cnx->libGles1 = NULL;
    }

    if (cnx->libGles2) {
        dlclose(cnx->libGles2);
        cnx->libGles2 = NULL;
    }
}

void Loader::init_api(void* dso, char const *const *api,
                      __eglMustCastToProperFunctionPointerType* curr,
                      getProcAddressType getProcAddress) {
    while (*api) {
        const char *name = *api;
        __eglMustCastToProperFunctionPointerType fptr =
            (__eglMustCastToProperFunctionPointerType)dlsym(dso, name);

        if (fptr == NULL) {
            fptr = getProcAddress(name);
        }
        if (fptr == NULL) {
            /*
            * GL_EXT_debug_label functions are special, we always report the extension as supported,
            * which gets handled by GLES_trace. If GLES_trace is not enabled, then these are no-ops.
            */
            if (!strcmp(name, "glInsertEventMarkerEXT") || !strcmp(name, "glPushGroupMarkerEXT") ||
                !strcmp(name, "glPopGroupMarkerEXT")) {
                fptr = gl_noop;
            } else {
                fptr = gl_unimplemented;
            }
        }
        *curr++ = fptr;
        api++;
    }
}

} // end of namespace android

#if defined(GLTRACE_DLOPEN_INTERCEPTION)

static void *(*_dlopen)(const char *filename, int flag) = NULL;
typedef void *(*dlopen_function_pointer)(const char *filename, int flag);

__attribute__((constructor))
static void init_dlopen() {
    ALOGD("<gfxspy> setting up dlopen interception...");
    if (_dlopen == NULL) _dlopen = (dlopen_function_pointer)dlsym(RTLD_NEXT, "dlopen");
    if (_dlopen == NULL) _dlopen = (dlopen_function_pointer)dlsym(RTLD_DEFAULT, "dlopen");
    ALOGD("<gfxspy> _dlopen = %p", _dlopen);
}

void *dlopen(const char *filename, int flag) {
    if (_dlopen == NULL) init_dlopen();
    ALOGD("<gfxspy> intercepting dlopen(\"%s\", %x)", filename, flag);
    void *handle = (*_dlopen)(filename, flag);
    if (handle == NULL) {
        ALOGD("<gfxspy> got a NULL handle from dlopen!\n");
    }
    return handle;
}

#endif // defined(GLTRACE_DLOPEN_INTERCEPTION)
