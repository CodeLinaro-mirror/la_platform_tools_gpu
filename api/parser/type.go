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

// { annotation } 'class' identifer '{' { field } '}'
func class(p *parse.Parser, cst *parse.Branch, a *ast.Annotations) *ast.Class {
	if !peekKeyword(ast.KeywordClass, p) {
		return nil
	}
	c := &ast.Class{}
	consumeAnnotations(&c.Annotations, a)
	p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
		c.CST = cst
		requireKeyword(ast.KeywordClass, p, cst)
		c.Name = requireIdentifier(p, cst)
		requireOperator(ast.OpBlockStart, p, cst)
		for !operator(ast.OpBlockEnd, p, cst) {
			c.Fields = append(c.Fields, requireField(p, cst, nil))
		}
	})
	return c
}

// { annotation } type identifier [ '=' expression ] [ ',' ]
func requireField(p *parse.Parser, cst *parse.Branch, a *ast.Annotations) *ast.Field {
	f := &ast.Field{}
	consumeAnnotations(&f.Annotations, a)
	p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
		f.CST = cst
		parseAnnotations(&f.Annotations, p, cst)
		f.Type = requireTypeRef(p, cst)
		f.Name = requireIdentifier(p, cst)
		if operator(ast.OpAssign, p, cst) {
			f.Default = requireExpression(p, cst)
		}
		operator(ast.OpListSeparator, p, cst)
	})
	return f
}

// { annotation } 'define' identifier expression
func definition(p *parse.Parser, cst *parse.Branch, a *ast.Annotations) *ast.Definition {
	if !peekKeyword(ast.KeywordDefine, p) {
		return nil
	}
	d := &ast.Definition{}
	consumeAnnotations(&d.Annotations, a)
	p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
		d.CST = cst
		requireKeyword(ast.KeywordDefine, p, cst)
		d.Name = requireIdentifier(p, cst)
		d.Expression = requireExpression(p, cst)
	})
	return d
}

// { annotation } ( 'enum' | 'bitfield' ) [ : identifier { ',' identifer} ] '{' { identifier '=' expression [ ',' ] } '}'
func enum(p *parse.Parser, cst *parse.Branch, a *ast.Annotations) *ast.Enum {
	if !peekKeyword(ast.KeywordEnum, p) && !peekKeyword(ast.KeywordBitfield, p) {
		return nil
	}
	s := &ast.Enum{}
	consumeAnnotations(&s.Annotations, a)
	p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
		s.CST = cst
		if keyword(ast.KeywordEnum, p, cst) == nil {
			requireKeyword(ast.KeywordBitfield, p, cst)
			s.IsBitfield = true
		}
		s.Name = requireIdentifier(p, cst)
		if operator(ast.OpExtends, p, cst) {
			for !peekOperator(ast.OpBlockStart, p) {
				if len(s.Extends) > 0 {
					requireOperator(ast.OpListSeparator, p, cst)
				}
				extend := requireIdentifier(p, cst)
				s.Extends = append(s.Extends, extend)
			}
		}
		requireOperator(ast.OpBlockStart, p, cst)
		for !operator(ast.OpBlockEnd, p, cst) {
			p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
				entry := &ast.EnumEntry{}
				entry.CST = cst
				entry.Name = requireIdentifier(p, cst)
				requireOperator(ast.OpAssign, p, cst)
				entry.Value = requireNumber(p, cst)
				operator(ast.OpListSeparator, p, cst)
				s.Entries = append(s.Entries, entry)
			})
		}
	})
	return s
}

// { annotation } 'alias' type identifier
func alias(p *parse.Parser, cst *parse.Branch, a *ast.Annotations) *ast.Alias {
	if !peekKeyword(ast.KeywordAlias, p) {
		return nil
	}
	s := &ast.Alias{}
	consumeAnnotations(&s.Annotations, a)
	p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
		s.CST = cst
		requireKeyword(ast.KeywordAlias, p, cst)
		s.To = requireTypeRef(p, cst)
		s.Name = requireIdentifier(p, cst)
	})
	return s
}

// { annotation } 'type' type identifier
func pseudonym(p *parse.Parser, cst *parse.Branch, a *ast.Annotations) *ast.Pseudonym {
	if !peekKeyword(ast.KeywordPseudonym, p) {
		return nil
	}
	s := &ast.Pseudonym{}
	consumeAnnotations(&s.Annotations, a)
	p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
		s.CST = cst
		requireKeyword(ast.KeywordPseudonym, p, cst)
		s.To = requireTypeRef(p, cst)
		s.Name = requireIdentifier(p, cst)
	})
	return s
}

// generic { extend_type }
func typeRef(p *parse.Parser, cst *parse.Branch) ast.Node {
	var ref ast.Node
	preconst := keyword(ast.KeywordConst, p, cst) != nil
	if g := generic(p, cst); g == nil {
		return nil
	} else {
		ref = g
	}
	if id, isid := ref.(*ast.Generic); isid && peekOperator(ast.OpMember, p) {
		t := &ast.Imported{From: id.Name}
		p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
			t.CST = cst
			requireOperator(ast.OpMember, p, cst)
			t.Name = requireIdentifier(p, cst)
		})
		ref = t
	}
	for {
		if t := extendTypeRef(p, cst, ref, preconst); t != nil {
			preconst = false
			ref = t
		} else {
			break
		}
	}
	if preconst {
		p.ErrorAt(ref.Fragment(), "const only applies to pointer types")
	}
	return ref
}

// generic ( pointer_type | static_array_type )
func extendTypeRef(p *parse.Parser, cst *parse.Branch, ref ast.Node, preconst bool) ast.Node {
	if e := pointerType(p, cst, ref, preconst); e != nil {
		return e
	}
	if preconst {
		p.ErrorAt(ref.Fragment(), "const only applies to pointer types")
	}
	if s := indexedType(p, cst, ref); s != nil {
		return s
	}
	return nil
}

func requireTypeRef(p *parse.Parser, cst *parse.Branch) ast.Node {
	t := typeRef(p, cst)
	if t == nil {
		p.Expected("type reference")
	}
	return t
}

// lhs_type ['const'] '*'
func pointerType(p *parse.Parser, cst *parse.Branch, ref ast.Node, preconst bool) *ast.PointerType {
	if !peekOperator(ast.OpPointer, p) && !peekKeyword(ast.KeywordConst, p) {
		return nil
	}
	t := &ast.PointerType{To: ref, Const: preconst}
	p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
		t.CST = cst
		if !t.Const {
			t.Const = keyword(ast.KeywordConst, p, cst) != nil
		}
		requireOperator(ast.OpPointer, p, cst)
	})
	return t
}

// lhs_type '[' [ expression ] ']'
func indexedType(p *parse.Parser, cst *parse.Branch, ref ast.Node) *ast.IndexedType {
	if !peekOperator(ast.OpIndexStart, p) {
		return nil
	}
	t := &ast.IndexedType{ValueType: ref}
	p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
		t.CST = cst
		requireOperator(ast.OpIndexStart, p, cst)
		if !peekOperator(ast.OpIndexEnd, p) {
			t.Index = requireExpression(p, cst)
		}
		requireOperator(ast.OpIndexEnd, p, cst)
	})
	return t
}
