////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

package test

import (
	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/replay/builder"
	"android.googlesource.com/platform/tools/gpu/replay/protocol"
	"android.googlesource.com/platform/tools/gpu/replay/value"
)

func readBytes(r binary.Reader, c uint64) ([]byte, error) {
	b := make([]byte, c)
	err := r.Data(b)
	return b, err
}
func readString(r binary.Reader, c uint64) (string, error) {
	if buf, err := readBytes(r, c); err == nil {
		str := string(buf)
		for i, c := range str {
			if c == 0 {
				return str[:i], nil
			}
		}
		return str, nil
	} else {
		return "", err
	}
}

var funcInfoCmdVoid = builder.FunctionInfo{ID: 0, ReturnType: protocol.TypeVoid, Parameters: 0}
var funcInfoCmdVoidU8 = builder.FunctionInfo{ID: 1, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidS8 = builder.FunctionInfo{ID: 2, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidU16 = builder.FunctionInfo{ID: 3, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidS16 = builder.FunctionInfo{ID: 4, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidF32 = builder.FunctionInfo{ID: 5, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidU32 = builder.FunctionInfo{ID: 6, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidS32 = builder.FunctionInfo{ID: 7, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidF64 = builder.FunctionInfo{ID: 8, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidU64 = builder.FunctionInfo{ID: 9, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidS64 = builder.FunctionInfo{ID: 10, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidBool = builder.FunctionInfo{ID: 11, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidString = builder.FunctionInfo{ID: 12, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoid3Strings = builder.FunctionInfo{ID: 13, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoCmdVoid3Arrays = builder.FunctionInfo{ID: 14, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoCmdVoidArrayOfStrings = builder.FunctionInfo{ID: 15, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdU8 = builder.FunctionInfo{ID: 16, ReturnType: protocol.TypeUint8, Parameters: 0}
var funcInfoCmdS8 = builder.FunctionInfo{ID: 17, ReturnType: protocol.TypeInt8, Parameters: 0}
var funcInfoCmdU16 = builder.FunctionInfo{ID: 18, ReturnType: protocol.TypeUint16, Parameters: 0}
var funcInfoCmdS16 = builder.FunctionInfo{ID: 19, ReturnType: protocol.TypeInt16, Parameters: 0}
var funcInfoCmdF32 = builder.FunctionInfo{ID: 20, ReturnType: protocol.TypeFloat, Parameters: 0}
var funcInfoCmdU32 = builder.FunctionInfo{ID: 21, ReturnType: protocol.TypeUint32, Parameters: 0}
var funcInfoCmdS32 = builder.FunctionInfo{ID: 22, ReturnType: protocol.TypeInt32, Parameters: 0}
var funcInfoCmdF64 = builder.FunctionInfo{ID: 23, ReturnType: protocol.TypeDouble, Parameters: 0}
var funcInfoCmdU64 = builder.FunctionInfo{ID: 24, ReturnType: protocol.TypeUint64, Parameters: 0}
var funcInfoCmdS64 = builder.FunctionInfo{ID: 25, ReturnType: protocol.TypeInt64, Parameters: 0}
var funcInfoCmdBool = builder.FunctionInfo{ID: 26, ReturnType: protocol.TypeBool, Parameters: 0}
var funcInfoCmdString = builder.FunctionInfo{ID: 27, ReturnType: protocol.TypeAbsolutePointer, Parameters: 0}
var funcInfoCmdArrayOfFloat = builder.FunctionInfo{ID: 28, ReturnType: protocol.TypeAbsolutePointer, Parameters: 0}
var funcInfoCmdPointer = builder.FunctionInfo{ID: 29, ReturnType: protocol.TypeAbsolutePointer, Parameters: 0}
var funcInfoCmdVoidOutU8 = builder.FunctionInfo{ID: 30, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidOutS8 = builder.FunctionInfo{ID: 31, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidOutU16 = builder.FunctionInfo{ID: 32, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidOutS16 = builder.FunctionInfo{ID: 33, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidOutF32 = builder.FunctionInfo{ID: 34, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidOutU32 = builder.FunctionInfo{ID: 35, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidOutS32 = builder.FunctionInfo{ID: 36, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidOutF64 = builder.FunctionInfo{ID: 37, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidOutU64 = builder.FunctionInfo{ID: 38, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidOutS64 = builder.FunctionInfo{ID: 39, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidOutBool = builder.FunctionInfo{ID: 40, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidOutString = builder.FunctionInfo{ID: 41, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidOutFixedSizeBuffer = builder.FunctionInfo{ID: 42, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidOut3Strings = builder.FunctionInfo{ID: 43, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoCmdVoid3Remapped = builder.FunctionInfo{ID: 44, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoCmdVoidOut3Remapped = builder.FunctionInfo{ID: 45, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoCmdVoidOutArrayOfRemapped = builder.FunctionInfo{ID: 46, ReturnType: protocol.TypeVoid, Parameters: 1}

func (c remapped) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U32(uint32(c))
}
func (arr BoolArray) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			ϟb.Push(value.Bool(e))
		}
		return ϟb.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr F32Array) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			ϟb.Push(value.F32(e))
		}
		return ϟb.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr RemappedArray) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			if key, remap := e.remap(ϟa, ϟs); remap {
				loadRemap(ϟb, key, e.value(ϟb, ϟa, ϟs))
			} else {
				ϟb.Push(e.value(ϟb, ϟa, ϟs))
			}
		}
		return ϟb.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr S8Array) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			ϟb.Push(value.S8(e))
		}
		return ϟb.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr StringArray) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			ϟb.Push(ϟb.String(e))
		}
		return ϟb.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}

type CmdU8_Postback struct {
	Result uint8
}

func (o *CmdU8_Postback) Decode(d binary.Decoder) error {
	if v, err := d.Uint8(); err == nil {
		o.Result = v
	} else {
		return err
	}
	return nil
}

type CmdS8_Postback struct {
	Result int8
}

func (o *CmdS8_Postback) Decode(d binary.Decoder) error {
	if v, err := d.Int8(); err == nil {
		o.Result = v
	} else {
		return err
	}
	return nil
}

type CmdU16_Postback struct {
	Result uint16
}

func (o *CmdU16_Postback) Decode(d binary.Decoder) error {
	if v, err := d.Uint16(); err == nil {
		o.Result = v
	} else {
		return err
	}
	return nil
}

type CmdS16_Postback struct {
	Result int16
}

func (o *CmdS16_Postback) Decode(d binary.Decoder) error {
	if v, err := d.Int16(); err == nil {
		o.Result = v
	} else {
		return err
	}
	return nil
}

type CmdF32_Postback struct {
	Result float32
}

func (o *CmdF32_Postback) Decode(d binary.Decoder) error {
	if v, err := d.Float32(); err == nil {
		o.Result = v
	} else {
		return err
	}
	return nil
}

type CmdU32_Postback struct {
	Result uint32
}

func (o *CmdU32_Postback) Decode(d binary.Decoder) error {
	if v, err := d.Uint32(); err == nil {
		o.Result = v
	} else {
		return err
	}
	return nil
}

type CmdS32_Postback struct {
	Result int32
}

func (o *CmdS32_Postback) Decode(d binary.Decoder) error {
	if v, err := d.Int32(); err == nil {
		o.Result = v
	} else {
		return err
	}
	return nil
}

type CmdF64_Postback struct {
	Result float64
}

func (o *CmdF64_Postback) Decode(d binary.Decoder) error {
	if v, err := d.Float64(); err == nil {
		o.Result = v
	} else {
		return err
	}
	return nil
}

type CmdU64_Postback struct {
	Result uint64
}

func (o *CmdU64_Postback) Decode(d binary.Decoder) error {
	if v, err := d.Uint64(); err == nil {
		o.Result = v
	} else {
		return err
	}
	return nil
}

type CmdS64_Postback struct {
	Result int64
}

func (o *CmdS64_Postback) Decode(d binary.Decoder) error {
	if v, err := d.Int64(); err == nil {
		o.Result = v
	} else {
		return err
	}
	return nil
}

type CmdBool_Postback struct {
	Result bool
}

func (o *CmdBool_Postback) Decode(d binary.Decoder) error {
	if v, err := d.Bool(); err == nil {
		o.Result = v
	} else {
		return err
	}
	return nil
}

type CmdString_Postback struct {
	Result string
}

func (o *CmdString_Postback) Decode(result_cnt uint64, d binary.Decoder) error {
	if val, err := readString(d, result_cnt); err == nil {
		o.Result = val
	} else {
		return err
	}
	return nil
}

type CmdArrayOfFloat_Postback struct {
	Result F32Array
}

func (o *CmdArrayOfFloat_Postback) Decode(result_cnt uint64, d binary.Decoder) error {
	o.Result = make(F32Array, result_cnt)
	for i := range o.Result {
		if v, err := d.Float32(); err == nil {
			o.Result[i] = v
		} else {
			return err
		}
	}
	return nil
}

type CmdPointer_Postback struct {
	Result []byte
}

func (o *CmdPointer_Postback) Decode(result_cnt uint64, d binary.Decoder) error {
	if val, err := readBytes(d, result_cnt); err == nil {
		o.Result = val
	} else {
		return err
	}
	return nil
}

type CmdVoidOutU8_Postback struct {
	A uint8
}

func (o *CmdVoidOutU8_Postback) Decode(d binary.Decoder) error {
	if v, err := d.Uint8(); err == nil {
		o.A = v
	} else {
		return err
	}
	return nil
}

type CmdVoidOutS8_Postback struct {
	A int8
}

func (o *CmdVoidOutS8_Postback) Decode(d binary.Decoder) error {
	if v, err := d.Int8(); err == nil {
		o.A = v
	} else {
		return err
	}
	return nil
}

type CmdVoidOutU16_Postback struct {
	A uint16
}

func (o *CmdVoidOutU16_Postback) Decode(d binary.Decoder) error {
	if v, err := d.Uint16(); err == nil {
		o.A = v
	} else {
		return err
	}
	return nil
}

type CmdVoidOutS16_Postback struct {
	A int16
}

func (o *CmdVoidOutS16_Postback) Decode(d binary.Decoder) error {
	if v, err := d.Int16(); err == nil {
		o.A = v
	} else {
		return err
	}
	return nil
}

type CmdVoidOutF32_Postback struct {
	A float32
}

func (o *CmdVoidOutF32_Postback) Decode(d binary.Decoder) error {
	if v, err := d.Float32(); err == nil {
		o.A = v
	} else {
		return err
	}
	return nil
}

type CmdVoidOutU32_Postback struct {
	A uint32
}

func (o *CmdVoidOutU32_Postback) Decode(d binary.Decoder) error {
	if v, err := d.Uint32(); err == nil {
		o.A = v
	} else {
		return err
	}
	return nil
}

type CmdVoidOutS32_Postback struct {
	A int32
}

func (o *CmdVoidOutS32_Postback) Decode(d binary.Decoder) error {
	if v, err := d.Int32(); err == nil {
		o.A = v
	} else {
		return err
	}
	return nil
}

type CmdVoidOutF64_Postback struct {
	A float64
}

func (o *CmdVoidOutF64_Postback) Decode(d binary.Decoder) error {
	if v, err := d.Float64(); err == nil {
		o.A = v
	} else {
		return err
	}
	return nil
}

type CmdVoidOutU64_Postback struct {
	A uint64
}

func (o *CmdVoidOutU64_Postback) Decode(d binary.Decoder) error {
	if v, err := d.Uint64(); err == nil {
		o.A = v
	} else {
		return err
	}
	return nil
}

type CmdVoidOutS64_Postback struct {
	A int64
}

func (o *CmdVoidOutS64_Postback) Decode(d binary.Decoder) error {
	if v, err := d.Int64(); err == nil {
		o.A = v
	} else {
		return err
	}
	return nil
}

type CmdVoidOutBool_Postback struct {
	A bool
}

func (o *CmdVoidOutBool_Postback) Decode(d binary.Decoder) error {
	if v, err := d.Bool(); err == nil {
		o.A = v
	} else {
		return err
	}
	return nil
}

type CmdVoidOutString_Postback struct {
	A string
}

func (o *CmdVoidOutString_Postback) Decode(a_cnt uint64, d binary.Decoder) error {
	if val, err := readString(d, a_cnt); err == nil {
		o.A = val
	} else {
		return err
	}
	return nil
}

type CmdVoidOutFixedSizeBuffer_Postback struct {
	A []byte
}

func (o *CmdVoidOutFixedSizeBuffer_Postback) Decode(a_cnt uint64, d binary.Decoder) error {
	if val, err := readBytes(d, a_cnt); err == nil {
		o.A = val
	} else {
		return err
	}
	return nil
}

type CmdVoidOut3Strings_Postback struct {
	A string
	B string
	C string
}

func (o *CmdVoidOut3Strings_Postback) Decode(a_cnt uint64,
	b_cnt uint64,
	c_cnt uint64, d binary.Decoder) error {
	if val, err := readString(d, a_cnt); err == nil {
		o.A = val
	} else {
		return err
	}
	if val, err := readString(d, b_cnt); err == nil {
		o.B = val
	} else {
		return err
	}
	if val, err := readString(d, c_cnt); err == nil {
		o.C = val
	} else {
		return err
	}
	return nil
}

type CmdVoidOut3Remapped_Postback struct {
	A remapped
	B remapped
	C remapped
}

func (o *CmdVoidOut3Remapped_Postback) Decode(d binary.Decoder) error {
	{
		var x uint32
		if v, err := d.Uint32(); err == nil {
			x = v
		} else {
			return err
		}
		o.A = remapped(x)
	}
	{
		var x uint32
		if v, err := d.Uint32(); err == nil {
			x = v
		} else {
			return err
		}
		o.B = remapped(x)
	}
	{
		var x uint32
		if v, err := d.Uint32(); err == nil {
			x = v
		} else {
			return err
		}
		o.C = remapped(x)
	}
	return nil
}

type CmdVoidOutArrayOfRemapped_Postback struct {
	A RemappedArray
}

func (o *CmdVoidOutArrayOfRemapped_Postback) Decode(a_cnt uint64, d binary.Decoder) error {
	o.A = make(RemappedArray, a_cnt)
	for i := range o.A {
		{
			var x uint32
			if v, err := d.Uint32(); err == nil {
				x = v
			} else {
				return err
			}
			o.A[i] = remapped(x)
		}
	}
	return nil
}
func loadRemap(b *builder.Builder, key interface{}, val value.Value) {
	ptr, found := b.Remappings[key]
	if found {
		b.Load(val.Type(), ptr)
	} else {
		ptr = b.AllocateMemory(uint64(val.Type().Size(b.PointerSize())))
		b.Push(val) // We have an input to an unknown id, use the unmapped value.
		b.Clone(0)
		b.Store(ptr)
		b.Remappings[key] = ptr
	}
}
func storeRemap(b *builder.Builder, key interface{}, val value.Pointer, ty protocol.Type) {
	ptr, found := b.Remappings[key]
	if !found {
		ptr = b.AllocateMemory(uint64(ty.Size(b.PointerSize())))
		b.Load(ty, val)
		b.Store(ptr)
		b.Remappings[key] = ptr
	}
}

var _ = replay.Replayer(&CmdVoid{}) // interface compliance check
func (ϟa *CmdVoid) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.CallNoPush(funcInfoCmdVoid)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdVoidU8{}) // interface compliance check
func (ϟa *CmdVoidU8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U8(ϟa.A))
	ϟb.CallNoPush(funcInfoCmdVoidU8)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdVoidS8{}) // interface compliance check
func (ϟa *CmdVoidS8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.S8(ϟa.A))
	ϟb.CallNoPush(funcInfoCmdVoidS8)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdVoidU16{}) // interface compliance check
func (ϟa *CmdVoidU16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U16(ϟa.A))
	ϟb.CallNoPush(funcInfoCmdVoidU16)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdVoidS16{}) // interface compliance check
func (ϟa *CmdVoidS16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.S16(ϟa.A))
	ϟb.CallNoPush(funcInfoCmdVoidS16)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdVoidF32{}) // interface compliance check
func (ϟa *CmdVoidF32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.F32(ϟa.A))
	ϟb.CallNoPush(funcInfoCmdVoidF32)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdVoidU32{}) // interface compliance check
func (ϟa *CmdVoidU32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U32(ϟa.A))
	ϟb.CallNoPush(funcInfoCmdVoidU32)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdVoidS32{}) // interface compliance check
func (ϟa *CmdVoidS32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.S32(ϟa.A))
	ϟb.CallNoPush(funcInfoCmdVoidS32)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdVoidF64{}) // interface compliance check
func (ϟa *CmdVoidF64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.F64(ϟa.A))
	ϟb.CallNoPush(funcInfoCmdVoidF64)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdVoidU64{}) // interface compliance check
func (ϟa *CmdVoidU64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.U64(ϟa.A))
	ϟb.CallNoPush(funcInfoCmdVoidU64)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdVoidS64{}) // interface compliance check
func (ϟa *CmdVoidS64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.S64(ϟa.A))
	ϟb.CallNoPush(funcInfoCmdVoidS64)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdVoidBool{}) // interface compliance check
func (ϟa *CmdVoidBool) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(value.Bool(ϟa.A))
	ϟb.CallNoPush(funcInfoCmdVoidBool)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdVoidString{}) // interface compliance check
func (ϟa *CmdVoidString) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(ϟb.String(ϟa.A))
	ϟb.CallNoPush(funcInfoCmdVoidString)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdVoid3Strings{}) // interface compliance check
func (ϟa *CmdVoid3Strings) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(ϟb.String(ϟa.A))
	ϟb.Push(ϟb.String(ϟa.B))
	ϟb.Push(ϟb.String(ϟa.C))
	ϟb.CallNoPush(funcInfoCmdVoid3Strings)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdVoid3Arrays{}) // interface compliance check
func (ϟa *CmdVoid3Arrays) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(ϟa.A.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.B.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.C.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoCmdVoid3Arrays)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdVoidArrayOfStrings{}) // interface compliance check
func (ϟa *CmdVoidArrayOfStrings) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	ϟb.Push(ϟa.A.value(ϟb, ϟa, ϟs))
	ϟb.CallNoPush(funcInfoCmdVoidArrayOfStrings)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdU8{}) // interface compliance check
func (ϟa *CmdU8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{1 /* result */})
	ϟb.CallPush(funcInfoCmdU8)
	ϟb.Store(outputs[0])
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := CmdU8_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdS8{}) // interface compliance check
func (ϟa *CmdS8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{1 /* result */})
	ϟb.CallPush(funcInfoCmdS8)
	ϟb.Store(outputs[0])
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := CmdS8_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdU16{}) // interface compliance check
func (ϟa *CmdU16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{2 /* result */})
	ϟb.CallPush(funcInfoCmdU16)
	ϟb.Store(outputs[0])
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := CmdU16_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdS16{}) // interface compliance check
func (ϟa *CmdS16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{2 /* result */})
	ϟb.CallPush(funcInfoCmdS16)
	ϟb.Store(outputs[0])
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := CmdS16_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdF32{}) // interface compliance check
func (ϟa *CmdF32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{4 /* result */})
	ϟb.CallPush(funcInfoCmdF32)
	ϟb.Store(outputs[0])
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := CmdF32_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdU32{}) // interface compliance check
func (ϟa *CmdU32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{4 /* result */})
	ϟb.CallPush(funcInfoCmdU32)
	ϟb.Store(outputs[0])
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := CmdU32_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdS32{}) // interface compliance check
func (ϟa *CmdS32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{4 /* result */})
	ϟb.CallPush(funcInfoCmdS32)
	ϟb.Store(outputs[0])
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := CmdS32_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdF64{}) // interface compliance check
func (ϟa *CmdF64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{8 /* result */})
	ϟb.CallPush(funcInfoCmdF64)
	ϟb.Store(outputs[0])
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := CmdF64_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdU64{}) // interface compliance check
func (ϟa *CmdU64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{8 /* result */})
	ϟb.CallPush(funcInfoCmdU64)
	ϟb.Store(outputs[0])
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := CmdU64_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdS64{}) // interface compliance check
func (ϟa *CmdS64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{8 /* result */})
	ϟb.CallPush(funcInfoCmdS64)
	ϟb.Store(outputs[0])
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := CmdS64_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdBool{}) // interface compliance check
func (ϟa *CmdBool) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{1 /* result */})
	ϟb.CallPush(funcInfoCmdBool)
	ϟb.Store(outputs[0])
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := CmdBool_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdString{}) // interface compliance check
func (ϟa *CmdString) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	result_cnt := uint64(int32(10))
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{result_cnt /* result */})
	ϟb.CallPush(funcInfoCmdString)
	ϟb.Push(outputs[0])
	ϟb.Strcpy(result_cnt)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := CmdString_Postback{}
			if err := postback.Decode(result_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdArrayOfFloat{}) // interface compliance check
