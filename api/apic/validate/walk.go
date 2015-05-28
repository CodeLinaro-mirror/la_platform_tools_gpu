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

package validate

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/api/semantic"
)

func walk(n semantic.Node, v func(semantic.Node) bool) {
	if !v(n) {
		return
	}

	switch n := n.(type) {
	case *semantic.API:
		for _, c := range n.Members {
			walk(c, v)
		}
	case *semantic.Slice:
		walk(n.To, v)
	case *semantic.SliceIndex:
		walk(n.Slice, v)
		walk(n.Index, v)
	case *semantic.Assert:
		walk(n.Condition, v)
	case *semantic.Assign:
		walk(n.LHS, v)
		walk(n.RHS, v)
	case *semantic.Annotation:
		for _, c := range n.Arguments {
			walk(c, v)
		}
	case *semantic.Block:
		for _, c := range n.Statements {
			walk(c, v)
		}
	case semantic.BoolValue:
	case *semantic.BinaryOp:
		walk(n.LHS, v)
		walk(n.RHS, v)
	case *semantic.Branch:
		walk(n.Condition, v)
		walk(n.True, v)
		if n.False != nil {
			walk(n.False, v)
		}
	case *semantic.Builtin:
	case *semantic.Reference:
		walk(n.To, v)
	case *semantic.Call:
		walk(n.Type, v)
		walk(n.Target, v)
		for _, a := range n.Arguments {
			walk(a, v)
		}
	case *semantic.Callable:
		if n.Object != nil {
			walk(n.Object, v)
		}
	case *semantic.Case:
		for _, c := range n.Conditions {
			walk(c, v)
		}
		walk(n.Block, v)
	case *semantic.Cast:
		walk(n.Object, v)
		walk(n.Type, v)
	case *semantic.Class:
		for _, e := range n.Extends {
			walk(e, v)
		}
		for _, f := range n.Fields {
			walk(f, v)
		}
		for _, m := range n.Methods {
			walk(m, v)
		}
	case *semantic.ClassInitializer:
		for _, f := range n.Fields {
			walk(f, v)
		}
	case *semantic.Choice:
		for _, c := range n.Conditions {
			walk(c, v)
		}
		walk(n.Expression, v)
	case *semantic.DeclareLocal:
		walk(n.Local, v)
	case *semantic.Enum:
		for _, e := range n.Extends {
			walk(e, v)
		}
		for _, e := range n.Entries {
			walk(e, v)
		}
	case *semantic.EnumEntry:
	case *semantic.Pseudonym:
		walk(n.To, v)
		for _, m := range n.Methods {
			walk(m, v)
		}
	case *semantic.Field:
		walk(n.Type, v)
		if n.Default != nil {
			walk(n.Default, v)
		}
	case *semantic.FieldInitializer:
		walk(n.Value, v)
	case semantic.Float32Value:
	case semantic.Float64Value:
	case *semantic.Function:
		v(n.Docs)
		walk(n.Return, v)
		for _, c := range n.FullParameters {
			walk(c, v)
		}
		walk(n.Block, v)
		walk(n.Signature, v)
	case *semantic.Parameter:
		for _, c := range n.Annotations {
			walk(c, v)
		}
		walk(n.Type, v)
	case *semantic.Global:
	case *semantic.StaticArray:
	case *semantic.Signature:
	case semantic.Int8Value:
	case semantic.Int16Value:
	case semantic.Int32Value:
	case semantic.Int64Value:
	case *semantic.Iteration:
		walk(n.Iterator, v)
		walk(n.Iterable, v)
		walk(n.Block, v)
	case *semantic.Length:
		walk(n.Object, v)
	case *semantic.Local:
		walk(n.Type, v)
		if n.Value != nil {
			walk(n.Value, v)
		}
	case *semantic.Map:
		walk(n.KeyType, v)
		walk(n.ValueType, v)
	case *semantic.MapAssign:
		walk(n.To, v)
		walk(n.Value, v)
	case *semantic.MapContains:
		walk(n.Key, v)
		walk(n.Map, v)
	case *semantic.MapIndex:
		walk(n.Map, v)
		walk(n.Index, v)
	case *semantic.Member:
		walk(n.Object, v)
		walk(n.Field, v)
	case *semantic.Pointer:
		walk(n.To, v)
	case *semantic.Return:
		if n.Value != nil {
			walk(n.Value, v)
		}
	case *semantic.Select:
		walk(n.Value, v)
		for _, c := range n.Choices {
			walk(c, v)
		}
	case *semantic.Switch:
		walk(n.Value, v)
		for _, c := range n.Cases {
			walk(c, v)
		}
	case semantic.Uint8Value:
	case semantic.Uint16Value:
	case semantic.Uint32Value:
	case semantic.Uint64Value:
	case *semantic.Unknown:
	case *semantic.Clone:
		walk(n.Slice, v)
	case *semantic.Copy:
		walk(n.Src, v)
		walk(n.Dst, v)
	case *semantic.Create:
	case *semantic.Ignore:
	case *semantic.Make:
		walk(n.Size, v)
	case semantic.Null:
	case *semantic.PointerRange:
		walk(n.Pointer, v)
	case *semantic.Read:
		walk(n.Slice, v)
	case *semantic.SliceRange:
		walk(n.Slice, v)
	case *semantic.Write:
		walk(n.Slice, v)
	default:
		panic(fmt.Errorf("Unsupported semantic node type %T", n))
	}
}
