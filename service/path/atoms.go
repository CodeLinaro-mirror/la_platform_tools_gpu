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

// Atoms is a path that refers to the full list of atoms in a capture.
type Atoms struct {
	binary.Generate
	Capture *Capture // The path to the capture containing the atoms.
}

// String returns the string representation of the path.
func (n *Atoms) String() string { return n.Path() }

// Path implements the Path interface.
func (n *Atoms) Path() string {
	return fmt.Sprintf("%v.Atoms", n.Capture)
}

// Base implements the Path interface, returning the path to the atoms.
func (n *Atoms) Base() Path {
	return n.Capture
}

// Clone implements the Path interface, returning a deep-copy of this path.
func (n *Atoms) Clone() Path {
	return &Atoms{Capture: n.Capture.Clone().(*Capture)}
}

// Validate implements the Path interface.
func (n *Atoms) Validate() error {
	switch {
	case n == nil:
		return fmt.Errorf("Atoms is nil")
	case n.Capture == nil:
		return fmt.Errorf("Atoms.Capture is nil")
	}
	return n.Capture.Validate()
}

// Index returns the path to the i'th atom in the atom list.
func (n *Atoms) Index(i uint64) *Atom {
	return &Atom{Atoms: n, Index: i}
}

// Slice returns the path to the sliced subset of the atom list.
func (n *Atoms) Slice(start, end uint64) *Slice {
	return &Slice{Array: n, Start: start, End: end}
}
