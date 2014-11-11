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
	"fmt"
	"reflect"
	"testing"
	"unicode/utf8"
)

func peek(r *Reader, value string) bool {
	result := r.String(value)
	r.Rollback()
	return result
}

func verifyTokens(t *testing.T, got Node) {
	next := 0
	verifyNodeTokens(t, got, &next)
}

func verifyNodeTokens(t *testing.T, n Node, next *int) {
	// walk the prefix
	for _, f := range n.Prefix() {
		verifyFragmentTokens(t, f, next)
	}
	start := *next
	if b, ok := n.(*Branch); ok {
		for _, c := range b.Children {
			verifyNodeTokens(t, c, next)
		}
	} else {
		verifyFragmentTokens(t, n, next)
	}
	end := *next
	// walk the suffix
	for _, f := range n.Suffix() {
		verifyFragmentTokens(t, f, next)
	}
	tok := n.Token()
	if start != end {
		if tok.Start != start || tok.End != end {
			t.Fatalf("branch token range %v:%v expected, got %v:%v",
				start, end, tok.Start, tok.End)
		}
	}
}

func verifyFragmentTokens(t *testing.T, f Fragment, next *int) {
	tok := f.Token()
	str := tok.String()
	length := utf8.RuneCountInString(str)
	if tok.Start != *next {
		t.Fatalf("token start %v expected, got %v for %q", *next, tok.Start, str)
	}
	if tok.Start+length != tok.End {
		t.Fatalf("token range %v:%v but token length %v for %q",
			tok.Start, tok.End, length, str)
	}
	*next = tok.End
}

func compareCST(t *testing.T, expect, got Node) {
	compareNodes(t, "root", expect, got)
}

func compareNodes(t *testing.T, in string, expect, got Node) {
	compareSeparators(t, fmt.Sprintf("%s#prefix", in), expect.Prefix(), got.Prefix())
	if reflect.TypeOf(expect) != reflect.TypeOf(got) {
		t.Fatalf("expected type %T got %T in %s", expect, got, in)
	}
	if b, ok := expect.(*Branch); ok {
		compareBranches(t, in, b, got.(*Branch))
	} else {
		compareLeaves(t, in, expect.(*Leaf), got.(*Leaf))
	}
	compareSeparators(t, fmt.Sprintf("%s#suffix", in), expect.Suffix(), got.Suffix())
}

func compareBranches(t *testing.T, in string, expect, got *Branch) {
	if len(expect.Children) != len(got.Children) {
		t.Fatalf("expected %v children got %v in %s", len(expect.Children), len(got.Children), in)
	}
	for i, e := range expect.Children {
		compareNodes(t, fmt.Sprintf("%s.%v", in, i), e, got.Children[i])
	}
}

func compareLeaves(t *testing.T, in string, expect, got *Leaf) {
	e := expect.Token().String()
	g := got.Token().String()
	if e != g {
		t.Fatalf("expected value %q got %q in %s", e, g, in)
	}
}

func compareSeparators(t *testing.T, in string, expect, got Separator) {
	if len(expect) != len(got) {
		t.Fatalf("expected %v fragments got %v in %s", len(expect), len(got), in)
	}
	for i, e := range expect {
		g := got[i]
		if reflect.TypeOf(e) != reflect.TypeOf(g) {
			t.Fatalf("expected type %T got %T in %s", e, g, in)
		}
		et := e.Token().String()
		gt := g.Token().String()
		if et != gt {
			t.Fatalf("expected fragment %q got %q in %s", et, gt, in)
		}
	}
}

func newStringToken(value string) Token {
	runes := bytes.Runes([]byte(value))
	return Token{Runes: runes, Start: 0, End: len(runes)}
}

func nodeOf(v interface{}) Node {
	if n, ok := v.(Node); ok {
		return n
	}
	l := &Leaf{}
	l.SetToken(newStringToken(v.(string)))
	return l
}

func b(nodes ...interface{}) *Branch {
	n := &Branch{}
	for _, v := range nodes {
		c := nodeOf(v)
		n.Children = append(n.Children, c)
	}
	return n
}

func space(tok string) Fragment {
	return NewFragment(newStringToken(tok))
}

func c(tok string) Fragment {
	return NewFragment(newStringToken(tok))
}

func s(list ...interface{}) Separator {
	var sep Separator
	for _, e := range list {
		if f, ok := e.(Fragment); ok {
			sep = append(sep, f)
		} else {
			sep = append(sep, space(e.(string)))
		}
	}
	return sep
}

func ps(prefix interface{}, v interface{}, suffix interface{}) Node {
	n := nodeOf(v)
	if prefix != nil {
		if sep, ok := prefix.(Separator); ok {
			n.AddPrefix(sep)
		} else {
			n.AddPrefix(s(prefix))
		}
	}
	if suffix != nil {
		if sep, ok := suffix.(Separator); ok {
			n.AddSuffix(sep)
		} else {
			n.AddSuffix(s(suffix))
		}
	}
	return n
}
