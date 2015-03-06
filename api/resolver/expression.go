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

	"android.googlesource.com/platform/tools/gpu/api/ast"
	"android.googlesource.com/platform/tools/gpu/api/semantic"
)

type invalid struct{}

func (invalid) ExpressionType() semantic.Type { return semantic.VoidType }

// expression translates the ast expression to a semantic expression.
func expression(ctx *context, in ast.Node) semantic.Expression {
	switch in := in.(type) {
	case *ast.UnaryOp:
		return unaryOp(ctx, in)
	case *ast.BinaryOp:
		return binaryOp(ctx, in)
	case *ast.Call:
		return call(ctx, in)
	case *ast.Switch:
		return select_(ctx, in)
	case *ast.Member:
		return member(ctx, in)
	case *ast.Index:
		return index(ctx, in)
	case *ast.Identifier:
		return identifier(ctx, in)
	case *ast.ClassInitializer:
		return classInitializer(ctx, in)
	case *ast.New:
		return new(ctx, in)
	case *ast.Group:
		return expression(ctx, in.Expression)
	case *ast.Cast:
		return cast(ctx, in)
	case *ast.Length:
		return length(ctx, in)
	case *ast.Unknown:
		return &semantic.Unknown{AST: in}
	case *ast.Number:
		return number(ctx, in)
	case *ast.Bool:
		return semantic.BoolValue(in.Value)
	case *ast.String:
		return semantic.StringValue(in.Value)
	default:
		ctx.icef(in, "Unhandled expression type %T found", in)
		return invalid{}
	}
}

func call(ctx *context, in *ast.Call) semantic.Expression {
	target := expression(ctx, in.Target)
	switch target := target.(type) {
	case *macroStub:
		return macroCall(ctx, in, target)
	case *semantic.Callable:
		return functionCall(ctx, in, target)
	default:
		ctx.errorf(in, "Invalid method call target %T found", target)
		return invalid{}
	}
}

func callArguments(ctx *context, at ast.Node, in []ast.Node, params []*semantic.Parameter, name string) []semantic.Expression {
	out := []semantic.Expression{}
	if len(params) != len(in) {
		ctx.errorf(at, "wrong number of arguments to %s, expected %v got %v", name, len(params), len(in))
		return out
	}
	for i, a := range in {
		p := params[i]
		ctx.with(p.Type, func() {
			arg := expression(ctx, a)
			at := arg.ExpressionType()
			out = append(out, arg)
			if !assignable(p.Type, at) {
				ctx.errorf(a, "argument %d to %s is wrong type, expected %s got %s", i, name, typename(p.Type), typename(at))
			}
		})
	}
	return out
}

func functionCall(ctx *context, in *ast.Call, target *semantic.Callable) *semantic.Call {
	out := &semantic.Call{AST: in, Target: target, Type: semantic.VoidType}
	params := target.Function.FullParameters
	if target.Object != nil {
		if target.Function.This == nil {
			ctx.errorf(in, "method call on non method %s of %T", target.Function.Name, target.Object)
			return out
		}
		params = params[1:len(params)]
	}
	if !isVoid(target.Function.Return.Type) {
		params = params[0 : len(params)-1]
	}
	out.Arguments = callArguments(ctx, in, in.Arguments, params, target.Function.Name)
	out.Type = out.Target.Function.Return.Type
	ctx.mappings[in] = out
	return out
}

func macroCall(ctx *context, in *ast.Call, stub *macroStub) semantic.Expression {
	if ctx.scope.block == nil {
		ctx.errorf(in, "macro call outside of block scope")
		return invalid{}
	}
	// generate a globally unique naming prefix to prevent symbol collisions
	prefix := fmt.Sprintf("%s_%v_", stub.function.Name, ctx.uid())
	params := stub.function.CallParameters()
	var result *semantic.DeclareLocal
	args := callArguments(ctx, in, in.Arguments, params, stub.function.Name)
	// switch scopes back to the one the macro was declared in to prevent symbol leak
	callScope := ctx.scope
	ctx.scope = stub.scope
	defer func() { ctx.scope = callScope }()
	ctx.with(semantic.VoidType, func() {
		// put the block back so we inject directly in to it
		ctx.scope.block = callScope.block
		// replace parameters with a new uniquely named local variable
		for i, p := range params {
			l := addLocal(ctx, nil, p.Name, args[i])
			// set the unique name after symbol table injection
			// this means that the lookups inside the macro match the local correctly
			// but the semantic graph as a globally unique name
			l.Local.Name = prefix + l.Local.Name
			ctx.addStatement(l)
		}
		// evaluate the macro body in place
		r := body(ctx, stub.function.AST.Block.Statements, stub.function)
		// substitute the return statement for a local assignment
		if r != nil {
			result = addLocal(ctx, nil, prefix+"result", r.Value)
			ctx.addStatement(result)
		}
	})
	if result == nil {
		return invalid{}
	}
	return result.Local
}

