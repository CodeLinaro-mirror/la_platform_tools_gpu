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

// Package semantic holds the set of types used in the abstract semantic graph
// representation of the api language.
package semantic

import "android.googlesource.com/platform/tools/gpu/api/ast"

// API is the root of the ASG, and holds a fully resolved api.
type API struct {
	AST          *ast.API       // the underlying syntax node this was built from
	Enums        []*Enum        // the set of enums
	Classes      []*Class       // the set of classes
	Pseudonyms   []*Pseudonym   // the set of pseudo types
	Macros       []*Function    // the set of macros
	Externs      []*Function    // the external function references
	Functions    []*Function    // the global functions
	Globals      []*Global      // the global variables
	Arrays       []*Array       // the array types used
	StaticArrays []*StaticArray // the fixed size array types used
	Maps         []*Map         // the map types used
	Pointers     []*Pointer     // the pointer types used
	Members                     // a map of name to member for top level symbols
}

type Annotation struct {
	AST       *ast.Annotation // the underlying syntax node this was built from
	Name      string          // the name of the annotation
	Arguments []Expression    // the arguments to the annotation
}

// Annotated is the common interface to objects that can carry annotations.
type Annotated interface {
	// GetAnnotation returns the annotation with the matching name, if present.
	GetAnnotation(name string) *Annotation
}

// Annotations is an array of Annotation objects that implements the Annotated
// interface. It is used as an anonymous field on objects that carry
// annotations.
type Annotations []*Annotation

// GetAnnotation implements the Annotated interface for the Annotations type.
func (a Annotations) GetAnnotation(name string) *Annotation {
	for _, entry := range a {
		if entry.Name == name {
			return entry
		}
	}
	return nil
}

// Global represents a global variable.
type Global struct {
	AST         *ast.Field // the underlying syntax node this was built from
	Annotations            // the annotations applied to this global
	Type        Type       // the type the global stores
	Name        string     // the name of the global
	Default     Expression // the initial value of the global
}

// Implements Expression to return the type stored in the global.
func (g *Global) ExpressionType() Type { return g.Type }
