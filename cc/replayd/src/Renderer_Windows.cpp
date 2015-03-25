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

#if TARGET_OS == CAZE_OS_WINDOWS

#include "GfxApi.h"
#include "Log.h"
#include "Renderer.h"

#include <windows.h>

namespace android {
namespace caze {
namespace {

const TCHAR* wndClassName = TEXT("replayd");
int gDepthSize = 0;
int gStencilSize = 0;

HGLRC renderingContext = nullptr;

LRESULT CALLBACK WndProc(HWND hWnd, UINT message, WPARAM wParam, LPARAM lParam) {
    switch(message) {
        case WM_CREATE: {
            PIXELFORMATDESCRIPTOR pfd;
            memset(&pfd, 0, sizeof(pfd));

            pfd.nSize = sizeof(PIXELFORMATDESCRIPTOR);
            pfd.nVersion = 1;
            pfd.dwFlags = PFD_DRAW_TO_WINDOW | PFD_SUPPORT_OPENGL | PFD_DOUBLEBUFFER;
            pfd.iPixelType = PFD_TYPE_RGBA;
            pfd.cColorBits = 32;
            pfd.cDepthBits = gDepthSize;
            pfd.cStencilBits = gStencilSize;
            pfd.cAlphaBits = 8;
            pfd.iLayerType = PFD_MAIN_PLANE;

            HDC deviceContext = GetDC(hWnd);

            int pixelFormat;
            pixelFormat = ChoosePixelFormat(deviceContext, &pfd);
            SetPixelFormat(deviceContext, pixelFormat, &pfd);

            renderingContext = wglCreateContext(deviceContext);
            if (renderingContext == nullptr) {
                CAZE_FATAL("Failed to create GL context. Error: %d", GetLastError());
            }
            if (!wglMakeCurrent(deviceContext, renderingContext)) {
                CAZE_FATAL("Failed to attach GL context. Error: %d", GetLastError());
            }
            break;
        }
        case WM_DESTROY: {
            HDC deviceContext = GetDC(hWnd);
            if (!wglMakeCurrent(deviceContext, nullptr)) {
                CAZE_FATAL("Failed to detach GL context. Error: %d", GetLastError());
            }
            if (!wglDeleteContext(renderingContext)) {
                CAZE_FATAL("Failed to delete GL context. Error: %d", GetLastError());
            }
            renderingContext = nullptr;
            break;
        }
    default:
        return DefWindowProc(hWnd, message, wParam, lParam);
    }
    return 0;
}

WNDCLASS registerWindowClass() {
    WNDCLASS wc;
    memset(&wc, 0, sizeof(wc));

    wc.style         = 0;
    wc.lpfnWndProc   = WndProc;
    wc.hInstance     = GetModuleHandle(0);
    wc.hCursor       = LoadCursor(0, IDC_ARROW); // TODO: Needed?
    wc.hbrBackground = HBRUSH(COLOR_WINDOW + 1);
    wc.lpszMenuName  = TEXT("");
    wc.lpszClassName = wndClassName;

    if (RegisterClass(&wc) == 0) {
        CAZE_FATAL("Failed to register window class. Error: %d", GetLastError());
    }

    return wc;
}

class RendererImpl : public Renderer {
public:
    RendererImpl(int width, int height, int depthSize, int stencilSize);
    virtual ~RendererImpl() override;

    virtual const char* name() override;
    virtual const char* extensions() override;
    virtual const char* vendor() override;
    virtual const char* version() override;

private:
    HWND mWindow;
};

RendererImpl::RendererImpl(int width, int height, int depthSize, int stencilSize) {
    if (renderingContext != 0) {
        CAZE_FATAL("Renderer already created. Only one instance can be created at any given time.\n");
    }
    gDepthSize = depthSize;
    gStencilSize = stencilSize;

    static WNDCLASS wc = registerWindowClass(); // Only needs to be done once per app life-time.

    mWindow = CreateWindow(wndClassName, TEXT(""), WS_POPUP, 0, 0, width, height, 0, 0, GetModuleHandle(0), 0);
    if (mWindow == 0) {
        CAZE_FATAL("Failed to create window. Error: %d", GetLastError());
    }

    // Initialize the graphics API
    gfxapi::Initialize();
}

RendererImpl::~RendererImpl() {
    if (!DestroyWindow(mWindow)) {
        CAZE_FATAL("Failed to destroy window. Error: %d", GetLastError());
    }
}

const char* RendererImpl::name() {
    return reinterpret_cast<const char*>(
        gfxapi::glGetString(gfxapi::StringConstant::GL_RENDERER));
}

const char* RendererImpl::extensions() {
    return reinterpret_cast<const char*>(
        gfxapi::glGetString(gfxapi::StringConstant::GL_EXTENSIONS));
}

const char* RendererImpl::vendor() {
    return reinterpret_cast<const char*>(
        gfxapi::glGetString(gfxapi::StringConstant::GL_VENDOR));
}

const char* RendererImpl::version() {
    return reinterpret_cast<const char*>(
        gfxapi::glGetString(gfxapi::StringConstant::GL_VERSION));
}

} // end of anonymous namespace

std::unique_ptr<Renderer> Renderer::create(int width, int height, int depthSize, int stencilSize) {
    return std::unique_ptr<Renderer>(new RendererImpl(width, height, depthSize, stencilSize));
}

}  // end of namespace caze
}  // end of namespace android

#endif // TARGET_OS == CAZE_OS_WINDOWS

