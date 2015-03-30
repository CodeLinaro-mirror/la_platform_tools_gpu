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

#ifndef ANDROID_CAZE_RESOURCE_DISK_CACHE_H
#define ANDROID_CAZE_RESOURCE_DISK_CACHE_H

#include "ResourceProvider.h"

#include <memory>
#include <string>

namespace android {
namespace caze {

// Unlimited size disk cache for resources
class ResourceDiskCache : public ResourceProvider {
public:
    // Creates new disk cache with the specified base path. If the base path is not readable or it
    // can't be created then returns the fall back provider.
    static std::unique_ptr<ResourceProvider> create(
            std::unique_ptr<ResourceProvider> fallbackProvider, const std::string& path);

    // If the requested resource is on the disk and it's size match with the requested size then
    // load it from disk. Otherwise request it from the fall back provider and save it to disk.
    bool get(const ResourceId& id, const GazerConnection& gazer, void* target,
             uint32_t size) override;

    // Prefetch the resources because it is significantly faster to request multiple resource form
    // the underlying provider (resource requester) than requesting each resource one by one.
    bool prefetch(const ResourceList& resources,
                  const GazerConnection& gazer, void* buffer, uint32_t size) override;

private:
    ResourceDiskCache(std::unique_ptr<ResourceProvider> fallbackProvider, const std::string& path);

    // Fetch the resources from the underlying resource provider and saves them into the disk cache.
    // When fetching multiply resources it calculates the offset of the resources in the response
    // got from the fall back provider to save the correct data for each resource
    bool fetch(const GazerConnection& gazer, void* buffer, uint32_t size, const ResourceList& query);

    // Fall back resource provider for the cases when the requested resource is not in the disk
    // cache
    std::unique_ptr<ResourceProvider> mFallbackProvider;

    // The base path of the disk cache
    const std::string mPath;
};

}  // end of namespace caze
}  // end of namespace android

#endif  // ANDROID_CAZE_RESOURCE_DISK_CACHE_H
