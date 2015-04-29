////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

package test

import (
	"fmt"
	"strings"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/service"
)

type remapped uint32

func (c *remapped) Less(rhs remapped) bool  { return uint32(*c) < uint32(rhs) }
func (c *remapped) Equal(rhs remapped) bool { return uint32(*c) == uint32(rhs) }

type BoolArray []bool

func (s BoolArray) Len() int      { return len(s) }
func (s BoolArray) Range() []bool { return s }

type F32Array []float32

func (s F32Array) Len() int         { return len(s) }
func (s F32Array) Range() []float32 { return s }

type RemappedArray []remapped

func (s RemappedArray) Len() int          { return len(s) }
func (s RemappedArray) Range() []remapped { return s }

type S8Array []int8

func (s S8Array) Len() int      { return len(s) }
func (s S8Array) Range() []int8 { return s }

type StringArray []string

func (s StringArray) Len() int        { return len(s) }
func (s StringArray) Range() []string { return s }

////////////////////////////////////////////////////////////////////////////////
// CmdVoid
////////////////////////////////////////////////////////////////////////////////
type CmdVoid struct {
	binary.Generate
	InContext atom.ContextID
}

func (c *CmdVoid) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void(")
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoid) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdVoid) TypeID() atom.TypeID {
	return 0
}
func (c *CmdVoid) Flags() atom.Flags {
	return 0
}
func (CmdVoid) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidU8
////////////////////////////////////////////////////////////////////////////////
type CmdVoidU8 struct {
	binary.Generate
	InContext atom.ContextID
	A         uint8
}

func (c *CmdVoidU8) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_u8(",
		fmt.Sprintf("a:%v", c.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidU8) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdVoidU8) TypeID() atom.TypeID {
	return 1
}
func (c *CmdVoidU8) Flags() atom.Flags {
	return 0
}
func (CmdVoidU8) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidS8
////////////////////////////////////////////////////////////////////////////////
type CmdVoidS8 struct {
	binary.Generate
	InContext atom.ContextID
	A         int8
}

func (c *CmdVoidS8) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_s8(",
		fmt.Sprintf("a:%v", c.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidS8) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdVoidS8) TypeID() atom.TypeID {
	return 2
}
func (c *CmdVoidS8) Flags() atom.Flags {
	return 0
}
func (CmdVoidS8) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidU16
////////////////////////////////////////////////////////////////////////////////
type CmdVoidU16 struct {
	binary.Generate
	InContext atom.ContextID
	A         uint16
}

func (c *CmdVoidU16) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_u16(",
		fmt.Sprintf("a:%v", c.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidU16) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdVoidU16) TypeID() atom.TypeID {
	return 3
}
func (c *CmdVoidU16) Flags() atom.Flags {
	return 0
}
func (CmdVoidU16) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidS16
////////////////////////////////////////////////////////////////////////////////
type CmdVoidS16 struct {
	binary.Generate
	InContext atom.ContextID
	A         int16
}

func (c *CmdVoidS16) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_s16(",
		fmt.Sprintf("a:%v", c.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidS16) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdVoidS16) TypeID() atom.TypeID {
	return 4
}
func (c *CmdVoidS16) Flags() atom.Flags {
	return 0
}
func (CmdVoidS16) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidF32
////////////////////////////////////////////////////////////////////////////////
type CmdVoidF32 struct {
	binary.Generate
	InContext atom.ContextID
	A         float32
}

func (c *CmdVoidF32) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_f32(",
		fmt.Sprintf("a:%v", c.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidF32) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdVoidF32) TypeID() atom.TypeID {
	return 5
}
func (c *CmdVoidF32) Flags() atom.Flags {
	return 0
}
func (CmdVoidF32) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidU32
////////////////////////////////////////////////////////////////////////////////
type CmdVoidU32 struct {
	binary.Generate
	InContext atom.ContextID
	A         uint32
}

func (c *CmdVoidU32) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_u32(",
		fmt.Sprintf("a:%v", c.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidU32) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdVoidU32) TypeID() atom.TypeID {
	return 6
}
func (c *CmdVoidU32) Flags() atom.Flags {
	return 0
}
func (CmdVoidU32) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidS32
////////////////////////////////////////////////////////////////////////////////
type CmdVoidS32 struct {
	binary.Generate
	InContext atom.ContextID
	A         int32
}

func (c *CmdVoidS32) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_s32(",
		fmt.Sprintf("a:%v", c.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidS32) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdVoidS32) TypeID() atom.TypeID {
	return 7
}
func (c *CmdVoidS32) Flags() atom.Flags {
	return 0
}
func (CmdVoidS32) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidF64
////////////////////////////////////////////////////////////////////////////////
type CmdVoidF64 struct {
	binary.Generate
	InContext atom.ContextID
	A         float64
}

