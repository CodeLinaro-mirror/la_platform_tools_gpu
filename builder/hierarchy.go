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

package builder

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service"
)

// BuildLazy returns the *service.Hierarchy resulting from the given
// GetHierarchy request.
func (r *GetHierarchy) BuildLazy(c interface{}, d database.Database, l log.Logger) (binary.Object, error) {
	capture, err := service.ResolveCapture(r.Capture, d, l)
	if err != nil {
		return nil, err
	}

	atoms, err := loadAtoms(capture.Atoms, d, l)
	if err != nil {
		return nil, err
	}

	root := atom.Group{
		Range: atom.Range{Start: 0, End: atom.ID(len(atoms))},
	}
	var frameIndex, drawIndex int
	var frameStartID, drawStartID atom.ID
	for i, a := range atoms {
		endID := atom.ID(i + 1) // Increment by one, since atom.Range's end is non-inclusive.
		if a.Flags().IsEndOfFrame() {
			root.SubGroups.Add(frameStartID, endID, fmt.Sprintf("Frame %d", frameIndex))
			frameStartID = endID
			frameIndex++
			drawStartID = endID
			drawIndex = 0 // Reset the draw index, it is relative to the new frame index.
		}
		if a.Flags().IsDrawCall() {
			root.SubGroups.Add(drawStartID, endID, fmt.Sprintf("Draw %d", drawIndex))
			drawStartID = endID
			drawIndex++
		}
	}

	return &service.Hierarchy{
		Root: root,
	}, nil
}
