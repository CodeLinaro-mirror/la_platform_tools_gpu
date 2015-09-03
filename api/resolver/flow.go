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

// order is a bitfield describing whether a statement or block belongs before
// (pre) or after (post) the command's fence. Some statements and blocks
// straddle the fence, in which case both the pre and post bits will be set.
type order int

const (
	pre  = order(1) // statement is pre-fence
	post = order(2) // statement is post-fence
)

func (o order) isPre() bool  { return (o & pre) != 0 }
func (o order) isPost() bool { return (o & post) != 0 }

type fenceTracker struct {
	ctx      *context
	explicit bool // explcitly declared fence found.
	orders   map[semantic.Node]order
}

func addFence(ctx *context, block *semantic.Block) {
	t := fenceTracker{ctx: ctx, orders: map[semantic.Node]order{}}
	t.analyse(block)
	if !t.explicit && !t.insertFence(block) {
		block.Statements = append(block.Statements, &semantic.Fence{})
	}
}

// analyse traverses the statements and expressions for semantic node n and its
// children, assessing and returning whether each node is pre or post w.r.t the
// fence. An entry is added to the fenceTracker.orders map for each visited
// node.
func (t *fenceTracker) analyse(n semantic.Node) order {
	o := order(0)
	switch n := n.(type) {
	case *semantic.Fence:
		if !n.Explicit {
			t.ctx.icef(n.AST, "unexpected fence found")
		}
		if t.explicit {
			t.ctx.errorf(n, "multiple explicit fences found")
		}
		t.explicit = true
	case *semantic.Read:
		o = pre | t.analyse(n.Slice)
	case *semantic.Unknown:
		o = post
	case *semantic.Write:
		o = post | t.analyse(n.Slice)
	case *semantic.Copy:
		o = pre | post
	case *semantic.SliceIndex:
		o = pre | t.analyse(n.Index) | t.analyse(n.Slice)
	case *semantic.SliceAssign:
		o = post | t.analyse(n.Value)
	case *semantic.Return:
		o = post
	case semantic.Type:
	// Don't traverse types
	default:
		semantic.Visit(n, func(c semantic.Node) { o |= t.analyse(c) })
	}
	if o != 0 {
		t.orders[n] = o
	}
	return o
}

func (t *fenceTracker) insertFence(n semantic.Node) (inserted bool) {
	switch n := n.(type) {
	case *semantic.Block:
		for i := 0; i < len(n.Statements); i++ {
			s := n.Statements[i]
			o := t.orders[s]
			if inserted {
				// fence already found, but continue looking over the statements
				// to ensure that no more pre-statements are found.
				if o.isPre() {
					t.ctx.errorf(s, "pre-statement after fence")
					return true
				}
				continue
			}

			if o.isPost() {
				// first post statement or block containing post statement found.
				// insert a fence.
				switch {
				case !o.isPre(): // pre -> post transision between statements.
					n.Statements = append(n.Statements, nil)
					copy(n.Statements[i+1:], n.Statements[i:])
					n.Statements[i] = &semantic.Fence{}
					i++ // step past the newly inserted fence

				case t.insertFence(s): // inserted in sub-block

				default: // pre and post statement found.
					n.Statements[i] = &semantic.Fence{Statement: n.Statements[i]}
				}
				inserted = true
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