func select_(ctx *context, in *ast.Switch) *semantic.Select {
	out := &semantic.Select{AST: in}
	out.Type = nil
	out.Value = expression(ctx, in.Value)
	vt := out.Value.ExpressionType()
	for _, c := range in.Cases {
		e := choice(ctx, c, vt)
		out.Choices = append(out.Choices, e)
		if out.Type == nil {
			out.Type = e.Expression.ExpressionType()
		} else if !equal(out.Type, e.Expression.ExpressionType()) {
			// TODO: This could be a common ancestor type instead?
			out.Type = semantic.VoidType
		}
	}
	if out.Type == nil {
		ctx.errorf(in, "could not determine type of switch")
		out.Type = semantic.VoidType
	}
	ctx.mappings[in] = out
	return out
}

// choice translates Case in to a select Choice.
// vt is the resolved type of the select value being compared against, and can
// be used to infer the choice condition type.
func choice(ctx *context, in *ast.Case, vt semantic.Type) *semantic.Choice {
	out := &semantic.Choice{AST: in}
	ctx.with(vt, func() {
		for _, cond := range in.Conditions {
			exp := expression(ctx, cond)
			out.Conditions = append(out.Conditions, exp)
			ct := exp.ExpressionType()
			if !comparable(vt, ct) {
				ctx.errorf(cond, "select value %s is not comparable with choice condition %s", typename(vt), typename(ct))
			}
		}
	})
	if len(in.Block.Statements) != 1 {
		ctx.errorf(in, "switch case is not a single expression")
		out.Expression = invalid{}
		return out
	}
	out.Expression = expression(ctx, in.Block.Statements[0])
	ctx.mappings[in] = out
	return out
}

func member(ctx *context, in *ast.Member) semantic.Expression {
	obj := expression(ctx, in.Object)
	ot := obj.ExpressionType()
	entry := ot.Member(in.Name.Value)
	if entry == nil {
		ctx.errorf(in, "%s is not a member of %s", in.Name.Value, typename(ot))
		return invalid{}
	}
	var out semantic.Expression
	switch entry := entry.(type) {
	case *semantic.Field:
		out = &semantic.Member{AST: in, Object: obj, Field: entry}
	case *semantic.Function:
		out = &semantic.Callable{Object: obj, Function: entry}
	default:
		ctx.errorf(in, "Invalid member lookup type %T found", entry)
		return invalid{}
	}
	ctx.mappings[in] = out
	return out
}

func index(ctx *context, in *ast.Index) semantic.Expression {
	object := expression(ctx, in.Object)
	at := object.ExpressionType()
	switch at := at.(type) {
	case *semantic.Array:
		out := &semantic.ArrayIndex{AST: in, Array: object}
		ctx.with(semantic.Int32Type, func() {
			out.Index = expression(ctx, in.Index)
		})
		it := out.Index.ExpressionType()
		if !equal(it, semantic.Int32Type) && !equal(it, semantic.Uint32Type) {
			ctx.errorf(in, "type %s not valid indexing array", typename(it))
		}
		out.ValueType = at.ValueType
		ctx.mappings[in] = out
		return out
	case *semantic.Map:
		out := &semantic.MapIndex{AST: in, Map: object}
		ctx.with(at.KeyType, func() {
			out.Index = expression(ctx, in.Index)
		})
		it := out.Index.ExpressionType()
		if !comparable(it, at.KeyType) {
			ctx.errorf(in, "type %s not valid indexing map", typename(it))
		}
		out.ValueType = at.ValueType
		ctx.mappings[in] = out
		return out
	default:
		ctx.errorf(in, "index operation on non indexable type %s", typename(at))
		return invalid{}
	}
}

func identifier(ctx *context, in *ast.Identifier) semantic.Expression {
	out := ctx.get(in, in.Value)
	switch out := out.(type) {
	case *semantic.Function:
		s := &semantic.Callable{Function: out}
		ctx.mappings[in] = s
		return s
	case semantic.Expression:
		ctx.mappings[in] = out
		return out
	default:
		ctx.errorf(in, "Symbol %s was non expression %T", in.Value, out)
		return invalid{}
	}
}

