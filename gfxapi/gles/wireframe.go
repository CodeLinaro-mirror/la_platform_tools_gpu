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
	"fmt"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/endian"
	"android.googlesource.com/platform/tools/gpu/binary/flat"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
)

// wireframe returns an atom transform that replaces all draw calls of triangle
// primitives with draw calls of a wireframe equivalent.
func wireframe(d database.Database, l log.Logger) atom.Transformer {
	l = log.Enter(l, "Wireframe")

	s := gfxapi.NewState()
	return atom.Transform("Wireframe", func(i atom.ID, a atom.Atom, out atom.Writer) {
		if err := a.Mutate(s, d, l); err != nil {
			log.Errorf(l, "%v", err)
		}

		if a.Flags().IsDrawCall() {
			drawWireframe(i, a, s, d, l, out)
		} else {
			out.Write(i, a)
		}
	})
}

// wireframeOverlay returns an atom transform that renders the wireframe of the
// mesh over of the specified draw call.
func wireframeOverlay(id atom.ID, d database.Database, l log.Logger) atom.Transformer {
	l = log.Enter(l, "WireframeOverlay")

	s := gfxapi.NewState()
	return atom.Transform("WireframeOverlay", func(i atom.ID, a atom.Atom, out atom.Writer) {
		if err := a.Mutate(s, d, l); err != nil {
			log.Errorf(l, "%v", err)
		}

		if i == id && a.Flags().IsDrawCall() {
			out.Write(atom.NoID, a)

			t := tweaker{out: out, ctx: getContext(s)}
			t.glEnable(GLenum_GL_BLEND)
			t.glBlendColor(1.5, 0.5, 1.0, 1.0)
			t.glBlendFunc(GLenum_GL_CONSTANT_COLOR, GLenum_GL_ZERO)
			if t.ctx.Rasterizing.DepthMask && t.ctx.Capabilities[GLenum_GL_DEPTH_TEST] {
				t.glDepthMask(false)
				t.glDepthFunc(GLenum_GL_EQUAL)
			}

			drawWireframe(i, a, s, d, l, out)

			t.revert()
		} else {
			out.Write(i, a)
		}
	})
}

func drawWireframe(i atom.ID, a atom.Atom, s *gfxapi.State, d database.Database, l log.Logger, out atom.Writer) {
	c := getContext(s)
	indices, drawMode, err := getIndices(a, c, s, d, l)
	if err != nil {
		log.Errorf(l, err.Error())
		return
	}
	indices, drawMode, err = makeWireframe(indices, drawMode)
	if err != nil {
		log.Errorf(l, err.Error())
		return
	}

	// Store the wire-frame data to a temporary address.
	wireframeData, wireframeDataType := encodeIndices(indices)
	resID, err := database.Store(wireframeData, d, l)
	if err != nil {
		panic(err)
	}

	// Unbind the index buffer
	oldIndexBufferID := c.BoundBuffers[GLenum_GL_ELEMENT_ARRAY_BUFFER]
	out.Write(atom.NoID,
		NewGlBindBuffer(GLenum_GL_ELEMENT_ARRAY_BUFFER, 0).
			AddRead(memory.Tmp.Range(uint64(len(wireframeData))), resID))

	// Draw the wire-frame
	out.Write(i, NewGlDrawElements(
		drawMode, GLsizei(len(indices)), wireframeDataType, memory.Tmp))

	// Rebind the old index buffer
	out.Write(atom.NoID, NewGlBindBuffer(
		GLenum_GL_ELEMENT_ARRAY_BUFFER, oldIndexBufferID))
}

type index uint32

// TODO: The decode/encode methods below assume little endian

func decodeIndices(d binary.Decoder, indicesType GLenum) ([]index, error) {
	indices := make([]index, 0)
	switch indicesType {
	case GLenum_GL_UNSIGNED_BYTE:
		for {
			if val, err := d.Uint8(); err == nil {
				indices = append(indices, index(val))
			} else {
				return indices, nil
			}
		}

	case GLenum_GL_UNSIGNED_SHORT:
		for {
			if val, err := d.Uint16(); err == nil {
				indices = append(indices, index(val))
			} else {
				return indices, nil
			}
		}

	case GLenum_GL_UNSIGNED_INT:
		for {
			if val, err := d.Uint32(); err == nil {
				indices = append(indices, index(val))
			} else {
				return indices, nil
			}
		}

	default:
		return nil, fmt.Errorf("Invalid index type: %v", indicesType)
	}
}

