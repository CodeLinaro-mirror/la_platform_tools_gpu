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
	"android.googlesource.com/platform/tools/gpu/service/path"
	"android.googlesource.com/platform/tools/gpu/task"
	"github.com/google/gxui"
)

func CreateFramesPanel(appCtx *ApplicationContext) gxui.Control {
	adapter := CreateFilmStripAdapter(appCtx)

	list := appCtx.theme.CreateList()
	list.SetAdapter(adapter)
	list.SetOrientation(gxui.Horizontal)

	var capture *path.Capture
	var device *path.Device

	t := task.New()
	update := func() {
		t.Run(updateFilmStripAdapter{appCtx, capture, adapter})
	}

	appCtx.events.OnSelect(func(p path.Path) {
		if c := path.FindCapture(p); c != nil && !path.Equal(c, capture) {
			capture = c
			update()
		}
		if d := path.FindDevice(p); d != nil && !path.Equal(d, device) {
			device = d
			adapter.UpdateDevice(d)
		}
	})

	list.OnSelectionChanged(func(item gxui.AdapterItem) {
		appCtx.events.Select(capture.Atoms().Index(item.(uint64)))
	})

	return list
}

type updateFilmStripAdapter struct {
	context *ApplicationContext
	capture *path.Capture
	adapter *FilmStripAdapter
}

func (t updateFilmStripAdapter) Run(c task.CancelSignal) {
	atoms, err := t.context.rpc.LoadAtoms(t.capture.Atoms())
	if err != nil {
		return
	}
	c.Check()
	frames := []uint64{}
	for i, t := range atoms {
		if t.Flags().IsEndOfFrame() {
			frames = append(frames, uint64(i))
		}
	}
	c.Check()
	t.context.Run(func() {
		t.adapter.UpdateFrames(t.capture, frames)
	})
}
