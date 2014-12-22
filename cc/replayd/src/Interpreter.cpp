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

#include "Interpreter.h"
#include "MemoryManager.h"
#include "Log.h"

#include <utility>
#include <vector>

namespace android {
namespace caze {

Interpreter::Interpreter(const MemoryManager* memoryManager, uint32_t stackDepth) :
        mMemoryManager(memoryManager), mStack(stackDepth, mMemoryManager) {
    registerFunction(PRINT_STACK_FUNCTION_ID, [](Stack* stack, bool) {
        stack->printStack();
        return true;
    });
}

bool Interpreter::run(const std::pair<const uint32_t*, uint32_t>& instructions) {
    for (uint32_t i = 0; i < instructions.second; ++i) {
        if (!interpret(instructions.first[i])) {
            CAZE_WARNING(
                    "Interpreter is stopped because of an interpretation error in opcode %u (%u)\n",
                    i, instructions.first[i]);
            return false;
        }
    }
    return true;
}

void Interpreter::registerFunction(uint16_t functionCode, Function function) {
    mFunctions[functionCode] = function;
}

BaseType Interpreter::extractType(uint32_t opcode) const {
    return BaseType((opcode & TYPE_MASK) >> TYPE_BIT_SHIFT);
}

uint32_t Interpreter::extract20bitData(uint32_t opcode) const {
    return opcode & DATA_MASK20;
}

uint32_t Interpreter::extract26bitData(uint32_t opcode) const {
    return opcode & DATA_MASK26;
}

bool Interpreter::call(uint32_t opcode) {
    auto func = mFunctions.find(opcode & FUNCTION_ID_MASK);
    if (func == mFunctions.end()) {
        CAZE_WARNING("Invalid function id: %u\n", opcode & FUNCTION_ID_MASK);
        return false;
    } else {
        return func->second(&mStack, opcode & PUSH_RETURN_MASK);
    }
}

bool Interpreter::pushI(uint32_t opcode) {
    BaseType type = extractType(opcode);
    uint64_t data = extract20bitData(opcode);
    switch (type) {
        // Sign extension for signed types
        case BaseType::Int32:
        case BaseType::Int64:
            if (data & 0x80000) {
                data |= 0xfffffffffff00000ULL;
            }
            break;
        // Shifting the value into the exponent for floating point types
        case BaseType::Float:
            data <<= 23;
            break;
        case BaseType::Double:
            data <<= 52;
            break;
        default:
            break;
    }
    mStack.pushFrom(type, &data);
    return mStack.isValid();
}

bool Interpreter::loadC(uint32_t opcode) {
    BaseType type = extractType(opcode);
    const void* address = mMemoryManager->constantToAbsolute(extract20bitData(opcode));
    mStack.pushFrom(type, address);
    return mStack.isValid();
}

bool Interpreter::loadV(uint32_t opcode) {
    BaseType type = extractType(opcode);
    const void* address = mMemoryManager->volatileToAbsolute(extract20bitData(opcode));
    mStack.pushFrom(type, address);
    return mStack.isValid();
}

bool Interpreter::load(uint32_t opcode) {
    BaseType type = extractType(opcode);
    const void* address = mStack.pop<const void*>();
    mStack.pushFrom(type, address);
    return mStack.isValid();
}

bool Interpreter::pop(uint32_t opcode) {
    mStack.discard(extract26bitData(opcode));
    return mStack.isValid();
}

bool Interpreter::storeV(uint32_t opcode) {
    void* address = mMemoryManager->volatileToAbsolute(extract26bitData(opcode));
    mStack.popTo(address);
    return mStack.isValid();
}

bool Interpreter::store() {
    void* address = mStack.pop<void*>();
    mStack.popTo(address);
    return mStack.isValid();
}

bool Interpreter::resource(uint32_t opcode) {
    mStack.push<uint32_t>(extract26bitData(opcode));
    return this->call(Interpreter::RESOURCE_FUNCTION_ID);
}

bool Interpreter::post() {
    return this->call(Interpreter::POST_FUNCTION_ID);
}

bool Interpreter::copy(uint32_t opcode) {
    uint32_t count = extract26bitData(opcode);
    const void* source = mStack.pop<const void*>();
    void* target = mStack.pop<void*>();
    memcpy(target, source, count);
    return mStack.isValid();
}

bool Interpreter::clone(uint32_t opcode) {
    mStack.clone(extract26bitData(opcode));
    return mStack.isValid();
}

bool Interpreter::extend(uint32_t opcode) {
    BaseType type = mStack.getTopType();
    uint32_t data = extract26bitData(opcode);
    uint64_t value;
    mStack.popTo(&value);
    switch (type) {
        // Masking out the mantissa end extending it with the new bits for floating point types
        case BaseType::Float: {
            value |= (data & 0x007fffffULL);
            break;
        }
        case BaseType::Double: {
            uint64_t exponent = value & 0xfff0000000000000ULL;
            value <<= 26;
            value |= data;
            value &= 0x000fffffffffffffULL;
            value |= exponent;
            break;
        }
        // Extending the value with 26 new LSB
        default: {
            value = (value << 26) | data;
            break;
        }
    }
    mStack.pushFrom(type, &value);
    return mStack.isValid();
}

bool Interpreter::type() {
    BaseType type = mStack.getTopType();
    mStack.discard(1);
    mStack.push<uint8_t>(static_cast<uint8_t>(type));
    return mStack.isValid();
}

bool Interpreter::strlen() {
    const char* address = mStack.pop<const char*>();
    mStack.push<uint32_t>(::strlen(address));
    return mStack.isValid();
}

bool Interpreter::relAddress() {
    void* address = mStack.pop<void*>();
    if (mMemoryManager->isConstantAddress(address)) {
        uint32_t offset = mMemoryManager->absoluteToConstant(address);
        mStack.pushFrom(BaseType::ConstantPointer, &offset);
    } else if (mMemoryManager->isVolatileAddress(address)) {
        uint32_t offset = mMemoryManager->absoluteToVolatile(address);
        mStack.pushFrom(BaseType::VolatilePointer, &offset);
    } else {
        mStack.pushFrom(BaseType::AbsolutePointer, &address);
    }
    return mStack.isValid();
}

bool Interpreter::absAddress() {
    void* address = mStack.pop<void*>();
    mStack.pushFrom(BaseType::AbsolutePointer, &address);
    return mStack.isValid();
}

#define DEBUG_OPCODE(name, value) CAZE_DEBUG(name ":\t%#010x\n", value)

bool Interpreter::interpret(uint32_t opcode) {
    InstructionCode code = static_cast<InstructionCode>(opcode >> OPCODE_BIT_SHIFT);
    switch (code) {
        case InstructionCode::CALL:
            DEBUG_OPCODE("CALL", opcode);
            return this->call(opcode);
        case InstructionCode::PUSH_I:
            DEBUG_OPCODE("PUSH_I", opcode);
            return this->pushI(opcode);
        case InstructionCode::LOAD_C:
            DEBUG_OPCODE("LOAD_C", opcode);
            return this->loadC(opcode);
        case InstructionCode::LOAD_V:
            DEBUG_OPCODE("LOAD_V", opcode);
            return this->loadV(opcode);
        case InstructionCode::LOAD:
            DEBUG_OPCODE("LOAD", opcode);
            return this->load(opcode);
        case InstructionCode::POP:
            DEBUG_OPCODE("POP", opcode);
            return this->pop(opcode);
        case InstructionCode::STORE_V:
            DEBUG_OPCODE("STORE_V", opcode);
            return this->storeV(opcode);
        case InstructionCode::STORE:
            DEBUG_OPCODE("STORE", opcode);
            return this->store();
        case InstructionCode::RESOURCE:
            DEBUG_OPCODE("RESOURCE", opcode);
            return this->resource(opcode);
        case InstructionCode::POST:
            DEBUG_OPCODE("POST", opcode);
            return this->post();
        case InstructionCode::COPY:
            DEBUG_OPCODE("COPY", opcode);
            return this->copy(opcode);
        case InstructionCode::CLONE:
            DEBUG_OPCODE("CLONE", opcode);
            return this->clone(opcode);
        case InstructionCode::EXTEND:
            DEBUG_OPCODE("EXTEND", opcode);
            return this->extend(opcode);
        case InstructionCode::TYPE:
            DEBUG_OPCODE("TYPE", opcode);
            return this->type();
        case InstructionCode::STRLEN:
            DEBUG_OPCODE("STRLEN", opcode);
            return this->strlen();
        case InstructionCode::REL_ADDRESS:
            DEBUG_OPCODE("REL_ADDRESS", opcode);
            return this->relAddress();
        case InstructionCode::ABS_ADDRESS:
            DEBUG_OPCODE("ABS_ADDRESS", opcode);
            return this->absAddress();
        default:
            CAZE_WARNING("Unknown opcode! %#010x\n", opcode);
            return false;
    }
}

#undef DEBUG_OPCODE

}  // end of namespace caze
}  // end of namespace android
