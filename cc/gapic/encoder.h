/*
 * Copyright 2015, The Android Open Source Project
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

#ifndef GAPIC_ENCODER_H
#define GAPIC_ENCODER_H

#include "id.h"

#include <memory>
#include <stdint.h>
#include <unordered_map>

namespace gapic {

class StreamWriter;
class Encodable;
class Encoder;

class Encodable {
public:
    virtual void Encode(Encoder* to) const = 0;
    virtual const gapic::Id& Id() const = 0;
};

// Encoder provides methods for encoding values to the provided StreamWriter
// using variable-length-encoding.
class Encoder {
public:
    Encoder(std::shared_ptr<StreamWriter> output);

    void Bool(bool);
    void Int8(int8_t);
    void Uint8(uint8_t);
    void Uint16(uint16_t);
    void Int16(int16_t);
    void Float32(float);
    void Uint32(uint32_t);
    void Int32(int32_t);
    void Float64(double);
    void Uint64(uint64_t);
    void Int64(int64_t);
    void String(const char*);
    void Data(const void* ptr, int32_t size);
    void Id(const gapic::Id&);

    void Value(const Encodable* obj);
    void Variant(const Encodable* obj);
    void Object(const Encodable* obj);
private:
    std::unordered_map<gapic::Id, uint32_t> mIds;
    std::shared_ptr<StreamWriter> mOutput;
    uint32_t mLastObjectId;
};

} // namespace gapic

#endif // GAPIC_ENCODER_H
