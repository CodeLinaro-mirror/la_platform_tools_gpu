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
#include "MockConnection.h"
#include "TestUtilities.h"

#include <memory>
#include <string>
#include <vector>

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
const std::string replayId = "ABCDE";

class GazerConnectionTest : public ::testing::Test {
protected:
    virtual void SetUp() {
        mConnection = new StrictMock<MockConnection>();
        mGazerConnection = createGazerConnection(mConnection, replayId, 0);
        mBuffer.reset(new uint8_t[BUFFER_SIZE]);
    }

    StrictMock<MockConnection>* mConnection;
    std::unique_ptr<GazerConnection> mGazerConnection;
    std::unique_ptr<uint8_t[]> mBuffer;
};
}  // end of anonymous namespace

TEST(GazerConnectionTestStatic, Create) {
    auto connection = new StrictMock<MockConnection>();
    uint32_t replayLength = 0x56003412;

    EXPECT_CALL(*connection, recv(_, 4))
            // Replay id length
            .WillOnce(DoAll(WithArg<0>(SetVoidPointee(toByteVector<uint32_t>(replayId.size()))),
                            ReturnArg<1>()))
            // Replay length
            .WillOnce(DoAll(WithArg<0>(SetVoidPointee(toByteVector<uint32_t>(replayLength))),
                            ReturnArg<1>()));

    EXPECT_CALL(*connection, recv(_, replayId.size()))
            // Replay id
            .WillOnce(DoAll(WithArg<0>(SetVoidPointee(toByteVector<std::string>(replayId))),
                            ReturnArg<1>()));

    auto gazerConnection = GazerConnection::create(std::unique_ptr<Connection>(connection));

    EXPECT_THAT(gazerConnection, NotNull());
    EXPECT_EQ(replayId, gazerConnection->replayId());
    EXPECT_EQ(replayLength, gazerConnection->replayLength());
}

TEST(GazerConnectionTestStatic, CreateErrorReadReplayId) {
    auto connection = new StrictMock<MockConnection>();

    EXPECT_CALL(*connection, recv(_, 4))
            // Replay id length
            .WillOnce(DoAll(WithArg<0>(SetVoidPointee(toByteVector<uint32_t>(replayId.size()))),
                            ReturnArg<1>()));

    // Replay id read failed (4 byte read out of 5)
    EXPECT_CALL(*connection, recv(_, replayId.size())).WillOnce(Return(4));

    auto gazerConnection = GazerConnection::create(std::unique_ptr<Connection>(connection));

    EXPECT_THAT(gazerConnection, IsNull());
}

TEST(GazerConnectionTestStatic, CreateErrorReadReplayLength) {
    auto connection = new StrictMock<MockConnection>();

    EXPECT_CALL(*connection, recv(_, 4))
            // Replay id length
            .WillOnce(DoAll(WithArg<0>(SetVoidPointee(toByteVector<uint32_t>(replayId.size()))),
                            ReturnArg<1>()))
            // Replay length read failed (1 byte out of 4 read)
            .WillOnce(Return(1));

    EXPECT_CALL(*connection, recv(_, replayId.size()))
            // Replay id
            .WillOnce(DoAll(WithArg<0>(SetVoidPointee(toByteVector(replayId))), ReturnArg<1>()));

    auto gazerConnection = GazerConnection::create(std::unique_ptr<Connection>(connection));

    EXPECT_THAT(gazerConnection, IsNull());
}

TEST_F(GazerConnectionTest, Get) {
    std::vector<uint8_t> resourceContent{1, 2, 3};

    // Message direction (GET: 0)
    EXPECT_CALL(*mConnection, send(VoidPointee(std::vector<uint8_t>{0}), 1))
            .WillOnce(ReturnArg<1>());

    // Resource count
    EXPECT_CALL(*mConnection, send(VoidPointee(toByteVector<uint32_t>(2)), 4))
            .WillOnce(ReturnArg<1>());

    // Length of the two resource id
    EXPECT_CALL(*mConnection, send(VoidPointee(toByteVector<uint32_t>(1)), 4))
            .WillOnce(ReturnArg<1>())
            .WillOnce(ReturnArg<1>());

    // First resource id ("A")
    EXPECT_CALL(*mConnection, send(VoidPointee(toByteVector<std::string>("A")), 1))
            .WillOnce(ReturnArg<1>());

    // Second resource id ("B")
    EXPECT_CALL(*mConnection, send(VoidPointee(toByteVector<std::string>("B")), 1))
            .WillOnce(ReturnArg<1>());

    // Resource content
    EXPECT_CALL(*mConnection, recv(_, 3))
            .WillOnce(DoAll(WithArg<0>(SetVoidPointee(resourceContent)), ReturnArg<1>()));

    EXPECT_TRUE(mGazerConnection->getResources({"A", "B"}, mBuffer.get(), 3));
    EXPECT_THAT(mBuffer.get(), VoidPointee(resourceContent));
}

TEST_F(GazerConnectionTest, GetErrorMessageType) {
    // Message direction (GET: 0) failed
    EXPECT_CALL(*mConnection, send(VoidPointee(std::vector<uint8_t>{0}), 1)).WillOnce(Return(0));

    EXPECT_FALSE(mGazerConnection->getResources({"A", "B"}, mBuffer.get(), 3));
}

