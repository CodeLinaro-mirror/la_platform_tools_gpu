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
	"android.googlesource.com/platform/tools/gpu/_experimental/client/charts"
	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/service"
	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

const targetFrameTime = μs(16666)

func CreateProfilerPanel(appCtx *ApplicationContext) gxui.Control {
	gridlines := charts.DefaultGridlines
	gridlines.Format = func(t int) string { return μs(t).String() }
	gridlines.Multiples = []int{20, 50}
	chart := charts.NewBarChart(appCtx.Theme())
	chart.SetOrientation(gxui.Horizontal)
	chart.SetBackgroundBrush(gxui.CreateBrush(gxui.Gray10))
	chart.SetGridlines(gridlines)
	chart.OnBarDoubleClicked(func(idx int, ev gxui.MouseEvent) {
		frameId := appCtx.timingInfo.PerFrame[idx].ToAtomId
		appCtx.SelectAtom(atom.ID(frameId))
	})

	appCtx.OnTimingInfoUpdated(func() {
		chart.SetData(NewTimingData(appCtx.timingInfo))
	})

	return chart
}

type TimingData struct {
	frames   []int
	min, max int
}

func (d TimingData) Count() int {
	return len(d.frames)
}

func (d TimingData) Values(i int) []int {
	return []int{d.frames[i]}
}

func (d TimingData) Limits() (int, int) {
	rng := d.max - d.min
	return d.min - rng/10, d.max + rng/10
}

func (d TimingData) BarBrush(bar int, stack int, highlighted bool) gxui.Brush {
	if highlighted {
		return gxui.CreateBrush(d.BarBrush(bar, stack, false).Color.MulRGB(1.5).Saturate())
	}
	return gxui.CreateBrush(μs(d.Values(bar)[0]).ThesholdColor(targetFrameTime))
}

func (d TimingData) LabelBackgroundBrush(bar int, stack int) gxui.Brush {
	if stack < len(d.Values(bar)) {
		return gxui.CreateBrush(d.BarBrush(bar, stack, false).Color.MulRGB(0.5))
	} else {
		return gxui.CreateBrush(gxui.Gray20)
	}
}

func (d TimingData) LabelTextColor(bar int, stack int) gxui.Color {
	return gxui.Gray80
}

func NewTimingData(t service.TimingInfo) TimingData {
	frames := []int{}
	min, max := math.MaxInt, math.MinInt
	for _, timer := range t.PerFrame {
		μs := int(timer.Nanoseconds / 1e3)
		frames = append(frames, μs)
		min = math.Min(min, μs)
		max = math.Max(max, μs)
	}
	return TimingData{
		frames: frames,
		min:    min,
		max:    max,
	}
}
