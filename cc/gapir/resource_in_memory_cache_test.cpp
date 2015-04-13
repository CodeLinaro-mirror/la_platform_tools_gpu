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
#include "mock_resource_provider.h"
#include "resource_in_memory_cache.h"
#include "resource_provider.h"
#include "server_connection.h"
#include "test_utilities.h"

#include <memory>
#include <vector>

#include <gmock/gmock.h>
#include <gtest/gtest.h>

using ::testing::_;
using ::testing::ContainerEq;
using ::testing::DoAll;
using ::testing::Eq;
using ::testing::Return;
using ::testing::ReturnArg;
using ::testing::StrictMock;
using ::testing::WithArg;

namespace gapir {
namespace test {
namespace {

const uint32_t MEMORY_SIZE = 4096;
const uint32_t CACHE_SIZE = 2048;

class ResourceInMemoryCacheTest : public ::testing::Test {
protected:
    virtual void SetUp() {
        mMemoryManager.reset(new MemoryManager({MEMORY_SIZE}));
        mMemoryManager->setVolatileMemory(MEMORY_SIZE - CACHE_SIZE);

        mFallbackProvider = new StrictMock<MockResourceProvider>();

        mResourceInMemoryCache = ResourceInMemoryCache::create(
                std::unique_ptr<StrictMock<MockResourceProvider>>(mFallbackProvider),
                mMemoryManager->getBaseAddress());
        mResourceInMemoryCache->updateSize(CACHE_SIZE);
        mServer = createServerConnection("", 0);
    }

    StrictMock<MockResourceProvider>* mFallbackProvider;

    std::unique_ptr<MemoryManager> mMemoryManager;
    std::unique_ptr<ResourceInMemoryCache> mResourceInMemoryCache;
    std::unique_ptr<ServerConnection> mServer;
};

}  // anonymous namespace

TEST_F(ResourceInMemoryCacheTest, Prefetch) {
    ResourceProvider::ResourceList expected{{"A", 64}, {"D", 1024}, {"E", 2048}};

    EXPECT_CALL(*mFallbackProvider, get(_, _, _, _)).WillRepeatedly(Return(true));
    EXPECT_CALL(*mFallbackProvider, prefetch(ContainerEq(expected), _, _, _))
            .WillOnce(Return(true));

    mResourceInMemoryCache->get("C", *mServer, mMemoryManager->getVolatileAddress(), 512);
    mResourceInMemoryCache->get("B", *mServer, mMemoryManager->getVolatileAddress(), 256);
    mResourceInMemoryCache->prefetch({{"A", 64}, {"B", 256}, {"C", 512}, {"D", 1024}, {"E", 2048}},
                                     *mServer, mMemoryManager->getVolatileAddress(),
                                     MEMORY_SIZE - CACHE_SIZE);
}

TEST_F(ResourceInMemoryCacheTest, Resize) {
    bool res;

    // Fill the cache with three element
    EXPECT_CALL(*mFallbackProvider, get(_, _, _, _)).WillRepeatedly(Return(true));

    res = mResourceInMemoryCache->get("A", *mServer, mMemoryManager->getVolatileAddress(), 256);
    EXPECT_TRUE(res);

    res = mResourceInMemoryCache->get("B", *mServer, mMemoryManager->getVolatileAddress(), 512);
    EXPECT_TRUE(res);

    res = mResourceInMemoryCache->get("C", *mServer, mMemoryManager->getVolatileAddress(), 1024);
    EXPECT_TRUE(res);

    // New elements are not loaded
    EXPECT_CALL(*mFallbackProvider, get(_, _, _, _)).WillRepeatedly(Return(false));

    // Reduce cache size what removes the last two element from the cache
    mResourceInMemoryCache->updateSize(512);

    res = mResourceInMemoryCache->get("A", *mServer, mMemoryManager->getVolatileAddress(), 256);
    EXPECT_TRUE(res);

    res = mResourceInMemoryCache->get("B", *mServer, mMemoryManager->getVolatileAddress(), 512);
    EXPECT_FALSE(res);

    res = mResourceInMemoryCache->get("C", *mServer, mMemoryManager->getVolatileAddress(), 1024);
    EXPECT_FALSE(res);

    // Reduce cache size what removes all remaining element from the cache
    mResourceInMemoryCache->updateSize(128);

    res = mResourceInMemoryCache->get("A", *mServer, mMemoryManager->getVolatileAddress(), 256);
    EXPECT_FALSE(res);

    res = mResourceInMemoryCache->get("B", *mServer, mMemoryManager->getVolatileAddress(), 512);
    EXPECT_FALSE(res);

    res = mResourceInMemoryCache->get("C", *mServer, mMemoryManager->getVolatileAddress(), 1024);
    EXPECT_FALSE(res);
}

TEST_F(ResourceInMemoryCacheTest, CachingLogic) {
    bool res;

    EXPECT_CALL(*mFallbackProvider, get(_, _, _, _)).Times(6).WillRepeatedly(Return(true));

    // Call get {1st} (not save to cache because not fit)
    res = mResourceInMemoryCache->get("A", *mServer, mMemoryManager->getVolatileAddress(), 4096);
    EXPECT_TRUE(res);

    // Call get {2nd}
    res = mResourceInMemoryCache->get("A", *mServer, mMemoryManager->getVolatileAddress(), 512);
    EXPECT_TRUE(res);

    // Not call get
    res = mResourceInMemoryCache->get("A", *mServer, mMemoryManager->getVolatileAddress(), 512);
    EXPECT_TRUE(res);

    // Call get {3rd}
    res = mResourceInMemoryCache->get("B", *mServer, mMemoryManager->getVolatileAddress(), 1024);
    EXPECT_TRUE(res);

    // Not call get
    res = mResourceInMemoryCache->get("A", *mServer, mMemoryManager->getVolatileAddress(), 512);
    EXPECT_TRUE(res);

    // Not call get
    res = mResourceInMemoryCache->get("B", *mServer, mMemoryManager->getVolatileAddress(), 1024);
    EXPECT_TRUE(res);

    // Call get {4th}
    res = mResourceInMemoryCache->get("C", *mServer, mMemoryManager->getVolatileAddress(), 128);
    EXPECT_TRUE(res);

    // Not call get
    res = mResourceInMemoryCache->get("A", *mServer, mMemoryManager->getVolatileAddress(), 512);
    EXPECT_TRUE(res);

    // Not call get
    res = mResourceInMemoryCache->get("B", *mServer, mMemoryManager->getVolatileAddress(), 1024);
    EXPECT_TRUE(res);

    // Not call get
    res = mResourceInMemoryCache->get("C", *mServer, mMemoryManager->getVolatileAddress(), 128);
    EXPECT_TRUE(res);

    // Call get {5th}
    res = mResourceInMemoryCache->get("D", *mServer, mMemoryManager->getVolatileAddress(), 768);
    EXPECT_TRUE(res);

    // Not call get
    res = mResourceInMemoryCache->get("C", *mServer, mMemoryManager->getVolatileAddress(), 128);
    EXPECT_TRUE(res);

    // Call get {6th}
    res = mResourceInMemoryCache->get("E", *mServer, mMemoryManager->getVolatileAddress(), 1024);
    EXPECT_TRUE(res);

    // Not call get
    res = mResourceInMemoryCache->get("D", *mServer, mMemoryManager->getVolatileAddress(), 768);
    EXPECT_TRUE(res);
}

}  // namespace test
}  // namespace gapir
