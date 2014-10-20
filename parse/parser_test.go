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

package parse

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestEmpty(t *testing.T) {
	testParse(t, ``, b(), list())
}

func TestErrorInvalid(t *testing.T) {
	testFail(t, `@`)
}

func TestErrorExpected(t *testing.T) {
	testFail(t, `[,]`)
	testFail(t, `[ 0`)
}

func TestUnconsumed(t *testing.T) {
	testCustomFail(t, "a", func(p *Parser, cst *Branch) {})
}

func TestUnconsumedByBranch(t *testing.T) {
	testCustomFail(t, "a", func(p *Parser, cst *Branch) {
		p.ParseBranch(cst, func(p *Parser, cst *Branch) {
			p.NotSpace()
		})
	})
}

func TestUnconsumedOnBranch(t *testing.T) {
	testCustomFail(t, "a", func(p *Parser, cst *Branch) {
		p.NotSpace()
		p.ParseBranch(cst, func(p *Parser, cst *Branch) {})
	})
}

func TestSpace(t *testing.T) {
	testParse(t, ` `, ps(" ", b(), nil), list())
	testParse(t, `  `, ps("  ", b(), nil), list())
}

func TestLineComment(t *testing.T) {
	testParse(t, `//note`, ps(s(c("//note")), b(), nil), list())
	testParse(t, `//note
`, ps(s(c("//note"), "\n"), b(), nil), list())
}

func TestBlockComment(t *testing.T) {
	testParse(t, `/*note*/`, ps(s(c("/*note*/")), b(), nil), list())
}

func TestCommentUnclosed(t *testing.T) {
	testFail(t, `/*a`)
}

func TestComplexComment(t *testing.T) {
	testParse(t, `
///*a*/
/*b//c*
//*/ //d
d //e
`, ps(
		s("\n", c("///*a*/"), "\n", c("/*b//c*\n//*/"), " ", c("//d"), "\n"),
		b("d"),
		s(" ", c("//e"), "\n")),
		list("d"))
}

func TestValue(t *testing.T) {
	testParse(t, `a`, b("a"), list("a"))
	testParse(t, `0xA4`, b("0xA4"), list(164))
	testParse(t, ` a`, ps(" ", b("a"), nil), nil)
}

func TestArray(t *testing.T) {
	testParse(t, `[]`, b(b("[", b(), "]")), list(array()))
	testParse(t, `[a]`, b(b("[", b("a"), "]")), list(array("a")))
	testParse(t, `[a,b]`, b(b("[", b("a", ",", "b"), "]")), list(array("a", "b")))
}

func TestCall1(t *testing.T) {
	testParse(t, `a(b)`, b("a", b("(", b("b"), ")")), list(call("a", "b")))
	testParse(t, `a(b,c)`, b("a", b("(", b("b", ",", "c"), ")")), list(call("a", "b", "c")))
}

func TestComplex(t *testing.T) {
	testParse(t,
		` a  (b   ,c)`,
		ps(" ", b(
			ps(nil, "a", "  "),
			b("(", b(ps(nil, "b", "   "), ",", "c"), ")"),
		), nil),
		list(call("a", "b", "c")))
}

func TestErrorLimit(t *testing.T) {
	errs := Parse(func(p *Parser, cst *Branch) {
		for i := 0; true; i++ {
			p.ErrorAt(cst, "failure")
			if i >= ParseErrorLimit {
				t.Fatalf("Parsing not terminated after %v errors", i)
			}
		}
	}, "", NewSkip("//", "/*", "*/"))
	if len(errs) != ParseErrorLimit {
		t.Fatalf("Expected %v errors, got %v", ParseErrorLimit, len(errs))
	}
}

func TestCursor(t *testing.T) {
	root := &listNode{}
	line, column := 3, 5
	content := ""
	for i := 1; i < line; i++ {
		content += "\n"
	}
	for i := 1; i < column; i++ {
		content += " "
	}
	content += "@  \n  "
	errs := Parse(root.parse, content, NewSkip("//", "/*", "*/"))
	if len(errs) == 0 {
		t.Fatalf("Expected errors")
	}
	l, c := errs[0].At.Token().Cursor()
	if line != l || column != c {
		t.Fatalf("Expected line:column of %v:%v and got %v:%v\n", line, column, l, c)
	}
	str := fmt.Sprintf("%s", errs[0])
	prefix := fmt.Sprintf("%v:%v: Unexpected", line, column)
	if !strings.HasPrefix(str, prefix) {
		t.Fatalf("Expected error message to start with %q got %q\n", prefix, str)
	}
}

func TestCustomPanic(t *testing.T) {
	custom := errors.New("custom")
	defer func() {
		r := recover()
		if r != custom {
			t.Fatalf("Expected custom panic recovery")
		}
	}()
	Parse(func(p *Parser, _ *Branch) { panic(custom) }, "", NewSkip("//", "/*", "*/"))
}

func testParse(t *testing.T, content string, cst Node, ast *listNode) {
	root := &listNode{}
	var gotCst *Branch
	rootParse := func(p *Parser, cst *Branch) {
		gotCst = cst
		root.parse(p, cst)
	}
	errs := Parse(rootParse, content, NewSkip("//", "/*", "*/"))
	if len(errs) > 0 {
		for _, e := range errs {
			line, column := e.At.Token().Cursor()
			t.Logf("%v:%v: %s\n", line, column, e.Message)
		}
		t.Fatalf("expected success, got %d errors", len(errs))
		return
	}
	out := &bytes.Buffer{}
	gotCst.WriteTo(out)
	written := out.String()
	if written != content {
		t.Errorf("Ast does not match input")
		t.Errorf("expected '%s' got '%s'", content, written)
		t.Fail()
	}
	verifyTokens(t, gotCst)
	if cst != nil {
		compareCST(t, cst, gotCst)
	}
	if ast != nil {
		compareAST(t, ast, root)
	}
}

func testCustomFail(t *testing.T, content string, do BranchParser) {
	errs := Parse(do, content, NewSkip("//", "/*", "*/"))
	if len(errs) == 0 {
		t.Fatalf("Expected errors")
	} else {
		for _, e := range errs {
			line, column := e.At.Token().Cursor()
			t.Logf("%v:%v: %s\n", line, column, e.Message)
		}
	}
}

func testFail(t *testing.T, content string) {
	root := &listNode{}
	testCustomFail(t, content, root.parse)
}
