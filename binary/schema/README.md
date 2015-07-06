# schema
--
    import "android.googlesource.com/platform/tools/gpu/binary/schema"

Package schema implements rtti for the binary system.

## Usage

```go
var Namespace = registry.NewNamespace()
```

#### func  Underlying

```go
func Underlying(v interface{}) interface{}
```
Underlying traverses the single, anonymous fields nested in v, returning the
deepest-nested value that is not an object or does not have a single, anonymous
field. If v is not an Object or does not have a single, anonymous field then v
is returned.

#### type Array

```go
type Array struct {
	binary.Generate
	Alias     string // The alias this array type was given, if present
	ValueType Type   // The value type stored in the array
	Size      uint32 // The fixed size of the array
}
```

Array is the Type descriptor for fixed size buffers of known type.

#### func (*Array) Basename

```go
func (a *Array) Basename() string
```

#### func (*Array) Class

```go
func (*Array) Class() binary.Class
```

#### func (*Array) Decode

```go
func (a *Array) Decode(d binary.Decoder) (interface{}, error)
```

#### func (*Array) Encode

```go
func (a *Array) Encode(e binary.Encoder, value interface{}) error
```

#### func (*Array) Skip

```go
func (a *Array) Skip(d binary.Decoder) error
```

#### func (*Array) String

```go
func (a *Array) String() string
```

#### func (*Array) Typename

```go
func (a *Array) Typename() string
```

#### type Class

```go
type Class struct {
	binary.Generate
	TypeID   binary.ID       // The unique type identifier for the Object.
	Package  string          // The package that declared the struct.
	Name     string          // The simple name of the Object.
	Fields   FieldList       // Descriptions of the fields of the Object.
	Metadata []binary.Object // The metadata for the class.
}
```

Class represents an encodable object type with a type ID.

#### func  Lookup

```go
func Lookup(id binary.ID) *Class
```
Lookup looks up a Class by the given type id. If there is no match, it will
return nil.

#### func  Of

```go
func Of(class binary.Class) *Class
```
Returns the schema class for a binary class, if it has one.

#### func (*Class) Class

```go
func (*Class) Class() binary.Class
```

#### func (*Class) Decode

```go
func (c *Class) Decode(d binary.Decoder) (binary.Object, error)
```

#### func (*Class) DecodeTo

```go
func (c *Class) DecodeTo(d binary.Decoder, object binary.Object) error
```

#### func (*Class) Encode

```go
func (c *Class) Encode(e binary.Encoder, object binary.Object) error
```

#### func (*Class) ID

```go
func (c *Class) ID() binary.ID
```

#### func (*Class) New

```go
func (c *Class) New() binary.Object
```

#### func (*Class) Skip

```go
func (c *Class) Skip(d binary.Decoder) error
```

#### type Constant

```go
type Constant struct {
	binary.Generate
	Name  string
	Value interface{}
}
```


#### func (*Constant) Class

```go
func (*Constant) Class() binary.Class
```

#### type ConstantSet

```go
type ConstantSet struct {
	binary.Generate
	Type    Type       // The type of the constant.
	Entries []Constant // The constant values
}
```


#### func (*ConstantSet) Class

```go
func (*ConstantSet) Class() binary.Class
```

#### func (*ConstantSet) Len

```go
func (s *ConstantSet) Len() int
```

#### func (*ConstantSet) Less

```go
func (s *ConstantSet) Less(i, j int) bool
```

#### func (*ConstantSet) Swap

```go
func (s *ConstantSet) Swap(i, j int)
```

#### type Constants

```go
type Constants []ConstantSet
```


#### func (*Constants) Add

```go
func (c *Constants) Add(t Type, v Constant)
```

#### func (Constants) Len

```go
func (c Constants) Len() int
```

#### func (Constants) Less

```go
func (c Constants) Less(i, j int) bool
```

#### func (Constants) Swap

```go
func (c Constants) Swap(i, j int)
```

#### type Field

```go
type Field struct {
	binary.Generate
	Declared string // The name of the field.
	Type     Type   // The type stored in the field.
}
```

Field represents a name/type pair for a field in an Object.

#### func (*Field) Class

