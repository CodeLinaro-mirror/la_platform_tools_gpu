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

#ifndef GAPIR_RESOURCE_DISK_CACHE_H
#define GAPIR_RESOURCE_DISK_CACHE_H

#include "resource_provider.h"

#include <gapic/archive.h>

#include <memory>
#include <string>

namespace gapir {

// Unlimited size disk cache for resources
class ResourceDiskCache : public ResourceProvider {
public:
    // Creates new disk cache with the specified base path. If the base path is not readable or it
    // can't be created then returns the fall back provider.
    static std::unique_ptr<ResourceProvider> create(
            std::unique_ptr<ResourceProvider> fallbackProvider, const std::string& path);

    // If the requested resource is on the disk and its size matches the requested size then load
    // it from the archive. Otherwise request it from the fall back provider and save it to disk.
    bool get(const ResourceId& id, const ServerConnection& gazer, void* target,
             uint32_t size) override;

    // Forwards the get request to the fallback provider without caching its result.
    bool getUncached(const ResourceId& id, const ServerConnection& gazer, void* target,
                     uint32_t size) override;

    // Prefetch the resources because it is significantly faster to request multiple resource form
    // the underlying provider (resource requester) than requesting each resource one by one.
    bool prefetch(const ResourceList& resources,
                  const ServerConnection& gazer, void* buffer, uint32_t size) override;

private:
    ResourceDiskCache(std::unique_ptr<ResourceProvider> fallbackProvider, const std::string& path);

    // Fetch the resources from the underlying resource provider and save them into the disk cache.
    // When fetching multiple resources, write them in-order contiguously to the given buffer and
    // return the total size.
    bool fetch(const ServerConnection& gazer, void* buffer, uint32_t size, const ResourceList& query);

    // Fall back resource provider for the cases when the requested resource is not in the disk cache.
    std::unique_ptr<ResourceProvider> mFallbackProvider;

    // Disk-backed archive holding the cached resources.
    gapic::Archive mArchive;
};

}  // namespace gapir

#endif  // GAPIR_RESOURCE_DISK_CACHE_H
