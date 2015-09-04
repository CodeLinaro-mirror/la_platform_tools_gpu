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
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/gfxapi/gles/glsl/ast"
	"android.googlesource.com/platform/tools/gpu/interval"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/replay/builder"
	"android.googlesource.com/platform/tools/gpu/service"
)

type support int

const (
	unsupported support = iota
	supported
	required
)

func (s support) String() string {
	switch s {
	case unsupported:
		return "unsupported"
	case supported:
		return "supported"
	case required:
		return "required"
	default:
		return fmt.Sprintf("support<%d>", s)
	}
}

type features struct {
	vertexArrayObjects         support // support for VBOs
	uncompressedTextureFormats map[GLenum]struct{}
	compressedTextureFormats   map[GLenum]struct{}
}

func getFeatures(version, extensions string, l log.Logger) (features, error) {
	v, err := ParseVersion(version)
	if err != nil {
		return features{}, err
	}

	utfs, err := getSupportedUncompressedTextureFormats(*v, extensions)
	if err != nil {
		log.W(l, "getSupportedUncompressedTextureFormats returned error: %v", err)
	}

	f := features{
		uncompressedTextureFormats: utfs,
		compressedTextureFormats:   getSupportedCompressedTextureFormats(extensions),
	}

	// TODO: Properly check the specifications for these flags.
	switch {
	case v.IsES && v.Major >= 3:
		f.vertexArrayObjects = supported
	case !v.IsES && v.Major >= 3:
		f.vertexArrayObjects = required
	}

	return f, nil
}

