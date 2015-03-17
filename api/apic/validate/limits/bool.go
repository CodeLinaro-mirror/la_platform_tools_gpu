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

package limits

import "android.googlesource.com/platform/tools/gpu/api/ast"

type boolLimit int

const (
	False = boolLimit(iota) // A boolean limit that is false.
	True                    // A boolean limit that is true.
	Maybe                   // A boolean limit that can be either true of false.
)

func (l boolLimit) Unary(operator string) Limits {
	switch operator {
	case ast.OpNot:
		switch l {
		case False:
			return True
		case True:
			return False
		default:
			return Maybe
		}

	default:
		return Maybe
	}
}

func (a boolLimit) Binary(operator string, rhs Limits) Limits {
	b := rhs.(boolLimit)
	switch operator {
	case ast.OpAnd:
		switch {
		case a == False, b == False:
			return False
		case a == Maybe, b == Maybe:
			return Maybe
		default:
			return True
		}

	case ast.OpOr:
		switch {
		case a == True, b == True:
			return True
		case a == Maybe, b == Maybe:
			return Maybe
		default:
			return False
		}

	case ast.OpEQ:
		switch {
		case a == Maybe, b == Maybe:
			return Maybe
		default:
			if a == b {
				return True
			} else {
				return False
			}
		}

	default:
		return Maybe
	}
}
