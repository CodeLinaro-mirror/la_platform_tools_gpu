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

#include "../dl_loader.h"
#include "../log.h"

namespace gapic {

#define FRAMEWORK_ROOT "/System/Library/Frameworks/OpenGL.framework/"

void* GetGfxProcAddress(const char *name, bool bypassLocal) {
   if (bypassLocal) {
        static DlLoader opengl(FRAMEWORK_ROOT "OpenGL");
        if (void* proc = opengl.lookup(name)) {
            GAPID_INFO("GetGfxProcAddress(%s, %d) -> 0x%x (from OpenGL dlsym)", name, bypassLocal, proc);
            return proc;
        }

        static DlLoader libgl(FRAMEWORK_ROOT "Libraries/libGL.dylib");
        if (void* proc = libgl.lookup(name)) {
            GAPID_INFO("GetGfxProcAddress(%s, %d) -> 0x%x (from libGL dlsym)", name, bypassLocal, proc);
            return proc;
        }

        static DlLoader libglu(FRAMEWORK_ROOT "Libraries/libGLU.dylib");
        if (void* proc = libglu.lookup(name)) {
            GAPID_INFO("GetGfxProcAddress(%s, %d) -> 0x%x (from libGLU dlsym)", name, bypassLocal, proc);
            return proc;
        }
    } else {
        static DlLoader local(nullptr);
        return local.lookup(name);
    }

    GAPID_INFO("GetGfxProcAddress(%s, %d) -> not found", name, bypassLocal);
    return nullptr;
}

} // namespace gapic
