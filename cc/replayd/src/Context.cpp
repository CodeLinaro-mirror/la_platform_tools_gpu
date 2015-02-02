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

#include "Context.h"
#include "Log.h"
#include "GazerConnection.h"
#include "GlFunctions.h"
#include "Interpreter.h"
#include "MemoryManager.h"
#include "ReplayRequest.h"
#include "ResourceProvider.h"
#include "Stack.h"
#include "Target.h"

#include <cstdlib>
#include <sstream>
#include <string>

namespace android {
namespace caze {
namespace {

void glfwErrorCallback(int error, const char* description) {
    CAZE_WARNING("Glfw error occurred (error code: %d): %s\n", error, description);
}

}  // end of anonymous namespace

std::unique_ptr<Context> Context::create(const GazerConnection& gazer,
                                         ResourceProvider* resourceProvider,
                                         MemoryManager* memoryManager) {
    std::unique_ptr<Context> context(new Context(gazer, resourceProvider, memoryManager));

    if (context->initialize()) {
        return context;
    } else {
        return nullptr;
    }
}

Context::Context(const GazerConnection& gazer, ResourceProvider* resourceProvider,
                 MemoryManager* memoryManager) :
        mGazer(gazer), mResourceProvider(resourceProvider), mMemoryManager(memoryManager)
#if TARGET_OS == CAZE_OS_LINUX || TARGET_OS == CAZE_OS_OSX || TARGET_OS == CAZE_OS_WINDOWS
        , mGlfwWindow(nullptr)
#endif  // TARGET_OS == CAZE_OS_LINUX || TARGET_OS == CAZE_OS_OSX || TARGET_OS == CAZE_OS_WINDOWS
{}

Context::~Context() {
    destroyGl();
}

bool Context::initialize() {
    mReplayRequest = ReplayRequest::create(mGazer, mResourceProvider, mMemoryManager);
    if (mReplayRequest == nullptr) {
        return false;
    }

    if (!mMemoryManager->setVolatileMemory(mReplayRequest->getVolatileMemorySize())) {
        CAZE_WARNING("Setting the volatile memory size failed (size: %u)\n",
                     mReplayRequest->getVolatileMemorySize());
        return false;
    }

    mMemoryManager->setConstantMemory(mReplayRequest->getConstantMemory());

    CAZE_INFO("Prefetching resources...\n");
    mResourceProvider->prefetch(mReplayRequest->getResources(), mGazer,
                                mMemoryManager->getVolatileAddress(),
                                mReplayRequest->getVolatileMemorySize());
    CAZE_INFO("Prefetching ready\n");

    mInMemoryCacheSize = static_cast<uint8_t*>(mMemoryManager->getVolatileAddress()) -
                         static_cast<uint8_t*>(mMemoryManager->getBaseAddress());
    return true;
}

bool Context::interpret() {
    Interpreter interpreter(mMemoryManager, mReplayRequest->getStackSize());
    registerCallbacks(&interpreter);
    return interpreter.run(mReplayRequest->getInstructionList());
}

uint32_t Context::getInMemoryCacheSize() const {
    return mInMemoryCacheSize;
}

void Context::registerCallbacks(Interpreter* interpreter) {
    GlFunctions::Register(interpreter);

    // Custom function for posting and fetching resources to and from the server
    interpreter->registerFunction(Interpreter::POST_FUNCTION_ID,
                                  [this](Stack* stack, bool) { return this->postData(stack); });
    interpreter->registerFunction(Interpreter::RESOURCE_FUNCTION_ID,
                                  [this](Stack* stack, bool) { return this->loadResource(stack); });

    // Function for initializing the Gl context. The first three values from the stack (last three
    // arguments) are not used.
    interpreter->registerFunction(GlFunctions::Init, [this](Stack* stack, bool) {
        stack->discard(3);
        int32_t height = stack->pop<int32_t>();
        int32_t width = stack->pop<int32_t>();

        if (stack->isValid()) {
            CAZE_DEBUG("initGl(%d, %d)\n", height, width);
            return this->initGl(width, height);
        } else {
            CAZE_WARNING("Error during calling function initGl\n");
            return false;
        }
    });
}

bool Context::initGl(int width, int height) {
#ifdef EGL_VERSION_1_0
    EGLint error;

    mEglDisplay = eglGetDisplay(EGL_DEFAULT_DISPLAY);
    error = eglGetError();
    if (error != EGL_SUCCESS) {
        CAZE_WARNING("Failed to get EGL display: %d\n", error);
        return false;
    }

    eglInitialize(mEglDisplay, nullptr, nullptr);
    error = eglGetError();
    if (error != EGL_SUCCESS) {
        CAZE_WARNING("Failed to initialize EGL: %d\n", error);
        return false;
    }

    eglBindAPI(EGL_OPENGL_ES_API);
    error = eglGetError();
    if (error != EGL_SUCCESS) {
        CAZE_WARNING("Failed to bind EGL API: %d\n", error);
        return false;
    }

    // Find a supported EGL context config.
    const int configAttribList[] = {
        // RGBA8 buffer
        EGL_RED_SIZE, 8,
        EGL_GREEN_SIZE, 8,
        EGL_BLUE_SIZE, 8,
        EGL_ALPHA_SIZE, 8,
        EGL_BUFFER_SIZE, 32,

        // D8S8 buffer
        EGL_DEPTH_SIZE, 8,
        EGL_STENCIL_SIZE, 8,

        // GL|ES API version
        // Note: EGL_OPENGL_ES3_BIT_KHR is undefined on Android NDK.
        EGL_RENDERABLE_TYPE, EGL_OPENGL_ES2_BIT,

        EGL_NONE
    };
    int one = 1;
    EGLConfig eglConfig;
    eglChooseConfig(mEglDisplay, configAttribList, &eglConfig, 1, &one);
    error = eglGetError();
    if (error != EGL_SUCCESS) {
        CAZE_WARNING("Failed to choose EGL config: %d\n", error);
        return false;
    }

    // Create an EGL context.
    const int contextAttribList[] = {
        EGL_CONTEXT_CLIENT_VERSION, 2,
        EGL_NONE
    };
    mEglContext = eglCreateContext(mEglDisplay, eglConfig, EGL_NO_CONTEXT, contextAttribList);
    error = eglGetError();
    if (error != EGL_SUCCESS) {
        CAZE_WARNING("Failed to create EGL context: %d\n", error);
        return false;
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
        CAZE_WARNING("Failed to create EGL pbuffer surface: %d\n", error);
        return false;
    }

    eglMakeCurrent(mEglDisplay, mEglSurface, mEglSurface, mEglContext);
    error = eglGetError();
    if (error != EGL_SUCCESS) {
        CAZE_WARNING("Failed to make EGL current: %d\n", error);
        return false;
    }

    mDeviceInfo["egl.vendor"] = eglQueryString(mEglDisplay, EGL_VENDOR);
    mDeviceInfo["egl.version"] = eglQueryString(mEglDisplay, EGL_VERSION);
    mDeviceInfo["egl.extensions"] = eglQueryString(mEglDisplay, EGL_EXTENSIONS);
    mDeviceInfo["egl.client_apis"] = eglQueryString(mEglDisplay, EGL_CLIENT_APIS);
#endif  // EGL_VERSION_1_0

#ifdef GLFW_VERSION_MAJOR
    glfwSetErrorCallback(glfwErrorCallback);

    if (!glfwInit()) {
        CAZE_WARNING("Failed to init glfw\n");
        return false;
    }

    glfwWindowHint(GLFW_VISIBLE, GL_FALSE);
    mGlfwWindow = glfwCreateWindow(width, height, "Replay", nullptr, nullptr);
    if (nullptr == mGlfwWindow) {
        CAZE_WARNING("Failed to create glfw window\n");
        return false;
    }

    glfwMakeContextCurrent(mGlfwWindow);
    glfwSwapInterval(0);  // disable vsync
    mDeviceInfo["glfw.version"] = glfwGetVersionString();
#endif  // GLFW_VERSION_MAJOR

#if GLEW_VERSION
    glewExperimental = GL_TRUE;
    if (glewInit() != GLEW_OK) {
        CAZE_WARNING("Failed to init glew\n");
        return false;
    }

    mDeviceInfo["glew.version"] = reinterpret_cast<char const*>(glewGetString(GLEW_VERSION));
#endif  // GLEW_VERSION

    mDeviceInfo["gl.vendor"] = reinterpret_cast<char const*>(glGetString(GL_VENDOR));
    mDeviceInfo["gl.version"] = reinterpret_cast<char const*>(glGetString(GL_VERSION));
    mDeviceInfo["gl.renderer"] = reinterpret_cast<char const*>(glGetString(GL_RENDERER));
    mDeviceInfo["gl.extensions"] = reinterpret_cast<char const*>(glGetString(GL_EXTENSIONS));

    return true;
}

void Context::destroyGl() {
#ifdef EGL_VERSION_1_0
    EGLint error;

    eglMakeCurrent(mEglDisplay, EGL_NO_SURFACE, EGL_NO_SURFACE, EGL_NO_CONTEXT);
    error = eglGetError();
    if (error != EGL_SUCCESS) {
        CAZE_WARNING("Failed to make EGL current during destroy: %d\n", error);
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
#endif  // EGL_VERSION_1_0

#ifdef GLFW_VERSION_MAJOR
    if (mGlfwWindow) {
        glfwDestroyWindow(mGlfwWindow);
        mGlfwWindow = nullptr;
        glfwTerminate();
    }
#endif  // GLFW_VERSION_MAJOR
}

const DeviceInfo& Context::getDeviceInfo() const {
    return mDeviceInfo;
}

bool Context::loadResource(Stack* stack) {
    uint32_t resourceId = stack->pop<uint32_t>();
    void* address = stack->pop<void*>();

    if (!stack->isValid()) {
        CAZE_WARNING("Error during loadResource\n");
        return false;
    }

    const auto& resourceData = mReplayRequest->getResourceData(resourceId);
    if (!mResourceProvider->get(resourceData.first, mGazer, address, resourceData.second)) {
        CAZE_WARNING("Can't fetch resource: %s\n", resourceData.first.c_str());
        return false;
    }

    return true;
}

bool Context::postData(Stack* stack) {
    const uint32_t count = stack->pop<uint32_t>();
    const void* address = stack->pop<const void*>();

    if (!stack->isValid()) {
        CAZE_WARNING("Error during postData\n");
        return false;
    }

    if (!mGazer.post(address, count)) {
        CAZE_WARNING("Failed to post data to the server\n");
        return false;
    }

    return true;
}

}  // end of namespace caze
}  // end of namespace android
