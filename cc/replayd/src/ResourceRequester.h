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

#ifndef ANDROID_CAZE_RESOURCE_REQUESTER_H
#define ANDROID_CAZE_RESOURCE_REQUESTER_H

#include "ResourceProvider.h"

#include <utility>
#include <vector>

namespace android {
namespace caze {

class GazerConnection;

// Resource provider which use the GazerConnection to fetch the resources from the server
class ResourceRequester : public ResourceProvider {
public:
    static std::unique_ptr<ResourceRequester> create();

    // Request the resource from the GazerConnection with a GET request
    bool get(const ResourceId& id, const GazerConnection& gazer, void* target,
             size_t size) override;

    // Request all of the requested resources from the GazerConnection with a single GET request
    bool get(const std::vector<std::pair<ResourceId, size_t>>& id, const GazerConnection& gazer,
             void* target) override;

    // No prefetching is supported because there is no storage layer in this resource provider
    bool prefetch(const std::vector<std::pair<ResourceId, size_t>>& resources,
                  const GazerConnection& gazer, void* buffer, size_t size) override;

private:
    ResourceRequester() = default;
};

}  // end of namespace caze
}  // end of namespace android

#endif  // ANDROID_CAZE_RESOURCE_REQUESTER_H
