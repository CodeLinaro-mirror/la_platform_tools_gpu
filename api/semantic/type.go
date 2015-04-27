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

// Type is the interface to any object that can act as a type to the api
// langauge.
type Type interface {
	Node
	Typename() string        // returns the full name of the type, must be unique
	Member(Name string) Node // looks up a member by name from a type
}

// Expression represents anything that can act as an expression in the api
// language, it must be able to correctly report the type of value it would
// return if executed.
type Expression interface {
	Node
	ExpressionType() Type // returns the expression value type.
}

// Members wraps a map and implements part of the Type interface.
// It is used as a mixin helper.
type Members map[string]Node

// Member returns the entry in the map that matches name, or nil if none does.
func (t Members) Member(name string) Node {
	m, ok := t[name]
	if !ok {
		return nil
	}
	return m
}

// Class represents an api class construct.
type Class struct {
	AST         *ast.Class  // the underlying syntax node this was built from
	Annotations             // the annotations applied to this class
	Name        string      // the name of the class
	Docs        []string    // the documentation for the class
	Extends     []*Class    // the classes this extends
	ExtendedBy  []*Class    // the classes that declared they extended this class
	Fields      []*Field    // the set of fields the class declares
	Methods     []*Function // the set of functions associated with the class
	Members                 // the name->member map, to implement Type
}

// Implements Type to return the class name as the type name
func (t Class) Typename() string { return t.Name }

// Field represents a field entry in a class.
type Field struct {
	AST         *ast.Field // the underlying syntax node this was built from
	Annotations            // the annotations applied to this field
	Class       *Class     // the class this field belongs to
	Type        Type       // the type the field stores
	Name        string     // the name of the field
	Docs        []string   // the documentation for the field
	Default     Expression // the default value of the field
}

// Implements Expression to return the type stored in the field.
func (f *Field) ExpressionType() Type { return f.Type }

// ClassInitializer represents an expression that can assign values to multiple
// fields of a class.
type ClassInitializer struct {
	AST    *ast.ClassInitializer // the underlying syntax node this was built from
	Class  *Class                // the class to initialize
	Fields []*FieldInitializer   // the set of field assignments
}

// ExpressionType implements Expression returning the class type being initialized.
func (c *ClassInitializer) ExpressionType() Type {
	if c.Class != nil {
		return c.Class
	} else {
		return nil
	}
}

// FieldInitializer
type FieldInitializer struct {
	AST   *ast.FieldInitializer // the underlying syntax node this was built from
	Field *Field                // the field to assign to
	Value Expression            // the value to assign
}

// New represents an expression that allocates a new class instance.
type New struct {
	AST         *ast.New          // the underlying syntax node this was built from
	Initializer *ClassInitializer // The initialization for the new instance
	Type        Type              // The pointer type returned from the new
}

// ExpressionType implements Expression returning a pointer to the class type being initialized.
func (n *New) ExpressionType() Type { return n.Type }

// Cast represents a type coercion expression.
// It reports it's type as the one specified, rather than the expression it
// wraps.
type Cast struct {
	AST    *ast.Cast  // the underlying syntax node this was built from
	Object Expression // the actual expression being wrapped
	Type   Type       // the type to coerce the expression to
}

// ExpressionType implements Expression returning the type being cast to.
func (c *Cast) ExpressionType() Type { return c.Type }

// Enum represents the api enum construct.
type Enum struct {
	AST         *ast.Enum    // the underlying syntax node this was built from
	Annotations              // the annotations applied to this enum
	Name        string       // the type name of the enum
	Docs        []string     // the documentation for the enum
	IsBitfield  bool         // whether this enum is actually a bitfield
	Extends     []*Enum      // the enums this enum extends
	Entries     []*EnumEntry // the entries of this enum
	AllEntries  []*EnumEntry // the flattened list of all entries including inherited ones
}

// Implements Type to return the enum name as the type name
func (t Enum) Typename() string { return t.Name }

// Implements Type returning the matching enum entry if there is one.
func (t Enum) Member(name string) Node {
	for _, e := range t.AllEntries {
		if e.Name == name {
			return e
		}
	}
	return nil
}

