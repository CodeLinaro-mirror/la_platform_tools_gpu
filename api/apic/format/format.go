// Copyright (C) 2015 The Android Open Source Project
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

// Package format registers and implements the "format" apic command.
//
// The format command re-formats an API file to a consistent style.
package format

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"text/tabwriter"

	"android.googlesource.com/platform/tools/gpu/api/apic/commands"
	"android.googlesource.com/platform/tools/gpu/api/ast"
	"android.googlesource.com/platform/tools/gpu/api/parser"
	"android.googlesource.com/platform/tools/gpu/parse"
)

const debug = false

var (
	command = &commands.Command{
		Name:      "format",
		ShortHelp: "Formats an api file",
		Run:       doFormat,
	}
)

func init() {
	commands.Register(command)
}

func doFormat(flags flag.FlagSet) {
	args := flags.Args()
	if len(args) < 1 {
		commands.Usage("Missing api file\n")
	}
	for _, path := range args {
		f, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Printf("Failed to read api file '%s': %v\n", path, err)
			continue
		}

		api, errs := parser.Parse(string(f))
		if len(errs) > 0 {
			fmt.Printf("Errors while parsing '%s':\n", path)
			for i, e := range errs {
				fmt.Printf("%d: %v", i, e)
			}
			continue
		}

		buf := &bytes.Buffer{}
		Format(api, buf)
		if err = ioutil.WriteFile(path, buf.Bytes(), 0777); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to write formatted api file '%s': %v\n", path, err)
		}
	}
}

// Format prints the full re-formatted AST tree to w.
func Format(api *ast.API, w io.Writer) {
	p := printer{}
	// traverse the AST, applying markup for the CST nodes.
	p.markup(api)

	// Writers are chained like so:
	//    indenter -> [tabwriter -> tabwriter -> ...] -> wsTrimmer -> w
	// initially there are no tabwriters, so the initial chain is:
	//    indenter -> wsTrimmer -> w
	trimmer := &wsTrimmer{out: w}
	p.indenter.out = trimmer
	p.out = trimmer

	// print the CST using the markup generated from the AST.
	p.print(api.Node())
}

type printer struct {
	tabbers    []*tabwriter.Writer
	indenter   indenter
	out        io.Writer
	injections map[injectKey]string
	aligns     map[parse.Node]struct{}
}

// isNewline returns true if n starts on a new line.
func isNewline(n parse.Node) bool {
	for _, s := range n.Prefix() {
		if strings.Contains(s.Token().String(), "\n") {
			return true
		}
	}
	if b, ok := n.(*parse.Branch); ok && len(b.Children) > 0 {
		return isNewline(b.Children[0])
	}
	return false
}

// bracketed returns true if the block b was declared with {} brackets.
func bracketed(b *ast.Block) bool {
	if len(b.CST.Children) == 0 {
		return false
	}
	return b.CST.Children[0].Token().String() == ast.OpBlockStart
}

