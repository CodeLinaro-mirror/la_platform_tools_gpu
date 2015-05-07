// Copyright (C) 2014 The Android Open Source Project
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

// copyright is a tool to maintain copyright headers.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"android.googlesource.com/platform/tools/gpu/tools/copyright"
)

var (
	usage_header = `clean_generated finds and deletes generated source files

usage: clean_generated [options]
options:`
	noactions = flag.Bool("n", false,
		"don't perform any actions, just print information")
	usage_footer = `
The search is rooted at the current working directory.
It finds files with a known extension and a known generated file header comment.
If the -n flag is not specified, the file will then be removed.
`
)

func run() error {
	flag.Usage = func() {
		fmt.Println(usage_header)
		flag.PrintDefaults()
		fmt.Print(usage_footer)
	}
	flag.Parse()
	paths := flag.Args()
	if len(paths) == 0 {
		paths = []string{"."}
	}
	for _, path := range paths {
		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			if copyright.FindExtension(filepath.Ext(path)) == nil {
				return nil
			}
			file, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			if copyright.MatchGenerated(file) == 0 {
				return nil
			}
			fmt.Printf("rm %s\n", path)
			if !*noactions {
				os.Remove(path)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "generated failed: %v\n", err)
		os.Exit(1)
	}
}