func compat(device *service.Device, d database.Database, l log.Logger) (atom.Transformer, error) {
	l = log.Enter(l, "compat")

	target, err := getFeatures(device.Version, device.Extensions, l)
	if err != nil {
		return nil, fmt.Errorf(
			"Error '%v' when getting feature list for version: '%s', extensions: '%s'.",
			err, device.Version, device.Extensions)
	}

	contexts := map[*Context]features{}

	s := gfxapi.NewState()
	return atom.Transform("compat", func(i atom.ID, a atom.Atom, out atom.Writer) {
		switch a := a.(type) {
		case *ContextInfo:
			ctx := getContext(s)
			if _, found := contexts[ctx]; found {
				break
			}

			source, err := getFeatures(a.Version, a.Extensions, l)
			if err != nil {
				log.E(l, "Error '%v' when getting feature list for version: '%s', extensions: '%s'.",
					err, a.Version, a.Extensions)
				break
			}

			contexts[ctx] = source

			if target.vertexArrayObjects == required &&
				source.vertexArrayObjects != required {
				// Capture does not have to support VAO, but replay device requires it.
				// Satisfy the target by creating and binding a single VAO.
				a.Mutate(s, d, l)
				out.Write(i, a)
				out.Write(atom.NoID, NewGlGenVertexArrays(1, memory.Tmp).
					AddRead(atom.Data(s.Architecture, d, l, memory.Tmp, VertexArrayId(1))))
				out.Write(atom.NoID, NewGlBindVertexArray(1))
				return
			}

		case *GlShaderSource:
			// Apply the state mutation of the unmodified glShaderSource atom.
			// This is so we can grab the source string from the Shader object.
			a.Mutate(s, d, l)
			shader := getContext(s).Instances.Shaders.Get(a.Shader)

			lang := ast.LangVertexShader
			switch shader.Type {
			case GLenum_GL_VERTEX_SHADER:
			case GLenum_GL_FRAGMENT_SHADER:
				lang = ast.LangFragmentShader
			default:
				log.W(l, "Unknown shader type %v", shader.Type)
			}

			src, err := glslCompat(shader.Source, lang, device)
			if err != nil {
				log.E(l, "Failed to reformat GLSL source for atom %d: %v", i, err)
			}

			a = NewGlShaderSource(a.Shader, 1, memory.Tmp, memory.Nullptr).
				AddRead(atom.Data(s.Architecture, d, l, memory.Tmp.Offset(8), src)).
				AddRead(atom.Data(s.Architecture, d, l, memory.Tmp, memory.Tmp.Offset(8)))
			a.Mutate(s, d, l)
			out.Write(i, a)
			return

		case *GlVertexAttribPointer:
			if target.vertexArrayObjects == required &&
				getContext(s).BoundBuffers[GLenum_GL_ARRAY_BUFFER] == 0 {
				// Client-pointers are not supported, we need to copy this data to a buffer.
				// However, we can't do this now as the observation only happens at the draw call.
				// Apply the state changes, but don't write the emit the atom - we need to defer
				// the trickery to the draw call.
				a.Mutate(s, d, l)
				return
			}

		case *GlDrawArrays:
			if target.vertexArrayObjects == required {
				if c := getContext(s); clientVAsBound(c) {
					first := int(a.FirstIndex)
					last := first + int(a.IndexCount) - 1
					defer moveClientVBsToVAs(first, last, i, a, s, c, d, l, out)()
				}
			}

		case *GlDrawElements:
			if target.vertexArrayObjects == required {
				c := getContext(s)
				e := externs{a: a, s: s, d: d, l: l}

				ib := c.BoundBuffers[GLenum_GL_ELEMENT_ARRAY_BUFFER]
				clientIB := ib == 0
				clientVB := clientVAsBound(c)
				if clientIB {
					// The indices for the glDrawElements call is in client memory.
					// We need to move this into a temporary buffer.

					// Generate a new element array buffer and bind it.
					id := BufferId(newUnusedID(func(x uint32) bool { _, ok := c.Instances.Buffers[BufferId(x)]; return ok }))
					c.Instances.Buffers[id] = &Buffer{} // Not used aside from reserving the ID.
					out.Write(atom.NoID, NewGlGenBuffers(1, memory.Tmp).
						AddRead(atom.Data(s.Architecture, d, l, memory.Tmp, id)))
					out.Write(atom.NoID, NewGlBindBuffer(GLenum_GL_ELEMENT_ARRAY_BUFFER, id))

					// By moving the draw call's observations earlier, populate the element array buffer.
					size, base := DataTypeSize(a.IndicesType)*int(a.ElementCount), a.Indices.Pointer
					glBufferData := NewGlBufferData(GLenum_GL_ELEMENT_ARRAY_BUFFER, GLsizeiptr(size), base, GLenum_GL_STATIC_DRAW)
					glBufferData.observations = a.observations
					out.Write(atom.NoID, glBufferData)

					// Clean-up
					defer func() {
						out.Write(atom.NoID, NewGlBindBuffer(GLenum_GL_ELEMENT_ARRAY_BUFFER, ib))
						delete(c.Instances.Buffers, id)
					}()

					if clientVB {
						// Some of the vertex arrays for the glDrawElements call is in
						// client memory and we need to move this into temporary buffer(s).
						// The indices are also in client memory, so we need to apply the
						// atom's reads now so that the indices can be read from the
						// application pool.
						a.Observations().ApplyReads(s.Memory[memory.ApplicationPool])
						limits := e.calcIndexLimits(U8ᵖ(a.Indices), a.IndicesType, 0, uint32(a.ElementCount))
						defer moveClientVBsToVAs(int(limits.Min), int(limits.Max), i, a, s, c, d, l, out)()
					}

					glDrawElements := *a
					glDrawElements.Indices.Pointer.Address = 0
					glDrawElements.Mutate(s, d, l)
					out.Write(i, &glDrawElements)
					return

				} else if clientVB { // GL_ELEMENT_ARRAY_BUFFER is bound
					// Some of the vertex arrays for the glDrawElements call is in
					// client memory and we need to move this into temporary buffer(s).
					// The indices are server-side, so can just be read from the internal
					// pooled buffer.
					data := c.Instances.Buffers[ib].Data.Index(0, s)
					base := uint32(a.Indices.Pointer.Address)
					limits := e.calcIndexLimits(data, a.IndicesType, base, uint32(a.ElementCount))
					defer moveClientVBsToVAs(int(limits.Min), int(limits.Max), i, a, s, c, d, l, out)()
				}
			}

		case *GlTexImage2D:
			if _, supported := target.uncompressedTextureFormats[a.Format]; !supported {
				if err := convertTexImage2D(i, a, s, d, l, out); err == nil {
					return
				} else {
					log.E(l, "Failed to convert texture: %v", err)
				}
			}

		case *GlTexSubImage2D:
			if _, supported := target.uncompressedTextureFormats[a.Format]; !supported {
				if err := convertTexSubImage2D(i, a, s, d, l, out); err == nil {
					return
				} else {
					log.E(l, "Failed to convert texture: %v", err)
				}
			}

		case *GlCompressedTexImage2D:
			if _, supported := target.compressedTextureFormats[a.Format]; !supported {
				if err := decompressTexImage2D(i, a, s, d, l, out); err == nil {
					return
				} else {
					log.E(l, "Failed to decompress texture: %v", err)
				}
			}
		}

		a.Mutate(s, d, l)
		out.Write(i, a)

	}), nil
}

