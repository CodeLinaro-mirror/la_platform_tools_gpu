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
#include "connection_writer.h"

#include <gapic/encoder.h>
#include <gapic/target.h>

#if TARGET_OS == GAPID_OS_WINDOWS
#include "windows/wgl.h"
#endif // TARGET_OS

namespace gapii {

// Use a "localabstract" pipe on Android to prevent depending on the traced application
// having the INTERNET permission set, required for opening and listening on a TCP socket.
Spy::Spy() {
#if TARGET_OS == GAPID_OS_ANDROID
    auto writer = ConnectionWriter::listenPipe("gfxspy", true);
#else // TARGET_OS
    auto writer = ConnectionWriter::listenSocket("127.0.0.1", "9286");
#endif
    auto encoder = std::shared_ptr<gapic::Encoder>(new gapic::Encoder(writer));
    mEncoder = encoder;
    GlesSpy::init(encoder);
}

Spy::~Spy() {
}

void Spy::init(int32_t width, int32_t height,
        uint32_t colorFormat, uint32_t depthFormat, uint32_t stencilFormat) {
    mImports.Resolve();
    mState.init(width, height, colorFormat, depthFormat, stencilFormat);
    mEncoder->Uint16(0); // INIT_ID
    mEncoder->Uint32(0); // ContextID
    mEncoder->Int32(width);
    mEncoder->Int32(height);
    mEncoder->Uint32(colorFormat);
    mEncoder->Uint32(depthFormat);
    mEncoder->Uint32(stencilFormat);
}

EGLBoolean Spy::eglInitialize(EGLDisplay const dpy, EGLint* const major, EGLint* const minor) {
    using namespace RenderbufferFormat;
    // TODO: Fetch dimensions and formats from OS.
    init(256, 256, GL_RGBA8, GL_DEPTH_COMPONENT16, GL_STENCIL_INDEX8);
    return GlesSpy::eglInitialize(dpy, major, minor);
}

HGLRC Spy::wglCreateContext(HDC hdc) {
#if TARGET_OS == GAPID_OS_WINDOWS
    wgl::FramebufferInfo info;
    wgl::getFramebufferInfo(hdc, info);
    init(info.width, info.height, info.colorFormat, info.depthFormat, info.stencilFormat);
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

GLXContext Spy::glXCreateContext(const void* display, const void* vis,
                            GLXContext shareList, bool direct) {
    using namespace RenderbufferFormat;
    // TODO: Fetch dimensions and formats from OS.
    init(256, 256, GL_RGBA8, GL_DEPTH_COMPONENT16, GL_STENCIL_INDEX8);
    return GlesSpy::glXCreateContext(display, vis, shareList, direct);
}

GLXContext Spy::glXCreateNewContext(const void* display, const void* fbconfig,
                                    uint32_t type, GLXContext shared, bool direct) {
    using namespace RenderbufferFormat;
    // TODO: Fetch dimensions and formats from OS.
    init(256, 256, GL_RGBA8, GL_DEPTH_COMPONENT16, GL_STENCIL_INDEX8);
    return GlesSpy::glXCreateNewContext(display, fbconfig, type, shared, direct);
}

} // namespace gapii
