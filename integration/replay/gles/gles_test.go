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

package gles

import (
	"flag"
	"image"
	"testing"
	"time"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/gfxapi/gles"
	"android.googlesource.com/platform/tools/gpu/integration/replay/utils"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/replay"
)

const replayTimeout = time.Second * 10

var generateReferenceImages = flag.Bool("generate", false, "generate reference images")

func checkColorBuffer(t *testing.T, ctx *replay.Context, mgr *replay.Manager, after atom.ID, w, h uint32, name string, threshold float64) {
	select {
	case img := <-gles.API().ColorBuffer(ctx, mgr, after, w, h, false):
		if img.Error != nil {
			t.Errorf("Failed to read ColorBuffer at %d for %s. Reason: %v", after, name, img.Error)
			return
		}
		got := &image.NRGBA{
			Pix:  img.Data,
			Rect: image.Rect(0, 0, int(w), int(h)),
		}
		if *generateReferenceImages {
			storeReferenceImage(t, name, got)
		} else {
			expected := loadReferenceImage(t, name)
			err := compareImages(t, expected, got)
			if err > threshold {
				t.Errorf("%v had error of %v%% which is above the threshold of %v%%", name, err*100, threshold*100)
			}
		}
	case <-time.Tick(replayTimeout):
		t.Errorf("Timeout reading ColorBuffer at %d for %s", after, name)
	}
}

func TestClear(t *testing.T) {
	cid := atom.ContextID(0)
	db, logger := utils.NewInMemoryDatabase(), log.Testing(t)
	w, h := uint32(64), uint32(64)
	atoms := atom.List{
		/* 0 */ gles.NewInit(cid, int32(w), int32(h),
			gles.RenderbufferFormat_GL_RGB565,
			gles.RenderbufferFormat_GL_DEPTH_COMPONENT16,
			gles.RenderbufferFormat_GL_STENCIL_INDEX8,
		),
		/* 1 */ gles.NewGlClearColor(cid, 1.0, 0.0, 0.0, 1.0),
		/* 2 */ gles.NewGlClear(cid, gles.ClearMask_GL_COLOR_BUFFER_BIT),
		/* 3 */ gles.NewGlClearColor(cid, 0.0, 1.0, 0.0, 1.0),
		/* 4 */ gles.NewGlClear(cid, gles.ClearMask_GL_COLOR_BUFFER_BIT),
		/* 5 */ gles.NewGlClearColor(cid, 0.0, 0.0, 1.0, 1.0),
		/* 6 */ gles.NewGlClear(cid, gles.ClearMask_GL_COLOR_BUFFER_BIT),
		/* 7 */ gles.NewGlClearColor(cid, 0.0, 0.0, 0.0, 1.0),
		/* 8 */ gles.NewGlClear(cid, gles.ClearMask_GL_COLOR_BUFFER_BIT),
	}

	mgr := replay.New(db, logger)
	device := utils.FindLocalDevice(t, mgr)

	ctx := &replay.Context{
		CaptureID: utils.StoreCapture(t, atoms, db, logger),
		DeviceID:  device.ID(),
	}

	checkColorBuffer(t, ctx, mgr, 2, w, h, "solid-red", 0)
	checkColorBuffer(t, ctx, mgr, 4, w, h, "solid-green", 0)
	checkColorBuffer(t, ctx, mgr, 6, w, h, "solid-blue", 0)
	checkColorBuffer(t, ctx, mgr, 8, w, h, "transparent", 0)
}
