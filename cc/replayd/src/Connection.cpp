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

#include "Connection.h"

#include <string>

namespace android {
namespace caze {

bool Connection::sendString(const std::string& s) {
    uint32_t length = static_cast<uint32_t>(s.size());
    if (this->send(&length, sizeof(length)) != sizeof(length)) {
        return false;
    }
    return this->send(s.c_str(), length) == length;
}

bool Connection::readString(std::string* s) {
    uint32_t length;
    if (this->recv(&length, sizeof(length)) != sizeof(length)) {
        return false;
    }

    s->resize(length);
    return this->recv(&s->front(), length) == length;
}

}  // end of namespace caze
}  // end of namespace android