TEST_F(GazerConnectionTest, GetErrorCount) {
    // Message direction (GET: 0)
    EXPECT_CALL(*mConnection, send(VoidPointee(std::vector<uint8_t>{0}), 1))
            .WillOnce(ReturnArg<1>());

    // Resource count failed
    EXPECT_CALL(*mConnection, send(VoidPointee(toByteVector<uint32_t>(2)), 4)).WillOnce(Return(2));

    EXPECT_FALSE(mGazerConnection->getResources({"A", "B"}, mBuffer.get(), 3));
}

TEST_F(GazerConnectionTest, GetErrorId) {
    // Message direction (GET: 0)
    EXPECT_CALL(*mConnection, send(VoidPointee(std::vector<uint8_t>{0}), 1))
            .WillOnce(ReturnArg<1>());

    // Resource count
    EXPECT_CALL(*mConnection, send(VoidPointee(toByteVector<uint32_t>(2)), 4))
            .WillOnce(ReturnArg<1>());

    // Length of the two resource id, second one fails
    EXPECT_CALL(*mConnection, send(VoidPointee(toByteVector<uint32_t>(1)), 4))
            .WillOnce(ReturnArg<1>())
            .WillOnce(Return(1));

    // First resource id ("A")
    EXPECT_CALL(*mConnection, send(VoidPointee(toByteVector<std::string>("A")), 1))
            .WillOnce(ReturnArg<1>());

    EXPECT_FALSE(mGazerConnection->getResources({"A", "B"}, mBuffer.get(), 3));
}

TEST_F(GazerConnectionTest, GetErrorContent) {
    // Message direction (GET: 0)
    EXPECT_CALL(*mConnection, send(VoidPointee(std::vector<uint8_t>{0}), 1))
            .WillOnce(ReturnArg<1>());

    // Resource count
    EXPECT_CALL(*mConnection, send(VoidPointee(toByteVector<uint32_t>(2)), 4))
            .WillOnce(ReturnArg<1>());

    // Length of the two resource id
    EXPECT_CALL(*mConnection, send(VoidPointee(toByteVector<uint32_t>(1)), 4))
            .WillOnce(ReturnArg<1>())
            .WillOnce(ReturnArg<1>());

    // First resource id ("A")
    EXPECT_CALL(*mConnection, send(VoidPointee(toByteVector<std::string>("A")), 1))
            .WillOnce(ReturnArg<1>());

    // Second resource id ("B")
    EXPECT_CALL(*mConnection, send(VoidPointee(toByteVector<std::string>("B")), 1))
            .WillOnce(ReturnArg<1>());

    // Resource content failed
    EXPECT_CALL(*mConnection, recv(_, 3)).WillOnce(Return(0));

    EXPECT_FALSE(mGazerConnection->getResources({"A", "B"}, mBuffer.get(), 3));
}

TEST_F(GazerConnectionTest, Post) {
    std::vector<uint8_t> postData{1, 2, 3};

    // Message direction (POST: 1)
    EXPECT_CALL(*mConnection, send(VoidPointee(std::vector<uint8_t>{1}), 1))
            .WillOnce(ReturnArg<1>());

    // Data size (3)
    EXPECT_CALL(*mConnection, send(VoidPointee(toByteVector<uint32_t>(postData.size())), 4))
            .WillOnce(ReturnArg<1>());

    // Data content
    EXPECT_CALL(*mConnection, send(VoidPointee(postData), postData.size()))
            .WillOnce(ReturnArg<1>());

    EXPECT_TRUE(mGazerConnection->post(&postData.front(), postData.size()));
}

TEST_F(GazerConnectionTest, PostErrorMessageType) {
    std::vector<uint8_t> postData{1, 2, 3};

    // Message direction (POST: 1) failed
    EXPECT_CALL(*mConnection, send(VoidPointee(std::vector<uint8_t>{1}), 1)).WillOnce(Return(0));

    EXPECT_FALSE(mGazerConnection->post(&postData.front(), postData.size()));
}

TEST_F(GazerConnectionTest, PostErrorSize) {
    std::vector<uint8_t> postData{1, 2, 3};

    // Message direction (POST: 1)
    EXPECT_CALL(*mConnection, send(VoidPointee(std::vector<uint8_t>{1}), 1))
            .WillOnce(ReturnArg<1>());

    // Data size (3) failed
    EXPECT_CALL(*mConnection, send(VoidPointee(toByteVector<uint32_t>(3)), 4)).WillOnce(Return(2));

    EXPECT_FALSE(mGazerConnection->post(&postData.front(), postData.size()));
}

TEST_F(GazerConnectionTest, PostErrorData) {
    std::vector<uint8_t> postData{1, 2, 3};

    // Message direction (POST: 1)
    EXPECT_CALL(*mConnection, send(VoidPointee(std::vector<uint8_t>{1}), 1))
            .WillOnce(ReturnArg<1>());

    // Data size (3)
    EXPECT_CALL(*mConnection, send(VoidPointee(toByteVector<uint32_t>(postData.size())), 4))
            .WillOnce(ReturnArg<1>());

    // Data content
    EXPECT_CALL(*mConnection, send(VoidPointee(postData), postData.size())).WillOnce(Return(0));

    EXPECT_FALSE(mGazerConnection->post(&postData.front(), postData.size()));
}

}  // end of namespace test
}  // end of namespace caze
}  // end of namespace android
