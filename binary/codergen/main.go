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

// The codergen command parses go code to automatically generate encoders and
// decoders for the structs it finds.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/build"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"android.googlesource.com/platform/tools/gpu/binary/generate"

	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/go/types"
)

var (
	verbose = flag.Bool("v", false, "verbose messages")
	golang  = flag.Bool("go", false, "generate go code")
	java    = flag.String("java", "", "the path to generate files in")
)

const usage = `codergen: A tool to generate coders for go structs.
Usage: codergen [--go] [--java=file] <args>...
  -help: show this help message
`

type Entry struct {
	Output    string
	File      generate.File
	Generator func(*generate.File) ([]byte, error)
}

func scan(entry string, wd string, config *loader.Config) error {
	base := strings.TrimSuffix(entry, "...")
	pkg, err := build.Default.Import(base, wd, build.FindOnly)
	if err != nil {
		return err
	}
	if *verbose {
		fmt.Printf("%s from %s\n", pkg.ImportPath, pkg.Dir)
	}
	if len(base) == len(entry) {
		config.ImportWithTests(pkg.ImportPath)
	} else {
		err = filepath.Walk(pkg.Dir, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				return nil
			}
			if filepath.Base(path)[0] == '.' {
				return filepath.SkipDir
			}
			name := pkg.ImportPath + strings.TrimPrefix(path, pkg.Dir)
			config.ImportWithTests(name)
			return err
		})
	}
	return err
}

func filterStructs(pkg *loader.PackageInfo) []*types.TypeName {
	result := []*types.TypeName{}
	for _, def := range pkg.Defs {
		name, ok := def.(*types.TypeName)
		if !ok {
			continue
		}
		t, ok := name.Type().(*types.Named)
		if !ok {
			continue
		}
		if _, ok := t.Underlying().(*types.Struct); !ok {
			continue
		}
		result = append(result, name)
	}
	return result
}

func run() error {
	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	flag.Usage = func() {
		fmt.Printf(usage)
		flag.PrintDefaults()
	}
	flag.Parse()
	config := &loader.Config{}
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	config.AllowErrors = true
	if !*verbose {
		config.TypeChecker.Error = func(e error) {}
	}
	if *verbose {
		fmt.Printf("Scanning\n")
	}
	args := flag.Args()
	if len(args) == 0 {
		args = append(args, "./...")
	}
	for _, arg := range args {
		if err := scan(arg, wd, config); err != nil {
			return err
		}
	}
	if *verbose {
		fmt.Printf("Loading\n")
	}
	info, err := config.Load()
	if err != nil {
		return err
	}
	if *verbose {
		fmt.Printf("Generating\n")
	}
	files := map[string]*generate.File{}
	for _, pkg := range info.Imported {
		pkgName := pkg.Pkg.Name()
		structs := filterStructs(pkg)
		for _, name := range structs {
			fileName := pkgName
			fromFile := config.Fset.File(name.Pos()).Name()
			path := filepath.Dir(fromFile)
			isTest := strings.HasSuffix(fromFile, "_test.go")
			if isTest {
				fileName += "#test"
			}
			fileName = filepath.Join(path, fileName)
			file, found := files[fileName]
			if !found {
				file = &generate.File{}
				file.Package = pkgName
				file.IsTest = isTest
				file.Path = path
				file.Imports = make(map[string]struct{})
				files[fileName] = file
			}
			s := generate.FromTypename(pkg.Pkg, name, file.Imports)
			if s != nil {
				file.Structs = append(file.Structs, s)
			}
		}
	}
	for _, file := range files {
		if len(file.Structs) == 0 {
			continue
		}
		generate.Sort(file.Structs)
		if *golang {
			entry := Entry{
				File:      *file,
				Output:    file.Package + "_binary.go",
				Generator: generate.GoFile,
			}
			entry.File.Generated = fmt.Sprintf("codergen -go")
			if file.IsTest {
				entry.Output = file.Package + "_binary_test.go"
			}
			entry.Output = path.Join(file.Path, entry.Output)
			if err := entry.Generate(); err != nil {
				return err
			}
		}
		if *java != "" && !file.IsTest {
			entry := Entry{
				File:      *file,
				Output:    filepath.Join(*java, "ObjectFactory.java"),
				Generator: generate.JavaFile,
			}
			entry.File.Generated = fmt.Sprintf("codergen -java=%s", filepath.Base(*java))
			entry.File.ClassPrefix = strings.Title(file.Package)
			i := strings.LastIndex(*java, "/com/")
			if i >= 0 {
				entry.File.Package = strings.Replace((*java)[i+1:], "/", ".", -1)
			}
			if err := entry.Generate(); err != nil {
				return err
			}
		}
	}
	return nil
}

func (e *Entry) Generate() error {
	result, err := e.Generator(&e.File)
	if err != nil {
		return err
	}
	current, err := ioutil.ReadFile(e.Output)
	if err == nil && bytes.Equal(result, current) {
		if *verbose {
			fmt.Printf("No change for %s\n", e.Output)
		}
		return nil
	}
	if *verbose {
		fmt.Printf("Generate %s\n", e.Output)
	}
	return ioutil.WriteFile(e.Output, result, os.ModePerm)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "codergen failed: %v\n", err)
		os.Exit(1)
	}
}
