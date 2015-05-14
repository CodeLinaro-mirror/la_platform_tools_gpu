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
	"fmt"

	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

var DefaultGridlines = GridLines{
	Brush:     gxui.CreateBrush(gxui.Gray40),
	TextColor: gxui.Gray70,
	Font:      nil,
	Format:    func(v int) string { return fmt.Sprintf("%d", v) },
	Density:   30,
	Multiples: []int{2, 5},
}

type GridLines struct {
	Brush     gxui.Brush
	TextColor gxui.Color
	Font      gxui.Font
	Format    func(int) string
	Density   int // DIPs per line
	Multiples []int
}

func (g GridLines) PaintVertical(theme gxui.Theme, canvas gxui.Canvas, zoomWindow math.Rect, viewBounds math.Rect) {
	font := g.Font
	if font == nil {
		font = theme.DefaultFont()
	}
	step := quantize(zoomWindow.H(), viewBounds.Size().H/g.Density, g.Multiples...)
	from := (zoomWindow.Min.Y / step) * step
	for i := from; i < zoomWindow.Max.Y; i += step {
		y := math.Point{Y: i}.Remap(zoomWindow, viewBounds.Size().Rect()).Y
		b := &gxui.TextBlock{
			Runes:     []rune(g.Format(-i)),
			AlignRect: math.CreateRect(0, y, viewBounds.Size().W, y+1),
		}
		canvas.DrawRect(b.AlignRect, g.Brush)
		canvas.DrawRunes(font, b.Runes, font.Layout(b), g.TextColor)
	}
}

func (g GridLines) PaintHorizontal(theme gxui.Theme, canvas gxui.Canvas, zoomWindow math.Rect, viewBounds math.Rect) {

}
