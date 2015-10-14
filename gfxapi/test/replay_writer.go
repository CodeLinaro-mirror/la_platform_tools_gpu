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
var funcInfoCmdSliceCasts = builder.FunctionInfo{ID: 5, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoCmdVoid = builder.FunctionInfo{ID: 6, ReturnType: protocol.TypeVoid, Parameters: 0}
var funcInfoCmdUnknownRet = builder.FunctionInfo{ID: 7, ReturnType: protocol.TypeInt64, Parameters: 0}
var funcInfoCmdUnknownWritePtr = builder.FunctionInfo{ID: 8, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdUnknownWriteSlice = builder.FunctionInfo{ID: 9, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidU8 = builder.FunctionInfo{ID: 10, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidS8 = builder.FunctionInfo{ID: 11, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidU16 = builder.FunctionInfo{ID: 12, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidS16 = builder.FunctionInfo{ID: 13, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidF32 = builder.FunctionInfo{ID: 14, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidU32 = builder.FunctionInfo{ID: 15, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidS32 = builder.FunctionInfo{ID: 16, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidF64 = builder.FunctionInfo{ID: 17, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidU64 = builder.FunctionInfo{ID: 18, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidS64 = builder.FunctionInfo{ID: 19, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidBool = builder.FunctionInfo{ID: 20, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidString = builder.FunctionInfo{ID: 21, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoid3Strings = builder.FunctionInfo{ID: 22, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoCmdVoid3InArrays = builder.FunctionInfo{ID: 23, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoCmdVoidInArrayOfStrings = builder.FunctionInfo{ID: 24, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoCmdVoidReadU8 = builder.FunctionInfo{ID: 25, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidReadS8 = builder.FunctionInfo{ID: 26, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidReadU16 = builder.FunctionInfo{ID: 27, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidReadS16 = builder.FunctionInfo{ID: 28, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidReadF32 = builder.FunctionInfo{ID: 29, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidReadU32 = builder.FunctionInfo{ID: 30, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidReadS32 = builder.FunctionInfo{ID: 31, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidReadF64 = builder.FunctionInfo{ID: 32, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidReadU64 = builder.FunctionInfo{ID: 33, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidReadS64 = builder.FunctionInfo{ID: 34, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidReadBool = builder.FunctionInfo{ID: 35, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidReadPtrs = builder.FunctionInfo{ID: 36, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoCmdVoidWriteU8 = builder.FunctionInfo{ID: 37, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidWriteS8 = builder.FunctionInfo{ID: 38, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidWriteU16 = builder.FunctionInfo{ID: 39, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidWriteS16 = builder.FunctionInfo{ID: 40, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidWriteF32 = builder.FunctionInfo{ID: 41, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidWriteU32 = builder.FunctionInfo{ID: 42, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidWriteS32 = builder.FunctionInfo{ID: 43, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidWriteF64 = builder.FunctionInfo{ID: 44, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidWriteU64 = builder.FunctionInfo{ID: 45, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidWriteS64 = builder.FunctionInfo{ID: 46, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidWriteBool = builder.FunctionInfo{ID: 47, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidWritePtrs = builder.FunctionInfo{ID: 48, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoCmdU8 = builder.FunctionInfo{ID: 49, ReturnType: protocol.TypeUint8, Parameters: 0}
var funcInfoCmdS8 = builder.FunctionInfo{ID: 50, ReturnType: protocol.TypeInt8, Parameters: 0}
var funcInfoCmdU16 = builder.FunctionInfo{ID: 51, ReturnType: protocol.TypeUint16, Parameters: 0}
var funcInfoCmdS16 = builder.FunctionInfo{ID: 52, ReturnType: protocol.TypeInt16, Parameters: 0}
var funcInfoCmdF32 = builder.FunctionInfo{ID: 53, ReturnType: protocol.TypeFloat, Parameters: 0}
var funcInfoCmdU32 = builder.FunctionInfo{ID: 54, ReturnType: protocol.TypeUint32, Parameters: 0}
var funcInfoCmdS32 = builder.FunctionInfo{ID: 55, ReturnType: protocol.TypeInt32, Parameters: 0}
var funcInfoCmdF64 = builder.FunctionInfo{ID: 56, ReturnType: protocol.TypeDouble, Parameters: 0}
var funcInfoCmdU64 = builder.FunctionInfo{ID: 57, ReturnType: protocol.TypeUint64, Parameters: 0}
var funcInfoCmdS64 = builder.FunctionInfo{ID: 58, ReturnType: protocol.TypeInt64, Parameters: 0}
var funcInfoCmdBool = builder.FunctionInfo{ID: 59, ReturnType: protocol.TypeBool, Parameters: 0}
var funcInfoCmdString = builder.FunctionInfo{ID: 60, ReturnType: protocol.TypeVolatilePointer, Parameters: 0}
var funcInfoCmdPointer = builder.FunctionInfo{ID: 61, ReturnType: protocol.TypeAbsolutePointer, Parameters: 0}
var funcInfoCmdVoid3Remapped = builder.FunctionInfo{ID: 62, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoCmdVoidInArrayOfRemapped = builder.FunctionInfo{ID: 63, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidOutArrayOfRemapped = builder.FunctionInfo{ID: 64, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdVoidOutArrayOfUnknownRemapped = builder.FunctionInfo{ID: 65, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCmdRemapped = builder.FunctionInfo{ID: 66, ReturnType: protocol.TypeUint32, Parameters: 0}

func (c remapped) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U32(uint32(c))
}

var _ = replay.Replayer(&CmdClone{}) // interface compliance check
// Replay emits the replay instructions to call cmdClone(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdClone) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟc.U8s = ϟa.Src.Slice(uint64(uint32(0)), uint64(ϟa.Cnt), ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdClone().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdClone) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.Src.value())
	ϟb.Push(value.U32(ϟa.Cnt))
	ϟb.Call(funcInfoCmdClone)
}

var _ = replay.Replayer(&CmdMake{}) // interface compliance check
// Replay emits the replay instructions to call cmdMake(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdMake) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟc.U8s = MakeU8ˢ(uint64(ϟa.Cnt), ϟs)
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdMake().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdMake) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(value.U32(ϟa.Cnt))
	ϟb.Call(funcInfoCmdMake)
}

var _ = replay.Replayer(&CmdCopy{}) // interface compliance check
// Replay emits the replay instructions to call cmdCopy(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdCopy) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟc.U8s = MakeU8ˢ(uint64(ϟa.Cnt), ϟs)
	ϟdst, ϟsrc := ϟc.U8s, ϟa.Src.Slice(uint64(uint32(0)), uint64(ϟa.Cnt), ϟs)
	ϟcount := min(ϟdst.Count, ϟsrc.Count)
	ϟdst, ϟsrc = ϟdst.Slice(0, ϟcount, ϟs), ϟsrc.Slice(0, ϟcount, ϟs)
	ϟsrcElems := ϟsrc.Read(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟdst.Write(ϟsrcElems, ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdCopy().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdCopy) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.Src.value())
	ϟb.Push(value.U32(ϟa.Cnt))
	ϟb.Call(funcInfoCmdCopy)
}

var _ = replay.Replayer(&CmdCharsliceToString{}) // interface compliance check
// Replay emits the replay instructions to call cmdCharsliceToString(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdCharsliceToString) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟc.Str = string(ϟa.S.Slice(uint64(uint32(0)), uint64(ϟa.Len), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdCharsliceToString().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdCharsliceToString) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.S.value())
	ϟb.Push(value.U32(ϟa.Len))
	ϟb.Call(funcInfoCmdCharsliceToString)
}

var _ = replay.Replayer(&CmdCharptrToString{}) // interface compliance check
// Replay emits the replay instructions to call cmdCharptrToString(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdCharptrToString) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟc.Str = strings.TrimRight(string(ϟa.S.StringSlice(ϟs, ϟd, ϟl).Read(ϟa, ϟs, ϟd, ϟl, ϟb)), "\x00")
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdCharptrToString().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdCharptrToString) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.S.value())
	ϟb.Call(funcInfoCmdCharptrToString)
}

var _ = replay.Replayer(&CmdSliceCasts{}) // interface compliance check
// Replay emits the replay instructions to call cmdSliceCasts(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdSliceCasts) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟc.U8s = AsU8ˢ(ϟa.S.Slice(uint64(uint32(0)), uint64(ϟa.L), ϟs), ϟs)
	ϟc.U16s = ϟa.S.Slice(uint64(uint32(0)), uint64(ϟa.L), ϟs)
	ϟc.U32s = AsU32ˢ(ϟa.S.Slice(uint64(uint32(0)), uint64(ϟa.L), ϟs), ϟs)
	ϟc.Ints = AsIntˢ(ϟa.S.Slice(uint64(uint32(0)), uint64(ϟa.L), ϟs), ϟs)
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdSliceCasts().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdSliceCasts) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.S.value())
	ϟb.Push(value.U32(ϟa.L))
	ϟb.Call(funcInfoCmdSliceCasts)
}

var _ = replay.Replayer(&CmdVoid{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoid(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoid) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoid().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoid) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Call(funcInfoCmdVoid)
}

var _ = replay.Replayer(&CmdUnknownRet{}) // interface compliance check
// Replay emits the replay instructions to call cmdUnknownRet(), and performs the
// necessary state-mutation and memory observations to ϟs.
// The cmdUnknownRet() return value will be stored on the stack.
func (ϟa *CmdUnknownRet) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdUnknownRet().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
// The cmdUnknownRet() return value will be stored on the stack.
func (ϟa *CmdUnknownRet) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Call(funcInfoCmdUnknownRet)
}

var _ = replay.Replayer(&CmdUnknownWritePtr{}) // interface compliance check
// Replay emits the replay instructions to call cmdUnknownWritePtr(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdUnknownWritePtr) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.P.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.P.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdUnknownWritePtr().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdUnknownWritePtr) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.P.value())
	ϟb.Call(funcInfoCmdUnknownWritePtr)
}

var _ = replay.Replayer(&CmdUnknownWriteSlice{}) // interface compliance check
// Replay emits the replay instructions to call cmdUnknownWriteSlice(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdUnknownWriteSlice) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	count := int32(5)                                        // s32
	slice := ϟa.A.Slice(uint64(int32(0)), uint64(count), ϟs) // Intˢ
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := int32(int32(0)); i < count; i++ {
		unknown := int64(ϟa.A.Slice(uint64(int32(0)), uint64(count), ϟs).Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // int
		slice.Index(uint64(i), ϟs).Write(unknown, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = unknown
	}
	_, _ = count, slice
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdUnknownWriteSlice().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdUnknownWriteSlice) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdUnknownWriteSlice)
}

var _ = replay.Replayer(&CmdVoidU8{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidU8(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidU8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidU8().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidU8) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(value.U8(ϟa.A))
	ϟb.Call(funcInfoCmdVoidU8)
}

var _ = replay.Replayer(&CmdVoidS8{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidS8(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidS8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidS8().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidS8) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(value.S8(ϟa.A))
	ϟb.Call(funcInfoCmdVoidS8)
}

var _ = replay.Replayer(&CmdVoidU16{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidU16(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidU16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidU16().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidU16) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(value.U16(ϟa.A))
	ϟb.Call(funcInfoCmdVoidU16)
}

var _ = replay.Replayer(&CmdVoidS16{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidS16(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidS16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidS16().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidS16) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(value.S16(ϟa.A))
	ϟb.Call(funcInfoCmdVoidS16)
}

var _ = replay.Replayer(&CmdVoidF32{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidF32(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidF32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidF32().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidF32) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(value.F32(ϟa.A))
	ϟb.Call(funcInfoCmdVoidF32)
}

var _ = replay.Replayer(&CmdVoidU32{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidU32(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidU32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidU32().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidU32) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(value.U32(ϟa.A))
	ϟb.Call(funcInfoCmdVoidU32)
}

var _ = replay.Replayer(&CmdVoidS32{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidS32(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidS32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidS32().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidS32) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(value.S32(ϟa.A))
	ϟb.Call(funcInfoCmdVoidS32)
}

var _ = replay.Replayer(&CmdVoidF64{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidF64(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidF64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidF64().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidF64) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(value.F64(ϟa.A))
	ϟb.Call(funcInfoCmdVoidF64)
}

var _ = replay.Replayer(&CmdVoidU64{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidU64(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidU64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidU64().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidU64) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(value.U64(ϟa.A))
	ϟb.Call(funcInfoCmdVoidU64)
}

var _ = replay.Replayer(&CmdVoidS64{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidS64(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidS64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidS64().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidS64) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(value.S64(ϟa.A))
	ϟb.Call(funcInfoCmdVoidS64)
}

var _ = replay.Replayer(&CmdVoidBool{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidBool(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidBool) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidBool().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidBool) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(value.Bool(ϟa.A))
	ϟb.Call(funcInfoCmdVoidBool)
}

var _ = replay.Replayer(&CmdVoidString{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidString(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidString) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidString().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidString) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟb.String(ϟa.A))
	ϟb.Call(funcInfoCmdVoidString)
}

var _ = replay.Replayer(&CmdVoid3Strings{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoid3Strings(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoid3Strings) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoid3Strings().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoid3Strings) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟb.String(ϟa.A))
	ϟb.Push(ϟb.String(ϟa.B))
	ϟb.Push(ϟb.String(ϟa.C))
	ϟb.Call(funcInfoCmdVoid3Strings)
}

var _ = replay.Replayer(&CmdVoid3InArrays{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoid3InArrays(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoid3InArrays) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟc.U8s = MakeU8ˢ(uint64(10), ϟs)
	ϟa.B.Slice(uint64(5), uint64(15), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.C.Slice(uint64(5), uint64(15), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟdst, ϟsrc := ϟc.U8s, ϟa.A.Slice(uint64(5), uint64(25), ϟs)
	ϟcount := min(ϟdst.Count, ϟsrc.Count)
	ϟdst, ϟsrc = ϟdst.Slice(0, ϟcount, ϟs), ϟsrc.Slice(0, ϟcount, ϟs)
	ϟsrcElems := ϟsrc.Read(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟdst.Write(ϟsrcElems, ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoid3InArrays().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoid3InArrays) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.A.value())
	ϟb.Push(ϟa.B.value())
	ϟb.Push(ϟa.C.value())
	ϟb.Call(funcInfoCmdVoid3InArrays)
}

var _ = replay.Replayer(&CmdVoidInArrayOfStrings{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidInArrayOfStrings(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidInArrayOfStrings) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	names := ϟa.Strings.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // Charᶜᵖˢ
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		name := strings.TrimRight(string(Charᵖ(names.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb)).StringSlice(ϟs, ϟd, ϟl).Read(ϟa, ϟs, ϟd, ϟl, ϟb)), "\x00") // string
		_ = name
	}
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = names
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidInArrayOfStrings().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidInArrayOfStrings) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.Strings.value())
	ϟb.Push(value.S32(ϟa.Count))
	ϟb.Call(funcInfoCmdVoidInArrayOfStrings)
}

var _ = replay.Replayer(&CmdVoidReadU8{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidReadU8(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidReadU8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb) // u8
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = x
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidReadU8().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidReadU8) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidReadU8)
}

var _ = replay.Replayer(&CmdVoidReadS8{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidReadS8(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidReadS8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb) // s8
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = x
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidReadS8().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidReadS8) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidReadS8)
}

var _ = replay.Replayer(&CmdVoidReadU16{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidReadU16(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidReadU16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb) // u16
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = x
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidReadU16().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidReadU16) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidReadU16)
}

var _ = replay.Replayer(&CmdVoidReadS16{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidReadS16(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidReadS16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb) // s16
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = x
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidReadS16().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidReadS16) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidReadS16)
}

var _ = replay.Replayer(&CmdVoidReadF32{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidReadF32(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidReadF32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb) // f32
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = x
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidReadF32().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidReadF32) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidReadF32)
}

var _ = replay.Replayer(&CmdVoidReadU32{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidReadU32(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidReadU32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb) // u32
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = x
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidReadU32().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidReadU32) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidReadU32)
}

var _ = replay.Replayer(&CmdVoidReadS32{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidReadS32(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidReadS32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb) // s32
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = x
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidReadS32().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidReadS32) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidReadS32)
}

var _ = replay.Replayer(&CmdVoidReadF64{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidReadF64(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidReadF64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb) // f64
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = x
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidReadF64().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidReadF64) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidReadF64)
}

var _ = replay.Replayer(&CmdVoidReadU64{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidReadU64(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidReadU64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb) // u64
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = x
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidReadU64().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidReadU64) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidReadU64)
}

var _ = replay.Replayer(&CmdVoidReadS64{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidReadS64(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidReadS64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb) // s64
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = x
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidReadS64().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidReadS64) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidReadS64)
}

var _ = replay.Replayer(&CmdVoidReadBool{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidReadBool(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidReadBool) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb) // bool
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = x
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidReadBool().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidReadBool) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidReadBool)
}

var _ = replay.Replayer(&CmdVoidReadPtrs{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidReadPtrs(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidReadPtrs) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb) // f32
	y := ϟa.B.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb) // u16
	z := ϟa.C.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb) // bool
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = x, y, z
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidReadPtrs().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidReadPtrs) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.A.value())
	ϟb.Push(ϟa.B.value())
	ϟb.Push(ϟa.C.value())
	ϟb.Call(funcInfoCmdVoidReadPtrs)
}

var _ = replay.Replayer(&CmdVoidWriteU8{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidWriteU8(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidWriteU8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(uint8(1), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidWriteU8().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidWriteU8) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidWriteU8)
}

var _ = replay.Replayer(&CmdVoidWriteS8{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidWriteS8(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidWriteS8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(int8(1), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidWriteS8().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidWriteS8) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidWriteS8)
}

var _ = replay.Replayer(&CmdVoidWriteU16{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidWriteU16(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidWriteU16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(uint16(1), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidWriteU16().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidWriteU16) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidWriteU16)
}

var _ = replay.Replayer(&CmdVoidWriteS16{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidWriteS16(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidWriteS16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(int16(1), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidWriteS16().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidWriteS16) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidWriteS16)
}

var _ = replay.Replayer(&CmdVoidWriteF32{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidWriteF32(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidWriteF32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(float32(1), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidWriteF32().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidWriteF32) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidWriteF32)
}

var _ = replay.Replayer(&CmdVoidWriteU32{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidWriteU32(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidWriteU32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(uint32(1), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidWriteU32().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidWriteU32) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidWriteU32)
}

var _ = replay.Replayer(&CmdVoidWriteS32{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidWriteS32(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidWriteS32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(int32(1), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidWriteS32().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidWriteS32) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidWriteS32)
}

var _ = replay.Replayer(&CmdVoidWriteF64{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidWriteF64(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidWriteF64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(float64(1), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidWriteF64().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidWriteF64) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidWriteF64)
}

var _ = replay.Replayer(&CmdVoidWriteU64{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidWriteU64(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidWriteU64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(uint64(1), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidWriteU64().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidWriteU64) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidWriteU64)
}

var _ = replay.Replayer(&CmdVoidWriteS64{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidWriteS64(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidWriteS64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(int64(1), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidWriteS64().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidWriteS64) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidWriteS64)
}

var _ = replay.Replayer(&CmdVoidWriteBool{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidWriteBool(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidWriteBool) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(true, ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidWriteBool().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidWriteBool) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidWriteBool)
}

var _ = replay.Replayer(&CmdVoidWritePtrs{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidWritePtrs(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidWritePtrs) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(float32(10), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.B.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(uint16(20), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.C.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(false, ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidWritePtrs().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidWritePtrs) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.A.value())
	ϟb.Push(ϟa.B.value())
	ϟb.Push(ϟa.C.value())
	ϟb.Call(funcInfoCmdVoidWritePtrs)
}

var _ = replay.Replayer(&CmdU8{}) // interface compliance check
// Replay emits the replay instructions to call cmdU8(), and performs the
// necessary state-mutation and memory observations to ϟs.
// The cmdU8() return value will be stored on the stack.
func (ϟa *CmdU8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdU8().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
// The cmdU8() return value will be stored on the stack.
func (ϟa *CmdU8) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Call(funcInfoCmdU8)
}

var _ = replay.Replayer(&CmdS8{}) // interface compliance check
// Replay emits the replay instructions to call cmdS8(), and performs the
// necessary state-mutation and memory observations to ϟs.
// The cmdS8() return value will be stored on the stack.
func (ϟa *CmdS8) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdS8().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
// The cmdS8() return value will be stored on the stack.
func (ϟa *CmdS8) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Call(funcInfoCmdS8)
}

var _ = replay.Replayer(&CmdU16{}) // interface compliance check
// Replay emits the replay instructions to call cmdU16(), and performs the
// necessary state-mutation and memory observations to ϟs.
// The cmdU16() return value will be stored on the stack.
func (ϟa *CmdU16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdU16().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
// The cmdU16() return value will be stored on the stack.
func (ϟa *CmdU16) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Call(funcInfoCmdU16)
}

var _ = replay.Replayer(&CmdS16{}) // interface compliance check
// Replay emits the replay instructions to call cmdS16(), and performs the
// necessary state-mutation and memory observations to ϟs.
// The cmdS16() return value will be stored on the stack.
func (ϟa *CmdS16) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdS16().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
// The cmdS16() return value will be stored on the stack.
func (ϟa *CmdS16) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Call(funcInfoCmdS16)
}

var _ = replay.Replayer(&CmdF32{}) // interface compliance check
// Replay emits the replay instructions to call cmdF32(), and performs the
// necessary state-mutation and memory observations to ϟs.
// The cmdF32() return value will be stored on the stack.
func (ϟa *CmdF32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdF32().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
// The cmdF32() return value will be stored on the stack.
func (ϟa *CmdF32) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Call(funcInfoCmdF32)
}

var _ = replay.Replayer(&CmdU32{}) // interface compliance check
// Replay emits the replay instructions to call cmdU32(), and performs the
// necessary state-mutation and memory observations to ϟs.
// The cmdU32() return value will be stored on the stack.
func (ϟa *CmdU32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdU32().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
// The cmdU32() return value will be stored on the stack.
func (ϟa *CmdU32) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Call(funcInfoCmdU32)
}

var _ = replay.Replayer(&CmdS32{}) // interface compliance check
// Replay emits the replay instructions to call cmdS32(), and performs the
// necessary state-mutation and memory observations to ϟs.
// The cmdS32() return value will be stored on the stack.
func (ϟa *CmdS32) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdS32().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
// The cmdS32() return value will be stored on the stack.
func (ϟa *CmdS32) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Call(funcInfoCmdS32)
}

var _ = replay.Replayer(&CmdF64{}) // interface compliance check
// Replay emits the replay instructions to call cmdF64(), and performs the
// necessary state-mutation and memory observations to ϟs.
// The cmdF64() return value will be stored on the stack.
func (ϟa *CmdF64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdF64().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
// The cmdF64() return value will be stored on the stack.
func (ϟa *CmdF64) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Call(funcInfoCmdF64)
}

var _ = replay.Replayer(&CmdU64{}) // interface compliance check
// Replay emits the replay instructions to call cmdU64(), and performs the
// necessary state-mutation and memory observations to ϟs.
// The cmdU64() return value will be stored on the stack.
func (ϟa *CmdU64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdU64().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
// The cmdU64() return value will be stored on the stack.
func (ϟa *CmdU64) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Call(funcInfoCmdU64)
}

var _ = replay.Replayer(&CmdS64{}) // interface compliance check
// Replay emits the replay instructions to call cmdS64(), and performs the
// necessary state-mutation and memory observations to ϟs.
// The cmdS64() return value will be stored on the stack.
func (ϟa *CmdS64) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdS64().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
// The cmdS64() return value will be stored on the stack.
func (ϟa *CmdS64) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Call(funcInfoCmdS64)
}

var _ = replay.Replayer(&CmdBool{}) // interface compliance check
// Replay emits the replay instructions to call cmdBool(), and performs the
// necessary state-mutation and memory observations to ϟs.
// The cmdBool() return value will be stored on the stack.
func (ϟa *CmdBool) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdBool().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
// The cmdBool() return value will be stored on the stack.
func (ϟa *CmdBool) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Call(funcInfoCmdBool)
}

var _ = replay.Replayer(&CmdString{}) // interface compliance check
// Replay emits the replay instructions to call cmdString(), and performs the
// necessary state-mutation and memory observations to ϟs.
// The cmdString() return value will be stored on the stack.
func (ϟa *CmdString) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdString().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
// The cmdString() return value will be stored on the stack.
func (ϟa *CmdString) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Call(funcInfoCmdString)
}

var _ = replay.Replayer(&CmdPointer{}) // interface compliance check
// Replay emits the replay instructions to call cmdPointer(), and performs the
// necessary state-mutation and memory observations to ϟs.
// The cmdPointer() return value will be stored on the stack.
func (ϟa *CmdPointer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdPointer().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
// The cmdPointer() return value will be stored on the stack.
func (ϟa *CmdPointer) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Call(funcInfoCmdPointer)
}

var _ = replay.Replayer(&CmdVoid3Remapped{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoid3Remapped(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoid3Remapped) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoid3Remapped().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoid3Remapped) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
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
}

var _ = replay.Replayer(&CmdVoidInArrayOfRemapped{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidInArrayOfRemapped(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidInArrayOfRemapped) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(5), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidInArrayOfRemapped().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidInArrayOfRemapped) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidInArrayOfRemapped)
}

var _ = replay.Replayer(&CmdVoidOutArrayOfRemapped{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidOutArrayOfRemapped(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidOutArrayOfRemapped) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(5), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidOutArrayOfRemapped().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidOutArrayOfRemapped) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidOutArrayOfRemapped)
}

var _ = replay.Replayer(&CmdVoidOutArrayOfUnknownRemapped{}) // interface compliance check
// Replay emits the replay instructions to call cmdVoidOutArrayOfUnknownRemapped(), and performs the
// necessary state-mutation and memory observations to ϟs.
func (ϟa *CmdVoidOutArrayOfUnknownRemapped) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	count := int32(5)                                        // s32
	slice := ϟa.A.Slice(uint64(int32(0)), uint64(count), ϟs) // Remappedˢ
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := int32(int32(0)); i < count; i++ {
		unknown := remapped(ϟa.A.Slice(uint64(int32(0)), uint64(count), ϟs).Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // remapped
		slice.Index(uint64(i), ϟs).Write(unknown, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = unknown
	}
	_, _ = count, slice
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdVoidOutArrayOfUnknownRemapped().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
func (ϟa *CmdVoidOutArrayOfUnknownRemapped) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	ϟb.Push(ϟa.A.value())
	ϟb.Call(funcInfoCmdVoidOutArrayOfUnknownRemapped)
}

var _ = replay.Replayer(&CmdRemapped{}) // interface compliance check
// Replay emits the replay instructions to call cmdRemapped(), and performs the
// necessary state-mutation and memory observations to ϟs.
// The cmdRemapped() return value will be stored on the stack.
func (ϟa *CmdRemapped) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (ϟe error) {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.Call(ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

// Call builds the replay instructions to push the arguments to the stack and invoke cmdRemapped().
// Unlike Replay(), Call() does not perform any state-mutation or memory observations to ϟs.
// The cmdRemapped() return value will be stored on the stack.
func (ϟa *CmdRemapped) Call(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
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
}
func (p U8ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p U16ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p U32ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p Intᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p Charᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p Charᶜᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p Charᶜᵖᶜᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p Charᶜᵖᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p S8ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p S16ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p F32ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p S32ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p F64ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p U64ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p S64ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
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
func (p Remappedᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
