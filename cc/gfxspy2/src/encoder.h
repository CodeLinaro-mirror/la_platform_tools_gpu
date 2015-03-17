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

#ifndef ANDROID_GFXSPY_ENCODER_H
#define ANDROID_GFXSPY_ENCODER_H

#include <stdint.h>

#include "id.h"

class StreamWriter;

class Encoder {
public:
    Encoder(StreamWriter* output);

    void Bool(bool);
    void S8(int8_t);
    void U8(uint8_t);
    void U16(uint16_t);
    void S16(int16_t);
    void F32(float);
    void U32(uint32_t);
    void S32(int32_t);
    void F64(double);
    void U64(uint64_t);
    void S64(int64_t);
    void String(const char*);
    void Data(const void*, uint64_t);
    void Id(const Id&);

private:
    StreamWriter* mOutput;
};

#endif // ANDROID_GFXSPY_ENCODER_H
