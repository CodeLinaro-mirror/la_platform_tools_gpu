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
	Struct Value  // The path to the structure holding the field.
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
	return &Field{Struct: n.Struct.Clone().(Value), Name: n.Name}
}

// Field implements the Value interface.
func (n *Field) Field(name string) Value {
	return &Field{Struct: n, Name: name}
}

// ArrayIndex implements the Value interface.
func (n *Field) ArrayIndex(index uint64) Value {
	return &ArrayIndex{Array: n, Index: index}
}

// MapIndex implements the Value interface.
func (n *Field) MapIndex(key interface{}) Value {
	return &MapIndex{Map: n, Key: key}
}