func encodeIndices(indices []index) ([]byte, GLenum) {
	maxIndex := index(0)
	for _, v := range indices {
		if v > maxIndex {
			maxIndex = v
		}
	}
	buf := &bytes.Buffer{}
	enc := flat.Encoder(endian.Writer(buf, endian.Little))
	switch {
	case maxIndex > 0xFFFF:
		// TODO: GL_UNSIGNED_INT in glDrawElements is supported only since GLES 3.0
		for _, v := range indices {
			enc.Uint32(uint32(v))
		}
		return buf.Bytes(), GLenum_GL_UNSIGNED_INT

	case maxIndex > 0xFF:
		for _, v := range indices {
			enc.Uint16(uint16(v))
		}
		return buf.Bytes(), GLenum_GL_UNSIGNED_SHORT

	default:
		for _, v := range indices {
			enc.Uint8(uint8(v))
		}
		return buf.Bytes(), GLenum_GL_UNSIGNED_BYTE
	}
}

// Get the effective index buffer and primitive type for draw call
func getIndices(
	a atom.Atom,
	c *Context,
	s *gfxapi.State,
	d database.Database,
	l log.Logger) ([]index, GLenum, error) {

	switch a := a.(type) {
	case *GlDrawArrays:
		indices := make([]index, a.IndexCount)
		for i := range indices {
			indices[i] = index(a.FirstIndex) + index(i)
		}
		return indices, a.DrawMode, nil

	case *GlDrawElements:
		indexSize := map[GLenum]uint64{
			GLenum_GL_UNSIGNED_BYTE:  1,
			GLenum_GL_UNSIGNED_SHORT: 2,
			GLenum_GL_UNSIGNED_INT:   4,
		}[a.IndicesType]
		indexBufferID := c.BoundBuffers[GLenum_GL_ELEMENT_ARRAY_BUFFER]
		size := uint64(a.ElementCount) * indexSize

		var decoder binary.Decoder
		if indexBufferID == 0 {
			// Get the index buffer data from pointer
			decoder = a.Indices.Slice(0, size, s).Decoder(s, d, l)
		} else {
			// Get the index buffer data from buffer, offset by the 'indices' pointer.
			indexBuffer := c.Instances.Buffers[indexBufferID]
			if indexBuffer == nil {
				return nil, 0, fmt.Errorf("Can not find buffer %v", indexBufferID)
			}
			offset := uint64(a.Indices.Address)
			decoder = indexBuffer.Data.Slice(offset, offset+size, s).Decoder(s, d, l)
		}

		indices, err := decodeIndices(decoder, a.IndicesType)
		return indices, a.DrawMode, err

	default:
		return nil, 0, fmt.Errorf("Unknown draw command %v", a)
	}
}

func appendWireframeOfTriangle(lines []index, v0, v1, v2 index) []index {
	if v0 == v1 || v1 == v2 || v2 == v0 {
		return lines // Ignore degenerate triangle
	} else {
		return append(lines, v0, v1, v1, v2, v2, v0)
	}
}

func makeWireframe(indices []index, drawMode GLenum) ([]index, GLenum, error) {
	switch drawMode {
	case GLenum_GL_POINTS, GLenum_GL_LINES, GLenum_GL_LINE_STRIP, GLenum_GL_LINE_LOOP:
		return indices, drawMode, nil

	case GLenum_GL_TRIANGLES:
		numTriangles := len(indices) / 3
		lines := make([]index, 0, numTriangles*6)
		for i := 0; i < numTriangles; i++ {
			lines = appendWireframeOfTriangle(lines, indices[i*3], indices[i*3+1], indices[i*3+2])
		}
		return lines, GLenum_GL_LINES, nil

	case GLenum_GL_TRIANGLE_STRIP:
		numTriangles := len(indices) - 2
		if numTriangles > 0 {
			lines := make([]index, 0, numTriangles*6)
			for i := 0; i < numTriangles; i++ {
				lines = appendWireframeOfTriangle(lines, indices[i], indices[i+1], indices[i+2])
			}
			return lines, GLenum_GL_LINES, nil
		}
		return []index{}, GLenum_GL_LINES, nil

	case GLenum_GL_TRIANGLE_FAN:
		numTriangles := len(indices) - 2
		if numTriangles > 0 {
			lines := make([]index, 0, numTriangles*6)
			for i := 0; i < numTriangles; i++ {
				lines = appendWireframeOfTriangle(lines, indices[0], indices[i+1], indices[i+2])
			}
			return lines, GLenum_GL_LINES, nil
		}
		return []index{}, GLenum_GL_LINES, nil

	default:
		return nil, 0, fmt.Errorf("Unknown mode: %v", drawMode)
	}
}
