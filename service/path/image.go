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

// ImageInfo is a path that refers to a image information.
type ImageInfo struct {
	binary.Generate
	ID binary.ID // The ImageInfo's unique identifier.
}

// String returns the string representation of the path.
func (i *ImageInfo) String() string { return i.Path() }

// Path implements the Path interface.
func (i *ImageInfo) Path() string {
	return fmt.Sprintf("ImageInfo(%v)", i.ID)
}

// Base implements the Path interface, returning nil as this is a root.
func (i *ImageInfo) Base() Path {
	return nil
}

// Clone implements the Path interface, returning a deep-copy of this path.
func (i *ImageInfo) Clone() Path {
	return &ImageInfo{ID: i.ID}
}

// Validate implements the Path interface.
func (i *ImageInfo) Validate() error {
	switch {
	case i == nil:
		return fmt.Errorf("ImageInfo is nil")
	case !i.ID.Valid():
		return fmt.Errorf("ImageInfo.ID is invalid")
	}
	return nil
}

// FindImageInfo returns the first ImageInfo found traversing the path p.
// If no ImageInfo was found, then nil is returned.
func FindImageInfo(p Path) *ImageInfo {
	for p != nil {
		if p, ok := p.(*ImageInfo); ok {
			return p
		}
		p = p.Base()
	}
	return nil
}
