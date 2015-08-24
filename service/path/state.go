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

// State is a path that refers to the driver state immediately after an atom.
type State struct {
	binary.Generate
	After *Atom // The path to the atom the state immediately follows.
}

// String returns the string representation of the path.
func (n *State) String() string { return n.Path() }

// Path implements the Path interface.
func (n *State) Path() string {
	return fmt.Sprintf("%v.State", n.After)
}

// Base implements the Path interface, returning the path to the atom the state
// is after.
func (n *State) Base() Path {
	return n.After
}

// Clone implements the Path interface, returning a deep-copy of this path.
func (n *State) Clone() Path {
	return &State{After: n.After.Clone().(*Atom)}
}

// Validate implements the Path interface.
func (n *State) Validate() error {
	switch {
	case n == nil:
		return fmt.Errorf("State is nil")
	case n.After == nil:
		return fmt.Errorf("State.After is nil")
	}
	return n.After.Validate()
}

// Field returns the path to the field value with the specified name on the
// struct object represented by this path.
// The represented value type must be of type struct, otherwise the returned
// path is invalid.
func (n *State) Field(name string) *Field {
	return &Field{Struct: n, Name: name}
}

// Slice returns the path to the sliced subset of this array or slice
// represented by this path.
// The represented value type must be of type array or slice, otherwise the
// returned path is invalid.
func (n *State) Slice(start, end uint64) *Slice {
	return &Slice{Array: n, Start: start, End: end}
}

// ArrayIndex returns the path to the i'th element on the array or slice
// represented by this path.
// The represented value type must be of type array or slice, otherwise the
// returned path is invalid.
func (n *State) ArrayIndex(index uint64) *ArrayIndex {
	return &ArrayIndex{Array: n, Index: index}
}

// As returns the path to the state object converted to the requested type.
// If the represented value does not support converting to the requested type
// then the returned path is invalid.
func (n *State) As(ty interface{}) Value {
	return &As{Object: n, Type: ty}
}

// MapIndex returns the path to the map element with key k on the map object
// represented by this path.
// The represented value type must be of type map, otherwise the returned path
// is invalid.
func (n *State) MapIndex(key interface{}) *MapIndex {
	return &MapIndex{Map: n, Key: key}
}

// FindState returns the first State found traversing the path p.
// If no State was found, then nil is returned.
func FindState(p Path) *State {
	for p != nil {
		if p, ok := p.(*State); ok {
			return p
		}
		p = p.Base()
	}
	return nil
}
