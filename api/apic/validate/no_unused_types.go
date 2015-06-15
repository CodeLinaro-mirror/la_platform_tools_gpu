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

package validate

import (
	"android.googlesource.com/platform/tools/gpu/api/semantic"
	"android.googlesource.com/platform/tools/gpu/parse"
)

// noUnusedTypes verifies that all declared types are used.
func noUnusedTypes(apiName string, api *semantic.API) []error {
	used := map[semantic.Type]bool{}
	tokens := map[semantic.Type]parse.Token{}

	// Gather all declared types
	for _, t := range api.Classes {
		used[t] = false
		tokens[t] = t.AST.CST.Token()
	}
	for _, t := range api.Enums {
		used[t] = false
		tokens[t] = t.AST.CST.Token()
	}
	for _, t := range api.Pseudonyms {
		used[t] = false
		tokens[t] = t.AST.CST.Token()
	}

	// Functions for marking types as used
	var markUsed func(t semantic.Type)
	markUsed = func(t semantic.Type) {
		if used[t] {
			return
		}
		used[t] = true
		switch t := t.(type) {
		case *semantic.Reference:
			markUsed(t.To)
		case *semantic.Slice:
			markUsed(t.To)
		case *semantic.StaticArray:
			markUsed(t.ValueType)
		case *semantic.Map:
			markUsed(t.ValueType)
			markUsed(t.KeyType)
		case *semantic.Pointer:
			markUsed(t.To)
		case *semantic.Pseudonym:
			markUsed(t.To)
		case *semantic.Class:
			for _, f := range t.Fields {
				markUsed(f.Type)
			}
			for _, e := range t.Extends {
				markUsed(e)
			}
		case *semantic.Enum:
			for _, e := range t.Extends {
				markUsed(e)
			}
		}
	}
	var traverseExpression func(e semantic.Expression)
	traverseExpression = func(e semantic.Expression) {
		switch e := e.(type) {
		case *semantic.UnaryOp:
			traverseExpression(e.Expression)
		case *semantic.BinaryOp:
			traverseExpression(e.LHS)
			traverseExpression(e.RHS)
		case *semantic.ClassInitializer:
			markUsed(e.Class)
			// TODO: Insert more expression cases
		}
	}
	var traverseStatement func(s interface{})
	traverseStatement = func(s interface{}) {
		switch s := s.(type) {
		case *semantic.DeclareLocal:
			traverseExpression(s.Local.Value)
		case *semantic.Assign:
			traverseExpression(s.LHS)
			traverseExpression(s.RHS)
		case *semantic.Call:
			for _, a := range s.Arguments {
				traverseExpression(a)
			}
			// TODO: Insert more statement cases
		}
	}
	var traverseFunction func(f *semantic.Function)
	traverseFunction = func(f *semantic.Function) {
		for _, p := range f.FullParameters {
			markUsed(p.Type)
		}
		for _, s := range f.Block.Statements {
			traverseStatement(s)
		}
	}

	// Traverse the API finding all used types
	for _, g := range api.Globals {
		markUsed(g.Type)
	}
	for _, f := range api.Functions {
		traverseFunction(f)
	}
	for _, c := range api.Classes {
		for _, m := range c.Methods {
			traverseFunction(m)
		}
	}
	for _, c := range api.Pseudonyms {
		for _, m := range c.Methods {
			traverseFunction(m)
		}
	}

	// Report all types declared but not used as errors
	errors := []error{}
	for t, used := range used {
		if !used {
			e := err(apiName, tokens[t], "%s declared but never used", t.Typename())
			errors = append(errors, e)
		}
	}
	return errors
}
