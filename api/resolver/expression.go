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
	case *ast.Generic:
		id := identifier(ctx, in.Name)
		if len(in.Arguments) > 0 {
			ctx.errorf(in, "identifier %s does not support type arguments", in.Name.Value)
		}
		return id
	case *ast.Group:
		return expression(ctx, in.Expression)
	case *ast.Unknown:
		return &semantic.Unknown{AST: in}
	case *ast.Number:
		return number(ctx, in)
	case *ast.Bool:
		return semantic.BoolValue(in.Value)
	case *ast.String:
		return semantic.StringValue(in.Value)
	case *ast.Null:
		return semantic.Null{AST: in, Type: ctx.scope.inferType}
	default:
		ctx.icef(in, "Unhandled expression type %T found", in)
		return invalid{}
	}
}

func call(ctx *context, in *ast.Call) semantic.Expression {
	if b := internalCall(ctx, in); b != nil {
		return b
	}
	if c := classCall(ctx, in); c != nil {
		return c
	}
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
			if i >= len(args) {
				break // will have already errored
			}
			if args[i] == nil {
				continue // will have already errored
			}
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

func castToU64(ctx *context, in ast.Node, expr semantic.Expression) semantic.Expression {
	ty := expr.ExpressionType()
	if equal(ty, semantic.Uint64Type) {
		return expr
	}
	if !castable(ty, semantic.Uint64Type) {
		ctx.errorf(in, "cannot cast %s to u64", typename(ty))
	}
	return &semantic.Cast{Object: expr, Type: semantic.Uint64Type}
}

func index(ctx *context, in *ast.Index) semantic.Expression {
	object := expression(ctx, in.Object)
	at := baseType(object.ExpressionType())
	var index semantic.Expression

	switch at := at.(type) {
	case *semantic.Pointer:
		ctx.with(semantic.Uint64Type, func() {
			index = expression(ctx, in.Index)
		})
		if bop, ok := index.(*semantic.BinaryOp); ok && bop.Operator == ast.OpSlice {
			// pointer[a:b]
			bop.LHS = castToU64(ctx, bop.AST.LHS, bop.LHS)
			bop.RHS = castToU64(ctx, bop.AST.RHS, bop.RHS)
			out := &semantic.PointerRange{AST: in, Pointer: object, Type: at.Slice, Range: bop}
			ctx.mappings[in] = out
			return out
		}
		if n, ok := index.(semantic.Uint64Value); ok && n == 0 {
			// pointer[0]
			// TODO: clean up the magical 0 index on pointers
			r := &semantic.BinaryOp{LHS: n, Operator: ast.OpSlice, RHS: n + 1}
			slice := &semantic.PointerRange{AST: in, Pointer: object, Type: at.Slice, Range: r}
			out := &semantic.SliceIndex{AST: in, Slice: slice, Type: at.Slice, Index: n}
			ctx.mappings[in] = out
			return out
		}
		ctx.errorf(in, "type %s not valid slicing pointer", typename(index.ExpressionType()))
		return invalid{}
	case *semantic.Slice:
		ctx.with(semantic.Uint64Type, func() {
			index = expression(ctx, in.Index)
		})
		if bop, ok := index.(*semantic.BinaryOp); ok && bop.Operator == ast.OpSlice {
			// slice[a:b]
			bop.LHS = castToU64(ctx, bop.AST.LHS, bop.LHS)
			bop.RHS = castToU64(ctx, bop.AST.RHS, bop.RHS)
			out := &semantic.SliceRange{AST: in, Slice: object, Type: at, Range: bop}
			ctx.mappings[in] = out
			return out
		}
		// slice[a]
		index = castToU64(ctx, in, index)
		out := &semantic.SliceIndex{AST: in, Slice: object, Type: at, Index: index}
		ctx.mappings[in] = out
		return out
	case *semantic.Map:
		// map[k]
		ctx.with(at.KeyType, func() {
			index = expression(ctx, in.Index)
		})
		it := index.ExpressionType()
		if !comparable(it, at.KeyType) {
			ctx.errorf(in, "type %s not valid indexing map", typename(it))
		}
		out := &semantic.MapIndex{AST: in, Map: object, Type: at, Index: index}
		ctx.mappings[in] = out
		return out
	}
	ctx.errorf(in, "index operation on non indexable type %s", typename(at))
	return invalid{}
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

func classCall(ctx *context, in *ast.Call) semantic.Expression {
	g, ok := in.Target.(*ast.Generic)
	if !ok {
		return nil
	}
	t := ctx.findType(in, g.Name.Value)
	class, ok := t.(*semantic.Class)
	if !ok {
		return nil
	}
	return classInitializer(ctx, class, in)
}

func classInitializer(ctx *context, class *semantic.Class, in *ast.Call) *semantic.ClassInitializer {
	out := &semantic.ClassInitializer{AST: in, Class: class}
	ctx.mappings[in] = out
	if len(in.Arguments) == 0 {
		return out
	}
	if _, named := in.Arguments[0].(*ast.NamedArg); named {
		for _, a := range in.Arguments {
			n, ok := a.(*ast.NamedArg)
			if !ok {
				ctx.errorf(a, "class %s has no field %s", class.Name, n.Name.Value)
				return out
			}
			m := class.Member(n.Name.Value)
			if m == nil {
				ctx.errorf(n.Name, "class %s has no field %s", class.Name, n.Name.Value)
				return out
			}
			f, ok := m.(*semantic.Field)
			if !ok {
				ctx.errorf(n.Name, "member %s of class %s is not a field [%T]", n.Name.Value, class.Name, m)
				return out
			}
			ctx.mappings[n.Name] = f
			out.Fields = append(out.Fields, fieldInitializer(ctx, class, f, n.Value))
		}
		return out
	}
	if len(in.Arguments) > len(class.Fields) {
		ctx.errorf(in, "too many arguments to class %s constructor, expected %d got %d", class.Name, len(class.Fields), len(in.Arguments))
		return out
	}
	for i, a := range in.Arguments {
		out.Fields = append(out.Fields, fieldInitializer(ctx, class, class.Fields[i], a))
	}
	return out
}

func fieldInitializer(ctx *context, class *semantic.Class, field *semantic.Field, in ast.Node) *semantic.FieldInitializer {
	out := &semantic.FieldInitializer{AST: in, Field: field}
	ctx.with(field.Type, func() {
		out.Value = expression(ctx, in)
	})
	ft := field.Type
	vt := out.Value.ExpressionType()
	if !assignable(ft, vt) {
		ctx.errorf(in, "field %s cannot assign %s to %s", field.Name, typename(vt), typename(ft))
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
