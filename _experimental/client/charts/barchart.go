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
	"github.com/google/gxui/mixins/parts"
)

type BarChartData interface {
	Count() int
	Values(bar int) []int
	Limits() (int, int)
	BarBrush(bar int, stack int, highlighted bool) gxui.Brush
	LabelBackgroundBrush(bar int, stack int) gxui.Brush
	LabelTextColor(bar int, stack int) gxui.Color
}

type BarChart struct {
	base.Container
	ZoomWindow
	parts.BackgroundBorderPainter

	theme              gxui.Theme
	data               BarChartData
	barWidth           int
	barPadding         int
	barPen             gxui.Pen
	highlightedBarPen  gxui.Pen
	orientation        gxui.Orientation
	gridLines          GridLines
	highlightedIdx     int
	onBarClicked       gxui.Event
	onBarDoubleClicked gxui.Event
}

func NewBarChart(theme gxui.Theme) *BarChart {
	c := &BarChart{}
	c.Container.Init(c, theme)
	c.ZoomWindow.Init(c, theme)
	c.BackgroundBorderPainter.Init(c)
	c.theme = theme
	c.barWidth = 50
	c.barPadding = 4
	c.barPen = gxui.CreatePen(1, gxui.Gray50)
	c.highlightedBarPen = gxui.CreatePen(1, gxui.Gray70)
	c.gridLines = DefaultGridlines
	c.SetMouseEventTarget(true)
	return c
}

func (c *BarChart) OnBarClicked(f func(idx int, ev gxui.MouseEvent)) gxui.EventSubscription {
	if c.onBarClicked == nil {
		c.onBarClicked = gxui.CreateEvent(f)
	}
	return c.onBarClicked.Listen(f)
}

func (c *BarChart) OnBarDoubleClicked(f func(idx int, ev gxui.MouseEvent)) gxui.EventSubscription {
	if c.onBarDoubleClicked == nil {
		c.onBarDoubleClicked = gxui.CreateEvent(f)
	}
	return c.onBarDoubleClicked.Listen(f)
}

func (c *BarChart) DesiredSize(min, max math.Size) math.Size {
	return max
}

func (c *BarChart) SetData(data BarChartData) {
	c.data = data
	c.resetZoom()
}

func (c *BarChart) Data() BarChartData {
	return c.data
}

func (c *BarChart) Orientation() gxui.Orientation {
	return c.orientation
}

func (c *BarChart) SetOrientation(o gxui.Orientation) {
	if c.orientation != o {
		c.orientation = o
		c.resetZoom()
	}
}

func (c *BarChart) SetBarPen(pen gxui.Pen) {
	if c.barPen != pen {
		c.barPen = pen
		c.Redraw()
	}
}

func (c *BarChart) SetHighlightedBarPen(pen gxui.Pen) {
	if c.highlightedBarPen != pen {
		c.highlightedBarPen = pen
		c.Redraw()
	}
}

func (c *BarChart) SetGridlines(gridlines GridLines) {
	c.gridLines = gridlines
	c.Redraw()
}

func (c *BarChart) SetSize(s math.Size) {
	c.SetMinZoomWindowSize(s)
	c.Container.SetSize(s)
}

func (c *BarChart) Paint(canvas gxui.Canvas) {
	r := canvas.Size().Rect()
	c.PaintBackground(canvas, r)

	if c.data != nil {
		viewBounds := c.ViewBounds()
		paddingWeight, barWeight := c.paddingWeight(), c.barWeight()

		if c.data.Count() > 0 {
			min, max := c.windowDataRange()
			for i := min; i <= max; i++ {
				c.drawBars(canvas, i, viewBounds, paddingWeight, barWeight)
			}
		}
		if c.orientation.Horizontal() {
			c.gridLines.PaintVertical(c.theme, canvas, c.zoomWindow, viewBounds)
		} else {
			c.gridLines.PaintHorizontal(c.theme, canvas, c.zoomWindow, viewBounds)
		}
		if c.highlightedIdx > 0 && c.highlightedIdx < c.data.Count() {
			c.drawLabels(canvas, c.highlightedIdx, viewBounds, paddingWeight)
		}
	}
	c.Container.Paint(canvas)
	c.PaintBorder(canvas, r)
}

