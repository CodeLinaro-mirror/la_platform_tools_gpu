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
#include "../target.h" // snprintf

#include <stdio.h>
#include <string>
#include <windows.h>
#include <wingdi.h>

namespace {

typedef void* (__stdcall *PFNWGLGETPROCADDRESS)(const char* name);

std::string opengl32Path() {
    char sysdir[MAX_PATH];
    GetSystemDirectoryA(sysdir, MAX_PATH-1);

    char dllpath[MAX_PATH];
    snprintf(dllpath, MAX_PATH, "%s\\opengl32.dll", sysdir);
    return std::string(dllpath);
}

} // anonymous namespace

namespace gapic {

void* GetGfxProcAddress(const char* name) {
    static DlLoader opengl(opengl32Path().c_str());

    if (void* f = opengl.lookup(name)) {
        return f;
    }

    auto gpa = reinterpret_cast<PFNWGLGETPROCADDRESS>(opengl.lookup("wglGetProcAddress"));
    if (gpa != nullptr) {
        return gpa(name);
    }

    return nullptr;
}

} // namespace gapic
