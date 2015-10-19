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

#ifndef GAPIC_SCRATCH_ALLOCATOR_H
#define GAPIC_SCRATCH_ALLOCATOR_H

#include "map.h"
#include "vector.h"

namespace gapic {

// ScratchAllocator is a simple linear allocator that uses a single, static heap allocation
// that is pre-allocated at construction.
class ScratchAllocator {
public:
    // Constructs a ScratchAllocator of the specified size in bytes.
    inline ScratchAllocator(size_t size);

    // Destructs the ScratchAllocator, freeing the memory allocated in the constructor.
    inline ~ScratchAllocator();

    // reset sets the head pointer back to 0.
    inline void reset();

    // allocate reserves size bytes from the allocator.
    // The pointer returned will be aligned to the specified number of bytes.
    inline void* allocate(size_t size, size_t alignment);

    // create constructs count instances of T, returning the pointer to the first instance.
    template<typename T>
    inline T* create(size_t count = 1);

    // vector returns a gapic::Vector with the specified maximum capacity.
    template<typename T>
    inline Vector<T> vector(size_t capacity);

    // map returns a gapic::Map with the specified maximum capacity.
    template<typename K, typename V>
    inline Map<K, V> map(size_t capacity);

private:
    ScratchAllocator() = delete;

    uint8_t* mBase; // Base address of the allocator's heap.
    uint8_t* mEnd;  // One byte beyond the end of the allocator's heap.
    uint8_t* mHead; // Next allocation's address.
};


ScratchAllocator::ScratchAllocator(size_t size) {
    mBase = new uint8_t[size];
    mEnd = mBase + size;
    mHead = mBase;
}

ScratchAllocator::~ScratchAllocator() {
    delete[] mBase;
}

inline void ScratchAllocator::reset() {
    mHead = mBase;
}

inline void* ScratchAllocator::allocate(size_t size, size_t alignment) {
    uintptr_t a = reinterpret_cast<uintptr_t>(alignment);
    uintptr_t p = reinterpret_cast<uintptr_t>(mHead);
    if (uintptr_t o = p % a) {
      mHead += a - o; // Alignment
    }
    uint8_t* out = mHead;
    mHead += size;
    if (mHead > mEnd) {
        GAPID_FATAL("ScratchAllocator of size 0x%x is out of memory by 0x%x bytes", mEnd - mBase, mHead - mEnd);
    }
    return out;
}

template<typename T>
inline T* ScratchAllocator::create(size_t count /* = 1 */) {
    void* buffer = allocate(sizeof(T) * count, alignof(T));
    return new(buffer) T();
}

template<typename T>
inline Vector<T> ScratchAllocator::vector(size_t capacity) {
    T* first = create<T>(capacity);
    return Vector<T>(first, 0, capacity);
}

template<typename K, typename V>
inline Map<K, V> ScratchAllocator::map(size_t capacity) {
    typedef typename Map<K, V>::Entry T;
    T* first = create<T>(capacity);
    return Map<K, V>(first, capacity);
}

} // namespace gapic

#endif  // GAPIC_SCRATCH_ALLOCATOR_H
