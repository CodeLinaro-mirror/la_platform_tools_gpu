////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

package test

import (
	"fmt"
	"io"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/replay/builder"
	"android.googlesource.com/platform/tools/gpu/replay/protocol"
	"android.googlesource.com/platform/tools/gpu/replay/value"
)

type replayer interface {
	replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool)
}

func readBytes(r io.Reader, c uint64) ([]byte, error) {
	b := make([]byte, c)
	_, err := io.ReadFull(r, b)
	return b, err
}
func readString(r io.Reader, c uint64) (string, error) {
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

func (c remapped) value(b *builder.Builder, ω atom.Atom, s *state) value.Value {
	return value.U32(uint32(c))
}
func (arr BoolArray) value(b *builder.Builder, ω atom.Atom, s *state) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			b.Push(value.Bool(e))
		}
		return b.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr F32Array) value(b *builder.Builder, ω atom.Atom, s *state) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			b.Push(value.F32(e))
		}
		return b.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr RemappedArray) value(b *builder.Builder, ω atom.Atom, s *state) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			if key, remap := e.remap(ω, s); remap {
				loadRemap(b, key, e.value(b, ω, s))
			} else {
				b.Push(e.value(b, ω, s))
			}
		}
		return b.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr S8Array) value(b *builder.Builder, ω atom.Atom, s *state) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			b.Push(value.S8(e))
		}
		return b.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}
func (arr StringArray) value(b *builder.Builder, ω atom.Atom, s *state) value.Pointer {
	if len(arr) > 0 {
		for _, e := range arr {
			b.Push(b.String(e))
		}
		return b.Buffer(len(arr))
	} else {
		return value.AbsolutePointer(0)
	}
}

type CmdU8_Postback struct {
	Result uint8
}

