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
	"go/build"
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
	t := generate.NewTemplates()
	for _, dir := range loader.Directories {
		if !dir.Scan {
			continue
		}
		if dir.Module.Output != nil {
			if err := output(t, dir.Module.Output); err != nil {
				return err
			}
		}
		if dir.Test.Output != nil {
			if err := output(t, dir.Test.Output); err != nil {
				return err
			}
		}
	}
	return nil
}

func output(t *generate.Templates, file *generate.File) error {
	if len(file.Structs) == 0 && len(file.Constants) == 0 {
		return nil
	}
	if _, ignored := file.Directives["ignore"]; ignored {
		return nil
	}
	generate.Sort(file.Structs)
	sort.Sort(&file.Constants)
	for i := range file.Constants {
		sort.Sort(&file.Constants[i])
	}
	if *golang {
		gen := generate.NewGo(file)
		gen.Copyright = copyright.Build(
			"generated_by", copyright.Info{
				Tool: "codergen -go",
				Year: "2015",
			})
		out := file.Name + "_binary.go"
		if file.IsTest {
			out = file.Name + "_binary_test.go"
		}
		out = path.Join(file.Path, out)
		if err := Generate(gen, t, out); err != nil {
			return err
		}
	}
	javaPackage, doJava := file.Directives["java.package"]
	if *java != "" && !file.IsTest && doJava {
		gen := generate.NewJava(file)
		gen.JavaPackage = javaPackage
		source, _ := file.Directives["java.source"]
		indent, _ := file.Directives["java.indent"]
		gen.MemberPrefix, _ = file.Directives["java.member_prefix"]
		gen.Copyright = strings.TrimSpace(copyright.Build(
			"generated_aosp_java", copyright.Info{
				Year: "2015",
			}))
		gen.Indent = strings.Trim(indent, `"`)
		if gen.Indent == "" {
			gen.Indent = "    "
		}
		pkgPath := strings.Replace(javaPackage, ".", "/", -1)
		for _, s := range file.Structs {
			gen.Struct.Struct = s
			out := filepath.Join(*java, source, pkgPath, gen.Struct.Name()+".java")
			if err := Generate(gen, t, out); err != nil {
				return err
			}
		}
	}
	cppNamespace, doCpp := file.Directives["cpp"]
	if *cpp != "" && !file.IsTest && doCpp {
		gen := generate.NewCpp(file)
		gen.Namespace = cppNamespace
		gen.Copyright = copyright.Build(
			"generated_by", copyright.Info{
				Tool: fmt.Sprintf("codergen -cpp=%s", filepath.Base(*cpp)),
				Year: "2015",
			})
		out := filepath.Join(*cpp, cppNamespace+".h")
		if err := Generate(gen, t, out); err != nil {
			return err
		}
	}
	return nil
}

type generator interface {
	Run(t *generate.Templates, out string) (bool, error)
}

func Generate(g generator, t *generate.Templates, path string) error {
	out := path
	if *nowrite {
		out = ""
	}
	changed, err := g.Run(t, out)
	if err != nil {
		return err
	}
	if changed {
		if *nowrite {
			fmt.Printf("Not writing %s\n", path)
		} else if *verbose {
			fmt.Printf("Generated %s\n", path)
		}
	} else if *verbose {
		fmt.Printf("No change for %s\n", path)
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "codergen failed: %v\n", err)
		os.Exit(1)
	}
}
