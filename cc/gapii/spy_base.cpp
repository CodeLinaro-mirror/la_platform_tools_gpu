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

using gapic::coder::memory::Range;
using gapic::Interval;

namespace {

const size_t SCRATCH_BUFFER_SIZE = 64*1024;

}  // anonymous namespace

namespace gapii {

SpyBase::SpyBase() : mScratch(SCRATCH_BUFFER_SIZE) {}

void SpyBase::init(std::shared_ptr<gapic::Encoder> encoder) {
    mEncoder = encoder;
}

void SpyBase::read(const void* base, uint64_t size) {
    if (size > 0) {
        uintptr_t start = reinterpret_cast<uintptr_t>(base);
        uintptr_t end = start + static_cast<uintptr_t>(size);
        mPendingObservations.merge(Interval<uintptr_t>{start, end});
    }
}

void SpyBase::write(const void* base, uint64_t size) {
    if (size > 0) {
        uintptr_t start = reinterpret_cast<uintptr_t>(base);
        uintptr_t end = start + static_cast<uintptr_t>(size);
        mPendingObservations.merge(Interval<uintptr_t>{start, end});
    }
}

void SpyBase::observe(gapic::Vector<Observation>& observations) {
    observations = mScratch.vector<Observation>(mPendingObservations.count());
    for (auto p : mPendingObservations) {
        gapic::Vector<uint8_t> data(reinterpret_cast<uint8_t*>(p.start), p.end - p.start);
        gapic::Id id = gapic::Id::Hash(data.data(), data.count());
        if (mResources.count(id) == 0) {
            gapic::coder::atom::Resource resource(id, data);
            mEncoder->Variant(&resource);
            mResources.emplace(id);
        }
        observations.append(Observation(Range(p.start, data.count()), id));
    }
    mPendingObservations.clear();
}

}  // namespace gapii
