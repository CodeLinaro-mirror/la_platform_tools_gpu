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

func block(ctx *context, in *ast.Block, owner semantic.Node) *semantic.Block {
	out := &semantic.Block{AST: in}
	if in == nil {
		return out
	}
	ctx.with(semantic.VoidType, func() {
		ctx.scope.block = &out.Statements
		r := body(ctx, in.Statements, owner)
		if r != nil {
			ctx.addStatement(r)
		}
	})
	ctx.mappings[in] = out
	return out
}

// body is a resolve function that processes a list of statements and injects them
// into the context's current block.
// the final return statement, if present, is not injected, but returned from the
// function, as it often needs special handling depending on the owner of the
// statements
func body(ctx *context, in []ast.Node, owner semantic.Node) *semantic.Return {
	f, isFunction := owner.(*semantic.Function)
	var returnStatement *ast.Return
	// we need to check and strip the "return" if the function is supposed to have one
	if isFunction && !isVoid(f.Return.Type) {
		if len(in) == 0 {
			ctx.errorf(f.AST, "Missing return statement")
		} else if r, ok := in[len(in)-1].(*ast.Return); !ok {
			ctx.errorf(f.AST, "Last statement must be a return")
		} else {
			in = in[0 : len(in)-1]
			returnStatement = r
		}
	}
	// now process the non return statements
	for _, s := range in {
		ctx.addStatement(statement(ctx, s))
	}
	// and special case the return statement allowing access to the return parameter
	if returnStatement != nil {
		return return_(ctx, returnStatement, f)
	}
	return nil
}

func statement(ctx *context, in ast.Node) semantic.Node {
	switch in := in.(type) {
	case *ast.Assign:
		return assign(ctx, in)
	case *ast.DeclareLocal:
		return declareLocal(ctx, in)
	case *ast.Branch:
		return branch(ctx, in)
	case *ast.Switch:
		return switch_(ctx, in)
	case *ast.Iteration:
		return iteration(ctx, in)
	case *ast.Call:
		e := call(ctx, in)
		if !isVoid(e.ExpressionType()) {
			ctx.errorf(in, "function with return type as statement not allowed")
			return invalid{}
		}
		return e
	case *ast.Return:
		ctx.errorf(in, "unexpected return")
		return invalid{}
	case *ast.Generic:
		ctx.errorf(in.Name, "unexpected identifier %s", in.Name.Value)
		return invalid{}
	case *ast.Fence:
		return &semantic.Fence{AST: in, Explicit: true}
	default:
		ctx.errorf(in, "not a statement (%T)", in)
		return invalid{}
	}
}

func assign(ctx *context, in *ast.Assign) semantic.Node {
	lhs := expression(ctx, in.LHS)
	var rhs semantic.Expression
	ctx.with(lhs.ExpressionType(), func() {
		rhs = expression(ctx, in.RHS)
	})
	var out semantic.Node
	inferUnknown(ctx, lhs, rhs)
	lt := lhs.ExpressionType()
	rt := rhs.ExpressionType()
	if !assignable(lt, rt) {
		ctx.errorf(in, "cannot assign %s to %s", typename(rt), typename(lt))
	}
	switch lhs := lhs.(type) {
	case *semantic.ArrayIndex:
		out = &semantic.ArrayAssign{AST: in, To: lhs, Value: rhs, Operator: in.Operator}
	case *semantic.MapIndex:
		out = &semantic.MapAssign{AST: in, To: lhs, Value: rhs, Operator: in.Operator}
	case *semantic.SliceIndex:
		out = &semantic.SliceAssign{AST: in, To: lhs, Value: rhs, Operator: in.Operator}
	default:
		out = &semantic.Assign{AST: in, LHS: lhs, Operator: in.Operator, RHS: rhs}
	}
	ctx.mappings[in] = out
	return out
}

