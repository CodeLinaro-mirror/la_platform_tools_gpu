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

import (
	"fmt"
	"sort"
)

// A type for binary operators of the language.
type BinaryOp struct {
	op *string
}

func (op BinaryOp) String() string { return *op.op }

// All operators of the language, in an array, sorted according to string length (longest first).
// Do not modify.
var Operators []fmt.Stringer

func appendBinaryOperator(opname string) BinaryOp {
	op := BinaryOp{&opname}
	Operators = append(Operators, op)
	return op
}

// A type for unary operators of the language.
type UnaryOp BinaryOp

func (o UnaryOp) String() string { return BinaryOp(o).String() }

func appendUnaryOperator(opname string) UnaryOp {
	op := UnaryOp{&opname}
	Operators = append(Operators, op)
	return op
}

var prefixOps = map[string]UnaryOp{}

func appendPrefixOp(opname string) UnaryOp {
	op := appendUnaryOperator(opname)
	prefixOps[opname] = op
	return op
}

// GetPrefixOp returns the prefix unary operator associated with the given string. For "++" it
// returns UoPreinc instead of UoPostinc. If an operator is not found, the second result is
// false.
func GetPrefixOp(opname string) (op UnaryOp, present bool) { op, present = prefixOps[opname]; return }

var assignmentOps = map[BinaryOp]struct{}{}

func appendAssignOp(opname string) BinaryOp {
	op := appendBinaryOperator(opname)
	assignmentOps[op] = struct{}{}
	return BinaryOp(op)
}

// Is op an assignment operator?
func IsAssignmentOp(op BinaryOp) bool { _, present := assignmentOps[op]; return present }

// All the operators of the language, one variable per operator.
var (
	BoAdd    = appendBinaryOperator("+")
	BoBand   = appendBinaryOperator("&")
	BoBor    = appendBinaryOperator("|")
	BoBxor   = appendBinaryOperator("^")
	BoComma  = appendBinaryOperator(",")
	BoDiv    = appendBinaryOperator("/")
	BoEq     = appendBinaryOperator("==")
	BoLand   = appendBinaryOperator("&&")
	BoLess   = appendBinaryOperator("<")
	BoLessEq = appendBinaryOperator("<=")
	BoLor    = appendBinaryOperator("||")
	BoLxor   = appendBinaryOperator("^^")
	BoMod    = appendBinaryOperator("%")
	BoMore   = appendBinaryOperator(">")
	BoMoreEq = appendBinaryOperator(">=")
	BoMul    = appendBinaryOperator("*")
	BoNotEq  = appendBinaryOperator("!=")
	BoShl    = appendBinaryOperator("<<")
	BoShr    = appendBinaryOperator(">>")
	BoSub    = appendBinaryOperator("-")

	BoAssign    = appendAssignOp("=")
	BoAddAssign = appendAssignOp("+=")
	BoAndAssign = appendAssignOp("&=")
	BoDivAssign = appendAssignOp("/=")
	BoModAssign = appendAssignOp("%=")
	BoMulAssign = appendAssignOp("*=")
	BoOrAssign  = appendAssignOp("|=")
	BoShlAssign = appendAssignOp("<<=")
	BoShrAssign = appendAssignOp(">>=")
	BoSubAssign = appendAssignOp("-=")
	BoXorAssign = appendAssignOp("^=")

	UoBnot    = appendPrefixOp("~")
	UoLnot    = appendPrefixOp("!")
	UoMinus   = appendPrefixOp("-")
	UoPlus    = appendPrefixOp("+")
	UoPredec  = appendPrefixOp("--")
	UoPreinc  = appendPrefixOp("++")
	UoPostdec = appendUnaryOperator("--")
	UoPostinc = appendUnaryOperator("++")
)

type sortOperators []fmt.Stringer

func (s sortOperators) Len() int           { return len(s) }
func (s sortOperators) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortOperators) Less(i, j int) bool { return len(s[i].String()) > len(s[j].String()) }

func init() {
	// Sort operators according to their length, so that e.g. <<= gets picked up before <<
	sort.Sort(sortOperators(Operators))
}
