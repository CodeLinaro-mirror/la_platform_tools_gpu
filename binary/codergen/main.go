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
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"android.googlesource.com/platform/tools/gpu/binary/generate"

	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/go/types"
)

var (
	golang = flag.Bool("go", false, "generate go code")
	java   = flag.String("java", "", "the java file to generate")
)

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
	flag.Parse()
	config := loader.Config{}
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	config.AllowErrors = true
	if len(flag.Args()) == 0 {
		if filenames, err := filepath.Glob(path.Join(wd, "*.go")); err != nil {
			return err
		} else {
			config.CreateFromFilenames(wd, filenames...)
		}
	} else if _, err := config.FromArgs(flag.Args(), false); err != nil {
		return err
	}
	info, err := config.Load()
	if err != nil {
		return err
	}
	files := map[string]*generate.File{}
	for _, pkg := range info.Created {
		pkgName := pkg.Pkg.Name()
		structs := filterStructs(pkg)
		for _, name := range structs {
			fileName := pkgName
			fromFile := config.Fset.File(name.Pos()).Name()
			isTest := strings.HasSuffix(fromFile, "_test.go")
			if isTest {
				fileName += "#test"
			}
			s := generate.FromTypename(pkg.Pkg, name)
			if s != nil {
				file, found := files[fileName]
				if !found {
					file = &generate.File{}
					file.Package = pkgName
					file.IsTest = isTest
					files[fileName] = file
				}
				file.Structs = append(file.Structs, s)
			}
		}
	}
	for _, file := range files {
		generate.Sort(file.Structs)
		if *golang {
			file := file
			file.Generated = fmt.Sprintf("codergen -go")
			result, err := generate.GoFile(file)
			if err != nil {
				return err
			}
			filename := file.Package+"_binary.go"
			if file.IsTest {
				filename = file.Package+"_binary_test.go"
			}
			err = ioutil.WriteFile(path.Join(wd, filename), result, os.ModePerm)
			if err != nil {
				return err
			}
		}
		if *java != "" && !file.IsTest {
			file := file
			file.Generated = fmt.Sprintf("codergen -java=%s", filepath.Base(*java))
			file.ClassPrefix = strings.Title(file.Package)
			i := strings.LastIndex(*java, "/com/")
			if i >= 0 {
				file.Package = strings.Replace((*java)[i+1:], "/", ".", -1)
			}
			filename := filepath.Join(*java, "ObjectFactory.java")
			result, err := generate.JavaFile(file)
			if err != nil {
				return err
			}
			err = ioutil.WriteFile(filename, result, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "codergen failed: %v\n", err)
		os.Exit(1)
	}
}
