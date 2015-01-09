/*
 ** Copyright 2009, The Android Open Source Project
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

#ifndef ANDROID_GLTRACE_LOADER_H
#define ANDROID_GLTRACE_LOADER_H

#include <EGL/egl.h>

namespace android {
namespace gltrace {

struct egl_connection_t;

class Loader {
public:
    static void open(egl_connection_t* cnx);
    static void close(egl_connection_t* cnx);

    static void* dlopen(const char *filename, int flag);

private:
    typedef void *(*DlopenFunctionPointerType)(const char *filename, int flag);
    typedef __eglMustCastToProperFunctionPointerType(*GetProcAddressType)(const char*);

    // Always points to the system version of dlopen (to bypass dlopen interception).
    static DlopenFunctionPointerType sRealDlopenPointer;

    // When true, dlopen calls for GLES and EGL will be redirected to this library.
    // Must be false during bootstrapping to avoid redirecting dlopen calls from
    // the system and vendor drivers.
    static bool sEnableDlopenRetargeting;

    static void initApi(void *dso, char const *const *api,
            __eglMustCastToProperFunctionPointerType* curr,
            GetProcAddressType getProcAddress);

    static void initDlopen();
    static void* realDlopen(const char *filename, int flag);

    Loader();
    ~Loader();
};

} // end of namespace gltrace
} // end of namespace android

#endif // ANDROID_GLTRACE_LOADER_H
