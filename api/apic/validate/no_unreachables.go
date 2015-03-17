// Copyright (C) 2015 The Android Open Source Project
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

package validate

import (
	"android.googlesource.com/platform/tools/gpu/api/apic/validate/limits"
	"android.googlesource.com/platform/tools/gpu/api/semantic"
	"android.googlesource.com/platform/tools/gpu/parse"
)

// validateNoUnreachables checks there are no unreachable blocks.
func validateNoUnreachables(api *semantic.API) []parse.Error {
	ctx := context{
		locals:     make(map[*semantic.Local]limits.Limits),
		parameters: make(map[*semantic.Parameter]limits.Limits),
		errors:     &errors{},
	}
	walk(api, ctx.traverse)
	return *ctx.errors
}

func (ctx *context) traverse(n semantic.Node) bool {
	switch n := n.(type) {
	case *semantic.Function:
		ctx := ctx.clone()
		for _, p := range n.FullParameters {
			ctx.parameters[p] = limits.Unbound(p.ExpressionType())
		}
		if n.Block != nil {
			walk(n.Block, ctx.traverse)
		}
		return false

	case *semantic.Assert:
		*ctx = *ctx.setTrue(n.Condition)
		return false

	case *semantic.DeclareLocal:
		ctx.locals[n.Local] = ctx.limits(n.Local.Value)
		return false

	case *semantic.Branch:
		limit := ctx.limits(n.Condition)
		if n.True != nil && len(n.True.Statements) > 0 {
			if limit == limits.False {
				ctx.errors.add(n.True.AST.CST, "Unreachable block")
			} else {
				walk(n.True, ctx.clone().setTrue(n.Condition).traverse)
			}
		}

		if n.False != nil && len(n.False.Statements) > 0 {
			if limit == limits.True {
				ctx.errors.add(n.False.AST.CST, "Unreachable block")
			} else {
				walk(n.False, ctx.clone().setFalse(n.Condition).traverse)
			}
		}
		return false

	default:
		return true
	}
}

type errors []parse.Error

func (l *errors) add(at parse.Fragment, msg string) {
	(*l) = append(*l, parse.Error{At: at, Message: msg})
}

type context struct {
	locals     map[*semantic.Local]limits.Limits
	parameters map[*semantic.Parameter]limits.Limits
	errors     *errors
}

func (ctx *context) clone() *context {
	c := context{
		locals:     make(map[*semantic.Local]limits.Limits),
		parameters: make(map[*semantic.Parameter]limits.Limits),
		errors:     ctx.errors,
	}
	for local, limit := range ctx.locals {
		c.locals[local] = limit
	}
	for parameter, limit := range ctx.parameters {
		c.parameters[parameter] = limit
	}
	return &c
}

func (ctx *context) limits(n semantic.Expression) limits.Limits {
	switch n := n.(type) {
	case *semantic.Local:
		return ctx.locals[n]

	case *semantic.Parameter:
		return ctx.parameters[n]

	case semantic.BoolValue:
		if n {
			return limits.True
		} else {
			return limits.False
		}

	case semantic.Uint16Value:
		return limits.Uint(uint64(n))

	case semantic.Uint32Value:
		return limits.Uint(uint64(n))

	case *semantic.BinaryOp:
		lhs, rhs := ctx.limits(n.LHS), ctx.limits(n.RHS)
		if lhs != nil && rhs != nil {
			return lhs.Binary(n.Operator, rhs)
		}

	case *semantic.UnaryOp:
		expr := ctx.limits(n.Expression)
		if expr != nil {
			return expr.Unary(n.Operator)
		}
	}

	return nil
}

func (ctx *context) setTrue(n semantic.Expression) *context {
	switch n := n.(type) {
	case *semantic.Local:
		ctx.locals[n] = limits.True

	case *semantic.Parameter:
		ctx.parameters[n] = limits.True

	case *semantic.BinaryOp:
		switch n.Operator {
		case "&&":
			ctx.setTrue(n.LHS)
			ctx.setTrue(n.RHS)
		}

	case *semantic.UnaryOp:
		switch n.Operator {
		case "!":
			ctx.setFalse(n.Expression)
		}

	case semantic.BoolValue:
		if n != true {
			panic("Asserting true value is false")
		}
	}
	return ctx
}

func (ctx *context) setFalse(n semantic.Expression) *context {
	switch n := n.(type) {
	case *semantic.Local:
		ctx.locals[n] = limits.False

	case *semantic.Parameter:
		ctx.parameters[n] = limits.False

	case *semantic.BinaryOp:
		switch n.Operator {
		case "||":
			ctx.setFalse(n.LHS)
			ctx.setFalse(n.RHS)
		}

	case *semantic.UnaryOp:
		switch n.Operator {
		case "!":
			ctx.setTrue(n.Expression)
		}

	case semantic.BoolValue:
		if n != false {
			panic("Asserting false value is true")
		}
	}
	return ctx
}
