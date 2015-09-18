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

import (
	"fmt"
	"os"
	"path/filepath"
)

// Root returns the repo-root of the project.
var RepoRoot = func() Root {
	gopaths := filepath.SplitList(os.Getenv("GOPATH"))
	for _, gopath := range gopaths {
		if File(gopath).Join("src", "android.googlesource.com", "platform", "tools", "gpu", "maker", "build", "root.go").Exists() {
			return Root{Name: "repo", Path: File(File(gopath).Join("..", "..").Absolute())}
		}
	}
	panic(fmt.Errorf("Project could not be found in any GOPATH: %v", gopaths))
}()
