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
	"image"
	"image/color"
	"os"
	"path/filepath"
	"time"

	"android.googlesource.com/platform/tools/gpu/atexit"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service"
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
	theme, driver := appCtx.Theme(), appCtx.Theme().Driver()

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
			holder.AddPanel(CreateDocsPanel(appCtx), "Docs")
			holder.AddPanel(CreateLogPanel(appCtx), "Log")
			hSplitter.AddChild(holder)
		}
		vSplitter.AddChild(hSplitter)
	}

	return vSplitter
}

type capture struct {
	id   service.CaptureId
	info service.Capture
}

func (d *capture) String() string { return d.info.Name }

func createCaptureList(appCtx *ApplicationContext) gxui.DropDownList {
	theme, r := appCtx.Theme(), appCtx.Rpc()

	adapter := gxui.CreateDefaultAdapter()

	list := theme.CreateDropDownList()
	list.SetBubbleOverlay(appCtx.DropDownOverlay())
	list.SetAdapter(adapter)
	list.OnSelectionChanged(func(item gxui.AdapterItem) {
		appCtx.LoadCapture(item.(*capture).id, true)
	})

	go func() {
		l := appCtx.Logger().Fork().Enter("CreateCaptureList")
		ids, err := r.GetCaptures(l)
		if err != nil {
			return
		}
		captures := make([]*capture, 0, len(ids))
		for _, id := range ids {
			id := id
			go func() {
				if info, err := r.ResolveCapture(id, l); err == nil {
					appCtx.Run(func() {
						captures = append(captures, &capture{id, info})
						adapter.SetItems(captures)
					})
				}
			}()
		}
	}()
	return list
}

type device struct {
	id   service.DeviceId
	info service.Device
}

func (d *device) String() string { return d.info.Name }

func createDeviceList(appCtx *ApplicationContext) gxui.DropDownList {
	theme, r := appCtx.Theme(), appCtx.Rpc()
	driver := theme.Driver()

	adapter := gxui.CreateDefaultAdapter()
	wanted := appCtx.ReplayDevice

	list := theme.CreateDropDownList()
	list.SetBubbleOverlay(appCtx.DropDownOverlay())
	list.SetAdapter(adapter)
	list.OnSelectionChanged(func(item gxui.AdapterItem) {
		appCtx.SelectDevice(item.(*device).id)
		wanted = ""
	})

	list.OnAttach(func() {
		go func() {
			l := appCtx.Logger().Fork().Enter("DeviceListUpdate")

			for list.Attached() { // While the list control is visible
				if ids, err := r.GetDevices(l); err == nil {
					devices := make([]*device, 0, len(ids))
					found := -1
					for _, id := range ids {
						if info, err := r.ResolveDevice(id, l); err == nil {
							devices = append(devices, &device{id, info})
							if info.Name == wanted {
								found = len(devices) - 1
								wanted = ""
							}
						}
					}

					appCtx.Run(func() {
						adapter.SetItems(devices)
						if found >= 0 {
							list.Select(devices[found])
						}
					})
				}

				var selected service.DeviceId
				driver.CallSync(func() { selected = appCtx.SelectedDevice() })

				if !selected.Valid() {
					time.Sleep(250 * time.Millisecond)
				} else {
					time.Sleep(30 * time.Second)
				}
			}
		}()
	})

	return list
}

