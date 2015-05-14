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

	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

func AddTextToolTip(appCtx *ApplicationContext, target gxui.Control, text string) {
	appCtx.ToolTipController().AddToolTip(target, 0.7, func(p math.Point) gxui.Control {
		label := appCtx.Theme().CreateLabel()
		label.SetText(text)
		return label
	})
}

func CreateWireframeButton(appCtx *ApplicationContext) gxui.Button {
	theme := appCtx.Theme()

	canvas := theme.Driver().CreateCanvas(math.Size{W: 15, H: 15})
	canvas.DrawPolygon(wireframeIconPoly, gxui.TransparentPen, gxui.CreateBrush(gxui.Gray80))
	canvas.Complete()

	icon := theme.CreateImage()
	icon.SetCanvas(canvas)

	button := theme.CreateButton()
	button.AddChild(icon)
	button.SetType(gxui.ToggleButton)
	button.OnClick(func(gxui.MouseEvent) {
		appCtx.SetWireframe(button.IsChecked())
		appCtx.RequestReplay()
	})
	appCtx.OnWireframeChanged(func() {
		button.SetChecked(appCtx.Wireframe())
	})
	button.SetPadding(math.CreateSpacing(4))

	AddTextToolTip(appCtx, button, "Wireframe")

	return button
}

func CreateColorBufferPanel(appCtx *ApplicationContext) gxui.Control {
	theme := appCtx.Theme()

	image := theme.CreateImage()
	image.SetScalingMode(gxui.ScalingExpandGreedy)
	image.SetAspectMode(gxui.AspectCorrectLetterbox)

	appCtx.ToolTipController().AddToolTip(image, 1.0, func(p math.Point) gxui.Control {
		tex := image.Texture()
		if tex == nil {
			return nil
		}
		img := tex.Image()
		if img == nil {
			return nil
		}
		p, inImage := image.PixelAt(p)
		if !inImage {
			return nil
		}
		x, y := p.X, tex.SizePixels().H-p.Y
		c := img.At(x, y)

		r, g, b, a := c.RGBA()

		color := theme.CreateImage()
		color.SetExplicitSize(math.Size{W: 16, H: 16})
		color.SetBorderPen(gxui.Pen{
			Width: 1,
			Color: gxui.Gray40,
		})
		color.SetBackgroundBrush(gxui.Brush{
			Color: gxui.Color{
				R: float32(r) / 65535.0,
				G: float32(g) / 65535.0,
				B: float32(b) / 65535.0,
				A: float32(a) / 65535.0,
			},
		})

		label := theme.CreateLabel()
		label.SetText(fmt.Sprintf("X:%d, Y:%d\nR:%.3d G:%.3d B:%.3d A:%.3d",
			x, y, r/256, g/256, b/256, a/256))
		label.SetMultiline(true)

		layout := theme.CreateLinearLayout()
		layout.SetDirection(gxui.LeftToRight)
		layout.SetVerticalAlignment(gxui.AlignMiddle)
		layout.AddChild(color)
		layout.AddChild(label)

		return layout
	})

	appCtx.OnColorBufferUpdate(func() {
		image.SetTexture(appCtx.ColorBuffer())
	})

	layout := theme.CreateLinearLayout()
	layout.SetDirection(gxui.TopToBottom)
	layout.AddChild(CreateWireframeButton(appCtx))
	layout.AddChild(image)

	return layout
}
