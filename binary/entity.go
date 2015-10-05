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

package binary

import (
	"fmt"
	"strings"
)

// Entity represents the encodable type information for an object.
// In it's compact form, the entity contains only the information strictly required to generate it's signature, and not
// any of metadata or display names.
type Entity struct {
	Package  string    // The package that declared the struct.
	Display  string    // The display name of the class, not set in compact form.
	Identity string    // The true name of the class.
	Version  string    // The version string of the class, if set.
	Exported bool      // Whether the class is exported from it's package, not set in compact form.
	Fields   FieldList // Descriptions of the fields of the class.
	Metadata []Object  // The metadata for the class, not set in compact form

	// signature is the cached result of a call to Signature.
	// It is never encoded to a stream.
	signature string
}

// Name returns the name of the Entity.
func (e *Entity) Name() string {
	if e.Display != "" {
		return e.Display
	}
	return e.Identity
}

type Signature string

// Signature returns a canonical string representations of an entities signature.
// If two entities have the same Signature, the are assumed to represent the same type
func (e *Entity) Signature() Signature {
	if e.signature == "" {
		// Internally signature is implemented by the fmt.Formatter interface with the format specifier 'z'
		e.signature = fmt.Sprintf("%z", e)
	}
	return Signature(e.signature)
}

// Format implements the fmt.Formatter interface
func (e *Entity) Format(f fmt.State, c rune) {
	// if c is 'z' then we are printing in signature format.
	// this is an internal implementation detail, the only code that should ever do this is the Signature method.
	fmt.Fprint(f, e.Package, ".", e.Identity)
	if e.Version != "" {
		fmt.Fprint(f, '@', e.Version)
	}
	if c != 'z' && e.Display != "" {
		fmt.Fprint(f, '(', e.Display, ")")
	}
	fmt.Fprint(f, "{")
	for i, field := range e.Fields {
		if i != 0 {
			fmt.Fprint(f, ",")
		}
		if c != 'z' && field.Declared != "" {
			fmt.Fprint(f, field.Declared, " ")
		}
		field.Type.Format(f, c)
	}
	fmt.Fprint(f, "}")
}

// FieldList is a slice of fields.
type FieldList []Field

// Field represents a name/type pair for a field in an Object.
type Field struct {
	Declared string // The name of the field, not set in compact form.
	Type     Type   // The type stored in the field.
}

// Type represents the common interface to all type objects in the schema.
type Type interface {
	String() string         // The true name of the type.
	Representation() string // The encoded representation of the type.
	EncodeValue(e Encoder, value interface{})
	DecodeValue(d Decoder) interface{}
	Format(f fmt.State, c rune)
}

func trimPackage(n string) string {
	i := strings.LastIndex(n, ".")
	if i < 0 {
		return n
	}
	return n[i+1:]
}

func (f Field) Name() string {
	if f.Declared == "" {
		return trimPackage(f.Type.String())
	}
	return f.Declared
}

// Find searches the field list of the field with the specified name, returning
// the index of the field if found, otherwise -1.
func (l FieldList) Find(name string) int {
	for i, f := range l {
		if f.Name() == name {
			return i
		}
	}
	return -1
}
