// Copyright (C) 2014 The Android Open Source Project
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

package generate

import (
	"bytes"
	"fmt"
	"strings"

	"android.googlesource.com/platform/tools/gpu/api/resolver"
	"android.googlesource.com/platform/tools/gpu/binary/schema"
)

var (
	cppTypeMap = map[string]string{
		"int8":       "int8_t",
		"uint8":      "uint8_t",
		"int16":      "int16_t",
		"uint16":     "uint16_t",
		"int32":      "int32_t",
		"uint32":     "uint32_t",
		"int64":      "int64_t",
		"uint64":     "uint64_t",
		"float32":    "float",
		"float64":    "double",
		"string":     "char*",
		"binary::ID": "gapic::Id",
	}
	cppMethodMap = map[string]string{
		"ID": "Id",
	}
)

func (f *functions) FixupName(n string) string {
	n = strings.Replace(n, ".", "::", -1)
	n = strings.Replace(n, resolver.PointerSuffix, "__P", -1)
	n = strings.Replace(n, resolver.SliceSuffix, "__S", -1)
	n = strings.Replace(n, resolver.ArraySuffix, "__A", -1)
	n = strings.Replace(n, resolver.TypeInfix, "__", -1)
	if result, ok := cppTypeMap[n]; ok {
		return result
	}
	return n
}

func (f *functions) CppStorage(t schema.Type) string {
	switch t := t.(type) {
	case *schema.Primitive:
		return f.FixupName(t.Native())
	case *schema.Struct:
		return f.FixupName(t.Name)
	case *schema.Interface:
		return "gapic::Encodable*"
	case *schema.Pointer:
		return f.CppStorage(t.Type) + "*"
	case *schema.Array:
		return f.CppStorage(t.ValueType) + "*"
	case *schema.Slice:
		return fmt.Sprintf("std::vector<%s>", f.CppStorage(t.ValueType))
	case *schema.Stream:
		return fmt.Sprintf("std::vector<%s>", f.CppStorage(t.ValueType))
	case *schema.Map:
		return fmt.Sprintf("std::unordered_map<%s,%s>*", f.CppStorage(t.KeyType), f.CppStorage(t.ValueType))
	default:
		panic(fmt.Errorf("Unknown value type %T", t))
	}
}

func (f *functions) CppMethod(t *schema.Primitive) string {
	n := t.Method.String()
	if result, ok := cppMethodMap[n]; ok {
		return result
	}
	return n
}

// CppFile generates the all the cpp code for a file with a set of structs.
func (g *Generator) CppFile(file *File) ([]byte, error) {
	g.f.prefix = "Cpp."
	f := *file
	if f.Indent == "" {
		f.Indent = "    "
	}
	b := &bytes.Buffer{}
	if err := g.f.execute(g.f.prefix+"File", b, &f); err != nil {
		return nil, err
	}
	s := b.String()
	s = strings.Replace(s, indent, f.Indent, -1)
	return []byte(s), nil
}
