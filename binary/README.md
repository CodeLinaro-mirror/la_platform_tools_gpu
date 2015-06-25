# binary
--
    import "android.googlesource.com/platform/tools/gpu/binary"

Package binary implements encoding and decoding of various primitive data types
to and from a binary stream. The package holds BitStream for packing and
unpacking sequences of bits, Float16 for dealing with 16 bit floating- point
values and Reader/Writer for encoding and decoding various value types to a
binary stream. There is also the higher level Encoder/Decoder that can be used
for serializing object hierarchies.

binary.Reader and binary.Writer provide a symmetrical pair of methods for
encoding and decoding various data types to a binary stream. For performance
reasons, each data type has a separate method for encoding and decoding rather
than having a single pair of methods encoding and decoding boxed values in an
interface{}.

binary.Encoder and binary.Decoder extend the binary.Reader and binary.Writer
interfaces by also providing a symmetrical pair of methods for encoding and
decoding object types.

## Usage

```go
const IDSize = 20
```
IDSize is the size of an ID.

#### func  ReadInt

```go
func ReadInt(r Reader, bits int) (int64, error)
```
ReadInt reads a signed integer of either 8, 16, 32 or 64 bits from r, returning
the result as a int64.

#### func  ReadUint

```go
func ReadUint(r Reader, bits int) (uint64, error)
```
ReadUint reads an unsigned integer of either 8, 16, 32 or 64 bits from r,
returning the result as a uint64.

#### func  WriteInt

```go
func WriteInt(w Writer, bits int, v int64) error
```
WriteInt writes the signed integer v of either 8, 16, 32 or 64 bits to w.

#### func  WriteUint

```go
func WriteUint(w Writer, bits int, v uint64) error
```
WriteUint writes the unsigned integer v of either 8, 16, 32 or 64 bits to w.

#### type BitStream

```go
type BitStream struct {
	Data     []byte // The byte slice containing the bits
	ReadPos  uint32 // The current read offset from the start of the Data slice (in bits)
	WritePos uint32 // The current write offset from the start of the Data slice (in bits)
}
```

BitStream provides methods for reading and writing bits to a slice of bytes.
Bits are packed in a least-significant-bit to most-significant-bit order.

#### func (*BitStream) Read

```go
func (s *BitStream) Read(count uint32) uint32
```
Read reads the specified number of bits from the BitStream, increamenting the
ReadPos by the specified number of bits and returning the bits packed into a
uint32. The bits are packed into the uint32 from LSB to MSB.

#### func (*BitStream) ReadBit

```go
func (s *BitStream) ReadBit() uint32
```
ReadBit reads a single bit from the BitStream, incrementing ReadPos by one.

#### func (*BitStream) Write

```go
func (s *BitStream) Write(bits, count uint32)
```
Write writes the specified number of bits from the packed uint32, increamenting
the WritePos by the specified number of bits. The bits are read from the uint32
from LSB to MSB.

#### func (*BitStream) WriteBit

```go
func (s *BitStream) WriteBit(bit uint32)
```
WriteBit writes a single bit to the BitStream, incrementing WritePos by one.

#### type Class

```go
type Class interface {
	// ID should be a sha1 has of the types signature, such that
	// no two classes generate the same ID, and any change to the types name or
	// fields causes it's id to change.
	ID() ID

	// New can be used to build a new default initialized instance of the type.
	New() Object

	// Encode writes the supplied object to the supplied Encoder.
	// The object must be a type the Class understands, the implementation is
	// allowed to panic if it is not.
	Encode(Encoder, Object) error

	// Decode reads a single object from the supplied Decoder.
	Decode(Decoder) (Object, error)

	// DecodeTo reads into the supplied object from the supplied Decoder.
	// The object must be a type the Class understands, the implementation is
	// allowed to panic if it is not.
	DecodeTo(Decoder, Object) error

	// Skip moves over a single object from the supplied Decoder.
	// This must skip the same data that Decode would have read.
	Skip(Decoder) error
}
```

Class represents a struct type in the binary registry.

#### type Decoder

