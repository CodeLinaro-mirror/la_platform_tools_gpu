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

#ifndef GAPII_STATE_BASE_H
#define GAPII_STATE_BASE_H

#include "slice.h"

#include <functional>
#include <stdint.h>

namespace gapii {

class StateBase {
public:
    typedef std::function<void(const void* ptr, uint64_t offset, uint64_t size)> ObserveFunc;

    inline StateBase(ObserveFunc observe);

protected:
    template <typename T>
    inline void read(const Slice<T>& slice) const;

    template <typename T>
    inline void write(const Slice<T>& slice) const;

    template <typename T>
    inline void copy(const Slice<T>& dst, const Slice<T>& src) const;

    template<typename T>
    inline Slice<T> make(uint64_t count) const;

    template<typename T>
    inline Slice<T> clone(const Slice<T>& src) const;

    template<typename T>
    inline Slice<T> slice(T* src, uint64_t s, uint64_t e) const;

    inline Slice<uint8_t> slice(void* src, uint64_t s, uint64_t e) const;

    inline Slice<char> slice(const std::string& src) const;

    template<typename T>
    inline Slice<T> slice(const Slice<T>& src, uint64_t s, uint64_t e) const;

    inline std::string string(const Slice<char>& slice) const;

    ObserveFunc mObserve;
};

inline StateBase::StateBase(ObserveFunc observe)
    : mObserve(observe) {}

template <typename T>
inline void StateBase::read(const Slice<T>& slice) const {
    return mObserve(slice.begin(), 0, slice.count() * sizeof(T));
}

template <typename T>
inline void StateBase::write(const Slice<T>& slice) const {
    return mObserve(slice.begin(), 0, slice.count() * sizeof(T));
}

template <typename T>
inline void StateBase::copy(const Slice<T>& dst, const Slice<T>& src) const {
    read(src);
    uint64_t c = (src.count() < dst.count()) ? src.count() : dst.count();
    for (uint64_t i = 0; i < c; i++) {
      dst[i] = src[i];
    }
    write(dst);
}

template<typename T>
inline Slice<T> StateBase::make(uint64_t count) const {
    auto pool = Pool::create(count * sizeof(T));
    return Slice<T>(reinterpret_cast<T*>(pool->base()), count, pool);
}

template<typename T>
inline Slice<T> StateBase::clone(const Slice<T>& src) const {
    Slice<T> dst = make<T>(src.count());
    copy(dst, src);
    return dst;
}

template<typename T>
inline Slice<T> StateBase::slice(T* src, uint64_t s, uint64_t e) const {
    // TODO: Find the pool containing src
    return Slice<T>(src+s, e-s, std::shared_ptr<Pool>());
}

inline Slice<uint8_t> StateBase::slice(void* src, uint64_t s, uint64_t e) const {
    return slice(reinterpret_cast<uint8_t*>(src), s, e);
}

inline Slice<char> StateBase::slice(const std::string& src) const {
    Slice<char> dst = make<char>(src.length());
    for (uint64_t i = 0; i < src.length(); i++) {
        dst[i] = src[i];
    }
    return dst;
}

template<typename T>
inline Slice<T> StateBase::slice(const Slice<T>& src, uint64_t s, uint64_t e) const {
    return src(s, e);
}

inline std::string StateBase::string(const Slice<char>& slice) const {
    return std::string(slice.begin(), slice.end());
}

}  // namespace gapii

#endif // GAPII_STATE_BASE_H
