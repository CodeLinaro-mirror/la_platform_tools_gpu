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

#include <gapic/target.h>

#if TARGET_OS == TARGET_OS_WINDOWS
#include "windows/wgl.h"
#endif // TARGET_OS

namespace gapii {

Spy::Spy(std::shared_ptr<gapic::Encoder> encoder)
    : mEncoder(encoder)
    , GlesSpy(encoder) {
}

Spy::~Spy() {
    mEncoder->U16(0xffff); // Type ID -- TODO: mEncoder->Id(EOS_ID);
}

void Spy::init(int32_t width, int32_t height,
        uint32_t colorFormat, uint32_t depthFormat, uint32_t stencilFormat) {
    mState.init(width, height, colorFormat, depthFormat, stencilFormat);
    mEncoder->U16(0); // INIT_ID
    mEncoder->U32(0); // ContextID
    mEncoder->S32(width);
    mEncoder->S32(height);
    mEncoder->U32(colorFormat);
    mEncoder->U32(depthFormat);
    mEncoder->U32(stencilFormat);
}

HGLRC Spy::wglCreateContext(HDC hdc) {
#if TARGET_OS == TARGET_OS_WINDOWS
    wgl::FramebufferInfo info;
    wgl::getFramebufferInfo(hdc, info);
    init(info.width, info.height,
            info.colorFormat, info.depthFormat, info.stencilFormat);
#endif // TARGET_OS
    return GlesSpy::wglCreateContext(hdc);
}

BOOL Spy::wglMakeCurrent(HDC hdc, HGLRC hglrc) {
    BOOL res = GlesSpy::wglMakeCurrent(hdc, hglrc);
    mImports.Resolve();
    return res;
}

} // namespace gapii
