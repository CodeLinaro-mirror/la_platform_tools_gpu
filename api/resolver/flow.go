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

type fenceInfo struct {
	pre   bool
	post  bool
	fence *semantic.Fence
}

func (f *fenceInfo) merge(child fenceInfo) {
	f.pre = f.pre || child.pre
	f.post = f.post || child.post
	if f.fence == nil {
		f.fence = child.fence
	}
}

func addFence(ctx *context, block *semantic.Block) {
	t := detectFence(ctx, block, nil, true)
	if t.fence == nil {
		block.Statements = append(block.Statements, &semantic.Fence{})
	}
}

func detectFence(ctx *context, n semantic.Node, fence *semantic.Fence, istop bool) fenceInfo {
	switch n := n.(type) {
	case *semantic.Block:
		info := fenceInfo{fence: fence, pre: istop}
		for index, s := range n.Statements {
			child := detectFence(ctx, s, info.fence, false)
			info.merge(child)
			// if we are both pre and post without a fence, we have reached a boundary
			if info.pre && info.post && (info.fence == nil) {
				info.fence = &semantic.Fence{Statement: s}
				n.Statements[index] = info.fence
			}
		}
		return info
	case *semantic.Branch, *semantic.Select, *semantic.Switch, *semantic.Iteration:
		info := detectFenceChildren(ctx, n, fence)
		if fence == nil && info.fence != nil {
			ctx.errorf(info.fence.Statement, "fence not allowed in %T", n)
		}
		return info
	case *semantic.Copy:
		if fence != nil {
			ctx.errorf(n, "copy after fence")
		}
		info := detectFenceChildren(ctx, n, fence)
		// copy is both pre and post
		info.pre = true
		info.post = true
		return info
	case *semantic.Read, *semantic.SliceIndex:
		if fence != nil {
			ctx.errorf(n, "%T after fence", n)
		}
		info := detectFenceChildren(ctx, n, fence)
		info.pre = true
		return info
	case *semantic.SliceAssign:
		// Need to not directly test the slice index as it will look like a read
		info := detectFenceChildren(ctx, n.To, fence)
		info.merge(detectFence(ctx, n.Value, info.fence, false))
		info.post = true
		return info
	case *semantic.Write, *semantic.Unknown, *semantic.Return:
		info := detectFenceChildren(ctx, n, fence)
		info.post = true
		return info
	case semantic.Type:
		return fenceInfo{fence: fence}
	default:
		return detectFenceChildren(ctx, n, fence)
	}
}

func detectFenceChildren(ctx *context, n semantic.Node, fence *semantic.Fence) fenceInfo {
	info := fenceInfo{fence: fence}
	semantic.Visit(n, func(v semantic.Node) {
		info.merge(detectFence(ctx, v, info.fence, false))
	})
	return info
}
