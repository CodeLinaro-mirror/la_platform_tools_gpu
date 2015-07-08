// Copyright (C) 2015 The Android Open Source Project
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

package path

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/binary"
)

// Field is a path that refers to a single field of a struct object.
type Field struct {
	binary.Generate
	Struct Path   // The path to the structure holding the field.
	Name   string // The name of the field.
}

// String returns the string representation of the path.
func (n *Field) String() string { return n.Path() }

// Path implements the Path interface.
func (n *Field) Path() string {
	return fmt.Sprintf("%v.%s", n.Struct, n.Name)
}

// Base implements the Path interface, returning the path to the struct.
func (n *Field) Base() Path {
	return n.Struct
}

// Clone implements the Path interface, returning a deep-copy of this path.
func (n *Field) Clone() Path {
	return &Field{Struct: n.Struct.Clone(), Name: n.Name}
}

// Validate implements the Path interface.
func (n *Field) Validate() error {
	switch {
	case n == nil:
		return fmt.Errorf("Field is nil")
	case n.Struct == nil:
		return fmt.Errorf("Field.Struct is nil")
	case n.Name == "":
		return fmt.Errorf("Field.Name is empty")
	}
	return n.Struct.Validate()
}

// Field returns the path to the field value with the specified name on the
// struct object represented by this path.
// The represented value type must be of type struct, otherwise the returned
// path is invalid.
func (n *Field) Field(name string) *Field {
	return &Field{Struct: n, Name: name}
}

// Slice returns the path to the sliced subset of this array or slice
// represented by this path.
// The represented value type must be of type array or slice, otherwise the
// returned path is invalid.
func (n *Field) Slice(start, end uint64) *Slice {
	return &Slice{Array: n, Start: start, End: end}
}

// ArrayIndex returns the path to the i'th element on the array or slice
// represented by this path.
// The represented value type must be of type array or slice, otherwise the
// returned path is invalid.
func (n *Field) ArrayIndex(index uint64) *ArrayIndex {
	return &ArrayIndex{Array: n, Index: index}
}

// MapIndex returns the path to the map element with key k on the map object
// represented by this path.
// The represented value type must be of type map, otherwise the returned path
// is invalid.
func (n *Field) MapIndex(key interface{}) *MapIndex {
	return &MapIndex{Map: n, Key: key}
}

// FindField returns the first Field found traversing the path p.
// If no Atom was found, then nil is returned.
func FindField(p Path) *Field {
	for p != nil {
		if p, ok := p.(*Field); ok {
			return p
		}
		p = p.Base()
	}
	return nil
}
