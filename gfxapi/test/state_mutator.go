////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

package test

import (
	"fmt"
	"log"
	"reflect"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/gfxapi/state"
)

func getState(a atom.Atom, s *state.State) *State {
	id := a.ContextID()
	if state, ok := s.Contexts[id].(*State); ok {
		return state
	} else {
		panic(fmt.Errorf("State for atom %T with context id %d was %T, expected *gfxapi_test.State",
			a, id, s.Contexts[id]))
	}
}

func (ϟa *CmdVoid) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdVoid_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_void expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdVoidU8) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdVoidU8_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_void_u8 expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdVoidS8) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdVoidS8_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_void_s8 expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdVoidU16) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdVoidU16_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_void_u16 expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdVoidS16) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdVoidS16_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_void_s16 expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdVoidF32) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdVoidF32_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_void_f32 expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdVoidU32) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdVoidU32_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_void_u32 expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdVoidS32) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdVoidS32_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_void_s32 expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdVoidF64) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdVoidF64_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_void_f64 expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdVoidU64) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdVoidU64_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_void_u64 expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdVoidS64) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdVoidS64_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_void_s64 expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdVoidBool) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdVoidBool_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_void_bool expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdVoidString) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdVoidString_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_void_string expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdVoid3Strings) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdVoid3Strings_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_void_3_strings expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdVoid3Arrays) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdVoid3Arrays_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_void_3_arrays expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdVoidArrayOfStrings) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdVoidArrayOfStrings_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_void_array_of_strings expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdU8) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdU8_Out{}
	ϟo.Result = uint8(0)
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_u8 expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdS8) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdS8_Out{}
	ϟo.Result = int8(0)
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_s8 expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdU16) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdU16_Out{}
	ϟo.Result = uint16(0)
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_u16 expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdS16) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdS16_Out{}
	ϟo.Result = int16(0)
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_s16 expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdF32) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdF32_Out{}
	ϟo.Result = float32(0)
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_f32 expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdU32) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdU32_Out{}
	ϟo.Result = uint32(0)
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_u32 expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdS32) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdS32_Out{}
	ϟo.Result = int32(0)
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_s32 expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdF64) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdF64_Out{}
	ϟo.Result = float64(0)
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_f64 expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdU64) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdU64_Out{}
	ϟo.Result = uint64(0)
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_u64 expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdS64) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdS64_Out{}
	ϟo.Result = int64(0)
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_s64 expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdBool) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdBool_Out{}
	ϟo.Result = false
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_bool expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdString) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdString_Out{}
	ϟo.Result = ""
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_string expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdArrayOfFloat) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdArrayOfFloat_Out{}
	ϟo.Result = make(F32Array, int32(10))
	ϟo.Result = ϟa.Out.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_array_of_float expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdPointer) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdPointer_Out{}
	ϟo.Result = ϟa.Out.Result
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_pointer expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdVoidOutU8) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdVoidOutU8_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_void_out_u8 expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdVoidOutS8) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdVoidOutS8_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_void_out_s8 expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdVoidOutU16) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdVoidOutU16_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_void_out_u16 expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdVoidOutS16) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdVoidOutS16_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_void_out_s16 expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdVoidOutF32) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdVoidOutF32_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_void_out_f32 expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdVoidOutU32) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdVoidOutU32_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_void_out_u32 expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdVoidOutS32) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdVoidOutS32_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_void_out_s32 expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdVoidOutF64) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdVoidOutF64_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_void_out_f64 expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdVoidOutU64) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdVoidOutU64_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_void_out_u64 expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdVoidOutS64) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdVoidOutS64_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_void_out_s64 expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdVoidOutBool) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdVoidOutBool_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_void_out_bool expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdVoidOutString) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdVoidOutString_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_void_out_string expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdVoidOutFixedSizeBuffer) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdVoidOutFixedSizeBuffer_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_void_out_fixed_size_buffer expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdVoidOut3Strings) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdVoidOut3Strings_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_void_out_3_strings expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdVoid3Remapped) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdVoid3Remapped_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_void_3_remapped expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdVoidOut3Remapped) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdVoidOut3Remapped_Out{}
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_void_out_3_remapped expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
func (ϟa *CmdVoidOutArrayOfRemapped) Mutate(ϟs *state.State) error {
	ϟc := getState(ϟa, ϟs)
	ϟo := CmdVoidOutArrayOfRemapped_Out{}
	ϟo.A = make(RemappedArray, int32(5))
	if ϟc.ValidateOutput && !reflect.DeepEqual(ϟa.Out, ϟo) {
		log.Printf("Applying cmd_void_out_array_of_remapped expected %v got %v", ϟa.Out, ϟo)
	}
	return nil
}
