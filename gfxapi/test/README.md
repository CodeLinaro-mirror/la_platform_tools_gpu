# test
--
    import "android.googlesource.com/platform/tools/gpu/gfxapi/test"

Package test is the integration test suite for the api compiler and templates.

## Usage

```go
var Namespace = registry.NewNamespace()
```

#### func  API

```go
func API() gfxapi.API
```

#### type Boolˢ

```go
type Boolˢ struct {
	binary.Generate
	SliceInfo
}
```

Boolˢ is a slice of bool.

#### func  AsBoolˢ

```go
func AsBoolˢ(s Slice, ϟs *gfxapi.State) Boolˢ
```
AsBoolˢ returns s cast to a Boolˢ. The returned slice length will be calculated
so that the returned slice is no longer (in bytes) than s.

#### func  MakeBoolˢ

```go
func MakeBoolˢ(count uint64, ϟs *gfxapi.State) Boolˢ
```
MakeBoolˢ returns a Boolˢ backed by a new memory pool.

#### func (*Boolˢ) Class

```go
func (*Boolˢ) Class() binary.Class
```

#### func (Boolˢ) Clone

```go
func (s Boolˢ) Clone(ϟs *gfxapi.State) Boolˢ
```
Clone returns a copy of the Boolˢ in a new memory pool.

#### func (Boolˢ) Copy