func (c *CmdVoidF64) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_f64(",
		fmt.Sprintf("a:%v", c.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidF64) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdVoidF64) TypeID() atom.TypeID {
	return 8
}
func (c *CmdVoidF64) Flags() atom.Flags {
	return 0
}
func (CmdVoidF64) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidU64
////////////////////////////////////////////////////////////////////////////////
type CmdVoidU64 struct {
	binary.Generate
	InContext atom.ContextID
	A         uint64
}

func (c *CmdVoidU64) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_u64(",
		fmt.Sprintf("a:%v", c.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidU64) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdVoidU64) TypeID() atom.TypeID {
	return 9
}
func (c *CmdVoidU64) Flags() atom.Flags {
	return 0
}
func (CmdVoidU64) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidS64
////////////////////////////////////////////////////////////////////////////////
type CmdVoidS64 struct {
	binary.Generate
	InContext atom.ContextID
	A         int64
}

func (c *CmdVoidS64) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_s64(",
		fmt.Sprintf("a:%v", c.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidS64) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdVoidS64) TypeID() atom.TypeID {
	return 10
}
func (c *CmdVoidS64) Flags() atom.Flags {
	return 0
}
func (CmdVoidS64) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidBool
////////////////////////////////////////////////////////////////////////////////
type CmdVoidBool struct {
	binary.Generate
	InContext atom.ContextID
	A         bool
}

func (c *CmdVoidBool) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_bool(",
		fmt.Sprintf("a:%v", c.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidBool) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdVoidBool) TypeID() atom.TypeID {
	return 11
}
func (c *CmdVoidBool) Flags() atom.Flags {
	return 0
}
func (CmdVoidBool) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidString
////////////////////////////////////////////////////////////////////////////////
type CmdVoidString struct {
	binary.Generate
	InContext atom.ContextID
	A         string
}

func (c *CmdVoidString) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_string(",
		fmt.Sprintf("a:%v", c.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidString) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdVoidString) TypeID() atom.TypeID {
	return 12
}
func (c *CmdVoidString) Flags() atom.Flags {
	return 0
}
func (CmdVoidString) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdVoid3Strings
////////////////////////////////////////////////////////////////////////////////
type CmdVoid3Strings struct {
	binary.Generate
	InContext atom.ContextID
	A         string
	B         string
	C         string
}

func (c *CmdVoid3Strings) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_3_strings(",
		fmt.Sprintf("a:%v", c.A),
		", ",
		fmt.Sprintf("b:%v", c.B),
		", ",
		fmt.Sprintf("c:%v", c.C),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoid3Strings) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdVoid3Strings) TypeID() atom.TypeID {
	return 13
}
func (c *CmdVoid3Strings) Flags() atom.Flags {
	return 0
}
func (CmdVoid3Strings) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdVoid3Arrays
////////////////////////////////////////////////////////////////////////////////
type CmdVoid3Arrays struct {
	binary.Generate
	InContext atom.ContextID
	A         S8Array
	B         StringArray
	C         BoolArray
}

func (c *CmdVoid3Arrays) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_3_arrays(",
		fmt.Sprintf("%v", c.A),
		", ",
		fmt.Sprintf("%v", c.B),
		", ",
		fmt.Sprintf("%v", c.C),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoid3Arrays) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdVoid3Arrays) TypeID() atom.TypeID {
	return 14
}
func (c *CmdVoid3Arrays) Flags() atom.Flags {
	return 0
}
func (CmdVoid3Arrays) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidArrayOfStrings
////////////////////////////////////////////////////////////////////////////////
type CmdVoidArrayOfStrings struct {
	binary.Generate
	InContext atom.ContextID
	A         StringArray
}

func (c *CmdVoidArrayOfStrings) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_array_of_strings(",
		fmt.Sprintf("%v", c.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidArrayOfStrings) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdVoidArrayOfStrings) TypeID() atom.TypeID {
	return 15
}
func (c *CmdVoidArrayOfStrings) Flags() atom.Flags {
	return 0
}
func (CmdVoidArrayOfStrings) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdU8
////////////////////////////////////////////////////////////////////////////////
type CmdU8 struct {
	binary.Generate
	InContext atom.ContextID
	Result    uint8
}

func (c *CmdU8) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_u8(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *CmdU8) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdU8) TypeID() atom.TypeID {
	return 16
}
func (c *CmdU8) Flags() atom.Flags {
	return 0
}
func (CmdU8) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdS8
////////////////////////////////////////////////////////////////////////////////
type CmdS8 struct {
	binary.Generate
	InContext atom.ContextID
	Result    int8
}

func (c *CmdS8) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_s8(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *CmdS8) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdS8) TypeID() atom.TypeID {
	return 17
}
func (c *CmdS8) Flags() atom.Flags {
	return 0
}
func (CmdS8) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdU16
////////////////////////////////////////////////////////////////////////////////
type CmdU16 struct {
	binary.Generate
	InContext atom.ContextID
	Result    uint16
}

