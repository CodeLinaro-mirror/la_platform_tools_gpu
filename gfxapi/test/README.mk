# test
--
    import "android.googlesource.com/platform/tools/gpu/gfxapi/test"

Package test is the integration test suite for the api compiler and templates.

## Usage

#### func  API

```go
func API() gfxapi.API
```

#### type BoolArray

```go
type BoolArray []bool
```


#### func (BoolArray) Len

```go
func (s BoolArray) Len() int
```

#### func (BoolArray) Range

```go
func (s BoolArray) Range() []bool
```

#### type CmdArrayOfFloat

```go
type CmdArrayOfFloat struct {
	binary.Generate
	Result F32Array
}
```

//////////////////////////////////////////////////////////////////////////////
CmdArrayOfFloat
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdArrayOfFloat

```go
func NewCmdArrayOfFloat(
	pResult F32Array,
) *CmdArrayOfFloat
```

#### func (*CmdArrayOfFloat) API

```go
func (c *CmdArrayOfFloat) API() gfxapi.API
```

#### func (*CmdArrayOfFloat) Class

```go
func (*CmdArrayOfFloat) Class() binary.Class
```

#### func (*CmdArrayOfFloat) Flags

```go
func (c *CmdArrayOfFloat) Flags() atom.Flags
```

#### func (*CmdArrayOfFloat) Mutate

```go
func (ϟa *CmdArrayOfFloat) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdArrayOfFloat) Replay

```go
func (ϟa *CmdArrayOfFloat) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdArrayOfFloat) String

```go
func (c *CmdArrayOfFloat) String() string
```

#### func (*CmdArrayOfFloat) TypeID

```go
func (c *CmdArrayOfFloat) TypeID() atom.TypeID
```

#### type CmdArrayOfFloat_Postback

```go
type CmdArrayOfFloat_Postback struct {
	Result F32Array
}
```


#### func (*CmdArrayOfFloat_Postback) Decode

```go
func (o *CmdArrayOfFloat_Postback) Decode(result_cnt uint64, d binary.Decoder) error
```

#### type CmdBool

```go
type CmdBool struct {
	binary.Generate
	Result bool
}
```

//////////////////////////////////////////////////////////////////////////////
CmdBool
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdBool

```go
func NewCmdBool(
	pResult bool,
) *CmdBool
```

#### func (*CmdBool) API

```go
func (c *CmdBool) API() gfxapi.API
```

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
func (ϟa *CmdBool) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdBool) Replay

```go
func (ϟa *CmdBool) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdBool) String

```go
func (c *CmdBool) String() string
```

#### func (*CmdBool) TypeID

```go
func (c *CmdBool) TypeID() atom.TypeID
```

#### type CmdBool_Postback

```go
type CmdBool_Postback struct {
	Result bool
}
```


#### func (*CmdBool_Postback) Decode

```go
func (o *CmdBool_Postback) Decode(d binary.Decoder) error
```

#### type CmdF32

```go
type CmdF32 struct {
	binary.Generate
	Result float32
}
```

//////////////////////////////////////////////////////////////////////////////
CmdF32
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdF32

```go
func NewCmdF32(
	pResult float32,
) *CmdF32
```

#### func (*CmdF32) API

```go
func (c *CmdF32) API() gfxapi.API
```

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
func (ϟa *CmdF32) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdF32) Replay

```go
func (ϟa *CmdF32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdF32) String

```go
func (c *CmdF32) String() string
```

#### func (*CmdF32) TypeID

```go
func (c *CmdF32) TypeID() atom.TypeID
```

#### type CmdF32_Postback

```go
type CmdF32_Postback struct {
	Result float32
}
```


#### func (*CmdF32_Postback) Decode

```go
func (o *CmdF32_Postback) Decode(d binary.Decoder) error
```

#### type CmdF64

```go
type CmdF64 struct {
	binary.Generate
	Result float64
}
```

//////////////////////////////////////////////////////////////////////////////
CmdF64
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdF64

```go
func NewCmdF64(
	pResult float64,
) *CmdF64
```

#### func (*CmdF64) API

```go
func (c *CmdF64) API() gfxapi.API
```

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
func (ϟa *CmdF64) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdF64) Replay

```go
func (ϟa *CmdF64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdF64) String

```go
func (c *CmdF64) String() string
```

#### func (*CmdF64) TypeID

```go
func (c *CmdF64) TypeID() atom.TypeID
```

#### type CmdF64_Postback

```go
type CmdF64_Postback struct {
	Result float64
}
```


#### func (*CmdF64_Postback) Decode

```go
func (o *CmdF64_Postback) Decode(d binary.Decoder) error
```

#### type CmdPointer

```go
type CmdPointer struct {
	binary.Generate
	Result memory.Pointer
}
```

//////////////////////////////////////////////////////////////////////////////
CmdPointer
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdPointer

```go
func NewCmdPointer(
	pResult memory.Pointer,
) *CmdPointer
```

#### func (*CmdPointer) API

```go
func (c *CmdPointer) API() gfxapi.API
```

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
func (ϟa *CmdPointer) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdPointer) Replay

```go
func (ϟa *CmdPointer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdPointer) String

```go
func (c *CmdPointer) String() string
```

#### func (*CmdPointer) TypeID

```go
func (c *CmdPointer) TypeID() atom.TypeID
```

#### type CmdPointer_Postback

```go
type CmdPointer_Postback struct {
	Result []byte
}
```


#### func (*CmdPointer_Postback) Decode

```go
func (o *CmdPointer_Postback) Decode(result_cnt uint64, d binary.Decoder) error
```

#### type CmdS16

```go
type CmdS16 struct {
	binary.Generate
	Result int16
}
```

//////////////////////////////////////////////////////////////////////////////
CmdS16
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdS16

```go
func NewCmdS16(
	pResult int16,
) *CmdS16
```

#### func (*CmdS16) API

```go
func (c *CmdS16) API() gfxapi.API
```

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
func (ϟa *CmdS16) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdS16) Replay

