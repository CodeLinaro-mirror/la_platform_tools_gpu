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

// ResourceID is the unique identifier of a single resource.
type ResourceID binary.ID

// Resource is a path that refers to a single resource snapshot at a given point
// in an atom stream.
type Resource struct {
	binary.Generate
	ID    ResourceID // The resource unique identifier.
	After *Atom      // The atom this resource snapshot immediately follows.
}

// String returns the string representation of the path.
func (n *Resource) String() string { return n.Path() }

// Path implements the Path interface.
func (n *Resource) Path() string {
	return fmt.Sprintf("%v.Resource<%v>", n.After, binary.ID(n.ID))
}

// Base implements the Path interface, returning the path to the atom the state
// is after.
func (n *Resource) Base() Path {
	return n.After
}

// Clone implements the Path interface, returning a deep-copy of this path.
func (n *Resource) Clone() Path {
	return &Resource{After: n.After.Clone().(*Atom)}
}

// Validate implements the Path interface.
func (n *Resource) Validate() error {
	switch {
	case n == nil:
		return fmt.Errorf("Resource is nil")
	case n.After == nil:
		return fmt.Errorf("Resource.After is nil")
	}
	return n.After.Validate()
}

// Thumbnail returns the path to the thumbnail of this resource.
func (n *Resource) Thumbnail(desiredWidth, desiredHeight uint32) *Thumbnail {
	return &Thumbnail{Object: n, DesiredWidth: desiredWidth, DesiredHeight: desiredHeight}
}
