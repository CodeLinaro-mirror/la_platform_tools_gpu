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

#include "spy.h"
#include "connection_writer.h"

#include <gapic/encoder.h>
#include <gapic/log.h>
#include <gapic/target.h>
#include <gapic/thread.h>

#if TARGET_OS == GAPID_OS_WINDOWS
#include "windows/wgl.h"
#endif // TARGET_OS

namespace {

const uint32_t EGL_WIDTH  = 0x3057;
const uint32_t EGL_HEIGHT = 0x3056;

const uint32_t GLX_WIDTH  = 0x801D;
const uint32_t GLX_HEIGHT = 0x801E;

const uint32_t kCGLCPSurfaceBackingSize = 304;

bool isLittleEndian() {
    union {
        uint32_t i;
        char c[4];
    } u;
    u.i = 0x01020304;
    return u.c[0] == 4;
}

} // anonymous namespace

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
    GlesSpy::init(encoder);
    GlesSpy::architecture(alignof(void*), sizeof(void*), sizeof(int), isLittleEndian());
}

void Spy::lock() {
    SpyBase::lock();
    auto threadID = gapic::Thread::current().id();
    if (threadID != CurrentThread) {
        GAPID_INFO("Changing threads: %" PRIu64 "-> %" PRIu64 "\n", CurrentThread, threadID);
        GlesSpy::switchThread(threadID);
    }
}

EGLBoolean Spy::eglInitialize(EGLDisplay dpy, EGLint* major, EGLint* minor) {
    EGLBoolean res = GlesSpy::eglInitialize(dpy, major, minor);
    if (res != 0) {
        mImports.Resolve();
    }
    return res;
}

EGLBoolean Spy::eglMakeCurrent(EGLDisplay display, EGLSurface draw, EGLSurface read, EGLContext context) {
    using namespace GLenum;

    EGLBoolean res = GlesSpy::eglMakeCurrent(display, draw, read, context);
    if (res != 0 && draw != nullptr) {
        int width = 0;
        int height = 0;
        mImports.eglQuerySurface(display, draw, EGL_WIDTH, &width);
        mImports.eglQuerySurface(display, draw, EGL_HEIGHT, &height);

        // TODO: Probe formats
        GlesSpy::backbufferInfo(width, height, GL_RGBA8, GL_DEPTH_COMPONENT16, GL_STENCIL_INDEX8, true);
    }

    return res;
}

BOOL Spy::wglMakeCurrent(HDC hdc, HGLRC hglrc) {
    BOOL res = GlesSpy::wglMakeCurrent(hdc, hglrc);
    if (res != 0 && hglrc != nullptr) {
        mImports.Resolve();

#if TARGET_OS == GAPID_OS_WINDOWS
        wgl::FramebufferInfo info;
        wgl::getFramebufferInfo(hdc, info);
        GlesSpy::backbufferInfo(info.width, info.height,
                info.colorFormat, info.depthFormat, info.stencilFormat, true);
#endif // TARGET_OS
    }

    return res;
}

CGLError Spy::CGLSetCurrentContext(CGLContextObj ctx) {
    using namespace GLenum;

    CGLError err = GlesSpy::CGLSetCurrentContext(ctx);
    if (err == 0 && ctx != nullptr) {
        CGSConnectionID cid;
        CGSWindowID wid;
        CGSSurfaceID sid;
        double bounds[4];

        if (mImports.CGLGetSurface(ctx, &cid, &wid, &sid) == 0) {
            mImports.CGSGetSurfaceBounds(cid, wid, sid, bounds);
        }
        int width = bounds[2] - bounds[0];  // size.x - origin.x
        int height = bounds[3] - bounds[1]; // size.y - origin.y

        // TODO: Probe formats
        GlesSpy::backbufferInfo(width, height, GL_RGBA8, GL_DEPTH_COMPONENT16, GL_STENCIL_INDEX8, true);
    }
    return err;
}

Bool Spy::glXMakeContextCurrent(void* display, GLXDrawable draw, GLXDrawable read, GLXContext ctx) {
    using namespace GLenum;

    Bool res = GlesSpy::glXMakeContextCurrent(display, draw, read, ctx);
    if (res != 0 && display != nullptr) {
        int width = 0;
        int height = 0;
        mImports.glXQueryDrawable(display, draw, GLX_WIDTH, &width);
        mImports.glXQueryDrawable(display, draw, GLX_HEIGHT, &height);

        // TODO: Probe formats
        GlesSpy::backbufferInfo(width, height, GL_RGBA8, GL_DEPTH_COMPONENT16, GL_STENCIL_INDEX8, true);
    }

    return res;
}

Bool Spy::glXMakeCurrent(void* display, GLXDrawable drawable, GLXContext ctx) {
    using namespace GLenum;

    Bool res = GlesSpy::glXMakeCurrent(display, drawable, ctx);
    if (res != 0 && display != nullptr) {
        int width = 0;
        int height = 0;
        mImports.glXQueryDrawable(display, drawable, GLX_WIDTH, &width);
        mImports.glXQueryDrawable(display, drawable, GLX_HEIGHT, &height);

        // TODO: Probe formats
        GlesSpy::backbufferInfo(
            width, height, GL_RGBA8, GL_DEPTH_COMPONENT16, GL_STENCIL_INDEX8, true);
    }

    return res;
}

} // namespace gapii
