# builder
--
    import "android.googlesource.com/platform/tools/gpu/replay/builder"

Package builder contains the Builder type to build replay payloads.

## Usage

#### type Builder

```go
type Builder struct {

	// Remappings is a map of a arbitrary keys to pointers. Typically, this is
	// used as a map of observed values to values that are only known at replay
	// execution time, such as driver generated handles.
	// The Remappings field is not accessed by the Builder and can be used in any
	// way the developer requires.
	Remappings map[interface{}]value.Pointer
}
```

Builder is used to build the Payload to send to the replay virtual machine. The
builder has a number of methods for mutating the virtual machine stack, invoking
functions and posting back data.

#### func  New

```go
func New(ptrSize, ptrAlignment int, byteOrder endian.ByteOrder) *Builder
```
New returns a newly constructed Builder configured to replay on a target
architecture that has a pointer size of ptrSize bytes and and alignment of
ptrAlignment bytes, with byte ordering that matches byteOrder.

#### func (*Builder) AllocateMemory

```go
func (b *Builder) AllocateMemory(size uint64) value.Pointer
```
AllocateMemory allocates and returns a pointer to a block of memory in the
volatile address-space big enough to hold size bytes. The memory will be
allocated for the entire replay duration and cannot be freed.

#### func (*Builder) AllocateTemporaryMemory

```go
func (b *Builder) AllocateTemporaryMemory(size uint64) value.Pointer
```
AllocateTemporaryMemory allocates and returns a pointer to a block of memory in
the temporary volatile address-space big enough to hold size bytes. The memory
block will be freed on the next call to EndAtom, upon which reading or writing
to this memory will result in undefined behavior.

#### func (*Builder) AllocateTemporaryMemoryChunks

```go
func (b *Builder) AllocateTemporaryMemoryChunks(sizes []uint64) (ptrs []value.Pointer, size uint64)
```
AllocateTemporaryMemoryChunks allocates a contiguous block of memory in the
temporary volatile address-space big enough to hold all the specified chunks
sizes, in sequential order. AllocateTemporaryMemoryChunks returns a pointer to
each of the allocated chunks and the size of the entire allocation. The
allocation block will be freed on the next call to EndAtom, upon which reading
or writing to this memory will result in undefined behavior.

#### func (*Builder) BeginAtom

```go
func (b *Builder) BeginAtom(id atom.ID)
```
BeginAtom should be called before building any replay instructions.

#### func (*Builder) Buffer

```go
func (b *Builder) Buffer(count int) value.Pointer
```
Buffer returns a pointer to a block of memory in holding the count number of
previously pushed values. If all the values are constant, then the buffer will
be held in the constant address-space, otherwise the buffer will be built in the
temporary address-space.

#### func (*Builder) Build

```go
func (b *Builder) Build(logger log.Logger) (protocol.Payload, ResponseDecoder, error)
```
Build compiles the replay instructions, returning a Payload that can be sent to
the replay virtual-machine and a ResponseDecoder for interpreting the responses.

#### func (*Builder) CallNoPush

```go
func (b *Builder) CallNoPush(f FunctionInfo)
```
CallNoPush will invoke the function f, popping all parameter values previously
pushed to the stack with Push, starting with the first parameter. Unlike
CallPush the return value of the function will not be pushed on to the stack and
functions with void return types are accepted.

#### func (*Builder) CallPush

```go
func (b *Builder) CallPush(f FunctionInfo)
```
CallPush will invoke the function f, popping all parameter values previously
pushed to the stack with Push, starting with the first parameter. After invoking
the function the return value of the function will be pushed on to the stack. If
f has a void return type then CallPush will panic.

#### func (*Builder) Clone

```go
func (b *Builder) Clone(index int)
```
Clone makes a copy of the n-th element from the top of the stack and pushes the
copy to the top of the stack.

#### func (*Builder) Copy

```go
func (b *Builder) Copy(size uint64)
```
Copy pops the target address and then the source address from the top of the
stack, and then copies Count bytes from source to target.

#### func (*Builder) EndAtom

```go
func (b *Builder) EndAtom()
```
EndAtom should be called after emitting the commands to replay a single atom.
EndAtom frees all allocated temporary memory, and asserts the stack should be
empty.

#### func (*Builder) Load

```go
func (b *Builder) Load(ty protocol.Type, addr value.Pointer)
```
Load loads the value of type ty from addr and then pushes the loaded value to
the top of the stack.

#### func (*Builder) Observation

```go
func (b *Builder) Observation(rng memory.Range, resourceID binary.ID)
```
Observation fills the memory range in capture address-space rng with the data of
resourceID.

#### func (*Builder) PointerAlignment

```go
func (b *Builder) PointerAlignment() int
```
PointerAlignment returns the required alignment of a pointer in bytes for the
replay target architecture.

#### func (*Builder) PointerSize

```go
func (b *Builder) PointerSize() int
```
PointerSize returns the size of a pointer in bytes for the target replay
architecture.

#### func (*Builder) Pop

```go
func (b *Builder) Pop(count uint32)
```
Pop removes the top count values from the top of the stack.

#### func (*Builder) Post

```go
func (b *Builder) Post(addr value.Pointer, size uint64, id atom.ID, d PostDecoder)
```
Post posts size bytes from addr to the decoder d. The decoder d must consume all
size bytes before returning; failure to do this will corrupt all subsequent
postbacks.

#### func (*Builder) Push

```go
func (b *Builder) Push(val value.Value)
```
Push pushes val to the top of the stack.

#### func (*Builder) Store

```go
func (b *Builder) Store(addr value.Pointer)
```
Store pops the value from the top of the stack and writes the value to addr.

#### func (*Builder) Strcpy

```go
func (b *Builder) Strcpy(maxCount uint64)
```
Strcpy pops the source address then the target address from the top of the
stack, and then copies at most maxCount-1 bytes from source to target. If
maxCount is greater than the source string length, then the target will be
padded with 0s. The destination buffer will always be 0-terminated.

#### func (*Builder) String

```go
func (b *Builder) String(s string) value.Pointer
```
String returns a pointer to a block of memory in the constant address-space
holding the string s. The string will be stored with a null-terminating byte.

#### type FunctionInfo

```go
type FunctionInfo struct {
	ID         uint16        // The unique identifier for the function.
	ReturnType protocol.Type // The returns type of the function.
	Parameters int           // The number of parameters for the function.
}
```

FunctionInfo holds the information about a function that can be called by the
replay virtual-machine.

#### type PostDecoder

```go
type PostDecoder func(binary.Decoder) (interface{}, error)
```

PostDecoder decodes a single atom's postback, returning the postback data or an
error. The PostDecoder must decode all the data that was issued in the Post call
before returning.

#### type Postback

```go
type Postback struct {
	ID    atom.ID     // The associated atom for this Postback.
	Data  interface{} // The postback data. Nil if Error is non-nil.
	Error error       // Error raised decoding the postback, or nil if there was no error.
}
```

Postback holds the information for a single atom's postback data.

#### type ResponseDecoder

```go
type ResponseDecoder func(r io.Reader) <-chan Postback
```

ResponseDecoder decodes all postback responses from the replay virtual machine,
writing each to the returned chan. The chan will be closed once all postbacks
have been read, or after the first Postback with a non-nil Error.