```go
func (ϟa *CmdS16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdS16) String

```go
func (c *CmdS16) String() string
```

#### func (*CmdS16) TypeID

```go
func (c *CmdS16) TypeID() atom.TypeID
```

#### type CmdS16_Postback

```go
type CmdS16_Postback struct {
	Result int16
}
```


#### func (*CmdS16_Postback) Decode

```go
func (o *CmdS16_Postback) Decode(d binary.Decoder) error
```

#### type CmdS32

```go
type CmdS32 struct {
	binary.Generate
	Result int32
}
```

//////////////////////////////////////////////////////////////////////////////
CmdS32
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdS32

```go
func NewCmdS32(
	pResult int32,
) *CmdS32
```

#### func (*CmdS32) API

```go
func (c *CmdS32) API() gfxapi.API
```

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
func (ϟa *CmdS32) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdS32) Replay

```go
func (ϟa *CmdS32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdS32) String

```go
func (c *CmdS32) String() string
```

#### func (*CmdS32) TypeID

```go
func (c *CmdS32) TypeID() atom.TypeID
```

#### type CmdS32_Postback

```go
type CmdS32_Postback struct {
	Result int32
}
```


#### func (*CmdS32_Postback) Decode

```go
func (o *CmdS32_Postback) Decode(d binary.Decoder) error
```

#### type CmdS64

```go
type CmdS64 struct {
	binary.Generate
	Result int64
}
```

//////////////////////////////////////////////////////////////////////////////
CmdS64
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdS64

```go
func NewCmdS64(
	pResult int64,
) *CmdS64
```

#### func (*CmdS64) API

```go
func (c *CmdS64) API() gfxapi.API
```

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
func (ϟa *CmdS64) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdS64) Replay

```go
func (ϟa *CmdS64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdS64) String

```go
func (c *CmdS64) String() string
```

#### func (*CmdS64) TypeID

```go
func (c *CmdS64) TypeID() atom.TypeID
```

#### type CmdS64_Postback

```go
type CmdS64_Postback struct {
	Result int64
}
```


#### func (*CmdS64_Postback) Decode

```go
func (o *CmdS64_Postback) Decode(d binary.Decoder) error
```

#### type CmdS8

```go
type CmdS8 struct {
	binary.Generate
	Result int8
}
```

//////////////////////////////////////////////////////////////////////////////
CmdS8
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdS8

```go
func NewCmdS8(
	pResult int8,
) *CmdS8
```

#### func (*CmdS8) API

```go
func (c *CmdS8) API() gfxapi.API
```

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
func (ϟa *CmdS8) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdS8) Replay

```go
func (ϟa *CmdS8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdS8) String

```go
func (c *CmdS8) String() string
```

#### func (*CmdS8) TypeID

```go
func (c *CmdS8) TypeID() atom.TypeID
```

#### type CmdS8_Postback

```go
type CmdS8_Postback struct {
	Result int8
}
```


#### func (*CmdS8_Postback) Decode

```go
func (o *CmdS8_Postback) Decode(d binary.Decoder) error
```

#### type CmdString

```go
type CmdString struct {
	binary.Generate
	Result string
}
```

//////////////////////////////////////////////////////////////////////////////
CmdString
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdString

```go
func NewCmdString(
	pResult string,
) *CmdString
```

#### func (*CmdString) API

```go
func (c *CmdString) API() gfxapi.API
```

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
func (ϟa *CmdString) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdString) Replay

```go
func (ϟa *CmdString) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdString) String

```go
func (c *CmdString) String() string
```

#### func (*CmdString) TypeID

```go
func (c *CmdString) TypeID() atom.TypeID
```

#### type CmdString_Postback

```go
type CmdString_Postback struct {
	Result string
}
```


#### func (*CmdString_Postback) Decode

```go
func (o *CmdString_Postback) Decode(result_cnt uint64, d binary.Decoder) error
```

#### type CmdU16

```go
type CmdU16 struct {
	binary.Generate
	Result uint16
}
```

//////////////////////////////////////////////////////////////////////////////
CmdU16
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdU16

```go
func NewCmdU16(
	pResult uint16,
) *CmdU16
```

#### func (*CmdU16) API

```go
func (c *CmdU16) API() gfxapi.API
```

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
func (ϟa *CmdU16) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdU16) Replay

```go
func (ϟa *CmdU16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdU16) String

```go
func (c *CmdU16) String() string
```

#### func (*CmdU16) TypeID

```go
func (c *CmdU16) TypeID() atom.TypeID
```

#### type CmdU16_Postback

```go
type CmdU16_Postback struct {
	Result uint16
}
```


#### func (*CmdU16_Postback) Decode

```go
func (o *CmdU16_Postback) Decode(d binary.Decoder) error
```

#### type CmdU32

```go
type CmdU32 struct {
	binary.Generate
	Result uint32
}
```

//////////////////////////////////////////////////////////////////////////////
CmdU32
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdU32

```go
func NewCmdU32(
	pResult uint32,
) *CmdU32
```

#### func (*CmdU32) API

```go
func (c *CmdU32) API() gfxapi.API
```

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
func (ϟa *CmdU32) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdU32) Replay

```go
func (ϟa *CmdU32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdU32) String

```go
func (c *CmdU32) String() string
```

#### func (*CmdU32) TypeID

```go
func (c *CmdU32) TypeID() atom.TypeID
```

#### type CmdU32_Postback

```go
type CmdU32_Postback struct {
	Result uint32
}
```


#### func (*CmdU32_Postback) Decode

```go
func (o *CmdU32_Postback) Decode(d binary.Decoder) error
```

#### type CmdU64

```go
type CmdU64 struct {
	binary.Generate
	Result uint64
}
```

//////////////////////////////////////////////////////////////////////////////
CmdU64
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdU64

```go
func NewCmdU64(
	pResult uint64,
) *CmdU64
```

#### func (*CmdU64) API

```go
func (c *CmdU64) API() gfxapi.API
```

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
func (ϟa *CmdU64) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdU64) Replay

```go
func (ϟa *CmdU64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdU64) String

```go
func (c *CmdU64) String() string
```

#### func (*CmdU64) TypeID

```go
func (c *CmdU64) TypeID() atom.TypeID
```

