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
	"fmt"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary/endian"
	"android.googlesource.com/platform/tools/gpu/builder"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	rb "android.googlesource.com/platform/tools/gpu/replay/builder"
)

type externs struct {
	a atom.Atom
	s *gfxapi.State
	d database.Database
	l log.Logger
	b *rb.Builder
}

func (e externs) mapMemory(slice Slice) {
	if b := e.b; b != nil {
		switch e.a.(type) {
		case *GlMapBufferRange:
			// Base address is on the stack.
			b.MapMemory(slice.Range(e.s))

		default:
			log.E(e.l, "mapBuffer extern called for unsupported atom %T", e.a)
		}
	}
}

func (e externs) unmapMemory(slice Slice) {
	if b := e.b; b != nil {
		b.UnmapMemory(slice.Range(e.s))
	}
}

func (e externs) elSize(ty GLenum) uint64 {
	return uint64(DataTypeSize(ty))
}

func (e externs) calcIndexLimits(data U8ᵖ, ty GLenum, offset, count uint32) builder.IndexLimits {
	elSize := e.elSize(ty)
	id := data.Slice(uint64(offset), uint64(offset)+uint64(count)*elSize, e.s).ResourceID(e.s, e.d, e.l)
	littleEndian := e.s.Architecture.ByteOrder == endian.Little
	limits, err := builder.CalcIndexLimits(id, int(count), int(elSize), littleEndian, e.d, e.l)
	if err != nil {
		panic(fmt.Errorf("Could not calculate index limits: %v", err))
	}
	return limits
}

func (e externs) minIndex(data U8ᵖ, ty GLenum, offset, count uint32) uint32 {
	return e.calcIndexLimits(data, ty, offset, count).Min
}

func (e externs) maxIndex(data U8ᵖ, ty GLenum, offset, count uint32) uint32 {
	return e.calcIndexLimits(data, ty, offset, count).Max
}

func (e externs) substr(str string, start, end int32) string {
	return str[start:end]
}

func (e externs) pixelSize(f GLenum, ty GLenum) uint32 {
	return pixelSize(f, ty)
}

