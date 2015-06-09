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

// Inline methods
void SpyBase::init(std::shared_ptr<gapic::Encoder> encoder) {
    mEncoder = encoder;
}

void SpyBase::read(const void* base, uint64_t size) {
  mObservations.mReads.push_back(observe(base, size));
}

void SpyBase::write(const void* base, uint64_t size) {
  mObservations.mWrites.push_back(observe(base, size));
}

void SpyBase::encodeObservations() {
    mEncoder->Value(&mObservations);
    mObservations.mReads.clear();
    mObservations.mWrites.clear();
}

Observation SpyBase::observe(const void* base, uint64_t size) {
    gapic::Id id = gapic::Id::Hash(base, size);
    if (mResources.count(id) == 0) {
        mEncoder->Uint16(0xfffd);  // Type ID -- TODO: mEncoder->Id(RESOURCE_ID);
        mEncoder->Id(id);
        mEncoder->Data(base, size);
        mResources.emplace(id);
    }
    return Observation(Range(reinterpret_cast<uintptr_t>(base), size), id);
}

}  // namespace gapii
