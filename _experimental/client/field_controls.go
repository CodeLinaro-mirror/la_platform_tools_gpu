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
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/schema"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/service/path"
	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

type parse func(s string) (interface{}, bool)
type commit func(interface{})

func findConstants(t binary.Type, appCtx *ApplicationContext) schema.ConstantSet {
	if t == nil {
		return schema.ConstantSet{}
	}
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

func createLabel(appCtx *ApplicationContext, s string, c gxui.Color) gxui.Label {
	l := appCtx.theme.CreateLabel()
	l.SetText(s)
	l.SetColor(c)
	l.SetMargin(math.Spacing{})
	return l
}

func createMonospaceLabel(appCtx *ApplicationContext, s string, c gxui.Color) gxui.Label {
	l := createLabel(appCtx, s, c)
	l.SetFont(appCtx.monospace)
	return l
}

func createEnumList(appCtx *ApplicationContext, values interface{}, selected interface{}, commit commit) gxui.DropDownList {
	a := gxui.CreateDefaultAdapter()
	a.SetItems(values)
	a.SetSizeAsLargest(appCtx.theme)
	a.SetStyleLabel(func(t gxui.Theme, l gxui.Label) { l.SetColor(CONSTANT_COLOR) })
	l := appCtx.theme.CreateDropDownList()
	l.SetAdapter(a)
	l.SetBubbleOverlay(appCtx.dropDownOverlay)
	l.Select(selected)
	l.OnSelectionChanged(func(i gxui.AdapterItem) { commit(i) })
	return l
}

func createTextbox(appCtx *ApplicationContext, value interface{}, parse parse, commit commit) gxui.TextBox {
	newvalue := value
	tb := appCtx.theme.CreateTextBox()
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
	tb.SetTextColor(CONSTANT_COLOR)
	return tb
}

func createIntField(appCtx *ApplicationContext, p path.Path, v interface{}, s schema.ConstantSet) gxui.Control {
	if len(s.Entries) > 0 {
		c := findConstant(s, v)
		return createEnumList(appCtx, s.Entries, c, func(v interface{}) {
			appCtx.Change(p, v.(schema.Constant).Value)
		})
	}
	ty := reflect.TypeOf(v)
	return createTextbox(appCtx, v, func(s string) (interface{}, bool) {
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

func createUintField(appCtx *ApplicationContext, p path.Path, v interface{}, s schema.ConstantSet) gxui.Control {
	if len(s.Entries) > 0 {
		c := findConstant(s, v)
		return createEnumList(appCtx, s.Entries, c, func(v interface{}) {
			appCtx.Change(p, v.(schema.Constant).Value)
		})
	}
	ty := reflect.TypeOf(v)
	return createTextbox(appCtx, v, func(s string) (interface{}, bool) {
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

func createFloatField(appCtx *ApplicationContext, p path.Path, v interface{}) gxui.Control {
	ty := reflect.TypeOf(v)
	return createTextbox(appCtx, v, func(s string) (interface{}, bool) {
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

func createField(appCtx *ApplicationContext, p path.Path, t binary.Type, v interface{}) gxui.Control {
	var c gxui.Control

	switch v := schema.Underlying(v).(type) {
	case *memory.Pointer:
		b := appCtx.theme.CreateButton()
		b.SetMargin(math.Spacing{})
		b.AddChild(createLabel(appCtx, v.String(), CONSTANT_COLOR))
		c = b

	case *atom.Observations:
		return nil // don't display observation fields

	case bool:
		c = createEnumList(appCtx, []bool{false, true}, v, func(v interface{}) {
			appCtx.Change(p, v)
		})

	case int8, int16, int32, int64:
		c = createIntField(appCtx, p, v, findConstants(t, appCtx))

	case uint8, uint16, uint32, uint64:
		c = createUintField(appCtx, p, v, findConstants(t, appCtx))

	case float32, float64:
		c = createFloatField(appCtx, p, v)

	default:
		c = createLabel(appCtx, fmt.Sprintf("%v", v), CONSTANT_COLOR)
	}

	c.OnClick(func(gxui.MouseEvent) {
		l, err := appCtx.rpc.Follow(p)
		if err == nil && l != nil {
			appCtx.events.Select(l)
		}
	})

	appCtx.toolTipController.AddToolTip(c, 0.7, func(math.Point) gxui.Control {
		l := appCtx.theme.CreateLabel()
		l.SetText(p.Path())
		return l
	})

	return c
}
