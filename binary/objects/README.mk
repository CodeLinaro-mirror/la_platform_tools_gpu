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
	TerminatorID = binary.ID{0x67, 0x56, 0x42, 0x64, 0x1b, 0xc7, 0xa0, 0xf4, 0x8d, 0xe1, 0x64, 0xc0, 0x4d, 0x22, 0x9b, 0xa8, 0x13, 0xb2, 0x70, 0xea}
)
```

#### type Terminator

```go
type Terminator struct {
	binary.Generate `id:"TerminatorID"`
}
```

Terminator is an object with no payload who's purpose is to mark the end of a an
object stream.

#### func (*Terminator) Class

```go
func (*Terminator) Class() binary.Class
```
