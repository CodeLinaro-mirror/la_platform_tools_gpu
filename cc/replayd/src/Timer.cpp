/*
 * Copyright 2015, The Android Open Source Project
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

#include "Log.h"
#include "Target.h"
#include "Timer.h"

#include <ctime>

#if TARGET_OS == CAZE_OS_WINDOWS
#include "windows.h"
#else
#include <errno.h>
#include <sys/time.h>
#endif

namespace android {
namespace caze {

namespace {

const uint64_t SEC_TO_NANO   = 1000000000;
const uint64_t MICRO_TO_NANO = 1000;

// Returns a monotonic clock reading in a platform-specific unit.
// Use platformDurationToNanosecods() to convert the difference in values returned from two calls to
// platformGetTime() into nanoseconds.
inline uint64_t platformGetTime() {
#if TARGET_OS == CAZE_OS_OSX
    timeval tv = {0, 0};
    if (gettimeofday(&tv, NULL) != 0) {
        CAZE_FATAL("Unable to start timer. Error: %d", errno);
    }
    return uint64_t(tv.tv_usec) * MICRO_TO_NANO +
            uint64_t(tv.tv_sec) * SEC_TO_NANO;
#elif TARGET_OS == CAZE_OS_WINDOWS
    LARGE_INTEGER i;
    if (!QueryPerformanceCounter(&i)) {
        CAZE_FATAL("Unable to start timer. Error: %d", GetLastError());
    }
    return i.QuadPart;
#else
    timespec ts = {0, 0};
    if (clock_gettime(CLOCK_MONOTONIC, &ts) != 0) {
        CAZE_FATAL("Unable to start timer. Error: %d", errno);
    }
    return uint64_t(ts.tv_nsec) + uint64_t(ts.tv_sec) * SEC_TO_NANO;
#endif
}

// See platformGetTime().
inline uint64_t platformDurationToNanosecods(uint64_t duration) {
#if TARGET_OS == CAZE_OS_WINDOWS
    static LARGE_INTEGER sFreq = 0;
    if (sFreq.QuadPart == 0) {
        QueryPerformanceFrequency(&sFreq);
        if (sFreq.QuadPart == 0) {
            CAZE_FATAL("Unable to query performance frequency. Error: %d", GetLastError());
        }
    }
    return (duration * SEC_TO_NANO) / sFreq.QuadPart;
#else
    return duration;
#endif
}

} // end of anonymous namespace

void Timer::Start() {
    mStartTime = platformGetTime();
}

uint64_t Timer::Stop() {
    uint64_t endTime = platformGetTime();
    return platformDurationToNanosecods(endTime - mStartTime);
}

} // end of namespace caze
} // end of namespace android
