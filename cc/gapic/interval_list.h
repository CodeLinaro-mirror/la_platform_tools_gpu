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

#ifndef GAPIC_INTERVAL_LIST_H
#define GAPIC_INTERVAL_LIST_H

#include <algorithm>
#include <vector>

namespace gapic {

// Interval represents a single interval range of type T.
template <typename T>
struct Interval {
    T start; // The index of the first item in the interval range.
    T end;   // The index of the one-past the last item in the interval range.
};

// Interval equality operator.
template <typename T>
inline bool operator == (const Interval<T>& lhs, const Interval<T>& rhs) {
    return lhs.start == rhs.start && lhs.end == rhs.end;
}

// IntervalList holds a ascendingly-sorted list of Intervals.
// IntervalList supports for-range looping.
template <typename T>
class IntervalList {
public:
    typedef Interval<T>         value_type;
    typedef const Interval<T>*  const_iterator;

    // merge adds the interval i to this list, merging any overlapping intervals.
    inline void merge(const Interval<T>& i);

    // clear removes all intervals from the list.
    inline void clear();

    // count returns the number of intervals in the list.
    inline uint32_t count() const;

    // begin() returns the pointer to the first interval in the list.
    inline const Interval<T>* begin() const;

    // end() returns the pointer to one-past the last interval in the list.
    inline const Interval<T>* end() const;

protected:
    // mergeStart returns the index of the first interval to use in a merge.
    inline size_t mergeStart(const Interval<T>& i);

    // mergeEnd returns the index of one-past the last interval to use in a merge.
    inline size_t mergeEnd(const Interval<T>& i);

    std::vector< Interval<T> > mIntervals;
};

template<typename T>
inline void IntervalList<T>::merge(const Interval<T>& i) {
    if (count() == 0) {
        mIntervals.push_back(i);
    } else {
        auto start = mIntervals.begin() + mergeStart(i);
        auto end   = mIntervals.begin() + mergeEnd(i);
        if (start <= end) {
            auto low = std::min(start->start, i.start);
            auto high = std::max(end->end, i.end);
            mIntervals.erase(start, end);
            *start = Interval<T>{low, high};
        } else {
            mIntervals.insert(start, i);
        }
    }
}

template<typename T>
inline void IntervalList<T>::clear() {
    mIntervals.clear();
}

template<typename T>
inline uint32_t IntervalList<T>::count() const {
    return mIntervals.size();
}

template<typename T>
inline const Interval<T>* IntervalList<T>::begin() const {
    if (count() > 0) {
        return &mIntervals[0];
    } else {
        return nullptr;
    }
}

template<typename T>
inline const Interval<T>* IntervalList<T>::end() const {
    size_t c = mIntervals.size();
    if (c > 0) {
        return &mIntervals[c];
    } else {
        return nullptr;
    }
}

template<typename T>
inline size_t IntervalList<T>::mergeStart(const Interval<T>& i) {
    size_t l = 0;
    size_t h = mIntervals.size();
    while (l != h) {
        size_t m = (l + h) / 2;
        if (mIntervals[m].end >= i.start) {
            h = m;
        } else {
            l = m + 1;
        }
    }
    return l;
}

template<typename T>
inline size_t IntervalList<T>::mergeEnd(const Interval<T>& i) {
    size_t l = -1;
    size_t h = mIntervals.size() - 1;
    while (l != h) {
        size_t m = (l + h + 1) / 2;
        if (mIntervals[m].start <= i.end) {
            l = m;
        } else {
            h = m - 1;
        }
    }
    return l;
}

} // namespace gapic

#endif // GAPIC_INTERVAL_LIST_H
