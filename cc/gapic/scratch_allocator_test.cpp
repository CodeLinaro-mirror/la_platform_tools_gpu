/*
 * Copyright (C) 2015 The Android Open Source Project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#include "scratch_allocator.h"

#include <gmock/gmock.h>
#include <gtest/gtest.h>

using ::testing::ElementsAreArray;

namespace gapic {

namespace test {

namespace {

enum Type { TypeA, TypeB, TypeC };
struct I { virtual Type getType() = 0; };
struct A : public I { Type getType() override { return TypeA; } };
struct B : public I { Type getType() override { return TypeB; } };
struct C : public I { Type getType() override { return TypeC; } };

}  // anonymous namespace

class ScratchAllocatorTest : public ::testing::Test {};

TEST_F(ScratchAllocatorTest, BaseAddresses) {
    ScratchAllocator sa(0x1000);
    int* base = sa.create<int>(1);
    sa.reset();
    int* a = sa.create<int>(1);
    int* b = sa.create<int>(1);
    int* c = sa.create<int>(1);
    EXPECT_EQ(base+0, a);
    EXPECT_EQ(base+1, b);
    EXPECT_EQ(base+2, c);
}

TEST_F(ScratchAllocatorTest, Vectors) {
    ScratchAllocator sa(0x1000);
    auto v = sa.vector<I*>(3);

    for (auto e : v) {
        FAIL() << "Empty vector should not interate.";
    }

    v.append(sa.create<A>());
    v.append(sa.create<B>());
    v.append(sa.create<C>());

    EXPECT_EQ(TypeA, v[0]->getType());
    EXPECT_EQ(TypeB, v[1]->getType());
    EXPECT_EQ(TypeC, v[2]->getType());

    Type expected[3] = {TypeA, TypeB, TypeC};
    int i = 0;
    for (auto c : v) {
        EXPECT_EQ(expected[i++], c->getType());
    }
}

} // namespace test
} // namespace gapic