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

#ifndef GAPII_SPY_BASE_H
#define GAPII_SPY_BASE_H

#include "slice.h"

#include <gapic/encoder.h>
#include <gapic/interval_list.h>
#include <gapic/mutex.h>

#include <gapic/coder/memory.h>
#include <gapic/coder/atom.h>

#include <stdint.h>

#include <memory>
#include <string>
#include <unordered_set>
#include <vector>

namespace gapii {

class SpyBase {
public:
    void init(std::shared_ptr<gapic::Encoder> encoder);

    // lock must be called before invoking any command.
    inline void lock();

    // unlock must be called after invoking any command.
    inline void unlock();

protected:
    typedef gapic::coder::atom::Observation Observation;

    typedef std::unordered_set<gapic::Id> IdSet;
    typedef std::shared_ptr<gapic::Encoder> EncoderSPtr;

    // read is called to make a read memory observation of size bytes, starting at base.
    void read(const void* base, uint64_t size);

    // write is called to make a write memory observation of size bytes, starting at base.
    void write(const void* base, uint64_t size);

    // observe observes all the pending memory observations, returning the list of observations made.
    // The list of pending memory observations is cleared on returning.
    void observe(gapic::Array<Observation>& observations);

    // read observes the memory for the given slice as a read operation.
    template <typename T>
    inline void read(const Slice<T>& slice);

    // read reads and returns the i'th element from the slice src.
    template <typename T>
    inline T read(const Slice<T>& src, uint64_t i);

    // write observes the memory for the given slice as a write operation.
    template <typename T>
    inline void write(const Slice<T>& slice);

    // writes a value to i'th element in the slice dst.
    template <typename T>
    inline void write(const Slice<T>& dst, uint64_t i, const T& value);

    // copy copies N elements from src to dst, where N is the smaller of src.count() and
    // dst.count().
    // copy observes the sub-slice of src as a read operation.
    // The sub-slice of dst is returned so that the write observation can be made after the call to
    // the imported function.
    template <typename T>
    inline Slice<T> copy(const Slice<T>& dst, const Slice<T>& src);

    // clone observes src as a read operation and returns a copy of src in a new Pool.
    template<typename T>
    inline Slice<T> clone(const Slice<T>& src);

    // make constructs and returns a Slice backed by a new pool.
    template<typename T>
    inline Slice<T> make(uint64_t count) const;

    // slice returns a slice wrapping the application-pool pointer src, starting at elements s
    // ending at one element before e.
    template<typename T>
    inline Slice<T> slice(T* src, uint64_t s, uint64_t e) const;

    // slice returns a slice wrapping the application-pool pointer src, starting at s bytes
    // from src and ending at one byte before e.
    inline Slice<uint8_t> slice(void* src, uint64_t s, uint64_t e) const;

    // slice returns a Slice<char>, backed by a new pool, holding a copy of the string src.
    // src is observed as a read operation.
    inline Slice<char> slice(const std::string& src) const;

    // slice returns a sub-slice of src, starting at elements s and ending at one element before e.
    template<typename T>
    inline Slice<T> slice(const Slice<T>& src, uint64_t s, uint64_t e) const;

    // string returns a std::string from the null-terminated string str.
    // str is observed as a read operation.
    inline std::string string(const char* str);

    // string returns a std::string from the Slice<char> slice.
    // slice is observed as a read operation.
    inline std::string string(const Slice<char>& slice);

    // The output stream encoder.
    EncoderSPtr mEncoder;

private:
    // writes a value to i'th element in the slice dst.
    // To disambiguate the overloads of write_, void* is used as a dummy last parameter so that the
    // T[N] overload is preferred (but not an option for non-array T types).
    template <typename T>
    inline void write_(const Slice<T>& dst, uint64_t i, const T& value, void*);

    // writes an array-value to i'th element in the slice dst.
    // To disambiguate the overloads of write_, int is used as a dummy last parameter so that the
    // T[N] overload is preferred (but not an option for non-array T types).
    template <typename T, size_t N>
    inline void write_(const Slice<T[N]>& dst, uint64_t i, const T(&value)[N], int);

