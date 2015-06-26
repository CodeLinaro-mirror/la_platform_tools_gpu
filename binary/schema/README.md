# schema
--
    import "android.googlesource.com/platform/tools/gpu/binary/schema"

Package schema implements rtti for the binary system.

## Usage

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
	TypeID  binary.ID // The unique type identifier for the Object.
	Package string    // The package that declared the struct.
	Name    string    // The simple name of the Object.
	Display string    // The display name of the Object.
	Fields  []Field   // Descriptions of the fields of the Object.
}
```

Class represents an encodable object type with a type ID.

#### func  Lookup

```go
func Lookup(id binary.ID) *Class
```
Lookup looks up a Class by the given type id. If there is no match, it will
return nil.

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

#### type Constants

```go
type Constants []interface{}
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

#### type Int16Constant

```go
type Int16Constant struct {
	binary.Generate
	Name  string
	Value int16
}
```


#### func (Int16Constant) Add

```go
func (v Int16Constant) Add(c *Constants, t Type)
```

#### func (*Int16Constant) Class

```go
func (*Int16Constant) Class() binary.Class
```

#### type Int16Constants

```go
type Int16Constants struct {
	binary.Generate
	Type   Type            // The type of the constant.
	Values []Int16Constant // The constant values
}
```


#### func (*Int16Constants) Class

```go
func (*Int16Constants) Class() binary.Class
```

#### func (Int16Constants) Len

```go
func (s Int16Constants) Len() int
```

#### func (Int16Constants) Less

```go
func (s Int16Constants) Less(i, j int) bool
```

#### func (Int16Constants) Swap

```go
func (s Int16Constants) Swap(i, j int)
```

#### type Int32Constant

```go
type Int32Constant struct {
	binary.Generate
	Name  string
	Value int32
}
```


#### func (Int32Constant) Add

```go
func (v Int32Constant) Add(c *Constants, t Type)
```

#### func (*Int32Constant) Class

```go
func (*Int32Constant) Class() binary.Class
```

#### type Int32Constants

```go
type Int32Constants struct {
	binary.Generate
	Type   Type            // The type of the constant.
	Values []Int32Constant // The constant values
}
```


#### func (*Int32Constants) Class

```go
func (*Int32Constants) Class() binary.Class
```

#### func (Int32Constants) Len

```go
func (s Int32Constants) Len() int
```

#### func (Int32Constants) Less

```go
func (s Int32Constants) Less(i, j int) bool
```

#### func (Int32Constants) Swap

```go
func (s Int32Constants) Swap(i, j int)
```

#### type Int64Constant

```go
type Int64Constant struct {
	binary.Generate
	Name  string
	Value int64
}
```


#### func (Int64Constant) Add

```go
func (v Int64Constant) Add(c *Constants, t Type)
```

#### func (*Int64Constant) Class

```go
func (*Int64Constant) Class() binary.Class
```

#### type Int64Constants

```go
type Int64Constants struct {
	binary.Generate
	Type   Type            // The type of the constant.
	Values []Int64Constant // The constant values
}
```


#### func (*Int64Constants) Class

```go
func (*Int64Constants) Class() binary.Class
```

#### func (Int64Constants) Len

```go
func (s Int64Constants) Len() int
```

#### func (Int64Constants) Less

```go
func (s Int64Constants) Less(i, j int) bool
```

#### func (Int64Constants) Swap

```go
func (s Int64Constants) Swap(i, j int)
```

#### type Int8Constant

```go
type Int8Constant struct {
	binary.Generate
	Name  string
	Value int8
}
```


#### func (Int8Constant) Add

```go
func (v Int8Constant) Add(c *Constants, t Type)
```

#### func (*Int8Constant) Class

```go
func (*Int8Constant) Class() binary.Class
```

#### type Int8Constants

```go
type Int8Constants struct {
	binary.Generate
	Type   Type           // The type of the constant.
	Values []Int8Constant // The constant values
}
```


#### func (*Int8Constants) Class

```go
func (*Int8Constants) Class() binary.Class
```

