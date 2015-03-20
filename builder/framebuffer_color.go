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
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/database/store"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/service"
)

// build writes to out the ImageInfo resource resulting from the given GetFramebufferColor request.
func (request *GetFramebufferColor) build(db database.Database, logger log.Logger, out binary.Object) error {
	fbWidth, fbHeight, err := getAtomFramebufferDimensions(request.Capture, request.Context, request.After, db, logger)
	if err != nil {
		return err
	}
	imgWidth, imgHeight := uniformScale(fbWidth, fbHeight, request.Settings.MaxWidth, request.Settings.MaxHeight)

	data, err := db.StoreRequest(&RenderFramebufferColor{
		Capture:   request.Capture,
		Context:   request.Context,
		Device:    request.Device,
		After:     request.After,
		Width:     imgWidth,
		Height:    imgHeight,
		Wireframe: request.Settings.Wireframe,
	}, logger)

	if err != nil {
		return err
	}

	store.CopyResource(out, &service.ImageInfo{
		Format: service.ImageFormatRGBA8, // TODO: Add support for other formats.
		Width:  imgWidth,
		Height: imgHeight,
		Data:   service.BinaryId{data},
	})
	return nil
}

// build computes and writes the output of the given RenderFramebufferColor request to the given out.
func (request *RenderFramebufferColor) build(mgr *replay.Manager, db database.Database, logger log.Logger, out binary.Object) error {
	ctx := &replay.Context{
		DeviceID:  request.Device,
		CaptureID: request.Capture,
		ContextID: request.Context,
	}

	capture, err := loadCapture(request.Capture, db, logger)
	if err != nil {
		return err
	}

	api, err := getAPI(capture.Contexts, request.Context)
	if err != nil {
		return err
	}

	img := <-api.ColorBuffer(ctx, mgr, request.After, request.Width, request.Height, request.Wireframe)
	if img.Error != nil {
		logger.Error("%v", img.Error)
		return img.Error
	}

	store.CopyResource(out, &service.Binary{Data: img.Data})
	return nil
}
