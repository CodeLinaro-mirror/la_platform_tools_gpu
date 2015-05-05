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

package maker

import "path/filepath"

var (
	// Config holds the current configuration of the maker system.
	Config struct {
		// Verbose enables increased logging output.
		Verbose int
		// TargetArchitecture is the architecture to build binaries for.
		TargetArchitecture string
		// TargetOS is the OS to build for.
		TargetOS string
		// RootPath is the root directory to work in.
		RootPath string
		// DisableParallel turns of all parallel build support.
		DisableParallel bool
	}
)

// DepsPath joins the supplied path to the dependancy cache directory.
func DepsPath(path string) string {
	return filepath.Join(Config.RootPath, "deps", path)
}

// DataPath joins the supplied path to the application data directory.
func DataPath(path string) string {
	return filepath.Join(Config.RootPath, "data", path)
}
