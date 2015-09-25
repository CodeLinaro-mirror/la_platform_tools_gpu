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

#include "connection_header.h"

#include <gapic/log.h>
#include <gapic/stream_reader.h>

namespace gapii {

ConnectionHeader::ConnectionHeader()
    : mVersion(0)
    , mObserveFramebufferOnEOF(0)
    , mObserveFramebufferOnDrawCall(0) {}

bool ConnectionHeader::read(gapic::StreamReader* reader) {
    if (!reader->read(mMagic)) {
        return false;
    }
    if (mMagic[0] != 's'
     || mMagic[1] != 'p'
     || mMagic[2] != 'y'
     || mMagic[3] != '0') {
        GAPID_WARNING("ConnectionHeader magic was not as expected. Got %c%c%c%c",
            mMagic[0], mMagic[1], mMagic[2], mMagic[3]);
        return false;
     }

    if (!reader->read(mVersion)) {
        return false;
    }
    // TODO: Endian-swap version if GAPII is running on a big-endian architecture.
    if (mVersion != 1) {
        GAPID_WARNING("Unsupported ConnectionHeader version. Got %d. Only understand 1.",
            mVersion);
        return false;
    }

    return reader->read(mObserveFramebufferOnEOF) &&
           reader->read(mObserveFramebufferOnDrawCall);
}

} // namespace gapii
