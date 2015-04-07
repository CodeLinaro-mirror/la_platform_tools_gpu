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

#if TARGET_OS == GAPID_OS_OSX

#include "GfxApi.h"
#include "Log.h"
#include "Renderer.h"

#import <OpenGL/OpenGL.h>
#import <AppKit/AppKit.h>

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
    NSWindow* mWindow;
    NSOpenGLContext* mContext;
};

RendererImpl::RendererImpl(int width, int height, int depthSize, int stencilSize) {
    [NSApplication sharedApplication];

    NSRect rect = NSMakeRect(0, 0, width, height);

    mWindow = [[NSWindow alloc]
        initWithContentRect:rect
        styleMask:NSTitledWindowMask
        backing:NSBackingStoreBuffered
        defer:NO
    ];
    if (mWindow == nullptr) {
        GAPID_FATAL("Unable to create NSWindow");
    }

    NSOpenGLPixelFormatAttribute attributes[] = {
        NSOpenGLPFANoRecovery,
        NSOpenGLPFAColorSize, (NSOpenGLPixelFormatAttribute)32,
        NSOpenGLPFADepthSize, (NSOpenGLPixelFormatAttribute)depthSize,
        NSOpenGLPFAStencilSize, (NSOpenGLPixelFormatAttribute)stencilSize,
        NSOpenGLPFAAccelerated,
        NSOpenGLPFABackingStore,
        (NSOpenGLPixelFormatAttribute)0
    };

    NSOpenGLPixelFormat* format = [[NSOpenGLPixelFormat alloc] initWithAttributes:attributes];
    if (format == nullptr) {
        GAPID_FATAL("Unable to create NSOpenGLPixelFormat");
    }

    mContext = [[NSOpenGLContext alloc] initWithFormat:format shareContext:nil];
    if (mContext == nullptr) {
        GAPID_FATAL("Unable to create NSOpenGLContext");
    }

    [mContext setView:[mWindow contentView]];
    [mContext makeCurrentContext];

    // Initialize the graphics API
    gfxapi::Initialize();
}

RendererImpl::~RendererImpl() {
    [NSOpenGLContext clearCurrentContext];
    [mContext release];
    [mWindow close];
    [mWindow release];
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

#endif // TARGET_OS == GAPID_OS_OSX

