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

#ifndef GAPII_CONNECTION_HEADER_H
#define GAPII_CONNECTION_HEADER_H

#include <stdint.h>

namespace gapic {

class StreamReader;

} // namespace gapic

namespace gapii {

// ConnectionHeader is the first packet of data sent from the tool controlling
// the capture to the interceptor.
// All fields are encoded little-endian with no compression, regardless of
// architecture.
class ConnectionHeader {
public:
    ConnectionHeader();

    // read reads the ConnectionHeader from the provided stream, returning true
    // on success or false on error.
    bool read(gapic::StreamReader* reader);

    uint8_t  mMagic[4];                     // 's', 'p', 'y', '0'
    uint32_t mVersion;                      // 1
    uint8_t  mObserveFramebufferOnEOF;      // non-zero == enabled
    uint8_t  mObserveFramebufferOnDrawCall; // non-zero == enabled
};

} // namespace gapii

#endif // GAPII_CONNECTION_HEADER_H
