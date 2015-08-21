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

#include "../gfx_api.h"
#include "../renderer.h"

#include <gapic/log.h>

#include <cstring>
#include <X11/Xresource.h>

namespace gapir {
namespace {

typedef XID GLXPbuffer;
typedef XID GLXDrawable;
typedef /*struct __GLXcontextRec*/ void *GLXContext;
typedef /*struct __GLXFBConfigRec*/ void *GLXFBConfig;

enum {
    // Used by glXChooseFBConfig.
    GLX_RED_SIZE      = 8,
    GLX_GREEN_SIZE    = 9,
    GLX_BLUE_SIZE     = 10,
    GLX_ALPHA_SIZE    = 11,
    GLX_DEPTH_SIZE    = 12,
    GLX_STENCIL_SIZE  = 13,
    GLX_DRAWABLE_TYPE = 0x8010,
    GLX_RENDER_TYPE   = 0x8011,
    GLX_RGBA_BIT      = 0x00000001,
    GLX_PBUFFER_BIT   = 0x00000004,

    // Used by glXCreateNewContext.
    GLX_RGBA_TYPE = 0x8014,

    // Used by glXCreatePbuffer.
    GLX_PBUFFER_HEIGHT = 0x8040,
    GLX_PBUFFER_WIDTH  = 0x8041
};

extern "C" {

GLXFBConfig *glXChooseFBConfig(Display *dpy, int screen, const int *attrib_list, int *nelements);
GLXContext glXCreateNewContext(Display *dpy, GLXFBConfig config, int render_type,
                               GLXContext share_list, Bool direct);
GLXPbuffer glXCreatePbuffer(Display *dpy, GLXFBConfig config, const int *attrib_list);
void glXDestroyPbuffer(Display *dpy, GLXPbuffer pbuf);
Bool glXMakeContextCurrent(Display *dpy, GLXDrawable draw, GLXDrawable read, GLXContext ctx);
Bool glXQueryVersion(Display *dpy, int *maj, int *min);
void glXDestroyContext(Display *dpy, GLXContext ctx);

} // extern "C"

class RendererImpl : public Renderer {
public:
    RendererImpl();
    virtual ~RendererImpl() override;

    virtual void setBackbuffer(int width, int height, int depthSize, int stencilSize);
    virtual void bind() override;
    virtual void unbind() override;
    virtual const char* name() override;
    virtual const char* extensions() override;
    virtual const char* vendor() override;
    virtual const char* version() override;

private:
    void reset();
    void createPbuffer(int width, int height);

    int mWidth;
    int mHeight;
    int mDepthSize;
    int mStencilSize;
    bool mBound;

