package gles

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/gfxapi/gles/glsl"
	"android.googlesource.com/platform/tools/gpu/gfxapi/gles/glsl/ast"
)

// precisionStrip returns a transform that removes all precision specifiers from
// shader programs.
func precisionStrip() atom.Transform {
	return func(id atom.ID, a atom.Atom, out atom.Writer) {
		if cmd, ok := a.(*GlShaderSource); ok {
			var src string

			for _, s := range cmd.In.Source {
				src = src + s
			}

			tree, err := glsl.Parse(src, ast.LangVertexShader)
			if len(err) > 0 {
				panic(fmt.Errorf("Failed to parse shader '%s': %s", src, err[0]))
			}

			precisionStripVisit(tree)

			src = fmt.Sprint(glsl.Formatter(tree))

			out.Write(id, &GlShaderSource{
				In: GlShaderSource_In{
					Shader: cmd.In.Shader,
					Count:  1,
					Source: []string{src},
				},
			})
		} else {
			out.Write(id, a)
		}
	}
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
