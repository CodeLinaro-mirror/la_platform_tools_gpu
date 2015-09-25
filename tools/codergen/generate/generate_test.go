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
	"reflect"
	"strings"
	"testing"

	"android.googlesource.com/platform/tools/gpu/binary/schema"
	"android.googlesource.com/platform/tools/gpu/tools/codergen/scan"
)

var (
	structType = &schema.Struct{Name: "TestObject"}
	fields     = []schema.Field{
		{Declared: "u8", Type: &schema.Primitive{Name: "uint8", Method: schema.Uint8}},
		{Declared: "u16", Type: &schema.Primitive{Name: "uint16", Method: schema.Uint16}},
		{Declared: "u32", Type: &schema.Primitive{Name: "uint32", Method: schema.Uint32}},
		{Declared: "u64", Type: &schema.Primitive{Name: "uint64", Method: schema.Uint64}},
		{Declared: "i8", Type: &schema.Primitive{Name: "int8", Method: schema.Int8}},
		{Declared: "i16", Type: &schema.Primitive{Name: "int16", Method: schema.Int16}},
		{Declared: "i32", Type: &schema.Primitive{Name: "int32", Method: schema.Int32}},
		{Declared: "i64", Type: &schema.Primitive{Name: "int64", Method: schema.Int64}},
		{Declared: "f32", Type: &schema.Primitive{Name: "float32", Method: schema.Float32}},
		{Declared: "f64", Type: &schema.Primitive{Name: "float64", Method: schema.Float64}},
		{Declared: "bool", Type: &schema.Primitive{Name: "bool", Method: schema.Bool}},
		{Declared: "byte", Type: &schema.Primitive{Name: "byte", Method: schema.Uint8}},
		{Declared: "int", Type: &schema.Primitive{Name: "int", Method: schema.Int32}},
		{Declared: "str", Type: &schema.Primitive{Name: "string", Method: schema.String}},
		{Declared: "codeable", Type: structType},
		{Declared: "pointer", Type: &schema.Pointer{Type: structType}},
		{Declared: "object", Type: &schema.Interface{Name: "binary.Object"}},
		{Declared: "slice", Type: &schema.Slice{ValueType: structType}},
		{Declared: "alias", Type: &schema.Slice{Alias: "Other", ValueType: &schema.Primitive{Name: "int", Method: schema.Int32}}},
		{Declared: "array", Type: &schema.Array{ValueType: &schema.Primitive{Name: "int", Method: schema.Int32}, Size: 10}},
		{Declared: "dict", Type: &schema.Map{KeyType: &schema.Primitive{Name: "string", Method: schema.String}, ValueType: structType}},
		{Declared: "data", Type: &schema.Slice{ValueType: &schema.Primitive{Name: "uint8", Method: schema.Uint8}}},
		{Declared: "id", Type: &schema.Primitive{Name: "binary.ID", Method: schema.ID}},
		{Declared: "", Type: structType},
	}
)
var testId int

