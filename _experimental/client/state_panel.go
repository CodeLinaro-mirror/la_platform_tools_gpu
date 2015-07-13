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

func CreateStatePanel(appCtx *ApplicationContext) gxui.Control {
	adapter := NewStateAdapter(appCtx)
	tree := appCtx.theme.CreateTree()
	tree.SetAdapter(adapter)

	var state *path.State

	t := task.New()
	update := func() {
		t.Run(updateStateAdapter{appCtx, state, adapter})
	}

	appCtx.events.OnSelect(func(p path.Path) {
		if a := path.FindAtom(p); a != nil {
			if s := a.StateAfter(); !path.Equal(s, state) {
				state = s
				update()
			}
		}
		if s, a := path.FindAtomSlice(p); s != nil && a != nil {
			if s := a.Index(s.End - 1).StateAfter(); !path.Equal(s, state) {
				state = s
				update()
			}
		}
	})

	return tree
}

type updateStateAdapter struct {
	context *ApplicationContext
	state   *path.State
	adapter *StateAdapter
}

func (t updateStateAdapter) Run(c task.CancelSignal) {
	state, err := t.context.rpc.LoadState(t.state)
	if err != nil {
		return
	}
	c.Check()
	t.context.Run(func() {
		t.adapter.Update(state, t.state)
	})
}
