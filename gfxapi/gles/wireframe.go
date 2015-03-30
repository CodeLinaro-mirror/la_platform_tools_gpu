package gles

import (
	"bytes"
	eb "encoding/binary"
	"fmt"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary/endian"
	"android.googlesource.com/platform/tools/gpu/binary/flat"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/database/store"
	"android.googlesource.com/platform/tools/gpu/gfxapi/state"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
)

// wireframe returns an atom transform that replaces all draw calls of triangle
// primitives with draw calls of a wireframe equivalent.
func wireframe(db database.Database, logger log.Logger) atom.Transformer {
	logger = logger.Enter("Wireframe")

	s := state.New()
	return atom.Transform("Wireframe", func(id atom.ID, a atom.Atom, out atom.Writer) {
		if err := s.Mutate(a); err != nil {
			logger.Error("%v", err)
		}

		if a.Flags().IsDrawCall() {
			c := getState(a, s)
			cid := a.ContextID()
			indices, drawMode, err := getIndices(id, a, db, c, &s.Memory, logger)
			if err != nil {
				logger.Error(err.Error())
				return
			}
			indices, drawMode, err = makeWireframe(indices, drawMode)
			if err != nil {
				logger.Error(err.Error())
				return
			}

			// Store the wire-frame at virtual address
			address := memory.Pointer(0x5746000000000000)
			wireframeData, wireframeDataType := encodeIndices(indices)
			res := store.Blob{Data: wireframeData}
			resID, err := db.Store(&res, logger)
			if err != nil {
				panic(err)
			}
			out.Write(id, &atom.Observation{
				Range:      memory.Range{Base: address, Size: uint64(len(wireframeData))},
				ResourceID: resID,
				Context:    a.ContextID(),
			})

			// Unbind the index buffer
			oldIndexBufferID := c.BoundBuffers[BufferTarget_GL_ELEMENT_ARRAY_BUFFER]
			out.Write(id, NewGlBindBuffer(
				cid, BufferTarget(BufferTarget_GL_ELEMENT_ARRAY_BUFFER), 0))

			// Draw the wire-frame
			out.Write(id, NewGlDrawElements(
				cid, drawMode, int32(len(indices)), wireframeDataType, IndicesPointer(address)))

			// Rebind the old index buffer
			out.Write(id, NewGlBindBuffer(
				cid, BufferTarget(BufferTarget_GL_ELEMENT_ARRAY_BUFFER), oldIndexBufferID))
		} else {
			out.Write(id, a)
		}
	})
}

type index uint32

// TODO: The decode/encode methods below assume little endian

func decodeIndices(data []byte, indicesType IndicesType) ([]index, error) {
	dec := flat.Decoder(endian.Reader(bytes.NewBuffer(data), eb.LittleEndian))
	indices := make([]index, 0)
	switch indicesType {
	case IndicesType_GL_UNSIGNED_BYTE:
		for {
			if val, err := dec.Uint8(); err == nil {
				indices = append(indices, index(val))
			} else {
				return indices, nil
			}
		}

	case IndicesType_GL_UNSIGNED_SHORT:
		for {
			if val, err := dec.Uint16(); err == nil {
				indices = append(indices, index(val))
			} else {
				return indices, nil
			}
		}

	case IndicesType_GL_UNSIGNED_INT:
		for {
			if val, err := dec.Uint32(); err == nil {
				indices = append(indices, index(val))
			} else {
				return indices, nil
			}
		}

	default:
		return nil, fmt.Errorf("Invalid index type: %v", indicesType)
	}
}

func encodeIndices(indices []index) ([]byte, IndicesType) {
	maxIndex := index(0)
	for _, v := range indices {
		if v > maxIndex {
			maxIndex = v
		}
	}
	buf := &bytes.Buffer{}
	enc := flat.Encoder(endian.Writer(buf, eb.LittleEndian))
	switch {
	case maxIndex > 0xFFFF:
		// TODO: GL_UNSIGNED_INT in glDrawElements is supported only since GLES 3.0
		for _, v := range indices {
			enc.Uint32(uint32(v))
		}
		return buf.Bytes(), IndicesType_GL_UNSIGNED_INT

	case maxIndex > 0xFF:
		for _, v := range indices {
			enc.Uint16(uint16(v))
		}
		return buf.Bytes(), IndicesType_GL_UNSIGNED_SHORT

	default:
		for _, v := range indices {
			enc.Uint8(uint8(v))
		}
		return buf.Bytes(), IndicesType_GL_UNSIGNED_BYTE
	}
}

