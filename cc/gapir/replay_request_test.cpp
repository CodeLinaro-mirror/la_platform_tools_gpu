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

#include "memory_manager.h"
#include "mock_connection.h"
#include "mock_resource_provider.h"
#include "replay_request.h"
#include "server_connection.h"
#include "server_listener.h"
#include "test_utilities.h"

#include <gmock/gmock.h>
#include <gtest/gtest.h>

#include <memory>
#include <string>
#include <vector>

using ::testing::_;
using ::testing::DoAll;
using ::testing::Eq;
using ::testing::NotNull;
using ::testing::Return;
using ::testing::ReturnArg;
using ::testing::StrictMock;
using ::testing::WithArg;
using ::testing::ElementsAreArray;

namespace gapir {
namespace test {
namespace {

const uint32_t MEMORY_SIZE = 4096;
const std::string replayId = "ABCDE";

}  // anonymous namespace

TEST(ReplayRequestTestStatic, Create) {
    uint32_t stackSize = 128;
    uint32_t volatileMemorySize = 1024;
    std::vector<uint8_t> constantMemory =
        {'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'};
    ResourceProvider::ResourceList resources{{"ZYX", 16}, {"1234", 32}};
    std::vector<uint32_t> instructionList{0, 1, 2};

    auto replayData = createReplayData(stackSize, volatileMemorySize, constantMemory, resources,
                                       instructionList);

    auto connection = new MockConnection();
    std::unique_ptr<StrictMock<MockResourceProvider>> resourceProvider(
            new StrictMock<MockResourceProvider>());
    pushString(&connection->in, replayId);
    pushUint32(&connection->in, replayData.size());

    EXPECT_CALL(*resourceProvider, get(Eq(replayId), _, _, replayData.size()))
            // Replay data
            .WillOnce(DoAll(WithArg<2>(SetVoidPointee(replayData)), ReturnArg<3>()));

    std::vector<uint32_t> memorySizes = {MEMORY_SIZE};
    std::unique_ptr<MemoryManager> memoryManager(new MemoryManager(memorySizes));

    auto gazerConnection = ServerConnection::create(std::unique_ptr<gapic::Connection>(connection));
    auto replayRequest =
            ReplayRequest::create(*gazerConnection, resourceProvider.get(), memoryManager.get());

    EXPECT_THAT(gazerConnection, NotNull());
    EXPECT_THAT(replayRequest, NotNull());

    EXPECT_EQ(stackSize, replayRequest->getStackSize());
    EXPECT_EQ(volatileMemorySize, replayRequest->getVolatileMemorySize());
    EXPECT_EQ(resources, replayRequest->getResources());
    EXPECT_THAT(constantMemory,
        ElementsAreArray((uint8_t*)(replayRequest->getConstantMemory().first), replayRequest->getConstantMemory().second));
    EXPECT_THAT(instructionList,
        ElementsAreArray(replayRequest->getInstructionList().first, replayRequest->getInstructionList().second));
}

TEST(ReplayRequestTestStatic, CreateErrorGet) {
    uint32_t replayLength = 255;
    MockConnection* connection = new MockConnection();
    std::unique_ptr<StrictMock<MockResourceProvider>> resourceProvider(
            new StrictMock<MockResourceProvider>());

    pushString(&connection->in, replayId);
    pushUint32(&connection->in, replayLength);

    // Get replay request from resource provider fail
    EXPECT_CALL(*resourceProvider, get(Eq(replayId), _, _, replayLength)).WillOnce(Return(0));

    std::vector<uint32_t> memorySizes = {MEMORY_SIZE};
    std::unique_ptr<MemoryManager> memoryManager(new MemoryManager(memorySizes));

    auto gazerConnection = ServerConnection::create(std::unique_ptr<gapic::Connection>(connection));
    auto replayRequest =
            ReplayRequest::create(*gazerConnection, resourceProvider.get(), memoryManager.get());

    EXPECT_THAT(gazerConnection, NotNull());
    EXPECT_EQ(nullptr, replayRequest);
}

}  // namespace test
}  // namespace gapir
