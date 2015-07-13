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

	"android.googlesource.com/platform/tools/gpu/service"
	"android.googlesource.com/platform/tools/gpu/service/path"
	"android.googlesource.com/platform/tools/gpu/task"
	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

func AddTextToolTip(appCtx *ApplicationContext, target gxui.Control, text string) {
	appCtx.toolTipController.AddToolTip(target, 0.7, func(p math.Point) gxui.Control {
		label := appCtx.theme.CreateLabel()
		label.SetText(text)
		return label
	})
}

func CreateWireframeButton(appCtx *ApplicationContext, changed func(bool)) gxui.Button {
	canvas := appCtx.theme.Driver().CreateCanvas(math.Size{W: 15, H: 15})
	canvas.DrawPolygon(wireframeIconPoly, gxui.TransparentPen, gxui.CreateBrush(gxui.Gray80))
	canvas.Complete()

	icon := appCtx.theme.CreateImage()
	icon.SetCanvas(canvas)

	button := appCtx.theme.CreateButton()
	button.AddChild(icon)
	button.SetType(gxui.ToggleButton)
	button.OnClick(func(gxui.MouseEvent) { changed(button.IsChecked()) })
	button.SetPadding(math.CreateSpacing(4))

	AddTextToolTip(appCtx, button, "Wireframe")

	return button
}

func CreateColorBufferPanel(appCtx *ApplicationContext) gxui.Control {
	theme := appCtx.theme

	image := theme.CreateImage()
	image.SetScalingMode(gxui.ScalingExpandGreedy)
	image.SetAspectMode(gxui.AspectCorrectLetterbox)

	appCtx.toolTipController.AddToolTip(image, 1.0, func(p math.Point) gxui.Control {
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

	var device *path.Device
	var after *path.Atom
	var wireframe bool

	t := task.New()
	update := func() {
		if device != nil && after != nil {
			t.Run(updateColorBuffer{appCtx, device, after, wireframe, image})
		}
	}

	appCtx.events.OnSelect(func(p path.Path) {
		if d := path.FindDevice(p); d != nil && !path.Equal(d, device) {
			device = d
			update()
		}
		if a := path.FindAtom(p); a != nil && !path.Equal(a, after) {
			after = a
			update()
		}
		if s, a := path.FindAtomSlice(p); s != nil {
			if i := a.Index(s.End - 1); !path.Equal(i, after) {
				after = i
				update()
			}
		}
	})

	wireframeChanged := func(w bool) {
		wireframe = w
		update()
	}

	layout := theme.CreateLinearLayout()
	layout.SetDirection(gxui.TopToBottom)
	layout.AddChild(CreateWireframeButton(appCtx, wireframeChanged))
	layout.AddChild(image)

	return layout
}

type updateColorBuffer struct {
	context   *ApplicationContext
	device    *path.Device
	after     *path.Atom
	wireframe bool
	image     gxui.Image
}

func (t updateColorBuffer) Run(c task.CancelSignal) {
	settings := service.RenderSettings{
		MaxWidth:  0xffff,
		MaxHeight: 0xffff,
		Wireframe: t.wireframe,
	}
	if w, h, d, err := t.context.rpc.RequestColorBuffer(t.device, t.after, settings); err == nil {
		c.Check()
		t.context.Run(func() {
			t.image.SetTexture(NewColorTexture(t.context.theme.Driver(), w, h, d))
		})
	}
}