```go
func (dst Boolˢ) Copy(src Boolˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Boolˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (Boolˢ) Decoder

```go
func (s Boolˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (Boolˢ) ElementSize

```go
func (s Boolˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Boolˢ points to.

#### func (Boolˢ) Encoder

```go
func (s Boolˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (Boolˢ) Index

```go
func (s Boolˢ) Index(i uint64, ϟs *gfxapi.State) Boolᵖ
```
Index returns a Boolᵖ to the i'th element in this Boolˢ.

#### func (Boolˢ) OnRead

```go
func (s Boolˢ) OnRead(ϟs *gfxapi.State) Boolˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (Boolˢ) OnWrite

```go
func (s Boolˢ) OnWrite(ϟs *gfxapi.State) Boolˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (Boolˢ) Range

```go
func (s Boolˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (Boolˢ) Read

```go
func (s Boolˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []bool
```
Read reads and returns all the bool elements in this Boolˢ.

#### func (Boolˢ) ResourceID

```go
func (s Boolˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (Boolˢ) Slice

```go
func (s Boolˢ) Slice(start, end uint64, ϟs *gfxapi.State) Boolˢ
```
Slice returns a sub-slice from the Boolˢ using start and end indices.

#### func (Boolˢ) String

```go
func (s Boolˢ) String() string
```
String returns a string description of the Boolˢ slice.

#### func (Boolˢ) Write

```go
func (s Boolˢ) Write(src []bool, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type Boolᵖ

```go
type Boolᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}
```

Boolᵖ is a pointer to a bool element.

#### func  NewBoolᵖ

```go
func NewBoolᵖ(addr memory.Pointer) Boolᵖ
```
NewBoolᵖ returns a Boolᵖ that points to addr in the application pool.

#### func (*Boolᵖ) Class

```go
func (*Boolᵖ) Class() binary.Class
```

#### func (Boolᵖ) ElementSize

```go
func (p Boolᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Boolᵖ points to.

#### func (Boolᵖ) OnRead

```go
func (p Boolᵖ) OnRead(ϟs *gfxapi.State) Boolᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (Boolᵖ) OnWrite

```go
func (p Boolᵖ) OnWrite(ϟs *gfxapi.State) Boolᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (Boolᵖ) Read

```go
func (p Boolᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) bool
```
Read reads and returns the bool element at the pointer.

#### func (Boolᵖ) Slice

```go
func (p Boolᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Boolˢ
```
Slice returns a new Boolˢ from the pointer using start and end indices.

#### func (Boolᵖ) String

```go
func (p Boolᵖ) String() string
```
String returns a string description of the Boolᵖ pointer.

#### func (Boolᵖ) Write

```go
func (p Boolᵖ) Write(value bool, ϟs *gfxapi.State)
```
Write writes value to the bool element at the pointer.

#### type Charˢ

```go
type Charˢ struct {
	binary.Generate
	SliceInfo
}
```

Charˢ is a slice of byte.

#### func  AsCharˢ

```go
func AsCharˢ(s Slice, ϟs *gfxapi.State) Charˢ
```
AsCharˢ returns s cast to a Charˢ. The returned slice length will be calculated
so that the returned slice is no longer (in bytes) than s.

#### func  MakeCharˢ

```go
func MakeCharˢ(count uint64, ϟs *gfxapi.State) Charˢ
```
MakeCharˢ returns a Charˢ backed by a new memory pool.

#### func  MakeCharˢFromString

```go
func MakeCharˢFromString(str string, ϟs *gfxapi.State) Charˢ
```
MakeCharˢFromString returns a Charˢ backed by a new memory pool containing a
copy of str.

#### func (*Charˢ) Class

```go
func (*Charˢ) Class() binary.Class
```

#### func (Charˢ) Clone

```go
func (s Charˢ) Clone(ϟs *gfxapi.State) Charˢ
```
Clone returns a copy of the Charˢ in a new memory pool.

#### func (Charˢ) Copy

```go
func (dst Charˢ) Copy(src Charˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Charˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (Charˢ) Decoder

```go
func (s Charˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (Charˢ) ElementSize

```go
func (s Charˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Charˢ points to.

#### func (Charˢ) Encoder

```go
func (s Charˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (Charˢ) Index

```go
func (s Charˢ) Index(i uint64, ϟs *gfxapi.State) Charᵖ
```
Index returns a Charᵖ to the i'th element in this Charˢ.

#### func (Charˢ) OnRead

```go
func (s Charˢ) OnRead(ϟs *gfxapi.State) Charˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (Charˢ) OnWrite

```go
func (s Charˢ) OnWrite(ϟs *gfxapi.State) Charˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (Charˢ) Range

```go
func (s Charˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (Charˢ) Read

```go
func (s Charˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []byte
```
Read reads and returns all the byte elements in this Charˢ.

#### func (Charˢ) ResourceID

```go
func (s Charˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (Charˢ) Slice

```go
func (s Charˢ) Slice(start, end uint64, ϟs *gfxapi.State) Charˢ
```
Slice returns a sub-slice from the Charˢ using start and end indices.

#### func (Charˢ) String

```go
func (s Charˢ) String() string
```
String returns a string description of the Charˢ slice.

#### func (Charˢ) Write

```go
func (s Charˢ) Write(src []byte, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type Charᵖ

```go
type Charᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}
```

Charᵖ is a pointer to a byte element.

#### func  NewCharᵖ

```go
func NewCharᵖ(addr memory.Pointer) Charᵖ
```
NewCharᵖ returns a Charᵖ that points to addr in the application pool.

#### func (*Charᵖ) Class

```go
func (*Charᵖ) Class() binary.Class
```

#### func (Charᵖ) ElementSize

```go
func (p Charᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Charᵖ points to.

#### func (Charᵖ) OnRead

```go
func (p Charᵖ) OnRead(ϟs *gfxapi.State) Charᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (Charᵖ) OnWrite

```go
func (p Charᵖ) OnWrite(ϟs *gfxapi.State) Charᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (Charᵖ) Read

```go
func (p Charᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) byte
```
Read reads and returns the byte element at the pointer.

#### func (Charᵖ) Slice

```go
func (p Charᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Charˢ
```
Slice returns a new Charˢ from the pointer using start and end indices.

#### func (Charᵖ) String

```go
func (p Charᵖ) String() string
```
String returns a string description of the Charᵖ pointer.

#### func (Charᵖ) StringSlice

```go
func (p Charᵖ) StringSlice(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, incNullTerm bool) Charˢ
```
StringSlice returns a slice starting at p and ending at the first 0 byte
null-terminator. If incNullTerm is true then the null-terminator is included in
the slice.

#### func (Charᵖ) Write

```go
func (p Charᵖ) Write(value byte, ϟs *gfxapi.State)
```
Write writes value to the byte element at the pointer.

#### type Charᵖˢ

```go
type Charᵖˢ struct {
	binary.Generate
	SliceInfo
}
```

Charᵖˢ is a slice of Charᵖ.

#### func  AsCharᵖˢ

```go
func AsCharᵖˢ(s Slice, ϟs *gfxapi.State) Charᵖˢ
```
AsCharᵖˢ returns s cast to a Charᵖˢ. The returned slice length will be
calculated so that the returned slice is no longer (in bytes) than s.

#### func  MakeCharᵖˢ

```go
func MakeCharᵖˢ(count uint64, ϟs *gfxapi.State) Charᵖˢ
```
MakeCharᵖˢ returns a Charᵖˢ backed by a new memory pool.

#### func (*Charᵖˢ) Class

```go
func (*Charᵖˢ) Class() binary.Class
```

#### func (Charᵖˢ) Clone

```go
func (s Charᵖˢ) Clone(ϟs *gfxapi.State) Charᵖˢ
```
Clone returns a copy of the Charᵖˢ in a new memory pool.

#### func (Charᵖˢ) Copy

```go
func (dst Charᵖˢ) Copy(src Charᵖˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Charᵖˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (Charᵖˢ) Decoder

```go
func (s Charᵖˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (Charᵖˢ) ElementSize

```go
func (s Charᵖˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Charᵖˢ points to.

#### func (Charᵖˢ) Encoder

```go
func (s Charᵖˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (Charᵖˢ) Index

```go
func (s Charᵖˢ) Index(i uint64, ϟs *gfxapi.State) Charᵖᵖ
```
Index returns a Charᵖᵖ to the i'th element in this Charᵖˢ.

#### func (Charᵖˢ) OnRead

```go
func (s Charᵖˢ) OnRead(ϟs *gfxapi.State) Charᵖˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (Charᵖˢ) OnWrite

```go
func (s Charᵖˢ) OnWrite(ϟs *gfxapi.State) Charᵖˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (Charᵖˢ) Range

```go
func (s Charᵖˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (Charᵖˢ) Read

```go
func (s Charᵖˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []Charᵖ
```
Read reads and returns all the Charᵖ elements in this Charᵖˢ.

#### func (Charᵖˢ) ResourceID

```go
func (s Charᵖˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (Charᵖˢ) Slice

```go
func (s Charᵖˢ) Slice(start, end uint64, ϟs *gfxapi.State) Charᵖˢ
```
Slice returns a sub-slice from the Charᵖˢ using start and end indices.

#### func (Charᵖˢ) String

```go
func (s Charᵖˢ) String() string
```
String returns a string description of the Charᵖˢ slice.

#### func (Charᵖˢ) Write

```go
func (s Charᵖˢ) Write(src []Charᵖ, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type Charᵖᵖ

```go
type Charᵖᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}
```

Charᵖᵖ is a pointer to a Charᵖ element. Note: Pointers are stored differently
between the application pool and internal pools.

    * The application pool stores pointers as an address an architecture-dependant size.
    * Internal pools store pointers as an 64-bit unsigned address and a 32-bit unsigned
      pool identifier.

#### func  NewCharᵖᵖ

```go
func NewCharᵖᵖ(addr memory.Pointer) Charᵖᵖ
```
NewCharᵖᵖ returns a Charᵖᵖ that points to addr in the application pool.

#### func (*Charᵖᵖ) Class

```go
func (*Charᵖᵖ) Class() binary.Class
```

#### func (Charᵖᵖ) ElementSize

```go
func (p Charᵖᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Charᵖᵖ points to.

#### func (Charᵖᵖ) OnRead

```go
func (p Charᵖᵖ) OnRead(ϟs *gfxapi.State) Charᵖᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (Charᵖᵖ) OnWrite

```go
func (p Charᵖᵖ) OnWrite(ϟs *gfxapi.State) Charᵖᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (Charᵖᵖ) Read

```go
func (p Charᵖᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) Charᵖ
```
Read reads and returns the Charᵖ element at the pointer.

#### func (Charᵖᵖ) Slice

```go
func (p Charᵖᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Charᵖˢ
```
Slice returns a new Charᵖˢ from the pointer using start and end indices.

#### func (Charᵖᵖ) String

```go
func (p Charᵖᵖ) String() string
```
String returns a string description of the Charᵖᵖ pointer.

#### func (Charᵖᵖ) Write

```go
func (p Charᵖᵖ) Write(value Charᵖ, ϟs *gfxapi.State)
```
Write writes value to the Charᵖ element at the pointer.

#### type CmdBool

```go
type CmdBool struct {
	binary.Generate `display:"cmd_bool"`

	Result bool
}
```

//////////////////////////////////////////////////////////////////////////////
CmdBool
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdBool

```go
func NewCmdBool(Result bool) *CmdBool
```

#### func (*CmdBool) API

```go
func (c *CmdBool) API() gfxapi.API
```

#### func (*CmdBool) AddRead

```go
func (a *CmdBool) AddRead(rng memory.Range, id binary.ID) *CmdBool
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdBool pointer is returned so that calls can be chained.

#### func (*CmdBool) AddWrite

```go
func (a *CmdBool) AddWrite(rng memory.Range, id binary.ID) *CmdBool
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdBool pointer is returned so that calls can be chained.

#### func (*CmdBool) Class

```go
func (*CmdBool) Class() binary.Class
```

#### func (*CmdBool) Flags

```go
func (c *CmdBool) Flags() atom.Flags
```

#### func (*CmdBool) Mutate

```go
func (ϟa *CmdBool) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdBool) Observations

```go
func (a *CmdBool) Observations() *atom.Observations
```

#### func (*CmdBool) Replay

```go
func (ϟa *CmdBool) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdBool) String

```go
func (a *CmdBool) String() string
```

#### func (*CmdBool) TypeID

```go
func (c *CmdBool) TypeID() atom.TypeID
```

#### type CmdCharptrToString

```go
type CmdCharptrToString struct {
	binary.Generate `display:"cmd_charptr_to_string"`

	S Charᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
CmdCharptrToString
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdCharptrToString

```go
func NewCmdCharptrToString(S memory.Pointer) *CmdCharptrToString
```

#### func (*CmdCharptrToString) API

```go
func (c *CmdCharptrToString) API() gfxapi.API
```

#### func (*CmdCharptrToString) AddRead

```go
func (a *CmdCharptrToString) AddRead(rng memory.Range, id binary.ID) *CmdCharptrToString
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdCharptrToString pointer is returned so that calls can be
chained.

#### func (*CmdCharptrToString) AddWrite

```go
func (a *CmdCharptrToString) AddWrite(rng memory.Range, id binary.ID) *CmdCharptrToString
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdCharptrToString pointer is returned so that calls can be
chained.

#### func (*CmdCharptrToString) Class

```go
func (*CmdCharptrToString) Class() binary.Class
```

#### func (*CmdCharptrToString) Flags

```go
func (c *CmdCharptrToString) Flags() atom.Flags
```

#### func (*CmdCharptrToString) Mutate

```go
func (ϟa *CmdCharptrToString) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdCharptrToString) Observations

```go
func (a *CmdCharptrToString) Observations() *atom.Observations
```

#### func (*CmdCharptrToString) Replay

```go
func (ϟa *CmdCharptrToString) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdCharptrToString) String

```go
func (a *CmdCharptrToString) String() string
```

#### func (*CmdCharptrToString) TypeID

```go
func (c *CmdCharptrToString) TypeID() atom.TypeID
```

#### type CmdCharsliceToString

```go
type CmdCharsliceToString struct {
	binary.Generate `display:"cmd_charslice_to_string"`

	S   Charᵖ
	Len uint32
}
```

//////////////////////////////////////////////////////////////////////////////
CmdCharsliceToString
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdCharsliceToString

```go
func NewCmdCharsliceToString(S memory.Pointer, Len uint32) *CmdCharsliceToString
```

#### func (*CmdCharsliceToString) API

```go
func (c *CmdCharsliceToString) API() gfxapi.API
```

#### func (*CmdCharsliceToString) AddRead

```go
func (a *CmdCharsliceToString) AddRead(rng memory.Range, id binary.ID) *CmdCharsliceToString
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdCharsliceToString pointer is returned so that calls can be
chained.

#### func (*CmdCharsliceToString) AddWrite

```go
func (a *CmdCharsliceToString) AddWrite(rng memory.Range, id binary.ID) *CmdCharsliceToString
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdCharsliceToString pointer is returned so that calls can be
chained.

#### func (*CmdCharsliceToString) Class

```go
func (*CmdCharsliceToString) Class() binary.Class
```

#### func (*CmdCharsliceToString) Flags

```go
func (c *CmdCharsliceToString) Flags() atom.Flags
```

#### func (*CmdCharsliceToString) Mutate

```go
func (ϟa *CmdCharsliceToString) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdCharsliceToString) Observations

```go
func (a *CmdCharsliceToString) Observations() *atom.Observations
```

#### func (*CmdCharsliceToString) Replay

```go
func (ϟa *CmdCharsliceToString) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdCharsliceToString) String

```go
func (a *CmdCharsliceToString) String() string
```

#### func (*CmdCharsliceToString) TypeID

```go
func (c *CmdCharsliceToString) TypeID() atom.TypeID
```

#### type CmdClone

```go
type CmdClone struct {
	binary.Generate `display:"cmd_clone"`

	Src U8ᵖ
	Cnt uint32
}
```

//////////////////////////////////////////////////////////////////////////////
CmdClone
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdClone

```go
func NewCmdClone(Src memory.Pointer, Cnt uint32) *CmdClone
```

#### func (*CmdClone) API

```go
func (c *CmdClone) API() gfxapi.API
```

#### func (*CmdClone) AddRead

```go
func (a *CmdClone) AddRead(rng memory.Range, id binary.ID) *CmdClone
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdClone pointer is returned so that calls can be chained.

#### func (*CmdClone) AddWrite

```go
func (a *CmdClone) AddWrite(rng memory.Range, id binary.ID) *CmdClone
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdClone pointer is returned so that calls can be chained.

#### func (*CmdClone) Class

```go
func (*CmdClone) Class() binary.Class
```

#### func (*CmdClone) Flags

```go
func (c *CmdClone) Flags() atom.Flags
```

#### func (*CmdClone) Mutate

```go
func (ϟa *CmdClone) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdClone) Observations

```go
func (a *CmdClone) Observations() *atom.Observations
```

#### func (*CmdClone) Replay

```go
func (ϟa *CmdClone) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdClone) String

```go
func (a *CmdClone) String() string
```

#### func (*CmdClone) TypeID

```go
func (c *CmdClone) TypeID() atom.TypeID
```

#### type CmdCopy

```go
type CmdCopy struct {
	binary.Generate `display:"cmd_copy"`

	Src U8ᵖ
	Cnt uint32
}
```

//////////////////////////////////////////////////////////////////////////////
CmdCopy
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdCopy

```go
func NewCmdCopy(Src memory.Pointer, Cnt uint32) *CmdCopy
```

#### func (*CmdCopy) API

```go
func (c *CmdCopy) API() gfxapi.API
```

#### func (*CmdCopy) AddRead

```go
func (a *CmdCopy) AddRead(rng memory.Range, id binary.ID) *CmdCopy
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdCopy pointer is returned so that calls can be chained.

#### func (*CmdCopy) AddWrite

```go
func (a *CmdCopy) AddWrite(rng memory.Range, id binary.ID) *CmdCopy
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdCopy pointer is returned so that calls can be chained.

#### func (*CmdCopy) Class

```go
func (*CmdCopy) Class() binary.Class
```

#### func (*CmdCopy) Flags

```go
func (c *CmdCopy) Flags() atom.Flags
```

#### func (*CmdCopy) Mutate

```go
func (ϟa *CmdCopy) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdCopy) Observations

```go
func (a *CmdCopy) Observations() *atom.Observations
```

#### func (*CmdCopy) Replay

```go
func (ϟa *CmdCopy) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdCopy) String

```go
func (a *CmdCopy) String() string
```

#### func (*CmdCopy) TypeID

```go
func (c *CmdCopy) TypeID() atom.TypeID
```

#### type CmdF32

```go
type CmdF32 struct {
	binary.Generate `display:"cmd_f32"`

	Result float32
}
```

//////////////////////////////////////////////////////////////////////////////
CmdF32
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdF32

```go
func NewCmdF32(Result float32) *CmdF32
```

#### func (*CmdF32) API

```go
func (c *CmdF32) API() gfxapi.API
```

#### func (*CmdF32) AddRead

```go
func (a *CmdF32) AddRead(rng memory.Range, id binary.ID) *CmdF32
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdF32 pointer is returned so that calls can be chained.

#### func (*CmdF32) AddWrite

```go
func (a *CmdF32) AddWrite(rng memory.Range, id binary.ID) *CmdF32
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdF32 pointer is returned so that calls can be chained.

#### func (*CmdF32) Class

```go
func (*CmdF32) Class() binary.Class
```

#### func (*CmdF32) Flags

```go
func (c *CmdF32) Flags() atom.Flags
```

#### func (*CmdF32) Mutate

```go
func (ϟa *CmdF32) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdF32) Observations

```go
func (a *CmdF32) Observations() *atom.Observations
```

#### func (*CmdF32) Replay

```go
func (ϟa *CmdF32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdF32) String

```go
func (a *CmdF32) String() string
```

#### func (*CmdF32) TypeID

```go
func (c *CmdF32) TypeID() atom.TypeID
```

#### type CmdF64

```go
type CmdF64 struct {
	binary.Generate `display:"cmd_f64"`

	Result float64
}
```

//////////////////////////////////////////////////////////////////////////////
CmdF64
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdF64

```go
func NewCmdF64(Result float64) *CmdF64
```

#### func (*CmdF64) API

```go
func (c *CmdF64) API() gfxapi.API
```

#### func (*CmdF64) AddRead

```go
func (a *CmdF64) AddRead(rng memory.Range, id binary.ID) *CmdF64
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdF64 pointer is returned so that calls can be chained.

#### func (*CmdF64) AddWrite

```go
func (a *CmdF64) AddWrite(rng memory.Range, id binary.ID) *CmdF64
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdF64 pointer is returned so that calls can be chained.

#### func (*CmdF64) Class

```go
func (*CmdF64) Class() binary.Class
```

#### func (*CmdF64) Flags

```go
func (c *CmdF64) Flags() atom.Flags
```

#### func (*CmdF64) Mutate

```go
func (ϟa *CmdF64) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdF64) Observations

```go
func (a *CmdF64) Observations() *atom.Observations
```

#### func (*CmdF64) Replay

```go
func (ϟa *CmdF64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdF64) String

```go
func (a *CmdF64) String() string
```

#### func (*CmdF64) TypeID

```go
func (c *CmdF64) TypeID() atom.TypeID
```

#### type CmdMake

```go
type CmdMake struct {
	binary.Generate `display:"cmd_make"`

	Cnt uint32
}
```

//////////////////////////////////////////////////////////////////////////////
CmdMake
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdMake

```go
func NewCmdMake(Cnt uint32) *CmdMake
```

#### func (*CmdMake) API

```go
func (c *CmdMake) API() gfxapi.API
```

#### func (*CmdMake) AddRead

```go
func (a *CmdMake) AddRead(rng memory.Range, id binary.ID) *CmdMake
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdMake pointer is returned so that calls can be chained.

#### func (*CmdMake) AddWrite

```go
func (a *CmdMake) AddWrite(rng memory.Range, id binary.ID) *CmdMake
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdMake pointer is returned so that calls can be chained.

#### func (*CmdMake) Class

```go
func (*CmdMake) Class() binary.Class
```

#### func (*CmdMake) Flags

```go
func (c *CmdMake) Flags() atom.Flags
```

#### func (*CmdMake) Mutate

```go
func (ϟa *CmdMake) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdMake) Observations

```go
func (a *CmdMake) Observations() *atom.Observations
```

#### func (*CmdMake) Replay

```go
func (ϟa *CmdMake) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdMake) String

```go
func (a *CmdMake) String() string
```

#### func (*CmdMake) TypeID

```go
func (c *CmdMake) TypeID() atom.TypeID
```

#### type CmdPointer

```go
type CmdPointer struct {
	binary.Generate `display:"cmd_pointer"`

	Result Voidᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
CmdPointer
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdPointer

```go
func NewCmdPointer(Result memory.Pointer) *CmdPointer
```

#### func (*CmdPointer) API

```go
func (c *CmdPointer) API() gfxapi.API
```

#### func (*CmdPointer) AddRead

```go
func (a *CmdPointer) AddRead(rng memory.Range, id binary.ID) *CmdPointer
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdPointer pointer is returned so that calls can be chained.

#### func (*CmdPointer) AddWrite

```go
func (a *CmdPointer) AddWrite(rng memory.Range, id binary.ID) *CmdPointer
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdPointer pointer is returned so that calls can be chained.

#### func (*CmdPointer) Class

```go
func (*CmdPointer) Class() binary.Class
```

#### func (*CmdPointer) Flags

```go
func (c *CmdPointer) Flags() atom.Flags
```

#### func (*CmdPointer) Mutate

```go
func (ϟa *CmdPointer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdPointer) Observations

```go
func (a *CmdPointer) Observations() *atom.Observations
```

#### func (*CmdPointer) Replay

```go
func (ϟa *CmdPointer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdPointer) String

```go
func (a *CmdPointer) String() string
```

#### func (*CmdPointer) TypeID

```go
func (c *CmdPointer) TypeID() atom.TypeID
```

#### type CmdRemapped

```go
type CmdRemapped struct {
	binary.Generate `display:"cmd_remapped"`

	Result remapped
}
```

//////////////////////////////////////////////////////////////////////////////
CmdRemapped
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdRemapped

```go
func NewCmdRemapped(Result remapped) *CmdRemapped
```

#### func (*CmdRemapped) API

```go
func (c *CmdRemapped) API() gfxapi.API
```

#### func (*CmdRemapped) AddRead

```go
func (a *CmdRemapped) AddRead(rng memory.Range, id binary.ID) *CmdRemapped
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdRemapped pointer is returned so that calls can be chained.

#### func (*CmdRemapped) AddWrite

```go
func (a *CmdRemapped) AddWrite(rng memory.Range, id binary.ID) *CmdRemapped
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdRemapped pointer is returned so that calls can be chained.

#### func (*CmdRemapped) Class

```go
func (*CmdRemapped) Class() binary.Class
```

#### func (*CmdRemapped) Flags

```go
func (c *CmdRemapped) Flags() atom.Flags
```

#### func (*CmdRemapped) Mutate

```go
func (ϟa *CmdRemapped) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdRemapped) Observations

```go
func (a *CmdRemapped) Observations() *atom.Observations
```

#### func (*CmdRemapped) Replay

```go
func (ϟa *CmdRemapped) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdRemapped) String

```go
func (a *CmdRemapped) String() string
```

#### func (*CmdRemapped) TypeID

```go
func (c *CmdRemapped) TypeID() atom.TypeID
```

#### type CmdS16

```go
type CmdS16 struct {
	binary.Generate `display:"cmd_s16"`

	Result int16
}
```

//////////////////////////////////////////////////////////////////////////////
CmdS16
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdS16

```go
func NewCmdS16(Result int16) *CmdS16
```

#### func (*CmdS16) API

```go
func (c *CmdS16) API() gfxapi.API
```

#### func (*CmdS16) AddRead

```go
func (a *CmdS16) AddRead(rng memory.Range, id binary.ID) *CmdS16
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdS16 pointer is returned so that calls can be chained.

#### func (*CmdS16) AddWrite

```go
func (a *CmdS16) AddWrite(rng memory.Range, id binary.ID) *CmdS16
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdS16 pointer is returned so that calls can be chained.

#### func (*CmdS16) Class

```go
func (*CmdS16) Class() binary.Class
```

#### func (*CmdS16) Flags

```go
func (c *CmdS16) Flags() atom.Flags
```

#### func (*CmdS16) Mutate

```go
func (ϟa *CmdS16) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdS16) Observations

```go
func (a *CmdS16) Observations() *atom.Observations
```

#### func (*CmdS16) Replay

```go
func (ϟa *CmdS16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdS16) String

```go
func (a *CmdS16) String() string
```

#### func (*CmdS16) TypeID

```go
func (c *CmdS16) TypeID() atom.TypeID
```

#### type CmdS32

```go
type CmdS32 struct {
	binary.Generate `display:"cmd_s32"`

	Result int32
}
```

//////////////////////////////////////////////////////////////////////////////
CmdS32
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdS32

```go
func NewCmdS32(Result int32) *CmdS32
```

#### func (*CmdS32) API

```go
func (c *CmdS32) API() gfxapi.API
```

#### func (*CmdS32) AddRead

```go
func (a *CmdS32) AddRead(rng memory.Range, id binary.ID) *CmdS32
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdS32 pointer is returned so that calls can be chained.

#### func (*CmdS32) AddWrite

```go
func (a *CmdS32) AddWrite(rng memory.Range, id binary.ID) *CmdS32
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdS32 pointer is returned so that calls can be chained.

#### func (*CmdS32) Class

```go
func (*CmdS32) Class() binary.Class
```

#### func (*CmdS32) Flags

```go
func (c *CmdS32) Flags() atom.Flags
```

#### func (*CmdS32) Mutate

```go
func (ϟa *CmdS32) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdS32) Observations

```go
func (a *CmdS32) Observations() *atom.Observations
```

#### func (*CmdS32) Replay

```go
func (ϟa *CmdS32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdS32) String

```go
func (a *CmdS32) String() string
```

#### func (*CmdS32) TypeID

```go
func (c *CmdS32) TypeID() atom.TypeID
```

#### type CmdS64

```go
type CmdS64 struct {
	binary.Generate `display:"cmd_s64"`

	Result int64
}
```

//////////////////////////////////////////////////////////////////////////////
CmdS64
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdS64

```go
func NewCmdS64(Result int64) *CmdS64
```

#### func (*CmdS64) API

```go
func (c *CmdS64) API() gfxapi.API
```

#### func (*CmdS64) AddRead

```go
func (a *CmdS64) AddRead(rng memory.Range, id binary.ID) *CmdS64
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdS64 pointer is returned so that calls can be chained.

#### func (*CmdS64) AddWrite

```go
func (a *CmdS64) AddWrite(rng memory.Range, id binary.ID) *CmdS64
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdS64 pointer is returned so that calls can be chained.

#### func (*CmdS64) Class

```go
func (*CmdS64) Class() binary.Class
```

#### func (*CmdS64) Flags

```go
func (c *CmdS64) Flags() atom.Flags
```

#### func (*CmdS64) Mutate

```go
func (ϟa *CmdS64) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdS64) Observations

```go
func (a *CmdS64) Observations() *atom.Observations
```

#### func (*CmdS64) Replay

```go
func (ϟa *CmdS64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdS64) String

```go
func (a *CmdS64) String() string
```

#### func (*CmdS64) TypeID

```go
func (c *CmdS64) TypeID() atom.TypeID
```

#### type CmdS8

```go
type CmdS8 struct {
	binary.Generate `display:"cmd_s8"`

	Result int8
}
```

//////////////////////////////////////////////////////////////////////////////
CmdS8
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdS8

```go
func NewCmdS8(Result int8) *CmdS8
```

#### func (*CmdS8) API

```go
func (c *CmdS8) API() gfxapi.API
```

#### func (*CmdS8) AddRead

```go
func (a *CmdS8) AddRead(rng memory.Range, id binary.ID) *CmdS8
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdS8 pointer is returned so that calls can be chained.

#### func (*CmdS8) AddWrite

```go
func (a *CmdS8) AddWrite(rng memory.Range, id binary.ID) *CmdS8
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdS8 pointer is returned so that calls can be chained.

#### func (*CmdS8) Class

```go
func (*CmdS8) Class() binary.Class
```

#### func (*CmdS8) Flags

```go
func (c *CmdS8) Flags() atom.Flags
```

#### func (*CmdS8) Mutate

```go
func (ϟa *CmdS8) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdS8) Observations

```go
func (a *CmdS8) Observations() *atom.Observations
```

#### func (*CmdS8) Replay

```go
func (ϟa *CmdS8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdS8) String

```go
func (a *CmdS8) String() string
```

#### func (*CmdS8) TypeID

```go
func (c *CmdS8) TypeID() atom.TypeID
```

#### type CmdSliceCasts

```go
type CmdSliceCasts struct {
	binary.Generate `display:"cmd_slice_casts"`

	S U16ᵖ
	L uint32
}
```

//////////////////////////////////////////////////////////////////////////////
CmdSliceCasts
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdSliceCasts

```go
func NewCmdSliceCasts(S memory.Pointer, L uint32) *CmdSliceCasts
```

#### func (*CmdSliceCasts) API

```go
func (c *CmdSliceCasts) API() gfxapi.API
```

#### func (*CmdSliceCasts) AddRead

```go
func (a *CmdSliceCasts) AddRead(rng memory.Range, id binary.ID) *CmdSliceCasts
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdSliceCasts pointer is returned so that calls can be chained.

#### func (*CmdSliceCasts) AddWrite

```go
func (a *CmdSliceCasts) AddWrite(rng memory.Range, id binary.ID) *CmdSliceCasts
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdSliceCasts pointer is returned so that calls can be chained.

#### func (*CmdSliceCasts) Class

```go
func (*CmdSliceCasts) Class() binary.Class
```

#### func (*CmdSliceCasts) Flags

```go
func (c *CmdSliceCasts) Flags() atom.Flags
```

#### func (*CmdSliceCasts) Mutate

```go
func (ϟa *CmdSliceCasts) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdSliceCasts) Observations

```go
func (a *CmdSliceCasts) Observations() *atom.Observations
```

#### func (*CmdSliceCasts) Replay

```go
func (ϟa *CmdSliceCasts) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdSliceCasts) String

```go
func (a *CmdSliceCasts) String() string
```

#### func (*CmdSliceCasts) TypeID

```go
func (c *CmdSliceCasts) TypeID() atom.TypeID
```

#### type CmdString

```go
type CmdString struct {
	binary.Generate `display:"cmd_string"`

	Result string
}
```

//////////////////////////////////////////////////////////////////////////////
CmdString
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdString

```go
func NewCmdString(Result string) *CmdString
```

#### func (*CmdString) API

```go
func (c *CmdString) API() gfxapi.API
```

#### func (*CmdString) AddRead

```go
func (a *CmdString) AddRead(rng memory.Range, id binary.ID) *CmdString
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdString pointer is returned so that calls can be chained.

#### func (*CmdString) AddWrite

```go
func (a *CmdString) AddWrite(rng memory.Range, id binary.ID) *CmdString
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdString pointer is returned so that calls can be chained.

#### func (*CmdString) Class

```go
func (*CmdString) Class() binary.Class
```

#### func (*CmdString) Flags

```go
func (c *CmdString) Flags() atom.Flags
```

#### func (*CmdString) Mutate

```go
func (ϟa *CmdString) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdString) Observations

```go
func (a *CmdString) Observations() *atom.Observations
```

#### func (*CmdString) Replay

```go
func (ϟa *CmdString) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdString) String

```go
func (a *CmdString) String() string
```

#### func (*CmdString) TypeID

```go
func (c *CmdString) TypeID() atom.TypeID
```

#### type CmdU16

```go
type CmdU16 struct {
	binary.Generate `display:"cmd_u16"`

	Result uint16
}
```

//////////////////////////////////////////////////////////////////////////////
CmdU16
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdU16

```go
func NewCmdU16(Result uint16) *CmdU16
```

#### func (*CmdU16) API

```go
func (c *CmdU16) API() gfxapi.API
```

#### func (*CmdU16) AddRead

```go
func (a *CmdU16) AddRead(rng memory.Range, id binary.ID) *CmdU16
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdU16 pointer is returned so that calls can be chained.

#### func (*CmdU16) AddWrite

```go
func (a *CmdU16) AddWrite(rng memory.Range, id binary.ID) *CmdU16
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdU16 pointer is returned so that calls can be chained.

#### func (*CmdU16) Class

```go
func (*CmdU16) Class() binary.Class
```

#### func (*CmdU16) Flags

```go
func (c *CmdU16) Flags() atom.Flags
```

#### func (*CmdU16) Mutate

```go
func (ϟa *CmdU16) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdU16) Observations

```go
func (a *CmdU16) Observations() *atom.Observations
```

#### func (*CmdU16) Replay

```go
func (ϟa *CmdU16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdU16) String

```go
func (a *CmdU16) String() string
```

#### func (*CmdU16) TypeID

```go
func (c *CmdU16) TypeID() atom.TypeID
```

#### type CmdU32

```go
type CmdU32 struct {
	binary.Generate `display:"cmd_u32"`

	Result uint32
}
```

//////////////////////////////////////////////////////////////////////////////
CmdU32
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdU32

```go
func NewCmdU32(Result uint32) *CmdU32
```

#### func (*CmdU32) API

```go
func (c *CmdU32) API() gfxapi.API
```

#### func (*CmdU32) AddRead

```go
func (a *CmdU32) AddRead(rng memory.Range, id binary.ID) *CmdU32
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdU32 pointer is returned so that calls can be chained.

#### func (*CmdU32) AddWrite

```go
func (a *CmdU32) AddWrite(rng memory.Range, id binary.ID) *CmdU32
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdU32 pointer is returned so that calls can be chained.

#### func (*CmdU32) Class

```go
func (*CmdU32) Class() binary.Class
```

#### func (*CmdU32) Flags

```go
func (c *CmdU32) Flags() atom.Flags
```

#### func (*CmdU32) Mutate

```go
func (ϟa *CmdU32) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdU32) Observations

```go
func (a *CmdU32) Observations() *atom.Observations
```

#### func (*CmdU32) Replay

```go
func (ϟa *CmdU32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdU32) String

```go
func (a *CmdU32) String() string
```

#### func (*CmdU32) TypeID

```go
func (c *CmdU32) TypeID() atom.TypeID
```

#### type CmdU64

```go
type CmdU64 struct {
	binary.Generate `display:"cmd_u64"`

	Result uint64
}
```

//////////////////////////////////////////////////////////////////////////////
CmdU64
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdU64

```go
func NewCmdU64(Result uint64) *CmdU64
```

#### func (*CmdU64) API

```go
func (c *CmdU64) API() gfxapi.API
```

#### func (*CmdU64) AddRead

```go
func (a *CmdU64) AddRead(rng memory.Range, id binary.ID) *CmdU64
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdU64 pointer is returned so that calls can be chained.

#### func (*CmdU64) AddWrite

```go
func (a *CmdU64) AddWrite(rng memory.Range, id binary.ID) *CmdU64
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdU64 pointer is returned so that calls can be chained.

#### func (*CmdU64) Class

```go
func (*CmdU64) Class() binary.Class
```

#### func (*CmdU64) Flags

```go
func (c *CmdU64) Flags() atom.Flags
```

#### func (*CmdU64) Mutate

```go
func (ϟa *CmdU64) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdU64) Observations

```go
func (a *CmdU64) Observations() *atom.Observations
```

#### func (*CmdU64) Replay

```go
func (ϟa *CmdU64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdU64) String

```go
func (a *CmdU64) String() string
```

#### func (*CmdU64) TypeID

```go
func (c *CmdU64) TypeID() atom.TypeID
```

#### type CmdU8

```go
type CmdU8 struct {
	binary.Generate `display:"cmd_u8"`

	Result uint8
}
```

//////////////////////////////////////////////////////////////////////////////
CmdU8
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdU8

```go
func NewCmdU8(Result uint8) *CmdU8
```

#### func (*CmdU8) API

```go
func (c *CmdU8) API() gfxapi.API
```

#### func (*CmdU8) AddRead

```go
func (a *CmdU8) AddRead(rng memory.Range, id binary.ID) *CmdU8
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdU8 pointer is returned so that calls can be chained.

#### func (*CmdU8) AddWrite

```go
func (a *CmdU8) AddWrite(rng memory.Range, id binary.ID) *CmdU8
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdU8 pointer is returned so that calls can be chained.

#### func (*CmdU8) Class

```go
func (*CmdU8) Class() binary.Class
```

#### func (*CmdU8) Flags

```go
func (c *CmdU8) Flags() atom.Flags
```

#### func (*CmdU8) Mutate

```go
func (ϟa *CmdU8) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdU8) Observations

```go
func (a *CmdU8) Observations() *atom.Observations
```

#### func (*CmdU8) Replay

```go
func (ϟa *CmdU8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdU8) String

```go
func (a *CmdU8) String() string
```

#### func (*CmdU8) TypeID

```go
func (c *CmdU8) TypeID() atom.TypeID
```

#### type CmdUnknownRet

```go
type CmdUnknownRet struct {
	binary.Generate `display:"cmd_unknown_ret"`

	Result int64
}
```

//////////////////////////////////////////////////////////////////////////////
CmdUnknownRet
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdUnknownRet

```go
func NewCmdUnknownRet(Result int64) *CmdUnknownRet
```

#### func (*CmdUnknownRet) API

```go
func (c *CmdUnknownRet) API() gfxapi.API
```

#### func (*CmdUnknownRet) AddRead

```go
func (a *CmdUnknownRet) AddRead(rng memory.Range, id binary.ID) *CmdUnknownRet
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdUnknownRet pointer is returned so that calls can be chained.

#### func (*CmdUnknownRet) AddWrite

```go
func (a *CmdUnknownRet) AddWrite(rng memory.Range, id binary.ID) *CmdUnknownRet
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdUnknownRet pointer is returned so that calls can be chained.

#### func (*CmdUnknownRet) Class

```go
func (*CmdUnknownRet) Class() binary.Class
```

#### func (*CmdUnknownRet) Flags

```go
func (c *CmdUnknownRet) Flags() atom.Flags
```

#### func (*CmdUnknownRet) Mutate

```go
func (ϟa *CmdUnknownRet) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdUnknownRet) Observations

```go
func (a *CmdUnknownRet) Observations() *atom.Observations
```

#### func (*CmdUnknownRet) Replay

```go
func (ϟa *CmdUnknownRet) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdUnknownRet) String

```go
func (a *CmdUnknownRet) String() string
```

#### func (*CmdUnknownRet) TypeID

```go
func (c *CmdUnknownRet) TypeID() atom.TypeID
```

#### type CmdUnknownWritePtr

```go
type CmdUnknownWritePtr struct {
	binary.Generate `display:"cmd_unknown_write_ptr"`

	P Intᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
CmdUnknownWritePtr
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdUnknownWritePtr

```go
func NewCmdUnknownWritePtr(P memory.Pointer) *CmdUnknownWritePtr
```

#### func (*CmdUnknownWritePtr) API

```go
func (c *CmdUnknownWritePtr) API() gfxapi.API
```

#### func (*CmdUnknownWritePtr) AddRead

```go
func (a *CmdUnknownWritePtr) AddRead(rng memory.Range, id binary.ID) *CmdUnknownWritePtr
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdUnknownWritePtr pointer is returned so that calls can be
chained.

#### func (*CmdUnknownWritePtr) AddWrite

```go
func (a *CmdUnknownWritePtr) AddWrite(rng memory.Range, id binary.ID) *CmdUnknownWritePtr
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdUnknownWritePtr pointer is returned so that calls can be
chained.

#### func (*CmdUnknownWritePtr) Class

```go
func (*CmdUnknownWritePtr) Class() binary.Class
```

#### func (*CmdUnknownWritePtr) Flags

```go
func (c *CmdUnknownWritePtr) Flags() atom.Flags
```

#### func (*CmdUnknownWritePtr) Mutate

```go
func (ϟa *CmdUnknownWritePtr) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdUnknownWritePtr) Observations

