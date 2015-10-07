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
	"sort"

	"android.googlesource.com/platform/tools/gpu/service"
	"android.googlesource.com/platform/tools/gpu/service/path"
	"android.googlesource.com/platform/tools/gpu/task"
	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

const kFilmStripAdapterItemWidth = 200
const kFilmStripAdapterItemHeight = 150

type FilmStripAdapter struct {
	gxui.AdapterBase
	appCtx  *ApplicationContext
	device  *path.Device
	capture *path.Capture
	frames  []uint64
}

func CreateFilmStripAdapter(appCtx *ApplicationContext) *FilmStripAdapter {
	return &FilmStripAdapter{appCtx: appCtx}
}

func (a *FilmStripAdapter) UpdateFrames(capture *path.Capture, frames []uint64) {
	a.capture = capture
	a.frames = frames
	a.DataReplaced()
}

func (a *FilmStripAdapter) UpdateDevice(device *path.Device) {
	a.device = device
	a.DataReplaced()
}

func (a *FilmStripAdapter) Size(theme gxui.Theme) math.Size {
	return math.Size{W: kFilmStripAdapterItemWidth, H: kFilmStripAdapterItemHeight}
}

func (a *FilmStripAdapter) ItemAt(index int) gxui.AdapterItem {
	return a.frames[index]
}

func (a *FilmStripAdapter) ItemIndex(item gxui.AdapterItem) int {
	index := item.(uint64)
	return sort.Search(len(a.frames), func(i int) bool {
		return a.frames[i] >= index
	})
}

func (a *FilmStripAdapter) Count() int {
	return len(a.frames)
}

func (a *FilmStripAdapter) Create(theme gxui.Theme, index int) gxui.Control {
	w, h := kFilmStripAdapterItemWidth, kFilmStripAdapterItemHeight
	atomIndex := a.frames[index]
	p := a.capture.Atoms().Index(atomIndex)

	i := theme.CreateImage()
	i.SetAspectMode(gxui.AspectCorrectLetterbox)
	i.SetExplicitSize(math.Size{W: w, H: h})

	b := theme.CreateButton()
	b.SetDirection(gxui.TopToBottom)
	b.SetText(fmt.Sprintf("%.6d - Frame %d", atomIndex, index))
	b.SetPadding(math.Spacing{L: 3, T: 3, R: 3, B: 3})
	b.SetMargin(math.Spacing{L: 3, T: 3, R: 3, B: 3})
	b.OnClick(func(gxui.MouseEvent) { a.appCtx.events.Select(p) })
	b.AddChild(i)

	t := task.New()
	b.OnAttach(func() { t.Run(updateFramebufferThumbnail{a.appCtx, a.device, p, i}) })
	b.OnDetach(t.Cancel)

	return b
}

type updateFramebufferThumbnail struct {
	context *ApplicationContext
	device  *path.Device
	after   *path.Atom
	image   gxui.Image
}

func (t updateFramebufferThumbnail) Run(c task.CancelSignal) {
	settings := service.RenderSettings{
		MaxWidth:  kFilmStripAdapterItemWidth,
		MaxHeight: kFilmStripAdapterItemHeight,
	}
	if w, h, d, err := t.context.rpc.RequestColorBuffer(t.device, t.after, settings); err == nil {
		c.Check()
		t.context.Run(func() {
			tex := NewTexture(t.context.theme.Driver(), w, h, d)
			t.image.SetTexture(tex)
			t.context.toolTipController.AddToolTip(t.image, 0.7, func(math.Point) gxui.Control {
				large := t.context.theme.CreateImage()
				large.SetTexture(tex)
				large.SetAspectMode(gxui.AspectCorrectLetterbox)
				return large
			})
		})
	}
}