// Get the effective index buffer and primitive type for draw call
func getIndices(id atom.ID, a atom.Atom, db database.Database, s *State, m *memory.Memory, logger log.Logger) ([]index, DrawMode, error) {
	switch a := a.(type) {
	case *GlDrawArrays:
		indices := make([]index, a.In.IndexCount)
		for i := range indices {
			indices[i] = index(a.In.FirstIndex) + index(i)
		}
		return indices, a.In.DrawMode, nil

	case *GlDrawElements:
		indexSize := map[IndicesType]uint64{
			IndicesType_GL_UNSIGNED_BYTE:  1,
			IndicesType_GL_UNSIGNED_SHORT: 2,
			IndicesType_GL_UNSIGNED_INT:   4,
		}[a.In.IndicesType]
		indexBufferID := s.BoundBuffers[BufferTarget_GL_ELEMENT_ARRAY_BUFFER]
		if indexBufferID == 0 {
			// Get the index buffer data from pointer
			size := uint64(a.In.ElementCount) * indexSize
			mem := m.Slice(memory.Range{Base: memory.Pointer(a.In.Indices), Size: size})
			data, err := mem.Get(db, logger)
			if err != nil {
				return nil, a.In.DrawMode, err
			}
			indices, err := decodeIndices(data, a.In.IndicesType)
			return indices, a.In.DrawMode, err
		} else {
			// Get the index buffer data from buffer
			indexBuffer := s.Instances.Buffers[indexBufferID]
			if indexBuffer == nil {
				return nil, 0, fmt.Errorf("Can not find buffer %v", indexBufferID)
			}
			offset := memory.Pointer(a.In.Indices)
			size := uint64(a.In.ElementCount) * indexSize
			mem := indexBuffer.Data.Slice(memory.Range{Base: offset, Size: size})
			data, err := mem.Get(db, logger)
			if err != nil {
				return nil, a.In.DrawMode, err
			}
			indices, err := decodeIndices(data, a.In.IndicesType)
			return indices, a.In.DrawMode, err
		}

	default:
		return nil, 0, fmt.Errorf("Unknown draw command %v", a)
	}
}

func appendWireframeOfTriangle(lines []index, v0, v1, v2 index) []index {
	if v0 == v1 || v1 == v2 || v2 == v0 {
		return lines // Ignore degenerate triangle
	} else {
		return append(lines, v0, v1, v1, v2, v2, v0)
	}
}

func makeWireframe(indices []index, drawMode DrawMode) ([]index, DrawMode, error) {
	switch drawMode {
	case DrawMode_GL_POINTS, DrawMode_GL_LINES, DrawMode_GL_LINE_STRIP, DrawMode_GL_LINE_LOOP:
		return indices, drawMode, nil

	case DrawMode_GL_TRIANGLES:
		numTriangles := len(indices) / 3
		lines := make([]index, 0, numTriangles*6)
		for i := 0; i < numTriangles; i++ {
			lines = appendWireframeOfTriangle(lines, indices[i*3], indices[i*3+1], indices[i*3+2])
		}
		return lines, DrawMode_GL_LINES, nil

	case DrawMode_GL_TRIANGLE_STRIP:
		numTriangles := len(indices) - 2
		if numTriangles > 0 {
			lines := make([]index, 0, numTriangles*6)
			for i := 0; i < numTriangles; i++ {
				lines = appendWireframeOfTriangle(lines, indices[i], indices[i+1], indices[i+2])
			}
			return lines, DrawMode_GL_LINES, nil
		}
		return []index{}, DrawMode_GL_LINES, nil

	case DrawMode_GL_TRIANGLE_FAN:
		numTriangles := len(indices) - 2
		if numTriangles > 0 {
			lines := make([]index, 0, numTriangles*6)
			for i := 0; i < numTriangles; i++ {
				lines = appendWireframeOfTriangle(lines, indices[0], indices[i+1], indices[i+2])
			}
			return lines, DrawMode_GL_LINES, nil
		}
		return []index{}, DrawMode_GL_LINES, nil

	default:
		return nil, 0, fmt.Errorf("Unknown mode: %v", drawMode)
	}
}