func addLocal(ctx *context, in *ast.DeclareLocal, name string, value semantic.Expression) *semantic.DeclareLocal {
	out := &semantic.DeclareLocal{AST: in}
	out.Local = &semantic.Local{
		Declaration: out,
		Named:       semantic.Named(name),
		Value:       value,
		Type:        value.ExpressionType(),
	}
	if isVoid(out.Local.Type) {
		ctx.errorf(in, "void in local declaration")
	}
	ctx.addNamed(out.Local)
	ctx.mappings[in] = out
	return out
}

func declareLocal(ctx *context, in *ast.DeclareLocal) *semantic.DeclareLocal {
	out := addLocal(ctx, in, in.Name.Value, expression(ctx, in.RHS))
	ctx.mappings[in] = out
	return out
}

func branch(ctx *context, in *ast.Branch) *semantic.Branch {
	out := &semantic.Branch{AST: in}
	out.Condition = expression(ctx, in.Condition)
	ct := out.Condition.ExpressionType()
	if ct == nil {
		ctx.errorf(in, "condition was not valid")
		return out
	}
	if !equal(ct, semantic.BoolType) {
		ctx.errorf(in, "if condition must be boolean (got %s)", typename(ct))
	}
	out.True = block(ctx, in.True, out)
	out.False = block(ctx, in.False, out)
	ctx.mappings[in] = out
	return out
}

func switch_(ctx *context, in *ast.Switch) *semantic.Switch {
	out := &semantic.Switch{AST: in}
	out.Value = expression(ctx, in.Value)
	vt := out.Value.ExpressionType()
	for _, c := range in.Cases {
		out.Cases = append(out.Cases, case_(ctx, c, vt))
	}
	if in.Default != nil {
		out.Default = block(ctx, in.Default.Block, out)
	}
	ctx.mappings[in] = out
	return out
}

// case_ translates Case in to a switch Case.
// vt is the resolved type of the switch value being compared against, and can
// be used to infer the case condition type.
func case_(ctx *context, in *ast.Case, vt semantic.Type) *semantic.Case {
	out := &semantic.Case{AST: in}
	ctx.with(vt, func() {
		for _, cond := range in.Conditions {
			exp := expression(ctx, cond)
			out.Conditions = append(out.Conditions, exp)
			ct := exp.ExpressionType()
			if !comparable(vt, ct) {
				ctx.errorf(cond, "switch value %s is not comparable with case condition %s", typename(vt), typename(ct))
			}
		}
	})
	out.Block = block(ctx, in.Block, out)
	ctx.mappings[in] = out
	return out
}

func iteration(ctx *context, in *ast.Iteration) *semantic.Iteration {
	v := &semantic.Local{Named: semantic.Named(in.Variable.Value)}
	ctx.mappings[in.Variable] = v
	out := &semantic.Iteration{AST: in, Iterator: v}
	out.Iterable = expression(ctx, in.Iterable)
	if b, ok := out.Iterable.(*semantic.BinaryOp); !ok {
		ctx.errorf(in, "iterable can only be range operator, got %T", b)
	} else if b.Operator != ast.OpRange {
		ctx.errorf(in, "iterable can only be range operator, got %s\n", b.Operator)
	}
	v.Type = out.Iterable.ExpressionType()
	ctx.with(semantic.VoidType, func() {
		ctx.addNamed(v)
		out.Block = block(ctx, in.Block, out)
	})
	ctx.mappings[in] = out
	return out
}

func return_(ctx *context, in *ast.Return, f *semantic.Function) *semantic.Return {
	out := &semantic.Return{AST: in}
	out.Function = f
	ctx.with(f.Return.Type, func() {
		out.Value = expression(ctx, in.Value)
	})
	inferUnknown(ctx, f.Return, out.Value)
	rt := out.Value.ExpressionType()
	if !assignable(f.Return.Type, rt) {
		ctx.errorf(in, "cannot assign %s to %s", typename(rt), typename(f.Return.Type))
	}
	ctx.mappings[in] = out
	return out
}