    // The list of pending reads or writes observations that are yet to be made.
    gapic::IntervalList<uintptr_t> mPendingObservations;

    // The list of observations that have already been encoded.
    IdSet mResources;

    // The mutex that should be locked for the duration of each of the intercepted commands.
    gapic::Mutex mMutex;
};

inline void SpyBase::lock() {
    mMutex.lock();
}

inline void SpyBase::unlock() {
    mMutex.unlock();
}

template <typename T>
inline void SpyBase::read(const Slice<T>& slice) {
    if (slice.isApplicationPool()) {
        read(slice.begin(), slice.count() * sizeof(T));
    }
}

template<typename T>
inline T SpyBase::read(const Slice<T>& src, uint64_t index) {
    T& elem = src[index];
    if (src.isApplicationPool()) {
        read(&elem, sizeof(T));
    }
    return elem;
}

template <typename T>
inline void SpyBase::write(const Slice<T>& slice) {
    if (slice.isApplicationPool()) {
        return write(slice.begin(), slice.count() * sizeof(T));
    }
}

template<typename T>
inline void SpyBase::write(const Slice<T>& dst, uint64_t index, const T& value) {
    write_(dst, index, value, 0);
}

template<typename T>
inline void SpyBase::write_(const Slice<T>& dst, uint64_t index, const T& value, void*) {
    if (!dst.isApplicationPool()) { // The spy must not mutate data in the application pool.
        dst[index] = value;
    } else {
        write(&dst[index], sizeof(T));
    }
}

template <typename T, size_t N>
inline void SpyBase::write_(const Slice<T[N]>& dst, uint64_t index, const T(&value)[N], int) {
    if (!dst.isApplicationPool()) { // The spy must not mutate data in the application pool.
        for (size_t i = 0; i < N; i++) {
            dst[index][i] = value[i];
        }
    } else {
        write(&dst[index], sizeof(T[N]));
    }
}

template <typename T>
inline Slice<T> SpyBase::copy(const Slice<T>& dst, const Slice<T>& src) {
    read(src);
    if (!dst.isApplicationPool()) { // The spy must not mutate data in the application pool.
        uint64_t c = (src.count() < dst.count()) ? src.count() : dst.count();
        for (uint64_t i = 0; i < c; i++) {
          dst[i] = src[i];
        }
    }
    return dst;
}

template<typename T>
inline Slice<T> SpyBase::clone(const Slice<T>& src) {
    Slice<T> dst = make<T>(src.count());
    copy(dst, src);
    return dst;
}

template<typename T>
inline Slice<T> SpyBase::make(uint64_t count) const {
    auto pool = Pool::create(count * sizeof(T));
    return Slice<T>(reinterpret_cast<T*>(pool->base()), count, pool);
}

template<typename T>
inline Slice<T> SpyBase::slice(T* src, uint64_t s, uint64_t e) const {
    // TODO: Find the pool containing src
    return Slice<T>(src+s, e-s, std::shared_ptr<Pool>());
}

inline Slice<uint8_t> SpyBase::slice(void* src, uint64_t s, uint64_t e) const {
    return slice(reinterpret_cast<uint8_t*>(src), s, e);
}

inline Slice<char> SpyBase::slice(const std::string& src) const {
    Slice<char> dst = make<char>(src.length());
    for (uint64_t i = 0; i < src.length(); i++) {
        dst[i] = src[i];
    }
    return dst;
}

template<typename T>
inline Slice<T> SpyBase::slice(const Slice<T>& src, uint64_t s, uint64_t e) const {
    return src(s, e);
}

inline std::string SpyBase::string(const char* str) {
    for (uint64_t i = 0; ; i++) {
        if (str[i] == 0) {
            read(str, i + 1);
            return std::string(str, str + i);
        }
    }
}

inline std::string SpyBase::string(const Slice<char>& slice) {
    read(slice);
    return std::string(slice.begin(), slice.end());
}

}  // namespace gapii

#endif // GAPII_SPY_BASE_H
