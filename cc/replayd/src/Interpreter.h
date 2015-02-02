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

#ifndef ANDROID_CAZE_INTERPRETER_H
#define ANDROID_CAZE_INTERPRETER_H

#include "Stack.h"

#include <stdint.h>

#include <functional>
#include <unordered_map>
#include <utility>

namespace android {
namespace caze {

class MemoryManager;

// Implementation of a (fix sized) stack based virtual machine to interpret the instructions in the
// given opcode stream.
//
// The list of supported opcodes and their detailed definition is described in the
// Interpreter_doc.txt documentation file
class Interpreter {
public:
    // Function ids for implementation specific functions and special debugging functions. These
    // functions shouldn't be called by the opcode stream
    enum FunctionIds : uint16_t {
        // Custom function Ids
        POST_FUNCTION_ID        = 0xff00,
        RESOURCE_FUNCTION_ID    = 0xff01,
        // Debug function Ids
        PRINT_STACK_FUNCTION_ID = 0xff80,
    };

    // Instruction codes for the different instructions. The codes have to be consistent with the
    // codes on the server side.
    enum class InstructionCode : uint8_t {
        CALL        = 0,
        PUSH_I      = 1,
        LOAD_C      = 2,
        LOAD_V      = 3,
        LOAD        = 4,
        POP         = 5,
        STORE_V     = 6,
        STORE       = 7,
        RESOURCE    = 8,
        POST        = 9,
        COPY        = 10,
        CLONE       = 11,
        STRCPY      = 12,
        EXTEND      = 13,
    };

    // General signature for functions callable by the interpreter with a function call instruction.
    // The first argument is a pointer to the stack of the Virtual Machine and the second argument
    // is true if the caller expect the return value of the function to be pushed to the stack. The
    // function should return true if the function call was successful, false otherwise
    typedef std::function<bool(Stack*, bool)> Function;

    // Creates a new interpreter with the specified memory manager (for resolving memory addresses)
    // and with the specified maximum stack size
    Interpreter(const MemoryManager* memoryManager, uint32_t stackDepth);

    // Register a new function for the specific opcode
    void registerFunction(uint16_t opcode, Function function);

    // Runs the interpreter on the instruction list specified by the pointer and by its size.
    bool run(const std::pair<const uint32_t*, uint32_t>& instructions);

private:
    enum : uint32_t {
        TYPE_MASK        = 0x03f00000U,
        FUNCTION_ID_MASK = 0x0000ffffU,
        PUSH_RETURN_MASK = 0x01000000U,
        DATA_MASK20      = 0x000fffffU,
        DATA_MASK26      = 0x03ffffffU,
        TYPE_BIT_SHIFT   = 20,
        OPCODE_BIT_SHIFT = 26,
    };

    // Get type information out from an opcode. The type is always stored in the 7th to 13th MSB
    // (both inclusive) of the opcode
    BaseType extractType(uint32_t opcode) const;

    // Get 20 bit data out from an opcode located in the 20 LSB of the opcode.
    uint32_t extract20bitData(uint32_t opcode) const;

    // Get 26 bit data out from an opcode located in the 26 LSB of the opcode.
    uint32_t extract26bitData(uint32_t opcode) const;

    // Implementation of the opcodes supported by the interpreter. Each function returns true if the
    // operation was successful, false otherwise
    bool call(uint32_t opcode);
    bool pushI(uint32_t opcode);
    bool loadC(uint32_t opcode);
    bool loadV(uint32_t opcode);
    bool load(uint32_t opcode);
    bool pop(uint32_t opcode);
    bool storeV(uint32_t opcode);
    bool store();
    bool resource(uint32_t);
    bool post();
    bool copy(uint32_t opcode);
    bool clone(uint32_t opcode);
    bool strcpy(uint32_t opcode);
    bool extend(uint32_t opcode);

    // Interpret one specific opcode. Returns true if it was successful false otherwise
    bool interpret(uint32_t opcode);

    // Memory manager which managing the memory used during the interpretation
    const MemoryManager* mMemoryManager;

    // The stack of the Virtual Machine
    Stack mStack;

    // Map of the supported function ids to the actual function implementations
    std::unordered_map<uint16_t, Function> mFunctions;
};

}  // end of namespace caze
}  // end of namespace android

#endif  // ANDROID_CAZE_INTERPRETER_H
