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

// Blob is a path that refers to a byte array.
type Blob struct {
	binary.Generate
	ID binary.ID // The blob's unique identifier.
}

// String returns the string representation of the path.
func (b *Blob) String() string { return b.Path() }

// Path implements the Path interface.
func (b *Blob) Path() string {
	return fmt.Sprintf("Blob(%v)", b.ID)
}

// Base implements the Path interface, returning nil as this is a root.
func (b *Blob) Base() Path {
	return nil
}

// Clone implements the Path interface, returning a deep-copy of this path.
func (b *Blob) Clone() Path {
	return &Blob{ID: b.ID}
}

// Validate implements the Path interface.
func (b *Blob) Validate() error {
	switch {
	case b == nil:
		return fmt.Errorf("Blob is nil")
	case !b.ID.Valid():
		return fmt.Errorf("Blob.ID is invalid")
	}
	return nil
}

// FindBlob returns the first Blob found traversing the path p.
// If no Blob was found, then nil is returned.
func FindBlob(p Path) *Blob {
	for p != nil {
		if p, ok := p.(*Blob); ok {
			return p
		}
		p = p.Base()
	}
	return nil
}
