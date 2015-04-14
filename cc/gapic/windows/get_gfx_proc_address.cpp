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

#include <stdio.h>
#include <windows.h>
#include <wingdi.h>

#include <gapic/target.h> // snprintf

namespace {

typedef void* (__stdcall *PFNWGLGETPROCADDRESS)(const char* name);

HMODULE loadOpengl32() {
    char sysdir[MAX_PATH];
    GetSystemDirectoryA(sysdir, MAX_PATH-1);

    char dllpath[MAX_PATH];
    snprintf(dllpath, MAX_PATH, "%s\\opengl32.dll", sysdir);
    return LoadLibraryExA(dllpath, NULL, 0);
}

} // anonymous namespace

namespace gapic {

void* GetGfxProcAddress(const char* name) {
    static HMODULE module = loadOpengl32();

    if (void* f = reinterpret_cast<void*>(GetProcAddress(module, name))) {
        return f;
    }

    static PFNWGLGETPROCADDRESS gpa = reinterpret_cast<PFNWGLGETPROCADDRESS>(
            GetProcAddress(module, "wglGetProcAddress"));
    if (gpa != nullptr) {
        return gpa(name);
    }

    return nullptr;
}

} // namespace gapic
