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

// '{' { statements } '}'
func requireBlock(p *parse.Parser, cst *parse.Branch) *ast.Block {
	block := &ast.Block{}
	p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
		block.CST = cst
		if operator(ast.OpBlockStart, p, cst) {
			for !operator(ast.OpBlockEnd, p, cst) {
				block.Statements = append(block.Statements, requireStatement(p, cst))
			}
		} else {
			block.Statements = append(block.Statements, requireStatement(p, cst))
		}
	})
	return block
}

// ( branch | iteration | return | expression ) [ declare_local | assign ]
func requireStatement(p *parse.Parser, cst *parse.Branch) ast.Node {
	if g := branch(p, cst); g != nil {
		return g
	}
	if g := iteration(p, cst); g != nil {
		return g
	}
	if g := return_(p, cst); g != nil {
		return g
	}
	if g := fence(p, cst); g != nil {
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

// 'if' expression block [ 'else' block ]
func branch(p *parse.Parser, cst *parse.Branch) *ast.Branch {
	if !peekKeyword(ast.KeywordIf, p) {
		return nil
	}
	s := &ast.Branch{}
	p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
		s.CST = cst
		requireKeyword(ast.KeywordIf, p, cst)
		s.Condition = requireExpression(p, cst)
		s.True = requireBlock(p, cst)
		if keyword(ast.KeywordElse, p, cst) != nil {
			s.False = requireBlock(p, cst)
		}
	})
	return s
}

// 'for' identifier 'in' expresion block
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
		s.Iterable = requireExpression(p, cst)
		s.Block = requireBlock(p, cst)
	})
	return s
}

// lhs ':=' expression
func declareLocal(p *parse.Parser, cst *parse.Branch, lhs ast.Node) *ast.DeclareLocal {
	l, ok := lhs.(*ast.Generic)
	if !ok || len(l.Arguments) > 0 || !peekOperator(ast.OpDeclare, p) {
		return nil
	}
	s := &ast.DeclareLocal{Name: l.Name}
	p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
		s.CST = cst
		requireOperator(ast.OpDeclare, p, cst)
		s.RHS = requireExpression(p, cst)
	})
	return s
}

var assignments = []string{
	ast.OpAssign,
	ast.OpAssignPlus,
	ast.OpAssignMinus,
}

// lhs ( '=' | '+=' | '-=' )  expression
func assign(p *parse.Parser, cst *parse.Branch, lhs ast.Node) *ast.Assign {
	op := ""
	for _, test := range assignments {
		if peekOperator(test, p) {
			op = test
			break
		}
	}
	if op == "" {
		return nil
	}
	s := &ast.Assign{LHS: lhs, Operator: op}
	p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
		s.CST = cst
		requireOperator(op, p, cst)
		s.RHS = requireExpression(p, cst)
	})
	return s
}

// 'return' expresssion
func return_(p *parse.Parser, cst *parse.Branch) *ast.Return {
	if !peekKeyword(ast.KeywordReturn, p) {
		return nil
	}
	s := &ast.Return{}
	p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
		s.CST = cst
		requireKeyword(ast.KeywordReturn, p, cst)
		s.Value = requireExpression(p, cst)
	})
	return s
}

// 'fence' statement
func fence(p *parse.Parser, cst *parse.Branch) *ast.Fence {
	if !peekKeyword(ast.KeywordFence, p) {
		return nil
	}
	f := &ast.Fence{}
	p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
		requireKeyword(ast.KeywordFence, p, cst)
		f.CST = cst
	})
	return f
}
