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
	"fmt"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/task"
	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

type hierarchyItem struct {
	rng atom.Range
}

// hierarchyTreeNode is a gxui.TreeNode representing a range of atoms.
// It has the item of type hierarchyItem.
type hierarchyTreeNode struct {
	ctx   *commandAdapterCtx
	group atom.Group
}

func (n hierarchyTreeNode) Count() int {
	return int(n.group.Count())
}

func (n hierarchyTreeNode) NodeAt(index int) gxui.TreeNode {
	if id, subgroup := n.group.Index(uint64(index)); subgroup != nil {
		return hierarchyTreeNode{
			ctx:   n.ctx,
			group: *subgroup,
		}
	} else {
		return atomTreeNode{
			ctx:  n.ctx,
			item: atomItem{atomIndex: id},
		}
	}
}

func (n hierarchyTreeNode) Item() gxui.AdapterItem {
	return hierarchyItem{rng: n.group.Range}
}

func (n hierarchyTreeNode) ItemIndex(item gxui.AdapterItem) int {
	var id uint64

	switch i := item.(type) {
	case hierarchyItem:
		id = i.rng.Start
	case atomItem:
		id = i.atomIndex
	case observationItem:
		id = i.atomIndex
	default:
		panic(fmt.Errorf("Unknown item type %T", i))
	}

	return int(n.group.IndexOf(id))
}

func (n hierarchyTreeNode) Create(theme gxui.Theme) gxui.Control {
	layout := theme.CreateLinearLayout()
	layout.SetDirection(gxui.LeftToRight)

	img := theme.CreateImage()
	img.SetExplicitSize(math.Size{W: kCommandAdapterItemHeight, H: kCommandAdapterItemHeight})
	img.SetAspectMode(gxui.AspectCorrectLetterbox)
	layout.AddChild(img)

	if ns, ok := n.ctx.timings.RangeDuration(n.group.Range); ok {
		timeLbl := theme.CreateLabel()
		milliseconds := float64(ns) / 1000000.
		timeLbl.SetText(fmt.Sprintf("%6.3f ms ", milliseconds))
		switch {
		case milliseconds >= 1.0:
			timeLbl.SetColor(gxui.ColorFromHex(0xFFFC19 + 0xFF<<24))
		case milliseconds >= 5.0:
			timeLbl.SetColor(gxui.ColorFromHex(0xD21212 + 0xFF<<24))
		}
		layout.AddChild(timeLbl)
	}

	atomIndex := n.group.Range.Last()
	appCtx := n.ctx.appCtx

	t := task.New()
	update := func() {
		t.Cancel()
		if n.ctx.device != nil {
			if flags := n.ctx.atoms[atomIndex].Flags(); flags.IsDrawCall() || flags.IsEndOfFrame() {
				after := n.ctx.capture.Atoms().Index(atomIndex)
				t.Run(updateFramebufferThumbnail{appCtx, n.ctx.device, after, img})
			}
		}
	}

	// TODO: When trees can be updated, this should listen for capture / device changes.
	img.OnAttach(update)
	img.OnDetach(t.Cancel)

	label := theme.CreateLabel()
	label.SetText(n.group.Name)
	layout.AddChild(label)

	return layout
}
