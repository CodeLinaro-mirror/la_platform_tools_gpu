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

#include "MockConnection.h"
#include "TestUtilities.h"

#include <string>

#include <gmock/gmock.h>
#include <gtest/gtest.h>

using ::testing::_;
using ::testing::DoAll;
using ::testing::Return;
using ::testing::ReturnArg;
using ::testing::StrictMock;
using ::testing::WithArg;

namespace android {
namespace caze {
namespace test {
namespace {

const std::string testString = "ABCDE";

class ConnectionTest : public ::testing::Test {
protected:
    virtual void SetUp() {
        mConnection.reset(new StrictMock<MockConnection>());
    }

    std::unique_ptr<StrictMock<MockConnection>> mConnection;
};

}  // end of anonymous namespace

TEST_F(ConnectionTest, SendEmptyString) {
    // Length of the string
    EXPECT_CALL(*mConnection, send(VoidPointee(toByteVector<uint32_t>(0)), 4))
            .WillOnce(ReturnArg<1>());
    // Content of the string
    EXPECT_CALL(*mConnection, send(_, 0)).WillOnce(ReturnArg<1>());

    EXPECT_TRUE(mConnection->sendString(""));
}

TEST_F(ConnectionTest, SendString) {
    // Length of the string
    EXPECT_CALL(*mConnection, send(VoidPointee(toByteVector<uint32_t>(testString.size())), 4))
            .WillOnce(ReturnArg<1>());
    // Content of the string
    EXPECT_CALL(*mConnection, send(VoidPointee(toByteVector(testString)), testString.size()))
            .WillOnce(ReturnArg<1>());

    EXPECT_TRUE(mConnection->sendString(testString));
}

TEST_F(ConnectionTest, SendStringErrorLength) {
    // Length of the string send failed (only 2 byte sent out of 4)
    EXPECT_CALL(*mConnection, send(VoidPointee(toByteVector<uint32_t>(testString.size())), 4))
            .WillOnce(Return(2));

    EXPECT_FALSE(mConnection->sendString(testString));
}

TEST_F(ConnectionTest, SendStringErrorContent) {
    // Length of the string
    EXPECT_CALL(*mConnection, send(VoidPointee(toByteVector<uint32_t>(testString.size())), 4))
            .WillOnce(ReturnArg<1>());
    // Content of the string send failed (0 byte sent out of 5)
    EXPECT_CALL(*mConnection, send(VoidPointee(toByteVector(testString)), testString.size()))
            .WillOnce(Return(0));

    EXPECT_FALSE(mConnection->sendString(testString));
}

TEST_F(ConnectionTest, ReadEmptyString) {
    // Length of the string
    EXPECT_CALL(*mConnection, recv(_, 4))
            .WillOnce(DoAll(WithArg<0>(SetVoidPointee(toByteVector<uint32_t>(0))), ReturnArg<1>()));
    // Content of the string
    EXPECT_CALL(*mConnection, recv(_, 0)).WillOnce(ReturnArg<1>());

    std::string s;
    EXPECT_TRUE(mConnection->readString(&s));
    EXPECT_EQ("", s);
}

TEST_F(ConnectionTest, ReadString) {
    // Length of the string
    EXPECT_CALL(*mConnection, recv(_, 4))
            .WillOnce(DoAll(WithArg<0>(SetVoidPointee(toByteVector<uint32_t>(testString.size()))),
                            ReturnArg<1>()));
    // Content of the string
    EXPECT_CALL(*mConnection, recv(_, testString.size()))
            .WillOnce(DoAll(WithArg<0>(SetVoidPointee(toByteVector(testString))), ReturnArg<1>()));

    std::string s;
    EXPECT_TRUE(mConnection->readString(&s));
    EXPECT_EQ(testString, s);
}

TEST_F(ConnectionTest, ReadStringLengthError) {
    // Length of the string read failed (2 out of 4 byte read)
    EXPECT_CALL(*mConnection, recv(_, 4)).WillOnce(Return(2));

    std::string s;
    EXPECT_FALSE(mConnection->readString(&s));
}

TEST_F(ConnectionTest, ReadStringContentError) {
    // Length of the string
    EXPECT_CALL(*mConnection, recv(_, 4))
            .WillOnce(DoAll(WithArg<0>(SetVoidPointee(toByteVector<uint32_t>(testString.size()))),
                            ReturnArg<1>()));
    // Content of the string read failed (3 out of 5 byte read)
    EXPECT_CALL(*mConnection, recv(_, testString.size())).WillOnce(Return(3));

    std::string s;
    EXPECT_FALSE(mConnection->readString(&s));
}

}  // end of namespace test
}  // end of namespace caze
}  // end of namespace android
