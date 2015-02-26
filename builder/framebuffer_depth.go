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

// build writes to out the ImageInfo resource resulting from the given GetFramebufferDepth request.
func (request *GetFramebufferDepth) build(db database.Database, logger log.Logger, out binary.Object) error {
	fbWidth, fbHeight, err := getAtomFramebufferDimensions(request.Capture, request.Context, request.After, db, logger)
	if err != nil {
		return err
	}

	data, err := db.StoreRequest(&RenderFramebufferDepth{
		Capture: request.Capture,
		Context: request.Context,
		Device:  request.Device,
		After:   request.After,
	}, logger)

	if err != nil {
		return err
	}

	store.CopyResource(out, &service.ImageInfo{
		Format: service.ImageFormatFloat32, // TODO: Add support for other formats.
		Width:  fbWidth,
		Height: fbHeight,
		Data:   service.BinaryId{data},
	})
	return nil
}

// build computes and writes the output of the given RenderFramebufferDepth request to the given out.
func (request *RenderFramebufferDepth) build(mgr *replay.Manager, db database.Database, logger log.Logger, out binary.Object) error {
	ctx := &replay.Context{
		DeviceID:  request.Device,
		CaptureID: request.Capture,
		ContextID: request.Context,
	}

	api, err := getAPI(request.Capture, db, logger)
	if err != nil {
		return err
	}

	img := <-api.DepthBuffer(ctx, mgr, request.After)
	if img.Error != nil {
		logger.Error("%v", img.Error)
		return img.Error
	}

	store.CopyResource(out, &service.Binary{Data: img.Data})
	return nil
}
