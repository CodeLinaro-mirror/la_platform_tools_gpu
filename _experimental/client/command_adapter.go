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
	"reflect"
	"strconv"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary/schema"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/service/path"
	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

const kCommandAdapterItemHeight = 18

type parser func(s string) (interface{}, bool)
type committer func(interface{})

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

func createTextbox(t gxui.Theme, value interface{}, active bool, parse parser, commit committer) gxui.TextBox {
	newvalue := value
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
		if v, ok := parse(text); ok {
			tb.SetTextColor(CONSTANT_COLOR)
			newvalue = v
		} else {
			tb.SetTextColor(gxui.Red)
		}
	})
	tb.OnLostFocus(func() {
		if value != newvalue {
			commit(newvalue)
		}
	})
	if active {
		tb.SetTextColor(CONSTANT_COLOR)
	} else {
		tb.SetTextColor(INACTIVE_COLOR)
	}
	return tb
}

func findConstants(t schema.Type, appCtx *ApplicationContext) schema.ConstantSet {
	set, _ := appCtx.constants[t.String()]
	return set
}

func findConstant(s schema.ConstantSet, v interface{}) schema.Constant {
	for _, entry := range s.Entries {
		if entry.Value == v {
			return entry
		}
	}
	return schema.Constant{}
}

func createIntField(t gxui.Theme, appCtx *ApplicationContext, p path.Path, v interface{}, s schema.ConstantSet) gxui.Control {
	if len(s.Entries) > 0 {
		c := findConstant(s, v)
		return createEnumList(t, appCtx, s.Entries, c, true, func(v gxui.AdapterItem) {
			appCtx.Change(p, v.(schema.Constant).Value)
		})
	}
	ty := reflect.TypeOf(v)
	return createTextbox(t, v, true, func(s string) (interface{}, bool) {
		if i, err := strconv.ParseInt(s, 0, ty.Bits()); err == nil {
			v := reflect.New(ty).Elem()
			v.SetInt(i)
			return v.Interface(), true
		} else {
			return nil, false
		}
	}, func(v interface{}) {
		appCtx.Change(p, v)
	})
}

func createUintField(t gxui.Theme, appCtx *ApplicationContext, p path.Path, v interface{}, s schema.ConstantSet) gxui.Control {
	if len(s.Entries) > 0 {
		c := findConstant(s, v)
		return createEnumList(t, appCtx, s.Entries, c, true, func(v gxui.AdapterItem) {
			appCtx.Change(p, v.(schema.Constant).Value)
		})
	}
	ty := reflect.TypeOf(v)
	return createTextbox(t, v, true, func(s string) (interface{}, bool) {
		if i, err := strconv.ParseUint(s, 0, ty.Bits()); err == nil {
			v := reflect.New(ty).Elem()
			v.SetUint(i)
			return v.Interface(), true
		} else {
			return nil, false
		}
	}, func(v interface{}) {
		appCtx.Change(p, v)
	})
}

func createFloatField(t gxui.Theme, appCtx *ApplicationContext, p path.Path, v interface{}) gxui.Control {
	ty := reflect.TypeOf(v)
	return createTextbox(t, v, true, func(s string) (interface{}, bool) {
		if i, err := strconv.ParseFloat(s, ty.Bits()); err == nil {
			v := reflect.New(ty).Elem()
			v.SetFloat(i)
			return v.Interface(), true
		} else {
			return nil, false
		}
	}, func(v interface{}) {
		appCtx.Change(p, v)
	})
}

