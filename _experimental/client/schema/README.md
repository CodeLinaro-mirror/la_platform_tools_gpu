# schema
--
    import "android.googlesource.com/platform/tools/gpu/_experimental/client/schema"


## Usage

#### func  AllEnumEntries

```go
func AllEnumEntries(enum *service.EnumInfo) service.EnumEntryArray
```
AllEnumEntries returns the flattened list of enum entries for the given EnumInfo
and all enums it extends.

#### func  DecodeAtoms

```go
func DecodeAtoms(stream service.AtomStream, schema service.Schema, atomMap AtomMap) ([]Atom, error)
```
DecodeAtoms decodes all atoms from the AtomStream stream.

#### func  IndexOfValue

```go
func IndexOfValue(l service.EnumEntryArray, value uint32) int
```
IndexOfValue returns the index of the enum entry with the specified value, or -1
if no value exists in the array.

#### func  ReadType

```go
func ReadType(ty service.TypeInfo, d binary.Decoder) (interface{}, error)
```
ReadType reads the schema type ty from the decoder d.

#### func  WriteType

```go
func WriteType(v interface{}, e binary.Encoder) error
```
WriteType writes v to the encoder e. The value v must be one of the following
types:

    bool
    int8
    uint8
    int16
    uint16
    int32
    uint32
    float32
    int64
    uint64
    float64
    string
    EnumValue
    Struct
    Class
    Array
    Map
    memory.Pointer

#### type Array

```go
type Array struct {
	Type     *service.ArrayInfo // The array type info.
	Elements []ArrayElement     // The array elements.
}
```

Array is a schema-typed Array value.

#### func (Array) String

```go
func (a Array) String() string
```

#### type ArrayElement

```go
type ArrayElement interface{}
```

ArrayElement is a single element held by an Array.

#### type Atom

```go
type Atom struct {
	Info         service.AtomInfo
	Observations Observations
	Arguments    []interface{}
}
```

Atom is a schema-typed Atom value.

#### func  UnpackAtom

```go
func UnpackAtom(d binary.Decoder, i service.AtomInfo) (Atom, error)
```
UnpackAtom unpacks and returns an Atom of the AtomInfo type from the decoder d.

#### func (Atom) Pack

```go
func (a Atom) Pack(e binary.Encoder) error
```
Pack encodes the Atom to the encoder e. The Atom can be decoded using the Unpack
method of the atom's AtomInfo.

#### type AtomMap

```go
type AtomMap map[uint16]binary.Class
```


#### type Class

```go
type Class struct {
	Type   *service.ClassInfo // The class type info.
	Fields []Field            // The class field values.
}
```

Class is a schema-typed Class object instance.

#### type EnumValue

```go
type EnumValue struct {
	Type  *service.EnumInfo // The enum this value belongs to.
	Value uint32            // The numerical value.
}
```

EnumValue is a schema-typed Enum value. Unlike EnumEntry, EnumValue can hold
values outside of the acceptable enum values.

#### func (EnumValue) String

```go
func (e EnumValue) String() string
```

#### type Field

```go
type Field struct {
	Info  *service.FieldInfo // The field type info.
	Value interface{}        // The field value.
}
```

Field is a schema-typed Class or Struct field instance.

#### type Map

```go
type Map struct {
	Type     *service.MapInfo // The map type info.
	Elements []MapElement     // The map elements.
}
```

Map is a schema-typed Map value.

#### type MapElement

```go
type MapElement struct {
	Key, Value interface{}
}
```

MapElement is a single key-value pair element held by a Map.

#### type Observation

```go
type Observation struct {
	Range memory.Range // Memory range that was observed.
	ID    binary.ID    // The resource identifier of the observed data.
}
```


#### func (*Observation) Decode

```go
func (o *Observation) Decode(d binary.Decoder) error
```
Decode decodes an Observation structure from the decoder d.

#### func (*Observation) Encode

```go
func (o *Observation) Encode(e binary.Encoder) error
```
Encode encodes an Observation structure to the encoder e.

#### type Observations

```go
type Observations struct {
	Reads  []Observation
	Writes []Observation
}
```


#### func (*Observations) Decode

```go
func (o *Observations) Decode(d binary.Decoder) error
```
Decode decodes an Observations structure from the decoder d.

#### func (*Observations) Encode

```go
func (o *Observations) Encode(e binary.Encoder) error
```
Encode encodes an Observations structure to the encoder e.

#### type StaticArray

```go
type StaticArray struct {
	Type     *service.StaticArrayInfo // The static array type info.
	Elements []ArrayElement           // The static array elements.
}
```

StaticArray is a schema-typed StaticArray value.

#### func (StaticArray) String

```go
func (a StaticArray) String() string
```

#### type Struct

```go
type Struct struct {
	Type   *service.StructInfo // The struct type info.
	Fields []Field             // The struct field values.
}
```

Struct is a schema-typed Struct object instance.
