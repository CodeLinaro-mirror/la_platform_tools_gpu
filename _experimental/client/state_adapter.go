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
	"android.googlesource.com/platform/tools/gpu/binary/schema"
	"android.googlesource.com/platform/tools/gpu/memory"
	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

const kStateAdapterItemHeight = 18

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
	addButton := func(format string, args ...interface{}) gxui.Button {
		button := theme.CreateButton()
		button.SetText(fmt.Sprintf(format, args...))
		// button.SetPadding(math.ZeroSpacing)
		button.SetMargin(math.ZeroSpacing)
		layout.AddChild(button)
		return button
	}

	addLabel("%s: ", name)

	switch ty := value.(type) {
	case atom.ID:
		addButton(fmt.Sprintf("0x%.8x", ty)).OnClick(func(gxui.MouseEvent) {
			appCtx.SelectAtom(ty)
		})
	case memory.Pointer:
		addButton(fmt.Sprintf("0x%.8x", ty)).OnClick(func(gxui.MouseEvent) {
			appCtx.SelectAddress(ty)
		})
	/*
		case schema.EnumValue:
			addLabel(fmt.Sprintf("%v", ty))
	*/
	case string, int, uint, int32, uint32, int16, uint16, int8, uint8, float32, float64, bool:
		// TODO: Click to select object
		// obj := state.LookupId(ty)
		// addButton(ty.String()).OnClick(func(gxui.MouseEvent) {
		// 	appCtx.SelectObject(obj)
		// })
		addLabel(fmt.Sprintf("%v", ty))
	}
	return layout
}

type StateAdapterItem struct {
	appCtx   *ApplicationContext
	key      string
	value    interface{}
	children []*StateAdapterItem
}

func (i *StateAdapterItem) Init(appCtx *ApplicationContext, key string, value interface{}) {
	i.appCtx = appCtx
	i.key = key
	i.value = value
	/*
		switch ty := value.(type) {
		case schema.Struct:
			for _, f := range ty.Fields {
				child := &StateAdapterItem{}
				child.Init(appCtx, f.Info.Name, f.Value)
				i.children = append(i.children, child)
			}
		case schema.Class:
			for _, f := range ty.Fields {
				child := &StateAdapterItem{}
				child.Init(appCtx, f.Info.Name, f.Value)
				i.children = append(i.children, child)
			}
		case schema.Array:
			for k, e := range ty.Elements {
				child := &StateAdapterItem{}
				child.Init(appCtx, fmt.Sprintf("%d", k), e)
				i.children = append(i.children, child)
			}
		case schema.Map:
			for _, e := range ty.Elements {
				child := &StateAdapterItem{}
				child.Init(appCtx, fmt.Sprintf("%v", e.Key), e.Value)
				i.children = append(i.children, child)
			}
		}
	*/
}

func (i *StateAdapterItem) Count() int {
	return len(i.children)
}

func (i *StateAdapterItem) NodeAt(index int) gxui.TreeNode {
	if len(i.children[index].children) > 0 {
		return i.children[index]
	} else {
		return nil
	}
}

func (i *StateAdapterItem) ItemAt(index int) gxui.AdapterItem {
	return i.children[index]
}

func (i *StateAdapterItem) ItemIndex(item gxui.AdapterItem) int {
	// Brute force search
	sai := item.(*StateAdapterItem)
	for i, c := range i.children {
		if c == sai {
			return i
		}
		if idx := c.ItemIndex(item); idx >= 0 {
			return i
		}
	}
	return -1
}

func (i *StateAdapterItem) Create(t gxui.Theme, index int) gxui.Control {
	return createControls(i.appCtx, i.children[index].key, i.children[index].value)
}

type StateAdapter struct {
	StateAdapterItem
	gxui.AdapterBase
}

func CreateStateAdapter(appCtx *ApplicationContext, state schema.Struct) *StateAdapter {
	adapter := &StateAdapter{}
	adapter.Init(appCtx, "state", state)
	return adapter
}

func (r *StateAdapter) Size(theme gxui.Theme) math.Size {
	return math.Size{W: math.MaxSize.W, H: kStateAdapterItemHeight}
}
