////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

package test

import (
	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/replay/builder"
	"android.googlesource.com/platform/tools/gpu/replay/protocol"
	"android.googlesource.com/platform/tools/gpu/replay/value"
)

type postCall []func()

func (p *postCall) add(f func()) {
	*p = append(*p, f)
}
func (p *postCall) exec() {
	for _, f := range *p {
		f()
	}
}
func loadRemap(ϟb *builder.Builder, key interface{}, val value.Value) {
	if ptr, found := ϟb.Remappings[key]; found {
		ϟb.Load(val.Type(), ptr)
	} else {
		ptr = ϟb.AllocateMemory(uint64(val.Type().Size(ϟb.Architecture().PointerSize)))
		ϟb.Push(val) // We have an input to an unknown id, use the unmapped value.
		ϟb.Clone(0)
		ϟb.Store(ptr)
		ϟb.Remappings[key] = ptr
	}
}
func storeRemap(ϟb *builder.Builder, key interface{}, val value.Pointer, ty protocol.Type) {
	if ptr, found := ϟb.Remappings[key]; !found {
		ptr = ϟb.AllocateMemory(uint64(ty.Size(ϟb.Architecture().PointerSize)))
		ϟb.Load(ty, val)
		ϟb.Store(ptr)
		ϟb.Remappings[key] = ptr
	}
}

