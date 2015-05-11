// Copyright (C) 2015 The Android Open Source Project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package schema

import (
	"sort"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/service"
)

func init() {
	RegisterAtom(service.AtomInfo{
		Type: uint16(atom.TypeIDObservation),
		Name: "Observation",
		Parameters: service.ParameterInfoArray{
			service.ParameterInfo{Name: "Base", Type: Pointer},
			service.ParameterInfo{Name: "Size", Type: U64},
			service.ParameterInfo{Name: "ResourceID", Type: ID},
		},
	})
}

var schema service.Schema
var atomsNeedSorting = false

var Bool = service.CreateSimpleInfo("bool", service.TypeKindBool)
var Float = service.CreateSimpleInfo("float", service.TypeKindF32)
var Double = service.CreateSimpleInfo("double", service.TypeKindF64)
var Int = service.CreateSimpleInfo("int", service.TypeKindS64)
var Uint = service.CreateSimpleInfo("uint", service.TypeKindU64)
var S8 = service.CreateSimpleInfo("s8", service.TypeKindS8)
var U8 = service.CreateSimpleInfo("u8", service.TypeKindU8)
var S16 = service.CreateSimpleInfo("s16", service.TypeKindS16)
var U16 = service.CreateSimpleInfo("u16", service.TypeKindU16)
var S32 = service.CreateSimpleInfo("s32", service.TypeKindS32)
var U32 = service.CreateSimpleInfo("u32", service.TypeKindU32)
var S64 = service.CreateSimpleInfo("s64", service.TypeKindS64)
var U64 = service.CreateSimpleInfo("u64", service.TypeKindU64)
var Pointer = service.CreateSimpleInfo("pointer", service.TypeKindPointer)
var Memory = service.CreateSimpleInfo("memory", service.TypeKindMemory)
var String = service.CreateSimpleInfo("string", service.TypeKindString)
var Any = service.CreateSimpleInfo("any", service.TypeKindAny)
var ID = service.CreateSimpleInfo("id", service.TypeKindID)

type atomSorter service.AtomInfoArray

func (s atomSorter) Len() int           { return len(s) }
func (s atomSorter) Less(i, j int) bool { return s[i].Type < s[j].Type }
func (s atomSorter) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// Schema returns the schema of all registered atoms and APIs.
func Schema() service.Schema {
	if atomsNeedSorting {
		sort.Sort(atomSorter(schema.Atoms))
		atomsNeedSorting = false
	}
	return schema
}

// RegisterAtom registers the atom info a with the schema.
func RegisterAtom(a service.AtomInfo) {
	schema.Atoms = append(schema.Atoms, a)
	atomsNeedSorting = true
}

// RegisterAtom registers the graphics API api with the schema.
func RegisterAPI(api gfxapi.API, state service.StructInfo) {
	schema.Apis = append(schema.Apis, service.ApiSchema{
		Api:   service.ApiId{ID: binary.ID(api.ID())},
		State: state,
	})
}
