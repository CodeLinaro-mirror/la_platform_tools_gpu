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

import "android.googlesource.com/platform/tools/gpu/atom"

// tweaker provides a set of methods for temporarily changing the GLES state.
type tweaker struct {
	out  atom.Writer
	ctx  *Context
	undo []func()
}

// revert undoes all the changes made by the tweaker.
func (t *tweaker) revert() {
	for i := len(t.undo) - 1; i >= 0; i-- {
		t.undo[i]()
	}
	t.undo = nil
}

func (t *tweaker) glEnable(v GLenum) {
	if !t.ctx.Capabilities[v] {
		t.undo = append(t.undo, func() {
			t.out.Write(atom.NoID, NewGlDisable(v))
			t.ctx.Capabilities[v] = false
		})
		t.out.Write(atom.NoID, NewGlEnable(v))
		t.ctx.Capabilities[v] = true
	}
}

func (t *tweaker) glDisable(v GLenum) {
	if t.ctx.Capabilities[v] {
		t.undo = append(t.undo, func() {
			t.out.Write(atom.NoID, NewGlEnable(v))
			t.ctx.Capabilities[v] = true
		})
		t.out.Write(atom.NoID, NewGlDisable(v))
		t.ctx.Capabilities[v] = false
	}
}

func (t *tweaker) glDepthMask(v bool) {
	if o := t.ctx.Rasterizing.DepthMask; o != v {
		t.undo = append(t.undo, func() {
			t.out.Write(atom.NoID, NewGlDepthMask(o))
			t.ctx.Rasterizing.DepthMask = o
		})
		t.out.Write(atom.NoID, NewGlDepthMask(v))
		t.ctx.Rasterizing.DepthMask = v
	}
}

func (t *tweaker) glDepthFunc(v GLenum) {
	if o := t.ctx.Rasterizing.DepthTestFunction; o != v {
		t.undo = append(t.undo, func() {
			t.out.Write(atom.NoID, NewGlDepthFunc(o))
			t.ctx.Rasterizing.DepthTestFunction = o
		})
		t.out.Write(atom.NoID, NewGlDepthFunc(v))
		t.ctx.Rasterizing.DepthTestFunction = v
	}
}

func (t *tweaker) glBlendColor(r, g, b, a float32) {
	n := Color{Red: r, Green: g, Blue: b, Alpha: a}
	if o := t.ctx.Blending.BlendColor; o != n {
		t.undo = append(t.undo, func() {
			t.out.Write(atom.NoID, NewGlBlendColor(o.Red, o.Green, o.Blue, o.Alpha))
			t.ctx.Blending.BlendColor = o
		})
		t.out.Write(atom.NoID, NewGlBlendColor(r, g, b, a))
		t.ctx.Blending.BlendColor = n
	}
}

func (t *tweaker) glBlendFunc(src, dst GLenum) {
	t.glBlendFuncSeparate(src, dst, src, dst)
}

func (t *tweaker) glBlendFuncSeparate(srcRGB, dstRGB, srcA, dstA GLenum) {
	oSrcRGB, oDstRGB, oSrcA, oDstA :=
		t.ctx.Blending.SrcRgbBlendFactor,
		t.ctx.Blending.DstRgbBlendFactor,
		t.ctx.Blending.SrcAlphaBlendFactor,
		t.ctx.Blending.DstAlphaBlendFactor
	if oSrcRGB != srcRGB || oDstRGB != dstRGB || oSrcA != srcA || oDstA != dstA {
		t.undo = append(t.undo, func() {
			t.out.Write(atom.NoID, NewGlBlendFuncSeparate(oSrcRGB, oDstRGB, oSrcA, oDstA))
			t.ctx.Blending.SrcRgbBlendFactor,
				t.ctx.Blending.DstRgbBlendFactor,
				t.ctx.Blending.SrcAlphaBlendFactor,
				t.ctx.Blending.DstAlphaBlendFactor = oSrcRGB, oDstRGB, oSrcA, oDstA
		})
		t.out.Write(atom.NoID, NewGlBlendFuncSeparate(srcRGB, dstRGB, srcA, dstA))
		t.ctx.Blending.SrcRgbBlendFactor,
			t.ctx.Blending.DstRgbBlendFactor,
			t.ctx.Blending.SrcAlphaBlendFactor,
			t.ctx.Blending.DstAlphaBlendFactor = srcRGB, dstRGB, srcA, dstA
	}
}
