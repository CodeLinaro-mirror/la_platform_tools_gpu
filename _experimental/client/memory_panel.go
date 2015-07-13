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
	"android.googlesource.com/platform/tools/gpu/service"
	"android.googlesource.com/platform/tools/gpu/service/path"
	"android.googlesource.com/platform/tools/gpu/task"
	"github.com/google/gxui"
)

func CreateMemoryPanelReinterpretButtons(
	appCtx *ApplicationContext,
	memory_list gxui.List,
	rawAdapter *MemoryAdapter,
	imgAdapter *MemoryImageAdapter) gxui.Control {

	modes := []string{}
	actions := map[string]func(){}

	for _, t := range []DataType{
		U8(0), S8(0),
		U16(0), S16(0),
		U32(0), S32(0),
		U64(0), S64(0),
		F32(0), F64(0),
		ASCII(0),
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

	label := appCtx.theme.CreateLabel()
	label.SetText("Type:")

	list := appCtx.theme.CreateDropDownList()
	list.SetBubbleOverlay(appCtx.dropDownOverlay)
	list.SetAdapter(adapter)
	list.OnSelectionChanged(func(item gxui.AdapterItem) { actions[item.(string)]() })
	list.Select(modes[0])

	layout := appCtx.theme.CreateLinearLayout()
	layout.SetDirection(gxui.LeftToRight)
	layout.AddChild(label)
	layout.AddChild(list)
	return layout
}

func CreateMemoryPanel(appCtx *ApplicationContext) gxui.Control {
	rawAdapter := CreateMemoryAdapter(appCtx)
	imgAdapter := CreateMemoryImageAdapter(appCtx)

	list := appCtx.theme.CreateList()
	list.SetAdapter(rawAdapter)

	var address uint64
	var after *path.Atom

	update := func() {
		rawAdapter.Update(after, address)
		imgAdapter.Update(after, address)
	}

	appCtx.events.OnSelect(func(p path.Path) {
		if a := path.FindAtom(p); a != nil && !path.Equal(a, after) {
			after = a
			update()
		}
		if s, a := path.FindAtomSlice(p); s != nil && a != nil {
			if a := a.Index(s.End - 1); !path.Equal(a, after) {
				after = a
				update()
			}
		}
		// TODO: Address
	})

	layout := appCtx.theme.CreateLinearLayout()
	layout.SetDirection(gxui.TopToBottom)
	layout.AddChild(CreateMemoryPanelReinterpretButtons(appCtx, list, rawAdapter, imgAdapter))
	layout.AddChild(list)
	return layout
}

type requestMemory struct {
	context  *ApplicationContext
	after    *path.Atom
	address  uint64
	size     uint64
	callback func(service.MemoryInfo)
}

func (t requestMemory) Run(c task.CancelSignal) {
	res, err := t.context.rpc.RequestMemory(t.after, t.address, t.size)
	if err != nil {
		return
	}
	c.Check()
	t.context.Run(func() {
		t.callback(res)
	})
}
