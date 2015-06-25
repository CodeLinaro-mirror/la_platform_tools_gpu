# utils
--
    import "android.googlesource.com/platform/tools/gpu/integration/replay/utils"

Package utils contains utilities for replay integration tests.

## Usage

#### func  FindLocalDevice

```go
func FindLocalDevice(t *testing.T, mgr *replay.Manager) replay.Device
```
FindLocalDevice returns the replay Device for the local host. If the local host
cannot be found then the test fails and nil is returned.

#### func  StoreCapture

```go
func StoreCapture(t *testing.T, atoms atom.List, d database.Database, l log.Logger) service.CaptureId
```
StoreCapture encodes and writes the atom list to the database, returning an
identifier to the newly constructed and stored Capture.
