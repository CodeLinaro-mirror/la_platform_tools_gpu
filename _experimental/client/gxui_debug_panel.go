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
	"runtime"
	"runtime/debug"

	"github.com/google/gxui"
)

func CreateGxuiDebug(appCtx *ApplicationContext, window gxui.Window, driver gxui.Driver) gxui.Control {
	theme := appCtx.Theme()

	layout := theme.CreateLinearLayout()
	layout.SetDirection(gxui.TopToBottom)

	dumpButton := theme.CreateButton()
	dumpButton.SetText("Dump hierarchy")
	dumpButton.OnClick(func(gxui.MouseEvent) { gxui.Dump(window) })

	breadcrumbLabel := theme.CreateLabel()
	window.OnMouseMove(func(ev gxui.MouseEvent) {
		breadcrumbLabel.SetText(gxui.BreadcrumbsAt(window, ev.Point))
	})

	debugStatsLabel := theme.CreateLabel()
	debugStatsLabel.SetMultiline(true)
	debugStatsLabel.SetVerticalAlignment(gxui.AlignTop)

	gcNowButton := theme.CreateButton()
	gcNowButton.SetText("GC now")
	gcNowButton.OnClick(func(gxui.MouseEvent) {
		runtime.GC()
	})

	disableGCButton := theme.CreateButton()
	disableGCButton.SetText("Disable GC")
	disableGCButton.OnClick(func(gxui.MouseEvent) {
		disabled := !disableGCButton.IsChecked()
		disableGCButton.SetChecked(disabled)
		if disabled {
			debug.SetGCPercent(-1)
		} else {
			debug.SetGCPercent(100)
		}
	})

	//	updateStatsButton := theme.CreateButton()
	//	updateStatsButton.SetText("Update stats")
	//	updateStatsButton.OnClick(func(gxui.MouseEvent) {
	//		debugStatsLabel.SetText(window.Viewport().Stats())
	//	})

	layout.AddChild(dumpButton)
	layout.AddChild(breadcrumbLabel)
	layout.AddChild(gcNowButton)
	layout.AddChild(disableGCButton)
	//	layout.AddChild(updateStatsButton)
	layout.AddChild(debugStatsLabel)

	scrollLayout := theme.CreateScrollLayout()
	scrollLayout.SetChild(layout)
	return scrollLayout
}
