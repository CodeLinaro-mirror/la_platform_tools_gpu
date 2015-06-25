# objects
--
    import "android.googlesource.com/platform/tools/gpu/binary/objects"

Package objects holds standard common binary object implementations.

## Usage

```go
var (
	NilClass = &binaryClassNil{}
)
```

```go
var (
	TerminatorID = binary.ID{0x01}
)
```

#### type Terminator

```go
type Terminator struct{}
```

Terminator is an object with no payload who's purpose is to mark the end of a an
object stream.

#### func (*Terminator) Class

```go
func (*Terminator) Class() binary.Class
```
