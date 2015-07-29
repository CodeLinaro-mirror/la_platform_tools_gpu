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

package main

import (
	"flag"
	"fmt"
	"os"

	"android.googlesource.com/platform/tools/gpu/api/apic/commands"
	_ "android.googlesource.com/platform/tools/gpu/api/apic/format"
	_ "android.googlesource.com/platform/tools/gpu/api/apic/template"
	_ "android.googlesource.com/platform/tools/gpu/api/apic/validate"
	"android.googlesource.com/platform/tools/gpu/tools/profile"
)

func run() error {
	flag.Parse()
	defer profile.CPU()()
	args := flag.Args()
	if len(args) < 1 {
		return commands.Usage("Must supply a verb\n")
	}
	verb := args[0]
	matches := commands.Filter(verb)
	switch len(matches) {
	case 1:
		c := matches[0]
		commands.Logf("Running %q\n", c.Name)
		c.Flags.Parse(args[1:])
		return c.Run(c.Flags)
	case 0:
		return commands.Usage("Verb '%s' is unknown\n", verb)
	default:
		return commands.Usage("Verb '%s' is ambiguous\n", verb)
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "apic failed: %v\n", err)
		os.Exit(1)
	}
}
