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

package schema

import (
	"strings"

	"android.googlesource.com/platform/tools/gpu/binary"
)

// Class represents an encodable object type with a type ID.
type Class struct {
	TypeID   binary.ID       // The unique type identifier for the Object.
	Package  string          // The package that declared the struct.
	Name     string          // The simple name of the Object.
	Exported bool            // Whether the class is exported from it's package
	Fields   FieldList       // Descriptions of the fields of the Object.
	Metadata []binary.Object // The metadata for the class.
}

// Field represents a name/type pair for a field in an Object.
type Field struct {
	Declared string // The name of the field.
	Type     Type   // The type stored in the field.
}

// FieldList is a slice of fields.
type FieldList []Field

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

func (c *Class) ID() binary.ID {
	return c.TypeID
}

func (c *Class) New() binary.Object { return &Object{Type: c} }

func (c *Class) Encode(e binary.Encoder, object binary.Object) {
	o := object.(*Object)
	for i, f := range c.Fields {
		f.Type.Encode(e, o.Fields[i])
	}
}

func (c *Class) doDecode(d binary.Decoder, o *Object) {
	o.Fields = make([]interface{}, len(c.Fields))
	for i, f := range c.Fields {
		o.Fields[i] = f.Type.Decode(d)
	}
}

func (c *Class) Decode(d binary.Decoder) binary.Object {
	o := &Object{Type: c}
	c.doDecode(d, o)
	return o
}

func (c *Class) DecodeTo(d binary.Decoder, object binary.Object) {
	c.doDecode(d, object.(*Object))
}