#### type CmdU64_Postback

```go
type CmdU64_Postback struct {
	Result uint64
}
```


#### func (*CmdU64_Postback) Decode

```go
func (o *CmdU64_Postback) Decode(d binary.Decoder) error
```

#### type CmdU8

```go
type CmdU8 struct {
	binary.Generate
	Result uint8
}
```

//////////////////////////////////////////////////////////////////////////////
CmdU8
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdU8

```go
func NewCmdU8(
	pResult uint8,
) *CmdU8
```

#### func (*CmdU8) API

```go
func (c *CmdU8) API() gfxapi.API
```

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
func (ϟa *CmdU8) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdU8) Replay

```go
func (ϟa *CmdU8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdU8) String

```go
func (c *CmdU8) String() string
```

#### func (*CmdU8) TypeID

```go
func (c *CmdU8) TypeID() atom.TypeID
```

#### type CmdU8_Postback

```go
type CmdU8_Postback struct {
	Result uint8
}
```


#### func (*CmdU8_Postback) Decode

```go
func (o *CmdU8_Postback) Decode(d binary.Decoder) error
```

#### type CmdVoid

```go
type CmdVoid struct {
	binary.Generate
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
func (ϟa *CmdVoid) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdVoid) Replay

```go
func (ϟa *CmdVoid) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdVoid) String

```go
func (c *CmdVoid) String() string
```

#### func (*CmdVoid) TypeID

```go
func (c *CmdVoid) TypeID() atom.TypeID
```

#### type CmdVoid3Arrays

```go
type CmdVoid3Arrays struct {
	binary.Generate
	A S8Array
	B StringArray
	C BoolArray
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoid3Arrays
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoid3Arrays

```go
func NewCmdVoid3Arrays(
	pA S8Array,
	pB StringArray,
	pC BoolArray,
) *CmdVoid3Arrays
```

#### func (*CmdVoid3Arrays) API

```go
func (c *CmdVoid3Arrays) API() gfxapi.API
```

#### func (*CmdVoid3Arrays) Class

```go
func (*CmdVoid3Arrays) Class() binary.Class
```

#### func (*CmdVoid3Arrays) Flags

```go
func (c *CmdVoid3Arrays) Flags() atom.Flags
```

#### func (*CmdVoid3Arrays) Mutate

```go
func (ϟa *CmdVoid3Arrays) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdVoid3Arrays) Replay

```go
func (ϟa *CmdVoid3Arrays) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdVoid3Arrays) String

```go
func (c *CmdVoid3Arrays) String() string
```

#### func (*CmdVoid3Arrays) TypeID

```go
func (c *CmdVoid3Arrays) TypeID() atom.TypeID
```

#### type CmdVoid3Remapped

```go
type CmdVoid3Remapped struct {
	binary.Generate
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
func NewCmdVoid3Remapped(
	pA remapped,
	pB remapped,
	pC remapped,
) *CmdVoid3Remapped
```

#### func (*CmdVoid3Remapped) API

```go
func (c *CmdVoid3Remapped) API() gfxapi.API
```

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
func (ϟa *CmdVoid3Remapped) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdVoid3Remapped) Replay

```go
func (ϟa *CmdVoid3Remapped) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdVoid3Remapped) String

```go
func (c *CmdVoid3Remapped) String() string
```

#### func (*CmdVoid3Remapped) TypeID

```go
func (c *CmdVoid3Remapped) TypeID() atom.TypeID
```

#### type CmdVoid3Strings

```go
type CmdVoid3Strings struct {
	binary.Generate
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
func NewCmdVoid3Strings(
	pA string,
	pB string,
	pC string,
) *CmdVoid3Strings
```

#### func (*CmdVoid3Strings) API

```go
func (c *CmdVoid3Strings) API() gfxapi.API
```

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
func (ϟa *CmdVoid3Strings) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdVoid3Strings) Replay

```go
func (ϟa *CmdVoid3Strings) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdVoid3Strings) String

```go
func (c *CmdVoid3Strings) String() string
```

#### func (*CmdVoid3Strings) TypeID

```go
func (c *CmdVoid3Strings) TypeID() atom.TypeID
```

#### type CmdVoidArrayOfStrings

```go
type CmdVoidArrayOfStrings struct {
	binary.Generate
	A StringArray
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidArrayOfStrings
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidArrayOfStrings

```go
func NewCmdVoidArrayOfStrings(
	pA StringArray,
) *CmdVoidArrayOfStrings
```

#### func (*CmdVoidArrayOfStrings) API

```go
func (c *CmdVoidArrayOfStrings) API() gfxapi.API
```

#### func (*CmdVoidArrayOfStrings) Class

```go
func (*CmdVoidArrayOfStrings) Class() binary.Class
```

#### func (*CmdVoidArrayOfStrings) Flags

```go
func (c *CmdVoidArrayOfStrings) Flags() atom.Flags
```

#### func (*CmdVoidArrayOfStrings) Mutate

```go
func (ϟa *CmdVoidArrayOfStrings) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdVoidArrayOfStrings) Replay

```go
func (ϟa *CmdVoidArrayOfStrings) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdVoidArrayOfStrings) String

```go
func (c *CmdVoidArrayOfStrings) String() string
```

#### func (*CmdVoidArrayOfStrings) TypeID

```go
func (c *CmdVoidArrayOfStrings) TypeID() atom.TypeID
```

#### type CmdVoidBool

```go
type CmdVoidBool struct {
	binary.Generate
	A bool
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidBool
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidBool

```go
func NewCmdVoidBool(
	pA bool,
) *CmdVoidBool
```

#### func (*CmdVoidBool) API

```go
func (c *CmdVoidBool) API() gfxapi.API
```

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
func (ϟa *CmdVoidBool) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdVoidBool) Replay

