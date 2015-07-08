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
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/service"
)

// BuildLazy returns the *service.ImageInfo resulting from the given
// GetFramebufferColor request.
func (r *GetFramebufferColor) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error) {
	fbWidth, fbHeight, err := getAtomFramebufferDimensions(r.After, d, l)
	if err != nil {
		return nil, err
	}
	imgWidth, imgHeight := uniformScale(fbWidth, fbHeight, r.Settings.MaxWidth, r.Settings.MaxHeight)

	data, err := database.Store(&RenderFramebufferColor{
		Device:    r.Device,
		After:     r.After,
		Width:     imgWidth,
		Height:    imgHeight,
		Wireframe: r.Settings.Wireframe,
	}, d, l)

	if err != nil {
		return nil, err
	}

	return &service.ImageInfo{
		Format: service.ImageFormatRGBA8, // TODO: Add support for other formats.
		Width:  imgWidth,
		Height: imgHeight,
		Data:   service.BinaryId{ID: data},
	}, nil
}

// BuildLazy returns the *service.Binary data for the given RenderFramebufferColor
// request.
func (r *RenderFramebufferColor) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error) {
	mgr := c.(*Context).ReplayManager

	ctx := &replay.Context{
		DeviceID:  service.DeviceId{ID: r.Device.ID},
		CaptureID: service.CaptureId{ID: r.After.Atoms.Capture.ID},
	}

	after, err := ResolveAtom(r.After, d, l)
	if err != nil {
		return nil, err
	}

	apiID := after.API()
	api := gfxapi.Find(apiID)
	if api == nil {
		return nil, fmt.Errorf("Unknown graphics API '%v'", apiID)
	}

	query, ok := api.(replay.QueryColorBuffer)
	if !ok {
		return nil, fmt.Errorf("The graphics API %s does not support reading color buffers", api.Name())
	}

	img := <-query.QueryColorBuffer(ctx, mgr, atom.ID(r.After.Index), r.Width, r.Height, r.Wireframe)
	if img.Error != nil {
		log.Errorf(l, "%v", img.Error)
		return nil, img.Error
	}

	return img.Data, nil
}
