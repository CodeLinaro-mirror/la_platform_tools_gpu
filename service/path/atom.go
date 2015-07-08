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

// String returns the string representation of the path.
func (n *Atom) String() string { return n.Path() }

// Path implements the Path interface.
func (n *Atom) Path() string {
	return fmt.Sprintf("%v[%d]", n.Atoms, n.Index)
}

// Base implements the Path interface, returning the path to the atom list.
func (n *Atom) Base() Path {
	return n.Atoms
}

// Clone implements the Path interface, returning a deep-copy of this path.
func (n *Atom) Clone() Path {
	return &Atom{Atoms: n.Atoms.Clone().(*Atoms), Index: n.Index}
}

// Field returns the path to the field value with the specified name on the
// atom represented by this path.
func (n *Atom) Field(name string) *Field {
	return &Field{Struct: n, Name: name}
}

// StateAfter returns the path to the state immediately following this atom.
func (n *Atom) StateAfter() *State {
	return &State{After: n}
}
