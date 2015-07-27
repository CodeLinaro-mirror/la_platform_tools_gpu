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
	"android.googlesource.com/platform/tools/gpu/builder"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/image"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/service"
	"android.googlesource.com/platform/tools/gpu/service/path"
)

// decompressTextures returns an atom transform that replaces GlCompressedTexImage2D atoms with
// GlTexImage2D atoms if unsupported by the target device, for the following compressed texture formats:
//   GL_ATC_RGB_AMD
//   GL_ATC_RGBA_EXPLICIT_ALPHA_AMD
//   GL_ETC1_RGB8_OES
func decompressTextures(device *service.Device, capture *path.Capture, d database.Database, l log.Logger) atom.Transformer {
	l = log.Enter(l, "decompressTextures")
	s := gfxapi.NewState()
	supportedFormats := getCompressedFormats(device)
	return atom.Transform("DecompressTextures", func(i atom.ID, a atom.Atom, out atom.Writer) {
		if err := a.Mutate(s, d, l); err != nil {
			log.Errorf(l, "%v", err)
		}

		switch a := a.(type) {
		case *GlCompressedTexImage2D:
			// No-op if the compressed texture format is supported by the target device
			if _, supported := supportedFormats[a.Format]; supported {
				out.Write(i, a)
				return
			}

			id, err := database.Store(&builder.ConvertImage{
				Data:       a.Data.Slice(0, uint64(a.ImageSize), s).ResourceID(s, d, l),
				Width:      int(a.Width),
				Height:     int(a.Height),
				FormatFrom: getImageFormat(a.Format),
				FormatTo:   image.RGBA(),
			}, d, l)
			if err != nil {
				panic(err)
			}

			size := a.Width * a.Height * 4
			out.Write(i, NewGlTexImage2D(
				a.Target,
				a.Level,
				GLenum_GL_RGBA,
				a.Width,
				a.Height,
				a.Border,
				GLenum_GL_RGBA,
				GLenum_GL_UNSIGNED_BYTE,
				memory.Tmp,
			).AddRead(memory.Tmp.Range(uint64(size)), id))

		default:
			out.Write(i, a)
		}
	})
}

func getImageFormat(f GLenum) image.Format {
	switch f {
	case GLenum_GL_ATC_RGB_AMD:
		return image.ATC_RGB_AMD()
	case GLenum_GL_ATC_RGBA_EXPLICIT_ALPHA_AMD:
		return image.ATC_RGBA_EXPLICIT_ALPHA_AMD()
	case GLenum_GL_ETC1_RGB8_OES:
		return image.ETC1_RGB8_OES()
	}
	panic(fmt.Errorf("Unsupported input format: %s", f.String()))
}

// getCompressedFormats returns the set of supported compressed texture formats for a given device
func getCompressedFormats(device *service.Device) map[GLenum]struct{} {
	ret := map[GLenum]struct{}{}
	for _, extension := range strings.Split(device.Extensions, " ") {
		for _, format := range getExtensionFormats(extension) {
			ret[format] = struct{}{}
		}
	}
	return ret
}

// getExtensionFormats returns the list of compressed texture formats enabled by a given extension
func getExtensionFormats(extension string) []GLenum {
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
			GLenum_GL_COMPRESSED_LUMINANCE_LATC1_NV,
			GLenum_GL_COMPRESSED_SIGNED_LUMINANCE_LATC1_NV,
			GLenum_GL_COMPRESSED_LUMINANCE_ALPHA_LATC2_NV,
			GLenum_GL_COMPRESSED_SIGNED_LUMINANCE_ALPHA_LATC2_NV,
		}
	default:
		return []GLenum{}
	}
}
