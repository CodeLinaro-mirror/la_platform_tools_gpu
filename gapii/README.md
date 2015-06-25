# gapii
--
    import "android.googlesource.com/platform/tools/gpu/gapii"

Package gapii provides functions for launching and communicating with the gapii
tracer.

## Usage

#### func  AdbStart

```go
func AdbStart(l log.Logger, a *adb.Action, spyport adb.Port) error
```
AdbStart launches an activity on an android device with the gapii tracer
enabled. Gapii will attempt to connect back on the specified host port to write
the trace.

#### func  Capture

```go
func Capture(logger log.Logger, port int, w io.Writer, stop chan struct{}) (int64, error)
```
Capture opens up the specified port and then waits for a capture to be
delivered. It copies the capture into the supplied writer.
