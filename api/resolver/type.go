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

package resolver

import (
	"fmt"
	"strconv"
	"strings"

	"android.googlesource.com/platform/tools/gpu/api/ast"
	"android.googlesource.com/platform/tools/gpu/api/semantic"
)

func type_(ctx *context, in interface{}) semantic.Type {
	switch in := in.(type) {
	case *ast.Generic:
		return genericType(ctx, in)
	case *ast.IndexedType:
		of := type_(ctx, in.ValueType)
		if in.Index == nil {
			return getSliceType(ctx, in, of)
		}
		size := uint32(0)
		ctx.with(semantic.Uint32Type, func() {
			e := expression(ctx, in.Index)
			if n, ok := e.(semantic.Uint32Value); ok {
				size = uint32(n)
			} else {
				ctx.errorf(in.Index, "Array dimension must be a constant number, got %T", e)
			}
		})
		return getStaticArrayType(ctx, in, of, size)
	case *ast.PointerType:
		to := type_(ctx, in.To)
		return getPointerType(ctx, in, to, in.Const)
	case *ast.Imported:
		return importedType(ctx, in)
	case ast.Node:
		ctx.errorf(in, "Unhandled typeref %T found", in)
		return semantic.VoidType
	default:
		ctx.icef(nil, "Non-node (%T) typeref found", in)
		return semantic.VoidType
	}
}

func genericType(ctx *context, in *ast.Generic) semantic.Type {
	switch in.Name.Value {
	case "map":
		if len(in.Arguments) != 2 {
			ctx.errorf(in, "Map requires 2 args, got %d", len(in.Arguments))
			return semantic.VoidType
		}
		kt := type_(ctx, in.Arguments[0])
		vt := type_(ctx, in.Arguments[1])
		return getMapType(ctx, in, kt, vt)
	case "ref":
		if len(in.Arguments) != 1 {
			ctx.errorf(in, "Ref requires 1 arg, got %d", len(in.Arguments))
			return semantic.VoidType
		}
		vt := type_(ctx, in.Arguments[0])
		return getRefType(ctx, in, vt)
	default:
		if len(in.Arguments) != 0 {
			ctx.errorf(in, "Type %s is not parameterised, got %d", in.Name, len(in.Arguments))
			return semantic.VoidType
		}
		return getSimpleType(ctx, in.Name)
	}
}

func getSimpleType(ctx *context, in *ast.Identifier) semantic.Type {
	name := in.Value
	out := ctx.findType(in, name)
	if out == nil {
		ctx.errorf(in, "Type %s not found", name)
		return semantic.VoidType
	}
	ctx.mappings[in] = out
	if a, ok := out.(*Alias); ok {
		if a.To == nil {
			a.To = type_(ctx, a.AST.To)
		}
		return a.To
	}
	return out
}

func getMapType(ctx *context, at ast.Node, kt, vt semantic.Type) *semantic.Map {
	name := strings.Title(vt.Typename()) + "_" + kt.Typename() + "Map"
	for _, m := range ctx.api.Maps {
		if m.Name == name {
			if !equal(kt, m.KeyType) {
				ctx.icef(at, "Map %s found with non matching key, got %s expected %s", name, typename(m.KeyType), typename(kt))
			}
			if !equal(vt, m.ValueType) {
				ctx.icef(at, "Map %s found with non matching value, got %s expected %s", name, typename(m.ValueType), typename(vt))
			}
			ctx.mappings[at] = m
			return m
		}
	}
	out := &semantic.Map{
		Name:      name,
		KeyType:   kt,
		ValueType: vt,
		Members:   semantic.Members{},
	}
	ctx.api.Maps = append(ctx.api.Maps, out)
	ctx.mappings[at] = out
	return out
}

func getStaticArrayType(ctx *context, at ast.Node, of semantic.Type, size uint32) *semantic.StaticArray {
	name := fmt.Sprintf("%sStaticArray_%d", of.Typename(), size)
	for _, a := range ctx.api.StaticArrays {
		if a.Name == name {
			if !equal(a.ValueType, of) {
				ctx.icef(at, "Static array %s found with non matching value, got %s expected %s",
					a.Name, typename(a.ValueType), typename(of))
			}
			ctx.mappings[at] = a
			return a
		}
	}
	out := &semantic.StaticArray{
		Name:      name,
		ValueType: of,
		Size:      size,
	}
	ctx.api.StaticArrays = append(ctx.api.StaticArrays, out)
	ctx.mappings[at] = out
	return out
}

func getPointerType(ctx *context, at ast.Node, to semantic.Type, constant bool) *semantic.Pointer {
	name := strings.Title(to.Typename())
	if constant {
		name += "Const"
	}
	name += "Pointer"
	for _, p := range ctx.api.Pointers {
		if p.Name == name {
			if !equal(to, p.To) {
				ctx.icef(at, "Pointer %s found with non matching value, got %s expected %s", name, typename(p.To), typename(to))
			}
			ctx.mappings[at] = p
			return p
		}
	}
	out := &semantic.Pointer{
		Name:  name,
		To:    to,
		Const: constant,
	}
	ctx.api.Pointers = append(ctx.api.Pointers, out)
	ctx.mappings[at] = out
	return out
}

