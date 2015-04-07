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

#ifndef ANDROID_CAZE_STACK_H
#define ANDROID_CAZE_STACK_H

#include "BaseType.h"
#include "Log.h"
#include "MemoryManager.h"

#include <stdint.h>
#include <string.h>

#include <vector>

namespace android {
namespace caze {

// Strongly typed, limited size stack for the stack based virtual machine. If an invalid operation
// is called on the stack then the stack will go into an invalid state where each operation is no op
// (except print stack). From an invalid state the stack can't go back to a valid state again.
class Stack {
public:
    // Construct a new stack with the given size and memory manager
    // The memory manager is needed to resolve constant and volatile pointers to absolute pointers
    Stack(uint32_t size, const MemoryManager* memoryManager);

    // Pop the element from the top of the stack as the given type if its type is matching with the
    // given type. For pointer types convert the pointer to an absolute pointer before returning it.
    // Puts the stack into an invalid state if called on an empty stack or if the requested type
    // doesn't match with the type of the value at the top of the stack.
    template <typename T>
    T pop() {
        if (!mValid) {
            GAPID_WARNING("Pop on invalid stack\n");
            return T();
        }

        if (mTop == 0 || mTop > mStack.size()) {
            mValid = false;
            GAPID_WARNING("Stack head is pointing outside of the stack, offset: %d\n", mTop);
            return T();
        }

        GAPID_DEBUG("-%s pop()\n", mStack[mTop-1].debugInfo(mMemoryManager));
        return PopImpl<T>::pop(this);
    }

    // Push the given value to the top of the stack with a type determined by the type of the value
    // given. Put the stack into invalid state if called on a full stack.
    template <typename T>
    void push(T value) {
        if (!mValid) {
            GAPID_WARNING("Push on invalid stack\n");
            return;
        }

        if (mTop > mStack.size() - 1) {
            mValid = false;
            GAPID_WARNING("Stack head is pointing outside of the stack, offset: %d\n", mTop);
            return;
        }

        mStack[mTop].set(TypeToBaseType<typename std::remove_cv<T>::type>::type, &value);
        GAPID_DEBUG("+%s push()\n", mStack[mTop].debugInfo(mMemoryManager));
        mTop++;
    }

    // Returns the type of the element at the top of the stack. Put the stack into invalid state if
    // called on an empty stack.
    BaseType getTopType();

    // Push a new item to the stack with the given type from the given memory address. Put the stack
    // into invalid state if it is already full before the call.
    void pushFrom(BaseType type, const void* data);

    // Pop the item from the top of the stack to the given memory address. The number of bytes
    // written to the address is determined by the type of the element at the top of the stack.
    // If convertPtrToAbsolute is true and the top of the stack is a pointer, then the pointer is
    // converted to an absolute address before writing to address.
    // The stack will enter an invalid state if called on an empty stack.
    void popTo(void* address, bool convertPtrToAbsolute);

    // Discards count amount of elements from the top of the stack. Put the stack into invalid state
    // if the stack contains fewer element then the given count.
    void discard(uint32_t count);

    // Clone the n-th element from the top of the stack to the top of the stack. The currently top
    // most element is the 0th and the index is increasing with going down in the stack. Put the
    // stack into invalid state if called on a full stacked or if the given index points out from
    // the stack (greater then current size minus one).
    void clone(uint32_t n);

    // Print the content of the stack to the log output. The output is only written to the log
    // output if the debug level is at least DEBUG. This function will work even if the stack is not
    // in a valid state.
    void printStack() const;

    // Returns if the stack is in a valid state or not.
    bool isValid() const;

private:
    // Implementation of the pop method for non pointer types.
    template <typename T>
    struct PopImpl {
        static T pop(Stack* stack) {
            stack->mTop--;
            BaseType baseType = TypeToBaseType<typename std::remove_cv<T>::type>::type;
            if (stack->mStack[stack->mTop].type() != baseType) {
                stack->mValid = false;
                GAPID_WARNING(
                        "Pop type (%s) doesn't match with the type at the top of the stack (%s)\n",
                        baseTypeName(baseType),
                        baseTypeName(stack->mStack[stack->mTop].type()));
                return T();
            }
            return stack->mStack[stack->mTop].value<T>();
        }
    };

    // Implementation of the pop method for pointer types.
    template <typename T>
    struct PopImpl<T*> {
        static T* pop(Stack* stack) {
            stack->mTop--;
            switch (stack->mStack[stack->mTop].type()) {
                case BaseType::AbsolutePointer:
                    return stack->mStack[stack->mTop].value<T*>();
                case BaseType::ConstantPointer: {
                    uint32_t offset = stack->mStack[stack->mTop].value<uint32_t>();
                    const void* address = stack->mMemoryManager->constantToAbsolute(offset);
                    return const_cast<T*>(static_cast<const T*>(address));
                }
                case BaseType::VolatilePointer: {
                    uint32_t offset = stack->mStack[stack->mTop].value<uint32_t>();
                    return static_cast<T*>(stack->mMemoryManager->volatileToAbsolute(offset));
                }
                default:
                    GAPID_WARNING("Pop pointer from non pointer type (%s) is not allowed!\n",
                                 baseTypeName(stack->mStack[stack->mTop].type()));
                    stack->mValid = false;
                    return nullptr;
            }
        }
    };

    class Entry {
    public:
        const void* valuePtr() const {
            return &mValue;
        }

        template <typename T>
        const T& value() const {
            return *reinterpret_cast<const T*>(&mValue);
        }

        const BaseType& type() const {
            return mType;
        }

        void set(BaseType type, const void* data) {
            memcpy(&mValue, data, baseTypeSize(type));
            mType = type;
        }

        // Return a string describing the stack entry.
        // The pointer returned is only valid until the next call to debugInfo, regardless of the
        // Entry instance.
        // This function is not thread safe.
        const char* debugInfo(const MemoryManager* memoryManager) const;
    private:
        // Type of the element stored by this entry
        BaseType mType;

        // Union of all possible types stored on the stack for creating a unified value type with
        // getter function to access the value as a specific type
        union {
            bool     b;
            int8_t   i8;
            int16_t  i16;
            int32_t  i32;
            int64_t  i64;
            uint8_t  u8;
            uint16_t u16;
            uint32_t u32;
            uint64_t u64;
            float    f;
            double   d;
            void*    p;
        } mValue;
    };

    // Indicates if the stack is in a consistent state (true value) or not (false value). The stack
    // go into an inconsistent state after an invalid operation. When mValid is false then all of
    // the operation on the stack (expect printing the stack) produce a warning message and falls
    // back to no op (with zero initialized return value where necessary). The stack can't go back
    // from an invalid state to a valid state again.
    bool mValid;

    // Contains the offset of the first empty slot in the stack from the bottom of the stack. The
    // value of it indicates the number of elements currently in the stack.
    uint32_t mTop;

    // mStack stores the entries currently in the stack. The 0th element corresponds to the bottom
    // of the stack.
    std::vector<Entry> mStack;

    // Reference to the memory manager used to resolve constant and volatile pointers to absolute
    // pointers when thy are popped from the stack.
    const MemoryManager* mMemoryManager;
};

}  // end of namespace caze
}  // end of namespace android

#endif  // ANDROID_CAZE_STACK_H
