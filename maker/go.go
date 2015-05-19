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

import "sync"

var golock sync.Mutex

// GoCommand runs "go" with the specified arguments.
func GoCommand(args ...string) *Step {
	golock.Lock()
	defer golock.Unlock()
	return Command(goTool, args...)
}

// GoInstall builds a new Step that runs "go install" on the supplied module.
// It will return the resulting binary entity.
// The step will depend on the go tool, and will be set to always run if
// depended on.
func GoInstall(module string) Entity {
	_, name := PathSplit(module)
	dst := File(Paths.Bin, name+HostExecutableExtension)
	if Creator(dst) == nil {
		GoCommand("install", module).Creates(dst).AlwaysRun()
	}
	return dst
}

// GoTest creates a new Step that runs "go test" on the supplied module.
// It returns a virtual entity that represents the test output.
func GoTest(module string) Entity {
	test := Virtual("")
	GoCommand("test", module).Creates(test)
	return test
}

// GoRun returns a Step that runs "go run" with the supplied go file
// and arguments.
func GoRun(gofile string, args ...string) *Step {
	return GoCommand(append([]string{"run", gofile}, args...)...)
}

// GoSrcPath returns the full path to a file or directory inside the GoPath.
func GoSrcPath(path string) string {
	// TODO: search GoPath
	return Path(GoPath[0], "src", path)
}