func (c *BarChart) MouseMove(e gxui.MouseEvent) {
	c.setHighlighted(c.barAt(e.Point))
	c.ZoomWindow.MouseMove(e)
}

func (c *BarChart) Click(e gxui.MouseEvent) bool {
	if idx := c.barAt(e.Point); idx >= 0 {
		if c.onBarClicked != nil {
			c.onBarClicked.Fire(idx, e)
		}
		return true
	}
	return false
}

func (c *BarChart) DoubleClick(e gxui.MouseEvent) bool {
	if idx := c.barAt(e.Point); idx >= 0 {
		if c.onBarDoubleClicked != nil {
			c.onBarDoubleClicked.Fire(idx, e)
		}
		return true
	}
	return false
}

func (c *BarChart) barAt(p math.Point) int {
	p = p.Remap(c.ViewBounds(), c.zoomWindow)
	if c.orientation.Horizontal() {
		idx := p.X / c.barSeparation()
		flatRect := c.barFlatRect(idx, c.paddingWeight())
		if p.X >= flatRect.Min.X && p.X < flatRect.Max.X {
			return idx
		}
	} else {
		idx := p.Y / c.barSeparation()
		flatRect := c.barFlatRect(idx, c.paddingWeight())
		if p.Y >= flatRect.Min.Y && p.Y < flatRect.Max.Y {
			return idx
		}
	}
	return -1
}

func (c *BarChart) resetZoom() {
	if c.data != nil {
		var r math.Rect
		min, max := c.data.Limits()
		if c.orientation.Horizontal() {
			r = math.CreateRect(0, -max, c.barWidth*c.data.Count(), -min)
		} else {
			r = math.CreateRect(min, 0, max, c.barWidth*c.data.Count())
		}
		c.SetZoomBounds(r)
		c.SetZoomWindow(r)
		c.Redraw()
	}
}

func (c *BarChart) majorAxisUnitsPerDip() float32 {
	viewBounds := c.ViewBounds()
	if c.orientation.Horizontal() {
		return float32(c.zoomWindow.W()) / float32(viewBounds.Size().W)
	} else {
		return float32(c.zoomWindow.H()) / float32(viewBounds.Size().H)
	}
}

func (c *BarChart) paddingWeight() float32 {
	paddingDips := float32(c.barPadding) / c.majorAxisUnitsPerDip()
	return math.SmoothStep(paddingDips, 1.5, 2)
}

func (c *BarChart) barWeight() float32 {
	barDips := float32(c.barSeparation()) / c.majorAxisUnitsPerDip()
	return math.SmoothStep(barDips, 2+c.barPen.Width*2, 3+c.barPen.Width*2)
}

func (c *BarChart) barSeparation() int {
	return c.barWidth + c.barPadding
}

func (c *BarChart) setHighlighted(idx int) {
	if c.highlightedIdx != idx {
		c.highlightedIdx = idx
		c.Redraw()
	}
}

func (c *BarChart) barRects(barIdx int, viewBounds math.Rect, paddingWeight float32) []math.Rect {
	r := c.barFlatRect(barIdx, paddingWeight)

	values := c.data.Values(barIdx)
	rects := make([]math.Rect, 0, len(values))

	for _, v := range values {
		if c.orientation.Horizontal() {
			r.Max.Y = r.Min.Y
			r.Min.Y -= v
		} else {
			r.Min.X = r.Max.X
			r.Max.X += v
		}
		if r.Size().Area() > 0 {
			rects = append(rects, r.Remap(c.zoomWindow, viewBounds))
		}
	}

	return rects
}

