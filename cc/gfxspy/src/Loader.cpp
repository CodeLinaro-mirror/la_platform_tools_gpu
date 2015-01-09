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

#include "egldefs.h"
#include "glestrace.h"
#include "Loader.h"
#include "log/log.h"

#include <EGL/egl.h>
#include <dlfcn.h>

#if defined(__LP64__)
#define SYSTEM_LIB_PATH "/system/lib64/"
#else
#define SYSTEM_LIB_PATH "/system/lib/"
#endif

#if defined(GLTRACE_DLOPEN_INTERCEPTION)
void* dlopen(const char *filename, int flag) {
    return android::gltrace::Loader::dlopen(filename, flag);
}
#endif // defined(GLTRACE_DLOPEN_INTERCEPTION)

namespace android {
namespace gltrace {

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

static const size_t G_HOOKS_SIZE = 2;
gl_hooks_t gHooks[G_HOOKS_SIZE];

Loader::DlopenFunctionPointerType Loader::sRealDlopenPointer = NULL;
bool Loader::sEnableDlopenRetargeting = false;

void Loader::open(egl_connection_t* cnx) {
    if (cnx->libEgl) return; // Already loaded.

    initDlopen();

    // Don't want retargeting while loading our function tables.
    sEnableDlopenRetargeting = false;

    cnx->libEgl = realDlopen(SYSTEM_LIB_PATH "libEGL.so", RTLD_NOW | RTLD_LOCAL);
    LOG_ALWAYS_FATAL_IF(!cnx->libEgl, "can't load system EGL library: %s", dlerror());

    cnx->libGles1 = realDlopen(SYSTEM_LIB_PATH "libGLESv1_CM.so", RTLD_NOW | RTLD_LOCAL);
    cnx->libGles2 = realDlopen(SYSTEM_LIB_PATH "libGLESv2.so", RTLD_NOW | RTLD_LOCAL);
    LOG_ALWAYS_FATAL_IF(!cnx->libGles2 || !cnx->libGles1,
                        "can't load system GLES library: %s", dlerror());

    GetProcAddressType getProcAddress = (GetProcAddressType)dlsym(cnx->libEgl, "eglGetProcAddress");
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
        initApi(cnx->libGles1, gl_names,
            (__eglMustCastToProperFunctionPointerType*)&gHooks[egl_connection_t::GLESv1_INDEX].gl,
            getProcAddress);
        cnx->hooks[egl_connection_t::GLESv1_INDEX] = &gHooks[egl_connection_t::GLESv1_INDEX];
    }

    if (cnx->libGles2) {
        initApi(cnx->libGles2, gl_names,
            (__eglMustCastToProperFunctionPointerType*)&gHooks[egl_connection_t::GLESv2_INDEX].gl,
            getProcAddress);
        cnx->hooks[egl_connection_t::GLESv2_INDEX] = &gHooks[egl_connection_t::GLESv2_INDEX];
    }
    // Now that we've initialized our function tables, enable the retargeting.
    sEnableDlopenRetargeting = true;
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

void* Loader::dlopen(const char *filename, int flag) {
    if (sRealDlopenPointer == NULL) {
        ALOGW("dlopen() called before initializing sRealDlopenPointer.");
        Loader::initDlopen();
    }
#if defined(GLTRACE_DLOPEN_INTERCEPTION)
    // We won't make any attempt to intercept calls without a file name
    // (e.g. when using RTLD_DEFAULT or RTLD_NEXT).
    if (sEnableDlopenRetargeting && filename != NULL) {
        const char* glesLibnamePrefix = "libGLESv";
        const char* eglLibnamePrefix = "libEGL";
        // See if the .so being opened is the GLES or EGL library.
        const char* matchedString = strstr(filename, glesLibnamePrefix);
        if (matchedString == NULL) {
            matchedString = strstr(filename, eglLibnamePrefix);
        }
        if (matchedString != NULL) {
            // We don't want to intercept calls that could match similarly named libraries
            // (e.g.  "foolibEGL.so"), but we do want to intercept calls that begin
            // with a path (e.g. "/system/lib/libEGL.so").
            if (matchedString == filename || (*(matchedString - 1)) == '/') {
                ALOGI("Changing dlopen(\"%s\", %d) to sRealDlopenPointer(\"libgfxspy.so\", %d).",
                        filename, flag, flag);
                filename = "libgfxspy.so";
            }
        }
    }
#endif // defined(GLTRACE_DLOPEN_INTERCEPTION)
    void *handle = realDlopen(filename, flag);
    if (sEnableDlopenRetargeting && handle == NULL) {
        ALOGD("realDlopen(%s, %d) returned <NULL>.", filename ? filename : "<NULL>", flag);
    }
    return handle;
}

void Loader::initApi(void* dso, char const *const *api,
                      __eglMustCastToProperFunctionPointerType* curr,
                      GetProcAddressType getProcAddress) {
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

void Loader::initDlopen() {
    if (sRealDlopenPointer == NULL) {
#if defined(GLTRACE_DLOPEN_INTERCEPTION)
        ALOGD("Setting up dlopen() interception...");
        sRealDlopenPointer = reinterpret_cast<DlopenFunctionPointerType>(dlsym(
                RTLD_NEXT, "dlopen"));
        if (sRealDlopenPointer == NULL) {
            sRealDlopenPointer = reinterpret_cast<DlopenFunctionPointerType>(dlsym(
                    RTLD_DEFAULT, "dlopen"));
        }
        ALOGD("sRealDlopenPointer = %p.", sRealDlopenPointer);
        LOG_ALWAYS_FATAL_IF(sRealDlopenPointer == NULL,
                "Couldn't find system version of dlopen().");
#else
        ALOGD("dlopen interception is disabled.");
        sRealDlopenPointer = &dlopen;
#endif // defined(GLTRACE_DLOPEN_INTERCEPTION)
    } else {
        ALOGD("Redundant call to initDlopen().");
    }
}

void* Loader::realDlopen(const char *filename, int flag) {
    LOG_ALWAYS_FATAL_IF(sRealDlopenPointer == NULL,
            "dlopen() interception has not been initialized.");
    return (*sRealDlopenPointer)(filename, flag);
}

} // end of namespace gltrace
} // end of namespace android