// markup populates the parse.Node maps with information based on the ast tree.
func (p *printer) markup(n ast.Node) {
	switch n := n.(type) {
	case *ast.Alias:
		p.inject(n.To, afterPrefix, "•")
		p.inject(n.Name, afterPrefix, "\t•")

	case *ast.Annotation:
		p.inject(n, beforeSuffix, "•")

	case *ast.API:
		p.align(n)

	case *ast.Assign:
		p.inject(n.LHS, beforeSuffix, "•")
		p.inject(n.RHS, afterPrefix, "•")

	case *ast.BinaryOp:
		if n.Operator != ":" {
			p.inject(n.LHS, beforeSuffix, "•")
			p.inject(n.RHS, afterPrefix, "•")
		}

	case *ast.Block:
		p.align(n)

		// •{}•
		p.inject(n, beforePrefix, "•")
		p.inject(n, afterSuffix, "•")

		if c := len(n.CST.Children); c > 0 {
			if bracketed(n) {
				// {•» statements... «•}
				p.inject(n.CST.Children[0], beforeSuffix, "•»")
				p.inject(n.CST.Children[c-1], afterPrefix, "«•")
			} else {
				// »statement«
				p.inject(n, beforePrefix, "»")
				p.inject(n, afterSuffix, "«")
			}
		}

	case *ast.Branch:
		p.inject(n.Condition, afterPrefix, "•")

	case *ast.Call:
		p.align(n)
		if c := len(n.Arguments); c > 0 {
			for i, v := range n.Arguments {
				if i == 0 {
					p.inject(v, afterPrefix, "\t")
				} else {
					p.inject(v, afterPrefix, "•\t")
				}
			}
			p.inject(n.Target, beforeSuffix, "»")
			p.inject(n.Arguments[c-1], beforeSuffix, "«")
		}

	case *ast.Case:
		p.inject(n, afterPrefix, "•")
		if c := len(n.Conditions); c > 0 {
			for _, c := range n.Conditions {
				p.inject(c, afterPrefix, "•")
			}
			p.inject(n.Conditions[0], afterPrefix, "»»")
			p.inject(n.Conditions[c-1], beforeSuffix, "««")
		}
		if !isNewline(n.Block.CST) {
			// align:
			// case Foo: |•{ ... }•
			// case Blah:|•{ ... }•
			p.inject(n.Block, afterPrefix, "\t")
		}

	case *ast.Class:
		p.align(n)
		p.inject(n.Name, afterPrefix, "•")
		p.inject(n.Name, beforeSuffix, "•")
		if c := len(n.Fields); c > 0 {
			p.inject(n.Name, beforeSuffix, "»")
			p.inject(n.Fields[c-1], beforeSuffix, "«")
		}

	case *ast.Default:
		p.inject(n, afterPrefix, "•")
		if !isNewline(n.Block.CST) {
			// align:
			// case Foo: |•{ ... }•
			// case Blah:|•{ ... }•
			p.inject(n.Block, afterPrefix, "\t")
		}

	case *ast.DeclareLocal:
		p.inject(n.Name, beforeSuffix, "•")
		p.inject(n.RHS, afterPrefix, "•")

	case *ast.Enum:
		p.align(n)
		p.inject(n.Name, afterPrefix, "•")
		p.inject(n.Name, beforeSuffix, "•")
		if c := len(n.Entries); c > 0 {
			p.inject(n.Name, beforeSuffix, "»")
			p.inject(n.Entries[c-1], beforeSuffix, "«")
		}

	case *ast.EnumEntry:
		//name[A•   ]=[B•]value[C    ]••// comment
		p.inject(n.Name, beforeSuffix, "\t•") // A
		p.inject(n.Value, afterPrefix, "\t•") // B
		p.inject(n.Value, beforeSuffix, "\t") // C

	case *ast.Field:
		//type[A•   ]name[B•   ]=[C•]default[D    ]•// comment
		p.inject(n.Name, afterPrefix, "\t•") // A
		if n.Default != nil {
			p.inject(n.Name, beforeSuffix, "\t•")   // B
			p.inject(n.Default, afterPrefix, "•")   // C
			p.inject(n.Default, beforeSuffix, "\t") // D
		} else {
			p.inject(n.Name, beforeSuffix, "\t\t")
		}

	case *ast.Function:
		p.align(n)
		ret := n.Parameters[len(n.Parameters)-1]
		p.inject(n.Name, afterPrefix, "•")
		p.inject(ret, afterPrefix, "•")
		for i, v := range n.Parameters[:len(n.Parameters)-1] {
			if i == 0 {
				p.inject(v.Type, afterPrefix, "\t")
			} else {
				p.inject(v.Type, afterPrefix, "•\t")
			}
			p.inject(v.Name, afterPrefix, "•\t")
		}

	case *ast.Generic:
		if len(n.Arguments) > 0 {
			for _, a := range n.Arguments[1:] {
				p.inject(a, afterPrefix, "•")
			}
		}

	case *ast.Iteration:
		p.inject(n.Variable, afterPrefix, "•")
		p.inject(n.Variable, beforeSuffix, "•")
		p.inject(n.Iterable, afterPrefix, "•")

	case *ast.Import:
		if n.Name != nil {
			p.inject(n.Name, afterPrefix, "•")
		}
		p.inject(n.Path, afterPrefix, "•")

	case *ast.NamedArg:
		p.inject(n.Value, afterPrefix, "•\t")

	case *ast.PointerType:
		if n.Const {
			p.inject(n.To, beforeSuffix, "•")
		}

	case *ast.PreConst:
		p.inject(n.Type, afterPrefix, "•")

	case *ast.Pseudonym:
		p.inject(n.To, afterPrefix, "•")
		p.inject(n.Name, afterPrefix, "\t•")

	case *ast.Return:
		p.inject(n.Value, afterPrefix, "•")

	case *ast.Switch:
		p.align(n)
		p.inject(n.Value, afterPrefix, "•")
		p.inject(n.Value, beforeSuffix, "•")
		p.inject(n.Value, afterSuffix, "»")
		p.inject(n.CST.Children[len(n.CST.Children)-1], beforePrefix, "«")
	}

	ast.Visit(n, p.markup)
}

