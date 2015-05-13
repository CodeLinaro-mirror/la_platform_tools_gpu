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

#import <OpenGL/OpenGL.h>
#import <AppKit/AppKit.h>

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
    NSWindow* mWindow;
    NSOpenGLContext* mContext;
};

RendererImpl::RendererImpl()
        : mWidth(0)
        , mHeight(0)
        , mDepthSize(0)
        , mStencilSize(0)
        , mBound(false)
        , mWindow(nullptr)
        , mContext(nullptr) {

    // Initialize with a default target.
    setBackbuffer(8, 8, 24, 8);
}

RendererImpl::~RendererImpl() {
    reset();
}

void RendererImpl::reset() {
    unbind();

    if (mWindow != nullptr) {
        [mWindow close];
        [mWindow release];
        mWindow = nullptr;
    }

    if (mContext != nullptr) {
        [mContext release];
        mContext = nullptr;
    }

    mWidth = 0;
    mHeight = 0;
    mDepthSize = 0;
    mStencilSize = 0;
}

void RendererImpl::setBackbuffer(int width, int height, int depthSize, int stencilSize) {
    if (mContext != nullptr &&
        mWidth == width &&
        mHeight == height &&
        mDepthSize == depthSize &&
        mStencilSize == stencilSize) {

        return;
    }

    const bool wasBound = mBound;

    [NSApplication sharedApplication];

    reset();

    NSRect rect = NSMakeRect(0, 0, width, height);

    mWindow = [[NSWindow alloc]
        initWithContentRect:rect
        styleMask:NSTitledWindowMask
        backing:NSBackingStoreBuffered
        defer:NO
    ];
    if (mWindow == nullptr) {
        GAPID_FATAL("Unable to create NSWindow\n");
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
        GAPID_FATAL("Unable to create NSOpenGLPixelFormat\n");
    }

    mContext = [[NSOpenGLContext alloc] initWithFormat:format shareContext:nil];
    if (mContext == nullptr) {
        GAPID_FATAL("Unable to create NSOpenGLContext\n");
    }

    [mContext setView:[mWindow contentView]];

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
        [mContext makeCurrentContext];
        mBound = true;

        // Initialize the graphics API
        // TODO: Inefficient - consider moving the imports into this renderer
        gfxapi::Initialize();
    }
}

void RendererImpl::unbind() {
    if (mBound) {
        [NSOpenGLContext clearCurrentContext];
        mBound = false;
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

} // anonymous namespace

Renderer* Renderer::create() {
    return new RendererImpl();
}

}  // namespace gapir

