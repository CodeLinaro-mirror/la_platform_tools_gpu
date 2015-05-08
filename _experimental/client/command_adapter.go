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
	"strconv"

	"android.googlesource.com/platform/tools/gpu/_experimental/client/schema"
	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/service"
	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

const kCommandAdapterItemHeight = 18

func createEnumList(t gxui.Theme, appCtx *ApplicationContext, values interface{}, selected gxui.AdapterItem, active bool, onChange func(item gxui.AdapterItem)) gxui.DropDownList {
	a := gxui.CreateDefaultAdapter()
	a.SetItems(values)
	a.SetSizeAsLargest(t)
	a.SetStyleLabel(func(t gxui.Theme, l gxui.Label) {
		if active {
			l.SetColor(CONSTANT_COLOR)
		} else {
			l.SetColor(INACTIVE_COLOR)
		}
	})
	l := t.CreateDropDownList()
	l.SetAdapter(a)
	l.SetBubbleOverlay(appCtx.DropDownOverlay())
	l.Select(selected)
	l.OnSelectionChanged(onChange)
	return l
}

func createTextbox(t gxui.Theme, appCtx *ApplicationContext, value interface{}, active bool, onChange func(s string) bool) gxui.TextBox {
	tb := t.CreateTextBox()
	tb.SetMargin(math.Spacing{})
	tb.SetPadding(math.Spacing{})
	tb.SetText(fmt.Sprintf("%v", value))
	b := &gxui.TextBlock{Runes: []rune(tb.Text())}
	tb.SetDesiredWidth(tb.Font().Measure(b).W + 4)
	tb.OnTextChanged(func([]gxui.TextBoxEdit) {
		text := tb.Text()
		b := &gxui.TextBlock{Runes: []rune(text)}
		tb.SetDesiredWidth(tb.Font().Measure(b).W + 4)
		if onChange(text) {
			tb.SetTextColor(CONSTANT_COLOR)
		} else {
			tb.SetTextColor(gxui.Red)
		}
	})
	tb.OnLostFocus(func() {
		text := fmt.Sprintf("%v", value)
		if text != tb.Text() {
			tb.SetText(text)
		}
	})
	if active {
		tb.SetTextColor(CONSTANT_COLOR)
	} else {
		tb.SetTextColor(INACTIVE_COLOR)
	}
	return tb
}

