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
	"log"
	"testing"

	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/go/types"
)

func parseStructs(source string) []*Struct {
	config := loader.Config{}
	fakeFile := fmt.Sprintf(`
	package fake
	import "android.googlesource.com/platform/tools/gpu/binary"
	%s`, source)
	file, err := config.ParseFile("", fakeFile)
	if err != nil {
		log.Fatalf("invalid source: %s", err)
	}
	config.CreateFromFiles("", file)
	info, err := config.Load()
	if err != nil {
		log.Fatalf("load failed: %s", err)
	}
	result := []*Struct{}
	for _, pkg := range info.Created {
		for _, def := range pkg.Defs {
			if n, ok := def.(*types.TypeName); ok {
				if t, ok := n.Type().(*types.Named); ok {
					if _, ok := t.Underlying().(*types.Struct); ok {
						result = append(result, FromTypename(pkg.Pkg, n))
					}
				}
			}
		}
	}
	return result
}

func parseStruct(t *testing.T, name string, source string) *Struct {
	s := parseStructs(source)
	if len(s) != 1 {
		log.Fatalf("Parsed %d structs, expected 1", len(s))
	}
	if s[0].Name != name {
		t.Errorf("Got struct %s, expected %s", s[0].Name, name)
	}
	return s[0]
}

func TestEmpty(t *testing.T) {
	s := parseStruct(t, "MyStruct", "type MyStruct struct {binary.Generate}")
	if len(s.Fields) != 0 {
		t.Errorf("Got %d fields, expected none", len(s.Fields))
	}
}

func TestStableID(t *testing.T) {
	source := "type MyStruct struct {binary.Generate}"
	a := parseStruct(t, "MyStruct", source)
	b := parseStruct(t, "MyStruct", source)
	if a.ID != b.ID {
		t.Errorf("ID was not stable")
	}
}

func TestNameAffectsID(t *testing.T) {
	a := parseStruct(t, "MyStruct", "type MyStruct struct {binary.Generate}")
	b := parseStruct(t, "YourStruct", "type YourStruct struct {binary.Generate}")
	if a.ID == b.ID {
		t.Errorf("Name change did not change ID")
	}
}

func TestFieldCountAffectsID(t *testing.T) {
	a := parseStruct(t, "MyStruct", "type MyStruct struct {binary.Generate; a int}")
	b := parseStruct(t, "MyStruct", "type MyStruct struct {binary.Generate}")
	if a.ID == b.ID {
		t.Errorf("Field count did not change ID")
	}
}

func TestFieldNameAffectsID(t *testing.T) {
	a := parseStruct(t, "MyStruct", "type MyStruct struct {binary.Generate; a int}")
	b := parseStruct(t, "MyStruct", "type MyStruct struct {binary.Generate; b int}")
	if a.ID == b.ID {
		t.Errorf("Field name did not change ID")
	}
}

func TestFieldTypeAffectsID(t *testing.T) {
	a := parseStruct(t, "MyStruct", "type MyStruct struct {binary.Generate; a int}")
	b := parseStruct(t, "MyStruct", "type MyStruct struct {binary.Generate; a byte}")
	if a.ID == b.ID {
		t.Errorf("Field type did not change ID")
	}
}

func TestTypes(t *testing.T) {
	prefix := "type Other [10]int\n"
	fields := []Field{
		{"a", &Type{"uint8", "uint8", Native, nil, nil, "Uint8", ""}, false},
		{"b", &Type{"uint16", "uint16", Native, nil, nil, "Uint16", ""}, false},
		{"c", &Type{"uint32", "uint32", Native, nil, nil, "Uint32", ""}, false},
		{"d", &Type{"uint64", "uint64", Native, nil, nil, "Uint64", ""}, false},
		{"e", &Type{"int8", "int8", Native, nil, nil, "Int8", ""}, false},
		{"f", &Type{"int16", "int16", Native, nil, nil, "Int16", ""}, false},
		{"g", &Type{"int32", "int32", Native, nil, nil, "Int32", ""}, false},
		{"h", &Type{"int64", "int64", Native, nil, nil, "Int64", ""}, false},
		{"i", &Type{"float32", "float32", Native, nil, nil, "Float32", ""}, false},
		{"j", &Type{"float64", "float64", Native, nil, nil, "Float64", ""}, false},
		{"k", &Type{"byte", "uint8", Remap, nil, nil, "Uint8", ""}, false},
		{"l", &Type{"int", "int32", Remap, nil, nil, "Int32", ""}, false},
		{"m", &Type{"bool", "bool", Native, nil, nil, "Bool", ""}, false},
		{"n", &Type{"string", "string", Native, nil, nil, "String", "SkipString"}, false},
		{"o", &Type{"struct{}", "struct{}", Codeable, nil, nil, "", ""}, false},
		{"p", &Type{"*struct{}", "*struct{}", Pointer, nil, nil, "", ""}, false},
		{"q", &Type{"[]struct{}", "[]struct{}", Array, nil, nil, "", ""}, false},
		{"r", &Type{"interface{}", "interface{}", Interface, nil, nil, "", ""}, false},
		{"s", &Type{"map[string]struct{}", "map[string]struct{}", Map, nil, nil, "", ""}, false},
		{"u", &Type{"[]byte", "[]byte", Array, nil, nil, "Data", ""}, false},
		{"v", &Type{"binary.ID", "[20]byte", Native, nil, nil, "ID", "SkipID"}, false},
		{"", &Type{"Other", "[10]int", Array, nil, nil, "", ""}, true},
	}
	source := &bytes.Buffer{}
	fmt.Fprintln(source, prefix)
	fmt.Fprint(source, "type MyStruct struct {binary.Generate;\n")
	for _, f := range fields {
		fmt.Fprintf(source, "  %s %s\n", f.Name, f.Type.Name)
	}
	fmt.Fprint(source, "}\n")
	s := parseStruct(t, "MyStruct", source.String())
	if len(s.Fields) != len(fields) {
		t.Errorf("Got %d fields, expected %d", len(s.Fields), len(fields))
	}
	for i, got := range s.Fields {
		expected := fields[i]
		if expected.Name == "" {
			expected.Name = expected.Type.Name
		}
		if got.Name != expected.Name {
			t.Errorf("Got field %s, expected %s", got.Name, expected.Name)
		}
		if got.Type.Kind != expected.Type.Kind {
			t.Errorf("Got field kind %d, expected %d for %s %s",
				got.Type.Kind, expected.Type.Kind, expected.Name, expected.Type.Name)
		}
		if got.Type.Native != expected.Type.Native {
			t.Errorf("Got field native type %s, expected %s for %s %s",
				got.Type.Native, expected.Type.Native, expected.Name, expected.Type.Name)
		}
		if got.Type.Method != expected.Type.Method {
			t.Errorf("Got encoder method %s, expected %s for %s %s",
				got.Type.Method, expected.Type.Method, expected.Name, expected.Type.Name)
		}
		if got.Type.SkipMethod != expected.Type.SkipMethod {
			t.Errorf("Got decoder skip method %s, expected %s for %s %s",
				got.Type.SkipMethod, expected.Type.SkipMethod, expected.Name, expected.Type.Name)
		}
		if got.Anonymous != expected.Anonymous {
			t.Errorf("Got anonymous %v, expected %v for %s %s", got.Anonymous, expected.Anonymous, expected.Name, expected.Type.Name)
		}
	}
}
