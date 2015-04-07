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

#ifndef GAPIR_SOCKET_CONNECTION_H
#define GAPIR_SOCKET_CONNECTION_H

#include "connection.h"

#include <stdint.h>

#include <cstddef>
#include <memory>
#include <string>

namespace gapir {

// Connection object using a native socket
class SocketConnection : public Connection {
public:
    ~SocketConnection();

    // Creates a new socket connection listening on the specified hostname and port. Returns a
    // connection object on successful open or a nullptr if opening the connection is unsuccessful
    static std::unique_ptr<Connection> create(const char* hostname, const char* port);

    // Implementation of the Connection interface
    size_t send(const void* data, size_t size) override;
    size_t recv(void* data, size_t size) override;
    const char* error() override;
    std::unique_ptr<Connection> accept() override;

private:
    // Private constructor used only by the static create function
    explicit SocketConnection(int socket);

    // Network initializer class to handle the initialization of the network driver. A socket
    // function can be called only if there is at least one active network initializer is in
    // the system
    struct NetworkInitializer {
        NetworkInitializer();
        ~NetworkInitializer();
    };

    // The underlying socket for the connection
    int mSocket;

    // Network initializer instance lasts for the lifetime of the connection
    NetworkInitializer mNetworkInitializer;
};

}  // namespace gapir

#endif
