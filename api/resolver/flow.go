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

type fenceTracker struct {
	ctx  *context
	post bool
}

func addFence(ctx *context, block *semantic.Block) {
	t := fenceTracker{ctx: ctx}
	if !t.boundary(block) {
		block.Statements = append(block.Statements, &semantic.Fence{})
	}
}

func (t *fenceTracker) scan(n semantic.Node) {
	switch n := n.(type) {
	case *semantic.Block:
		for i := 0; i < len(n.Statements); i++ {
			s := n.Statements[i]
			if t.boundary(s) {
				// first post operation detected, insert fence marker here
				n.Statements = append(n.Statements, nil)
				copy(n.Statements[i+1:], n.Statements[i:])
				n.Statements[i] = &semantic.Fence{}
				// step past the newly inserted fence
				i++
				// Continue calling boundary() on later statements to detect post-after-pre errors.
			}
		}
	case *semantic.Branch:
		t.scan(n.Condition)
		if t.boundary(n.True) {
			t.ctx.errorf(n.True, "pre-post boundary in true condition")
		}
		if n.False != nil {
			if t.boundary(n.False) {
				t.ctx.errorf(n.False, "pre-post boundary in condition")
			}
		}
	case *semantic.Select, *semantic.Switch, *semantic.Iteration:
		ct := *t
		semantic.Visit(n, t.scan)
		if ct.post && !t.post {
			t.post = true
			t.ctx.errorf(n, "pre-post boundary in %T", n)
		}
	case *semantic.Copy:
		semantic.Visit(n, t.scan)
		if t.post {
			t.ctx.errorf(n, "copy after fence")
		}
		t.post = true
	case *semantic.Read:
		semantic.Visit(n, t.scan)
		if t.post {
			t.ctx.errorf(n, "read after fence")
		}
	case *semantic.Unknown:
		t.post = true
	case *semantic.Write:
		semantic.Visit(n, t.scan)
		t.post = true
	case *semantic.SliceIndex:
		semantic.Visit(n, t.scan)
		if t.post {
			t.ctx.errorf(n, "slice index after fence")
		}
	case *semantic.SliceAssign:
		semantic.Visit(n.To, t.scan)
		t.scan(n.Value)
		t.post = true
	default:
		semantic.Visit(n, t.scan)
	}
}

func (t *fenceTracker) boundary(n semantic.Node) bool {
	ct := *t
	ct.scan(n)
	if ct.post && !t.post {
		t.post = true
		return true
	}
	return false
}
