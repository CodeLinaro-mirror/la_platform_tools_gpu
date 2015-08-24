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
	"android.googlesource.com/platform/tools/gpu/api/ast"
	"android.googlesource.com/platform/tools/gpu/api/semantic"
)

var (
	internals = map[string]func(*context, *ast.Call, *ast.Generic) semantic.Expression{}
)

func init() {
	internals["assert"] = assert
	internals["as"] = cast
	internals["new"] = new_
	internals["make"] = make_
	internals["clone"] = clone
	internals["read"] = read
	internals["write"] = write
	internals["copy"] = copy_
	internals["len"] = length
}

func internalCall(ctx *context, in *ast.Call) semantic.Expression {
	g, ok := in.Target.(*ast.Generic)
	if !ok {
		return nil
	}
	p, ok := internals[g.Name.Value]
	if !ok {
		return nil
	}
	return p(ctx, in, g)
}

func checkInternalFunc(ctx *context, in *ast.Call, g *ast.Generic, typeCount, paramCount int) bool {
	if typeCount >= 0 && len(g.Arguments) != typeCount {
		ctx.errorf(in, "wrong number of types to %s, expected %d got %v", g.Name.Value, typeCount, len(g.Arguments))
		return false
	}
	if paramCount >= 0 && len(in.Arguments) != paramCount {
		ctx.errorf(in, "wrong number of arguments to %s, expected %d got %v", g.Name.Value, paramCount, len(in.Arguments))
		return false
	}
	return true
}

func assert(ctx *context, in *ast.Call, g *ast.Generic) semantic.Expression {
	if !checkInternalFunc(ctx, in, g, 0, 1) {
		return invalid{}
	}
	condition := expression(ctx, in.Arguments[0])
	t := condition.ExpressionType()
	if !equal(t, semantic.BoolType) {
		ctx.errorf(in, "assert expression must be a bool, got %s", typename(t))
	}
	out := &semantic.Assert{AST: in, Condition: condition}
	ctx.mappings[in] = out
	return out
}

func cast(ctx *context, in *ast.Call, g *ast.Generic) semantic.Expression {
	if !checkInternalFunc(ctx, in, g, 1, 1) {
		return invalid{}
	}
	t := type_(ctx, g.Arguments[0])
	var obj semantic.Expression
	ctx.with(t, func() {
		obj = expression(ctx, in.Arguments[0])
	})
	if equal(t, obj.ExpressionType()) {
		ctx.mappings[in] = obj
		return obj
	}
	if !castable(obj.ExpressionType(), t) {
		ctx.errorf(in, "cannot cast from %s to %s", typename(obj.ExpressionType()), typename(t))
	}
	out := &semantic.Cast{AST: in, Object: obj, Type: t}
	ctx.mappings[in] = out
	return out
}

func new_(ctx *context, in *ast.Call, g *ast.Generic) semantic.Expression {
	if !checkInternalFunc(ctx, in, g, 1, -1) {
		return invalid{}
	}
	t := type_(ctx, g.Arguments[0])
	rt := getRefType(ctx, in, t)
	if c, isclass := t.(*semantic.Class); isclass {
		i := classInitializer(ctx, c, in)
		out := &semantic.Create{AST: in, Type: rt, Initializer: i}
		ctx.mappings[in] = out
		return out
	}
	out := &semantic.New{AST: in, Type: rt}
	ctx.mappings[in] = out
	return out
}

func make_(ctx *context, in *ast.Call, g *ast.Generic) semantic.Expression {
	if !checkInternalFunc(ctx, in, g, 1, 1) {
		return invalid{}
	}
	t := type_(ctx, g.Arguments[0])

	var size semantic.Expression
	ctx.with(semantic.Uint64Type, func() {
		size = expression(ctx, in.Arguments[0])
	})
	size = castTo(ctx, in.Arguments[0], size, semantic.Uint64Type)
	out := &semantic.Make{AST: in, Type: getSliceType(ctx, in, t), Size: size}
	ctx.mappings[in] = out
	return out
}

