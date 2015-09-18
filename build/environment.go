// Copyright (C) 2015 The Android Open Source Project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package build

import "android.googlesource.com/platform/tools/gpu/log"

// Environment holds the global environment configuration for the build.
type Environment struct {
	Intermediates File       // The intermediates directory
	Roots         RootList   // List of root directories.
	Keystore      File       // The path to the keystore used to sign APKs.
	Storepass     string     // The password to the keystore.
	Keypass       string     // The password to the key in the keystore.
	Keyalias      string     // The alias of the key used to sign APKs.
	Logger        log.Logger // The logger to emit log messages to.
	ForceBuild    bool       // If true, all build steps will be forced.
	Verbose       int        // The verbosity level.
}

// Root describes a named root directory.
type Root struct {
	Name string // Name of this root.
	Path File   // Root path
}

// RootList is a list of roots.
type RootList []Root

// Find returns the first Root that contains the file f, or nil if no roots
// contain the file.
func (l RootList) Find(f File) *Root {
	for _, r := range l {
		if r.Path.Contains(f) {
			return &r
		}
	}
	return nil
}
