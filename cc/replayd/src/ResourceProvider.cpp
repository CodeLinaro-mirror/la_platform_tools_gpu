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

#include "Log.h"
#include "ResourceProvider.h"
#include "Target.h"

#include <utility>
#include <vector>

namespace android {
namespace caze {

bool ResourceProvider::get(const std::vector<std::pair<ResourceId, size_t>>& id,
                           const GazerConnection& gazer, void* target) {
    size_t offset = 0;
    for (const auto& it : id) {
        if (!get(it.first, gazer, static_cast<uint8_t*>(target) + offset, it.second)) {
            return false;
        }
        offset += it.second;
    }
    return true;
}

}  // end of namespace caze
}  // end of namespace android
