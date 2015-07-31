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

package parser

import (
	"unicode"
	"unicode/utf8"

	"android.googlesource.com/platform/tools/gpu/api/ast"
	"android.googlesource.com/platform/tools/gpu/parse"
)

func peekOperator(op string, p *parse.Parser) bool {
	scanned := scanOperator(p)
	p.Rollback()
	return op == scanned
}

func operator(op string, p *parse.Parser, cst *parse.Branch) bool {
	scanned := scanOperator(p)
	if op != scanned {
		p.Rollback()
		return false
	}
	p.ParseLeaf(cst, nil)
	return true
}

func requireOperator(op string, p *parse.Parser, cst *parse.Branch) {
	if !operator(op, p, cst) {
		p.Expected(string(op))
	}
}

func scanOperator(p *parse.Parser) string {
	for _, op := range ast.Operators {
		if p.String(string(op)) {
			r, _ := utf8.DecodeLastRuneInString(string(op))
			if !unicode.IsLetter(r) || !unicode.IsLetter(p.Peek()) {
				return op
			}
		}
	}
	return ""
}

// lhs operator expression
func binaryOp(p *parse.Parser, lhs ast.Node) *ast.BinaryOp {
	op := scanOperator(p)
	if _, found := ast.BinaryOperators[op]; !found {
		p.Rollback()
		return nil
	}
	n := &ast.BinaryOp{LHS: lhs, Operator: op}
	p.Extend(lhs.Node(), func(p *parse.Parser, cst *parse.Branch) {
		n.CST = cst
		p.ParseLeaf(cst, nil)
		n.RHS = requireExpression(p, cst)
	})
	return n
}

// operator expression
func unaryOp(p *parse.Parser, cst *parse.Branch) *ast.UnaryOp {
	op := scanOperator(p)
	if _, found := ast.UnaryOperators[op]; !found {
		p.Rollback()
		return nil
	}
	p.ParseLeaf(cst, nil)
	n := &ast.UnaryOp{Operator: op}
	p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
		n.CST = cst
		n.Expression = requireExpression(p, cst)
	})
	return n
}
