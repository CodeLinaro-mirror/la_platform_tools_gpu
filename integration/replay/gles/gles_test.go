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
	"sync"
	"testing"
	"time"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/check"
	"android.googlesource.com/platform/tools/gpu/client/gapir"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/device"
	"android.googlesource.com/platform/tools/gpu/gfxapi/gles"
	"android.googlesource.com/platform/tools/gpu/image"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/replay"
)

const (
	replayTimeout = time.Second * 5

	simpleVSSource = `
		precision mediump float;
		attribute vec3 position;
		void main() {
			gl_Position = vec4(position, 1.0);
		}`

	simpleFSSource = `
		precision mediump float;
		void main() {
			gl_FragColor = vec4(1.0, 0.0, 0.0, 1.0);
		}`
)

var (
	triangleVertices = []float32{
		+0.0, -0.5, 0.1,
		-0.5, +0.5, 0.5,
		+0.5, +0.5, 0.9,
	}
)

var generateReferenceImages = flag.Bool("generate", false, "generate reference images")

func p(addr uint64) memory.Pointer {
	return memory.Pointer{Address: addr, Pool: memory.ApplicationPool}
}

func checkImage(t *testing.T, name string, got *image.Image, threshold float64) {
	if *generateReferenceImages {
		storeReferenceImage(t, name, got)
	} else {
		expected := loadReferenceImage(t, name)
		diff, err := image.Difference(got, expected)
		if err != nil {
			t.Errorf("image.Difference returned error: %v", err)
			return
		}
		if diff > threshold {
			t.Errorf("%v had error of %v%% which is above the threshold of %v%%", name, diff*100, threshold*100)
		}
	}
}

func checkIssues(t *testing.T, ctx replay.Context, mgr *replay.Manager, expected []replay.Issue, done *sync.WaitGroup) {
	if done != nil {
		defer done.Done()
	}
	issues := gles.API().(replay.QueryIssues).QueryIssues(ctx, mgr)
	timeout := time.Tick(replayTimeout)
	got := []replay.Issue{}
	for {
		select {
		case issue, more := <-issues:
			if more {
				got = append(got, issue)
			} else {
				check.SlicesDeepEqual(t, got, expected)
				return
			}
		case <-timeout:
			// Panic instead of erroring so we see the status of the go-routine we're waiting for.
			panic(fmt.Errorf("Timeout querying for issue"))
		}
	}
}

func checkColorBuffer(t *testing.T, ctx replay.Context, mgr *replay.Manager, w, h uint32, threshold float64, name string, after atom.ID, done *sync.WaitGroup) {
	if done != nil {
		defer done.Done()
	}
	select {
	case res := <-gles.API().(replay.QueryColorBuffer).QueryColorBuffer(ctx, mgr, after, w, h, replay.NoWireframe):
		if res.Error != nil {
			t.Errorf("Failed to read ColorBuffer at %d for %s. Reason: %v", after, name, res.Error)
			return
		}
		checkImage(t, name, res.Image, threshold)
	case <-time.Tick(replayTimeout):
		// Panic instead of erroring so we see the status of the go-routine we're waiting for.
		panic(fmt.Errorf("Timeout reading ColorBuffer at %d for %s", after, name))
	}
}

func checkDepthBuffer(t *testing.T, ctx replay.Context, mgr *replay.Manager, w, h uint32, threshold float64, name string, after atom.ID, done *sync.WaitGroup) {
	if done != nil {
		defer done.Done()
	}
	select {
	case res := <-gles.API().(replay.QueryDepthBuffer).QueryDepthBuffer(ctx, mgr, after):
		if res.Error != nil {
			t.Errorf("Failed to read DepthBuffer at %d for %s. Reason: %v", after, name, res.Error)
			return
		}
		checkImage(t, name, res.Image, threshold)
	case <-time.Tick(replayTimeout):
		// Panic instead of erroring so we see the status of the go-routine we're waiting for.
		panic(fmt.Errorf("Timeout reading DepthBuffer at %d for %s", after, name))
	}
}

type ctxCfg struct {
	context replay.Context
	config  replay.Config
}

func (c ctxCfg) String() string { return fmt.Sprintf("Context: %+v, Config: %+v", c.context, c.config) }

func checkReplay(t *testing.T, expectedContext replay.Context, expectedBatchCount int) func() {
	batchCount := 0
	uniqueCtxCfgs := map[ctxCfg]struct{}{}
	replay.Events.OnReplay = func(device gapir.Device, context replay.Context, config replay.Config, requests []replay.Request) {
		if expectedContext != context {
			t.Errorf("Expected replay context: %v, got: %v", expectedContext, context)
		}
		batchCount++
		uniqueCtxCfgs[ctxCfg{context, config}] = struct{}{}
	}
	return func() {
		if batchCount != expectedBatchCount {
			t.Errorf("Expected %v replay batches, got %v", expectedBatchCount, batchCount)
			t.Errorf("%d unique context-config pairs:", len(uniqueCtxCfgs))
			for cc := range uniqueCtxCfgs {
				t.Errorf(" • %v", cc)
			}
		}
	}
}

