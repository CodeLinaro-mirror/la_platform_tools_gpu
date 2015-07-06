# log
--
    import "android.googlesource.com/platform/tools/gpu/log"

Package log provides a hierarchical logger interface and implementations of the
interface.

## Usage

```go
var ConstantValues schema.Constants
```

#### func  Alertf

```go
func Alertf(l Interface, msg string, args ...interface{})
```
Alertf calls l.Log with the Alert severity.

#### func  Criticalf

```go
func Criticalf(l Interface, msg string, args ...interface{})
```
Criticalf calls l.Log with the Critical severity.

#### func  Debugf

```go
func Debugf(l Interface, msg string, args ...interface{})
```
Debugf calls l.Log with the Debug severity.

#### func  E

```go
func E(l Interface, msg string, args ...interface{})
```
E calls l.Log with the Error severity.

#### func  Emergencyf

```go
func Emergencyf(l Interface, msg string, args ...interface{})
```
Emergencyf calls l.Log with the Emergency severity.

#### func  Errorf

```go
func Errorf(l Interface, msg string, args ...interface{})
```
Errorf calls l.Log with the Error severity.

#### func  I

```go
func I(l Interface, msg string, args ...interface{})
```
I calls l.Log with the Info severity.

#### func  Infof

```go
func Infof(l Interface, msg string, args ...interface{})
```
Infof calls l.Log with the Info severity.

#### func  Noticef

```go
func Noticef(l Interface, msg string, args ...interface{})
```
Noticef calls l.Log with the Notice severity.

#### func  W

```go
func W(l Interface, msg string, args ...interface{})
```
W calls l.Log with the Warning severity.

#### func  Warningf

```go
func Warningf(l Interface, msg string, args ...interface{})
```
Warningf calls l.Log with the Warning severity.

#### type Entry

```go
type Entry struct {
	Severity  Severity  // The Entry severity
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

#### type Interface

```go
type Interface interface {
	// Log writes an error message to the logger with the specified severity.
	// Arguments are handled in the manner of fmt.Printf.
	Log(severity Severity, msg string, args ...interface{})
}
```

Interface declares the methods for an object that accepts logging messages.

#### type Logger

```go
type Logger interface {
	Interface

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
func Testing(t delegate) Logger
```
Testing returns a Logger that writes to t's log methods.

#### func  Writer

```go
func Writer(info, warn, err io.Writer, done func()) Logger
```
Writer returns a Logger that writes to the supplied streams.

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

#### func (Nop) Log

```go
func (Nop) Log(s Severity, msg string, args ...interface{})
```
Log does nothing

#### type Severity

```go
type Severity int
```

Severity defines the severity of a logging message. The levels match the ones
defined in rfc5424 for syslog.

```go
const (
	// Emergency indicates the system is unusable, no further data should be trusted.
	Emergency Severity = 0
	// Alert indicates action must be taken immediately.
	Alert Severity = 1
	// Critical indicates errors severe enough to terminate processing.
	Critical Severity = 2
	// Error indicates non terminal failure conditions that may have an effect on results.
	Error Severity = 3
	// Warning indicates issues that might affect performance or compatibility, but could be ignored.
	Warning Severity = 4
	// Notice indicates normal but significant conditions.
	Notice Severity = 5
	// Info indicates minor informational messages that should generally be ignored.
	Info Severity = 6
	// Debug indicates verbose debug-level messages.
	Debug Severity = 7
)
```

#### func (*Severity) Parse

```go
func (v *Severity) Parse(s string) error
```

#### func (Severity) String

```go
func (v Severity) String() string
```

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

#### func (*Splitter) Log

```go
func (s *Splitter) Log(severity Severity, msg string, args ...interface{})
```
Logf will call Logf with the same arguments on all logs passed to Add.
