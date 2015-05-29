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
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/device"
	"android.googlesource.com/platform/tools/gpu/gfxapi/gles"
	"android.googlesource.com/platform/tools/gpu/integration/replay/utils"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/replay"
)

const replayTimeout = time.Second * 5

var generateReferenceImages = flag.Bool("generate", false, "generate reference images")

func checkColorBuffer(t *testing.T, ctx *replay.Context, mgr *replay.Manager, w, h uint32, threshold float64, name string, after atom.ID) {
	select {
	case img := <-gles.API().(replay.QueryColorBuffer).QueryColorBuffer(ctx, mgr, after, w, h, false):
		if img.Error != nil {
			t.Errorf("Failed to read ColorBuffer at %d for %s. Reason: %v", after, name, img.Error)
			return
		}
		if w*h*4 != uint32(len(img.Data)) {
			t.Errorf("ColorBuffer does not contain the expected number of bytes. Expected: %v, got: %v", w*h*4, len(img.Data))
			return
		}
		got := &image.NRGBA{
			Pix:    img.Data,
			Stride: int(w * 4),
			Rect:   image.Rect(0, 0, int(w), int(h)),
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

func initContext(a device.Architecture, d database.Database, l log.Logger, width, height uint32) atom.List {
	eglDisplay := memory.Pointer(0x1000)
	eglConfig := memory.Pointer(0x2000)
	eglShareContext := memory.Pointer(0)
	eglAttribList := []gles.EGLint{0}
	eglSurface := memory.Pointer(0x3000)
	eglContext := memory.Pointer(0x5000)
	eglTrue := gles.EGLBoolean(1)
	color := gles.RenderbufferFormat_GL_RGB565
	depth := gles.RenderbufferFormat_GL_DEPTH_COMPONENT16
	stencil := gles.RenderbufferFormat_GL_STENCIL_INDEX8
	return atom.List{
		gles.NewEglCreateContext(eglDisplay, eglConfig, eglShareContext, 0x1000000, eglContext).
			AddRead(atom.Data(a, d, l, 0x1000000, eglAttribList)),
		gles.NewEglMakeCurrent(eglDisplay, eglSurface, eglSurface, eglContext, eglTrue),
		gles.NewBackbufferInfo(int32(width), int32(height), color, depth, stencil, true /* resetViewportScissor */),
	}
}

func TestClear(t *testing.T) {
	d, l := database.InMemory(), log.Testing(t)
	mgr := replay.New(d, l)
	device := utils.FindLocalDevice(t, mgr)
	w, h := uint32(64), uint32(64)
	atoms := initContext(device.Info().Architecture(), d, l, w, h)
	red := atoms.Add(
		gles.NewGlClearColor(1.0, 0.0, 0.0, 1.0),
		gles.NewGlClear(gles.ClearMask_GL_COLOR_BUFFER_BIT),
	)
	green := atoms.Add(
		gles.NewGlClearColor(0.0, 1.0, 0.0, 1.0),
		gles.NewGlClear(gles.ClearMask_GL_COLOR_BUFFER_BIT),
	)
	blue := atoms.Add(
		gles.NewGlClearColor(0.0, 0.0, 1.0, 1.0),
		gles.NewGlClear(gles.ClearMask_GL_COLOR_BUFFER_BIT),
	)
	black := atoms.Add(
		gles.NewGlClearColor(0.0, 0.0, 0.0, 1.0),
		gles.NewGlClear(gles.ClearMask_GL_COLOR_BUFFER_BIT),
	)

	ctx := &replay.Context{
		CaptureID: utils.StoreCapture(t, atom.List(atoms), d, l),
		DeviceID:  device.ID(),
	}

	checkColorBuffer(t, ctx, mgr, w, h, 0, "solid-red", red)
	checkColorBuffer(t, ctx, mgr, w, h, 0, "solid-green", green)
	checkColorBuffer(t, ctx, mgr, w, h, 0, "solid-blue", blue)
	checkColorBuffer(t, ctx, mgr, w, h, 0, "solid-black", black)
}

func TestDrawTriangle(t *testing.T) {
	d, l := database.InMemory(), log.Testing(t)
	mgr := replay.New(d, l)
	device := utils.FindLocalDevice(t, mgr)
	a := device.Info().Architecture()
	w, h := uint32(64), uint32(64)
	vs, fs := gles.ShaderId(0x10), gles.ShaderId(0x20)
	program := gles.ProgramId(0x30)
	position := gles.AttributeLocation(0)
	vsSource := `
		precision mediump float;
		attribute vec2 position;
		void main() {
			gl_Position = vec4(position, 0.5, 1.0);
		}`
	fsSource := `
		precision mediump float;
		void main() {
			gl_FragColor = vec4(1.0, 0.0, 0.0, 1.0);
		}`
	atoms := initContext(a, d, l, w, h)
	vertices := []float32{
		+0.0, -0.5,
		-0.5, +0.5,
		+0.5, +0.5,
	}
	clear := atoms.Add(
		gles.NewGlClearColor(0.0, 1.0, 0.0, 1.0),
		gles.NewGlClear(gles.ClearMask_GL_COLOR_BUFFER_BIT),
	)
	triangle := atoms.Add(
		gles.NewGlCreateShader(gles.ShaderType_GL_VERTEX_SHADER, vs),
		gles.NewGlShaderSource(vs, 1, 0x100000, 0x100010).
			AddRead(atom.Data(a, d, l, 0x100000, memory.Pointer(0x100020))).
			AddRead(atom.Data(a, d, l, 0x100010, int32(len(vsSource)))).
			AddRead(atom.Data(a, d, l, 0x100020, vsSource)),
		gles.NewGlCompileShader(vs),
		gles.NewGlCreateShader(gles.ShaderType_GL_FRAGMENT_SHADER, fs),
		gles.NewGlShaderSource(fs, 1, 0x100000, 0x100010).
			AddRead(atom.Data(a, d, l, 0x100000, memory.Pointer(0x100020))).
			AddRead(atom.Data(a, d, l, 0x100010, int32(len(fsSource)))).
			AddRead(atom.Data(a, d, l, 0x100020, fsSource)),
		gles.NewGlCompileShader(fs),
		gles.NewGlCreateProgram(program),
		gles.NewGlAttachShader(program, vs),
		gles.NewGlAttachShader(program, fs),
		gles.NewGlLinkProgram(program),
		gles.NewGlUseProgram(program),
		gles.NewGlGetAttribLocation(program, "position", position),
		gles.NewGlEnableVertexAttribArray(position),
		gles.NewGlVertexAttribPointer(position, 2, gles.VertexAttribType_GL_FLOAT, false, 0, 0x100000).
			AddRead(atom.Data(a, d, l, 0x100000, vertices)),
		gles.NewGlDrawArrays(gles.DrawMode_GL_TRIANGLES, 0, 3),
	)

	ctx := &replay.Context{
		CaptureID: utils.StoreCapture(t, atom.List(atoms), d, l),
		DeviceID:  device.ID(),
	}

	checkColorBuffer(t, ctx, mgr, w, h, 0.0, "solid-green", clear)
	checkColorBuffer(t, ctx, mgr, w, h, 0.1, "triangle", triangle)
}
