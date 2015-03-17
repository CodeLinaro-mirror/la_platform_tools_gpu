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
)

type remapped uint32

func (c *remapped) Encode(e *binary.Encoder) error {
	x := uint32(*c)
	if err := e.Uint32(x); err != nil {
		return err
	}
	return nil
}
func (c *remapped) Decode(d *binary.Decoder) error {
	var x uint32
	if v, err := d.Uint32(); err == nil {
		x = v
	} else {
		return err
	}
	*c = remapped(x)
	return nil
}
func (c *remapped) Less(rhs remapped) bool  { return uint32(*c) < uint32(rhs) }
func (c *remapped) Equal(rhs remapped) bool { return uint32(*c) == uint32(rhs) }

type BoolArray []bool

func (s BoolArray) Len() int      { return len(s) }
func (s BoolArray) Range() []bool { return s }
func (s BoolArray) Encode(e *binary.Encoder) error {
	e.Uint32(uint32(len(s)))
	for _, v := range s {
		if err := e.Bool(v); err != nil {
			return err
		}
	}
	return nil
}
func (s *BoolArray) Decode(d *binary.Decoder) error {
	c, err := d.Uint32()
	if err != nil {
		return err
	}
	*s = make(BoolArray, c)
	for i := range *s {
		if v, err := d.Bool(); err == nil {
			(*s)[i] = v
		} else {
			return err
		}
	}
	return nil
}

type F32Array []float32

func (s F32Array) Len() int         { return len(s) }
func (s F32Array) Range() []float32 { return s }
func (s F32Array) Encode(e *binary.Encoder) error {
	e.Uint32(uint32(len(s)))
	for _, v := range s {
		if err := e.Float32(v); err != nil {
			return err
		}
	}
	return nil
}
func (s *F32Array) Decode(d *binary.Decoder) error {
	c, err := d.Uint32()
	if err != nil {
		return err
	}
	*s = make(F32Array, c)
	for i := range *s {
		if v, err := d.Float32(); err == nil {
			(*s)[i] = v
		} else {
			return err
		}
	}
	return nil
}

type RemappedArray []remapped

