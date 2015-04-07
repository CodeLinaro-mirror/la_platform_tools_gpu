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

#include <gapic/target.h>

#if TARGET_OS == GAPID_OS_WINDOWS

#include <windows.h>
#include <wingdi.h>

namespace android {
namespace caze {

void* GetGfxProcAddress(const char *name) {
    void* p = (void*) wglGetProcAddress(name);

    // Function not found with wglGetProcAddress - try opengl32.dll directly.
    if (p == nullptr ||
        (p == (void*) 0x1) || (p == (void*) 0x2) || (p == (void*) 0x3) || (p == (void*) -1)) {
        static HMODULE module = LoadLibrary(TEXT("opengl32.dll"));
        p = (void*)GetProcAddress(module, name);
    }

    return p;
}

}  // end of namespace caze
}  // end of namespace android

#endif // TARGET_OS == GAPID_OS_WINDOWS
