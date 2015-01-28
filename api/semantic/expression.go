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

// Select is the expression form of a switch.
type Select struct {
	AST     *ast.Switch // the underlying syntax node this was built from
	Type    Type        // the return type of the select if valid
	Value   Expression  // the value to match the cases against
	Choices []*Choice   // The set of possible choices to match
}

// ExpressionType implements Expression with the unified type of the choices
func (s *Select) ExpressionType() Type { return s.Type }

// Choice represents a possible choice in a select
type Choice struct {
	AST        *ast.Case    // the underlying syntax node this was built from
	Conditions []Expression // the set of expressions to match the select value against
	Expression Expression   // the expression to use if a condition matches
}

// Local represents an immutable local storage slot, created by a DeclareLocal,
// and referred to by identifiers that resolve to that slot.
type Local struct {
	Declaration *DeclareLocal // the statement that created the local
	Type        Type          // the type of the storage
	Name        string        // the identifier that will resolve to this local
	Value       Expression    // the expression the local was assigned on creation
}

// ExpressionType implements Expression
func (l *Local) ExpressionType() Type { return l.Type }

// UnaryOp represents an operator applied to a single expression.
// It's type depends on the operator and the type of the expression it is
// being applied to.
type UnaryOp struct {
	AST        *ast.UnaryOp // the underlying syntax node this was built from
	Type       Type         // the resolved type of the operation
	Operator   string       // the operator being applied
	Expression Expression   // the expression to apply the operator to
}

// ExpressionType implements Expression
func (b *UnaryOp) ExpressionType() Type { return b.Type }

// BinaryOp represents any operator applied to two arguments.
// The resolved type of the expression depends on the types of the two
// arguments and which operator it represents.
type BinaryOp struct {
	AST      *ast.BinaryOp // the underlying syntax node this was built from
	Type     Type          // the resolved type of this binary expression
	LHS      Expression    // the expression that appears on the left of the operator
	Operator string        // the operator being applied
	RHS      Expression    // the expression that appears on the right of the operator
}

// ExpressionType implements Expression
func (b *BinaryOp) ExpressionType() Type { return b.Type }

// Member is an expression that looks up a field by name from an object.
type Member struct {
	AST    *ast.Member // the underlying syntax node this was built from
	Object Expression  // the object to look up a field of
	Field  *Field      // the field to look up
}

// ExpressionType implements Expression returning the type of the field.
func (m *Member) ExpressionType() Type { return m.Field.Type }

// ArrayIndex represents using the indexing operator on an array type.
type ArrayIndex struct {
	AST       *ast.Index // the underlying syntax node this was built from
	ValueType Type       // the value type of the array being indexed
	Array     Expression // the expression that returns the array to be indexed
	Index     Expression // the index to use on the array
}

// ExpressionType implements Expression returning the value type of the array.
func (i *ArrayIndex) ExpressionType() Type { return i.ValueType }

// MapIndex represents using the indexing operator on an array type.
type MapIndex struct {
	AST       *ast.Index // the underlying syntax node this was built from
	ValueType Type       // the value type of the array being indexed
	Map       Expression // the expression that returns the map to be indexed
	Index     Expression // the index to use on the map
}

// ExpressionType implements Expression returning the value type of the map.
func (i *MapIndex) ExpressionType() Type { return i.ValueType }

// Unknown represents a value that cannot be predicted.
// These values are non-deterministic with regard to the API specification and
// may vary between implementations of the API.
type Unknown struct {
	AST      *ast.Unknown // the underlying syntax node this was built from
	Inferred Expression   // the inferred expression to derive the unknown from the outputs
}

// ExpressionType implements Expression with the inferred type of the unknown.
// If the unknown could not be inferred, it will be of type "any" so allow
// expressions using it to resolve anyway.
func (u Unknown) ExpressionType() Type {
	if u.Inferred == nil {
		return AnyType
	}
	return u.Inferred.ExpressionType()
}

// Represents a length of object expression.
// Object must be of either Array, Map or string type.
// The length expression is allowed to be of any numeric type
type Length struct {
	AST    *ast.Length // the underlying syntax node this was built from
	Object Expression  // the object go get the length of
	Type   Type        // the resolved type of the length operation
}

// ExpressionType implements Expression
func (l Length) ExpressionType() Type {
	return l.Type
}

// Int32 is an int32 that implements Expression so it can be in the semantic graph
type Int32 int32

// ExpressionType implements Expression with a type of Int32Type
func (v Int32) ExpressionType() Type { return Int32Type }

// Uint32 is a uint32 that implements Expression so it can be in the semantic graph
type Uint32 uint32

// ExpressionType implements Expression with a type of Uint32Type
func (v Uint32) ExpressionType() Type { return Uint32Type }

// Int64 is an int64 that implements Expression so it can be in the semantic graph
type Int64 int64

// ExpressionType implements Expression with a type of Int64Type
func (v Int64) ExpressionType() Type { return Int64Type }

// Uint64 is a uint64 that implements Expression so it can be in the semantic graph
type Uint64 uint64

// ExpressionType implements Expression with a type of Uint64Type
func (v Uint64) ExpressionType() Type { return Uint64Type }

// Bool is a bool that implements Expression so it can be in the semantic graph
type Bool bool

// ExpressionType implements Expression with a type of BoolType
func (v Bool) ExpressionType() Type { return BoolType }

// String is a string that implements Expression so it can be in the semantic graph
type String string

// ExpressionType implements Expression with a type of StringType
func (v String) ExpressionType() Type { return StringType }

// Float64 is a float64 that implements Expression so it can be in the semantic graph
type Float64 float64

// ExpressionType implements Expression with a type of Float64Type
func (v Float64) ExpressionType() Type { return Float64Type }

// Float32 is a float32 that implements Expression so it can be in the semantic graph
type Float32 float32

// ExpressionType implements Expression with a type of Float32Type
func (v Float32) ExpressionType() Type { return Float32Type }
