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
	"android.googlesource.com/platform/tools/gpu/service"
	"android.googlesource.com/platform/tools/gpu/service/path"
)

// BuildLazy returns the *service.ImageInfo resulting from the given
// GetFramebufferColor request.
func (r *GetFramebufferColor) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error) {
	fbWidth, fbHeight, err := getAtomFramebufferDimensions(r.After, d, l)
	if err != nil {
		return nil, fmt.Errorf("Failed to get framebuffer dimensions: %v\n", err)
	}
	imgWidth, imgHeight := uniformScale(fbWidth, fbHeight, r.Settings.MaxWidth, r.Settings.MaxHeight)

	data, err := database.Store(&RenderFramebufferColor{
		Device:        r.Device,
		After:         r.After,
		Width:         imgWidth,
		Height:        imgHeight,
		WireframeMode: r.Settings.WireframeMode,
	}, d, l)

	if err != nil {
		return nil, err
	}

	return &image.Info{
		Format: image.RGBA(), // TODO: Add support for other formats.
		Width:  int(imgWidth),
		Height: int(imgHeight),
		Data:   &path.Blob{ID: data},
	}, nil
}

// BuildLazy returns the *service.Binary data for the given RenderFramebufferColor
// request.
func (r *RenderFramebufferColor) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error) {
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

	query, ok := api.(replay.QueryColorBuffer)
	if !ok {
		return nil, fmt.Errorf("The graphics API %s does not support reading color buffers", api.Name())
	}

	wireframeMode := replay.NoWireframe
	switch r.WireframeMode {
	case service.NoWireframe:
	case service.AllWireframe:
		wireframeMode = replay.AllWireframe
	case service.WireframeOverlay:
		wireframeMode = replay.WireframeOverlay
	default:
		return nil, fmt.Errorf("Unknown wireframe mode %v", r.WireframeMode)
	}

	img := <-query.QueryColorBuffer(ctx, mgr, atom.ID(r.After.Index), r.Width, r.Height, wireframeMode)
	if img.Error != nil {
		err := fmt.Errorf("Failed to retrieve framebuffer: %v", img.Error)
		log.Errorf(l, "%v", err)
		return nil, err
	}

	return img.Data, nil
}
