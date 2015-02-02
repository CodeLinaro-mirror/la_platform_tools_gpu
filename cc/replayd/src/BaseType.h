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

#ifndef ANDROID_CAZE_BASE_TYPE_H
#define ANDROID_CAZE_BASE_TYPE_H

#include <stdint.h>

namespace android {
namespace caze {

// Unique ID for each supported data type. The ID have to fit into 6 bits (0-63) to fit into the
// opcode stream and the values have to be consistent with the values on the server side
enum class BaseType : uint8_t {
    Bool            = 0,
    Int8            = 1,
    Int16           = 2,
    Int32           = 3,
    Int64           = 4,
    Uint8           = 5,
    Uint16          = 6,
    Uint32          = 7,
    Uint64          = 8,
    Float           = 9,
    Double          = 10,
    AbsolutePointer = 11,
    ConstantPointer = 12,
    VolatilePointer = 13,
};

// Return the size of the underlying type for the given BaseType
uint32_t baseTypeSize(BaseType type);

// Return the name of the given BaseType
const char* baseTypeName(BaseType type);

// Provide the BaseType value corresponding to the type specified in T. For pointers the
// corresponding base type is AbsolutePointer
template<typename T> struct TypeToBaseType;

template<typename T> struct TypeToBaseType<T*> {
    static const BaseType type = BaseType::AbsolutePointer;
};

template<> struct TypeToBaseType<bool>     { static const BaseType type = BaseType::Bool; };
template<> struct TypeToBaseType<int8_t>   { static const BaseType type = BaseType::Int8; };
template<> struct TypeToBaseType<int16_t>  { static const BaseType type = BaseType::Int16; };
template<> struct TypeToBaseType<int32_t>  { static const BaseType type = BaseType::Int32; };
template<> struct TypeToBaseType<int64_t>  { static const BaseType type = BaseType::Int64; };
template<> struct TypeToBaseType<uint8_t>  { static const BaseType type = BaseType::Uint8; };
template<> struct TypeToBaseType<uint16_t> { static const BaseType type = BaseType::Uint16; };
template<> struct TypeToBaseType<uint32_t> { static const BaseType type = BaseType::Uint32; };
template<> struct TypeToBaseType<uint64_t> { static const BaseType type = BaseType::Uint64; };
template<> struct TypeToBaseType<float>    { static const BaseType type = BaseType::Float; };
template<> struct TypeToBaseType<double>   { static const BaseType type = BaseType::Double; };

}  // end of namespace caze
}  // end of namespace android

#endif  // ANDROID_CAZE_BASE_TYPE_H
