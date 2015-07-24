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
	tabbers  []*tabwriter.Writer
	indenter indenter
	out      io.Writer
	prefixes map[parse.Node]string
	suffixes map[parse.Node]string
	aligns   map[parse.Node]struct{}
}

// markup populates the parse.Node maps with information based on the ast tree.
func (p *printer) markup(n ast.Node) {
	switch n := n.(type) {
	case *ast.Alias:
		p.prefix(n.To, "•")
		p.prefix(n.Name, "\t•")

	case *ast.Annotation:
		p.suffix(n, "•")

	case *ast.API:
		p.align(n)

	case *ast.Assign:
		p.suffix(n.LHS, "•")
		p.prefix(n.RHS, "•")

	case *ast.BinaryOp:
		if n.Operator != ":" {
			p.suffix(n.LHS, "•")
			p.prefix(n.RHS, "•")
		}

	case *ast.Block:
		p.align(n)
		p.prefix(n, "•")
		if c := len(n.Statements); c > 0 {
			p.prefix(n, "»")
			p.suffix(n.Statements[c-1], "«")
		}

	case *ast.Branch:
		p.prefix(n.Condition, "•")
		if n.False != nil {
			p.suffix(n.True, "•")
		}

	case *ast.Call:
		p.align(n)
		if c := len(n.Arguments); c > 0 {
			for i, v := range n.Arguments {
				if i == 0 {
					p.prefix(v, "\t")
				} else {
					p.prefix(v, "•\t")
				}
			}
			p.suffix(n.Target, "»")
			p.suffix(n.Arguments[c-1], "«")
		}

	case *ast.Case:
		p.prefix(n, "•")
		if c := len(n.Conditions); c > 0 {
			for _, c := range n.Conditions {
				p.prefix(c, "•")
			}
			p.prefix(n.Conditions[0], "»»")
			p.suffix(n.Conditions[c-1], "««")
		}
		p.prefix(n.Block, "\t")
		p.suffix(n.Block, "•")

	case *ast.Class:
		p.align(n)
		p.prefix(n.Name, "•")
		p.suffix(n.Name, "•")
		if c := len(n.Fields); c > 0 {
			p.suffix(n.Name, "»")
			p.suffix(n.Fields[c-1], "«")
		}

	case *ast.DeclareLocal:
		p.suffix(n.Name, "•")
		p.prefix(n.RHS, "•")

	case *ast.Enum:
		p.align(n)
		p.prefix(n.Name, "•")
		p.suffix(n.Name, "•")
		if c := len(n.Entries); c > 0 {
			p.suffix(n.Name, "»")
			p.suffix(n.Entries[c-1], "«")
		}

	case *ast.EnumEntry:
		//name[A•   ]=[B•]value[C    ]••// comment
		p.suffix(n.Name, "\t•")  // A
		p.prefix(n.Value, "\t•") // B
		p.suffix(n.Value, "\t")  // C

	case *ast.Field:
		//type[A•   ]name[B•   ]=[C•]default[D    ]••// comment
		p.prefix(n.Name, "\t•") // A
		if n.Default != nil {
			p.suffix(n.Name, "\t•")   // B
			p.prefix(n.Default, "•")  // C
			p.suffix(n.Default, "\t") // D
		} else {
			p.suffix(n.Name, "\t\t")
		}

	case *ast.Function:
		p.align(n)
		ret := n.Parameters[len(n.Parameters)-1]
		p.prefix(n.Name, "•")
		p.prefix(ret, "•")
		for i, v := range n.Parameters[:len(n.Parameters)-1] {
			if i == 0 {
				p.prefix(v.Type, "\t")
			} else {
				p.prefix(v.Type, "•\t")
			}
			p.prefix(v.Name, "•\t")
		}

	case *ast.Generic:
		if len(n.Arguments) > 0 {
			for _, a := range n.Arguments[1:] {
				p.prefix(a, "•")
			}
		}

	case *ast.Iteration:
		p.prefix(n.Variable, "•")
		p.suffix(n.Variable, "•")
		p.prefix(n.Iterable, "•")

	case *ast.Import:
		if n.Name != nil {
			p.prefix(n.Name, "•")
		}
		p.prefix(n.Path, "•")

	case *ast.NamedArg:
		p.prefix(n.Value, "•\t")

	case *ast.PointerType:
		if n.Const {
			p.prefix(n.To, "•")
		}

	case *ast.Pseudonym:
		p.prefix(n.To, "•")
		p.prefix(n.Name, "\t•")

	case *ast.Return:
		p.prefix(n.Value, "•")

	case *ast.Switch:
		p.align(n)
		p.prefix(n.Value, "•")
		p.suffix(n.Value, "•")
		if c := len(n.Cases); c > 0 {
			p.suffix(n.Value, "»")
			p.suffix(n.Cases[c-1], "«")
		}
	}

	ast.Visit(n, p.markup)
}

// print traverses and prints the CST, applying modifications based on the
// markup pass.
func (p *printer) print(n parse.Node) {
	// emit any custom prefixes.
	if prefix, ok := p.prefixes[n]; ok {
		p.write(prefix)
	}

	// print the prefix comments.
	p.separator(n.Prefix())

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

	// print the suffix comments.
	p.separator(n.Suffix())

	// print any custom suffixes.
	if suffix, ok := p.suffixes[n]; ok {
		p.write(suffix)
	}
}

// write prints the string s to the indenter which is always the head of the
// writer chain.
func (p *printer) write(s string) {
	p.indenter.Write([]byte(s))
}

// pushTabber injects a new tabwriter after the indenter in the writer chain.
func (p *printer) pushTabber() {
	var out io.Writer
	if c := len(p.tabbers); c > 0 {
		out = p.tabbers[c-1]
	} else {
		out = p.out
	}

	var t *tabwriter.Writer
	if debug {
		t = tabwriter.NewWriter(out, 0, 2, 0, ' ', tabwriter.Debug)
	} else {
		t = tabwriter.NewWriter(out, 0, 2, 0, ' ', 0)
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

// prefix prefixes the node n with s when it is printed.
func (p *printer) prefix(n ast.Node, s string) {
	if p.prefixes == nil {
		p.prefixes = make(map[parse.Node]string)
	}
	p.prefixes[n.Node()] = p.prefixes[n.Node()] + s
}

// prefix suffixes the node n with s when it is printed.
func (p *printer) suffix(n ast.Node, s string) {
	if p.suffixes == nil {
		p.suffixes = make(map[parse.Node]string)
	}
	p.suffixes[n.Node()] = p.suffixes[n.Node()] + s
}

// align marks up n's children to be printed with a new tabwriter.
func (p *printer) align(n ast.Node) {
	if p.aligns == nil {
		p.aligns = make(map[parse.Node]struct{})
	}
	p.aligns[n.Node()] = struct{}{}
}

// separator writes sep to the indenter iff it is a comment.
// All comments are preceeded with two soft whitespaces.
func (p *printer) separator(sep parse.Separator) {
	for _, sep := range sep {
		s := sep.Token().String()
		switch {
		case strings.HasPrefix(s, "//"), strings.HasPrefix(s, "/*"):
			p.write("••")
			p.write(s)

		case strings.HasPrefix(s, "\n"):
			p.write(s)
		}
	}
}
