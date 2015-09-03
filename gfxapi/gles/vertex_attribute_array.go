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

	"android.googlesource.com/platform/tools/gpu/memory"
)

// DataTypeSize returns the size in bytes of the the specified data type.
func DataTypeSize(t GLenum) int {
	switch t {
	case GLenum_GL_BYTE:
		return 1
	case GLenum_GL_UNSIGNED_BYTE:
		return 1
	case GLenum_GL_SHORT:
		return 2
	case GLenum_GL_UNSIGNED_SHORT:
		return 2
	case GLenum_GL_HALF_FLOAT_ARB:
		return 2
	case GLenum_GL_HALF_FLOAT_OES:
		return 2
	case GLenum_GL_FIXED:
		return 4
	case GLenum_GL_FLOAT:
		return 4
	case GLenum_GL_UNSIGNED_INT:
		return 4
	default:
		panic(fmt.Errorf("Unknown data type %v", t))
	}
}

// MemoryRange returns the memory range for the specified vertex index range,
// either in client memory, or if v.Buffer != 0 an offset within the associated
// buffer's memory.
func (v VertexAttributeArray) MemoryRange(first, last int) memory.Range {
	size := DataTypeSize(v.Type) * int(v.Size)
	stride := int(v.Stride)
	if stride == 0 {
		stride = size
	}
	count := (last - first) + 1
	return v.Pointer.Range(uint64(size + (count-1)*stride))
}
