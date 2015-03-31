/*
 * Copyright 2015, The Android Open Source Project
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

#include "id.h"
#include "hash.h"

namespace {

void hash(const void* ptr, uint64_t size, gapic::Id& out) {
    int s = static_cast<int>(size);
    MurmurHash3_x86_128(ptr, s, 0x342d23f2, &out.data[0]);
    MurmurHash3_x86_32(ptr, s, 0x4dd61236, &out.data[16]);
}

} // anonymous namespace

namespace gapic {

Id::Id() {}

Id::Id(const void* ptr, uint64_t size) {
    hash(ptr, size, *this);
}

bool Id::operator == (const Id& rhs) const {
    return memcmp(data, rhs.data, sizeof(data)) == 0;
}

} // namespace gapic