var funcInfoCmdClone = builder.FunctionInfo{ID: 0, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoCmdMake = builder.FunctionInfo{ID: 1, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdCopy = builder.FunctionInfo{ID: 2, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoCmdCharsliceToString = builder.FunctionInfo{ID: 3, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoCmdCharptrToString = builder.FunctionInfo{ID: 4, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoid = builder.FunctionInfo{ID: 5, ReturnType: protocol.TypeVoid, Parameters: 0}
var funcInfoCmdUnknownRet = builder.FunctionInfo{ID: 6, ReturnType: protocol.TypeInt64, Parameters: 0}
var funcInfoCmdUnknownWritePtr = builder.FunctionInfo{ID: 7, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdUnknownWriteSlice = builder.FunctionInfo{ID: 8, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidU8 = builder.FunctionInfo{ID: 9, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidS8 = builder.FunctionInfo{ID: 10, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidU16 = builder.FunctionInfo{ID: 11, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidS16 = builder.FunctionInfo{ID: 12, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidF32 = builder.FunctionInfo{ID: 13, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidU32 = builder.FunctionInfo{ID: 14, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidS32 = builder.FunctionInfo{ID: 15, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidF64 = builder.FunctionInfo{ID: 16, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidU64 = builder.FunctionInfo{ID: 17, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidS64 = builder.FunctionInfo{ID: 18, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidBool = builder.FunctionInfo{ID: 19, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidString = builder.FunctionInfo{ID: 20, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoid3Strings = builder.FunctionInfo{ID: 21, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoCmdVoid3InArrays = builder.FunctionInfo{ID: 22, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoCmdVoidReadU8 = builder.FunctionInfo{ID: 23, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidReadS8 = builder.FunctionInfo{ID: 24, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidReadU16 = builder.FunctionInfo{ID: 25, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidReadS16 = builder.FunctionInfo{ID: 26, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidReadF32 = builder.FunctionInfo{ID: 27, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidReadU32 = builder.FunctionInfo{ID: 28, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidReadS32 = builder.FunctionInfo{ID: 29, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidReadF64 = builder.FunctionInfo{ID: 30, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidReadU64 = builder.FunctionInfo{ID: 31, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidReadS64 = builder.FunctionInfo{ID: 32, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidReadBool = builder.FunctionInfo{ID: 33, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidReadPtrs = builder.FunctionInfo{ID: 34, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoCmdVoidWriteU8 = builder.FunctionInfo{ID: 35, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidWriteS8 = builder.FunctionInfo{ID: 36, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidWriteU16 = builder.FunctionInfo{ID: 37, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidWriteS16 = builder.FunctionInfo{ID: 38, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidWriteF32 = builder.FunctionInfo{ID: 39, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidWriteU32 = builder.FunctionInfo{ID: 40, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidWriteS32 = builder.FunctionInfo{ID: 41, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidWriteF64 = builder.FunctionInfo{ID: 42, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidWriteU64 = builder.FunctionInfo{ID: 43, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidWriteS64 = builder.FunctionInfo{ID: 44, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidWriteBool = builder.FunctionInfo{ID: 45, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidWritePtrs = builder.FunctionInfo{ID: 46, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoCmdU8 = builder.FunctionInfo{ID: 47, ReturnType: protocol.TypeUint8, Parameters: 0}
var funcInfoCmdS8 = builder.FunctionInfo{ID: 48, ReturnType: protocol.TypeInt8, Parameters: 0}
var funcInfoCmdU16 = builder.FunctionInfo{ID: 49, ReturnType: protocol.TypeUint16, Parameters: 0}
var funcInfoCmdS16 = builder.FunctionInfo{ID: 50, ReturnType: protocol.TypeInt16, Parameters: 0}
var funcInfoCmdF32 = builder.FunctionInfo{ID: 51, ReturnType: protocol.TypeFloat, Parameters: 0}
var funcInfoCmdU32 = builder.FunctionInfo{ID: 52, ReturnType: protocol.TypeUint32, Parameters: 0}
var funcInfoCmdS32 = builder.FunctionInfo{ID: 53, ReturnType: protocol.TypeInt32, Parameters: 0}
var funcInfoCmdF64 = builder.FunctionInfo{ID: 54, ReturnType: protocol.TypeDouble, Parameters: 0}
var funcInfoCmdU64 = builder.FunctionInfo{ID: 55, ReturnType: protocol.TypeUint64, Parameters: 0}
var funcInfoCmdS64 = builder.FunctionInfo{ID: 56, ReturnType: protocol.TypeInt64, Parameters: 0}
var funcInfoCmdBool = builder.FunctionInfo{ID: 57, ReturnType: protocol.TypeBool, Parameters: 0}
var funcInfoCmdString = builder.FunctionInfo{ID: 58, ReturnType: protocol.TypeVolatilePointer, Parameters: 0}
var funcInfoCmdPointer = builder.FunctionInfo{ID: 59, ReturnType: protocol.TypeVolatilePointer, Parameters: 0}
var funcInfoCmdVoid3Remapped = builder.FunctionInfo{ID: 60, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoCmdVoidInArrayOfRemapped = builder.FunctionInfo{ID: 61, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidOutArrayOfRemapped = builder.FunctionInfo{ID: 62, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidOutArrayOfUnknownRemapped = builder.FunctionInfo{ID: 63, ReturnType: protocol.TypeVoid, Parameters: 1}

func (c remapped) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U32(uint32(c))
}

var _ = replay.Replayer(&CmdClone{}) // interface compliance check
func (ϟa *CmdClone) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟc.Buf = ϟa.Src.Slice(uint64(uint32(0)), uint64(ϟa.Cnt), ϟs).Clone(ϟs)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.Src.value())
	ϟb.Push(value.U32(ϟa.Cnt))
	ϟb.Call(funcInfoCmdClone)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdMake{}) // interface compliance check
func (ϟa *CmdMake) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟc.Buf = MakeU8ˢ(uint64(ϟa.Cnt), ϟs)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U32(ϟa.Cnt))
	ϟb.Call(funcInfoCmdMake)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdCopy{}) // interface compliance check
func (ϟa *CmdCopy) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟc.Buf = MakeU8ˢ(uint64(ϟa.Cnt), ϟs)
	ϟc.Buf.replayCopy(ϟa.Src.Slice(uint64(uint32(0)), uint64(ϟa.Cnt), ϟs), ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.Src.value())
	ϟb.Push(value.U32(ϟa.Cnt))
	ϟb.Call(funcInfoCmdCopy)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdCharsliceToString{}) // interface compliance check
func (ϟa *CmdCharsliceToString) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟc.Str = string(ϟa.S.Slice(uint64(uint32(0)), uint64(ϟa.Len), ϟs).Read(ϟs, ϟd, ϟl))
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.S.value())
	ϟb.Push(value.U32(ϟa.Len))
	ϟb.Call(funcInfoCmdCharsliceToString)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdCharptrToString{}) // interface compliance check
func (ϟa *CmdCharptrToString) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟc.Str = string(ϟa.S.StringSlice(ϟs, ϟd, ϟl, false).Read(ϟs, ϟd, ϟl))
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.S.value())
	ϟb.Call(funcInfoCmdCharptrToString)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoid{}) // interface compliance check
func (ϟa *CmdVoid) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoCmdVoid)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdUnknownRet{}) // interface compliance check
func (ϟa *CmdUnknownRet) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoCmdUnknownRet)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdUnknownWritePtr{}) // interface compliance check
func (ϟa *CmdUnknownWritePtr) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.P.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(ϟa.P.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.P.value())
	ϟb.Call(funcInfoCmdUnknownWritePtr)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdUnknownWriteSlice{}) // interface compliance check
func (ϟa *CmdUnknownWriteSlice) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	count := int32(5)                                        // s32
	slice := ϟa.A.Slice(uint64(int32(0)), uint64(count), ϟs) // Intˢ
	for i := int32(int32(0)); i < count; i++ {
		unknown := int64(ϟa.A.Slice(uint64(int32(0)), uint64(count), ϟs).Index(uint64(i), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)) // int
		slice.Index(uint64(i), ϟs).replayWrite(unknown, ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
		_ = unknown
	}
	_, _ = count, slice
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdUnknownWriteSlice)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidU8{}) // interface compliance check
func (ϟa *CmdVoidU8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U8(ϟa.A))
	ϟb.Call(funcInfoCmdVoidU8)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidS8{}) // interface compliance check
func (ϟa *CmdVoidS8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.S8(ϟa.A))
	ϟb.Call(funcInfoCmdVoidS8)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidU16{}) // interface compliance check
func (ϟa *CmdVoidU16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U16(ϟa.A))
	ϟb.Call(funcInfoCmdVoidU16)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidS16{}) // interface compliance check
func (ϟa *CmdVoidS16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.S16(ϟa.A))
	ϟb.Call(funcInfoCmdVoidS16)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidF32{}) // interface compliance check
func (ϟa *CmdVoidF32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.F32(ϟa.A))
	ϟb.Call(funcInfoCmdVoidF32)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidU32{}) // interface compliance check
func (ϟa *CmdVoidU32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U32(ϟa.A))
	ϟb.Call(funcInfoCmdVoidU32)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidS32{}) // interface compliance check
func (ϟa *CmdVoidS32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.S32(ϟa.A))
	ϟb.Call(funcInfoCmdVoidS32)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidF64{}) // interface compliance check
func (ϟa *CmdVoidF64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.F64(ϟa.A))
	ϟb.Call(funcInfoCmdVoidF64)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidU64{}) // interface compliance check
func (ϟa *CmdVoidU64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U64(ϟa.A))
	ϟb.Call(funcInfoCmdVoidU64)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidS64{}) // interface compliance check
func (ϟa *CmdVoidS64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.S64(ϟa.A))
	ϟb.Call(funcInfoCmdVoidS64)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidBool{}) // interface compliance check
func (ϟa *CmdVoidBool) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.Bool(ϟa.A))
	ϟb.Call(funcInfoCmdVoidBool)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidString{}) // interface compliance check
func (ϟa *CmdVoidString) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟb.String(ϟa.A))
	ϟb.Call(funcInfoCmdVoidString)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoid3Strings{}) // interface compliance check
func (ϟa *CmdVoid3Strings) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟb.String(ϟa.A))
	ϟb.Push(ϟb.String(ϟa.B))
	ϟb.Push(ϟb.String(ϟa.C))
	ϟb.Call(funcInfoCmdVoid3Strings)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoid3InArrays{}) // interface compliance check
func (ϟa *CmdVoid3InArrays) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟc.Buf = MakeU8ˢ(uint64(10), ϟs)
	ϟa.B.Slice(uint64(5), uint64(15), ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.C.Slice(uint64(5), uint64(15), ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟc.Buf.replayCopy(ϟa.A.Slice(uint64(5), uint64(25), ϟs), ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Push(ϟa.B.value())
	ϟb.Push(ϟa.C.value())
	ϟb.Call(funcInfoCmdVoid3InArrays)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidReadU8{}) // interface compliance check
func (ϟa *CmdVoidReadU8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb) // u8
	_ = x
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidReadU8)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidReadS8{}) // interface compliance check
func (ϟa *CmdVoidReadS8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb) // s8
	_ = x
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidReadS8)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidReadU16{}) // interface compliance check
func (ϟa *CmdVoidReadU16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb) // u16
	_ = x
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidReadU16)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidReadS16{}) // interface compliance check
func (ϟa *CmdVoidReadS16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb) // s16
	_ = x
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidReadS16)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidReadF32{}) // interface compliance check
func (ϟa *CmdVoidReadF32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb) // f32
	_ = x
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidReadF32)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidReadU32{}) // interface compliance check
func (ϟa *CmdVoidReadU32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb) // u32
	_ = x
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidReadU32)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidReadS32{}) // interface compliance check
func (ϟa *CmdVoidReadS32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb) // s32
	_ = x
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidReadS32)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidReadF64{}) // interface compliance check
func (ϟa *CmdVoidReadF64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb) // f64
	_ = x
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidReadF64)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidReadU64{}) // interface compliance check
func (ϟa *CmdVoidReadU64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb) // u64
	_ = x
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidReadU64)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidReadS64{}) // interface compliance check
func (ϟa *CmdVoidReadS64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb) // s64
	_ = x
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidReadS64)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidReadBool{}) // interface compliance check
func (ϟa *CmdVoidReadBool) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb) // bool
	_ = x
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidReadBool)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidReadPtrs{}) // interface compliance check
func (ϟa *CmdVoidReadPtrs) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb) // f32
	y := ϟa.B.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb) // u16
	z := ϟa.C.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb) // bool
	_, _, _ = x, y, z
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Push(ϟa.B.value())
	ϟb.Push(ϟa.C.value())
	ϟb.Call(funcInfoCmdVoidReadPtrs)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidWriteU8{}) // interface compliance check