#### func (Int8Constants) Len

```go
func (s Int8Constants) Len() int
```

#### func (Int8Constants) Less

```go
func (s Int8Constants) Less(i, j int) bool
```

#### func (Int8Constants) Swap

```go
func (s Int8Constants) Swap(i, j int)
```

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

#### type Uint16Constant

```go
type Uint16Constant struct {
	binary.Generate
	Name  string
	Value uint16
}
```


#### func (Uint16Constant) Add

```go
func (v Uint16Constant) Add(c *Constants, t Type)
```

#### func (*Uint16Constant) Class

```go
func (*Uint16Constant) Class() binary.Class
```

#### type Uint16Constants

```go
type Uint16Constants struct {
	binary.Generate
	Type   Type             // The type of the constant.
	Values []Uint16Constant // The constant values
}
```


#### func (*Uint16Constants) Class

```go
func (*Uint16Constants) Class() binary.Class
```

#### func (Uint16Constants) Len

```go
func (s Uint16Constants) Len() int
```

#### func (Uint16Constants) Less

```go
func (s Uint16Constants) Less(i, j int) bool
```

#### func (Uint16Constants) Swap

```go
func (s Uint16Constants) Swap(i, j int)
```

#### type Uint32Constant

```go
type Uint32Constant struct {
	binary.Generate
	Name  string
	Value uint32
}
```


#### func (Uint32Constant) Add

```go
func (v Uint32Constant) Add(c *Constants, t Type)
```

#### func (*Uint32Constant) Class

```go
func (*Uint32Constant) Class() binary.Class
```

#### type Uint32Constants

```go
type Uint32Constants struct {
	binary.Generate
	Type   Type             // The type of the constant.
	Values []Uint32Constant // The constant values
}
```


#### func (*Uint32Constants) Class

```go
func (*Uint32Constants) Class() binary.Class
```

#### func (Uint32Constants) Len

```go
func (s Uint32Constants) Len() int
```

#### func (Uint32Constants) Less

```go
func (s Uint32Constants) Less(i, j int) bool
```

#### func (Uint32Constants) Swap

```go
func (s Uint32Constants) Swap(i, j int)
```

#### type Uint64Constant

```go
type Uint64Constant struct {
	binary.Generate
	Name  string
	Value uint64
}
```


#### func (Uint64Constant) Add

```go
func (v Uint64Constant) Add(c *Constants, t Type)
```

#### func (*Uint64Constant) Class

```go
func (*Uint64Constant) Class() binary.Class
```

#### type Uint64Constants

```go
type Uint64Constants struct {
	binary.Generate
	Type   Type             // The type of the constant.
	Values []Uint64Constant // The constant values
}
```


#### func (*Uint64Constants) Class

```go
func (*Uint64Constants) Class() binary.Class
```

#### func (Uint64Constants) Len

```go
func (s Uint64Constants) Len() int
```

#### func (Uint64Constants) Less

```go
func (s Uint64Constants) Less(i, j int) bool
```

#### func (Uint64Constants) Swap

```go
func (s Uint64Constants) Swap(i, j int)
```

#### type Uint8Constant

```go
type Uint8Constant struct {
	binary.Generate
	Name  string
	Value uint8
}
```


#### func (Uint8Constant) Add

```go
func (v Uint8Constant) Add(c *Constants, t Type)
```

#### func (*Uint8Constant) Class

```go
func (*Uint8Constant) Class() binary.Class
```

#### type Uint8Constants

```go
type Uint8Constants struct {
	binary.Generate
	Type   Type            // The type of the constant.
	Values []Uint8Constant // The constant values
}
```


#### func (*Uint8Constants) Class

```go
func (*Uint8Constants) Class() binary.Class
```

#### func (Uint8Constants) Len

```go
func (s Uint8Constants) Len() int
```

#### func (Uint8Constants) Less

```go
func (s Uint8Constants) Less(i, j int) bool
```

#### func (Uint8Constants) Swap

```go
func (s Uint8Constants) Swap(i, j int)
```
