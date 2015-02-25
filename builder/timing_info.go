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
)

// build computes and writes the output of the given GetTimingInfo request to the given out.
func (request *GetTimingInfo) build(mgr *replay.Manager, db database.Database, logger log.Logger, out binary.Object) error {
	ctx := &replay.Context{
		DeviceID:  request.Device,
		CaptureID: request.Capture,
		ContextID: request.Context,
	}

	api, err := getAPI(request.Capture, db, logger)
	if err != nil {
		return err
	}

	timings := <-api.TimeCalls(ctx, mgr, request.TimingMask)
	if timings.Error != nil {
		logger.Error("%v", timings.Error)
		return timings.Error
	}

	store.CopyResource(out, &timings.TimingInfo)
	return nil
}
