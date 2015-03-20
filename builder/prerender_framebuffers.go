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
	"sync"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/database/store"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service"
)

// build writes to out an empty Binary resource after processing the given PrerenderFramebuffers request.
func (request *PrerenderFramebuffers) build(db database.Database, logger log.Logger, out binary.Object) error {
	capture, err := loadCapture(request.Capture, db, logger)
	if err != nil {
		return err
	}

	atoms, err := loadAtoms(capture.Atoms, db, logger)
	if err != nil {
		return err
	}

	atomIDsByContext := make(map[atom.ContextID][]atom.ID)
	for _, id := range request.AtomIDs {
		if id < uint64(len(atoms)) {
			ctx := atoms[id].ContextID()
			atomIDsByContext[ctx] = append(atomIDsByContext[ctx], atom.ID(id))
		}
	}

	renderSettings := service.RenderSettings{
		MaxWidth:  request.Width,
		MaxHeight: request.Height,
		Wireframe: false,
	}

	var wg sync.WaitGroup
	for contextID, atomIDs := range atomIDsByContext {
		var binaryIDs []binary.ID
		for _, atomID := range atomIDs {
			if imageInfoID, err := db.StoreRequest(&GetFramebufferColor{
				Capture:  request.Capture,
				Context:  contextID,
				Device:   request.Device,
				After:    atomID,
				Settings: renderSettings,
			}, logger); err == nil {
				var imageInfo service.ImageInfo
				if err = db.Load(imageInfoID, logger, &imageInfo); err == nil {
					binaryIDs = append(binaryIDs, imageInfo.Data.ID)
				}
			}
		}

		wg.Add(len(binaryIDs))
		for _, binaryID := range binaryIDs {
			go func(id binary.ID) {
				defer wg.Done()
				var dummy service.Binary
				db.Load(id, logger, &dummy)
			}(binaryID)
		}
		wg.Wait()
	}

	store.CopyResource(out, &service.Binary{})
	return nil
}
