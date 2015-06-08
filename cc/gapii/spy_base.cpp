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

#include "spy_base.h"

namespace gapii {

void SpyBase::Observation::encode(gapic::Encoder* e) const {
    e->Uint64(reinterpret_cast<uint64_t>(mBase));
    e->Uint64(mSize);
    e->Id(mId);
}

// Inline methods
void SpyBase::init(std::shared_ptr<gapic::Encoder> encoder) {
    mEncoder = encoder;
}

void SpyBase::read(const void* base, uint64_t size) {
    mReads.push_back(observe(base, size));
}

void SpyBase::write(const void* base, uint64_t size) {
    mWrites.push_back(observe(base, size));
}

void SpyBase::encodeObservations() {
    mEncoder->Uint32(mReads.size());
    for (auto r : mReads) {
        r.encode(mEncoder.get());
    }
    mEncoder->Uint32(mWrites.size());
    for (auto r : mWrites) {
        r.encode(mEncoder.get());
    }
    mReads.clear();
    mWrites.clear();
}

SpyBase::Observation SpyBase::observe(const void* base, uint64_t size) {
    gapic::Id id = gapic::Id::Hash(base, size);
    if (mResources.count(id) == 0) {
        mEncoder->Uint16(0xfffd);  // Type ID -- TODO: mEncoder->Id(RESOURCE_ID);
        mEncoder->Id(id);
        mEncoder->Data(base, size);
        mResources.emplace(id);
    }
    return Observation{ base, size, id };
}

}  // namespace gapii

