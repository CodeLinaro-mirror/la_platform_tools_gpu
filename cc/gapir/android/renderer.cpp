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

#include "../gfx_api.h"
#include "../renderer.h"

#include <gapic/log.h>

#include <EGL/egl.h>

namespace gapir {
namespace {

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

    int mWidth;
    int mHeight;
    int mDepthSize;
    int mStencilSize;
    bool mBound;

    EGLContext mContext;
    EGLSurface mSurface;
    EGLDisplay mDisplay;
};

RendererImpl::RendererImpl()
        : mWidth(0)
        , mHeight(0)
        , mDepthSize(0)
        , mStencilSize(0)
        , mBound(false)
        , mContext(EGL_NO_CONTEXT)
        , mSurface(EGL_NO_SURFACE) {

    mDisplay = eglGetDisplay(EGL_DEFAULT_DISPLAY);
    EGLint error = eglGetError();
    if (error != EGL_SUCCESS) {
        GAPID_FATAL("Failed to get EGL display: %d\n", error);
    }

    eglInitialize(mDisplay, nullptr, nullptr);
    error = eglGetError();
    if (error != EGL_SUCCESS) {
        GAPID_FATAL("Failed to initialize EGL: %d\n", error);
    }

    eglBindAPI(EGL_OPENGL_ES_API);
    error = eglGetError();
    if (error != EGL_SUCCESS) {
        GAPID_FATAL("Failed to bind EGL API: %d\n", error);
    }

    // Initialize with a default target.
    setBackbuffer(8, 8, 24, 8);
}

RendererImpl::~RendererImpl() {
    reset();

    eglTerminate(mDisplay);
    EGLint error = eglGetError();
    if (error != EGL_SUCCESS) {
        GAPID_WARNING("Failed to terminate EGL: %d\n", error);
    }

    eglReleaseThread();
    error = eglGetError();
    if (error != EGL_SUCCESS) {
        GAPID_WARNING("Failed to release EGL thread: %d\n", error);
    }
}

void RendererImpl::reset() {
    unbind();

    if (mSurface != EGL_NO_SURFACE) {
        eglDestroySurface(mDisplay, mSurface);
        EGLint error = eglGetError();
        if (error != EGL_SUCCESS) {
            GAPID_WARNING("Failed to destroy EGL surface: %d\n", error);
        }
        mSurface = EGL_NO_SURFACE;
    }

    if (mContext != EGL_NO_CONTEXT) {
        eglDestroyContext(mDisplay, mContext);
        EGLint error = eglGetError();
        if (error != EGL_SUCCESS) {
            GAPID_WARNING("Failed to destroy EGL context: %d\n", error);
        }
        mContext = EGL_NO_CONTEXT;
    }

    mWidth = 0;
    mHeight = 0;
    mDepthSize = 0;
    mStencilSize = 0;
}

void RendererImpl::setBackbuffer(int width, int height, int depthSize, int stencilSize) {
    if (mContext != EGL_NO_CONTEXT &&
        mWidth == width &&
        mHeight == height &&
        mDepthSize == depthSize &&
        mStencilSize == stencilSize) {

        return;
    }

    const bool wasBound = mBound;

    reset();

    // Find a supported EGL context config.
    const int configAttribList[] = {
        EGL_RED_SIZE, 8,
        EGL_GREEN_SIZE, 8,
        EGL_BLUE_SIZE, 8,
        EGL_ALPHA_SIZE, 8,
        EGL_BUFFER_SIZE, 32,
        EGL_DEPTH_SIZE, depthSize,
        EGL_STENCIL_SIZE, stencilSize,
        EGL_SURFACE_TYPE, EGL_PBUFFER_BIT,
        EGL_RENDERABLE_TYPE, EGL_OPENGL_ES2_BIT,
        EGL_NONE
    };
    int one = 1;
    EGLConfig eglConfig;
    eglChooseConfig(mDisplay, configAttribList, &eglConfig, 1, &one);
    EGLint error = eglGetError();
    if (error != EGL_SUCCESS) {
        GAPID_FATAL("Failed to choose EGL config: %d\n", error);
    }

    // Create an EGL context.
    const int contextAttribList[] = {
        EGL_CONTEXT_CLIENT_VERSION, 2,
        EGL_NONE
    };
    mContext = eglCreateContext(mDisplay, eglConfig, EGL_NO_CONTEXT, contextAttribList);
    error = eglGetError();
    if (error != EGL_SUCCESS) {
        GAPID_FATAL("Failed to create EGL context: %d\n", error);
    }

    // Create an EGL surface for the read/draw framebuffer.
    const int surfaceAttribList[] = {
        EGL_WIDTH, width,
        EGL_HEIGHT, height,
        EGL_NONE
    };
    mSurface = eglCreatePbufferSurface(mDisplay, eglConfig, surfaceAttribList);
    error = eglGetError();
    if (error != EGL_SUCCESS) {
        GAPID_FATAL("Failed to create EGL pbuffer surface: %d\n", error);
    }

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
        eglMakeCurrent(mDisplay, mSurface, mSurface, mContext);
        EGLint error = eglGetError();
        if (error != EGL_SUCCESS) {
            GAPID_FATAL("Failed to make EGL current: %d\n", error);
        }

        mBound = true;

        // Initialize the graphics API
        // TODO: Inefficient - consider moving the imports into this renderer
        gfxapi::Initialize();
    }
}

void RendererImpl::unbind() {
    if (mBound) {
        eglMakeCurrent(mDisplay, EGL_NO_SURFACE, EGL_NO_SURFACE, EGL_NO_CONTEXT);
        EGLint error = eglGetError();
        if (error != EGL_SUCCESS) {
            GAPID_WARNING("Failed to release EGL context: %d\n", error);
        }
        mBound = false;
    }
}

const char* RendererImpl::name() {
    return eglQueryString(mDisplay, EGL_EXTENSIONS);
}

const char* RendererImpl::extensions() {
    return eglQueryString(mDisplay, EGL_CLIENT_APIS);
}

const char* RendererImpl::vendor() {
    return eglQueryString(mDisplay, EGL_VERSION);
}

const char* RendererImpl::version() {
    return eglQueryString(mDisplay, EGL_VENDOR);
}

} // anonymous namespace

Renderer* Renderer::create() {
    return new RendererImpl();
}

}  // namespace gapir