func (ϟa *CmdVoidWriteU8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(uint8(1), ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidWriteU8)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidWriteS8{}) // interface compliance check
func (ϟa *CmdVoidWriteS8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(int8(1), ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidWriteS8)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidWriteU16{}) // interface compliance check
func (ϟa *CmdVoidWriteU16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(uint16(1), ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidWriteU16)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidWriteS16{}) // interface compliance check
func (ϟa *CmdVoidWriteS16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(int16(1), ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidWriteS16)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidWriteF32{}) // interface compliance check
func (ϟa *CmdVoidWriteF32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(float32(1), ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidWriteF32)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidWriteU32{}) // interface compliance check
func (ϟa *CmdVoidWriteU32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(uint32(1), ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidWriteU32)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidWriteS32{}) // interface compliance check
func (ϟa *CmdVoidWriteS32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(int32(1), ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidWriteS32)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidWriteF64{}) // interface compliance check
func (ϟa *CmdVoidWriteF64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(float64(1), ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidWriteF64)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidWriteU64{}) // interface compliance check
func (ϟa *CmdVoidWriteU64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(uint64(1), ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidWriteU64)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidWriteS64{}) // interface compliance check
func (ϟa *CmdVoidWriteS64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(int64(1), ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidWriteS64)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidWriteBool{}) // interface compliance check
func (ϟa *CmdVoidWriteBool) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(true, ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidWriteBool)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidWritePtrs{}) // interface compliance check
func (ϟa *CmdVoidWritePtrs) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(float32(10), ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
	ϟa.B.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(uint16(20), ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
	ϟa.C.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(false, ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Push(ϟa.B.value())
	ϟb.Push(ϟa.C.value())
	ϟb.Call(funcInfoCmdVoidWritePtrs)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdU8{}) // interface compliance check
func (ϟa *CmdU8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = uint8(0)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoCmdU8)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdS8{}) // interface compliance check
func (ϟa *CmdS8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = int8(0)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoCmdS8)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdU16{}) // interface compliance check
func (ϟa *CmdU16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = uint16(0)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoCmdU16)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdS16{}) // interface compliance check
func (ϟa *CmdS16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = int16(0)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoCmdS16)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdF32{}) // interface compliance check
func (ϟa *CmdF32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = float32(0)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoCmdF32)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdU32{}) // interface compliance check
func (ϟa *CmdU32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = uint32(0)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoCmdU32)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdS32{}) // interface compliance check
func (ϟa *CmdS32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = int32(0)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoCmdS32)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdF64{}) // interface compliance check
func (ϟa *CmdF64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = float64(0)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoCmdF64)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdU64{}) // interface compliance check
func (ϟa *CmdU64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = uint64(0)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoCmdU64)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdS64{}) // interface compliance check
func (ϟa *CmdS64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = int64(0)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoCmdS64)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdBool{}) // interface compliance check
func (ϟa *CmdBool) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = false
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoCmdBool)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdString{}) // interface compliance check
func (ϟa *CmdString) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ""
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoCmdString)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdPointer{}) // interface compliance check
func (ϟa *CmdPointer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoCmdPointer)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoid3Remapped{}) // interface compliance check
func (ϟa *CmdVoid3Remapped) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
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
	ϟb.Call(funcInfoCmdVoid3Remapped)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidInArrayOfRemapped{}) // interface compliance check
