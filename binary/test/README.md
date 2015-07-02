# test
--
    import "android.googlesource.com/platform/tools/gpu/binary/test"


## Usage

```go
var (
	TypeAID = binary.ID{0x29, 0x0a, 0x4b, 0x25, 0x7d, 0x55, 0xab, 0x2b, 0x8f, 0x03, 0x32, 0x53, 0x7f, 0xd8, 0x66, 0x69, 0xbc, 0x77, 0x07, 0x98}
	TypeBID = binary.ID{0x04, 0x84, 0xdf, 0x7d, 0x88, 0x25, 0xef, 0x43, 0xf7, 0x71, 0x3c, 0x5c, 0x03, 0x2b, 0xe0, 0xfd, 0x42, 0x1f, 0x6e, 0x87}
)
```

```go
var BadObject = &BadType{Data: "BadObject"}
```

```go
var Namespace = registry.NewNamespace()
```

```go
var ObjectA = &TypeA{Data: "ObjectA"}
```

```go
var ObjectB = &TypeB{Data: "ObjectB"}
```

#### func  DecodeObject

```go
func DecodeObject(t *testing.T, entry Entry, d binary.Decoder, reader *bytes.Reader)
```

#### func  DecodeValue

```go
func DecodeValue(t *testing.T, entry Entry, d binary.Decoder, reader *bytes.Reader)
```

#### func  EncodeObject

```go
func EncodeObject(t *testing.T, entry Entry, e binary.Encoder, buf *bytes.Buffer)
```

#### func  EncodeValue

```go
func EncodeValue(t *testing.T, entry Entry, e binary.Encoder, buf *bytes.Buffer)
```

#### func  VerifyData

```go
func VerifyData(t *testing.T, entry Entry, got *bytes.Buffer)
```

#### type BadType

```go
type BadType struct {
	binary.Generate `disable:"true"`
	Data            string
}
```


#### type Bytes

```go
type Bytes struct {
	Data []byte
}
```


#### func (Bytes) Add

```go
func (b Bytes) Add(data ...byte) Bytes
```

#### func (Bytes) ID

```go
func (b Bytes) ID(id binary.ID) Bytes
```

#### type Entry

```go
type Entry struct {
	Name   string
	Values []binary.Object
	Data   []byte
}
```


#### type TypeA

```go
type TypeA struct {
	binary.Generate `id:"TypeAID"`
	Data            string
}
```


#### func (*TypeA) Class

```go
func (*TypeA) Class() binary.Class
```

#### type TypeB

```go
type TypeB struct {
	binary.Generate `id:"TypeBID"`
	Data            string
}
```


#### func (*TypeB) Class

```go
func (*TypeB) Class() binary.Class
```
