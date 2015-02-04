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
#include "Connection.h"
#include "GazerConnection.h"
#include "Interpreter.h"
#include "MockConnection.h"
#include "TestUtilities.h"

#include <memory>
#include <string>
#include <vector>

#include <gmock/gmock.h>

using ::testing::_;
using ::testing::DoAll;
using ::testing::NotNull;
using ::testing::Return;
using ::testing::ReturnArg;
using ::testing::StrictMock;
using ::testing::WithArg;

namespace android {
namespace caze {
namespace test {

uint32_t instruction(Interpreter::InstructionCode code) {
    return static_cast<uint32_t>(code) << 26;
}

uint32_t instruction(Interpreter::InstructionCode code, uint32_t data) {
    return (static_cast<uint32_t>(code) << 26) | (data & 0x03ffffffU);
}

uint32_t instruction(Interpreter::InstructionCode code, BaseType type, uint32_t data) {
    return (static_cast<uint32_t>(code) << 26) | (static_cast<uint32_t>(type) << 20) |
           (data & 0x000fffff);
}

namespace {
void addUint32ToUint8Vector(std::vector<uint8_t>* v, uint32_t x) {
    for (uint8_t i = 0; i < 32; i += 8) {
        v->push_back((x >> i) & 0xff);
    }
}
}

template <>
std::vector<uint8_t> toByteVector<uint8_t>(const uint8_t& x) {
    return {x};
}

template <>
std::vector<uint8_t> toByteVector<uint32_t>(const uint32_t& x) {
    std::vector<uint8_t> v;
    addUint32ToUint8Vector(&v, x);
    return v;
}

template <>
std::vector<uint8_t> toByteVector<std::string>(const std::string& x) {
    std::vector<uint8_t> v;
    v.insert(v.end(), x.begin(), x.end());
    return v;
}

template <>
std::vector<uint8_t> toByteVector<std::vector<uint32_t>>(const std::vector<uint32_t>& x) {
    std::vector<uint8_t> v;
    for (auto& it : x) {
        std::vector<uint8_t> t = toByteVector(it);
        v.insert(v.end(), t.begin(), t.end());
    }
    return v;
}

std::vector<uint8_t> createReplayData(uint32_t stackSize, uint32_t volatileMemorySize,
                                      const std::vector<uint8_t>& constantMemory,
                                      const ResourceProvider::ResourceList& resources,
                                      const std::vector<uint32_t>& instructions) {
    std::vector<uint8_t> replayData;
    addUint32ToUint8Vector(&replayData, stackSize);
    addUint32ToUint8Vector(&replayData, volatileMemorySize);
    addUint32ToUint8Vector(&replayData, constantMemory.size());
    replayData.insert(replayData.end(), constantMemory.begin(), constantMemory.end());
    addUint32ToUint8Vector(&replayData, resources.size());
    for (auto& it : resources) {
        addUint32ToUint8Vector(&replayData, it.first.size());
        replayData.insert(replayData.end(), it.first.begin(), it.first.end());
        addUint32ToUint8Vector(&replayData, it.second);
    }
    addUint32ToUint8Vector(&replayData, instructions.size() * sizeof(uint32_t));
    for (auto it : instructions) {
        addUint32ToUint8Vector(&replayData, it);
    }
    return replayData;
}

std::unique_ptr<GazerConnection> createGazerConnection(MockConnection* connection,
                                                       const std::string& replayId,
                                                       uint32_t replayLength) {
    std::vector<uint8_t> replayIdData;
    std::vector<uint8_t> replayIdLengthData;
    std::vector<uint8_t> replayLengthData;

    replayIdData.insert(replayIdData.end(), replayId.begin(), replayId.end());
    addUint32ToUint8Vector(&replayIdLengthData, replayId.size());
    addUint32ToUint8Vector(&replayLengthData, replayLength);

    EXPECT_CALL(*connection, recv(_, _))
            // Replay id length
            .WillOnce(DoAll(WithArg<0>(SetVoidPointee(replayIdLengthData)), ReturnArg<1>()))
            // Replay id
            .WillOnce(DoAll(WithArg<0>(SetVoidPointee(replayIdData)), ReturnArg<1>()))
            // Replay length
            .WillOnce(DoAll(WithArg<0>(SetVoidPointee(replayLengthData)), ReturnArg<1>()));

    std::unique_ptr<GazerConnection> gazerConnection =
            GazerConnection::create(std::unique_ptr<Connection>(connection));

    EXPECT_THAT(gazerConnection, NotNull());

    return std::move(gazerConnection);
}

std::unique_ptr<GazerConnection> createGazerConnection(const std::string& replayId,
                                                       uint32_t replayLength) {
    StrictMock<MockConnection>* connection = new StrictMock<MockConnection>();
    return createGazerConnection(connection, replayId, replayLength);
}

}  // end of namespace test
}  // end of namespace caze
}  // end of namespace android
