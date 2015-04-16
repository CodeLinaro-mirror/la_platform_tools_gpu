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

#include "encoder.h"
#include "stream_writer.h"

#include <cstring>

namespace gapic {

Encoder::Encoder(StreamWriter* output) : mOutput(output) {}

void Encoder::Bool(bool v) {
    uint8_t b = v ? 1 : 0;
    mOutput->Write(&b, 1);
}

void Encoder::S8(int8_t v) {
    mOutput->Write(&v, 1);
}

void Encoder::U8(uint8_t v) {
    mOutput->Write(&v, 1);
}

void Encoder::U16(uint16_t v) {
    uint8_t buf[9];
    uint16_t space = 0x7f;
    uint8_t tag = 0;
    for (int o = 8; true; o--) {
        if (v <= space) {
            buf[o] = uint8_t(v) | uint8_t(tag);
            mOutput->Write(buf + o, 9 - o);
            return;
        }
        buf[o] = uint8_t(v);
        v = v >> 8;
        space >>= 1;
        tag = (tag >> 1) | 0x80;
    }
}

void Encoder::S16(int16_t v) {
    uint16_t uv = uint16_t(v) << 1;
    U16((v < 0) ? ~uv : uv);
}

void Encoder::F32(float v) {
    uint32_t bits = *reinterpret_cast<uint32_t*>(&v);
    uint32_t shuffled =
            ((bits & 0x000000ff) << 24) |
            ((bits & 0x0000ff00) << 8)  |
            ((bits & 0x00ff0000) >> 8)  |
            ((bits & 0xff000000) >> 24);
    return U32(shuffled);
}

void Encoder::U32(uint32_t v) {
    uint8_t buf[9];
    uint32_t space = 0x7f;
    uint8_t tag = 0;
    for (int o = 8; true; o--) {
        if (v <= space) {
            buf[o] = uint8_t(v) | uint8_t(tag);
            mOutput->Write(buf + o, 9 - o);
            return;
        }
        buf[o] = uint8_t(v);
        v = v >> 8;
        space >>= 1;
        tag = (tag >> 1) | 0x80;
    }
}

void Encoder::S32(int32_t v) {
    uint32_t uv = uint32_t(v) << 1;
    U32((v < 0) ? ~uv : uv);
}

void Encoder::F64(double v) {
    uint64_t bits = *reinterpret_cast<uint64_t*>(&v);
    uint64_t shuffled =
            ((bits & 0x00000000000000ffULL) << 56) |
            ((bits & 0x000000000000ff00ULL) << 40) |
            ((bits & 0x0000000000ff0000ULL) << 24) |
            ((bits & 0x00000000ff000000ULL) << 8)  |
            ((bits & 0x000000ff00000000ULL) >> 8)  |
            ((bits & 0x0000ff0000000000ULL) >> 24) |
            ((bits & 0x00ff000000000000ULL) >> 40) |
            ((bits & 0xff00000000000000ULL) >> 56);
    return U64(shuffled);
}

void Encoder::U64(uint64_t v) {
    uint8_t buf[9];
    uint64_t space = 0x7f;
    uint8_t tag = 0;
    for (int o = 8; true; o--) {
        if (v <= space) {
            buf[o] = uint8_t(v) | uint8_t(tag);
            mOutput->Write(buf + o, 9 - o);
            return;
        }
        buf[o] = uint8_t(v);
        v = v >> 8;
        space >>= 1;
        tag = (tag >> 1) | 0x80;
    }
}

void Encoder::S64(int64_t v) {
    uint64_t uv = uint64_t(v) << 1;
    U64((v < 0) ? ~uv : uv);
}

void Encoder::String(const char* v) {
    uint32_t len = v != nullptr ? static_cast<uint32_t>(strlen(v)) : 0;
    U32(len);
    mOutput->Write(v, len);
}

void Encoder::Data(const void* ptr, int32_t size) {
    U32(size);
    mOutput->Write(ptr, size);
}

void Encoder::Id(const gapic::Id& id) {
    auto it = mIds.find(id);
    if (it != mIds.end()) {
        U32(it->second << 1);
    } else {
        uint32_t sid = mIds.size() + 1;
        mIds[id] = sid;
        U32((sid << 1) | 1);
        mOutput->Write(&id.data, 20);
    }
}

} // namespace gapic
