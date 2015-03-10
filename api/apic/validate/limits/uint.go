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

import (
	"android.googlesource.com/platform/tools/gpu/api/ast"
	"android.googlesource.com/platform/tools/gpu/interval"
)

type uintLimit struct {
	spans interval.U64SpanList
}

func Uint(v uint64) Limits { return newUint(v, v) }

func newUint(minValue, maxValue uint64) uintLimit {
	return uintLimit{
		spans: interval.U64SpanList{
			interval.U64Span{
				Start: minValue,
				End:   maxValue + 1,
			},
		},
	}
}

func (l uintLimit) Unary(operator string) Limits {
	return Maybe
}

func (a uintLimit) Binary(operator string, rhs Limits) Limits {
	b := rhs.(uintLimit)
	switch operator {
	case ast.OpGT:
		a, b := a.span(), b.span()
		switch {
		case a.Start > b.End-1:
			return True
		case a.End-1 <= b.Start:
			return False
		default:
			return Maybe
		}

	case ast.OpGE:
		a, b := a.span(), b.span()
		switch {
		case a.Start >= b.End-1:
			return True
		case a.End-1 < b.Start:
			return False
		default:
			return Maybe
		}

	case ast.OpLT:
		a, b := a.span(), b.span()
		switch {
		case a.End-1 < b.Start:
			return True
		case a.Start >= b.End-1:
			return False
		default:
			return Maybe
		}

	case ast.OpLE:
		a, b := a.span(), b.span()
		switch {
		case a.End-1 <= b.Start:
			return True
		case a.Start > b.End-1:
			return False
		default:
			return Maybe
		}

	case ast.OpEQ:
		if len(a.spans) == 1 && len(b.spans) == 1 && // Only 1 interval in LHS and RHS
			a.spans[0].Start == a.spans[0].End-1 && // Only 1 value in LHS
			b.spans[0].Start == b.spans[0].End-1 && // Only 1 value in RHS
			a.spans[0].Start == b.spans[0].Start { // Value the same in LHS and RHS
			return True
		}
		for _, span := range b.spans {
			if _, c := interval.Intersect(&a.spans, span); c > 0 {
				return Maybe
			}
		}
		return False

	default:
		return Maybe
	}
}

func (a uintLimit) span() interval.U64Span {
	return interval.U64Span{
		Start: a.spans[0].Start,
		End:   a.spans[len(a.spans)-1].End,
	}
}
