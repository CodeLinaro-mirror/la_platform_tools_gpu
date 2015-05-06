/*
 * Copyright 2015, The Android Open Source Project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
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
#include <windows.h>
#else
#include <dlfcn.h>
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
        mLibrary = dlopen(name, RTLD_NOW | RTLD_LOCAL);
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

