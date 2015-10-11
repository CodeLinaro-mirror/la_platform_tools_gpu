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
	"strings"

	"android.googlesource.com/platform/tools/gpu/gfxapi/gles/glsl"
	"android.googlesource.com/platform/tools/gpu/gfxapi/gles/glsl/ast"
	"android.googlesource.com/platform/tools/gpu/gfxapi/gles/glsl/parser"
	"android.googlesource.com/platform/tools/gpu/service"
)

type glslTransform int

const (
	glslStripPrecision   glslTransform = 0x01
	glslAttributeToIn    glslTransform = 0x02
	glslVaryingToIn      glslTransform = 0x04
	glslVaryingToOut     glslTransform = 0x08
	glslDepTextureFuncs  glslTransform = 0x10
	glslDeclareFragColor glslTransform = 0x20
	glslDeclareFragData  glslTransform = 0x40

	compatFragColor = "FragColor"
	compatFragData  = "FragData"
)

func (t glslTransform) stripPrecision() bool   { return t&glslStripPrecision != 0 }
func (t glslTransform) attributeToIn() bool    { return t&glslAttributeToIn != 0 }
func (t glslTransform) varyingToIn() bool      { return t&glslVaryingToIn != 0 }
func (t glslTransform) varyingToOut() bool     { return t&glslVaryingToOut != 0 }
func (t glslTransform) depTextureFuncs() bool  { return t&glslDepTextureFuncs != 0 }
func (t glslTransform) declareFragColor() bool { return t&glslDeclareFragColor != 0 }
func (t glslTransform) declareFragData() bool  { return t&glslDeclareFragData != 0 }

type glslTransformState struct {
	glslTransform
	FragColor *ast.VariableSym
	FragData  map[int]*ast.VariableSym
}

func (t glslTransform) apply(n interface{}) {
	s := glslTransformState{
		glslTransform: t,
		FragData:      map[int]*ast.VariableSym{},
	}
	s.apply(n, nil)
}

var vec4 = &ast.BuiltinType{Type: ast.TVec4}

