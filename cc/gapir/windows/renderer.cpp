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

#include <windows.h>

namespace gapir {
namespace {

const TCHAR* wndClassName = TEXT("replayd");

WNDCLASS registerWindowClass() {
    WNDCLASS wc;
    memset(&wc, 0, sizeof(wc));

    wc.style         = 0;
    wc.lpfnWndProc   = DefWindowProc;
    wc.hInstance     = GetModuleHandle(0);
    wc.hCursor       = LoadCursor(0, IDC_ARROW); // TODO: Needed?
    wc.hbrBackground = HBRUSH(COLOR_WINDOW + 1);
    wc.lpszMenuName  = TEXT("");
    wc.lpszClassName = wndClassName;

    if (RegisterClass(&wc) == 0) {
        GAPID_FATAL("Failed to register window class. Error: %d", GetLastError());
    }

    return wc;
}

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

    HGLRC mRenderingContext;
    HDC mDeviceContext;
    HWND mWindow;
};

RendererImpl::RendererImpl()
        : mWidth(0)
        , mHeight(0)
        , mDepthSize(0)
        , mStencilSize(0)
        , mBound(false)
        , mRenderingContext(nullptr)
        , mDeviceContext(0)
        , mWindow(0) {

    // Initialize with a default target.
    setBackbuffer(8, 8, 24, 8);
}

RendererImpl::~RendererImpl() {
    reset();
}

void RendererImpl::reset() {
    unbind();

    if (mRenderingContext != nullptr) {
        if (!wglDeleteContext(mRenderingContext)) {
            GAPID_FATAL("Failed to delete GL context. Error: %d", GetLastError());
        }
        mRenderingContext = nullptr;
    }

    if (mDeviceContext != nullptr) {
        // TODO: Does this need to be released?
        mDeviceContext = nullptr;
    }

    if (mWindow != nullptr) {
        if (!DestroyWindow(mWindow)) {
            GAPID_FATAL("Failed to destroy window. Error: %d", GetLastError());
        }
        mWindow = nullptr;
    }

    mWidth = 0;
    mHeight = 0;
    mDepthSize = 0;
    mStencilSize = 0;
}

void RendererImpl::setBackbuffer(int width, int height, int depthSize, int stencilSize) {
    if (mDeviceContext != nullptr &&
        mWidth == width &&
        mHeight == height &&
        mDepthSize == depthSize &&
        mStencilSize == stencilSize) {
        // No change
        return;
    }

    if (mWindow != nullptr &&
        mDepthSize == depthSize &&
        mStencilSize == stencilSize) {
        // Resize only
        GAPID_INFO("Resizing renderer: %dx%d -> %dx%d\n", mWidth, mHeight, width, height);
        SetWindowPos(mWindow, nullptr, 0, 0, width, height, SWP_NOMOVE);
        mWidth = width;
        mHeight = height;
        return;
    }

    const bool wasBound = mBound;

    reset();

    static WNDCLASS wc = registerWindowClass(); // Only needs to be done once per app life-time.

    mWindow = CreateWindow(wndClassName, TEXT(""), WS_POPUP, 0, 0,
            width, height, 0, 0, GetModuleHandle(0), 0);
    if (mWindow == 0) {
        GAPID_FATAL("Failed to create window. Error: %d", GetLastError());
    }

    PIXELFORMATDESCRIPTOR pfd;
    memset(&pfd, 0, sizeof(pfd));

    pfd.nSize = sizeof(PIXELFORMATDESCRIPTOR);
    pfd.nVersion = 1;
    pfd.dwFlags = PFD_DRAW_TO_WINDOW | PFD_SUPPORT_OPENGL;
    pfd.iPixelType = PFD_TYPE_RGBA;
    pfd.cColorBits = 32;
    pfd.cDepthBits = depthSize;
    pfd.cStencilBits = stencilSize;
    pfd.cAlphaBits = 8;
    pfd.iLayerType = PFD_MAIN_PLANE;

    mDeviceContext = GetDC(mWindow);

    int pixelFormat = ChoosePixelFormat(mDeviceContext, &pfd);
    SetPixelFormat(mDeviceContext, pixelFormat, &pfd);

    mRenderingContext = wglCreateContext(mDeviceContext);
    if (mRenderingContext == nullptr) {
        GAPID_FATAL("Failed to create GL context. Error: %d", GetLastError());
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
        if (!wglMakeCurrent(mDeviceContext, mRenderingContext)) {
            GAPID_FATAL("Failed to attach GL context. Error: %d", GetLastError());
        }

        mBound = true;

        // Initialize the graphics API
        // TODO: Inefficient - consider moving the imports into this renderer
        gfxapi::Initialize();
    }
}

void RendererImpl::unbind() {
    if (mBound) {
        if (!wglMakeCurrent(mDeviceContext, nullptr)) {
            GAPID_FATAL("Failed to detach GL context. Error: %d", GetLastError());
        }

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
