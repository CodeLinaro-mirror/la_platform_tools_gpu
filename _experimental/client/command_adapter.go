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
	"android.googlesource.com/platform/tools/gpu/service"
	"android.googlesource.com/platform/tools/gpu/service/path"
	"android.googlesource.com/platform/tools/gpu/task"
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
	l.SetBubbleOverlay(appCtx.dropDownOverlay)
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

func createAtomControls(ctx *commandAdapterCtx, id atom.ID) gxui.Control {
	t := ctx.appCtx.theme
	p := ctx.capture.Atoms().Index(uint64(id))
	a := ctx.atoms[p.Index].(*Atom)
	active := true

	layout := t.CreateLinearLayout()
	layout.SetDirection(gxui.LeftToRight)
	layout.AddChild(CreateLabel(t, fmt.Sprintf("%.6d ", p.Index), LINE_NUMBER_COLOR, active))

	if ns, ok := ctx.timings.AtomDuration(id); ok {
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

	nameLbl := CreateLabel(t, atom.MetadataOf(a).DisplayName, COMMAND_COLOR, active)
	layout.AddChild(nameLbl)

	layout.AddChild(CreateLabel(t, "(", CODE_COLOR, active))
	needcomma := false
	for i := 0; i < a.FieldCount(); i++ {
		argIdx := i // capture for closures
		info, v := a.Field(argIdx)
		constants := findConstants(info.Type, ctx.appCtx)
		p := p.Field(info.Name())
		if needcomma {
			layout.AddChild(CreateLabel(t, ", ", CODE_COLOR, active))
		}
		var c gxui.Control

		switch v := schema.Underlying(v).(type) {
		case *memory.Pointer:
			b := t.CreateButton()
			b.SetMargin(math.Spacing{})
			//b.SetPadding(math.Spacing{})
			b.AddChild(CreateLabel(t, v.String(), CONSTANT_COLOR, active))
			//b.OnClick(func(gxui.MouseEvent) { appCtx.SelectPointer(*v) }) // [BENC]: TODO
			c = b

		case *atom.Observations:
			continue //don't display observations as a parameter

		case bool:
			c = createEnumList(t, ctx.appCtx, []bool{false, true}, v, active, func(v gxui.AdapterItem) {
				ctx.appCtx.Change(p, v)
			})

		case int8, int16, int32, int64:
			c = createIntField(t, ctx.appCtx, p, v, constants)

		case uint8, uint16, uint32, uint64:
			c = createUintField(t, ctx.appCtx, p, v, constants)

		case float32, float64:
			c = createFloatField(t, ctx.appCtx, p, v)

		default:
			c = CreateLabel(t, fmt.Sprintf("%v", v), CONSTANT_COLOR, active)
		}

		layout.AddChild(c)
		needcomma = true

		ctx.appCtx.toolTipController.AddToolTip(c, 0.7, func(math.Point) gxui.Control {
			l := t.CreateLabel()
			l.SetText(p.Path())
			return l
		})
	}

	layout.AddChild(CreateLabel(t, ")", CODE_COLOR, active))
	return layout
}

func createAtomGroupControls(ctx *commandAdapterCtx, g atom.Group) gxui.Control {
	theme := ctx.appCtx.theme

	layout := theme.CreateLinearLayout()
	layout.SetDirection(gxui.LeftToRight)

	img := theme.CreateImage()
	img.SetExplicitSize(math.Size{W: kCommandAdapterItemHeight, H: kCommandAdapterItemHeight})
	img.SetAspectMode(gxui.AspectCorrectLetterbox)
	layout.AddChild(img)

	if ns, ok := ctx.timings.RangeDuration(g.Range); ok {
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

	atomID := g.Range.Last()
	appCtx := ctx.appCtx

	t := task.New()
	update := func() {
		t.Cancel()
		if ctx.device != nil {
			if flags := ctx.atoms[atomID].Flags(); flags.IsDrawCall() || flags.IsEndOfFrame() {
				after := ctx.capture.Atoms().Index(uint64(atomID))
				t.Run(updateThumbnail{appCtx, ctx.device, after, img})
			}
		}
	}

	// TODO: When trees can be updated, this should listen for capture / device changes.
	img.OnAttach(update)
	img.OnDetach(t.Cancel)

	label := theme.CreateLabel()
	label.SetText(g.Name)
	layout.AddChild(label)

	return layout
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
	ctx    *commandAdapterCtx
	atomID atom.ID
}

func (n observationsNode) Count() int {
	atom := n.ctx.atoms[n.atomID]
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
	atom := n.ctx.atoms[n.atomID]
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

	b := theme.CreateButton()
	b.SetMargin(math.Spacing{})
	b.AddChild(CreateLabel(theme, r.String(), c, true))
	// ptr := memory.Pointer{Address: r.Base, Pool: memory.ApplicationPool}
	// b.OnClick(func(gxui.MouseEvent) { n.appCtx.SelectPointer(ptr) }) // [BENC]: TODO
	return b
}

type hierarchyNode struct {
	ctx   *commandAdapterCtx
	group atom.Group
	depth uint
}

func (n hierarchyNode) Count() int {
	return int(n.group.Count())
}

func (n hierarchyNode) NodeAt(index int) gxui.TreeNode {
	if id, subgroup := n.group.Index(uint64(index)); subgroup != nil {
		return &hierarchyNode{
			ctx:   n.ctx,
			group: *subgroup,
			depth: n.depth + 1,
		}
	} else {
		return observationsNode{n.ctx, id}
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
	var id atom.ID

	switch i := item.(type) {
	case hierarchyItem:
		id = i.Start
	case observationsItem:
		id = i.atomID
	default:
		panic(fmt.Errorf("Unknown item type %T", i))
	}

	return int(n.group.IndexOf(id))
}

func (n hierarchyNode) Create(theme gxui.Theme, index int) gxui.Control {
	id, subgroup := n.group.Index(uint64(index))
	if subgroup != nil {
		return createAtomGroupControls(n.ctx, *subgroup)
	} else {
		return createAtomControls(n.ctx, id)
	}
}

type CommandAdapter struct {
	gxui.AdapterBase
	hierarchyNode
	ctx commandAdapterCtx
}

func CreateCommandAdapter(appCtx *ApplicationContext) *CommandAdapter {
	a := &CommandAdapter{}
	a.ctx = commandAdapterCtx{appCtx: appCtx}
	a.hierarchyNode.ctx = &a.ctx
	return a
}

func (a *CommandAdapter) UpdateAtoms(capture *path.Capture, atoms []atom.Atom, root atom.Group) {
	a.ctx.capture = capture
	a.ctx.atoms = atoms
	a.hierarchyNode.group = root
	a.DataReplaced()
}

func (a *CommandAdapter) UpdateTimings(timings service.TimingInfo) {
	a.ctx.timings = timings
	a.DataReplaced()
}

func (a *CommandAdapter) UpdateDevice(device *path.Device) {
	a.ctx.device = device
	a.DataReplaced()
}

func (a CommandAdapter) Path(item gxui.AdapterItem) path.Path {
	switch i := item.(type) {
	case nil:
		return nil
	case hierarchyItem:
		span := i.atomRange().Span()
		if span.Start+1 != span.End {
			return a.ctx.capture.Atoms().Slice(span.Start, span.End)
		} else {
			return a.ctx.capture.Atoms().Index(span.Start)
		}
	case observationsItem:
		return a.ctx.capture.Atoms().Index(uint64(i.atomID)).Field("Observations").ArrayIndex(uint64(i.index))
	default:
		panic(fmt.Errorf("Unknown item type %T", i))
	}
}

func (a CommandAdapter) Item(p path.Path) gxui.AdapterItem {
	if s, _ := path.FindAtomSlice(p); s != nil {
		return hierarchyItem{Start: atom.ID(s.Start), End: atom.ID(s.End)}
	}
	if i := path.FindArrayIndex(p); i != nil {
		if f, ok := i.Array.(*path.Field); ok && f.Name == "Observations" {
			if a, ok := f.Struct.(*path.Atom); ok { // Observations of an atom
				id := atom.ID(a.Index)
				return observationsItem{atomID: id, index: int(i.Index)}
			}
		}
	}
	if a := path.FindAtom(p); a != nil {
		id := atom.ID(a.Index)
		return hierarchyItem{Start: id, End: id + 1}
	}
	return nil
}

// gxui.TreeAdapter compliance
func (a CommandAdapter) Size(theme gxui.Theme) math.Size {
	return math.Size{W: math.MaxSize.W, H: kCommandAdapterItemHeight}
}

type commandAdapterCtx struct {
	appCtx  *ApplicationContext
	capture *path.Capture
	device  *path.Device
	atoms   []atom.Atom
	timings service.TimingInfo
}
