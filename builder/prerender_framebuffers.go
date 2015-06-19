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
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service"
)

// build writes to out an empty Binary resource after processing the given PrerenderFramebuffers request.
func (request *PrerenderFramebuffers) build(db database.Database, logger log.Logger, out binary.Object) error {
	renderSettings := service.RenderSettings{
		MaxWidth:  request.Width,
		MaxHeight: request.Height,
		Wireframe: false,
	}

	var wg sync.WaitGroup
	for _, atomID := range request.AtomIDs {
		imageInfoID, err := db.StoreRequest(&GetFramebufferColor{
			Capture:  request.Capture,
			Device:   request.Device,
			API:      request.API,
			After:    atom.ID(atomID),
			Settings: renderSettings,
		}, logger)

		if err == nil {
			wg.Add(1)

			go func() {
				defer wg.Done()

				var imageInfo service.ImageInfo
				if err := db.Load(imageInfoID, logger, &imageInfo); err == nil {
					var dummy service.Binary
					db.Load(imageInfo.Data.ID, logger, &dummy)
				}
			}()
		}
	}

	wg.Wait()

	database.CopyResource(out, &service.Binary{})
	return nil
}
