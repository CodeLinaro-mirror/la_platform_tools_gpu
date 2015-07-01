////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

package test

import (
	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	bschema "android.googlesource.com/platform/tools/gpu/binary/schema"
	"android.googlesource.com/platform/tools/gpu/gfxapi/schema"
	"android.googlesource.com/platform/tools/gpu/service"
)

func init() {
	sc_CmdClone := bschema.Of((*CmdClone)(nil).Class())
	sc_CmdClone.Metadata = append(sc_CmdClone.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdMake := bschema.Of((*CmdMake)(nil).Class())
	sc_CmdMake.Metadata = append(sc_CmdMake.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdCopy := bschema.Of((*CmdCopy)(nil).Class())
	sc_CmdCopy.Metadata = append(sc_CmdCopy.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdCharsliceToString := bschema.Of((*CmdCharsliceToString)(nil).Class())
	sc_CmdCharsliceToString.Metadata = append(sc_CmdCharsliceToString.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdCharptrToString := bschema.Of((*CmdCharptrToString)(nil).Class())
	sc_CmdCharptrToString.Metadata = append(sc_CmdCharptrToString.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdSliceCasts := bschema.Of((*CmdSliceCasts)(nil).Class())
	sc_CmdSliceCasts.Metadata = append(sc_CmdSliceCasts.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoid := bschema.Of((*CmdVoid)(nil).Class())
	sc_CmdVoid.Metadata = append(sc_CmdVoid.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdUnknownRet := bschema.Of((*CmdUnknownRet)(nil).Class())
	sc_CmdUnknownRet.Metadata = append(sc_CmdUnknownRet.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdUnknownWritePtr := bschema.Of((*CmdUnknownWritePtr)(nil).Class())
	sc_CmdUnknownWritePtr.Metadata = append(sc_CmdUnknownWritePtr.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdUnknownWriteSlice := bschema.Of((*CmdUnknownWriteSlice)(nil).Class())
	sc_CmdUnknownWriteSlice.Metadata = append(sc_CmdUnknownWriteSlice.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidU8 := bschema.Of((*CmdVoidU8)(nil).Class())
	sc_CmdVoidU8.Metadata = append(sc_CmdVoidU8.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidS8 := bschema.Of((*CmdVoidS8)(nil).Class())
	sc_CmdVoidS8.Metadata = append(sc_CmdVoidS8.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidU16 := bschema.Of((*CmdVoidU16)(nil).Class())
	sc_CmdVoidU16.Metadata = append(sc_CmdVoidU16.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidS16 := bschema.Of((*CmdVoidS16)(nil).Class())
	sc_CmdVoidS16.Metadata = append(sc_CmdVoidS16.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidF32 := bschema.Of((*CmdVoidF32)(nil).Class())
	sc_CmdVoidF32.Metadata = append(sc_CmdVoidF32.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidU32 := bschema.Of((*CmdVoidU32)(nil).Class())
	sc_CmdVoidU32.Metadata = append(sc_CmdVoidU32.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidS32 := bschema.Of((*CmdVoidS32)(nil).Class())
	sc_CmdVoidS32.Metadata = append(sc_CmdVoidS32.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidF64 := bschema.Of((*CmdVoidF64)(nil).Class())
	sc_CmdVoidF64.Metadata = append(sc_CmdVoidF64.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidU64 := bschema.Of((*CmdVoidU64)(nil).Class())
	sc_CmdVoidU64.Metadata = append(sc_CmdVoidU64.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidS64 := bschema.Of((*CmdVoidS64)(nil).Class())
	sc_CmdVoidS64.Metadata = append(sc_CmdVoidS64.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidBool := bschema.Of((*CmdVoidBool)(nil).Class())
	sc_CmdVoidBool.Metadata = append(sc_CmdVoidBool.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidString := bschema.Of((*CmdVoidString)(nil).Class())
	sc_CmdVoidString.Metadata = append(sc_CmdVoidString.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoid3Strings := bschema.Of((*CmdVoid3Strings)(nil).Class())
	sc_CmdVoid3Strings.Metadata = append(sc_CmdVoid3Strings.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoid3InArrays := bschema.Of((*CmdVoid3InArrays)(nil).Class())
	sc_CmdVoid3InArrays.Metadata = append(sc_CmdVoid3InArrays.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidInArrayOfPointers := bschema.Of((*CmdVoidInArrayOfPointers)(nil).Class())
	sc_CmdVoidInArrayOfPointers.Metadata = append(sc_CmdVoidInArrayOfPointers.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidReadU8 := bschema.Of((*CmdVoidReadU8)(nil).Class())
	sc_CmdVoidReadU8.Metadata = append(sc_CmdVoidReadU8.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidReadS8 := bschema.Of((*CmdVoidReadS8)(nil).Class())
	sc_CmdVoidReadS8.Metadata = append(sc_CmdVoidReadS8.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidReadU16 := bschema.Of((*CmdVoidReadU16)(nil).Class())
	sc_CmdVoidReadU16.Metadata = append(sc_CmdVoidReadU16.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidReadS16 := bschema.Of((*CmdVoidReadS16)(nil).Class())
	sc_CmdVoidReadS16.Metadata = append(sc_CmdVoidReadS16.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidReadF32 := bschema.Of((*CmdVoidReadF32)(nil).Class())
	sc_CmdVoidReadF32.Metadata = append(sc_CmdVoidReadF32.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidReadU32 := bschema.Of((*CmdVoidReadU32)(nil).Class())
	sc_CmdVoidReadU32.Metadata = append(sc_CmdVoidReadU32.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidReadS32 := bschema.Of((*CmdVoidReadS32)(nil).Class())
	sc_CmdVoidReadS32.Metadata = append(sc_CmdVoidReadS32.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidReadF64 := bschema.Of((*CmdVoidReadF64)(nil).Class())
	sc_CmdVoidReadF64.Metadata = append(sc_CmdVoidReadF64.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidReadU64 := bschema.Of((*CmdVoidReadU64)(nil).Class())
	sc_CmdVoidReadU64.Metadata = append(sc_CmdVoidReadU64.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidReadS64 := bschema.Of((*CmdVoidReadS64)(nil).Class())
	sc_CmdVoidReadS64.Metadata = append(sc_CmdVoidReadS64.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidReadBool := bschema.Of((*CmdVoidReadBool)(nil).Class())
	sc_CmdVoidReadBool.Metadata = append(sc_CmdVoidReadBool.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidReadPtrs := bschema.Of((*CmdVoidReadPtrs)(nil).Class())
	sc_CmdVoidReadPtrs.Metadata = append(sc_CmdVoidReadPtrs.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidWriteU8 := bschema.Of((*CmdVoidWriteU8)(nil).Class())
	sc_CmdVoidWriteU8.Metadata = append(sc_CmdVoidWriteU8.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidWriteS8 := bschema.Of((*CmdVoidWriteS8)(nil).Class())
	sc_CmdVoidWriteS8.Metadata = append(sc_CmdVoidWriteS8.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidWriteU16 := bschema.Of((*CmdVoidWriteU16)(nil).Class())
	sc_CmdVoidWriteU16.Metadata = append(sc_CmdVoidWriteU16.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidWriteS16 := bschema.Of((*CmdVoidWriteS16)(nil).Class())
	sc_CmdVoidWriteS16.Metadata = append(sc_CmdVoidWriteS16.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidWriteF32 := bschema.Of((*CmdVoidWriteF32)(nil).Class())
	sc_CmdVoidWriteF32.Metadata = append(sc_CmdVoidWriteF32.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidWriteU32 := bschema.Of((*CmdVoidWriteU32)(nil).Class())
	sc_CmdVoidWriteU32.Metadata = append(sc_CmdVoidWriteU32.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidWriteS32 := bschema.Of((*CmdVoidWriteS32)(nil).Class())
	sc_CmdVoidWriteS32.Metadata = append(sc_CmdVoidWriteS32.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidWriteF64 := bschema.Of((*CmdVoidWriteF64)(nil).Class())
	sc_CmdVoidWriteF64.Metadata = append(sc_CmdVoidWriteF64.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidWriteU64 := bschema.Of((*CmdVoidWriteU64)(nil).Class())
	sc_CmdVoidWriteU64.Metadata = append(sc_CmdVoidWriteU64.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidWriteS64 := bschema.Of((*CmdVoidWriteS64)(nil).Class())
	sc_CmdVoidWriteS64.Metadata = append(sc_CmdVoidWriteS64.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidWriteBool := bschema.Of((*CmdVoidWriteBool)(nil).Class())
	sc_CmdVoidWriteBool.Metadata = append(sc_CmdVoidWriteBool.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidWritePtrs := bschema.Of((*CmdVoidWritePtrs)(nil).Class())
	sc_CmdVoidWritePtrs.Metadata = append(sc_CmdVoidWritePtrs.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdU8 := bschema.Of((*CmdU8)(nil).Class())
	sc_CmdU8.Metadata = append(sc_CmdU8.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdS8 := bschema.Of((*CmdS8)(nil).Class())
	sc_CmdS8.Metadata = append(sc_CmdS8.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdU16 := bschema.Of((*CmdU16)(nil).Class())
	sc_CmdU16.Metadata = append(sc_CmdU16.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdS16 := bschema.Of((*CmdS16)(nil).Class())
	sc_CmdS16.Metadata = append(sc_CmdS16.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdF32 := bschema.Of((*CmdF32)(nil).Class())
	sc_CmdF32.Metadata = append(sc_CmdF32.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdU32 := bschema.Of((*CmdU32)(nil).Class())
	sc_CmdU32.Metadata = append(sc_CmdU32.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdS32 := bschema.Of((*CmdS32)(nil).Class())
	sc_CmdS32.Metadata = append(sc_CmdS32.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdF64 := bschema.Of((*CmdF64)(nil).Class())
	sc_CmdF64.Metadata = append(sc_CmdF64.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdU64 := bschema.Of((*CmdU64)(nil).Class())
	sc_CmdU64.Metadata = append(sc_CmdU64.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdS64 := bschema.Of((*CmdS64)(nil).Class())
	sc_CmdS64.Metadata = append(sc_CmdS64.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdBool := bschema.Of((*CmdBool)(nil).Class())
	sc_CmdBool.Metadata = append(sc_CmdBool.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdString := bschema.Of((*CmdString)(nil).Class())
	sc_CmdString.Metadata = append(sc_CmdString.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdPointer := bschema.Of((*CmdPointer)(nil).Class())
	sc_CmdPointer.Metadata = append(sc_CmdPointer.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoid3Remapped := bschema.Of((*CmdVoid3Remapped)(nil).Class())
	sc_CmdVoid3Remapped.Metadata = append(sc_CmdVoid3Remapped.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidInArrayOfRemapped := bschema.Of((*CmdVoidInArrayOfRemapped)(nil).Class())
	sc_CmdVoidInArrayOfRemapped.Metadata = append(sc_CmdVoidInArrayOfRemapped.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidOutArrayOfRemapped := bschema.Of((*CmdVoidOutArrayOfRemapped)(nil).Class())
	sc_CmdVoidOutArrayOfRemapped.Metadata = append(sc_CmdVoidOutArrayOfRemapped.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdVoidOutArrayOfUnknownRemapped := bschema.Of((*CmdVoidOutArrayOfUnknownRemapped)(nil).Class())
	sc_CmdVoidOutArrayOfUnknownRemapped.Metadata = append(sc_CmdVoidOutArrayOfUnknownRemapped.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})

	sc_CmdRemapped := bschema.Of((*CmdRemapped)(nil).Class())
	sc_CmdRemapped.Metadata = append(sc_CmdRemapped.Metadata, &atom.Metadata{
		Api:              binary.ID(apiID),
		Flags:            0,
		DocumentationUrl: "[]",
	})
}
func init() {
	s := schemaBuilder{
		staticArrays: make(map[int]*service.StaticArrayInfo),
		maps:         make(map[int]*service.MapInfo),
		enums:        make(map[int]*service.EnumInfo),
		structs:      make(map[int]*service.StructInfo),
		classes:      make(map[int]*service.ClassInfo),
	}
	_ = s
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 0,
		Name: "cmd_clone",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "src",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "cnt",
				Type: schema.U32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 1,
		Name: "cmd_make",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "cnt",
				Type: schema.U32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 2,
		Name: "cmd_copy",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "src",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "cnt",
				Type: schema.U32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 3,
		Name: "cmd_charslice_to_string",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "s",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "len",
				Type: schema.U32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 4,
		Name: "cmd_charptr_to_string",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "s",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 5,
		Name: "cmd_slice_casts",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "s",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "l",
				Type: schema.U32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:              service.ApiId{ID: binary.ID(apiID)},
		Type:             6,
		Name:             "cmd_void",
		Parameters:       []service.ParameterInfo{},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 7,
		Name: "cmd_unknown_ret",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "result",
				Type: schema.Int,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 8,
		Name: "cmd_unknown_write_ptr",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "p",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 9,
		Name: "cmd_unknown_write_slice",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 10,
		Name: "cmd_void_u8",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.U8,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 11,
		Name: "cmd_void_s8",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.S8,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 12,
		Name: "cmd_void_u16",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.U16,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 13,
		Name: "cmd_void_s16",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.S16,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 14,
		Name: "cmd_void_f32",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.Float,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 15,
		Name: "cmd_void_u32",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.U32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 16,
		Name: "cmd_void_s32",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.S32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 17,
		Name: "cmd_void_f64",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.Double,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 18,
		Name: "cmd_void_u64",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.U64,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 19,
		Name: "cmd_void_s64",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.S64,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 20,
		Name: "cmd_void_bool",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.Bool,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 21,
		Name: "cmd_void_string",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.String,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 22,
		Name: "cmd_void_3_strings",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.String,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "b",
				Type: schema.String,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "c",
				Type: schema.String,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 23,
		Name: "cmd_void_3_in_arrays",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "b",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "c",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 24,
		Name: "cmd_void_in_array_of_pointers",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "count",
				Type: schema.S32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 25,
		Name: "cmd_void_read_u8",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 26,
		Name: "cmd_void_read_s8",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 27,
		Name: "cmd_void_read_u16",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 28,
		Name: "cmd_void_read_s16",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 29,
		Name: "cmd_void_read_f32",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 30,
		Name: "cmd_void_read_u32",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 31,
		Name: "cmd_void_read_s32",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 32,
		Name: "cmd_void_read_f64",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 33,
		Name: "cmd_void_read_u64",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 34,
		Name: "cmd_void_read_s64",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 35,
		Name: "cmd_void_read_bool",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 36,
		Name: "cmd_void_read_ptrs",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "b",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "c",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 37,
		Name: "cmd_void_write_u8",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 38,
		Name: "cmd_void_write_s8",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 39,
		Name: "cmd_void_write_u16",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 40,
		Name: "cmd_void_write_s16",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 41,
		Name: "cmd_void_write_f32",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 42,
		Name: "cmd_void_write_u32",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 43,
		Name: "cmd_void_write_s32",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 44,
		Name: "cmd_void_write_f64",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 45,
		Name: "cmd_void_write_u64",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 46,
		Name: "cmd_void_write_s64",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 47,
		Name: "cmd_void_write_bool",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 48,
		Name: "cmd_void_write_ptrs",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "b",
				Type: schema.Pointer,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "c",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 49,
		Name: "cmd_u8",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "result",
				Type: schema.U8,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 50,
		Name: "cmd_s8",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "result",
				Type: schema.S8,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 51,
		Name: "cmd_u16",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "result",
				Type: schema.U16,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 52,
		Name: "cmd_s16",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "result",
				Type: schema.S16,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 53,
		Name: "cmd_f32",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "result",
				Type: schema.Float,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 54,
		Name: "cmd_u32",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "result",
				Type: schema.U32,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 55,
		Name: "cmd_s32",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "result",
				Type: schema.S32,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 56,
		Name: "cmd_f64",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "result",
				Type: schema.Double,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 57,
		Name: "cmd_u64",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "result",
				Type: schema.U64,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 58,
		Name: "cmd_s64",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "result",
				Type: schema.S64,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 59,
		Name: "cmd_bool",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "result",
				Type: schema.Bool,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 60,
		Name: "cmd_string",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "result",
				Type: schema.String,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 61,
		Name: "cmd_pointer",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "result",
				Type: schema.Pointer,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 62,
		Name: "cmd_void_3_remapped",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "b",
				Type: schema.U32,
				Out:  false,
			},
			service.ParameterInfo{
				Name: "c",
				Type: schema.U32,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 63,
		Name: "cmd_void_in_array_of_remapped",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 64,
		Name: "cmd_void_out_array_of_remapped",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 65,
		Name: "cmd_void_out_array_of_unknown_remapped",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.Pointer,
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 66,
		Name: "cmd_remapped",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "result",
				Type: schema.U32,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
}

type schemaBuilder struct {
	staticArrays map[int]*service.StaticArrayInfo
	maps         map[int]*service.MapInfo
	enums        map[int]*service.EnumInfo
	structs      map[int]*service.StructInfo
	classes      map[int]*service.ClassInfo
}

func (s schemaBuilder) getStaticArrayInfo(id int) *service.StaticArrayInfo {
	e, f := s.staticArrays[id]
	if f {
		return e
	}
	switch id {
	}
	s.staticArrays[id] = e
	return e
}
func (s schemaBuilder) getMapInfo(id int) *service.MapInfo {
	e, f := s.maps[id]
	if f {
		return e
	}
	switch id {
	}
	s.maps[id] = e
	return e
}
func (s schemaBuilder) getEnumInfo(id int) *service.EnumInfo {
	e, f := s.enums[id]
	if f {
		return e
	}
	switch id {
	}
	s.enums[id] = e
	return e
}
func (s schemaBuilder) getClassInfo(id int) *service.ClassInfo {
	e, f := s.classes[id]
	if f {
		return e
	}
	switch id {
	case 0:
		e = service.CreateClassInfo(
			"Tester",
			service.TypeKindClass,
			service.FieldInfoPtrArray{
				&service.FieldInfo{
					Name: "A",
					Type: s.getClassInfo(-1),
				},
				&service.FieldInfo{
					Name: "B",
					Type: s.getClassInfo(1),
				},
			},
			service.ClassInfoPtrArray{},
		)
	case 1:
		e = service.CreateClassInfo(
			"Included",
			service.TypeKindClass,
			service.FieldInfoPtrArray{
				&service.FieldInfo{
					Name: "S",
					Type: schema.String,
				},
			},
			service.ClassInfoPtrArray{},
		)
	}
	s.classes[id] = e
	return e
}
