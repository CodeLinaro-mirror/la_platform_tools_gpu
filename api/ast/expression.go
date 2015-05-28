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
	Statements []Node        // The set of statements that make up the block
}

func (t Block) Fragment() parse.Fragment { return t.CST }

// Branch represents an «"if" condition { trueblock } "else" { falseblock }» structure.
type Branch struct {
	CST       *parse.Branch // underlying parse structure for this node
	Condition Node          // the condition to use to select which block is active
	True      *Block        // the block to use if condition is true
	False     *Block        // the block to use if condition is false
}

func (t Branch) Fragment() parse.Fragment { return t.CST }

// Iteration represents a «"for" variable "in" iterable { block }» structure.
type Iteration struct {
	CST      *parse.Branch // underlying parse structure for this node
	Variable *Identifier   // the variable to use for the iteration value
	Iterable Node          // the expression that produces the iterable to loop over
	Block    *Block        // the block to run once per item in the iterable
}

func (t Iteration) Fragment() parse.Fragment { return t.CST }

// Switch represents a «"switch" value { cases }» structure.
// The first matching case is selected.
// If a switch is used as an expression, the case blocks must all be a single
// expression.
type Switch struct {
	CST   *parse.Branch // underlying parse structure for this node
	Value Node          // the value to match against
	Cases []*Case       // the set of cases to match the value with
}

func (t Switch) Fragment() parse.Fragment { return t.CST }

// Case represents a «"case" conditions: block» structure within a switch statement.
// The conditions are a comma separated list of expressions the switch statement
// value will be compared against.
type Case struct {
	CST        *parse.Branch // underlying parse structure for this node.
	Conditions []Node        // the set of conditions that would select this case
	Block      *Block        // the block to run if this case is selected
}

func (t Case) Fragment() parse.Fragment { return t.CST }

// Group represents the «(expression)» construct, a single parenthesized expression.
type Group struct {
	CST        *parse.Branch // underlying parse structure for this node
	Expression Node          // the expression within the parentheses
}

func (t Group) Fragment() parse.Fragment { return t.CST }

// DeclareLocal represents a «name := value» statement that declares a new
// immutable local variable with the specified value and inferred type.
type DeclareLocal struct {
	CST  *parse.Branch // underlying parse structure for this node
	Name *Identifier   // the name to give the new local
	RHS  Node          // the value to store in that local
}

func (t DeclareLocal) Fragment() parse.Fragment { return t.CST }

// Assign represents a «location {,+,-}= value» statement that assigns a value to
// an existing mutable location.
type Assign struct {
	CST      *parse.Branch // underlying parse structure for this node
	LHS      Node          // the location to store the value into
	Operator string        // the assignment operator being applied
	RHS      Node          // the value to store
}

func (t Assign) Fragment() parse.Fragment { return t.CST }

// Return represents the «"return" value» construct, that assigns the value to
// the result slot of the function.
type Return struct {
	CST   *parse.Branch // underlying parse structure for this node.
	Value Node          // the value to return
}

func (t Return) Fragment() parse.Fragment { return t.CST }

// Member represents an expressions that access members of objects.
// Always of the form «object.name» where object is an expression.
type Member struct {
	CST    *parse.Branch // underlying parse structure for this node
	Object Node          // the object to get a member of
	Name   *Identifier   // the name of the member to get
}

func (t Member) Fragment() parse.Fragment { return t.CST }

// Index represents any expression of the form «object[index]»
// Used for arrays, maps and bitfields.
type Index struct {
	CST    *parse.Branch // underlying parse structure for this node
	Object Node          // the object to index
	Index  Node          // the index to lookup
}

func (t Index) Fragment() parse.Fragment { return t.CST }
