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
	"sort"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/schema"
	"golang.org/x/tools/go/types"
)

// Struct is a description of an encodable struct.
// Signature includes the package, name and name and type of all the fields.
// Any change to the Signature will cause the ID to change.
type Struct struct {
	schema.Class
	Tags      Tags   // The tags associated with the type.
	Signature string // The full string type signature of the Struct.
}

// FromTypename creates and initializes a Struct from a types.Typename.
// It assumes that the typename will map to a types.Struct, and adds all the
// fields of that struct to the Struct information.
func NewStruct(pkg *types.Package, n *types.TypeName, imports Imports) *Struct {
	b := findBinaryObject(pkg)
	t := n.Type().Underlying().(*types.Struct)
	s := &Struct{Class: schema.Class{
		Name:    n.Name(),
		Package: pkg.Name(),
	}}
	tagged := false
	for i := 0; i < t.NumFields(); i++ {
		decl := t.Field(i)
		tags := Tags(t.Tag(i))
		if decl.Anonymous() &&
			decl.Type().String() == binaryGenerate &&
			!tags.Flag("disable") {
			tagged = true
			s.Tags = tags
			continue
		}
		f := schema.Field{}
		if !decl.Anonymous() {
			f.Declared = decl.Name()
		}
		f.Type = fromType(pkg, decl.Type(), tags, imports, b)
		delete(imports, pkg.Path())
		s.Fields = append(s.Fields, f)
	}
	if !tagged {
		return nil
	}
	s.UpdateID()
	return s
}

// IDName returns the name to give the ID of the type.
func (s *Struct) IDName() string {
	name := s.Tags.Get("id")
	if name == "" {
		name = "binaryID" + s.Name
	}
	return name
}

// UpdateID recalculates the struct ID from the current signature.
func (s *Struct) UpdateID() {
	b := &bytes.Buffer{}
	fmt.Fprintf(b, "struct %s.%s {", s.Package, s.Name)
	for i, f := range s.Fields {
		if i != 0 {
			fmt.Fprint(b, ",")
		}
		fmt.Fprintf(b, " %s:%s", f.Name(), f.Type)
	}
	fmt.Fprint(b, " }")
	s.Signature = b.String()
	s.TypeID = binary.NewID([]byte(s.Signature))
}

type sortEntry struct {
	s       *Struct
	visited bool
}

func walkType(t schema.Type, byname map[string]*sortEntry, structs []*Struct, i int) int {
	switch t := t.(type) {
	case *schema.Primitive:
	case *schema.Struct:
		i = walkStructs(t.Name, byname, structs, i)
	case *schema.Interface:
		i = walkStructs(t.Name, byname, structs, i)
	case *schema.Pointer:
		i = walkType(t.Type, byname, structs, i)
	case *schema.Array:
		i = walkType(t.ValueType, byname, structs, i)
	case *schema.Slice:
		i = walkType(t.ValueType, byname, structs, i)
	case *schema.Map:
		i = walkType(t.KeyType, byname, structs, i)
		i = walkType(t.ValueType, byname, structs, i)
	}
	return i
}

func walkStructs(name string, byname map[string]*sortEntry, structs []*Struct, i int) int {
	entry, found := byname[name]
	if !found || entry.visited {
		return i
	}
	entry.visited = true
	for _, f := range entry.s.Fields {
		i = walkType(f.Type, byname, structs, i)
	}
	structs[i] = entry.s
	return i + 1
}

// sortStructs is used to ensure stable ordering of Struct slices.
// This is to ensure automatically generated code has minimum diffs.
// The sort order is by Struct name, but guarantees dependencies occur first.
func sortStructs(structs []*Struct) {
	names := make(sort.StringSlice, len(structs))
	byname := make(map[string]*sortEntry, len(structs))
	for i, s := range structs {
		names[i] = s.Name
		byname[s.Name] = &sortEntry{s, false}
	}
	names.Sort()
	i := 0
	for _, name := range names {
		i = walkStructs(name, byname, structs, i)
	}
}
