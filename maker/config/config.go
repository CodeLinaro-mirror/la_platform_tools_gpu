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

package config

import (
	"log"
	"os"
	"path/filepath"

	"android.googlesource.com/platform/tools/gpu/maker"
)

var (
	// Verbose enables increased logging output.
	Verbose int
	// StopOnError makes the system quit faster once an error has been found.
	StopOnError bool

	// Paths holds the set of path roots for the build.
	Paths struct {
		// The root path of the build
		Root string
		// The dependancy cache directory
		Deps string
		// The application binary directory.
		Bin string
	}

	//GoPath is the GOPATH environment setting
	GoPath []string

	// EnvVars holds the environment overrides used when spawning external commands.
	EnvVars = map[string][]string{}

	TargetOS string = HostOS
)

func init() {
	GoPath = filepath.SplitList(os.Getenv("GOPATH"))
	if len(GoPath) == 0 {
		log.Fatalf("GOPATH %q not valid", os.Getenv("GOPATH"))
	}
	for i := range GoPath {
		GoPath[i] = maker.CommonPath(GoPath[i])
	}
	root := GoPath[0]
	Paths.Root = root
	Paths.Deps = maker.Path(root, "deps")
	Paths.Bin = maker.Path(root, "bin")
	EnvVars["PATH"] = []string{Paths.Bin}
}