```go
func (a *CmdUnknownWritePtr) Observations() *atom.Observations
```

#### func (*CmdUnknownWritePtr) Replay

```go
func (ϟa *CmdUnknownWritePtr) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdUnknownWritePtr) String

```go
func (a *CmdUnknownWritePtr) String() string
```

#### func (*CmdUnknownWritePtr) TypeID

```go
func (c *CmdUnknownWritePtr) TypeID() atom.TypeID
```

#### type CmdUnknownWriteSlice

```go
type CmdUnknownWriteSlice struct {
	binary.Generate `display:"cmd_unknown_write_slice"`

	A Intᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
CmdUnknownWriteSlice
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdUnknownWriteSlice

```go
func NewCmdUnknownWriteSlice(A memory.Pointer) *CmdUnknownWriteSlice
```

#### func (*CmdUnknownWriteSlice) API

```go
func (c *CmdUnknownWriteSlice) API() gfxapi.API
```

#### func (*CmdUnknownWriteSlice) AddRead

```go
func (a *CmdUnknownWriteSlice) AddRead(rng memory.Range, id binary.ID) *CmdUnknownWriteSlice
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdUnknownWriteSlice pointer is returned so that calls can be
chained.

#### func (*CmdUnknownWriteSlice) AddWrite

```go
func (a *CmdUnknownWriteSlice) AddWrite(rng memory.Range, id binary.ID) *CmdUnknownWriteSlice
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdUnknownWriteSlice pointer is returned so that calls can be
chained.

#### func (*CmdUnknownWriteSlice) Class

```go
func (*CmdUnknownWriteSlice) Class() binary.Class
```

#### func (*CmdUnknownWriteSlice) Flags

```go
func (c *CmdUnknownWriteSlice) Flags() atom.Flags
```

#### func (*CmdUnknownWriteSlice) Mutate

```go
func (ϟa *CmdUnknownWriteSlice) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdUnknownWriteSlice) Observations

```go
func (a *CmdUnknownWriteSlice) Observations() *atom.Observations
```

#### func (*CmdUnknownWriteSlice) Replay

```go
func (ϟa *CmdUnknownWriteSlice) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdUnknownWriteSlice) String

```go
func (a *CmdUnknownWriteSlice) String() string
```

#### func (*CmdUnknownWriteSlice) TypeID

```go
func (c *CmdUnknownWriteSlice) TypeID() atom.TypeID
```

#### type CmdVoid

```go
type CmdVoid struct {
	binary.Generate `display:"cmd_void"`
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoid
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoid

```go
func NewCmdVoid() *CmdVoid
```

#### func (*CmdVoid) API

```go
func (c *CmdVoid) API() gfxapi.API
```

#### func (*CmdVoid) AddRead

```go
func (a *CmdVoid) AddRead(rng memory.Range, id binary.ID) *CmdVoid
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoid pointer is returned so that calls can be chained.

#### func (*CmdVoid) AddWrite

```go
func (a *CmdVoid) AddWrite(rng memory.Range, id binary.ID) *CmdVoid
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoid pointer is returned so that calls can be chained.

#### func (*CmdVoid) Class

```go
func (*CmdVoid) Class() binary.Class
```

#### func (*CmdVoid) Flags

```go
func (c *CmdVoid) Flags() atom.Flags
```

#### func (*CmdVoid) Mutate

```go
func (ϟa *CmdVoid) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoid) Observations

```go
func (a *CmdVoid) Observations() *atom.Observations
```

#### func (*CmdVoid) Replay

```go
func (ϟa *CmdVoid) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoid) String

```go
func (a *CmdVoid) String() string
```

#### func (*CmdVoid) TypeID

```go
func (c *CmdVoid) TypeID() atom.TypeID
```

#### type CmdVoid3InArrays

```go
type CmdVoid3InArrays struct {
	binary.Generate `display:"cmd_void_3_in_arrays"`

	A U8ᵖ
	B U32ᵖ
	C Intᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoid3InArrays
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoid3InArrays

```go
func NewCmdVoid3InArrays(A memory.Pointer, B memory.Pointer, C memory.Pointer) *CmdVoid3InArrays
```

#### func (*CmdVoid3InArrays) API

```go
func (c *CmdVoid3InArrays) API() gfxapi.API
```

#### func (*CmdVoid3InArrays) AddRead

```go
func (a *CmdVoid3InArrays) AddRead(rng memory.Range, id binary.ID) *CmdVoid3InArrays
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoid3InArrays pointer is returned so that calls can be chained.

#### func (*CmdVoid3InArrays) AddWrite

```go
func (a *CmdVoid3InArrays) AddWrite(rng memory.Range, id binary.ID) *CmdVoid3InArrays
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoid3InArrays pointer is returned so that calls can be chained.

#### func (*CmdVoid3InArrays) Class

```go
func (*CmdVoid3InArrays) Class() binary.Class
```

#### func (*CmdVoid3InArrays) Flags

```go
func (c *CmdVoid3InArrays) Flags() atom.Flags
```

#### func (*CmdVoid3InArrays) Mutate

```go
func (ϟa *CmdVoid3InArrays) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoid3InArrays) Observations

```go
func (a *CmdVoid3InArrays) Observations() *atom.Observations
```

#### func (*CmdVoid3InArrays) Replay

```go
func (ϟa *CmdVoid3InArrays) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoid3InArrays) String

```go
func (a *CmdVoid3InArrays) String() string
```

#### func (*CmdVoid3InArrays) TypeID

```go
func (c *CmdVoid3InArrays) TypeID() atom.TypeID
```

#### type CmdVoid3Remapped

```go
type CmdVoid3Remapped struct {
	binary.Generate `display:"cmd_void_3_remapped"`

	A remapped
	B remapped
	C remapped
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoid3Remapped
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoid3Remapped

```go
func NewCmdVoid3Remapped(A remapped, B remapped, C remapped) *CmdVoid3Remapped
```

#### func (*CmdVoid3Remapped) API

```go
func (c *CmdVoid3Remapped) API() gfxapi.API
```

#### func (*CmdVoid3Remapped) AddRead

```go
func (a *CmdVoid3Remapped) AddRead(rng memory.Range, id binary.ID) *CmdVoid3Remapped
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoid3Remapped pointer is returned so that calls can be chained.

#### func (*CmdVoid3Remapped) AddWrite

```go
func (a *CmdVoid3Remapped) AddWrite(rng memory.Range, id binary.ID) *CmdVoid3Remapped
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoid3Remapped pointer is returned so that calls can be chained.

#### func (*CmdVoid3Remapped) Class

```go
func (*CmdVoid3Remapped) Class() binary.Class
```

#### func (*CmdVoid3Remapped) Flags

```go
func (c *CmdVoid3Remapped) Flags() atom.Flags
```

#### func (*CmdVoid3Remapped) Mutate

```go
func (ϟa *CmdVoid3Remapped) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoid3Remapped) Observations

```go
func (a *CmdVoid3Remapped) Observations() *atom.Observations
```

#### func (*CmdVoid3Remapped) Replay

```go
func (ϟa *CmdVoid3Remapped) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoid3Remapped) String

```go
func (a *CmdVoid3Remapped) String() string
```

#### func (*CmdVoid3Remapped) TypeID

```go
func (c *CmdVoid3Remapped) TypeID() atom.TypeID
```

#### type CmdVoid3Strings

```go
type CmdVoid3Strings struct {
	binary.Generate `display:"cmd_void_3_strings"`

	A string
	B string
	C string
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoid3Strings
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoid3Strings

```go
func NewCmdVoid3Strings(A string, B string, C string) *CmdVoid3Strings
```

#### func (*CmdVoid3Strings) API

```go
func (c *CmdVoid3Strings) API() gfxapi.API
```

#### func (*CmdVoid3Strings) AddRead

```go
func (a *CmdVoid3Strings) AddRead(rng memory.Range, id binary.ID) *CmdVoid3Strings
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoid3Strings pointer is returned so that calls can be chained.

#### func (*CmdVoid3Strings) AddWrite

```go
func (a *CmdVoid3Strings) AddWrite(rng memory.Range, id binary.ID) *CmdVoid3Strings
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoid3Strings pointer is returned so that calls can be chained.

#### func (*CmdVoid3Strings) Class

```go
func (*CmdVoid3Strings) Class() binary.Class
```

#### func (*CmdVoid3Strings) Flags

```go
func (c *CmdVoid3Strings) Flags() atom.Flags
```

#### func (*CmdVoid3Strings) Mutate

```go
func (ϟa *CmdVoid3Strings) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoid3Strings) Observations

```go
func (a *CmdVoid3Strings) Observations() *atom.Observations
```

#### func (*CmdVoid3Strings) Replay

```go
func (ϟa *CmdVoid3Strings) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoid3Strings) String

```go
func (a *CmdVoid3Strings) String() string
```

#### func (*CmdVoid3Strings) TypeID

```go
func (c *CmdVoid3Strings) TypeID() atom.TypeID
```

#### type CmdVoidBool

```go
type CmdVoidBool struct {
	binary.Generate `display:"cmd_void_bool"`

	A bool
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidBool
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidBool

```go
func NewCmdVoidBool(A bool) *CmdVoidBool
```

#### func (*CmdVoidBool) API

```go
func (c *CmdVoidBool) API() gfxapi.API
```

#### func (*CmdVoidBool) AddRead

```go
func (a *CmdVoidBool) AddRead(rng memory.Range, id binary.ID) *CmdVoidBool
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidBool pointer is returned so that calls can be chained.

#### func (*CmdVoidBool) AddWrite

```go
func (a *CmdVoidBool) AddWrite(rng memory.Range, id binary.ID) *CmdVoidBool
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidBool pointer is returned so that calls can be chained.

#### func (*CmdVoidBool) Class

```go
func (*CmdVoidBool) Class() binary.Class
```

#### func (*CmdVoidBool) Flags

```go
func (c *CmdVoidBool) Flags() atom.Flags
```

#### func (*CmdVoidBool) Mutate

```go
func (ϟa *CmdVoidBool) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidBool) Observations

```go
func (a *CmdVoidBool) Observations() *atom.Observations
```

#### func (*CmdVoidBool) Replay

```go
func (ϟa *CmdVoidBool) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidBool) String

```go
func (a *CmdVoidBool) String() string
```

#### func (*CmdVoidBool) TypeID

```go
func (c *CmdVoidBool) TypeID() atom.TypeID
```

#### type CmdVoidF32

```go
type CmdVoidF32 struct {
	binary.Generate `display:"cmd_void_f32"`

	A float32
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidF32
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidF32

```go
func NewCmdVoidF32(A float32) *CmdVoidF32
```

#### func (*CmdVoidF32) API

```go
func (c *CmdVoidF32) API() gfxapi.API
```

#### func (*CmdVoidF32) AddRead

```go
func (a *CmdVoidF32) AddRead(rng memory.Range, id binary.ID) *CmdVoidF32
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidF32 pointer is returned so that calls can be chained.

#### func (*CmdVoidF32) AddWrite

```go
func (a *CmdVoidF32) AddWrite(rng memory.Range, id binary.ID) *CmdVoidF32
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidF32 pointer is returned so that calls can be chained.

#### func (*CmdVoidF32) Class

```go
func (*CmdVoidF32) Class() binary.Class
```

#### func (*CmdVoidF32) Flags

```go
func (c *CmdVoidF32) Flags() atom.Flags
```

#### func (*CmdVoidF32) Mutate

```go
func (ϟa *CmdVoidF32) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidF32) Observations

```go
func (a *CmdVoidF32) Observations() *atom.Observations
```

#### func (*CmdVoidF32) Replay

```go
func (ϟa *CmdVoidF32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidF32) String

```go
func (a *CmdVoidF32) String() string
```

#### func (*CmdVoidF32) TypeID

```go
func (c *CmdVoidF32) TypeID() atom.TypeID
```

#### type CmdVoidF64

```go
type CmdVoidF64 struct {
	binary.Generate `display:"cmd_void_f64"`

	A float64
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidF64
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidF64

```go
func NewCmdVoidF64(A float64) *CmdVoidF64
```

#### func (*CmdVoidF64) API

```go
func (c *CmdVoidF64) API() gfxapi.API
```

#### func (*CmdVoidF64) AddRead

```go
func (a *CmdVoidF64) AddRead(rng memory.Range, id binary.ID) *CmdVoidF64
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidF64 pointer is returned so that calls can be chained.

#### func (*CmdVoidF64) AddWrite

```go
func (a *CmdVoidF64) AddWrite(rng memory.Range, id binary.ID) *CmdVoidF64
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidF64 pointer is returned so that calls can be chained.

#### func (*CmdVoidF64) Class

```go
func (*CmdVoidF64) Class() binary.Class
```

#### func (*CmdVoidF64) Flags

```go
func (c *CmdVoidF64) Flags() atom.Flags
```

#### func (*CmdVoidF64) Mutate

```go
func (ϟa *CmdVoidF64) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidF64) Observations

```go
func (a *CmdVoidF64) Observations() *atom.Observations
```

#### func (*CmdVoidF64) Replay

```go
func (ϟa *CmdVoidF64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidF64) String

```go
func (a *CmdVoidF64) String() string
```

#### func (*CmdVoidF64) TypeID

```go
func (c *CmdVoidF64) TypeID() atom.TypeID
```

#### type CmdVoidInArrayOfPointers

```go
type CmdVoidInArrayOfPointers struct {
	binary.Generate `display:"cmd_void_in_array_of_pointers"`

	A     Charᵖᵖ
	Count int32
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidInArrayOfPointers
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidInArrayOfPointers

```go
func NewCmdVoidInArrayOfPointers(A memory.Pointer, Count int32) *CmdVoidInArrayOfPointers
```

#### func (*CmdVoidInArrayOfPointers) API

```go
func (c *CmdVoidInArrayOfPointers) API() gfxapi.API
```

#### func (*CmdVoidInArrayOfPointers) AddRead

```go
func (a *CmdVoidInArrayOfPointers) AddRead(rng memory.Range, id binary.ID) *CmdVoidInArrayOfPointers
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidInArrayOfPointers pointer is returned so that calls can be
chained.

#### func (*CmdVoidInArrayOfPointers) AddWrite

```go
func (a *CmdVoidInArrayOfPointers) AddWrite(rng memory.Range, id binary.ID) *CmdVoidInArrayOfPointers
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidInArrayOfPointers pointer is returned so that calls can be
chained.

#### func (*CmdVoidInArrayOfPointers) Class

```go
func (*CmdVoidInArrayOfPointers) Class() binary.Class
```

#### func (*CmdVoidInArrayOfPointers) Flags

```go
func (c *CmdVoidInArrayOfPointers) Flags() atom.Flags
```

#### func (*CmdVoidInArrayOfPointers) Mutate

```go
func (ϟa *CmdVoidInArrayOfPointers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidInArrayOfPointers) Observations

```go
func (a *CmdVoidInArrayOfPointers) Observations() *atom.Observations
```

#### func (*CmdVoidInArrayOfPointers) Replay

```go
func (ϟa *CmdVoidInArrayOfPointers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidInArrayOfPointers) String

```go
func (a *CmdVoidInArrayOfPointers) String() string
```

#### func (*CmdVoidInArrayOfPointers) TypeID

```go
func (c *CmdVoidInArrayOfPointers) TypeID() atom.TypeID
```

#### type CmdVoidInArrayOfRemapped

```go
type CmdVoidInArrayOfRemapped struct {
	binary.Generate `display:"cmd_void_in_array_of_remapped"`

	A Remappedᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidInArrayOfRemapped
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidInArrayOfRemapped

```go
func NewCmdVoidInArrayOfRemapped(A memory.Pointer) *CmdVoidInArrayOfRemapped
```

#### func (*CmdVoidInArrayOfRemapped) API

```go
func (c *CmdVoidInArrayOfRemapped) API() gfxapi.API
```

#### func (*CmdVoidInArrayOfRemapped) AddRead

```go
func (a *CmdVoidInArrayOfRemapped) AddRead(rng memory.Range, id binary.ID) *CmdVoidInArrayOfRemapped
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidInArrayOfRemapped pointer is returned so that calls can be
chained.

#### func (*CmdVoidInArrayOfRemapped) AddWrite

```go
func (a *CmdVoidInArrayOfRemapped) AddWrite(rng memory.Range, id binary.ID) *CmdVoidInArrayOfRemapped
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidInArrayOfRemapped pointer is returned so that calls can be
chained.

#### func (*CmdVoidInArrayOfRemapped) Class

```go
func (*CmdVoidInArrayOfRemapped) Class() binary.Class
```

#### func (*CmdVoidInArrayOfRemapped) Flags

```go
func (c *CmdVoidInArrayOfRemapped) Flags() atom.Flags
```

#### func (*CmdVoidInArrayOfRemapped) Mutate

```go
func (ϟa *CmdVoidInArrayOfRemapped) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidInArrayOfRemapped) Observations

```go
func (a *CmdVoidInArrayOfRemapped) Observations() *atom.Observations
```

#### func (*CmdVoidInArrayOfRemapped) Replay

```go
func (ϟa *CmdVoidInArrayOfRemapped) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidInArrayOfRemapped) String

```go
func (a *CmdVoidInArrayOfRemapped) String() string
```

#### func (*CmdVoidInArrayOfRemapped) TypeID

```go
func (c *CmdVoidInArrayOfRemapped) TypeID() atom.TypeID
```

#### type CmdVoidOutArrayOfRemapped

```go
type CmdVoidOutArrayOfRemapped struct {
	binary.Generate `display:"cmd_void_out_array_of_remapped"`

	A Remappedᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidOutArrayOfRemapped
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidOutArrayOfRemapped

```go
func NewCmdVoidOutArrayOfRemapped(A memory.Pointer) *CmdVoidOutArrayOfRemapped
```

#### func (*CmdVoidOutArrayOfRemapped) API

```go
func (c *CmdVoidOutArrayOfRemapped) API() gfxapi.API
```

#### func (*CmdVoidOutArrayOfRemapped) AddRead

```go
func (a *CmdVoidOutArrayOfRemapped) AddRead(rng memory.Range, id binary.ID) *CmdVoidOutArrayOfRemapped
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidOutArrayOfRemapped pointer is returned so that calls can be
chained.

#### func (*CmdVoidOutArrayOfRemapped) AddWrite

```go
func (a *CmdVoidOutArrayOfRemapped) AddWrite(rng memory.Range, id binary.ID) *CmdVoidOutArrayOfRemapped
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidOutArrayOfRemapped pointer is returned so that calls can be
chained.

#### func (*CmdVoidOutArrayOfRemapped) Class

```go
func (*CmdVoidOutArrayOfRemapped) Class() binary.Class
```

#### func (*CmdVoidOutArrayOfRemapped) Flags

```go
func (c *CmdVoidOutArrayOfRemapped) Flags() atom.Flags
```

#### func (*CmdVoidOutArrayOfRemapped) Mutate

```go
func (ϟa *CmdVoidOutArrayOfRemapped) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidOutArrayOfRemapped) Observations

```go
func (a *CmdVoidOutArrayOfRemapped) Observations() *atom.Observations
```

#### func (*CmdVoidOutArrayOfRemapped) Replay

```go
func (ϟa *CmdVoidOutArrayOfRemapped) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidOutArrayOfRemapped) String

