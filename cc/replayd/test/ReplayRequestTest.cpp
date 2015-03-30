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
#include "MemoryManager.h"
#include "MockConnection.h"
#include "MockResourceProvider.h"
#include "ReplayRequest.h"
#include "TestUtilities.h"

#include <memory>
#include <string>
#include <vector>

#include <gmock/gmock.h>
#include <gtest/gtest.h>

using ::testing::_;
using ::testing::DoAll;
using ::testing::Eq;
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

const size_t MEMORY_SIZE = 4096;
const std::string replayId = "ABCDE";

}  // end of anonymous namespace

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

    std::unique_ptr<MemoryManager> memoryManager(new MemoryManager({MEMORY_SIZE}));

    auto gazerConnection = GazerConnection::create(std::unique_ptr<Connection>(connection));
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

    std::unique_ptr<MemoryManager> memoryManager(new MemoryManager({MEMORY_SIZE}));

    auto gazerConnection = GazerConnection::create(std::unique_ptr<Connection>(connection));
    auto replayRequest =
            ReplayRequest::create(*gazerConnection, resourceProvider.get(), memoryManager.get());

    EXPECT_THAT(gazerConnection, NotNull());
    EXPECT_EQ(nullptr, replayRequest);
}

}  // end of namespace test
}  // end of namespace caze
}  // end of namespace android
