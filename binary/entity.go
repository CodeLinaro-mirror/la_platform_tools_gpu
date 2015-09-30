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

import "strings"

// Entity represents the encodable type information for an object.
type Entity struct {
	TypeID   ID        // The unique type identifier for the class.
	Package  string    // The package that declared the struct.
	Display  string    // The display name of the class, does not affect the signature.
	Identity string    // The true name of the class.
	Version  string    // The version string of the class, if set.
	Exported bool      // Whether the class is exported from it's package
	Fields   FieldList // Descriptions of the fields of the class.
	Metadata []Object  // The metadata for the class.
}

func (e *Entity) Name() string {
	if e.Display != "" {
		return e.Display
	}
	return e.Identity
}

// FieldList is a slice of fields.
type FieldList []Field

// Field represents a name/type pair for a field in an Object.
type Field struct {
	Declared string // The name of the field, does not affect the signature.
	Type     Type   // The type stored in the field.
}

// Type represents the common interface to all type objects in the schema.
type Type interface {
	String() string         // The true name of the type.
	Representation() string // The encoded representation of the type.
	EncodeValue(e Encoder, value interface{})
	DecodeValue(d Decoder) interface{}
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
