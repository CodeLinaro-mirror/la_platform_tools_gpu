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
	"android.googlesource.com/platform/tools/gpu/gfxapi/gles/glsl"
	"android.googlesource.com/platform/tools/gpu/gfxapi/gles/glsl/ast"
)

// precisionStrip returns a transform that removes all precision specifiers from
// shader programs.
func precisionStrip() atom.Transformer {
	return atom.Transform("PrecisionStrip", func(id atom.ID, a atom.Atom, out atom.Writer) {
		if cmd, ok := a.(*GlShaderSource); ok {
			var src string

			for _, s := range cmd.Source {
				src = src + s
			}

			tree, err := glsl.Parse(src, ast.LangVertexShader)
			if len(err) > 0 {
				panic(fmt.Errorf("Failed to parse shader '%s': %s", src, err[0]))
			}

			precisionStripVisit(tree)

			src = fmt.Sprint(glsl.Formatter(tree))

			out.Write(id, &GlShaderSource{
				Shader: cmd.Shader,
				Count:  1,
				Source: []string{src},
			})
		} else {
			out.Write(id, a)
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