func (c *CmdU16) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_u16(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *CmdU16) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdU16) TypeID() atom.TypeID {
	return 18
}
func (c *CmdU16) Flags() atom.Flags {
	return 0
}
func (CmdU16) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdS16
////////////////////////////////////////////////////////////////////////////////
type CmdS16 struct {
	binary.Generate
	InContext atom.ContextID
	Result    int16
}

func (c *CmdS16) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_s16(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *CmdS16) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdS16) TypeID() atom.TypeID {
	return 19
}
func (c *CmdS16) Flags() atom.Flags {
	return 0
}
func (CmdS16) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdF32
////////////////////////////////////////////////////////////////////////////////
type CmdF32 struct {
	binary.Generate
	InContext atom.ContextID
	Result    float32
}

func (c *CmdF32) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_f32(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *CmdF32) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdF32) TypeID() atom.TypeID {
	return 20
}
func (c *CmdF32) Flags() atom.Flags {
	return 0
}
func (CmdF32) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdU32
////////////////////////////////////////////////////////////////////////////////
type CmdU32 struct {
	binary.Generate
	InContext atom.ContextID
	Result    uint32
}

func (c *CmdU32) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_u32(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *CmdU32) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdU32) TypeID() atom.TypeID {
	return 21
}
func (c *CmdU32) Flags() atom.Flags {
	return 0
}
func (CmdU32) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdS32
////////////////////////////////////////////////////////////////////////////////
type CmdS32 struct {
	binary.Generate
	InContext atom.ContextID
	Result    int32
}

func (c *CmdS32) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_s32(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *CmdS32) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdS32) TypeID() atom.TypeID {
	return 22
}
func (c *CmdS32) Flags() atom.Flags {
	return 0
}
func (CmdS32) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdF64
////////////////////////////////////////////////////////////////////////////////
type CmdF64 struct {
	binary.Generate
	InContext atom.ContextID
	Result    float64
}

func (c *CmdF64) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_f64(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *CmdF64) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdF64) TypeID() atom.TypeID {
	return 23
}
func (c *CmdF64) Flags() atom.Flags {
	return 0
}
func (CmdF64) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdU64
////////////////////////////////////////////////////////////////////////////////
type CmdU64 struct {
	binary.Generate
	InContext atom.ContextID
	Result    uint64
}

func (c *CmdU64) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_u64(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *CmdU64) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdU64) TypeID() atom.TypeID {
	return 24
}
func (c *CmdU64) Flags() atom.Flags {
	return 0
}
func (CmdU64) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdS64
////////////////////////////////////////////////////////////////////////////////
type CmdS64 struct {
	binary.Generate
	InContext atom.ContextID
	Result    int64
}

func (c *CmdS64) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_s64(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *CmdS64) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdS64) TypeID() atom.TypeID {
	return 25
}
func (c *CmdS64) Flags() atom.Flags {
	return 0
}
func (CmdS64) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdBool
////////////////////////////////////////////////////////////////////////////////
type CmdBool struct {
	binary.Generate
	InContext atom.ContextID
	Result    bool
}

func (c *CmdBool) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_bool(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *CmdBool) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdBool) TypeID() atom.TypeID {
	return 26
}
func (c *CmdBool) Flags() atom.Flags {
	return 0
}
func (CmdBool) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdString
////////////////////////////////////////////////////////////////////////////////
type CmdString struct {
	binary.Generate
	InContext atom.ContextID
	Result    string
}

func (c *CmdString) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_string(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *CmdString) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdString) TypeID() atom.TypeID {
	return 27
}
func (c *CmdString) Flags() atom.Flags {
	return 0
}
func (CmdString) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdArrayOfFloat
////////////////////////////////////////////////////////////////////////////////
type CmdArrayOfFloat struct {
	binary.Generate
	InContext atom.ContextID
	Result    F32Array
}

func (c *CmdArrayOfFloat) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_array_of_float(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *CmdArrayOfFloat) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdArrayOfFloat) TypeID() atom.TypeID {
	return 28
}
func (c *CmdArrayOfFloat) Flags() atom.Flags {
	return 0
}
func (CmdArrayOfFloat) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdPointer
////////////////////////////////////////////////////////////////////////////////
type CmdPointer struct {
	binary.Generate
	InContext atom.ContextID
	Result    memory.Pointer
}

func (c *CmdPointer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_pointer(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *CmdPointer) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdPointer) TypeID() atom.TypeID {
	return 29
}
func (c *CmdPointer) Flags() atom.Flags {
	return 0
}
func (CmdPointer) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOutU8
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOutU8 struct {
	binary.Generate
	InContext atom.ContextID
	A         uint8
}

func (c *CmdVoidOutU8) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_out_u8(",
		fmt.Sprintf("a:%v", c.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidOutU8) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdVoidOutU8) TypeID() atom.TypeID {
	return 30
}
func (c *CmdVoidOutU8) Flags() atom.Flags {
	return 0
}
func (CmdVoidOutU8) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOutS8
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOutS8 struct {
	binary.Generate
	InContext atom.ContextID
	A         int8
}

