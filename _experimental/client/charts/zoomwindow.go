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

package charts

import (
	"github.com/google/gxui"
	"github.com/google/gxui/math"
	"github.com/google/gxui/mixins/base"
)

const zoomScale = 1.01

type InputKind int

const (
	ScrollX = InputKind(iota)
	ScrollY
	DragX
	DragY
)

type ActionKind int

const (
	Nop = ActionKind(iota)
	PanX
	PanY
	ZoomX
	ZoomY
	ZoomXY
)

type InputHandler func(InputKind, gxui.Orientation, gxui.MouseEvent) ActionKind

func DefaultInputHandler(input InputKind, o gxui.Orientation, ev gxui.MouseEvent) ActionKind {
	switch input {
	case ScrollY:
		return ZoomXY
	case DragX:
		return PanX
	case DragY:
		return PanY
	default:
		return Nop
	}
}

type ZoomWindowOuter interface {
	base.ContainerOuter
	Orientation() gxui.Orientation
}

type ZoomWindow struct {
	outer             ZoomWindowOuter
	minZoomWindowSize math.Size
	zoomBounds        math.Rect
	zoomWindow        math.Rect
	hScrollbar        gxui.ScrollBar
	vScrollbar        gxui.ScrollBar
	hScrollbarChild   *gxui.Child
	vScrollbarChild   *gxui.Child
	mouseDown         bool
	mousePos          math.Point

	InputHandler InputHandler
}

func (z *ZoomWindow) Init(outer ZoomWindowOuter, theme gxui.Theme) {
	z.outer = outer
	z.minZoomWindowSize.W = 10
	z.minZoomWindowSize.H = 10
	z.hScrollbar = theme.CreateScrollBar()
	z.vScrollbar = theme.CreateScrollBar()
	z.hScrollbar.SetOrientation(gxui.Horizontal)
	z.vScrollbar.SetOrientation(gxui.Vertical)
	z.hScrollbarChild = z.outer.AddChild(z.hScrollbar)
	z.vScrollbarChild = z.outer.AddChild(z.vScrollbar)
	z.hScrollbar.OnScroll(func(from, to int) {
		w := z.zoomWindow
		b := z.zoomBounds.Min.X
		z.SetZoomWindow(math.CreateRect(from+b, w.Min.Y, to+b, w.Max.Y))
	})
	z.vScrollbar.OnScroll(func(from, to int) {
		w := z.zoomWindow
		b := z.zoomBounds.Min.Y
		z.SetZoomWindow(math.CreateRect(w.Min.X, from+b, w.Max.X, to+b))
	})
	z.InputHandler = DefaultInputHandler
}

func (c *BarChart) LayoutChildren() {
	s := c.Size()
	hs := c.hScrollbar.DesiredSize(math.ZeroSize, s)
	vs := c.vScrollbar.DesiredSize(math.ZeroSize, s)
	c.hScrollbarChild.Layout(math.CreateRect(0, s.H-hs.H, s.W-vs.W, s.H))
	c.vScrollbarChild.Layout(math.CreateRect(s.W-vs.W, 0, s.W, s.H-vs.H))
}

func (z *ZoomWindow) SetMinZoomWindowSize(size math.Size) {
	z.minZoomWindowSize = size
	z.SetZoomWindow(z.zoomWindow)
}

func (z *ZoomWindow) SetZoomWindow(zoomWindow math.Rect) {
	minRect := z.minZoomWindowSize.CenteredRect().Offset(zoomWindow.Mid())
	zoomWindow = zoomWindow.Union(minRect).Constrain(z.zoomBounds)
	if z.zoomWindow != zoomWindow {
		z.zoomWindow = zoomWindow
		z.hScrollbar.SetScrollPosition(zoomWindow.Min.X-z.zoomBounds.Min.X, zoomWindow.Max.X-z.zoomBounds.Min.X)
		z.vScrollbar.SetScrollPosition(zoomWindow.Min.Y-z.zoomBounds.Min.Y, zoomWindow.Max.Y-z.zoomBounds.Min.Y)
		z.outer.Redraw()
	}
}

func (z *ZoomWindow) SetZoomBounds(zoomBounds math.Rect) {
	z.zoomBounds = zoomBounds
	z.hScrollbar.SetScrollLimit(zoomBounds.W())
	z.vScrollbar.SetScrollLimit(zoomBounds.H())
	z.SetZoomWindow(z.zoomWindow)
}

func (z *ZoomWindow) ViewBounds() math.Rect {
	bounds := z.outer.Size().Rect()
	if z.vScrollbar.IsVisible() {
		bounds.Max.X -= z.vScrollbar.Size().W
	}
	if z.hScrollbar.IsVisible() {
		bounds.Max.Y -= z.hScrollbar.Size().H
	}
	return bounds
}

func (z *ZoomWindow) MouseDown(e gxui.MouseEvent) {
	for _, c := range z.outer.Children() {
		if c.Bounds().Contains(e.Point) {
			return
		}
	}
	z.mouseDown = true
	z.mousePos = e.Point
}

func (z *ZoomWindow) MouseUp(e gxui.MouseEvent) {
	z.mouseDown = false
}

func (z *ZoomWindow) MouseMove(e gxui.MouseEvent) {
	lastPos := z.mousePos
	z.mousePos = e.Point
	if z.mouseDown {
		delta := lastPos.Sub(e.Point).Vec2()
		if delta.X != 0 {
			z.handleAction(z.InputHandler(DragX, z.outer.Orientation(), e), delta.X)
		}
		if delta.Y != 0 {
			z.handleAction(z.InputHandler(DragY, z.outer.Orientation(), e), delta.Y)
		}
	}
}

func (z *ZoomWindow) MouseScroll(e gxui.MouseEvent) bool {
	if e.ScrollX != 0 {
		z.handleAction(z.InputHandler(ScrollX, z.outer.Orientation(), e), float32(e.ScrollX))
	}
	if e.ScrollY != 0 {
		z.handleAction(z.InputHandler(ScrollY, z.outer.Orientation(), e), float32(e.ScrollY))
	}
	return true
}

func (z *ZoomWindow) handleAction(action ActionKind, value float32) {
	switch action {
	case PanX:
		value *= float32(z.zoomWindow.W()) / float32(z.ViewBounds().W())
		z.SetZoomWindow(z.zoomWindow.OffsetX(int(value)))
	case PanY:
		value *= float32(z.zoomWindow.H()) / float32(z.ViewBounds().H())
		z.SetZoomWindow(z.zoomWindow.OffsetY(int(value)))
	case ZoomX, ZoomY, ZoomXY:
		center := z.mousePos.Remap(z.ViewBounds(), z.zoomWindow)
		scale := math.Powf(zoomScale, float32(value))
		s := math.Vec2{X: 1, Y: 1}
		if action != PanY {
			s.X = math.Maxf(scale, float32(z.minZoomWindowSize.W)/float32(z.zoomWindow.W()))
		}
		if action != PanX {
			s.Y = math.Maxf(scale, float32(z.minZoomWindowSize.H)/float32(z.zoomWindow.H()))
		}
		z.SetZoomWindow(z.zoomWindow.ScaleAt(center, s))
	}
}
