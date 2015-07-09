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
	"unicode"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/any"
	"android.googlesource.com/platform/tools/gpu/binary/schema"
	"golang.org/x/tools/go/types"
)

const (
	indent       = "»"
	classPrefix  = "⊹"
	memberPrefix = "∍"
)

type Imports map[string]struct{}

type File struct {
	Name       string
	Import     string
	IsTest     bool
	Path       string
	Directives map[string]string
	Structs    []*Struct
	Constants  schema.Constants
	Imports    Imports
}

const (
	binaryPackage  = "android.googlesource.com/platform/tools/gpu/binary"
	binaryGenerate = binaryPackage + ".Generate"
)

// Struct is a description of an encodable struct.
// Signature includes the package, name and name and type of all the fields.
// Any change to the Signature will cause the ID to change.
type Struct struct {
	schema.Class
	IDName    string // The name to give the ID of the type.
	Signature string // The full string type signature of the Struct.
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
	b := findBinaryObject(pkg)
	t := n.Type().Underlying().(*types.Struct)
	s := &Struct{Class: schema.Class{
		Name:    n.Name(),
		Package: pkg.Name(),
	}}
	tagged := false
	for i := 0; i < t.NumFields(); i++ {
		decl := t.Field(i)
		tag := tag(t.Tag(i))
		if decl.Anonymous() &&
			decl.Type().String() == binaryGenerate &&
			!tag.Flag("disable") {
			tagged = true
			s.IDName = tag.Get("id")
			continue
		}
		f := schema.Field{}
		if !decl.Anonymous() {
			f.Declared = decl.Name()
		}
		f.Type = fromType(pkg, decl.Type(), tag, imports, b)
		delete(imports, pkg.Path())
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
		fmt.Fprintf(b, " %s:%s", f.Name(), f.Type)
	}
	fmt.Fprint(b, " }")
	s.Signature = b.String()
	s.TypeID = binary.NewID([]byte(s.Signature))
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

// findBinaryObject looks for the binary.Object type in the imports, returning it
// if it is found, or nil if it is not.
func findBinaryObject(pkg *types.Package) *types.Interface {
	for _, p := range pkg.Imports() {
		if p.Path() == binaryPackage {
			if o := p.Scope().Lookup("Object"); o != nil {
				return o.Type().Underlying().(*types.Interface)
			}
		}
	}
	return nil
}

// fromType creates a appropriate schema.Type object from a types.Type.
func fromType(pkg *types.Package, from types.Type, tag tag, imports Imports, binObj *types.Interface) schema.Type {
	alias := ""
	fullname := types.TypeString(pkg, from) // fully-qualified name including full package path
	name := strings.Map(spaceToUnderscore, path.Base(fullname))
	if named, isNamed := from.(*types.Named); isNamed {
		alias = name
		from = from.Underlying()
		p := named.Obj().Pkg()
		if p != nil && p != pkg {
			imports[p.Path()] = struct{}{}
		}
	}
	gotype := strings.Map(spaceToUnderscore, from.String())
	switch from := from.(type) {
	case *types.Basic:
		switch from.Kind() {
		case types.Int:
			return &schema.Primitive{Name: name, Method: schema.Int32}
		case types.Byte:
			return &schema.Primitive{Name: name, Method: schema.Uint8}
		case types.Rune:
			return &schema.Primitive{Name: name, Method: schema.Int32}
		default:
			m, err := schema.ParseMethod(strings.Title(gotype))
			if err != nil {
				return &schema.Primitive{Name: fmt.Sprintf("%s_bad_%s", name, gotype), Method: schema.String}
			}
			return &schema.Primitive{Name: name, Method: m}
		}
	case *types.Pointer:
		return &schema.Pointer{Type: fromType(pkg, from.Elem(), tag, imports, binObj)}
	case *types.Interface:
		if binObj != nil && !types.Implements(from, binObj) {
			return &any.Any{}
		} else {
			return &schema.Interface{Name: name}
		}
	case *types.Slice:
		vt := fromType(pkg, from.Elem(), "", imports, binObj)
		if tag.Flag("stream") {
			return &schema.Stream{Alias: alias, ValueType: vt}
		}
		return &schema.Slice{Alias: alias, ValueType: vt}
	case *types.Array:
		length := uint32(from.Len())
		if elem, ok := from.Elem().(*types.Basic); ok {
			if elem.Kind() == types.Byte && length == binary.IDSize {
				return &schema.Primitive{Name: name, Method: schema.ID}
			}
		}
		return &schema.Array{
			Alias:     alias,
			ValueType: fromType(pkg, from.Elem(), "", imports, binObj),
			Size:      length,
		}
	case *types.Map:
		return &schema.Map{
			Alias:     alias,
			KeyType:   fromType(pkg, from.Key(), "", imports, binObj),
			ValueType: fromType(pkg, from.Elem(), "", imports, binObj),
		}
	default:
		return &schema.Struct{Name: name}
	}
}

type sortEntry struct {
	s       *Struct
	visited bool
}

func walkType(t schema.Type, byname map[string]*sortEntry, structs []*Struct, i int) int {
	switch t := t.(type) {
	case *schema.Primitive:
	case *schema.Struct:
		i = walk(t.Name, byname, structs, i)
	case *schema.Interface:
		i = walk(t.Name, byname, structs, i)
	case *schema.Pointer:
		i = walkType(t.Type, byname, structs, i)
	case *schema.Array:
		i = walkType(t.ValueType, byname, structs, i)
	case *schema.Slice:
		i = walkType(t.ValueType, byname, structs, i)
	case *schema.Stream:
		i = walkType(t.ValueType, byname, structs, i)
	case *schema.Map:
		i = walkType(t.KeyType, byname, structs, i)
		i = walkType(t.ValueType, byname, structs, i)
	}
	return i
}

func walk(name string, byname map[string]*sortEntry, structs []*Struct, i int) int {
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