func (e externs) stateVariableSize(v GLenum) int32 {
	switch v {
	case GLenum_GL_ACTIVE_TEXTURE:
		return 1
	case GLenum_GL_ALIASED_LINE_WIDTH_RANGE:
		return 2
	case GLenum_GL_ALIASED_POINT_SIZE_RANGE:
		return 2
	case GLenum_GL_ALPHA_BITS:
		return 1
	case GLenum_GL_ARRAY_BUFFER_BINDING:
		return 1
	case GLenum_GL_BLEND:
		return 1
	case GLenum_GL_BLEND_COLOR:
		return 4
	case GLenum_GL_BLEND_DST_ALPHA:
		return 1
	case GLenum_GL_BLEND_DST_RGB:
		return 1
	case GLenum_GL_BLEND_EQUATION_ALPHA:
		return 1
	case GLenum_GL_BLEND_EQUATION_RGB:
		return 1
	case GLenum_GL_BLEND_SRC_ALPHA:
		return 1
	case GLenum_GL_BLEND_SRC_RGB:
		return 1
	case GLenum_GL_BLUE_BITS:
		return 1
	case GLenum_GL_COLOR_CLEAR_VALUE:
		return 4
	case GLenum_GL_COLOR_WRITEMASK:
		return 4
	case GLenum_GL_CULL_FACE:
		return 1
	case GLenum_GL_CULL_FACE_MODE:
		return 1
	case GLenum_GL_CURRENT_PROGRAM:
		return 1
	case GLenum_GL_DEPTH_BITS:
		return 1
	case GLenum_GL_DEPTH_CLEAR_VALUE:
		return 1
	case GLenum_GL_DEPTH_FUNC:
		return 1
	case GLenum_GL_DEPTH_RANGE:
		return 2
	case GLenum_GL_DEPTH_TEST:
		return 1
	case GLenum_GL_DEPTH_WRITEMASK:
		return 1
	case GLenum_GL_DITHER:
		return 1
	case GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING:
		return 1
	case GLenum_GL_FRAMEBUFFER_BINDING:
		return 1
	case GLenum_GL_FRONT_FACE:
		return 1
	case GLenum_GL_GENERATE_MIPMAP_HINT:
		return 1
	case GLenum_GL_GREEN_BITS:
		return 1
	case GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT:
		return 1
	case GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE:
		return 1
	case GLenum_GL_LINE_WIDTH:
		return 1
	case GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS:
		return 1
	case GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE:
		return 1
	case GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS:
		return 1
	case GLenum_GL_MAX_RENDERBUFFER_SIZE:
		return 1
	case GLenum_GL_MAX_TEXTURE_IMAGE_UNITS:
		return 1
	case GLenum_GL_MAX_TEXTURE_SIZE:
		return 1
	case GLenum_GL_MAX_VARYING_VECTORS:
		return 1
	case GLenum_GL_MAX_VERTEX_ATTRIBS:
		return 1
	case GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS:
		return 1
	case GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS:
		return 1
	case GLenum_GL_MAX_VIEWPORT_DIMS:
		return 2
	case GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS:
		return 1
	case GLenum_GL_NUM_SHADER_BINARY_FORMATS:
		return 1
	case GLenum_GL_PACK_ALIGNMENT:
		return 1
	case GLenum_GL_POLYGON_OFFSET_FACTOR:
		return 1
	case GLenum_GL_POLYGON_OFFSET_FILL:
		return 1
	case GLenum_GL_POLYGON_OFFSET_UNITS:
		return 1
	case GLenum_GL_RED_BITS:
		return 1
	case GLenum_GL_RENDERBUFFER_BINDING:
		return 1
	case GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE:
		return 1
	case GLenum_GL_SAMPLE_BUFFERS:
		return 1
	case GLenum_GL_SAMPLE_COVERAGE:
		return 1
	case GLenum_GL_SAMPLE_COVERAGE_INVERT:
		return 1
	case GLenum_GL_SAMPLE_COVERAGE_VALUE:
		return 1
	case GLenum_GL_SAMPLES:
		return 1
	case GLenum_GL_MAX_SAMPLES:
		return 1
	case GLenum_GL_NUM_PROGRAM_BINARY_FORMATS:
		return 1
	case GLenum_GL_MAX_COLOR_ATTACHMENTS:
		return 1
	case GLenum_GL_MAX_UNIFORM_BLOCK_SIZE:
		return 1
	case GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS:
		return 1
	case GLenum_GL_SCISSOR_BOX:
		return 4
	case GLenum_GL_SCISSOR_TEST:
		return 1
	case GLenum_GL_SHADER_COMPILER:
		return 1
	case GLenum_GL_STENCIL_BACK_FAIL:
		return 1
	case GLenum_GL_STENCIL_BACK_FUNC:
		return 1
	case GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL:
		return 1
	case GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS:
		return 1
	case GLenum_GL_STENCIL_BACK_REF:
		return 1
	case GLenum_GL_STENCIL_BACK_VALUE_MASK:
		return 1
	case GLenum_GL_STENCIL_BACK_WRITEMASK:
		return 1
	case GLenum_GL_STENCIL_BITS:
		return 1
	case GLenum_GL_STENCIL_CLEAR_VALUE:
		return 1
	case GLenum_GL_STENCIL_FAIL:
		return 1
	case GLenum_GL_STENCIL_FUNC:
		return 1
	case GLenum_GL_STENCIL_PASS_DEPTH_FAIL:
		return 1
	case GLenum_GL_STENCIL_PASS_DEPTH_PASS:
		return 1
	case GLenum_GL_STENCIL_REF:
		return 1
	case GLenum_GL_STENCIL_TEST:
		return 1
	case GLenum_GL_STENCIL_VALUE_MASK:
		return 1
	case GLenum_GL_STENCIL_WRITEMASK:
		return 1
	case GLenum_GL_SUBPIXEL_BITS:
		return 1
	case GLenum_GL_TEXTURE_BINDING_2D:
		return 1
	case GLenum_GL_TEXTURE_BINDING_CUBE_MAP:
		return 1
	case GLenum_GL_UNPACK_ALIGNMENT:
		return 1
	case GLenum_GL_VIEWPORT:
		return 4
	case GLenum_GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT:
		return 1
	case GLenum_GL_GPU_DISJOINT_EXT:
		return 1
	default:
		fmt.Printf("Warning: Unknown size for state variable: %s\n", v.String())
		return 0
	}
}

func (e externs) resolveAttributesAndUniforms(p *Program, pid ProgramId) {
	for _, extra := range e.a.Extras() {
		if pi, ok := extra.(*ProgramInfo); ok {
			for l, a := range pi.Attributes {
				p.Attributes[l] = VertexAttribute{Name: a.Name, VectorCount: a.VectorCount, Type: a.Type}
				if e.b != nil {
					// Force the attribute to use the capture-observed location for replay.
					NewGlBindAttribLocation(pid, l, a.Name).Replay(atom.NoID, e.s, e.d, e.l, e.b)
				}
			}
			for l, u := range pi.Uniforms {
				p.Uniforms[l] = Uniform{Name: u.Name, VectorCount: u.VectorCount, Type: u.Type}
			}
			return
		}
	}
	log.W(e.l, "%T missing ProgramInfo extra. Found extras:", e.a)
	for _, extra := range e.a.Extras() {
		log.W(e.l, "%T", extra)
	}
}
