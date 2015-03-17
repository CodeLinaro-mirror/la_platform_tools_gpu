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

// Package limits is used to calculate the possible values for a variable.
package limits

import "android.googlesource.com/platform/tools/gpu/api/semantic"

// Limits represent the possible values for a given type.
type Limits interface {
	// Binary returns the unary operator op performed with the limits.
	Unary(op string) Limits
	// Binary returns the binary operator op performed with the limits and rhs.
	Binary(op string, rhs Limits) Limits
}

// Unbound returns the unbound limits for the given semantic type.
func Unbound(ty semantic.Type) Limits {
	switch ty {
	case semantic.BoolType:
		return Maybe
	case semantic.Uint8Type:
		return newUint(0, (1<<8)-1)
	case semantic.Uint16Type:
		return newUint(0, (1<<16)-1)
	case semantic.Uint32Type:
		return newUint(0, (1<<32)-1)
	case semantic.Uint64Type, semantic.UintType:
		return newUint(0, (1<<64)-1)
	}
	return nil
}
