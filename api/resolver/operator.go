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

func unaryOp(ctx *context, in *ast.UnaryOp) *semantic.UnaryOp {
	out := &semantic.UnaryOp{AST: in}
	out.Operator = in.Operator
	out.Expression = expression(ctx, in.Expression)
	et := out.Expression.ExpressionType()
	switch out.Operator {
	case ast.OpNot:
		out.Type = semantic.BoolType
		if !equal(et, semantic.BoolType) {
			ctx.errorf(in, "operator %s applied to non bool type %s", out.Operator, typename(et))
		}
	default:
		ctx.icef(in, "unhandled unary operator %s", out.Operator)
	}
	return out
}

func binaryOp(ctx *context, in *ast.BinaryOp) *semantic.BinaryOp {
	out := &semantic.BinaryOp{AST: in}
	switch in.LHS.(type) {
	case *ast.Number, *ast.Length:
		// leave these to be inferred after the rhs is known
	default:
		out.LHS = expression(ctx, in.LHS)
	}
	out.Operator = in.Operator
	if out.LHS != nil {
		ctx.with(out.LHS.ExpressionType(), func() {
			out.RHS = expression(ctx, in.RHS)
		})
	} else {
		out.RHS = expression(ctx, in.RHS)
		ctx.with(out.RHS.ExpressionType(), func() {
			out.LHS = expression(ctx, in.LHS)
		})
	}
	lt := out.LHS.ExpressionType()
	rt := out.RHS.ExpressionType()
	switch out.Operator {
	case ast.OpIn:
		switch rt := rt.(type) {
		case *semantic.Map:
			if !comparable(lt, rt.KeyType) {
				ctx.errorf(in, "%s with type %s, but map key type is %s", out.Operator, typename(lt), typename(rt.KeyType))
			}
			out.Type = rt.ValueType
		case *semantic.Enum:
			if !equal(lt, rt) {
				ctx.errorf(in, "enum bittest on %s with %s is not allowed", typename(lt), typename(rt))
			}
			out.Type = semantic.BoolType
		default:
			ctx.errorf(in, "%s only allowed on maps, not %s", out.Operator, typename(rt))
		}
	case ast.OpEQ, ast.OpGT, ast.OpLT, ast.OpGE, ast.OpLE, ast.OpNE:
		if !comparable(lt, rt) {
			ctx.errorf(in, "comparison %s of %s against %s not allowed", out.Operator, typename(lt), typename(rt))
		}
		out.Type = semantic.BoolType
	case ast.OpOr, ast.OpAnd:
		if !equal(lt, semantic.BoolType) {
			ctx.errorf(in, "lhs of %s is %s not boolean", out.Operator, typename(lt))
		}
		if !equal(rt, semantic.BoolType) {
			ctx.errorf(in, "rhs of %s is %s not boolean", out.Operator, typename(rt))
		}
		out.Type = semantic.BoolType
	case ast.OpPlus, ast.OpMinus, ast.OpMultiply, ast.OpDivide:
		if !equal(lt, rt) {
			ctx.errorf(in, "operator %s on %s and %s not allowed", out.Operator, typename(lt), typename(rt))
		}
		out.Type = lt
	case ast.OpRange:
		if !equal(lt, rt) {
			ctx.errorf(in, "operator %s on %s and %s not allowed", out.Operator, typename(lt), typename(rt))
		}
		out.Type = lt
	default:
		ctx.icef(in, "unknown binary operator %s", out.Operator)
	}
	return out
}
