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

#ifndef ANDROID_CAZE_GAZER_CONNECTION_H
#define ANDROID_CAZE_GAZER_CONNECTION_H

#include <memory>
#include <string>
#include <vector>

namespace android {
namespace caze {

class Connection;

// Class for managing the communication between the replay daemon and the server (gazer)
class GazerConnection {
public:
    // Creates a gazer connection using the given connection
    static std::unique_ptr<GazerConnection> create(std::unique_ptr<Connection> conn);

    ~GazerConnection();

    // Returns the resource id of the replay data
    const std::string& replayId() const;

    // Returns the length of the replay data
    uint32_t replayLength() const;

    // Fetch the specified resources to the specified target address from the server. The resources
    // are loaded into the memory address continuously in the order they are specified in the id
    // list. Size have to specify the sum size of the requested resources. The function returns true
    // if fetching of the resources was successful false otherwise.
    bool getResources(const std::vector<std::string>& resourceIds, void* target,
                      uint32_t size) const;

    // Post a blob of data from the given address with the given size to the server. Returns true if
    // the posting was successful false otherwise.
    bool post(const void* postData, uint32_t postSize) const;

    // Type of the message sent to the server. It have to be consistent with the values expected by
    // the server
    enum MessageType : uint8_t {
        MESSAGE_TYPE_GET  = 0,
        MESSAGE_TYPE_POST = 1,
    };

private:
    // Initialize the member variables of the GazerConnection object
    GazerConnection(std::unique_ptr<Connection> conn, const std::string& replayId,
                    uint32_t replayLen);

    // The connection used for sending and receiving data to and from the server.
    std::unique_ptr<Connection> mConn;

    // The length of the replay this connection belongs to.
    uint32_t mReplayLen;

    // The resource id of the replay this request belongs to.
    std::string mReplayId;
};

}  // end of namespace caze
}  // end of namespace android

#endif  // ANDROID_CAZE_GAZER_CONNECTION_H
