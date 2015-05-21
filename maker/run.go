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

package maker

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"sort"
)

const (
	// Default is the name of the entity that is built if none are supplied on the
	// command line.
	Default = "default"
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

// Run should be invoked once from main.
// It parses the command line, builds the graph, and then performs the required
// action.
func Run() {
	// Get the configuration
	verbose := flag.Int("v", 1, "Verbose mode")
	targetArch := flag.String("arch", runtime.GOARCH, "Target architecture.")
	targetOS := flag.String("os", runtime.GOOS, "Target OS.")
	do := flag.String("do", "make", "The action to perform, one of make, show or clean.")
	threads := flag.Int("threads", runtime.NumCPU(), "Set number of go routines to use. 0 disables parallel builds.")
	flag.Parse()
	if *threads > 0 {
		runtime.GOMAXPROCS(*threads)
	} else {
		Config.DisableParallel = true
	}
	Config.Verbose = *verbose
	Config.TargetArchitecture = *targetArch
	Config.TargetOS = *targetOS
	// Build the entity graph
	for _, f := range prepares {
		f()
	}
	// Prepare the active path
	targets := flag.Args()
	meta := List("")
	for _, match := range targets {
		e := FindPathEntity(match)
		if e != nil {
			if Creator(e) != nil {
				meta.DependsOn(e)
			}
		} else {
			// not an exact entry, so fuzzy search time
			entities := FindEntities(match)
			if len(entities) == 0 {
				log.Fatalf("no entities match for %q", match)
			}
			for _, e := range entities {
				if Creator(e) != nil {
					meta.DependsOn(e)
				}
			}
		}
	}
	// Perform the requested action
	switch *do {
	case "make":
		if len(meta.inputs) == 0 {
			meta.DependsOn(Default)
		}
		meta.start()
		<-meta.done
		if Errors.Failed() {
			fmt.Printf("Failed:%s\n", Errors.First())
			os.Exit(1)
		} else {
			fmt.Printf("Succeeded\n")
		}
	case "show":
		if len(meta.inputs) == 0 {
			fmt.Printf("targets available are:\n")
			strings := sort.StringSlice{}
			for _, e := range entities {
				if IsVirtual(e) {
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

type dumper map[*Step]struct{}

func (d dumper) dump(s *Step, seen []*Step) {
	if s == nil {
		fmt.Println()
		return
	}
	fmt.Printf(" [%d]", len(s.inputs))
	if _, done := d[s]; done {
		fmt.Println(" - already seen")
		for i := range seen {
			if seen[i] == s {
				err := "Error: Cyclic dependency chain found:\n"
				for i := range seen {
					err += fmt.Sprintf("  [%d]: %v\n", i, seen[i])
				}
				panic(err)
			}
		}
		return
	}
	fmt.Println()
	d[s] = struct{}{}
	for i, e := range s.inputs {
		for i := 0; i < len(seen); i++ {
			fmt.Print("  ")
		}
		fmt.Printf("(%d) %s", i+1, e)
		d.dump(Creator(e), append(seen, s))
	}
}