```go
func (*Field) Class() binary.Class
```

#### func (Field) Name

```go
func (f Field) Name() string
```

#### type FieldList

```go
type FieldList []Field
```

FieldList is a slice of fields.

#### func (FieldList) Find

```go
func (l FieldList) Find(name string) int
```
Find searches the field list of the field with the specified name, returning the
index of the field if found, otherwise -1.

#### type Interface

```go
type Interface struct {
	binary.Generate
	Name string // The simple name of the type.
}
```

Interface is the Type descriptor for a field who's underlying type is dynamic.

#### func (*Interface) Basename

```go
func (i *Interface) Basename() string
```

#### func (*Interface) Class

```go
func (*Interface) Class() binary.Class
```

#### func (*Interface) Decode

```go
func (i *Interface) Decode(d binary.Decoder) (interface{}, error)
```

#### func (*Interface) Encode

```go
func (i *Interface) Encode(e binary.Encoder, value interface{}) error
```

#### func (*Interface) Skip

```go
func (i *Interface) Skip(d binary.Decoder) error
```

#### func (*Interface) String

```go
func (i *Interface) String() string
```

#### func (*Interface) Typename

```go
func (i *Interface) Typename() string
```

#### type Map

```go
type Map struct {
	binary.Generate
	Alias     string // The alias this array type was given, if present
	KeyType   Type   // The key type used.
	ValueType Type   // The value type stored in the map.
}
```

Map is the Type descriptor for key/value stores.

#### func (*Map) Basename

```go
func (m *Map) Basename() string
```

#### func (*Map) Class

```go
func (*Map) Class() binary.Class
```

#### func (*Map) Decode

```go
func (m *Map) Decode(d binary.Decoder) (interface{}, error)
```

#### func (*Map) Encode

```go
func (m *Map) Encode(e binary.Encoder, value interface{}) error
```

#### func (*Map) Skip

```go
func (m *Map) Skip(d binary.Decoder) error
```

#### func (*Map) String

```go
func (m *Map) String() string
```

#### func (*Map) Typename

```go
func (m *Map) Typename() string
```

#### type Method

```go
type Method int
```

Method denotes the encoding/decoding method a primitive type will use.

```go
const (
	ID Method = iota
	Bool
	Int8
	Uint8
	Int16
	Uint16
	Int32
	Uint32
	Int64
	Uint64
	Float32
	Float64
	String
)
```

#### func  ParseMethod

```go
func ParseMethod(s string) (Method, error)
```
This will convert a string to a Method, or return an error if the string was not
a valid method name.

#### func (Method) Skippable

```go
func (m Method) Skippable() bool
```
Skippable returns true if the method has a complimentary skip method on the
decoder interface. If this is not true, the normal decoding method is used
during skipping.

#### func (Method) String

```go
func (m Method) String() string
```

#### type Object

```go
type Object struct {
	Type   *Class
	Fields []interface{}
}
```

Object is an instance of a Class.

#### func (*Object) Base

```go
func (o *Object) Base() interface{}
```
Base returns the value of the single, anonymous field of o. If o does not have a
single anonymous field, then Base returns nil.

#### func (*Object) Class

```go
func (o *Object) Class() binary.Class
```
Class implements binary.Object using the schema system to do the encoding and
decoding of fields.

#### type Pointer

```go
type Pointer struct {
	binary.Generate
	Type Type // The pointed to type.
}
```

Pointer is the Type descriptor for pointers.

#### func (*Pointer) Basename

```go
func (p *Pointer) Basename() string
```

#### func (*Pointer) Class

```go
func (*Pointer) Class() binary.Class
```

#### func (*Pointer) Decode

```go
func (p *Pointer) Decode(d binary.Decoder) (interface{}, error)
```

#### func (*Pointer) Encode

```go
func (p *Pointer) Encode(e binary.Encoder, value interface{}) error
```

#### func (*Pointer) Skip

```go
func (p *Pointer) Skip(d binary.Decoder) error
```

#### func (*Pointer) String

```go
func (p *Pointer) String() string
```

#### func (*Pointer) Typename

```go
func (p *Pointer) Typename() string
```

#### type Primitive

