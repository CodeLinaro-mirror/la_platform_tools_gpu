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
	"github.com/google/gxui"
)

func CreateMemoryPanelReinterpretButtons(
	appCtx *ApplicationContext,
	memory_list gxui.List,
	rawAdapter *MemoryAdapter,
	imgAdapter *MemoryImageAdapter) gxui.Control {

	theme := appCtx.Theme()
	modes := []string{}
	actions := map[string]func(){}

	for _, t := range []DataType{
		U8(0), S8(0),
		U16(0), S16(0),
		U32(0), S32(0),
		F32(0), F64(0),
	} {
		t := t
		modes = append(modes, t.Name())
		actions[t.Name()] = func() {
			rawAdapter.SetDataType(t)
			memory_list.SetAdapter(rawAdapter)
		}
	}

	for _, t := range []PixelType{
		RGB888(0),
		RGBA8888(0),
	} {
		t := t
		modes = append(modes, t.Name())
		actions[t.Name()] = func() {
			imgAdapter.SetPixelType(t)
			memory_list.SetAdapter(imgAdapter)
		}
	}

	adapter := gxui.CreateDefaultAdapter()
	adapter.SetItems(modes)

	label := theme.CreateLabel()
	label.SetText("Type:")

	list := theme.CreateDropDownList()
	list.SetBubbleOverlay(appCtx.DropDownOverlay())
	list.SetAdapter(adapter)
	list.OnSelectionChanged(func(item gxui.AdapterItem) { actions[item.(string)]() })
	list.Select(modes[0])

	layout := theme.CreateLinearLayout()
	layout.SetDirection(gxui.LeftToRight)
	layout.AddChild(label)
	layout.AddChild(list)
	return layout
}

func CreateMemoryPanel(appCtx *ApplicationContext) gxui.Control {
	theme := appCtx.Theme()

	rawAdapter := CreateMemoryAdapter(appCtx)
	imgAdapter := CreateMemoryImageAdapter(appCtx)

	list := theme.CreateList()
	list.SetAdapter(rawAdapter)
	appCtx.OnAtomsUpdated(func() {
		rawAdapter.DataReplaced()
		imgAdapter.DataReplaced()
	})
	appCtx.OnAtomSelected(func() {
		selectedAtomID, selectedAddress := appCtx.SelectedAtomID(), appCtx.SelectedAddress()
		rawAdapter.SetData(selectedAtomID, selectedAddress)
		imgAdapter.SetData(selectedAtomID, selectedAddress)
	})
	appCtx.OnAddressSelected(func() {
		selectedAtomID, selectedAddress := appCtx.SelectedAtomID(), appCtx.SelectedAddress()
		rawAdapter.SetData(selectedAtomID, selectedAddress)
		imgAdapter.SetData(selectedAtomID, selectedAddress)
		list.ScrollTo(selectedAddress)
	})
	layout := theme.CreateLinearLayout()
	layout.SetDirection(gxui.TopToBottom)
	layout.AddChild(CreateMemoryPanelReinterpretButtons(appCtx, list, rawAdapter, imgAdapter))
	layout.AddChild(list)
	return layout
}