func createAtomControls(t gxui.Theme, appCtx *ApplicationContext, id atom.ID) gxui.Control {
	a := appCtx.Atoms()[id]
	active := true

	ll := t.CreateLinearLayout()
	ll.SetDirection(gxui.LeftToRight)
	ll.AddChild(CreateLabel(t, fmt.Sprintf("%.6d ", id), LINE_NUMBER_COLOR, active))

	switch {
	case a.Info.IsCommand:
		if active {
			//appCtx.OnTimingInfoUpdated(func() {
			timeLbl := t.CreateLabel()
			milliseconds := float64(appCtx.timingPerCommand[uint64(id)]) / 1000000.
			timeLbl.SetText(fmt.Sprintf("%6.3f ms ", milliseconds))
			if milliseconds >= 1. {
				timeLbl.SetColor(gxui.ColorFromHex(0xFFFC19 + 0xFF<<24))
			}
			if milliseconds >= 5. {
				timeLbl.SetColor(gxui.ColorFromHex(0xD21212 + 0xFF<<24))
			}
			ll.AddChild(timeLbl)
			//})
		}

		nameLbl := CreateLabel(t, a.Info.Name, COMMAND_COLOR, active)
		ll.AddChild(nameLbl)

		ll.AddChild(CreateLabel(t, "(", CODE_COLOR, active))
		for i, p := range a.Info.Parameters {
			argIdx := i
			v := a.Arguments[argIdx]
			if argIdx > 0 {
				ll.AddChild(CreateLabel(t, ", ", CODE_COLOR, active))
			}
			var c gxui.Control
			switch p.Type.GetKind() {
			case service.TypeKindPointer:
				b := t.CreateButton()
				b.SetMargin(math.Spacing{})
				//b.SetPadding(math.Spacing{})
				b.AddChild(CreateLabel(t, fmt.Sprintf("0x%x", v), CONSTANT_COLOR, active))
				b.OnClick(func(gxui.MouseEvent) {
					appCtx.SelectAddress(v.(memory.Pointer))
				})
				c = b
			case service.TypeKindEnum:
				enum := p.Type.(*service.EnumInfo)
				entries := schema.AllEnumEntries(enum)
				value := v.(schema.EnumValue)
				selected := service.EnumEntry{Name: value.String(), Value: value.Value}
				l := createEnumList(t, appCtx, entries, selected, active, func(item gxui.AdapterItem) {
					entry := item.(service.EnumEntry)
					a.Arguments[argIdx] = schema.EnumValue{Type: value.Type, Value: entry.Value}
					appCtx.ReplaceAtom(a, id)
				})
				c = l
			case service.TypeKindBool:
				c = createEnumList(t, appCtx, []bool{false, true}, v, active, func(v gxui.AdapterItem) {
					a.Arguments[argIdx] = v
					appCtx.ReplaceAtom(a, id)
				})
			case service.TypeKindS32:
				c = createTextbox(t, appCtx, v, active, func(s string) bool {
					if i, err := strconv.ParseInt(s, 0, 32); err == nil {
						a.Arguments[argIdx] = int32(i)
						appCtx.ReplaceAtom(a, id)
						return true
					} else {
						return false
					}
				})
			case service.TypeKindS64:
				c = createTextbox(t, appCtx, v, active, func(s string) bool {
					if i, err := strconv.ParseInt(s, 0, 64); err == nil {
						a.Arguments[argIdx] = int64(i)
						appCtx.ReplaceAtom(a, id)
						return true
					} else {
						return false
					}
				})
			case service.TypeKindU32:
				c = createTextbox(t, appCtx, v, active, func(s string) bool {
					if i, err := strconv.ParseUint(s, 0, 32); err == nil {
						a.Arguments[argIdx] = uint32(i)
						appCtx.ReplaceAtom(a, id)
						return true
					} else {
						return false
					}
				})
			case service.TypeKindU64:
				c = createTextbox(t, appCtx, v, active, func(s string) bool {
					if i, err := strconv.ParseUint(s, 0, 64); err == nil {
						a.Arguments[argIdx] = uint64(i)
						appCtx.ReplaceAtom(a, id)
						return true
					} else {
						return false
					}
				})
			case service.TypeKindF32:
				c = createTextbox(t, appCtx, v, active, func(s string) bool {
					if f, err := strconv.ParseFloat(s, 32); err == nil {
						a.Arguments[argIdx] = float32(f)
						appCtx.ReplaceAtom(a, id)
						return true
					} else {
						return false
					}
				})
			default:
				c = CreateLabel(t, fmt.Sprintf("%v", v), CONSTANT_COLOR, active)
			}

			ll.AddChild(c)

			pName := p.Name
			appCtx.ToolTipController().AddToolTip(c, 0.7, func(math.Point) gxui.Control {
				l := t.CreateLabel()
				l.SetText(pName)
				return l
			})
		}

		ll.AddChild(CreateLabel(t, ")", CODE_COLOR, active))

		/* TODO: Return value
		res := ty.ReturnValue()
		if res != nil {
			ll.AddChild(CreateLabel(t, fmt.Sprintf(" → %v", res.Value()), CONSTANT_COLOR, active))
		}
		*/
	}
	return ll
}

