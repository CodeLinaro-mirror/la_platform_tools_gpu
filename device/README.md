# device
--
    import "android.googlesource.com/platform/tools/gpu/device"


## Usage

#### type Architecture

```go
type Architecture struct {
	PointerAlignment int              // The alignment in bytes of a pointer type.
	PointerSize      int              // The size in bytes of a pointer type.
	IntegerSize      int              // The size in bytes of a int or unsigned int.
	ByteOrder        endian.ByteOrder // The byte ordering for the target.
}
```

Architecture holds architecture information about a device.

#### func (Architecture) String

```go
func (a Architecture) String() string
```
