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

package run

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"sort"
	"strings"

	"android.googlesource.com/platform/tools/gpu/maker/config"
	"android.googlesource.com/platform/tools/gpu/maker/graph"
)

var (
	prepares = []func(){}
)

// Register a new graph building function with the maker system.
// The function will be invoked during Run to add entities and steps to the
// build graph.
func Register(f func()) {
	prepares = append(prepares, f)
}

type stringSetFlag []string

func (f *stringSetFlag) String() string    { return strings.Join(f.Strings(), ":") }
func (f *stringSetFlag) Strings() []string { return ([]string)(*f) }

func (f *stringSetFlag) Set(value string) error {
	*f = append(*f, value)
	return nil
}

// Run should be invoked once from main.
// It parses the command line, builds the graph, and then performs the required
// action.
func Run() {
	// Get the configuration
	verbose := flag.Int("v", 1, "Verbose mode")
	do := flag.String("do", "make", "The action to perform, one of make, show or clean.")
	early := flag.Bool("early", false, "Stops the build at the first error, also disables parallel builds.")
	threads := flag.Int("threads", runtime.NumCPU(), "Set number of OS threads to use. 0 disables parallel builds.")
	targetOS := flag.String("targetos", config.TargetOS.Name, "target OS to build")
	var disables stringSetFlag
	flag.Var(&disables, "disable", "Disable a specific node")
	flag.Parse()
	var runner graph.Runner = &parallelRunner{}
	if *threads > 0 {
		runtime.GOMAXPROCS(*threads)
	} else {
		runner = &blockingRunner{}
	}
	if *early {
		config.StopOnError = true
		runner = &blockingRunner{}
	}

	config.Verbose = *verbose
	config.TargetOS = config.FindOS(*targetOS)
	// Build the entity graph
	for _, f := range prepares {
		f()
	}
	// Force disabled status from the command line
	for _, d := range disables.Strings() {
		if s := graph.Creator(d); s != nil {
			s.Disable()
		}
	}
	// Prepare the active path
	targets := flag.Args()
	meta := graph.List("")
	for _, match := range targets {
		e := graph.FindPathEntity(match)
		if e != nil {
			if graph.Creator(e) != nil {
				meta.DependsOn(e)
			}
		} else {
			// not an exact entry, so fuzzy search time
			entities := graph.FindEntities(match)
			if len(entities) == 0 {
				log.Fatalf("no entities match for %q", match)
			}
			for _, e := range entities {
				if graph.Creator(e) != nil {
					meta.DependsOn(e)
				}
			}
		}
	}
	// Perform the requested action
	switch *do {
	case "make":
		if len(meta.Inputs) == 0 {
			meta.DependsOn(graph.Default)
		}
		meta.Process(runner)
		err := meta.Wait()
		if err != nil {
			fmt.Printf("Failed: %s\n", err)
			os.Exit(1)
		} else {
			fmt.Printf("Succeeded\n")
		}
	case "show":
		if len(meta.Inputs) == 0 {
			fmt.Printf("targets available are:\n")
			strings := sort.StringSlice{}
			for _, e := range graph.Entities {
				if graph.IsVirtual(e) {
					strings = append(strings, e.Name())
				}
			}
			strings.Sort()
			for _, n := range strings {
				fmt.Printf("    %s\n", n)
			}
		} else {
			fmt.Printf("active dependancy graph is:\n")
			dumper{}.dump(meta, nil)
		}
	case "clean":
		log.Fatalf("Clean not yet supported")
	default:
		log.Fatalf("Unknown action %q", *do)
	}
}
