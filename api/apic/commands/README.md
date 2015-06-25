# commands
--
    import "android.googlesource.com/platform/tools/gpu/api/apic/commands"


## Usage

#### func  CheckErrors

```go
func CheckErrors(apiName string, errs parse.ErrorList)
```
CheckErrors will, if len(errs) > 0, print each of the error messages for the
specified api and then terminate the program. If errs is zero length,
CheckErrors does nothing.

#### func  Filter

```go
func Filter(prefix string) (result []*Command)
```
Filter returns the filtered list of commands who's names match the specified
prefix.

#### func  Logf

```go
func Logf(message string, args ...interface{})
```
Log prints message with the formatting args to stdout.

#### func  MaybeError

```go
func MaybeError(file string, err error)
```
MaybeError will, if err is not nil, print the error err along with the optional
file path to stderr and then terminate the program. If err is nil, MaybeError
does nothing.

#### func  Register

```go
func Register(c *Command)
```
Register adds a new command to the supported set, it will panic if a duplicate
name is encountered.

#### func  Usage

```go
func Usage(message string, args ...interface{})
```
Usage prints message with the formatting args (if not empty) to stderr, prints
the command usage information to stderr and then terminates the program.

#### type Command

```go
type Command struct {
	Name      string                   // The name of the command
	Run       func(flags flag.FlagSet) // the action for the command
	ShortHelp string                   // Help for how to use the command
	Flags     flag.FlagSet             // The command line flags it accepts
}
```

Command holds information about a runnable api command.
