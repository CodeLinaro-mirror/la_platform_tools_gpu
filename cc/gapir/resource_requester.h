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

#ifndef GAPIR_RESOURCE_REQUESTER_H
#define GAPIR_RESOURCE_REQUESTER_H

#include "resource_provider.h"

#include <memory> // std::unique_ptr

namespace gapir {

class ServerConnection;

// Resource provider which use the ServerConnection to fetch the resources from the server
class ResourceRequester : public ResourceProvider {
public:
    static std::unique_ptr<ResourceRequester> create();

    // Request the resource from the ServerConnection with a GET request
    bool get(const ResourceId& id, const ServerConnection& gazer,
             void* target, uint32_t size) override;

    // Request all of the requested resources from the ServerConnection with a single GET request
    bool get(const ResourceList& resources, const ServerConnection& gazer,
             void* target, uint32_t size) override;

    // No prefetching is supported because there is no storage layer in this resource provider
    bool prefetch(const ResourceList& resources,
                  const ServerConnection& gazer, void* buffer, uint32_t size) override;

private:
    ResourceRequester() = default;
};

}  // namespace gapir

#endif  // GAPIR_RESOURCE_REQUESTER_H