// print traverses and prints the CST, applying modifications based on the
// markup pass.
func (p *printer) print(n parse.Node) {
	// emit any beforePrefix injections.
	if s, ok := p.injections[injectKey{n, beforePrefix}]; ok {
		p.write(s)
	}

	// print the prefix comments.
	p.separator(n.Prefix())

	// emit any afterPrefix injections.
	if s, ok := p.injections[injectKey{n, afterPrefix}]; ok {
		p.write(s)
	}

	switch n := n.(type) {
	case *parse.Branch:
		// if this node should align the children, push a new tabber.
		_, align := p.aligns[n]
		if align {
			p.pushTabber()
		}
		// print the child CST nodes.
		for _, c := range n.Children {
			p.print(c)
		}
		if align {
			p.popTabber()
		}
	case *parse.Leaf:
		p.write(n.Token().String())

	default:
		panic("Unknown parse node type")
	}

	// emit any beforeSuffix injections.
	if s, ok := p.injections[injectKey{n, beforeSuffix}]; ok {
		p.write(s)
	}

	// print the suffix comments.
	p.separator(n.Suffix())

	// emit any afterSuffix injections.
	if s, ok := p.injections[injectKey{n, afterSuffix}]; ok {
		p.write(s)
	}
}

// write prints the string s to the indenter which is always the head of the
// writer chain.
func (p *printer) write(s string) {
	p.indenter.Write([]byte(s))
}

// pushTabber injects a new tabwriter after the indenter in the writer chain.
func (p *printer) pushTabber() {
	var t *tabwriter.Writer
	if debug {
		t = tabwriter.NewWriter(p.indenter.out, 0, 2, 0, ' ', tabwriter.Debug)
	} else {
		t = tabwriter.NewWriter(p.indenter.out, 0, 2, 0, ' ', 0)
	}
	p.tabbers = append(p.tabbers, t)
	p.indenter.out = t
}

// popTabber removes the tabwriter after the indenter in the writer chain.
func (p *printer) popTabber() {
	c := len(p.tabbers)
	p.tabbers[c-1].Flush()
	p.tabbers = p.tabbers[:c-1]
	if c := len(p.tabbers); c > 0 {
		p.indenter.out = p.tabbers[c-1]
	} else {
		p.indenter.out = p.out
	}
}

type position int

const (
	beforePrefix = position(iota)
	afterPrefix
	beforeSuffix
	afterSuffix
)

type injectKey struct {
	n parse.Node
	p position
}

// inject adds s using r relative to the AST or CST node n.
func (p *printer) inject(n interface{}, r position, s string) {
	if p.injections == nil {
		p.injections = make(map[injectKey]string)
	}
	var key injectKey
	key.p = r
	switch n := n.(type) {
	case parse.Node:
		key.n = n
	case ast.Node:
		key.n = n.Node()
	default:
		panic(fmt.Errorf("n must be a parse.Node or ast.Node. Got %T", n))
	}
	p.injections[key] = p.injections[key] + s
}

// align marks up n's children to be printed with a new tabwriter.
func (p *printer) align(n ast.Node) {
	if p.aligns == nil {
		p.aligns = make(map[parse.Node]struct{})
	}
	p.aligns[n.Node()] = struct{}{}
}

// separator writes sep to the indenter iff it is a comment.
// All comments are preceeded with a soft whitespace.
func (p *printer) separator(sep parse.Separator) {
	for _, sep := range sep {
		s := sep.Token().String()
		switch {
		case strings.HasPrefix(s, "//"), strings.HasPrefix(s, "/*"):
			p.write("•")
			// Write the '/' to the indenter to get new lines indented.
			p.write(s[:1])
			// Write the rest of the comment skipping the indenter, as we don't want
			// to change new-line indentation within the comments.
			p.indenter.out.Write([]byte(s[1:]))

		case strings.HasPrefix(s, "\n"):
			p.write(s)
		}
	}
}
