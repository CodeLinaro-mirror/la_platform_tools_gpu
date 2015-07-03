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
	"sort"
	"strings"

	"android.googlesource.com/platform/tools/gpu/binary/generate"
	"android.googlesource.com/platform/tools/gpu/tools/copyright"
)

var (
	verbose     = flag.Bool("v", false, "verbose messages")
	nowrite     = flag.Bool("n", false, "don't write the files")
	forceSource = flag.Bool("s", true, "force source only")
	golang      = flag.Bool("go", false, "generate go code")
	java        = flag.String("java", "", "the path to generate files in")
	cpp         = flag.String("cpp", "", "the path to generate files in")
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

func scan(entry string, loader *generate.Loader) error {
	base := strings.TrimSuffix(entry, "...")
	pkg, err := build.Default.Import(base, loader.Path, build.FindOnly)
	if err != nil {
		return err
	}
	if *verbose {
		fmt.Printf("%s from %s\n", pkg.ImportPath, pkg.Dir)
	}
	if len(base) == len(entry) {
		loader.ScanPackage(pkg.ImportPath)
		return nil
	} else {
		return filepath.Walk(pkg.Dir, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				return nil
			}
			if filepath.Base(path)[0] == '.' || filepath.Base(path)[0] == '_' {
				return filepath.SkipDir
			}
			name := pkg.ImportPath + strings.TrimPrefix(path, pkg.Dir)
			if *verbose {
				fmt.Printf("Reading %s\n", name)
			}
			loader.ScanPackage(name)
			return nil
		})
	}
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
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	loader := generate.NewLoader(wd, *forceSource)
	if *verbose {
		fmt.Printf("Scanning\n")
	}
	args := flag.Args()
	if len(args) == 0 {
		args = append(args, "./...")
	}
	for _, arg := range args {
		if err := scan(arg, loader); err != nil {
			return err
		}
	}
	if *verbose {
		fmt.Printf("Processing\n")
	}
	if err := loader.Process(); err != nil {
		return err
	}
	if *verbose {
		fmt.Printf("Generating\n")
	}
	gen := generate.NewGenerator()
	for _, dir := range loader.Directories {
		if !dir.Scan {
			continue
		}
		if dir.Module.Output != nil {
			if err := output(gen, dir.Module.Output); err != nil {
				return err
			}
		}
		if dir.Test.Output != nil {
			if err := output(gen, dir.Test.Output); err != nil {
				return err
			}
		}
	}
	return nil
}

func output(gen *generate.Generator, file *generate.File) error {
	if len(file.Structs) == 0 && len(file.Constants) == 0 {
		return nil
	}
	if _, ignored := file.Directives["ignore"]; ignored {
		return nil
	}
	generate.Sort(file.Structs)
	for _, c := range file.Constants {
		sort.Sort(c.(sort.Interface))
	}
	if *golang {
		entry := Entry{
			File:      *file,
			Output:    file.Package + "_binary.go",
			Generator: gen.GoFile,
		}
		entry.File.Copyright = copyright.Build(
			"generated_by", copyright.Info{
				Tool: "codergen -go",
				Year: "2015",
			})
		if file.IsTest {
			entry.Output = file.Package + "_binary_test.go"
		}
		entry.Output = path.Join(file.Path, entry.Output)
		if err := entry.Generate(); err != nil {
			return err
		}
	}
	javaPackage, doJava := file.Directives["java.package"]
	if *java != "" && !file.IsTest && doJava {
		source, _ := file.Directives["java.source"]
		indent, _ := file.Directives["java.indent"]
		member, _ := file.Directives["java.member_prefix"]
		class, _ := file.Directives["java.class_prefix"]
		pkgPath := strings.Replace(javaPackage, ".", "/", -1)
		entry := Entry{
			File:      *file,
			Output:    filepath.Join(*java, source, pkgPath, "ObjectFactory.java"),
			Generator: gen.JavaFile,
		}
		entry.File.Copyright = copyright.Build(
			"generated_aosp_java", copyright.Info{
				Year: "2015",
			})
		entry.File.Package = javaPackage
		entry.File.Indent = strings.Trim(indent, `"`)
		if entry.File.Indent == "" {
			entry.File.Indent = "    "
		}
		entry.File.MemberPrefix = member
		entry.File.ClassPrefix = class
		if err := entry.Generate(); err != nil {
			return err
		}
	}
	cppNamespace, doCpp := file.Directives["cpp"]
	if *cpp != "" && !file.IsTest && doCpp {
		entry := Entry{
			File:      *file,
			Output:    filepath.Join(*cpp, cppNamespace+".h"),
			Generator: gen.CppFile,
		}
		entry.File.Package = cppNamespace
		entry.File.Copyright = copyright.Build(
			"generated_by", copyright.Info{
				Tool: fmt.Sprintf("codergen -cpp=%s", filepath.Base(*cpp)),
				Year: "2015",
			})
		if err := entry.Generate(); err != nil {
			return err
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
	if *nowrite {
		fmt.Printf("Not writing %s\n", e.Output)
		return nil
	}
	if *verbose {
		fmt.Printf("Generate %s\n", e.Output)
	}
	return ioutil.WriteFile(e.Output, result, 0666)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "codergen failed: %v\n", err)
		os.Exit(1)
	}
}
