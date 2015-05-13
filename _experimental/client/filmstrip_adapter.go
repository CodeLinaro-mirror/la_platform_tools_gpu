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

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/_experimental/client/schema"
	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

const kFilmStripAdapterItemWidth = 200
const kFilmStripAdapterItemHeight = 150

type FilmStripAdapter struct {
	gxui.AdapterBase
	appCtx *ApplicationContext
	frames []atom.ID
}

func CreateFilmStripAdapter(appCtx *ApplicationContext) *FilmStripAdapter {
	return &FilmStripAdapter{appCtx: appCtx}
}

func (a *FilmStripAdapter) SetAtoms(atoms []schema.Atom) {
	a.frames = []atom.ID{}
	for i, t := range atoms {
		if t.Info.IsEndOfFrame {
			a.frames = append(a.frames, atom.ID(i))
		}
	}
	a.DataReplaced()
}

func (a *FilmStripAdapter) Size(theme gxui.Theme) math.Size {
	return math.Size{W: kFilmStripAdapterItemWidth, H: kFilmStripAdapterItemHeight}
}

func (a *FilmStripAdapter) ItemAt(index int) gxui.AdapterItem {
	return a.frames[index]
}

func (a *FilmStripAdapter) ItemIndex(item gxui.AdapterItem) int {
	atomID := item.(atom.ID)
	return sort.Search(len(a.frames), func(i int) bool {
		return a.frames[i] >= atomID
	})
}

func (a *FilmStripAdapter) Count() int {
	return len(a.frames)
}

func (a *FilmStripAdapter) Create(t gxui.Theme, index int) gxui.Control {
	atomID := a.frames[index]
	i := t.CreateImage()
	i.SetAspectMode(gxui.AspectCorrectCrop)
	i.SetExplicitSize(math.Size{W: kFilmStripAdapterItemWidth, H: kFilmStripAdapterItemHeight})

	b := t.CreateButton()
	b.SetDirection(gxui.TopToBottom)
	b.SetText(fmt.Sprintf("%.6d - Frame %d", atomID, index))
	b.SetPadding(math.Spacing{L: 3, T: 3, R: 3, B: 3})
	b.SetMargin(math.Spacing{L: 3, T: 3, R: 3, B: 3})
	b.OnClick(func(gxui.MouseEvent) {
		a.appCtx.SelectAtom(atomID)
	})
	b.AddChild(i)

	var cancel chan<- struct{}
	cancelThumbnail := func() {
		if cancel != nil {
			close(cancel)
			cancel = nil
		}
	}
	requestThumbnail := func() {
		cancelThumbnail()
		cancel = a.appCtx.RequestThumbnail(atomID, kFilmStripAdapterItemWidth, kFilmStripAdapterItemHeight, i.SetTexture)
	}

	var subscription gxui.EventSubscription
	b.OnAttach(func() {
		requestThumbnail()
		subscription = a.appCtx.OnDeviceSelected(requestThumbnail)
	})
	b.OnDetach(func() {
		cancelThumbnail()
		subscription.Unlisten()
	})

	return b
}
