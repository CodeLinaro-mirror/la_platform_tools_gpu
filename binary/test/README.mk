# test
--
    import "android.googlesource.com/platform/tools/gpu/binary/test"


## Usage

```go
var BadObject = &BadType{data: "BadObject"}
```

```go
var ObjectA = &TypeA{data: "ObjectA"}
```

```go
var ObjectB = &TypeB{data: "ObjectB"}
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
}
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
	binary.Generate
}
```


#### func (*TypeA) Class

```go
func (*TypeA) Class() binary.Class
```

#### type TypeB

```go
type TypeB struct {
	binary.Generate
}
```


#### func (*TypeB) Class

```go
func (*TypeB) Class() binary.Class
```