func createAtomGroupControls(t gxui.Theme, appCtx *ApplicationContext, g atom.Group) gxui.Control {
	layout := t.CreateLinearLayout()
	layout.SetDirection(gxui.LeftToRight)

	img := t.CreateImage()
	img.SetExplicitSize(math.Size{W: kCommandAdapterItemHeight, H: kCommandAdapterItemHeight})
	img.SetAspectMode(gxui.AspectCorrectLetterbox)
	layout.AddChild(img)

	atomID := g.Range.Last()
	atom := appCtx.Atoms()[atomID]
	if atom.Info.IsDrawCall || atom.Info.IsEndOfFrame {
		var cancel chan<- struct{}
		cancelThumbnail := func() {
			if cancel != nil {
				close(cancel)
				cancel = nil
			}
		}
		requestThumbnail := func() {
			cancelThumbnail()
			cancel = appCtx.RequestThumbnail(atomID, kFilmStripAdapterItemWidth, kFilmStripAdapterItemHeight, func(tex gxui.Texture) {
				img.SetTexture(tex)
				appCtx.ToolTipController().AddToolTip(img, 0.7, func(math.Point) gxui.Control {
					large := t.CreateImage()
					large.SetTexture(tex)
					large.SetAspectMode(gxui.AspectCorrectLetterbox)
					return large
				})
			})
		}

		var subscription gxui.EventSubscription
		img.OnAttach(func() {
			requestThumbnail()
			subscription = appCtx.OnDeviceSelected(requestThumbnail)
		})
		img.OnDetach(func() {
			cancelThumbnail()
			subscription.Unlisten()
		})
	}

	label := t.CreateLabel()
	label.SetText(g.Name)
	layout.AddChild(label)

	return layout
}

type cmdAdapterItem struct {
	depth  uint
	atomID atom.ID
}

type CommandAdapterNode struct {
	appCtx *ApplicationContext
	group  atom.Group
	depth  uint
}

func (n CommandAdapterNode) Count() int {
	return int(n.group.Count())
}

func (n CommandAdapterNode) NodeAt(index int) gxui.TreeNode {
	_, subgroup := n.group.Index(uint64(index))
	if subgroup != nil {
		return &CommandAdapterNode{
			appCtx: n.appCtx,
			group:  *subgroup,
			depth:  n.depth + 1,
		}
	} else {
		return nil
	}
}

func (n CommandAdapterNode) ItemAt(index int) gxui.AdapterItem {
	atomID, _ := n.group.Index(uint64(index))
	return cmdAdapterItem{n.depth, atomID}
}

func (n CommandAdapterNode) ItemIndex(item gxui.AdapterItem) int {
	return int(n.group.IndexOf(item.(cmdAdapterItem).atomID))
}

func (n CommandAdapterNode) Create(theme gxui.Theme, index int) gxui.Control {
	id, subgroup := n.group.Index(uint64(index))
	if subgroup != nil {
		return createAtomGroupControls(theme, n.appCtx, *subgroup)
	} else {
		return createAtomControls(theme, n.appCtx, id)
	}
}

type CommandAdapter struct {
	gxui.AdapterBase
	CommandAdapterNode
	appCtx *ApplicationContext
}

func CreateCommandAdapter(appCtx *ApplicationContext) *CommandAdapter {
	a := &CommandAdapter{
		CommandAdapterNode: CommandAdapterNode{appCtx: appCtx},
	}
	return a
}

func (a *CommandAdapter) SetRoot(root atom.Group) {
	a.group = root
	a.DataReplaced()
}

func (a CommandAdapter) AtomRange(item gxui.AdapterItem) atom.Range {
	if item == nil {
		return atom.Range{}
	}
	cai := item.(cmdAdapterItem)
	depth, atomID := cai.depth, cai.atomID
	g := a.group
	for {
		idx := g.SubGroups.IndexOf(atomID)
		if idx < 0 {
			return atom.Range{Start: atomID, End: atomID + 1}
		} else {
			g = g.SubGroups[idx]
		}
		if depth == 0 {
			return g.Range
		}
		depth--
	}
}

func (a CommandAdapter) Item(id atom.ID) gxui.AdapterItem {
	depth := uint(0)
	g := a.group
	for {
		idx := g.SubGroups.IndexOf(id)
		if idx < 0 {
			return cmdAdapterItem{depth, id}
		}
		g = g.SubGroups[idx]
		depth++
	}
}

// gxui.TreeAdapter compliance
func (a CommandAdapter) Size(theme gxui.Theme) math.Size {
	return math.Size{W: math.MaxSize.W, H: kCommandAdapterItemHeight}
}