func (t *glslTransformState) apply(child, parent interface{}) interface{} {
	ast.TransformChildren(child, t.apply)

	addOut := func(n *ast.Ast, v *ast.VariableSym) *ast.MultiVarDecl {
		decl := &ast.MultiVarDecl{
			Quals: &ast.TypeQualifiers{Storage: ast.StorOut},
			Type:  v.Type(),
			Vars:  []*ast.VariableSym{v},
		}
		decls := make([]interface{}, len(n.Decls)+1)
		copy(decls[1:], n.Decls)
		n.Decls = decls
		n.Decls[0] = decl
		return decl
	}
	fixReservedNames := func(name *string) {
		if t.depTextureFuncs() {
			switch *name {
			case "texture", "textureProj", "textureLod", "textureProjLod",
				"shadow", "shadowProj", "shadowLod", "shadowProjLod":
				// These weren't keywords before the texture-lookup renames, but are now.
				// Rename these usages.
				*name += "__"
			}
		}
	}

	switch n := child.(type) {
	case *ast.Ast:
		if t.stripPrecision() {
			n.Decls = removePrecisions(n.Decls)
		}
		if t.declareFragColor() && t.FragColor != nil {
			addOut(n, t.FragColor) // out vec4 FragColor;
		}
		if t.declareFragData() {
			for idx, fragdata := range t.FragData {
				decl := addOut(n, fragdata) // layout(location = N) out vec4 FragDataN;
				decl.Quals.Layout = &ast.LayoutQualifier{
					Ids: []ast.LayoutQualifierID{
						ast.LayoutQualifierID{Name: "location", Value: ast.IntValue(idx)},
					},
				}
			}
		}
	case *ast.IfStmt:
		if t.stripPrecision() {
			clearIfPrecision(&n.ThenStmt)
			clearIfPrecision(&n.ElseStmt)
		}
	case *ast.CompoundStmt:
		if t.stripPrecision() {
			n.Stmts = removePrecisions(n.Stmts)
		}
	case *ast.WhileStmt:
		if t.stripPrecision() {
			clearIfPrecision(&n.Stmt)
		}
	case *ast.DoStmt:
		if t.stripPrecision() {
			clearIfPrecision(&n.Stmt)
		}
	case *ast.ForStmt:
		if t.stripPrecision() {
			clearIfPrecision(&n.Init)
			clearIfPrecision(&n.Body)
		}
	case *ast.BuiltinType:
		if t.stripPrecision() {
			n.Precision = ast.NoneP
		}
	case *ast.IndexExpr:
		if t.declareFragData() {
			if vr, ok := n.Base.(*ast.VarRefExpr); ok {
				if v, ok := vr.Sym.(*ast.VariableSym); ok {
					if constExpr, ok := n.Index.(*ast.ConstantExpr); ok {
						if index, ok := constExpr.Value.(ast.IntValue); ok {
							idx := int(index)
							if v.Name() == "gl_FragData" {
								name := fmt.Sprintf("%s%d", compatFragData, idx)
								sym, found := t.FragData[idx]
								if !found {
									sym = &ast.VariableSym{SymType: vec4, SymName: name}
								}
								t.FragData[idx] = sym
								return &ast.VarRefExpr{Sym: sym}
							}
						}
					}
				}
			}
		}
	case *ast.FunctionDecl:
		if t.depTextureFuncs() {
			// Texture-lookup functions got renamed.
			// Fix up old names to match the new names.
			var sym ast.Symbol
			switch n.Name() {
			case "texture1D", "texture1DProj", "texture1DLod", "texture1DProjLod",
				"shadow1D", "shadow1DProj", "shadow1DLod", "shadow1DProjLod":
				sym = parser.FindBuiltin(strings.Replace(n.Name(), "1D", "", -1)) // drop 1D

			case "texture2D", "texture2DProj", "texture2DLod", "texture2DProjLod",
				"shadow2D", "shadow2DProj", "shadow2DLod", "shadow2DProjLod":
				sym = parser.FindBuiltin(strings.Replace(n.Name(), "2D", "", -1)) // drop 2D

			case "texture3D", "texture3DProj", "texture3DLod", "texture3DProjLod":
				// TODO: For the proj versions, the texture coordinate is divided by coord.q.
				sym = parser.FindBuiltin(strings.Replace(n.Name(), "3D", "", -1)) // drop 3D

			case "textureCube", "textureCubeLod":
				sym = parser.FindBuiltin(strings.Replace(n.Name(), "Cube", "", -1)) // drop Cube
			}
			if v, ok := sym.(ast.ValueSymbol); ok {
				return v // replace usage
			}
		}
	case *ast.VariableSym:
		fixReservedNames(&n.SymName)
		if n.Quals != nil {
			switch {
			case n.Quals.Storage == ast.StorAttribute && t.attributeToIn():
				n.Quals.Storage = ast.StorIn
			case n.Quals.Storage == ast.StorVarying && t.varyingToOut():
				n.Quals.Storage = ast.StorOut
			case n.Quals.Storage == ast.StorVarying && t.varyingToIn():
				n.Quals.Storage = ast.StorIn
			}
		}
		if t.declareFragColor() {
			if n.Name() == "gl_FragColor" {
				if t.FragColor == nil {
					t.FragColor = &ast.VariableSym{SymType: vec4, SymName: compatFragColor}
				}
				return t.FragColor // replace usage
			}
		}
	case *ast.FuncParameterSym:
		fixReservedNames(&n.SymName)
	case *ast.BinaryExpr:
		if left, ok := n.Left.(*ast.VarRefExpr); ok && left.Sym == t.FragColor {
			// Force cast RHS to vec4.
			// TODO: Only add cast if necessary.
			n.Right = &ast.CallExpr{
				Args:   []ast.Expression{n.Right},
				Callee: &ast.VarRefExpr{Sym: &ast.VariableSym{SymType: vec4, SymName: "vec4"}},
			}
		}
	}

	return child
}

func glslCompat(src string, lang ast.Language, device *service.Device) (string, error) {
	tree, _, errs := glsl.Parse(src, lang)
	if len(errs) > 0 {
		return "", fmt.Errorf("Failed to parse shader source:\n%s\n%s", src, errs)
	}

	deviceGLVersion, err := ParseVersion(device.Version)
	if err != nil {
		return "", fmt.Errorf("Could not parse GL version from '%s'. Error: %v",
			device.Version, err)
	}

	deviceGLSLVersion, err := GLSLVersion(device.Version)
	if err != nil {
		return "", fmt.Errorf("Could not get GLSL version from GL version: '%s'. Error: %v",
			device.Version, err)
	}

	// TODO: set targetGLSLVersion to a version closest to the original source
	// version, while still being maintained by the target device.
	targetGLSLVersion := deviceGLSLVersion

	transform := glslTransform(0)
	if !deviceGLVersion.IsES { // Strip any precision specifiers
		transform |= glslStripPrecision
	}

	if targetGLSLVersion.GreaterThan(1, 2) {
		switch lang {
		case ast.LangVertexShader:
			transform |= glslAttributeToIn
			transform |= glslVaryingToOut
		case ast.LangFragmentShader:
			transform |= glslVaryingToIn
			transform |= glslDeclareFragColor
			transform |= glslDeclareFragData
		}
		transform |= glslDepTextureFuncs
	}

	if transform == 0 {
		return src, nil
	}

	transform.apply(tree)

	return glsl.Format(tree, targetGLSLVersion), nil
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