func (ϟa *CmdArrayOfFloat) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	result_cnt := uint64(int32(10))
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{result_cnt * 4 /* result */})
	ϟb.CallPush(funcInfoCmdArrayOfFloat)
	ϟb.Push(outputs[0])
	ϟb.Copy(result_cnt * 4)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := CmdArrayOfFloat_Postback{}
			if err := postback.Decode(result_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdPointer{}) // interface compliance check
func (ϟa *CmdPointer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	result_cnt := uint64(int32(10))
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{result_cnt /* result */})
	ϟb.CallPush(funcInfoCmdPointer)
	ϟb.Push(outputs[0])
	ϟb.Copy(result_cnt)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := CmdPointer_Postback{}
			if err := postback.Decode(result_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdVoidOutU8{}) // interface compliance check
func (ϟa *CmdVoidOutU8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{1 /* a */})
	ϟb.Push(outputs[0]) // a
	ϟb.CallNoPush(funcInfoCmdVoidOutU8)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := CmdVoidOutU8_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdVoidOutS8{}) // interface compliance check
func (ϟa *CmdVoidOutS8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{1 /* a */})
	ϟb.Push(outputs[0]) // a
	ϟb.CallNoPush(funcInfoCmdVoidOutS8)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := CmdVoidOutS8_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdVoidOutU16{}) // interface compliance check