```go
func (a *CmdVoidOutArrayOfRemapped) String() string
```

#### func (*CmdVoidOutArrayOfRemapped) TypeID

```go
func (c *CmdVoidOutArrayOfRemapped) TypeID() atom.TypeID
```

#### type CmdVoidOutArrayOfUnknownRemapped

```go
type CmdVoidOutArrayOfUnknownRemapped struct {
	binary.Generate `display:"cmd_void_out_array_of_unknown_remapped"`

	A Remappedᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidOutArrayOfUnknownRemapped
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidOutArrayOfUnknownRemapped

```go
func NewCmdVoidOutArrayOfUnknownRemapped(A memory.Pointer) *CmdVoidOutArrayOfUnknownRemapped
```

#### func (*CmdVoidOutArrayOfUnknownRemapped) API

```go
func (c *CmdVoidOutArrayOfUnknownRemapped) API() gfxapi.API
```

#### func (*CmdVoidOutArrayOfUnknownRemapped) AddRead

```go
func (a *CmdVoidOutArrayOfUnknownRemapped) AddRead(rng memory.Range, id binary.ID) *CmdVoidOutArrayOfUnknownRemapped
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidOutArrayOfUnknownRemapped pointer is returned so that calls
can be chained.

#### func (*CmdVoidOutArrayOfUnknownRemapped) AddWrite

```go
func (a *CmdVoidOutArrayOfUnknownRemapped) AddWrite(rng memory.Range, id binary.ID) *CmdVoidOutArrayOfUnknownRemapped
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidOutArrayOfUnknownRemapped pointer is returned so that calls
can be chained.

#### func (*CmdVoidOutArrayOfUnknownRemapped) Class

```go
func (*CmdVoidOutArrayOfUnknownRemapped) Class() binary.Class
```

#### func (*CmdVoidOutArrayOfUnknownRemapped) Flags

```go
func (c *CmdVoidOutArrayOfUnknownRemapped) Flags() atom.Flags
```

#### func (*CmdVoidOutArrayOfUnknownRemapped) Mutate

```go
func (ϟa *CmdVoidOutArrayOfUnknownRemapped) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidOutArrayOfUnknownRemapped) Observations

```go
func (a *CmdVoidOutArrayOfUnknownRemapped) Observations() *atom.Observations
```

#### func (*CmdVoidOutArrayOfUnknownRemapped) Replay

```go
func (ϟa *CmdVoidOutArrayOfUnknownRemapped) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidOutArrayOfUnknownRemapped) String

```go
func (a *CmdVoidOutArrayOfUnknownRemapped) String() string
```

#### func (*CmdVoidOutArrayOfUnknownRemapped) TypeID

```go
func (c *CmdVoidOutArrayOfUnknownRemapped) TypeID() atom.TypeID
```

#### type CmdVoidReadBool

```go
type CmdVoidReadBool struct {
	binary.Generate `display:"cmd_void_read_bool"`

	A Boolᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidReadBool
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidReadBool

```go
func NewCmdVoidReadBool(A memory.Pointer) *CmdVoidReadBool
```

#### func (*CmdVoidReadBool) API

```go
func (c *CmdVoidReadBool) API() gfxapi.API
```

#### func (*CmdVoidReadBool) AddRead

```go
func (a *CmdVoidReadBool) AddRead(rng memory.Range, id binary.ID) *CmdVoidReadBool
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidReadBool pointer is returned so that calls can be chained.

#### func (*CmdVoidReadBool) AddWrite

```go
func (a *CmdVoidReadBool) AddWrite(rng memory.Range, id binary.ID) *CmdVoidReadBool
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidReadBool pointer is returned so that calls can be chained.

#### func (*CmdVoidReadBool) Class

```go
func (*CmdVoidReadBool) Class() binary.Class
```

#### func (*CmdVoidReadBool) Flags

```go
func (c *CmdVoidReadBool) Flags() atom.Flags
```

#### func (*CmdVoidReadBool) Mutate

```go
func (ϟa *CmdVoidReadBool) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidReadBool) Observations

```go
func (a *CmdVoidReadBool) Observations() *atom.Observations
```

#### func (*CmdVoidReadBool) Replay

```go
func (ϟa *CmdVoidReadBool) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidReadBool) String

```go
func (a *CmdVoidReadBool) String() string
```

#### func (*CmdVoidReadBool) TypeID

```go
func (c *CmdVoidReadBool) TypeID() atom.TypeID
```

#### type CmdVoidReadF32

```go
type CmdVoidReadF32 struct {
	binary.Generate `display:"cmd_void_read_f32"`

	A F32ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidReadF32
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidReadF32

```go
func NewCmdVoidReadF32(A memory.Pointer) *CmdVoidReadF32
```

#### func (*CmdVoidReadF32) API

```go
func (c *CmdVoidReadF32) API() gfxapi.API
```

#### func (*CmdVoidReadF32) AddRead

```go
func (a *CmdVoidReadF32) AddRead(rng memory.Range, id binary.ID) *CmdVoidReadF32
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidReadF32 pointer is returned so that calls can be chained.

#### func (*CmdVoidReadF32) AddWrite

```go
func (a *CmdVoidReadF32) AddWrite(rng memory.Range, id binary.ID) *CmdVoidReadF32
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidReadF32 pointer is returned so that calls can be chained.

#### func (*CmdVoidReadF32) Class

```go
func (*CmdVoidReadF32) Class() binary.Class
```

#### func (*CmdVoidReadF32) Flags

```go
func (c *CmdVoidReadF32) Flags() atom.Flags
```

#### func (*CmdVoidReadF32) Mutate

```go
func (ϟa *CmdVoidReadF32) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidReadF32) Observations

```go
func (a *CmdVoidReadF32) Observations() *atom.Observations
```

#### func (*CmdVoidReadF32) Replay

```go
func (ϟa *CmdVoidReadF32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidReadF32) String

```go
func (a *CmdVoidReadF32) String() string
```

#### func (*CmdVoidReadF32) TypeID

```go
func (c *CmdVoidReadF32) TypeID() atom.TypeID
```

#### type CmdVoidReadF64

```go
type CmdVoidReadF64 struct {
	binary.Generate `display:"cmd_void_read_f64"`

	A F64ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidReadF64
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidReadF64

```go
func NewCmdVoidReadF64(A memory.Pointer) *CmdVoidReadF64
```

#### func (*CmdVoidReadF64) API

```go
func (c *CmdVoidReadF64) API() gfxapi.API
```

#### func (*CmdVoidReadF64) AddRead

```go
func (a *CmdVoidReadF64) AddRead(rng memory.Range, id binary.ID) *CmdVoidReadF64
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidReadF64 pointer is returned so that calls can be chained.

#### func (*CmdVoidReadF64) AddWrite

```go
func (a *CmdVoidReadF64) AddWrite(rng memory.Range, id binary.ID) *CmdVoidReadF64
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidReadF64 pointer is returned so that calls can be chained.

#### func (*CmdVoidReadF64) Class

```go
func (*CmdVoidReadF64) Class() binary.Class
```

#### func (*CmdVoidReadF64) Flags

```go
func (c *CmdVoidReadF64) Flags() atom.Flags
```

#### func (*CmdVoidReadF64) Mutate

```go
func (ϟa *CmdVoidReadF64) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidReadF64) Observations

```go
func (a *CmdVoidReadF64) Observations() *atom.Observations
```

#### func (*CmdVoidReadF64) Replay

```go
func (ϟa *CmdVoidReadF64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidReadF64) String

```go
func (a *CmdVoidReadF64) String() string
```

#### func (*CmdVoidReadF64) TypeID

```go
func (c *CmdVoidReadF64) TypeID() atom.TypeID
```

#### type CmdVoidReadPtrs

```go
type CmdVoidReadPtrs struct {
	binary.Generate `display:"cmd_void_read_ptrs"`

	A F32ᵖ
	B U16ᵖ
	C Boolᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidReadPtrs
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidReadPtrs

```go
func NewCmdVoidReadPtrs(A memory.Pointer, B memory.Pointer, C memory.Pointer) *CmdVoidReadPtrs
```

#### func (*CmdVoidReadPtrs) API

```go
func (c *CmdVoidReadPtrs) API() gfxapi.API
```

#### func (*CmdVoidReadPtrs) AddRead

```go
func (a *CmdVoidReadPtrs) AddRead(rng memory.Range, id binary.ID) *CmdVoidReadPtrs
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidReadPtrs pointer is returned so that calls can be chained.

#### func (*CmdVoidReadPtrs) AddWrite

```go
func (a *CmdVoidReadPtrs) AddWrite(rng memory.Range, id binary.ID) *CmdVoidReadPtrs
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidReadPtrs pointer is returned so that calls can be chained.

#### func (*CmdVoidReadPtrs) Class

```go
func (*CmdVoidReadPtrs) Class() binary.Class
```

#### func (*CmdVoidReadPtrs) Flags

```go
func (c *CmdVoidReadPtrs) Flags() atom.Flags
```

#### func (*CmdVoidReadPtrs) Mutate

```go
func (ϟa *CmdVoidReadPtrs) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidReadPtrs) Observations

```go
func (a *CmdVoidReadPtrs) Observations() *atom.Observations
```

#### func (*CmdVoidReadPtrs) Replay

```go
func (ϟa *CmdVoidReadPtrs) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidReadPtrs) String

```go
func (a *CmdVoidReadPtrs) String() string
```

#### func (*CmdVoidReadPtrs) TypeID

```go
func (c *CmdVoidReadPtrs) TypeID() atom.TypeID
```

#### type CmdVoidReadS16

```go
type CmdVoidReadS16 struct {
	binary.Generate `display:"cmd_void_read_s16"`

	A S16ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidReadS16
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidReadS16

```go
func NewCmdVoidReadS16(A memory.Pointer) *CmdVoidReadS16
```

#### func (*CmdVoidReadS16) API

```go
func (c *CmdVoidReadS16) API() gfxapi.API
```

#### func (*CmdVoidReadS16) AddRead

```go
func (a *CmdVoidReadS16) AddRead(rng memory.Range, id binary.ID) *CmdVoidReadS16
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidReadS16 pointer is returned so that calls can be chained.

#### func (*CmdVoidReadS16) AddWrite

```go
func (a *CmdVoidReadS16) AddWrite(rng memory.Range, id binary.ID) *CmdVoidReadS16
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidReadS16 pointer is returned so that calls can be chained.

#### func (*CmdVoidReadS16) Class

```go
func (*CmdVoidReadS16) Class() binary.Class
```

#### func (*CmdVoidReadS16) Flags

```go
func (c *CmdVoidReadS16) Flags() atom.Flags
```

#### func (*CmdVoidReadS16) Mutate

```go
func (ϟa *CmdVoidReadS16) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidReadS16) Observations

```go
func (a *CmdVoidReadS16) Observations() *atom.Observations
```

#### func (*CmdVoidReadS16) Replay

```go
func (ϟa *CmdVoidReadS16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidReadS16) String

```go
func (a *CmdVoidReadS16) String() string
```

#### func (*CmdVoidReadS16) TypeID

```go
func (c *CmdVoidReadS16) TypeID() atom.TypeID
```

#### type CmdVoidReadS32

```go
type CmdVoidReadS32 struct {
	binary.Generate `display:"cmd_void_read_s32"`

	A S32ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidReadS32
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidReadS32

```go
func NewCmdVoidReadS32(A memory.Pointer) *CmdVoidReadS32
```

#### func (*CmdVoidReadS32) API

```go
func (c *CmdVoidReadS32) API() gfxapi.API
```

#### func (*CmdVoidReadS32) AddRead

```go
func (a *CmdVoidReadS32) AddRead(rng memory.Range, id binary.ID) *CmdVoidReadS32
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidReadS32 pointer is returned so that calls can be chained.

#### func (*CmdVoidReadS32) AddWrite

```go
func (a *CmdVoidReadS32) AddWrite(rng memory.Range, id binary.ID) *CmdVoidReadS32
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidReadS32 pointer is returned so that calls can be chained.

#### func (*CmdVoidReadS32) Class

```go
func (*CmdVoidReadS32) Class() binary.Class
```

#### func (*CmdVoidReadS32) Flags

```go
func (c *CmdVoidReadS32) Flags() atom.Flags
```

#### func (*CmdVoidReadS32) Mutate

```go
func (ϟa *CmdVoidReadS32) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidReadS32) Observations

```go
func (a *CmdVoidReadS32) Observations() *atom.Observations
```

#### func (*CmdVoidReadS32) Replay

```go
func (ϟa *CmdVoidReadS32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidReadS32) String

```go
func (a *CmdVoidReadS32) String() string
```

#### func (*CmdVoidReadS32) TypeID

```go
func (c *CmdVoidReadS32) TypeID() atom.TypeID
```

#### type CmdVoidReadS64

```go
type CmdVoidReadS64 struct {
	binary.Generate `display:"cmd_void_read_s64"`

	A S64ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidReadS64
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidReadS64

```go
func NewCmdVoidReadS64(A memory.Pointer) *CmdVoidReadS64
```

#### func (*CmdVoidReadS64) API

```go
func (c *CmdVoidReadS64) API() gfxapi.API
```

#### func (*CmdVoidReadS64) AddRead

```go
func (a *CmdVoidReadS64) AddRead(rng memory.Range, id binary.ID) *CmdVoidReadS64
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidReadS64 pointer is returned so that calls can be chained.

#### func (*CmdVoidReadS64) AddWrite

```go
func (a *CmdVoidReadS64) AddWrite(rng memory.Range, id binary.ID) *CmdVoidReadS64
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidReadS64 pointer is returned so that calls can be chained.

#### func (*CmdVoidReadS64) Class

```go
func (*CmdVoidReadS64) Class() binary.Class
```

#### func (*CmdVoidReadS64) Flags

```go
func (c *CmdVoidReadS64) Flags() atom.Flags
```

#### func (*CmdVoidReadS64) Mutate

```go
func (ϟa *CmdVoidReadS64) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidReadS64) Observations

```go
func (a *CmdVoidReadS64) Observations() *atom.Observations
```

#### func (*CmdVoidReadS64) Replay

```go
func (ϟa *CmdVoidReadS64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidReadS64) String

```go
func (a *CmdVoidReadS64) String() string
```

#### func (*CmdVoidReadS64) TypeID

```go
func (c *CmdVoidReadS64) TypeID() atom.TypeID
```

#### type CmdVoidReadS8

```go
type CmdVoidReadS8 struct {
	binary.Generate `display:"cmd_void_read_s8"`

	A S8ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidReadS8
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidReadS8

```go
func NewCmdVoidReadS8(A memory.Pointer) *CmdVoidReadS8
```

#### func (*CmdVoidReadS8) API

```go
func (c *CmdVoidReadS8) API() gfxapi.API
```

#### func (*CmdVoidReadS8) AddRead

```go
func (a *CmdVoidReadS8) AddRead(rng memory.Range, id binary.ID) *CmdVoidReadS8
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidReadS8 pointer is returned so that calls can be chained.

#### func (*CmdVoidReadS8) AddWrite

```go
func (a *CmdVoidReadS8) AddWrite(rng memory.Range, id binary.ID) *CmdVoidReadS8
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidReadS8 pointer is returned so that calls can be chained.

#### func (*CmdVoidReadS8) Class

```go
func (*CmdVoidReadS8) Class() binary.Class
```

#### func (*CmdVoidReadS8) Flags

```go
func (c *CmdVoidReadS8) Flags() atom.Flags
```

#### func (*CmdVoidReadS8) Mutate

```go
func (ϟa *CmdVoidReadS8) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidReadS8) Observations

```go
func (a *CmdVoidReadS8) Observations() *atom.Observations
```

#### func (*CmdVoidReadS8) Replay

```go
func (ϟa *CmdVoidReadS8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidReadS8) String

```go
func (a *CmdVoidReadS8) String() string
```

#### func (*CmdVoidReadS8) TypeID

```go
func (c *CmdVoidReadS8) TypeID() atom.TypeID
```

#### type CmdVoidReadU16

```go
type CmdVoidReadU16 struct {
	binary.Generate `display:"cmd_void_read_u16"`

	A U16ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidReadU16
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidReadU16

```go
func NewCmdVoidReadU16(A memory.Pointer) *CmdVoidReadU16
```

#### func (*CmdVoidReadU16) API

```go
func (c *CmdVoidReadU16) API() gfxapi.API
```

#### func (*CmdVoidReadU16) AddRead

```go
func (a *CmdVoidReadU16) AddRead(rng memory.Range, id binary.ID) *CmdVoidReadU16
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidReadU16 pointer is returned so that calls can be chained.

#### func (*CmdVoidReadU16) AddWrite

```go
func (a *CmdVoidReadU16) AddWrite(rng memory.Range, id binary.ID) *CmdVoidReadU16
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidReadU16 pointer is returned so that calls can be chained.

#### func (*CmdVoidReadU16) Class

```go
func (*CmdVoidReadU16) Class() binary.Class
```

#### func (*CmdVoidReadU16) Flags

```go
func (c *CmdVoidReadU16) Flags() atom.Flags
```

#### func (*CmdVoidReadU16) Mutate

```go
func (ϟa *CmdVoidReadU16) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidReadU16) Observations

```go
func (a *CmdVoidReadU16) Observations() *atom.Observations
```

#### func (*CmdVoidReadU16) Replay

```go
func (ϟa *CmdVoidReadU16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidReadU16) String

```go
func (a *CmdVoidReadU16) String() string
```

#### func (*CmdVoidReadU16) TypeID

```go
func (c *CmdVoidReadU16) TypeID() atom.TypeID
```

#### type CmdVoidReadU32

```go
type CmdVoidReadU32 struct {
	binary.Generate `display:"cmd_void_read_u32"`

	A U32ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidReadU32
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidReadU32

```go
func NewCmdVoidReadU32(A memory.Pointer) *CmdVoidReadU32
```

#### func (*CmdVoidReadU32) API

```go
func (c *CmdVoidReadU32) API() gfxapi.API
```

#### func (*CmdVoidReadU32) AddRead

```go
func (a *CmdVoidReadU32) AddRead(rng memory.Range, id binary.ID) *CmdVoidReadU32
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidReadU32 pointer is returned so that calls can be chained.

#### func (*CmdVoidReadU32) AddWrite

```go
func (a *CmdVoidReadU32) AddWrite(rng memory.Range, id binary.ID) *CmdVoidReadU32
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidReadU32 pointer is returned so that calls can be chained.

#### func (*CmdVoidReadU32) Class

```go
func (*CmdVoidReadU32) Class() binary.Class
```

#### func (*CmdVoidReadU32) Flags

```go
func (c *CmdVoidReadU32) Flags() atom.Flags
```

#### func (*CmdVoidReadU32) Mutate

```go
func (ϟa *CmdVoidReadU32) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidReadU32) Observations

```go
func (a *CmdVoidReadU32) Observations() *atom.Observations
```

#### func (*CmdVoidReadU32) Replay

```go
func (ϟa *CmdVoidReadU32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidReadU32) String

```go
func (a *CmdVoidReadU32) String() string
```

#### func (*CmdVoidReadU32) TypeID

```go
func (c *CmdVoidReadU32) TypeID() atom.TypeID
```

#### type CmdVoidReadU64

```go
type CmdVoidReadU64 struct {
	binary.Generate `display:"cmd_void_read_u64"`

	A U64ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidReadU64
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidReadU64

```go
func NewCmdVoidReadU64(A memory.Pointer) *CmdVoidReadU64
```

#### func (*CmdVoidReadU64) API

```go
func (c *CmdVoidReadU64) API() gfxapi.API
```

#### func (*CmdVoidReadU64) AddRead

```go
func (a *CmdVoidReadU64) AddRead(rng memory.Range, id binary.ID) *CmdVoidReadU64
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidReadU64 pointer is returned so that calls can be chained.

#### func (*CmdVoidReadU64) AddWrite

```go
func (a *CmdVoidReadU64) AddWrite(rng memory.Range, id binary.ID) *CmdVoidReadU64
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidReadU64 pointer is returned so that calls can be chained.

#### func (*CmdVoidReadU64) Class

```go
func (*CmdVoidReadU64) Class() binary.Class
```

#### func (*CmdVoidReadU64) Flags

```go
func (c *CmdVoidReadU64) Flags() atom.Flags
```

#### func (*CmdVoidReadU64) Mutate

```go
func (ϟa *CmdVoidReadU64) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidReadU64) Observations

```go
func (a *CmdVoidReadU64) Observations() *atom.Observations
```

#### func (*CmdVoidReadU64) Replay

```go
func (ϟa *CmdVoidReadU64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidReadU64) String

```go
func (a *CmdVoidReadU64) String() string
```

#### func (*CmdVoidReadU64) TypeID

```go
func (c *CmdVoidReadU64) TypeID() atom.TypeID
```

#### type CmdVoidReadU8

```go
type CmdVoidReadU8 struct {
	binary.Generate `display:"cmd_void_read_u8"`

	A U8ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidReadU8
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidReadU8

```go
func NewCmdVoidReadU8(A memory.Pointer) *CmdVoidReadU8
```

#### func (*CmdVoidReadU8) API

```go
func (c *CmdVoidReadU8) API() gfxapi.API
```

#### func (*CmdVoidReadU8) AddRead

```go
func (a *CmdVoidReadU8) AddRead(rng memory.Range, id binary.ID) *CmdVoidReadU8
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidReadU8 pointer is returned so that calls can be chained.

#### func (*CmdVoidReadU8) AddWrite

```go
func (a *CmdVoidReadU8) AddWrite(rng memory.Range, id binary.ID) *CmdVoidReadU8
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidReadU8 pointer is returned so that calls can be chained.

#### func (*CmdVoidReadU8) Class

```go
func (*CmdVoidReadU8) Class() binary.Class
```

#### func (*CmdVoidReadU8) Flags

```go
func (c *CmdVoidReadU8) Flags() atom.Flags
```

#### func (*CmdVoidReadU8) Mutate

```go
func (ϟa *CmdVoidReadU8) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidReadU8) Observations

