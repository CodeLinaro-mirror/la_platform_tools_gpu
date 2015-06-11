////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

package test

import (
	"strings"

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

func loadRemap(ϟb *builder.Builder, key interface{}, ty protocol.Type, val value.Value) {
	if ptr, found := ϟb.Remappings[key]; found {
		ϟb.Load(ty, ptr)
	} else {
		ptr = ϟb.AllocateMemory(uint64(ty.Size(ϟb.Architecture().PointerSize)))
		ϟb.Push(val) // We have an input to an unknown id, use the unmapped value.
		ϟb.Clone(0)
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
var funcInfoCmdVoidInArrayOfPointers = builder.FunctionInfo{ID: 23, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoCmdVoidReadU8 = builder.FunctionInfo{ID: 24, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidReadS8 = builder.FunctionInfo{ID: 25, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidReadU16 = builder.FunctionInfo{ID: 26, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidReadS16 = builder.FunctionInfo{ID: 27, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidReadF32 = builder.FunctionInfo{ID: 28, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidReadU32 = builder.FunctionInfo{ID: 29, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidReadS32 = builder.FunctionInfo{ID: 30, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidReadF64 = builder.FunctionInfo{ID: 31, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidReadU64 = builder.FunctionInfo{ID: 32, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidReadS64 = builder.FunctionInfo{ID: 33, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidReadBool = builder.FunctionInfo{ID: 34, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidReadPtrs = builder.FunctionInfo{ID: 35, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoCmdVoidWriteU8 = builder.FunctionInfo{ID: 36, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidWriteS8 = builder.FunctionInfo{ID: 37, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidWriteU16 = builder.FunctionInfo{ID: 38, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidWriteS16 = builder.FunctionInfo{ID: 39, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidWriteF32 = builder.FunctionInfo{ID: 40, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidWriteU32 = builder.FunctionInfo{ID: 41, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidWriteS32 = builder.FunctionInfo{ID: 42, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidWriteF64 = builder.FunctionInfo{ID: 43, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidWriteU64 = builder.FunctionInfo{ID: 44, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidWriteS64 = builder.FunctionInfo{ID: 45, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidWriteBool = builder.FunctionInfo{ID: 46, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidWritePtrs = builder.FunctionInfo{ID: 47, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoCmdU8 = builder.FunctionInfo{ID: 48, ReturnType: protocol.TypeUint8, Parameters: 0}
var funcInfoCmdS8 = builder.FunctionInfo{ID: 49, ReturnType: protocol.TypeInt8, Parameters: 0}
var funcInfoCmdU16 = builder.FunctionInfo{ID: 50, ReturnType: protocol.TypeUint16, Parameters: 0}
var funcInfoCmdS16 = builder.FunctionInfo{ID: 51, ReturnType: protocol.TypeInt16, Parameters: 0}
var funcInfoCmdF32 = builder.FunctionInfo{ID: 52, ReturnType: protocol.TypeFloat, Parameters: 0}
var funcInfoCmdU32 = builder.FunctionInfo{ID: 53, ReturnType: protocol.TypeUint32, Parameters: 0}
var funcInfoCmdS32 = builder.FunctionInfo{ID: 54, ReturnType: protocol.TypeInt32, Parameters: 0}
var funcInfoCmdF64 = builder.FunctionInfo{ID: 55, ReturnType: protocol.TypeDouble, Parameters: 0}
var funcInfoCmdU64 = builder.FunctionInfo{ID: 56, ReturnType: protocol.TypeUint64, Parameters: 0}
var funcInfoCmdS64 = builder.FunctionInfo{ID: 57, ReturnType: protocol.TypeInt64, Parameters: 0}
var funcInfoCmdBool = builder.FunctionInfo{ID: 58, ReturnType: protocol.TypeBool, Parameters: 0}
var funcInfoCmdString = builder.FunctionInfo{ID: 59, ReturnType: protocol.TypeVolatilePointer, Parameters: 0}
var funcInfoCmdPointer = builder.FunctionInfo{ID: 60, ReturnType: protocol.TypeVolatilePointer, Parameters: 0}
var funcInfoCmdVoid3Remapped = builder.FunctionInfo{ID: 61, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoCmdVoidInArrayOfRemapped = builder.FunctionInfo{ID: 62, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidOutArrayOfRemapped = builder.FunctionInfo{ID: 63, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidOutArrayOfUnknownRemapped = builder.FunctionInfo{ID: 64, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdRemapped = builder.FunctionInfo{ID: 65, ReturnType: protocol.TypeUint32, Parameters: 0}

func (c remapped) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U32(uint32(c))
}

var _ = replay.Replayer(&CmdClone{}) // interface compliance check
func (ϟa *CmdClone) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟc.Buf = ϟa.Src.Slice(uint64(uint32(0)), uint64(ϟa.Cnt), ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb).Clone(ϟs)
	ϟb.Push(ϟa.Src.value())
	ϟb.Push(value.U32(ϟa.Cnt))
	ϟb.Call(funcInfoCmdClone)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdMake{}) // interface compliance check
func (ϟa *CmdMake) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟc.Buf = MakeU8ˢ(uint64(ϟa.Cnt), ϟs)
	ϟb.Push(value.U32(ϟa.Cnt))
	ϟb.Call(funcInfoCmdMake)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdCopy{}) // interface compliance check
func (ϟa *CmdCopy) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟc.Buf = MakeU8ˢ(uint64(ϟa.Cnt), ϟs)
	ϟdst, ϟsrc := ϟc.Buf.Copy(ϟa.Src.Slice(uint64(uint32(0)), uint64(ϟa.Cnt), ϟs), ϟs, ϟd, ϟl)
	ϟsrc.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟb.Push(ϟa.Src.value())
	ϟb.Push(value.U32(ϟa.Cnt))
	ϟb.Call(funcInfoCmdCopy)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟdst.onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

var _ = replay.Replayer(&CmdCharsliceToString{}) // interface compliance check
func (ϟa *CmdCharsliceToString) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟc.Str = string(ϟa.S.Slice(uint64(uint32(0)), uint64(ϟa.Len), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb))
	ϟb.Push(ϟa.S.value())
	ϟb.Push(value.U32(ϟa.Len))
	ϟb.Call(funcInfoCmdCharsliceToString)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdCharptrToString{}) // interface compliance check
func (ϟa *CmdCharptrToString) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟc.Str = strings.TrimRight(string(ϟa.S.StringSlice(ϟs, ϟd, ϟl, true).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)), "\x00")
	ϟb.Push(ϟa.S.value())
	ϟb.Call(funcInfoCmdCharptrToString)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdVoid{}) // interface compliance check
func (ϟa *CmdVoid) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoCmdVoid)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdUnknownRet{}) // interface compliance check
func (ϟa *CmdUnknownRet) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoCmdUnknownRet)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdUnknownWritePtr{}) // interface compliance check
func (ϟa *CmdUnknownWritePtr) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.P.value())
	ϟb.Call(funcInfoCmdUnknownWritePtr)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.P.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(ϟa.P.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

var _ = replay.Replayer(&CmdUnknownWriteSlice{}) // interface compliance check
func (ϟa *CmdUnknownWriteSlice) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	count := int32(5)                                        // s32
	slice := ϟa.A.Slice(uint64(int32(0)), uint64(count), ϟs) // Intˢ
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdUnknownWriteSlice)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := int32(int32(0)); i < count; i++ {
		unknown := int64(ϟa.A.Slice(uint64(int32(0)), uint64(count), ϟs).Index(uint64(i), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)) // int
		slice.Index(uint64(i), ϟs).replayWrite(unknown, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = unknown
	}
	_, _ = count, slice
	return nil
}

var _ = replay.Replayer(&CmdVoidU8{}) // interface compliance check
func (ϟa *CmdVoidU8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U8(ϟa.A))
	ϟb.Call(funcInfoCmdVoidU8)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdVoidS8{}) // interface compliance check
func (ϟa *CmdVoidS8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.S8(ϟa.A))
	ϟb.Call(funcInfoCmdVoidS8)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdVoidU16{}) // interface compliance check
func (ϟa *CmdVoidU16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U16(ϟa.A))
	ϟb.Call(funcInfoCmdVoidU16)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdVoidS16{}) // interface compliance check
func (ϟa *CmdVoidS16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.S16(ϟa.A))
	ϟb.Call(funcInfoCmdVoidS16)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdVoidF32{}) // interface compliance check
func (ϟa *CmdVoidF32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.F32(ϟa.A))
	ϟb.Call(funcInfoCmdVoidF32)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdVoidU32{}) // interface compliance check
func (ϟa *CmdVoidU32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U32(ϟa.A))
	ϟb.Call(funcInfoCmdVoidU32)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdVoidS32{}) // interface compliance check
func (ϟa *CmdVoidS32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.S32(ϟa.A))
	ϟb.Call(funcInfoCmdVoidS32)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdVoidF64{}) // interface compliance check
func (ϟa *CmdVoidF64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.F64(ϟa.A))
	ϟb.Call(funcInfoCmdVoidF64)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdVoidU64{}) // interface compliance check
func (ϟa *CmdVoidU64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U64(ϟa.A))
	ϟb.Call(funcInfoCmdVoidU64)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdVoidS64{}) // interface compliance check
func (ϟa *CmdVoidS64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.S64(ϟa.A))
	ϟb.Call(funcInfoCmdVoidS64)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdVoidBool{}) // interface compliance check
func (ϟa *CmdVoidBool) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.Bool(ϟa.A))
	ϟb.Call(funcInfoCmdVoidBool)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdVoidString{}) // interface compliance check
func (ϟa *CmdVoidString) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟb.String(ϟa.A))
	ϟb.Call(funcInfoCmdVoidString)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdVoid3Strings{}) // interface compliance check
func (ϟa *CmdVoid3Strings) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟb.String(ϟa.A))
	ϟb.Push(ϟb.String(ϟa.B))
	ϟb.Push(ϟb.String(ϟa.C))
	ϟb.Call(funcInfoCmdVoid3Strings)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdVoid3InArrays{}) // interface compliance check
func (ϟa *CmdVoid3InArrays) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟc.Buf = MakeU8ˢ(uint64(10), ϟs)
	ϟa.B.Slice(uint64(5), uint64(15), ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.C.Slice(uint64(5), uint64(15), ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟdst, ϟsrc := ϟc.Buf.Copy(ϟa.A.Slice(uint64(5), uint64(25), ϟs), ϟs, ϟd, ϟl)
	ϟsrc.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟb.Push(ϟa.A.value())
	ϟb.Push(ϟa.B.value())
	ϟb.Push(ϟa.C.value())
	ϟb.Call(funcInfoCmdVoid3InArrays)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟdst.onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

var _ = replay.Replayer(&CmdVoidInArrayOfPointers{}) // interface compliance check
func (ϟa *CmdVoidInArrayOfPointers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	slice := ϟa.A.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // Charᵖˢ
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		x := slice.Index(uint64(i), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb).Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb) // char
		_ = x
	}
	ϟb.Push(ϟa.A.value())
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Call(funcInfoCmdVoidInArrayOfPointers)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = slice
	return nil
}

var _ = replay.Replayer(&CmdVoidReadU8{}) // interface compliance check
func (ϟa *CmdVoidReadU8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb) // u8
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidReadU8)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = x
	return nil
}

var _ = replay.Replayer(&CmdVoidReadS8{}) // interface compliance check
func (ϟa *CmdVoidReadS8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb) // s8
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidReadS8)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = x
	return nil
}

var _ = replay.Replayer(&CmdVoidReadU16{}) // interface compliance check
func (ϟa *CmdVoidReadU16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb) // u16
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidReadU16)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = x
	return nil
}

