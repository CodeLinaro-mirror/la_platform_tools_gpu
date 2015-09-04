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
	"strings"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/image"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
)

// decompressTexImage2D writes a glTexImage2D using the decompressed data for
// the given glCompressedTexImage2D.
func decompressTexImage2D(i atom.ID, a *GlCompressedTexImage2D, s *gfxapi.State, d database.Database, l log.Logger, out atom.Writer) error {
	l = log.Enter(l, "decompressTexImage2D")
	c := getContext(s)

	data := a.Data
	if pb := c.BoundBuffers[GLenum_GL_PIXEL_UNPACK_BUFFER]; pb != 0 {
		base := a.Data.Pointer.Address
		data = TexturePointer(c.Instances.Buffers[pb].Data.Index(base, s))
		out.Write(atom.NoID, NewGlBindBuffer(GLenum_GL_PIXEL_UNPACK_BUFFER, 0))
		defer out.Write(atom.NoID, NewGlBindBuffer(GLenum_GL_PIXEL_UNPACK_BUFFER, pb))
	} else {
		a.Observations().ApplyReads(s.Memory[memory.ApplicationPool])
	}

	id, err := database.Store(&image.LazyConverter{
		Data:       data.Slice(0, uint64(a.ImageSize), s).ResourceID(s, d, l),
		Width:      uint32(a.Width),
		Height:     uint32(a.Height),
		FormatFrom: imageFormat(a.Format),
		FormatTo:   image.RGBA(),
	}, d, l)
	if err != nil {
		return err
	}

	dstSize := a.Width * a.Height * 4

	out.Write(i, NewGlTexImage2D(
		a.Target,
		a.Level,
		GLint(GLenum_GL_RGBA),
		a.Width,
		a.Height,
		a.Border,
		GLenum_GL_RGBA,
		GLenum_GL_UNSIGNED_BYTE,
		memory.Tmp,
	).AddRead(memory.Tmp.Range(uint64(dstSize)), id))

	return nil
}

// convertTexImage2D writes the glTexImage2D with the format of changed to RGBA.
func convertTexImage2D(i atom.ID, a *GlTexImage2D, s *gfxapi.State, d database.Database, l log.Logger, out atom.Writer) error {
	l = log.Enter(l, "convertTexImage2D")
	c := getContext(s)

	url := c.PixelStorage[GLenum_GL_UNPACK_ROW_LENGTH]
	srcStridePixels := int(url)
	if srcStridePixels == 0 {
		srcStridePixels = int(a.Width)
	}

	data := a.Data
	if pb := c.BoundBuffers[GLenum_GL_PIXEL_UNPACK_BUFFER]; pb != 0 {
		base := a.Data.Pointer.Address
		data = TexturePointer(c.Instances.Buffers[pb].Data.Index(base, s))
		out.Write(atom.NoID, NewGlBindBuffer(GLenum_GL_PIXEL_UNPACK_BUFFER, 0))
		defer out.Write(atom.NoID, NewGlBindBuffer(GLenum_GL_PIXEL_UNPACK_BUFFER, pb))
	} else {
		if a.Data.Pointer == memory.Nullptr {
			out.Write(i, NewGlTexImage2D(
				a.Target,
				a.Level,
				GLint(GLenum_GL_RGBA),
				a.Width,
				a.Height,
				a.Border,
				GLenum_GL_RGBA,
				GLenum_GL_UNSIGNED_BYTE,
				memory.Nullptr,
			))
			return nil
		}
		a.Observations().ApplyReads(s.Memory[memory.ApplicationPool])
	}

	srcFmt := imageFormat(a.Format)
	srcSize := srcFmt.Size(srcStridePixels, int(a.Height))
	dstSize := a.Width * a.Height * 4

	id, err := database.Store(&image.LazyConverter{
		Data:       data.Slice(0, uint64(srcSize), s).ResourceID(s, d, l),
		Width:      uint32(a.Width),
		Height:     uint32(a.Height),
		StrideFrom: srcFmt.Size(srcStridePixels, 1),
		FormatFrom: srcFmt,
		FormatTo:   image.RGBA(),
	}, d, l)
	if err != nil {
		return err
	}

	out.Write(atom.NoID, NewGlPixelStorei(GLenum_GL_UNPACK_ROW_LENGTH, 0))
	out.Write(i, NewGlTexImage2D(
		a.Target,
		a.Level,
		GLint(GLenum_GL_RGBA),
		a.Width,
		a.Height,
		a.Border,
		GLenum_GL_RGBA,
		GLenum_GL_UNSIGNED_BYTE,
		memory.Tmp,
	).AddRead(memory.Tmp.Range(uint64(dstSize)), id))
	out.Write(atom.NoID, NewGlPixelStorei(GLenum_GL_UNPACK_ROW_LENGTH, url))

	return nil
}

