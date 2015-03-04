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
	"path/filepath"

	"android.googlesource.com/platform/tools/gpu/tools/copyright"
)

var (
	usage_header = `copyright finds and fixes bad copyright headers

usage: copyright [options]
options:`
	noactions = flag.Bool("n", false,
		"don't perform any actions, just print information")
	usage_footer = `
The search is rooted at the current working directory.
It will attempt to fix incorrect copyright headers unless you specify the
-n flag, and will only replace copyright headers patterns that it knows about.
For unknown comments it will just prepend the correct copyright header.
It will not touch file extensions that it does not know the header type for.
`
)

func update(path string, reason string, header string, body []byte) error {
	fmt.Printf("Copyright on %s was %s\n", path, reason)
	if *noactions {
		return nil
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(header)
	if err != nil {
		return err
	}
	_, err = file.Write(body)
	if err != nil {
		return err
	}
	return nil
}

func run() error {
	flag.Usage = func() {
		fmt.Println(usage_header)
		flag.PrintDefaults()
		fmt.Print(usage_footer)
	}
	flag.Parse()
	return filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		extension := filepath.Ext(path)
		l := copyright.FindExtension(extension)
		if l == nil {
			return nil
		}
		file, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		if i := l.MatchCurrent(file); i > 0 {
			return nil
		}
		if i := copyright.MatchGenerated(file); i > 0 {
			return nil
		}
		if i := l.MatchOld(file); i > 0 {
			return update(path, "out of date", l.Emit, file[i:])
		}
		if i := copyright.MatchNormal(file); i > 0 {
			return update(path, "invalid", l.Emit, file[i:])
		}
		return update(path, "missing", l.Emit, file)
	})
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "copyright failed: %v\n", err)
		os.Exit(1)
	}
}