var _ = replay.Replayer(&CmdVoidReadS16{}) // interface compliance check
func (ϟa *CmdVoidReadS16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb) // s16
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidReadS16)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = x
	return nil
}

var _ = replay.Replayer(&CmdVoidReadF32{}) // interface compliance check
func (ϟa *CmdVoidReadF32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb) // f32
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidReadF32)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = x
	return nil
}

var _ = replay.Replayer(&CmdVoidReadU32{}) // interface compliance check
func (ϟa *CmdVoidReadU32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb) // u32
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidReadU32)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = x
	return nil
}

var _ = replay.Replayer(&CmdVoidReadS32{}) // interface compliance check
func (ϟa *CmdVoidReadS32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb) // s32
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidReadS32)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = x
	return nil
}

var _ = replay.Replayer(&CmdVoidReadF64{}) // interface compliance check
func (ϟa *CmdVoidReadF64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb) // f64
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidReadF64)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = x
	return nil
}

var _ = replay.Replayer(&CmdVoidReadU64{}) // interface compliance check
func (ϟa *CmdVoidReadU64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb) // u64
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidReadU64)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = x
	return nil
}

var _ = replay.Replayer(&CmdVoidReadS64{}) // interface compliance check
func (ϟa *CmdVoidReadS64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb) // s64
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidReadS64)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = x
	return nil
}

