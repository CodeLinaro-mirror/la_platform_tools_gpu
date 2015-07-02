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

	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/service"
)

// Build returns the *service.ImageInfo resulting from the given
// GetFramebufferDepth request.
func (r *GetFramebufferDepth) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error) {
	if !r.API.Valid() {
		return nil, fmt.Errorf("API must be valid")
	}

	fbWidth, fbHeight, err := getAtomFramebufferDimensions(r.Capture, r.After, d, l)
	if err != nil {
		return nil, err
	}

	data, err := database.Store(&RenderFramebufferDepth{
		Capture: r.Capture,
		Device:  r.Device,
		API:     r.API,
		After:   r.After,
	}, d, l)

	if err != nil {
		return nil, err
	}

	return &service.ImageInfo{
		Format: service.ImageFormatFloat32, // TODO: Add support for other formats.
		Width:  fbWidth,
		Height: fbHeight,
		Data:   service.BinaryId{ID: data},
	}, nil
}

// BuildLazy returns the *service.Binary data for the given RenderFramebufferDepth
// request.
func (r *RenderFramebufferDepth) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error) {
	mgr := c.(*Context).ReplayManager

	ctx := &replay.Context{
		DeviceID:  r.Device,
		CaptureID: r.Capture,
	}

	api := gfxapi.Find(gfxapi.ID(r.API.ID))
	if api == nil {
		return nil, fmt.Errorf("Unknown graphics API '%v'", r.API.ID)
	}

	query, ok := api.(replay.QueryDepthBuffer)
	if !ok {
		return nil, fmt.Errorf("The graphics API %s does not support reading depth buffers", api.Name())
	}

	img := <-query.QueryDepthBuffer(ctx, mgr, r.After)
	if img.Error != nil {
		log.Errorf(l, "%v", img.Error)
		return nil, img.Error
	}

	return &service.Binary{Data: img.Data}, nil
}
