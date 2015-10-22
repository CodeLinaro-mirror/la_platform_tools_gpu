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
 *
 */

#include "gles_constants.h"

#include "gles_types.h"

namespace gapii {

using namespace GLenum;

const ConstantDesc kGLESConstantDescs[] = {
    { GL_SUBPIXEL_BITS, GL_INT, true },
    { GL_MAX_ELEMENT_INDEX, GL_INT64_ARB, true },
    { GL_MAX_3D_TEXTURE_SIZE, GL_INT, true },
    { GL_MAX_TEXTURE_SIZE, GL_INT, true },
    { GL_MAX_ARRAY_TEXTURE_LAYERS, GL_INT, true },
    { GL_MAX_TEXTURE_LOD_BIAS, GL_FLOAT, true },
    { GL_MAX_CUBE_MAP_TEXTURE_SIZE, GL_INT, true },
    { GL_MAX_RENDERBUFFER_SIZE, GL_INT, true },
    { GL_ALIASED_POINT_SIZE_RANGE, GL_FLOAT, false },
    { GL_ALIASED_LINE_WIDTH_RANGE, GL_FLOAT, false },
    { GL_MULTISAMPLE_LINE_WIDTH_RANGE, GL_FLOAT, false },
    { GL_MULTISAMPLE_LINE_WIDTH_GRANULARITY, GL_FLOAT, true },
    { GL_MAX_DRAW_BUFFERS, GL_INT, true },
    { GL_MAX_FRAMEBUFFER_WIDTH, GL_INT, true },
    { GL_MAX_FRAMEBUFFER_HEIGHT, GL_INT, true },
    { GL_MAX_FRAMEBUFFER_LAYERS, GL_INT, true },
    { GL_MAX_FRAMEBUFFER_SAMPLES, GL_INT, true },
    { GL_MAX_COLOR_ATTACHMENTS, GL_INT, true },
    { GL_MIN_FRAGMENT_INTERPOLATION_OFFSET, GL_FLOAT, true },
    { GL_MAX_FRAGMENT_INTERPOLATION_OFFSET, GL_FLOAT, true },
    { GL_FRAGMENT_INTERPOLATION_OFFSET_BITS, GL_INT, true },
    { GL_MAX_VIEWPORT_DIMS, GL_INT, false },
    { GL_MAX_SAMPLE_MASK_WORDS, GL_INT, true },
    { GL_MAX_COLOR_TEXTURE_SAMPLES, GL_INT, true },
    { GL_MAX_DEPTH_TEXTURE_SAMPLES, GL_INT, true },
    { GL_MAX_INTEGER_SAMPLES, GL_INT, true },
    { GL_MAX_SERVER_WAIT_TIMEOUT, GL_INT64_ARB, true },
    { GL_LAYER_PROVOKING_VERTEX, GL_INT, true },
    { GL_PRIMITIVE_RESTART_FOR_PATCHES_SUPPORTED, GL_BOOL, true },
    { GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GL_INT, true },
    { GL_MAX_VERTEX_ATTRIB_BINDINGS, GL_INT, true },
    { GL_MAX_VERTEX_ATTRIB_STRIDE, GL_INT, true },
    { GL_MAX_ELEMENTS_INDICES, GL_INT, true },
    { GL_MAX_ELEMENTS_VERTICES, GL_INT, true },
    { GL_MAX_TEXTURE_BUFFER_SIZE, GL_INT, true },
    { GL_NUM_COMPRESSED_TEXTURE_FORMATS, GL_INT, true },
    { GL_COMPRESSED_TEXTURE_FORMATS, GL_INT, false },
    { GL_NUM_PROGRAM_BINARY_FORMATS, GL_INT, true },
    { GL_PROGRAM_BINARY_FORMATS, GL_INT, false },
    { GL_NUM_SHADER_BINARY_FORMATS, GL_INT, true },
    { GL_SHADER_BINARY_FORMATS, GL_INT, false },
    { GL_SHADER_COMPILER, GL_BOOL, true },
    { GL_TEXTURE_BUFFER_OFFSET_ALIGNMENT, GL_INT, true },
    { GL_NUM_EXTENSIONS, GL_INT, true },
    { GL_MAJOR_VERSION, GL_INT, true },
    { GL_MINOR_VERSION, GL_INT, true },
    { GL_CONTEXT_FLAGS, GL_INT, true },
    { GL_MAX_VERTEX_ATTRIBS, GL_INT, true },
    { GL_MAX_VERTEX_UNIFORM_COMPONENTS, GL_INT, true },
    { GL_MAX_VERTEX_UNIFORM_VECTORS, GL_INT, true },
    { GL_MAX_VERTEX_UNIFORM_BLOCKS, GL_INT, true },
    { GL_MAX_VERTEX_OUTPUT_COMPONENTS, GL_INT, true },
    { GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GL_INT, true },
    { GL_MAX_VERTEX_ATOMIC_COUNTER_BUFFERS, GL_INT, true },
    { GL_MAX_VERTEX_ATOMIC_COUNTERS, GL_INT, true },
    { GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GL_INT, true },
    { GL_MAX_TESS_GEN_LEVEL, GL_INT, true },
    { GL_MAX_PATCH_VERTICES, GL_INT, true },
    { GL_MAX_TESS_CONTROL_UNIFORM_COMPONENTS, GL_INT, true },
    { GL_MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS, GL_INT, true },
    { GL_MAX_TESS_CONTROL_OUTPUT_COMPONENTS, GL_INT, true },
    { GL_MAX_TESS_PATCH_COMPONENTS, GL_INT, true },
    { GL_MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS, GL_INT, true },
    { GL_MAX_TESS_CONTROL_INPUT_COMPONENTS, GL_INT, true },
    { GL_MAX_TESS_CONTROL_UNIFORM_BLOCKS, GL_INT, true },
    { GL_MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS, GL_INT, true },
    { GL_MAX_TESS_CONTROL_ATOMIC_COUNTERS, GL_INT, true },
    { GL_MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS, GL_INT, true },
    { GL_MAX_TESS_EVALUATION_UNIFORM_COMPONENTS, GL_INT, true },
    { GL_MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS, GL_INT, true },
    { GL_MAX_TESS_EVALUATION_OUTPUT_COMPONENTS, GL_INT, true },
    { GL_MAX_TESS_EVALUATION_INPUT_COMPONENTS, GL_INT, true },
    { GL_MAX_TESS_EVALUATION_UNIFORM_BLOCKS, GL_INT, true },
    { GL_MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS, GL_INT, true },
    { GL_MAX_TESS_EVALUATION_ATOMIC_COUNTERS, GL_INT, true },
    { GL_MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS, GL_INT, true },
    { GL_MAX_GEOMETRY_UNIFORM_COMPONENTS, GL_INT, true },
    { GL_MAX_GEOMETRY_UNIFORM_BLOCKS, GL_INT, true },
    { GL_MAX_GEOMETRY_INPUT_COMPONENTS, GL_INT, true },
    { GL_MAX_GEOMETRY_OUTPUT_COMPONENTS, GL_INT, true },
    { GL_MAX_GEOMETRY_OUTPUT_VERTICES, GL_INT, true },
    { GL_MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS, GL_INT, true },
    { GL_MAX_GEOMETRY_TEXTURE_IMAGE_UNITS, GL_INT, true },
    { GL_MAX_GEOMETRY_SHADER_INVOCATIONS, GL_INT, true },
    { GL_MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS, GL_INT, true },
    { GL_MAX_GEOMETRY_ATOMIC_COUNTERS, GL_INT, true },
    { GL_MAX_GEOMETRY_SHADER_STORAGE_BLOCKS, GL_INT, true },
    { GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GL_INT, true },
    { GL_MAX_FRAGMENT_UNIFORM_VECTORS, GL_INT, true },
    { GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GL_INT, true },
    { GL_MAX_FRAGMENT_INPUT_COMPONENTS, GL_INT, true },
    { GL_MAX_TEXTURE_IMAGE_UNITS, GL_INT, true },
    { GL_MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS, GL_INT, true },
    { GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GL_INT, true },
    { GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GL_INT, true },
    { GL_MIN_PROGRAM_TEXTURE_GATHER_OFFSET, GL_INT, true },
    { GL_MAX_PROGRAM_TEXTURE_GATHER_OFFSET, GL_INT, true },
    { GL_MIN_PROGRAM_TEXEL_OFFSET, GL_INT, true },
    { GL_MAX_PROGRAM_TEXEL_OFFSET, GL_INT, true },
    { GL_MAX_COMPUTE_WORK_GROUP_COUNT, GL_INT, false },
    { GL_MAX_COMPUTE_WORK_GROUP_SIZE, GL_INT, false },
    { GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GL_INT, true },
    { GL_MAX_COMPUTE_UNIFORM_BLOCKS, GL_INT, true },
    { GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GL_INT, true },
    { GL_MAX_COMPUTE_SHARED_MEMORY_SIZE, GL_INT, true },
    { GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GL_INT, true },
    { GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GL_INT, true },
    { GL_MAX_COMPUTE_ATOMIC_COUNTERS, GL_INT, true },
    { GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GL_INT, true },
    { GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GL_INT, true },
    { GL_MAX_UNIFORM_BUFFER_BINDINGS, GL_INT, true },
    { GL_MAX_UNIFORM_BLOCK_SIZE, GL_INT64_ARB, true },
    { GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GL_INT, true },
    { GL_MAX_COMBINED_UNIFORM_BLOCKS, GL_INT, true },
    { GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GL_INT64_ARB, true },
    { GL_MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS, GL_INT, true },
    { GL_MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS, GL_INT, true },
    { GL_MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS, GL_INT, true },
    { GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GL_INT64_ARB, true },
    { GL_MAX_VARYING_COMPONENTS, GL_INT, true },
    { GL_MAX_VARYING_VECTORS, GL_INT, true },
    { GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GL_INT, true },
    { GL_MAX_COMBINED_SHADER_OUTPUT_RESOURCES, GL_INT, true },
    { GL_MAX_UNIFORM_LOCATIONS, GL_INT, true },
    { GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GL_INT, true },
    { GL_MAX_ATOMIC_COUNTER_BUFFER_SIZE, GL_INT, true },
    { GL_MAX_COMBINED_ATOMIC_COUNTER_BUFFERS, GL_INT, true },
    { GL_MAX_COMBINED_ATOMIC_COUNTERS, GL_INT, true },
    { GL_MAX_IMAGE_UNITS, GL_INT, true },
    { GL_MAX_VERTEX_IMAGE_UNIFORMS, GL_INT, true },
    { GL_MAX_TESS_CONTROL_IMAGE_UNIFORMS, GL_INT, true },
    { GL_MAX_TESS_EVALUATION_IMAGE_UNIFORMS, GL_INT, true },
    { GL_MAX_GEOMETRY_IMAGE_UNIFORMS, GL_INT, true },
    { GL_MAX_FRAGMENT_IMAGE_UNIFORMS, GL_INT, true },
    { GL_MAX_COMPUTE_IMAGE_UNIFORMS, GL_INT, true },
    { GL_MAX_COMBINED_IMAGE_UNIFORMS, GL_INT, true },
    { GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GL_INT, true },
    { GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GL_INT64_ARB, true },
    { GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GL_INT, true },
    { GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GL_INT, true },
    { 0, 0, false },
};

int GetSizeOfGLType(uint32_t type) {
    switch(type) {
    case GL_BOOL:
        return sizeof(GLboolean);
    case GL_FLOAT:
        return sizeof(GLfloat);
    case GL_INT:
        return sizeof(GLint);
    case GL_INT64_ARB:
        return sizeof(GLint64);
    default:
        GAPID_FATAL("Unknown GL type %d for constant", int(type));
    }
}

// Helper method to prepare the output buffer.
// Returns typed pointer to the start of the buffer.
template<typename T>
static T* resizeBuffer(std::vector<uint8_t>* buffer, int component_count) {
    buffer->clear();
    buffer->resize(sizeof(T) * component_count, 0);
    return reinterpret_cast<T*>(buffer->data());
}

bool GetConstant(const GlesImports& imports,
                 bool is_gles_30_or_newer,
                 const ConstantDesc* desc,
                 std::vector<uint8_t>* values) {
    if (desc->scalar) {
        switch(desc->type) {
        case GL_BOOL:
            imports.glGetBooleanv(desc->name, resizeBuffer<GLboolean>(values, 1));
            break;
        case GL_FLOAT:
            imports.glGetFloatv(desc->name, resizeBuffer<GLfloat>(values, 1));
            break;
        case GL_INT:
            imports.glGetIntegerv(desc->name, resizeBuffer<GLint>(values, 1));
            break;
        case GL_INT64_ARB:
            if (!is_gles_30_or_newer) {
                return false;
            }
            imports.glGetInteger64v(desc->name, resizeBuffer<GLint64>(values, 1));
            break;
        default:
            return false;
        }
    } else {
        GLint num;
        switch(desc->name) {
        case GL_ALIASED_POINT_SIZE_RANGE:
        case GL_ALIASED_LINE_WIDTH_RANGE:
        case GL_MULTISAMPLE_LINE_WIDTH_RANGE:
            imports.glGetFloatv(desc->name, resizeBuffer<GLfloat>(values, 2));
            break;
        case GL_MAX_VIEWPORT_DIMS:
            imports.glGetIntegerv(desc->name, resizeBuffer<GLint>(values, 2));
            break;
        case GL_COMPRESSED_TEXTURE_FORMATS:
            imports.glGetIntegerv(GL_NUM_COMPRESSED_TEXTURE_FORMATS, &num);
            if (imports.glGetError() != GL_NO_ERROR) {
                return false;
            }
            imports.glGetIntegerv(desc->name, resizeBuffer<GLint>(values, num));
            break;
        case GL_PROGRAM_BINARY_FORMATS:
            imports.glGetIntegerv(GL_NUM_PROGRAM_BINARY_FORMATS, &num);
            if (imports.glGetError() != GL_NO_ERROR) {
                return false;
            }
            imports.glGetIntegerv(desc->name, resizeBuffer<GLint>(values, num));
            break;
        case GL_SHADER_BINARY_FORMATS:
            imports.glGetIntegerv(GL_NUM_SHADER_BINARY_FORMATS, &num);
            if (imports.glGetError() != GL_NO_ERROR) {
                return false;
            }
            imports.glGetIntegerv(desc->name, resizeBuffer<GLint>(values, num));
            break;
        case GL_MAX_COMPUTE_WORK_GROUP_COUNT: {
            if (!is_gles_30_or_newer) {
                return false;
            }
            GLint* dims = resizeBuffer<GLint>(values, 3);
            for (int i = 0; i < 3; i++) {
                imports.glGetIntegeri_v(GL_MAX_COMPUTE_WORK_GROUP_COUNT, i, &dims[i]);
            }
            break;
        }
        case GL_MAX_COMPUTE_WORK_GROUP_SIZE: {
            if (!is_gles_30_or_newer) {
                return false;
            }
            GLint* dims = resizeBuffer<GLint>(values, 3);
            for (int i = 0; i < 3; i++) {
                imports.glGetIntegeri_v(GL_MAX_COMPUTE_WORK_GROUP_SIZE, i, &dims[i]);
            }
            break;
        }
        default:
            return false;
        }
    }
    return (imports.glGetError() == GL_NO_ERROR);
}

}
