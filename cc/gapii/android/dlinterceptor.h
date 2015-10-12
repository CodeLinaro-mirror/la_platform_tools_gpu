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

#ifndef GAPII_GLES_INTERCEPTOR_H
#define GAPII_GLES_INTERCEPTOR_H

#include <gapic/dl_loader.h>

#include <set>
#include <vector>

namespace gapii {

class DlInterceptor {
public:
    typedef void (OnLoadTy) (void* handle, const char* name);
    typedef void* (DlopenTy) (const char* path, int flags);
    typedef void* (DlsymTy) (void* handle, const char* name);

    static void init(DlopenTy* dlopen, DlsymTy* dlsym, OnLoadTy* onload);

    // isDriver returns true if the specified .so path refers to a graphics driver.
    static bool isDriver(const char* sopath);

    // dlopen is a replacement for dlopen() that redirects loads of graphics
    // driver libraries to GAPII.
    static void* dlopen(const char* filename, int flag);

    // dlsym is a replacement for dlsym() that redirects function lookups to the
    // real driver libraries when they're not found in GAPIS.
    static void* dlsym(void* handle, const char* name);

    // load can be used to load a library with the specified path.
    // loader matches the gapic::DlLoader::Loader signature and can be passed to
    // gapic::DlLoader::setCustomLoader().
    static void* load(const char* name);

    // resolve can be used to find a function in a library.
    // resolver matches the gapic::DlLoader::Resolver signature and can be passed to
    // gapic::DlLoader::setCustomResolver().
    static void* resolve(void* handle, const char* name);

private:
    static void* getLibGAPII();

    // sImports contains the opaque handles to the real graphics libraries
    // that GAPII will intercept when calling dlopen().
    static std::vector<void*> sImports;

    // sExports contains the opaque handles to the GAPII library returned by
    // dlopen().
    static std::set<void*> sExports;

    // sDlopen holds the address of the regular dlopen() function.
    static DlopenTy* sDlopen;

    // sDlsym holds the address of the regular dlsym() function.
    static DlsymTy* sDlsym;

    // sOnLoad is the callback function that will be invoked whenever
    // an intercepted call to dlopen() is made.
    static OnLoadTy* sOnLoad;
};

}  // namespace gapii

#endif  // GAPII_GLES_INTERCEPTOR_H