var _ = replay.Replayer(&CmdVoidReadBool{}) // interface compliance check
func (ϟa *CmdVoidReadBool) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb) // bool
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidReadBool)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = x
	return nil
}

var _ = replay.Replayer(&CmdVoidReadPtrs{}) // interface compliance check
func (ϟa *CmdVoidReadPtrs) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb) // f32
	y := ϟa.B.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb) // u16
	z := ϟa.C.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb) // bool
	ϟb.Push(ϟa.A.value())
	ϟb.Push(ϟa.B.value())
	ϟb.Push(ϟa.C.value())
	ϟb.Call(funcInfoCmdVoidReadPtrs)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = x, y, z
	return nil
}

var _ = replay.Replayer(&CmdVoidWriteU8{}) // interface compliance check
func (ϟa *CmdVoidWriteU8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidWriteU8)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(uint8(1), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

var _ = replay.Replayer(&CmdVoidWriteS8{}) // interface compliance check
func (ϟa *CmdVoidWriteS8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidWriteS8)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(int8(1), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

var _ = replay.Replayer(&CmdVoidWriteU16{}) // interface compliance check
func (ϟa *CmdVoidWriteU16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidWriteU16)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(uint16(1), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

var _ = replay.Replayer(&CmdVoidWriteS16{}) // interface compliance check
func (ϟa *CmdVoidWriteS16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidWriteS16)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(int16(1), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

var _ = replay.Replayer(&CmdVoidWriteF32{}) // interface compliance check
func (ϟa *CmdVoidWriteF32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidWriteF32)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(float32(1), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

var _ = replay.Replayer(&CmdVoidWriteU32{}) // interface compliance check
func (ϟa *CmdVoidWriteU32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidWriteU32)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(uint32(1), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

var _ = replay.Replayer(&CmdVoidWriteS32{}) // interface compliance check
func (ϟa *CmdVoidWriteS32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidWriteS32)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(int32(1), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

var _ = replay.Replayer(&CmdVoidWriteF64{}) // interface compliance check
func (ϟa *CmdVoidWriteF64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidWriteF64)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(float64(1), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

var _ = replay.Replayer(&CmdVoidWriteU64{}) // interface compliance check
func (ϟa *CmdVoidWriteU64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidWriteU64)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(uint64(1), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

var _ = replay.Replayer(&CmdVoidWriteS64{}) // interface compliance check
func (ϟa *CmdVoidWriteS64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidWriteS64)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(int64(1), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

var _ = replay.Replayer(&CmdVoidWriteBool{}) // interface compliance check
func (ϟa *CmdVoidWriteBool) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidWriteBool)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(true, ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

var _ = replay.Replayer(&CmdVoidWritePtrs{}) // interface compliance check
func (ϟa *CmdVoidWritePtrs) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Push(ϟa.B.value())
	ϟb.Push(ϟa.C.value())
	ϟb.Call(funcInfoCmdVoidWritePtrs)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(float32(10), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.B.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(uint16(20), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.C.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(false, ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

var _ = replay.Replayer(&CmdU8{}) // interface compliance check
func (ϟa *CmdU8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoCmdU8)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdS8{}) // interface compliance check
func (ϟa *CmdS8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoCmdS8)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdU16{}) // interface compliance check
func (ϟa *CmdU16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoCmdU16)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdS16{}) // interface compliance check
func (ϟa *CmdS16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoCmdS16)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdF32{}) // interface compliance check
func (ϟa *CmdF32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoCmdF32)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdU32{}) // interface compliance check
func (ϟa *CmdU32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoCmdU32)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdS32{}) // interface compliance check
func (ϟa *CmdS32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoCmdS32)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdF64{}) // interface compliance check
func (ϟa *CmdF64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoCmdF64)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdU64{}) // interface compliance check
func (ϟa *CmdU64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoCmdU64)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdS64{}) // interface compliance check
func (ϟa *CmdS64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoCmdS64)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdBool{}) // interface compliance check
func (ϟa *CmdBool) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoCmdBool)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdString{}) // interface compliance check
func (ϟa *CmdString) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoCmdString)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdPointer{}) // interface compliance check
func (ϟa *CmdPointer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoCmdPointer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdVoid3Remapped{}) // interface compliance check
func (ϟa *CmdVoid3Remapped) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	if key, remap := ϟa.A.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.A.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.A.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.B.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.B.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.B.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.C.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.C.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.C.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoCmdVoid3Remapped)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdVoidInArrayOfRemapped{}) // interface compliance check
func (ϟa *CmdVoidInArrayOfRemapped) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(5), ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidInArrayOfRemapped)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&CmdVoidOutArrayOfRemapped{}) // interface compliance check
func (ϟa *CmdVoidOutArrayOfRemapped) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidOutArrayOfRemapped)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(5), ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

