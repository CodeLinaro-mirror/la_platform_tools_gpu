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

// Package generate contains the rpc code generation functionality.
package generate

import (
	"fmt"
	"path/filepath"

	"android.googlesource.com/platform/tools/gpu/api/apic/template"
	"android.googlesource.com/platform/tools/gpu/api/semantic"
	"android.googlesource.com/platform/tools/gpu/binary/schema"
)

func loader(filename string) ([]byte, error) {
	s, ok := embedded[filename]
	if !ok {
		return nil, fmt.Errorf("Unknown embedded filename %s", filename)
	}
	return []byte(s), nil
}

func fromType(from semantic.Type) schema.Type {
	switch from {
	case semantic.BoolType:
		return &schema.Primitive{Name: "bool", Method: schema.Bool}
	case semantic.IntType:
		return &schema.Primitive{Name: "int", Method: schema.Int32}
	case semantic.UintType:
		return &schema.Primitive{Name: "uint", Method: schema.Uint32}
	case semantic.Int8Type:
		return &schema.Primitive{Name: "int8", Method: schema.Int8}
	case semantic.Uint8Type:
		return &schema.Primitive{Name: "uint8", Method: schema.Uint8}
	case semantic.Int16Type:
		return &schema.Primitive{Name: "int16", Method: schema.Int16}
	case semantic.Uint16Type:
		return &schema.Primitive{Name: "uint16", Method: schema.Uint16}
	case semantic.Int32Type:
		return &schema.Primitive{Name: "int32", Method: schema.Int32}
	case semantic.Uint32Type:
		return &schema.Primitive{Name: "uint32", Method: schema.Uint32}
	case semantic.Float32Type:
		return &schema.Primitive{Name: "float32", Method: schema.Float32}
	case semantic.Float64Type:
		return &schema.Primitive{Name: "float64", Method: schema.Float64}
	case semantic.Int64Type:
		return &schema.Primitive{Name: "int64", Method: schema.Int64}
	case semantic.Uint64Type:
		return &schema.Primitive{Name: "uint64", Method: schema.Uint64}
	case semantic.StringType:
		return &schema.Primitive{Name: "string", Method: schema.String}
	}
	switch from := from.(type) {
	case *semantic.Enum:
		return &schema.Primitive{Name: from.Name(), Method: schema.Int32}
	case *semantic.Pointer:
		return &schema.Pointer{Type: fromType(from.To)}
	case *semantic.Slice:
		return &schema.Slice{ValueType: fromType(from.To)}
	case *semantic.Class:
		if from.GetAnnotation("Interface") != nil {
			return &schema.Interface{Name: from.Name()}
		}
		return &schema.Struct{Name: from.Name()}
	default:
		return &schema.Struct{Name: from.Name()}
	}
}

// Init prepares a new template processor that layers the apic one with
// the functions from the binary codec generate package.
func Init(apiFile string, api *semantic.API) *template.Functions {
	apiFile, _ = filepath.Abs(apiFile)
	return template.NewFunctions(apiFile, api, loader, map[string]interface{}{})
}

// Go invokes the main go code generation template.
func Go(f *template.Functions) error {
	return f.Include(rpc_go_tmpl_file)
}

// Java invokes the main go code generation template.
func Java(f *template.Functions) error {
	return f.Include(rpc_java_tmpl_file)
}
