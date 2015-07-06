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
	Array Value  // The path to the array.
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
	return &ArrayIndex{Array: n.Array.Clone().(Value), Index: n.Index}
}

// Field implements the Value interface.
func (n *ArrayIndex) Field(name string) Value {
	return &Field{Struct: n, Name: name}
}

// ArrayIndex implements the Value interface.
func (n *ArrayIndex) ArrayIndex(index uint64) Value {
	return &ArrayIndex{Array: n, Index: index}
}

// MapIndex implements the Value interface.
func (n *ArrayIndex) MapIndex(key interface{}) Value {
	return &MapIndex{Map: n, Key: key}
}
