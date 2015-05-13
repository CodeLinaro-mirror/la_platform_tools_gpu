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

#include "context.h"
#include "gfx_api.h"
#include "interpreter.h"
#include "memory_manager.h"
#include "post_buffer.h"
#include "renderer.h"
#include "replay_request.h"
#include "resource_provider.h"
#include "server_connection.h"
#include "stack.h"

#include <gapic/log.h>
#include <gapic/target.h>

#include <cstdlib>
#include <sstream>
#include <string>

namespace gapir {

std::unique_ptr<Context> Context::create(const ServerConnection& gazer,
                                         ResourceProvider* resourceProvider,
                                         MemoryManager* memoryManager) {
    std::unique_ptr<Context> context(new Context(gazer, resourceProvider, memoryManager));

    if (context->initialize()) {
        return context;
    } else {
        return nullptr;
    }
}

// TODO: Make the PostBuffer size dynamic? It currently holds 2MB of data.
Context::Context(const ServerConnection& gazer, ResourceProvider* resourceProvider,
                 MemoryManager* memoryManager) :
        mServer(gazer), mResourceProvider(resourceProvider), mMemoryManager(memoryManager),
        mBoundRenderer(nullptr),
        mPostBuffer(new PostBuffer(POST_BUFFER_SIZE, [this](const void* address, uint32_t count) {
            return this->mServer.post(address, count);
        })) {
}

Context::~Context() {
    for (auto it = mRenderers.begin(); it != mRenderers.end(); it++) {
        delete it->second;
    }
}

bool Context::initialize() {
    mReplayRequest = ReplayRequest::create(mServer, mResourceProvider, mMemoryManager);
    if (mReplayRequest == nullptr) {
        return false;
    }

    if (!mMemoryManager->setVolatileMemory(mReplayRequest->getVolatileMemorySize())) {
        GAPID_WARNING("Setting the volatile memory size failed (size: %u)\n",
                     mReplayRequest->getVolatileMemorySize());
        return false;
    }

    mMemoryManager->setConstantMemory(mReplayRequest->getConstantMemory());

    GAPID_INFO("Prefetching resources...\n");
    mResourceProvider->prefetch(mReplayRequest->getResources(), mServer,
                                mMemoryManager->getVolatileAddress(),
                                mReplayRequest->getVolatileMemorySize());

    GAPID_INFO("Prefetching ready\n");
    mInMemoryCacheSize = static_cast<uint32_t>(
            static_cast<uint8_t*>(mMemoryManager->getVolatileAddress()) -
            static_cast<uint8_t*>(mMemoryManager->getBaseAddress()));
    return true;
}

bool Context::interpret() {
    Interpreter interpreter(mMemoryManager, mReplayRequest->getStackSize());
    registerCallbacks(&interpreter);
    return interpreter.run(mReplayRequest->getInstructionList()) && mPostBuffer->flush();
}

uint32_t Context::getInMemoryCacheSize() const {
    return mInMemoryCacheSize;
}

void Context::registerCallbacks(Interpreter* interpreter) {
    gfxapi::Register(interpreter);

    // Custom function for posting and fetching resources to and from the server
    interpreter->registerFunction(Interpreter::POST_FUNCTION_ID,
                                  [this](Stack* stack, bool) { return this->postData(stack); });
    interpreter->registerFunction(Interpreter::RESOURCE_FUNCTION_ID,
                                  [this](Stack* stack, bool) { return this->loadResource(stack); });

    // Registering custom synthetic functions
    interpreter->registerFunction(gfxapi::Ids::StartTimer,
                                  [this](Stack* stack, bool) { return this->startTimer(stack); });
    interpreter->registerFunction(gfxapi::Ids::StopTimer,
                                  [this](Stack* stack, bool pushReturn) {
        return this->stopTimer(stack, pushReturn);
    });
    interpreter->registerFunction(gfxapi::Ids::FlushPostBuffer, [this](Stack* stack, bool) {
        return this->flushPostBuffer(stack);
    });

    interpreter->registerFunction(gfxapi::Ids::ReplayCreateRenderer, [this](Stack* stack, bool) {
        int32_t id = stack->pop<uint32_t>();
        if (stack->isValid()) {
            GAPID_INFO("replayCreateRenderer(%d)\n", id);
            if (Renderer* prev = mRenderers[id]) {
                if (mBoundRenderer == prev) {
                    mBoundRenderer = nullptr;
                }
                delete prev;
            }
            mRenderers[id] = Renderer::create();
            return true;
        } else {
            GAPID_WARNING("Error during calling function replayCreateRenderer\n");
            return false;
        }
    });

    interpreter->registerFunction(gfxapi::Ids::ReplayBindRenderer, [this](Stack* stack, bool) {
        int32_t id = stack->pop<uint32_t>();
        if (stack->isValid()) {
            GAPID_INFO("replayBindRenderer(%d)\n", id);
            if (mBoundRenderer != nullptr) {
                mBoundRenderer->unbind();
                mBoundRenderer = nullptr;
            }
            mBoundRenderer = mRenderers[id];
            mBoundRenderer->bind();
            return true;
        } else {
            GAPID_WARNING("Error during calling function replayBindRenderer\n");
            return false;
        }
    });

    interpreter->registerFunction(gfxapi::Ids::BackbufferInfo, [this](Stack* stack, bool) {
        stack->pop<bool>(); // ignored
        uint32_t stencil_fmt = stack->pop<uint32_t>();
        uint32_t depth_fmt = stack->pop<uint32_t>();
        uint32_t color_fmt = stack->pop<uint32_t>();
        int32_t height = stack->pop<int32_t>();
        int32_t width = stack->pop<int32_t>();
        uint32_t depthSize = 0;
        switch (depth_fmt) {
            case static_cast<uint32_t>(gfxapi::RenderbufferFormat::GL_DEPTH_COMPONENT16):
                depthSize = 16;
                break;
            case static_cast<uint32_t>(gfxapi::TexelFormat_GLES_3_0::GL_DEPTH24_STENCIL8):
                depthSize = 24;
                break;
        }
        uint32_t stencilSize = 0;
        switch (stencil_fmt) {
            case static_cast<uint32_t>(gfxapi::RenderbufferFormat::GL_STENCIL_INDEX8):
            case static_cast<uint32_t>(gfxapi::TexelFormat_GLES_3_0::GL_DEPTH24_STENCIL8):
                stencilSize = 8;
        }
         if (stack->isValid()) {
            GAPID_INFO("backbufferInfo(%d, %d, 0x%x, 0x%x, 0x%x)\n",
                    width, height, color_fmt, depth_fmt, stencil_fmt);
            if (mBoundRenderer == nullptr) {
                GAPID_INFO("backbufferInfo called without a bound renderer\n");
                return false;
            }
            mBoundRenderer->setBackbuffer(width, height, depthSize, stencilSize);
            return true;
        } else {
            GAPID_WARNING("Error during calling function replayCreateRenderer\n");
            return false;
        }
    });
}

bool Context::loadResource(Stack* stack) {
    uint32_t resourceId = stack->pop<uint32_t>();
    void* address = stack->pop<void*>();

    if (!stack->isValid()) {
        GAPID_WARNING("Error during loadResource\n");
        return false;
    }

    const auto& resourceData = mReplayRequest->getResourceData(resourceId);
    if (!mResourceProvider->get(resourceData.first, mServer, address, resourceData.second)) {
        GAPID_WARNING("Can't fetch resource: %s\n", resourceData.first.c_str());
        return false;
    }

    return true;
}

bool Context::postData(Stack* stack) {
    const uint32_t count = stack->pop<uint32_t>();
    const void* address = stack->pop<const void*>();

    if (!stack->isValid()) {
        GAPID_WARNING("Error during postData\n");
        return false;
    }

    return mPostBuffer->push(address, count);
}

bool Context::flushPostBuffer(Stack* stack) {
    if (!stack->isValid()) {
        GAPID_WARNING("Error during flushPostBuffer\n");
        return false;
    }

    return mPostBuffer->flush();
}

bool Context::startTimer(Stack* stack) {
    uint8_t index = stack->pop<uint8_t>();
    if (stack->isValid()) {
        if (index < MAX_TIMERS) {
            GAPID_INFO("startTimer(%d)\n", index);
            mTimers[index].Start();
            return true;
        } else {
            GAPID_WARNING("StartTimer called with invalid index %d\n", index);
        }
    } else {
        GAPID_WARNING("Error while calling function StartTimer\n");
    }
    return false;
}

bool Context::stopTimer(Stack* stack, bool pushReturn) {
    uint8_t index = stack->pop<uint8_t>();
    if (stack->isValid()) {
        if (index < MAX_TIMERS) {
            GAPID_INFO("stopTimer(%d)\n", index);
            uint64_t ns = mTimers[index].Stop();
            if (pushReturn) {
                stack->push(ns);
            }
            return true;
        } else {
            GAPID_WARNING("StopTimer called with invalid index %d\n", index);
        }
    } else {
        GAPID_WARNING("Error while calling function StopTimer\n");
    }
    return false;
}

}  // namespace gapir
