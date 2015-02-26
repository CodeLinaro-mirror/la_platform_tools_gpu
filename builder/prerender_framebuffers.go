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
	atoms, _, err := getAtoms(request.Capture, db, logger)
	if err != nil {
		return err
	}

	atomIdsByContext := make(map[atom.ContextId][]atom.Id)
	for _, id := range request.AtomIds {
		if id < uint64(len(atoms)) {
			ctx := atoms[id].ContextId()
			atomIdsByContext[ctx] = append(atomIdsByContext[ctx], atom.Id(id))
		}
	}

	renderSettings := service.RenderSettings{
		MaxWidth:  request.Width,
		MaxHeight: request.Height,
		Wireframe: false,
	}

	var wg sync.WaitGroup
	for contextId, atomIds := range atomIdsByContext {
		var binaryIds []binary.ID
		for _, atomId := range atomIds {
			if imageInfoId, err := db.StoreRequest(&GetFramebufferColor{
				Capture:  request.Capture,
				Context:  contextId,
				Device:   request.Device,
				After:    atomId,
				Settings: renderSettings,
			}, logger); err == nil {
				var imageInfo service.ImageInfo
				if err = db.Load(imageInfoId, logger, &imageInfo); err == nil {
					binaryIds = append(binaryIds, imageInfo.Data.ID)
				}
			}
		}

		wg.Add(len(binaryIds))
		for _, binaryId := range binaryIds {
			go func(id binary.ID) {
				defer wg.Done()
				var dummy service.Binary
				db.Load(id, logger, &dummy)
			}(binaryId)
		}
		wg.Wait()
	}

	store.CopyResource(out, &service.Binary{})
	return nil
}
