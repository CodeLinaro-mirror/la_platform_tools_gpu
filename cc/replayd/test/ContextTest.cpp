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

#include "BaseType.h"
#include "Context.h"
#include "Interpreter.h"
#include "MemoryManager.h"
#include "MockConnection.h"
#include "MockResourceProvider.h"
#include "ResourceProvider.h"
#include "TestUtilities.h"

#include <memory>
#include <vector>

#include <gmock/gmock.h>
#include <gtest/gtest.h>

using ::testing::_;
using ::testing::DoAll;
using ::testing::Eq;
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

const size_t MEMORY_SIZE = 4096;

class ContextTest : public ::testing::Test {
protected:
    virtual void SetUp() {
        mMemoryManager.reset(new MemoryManager({MEMORY_SIZE}));
        mResourceProvider.reset(new StrictMock<MockResourceProvider>());
    }

    std::unique_ptr<MemoryManager> mMemoryManager;
    std::unique_ptr<StrictMock<MockResourceProvider>> mResourceProvider;
};

void resourceProviderLoadReplay(MockResourceProvider* resourceProvider,
                                const std::vector<uint8_t>& replayData) {
    EXPECT_CALL(*resourceProvider, get(Eq(""), _, _, replayData.size()))
            .WillOnce(DoAll(WithArg<2>(SetVoidPointee(replayData)), ReturnArg<3>()));
    EXPECT_CALL(*resourceProvider, prefetch(_, _, _, _)).WillOnce(Return(true));
}

}  // end of anonymous namespace

TEST_F(ContextTest, Create) {
    auto replayData = createReplayData(0, 0, {}, {}, {});

    resourceProviderLoadReplay(mResourceProvider.get(), replayData);
    auto gazerConnection = createGazerConnection("", replayData.size());
    auto context = Context::create(*gazerConnection, mResourceProvider.get(), mMemoryManager.get());

    EXPECT_THAT(context, NotNull());
}

TEST_F(ContextTest, CreateErrorReplayRequest) {
    // Failed to load
    EXPECT_CALL(*mResourceProvider, get(Eq(""), _, _, 10)).WillOnce(Return(0));

    auto gazerConnection = createGazerConnection("", 10);
    auto context = Context::create(*gazerConnection, mResourceProvider.get(), mMemoryManager.get());

    EXPECT_THAT(context, IsNull());
}

TEST_F(ContextTest, CreateErrorVolatileMemory) {
    auto replayData = createReplayData(0, MEMORY_SIZE, {}, {}, {});

    EXPECT_CALL(*mResourceProvider, get(Eq(""), _, _, replayData.size()))
            .WillOnce(DoAll(WithArg<2>(SetVoidPointee(replayData)), ReturnArg<3>()));

    auto gazerConnection = createGazerConnection("", replayData.size());
    auto context = Context::create(*gazerConnection, mResourceProvider.get(), mMemoryManager.get());

    EXPECT_THAT(context, IsNull());
}

TEST_F(ContextTest, LoadResource) {
    auto replayData = createReplayData(
            128, 1024, {}, {{"A", 4}},
            {instruction(Interpreter::InstructionCode::PUSH_I, BaseType::VolatilePointer, 0),
             instruction(Interpreter::InstructionCode::RESOURCE, 0)});

    std::vector<uint8_t> resourceA{1, 2, 3, 4};

    resourceProviderLoadReplay(mResourceProvider.get(), replayData);

    EXPECT_CALL(*mResourceProvider, get(Eq("A"), _, _, 4))
            .WillOnce(DoAll(WithArg<2>(SetVoidPointee(resourceA)), ReturnArg<3>()));

    auto gazerConnection = createGazerConnection("", replayData.size());
    auto context = Context::create(*gazerConnection, mResourceProvider.get(), mMemoryManager.get());

    EXPECT_THAT(context, NotNull());
    EXPECT_TRUE(context->interpret());
    EXPECT_THAT(mMemoryManager->volatileToAbsolute(0), VoidPointee(resourceA));
}

TEST_F(ContextTest, LoadResourcePopFailed) {
    auto replayData = createReplayData(128, 1024, {}, {{"A", 4}},
                                       {instruction(Interpreter::InstructionCode::RESOURCE, 0)});

    resourceProviderLoadReplay(mResourceProvider.get(), replayData);
    auto gazerConnection = createGazerConnection("", replayData.size());
    auto context = Context::create(*gazerConnection, mResourceProvider.get(), mMemoryManager.get());

    EXPECT_THAT(context, NotNull());
    EXPECT_FALSE(context->interpret());
}

