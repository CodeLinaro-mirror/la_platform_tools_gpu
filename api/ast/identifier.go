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

import "android.googlesource.com/platform/tools/gpu/parse"

// Identifier holds a parsed identifier in the parse tree.
type Identifier struct {
	CST   *parse.Leaf // underlying parse leaf for this node
	Value string      // the identifier
}

func (t Identifier) Fragment() parse.Fragment { return t.CST }

// Generic represents a identifier modified by type arguments. It looks like:
// «identifier ! ( arg | <arg {, arg} )>»
type Generic struct {
	CST       *parse.Branch // underlying parse structure for this node
	Name      *Identifier   // the generic identifier.
	Arguments []Node        // the type arguments to the generic.
}

func (t Generic) Fragment() parse.Fragment { return t.CST }

const (
	// Keyword strings represent places in the syntax where a word has special
	// meaning.
	KeywordAPI       = "api"
	KeywordAlias     = "alias"
	KeywordBitfield  = "bitfield"
	KeywordCase      = "case"
	KeywordClass     = "class"
	KeywordCmd       = "cmd"
	KeywordConst     = "const"
	KeywordDefine    = "define"
	KeywordElse      = "else"
	KeywordEnum      = "enum"
	KeywordExtern    = "extern"
	KeywordFalse     = "false"
	KeywordFor       = "for"
	KeywordIf        = "if"
	KeywordImport    = "import"
	KeywordIn        = "in"
	KeywordMacro     = "macro"
	KeywordNull      = "null"
	KeywordReturn    = "return"
	KeywordPseudonym = "type"
	KeywordSwitch    = "switch"
	KeywordThis      = "this"
	KeywordTrue      = "true"
	KeywordWhen      = "when"
)
