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
func (p AbsolutePointer) Get(PointerResolver) (uint64, error)
```
Get returns the uint64 value of the absolute pointer.

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

#### func (AbsolutePointer) Type

```go
func (p AbsolutePointer) Type() protocol.Type
```
Type returns TypeAbsolutePointer.

#### type Bool

```go
type Bool bool
```

Bool is a Value of type TypeBool.

#### func (Bool) Get

```go
func (v Bool) Get(PointerResolver) (uint64, error)
```
Get returns 1 if the Bool is true, otherwise 0.

#### func (Bool) Type

```go
func (v Bool) Type() protocol.Type
```
Type returns TypeBool.

#### type ConstantPointer

```go
type ConstantPointer uint64
```

ConstantPointer is a pointer in the constant address-space that will not be
altered before being passed to the protocol.

#### func (ConstantPointer) Get

```go
func (p ConstantPointer) Get(PointerResolver) (uint64, error)
```
Get returns the uint64 value of the pointer in constant address-space.

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

#### func (ConstantPointer) Type

```go
func (p ConstantPointer) Type() protocol.Type
```
Type returns TypeConstantPointer.

#### type F32

```go
type F32 float32
```

F32 is a Value of type TypeFloat.

#### func (F32) Get

```go
func (v F32) Get(PointerResolver) (uint64, error)
```
Get returns the IEEE 754 representation of the value packed into the low part of
a uint64.

#### func (F32) Type

```go
func (v F32) Type() protocol.Type
```
Type returns TypeFloat.

#### type F64

```go
type F64 float64
```

F64 is a Value of type TypeDouble.

#### func (F64) Get

```go
func (v F64) Get(PointerResolver) (uint64, error)
```
Get returns the IEEE 754 representation of the value packed into a uint64.

#### func (F64) Type

```go
func (v F64) Type() protocol.Type
```
Type returns TypeDouble.

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
	TranslateTemporaryPointer(ptr uint64) (uint64, error)

	// TranslateCapturePointer returns the capture-observed pointer
	// translated to volatile address-space.
	TranslateCapturePointer(ptr uint64) (uint64, error)
}
```

PointerResolver is used to translate pointers into the volatile address-space.

#### type S16

```go
type S16 int16
```

S16 is a Value of type TypeInt16.

#### func (S16) Get

```go
func (v S16) Get(PointerResolver) (uint64, error)
```
Get returns the value sign-extended to a uint64.

#### func (S16) Type

```go
func (v S16) Type() protocol.Type
```
Type returns TypeInt16.

#### type S32

```go
type S32 int32
```

S32 is a Value of type TypeInt32.

#### func (S32) Get

```go
func (v S32) Get(PointerResolver) (uint64, error)
```
Get returns the value sign-extended to a uint64.

#### func (S32) Type

```go
func (v S32) Type() protocol.Type
```
Type returns TypeInt32.

#### type S64

```go
type S64 int64
```

S64 is a Value of type TypeInt64.

#### func (S64) Get

```go
func (v S64) Get(PointerResolver) (uint64, error)
```
Get returns the value reinterpreted as a uint64.

#### func (S64) Type

```go
func (v S64) Type() protocol.Type
```
Type returns TypeInt64.

#### type S8

```go
type S8 int8
```

S8 is a Value of type TypeInt8.

#### func (S8) Get

```go
func (v S8) Get(PointerResolver) (uint64, error)
```
Get returns the value sign-extended to a uint64.

#### func (S8) Type

```go
func (v S8) Type() protocol.Type
```
Type returns TypeInt8.

#### type U16

```go
type U16 uint16
```

U16 is a Value of type TypeUint16.

#### func (U16) Get

```go
func (v U16) Get(PointerResolver) (uint64, error)
```
Get returns the value zero-extended to a uint64.

#### func (U16) Type

```go
func (v U16) Type() protocol.Type
```
Type returns TypeUint16.

#### type U32

```go
type U32 uint32
```

U32 is a Value of type TypeUint32.

#### func (U32) Get

```go
func (v U32) Get(PointerResolver) (uint64, error)
```
Get returns the value zero-extended to a uint64.

#### func (U32) Type

```go
func (v U32) Type() protocol.Type
```
Type returns TypeUint32.

#### type U64

```go
type U64 uint64
```

U64 is a Value of type TypeUint64.

#### func (U64) Get

```go
func (v U64) Get(PointerResolver) (uint64, error)
```
Get returns the value zero-extended to a uint64.

#### func (U64) Type

```go
func (v U64) Type() protocol.Type
```
Type returns TypeUint64.

#### type U8

```go
type U8 uint8
```

U8 is a Value of type TypeUint8.

#### func (U8) Get

```go
func (v U8) Get(PointerResolver) (uint64, error)
```
Get returns the value zero-extended to a uint64.

#### func (U8) Type

```go
func (v U8) Type() protocol.Type
```
Type returns TypeUint8.

#### type Value

```go
type Value interface {
	// Type returns the virtual-machine type of the Value.
	Type() protocol.Type

	// Get returns the bit-representation of the value. For example a boolean
	// value would either be 0 or 1, a uint32 value would be zero-extended, a
	// float64 would be the IEEE 754 representation reinterpreted as a uint64.
	Get(PointerResolver) (uint64, error)
}
```

Value is the interface for all values to be passed either in opcodes or constant
memory to the replay virtual machine.

#### type VolatileCapturePointer

```go
type VolatileCapturePointer uint64
```

VolatileCapturePointer is a pointer that was observed at capture time. Pointers
of this type are remapped to an equivalent volatile address-space pointer before
being passed to the protocol.

#### func (VolatileCapturePointer) Get

```go
func (p VolatileCapturePointer) Get(r PointerResolver) (uint64, error)
```
Get returns the observed pointer translated to an equivalent volatile
address-space pointer.

#### func (VolatileCapturePointer) IsValid

```go
func (p VolatileCapturePointer) IsValid() bool
```
IsValid returns true if the pointer considered valid. Currently this is a test
for the pointer being greater than 0x10000 as low addresses are likely to be a
wrong interpretation of the value. This may change in the future.

#### func (VolatileCapturePointer) Offset

```go
func (p VolatileCapturePointer) Offset(offset uint64) Pointer
```
Offset returns the sum of the pointer with offset.

#### func (VolatileCapturePointer) Type

```go
func (p VolatileCapturePointer) Type() protocol.Type
```
Type returns TypeVolatilePointer.

#### type VolatilePointer

```go
type VolatilePointer uint64
```

VolatilePointer is a pointer to the volatile address-space. Unlike
VolatileCapturePointer, there is no remapping.

#### func (VolatilePointer) Get

```go
func (p VolatilePointer) Get(PointerResolver) (uint64, error)
```
Get returns the uint64 value of the pointer in volatile address-space.

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

#### func (VolatilePointer) Type

```go
func (p VolatilePointer) Type() protocol.Type
```
Type returns TypeVolatilePointer.

#### type VolatileTemporaryPointer

```go
type VolatileTemporaryPointer uint64
```

VolatileTemporaryPointer is a pointer to in temporary address-space. The
temporary address-space sits within a reserved area of the the volatile address
space and its offset is calculated dynamically.

#### func (VolatileTemporaryPointer) Get

```go
func (p VolatileTemporaryPointer) Get(r PointerResolver) (uint64, error)
```
Get returns the dynamically calculated offset of the temporary pointer within
volatile address-space.

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

#### func (VolatileTemporaryPointer) Type

```go
func (p VolatileTemporaryPointer) Type() protocol.Type
```
Type returns TypeVolatilePointer.
