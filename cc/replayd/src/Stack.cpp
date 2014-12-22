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

Stack::Stack(uint32_t size, const MemoryManager* memoryManager) :
        mValid(true), mTop(0), mStack(size), mMemoryManager(memoryManager) {
}

void Stack::printStack() const {
    CAZE_DEBUG("Stack size: %u\n", mTop);
    for (uint32_t i = 0; i < mTop; ++i) {
        switch (mStack[mTop - i - 1].type()) {
            case BaseType::Bool:
                CAZE_DEBUG("Bool\t%d\n", mStack[mTop - i - 1].value<bool>());
                break;
            case BaseType::Int8:
                CAZE_DEBUG("Int8\t%" PRId8 "\n", mStack[mTop - i - 1].value<int8_t>());
                break;
            case BaseType::Int16:
                CAZE_DEBUG("Int16\t%" PRId16 "\n", mStack[mTop - i - 1].value<int16_t>());
                break;
            case BaseType::Int32:
                CAZE_DEBUG("Int32\t%" PRId32 "\n", mStack[mTop - i - 1].value<int32_t>());
                break;
            case BaseType::Int64:
                CAZE_DEBUG("Int64\t%" PRId64 "\n", mStack[mTop - i - 1].value<int64_t>());
                break;
            case BaseType::Uint8:
                CAZE_DEBUG("Uint8\t%" PRIu8 "\n", mStack[mTop - i - 1].value<uint8_t>());
                break;
            case BaseType::Uint16:
                CAZE_DEBUG("Uint16\t%" PRIu16 "\n", mStack[mTop - i - 1].value<uint16_t>());
                break;
            case BaseType::Uint32:
                CAZE_DEBUG("Uint32\t%" PRIu32 "\n", mStack[mTop - i - 1].value<uint32_t>());
                break;
            case BaseType::Uint64:
                CAZE_DEBUG("Uint64\t%" PRIu64 "\n", mStack[mTop - i - 1].value<uint64_t>());
                break;
            case BaseType::Float:
                CAZE_DEBUG("Float\t%f\n", mStack[mTop - i - 1].value<float>());
                break;
            case BaseType::Double:
                CAZE_DEBUG("Double\t%f\n", mStack[mTop - i - 1].value<double>());
                break;
            case BaseType::AbsolutePointer:
                CAZE_DEBUG("AbsPtr\t%p\n", mStack[mTop - i - 1].value<void*>());
                break;
            case BaseType::ConstantPointer:
                CAZE_DEBUG("ConPtr\t%" PRIu32 "\n", mStack[mTop - i - 1].value<uint32_t>());
                break;
            case BaseType::VolatilePointer:
                CAZE_DEBUG("VolPtr\t%" PRIu32 "\n", mStack[mTop - i - 1].value<uint32_t>());
                break;
        }
    }
}

BaseType Stack::getTopType() {
    if (!mValid) {
        CAZE_WARNING("GetTopType on invalid stack\n");
        return BaseType::Bool;
    }

    if (mTop == 0 || mTop > mStack.size()) {
        mValid = false;
        CAZE_WARNING("GetTopType with invalid stack head: %u (size: %zu)\n", mTop, mStack.size());
        return BaseType::Bool;
    }

    return mStack[mTop - 1].type();
}

void Stack::pushFrom(BaseType type, const void* data) {
    if (!mValid) {
        CAZE_WARNING("PushFrom on invalid stack\n");
        return;
    }

    if (mTop > mStack.size() - 1) {
        mValid = false;
        CAZE_WARNING("PushFrom with invalid stack head: %u (size: %zu)\n", mTop, mStack.size());
        return;
    }

    memset(&mStack[mTop].value<uint32_t>(), 0, sizeof(Entry));
    memcpy(&mStack[mTop].value<uint32_t>(), data, baseTypeSize(type));
    mStack[mTop].type() = type;
    mTop++;
}

void Stack::popTo(void* address) {
    if (!mValid) {
        CAZE_WARNING("PopTo on invalid stack\n");
        return;
    }

    if (mTop == 0 || mTop > mStack.size()) {
        mValid = false;
        CAZE_WARNING("PopTo with invalid stack head: %u (size: %zu)\n", mTop, mStack.size());
        return;
    }

    mTop--;
    memcpy(address, &mStack[mTop].value<uint32_t>(), baseTypeSize(mStack[mTop].type()));
}

void Stack::discard(uint32_t count) {
    if (!mValid) {
        CAZE_WARNING("Discard on invalid stack\n");
        return;
    }

    if (count > mTop) {
        mValid = false;
        CAZE_WARNING("Discarding more element (%u) then in the stack (%u)\n", count, mTop);
        return;
    }

    mTop -= count;
}

void Stack::clone(uint32_t n) {
    if (!mValid) {
        CAZE_WARNING("Clone on invalid stack\n");
        return;
    }

    if (mTop >= mStack.size()) {
        mValid = false;
        CAZE_WARNING("Cloning to full stack\n");
        return;
    }

    if (mTop < n + 1) {
        mValid = false;
        CAZE_WARNING("Cloning from invalid index: %u (head: %u)\n", n, mTop);
        return;
    }

    mStack[mTop] = mStack[mTop - n - 1];
    mTop++;
}

bool Stack::isValid() const {
    return mValid;
}

}  // end of namespace caze
}  // end of namespace android
