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

#ifndef GAPIC_WRITER_H
#define GAPIC_WRITER_H

#include <stdint.h>

namespace gapic {

// StreamWriter is a pure-virtual interface used to write data streams.
class StreamWriter {
public:
    virtual void Write(const void* data, uint64_t size) = 0;

protected:
    virtual ~StreamWriter() {}
};

} // namespace gapic

#endif // GAPIC_WRITER_H
