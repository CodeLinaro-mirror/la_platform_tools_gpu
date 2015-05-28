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

import (
	"strconv"

	"android.googlesource.com/platform/tools/gpu/api/ast"
	"android.googlesource.com/platform/tools/gpu/api/semantic"
)

func inferNumber(ctx *context, in *ast.Number, infer semantic.Type) semantic.Expression {
	var out semantic.Expression
	switch infer {
	case semantic.Int8Type:
		if v, err := strconv.ParseInt(in.Value, 0, 8); err == nil {
			out = semantic.Int8Value(v)
		}
	case semantic.Uint8Type:
		if v, err := strconv.ParseUint(in.Value, 0, 8); err == nil {
			out = semantic.Uint8Value(v)
		}
	case semantic.Int16Type:
		if v, err := strconv.ParseInt(in.Value, 0, 16); err == nil {
			out = semantic.Int16Value(v)
		}
	case semantic.Uint16Type:
		if v, err := strconv.ParseUint(in.Value, 0, 16); err == nil {
			out = semantic.Uint16Value(v)
		}
	case semantic.Int32Type:
		if v, err := strconv.ParseInt(in.Value, 0, 32); err == nil {
			out = semantic.Int32Value(v)
		}
	case semantic.Uint32Type:
		if v, err := strconv.ParseUint(in.Value, 0, 32); err == nil {
			out = semantic.Uint32Value(v)
		}
	case semantic.Int64Type:
		if v, err := strconv.ParseInt(in.Value, 0, 64); err == nil {
			out = semantic.Int64Value(v)
		}
	case semantic.Uint64Type:
		if v, err := strconv.ParseUint(in.Value, 0, 64); err == nil {
			out = semantic.Uint64Value(v)
		}
	case semantic.Float32Type:
		if v, err := strconv.ParseFloat(in.Value, 32); err == nil {
			out = semantic.Float32Value(v)
		}
	case semantic.Float64Type:
		if v, err := strconv.ParseFloat(in.Value, 64); err == nil {
			out = semantic.Float64Value(v)
		}
	default:
		return nil
	}
	ctx.mappings[in] = out
	return out
}

func inferUnknown(ctx *context, lhs semantic.Node, rhs semantic.Node) {
	u := findUnknown(ctx, rhs)
	if u != nil && u.Inferred == nil {
		u.Inferred = lhsToObserved(ctx, lhs)
	}
}

func findUnknown(ctx *context, rhs semantic.Node) *semantic.Unknown {
	switch rhs := rhs.(type) {
	case *semantic.Unknown:
		return rhs
	case *semantic.Cast:
		return findUnknown(ctx, rhs.Object)
	case *semantic.Local:
		return findUnknown(ctx, rhs.Value)
	default:
		return nil
	}
}

// lhsToObserved takes an expression that is a valid lhs of an assignment, and
// creates a new expression that would read the observed output.
// This is used when attempting to infer the value an Unknown had from the
// observed outputs.
func lhsToObserved(ctx *context, lhs semantic.Node) semantic.Expression {
	switch lhs := lhs.(type) {
	case *semantic.SliceIndex:
		o := lhsToObserved(ctx, lhs.Slice)
		if o == nil {
			return nil
		}
		return &semantic.SliceIndex{Slice: o, Index: lhs.Index, Type: lhs.Type}
	case *semantic.PointerRange:
		o := lhsToObserved(ctx, lhs.Pointer)
		if o == nil {
			return nil
		}
		return &semantic.PointerRange{Pointer: o, Range: lhs.Range, Type: lhs.Type}
	case *semantic.MapIndex:
		o := lhsToObserved(ctx, lhs.Map)
		if o == nil {
			return nil
		}
		return &semantic.MapIndex{Map: o, Index: lhs.Index, Type: lhs.Type}
	case *semantic.Parameter:
		return &semantic.Observed{Parameter: lhs}
	case *semantic.Local:
		return lhsToObserved(ctx, lhs.Value)
	default:
		ctx.errorf(lhs, "Cannot infer unknown from %T", lhs)
		return nil
	}
}