// convertTexSubImage2D writes the glTexSubImage2D with the format of changed to RGBA.
func convertTexSubImage2D(i atom.ID, a *GlTexSubImage2D, s *gfxapi.State, d database.Database, l log.Logger, out atom.Writer) error {
	l = log.Enter(l, "convertTexSubImage2D")
	c := getContext(s)

	url := c.PixelStorage[GLenum_GL_UNPACK_ROW_LENGTH]
	srcStridePixels := int(url)
	if srcStridePixels == 0 {
		srcStridePixels = int(a.Width)
	}

	srcFmt := imageFormat(a.Format)
	srcSize := srcFmt.Size(srcStridePixels, int(a.Height))
	dstSize := int(a.Width) * int(a.Height) * 4

	data := a.Data
	if pb := c.BoundBuffers[GLenum_GL_PIXEL_UNPACK_BUFFER]; pb != 0 {
		base := a.Data.Pointer.Address
		data = TexturePointer(c.Instances.Buffers[pb].Data.Index(base, s))
		out.Write(atom.NoID, NewGlBindBuffer(GLenum_GL_PIXEL_UNPACK_BUFFER, 0))
		defer out.Write(atom.NoID, NewGlBindBuffer(GLenum_GL_PIXEL_UNPACK_BUFFER, pb))
	} else {
		a.Observations().ApplyReads(s.Memory[memory.ApplicationPool])
	}

	id, err := database.Store(&image.LazyConverter{
		Data:       data.Slice(0, uint64(srcSize), s).ResourceID(s, d, l),
		Width:      uint32(a.Width),
		Height:     uint32(a.Height),
		StrideFrom: srcFmt.Size(srcStridePixels, 1),
		FormatFrom: srcFmt,
		FormatTo:   image.RGBA(),
	}, d, l)
	if err != nil {
		return err
	}

	out.Write(atom.NoID, NewGlPixelStorei(GLenum_GL_UNPACK_ROW_LENGTH, 0))
	out.Write(i, NewGlTexSubImage2D(
		a.Target,
		a.Level,
		a.Xoffset,
		a.Yoffset,
		a.Width,
		a.Height,
		GLenum_GL_RGBA,
		GLenum_GL_UNSIGNED_BYTE,
		memory.Tmp,
	).AddRead(memory.Tmp.Range(uint64(dstSize)), id))
	out.Write(atom.NoID, NewGlPixelStorei(GLenum_GL_UNPACK_ROW_LENGTH, url))

	return nil
}

// getSupportedUncompressedTextureFormats returns the set of supported
// uncompressed texture formats for a given version and extension list.
func getSupportedUncompressedTextureFormats(version Version, extensions string) (map[GLenum]struct{}, error) {
	s := struct{}{}
	if version.IsES {
		switch { // OpenGL ES
		case version.Major >= 3:
			return map[GLenum]struct{}{
				GLenum_GL_RED:             s,
				GLenum_GL_RED_INTEGER:     s,
				GLenum_GL_RG:              s,
				GLenum_GL_RG_INTEGER:      s,
				GLenum_GL_RGB:             s,
				GLenum_GL_RGB_INTEGER:     s,
				GLenum_GL_RGBA:            s,
				GLenum_GL_RGBA_INTEGER:    s,
				GLenum_GL_DEPTH_COMPONENT: s,
				GLenum_GL_DEPTH_STENCIL:   s,
				GLenum_GL_LUMINANCE_ALPHA: s,
				GLenum_GL_LUMINANCE:       s,
				GLenum_GL_ALPHA:           s,
			}, nil

		case version.Major >= 2:
			return map[GLenum]struct{}{
				GLenum_GL_ALPHA:           s,
				GLenum_GL_RGB:             s,
				GLenum_GL_RGBA:            s,
				GLenum_GL_LUMINANCE:       s,
				GLenum_GL_LUMINANCE_ALPHA: s,
			}, nil
		}
	} else {
		switch { // OpenGL
		case version.Major >= 4:
			return map[GLenum]struct{}{
				GLenum_GL_RED:             s,
				GLenum_GL_RG:              s,
				GLenum_GL_RGB:             s,
				GLenum_GL_BGR:             s,
				GLenum_GL_RGBA:            s,
				GLenum_GL_BGRA:            s,
				GLenum_GL_RED_INTEGER:     s,
				GLenum_GL_RG_INTEGER:      s,
				GLenum_GL_RGB_INTEGER:     s,
				GLenum_GL_BGR_INTEGER:     s,
				GLenum_GL_RGBA_INTEGER:    s,
				GLenum_GL_BGRA_INTEGER:    s,
				GLenum_GL_STENCIL_INDEX:   s,
				GLenum_GL_DEPTH_COMPONENT: s,
				GLenum_GL_DEPTH_STENCIL:   s,
			}, nil

		case version.Major == 3 && version.Minor >= 3:
			return map[GLenum]struct{}{
				GLenum_GL_RED:             s,
				GLenum_GL_RG:              s,
				GLenum_GL_RGB:             s,
				GLenum_GL_BGR:             s,
				GLenum_GL_RGBA:            s,
				GLenum_GL_BGRA:            s,
				GLenum_GL_DEPTH_COMPONENT: s,
				GLenum_GL_DEPTH_STENCIL:   s,
			}, nil

		case version.Major >= 2:
			return map[GLenum]struct{}{
				GLenum_GL_COLOR_INDEX:     s,
				GLenum_GL_RED:             s,
				GLenum_GL_GREEN:           s,
				GLenum_GL_BLUE:            s,
				GLenum_GL_ALPHA:           s,
				GLenum_GL_RGB:             s,
				GLenum_GL_BGR:             s,
				GLenum_GL_RGBA:            s,
				GLenum_GL_BGRA:            s,
				GLenum_GL_LUMINANCE:       s,
				GLenum_GL_LUMINANCE_ALPHA: s,

				// HACK: Not officially supported, but is with most compatability modes.
				// Required for depth-buffer readback. Consider removing.
				GLenum_GL_DEPTH_STENCIL: s,
			}, nil
		}
	}

	return nil, fmt.Errorf("Unsupported version %+v", version)
}