// EnumEntry represents a single entry in an Enum.
type EnumEntry struct {
	AST   *ast.EnumEntry // the underlying syntax node this was built from
	Enum  *Enum          // the enum this entry belongs to
	Name  string         // the name of this entry
	Docs  []string       // the documentation for the enum entry
	Value uint32         // the value this entry represents
}

// ExpressionType implements Expression returning the enum type.
func (e *EnumEntry) ExpressionType() Type {
	if e.Enum != nil {
		return e.Enum
	} else {
		return nil
	}
}

// Pseudonym represents the type construct.
// It acts as a type in it's own right that can carry methods, but is defined
// in terms of another type.
type Pseudonym struct {
	AST         *ast.Pseudonym // the underlying syntax node this was built from
	Annotations                // the annotations applied to this pseudonym
	Name        string         // the type name
	Docs        []string       // the documentation for the pseudonym
	To          Type           // the underlying type
	Methods     []*Function    // the methods added directly to the pseudonym
	Members                    // the direct members
}

// Implements Type to return the type name
func (t Pseudonym) Typename() string { return t.Name }

// Implements Type returning the direct member if it has it, otherwise
// delegating the lookup to the underlying type.
func (t Pseudonym) Member(name string) Node {
	m := t.Members.Member(name)
	if m == nil {
		m = t.To.Member(name)
	}
	return m
}

// Array represents an array type declaration.
type Array struct {
	Name      string // the full name of the array type
	ValueType Type   // the value type stored in the array
}

func (t Array) Typename() string        { return t.Name }
func (t Array) Member(name string) Node { return nil }

// StaticArray represents a multi-dimensional fixed size array type.
type StaticArray struct {
	Name      string // the full type name
	ValueType Type   // the storage type of the elements
	Size      uint32 // the dimension of the array
}

func (t StaticArray) Typename() string        { return t.Name }
func (t StaticArray) Member(name string) Node { return nil }

// Map represents an api map type declaration.
type Map struct {
	Name      string // the full type name
	KeyType   Type   // the type used as an indexing key
	ValueType Type   // the type stored in the map
	Members          // holds the map built-in methods
}

func (t Map) Typename() string { return t.Name }

// Pointer represents an api pointer type declaration.
type Pointer struct {
	Name  string // the full type name
	To    Type   // the type this is a pointer to
	Array bool   // points to multiple elements, rather than one
}

func (t Pointer) Typename() string { return t.Name }
func (t Pointer) Member(name string) Node {
	if t.Array {
		return nil
	}
	return t.To.Member(name)
}

// Buffer represents a state pointer type declaration.
type Buffer struct {
	Name      string // the full type name
	To        Type   // the type this is a pointer to
	Array     bool   // points to multiple elements, rather than one
	FakeArray *Array // TODO:Remove - Hold a fake array for now as a schema compatibility measure
}

func (t Buffer) Typename() string { return t.Name }
func (t Buffer) Member(name string) Node {
	if t.Array {
		return nil
	}
	return t.To.Member(name)
}

// Builtin represents one of the primitive types.
type Builtin struct {
	Name string // the primitive type name
}

func (t Builtin) Typename() string        { return t.Name }
func (t Builtin) Member(name string) Node { return nil }

func builtin(name string) *Builtin {
	b := &Builtin{Name: name}
	BuiltinTypes = append(BuiltinTypes, b)
	return b
}

var (
	// These are all the fundamental primitive types of the api language

	// Special types
	VoidType   = builtin("void")
	AnyType    = builtin("any")
	StringType = builtin("string")
	// Unsized primitives
	BoolType = builtin("bool")
	CharType = builtin("char")
	IntType  = builtin("int")
	UintType = builtin("uint")
	// Fixed size integer forms
	Int8Type   = builtin("s8")
	Uint8Type  = builtin("u8")
	Int16Type  = builtin("s16")
	Uint16Type = builtin("u16")
	Int32Type  = builtin("s32")
	Uint32Type = builtin("u32")
	Int64Type  = builtin("s64")
	Uint64Type = builtin("u64")
	// Floating point forms
	Float32Type = builtin("f32")
	Float64Type = builtin("f64")
)

var BuiltinTypes []*Builtin
