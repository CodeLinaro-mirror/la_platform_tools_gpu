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

#ifndef ANDROID_CAZE_MOCK_CONNECTION_H
#define ANDROID_CAZE_MOCK_CONNECTION_H

#include "Connection.h"

#include <memory>

#include <gmock/gmock.h>

namespace android {
namespace caze {
namespace test {

class MockConnection : public Connection {
public:
    MOCK_METHOD2(send, size_t(const void* data, size_t size));
    MOCK_METHOD2(recv, size_t(void* data, size_t size));
    MOCK_METHOD0(acceptProxy, Connection*());

    const char* error() override { return ""; }

    std::unique_ptr<Connection> accept() override {
        return std::unique_ptr<Connection>(this->acceptProxy());
    }
};

}  // end of namespace test
}  // end of namespace caze
}  // end of namespace android

#endif  // ANDROID_CAZE_MOCK_CONNECTION_H
