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
	Name      string                   // The name of the command
	Run       func(flags flag.FlagSet) // the action for the command
	ShortHelp string                   // Help for how to use the command
	Flags     flag.FlagSet             // The command line flags it accepts
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

// MaybeError will, if err is not nil, print the error err along with the
// optional file path to stderr and then terminate the program. If err is nil,
// MaybeError does nothing.
func MaybeError(file string, err error) {
	if err == nil {
		return
	}
	if len(file) == 0 {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	} else {
		fmt.Fprintf(os.Stderr, "%s: %s\n", file, err)
	}
	os.Exit(1)
}

// Usage prints message with the formatting args (if not empty) to stderr,
// prints the command usage information to stderr and then terminates the program.
func Usage(message string, args ...interface{}) {
	if len(message) > 0 {
		fmt.Fprintf(os.Stderr, message, args...)
	}
	fmt.Fprintf(os.Stderr, "\nApic is a tool for managing api source files\n\n")
	fmt.Fprintf(os.Stderr, "Available commands\n")
	for _, c := range commands {
		fmt.Fprintf(os.Stderr, "  %s : %s\n", c.Name, c.ShortHelp)
	}
	os.Exit(1)
}

// Log prints message with the formatting args to stdout.
func Log(message string, args ...interface{}) {
	if *verbose {
		fmt.Fprintf(os.Stdout, message, args...)
	}
}

// CheckErrors will, if len(errs) > 0, print each of the error messages for the
// specified api and then terminate the program. If errs is zero length,
// CheckErrors does nothing.
func CheckErrors(apiName string, errs parse.ErrorList) {
	if len(errs) == 0 {
		return
	}
	if len(errs) > maxErrors {
		errs = errs[:maxErrors]
	}
	for _, e := range errs {
		if e.At != nil {
			line, column := e.At.Token().Cursor()
			fmt.Fprintf(os.Stderr, "%s:%v:%v: %s\n", apiName, line, column, e.Message)
		} else {
			fmt.Fprintf(os.Stderr, "%s: %s\n", apiName, e.Message)
		}
	}
	if len(errs) > maxErrors {
		fmt.Fprintf(os.Stderr, "And %d more errors\n", len(errs)-maxErrors)
	}
	fmt.Fprintf(os.Stderr, "Stack of first error:\n%s\n", errs[0].Stack)
	os.Exit(3)
}
