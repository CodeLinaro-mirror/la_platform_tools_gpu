# task
--
    import "android.googlesource.com/platform/tools/gpu/task"

Package task contains types that can be used to create cancelable tasks.

## Usage

#### type CancelSignal

```go
type CancelSignal <-chan hidden
```

CancelSignal is used by Runner to check whether they have been canceled.

#### func (CancelSignal) Check

```go
func (c CancelSignal) Check()
```
Check stops execution of the Runner if it has been canceled. If Check is called
and the Runner has been canceled then the deferred functions of the Runner will
be executed and then the Runner will immediately cease execution. If the Runner
has not been canceled, then Check simply returns and the Runable continues to
execute.

#### type Runner

```go
type Runner interface {
	// Run is called by a Task to execute lengthy logic on a new go-routine.
	// The Run method should periodically check whether it should continue
	// execution by calling Check() on the CancelSignal.
	Run(CancelSignal)
}
```

Runner is the interface implemented by types that implement the Run method.

#### type Task

```go
type Task struct {
}
```

Tasks are the executors of Runners. All methods on a Task support concurrent
go-routine usage.

#### func  New

```go
func New() Task
```
New constructs and returns a new Task.

#### func (*Task) Cancel

```go
func (t *Task) Cancel()
```
Cancel signals to the last Runner started with Run() that it should stop.
Repeated calls to Cancel() will do nothing.

#### func (*Task) Run

```go
func (t *Task) Run(r Runner)
```
Run begins execution of the Runner r on a new go-routine. If a previous Runner
was started by this Task, then it is automatically canceled. Note however, that
Run does not wait for the previous Runner to return before r is started. The
Runner can be canceled by calling Cancel on this Task.
