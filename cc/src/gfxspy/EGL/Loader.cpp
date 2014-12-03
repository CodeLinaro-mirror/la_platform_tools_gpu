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

#include <ctype.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <errno.h>
#include <dlfcn.h>
#include <limits.h>
#include <dirent.h>

#include <cutils/log.h>

#include <EGL/egl.h>

#include "../glestrace.h"

#include "EGL/egldefs.h"
#include "EGL/Loader.h"

// ----------------------------------------------------------------------------
namespace android {
// ----------------------------------------------------------------------------

static pthread_mutex_t sLogPrintMutex = PTHREAD_MUTEX_INITIALIZER;
static nsecs_t sLogPrintTime = 0;
#define NSECS_DURATION 1000000000

void gl_unimplemented() {
    bool printLog = false;
    nsecs_t now = systemTime();
    pthread_mutex_lock(&sLogPrintMutex);
    if ((now - sLogPrintTime) > NSECS_DURATION) {
        sLogPrintTime = now;
        printLog = true;
    }
    pthread_mutex_unlock(&sLogPrintMutex);
    if (printLog) {
        ALOGE("called unimplemented OpenGL ES API");
    }
}

void gl_noop() {
}


// ----------------------------------------------------------------------------
// GL / EGL hooks
// ----------------------------------------------------------------------------

#undef GL_ENTRY
#undef EGL_ENTRY
#define GL_ENTRY(_r, _api, ...) #_api,
#define EGL_ENTRY(_r, _api, ...) #_api,

    char const * const gl_names[] = {
#include "../entries.in"
        NULL
    };