```go
func (ϟa *CmdVoidBool) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdVoidBool) String

```go
func (c *CmdVoidBool) String() string
```

#### func (*CmdVoidBool) TypeID

```go
func (c *CmdVoidBool) TypeID() atom.TypeID
```

#### type CmdVoidF32

```go
type CmdVoidF32 struct {
	binary.Generate
	A float32
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidF32
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidF32

```go
func NewCmdVoidF32(
	pA float32,
) *CmdVoidF32
```

#### func (*CmdVoidF32) API

```go
func (c *CmdVoidF32) API() gfxapi.API
```

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
func (ϟa *CmdVoidF32) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdVoidF32) Replay

```go
func (ϟa *CmdVoidF32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdVoidF32) String

```go
func (c *CmdVoidF32) String() string
```

#### func (*CmdVoidF32) TypeID

```go
func (c *CmdVoidF32) TypeID() atom.TypeID
```

#### type CmdVoidF64

```go
type CmdVoidF64 struct {
	binary.Generate
	A float64
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidF64
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidF64

```go
func NewCmdVoidF64(
	pA float64,
) *CmdVoidF64
```

#### func (*CmdVoidF64) API

```go
func (c *CmdVoidF64) API() gfxapi.API
```

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
func (ϟa *CmdVoidF64) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdVoidF64) Replay

```go
func (ϟa *CmdVoidF64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdVoidF64) String

```go
func (c *CmdVoidF64) String() string
```

#### func (*CmdVoidF64) TypeID

```go
func (c *CmdVoidF64) TypeID() atom.TypeID
```

#### type CmdVoidOut3Remapped

```go
type CmdVoidOut3Remapped struct {
	binary.Generate
	A remapped
	B remapped
	C remapped
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidOut3Remapped
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidOut3Remapped

```go
func NewCmdVoidOut3Remapped(
	pA remapped,
	pB remapped,
	pC remapped,
) *CmdVoidOut3Remapped
```

#### func (*CmdVoidOut3Remapped) API

```go
func (c *CmdVoidOut3Remapped) API() gfxapi.API
```

#### func (*CmdVoidOut3Remapped) Class

```go
func (*CmdVoidOut3Remapped) Class() binary.Class
```

#### func (*CmdVoidOut3Remapped) Flags

```go
func (c *CmdVoidOut3Remapped) Flags() atom.Flags
```

#### func (*CmdVoidOut3Remapped) Mutate

```go
func (ϟa *CmdVoidOut3Remapped) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdVoidOut3Remapped) Replay

```go
func (ϟa *CmdVoidOut3Remapped) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdVoidOut3Remapped) String

```go
func (c *CmdVoidOut3Remapped) String() string
```

#### func (*CmdVoidOut3Remapped) TypeID

```go
func (c *CmdVoidOut3Remapped) TypeID() atom.TypeID
```

#### type CmdVoidOut3Remapped_Postback

```go
type CmdVoidOut3Remapped_Postback struct {
	A remapped
	B remapped
	C remapped
}
```


#### func (*CmdVoidOut3Remapped_Postback) Decode

```go
func (o *CmdVoidOut3Remapped_Postback) Decode(d binary.Decoder) error
```

#### type CmdVoidOut3Strings

```go
type CmdVoidOut3Strings struct {
	binary.Generate
	A string
	B string
	C string
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidOut3Strings
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidOut3Strings

```go
func NewCmdVoidOut3Strings(
	pA string,
	pB string,
	pC string,
) *CmdVoidOut3Strings
```

#### func (*CmdVoidOut3Strings) API

```go
func (c *CmdVoidOut3Strings) API() gfxapi.API
```

#### func (*CmdVoidOut3Strings) Class

```go
func (*CmdVoidOut3Strings) Class() binary.Class
```

#### func (*CmdVoidOut3Strings) Flags

```go
func (c *CmdVoidOut3Strings) Flags() atom.Flags
```

#### func (*CmdVoidOut3Strings) Mutate

```go
func (ϟa *CmdVoidOut3Strings) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdVoidOut3Strings) Replay

```go
func (ϟa *CmdVoidOut3Strings) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdVoidOut3Strings) String

```go
func (c *CmdVoidOut3Strings) String() string
```

#### func (*CmdVoidOut3Strings) TypeID

```go
func (c *CmdVoidOut3Strings) TypeID() atom.TypeID
```

#### type CmdVoidOut3Strings_Postback

```go
type CmdVoidOut3Strings_Postback struct {
	A string
	B string
	C string
}
```


#### func (*CmdVoidOut3Strings_Postback) Decode

```go
func (o *CmdVoidOut3Strings_Postback) Decode(a_cnt uint64,
	b_cnt uint64,
	c_cnt uint64, d binary.Decoder) error
```

#### type CmdVoidOutArrayOfRemapped

```go
type CmdVoidOutArrayOfRemapped struct {
	binary.Generate
	A RemappedArray
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidOutArrayOfRemapped
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidOutArrayOfRemapped

```go
func NewCmdVoidOutArrayOfRemapped(
	pA RemappedArray,
) *CmdVoidOutArrayOfRemapped
```

#### func (*CmdVoidOutArrayOfRemapped) API

```go
func (c *CmdVoidOutArrayOfRemapped) API() gfxapi.API
```

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
func (ϟa *CmdVoidOutArrayOfRemapped) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdVoidOutArrayOfRemapped) Replay

```go
func (ϟa *CmdVoidOutArrayOfRemapped) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdVoidOutArrayOfRemapped) String

```go
func (c *CmdVoidOutArrayOfRemapped) String() string
```

#### func (*CmdVoidOutArrayOfRemapped) TypeID

```go
func (c *CmdVoidOutArrayOfRemapped) TypeID() atom.TypeID
```

#### type CmdVoidOutArrayOfRemapped_Postback

```go
type CmdVoidOutArrayOfRemapped_Postback struct {
	A RemappedArray
}
```


#### func (*CmdVoidOutArrayOfRemapped_Postback) Decode

```go
func (o *CmdVoidOutArrayOfRemapped_Postback) Decode(a_cnt uint64, d binary.Decoder) error
```

#### type CmdVoidOutBool

```go
type CmdVoidOutBool struct {
	binary.Generate
	A bool
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidOutBool
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidOutBool

```go
func NewCmdVoidOutBool(
	pA bool,
) *CmdVoidOutBool
```

#### func (*CmdVoidOutBool) API

```go
func (c *CmdVoidOutBool) API() gfxapi.API
```

#### func (*CmdVoidOutBool) Class

```go
func (*CmdVoidOutBool) Class() binary.Class
```

#### func (*CmdVoidOutBool) Flags

```go
func (c *CmdVoidOutBool) Flags() atom.Flags
```

#### func (*CmdVoidOutBool) Mutate

```go
func (ϟa *CmdVoidOutBool) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdVoidOutBool) Replay

```go
func (ϟa *CmdVoidOutBool) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdVoidOutBool) String

