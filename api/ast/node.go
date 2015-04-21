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

package ast

import "android.googlesource.com/platform/tools/gpu/parse"

// Node represents any AST-node type.
type Node interface {
	// Fragment returns the CST parse Fragment for this AST node.
	Fragment() parse.Fragment
}

func (t Alias) Fragment() parse.Fragment            { return t.CST }
func (t Annotation) Fragment() parse.Fragment       { return t.CST }
func (t API) Fragment() parse.Fragment              { return t.CST }
func (t ArrayType) Fragment() parse.Fragment        { return t.CST }
func (t Assert) Fragment() parse.Fragment           { return t.CST }
func (t Assign) Fragment() parse.Fragment           { return t.CST }
func (t BinaryOp) Fragment() parse.Fragment         { return t.CST }
func (t Block) Fragment() parse.Fragment            { return t.CST }
func (t Bool) Fragment() parse.Fragment             { return t.CST }
func (t Branch) Fragment() parse.Fragment           { return t.CST }
func (t Call) Fragment() parse.Fragment             { return t.CST }
func (t Case) Fragment() parse.Fragment             { return t.CST }
func (t Cast) Fragment() parse.Fragment             { return t.CST }
func (t Class) Fragment() parse.Fragment            { return t.CST }
func (t ClassInitializer) Fragment() parse.Fragment { return t.CST }
func (t DeclareLocal) Fragment() parse.Fragment     { return t.CST }
func (t Enum) Fragment() parse.Fragment             { return t.CST }
func (t EnumEntry) Fragment() parse.Fragment        { return t.CST }
func (t Field) Fragment() parse.Fragment            { return t.CST }
func (t FieldInitializer) Fragment() parse.Fragment { return t.CST }
func (t Function) Fragment() parse.Fragment         { return t.CST }
func (t Group) Fragment() parse.Fragment            { return t.CST }
func (t Identifier) Fragment() parse.Fragment       { return t.CST }
func (t Index) Fragment() parse.Fragment            { return t.CST }
func (t Iteration) Fragment() parse.Fragment        { return t.CST }
func (t Length) Fragment() parse.Fragment           { return t.CST }
func (t MapType) Fragment() parse.Fragment          { return t.CST }
func (t Member) Fragment() parse.Fragment           { return t.CST }
func (t New) Fragment() parse.Fragment              { return t.CST }
func (t Number) Fragment() parse.Fragment           { return t.CST }
func (t Parameter) Fragment() parse.Fragment        { return t.CST }
func (t PointerType) Fragment() parse.Fragment      { return t.CST }
func (t Pseudonym) Fragment() parse.Fragment        { return t.CST }
func (t Return) Fragment() parse.Fragment           { return t.CST }
func (t StaticArrayType) Fragment() parse.Fragment  { return t.CST }
func (t String) Fragment() parse.Fragment           { return t.CST }
func (t Switch) Fragment() parse.Fragment           { return t.CST }
func (t UnaryOp) Fragment() parse.Fragment          { return t.CST }
func (t Unknown) Fragment() parse.Fragment          { return t.CST }
func (t Null) Fragment() parse.Fragment             { return t.CST }

func (t Invalid) Fragment() parse.Fragment { return nil }
