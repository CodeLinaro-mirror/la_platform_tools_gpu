# any
--
    import "android.googlesource.com/platform/tools/gpu/binary/any"

Package any contains Object wrappers for Plain-Old-Data types.

## Usage

```go
var Namespace = registry.NewNamespace()
```

#### func  Box

```go
func Box(v interface{}) (binary.Object, error)
```
Box returns v wrapped by a struct implementing binary.Object. If v is not
boxable then ErrUnboxable is returned.

#### func  Unbox

```go
func Unbox(o binary.Object) (interface{}, error)
```
Unbox returns the value in o wrapped by a call to Box. If o is not a boxed value
ErrNotBoxedValue is returned.

#### type Any

```go
type Any struct {
	binary.Generate
}
```

Any is the schema Type descriptor for a field who's underlying type requires
boxing and unboxing. The type is usually declared as an empty interface.

#### func (*Any) Basename

```go
func (i *Any) Basename() string
```

#### func (*Any) Class

```go
func (*Any) Class() binary.Class
```

#### func (*Any) Decode

```go
func (i *Any) Decode(d binary.Decoder) (interface{}, error)
```

#### func (*Any) Encode

```go
func (i *Any) Encode(e binary.Encoder, value interface{}) error
```

#### func (*Any) Skip

```go
func (i *Any) Skip(d binary.Decoder) error
```

#### func (*Any) String

```go
func (i *Any) String() string
```

#### func (*Any) Typename

```go
func (i *Any) Typename() string
```

#### type ErrNotBoxedValue

```go
type ErrNotBoxedValue struct {
	Object binary.Object // The object that is not a boxed value.
}
```

ErrNotBoxedValue is returned when an Object is passed to Unbox that was not
previously returned by a call to Box.

#### func (ErrNotBoxedValue) Error

```go
func (e ErrNotBoxedValue) Error() string
```
Error returns the error message.

#### type ErrUnboxable

```go
type ErrUnboxable struct {
	Value interface{} // The value that could not be encoded.
}
```

ErrUnboxable is returned when a non-boxable value type is passed to Box.

#### func (ErrUnboxable) Error

```go
func (e ErrUnboxable) Error() string
```
Error returns the error message.
