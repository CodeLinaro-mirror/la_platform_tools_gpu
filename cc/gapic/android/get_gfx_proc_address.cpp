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

#include "../dl_loader.h"
#include "../log.h"

#include <string>
#include <unordered_map>

#if defined(__LP64__)
#define SYSTEM_LIB_PATH "/system/lib64/"
#else
#define SYSTEM_LIB_PATH "/system/lib/"
#endif

namespace gapic {

void* GetGfxProcAddress(const char* name) {
    GAPID_INFO("GetGfxProcAddress(%s)", name);

    static std::unordered_map<std::string, void*> cache;
    auto it = cache.find(name);
    if (it != cache.end()) {
        GAPID_INFO("GetGfxProcAddress(%s) -> 0x%x (from cache)", name, it->second);
        return it->second;
    }

    static DlLoader libEGL(SYSTEM_LIB_PATH "libEGL.so");
    if (void* proc = libEGL.lookup(name)) {
        GAPID_INFO("GetGfxProcAddress(%s) -> 0x%x (from libEGL dlsym)", name, proc);
        cache[name] = proc;
        return proc;
    }

    static DlLoader libGLESv2(SYSTEM_LIB_PATH "libGLESv2.so");
    if (void* proc = libGLESv2.lookup(name)) {
        GAPID_INFO("GetGfxProcAddress(%s) -> 0x%x (from libGLESv2 dlsym)", name, proc);
        cache[name] = proc;
        return proc;
    }

    static DlLoader libGLESv3(SYSTEM_LIB_PATH "libGLESv3.so");
    if (void* proc = libGLESv3.lookup(name)) {
        GAPID_INFO("GetGfxProcAddress(%s) -> 0x%x (from libGLESv3 dlsym)", name, proc);
        cache[name] = proc;
        return proc;
    }

    typedef void*(*getProcAddressType)(const char*);
    static getProcAddressType eglGetProcAddress =
            reinterpret_cast<getProcAddressType>(libEGL.lookup("eglGetProcAddress"));
    if (eglGetProcAddress == nullptr) {
        return nullptr;
    }

    if (void* proc = eglGetProcAddress(name)) {
        GAPID_INFO("GetGfxProcAddress(%s) -> 0x%x (via eglGetProcAddress)", name, proc);
        cache[name] = proc;
        return proc;
    }

    GAPID_INFO("GetGfxProcAddress(%s) -> not found", name);
    cache[name] = nullptr;
    return nullptr;
}

}  // namespace gapic