// clientVAsBound returns true if there are any vertex attribute arrays enabled
// with pointers to client-side memory.
func clientVAsBound(c *Context) bool {
	for _, arr := range c.VertexAttributeArrays {
		if arr.Enabled && arr.Buffer == 0 {
			return true
		}
	}
	return false
}

// moveClientVBsToVAs is a compatability helper for transforming client-side
// vertex array data (which is not supported by glVertexAttribPointer in later
// versions of GL), into array-buffers.
func moveClientVBsToVAs(
	first, last int, // vertex indices
	i atom.ID,
	a atom.Atom,
	s *gfxapi.State,
	c *Context,
	d database.Database,
	l log.Logger,
	out atom.Writer) (revert func()) {

	rngs := interval.U64RangeList{}

	// Gather together all the client-buffers in use by the vertex-attribs.
	// Merge together all the memory intervals that these use.
	for _, arr := range c.VertexAttributeArrays {
		if arr.Enabled && arr.Buffer == 0 {
			interval.Merge(&rngs, arr.MemoryRange(first, last).Span(), true)
		}
	}

	if len(rngs) == 0 {
		// Draw call does not use client-side buffers. Just draw.
		out.Write(i, a)
		a.Mutate(s, d, l)
		return
	}

	// Create an array-buffer for each chunk of overlapping client-side buffers in
	// use. These are populated with data below.
	ids := make([]BufferId, len(rngs))
	for i := range rngs {
		id := BufferId(newUnusedID(func(x uint32) bool { _, ok := c.Instances.Buffers[BufferId(x)]; return ok }))
		c.Instances.Buffers[id] = &Buffer{} // Not used aside from reserving the ID.
		ids[i] = id
	}
	out.Write(atom.NoID, NewGlGenBuffers(GLsizei(len(ids)), memory.Tmp).
		AddRead(atom.Data(s.Architecture, d, l, memory.Tmp, ids)))

	// Apply the memory observations that were made by the draw call now.
	// We need to do this as the glBufferData calls below will require the data.
	out.Write(atom.NoID, replay.Custom(func(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error {
		a.Observations().ApplyReads(s.Memory[memory.ApplicationPool])
		return nil
	}))

	// Note: be careful of overwriting the observations made above, before the
	// calls to glBufferData below.

	// Fill the array-buffers with the observed memory data.
	for i, rng := range rngs {
		base := memory.Pointer{Address: rng.First, Pool: memory.ApplicationPool}
		size := GLsizeiptr(rng.Count)
		out.Write(atom.NoID, NewGlBindBuffer(GLenum_GL_ARRAY_BUFFER, ids[i]))
		out.Write(atom.NoID, NewGlBufferData(GLenum_GL_ARRAY_BUFFER, size, base, GLenum_GL_STATIC_DRAW))
	}

	// Redirect all the vertex attrib arrays to point to the array-buffer data.
	for l, arr := range c.VertexAttributeArrays {
		if arr.Enabled && arr.Buffer == 0 {
			i := interval.IndexOf(&rngs, arr.Pointer.Address)
			offset := arr.Pointer.Address - rngs[i].First
			out.Write(atom.NoID, NewGlBindBuffer(GLenum_GL_ARRAY_BUFFER, ids[i]))
			out.Write(atom.NoID, &GlVertexAttribPointer{
				Location:   l,
				Size:       GLint(arr.Size),
				Type:       arr.Type,
				Normalized: arr.Normalized,
				Stride:     arr.Stride,
				Data:       NewVertexPointer(offset),
			})
		}
	}

	// Restore original state.
	return func() {
		out.Write(atom.NoID, NewGlBindBuffer(GLenum_GL_ARRAY_BUFFER, c.BoundBuffers[GLenum_GL_ARRAY_BUFFER]))
		for _, id := range ids {
			delete(c.Instances.Buffers, id)
			out.Write(atom.NoID, NewGlDeleteBuffers(1, memory.Tmp).
				AddRead(atom.Data(s.Architecture, d, l, memory.Tmp, id)))
		}
	}
}
