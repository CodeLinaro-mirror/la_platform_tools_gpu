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

// { annotation } 'class' identifer [ : identifier { ',' identifer} ] '{' { field } '}'
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
		if operator(ast.OpExtends, p, cst) {
			for !peekOperator(ast.OpBlockStart, p) {
				if len(c.Extends) > 0 {
					requireOperator(ast.OpListSeparator, p, cst)
				}
				extend := requireIdentifier(p, cst)
				c.Extends = append(c.Extends, extend)
			}
		}
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

// lhs_type { extend_type }
func typeRef(p *parse.Parser, cst *parse.Branch) ast.Node {
	ref := typeRefLHS(p, cst)
	if ref == nil {
		return nil
	}
	if id, isid := ref.(*ast.Identifier); isid && peekOperator(ast.OpMember, p) {
		t := &ast.Imported{From: id}
		p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
			t.CST = cst
			requireOperator(ast.OpMember, p, cst)
			t.Name = requireIdentifier(p, cst)
		})
		ref = t
	}
	for {
		if t := extendTypeRef(p, cst, ref); t != nil {
			ref = t
		} else {
			break
		}
	}
	return ref
}

// generic_type | identifier
func typeRefLHS(p *parse.Parser, cst *parse.Branch) ast.Node {
	if g := genericType(p, cst); g != nil {
		return g
	}
	if i := identifier(p, cst); i != nil {
		return i
	}
	return nil
}

// lhs_type ( pointer_type | static_array_type )
func extendTypeRef(p *parse.Parser, cst *parse.Branch, ref ast.Node) ast.Node {
	if e := pointerType(p, cst, ref); e != nil {
		return e
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

// ( 'array' | 'map' | 'buffer' | 'ptr' ) '<' type { ',' type } '>'
func genericType(p *parse.Parser, cst *parse.Branch) *ast.GenericType {
	found := false
	for _, word := range []string{ast.KeywordArray, ast.KeywordMap, ast.KeywordBuffer, ast.KeywordPointer} {
		if peekKeyword(word, p) {
			found = true
			break
		}
	}
	if !found {
		return nil
	}
	t := &ast.GenericType{}
	p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
		t.CST = cst
		t.Generic = requireIdentifier(p, cst)
		requireOperator(ast.OpMetaStart, p, cst)
		for {
			t.Args = append(t.Args, requireTypeRef(p, cst))
			if !operator(ast.OpListSeparator, p, cst) {
				break
			}
		}
		requireOperator(ast.OpMetaEnd, p, cst)
	})
	return t
}

// lhs_type '*'
func pointerType(p *parse.Parser, cst *parse.Branch, ref ast.Node) *ast.PointerType {
	if !peekOperator(ast.OpPointer, p) {
		return nil
	}
	t := &ast.PointerType{}
	p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
		t.CST = cst
		requireOperator(ast.OpPointer, p, cst)
	})
	t.To = ref
	return t
}

// lhs_type '[' [ expression ] ']'
func indexedType(p *parse.Parser, cst *parse.Branch, ref ast.Node) *ast.IndexedType {
	if !peekOperator(ast.OpIndexStart, p) {
		return nil
	}
	t := &ast.IndexedType{}
	p.ParseBranch(cst, func(p *parse.Parser, cst *parse.Branch) {
		t.CST = cst
		requireOperator(ast.OpIndexStart, p, cst)
		if !peekOperator(ast.OpIndexEnd, p) {
			t.Index = requireExpression(p, cst)
		}
		requireOperator(ast.OpIndexEnd, p, cst)
	})
	t.ValueType = ref
	return t
}
