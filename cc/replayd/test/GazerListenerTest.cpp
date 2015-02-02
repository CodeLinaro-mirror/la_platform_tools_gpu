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

const size_t BUFFER_SIZE = 4096;

class GazerListenerTest : public ::testing::Test {
protected:
    virtual void SetUp() {
        mConnection = new StrictMock<MockConnection>();
        mGazerListener.reset(new GazerListener(std::unique_ptr<Connection>(mConnection)));
    }

    StrictMock<MockConnection>* mConnection;
    std::unique_ptr<GazerListener> mGazerListener;
};

void registerMocksForCreateGazerConnection(MockConnection* connection) {
    EXPECT_CALL(*connection, recv(_, 4))
            // Replay id length
            .WillOnce(DoAll(WithArg<0>(SetVoidPointee(toByteVector<uint32_t>(0))), ReturnArg<1>()))
            // Replay length
            .WillOnce(DoAll(WithArg<0>(SetVoidPointee(toByteVector<uint32_t>(0))), ReturnArg<1>()));

    // Replay id
    EXPECT_CALL(*connection, recv(_, 0)).WillOnce(ReturnArg<1>());
}

}  // end of anonymous namespace

TEST_F(GazerListenerTest, AcceptConnection) {
    auto clientConnection = new StrictMock<MockConnection>();

    EXPECT_CALL(*mConnection, acceptProxy()).WillOnce(Return(clientConnection));

    registerMocksForCreateGazerConnection(clientConnection);

    EXPECT_THAT(mGazerListener->acceptConnection(), NotNull());
}

TEST_F(GazerListenerTest, AcceptConnectionErrorAccept) {
    EXPECT_CALL(*mConnection, acceptProxy()).WillOnce(Return(nullptr));

    EXPECT_THAT(mGazerListener->acceptConnection(), IsNull());
}

TEST_F(GazerListenerTest, AcceptConnectionErrorGazerConnection) {
    auto clientConnection1 = new StrictMock<MockConnection>();
    auto clientConnection2 = new StrictMock<MockConnection>();

    EXPECT_CALL(*mConnection, acceptProxy())
            .WillOnce(Return(clientConnection1))
            .WillOnce(Return(clientConnection2));

    // Replay id length failed
    EXPECT_CALL(*clientConnection1, recv(_, 4)).WillOnce(Return(2));

    registerMocksForCreateGazerConnection(clientConnection2);

    EXPECT_THAT(mGazerListener->acceptConnection(), NotNull());
}

}  // end of namespace test
}  // end of namespace caze
}  // end of namespace android
