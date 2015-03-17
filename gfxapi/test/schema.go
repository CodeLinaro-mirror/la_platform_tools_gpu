////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

package test

import (
	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/service"
)

func (a api) Schema() service.Schema {
	s := schemaBuilder{
		arrays:       make(map[int]*service.ArrayInfo),
		staticArrays: make(map[int]*service.StaticArrayInfo),
		maps:         make(map[int]*service.MapInfo),
		enums:        make(map[int]*service.EnumInfo),
		structs:      make(map[int]*service.StructInfo),
		classes:      make(map[int]*service.ClassInfo),
		boolInfo:     service.CreateSimpleInfo("bool", service.TypeKindBool),
		floatInfo:    service.CreateSimpleInfo("float", service.TypeKindF32),
		doubleInfo:   service.CreateSimpleInfo("double", service.TypeKindF64),
		intInfo:      service.CreateSimpleInfo("int", service.TypeKindS8),
		uintInfo:     service.CreateSimpleInfo("uint", service.TypeKindU8),
		s8Info:       service.CreateSimpleInfo("s8", service.TypeKindS8),
		u8Info:       service.CreateSimpleInfo("u8", service.TypeKindU8),
		s16Info:      service.CreateSimpleInfo("s16", service.TypeKindS16),
		u16Info:      service.CreateSimpleInfo("u16", service.TypeKindU16),
		s32Info:      service.CreateSimpleInfo("s32", service.TypeKindS32),
		u32Info:      service.CreateSimpleInfo("u32", service.TypeKindU32),
		s64Info:      service.CreateSimpleInfo("s64", service.TypeKindS64),
		u64Info:      service.CreateSimpleInfo("u64", service.TypeKindU64),
		pointerInfo:  service.CreateSimpleInfo("pointer", service.TypeKindPointer),
		memoryInfo:   service.CreateSimpleInfo("memory", service.TypeKindMemory),
		stringInfo:   service.CreateSimpleInfo("string", service.TypeKindString),
		anyInfo:      service.CreateSimpleInfo("any", service.TypeKindAny),
		idInfo:       service.CreateSimpleInfo("id", service.TypeKindID),
	}
	return service.Schema{
		Arrays: service.ArrayInfoArray{
			s.getArrayInfo(0),
			s.getArrayInfo(1),
			s.getArrayInfo(2),
			s.getArrayInfo(3),
			s.getArrayInfo(4),
		},
		StaticArrays: service.StaticArrayInfoArray{},
		Maps:         service.MapInfoArray{},
		Enums:        service.EnumInfoArray{},
		Classes:      service.ClassInfoArray{},
		Atoms: service.AtomInfoArray{
			service.AtomInfo{
				Type:             0,
				Name:             "cmd_void",
				Parameters:       []service.ParameterInfo{},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 1,
				Name: "cmd_void_u8",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "a",
						Type: s.u8Info,
						Out:  false,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 2,
				Name: "cmd_void_s8",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "a",
						Type: s.s8Info,
						Out:  false,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 3,
				Name: "cmd_void_u16",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "a",
						Type: s.u16Info,
						Out:  false,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 4,
				Name: "cmd_void_s16",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "a",
						Type: s.s16Info,
						Out:  false,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 5,
				Name: "cmd_void_f32",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "a",
						Type: s.floatInfo,
						Out:  false,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 6,
				Name: "cmd_void_u32",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "a",
						Type: s.u32Info,
						Out:  false,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 7,
				Name: "cmd_void_s32",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "a",
						Type: s.s32Info,
						Out:  false,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 8,
				Name: "cmd_void_f64",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "a",
						Type: s.doubleInfo,
						Out:  false,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 9,
				Name: "cmd_void_u64",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "a",
						Type: s.u64Info,
						Out:  false,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 10,
				Name: "cmd_void_s64",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "a",
						Type: s.s64Info,
						Out:  false,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 11,
				Name: "cmd_void_bool",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "a",
						Type: s.boolInfo,
						Out:  false,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 12,
				Name: "cmd_void_string",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "a",
						Type: s.stringInfo,
						Out:  false,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 13,
				Name: "cmd_void_3_strings",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "a",
						Type: s.stringInfo,
						Out:  false,
					},
					service.ParameterInfo{
						Name: "b",
						Type: s.stringInfo,
						Out:  false,
					},
					service.ParameterInfo{
						Name: "c",
						Type: s.stringInfo,
						Out:  false,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
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
				DocumentationUrl: "",
			}, service.AtomInfo{
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
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 16,
				Name: "cmd_u8",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "result",
						Type: s.u8Info,
						Out:  true,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 17,
				Name: "cmd_s8",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "result",
						Type: s.s8Info,
						Out:  true,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 18,
				Name: "cmd_u16",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "result",
						Type: s.u16Info,
						Out:  true,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 19,
				Name: "cmd_s16",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "result",
						Type: s.s16Info,
						Out:  true,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 20,
				Name: "cmd_f32",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "result",
						Type: s.floatInfo,
						Out:  true,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 21,
				Name: "cmd_u32",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "result",
						Type: s.u32Info,
						Out:  true,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 22,
				Name: "cmd_s32",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "result",
						Type: s.s32Info,
						Out:  true,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 23,
				Name: "cmd_f64",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "result",
						Type: s.doubleInfo,
						Out:  true,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 24,
				Name: "cmd_u64",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "result",
						Type: s.u64Info,
						Out:  true,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 25,
				Name: "cmd_s64",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "result",
						Type: s.s64Info,
						Out:  true,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 26,
				Name: "cmd_bool",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "result",
						Type: s.boolInfo,
						Out:  true,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 27,
				Name: "cmd_string",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "result",
						Type: s.stringInfo,
						Out:  true,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
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
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 29,
				Name: "cmd_pointer",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "result",
						Type: s.pointerInfo,
						Out:  true,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 30,
				Name: "cmd_void_out_u8",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "a",
						Type: s.u8Info,
						Out:  true,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 31,
				Name: "cmd_void_out_s8",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "a",
						Type: s.s8Info,
						Out:  true,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 32,
				Name: "cmd_void_out_u16",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "a",
						Type: s.u16Info,
						Out:  true,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 33,
				Name: "cmd_void_out_s16",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "a",
						Type: s.s16Info,
						Out:  true,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 34,
				Name: "cmd_void_out_f32",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "a",
						Type: s.floatInfo,
						Out:  true,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 35,
				Name: "cmd_void_out_u32",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "a",
						Type: s.u32Info,
						Out:  true,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 36,
				Name: "cmd_void_out_s32",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "a",
						Type: s.s32Info,
						Out:  true,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 37,
				Name: "cmd_void_out_f64",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "a",
						Type: s.doubleInfo,
						Out:  true,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 38,
				Name: "cmd_void_out_u64",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "a",
						Type: s.u64Info,
						Out:  true,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 39,
				Name: "cmd_void_out_s64",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "a",
						Type: s.s64Info,
						Out:  true,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 40,
				Name: "cmd_void_out_bool",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "a",
						Type: s.boolInfo,
						Out:  true,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 41,
				Name: "cmd_void_out_string",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "a",
						Type: s.stringInfo,
						Out:  true,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 42,
				Name: "cmd_void_out_fixed_size_buffer",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "a",
						Type: s.pointerInfo,
						Out:  true,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 43,
				Name: "cmd_void_out_3_strings",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "a",
						Type: s.stringInfo,
						Out:  true,
					},
					service.ParameterInfo{
						Name: "b",
						Type: s.stringInfo,
						Out:  true,
					},
					service.ParameterInfo{
						Name: "c",
						Type: s.stringInfo,
						Out:  true,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 44,
				Name: "cmd_void_3_remapped",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "a",
						Type: s.u32Info,
						Out:  false,
					},
					service.ParameterInfo{
						Name: "b",
						Type: s.u32Info,
						Out:  false,
					},
					service.ParameterInfo{
						Name: "c",
						Type: s.u32Info,
						Out:  false,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
				Type: 45,
				Name: "cmd_void_out_3_remapped",
				Parameters: []service.ParameterInfo{
					service.ParameterInfo{
						Name: "a",
						Type: s.u32Info,
						Out:  true,
					},
					service.ParameterInfo{
						Name: "b",
						Type: s.u32Info,
						Out:  true,
					},
					service.ParameterInfo{
						Name: "c",
						Type: s.u32Info,
						Out:  true,
					},
				},
				IsCommand:        true,
				IsDrawCall:       false,
				IsEndOfFrame:     false,
				DocumentationUrl: "",
			}, service.AtomInfo{
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
				DocumentationUrl: "",
			},
			service.AtomInfo{
				Type: uint16(memory.TypeIDObservation),
				Name: "MemoryObservation",
				Parameters: service.ParameterInfoArray{
					service.ParameterInfo{
						Name: "Pointer",
						Type: s.pointerInfo,
					},
					service.ParameterInfo{
						Name: "Size",
						Type: s.u64Info,
					},
					service.ParameterInfo{
						Name: "ResourceID",
						Type: s.idInfo,
					},
				},
			},
			service.AtomInfo{
				Type: uint16(atom.TypeIDEos),
				Name: "EndOfStream",
			},
		},
		State: service.CreateStructInfo(
			"globals",
			service.TypeKindStruct,
			service.FieldInfoArray{},
		),
	}
}