func (s RemappedArray) Len() int          { return len(s) }
func (s RemappedArray) Range() []remapped { return s }
func (s RemappedArray) Encode(e *binary.Encoder) error {
	e.Uint32(uint32(len(s)))
	for _, v := range s {
		if err := v.Encode(e); err != nil {
			return err
		}
	}
	return nil
}
func (s *RemappedArray) Decode(d *binary.Decoder) error {
	c, err := d.Uint32()
	if err != nil {
		return err
	}
	*s = make(RemappedArray, c)
	for i := range *s {
		if err := (*s)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}

type S8Array []int8

func (s S8Array) Len() int      { return len(s) }
func (s S8Array) Range() []int8 { return s }
func (s S8Array) Encode(e *binary.Encoder) error {
	e.Uint32(uint32(len(s)))
	for _, v := range s {
		if err := e.Int8(v); err != nil {
			return err
		}
	}
	return nil
}
func (s *S8Array) Decode(d *binary.Decoder) error {
	c, err := d.Uint32()
	if err != nil {
		return err
	}
	*s = make(S8Array, c)
	for i := range *s {
		if v, err := d.Int8(); err == nil {
			(*s)[i] = v
		} else {
			return err
		}
	}
	return nil
}

type StringArray []string

func (s StringArray) Len() int        { return len(s) }
func (s StringArray) Range() []string { return s }
func (s StringArray) Encode(e *binary.Encoder) error {
	e.Uint32(uint32(len(s)))
	for _, v := range s {
		if err := e.String(v); err != nil {
			return err
		}
	}
	return nil
}
func (s *StringArray) Decode(d *binary.Decoder) error {
	c, err := d.Uint32()
	if err != nil {
		return err
	}
	*s = make(StringArray, c)
	for i := range *s {
		if v, err := d.String(); err == nil {
			(*s)[i] = v
		} else {
			return err
		}
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdVoid
////////////////////////////////////////////////////////////////////////////////
type CmdVoid_In struct {
}
type CmdVoid_Out struct {
}
type CmdVoid struct {
	Context atom.ContextID
	In      CmdVoid_In
	Out     CmdVoid_Out
}

func (c *CmdVoid) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void(")
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoid) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdVoid) TypeID() atom.TypeID {
	return 0
}
func (c *CmdVoid) Flags() atom.Flags {
	return 0
}
func (c *CmdVoid) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	return nil
}
func (c *CmdVoid) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdVoidU8
////////////////////////////////////////////////////////////////////////////////
type CmdVoidU8_In struct {
	A uint8
}
type CmdVoidU8_Out struct {
}
type CmdVoidU8 struct {
	Context atom.ContextID
	In      CmdVoidU8_In
	Out     CmdVoidU8_Out
}

func (c *CmdVoidU8) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_u8(",
		fmt.Sprintf("a:%v", c.In.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidU8) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdVoidU8) TypeID() atom.TypeID {
	return 1
}
func (c *CmdVoidU8) Flags() atom.Flags {
	return 0
}
func (c *CmdVoidU8) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Uint8(c.In.A); err != nil {
		return err
	}
	return nil
}
func (c *CmdVoidU8) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Uint8(); err == nil {
		c.In.A = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdVoidS8
////////////////////////////////////////////////////////////////////////////////
type CmdVoidS8_In struct {
	A int8
}
type CmdVoidS8_Out struct {
}
type CmdVoidS8 struct {
	Context atom.ContextID
	In      CmdVoidS8_In
	Out     CmdVoidS8_Out
}

func (c *CmdVoidS8) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_s8(",
		fmt.Sprintf("a:%v", c.In.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidS8) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdVoidS8) TypeID() atom.TypeID {
	return 2
}
func (c *CmdVoidS8) Flags() atom.Flags {
	return 0
}
func (c *CmdVoidS8) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Int8(c.In.A); err != nil {
		return err
	}
	return nil
}
func (c *CmdVoidS8) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Int8(); err == nil {
		c.In.A = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdVoidU16
////////////////////////////////////////////////////////////////////////////////
type CmdVoidU16_In struct {
	A uint16
}
type CmdVoidU16_Out struct {
}
type CmdVoidU16 struct {
	Context atom.ContextID
	In      CmdVoidU16_In
	Out     CmdVoidU16_Out
}

func (c *CmdVoidU16) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_u16(",
		fmt.Sprintf("a:%v", c.In.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidU16) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdVoidU16) TypeID() atom.TypeID {
	return 3
}
func (c *CmdVoidU16) Flags() atom.Flags {
	return 0
}
func (c *CmdVoidU16) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Uint16(c.In.A); err != nil {
		return err
	}
	return nil
}
func (c *CmdVoidU16) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Uint16(); err == nil {
		c.In.A = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdVoidS16
////////////////////////////////////////////////////////////////////////////////
type CmdVoidS16_In struct {
	A int16
}
type CmdVoidS16_Out struct {
}
type CmdVoidS16 struct {
	Context atom.ContextID
	In      CmdVoidS16_In
	Out     CmdVoidS16_Out
}

func (c *CmdVoidS16) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_s16(",
		fmt.Sprintf("a:%v", c.In.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidS16) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdVoidS16) TypeID() atom.TypeID {
	return 4
}
func (c *CmdVoidS16) Flags() atom.Flags {
	return 0
}
func (c *CmdVoidS16) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Int16(c.In.A); err != nil {
		return err
	}
	return nil
}
func (c *CmdVoidS16) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Int16(); err == nil {
		c.In.A = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdVoidF32
////////////////////////////////////////////////////////////////////////////////
type CmdVoidF32_In struct {
	A float32
}
type CmdVoidF32_Out struct {
}
type CmdVoidF32 struct {
	Context atom.ContextID
	In      CmdVoidF32_In
	Out     CmdVoidF32_Out
}

func (c *CmdVoidF32) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_f32(",
		fmt.Sprintf("a:%v", c.In.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidF32) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdVoidF32) TypeID() atom.TypeID {
	return 5
}
func (c *CmdVoidF32) Flags() atom.Flags {
	return 0
}
func (c *CmdVoidF32) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Float32(c.In.A); err != nil {
		return err
	}
	return nil
}
func (c *CmdVoidF32) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Float32(); err == nil {
		c.In.A = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdVoidU32
////////////////////////////////////////////////////////////////////////////////
type CmdVoidU32_In struct {
	A uint32
}
type CmdVoidU32_Out struct {
}
type CmdVoidU32 struct {
	Context atom.ContextID
	In      CmdVoidU32_In
	Out     CmdVoidU32_Out
}

func (c *CmdVoidU32) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_u32(",
		fmt.Sprintf("a:%v", c.In.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidU32) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdVoidU32) TypeID() atom.TypeID {
	return 6
}
func (c *CmdVoidU32) Flags() atom.Flags {
	return 0
}
func (c *CmdVoidU32) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Uint32(c.In.A); err != nil {
		return err
	}
	return nil
}
func (c *CmdVoidU32) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Uint32(); err == nil {
		c.In.A = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdVoidS32
////////////////////////////////////////////////////////////////////////////////
type CmdVoidS32_In struct {
	A int32
}
type CmdVoidS32_Out struct {
}
type CmdVoidS32 struct {
	Context atom.ContextID
	In      CmdVoidS32_In
	Out     CmdVoidS32_Out
}

func (c *CmdVoidS32) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_s32(",
		fmt.Sprintf("a:%v", c.In.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidS32) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdVoidS32) TypeID() atom.TypeID {
	return 7
}
func (c *CmdVoidS32) Flags() atom.Flags {
	return 0
}
func (c *CmdVoidS32) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Int32(c.In.A); err != nil {
		return err
	}
	return nil
}
func (c *CmdVoidS32) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Int32(); err == nil {
		c.In.A = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdVoidF64
////////////////////////////////////////////////////////////////////////////////
type CmdVoidF64_In struct {
	A float64
}
type CmdVoidF64_Out struct {
}
type CmdVoidF64 struct {
	Context atom.ContextID
	In      CmdVoidF64_In
	Out     CmdVoidF64_Out
}

func (c *CmdVoidF64) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_f64(",
		fmt.Sprintf("a:%v", c.In.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidF64) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdVoidF64) TypeID() atom.TypeID {
	return 8
}
func (c *CmdVoidF64) Flags() atom.Flags {
	return 0
}
func (c *CmdVoidF64) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Float64(c.In.A); err != nil {
		return err
	}
	return nil
}
func (c *CmdVoidF64) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Float64(); err == nil {
		c.In.A = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdVoidU64
////////////////////////////////////////////////////////////////////////////////
type CmdVoidU64_In struct {
	A uint64
}
type CmdVoidU64_Out struct {
}
type CmdVoidU64 struct {
	Context atom.ContextID
	In      CmdVoidU64_In
	Out     CmdVoidU64_Out
}

func (c *CmdVoidU64) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_u64(",
		fmt.Sprintf("a:%v", c.In.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidU64) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdVoidU64) TypeID() atom.TypeID {
	return 9
}
func (c *CmdVoidU64) Flags() atom.Flags {
	return 0
}
func (c *CmdVoidU64) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Uint64(c.In.A); err != nil {
		return err
	}
	return nil
}
func (c *CmdVoidU64) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Uint64(); err == nil {
		c.In.A = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdVoidS64
////////////////////////////////////////////////////////////////////////////////
type CmdVoidS64_In struct {
	A int64
}
type CmdVoidS64_Out struct {
}
type CmdVoidS64 struct {
	Context atom.ContextID
	In      CmdVoidS64_In
	Out     CmdVoidS64_Out
}

func (c *CmdVoidS64) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_s64(",
		fmt.Sprintf("a:%v", c.In.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidS64) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdVoidS64) TypeID() atom.TypeID {
	return 10
}
func (c *CmdVoidS64) Flags() atom.Flags {
	return 0
}
func (c *CmdVoidS64) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Int64(c.In.A); err != nil {
		return err
	}
	return nil
}
func (c *CmdVoidS64) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Int64(); err == nil {
		c.In.A = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdVoidBool
////////////////////////////////////////////////////////////////////////////////
type CmdVoidBool_In struct {
	A bool
}
type CmdVoidBool_Out struct {
}
type CmdVoidBool struct {
	Context atom.ContextID
	In      CmdVoidBool_In
	Out     CmdVoidBool_Out
}

func (c *CmdVoidBool) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_bool(",
		fmt.Sprintf("a:%v", c.In.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidBool) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdVoidBool) TypeID() atom.TypeID {
	return 11
}
func (c *CmdVoidBool) Flags() atom.Flags {
	return 0
}
func (c *CmdVoidBool) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Bool(c.In.A); err != nil {
		return err
	}
	return nil
}
func (c *CmdVoidBool) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Bool(); err == nil {
		c.In.A = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdVoidString
////////////////////////////////////////////////////////////////////////////////
type CmdVoidString_In struct {
	A string
}
type CmdVoidString_Out struct {
}
type CmdVoidString struct {
	Context atom.ContextID
	In      CmdVoidString_In
	Out     CmdVoidString_Out
}

func (c *CmdVoidString) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_string(",
		fmt.Sprintf("a:%v", c.In.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidString) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdVoidString) TypeID() atom.TypeID {
	return 12
}
func (c *CmdVoidString) Flags() atom.Flags {
	return 0
}
func (c *CmdVoidString) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.String(c.In.A); err != nil {
		return err
	}
	return nil
}
func (c *CmdVoidString) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.String(); err == nil {
		c.In.A = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdVoid3Strings
////////////////////////////////////////////////////////////////////////////////
type CmdVoid3Strings_In struct {
	A string
	B string
	C string
}
type CmdVoid3Strings_Out struct {
}
type CmdVoid3Strings struct {
	Context atom.ContextID
	In      CmdVoid3Strings_In
	Out     CmdVoid3Strings_Out
}

func (c *CmdVoid3Strings) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_3_strings(",
		fmt.Sprintf("a:%v", c.In.A),
		", ",
		fmt.Sprintf("b:%v", c.In.B),
		", ",
		fmt.Sprintf("c:%v", c.In.C),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoid3Strings) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdVoid3Strings) TypeID() atom.TypeID {
	return 13
}
func (c *CmdVoid3Strings) Flags() atom.Flags {
	return 0
}
func (c *CmdVoid3Strings) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.String(c.In.A); err != nil {
		return err
	}
	if err := e.String(c.In.B); err != nil {
		return err
	}
	if err := e.String(c.In.C); err != nil {
		return err
	}
	return nil
}
func (c *CmdVoid3Strings) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.String(); err == nil {
		c.In.A = v
	} else {
		return err
	}
	if v, err := d.String(); err == nil {
		c.In.B = v
	} else {
		return err
	}
	if v, err := d.String(); err == nil {
		c.In.C = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdVoid3Arrays
////////////////////////////////////////////////////////////////////////////////
type CmdVoid3Arrays_In struct {
	A S8Array
	B StringArray
	C BoolArray
}
type CmdVoid3Arrays_Out struct {
}
type CmdVoid3Arrays struct {
	Context atom.ContextID
	In      CmdVoid3Arrays_In
	Out     CmdVoid3Arrays_Out
}

func (c *CmdVoid3Arrays) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_3_arrays(",
		fmt.Sprintf("%v", c.In.A),
		", ",
		fmt.Sprintf("%v", c.In.B),
		", ",
		fmt.Sprintf("%v", c.In.C),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoid3Arrays) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdVoid3Arrays) TypeID() atom.TypeID {
	return 14
}
func (c *CmdVoid3Arrays) Flags() atom.Flags {
	return 0
}
func (c *CmdVoid3Arrays) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := c.In.A.Encode(e); err != nil {
		return err
	}
	if err := c.In.B.Encode(e); err != nil {
		return err
	}
	if err := c.In.C.Encode(e); err != nil {
		return err
	}
	return nil
}
func (c *CmdVoid3Arrays) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if err := c.In.A.Decode(d); err != nil {
		return err
	}
	if err := c.In.B.Decode(d); err != nil {
		return err
	}
	if err := c.In.C.Decode(d); err != nil {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdVoidArrayOfStrings
////////////////////////////////////////////////////////////////////////////////
type CmdVoidArrayOfStrings_In struct {
	A StringArray
}
type CmdVoidArrayOfStrings_Out struct {
}
type CmdVoidArrayOfStrings struct {
	Context atom.ContextID
	In      CmdVoidArrayOfStrings_In
	Out     CmdVoidArrayOfStrings_Out
}

func (c *CmdVoidArrayOfStrings) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_array_of_strings(",
		fmt.Sprintf("%v", c.In.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidArrayOfStrings) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdVoidArrayOfStrings) TypeID() atom.TypeID {
	return 15
}
func (c *CmdVoidArrayOfStrings) Flags() atom.Flags {
	return 0
}
func (c *CmdVoidArrayOfStrings) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := c.In.A.Encode(e); err != nil {
		return err
	}
	return nil
}
func (c *CmdVoidArrayOfStrings) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if err := c.In.A.Decode(d); err != nil {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdU8
////////////////////////////////////////////////////////////////////////////////
type CmdU8_In struct {
}
type CmdU8_Out struct {
	Result uint8
}
type CmdU8 struct {
	Context atom.ContextID
	In      CmdU8_In
	Out     CmdU8_Out
}

func (c *CmdU8) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_u8(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *CmdU8) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdU8) TypeID() atom.TypeID {
	return 16
}
func (c *CmdU8) Flags() atom.Flags {
	return 0
}
func (c *CmdU8) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Uint8(c.Out.Result); err != nil {
		return err
	}
	return nil
}
func (c *CmdU8) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Uint8(); err == nil {
		c.Out.Result = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdS8
////////////////////////////////////////////////////////////////////////////////
type CmdS8_In struct {
}
type CmdS8_Out struct {
	Result int8
}
type CmdS8 struct {
	Context atom.ContextID
	In      CmdS8_In
	Out     CmdS8_Out
}

func (c *CmdS8) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_s8(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *CmdS8) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdS8) TypeID() atom.TypeID {
	return 17
}
func (c *CmdS8) Flags() atom.Flags {
	return 0
}
func (c *CmdS8) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Int8(c.Out.Result); err != nil {
		return err
	}
	return nil
}
func (c *CmdS8) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Int8(); err == nil {
		c.Out.Result = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdU16
////////////////////////////////////////////////////////////////////////////////
type CmdU16_In struct {
}
type CmdU16_Out struct {
	Result uint16
}
type CmdU16 struct {
	Context atom.ContextID
	In      CmdU16_In
	Out     CmdU16_Out
}

func (c *CmdU16) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_u16(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *CmdU16) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdU16) TypeID() atom.TypeID {
	return 18
}
func (c *CmdU16) Flags() atom.Flags {
	return 0
}
func (c *CmdU16) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Uint16(c.Out.Result); err != nil {
		return err
	}
	return nil
}
func (c *CmdU16) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Uint16(); err == nil {
		c.Out.Result = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdS16
////////////////////////////////////////////////////////////////////////////////
type CmdS16_In struct {
}
type CmdS16_Out struct {
	Result int16
}
type CmdS16 struct {
	Context atom.ContextID
	In      CmdS16_In
	Out     CmdS16_Out
}

func (c *CmdS16) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_s16(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *CmdS16) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdS16) TypeID() atom.TypeID {
	return 19
}
func (c *CmdS16) Flags() atom.Flags {
	return 0
}
func (c *CmdS16) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Int16(c.Out.Result); err != nil {
		return err
	}
	return nil
}
func (c *CmdS16) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Int16(); err == nil {
		c.Out.Result = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdF32
////////////////////////////////////////////////////////////////////////////////
type CmdF32_In struct {
}
type CmdF32_Out struct {
	Result float32
}
type CmdF32 struct {
	Context atom.ContextID
	In      CmdF32_In
	Out     CmdF32_Out
}

func (c *CmdF32) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_f32(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *CmdF32) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdF32) TypeID() atom.TypeID {
	return 20
}
func (c *CmdF32) Flags() atom.Flags {
	return 0
}
func (c *CmdF32) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Float32(c.Out.Result); err != nil {
		return err
	}
	return nil
}
func (c *CmdF32) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Float32(); err == nil {
		c.Out.Result = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdU32
////////////////////////////////////////////////////////////////////////////////
type CmdU32_In struct {
}
type CmdU32_Out struct {
	Result uint32
}
type CmdU32 struct {
	Context atom.ContextID
	In      CmdU32_In
	Out     CmdU32_Out
}

func (c *CmdU32) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_u32(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *CmdU32) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdU32) TypeID() atom.TypeID {
	return 21
}
func (c *CmdU32) Flags() atom.Flags {
	return 0
}
func (c *CmdU32) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Uint32(c.Out.Result); err != nil {
		return err
	}
	return nil
}
func (c *CmdU32) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Uint32(); err == nil {
		c.Out.Result = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdS32
////////////////////////////////////////////////////////////////////////////////
type CmdS32_In struct {
}
type CmdS32_Out struct {
	Result int32
}
type CmdS32 struct {
	Context atom.ContextID
	In      CmdS32_In
	Out     CmdS32_Out
}

func (c *CmdS32) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_s32(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *CmdS32) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdS32) TypeID() atom.TypeID {
	return 22
}
func (c *CmdS32) Flags() atom.Flags {
	return 0
}
func (c *CmdS32) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Int32(c.Out.Result); err != nil {
		return err
	}
	return nil
}
func (c *CmdS32) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Int32(); err == nil {
		c.Out.Result = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdF64
////////////////////////////////////////////////////////////////////////////////
type CmdF64_In struct {
}
type CmdF64_Out struct {
	Result float64
}
type CmdF64 struct {
	Context atom.ContextID
	In      CmdF64_In
	Out     CmdF64_Out
}

func (c *CmdF64) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_f64(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *CmdF64) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdF64) TypeID() atom.TypeID {
	return 23
}
func (c *CmdF64) Flags() atom.Flags {
	return 0
}
func (c *CmdF64) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Float64(c.Out.Result); err != nil {
		return err
	}
	return nil
}
func (c *CmdF64) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Float64(); err == nil {
		c.Out.Result = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdU64
////////////////////////////////////////////////////////////////////////////////
type CmdU64_In struct {
}
type CmdU64_Out struct {
	Result uint64
}
type CmdU64 struct {
	Context atom.ContextID
	In      CmdU64_In
	Out     CmdU64_Out
}

func (c *CmdU64) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_u64(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *CmdU64) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdU64) TypeID() atom.TypeID {
	return 24
}
func (c *CmdU64) Flags() atom.Flags {
	return 0
}
func (c *CmdU64) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Uint64(c.Out.Result); err != nil {
		return err
	}
	return nil
}
func (c *CmdU64) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Uint64(); err == nil {
		c.Out.Result = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdS64
////////////////////////////////////////////////////////////////////////////////
type CmdS64_In struct {
}
type CmdS64_Out struct {
	Result int64
}
type CmdS64 struct {
	Context atom.ContextID
	In      CmdS64_In
	Out     CmdS64_Out
}

func (c *CmdS64) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_s64(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *CmdS64) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdS64) TypeID() atom.TypeID {
	return 25
}
func (c *CmdS64) Flags() atom.Flags {
	return 0
}
func (c *CmdS64) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Int64(c.Out.Result); err != nil {
		return err
	}
	return nil
}
func (c *CmdS64) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Int64(); err == nil {
		c.Out.Result = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdBool
////////////////////////////////////////////////////////////////////////////////
type CmdBool_In struct {
}
type CmdBool_Out struct {
	Result bool
}
type CmdBool struct {
	Context atom.ContextID
	In      CmdBool_In
	Out     CmdBool_Out
}

func (c *CmdBool) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_bool(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *CmdBool) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdBool) TypeID() atom.TypeID {
	return 26
}
func (c *CmdBool) Flags() atom.Flags {
	return 0
}
func (c *CmdBool) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Bool(c.Out.Result); err != nil {
		return err
	}
	return nil
}
func (c *CmdBool) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Bool(); err == nil {
		c.Out.Result = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdString
////////////////////////////////////////////////////////////////////////////////
type CmdString_In struct {
}
type CmdString_Out struct {
	Result string
}
type CmdString struct {
	Context atom.ContextID
	In      CmdString_In
	Out     CmdString_Out
}

func (c *CmdString) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_string(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *CmdString) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdString) TypeID() atom.TypeID {
	return 27
}
func (c *CmdString) Flags() atom.Flags {
	return 0
}
func (c *CmdString) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.String(c.Out.Result); err != nil {
		return err
	}
	return nil
}
func (c *CmdString) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.String(); err == nil {
		c.Out.Result = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdArrayOfFloat
////////////////////////////////////////////////////////////////////////////////
type CmdArrayOfFloat_In struct {
}
type CmdArrayOfFloat_Out struct {
	Result F32Array
}
type CmdArrayOfFloat struct {
	Context atom.ContextID
	In      CmdArrayOfFloat_In
	Out     CmdArrayOfFloat_Out
}

func (c *CmdArrayOfFloat) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_array_of_float(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *CmdArrayOfFloat) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdArrayOfFloat) TypeID() atom.TypeID {
	return 28
}
func (c *CmdArrayOfFloat) Flags() atom.Flags {
	return 0
}
func (c *CmdArrayOfFloat) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := c.Out.Result.Encode(e); err != nil {
		return err
	}
	return nil
}
func (c *CmdArrayOfFloat) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if err := c.Out.Result.Decode(d); err != nil {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdPointer
////////////////////////////////////////////////////////////////////////////////
type CmdPointer_In struct {
}
type CmdPointer_Out struct {
	Result memory.Pointer
}
type CmdPointer struct {
	Context atom.ContextID
	In      CmdPointer_In
	Out     CmdPointer_Out
}

func (c *CmdPointer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_pointer(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *CmdPointer) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdPointer) TypeID() atom.TypeID {
	return 29
}
func (c *CmdPointer) Flags() atom.Flags {
	return 0
}
func (c *CmdPointer) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Uint64(uint64(c.Out.Result)); err != nil {
		return err
	}
	return nil
}
func (c *CmdPointer) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Uint64(); err == nil {
		c.Out.Result = memory.Pointer(v)
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOutU8
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOutU8_In struct {
}
type CmdVoidOutU8_Out struct {
	A uint8
}
type CmdVoidOutU8 struct {
	Context atom.ContextID
	In      CmdVoidOutU8_In
	Out     CmdVoidOutU8_Out
}

func (c *CmdVoidOutU8) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_out_u8(",
		fmt.Sprintf("a:%v", c.Out.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidOutU8) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdVoidOutU8) TypeID() atom.TypeID {
	return 30
}
func (c *CmdVoidOutU8) Flags() atom.Flags {
	return 0
}
func (c *CmdVoidOutU8) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Uint8(c.Out.A); err != nil {
		return err
	}
	return nil
}
func (c *CmdVoidOutU8) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Uint8(); err == nil {
		c.Out.A = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOutS8
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOutS8_In struct {
}
type CmdVoidOutS8_Out struct {
	A int8
}
type CmdVoidOutS8 struct {
	Context atom.ContextID
	In      CmdVoidOutS8_In
	Out     CmdVoidOutS8_Out
}

func (c *CmdVoidOutS8) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_out_s8(",
		fmt.Sprintf("a:%v", c.Out.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidOutS8) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdVoidOutS8) TypeID() atom.TypeID {
	return 31
}
func (c *CmdVoidOutS8) Flags() atom.Flags {
	return 0
}
func (c *CmdVoidOutS8) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Int8(c.Out.A); err != nil {
		return err
	}
	return nil
}
func (c *CmdVoidOutS8) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Int8(); err == nil {
		c.Out.A = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOutU16
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOutU16_In struct {
}
type CmdVoidOutU16_Out struct {
	A uint16
}
type CmdVoidOutU16 struct {
	Context atom.ContextID
	In      CmdVoidOutU16_In
	Out     CmdVoidOutU16_Out
}

func (c *CmdVoidOutU16) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_out_u16(",
		fmt.Sprintf("a:%v", c.Out.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidOutU16) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdVoidOutU16) TypeID() atom.TypeID {
	return 32
}
func (c *CmdVoidOutU16) Flags() atom.Flags {
	return 0
}
func (c *CmdVoidOutU16) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Uint16(c.Out.A); err != nil {
		return err
	}
	return nil
}
func (c *CmdVoidOutU16) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Uint16(); err == nil {
		c.Out.A = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOutS16
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOutS16_In struct {
}
type CmdVoidOutS16_Out struct {
	A int16
}
type CmdVoidOutS16 struct {
	Context atom.ContextID
	In      CmdVoidOutS16_In
	Out     CmdVoidOutS16_Out
}

func (c *CmdVoidOutS16) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_out_s16(",
		fmt.Sprintf("a:%v", c.Out.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidOutS16) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdVoidOutS16) TypeID() atom.TypeID {
	return 33
}
func (c *CmdVoidOutS16) Flags() atom.Flags {
	return 0
}
func (c *CmdVoidOutS16) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Int16(c.Out.A); err != nil {
		return err
	}
	return nil
}
func (c *CmdVoidOutS16) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Int16(); err == nil {
		c.Out.A = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOutF32
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOutF32_In struct {
}
type CmdVoidOutF32_Out struct {
	A float32
}
type CmdVoidOutF32 struct {
	Context atom.ContextID
	In      CmdVoidOutF32_In
	Out     CmdVoidOutF32_Out
}

func (c *CmdVoidOutF32) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_out_f32(",
		fmt.Sprintf("a:%v", c.Out.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidOutF32) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdVoidOutF32) TypeID() atom.TypeID {
	return 34
}
func (c *CmdVoidOutF32) Flags() atom.Flags {
	return 0
}
func (c *CmdVoidOutF32) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Float32(c.Out.A); err != nil {
		return err
	}
	return nil
}
func (c *CmdVoidOutF32) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Float32(); err == nil {
		c.Out.A = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOutU32
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOutU32_In struct {
}
type CmdVoidOutU32_Out struct {
	A uint32
}
type CmdVoidOutU32 struct {
	Context atom.ContextID
	In      CmdVoidOutU32_In
	Out     CmdVoidOutU32_Out
}

func (c *CmdVoidOutU32) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_out_u32(",
		fmt.Sprintf("a:%v", c.Out.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidOutU32) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdVoidOutU32) TypeID() atom.TypeID {
	return 35
}
func (c *CmdVoidOutU32) Flags() atom.Flags {
	return 0
}
func (c *CmdVoidOutU32) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Uint32(c.Out.A); err != nil {
		return err
	}
	return nil
}
func (c *CmdVoidOutU32) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Uint32(); err == nil {
		c.Out.A = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOutS32
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOutS32_In struct {
}
type CmdVoidOutS32_Out struct {
	A int32
}
type CmdVoidOutS32 struct {
	Context atom.ContextID
	In      CmdVoidOutS32_In
	Out     CmdVoidOutS32_Out
}

func (c *CmdVoidOutS32) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_out_s32(",
		fmt.Sprintf("a:%v", c.Out.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidOutS32) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdVoidOutS32) TypeID() atom.TypeID {
	return 36
}
func (c *CmdVoidOutS32) Flags() atom.Flags {
	return 0
}
func (c *CmdVoidOutS32) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Int32(c.Out.A); err != nil {
		return err
	}
	return nil
}
func (c *CmdVoidOutS32) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Int32(); err == nil {
		c.Out.A = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOutF64
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOutF64_In struct {
}
type CmdVoidOutF64_Out struct {
	A float64
}
type CmdVoidOutF64 struct {
	Context atom.ContextID
	In      CmdVoidOutF64_In
	Out     CmdVoidOutF64_Out
}

func (c *CmdVoidOutF64) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_out_f64(",
		fmt.Sprintf("a:%v", c.Out.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidOutF64) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdVoidOutF64) TypeID() atom.TypeID {
	return 37
}
func (c *CmdVoidOutF64) Flags() atom.Flags {
	return 0
}
func (c *CmdVoidOutF64) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Float64(c.Out.A); err != nil {
		return err
	}
	return nil
}
func (c *CmdVoidOutF64) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Float64(); err == nil {
		c.Out.A = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOutU64
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOutU64_In struct {
}
type CmdVoidOutU64_Out struct {
	A uint64
}
type CmdVoidOutU64 struct {
	Context atom.ContextID
	In      CmdVoidOutU64_In
	Out     CmdVoidOutU64_Out
}

func (c *CmdVoidOutU64) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_out_u64(",
		fmt.Sprintf("a:%v", c.Out.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidOutU64) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdVoidOutU64) TypeID() atom.TypeID {
	return 38
}
func (c *CmdVoidOutU64) Flags() atom.Flags {
	return 0
}
func (c *CmdVoidOutU64) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Uint64(c.Out.A); err != nil {
		return err
	}
	return nil
}
func (c *CmdVoidOutU64) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Uint64(); err == nil {
		c.Out.A = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOutS64
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOutS64_In struct {
}
type CmdVoidOutS64_Out struct {
	A int64
}
type CmdVoidOutS64 struct {
	Context atom.ContextID
	In      CmdVoidOutS64_In
	Out     CmdVoidOutS64_Out
}

func (c *CmdVoidOutS64) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_out_s64(",
		fmt.Sprintf("a:%v", c.Out.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidOutS64) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdVoidOutS64) TypeID() atom.TypeID {
	return 39
}
func (c *CmdVoidOutS64) Flags() atom.Flags {
	return 0
}
func (c *CmdVoidOutS64) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Int64(c.Out.A); err != nil {
		return err
	}
	return nil
}
func (c *CmdVoidOutS64) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Int64(); err == nil {
		c.Out.A = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOutBool
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOutBool_In struct {
}
type CmdVoidOutBool_Out struct {
	A bool
}
type CmdVoidOutBool struct {
	Context atom.ContextID
	In      CmdVoidOutBool_In
	Out     CmdVoidOutBool_Out
}

func (c *CmdVoidOutBool) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_out_bool(",
		fmt.Sprintf("a:%v", c.Out.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidOutBool) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdVoidOutBool) TypeID() atom.TypeID {
	return 40
}
func (c *CmdVoidOutBool) Flags() atom.Flags {
	return 0
}
func (c *CmdVoidOutBool) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Bool(c.Out.A); err != nil {
		return err
	}
	return nil
}
func (c *CmdVoidOutBool) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Bool(); err == nil {
		c.Out.A = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOutString
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOutString_In struct {
}
type CmdVoidOutString_Out struct {
	A string
}
type CmdVoidOutString struct {
	Context atom.ContextID
	In      CmdVoidOutString_In
	Out     CmdVoidOutString_Out
}

func (c *CmdVoidOutString) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_out_string(",
		fmt.Sprintf("a:%v", c.Out.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidOutString) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdVoidOutString) TypeID() atom.TypeID {
	return 41
}
func (c *CmdVoidOutString) Flags() atom.Flags {
	return 0
}
func (c *CmdVoidOutString) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.String(c.Out.A); err != nil {
		return err
	}
	return nil
}
func (c *CmdVoidOutString) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.String(); err == nil {
		c.Out.A = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOutFixedSizeBuffer
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOutFixedSizeBuffer_In struct {
}
type CmdVoidOutFixedSizeBuffer_Out struct {
	A memory.Pointer
}
type CmdVoidOutFixedSizeBuffer struct {
	Context atom.ContextID
	In      CmdVoidOutFixedSizeBuffer_In
	Out     CmdVoidOutFixedSizeBuffer_Out
}

func (c *CmdVoidOutFixedSizeBuffer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_out_fixed_size_buffer(",
		fmt.Sprintf("0x%x", c.Out.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidOutFixedSizeBuffer) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdVoidOutFixedSizeBuffer) TypeID() atom.TypeID {
	return 42
}
func (c *CmdVoidOutFixedSizeBuffer) Flags() atom.Flags {
	return 0
}
func (c *CmdVoidOutFixedSizeBuffer) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.Uint64(uint64(c.Out.A)); err != nil {
		return err
	}
	return nil
}
func (c *CmdVoidOutFixedSizeBuffer) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.Uint64(); err == nil {
		c.Out.A = memory.Pointer(v)
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOut3Strings
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOut3Strings_In struct {
}
type CmdVoidOut3Strings_Out struct {
	A string
	B string
	C string
}
type CmdVoidOut3Strings struct {
	Context atom.ContextID
	In      CmdVoidOut3Strings_In
	Out     CmdVoidOut3Strings_Out
}

func (c *CmdVoidOut3Strings) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_out_3_strings(",
		fmt.Sprintf("a:%v", c.Out.A),
		", ",
		fmt.Sprintf("b:%v", c.Out.B),
		", ",
		fmt.Sprintf("c:%v", c.Out.C),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidOut3Strings) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdVoidOut3Strings) TypeID() atom.TypeID {
	return 43
}
func (c *CmdVoidOut3Strings) Flags() atom.Flags {
	return 0
}
func (c *CmdVoidOut3Strings) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := e.String(c.Out.A); err != nil {
		return err
	}
	if err := e.String(c.Out.B); err != nil {
		return err
	}
	if err := e.String(c.Out.C); err != nil {
		return err
	}
	return nil
}
func (c *CmdVoidOut3Strings) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if v, err := d.String(); err == nil {
		c.Out.A = v
	} else {
		return err
	}
	if v, err := d.String(); err == nil {
		c.Out.B = v
	} else {
		return err
	}
	if v, err := d.String(); err == nil {
		c.Out.C = v
	} else {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdVoid3Remapped
////////////////////////////////////////////////////////////////////////////////
type CmdVoid3Remapped_In struct {
	A remapped
	B remapped
	C remapped
}
type CmdVoid3Remapped_Out struct {
}
type CmdVoid3Remapped struct {
	Context atom.ContextID
	In      CmdVoid3Remapped_In
	Out     CmdVoid3Remapped_Out
}

func (c *CmdVoid3Remapped) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_3_remapped(",
		fmt.Sprintf("a:%v", c.In.A),
		", ",
		fmt.Sprintf("b:%v", c.In.B),
		", ",
		fmt.Sprintf("c:%v", c.In.C),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoid3Remapped) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdVoid3Remapped) TypeID() atom.TypeID {
	return 44
}
func (c *CmdVoid3Remapped) Flags() atom.Flags {
	return 0
}
func (c *CmdVoid3Remapped) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := c.In.A.Encode(e); err != nil {
		return err
	}
	if err := c.In.B.Encode(e); err != nil {
		return err
	}
	if err := c.In.C.Encode(e); err != nil {
		return err
	}
	return nil
}
func (c *CmdVoid3Remapped) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if err := c.In.A.Decode(d); err != nil {
		return err
	}
	if err := c.In.B.Decode(d); err != nil {
		return err
	}
	if err := c.In.C.Decode(d); err != nil {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOut3Remapped
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOut3Remapped_In struct {
}
type CmdVoidOut3Remapped_Out struct {
	A remapped
	B remapped
	C remapped
}
type CmdVoidOut3Remapped struct {
	Context atom.ContextID
	In      CmdVoidOut3Remapped_In
	Out     CmdVoidOut3Remapped_Out
}

func (c *CmdVoidOut3Remapped) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_out_3_remapped(",
		fmt.Sprintf("a:%v", c.Out.A),
		", ",
		fmt.Sprintf("b:%v", c.Out.B),
		", ",
		fmt.Sprintf("c:%v", c.Out.C),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidOut3Remapped) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdVoidOut3Remapped) TypeID() atom.TypeID {
	return 45
}
func (c *CmdVoidOut3Remapped) Flags() atom.Flags {
	return 0
}
func (c *CmdVoidOut3Remapped) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := c.Out.A.Encode(e); err != nil {
		return err
	}
	if err := c.Out.B.Encode(e); err != nil {
		return err
	}
	if err := c.Out.C.Encode(e); err != nil {
		return err
	}
	return nil
}
func (c *CmdVoidOut3Remapped) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if err := c.Out.A.Decode(d); err != nil {
		return err
	}
	if err := c.Out.B.Decode(d); err != nil {
		return err
	}
	if err := c.Out.C.Decode(d); err != nil {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOutArrayOfRemapped
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOutArrayOfRemapped_In struct {
}
type CmdVoidOutArrayOfRemapped_Out struct {
	A RemappedArray
}
type CmdVoidOutArrayOfRemapped struct {
	Context atom.ContextID
	In      CmdVoidOutArrayOfRemapped_In
	Out     CmdVoidOutArrayOfRemapped_Out
}

func (c *CmdVoidOutArrayOfRemapped) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "cmd_void_out_array_of_remapped(",
		fmt.Sprintf("%v", c.Out.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *CmdVoidOutArrayOfRemapped) ContextID() atom.ContextID {
	return c.Context
}
func (c *CmdVoidOutArrayOfRemapped) TypeID() atom.TypeID {
	return 46
}
func (c *CmdVoidOutArrayOfRemapped) Flags() atom.Flags {
	return 0
}
func (c *CmdVoidOutArrayOfRemapped) Encode(e *binary.Encoder) error {
	if err := c.Context.Encode(e); err != nil {
		return err
	}
	if err := c.Out.A.Encode(e); err != nil {
		return err
	}
	return nil
}
func (c *CmdVoidOutArrayOfRemapped) Decode(d *binary.Decoder) error {
	if err := c.Context.Decode(d); err != nil {
		return err
	}
	if err := c.Out.A.Decode(d); err != nil {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Globals
////////////////////////////////////////////////////////////////////////////////
type Globals struct {
}

func (g *Globals) Init() {
}
func (g *Globals) Encode(e *binary.Encoder) error {
	return nil
}
func (g *Globals) Decode(d *binary.Decoder) error {
	return nil
}
func NewCmdVoid() *CmdVoid {
	return &CmdVoid{
		In:  CmdVoid_In{},
		Out: CmdVoid_Out{},
	}
}
func NewCmdVoidU8(
	pA uint8,
) *CmdVoidU8 {
	return &CmdVoidU8{
		In:  CmdVoidU8_In{pA},
		Out: CmdVoidU8_Out{},
	}
}
func NewCmdVoidS8(
	pA int8,
) *CmdVoidS8 {
	return &CmdVoidS8{
		In:  CmdVoidS8_In{pA},
		Out: CmdVoidS8_Out{},
	}
}
func NewCmdVoidU16(
	pA uint16,
) *CmdVoidU16 {
	return &CmdVoidU16{
		In:  CmdVoidU16_In{pA},
		Out: CmdVoidU16_Out{},
	}
}
func NewCmdVoidS16(
	pA int16,
) *CmdVoidS16 {
	return &CmdVoidS16{
		In:  CmdVoidS16_In{pA},
		Out: CmdVoidS16_Out{},
	}
}
func NewCmdVoidF32(
	pA float32,
) *CmdVoidF32 {
	return &CmdVoidF32{
		In:  CmdVoidF32_In{pA},
		Out: CmdVoidF32_Out{},
	}
}
func NewCmdVoidU32(
	pA uint32,
) *CmdVoidU32 {
	return &CmdVoidU32{
		In:  CmdVoidU32_In{pA},
		Out: CmdVoidU32_Out{},
	}
}
func NewCmdVoidS32(
	pA int32,
) *CmdVoidS32 {
	return &CmdVoidS32{
		In:  CmdVoidS32_In{pA},
		Out: CmdVoidS32_Out{},
	}
}
func NewCmdVoidF64(
	pA float64,
) *CmdVoidF64 {
	return &CmdVoidF64{
		In:  CmdVoidF64_In{pA},
		Out: CmdVoidF64_Out{},
	}
}
func NewCmdVoidU64(
	pA uint64,
) *CmdVoidU64 {
	return &CmdVoidU64{
		In:  CmdVoidU64_In{pA},
		Out: CmdVoidU64_Out{},
	}
}
func NewCmdVoidS64(
	pA int64,
) *CmdVoidS64 {
	return &CmdVoidS64{
		In:  CmdVoidS64_In{pA},
		Out: CmdVoidS64_Out{},
	}
}
func NewCmdVoidBool(
	pA bool,
) *CmdVoidBool {
	return &CmdVoidBool{
		In:  CmdVoidBool_In{pA},
		Out: CmdVoidBool_Out{},
	}
}
func NewCmdVoidString(
	pA string,
) *CmdVoidString {
	return &CmdVoidString{
		In:  CmdVoidString_In{pA},
		Out: CmdVoidString_Out{},
	}
}
func NewCmdVoid3Strings(
	pA string,
	pB string,
	pC string,
) *CmdVoid3Strings {
	return &CmdVoid3Strings{
		In:  CmdVoid3Strings_In{pA, pB, pC},
		Out: CmdVoid3Strings_Out{},
	}
}
func NewCmdVoid3Arrays(
	pA S8Array,
	pB StringArray,
	pC BoolArray,
) *CmdVoid3Arrays {
	return &CmdVoid3Arrays{
		In:  CmdVoid3Arrays_In{pA, pB, pC},
		Out: CmdVoid3Arrays_Out{},
	}
}
func NewCmdVoidArrayOfStrings(
	pA StringArray,
) *CmdVoidArrayOfStrings {
	return &CmdVoidArrayOfStrings{
		In:  CmdVoidArrayOfStrings_In{pA},
		Out: CmdVoidArrayOfStrings_Out{},
	}
}
func NewCmdU8(
	pResult uint8,
) *CmdU8 {
	return &CmdU8{
		In:  CmdU8_In{},
		Out: CmdU8_Out{pResult},
	}
}
func NewCmdS8(
	pResult int8,
) *CmdS8 {
	return &CmdS8{
		In:  CmdS8_In{},
		Out: CmdS8_Out{pResult},
	}
}
func NewCmdU16(
	pResult uint16,
) *CmdU16 {
	return &CmdU16{
		In:  CmdU16_In{},
		Out: CmdU16_Out{pResult},
	}
}
func NewCmdS16(
	pResult int16,
) *CmdS16 {
	return &CmdS16{
		In:  CmdS16_In{},
		Out: CmdS16_Out{pResult},
	}
}
func NewCmdF32(
	pResult float32,
) *CmdF32 {
	return &CmdF32{
		In:  CmdF32_In{},
		Out: CmdF32_Out{pResult},
	}
}
func NewCmdU32(
	pResult uint32,
) *CmdU32 {
	return &CmdU32{
		In:  CmdU32_In{},
		Out: CmdU32_Out{pResult},
	}
}
func NewCmdS32(
	pResult int32,
) *CmdS32 {
	return &CmdS32{
		In:  CmdS32_In{},
		Out: CmdS32_Out{pResult},
	}
}
func NewCmdF64(
	pResult float64,
) *CmdF64 {
	return &CmdF64{
		In:  CmdF64_In{},
		Out: CmdF64_Out{pResult},
	}
}
func NewCmdU64(
	pResult uint64,
) *CmdU64 {
	return &CmdU64{
		In:  CmdU64_In{},
		Out: CmdU64_Out{pResult},
	}
}
func NewCmdS64(
	pResult int64,
) *CmdS64 {
	return &CmdS64{
		In:  CmdS64_In{},
		Out: CmdS64_Out{pResult},
	}
}
func NewCmdBool(
	pResult bool,
) *CmdBool {
	return &CmdBool{
		In:  CmdBool_In{},
		Out: CmdBool_Out{pResult},
	}
}
func NewCmdString(
	pResult string,
) *CmdString {
	return &CmdString{
		In:  CmdString_In{},
		Out: CmdString_Out{pResult},
	}
}
func NewCmdArrayOfFloat(
	pResult F32Array,
) *CmdArrayOfFloat {
	return &CmdArrayOfFloat{
		In:  CmdArrayOfFloat_In{},
		Out: CmdArrayOfFloat_Out{pResult},
	}
}
func NewCmdPointer(
	pResult memory.Pointer,
) *CmdPointer {
	return &CmdPointer{
		In:  CmdPointer_In{},
		Out: CmdPointer_Out{pResult},
	}
}
func NewCmdVoidOutU8(
	pA uint8,
) *CmdVoidOutU8 {
	return &CmdVoidOutU8{
		In:  CmdVoidOutU8_In{},
		Out: CmdVoidOutU8_Out{pA},
	}
}
func NewCmdVoidOutS8(
	pA int8,
) *CmdVoidOutS8 {
	return &CmdVoidOutS8{
		In:  CmdVoidOutS8_In{},
		Out: CmdVoidOutS8_Out{pA},
	}
}
func NewCmdVoidOutU16(
	pA uint16,
) *CmdVoidOutU16 {
	return &CmdVoidOutU16{
		In:  CmdVoidOutU16_In{},
		Out: CmdVoidOutU16_Out{pA},
	}
}
func NewCmdVoidOutS16(
	pA int16,
) *CmdVoidOutS16 {
	return &CmdVoidOutS16{
		In:  CmdVoidOutS16_In{},
		Out: CmdVoidOutS16_Out{pA},
	}
}
func NewCmdVoidOutF32(
	pA float32,
) *CmdVoidOutF32 {
	return &CmdVoidOutF32{
		In:  CmdVoidOutF32_In{},
		Out: CmdVoidOutF32_Out{pA},
	}
}
func NewCmdVoidOutU32(
	pA uint32,
) *CmdVoidOutU32 {
	return &CmdVoidOutU32{
		In:  CmdVoidOutU32_In{},
		Out: CmdVoidOutU32_Out{pA},
	}
}
func NewCmdVoidOutS32(
	pA int32,
) *CmdVoidOutS32 {
	return &CmdVoidOutS32{
		In:  CmdVoidOutS32_In{},
		Out: CmdVoidOutS32_Out{pA},
	}
}
func NewCmdVoidOutF64(
	pA float64,
) *CmdVoidOutF64 {
	return &CmdVoidOutF64{
		In:  CmdVoidOutF64_In{},
		Out: CmdVoidOutF64_Out{pA},
	}
}
func NewCmdVoidOutU64(
	pA uint64,
) *CmdVoidOutU64 {
	return &CmdVoidOutU64{
		In:  CmdVoidOutU64_In{},
		Out: CmdVoidOutU64_Out{pA},
	}
}
func NewCmdVoidOutS64(
	pA int64,
) *CmdVoidOutS64 {
	return &CmdVoidOutS64{
		In:  CmdVoidOutS64_In{},
		Out: CmdVoidOutS64_Out{pA},
	}
}
func NewCmdVoidOutBool(
	pA bool,
) *CmdVoidOutBool {
	return &CmdVoidOutBool{
		In:  CmdVoidOutBool_In{},
		Out: CmdVoidOutBool_Out{pA},
	}
}
func NewCmdVoidOutString(
	pA string,
) *CmdVoidOutString {
	return &CmdVoidOutString{
		In:  CmdVoidOutString_In{},
		Out: CmdVoidOutString_Out{pA},
	}
}
func NewCmdVoidOutFixedSizeBuffer(
	pA memory.Pointer,
) *CmdVoidOutFixedSizeBuffer {
	return &CmdVoidOutFixedSizeBuffer{
		In:  CmdVoidOutFixedSizeBuffer_In{},
		Out: CmdVoidOutFixedSizeBuffer_Out{pA},
	}
}
func NewCmdVoidOut3Strings(
	pA string,
	pB string,
	pC string,
) *CmdVoidOut3Strings {
	return &CmdVoidOut3Strings{
		In:  CmdVoidOut3Strings_In{},
		Out: CmdVoidOut3Strings_Out{pA, pB, pC},
	}
}
func NewCmdVoid3Remapped(
	pA remapped,
	pB remapped,
	pC remapped,
) *CmdVoid3Remapped {
	return &CmdVoid3Remapped{
		In:  CmdVoid3Remapped_In{pA, pB, pC},
		Out: CmdVoid3Remapped_Out{},
	}
}
func NewCmdVoidOut3Remapped(
	pA remapped,
	pB remapped,
	pC remapped,
) *CmdVoidOut3Remapped {
	return &CmdVoidOut3Remapped{
		In:  CmdVoidOut3Remapped_In{},
		Out: CmdVoidOut3Remapped_Out{pA, pB, pC},
	}
}
func NewCmdVoidOutArrayOfRemapped(
	pA RemappedArray,
) *CmdVoidOutArrayOfRemapped {
	return &CmdVoidOutArrayOfRemapped{
		In:  CmdVoidOutArrayOfRemapped_In{},
		Out: CmdVoidOutArrayOfRemapped_Out{pA},
	}
}

////////////////////////////////////////////////////////////////////////////////
// API
////////////////////////////////////////////////////////////////////////////////
type api struct{}

func (a api) Name() string {
	return "gfxapi_test"
}
func (a api) InitialState() gfxapi.State {
	return initialState()
}
func (a api) StateMutator(s gfxapi.State) atom.Writer {
	return &StateMutator{State: s.(*state)}
}
func API() gfxapi.API {
	return api{}
}
func init() {
	gfxapi.Register(API())
	atom.Register(atom.TypeInfo{
		Name: "CmdVoid",
		Docs: "",
		ID:   0,
		New:  func() atom.Atom { return &CmdVoid{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidU8",
		Docs: "",
		ID:   1,
		New:  func() atom.Atom { return &CmdVoidU8{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidS8",
		Docs: "",
		ID:   2,
		New:  func() atom.Atom { return &CmdVoidS8{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidU16",
		Docs: "",
		ID:   3,
		New:  func() atom.Atom { return &CmdVoidU16{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidS16",
		Docs: "",
		ID:   4,
		New:  func() atom.Atom { return &CmdVoidS16{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidF32",
		Docs: "",
		ID:   5,
		New:  func() atom.Atom { return &CmdVoidF32{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidU32",
		Docs: "",
		ID:   6,
		New:  func() atom.Atom { return &CmdVoidU32{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidS32",
		Docs: "",
		ID:   7,
		New:  func() atom.Atom { return &CmdVoidS32{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidF64",
		Docs: "",
		ID:   8,
		New:  func() atom.Atom { return &CmdVoidF64{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidU64",
		Docs: "",
		ID:   9,
		New:  func() atom.Atom { return &CmdVoidU64{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidS64",
		Docs: "",
		ID:   10,
		New:  func() atom.Atom { return &CmdVoidS64{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidBool",
		Docs: "",
		ID:   11,
		New:  func() atom.Atom { return &CmdVoidBool{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidString",
		Docs: "",
		ID:   12,
		New:  func() atom.Atom { return &CmdVoidString{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoid3Strings",
		Docs: "",
		ID:   13,
		New:  func() atom.Atom { return &CmdVoid3Strings{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoid3Arrays",
		Docs: "",
		ID:   14,
		New:  func() atom.Atom { return &CmdVoid3Arrays{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidArrayOfStrings",
		Docs: "",
		ID:   15,
		New:  func() atom.Atom { return &CmdVoidArrayOfStrings{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdU8",
		Docs: "",
		ID:   16,
		New:  func() atom.Atom { return &CmdU8{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdS8",
		Docs: "",
		ID:   17,
		New:  func() atom.Atom { return &CmdS8{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdU16",
		Docs: "",
		ID:   18,
		New:  func() atom.Atom { return &CmdU16{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdS16",
		Docs: "",
		ID:   19,
		New:  func() atom.Atom { return &CmdS16{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdF32",
		Docs: "",
		ID:   20,
		New:  func() atom.Atom { return &CmdF32{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdU32",
		Docs: "",
		ID:   21,
		New:  func() atom.Atom { return &CmdU32{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdS32",
		Docs: "",
		ID:   22,
		New:  func() atom.Atom { return &CmdS32{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdF64",
		Docs: "",
		ID:   23,
		New:  func() atom.Atom { return &CmdF64{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdU64",
		Docs: "",
		ID:   24,
		New:  func() atom.Atom { return &CmdU64{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdS64",
		Docs: "",
		ID:   25,
		New:  func() atom.Atom { return &CmdS64{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdBool",
		Docs: "",
		ID:   26,
		New:  func() atom.Atom { return &CmdBool{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdString",
		Docs: "",
		ID:   27,
		New:  func() atom.Atom { return &CmdString{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdArrayOfFloat",
		Docs: "",
		ID:   28,
		New:  func() atom.Atom { return &CmdArrayOfFloat{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdPointer",
		Docs: "",
		ID:   29,
		New:  func() atom.Atom { return &CmdPointer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidOutU8",
		Docs: "",
		ID:   30,
		New:  func() atom.Atom { return &CmdVoidOutU8{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidOutS8",
		Docs: "",
		ID:   31,
		New:  func() atom.Atom { return &CmdVoidOutS8{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidOutU16",
		Docs: "",
		ID:   32,
		New:  func() atom.Atom { return &CmdVoidOutU16{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidOutS16",
		Docs: "",
		ID:   33,
		New:  func() atom.Atom { return &CmdVoidOutS16{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidOutF32",
		Docs: "",
		ID:   34,
		New:  func() atom.Atom { return &CmdVoidOutF32{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidOutU32",
		Docs: "",
		ID:   35,
		New:  func() atom.Atom { return &CmdVoidOutU32{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidOutS32",
		Docs: "",
		ID:   36,
		New:  func() atom.Atom { return &CmdVoidOutS32{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidOutF64",
		Docs: "",
		ID:   37,
		New:  func() atom.Atom { return &CmdVoidOutF64{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidOutU64",
		Docs: "",
		ID:   38,
		New:  func() atom.Atom { return &CmdVoidOutU64{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidOutS64",
		Docs: "",
		ID:   39,
		New:  func() atom.Atom { return &CmdVoidOutS64{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidOutBool",
		Docs: "",
		ID:   40,
		New:  func() atom.Atom { return &CmdVoidOutBool{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidOutString",
		Docs: "",
		ID:   41,
		New:  func() atom.Atom { return &CmdVoidOutString{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidOutFixedSizeBuffer",
		Docs: "",
		ID:   42,
		New:  func() atom.Atom { return &CmdVoidOutFixedSizeBuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidOut3Strings",
		Docs: "",
		ID:   43,
		New:  func() atom.Atom { return &CmdVoidOut3Strings{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoid3Remapped",
		Docs: "",
		ID:   44,
		New:  func() atom.Atom { return &CmdVoid3Remapped{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidOut3Remapped",
		Docs: "",
		ID:   45,
		New:  func() atom.Atom { return &CmdVoidOut3Remapped{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CmdVoidOutArrayOfRemapped",
		Docs: "",
		ID:   46,
		New:  func() atom.Atom { return &CmdVoidOutArrayOfRemapped{} },
	})
}
