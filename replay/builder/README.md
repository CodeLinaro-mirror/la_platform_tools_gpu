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
func New(architecture device.Architecture) *Builder
```
New returns a newly constructed Builder configured to replay on a target with
the specified Architecture.

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

#### func (*Builder) Architecture

```go
func (b *Builder) Architecture() device.Architecture
```
Architecture returns the architecture for the target replay device.

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

#### func (*Builder) Call

```go
func (b *Builder) Call(f FunctionInfo)
```
Call will invoke the function f, popping all parameter values previously pushed
to the stack with Push, starting with the first parameter. If f has a non-void
return type, after invoking the function the return value of the function will
be pushed on to the stack.

#### func (*Builder) Clone

```go
func (b *Builder) Clone(index int)
```
Clone makes a copy of the n-th element from the top of the stack and pushes the
copy to the top of the stack.

#### func (*Builder) CommitAtom

```go
func (b *Builder) CommitAtom()
```
CommitAtom should be called after emitting the commands to replay a single atom.
CommitAtom frees all temporary allocated memory and clears the stack.

#### func (*Builder) Copy

```go
func (b *Builder) Copy(size uint64)
```
Copy pops the target address and then the source address from the top of the
stack, and then copies Count bytes from source to target.

#### func (*Builder) Load

```go
func (b *Builder) Load(ty protocol.Type, addr value.Pointer)
```
Load loads the value of type ty from addr and then pushes the loaded value to
the top of the stack.

#### func (*Builder) MapMemory

```go
func (b *Builder) MapMemory(rng memory.Range)
```
MapMemory adds rng as a memory range that needs allocating for replay.

#### func (*Builder) Pop

```go
func (b *Builder) Pop(count uint32)
```
Pop removes the top count values from the top of the stack.

#### func (*Builder) Post

```go
func (b *Builder) Post(addr value.Pointer, size uint64, p Postback)
```
Post posts size bytes from addr to the decoder d. The decoder d must consume all
size bytes before returning; failure to do this will corrupt all subsequent
postbacks.

#### func (*Builder) Push

```go
func (b *Builder) Push(val value.Value)
```
Push pushes val to the top of the stack.

#### func (*Builder) RevertAtom

```go
func (b *Builder) RevertAtom(err error)
```
RevertAtom reverts all the instructions since the last call to BeginAtom. Any
postbacks issued since the last call to BeginAtom will be called with the error
err and a nil decoder.

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

#### func (*Builder) Write

```go
func (b *Builder) Write(rng memory.Range, resourceID binary.ID)
```
Write fills the memory range in capture address-space rng with the data of
resourceID.

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

#### type Postback

```go
type Postback func(d binary.Decoder, err error) error
```

Postback decodes a single atom's postback, returning and carrying over errors.
The Postback must decode all the data that was issued in the Post call before
returning. If err is nil, then d is the Decoder to the postback data. If d is
nil, then a previous postback failed to decode before decoding could begin for
this postback and err holds the error.

#### type ResponseDecoder

```go
type ResponseDecoder func(r io.Reader, err error)
```

ResponseDecoder decodes all postback responses from the replay virtual machine.
If err is nil, then r is the Reader to the sequential postback data. If r is
nil, then the postback data was absent or corrupted and err holds the error.
