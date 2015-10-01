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
	binary.Entity
	Tags       Tags // The tags associated with the type.
	Raw        *types.Struct
	unresolved []unresolved // The set of unresolved schema.Struct objects.
}

type unresolved struct {
	s *schema.Struct
	t *types.Struct
}

func (m *Module) addStruct(n *types.TypeName) {
	t := n.Type().Underlying().(*types.Struct)
	s := &Struct{
		Entity: binary.Entity{
			Display:  n.Name(),
			Package:  m.Source.Types.Name(),
			Exported: n.Exported(),
		},
		Raw: t,
	}
	tagged := false
	var invalidStruct error
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
		f := binary.Field{}
		if !decl.Anonymous() {
			f.Declared = decl.Name()
		}
		defer func() {
			if r := recover(); r != nil && invalidStruct == nil {
				invalidStruct = fmt.Errorf("Handling %v.%v gave error %v", s.Name(), f.Name(), r)
			}
		}()
		f.Type = m.fromType(decl.Type(), s, tags)
		s.Fields = append(s.Fields, f)
	}
	if tagged {
		if invalidStruct != nil {
			panic(invalidStruct)
		}
		s.Identity = s.Tag("identity", s.Display)
		if s.Display == s.Identity {
			s.Display = ""
		}
		s.Version = s.Tag("version", "")
		m.Structs = append(m.Structs, s)
	}
}

// Tag returns the named tag if present, missing otherwise.
func (s *Struct) Tag(name string, missing string) string {
	result := s.Tags.Get(name)
	if result == "" {
		result = missing
	}
	return result
}

// IDName returns the name to give the ID of the type.
func (s *Struct) IDName() string {
	return s.Tag("id", "binaryID"+s.Name())
}

// ID returns the type id for the entity.
func (s *Struct) ID() binary.ID {
	return s.TypeID
}

// HasStructTag returns true if any struct in the module has the named tag.
func (m *Module) HasStructTag(name string) bool {
	for _, s := range m.Structs {
		result := s.Tags.Get(name)
		if result != "" {
			return true
		}
	}
	return false
}

// UpdateID recalculates the struct ID from the current signature.
func (s *Struct) UpdateID() {
	s.TypeID = binary.NewID([]byte(s.Entity.Signature()))
}

type sortEntry struct {
	s       *Struct
	visited bool
}

func walkType(t binary.Type, byname map[string]*sortEntry, structs []*Struct, i int) int {
	switch t := t.(type) {
	case *schema.Primitive:
	case *schema.Struct:
		i = walkStructs(t.String(), byname, structs, i)
	case *schema.Interface:
		i = walkStructs(t.Name, byname, structs, i)
	case *schema.Variant:
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

func (m *Module) finaliseStructs() {
	names := make(sort.StringSlice, len(m.Structs))
	byname := make(map[string]*sortEntry, len(m.Structs))
	for i, s := range m.Structs {
		s.UpdateID()
		name := s.Name()
		names[i] = name
		byname[name] = &sortEntry{s, false}
	}
	names.Sort()
	i := 0
	for _, name := range names {
		i = walkStructs(name, byname, m.Structs, i)
	}
}
