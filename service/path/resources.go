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

// Resources is a path that refers to all the active resources in given state.
type Resources struct {
	binary.Generate
	Capture *Capture // The path to the Capture these resources belong to.
}

// String returns the string representation of the path.
func (n *Resources) String() string { return n.Path() }

// Path implements the Path interface.
func (n *Resources) Path() string {
	return fmt.Sprintf("%v.Capture", n.Capture)
}

// Base implements the Path interface, returning the path to the atom the state
// is after.
func (n *Resources) Base() Path {
	return n.Capture
}

// Clone implements the Path interface, returning a deep-copy of this path.
func (n *Resources) Clone() Path {
	return &Resources{Capture: n.Capture.Clone().(*Capture)}
}

// Validate implements the Path interface.
func (n *Resources) Validate() error {
	switch {
	case n == nil:
		return fmt.Errorf("Resources is nil")
	case n.Capture == nil:
		return fmt.Errorf("Resources.Capture is nil")
	}
	return n.Capture.Validate()
}