    char const * const egl_names[] = {
#include "egl_entries.in"
        NULL
    };

#undef GL_ENTRY
#undef EGL_ENTRY


/*
 * EGL userspace drivers must be provided either:
 * - as a single library:
 *      /vendor/lib/egl/libGLES.so
 *
 * - as separate libraries:
 *      /vendor/lib/egl/libEGL.so
 *      /vendor/lib/egl/libGLESv1_CM.so
 *      /vendor/lib/egl/libGLESv2.so
 *
 * The software renderer for the emulator must be provided as a single
 * library at:
 *
 *      /system/lib/egl/libGLES_android.so
 *
 *
 * For backward compatibility and to facilitate the transition to
 * this new naming scheme, the loader will additionally look for:
 *
 *      /{vendor|system}/lib/egl/lib{GLES | [EGL|GLESv1_CM|GLESv2]}_*.so
 *
 */

ANDROID_SINGLETON_STATIC_INSTANCE( Loader )

// ----------------------------------------------------------------------------

Loader::driver_t::driver_t(void* gles)
{
    dso[0] = gles;
    for (size_t i=1 ; i<NELEM(dso) ; i++)
        dso[i] = 0;
}

Loader::driver_t::~driver_t()
{
    for (size_t i=0 ; i<NELEM(dso) ; i++) {
        if (dso[i]) {
            dlclose(dso[i]);
            dso[i] = 0;
        }
    }
}

status_t Loader::driver_t::set(void* hnd, int32_t api)
{
    switch (api) {
        case EGL:
            dso[0] = hnd;
            break;
        case GLESv1_CM:
            dso[1] = hnd;
            break;
        case GLESv2:
            dso[2] = hnd;
            break;
        default:
            return BAD_INDEX;
    }
    return NO_ERROR;
}

// ----------------------------------------------------------------------------

Loader::Loader()
    : getProcAddress(NULL) {
}

Loader::~Loader() {
    gltrace::GLTrace_stop();
}

static void* load_wrapper(const char* path) {
    void* so = dlopen(path, RTLD_NOW | RTLD_LOCAL);
    ALOGE_IF(!so, "dlopen(\"%s\") failed: %s", path, dlerror());
    return so;
}

void* Loader::open(egl_connection_t* cnx)
{
    void* dso;
    driver_t* hnd = 0;

#if defined(__LP64__)
    cnx->libEgl   = load_wrapper("/system/lib64/libEGL.so");
    cnx->libGles2 = load_wrapper("/system/lib64/libGLESv2.so");
    cnx->libGles1 = load_wrapper("/system/lib64/libGLESv1_CM.so");
#else
    cnx->libEgl   = load_wrapper("/system/lib/libEGL.so");
    cnx->libGles2 = load_wrapper("/system/lib/libGLESv2.so");
    cnx->libGles1 = load_wrapper("/system/lib/libGLESv1_CM.so");
#endif
    LOG_ALWAYS_FATAL_IF(!cnx->libEgl,
            "couldn't load system EGL wrapper libraries");

    LOG_ALWAYS_FATAL_IF(!cnx->libGles2 || !cnx->libGles1,
            "couldn't load system OpenGL ES wrapper libraries");

    dso = load_driver("GLES", cnx, EGL | GLESv1_CM | GLESv2);
    if (dso) {
        hnd = new driver_t(dso);
    } else {
        // Always load EGL first
        dso = load_driver("EGL", cnx, EGL);
        if (dso) {
            hnd = new driver_t(dso);
            hnd->set( load_driver("GLESv1_CM", cnx, GLESv1_CM), GLESv1_CM );
            hnd->set( load_driver("GLESv2",    cnx, GLESv2),    GLESv2 );
        }
    }

    LOG_ALWAYS_FATAL_IF(!hnd, "couldn't find an OpenGL ES implementation");

    return (void*)hnd;
}

status_t Loader::close(void* driver)
{
    driver_t* hnd = (driver_t*)driver;
    delete hnd;
    return NO_ERROR;
}

void Loader::init_api(void* dso,
        char const * const * api,
        __eglMustCastToProperFunctionPointerType* curr,
        getProcAddressType getProcAddress)
{
    const ssize_t SIZE = 256;
    char scrap[SIZE];
    while (*api) {
        char const * name = *api;
        __eglMustCastToProperFunctionPointerType f =
            (__eglMustCastToProperFunctionPointerType)dlsym(dso, name);
        if (f == NULL) {
            // couldn't find the entry-point, use eglGetProcAddress()
            f = getProcAddress(name);
        }
        if (f == NULL) {
            // Try without the OES postfix
            ssize_t index = ssize_t(strlen(name)) - 3;
            if ((index>0 && (index<SIZE-1)) && (!strcmp(name+index, "OES"))) {
                strncpy(scrap, name, index);
                scrap[index] = 0;
                f = (__eglMustCastToProperFunctionPointerType)dlsym(dso, scrap);
                //ALOGD_IF(f, "found <%s> instead", scrap);
            }
        }
        if (f == NULL) {
            // Try with the OES postfix
            ssize_t index = ssize_t(strlen(name)) - 3;
            if (index>0 && strcmp(name+index, "OES")) {
                snprintf(scrap, SIZE, "%sOES", name);
                f = (__eglMustCastToProperFunctionPointerType)dlsym(dso, scrap);
                //ALOGD_IF(f, "found <%s> instead", scrap);
            }
        }
        if (f == NULL) {
            //ALOGD("%s", name);
            f = (__eglMustCastToProperFunctionPointerType)gl_unimplemented;

            /*
             * GL_EXT_debug_label is special, we always report it as
             * supported, it's handled by GLES_trace. If GLES_trace is not
             * enabled, then these are no-ops.
             */
            if (!strcmp(name, "glInsertEventMarkerEXT")) {
                f = (__eglMustCastToProperFunctionPointerType)gl_noop;
            } else if (!strcmp(name, "glPushGroupMarkerEXT")) {
                f = (__eglMustCastToProperFunctionPointerType)gl_noop;
            } else if (!strcmp(name, "glPopGroupMarkerEXT")) {
                f = (__eglMustCastToProperFunctionPointerType)gl_noop;
            }
        }
        *curr++ = f;
        api++;
    }
}

void *Loader::load_driver(const char* kind,
        egl_connection_t* cnx, uint32_t mask)
{
    class MatchFile {
    public:
        static String8 find(const char* kind) {
            String8 result;
            String8 pattern;
            pattern.appendFormat("lib%s", kind);
            const char* const searchPaths[] = {
#if defined(__LP64__)
                    "/vendor/lib64/egl",
                    "/system/lib64/egl"
#else
                    "/vendor/lib/egl",
                    "/system/lib/egl"
#endif
            };

            // first, we search for the exact name of the GLES userspace
            // driver in both locations.
            // i.e.:
            //      libGLES.so, or:
            //      libEGL.so, libGLESv1_CM.so, libGLESv2.so

            for (size_t i=0 ; i<NELEM(searchPaths) ; i++) {
                if (find(result, pattern, searchPaths[i], true)) {
                    return result;
                }
            }

            // for compatibility with the old "egl.cfg" naming convention
            // we look for files that match:
            //      libGLES_*.so, or:
            //      libEGL_*.so, libGLESv1_CM_*.so, libGLESv2_*.so

            pattern.append("_");
            for (size_t i=0 ; i<NELEM(searchPaths) ; i++) {
                if (find(result, pattern, searchPaths[i], false)) {
                    return result;
                }
            }

            // we didn't find the driver. gah.
            result.clear();
            return result;
        }

    private:
        static bool find(String8& result,
                const String8& pattern, const char* const search, bool exact) {

            if (exact) {
                String8 absolutePath;
                absolutePath.appendFormat("%s/%s.so", search, pattern.string());
                if (!access(absolutePath.string(), R_OK)) {
                    result = absolutePath;
                    return true;
                }
                return false;
            }

            DIR* d = opendir(search);
            if (d != NULL) {
                struct dirent cur;
                struct dirent* e;
                while (readdir_r(d, &cur, &e) == 0 && e) {
                    if (e->d_type == DT_DIR) {
                        continue;
                    }
                    if (!strcmp(e->d_name, "libGLES_android.so")) {
                        // always skip the software renderer
                        continue;
                    }
                    if (strstr(e->d_name, pattern.string()) == e->d_name) {
                        if (!strcmp(e->d_name + strlen(e->d_name) - 3, ".so")) {
                            result.clear();
                            result.appendFormat("%s/%s", search, e->d_name);
                            closedir(d);
                            return true;
                        }
                    }
                }
                closedir(d);
            }
            return false;
        }
    };


    String8 absolutePath = MatchFile::find(kind);
    if (absolutePath.isEmpty()) {
        // this happens often, we don't want to log an error
        return 0;
    }
    const char* const driver_absolute_path = absolutePath.string();

    // This loads the vender version of the GLES driver.
    // TODO: Remove all the code associated with loading the vendor
    // driver.  This version of gltrace (gfxspy) never uses the vendor
    // version of the driver.  Deleting the vendor driver code
    // requires some modifications to code that calls this function
    // (and probably a few other places).
    void* dso = dlopen(driver_absolute_path, RTLD_NOW | RTLD_LOCAL);
    if (dso == 0) {
        const char* err = dlerror();
        ALOGE("load_driver(%s): %s", driver_absolute_path, err?err:"unknown");
        return 0;
    }

    ALOGD("loaded %s", driver_absolute_path);

    if (mask & EGL) {
        // Use the system version of EGL, not the vendor version.
        void* dso = cnx->libEgl;
        getProcAddress = (getProcAddressType)dlsym(dso, "eglGetProcAddress");

        ALOGE_IF(!getProcAddress,
                "can't find eglGetProcAddress() in %s", driver_absolute_path);

        egl_t* egl = &cnx->egl;
        __eglMustCastToProperFunctionPointerType* curr =
            (__eglMustCastToProperFunctionPointerType*)egl;
        char const * const * api = egl_names;
        while (*api) {
            char const * name = *api;
            __eglMustCastToProperFunctionPointerType f =
                (__eglMustCastToProperFunctionPointerType)dlsym(dso, name);
            if (f == NULL) {
                // couldn't find the entry-point, use eglGetProcAddress()
                f = getProcAddress(name);
                if (f == NULL) {
                    f = (__eglMustCastToProperFunctionPointerType)0;
                }
            }
            *curr++ = f;
            api++;
        }
    }

    if (mask & GLESv1_CM) {
        // Use the system version of the .so, and not the vendor version.
        void* dso = cnx->libGles1;
        init_api(dso, gl_names,
            (__eglMustCastToProperFunctionPointerType*)
                &cnx->hooks[egl_connection_t::GLESv1_INDEX]->gl,
            getProcAddress);
    }

    if (mask & GLESv2) {
        // Use the system version of the .so, and not the vendor version.
        void* dso = cnx->libGles2;
        init_api(dso, gl_names,
            (__eglMustCastToProperFunctionPointerType*)
                &cnx->hooks[egl_connection_t::GLESv2_INDEX]->gl,
            getProcAddress);
    }

    return dso;
}

// ----------------------------------------------------------------------------
} // namespace android
// ----------------------------------------------------------------------------
