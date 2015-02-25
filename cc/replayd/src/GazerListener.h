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

#ifndef ANDROID_CAZE_GAZER_LISTENER_H
#define ANDROID_CAZE_GAZER_LISTENER_H

#include <memory>

namespace android {
namespace caze {

class Connection;
class GazerConnection;

// Class for listening to incoming connections from the server (gazer)
class GazerListener {
public:
    // Construct a GazerListener using the specified connection.
    // maxMemorySize is the maximum memory size that can be reported as
    // supported by this device.
    explicit GazerListener(std::unique_ptr<Connection> conn, uint64_t maxMemorySize);

    // Accept a new incoming connection on the underlying socket and creates a GazerConnection over
    // the newly created socket object.
    std::unique_ptr<GazerConnection> acceptConnection();

private:
    enum ConnectionType {
        DEVICE_INFO    = 0,
        REPLAY_REQUEST = 1,
    };

    // The underlying server socket for the listener
    std::unique_ptr<Connection> mConn;
    // The maximum memory size that can be reported as supported by this device.
    uint64_t mMaxMemorySize;
};

}  // end of namespace caze
}  // end of namespace android

#endif  // ANDROID_CAZE_SOCKET_LISTENER_H
