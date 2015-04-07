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

#include <gapic/target.h>

#if TARGET_OS == GAPID_OS_LINUX

#include "GfxApi.h"
#include "Log.h"
#include "Renderer.h"

#include <cstring>
#include <X11/Xresource.h>

namespace android {
namespace caze {
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
    RendererImpl(int width, int height, int depthSize, int stencilSize);
    virtual ~RendererImpl() override;

    virtual const char* name() override;
    virtual const char* extensions() override;
    virtual const char* vendor() override;
    virtual const char* version() override;

private:
    Display *mDisplay;
    GLXContext mContext;
    GLXPbuffer mPbuffer;
};

RendererImpl::RendererImpl(int width, int height, int depthSize, int stencilSize) {
    mDisplay = XOpenDisplay(nullptr);
    if (mDisplay == nullptr) {
        GAPID_FATAL("Unable to to open X display\n");
    }

    int major;
    int minor;
    if (!glXQueryVersion(mDisplay, &major, &minor) || (major == 1 && minor < 3)) {
        GAPID_FATAL("GLX 1.3+ unsupported by X server (was %d.%d)\n", major, minor);
    }

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
        GAPID_FATAL("Unable to find a suitable X framebuffer config\n");
    }
    GLXFBConfig fbConfig = fbConfigs[0];
    XFree(fbConfigs);

    mContext = glXCreateNewContext(mDisplay, fbConfig, GLX_RGBA_TYPE, nullptr, True);
    if (mContext == nullptr) {
        GAPID_FATAL("Failed to create glX context\n");
    }
    XSync(mDisplay, False);

    const int pbufferAttribs[] = {
        GLX_PBUFFER_WIDTH, width,
        GLX_PBUFFER_HEIGHT, height,
        None
    };
    mPbuffer = glXCreatePbuffer(mDisplay, fbConfig, pbufferAttribs);

    if (!glXMakeContextCurrent(mDisplay, mPbuffer, mPbuffer, mContext)) {
        GAPID_FATAL("Unable to make GLX context current\n");
    }

    gfxapi::Initialize();
}

RendererImpl::~RendererImpl() {
    if (mDisplay != nullptr) {
        glXDestroyContext(mDisplay, mContext);
        glXDestroyPbuffer(mDisplay, mPbuffer);
        XCloseDisplay(mDisplay);
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

#endif // TARGET_OS == GAPID_OS_LINUX

