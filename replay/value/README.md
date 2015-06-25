# value
--
    import "android.googlesource.com/platform/tools/gpu/replay/value"

Package value contains the value types used by the replay virtual machine.

Each numerical and boolean value type is backed by a corresponding primitive Go
type for convenience of construction and usage. Pointer values can belong to
various different address spaces, and for compatibility with both 32 and 64 bit
architectures, are all backed by uint64.

## Usage

#### type AbsolutePointer

```go
type AbsolutePointer uint64
```

AbsolutePointer is a pointer in the absolute address-space that will not be
altered before being passed to the protocol.

#### func (AbsolutePointer) Get

```go
func (p AbsolutePointer) Get(PointerResolver) (protocol.Type, uint64)
```
Get returns TypeAbsolutePointer and the uint64 value of the absolute pointer.

#### func (AbsolutePointer) IsValid

```go
func (p AbsolutePointer) IsValid() bool
```
IsValid returns true for all absolute pointers.

#### func (AbsolutePointer) Offset

```go
func (p AbsolutePointer) Offset(offset uint64) Pointer
```
Offset returns the sum of the pointer with offset.

#### type Bool

```go
type Bool bool
```

Bool is a Value of type TypeBool.

#### func (Bool) Get

```go
func (v Bool) Get(PointerResolver) (protocol.Type, uint64)
```
Get returns TypeBool and 1 if the Bool is true, otherwise 0.

#### type ConstantPointer

```go
type ConstantPointer uint64
```

ConstantPointer is a pointer in the constant address-space that will not be
altered before being passed to the protocol.

#### func (ConstantPointer) Get

```go
func (p ConstantPointer) Get(PointerResolver) (protocol.Type, uint64)
```
Get returns TypeConstantPointer and the uint64 value of the pointer in constant
address-space.

#### func (ConstantPointer) IsValid

```go
func (p ConstantPointer) IsValid() bool
```
IsValid returns true.

#### func (ConstantPointer) Offset

```go
func (p ConstantPointer) Offset(offset uint64) Pointer
```
Offset returns the sum of the pointer with offset.

#### type F32

```go
type F32 float32
```

F32 is a Value of type TypeFloat.

#### func (F32) Get

```go
func (v F32) Get(PointerResolver) (protocol.Type, uint64)
```
Get returns TypeFloat and the IEEE 754 representation of the value packed into
the low part of a uint64.

#### type F64

```go
type F64 float64
```

F64 is a Value of type TypeDouble.

#### func (F64) Get

```go
func (v F64) Get(PointerResolver) (protocol.Type, uint64)
```
Get returns TypeDouble and the IEEE 754 representation of the value packed into
a uint64.

#### type Pointer

```go
type Pointer interface {
	Value

	// Add returns the Pointer offset by v.
	Offset(v uint64) Pointer

	// IsValid returns true if the pointer is within acceptable ranges.
	IsValid() bool
}
```

Pointer is a pointer-typed Value.

#### type PointerResolver

```go
type PointerResolver interface {
	// TranslateTemporaryPointer returns the temporary address-space pointer ptr
	// translated to volatile address-space.
	TranslateTemporaryPointer(ptr uint64) uint64

	// TranslateRemappedPointer returns the capture-observed pointer
	// translated to volatile or absolute address-space.
	TranslateRemappedPointer(ptr uint64) (protocol.Type, uint64)
}
```

PointerResolver is used to translate pointers into the volatile address-space.

#### type RemappedPointer

```go
type RemappedPointer uint64
```

RemappedPointer is a pointer that was observed at capture time. Pointers of this
type are remapped to an equivalent volatile address-space pointer, or absolute
address-space pointer before being passed to the protocol.

#### func (RemappedPointer) Get

```go
func (p RemappedPointer) Get(r PointerResolver) (protocol.Type, uint64)
```
Get returns the pointer type and the pointer translated to either an equivalent
volatile address-space pointer or absolute pointer.

#### func (RemappedPointer) IsValid

```go
func (p RemappedPointer) IsValid() bool
```
IsValid returns true if the pointer considered valid. Currently this is a test
for the pointer being greater than 0x1000 as low addresses are likely to be a
wrong interpretation of the value. This may change in the future.

#### func (RemappedPointer) Offset