func (o *CmdU8_Postback) Decode(d *protocol.Decoder) error {
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

func (o *CmdS8_Postback) Decode(d *protocol.Decoder) error {
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

func (o *CmdU16_Postback) Decode(d *protocol.Decoder) error {
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

func (o *CmdS16_Postback) Decode(d *protocol.Decoder) error {
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

func (o *CmdF32_Postback) Decode(d *protocol.Decoder) error {
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

func (o *CmdU32_Postback) Decode(d *protocol.Decoder) error {
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

func (o *CmdS32_Postback) Decode(d *protocol.Decoder) error {
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

func (o *CmdF64_Postback) Decode(d *protocol.Decoder) error {
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

func (o *CmdU64_Postback) Decode(d *protocol.Decoder) error {
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

func (o *CmdS64_Postback) Decode(d *protocol.Decoder) error {
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

func (o *CmdBool_Postback) Decode(d *protocol.Decoder) error {
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

func (o *CmdString_Postback) Decode(result_cnt uint64, d *protocol.Decoder) error {
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

func (o *CmdArrayOfFloat_Postback) Decode(result_cnt uint64, d *protocol.Decoder) error {
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

func (o *CmdPointer_Postback) Decode(result_cnt uint64, d *protocol.Decoder) error {
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

func (o *CmdVoidOutU8_Postback) Decode(d *protocol.Decoder) error {
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

func (o *CmdVoidOutS8_Postback) Decode(d *protocol.Decoder) error {
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

func (o *CmdVoidOutU16_Postback) Decode(d *protocol.Decoder) error {
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

func (o *CmdVoidOutS16_Postback) Decode(d *protocol.Decoder) error {
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

func (o *CmdVoidOutF32_Postback) Decode(d *protocol.Decoder) error {
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

func (o *CmdVoidOutU32_Postback) Decode(d *protocol.Decoder) error {
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

func (o *CmdVoidOutS32_Postback) Decode(d *protocol.Decoder) error {
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

func (o *CmdVoidOutF64_Postback) Decode(d *protocol.Decoder) error {
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

func (o *CmdVoidOutU64_Postback) Decode(d *protocol.Decoder) error {
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

func (o *CmdVoidOutS64_Postback) Decode(d *protocol.Decoder) error {
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

func (o *CmdVoidOutBool_Postback) Decode(d *protocol.Decoder) error {
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

func (o *CmdVoidOutString_Postback) Decode(a_cnt uint64, d *protocol.Decoder) error {
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

func (o *CmdVoidOutFixedSizeBuffer_Postback) Decode(a_cnt uint64, d *protocol.Decoder) error {
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
	c_cnt uint64, d *protocol.Decoder) error {
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

func (o *CmdVoidOut3Remapped_Postback) Decode(d *protocol.Decoder) error {
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

func (o *CmdVoidOutArrayOfRemapped_Postback) Decode(a_cnt uint64, d *protocol.Decoder) error {
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

type replayWriter struct {
	builder *builder.Builder
	state   *state
}

func newReplayWriter(b *builder.Builder) *replayWriter {
	return &replayWriter{
		builder: b,
		state:   initialState(),
	}
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
func (r *replayWriter) Write(id atom.ID, a atom.Atom, wantOutput bool) {
	b := r.builder
	switch ω := a.(type) {
	case *memory.Observation:
		b.Observation(ω.Range, ω.ResourceID)
	case replayer:
		ω.replay(id, r.state, b, wantOutput)
	case *atom.EOS:
	default:
		panic(fmt.Errorf("Unsupported atom type %T for Write", ω))
	}
	b.EndAtom()
}

var _ = replayer(&CmdVoid{}) // interface compliance check
func (ω *CmdVoid) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.CallNoPush(funcInfoCmdVoid)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&CmdVoidU8{}) // interface compliance check
func (ω *CmdVoidU8) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U8(ω.In.A))
	b.CallNoPush(funcInfoCmdVoidU8)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&CmdVoidS8{}) // interface compliance check
func (ω *CmdVoidS8) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.S8(ω.In.A))
	b.CallNoPush(funcInfoCmdVoidS8)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&CmdVoidU16{}) // interface compliance check
func (ω *CmdVoidU16) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U16(ω.In.A))
	b.CallNoPush(funcInfoCmdVoidU16)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&CmdVoidS16{}) // interface compliance check
func (ω *CmdVoidS16) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.S16(ω.In.A))
	b.CallNoPush(funcInfoCmdVoidS16)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&CmdVoidF32{}) // interface compliance check
func (ω *CmdVoidF32) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.F32(ω.In.A))
	b.CallNoPush(funcInfoCmdVoidF32)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&CmdVoidU32{}) // interface compliance check
func (ω *CmdVoidU32) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U32(ω.In.A))
	b.CallNoPush(funcInfoCmdVoidU32)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&CmdVoidS32{}) // interface compliance check
func (ω *CmdVoidS32) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.S32(ω.In.A))
	b.CallNoPush(funcInfoCmdVoidS32)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&CmdVoidF64{}) // interface compliance check
func (ω *CmdVoidF64) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.F64(ω.In.A))
	b.CallNoPush(funcInfoCmdVoidF64)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&CmdVoidU64{}) // interface compliance check
func (ω *CmdVoidU64) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.U64(ω.In.A))
	b.CallNoPush(funcInfoCmdVoidU64)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&CmdVoidS64{}) // interface compliance check
func (ω *CmdVoidS64) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.S64(ω.In.A))
	b.CallNoPush(funcInfoCmdVoidS64)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&CmdVoidBool{}) // interface compliance check
func (ω *CmdVoidBool) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(value.Bool(ω.In.A))
	b.CallNoPush(funcInfoCmdVoidBool)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&CmdVoidString{}) // interface compliance check
func (ω *CmdVoidString) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(b.String(ω.In.A))
	b.CallNoPush(funcInfoCmdVoidString)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&CmdVoid3Strings{}) // interface compliance check
func (ω *CmdVoid3Strings) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(b.String(ω.In.A))
	b.Push(b.String(ω.In.B))
	b.Push(b.String(ω.In.C))
	b.CallNoPush(funcInfoCmdVoid3Strings)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&CmdVoid3Arrays{}) // interface compliance check
func (ω *CmdVoid3Arrays) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(ω.In.A.value(b, ω, s))
	b.Push(ω.In.B.value(b, ω, s))
	b.Push(ω.In.C.value(b, ω, s))
	b.CallNoPush(funcInfoCmdVoid3Arrays)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&CmdVoidArrayOfStrings{}) // interface compliance check
func (ω *CmdVoidArrayOfStrings) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	b.Push(ω.In.A.value(b, ω, s))
	b.CallNoPush(funcInfoCmdVoidArrayOfStrings)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&CmdU8{}) // interface compliance check
func (ω *CmdU8) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{1 /* result */})
	b.CallPush(funcInfoCmdU8)
	b.Store(outputs[0])
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := CmdU8_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&CmdS8{}) // interface compliance check
func (ω *CmdS8) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{1 /* result */})
	b.CallPush(funcInfoCmdS8)
	b.Store(outputs[0])
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := CmdS8_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&CmdU16{}) // interface compliance check
func (ω *CmdU16) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{2 /* result */})
	b.CallPush(funcInfoCmdU16)
	b.Store(outputs[0])
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := CmdU16_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&CmdS16{}) // interface compliance check
func (ω *CmdS16) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{2 /* result */})
	b.CallPush(funcInfoCmdS16)
	b.Store(outputs[0])
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := CmdS16_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&CmdF32{}) // interface compliance check
func (ω *CmdF32) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{4 /* result */})
	b.CallPush(funcInfoCmdF32)
	b.Store(outputs[0])
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := CmdF32_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&CmdU32{}) // interface compliance check
func (ω *CmdU32) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{4 /* result */})
	b.CallPush(funcInfoCmdU32)
	b.Store(outputs[0])
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := CmdU32_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&CmdS32{}) // interface compliance check
func (ω *CmdS32) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{4 /* result */})
	b.CallPush(funcInfoCmdS32)
	b.Store(outputs[0])
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := CmdS32_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&CmdF64{}) // interface compliance check
func (ω *CmdF64) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{8 /* result */})
	b.CallPush(funcInfoCmdF64)
	b.Store(outputs[0])
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := CmdF64_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&CmdU64{}) // interface compliance check
func (ω *CmdU64) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{8 /* result */})
	b.CallPush(funcInfoCmdU64)
	b.Store(outputs[0])
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := CmdU64_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&CmdS64{}) // interface compliance check
func (ω *CmdS64) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{8 /* result */})
	b.CallPush(funcInfoCmdS64)
	b.Store(outputs[0])
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := CmdS64_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&CmdBool{}) // interface compliance check
func (ω *CmdBool) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{1 /* result */})
	b.CallPush(funcInfoCmdBool)
	b.Store(outputs[0])
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := CmdBool_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&CmdString{}) // interface compliance check
func (ω *CmdString) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	result_cnt := uint64(10)
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{result_cnt /* result */})
	b.CallPush(funcInfoCmdString)
	b.Push(outputs[0])
	b.Strcpy(result_cnt)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := CmdString_Postback{}
			if err := postback.Decode(result_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&CmdArrayOfFloat{}) // interface compliance check
