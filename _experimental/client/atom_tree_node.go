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
	"github.com/google/gxui"
)

type atomItem struct {
	atomID atom.ID
}

// atomTreeNode is a gxui.TreeNode representing a single atom.
// It has the item of type atomItem.
type atomTreeNode struct {
	ctx  *commandAdapterCtx
	item atomItem
}

func (n atomTreeNode) Count() int {
	atom := n.ctx.atoms[n.item.atomID]
	observations := atom.Observations()
	return len(observations.Reads) + len(observations.Writes)
}

func (n atomTreeNode) NodeAt(index int) gxui.TreeNode {
	atom := n.ctx.atoms[n.item.atomID]
	observations := atom.Observations()
	if index < len(observations.Reads) {
		return observationTreeNode{
			ctx: n.ctx,
			item: observationItem{
				atomID: n.item.atomID,
				isRead: true,
				index:  index,
			},
		}
	} else {
		return observationTreeNode{
			ctx: n.ctx,
			item: observationItem{
				atomID: n.item.atomID,
				isRead: false,
				index:  index - len(observations.Reads),
			},
		}
	}
}

func (n atomTreeNode) Item() gxui.AdapterItem {
	return n.item
}

func (n atomTreeNode) ItemIndex(item gxui.AdapterItem) int {
	if i := item.(observationItem); i.isRead {
		return i.index
	} else {
		atom := n.ctx.atoms[n.item.atomID]
		observations := atom.Observations()
		return len(observations.Reads) + i.index
	}
}

func (n atomTreeNode) Create(t gxui.Theme) gxui.Control {
	p := n.ctx.capture.Atoms().Index(uint64(n.item.atomID))
	a := n.ctx.atoms[p.Index].(*Atom)

	layout := t.CreateLinearLayout()
	layout.SetDirection(gxui.LeftToRight)
	layout.AddChild(createLabel(n.ctx.appCtx, fmt.Sprintf("%.6d ", p.Index), LINE_NUMBER_COLOR))

	if ns, ok := n.ctx.timings.AtomDuration(n.item.atomID); ok {
		timeLbl := t.CreateLabel()
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

	nameLbl := createLabel(n.ctx.appCtx, atom.MetadataOf(a).DisplayName, COMMAND_COLOR)
	layout.AddChild(nameLbl)

	layout.AddChild(createLabel(n.ctx.appCtx, "(", CODE_COLOR))
	needcomma := false
	for i := 0; i < a.FieldCount(); i++ {
		argIdx := i // capture for closures
		f, v := a.Field(argIdx)
		p := p.Field(f.Name())
		c := createField(n.ctx.appCtx, p, f.Type, v)

		if needcomma {
			layout.AddChild(createLabel(n.ctx.appCtx, ", ", CODE_COLOR))
		}

		if c == nil {
			continue
		}

		layout.AddChild(c)
		needcomma = true
	}

	layout.AddChild(createLabel(n.ctx.appCtx, ")", CODE_COLOR))
	return layout
}
