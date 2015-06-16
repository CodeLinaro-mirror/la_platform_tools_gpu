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

#ifndef GAPIR_MEMORY_MANAGER_H
#define GAPIR_MEMORY_MANAGER_H

#include <stdint.h>

#include <memory>
#include <type_traits>
#include <utility>
#include <vector>

namespace gapir {

// Memory manager class for managing the memory used by the replay system. Outside of this class
// there shouldn't be any significant heap allocation in the replay daemon itself (small allocations
// are possible for bookkeeping)
//
// The layout of the memory managed by the memory manager (extra paddings are possible between the
// different memory regions):
// | In memory resource cache | Volatile Memory | Replay data |
class MemoryManager {
public:
    // Creating a memory manager will try to allocate memory based on the size list provided, while
    // keeping at least size * kOverheadFactor free bytes for possible driver overhead allocations.
    // Stopping after the first successful allocation and cause a fatal error if none of the sizes
    // could be allocated.
    explicit MemoryManager(const std::vector<uint32_t>& sizeList);

    // Sets the size of the replay data. Returns true if the given size fits in the memory and false
    // otherwise
    bool setReplayDataSize(uint32_t size);

    // Sets the size of the volatile memory. Returns true if the given size fits in the memory and
    // false otherwise
    bool setVolatileMemory(uint32_t size);

    // Sets the size and the base address of the constant memory. The given memory range have to be
    // inside the range of the volatile memory that is accessible with the getVolatileBaseAddress()
    // and the getVloatileSize() function calls
    void setConstantMemory(const std::pair<const void*, uint32_t>& constantMemory);

    // Returns the size and the base address of the different memory regions managed by the memory
    // manager
    void* getBaseAddress() const { return mMemory.get(); }
    void* getReplayAddress() const { return mReplayData.base; }
    void* getVolatileAddress() const { return mVolatileMemory.base; }
    uint32_t getSize() const { return mSize; }
    uint32_t getConstantSize() const { return mConstantMemory.size; }
    uint32_t getVolatileSize() const { return mVolatileMemory.size; }

    // Converts a given relative (constant or volatile) pointer to an absolute pointer without
    // checking if the given offset is inside the range of that memory
    const void* constantToAbsolute(uint32_t offset) const;
    void* volatileToAbsolute(uint32_t offset) const;

    // Converts an absolute pointer to a relative (constant or volatile) pointer without checking if
    // the address is inside the range
    uint32_t absoluteToConstant(const void* address) const;
    uint32_t absoluteToVolatile(const void* address) const;

    // Checks if the given absolute pointer is points inside the constant or inside the volatile
    // memory
    bool isConstantAddress(const void* address) const;
    bool isVolatileAddress(const void* address) const;

private:
    // Struct to represent a memory interval inside the memory manager with its base address and its
    // size
    struct MemoryRange {
        MemoryRange();
        MemoryRange(uint8_t* base, uint32_t size);
        uint8_t* end() const { return base + size; }

        uint8_t* base;
        uint32_t size;
    };

    // Alignment used for each memory region in bytes
    static const uint32_t kAlignment = std::alignment_of<double>::value;

    // Align a given pointer based on the 'kAlignment' value. The return address is always smaller
    // or equal to the given address what is required by the memory layout used
    uint8_t* align(uint8_t* addr) const;

    // The size and the base address of the memory block managed by the memory manager. This pointer
    // owns the allocated memory
    uint32_t mSize;
    std::unique_ptr<uint8_t[]> mMemory;

    // The size and base address of the replay data. The memory range specified by these values have
    // to specify a subset of the memory managed by the memory manager.
    MemoryRange mReplayData;

    // The size and base address of the constant memory. The constant memory is located inside the
    // replay data, so the memory range specified by these values have to be the subset of the
    // memory range specified by the replay data variables (size and pointer)
    MemoryRange mConstantMemory;

    // The size and the base address of the volatile memory. The memory range specified by these
    // values have to specify a subset of the memory managed by the memory manager.
    MemoryRange mVolatileMemory;
};

}  // namespace gapir

#endif  // GAPIR_MEMORY_MANAGER_H
