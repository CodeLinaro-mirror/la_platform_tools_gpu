////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

package test

import (
	"log"
	"reflect"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/memory"
)

type StateMutator struct {
	State          *state
	ValidateOutput bool
}

var _ atom.Writer = StateMutator{} // Interface compliance test
func (m StateMutator) Write(ψ atom.ID, Θ atom.Atom) {
	switch ω := Θ.(type) {
	case *CmdVoid:
		Σ := CmdVoid_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_void expected %v got %v", ω.Out, Σ)
		}
	case *CmdVoidU8:
		Σ := CmdVoidU8_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_void_u8 expected %v got %v", ω.Out, Σ)
		}
	case *CmdVoidS8:
		Σ := CmdVoidS8_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_void_s8 expected %v got %v", ω.Out, Σ)
		}
	case *CmdVoidU16:
		Σ := CmdVoidU16_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_void_u16 expected %v got %v", ω.Out, Σ)
		}
	case *CmdVoidS16:
		Σ := CmdVoidS16_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_void_s16 expected %v got %v", ω.Out, Σ)
		}
	case *CmdVoidF32:
		Σ := CmdVoidF32_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_void_f32 expected %v got %v", ω.Out, Σ)
		}
	case *CmdVoidU32:
		Σ := CmdVoidU32_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_void_u32 expected %v got %v", ω.Out, Σ)
		}
	case *CmdVoidS32:
		Σ := CmdVoidS32_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_void_s32 expected %v got %v", ω.Out, Σ)
		}
	case *CmdVoidF64:
		Σ := CmdVoidF64_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_void_f64 expected %v got %v", ω.Out, Σ)
		}
	case *CmdVoidU64:
		Σ := CmdVoidU64_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_void_u64 expected %v got %v", ω.Out, Σ)
		}
	case *CmdVoidS64:
		Σ := CmdVoidS64_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_void_s64 expected %v got %v", ω.Out, Σ)
		}
	case *CmdVoidBool:
		Σ := CmdVoidBool_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_void_bool expected %v got %v", ω.Out, Σ)
		}
	case *CmdVoidString:
		Σ := CmdVoidString_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_void_string expected %v got %v", ω.Out, Σ)
		}
	case *CmdVoid3Strings:
		Σ := CmdVoid3Strings_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_void_3_strings expected %v got %v", ω.Out, Σ)
		}
	case *CmdVoid3Arrays:
		Σ := CmdVoid3Arrays_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_void_3_arrays expected %v got %v", ω.Out, Σ)
		}
	case *CmdVoidArrayOfStrings:
		Σ := CmdVoidArrayOfStrings_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_void_array_of_strings expected %v got %v", ω.Out, Σ)
		}
	case *CmdU8:
		Σ := CmdU8_Out{}
		Σ.Result = 0
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_u8 expected %v got %v", ω.Out, Σ)
		}
	case *CmdS8:
		Σ := CmdS8_Out{}
		Σ.Result = 0
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_s8 expected %v got %v", ω.Out, Σ)
		}
	case *CmdU16:
		Σ := CmdU16_Out{}
		Σ.Result = 0
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_u16 expected %v got %v", ω.Out, Σ)
		}
	case *CmdS16:
		Σ := CmdS16_Out{}
		Σ.Result = 0
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_s16 expected %v got %v", ω.Out, Σ)
		}
	case *CmdF32:
		Σ := CmdF32_Out{}
		Σ.Result = 0
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_f32 expected %v got %v", ω.Out, Σ)
		}
	case *CmdU32:
		Σ := CmdU32_Out{}
		Σ.Result = 0
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_u32 expected %v got %v", ω.Out, Σ)
		}
	case *CmdS32:
		Σ := CmdS32_Out{}
		Σ.Result = 0
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_s32 expected %v got %v", ω.Out, Σ)
		}
	case *CmdF64:
		Σ := CmdF64_Out{}
		Σ.Result = 0
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_f64 expected %v got %v", ω.Out, Σ)
		}
	case *CmdU64:
		Σ := CmdU64_Out{}
		Σ.Result = 0
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_u64 expected %v got %v", ω.Out, Σ)
		}
	case *CmdS64:
		Σ := CmdS64_Out{}
		Σ.Result = 0
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_s64 expected %v got %v", ω.Out, Σ)
		}
	case *CmdBool:
		Σ := CmdBool_Out{}
		Σ.Result = false
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_bool expected %v got %v", ω.Out, Σ)
		}
	case *CmdString:
		Σ := CmdString_Out{}
		Σ.Result = ""
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_string expected %v got %v", ω.Out, Σ)
		}
	case *CmdArrayOfFloat:
		Σ := CmdArrayOfFloat_Out{}
		Σ.Result = make(F32Array, 10)
		Σ.Result = ω.Out.Result
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_array_of_float expected %v got %v", ω.Out, Σ)
		}
	case *CmdPointer:
		Σ := CmdPointer_Out{}
		Σ.Result = ω.Out.Result
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_pointer expected %v got %v", ω.Out, Σ)
		}
	case *CmdVoidOutU8:
		Σ := CmdVoidOutU8_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_void_out_u8 expected %v got %v", ω.Out, Σ)
		}
	case *CmdVoidOutS8:
		Σ := CmdVoidOutS8_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_void_out_s8 expected %v got %v", ω.Out, Σ)
		}
	case *CmdVoidOutU16:
		Σ := CmdVoidOutU16_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_void_out_u16 expected %v got %v", ω.Out, Σ)
		}
	case *CmdVoidOutS16:
		Σ := CmdVoidOutS16_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_void_out_s16 expected %v got %v", ω.Out, Σ)
		}
	case *CmdVoidOutF32:
		Σ := CmdVoidOutF32_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_void_out_f32 expected %v got %v", ω.Out, Σ)
		}
	case *CmdVoidOutU32:
		Σ := CmdVoidOutU32_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_void_out_u32 expected %v got %v", ω.Out, Σ)
		}
	case *CmdVoidOutS32:
		Σ := CmdVoidOutS32_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_void_out_s32 expected %v got %v", ω.Out, Σ)
		}
	case *CmdVoidOutF64:
		Σ := CmdVoidOutF64_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_void_out_f64 expected %v got %v", ω.Out, Σ)
		}
	case *CmdVoidOutU64:
		Σ := CmdVoidOutU64_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_void_out_u64 expected %v got %v", ω.Out, Σ)
		}
	case *CmdVoidOutS64:
		Σ := CmdVoidOutS64_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_void_out_s64 expected %v got %v", ω.Out, Σ)
		}
	case *CmdVoidOutBool:
		Σ := CmdVoidOutBool_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_void_out_bool expected %v got %v", ω.Out, Σ)
		}
	case *CmdVoidOutString:
		Σ := CmdVoidOutString_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_void_out_string expected %v got %v", ω.Out, Σ)
		}
	case *CmdVoidOutFixedSizeBuffer:
		Σ := CmdVoidOutFixedSizeBuffer_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_void_out_fixed_size_buffer expected %v got %v", ω.Out, Σ)
		}
	case *CmdVoidOut3Strings:
		Σ := CmdVoidOut3Strings_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_void_out_3_strings expected %v got %v", ω.Out, Σ)
		}
	case *CmdVoid3Remapped:
		Σ := CmdVoid3Remapped_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_void_3_remapped expected %v got %v", ω.Out, Σ)
		}
	case *CmdVoidOut3Remapped:
		Σ := CmdVoidOut3Remapped_Out{}
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_void_out_3_remapped expected %v got %v", ω.Out, Σ)
		}
	case *CmdVoidOutArrayOfRemapped:
		Σ := CmdVoidOutArrayOfRemapped_Out{}
		Σ.A = make(RemappedArray, 5)
		if m.ValidateOutput && !reflect.DeepEqual(ω.Out, Σ) {
			log.Printf("Applying cmd_void_out_array_of_remapped expected %v got %v", ω.Out, Σ)
		}
	case *memory.Observation:
		m.State.Mem.Slice(ω.Range).Write(memory.ResourceData(ω.ResourceID, ω.Range.Size))
	}
}
