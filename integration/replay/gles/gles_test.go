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
	"fmt"
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

const (
	replayTimeout = time.Second * 5

	simpleVSSource = `
		precision mediump float;
		attribute vec2 position;
		void main() {
			gl_Position = vec4(position, 0.5, 1.0);
		}`

	simpleFSSource = `
		precision mediump float;
		void main() {
			gl_FragColor = vec4(1.0, 0.0, 0.0, 1.0);
		}`
)

var (
	triangleVertices = []float32{
		+0.0, -0.5,
		-0.5, +0.5,
		+0.5, +0.5,
	}
)

var generateReferenceImages = flag.Bool("generate", false, "generate reference images")

func p(addr uint64) memory.Pointer {
	return memory.Pointer{Address: addr, Pool: memory.ApplicationPool}
}

func checkColorBuffer(t *testing.T, ctx replay.Context, mgr *replay.Manager, w, h uint32, threshold float64, name string, after atom.ID) {
	select {
	case img := <-gles.API().(replay.QueryColorBuffer).QueryColorBuffer(ctx, mgr, after, w, h, replay.NoWireframe):
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
		// Panic instead of erroring so we see the status of the go-routine we're waiting for.
		panic(fmt.Errorf("Timeout reading ColorBuffer at %d for %s", after, name))
	}
}

func setBackbuffer(width, height int, preserveBuffersOnSwap bool) atom.Atom {
	color := gles.GLenum_GL_RGB565
	depth := gles.GLenum_GL_DEPTH_COMPONENT16
	stencil := gles.GLenum_GL_STENCIL_INDEX8
	return gles.NewBackbufferInfo(gles.GLsizei(width), gles.GLsizei(height), color, depth, stencil,
		true /* resetViewportScissor */, preserveBuffersOnSwap)
}

func initContext(a device.Architecture, d database.Database, l log.Logger, width, height int, preserveBuffersOnSwap bool) *atom.List {
	eglDisplay := p(0x1000)
	eglConfig := p(0x2000)
	eglShareContext := memory.Nullptr
	eglAttribList := []gles.EGLint{0}
	eglSurface := p(0x3000)
	eglContext := p(0x5000)
	eglTrue := gles.EGLBoolean(1)

	atoms := atom.NewList(
		gles.NewEglCreateContext(eglDisplay, eglConfig, eglShareContext, p(0x1000000), eglContext).
			AddRead(atom.Data(a, d, l, p(0x1000000), eglAttribList)),
		gles.NewEglMakeCurrent(eglDisplay, eglSurface, eglSurface, eglContext, eglTrue),
		setBackbuffer(width, height, preserveBuffersOnSwap),
	)
	return atoms
}

func TestClear(t *testing.T) {
	d, l := database.NewInMemory(nil), log.Testing(t)
	mgr := replay.New(d, l)
	device := utils.FindLocalDevice(t, mgr)
	atoms := initContext(device.Info().Architecture(), d, l, 64, 64, false)
	red := atoms.Add(
		gles.NewGlClearColor(1.0, 0.0, 0.0, 1.0),
		gles.NewGlClear(gles.GLbitfield_GL_COLOR_BUFFER_BIT),
	)
	green := atoms.Add(
		gles.NewGlClearColor(0.0, 1.0, 0.0, 1.0),
		gles.NewGlClear(gles.GLbitfield_GL_COLOR_BUFFER_BIT),
	)
	blue := atoms.Add(
		gles.NewGlClearColor(0.0, 0.0, 1.0, 1.0),
		gles.NewGlClear(gles.GLbitfield_GL_COLOR_BUFFER_BIT),
	)
	black := atoms.Add(
		gles.NewGlClearColor(0.0, 0.0, 0.0, 1.0),
		gles.NewGlClear(gles.GLbitfield_GL_COLOR_BUFFER_BIT),
	)

	ctx := replay.Context{
		Capture: utils.StoreCapture(t, atoms, d, l).ID,
		Device:  device.ID(),
	}

	checkColorBuffer(t, ctx, mgr, 64, 64, 0, "solid-red", red)
	checkColorBuffer(t, ctx, mgr, 64, 64, 0, "solid-green", green)
	checkColorBuffer(t, ctx, mgr, 64, 64, 0, "solid-blue", blue)
	checkColorBuffer(t, ctx, mgr, 64, 64, 0, "solid-black", black)
}

