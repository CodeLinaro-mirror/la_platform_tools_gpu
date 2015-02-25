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

// copyright is a tool to maintain copyright headers.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	noactions = flag.Bool("n", false,
		"don't perform any actions, just print information")
)

func check(path string) error {
	err := exec.Command("git", "ls-files", "--error-unmatch", path).Run()
	if err == nil {
		fmt.Printf("Generated file %s is checked in! Skipping.\n", path)
		return nil
	}
	err = exec.Command("git", "check-ignore", path).Run()
	if err != nil {
		fmt.Printf("Generated file %s was not ignored\n", path)
		dir := filepath.Dir(path)
		name := filepath.Base(path)
		ignore := filepath.Join(dir, ".gitignore")
		fmt.Printf("git ignore %s in %s\n", name, ignore)
		if !*noactions {
			f, err := os.OpenFile(ignore, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
			if err != nil {
				return err
			}
			_, err = f.WriteString(name + "\n")
			f.Close()
			if err != nil {
				return err
			}
		}
	}
	fmt.Printf("rm %s \n", path)
	if !*noactions {
		os.Remove(path)
	}
	return nil
}

func run() error {
	flag.Parse()
	patterns := make([]*regexp.Regexp, len(headers))
	for i, h := range headers {
		rx, err := regexp.Compile("^" + strings.TrimSpace(h))
		if err != nil {
			panic(err)
		}
		patterns[i] = rx
	}
	types := map[string]struct{}{}
	for _, e := range extensions {
		types[e] = struct{}{}
	}
	return filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		extension := filepath.Ext(path)
		if _, present := types[extension]; !present {
			return nil
		}
		file, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		generated := false
		for _, p := range patterns {
			if p.Match(file) {
				generated = true
				break
			}
		}
		if generated {
			return check(path)
		}
		return nil
	})
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "generated failed: %v\n", err)
		os.Exit(1)
	}
}
