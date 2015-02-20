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

//go:generate embed

// Package generate has support for generating encode and decode methods
// for the binary package automatically.
package generate

import (
	"bytes"
	"fmt"
	"path"
	"sort"
	"strings"
	"text/template"

	"android.googlesource.com/platform/tools/gpu/binary"
	"golang.org/x/tools/go/types"
)

type File struct {
	Generated string
	Package   string
	Structs   []*Struct
}

// Struct is a description of an encodable struct.
// Signature includes the package, name and name and type of all the fields.
// Any change to the Signature will cause the ID to change.
type Struct struct {
	Name       string    // The simple name of the type.
	Package    string    // The package name the struct belongs to.
	Fields     []Field   // Descriptions of the fields of the struct.
	Signature  string    // The full string type signature of the Struct.
	ID         binary.ID // The unique type identifier for the Struct.
	Delegating bool      // True if the struct is a pure delegating type.
}

// Kind describes the basic nature of a type.
type Kind int

const (
	// Native is the kind for primitive types with corresponding direct methods on
	// Encoder and Decoder
	Native Kind = iota
	// Remap is the kind for a type declared as alias to a primitive type.
	// For example: type U32 uint32.
	Remap
	// Codeable is the kind for a direct in place struct.
	Codeable
	// Pointer is the kind for a pointer to a struct type. If the struct instance
	// has equality (==) with a previously encoded object, then this struct will
	// be encoded as a reference to the first encoded object.
	Pointer
	// Array is the kind for an in place slice, with a dynamic length.
	Array
	// Interface is the kind for an object boxed in an binary.Object interface
	// (or superset of). If the object has equality (==) with a previously
	// encoded object, then this object will be encoded as a reference to the
	// first encoded object.
	Interface
)

// Field holds a description of a single Struct member.
type Field struct {
	// Name is the true field name.
	Name      string // The name the field was given.
	Type      *Type  // A description of the type of the field.
	Anonymous bool   // Whether the field was anonymous.
}

// Type is used to describe fields of a struct.
type Type struct {
	Name    string // The name of the type.
	Native  string // The go native name of the type.
	Kind    Kind   // The types basic Kind.
	SubType *Type  // If the type is an Array, holds the element type.
	Method  string // The encode/decode method to use.
}

// FromTypename creates and initializes a Struct from a types.Typename.
// It assumes that the typename will map to a types.Struct, and adds all the
// fields of that struct to the Struct information.
func FromTypename(pkg *types.Package, n *types.TypeName) *Struct {
	t := n.Type().Underlying().(*types.Struct)
	s := &Struct{Name: n.Name()}
	s.Fields = make([]Field, t.NumFields())
	s.Package = pkg.Name()
	for i := range s.Fields {
		decl := t.Field(i)
		f := &s.Fields[i]
		f.Name = decl.Name()
		f.Type = FromType(pkg, decl.Type())
		f.Anonymous = decl.Anonymous()
	}
	s.updateID()
	s.Delegating = len(s.Fields) == 1 && s.Fields[0].Anonymous
	return s
}

func (s *Struct) updateID() {
	b := &bytes.Buffer{}
	fmt.Fprintf(b, "struct %s.%s {", s.Package, s.Name)
	for i, f := range s.Fields {
		if i != 0 {
			fmt.Fprint(b, ",")
		}
		fmt.Fprintf(b, " %s:%s", f.Name, f.Type.Name)
	}
	fmt.Fprint(b, " }")
	s.Signature = b.String()
	s.ID = binary.NewID([]byte(s.Signature))
}

// FromType creates a appropriate Type object from a types.Type.
func FromType(pkg *types.Package, from types.Type) *Type {
	t := &Type{Name: path.Base(types.TypeString(pkg, from))}
	if _, isNamed := from.(*types.Named); isNamed {
		from = from.Underlying()
	}
	t.Native = from.String()
	switch from := from.(type) {
	case *types.Basic:
		t.Kind = Native
		switch from.Kind() {
		case types.Int:
			t.Native = "int32"
		case types.Byte:
			t.Native = "uint8"
		}
		t.Method = strings.Title(t.Native)
		if t.Native != t.Name {
			t.Kind = Remap
		}
	case *types.Pointer:
		t.Kind = Pointer
		t.SubType = FromType(pkg, from.Elem())
	case *types.Interface:
		t.Kind = Interface
	case *types.Slice:
		t.Kind = Array
		t.SubType = FromType(pkg, from.Elem())
	default:
		t.Kind = Codeable
	}
	return t
}

// Sort is used to ensure stable ordering of Struct slices.
// This is to ensure automatically generated code has minimum diffs.
// The sort order is by Struct name.
func Sort(structs []*Struct) {
	sort.Sort(structsByName(structs))
}

type structsByName []*Struct

func (a structsByName) Len() int           { return len(a) }
func (a structsByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a structsByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func getTemplate(t *template.Template, name string) *template.Template {
	result := t.Lookup(name)
	if result == nil {
		panic(fmt.Errorf("Could not find template %s", name))
	}
	return result
}

type kindToTemplate map[Kind]*template.Template

func kindDispatch(table kindToTemplate, name string, t *Type) string {
	b := &bytes.Buffer{}
	if err := table[t.Kind].Execute(b, Field{name, t, false}); err != nil {
		panic(err)
	}
	return b.String()
}
