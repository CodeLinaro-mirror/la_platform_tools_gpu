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
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os/user"
	"path/filepath"
	"time"

	"android.googlesource.com/platform/tools/gpu/adb"
	"android.googlesource.com/platform/tools/gpu/gapii"
	"android.googlesource.com/platform/tools/gpu/log"
	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

var (
	spyport = flag.Int("i", 9286, "gapii TCP port to connect to")
)

type launchItem struct {
	*adb.Action
}

func (i launchItem) String() string { return fmt.Sprintf("%s/%s", i.Package.Name, i.Activity) }

func CreateLaunchAndroidDialog(theme gxui.Theme, statusLogger log.Logger, capture func()) {
	driver := theme.Driver()
	window := theme.CreateWindow(500, 800, "Launch Android application...")

	overlay := theme.CreateBubbleOverlay()

	deviceAdapter := gxui.CreateDefaultAdapter()

	devices, _ := adb.Devices()
	deviceAdapter.SetItems(devices)

	deviceList := theme.CreateDropDownList()
	deviceList.SetAdapter(deviceAdapter)
	deviceList.SetBubbleOverlay(overlay)

	packageList := theme.CreateList()

	deviceList.OnSelectionChanged(func(sel gxui.AdapterItem) {
		device := sel.(*adb.Device)
		adapter := gxui.CreateDefaultAdapter()
		adapter.SetSize(math.Size{W: math.MaxSize.W, H: 16})
		packageList.SetAdapter(adapter)
		go func() (err error) {
			defer func() {
				if err != nil {
					driver.Call(func() { adapter.SetItems(err) })
				}
			}()

			items := []launchItem{}
			err = device.Root()
			switch err {
			case nil:
			case adb.ErrDeviceNotRooted:
				return err
			default:
				return fmt.Errorf("Failed to restart ADB as root: %v", err)
			}
			packages, err := device.InstalledPackages()
			if len(packages) == 0 || err != nil {
				return fmt.Errorf(fmt.Sprintf("Could not get list of installed packages: %v", err))
			}
			for _, pkg := range packages {
				for _, action := range pkg.Actions {
					if action.Name == "android.intent.action.MAIN" {
						driver.CallSync(func() {
							items = append(items, launchItem{action})
							adapter.SetItems(items)
						})
					}
				}
			}
			return nil
		}()
	})

	packageList.OnDoubleClick(func(gxui.MouseEvent) {
		if sel := packageList.Selected(); sel != nil {
			if item, ok := sel.(launchItem); ok {
				go func() {
					driver.Call(window.Close)
					gapii.AdbStart(statusLogger, item.Action, adb.TCPPort(*spyport))

					capture()
				}()
			}
		}
	})

	layout := theme.CreateLinearLayout()
	layout.AddChild(deviceList)
	layout.AddChild(packageList)

	window.AddChild(layout)
	window.AddChild(overlay)

	if len(devices) > 0 {
		deviceList.Select(devices[0])
	}
}

func CreateTakeCaptureDialog(appCtx *ApplicationContext) {
	theme := appCtx.Theme()
	window := theme.CreateWindow(500, 200, "Take capture")

	launch := theme.CreateButton()
	launch.SetText("Launch...")

	button := theme.CreateButton()
	button.SetText("Capture...")

	load := theme.CreateButton()
	load.SetText("Import...")

	namelbl := theme.CreateLabel()
	namelbl.SetText("Capture name:")

	name := theme.CreateTextBox()
	name.SetDesiredWidth(math.MaxSize.W)
	name.SetText(time.Now().String())

	top := theme.CreateLinearLayout()
	top.SetSizeMode(gxui.Fill)
	top.SetDirection(gxui.TopToBottom)

	row := theme.CreateLinearLayout()
	row.SetDirection(gxui.LeftToRight)
	row.AddChild(namelbl)
	row.AddChild(name)
	top.AddChild(row)

	statusAdapter := CreateLogAdapter(1024, appCtx.Theme().Driver().Call)
	statusLogger := statusAdapter.Logger()
	status := theme.CreateList()
	status.SetAdapter(statusAdapter)
	top.AddChild(status)

	bottom := theme.CreateLinearLayout()
	bottom.SetDirection(gxui.RightToLeft)
	bottom.AddChild(button)
	bottom.AddChild(launch)
	bottom.AddChild(load)

	layout := theme.CreateLinearLayout()
	layout.SetDirection(gxui.BottomToTop)
	layout.AddChild(bottom)
	layout.AddChild(top)

	window.AddChild(layout)

	var clickSubscription gxui.EventSubscription
	capture := func() {
		stop := make(chan struct{})

		clickSubscription.Unlisten()
		button.SetText("Stop")
		clickSubscription = button.OnClick(func(gxui.MouseEvent) {
			clickSubscription.Unlisten()
			close(stop)
		})

		go func() {
			buf := &bytes.Buffer{}
			count, _ := gapii.Capture(statusLogger, *spyport, buf, stop)
			if count > 0 {
				data := buf.Bytes()
				statusLogger.Infof("Importing...")
				id, err := appCtx.Rpc().Import(name.Text(), data, appCtx.Logger())
				if err != nil {
					panic(err)
				}

				statusLogger.Infof("Loading...")
				appCtx.LoadCapture(id, true)

				theme.Driver().Call(func() {
					window.Close()
				})
			} else {
				theme.Driver().Call(func() {
					button.SetText("Close")
					button.OnClick(func(gxui.MouseEvent) {
						window.Close()
					})
				})
			}
		}()
	}

	clickSubscription = button.OnClick(func(gxui.MouseEvent) { capture() })
	launch.OnClick(func(ev gxui.MouseEvent) {
		CreateLaunchAndroidDialog(theme, statusLogger, capture)
	})

	load.OnClick(func(ev gxui.MouseEvent) {
		go func() {
			path := name.Text()
			if path[:2] == "~/" {
				usr, _ := user.Current()
				path = filepath.Join(usr.HomeDir, path[2:])
			}
			statusLogger.Infof("Loading %s", path)
			data, err := ioutil.ReadFile(path)
			if err != nil {
				statusLogger.Infof("Failed opening file %s: %s", path, err)
			} else if len(data) == 0 {
				statusLogger.Infof("Zero size file %s", path)
			} else {
				statusLogger.Infof("Importing...")
				id, err := appCtx.Rpc().Import(name.Text(), data, appCtx.Logger())
				if err != nil {
					panic(err)
				}
				statusLogger.Infof("Loading...")
				appCtx.LoadCapture(id, true)
				theme.Driver().Call(func() {
					window.Close()
				})
			}
		}()
	})
}
