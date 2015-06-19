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
	"android.googlesource.com/platform/tools/gpu/builder"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/image"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/service"
)

// decompressTextures returns an atom transform that replaces all
// GlCompressedTexImage2D atoms with glTexImage2D atoms for the compressed
// texture formats:
//   GL_ATC_RGB_AMD
//   GL_ATC_RGBA_EXPLICIT_ALPHA_AMD
//   GL_ETC1_RGB8_OES
func decompressTextures(capture service.CaptureId, d database.Database, l log.Logger) atom.Transformer {
	l = l.Enter("decompressTextures")
	s := gfxapi.NewState()
	return atom.Transform("DecompressTextures", func(i atom.ID, a atom.Atom, out atom.Writer) {
		if err := a.Mutate(s, d, l); err != nil {
			l.Errorf("%v", err)
		}

		switch a := a.(type) {
		case *GlCompressedTexImage2D:
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

			address := memory.Tmp.Base
			size := a.Width * a.Height * 4
			out.Write(i, NewGlTexImage2D(
				a.Target,
				a.Level,
				TexelFormat_GL_RGBA,
				a.Width,
				a.Height,
				a.Border,
				TexelFormat_GL_RGBA,
				TexelType_GL_UNSIGNED_BYTE,
				address,
			).AddRead(address.Range(uint64(size)), id))

		default:
			out.Write(i, a)
		}
	})
}

func getImageFormat(f CompressedTexelFormat) image.Format {
	switch f {
	case CompressedTexelFormat_GL_ATC_RGB_AMD:
		return image.ATC_RGB_AMD()
	case CompressedTexelFormat_GL_ATC_RGBA_EXPLICIT_ALPHA_AMD:
		return image.ATC_RGBA_EXPLICIT_ALPHA_AMD()
	case CompressedTexelFormat_GL_ETC1_RGB8_OES:
		return image.ETC1_RGB8_OES()
	}
	panic(fmt.Errorf("Unsupported input format: %s", f.String()))
}
