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
using ::testing::ElementsAreArray;

namespace android {
namespace caze {
namespace test {
namespace {

const std::string replayId = "ABCDE";

class GazerConnectionTest : public ::testing::Test {
protected:
    virtual void SetUp() {
        mConnection = new MockConnection();
        mGazerConnection = createGazerConnection(mConnection, replayId, 0);
    }

    MockConnection* mConnection;
    std::unique_ptr<GazerConnection> mGazerConnection;
    std::vector<uint8_t> mBuffer;
};
}  // end of anonymous namespace

TEST(GazerConnectionTestStatic, Create) {
    auto connection = new MockConnection();
    uint32_t replayLength = 0x56003412;
    pushString(&connection->in, replayId);
    pushUint32(&connection->in, replayLength);

    auto gazerConnection = GazerConnection::create(std::unique_ptr<Connection>(connection));

    EXPECT_THAT(gazerConnection, NotNull());
    EXPECT_EQ(replayId, gazerConnection->replayId());
    EXPECT_EQ(replayLength, gazerConnection->replayLength());
}

TEST(GazerConnectionTestStatic, CreateErrorReadReplayId) {
    auto connection = new MockConnection();
    pushUint8(&connection->in, 'A');
    // Replay id read failed
    auto gazerConnection = GazerConnection::create(std::unique_ptr<Connection>(connection));

    EXPECT_THAT(gazerConnection, IsNull());
}

TEST_F(GazerConnectionTest, Get) {
    std::vector<uint8_t> resourceContent{1, 2, 3};
    mBuffer.resize(resourceContent.size());

    std::vector<uint8_t> expected;
    pushUint8(&expected, GazerConnection::MESSAGE_TYPE_GET);
    pushUint32(&expected, 2);
    pushString(&expected, "A");
    pushString(&expected, "B");

    pushBytes(&mConnection->in, resourceContent);

    EXPECT_TRUE(mGazerConnection->getResources({"A", "B"}, mBuffer.data(), mBuffer.size()));
    EXPECT_THAT(mBuffer, ElementsAreArray(resourceContent));
    EXPECT_EQ(mConnection->out, expected);
}

TEST_F(GazerConnectionTest, GetErrorMessageType) {
    mBuffer.resize(3);
    mConnection->out_limit = 0;
    EXPECT_FALSE(mGazerConnection->getResources({"A", "B"}, mBuffer.data(), mBuffer.size()));
}

TEST_F(GazerConnectionTest, GetErrorCount) {
    mBuffer.resize(3);
    mConnection->out_limit = 1;
    EXPECT_FALSE(mGazerConnection->getResources({"A", "B"}, mBuffer.data(), mBuffer.size()));
}

TEST_F(GazerConnectionTest, GetErrorId) {
    mBuffer.resize(3);
    mConnection->out_limit = 3;
    EXPECT_FALSE(mGazerConnection->getResources({"A", "B"}, mBuffer.data(), mBuffer.size()));
}

TEST_F(GazerConnectionTest, GetErrorContent) {
    mBuffer.resize(3);
    mConnection->out_limit = 6;
    EXPECT_FALSE(mGazerConnection->getResources({"A", "B"}, mBuffer.data(), mBuffer.size()));
}

TEST_F(GazerConnectionTest, Post) {
    std::vector<uint8_t> postData{1, 2, 3};

    std::vector<uint8_t> expected;
    pushUint8(&expected, GazerConnection::MESSAGE_TYPE_POST);
    pushUint32(&expected, 3);
    pushBytes(&expected, postData);

    EXPECT_TRUE(mGazerConnection->post(&postData.front(), postData.size()));
    EXPECT_EQ(mConnection->out, expected);
}

TEST_F(GazerConnectionTest, PostErrorMessageType) {
    std::vector<uint8_t> postData{1, 2, 3};
    mConnection->out_limit = 0;
    EXPECT_FALSE(mGazerConnection->post(&postData.front(), postData.size()));
}

TEST_F(GazerConnectionTest, PostErrorSize) {
    std::vector<uint8_t> postData{1, 2, 3};
    mConnection->out_limit = 1;
    EXPECT_FALSE(mGazerConnection->post(&postData.front(), postData.size()));
}

TEST_F(GazerConnectionTest, PostErrorData) {
    std::vector<uint8_t> postData{1, 2, 3};
    mConnection->out_limit = 4;
    EXPECT_FALSE(mGazerConnection->post(&postData.front(), postData.size()));
}

}  // end of namespace test
}  // end of namespace caze
}  // end of namespace android
