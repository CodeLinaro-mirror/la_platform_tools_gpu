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

// As refers to a struct, array, slice, map or POD type converted to ty.
type As struct {
	binary.Generate
	Object Path        // The path to the unconverted object.
	Type   interface{} // The requested type of the converted object.
}

// String returns the string representation of the path.
func (n *As) String() string { return n.Path() }

// Path implements the Path interface.
func (n *As) Path() string {
	return fmt.Sprintf("%v.As<%s>", n.Object, n.Type)
}

// Base implements the Path interface, returning the path to the struct.
func (n *As) Base() Path {
	return n.Object
}

// Clone implements the Path interface, returning a deep-copy of this path.
func (n *As) Clone() Path {
	return &As{Object: n.Object.Clone(), Type: n.Type}
}

// Validate implements the Path interface.
func (n *As) Validate() error {
	switch {
	case n == nil:
		return fmt.Errorf("Field is nil")
	case n.Object == nil:
		return fmt.Errorf("Field.Object is nil")
	case n.Type == nil:
		return fmt.Errorf("Field.Type is nil")
	}
	return n.Object.Validate()
}

// Field returns the path to the field value with the specified name on the
// struct object represented by this path.
// The represented value type must be of type struct, otherwise the returned
// path is invalid.
func (n *As) Field(name string) *Field {
	return &Field{Struct: n, Name: name}
}

// Slice returns the path to the sliced subset of this array or slice
// represented by this path.
// The represented value type must be of type array or slice, otherwise the
// returned path is invalid.
func (n *As) Slice(start, end uint64) *Slice {
	return &Slice{Array: n, Start: start, End: end}
}

// ArrayIndex returns the path to the i'th element on the array or slice
// represented by this path.
// The represented value type must be of type array or slice, otherwise the
// returned path is invalid.
func (n *As) ArrayIndex(index uint64) *ArrayIndex {
	return &ArrayIndex{Array: n, Index: index}
}

// MapIndex returns the path to the map element with key k on the map object
// represented by this path.
// The represented value type must be of type map, otherwise the returned path
// is invalid.
func (n *As) MapIndex(key interface{}) *MapIndex {
	return &MapIndex{Map: n, Key: key}
}

// As returns the path to the struct, array, slice, map or POD type, converted
// to the requested type.
// If the represented value does not support converting to the requested type
// then the returned path is invalid.
func (n *As) As(ty interface{}) Value {
	return &As{Object: n, Type: ty}
}
