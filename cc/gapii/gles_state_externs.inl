/*
 * Copyright 2015, The Android Open Source Project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

template<typename T> T inline min(T a, T b) { return (a < b) ? a : b; }
template<typename T> T inline max(T a, T b) { return (a > b) ? a : b; }

template<typename T> inline void mapMemory(const T&) {}
template<typename T> inline void unmapMemory(const T&) {}

inline uint32_t minIndex(const uint8_t* indices, uint32_t indices_type, uint32_t offset, uint32_t count) {
    uint32_t v = ~(uint32_t)0;
    switch (indices_type) {
        case gapii::GLenum::GL_UNSIGNED_BYTE: {
            const uint8_t* p = reinterpret_cast<const uint8_t*>(&indices[offset]);
            for (uint32_t i = 0; i < count; i++) { v = min<uint32_t>(v, p[i]); }
            break;
        }
        case gapii::GLenum::GL_UNSIGNED_SHORT: {
            const uint16_t* p = reinterpret_cast<const uint16_t*>(&indices[offset]);
            for (uint32_t i = 0; i < count; i++) { v = min<uint32_t>(v, p[i]); }
            break;
        }
        case gapii::GLenum::GL_UNSIGNED_INT: {
            const uint32_t* p = reinterpret_cast<const uint32_t*>(&indices[offset]);
            for (uint32_t i = 0; i < count; i++) { v = min<uint32_t>(v, p[i]); }
            break;
        }
    }
    return v;
}

inline uint32_t maxIndex(const uint8_t* indices, uint32_t indices_type, uint32_t offset, uint32_t count) {
    uint32_t v = 0;
    switch (indices_type) {
        case gapii::GLenum::GL_UNSIGNED_BYTE: {
            const uint8_t* p = reinterpret_cast<const uint8_t*>(&indices[offset]);
            for (uint32_t i = 0; i < count; i++) { v = max<uint32_t>(v, p[i]); }
            break;
        }
        case gapii::GLenum::GL_UNSIGNED_SHORT: {
            const uint16_t* p = reinterpret_cast<const uint16_t*>(&indices[offset]);
            for (uint32_t i = 0; i < count; i++) { v = max<uint32_t>(v, p[i]); }
            break;
        }
        case gapii::GLenum::GL_UNSIGNED_INT: {
            const uint32_t* p = reinterpret_cast<const uint32_t*>(&indices[offset]);
            for (uint32_t i = 0; i < count; i++) { v = max<uint32_t>(v, p[i]); }
            break;
        }
    }
    return v;
}

inline int stateVariableSize(uint32_t v) {
    switch (v) {
    case GLenum::GL_ACTIVE_TEXTURE:
        return 1;
    case GLenum::GL_ALIASED_LINE_WIDTH_RANGE:
        return 2;
    case GLenum::GL_ALIASED_POINT_SIZE_RANGE:
        return 2;
    case GLenum::GL_ALPHA_BITS:
        return 1;
    case GLenum::GL_ARRAY_BUFFER_BINDING:
        return 1;
    case GLenum::GL_BLEND:
        return 1;
    case GLenum::GL_BLEND_COLOR:
        return 4;
    case GLenum::GL_BLEND_DST_ALPHA:
        return 1;
    case GLenum::GL_BLEND_DST_RGB:
        return 1;
    case GLenum::GL_BLEND_EQUATION_ALPHA:
        return 1;
    case GLenum::GL_BLEND_EQUATION_RGB:
        return 1;
    case GLenum::GL_BLEND_SRC_ALPHA:
        return 1;
    case GLenum::GL_BLEND_SRC_RGB:
        return 1;
    case GLenum::GL_BLUE_BITS:
        return 1;
    case GLenum::GL_COLOR_CLEAR_VALUE:
        return 4;
    case GLenum::GL_COLOR_WRITEMASK:
        return 4;
    case GLenum::GL_CULL_FACE:
        return 1;
    case GLenum::GL_CULL_FACE_MODE:
        return 1;
    case GLenum::GL_CURRENT_PROGRAM:
        return 1;
    case GLenum::GL_DEPTH_BITS:
        return 1;
    case GLenum::GL_DEPTH_CLEAR_VALUE:
        return 1;
    case GLenum::GL_DEPTH_FUNC:
        return 1;
    case GLenum::GL_DEPTH_RANGE:
        return 2;
    case GLenum::GL_DEPTH_TEST:
        return 1;
    case GLenum::GL_DEPTH_WRITEMASK:
        return 1;
    case GLenum::GL_DITHER:
        return 1;
    case GLenum::GL_ELEMENT_ARRAY_BUFFER_BINDING:
        return 1;
    case GLenum::GL_FRAMEBUFFER_BINDING:
        return 1;
    case GLenum::GL_FRONT_FACE:
        return 1;
    case GLenum::GL_GENERATE_MIPMAP_HINT:
        return 1;
    case GLenum::GL_GREEN_BITS:
        return 1;
    case GLenum::GL_IMPLEMENTATION_COLOR_READ_FORMAT:
        return 1;
    case GLenum::GL_IMPLEMENTATION_COLOR_READ_TYPE:
        return 1;
    case GLenum::GL_LINE_WIDTH:
        return 1;
    case GLenum::GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS:
        return 1;
    case GLenum::GL_MAX_CUBE_MAP_TEXTURE_SIZE:
        return 1;
    case GLenum::GL_MAX_FRAGMENT_UNIFORM_VECTORS:
        return 1;
    case GLenum::GL_MAX_RENDERBUFFER_SIZE:
        return 1;
    case GLenum::GL_MAX_TEXTURE_IMAGE_UNITS:
        return 1;
    case GLenum::GL_MAX_TEXTURE_SIZE:
        return 1;
    case GLenum::GL_MAX_VARYING_VECTORS:
        return 1;
    case GLenum::GL_MAX_VERTEX_ATTRIBS:
        return 1;
    case GLenum::GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS:
        return 1;
    case GLenum::GL_MAX_VERTEX_UNIFORM_VECTORS:
        return 1;
    case GLenum::GL_MAX_VIEWPORT_DIMS:
        return 2;
    case GLenum::GL_NUM_COMPRESSED_TEXTURE_FORMATS:
        return 1;
    case GLenum::GL_NUM_SHADER_BINARY_FORMATS:
        return 1;
    case GLenum::GL_PACK_ALIGNMENT:
        return 1;
    case GLenum::GL_POLYGON_OFFSET_FACTOR:
        return 1;
    case GLenum::GL_POLYGON_OFFSET_FILL:
        return 1;
    case GLenum::GL_POLYGON_OFFSET_UNITS:
        return 1;
    case GLenum::GL_RED_BITS:
        return 1;
    case GLenum::GL_RENDERBUFFER_BINDING:
        return 1;
    case GLenum::GL_SAMPLE_ALPHA_TO_COVERAGE:
        return 1;
    case GLenum::GL_SAMPLE_BUFFERS:
        return 1;
    case GLenum::GL_SAMPLE_COVERAGE:
        return 1;
    case GLenum::GL_SAMPLE_COVERAGE_INVERT:
        return 1;
    case GLenum::GL_SAMPLE_COVERAGE_VALUE:
        return 1;
    case GLenum::GL_SAMPLES:
        return 1;
    case GLenum::GL_MAX_SAMPLES:
        return 1;
    case GLenum::GL_NUM_PROGRAM_BINARY_FORMATS:
        return 1;
    case GLenum::GL_MAX_COLOR_ATTACHMENTS:
        return 1;
    case GLenum::GL_MAX_UNIFORM_BLOCK_SIZE:
        return 1;
    case GLenum::GL_MAX_VERTEX_UNIFORM_COMPONENTS:
        return 1;
    case GLenum::GL_SCISSOR_BOX:
        return 4;
    case GLenum::GL_SCISSOR_TEST:
        return 1;
    case GLenum::GL_SHADER_COMPILER:
        return 1;
    case GLenum::GL_STENCIL_BACK_FAIL:
        return 1;
    case GLenum::GL_STENCIL_BACK_FUNC:
        return 1;
    case GLenum::GL_STENCIL_BACK_PASS_DEPTH_FAIL:
        return 1;
    case GLenum::GL_STENCIL_BACK_PASS_DEPTH_PASS:
        return 1;
    case GLenum::GL_STENCIL_BACK_REF:
        return 1;
    case GLenum::GL_STENCIL_BACK_VALUE_MASK:
        return 1;
    case GLenum::GL_STENCIL_BACK_WRITEMASK:
        return 1;
    case GLenum::GL_STENCIL_BITS:
        return 1;
    case GLenum::GL_STENCIL_CLEAR_VALUE:
        return 1;
    case GLenum::GL_STENCIL_FAIL:
        return 1;
    case GLenum::GL_STENCIL_FUNC:
        return 1;
    case GLenum::GL_STENCIL_PASS_DEPTH_FAIL:
        return 1;
    case GLenum::GL_STENCIL_PASS_DEPTH_PASS:
        return 1;
    case GLenum::GL_STENCIL_REF:
        return 1;
    case GLenum::GL_STENCIL_TEST:
        return 1;
    case GLenum::GL_STENCIL_VALUE_MASK:
        return 1;
    case GLenum::GL_STENCIL_WRITEMASK:
        return 1;
    case GLenum::GL_SUBPIXEL_BITS:
        return 1;
    case GLenum::GL_TEXTURE_BINDING_2D:
        return 1;
    case GLenum::GL_TEXTURE_BINDING_CUBE_MAP:
        return 1;
    case GLenum::GL_UNPACK_ALIGNMENT:
        return 1;
    case GLenum::GL_VIEWPORT:
        return 4;
    case GLenum::GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT:
        return 1;
    case GLenum::GL_GPU_DISJOINT_EXT:
        return 1;
    case GLenum::GL_MAJOR_VERSION:
        return 1;
    case GLenum::GL_MINOR_VERSION:
        return 1;
    default:
        return 0; // TODO: Assert?
    }
}

int pixelSize(uint32_t format, uint32_t type) {
    int num_components = 0;
    switch (format) {
    case GLenum::GL_ALPHA:
        num_components = 1;
        break;
    case GLenum::GL_LUMINANCE:
        num_components = 1;
        break;
    case GLenum::GL_LUMINANCE_ALPHA:
        num_components = 2;
        break;
    case GLenum::GL_RED:
        num_components = 1;
        break;
    case GLenum::GL_RED_INTEGER:
        num_components = 1;
        break;
    case GLenum::GL_RG:
        num_components = 2;
        break;
    case GLenum::GL_RG_INTEGER:
        num_components = 2;
        break;
    case GLenum::GL_RGB:
        num_components = 3;
        break;
    case GLenum::GL_RGB_INTEGER:
        num_components = 3;
        break;
    case GLenum::GL_RGBA:
        num_components = 4;
        break;
    case GLenum::GL_RGBA_INTEGER:
        num_components = 4;
        break;
    case GLenum::GL_DEPTH_COMPONENT:
        num_components = 1;
        break;
    case GLenum::GL_DEPTH_STENCIL:
        num_components = 2;
        break;
    case GLenum::GL_STENCIL_INDEX:
        num_components = 1;
        break;
    }

    switch(type) {
    case GLenum::GL_UNSIGNED_BYTE:
        return num_components * 1;
    case GLenum::GL_BYTE:
        return num_components * 1;
    case GLenum::GL_UNSIGNED_SHORT:
        return num_components * 2;
    case GLenum::GL_SHORT:
        return num_components * 2;
    case GLenum::GL_UNSIGNED_INT:
        return num_components * 4;
    case GLenum::GL_INT:
        return num_components * 4;
    case GLenum::GL_HALF_FLOAT:
        return num_components * 2;
    case GLenum::GL_HALF_FLOAT_OES:
        return num_components * 2;
    case GLenum::GL_FLOAT:
        return num_components * 4;
    case GLenum::GL_UNSIGNED_SHORT_5_6_5:
        return 2;
    case GLenum::GL_UNSIGNED_SHORT_4_4_4_4:
        return 2;
    case GLenum::GL_UNSIGNED_SHORT_5_5_5_1:
        return 2;
    case GLenum::GL_UNSIGNED_INT_2_10_10_10_REV:
        return 4;
    case GLenum::GL_UNSIGNED_INT_10F_11F_11F_REV:
        return 4;
    case GLenum::GL_UNSIGNED_INT_5_9_9_9_REV:
        return 4;
    case GLenum::GL_UNSIGNED_INT_24_8:
        return 4;
    case GLenum::GL_FLOAT_32_UNSIGNED_INT_24_8_REV:
        return 8;
    default:
        return 0; // TODO: Assert?
    }
}
