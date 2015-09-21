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

package do

import (
	"android.googlesource.com/platform/tools/gpu/maker"
	"android.googlesource.com/platform/tools/gpu/maker/config"
	"android.googlesource.com/platform/tools/gpu/maker/graph"
)

const GoPkgResources = "go_packages"

var (
	goTool graph.Entity
)

func init() {
	goTool = graph.FindTool("go")
}

// GoCommand runs "go" with the specified arguments.
func GoCommand(args ...string) *graph.Step {
	return Exec(goTool, args...).Access(GoPkgResources)
}

// GoInstall builds a new Step that runs "go install" on the supplied module.
// It will return the resulting binary entity.
// The step will depend on the go tool, and will be set to always run if
// depended on.
func GoInstall(root string, relative string) graph.Entity {
	module := maker.Path(root, relative)
	_, name := maker.PathSplit(module)
	dst := graph.File(config.Paths.Bin, name+config.HostOS.ExecutableExtension)
	if graph.Creator(dst) == nil {
		GoCommand("install", module).Creates(dst).AlwaysRun()
	}
	return dst
}

// GoTest creates a new Step that runs "go test" on the supplied module.
// It returns a virtual entity that represents the test output.
func GoTest(module string) graph.Entity {
	test := graph.Virtual("")
	GoCommand("test", module).Creates(test)
	graph.List("go_test").DependsOn(test)
	return test
}

// GoRun returns a Step that runs "go run" with the supplied go file
// and arguments.
func GoRun(gofile string, args ...string) *graph.Step {
	return GoCommand(append([]string{"run", gofile}, args...)...)
}

// GoSrcPath returns the full path to a file or directory inside the GoPath.
func GoSrcPath(path string) string {
	// TODO: search GoPath
	return maker.Path(config.GoPath[0], "src", path)
}
