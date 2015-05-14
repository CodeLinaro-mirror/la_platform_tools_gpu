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

import "github.com/google/gxui"

func CreateDepthBufferPanel(appCtx *ApplicationContext) gxui.Control {
	theme := appCtx.Theme()

	image := theme.CreateImage()
	image.SetScalingMode(gxui.ScalingExpandGreedy)
	image.SetAspectMode(gxui.AspectCorrectLetterbox)

	appCtx.OnDepthBufferUpdate(func() {
		image.SetTexture(appCtx.DepthBuffer())
	})

	layout := theme.CreateLinearLayout()
	layout.SetDirection(gxui.TopToBottom)
	layout.AddChild(image)
	return layout
}
