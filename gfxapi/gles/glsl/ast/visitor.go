// Copyright (C) 2014 The Android Open Source Project
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

package ast

// ChildVisitor is a callback function used by VisitChildren to visit nodes.
type ChildVisitor func(interface{})

// VisitChildren is a helper function which calls the provided callback for each child node of
// the argument.
func VisitChildren(n interface{}, v ChildVisitor) {
	switch n := n.(type) {
	case *BinaryExpr:
		v(n.Left)
		v(n.Right)
	case *IndexExpr:
		v(n.Base)
		v(n.Index)
	case *ConditionalExpr:
		v(n.Cond)
		v(n.TrueExpr)
		v(n.FalseExpr)
	case *UnaryExpr:
		v(n.Expr)
	case *DotExpr:
		v(n.Expr)
	case *ConstantExpr, *VarRefExpr:
	case *TypeConversionExpr:
		v(n.RetType)
	case *ParenExpr:
		v(n.Expr)
	case *CallExpr:
		v(n.Callee)
		for _, a := range n.Args {
			v(a)
		}

	case *DeclarationStmt:
		v(n.Decl)
	case *ExpressionStmt:
		v(n.Expr)
	case *CaseStmt:
		v(n.Expr)
	case *SwitchStmt:
		v(n.Expr)
		v(n.Stmts)
	case *WhileStmt:
		v(n.Cond)
		v(n.Stmt)
	case *DoStmt:
		v(n.Stmt)
		v(n.Expr)
	case *ForStmt:
		v(n.Init)
		v(n.Cond)
		v(n.Loop)
		v(n.Body)
	case *ContinueStmt, *BreakStmt, *DiscardStmt, *EmptyStmt, *DefaultStmt:
	case *IfStmt:
		v(n.IfExpr)
		v(n.ThenStmt)
		if n.ElseStmt != nil {
			v(n.ElseStmt)
		}
	case *CompoundStmt:
		for _, s := range n.Stmts {
			v(s)
		}
	case *ReturnStmt:
		if n.Expr != nil {
			v(n.Expr)
		}

	case *BuiltinType:
	case *ArrayType:
		v(n.Base)
		if n.Size != nil {
			v(n.Size)
		}
	case *StructType:
		if n.StructDef {
			v(n.Sym)
		}

	case *LayoutQualifier:
	case *TypeQualifiers:
		if n.Layout != nil {
			v(n.Layout)
		}

	case *PrecisionDecl:
		v(n.Type)
	case *FunctionDecl:
		v(n.RetType)
		for _, p := range n.Params {
			v(p)
		}
		if n.Stmts != nil {
			v(n.Stmts)
		}
	case *MultiVarDecl:
		if n.Quals != nil {
			v(n.Quals)
		}
		v(n.Type)
		for _, n := range n.Vars {
			v(n)
		}
	case *LayoutDecl:
		v(n.Layout)
	case *InvariantDecl:
		for _, n := range n.Vars {
			v(n)
		}
	case *UniformDecl:
		v(n.Block)
		if n.Size != nil {
			v(n.Size)
		}

	case *VariableSym:
		v(n.Type)
		if n.Init != nil {
			v(n.Init)
		}
	case *FuncParameterSym:
		v(n.SymType)
	case *StructSym:
		for _, n := range n.Vars {
			v(n)
		}
	case *UniformBlock:
		if n.Layout != nil {
			v(n.Layout)
		}
		for _, n := range n.Vars {
			v(n)
		}

	case *ExpressionCond:
		v(n.Expr)
	case *VarDeclCond:
		v(n.Sym)

	case *Ast:
		for _, n := range n.Decls {
			v(n)
		}
	}
}