func TestDrawTriangle(t *testing.T) {
	d, l := database.NewInMemory(nil), log.Testing(t)
	mgr := replay.New(d, l)
	device := utils.FindLocalDevice(t, mgr)
	a := device.Info().Architecture()
	vs, fs, prog, pos := gles.ShaderId(0x10), gles.ShaderId(0x20), gles.ProgramId(0x30), gles.AttributeLocation(0)
	atoms := initContext(a, d, l, 64, 64, false)
	clear := atoms.Add(
		gles.NewGlClearColor(0.0, 1.0, 0.0, 1.0),
		gles.NewGlClear(gles.GLbitfield_GL_COLOR_BUFFER_BIT),
	)
	atoms.Add(gles.BuildProgram(a, d, l, vs, fs, prog, simpleVSSource, simpleFSSource)...)
	triangle := atoms.Add(
		gles.NewGlLinkProgram(prog),
		gles.NewGlUseProgram(prog),
		gles.NewGlGetAttribLocation(prog, "position", gles.GLint(pos)),
		gles.NewGlEnableVertexAttribArray(pos),
		gles.NewGlVertexAttribPointer(pos, 2, gles.GLenum_GL_FLOAT, gles.GLboolean(0), 0, p(0x100000)).
			AddRead(atom.Data(a, d, l, p(0x100000), triangleVertices)),
		gles.NewGlDrawArrays(gles.GLenum_GL_TRIANGLES, 0, 3),
	)

	ctx := replay.Context{
		Capture: utils.StoreCapture(t, atoms, d, l).ID,
		Device:  device.ID(),
	}

	checkColorBuffer(t, ctx, mgr, 64, 64, 0.0, "solid-green", clear)
	checkColorBuffer(t, ctx, mgr, 64, 64, 0.01, "triangle", triangle)
}

// TestResizeRenderer checks that backbuffers can be resized without destroying
// the current context.
func TestResizeRenderer(t *testing.T) {
	d, l := database.NewInMemory(nil), log.Testing(t)
	mgr := replay.New(d, l)
	device := utils.FindLocalDevice(t, mgr)
	a := device.Info().Architecture()
	vs, fs, prog, pos := gles.ShaderId(0x10), gles.ShaderId(0x20), gles.ProgramId(0x30), gles.AttributeLocation(0)
	atoms := initContext(a, d, l, 8, 8, false) // start with a small backbuffer
	atoms.Add(gles.BuildProgram(a, d, l, vs, fs, prog, simpleVSSource, simpleFSSource)...)
	atoms.Add(
		gles.NewGlLinkProgram(prog),
		gles.NewGlUseProgram(prog),
		gles.NewGlGetAttribLocation(prog, "position", gles.GLint(pos)),
		gles.NewGlEnableVertexAttribArray(pos),
		gles.NewGlVertexAttribPointer(pos, 2, gles.GLenum_GL_FLOAT, gles.GLboolean(0), 0, p(0x100000)).
			AddRead(atom.Data(a, d, l, p(0x100000), triangleVertices)),
	)
	triangle := atoms.Add(
		setBackbuffer(64, 64, false), // Resize just before clearing and drawing.
		gles.NewGlClearColor(0.0, 0.0, 1.0, 1.0),
		gles.NewGlClear(gles.GLbitfield_GL_COLOR_BUFFER_BIT),
		gles.NewGlDrawArrays(gles.GLenum_GL_TRIANGLES, 0, 3),
	)

	ctx := replay.Context{
		Capture: utils.StoreCapture(t, atoms, d, l).ID,
		Device:  device.ID(),
	}

	checkColorBuffer(t, ctx, mgr, 64, 64, 0.01, "triangle_2", triangle)
}

// TestPreserveBuffersOnSwap checks that when the preserveBuffersOnSwap flag is
// set, the backbuffer is preserved between calls to eglSwapBuffers().
func TestPreserveBuffersOnSwap(t *testing.T) {
	d, l := database.NewInMemory(nil), log.Testing(t)
	mgr := replay.New(d, l)
	device := utils.FindLocalDevice(t, mgr)
	a := device.Info().Architecture()
	atoms := initContext(a, d, l, 64, 64, true)
	clear := atoms.Add(
		gles.NewGlClearColor(0.0, 0.0, 1.0, 1.0),
		gles.NewGlClear(gles.GLbitfield_GL_COLOR_BUFFER_BIT),
	)
	swapA := atoms.Add(gles.NewEglSwapBuffers(memory.Nullptr, memory.Nullptr, 1))
	swapB := atoms.Add(gles.NewEglSwapBuffers(memory.Nullptr, memory.Nullptr, 1))
	swapC := atoms.Add(gles.NewEglSwapBuffers(memory.Nullptr, memory.Nullptr, 1))

	ctx := replay.Context{
		Capture: utils.StoreCapture(t, atoms, d, l).ID,
		Device:  device.ID(),
	}

	checkColorBuffer(t, ctx, mgr, 64, 64, 0.0, "solid-blue", clear)
	checkColorBuffer(t, ctx, mgr, 64, 64, 0.0, "solid-blue", swapA)
	checkColorBuffer(t, ctx, mgr, 64, 64, 0.0, "solid-blue", swapB)
	checkColorBuffer(t, ctx, mgr, 64, 64, 0.0, "solid-blue", swapC)
}