```go
func (a *CmdVoidReadU8) Observations() *atom.Observations
```

#### func (*CmdVoidReadU8) Replay

```go
func (ϟa *CmdVoidReadU8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidReadU8) String

```go
func (a *CmdVoidReadU8) String() string
```

#### func (*CmdVoidReadU8) TypeID

```go
func (c *CmdVoidReadU8) TypeID() atom.TypeID
```

#### type CmdVoidS16

```go
type CmdVoidS16 struct {
	binary.Generate `display:"cmd_void_s16"`

	A int16
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidS16
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidS16

```go
func NewCmdVoidS16(A int16) *CmdVoidS16
```

#### func (*CmdVoidS16) API

```go
func (c *CmdVoidS16) API() gfxapi.API
```

#### func (*CmdVoidS16) AddRead

```go
func (a *CmdVoidS16) AddRead(rng memory.Range, id binary.ID) *CmdVoidS16
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidS16 pointer is returned so that calls can be chained.

#### func (*CmdVoidS16) AddWrite

```go
func (a *CmdVoidS16) AddWrite(rng memory.Range, id binary.ID) *CmdVoidS16
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidS16 pointer is returned so that calls can be chained.

#### func (*CmdVoidS16) Class

```go
func (*CmdVoidS16) Class() binary.Class
```

#### func (*CmdVoidS16) Flags

```go
func (c *CmdVoidS16) Flags() atom.Flags
```

#### func (*CmdVoidS16) Mutate

```go
func (ϟa *CmdVoidS16) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidS16) Observations

```go
func (a *CmdVoidS16) Observations() *atom.Observations
```

#### func (*CmdVoidS16) Replay

```go
func (ϟa *CmdVoidS16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidS16) String

```go
func (a *CmdVoidS16) String() string
```

#### func (*CmdVoidS16) TypeID

```go
func (c *CmdVoidS16) TypeID() atom.TypeID
```

#### type CmdVoidS32

```go
type CmdVoidS32 struct {
	binary.Generate `display:"cmd_void_s32"`

	A int32
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidS32
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidS32

```go
func NewCmdVoidS32(A int32) *CmdVoidS32
```

#### func (*CmdVoidS32) API

```go
func (c *CmdVoidS32) API() gfxapi.API
```

#### func (*CmdVoidS32) AddRead

```go
func (a *CmdVoidS32) AddRead(rng memory.Range, id binary.ID) *CmdVoidS32
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidS32 pointer is returned so that calls can be chained.

#### func (*CmdVoidS32) AddWrite

```go
func (a *CmdVoidS32) AddWrite(rng memory.Range, id binary.ID) *CmdVoidS32
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidS32 pointer is returned so that calls can be chained.

#### func (*CmdVoidS32) Class

```go
func (*CmdVoidS32) Class() binary.Class
```

#### func (*CmdVoidS32) Flags

```go
func (c *CmdVoidS32) Flags() atom.Flags
```

#### func (*CmdVoidS32) Mutate

```go
func (ϟa *CmdVoidS32) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidS32) Observations

```go
func (a *CmdVoidS32) Observations() *atom.Observations
```

#### func (*CmdVoidS32) Replay

```go
func (ϟa *CmdVoidS32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidS32) String

```go
func (a *CmdVoidS32) String() string
```

#### func (*CmdVoidS32) TypeID

```go
func (c *CmdVoidS32) TypeID() atom.TypeID
```

#### type CmdVoidS64

```go
type CmdVoidS64 struct {
	binary.Generate `display:"cmd_void_s64"`

	A int64
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidS64
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidS64

```go
func NewCmdVoidS64(A int64) *CmdVoidS64
```

#### func (*CmdVoidS64) API

```go
func (c *CmdVoidS64) API() gfxapi.API
```

#### func (*CmdVoidS64) AddRead

```go
func (a *CmdVoidS64) AddRead(rng memory.Range, id binary.ID) *CmdVoidS64
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidS64 pointer is returned so that calls can be chained.

#### func (*CmdVoidS64) AddWrite

```go
func (a *CmdVoidS64) AddWrite(rng memory.Range, id binary.ID) *CmdVoidS64
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidS64 pointer is returned so that calls can be chained.

#### func (*CmdVoidS64) Class

```go
func (*CmdVoidS64) Class() binary.Class
```

#### func (*CmdVoidS64) Flags

```go
func (c *CmdVoidS64) Flags() atom.Flags
```

#### func (*CmdVoidS64) Mutate

```go
func (ϟa *CmdVoidS64) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidS64) Observations

```go
func (a *CmdVoidS64) Observations() *atom.Observations
```

#### func (*CmdVoidS64) Replay

```go
func (ϟa *CmdVoidS64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidS64) String

```go
func (a *CmdVoidS64) String() string
```

#### func (*CmdVoidS64) TypeID

```go
func (c *CmdVoidS64) TypeID() atom.TypeID
```

#### type CmdVoidS8

```go
type CmdVoidS8 struct {
	binary.Generate `display:"cmd_void_s8"`

	A int8
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidS8
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidS8

```go
func NewCmdVoidS8(A int8) *CmdVoidS8
```

#### func (*CmdVoidS8) API

```go
func (c *CmdVoidS8) API() gfxapi.API
```

#### func (*CmdVoidS8) AddRead

```go
func (a *CmdVoidS8) AddRead(rng memory.Range, id binary.ID) *CmdVoidS8
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidS8 pointer is returned so that calls can be chained.

#### func (*CmdVoidS8) AddWrite

```go
func (a *CmdVoidS8) AddWrite(rng memory.Range, id binary.ID) *CmdVoidS8
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidS8 pointer is returned so that calls can be chained.

#### func (*CmdVoidS8) Class

```go
func (*CmdVoidS8) Class() binary.Class
```

#### func (*CmdVoidS8) Flags

```go
func (c *CmdVoidS8) Flags() atom.Flags
```

#### func (*CmdVoidS8) Mutate

```go
func (ϟa *CmdVoidS8) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidS8) Observations

```go
func (a *CmdVoidS8) Observations() *atom.Observations
```

#### func (*CmdVoidS8) Replay

```go
func (ϟa *CmdVoidS8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidS8) String

```go
func (a *CmdVoidS8) String() string
```

#### func (*CmdVoidS8) TypeID

```go
func (c *CmdVoidS8) TypeID() atom.TypeID
```

#### type CmdVoidString

```go
type CmdVoidString struct {
	binary.Generate `display:"cmd_void_string"`

	A string
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidString
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidString

```go
func NewCmdVoidString(A string) *CmdVoidString
```

#### func (*CmdVoidString) API

```go
func (c *CmdVoidString) API() gfxapi.API
```

#### func (*CmdVoidString) AddRead

```go
func (a *CmdVoidString) AddRead(rng memory.Range, id binary.ID) *CmdVoidString
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidString pointer is returned so that calls can be chained.

#### func (*CmdVoidString) AddWrite

```go
func (a *CmdVoidString) AddWrite(rng memory.Range, id binary.ID) *CmdVoidString
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidString pointer is returned so that calls can be chained.

#### func (*CmdVoidString) Class

```go
func (*CmdVoidString) Class() binary.Class
```

#### func (*CmdVoidString) Flags

```go
func (c *CmdVoidString) Flags() atom.Flags
```

#### func (*CmdVoidString) Mutate

```go
func (ϟa *CmdVoidString) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidString) Observations

```go
func (a *CmdVoidString) Observations() *atom.Observations
```

#### func (*CmdVoidString) Replay

```go
func (ϟa *CmdVoidString) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidString) String

```go
func (a *CmdVoidString) String() string
```

#### func (*CmdVoidString) TypeID

```go
func (c *CmdVoidString) TypeID() atom.TypeID
```

#### type CmdVoidU16

```go
type CmdVoidU16 struct {
	binary.Generate `display:"cmd_void_u16"`

	A uint16
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidU16
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidU16

```go
func NewCmdVoidU16(A uint16) *CmdVoidU16
```

#### func (*CmdVoidU16) API

```go
func (c *CmdVoidU16) API() gfxapi.API
```

#### func (*CmdVoidU16) AddRead

```go
func (a *CmdVoidU16) AddRead(rng memory.Range, id binary.ID) *CmdVoidU16
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidU16 pointer is returned so that calls can be chained.

#### func (*CmdVoidU16) AddWrite

```go
func (a *CmdVoidU16) AddWrite(rng memory.Range, id binary.ID) *CmdVoidU16
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidU16 pointer is returned so that calls can be chained.

#### func (*CmdVoidU16) Class

```go
func (*CmdVoidU16) Class() binary.Class
```

#### func (*CmdVoidU16) Flags

```go
func (c *CmdVoidU16) Flags() atom.Flags
```

#### func (*CmdVoidU16) Mutate

```go
func (ϟa *CmdVoidU16) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidU16) Observations

```go
func (a *CmdVoidU16) Observations() *atom.Observations
```

#### func (*CmdVoidU16) Replay

```go
func (ϟa *CmdVoidU16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidU16) String

```go
func (a *CmdVoidU16) String() string
```

#### func (*CmdVoidU16) TypeID

```go
func (c *CmdVoidU16) TypeID() atom.TypeID
```

#### type CmdVoidU32

```go
type CmdVoidU32 struct {
	binary.Generate `display:"cmd_void_u32"`

	A uint32
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidU32
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidU32

```go
func NewCmdVoidU32(A uint32) *CmdVoidU32
```

#### func (*CmdVoidU32) API

```go
func (c *CmdVoidU32) API() gfxapi.API
```

#### func (*CmdVoidU32) AddRead

```go
func (a *CmdVoidU32) AddRead(rng memory.Range, id binary.ID) *CmdVoidU32
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidU32 pointer is returned so that calls can be chained.

#### func (*CmdVoidU32) AddWrite

```go
func (a *CmdVoidU32) AddWrite(rng memory.Range, id binary.ID) *CmdVoidU32
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidU32 pointer is returned so that calls can be chained.

#### func (*CmdVoidU32) Class

```go
func (*CmdVoidU32) Class() binary.Class
```

#### func (*CmdVoidU32) Flags

```go
func (c *CmdVoidU32) Flags() atom.Flags
```

#### func (*CmdVoidU32) Mutate

```go
func (ϟa *CmdVoidU32) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidU32) Observations

```go
func (a *CmdVoidU32) Observations() *atom.Observations
```

#### func (*CmdVoidU32) Replay

```go
func (ϟa *CmdVoidU32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidU32) String

```go
func (a *CmdVoidU32) String() string
```

#### func (*CmdVoidU32) TypeID

```go
func (c *CmdVoidU32) TypeID() atom.TypeID
```

#### type CmdVoidU64

```go
type CmdVoidU64 struct {
	binary.Generate `display:"cmd_void_u64"`

	A uint64
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidU64
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidU64

```go
func NewCmdVoidU64(A uint64) *CmdVoidU64
```

#### func (*CmdVoidU64) API

```go
func (c *CmdVoidU64) API() gfxapi.API
```

#### func (*CmdVoidU64) AddRead

```go
func (a *CmdVoidU64) AddRead(rng memory.Range, id binary.ID) *CmdVoidU64
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidU64 pointer is returned so that calls can be chained.

#### func (*CmdVoidU64) AddWrite

```go
func (a *CmdVoidU64) AddWrite(rng memory.Range, id binary.ID) *CmdVoidU64
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidU64 pointer is returned so that calls can be chained.

#### func (*CmdVoidU64) Class

```go
func (*CmdVoidU64) Class() binary.Class
```

#### func (*CmdVoidU64) Flags

```go
func (c *CmdVoidU64) Flags() atom.Flags
```

#### func (*CmdVoidU64) Mutate

```go
func (ϟa *CmdVoidU64) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidU64) Observations

```go
func (a *CmdVoidU64) Observations() *atom.Observations
```

#### func (*CmdVoidU64) Replay

```go
func (ϟa *CmdVoidU64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidU64) String

```go
func (a *CmdVoidU64) String() string
```

#### func (*CmdVoidU64) TypeID

```go
func (c *CmdVoidU64) TypeID() atom.TypeID
```

#### type CmdVoidU8

```go
type CmdVoidU8 struct {
	binary.Generate `display:"cmd_void_u8"`

	A uint8
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidU8
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidU8

```go
func NewCmdVoidU8(A uint8) *CmdVoidU8
```

#### func (*CmdVoidU8) API

```go
func (c *CmdVoidU8) API() gfxapi.API
```

#### func (*CmdVoidU8) AddRead

```go
func (a *CmdVoidU8) AddRead(rng memory.Range, id binary.ID) *CmdVoidU8
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidU8 pointer is returned so that calls can be chained.

#### func (*CmdVoidU8) AddWrite

```go
func (a *CmdVoidU8) AddWrite(rng memory.Range, id binary.ID) *CmdVoidU8
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidU8 pointer is returned so that calls can be chained.

#### func (*CmdVoidU8) Class

```go
func (*CmdVoidU8) Class() binary.Class
```

#### func (*CmdVoidU8) Flags

```go
func (c *CmdVoidU8) Flags() atom.Flags
```

#### func (*CmdVoidU8) Mutate

```go
func (ϟa *CmdVoidU8) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidU8) Observations

```go
func (a *CmdVoidU8) Observations() *atom.Observations
```

#### func (*CmdVoidU8) Replay

```go
func (ϟa *CmdVoidU8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidU8) String

```go
func (a *CmdVoidU8) String() string
```

#### func (*CmdVoidU8) TypeID

```go
func (c *CmdVoidU8) TypeID() atom.TypeID
```

#### type CmdVoidWriteBool

```go
type CmdVoidWriteBool struct {
	binary.Generate `display:"cmd_void_write_bool"`

	A Boolᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidWriteBool
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidWriteBool

```go
func NewCmdVoidWriteBool(A memory.Pointer) *CmdVoidWriteBool
```

#### func (*CmdVoidWriteBool) API

```go
func (c *CmdVoidWriteBool) API() gfxapi.API
```

#### func (*CmdVoidWriteBool) AddRead

```go
func (a *CmdVoidWriteBool) AddRead(rng memory.Range, id binary.ID) *CmdVoidWriteBool
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidWriteBool pointer is returned so that calls can be chained.

#### func (*CmdVoidWriteBool) AddWrite

```go
func (a *CmdVoidWriteBool) AddWrite(rng memory.Range, id binary.ID) *CmdVoidWriteBool
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidWriteBool pointer is returned so that calls can be chained.

#### func (*CmdVoidWriteBool) Class

```go
func (*CmdVoidWriteBool) Class() binary.Class
```

#### func (*CmdVoidWriteBool) Flags

```go
func (c *CmdVoidWriteBool) Flags() atom.Flags
```

#### func (*CmdVoidWriteBool) Mutate

```go
func (ϟa *CmdVoidWriteBool) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidWriteBool) Observations

```go
func (a *CmdVoidWriteBool) Observations() *atom.Observations
```

#### func (*CmdVoidWriteBool) Replay

```go
func (ϟa *CmdVoidWriteBool) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidWriteBool) String

```go
func (a *CmdVoidWriteBool) String() string
```

#### func (*CmdVoidWriteBool) TypeID

```go
func (c *CmdVoidWriteBool) TypeID() atom.TypeID
```

#### type CmdVoidWriteF32

```go
type CmdVoidWriteF32 struct {
	binary.Generate `display:"cmd_void_write_f32"`

	A F32ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidWriteF32
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidWriteF32

```go
func NewCmdVoidWriteF32(A memory.Pointer) *CmdVoidWriteF32
```

#### func (*CmdVoidWriteF32) API

```go
func (c *CmdVoidWriteF32) API() gfxapi.API
```

#### func (*CmdVoidWriteF32) AddRead

```go
func (a *CmdVoidWriteF32) AddRead(rng memory.Range, id binary.ID) *CmdVoidWriteF32
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidWriteF32 pointer is returned so that calls can be chained.

#### func (*CmdVoidWriteF32) AddWrite

```go
func (a *CmdVoidWriteF32) AddWrite(rng memory.Range, id binary.ID) *CmdVoidWriteF32
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidWriteF32 pointer is returned so that calls can be chained.

#### func (*CmdVoidWriteF32) Class

```go
func (*CmdVoidWriteF32) Class() binary.Class
```

#### func (*CmdVoidWriteF32) Flags

```go
func (c *CmdVoidWriteF32) Flags() atom.Flags
```

#### func (*CmdVoidWriteF32) Mutate

```go
func (ϟa *CmdVoidWriteF32) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidWriteF32) Observations

```go
func (a *CmdVoidWriteF32) Observations() *atom.Observations
```

#### func (*CmdVoidWriteF32) Replay

```go
func (ϟa *CmdVoidWriteF32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidWriteF32) String

```go
func (a *CmdVoidWriteF32) String() string
```

#### func (*CmdVoidWriteF32) TypeID

```go
func (c *CmdVoidWriteF32) TypeID() atom.TypeID
```

#### type CmdVoidWriteF64

```go
type CmdVoidWriteF64 struct {
	binary.Generate `display:"cmd_void_write_f64"`

	A F64ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidWriteF64
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidWriteF64

```go
func NewCmdVoidWriteF64(A memory.Pointer) *CmdVoidWriteF64
```

#### func (*CmdVoidWriteF64) API

```go
func (c *CmdVoidWriteF64) API() gfxapi.API
```

#### func (*CmdVoidWriteF64) AddRead

```go
func (a *CmdVoidWriteF64) AddRead(rng memory.Range, id binary.ID) *CmdVoidWriteF64
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidWriteF64 pointer is returned so that calls can be chained.

#### func (*CmdVoidWriteF64) AddWrite

```go
func (a *CmdVoidWriteF64) AddWrite(rng memory.Range, id binary.ID) *CmdVoidWriteF64
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidWriteF64 pointer is returned so that calls can be chained.

#### func (*CmdVoidWriteF64) Class

```go
func (*CmdVoidWriteF64) Class() binary.Class
```

#### func (*CmdVoidWriteF64) Flags

```go
func (c *CmdVoidWriteF64) Flags() atom.Flags
```

#### func (*CmdVoidWriteF64) Mutate

```go
func (ϟa *CmdVoidWriteF64) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidWriteF64) Observations

```go
func (a *CmdVoidWriteF64) Observations() *atom.Observations
```

#### func (*CmdVoidWriteF64) Replay

```go
func (ϟa *CmdVoidWriteF64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidWriteF64) String

```go
func (a *CmdVoidWriteF64) String() string
```

#### func (*CmdVoidWriteF64) TypeID

```go
func (c *CmdVoidWriteF64) TypeID() atom.TypeID
```

#### type CmdVoidWritePtrs

```go
type CmdVoidWritePtrs struct {
	binary.Generate `display:"cmd_void_write_ptrs"`

	A F32ᵖ
	B U16ᵖ
	C Boolᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidWritePtrs
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidWritePtrs

```go
func NewCmdVoidWritePtrs(A memory.Pointer, B memory.Pointer, C memory.Pointer) *CmdVoidWritePtrs
```

#### func (*CmdVoidWritePtrs) API

```go
func (c *CmdVoidWritePtrs) API() gfxapi.API
```

#### func (*CmdVoidWritePtrs) AddRead

```go
func (a *CmdVoidWritePtrs) AddRead(rng memory.Range, id binary.ID) *CmdVoidWritePtrs
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidWritePtrs pointer is returned so that calls can be chained.

#### func (*CmdVoidWritePtrs) AddWrite

```go
func (a *CmdVoidWritePtrs) AddWrite(rng memory.Range, id binary.ID) *CmdVoidWritePtrs
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidWritePtrs pointer is returned so that calls can be chained.

#### func (*CmdVoidWritePtrs) Class

```go
func (*CmdVoidWritePtrs) Class() binary.Class
```

#### func (*CmdVoidWritePtrs) Flags

```go
func (c *CmdVoidWritePtrs) Flags() atom.Flags
```

#### func (*CmdVoidWritePtrs) Mutate

```go
func (ϟa *CmdVoidWritePtrs) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidWritePtrs) Observations

```go
func (a *CmdVoidWritePtrs) Observations() *atom.Observations
```

#### func (*CmdVoidWritePtrs) Replay

```go
func (ϟa *CmdVoidWritePtrs) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidWritePtrs) String

