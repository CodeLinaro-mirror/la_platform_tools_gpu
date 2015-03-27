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
using ::testing::ElementsAre;

namespace android {
namespace caze {
namespace test {
namespace {

const std::string testString = "ABCDE";

class ConnectionTest : public ::testing::Test {
protected:
    virtual void SetUp() {
        mConnection.reset(new MockConnection());
    }

    std::unique_ptr<MockConnection> mConnection;
};

}  // end of anonymous namespace

TEST_F(ConnectionTest, SendEmptyString) {
    EXPECT_TRUE(mConnection->sendString(""));
    EXPECT_THAT(mConnection->out, ElementsAre(0, 0, 0, 0));
}

TEST_F(ConnectionTest, SendString) {
    EXPECT_TRUE(mConnection->sendString(testString));
    EXPECT_THAT(mConnection->out, ElementsAre(
        5, 0, 0, 0, 'A', 'B', 'C', 'D', 'E'));
}

TEST_F(ConnectionTest, SendStringError) {
    mConnection->out_limit = 3;
    EXPECT_FALSE(mConnection->sendString(testString));
}

TEST_F(ConnectionTest, ReadEmptyString) {
    pushString(&mConnection->in, "");
    std::string s;
    EXPECT_TRUE(mConnection->readString(&s));
    EXPECT_EQ("", s);
}

TEST_F(ConnectionTest, ReadString) {
    pushString(&mConnection->in, testString);
    std::string s;
    EXPECT_TRUE(mConnection->readString(&s));
    EXPECT_EQ(testString, s);
}

TEST_F(ConnectionTest, ReadStringError) {
    pushBytes(&mConnection->in, {'A', 'B'});
    std::string s;
    EXPECT_FALSE(mConnection->readString(&s));
}

}  // end of namespace test
}  // end of namespace caze
}  // end of namespace android