func getRefType(ctx *context, at ast.Node, to semantic.Type) *semantic.Reference {
	name := strings.Title(to.Typename())
	name += "Ref"
	for _, p := range ctx.api.References {
		if p.Name == name {
			if !equal(to, p.To) {
				ctx.icef(at, "ref %s found with non matching value, got %s expected %s", name, typename(p.To), typename(to))
			}
			ctx.mappings[at] = p
			return p
		}
	}
	out := &semantic.Reference{
		Name: name,
		To:   to,
	}
	ctx.api.References = append(ctx.api.References, out)
	ctx.mappings[at] = out
	return out
}

func getSliceType(ctx *context, at ast.Node, to semantic.Type) *semantic.Slice {
	name := strings.Title(to.Typename())
	name += "Slice"
	for _, s := range ctx.api.Slices {
		if s.Name == name {
			if !equal(to, s.To) {
				ctx.icef(at, "Slice %s found with non matching value, got %s expected %s", name, typename(s.To), typename(to))
			}
			ctx.mappings[at] = s
			return s
		}
	}
	out := &semantic.Slice{
		Name: name,
		To:   to,
	}
	ctx.api.Slices = append(ctx.api.Slices, out)
	ctx.mappings[at] = out
	return out
}

func enum(ctx *context, out *semantic.Enum) {
	if len(out.AllEntries) > 0 {
		// Already resolved.
		return
	}
	in := out.AST
	out.Docs = findDocumentation(in.CST)
	out.Annotations = annotations(ctx, in.Annotations)
	out.IsBitfield = in.IsBitfield
	for _, e := range in.Entries {
		v, err := strconv.ParseUint(e.Value.Value, 0, 32)
		if err != nil {
			ctx.errorf(e, "could not parse %s as uint32", e.Value)
			continue
		}
		entry := &semantic.EnumEntry{
			AST:   e,
			Enum:  out,
			Name:  e.Name.Value,
			Docs:  findDocumentation(e.CST),
			Value: uint32(v),
		}
		out.Entries = append(out.Entries, entry)
		out.AllEntries = append(out.AllEntries, entry)
		ctx.mappings[e] = entry
	}
	for _, extends := range in.Extends {
		t := ctx.findType(extends, extends.Value)
		if e, ok := t.(*semantic.Enum); !ok {
			ctx.errorf(extends, "non enum entry %s in extension list", typename(t))
		} else {
			out.Extends = append(out.Extends, e)
			enum(ctx, e)
			for _, entry := range e.AllEntries {
				copy := *entry
				copy.Enum = out
				out.AllEntries = append(out.AllEntries, &copy)
			}
			ctx.mappings[extends] = e
		}
	}
	for _, entry := range out.AllEntries {
		ctx.add(entry.Name, entry)
	}
	ctx.mappings[in] = out
}

func class(ctx *context, out *semantic.Class) {
	in := out.AST
	out.Docs = findDocumentation(in.CST)
	out.Annotations = annotations(ctx, in.Annotations)
	for _, extends := range in.Extends {
		t := ctx.findType(extends, extends.Value)
		if c, ok := t.(*semantic.Class); !ok {
			ctx.errorf(extends, "non class entry %s in extension list", typename(t))
		} else {
			out.Extends = append(out.Extends, c)
			c.ExtendedBy = append(c.ExtendedBy, out)
			ctx.mappings[extends] = c
		}
	}
	out.Fields = make([]*semantic.Field, len(in.Fields))
	for i, f := range in.Fields {
		field := field(ctx, f, out)
		out.Fields[i] = field
		out.Members[field.Name] = field
	}
	ctx.mappings[in] = out
}

func field(ctx *context, in *ast.Field, class *semantic.Class) *semantic.Field {
	out := &semantic.Field{AST: in, Name: in.Name.Value, Class: class}
	out.Docs = findDocumentation(in.CST)
	out.Annotations = annotations(ctx, in.Annotations)
	out.Type = type_(ctx, in.Type)
	if isVoid(out.Type) {
		ctx.errorf(in, "void typed field %s on class %s", out.Name, out.Class.Name)
	}
	if in.Default != nil {
		ctx.with(out.Type, func() {
			out.Default = expression(ctx, in.Default)
		})
		dt := out.Default.ExpressionType()
		if !assignable(out.Type, dt) {
			ctx.errorf(in, "cannot assign %s to %s", typename(dt), typename(out.Type))
		}
	}
	ctx.mappings[in] = out
	return out
}

func pseudonym(ctx *context, out *semantic.Pseudonym) {
	in := out.AST
	out.Docs = findDocumentation(in.CST)
	out.Annotations = annotations(ctx, in.Annotations)
	out.To = type_(ctx, in.To)
}

func importedType(ctx *context, in *ast.Imported) semantic.Type {
	api, ok := ctx.get(in, in.From.Value).(*semantic.API)
	if !ok {
		ctx.errorf(in, "%s not an imported api", in.From.Value)
		return nil
	}
	t, ok := api.Member(in.Name.Value).(semantic.Type)
	if !ok {
		ctx.errorf(in, "%s not a type in %s", in.Name.Value, in.From.Value)
		return nil
	}
	return t
}

func typename(e semantic.Type) string {
	if e == nil {
		return "missing"
	} else {
		return e.Typename()
	}
}
