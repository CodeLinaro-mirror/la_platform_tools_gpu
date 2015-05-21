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

// Package ast holds the set of types used in the abstract syntax tree
// representation of the api language.
package ast

import "android.googlesource.com/platform/tools/gpu/parse"

// API is the root of the AST tree, and constitutes one entire parsed file.
// It holds the set of top level AST nodes, grouped by type.
type API struct {
	CST        *parse.Branch // underlying parse structure for this node
	Imports    []*Import     // api files imported with the "import" keyword
	Macros     []*Function   // functions declared with the "macro" keyword
	Externs    []*Function   // functions declared with the "extern" keyword
	Commands   []*Function   // functions declared with the "cmd" keyword
	Pseudonyms []*Pseudonym  // strong type aliases declared with the "type" keyword
	Aliases    []*Alias      // weak type aliases declared with the "alias" keyword
	Enums      []*Enum       // enumerated types, declared with the "enum" keyword
	Classes    []*Class      // class types, declared with the "class" keyword
	Fields     []*Field      // variables declared at the global scope
}

// Annotation is the AST node that represents «@name(arguments)» constructs
type Annotation struct {
	CST       *parse.Branch // underlying parse structure for this node
	Name      *Identifier   // the name part (between the @ and the brackets)
	Arguments []Node        // the list of arguments (the bit in brackets)
}

// Annotations represents the set of Annotation objects that apply to another
// AST node.
type Annotations []*Annotation

// Invalid is used when an error was encountered in the parsing, but we want to
// keep going. If there are no errors, this will never be in the tree.
type Invalid struct{}

// Import is the AST node that represents «import name "path"» constructs
type Import struct {
	CST         *parse.Branch // underlying parse structure for this node
	Annotations Annotations   // the annotations applied to the import
	Name        *Identifier   // the name to import an api file as
	Path        *String       // the relative path to the api file
}