func newContextInfo(a device.Architecture, d database.Database, l log.Logger, width, height int, preserveBuffersOnSwap bool) atom.Atom {
	names := []gles.GLenum{}
	offsets := []uint32{}
	sizes := []uint32{}
	data := ""
	for name, value := range map[gles.GLenum]string{
		gles.GLenum_GL_RENDERER: "test-driver",
		gles.GLenum_GL_VENDOR:   "Super-Awesome-Graphics-Inc",
		gles.GLenum_GL_VERSION:  "OpenGL ES 2.0",
	} {
		names = append(names, name)
		offsets = append(offsets, uint32(len(data)))
		sizes = append(sizes, uint32(len(value)))
		data = data + value
	}

	return gles.NewContextInfo(
		uint32(len(names)),
		p(0x10000),
		p(0x20000),
		p(0x30000),
		p(0x40000),
		gles.GLsizei(width),
		gles.GLsizei(height),
		gles.GLenum_GL_RGB565,
		gles.GLenum_GL_DEPTH_COMPONENT16,
		gles.GLenum_GL_STENCIL_INDEX8,
		true,
		preserveBuffersOnSwap).
		AddRead(atom.Data(a, d, l, p(0x10000), names)).
		AddRead(atom.Data(a, d, l, p(0x20000), offsets)).
		AddRead(atom.Data(a, d, l, p(0x30000), sizes)).
		AddRead(atom.Data(a, d, l, p(0x40000), data))
}

// firstAtomID is the identifier of the first atom after initContext.
const firstAtomID atom.ID = 3

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
		newContextInfo(a, d, l, width, height, preserveBuffersOnSwap),
	)
	return atoms
}

func TestClear(t *testing.T) {
	d, l := database.NewInMemory(nil), log.Testing(t)
	mgr := replay.New(d, l)
	device := gapir.FindLocalDevice(t, mgr.Discovery())
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
		Capture: gapir.StoreCapture(t, atoms, d, l).ID,
		Device:  device.ID(),
	}

	defer checkReplay(t, ctx, 1)() // expect a single replay batch.

	done := &sync.WaitGroup{}
	done.Add(4)
	go checkColorBuffer(t, ctx, mgr, 64, 64, 0, "solid-red", red, done)
	go checkColorBuffer(t, ctx, mgr, 64, 64, 0, "solid-green", green, done)
	go checkColorBuffer(t, ctx, mgr, 64, 64, 0, "solid-blue", blue, done)
	go checkColorBuffer(t, ctx, mgr, 64, 64, 0, "solid-black", black, done)
	done.Wait()
}

func TestDrawTriangle(t *testing.T) {
	d, l := database.NewInMemory(nil), log.Testing(t)
	mgr := replay.New(d, l)
	device := gapir.FindLocalDevice(t, mgr.Discovery())
	a := device.Info().Architecture()
	vs, fs, prog, pos := gles.ShaderId(0x10), gles.ShaderId(0x20), gles.ProgramId(0x30), gles.AttributeLocation(0)
	atoms := initContext(a, d, l, 64, 64, false)
	atoms.Add(gles.NewGlEnable(gles.GLenum_GL_DEPTH_TEST)) // Required for depth-writing
	clear := atoms.Add(
		gles.NewGlClearColor(0.0, 1.0, 0.0, 1.0),
		gles.NewGlClear(gles.GLbitfield_GL_COLOR_BUFFER_BIT|gles.GLbitfield_GL_DEPTH_BUFFER_BIT),
	)
	atoms.Add(gles.BuildProgram(a, d, l, vs, fs, prog, simpleVSSource, simpleFSSource)...)
	triangle := atoms.Add(
		gles.NewGlLinkProgram(prog),
		gles.NewGlUseProgram(prog),
		gles.NewGlGetAttribLocation(prog, "position", gles.GLint(pos)),
		gles.NewGlEnableVertexAttribArray(pos),
		gles.NewGlVertexAttribPointer(pos, 3, gles.GLenum_GL_FLOAT, gles.GLboolean(0), 0, p(0x100000)),
		gles.NewGlDrawArrays(gles.GLenum_GL_TRIANGLES, 0, 3).
			AddRead(atom.Data(a, d, l, p(0x100000), triangleVertices)),
	)

	ctx := replay.Context{
		Capture: gapir.StoreCapture(t, atoms, d, l).ID,
		Device:  device.ID(),
	}

	defer checkReplay(t, ctx, 1)() // expect a single replay batch.

	done := &sync.WaitGroup{}
	done.Add(4)
	go checkColorBuffer(t, ctx, mgr, 64, 64, 0.0, "solid-green", clear, done)
	go checkDepthBuffer(t, ctx, mgr, 64, 64, 0.0, "one-depth", clear, done)
	go checkColorBuffer(t, ctx, mgr, 64, 64, 0.01, "triangle", triangle, done)
	go checkDepthBuffer(t, ctx, mgr, 64, 64, 0.01, "triangle-depth", triangle, done)
	done.Wait()
}

