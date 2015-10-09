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

#ifndef GAPII_GLES_CONSTANTS_H
#define GAPII_GLES_CONSTANTS_H

#include "gles_imports.h"

#include <vector>
#include <stdint.h>

namespace gapii {

struct ConstantDesc {
    uint32_t name;  // GLenum
    uint32_t type;  // GLenum
    bool scalar;  // False if it is vector constant.
};

// Implementation dependent constants.
extern const ConstantDesc kGLESConstantDescs[];

int GetSizeOfGLType(uint32_t type);

// Obtain the value of given constant as raw bytes.
// Some constants may return vector of values.
// The components are sized based on the GL type.
bool GetConstant(const GlesImports& imports,
                 bool is_gles_30_or_newer,
                 const ConstantDesc* desc,
                 std::vector<uint8_t>* buffer);
}

#endif
