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
	"fmt"
	"strings"
	"testing"

	"android.googlesource.com/platform/tools/gpu/api/ast"
	"android.googlesource.com/platform/tools/gpu/check"
	"android.googlesource.com/platform/tools/gpu/parse"
)

// printsCST dumps the CST tree from n.
// Mappings from CST -> AST should be passed as hints.
func printCST(t *testing.T, n parse.Node, hints map[parse.Node][]ast.Node) {
	t.Error(cstExpr(n))
	printCSTTree(t, n, hints, "", true)
}

func cstExpr(n parse.Node) string {
	switch n := n.(type) {
	case *parse.Leaf:
		return fmt.Sprintf(`L("%s")`, n.Token().String())
	case *parse.Branch:
		args := make([]string, len(n.Children))
		for i, c := range n.Children {
			args[i] = cstExpr(c)
		}
		return fmt.Sprintf(`B(%v)`, strings.Join(args, ", "))
	default:
		return ""
	}
}

func printCSTTree(t *testing.T, n parse.Node, hints map[parse.Node][]ast.Node, tab string, last bool) {
	prefix := "├──"
	if last {
		prefix = "└──"
	}
	var hint string
	if astNodes, ok := hints[n]; ok {
		hint = " ("
		for i, ast := range astNodes {
			if i > 0 {
				hint += ", "
			}
			hint += fmt.Sprintf("%T", ast)
		}
		hint += ")"
	}
	str := strings.Replace(n.Token().String(), "\n", "¬", -1)
	t.Errorf(`%s%s"%s"%s`, tab, prefix, str, hint)
	if b, ok := n.(*parse.Branch); ok {
		for i, c := range b.Children {
			if last {
				printCSTTree(t, c, hints, tab+"   ", i == len(b.Children)-1)
			} else {
				printCSTTree(t, c, hints, tab+"│  ", i == len(b.Children)-1)
			}
		}
	}
}

func TestParsedCST(t *testing.T) {
	runes := make([]rune, 0, 0xffff) // needs to be big enough to not be resized
	B := func(n ...parse.Node) *parse.Branch { return &parse.Branch{Children: n} }
	L := func(str string) *parse.Leaf {
		r := []rune(str)
		s, e := len(runes), len(runes)+len(r)
		runes = append(append(runes, r...), '·')
		l := &parse.Leaf{}
		l.SetToken(parse.Token{Runes: runes, Start: s, End: e})
		return l
	}

	for _, test := range []struct {
		name     string
		expected parse.Node
		source   string
		errors   []string
	}{
		{
			name:     "const int*",
			source:   `const int* a`,
			expected: B(B(B(L("const"), B(B(L("int")), L("*"))), L("a"))),
		},
		{
			name:     "const char* const *",
			source:   `const char* const * a`,
			expected: B(B(B(B(L("const"), B(B(L("char")), L("*"))), L("const"), L("*")), L("a"))),
		},
	} {
		api, errs := Parse(test.source)
		m := map[parse.Node][]ast.Node{api.Node(): {api}}
		var traverse func(n ast.Node)
		traverse = func(n ast.Node) {
			m[n.Node()] = append(m[n.Node()], n)
			ast.Visit(n, traverse)
		}
		ast.Visit(api, traverse)
		check.SlicesEqual(t, errs, test.errors)
		if got := api.CST; cstExpr(got) != cstExpr(test.expected) {
			t.Errorf("Testing '%s' - Parsed API was not as expected:", test.name)
			t.Error("Expected:")
			printCST(t, test.expected, m)
			t.Error("Got:")
			printCST(t, got, m)
		}
	}
}
