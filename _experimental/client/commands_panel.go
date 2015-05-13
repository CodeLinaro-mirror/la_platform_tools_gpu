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

	"github.com/google/gxui"
)

func CreateCommandsPanel(appCtx *ApplicationContext) gxui.Control {
	theme := appCtx.Theme()

	adapter := CreateCommandAdapter(appCtx)

	tree := theme.CreateTree()
	tree.SetAdapter(adapter)

	timer := (*time.Timer)(nil)
	selectSelection := func() {
		appCtx.Run(func() {
			r := adapter.AtomRange(tree.Selected())
			if !r.Contains(appCtx.SelectedAtomID()) {
				appCtx.SelectAtom(r.Last())
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

	appCtx.OnHierarchyUpdated(func() {
		adapter.SetRoot(appCtx.Hierarchy())
	})
	appCtx.OnTimingInfoUpdated(func() {
		adapter.SetRoot(appCtx.Hierarchy())
	})
	appCtx.OnAtomSelected(func() {
		r := adapter.AtomRange(tree.Selected())
		if appCtx.SelectedAtomID() != r.Last() {
			tree.Select(adapter.Item(appCtx.SelectedAtomID()))
		}
	})

	tree.OnKeyPress(func(ev gxui.KeyboardEvent) {
		if ev.Key == gxui.KeyG && ev.Modifier&gxui.ModControl != 0 {
			CreateGotoCommandDialog(appCtx)
		}
	})

	return tree
}