```go
func (c *CmdVoidOutBool) String() string
```

#### func (*CmdVoidOutBool) TypeID

```go
func (c *CmdVoidOutBool) TypeID() atom.TypeID
```

#### type CmdVoidOutBool_Postback

```go
type CmdVoidOutBool_Postback struct {
	A bool
}
```


#### func (*CmdVoidOutBool_Postback) Decode

```go
func (o *CmdVoidOutBool_Postback) Decode(d binary.Decoder) error
```

#### type CmdVoidOutF32

```go
type CmdVoidOutF32 struct {
	binary.Generate
	A float32
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidOutF32
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidOutF32

```go
func NewCmdVoidOutF32(
	pA float32,
) *CmdVoidOutF32
```

#### func (*CmdVoidOutF32) API

```go
func (c *CmdVoidOutF32) API() gfxapi.API
```

#### func (*CmdVoidOutF32) Class

```go
func (*CmdVoidOutF32) Class() binary.Class
```

#### func (*CmdVoidOutF32) Flags

```go
func (c *CmdVoidOutF32) Flags() atom.Flags
```

#### func (*CmdVoidOutF32) Mutate

```go
func (ϟa *CmdVoidOutF32) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdVoidOutF32) Replay

```go
func (ϟa *CmdVoidOutF32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdVoidOutF32) String

```go
func (c *CmdVoidOutF32) String() string
```

#### func (*CmdVoidOutF32) TypeID

```go
func (c *CmdVoidOutF32) TypeID() atom.TypeID
```

#### type CmdVoidOutF32_Postback

```go
type CmdVoidOutF32_Postback struct {
	A float32
}
```


#### func (*CmdVoidOutF32_Postback) Decode

```go
func (o *CmdVoidOutF32_Postback) Decode(d binary.Decoder) error
```

#### type CmdVoidOutF64

```go
type CmdVoidOutF64 struct {
	binary.Generate
	A float64
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidOutF64
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidOutF64

```go
func NewCmdVoidOutF64(
	pA float64,
) *CmdVoidOutF64
```

#### func (*CmdVoidOutF64) API

```go
func (c *CmdVoidOutF64) API() gfxapi.API
```

#### func (*CmdVoidOutF64) Class

```go
func (*CmdVoidOutF64) Class() binary.Class
```

#### func (*CmdVoidOutF64) Flags

```go
func (c *CmdVoidOutF64) Flags() atom.Flags
```

#### func (*CmdVoidOutF64) Mutate

```go
func (ϟa *CmdVoidOutF64) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdVoidOutF64) Replay

```go
func (ϟa *CmdVoidOutF64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdVoidOutF64) String

```go
func (c *CmdVoidOutF64) String() string
```

#### func (*CmdVoidOutF64) TypeID

```go
func (c *CmdVoidOutF64) TypeID() atom.TypeID
```

#### type CmdVoidOutF64_Postback

```go
type CmdVoidOutF64_Postback struct {
	A float64
}
```


#### func (*CmdVoidOutF64_Postback) Decode

```go
func (o *CmdVoidOutF64_Postback) Decode(d binary.Decoder) error
```

#### type CmdVoidOutFixedSizeBuffer

```go
type CmdVoidOutFixedSizeBuffer struct {
	binary.Generate
	A memory.Pointer
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidOutFixedSizeBuffer
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidOutFixedSizeBuffer

```go
func NewCmdVoidOutFixedSizeBuffer(
	pA memory.Pointer,
) *CmdVoidOutFixedSizeBuffer
```

#### func (*CmdVoidOutFixedSizeBuffer) API

```go
func (c *CmdVoidOutFixedSizeBuffer) API() gfxapi.API
```

#### func (*CmdVoidOutFixedSizeBuffer) Class

```go
func (*CmdVoidOutFixedSizeBuffer) Class() binary.Class
```

#### func (*CmdVoidOutFixedSizeBuffer) Flags

```go
func (c *CmdVoidOutFixedSizeBuffer) Flags() atom.Flags
```

#### func (*CmdVoidOutFixedSizeBuffer) Mutate

```go
func (ϟa *CmdVoidOutFixedSizeBuffer) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdVoidOutFixedSizeBuffer) Replay

