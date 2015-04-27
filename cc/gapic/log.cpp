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

#include "log.h"

#include <gapic/target.h>

#if TARGET_OS != GAPID_OS_ANDROID

#include <errno.h>
#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include <chrono>
#include <ctime> // Required for MSVC.

namespace gapic {

Logger Logger::instance;

void Logger::init(const char* path) {
    if (FILE* f = fopen(path, "w")) {
        GAPID_INFO("Logging to %s\n", path);
        instance.mFile = f;
    } else {
        GAPID_WARNING("Can't open file for logging (%s): %s\n", path, strerror(errno));
    }
}

void Logger::log(unsigned level, const char* location, const char* format, ...) {
    va_list args;
    va_start(args, format);
    instance.logImpl(level, location, format, args);
    va_end(args);

    if (level == LOG_LEVEL_FATAL) {
        exit(EXIT_FAILURE);
    }
}

Logger::Logger() : mFile(stdout) {}

Logger::~Logger() {
    fclose(mFile);
}

void Logger::logImpl(unsigned level, const char* location, const char* format, va_list args) {
    // Get the current time with milliseconds precision
    auto t = std::chrono::system_clock::now();
    std::time_t now = std::chrono::system_clock::to_time_t(t);
    std::tm* loc = std::localtime(&now);
    auto ms = std::chrono::duration_cast<std::chrono::milliseconds>(t.time_since_epoch());

    // Print out the common part of the log messages
    fprintf(mFile, "%02d:%02d:%02d.%03d #Caze %c: %s -> ", loc->tm_hour, loc->tm_min, loc->tm_sec,
            static_cast<int>(ms.count() % 1000), "FWID"[level], location);

    // Print out the actual log message
    vfprintf(mFile, format, args);

    // Flush the log to ensure that every message is written out even if the application crashes
    fflush(mFile);
}

}  // namespace gapic

#endif  // TARGET_OS != GAPID_OS_ANDROID
