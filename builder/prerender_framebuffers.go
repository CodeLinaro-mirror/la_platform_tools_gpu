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
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service"
)

// BuildLazy renders and caches all the framebuffer color buffers in the
// GetFramebufferDepth request, returning an empty *service.Binary.
func (r *PrerenderFramebuffers) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error) {
	renderSettings := service.RenderSettings{
		MaxWidth:  r.Width,
		MaxHeight: r.Height,
		Wireframe: false,
	}

	var wg sync.WaitGroup
	for _, atomID := range r.AtomIDs {
		id, err := database.Store(&GetFramebufferColor{
			Capture:  r.Capture,
			Device:   r.Device,
			API:      r.API,
			After:    atom.ID(atomID),
			Settings: renderSettings,
		}, d, l)

		if err == nil {
			wg.Add(1)

			go func() {
				defer wg.Done()

				imageInfo, err := service.ResolveImageInfo(service.ImageInfoId{ID: id}, d, l)
				if err == nil {
					service.ResolveBinary(imageInfo.Data, d, l)
				}
			}()
		}
	}

	wg.Wait()

	return &service.Binary{}, nil
}