func (c *CmdVoidOutS8) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_out_s8(",
		fmt.Sprintf("a:%v", c.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidOutS8) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdVoidOutS8) TypeID() atom.TypeID {
	return 31
}
func (c *CmdVoidOutS8) Flags() atom.Flags {
	return 0
}
func (CmdVoidOutS8) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOutU16
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOutU16 struct {
	binary.Generate
	InContext atom.ContextID
	A         uint16
}

func (c *CmdVoidOutU16) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_out_u16(",
		fmt.Sprintf("a:%v", c.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidOutU16) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdVoidOutU16) TypeID() atom.TypeID {
	return 32
}
func (c *CmdVoidOutU16) Flags() atom.Flags {
	return 0
}
func (CmdVoidOutU16) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOutS16
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOutS16 struct {
	binary.Generate
	InContext atom.ContextID
	A         int16
}

func (c *CmdVoidOutS16) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_out_s16(",
		fmt.Sprintf("a:%v", c.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidOutS16) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdVoidOutS16) TypeID() atom.TypeID {
	return 33
}
func (c *CmdVoidOutS16) Flags() atom.Flags {
	return 0
}
func (CmdVoidOutS16) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOutF32
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOutF32 struct {
	binary.Generate
	InContext atom.ContextID
	A         float32
}

func (c *CmdVoidOutF32) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_out_f32(",
		fmt.Sprintf("a:%v", c.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidOutF32) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdVoidOutF32) TypeID() atom.TypeID {
	return 34
}
func (c *CmdVoidOutF32) Flags() atom.Flags {
	return 0
}
func (CmdVoidOutF32) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOutU32
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOutU32 struct {
	binary.Generate
	InContext atom.ContextID
	A         uint32
}

func (c *CmdVoidOutU32) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_out_u32(",
		fmt.Sprintf("a:%v", c.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidOutU32) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdVoidOutU32) TypeID() atom.TypeID {
	return 35
}
func (c *CmdVoidOutU32) Flags() atom.Flags {
	return 0
}
func (CmdVoidOutU32) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOutS32
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOutS32 struct {
	binary.Generate
	InContext atom.ContextID
	A         int32
}

func (c *CmdVoidOutS32) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_out_s32(",
		fmt.Sprintf("a:%v", c.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidOutS32) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdVoidOutS32) TypeID() atom.TypeID {
	return 36
}
func (c *CmdVoidOutS32) Flags() atom.Flags {
	return 0
}
func (CmdVoidOutS32) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOutF64
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOutF64 struct {
	binary.Generate
	InContext atom.ContextID
	A         float64
}

func (c *CmdVoidOutF64) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_out_f64(",
		fmt.Sprintf("a:%v", c.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidOutF64) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdVoidOutF64) TypeID() atom.TypeID {
	return 37
}
func (c *CmdVoidOutF64) Flags() atom.Flags {
	return 0
}
func (CmdVoidOutF64) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOutU64
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOutU64 struct {
	binary.Generate
	InContext atom.ContextID
	A         uint64
}

func (c *CmdVoidOutU64) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_out_u64(",
		fmt.Sprintf("a:%v", c.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidOutU64) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdVoidOutU64) TypeID() atom.TypeID {
	return 38
}
func (c *CmdVoidOutU64) Flags() atom.Flags {
	return 0
}
func (CmdVoidOutU64) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOutS64
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOutS64 struct {
	binary.Generate
	InContext atom.ContextID
	A         int64
}

func (c *CmdVoidOutS64) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_out_s64(",
		fmt.Sprintf("a:%v", c.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidOutS64) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdVoidOutS64) TypeID() atom.TypeID {
	return 39
}
func (c *CmdVoidOutS64) Flags() atom.Flags {
	return 0
}
func (CmdVoidOutS64) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOutBool
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOutBool struct {
	binary.Generate
	InContext atom.ContextID
	A         bool
}

func (c *CmdVoidOutBool) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_out_bool(",
		fmt.Sprintf("a:%v", c.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidOutBool) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdVoidOutBool) TypeID() atom.TypeID {
	return 40
}
func (c *CmdVoidOutBool) Flags() atom.Flags {
	return 0
}
func (CmdVoidOutBool) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOutString
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOutString struct {
	binary.Generate
	InContext atom.ContextID
	A         string
}

func (c *CmdVoidOutString) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_out_string(",
		fmt.Sprintf("a:%v", c.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidOutString) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdVoidOutString) TypeID() atom.TypeID {
	return 41
}
func (c *CmdVoidOutString) Flags() atom.Flags {
	return 0
}
func (CmdVoidOutString) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOutFixedSizeBuffer
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOutFixedSizeBuffer struct {
	binary.Generate
	InContext atom.ContextID
	A         memory.Pointer
}

