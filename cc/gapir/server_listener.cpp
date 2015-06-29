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

#include "renderer.h"
#include "server_connection.h"
#include "server_listener.h"

#include <gapic/connection.h>
#include <gapic/log.h>

#include <string.h>

#include <memory>
#include <sstream>
#include <string>

namespace gapir {

const uint32_t PROTOCOL_VERSION = 1;
const int MAX_CONSECUTIVE_ACCEPT_FAILURES = 10;

ServerListener::ServerListener(std::unique_ptr<gapic::Connection> conn, uint64_t maxMemorySize) :
        mConn(std::move(conn)),
        mMaxMemorySize(maxMemorySize) {
}

std::unique_ptr<ServerConnection> ServerListener::acceptConnection() {
    int consecutiveAcceptFailures = 0;
    while (true) {
        std::unique_ptr<gapic::Connection> client = mConn->accept();
        if (client == nullptr) {
            GAPID_WARNING("Failed to accept incoming connection\n");
            consecutiveAcceptFailures++;
            if (consecutiveAcceptFailures >= MAX_CONSECUTIVE_ACCEPT_FAILURES) {
                GAPID_WARNING("Multiple consecutive failures (%d) accepting connection", consecutiveAcceptFailures);
                return nullptr;
            }
            continue;
        }
        consecutiveAcceptFailures = 0;

        uint8_t connectionType;
        if (client->recv(&connectionType, sizeof(connectionType)) != sizeof(connectionType)) {
            GAPID_WARNING("Failed to read connection type\n");
            continue;
        }

        switch (connectionType) {
            case DEVICE_INFO: {
                GAPID_INFO("Sending device info\n");
                uint8_t ptrSize = sizeof(void*);
                uint8_t ptrAlign = std::alignment_of<void*>::value;
                uint8_t targetOs = TARGET_OS;
                std::unique_ptr<Renderer> r(Renderer::create());
                r->bind();

                if (!client->send(PROTOCOL_VERSION) ||
                    !client->send(ptrSize) ||
                    !client->send(ptrAlign) ||
                    !client->send(mMaxMemorySize) ||
                    !client->send(targetOs) ||
                    !client->sendString(r->extensions()) ||
                    !client->sendString(r->name()) ||
                    !client->sendString(r->vendor()) ||
                    !client->sendString(r->version())) {
                    GAPID_WARNING("Failed to send connection header\n");
                }
                break;
            }
            case REPLAY_REQUEST: {
                std::unique_ptr<ServerConnection> conn = ServerConnection::create(std::move(client));
                if (conn != nullptr) {
                    return conn;
                } else {
                    GAPID_WARNING("Loading ServerConnection failed!\n");
                }
                break;
            }
            case SHUTDOWN_REQUEST: {
                GAPID_INFO("Shutdown request received!\n")
                return nullptr;
            }
            default: {
                GAPID_WARNING("Unknown connection type %d ignored\n", connectionType);
            }
        }
    }
}

}  // namespace gapir
