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
	"android.googlesource.com/platform/tools/gpu/service"
	"android.googlesource.com/platform/tools/gpu/service/path"
	"android.googlesource.com/platform/tools/gpu/task"
	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

const targetFrameTime = μs(16666)

func CreateProfilerPanel(appCtx *ApplicationContext) gxui.Control {
	gridlines := charts.DefaultGridlines
	gridlines.Format = func(t int) string { return μs(t).String() }
	gridlines.Multiples = []int{20, 50}
	chart := charts.NewBarChart(appCtx.theme)
	chart.SetOrientation(gxui.Horizontal)
	chart.SetBackgroundBrush(gxui.CreateBrush(gxui.Gray10))
	chart.SetGridlines(gridlines)

	var device *path.Device
	var capture *path.Capture

	t := task.New()
	update := func() {
		if device != nil && capture != nil {
			t.Run(updateProfilerPanel{appCtx, device, capture, chart})
		}
	}

	appCtx.events.OnSelect(func(p path.Path) {
		if d := path.FindDevice(p); d != nil && !path.Equal(d, device) {
			device = d
			update()
		}
		if c := path.FindCapture(p); c != nil && !path.Equal(c, capture) {
			capture = c
			update()
		}
	})

	chart.OnBarDoubleClicked(func(idx int, ev gxui.MouseEvent) {
		data := chart.Data().(TimingData)
		atomIndex := data.timings.PerFrame[idx].ToAtomID
		appCtx.events.Select(capture.Atoms().Index(atomIndex))
	})

	return chart
}

type updateProfilerPanel struct {
	context *ApplicationContext
	device  *path.Device
	capture *path.Capture
	chart   *charts.BarChart
}

func (t updateProfilerPanel) Run(c task.CancelSignal) {
	flags := service.TimingGPU |
		service.TimingPerFrame
	timings, err := t.context.rpc.LoadTiming(t.device, t.capture, flags)
	if err == nil {
		c.Check()
		μsMin, μsMax := math.MaxInt, math.MinInt
		for _, timer := range timings.PerFrame {
			μs := int(timer.Nanoseconds / 1e3)
			μsMin = math.Min(μsMin, μs)
			μsMax = math.Max(μsMax, μs)
		}
		data := TimingData{
			timings: timings,
			μsMin:   μsMin,
			μsMax:   μsMax,
		}
		c.Check()
		t.context.Run(func() { t.chart.SetData(data) })
	}
}

type TimingData struct {
	timings      service.TimingInfo
	μsMin, μsMax int
}

func (d TimingData) Count() int {
	return len(d.timings.PerFrame)
}

func (d TimingData) Values(i int) []int {
	μs := int(d.timings.PerFrame[i].Nanoseconds / 1e3)
	return []int{μs}
}

func (d TimingData) Limits() (int, int) {
	rng := d.μsMax - d.μsMin
	return d.μsMin - rng/10, d.μsMax + rng/10
}

func (d TimingData) BarBrush(bar int, stack int, highlighted bool) gxui.Brush {
	if highlighted {
		return gxui.CreateBrush(d.BarBrush(bar, stack, false).Color.MulRGB(1.5).Saturate())
	}
	return gxui.CreateBrush(μs(d.Values(bar)[0]).ThresholdColor(targetFrameTime))
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
