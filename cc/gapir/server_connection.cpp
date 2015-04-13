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

#include "connection.h"
#include "server_connection.h"

#include <gapic/log.h>

#include <memory>
#include <string>
#include <vector>

namespace gapir {

std::unique_ptr<ServerConnection> ServerConnection::create(std::unique_ptr<Connection> conn) {
    std::string replayId;
    if (!conn->readString(&replayId)) {
        GAPID_WARNING("Failed to read replay id. Error: %s\n", conn->error());
        return nullptr;
    }

    uint32_t replayLen;
    if (conn->recv(&replayLen, sizeof(replayLen)) != sizeof(replayLen)) {
        GAPID_WARNING("Failed to read replay length. Error: %s\n", conn->error());
        return nullptr;
    }

    return std::unique_ptr<ServerConnection>(
            new ServerConnection(std::move(conn), replayId, replayLen));
}

ServerConnection::ServerConnection(std::unique_ptr<Connection> conn, const std::string& replayId,
                                 uint32_t replayLen) :
        mConn(std::move(conn)),
        mReplayLen(replayLen),
        mReplayId(replayId) {
}

ServerConnection::~ServerConnection() {
}

const std::string& ServerConnection::replayId() const {
    return mReplayId;
}

uint32_t ServerConnection::replayLength() const {
    return mReplayLen;
}

bool ServerConnection::getResources(const std::vector<std::string>& resourceIds, void* target,
                          uint32_t size) const {
    GAPID_INFO("GET resources (count: %lu, size: %d, target: %p)\n",
              static_cast<unsigned long>(resourceIds.size()), size, target);

    MessageType type = MESSAGE_TYPE_GET;
    if (mConn->send(&type, sizeof(type)) != sizeof(type)) {
        GAPID_WARNING("Failed to send GET messageType to the server. Error: %s\n", mConn->error());
        return false;
    }

    uint32_t count = static_cast<uint32_t>(resourceIds.size());
    if (mConn->send(&count, sizeof(count)) != sizeof(count)) {
        GAPID_WARNING("Failed to send GET count to the server. Error: %s\n", mConn->error());
        return false;
    }

    for (const auto& res : resourceIds) {
        if (!mConn->sendString(res)) {
            GAPID_WARNING("Failed to send GET resource id to the server. Error: %s\n",
                mConn->error());
            return false;
        }
    }

    size_t received = mConn->recv(target, size);
    if (received != size) {
        GAPID_WARNING("GET resource returned unexpected size. "
            "Expected: 0x%x, Got: 0x%x. Error: %s\n", int(size), int(received), mConn->error());
        return false;
    }

    return true;
}

bool ServerConnection::post(const void* postData, uint32_t postSize) const {
    GAPID_INFO("POST: %p (%d)\n", postData, postSize);

    MessageType type = MESSAGE_TYPE_POST;
    if (mConn->send(&type, sizeof(type)) != sizeof(type)) {
        GAPID_WARNING("Failed to send POST messageType to the server. Error: %s\n", mConn->error());
        return false;
    }

    if (mConn->send(&postSize, sizeof(postSize)) != sizeof(postSize)) {
        GAPID_WARNING("Failed to send POST length to the server. Error: %s\n", mConn->error());
        return false;
    }

    if (mConn->send(postData, postSize) != postSize) {
        GAPID_WARNING("Failed to send POST content to the server. Error: %s\n", mConn->error());
        return false;
    }

    return true;
}

}  // namespace gapir
