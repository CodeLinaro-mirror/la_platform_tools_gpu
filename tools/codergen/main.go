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
	"sync"

	"android.googlesource.com/platform/tools/gpu/tools/codergen/generate"
	"android.googlesource.com/platform/tools/gpu/tools/codergen/scan"
	"android.googlesource.com/platform/tools/gpu/tools/codergen/template"
	"android.googlesource.com/platform/tools/gpu/tools/copyright"
	"android.googlesource.com/platform/tools/gpu/tools/profile"
)

var (
	verbose = flag.Bool("v", false, "verbose messages")
	nowrite = flag.Bool("n", false, "don't write the files")
	golang  = flag.Bool("go", false, "generate go code")
	java    = flag.String("java", "", "the path to generate files in")
	cpp     = flag.String("cpp", "", "the path to generate files in")
	workers = flag.Int("workers", 15, "The numer of output workers to use")
)

const usage = `codergen: A tool to generate coders for go structs.
Usage: codergen [--go] [--java=file] <args>...
  -help: show this help message
`

type errors struct {
	list []error
	mu   sync.Mutex
}

func (l *errors) Add(err error) {
	l.mu.Lock()
	l.list = append(l.list, err)
	l.mu.Unlock()
}

func (l *errors) String() string {
	return l.Error()
}

func (l *errors) Error() string {
	if len(l.list) > 0 {
		return l.list[0].Error()
	} else {
		return ""
	}
}

func worker(wg *sync.WaitGroup, errs *errors, tasks chan generate.Generate) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		t := template.New()
		for task := range tasks {
			out := task.Output
			if *nowrite {
				task.Output = ""
			}
			changed, err := t.Generate(task)
			if err != nil {
				errs.Add(err)
			} else if changed {
				if *nowrite {
					fmt.Printf("Not writing %s\n", out)
				} else if *verbose {
					fmt.Printf("Generated %s\n", out)
				}
			} else if *verbose {
				fmt.Printf("No change for %s\n", out)
			}
		}
	}()
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
	defer profile.CPU()()
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	scanner := scan.New(wd)
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
	wg := sync.WaitGroup{}
	errs := errors{}
	gen := make(chan generate.Generate)
	for i := 0; i < *workers; i++ {
		worker(&wg, &errs, gen)
	}
	info := copyright.Info{Tool: scan.Tool, Year: "2015"}
	for _, m := range modules {
		if *golang {
			generate.Go(m, info, gen)
		}
		_, doJava := m.Directives["java.package"]
		if *java != "" && !m.IsTest && doJava {
			generate.Java(m, info, gen, *java)
		}
		_, doCpp := m.Directives["cpp"]
		if *cpp != "" && !m.IsTest && doCpp {
			generate.Cpp(m, info, gen, *cpp)
		}
	}
	close(gen)
	wg.Wait()
	if len(errs.list) > 0 {
		for _, err := range errs.list {
			fmt.Print(err)
		}
		return errs.list[0]
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "codergen failed: %v\n", err)
		os.Exit(1)
	}
}
