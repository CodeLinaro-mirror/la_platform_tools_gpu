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
	"strconv"
	"strings"

	"android.googlesource.com/platform/tools/gpu/api/ast"
	"android.googlesource.com/platform/tools/gpu/api/semantic"
)

func type_(ctx *context, in interface{}) semantic.Type {
	switch in := in.(type) {
	case *ast.Identifier:
		return simpleType(ctx, in)
	case *ast.MapType:
		return mapType(ctx, in)
	case *ast.ArrayType:
		return arrayType(ctx, in)
	case *ast.StaticArrayType:
		return staticArrayType(ctx, in)
	case *ast.PointerType:
		return pointerType(ctx, in)
	default:
		ctx.errorf(in, "Unhandled typeref %T found", in)
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
	if a, ok := out.(*alias); ok {
		if a.To == nil {
			a.To = type_(ctx, a.AST.To)
		}
		return a.To
	}
	return out
}

const functionResult = "result"

func mapType(ctx *context, in *ast.MapType) *semantic.Map {
	kt := type_(ctx, in.KeyType)
	vt := type_(ctx, in.ValueType)
	name := strings.Title(vt.Typename()) + "_" + kt.Typename() + "Map"
	for _, m := range ctx.api.Maps {
		if m.Name == name {
			if !equal(kt, m.KeyType) {
				ctx.icef(in, "Map %s found with non matching key, got %s expected %s", name, typename(m.KeyType), typename(kt))
			}
			if !equal(vt, m.ValueType) {
				ctx.icef(in, "Map %s found with non matching value, got %s expected %s", name, typename(m.ValueType), typename(vt))
			}
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
				{Name: functionResult, Type: vt},
			},
		},
		&semantic.Function{Name: "GetOrError",
			FullParameters: []*semantic.Parameter{{},
				{Name: "key", Type: kt},
				{Name: functionResult, Type: vt},
			},
		},
		&semantic.Function{Name: "Set",
			FullParameters: []*semantic.Parameter{{},
				{Name: "key", Type: kt},
				{Name: "value", Type: vt},
			},
		},
		&semantic.Function{Name: "Contains",
			FullParameters: []*semantic.Parameter{{},
				{Name: "key", Type: kt},
				{Name: functionResult, Type: semantic.BoolType},
			},
		},
		&semantic.Function{Name: "Count",
			FullParameters: []*semantic.Parameter{{},
				{Name: functionResult, Type: semantic.Int32Type},
			},
		},
		&semantic.Function{Name: "Delete",
			FullParameters: []*semantic.Parameter{{},
				{Name: "key", Type: kt},
			},
		},
		&semantic.Function{Name: "Range",
			FullParameters: []*semantic.Parameter{{},
				{Name: functionResult, Type: semantic.AnyType},
			},
		},
	} {
		f.Owner = out
		f.This = f.FullParameters[0]
		f.This.Name = "self"
		f.This.Type = out
		last := f.FullParameters[len(f.FullParameters)-1]
		if last.Name == functionResult {
			f.Return = f.FullParameters[len(f.FullParameters)-1]
			f.Return.Output = true
			f.Outputs = append(f.Outputs, f.Return)
		} else {
			f.Return = &semantic.Parameter{Name: functionResult, Type: semantic.VoidType}
		}
		out.Members[f.Name] = f
	}
	ctx.api.Maps = append(ctx.api.Maps, out)
	return out
}

func arrayType(ctx *context, in *ast.ArrayType) *semantic.Array {
	vt := type_(ctx, in.ValueType)
	name := strings.Title(vt.Typename()) + "Array"
	for _, a := range ctx.api.Arrays {
		if a.Name == name {
			if !equal(vt, a.ValueType) {
				ctx.icef(in, "Array %s found with non matching value, got %s expected %s", name, typename(a.ValueType), typename(vt))
			}
			return a
		}
	}
	out := &semantic.Array{
		Name:      name,
		ValueType: vt,
	}
	ctx.api.Arrays = append(ctx.api.Arrays, out)
	return out
}

