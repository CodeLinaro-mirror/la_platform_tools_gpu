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
	"sort"

	"android.googlesource.com/platform/tools/gpu/binary/schema"
	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

const kStateAdapterNodeHeight = 18

func createControls(appCtx *ApplicationContext, name string, value interface{}) gxui.Control {
	theme := appCtx.Theme()

	layout := theme.CreateLinearLayout()
	layout.SetDirection(gxui.LeftToRight)

	addLabel := func(format string, args ...interface{}) {
		label := theme.CreateLabel()
		label.SetText(fmt.Sprintf(format, args...))
		label.SetMargin(math.ZeroSpacing)
		label.SetMultiline(false)
		layout.AddChild(label)
	}
	//	addButton := func(format string, args ...interface{}) gxui.Button {
	//		button := theme.CreateButton()
	//		button.SetText(fmt.Sprintf(format, args...))
	//		// button.SetPadding(math.ZeroSpacing)
	//		button.SetMargin(math.ZeroSpacing)
	//		layout.AddChild(button)
	//		return button
	//	}

	addLabel("%s: ", name)

	if value != nil {
		addLabel(fmt.Sprintf("%v", value))
	}
	//	switch ty := value.(type) {
	//	case atom.ID:
	//		addButton(fmt.Sprintf("0x%.8x", ty)).OnClick(func(gxui.MouseEvent) {
	//			appCtx.SelectAtom(ty)
	//		})
	//	case memory.Pointer:
	//		addButton(fmt.Sprintf("0x%.8x", ty)).OnClick(func(gxui.MouseEvent) {
	//			appCtx.SelectAddress(ty)
	//		})
	//	/*
	//		case schema.EnumValue:
	//			addLabel(fmt.Sprintf("%v", ty))
	//	*/
	//	case string, int, uint, int32, uint32, int16, uint16, int8, uint8, float32, float64, bool:
	//		// TODO: Click to select object
	//		// obj := state.LookupId(ty)
	//		// addButton(ty.String()).OnClick(func(gxui.MouseEvent) {
	//		// 	appCtx.SelectObject(obj)
	//		// })
	//		addLabel(fmt.Sprintf("%v", ty))
	//	}
	return layout
}

type StateAdapterNode struct {
	appCtx   *ApplicationContext
	name     string
	value    interface{}
	item     string
	children StateAdapterNodeList
	parent   *StateAdapterNode
}

type StateAdapterNodeList []*StateAdapterNode

func (l StateAdapterNodeList) Len() int           { return len(l) }
func (l StateAdapterNodeList) Less(a, b int) bool { return l[a].name < l[b].name }
func (l StateAdapterNodeList) Swap(a, b int)      { l[a], l[b] = l[b], l[a] }

func (n *StateAdapterNode) add(name string, value interface{}, item string) {
	n.children = append(n.children, &StateAdapterNode{
		appCtx: n.appCtx,
		name:   name,
		value:  value,
		item:   item,
		parent: n,
	})
}

func (n *StateAdapterNode) init() {
	if v, ok := n.value.(*schema.Object); ok {
		for i := range v.Fields {
			name := v.Type.Fields[i].Declared
			n.add(name, v.Fields[i], fmt.Sprintf("%s.%s", n.item, name))
		}
	} else {
		v := reflect.ValueOf(n.value)
		switch v.Kind() {
		case reflect.Array, reflect.Slice:
			for i, c := 0, v.Len(); i < c; i++ {
				name := fmt.Sprintf("%d", i)
				v := v.Index(i)
				n.add(name, v.Interface(), fmt.Sprintf("%s[%s]", n.item, name))
			}
		case reflect.Map:
			for _, k := range v.MapKeys() {
				name := fmt.Sprintf("%v", k.Interface())
				v := v.MapIndex(k)
				n.add(name, v.Interface(), fmt.Sprintf("%s[%s]", n.item, name))
			}
			sort.Sort(n.children)
		}
	}

	for _, c := range n.children {
		c.init()
	}
}

func (n *StateAdapterNode) Count() int {
	return len(n.children)
}

func (n *StateAdapterNode) NodeAt(index int) gxui.TreeNode {
	return n.children[index]
}

func (n *StateAdapterNode) ItemAt(index int) gxui.AdapterItem {
	return n.children[index].item
}

func (n *StateAdapterNode) ItemIndex(item gxui.AdapterItem) int {
	// Brute force search
	for i, c := range n.children {
		if c.item == item {
			return i
		}
		if c.ItemIndex(item) >= 0 {
			return i
		}
	}
	return -1
}

func (n *StateAdapterNode) Create(t gxui.Theme, index int) gxui.Control {
	c := n.children[index]
	if len(c.children) > 0 {
		return createControls(n.appCtx, c.name, nil)
	} else {
		return createControls(n.appCtx, c.name, c.value)
	}
}

type StateAdapter struct {
	StateAdapterNode
	gxui.AdapterBase
}

func (r *StateAdapter) Size(theme gxui.Theme) math.Size {
	return math.Size{W: math.MaxSize.W, H: kStateAdapterNodeHeight}
}

func NewStateAdapter(appCtx *ApplicationContext) *StateAdapter {
	a := &StateAdapter{
		StateAdapterNode: StateAdapterNode{
			appCtx: appCtx,
			name:   "state",
			item:   "state",
			value:  appCtx.state,
		},
	}
	a.init()
	return a
}
