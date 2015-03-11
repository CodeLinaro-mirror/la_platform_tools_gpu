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

#include "spy.h"
#include "file_writer.h"

void __stdcall init(int32_t, int32_t, gfxspy::RenderbufferFormat, gfxspy::RenderbufferFormat, gfxspy::RenderbufferFormat){}
void __stdcall glClearColor(float, float, float, float){}
void __stdcall glClear(gfxspy::ClearMask){}

int main(int argc, char** argv) {
    using namespace gfxspy;

    FileWriter file("atoms");
    Encoder encoder(&file);

    Spy spy(&encoder);

    encoder.U32(9);
    spy.init(init, 256, 256, RenderbufferFormat::GL_RGB565, RenderbufferFormat::GL_DEPTH_COMPONENT16, RenderbufferFormat::GL_STENCIL_INDEX8);
    spy.glClearColor(glClearColor, 1, 0, 0, 1);
    spy.glClear(glClear, ClearMask::GL_COLOR_BUFFER_BIT);
    spy.glClearColor(glClearColor, 0, 1, 0, 1);
    spy.glClear(glClear, ClearMask::GL_COLOR_BUFFER_BIT);
    spy.glClearColor(glClearColor, 0, 0, 1, 1);
    spy.glClear(glClear, ClearMask::GL_COLOR_BUFFER_BIT);
    spy.glClearColor(glClearColor, 0, 0, 0, 1);
    spy.glClear(glClear, ClearMask::GL_COLOR_BUFFER_BIT);
}