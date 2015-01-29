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
	"android.googlesource.com/platform/tools/gpu/api/ast"
	"android.googlesource.com/platform/tools/gpu/parse"
)

// '{' [ url ] { statements } '}' | { statements }
func requireBlock(p *parse.Parser, cst *parse.Branch) *ast.Block {
	block := &ast.Block{}
	p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
		block.CST = cst
		if operator(ast.OpBlockStart, p, cst) {
			block.Docs = url(p, cst)
			for !operator(ast.OpBlockEnd, p, cst) {
				block.Statements = append(block.Statements, requireStatement(p, cst))
			}
		} else {
			block.Statements = append(block.Statements, requireStatement(p, cst))
		}
	})
	return block
}

// ( assert | branch | iteration | expression ) [ declare_local | assign ]
func requireStatement(p *parse.Parser, cst *parse.Branch) interface{} {
	if g := assert(p, cst); g != nil {
		return g
	}
	if g := branch(p, cst); g != nil {
		return g
	}
	if g := iteration(p, cst); g != nil {
		return g
	}
	e := requireExpression(p, cst)
	if g := declareLocal(p, cst, e); g != nil {
		return g
	}
	if g := assign(p, cst, e); g != nil {
		return g
	}
	return e
}

// 'assert' simple_expresssion
func assert(p *parse.Parser, cst *parse.Branch) *ast.Assert {
	if !peekKeyword(ast.KeywordAssert, p) {
		return nil
	}
	s := &ast.Assert{}
	p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
		s.CST = cst
		requireKeyword(ast.KeywordAssert, p, cst)
		s.Condition = requireSimpleExpression(p, cst)
	})
	return s
}

// 'if' expression block [ 'else' block ]
func branch(p *parse.Parser, cst *parse.Branch) *ast.Branch {
	if !peekKeyword(ast.KeywordIf, p) {
		return nil
	}
	s := &ast.Branch{}
	p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
		s.CST = cst
		requireKeyword(ast.KeywordIf, p, cst)
		s.Condition = requireSimpleExpression(p, cst)
		s.True = requireBlock(p, cst)
		if keyword(ast.KeywordElse, p, cst) != nil {
			s.False = requireBlock(p, cst)
		}
	})
	return s
}

// 'for' identifier 'in' simple_expresion block
func iteration(p *parse.Parser, cst *parse.Branch) *ast.Iteration {
	if !peekKeyword(ast.KeywordFor, p) {
		return nil
	}
	s := &ast.Iteration{}
	p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
		s.CST = cst
		requireKeyword(ast.KeywordFor, p, cst)
		s.Variable = requireIdentifier(p, cst)
		requireKeyword(ast.KeywordIn, p, cst)
		s.Iterable = requireSimpleExpression(p, cst)
		s.Block = requireBlock(p, cst)
	})
	return s
}

// lhs ':=' expression
func declareLocal(p *parse.Parser, cst *parse.Branch, lhs interface{}) *ast.DeclareLocal {
	l, ok := lhs.(*ast.Identifier)
	if !ok || !peekOperator(ast.OpDeclare, p) {
		return nil
	}
	s := &ast.DeclareLocal{Name: l}
	p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
		s.CST = cst
		requireOperator(ast.OpDeclare, p, cst)
		s.RHS = requireExpression(p, cst)
	})
	return s
}

// lhs '=' expression
func assign(p *parse.Parser, cst *parse.Branch, lhs interface{}) *ast.Assign {
	if !peekOperator(ast.OpAssign, p) {
		return nil
	}
	s := &ast.Assign{LHS: lhs}
	p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
		s.CST = cst
		requireOperator(ast.OpAssign, p, cst)
		s.RHS = requireExpression(p, cst)
	})
	return s
}
