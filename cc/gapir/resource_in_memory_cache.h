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

#ifndef GAPIR_RESOURCE_IN_MEMORY_CACHE_H
#define GAPIR_RESOURCE_IN_MEMORY_CACHE_H

#include "resource_provider.h"

#include <list>
#include <memory>
#include <string>
#include <unordered_map>

namespace gapir {

// Fixed size in-memory resource cache with ability to resize itself to an arbitrary size while
// minimizing invalidation of resources in the cache. It uses a ring buffer to store the cache and
// starts invalidating cache entries from the oldest to the newest when more space is required.
//
// The state of the cache is one of the following, using the implementation's iterators:
// | free space                                         (empty cache, iterators are not valid) |
// | last and next block | free space                         (only one block is in the cache) |
// | free space          | next block | occupied space | last block | padding                  |
// | occupied space      | last block | free space     | next block | occupied space | padding |
class ResourceInMemoryCache : public ResourceProvider {
public:
    // Creates a new in-memory cache with the given fallback provider and base address. The initial
    // cache size is 0 byte.
    static std::unique_ptr<ResourceInMemoryCache> create(
            std::unique_ptr<ResourceProvider> fallbackProvider, void* baseAddress);

    // If the requested resource is in the cache, copies it to the given memory address. If it
    // isn't in the cache then requests it from the fallback provider and saves it into the cache
    // if it fits (smaller than or equal to the size of the full cache).
    bool get(const ResourceId& id, const ServerConnection& gazer, void* target,
             uint32_t size) override;

    // No prefetching is done for this provider because of the limited size of the in memory cache
    // and because there is no performance gain if we fetch multiple resources at the same time
    // compared to the case when we fetch them one by one. The prefetch request is forwarded to the
    // fallback provider for possible prefetching after filtering out resources already present in
    // the cache.
    bool prefetch(const ResourceList& resources,
                  const ServerConnection& gazer, void* buffer, uint32_t size) override;

    // Resizes the memory region used by the cache, modifying the end of its memory region.
    // Resources falling outside of the new memory region will be evicted from the cache.
    void updateSize(size_t newSize);

private:
    // Data structure representing a cached resource including its resource id and the offset of
    // its memory location, relative to the base address of the cache. Begin is the offset of the
    // first byte and end is the offset of the byte immediately after the end of the resource. The
    // size of the resource is (end - begin).
    struct Entry {
        Entry(const ResourceId& id, size_t begin, size_t end);

        ResourceId id;
        size_t begin;
        size_t end;
    };

    ResourceInMemoryCache(std::unique_ptr<ResourceProvider> fallbackProvider, void* baseAddress);

    // Allocates the given amount of memory inside the resource cache for the resource with the
    // provided resource id (required for later bookkeeping). If the allocation was successful (the
    // requested memory is smaller then the size of the cache) then returns the base address of the
    // allocated region, otherwise returns a null pointer,
    uint8_t* allocateMemory(const std::string& id, size_t size);

    // Fallback provider for the case when the requested resource is not in the cache.
    std::unique_ptr<ResourceProvider> mFallbackProvider;

    // The base address and the size of the memory used for caching. This memory region is owned by
    // the memory manager class, not by the cache itself.
    uint8_t* mBaseAddress;
    size_t mSize;

    // The list of resources currently in the cache with their location.
    std::list<Entry> mBlockList;

    // Iterator pointing to the last block inserted in the cache.
    std::list<Entry>::iterator mLastBlock;

    // Iterator pointing to the next block to be removed from the cache.
    std::list<Entry>::iterator mNextBlock;

    // The resources currently inside the cache. The value for each key is the offset of the
    // resource inside the cache
    std::unordered_map<ResourceId, size_t> mCache;
};

}  // namespace gapir

#endif  // GAPIR_RESOURCE_IN_MEMORY_CACHE_H
