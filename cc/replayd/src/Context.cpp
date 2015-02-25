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
#include "GfxApi.h"
#include "Interpreter.h"
#include "MemoryManager.h"
#include "PostBuffer.h"
#include "Renderer.h"
#include "ReplayRequest.h"
#include "ResourceProvider.h"
#include "Stack.h"
#include "Target.h"

#include <cstdlib>
#include <sstream>
#include <string>

namespace android {
namespace caze {

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

// TODO: Make the PostBuffer size dynamic? It currently holds 2MB of data.
Context::Context(const GazerConnection& gazer, ResourceProvider* resourceProvider,
                 MemoryManager* memoryManager) :
        mGazer(gazer), mResourceProvider(resourceProvider), mMemoryManager(memoryManager),
        mPostBuffer(new PostBuffer(POST_BUFFER_SIZE, [this](const void* address, uint32_t count) {
            return this->mGazer.post(address, count);
        })) {
}

Context::~Context() {
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

    // Function for initializing the API context. The first three argument are not currently used.
    interpreter->registerFunction(gfxapi::Ids::Init, [this](Stack* stack, bool) {
        stack->discard(3);
        int32_t height = stack->pop<int32_t>();
        int32_t width = stack->pop<int32_t>();

        if (stack->isValid()) {
            CAZE_INFO("init(%d, %d)\n", height, width);
            return this->init(width, height);
        } else {
            CAZE_WARNING("Error during calling function initGl\n");
            return false;
        }
    });
}

bool Context::init(int width, int height) {
    mRenderer = Renderer::create(width, height);
    return true;
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

    return mPostBuffer->push(address, count);
}

bool Context::flushPostBuffer(Stack* stack) {
    if (!stack->isValid()) {
        CAZE_WARNING("Error during flushPostBuffer\n");
        return false;
    }

    return mPostBuffer->flush();
}

bool Context::startTimer(Stack* stack) {
    uint8_t index = stack->pop<uint8_t>();
    if (stack->isValid()) {
        if (index < MAX_TIMERS) {
            CAZE_INFO("startTimer(%d)\n", index);
            mTimers[index].Start();
            return true;
        } else {
            CAZE_WARNING("StartTimer called with invalid index %d", index);
        }
    } else {
        CAZE_WARNING("Error while calling function StartTimer\n");
    }
    return false;
}

bool Context::stopTimer(Stack* stack, bool pushReturn) {
    uint8_t index = stack->pop<uint8_t>();
    if (stack->isValid()) {
        if (index < MAX_TIMERS) {
            CAZE_INFO("stopTimer(%d)\n", index);
            uint64_t ns = mTimers[index].Stop();
            if (pushReturn) {
                stack->push(ns);
            }
            return true;
        } else {
            CAZE_WARNING("StopTimer called with invalid index %d", index);
        }
    } else {
        CAZE_WARNING("Error while calling function StopTimer\n");
    }
    return false;
}

}  // end of namespace caze
}  // end of namespace android