func (ϟa *CmdVoidInArrayOfRemapped) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(5), ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidInArrayOfRemapped)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidOutArrayOfRemapped{}) // interface compliance check
func (ϟa *CmdVoidOutArrayOfRemapped) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(5), ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidOutArrayOfRemapped)
	ϟp.exec()
	return nil
}

var _ = replay.Replayer(&CmdVoidOutArrayOfUnknownRemapped{}) // interface compliance check
func (ϟa *CmdVoidOutArrayOfUnknownRemapped) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	ϟp := &postCall{}
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	count := int32(5)                                        // s32
	slice := ϟa.A.Slice(uint64(int32(0)), uint64(count), ϟs) // Remappedˢ
	for i := int32(int32(0)); i < count; i++ {
		unknown := remapped(ϟa.A.Slice(uint64(int32(0)), uint64(count), ϟs).Index(uint64(i), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)) // remapped
		slice.Index(uint64(i), ϟs).replayWrite(unknown, ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
		_ = unknown
	}
	_, _ = count, slice
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidOutArrayOfUnknownRemapped)
	ϟp.exec()
	return nil
}
func (p U8ᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint8 {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p U8ᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint8 {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p U8ᵖ) replayWrite(value uint8, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
}
func (p U8ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.VolatileCapturePointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p Charᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) byte {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Charᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) byte {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Charᵖ) replayWrite(value byte, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
}
func (p Charᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.VolatileCapturePointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p Intᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) int64 {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Intᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) int64 {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Intᵖ) replayWrite(value int64, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
}
func (p Intᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.VolatileCapturePointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p U8ᵖᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U8ᵖ {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p U8ᵖᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U8ᵖ {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p U8ᵖᵖ) replayWrite(value U8ᵖ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
}
func (p U8ᵖᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.VolatileCapturePointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p S8ᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) int8 {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p S8ᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) int8 {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p S8ᵖ) replayWrite(value int8, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
}
func (p S8ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.VolatileCapturePointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p U16ᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint16 {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p U16ᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint16 {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p U16ᵖ) replayWrite(value uint16, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
}
func (p U16ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.VolatileCapturePointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p S16ᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) int16 {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p S16ᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) int16 {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p S16ᵖ) replayWrite(value int16, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
}
func (p S16ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.VolatileCapturePointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p F32ᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) float32 {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p F32ᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) float32 {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p F32ᵖ) replayWrite(value float32, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
}
func (p F32ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.VolatileCapturePointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p U32ᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint32 {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p U32ᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint32 {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p U32ᵖ) replayWrite(value uint32, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
}
func (p U32ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.VolatileCapturePointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p S32ᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) int32 {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p S32ᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) int32 {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p S32ᵖ) replayWrite(value int32, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
}
func (p S32ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.VolatileCapturePointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p F64ᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) float64 {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p F64ᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) float64 {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p F64ᵖ) replayWrite(value float64, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
}
func (p F64ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.VolatileCapturePointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p U64ᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint64 {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p U64ᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint64 {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p U64ᵖ) replayWrite(value uint64, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
}
func (p U64ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.VolatileCapturePointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p S64ᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) int64 {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p S64ᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) int64 {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p S64ᵖ) replayWrite(value int64, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
}
func (p S64ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.VolatileCapturePointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p Boolᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) bool {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Boolᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) bool {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Boolᵖ) replayWrite(value bool, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
}
func (p Boolᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.VolatileCapturePointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p Voidᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.VolatileCapturePointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p Remappedᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) remapped {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Remappedᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) remapped {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Remappedᵖ) replayWrite(value remapped, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
}
func (p Remappedᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.VolatileCapturePointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (s Boolˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
}
func (s Boolˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
}
func (s Boolˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s Boolˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []bool {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (dst Boolˢ) replayCopy(src Boolˢ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	dst, src = dst.Copy(src, ϟs, ϟd, ϟl)
	src.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	dst.onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
}
func (s Charˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
}
func (s Charˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
}
func (s Charˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s Charˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []byte {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (dst Charˢ) replayCopy(src Charˢ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	dst, src = dst.Copy(src, ϟs, ϟd, ϟl)
	src.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	dst.onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
}
func (s F32ˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
}
func (s F32ˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
}
func (s F32ˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s F32ˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []float32 {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (dst F32ˢ) replayCopy(src F32ˢ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	dst, src = dst.Copy(src, ϟs, ϟd, ϟl)
	src.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	dst.onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
}
func (s F64ˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
}
func (s F64ˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
}
func (s F64ˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s F64ˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []float64 {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (dst F64ˢ) replayCopy(src F64ˢ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	dst, src = dst.Copy(src, ϟs, ϟd, ϟl)
	src.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	dst.onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
}
func (s Intˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
}
func (s Intˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
}
func (s Intˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s Intˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []int64 {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (dst Intˢ) replayCopy(src Intˢ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	dst, src = dst.Copy(src, ϟs, ϟd, ϟl)
	src.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	dst.onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
}
func (s Remappedˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ptr, step := value.VolatileCapturePointer(s.Base), value.VolatileCapturePointer(s.ElementSize(ϟs))
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
			if key, remap := v.remap(ϟa, ϟs); remap {
				loadRemap(ϟb, key, v.value(ϟb, ϟa, ϟs))
			} else {
				ϟb.Push(v.value(ϟb, ϟa, ϟs))
			}
			ϟb.Store(ptr)
			ptr += step
		}
	}
}
func (s Remappedˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
		ϟp.add(func() {
			ptr, step := value.VolatileCapturePointer(s.Base), value.VolatileCapturePointer(s.ElementSize(ϟs))
			for _, v := range s.Read(ϟs, ϟd, ϟl) {
				if key, remap := v.remap(ϟa, ϟs); remap {
					storeRemap(ϟb, key, ptr, protocol.TypeUint32)
				}
				ptr += step
			}
		})
	}
}
func (s Remappedˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s Remappedˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []remapped {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (dst Remappedˢ) replayCopy(src Remappedˢ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	dst, src = dst.Copy(src, ϟs, ϟd, ϟl)
	src.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	dst.onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
}
func (s S16ˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
}
func (s S16ˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
}
func (s S16ˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s S16ˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []int16 {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (dst S16ˢ) replayCopy(src S16ˢ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	dst, src = dst.Copy(src, ϟs, ϟd, ϟl)
	src.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	dst.onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
}
func (s S32ˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
}
func (s S32ˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
}
func (s S32ˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s S32ˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []int32 {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (dst S32ˢ) replayCopy(src S32ˢ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	dst, src = dst.Copy(src, ϟs, ϟd, ϟl)
	src.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	dst.onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
}
func (s S64ˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
}
func (s S64ˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
}
func (s S64ˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s S64ˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []int64 {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (dst S64ˢ) replayCopy(src S64ˢ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	dst, src = dst.Copy(src, ϟs, ϟd, ϟl)
	src.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	dst.onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
}
func (s S8ˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
}
func (s S8ˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
}
func (s S8ˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s S8ˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []int8 {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (dst S8ˢ) replayCopy(src S8ˢ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	dst, src = dst.Copy(src, ϟs, ϟd, ϟl)
	src.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	dst.onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
}
func (s U16ˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
}
func (s U16ˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
}
func (s U16ˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s U16ˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []uint16 {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (dst U16ˢ) replayCopy(src U16ˢ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	dst, src = dst.Copy(src, ϟs, ϟd, ϟl)
	src.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	dst.onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
}
func (s U32ˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
}
func (s U32ˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
}
func (s U32ˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s U32ˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []uint32 {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (dst U32ˢ) replayCopy(src U32ˢ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	dst, src = dst.Copy(src, ϟs, ϟd, ϟl)
	src.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	dst.onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
}
func (s U64ˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
}
func (s U64ˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
}
func (s U64ˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s U64ˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []uint64 {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (dst U64ˢ) replayCopy(src U64ˢ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	dst, src = dst.Copy(src, ϟs, ϟd, ϟl)
	src.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	dst.onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
}
func (s U8ˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
}
func (s U8ˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
}
func (s U8ˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s U8ˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []uint8 {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (dst U8ˢ) replayCopy(src U8ˢ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	dst, src = dst.Copy(src, ϟs, ϟd, ϟl)
	src.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	dst.onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
}
func (s U8ᵖˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
}
func (s U8ᵖˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
}
func (s U8ᵖˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s U8ᵖˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []U8ᵖ {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (dst U8ᵖˢ) replayCopy(src U8ᵖˢ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	dst, src = dst.Copy(src, ϟs, ϟd, ϟl)
	src.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	dst.onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb, ϟp)
}
func (s Voidˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
}
func (s Voidˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder, ϟp *postCall) {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
}
func (s Voidˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
