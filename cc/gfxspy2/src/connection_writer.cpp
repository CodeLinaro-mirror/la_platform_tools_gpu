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

#include "connection_writer.h"

#include <gapic/log.h>
#include <gapic/socket_connection.h>

namespace gapii {

std::shared_ptr<ConnectionWriter> ConnectionWriter::listenSocket(
        const char* hostname, const char* port) {
    auto c = gapic::SocketConnection::createSocket(hostname, port);
    GAPID_INFO("GAPII awaiting connection on socket %s:%s\n", hostname, port);
    return std::shared_ptr<ConnectionWriter>(new ConnectionWriter(c->accept()));
}

std::shared_ptr<ConnectionWriter> ConnectionWriter::listenPipe(
        const char* pipename, bool abstract) {
    auto c = gapic::SocketConnection::createPipe(pipename, abstract);
    GAPID_INFO("GAPII awaiting connection on pipe %s%s\n",
        pipename, (abstract ? " (abstract)" : ""));
    return std::shared_ptr<ConnectionWriter>(new ConnectionWriter(c->accept()));
}

ConnectionWriter::ConnectionWriter(std::unique_ptr<gapic::Connection> connection)
    : mConnection(std::move(connection)) {}

void ConnectionWriter::Write(const void* data, uint64_t size) {
    mConnection->send(data, size);
}

} // namespace gapii
