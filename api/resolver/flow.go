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

type order int

const (
	pre  = order(1)
	post = order(2)
)

func (o order) pre() bool  { return (o & pre) != 0 }
func (o order) post() bool { return (o & post) != 0 }

type fenceTracker struct {
	ctx    *context
	orders map[semantic.Node]order
}

func addFence(ctx *context, block *semantic.Block) {
	t := fenceTracker{ctx: ctx, orders: map[semantic.Node]order{}}
	t.order(block)
	if !t.insertFence(block) {
		block.Statements = append(block.Statements, &semantic.Fence{})
	}
}

func (t *fenceTracker) order(n semantic.Node) order {
	o := order(0)
	switch n := n.(type) {
	case *semantic.Read:
		o |= pre
	case *semantic.Unknown:
		o |= post
	case *semantic.Write:
		o |= post | t.order(n.Slice)
	case *semantic.Copy:
		o |= pre | t.order(n.Src)
	case *semantic.SliceIndex:
		o |= pre | t.order(n.Index) | t.order(n.Slice)
	case *semantic.SliceAssign:
		o |= post | t.order(n.Value)
	default:
		semantic.Visit(n, func(c semantic.Node) { o |= t.order(c) })
	}
	t.orders[n] = o
	return o
}

func (t *fenceTracker) insertFence(n semantic.Node) (inserted bool) {
	switch n := n.(type) {
	case *semantic.Block:
		for i := 0; i < len(n.Statements); i++ {
			s := n.Statements[i]
			o := t.orders[s]
			switch {
			case inserted:
				if o.pre() {
					t.ctx.errorf(s, "pre-statement after fence")
					return true
				}
			case o.post():
				if !o.pre() || !t.insertFence(s) {
					// first post operation detected, insert fence marker here
					n.Statements = append(n.Statements, nil)
					copy(n.Statements[i+1:], n.Statements[i:])
					n.Statements[i] = &semantic.Fence{}
					i++ // step past the newly inserted fence
					inserted = true
				}
			}
		}
		return inserted
	case *semantic.Iteration, *semantic.Switch, *semantic.Branch:
		t.ctx.errorf(n, "fence not permitted in %T", n)
		return true
	default:
		return false
	}
}
