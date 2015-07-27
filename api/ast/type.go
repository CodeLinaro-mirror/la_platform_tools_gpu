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
// «"class" name { fields }»
type Class struct {
	CST         *parse.Branch // underlying parse structure for this node
	Annotations Annotations   // the annotations applied to the class
	Name        *Identifier   // the name of the class
	Fields      []*Field      // the fields of the class
}

func (t Class) Node() parse.Node { return t.CST }

// Field represents a field of a class or api, with the structure
// «type name = expression»
type Field struct {
	CST         *parse.Branch // underlying parse structure for this node
	Annotations Annotations   // the annotations applied to the field
	Type        Node          // the type the field holds
	Name        *Identifier   // the name of the field
	Default     Node          // the default value expression for the field
}

func (t Field) Node() parse.Node { return t.CST }

// EnumEntry represents a single value in an enumerated type.
type EnumEntry struct {
	CST   *parse.Branch // underlying parse structure for this node
	Owner *Enum         // the enum this entry is a part of
	Name  *Identifier   // the name this entry is given
	Value *Number       // the value of this entry
}

func (t EnumEntry) Node() parse.Node { return t.CST }

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

func (t Enum) Node() parse.Node { return t.CST }

// IndexedType represents a type declaration with an indexing suffix,
// which looks like «type[index]»
type IndexedType struct {
	CST       *parse.Branch // underlying parse structure for this node
	ValueType Node          // The element type exposed by the indexed type
	Index     Node          // the index of the type
}

func (t IndexedType) Node() parse.Node { return t.CST }

// PreConst represents a pre-const type declaration, of the form «const type»
type PreConst struct {
	CST  *parse.Branch // underlying parse structure for this node
	Type Node          // the underlying type that is constant
}

func (t PreConst) Node() parse.Node { return t.CST }

// PointerType represents a pointer type declaration, of the form «type*»
type PointerType struct {
	CST   *parse.Branch // underlying parse structure for this node
	To    Node          // the underlying type this pointer points to
	Const bool          // whether the pointer type has the post-const modifier applied
}

func (t PointerType) Node() parse.Node { return t.CST }

// Alias represents a weak type alias, with structure «"alias" type name».
// An alias does not declare a new type, just a reusable name for a common type.
type Alias struct {
	CST         *parse.Branch // underlying parse structure for this node
	Annotations Annotations   // the annotations applied to the alias
	Name        *Identifier   // the name of the alias
	To          Node          // the type it is an alias for
}

func (t Alias) Node() parse.Node { return t.CST }

// Pseudonym declares a new type in terms of another type.
// Has the form «"type" type name»
// Pseydonyms are proper types, but the underlying type can be discovered.
type Pseudonym struct {
	CST         *parse.Branch // underlying parse structure for this node
	Annotations Annotations   // the annotations applied to the type
	Name        *Identifier   // the name of the type
	To          Node          // the underlying type
}

func (t Pseudonym) Node() parse.Node { return t.CST }

// Imported represents an imported type name.
type Imported struct {
	CST  *parse.Branch // underlying parse structure for this node
	From *Identifier   // the import this name is from
	Name *Identifier   // the name being imported
}

func (t Imported) Node() parse.Node { return t.CST }

// Definition declares a new named literal, has the form «"define" name value».
type Definition struct {
	CST         *parse.Branch // underlying parse structure for this node
	Annotations Annotations   // the annotations applied to this definition
	Name        *Identifier   // the name of this definition
	Expression  Node          // the expression this definition expands to
}

func (t Definition) Node() parse.Node { return t.CST }