```go
type Decoder interface {
	Reader
	// ID decodes a binary.ID from the stream.
	ID() (ID, error)
	// SkipID skips over a binary.ID in the stream.
	SkipID() error
	// Value decodes an Object from the stream.
	Value(Object) error
	// SkipValue must skip the same data that a call to Value would read.
	// The value may be a typed nil.
	SkipValue(Object) error
	// Variant decodes and returns an Object from the stream. The Class in the
	// stream must have been previously registered with binary.registry.Add.
	Variant() (Object, error)
	// SkipVariant must skip the same data that a call to Variant would read.
	SkipVariant() (ID, error)
	// Object decodes and returns an Object from the stream. Object instances
	// that were encoded multiple times may be decoded and returned as a shared,
	// single instance. The Class in the stream must have been previously
	// registered with binary.registry.Add.
	Object() (Object, error)
	// SkipObject must skip the same data that a call to Object would read.
	SkipObject() (ID, error)
}
```

Decoder extends Reader with additional methods for decoding objects.

#### type Encoder

```go
type Encoder interface {
	Writer
	// ID writes a binary.ID to the stream.
	ID(ID) error
	// Object encodes an Object with no type preamble and no sharing.
	Value(obj Object) error
	// Variant encodes an Object with no sharing. The type of obj must have
	// been previously registered with binary.registry.Add.
	Variant(obj Object) error
	// Object encodes an Object, optionally encoding objects only on the first
	// time it sees them. The type of obj must have been previously registered
	// with binary.registry.Add.
	Object(obj Object) error
}
```

Encoder extends Writer with additional methods for encoding objects.

#### type Float16

```go
type Float16 uint16
```

Float16 represents a 16-bit floating point number, containing a single sign bit,
5 exponent bits and 10 fractional bits. This corresponds to IEEE 754-2008
binary16 (or half precision float) type.

     MSB                                                                         LSB
    ╔════╦════╤════╤════╤════╤════╦════╤════╤════╤════╤════╤════╤════╤════╤════╤════╗
    ║Sign║ E₄ │ E₃ │ E₂ │ E₁ │ E₀ ║ F₉ │ F₈ │ F₇ │ F₆ │ F₅ │ F₄ │ F₃ │ F₂ │ F₁ │ F₀ ║
    ╚════╩════╧════╧════╧════╧════╩════╧════╧════╧════╧════╧════╧════╧════╧════╧════╝
    Where E is the exponent bits and F is the fractional bits.

#### func  NewFloat16

```go
func NewFloat16(f32 float32) Float16
```
NewFloat16 returns a Float16 encoding of a 32-bit floating point number.
Infinities and NaNs are encoded as such. Very large and very small numbers get
rounded to infinity and zero respectively.

#### func  NewFloat16Inf

```go
func NewFloat16Inf(sign int) Float16
```
Float16Inf returns positive infinity if sign >= 0, negative infinity if sign <
0.

#### func  NewFloat16NaN

```go
func NewFloat16NaN() Float16
```
Float16NaN returns an “not-a-number” value.

#### func (Float16) Float32

```go
func (f Float16) Float32() float32
```
Float32 returns the Float16 value expanded to a float32. Infinities and NaNs are
expanded as such.

#### func (Float16) IsInf

```go
func (f Float16) IsInf(sign int) bool
```
IsInf reports whether f is an infinity, according to sign. If sign > 0, IsInf
reports whether f is positive infinity. If sign < 0, IsInf reports whether f is
negative infinity. If sign == 0, IsInf reports whether f is either infinity.

#### func (Float16) IsNaN

```go
func (f Float16) IsNaN() bool
```
IsNaN reports whether f is an “not-a-number” value.

#### type Generate

```go
type Generate struct{}
```

Generate is used to tag structures that need an auto generated Class. The
codergen function searches packages for structs that have this type as an
anonymous field, and then automatically generates the encoding and decoding
functionality for those structures. For example, the following struct would
create the Class with methods needed to encode and decode the Name and Value
fields, and register that class. The embedding will also fully implement the
binary.Object interface, but with methods that panic. This will get overridden
with the generated Methods. This is important because it means the package is
resolvable without the generated code, which means the types can be correctly
evaluated during the generation process.

