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

package main

import (
	"flag"

	"android.googlesource.com/platform/tools/gpu/_experimental/client/charts"
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/math"
	"github.com/google/gxui/themes/dark"
)

func smoothrnd(key float32) float32 {
	freqs := []float32{3452.2341, 156.653, 13.345, 5.235, 0.1352}
	phase := []float32{34.136531, 9.48426, 2.3419, 7.691, 4.2769}
	v := float32(0)
	for i := range freqs {
		v += (0.5 + 0.5*math.Sinf(float32(key)*freqs[i]+phase[i])) * float32(i)
	}
	return v / float32(len(freqs))
}

type ChartData struct{}

func (ChartData) Count() int { return 1000 }
func (ChartData) Values(i int) []int {
	return []int{
		int(200 * smoothrnd(float32(i+16214)/100)),
		int(200 * smoothrnd(float32(i+43862)/100)),
		int(200 * smoothrnd(float32(i+29435)/100)),
	}
}
func (ChartData) Limits() (int, int) { return -100, 10000 }

func (ChartData) BarBrush(bar int, stack int, highlighted bool) gxui.Brush {
	color := []gxui.Color{
		{R: 0.00, G: 0.53, B: 0.74, A: 1},
		{R: 0.00, G: 0.73, B: 0.44, A: 1},
		{R: 0.63, G: 0.20, B: 0.24, A: 1},
	}[stack]
	if highlighted {
		color = color.MulRGB(1.5).Saturate()
	}
	return gxui.CreateBrush(color)
}

func (c ChartData) LabelBackgroundBrush(bar int, stack int) gxui.Brush {
	if stack < 3 {
		return c.BarBrush(bar, stack, false)
	} else {
		return gxui.CreateBrush(gxui.Gray80)
	}
}

func (c ChartData) LabelTextColor(bar int, stack int) gxui.Color {
	return gxui.Black
}

func appMain(driver gxui.Driver) {
	theme := dark.CreateTheme(driver)

	window := theme.CreateWindow(800, 600, "Profiler chart")
	window.OnClose(driver.Terminate)

	gridlines := charts.DefaultGridlines
	gridlines.Multiples = []int{2, 5}

	chart := charts.NewBarChart(theme)
	chart.SetData(ChartData{})
	chart.SetOrientation(gxui.Horizontal)
	chart.SetGridlines(gridlines)
	chart.SetBarPen(gxui.CreatePen(1.0, gxui.Color{R: 0, G: 0.4, B: 0.5, A: 1}))
	chart.OnDoubleClick(func(gxui.MouseEvent) {
		chart.SetOrientation(chart.Orientation().Flip())
	})
	window.AddChild(chart)
}

func maybeError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.Parse()
	gl.StartDriver(appMain)
}
