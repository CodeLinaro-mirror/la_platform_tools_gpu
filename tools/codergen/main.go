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
	"os"
	"runtime"

	"android.googlesource.com/platform/tools/gpu/tools/codergen/generate"
	"android.googlesource.com/platform/tools/gpu/tools/codergen/scan"
	"android.googlesource.com/platform/tools/gpu/tools/codergen/template"
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
	scanner := scan.New(wd, *forceSource)
	if *verbose {
		fmt.Printf("Scanning\n")
	}
	args := flag.Args()
	if len(args) == 0 {
		args = append(args, "./...")
	}
	for _, arg := range args {
		if err := scanner.Scan(arg, *verbose); err != nil {
			return err
		}
	}
	if *verbose {
		fmt.Printf("Processing\n")
	}
	if err := scanner.Process(); err != nil {
		return err
	}
	modules, err := generate.From(scanner)
	if err != nil {
		return err
	}
	if *verbose {
		fmt.Printf("Generating\n")
	}
	t := template.New()
	info := copyright.Info{Tool: "codergen", Year: "2015"}
	gen := func(name string, arg interface{}, output string, reflow template.PostProcess) error {
		out := output
		if *nowrite {
			out = ""
		}
		changed, err := t.Generate(arg, name, arg, out, reflow)
		if err != nil {
			return err
		}
		if changed {
			if *nowrite {
				fmt.Printf("Not writing %s\n", output)
			} else if *verbose {
				fmt.Printf("Generated %s\n", output)
			}
		} else if *verbose {
			fmt.Printf("No change for %s\n", output)
		}
		return nil
	}

	for _, m := range modules {
		if *golang {
			if err := generate.Go(m, info, gen); err != nil {
				return err
			}
		}
		_, doJava := m.Directives["java.package"]
		if *java != "" && !m.IsTest && doJava {
			if err := generate.Java(m, info, gen, *java); err != nil {
				return err
			}
		}
		_, doCpp := m.Directives["cpp"]
		if *cpp != "" && !m.IsTest && doCpp {
			if err := generate.Cpp(m, info, gen, *cpp); err != nil {
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
