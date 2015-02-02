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

#include "ReplayRequest.h"
#include "GazerConnection.h"
#include "Log.h"
#include "MemoryManager.h"
#include "ResourceProvider.h"

#include <string.h>

#include <string>
#include <utility>
#include <vector>

namespace android {
namespace caze {

std::unique_ptr<ReplayRequest> ReplayRequest::create(const GazerConnection& gazer,
                                                     ResourceProvider* resourceProvider,
                                                     MemoryManager* memoryManager) {
    CAZE_INFO("Load replay request...\n");

    memoryManager->setReplayDataSize(gazer.replayLength());

    // Request the replay data from the gazer connection
    if (!resourceProvider->get(gazer.replayId(), gazer, memoryManager->getReplayAddress(),
                               gazer.replayLength())) {
        CAZE_WARNING("Can't load replay request: %s\n", gazer.replayId().c_str());
        return nullptr;
    }

    std::unique_ptr<ReplayRequest> req(new ReplayRequest());
    if (req->load(memoryManager->getReplayAddress(), gazer.replayLength())) {
        return req;
    } else {
        return nullptr;
    }
}

uint32_t ReplayRequest::getStackSize() const {
    return mStackSize;
}

uint32_t ReplayRequest::getVolatileMemorySize() const {
    return mVolatileMemorySize;
}

const std::vector<std::pair<std::string, uint32_t>>& ReplayRequest::getResources() const {
    return mResources;
}

const std::pair<const void*, uint32_t>& ReplayRequest::getConstantMemory() const {
    return mConstantMemory;
}

const std::pair<std::string, uint32_t>& ReplayRequest::getResourceData(uint32_t resourceId) const {
    return mResources[resourceId];
}

const std::pair<const uint32_t*, uint32_t>& ReplayRequest::getInstructionList() const {
    return mInstructionList;
}

bool ReplayRequest::load(void* data, uint32_t size) {
    // Parse the data fetched from the gazer connection
    const uint8_t* ptr = static_cast<uint8_t*>(data);
    ptr = loadStackSize(ptr);
    ptr = loadVolatileMemorySize(ptr);
    ptr = loadConstantMemory(ptr);
    ptr = loadResourceIds(ptr);
    ptr = loadInstructionList(ptr);

    CAZE_INFO("Replay request loaded\n");

    return ptr - size == data;
}

const uint8_t* ReplayRequest::loadVolatileMemorySize(const uint8_t* ptr) {
    mVolatileMemorySize = *reinterpret_cast<const uint32_t*>(ptr);
    ptr += sizeof(uint32_t);
    CAZE_INFO("Volatile memory size: %d\n", mVolatileMemorySize);
    return ptr;
}

const uint8_t* ReplayRequest::loadStackSize(const uint8_t* ptr) {
    mStackSize = *reinterpret_cast<const uint32_t*>(ptr);
    ptr += sizeof(uint32_t);
    CAZE_INFO("Stack size: %d\n", mStackSize);
    return ptr;
}

const uint8_t* ReplayRequest::loadConstantMemory(const uint8_t* ptr) {
    uint32_t constantMemorySize = *reinterpret_cast<const uint32_t*>(ptr);
    ptr += sizeof(uint32_t);

    mConstantMemory = {ptr, constantMemorySize};
    CAZE_INFO("Constant memory size: %d\n", constantMemorySize);
    ptr += constantMemorySize;

    return ptr;
}

const uint8_t* ReplayRequest::loadResourceIds(const uint8_t* ptr) {
    uint32_t resourceCount = *reinterpret_cast<const uint32_t*>(ptr);
    ptr += sizeof(uint32_t);

    mResources.reserve(resourceCount);
    for (uint32_t i = 0; i < resourceCount; ++i) {
        uint32_t resourceNameLength = *reinterpret_cast<const uint32_t*>(ptr);
        ptr += sizeof(uint32_t);

        const std::string resourceName(reinterpret_cast<const char*>(ptr), resourceNameLength);
        ptr += resourceNameLength;

        uint32_t resourceSize = *reinterpret_cast<const uint32_t*>(ptr);
        ptr += sizeof(uint32_t);

        mResources.emplace_back(std::move(resourceName), resourceSize);
    }
    CAZE_INFO("Resources: %d\n", resourceCount);

    return ptr;
}

const uint8_t* ReplayRequest::loadInstructionList(const uint8_t* ptr) {
    uint32_t instructionListSize = *reinterpret_cast<const uint32_t*>(ptr);
    ptr += sizeof(uint32_t);

    const uint32_t instructionCount = instructionListSize / sizeof(uint32_t);
    mInstructionList = {reinterpret_cast<const uint32_t*>(ptr), instructionCount};
    CAZE_INFO("Instruction count: %d\n", instructionCount);
    ptr += instructionListSize;

    return ptr;
}

}  // end of namespace caze
}  // end of namespace android
