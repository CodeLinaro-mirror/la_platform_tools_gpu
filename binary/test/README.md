# test
--
    import "android.googlesource.com/platform/tools/gpu/binary/test"


## Usage

```go
var (
	TypeAID = binary.ID{0xa4, 0xbe, 0x00, 0x04, 0x4c, 0x84, 0x76, 0x86, 0xdc, 0x77, 0x63, 0x6d, 0x19, 0xdd, 0x63, 0x33, 0x17, 0x38, 0xbf, 0x24}
	TypeBID = binary.ID{0x73, 0xbd, 0xff, 0x55, 0x9c, 0xc4, 0x5b, 0xe3, 0xaf, 0x72, 0xfd, 0xb6, 0x97, 0xfb, 0x0e, 0xe1, 0x8d, 0x19, 0xa9, 0x67}
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
	Values []interface{}
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
