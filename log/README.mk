# log
--
    import "android.googlesource.com/platform/tools/gpu/log"

Package log provides a hierarchical logger interface and implementations of the
interface.

## Usage

#### type Entry

```go
type Entry struct {
	Kind      Kind      // The Entry kind
	Message   string    // The message text
	Scope     string    // The scope of the message
	Context   uint32    // The context identifier for the message
	Timestamp time.Time // The time at which the message was raised
}
```

Entry is a single Info, Warning or Error message written to the chan passed to
Channel.

#### func (*Entry) String

```go
func (e *Entry) String() string
```
String returns the string representation of the entry with timestamp.

#### func (*Entry) StringNoTimestamp

```go
func (e *Entry) StringNoTimestamp() string
```
StringNoTimestamp String returns the string representation of the entry without
the timestamp.

#### type FlushRequest

```go
type FlushRequest chan struct{}
```

FlushRequest is a signal to flush the Logger and is written to the output
channel passed to Channel whenever Flush() is called. On receiving a
FlushRequest, any pending messages should be flushed and the FlushRequest should
be closed.

#### type Kind

```go
type Kind int
```

Kind defines an entry's message kind - either Info, Warning or Error.

```go
const (
	Info Kind = iota
	Warning
	Error
)
```

#### func (Kind) String

```go
func (i Kind) String() string
```

#### type Logger

```go
type Logger interface {
	// Info writes an information message to the logger. These message types are intented to be used
	// for non-critial, expected events. Arguments are handled in the manner of fmt.Printf.
	Info(msg string, args ...interface{})

	// Warning writes a warning message to the logger. This is intented to be used for unexpected but
	// non-critical events. Arguments are handled in the manner of fmt.Printf.
	Warning(msg string, args ...interface{})

	// Error writes an error message to the logger. This is intented to be used for unexpected and
	// critical error events. Arguments are handled in the manner of fmt.Printf.
	Error(msg string, args ...interface{})

	// Enter creates a new logger scoped within the existing logger. This can be used to produce
	// hierarchical log messages.
	Enter(name string) Logger

	// Fork creates a new logger with the same scope as the existing logger, but with a new context
	// identifier. It is good practice to fork logs before passing to another goroutine so that
	// messages can be associated with their goroutine of execution.
	Fork() Logger

	// Flush ensures that any pending messages are written by the logger.
	Flush()

	// Close closes the logger, automatically flushing any remaining messages.
	// After calling Close, no other methods can be called on the logger.
	Close()
}
```

Logger is the interface for types that implement a hierarchical message logger.

#### func  Channel

```go
func Channel(out chan<- interface{}) Logger
```
Channel is an implementation of Logger interface that writes out an Entry to the
specified chan for every message, and a FlushRequest when Flush is called.

#### func  File

```go
func File(path string) (Logger, error)
```
File creates a new Logger that will write messages to the specified file path.
If a file exists at the specified path, then this file will be overwritten.

#### func  Std

```go
func Std() Logger
```
Std returns a Logger that writes to stdout and stderr.

#### func  Testing

```go
func Testing(t *testing.T) Logger
```
Testing returns a Logger that writes to t's log methods.

#### type Nop

```go
type Nop struct{}
```

Nop is an implementation of Logger interface that does nothing for Info,
Warning, Error and Flush. Enter and Fork both return Nops.

#### func (Nop) Close

```go
func (Nop) Close()
```
Close does nothing

#### func (Nop) Enter

```go
func (Nop) Enter(name string) Logger
```
Enter returns the same Nop implementation of Logger

#### func (Nop) Error

```go
func (Nop) Error(msg string, args ...interface{})
```
Error does nothing

#### func (Nop) Flush

```go
func (Nop) Flush()
```
Flush does nothing

#### func (Nop) Fork

```go
func (Nop) Fork() Logger
```
Fork returns the same Nop implementation of Logger

#### func (Nop) Info

```go
func (Nop) Info(msg string, args ...interface{})
```
Info does nothing

#### func (Nop) Warning

```go
func (Nop) Warning(msg string, args ...interface{})
```
Warning does nothing

#### type Splitter

```go
type Splitter struct {
}
```

Splitter is an implementation of the Logger interface that delegates all method
calls to all logs passed to Add.

#### func (*Splitter) Add

```go
func (s *Splitter) Add(l Logger)
```
Add adds l to the list of loggers that the Splitter will delegate calls to.

#### func (*Splitter) Close

```go
func (s *Splitter) Close()
```
Close will call Close on all logs passed to Add.

#### func (*Splitter) Enter

```go
func (s *Splitter) Enter(name string) Logger
```
Enter will call Enter with the same argument on all logs passed to Add.

#### func (*Splitter) Error

```go
func (s *Splitter) Error(msg string, args ...interface{})
```
Error will call Error with the same arguments on all logs passed to Add.

#### func (*Splitter) Flush

```go
func (s *Splitter) Flush()
```
Flush will call Flush on all logs passed to Add.

#### func (*Splitter) Fork

```go
func (s *Splitter) Fork() Logger
```
Fork will call Fork on all logs passed to Add.

#### func (*Splitter) Info

```go
func (s *Splitter) Info(msg string, args ...interface{})
```
Info will call Info with the same arguments on all logs passed to Add.

#### func (*Splitter) Warning

```go
func (s *Splitter) Warning(msg string, args ...interface{})
```
Warning will call Warning with the same arguments on all logs passed to Add.
