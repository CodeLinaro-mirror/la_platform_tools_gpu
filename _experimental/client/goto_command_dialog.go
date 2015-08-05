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

	"android.googlesource.com/platform/tools/gpu/service/path"
	"github.com/google/gxui"
)

func CreateGotoCommandDialog(appCtx *ApplicationContext, atoms *path.Atoms) gxui.Window {
	label := appCtx.theme.CreateLabel()
	label.SetText("Goto command:")

	textbox := appCtx.theme.CreateTextBox()
	// textbox.SetText(fmt.Sprintf("%v", appCtx.SelectedAtomIndex())) // TODO
	textbox.SelectAll()

	layout := appCtx.theme.CreateLinearLayout()
	layout.SetDirection(gxui.LeftToRight)
	layout.AddChild(label)
	layout.AddChild(textbox)

	wnd := appCtx.theme.CreateWindow(200, 30, "Goto Command")
	wnd.AddChild(layout)

	wnd.OnKeyDown(func(ev gxui.KeyboardEvent) {
		if ev.Key == gxui.KeyEnter {
			var index uint64
			_, err := fmt.Sscanf(textbox.Text(), "%d", &index)
			if err == nil {
				appCtx.events.Select(atoms.Index(index))
				wnd.Close()
			}
		}
	})

	wnd.SetFocus(textbox)

	return wnd
}