    Display *mDisplay;
    GLXContext mContext;
    GLXPbuffer mPbuffer;
    GLXFBConfig mFBConfig;
};

RendererImpl::RendererImpl()
        : mWidth(0)
        , mHeight(0)
        , mDepthSize(0)
        , mStencilSize(0)
        , mBound(false)
        , mDisplay(nullptr)
        , mContext(nullptr)
        , mPbuffer(0) {

    mDisplay = XOpenDisplay(nullptr);
    if (mDisplay == nullptr) {
        GAPID_FATAL("Unable to to open X display");
    }

    int major;
    int minor;
    if (!glXQueryVersion(mDisplay, &major, &minor) || (major == 1 && minor < 3)) {
        GAPID_FATAL("GLX 1.3+ unsupported by X server (was %d.%d)", major, minor);
    }

    // Initialize with a default target.
    setBackbuffer(8, 8, 24, 8);
}

RendererImpl::~RendererImpl() {
    reset();

    if (mDisplay != nullptr) {
        XCloseDisplay(mDisplay);
    }
}

void RendererImpl::reset() {
    unbind();

    if (mContext != nullptr) {
        glXDestroyContext(mDisplay, mContext);
        mContext = nullptr;
    }

    if (mPbuffer != 0) {
        glXDestroyPbuffer(mDisplay, mPbuffer);
        mPbuffer = 0;
    }

    mWidth = 0;
    mHeight = 0;
    mDepthSize = 0;
    mStencilSize = 0;
}

void RendererImpl::createPbuffer(int width, int height) {
    if (mPbuffer != 0) {
        glXDestroyPbuffer(mDisplay, mPbuffer);
        mPbuffer = 0;
    }
    const int pbufferAttribs[] = {
        GLX_PBUFFER_WIDTH, width,
        GLX_PBUFFER_HEIGHT, height,
        None
    };
    mPbuffer = glXCreatePbuffer(mDisplay, mFBConfig, pbufferAttribs);
}

void RendererImpl::setBackbuffer(int width, int height, int depthSize, int stencilSize) {
    if (mContext != nullptr &&
        mWidth == width &&
        mHeight == height &&
        mDepthSize == depthSize &&
        mStencilSize == stencilSize) {
        // No change
        return;
    }

    if (mContext != nullptr &&
        mDepthSize == depthSize &&
        mStencilSize == stencilSize) {
        // Resize only
        GAPID_INFO("Resizing renderer: %dx%d -> %dx%d", mWidth, mHeight, width, height);
        createPbuffer(width, height);
        glXMakeContextCurrent(mDisplay, mPbuffer, mPbuffer, mContext);
        mWidth = width;
        mHeight = height;
        return;
    }

    const bool wasBound = mBound;

    reset();

    const int visualAttribs[] = {
        GLX_RED_SIZE, 8,
        GLX_GREEN_SIZE, 8,
        GLX_BLUE_SIZE, 8,
        GLX_ALPHA_SIZE, 8,
        GLX_DEPTH_SIZE, depthSize,
        GLX_STENCIL_SIZE, stencilSize,
        GLX_RENDER_TYPE, GLX_RGBA_BIT,
        GLX_DRAWABLE_TYPE, GLX_PBUFFER_BIT,
        None
    };
    int fbConfigsCount;
    GLXFBConfig *fbConfigs = glXChooseFBConfig(
            mDisplay, DefaultScreen(mDisplay), visualAttribs, &fbConfigsCount);
    if (fbConfigs == nullptr) {
        GAPID_FATAL("Unable to find a suitable X framebuffer config");
    }
    mFBConfig = fbConfigs[0];
    XFree(fbConfigs);

    mContext = glXCreateNewContext(mDisplay, mFBConfig, GLX_RGBA_TYPE, nullptr, True);
    if (mContext == nullptr) {
        GAPID_FATAL("Failed to create glX context");
    }
    XSync(mDisplay, False);

    createPbuffer(width, height);

    mWidth = width;
    mHeight = height;
    mDepthSize = depthSize;
    mStencilSize = stencilSize;

    if (wasBound) {
        bind();
    }
}

void RendererImpl::bind() {
    if (!mBound) {
        if (!glXMakeContextCurrent(mDisplay, mPbuffer, mPbuffer, mContext)) {
            GAPID_FATAL("Unable to make GLX context current");
        }

        mBound = true;

        // Initialize the graphics API
        // TODO: Inefficient - consider moving the imports into this renderer
        gfxapi::Initialize();
    }
}

void RendererImpl::unbind() {
    if (mBound) {
        // TODO: glXMakeContextCurrent(...)
        mBound = false;
    }
}

const char* RendererImpl::name() {
    return reinterpret_cast<const char*>(
        gfxapi::glGetString(gfxapi::GLenum::GL_RENDERER));
}

const char* RendererImpl::extensions() {
    return reinterpret_cast<const char*>(
        gfxapi::glGetString(gfxapi::GLenum::GL_EXTENSIONS));
}

const char* RendererImpl::vendor() {
    return reinterpret_cast<const char*>(
        gfxapi::glGetString(gfxapi::GLenum::GL_VENDOR));
}

const char* RendererImpl::version() {
    return reinterpret_cast<const char*>(
        gfxapi::glGetString(gfxapi::GLenum::GL_VERSION));
}

} // anonymous namespace

Renderer* Renderer::create() {
    return new RendererImpl();
}

}  // namespace gapir

