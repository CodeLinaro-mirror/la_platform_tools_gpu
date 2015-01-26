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

package ast

import "android.googlesource.com/platform/tools/gpu/parse"

// Function represents the declaration of a callable entity, any of "cmd", "extern" or "macro".
// Its structure is «return_type name(parameters) body»
// where parameters is a comma separated list and body is an optional block.
type Function struct {
	CST         *parse.Branch // underlying parse structure for this node
	Annotations Annotations   // the annotations applied to this function
	Name        *Identifier   // the name of the function
	Parameters  []*Parameter  // the parameters the function takes
	Block       *Block        // the body of the function if present
}

// Parameter represents a single parameter in the set of parameters for a Function.
// It has the structure «["in"|"out"|"inout"|"this"] type name»
type Parameter struct {
	CST         *parse.Branch // underlying parse structure for this node
	Annotations Annotations   // the annotations applied to this parameter
	Input       bool          // true if the parameter is an input
	Output      bool          // true if the parameters is an output
	This        bool          // true if the parameter is the this pointer of a method
	Type        interface{}   // the type of the parameter
	Name        *Identifier   // the name the parameter as exposed to the body
}

// Call is an expression that invokes a function with a set of arguments.
// It has the structure «target(arguments)» where target must be a function and
// arguments is a comma separated list of expressions.
type Call struct {
	CST       *parse.Branch // underlying parse structure for this node
	Target    interface{}   // the function to invoke
	Arguments []interface{} // the arguments to the function
}