type MyNamedValue struct {

    binary.Generate
    Name  string
    Value []byte

}

#### func (Generate) Class

```go
func (Generate) Class() Class
```

#### type ID

```go
type ID [IDSize]byte
```

ID is a codeable unique identifier.

#### func  NewID

```go
func NewID(data ...[]byte) ID
```
Create a new ID that is the sha1 hash of the supplied data.

#### func  ParseID

```go
func ParseID(s string) (id ID, err error)
```
ParseID parses lowercase string s as a 20 byte hex-encoded ID.

#### func (ID) Format

```go
func (id ID) Format(f fmt.State, c rune)
```

#### func (ID) String

```go
func (id ID) String() string
```

#### func (ID) Valid

```go
func (id ID) Valid() bool
```
Valid returns true if the id is not the default value.

#### type Object

```go
type Object interface {
	// Class returns the serialize information and functionality for this type.
	// The method should be valid on a nil pointer.
	Class() Class
}
```

Object is the interface to any class that wants to be encoded/decoded.

#### type Reader

```go
type Reader interface {
	// Data reads the data bytes in their entirety.
	Data([]byte) error
	// Skip jumps past count bytes.
	Skip(count uint32) error
	// Bool decodes and returns a boolean value from the Reader.
	Bool() (bool, error)
	// Int8 decodes and returns a signed, 8 bit integer value from the Reader.
	Int8() (int8, error)
	// Uint8 decodes and returns an unsigned, 8 bit integer value from the Reader.
	Uint8() (uint8, error)
	// Int16 decodes and returns a signed, 16 bit integer value from the Reader.
	Int16() (int16, error)
	// Uint16 decodes and returns an unsigned, 16 bit integer value from the Reader.
	Uint16() (uint16, error)
	// Int32 decodes and returns a signed, 32 bit integer value from the Reader.
	Int32() (int32, error)
	// Uint32 decodes and returns an unsigned, 32 bit integer value from the Reader.
	Uint32() (uint32, error)
	// Float32 decodes and returns a 32 bit floating-point value from the Reader.
	Float32() (float32, error)
	// Int64 decodes and returns a signed, 64 bit integer value from the Reader.
	Int64() (int64, error)
	// Uint64 decodes and returns an unsigned, 64 bit integer value from the Reader.
	Uint64() (uint64, error)
	// Float64 decodes and returns a 64 bit floating-point value from the Reader.
	Float64() (float64, error)
	// String decodes and returns a string from the Reader.
	String() (string, error)
	// SkipString skips over a single string from the Reader.
	SkipString() error
}
```

Reader provides methods for decoding values.

#### type Writer

```go
type Writer interface {
	// Data writes the data bytes in their entirety.
	Data([]byte) error
	// Bool encodes a boolean value to the Writer.
	Bool(bool) error
	// Int8 encodes a signed, 8 bit integer value to the Writer.
	Int8(int8) error
	// Uint8 encodes an unsigned, 8 bit integer value to the Writer.
	Uint8(uint8) error
	// Int16 encodes a signed, 16 bit integer value to the Writer.
	Int16(int16) error
	// Uint16 encodes an unsigned, 16 bit integer value to the Writer.
	Uint16(uint16) error
	// Int32 encodes a signed, 32 bit integer value to the Writer.
	Int32(int32) error
	// Uint32 encodes an usigned, 32 bit integer value to the Writer.
	Uint32(uint32) error
	// Float32 encodes a 32 bit floating-point value to the Writer.
	Float32(float32) error
	// Int64 encodes a signed, 64 bit integer value to the Writer.
	Int64(int64) error
	// Uint64 encodes an unsigned, 64 bit integer value to the Encoders's io.Writer.
	Uint64(uint64) error
	// Float64 encodes a 64 bit floating-point value to the Writer.
	Float64(float64) error
	// String encodes a string to the Writer.
	String(string) error
}
```

Writer provides methods for encoding values.
