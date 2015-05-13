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

#include "mock_connection.h"
#include "server_connection.h"
#include "server_listener.h"
#include "test_utilities.h"

#include <gmock/gmock.h>
#include <gtest/gtest.h>

#include <memory>

using ::testing::_;
using ::testing::DoAll;
using ::testing::IsNull;
using ::testing::NotNull;
using ::testing::Return;
using ::testing::ReturnArg;
using ::testing::StrictMock;
using ::testing::WithArg;

namespace gapir {
namespace test {
namespace {

const uint64_t MAX_MEMORY_SIZE = 1024;

class ServerListenerTest : public ::testing::Test {
protected:
    virtual void SetUp() {
        mConnection = new MockConnection();
        mServerListener.reset(
            new ServerListener(std::unique_ptr<gapic::Connection>(mConnection), MAX_MEMORY_SIZE));
    }

    MockConnection* mConnection;
    std::unique_ptr<ServerListener> mServerListener;
};

}  // anonymous namespace

TEST_F(ServerListenerTest, AcceptConnection) {
    auto clientConnection = new MockConnection();
    mConnection->connections.push(clientConnection);
    pushUint8(&clientConnection->in, ServerListener::REPLAY_REQUEST);
    pushString(&clientConnection->in, "");
    pushUint32(&clientConnection->in, 0);
    EXPECT_THAT(mServerListener->acceptConnection(), NotNull());
}

TEST_F(ServerListenerTest, AcceptConnectionErrorAccept) {
    EXPECT_THAT(mServerListener->acceptConnection(), IsNull());
}

TEST_F(ServerListenerTest, AcceptConnectionErrorServerConnection) {
    std::string replayId = "Replay2";
    auto clientConnection1 = new MockConnection();
    auto clientConnection2 = new MockConnection();
    mConnection->connections.push(clientConnection1);
    mConnection->connections.push(clientConnection2);
    pushUint8(&clientConnection1->in, ServerListener::REPLAY_REQUEST);
    pushUint8(&clientConnection1->in, '1');
    pushUint8(&clientConnection2->in, ServerListener::REPLAY_REQUEST);
    pushString(&clientConnection2->in, replayId);
    pushUint32(&clientConnection2->in, 0);
    EXPECT_THAT(mServerListener->acceptConnection(), NotNull());
    //TODO: check we actually got connection 2
}

}  // namespace test
}  // namespace gapir
