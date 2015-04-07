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

static const uint32_t PROTOCOL_VERSION = 1;

GazerListener::GazerListener(std::unique_ptr<Connection> conn, uint64_t maxMemorySize) :
        mConn(std::move(conn)),
        mMaxMemorySize(maxMemorySize) {
}

std::unique_ptr<GazerConnection> GazerListener::acceptConnection() {
    while (true) {
        std::unique_ptr<Connection> client = mConn->accept();
        if (client == nullptr) {
            GAPID_WARNING("Failed to accept incoming connection\n");
            return nullptr;
        }

        uint8_t connectionType;
        if (client->recv(&connectionType, sizeof(connectionType)) != sizeof(connectionType)) {
            GAPID_WARNING("Failed to read connection type\n");
            return nullptr;
        }

        switch (connectionType) {
            case DEVICE_INFO: {
                GAPID_INFO("Sending device info\n");
                uint8_t ptrSize = sizeof(void*);
                uint8_t ptrAlign = std::alignment_of<void*>::value;
                uint8_t targetOs = TARGET_OS;

                if (!client->send(PROTOCOL_VERSION) ||
                    !client->send(ptrSize) ||
                    !client->send(ptrAlign) ||
                    !client->send(mMaxMemorySize) ||
                    !client->send(targetOs)) {
                    GAPID_WARNING("Failed to send connection header\n");
                    return nullptr;
                }
                break;
            }
            case REPLAY_REQUEST: {
                std::unique_ptr<GazerConnection> conn = GazerConnection::create(std::move(client));
                if (conn != nullptr) {
                    return conn;
                } else {
                    GAPID_WARNING("Loading GazerConnection failed!\n");
                }
                break;
            }
        }
    }
}

}  // end of namespace caze
}  // end of namespace android
