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
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/service"
)

// BuildLazy returns the *service.MemoryInfo resulting from the given
// GetMemoryInfo request.
func (r *GetMemoryInfo) BuildLazy(c interface{}, d database.Database, l log.Logger) (binary.Object, error) {
	capture, err := service.ResolveCapture(r.Capture, d, l)
	if err != nil {
		return nil, err
	}

	atoms, err := loadAtoms(capture.Atoms, d, l)
	if err != nil {
		return nil, err
	}

	if r.After >= atom.ID(len(atoms)) {
		return nil, fmt.Errorf("After (%d) parameter is out of bounds. [0-%d]", r.After, len(atoms))
	}

	s := gfxapi.NewState()
	for i, a := range atoms[:r.After] {
		if err := a.Mutate(s, d, l); err != nil {
			l.Warningf("Atom %d %v: %v", i, a, err)
		}
	}

	// TODO: Pool, Stale, Unknown
	data, err := s.Memory[memory.ApplicationPool].Slice(r.Range).Get(d, l)
	if err != nil {
		return nil, err
	}

	res := &service.MemoryInfo{Data: data}
	res.Current.Pack(memory.RangeList{r.Range})
	return res, nil
}
