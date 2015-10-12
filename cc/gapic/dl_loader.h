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

#ifndef GAPIC_DL_LOADER_H
#define GAPIC_DL_LOADER_H

namespace gapic {

// Utility class for retrieving function pointers from dynamic libraries.
class DlLoader {
public:
    // Loader is a function that can be used to load libraries.
    typedef void* (Loader)(const char* name);

    // Resolver is a function that can be used to find functions in libraries.
    typedef void* (Resolver)(void* handle, const char* name);

    // setCustomLoader replaces the default library loader with the specified loader for all
    // instances of DlLoader.
    static void setCustomLoader(Loader* loader);

    // setCustomResolver replaces the default function resolver with the specified resolver
    // for all instances of DlLoader.
    static void setCustomResolver(Resolver* resolver);

    // Loads the specified dynamic library.
    // If the library cannot be loaded then this is a fatal error.
    // For *nix systems, a nullptr can be used to search the application's functions.
    DlLoader(const char* name);

    // Unloads the library loaded in the constructor.
    ~DlLoader();

    // Looks up the function with the specified name from the library.
    // Returns nullptr if the function is not found.
    void* lookup(const char* name);

private:
    DlLoader() =default;
    DlLoader(const DlLoader&) =delete;
    DlLoader& operator=(const DlLoader&) =delete;

    void* mLibrary;

    static Loader* sLoader;
    static Resolver* sResolver;
};

}  // namespace gapic

#endif  // GAPIC_DL_LOADER_H
