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

#ifndef ANDROID_CAZE_RENDERER_H
#define ANDROID_CAZE_RENDERER_H

#include <memory>

namespace android {
namespace caze {

// Renderer is an interface to an off-screen rendering context. Constructing a renderer will perform
// any necessary hidden window construction and minimal event handling for the target platform.
// Renderer implementations currently create a OpenGL / GLES rendering context. In the future other
// graphics context may be supported.
class Renderer {
public:
    // Construct and return an offscreen renderer with the specified dimensions.
    static std::unique_ptr<Renderer> create(int width, int height);

    // Destroys the renderer and any associated off-screen windows.
    virtual ~Renderer() {}

    // Returns the name of the renderer's created graphics context.
    virtual const char* name() = 0;

    // Returns the list of extensions that the renderer's graphics context supports.
    virtual const char* extensions() = 0;

    // Returns the name of the vendor that has implemented the renderer's graphics context.
    virtual const char* vendor() = 0;

    // Returns the version of the renderer's graphics context.
    virtual const char* version() = 0;
};

}  // end of namespace caze
}  // end of namespace android

#endif  // ANDROID_CAZE_RENDERER_H
