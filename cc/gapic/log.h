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

#ifndef GAPIC_LOG_H
#define GAPIC_LOG_H

// General logging functions for the replay system. All logging should be done through these macros.
// On android the log is written to the general log output with tag name "Caze". On PCs it is
// written to "logs/replay.log" if possible or to the stderr otherwise.
//
// The replay system supports the following log levels with the specified meanings:
// * LOG_LEVEL_FATAL:   Fatal error, no recovery is possible, the system will die immediately
//                      (always should be enabled)
// * LOG_LEVEL_WARNING: Issue with the current replay. It have to be terminated (it is the callers
//                      task) but the replay system can accept further replay requests
// * LOG_LEVEL_INFO:    General purpose logging, the amount of them shouldn't effect the performance
//                      of the application. This is the default log level.
// * LOG_LEVEL_DEBUG    Verbose logs used only for debugging. They can have significant effect on
//                      the performance of the application

#include "target.h"

// Levels of logging from least verbose to most verbose.
#define LOG_LEVEL_FATAL   0
#define LOG_LEVEL_WARNING 1
#define LOG_LEVEL_INFO    2
#define LOG_LEVEL_DEBUG   3

// If no log level specified then use the default one.
#ifndef LOG_LEVEL
#   define LOG_LEVEL LOG_LEVEL_WARNING
#endif

#define GAPID_STR(S) GAPID_STR2(S)
#define GAPID_STR2(S) #S

#if TARGET_OS == GAPID_OS_ANDROID

#include <android/log.h>

#define GAPID_LOGGER_INIT(path)

#define GAPID_FATAL_IMPL(...) \
    __android_log_assert(nullptr, "GAPID", \
                         __FILE__ ":" GAPID_STR(__LINE__) ": " __VA_ARGS__);

#define GAPID_WARNING_IMPL(...) \
    __android_log_print(ANDROID_LOG_WARN, "GAPID", \
                        __FILE__ ":" GAPID_STR(__LINE__) ": " __VA_ARGS__);

#define GAPID_INFO_IMPL(...) \
    __android_log_print(ANDROID_LOG_INFO, "GAPID", \
                        __FILE__ ":" GAPID_STR(__LINE__) ": " __VA_ARGS__);

#define GAPID_DEBUG_IMPL(...) \
    __android_log_print(ANDROID_LOG_DEBUG, "GAPID", \
                        __FILE__ ":" GAPID_STR(__LINE__) ": " __VA_ARGS__);

#else  // TARGET_OS == GAPID_OS_ANDROID

#include <stdio.h>

namespace gapic {

// Singleton logger implementation for PCs to write formatted log messages.
class Logger {
public:
    // Initializes the logger to write to the log file at path.
    static void init(const char* path);

    // Write a log message to the log output with the specific log level. The location should
    // contain the place where the log is written from and the format is a standard C format string
    // If a message is logged with level LOG_LEVEL_FATAL, the program will terminate after the
    // message is printed.
    // Log messages take the form:
    // <Time stamp> #GAPID <log level>: <location> -> <message>
    static void log(unsigned level, const char* location, const char* format, ...);

private:
    // The single logger instance
    static Logger instance;

    Logger();
    ~Logger();
    void logImpl(unsigned level, const char* location, const char* format, va_list args);

    FILE* mFile;
};

}  // namespace gapic

#define GAPID_LOGGER_INIT(path) gapic::Logger::init(path)

#define GAPID_FATAL_IMPL(...) \
        ::gapic::Logger::log(LOG_LEVEL_FATAL, \
                             __FILE__ ":" GAPID_STR(__LINE__), __VA_ARGS__)

#define GAPID_WARNING_IMPL(...) \
        ::gapic::Logger::log(LOG_LEVEL_WARNING, \
                             __FILE__ ":" GAPID_STR(__LINE__), __VA_ARGS__)

#define GAPID_INFO_IMPL(...) \
        ::gapic::Logger::log(LOG_LEVEL_INFO, \
                             __FILE__ ":" GAPID_STR(__LINE__), __VA_ARGS__)

#define GAPID_DEBUG_IMPL(...) \
        ::gapic::Logger::log(LOG_LEVEL_DEBUG, \
                             __FILE__ ":" GAPID_STR(__LINE__), __VA_ARGS__)

#endif  // TARGET_OS == GAPID_OS_ANDROID

// Define the required log macros based on the specified log level

#define GAPID_FATAL(...) GAPID_FATAL_IMPL(__VA_ARGS__)

#if LOG_LEVEL >= LOG_LEVEL_WARNING
#   define GAPID_WARNING(...) GAPID_WARNING_IMPL(__VA_ARGS__)
#else
#   define GAPID_WARNING(...)
#endif

#if LOG_LEVEL >= LOG_LEVEL_INFO
#   define GAPID_INFO(...) GAPID_INFO_IMPL(__VA_ARGS__)
#else
#   define GAPID_INFO(...)
#endif

#if LOG_LEVEL >= LOG_LEVEL_DEBUG
#   define GAPID_DEBUG(...) GAPID_DEBUG_IMPL(__VA_ARGS__)
#else
#   define GAPID_DEBUG(...)
#endif


#endif  // GAPIC_LOG_H
