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
	ctx.mappings[in] = out
	return out
}

func binaryOp(ctx *context, in *ast.BinaryOp) semantic.Expression {
	var lhs semantic.Expression
	var rhs semantic.Expression
	switch in.LHS.(type) {
	case *ast.Number, *ast.Length:
		// leave these to be inferred after the rhs is known
	default:
		lhs = expression(ctx, in.LHS)
	}
	if lhs != nil {
		ctx.with(lhs.ExpressionType(), func() {
			rhs = expression(ctx, in.RHS)
		})
	} else {
		rhs = expression(ctx, in.RHS)
		ctx.with(rhs.ExpressionType(), func() {
			lhs = expression(ctx, in.LHS)
		})
	}
	lt := lhs.ExpressionType()
	rt := rhs.ExpressionType()
	var out semantic.Expression
	switch in.Operator {
	case ast.OpIn:
		switch rt := rt.(type) {
		case *semantic.Map:
			if !comparable(lt, rt.KeyType) {
				ctx.errorf(in, "%s with type %s, but map key type is %s", in.Operator, typename(lt), typename(rt.KeyType))
			}
			out = &semantic.MapContains{AST: in, Map: rhs, Key: lhs}
		case *semantic.Enum:
			if !equal(lt, rt) {
				ctx.errorf(in, "enum bittest on %s with %s is not allowed", typename(lt), typename(rt))
			}
			out = &semantic.BitTest{AST: in, Bitfield: rhs, Bits: lhs}
		default:
			ctx.errorf(in, "%s only allowed on maps, not %s", in.Operator, typename(rt))
			return invalid{}
		}
	case ast.OpEQ, ast.OpGT, ast.OpLT, ast.OpGE, ast.OpLE, ast.OpNE:
		if !comparable(lt, rt) {
			ctx.errorf(in, "comparison %s %s %s not allowed", typename(lt), in.Operator, typename(rt))
		}
		out = &semantic.BinaryOp{AST: in, LHS: lhs, RHS: rhs, Type: semantic.BoolType, Operator: in.Operator}
	case ast.OpOr, ast.OpAnd:
		if !equal(lt, semantic.BoolType) {
			ctx.errorf(in, "lhs of %s is %s not boolean", in.Operator, typename(lt))
		}
		if !equal(rt, semantic.BoolType) {
			ctx.errorf(in, "rhs of %s is %s not boolean", in.Operator, typename(rt))
		}
		out = &semantic.BinaryOp{AST: in, LHS: lhs, RHS: rhs, Type: semantic.BoolType, Operator: in.Operator}
	case ast.OpPlus, ast.OpMinus, ast.OpMultiply, ast.OpDivide:
		if !equal(lt, rt) {
			ctx.errorf(in, "incompatible types for maths %s %s %s", typename(lt), in.Operator, typename(rt))
		}
		out = &semantic.BinaryOp{AST: in, LHS: lhs, RHS: rhs, Type: lt, Operator: in.Operator}
	case ast.OpRange:
		if !equal(lt, rt) {
			ctx.errorf(in, "range %s %s %s not allowed", typename(lt), in.Operator, typename(rt))
		}
		out = &semantic.BinaryOp{AST: in, LHS: lhs, RHS: rhs, Type: lt, Operator: in.Operator}
	default:
		ctx.icef(in, "unknown binary operator %s", in.Operator)
		return invalid{}
	}
	ctx.mappings[in] = out
	return out
}
