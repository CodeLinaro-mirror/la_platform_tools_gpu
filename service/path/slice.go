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

// Slice is a path that refers to a subset of the elements in an array.
type Slice struct {
	binary.Generate
	Array Path // The path to the array.
	Start uint64
	End   uint64
}

// String returns the string representation of the path.
func (n *Slice) String() string { return n.Path() }

// Path implements the Path interface.
func (n *Slice) Path() string {
	return fmt.Sprintf("%v[%d:%d]", n.Array.Path(), n.Start, n.End)
}

// Base implements the Path interface, returning the path to the array.
func (n *Slice) Base() Path {
	return n.Array
}

// Clone implements the Path interface, returning a deep-copy of this path.
func (n *Slice) Clone() Path {
	return &Slice{Array: n.Array.Clone(), Start: n.Start, End: n.End}
}

// Validate implements the Path interface.
func (n *Slice) Validate() error {
	switch {
	case n == nil:
		return fmt.Errorf("Slice is nil")
	case n.Array == nil:
		return fmt.Errorf("Slice.Array is nil")
	case n.End < n.Start:
		return fmt.Errorf("Slice.End (%d) is before Slice.Start (%d)", n.End, n.Start)
	}
	return n.Array.Validate()
}

// Index returns the path to the i'th element in the slice.
func (n *Slice) Index(i uint64) Path {
	return &ArrayIndex{Array: n, Index: i}
}

// FindSlice returns the first Slice found traversing the path p.
// If no Slice was found, then nil is returned.
func FindSlice(p Path) *Slice {
	for p != nil {
		if p, ok := p.(*Slice); ok {
			return p
		}
		p = p.Base()
	}
	return nil
}
