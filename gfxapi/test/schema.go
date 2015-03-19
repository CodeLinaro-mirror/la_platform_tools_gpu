////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

package test

import (
	"android.googlesource.com/platform/tools/gpu/gfxapi/schema"
	"android.googlesource.com/platform/tools/gpu/service"
)

func init() {
	s := schemaBuilder{
		arrays:       make(map[int]*service.ArrayInfo),
		staticArrays: make(map[int]*service.StaticArrayInfo),
		maps:         make(map[int]*service.MapInfo),
		enums:        make(map[int]*service.EnumInfo),
		structs:      make(map[int]*service.StructInfo),
		classes:      make(map[int]*service.ClassInfo),
	}
	schema.RegisterAtom(service.AtomInfo{
		Type:             0,
		Name:             "cmd_void",
		Parameters:       []service.ParameterInfo{},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 1,
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
		DocumentationUrl: "[///////////////////////////////////////////////////////////// /////////////////////////////////////////////////////////////]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 2,
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
		Type: 3,
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
		Type: 4,
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
		Type: 5,
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
		Type: 6,
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
		Type: 7,
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
		Type: 8,
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
		Type: 9,
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
		Type: 10,
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
		Type: 11,
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
		Type: 12,
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
		Type: 13,
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
		DocumentationUrl: "[///////////////////////////////////////////////////////////// /////////////////////////////////////////////////////////////]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 14,
		Name: "cmd_void_3_arrays",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: s.getArrayInfo(3),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "b",
				Type: s.getArrayInfo(4),
				Out:  false,
			},
			service.ParameterInfo{
				Name: "c",
				Type: s.getArrayInfo(0),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[///////////////////////////////////////////////////////////// /////////////////////////////////////////////////////////////]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 15,
		Name: "cmd_void_array_of_strings",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: s.getArrayInfo(4),
				Out:  false,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[///////////////////////////////////////////////////////////// /////////////////////////////////////////////////////////////]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 16,
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
		DocumentationUrl: "[///////////////////////////////////////////////////////////// /////////////////////////////////////////////////////////////]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 17,
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
		Type: 18,
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
		Type: 19,
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
		Type: 20,
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
		Type: 21,
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
		Type: 22,
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
		Type: 23,
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
		Type: 24,
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
		Type: 25,
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
		Type: 26,
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
		Type: 27,
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
		Type: 28,
		Name: "cmd_array_of_float",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "result",
				Type: s.getArrayInfo(1),
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 29,
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
		Type: 30,
		Name: "cmd_void_out_u8",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.U8,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[///////////////////////////////////////////////////////////// /////////////////////////////////////////////////////////////]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 31,
		Name: "cmd_void_out_s8",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
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
		Type: 32,
		Name: "cmd_void_out_u16",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
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
		Type: 33,
		Name: "cmd_void_out_s16",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
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
		Type: 34,
		Name: "cmd_void_out_f32",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
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
		Type: 35,
		Name: "cmd_void_out_u32",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
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
		Type: 36,
		Name: "cmd_void_out_s32",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
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
		Type: 37,
		Name: "cmd_void_out_f64",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
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
		Type: 38,
		Name: "cmd_void_out_u64",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
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
		Type: 39,
		Name: "cmd_void_out_s64",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
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
		Type: 40,
		Name: "cmd_void_out_bool",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
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
		Type: 41,
		Name: "cmd_void_out_string",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
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
		Type: 42,
		Name: "cmd_void_out_fixed_size_buffer",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
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
		Type: 43,
		Name: "cmd_void_out_3_strings",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.String,
				Out:  true,
			},
			service.ParameterInfo{
				Name: "b",
				Type: schema.String,
				Out:  true,
			},
			service.ParameterInfo{
				Name: "c",
				Type: schema.String,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[///////////////////////////////////////////////////////////// /////////////////////////////////////////////////////////////]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 44,
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
		DocumentationUrl: "[///////////////////////////////////////////////////////////// /////////////////////////////////////////////////////////////]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 45,
		Name: "cmd_void_out_3_remapped",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: schema.U32,
				Out:  true,
			},
			service.ParameterInfo{
				Name: "b",
				Type: schema.U32,
				Out:  true,
			},
			service.ParameterInfo{
				Name: "c",
				Type: schema.U32,
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[///////////////////////////////////////////////////////////// /////////////////////////////////////////////////////////////]",
	})
	schema.RegisterAtom(service.AtomInfo{
		Type: 46,
		Name: "cmd_void_out_array_of_remapped",
		Parameters: []service.ParameterInfo{
			service.ParameterInfo{
				Name: "a",
				Type: s.getArrayInfo(2),
				Out:  true,
			},
		},
		IsCommand:        true,
		IsDrawCall:       false,
		IsEndOfFrame:     false,
		DocumentationUrl: "[]",
	})
	schema.RegisterAPI(api{}, service.StructInfo{
		Name:   "state",
		Kind:   service.TypeKindStruct,
		Fields: service.FieldInfoArray{},
	})
}

type schemaBuilder struct {
	arrays       map[int]*service.ArrayInfo
	staticArrays map[int]*service.StaticArrayInfo
	maps         map[int]*service.MapInfo
	enums        map[int]*service.EnumInfo
	structs      map[int]*service.StructInfo
	classes      map[int]*service.ClassInfo
}

func (s schemaBuilder) getArrayInfo(id int) *service.ArrayInfo {
	e, f := s.arrays[id]
	if f {
		return e
	}
	switch id {
	case 0:
		e = service.CreateArrayInfo("BoolArray", service.TypeKindArray, schema.Bool)
	case 1:
		e = service.CreateArrayInfo("F32Array", service.TypeKindArray, schema.Float)
	case 2:
		e = service.CreateArrayInfo("RemappedArray", service.TypeKindArray, schema.U32)
	case 3:
		e = service.CreateArrayInfo("S8Array", service.TypeKindArray, schema.S8)
	case 4:
		e = service.CreateArrayInfo("StringArray", service.TypeKindArray, schema.String)
	}
	s.arrays[id] = e
	return e
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
	}
	s.classes[id] = e
	return e
}