type schemaBuilder struct {
	arrays       map[int]*service.ArrayInfo
	staticArrays map[int]*service.StaticArrayInfo
	maps         map[int]*service.MapInfo
	enums        map[int]*service.EnumInfo
	structs      map[int]*service.StructInfo
	classes      map[int]*service.ClassInfo
	boolInfo     service.TypeInfo
	intInfo      service.TypeInfo
	uintInfo     service.TypeInfo
	floatInfo    service.TypeInfo
	doubleInfo   service.TypeInfo
	s8Info       service.TypeInfo
	u8Info       service.TypeInfo
	s16Info      service.TypeInfo
	u16Info      service.TypeInfo
	s32Info      service.TypeInfo
	u32Info      service.TypeInfo
	s64Info      service.TypeInfo
	u64Info      service.TypeInfo
	pointerInfo  service.TypeInfo
	memoryInfo   service.TypeInfo
	stringInfo   service.TypeInfo
	anyInfo      service.TypeInfo
	idInfo       service.TypeInfo
}

func (s schemaBuilder) getArrayInfo(id int) *service.ArrayInfo {
	e, f := s.arrays[id]
	if f {
		return e
	}
	switch id {
	case 0:
		e = service.CreateArrayInfo(
			"BoolArray",
			service.TypeKindArray, s.boolInfo,
		)
	case 1:
		e = service.CreateArrayInfo(
			"F32Array",
			service.TypeKindArray, s.floatInfo,
		)
	case 2:
		e = service.CreateArrayInfo(
			"RemappedArray",
			service.TypeKindArray, s.u32Info,
		)
	case 3:
		e = service.CreateArrayInfo(
			"S8Array",
			service.TypeKindArray, s.s8Info,
		)
	case 4:
		e = service.CreateArrayInfo(
			"StringArray",
			service.TypeKindArray, s.stringInfo,
		)
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
