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

#include <gapic/log.h>
#include <gapic/target.h>

#include <dlfcn.h>

#if defined(__LP64__)
#define SYSTEM_LIB_PATH "/system/lib64/"
#else
#define SYSTEM_LIB_PATH "/system/lib/"
#endif

namespace gapic {

namespace {

// Utility class for retrieving function pointers from dynamic libraries.
class DlLoader {
public:
  // Loads the specified dynamic library. If the library cannot be loaded then this is a fatal
  // error.
  inline DlLoader(const char* name) {
    mLibrary = dlopen(name, RTLD_NOW | RTLD_LOCAL);
    if (mLibrary == nullptr) {
      GAPID_FATAL("Can't load library %s: %s", name, dlerror());
    }
  }

  // Unloads the library loaded in the constructor.
  inline ~DlLoader() {
    dlclose(mLibrary);
  }

  // Looks up the function with the specified name from the library. Returns nullptr if the function
  // is not found.
  inline void* lookup(const char* name) {
    return dlsym(mLibrary, name);
  }

private:
  void* mLibrary;
};

} // anonymous namespace

void* GetGfxProcAddress(const char* name) {
    static DlLoader dlLoader(SYSTEM_LIB_PATH "libEGL.so");

    typedef void*(*getProcAddressType)(const char*);
    getProcAddressType getProcAddress =
            reinterpret_cast<getProcAddressType>(dlLoader.lookup("eglGetProcAddress"));
    if (getProcAddress == nullptr) {
      GAPID_FATAL("Can't find eglGetProcAddress(): %s", dlerror());
    }

    return getProcAddress(name);
}

}  // namespace gapic