func classInitializer(ctx *context, in *ast.ClassInitializer) *semantic.ClassInitializer {
	out := &semantic.ClassInitializer{AST: in}
	t := type_(ctx, in.Class)
	class, ok := t.(*semantic.Class)
	if !ok {
		ctx.errorf(in.Class, "non class entry %s in extension list", typename(t))
		return out
	}
	out.Class = class
	for _, f := range in.Fields {
		out.Fields = append(out.Fields, fieldInitializer(ctx, out.Class, f))
	}
	ctx.mappings[in] = out
	return out
}

func fieldInitializer(ctx *context, class *semantic.Class, in *ast.FieldInitializer) *semantic.FieldInitializer {
	out := &semantic.FieldInitializer{AST: in}
	m := class.Member(in.Name.Value)
	if m == nil {
		ctx.errorf(in.Name, "class %s has no field %s", class.Name, in.Name.Value)
		return out
	}
	field, ok := m.(*semantic.Field)
	if !ok {
		ctx.errorf(in.Name, "member %s of class %s is not a field [%T]", in.Name.Value, class.Name, m)
		return out
	}
	ctx.mappings[in.Name] = field
	out.Field = field
	ctx.with(field.Type, func() {
		out.Value = expression(ctx, in.Value)
	})
	ft := out.Field.Type
	vt := out.Value.ExpressionType()
	if !assignable(ft, vt) {
		ctx.errorf(in, "field %s cannot assign %s to %s", out.Field.Name, typename(vt), typename(ft))
	}
	ctx.mappings[in] = out
	return out
}

func new(ctx *context, in *ast.New) *semantic.New {
	out := &semantic.New{AST: in}
	out.Initializer = classInitializer(ctx, in.ClassInitializer)
	out.Type = getPointerType(ctx, in, out.Initializer.Class)
	ctx.mappings[in] = out
	return out
}

func cast(ctx *context, in *ast.Cast) semantic.Expression {
	t := type_(ctx, in.Type)
	var obj semantic.Expression
	ctx.with(t, func() {
		obj = expression(ctx, in.Object)
	})
	if equal(t, obj.ExpressionType()) {
		ctx.mappings[in] = obj
		return obj
	}
	out := &semantic.Cast{AST: in, Object: obj, Type: t}
	if !castable(obj.ExpressionType(), out.Type) {
		ctx.errorf(in, "cannot cast from %s to %s", typename(obj.ExpressionType()), typename(out.Type))
	}
	ctx.mappings[in] = out
	return out
}

func length(ctx *context, in *ast.Length) *semantic.Length {
	out := &semantic.Length{AST: in}
	out.Object = expression(ctx, in.Object)
	at := out.Object.ExpressionType()
	if at == nil {
		ctx.errorf(in, "condition was not valid")
		return out
	}
	ok := false
	switch at.(type) {
	case *semantic.Array:
		ok = true
	case *semantic.Map:
		ok = true
	case *semantic.Builtin:
		if at == semantic.StringType || at == semantic.PointerType {
			ok = true
		}
	}
	if !ok {
		ctx.errorf(in, "len cannot work out length of type %s", typename(at))
		return out
	}
	infer := baseType(ctx.scope.inferType)
	switch infer {
	case semantic.Int32Type, semantic.Uint32Type, semantic.Int64Type, semantic.Uint64Type:
		out.Type = infer
	default:
		out.Type = semantic.Int32Type
	}
	ctx.mappings[in] = out
	return out
}

func number(ctx *context, in *ast.Number) semantic.Expression {
	infer := baseType(ctx.scope.inferType)
	out := inferNumber(ctx, in, infer)
	if out != nil {
		if infer == ctx.scope.inferType {
			ctx.mappings[in] = out
			return out
		}
		return &semantic.Cast{Type: ctx.scope.inferType, Object: out}
	}
	if v, err := strconv.ParseInt(in.Value, 0, 32); err == nil {
		s := semantic.Int32Value(v)
		ctx.mappings[in] = s
		return s
	}
	if v, err := strconv.ParseFloat(in.Value, 64); err == nil {
		s := semantic.Float64Value(v)
		ctx.mappings[in] = s
		return s
	}
	ctx.errorf(in, "could not parse %s as a number (%s)", in.Value, typename(infer))
	return invalid{}
}
