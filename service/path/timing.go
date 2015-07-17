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

// TimingInfo is a path that refers to timing information.
type TimingInfo struct {
	binary.Generate
	ID binary.ID // The TimingInfo's unique identifier.
}

// String returns the string representation of the path.
func (t *TimingInfo) String() string { return t.Path() }

// Path implements the Path interface.
func (t *TimingInfo) Path() string {
	return fmt.Sprintf("TimingInfo(%v)", t.ID)
}

// Base implements the Path interface, returning nil as this is a root.
func (t *TimingInfo) Base() Path {
	return nil
}

// Clone implements the Path interface, returning a deep-copy of this path.
func (t *TimingInfo) Clone() Path {
	return &TimingInfo{ID: t.ID}
}

// Validate implements the Path interface.
func (t *TimingInfo) Validate() error {
	switch {
	case t == nil:
		return fmt.Errorf("TimingInfo is nil")
	case !t.ID.Valid():
		return fmt.Errorf("TimingInfo.ID is invalid")
	}
	return nil
}

// FindTimingInfo returns the first TimingInfo found traversing the path p.
// If no TimingInfo was found, then nil is returned.
func FindTimingInfo(p Path) *TimingInfo {
	for p != nil {
		if p, ok := p.(*TimingInfo); ok {
			return p
		}
		p = p.Base()
	}
	return nil
}
