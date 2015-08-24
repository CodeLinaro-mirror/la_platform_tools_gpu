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
	"os"
	"path/filepath"
	"time"

	"android.googlesource.com/platform/tools/gpu/atexit"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service"
	"android.googlesource.com/platform/tools/gpu/service/path"
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/math"
	"github.com/google/gxui/themes/dark"
)

type Config struct {
	DataPath       string
	Gapis          string
	GXUIDebug      bool
	InitialCapture string
	ReplayDevice   string
}

type app struct {
	Config
}

func createPanels(appCtx *ApplicationContext, window gxui.Window) gxui.Control {
	theme, driver := appCtx.theme, appCtx.theme.Driver()

	vSplitter := theme.CreateSplitterLayout()
	vSplitter.SetOrientation(gxui.Vertical)

	{
		holder := theme.CreatePanelHolder()
		holder.AddPanel(CreateFramesPanel(appCtx), "Frames")
		holder.AddPanel(CreateProfilerPanel(appCtx), "Profile")
		vSplitter.AddChild(holder)
	}

	{
		hSplitter := theme.CreateSplitterLayout()
		hSplitter.SetOrientation(gxui.Horizontal)
		{
			holder := theme.CreatePanelHolder()
			holder.AddPanel(CreateCommandsPanel(appCtx), "Commands")
			hSplitter.AddChild(holder)
		}
		{
			holder := theme.CreatePanelHolder()
			holder.AddPanel(CreateColorBufferPanel(appCtx), "Color")
			holder.AddPanel(CreateDepthBufferPanel(appCtx), "Depth")
			holder.AddPanel(theme.CreateLinearLayout(), "Stencil")
			hSplitter.AddChild(holder)
			hSplitter.SetChildWeight(holder, 2.0)
		}
		vSplitter.AddChild(hSplitter)
		vSplitter.SetChildWeight(hSplitter, 2.0)
	}

	{
		hSplitter := theme.CreateSplitterLayout()
		hSplitter.SetOrientation(gxui.Horizontal)
		{
			holder := theme.CreatePanelHolder()
			holder.AddPanel(CreateStatePanel(appCtx), "State")
			if appCtx.GXUIDebug {
				holder.AddPanel(CreateGxuiDebug(appCtx, window, driver), "GXUI debug")
			}
			hSplitter.AddChild(holder)
		}
		{
			holder := theme.CreatePanelHolder()
			holder.AddPanel(CreateReportPanel(appCtx), "Report")
			holder.AddPanel(CreateMemoryPanel(appCtx), "Memory")
			holder.AddPanel(CreateImageViewerPanel(appCtx), "Image")
			holder.AddPanel(CreateResourcesPanel(appCtx), "Resources")
			holder.AddPanel(CreateLogPanel(appCtx), "Log")
			hSplitter.AddChild(holder)
		}
		vSplitter.AddChild(hSplitter)
	}

	return vSplitter
}

type capture struct {
	path *path.Capture
	info *service.Capture
}

type captureAdapter struct {
	gxui.AdapterBase
	items []capture
}

func (a *captureAdapter) Count() int {
	return len(a.items)
}

func (a *captureAdapter) ItemAt(index int) gxui.AdapterItem {
	return a.items[index].path
}

func (a *captureAdapter) ItemIndex(item gxui.AdapterItem) int {
	p := item.(*path.Capture)
	for i := range a.items {
		if path.Equal(a.items[i].path, p) {
			return i
		}
	}
	return -1
}

func (a *captureAdapter) Create(theme gxui.Theme, index int) gxui.Control {
	l := theme.CreateLabel()
	l.SetText(a.items[index].info.Name)
	return l
}

func (a *captureAdapter) Size(gxui.Theme) math.Size {
	return math.Size{W: 250, H: 18}
}

func createCaptureList(appCtx *ApplicationContext) gxui.DropDownList {
	adapter := &captureAdapter{}
	list := appCtx.theme.CreateDropDownList()
	list.SetBubbleOverlay(appCtx.dropDownOverlay)
	list.SetAdapter(adapter)
	list.OnSelectionChanged(func(item gxui.AdapterItem) {
		appCtx.events.Select(item.(*path.Capture))
	})

	list.OnAttach(func() {
		go func() {
			for list.Attached() { // While the list control is visible
				if captures, err := appCtx.rpc.GetCaptures(); err == nil {
					appCtx.Run(func() {
						adapter.items = captures
						adapter.DataChanged()
					})
				}
				time.Sleep(10 * time.Second)
			}
		}()
	})

	return list
}

type device struct {
	path *path.Device
	info *service.Device
}

type deviceAdapter struct {
	gxui.AdapterBase
	items []device
}