```go
func (ϟa *CmdVoidOutFixedSizeBuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdVoidOutFixedSizeBuffer) String

```go
func (c *CmdVoidOutFixedSizeBuffer) String() string
```

#### func (*CmdVoidOutFixedSizeBuffer) TypeID

```go
func (c *CmdVoidOutFixedSizeBuffer) TypeID() atom.TypeID
```

#### type CmdVoidOutFixedSizeBuffer_Postback

```go
type CmdVoidOutFixedSizeBuffer_Postback struct {
	A []byte
}
```


#### func (*CmdVoidOutFixedSizeBuffer_Postback) Decode

```go
func (o *CmdVoidOutFixedSizeBuffer_Postback) Decode(a_cnt uint64, d binary.Decoder) error
```

#### type CmdVoidOutS16

```go
type CmdVoidOutS16 struct {
	binary.Generate
	A int16
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidOutS16
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidOutS16

```go
func NewCmdVoidOutS16(
	pA int16,
) *CmdVoidOutS16
```

#### func (*CmdVoidOutS16) API

```go
func (c *CmdVoidOutS16) API() gfxapi.API
```

#### func (*CmdVoidOutS16) Class

```go
func (*CmdVoidOutS16) Class() binary.Class
```

#### func (*CmdVoidOutS16) Flags

```go
func (c *CmdVoidOutS16) Flags() atom.Flags
```

#### func (*CmdVoidOutS16) Mutate

```go
func (ϟa *CmdVoidOutS16) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdVoidOutS16) Replay

```go
func (ϟa *CmdVoidOutS16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdVoidOutS16) String

```go
func (c *CmdVoidOutS16) String() string
```

#### func (*CmdVoidOutS16) TypeID

```go
func (c *CmdVoidOutS16) TypeID() atom.TypeID
```

#### type CmdVoidOutS16_Postback

```go
type CmdVoidOutS16_Postback struct {
	A int16
}
```


#### func (*CmdVoidOutS16_Postback) Decode

```go
func (o *CmdVoidOutS16_Postback) Decode(d binary.Decoder) error
```

#### type CmdVoidOutS32

```go
type CmdVoidOutS32 struct {
	binary.Generate
	A int32
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidOutS32
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidOutS32

```go
func NewCmdVoidOutS32(
	pA int32,
) *CmdVoidOutS32
```

#### func (*CmdVoidOutS32) API

```go
func (c *CmdVoidOutS32) API() gfxapi.API
```

#### func (*CmdVoidOutS32) Class

```go
func (*CmdVoidOutS32) Class() binary.Class
```

#### func (*CmdVoidOutS32) Flags

```go
func (c *CmdVoidOutS32) Flags() atom.Flags
```

#### func (*CmdVoidOutS32) Mutate

```go
func (ϟa *CmdVoidOutS32) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdVoidOutS32) Replay

```go
func (ϟa *CmdVoidOutS32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdVoidOutS32) String

```go
func (c *CmdVoidOutS32) String() string
```

#### func (*CmdVoidOutS32) TypeID

```go
func (c *CmdVoidOutS32) TypeID() atom.TypeID
```

#### type CmdVoidOutS32_Postback

```go
type CmdVoidOutS32_Postback struct {
	A int32
}
```


#### func (*CmdVoidOutS32_Postback) Decode

```go
func (o *CmdVoidOutS32_Postback) Decode(d binary.Decoder) error
```

#### type CmdVoidOutS64

```go
type CmdVoidOutS64 struct {
	binary.Generate
	A int64
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidOutS64
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidOutS64

```go
func NewCmdVoidOutS64(
	pA int64,
) *CmdVoidOutS64
```

#### func (*CmdVoidOutS64) API

```go
func (c *CmdVoidOutS64) API() gfxapi.API
```

#### func (*CmdVoidOutS64) Class

```go
func (*CmdVoidOutS64) Class() binary.Class
```

#### func (*CmdVoidOutS64) Flags

```go
func (c *CmdVoidOutS64) Flags() atom.Flags
```

#### func (*CmdVoidOutS64) Mutate

```go
func (ϟa *CmdVoidOutS64) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdVoidOutS64) Replay

```go
func (ϟa *CmdVoidOutS64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdVoidOutS64) String

```go
func (c *CmdVoidOutS64) String() string
```

#### func (*CmdVoidOutS64) TypeID

```go
func (c *CmdVoidOutS64) TypeID() atom.TypeID
```

#### type CmdVoidOutS64_Postback

```go
type CmdVoidOutS64_Postback struct {
	A int64
}
```


#### func (*CmdVoidOutS64_Postback) Decode

```go
func (o *CmdVoidOutS64_Postback) Decode(d binary.Decoder) error
```

#### type CmdVoidOutS8

```go
type CmdVoidOutS8 struct {
	binary.Generate
	A int8
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidOutS8
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidOutS8

```go
func NewCmdVoidOutS8(
	pA int8,
) *CmdVoidOutS8
```

#### func (*CmdVoidOutS8) API

```go
func (c *CmdVoidOutS8) API() gfxapi.API
```

#### func (*CmdVoidOutS8) Class

```go
func (*CmdVoidOutS8) Class() binary.Class
```

#### func (*CmdVoidOutS8) Flags

```go
func (c *CmdVoidOutS8) Flags() atom.Flags
```

#### func (*CmdVoidOutS8) Mutate

```go
func (ϟa *CmdVoidOutS8) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdVoidOutS8) Replay

```go
func (ϟa *CmdVoidOutS8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdVoidOutS8) String

```go
func (c *CmdVoidOutS8) String() string
```

#### func (*CmdVoidOutS8) TypeID

```go
func (c *CmdVoidOutS8) TypeID() atom.TypeID
```

#### type CmdVoidOutS8_Postback

```go
type CmdVoidOutS8_Postback struct {
	A int8
}
```


#### func (*CmdVoidOutS8_Postback) Decode

```go
func (o *CmdVoidOutS8_Postback) Decode(d binary.Decoder) error
```

#### type CmdVoidOutString

```go
type CmdVoidOutString struct {
	binary.Generate
	A string
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidOutString
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidOutString

```go
func NewCmdVoidOutString(
	pA string,
) *CmdVoidOutString
```

#### func (*CmdVoidOutString) API

```go
func (c *CmdVoidOutString) API() gfxapi.API
```

#### func (*CmdVoidOutString) Class

```go
func (*CmdVoidOutString) Class() binary.Class
```

#### func (*CmdVoidOutString) Flags

```go
func (c *CmdVoidOutString) Flags() atom.Flags
```

#### func (*CmdVoidOutString) Mutate

```go
func (ϟa *CmdVoidOutString) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdVoidOutString) Replay

```go
func (ϟa *CmdVoidOutString) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdVoidOutString) String

```go
func (c *CmdVoidOutString) String() string
```

#### func (*CmdVoidOutString) TypeID

```go
func (c *CmdVoidOutString) TypeID() atom.TypeID
```

#### type CmdVoidOutString_Postback

```go
type CmdVoidOutString_Postback struct {
	A string
}
```


#### func (*CmdVoidOutString_Postback) Decode

```go
func (o *CmdVoidOutString_Postback) Decode(a_cnt uint64, d binary.Decoder) error
```

#### type CmdVoidOutU16