func (c *CmdVoidOutFixedSizeBuffer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_out_fixed_size_buffer(",
		fmt.Sprintf("0x%x", c.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidOutFixedSizeBuffer) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdVoidOutFixedSizeBuffer) TypeID() atom.TypeID {
	return 42
}
func (c *CmdVoidOutFixedSizeBuffer) Flags() atom.Flags {
	return 0
}
func (CmdVoidOutFixedSizeBuffer) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOut3Strings
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOut3Strings struct {
	binary.Generate
	InContext atom.ContextID
	A         string
	B         string
	C         string
}

func (c *CmdVoidOut3Strings) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_out_3_strings(",
		fmt.Sprintf("a:%v", c.A),
		", ",
		fmt.Sprintf("b:%v", c.B),
		", ",
		fmt.Sprintf("c:%v", c.C),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidOut3Strings) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdVoidOut3Strings) TypeID() atom.TypeID {
	return 43
}
func (c *CmdVoidOut3Strings) Flags() atom.Flags {
	return 0
}
func (CmdVoidOut3Strings) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdVoid3Remapped
////////////////////////////////////////////////////////////////////////////////
type CmdVoid3Remapped struct {
	binary.Generate
	InContext atom.ContextID
	A         remapped
	B         remapped
	C         remapped
}

func (c *CmdVoid3Remapped) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_3_remapped(",
		fmt.Sprintf("a:%v", c.A),
		", ",
		fmt.Sprintf("b:%v", c.B),
		", ",
		fmt.Sprintf("c:%v", c.C),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoid3Remapped) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdVoid3Remapped) TypeID() atom.TypeID {
	return 44
}
func (c *CmdVoid3Remapped) Flags() atom.Flags {
	return 0
}
func (CmdVoid3Remapped) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOut3Remapped
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOut3Remapped struct {
	binary.Generate
	InContext atom.ContextID
	A         remapped
	B         remapped
	C         remapped
}

func (c *CmdVoidOut3Remapped) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_out_3_remapped(",
		fmt.Sprintf("a:%v", c.A),
		", ",
		fmt.Sprintf("b:%v", c.B),
		", ",
		fmt.Sprintf("c:%v", c.C),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidOut3Remapped) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdVoidOut3Remapped) TypeID() atom.TypeID {
	return 45
}
func (c *CmdVoidOut3Remapped) Flags() atom.Flags {
	return 0
}
func (CmdVoidOut3Remapped) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOutArrayOfRemapped
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOutArrayOfRemapped struct {
	binary.Generate
	InContext atom.ContextID
	A         RemappedArray
}

