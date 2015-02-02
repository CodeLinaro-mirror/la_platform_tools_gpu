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

#include "Log.h"
#include "MemoryManager.h"

#include <utility>
#include <vector>

namespace android {
namespace caze {

MemoryManager::MemoryRange::MemoryRange() : base(nullptr), size(0) {
}

MemoryManager::MemoryRange::MemoryRange(uint8_t* base, uint32_t size) : base(base), size(size) {
}

MemoryManager::MemoryManager(const std::vector<uint32_t>& sizeList) : mConstantMemory(nullptr, 0) {
    for (auto size : sizeList) {
        mSize = size;
        mMemory.reset(new uint8_t[mSize]);
        if (mMemory) {
            break;
        }
        CAZE_INFO("Failed to allocate %d of volatile memory, continuing...\n", size);
    }

    if (!mMemory) {
        CAZE_FATAL("Couldn't allocate any volatile memory size.\n");
    }

    CAZE_INFO("Base address: %p\n", mMemory.get());
    setReplayDataSize(0);
    setVolatileMemory(mSize);
}

bool MemoryManager::setReplayDataSize(uint32_t size) {
    if (size > mSize) {
        return false;
    }

    mReplayData = {align(mMemory.get() + mSize - size), size};
    return true;
}

bool MemoryManager::setVolatileMemory(uint32_t size) {
    if (size > mSize - mReplayData.size) {
        return false;
    }

    mVolatileMemory = {align(mReplayData.base - size), size};
    return true;
}

void MemoryManager::setConstantMemory(const std::pair<const void*, uint32_t>& constantMemory) {
    mConstantMemory.base = const_cast<uint8_t*>(static_cast<const uint8_t*>(constantMemory.first));
    mConstantMemory.size = constantMemory.second;
}

const void* MemoryManager::constantToAbsolute(uint32_t offset) const {
    return mConstantMemory.base + offset;
}

void* MemoryManager::volatileToAbsolute(uint32_t offset) const {
    return mVolatileMemory.base + offset;
}

uint32_t MemoryManager::absoluteToConstant(const void* address) const {
    return static_cast<const uint8_t*>(address) - mConstantMemory.base;
}

uint32_t MemoryManager::absoluteToVolatile(const void* address) const {
    return static_cast<const uint8_t*>(address) - mVolatileMemory.base;
}

bool MemoryManager::isConstantAddress(const void* address) const {
    return address >= mConstantMemory.base && address < mConstantMemory.end();
}

bool MemoryManager::isVolatileAddress(const void* address) const {
    return address >= mVolatileMemory.base && address < mVolatileMemory.end();
}

uint8_t* MemoryManager::align(uint8_t* addr) const {
    size_t x = reinterpret_cast<size_t>(addr);
    x -= x % MemoryManager::kAlignment;
    return reinterpret_cast<uint8_t*>(x);
}

}  // end of namespace caze
}  // end of namespace android
