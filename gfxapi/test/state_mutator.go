////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

package test

import (
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
)

func getState(s *gfxapi.State) *State {
	api := API()
	if state, ok := s.APIs[api].(*State); ok {
		return state
	} else {
		if s.APIs == nil {
			s.APIs = make(map[gfxapi.API]interface{})
		}
		state = &State{}
		state.Init()
		s.APIs[api] = state
		return state
	}
}

func (ϟa *CmdClone) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟc.U8s = ϟa.Src.Slice(uint64(uint32(0)), uint64(ϟa.Cnt), ϟs).Clone(ϟs)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *CmdMake) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟc.U8s = MakeU8ˢ(uint64(ϟa.Cnt), ϟs)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *CmdCopy) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟc.U8s = MakeU8ˢ(uint64(ϟa.Cnt), ϟs)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟc.U8s.Copy(ϟa.Src.Slice(uint64(uint32(0)), uint64(ϟa.Cnt), ϟs), ϟs, ϟd, ϟl)
	return nil
}
func (ϟa *CmdCharsliceToString) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟc.Str = string(ϟa.S.Slice(uint64(uint32(0)), uint64(ϟa.Len), ϟs).Read(ϟs, ϟd, ϟl))
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *CmdCharptrToString) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟc.Str = string(ϟa.S.StringSlice(ϟs, ϟd, ϟl, false).Read(ϟs, ϟd, ϟl))
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *CmdSliceCasts) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟc.U8s = AsU8ˢ(ϟa.S.Slice(uint64(uint32(0)), uint64(ϟa.L), ϟs), ϟs)
	ϟc.U16s = ϟa.S.Slice(uint64(uint32(0)), uint64(ϟa.L), ϟs)
	ϟc.U32s = AsU32ˢ(ϟa.S.Slice(uint64(uint32(0)), uint64(ϟa.L), ϟs), ϟs)
	ϟc.Ints = AsIntˢ(ϟa.S.Slice(uint64(uint32(0)), uint64(ϟa.L), ϟs), ϟs)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *CmdVoid) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *CmdUnknownRet) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	return nil
}
func (ϟa *CmdUnknownWritePtr) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.P.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.P.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl), ϟs)
	return nil
}
func (ϟa *CmdUnknownWriteSlice) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	count := int32(5)                                        // s32
	slice := ϟa.A.Slice(uint64(int32(0)), uint64(count), ϟs) // Intˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := int32(int32(0)); i < count; i++ {
		unknown := int64(ϟa.A.Slice(uint64(int32(0)), uint64(count), ϟs).Index(uint64(i), ϟs).Read(ϟs, ϟd, ϟl)) // int
		slice.Index(uint64(i), ϟs).Write(unknown, ϟs)
		_ = unknown
	}
	_, _ = count, slice
	return nil
}
func (ϟa *CmdVoidU8) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *CmdVoidS8) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *CmdVoidU16) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *CmdVoidS16) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *CmdVoidF32) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *CmdVoidU32) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *CmdVoidS32) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *CmdVoidF64) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *CmdVoidU64) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *CmdVoidS64) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *CmdVoidBool) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *CmdVoidString) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *CmdVoid3Strings) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *CmdVoid3InArrays) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟc.U8s = MakeU8ˢ(uint64(10), ϟs)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟc.U8s.Copy(ϟa.A.Slice(uint64(5), uint64(25), ϟs), ϟs, ϟd, ϟl)
	return nil
}
func (ϟa *CmdVoidInArrayOfPointers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	slice := ϟa.A.Slice(uint64(int32(0)), uint64(ϟa.Count), ϟs) // Charᵖˢ
	for i := int32(int32(0)); i < ϟa.Count; i++ {
		x := slice.Index(uint64(i), ϟs).Read(ϟs, ϟd, ϟl).Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl) // char
		_ = x
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = slice
	return nil
}
func (ϟa *CmdVoidReadU8) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl) // u8
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = x
	return nil
}
func (ϟa *CmdVoidReadS8) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl) // s8
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = x
	return nil
}
func (ϟa *CmdVoidReadU16) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl) // u16
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = x
	return nil
}
func (ϟa *CmdVoidReadS16) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl) // s16
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = x
	return nil
}
func (ϟa *CmdVoidReadF32) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl) // f32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = x
	return nil
}
func (ϟa *CmdVoidReadU32) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = x
	return nil
}
func (ϟa *CmdVoidReadS32) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl) // s32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = x
	return nil
}
func (ϟa *CmdVoidReadF64) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl) // f64
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = x
	return nil
}
func (ϟa *CmdVoidReadU64) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl) // u64
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = x
	return nil
}
func (ϟa *CmdVoidReadS64) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl) // s64
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = x
	return nil
}
func (ϟa *CmdVoidReadBool) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl) // bool
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = x
	return nil
}
func (ϟa *CmdVoidReadPtrs) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	x := ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl) // f32
	y := ϟa.B.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl) // u16
	z := ϟa.C.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟs, ϟd, ϟl) // bool
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = x, y, z
	return nil
}
func (ϟa *CmdVoidWriteU8) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(uint8(1), ϟs)
	return nil
}
func (ϟa *CmdVoidWriteS8) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(int8(1), ϟs)
	return nil
}
func (ϟa *CmdVoidWriteU16) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(uint16(1), ϟs)
	return nil
}
func (ϟa *CmdVoidWriteS16) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(int16(1), ϟs)
	return nil
}
func (ϟa *CmdVoidWriteF32) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(float32(1), ϟs)
	return nil
}
func (ϟa *CmdVoidWriteU32) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(uint32(1), ϟs)
	return nil
}
func (ϟa *CmdVoidWriteS32) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(int32(1), ϟs)
	return nil
}
func (ϟa *CmdVoidWriteF64) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(float64(1), ϟs)
	return nil
}
func (ϟa *CmdVoidWriteU64) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(uint64(1), ϟs)
	return nil
}
func (ϟa *CmdVoidWriteS64) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(int64(1), ϟs)
	return nil
}
func (ϟa *CmdVoidWriteBool) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(true, ϟs)
	return nil
}
func (ϟa *CmdVoidWritePtrs) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.A.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(float32(10), ϟs)
	ϟa.B.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(uint16(20), ϟs)
	ϟa.C.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(false, ϟs)
	return nil
}
func (ϟa *CmdU8) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = uint8(0)
	return nil
}
func (ϟa *CmdS8) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = int8(0)
	return nil
}
func (ϟa *CmdU16) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = uint16(0)
	return nil
}
func (ϟa *CmdS16) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = int16(0)
	return nil
}
func (ϟa *CmdF32) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = float32(0)
	return nil
}
func (ϟa *CmdU32) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = uint32(0)
	return nil
}
func (ϟa *CmdS32) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = int32(0)
	return nil
}
func (ϟa *CmdF64) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = float64(0)
	return nil
}
func (ϟa *CmdU64) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = uint64(0)
	return nil
}
func (ϟa *CmdS64) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = int64(0)
	return nil
}
func (ϟa *CmdBool) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = false
	return nil
}
func (ϟa *CmdString) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ""
	return nil
}
func (ϟa *CmdPointer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	return nil
}
func (ϟa *CmdVoid3Remapped) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *CmdVoidInArrayOfRemapped) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *CmdVoidOutArrayOfRemapped) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *CmdVoidOutArrayOfUnknownRemapped) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	count := int32(5)                                        // s32
	slice := ϟa.A.Slice(uint64(int32(0)), uint64(count), ϟs) // Remappedˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := int32(int32(0)); i < count; i++ {
		unknown := remapped(ϟa.A.Slice(uint64(int32(0)), uint64(count), ϟs).Index(uint64(i), ϟs).Read(ϟs, ϟd, ϟl)) // remapped
		slice.Index(uint64(i), ϟs).Write(unknown, ϟs)
		_ = unknown
	}
	_, _ = count, slice
	return nil
}
func (ϟa *CmdRemapped) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	return nil
}