```go
func (a *CmdVoidWritePtrs) String() string
```

#### func (*CmdVoidWritePtrs) TypeID

```go
func (c *CmdVoidWritePtrs) TypeID() atom.TypeID
```

#### type CmdVoidWriteS16

```go
type CmdVoidWriteS16 struct {
	binary.Generate `display:"cmd_void_write_s16"`

	A S16ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidWriteS16
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidWriteS16

```go
func NewCmdVoidWriteS16(A memory.Pointer) *CmdVoidWriteS16
```

#### func (*CmdVoidWriteS16) API

```go
func (c *CmdVoidWriteS16) API() gfxapi.API
```

#### func (*CmdVoidWriteS16) AddRead

```go
func (a *CmdVoidWriteS16) AddRead(rng memory.Range, id binary.ID) *CmdVoidWriteS16
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidWriteS16 pointer is returned so that calls can be chained.

#### func (*CmdVoidWriteS16) AddWrite

```go
func (a *CmdVoidWriteS16) AddWrite(rng memory.Range, id binary.ID) *CmdVoidWriteS16
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidWriteS16 pointer is returned so that calls can be chained.

#### func (*CmdVoidWriteS16) Class

```go
func (*CmdVoidWriteS16) Class() binary.Class
```

#### func (*CmdVoidWriteS16) Flags

```go
func (c *CmdVoidWriteS16) Flags() atom.Flags
```

#### func (*CmdVoidWriteS16) Mutate

```go
func (ϟa *CmdVoidWriteS16) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidWriteS16) Observations

```go
func (a *CmdVoidWriteS16) Observations() *atom.Observations
```

#### func (*CmdVoidWriteS16) Replay

```go
func (ϟa *CmdVoidWriteS16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidWriteS16) String

```go
func (a *CmdVoidWriteS16) String() string
```

#### func (*CmdVoidWriteS16) TypeID

```go
func (c *CmdVoidWriteS16) TypeID() atom.TypeID
```

#### type CmdVoidWriteS32

```go
type CmdVoidWriteS32 struct {
	binary.Generate `display:"cmd_void_write_s32"`

	A S32ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidWriteS32
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidWriteS32

```go
func NewCmdVoidWriteS32(A memory.Pointer) *CmdVoidWriteS32
```

#### func (*CmdVoidWriteS32) API

```go
func (c *CmdVoidWriteS32) API() gfxapi.API
```

#### func (*CmdVoidWriteS32) AddRead

```go
func (a *CmdVoidWriteS32) AddRead(rng memory.Range, id binary.ID) *CmdVoidWriteS32
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidWriteS32 pointer is returned so that calls can be chained.

#### func (*CmdVoidWriteS32) AddWrite

```go
func (a *CmdVoidWriteS32) AddWrite(rng memory.Range, id binary.ID) *CmdVoidWriteS32
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidWriteS32 pointer is returned so that calls can be chained.

#### func (*CmdVoidWriteS32) Class

```go
func (*CmdVoidWriteS32) Class() binary.Class
```

#### func (*CmdVoidWriteS32) Flags

```go
func (c *CmdVoidWriteS32) Flags() atom.Flags
```

#### func (*CmdVoidWriteS32) Mutate

```go
func (ϟa *CmdVoidWriteS32) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidWriteS32) Observations

```go
func (a *CmdVoidWriteS32) Observations() *atom.Observations
```

#### func (*CmdVoidWriteS32) Replay

```go
func (ϟa *CmdVoidWriteS32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidWriteS32) String

```go
func (a *CmdVoidWriteS32) String() string
```

#### func (*CmdVoidWriteS32) TypeID

```go
func (c *CmdVoidWriteS32) TypeID() atom.TypeID
```

#### type CmdVoidWriteS64

```go
type CmdVoidWriteS64 struct {
	binary.Generate `display:"cmd_void_write_s64"`

	A S64ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidWriteS64
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidWriteS64

```go
func NewCmdVoidWriteS64(A memory.Pointer) *CmdVoidWriteS64
```

#### func (*CmdVoidWriteS64) API

```go
func (c *CmdVoidWriteS64) API() gfxapi.API
```

#### func (*CmdVoidWriteS64) AddRead

```go
func (a *CmdVoidWriteS64) AddRead(rng memory.Range, id binary.ID) *CmdVoidWriteS64
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidWriteS64 pointer is returned so that calls can be chained.

#### func (*CmdVoidWriteS64) AddWrite

```go
func (a *CmdVoidWriteS64) AddWrite(rng memory.Range, id binary.ID) *CmdVoidWriteS64
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidWriteS64 pointer is returned so that calls can be chained.

#### func (*CmdVoidWriteS64) Class

```go
func (*CmdVoidWriteS64) Class() binary.Class
```

#### func (*CmdVoidWriteS64) Flags

```go
func (c *CmdVoidWriteS64) Flags() atom.Flags
```

#### func (*CmdVoidWriteS64) Mutate

```go
func (ϟa *CmdVoidWriteS64) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidWriteS64) Observations

```go
func (a *CmdVoidWriteS64) Observations() *atom.Observations
```

#### func (*CmdVoidWriteS64) Replay

```go
func (ϟa *CmdVoidWriteS64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidWriteS64) String

```go
func (a *CmdVoidWriteS64) String() string
```

#### func (*CmdVoidWriteS64) TypeID

```go
func (c *CmdVoidWriteS64) TypeID() atom.TypeID
```

#### type CmdVoidWriteS8

```go
type CmdVoidWriteS8 struct {
	binary.Generate `display:"cmd_void_write_s8"`

	A S8ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidWriteS8
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidWriteS8

```go
func NewCmdVoidWriteS8(A memory.Pointer) *CmdVoidWriteS8
```

#### func (*CmdVoidWriteS8) API

```go
func (c *CmdVoidWriteS8) API() gfxapi.API
```

#### func (*CmdVoidWriteS8) AddRead

```go
func (a *CmdVoidWriteS8) AddRead(rng memory.Range, id binary.ID) *CmdVoidWriteS8
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidWriteS8 pointer is returned so that calls can be chained.

#### func (*CmdVoidWriteS8) AddWrite

```go
func (a *CmdVoidWriteS8) AddWrite(rng memory.Range, id binary.ID) *CmdVoidWriteS8
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidWriteS8 pointer is returned so that calls can be chained.

#### func (*CmdVoidWriteS8) Class

```go
func (*CmdVoidWriteS8) Class() binary.Class
```

#### func (*CmdVoidWriteS8) Flags

```go
func (c *CmdVoidWriteS8) Flags() atom.Flags
```

#### func (*CmdVoidWriteS8) Mutate

```go
func (ϟa *CmdVoidWriteS8) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidWriteS8) Observations

```go
func (a *CmdVoidWriteS8) Observations() *atom.Observations
```

#### func (*CmdVoidWriteS8) Replay

```go
func (ϟa *CmdVoidWriteS8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidWriteS8) String

```go
func (a *CmdVoidWriteS8) String() string
```

#### func (*CmdVoidWriteS8) TypeID

```go
func (c *CmdVoidWriteS8) TypeID() atom.TypeID
```

#### type CmdVoidWriteU16

```go
type CmdVoidWriteU16 struct {
	binary.Generate `display:"cmd_void_write_u16"`

	A U16ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidWriteU16
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidWriteU16

```go
func NewCmdVoidWriteU16(A memory.Pointer) *CmdVoidWriteU16
```

#### func (*CmdVoidWriteU16) API

```go
func (c *CmdVoidWriteU16) API() gfxapi.API
```

#### func (*CmdVoidWriteU16) AddRead

```go
func (a *CmdVoidWriteU16) AddRead(rng memory.Range, id binary.ID) *CmdVoidWriteU16
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidWriteU16 pointer is returned so that calls can be chained.

#### func (*CmdVoidWriteU16) AddWrite

```go
func (a *CmdVoidWriteU16) AddWrite(rng memory.Range, id binary.ID) *CmdVoidWriteU16
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidWriteU16 pointer is returned so that calls can be chained.

#### func (*CmdVoidWriteU16) Class

```go
func (*CmdVoidWriteU16) Class() binary.Class
```

#### func (*CmdVoidWriteU16) Flags

```go
func (c *CmdVoidWriteU16) Flags() atom.Flags
```

#### func (*CmdVoidWriteU16) Mutate

```go
func (ϟa *CmdVoidWriteU16) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidWriteU16) Observations

```go
func (a *CmdVoidWriteU16) Observations() *atom.Observations
```

#### func (*CmdVoidWriteU16) Replay

```go
func (ϟa *CmdVoidWriteU16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidWriteU16) String

```go
func (a *CmdVoidWriteU16) String() string
```

#### func (*CmdVoidWriteU16) TypeID

```go
func (c *CmdVoidWriteU16) TypeID() atom.TypeID
```

#### type CmdVoidWriteU32

```go
type CmdVoidWriteU32 struct {
	binary.Generate `display:"cmd_void_write_u32"`

	A U32ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidWriteU32
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidWriteU32

```go
func NewCmdVoidWriteU32(A memory.Pointer) *CmdVoidWriteU32
```

#### func (*CmdVoidWriteU32) API

```go
func (c *CmdVoidWriteU32) API() gfxapi.API
```

#### func (*CmdVoidWriteU32) AddRead

```go
func (a *CmdVoidWriteU32) AddRead(rng memory.Range, id binary.ID) *CmdVoidWriteU32
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidWriteU32 pointer is returned so that calls can be chained.

#### func (*CmdVoidWriteU32) AddWrite

```go
func (a *CmdVoidWriteU32) AddWrite(rng memory.Range, id binary.ID) *CmdVoidWriteU32
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidWriteU32 pointer is returned so that calls can be chained.

#### func (*CmdVoidWriteU32) Class

```go
func (*CmdVoidWriteU32) Class() binary.Class
```

#### func (*CmdVoidWriteU32) Flags

```go
func (c *CmdVoidWriteU32) Flags() atom.Flags
```

#### func (*CmdVoidWriteU32) Mutate

```go
func (ϟa *CmdVoidWriteU32) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidWriteU32) Observations

```go
func (a *CmdVoidWriteU32) Observations() *atom.Observations
```

#### func (*CmdVoidWriteU32) Replay

```go
func (ϟa *CmdVoidWriteU32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidWriteU32) String

```go
func (a *CmdVoidWriteU32) String() string
```

#### func (*CmdVoidWriteU32) TypeID

```go
func (c *CmdVoidWriteU32) TypeID() atom.TypeID
```

#### type CmdVoidWriteU64

```go
type CmdVoidWriteU64 struct {
	binary.Generate `display:"cmd_void_write_u64"`

	A U64ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidWriteU64
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidWriteU64

```go
func NewCmdVoidWriteU64(A memory.Pointer) *CmdVoidWriteU64
```

#### func (*CmdVoidWriteU64) API

```go
func (c *CmdVoidWriteU64) API() gfxapi.API
```

#### func (*CmdVoidWriteU64) AddRead

```go
func (a *CmdVoidWriteU64) AddRead(rng memory.Range, id binary.ID) *CmdVoidWriteU64
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidWriteU64 pointer is returned so that calls can be chained.

#### func (*CmdVoidWriteU64) AddWrite

```go
func (a *CmdVoidWriteU64) AddWrite(rng memory.Range, id binary.ID) *CmdVoidWriteU64
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidWriteU64 pointer is returned so that calls can be chained.

#### func (*CmdVoidWriteU64) Class

```go
func (*CmdVoidWriteU64) Class() binary.Class
```

#### func (*CmdVoidWriteU64) Flags

```go
func (c *CmdVoidWriteU64) Flags() atom.Flags
```

#### func (*CmdVoidWriteU64) Mutate

```go
func (ϟa *CmdVoidWriteU64) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidWriteU64) Observations

```go
func (a *CmdVoidWriteU64) Observations() *atom.Observations
```

#### func (*CmdVoidWriteU64) Replay

