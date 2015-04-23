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

const replayTimeout = time.Second * 5

var generateReferenceImages = flag.Bool("generate", false, "generate reference images")

func checkColorBuffer(t *testing.T, ctx *replay.Context, mgr *replay.Manager, after atom.ID, w, h uint32, name string, threshold float64) {
	select {
	case img := <-gles.API().(replay.QueryColorBuffer).QueryColorBuffer(ctx, mgr, after, w, h, false):
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
	db, logger := utils.NewInMemoryDatabase(), log.Testing(t)
	w, h := uint32(64), uint32(64)
	eglDisplay := gles.EGLDisplay(0x1000)
	eglConfig := gles.EGLConfig(0x2000)
	eglShareContext := gles.EGLContext(0)
	eglAttribList := gles.EGLintArray{}
	eglSurface := gles.EGLSurface(0x3000)
	eglContext := gles.EGLContext(0x5000)
	eglTrue := gles.EGLBoolean(1)
	color := gles.RenderbufferFormat_GL_RGB565
	depth := gles.RenderbufferFormat_GL_DEPTH_COMPONENT16
	stencil := gles.RenderbufferFormat_GL_STENCIL_INDEX8
	atoms := atom.List{
		/* 0x0 */ gles.NewEglCreateContext(eglDisplay, eglConfig, eglShareContext, eglAttribList, eglContext),
		/* 0x1 */ gles.NewEglMakeCurrent(eglDisplay, eglSurface, eglSurface, eglContext, eglTrue),
		/* 0x2 */ gles.NewBackbufferInfo(int32(w), int32(h), color, depth, stencil, true /* resetViewportScissor */),
		/* 0x3 */ gles.NewGlClearColor(1.0, 0.0, 0.0, 1.0),
		/* 0x4 */ gles.NewGlClear(gles.ClearMask_GL_COLOR_BUFFER_BIT),
		/* 0x5 */ gles.NewGlClearColor(0.0, 1.0, 0.0, 1.0),
		/* 0x6 */ gles.NewGlClear(gles.ClearMask_GL_COLOR_BUFFER_BIT),
		/* 0x7 */ gles.NewGlClearColor(0.0, 0.0, 1.0, 1.0),
		/* 0x8 */ gles.NewGlClear(gles.ClearMask_GL_COLOR_BUFFER_BIT),
		/* 0x9 */ gles.NewGlClearColor(0.0, 0.0, 0.0, 1.0),
		/* 0xa */ gles.NewGlClear(gles.ClearMask_GL_COLOR_BUFFER_BIT),
	}

	mgr := replay.New(db, logger)
	device := utils.FindLocalDevice(t, mgr)

	ctx := &replay.Context{
		CaptureID: utils.StoreCapture(t, atoms, db, logger),
		DeviceID:  device.ID(),
	}

	checkColorBuffer(t, ctx, mgr, 0x4, w, h, "solid-red", 0)
	checkColorBuffer(t, ctx, mgr, 0x6, w, h, "solid-green", 0)
	checkColorBuffer(t, ctx, mgr, 0x8, w, h, "solid-blue", 0)
	checkColorBuffer(t, ctx, mgr, 0xa, w, h, "transparent", 0)
}
