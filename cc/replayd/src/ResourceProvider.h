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

#ifndef ANDROID_CAZE_RESOURCE_PROVIDER_H
#define ANDROID_CAZE_RESOURCE_PROVIDER_H

#include <stdint.h>
#include <string>
#include <utility>
#include <vector>

namespace android {
namespace caze {

class GazerConnection;
typedef std::string ResourceId;

class ResourceProvider {
public:
    typedef std::vector<std::pair<ResourceId, uint32_t>> ResourceList;

    // List of possible arguments can be set on the different resource providers
    enum class Argument {
        IN_MEMORY_CACHE_SIZE,
    };

    virtual ~ResourceProvider() {}

    // Loads resources from one of the available resource providers with the given <resourceId>.
    // Resources are written to the memory location given by <target> and at most <size> bytes
    // are written. Returns true if the resource was successfully loaded and all <size> bytes were
    // written, otherwise false.
    virtual bool get(const ResourceId& id, const GazerConnection& gazer,
                     void* target, uint32_t size) = 0;
    virtual bool get(const ResourceList& resources, const GazerConnection& gazer,
                     void* target, uint32_t size);

    // Prefetches the resources for resource providers where prefetching is available.
    // The resources vector have to contain (resource id, resource size) pairs and buffer should
    // point to a buffer at least as big as the largest resource. The buffer will be used for
    // temporary objects only with no guarantee about its final content. Returns true if prefetching
    // was successful, otherwise false.
    virtual bool prefetch(const ResourceList& resources,
                          const GazerConnection& gazer, void* buffer, uint32_t size) = 0;
};

}  // end of namespace caze
}  // end of namespace android

#endif  // ANDROID_CAZE_RESOURCE_PROVIDER_H
