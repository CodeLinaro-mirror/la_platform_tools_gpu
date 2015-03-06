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

// Class represents a class type declaration of the form
// «"class" name : extension_list { fields }»
type Class struct {
	CST         *parse.Branch // underlying parse structure for this node
	Annotations Annotations   // the annotations applied to the class
	Name        *Identifier   // the name of the class
	Extends     []*Identifier // the set of class names it extends
	Fields      []*Field      // the fields of the class
}

// ClassInitializer represents a class literal declaration, of the form
// «name { field_initializers }»
type ClassInitializer struct {
	CST    *parse.Branch       // underlying parse structure for this node
	Class  *Identifier         // the name of the class instantiate
	Fields []*FieldInitializer // the initializers for the class fields
}

// New represents an expression that allocates a new class instance and returns
// a pointer to it. It takes a class initializer to specify both the type and
// the initial value for the instance.
type New struct {
	CST              *parse.Branch // underlying parse structure for this node
	ClassInitializer *ClassInitializer
}

// Field represents a field of a class or api, with the structure
// «type name = expression»
type Field struct {
	CST         *parse.Branch // underlying parse structure for this node
	Annotations Annotations   // the annotations applied to the field
	Type        Node          // the type the field holds
	Name        *Identifier   // the name of the field
	Default     Node          // the default value expression for the field
}

// FieldInitializer is used as part of a ClassInitializer to specify the value a
// single field should have.
type FieldInitializer struct {
	CST   *parse.Branch // underlying parse structure for this node
	Name  *Identifier   // the name of the field
	Value Node          // the value the field should be given
}

// Cast represents a type coercion expression, of the form «expression "as" type»
type Cast struct {
	CST    *parse.Branch // underlying parse structure for this node
	Object Node          // the value to force the type of
	Type   Node          // the type it should be coerced to
}

// EnumEntry represents a single value in an enumerated type.
type EnumEntry struct {
	CST   *parse.Branch // underlying parse structure for this node
	Owner *Enum         // the enum this entry is a part of
	Name  *Identifier   // the name this entry is given
	Value *Number       // the value of this entry
}

// Enum represents an enumerated type declaration, of the form
// «"enum" name { entries }» where entries is a comma separated list of «name = value»
type Enum struct {
	CST         *parse.Branch // underlying parse structure for this node
	Annotations Annotations   // the annotations applied to the enum
	Name        *Identifier   // the name of the enum
	IsBitfield  bool          // whether this enum represents a bitfield form
	Entries     []*EnumEntry  // the set of valid entries for this enum
	Extends     []*Identifier // deprecated list of enums this extends
}

// Member represents an expressions that access members of objects.
// Always of the form «object.name» where object is an expression.
type Member struct {
	CST    *parse.Branch // underlying parse structure for this node
	Object Node          // the object to get a member of
	Name   *Identifier   // the name of the member to get
}

// Index represents any expression of the form «object[index]»
// Used for arrays, maps and bitfields.
type Index struct {
	CST    *parse.Branch // underlying parse structure for this node
	Object Node          // the object to index
	Index  Node          // the index to lookup
}

// ArrayType represents a type declaration for an array, which looks
// like «"array"<value_type>»
type ArrayType struct {
	CST       *parse.Branch // underlying parse structure for this node
	ValueType Node          // the type stored as elements of the array
}

// StaticArrayType represents a type declaration for a constant size array,
// which looks like «type[dimensions]»
// dimensions is a comma separated list of dimensions, for declaring
// multidimensional arrays, for instance f32[4,4] for a matrix.
type StaticArrayType struct {
	CST        *parse.Branch // underlying parse structure for this node
	ValueType  Node          // The type to store in the array
	Dimensions []Node        // the dimensions of the array
}

// MapType represents a type declaration of a map, of the form
// «map<key_type, value_type>»
type MapType struct {
	CST       *parse.Branch // underlying parse structure for this node
	KeyType   Node          // the type used to index the map
	ValueType Node          // the type stored against the index in the map
}

// PointerType represents a pointer type declaration, of the form «type*»
type PointerType struct {
	CST *parse.Branch // underlying parse structure for this node
	To  Node          // the underlying type this pointer points to
}

// Alias represents a weak type alias, with structure «"alias" type name».
// An alias does not declare a new type, just a reusable name for a common type.
type Alias struct {
	CST         *parse.Branch // underlying parse structure for this node
	Annotations Annotations   // the annotations applied to the alias
	Name        *Identifier   // the name of the alias
	To          Node          // the type it is an alias for
}

// Pseudonym declares a new type in terms of another type.
// Has the form «"type" type name»
// Pseydonyms are proper types, but the underlying type can be discovered.
type Pseudonym struct {
	CST         *parse.Branch // underlying parse structure for this node
	Annotations Annotations   // the annotations applied to the type
	Name        *Identifier   // the name of the type
	To          Node          // the underlying type
}
