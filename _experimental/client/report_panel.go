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

	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service"
	"android.googlesource.com/platform/tools/gpu/service/path"
	"android.googlesource.com/platform/tools/gpu/task"
	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

func CreateReportPanel(appCtx *ApplicationContext) gxui.Control {
	adapter := &ReportAdapter{}
	l := appCtx.theme.CreateList()
	l.SetAdapter(adapter)

	var capture *path.Capture

	t := task.New()
	update := func() {
		if capture != nil {
			t.Run(updateReportAdapter{appCtx, capture, adapter})
		}
	}

	appCtx.events.OnSelect(func(p path.Path) {
		if c := path.FindCapture(p); p != nil && !path.Equal(c, capture) {
			capture = c
			update()
		}
	})

	return l
}

type updateReportAdapter struct {
	context *ApplicationContext
	capture *path.Capture
	adapter *ReportAdapter
}

func (t updateReportAdapter) Run(c task.CancelSignal) {
	report, err := t.context.rpc.LoadReport(t.capture)
	if err != nil {
		return
	}
	c.Check()
	t.context.Run(func() {
		t.adapter.Update(report)
	})
}

type ReportAdapter struct {
	gxui.AdapterBase
	report service.Report
}

func (a *ReportAdapter) Update(report service.Report) {
	a.report = report
	a.DataReplaced()
}

func (a *ReportAdapter) Count() int {
	return len(a.report.Items)
}

func (a *ReportAdapter) Size(theme gxui.Theme) math.Size {
	return math.Size{math.MaxSize.W, 16}
}

func (a *ReportAdapter) ItemAt(index int) gxui.AdapterItem {
	return a.report.Items[index]
}

func (a *ReportAdapter) ItemIndex(item gxui.AdapterItem) int {
	for i, x := range a.report.Items {
		if x == item {
			return i
		}
	}
	return -1
}

func (a *ReportAdapter) Create(t gxui.Theme, index int) gxui.Control {
	item := a.report.Items[index]

	label := t.CreateLabel()
	label.SetText(fmt.Sprintf("%.5d %s %s", item.Atom, item.Severity, item.Message))

	switch item.Severity {
	case log.Warning:
		label.SetColor(gxui.Yellow)
	case log.Error, log.Critical:
		label.SetColor(gxui.Red)
	}

	return label
}