func (c *CmdVoidOutArrayOfRemapped) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_out_array_of_remapped(",
		fmt.Sprintf("%v", c.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidOutArrayOfRemapped) ContextID() atom.ContextID {
	return c.InContext
}
func (c *CmdVoidOutArrayOfRemapped) TypeID() atom.TypeID {
	return 46
}
func (c *CmdVoidOutArrayOfRemapped) Flags() atom.Flags {
	return 0
}
func (CmdVoidOutArrayOfRemapped) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// Globals
////////////////////////////////////////////////////////////////////////////////
type Globals struct {
	binary.Generate
}

func (g *Globals) Init() {
}
func NewCmdVoid(
	contextID atom.ContextID,
) *CmdVoid {
	return &CmdVoid{
		InContext: contextID,
	}
}
func NewCmdVoidU8(
	contextID atom.ContextID,
	pA uint8,
) *CmdVoidU8 {
	return &CmdVoidU8{
		InContext: contextID,
		A:         pA}
}
func NewCmdVoidS8(
	contextID atom.ContextID,
	pA int8,
) *CmdVoidS8 {
	return &CmdVoidS8{
		InContext: contextID,
		A:         pA}
}
func NewCmdVoidU16(
	contextID atom.ContextID,
	pA uint16,
) *CmdVoidU16 {
	return &CmdVoidU16{
		InContext: contextID,
		A:         pA}
}
func NewCmdVoidS16(
	contextID atom.ContextID,
	pA int16,
) *CmdVoidS16 {
	return &CmdVoidS16{
		InContext: contextID,
		A:         pA}
}
func NewCmdVoidF32(
	contextID atom.ContextID,
	pA float32,
) *CmdVoidF32 {
	return &CmdVoidF32{
		InContext: contextID,
		A:         pA}
}
func NewCmdVoidU32(
	contextID atom.ContextID,
	pA uint32,
) *CmdVoidU32 {
	return &CmdVoidU32{
		InContext: contextID,
		A:         pA}
}
func NewCmdVoidS32(
	contextID atom.ContextID,
	pA int32,
) *CmdVoidS32 {
	return &CmdVoidS32{
		InContext: contextID,
		A:         pA}
}
func NewCmdVoidF64(
	contextID atom.ContextID,
	pA float64,
) *CmdVoidF64 {
	return &CmdVoidF64{
		InContext: contextID,
		A:         pA}
}
func NewCmdVoidU64(
	contextID atom.ContextID,
	pA uint64,
) *CmdVoidU64 {
	return &CmdVoidU64{
		InContext: contextID,
		A:         pA}
}
func NewCmdVoidS64(
	contextID atom.ContextID,
	pA int64,
) *CmdVoidS64 {
	return &CmdVoidS64{
		InContext: contextID,
		A:         pA}
}
func NewCmdVoidBool(
	contextID atom.ContextID,
	pA bool,
) *CmdVoidBool {
	return &CmdVoidBool{
		InContext: contextID,
		A:         pA}
}
func NewCmdVoidString(
	contextID atom.ContextID,
	pA string,
) *CmdVoidString {
	return &CmdVoidString{
		InContext: contextID,
		A:         pA}
}
func NewCmdVoid3Strings(
	contextID atom.ContextID,
	pA string,
	pB string,
	pC string,
) *CmdVoid3Strings {
	return &CmdVoid3Strings{
		InContext: contextID,
		A:         pA, B: pB, C: pC}
}
func NewCmdVoid3Arrays(
	contextID atom.ContextID,
	pA S8Array,
	pB StringArray,
	pC BoolArray,
) *CmdVoid3Arrays {
	return &CmdVoid3Arrays{
		InContext: contextID,
		A:         pA, B: pB, C: pC}
}
func NewCmdVoidArrayOfStrings(
	contextID atom.ContextID,
	pA StringArray,
) *CmdVoidArrayOfStrings {
	return &CmdVoidArrayOfStrings{
		InContext: contextID,
		A:         pA}
}
func NewCmdU8(
	contextID atom.ContextID,
	pResult uint8,
) *CmdU8 {
	return &CmdU8{
		InContext: contextID,
		Result:    pResult}
}
func NewCmdS8(
	contextID atom.ContextID,
	pResult int8,
) *CmdS8 {
	return &CmdS8{
		InContext: contextID,
		Result:    pResult}
}
func NewCmdU16(
	contextID atom.ContextID,
	pResult uint16,
) *CmdU16 {
	return &CmdU16{
		InContext: contextID,
		Result:    pResult}
}
func NewCmdS16(
	contextID atom.ContextID,
	pResult int16,
) *CmdS16 {
	return &CmdS16{
		InContext: contextID,
		Result:    pResult}
}
func NewCmdF32(
	contextID atom.ContextID,
	pResult float32,
) *CmdF32 {
	return &CmdF32{
		InContext: contextID,
		Result:    pResult}
}
func NewCmdU32(
	contextID atom.ContextID,
	pResult uint32,
) *CmdU32 {
	return &CmdU32{
		InContext: contextID,
		Result:    pResult}
}
func NewCmdS32(
	contextID atom.ContextID,
	pResult int32,
) *CmdS32 {
	return &CmdS32{
		InContext: contextID,
		Result:    pResult}
}
func NewCmdF64(
	contextID atom.ContextID,
	pResult float64,
) *CmdF64 {
	return &CmdF64{
		InContext: contextID,
		Result:    pResult}
}
func NewCmdU64(
	contextID atom.ContextID,
	pResult uint64,
) *CmdU64 {
	return &CmdU64{
		InContext: contextID,
		Result:    pResult}
}
func NewCmdS64(
	contextID atom.ContextID,
	pResult int64,
) *CmdS64 {
	return &CmdS64{
		InContext: contextID,
		Result:    pResult}
}
func NewCmdBool(
	contextID atom.ContextID,
	pResult bool,
) *CmdBool {
	return &CmdBool{
		InContext: contextID,
		Result:    pResult}
}
func NewCmdString(
	contextID atom.ContextID,
	pResult string,
) *CmdString {
	return &CmdString{
		InContext: contextID,
		Result:    pResult}
}
func NewCmdArrayOfFloat(
	contextID atom.ContextID,
	pResult F32Array,
) *CmdArrayOfFloat {
	return &CmdArrayOfFloat{
		InContext: contextID,
		Result:    pResult}
}
func NewCmdPointer(
	contextID atom.ContextID,
	pResult memory.Pointer,
) *CmdPointer {
	return &CmdPointer{
		InContext: contextID,
		Result:    pResult}
}
func NewCmdVoidOutU8(
	contextID atom.ContextID,
	pA uint8,
) *CmdVoidOutU8 {
	return &CmdVoidOutU8{
		InContext: contextID,
		A:         pA}
}
func NewCmdVoidOutS8(
	contextID atom.ContextID,
	pA int8,
) *CmdVoidOutS8 {
	return &CmdVoidOutS8{
		InContext: contextID,
		A:         pA}
}
func NewCmdVoidOutU16(
	contextID atom.ContextID,
	pA uint16,
) *CmdVoidOutU16 {
	return &CmdVoidOutU16{
		InContext: contextID,
		A:         pA}
}
func NewCmdVoidOutS16(
	contextID atom.ContextID,
	pA int16,
) *CmdVoidOutS16 {
	return &CmdVoidOutS16{
		InContext: contextID,
		A:         pA}
}
func NewCmdVoidOutF32(
	contextID atom.ContextID,
	pA float32,
) *CmdVoidOutF32 {
	return &CmdVoidOutF32{
		InContext: contextID,
		A:         pA}
}
func NewCmdVoidOutU32(
	contextID atom.ContextID,
	pA uint32,
) *CmdVoidOutU32 {
	return &CmdVoidOutU32{
		InContext: contextID,
		A:         pA}
}
func NewCmdVoidOutS32(
	contextID atom.ContextID,
	pA int32,
) *CmdVoidOutS32 {
	return &CmdVoidOutS32{
		InContext: contextID,
		A:         pA}
}
func NewCmdVoidOutF64(
	contextID atom.ContextID,
	pA float64,
) *CmdVoidOutF64 {
	return &CmdVoidOutF64{
		InContext: contextID,
		A:         pA}
}
func NewCmdVoidOutU64(
	contextID atom.ContextID,
	pA uint64,
) *CmdVoidOutU64 {
	return &CmdVoidOutU64{
		InContext: contextID,
		A:         pA}
}
func NewCmdVoidOutS64(
	contextID atom.ContextID,
	pA int64,
) *CmdVoidOutS64 {
	return &CmdVoidOutS64{
		InContext: contextID,
		A:         pA}
}
func NewCmdVoidOutBool(
	contextID atom.ContextID,
	pA bool,
) *CmdVoidOutBool {
	return &CmdVoidOutBool{
		InContext: contextID,
		A:         pA}
}
func NewCmdVoidOutString(
	contextID atom.ContextID,
	pA string,
) *CmdVoidOutString {
	return &CmdVoidOutString{
		InContext: contextID,
		A:         pA}
}
func NewCmdVoidOutFixedSizeBuffer(
	contextID atom.ContextID,
	pA memory.Pointer,
) *CmdVoidOutFixedSizeBuffer {
	return &CmdVoidOutFixedSizeBuffer{
		InContext: contextID,
		A:         pA}
}
func NewCmdVoidOut3Strings(
	contextID atom.ContextID,
	pA string,
	pB string,
	pC string,
) *CmdVoidOut3Strings {
	return &CmdVoidOut3Strings{
		InContext: contextID,
		A:         pA, B: pB, C: pC}
}
func NewCmdVoid3Remapped(
	contextID atom.ContextID,
	pA remapped,
	pB remapped,
	pC remapped,
) *CmdVoid3Remapped {
	return &CmdVoid3Remapped{
		InContext: contextID,
		A:         pA, B: pB, C: pC}
}
func NewCmdVoidOut3Remapped(
	contextID atom.ContextID,
	pA remapped,
	pB remapped,
	pC remapped,
) *CmdVoidOut3Remapped {
	return &CmdVoidOut3Remapped{
		InContext: contextID,
		A:         pA, B: pB, C: pC}
}
func NewCmdVoidOutArrayOfRemapped(
	contextID atom.ContextID,
	pA RemappedArray,
) *CmdVoidOutArrayOfRemapped {
	return &CmdVoidOutArrayOfRemapped{
		InContext: contextID,
		A:         pA}
}

////////////////////////////////////////////////////////////////////////////////
// API
////////////////////////////////////////////////////////////////////////////////
type api struct{}

func (api) Name() string {
	return "gfxapi_test"
}
func (api) ID() service.ApiId {
	return service.ApiId{ID: binary.NewID([]byte("gfxapi_test"))}
}
func API() gfxapi.API {
	return api{}
}
func init() {
	gfxapi.Register(API())
	atom.Register(atom.TypeInfo{
		Name: "CmdVoid",
		Docs: "[]",
		ID:   0,
		New:  func() atom.Atom { return &CmdVoid{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidU8",
		Docs: "[]",
		ID:   1,
		New:  func() atom.Atom { return &CmdVoidU8{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidS8",
		Docs: "[]",
		ID:   2,
		New:  func() atom.Atom { return &CmdVoidS8{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidU16",
		Docs: "[]",
		ID:   3,
		New:  func() atom.Atom { return &CmdVoidU16{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidS16",
		Docs: "[]",
		ID:   4,
		New:  func() atom.Atom { return &CmdVoidS16{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidF32",
		Docs: "[]",
		ID:   5,
		New:  func() atom.Atom { return &CmdVoidF32{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidU32",
		Docs: "[]",
		ID:   6,
		New:  func() atom.Atom { return &CmdVoidU32{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidS32",
		Docs: "[]",
		ID:   7,
		New:  func() atom.Atom { return &CmdVoidS32{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidF64",
		Docs: "[]",
		ID:   8,
		New:  func() atom.Atom { return &CmdVoidF64{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidU64",
		Docs: "[]",
		ID:   9,
		New:  func() atom.Atom { return &CmdVoidU64{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidS64",
		Docs: "[]",
		ID:   10,
		New:  func() atom.Atom { return &CmdVoidS64{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidBool",
		Docs: "[]",
		ID:   11,
		New:  func() atom.Atom { return &CmdVoidBool{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidString",
		Docs: "[]",
		ID:   12,
		New:  func() atom.Atom { return &CmdVoidString{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoid3Strings",
		Docs: "[]",
		ID:   13,
		New:  func() atom.Atom { return &CmdVoid3Strings{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoid3Arrays",
		Docs: "[]",
		ID:   14,
		New:  func() atom.Atom { return &CmdVoid3Arrays{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidArrayOfStrings",
		Docs: "[]",
		ID:   15,
		New:  func() atom.Atom { return &CmdVoidArrayOfStrings{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdU8",
		Docs: "[]",
		ID:   16,
		New:  func() atom.Atom { return &CmdU8{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdS8",
		Docs: "[]",
		ID:   17,
		New:  func() atom.Atom { return &CmdS8{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdU16",
		Docs: "[]",
		ID:   18,
		New:  func() atom.Atom { return &CmdU16{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdS16",
		Docs: "[]",
		ID:   19,
		New:  func() atom.Atom { return &CmdS16{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdF32",
		Docs: "[]",
		ID:   20,
		New:  func() atom.Atom { return &CmdF32{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdU32",
		Docs: "[]",
		ID:   21,
		New:  func() atom.Atom { return &CmdU32{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdS32",
		Docs: "[]",
		ID:   22,
		New:  func() atom.Atom { return &CmdS32{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdF64",
		Docs: "[]",
		ID:   23,
		New:  func() atom.Atom { return &CmdF64{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdU64",
		Docs: "[]",
		ID:   24,
		New:  func() atom.Atom { return &CmdU64{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdS64",
		Docs: "[]",
		ID:   25,
		New:  func() atom.Atom { return &CmdS64{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdBool",
		Docs: "[]",
		ID:   26,
		New:  func() atom.Atom { return &CmdBool{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdString",
		Docs: "[]",
		ID:   27,
		New:  func() atom.Atom { return &CmdString{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdArrayOfFloat",
		Docs: "[]",
		ID:   28,
		New:  func() atom.Atom { return &CmdArrayOfFloat{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdPointer",
		Docs: "[]",
		ID:   29,
		New:  func() atom.Atom { return &CmdPointer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidOutU8",
		Docs: "[]",
		ID:   30,
		New:  func() atom.Atom { return &CmdVoidOutU8{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidOutS8",
		Docs: "[]",
		ID:   31,
		New:  func() atom.Atom { return &CmdVoidOutS8{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidOutU16",
		Docs: "[]",
		ID:   32,
		New:  func() atom.Atom { return &CmdVoidOutU16{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidOutS16",
		Docs: "[]",
		ID:   33,
		New:  func() atom.Atom { return &CmdVoidOutS16{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidOutF32",
		Docs: "[]",
		ID:   34,
		New:  func() atom.Atom { return &CmdVoidOutF32{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidOutU32",
		Docs: "[]",
		ID:   35,
		New:  func() atom.Atom { return &CmdVoidOutU32{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidOutS32",
		Docs: "[]",
		ID:   36,
		New:  func() atom.Atom { return &CmdVoidOutS32{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidOutF64",
		Docs: "[]",
		ID:   37,
		New:  func() atom.Atom { return &CmdVoidOutF64{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidOutU64",
		Docs: "[]",
		ID:   38,
		New:  func() atom.Atom { return &CmdVoidOutU64{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidOutS64",
		Docs: "[]",
		ID:   39,
		New:  func() atom.Atom { return &CmdVoidOutS64{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidOutBool",
		Docs: "[]",
		ID:   40,
		New:  func() atom.Atom { return &CmdVoidOutBool{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidOutString",
		Docs: "[]",
		ID:   41,
		New:  func() atom.Atom { return &CmdVoidOutString{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidOutFixedSizeBuffer",
		Docs: "[]",
		ID:   42,
		New:  func() atom.Atom { return &CmdVoidOutFixedSizeBuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidOut3Strings",
		Docs: "[]",
		ID:   43,
		New:  func() atom.Atom { return &CmdVoidOut3Strings{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoid3Remapped",
		Docs: "[]",
		ID:   44,
		New:  func() atom.Atom { return &CmdVoid3Remapped{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidOut3Remapped",
		Docs: "[]",
		ID:   45,
		New:  func() atom.Atom { return &CmdVoidOut3Remapped{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidOutArrayOfRemapped",
		Docs: "[]",
		ID:   46,
		New:  func() atom.Atom { return &CmdVoidOutArrayOfRemapped{} },
	})
}