func (ϟa *CmdVoidOutU16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{2 /* a */})
	ϟb.Push(outputs[0]) // a
	ϟb.CallNoPush(funcInfoCmdVoidOutU16)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := CmdVoidOutU16_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdVoidOutS16{}) // interface compliance check
func (ϟa *CmdVoidOutS16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{2 /* a */})
	ϟb.Push(outputs[0]) // a
	ϟb.CallNoPush(funcInfoCmdVoidOutS16)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := CmdVoidOutS16_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdVoidOutF32{}) // interface compliance check
func (ϟa *CmdVoidOutF32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{4 /* a */})
	ϟb.Push(outputs[0]) // a
	ϟb.CallNoPush(funcInfoCmdVoidOutF32)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := CmdVoidOutF32_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdVoidOutU32{}) // interface compliance check
func (ϟa *CmdVoidOutU32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{4 /* a */})
	ϟb.Push(outputs[0]) // a
	ϟb.CallNoPush(funcInfoCmdVoidOutU32)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := CmdVoidOutU32_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdVoidOutS32{}) // interface compliance check
func (ϟa *CmdVoidOutS32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{4 /* a */})
	ϟb.Push(outputs[0]) // a
	ϟb.CallNoPush(funcInfoCmdVoidOutS32)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := CmdVoidOutS32_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdVoidOutF64{}) // interface compliance check
