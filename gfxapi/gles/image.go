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

import "fmt"

// imageSize returns the image pixel data size in bytes for the given
// width, height, format and type.
func imageSize(width, height uint32, f GLenum, ty GLenum) uint32 {
	// TODO: Consider ty
	switch f {
	case GLenum_GL_ALPHA:
		return width * height
	case GLenum_GL_LUMINANCE:
		return width * height
	case GLenum_GL_LUMINANCE_ALPHA:
		return 2 * width * height
	case GLenum_GL_RGB:
		return 3 * width * height
	case GLenum_GL_RGBA:
		return 4 * width * height
	case GLenum_GL_RED:
		return width * height
	case GLenum_GL_RED_INTEGER:
		return width * height
	case GLenum_GL_RG:
		return width * height * 2
	case GLenum_GL_RG_INTEGER:
		return width * height * 2
	case GLenum_GL_RGB_INTEGER:
		return width * height * 3
	case GLenum_GL_RGBA_INTEGER:
		return width * height * 3
	case GLenum_GL_DEPTH_COMPONENT:
		return width * height
	case GLenum_GL_DEPTH_STENCIL:
		return width * height
	default:
		panic(fmt.Errorf("Unsupported image format: %v", f))
	}
}
