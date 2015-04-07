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

#ifndef GAPIR_MOCK_RESOURCE_PROVIDER_H
#define GAPIR_MOCK_RESOURCE_PROVIDER_H

#include "resource_provider.h"
#include "server_connection.h"

#include <vector>

#include <gmock/gmock.h>

namespace gapir {
namespace test {

class MockResourceProvider : public ResourceProvider {
public:
    MOCK_METHOD4(get, bool(const ResourceId& id, const ServerConnection& gazer, void* target,
                           uint32_t size));

    MOCK_METHOD3(get, bool(const ResourceList& resources,
                           const ServerConnection& gazer, void* target));

    MOCK_METHOD4(prefetch, bool(const ResourceList& resources,
                                const ServerConnection& gazer, void* buffer, uint32_t size));
};

}  // namespace test
}  // namespace gapir

#endif  // GAPIR_MOCK_RESOURCE_PROVIDER_H
