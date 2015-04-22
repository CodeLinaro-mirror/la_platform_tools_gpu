package gles

import "fmt"

// imageSize returns the image pixel data size in bytes for the given
// width, height, format and type.
func imageSize(width, height uint32, f TexelFormat, ty TexelType) uint32 {
	// TODO: Consider ty
	switch f {
	// TexelFormat_GLES1
	case TexelFormat_GL_ALPHA:
		return width * height
	case TexelFormat_GL_LUMINANCE:
		return width * height
	case TexelFormat_GL_LUMINANCE_ALPHA:
		return 2 * width * height
	case TexelFormat_GL_RGB:
		return 3 * width * height
	case TexelFormat_GL_RGBA:
		return 4 * width * height
	// TexelFormat_GLES3
	case TexelFormat_GL_RED:
		return width * height
	case TexelFormat_GL_RED_INTEGER:
		return width * height
	case TexelFormat_GL_RG:
		return width * height * 2
	case TexelFormat_GL_RG_INTEGER:
		return width * height * 2
	case TexelFormat_GL_RGB_INTEGER:
		return width * height * 3
	case TexelFormat_GL_RGBA_INTEGER:
		return width * height * 3
	case TexelFormat_GL_DEPTH_COMPONENT:
		return width * height
	case TexelFormat_GL_DEPTH_STENCIL:
		return width * height
	default:
		panic(fmt.Errorf("Unsupported image format: %v", f))
	}
}