// TestResizeRenderer checks that backbuffers can be resized without destroying
// the current context.
func TestResizeRenderer(t *testing.T) {
	d, l := database.NewInMemory(nil), log.Testing(t)
	mgr := replay.New(d, l)
	device := gapir.FindLocalDevice(t, mgr.Discovery())
	a := device.Info().Architecture()
	vs, fs, prog, pos := gles.ShaderId(0x10), gles.ShaderId(0x20), gles.ProgramId(0x30), gles.AttributeLocation(0)
	atoms := initContext(a, d, l, 8, 8, false) // start with a small backbuffer
	atoms.Add(gles.BuildProgram(a, d, l, vs, fs, prog, simpleVSSource, simpleFSSource)...)
	atoms.Add(
		gles.NewGlLinkProgram(prog),
		gles.NewGlUseProgram(prog),
		gles.NewGlGetAttribLocation(prog, "position", gles.GLint(pos)),
		gles.NewGlEnableVertexAttribArray(pos),
		gles.NewGlVertexAttribPointer(pos, 3, gles.GLenum_GL_FLOAT, gles.GLboolean(0), 0, p(0x100000)),
	)
	triangle := atoms.Add(
		newContextInfo(a, d, l, 64, 64, false), // Resize just before clearing and drawing.
		gles.NewGlClearColor(0.0, 0.0, 1.0, 1.0),
		gles.NewGlClear(gles.GLbitfield_GL_COLOR_BUFFER_BIT),
		gles.NewGlDrawArrays(gles.GLenum_GL_TRIANGLES, 0, 3).
			AddRead(atom.Data(a, d, l, p(0x100000), triangleVertices)),
	)

	ctx := replay.Context{
		Capture: gapir.StoreCapture(t, atoms, d, l).ID,
		Device:  device.ID(),
	}

	checkColorBuffer(t, ctx, mgr, 64, 64, 0.01, "triangle_2", triangle, nil)
}

// TestPreserveBuffersOnSwap checks that when the preserveBuffersOnSwap flag is
// set, the backbuffer is preserved between calls to eglSwapBuffers().
func TestPreserveBuffersOnSwap(t *testing.T) {
	d, l := database.NewInMemory(nil), log.Testing(t)
	mgr := replay.New(d, l)
	device := gapir.FindLocalDevice(t, mgr.Discovery())
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
		Capture: gapir.StoreCapture(t, atoms, d, l).ID,
		Device:  device.ID(),
	}

	done := &sync.WaitGroup{}
	done.Add(4)
	go checkColorBuffer(t, ctx, mgr, 64, 64, 0.0, "solid-blue", clear, done)
	go checkColorBuffer(t, ctx, mgr, 64, 64, 0.0, "solid-blue", swapA, done)
	go checkColorBuffer(t, ctx, mgr, 64, 64, 0.0, "solid-blue", swapB, done)
	go checkColorBuffer(t, ctx, mgr, 64, 64, 0.0, "solid-blue", swapC, done)
	done.Wait()
}

// TestIssues tests the QueryIssues replay command with various streams.
func TestIssues(t *testing.T) {
	d, l := database.NewInMemory(nil), log.Testing(t)
	mgr := replay.New(d, l)
	device := gapir.FindLocalDevice(t, mgr.Discovery())
	a := device.Info().Architecture()

	done := &sync.WaitGroup{}

	for _, test := range []struct {
		name     string
		atoms    []atom.Atom
		expected []replay.Issue
	}{
		{
			"glClear - no errors",
			[]atom.Atom{
				gles.NewGlClearColor(0.0, 0.0, 1.0, 1.0),
				gles.NewGlClear(gles.GLbitfield_GL_COLOR_BUFFER_BIT),
			},
			[]replay.Issue{},
		}, {
			"glActiveTexture - invalid enum",
			[]atom.Atom{&directCall{atom: gles.NewGlActiveTexture(gles.GLenum_GL_TEXTURE0 - 1)}},
			[]replay.Issue{
				replay.Issue{
					Atom:     firstAtomID,
					Severity: log.Error,
					Error:    fmt.Errorf("glGetError() returned %v", gles.GLenum_GL_INVALID_ENUM),
				},
			},
		},
	} {
		atoms := initContext(a, d, l, 64, 64, true)
		atoms.Add(test.atoms...)
		ctx := replay.Context{
			Capture: gapir.StoreCapture(t, atoms, d, l).ID,
			Device:  device.ID(),
		}
		done.Add(1)
		go checkIssues(t, ctx, mgr, test.expected, done)
	}

	done.Wait()
}
