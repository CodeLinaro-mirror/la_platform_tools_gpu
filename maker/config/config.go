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

	"android.googlesource.com/platform/tools/gpu/maker/graph"
)

var (
	// Verbose enables increased logging output.
	Verbose int
	// StopOnError makes the system quit faster once an error has been found.
	StopOnError bool

	// Paths holds the set of path roots for the build.
	Paths struct {
		// The root path of the repo workspace
		Repo *graph.Path
		// The default GOPATH entry
		GoBase *graph.Path
		// The dependancy cache directory
		Deps *graph.Path
		// The application binary directory.
		Bin *graph.Path
	}

	//GoPath is the GOPATH environment setting
	GoPath graph.Set

	// EnvVars holds the environment overrides used when spawning external commands.
	EnvVars = map[string][]string{}

	TargetOS = HostOS
)

func init() {
	gopath := filepath.SplitList(os.Getenv("GOPATH"))
	if len(gopath) == 0 {
		log.Fatalf("GOPATH %q not valid", os.Getenv("GOPATH"))
	}
	GoPath = make(graph.Set, len(gopath))
	for i, p := range gopath {
		dir := graph.Dir(p)
		if i == 0 {
			Paths.GoBase = dir
		}
		GoPath[i] = dir
	}
	Paths.Repo = Paths.GoBase.Parent().Parent()
	Paths.Deps = Paths.GoBase.Child("deps")
	Paths.Bin = Paths.GoBase.Child("bin")
	EnvVars["PATH"] = []string{Paths.Bin.Name()}
}
