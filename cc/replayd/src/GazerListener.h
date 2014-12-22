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
    // Opens a new socket for listening to incoming connections on the specific host and port.
    explicit GazerListener(std::unique_ptr<Connection> conn);

    // Accept a new incoming connection on the underlying socket and creates a GazerConnection over
    // the newly created socket object.
    std::unique_ptr<GazerConnection> acceptConnection();

private:
    // The underlying server socket for the listener
    std::unique_ptr<Connection> mConn;
};

}  // end of namespace caze
}  // end of namespace android

#endif  // ANDROID_CAZE_SOCKET_LISTENER_H
