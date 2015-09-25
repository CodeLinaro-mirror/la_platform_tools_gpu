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

#include "connection_header.h"
#include "connection_stream.h"
#include "spy.h"

#include <gapic/encoder.h>
#include <gapic/log.h>
#include <gapic/target.h>
#include <gapic/thread.h>

#if TARGET_OS == GAPID_OS_WINDOWS
#include "windows/wgl.h"
#endif // TARGET_OS

namespace {

const uint32_t EGL_WIDTH            = 0x3057;
const uint32_t EGL_HEIGHT           = 0x3056;
const uint32_t EGL_SWAP_BEHAVIOR    = 0x3093;
const uint32_t EGL_BUFFER_PRESERVED = 0x3094;

const uint32_t GLX_WIDTH  = 0x801D;
const uint32_t GLX_HEIGHT = 0x801E;

const uint32_t kCGLCPSurfaceBackingSize = 304;

inline bool isLittleEndian() {
    union {
        uint32_t i;
        char c[4];
    } u;
    u.i = 0x01020304;
    return u.c[0] == 4;
}

} // anonymous namespace

namespace gapii {

Spy::Spy()
  : mObserveFramebufferOnEOF(false)
  , mObserveFramebufferOnDrawCall(false) {

#if TARGET_OS == GAPID_OS_ANDROID
  // Use a "localabstract" pipe on Android to prevent depending on the traced application
  // having the INTERNET permission set, required for opening and listening on a TCP socket.
    auto conn = ConnectionStream::listenPipe("gapii", true);
#else // TARGET_OS
    auto conn = ConnectionStream::listenSocket("127.0.0.1", "9286");
#endif // TARGET_OS

    ConnectionHeader header;
    if (header.read(conn.get())) {
        mObserveFramebufferOnEOF = header.mObserveFramebufferOnEOF != 0;
        mObserveFramebufferOnDrawCall = header.mObserveFramebufferOnDrawCall != 0;
    } else {
        GAPID_WARNING("Failed to read connection header");
    }

    GAPID_INFO("GAPII connection established. Settings:");
    GAPID_INFO("Observe framebuffers on EOF:       %s", mObserveFramebufferOnEOF ? "yes" : "no");
    GAPID_INFO("Observe framebuffers on draw call: %s", mObserveFramebufferOnDrawCall ? "yes" : "no");

    mEncoder = std::shared_ptr<gapic::Encoder>(new gapic::Encoder(conn));
    mEncoder->String("GapiiTraceFile_V1.0");
    GlesSpy::init(mEncoder);
    GlesSpy::architecture(alignof(void*), sizeof(void*), sizeof(int), isLittleEndian());
}

void Spy::lock() {
    SpyBase::lock();
    auto threadID = gapic::Thread::current().id();
    if (threadID != CurrentThread) {
        GAPID_INFO("Changing threads: %" PRIu64 "-> %" PRIu64, CurrentThread, threadID);
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
        int swapBehavior = 0;
        mImports.eglQuerySurface(display, draw, EGL_WIDTH, &width);
        mImports.eglQuerySurface(display, draw, EGL_HEIGHT, &height);
        mImports.eglQuerySurface(display, draw, EGL_SWAP_BEHAVIOR, &swapBehavior);

        bool resetViewportScissor = true;
        bool preserveBuffersOnSwap = swapBehavior == EGL_BUFFER_PRESERVED;

        // TODO: Probe formats
        setContextInfo(width, height,
                GL_RGBA8, GL_DEPTH_COMPONENT16, GL_STENCIL_INDEX8,
                resetViewportScissor,
                preserveBuffersOnSwap);
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
        setContextInfo(info.width, info.height,
                info.colorFormat, info.depthFormat, info.stencilFormat,
                /* resetViewportScissor */ true,
                /* preserveBuffersOnSwap */ false);
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
        double bounds[4] = {0, 0, 0, 0};

        if (mImports.CGLGetSurface(ctx, &cid, &wid, &sid) == 0) {
            mImports.CGSGetSurfaceBounds(cid, wid, sid, bounds);
        } else {
            GAPID_WARNING("Could not get CGL surface");
        }
        int width = bounds[2] - bounds[0];  // size.x - origin.x
        int height = bounds[3] - bounds[1]; // size.y - origin.y

        // TODO: Probe formats
        setContextInfo(width, height,
                GL_RGBA8, GL_DEPTH_COMPONENT16, GL_STENCIL_INDEX8,
                /* resetViewportScissor */ true,
                /* preserveBuffersOnSwap */ false);
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
        setContextInfo(width, height,
                GL_RGBA8, GL_DEPTH_COMPONENT16, GL_STENCIL_INDEX8,
                /* resetViewportScissor */ true,
                /* preserveBuffersOnSwap */ false);
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
        setContextInfo(width, height,
                GL_RGBA8, GL_DEPTH_COMPONENT16, GL_STENCIL_INDEX8,
                /* resetViewportScissor */ true,
                /* preserveBuffersOnSwap */ false);
    }

    return res;
}

void Spy::setContextInfo(int32_t backbuffer_width, int32_t backbuffer_height,
                         uint32_t backbuffer_color_fmt, uint32_t backbuffer_depth_fmt,
                         uint32_t backbuffer_stencil_fmt, bool reset_viewport_scissor,
                         bool preserve_buffers_on_swap) {
    std::shared_ptr<Context> ctx = GlesSpy::Contexts[GlesSpy::CurrentThread];
    char* name = reinterpret_cast<char*>(mImports.glGetString(GLenum::GL_RENDERER));
    char* vendor = reinterpret_cast<char*>(mImports.glGetString(GLenum::GL_VENDOR));
    char* extensions = ""; // TODO
    char* version = reinterpret_cast<char*>(mImports.glGetString(GLenum::GL_VERSION));
    GlesSpy::contextInfo(name, vendor, extensions, version,
                         backbuffer_width, backbuffer_height,
                         backbuffer_color_fmt, backbuffer_depth_fmt,
                         backbuffer_stencil_fmt, reset_viewport_scissor,
                         preserve_buffers_on_swap);
}

int Spy::eglSwapBuffers(void* display, void* surface) {
    if (mObserveFramebufferOnEOF) { observeFramebuffer(); }
    return GlesSpy::eglSwapBuffers(display, surface);
}

void Spy::wglSwapBuffers(void* hdc) {
    if (mObserveFramebufferOnEOF) { observeFramebuffer(); }
    GlesSpy::wglSwapBuffers(hdc);
}

void Spy::glXSwapBuffers(void* display, void* drawable) {
    if (mObserveFramebufferOnEOF) { observeFramebuffer(); }
    GlesSpy::glXSwapBuffers(display, drawable);
}

int Spy::CGLFlushDrawable(void* ctx) {
    if (mObserveFramebufferOnEOF) { observeFramebuffer(); }
    return GlesSpy::CGLFlushDrawable(ctx);
}

// observeFramebuffer captures the currently bound framebuffer, and writes
// it to a FramebufferObservation atom.
void Spy::observeFramebuffer() {
    uint32_t w = 0;
    uint32_t h = 0;
    if (!getFramebufferAttachmentSize(w, h)) {
        return; // Could not get the framebuffer size.
    }
    uint32_t size = w * h * 4;
    uint8_t* data = new uint8_t[size];
    if (data != nullptr) {
        mImports.glReadPixels(0, 0, int32_t(w), int32_t(h),
                GLenum::GL_RGBA, GLenum::GL_UNSIGNED_BYTE, data);
        gapic::coder::atom::FramebufferObservation coder(w, h, gapic::Array<uint8_t>(data, size));
        mEncoder->Variant(&coder);
        delete [] data;
    } else {
        GAPID_WARNING("Failed to allocate buffer to observe framebuffer of size %ux%u", w, h);
    }
}

// TODO: When gfx api macros produce functions instead of inlining, move this logic
// to the gles.api file.
bool Spy::getFramebufferAttachmentSize(uint32_t& width, uint32_t& height) {
    std::shared_ptr<Context> ctx = GlesSpy::Contexts[GlesSpy::CurrentThread];
    if (ctx == nullptr) {
      return false;
    }

    auto framebufferID = ctx->mBoundFramebuffers.find(GLenum::GL_READ_FRAMEBUFFER);
    if (framebufferID == ctx->mBoundFramebuffers.end()) {
        return false;
    }

    auto framebuffer = ctx->mInstances.mFramebuffers.find(framebufferID->second);
    if (framebuffer == ctx->mInstances.mFramebuffers.end()) {
        return false;
    }

    auto attachment = framebuffer->second->mAttachments.find(GLenum::GL_COLOR_ATTACHMENT0);
    if (attachment == framebuffer->second->mAttachments.end()) {
        return false;
    }

    switch (attachment->second.mType) {
        case GLenum::GL_TEXTURE: {
            auto t = ctx->mInstances.mTextures.find(attachment->second.mObject);
            if (t == ctx->mInstances.mTextures.end()) {
                return false;
            }
            switch (t->second->mKind) {
                case TextureKind::TEXTURE2D: {
                    auto l = t->second->mTexture2D.find(attachment->second.mTextureLevel);
                    if (l == t->second->mTexture2D.end()) {
                        return false;
                    }
                    width = uint32_t(l->second.mWidth);
                    height = uint32_t(l->second.mHeight);
                    return true;
                }
                case TextureKind::CUBEMAP: {
                    auto l = t->second->mCubemap.find(attachment->second.mTextureLevel);
                    if (l == t->second->mCubemap.end()) {
                        return false;
                    }
                    auto f = l->second.mFaces.find(attachment->second.mCubeMapFace);
                    if (f == l->second.mFaces.end()) {
                        return false;
                    }
                    width = uint32_t(f->second.mWidth);
                    height = uint32_t(f->second.mHeight);
                    return true;
                }
            }
        }
        case GLenum::GL_RENDERBUFFER: {
            auto r = ctx->mInstances.mRenderbuffers.find(attachment->second.mObject);
            if (r == ctx->mInstances.mRenderbuffers.end()) {
                return false;
            }
            width = uint32_t(r->second->mWidth);
            height = uint32_t(r->second->mHeight);
            return true;
        }
    }
    return false;
}

} // namespace gapii