func staticArrayType(ctx *context, in *ast.StaticArrayType) *semantic.StaticArray {
	vt := type_(ctx, in.ValueType)
	name := strings.Title(vt.Typename()) + "StaticArray"
	for _, a := range ctx.api.StaticArrays {
		if a.Name == name {
			if !equal(vt, a.ValueType) {
				ctx.icef(in, "Static array %s found with non matching value, got %s expected %s", name, typename(a.ValueType), typename(vt))
			}
			return a
		}
	}
	out := &semantic.StaticArray{
		Name:      name,
		ValueType: vt,
	}
	ctx.with(semantic.Uint32Type, func() {
		for _, d := range in.Dimensions {
			e := expression(ctx, d)
			if n, ok := e.(semantic.Uint32Value); ok {
				out.Dimensions = append(out.Dimensions, n)
			} else {
				ctx.errorf(in, "Array dimension must be a constant number, got %T", e)
			}
		}
	})
	ctx.api.StaticArrays = append(ctx.api.StaticArrays, out)
	return out
}

func pointerType(ctx *context, in *ast.PointerType) *semantic.Pointer {
	vt := type_(ctx, in.To)
	name := strings.Title(vt.Typename()) + "Ref"
	for _, p := range ctx.api.Pointers {
		if p.Name == name {
			if !equal(vt, p.To) {
				ctx.icef(in, "Pointer %s found with non matching value, got %s expected %s", name, typename(p.To), typename(vt))
			}
			return p
		}
	}
	out := &semantic.Pointer{
		Name: name,
		To:   vt,
	}
	ctx.api.Pointers = append(ctx.api.Pointers, out)
	return out
}

func enum(ctx *context, out *semantic.Enum) {
	if len(out.AllEntries) > 0 {
		// Already resolved.
		return
	}
	in := out.AST
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
			Value: uint32(v),
		}
		out.Entries = append(out.Entries, entry)
		out.AllEntries = append(out.AllEntries, entry)
	}
	for _, extends := range in.Extends {
		t := ctx.findType(extends, extends.Value)
		if e, ok := t.(*semantic.Enum); !ok {
			ctx.errorf(extends, "non enum entry %s in extension list", t.Typename())
		} else {
			out.Extends = append(out.Extends, e)
			enum(ctx, e)
			for _, entry := range e.AllEntries {
				copy := *entry
				copy.Enum = out
				out.AllEntries = append(out.AllEntries, &copy)
			}
		}
	}
	for _, entry := range out.AllEntries {
		ctx.add(entry.Name, entry)
	}
}

func class(ctx *context, out *semantic.Class) {
	in := out.AST
	out.Annotations = annotations(ctx, in.Annotations)
	for _, extends := range in.Extends {
		t := ctx.findType(extends, extends.Value)
		if c, ok := t.(*semantic.Class); !ok {
			ctx.errorf(extends, "non class entry %s in extension list", t.Typename())
		} else {
			out.Extends = append(out.Extends, c)
			c.ExtendedBy = append(c.ExtendedBy, out)
		}
	}
	out.Fields = make([]*semantic.Field, len(in.Fields))
	for i, f := range in.Fields {
		field := field(ctx, f, out)
		out.Fields[i] = field
		out.Members[field.Name] = field
	}
}

func field(ctx *context, in *ast.Field, class *semantic.Class) *semantic.Field {
	out := &semantic.Field{AST: in, Name: in.Name.Value, Class: class}
	out.Annotations = annotations(ctx, in.Annotations)
	out.Type = type_(ctx, in.Type)
	if in.Default != nil {
		ctx.with(out.Type, func() {
			out.Default = expression(ctx, in.Default)
		})
		dt := out.Default.ExpressionType()
		if !assignable(out.Type, dt) {
			ctx.errorf(in, "cannot assign %s to %s", typename(dt), typename(out.Type))
		}
	}
	return out
}

func pseudonym(ctx *context, out *semantic.Pseudonym) {
	in := out.AST
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
