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
#include "Target.h"

#if TARGET_OS == CAZE_OS_ANDROID

#include "GfxApi.h"
#include "Log.h"
#include "Renderer.h"

#include <EGL/egl.h>

namespace android {
namespace caze {
namespace {

class RendererImpl : public Renderer {
public:
    RendererImpl(int width, int height, int depthSize, int stencilSize);
    virtual ~RendererImpl() override;

    virtual const char* name() override;
    virtual const char* extensions() override;
    virtual const char* vendor() override;
    virtual const char* version() override;

private:
    EGLContext mEglContext;
    EGLSurface mEglSurface;
    EGLDisplay mEglDisplay;
};

RendererImpl::RendererImpl(int width, int height, int depthSize, int stencilSize) {
    EGLint error;

    mEglDisplay = eglGetDisplay(EGL_DEFAULT_DISPLAY);
    error = eglGetError();
    if (error != EGL_SUCCESS) {
        CAZE_FATAL("Failed to get EGL display: %d\n", error);
    }

    eglInitialize(mEglDisplay, nullptr, nullptr);
    error = eglGetError();
    if (error != EGL_SUCCESS) {
        CAZE_FATAL("Failed to initialize EGL: %d\n", error);
    }

    eglBindAPI(EGL_OPENGL_ES_API);
    error = eglGetError();
    if (error != EGL_SUCCESS) {
        CAZE_FATAL("Failed to bind EGL API: %d\n", error);
    }

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
    eglChooseConfig(mEglDisplay, configAttribList, &eglConfig, 1, &one);
    error = eglGetError();
    if (error != EGL_SUCCESS) {
        CAZE_FATAL("Failed to choose EGL config: %d\n", error);
    }

    // Create an EGL context.
    const int contextAttribList[] = {
        EGL_CONTEXT_CLIENT_VERSION, 2,
        EGL_NONE
    };
    mEglContext = eglCreateContext(mEglDisplay, eglConfig, EGL_NO_CONTEXT, contextAttribList);
    error = eglGetError();
    if (error != EGL_SUCCESS) {
        CAZE_FATAL("Failed to create EGL context: %d\n", error);
    }

    // Create an EGL surface for the read/draw framebuffer.
    const int surfaceAttribList[] = {
        EGL_WIDTH, width,
        EGL_HEIGHT, height,
        EGL_NONE
    };
    mEglSurface = eglCreatePbufferSurface(mEglDisplay, eglConfig, surfaceAttribList);
    error = eglGetError();
    if (error != EGL_SUCCESS) {
        CAZE_FATAL("Failed to create EGL pbuffer surface: %d\n", error);
    }

    eglMakeCurrent(mEglDisplay, mEglSurface, mEglSurface, mEglContext);
    error = eglGetError();
    if (error != EGL_SUCCESS) {
        CAZE_FATAL("Failed to make EGL current: %d\n", error);
    }

    // Initialize the graphics API
    gfxapi::Initialize();
}

RendererImpl::~RendererImpl() {
    EGLint error;

    eglMakeCurrent(mEglDisplay, EGL_NO_SURFACE, EGL_NO_SURFACE, EGL_NO_CONTEXT);
    error = eglGetError();
    if (error != EGL_SUCCESS) {
        CAZE_WARNING("Failed to release EGL context: %d\n", error);
    }

    eglDestroySurface(mEglDisplay, mEglSurface);
    error = eglGetError();
    if (error != EGL_SUCCESS) {
        CAZE_WARNING("Failed to destroy EGL surface: %d\n", error);
    }

    eglDestroyContext(mEglDisplay, mEglContext);
    error = eglGetError();
    if (error != EGL_SUCCESS) {
        CAZE_WARNING("Failed to destroy EGL context: %d\n", error);
    }

    eglTerminate(mEglDisplay);
    error = eglGetError();
    if (error != EGL_SUCCESS) {
        CAZE_WARNING("Failed to terminate EGL: %d\n", error);
    }

    eglReleaseThread();
    error = eglGetError();
    if (error != EGL_SUCCESS) {
        CAZE_WARNING("Failed to release EGL thread: %d\n", error);
    }
}

const char* RendererImpl::name() {
    return eglQueryString(mEglDisplay, EGL_EXTENSIONS);
}

const char* RendererImpl::extensions() {
    return eglQueryString(mEglDisplay, EGL_CLIENT_APIS);
}

const char* RendererImpl::vendor() {
    return eglQueryString(mEglDisplay, EGL_VERSION);
}

const char* RendererImpl::version() {
    return eglQueryString(mEglDisplay, EGL_VENDOR);
}

} // end of anonymous namespace

std::unique_ptr<Renderer> Renderer::create(int width, int height, int depthSize, int stencilSize) {
    return std::unique_ptr<Renderer>(new RendererImpl(width, height, depthSize, stencilSize));
}

}  // end of namespace caze
}  // end of namespace android

#endif // TARGET_OS == CAZE_OS_ANDROID

