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

// MemoryRange is a path that refers to a range of memory for a specific pool at
// a specific point in the capture.
type MemoryRange struct {
	binary.Generate
	After *Atom // The path to the atom the memory snapshot immediately follows.

	Pool    uint64 // The pool identifier.
	Address uint64 // The memory base address.
	Size    uint64 // The size in bytes of this memory range.
}

// String returns the string representation of the path.
func (n *MemoryRange) String() string { return n.Path() }

// Path implements the Path interface.
func (n *MemoryRange) Path() string {
	return fmt.Sprintf("%v.MemoryRange<%v>[0x%x:0x%x]",
		n.After, n.Pool, n.Address, n.Address+n.Size-1)
}

// Base implements the Path interface, returning the path to the atom list.
func (n *MemoryRange) Base() Path {
	return n.After
}

// Clone implements the Path interface, returning a deep-copy of this path.
func (n *MemoryRange) Clone() Path {
	return &MemoryRange{
		After:   n.After.Clone().(*Atom),
		Pool:    n.Pool,
		Address: n.Address,
		Size:    n.Size,
	}
}

// Validate implements the Path interface.
func (n *MemoryRange) Validate() error {
	switch {
	case n == nil:
		return fmt.Errorf("MemoryRange is nil")
	case n.After == nil:
		return fmt.Errorf("MemoryRange.After is nil")
	}
	return n.After.Validate()
}

// FindMemoryRange returns the first MemoryRange found traversing the path p.
// If no MemoryRange was found, then nil is returned.
func FindMemoryRange(p Path) *MemoryRange {
	for p != nil {
		if p, ok := p.(*MemoryRange); ok {
			return p
		}
		p = p.Base()
	}
	return nil
}
