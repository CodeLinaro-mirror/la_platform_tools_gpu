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

package resolver

import (
	"fmt"
	"reflect"

	"android.googlesource.com/platform/tools/gpu/api/ast"
	"android.googlesource.com/platform/tools/gpu/api/semantic"
	"android.googlesource.com/platform/tools/gpu/parse"
)

type context struct {
	errors   parse.ErrorList
	api      *semantic.API
	types    map[string]semantic.Type
	macros   []*macroStub
	scope    *scope
	nextId   uint64
	mappings ASTToSemantic
}

type scope struct {
	semantic.Symbols
	outer     *scope
	inferType semantic.Type
	block     *[]semantic.Node
}

func (ctx *context) errorf(at interface{}, message string, args ...interface{}) {
	n, ok := at.(ast.Node)
	var f parse.Fragment
	if !ok {
		v := reflect.ValueOf(at)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		if v.Kind() == reflect.Struct {
			a := v.FieldByName("AST")
			if a.IsValid() {
				n, ok = a.Interface().(ast.Node)
			}
		}
	}
	if !ok {
		ctx.errors.Add(nil, nil, "Error non node %T", at)
	} else if n == nil || reflect.ValueOf(n).IsNil() {
		ctx.errors.Add(nil, nil, "Error nil node %T", at)
	} else {
		f = n.Node()
		if f == nil {
			ctx.errors.Add(nil, nil, "Error at %T node with nil fragment", at)
		}
	}
	ctx.errors.Add(nil, f, message, args...)
}

func (ctx *context) icef(at ast.Node, message string, args ...interface{}) {
	ctx.errorf(at, "INTERNAL ERROR: "+message, args...)
}

// with evaluates the action within a new nested scope.
// The scope is available as ctx.scope, and the original scope is
// restored before this function returns.
// Type t is used for type inference within the new scope (for
// instance when resolving untyped numeric constants)
func (ctx *context) with(t semantic.Type, action func()) {
	original := ctx.scope
	ctx.scope = &scope{
		outer:     ctx.scope,
		block:     ctx.scope.block,
		inferType: t,
	}
	defer func() { ctx.scope = original }()
	action()
}

// addNamed binds a named node within the current and nested scopes.
func (ctx *context) addNamed(value semantic.NamedNode) {
	ctx.scope.AddNamed(value)
}

// alias maps a name to a node without the node knowing the name.
func (ctx *context) add(name string, value semantic.Node) {
	ctx.scope.Add(name, value)
}

// addMembers adds all the members of owner directly to the curent scope.
func (ctx *context) addMembers(owner semantic.Owner) {
	owner.VisitMembers(func(m semantic.Owned) {
		ctx.scope.AddNamed(m)
	})
}

// addSymbols adds all the entries of symbols directly to the curent scope.
func (ctx *context) addSymbols(symbols *semantic.Symbols) {
	symbols.Visit(func(name string, node semantic.Node) {
		ctx.scope.Add(name, node)
	})
}

// find searches the scope stack for a bindings that matches the name.
func (ctx *context) find(name string) []semantic.Node {
	result := []semantic.Node{}
	for search := ctx.scope; search != nil; search = search.outer {
		list := search.FindAll(name)
		if len(list) > 0 {
			result = append(result, list...)
		}
	}
	return result
}

// disambiguate takes a list of possible scope values and attempts to see if
// one of them is unambiguously the right choice in the current context.
// For instance, if the values are all enum entries but only one of them
// matches the current enum inference.
func (ctx *context) disambiguate(matches []semantic.Node) []semantic.Node {
	if len(matches) <= 1 {
		return matches
	}
	var enum semantic.Owner
	for test := ctx.scope; test != nil; test = test.outer {
		if test.inferType != nil {
			if e, ok := test.inferType.(*semantic.Enum); ok {
				enum = e
				break
			}
		}
	}
	if enum == nil {
		// No disambiguating enums present
		return matches
	}
	var res *semantic.EnumEntry
	for _, m := range matches {
		if ev, ok := m.(*semantic.EnumEntry); ok {
			if enum == ev.Owner() {
				// We found a disambiguation match
				res = ev
			}
		} else {
			// Non enum match found
			return matches
		}
	}
	if res == nil {
		return matches
	}
	// Matched exactly once
	return []semantic.Node{res}
}

// get searches the scope stack for a bindings that matches the name.
// If it cannot find exactly 1 unambiguous match, it reports an error, and
// nil is returned.
func (ctx *context) get(at ast.Node, name string) semantic.Node {
	if name == "_" {
		return &semantic.Ignore{AST: at}
	}
	matches := ctx.disambiguate(ctx.find(name))
	switch len(matches) {
	case 0:
		ctx.errorf(at, "Unknown identifier %s", name)
		return nil
	case 1:
		return matches[0]
	default:
		possibilities := ""
		for i, m := range matches {
			if i > 0 {
				possibilities += ", "
			}
			switch t := m.(type) {
			case *semantic.EnumEntry:
				possibilities += fmt.Sprintf("%s.%s", t.Owner().Name(), t.Name())
			case *semantic.Parameter:
				possibilities += fmt.Sprintf("parameter %q", t.Name())
			case semantic.Type:
				possibilities += fmt.Sprintf("type %q [%T]", typename(t), t)
			default:
				possibilities += fmt.Sprintf("[%T]%v", t, t)
			}
		}
		ctx.errorf(at, "Ambiguous identifier %q [using %s].\n Could be: %s", name, typename(ctx.scope.inferType), possibilities)
		return nil
	}
}

func (ctx *context) addType(t semantic.Type) {
	name := t.Name()
	if _, present := ctx.types[name]; present {
		ctx.errorf(t, "Duplicate type %s", name)
	}
	ctx.types[name] = t
}

func (ctx *context) findType(at ast.Node, name string) semantic.Type {
	t, found := ctx.types[name]
	if !found {
		return nil
	}
	return t
}

func (ctx *context) addStatement(s semantic.Node) {
	if _, isInvalid := s.(invalid); isInvalid {
		return
	}
	*ctx.scope.block = append(*ctx.scope.block, s)
}

func (ctx *context) uid() uint64 {
	id := ctx.nextId
	ctx.nextId++
	return id
}
