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

	"android.googlesource.com/platform/tools/gpu/image"
)

// pixelSize returns the pixel data size in bytes for the given format and type.
func pixelSize(format GLenum, ty GLenum) uint32 {
	num_components := uint32(0)
	switch format {
	case GLenum_GL_ALPHA:
		num_components = 1
	case GLenum_GL_LUMINANCE:
		num_components = 1
	case GLenum_GL_LUMINANCE_ALPHA:
		num_components = 2
	case GLenum_GL_RED:
		num_components = 1
	case GLenum_GL_RED_INTEGER:
		num_components = 1
	case GLenum_GL_RG:
		num_components = 2
	case GLenum_GL_RG_INTEGER:
		num_components = 2
	case GLenum_GL_RGB:
		num_components = 3
	case GLenum_GL_RGB_INTEGER:
		num_components = 3
	case GLenum_GL_RGBA:
		num_components = 4
	case GLenum_GL_RGBA_INTEGER:
		num_components = 4
	case GLenum_GL_DEPTH_COMPONENT:
		num_components = 1
	case GLenum_GL_DEPTH_STENCIL:
		num_components = 2
	case GLenum_GL_STENCIL_INDEX:
		num_components = 1
	default:
		panic(fmt.Errorf("Unsupported image format: %v", format))
	}

	switch ty {
	case GLenum_GL_UNSIGNED_BYTE:
		return num_components * 1
	case GLenum_GL_BYTE:
		return num_components * 1
	case GLenum_GL_UNSIGNED_SHORT:
		return num_components * 2
	case GLenum_GL_SHORT:
		return num_components * 2
	case GLenum_GL_UNSIGNED_INT:
		return num_components * 4
	case GLenum_GL_INT:
		return num_components * 4
	case GLenum_GL_HALF_FLOAT:
		return num_components * 2
	case GLenum_GL_HALF_FLOAT_OES:
		return num_components * 2
	case GLenum_GL_FLOAT:
		return num_components * 4
	case GLenum_GL_UNSIGNED_SHORT_5_6_5:
		return 2
	case GLenum_GL_UNSIGNED_SHORT_4_4_4_4:
		return 2
	case GLenum_GL_UNSIGNED_SHORT_5_5_5_1:
		return 2
	case GLenum_GL_UNSIGNED_INT_2_10_10_10_REV:
		return 4
	case GLenum_GL_UNSIGNED_INT_10F_11F_11F_REV:
		return 4
	case GLenum_GL_UNSIGNED_INT_5_9_9_9_REV:
		return 4
	case GLenum_GL_UNSIGNED_INT_24_8:
		return 4
	case GLenum_GL_FLOAT_32_UNSIGNED_INT_24_8_REV:
		return 8
	default:
		panic(fmt.Errorf("Unsupported image type: %v", ty))
	}
}

// imageFormat returns the package image format for the given GL format and type.
func imageFormat(format GLenum, ty GLenum) image.Format {
	switch format {
	case GLenum_GL_ALPHA:
		switch ty {
		case GLenum_GL_UNSIGNED_BYTE:
			return image.Alpha()
		}
	case GLenum_GL_LUMINANCE:
		switch ty {
		case GLenum_GL_UNSIGNED_BYTE:
			return image.Luminance()
		}
	case GLenum_GL_LUMINANCE_ALPHA:
		switch ty {
		case GLenum_GL_UNSIGNED_BYTE:
			return image.LuminanceAlpha()
		}
	case GLenum_GL_RGB:
		switch ty {
		case GLenum_GL_UNSIGNED_BYTE:
			return image.RGB()
		case GLenum_GL_UNSIGNED_SHORT_5_6_5:
			return image.RGB565()
		}
	case GLenum_GL_RGBA:
		switch ty {
		case GLenum_GL_UNSIGNED_BYTE:
			return image.RGBA()
		case GLenum_GL_UNSIGNED_SHORT_5_5_5_1:
			return image.RGBA5551()
		case GLenum_GL_UNSIGNED_SHORT_4_4_4_4:
			return image.RGBA4444()
		}
	case GLenum_GL_ATC_RGB_AMD:
		return image.ATC_RGB_AMD()
	case GLenum_GL_ATC_RGBA_EXPLICIT_ALPHA_AMD:
		return image.ATC_RGBA_EXPLICIT_ALPHA_AMD()
	case GLenum_GL_ETC1_RGB8_OES:
		return image.ETC1_RGB8()
	case GLenum_GL_COMPRESSED_RGB8_ETC2:
		return image.ETC2_RGB8()
	case GLenum_GL_COMPRESSED_RGBA8_ETC2_EAC:
		return image.ETC2_RGBA8_EAC()
	}
	panic(fmt.Errorf("Unsupported input format-type pair: (%s, %s)", format, ty))
}
