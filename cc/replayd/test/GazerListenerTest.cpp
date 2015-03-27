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
#include "MockConnection.h"
#include "TestUtilities.h"

#include <memory>

#include <gmock/gmock.h>
#include <gtest/gtest.h>

using ::testing::_;
using ::testing::DoAll;
using ::testing::IsNull;
using ::testing::NotNull;
using ::testing::Return;
using ::testing::ReturnArg;
using ::testing::StrictMock;
using ::testing::WithArg;

namespace android {
namespace caze {
namespace test {
namespace {

const uint64_t MAX_MEMORY_SIZE = 1024;

class GazerListenerTest : public ::testing::Test {
protected:
    virtual void SetUp() {
        mConnection = new MockConnection();
        mGazerListener.reset(
            new GazerListener(std::unique_ptr<Connection>(mConnection), MAX_MEMORY_SIZE));
    }

    MockConnection* mConnection;
    std::unique_ptr<GazerListener> mGazerListener;
};

}  // end of anonymous namespace

TEST_F(GazerListenerTest, AcceptConnection) {
    auto clientConnection = new MockConnection();
    mConnection->connections.push(clientConnection);
    pushUint8(&clientConnection->in, GazerListener::REPLAY_REQUEST);
    pushString(&clientConnection->in, "");
    pushUint32(&clientConnection->in, 0);
    EXPECT_THAT(mGazerListener->acceptConnection(), NotNull());
}

TEST_F(GazerListenerTest, AcceptConnectionErrorAccept) {
    EXPECT_THAT(mGazerListener->acceptConnection(), IsNull());
}

TEST_F(GazerListenerTest, AcceptConnectionErrorGazerConnection) {
    std::string replayId = "Replay2";
    auto clientConnection1 = new MockConnection();
    auto clientConnection2 = new MockConnection();
    mConnection->connections.push(clientConnection1);
    mConnection->connections.push(clientConnection2);
    pushUint8(&clientConnection1->in, GazerListener::REPLAY_REQUEST);
    pushUint8(&clientConnection1->in, '1');
    pushUint8(&clientConnection2->in, GazerListener::REPLAY_REQUEST);
    pushString(&clientConnection2->in, replayId);
    pushUint32(&clientConnection2->in, 0);
    EXPECT_THAT(mGazerListener->acceptConnection(), NotNull());
    //TODO: check we actually got connection 2
}

}  // end of namespace test
}  // end of namespace caze
}  // end of namespace android
