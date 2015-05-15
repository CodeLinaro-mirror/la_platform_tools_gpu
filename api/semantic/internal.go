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

// Length represents a length of object expression.
// Object must be of either pointer, slice, map or string type.
// The length expression is allowed to be of any numeric type
type Length struct {
	AST    *ast.Call  // the underlying syntax node this was built from
	Object Expression // the object go get the length of
	Type   Type       // the resolved type of the length operation
}

// ExpressionType implements Expression
func (l *Length) ExpressionType() Type { return l.Type }

// Assert represents a runtime assertion.
// Assertions are also used to infer required behavior from the expressions.
type Assert struct {
	AST       *ast.Call  // the underlying syntax node this was built from
	Condition Expression // the condition is being asserted must be true
}

// ExpressionType implements Expression
func (a *Assert) ExpressionType() Type { return VoidType }

// Cast represents a type reinterpret expresssion.
type Cast struct {
	AST    *ast.Call  // the underlying syntax node this was built from
	Object Expression // the expression to cast the result of
	Type   Type       // the type to cast to
}

// ExpressionType implements Expression
func (c *Cast) ExpressionType() Type { return c.Type }

// New represents a call to new.
type New struct {
	AST  *ast.Call // the underlying syntax node this was built from
	Type *Reference
}

// ExpressionType implements Expression
func (n *New) ExpressionType() Type { return n.Type }

// Create represents a call to new on a class type.
type Create struct {
	AST         *ast.Call // the underlying syntax node this was built from
	Type        *Reference
	Initializer *ClassInitializer
}

// ExpressionType implements Expression
func (n *Create) ExpressionType() Type { return n.Type }

// Make represents a call to make.
type Make struct {
	AST  *ast.Call // the underlying syntax node this was built from
	Type *Slice
	Size Expression
}

// ExpressionType implements Expression
func (m *Make) ExpressionType() Type { return m.Type }

// Clone represents a call to make.
type Clone struct {
	AST   *ast.Call // the underlying syntax node this was built from
	Slice Expression
	Type  *Slice
}

// ExpressionType implements Expression
func (m *Clone) ExpressionType() Type { return m.Type }

// Read represents a call to make.
type Read struct {
	AST   *ast.Call // the underlying syntax node this was built from
	Slice Expression
}

// ExpressionType implements Expression
func (*Read) ExpressionType() Type { return VoidType }

// Write represents a call to make.
type Write struct {
	AST   *ast.Call // the underlying syntax node this was built from
	Slice Expression
}

// ExpressionType implements Expression
func (*Write) ExpressionType() Type { return VoidType }

// Copy represents a call to make.
type Copy struct {
	AST *ast.Call // the underlying syntax node this was built from
	Src Expression
	Dst Expression
}

// ExpressionType implements Expression
func (*Copy) ExpressionType() Type { return VoidType }
