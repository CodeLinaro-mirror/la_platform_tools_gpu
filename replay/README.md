# replay
--
    import "android.googlesource.com/platform/tools/gpu/replay"

Package replay is used to issue replay requests to replay devices.

## Usage

```go
var DisableLocalDeviceCache = false
```
DisableLocalDeviceCache can be used to disable the disk-cache for the local
device. If true, it is passed as a flag to replayd on spawning. This can be used
for disabling the cache for tests.

#### func  Replay

```go
func Replay(
	i atom.ID,
	a atom.Atom,
	s *gfxapi.State,
	d database.Database,
	l log.Logger,
	b *builder.Builder) error
```
Replay issues replay operations to the replay builder b for the given atom a
with identifier i, and graphics API state s. If replaying the Atom will have an
effect on the graphics driver state, then the call to Replay will also apply the
corresponding changes to the state s.

#### type CallTiming

```go
type CallTiming struct {
	TimingInfo service.TimingInfo // The timing data.
	Error      error              // The error that occurred generating the timing, if there was one.
}
```

CallTiming represents the call timing information for a replay.

#### type Config

```go
type Config interface{}
```

Config is a user-defined type used to describe the type of replay being
requested. Replay requests made with configs that have equality (==) will likely
be batched into the same replay pass. Configs can be used to force requests into
different replay passes. For example, by issuing requests with different configs
we can prevent a profiling Request from being issued in the same pass as a
Request to render all draw calls in wireframe.

#### type Context

```go
type Context struct {
	DeviceID  service.DeviceId  // The identifier of the device being used for replay.
	CaptureID service.CaptureId // The identifier of the capture that is being replayed.
}
```

Context describes the source capture and replay target information used for
issuing a replay request.

#### type Custom

```go
type Custom func(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error
```

Custom is an atom issuing custom replay operations to the replay builder b upon
Replay().

#### func (Custom) API

```go
func (Custom) API() gfxapi.API
```
atom.Atom compliance

#### func (Custom) Class

```go
func (Custom) Class() binary.Class
```

#### func (Custom) Flags

```go
func (Custom) Flags() atom.Flags
```

#### func (Custom) Mutate

```go
func (Custom) Mutate(s *gfxapi.State, d database.Database, l log.Logger) error
```

#### func (Custom) Observations

```go
func (Custom) Observations() *atom.Observations
```

#### func (Custom) Replay

```go
func (c Custom) Replay(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error
```

#### func (Custom) TypeID

```go
func (Custom) TypeID() atom.TypeID
```

#### type Device

```go
type Device interface {
	// ID returns the identifier for the replay device.
	ID() service.DeviceId
	// Info returns the service Device describing the replay device.
	Info() *service.Device
	// Connect opens a connection to the replay device.
	Connect() (io.ReadWriteCloser, error)
}
```

Device is the interface for a discovered replay device.

#### type Generator

```go
type Generator interface {
	// ReplayTransforms is called when a replay pass is ready to be sent to the
	// replay device. ReplayTransforms returns an atom transform list that
	// transforms the original, unaltered atom stream into a stream configured for
	// the replay pass. The transforms should satisfy all the specified requests
	// and config.
	ReplayTransforms(
		ctx Context,
		cfg Config,
		requests []Request,
		device *service.Device,
		db database.Database,
		logger log.Logger) atom.Transforms
}
```

Generator is the interface for types that support replay generation.

#### type Image

```go
type Image struct {
	Data  []byte // The pixel data for the image
	Error error  // The error that occurred generating the image if there was one.
}
```

Image represents pixel data from an api query. The exact format of the data
depends on the query that generated it.

#### type Manager

```go
type Manager struct {
}
```

Manager is used discover replay devices and to send replay requests to those
discovered devices.

#### func  New

```go
func New(d database.Database, l log.Logger) *Manager
```
New returns a new Manager instance using the database db and logger l.

#### func (*Manager) Devices

```go
func (m *Manager) Devices() []Device
```
DeviceIDs returns the list of devices that have been discovered.

#### func (*Manager) Replay

```go
func (m *Manager) Replay(ctx *Context, cfg Config, req Request, generator Generator) error
```
Replay requests that req is to be performed on the device described by ctx,
using the capture described by ctx. Replay is asynchronous, and the replay may
take some considerable time before it is executed. Replay requests made with
configs that have equality (==) will likely be batched into the same replay
pass.

#### type QueryCallDurations

```go
type QueryCallDurations interface {
	QueryCallDurations(ctx *Context, mgr *Manager, mask service.TimingMask) <-chan CallTiming
}
```

QueryCallDurations is the interface implemented by types that can time the
duration of each call in a capture.

#### type QueryColorBuffer

```go
type QueryColorBuffer interface {
	QueryColorBuffer(ctx *Context, mgr *Manager, after atom.ID, width, height uint32, wireframe bool) <-chan Image
}
```

QueryColorBuffer is the interface implemented by types that can return the
content of the color buffer at a particular point in a capture.

#### type QueryDepthBuffer

```go
type QueryDepthBuffer interface {
	QueryDepthBuffer(ctx *Context, mgr *Manager, after atom.ID) <-chan Image
}
```

QueryDepthBufferer is the interface implemented by types that can return the
content of the depth buffer at a particular point in a capture.

#### type Replayer

```go
type Replayer interface {
	// Replay issues replay operations to the replay builder b for the given atom
	// with identifier i, and graphics API state s. If the replay action will
	// have an effect on the graphics driver state, then the call to Replay should
	// also apply the corresponding changes to the state s.
	Replay(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error
}
```

Replayer is the interface that wraps the basic Replay method.

#### type Request

```go
type Request interface{}
```

Request is a user-defined type that holds information relevant to a single
replay request. An example Request would be one that informs ReplayTransforms to
insert a postback of the currently bound render-target content at a specific
atom.
