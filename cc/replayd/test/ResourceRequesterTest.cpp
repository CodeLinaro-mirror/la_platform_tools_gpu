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

#include "GazerConnection.h"
#include "MockConnection.h"
#include "ResourceProvider.h"
#include "ResourceRequester.h"
#include "TestUtilities.h"

#include <memory>
#include <string>
#include <vector>

#include <gmock/gmock.h>
#include <gtest/gtest.h>

using ::testing::_;
using ::testing::DoAll;
using ::testing::Return;
using ::testing::ReturnArg;
using ::testing::StrictMock;
using ::testing::ElementsAreArray;

namespace android {
namespace caze {
namespace test {
namespace {

class ResourceRequesterTest : public ::testing::Test {
protected:
    virtual void SetUp() {
        mConnection = new MockConnection();
        mGazer = createGazerConnection(mConnection, "", 0);
        mResourceProvider = ResourceRequester::create();
    }

    MockConnection* mConnection;
    std::unique_ptr<GazerConnection> mGazer;
    std::unique_ptr<ResourceProvider> mResourceProvider;
    std::vector<uint8_t> mBuffer;
};

}  // end of anonymous namespace

TEST_F(ResourceRequesterTest, Prefetch) {
    mBuffer.resize(4096);
    EXPECT_TRUE(mResourceProvider->prefetch({{"A", 16}, {"B", 32}}, *mGazer, mBuffer.data(),
                                            mBuffer.size()));
}

TEST_F(ResourceRequesterTest, SingleGet) {
    std::vector<uint8_t> payload = {'X', 'Y', 'Z'};
    mBuffer.resize(payload.size());
    std::vector<uint8_t> expected;
    pushUint8(&expected, GazerConnection::MESSAGE_TYPE_GET);
    pushUint32(&expected, 1);
    pushString(&expected, "A");

    pushBytes(&mConnection->in, payload);

    EXPECT_TRUE(mResourceProvider->get("A", *mGazer, mBuffer.data(), 3));
    EXPECT_THAT(mBuffer, ElementsAreArray(payload));
    EXPECT_EQ(mConnection->out, expected);
}

TEST_F(ResourceRequesterTest, MultiGet) {
    std::vector<uint8_t> payload = {'X', 'Y', 'Z', '1', '2', '3', '4', '5'};
    mBuffer.resize(payload.size());
    std::vector<uint8_t> expected;
    pushUint8(&expected, GazerConnection::MESSAGE_TYPE_GET);
    pushUint32(&expected, 2);
    pushString(&expected, "A");
    pushString(&expected, "B");

    pushBytes(&mConnection->in, payload);

    EXPECT_TRUE(mResourceProvider->get({{"A", 3}, {"B", 5}}, *mGazer, mBuffer.data(), mBuffer.size()));
    EXPECT_THAT(mBuffer, ElementsAreArray(payload));
    EXPECT_EQ(mConnection->out, expected);
}
}  // end of namespace test
}  // end of namespace caze
}  // end of namespace android