var _ = replay.Replayer(&CmdVoidOutArrayOfUnknownRemapped{}) // interface compliance check
func (ϟa *CmdVoidOutArrayOfUnknownRemapped) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	count := int32(5)                                        // s32
	slice := ϟa.A.Slice(uint64(int32(0)), uint64(count), ϟs) // Remappedˢ
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidOutArrayOfUnknownRemapped)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := int32(int32(0)); i < count; i++ {
		unknown := remapped(ϟa.A.Slice(uint64(int32(0)), uint64(count), ϟs).Index(uint64(i), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)) // remapped
		slice.Index(uint64(i), ϟs).replayWrite(unknown, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = unknown
	}
	_, _ = count, slice
	return nil
}

var _ = replay.Replayer(&CmdRemapped{}) // interface compliance check
func (ϟa *CmdRemapped) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoCmdRemapped)
	if key, remap := ϟa.Result.remap(ϟa, ϟs); remap {
		ptr, found := ϟb.Remappings[key]
		if !found {
			ptr = ϟb.AllocateMemory(uint64(4))
			ϟb.Remappings[key] = ptr
		}
		ϟb.Clone(0)
		ϟb.Store(ptr)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
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
func (p U8ᵖ) replayWrite(value uint8, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p U8ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
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
func (p Charᵖ) replayWrite(value byte, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p Charᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
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
func (p Intᵖ) replayWrite(value int64, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p Intᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
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
func (p U32ᵖ) replayWrite(value uint32, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p U32ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p Charᵖᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Charᵖ {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Charᵖᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Charᵖ {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Charᵖᵖ) replayWrite(value Charᵖ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p Charᵖᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
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
func (p S8ᵖ) replayWrite(value int8, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p S8ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
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
func (p U16ᵖ) replayWrite(value uint16, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p U16ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
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
func (p S16ᵖ) replayWrite(value int16, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p S16ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
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
func (p F32ᵖ) replayWrite(value float32, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p F32ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
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
func (p S32ᵖ) replayWrite(value int32, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p S32ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
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
func (p F64ᵖ) replayWrite(value float64, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p F64ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
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
func (p U64ᵖ) replayWrite(value uint64, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p U64ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
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
func (p S64ᵖ) replayWrite(value int64, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p S64ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
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
func (p Boolᵖ) replayWrite(value bool, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p Boolᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p Voidᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
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
func (p Remappedᵖ) replayWrite(value remapped, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p Remappedᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (s Boolˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Boolˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s Boolˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Boolˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
	return s
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
func (s Charˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Charˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s Charˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Charˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
	return s
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
func (s Charᵖˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Charᵖˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(s.ElementSize(ϟs))
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
			ϟb.Push(v.value())
			ϟb.Store(ptr)
			ptr += step
		}
	}
	return s
}
func (s Charᵖˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Charᵖˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
	return s
}
func (s Charᵖˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
func (s Charᵖˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []Charᵖ {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s F32ˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) F32ˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s F32ˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) F32ˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
	return s
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
func (s F64ˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) F64ˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s F64ˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) F64ˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
	return s
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
func (s Intˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Intˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s Intˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Intˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
	return s
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
func (s Remappedˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Remappedˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(s.ElementSize(ϟs))
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
			if key, remap := v.remap(ϟa, ϟs); remap {
				loadRemap(ϟb, key, protocol.TypeUint32, v.value(ϟb, ϟa, ϟs))
			} else {
				ϟb.Push(v.value(ϟb, ϟa, ϟs))
			}
			ϟb.Store(ptr)
			ptr += step
		}
	}
	return s
}
func (s Remappedˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Remappedˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
		size := s.ElementSize(ϟs)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(size)
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
			if key, remap := v.remap(ϟa, ϟs); remap {
				dst, found := ϟb.Remappings[key]
				if !found {
					dst = ϟb.AllocateMemory(size)
					ϟb.Remappings[key] = dst
				}
				ϟb.Load(protocol.TypeUint32, ptr)
				ϟb.Store(dst)
			}
			ptr += step
		}
	}
	return s
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
func (s S16ˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S16ˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s S16ˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S16ˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
	return s
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
func (s S32ˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S32ˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s S32ˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S32ˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
	return s
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
func (s S64ˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S64ˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s S64ˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S64ˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
	return s
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
func (s S8ˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S8ˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s S8ˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S8ˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
	return s
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
func (s U16ˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U16ˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s U16ˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U16ˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
	return s
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
func (s U32ˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U32ˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s U32ˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U32ˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
	return s
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
func (s U64ˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U64ˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s U64ˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U64ˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
	return s
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
func (s U8ˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U8ˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s U8ˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U8ˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
	return s
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
func (s Voidˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Voidˢ {
	if s.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s Voidˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Voidˢ {
	if s.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root)))
	}
	return s
}
func (s Voidˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root)))
	}
}
