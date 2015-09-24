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

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary/schema"
	"github.com/google/gxui"
)

type atomItem struct {
	atomIndex uint64
}

// atomTreeNode is a gxui.TreeNode representing a single atom.
// It has the item of type atomItem.
type atomTreeNode struct {
	ctx  *commandAdapterCtx
	item atomItem
}

func (n atomTreeNode) Count() int {
	atom := n.ctx.atoms[n.item.atomIndex]
	observations := atom.Observations()
	return len(observations.Reads) + len(observations.Writes)
}

func (n atomTreeNode) NodeAt(index int) gxui.TreeNode {
	atom := n.ctx.atoms[n.item.atomIndex]
	observations := atom.Observations()
	if index < len(observations.Reads) {
		return observationTreeNode{
			ctx: n.ctx,
			item: observationItem{
				atomIndex: n.item.atomIndex,
				isRead:    true,
				index:     index,
			},
		}
	} else {
		return observationTreeNode{
			ctx: n.ctx,
			item: observationItem{
				atomIndex: n.item.atomIndex,
				isRead:    false,
				index:     index - len(observations.Reads),
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
		atom := n.ctx.atoms[n.item.atomIndex]
		observations := atom.Observations()
		return len(observations.Reads) + i.index
	}
}

func (n atomTreeNode) Create(t gxui.Theme) gxui.Control {
	p := n.ctx.capture.Atoms().Index(n.item.atomIndex)
	a := n.ctx.atoms[p.Index].(atom.Atom)

	layout := t.CreateLinearLayout()
	layout.SetDirection(gxui.LeftToRight)
	layout.AddChild(createLabel(n.ctx.appCtx, fmt.Sprintf("%.6d ", p.Index), LINE_NUMBER_COLOR))

	if ns, ok := n.ctx.timings.AtomDuration(n.item.atomIndex); ok {
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
	for _, param := range atomParameters(a) {
		f, v := param.field, param.value

		if c := createField(n.ctx.appCtx, p.Field(f.Name()), f.Type, v); c != nil {
			if needcomma {
				layout.AddChild(createLabel(n.ctx.appCtx, ", ", CODE_COLOR))
			}

			layout.AddChild(c)
			needcomma = true
		}
	}

	layout.AddChild(createLabel(n.ctx.appCtx, ")", CODE_COLOR))

	if f, v := atomResult(a); f != nil {
		layout.AddChild(createLabel(n.ctx.appCtx, " -> ", COMMAND_COLOR))
		if c := createField(n.ctx.appCtx, p.Field(f.Name()), f.Type, v); c != nil {
			layout.AddChild(c)
			needcomma = true
		}
	}

	return layout
}

func underlyingType(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

func underlyingValue(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}

type atomParameter struct {
	field schema.Field
	value interface{}
}

func atomParameters(a atom.Atom) []atomParameter {
	switch a := a.(type) {
	case *Atom:
		c := a.ParameterCount()
		p := make([]atomParameter, c)
		for i := 0; i < c; i++ {
			p[i].field, p[i].value = a.Parameter(i)
		}
		return p
	default:
		v := underlyingValue(reflect.ValueOf(a))
		t := v.Type()
		c := t.NumField()
		p := make([]atomParameter, 0, c)
		for i := 0; i < c; i++ {
			if f := t.Field(i); !f.Anonymous {
				f := schema.Field{Declared: f.Name}
				p = append(p, atomParameter{field: f, value: v.Field(i).Interface()})
			}
		}
		return p
	}
}

func atomResult(a atom.Atom) (*schema.Field, interface{}) {
	switch a := a.(type) {
	case *Atom:
		return a.Result()
	default:
		for _, p := range atomParameters(a) {
			if p.field.Declared == "Result" {
				return &p.field, p.value
			}
		}
		return nil, nil
	}
}
