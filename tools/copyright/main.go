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
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type copyright struct {
	current string
	old     []string
	invalid []string
	skip    []string
}

var copyrights = map[string]copyright{}

type compiled struct {
	current []byte
	old     []*regexp.Regexp
	invalid []*regexp.Regexp
	skip    []*regexp.Regexp
}

func pattern(s string) *regexp.Regexp {
	rx, err := regexp.Compile("^" + strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}
	return rx
}

func compile() map[string]compiled {
	result := map[string]compiled{}
	for name, in := range copyrights {
		out := compiled{}
		out.current = []byte(strings.TrimSpace(in.current))
		for _, s := range in.old {
			out.old = append(out.old, pattern(s))
		}
		for _, s := range in.invalid {
			out.invalid = append(out.invalid, pattern(s))
		}
		for _, s := range in.skip {
			out.skip = append(out.skip, pattern(s))
		}
		result[name] = out
	}
	return result
}

func update(path string, reason string, header []byte, body []byte) error {
	fmt.Printf("Copyright on %s was %s\n", path, reason)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(header)
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
	headers := compile()
	return filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		extension := filepath.Ext(path)
		c, ok := headers[extension]
		if !ok {
			return nil
		}
		file, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		if bytes.HasPrefix(file, c.current) {
			return nil
		}
		for _, skip := range c.skip {
			if skip.Match(file) {
				return nil
			}
		}
		for _, p := range c.old {
			match := p.Find(file)
			if len(match) > 0 {
				return update(path, "out of date", c.current, file[len(match):])
			}
		}
		for _, p := range c.invalid {
			match := p.Find(file)
			if len(match) > 0 {
				return update(path, "invalid", c.current, file[len(match):])
			}
		}
		return update(path, "missing", c.current, file)
	})
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "copyright failed: %v\n", err)
		os.Exit(1)
	}
}
