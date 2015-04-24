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

package resolver

import "android.googlesource.com/platform/tools/gpu/api/semantic"

func assignable(lhs semantic.Type, rhs semantic.Type) bool {
	if isVoid(lhs) || isVoid(rhs) {
		return false
	}
	if lhs == semantic.AnyType || rhs == semantic.AnyType {
		return true
	}
	return lhs == rhs
}

func comparable(lhs semantic.Type, rhs semantic.Type) bool {
	if isVoid(lhs) || isVoid(rhs) {
		return false
	}
	if lhs == semantic.AnyType || rhs == semantic.AnyType {
		return true
	}
	return lhs == rhs
}

func equal(lhs semantic.Type, rhs semantic.Type) bool {
	return lhs == rhs
}

func enumExtends(t *semantic.Enum, base *semantic.Enum) bool {
	for _, e := range t.Extends {
		if equal(e, base) || enumExtends(e, base) {
			return true
		}
	}
	return false
}

func isNumber(t semantic.Type) bool {
	switch t {
	case semantic.IntType, semantic.UintType,
		semantic.Int8Type, semantic.Uint8Type,
		semantic.Int16Type, semantic.Uint16Type,
		semantic.Int32Type, semantic.Uint32Type,
		semantic.Int64Type, semantic.Uint64Type,
		semantic.Float32Type, semantic.Float64Type:
		return true
	default:
		return false
	}
}

func castable(from semantic.Type, to semantic.Type) bool {
	fromBase := baseType(from)
	toBase := baseType(to)
	if assignable(toBase, fromBase) {
		return true
	}
	fromEnum, fromIsEnum := fromBase.(*semantic.Enum)
	toEnum, toIsEnum := toBase.(*semantic.Enum)
	if fromIsEnum && toIsEnum {
		if enumExtends(toEnum, fromEnum) {
			return true // enum downcast
		}
		if enumExtends(fromEnum, toEnum) {
			return true // enum upcast, needed but unsafe
		}
	}
	fromIsNumber, toIsNumber := isNumber(fromBase), isNumber(toBase)
	if fromIsEnum && toIsNumber {
		return true // enum -> number
	}
	if fromIsNumber && toIsEnum {
		return true // number -> enum
	}
	_, fromIsPointer := fromBase.(*semantic.Pointer)
	if fromIsPointer && toIsNumber {
		return true // pointer -> number
	}
	if fromIsNumber && toIsNumber {
		return true // any numeric conversion
	}
	return false
}

func baseType(t semantic.Type) semantic.Type {
	for p, ok := t.(*semantic.Pseudonym); ok; p, ok = t.(*semantic.Pseudonym) {
		t = p.To
	}
	return t
}

func isVoid(t semantic.Type) bool {
	return baseType(t) == semantic.VoidType
}