```go
func (p RemappedPointer) Offset(offset uint64) Pointer
```
Offset returns the sum of the pointer with offset.

#### type S16

```go
type S16 int16
```

S16 is a Value of type TypeInt16.

#### func (S16) Get

```go
func (v S16) Get(PointerResolver) (protocol.Type, uint64)
```
Get returns TypeInt16 and the value sign-extended to a uint64.

#### type S32

```go
type S32 int32
```

S32 is a Value of type TypeInt32.

#### func (S32) Get

```go
func (v S32) Get(PointerResolver) (protocol.Type, uint64)
```
Get returns TypeInt32 and the value sign-extended to a uint64.

#### type S64

```go
type S64 int64
```

S64 is a Value of type TypeInt64.

#### func (S64) Get

```go
func (v S64) Get(PointerResolver) (protocol.Type, uint64)
```
Get returns TypeInt64 and the value reinterpreted as a uint64.

#### type S8

```go
type S8 int8
```

S8 is a Value of type TypeInt8.

#### func (S8) Get

```go
func (v S8) Get(PointerResolver) (protocol.Type, uint64)
```
Get returns TypeInt8 and the value sign-extended to a uint64.

#### type U16

```go
type U16 uint16
```

U16 is a Value of type TypeUint16.

#### func (U16) Get

```go
func (v U16) Get(PointerResolver) (protocol.Type, uint64)
```
Get returns TypeUint16 and the value zero-extended to a uint64.

#### type U32

```go
type U32 uint32
```

U32 is a Value of type TypeUint32.

#### func (U32) Get

```go
func (v U32) Get(PointerResolver) (protocol.Type, uint64)
```
Get returns TypeUint32 and the value zero-extended to a uint64.

#### type U64

```go
type U64 uint64
```

U64 is a Value of type TypeUint64.

#### func (U64) Get

```go
func (v U64) Get(PointerResolver) (protocol.Type, uint64)
```
Get returns TypeUint64 the value zero-extended to a uint64.

#### type U8

```go
type U8 uint8
```

U8 is a Value of type TypeUint8.

#### func (U8) Get

```go
func (v U8) Get(PointerResolver) (protocol.Type, uint64)
```
Get returns TypeUint8 and the value zero-extended to a uint64.

#### type Value

```go
type Value interface {
	// Get returns the protocol type and the bit-representation of the value.
	// For example a boolean value would either be 0 or 1, a uint32 value would be
	// zero-extended, a float64 would be the IEEE 754 representation
	// reinterpreted as a uint64.
	Get(PointerResolver) (protocol.Type, uint64)
}
```

Value is the interface for all values to be passed either in opcodes or constant
memory to the replay virtual machine.

#### type VolatilePointer

```go
type VolatilePointer uint64
```

VolatilePointer is a pointer to the volatile address-space. Unlike
RemappedPointer, there is no remapping.

#### func (VolatilePointer) Get

```go
func (p VolatilePointer) Get(PointerResolver) (protocol.Type, uint64)
```
Get returns TypeVolatilePointer and the uint64 value of the pointer in volatile
address-space.

#### func (VolatilePointer) IsValid

```go
func (p VolatilePointer) IsValid() bool
```
IsValid returns true.

#### func (VolatilePointer) Offset

```go
func (p VolatilePointer) Offset(offset uint64) Pointer
```
Offset returns the sum of the pointer with offset.

#### type VolatileTemporaryPointer

```go
type VolatileTemporaryPointer uint64
```

VolatileTemporaryPointer is a pointer to in temporary address-space. The
temporary address-space sits within a reserved area of the the volatile address
space and its offset is calculated dynamically.

#### func (VolatileTemporaryPointer) Get

```go
func (p VolatileTemporaryPointer) Get(r PointerResolver) (protocol.Type, uint64)
```
Get returns TypeVolatilePointer and the dynamically calculated offset of the
temporary pointer within volatile address-space.

#### func (VolatileTemporaryPointer) IsValid

```go
func (p VolatileTemporaryPointer) IsValid() bool
```
IsValid returns true.

#### func (VolatileTemporaryPointer) Offset

```go
func (p VolatileTemporaryPointer) Offset(offset uint64) Pointer
```
Offset returns the sum of the pointer with offset.
