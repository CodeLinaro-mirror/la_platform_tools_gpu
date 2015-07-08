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
	"strings"

	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/service"
)

// BuildLazy returns the *service.TimingInfo resulting from the given GetTimingInfo
// request.
func (r *GetTimingInfo) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error) {
	mgr := c.(*Context).ReplayManager

	ctx := &replay.Context{
		DeviceID:  r.Device,
		CaptureID: r.Capture,
	}

	capture, err := service.ResolveCapture(r.Capture, d, l)
	if err != nil {
		return nil, err
	}

	apis := capture.Apis

	results, count := make(chan replay.CallTiming, len(apis)), 0
	for _, apiID := range apis {
		api := gfxapi.Find(gfxapi.ID(apiID.ID))
		if api == nil {
			continue
		}

		query, ok := api.(replay.QueryCallDurations)
		if !ok {
			continue
		}

		go func() {
			results <- <-query.QueryCallDurations(ctx, mgr, r.Flags)
		}()
		count++
	}

	timings := service.TimingInfo{}
	errors := []string{}
	for i := 0; i < count; i++ {
		res := <-results
		if res.Error == nil {
			info := res.TimingInfo
			timings.PerCommand = append(timings.PerCommand, info.PerCommand...)
			timings.PerDrawCall = append(timings.PerDrawCall, info.PerDrawCall...)
			timings.PerFrame = append(timings.PerFrame, info.PerFrame...)
		} else {
			errors = append(errors, res.Error.Error())
		}
	}

	if len(errors) > 0 {
		return nil, fmt.Errorf("QueryCallDurations failed:\n%s", strings.Join(errors, "\n"))
	}

	// TODO: Sort timings

	return &timings, nil
}