func (a *deviceAdapter) Count() int {
	return len(a.items)
}

func (a *deviceAdapter) ItemAt(index int) gxui.AdapterItem {
	return a.items[index].path
}

func (a *deviceAdapter) ItemIndex(item gxui.AdapterItem) int {
	p := item.(*path.Device)
	for i := range a.items {
		if path.Equal(a.items[i].path, p) {
			return i
		}
	}
	return -1
}

func (a *deviceAdapter) Create(theme gxui.Theme, index int) gxui.Control {
	l := theme.CreateLabel()
	l.SetText(a.items[index].info.Name)
	return l
}

func (a *deviceAdapter) Size(gxui.Theme) math.Size {
	return math.Size{W: 250, H: 18}
}

func (d *device) String() string { return d.info.Name }

func createDeviceList(appCtx *ApplicationContext) gxui.DropDownList {
	adapter := &deviceAdapter{}
	list := appCtx.theme.CreateDropDownList()
	list.SetBubbleOverlay(appCtx.dropDownOverlay)
	list.SetAdapter(adapter)
	list.OnSelectionChanged(func(item gxui.AdapterItem) {
		appCtx.events.Select(item.(*path.Device))
	})

	list.OnAttach(func() {
		go func() {
			for list.Attached() { // While the list control is visible
				if devices, err := appCtx.rpc.GetDevices(); err == nil {
					appCtx.Run(func() {
						adapter.items = devices
						adapter.DataChanged()
						if list.Selected() == nil && len(devices) > 0 {
							list.Select(devices[0].path)
						}
					})
				}

				time.Sleep(10 * time.Second)
			}
		}()
	})

	return list
}

func createToolbar(appCtx *ApplicationContext) gxui.Control {
	theme := appCtx.theme

	takeCapture := theme.CreateButton()
	takeCapture.SetText("Take capture")
	takeCapture.OnClick(func(gxui.MouseEvent) {
		CreateTakeCaptureDialog(appCtx)
	})

	captureLabel := theme.CreateLabel()
	captureLabel.SetText("Capture: ")
	captureList := createCaptureList(appCtx)

	deviceLabel := theme.CreateLabel()
	deviceLabel.SetText("Device: ")
	deviceList := createDeviceList(appCtx)

	layout := theme.CreateLinearLayout()
	layout.SetDirection(gxui.LeftToRight)
	layout.AddChild(takeCapture)
	layout.AddChild(captureLabel)
	layout.AddChild(captureList)
	layout.AddChild(deviceLabel)
	layout.AddChild(deviceList)
	return layout
}

func createStatusBar(appCtx *ApplicationContext) gxui.Control {
	theme := appCtx.theme

	label := theme.CreateLabel()

	appCtx.events.OnSelect(func(p path.Path) {
		label.SetText(fmt.Sprintf("Selected: %v", p.Path()))
	})

	return label
}

type EnableDebugger interface {
	EnableDebug(bool)
}

func (a app) main(driver gxui.Driver) {
	if a.GXUIDebug {
		if d, ok := driver.(EnableDebugger); ok {
			d.EnableDebug(true)
		}
	}

	theme := dark.CreateTheme(driver)
	appCtx, err := CreateApplicationContext(theme, a.Config)
	if err != nil {
		fmt.Printf("Could not create application context: %v\n", err)
		os.Exit(1)
	}

	// Create the log file
	logPath, _ := filepath.Abs(filepath.Join(a.DataPath, "..", "logs", "client.log"))
	fmt.Printf("Client log file created at: %s\n", logPath)
	logFile, err := log.File(logPath)
	if err != nil {
		panic(err)
	}
	atexit.Register(func() { log.Close(logFile) }, time.Second)
	appCtx.logger.Add(logFile)

	window := theme.CreateWindow(800, 600, "Main")

	layout := theme.CreateLinearLayout()
	layout.SetDirection(gxui.TopToBottom)
	layout.AddChild(createToolbar(appCtx))
	layout.AddChild(createPanels(appCtx, window))

	status := theme.CreateLinearLayout()
	status.SetDirection(gxui.BottomToTop)
	status.AddChild(createStatusBar(appCtx))
	status.AddChild(layout)

	window.OnClose(driver.Terminate)
	window.AddChild(status)
	window.AddChild(appCtx.dropDownOverlay)
	window.AddChild(appCtx.toolTipOverlay)

	if appCtx.InitialCapture != "" {
		ImportCapture(appCtx, appCtx.InitialCapture, appCtx.logger)
	}

	appCtx.events.OnSelect(func(p path.Path) {
		log.Infof(appCtx.logger, "Select '%v'", p)
	})
}

func Run(config Config) {
	app := app{Config: config}
	gl.StartDriver(app.main)
}