func (ϟa *CmdVoidOutF64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{8 /* a */})
	ϟb.Push(outputs[0]) // a
	ϟb.CallNoPush(funcInfoCmdVoidOutF64)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := CmdVoidOutF64_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdVoidOutU64{}) // interface compliance check
func (ϟa *CmdVoidOutU64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{8 /* a */})
	ϟb.Push(outputs[0]) // a
	ϟb.CallNoPush(funcInfoCmdVoidOutU64)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := CmdVoidOutU64_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdVoidOutS64{}) // interface compliance check
func (ϟa *CmdVoidOutS64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{8 /* a */})
	ϟb.Push(outputs[0]) // a
	ϟb.CallNoPush(funcInfoCmdVoidOutS64)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := CmdVoidOutS64_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdVoidOutBool{}) // interface compliance check
func (ϟa *CmdVoidOutBool) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{1 /* a */})
	ϟb.Push(outputs[0]) // a
	ϟb.CallNoPush(funcInfoCmdVoidOutBool)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := CmdVoidOutBool_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdVoidOutString{}) // interface compliance check
func (ϟa *CmdVoidOutString) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	a_cnt := uint64(int32(10))
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{a_cnt /* a */})
	ϟb.Push(outputs[0]) // a
	ϟb.CallNoPush(funcInfoCmdVoidOutString)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := CmdVoidOutString_Postback{}
			if err := postback.Decode(a_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdVoidOutFixedSizeBuffer{}) // interface compliance check