func (ω *CmdArrayOfFloat) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	result_cnt := uint64(10)
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{result_cnt * 4 /* result */})
	b.CallPush(funcInfoCmdArrayOfFloat)
	b.Push(outputs[0])
	b.Copy(result_cnt * 4)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := CmdArrayOfFloat_Postback{}
			if err := postback.Decode(result_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&CmdPointer{}) // interface compliance check
func (ω *CmdPointer) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	result_cnt := uint64(10)
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{result_cnt /* result */})
	b.CallPush(funcInfoCmdPointer)
	b.Push(outputs[0])
	b.Copy(result_cnt)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := CmdPointer_Postback{}
			if err := postback.Decode(result_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&CmdVoidOutU8{}) // interface compliance check
func (ω *CmdVoidOutU8) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{1 /* a */})
	b.Push(outputs[0]) // a
	b.CallNoPush(funcInfoCmdVoidOutU8)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := CmdVoidOutU8_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&CmdVoidOutS8{}) // interface compliance check
func (ω *CmdVoidOutS8) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{1 /* a */})
	b.Push(outputs[0]) // a
	b.CallNoPush(funcInfoCmdVoidOutS8)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := CmdVoidOutS8_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&CmdVoidOutU16{}) // interface compliance check
func (ω *CmdVoidOutU16) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{2 /* a */})
	b.Push(outputs[0]) // a
	b.CallNoPush(funcInfoCmdVoidOutU16)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := CmdVoidOutU16_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&CmdVoidOutS16{}) // interface compliance check
func (ω *CmdVoidOutS16) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{2 /* a */})
	b.Push(outputs[0]) // a
	b.CallNoPush(funcInfoCmdVoidOutS16)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := CmdVoidOutS16_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&CmdVoidOutF32{}) // interface compliance check
func (ω *CmdVoidOutF32) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{4 /* a */})
	b.Push(outputs[0]) // a
	b.CallNoPush(funcInfoCmdVoidOutF32)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := CmdVoidOutF32_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&CmdVoidOutU32{}) // interface compliance check
func (ω *CmdVoidOutU32) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{4 /* a */})
	b.Push(outputs[0]) // a
	b.CallNoPush(funcInfoCmdVoidOutU32)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := CmdVoidOutU32_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&CmdVoidOutS32{}) // interface compliance check
func (ω *CmdVoidOutS32) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{4 /* a */})
	b.Push(outputs[0]) // a
	b.CallNoPush(funcInfoCmdVoidOutS32)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := CmdVoidOutS32_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&CmdVoidOutF64{}) // interface compliance check
func (ω *CmdVoidOutF64) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{8 /* a */})
	b.Push(outputs[0]) // a
	b.CallNoPush(funcInfoCmdVoidOutF64)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := CmdVoidOutF64_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&CmdVoidOutU64{}) // interface compliance check
func (ω *CmdVoidOutU64) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{8 /* a */})
	b.Push(outputs[0]) // a
	b.CallNoPush(funcInfoCmdVoidOutU64)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := CmdVoidOutU64_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&CmdVoidOutS64{}) // interface compliance check
