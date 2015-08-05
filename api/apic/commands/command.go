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

package commands

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"android.googlesource.com/platform/tools/gpu/parse"
)

// Command holds information about a runnable api command.
type Command struct {
	Name      string                         // The name of the command
	Run       func(flags flag.FlagSet) error // the action for the command
	ShortHelp string                         // Help for how to use the command
	Flags     flag.FlagSet                   // The command line flags it accepts
}

const (
	maxErrors = 10
)

var (
	commands = []*Command{}
	verbose  = flag.Bool("v", false, "Verbosity")
)

// Register adds a new command to the supported set, it will panic if a
// duplicate name is encountered.
func Register(c *Command) {
	if len(Filter(c.Name)) != 0 {
		panic(fmt.Errorf("Duplicate command name %s", c.Name))
	}
	commands = append(commands, c)
}

// Filter returns the filtered list of commands who's names match the specified
// prefix.
func Filter(prefix string) (result []*Command) {
	for _, c := range commands {
		if strings.HasPrefix(c.Name, prefix) {
			result = append(result, c)
		}
	}
	return result
}

// Usage prints message with the formatting args (if not empty) to stderr,
// prints the command usage information to stderr and then terminates the program.
func Usage(message string, args ...interface{}) error {
	err := ""
	if len(message) > 0 {
		err = fmt.Sprintf(message, args...)
		fmt.Fprint(os.Stderr, err)
	}
	fmt.Fprintf(os.Stderr, "\nApic is a tool for managing api source files\n\n")
	fmt.Fprintf(os.Stderr, "Available commands\n")
	for _, c := range commands {
		fmt.Fprintf(os.Stderr, "  %s : %s\n", c.Name, c.ShortHelp)
	}
	return fmt.Errorf(err)
}

// Log prints message with the formatting args to stdout.
func Logf(message string, args ...interface{}) {
	if *verbose {
		fmt.Fprintf(os.Stdout, message, args...)
	}
}

// CheckErrors will, if len(errs) > 0, print each of the error messages for the
// specified api and then terminate the program. If errs is zero length,
// CheckErrors does nothing.
func CheckErrors(apiName string, errs parse.ErrorList) error {
	if len(errs) == 0 {
		return nil
	}
	if len(errs) > maxErrors {
		errs = errs[:maxErrors]
	}
	for _, e := range errs {
		if e.At != nil {
			filename := e.At.Token().Source.Filename
			line, column := e.At.Token().Cursor()
			fmt.Fprintf(os.Stderr, "%s:%v:%v: %s\n", filename, line, column, e.Message)
		} else {
			fmt.Fprintf(os.Stderr, "%s: %s\n", apiName, e.Message)
		}
	}
	if len(errs) > maxErrors {
		fmt.Fprintf(os.Stderr, "And %d more errors\n", len(errs)-maxErrors)
	}
	fmt.Fprintf(os.Stderr, "Stack of first error:\n%s\n", errs[0].Stack)
	return errs
}
