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

// Codergen is a tool to parse go code and automatically generate encoders and
// decoders for the structs it finds.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"android.googlesource.com/platform/tools/gpu/binary/generate"

	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/go/types"
)

var (
	golang = flag.String("go", "", "the go file to generate")
	java   = flag.String("java", "", "the java file to generate")
)

func run() error {
	flag.Parse()
	config := loader.Config{SourceImports: true}
	_, err := config.FromArgs(flag.Args(), false)
	if err != nil {
		return err
	}
	config.AllowErrors = true
	info, err := config.Load()
	if err != nil {
		return err
	}
	file := generate.File{}
	for _, pkg := range info.Created {
		file.Package = pkg.Pkg.Name()
		for _, def := range pkg.Defs {
			if n, ok := def.(*types.TypeName); ok {
				if t, ok := n.Type().(*types.Named); ok {
					if _, ok := t.Underlying().(*types.Struct); ok {
						file.Structs = append(file.Structs, generate.FromTypename(pkg.Pkg, n))
					}
				}
			}
		}
	}
	generate.Sort(file.Structs)
	if *golang != "" {
		file := file
		file.Generated = fmt.Sprintf("codergen -go=%s", *golang)
		result, err := generate.GoFile(&file)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(*golang, result, os.ModePerm)
		if err != nil {
			return err
		}
	}
	if *java != "" {
		file := file
		file.Generated = fmt.Sprintf("codergen -java=%s", filepath.Base(*java))
		// pick off the package part
		smashed := strings.Split(*java, "/com/")
		file.Package = "com." + strings.Replace(smashed[len(smashed)-1], "/", ".", -1)
		filename := filepath.Join(*java, "ObjectFactory.java")
		result, err := generate.JavaFile(&file)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(filename, result, os.ModePerm)
		if err != nil {
			return err
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
