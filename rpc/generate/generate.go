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

//go:generate embed

// Package generate contains the rpc code generation functionality.
package generate

import (
	"fmt"
	"path/filepath"
	"strings"

	"android.googlesource.com/platform/tools/gpu/api/apic/template"
	"android.googlesource.com/platform/tools/gpu/api/semantic"
	binary "android.googlesource.com/platform/tools/gpu/binary/generate"
)

func loader(filename string) ([]byte, error) {
	s, ok := embedded[filename]
	if !ok {
		return nil, fmt.Errorf("Unknown embedded filename %s", filename)
	}
	return []byte(s), nil
}

func fromType(from semantic.Type) *binary.Type {
	native := ""
	switch from {
	case semantic.BoolType:
		native = "Bool"
	case semantic.IntType:
		native = "Int"
	case semantic.UintType:
		native = "Uint"
	case semantic.Int8Type:
		native = "Int8"
	case semantic.Uint8Type:
		native = "Uint8"
	case semantic.Int16Type:
		native = "Int16"
	case semantic.Uint16Type:
		native = "Uint16"
	case semantic.Int32Type:
		native = "Int32"
	case semantic.Uint32Type:
		native = "Uint32"
	case semantic.Float32Type:
		native = "Float32"
	case semantic.Float64Type:
		native = "Float64"
	case semantic.Int64Type:
		native = "Int64"
	case semantic.Uint64Type:
		native = "Uint64"
	case semantic.StringType:
		native = "String"
	}
	if native != "" {
		return &binary.Type{
			Name:   strings.ToLower(native),
			Kind:   binary.Native,
			Method: native,
		}
	}
	switch from := from.(type) {
	case *semantic.Enum:
		return &binary.Type{
			Name:   from.Typename(),
			Kind:   binary.Remap,
			Method: "Int32",
			Native: "int32",
		}
	case *semantic.Pointer:
		t := &binary.Type{
			Kind:    binary.Pointer,
			SubType: fromType(from.To),
		}
		if t.SubType.Kind != binary.Interface {
			t.Name = fmt.Sprintf("*%s", t.SubType.Name)
		} else {
			t.Name = t.SubType.Name
		}
		return t
	case *semantic.Array:
		t := &binary.Type{
			Kind:    binary.Array,
			SubType: fromType(from.ValueType),
		}
		name := t.SubType.Name
		switch t.SubType.Kind {
		case binary.Pointer:
			name = t.SubType.SubType.Name
		case binary.Native:
			name = t.SubType.Method
			switch name {
			case "Uint8":
				name = "U8"
			case "Int8":
				name = "S8"
			case "Uint16":
				name = "U16"
			case "Int16":
				name = "S16"
			case "Uint32":
				name = "U32"
			case "Int32":
				name = "S32"
			case "Uint64":
				name = "U64"
			case "Int64":
				name = "S64"
			case "Float32":
				name = "F32"
			case "Float64":
				name = "F64"
			}
		}
		t.Name = fmt.Sprintf("%sArray", name)
		return t
	case *semantic.Class:
		if from.GetAnnotation("Interface") != nil {
			return &binary.Type{
				Name: from.Typename(),
				Kind: binary.Interface,
			}
		}
		return &binary.Type{
			Name: from.Typename(),
			Kind: binary.Codeable,
		}
	default:
		return &binary.Type{
			Name: from.Typename(),
			Kind: binary.Codeable,
		}
	}
}

func addFields(s *binary.Struct, c *semantic.Class) {
	for _, e := range c.Extends {
		addFields(s, e)
	}
	for _, decl := range c.Fields {
		f := binary.Field{
			Name: decl.Name,
			Type: fromType(decl.Type),
		}
		s.Fields = append(s.Fields, f)
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
