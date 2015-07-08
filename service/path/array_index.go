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

// ArrayIndex is a path that refers to a single element of an array.
type ArrayIndex struct {
	binary.Generate
	Array Path   // The path to the array.
	Index uint64 // The index of the element in the array.
}

// String returns the string representation of the path.
func (n *ArrayIndex) String() string { return n.Path() }

// Path implements the Path interface.
func (n *ArrayIndex) Path() string {
	return fmt.Sprintf("%v[%d]", n.Array, n.Index)
}

// Base implements the Path interface, returning the path to the array.
func (n *ArrayIndex) Base() Path {
	return n.Array
}

// Clone implements the Path interface, returning a deep-copy of this path.
func (n *ArrayIndex) Clone() Path {
	return &ArrayIndex{Array: n.Array.Clone(), Index: n.Index}
}

// Validate implements the Path interface.
func (n *ArrayIndex) Validate() error {
	switch {
	case n == nil:
		return fmt.Errorf("ArrayIndex is nil")
	case n.Array == nil:
		return fmt.Errorf("ArrayIndex.Array is nil")
	}
	return n.Array.Validate()
}

// Field returns the path to the field value with the specified name on the
// struct object represented by this path.
// The represented value type must be of type struct, otherwise the returned
// path is invalid.
func (n *ArrayIndex) Field(name string) *Field {
	return &Field{Struct: n, Name: name}
}

// Slice returns the path to the sliced subset of this array or slice
// represented by this path.
// The represented value type must be of type array or slice, otherwise the
// returned path is invalid.
func (n *ArrayIndex) Slice(start, end uint64) *Slice {
	return &Slice{Array: n, Start: start, End: end}
}

// ArrayIndex returns the path to the i'th element on the array or slice
// represented by this path.
// The represented value type must be of type array or slice, otherwise the
// returned path is invalid.
func (n *ArrayIndex) ArrayIndex(index uint64) *ArrayIndex {
	return &ArrayIndex{Array: n, Index: index}
}

// MapIndex returns the path to the map element with key k on the map object
// represented by this path.
// The represented value type must be of type map, otherwise the returned path
// is invalid.
func (n *ArrayIndex) MapIndex(key interface{}) *MapIndex {
	return &MapIndex{Map: n, Key: key}
}

// FindArrayIndex returns the first ArrayIndex found traversing the path p.
// If no Atom was found, then nil is returned.
func FindArrayIndex(p Path) *ArrayIndex {
	for p != nil {
		if p, ok := p.(*ArrayIndex); ok {
			return p
		}
		p = p.Base()
	}
	return nil
}
