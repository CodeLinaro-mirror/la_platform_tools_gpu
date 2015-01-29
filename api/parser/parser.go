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

// Package parser implements a parser for converting the api language into
// abstract syntax trees.
package parser

import (
	"android.googlesource.com/platform/tools/gpu/api/ast"
	"android.googlesource.com/platform/tools/gpu/parse"
)

// Parse takes a string containing a complete api description and
// returns the abstract syntax tree representation of it.
// If the string is not syntactically valid, it will also return the
// errors encountered. If errors are returned, the ast returned will be
// the incomplete tree so far, and may not be structurally valid.
func Parse(data string) (*ast.API, []parse.Error) {
	var api *ast.API
	parser := func(p *parse.Parser, cst *parse.Branch) {
		api = requireAPI(p, cst)
	}
	errors := parse.Parse(parser, data, parse.NewSkip("//", "/*", "*/"))
	return api, errors
}
