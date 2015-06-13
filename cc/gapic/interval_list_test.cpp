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

#include "interval_list.h"

#include <gmock/gmock.h>
#include <gtest/gtest.h>

using ::testing::ElementsAreArray;

namespace gapic {

template <typename T>
::std::ostream& operator<<(::std::ostream& os, const Interval<T>& interval) {
    return os << "[" << interval.start << " - " << interval.end-1 << "]";
}

template <typename T>
::std::ostream& operator<<(::std::ostream& os, const IntervalList<T>& l) {
    os << "IntervalList{";
    for (auto i : l) {
        os << i;
    }
    return os << "}";
}

namespace test {

Interval<int> I(int first, int last) {
    return Interval<int>{first, last + 1};
}


class IntervalListTest : public ::testing::Test, public IntervalList<int> {
public:
    inline IntervalList<int>& L() {
        return *static_cast<IntervalList<int>*>(this);
    }
};

TEST_F(IntervalListTest, Empty) {
    EXPECT_EQ(count(), 0);
    EXPECT_EQ(begin(), end());
}

TEST_F(IntervalListTest, SingleMerge) {
    merge(I(1, 2));
    EXPECT_THAT(L(), ElementsAreArray({I(1, 2)}));
}

TEST_F(IntervalListTest, MergeSparseForward) {
    merge(I(1, 2));
    merge(I(4, 5));
    merge(I(7, 8));
    EXPECT_THAT(L(), ElementsAreArray({I(1, 2), I(4, 5), I(7, 8)}));
}

TEST_F(IntervalListTest, MergeSparseReverse) {
    merge(I(7, 8));
    merge(I(4, 5));
    merge(I(1, 2));
    EXPECT_THAT(L(), ElementsAreArray({I(1, 2), I(4, 5), I(7, 8)}));
}

//   0   1   2   3   4   5   6   7   8   9   A   B   C   D   E
//  -----------------------------------------------------------
//          ╭         ╮             ╭     ╮     ╭     ╮
//          │    0    │             │  1  │     │  2  │
//          ╰         ╯             ╰     ╯     ╰     ╯
//  -----------------------------------------------------------
//  ╭ ╮ ╭ ╮     ╭ ╮     ╭ ╮     ╭     ╮ ╭         ╮ ╭     ╮ ╭ ╮
//  │a│ │b│     │c│     │d│     │  e  │ │    f    │ │  g  │ │h│
//  ╰ ╯ ╰ ╯     ╰ ╯     ╰ ╯     ╰     ╯ ╰         ╯ ╰     ╯ ╰ ╯
//      ╭                 ╮         ╭                 ╮
//      │        i        │         │        j        │
//      ╰                 ╯         ╰                 ╯
//  ╭                                                         ╮
//  │                            k                            │
//  ╰                                                         ╯
auto a = I(0x0, 0x0);
auto b = I(0x1, 0x1);
auto c = I(0x3, 0x3);
auto d = I(0x5, 0x5);
auto e = I(0x7, 0x8);
auto f = I(0x9, 0xb);
auto g = I(0xc, 0xd);
auto h = I(0xe, 0xe);
auto i = I(0x1, 0x5);
auto j = I(0x8, 0xc);
auto k = I(0x0, 0xe);

TEST_F(IntervalListTest, Merge) {
    struct test {
        const char*                name;
        Interval<int>              interval;
        std::vector<Interval<int>> expected;
    };
    for (auto t : {
        test{"a", a, { I(0x0, 0x0), I(0x2, 0x4), I(0x8, 0x9), I(0xb, 0xc) } },
        test{"b", b, { I(0x1, 0x4), I(0x8, 0x9), I(0xb, 0xc) } },
        test{"c", c, { I(0x2, 0x4), I(0x8, 0x9), I(0xb, 0xc) } },
        test{"d", d, { I(0x2, 0x5), I(0x8, 0x9), I(0xb, 0xc) } },
        test{"e", e, { I(0x2, 0x4), I(0x7, 0x9), I(0xb, 0xc) } },
        test{"f", f, { I(0x2, 0x4), I(0x8, 0xc) } },
        test{"g", g, { I(0x2, 0x4), I(0x8, 0x9), I(0xb, 0xd) } },
        test{"h", h, { I(0x2, 0x4), I(0x8, 0x9), I(0xb, 0xc), I(0xe, 0xe) } },
        test{"i", i, { I(0x1, 0x5), I(0x8, 0x9), I(0xb, 0xc) } },
        test{"j", j, { I(0x2, 0x4), I(0x8, 0xc) } },
        test{"k", k, { I(0x0, 0xe) } },
    }) {
        clear();
        merge(I(0x2, 0x4)); // 0
        merge(I(0x8, 0x9)); // 1
        merge(I(0xb, 0xc)); // 2

        merge(t.interval);

        EXPECT_THAT(L(), ElementsAreArray(t.expected));
    }
}

TEST_F(IntervalListTest, MergeStartEnd) {
    merge(I(0x2, 0x4)); // 0
    merge(I(0x8, 0x9)); // 1
    merge(I(0xb, 0xc)); // 2

    struct test {
        const char*   name;
        Interval<int> interval;
        int           start;
        int           end;
    };
    for (auto t : {
        test{"a", a, 0, -1},
        test{"b", b, 0, 0},
        test{"c", c, 0, 0},
        test{"d", d, 0, 0},
        test{"e", e, 1, 1},
        test{"f", f, 1, 2},
        test{"g", g, 2, 2},
        test{"h", h, 3, 2},
        test{"i", i, 0, 0},
        test{"j", j, 1, 2},
        test{"k", k, 0, 2},
    }) {
        size_t s = mergeStart(t.interval);
        size_t e = mergeEnd(t.interval);
        if (t.start != s) {
            ADD_FAILURE() << t.name << ": l.mergeStart(" << t.interval << ") "
                    "returned " << s << ", expected " << t.start;
        }
        if (t.end != e) {
            ADD_FAILURE() << t.name << ": l.mergeEnd(" << t.interval << ") "
                    "returned " << e << ", expected " << t.end;
        }
    }
}

} // namespace test
} // namespace gapic
