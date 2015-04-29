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

// Package generate has support for generating encode and decode methods
// for the binary package automatically.
package generate

import (
	"bytes"
	"fmt"
	"path"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"unicode"

	"android.googlesource.com/platform/tools/gpu/binary"
	"golang.org/x/tools/go/types"
)

type Style struct {
	ClassPrefix  string
	MemberPrefix string
	Indent       string
}

type Imports map[string]struct{}

type File struct {
	Generated string
	Package   string
	Import    string
	IsTest    bool
	Path      string
	Structs   []*Struct
	Imports   Imports
	Style
}

// Struct is a description of an encodable struct.
// Signature includes the package, name and name and type of all the fields.
// Any change to the Signature will cause the ID to change.
type Struct struct {
	Name      string    // The simple name of the type.
	IDName    string    // The name to give the ID of the type.
	Package   string    // The package name the struct belongs to.
	Fields    []Field   // Descriptions of the fields of the struct.
	Signature string    // The full string type signature of the Struct.
	ID        binary.ID // The unique type identifier for the Struct.
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
	// StaticArray is the kind for an in place array, with a fixed length.
	StaticArray
	// Stream is the kind for an in place slice, with a Terminator.
	Stream
	// Interface is the kind for an object boxed in an binary.Object interface
	// (or superset of). If the object has equality (==) with a previously
	// encoded object, then this object may be encoded as a reference to the
	// first encoded object.
	Interface
	// Map is the kind for a key value map.
	Map
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
	Name       string // The name of the type.
	Native     string // The go native name of the type.
	Kind       Kind   // The types basic Kind.
	KeyType    *Type  // If the type is a Map, holds the key type.
	SubType    *Type  // If the type is an Array, Map, Pointer or StaticArray, holds the element type.
	Length     int    // If the type is a StaticArray, holds the fixed array size.
	Method     string // The encode/decode method to use.
	SkipMethod string // The skip method to use.
}

type tag string

func (t tag) Get(name string) string {
	return reflect.StructTag(t).Get(name)
}

func (t tag) Flag(name string) bool {
	v := reflect.StructTag(t).Get(name)
	if len(v) == 0 {
		return false
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		panic(fmt.Errorf("Malformed struct tag %q in %q: %v", name, t, err))
	}
	return b
}

// FromTypename creates and initializes a Struct from a types.Typename.
// It assumes that the typename will map to a types.Struct, and adds all the
// fields of that struct to the Struct information.
func FromTypename(pkg *types.Package, n *types.TypeName, imports Imports) *Struct {
	t := n.Type().Underlying().(*types.Struct)
	s := &Struct{Name: n.Name()}
	s.Package = pkg.Name()
	tagged := false
	for i := 0; i < t.NumFields(); i++ {
		decl := t.Field(i)
		tag := tag(t.Tag(i))
		if decl.Anonymous() &&
			decl.Type().String() == "android.googlesource.com/platform/tools/gpu/binary.Generate" &&
			!tag.Flag("disable") {
			tagged = true
			s.IDName = tag.Get("id")
			continue
		}
		f := Field{}
		f.Name = decl.Name()
		f.Type = fromType(pkg, decl.Type(), tag, imports)
		delete(imports, pkg.Path())
		f.Anonymous = decl.Anonymous()
		s.Fields = append(s.Fields, f)
	}
	if !tagged {
		return nil
	}
	s.UpdateID()
	return s
}

// UpdateID recalculates the struct ID from the current signature.
func (s *Struct) UpdateID() {
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
	if s.IDName == "" {
		s.IDName = "binaryID" + s.Name
	}
}

func spaceToUnderscore(r rune) rune {
	if unicode.IsSpace(r) {
		return '_'
	}
	return r
}

