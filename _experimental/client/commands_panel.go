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
	"time"

	"android.googlesource.com/platform/tools/gpu/service"
	"android.googlesource.com/platform/tools/gpu/service/path"
	"android.googlesource.com/platform/tools/gpu/task"
	"github.com/google/gxui"
)

func CreateCommandsPanel(appCtx *ApplicationContext) gxui.Control {
	adapter := CreateCommandAdapter(appCtx)

	tree := appCtx.theme.CreateTree()
	tree.SetAdapter(adapter)

	timer := (*time.Timer)(nil)
	selectSelection := func() {
		appCtx.Run(func() {
			if p := adapter.Path(tree.Selected()); p != nil {
				appCtx.events.Select(p)
			}
		})
	}

	tree.OnSelectionChanged(func(gxui.AdapterItem) {
		d := time.Millisecond * 500
		if timer == nil {
			timer = time.AfterFunc(d, selectSelection)
		} else {
			timer.Reset(d)
		}
	})

	var device *path.Device
	var capture *path.Capture
	t := task.New()
	update := func() {
		if device != nil && capture != nil {
			t.Run(updateCommandAdapter{appCtx, device, capture, adapter})
		}
	}

	appCtx.events.OnSelect(func(p path.Path) {
		if d := path.FindDevice(p); d != nil && !path.Equal(d, device) {
			device = d
			update()
		}
		if c := path.FindCapture(p); c != nil && !path.Equal(c, capture) {
			capture = c
			update()
		}
		if a := path.FindAtom(p); a != nil {
			if item := adapter.Item(a); item != nil {
				tree.Select(item)
			}
		}
	})

	tree.OnKeyPress(func(ev gxui.KeyboardEvent) {
		if ev.Key == gxui.KeyG && ev.Modifier&gxui.ModControl != 0 {
			CreateGotoCommandDialog(appCtx, capture.Atoms())
		}
	})

	return tree
}

type updateCommandAdapter struct {
	context *ApplicationContext
	device  *path.Device
	capture *path.Capture
	adapter *CommandAdapter
}

func (t updateCommandAdapter) Run(c task.CancelSignal) {
	atoms, err := t.context.rpc.LoadAtoms(t.capture.Atoms())
	if err != nil {
		return
	}
	c.Check()
	hierarchy, err := t.context.rpc.LoadHierarchy(t.capture.Hierarchy())
	if err != nil {
		return
	}
	c.Check()
	t.context.RunSync(func() {
		t.adapter.UpdateDevice(t.device)
		t.adapter.UpdateTimings(service.TimingInfo{})
		t.adapter.UpdateAtoms(t.capture, atoms, hierarchy)
	})
	flags := service.TimingFlagsTimingCPU |
		service.TimingFlagsTimingPerCommand |
		service.TimingFlagsTimingPerDrawCall |
		service.TimingFlagsTimingPerFrame
	timings, err := t.context.rpc.LoadTiming(t.device, t.capture, flags)
	if err != nil {
		return
	}
	c.Check()
	t.context.Run(func() {
		t.adapter.UpdateTimings(timings)
	})
}
