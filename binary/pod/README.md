# pod
--
    import "android.googlesource.com/platform/tools/gpu/binary/pod"

Package pod contains Object wrappers for Plain-Old-Data types.

## Usage

```go
var Namespace = registry.NewNamespace()
```

#### func  Unwrap

```go
func Unwrap(o binary.Object) interface{}
```
Unwrap returns the POD value wrapped in the binary.Object. If v is not a POD
type, then the function returns o unaltered.

#### func  Wrap

```go
func Wrap(v interface{}) binary.Object
```
Wrap returns v wrapped by a struct implementing binary.Object. If v already
conforms to binary.Object then v is returned. If v is not a POD type, then the
function returns nil.

#### type Bool

```go
type Bool struct {
	binary.Generate
	Value bool
}
```

Bool wraps a bool value into a binary.Object.

#### func (*Bool) Class

```go
func (*Bool) Class() binary.Class
```

#### type Float32

```go
type Float32 struct {
	binary.Generate
	Value float32
}
```

Float32 wraps a float32 value into a binary.Object.

#### func (*Float32) Class

```go
func (*Float32) Class() binary.Class
```

#### type Float64

```go
type Float64 struct {
	binary.Generate
	Value float64
}
```

Float64 wraps a float64 value into a binary.Object.

#### func (*Float64) Class

```go
func (*Float64) Class() binary.Class
```

#### type Int16

```go
type Int16 struct {
	binary.Generate
	Value int16
}
```

Int16 wraps a int16 value into a binary.Object.

#### func (*Int16) Class

```go
func (*Int16) Class() binary.Class
```

#### type Int32

```go
type Int32 struct {
	binary.Generate
	Value int32
}
```

Int32 wraps a int32 value into a binary.Object.

#### func (*Int32) Class

```go
func (*Int32) Class() binary.Class
```

#### type Int64

```go
type Int64 struct {
	binary.Generate
	Value int64
}
```

Int64 wraps a int64 value into a binary.Object.

#### func (*Int64) Class

```go
func (*Int64) Class() binary.Class
```

#### type Int8

```go
type Int8 struct {
	binary.Generate
	Value int8
}
```

Int8 wraps a int8 value into a binary.Object.

#### func (*Int8) Class

```go
func (*Int8) Class() binary.Class
```

#### type String

```go
type String struct {
	binary.Generate
	Value string
}
```

String wraps a string value into a binary.Object.

#### func (*String) Class

```go
func (*String) Class() binary.Class
```

#### type Uint16

```go
type Uint16 struct {
	binary.Generate
	Value uint16
}
```

Uint16 wraps a uint16 value into a binary.Object.

#### func (*Uint16) Class

```go
func (*Uint16) Class() binary.Class
```

#### type Uint32

```go
type Uint32 struct {
	binary.Generate
	Value uint32
}
```

Uint32 wraps a uint32 value into a binary.Object.

#### func (*Uint32) Class

```go
func (*Uint32) Class() binary.Class
```

#### type Uint64

```go
type Uint64 struct {
	binary.Generate
	Value uint64
}
```

Uint64 wraps a uint64 value into a binary.Object.

#### func (*Uint64) Class

```go
func (*Uint64) Class() binary.Class
```

#### type Uint8

```go
type Uint8 struct {
	binary.Generate
	Value uint8
}
```

Uint8 wraps a uint8 value into a binary.Object.

#### func (*Uint8) Class

```go
func (*Uint8) Class() binary.Class
```
