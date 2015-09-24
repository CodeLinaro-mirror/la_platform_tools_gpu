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

func CreateDepthBufferPanel(appCtx *ApplicationContext) gxui.Control {
	image := appCtx.theme.CreateImage()
	image.SetScalingMode(gxui.ScalingExpandGreedy)
	image.SetAspectMode(gxui.AspectCorrectLetterbox)

	var device *path.Device
	var after *path.Atom

	t := task.New()
	update := func() {
		if device != nil && after != nil {
			t.Run(updateDepthBuffer{appCtx, device, after, image})
		}
	}

	appCtx.events.OnSelect(func(p path.Path) {
		if d := path.FindDevice(p); d != nil && !path.Equal(d, device) {
			device = d
			update()
		}
		if a := lastAtom(p); a != nil && !path.Equal(a, after) {
			after = a
			update()
		}
	})

	layout := appCtx.theme.CreateLinearLayout()
	layout.SetDirection(gxui.TopToBottom)
	layout.AddChild(image)
	return layout
}

type updateDepthBuffer struct {
	context *ApplicationContext
	device  *path.Device
	after   *path.Atom
	image   gxui.Image
}

func (t updateDepthBuffer) Run(c task.CancelSignal) {
	if w, h, d, err := t.context.rpc.RequestDepthBuffer(t.device, t.after); err == nil {
		c.Check()
		t.context.Run(func() {
			t.image.SetTexture(NewTexture(t.context.theme.Driver(), w, h, d))
		})
	}
}
