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

namespace android {
namespace caze {
namespace test {
namespace {

const size_t BUFFER_SIZE = 4096;

class ResourceRequesterTest : public ::testing::Test {
protected:
    virtual void SetUp() {
        mConnection = new StrictMock<MockConnection>();
        mGazer = createGazerConnection(mConnection, "", 0);
        mResourceProvider = ResourceRequester::create();
        mBuffer.reset(new uint8_t[BUFFER_SIZE]);
    }

    StrictMock<MockConnection>* mConnection;
    std::unique_ptr<GazerConnection> mGazer;
    std::unique_ptr<ResourceProvider> mResourceProvider;
    std::unique_ptr<uint8_t[]> mBuffer;
};

}  // end of anonymous namespace

TEST_F(ResourceRequesterTest, Prefetch) {
    EXPECT_TRUE(mResourceProvider->prefetch({{"A", 16}, {"B", 32}}, *mGazer, mBuffer.get(),
                                            BUFFER_SIZE));
}

TEST_F(ResourceRequesterTest, SingleGet) {
    EXPECT_CALL(*mConnection, send(VoidPointee(std::vector<uint8_t>{0}), 1))
            .WillOnce(ReturnArg<1>());
    EXPECT_CALL(*mConnection, send(VoidPointee(toByteVector<uint32_t>(1)), 4))
            .WillOnce(ReturnArg<1>())   // Resource count
            .WillOnce(ReturnArg<1>());  // Resource id length
    EXPECT_CALL(*mConnection, send(VoidPointee(toByteVector<std::string>("A")), 1))
            .WillOnce(ReturnArg<1>());
    EXPECT_CALL(*mConnection, recv(_, 3))
            .WillOnce(DoAll(SetVoidPointee(toByteVector<std::string>("XYZ")), ReturnArg<1>()));
    EXPECT_TRUE(mResourceProvider->get("A", *mGazer, mBuffer.get(), 3));
    EXPECT_THAT(mBuffer.get(), VoidPointee(toByteVector<std::string>("XYZ")));
}

TEST_F(ResourceRequesterTest, MultiGet) {
    EXPECT_CALL(*mConnection, send(VoidPointee(std::vector<uint8_t>{0}), 1))
            .WillOnce(ReturnArg<1>());
    EXPECT_CALL(*mConnection, send(VoidPointee(std::vector<uint8_t>{2, 0, 0, 0}), 4))
            .WillOnce(ReturnArg<1>());  // Resource count
    EXPECT_CALL(*mConnection, send(VoidPointee(std::vector<uint8_t>{1, 0, 0, 0}), 4))
            .WillOnce(ReturnArg<1>())   // Resource 1 name length
            .WillOnce(ReturnArg<1>());  // Resource 2 name length
    // Resource 1 name
    EXPECT_CALL(*mConnection, send(VoidPointee(toByteVector<std::string>("A")), 1))
            .WillOnce(ReturnArg<1>());
    // Resource 2 name
    EXPECT_CALL(*mConnection, send(VoidPointee(toByteVector<std::string>("B")), 1))
            .WillOnce(ReturnArg<1>());
    // Resource content
    EXPECT_CALL(*mConnection, recv(_, 8))
            .WillOnce(DoAll(SetVoidPointee(toByteVector<std::string>("XYZ12345")), ReturnArg<1>()));
    EXPECT_TRUE(mResourceProvider->get({{"A", 3}, {"B", 5}}, *mGazer, mBuffer.get()));
    EXPECT_THAT(mBuffer.get(), VoidPointee(toByteVector<std::string>("XYZ12345")));
}

}  // end of namespace test
}  // end of namespace caze
}  // end of namespace android