func parseStructs(source string) []*Struct {
	testId++
	fakeFile := fmt.Sprintf(`
	package fake
	import "android.googlesource.com/platform/tools/gpu/binary"
	type TestObject struct{}
	func (*TestObject) Class() binary.Class { return nil }

	%s`, source)
	name := fmt.Sprintf("fake_%d.go", testId)
	pwd, _ := filepath.Abs(".")
	scanner := scan.New(pwd)
	scanner.ScanFile(name, fakeFile)
	if err := scanner.Process(); err != nil {
		log.Fatal("Process failed:", err)
	}
	modules, err := From(scanner)
	if err != nil {
		log.Fatal("Scan failed:", err)
	}
	for _, m := range modules {
		if m.Source.Directory.ImportPath == name && !m.IsTest {
			return m.Structs
		}
	}
	log.Fatal("Module find failed")
	return nil
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

func TestInterfaceVsAny(t *testing.T) {
	source := `
		type S struct {
			binary.Generate

			any0 interface{}
			any1 interface{ F() }
			any2 interface{ anyB; F2() }
			any3 anyA
			any4 anyB
			any5 anyC

			obj0 interface{ binary.Object }
			obj1 interface{ binary.Object; F() }
			obj2 interface{ objB; F3() }
			obj3 objA
			obj4 objB
			obj5 objC
		}

		type anyA interface{}
		type anyB interface{ F() }
		type anyC interface{ anyB; F2() }
		type objA interface{ binary.Object }
		type objB interface{ binary.Object; F() }
		type objC interface{ objB; F3() }
`

	s := parseStruct(t, "S", source)
	for _, f := range s.Fields {
		_, isAny := f.Type.(*schema.Any)
		_, isInt := f.Type.(*schema.Interface)
		if strings.HasPrefix(f.Name(), "any") && !isAny {
			t.Errorf("Field '%s' has unexpected type %T", f.Name(), f.Type)
		}
		if strings.HasPrefix(f.Name(), "obj") && !isInt {
			t.Errorf("Field '%s' has unexpected type %T", f.Name(), f.Type)
		}
	}
}
func TestStableID(t *testing.T) {
	source := "type MyStruct struct {binary.Generate}"
	a := parseStruct(t, "MyStruct", source)
	b := parseStruct(t, "MyStruct", source)
	if a.TypeID != b.TypeID {
		t.Errorf("ID was not stable")
	}
}

func TestNameAffectsID(t *testing.T) {
	a := parseStruct(t, "MyStruct", "type MyStruct struct {binary.Generate}")
	b := parseStruct(t, "YourStruct", "type YourStruct struct {binary.Generate}")
	if a.TypeID == b.TypeID {
		t.Errorf("Name change did not change ID")
	}
}

func TestFieldCountAffectsID(t *testing.T) {
	a := parseStruct(t, "MyStruct", "type MyStruct struct {binary.Generate; a int}")
	b := parseStruct(t, "MyStruct", "type MyStruct struct {binary.Generate}")
	if a.TypeID == b.TypeID {
		t.Errorf("Field count did not change ID")
	}
}

func TestFieldNameAffectsID(t *testing.T) {
	a := parseStruct(t, "MyStruct", "type MyStruct struct {binary.Generate; a int}")
	b := parseStruct(t, "MyStruct", "type MyStruct struct {binary.Generate; b int}")
	if a.TypeID == b.TypeID {
		t.Errorf("Field name did not change ID")
	}
}

func TestFieldTypeAffectsID(t *testing.T) {
	a := parseStruct(t, "MyStruct", "type MyStruct struct {binary.Generate; a int}")
	b := parseStruct(t, "MyStruct", "type MyStruct struct {binary.Generate; a byte}")
	if a.TypeID == b.TypeID {
		t.Errorf("Field type did not change ID")
	}
}

func TestTypes(t *testing.T) {
	prefix := "type Other []int\n"
	source := &bytes.Buffer{}
	fmt.Fprintln(source, prefix)
	fmt.Fprint(source, "type MyStruct struct {binary.Generate;\n")
	for _, f := range fields {
		fmt.Fprintf(source, "  %s %s\n", f.Declared, f.Type)
	}
	fmt.Fprint(source, "}\n")
	s := parseStruct(t, "MyStruct", source.String())
	if len(s.Fields) != len(fields) {
		t.Errorf("Got %d fields, expected %d", len(s.Fields), len(fields))
	}
	for i, got := range s.Fields {
		expected := fields[i]
		if got.Declared != expected.Declared {
			t.Errorf("Got field %s, expected %s", got.Declared, expected.Declared)
		}
		if reflect.TypeOf(got.Type) != reflect.TypeOf(expected.Type) {
			t.Errorf("Got field type %T, expected %T for %s", got.Type, expected.Type, expected.Name())
		}
		if got.Type.String() != expected.Type.String() {
			t.Errorf("Got field type %s, expected %s for %s", got.Type, expected.Type, expected.Name())
		}
	}
}
