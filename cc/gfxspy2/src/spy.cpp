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

#include <gapic/file_writer.h>
#include <gapic/target.h>

#if TARGET_OS == GAPID_OS_WINDOWS
#include "windows/wgl.h"
#endif // TARGET_OS

namespace gapii {

Spy::Spy() {
}

Spy::~Spy() {
    mEncoder->U16(0xffff); // Type ID -- TODO: mEncoder->Id(EOS_ID);
}

void Spy::init(int32_t width, int32_t height,
        uint32_t colorFormat, uint32_t depthFormat, uint32_t stencilFormat) {
    auto writer = std::shared_ptr<gapic::StreamWriter>(new gapic::FileWriter("atoms"));
    auto encoder = std::shared_ptr<gapic::Encoder>(new gapic::Encoder(writer));

    mEncoder = encoder;
    GlesSpy::init(encoder);

    mState.init(width, height, colorFormat, depthFormat, stencilFormat);
    encoder->U16(0); // INIT_ID
    encoder->U32(0); // ContextID
    encoder->S32(width);
    encoder->S32(height);
    encoder->U32(colorFormat);
    encoder->U32(depthFormat);
    encoder->U32(stencilFormat);
}

void Spy::eglInitialize(EGLDisplay display, int32_t* major, int32_t* minor) {
    using namespace RenderbufferFormat;
    // TODO: Fetch dimensions and formats from OS.
    init(256, 256, GL_RGBA8, GL_DEPTH_COMPONENT16, GL_STENCIL_INDEX8);
    GlesSpy::eglInitialize(display, major, minor);
}

void Spy::eglMakeCurrent(int32_t context) {
    GlesSpy::eglMakeCurrent(context);
    mImports.Resolve();
}

HGLRC Spy::wglCreateContext(HDC hdc) {
#if TARGET_OS == GAPID_OS_WINDOWS
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

CGLError Spy::CGLCreateContext(CGLPixelFormatObj pix, CGLContextObj share, CGLContextObj ctx) {
    using namespace RenderbufferFormat;
    // TODO: Fetch dimensions and formats from OS.
    init(256, 256, GL_RGBA8, GL_DEPTH_COMPONENT16, GL_STENCIL_INDEX8);
    return GlesSpy::CGLCreateContext(pix, share, ctx);
}

} // namespace gapii
