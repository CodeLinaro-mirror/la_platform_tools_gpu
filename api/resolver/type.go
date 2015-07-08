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

const (
	RefSuffix     = "ʳ"
	SliceSuffix   = "ˢ"
	ConstSuffix   = "ᶜ"
	PointerSuffix = "ᵖ"
	ArraySuffix   = "ᵃ"
	MapSuffix     = "ᵐ"
	TypeInfix     = "ː"
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
		var size uint32
		var sizeExpr semantic.Expression
		ctx.with(semantic.Uint32Type, func() {
			e := expression(ctx, in.Index)
			switch ty := e.(type) {
			case *semantic.DefinitionUsage:
				sizeExpr = ty.Definition
				e = ty.Expression
			case *semantic.EnumEntry:
				sizeExpr = e
				e = semantic.Uint32Value(ty.Value)
			default:
				sizeExpr = e
			}
			if n, ok := e.(semantic.Uint32Value); ok {
				size = uint32(n)
			} else {
				ctx.errorf(in.Index, "Array dimension must be a constant number, got %T", e)
			}
		})
		return getStaticArrayType(ctx, in, of, size, sizeExpr)
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
	if a, ok := out.(*semantic.Alias); ok {
		if a.To == nil {
			a.To = type_(ctx, a.AST.To)
		}
		return a.To
	}
	return out
}

func getMapType(ctx *context, at ast.Node, kt, vt semantic.Type) *semantic.Map {
	name := fmt.Sprintf("%s%s%s%s", strings.Title(kt.Name()), TypeInfix, vt.Name(), MapSuffix)
	for _, m := range ctx.api.Maps {
		if equal(kt, m.KeyType) && equal(vt, m.ValueType) {
			ctx.mappings[at] = m
			return m
		}
	}
	out := &semantic.Map{
		Named:     semantic.Named(name),
		KeyType:   kt,
		ValueType: vt,
	}
	ctx.api.Maps = append(ctx.api.Maps, out)
	ctx.mappings[at] = out
	return out
}

func getStaticArrayType(ctx *context, at ast.Node, of semantic.Type, size uint32, sizeExpr semantic.Expression) *semantic.StaticArray {
	name := fmt.Sprintf("%s%s%d%s", strings.Title(of.Name()), TypeInfix, size, ArraySuffix)
	for _, a := range ctx.api.StaticArrays {
		if equal(a.ValueType, of) && a.Size == size {
			ctx.mappings[at] = a
			return a
		}
	}
	out := &semantic.StaticArray{
		Named:     semantic.Named(name),
		ValueType: of,
		Size:      size,
		SizeExpr:  sizeExpr,
	}
	ctx.api.StaticArrays = append(ctx.api.StaticArrays, out)
	ctx.mappings[at] = out
	return out
}

func getRefType(ctx *context, at ast.Node, to semantic.Type) *semantic.Reference {
	name := strings.Title(to.Name()) + RefSuffix
	for _, p := range ctx.api.References {
		if equal(to, p.To) {
			ctx.mappings[at] = p
			return p
		}
	}
	out := &semantic.Reference{
		Named: semantic.Named(name),
		To:    to,
	}
	ctx.api.References = append(ctx.api.References, out)
	ctx.mappings[at] = out
	return out
}

func getPointerType(ctx *context, at ast.Node, to semantic.Type, constant bool) *semantic.Pointer {
	name := strings.Title(to.Name())
	if constant {
		name += ConstSuffix
	}
	name += PointerSuffix
	for _, p := range ctx.api.Pointers {
		if equal(to, p.To) {
			ctx.mappings[at] = p
			return p
		}
	}
	out := &semantic.Pointer{
		Named: semantic.Named(name),
		To:    to,
		Const: constant,
	}
	ctx.api.Pointers = append(ctx.api.Pointers, out)
	ctx.mappings[at] = out

	out.Slice = getSliceType(ctx, at, to)

	return out
}

func getSliceType(ctx *context, at ast.Node, to semantic.Type) *semantic.Slice {
	name := strings.Title(to.Name()) + SliceSuffix
	for _, s := range ctx.api.Slices {
		if equal(to, s.To) {
			ctx.mappings[at] = s
			return s
		}
	}
	out := &semantic.Slice{
		Named: semantic.Named(name),
		To:    to,
	}
	ctx.api.Slices = append(ctx.api.Slices, out)
	ctx.mappings[at] = out

	out.Pointer = getPointerType(ctx, at, to, false)

	return out
}

func definition(ctx *context, out *semantic.Definition) {
	in := out.AST
	out.Annotations = annotations(ctx, in.Annotations)
	out.Docs = findDocumentation(in.CST)
	out.Expression = expression(ctx, in.Expression)
	ctx.mappings[in] = out
}

func enum(ctx *context, out *semantic.Enum) {
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
			Named: semantic.Named(e.Name.Value),
			Docs:  findDocumentation(e.CST),
			Value: uint32(v),
		}
		out.Entries = append(out.Entries, entry)
		ctx.mappings[e] = entry
	}
	for _, extends := range in.Extends {
		t := ctx.findType(extends, extends.Value)
		if e, ok := t.(*semantic.Enum); !ok {
			ctx.errorf(extends, "non enum entry %s in extension list", typename(t))
		} else {
			out.Extends = append(out.Extends, e)
			ctx.mappings[extends] = e
		}
	}
	ctx.mappings[in] = out
}

func enumEntries(ctx *context, out, scan *semantic.Enum) {
	for _, e := range scan.Entries {
		if out != scan {
			e = &semantic.EnumEntry{
				AST:   e.AST,
				Named: e.Named,
				Value: e.Value,
			}
		}
		semantic.Add(out, e)
		ctx.addNamed(e)
	}
	for _, extends := range scan.Extends {
		enumEntries(ctx, out, extends)
	}
}

func class(ctx *context, out *semantic.Class) {
	in := out.AST
	out.Docs = findDocumentation(in.CST)
	out.Annotations = annotations(ctx, in.Annotations)
	out.Fields = make([]*semantic.Field, len(in.Fields))
	for i, f := range in.Fields {
		field := field(ctx, f, out)
		out.Fields[i] = field
	}
	ctx.mappings[in] = out
}

func field(ctx *context, in *ast.Field, class *semantic.Class) *semantic.Field {
	out := &semantic.Field{AST: in, Named: semantic.Named(in.Name.Value)}
	semantic.Add(class, out)
	out.Docs = findDocumentation(in.CST)
	out.Annotations = annotations(ctx, in.Annotations)
	out.Type = type_(ctx, in.Type)
	if isVoid(out.Type) {
		ctx.errorf(in, "void typed field %s on class %s", out.Name(), class.Name())
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
		return semantic.VoidType
	}
	t, ok := api.Member(in.Name.Value).(semantic.Type)
	if !ok {
		ctx.errorf(in, "%s not a type in %s", in.Name.Value, in.From.Value)
		return semantic.VoidType
	}
	return t
}

func typename(t semantic.Type) string {
	switch t := t.(type) {
	case nil:
		return "missing"
	case *semantic.StaticArray:
		return fmt.Sprintf("%s[%d]", typename(t.ValueType), t.Size)
	default:
		return t.Name()
	}
}
