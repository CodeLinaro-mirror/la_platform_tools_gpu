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
#include "GazerConnection.h"
#include "GazerListener.h"
#include "Log.h"

#include <string.h>

#include <memory>
#include <sstream>
#include <string>

namespace android {
namespace caze {

GazerListener::GazerListener(std::unique_ptr<Connection> conn) : mConn(std::move(conn)) {
}

std::unique_ptr<GazerConnection> GazerListener::acceptConnection() {
    while (true) {
        std::unique_ptr<Connection> client = mConn->accept();
        if (client == nullptr) {
            CAZE_WARNING("Failed to accept incoming connection\n");
            return nullptr;
        }

        std::unique_ptr<GazerConnection> conn = GazerConnection::create(std::move(client));
        if (conn != nullptr) {
            return conn;
        } else {
            CAZE_WARNING("Loading GazerConnection failed!\n");
        }
    }
}

}  // end of namespace caze
}  // end of namespace android
