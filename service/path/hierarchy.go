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

// Hierarchy is a path that refers to a capture's hierarchy.
type Hierarchy struct {
	binary.Generate
	Capture *Capture // The path to the capture containing the hierarchy.
}

// String returns the string representation of the path.
func (n *Hierarchy) String() string { return n.Path() }

// Path implements the Path interface.
func (n *Hierarchy) Path() string {
	return fmt.Sprintf("%v.Hierarchy", n.Capture)
}

// Base implements the Path interface, returning the path to the hierarchy.
func (n *Hierarchy) Base() Path {
	return n.Capture
}

// Clone implements the Path interface, returning a deep-copy of this path.
func (n *Hierarchy) Clone() Path {
	return &Hierarchy{Capture: n.Capture.Clone().(*Capture)}
}

// Validate implements the Path interface.
func (n *Hierarchy) Validate() error {
	switch {
	case n == nil:
		return fmt.Errorf("Hierarchy is nil")
	case n.Capture == nil:
		return fmt.Errorf("Hierarchy.Capture is nil")
	}
	return n.Capture.Validate()
}
