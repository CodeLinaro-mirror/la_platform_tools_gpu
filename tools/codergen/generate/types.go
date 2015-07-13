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
	"path"
	"strings"
	"unicode"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/any"
	"android.googlesource.com/platform/tools/gpu/binary/schema"
	"golang.org/x/tools/go/types"
)

const (
	binaryPackage  = "android.googlesource.com/platform/tools/gpu/binary"
	binaryGenerate = binaryPackage + ".Generate"
)

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

func spaceToUnderscore(r rune) rune {
	if unicode.IsSpace(r) {
		return '_'
	}
	return r
}

// fromType creates a appropriate schema.Type object from a types.Type.
func fromType(pkg *types.Package, from types.Type, tags Tags, imports Imports, binObj *types.Interface) schema.Type {
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
		return &schema.Pointer{Type: fromType(pkg, from.Elem(), tags, imports, binObj)}
	case *types.Interface:
		if binObj != nil && !types.Implements(from, binObj) {
			return &any.Any{}
		} else {
			return &schema.Interface{Name: name}
		}
	case *types.Slice:
		vt := fromType(pkg, from.Elem(), "", imports, binObj)
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