func (ϟa *CmdVoidOutFixedSizeBuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	a_cnt := uint64(int32(10))
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{a_cnt /* a */})
	ϟb.Push(outputs[0]) // a
	ϟb.CallNoPush(funcInfoCmdVoidOutFixedSizeBuffer)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := CmdVoidOutFixedSizeBuffer_Postback{}
			if err := postback.Decode(a_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdVoidOut3Strings{}) // interface compliance check
func (ϟa *CmdVoidOut3Strings) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	a_cnt := uint64(int32(15))
	b_cnt := uint64(int32(31))
	c_cnt := uint64(int32(47))
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{a_cnt /* a */, b_cnt /* b */, c_cnt /* c */})
	ϟb.Push(outputs[0]) // a
	ϟb.Push(outputs[1]) // b
	ϟb.Push(outputs[2]) // c
	ϟb.CallNoPush(funcInfoCmdVoidOut3Strings)
	ϟa.Mutate(ϟs)
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := CmdVoidOut3Strings_Postback{}
			if err := postback.Decode(a_cnt,
				b_cnt,
				c_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdVoid3Remapped{}) // interface compliance check
func (ϟa *CmdVoid3Remapped) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	if key, remap := ϟa.A.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.A.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.A.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.B.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.B.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.B.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.C.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, ϟa.C.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.C.value(ϟb, ϟa, ϟs))
	}
	ϟb.CallNoPush(funcInfoCmdVoid3Remapped)
	ϟa.Mutate(ϟs)
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdVoidOut3Remapped{}) // interface compliance check
func (ϟa *CmdVoidOut3Remapped) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{4 /* a */, 4 /* b */, 4 /* c */})
	ϟb.Push(outputs[0]) // a
	ϟb.Push(outputs[1]) // b
	ϟb.Push(outputs[2]) // c
	ϟb.CallNoPush(funcInfoCmdVoidOut3Remapped)
	ϟa.Mutate(ϟs)
	if key, remap := ϟa.A.remap(ϟa, ϟs); remap {
		storeRemap(ϟb, key, outputs[0], protocol.TypeUint32)
	}
	if key, remap := ϟa.B.remap(ϟa, ϟs); remap {
		storeRemap(ϟb, key, outputs[1], protocol.TypeUint32)
	}
	if key, remap := ϟa.C.remap(ϟa, ϟs); remap {
		storeRemap(ϟb, key, outputs[2], protocol.TypeUint32)
	}
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := CmdVoidOut3Remapped_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}

var _ = replay.Replayer(&CmdVoidOutArrayOfRemapped{}) // interface compliance check
func (ϟa *CmdVoidOutArrayOfRemapped) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟb.BeginAtom(ϟi)
	a_cnt := uint64(int32(5))
	outputs, size := ϟb.AllocateTemporaryMemoryChunks([]uint64{a_cnt * 4 /* a */})
	ϟb.Push(outputs[0]) // a
	ϟb.CallNoPush(funcInfoCmdVoidOutArrayOfRemapped)
	ϟa.Mutate(ϟs)
	for i, e := range ϟa.A {
		ptr := outputs[0].Offset(uint64(i * 4))
		if key, remap := e.remap(ϟa, ϟs); remap {
			storeRemap(ϟb, key, ptr, protocol.TypeUint32)
		}
	}
	if postback {
		ϟb.Post(outputs[0], size, ϟi, func(d binary.Decoder) (interface{}, error) {
			postback := CmdVoidOutArrayOfRemapped_Postback{}
			if err := postback.Decode(a_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
	ϟb.EndAtom()
}