// fromType creates a appropriate Type object from a types.Type.
func fromType(pkg *types.Package, from types.Type, tag tag, imports Imports) *Type {
	t := &Type{Name: strings.Map(spaceToUnderscore, path.Base(types.TypeString(pkg, from)))}
	if named, isNamed := from.(*types.Named); isNamed {
		from = from.Underlying()
		p := named.Obj().Pkg()
		if p != nil {
			imports[p.Path()] = struct{}{}
		}
	}
	t.Native = strings.Map(spaceToUnderscore, from.String())
	switch from := from.(type) {
	case *types.Basic:
		t.Kind = Native
		switch from.Kind() {
		case types.Int:
			t.Native = "int32"
		case types.Byte:
			t.Native = "uint8"
		case types.String:
			t.SkipMethod = "SkipString"
		}
		t.Method = strings.Title(t.Native)
		if t.Native != t.Name {
			t.Kind = Remap
		}
	case *types.Pointer:
		t.Kind = Pointer
		t.SubType = fromType(pkg, from.Elem(), tag, imports)
	case *types.Interface:
		t.Kind = Interface
	case *types.Slice:
		if tag.Flag("stream") {
			t.Kind = Stream
		} else {
			t.Kind = Array
		}
		t.SubType = fromType(pkg, from.Elem(), "", imports)
		switch elem := from.Elem().(type) {
		case *types.Basic:
			switch elem.Kind() {
			case types.Byte:
				t.Method = "Data"
			}
		}
	case *types.Array:
		t.Kind = StaticArray
		t.Length = int(from.Len())
		switch elem := from.Elem().(type) {
		case *types.Basic:
			switch elem.Kind() {
			case types.Byte:
				if from.Len() == binary.IDSize {
					t.Kind = Native
					t.SkipMethod = "SkipID"
					t.Method = "ID"
				}
			}
		}
		if t.Kind == StaticArray {
			t.SubType = fromType(pkg, from.Elem(), "", imports)
		}
	case *types.Map:
		t.Kind = Map
		t.KeyType = fromType(pkg, from.Key(), "", imports)
		t.SubType = fromType(pkg, from.Elem(), "", imports)
	default:
		t.Kind = Codeable
	}
	return t
}

type sortEntry struct {
	s       *Struct
	visited bool
}

func walk(name string, byname map[string]*sortEntry, structs []*Struct, i int) int {
	entry, found := byname[name]
	if !found || entry.visited {
		return i
	}
	entry.visited = true
	for _, f := range entry.s.Fields {
		i = walk(f.Type.Name, byname, structs, i)
		if f.Type.SubType != nil {
			i = walk(f.Type.SubType.Name, byname, structs, i)
		}
		if f.Type.KeyType != nil {
			i = walk(f.Type.KeyType.Name, byname, structs, i)
		}
	}
	structs[i] = entry.s
	return i + 1
}

// Sort is used to ensure stable ordering of Struct slices.
// This is to ensure automatically generated code has minimum diffs.
// The sort order is by Struct name, but guarantees dependencies occur first.
func Sort(structs []*Struct) {
	names := make(sort.StringSlice, len(structs))
	byname := make(map[string]*sortEntry, len(structs))
	for i, s := range structs {
		names[i] = s.Name
		byname[s.Name] = &sortEntry{s, false}
	}
	names.Sort()
	i := 0
	for _, name := range names {
		i = walk(name, byname, structs, i)
	}
}

func getTemplate(t *template.Template, name string) *template.Template {
	result := t.Lookup(name)
	if result == nil {
		panic(fmt.Errorf("Could not find template %s", name))
	}
	return result
}

type kindToTemplate map[Kind]*template.Template

func getTemplateMap(t *template.Template, prefix string) kindToTemplate {
	return kindToTemplate{
		Native:      getTemplate(t, prefix+"Native"),
		Remap:       getTemplate(t, prefix+"Remap"),
		Codeable:    getTemplate(t, prefix+"Codeable"),
		Pointer:     getTemplate(t, prefix+"Pointer"),
		Interface:   getTemplate(t, prefix+"Interface"),
		Array:       getTemplate(t, prefix+"Array"),
		StaticArray: getTemplate(t, prefix+"StaticArray"),
		Stream:      getTemplate(t, prefix+"Stream"),
		Map:         getTemplate(t, prefix+"Map"),
	}
}

func kindDispatch(table kindToTemplate, name string, t *Type) string {
	b := &bytes.Buffer{}
	if err := table[t.Kind].Execute(b, Field{name, t, false}); err != nil {
		panic(err)
	}
	return b.String()
}
