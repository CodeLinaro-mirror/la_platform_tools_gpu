# opcode
--
    import "android.googlesource.com/platform/tools/gpu/replay/opcode"

Package opcode holds all the opcodes that are to be interpreted by the replay
virtual machine.

Documentation on each of the opcodes can be found at:

    $GOPATH/src/android.googlesource.com/platform/tools/gpu/cc/replayd/src/Interpreter_doc.txt

## Usage

#### func  CheckDisassembly

```go
func CheckDisassembly(t *testing.T, got []interface{}, expected ...interface{})
```
CheckDisassembly is a test helper function that checks the list of got opcodes
matches those in expected. If any differences are found then these are logged to
t, and the test fails.

#### func  Decode

```go
func Decode(d binary.Decoder) (interface{}, error)
```
Decode returns the opcode decoded from decoder d.

#### func  Disassemble

```go
func Disassemble(r io.Reader, byteOrder endian.ByteOrder) ([]interface{}, error)
```
Disassemble disassembles and returns the stream of encoded Opcodes from r,
stopping once an EOF is reached.

#### type Call

```go
type Call struct {
	PushReturn bool   // Should the return value be pushed onto the stack?
	FunctionID uint16 // The function identifier to call.
}
```

Call represents the CALL virtual machine opcode.

#### func (Call) Encode

```go
func (c Call) Encode(e binary.Encoder) error
```

#### type Clone

```go
type Clone struct {
	Index uint32 // Index of element from top of stack to clone.
}
```

Clone represents the CLONE virtual machine opcode.

#### func (Clone) Encode

```go
func (c Clone) Encode(e binary.Encoder) error
```

#### type Copy

```go
type Copy struct {
	Count uint32 // Number of bytes to copy.
}
```

Copy represents the COPY virtual machine opcode.

#### func (Copy) Encode

```go
func (c Copy) Encode(e binary.Encoder) error
```

#### type Extend

```go
type Extend struct {
	Value uint32 // 26 bit value to extend the top of the stack by.
}
```

Extend represents the EXTEND virtual machine opcode.

#### func (Extend) Encode

```go
func (c Extend) Encode(e binary.Encoder) error
```

#### type Label

```go
type Label struct {
	Value uint32 // 26 bit label name.
}
```

Extend represents the LABEL virtual machine opcode.

#### func (Label) Encode

```go
func (c Label) Encode(e binary.Encoder) error
```

#### type Load

```go
type Load struct {
	DataType protocol.Type // The value types to load.
}
```

Load represents the LOAD virtual machine opcode.

#### func (Load) Encode

```go
func (c Load) Encode(e binary.Encoder) error
```

#### type LoadC

```go
type LoadC struct {
	DataType protocol.Type // The value type to load.
	Address  uint32        // The pointer to the value in constant address-space.
}
```

LoadC represents the LOAD_C virtual machine opcode.

#### func (LoadC) Encode

```go
func (c LoadC) Encode(e binary.Encoder) error
```

#### type LoadV

```go
type LoadV struct {
	DataType protocol.Type // The value type to load.
	Address  uint32        // The pointer to the value in volatile address-space.
}
```

LoadV represents the LOAD_V virtual machine opcode.

#### func (LoadV) Encode

```go
func (c LoadV) Encode(e binary.Encoder) error
```

#### type Pop

```go
type Pop struct {
	Count uint32 // Number of elements to pop from the top of the stack.
}
```

Pop represents the POP virtual machine opcode.

#### func (Pop) Encode

```go
func (c Pop) Encode(e binary.Encoder) error
```

#### type Post

```go
type Post struct{}
```

Post represents the POST virtual machine opcode.

#### func (Post) Encode

```go
func (c Post) Encode(e binary.Encoder) error
```

#### type PushI

```go
type PushI struct {
	DataType protocol.Type // The value type to push.
	Value    uint32        // The value to push packed into the low 20 bits.
}
```

PushI represents the PUSH_I virtual machine opcode.

#### func (PushI) Encode

```go
func (c PushI) Encode(e binary.Encoder) error
```

#### type Resource

```go
type Resource struct {
	ID uint32 // The index of the resource identifier.
}
```

Resource represents the RESOURCE virtual machine opcode.

#### func (Resource) Encode

```go
func (c Resource) Encode(e binary.Encoder) error
```

#### type Store

```go
type Store struct{}
```

Store represents the STORE virtual machine opcode.

#### func (Store) Encode

```go
func (c Store) Encode(e binary.Encoder) error
```

#### type StoreV

```go
type StoreV struct {
	Address uint32 // Pointer in volatile address-space.
}
```

StoreV represents the STORE_V virtual machine opcode.

#### func (StoreV) Encode

```go
func (c StoreV) Encode(e binary.Encoder) error
```

#### type Strcpy

```go
type Strcpy struct {
	MaxSize uint32 // Maximum size in bytes to copy.
}
```

Strcpy represents the STRCPY virtual machine opcode.

#### func (Strcpy) Encode

```go
func (c Strcpy) Encode(e binary.Encoder) error
```
