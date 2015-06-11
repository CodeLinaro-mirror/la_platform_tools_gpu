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
	"io"
	"net"
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

	layout := theme.CreateLinearLayout()
	layout.SetDirection(gxui.BottomToTop)
	layout.AddChild(bottom)
	layout.AddChild(top)

	window.AddChild(layout)

	var clickSubscription gxui.EventSubscription
	capture := func() {
		stop := make(signal)

		clickSubscription.Unlisten()
		button.SetText("Stop")
		clickSubscription = button.OnClick(func(gxui.MouseEvent) {
			clickSubscription.Unlisten()
			stop.raise()
		})

		go func() {
			if data := takeCapture(appCtx, stop, statusLogger); data != nil {
				statusLogger.Infof("Importing...")
				id, err := appCtx.Rpc().Import(appCtx.Logger(), name.Text(), data)
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
}

type signal chan struct{}

func (s signal) raise() {
	close(s)
}
func (s signal) signaled() bool {
	select {
	case <-s:
		return true
	default:
		return false
	}
}

type tcUpdate struct {
	msg  string
	data []byte
}

func takeCapture(appCtx *ApplicationContext, stop signal, statusLogger log.Logger) []byte {
	var conn net.Conn
	var err error

	statusLogger.Infof("Waiting for connection to localhost:%d...", *spyport)

waiting:
	for {
		if stop.signaled() {
			return nil
		}
		time.Sleep(500 * time.Millisecond)
		conn, err = net.Dial("tcp", fmt.Sprintf("localhost:%d", *spyport))
		if err == nil {
			buf := &bytes.Buffer{}
			bytesWritten := int64(0)
			for {
				if stop.signaled() {
					conn.Close()
					return buf.Bytes()
				}

				conn.SetReadDeadline(time.Now().Add(time.Millisecond * 100)) // Allow for stop event and UI refreshes.
				n, err := io.CopyN(buf, conn, 1024*64)

				if err, neterr := err.(net.Error); neterr {
					if err.Temporary() || err.Timeout() {
						bytesWritten += n
						statusLogger.Infof("Capturing...\n%v bytes", bytesWritten)
						continue
					}
				}

				switch err {
				case nil:
					bytesWritten += n
					statusLogger.Infof("Capturing...\n%v bytes", bytesWritten)

				case io.EOF:
					if len(buf.Bytes()) == 0 {
						// ADB has an annoying tendancy to insta-close forwarded sockets when
						// there's no application waiting for the connection. Treat this as
						// another waiting-for-connection case.
						continue waiting
					}

					statusLogger.Infof("Done")
					return buf.Bytes()

				default:
					statusLogger.Infof("Connection error: %v", err)
					return buf.Bytes()
				}
			}
		}
	}

	return nil
}
