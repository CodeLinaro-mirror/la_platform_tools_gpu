# asm
--
    import "android.googlesource.com/platform/tools/gpu/replay/asm"

Package asm contains high-level instructions to control the replay virtual
machine.

## Usage

#### type Call

```go
type Call struct {
	PushReturn bool   // If true, the return value is pushed to the VM stack.
	FunctionID uint16 // The function id registered with the VM to invoke.
}
```

Call is an Instruction to call a VM registered function. This instruction will
pop the parameters from the VM stack starting with the first parameter. If
PushReturn is true, then the return value of the function call will be pushed to
the top of the VM stack.

#### func (Call) Encode

```go
func (a Call) Encode(r value.PointerResolver, e binary.Encoder) error
```

#### type Clone

```go
type Clone struct {
	Index int
}
```

Clone is an Instruction that makes a copy of the the n-th element from the top
of the VM stack and pushes the copy to the top of the VM stack.

#### func (Clone) Encode

```go
func (a Clone) Encode(r value.PointerResolver, e binary.Encoder) error
```

#### type Copy

```go
type Copy struct {
	Count uint64 // Number of bytes to copy.
}
```

Copy is an Instruction that pops the target address and then the source address
from the top of the VM stack, and then copies Count bytes from source to target.

#### func (Copy) Encode

```go
func (a Copy) Encode(r value.PointerResolver, e binary.Encoder) error
```

#### type Instruction

```go
type Instruction interface {
	Encode(r value.PointerResolver, e binary.Encoder) error
}
```

Instruction is the interface of all instruction types.

Encode writes the instruction's opcodes to the binary encoder e, translating any
pointers to their final, resolved addresses using the PointerResolver r. An
instruction can produce zero, one or many opcodes.

#### type Label

```go
type Label struct {
	Value uint32
}
```

Label is an Instruction that holds a marker value, used for debugging.

#### func (Label) Encode

```go
func (a Label) Encode(r value.PointerResolver, e binary.Encoder) error
```

#### type Load

```go
type Load struct {
	DataType protocol.Type
	Source   value.Pointer
}
```

Load is an Instruction that loads the value of type DataType from pointer Source
and pushes the loaded value to the top of the VM stack.

#### func (Load) Encode

```go
func (a Load) Encode(r value.PointerResolver, e binary.Encoder) error
```

#### type Nop

```go
type Nop struct{}
```

Nop is a no-operation Instruction. Instructions of this type do nothing.

#### func (Nop) Encode

```go
func (Nop) Encode(r value.PointerResolver, e binary.Encoder) error
```

#### type Pop

```go
type Pop struct {
	Count uint32 // Number of values to discard from the top of the VM stack.
}
```

Pop is an Instruction that discards Count values from the top of the VM stack.

#### func (Pop) Encode

```go
func (a Pop) Encode(r value.PointerResolver, e binary.Encoder) error
```

#### type Post

```go
type Post struct {
	Source value.Pointer
	Size   uint64
}
```

Post is an Instruction that posts Size bytes from Source to the server.

#### func (Post) Encode

```go
func (a Post) Encode(r value.PointerResolver, e binary.Encoder) error
```

#### type Push

```go
type Push struct {
	Value value.Value // The value to push on to the VM stack.
}
```

Push is an Instruction to push Value to the top of the VM stack.

#### func (Push) Encode

```go
func (a Push) Encode(r value.PointerResolver, e binary.Encoder) error
```

#### type Resource

```go
type Resource struct {
	Index       uint32
	Destination memory.Pointer
}
```

Resource is an Instruction that loads the resource with index Index of Size
bytes and writes the resource to Destination.

#### func (Resource) Encode

```go
func (a Resource) Encode(r value.PointerResolver, e binary.Encoder) error
```

#### type Store

```go
type Store struct {
	Destination value.Pointer
}
```

Store is an Instruction that pops the value from the top of the VM stack and
writes the value to Destination.

#### func (Store) Encode

```go
func (a Store) Encode(r value.PointerResolver, e binary.Encoder) error
```

#### type Strcpy

```go
type Strcpy struct {
	MaxCount uint64
}
```

Strcpy is an Instruction that pops the target address then the source address
from the top of the VM stack, and then copies at most MaxCount-1 bytes from
source to target. If the MaxCount is greater than the source string length, then
the target will be padded with 0s. The destination buffer will always be
0-terminated.

#### func (Strcpy) Encode

```go
func (a Strcpy) Encode(r value.PointerResolver, e binary.Encoder) error
```
