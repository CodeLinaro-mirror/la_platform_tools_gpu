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
	"android.googlesource.com/platform/tools/gpu/database/store"
	"android.googlesource.com/platform/tools/gpu/gfxapi/state"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/service"
)

// build writes to out the MemoryInfo resource resulting from the given GetMemoryInfo request.
func (request *GetMemoryInfo) build(db database.Database, logger log.Logger, out binary.Object) error {
	capture, err := loadCapture(request.Capture, db, logger)
	if err != nil {
		return err
	}

	atoms, err := loadAtoms(capture.Atoms, db, logger)
	if err != nil {
		return err
	}

	if request.After >= atom.ID(len(atoms)) {
		return fmt.Errorf("After (%d) parameter is out of bounds. [0-%d]", request.After, len(atoms))
	}

	s := state.New()
	for _, a := range atoms[:request.After] {
		if err := s.Mutate(a); err != nil {
			return err
		}
	}

	// TODO: Stale, Unknown
	data, err := s.Memory.Slice(request.Range).Get(db, logger)
	if err != nil {
		return err
	}

	res := &service.MemoryInfo{Data: data}
	res.Current.Pack(memory.RangeList{request.Range})
	store.CopyResource(out, res)
	return nil
}