func (c *BarChart) drawBars(
	canvas gxui.Canvas,
	barIdx int,
	viewBounds math.Rect,
	paddingWeight, barWeight float32) {

	highlighted := barIdx == c.highlightedIdx

	pen := c.barPen
	if highlighted {
		pen = c.highlightedBarPen
	}

	if paddingWeight < 0.5 {
		pen.Width *= barWeight
	}

	for stackIdx, bar := range c.barRects(barIdx, viewBounds, paddingWeight) {
		brush := c.data.BarBrush(barIdx, stackIdx, highlighted)
		if barWeight == 0 {
			canvas.DrawRect(bar, brush)
		} else {
			r := 2 * barWeight
			if stackIdx > 0 {
				canvas.DrawRoundedRect(bar, r, r, r, r, pen, brush)
			} else {
				canvas.DrawRoundedRect(bar, r, r, 0, 0, pen, brush)
			}
		}
	}
}

func (c *BarChart) drawLabels(
	canvas gxui.Canvas,
	barIdx int,
	viewBounds math.Rect,
	paddingWeight float32) {

	font := c.theme.DefaultFont()
	values := c.data.Values(barIdx)
	rects := c.barRects(barIdx, viewBounds, paddingWeight)
	bounds := c.ViewBounds()

	// Draw each of the stack labels
	for stackIdx, bar := range rects {
		brush := c.data.LabelBackgroundBrush(barIdx, stackIdx)
		color := c.data.LabelTextColor(barIdx, stackIdx)
		b := &gxui.TextBlock{
			Runes: []rune(c.gridLines.Format(values[stackIdx])),
			H:     gxui.AlignCenter,
			V:     gxui.AlignMiddle,
		}
		b.AlignRect = font.Measure(b).CenteredRect().Offset(bar.Mid()).ExpandI(3).Constrain(c.ViewBounds())
		if c.orientation.Horizontal() {
			bounds.Max.Y = b.AlignRect.Min.Y
		} else {
			bounds.Min.X = b.AlignRect.Max.X
		}
		canvas.DrawRoundedRect(b.AlignRect, 3, 3, 3, 3, c.barPen, brush)
		canvas.DrawRunes(font, b.Runes, font.Layout(b), color)
	}

	// Draw the label at the top of the stack
	if count := len(values); count > 1 {
		r := rects[count-1]
		b := &gxui.TextBlock{
			Runes: []rune(c.gridLines.Format(sum(values))),
			H:     gxui.AlignCenter,
			V:     gxui.AlignMiddle,
		}
		if c.orientation.Horizontal() {
			b.AlignRect = font.Measure(b).CenteredRect().
				Offset(r.TC()).
				ExpandI(3).
				OffsetY(-r.H()).
				Constrain(c.ViewBounds())
		} else {
			b.AlignRect = font.Measure(b).CenteredRect().
				Offset(r.MR()).
				ExpandI(3).
				OffsetX(r.W()).
				Constrain(c.ViewBounds())
		}
		brush := c.data.LabelBackgroundBrush(barIdx, count)
		color := c.data.LabelTextColor(barIdx, count)
		canvas.DrawRoundedRect(b.AlignRect, 3, 3, 3, 3, c.barPen, brush)
		canvas.DrawRunes(font, b.Runes, font.Layout(b), color)
	}
}

func (c *BarChart) barFlatRect(barIdx int, paddingWeight float32) math.Rect {
	separation := c.barSeparation()
	padding := math.Round(float32(c.barPadding) * paddingWeight / 2)

	var r math.Rect
	if c.orientation.Horizontal() {
		r.Min.X = (barIdx+0)*separation + padding
		r.Max.X = (barIdx+1)*separation - padding
	} else {
		r.Min.Y = (barIdx+0)*separation + padding
		r.Max.Y = (barIdx+1)*separation - padding
	}
	return r
}

func (c *BarChart) windowDataRange() (int, int) {
	cnt := c.data.Count()
	w := c.barWidth + c.barPadding
	if c.orientation.Horizontal() {
		return math.Clamp(c.zoomWindow.Min.X/w, 0, cnt-1),
			math.Clamp((c.zoomWindow.Max.X+w-1)/w, 0, cnt-1)
	} else {
		return math.Clamp(c.zoomWindow.Min.Y/w, 0, cnt-1),
			math.Clamp((c.zoomWindow.Max.Y+w-1)/w, 0, cnt-1)
	}
}
