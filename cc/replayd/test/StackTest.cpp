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

#include "MemoryManager.h"
#include "Stack.h"
#include "TestUtilities.h"

#include <memory>

#include <gtest/gtest.h>

namespace android {
namespace caze {
namespace test {
namespace {

const size_t MEMORY_SIZE = 4096;
const uint32_t STACK_CAPACITY = 128;

class StackTest : public ::testing::Test {
protected:
    virtual void SetUp() {
        mMemoryManager.reset(new caze::MemoryManager({MEMORY_SIZE}));
        mStack.reset(new caze::Stack(STACK_CAPACITY, mMemoryManager.get()));
    }

    std::unique_ptr<caze::MemoryManager> mMemoryManager;
    std::unique_ptr<caze::Stack> mStack;
};

void fillStack(caze::Stack* stack, uint32_t n) {
    for (unsigned i = 0; i < n; ++i) {
        stack->push<uint32_t>(0);
    }
}
}  // end of anonymous namespace

TEST_F(StackTest, IsValid) {
    EXPECT_TRUE(mStack->isValid());
}

TEST_F(StackTest, GetType) {
    mStack->push<uint32_t>(123);
    EXPECT_EQ(caze::BaseType::Uint32, mStack->getTopType());
    EXPECT_TRUE(mStack->isValid());
    mStack->discard(1);
    EXPECT_TRUE(mStack->isValid());

    mStack->push<void*>(nullptr);
    EXPECT_EQ(caze::BaseType::AbsolutePointer, mStack->getTopType());
    EXPECT_TRUE(mStack->isValid());
    mStack->discard(1);
    EXPECT_TRUE(mStack->isValid());
}

TEST_F(StackTest, GetTypeErrorEmptyStack) {
    mStack->getTopType();
    EXPECT_FALSE(mStack->isValid());

    mStack->getTopType();
    EXPECT_FALSE(mStack->isValid());
}

TEST_F(StackTest, PushFrom) {
    uint32_t x = 123456789;
    mStack->pushFrom(caze::BaseType::Uint32, &x);
    EXPECT_EQ(123456789, mStack->pop<uint32_t>());
    EXPECT_TRUE(mStack->isValid());
}

TEST_F(StackTest, PushFromErrorStackOverflow) {
    fillStack(mStack.get(), STACK_CAPACITY);
    EXPECT_TRUE(mStack->isValid());

    uint32_t x = 123456789;
    mStack->pushFrom(caze::BaseType::Uint32, &x);
    EXPECT_FALSE(mStack->isValid());

    mStack->pushFrom(caze::BaseType::Uint32, &x);
    EXPECT_FALSE(mStack->isValid());
}

TEST_F(StackTest, PopTo) {
    uint32_t x;
    mStack->push<uint32_t>(123456);
    mStack->popTo(&x);
    EXPECT_EQ(123456, x);
}

TEST_F(StackTest, PopToErrorEmpytStack) {
    uint32_t x;

    mStack->popTo(&x);
    EXPECT_FALSE(mStack->isValid());

    mStack->popTo(&x);
    EXPECT_FALSE(mStack->isValid());
}

TEST_F(StackTest, Discard) {
    mStack->push<uint32_t>(1234);
    mStack->push<uint32_t>(2345);
    mStack->push<uint32_t>(3356);
    EXPECT_TRUE(mStack->isValid());

    mStack->discard(2);
    EXPECT_TRUE(mStack->isValid());
    EXPECT_EQ(1234, mStack->pop<uint32_t>());
}

TEST_F(StackTest, SequentialDiscard) {
    mStack->push<uint32_t>(12);
    mStack->push<uint32_t>(123);
    mStack->push<uint32_t>(234);
    mStack->discard(1);
    mStack->push<uint32_t>(345);
    mStack->discard(1);

    EXPECT_TRUE(mStack->isValid());
    EXPECT_EQ(123, mStack->pop<uint32_t>());
}

TEST_F(StackTest, DiscardErrorStackUnderflow) {
    mStack->push<uint32_t>(1234);
    EXPECT_TRUE(mStack->isValid());

    mStack->discard(2);
    EXPECT_FALSE(mStack->isValid());

    mStack->discard(1);
    EXPECT_FALSE(mStack->isValid());
}

TEST_F(StackTest, Clone) {
    mStack->push<uint32_t>(1234);
    mStack->push<uint32_t>(2345);
    EXPECT_TRUE(mStack->isValid());

    mStack->clone(1);
    EXPECT_TRUE(mStack->isValid());

    EXPECT_EQ(1234, mStack->pop<uint32_t>());
    EXPECT_EQ(2345, mStack->pop<uint32_t>());
    EXPECT_EQ(1234, mStack->pop<uint32_t>());
    EXPECT_TRUE(mStack->isValid());
}

TEST_F(StackTest, CloneErrorNonExistantIndex) {
    mStack->clone(1);
    EXPECT_FALSE(mStack->isValid());

    mStack->clone(1);
    EXPECT_FALSE(mStack->isValid());
}

TEST_F(StackTest, CloneErrorOutsideOfCapacity) {
    mStack->push<uint32_t>(1234);
    mStack->push<uint32_t>(2345);
    EXPECT_TRUE(mStack->isValid());

    mStack->clone(3);
    EXPECT_FALSE(mStack->isValid());
}

TEST_F(StackTest, CloneErrorStackOverflow) {
    fillStack(mStack.get(), STACK_CAPACITY);
    EXPECT_TRUE(mStack->isValid());

    mStack->clone(1);
    EXPECT_FALSE(mStack->isValid());

    mStack->clone(1);
    EXPECT_FALSE(mStack->isValid());
}

TEST_F(StackTest, PushPop) {
    mStack->push<uint32_t>(123);
    EXPECT_TRUE(mStack->isValid());

    ASSERT_EQ(123, mStack->pop<uint32_t>());
    EXPECT_TRUE(mStack->isValid());
}

TEST_F(StackTest, PopVolatilePointer) {
    uint32_t a = 0;
    mStack->pushFrom(caze::BaseType::VolatilePointer, &a);
    EXPECT_TRUE(mStack->isValid());
    EXPECT_EQ(mMemoryManager->volatileToAbsolute(0), mStack->pop<void*>());
}

TEST_F(StackTest, PopConstantPointer) {
    uint32_t a = 0;
    mStack->pushFrom(caze::BaseType::ConstantPointer, &a);
    EXPECT_TRUE(mStack->isValid());
    EXPECT_EQ(mMemoryManager->constantToAbsolute(0), mStack->pop<void*>());
}

TEST_F(StackTest, PopAbsolutePointer) {
    uint32_t a = 0;
    mStack->push<void*>(nullptr);
    EXPECT_TRUE(mStack->isValid());
    EXPECT_EQ(nullptr, mStack->pop<void*>());
}

TEST_F(StackTest, PopErrorStackUnderflow) {
    mStack->pop<uint32_t>();
    EXPECT_FALSE(mStack->isValid());

    mStack->pop<uint32_t>();
    EXPECT_FALSE(mStack->isValid());
}

TEST_F(StackTest, PopErrorTypeMissmatch) {
    mStack->push<uint16_t>(123);
    EXPECT_TRUE(mStack->isValid());

    mStack->pop<uint32_t>();
    EXPECT_FALSE(mStack->isValid());
}

TEST_F(StackTest, PopErrorMismatchingPointerType) {
    mStack->push<uint32_t>(123);
    EXPECT_TRUE(mStack->isValid());

    mStack->pop<uint32_t*>();
    EXPECT_FALSE(mStack->isValid());
}

TEST_F(StackTest, PushErrorOverCapacity) {
    fillStack(mStack.get(), STACK_CAPACITY);
    EXPECT_TRUE(mStack->isValid());

    mStack->push<uint32_t>(1);
    EXPECT_FALSE(mStack->isValid());

    mStack->push<uint32_t>(1);
    EXPECT_FALSE(mStack->isValid());
}

}  // end of namespace test
}  // end of namespace caze
}  // end of namespace android
