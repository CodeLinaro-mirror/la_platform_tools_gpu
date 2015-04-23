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
	case *ast.Identifier:
		return simpleType(ctx, in)
	case *ast.GenericType:
		return genericType(ctx, in)
	case *ast.IndexedType:
		return indexedType(ctx, in)
	case *ast.PointerType:
		return pointerType(ctx, in)
	case ast.Node:
		ctx.errorf(in, "Unhandled typeref %T found", in)
		return semantic.VoidType
	default:
		ctx.icef(nil, "Non-node (%T) typeref found", in)
		return semantic.VoidType
	}
}

func simpleType(ctx *context, in *ast.Identifier) semantic.Type {
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

func mapType(ctx *context, in *ast.GenericType) semantic.Type {
	if len(in.Args) != 2 {
		ctx.errorf(in, "Map requires 2 args, got %d", len(in.Args))
		return semantic.VoidType
	}
	kt := type_(ctx, in.Args[0])
	vt := type_(ctx, in.Args[1])
	name := strings.Title(vt.Typename()) + "_" + kt.Typename() + "Map"
	for _, m := range ctx.api.Maps {
		if m.Name == name {
			if !equal(kt, m.KeyType) {
				ctx.icef(in, "Map %s found with non matching key, got %s expected %s", name, typename(m.KeyType), typename(kt))
			}
			if !equal(vt, m.ValueType) {
				ctx.icef(in, "Map %s found with non matching value, got %s expected %s", name, typename(m.ValueType), typename(vt))
			}
			ctx.mappings[in] = m
			return m
		}
	}
	out := &semantic.Map{
		Name:      name,
		KeyType:   kt,
		ValueType: vt,
		Members:   semantic.Members{},
	}
	for _, f := range []*semantic.Function{
		&semantic.Function{Name: "Get",
			FullParameters: []*semantic.Parameter{{},
				{Name: "key", Type: kt},
				{Name: "value", Type: vt},
				{Type: vt},
			},
		},
		&semantic.Function{Name: "Delete",
			FullParameters: []*semantic.Parameter{{},
				{Name: "key", Type: kt},
			},
		},
		&semantic.Function{Name: "Range",
			FullParameters: []*semantic.Parameter{{},
				{Type: semantic.AnyType},
			},
		},
	} {
		f.Owner = out
		f.This = f.FullParameters[0]
		f.This.Name = "self"
		f.This.Type = out
		last := f.FullParameters[len(f.FullParameters)-1]
		if last.Name == "" {
			f.Return = f.FullParameters[len(f.FullParameters)-1]
			f.Return.Output = true
			f.Outputs = append(f.Outputs, f.Return)
		} else {
			f.Return = &semantic.Parameter{Type: semantic.VoidType}
		}
		out.Members[f.Name] = f
	}
	ctx.api.Maps = append(ctx.api.Maps, out)
	ctx.mappings[in] = out
	return out
}

func getArrayType(ctx *context, at ast.Node, of semantic.Type) *semantic.Array {
	name := strings.Title(of.Typename()) + "Array"
	for _, a := range ctx.api.Arrays {
		if a.Name == name {
			if !equal(of, a.ValueType) {
				ctx.icef(at, "Array %s found with non matching value, got %s expected %s", name, typename(a.ValueType), typename(of))
			}
			ctx.mappings[at] = a
			return a
		}
	}
	out := &semantic.Array{
		Name:      name,
		ValueType: of,
	}
	ctx.api.Arrays = append(ctx.api.Arrays, out)
	ctx.mappings[at] = out
	return out
}

func arrayType(ctx *context, in *ast.GenericType) semantic.Type {
	if len(in.Args) != 1 {
		ctx.errorf(in, "Array requires 1 arg, got %d", len(in.Args))
		return semantic.VoidType
	}
	return getArrayType(ctx, in, type_(ctx, in.Args[0]))
}

func genericType(ctx *context, in *ast.GenericType) semantic.Type {
	switch in.Generic.Value {
	case ast.KeywordArray:
		return arrayType(ctx, in)
	case ast.KeywordMap:
		return mapType(ctx, in)
	default:
		ctx.icef(in, "Generic type %s not handled", in.Generic.Value)
		return semantic.VoidType
	}
}

func indexedType(ctx *context, in *ast.IndexedType) semantic.Type {
	if in.Index == nil {
		return getPointerType(ctx, in, type_(ctx, in.ValueType), true)
	}
	out := &semantic.StaticArray{ValueType: type_(ctx, in.ValueType)}
	ctx.with(semantic.Uint32Type, func() {
		e := expression(ctx, in.Index)
		if n, ok := e.(semantic.Uint32Value); ok {
			out.Size = uint32(n)
		} else {
			ctx.errorf(in, "Array dimension must be a constant number, got %T", e)
		}
	})
	out.Name = fmt.Sprintf("%sStaticArray_%d", out.Size)
	for _, a := range ctx.api.StaticArrays {
		if a.Name == out.Name {
			if !equal(out.ValueType, a.ValueType) {
				ctx.icef(in, "Static array %s found with non matching value, got %s expected %s",
					out.Name, typename(a.ValueType), typename(out.ValueType))
			}
			ctx.mappings[in] = a
			return a
		}
	}

	ctx.api.StaticArrays = append(ctx.api.StaticArrays, out)
	ctx.mappings[in] = out
	return out
}

func getPointerType(ctx *context, at ast.Node, to semantic.Type, array bool) *semantic.Pointer {
	name := strings.Title(to.Typename())
	if array {
		name += "Array"
	} else {
		name += "Ref"
	}
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
		Array: array,
	}
	ctx.api.Pointers = append(ctx.api.Pointers, out)
	ctx.mappings[at] = out
	return out
}

func pointerType(ctx *context, in *ast.PointerType) *semantic.Pointer {
	return getPointerType(ctx, in, type_(ctx, in.To), false)
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

func typename(e semantic.Type) string {
	if e == nil {
		return "missing"
	} else {
		return e.Typename()
	}
}