```go
type Primitive struct {
	binary.Generate
	Name   string // The simple name of the type.
	Method Method // The enocde/decode method to use.
}
```

Primitive is the kind for primitive types with corresponding direct methods on
Encoder and Decoder

#### func (*Primitive) Basename

```go
func (p *Primitive) Basename() string
```

#### func (*Primitive) Class

```go
func (*Primitive) Class() binary.Class
```

#### func (*Primitive) Decode

```go
func (p *Primitive) Decode(d binary.Decoder) (interface{}, error)
```

#### func (*Primitive) Encode

```go
func (p *Primitive) Encode(e binary.Encoder, value interface{}) error
```

#### func (*Primitive) Native

```go
func (p *Primitive) Native() string
```
Native returns the go native type name for this primitive.

#### func (*Primitive) Skip

```go
func (p *Primitive) Skip(d binary.Decoder) error
```
Implements binary.Class

#### func (*Primitive) String

```go
func (p *Primitive) String() string
```

#### func (*Primitive) Typename

```go
func (p *Primitive) Typename() string
```

#### type Slice

```go
type Slice struct {
	binary.Generate
	Alias     string // The alias this array type was given, if present
	ValueType Type   // The value type stored in the slice.
}
```

Slice is the Type descriptor for dynamically sized buffers of known type,
encoded with a preceding count.

#### func (*Slice) Basename

```go
func (s *Slice) Basename() string
```

#### func (*Slice) Class

```go
func (*Slice) Class() binary.Class
```

#### func (*Slice) Decode

```go
func (s *Slice) Decode(d binary.Decoder) (interface{}, error)
```

#### func (*Slice) Encode

```go
func (s *Slice) Encode(e binary.Encoder, value interface{}) error
```

#### func (*Slice) Skip

```go
func (s *Slice) Skip(d binary.Decoder) error
```

#### func (*Slice) String

```go
func (s *Slice) String() string
```

#### func (*Slice) Typename

```go
func (s *Slice) Typename() string
```

#### type Stream

```go
type Stream struct {
	binary.Generate
	Alias     string // The alias this array type was given, if present
	ValueType Type   // The value type stored in the stream.
}
```

Stream is the Type descriptor for dynamically sized streams of a known type,
where the stream is terminated with a speical token rather than prefixed by a
count.

#### func (*Stream) Basename

```go
func (s *Stream) Basename() string
```

#### func (*Stream) Class

```go
func (*Stream) Class() binary.Class
```

#### func (*Stream) Decode

```go
func (s *Stream) Decode(d binary.Decoder) (interface{}, error)
```

#### func (*Stream) Encode

```go
func (s *Stream) Encode(e binary.Encoder, value interface{}) error
```

#### func (*Stream) Skip

```go
func (s *Stream) Skip(d binary.Decoder) error
```

#### func (*Stream) String

```go
func (s *Stream) String() string
```

#### func (*Stream) Typename

```go
func (s *Stream) Typename() string
```

#### type Struct

```go
type Struct struct {
	binary.Generate
	Name string    // The simple name of the type.
	ID   binary.ID // The unique type identifier for the Object.
}
```

Struct is the Type descriptor for an binary.Object typed value.

#### func (*Struct) Basename

```go
func (s *Struct) Basename() string
```

#### func (*Struct) Class

```go
func (*Struct) Class() binary.Class
```

#### func (*Struct) Decode

```go
func (s *Struct) Decode(d binary.Decoder) (interface{}, error)
```

#### func (*Struct) Encode

```go
func (s *Struct) Encode(e binary.Encoder, value interface{}) error
```

#### func (*Struct) Skip

```go
func (s *Struct) Skip(d binary.Decoder) error
```

#### func (*Struct) String

```go
func (s *Struct) String() string
```

#### func (*Struct) Typename

```go
func (s *Struct) Typename() string
```

#### type Type

```go
type Type interface {
	binary.Object
	String() string
	Encode(e binary.Encoder, value interface{}) error
	Decode(d binary.Decoder) (interface{}, error)
	Skip(d binary.Decoder) error
	Typename() string
	Basename() string
}
```

Type represents the common iterface to all type objects in the schema.
