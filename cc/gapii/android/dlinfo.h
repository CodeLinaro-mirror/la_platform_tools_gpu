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

#ifndef GAPII_ANDROID_DL_INFO_H
#define GAPII_ANDROID_DL_INFO_H

namespace gapii {

// DlInfo is a helper class that wraps dladdr() and Dl_info for more convenient usage.
class DlInfo {
public:
  // self assigns to out the DlInfo describing the GAPII library.
  // self returns null on success or an error message.
  static const char* self(DlInfo& out);

  // find assigns to out the DlInfo of the library that holds the specified address.
  // find returns null on success or an error message.
  static const char* find(const void* addr, DlInfo& out);

  const char* mPath;    // Path to the shared object.
  const void* mAddress; // Address at which shared object.
};

}  // namespace gapii

#endif  // GAPII_ANDROID_DL_INFO_H
