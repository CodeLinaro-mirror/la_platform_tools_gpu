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

// Block represents a linear sequence of statements, most often the contents
// of a {} pair.
type Block struct {
	CST        *parse.Branch // underlying parse structure for this node
	Docs       *URL          // the url to the documentation for this block
	Statements []Node        // The set of statements that make up the block
}

// Branch represents an «"if" condition { trueblock } "else" { falseblock }» structure.
type Branch struct {
	CST       *parse.Branch // underlying parse structure for this node
	Condition Node          // the condition to use to select which block is active
	True      *Block        // the block to use if condition is true
	False     *Block        // the block to use if condition is false
}

// Iteration represents a «"for" variable "in" iterable { block }» structure.
type Iteration struct {
	CST      *parse.Branch // underlying parse structure for this node
	Variable *Identifier   // the variable to use for the iteration value
	Iterable Node          // the expression that produces the iterable to loop over
	Block    *Block        // the block to run once per item in the iterable
}

// Switch represents a «"switch" value { cases }» structure.
// The first matching case is selected.
// If a switch is used as an expression, the case blocks must all be a single
// expression.
type Switch struct {
	CST   *parse.Branch // underlying parse structure for this node
	Value Node          // the value to match against
	Cases []*Case       // the set of cases to match the value with
}

// Case represents a «"case" conditions: block» structure within a switch statement.
// The conditions are a comma separated list of expressions the switch statement
// value will be compared against.
type Case struct {
	CST        *parse.Branch // underlying parse structure for this node.
	Conditions []Node        // the set of conditions that would select this case
	Block      *Block        // the block to run if this case is selected
}

// Group represents the «(expression)» construct, a single parenthesized expression.
type Group struct {
	CST        *parse.Branch // underlying parse structure for this node
	Expression Node          // the expression within the parentheses
}

// DeclareLocal represents a «name := value» statement that declares a new
// immutable local variable with the specified value and inferred type.
type DeclareLocal struct {
	CST  *parse.Branch // underlying parse structure for this node
	Name *Identifier   // the name to give the new local
	RHS  Node          // the value to store in that local
}

// Assign represents a «location {,+,-}= value» statement that assigns a value to
// an existing mutable location.
type Assign struct {
	CST      *parse.Branch // underlying parse structure for this node
	LHS      Node          // the location to store the value into
	Operator string        // the assignment operator being applied
	RHS      Node          // the value to store
}

// Assert represents the «"assert" condition» statement.
// Used mostly to express the pre-conditions of api commands, such as acceptable
// values for parameters that cannot be expressed in the type system.
type Assert struct {
	CST       *parse.Branch // underlying parse structure for this node.
	Condition Node          // the condition to check, should be true
}

// Length represents the «"len"(value)» construct, were value should be an
// expresssion that returns an object of array, string or map type.
type Length struct {
	CST    *parse.Branch // underlying parse structure for this node.
	Object Node          // the object to query the length of
}

// Return represents the «"return" value» construct, that assigns the value to
// the result slot of the function.
type Return struct {
	CST   *parse.Branch // underlying parse structure for this node.
	Value Node          // the value to return
}
