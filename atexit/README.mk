# atexit
--
    import "android.googlesource.com/platform/tools/gpu/atexit"

Package atexit enables registration of cleanup goroutines to be run at program
exit or when the process receives an interruption or kill signal.

## Usage

#### func  Exit

```go
func Exit(code int)
```
Exit calls all registered callbacks once in separate goroutines, waiting for
them to complete for a duration of at least the maximum timeout value given to
Register, then calls os.Exit with the given return code.

#### func  Register

```go
func Register(f func(), timeout time.Duration)
```
Register adds the given function f to a list of callbacks that will get called
when the current process receives an interruption or kill signal, or when Exit
is explicitly called. f will be given a duration of at least timeout to
complete.
