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

/*
Package glsl contains routines for manipulation of OpenGL ES Shading Language programs.

It exposes functions for parsing, serializing and evaluating GLES Shading Language programs.
While this package contains a number of sub-packages, the only sub-package which is expected
to be imported directly is the ast package, which contains the definitions of the AST of the
parsed program. The main functionality of the other packages is exposed through the functions
of this package.
*/
package glsl

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/gfxapi/gles/glsl/ast"
	"android.googlesource.com/platform/tools/gpu/gfxapi/gles/glsl/evaluator"
	"android.googlesource.com/platform/tools/gpu/gfxapi/gles/glsl/parser"
	"android.googlesource.com/platform/tools/gpu/gfxapi/gles/glsl/sema"
)

// Parse preprocesses and parses an OpenGL ES Shading language program present in the first
// argument. The second argument specifies the language, whose syntax to employ during parsing.
// The parsed AST is returned in the first result. If any parsing errors are encountered, they
// are returned in the second result.
func Parse(src string, lang ast.Language) (program interface{}, err []error) {
	return parser.Parse(src, lang, evaluator.EvaluatePreprocessorExpression)
}

// Formatter is a helper function which turns any AST node into something that can be printed with
// %v. The returned object's default format will print the tree under the ast node in a reindented
// form. The alternate format flag (%#v) will print the node while preserving original whitespace,
// if this is present in the ***Cst nodes of the tree.
func Formatter(node interface{}) fmt.Formatter { return parser.Formatter(node) }

// Analyze performs semantic analysis on the parsed program AST. It computes the types of all program
// expression, array sizes and values of constant variables. Any encountered errors are returned
// as a result.
func Analyze(program interface{}) (err []error) { return sema.Analyze(program, evaluator.Evaluate) }
