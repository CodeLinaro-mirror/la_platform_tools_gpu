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

const (
	// Keyword strings represent places in the syntax where a word has special
	// meaning.
	KeywordAPI       = "api"
	KeywordAlias     = "alias"
	KeywordArray     = "array"
	KeywordAs        = "as"
	KeywordAssert    = "assert"
	KeywordBitfield  = "bitfield"
	KeywordCase      = "case"
	KeywordClass     = "class"
	KeywordCmd       = "cmd"
	KeywordElse      = "else"
	KeywordEnum      = "enum"
	KeywordExtern    = "extern"
	KeywordFalse     = "false"
	KeywordFor       = "for"
	KeywordIf        = "if"
	KeywordIn        = "in"
	KeywordInout     = "inout"
	KeywordLength    = "len"
	KeywordMacro     = "macro"
	KeywordMap       = "map"
	KeywordNew       = "new"
	KeywordOut       = "out"
	KeywordPseudonym = "type"
	KeywordSwitch    = "switch"
	KeywordThis      = "this"
	KeywordTrue      = "true"
	KeywordWhen      = "when"
)
