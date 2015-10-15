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

#include "dlinfo.h"
#include "dlinterceptor.h"

#include <gapic/log.h>

#include <cstring>
#include <string>

#include <dlfcn.h>

namespace {

// intercept is a null-terminated list of library names that GAPII will
// intercept when loaded with dlopen().
static const char* intercept[] = {
  "libGLES.so",
  "libEGL.so",
  "libGLESv1_CM.so",
  "libGLESv2.so",
  "libGLESv3.so",
  nullptr,
};

// drivers is a null-terminated list of library names that GAPII considers a
// driver library. dlopen() and dlsym() must not be intercepted for these
// libraries.
static const char* drivers[] = {
  "gralloc.msm8960.so",
  "libadreno_utils",
  "libadreno_utils.so",
  "libEGL.so",
  "libEGL_adreno.so",
  "libEGL_adreno200.so",
  "libGLES_mali.so",
  "libGLESv1_CM.so",
  "libGLESv1_CM_adreno.so",
  "libGLESv2.so",
  "libGLESv2_adreno.so",
  "libGLESv2_adreno200.so",
  "libGLESv3.so",
  "libgsl.so",
  "libq3dtools_adreno200.so",
  "libsc-a3xx.so",
  nullptr,
};

// isIn returns true if the specified path is found in the null-terminated array.
bool isIn(const char* path, const char* array[]) {
    for (int i = 0; array[i] != nullptr; i++) {
        if (const char* match = strstr(path, array[i])) {
            // We don't want to intercept calls that could match similarly named libraries
            // (e.g. "foolibEGL.so"), but we do want to intercept calls that begin
            // with a path (e.g. "/system/lib/libEGL.so").
            if (match == path || (*(match - 1)) == '/') {
                return true;
            }
        }
    }
    return false;
}

}  // anonymous namespace

namespace gapii {

std::vector<void*>        DlInterceptor::sImports;
std::set<void*>           DlInterceptor::sExports;
DlInterceptor::DlopenTy*  DlInterceptor::sDlopen = nullptr;
DlInterceptor::DlsymTy*   DlInterceptor::sDlsym = nullptr;
DlInterceptor::OnLoadTy*  DlInterceptor::sOnLoad = nullptr;

void DlInterceptor::init(DlopenTy* dlopen, DlsymTy* dlsym, OnLoadTy onload) {
    sDlopen = dlopen;
    sDlsym = dlsym;
    sOnLoad = onload;
}

bool DlInterceptor::isDriver(const char* sopath) {
    return isIn(sopath, drivers);
}

void* DlInterceptor::dlopen(const char* filename, int flag) {
    GAPID_INFO("dlopen(\"%s\")", filename);
    if (filename == nullptr) {
        // nullptr filename means application handle.
        return sDlopen(filename, flag);
    }

    if (isIn(filename, intercept)) {
      GAPID_INFO("dlopen() interception for \"%s\"", filename);
      return getLibGAPII();
    }

    // Not something we want to intercept.
    // Pass through to regular dlopen().
    if (void* lib = sDlopen(filename, flag)) {
        sOnLoad(lib, filename);
        return lib;
    }

    return nullptr; // Not found.
}

void* DlInterceptor::dlsym(void* handle, const char* name) {
    GAPID_INFO("dlsym(\"%s\")", name);

    if (sExports.find(handle) == sExports.end()) {
        // Handle doesn't belong to GAPII. Pass through to the real dlsym().
        return sDlsym(handle, name);
    }

    // Handle belongs to GAPII. Search GAPII for symbol.
    static void* gapii = getLibGAPII();
    if (void* res = sDlsym(gapii, name)) {
        GAPID_INFO("dlsym() interception for \"%s\"", name);
        return res;
    }

    // GAPII did not contain the symbol. Fall-back to the intercepted libraries.
    static bool loadedImports = false;
    if (!loadedImports) {
        GAPID_INFO("Loading intercepted graphics libraries...");
        // Load all the libraries GAPII will intercept.
        for (int i = 0; intercept[i] != nullptr; i++) {
            if (void* lib = load(intercept[i])) {
                GAPID_INFO("  Loaded library '%s'.", intercept[i]);
                sOnLoad(lib, intercept[i]);
                sImports.push_back(lib);
            } else {
                GAPID_INFO("  Library '%s' was not found.", intercept[i]);
            }
        }
        GAPID_INFO("Done.");
        loadedImports = true;
    }

    for (void* lib : sImports) {
        if (void* res = sDlsym(lib, name)) {
            GAPID_WARNING("GAPII did not expose '%s', but was exposed by intercepted library.", name);
            return res;
        }
    }

    return nullptr; // Not found.
}

void* DlInterceptor::load(const char* name) {
    GAPID_INFO("load(\"%s\")", name);
    return sDlopen(name, RTLD_NOW | RTLD_LOCAL);
}

void* DlInterceptor::resolve(void* handle, const char* name) {
    GAPID_INFO("resolve(%p, \"%s\")", handle, name);
    return sDlsym(handle, name);
}

void* DlInterceptor::getLibGAPII() {
    static std::string sGAPIIPath;
    if (sGAPIIPath.size() == 0) {
        GAPID_INFO("Searching for GAPII library...");
        DlInfo dlinfo;
        if (const char* error = DlInfo::self(dlinfo)) {
            GAPID_FATAL("GAPII library path could not be found: %s", error);
        }
        sGAPIIPath = dlinfo.mPath;
        GAPID_INFO("GAPII library found at: '%s'", sGAPIIPath.c_str());
    }

    if (void* gapii = load(sGAPIIPath.c_str())) {
        sExports.insert(gapii); // TODO: ref-count and remove unreferenced entries?
        return gapii;
    }

    GAPID_FATAL("GAPII library could not be dlopen()'d from '%s'", sGAPIIPath.c_str());
    return nullptr; // Unreachable
}

}  // namespace gapii
