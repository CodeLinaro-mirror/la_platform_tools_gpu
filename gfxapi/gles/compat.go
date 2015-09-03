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
	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/gfxapi/gles/glsl/ast"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/service"
)

func compat(device *service.Device, d database.Database, l log.Logger) (atom.Transformer, error) {
	l = log.Enter(l, "compat")

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
		}

		a.Mutate(s, d, l)
		out.Write(i, a)

	}), nil
}
