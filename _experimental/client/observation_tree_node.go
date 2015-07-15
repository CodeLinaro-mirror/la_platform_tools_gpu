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

package client

import (
	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/memory"
	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

type observationItem struct {
	atomID atom.ID
	isRead bool
	index  int
}

// observationTreeNode is a gxui.TreeNode representing an atom's memory
// observations. It has the item of type observationsItem.
type observationTreeNode struct {
	ctx  *commandAdapterCtx
	item observationItem
}

func (n observationTreeNode) Count() int                          { return 0 }
func (n observationTreeNode) NodeAt(index int) gxui.TreeNode      { return nil }
func (n observationTreeNode) Item() gxui.AdapterItem              { return n.item }
func (n observationTreeNode) ItemIndex(item gxui.AdapterItem) int { return -1 }

func (n observationTreeNode) Create(theme gxui.Theme) gxui.Control {
	atom := n.ctx.atoms[n.item.atomID]
	observations := atom.Observations()
	var r memory.Range
	var c gxui.Color
	if n.item.isRead {
		r = observations.Reads[n.item.index].Range
		c = gxui.Green
	} else {
		r = observations.Writes[n.item.index].Range
		c = gxui.Red
	}

	b := theme.CreateButton()
	b.SetMargin(math.Spacing{})
	b.AddChild(createLabel(n.ctx.appCtx, r.String(), c))
	// ptr := memory.Pointer{Address: r.Base, Pool: memory.ApplicationPool}
	// b.OnClick(func(gxui.MouseEvent) { n.appCtx.SelectPointer(ptr) }) // [BENC]: TODO
	return b
}
