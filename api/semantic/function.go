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

package semantic

import (
	"net/url"

	"android.googlesource.com/platform/tools/gpu/api/ast"
)

// Docs represents a documentation link
type Docs url.URL

func (d *Docs) String() string { return (*url.URL)(d).String() }

// Function represents function like objects in the semantic graph.
type Function struct {
	AST            *ast.Function // the underlying syntax node this was built from
	Annotations                  // the annotations applied to the function
	Name           string        // the name of the function
	Docs           Docs          // the documentation link for the function
	Owner          Type          // the owner of the function
	Return         *Parameter    // the return parameter
	This           *Parameter    // the this parameter, missing for non method functions
	FullParameters []*Parameter  // all the parameters, including This at the start if valid, and Return at the end if not void
	Outputs        []*Parameter  // only the output parameters
	Block          *Block        // the body of the function, missing for externs
	Signature      *Signature    // the type signature of the function
}

// CallParameters returns the full set of parameters with the return value
// filtered out.
func (f *Function) CallParameters() []*Parameter {
	if f.Return.Type == VoidType {
		return f.FullParameters
	}
	return f.FullParameters[0 : len(f.FullParameters)-1]
}

// Parameter represents a single parameter declaration for a function.
type Parameter struct {
	AST         *ast.Parameter // the underlying syntax node this was built from
	Annotations                // the annotations applied to the parameter
	Function    *Function      // the function this parameter belongs to
	Name        string         // the name of the parameter
	Input       bool           // true if the parameter is an input
	Output      bool           // true if the parameter is an output
	Type        Type           // the type of the parameter
}

// ExpressionType implements Expression for parameter lookup.
func (p *Parameter) ExpressionType() Type { return p.Type }

// IsThis returns true if this parameter is the This parameter of it's function.
func (p *Parameter) IsThis() bool { return p == p.Function.This }

// Observed represents the final observed value of an output parameter.
// It is never produced directly from the ast, but is inserted when inferring
// the value of an unknown from observed outputs.
type Observed struct {
	Parameter *Parameter // the output parameter to infer from
}

// ExpressionType implements Expression for observed parameter lookup.
func (e *Observed) ExpressionType() Type { return e.Parameter.Type }

// Callable wraps a Function declaration into a first class function value expression,
// optionally binding to an object if its a method.
type Callable struct {
	Object   Expression // the object to use as the this parameter for a method
	Function *Function  // the function this expression represents
}

// ExpressionType implements Expression returning the function type signature.
func (c *Callable) ExpressionType() Type { return c.Function.Signature }

// Call represents a function call. It binds an Callable to the arguments it
// will be passed.
type Call struct {
	AST       *ast.Call    // the underlying syntax node this was built from
	Target    *Callable    // the function expression this invokes
	Arguments []Expression // the arguments to pass to the function
	Type      Type         // the return type of the call
}

// ExpressionType implements Expression returning the underlying function return type.
func (c *Call) ExpressionType() Type { return c.Type }

// Signature represents a callable type signature
type Signature struct {
	Name      string // the full type name
	Return    Type   // the return type of the callable
	Arguments []Type // the required callable arguments
}

func (t Signature) Typename() string      { return t.Name }
func (Signature) Member(name string) Node { return nil }