```go
type CmdVoidOutU16 struct {
	binary.Generate
	A uint16
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidOutU16
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidOutU16

```go
func NewCmdVoidOutU16(
	pA uint16,
) *CmdVoidOutU16
```

#### func (*CmdVoidOutU16) API

```go
func (c *CmdVoidOutU16) API() gfxapi.API
```

#### func (*CmdVoidOutU16) Class

```go
func (*CmdVoidOutU16) Class() binary.Class
```

#### func (*CmdVoidOutU16) Flags

```go
func (c *CmdVoidOutU16) Flags() atom.Flags
```

#### func (*CmdVoidOutU16) Mutate

```go
func (ϟa *CmdVoidOutU16) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdVoidOutU16) Replay

```go
func (ϟa *CmdVoidOutU16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdVoidOutU16) String

```go
func (c *CmdVoidOutU16) String() string
```

#### func (*CmdVoidOutU16) TypeID

```go
func (c *CmdVoidOutU16) TypeID() atom.TypeID
```

#### type CmdVoidOutU16_Postback

```go
type CmdVoidOutU16_Postback struct {
	A uint16
}
```


#### func (*CmdVoidOutU16_Postback) Decode

```go
func (o *CmdVoidOutU16_Postback) Decode(d binary.Decoder) error
```

#### type CmdVoidOutU32

```go
type CmdVoidOutU32 struct {
	binary.Generate
	A uint32
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidOutU32
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidOutU32

```go
func NewCmdVoidOutU32(
	pA uint32,
) *CmdVoidOutU32
```

#### func (*CmdVoidOutU32) API

```go
func (c *CmdVoidOutU32) API() gfxapi.API
```

#### func (*CmdVoidOutU32) Class

```go
func (*CmdVoidOutU32) Class() binary.Class
```

#### func (*CmdVoidOutU32) Flags

```go
func (c *CmdVoidOutU32) Flags() atom.Flags
```

#### func (*CmdVoidOutU32) Mutate

```go
func (ϟa *CmdVoidOutU32) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdVoidOutU32) Replay

```go
func (ϟa *CmdVoidOutU32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdVoidOutU32) String

```go
func (c *CmdVoidOutU32) String() string
```

#### func (*CmdVoidOutU32) TypeID

```go
func (c *CmdVoidOutU32) TypeID() atom.TypeID
```

#### type CmdVoidOutU32_Postback

```go
type CmdVoidOutU32_Postback struct {
	A uint32
}
```


#### func (*CmdVoidOutU32_Postback) Decode

```go
func (o *CmdVoidOutU32_Postback) Decode(d binary.Decoder) error
```

#### type CmdVoidOutU64

```go
type CmdVoidOutU64 struct {
	binary.Generate
	A uint64
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidOutU64
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidOutU64

```go
func NewCmdVoidOutU64(
	pA uint64,
) *CmdVoidOutU64
```

#### func (*CmdVoidOutU64) API

```go
func (c *CmdVoidOutU64) API() gfxapi.API
```

#### func (*CmdVoidOutU64) Class

```go
func (*CmdVoidOutU64) Class() binary.Class
```

#### func (*CmdVoidOutU64) Flags

```go
func (c *CmdVoidOutU64) Flags() atom.Flags
```

#### func (*CmdVoidOutU64) Mutate

```go
func (ϟa *CmdVoidOutU64) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdVoidOutU64) Replay

```go
func (ϟa *CmdVoidOutU64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdVoidOutU64) String

```go
func (c *CmdVoidOutU64) String() string
```

#### func (*CmdVoidOutU64) TypeID

```go
func (c *CmdVoidOutU64) TypeID() atom.TypeID
```

#### type CmdVoidOutU64_Postback

```go
type CmdVoidOutU64_Postback struct {
	A uint64
}
```


#### func (*CmdVoidOutU64_Postback) Decode

```go
func (o *CmdVoidOutU64_Postback) Decode(d binary.Decoder) error
```

#### type CmdVoidOutU8

```go
type CmdVoidOutU8 struct {
	binary.Generate
	A uint8
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidOutU8
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidOutU8

```go
func NewCmdVoidOutU8(
	pA uint8,
) *CmdVoidOutU8
```

#### func (*CmdVoidOutU8) API

```go
func (c *CmdVoidOutU8) API() gfxapi.API
```

#### func (*CmdVoidOutU8) Class

```go
func (*CmdVoidOutU8) Class() binary.Class
```

#### func (*CmdVoidOutU8) Flags

```go
func (c *CmdVoidOutU8) Flags() atom.Flags
```

#### func (*CmdVoidOutU8) Mutate

```go
func (ϟa *CmdVoidOutU8) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdVoidOutU8) Replay

```go
func (ϟa *CmdVoidOutU8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdVoidOutU8) String

```go
func (c *CmdVoidOutU8) String() string
```

#### func (*CmdVoidOutU8) TypeID

```go
func (c *CmdVoidOutU8) TypeID() atom.TypeID
```

#### type CmdVoidOutU8_Postback

```go
type CmdVoidOutU8_Postback struct {
	A uint8
}
```


#### func (*CmdVoidOutU8_Postback) Decode

```go
func (o *CmdVoidOutU8_Postback) Decode(d binary.Decoder) error
```

#### type CmdVoidS16

```go
type CmdVoidS16 struct {
	binary.Generate
	A int16
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidS16
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidS16

```go
func NewCmdVoidS16(
	pA int16,
) *CmdVoidS16
```

#### func (*CmdVoidS16) API

```go
func (c *CmdVoidS16) API() gfxapi.API
```

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
func (ϟa *CmdVoidS16) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdVoidS16) Replay

```go
func (ϟa *CmdVoidS16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdVoidS16) String

```go
func (c *CmdVoidS16) String() string
```

#### func (*CmdVoidS16) TypeID

```go
func (c *CmdVoidS16) TypeID() atom.TypeID
```

#### type CmdVoidS32

```go
type CmdVoidS32 struct {
	binary.Generate
	A int32
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidS32
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidS32

```go
func NewCmdVoidS32(
	pA int32,
) *CmdVoidS32
```

#### func (*CmdVoidS32) API

```go
func (c *CmdVoidS32) API() gfxapi.API
```

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
func (ϟa *CmdVoidS32) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdVoidS32) Replay

```go
func (ϟa *CmdVoidS32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdVoidS32) String

```go
func (c *CmdVoidS32) String() string
```

#### func (*CmdVoidS32) TypeID

```go
func (c *CmdVoidS32) TypeID() atom.TypeID
```

#### type CmdVoidS64

```go
type CmdVoidS64 struct {
	binary.Generate
	A int64
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidS64
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidS64

```go
func NewCmdVoidS64(
	pA int64,
) *CmdVoidS64
```

#### func (*CmdVoidS64) API

```go
func (c *CmdVoidS64) API() gfxapi.API
```

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
func (ϟa *CmdVoidS64) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdVoidS64) Replay

```go
func (ϟa *CmdVoidS64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdVoidS64) String

```go
func (c *CmdVoidS64) String() string
```

#### func (*CmdVoidS64) TypeID

```go
func (c *CmdVoidS64) TypeID() atom.TypeID
```

#### type CmdVoidS8

```go
type CmdVoidS8 struct {
	binary.Generate
	A int8
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidS8
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidS8

```go
func NewCmdVoidS8(
	pA int8,
) *CmdVoidS8
```

#### func (*CmdVoidS8) API

```go
func (c *CmdVoidS8) API() gfxapi.API
```

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
func (ϟa *CmdVoidS8) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdVoidS8) Replay

```go
func (ϟa *CmdVoidS8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdVoidS8) String

```go
func (c *CmdVoidS8) String() string
```

#### func (*CmdVoidS8) TypeID

```go
func (c *CmdVoidS8) TypeID() atom.TypeID
```

#### type CmdVoidString

```go
type CmdVoidString struct {
	binary.Generate
	A string
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidString
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidString

```go
func NewCmdVoidString(
	pA string,
) *CmdVoidString
```

#### func (*CmdVoidString) API

```go
func (c *CmdVoidString) API() gfxapi.API
```

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
func (ϟa *CmdVoidString) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdVoidString) Replay

```go
func (ϟa *CmdVoidString) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdVoidString) String

