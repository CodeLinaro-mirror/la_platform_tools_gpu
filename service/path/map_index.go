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

// MapIndex is a path that refers to a single value in a map.
type MapIndex struct {
	binary.Generate
	Map Path        // The path to the map containing the value.
	Key interface{} // The key to the value in the map.
}

// String returns the string representation of the path.
func (n *MapIndex) String() string { return n.Path() }

// Path implements the Path interface.
func (n *MapIndex) Path() string {
	return fmt.Sprintf("%v[%v]", n.Map, n.Key)
}

// Base implements the Path interface, returning the path to the map.
func (n *MapIndex) Base() Path {
	return n.Map
}

// Clone implements the Path interface, returning a deep-copy of this path.
func (n *MapIndex) Clone() Path {
	return &MapIndex{Map: n.Map.Clone(), Key: n.Key}
}

// Validate implements the Path interface.
func (n *MapIndex) Validate() error {
	switch {
	case n == nil:
		return fmt.Errorf("MapIndex is nil")
	case n.Map == nil:
		return fmt.Errorf("MapIndex.Map is nil")
	case n.Key == nil:
		return fmt.Errorf("MapIndex.Key is nil")
	}
	return n.Map.Validate()
}

// Field returns the path to the field value with the specified name on the
// struct object represented by this path.
// The represented value type must be of type struct, otherwise the returned
// path is invalid.
func (n *MapIndex) Field(name string) *Field {
	return &Field{Struct: n, Name: name}
}

// Slice returns the path to the sliced subset of this array or slice
// represented by this path.
// The represented value type must be of type array or slice, otherwise the
// returned path is invalid.
func (n *MapIndex) Slice(start, end uint64) *Slice {
	return &Slice{Array: n, Start: start, End: end}
}

// ArrayIndex returns the path to the i'th element on the array or slice
// represented by this path.
// The represented value type must be of type array or slice, otherwise the
// returned path is invalid.
func (n *MapIndex) ArrayIndex(index uint64) *ArrayIndex {
	return &ArrayIndex{Array: n, Index: index}
}

// MapIndex returns the path to the map element with key k on the map object
// represented by this path.
// The represented value type must be of type map, otherwise the returned path
// is invalid.
func (n *MapIndex) MapIndex(key interface{}) *MapIndex {
	return &MapIndex{Map: n, Key: key}
}