TEST_F(ContextTest, LoadResourceGetFailed) {
    auto replayData = createReplayData(
            128, 1024, {}, {{"A", 4}},
            {instruction(Interpreter::InstructionCode::PUSH_I, BaseType::VolatilePointer, 0),
             instruction(Interpreter::InstructionCode::RESOURCE, 0)});

    resourceProviderLoadReplay(mResourceProvider.get(), replayData);

    EXPECT_CALL(*mResourceProvider, get(Eq("A"), _, _, 4)).WillOnce(Return(0));

    auto gazerConnection = createGazerConnection("", replayData.size());
    auto context = Context::create(*gazerConnection, mResourceProvider.get(), mMemoryManager.get());

    EXPECT_THAT(context, NotNull());
    EXPECT_FALSE(context->interpret());
}

TEST_F(ContextTest, PostData) {
    auto replayData = createReplayData(
            128, 1024, {0, 1, 2, 3, 4, 5, 6, 7}, {},
            {instruction(Interpreter::InstructionCode::PUSH_I, BaseType::ConstantPointer, 1),
             instruction(Interpreter::InstructionCode::PUSH_I, BaseType::Uint32, 6),
             instruction(Interpreter::InstructionCode::POST)});

    auto connection = new StrictMock<MockConnection>();

    EXPECT_CALL(*connection, send(VoidPointee(std::vector<uint8_t>{1}), 1))
            .WillOnce(ReturnArg<1>());
    EXPECT_CALL(*connection, send(VoidPointee(toByteVector<uint32_t>(6)), 4))
            .WillOnce(ReturnArg<1>());
    EXPECT_CALL(*connection, send(VoidPointee(std::vector<uint8_t>{1, 2, 3, 4, 5, 6}), 6))
            .WillOnce(ReturnArg<1>());

    resourceProviderLoadReplay(mResourceProvider.get(), replayData);
    auto gazerConnection = createGazerConnection(connection, "", replayData.size());
    auto context = Context::create(*gazerConnection, mResourceProvider.get(), mMemoryManager.get());

    EXPECT_THAT(context, NotNull());
    EXPECT_TRUE(context->interpret());
}

TEST_F(ContextTest, PostDataErrorPop) {
    auto replayData = createReplayData(
            128, 1024, {0, 1, 2, 3, 4, 5, 6, 7}, {},
            {instruction(Interpreter::InstructionCode::PUSH_I, BaseType::ConstantPointer, 1),
             instruction(Interpreter::InstructionCode::PUSH_I, BaseType::Uint8, 6),  // Wrong type
             instruction(Interpreter::InstructionCode::POST)});

    resourceProviderLoadReplay(mResourceProvider.get(), replayData);
    auto gazerConnection = createGazerConnection("", replayData.size());
    auto context = Context::create(*gazerConnection, mResourceProvider.get(), mMemoryManager.get());

    EXPECT_THAT(context, NotNull());
    EXPECT_FALSE(context->interpret());
}

TEST_F(ContextTest, PostDataErrorPost) {
    auto replayData = createReplayData(
            128, 1024, {0, 1, 2, 3, 4, 5, 6, 7}, {},
            {instruction(Interpreter::InstructionCode::PUSH_I, BaseType::ConstantPointer, 1),
             instruction(Interpreter::InstructionCode::PUSH_I, BaseType::Uint32, 6),
             instruction(Interpreter::InstructionCode::POST)});

    auto connection = new StrictMock<MockConnection>();

    EXPECT_CALL(*connection, send(VoidPointee(std::vector<uint8_t>{1}), 1))
            .WillOnce(ReturnArg<1>());
    EXPECT_CALL(*connection, send(VoidPointee(std::vector<uint8_t>{6, 0, 0, 0}), 4))
            .WillOnce(ReturnArg<1>());
    // Send failed
    EXPECT_CALL(*connection, send(VoidPointee(std::vector<uint8_t>{1, 2, 3, 4, 5, 6}), 6))
            .WillOnce(Return(0));

    resourceProviderLoadReplay(mResourceProvider.get(), replayData);

    auto gazerConnection = createGazerConnection(connection, "", replayData.size());
    auto context = Context::create(*gazerConnection, mResourceProvider.get(), mMemoryManager.get());

    EXPECT_THAT(context, NotNull());
    EXPECT_FALSE(context->interpret());
}

}  // end of namespace test
}  // end of namespace caze
}  // end of namespace android
