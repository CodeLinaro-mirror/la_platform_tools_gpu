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
	"path/filepath"
	"testing"
)

var fields = []Field{
	{Name: "u8", Type: &Type{Name: "uint8", Kind: Native, Method: "Uint8"}},
	{Name: "u16", Type: &Type{Name: "uint16", Kind: Native, Method: "Uint16"}},
	{Name: "u32", Type: &Type{Name: "uint32", Kind: Native, Method: "Uint32"}},
	{Name: "u64", Type: &Type{Name: "uint64", Kind: Native, Method: "Uint64"}},
	{Name: "i8", Type: &Type{Name: "int8", Kind: Native, Method: "Int8"}},
	{Name: "i16", Type: &Type{Name: "int16", Kind: Native, Method: "Int16"}},
	{Name: "i32", Type: &Type{Name: "int32", Kind: Native, Method: "Int32"}},
	{Name: "i64", Type: &Type{Name: "int64", Kind: Native, Method: "Int64"}},
	{Name: "f32", Type: &Type{Name: "float32", Kind: Native, Method: "Float32"}},
	{Name: "f64", Type: &Type{Name: "float64", Kind: Native, Method: "Float64"}},
	{Name: "bool", Type: &Type{Name: "bool", Kind: Native, Method: "Bool"}},
	{Name: "byte", Type: &Type{Name: "byte", Native: "uint8", Kind: Remap, Method: "Uint8"}},
	{Name: "int", Type: &Type{Name: "int", Native: "int32", Kind: Remap, Method: "Int32"}},
	{Name: "str", Type: &Type{Name: "string", Kind: Native, Method: "String", SkipMethod: "SkipString"}},
	{Name: "codeable", Type: &Type{Name: "struct{}", Kind: Codeable}},
	{Name: "pointer", Type: &Type{Name: "*struct{}", Kind: Pointer}},
	{Name: "slice", Type: &Type{Name: "[]struct{}", Kind: Array}},
	{Name: "stream", Type: &Type{Name: "[]struct{}", Kind: Stream}},
	{Name: "object", Type: &Type{Name: "interface{}", Kind: Interface}},
	{Name: "dict", Type: &Type{Name: "map[string]struct{}", Kind: Map}},
	{Name: "data", Type: &Type{Name: "[]byte", Kind: Array, Method: "Data"}},
	{Name: "id", Type: &Type{Name: "binary.ID", Native: "[20]byte", Kind: Native, Method: "ID", SkipMethod: "SkipID"}},
	{Name: "array", Type: &Type{Name: "Other", Native: "[10]int", Kind: StaticArray}},
	{Name: "", Type: &Type{Name: "Other", Native: "[10]int", Kind: StaticArray}, Anonymous: true},
}

var loader *Loader

func init() {
	pwd, _ := filepath.Abs(".")
	loader = NewLoader(pwd)
}

func parseStructs(source string) []*Struct {
	fakeFile := fmt.Sprintf(`
	package fake
	import "android.googlesource.com/platform/tools/gpu/binary"
	%s`, source)

	file, err := loader.ScanFile("internal.go", fakeFile)
	if err != nil {
		log.Fatalf("Parse failed:", err)
	}
	return file.Structs
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

func TestDisable(t *testing.T) {
	s := parseStructs("type MyStruct struct {binary.Generate `disable:\"true\"`}")
	if len(s) != 0 {
		t.Errorf("Got %d structs, expected none", len(s))
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
	source := &bytes.Buffer{}
	fmt.Fprintln(source, prefix)
	fmt.Fprint(source, "type MyStruct struct {binary.Generate;\n")
	for _, f := range fields {
		fmt.Fprintf(source, "  %s %s", f.Name, f.Type.Name)
		if f.Type.Kind == Stream {
			fmt.Fprint(source, " `stream:\"true\"`")
		}
		fmt.Fprintln(source)
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
		native := expected.Type.Native
		if len(native) == 0 {
			native = expected.Type.Name
		}
		if got.Type.Native != native {
			t.Errorf("Got field native type %s, expected %s for %s %s",
				got.Type.Native, native, expected.Name, expected.Type.Name)
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
