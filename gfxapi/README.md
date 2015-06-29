# gfxapi
--
    import "android.googlesource.com/platform/tools/gpu/gfxapi"

Package gfxapi exposes the shared behavior of all graphics api's.

## Usage

#### func  Register

```go
func Register(api API)
```
Register adds an api to the understood set. It is illegal to register the same
name twice.

#### type API

```go
type API interface {
	// Name returns the official name of the api.
	Name() string

	// ID returns the unique API identifier.
	ID() ID

	// GetFramebufferAttachmentSize returns the width and height of the framebuffer at the given attachment.
	GetFramebufferAttachmentSize(state *State, attachment FramebufferAttachment) (width uint32, height uint32, err error)
}
```

API is the common interface to a graphics programming api.

#### func  Find

```go
func Find(id ID) API
```
Find looks up a graphics API by identifier. If the id has not been registered,
it returns nil.

#### type FramebufferAttachment

```go
type FramebufferAttachment uint32
```

FramebufferAttachment values indicate the type of frame buffer attachment.

```go
const (
	FramebufferAttachmentColor FramebufferAttachment = iota
	FramebufferAttachmentDepth
	FramebufferAttachmentStencil
)
```

#### func (*FramebufferAttachment) Parse

```go
func (v *FramebufferAttachment) Parse(s string) error
```

#### func (FramebufferAttachment) String

```go
func (v FramebufferAttachment) String() string
```

#### type ID

```go
type ID binary.ID
```

ID is an API identifier

#### func (ID) Valid

```go
func (i ID) Valid() bool
```
Valid returns true if the id is not the default zero value.

#### type State

```go
type State struct {
	binary.Object

	// Architecture holds information about the device architecture that was used
	// to create the capture.
	Architecture device.Architecture

	// Memory holds the memory state of the application.
	Memory map[memory.PoolID]*memory.Pool

	// NextPoolID hold the identifier of the next Pool to be created.
	NextPoolID memory.PoolID

	// APIs holds the per-API context states.
	APIs map[API]binary.Object
}
```

State represents the graphics state across all contexts.

#### func  NewState

```go
func NewState() *State
```

#### func (State) MemoryDecoder

```go
func (st State) MemoryDecoder(s memory.Slice, d database.Database, l log.Logger) binary.Decoder
```
MemoryDecoder returns a flat decoder backed by an endian reader that uses the
byte-order of the capture device to decode from the slice s.

#### func (State) MemoryEncoder

```go
func (st State) MemoryEncoder(p *memory.Pool, rng memory.Range) binary.Encoder
```
MemoryEncoder returns a flat encoder backed by an endian reader that uses the
byte-order of the capture device to encode to the pool p, for the range rng.

#### func (State) String

```go
func (s State) String() string
```
