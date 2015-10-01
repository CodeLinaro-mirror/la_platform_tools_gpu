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

// Thumbnail is a path to the thumbnail for a given object.
type Thumbnail struct {
	binary.Generate
	// The path to the thumbnail's object
	Object Path
	// The desired maximum width of the thumbnail image.
	// If DesiredMaxWidth <= 0, then no limits will be placed on the width.
	DesiredMaxWidth uint32
	// The desired maximum height of the thumbnail image.
	// If DesiredMaxHeight <= 0, then no limits will be placed on the height.
	DesiredMaxHeight uint32
	// If requested thumbnail format. If nil, then return the native format.
	DesiredFormat interface{}
}

// String returns the string representation of the path.
func (n *Thumbnail) String() string { return n.Path() }

// Path implements the Path interface.
func (n *Thumbnail) Path() string {
	return fmt.Sprintf("%v.Thumbnail<%v, %d x %d>", n.Object,
		n.DesiredFormat, n.DesiredMaxWidth, n.DesiredMaxHeight)
}

// Base implements the Path interface, returning the path to the atom the state
// is after.
func (n *Thumbnail) Base() Path {
	return n.Object
}

// Clone implements the Path interface, returning a deep-copy of this path.
func (n *Thumbnail) Clone() Path {
	return &Thumbnail{
		Object:           n.Object.Clone(),
		DesiredMaxWidth:  n.DesiredMaxWidth,
		DesiredMaxHeight: n.DesiredMaxHeight,
		DesiredFormat:    n.DesiredFormat,
	}
}

// Validate implements the Path interface.
func (n *Thumbnail) Validate() error {
	switch {
	case n == nil:
		return fmt.Errorf("Thumbnail is nil")
	case n.Object == nil:
		return fmt.Errorf("Thumbnail.Resources is nil")
	}
	return n.Object.Validate()
}

// As returns the path to the thumbnail converted to the requested format.
// If the represented value does not support converting to the requested type
// then the returned path is invalid.
func (n *Thumbnail) As(ty interface{}) Value {
	return &As{Object: n, Type: ty}
}
