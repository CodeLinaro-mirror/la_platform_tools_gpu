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

// lhs { extend }
func requireExpression(p *parse.Parser, cst *parse.Branch) ast.Node {
	lhs := requireLHSExpression(p, cst)
	for {
		if e := extendExpression(p, cst, lhs); e != nil {
			lhs = e
		} else {
			break
		}
	}
	return lhs
}

// ( group | switch | literal | unary_op | generic)
func requireLHSExpression(p *parse.Parser, cst *parse.Branch) ast.Node {
	if g := group(p, cst); g != nil {
		return g
	}
	if s := switch_(p, cst); s != nil {
		return s
	}
	if l := literal(p, cst); l != nil {
		return l
	}
	if u := unaryOp(p, cst); u != nil {
		return u
	}
	if g := generic(p, cst); g != nil {
		return g
	}
	p.Expected("expression")
	return &ast.Invalid{}
}

// lhs (index | call | binary_op | member)
func extendExpression(p *parse.Parser, cst *parse.Branch, lhs ast.Node) ast.Node {
	if i := index(p, cst, lhs); i != nil {
		return i
	}
	if c := call(p, cst, lhs); c != nil {
		return c
	}
	if e := binaryOp(p, cst, lhs); e != nil {
		return e
	}
	if m := member(p, cst, lhs); m != nil {
		return m
	}
	return nil
}

// 'null' | 'true' | 'false' | '"' string '"' | '?' | number
func literal(p *parse.Parser, cst *parse.Branch) ast.Node {
	if l := keyword(ast.KeywordNull, p, cst); l != nil {
		return &ast.Null{CST: l}
	}
	if l := keyword(ast.KeywordTrue, p, cst); l != nil {
		return &ast.Bool{CST: l, Value: true}
	}
	if l := keyword(ast.KeywordFalse, p, cst); l != nil {
		return &ast.Bool{CST: l, Value: false}
	}
	if s := string_(p, cst); s != nil {
		return s
	}
	if peekOperator(ast.OpUnknown, p) {
		n := &ast.Unknown{}
		p.ParseLeaf(cst, func(p *parse.Parser, l *parse.Leaf) {
			n.CST = l
			requireOperator(ast.OpUnknown, p, cst)
		})
		return n
	}
	if n := number(p, cst); n != nil {
		return n
	}
	return nil
}

// '"' string '"'
func string_(p *parse.Parser, cst *parse.Branch) *ast.String {
	if !p.Rune(ast.Quote) {
		return nil
	}
	n := &ast.String{}
	p.ParseLeaf(cst, func(p *parse.Parser, l *parse.Leaf) {
		n.CST = l
		p.SeekRune(ast.Quote)
		p.Rune(ast.Quote)
		l.SetToken(p.Consume())
		v := l.Token().String()
		n.Value = v[1 : len(v)-1]
	})
	return n
}

func requireString(p *parse.Parser, cst *parse.Branch) *ast.String {
	s := string_(p, cst)
	if s == nil {
		p.Expected("string")
	}
	return s
}

// standard numeric formats
func number(p *parse.Parser, cst *parse.Branch) *ast.Number {
	_ = p.Rune('+') || p.Rune('-') // optional sign
	if p.Numeric() == parse.NotNumeric {
		p.Rollback()
		return nil
	}
	n := &ast.Number{}
	p.ParseLeaf(cst, func(p *parse.Parser, l *parse.Leaf) { n.CST = l })
	n.Value = n.CST.Token().String()
	return n
}

func requireNumber(p *parse.Parser, cst *parse.Branch) *ast.Number {
	n := number(p, cst)
	if n == nil {
		p.Expected("number")
	}
	return n
}

// '(' expression ')'
func group(p *parse.Parser, cst *parse.Branch) *ast.Group {
	if !peekOperator(ast.OpListStart, p) {
		return nil
	}
	e := &ast.Group{}
	p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
		e.CST = cst
		requireOperator(ast.OpListStart, p, cst)
		e.Expression = requireExpression(p, cst)
		requireOperator(ast.OpListEnd, p, cst)
	})
	return e
}

// switch '{' { 'case' { expresion } ':' block } '}'
func switch_(p *parse.Parser, cst *parse.Branch) *ast.Switch {
	if !peekKeyword(ast.KeywordSwitch, p) {
		return nil
	}
	e := &ast.Switch{}
	p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
		e.CST = cst
		requireKeyword(ast.KeywordSwitch, p, cst)
		e.Value = requireExpression(p, cst)
		requireOperator(ast.OpBlockStart, p, cst)
		for !operator(ast.OpBlockEnd, p, cst) {
			p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
				entry := &ast.Case{}
				entry.CST = cst
				requireKeyword(ast.KeywordCase, p, cst)
				for !operator(ast.OpExtends, p, cst) {
					if len(entry.Conditions) > 0 {
						requireOperator(ast.OpListSeparator, p, cst)
					}
					entry.Conditions = append(entry.Conditions, requireExpression(p, cst))
				}
				entry.Block = requireBlock(p, cst)
				e.Cases = append(e.Cases, entry)
			})
		}
	})
	return e
}

// lhs '[' expression [ ':' [ expression ] ] ']'
func index(p *parse.Parser, cst *parse.Branch, lhs ast.Node) *ast.Index {
	if !peekOperator(ast.OpIndexStart, p) {
		return nil
	}
	e := &ast.Index{Object: lhs}
	p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
		e.CST = cst
		requireOperator(ast.OpIndexStart, p, cst)
		e.Index = requireExpression(p, cst)
		if operator(ast.OpSlice, p, cst) {
			n := &ast.BinaryOp{LHS: e.Index, Operator: ast.OpSlice}
			if !peekOperator(ast.OpIndexEnd, p) {
				p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
					n.CST = cst
					n.RHS = requireExpression(p, cst)
				})
			}
			e.Index = n
		}
		requireOperator(ast.OpIndexEnd, p, cst)
	})
	return e
}

// lhs '(' [ expression { ',' expression } ] ')'
func call(p *parse.Parser, cst *parse.Branch, lhs ast.Node) *ast.Call {
	if !peekOperator(ast.OpListStart, p) {
		return nil
	}
	e := &ast.Call{Target: lhs}
	p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
		e.CST = cst
		requireOperator(ast.OpListStart, p, cst)
		for !operator(ast.OpListEnd, p, cst) {
			arg := requireExpression(p, cst)
			if i, ok := arg.(*ast.Generic); ok && operator(ast.OpInitialise, p, cst) {
				n := &ast.NamedArg{Name: i.Name}
				p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
					n.CST = cst
					n.Value = requireExpression(p, cst)
				})
				arg = n
			}
			e.Arguments = append(e.Arguments, arg)
			if operator(ast.OpListEnd, p, cst) {
				break
			}
			requireOperator(ast.OpListSeparator, p, cst)
		}
	})
	return e
}

// lhs '.' identifier
func member(p *parse.Parser, cst *parse.Branch, lhs ast.Node) *ast.Member {
	if !peekOperator(ast.OpMember, p) {
		return nil
	}
	e := &ast.Member{Object: lhs}
	p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
		e.CST = cst
		requireOperator(ast.OpMember, p, cst)
		e.Name = requireIdentifier(p, cst)
	})
	return e
}
