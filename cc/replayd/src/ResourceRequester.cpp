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

#include "GazerConnection.h"
#include "ResourceRequester.h"

#include <utility>
#include <vector>

namespace android {
namespace caze {

std::unique_ptr<ResourceRequester> ResourceRequester::create() {
    return std::unique_ptr<ResourceRequester>(new ResourceRequester());
}

bool ResourceRequester::get(const ResourceId& id, const GazerConnection& gazer, void* target,
                            size_t size) {
    return gazer.getResources({id}, target, size);
}

bool ResourceRequester::get(const std::vector<std::pair<ResourceId, size_t>>& id,
                            const GazerConnection& gazer, void* target) {
    size_t querySize = 0;
    std::vector<ResourceId> query;
    for (const auto& r : id) {
        query.push_back(r.first);
        querySize += r.second;
    }
    return gazer.getResources(query, target, querySize);
}

bool ResourceRequester::prefetch(const std::vector<std::pair<ResourceId, size_t>>& resources,
                                 const GazerConnection& gazer, void* buffer, size_t size) {
    return true;
}

}  // end of namespace caze
}  // end of namespace android
