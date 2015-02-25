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

#ifndef ANDROID_CAZE_TARGET_H
#define ANDROID_CAZE_TARGET_H

// These values are used for target differentiation on session initialization.
// Do not change them without also updating the server-side replay code.
#define CAZE_OS_LINUX   1
#define CAZE_OS_OSX     2
#define CAZE_OS_WINDOWS 3
#define CAZE_OS_ANDROID 4

#if defined(TARGET_OS_LINUX)
#   define TARGET_OS CAZE_OS_LINUX
#   define STDCALL
#   define PATH_DELIMITER '/'
#   define PATH_DELIMITER_STR "/"
#endif

#if defined(TARGET_OS_OSX)
#   define TARGET_OS CAZE_OS_OSX
#   define STDCALL
#   define PATH_DELIMITER '/'
#   define PATH_DELIMITER_STR "/"
#endif

#if defined(TARGET_OS_ANDROID)
#   define TARGET_OS CAZE_OS_ANDROID
#   define STDCALL
#   define PATH_DELIMITER '/'
#   define PATH_DELIMITER_STR "/"
#endif

#if defined(TARGET_OS_WINDOWS)
#   define TARGET_OS CAZE_OS_WINDOWS
#   define STDCALL __stdcall
#   define PATH_DELIMITER '\\'
#   define PATH_DELIMITER_STR "\\"
#endif

#ifndef TARGET_OS
#   error "OS not defined correctly."
#   error "Exactly one of the following macro have to be defined:" \
           "TARGET_OS_LINUX, TARGET_OS_OSX, TARGET_OS_WINDOWS, TARGET_OS_ANDROID"
#endif

#ifdef _MSC_VER // MSVC
#   define snprintf _snprintf
#endif // _MSC_VER

#endif  // ANDROID_CAZE_TARGET_H
