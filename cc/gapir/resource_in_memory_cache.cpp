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

#include "resource_in_memory_cache.h"

#include <gapic/Log.h>

#include <string.h>

#include <algorithm>
#include <string>
#include <utility>
#include <vector>

namespace gapir {

std::unique_ptr<ResourceInMemoryCache> ResourceInMemoryCache::create(
        std::unique_ptr<ResourceProvider> fallbackProvider, void* baseAddress) {
    return std::unique_ptr<ResourceInMemoryCache>(
            new ResourceInMemoryCache(std::move(fallbackProvider), baseAddress));
}

ResourceInMemoryCache::ResourceInMemoryCache(std::unique_ptr<ResourceProvider> fallbackProvider,
                                             void* baseAddress) :
        mFallbackProvider(std::move(fallbackProvider)),
        mBaseAddress(static_cast<uint8_t*>(baseAddress)),
        mSize(0) {
}

ResourceInMemoryCache::Entry::Entry(const ResourceId& id, size_t begin, size_t end) :
        id(id), begin(begin), end(end) {
}

bool ResourceInMemoryCache::get(const ResourceId& id, const ServerConnection& gazer, void* target,
                                uint32_t size) {
    const auto& iter = mCache.find(id);
    if (iter != mCache.end()) {
        // Requested resource is in the cache. Do a memory copy only
        memcpy(target, mBaseAddress + iter->second, size);
        return true;
    } else {
        // Requested resource not in the cache. Fetch it from the fall back provider
        bool fallbackresult = mFallbackProvider->get(id, gazer, target, size);
        if (!fallbackresult) {
            // Fall back provider returned an error.
            return false;
        }

        // Allocate memory for the resource in the cache and save it if the allocation was
        // successful
        uint8_t* ptr = allocateMemory(id, size);
        if (ptr != nullptr) {
            mCache.emplace(id, ptr - mBaseAddress);
            memcpy(ptr, target, size);
        }
        return true;
    }
}

bool ResourceInMemoryCache::prefetch(const ResourceList& resources,
                                     const ServerConnection& gazer, void* buffer, uint32_t size) {
    // Filter out the resources from the prefetch request which are already in cache before
    // forwarding the request to the fall back provider
    ResourceList missingResources;
    std::copy_if(resources.begin(), resources.end(), std::back_inserter(missingResources),
            [this](const std::pair<ResourceId, uint32_t>& r) {
                return this->mCache.count(r.first) == 0;
            });
    return mFallbackProvider->prefetch(missingResources, gazer, buffer, size);
}

void ResourceInMemoryCache::updateSize(size_t newSize) {
    // If the size will be decreased then remove the resources not fitting into the cache anymore
    if (newSize <= mSize) {
        auto it = mBlockList.end();  // Iterator pointing after the last block before each iteration
        while (true) {
            // If there is no more block in the cache then we have nothing left to remove
            if (mBlockList.size() == 0) {
                break;
            }
            --it;  // Move to the next block from the end of the block list

            // If a block (at least partially) lies outside the new memory range then remove it from
            // the cache
            if (it->end > newSize) {
                mCache.erase(it->id);
                it = mBlockList.erase(it);
            } else {
                break;
            }
        }
    }

    mSize = newSize;
}

uint8_t* ResourceInMemoryCache::allocateMemory(const std::string& id, size_t size) {
    // If the requested size is greater then the cache size the the allocation is unsuccessful
    if (size > mSize) {
        return nullptr;
    }

    while (true) {
        // If there is no block in the cache then the new block can be allocated at the base address
        if (mBlockList.empty()) {
            mBlockList.emplace_back(id, 0, size);
            mLastBlock = mBlockList.begin();
            mNextBlock = mBlockList.begin();
            return mBaseAddress;
        }

        // Calculate the begin and the end offset of the currently empty (and active) memory range
        size_t begin, end;
        if (mNextBlock->begin > mLastBlock->end) {
            // | occupied space | last block | free space | next block | occupied space | padding |
            begin = mLastBlock->end;
            end = mNextBlock->begin;
        } else if (mNextBlock->begin > 0) {
            // | free space | next block | occupied space | last block | padding |
            begin = 0;
            end = mNextBlock->begin;
        } else {
            // | last and next block | free space |
            begin = mLastBlock->end;
            end = mSize;
        }

        // If the requested size fits into the current free space then the allocation is successful
        if (end - begin >= size) {
            if (mNextBlock->begin == 0) {
                // Next block is at the beginning of the cache. New blocks have to be inserted after
                // last block what is the last element in the cache
                mLastBlock = mBlockList.emplace(mBlockList.end(), id, begin, begin + size);
            } else {
                mLastBlock = mBlockList.emplace(mNextBlock, id, begin, begin + size);
            }
            return mBaseAddress + mLastBlock->begin;
        } else {
            // The size of the free space isn't enough. Remove the block pointed by the next block
            // iterator and then try again (continue the iteration)
            mCache.erase(mNextBlock->id);
            mNextBlock = mBlockList.erase(mNextBlock);
            if (mNextBlock == mBlockList.end()) {
                mNextBlock = mBlockList.begin();
            }
        }
    }
}

}  // namespace gapir
