////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

package test

import (
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/gfxapi/schema"
	"android.googlesource.com/platform/tools/gpu/service"
)

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
		Api:              service.ApiId{ID: binary.ID(apiID)},
		Type:             5,
		Name:             "cmd_void",
		Parameters:       []service.ParameterInfo{},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Api:  service.ApiId{ID: binary.ID(apiID)},
		Type: 6,
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
		Type: 7,
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
		Type: 8,
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
		Type: 9,
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
		Type: 10,
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
		Type: 11,
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
		Type: 12,
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
		Type: 13,
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
		Type: 14,
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
		Type: 15,
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
		Type: 16,
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
		Type: 17,
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
		Type: 18,
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
		Type: 19,
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
		Type: 20,
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
		Type: 21,
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
		Type: 22,
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
		Type: 23,
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
		Type: 24,
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
		Type: 25,
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
		Type: 26,
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
		Type: 27,
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
		Type: 28,
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
		Type: 29,
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
		Type: 30,
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
		Type: 31,
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
		Type: 32,
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
		Type: 33,
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
		Type: 34,
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
		Type: 35,
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
		Type: 36,
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
		Type: 37,
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
		Type: 38,
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
		Type: 39,
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
		Type: 40,
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
		Type: 41,
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
		Type: 42,
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
		Type: 43,
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
		Type: 44,
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
		Type: 45,
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
		Type: 46,
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
		Type: 47,
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
		Type: 48,
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
		Type: 49,
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
		Type: 50,
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
		Type: 51,
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
		Type: 52,
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
		Type: 53,
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
		Type: 54,
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
		Type: 55,
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
		Type: 56,
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
		Type: 57,
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
		Type: 58,
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
		Type: 59,
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
		Type: 60,
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
		Type: 61,
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
		Type: 62,
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
		Type: 63,
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
		Type: 64,
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
		Type: 65,
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
	schema.RegisterAPI(api{}, service.StructInfo{
		Name: "state",
		Kind: service.TypeKindStruct,
		Fields: service.FieldInfoPtrArray{
			&service.FieldInfo{
				Name: "buf",
				Type: schema.Int, /* TODO: Slice */
			},
			&service.FieldInfo{
				Name: "str",
				Type: schema.String,
			},
		},
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
