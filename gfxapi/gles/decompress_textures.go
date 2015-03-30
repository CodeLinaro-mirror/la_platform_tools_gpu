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
	"android.googlesource.com/platform/tools/gpu/gfxapi/state"
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
func decompressTextures(capture service.CaptureId, db database.Database, logger log.Logger) atom.Transformer {
	logger = logger.Enter("decompressTextures")
	s := state.New()
	return atom.Transform("DecompressTextures", func(id atom.ID, a atom.Atom, out atom.Writer) {
		if err := s.Mutate(a); err != nil {
			logger.Error("%v", err)
		}

		switch a := a.(type) {
		case *GlCompressedTexImage2D:
			resourceID := calcTextureID(capture, id, a)
			var blob store.Blob
			if db.Load(resourceID, logger, &blob) != nil {
				decompressed, err := decompress(db, logger, a, &s.Memory)
				if err != nil {
					panic(err)
				}
				data := store.Blob{Data: decompressed}
				decompressedID, err := db.Store(&data, logger)
				if err != nil {
					panic(err)
				}
				err = db.StoreLink(decompressedID, resourceID, logger)
				if err != nil {
					panic(err)
				}
			}

			address := memory.Pointer(0xF000000000000000)
			out.Write(id, &atom.Observation{
				Range:      memory.Range{Base: address, Size: uint64(len(blob.Data))},
				ResourceID: resourceID,
				Context:    a.ContextID(),
			})

			out.Write(id, NewGlTexImage2D(
				a.ContextID(),
				a.In.Target,
				a.In.Level,
				TexelFormat_GL_RGBA,
				a.In.Width,
				a.In.Height,
				a.In.Border,
				TexelFormat_GL_RGBA,
				TexelType_GL_UNSIGNED_BYTE,
				TexturePointer(address),
			))

		default:
			out.Write(id, a)
		}
	})
}

func calcTextureID(capture service.CaptureId, id atom.ID, a atom.Atom) binary.ID {
	buf := &bytes.Buffer{}
	e := cyclic.Encoder(vle.Writer(buf))
	capture.Encode(e)
	e.Uint64(uint64(id))
	a.Encode(e)
	return binary.NewID(buf.Bytes())
}

func decompress(db database.Database, logger log.Logger, a *GlCompressedTexImage2D, m *memory.Memory) ([]byte, error) {
	pointer := memory.Pointer(a.In.Data)
	compressedSize := uint64(a.In.ImageSize)
	compressed, err := m.Slice(memory.Range{Base: pointer, Size: compressedSize}).Get(db, logger)
	if err != nil {
		panic(err)
	}
	switch a.In.Format {
	case CompressedTexelFormat_GL_ATC_RGB_AMD:
		return image.Convert(compressed[:compressedSize], int(a.In.Width), int(a.In.Height), image.ATC_RGB_AMD(), image.RGBA())
	case CompressedTexelFormat_GL_ATC_RGBA_EXPLICIT_ALPHA_AMD:
		return image.Convert(compressed[:compressedSize], int(a.In.Width), int(a.In.Height), image.ATC_RGBA_EXPLICIT_ALPHA_AMD(), image.RGBA())
	case CompressedTexelFormat_GL_ETC1_RGB8_OES:
		return image.Convert(compressed[:compressedSize], int(a.In.Width), int(a.In.Height), image.ETC1_RGB8_OES(), image.RGBA())
	}
	return nil, fmt.Errorf("Unsupported input format: %s", a.In.Format.String())
}
