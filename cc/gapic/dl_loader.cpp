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

#include "dl_loader.h"
#include "log.h"
#include "target.h"

#if TARGET_OS == GAPID_OS_WINDOWS
#   include <windows.h>
#elif TARGET_OS == GAPID_OS_OSX
#   include <dlfcn.h>
#   include <unistd.h>
#else
#   include <dlfcn.h>
#endif

namespace gapic {

#if TARGET_OS == GAPID_OS_WINDOWS

DlLoader::DlLoader(const char* name) {
    mLibrary = reinterpret_cast<void*>(LoadLibraryExA(name, NULL, 0));
    if (mLibrary == nullptr) {
        GAPID_FATAL("Can't load library %s: %d", name, GetLastError());
    }
}

DlLoader::~DlLoader() {
    if (mLibrary != nullptr) {
        FreeLibrary(reinterpret_cast<HMODULE>(mLibrary));
    }
}

void* DlLoader::lookup(const char* name) {
    return reinterpret_cast<void*>(GetProcAddress(reinterpret_cast<HMODULE>(mLibrary), name));
}

#else // if TARGET_OS == GAPID_OS_WINDOWS

DlLoader::DlLoader(const char* name) {
    if (name == nullptr) {
        mLibrary = nullptr;
    } else {
#if TARGET_OS == GAPID_OS_OSX
        // DYLD_FRAMEWORK_PATH takes precedence even with absolute paths.
        // Use a symlink to get to the real library.
        // Credit to apitrace (https://github.com/apitrace) for this nasty, but
        // effective work-around.
        // TODO: not thread-safe.
        char tmp[] = "/tmp/dlopen.XXXXXX";
        if (mktemp(tmp) != nullptr) {
            if (symlink(name, tmp) == 0) {
                mLibrary = dlopen(tmp, RTLD_NOW | RTLD_LOCAL | RTLD_FIRST);
                remove(tmp);
            }
        }
#else // TARGET_OS == GAPID_OS_OSX
        mLibrary = dlopen(name, RTLD_NOW | RTLD_LOCAL);
#endif // TARGET_OS == GAPID_OS_OSX
        if (mLibrary == nullptr) {
            GAPID_FATAL("Can't load library %s: %s", name, dlerror());
        }
    }
}

DlLoader::~DlLoader() {
    if (mLibrary != nullptr) {
        dlclose(mLibrary);
    }
}

void* DlLoader::lookup(const char* name) {
    return dlsym((mLibrary ? mLibrary : RTLD_DEFAULT), name);
}

#endif

}  // namespace gapic
