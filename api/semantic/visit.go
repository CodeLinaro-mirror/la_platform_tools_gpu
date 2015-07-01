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

package semantic

import "fmt"

// Visit invokes visitor for all the children of the supplied node.
func Visit(node Node, visitor func(Node)) {
	switch n := node.(type) {
	case *API:
		(*Symbols)(&n.members).Visit(func(_ string, n Node) { visitor(n) })
	case *ArrayAssign:
		visitor(n.To)
		visitor(n.Value)
	case *ArrayIndex:
		visitor(n.Array)
		visitor(n.Index)
	case *ArrayInitializer:
		visitor(n.Array)
		for _, c := range n.Values {
			visitor(c)
		}
	case *Slice:
		visitor(n.To)
	case *SliceIndex:
		visitor(n.Slice)
		visitor(n.Index)
	case *SliceAssign:
		visitor(n.To)
		visitor(n.Value)
	case *Assert:
		visitor(n.Condition)
	case *Assign:
		visitor(n.LHS)
		visitor(n.RHS)
	case *Annotation:
		for _, c := range n.Arguments {
			visitor(c)
		}
	case *Block:
		for _, c := range n.Statements {
			visitor(c)
		}
	case BoolValue:
	case *BinaryOp:
		visitor(n.LHS)
		visitor(n.RHS)
	case *BitTest:
		visitor(n.Bitfield)
		visitor(n.Bits)
	case *UnaryOp:
		visitor(n.Expression)
	case *Branch:
		visitor(n.Condition)
		visitor(n.True)
		if n.False != nil {
			visitor(n.False)
		}
	case *Builtin:
	case *Reference:
		visitor(n.To)
	case *Call:
		visitor(n.Type)
		visitor(n.Target)
		for _, a := range n.Arguments {
			visitor(a)
		}
	case *Callable:
		if n.Object != nil {
			visitor(n.Object)
		}
	case *Case:
		for _, c := range n.Conditions {
			visitor(c)
		}
		visitor(n.Block)
	case *Cast:
		visitor(n.Object)
		visitor(n.Type)
	case *Class:
		for _, f := range n.Fields {
			visitor(f)
		}
		for _, m := range n.Methods {
			visitor(m)
		}
	case *ClassInitializer:
		for _, f := range n.Fields {
			visitor(f)
		}
	case *Choice:
		for _, c := range n.Conditions {
			visitor(c)
		}
		visitor(n.Expression)
	case *DeclareLocal:
		visitor(n.Local)
	case *Enum:
		for _, e := range n.Extends {
			visitor(e)
		}
		for _, e := range n.Entries {
			visitor(e)
		}
	case *EnumEntry:
	case *Pseudonym:
		visitor(n.To)
		for _, m := range n.Methods {
			visitor(m)
		}
	case *Fence:
	case *Field:
		visitor(n.Type)
		if n.Default != nil {
			visitor(n.Default)
		}
	case *FieldInitializer:
		visitor(n.Value)
	case Float32Value:
	case Float64Value:
	case *Function:
		visitor(n.Docs)
		visitor(n.Return)
		for _, c := range n.FullParameters {
			visitor(c)
		}
		visitor(n.Block)
		visitor(n.Signature)
	case *Parameter:
		for _, c := range n.Annotations {
			visitor(c)
		}
		visitor(n.Type)
	case *Global:
	case *StaticArray:
	case *Signature:
	case Int8Value:
	case Int16Value:
	case Int32Value:
	case Int64Value:
	case *Iteration:
		visitor(n.Iterator)
		visitor(n.Iterable)
		visitor(n.Block)
	case *Length:
		visitor(n.Object)
	case *Local:
		visitor(n.Type)
		if n.Value != nil {
			visitor(n.Value)
		}
	case *Map:
		visitor(n.KeyType)
		visitor(n.ValueType)
	case *MapAssign:
		visitor(n.To)
		visitor(n.Value)
	case *MapContains:
		visitor(n.Key)
		visitor(n.Map)
	case *MapIndex:
		visitor(n.Map)
		visitor(n.Index)
	case *Member:
		visitor(n.Object)
		visitor(n.Field)
	case *Pointer:
		visitor(n.To)
	case *Return:
		if n.Value != nil {
			visitor(n.Value)
		}
	case *Select:
		visitor(n.Value)
		for _, c := range n.Choices {
			visitor(c)
		}
	case StringValue:
	case *Switch:
		visitor(n.Value)
		for _, c := range n.Cases {
			visitor(c)
		}
	case Uint8Value:
	case Uint16Value:
	case Uint32Value:
	case Uint64Value:
	case *Unknown:
	case *Clone:
		visitor(n.Slice)
	case *Copy:
		visitor(n.Src)
		visitor(n.Dst)
	case *Create:
	case *Ignore:
	case *Make:
		visitor(n.Size)
	case Null:
	case *PointerRange:
		visitor(n.Pointer)
	case *Read:
		visitor(n.Slice)
	case *SliceRange:
		visitor(n.Slice)
	case *Write:
		visitor(n.Slice)
	default:
		panic(fmt.Errorf("Unsupported semantic node type %T", n))
	}
}
