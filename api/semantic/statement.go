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

import "android.googlesource.com/platform/tools/gpu/api/ast"

// Block represents a collection of statements, used as the body of other
// nodes.
type Block struct {
	AST        *ast.Block // the underlying syntax node this was built from
	Statements []Node     // the set of statements this block represents
}

// Branch represents the basic conditional execution statement.
// If Condition is true we use the True block, otherwise the False block.
type Branch struct {
	AST       *ast.Branch // the underlying syntax node this was built from
	Condition Expression  // the condition to select on
	True      *Block      // use if Condition is true
	False     *Block      // used if Condition is false
}

// Switch represents a resolved ast.Switch statement.
type Switch struct {
	AST     *ast.Switch // the underlying syntax node this was built from
	Value   Expression  // the value to match the cases against
	Cases   []*Case     // the set of case statements to choose from
	Default *Block      // the block to use if no condition matches
}

// Case represents a possible choice in a switch.
type Case struct {
	AST        *ast.Case    // the underlying syntax node this was built from
	Conditions []Expression // the set of expressions to match the switch value against
	Block      *Block       // the block to use if a condition matches
}

// Iteration is the basic looping construct.
// It will set Iterator to each value from Iterable in turn, and run Block for each one.
type Iteration struct {
	AST      *ast.Iteration // the underlying syntax node this was built from
	Iterator *Local         // the iteration control variable
	Iterable Expression     // the expression to iterate over
	Block    *Block         // the block to run for each entry from Iterable
}

// Assign is the only "mutating" construct.
// It assigns the value from the rhs into the slot described by the lhs, as defined
// by the operator.
type Assign struct {
	AST      *ast.Assign // the underlying syntax node this was built from
	LHS      Expression  // the expression that gives the location to store into
	Operator string      // the assignment operator being applied
	RHS      Expression  // the value to store
}

// ArrayAssign represents assigning to a static-array index expression.
type ArrayAssign struct {
	AST      *ast.Assign // the underlying syntax node this was built from
	To       *ArrayIndex // the array index to assign to
	Operator string      // the assignment operator being applied
	Value    Expression  // the value to set in the array
}

// MapAssign represents assigning to a map index expression.
type MapAssign struct {
	AST      *ast.Assign // the underlying syntax node this was built from
	To       *MapIndex   // the map index to assign to
	Operator string      // the assignment operator being applied
	Value    Expression  // the value to set in the map
}

// SliceAssign represents assigning to a slice index expression.
type SliceAssign struct {
	AST      *ast.Assign // the underlying syntax node this was built from
	To       *SliceIndex // the slice index to assign to
	Operator string      // the assignment operator being applied
	Value    Expression  // the value to set in the slice
}

// DeclareLocal represents a local variable declaration statement.
// Variables cannot be modified after declaration.
type DeclareLocal struct {
	AST   *ast.DeclareLocal // the underlying syntax node this was built from
	Local *Local            // the local variable that was declared by this statement
}

// Return represents return statement for a function or macro.
type Return struct {
	AST      *ast.Return // the underlying syntax node this was built from
	Function *Function   // the function this statement returns from
	Value    Expression  // the value to be returned
}

// Fence is a marker to indicate the point between all statements to be
// executed before (pre-fence) the call to the API function and all statements
// to be executed after (post-fence) the call to the API function.
//
// The Statement member is the first statement that is classified as post-fence,
// but may be nil if the fence is being added at the end of a function that has
// no post operations.
//
// Note that some statements are classified as both pre-fence and post-fence,
// and require logic to be executed either side of the API function call.
type Fence struct {
	Statement Node
}
