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

#ifndef GAPII_SPY_H
#define GAPII_SPY_H

#include "gles_spy.h"

#include <memory>
#include <unordered_map>

namespace gapii {

class Spy : public GlesSpy {
public:
    Spy();

    EGLBoolean eglInitialize(EGLDisplay dpy, EGLint* major, EGLint* minor);
    EGLBoolean eglMakeCurrent(EGLDisplay display, EGLSurface draw, EGLSurface read,
                              EGLContext context);
    BOOL wglMakeCurrent(HDC hdc, HGLRC hglrc);
    CGLError CGLSetCurrentContext(CGLContextObj ctx);
    void glXMakeContextCurrent(const void* display, GLXDrawable draw, GLXDrawable read,
                               GLXContext ctx);

    inline void RegisterSymbol(const std::string& name, void* symbol) {
        mSymbols.emplace(name, symbol);
    }
    inline void* LookupSymbol(const std::string& name) const {
        const auto symbol = mSymbols.find(name);
        return (symbol == mSymbols.end()) ? nullptr : symbol->second;
    }

private:
    void backbufferInfo(int32_t width, int32_t height,
                        uint32_t color_fmt, uint32_t depth_fmt, uint32_t stencil_fmt,
                        bool resetViewportScissor);

    std::shared_ptr<gapic::Encoder> mEncoder;
    std::unordered_map<std::string, void*> mSymbols;
};

} // namespace gapii

#endif // GAPII_SPY_H
