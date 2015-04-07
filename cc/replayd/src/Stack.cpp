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
#include "Stack.h"

#define __STDC_FORMAT_MACROS
#include <inttypes.h>

namespace android {
namespace caze {

const char* Stack::Entry::debugInfo(const MemoryManager* memoryManager) const {
    static const size_t size = 256;
    static char buf[size];
    switch (mType) {
    case BaseType::Bool:
        snprintf(buf, size, "bool<%d>", value<bool>());
        break;
    case BaseType::Int8:
        snprintf(buf, size, "int8<%" PRId8 ">", value<int8_t>());
        break;
    case BaseType::Int16:
        snprintf(buf, size, "int16<%" PRId16 ">", value<int16_t>());
        break;
    case BaseType::Int32:
        snprintf(buf, size, "int32<%" PRId32 ">", value<int32_t>());
        break;
    case BaseType::Int64:
        snprintf(buf, size, "int64<%" PRId64 ">", value<int64_t>());
        break;
    case BaseType::Uint8:
        snprintf(buf, size, "uint8<%" PRIu8 ">", value<uint8_t>());
        break;
    case BaseType::Uint16:
        snprintf(buf, size, "uint16<%" PRIu16 ">", value<uint16_t>());
        break;
    case BaseType::Uint32:
        snprintf(buf, size, "uint32<%" PRIu32 ">", value<uint32_t>());
        break;
    case BaseType::Uint64:
        snprintf(buf, size, "uint64<%" PRIu64 ">", value<uint64_t>());
        break;
    case BaseType::Float:
        snprintf(buf, size, "float<%f>", value<float>());
        break;
    case BaseType::Double:
        snprintf(buf, size, "double<%f>", value<double>());
        break;
    case BaseType::AbsolutePointer:
        snprintf(buf, size, "absolute-ptr<%p>", value<void*>());
        break;
    case BaseType::ConstantPointer: {
        uint32_t offset = value<uint32_t>();
        const void* pointer = memoryManager->constantToAbsolute(offset);
        snprintf(buf, size, "constant-ptr<0x%x> (%p)", offset, pointer);
        break;
    }
    case BaseType::VolatilePointer: {
        uint32_t offset = value<uint32_t>();
        const void* pointer = memoryManager->volatileToAbsolute(offset);
        snprintf(buf, size, "volatile-ptr<0x%x> (%p)", offset, pointer);
        break;
    }
    default:
        snprintf(buf, size, "unknown type<%d>", int(mType));
        break;
    }
    return buf;
}

Stack::Stack(uint32_t size, const MemoryManager* memoryManager) :
        mValid(true), mTop(0), mStack(size), mMemoryManager(memoryManager) {
}

void Stack::printStack() const {
    GAPID_DEBUG("Stack size: %u\n", mTop);
    for (uint32_t i = 0; i < mTop; ++i) {
        GAPID_DEBUG("(%d) %s\n", i, mStack[i].debugInfo(mMemoryManager));
    }
}

BaseType Stack::getTopType() {
    if (!mValid) {
        GAPID_WARNING("GetTopType on invalid stack\n");
        return BaseType::Bool;
    }

    if (mTop == 0 || mTop > mStack.size()) {
        mValid = false;
        GAPID_WARNING("GetTopType with invalid stack head: %u (size: %zu)\n", mTop, mStack.size());
        return BaseType::Bool;
    }

    return mStack[mTop - 1].type();
}

void Stack::pushFrom(BaseType type, const void* data) {
    if (!mValid) {
        GAPID_WARNING("PushFrom on invalid stack\n");
        return;
    }

    if (mTop > mStack.size() - 1) {
        mValid = false;
        GAPID_WARNING("PushFrom with invalid stack head: %u (size: %lu)\n", mTop,
                     static_cast<unsigned long>(mStack.size()));
        return;
    }

    mStack[mTop].set(type, data);
    GAPID_DEBUG("+%s pushFrom(%p)\n", mStack[mTop].debugInfo(mMemoryManager), data);
    mTop++;
}

void Stack::popTo(void* address, bool castPtrsToAbsolute) {
    if (!mValid) {
        GAPID_WARNING("PopTo on invalid stack\n");
        return;
    }

    if (mTop == 0 || mTop > mStack.size()) {
        mValid = false;
        GAPID_WARNING("PopTo with invalid stack head: %u (size: %lu)\n", mTop,
                     static_cast<unsigned long>(mStack.size()));
        return;
    }

    mTop--;
    GAPID_DEBUG("-%s popTo(%p)\n", mStack[mTop].debugInfo(mMemoryManager), address);

    if (castPtrsToAbsolute) {
        switch (mStack[mTop].type()) {
            case BaseType::ConstantPointer: {
                uint32_t offset = mStack[mTop].value<uint32_t>();
                void* pointer = const_cast<void*>(mMemoryManager->constantToAbsolute(offset));
                *reinterpret_cast<void**>(address) = pointer;
                return;
            }
            case BaseType::VolatilePointer: {
                uint32_t offset = mStack[mTop].value<uint32_t>();
                void* pointer = mMemoryManager->volatileToAbsolute(offset);
                *reinterpret_cast<void**>(address) = pointer;
                return;
            }
            default:
                break;
        }
    }

    memcpy(address, mStack[mTop].valuePtr(), baseTypeSize(mStack[mTop].type()));
}

void Stack::discard(uint32_t count) {
    if (!mValid) {
        GAPID_WARNING("Discard on invalid stack\n");
        return;
    }

    if (count > mTop) {
        mValid = false;
        GAPID_WARNING("Discarding more element (%u) then in the stack (%u)\n", count, mTop);
        return;
    }

    for (uint32_t i = 0; i < count; i++) {
        GAPID_DEBUG("-%s discard()\n", mStack[mTop - i - 1].debugInfo(mMemoryManager));
    }

    mTop -= count;
}

void Stack::clone(uint32_t n) {
    if (!mValid) {
        GAPID_WARNING("Clone on invalid stack\n");
        return;
    }

    if (mTop >= mStack.size()) {
        mValid = false;
        GAPID_WARNING("Cloning to full stack\n");
        return;
    }

    if (mTop < n + 1) {
        mValid = false;
        GAPID_WARNING("Cloning from invalid index: %u (head: %u)\n", n, mTop);
        return;
    }

    mStack[mTop] = mStack[mTop - n - 1];
    GAPID_DEBUG("+%s clone(%d)\n", mStack[mTop].debugInfo(mMemoryManager), n);
    mTop++;
}

bool Stack::isValid() const {
    return mValid;
}

}  // end of namespace caze
}  // end of namespace android