func createToolbar(appCtx *ApplicationContext) gxui.Control {
	theme := appCtx.Theme()

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

func loadTiming(appCtx *ApplicationContext) {
	deviceID := appCtx.SelectedDevice()
	captureID := appCtx.CaptureID()

	go func() {
		l := appCtx.Logger().Fork().Enter("Replay: timing")
		appCtx.timingInfo = service.TimingInfo{}

		mask := service.TimingMaskTimingPerFrame |
			service.TimingMaskTimingPerDrawCall |
			service.TimingMaskTimingPerCommand
		timingInfoID, err := appCtx.rpc.GetTimingInfo(deviceID, captureID, mask, l)
		if err != nil {
			return
		}

		timingInfo, err := appCtx.rpc.ResolveTimingInfo(timingInfoID, l)
		if err != nil {
			return
		}

		timingPerCommand := make(map[uint64]uint64)
		for _, t := range timingInfo.PerCommand {
			timingPerCommand[t.AtomId] = t.Nanoseconds
		}

		appCtx.Run(func() {
			appCtx.timingInfo = timingInfo
			appCtx.timingPerCommand = timingPerCommand
			appCtx.onTimingInfoUpdated.Fire()
		})
	}()
}

func DoReplay(appCtx *ApplicationContext) {
	driver, r := appCtx.Theme().Driver(), appCtx.Rpc()
	deviceID := appCtx.SelectedDevice()
	captureID := appCtx.CaptureID()
	atomID := appCtx.SelectedAtomID()
	settings := service.RenderSettings{
		MaxWidth:  65536,
		MaxHeight: 65536,
		Wireframe: appCtx.Wireframe(),
	}

	if atomID == InvalidAtomID {
		return
	}

	atom := appCtx.Atoms()[atomID]
	apiID := service.ApiId{ID: binary.ID(atom.API())}

	go func() {
		l := appCtx.Logger().Fork().Enter("Replay: color-buffer")
		imageID, err := r.GetFramebufferColor(deviceID, captureID, apiID, uint64(atomID), settings, l)
		if err != nil {
			return
		}

		imageInfo, err := r.ResolveImageInfo(imageID, l)
		if err != nil {
			return
		}

		imageData, err := r.ResolveBinary(imageInfo.Data, l)
		if err != nil {
			return
		}

		width, height := int(imageInfo.Width), int(imageInfo.Height)

		var img *image.RGBA
		if width > 0 && height > 0 {
			img = image.NewRGBA(image.Rect(0, 0, width, height))
			img.Pix = []byte(imageData.Data)
		}

		appCtx.Run(func() {
			if appCtx.colorBuffer != nil {
				appCtx.colorBuffer.Release()
			}
			appCtx.colorBuffer = nil
			if img != nil {
				appCtx.colorBuffer = driver.CreateTexture(img, 1)
				appCtx.colorBuffer.SetFlipY(true)
			}
			appCtx.onColorBufferUpdate.Fire()
		})

	}()

	go func() {
		l := appCtx.Logger().Fork().Enter("Replay: depth-buffer")
		imageID, err := r.GetFramebufferDepth(deviceID, captureID, apiID, uint64(atomID), l)
		if err != nil {
			return
		}

		imageInfo, err := r.ResolveImageInfo(imageID, l)
		if err != nil {
			return
		}

		imageData, err := r.ResolveBinary(imageInfo.Data, l)
		if err != nil {
			return
		}

		buffer := imageData.Data

		width, height := int(imageInfo.Width), int(imageInfo.Height)

		var img *image.RGBA
		if width > 0 && height > 0 {
			img = image.NewRGBA(image.Rect(0, 0, width, height))
			for y := 0; y < height; y++ {
				for x := 0; x < width; x++ {
					r, g, b, a := float32(buffer[0]), float32(buffer[1])/255.0, float32(buffer[2])/65025.0, float32(buffer[3])/160581375.0
					depth := (r + g + b + a) / 255.0
					buffer = buffer[4:]
					d := 0.01 / (1.0 - depth)
					c := color.RGBA{
						R: byte(math.Cosf(d+math.TwoPi*0.000)*127.0 + 128.0),
						G: byte(math.Cosf(d+math.TwoPi*0.333)*127.0 + 128.0),
						B: byte(math.Cosf(d+math.TwoPi*0.666)*127.0 + 128.0),
						A: byte(0xFF),
					}
					img.Set(x, y, c)
				}
			}
		}

		appCtx.Run(func() {
			if appCtx.depthBuffer != nil {
				appCtx.depthBuffer.Release()
			}
			appCtx.depthBuffer = nil
			if img != nil {
				appCtx.depthBuffer = driver.CreateTexture(img, 1)
				appCtx.depthBuffer.SetFlipY(true)
			}
			appCtx.onDepthBufferUpdate.Fire()
		})
	}()
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
	atexit.Register(logFile.Close, time.Second)
	appCtx.Logger().Add(logFile)

	window := theme.CreateWindow(800, 600, "Main")

	layout := theme.CreateLinearLayout()
	layout.SetDirection(gxui.TopToBottom)
	layout.AddChild(createToolbar(appCtx))
	layout.AddChild(createPanels(appCtx, window))

	window.OnClose(driver.Terminate)
	window.AddChild(layout)
	window.AddChild(appCtx.DropDownOverlay())
	window.AddChild(appCtx.ToolTipOverlay())

	// Perhaps add button to disable this?
	appCtx.OnAtomSelected(appCtx.RequestReplay)
	appCtx.OnAtomsUpdated(appCtx.LoadHierarchy)
	appCtx.OnAtomsUpdated(appCtx.LoadReport)
	appCtx.OnAtomsUpdated(func() { loadTiming(appCtx) })
	appCtx.OnRequestReplay(func() { DoReplay(appCtx) })
	appCtx.OnAtomSelected(appCtx.LoadState)

	appCtx.UpdateSchema()

	if appCtx.InitialCapture != "" {
		ImportCapture(appCtx, appCtx.InitialCapture, appCtx.logger)
	}
}

func Run(config Config) {
	app := app{Config: config}
	gl.StartDriver(app.main)
}
