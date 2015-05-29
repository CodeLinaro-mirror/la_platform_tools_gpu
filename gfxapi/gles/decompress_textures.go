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
	"bytes"
	"fmt"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/cyclic"
	"android.googlesource.com/platform/tools/gpu/binary/vle"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/database/store"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/image"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/service"
)

// uncompressTextures returns an atom transform that replaces all
// GlCompressedTexImage2D atoms with glTexImage2D atoms for the compressed
// texture formats:
//   GL_ATC_RGB_AMD
//   GL_ATC_RGBA_EXPLICIT_ALPHA_AMD
//   GL_ETC1_RGB8_OES
func decompressTextures(capture service.CaptureId, d database.Database, l log.Logger) atom.Transformer {
	l = l.Enter("decompressTextures")
	s := gfxapi.NewState()
	return atom.Transform("DecompressTextures", func(id atom.ID, a atom.Atom, out atom.Writer) {
		if err := a.Mutate(s, d, l); err != nil {
			l.Errorf("%v", err)
		}

		switch a := a.(type) {
		case *GlCompressedTexImage2D:
			resourceID := calcTextureID(capture, id, a)
			var blob store.Blob
			if d.Load(resourceID, l, &blob) != nil {
				decompressed, err := decompress(d, l, a, s.Memory[memory.ApplicationPool])
				if err != nil {
					panic(err)
				}
				blob = store.Blob{Data: decompressed}
				decompressedID, err := d.Store(&blob, l)
				if err != nil {
					panic(err)
				}
				err = d.StoreLink(decompressedID, resourceID, l)
				if err != nil {
					panic(err)
				}
			}

			address := memory.Pointer(0xF000000000000000)

			out.Write(id, NewGlTexImage2D(
				a.Target,
				a.Level,
				TexelFormat_GL_RGBA,
				a.Width,
				a.Height,
				a.Border,
				TexelFormat_GL_RGBA,
				TexelType_GL_UNSIGNED_BYTE,
				address,
			).AddRead(address.Range(uint64(len(blob.Data))), resourceID))

		default:
			out.Write(id, a)
		}
	})
}

func calcTextureID(capture service.CaptureId, id atom.ID, a atom.Atom) binary.ID {
	buf := &bytes.Buffer{}
	e := cyclic.Encoder(vle.Writer(buf))
	e.Value(&capture)
	e.Uint64(uint64(id))
	e.Value(a)
	return binary.NewID(buf.Bytes())
}

func decompress(d database.Database, l log.Logger, a *GlCompressedTexImage2D, m *memory.Pool) ([]byte, error) {
	compressedSize := uint64(a.ImageSize)
	compressed, err := m.Slice(memory.Range{Base: a.Data.Address, Size: compressedSize}).Get(d, l)
	if err != nil {
		panic(err)
	}
	switch a.Format {
	case CompressedTexelFormat_GL_ATC_RGB_AMD:
		return image.Convert(compressed[:compressedSize], int(a.Width), int(a.Height), image.ATC_RGB_AMD(), image.RGBA())
	case CompressedTexelFormat_GL_ATC_RGBA_EXPLICIT_ALPHA_AMD:
		return image.Convert(compressed[:compressedSize], int(a.Width), int(a.Height), image.ATC_RGBA_EXPLICIT_ALPHA_AMD(), image.RGBA())
	case CompressedTexelFormat_GL_ETC1_RGB8_OES:
		return image.Convert(compressed[:compressedSize], int(a.Width), int(a.Height), image.ETC1_RGB8_OES(), image.RGBA())
	}
	return nil, fmt.Errorf("Unsupported input format: %s", a.Format.String())
}
