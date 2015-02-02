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

func block(ctx *context, in *ast.Block, owner interface{}) *semantic.Block {
	out := &semantic.Block{AST: in}
	if in == nil {
		return out
	}
	ctx.with(semantic.VoidType, func() {
		f, isFunction := owner.(*semantic.Function)
		var returnStatement *ast.Return
		statements := in.Statements
		// we need to check and strip the "return" if the function is supposed to have one
		if isFunction && f.Return.Type != semantic.VoidType {
			if len(statements) == 0 {
				ctx.errorf(in, "Missing return statement")
			} else if r, ok := statements[len(statements)-1].(*ast.Return); !ok {
				ctx.errorf(in, "Last statement must be a return")
			} else {
				statements = statements[0 : len(statements)-1]
				returnStatement = r
			}
		}
		// now process the non return statements
		for _, s := range statements {
			out.Statements = append(out.Statements, statement(ctx, s))
		}
		// and special case the return statement allowing access to the return parameter
		if returnStatement != nil {
			out.Statements = append(out.Statements, return_(ctx, returnStatement, f))
		}
	})
	return out
}

func statement(ctx *context, in interface{}) interface{} {
	switch in := in.(type) {
	case *ast.Assert:
		return assert(ctx, in)
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
	case *ast.Return:
		ctx.errorf(in, "unexpected return")
		return invalid{}
	default:
		return expression(ctx, in)
	}
}

func assert(ctx *context, in *ast.Assert) *semantic.Assert {
	out := &semantic.Assert{AST: in}
	out.Condition = expression(ctx, in.Condition)
	t := out.Condition.ExpressionType()
	if !equal(t, semantic.BoolType) {
		ctx.errorf(in, "assert expression must be a bool, got %s", t.Typename())
	}
	return out
}

func assign(ctx *context, in *ast.Assign) *semantic.Assign {
	out := &semantic.Assign{AST: in}
	out.LHS = expression(ctx, in.LHS)
	ctx.with(out.LHS.ExpressionType(), func() {
		out.RHS = expression(ctx, in.RHS)
	})
	inferUnknown(ctx, out.LHS, out.RHS)
	rt := out.RHS.ExpressionType()
	lt := out.LHS.ExpressionType()
	if !assignable(lt, rt) {
		ctx.errorf(in, "cannot assign %s to %s", typename(rt), typename(lt))
	}
	return out
}

func declareLocal(ctx *context, in *ast.DeclareLocal) *semantic.DeclareLocal {
	out := &semantic.DeclareLocal{AST: in}
	out.Local = &semantic.Local{Declaration: out, Name: in.Name.Value}
	out.Local.Value = expression(ctx, in.RHS)
	out.Local.Type = out.Local.Value.ExpressionType()
	if equal(out.Local.Type, semantic.VoidType) {
		ctx.errorf(in, "void in local declaration")
	}
	ctx.add(out.Local.Name, out.Local)
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
	return out
}

func switch_(ctx *context, in *ast.Switch) *semantic.Switch {
	out := &semantic.Switch{AST: in}
	out.Value = expression(ctx, in.Value)
	vt := out.Value.ExpressionType()
	for _, c := range in.Cases {
		out.Cases = append(out.Cases, case_(ctx, c, vt))
	}
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
	return out
}

func iteration(ctx *context, in *ast.Iteration) *semantic.Iteration {
	v := &semantic.Local{Name: in.Variable.Value}
	out := &semantic.Iteration{AST: in, Iterator: v}
	out.Iterable = expression(ctx, in.Iterable)
	if b, ok := out.Iterable.(*semantic.BinaryOp); !ok {
		ctx.errorf(in, "iterable can only be range operator, got %T", b)
	} else if b.Operator != ast.OpRange {
		ctx.errorf(in, "iterable can only be range operator, got %s\n", b.Operator)
	}
	v.Type = out.Iterable.ExpressionType()
	ctx.with(semantic.VoidType, func() {
		ctx.add(v.Name, v)
		out.Block = block(ctx, in.Block, out)
	})
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
	return out
}
