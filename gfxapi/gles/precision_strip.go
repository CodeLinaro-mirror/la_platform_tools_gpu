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
	"android.googlesource.com/platform/tools/gpu/gfxapi/gles/glsl"
	"android.googlesource.com/platform/tools/gpu/gfxapi/gles/glsl/ast"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/service"
)

// precisionStrip returns a transform that removes all precision specifiers from
// shader programs if the device target doesn't support them.
func precisionStrip(device *service.Device, d database.Database, l log.Logger) atom.Transformer {
	s := gfxapi.NewState()
	if v, err := ParseVersion(device.Version); err == nil {
		if v.IsES {
			return nil
		}
	}
	return atom.Transform("PrecisionStrip", func(i atom.ID, a atom.Atom, out atom.Writer) {
		a.Mutate(s, d, l)
		if cmd, ok := a.(*GlShaderSource); ok {
			shader := getContext(s).Instances.Shaders.Get(cmd.Shader)

			tree, err := glsl.Parse(shader.Source, ast.LangVertexShader)
			if len(err) > 0 {
				panic(fmt.Errorf("Failed to parse shader source at atom %d '%s': %s", i, shader.Source, err[0]))
			}

			precisionStripVisit(tree)

			src := fmt.Sprint(glsl.Formatter(tree))

			out.Write(i,
				NewGlShaderSource(cmd.Shader, 1, memory.Tmp.Base, 0).
					AddRead(atom.Data(s.Architecture, d, l, memory.Tmp.Base+8, src)).
					AddRead(atom.Data(s.Architecture, d, l, memory.Tmp.Base+0, memory.Tmp.Base+8)))
		} else {
			out.Write(i, a)
		}
	})
}

func isPrecision(n interface{}) bool {
	switch n := n.(type) {
	case *ast.DeclarationStmt:
		return isPrecision(n.Decl)
	case *ast.PrecisionDecl:
		return true
	}
	return false
}

func removePrecisions(arr []interface{}) (ret []interface{}) {
	for _, n := range arr {
		if !isPrecision(n) {
			ret = append(ret, n)
		}
	}
	return
}

func clearIfPrecision(d *interface{}) {
	if isPrecision(*d) {
		*d = &ast.EmptyStmt{}
	}
}

func precisionStripVisit(n interface{}) {
	switch n := n.(type) {
	case *ast.Ast:
		n.Decls = removePrecisions(n.Decls)
	case *ast.IfStmt:
		clearIfPrecision(&n.ThenStmt)
		clearIfPrecision(&n.ElseStmt)
	case *ast.CompoundStmt:
		n.Stmts = removePrecisions(n.Stmts)
	case *ast.WhileStmt:
		clearIfPrecision(&n.Stmt)
	case *ast.DoStmt:
		clearIfPrecision(&n.Stmt)
	case *ast.ForStmt:
		clearIfPrecision(&n.Init)
		clearIfPrecision(&n.Body)
	case *ast.BuiltinType:
		n.Precision = ast.NoneP
	}
	ast.VisitChildren(n, precisionStripVisit)
}
