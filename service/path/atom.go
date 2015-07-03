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

// Atom is a path that refers to a single atom in an atom list.
type Atom struct {
	binary.Generate
	Atoms *Atoms // The path to the list of atoms.
	Index uint64 // The index of the atom in the array.
}

// Path implements the Path interface.
func (n *Atom) Path() string {
	return fmt.Sprintf("%v[%d]", n.Atoms, n.Index)
}

// Field implements the Value interface.
func (n *Atom) Field(name string) Value {
	return &Field{Struct: n, Name: name}
}

// ArrayIndex implements the Value interface.
func (n *Atom) ArrayIndex(index uint64) Value {
	return &ArrayIndex{Array: n, Index: index}
}

// MapIndex implements the Value interface.
func (n *Atom) MapIndex(key interface{}) Value {
	return &MapIndex{Map: n, Key: key}
}

// StateAfter returns the path to the state immediately following this atom.
func (n *Atom) StateAfter() *State {
	return &State{After: n}
}