func clone(ctx *context, in *ast.Call, g *ast.Generic) semantic.Expression {
	if !checkInternalFunc(ctx, in, g, 0, 1) {
		return invalid{}
	}
	slice := expression(ctx, in.Arguments[0])
	st, ok := slice.ExpressionType().(*semantic.Slice)
	if !ok {
		ctx.errorf(in, "%s only works on slice types, got type %v", g.Name.Value, typename(slice.ExpressionType()))
		return invalid{}
	}
	out := &semantic.Clone{AST: in, Slice: slice, Type: st}
	ctx.mappings[in] = out
	return out
}

func read(ctx *context, in *ast.Call, g *ast.Generic) semantic.Expression {
	if !checkInternalFunc(ctx, in, g, 0, 1) {
		return invalid{}
	}
	slice := expression(ctx, in.Arguments[0])
	if _, ok := slice.ExpressionType().(*semantic.Slice); !ok {
		ctx.errorf(in, "%s only works on slice types, got type %v", g.Name.Value, typename(slice.ExpressionType()))
		return invalid{}
	}
	out := &semantic.Read{AST: in, Slice: slice}
	ctx.mappings[in] = out
	return out
}

func write(ctx *context, in *ast.Call, g *ast.Generic) semantic.Expression {
	if !checkInternalFunc(ctx, in, g, 0, 1) {
		return invalid{}
	}
	slice := expression(ctx, in.Arguments[0])
	if _, ok := slice.ExpressionType().(*semantic.Slice); !ok {
		ctx.errorf(in, "%s only works on slice types, got type %v", g.Name.Value, typename(slice.ExpressionType()))
		return invalid{}
	}
	out := &semantic.Write{AST: in, Slice: slice}
	ctx.mappings[in] = out
	return out
}

func copy_(ctx *context, in *ast.Call, g *ast.Generic) semantic.Expression {
	if !checkInternalFunc(ctx, in, g, 0, 2) {
		return invalid{}
	}
	src := expression(ctx, in.Arguments[1])
	srct, ok := src.ExpressionType().(*semantic.Slice)
	if !ok {
		ctx.errorf(in, "%s only works on slice types, got type %v", g.Name.Value, typename(src.ExpressionType()))
		return invalid{}
	}
	dst := expression(ctx, in.Arguments[0])
	if !equal(srct, dst.ExpressionType()) {
		ctx.errorf(in, "%s slice types do't match, got %v and %v", g.Name.Value, typename(srct), typename(dst.ExpressionType()))
		return invalid{}
	}
	out := &semantic.Copy{AST: in, Src: src, Dst: dst}
	ctx.mappings[in] = out
	return out
}

func length(ctx *context, in *ast.Call, g *ast.Generic) semantic.Expression {
	if !checkInternalFunc(ctx, in, g, 0, 1) {
		return invalid{}
	}
	obj := expression(ctx, in.Arguments[0])
	at := obj.ExpressionType()
	if at == nil {
		ctx.errorf(in, "object was not valid")
		return invalid{}
	}
	ok := false
	ty := baseType(at)
	switch ty := ty.(type) {
	case *semantic.Slice:
		ok = true
	case *semantic.Map:
		ok = true
	case *semantic.Builtin:
		if ty == semantic.StringType {
			ok = true
		}
	}
	if !ok {
		ctx.errorf(in, "len cannot work out length of type %s", typename(at))
		return invalid{}
	}
	infer := baseType(ctx.scope.inferType)
	var t semantic.Type = semantic.Int32Type
	switch infer {
	case semantic.Int32Type, semantic.Uint32Type, semantic.Int64Type, semantic.Uint64Type:
		t = infer
	default:
		t = semantic.Int32Type
	}
	out := &semantic.Length{AST: in, Object: obj, Type: t}
	ctx.mappings[in] = out
	return out
}
