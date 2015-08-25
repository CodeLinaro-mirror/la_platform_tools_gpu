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
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/image"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/service/path"
)

// BuildLazy returns the *service.ImageInfo resulting from the given
// GetFramebufferDepth request.
func (r *GetFramebufferDepth) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error) {
	fbWidth, fbHeight, err := getAtomFramebufferDimensions(r.After, d, l)
	if err != nil {
		return nil, fmt.Errorf("Failed to get framebuffer dimensions: %v\n", err)
	}

	data, err := database.Store(&RenderFramebufferDepth{
		Device: r.Device,
		After:  r.After,
	}, d, l)

	if err != nil {
		return nil, err
	}

	return &image.Info{
		Format: image.Float32(), // TODO: Add support for other formats.
		Width:  fbWidth,
		Height: fbHeight,
		Data:   &path.Blob{ID: data},
	}, nil
}

// BuildLazy returns the *service.Binary data for the given RenderFramebufferDepth
// request.
func (r *RenderFramebufferDepth) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error) {
	mgr := c.(*Context).ReplayManager

	ctx := replay.Context{
		Device:  r.Device.ID,
		Capture: r.After.Atoms.Capture.ID,
	}

	after, err := ResolveAtom(r.After, d, l)
	if err != nil {
		return nil, fmt.Errorf("Could not resolve atom '%v': %v", r.After, err)
	}

	apiID := after.API()
	api := gfxapi.Find(apiID)
	if api == nil {
		return nil, fmt.Errorf("Unknown graphics API '%v'", apiID)
	}

	query, ok := api.(replay.QueryDepthBuffer)
	if !ok {
		return nil, fmt.Errorf("The graphics API %s does not support reading depth buffers", api.Name())
	}

	img := <-query.QueryDepthBuffer(ctx, mgr, atom.ID(r.After.Index))
	if img.Error != nil {
		err := fmt.Errorf("Failed to retrieve framebuffer: %v", img.Error)
		log.Errorf(l, "%v", err)
		return nil, err
	}

	return img.Data, nil
}
