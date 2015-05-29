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

#include "pool.h"

#include <cstdlib>

namespace gapii {

std::shared_ptr<Pool> Pool::create(uint64_t size) {
    return std::shared_ptr<Pool>(new Pool(size));
}

Pool::Pool(uint64_t size) : mData(malloc(size)), mSize(size) {}

Pool::~Pool() {
    free(mData);
}

}  // namespace gapii