// getSupportedCompressedTextureFormats returns the set of supported compressed
// texture formats for a given extension list.
func getSupportedCompressedTextureFormats(extensions string) map[GLenum]struct{} {
	supported := map[GLenum]struct{}{}
	for _, extension := range strings.Split(extensions, " ") {
		for _, format := range getExtensionTextureFormats(extension) {
			supported[format] = struct{}{}
		}
	}
	return supported
}

// getExtensionTextureFormats returns the list of compressed texture formats
// enabled by a given extension
func getExtensionTextureFormats(extension string) []GLenum {
	switch extension {
	case "GL_AMD_compressed_ATC_texture":
		return []GLenum{
			GLenum_GL_ATC_RGB_AMD,
			GLenum_GL_ATC_RGBA_EXPLICIT_ALPHA_AMD,
			GLenum_GL_ATC_RGBA_INTERPOLATED_ALPHA_AMD,
		}
	case "GL_OES_compressed_ETC1_RGB8_texture":
		return []GLenum{
			GLenum_GL_ETC1_RGB8_OES,
		}
	case "GL_EXT_texture_compression_dxt1":
		return []GLenum{
			GLenum_GL_COMPRESSED_RGB_S3TC_DXT1_EXT,
			GLenum_GL_COMPRESSED_RGBA_S3TC_DXT1_EXT,
		}
	case "GL_EXT_texture_compression_s3tc", "GL_NV_texture_compression_s3tc":
		return []GLenum{
			GLenum_GL_COMPRESSED_RGB_S3TC_DXT1_EXT,
			GLenum_GL_COMPRESSED_RGBA_S3TC_DXT1_EXT,
			GLenum_GL_COMPRESSED_RGBA_S3TC_DXT3_EXT,
			GLenum_GL_COMPRESSED_RGBA_S3TC_DXT5_EXT,
		}
	case "GL_KHR_texture_compression_astc_ldr":
		return []GLenum{
			GLenum_GL_COMPRESSED_RGBA_ASTC_4x4_KHR,
			GLenum_GL_COMPRESSED_RGBA_ASTC_5x4_KHR,
			GLenum_GL_COMPRESSED_RGBA_ASTC_5x5_KHR,
			GLenum_GL_COMPRESSED_RGBA_ASTC_6x5_KHR,
			GLenum_GL_COMPRESSED_RGBA_ASTC_6x6_KHR,
			GLenum_GL_COMPRESSED_RGBA_ASTC_8x5_KHR,
			GLenum_GL_COMPRESSED_RGBA_ASTC_8x6_KHR,
			GLenum_GL_COMPRESSED_RGBA_ASTC_8x8_KHR,
			GLenum_GL_COMPRESSED_RGBA_ASTC_10x5_KHR,
			GLenum_GL_COMPRESSED_RGBA_ASTC_10x6_KHR,
			GLenum_GL_COMPRESSED_RGBA_ASTC_10x8_KHR,
			GLenum_GL_COMPRESSED_RGBA_ASTC_10x10_KHR,
			GLenum_GL_COMPRESSED_RGBA_ASTC_12x10_KHR,
			GLenum_GL_COMPRESSED_RGBA_ASTC_12x12_KHR,
			GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_4x4_KHR,
			GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x4_KHR,
			GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x5_KHR,
			GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x5_KHR,
			GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x6_KHR,
			GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x5_KHR,
			GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x6_KHR,
			GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x8_KHR,
			GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x5_KHR,
			GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x6_KHR,
			GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x8_KHR,
			GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x10_KHR,
			GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x10_KHR,
			GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x12_KHR,
		}
	case "GL_EXT_texture_compression_latc", "GL_NV_texture_compression_latc":
		return []GLenum{
			GLenum_GL_COMPRESSED_LUMINANCE_LATC1_EXT,
			GLenum_GL_COMPRESSED_SIGNED_LUMINANCE_LATC1_EXT,
			GLenum_GL_COMPRESSED_LUMINANCE_ALPHA_LATC2_EXT,
			GLenum_GL_COMPRESSED_SIGNED_LUMINANCE_ALPHA_LATC2_EXT,
		}
	default:
		return []GLenum{}
	}
}
