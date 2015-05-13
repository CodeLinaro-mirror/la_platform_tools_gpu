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

#include "resource_requester.h"
#include "server_connection.h"

#include <vector>

namespace gapir {

std::unique_ptr<ResourceRequester> ResourceRequester::create() {
    return std::unique_ptr<ResourceRequester>(new ResourceRequester());
}

bool ResourceRequester::get(const ResourceId& id, const ServerConnection& gazer,
                            void* target, uint32_t size) {
    return gazer.getResources({id}, target, size);
}

bool ResourceRequester::get(const ResourceList& resources,
                            const ServerConnection& gazer,
                            void* target, uint32_t size) {
    uint32_t querySize = 0;
    std::vector<ResourceId> query;
    for (const auto& r : resources) {
        query.push_back(r.first);
        querySize += r.second;
    }
    if (querySize > size) {
        // Not enough space, don't even try!
        return false;
    }
    return gazer.getResources(query, target, querySize);
}

bool ResourceRequester::prefetch(const ResourceList& resources,
                                 const ServerConnection& gazer, void* buffer, uint32_t size) {
    return true;
}

}  // namespace gapir