```go
func (ϟa *CmdVoidWriteU64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidWriteU64) String

```go
func (a *CmdVoidWriteU64) String() string
```

#### func (*CmdVoidWriteU64) TypeID

```go
func (c *CmdVoidWriteU64) TypeID() atom.TypeID
```

#### type CmdVoidWriteU8

```go
type CmdVoidWriteU8 struct {
	binary.Generate `display:"cmd_void_write_u8"`

	A U8ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidWriteU8
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidWriteU8

```go
func NewCmdVoidWriteU8(A memory.Pointer) *CmdVoidWriteU8
```

#### func (*CmdVoidWriteU8) API

```go
func (c *CmdVoidWriteU8) API() gfxapi.API
```

#### func (*CmdVoidWriteU8) AddRead

```go
func (a *CmdVoidWriteU8) AddRead(rng memory.Range, id binary.ID) *CmdVoidWriteU8
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CmdVoidWriteU8 pointer is returned so that calls can be chained.

#### func (*CmdVoidWriteU8) AddWrite

```go
func (a *CmdVoidWriteU8) AddWrite(rng memory.Range, id binary.ID) *CmdVoidWriteU8
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CmdVoidWriteU8 pointer is returned so that calls can be chained.

#### func (*CmdVoidWriteU8) Class

```go
func (*CmdVoidWriteU8) Class() binary.Class
```

#### func (*CmdVoidWriteU8) Flags

```go
func (c *CmdVoidWriteU8) Flags() atom.Flags
```

#### func (*CmdVoidWriteU8) Mutate

```go
func (ϟa *CmdVoidWriteU8) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CmdVoidWriteU8) Observations

```go
func (a *CmdVoidWriteU8) Observations() *atom.Observations
```

#### func (*CmdVoidWriteU8) Replay

```go
func (ϟa *CmdVoidWriteU8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*CmdVoidWriteU8) String

```go
func (a *CmdVoidWriteU8) String() string
```

#### func (*CmdVoidWriteU8) TypeID

```go
func (c *CmdVoidWriteU8) TypeID() atom.TypeID
```

#### type F32ˢ

```go
type F32ˢ struct {
	binary.Generate
	SliceInfo
}
```

F32ˢ is a slice of float32.

#### func  AsF32ˢ

```go
func AsF32ˢ(s Slice, ϟs *gfxapi.State) F32ˢ
```
AsF32ˢ returns s cast to a F32ˢ. The returned slice length will be calculated so
that the returned slice is no longer (in bytes) than s.

#### func  MakeF32ˢ

```go
func MakeF32ˢ(count uint64, ϟs *gfxapi.State) F32ˢ
```
MakeF32ˢ returns a F32ˢ backed by a new memory pool.

#### func (*F32ˢ) Class

```go
func (*F32ˢ) Class() binary.Class
```

#### func (F32ˢ) Clone

```go
func (s F32ˢ) Clone(ϟs *gfxapi.State) F32ˢ
```
Clone returns a copy of the F32ˢ in a new memory pool.

#### func (F32ˢ) Copy

```go
func (dst F32ˢ) Copy(src F32ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s F32ˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (F32ˢ) Decoder

```go
func (s F32ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (F32ˢ) ElementSize

```go
func (s F32ˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that F32ˢ points to.

#### func (F32ˢ) Encoder

```go
func (s F32ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (F32ˢ) Index

```go
func (s F32ˢ) Index(i uint64, ϟs *gfxapi.State) F32ᵖ
```
Index returns a F32ᵖ to the i'th element in this F32ˢ.

#### func (F32ˢ) OnRead

```go
func (s F32ˢ) OnRead(ϟs *gfxapi.State) F32ˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (F32ˢ) OnWrite

```go
func (s F32ˢ) OnWrite(ϟs *gfxapi.State) F32ˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (F32ˢ) Range

```go
func (s F32ˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (F32ˢ) Read

```go
func (s F32ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []float32
```
Read reads and returns all the float32 elements in this F32ˢ.

#### func (F32ˢ) ResourceID

```go
func (s F32ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (F32ˢ) Slice

```go
func (s F32ˢ) Slice(start, end uint64, ϟs *gfxapi.State) F32ˢ
```
Slice returns a sub-slice from the F32ˢ using start and end indices.

#### func (F32ˢ) String

```go
func (s F32ˢ) String() string
```
String returns a string description of the F32ˢ slice.

#### func (F32ˢ) Write

```go
func (s F32ˢ) Write(src []float32, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type F32ᵖ

```go
type F32ᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}
```

F32ᵖ is a pointer to a float32 element.

#### func  NewF32ᵖ

```go
func NewF32ᵖ(addr memory.Pointer) F32ᵖ
```
NewF32ᵖ returns a F32ᵖ that points to addr in the application pool.

#### func (*F32ᵖ) Class

```go
func (*F32ᵖ) Class() binary.Class
```

#### func (F32ᵖ) ElementSize

```go
func (p F32ᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that F32ᵖ points to.

#### func (F32ᵖ) OnRead

```go
func (p F32ᵖ) OnRead(ϟs *gfxapi.State) F32ᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (F32ᵖ) OnWrite

```go
func (p F32ᵖ) OnWrite(ϟs *gfxapi.State) F32ᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (F32ᵖ) Read

```go
func (p F32ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) float32
```
Read reads and returns the float32 element at the pointer.

#### func (F32ᵖ) Slice

```go
func (p F32ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) F32ˢ
```
Slice returns a new F32ˢ from the pointer using start and end indices.

#### func (F32ᵖ) String

```go
func (p F32ᵖ) String() string
```
String returns a string description of the F32ᵖ pointer.

#### func (F32ᵖ) Write

```go
func (p F32ᵖ) Write(value float32, ϟs *gfxapi.State)
```
Write writes value to the float32 element at the pointer.

#### type F64ˢ

```go
type F64ˢ struct {
	binary.Generate
	SliceInfo
}
```

F64ˢ is a slice of float64.

#### func  AsF64ˢ

```go
func AsF64ˢ(s Slice, ϟs *gfxapi.State) F64ˢ
```
AsF64ˢ returns s cast to a F64ˢ. The returned slice length will be calculated so
that the returned slice is no longer (in bytes) than s.

#### func  MakeF64ˢ

```go
func MakeF64ˢ(count uint64, ϟs *gfxapi.State) F64ˢ
```
MakeF64ˢ returns a F64ˢ backed by a new memory pool.

#### func (*F64ˢ) Class

```go
func (*F64ˢ) Class() binary.Class
```

#### func (F64ˢ) Clone

```go
func (s F64ˢ) Clone(ϟs *gfxapi.State) F64ˢ
```
Clone returns a copy of the F64ˢ in a new memory pool.

#### func (F64ˢ) Copy

```go
func (dst F64ˢ) Copy(src F64ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s F64ˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (F64ˢ) Decoder

```go
func (s F64ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (F64ˢ) ElementSize

```go
func (s F64ˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that F64ˢ points to.

#### func (F64ˢ) Encoder

```go
func (s F64ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (F64ˢ) Index

```go
func (s F64ˢ) Index(i uint64, ϟs *gfxapi.State) F64ᵖ
```
Index returns a F64ᵖ to the i'th element in this F64ˢ.

#### func (F64ˢ) OnRead

```go
func (s F64ˢ) OnRead(ϟs *gfxapi.State) F64ˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (F64ˢ) OnWrite

```go
func (s F64ˢ) OnWrite(ϟs *gfxapi.State) F64ˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (F64ˢ) Range

```go
func (s F64ˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (F64ˢ) Read

```go
func (s F64ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []float64
```
Read reads and returns all the float64 elements in this F64ˢ.

#### func (F64ˢ) ResourceID

```go
func (s F64ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (F64ˢ) Slice

```go
func (s F64ˢ) Slice(start, end uint64, ϟs *gfxapi.State) F64ˢ
```
Slice returns a sub-slice from the F64ˢ using start and end indices.

#### func (F64ˢ) String

```go
func (s F64ˢ) String() string
```
String returns a string description of the F64ˢ slice.

#### func (F64ˢ) Write

```go
func (s F64ˢ) Write(src []float64, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type F64ᵖ

```go
type F64ᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}
```

F64ᵖ is a pointer to a float64 element.

#### func  NewF64ᵖ

```go
func NewF64ᵖ(addr memory.Pointer) F64ᵖ
```
NewF64ᵖ returns a F64ᵖ that points to addr in the application pool.

#### func (*F64ᵖ) Class

```go
func (*F64ᵖ) Class() binary.Class
```

#### func (F64ᵖ) ElementSize

```go
func (p F64ᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that F64ᵖ points to.

#### func (F64ᵖ) OnRead

```go
func (p F64ᵖ) OnRead(ϟs *gfxapi.State) F64ᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (F64ᵖ) OnWrite

```go
func (p F64ᵖ) OnWrite(ϟs *gfxapi.State) F64ᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (F64ᵖ) Read

```go
func (p F64ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) float64
```
Read reads and returns the float64 element at the pointer.

#### func (F64ᵖ) Slice

```go
func (p F64ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) F64ˢ
```
Slice returns a new F64ˢ from the pointer using start and end indices.

#### func (F64ᵖ) String

```go
func (p F64ᵖ) String() string
```
String returns a string description of the F64ᵖ pointer.

#### func (F64ᵖ) Write

```go
func (p F64ᵖ) Write(value float64, ϟs *gfxapi.State)
```
Write writes value to the float64 element at the pointer.

#### type Globals

```go
type Globals struct {
	binary.Generate
	U8s  U8ˢ
	U16s U16ˢ
	U32s U32ˢ
	Ints Intˢ
	Str  string
}
```

//////////////////////////////////////////////////////////////////////////////
Globals
//////////////////////////////////////////////////////////////////////////////

#### func (*Globals) Class

```go
func (*Globals) Class() binary.Class
```

#### func (*Globals) Init

```go
func (g *Globals) Init()
```

#### type Imported

```go
type Imported struct {
	binary.Generate
	Value uint32
}
```


#### func (*Imported) Class

```go
func (*Imported) Class() binary.Class
```

#### func (Imported) Init

```go
func (Imported) Init()
```

#### type Included

```go
type Included struct {
	binary.Generate
	CreatedAt atom.ID
	S         string
}
```

//////////////////////////////////////////////////////////////////////////////
class Included
//////////////////////////////////////////////////////////////////////////////

#### func (*Included) Class

```go
func (*Included) Class() binary.Class
```

#### func (*Included) GetCreatedAt

```go
func (c *Included) GetCreatedAt() atom.ID
```

#### func (*Included) Init

```go
func (c *Included) Init()
```

#### type Intˢ

```go
type Intˢ struct {
	binary.Generate
	SliceInfo
}
```

Intˢ is a slice of int64.

#### func  AsIntˢ

```go
func AsIntˢ(s Slice, ϟs *gfxapi.State) Intˢ
```
AsIntˢ returns s cast to a Intˢ. The returned slice length will be calculated so
that the returned slice is no longer (in bytes) than s.

#### func  MakeIntˢ

```go
func MakeIntˢ(count uint64, ϟs *gfxapi.State) Intˢ
```
MakeIntˢ returns a Intˢ backed by a new memory pool.

#### func (*Intˢ) Class

```go
func (*Intˢ) Class() binary.Class
```

#### func (Intˢ) Clone

```go
func (s Intˢ) Clone(ϟs *gfxapi.State) Intˢ
```
Clone returns a copy of the Intˢ in a new memory pool.

#### func (Intˢ) Copy

```go
func (dst Intˢ) Copy(src Intˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Intˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (Intˢ) Decoder

```go
func (s Intˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (Intˢ) ElementSize

```go
func (s Intˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Intˢ points to.

#### func (Intˢ) Encoder

```go
func (s Intˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (Intˢ) Index

```go
func (s Intˢ) Index(i uint64, ϟs *gfxapi.State) Intᵖ
```
Index returns a Intᵖ to the i'th element in this Intˢ.

#### func (Intˢ) OnRead

```go
func (s Intˢ) OnRead(ϟs *gfxapi.State) Intˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (Intˢ) OnWrite

```go
func (s Intˢ) OnWrite(ϟs *gfxapi.State) Intˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (Intˢ) Range

```go
func (s Intˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (Intˢ) Read

```go
func (s Intˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []int64
```
Read reads and returns all the int64 elements in this Intˢ.

#### func (Intˢ) ResourceID

```go
func (s Intˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (Intˢ) Slice

```go
func (s Intˢ) Slice(start, end uint64, ϟs *gfxapi.State) Intˢ
```
Slice returns a sub-slice from the Intˢ using start and end indices.

#### func (Intˢ) String

```go
func (s Intˢ) String() string
```
String returns a string description of the Intˢ slice.

#### func (Intˢ) Write

```go
func (s Intˢ) Write(src []int64, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type Intᵖ

```go
type Intᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}
```

Intᵖ is a pointer to a int64 element.

#### func  NewIntᵖ

```go
func NewIntᵖ(addr memory.Pointer) Intᵖ
```
NewIntᵖ returns a Intᵖ that points to addr in the application pool.

#### func (*Intᵖ) Class

```go
func (*Intᵖ) Class() binary.Class
```

#### func (Intᵖ) ElementSize

```go
func (p Intᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Intᵖ points to.

#### func (Intᵖ) OnRead

```go
func (p Intᵖ) OnRead(ϟs *gfxapi.State) Intᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (Intᵖ) OnWrite

```go
func (p Intᵖ) OnWrite(ϟs *gfxapi.State) Intᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (Intᵖ) Read

```go
func (p Intᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) int64
```
Read reads and returns the int64 element at the pointer.

#### func (Intᵖ) Slice

```go
func (p Intᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Intˢ
```
Slice returns a new Intˢ from the pointer using start and end indices.

#### func (Intᵖ) String

```go
func (p Intᵖ) String() string
```
String returns a string description of the Intᵖ pointer.

#### func (Intᵖ) Write

```go
func (p Intᵖ) Write(value int64, ϟs *gfxapi.State)
```
Write writes value to the int64 element at the pointer.

#### type Remappedˢ

```go
type Remappedˢ struct {
	binary.Generate
	SliceInfo
}
```

Remappedˢ is a slice of remapped.

#### func  AsRemappedˢ

```go
func AsRemappedˢ(s Slice, ϟs *gfxapi.State) Remappedˢ
```
AsRemappedˢ returns s cast to a Remappedˢ. The returned slice length will be
calculated so that the returned slice is no longer (in bytes) than s.

#### func  MakeRemappedˢ

```go
func MakeRemappedˢ(count uint64, ϟs *gfxapi.State) Remappedˢ
```
MakeRemappedˢ returns a Remappedˢ backed by a new memory pool.

#### func (*Remappedˢ) Class

```go
func (*Remappedˢ) Class() binary.Class
```

#### func (Remappedˢ) Clone

```go
func (s Remappedˢ) Clone(ϟs *gfxapi.State) Remappedˢ
```
Clone returns a copy of the Remappedˢ in a new memory pool.

#### func (Remappedˢ) Copy

```go
func (dst Remappedˢ) Copy(src Remappedˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Remappedˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (Remappedˢ) Decoder

```go
func (s Remappedˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (Remappedˢ) ElementSize

```go
func (s Remappedˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Remappedˢ points to.

#### func (Remappedˢ) Encoder

```go
func (s Remappedˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (Remappedˢ) Index

```go
func (s Remappedˢ) Index(i uint64, ϟs *gfxapi.State) Remappedᵖ
```
Index returns a Remappedᵖ to the i'th element in this Remappedˢ.

#### func (Remappedˢ) OnRead

```go
func (s Remappedˢ) OnRead(ϟs *gfxapi.State) Remappedˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (Remappedˢ) OnWrite

```go
func (s Remappedˢ) OnWrite(ϟs *gfxapi.State) Remappedˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (Remappedˢ) Range

```go
func (s Remappedˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (Remappedˢ) Read

```go
func (s Remappedˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []remapped
```
Read reads and returns all the remapped elements in this Remappedˢ.

#### func (Remappedˢ) ResourceID

```go
func (s Remappedˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (Remappedˢ) Slice

```go
func (s Remappedˢ) Slice(start, end uint64, ϟs *gfxapi.State) Remappedˢ
```
Slice returns a sub-slice from the Remappedˢ using start and end indices.

#### func (Remappedˢ) String

```go
func (s Remappedˢ) String() string
```
String returns a string description of the Remappedˢ slice.

#### func (Remappedˢ) Write

```go
func (s Remappedˢ) Write(src []remapped, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type Remappedᵖ

```go
type Remappedᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}
```

Remappedᵖ is a pointer to a remapped element.

#### func  NewRemappedᵖ

```go
func NewRemappedᵖ(addr memory.Pointer) Remappedᵖ
```
NewRemappedᵖ returns a Remappedᵖ that points to addr in the application pool.

#### func (*Remappedᵖ) Class

```go
func (*Remappedᵖ) Class() binary.Class
```

#### func (Remappedᵖ) ElementSize

```go
func (p Remappedᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Remappedᵖ points to.

#### func (Remappedᵖ) OnRead

```go
func (p Remappedᵖ) OnRead(ϟs *gfxapi.State) Remappedᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (Remappedᵖ) OnWrite

```go
func (p Remappedᵖ) OnWrite(ϟs *gfxapi.State) Remappedᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (Remappedᵖ) Read

```go
func (p Remappedᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) remapped
```
Read reads and returns the remapped element at the pointer.

#### func (Remappedᵖ) Slice

```go
func (p Remappedᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Remappedˢ
```
Slice returns a new Remappedˢ from the pointer using start and end indices.

#### func (Remappedᵖ) String

```go
func (p Remappedᵖ) String() string
```
String returns a string description of the Remappedᵖ pointer.

#### func (Remappedᵖ) Write

```go
func (p Remappedᵖ) Write(value remapped, ϟs *gfxapi.State)
```
Write writes value to the remapped element at the pointer.

#### type S16ˢ

```go
type S16ˢ struct {
	binary.Generate
	SliceInfo
}
```

S16ˢ is a slice of int16.

#### func  AsS16ˢ

```go
func AsS16ˢ(s Slice, ϟs *gfxapi.State) S16ˢ
```
AsS16ˢ returns s cast to a S16ˢ. The returned slice length will be calculated so
that the returned slice is no longer (in bytes) than s.

#### func  MakeS16ˢ

```go
func MakeS16ˢ(count uint64, ϟs *gfxapi.State) S16ˢ
```
MakeS16ˢ returns a S16ˢ backed by a new memory pool.

#### func (*S16ˢ) Class

```go
func (*S16ˢ) Class() binary.Class
```

#### func (S16ˢ) Clone

```go
func (s S16ˢ) Clone(ϟs *gfxapi.State) S16ˢ
```
Clone returns a copy of the S16ˢ in a new memory pool.

#### func (S16ˢ) Copy

```go
func (dst S16ˢ) Copy(src S16ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s S16ˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (S16ˢ) Decoder

```go
func (s S16ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (S16ˢ) ElementSize

```go
func (s S16ˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that S16ˢ points to.

#### func (S16ˢ) Encoder

```go
func (s S16ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (S16ˢ) Index

```go
func (s S16ˢ) Index(i uint64, ϟs *gfxapi.State) S16ᵖ
```
Index returns a S16ᵖ to the i'th element in this S16ˢ.

#### func (S16ˢ) OnRead

```go
func (s S16ˢ) OnRead(ϟs *gfxapi.State) S16ˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (S16ˢ) OnWrite

```go
func (s S16ˢ) OnWrite(ϟs *gfxapi.State) S16ˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (S16ˢ) Range

```go
func (s S16ˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (S16ˢ) Read

```go
func (s S16ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []int16
```
Read reads and returns all the int16 elements in this S16ˢ.

#### func (S16ˢ) ResourceID

```go
func (s S16ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (S16ˢ) Slice

```go
func (s S16ˢ) Slice(start, end uint64, ϟs *gfxapi.State) S16ˢ
```
Slice returns a sub-slice from the S16ˢ using start and end indices.

#### func (S16ˢ) String

```go
func (s S16ˢ) String() string
```
String returns a string description of the S16ˢ slice.

#### func (S16ˢ) Write

```go
func (s S16ˢ) Write(src []int16, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type S16ᵖ

```go
type S16ᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}
```

S16ᵖ is a pointer to a int16 element.

#### func  NewS16ᵖ

```go
func NewS16ᵖ(addr memory.Pointer) S16ᵖ
```
NewS16ᵖ returns a S16ᵖ that points to addr in the application pool.

#### func (*S16ᵖ) Class

```go
func (*S16ᵖ) Class() binary.Class
```

#### func (S16ᵖ) ElementSize

```go
func (p S16ᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that S16ᵖ points to.

#### func (S16ᵖ) OnRead

```go
func (p S16ᵖ) OnRead(ϟs *gfxapi.State) S16ᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (S16ᵖ) OnWrite

```go
func (p S16ᵖ) OnWrite(ϟs *gfxapi.State) S16ᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (S16ᵖ) Read

```go
func (p S16ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) int16
```
Read reads and returns the int16 element at the pointer.

#### func (S16ᵖ) Slice

```go
func (p S16ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) S16ˢ
```
Slice returns a new S16ˢ from the pointer using start and end indices.

#### func (S16ᵖ) String

```go
func (p S16ᵖ) String() string
```
String returns a string description of the S16ᵖ pointer.

#### func (S16ᵖ) Write

```go
func (p S16ᵖ) Write(value int16, ϟs *gfxapi.State)
```
Write writes value to the int16 element at the pointer.

#### type S32ˢ

```go
type S32ˢ struct {
	binary.Generate
	SliceInfo
}
```

S32ˢ is a slice of int32.

#### func  AsS32ˢ

```go
func AsS32ˢ(s Slice, ϟs *gfxapi.State) S32ˢ
```
AsS32ˢ returns s cast to a S32ˢ. The returned slice length will be calculated so
that the returned slice is no longer (in bytes) than s.

#### func  MakeS32ˢ

```go
func MakeS32ˢ(count uint64, ϟs *gfxapi.State) S32ˢ
```
MakeS32ˢ returns a S32ˢ backed by a new memory pool.

#### func (*S32ˢ) Class

```go
func (*S32ˢ) Class() binary.Class
```

#### func (S32ˢ) Clone

```go
func (s S32ˢ) Clone(ϟs *gfxapi.State) S32ˢ
```
Clone returns a copy of the S32ˢ in a new memory pool.

#### func (S32ˢ) Copy

```go
func (dst S32ˢ) Copy(src S32ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s S32ˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (S32ˢ) Decoder

```go
func (s S32ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (S32ˢ) ElementSize

```go
func (s S32ˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that S32ˢ points to.

#### func (S32ˢ) Encoder

```go
func (s S32ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (S32ˢ) Index

```go
func (s S32ˢ) Index(i uint64, ϟs *gfxapi.State) S32ᵖ
```
Index returns a S32ᵖ to the i'th element in this S32ˢ.

#### func (S32ˢ) OnRead

```go
func (s S32ˢ) OnRead(ϟs *gfxapi.State) S32ˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (S32ˢ) OnWrite

```go
func (s S32ˢ) OnWrite(ϟs *gfxapi.State) S32ˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (S32ˢ) Range

```go
func (s S32ˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (S32ˢ) Read

```go
func (s S32ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []int32
```
Read reads and returns all the int32 elements in this S32ˢ.

#### func (S32ˢ) ResourceID

```go
func (s S32ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (S32ˢ) Slice

```go
func (s S32ˢ) Slice(start, end uint64, ϟs *gfxapi.State) S32ˢ
```
Slice returns a sub-slice from the S32ˢ using start and end indices.

#### func (S32ˢ) String

```go
func (s S32ˢ) String() string
```
String returns a string description of the S32ˢ slice.

#### func (S32ˢ) Write

```go
func (s S32ˢ) Write(src []int32, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type S32ᵖ

```go
type S32ᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}
```

S32ᵖ is a pointer to a int32 element.

#### func  NewS32ᵖ

```go
func NewS32ᵖ(addr memory.Pointer) S32ᵖ
```
NewS32ᵖ returns a S32ᵖ that points to addr in the application pool.

#### func (*S32ᵖ) Class

```go
func (*S32ᵖ) Class() binary.Class
```

#### func (S32ᵖ) ElementSize

```go
func (p S32ᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that S32ᵖ points to.

#### func (S32ᵖ) OnRead

```go
func (p S32ᵖ) OnRead(ϟs *gfxapi.State) S32ᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (S32ᵖ) OnWrite

```go
func (p S32ᵖ) OnWrite(ϟs *gfxapi.State) S32ᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (S32ᵖ) Read

```go
func (p S32ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) int32
```
Read reads and returns the int32 element at the pointer.

#### func (S32ᵖ) Slice

```go
func (p S32ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) S32ˢ
```
Slice returns a new S32ˢ from the pointer using start and end indices.

#### func (S32ᵖ) String

```go
func (p S32ᵖ) String() string
```
String returns a string description of the S32ᵖ pointer.

#### func (S32ᵖ) Write

```go
func (p S32ᵖ) Write(value int32, ϟs *gfxapi.State)
```
Write writes value to the int32 element at the pointer.

#### type S64ˢ

```go
type S64ˢ struct {
	binary.Generate
	SliceInfo
}
```

S64ˢ is a slice of int64.

#### func  AsS64ˢ

```go
func AsS64ˢ(s Slice, ϟs *gfxapi.State) S64ˢ
```
AsS64ˢ returns s cast to a S64ˢ. The returned slice length will be calculated so
that the returned slice is no longer (in bytes) than s.

#### func  MakeS64ˢ

```go
func MakeS64ˢ(count uint64, ϟs *gfxapi.State) S64ˢ
```
MakeS64ˢ returns a S64ˢ backed by a new memory pool.

#### func (*S64ˢ) Class

```go
func (*S64ˢ) Class() binary.Class
```

#### func (S64ˢ) Clone

```go
func (s S64ˢ) Clone(ϟs *gfxapi.State) S64ˢ
```
Clone returns a copy of the S64ˢ in a new memory pool.

#### func (S64ˢ) Copy

```go
func (dst S64ˢ) Copy(src S64ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s S64ˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (S64ˢ) Decoder

```go
func (s S64ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (S64ˢ) ElementSize

```go
func (s S64ˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that S64ˢ points to.

#### func (S64ˢ) Encoder

```go
func (s S64ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (S64ˢ) Index

```go
func (s S64ˢ) Index(i uint64, ϟs *gfxapi.State) S64ᵖ
```
Index returns a S64ᵖ to the i'th element in this S64ˢ.

#### func (S64ˢ) OnRead

```go
func (s S64ˢ) OnRead(ϟs *gfxapi.State) S64ˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (S64ˢ) OnWrite

```go
func (s S64ˢ) OnWrite(ϟs *gfxapi.State) S64ˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (S64ˢ) Range

```go
func (s S64ˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (S64ˢ) Read

```go
func (s S64ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []int64
```
Read reads and returns all the int64 elements in this S64ˢ.

#### func (S64ˢ) ResourceID

```go
func (s S64ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (S64ˢ) Slice

```go
func (s S64ˢ) Slice(start, end uint64, ϟs *gfxapi.State) S64ˢ
```
Slice returns a sub-slice from the S64ˢ using start and end indices.

#### func (S64ˢ) String

```go
func (s S64ˢ) String() string
```
String returns a string description of the S64ˢ slice.

#### func (S64ˢ) Write

```go
func (s S64ˢ) Write(src []int64, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type S64ᵖ

```go
type S64ᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}
```

S64ᵖ is a pointer to a int64 element.

#### func  NewS64ᵖ

```go
func NewS64ᵖ(addr memory.Pointer) S64ᵖ
```
NewS64ᵖ returns a S64ᵖ that points to addr in the application pool.

#### func (*S64ᵖ) Class

```go
func (*S64ᵖ) Class() binary.Class
```

#### func (S64ᵖ) ElementSize

```go
func (p S64ᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that S64ᵖ points to.

#### func (S64ᵖ) OnRead

```go
func (p S64ᵖ) OnRead(ϟs *gfxapi.State) S64ᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (S64ᵖ) OnWrite

```go
func (p S64ᵖ) OnWrite(ϟs *gfxapi.State) S64ᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (S64ᵖ) Read

```go
func (p S64ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) int64
```
Read reads and returns the int64 element at the pointer.

#### func (S64ᵖ) Slice

```go
func (p S64ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) S64ˢ
```
Slice returns a new S64ˢ from the pointer using start and end indices.

#### func (S64ᵖ) String

```go
func (p S64ᵖ) String() string
```
String returns a string description of the S64ᵖ pointer.

#### func (S64ᵖ) Write

```go
func (p S64ᵖ) Write(value int64, ϟs *gfxapi.State)
```
Write writes value to the int64 element at the pointer.

#### type S8ˢ

```go
type S8ˢ struct {
	binary.Generate
	SliceInfo
}
```

S8ˢ is a slice of int8.

#### func  AsS8ˢ

```go
func AsS8ˢ(s Slice, ϟs *gfxapi.State) S8ˢ
```
AsS8ˢ returns s cast to a S8ˢ. The returned slice length will be calculated so
that the returned slice is no longer (in bytes) than s.

#### func  MakeS8ˢ

```go
func MakeS8ˢ(count uint64, ϟs *gfxapi.State) S8ˢ
```
MakeS8ˢ returns a S8ˢ backed by a new memory pool.

#### func (*S8ˢ) Class

```go
func (*S8ˢ) Class() binary.Class
```

#### func (S8ˢ) Clone

```go
func (s S8ˢ) Clone(ϟs *gfxapi.State) S8ˢ
```
Clone returns a copy of the S8ˢ in a new memory pool.

#### func (S8ˢ) Copy

```go
func (dst S8ˢ) Copy(src S8ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s S8ˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (S8ˢ) Decoder

```go
func (s S8ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (S8ˢ) ElementSize

```go
func (s S8ˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that S8ˢ points to.

#### func (S8ˢ) Encoder

```go
func (s S8ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (S8ˢ) Index

```go
func (s S8ˢ) Index(i uint64, ϟs *gfxapi.State) S8ᵖ
```
Index returns a S8ᵖ to the i'th element in this S8ˢ.

#### func (S8ˢ) OnRead

```go
func (s S8ˢ) OnRead(ϟs *gfxapi.State) S8ˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (S8ˢ) OnWrite

```go
func (s S8ˢ) OnWrite(ϟs *gfxapi.State) S8ˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (S8ˢ) Range

```go
func (s S8ˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (S8ˢ) Read

```go
func (s S8ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []int8
```
Read reads and returns all the int8 elements in this S8ˢ.

#### func (S8ˢ) ResourceID

```go
func (s S8ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (S8ˢ) Slice

```go
func (s S8ˢ) Slice(start, end uint64, ϟs *gfxapi.State) S8ˢ
```
Slice returns a sub-slice from the S8ˢ using start and end indices.

#### func (S8ˢ) String

```go
func (s S8ˢ) String() string
```
String returns a string description of the S8ˢ slice.

#### func (S8ˢ) Write

```go
func (s S8ˢ) Write(src []int8, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type S8ᵖ

```go
type S8ᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}
```

S8ᵖ is a pointer to a int8 element.

#### func  NewS8ᵖ

```go
func NewS8ᵖ(addr memory.Pointer) S8ᵖ
```
NewS8ᵖ returns a S8ᵖ that points to addr in the application pool.

#### func (*S8ᵖ) Class

```go
func (*S8ᵖ) Class() binary.Class
```

#### func (S8ᵖ) ElementSize

```go
func (p S8ᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that S8ᵖ points to.

#### func (S8ᵖ) OnRead

```go
func (p S8ᵖ) OnRead(ϟs *gfxapi.State) S8ᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (S8ᵖ) OnWrite

```go
func (p S8ᵖ) OnWrite(ϟs *gfxapi.State) S8ᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (S8ᵖ) Read

```go
func (p S8ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) int8
```
Read reads and returns the int8 element at the pointer.

#### func (S8ᵖ) Slice

```go
func (p S8ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) S8ˢ
```
Slice returns a new S8ˢ from the pointer using start and end indices.

#### func (S8ᵖ) String

```go
func (p S8ᵖ) String() string
```
String returns a string description of the S8ᵖ pointer.

#### func (S8ᵖ) Write

```go
func (p S8ᵖ) Write(value int8, ϟs *gfxapi.State)
```
Write writes value to the int8 element at the pointer.

#### type Slice

```go
type Slice interface {
	// Info returns the SliceInfo of this slice.
	Info() SliceInfo
	// ElementSize returns the size in bytes of a single element in the slice.
	ElementSize(ϟs *gfxapi.State) uint64
}
```

Slice is the interface implemented by all slice types

#### type SliceInfo

```go
type SliceInfo struct {
	binary.Generate
	Root  memory.Pointer // Original pointer this slice derives from.
	Base  memory.Pointer // Address of first element.
	Count uint64         // Number of elements in the slice.
	Pool  memory.PoolID  // Pool that holds the element data.
}
```

SliceInfo is the common data between all slice types.

#### func (*SliceInfo) Class

```go
func (*SliceInfo) Class() binary.Class
```

#### func (SliceInfo) Info

```go
func (s SliceInfo) Info() SliceInfo
```
Info returns the SliceInfo. It is used to conform to the Slice interface.

#### type State

```go
type State struct {
	Globals
	ValidateOutput bool
}
```


#### type Tester

```go
type Tester struct {
	binary.Generate
	CreatedAt atom.ID
	A         Imported
	B         Included
}
```

//////////////////////////////////////////////////////////////////////////////
class Tester
//////////////////////////////////////////////////////////////////////////////

#### func (*Tester) Class

```go
func (*Tester) Class() binary.Class
```

#### func (*Tester) GetCreatedAt

```go
func (c *Tester) GetCreatedAt() atom.ID
```

#### func (*Tester) Init

```go
func (c *Tester) Init()
```

#### type U16ˢ

```go
type U16ˢ struct {
	binary.Generate
	SliceInfo
}
```

U16ˢ is a slice of uint16.

#### func  AsU16ˢ

```go
func AsU16ˢ(s Slice, ϟs *gfxapi.State) U16ˢ
```
AsU16ˢ returns s cast to a U16ˢ. The returned slice length will be calculated so
that the returned slice is no longer (in bytes) than s.

#### func  MakeU16ˢ

```go
func MakeU16ˢ(count uint64, ϟs *gfxapi.State) U16ˢ
```
MakeU16ˢ returns a U16ˢ backed by a new memory pool.

#### func (*U16ˢ) Class

```go
func (*U16ˢ) Class() binary.Class
```

#### func (U16ˢ) Clone

```go
func (s U16ˢ) Clone(ϟs *gfxapi.State) U16ˢ
```
Clone returns a copy of the U16ˢ in a new memory pool.

#### func (U16ˢ) Copy

```go
func (dst U16ˢ) Copy(src U16ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s U16ˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (U16ˢ) Decoder

```go
func (s U16ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (U16ˢ) ElementSize

```go
func (s U16ˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that U16ˢ points to.

#### func (U16ˢ) Encoder

```go
func (s U16ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (U16ˢ) Index

```go
func (s U16ˢ) Index(i uint64, ϟs *gfxapi.State) U16ᵖ
```
Index returns a U16ᵖ to the i'th element in this U16ˢ.

#### func (U16ˢ) OnRead

```go
func (s U16ˢ) OnRead(ϟs *gfxapi.State) U16ˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (U16ˢ) OnWrite

```go
func (s U16ˢ) OnWrite(ϟs *gfxapi.State) U16ˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (U16ˢ) Range

```go
func (s U16ˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (U16ˢ) Read

```go
func (s U16ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []uint16
```
Read reads and returns all the uint16 elements in this U16ˢ.

#### func (U16ˢ) ResourceID

```go
func (s U16ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (U16ˢ) Slice

```go
func (s U16ˢ) Slice(start, end uint64, ϟs *gfxapi.State) U16ˢ
```
Slice returns a sub-slice from the U16ˢ using start and end indices.

#### func (U16ˢ) String

```go
func (s U16ˢ) String() string
```
String returns a string description of the U16ˢ slice.

#### func (U16ˢ) Write

```go
func (s U16ˢ) Write(src []uint16, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type U16ᵖ

```go
type U16ᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}
```

U16ᵖ is a pointer to a uint16 element.

#### func  NewU16ᵖ

```go
func NewU16ᵖ(addr memory.Pointer) U16ᵖ
```
NewU16ᵖ returns a U16ᵖ that points to addr in the application pool.

#### func (*U16ᵖ) Class

```go
func (*U16ᵖ) Class() binary.Class
```

#### func (U16ᵖ) ElementSize

```go
func (p U16ᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that U16ᵖ points to.

#### func (U16ᵖ) OnRead

```go
func (p U16ᵖ) OnRead(ϟs *gfxapi.State) U16ᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (U16ᵖ) OnWrite

```go
func (p U16ᵖ) OnWrite(ϟs *gfxapi.State) U16ᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (U16ᵖ) Read

```go
func (p U16ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) uint16
```
Read reads and returns the uint16 element at the pointer.

#### func (U16ᵖ) Slice

```go
func (p U16ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) U16ˢ
```
Slice returns a new U16ˢ from the pointer using start and end indices.

#### func (U16ᵖ) String

```go
func (p U16ᵖ) String() string
```
String returns a string description of the U16ᵖ pointer.

#### func (U16ᵖ) Write

```go
func (p U16ᵖ) Write(value uint16, ϟs *gfxapi.State)
```
Write writes value to the uint16 element at the pointer.

#### type U32ˢ

```go
type U32ˢ struct {
	binary.Generate
	SliceInfo
}
```

U32ˢ is a slice of uint32.

#### func  AsU32ˢ

```go
func AsU32ˢ(s Slice, ϟs *gfxapi.State) U32ˢ
```
AsU32ˢ returns s cast to a U32ˢ. The returned slice length will be calculated so
that the returned slice is no longer (in bytes) than s.

#### func  MakeU32ˢ

```go
func MakeU32ˢ(count uint64, ϟs *gfxapi.State) U32ˢ
```
MakeU32ˢ returns a U32ˢ backed by a new memory pool.

#### func (*U32ˢ) Class

```go
func (*U32ˢ) Class() binary.Class
```

#### func (U32ˢ) Clone

```go
func (s U32ˢ) Clone(ϟs *gfxapi.State) U32ˢ
```
Clone returns a copy of the U32ˢ in a new memory pool.

#### func (U32ˢ) Copy

```go
func (dst U32ˢ) Copy(src U32ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s U32ˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (U32ˢ) Decoder

```go
func (s U32ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (U32ˢ) ElementSize

```go
func (s U32ˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that U32ˢ points to.

#### func (U32ˢ) Encoder

```go
func (s U32ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (U32ˢ) Index

```go
func (s U32ˢ) Index(i uint64, ϟs *gfxapi.State) U32ᵖ
```
Index returns a U32ᵖ to the i'th element in this U32ˢ.

#### func (U32ˢ) OnRead

```go
func (s U32ˢ) OnRead(ϟs *gfxapi.State) U32ˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (U32ˢ) OnWrite

```go
func (s U32ˢ) OnWrite(ϟs *gfxapi.State) U32ˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (U32ˢ) Range

```go
func (s U32ˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (U32ˢ) Read

```go
func (s U32ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []uint32
```
Read reads and returns all the uint32 elements in this U32ˢ.

#### func (U32ˢ) ResourceID

```go
func (s U32ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (U32ˢ) Slice

```go
func (s U32ˢ) Slice(start, end uint64, ϟs *gfxapi.State) U32ˢ
```
Slice returns a sub-slice from the U32ˢ using start and end indices.

#### func (U32ˢ) String

```go
func (s U32ˢ) String() string
```
String returns a string description of the U32ˢ slice.

#### func (U32ˢ) Write

```go
func (s U32ˢ) Write(src []uint32, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type U32ᵖ

```go
type U32ᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}
```

U32ᵖ is a pointer to a uint32 element.

#### func  NewU32ᵖ

```go
func NewU32ᵖ(addr memory.Pointer) U32ᵖ
```
NewU32ᵖ returns a U32ᵖ that points to addr in the application pool.

#### func (*U32ᵖ) Class

```go
func (*U32ᵖ) Class() binary.Class
```

#### func (U32ᵖ) ElementSize

```go
func (p U32ᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that U32ᵖ points to.

#### func (U32ᵖ) OnRead

```go
func (p U32ᵖ) OnRead(ϟs *gfxapi.State) U32ᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (U32ᵖ) OnWrite

```go
func (p U32ᵖ) OnWrite(ϟs *gfxapi.State) U32ᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (U32ᵖ) Read

```go
func (p U32ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) uint32
```
Read reads and returns the uint32 element at the pointer.

#### func (U32ᵖ) Slice

```go
func (p U32ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) U32ˢ
```
Slice returns a new U32ˢ from the pointer using start and end indices.

#### func (U32ᵖ) String

```go
func (p U32ᵖ) String() string
```
String returns a string description of the U32ᵖ pointer.

#### func (U32ᵖ) Write

```go
func (p U32ᵖ) Write(value uint32, ϟs *gfxapi.State)
```
Write writes value to the uint32 element at the pointer.

#### type U64ˢ

```go
type U64ˢ struct {
	binary.Generate
	SliceInfo
}
```

U64ˢ is a slice of uint64.

#### func  AsU64ˢ

```go
func AsU64ˢ(s Slice, ϟs *gfxapi.State) U64ˢ
```
AsU64ˢ returns s cast to a U64ˢ. The returned slice length will be calculated so
that the returned slice is no longer (in bytes) than s.

#### func  MakeU64ˢ

```go
func MakeU64ˢ(count uint64, ϟs *gfxapi.State) U64ˢ
```
MakeU64ˢ returns a U64ˢ backed by a new memory pool.

#### func (*U64ˢ) Class

```go
func (*U64ˢ) Class() binary.Class
```

#### func (U64ˢ) Clone

```go
func (s U64ˢ) Clone(ϟs *gfxapi.State) U64ˢ
```
Clone returns a copy of the U64ˢ in a new memory pool.

#### func (U64ˢ) Copy

```go
func (dst U64ˢ) Copy(src U64ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s U64ˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (U64ˢ) Decoder

```go
func (s U64ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (U64ˢ) ElementSize

```go
func (s U64ˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that U64ˢ points to.

#### func (U64ˢ) Encoder

```go
func (s U64ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (U64ˢ) Index

```go
func (s U64ˢ) Index(i uint64, ϟs *gfxapi.State) U64ᵖ
```
Index returns a U64ᵖ to the i'th element in this U64ˢ.

#### func (U64ˢ) OnRead

```go
func (s U64ˢ) OnRead(ϟs *gfxapi.State) U64ˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (U64ˢ) OnWrite

```go
func (s U64ˢ) OnWrite(ϟs *gfxapi.State) U64ˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (U64ˢ) Range

```go
func (s U64ˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (U64ˢ) Read

```go
func (s U64ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []uint64
```
Read reads and returns all the uint64 elements in this U64ˢ.

#### func (U64ˢ) ResourceID

```go
func (s U64ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (U64ˢ) Slice

```go
func (s U64ˢ) Slice(start, end uint64, ϟs *gfxapi.State) U64ˢ
```
Slice returns a sub-slice from the U64ˢ using start and end indices.

#### func (U64ˢ) String

```go
func (s U64ˢ) String() string
```
String returns a string description of the U64ˢ slice.

#### func (U64ˢ) Write

```go
func (s U64ˢ) Write(src []uint64, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type U64ᵖ

```go
type U64ᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}
```

U64ᵖ is a pointer to a uint64 element.

#### func  NewU64ᵖ

```go
func NewU64ᵖ(addr memory.Pointer) U64ᵖ
```
NewU64ᵖ returns a U64ᵖ that points to addr in the application pool.

#### func (*U64ᵖ) Class

```go
func (*U64ᵖ) Class() binary.Class
```

#### func (U64ᵖ) ElementSize

```go
func (p U64ᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that U64ᵖ points to.

#### func (U64ᵖ) OnRead

```go
func (p U64ᵖ) OnRead(ϟs *gfxapi.State) U64ᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (U64ᵖ) OnWrite

```go
func (p U64ᵖ) OnWrite(ϟs *gfxapi.State) U64ᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (U64ᵖ) Read

```go
func (p U64ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) uint64
```
Read reads and returns the uint64 element at the pointer.

#### func (U64ᵖ) Slice

```go
func (p U64ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) U64ˢ
```
Slice returns a new U64ˢ from the pointer using start and end indices.

#### func (U64ᵖ) String

```go
func (p U64ᵖ) String() string
```
String returns a string description of the U64ᵖ pointer.

#### func (U64ᵖ) Write

```go
func (p U64ᵖ) Write(value uint64, ϟs *gfxapi.State)
```
Write writes value to the uint64 element at the pointer.

#### type U8ˢ

```go
type U8ˢ struct {
	binary.Generate
	SliceInfo
}
```

U8ˢ is a slice of uint8.

#### func  AsU8ˢ

```go
func AsU8ˢ(s Slice, ϟs *gfxapi.State) U8ˢ
```
AsU8ˢ returns s cast to a U8ˢ. The returned slice length will be calculated so
that the returned slice is no longer (in bytes) than s.

#### func  MakeU8ˢ

```go
func MakeU8ˢ(count uint64, ϟs *gfxapi.State) U8ˢ
```
MakeU8ˢ returns a U8ˢ backed by a new memory pool.

#### func (*U8ˢ) Class

```go
func (*U8ˢ) Class() binary.Class
```

#### func (U8ˢ) Clone

```go
func (s U8ˢ) Clone(ϟs *gfxapi.State) U8ˢ
```
Clone returns a copy of the U8ˢ in a new memory pool.

#### func (U8ˢ) Copy

```go
func (dst U8ˢ) Copy(src U8ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s U8ˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (U8ˢ) Decoder

```go
func (s U8ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (U8ˢ) ElementSize

```go
func (s U8ˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that U8ˢ points to.

#### func (U8ˢ) Encoder

```go
func (s U8ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (U8ˢ) Index

```go
func (s U8ˢ) Index(i uint64, ϟs *gfxapi.State) U8ᵖ
```
Index returns a U8ᵖ to the i'th element in this U8ˢ.

#### func (U8ˢ) OnRead

```go
func (s U8ˢ) OnRead(ϟs *gfxapi.State) U8ˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (U8ˢ) OnWrite

```go
func (s U8ˢ) OnWrite(ϟs *gfxapi.State) U8ˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (U8ˢ) Range

```go
func (s U8ˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (U8ˢ) Read

```go
func (s U8ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []uint8
```
Read reads and returns all the uint8 elements in this U8ˢ.

#### func (U8ˢ) ResourceID

```go
func (s U8ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (U8ˢ) Slice

```go
func (s U8ˢ) Slice(start, end uint64, ϟs *gfxapi.State) U8ˢ
```
Slice returns a sub-slice from the U8ˢ using start and end indices.

#### func (U8ˢ) String

```go
func (s U8ˢ) String() string
```
String returns a string description of the U8ˢ slice.

#### func (U8ˢ) Write

```go
func (s U8ˢ) Write(src []uint8, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type U8ᵖ

```go
type U8ᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}
```

U8ᵖ is a pointer to a uint8 element.

#### func  NewU8ᵖ

```go
func NewU8ᵖ(addr memory.Pointer) U8ᵖ
```
NewU8ᵖ returns a U8ᵖ that points to addr in the application pool.

#### func (*U8ᵖ) Class

```go
func (*U8ᵖ) Class() binary.Class
```

#### func (U8ᵖ) ElementSize

```go
func (p U8ᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that U8ᵖ points to.

#### func (U8ᵖ) OnRead

```go
func (p U8ᵖ) OnRead(ϟs *gfxapi.State) U8ᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (U8ᵖ) OnWrite

```go
func (p U8ᵖ) OnWrite(ϟs *gfxapi.State) U8ᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (U8ᵖ) Read

```go
func (p U8ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) uint8
```
Read reads and returns the uint8 element at the pointer.

#### func (U8ᵖ) Slice

```go
func (p U8ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) U8ˢ
```
Slice returns a new U8ˢ from the pointer using start and end indices.

#### func (U8ᵖ) String

```go
func (p U8ᵖ) String() string
```
String returns a string description of the U8ᵖ pointer.

#### func (U8ᵖ) Write

```go
func (p U8ᵖ) Write(value uint8, ϟs *gfxapi.State)
```
Write writes value to the uint8 element at the pointer.

#### type Voidˢ

```go
type Voidˢ struct {
	binary.Generate
	SliceInfo
}
```

Voidˢ is a slice of void.

#### func  MakeVoidˢ

```go
func MakeVoidˢ(count uint64, ϟs *gfxapi.State) Voidˢ
```
MakeVoidˢ returns a Voidˢ backed by a new memory pool.

#### func (*Voidˢ) Class

```go
func (*Voidˢ) Class() binary.Class
```

#### func (Voidˢ) Clone

```go
func (s Voidˢ) Clone(ϟs *gfxapi.State) Voidˢ
```
Clone returns a copy of the Voidˢ in a new memory pool.

#### func (Voidˢ) Decoder

```go
func (s Voidˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (Voidˢ) ElementSize

```go
func (s Voidˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Voidˢ points to.

#### func (Voidˢ) Encoder

```go
func (s Voidˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (Voidˢ) Index

```go
func (s Voidˢ) Index(i uint64, ϟs *gfxapi.State) Voidᵖ
```
Index returns a Voidᵖ to the i'th element in this Voidˢ.

#### func (Voidˢ) OnRead

```go
func (s Voidˢ) OnRead(ϟs *gfxapi.State) Voidˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (Voidˢ) OnWrite

```go
func (s Voidˢ) OnWrite(ϟs *gfxapi.State) Voidˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (Voidˢ) Range

```go
func (s Voidˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (Voidˢ) ResourceID

```go
func (s Voidˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (Voidˢ) Slice

```go
func (s Voidˢ) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ
```
Slice returns a sub-slice from the Voidˢ using start and end indices.

#### func (Voidˢ) String

```go
func (s Voidˢ) String() string
```
String returns a string description of the Voidˢ slice.

#### type Voidᵖ

```go
type Voidᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}
```

Voidᵖ is a pointer to a void element.

#### func  NewVoidᵖ

```go
func NewVoidᵖ(addr memory.Pointer) Voidᵖ
```
NewVoidᵖ returns a Voidᵖ that points to addr in the application pool.

#### func (*Voidᵖ) Class

```go
func (*Voidᵖ) Class() binary.Class
```

#### func (Voidᵖ) ElementSize

```go
func (p Voidᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Voidᵖ points to.

#### func (Voidᵖ) OnRead

```go
func (p Voidᵖ) OnRead(ϟs *gfxapi.State) Voidᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (Voidᵖ) OnWrite

```go
func (p Voidᵖ) OnWrite(ϟs *gfxapi.State) Voidᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (Voidᵖ) Slice

```go
func (p Voidᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ
```
Slice returns a new Voidˢ from the pointer using start and end indices.

#### func (Voidᵖ) String

```go
func (p Voidᵖ) String() string
```
String returns a string description of the Voidᵖ pointer.
