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
	"bytes"
	eb "encoding/binary"
	"flag"
	"fmt"
	"image"
	"testing"
	"time"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary/endian"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/database/store"
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

type atoms atom.List

func (l *atoms) add(atoms ...atom.Atom) atom.ID {
	*l = append(*l, atoms...)
	return atom.ID(len(*l) - 1)
}

func (l *atoms) data(db database.Database, logger log.Logger, data ...interface{}) memory.Pointer {
	buf := &bytes.Buffer{}
	w := endian.Writer(buf, eb.LittleEndian)
	for _, d := range data {
		switch d := d.(type) {
		case float32:
			w.Float32(d)
		default:
			panic(fmt.Errorf("Unsupported data type %T", d))
		}
	}

	blob := &store.Blob{Data: buf.Bytes()}
	id, err := db.Store(blob, logger)
	if err != nil {
		panic(err)
	}

	ptr := memory.Pointer(0x100000)
	l.add(&atom.Observation{
		ResourceID: id,
		Range: memory.Range{
			Base: ptr,
			Size: uint64(len(buf.Bytes())),
		},
	})

	return ptr
}

func initContext(width, height uint32) atoms {
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
	return atoms{
		gles.NewEglCreateContext(eglDisplay, eglConfig, eglShareContext, eglAttribList, eglContext),
		gles.NewEglMakeCurrent(eglDisplay, eglSurface, eglSurface, eglContext, eglTrue),
		gles.NewBackbufferInfo(int32(width), int32(height), color, depth, stencil, true /* resetViewportScissor */),
	}
}

func TestClear(t *testing.T) {
	db, logger := utils.NewInMemoryDatabase(), log.Testing(t)
	w, h := uint32(64), uint32(64)
	atoms := initContext(w, h)
	red := atoms.add(
		gles.NewGlClearColor(1.0, 0.0, 0.0, 1.0),
		gles.NewGlClear(gles.ClearMask_GL_COLOR_BUFFER_BIT),
	)
	green := atoms.add(
		gles.NewGlClearColor(0.0, 1.0, 0.0, 1.0),
		gles.NewGlClear(gles.ClearMask_GL_COLOR_BUFFER_BIT),
	)
	blue := atoms.add(
		gles.NewGlClearColor(0.0, 0.0, 1.0, 1.0),
		gles.NewGlClear(gles.ClearMask_GL_COLOR_BUFFER_BIT),
	)
	black := atoms.add(
		gles.NewGlClearColor(0.0, 0.0, 0.0, 1.0),
		gles.NewGlClear(gles.ClearMask_GL_COLOR_BUFFER_BIT),
	)

	mgr := replay.New(db, logger)
	device := utils.FindLocalDevice(t, mgr)

	ctx := &replay.Context{
		CaptureID: utils.StoreCapture(t, atom.List(atoms), db, logger),
		DeviceID:  device.ID(),
	}

	checkColorBuffer(t, ctx, mgr, w, h, 0, "solid-red", red)
	checkColorBuffer(t, ctx, mgr, w, h, 0, "solid-green", green)
	checkColorBuffer(t, ctx, mgr, w, h, 0, "solid-blue", blue)
	checkColorBuffer(t, ctx, mgr, w, h, 0, "solid-black", black)
}

func TestDrawTriangle(t *testing.T) {
	db, logger := utils.NewInMemoryDatabase(), log.Testing(t)
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
	atoms := initContext(w, h)
	vertices := gles.VertexPointer(
		atoms.data(db, logger,
			float32(+0.0), float32(-0.5),
			float32(-0.5), float32(+0.5),
			float32(+0.5), float32(+0.5),
		),
	)
	clear := atoms.add(
		gles.NewGlClearColor(0.0, 1.0, 0.0, 1.0),
		gles.NewGlClear(gles.ClearMask_GL_COLOR_BUFFER_BIT),
	)
	triangle := atoms.add(
		gles.NewGlCreateShader(gles.ShaderType_GL_VERTEX_SHADER, vs),
		gles.NewGlShaderSource(vs, 1, gles.StringArray{vsSource}, gles.S32Array{int32(len(vsSource))}),
		gles.NewGlCompileShader(vs),
		gles.NewGlCreateShader(gles.ShaderType_GL_FRAGMENT_SHADER, fs),
		gles.NewGlShaderSource(fs, 1, gles.StringArray{fsSource}, gles.S32Array{int32(len(fsSource))}),
		gles.NewGlCompileShader(fs),
		gles.NewGlCreateProgram(program),
		gles.NewGlAttachShader(program, vs),
		gles.NewGlAttachShader(program, fs),
		gles.NewGlLinkProgram(program),
		gles.NewGlUseProgram(program),
		gles.NewGlGetAttribLocation(program, "position", position),
		gles.NewGlEnableVertexAttribArray(position),
		gles.NewGlVertexAttribPointer(position, 2, gles.VertexAttribType_GL_FLOAT, false, 0, vertices),
		gles.NewGlDrawArrays(gles.DrawMode_GL_TRIANGLES, 0, 3),
	)
	mgr := replay.New(db, logger)
	device := utils.FindLocalDevice(t, mgr)

	ctx := &replay.Context{
		CaptureID: utils.StoreCapture(t, atom.List(atoms), db, logger),
		DeviceID:  device.ID(),
	}

	checkColorBuffer(t, ctx, mgr, w, h, 0.0, "solid-green", clear)
	checkColorBuffer(t, ctx, mgr, w, h, 0.1, "triangle", triangle)
}