```go
func (c *CmdVoidString) String() string
```

#### func (*CmdVoidString) TypeID

```go
func (c *CmdVoidString) TypeID() atom.TypeID
```

#### type CmdVoidU16

```go
type CmdVoidU16 struct {
	binary.Generate
	A uint16
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidU16
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidU16

```go
func NewCmdVoidU16(
	pA uint16,
) *CmdVoidU16
```

#### func (*CmdVoidU16) API

```go
func (c *CmdVoidU16) API() gfxapi.API
```

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
func (ϟa *CmdVoidU16) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdVoidU16) Replay

```go
func (ϟa *CmdVoidU16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdVoidU16) String

```go
func (c *CmdVoidU16) String() string
```

#### func (*CmdVoidU16) TypeID

```go
func (c *CmdVoidU16) TypeID() atom.TypeID
```

#### type CmdVoidU32

```go
type CmdVoidU32 struct {
	binary.Generate
	A uint32
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidU32
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidU32

```go
func NewCmdVoidU32(
	pA uint32,
) *CmdVoidU32
```

#### func (*CmdVoidU32) API

```go
func (c *CmdVoidU32) API() gfxapi.API
```

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
func (ϟa *CmdVoidU32) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdVoidU32) Replay

```go
func (ϟa *CmdVoidU32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdVoidU32) String

```go
func (c *CmdVoidU32) String() string
```

#### func (*CmdVoidU32) TypeID

```go
func (c *CmdVoidU32) TypeID() atom.TypeID
```

#### type CmdVoidU64

```go
type CmdVoidU64 struct {
	binary.Generate
	A uint64
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidU64
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidU64

```go
func NewCmdVoidU64(
	pA uint64,
) *CmdVoidU64
```

#### func (*CmdVoidU64) API

```go
func (c *CmdVoidU64) API() gfxapi.API
```

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
func (ϟa *CmdVoidU64) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdVoidU64) Replay

```go
func (ϟa *CmdVoidU64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdVoidU64) String

```go
func (c *CmdVoidU64) String() string
```

#### func (*CmdVoidU64) TypeID

```go
func (c *CmdVoidU64) TypeID() atom.TypeID
```

#### type CmdVoidU8

```go
type CmdVoidU8 struct {
	binary.Generate
	A uint8
}
```

//////////////////////////////////////////////////////////////////////////////
CmdVoidU8
//////////////////////////////////////////////////////////////////////////////

#### func  NewCmdVoidU8

```go
func NewCmdVoidU8(
	pA uint8,
) *CmdVoidU8
```

#### func (*CmdVoidU8) API

```go
func (c *CmdVoidU8) API() gfxapi.API
```

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
func (ϟa *CmdVoidU8) Mutate(ϟs *gfxapi.State) error
```

#### func (*CmdVoidU8) Replay

```go
func (ϟa *CmdVoidU8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*CmdVoidU8) String

```go
func (c *CmdVoidU8) String() string
```

#### func (*CmdVoidU8) TypeID

```go
func (c *CmdVoidU8) TypeID() atom.TypeID
```

#### type F32Array

```go
type F32Array []float32
```


#### func (F32Array) Len

```go
func (s F32Array) Len() int
```

#### func (F32Array) Range

```go
func (s F32Array) Range() []float32
```

#### type Globals

```go
type Globals struct {
	binary.Generate
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

#### type RemappedArray

```go
type RemappedArray []remapped
```


#### func (RemappedArray) Len

```go
func (s RemappedArray) Len() int
```

#### func (RemappedArray) Range

```go
func (s RemappedArray) Range() []remapped
```

#### type S8Array

```go
type S8Array []int8
```


#### func (S8Array) Len

```go
func (s S8Array) Len() int
```

#### func (S8Array) Range

```go
func (s S8Array) Range() []int8
```

#### type State

```go
type State struct {
	Globals
	ValidateOutput bool
}
```


#### type StringArray

```go
type StringArray []string
```


#### func (StringArray) Len

```go
func (s StringArray) Len() int
```

#### func (StringArray) Range

```go
func (s StringArray) Range() []string
```