func (ω *CmdVoidOutS64) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{8 /* a */})
	b.Push(outputs[0]) // a
	b.CallNoPush(funcInfoCmdVoidOutS64)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := CmdVoidOutS64_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&CmdVoidOutBool{}) // interface compliance check
func (ω *CmdVoidOutBool) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{1 /* a */})
	b.Push(outputs[0]) // a
	b.CallNoPush(funcInfoCmdVoidOutBool)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := CmdVoidOutBool_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&CmdVoidOutString{}) // interface compliance check
func (ω *CmdVoidOutString) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	a_cnt := uint64(10)
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{a_cnt /* a */})
	b.Push(outputs[0]) // a
	b.CallNoPush(funcInfoCmdVoidOutString)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := CmdVoidOutString_Postback{}
			if err := postback.Decode(a_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&CmdVoidOutFixedSizeBuffer{}) // interface compliance check
func (ω *CmdVoidOutFixedSizeBuffer) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	a_cnt := uint64(10)
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{a_cnt /* a */})
	b.Push(outputs[0]) // a
	b.CallNoPush(funcInfoCmdVoidOutFixedSizeBuffer)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := CmdVoidOutFixedSizeBuffer_Postback{}
			if err := postback.Decode(a_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&CmdVoidOut3Strings{}) // interface compliance check
func (ω *CmdVoidOut3Strings) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	a_cnt := uint64(15)
	b_cnt := uint64(31)
	c_cnt := uint64(47)
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{a_cnt /* a */, b_cnt /* b */, c_cnt /* c */})
	b.Push(outputs[0]) // a
	b.Push(outputs[1]) // b
	b.Push(outputs[2]) // c
	b.CallNoPush(funcInfoCmdVoidOut3Strings)
	StateMutator{State: s}.Write(id, ω)
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := CmdVoidOut3Strings_Postback{}
			if err := postback.Decode(a_cnt,
				b_cnt,
				c_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&CmdVoid3Remapped{}) // interface compliance check
func (ω *CmdVoid3Remapped) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if key, remap := ω.In.A.remap(ω, s); remap {
		loadRemap(b, key, ω.In.A.value(b, ω, s))
	} else {
		b.Push(ω.In.A.value(b, ω, s))
	}
	if key, remap := ω.In.B.remap(ω, s); remap {
		loadRemap(b, key, ω.In.B.value(b, ω, s))
	} else {
		b.Push(ω.In.B.value(b, ω, s))
	}
	if key, remap := ω.In.C.remap(ω, s); remap {
		loadRemap(b, key, ω.In.C.value(b, ω, s))
	} else {
		b.Push(ω.In.C.value(b, ω, s))
	}
	b.CallNoPush(funcInfoCmdVoid3Remapped)
	StateMutator{State: s}.Write(id, ω)
}

var _ = replayer(&CmdVoidOut3Remapped{}) // interface compliance check
func (ω *CmdVoidOut3Remapped) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{4 /* a */, 4 /* b */, 4 /* c */})
	b.Push(outputs[0]) // a
	b.Push(outputs[1]) // b
	b.Push(outputs[2]) // c
	b.CallNoPush(funcInfoCmdVoidOut3Remapped)
	StateMutator{State: s}.Write(id, ω)
	if key, remap := ω.Out.A.remap(ω, s); remap {
		storeRemap(b, key, outputs[0], protocol.TypeUint32)
	}
	if key, remap := ω.Out.B.remap(ω, s); remap {
		storeRemap(b, key, outputs[1], protocol.TypeUint32)
	}
	if key, remap := ω.Out.C.remap(ω, s); remap {
		storeRemap(b, key, outputs[2], protocol.TypeUint32)
	}
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := CmdVoidOut3Remapped_Postback{}
			if err := postback.Decode(d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}

var _ = replayer(&CmdVoidOutArrayOfRemapped{}) // interface compliance check
func (ω *CmdVoidOutArrayOfRemapped) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	a_cnt := uint64(5)
	outputs, size := b.AllocateTemporaryMemoryChunks([]uint64{a_cnt * 4 /* a */})
	b.Push(outputs[0]) // a
	b.CallNoPush(funcInfoCmdVoidOutArrayOfRemapped)
	StateMutator{State: s}.Write(id, ω)
	for i, e := range ω.Out.A {
		ptr := outputs[0].Offset(uint64(i * 4))
		if key, remap := e.remap(ω, s); remap {
			storeRemap(b, key, ptr, protocol.TypeUint32)
		}
	}
	if wantOutput {
		b.Post(outputs[0], size, id, func(d *protocol.Decoder) (interface{}, error) {
			postback := CmdVoidOutArrayOfRemapped_Postback{}
			if err := postback.Decode(a_cnt, d); err != nil {
				return nil, err
			}
			return postback, nil
		})
	}
}
