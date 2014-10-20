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
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

type listNode []interface{}

type arrayNode listNode

type callNode struct {
	name *valueNode
	args listNode
}

type numberNode uint64
type valueNode string

const (
	tagComma      = ","
	tagBeginArray = "["
	tagEndArray   = "]"
	tagBeginCall  = "("
	tagEndCall    = ")"
)

var (
	comma      = opParser(tagComma)
	beginArray = opParser(tagBeginArray)
	endArray   = opParser(tagEndArray)
	beginCall  = opParser(tagBeginCall)
	endCall    = opParser(tagEndCall)
)

func opParser(op string) LeafParser {
	return func(p *Parser, _ *Leaf) {
		if !p.String(op) {
			p.Expected(op)
		}
	}
}

func maybeValue(p *Parser, in *Branch) interface{} {
	switch {
	case peek(&p.Reader, tagBeginArray):
		a := &arrayNode{}
		p.ParseBranch(in, a.parse)
		return a
	case p.Numeric():
		var n numberNode
		p.ParseLeaf(in, n.consume)
		return &n
	case p.AlphaNumeric():
		var n valueNode
		p.ParseLeaf(in, n.consume)
		if peek(&p.Reader, tagBeginCall) {
			c := &callNode{name: &n}
			p.ParseBranch(in, c.parse)
			return c
		} else {
			return &n
		}
	}
	return nil
}

func parseValue(p *Parser, in *Branch) interface{} {
	v := maybeValue(p, in)
	if v == nil {
		p.Expected("value")
	}
	return v
}

func (n *valueNode) parse(p *Parser, l *Leaf) {
	if !p.AlphaNumeric() {
		p.Expected("value")
	}
	n.consume(p, l)
}

func (n *valueNode) consume(p *Parser, l *Leaf) {
	l.SetToken(p.Consume())
	*n = valueNode(l.Token().String())
}

func (n *numberNode) consume(p *Parser, l *Leaf) {
	l.SetToken(p.Consume())
	v, _ := strconv.ParseUint(l.Token().String(), 0, 32)
	*n = numberNode(v)
}

func (n *listNode) parse(p *Parser, cst *Branch) {
	v := maybeValue(p, cst)
	if v == nil {
		return
	}
	*n = append(*n, v)
	for p.String(tagComma) {
		p.ParseLeaf(cst, nil)
		*n = append(*n, parseValue(p, cst))
	}
}

func (n *arrayNode) parse(p *Parser, cst *Branch) {
	p.ParseLeaf(cst, beginArray)
	p.ParseBranch(cst, (*listNode)(n).parse)
	p.ParseLeaf(cst, endArray)
}

func (n *callNode) parse(p *Parser, cst *Branch) {
	p.ParseLeaf(cst, beginCall)
	p.ParseBranch(cst, n.args.parse)
	p.ParseLeaf(cst, endCall)
}

func compareAST(t *testing.T, expect, got *listNode) {
	compareList(t, "root", expect, got)
}

func compareValue(t *testing.T, in string, expect, got *valueNode) {
	if *expect != *got {
		t.Fatalf("expected %q got %q in %s", *expect, *got, in)
	}
}

func compareNumber(t *testing.T, in string, expect, got *numberNode) {
	if *expect != *got {
		t.Fatalf("expected %d got %d in %s", *expect, *got, in)
	}
}

func compareList(t *testing.T, in string, expect, got *listNode) {
	if len(*expect) != len(*got) {
		t.Fatalf("expected %v entries got %v in %s", len(*expect), len(*got), in)
	}
	for i, e := range *expect {
		in := fmt.Sprintf("%s#%v", in, i)
		g := (*got)[i]
		if reflect.TypeOf(e) != reflect.TypeOf(g) {
			t.Fatalf("expected type %T got %T in %s", e, g, in)
		}
		switch v := e.(type) {
		case *callNode:
			compareCall(t, in, v, g.(*callNode))
		case *arrayNode:
			compareArray(t, in, v, g.(*arrayNode))
		case *listNode:
			compareList(t, in, v, g.(*listNode))
		case *valueNode:
			compareValue(t, in, v, g.(*valueNode))
		case *numberNode:
			compareNumber(t, in, v, g.(*numberNode))
		default:
			t.Fatalf("Unknown types %T", e)
		}
	}
}

func compareArray(t *testing.T, in string, expect, got *arrayNode) {
	compareList(t, fmt.Sprintf("%s#array", in), (*listNode)(expect), (*listNode)(got))
}

func compareCall(t *testing.T, in string, expect, got *callNode) {
	compareValue(t, fmt.Sprintf("%s#call", in), expect.name, got.name)
	compareList(t, fmt.Sprintf("%s#call %s", in, *expect.name), &expect.args, &got.args)
}

func valueOf(v interface{}) interface{} {
	switch t := v.(type) {
	case string:
		n := valueNode(t)
		v = &n
	case int:
		n := numberNode(t)
		v = &n
	}
	return v
}

func list(values ...interface{}) *listNode {
	n := &listNode{}
	for _, v := range values {
		*n = append(*n, valueOf(v))
	}
	return n
}

func array(values ...interface{}) *arrayNode {
	n := &arrayNode{}
	for _, v := range values {
		*n = append(*n, valueOf(v))
	}
	return n
}

func call(name string, values ...interface{}) *callNode {
	n := &callNode{}
	nv := valueNode(name)
	n.name = &nv
	for _, v := range values {
		n.args = append(n.args, valueOf(v))
	}
	return n
}
