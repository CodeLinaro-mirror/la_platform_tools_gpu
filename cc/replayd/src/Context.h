/*
 * Copyright 2014, The Android Open Source Project
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

#ifndef ANDROID_CAZE_CONTEXT_H
#define ANDROID_CAZE_CONTEXT_H

#include "GlInclude.h"
#include "Target.h"

#include <memory>
#include <string>
#include <unordered_map>

namespace android {
namespace caze {

class GazerConnection;
class Interpreter;
class MemoryManager;
class ReplayRequest;
class ResourceProvider;
class Stack;

// Type for storing the device info (key: value)
typedef std::unordered_map<std::string, std::string> DeviceInfo;

// Context object for the replay containing the Gl context, the memory manager and the replay
// specific functions to handle network communication of the interpreter
class Context {
public:
    // Creates a new Context object and initialize it with loading the replay request, setting up
    // the memory manager, setting up the caches and prefetching the resources
    static std::unique_ptr<Context> create(const GazerConnection& gazer,
                                           ResourceProvider* resourceProvider,
                                           MemoryManager* memoryManager);

    ~Context();

    // Run the interpreter over the opcode stream of the replay request and returns true if the
    // interpretation was successful false otherwise
    bool interpret();

    // Get the device info belongs to the current device
    const DeviceInfo& getDeviceInfo() const;

    uint32_t getInMemoryCacheSize() const;

private:
    Context(const GazerConnection& gazer, ResourceProvider* resourceProvider,
            MemoryManager* memoryManager);

    // Initialize the context object with loading the replay request, setting up the memory manager,
    // setting up the caches and prefetching the resources
    bool initialize();

    // Register the callbacks for the interpreter (Gl functions, load resource, post resource)
    void registerCallbacks(Interpreter* interpreter);

    // Initialize the Gl context
    bool initGl(int width, int height);

    // Destroy the Gl context
    void destroyGl();

    // Post a chunk of data where the number of bytes is on the top of the stack (uint32_t) and the
    // address for the data is the second element on the stack (void*)
    bool postData(Stack* stack);

    // Load a resource from the resource provider where the index of the resource is at the top of
    // the stack (uint32_t) and the target address is at the second element of the stack (void*)
    bool loadResource(Stack* stack);

    // Gazer connection object to fetch and post resources back to the server
    const GazerConnection& mGazer;

    // Resource provider (possibly with caching) to fetch the resources required by the replay. It
    // is owned by the creator of the Context object.
    ResourceProvider* mResourceProvider;

    // Stores the current size of the in memory resource cache determined by the amount of memory
    // available in the memory manager and the amount of memory required by the replay.
    uint32_t mInMemoryCacheSize;

    // Memory manager to manage the memory used by the replay and by the interpreter. Is is owned by
    // the creator of the Context object.
    MemoryManager* mMemoryManager;

    // The data of the request for this context belongs to
    std::unique_ptr<ReplayRequest> mReplayRequest;

    // The device information object containing the basic info about the current replay device and
    // the details about the Gl version
    DeviceInfo mDeviceInfo;

#ifdef EGL_VERSION_1_0
    // EGL context variables
    EGLContext mEglContext;
    EGLSurface mEglSurface;
    EGLDisplay mEglDisplay;
#endif

#if TARGET_OS == CAZE_OS_LINUX || TARGET_OS == CAZE_OS_OSX || TARGET_OS == CAZE_OS_WINDOWS
    // GLFW window on PCs
    GLFWwindow* mGlfwWindow;
#endif  // TARGET_OS == CAZE_OS_LINUX || TARGET_OS == CAZE_OS_OSX || TARGET_OS == CAZE_OS_WINDOWS
};

}  // end of namespace caze
}  // end of namespace android

#endif  // ANDROID_CAZE_CONTEXT_H
