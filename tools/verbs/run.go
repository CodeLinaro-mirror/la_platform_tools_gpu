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

package verbs

import (
	"flag"
	"fmt"
	"os"

	"android.googlesource.com/platform/tools/gpu/tools/profile"
)

var (
	verbose   = flag.Bool("v", false, "Verbosity")
	Verbosity = 0
	ShortHelp = ""
)

func run() error {
	flag.Parse()
	defer profile.CPU()()
	args := flag.Args()
	if len(args) < 1 {
		return Usage("Must supply a verb\n")
	}
	verb := args[0]
	matches := Filter(verb)
	switch len(matches) {
	case 1:
		v := matches[0]
		v.Flags.Parse(args[1:])
		if *verbose {
			Verbosity = 1
		}
		return v.Run(v.Flags)
	case 0:
		return Usage("Verb '%s' is unknown\n", verb)
	default:
		return Usage("Verb '%s' is ambiguous\n", verb)
	}
}

// Run parses the command line and then dispatchs the the verb it specifies.
func Run() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s failed: %v\n", os.Args[0], err)
		os.Exit(1)
	}
}

// Usage prints message with the formatting args (if not empty) to stderr,
// prints the command usage information to stderr and then terminates the program.
func Usage(message string, args ...interface{}) error {
	err := ""
	if len(message) > 0 {
		err = fmt.Sprintf(message, args...)
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Fprintln(os.Stderr, "")
	if ShortHelp != "" {
		fmt.Fprintln(os.Stderr, ShortHelp)
		fmt.Fprintln(os.Stderr, "")
	}
	fmt.Fprintln(os.Stderr, "Available verbs\n")
	for _, v := range verbs {
		fmt.Fprintf(os.Stderr, "  %s : %s\n", v.Name, v.ShortHelp)
	}
	return fmt.Errorf(err)
}

// Log prints message with the formatting args to stdout if verbose is on.
func Logf(message string, args ...interface{}) {
	if Verbosity > 0 {
		fmt.Fprintf(os.Stdout, message, args...)
	}
}
