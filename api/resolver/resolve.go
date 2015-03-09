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

// Package resolver implements a semantic resolving for the api language.
// It is responsible for converting from an abstract syntax tree to a typed
// semantic graph ready for code generation.
package resolver

import (
	"android.googlesource.com/platform/tools/gpu/api/ast"
	"android.googlesource.com/platform/tools/gpu/api/semantic"
	"android.googlesource.com/platform/tools/gpu/parse"
)

// ASTToSemantic is a relational map of AST nodes to semantic nodes.
type ASTToSemantic map[ast.Node]semantic.Node

// Resolve takes a valid ast as produced by the parser and converts it to the
// semantic graph form.
// If the ast is not fully valid (ie there were parse errors) then the results
// are undefined, and may include null pointer access.
// If there are semantic problems with the ast, Resolve will return the set of
// errors it finds, and the returned graph may be incomplete/invalid.
func Resolve(compiled *ast.API) (*semantic.API, parse.ErrorList, ASTToSemantic) {
	ctx := &context{
		api: &semantic.API{
			AST:     compiled,
			Members: semantic.Members{},
		},
		types:    map[string]semantic.Type{},
		scope:    &scope{entries: map[string][]semantic.Node{}},
		mappings: make(ASTToSemantic),
	}
	func() {
		defer func() {
			err := recover()
			if err != nil && err != parse.AbortParse {
				panic(err)
			}
		}()
		// Register all the built in symbols
		for _, t := range semantic.BuiltinTypes {
			ctx.addType(t)
		}
		api(ctx, ctx.api)
	}()
	return ctx.api, ctx.errors, ctx.mappings
}
