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

// embed is a tool to embed text files as resource.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const header = `
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

package %s

`

func run() error {
	pwd, err := filepath.Abs(".")
	if err != nil {
		return err
	}
	_, pkg := filepath.Split(pwd)
	output, err := os.Create("embed.go")
	if err != nil {
		return err
	}
	defer func() { output.Close() }()
	fmt.Fprintf(output, header, pkg)
	files, err := ioutil.ReadDir(".")
	if err != nil {
		return err
	}
	for _, info := range files {
		if info.IsDir() {
			continue
		}
		extension := filepath.Ext(info.Name())
		if extension == ".go" {
			continue
		}
		name := strings.Replace(info.Name(), ".", "_", -1)
		data, err := ioutil.ReadFile(info.Name())
		if err != nil {
			return err
		}
		fmt.Fprintf(output, "const %s = \"", name)
		for _, b := range data {
			fmt.Fprintf(output, "\\x%02x", b)
		}
		fmt.Fprintf(output, "\"\n")
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "embed failed: %v\n", err)
		os.Exit(1)
	}
}