func createAtomControls(t gxui.Theme, appCtx *ApplicationContext, id atom.ID) gxui.Control {
	atomPath := appCtx.CaptureID().Path().Atoms().Index(uint64(id))
	a := appCtx.Atoms()[id].(*Atom)
	active := true

	ll := t.CreateLinearLayout()
	ll.SetDirection(gxui.LeftToRight)
	ll.AddChild(CreateLabel(t, fmt.Sprintf("%.6d ", id), LINE_NUMBER_COLOR, active))

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

	nameLbl := CreateLabel(t, atom.MetadataOf(a).DisplayName, COMMAND_COLOR, active)
	ll.AddChild(nameLbl)

	ll.AddChild(CreateLabel(t, "(", CODE_COLOR, active))
	needcomma := false
	for i := 0; i < a.FieldCount(); i++ {
		argIdx := i // capture for closures
		info, v := a.Field(argIdx)
		constants := findConstants(info.Type, appCtx)
		p := atomPath.Field(info.Name())
		if needcomma {
			ll.AddChild(CreateLabel(t, ", ", CODE_COLOR, active))
		}
		var c gxui.Control

		switch v := schema.Underlying(v).(type) {
		case *memory.Pointer:
			b := t.CreateButton()
			b.SetMargin(math.Spacing{})
			//b.SetPadding(math.Spacing{})
			b.AddChild(CreateLabel(t, v.String(), CONSTANT_COLOR, active))
			b.OnClick(func(gxui.MouseEvent) { appCtx.SelectPointer(*v) })
			c = b

		case *atom.Observations:
			continue //don't display observations as a parameter

		case bool:
			c = createEnumList(t, appCtx, []bool{false, true}, v, active, func(v gxui.AdapterItem) {
				appCtx.Change(p, v)
			})

		case int8, int16, int32, int64:
			c = createIntField(t, appCtx, p, v, constants)

		case uint8, uint16, uint32, uint64:
			c = createUintField(t, appCtx, p, v, constants)

		case float32, float64:
			c = createFloatField(t, appCtx, p, v)

		default:
			c = CreateLabel(t, fmt.Sprintf("%v", v), CONSTANT_COLOR, active)
		}

		ll.AddChild(c)
		needcomma = true
		appCtx.ToolTipController().AddToolTip(c, 0.7, func(math.Point) gxui.Control {
			l := t.CreateLabel()
			l.SetText(p.Path())
			return l
		})
	}

	ll.AddChild(CreateLabel(t, ")", CODE_COLOR, active))
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
	if atom.Flags().IsDrawCall() || atom.Flags().IsEndOfFrame() {
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

type cmdNode interface {
	atomRange() atom.Range
}

type observationsItem struct {
	atomID atom.ID
	index  int
}

func (i observationsItem) atomRange() atom.Range {
	return atom.Range{Start: i.atomID, End: i.atomID + 1}
}

type hierarchyItem atom.Range

func (i hierarchyItem) atomRange() atom.Range {
	return atom.Range(i)
}

type observationsNode struct {
	appCtx *ApplicationContext
	atomID atom.ID
}

func (n observationsNode) Count() int {
	atom := n.appCtx.Atoms()[n.atomID]
	observations := atom.Observations()
	return len(observations.Reads) + len(observations.Writes)
}

func (n observationsNode) NodeAt(index int) gxui.TreeNode {
	return nil
}

func (n observationsNode) ItemAt(index int) gxui.AdapterItem {
	return observationsItem{n.atomID, index}
}

func (n observationsNode) ItemIndex(item gxui.AdapterItem) int {
	return item.(observationsItem).index
}

func (n observationsNode) Create(theme gxui.Theme, index int) gxui.Control {
	atom := n.appCtx.Atoms()[n.atomID]
	observations := atom.Observations()
	var r memory.Range
	var c gxui.Color
	if index < len(observations.Reads) {
		r = observations.Reads[index].Range
		c = gxui.Green
	} else {
		index -= len(observations.Reads)
		r = observations.Writes[index].Range
		c = gxui.Red
	}

	ptr := memory.Pointer{Address: r.Base, Pool: memory.ApplicationPool}
	b := theme.CreateButton()
	b.SetMargin(math.Spacing{})
	b.AddChild(CreateLabel(theme, r.String(), c, true))
	b.OnClick(func(gxui.MouseEvent) { n.appCtx.SelectPointer(ptr) })
	return b
}

type hierarchyNode struct {
	appCtx *ApplicationContext
	group  atom.Group
	depth  uint
}

func (n hierarchyNode) Count() int {
	return int(n.group.Count())
}

func (n hierarchyNode) NodeAt(index int) gxui.TreeNode {
	if id, subgroup := n.group.Index(uint64(index)); subgroup != nil {
		return &hierarchyNode{
			appCtx: n.appCtx,
			group:  *subgroup,
			depth:  n.depth + 1,
		}
	} else {
		return observationsNode{n.appCtx, id}
	}
}

func (n hierarchyNode) ItemAt(index int) gxui.AdapterItem {
	if id, group := n.group.Index(uint64(index)); group == nil {
		return hierarchyItem{Start: id, End: id + 1}
	} else {
		return hierarchyItem(group.Range)
	}
}

func (n hierarchyNode) ItemIndex(item gxui.AdapterItem) int {
	return int(n.group.IndexOf(item.(cmdNode).atomRange().Start))
}

func (n hierarchyNode) Create(theme gxui.Theme, index int) gxui.Control {
	id, subgroup := n.group.Index(uint64(index))
	if subgroup != nil {
		return createAtomGroupControls(theme, n.appCtx, *subgroup)
	} else {
		return createAtomControls(theme, n.appCtx, id)
	}
}

type CommandAdapter struct {
	gxui.AdapterBase
	hierarchyNode
	appCtx *ApplicationContext
}

func CreateCommandAdapter(appCtx *ApplicationContext) *CommandAdapter {
	a := &CommandAdapter{
		hierarchyNode: hierarchyNode{appCtx: appCtx},
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
	return item.(cmdNode).atomRange()
}

func (a CommandAdapter) Item(id atom.ID) gxui.AdapterItem {
	return hierarchyItem{Start: id, End: id + 1}
}

// gxui.TreeAdapter compliance
func (a CommandAdapter) Size(theme gxui.Theme) math.Size {
	return math.Size{W: math.MaxSize.W, H: kCommandAdapterItemHeight}
}
