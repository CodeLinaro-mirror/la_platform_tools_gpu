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

#ifndef GAPII_CONNECTION_WRITER_H
#define GAPII_CONNECTION_WRITER_H

#include <gapic/stream_writer.h>

#include <memory>

namespace gapic {

class Connection;

} // namespace gapic


namespace gapii {

// ConnectionWriter is an implementation of the StreamWriter interface that writes to
// an incoming TCP connection.
class ConnectionWriter : public gapic::StreamWriter {
public:
    // listenSocket blocks and waits for a TCP connection to be made on the specified host and
    // port, returning a ConnectionWriter once a connection is established.
    static std::shared_ptr<ConnectionWriter> listenSocket(const char* hostname, const char* port);

    // listenPipe blocks and waits for a UNIX connection to be made on the specified pipe name,
    // optionally abstract, returning a ConnectionWriter once a connection is established.
    static std::shared_ptr<ConnectionWriter> listenPipe(const char* pipename, bool abstract);

    // gapic::StreamWriter compliance
    virtual void Write(const void* data, uint64_t size) override;

private:
    ConnectionWriter(std::unique_ptr<gapic::Connection>);

    std::unique_ptr<gapic::Connection> mConnection;
};

} // namespace gapii

#endif  // GAPII_CONNECTION_WRITER_H

