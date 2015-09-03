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
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/service"
)

type features struct {
	uncompressedTextureFormats map[GLenum]struct{}
	compressedTextureFormats   map[GLenum]struct{}
}

func getFeatures(version, extensions string) (features, error) {
	v, err := ParseVersion(version)
	if err != nil {
		return features{}, err
	}

	f := features{
		uncompressedTextureFormats: getSupportedUncompressedTextureFormats(*v, extensions),
		compressedTextureFormats:   getSupportedCompressedTextureFormats(extensions),
	}

	return f, nil
}

func compat(device *service.Device, d database.Database, l log.Logger) (atom.Transformer, error) {
	l = log.Enter(l, "compat")

	target, err := getFeatures(device.Version, device.Extensions)
	if err != nil {
		return nil, fmt.Errorf(
			"Error '%v' when getting feature list for version: '%s', extensions: '%s'.",
			err, device.Version, device.Extensions)
	}

	s := gfxapi.NewState()
	return atom.Transform("compat", func(i atom.ID, a atom.Atom, out atom.Writer) {
		switch a := a.(type) {

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
