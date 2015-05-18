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
	"android.googlesource.com/platform/tools/gpu/binary/vle"
	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

var (
	spyport = flag.Int("i", 9286, "gapii TCP port to connect to")
)

func CreateLaunchAndroidDialog(theme gxui.Theme, status func(s string), launched func()) {
	window := theme.CreateWindow(500, 800, "Launch Android application...")

	overlay := theme.CreateBubbleOverlay()

	deviceAdapter := gxui.CreateDefaultAdapter()
	packageAdapter := gxui.CreateDefaultAdapter()

	devices, _ := adb.Devices()
	deviceAdapter.SetItems(devices)

	deviceList := theme.CreateDropDownList()
	deviceList.SetAdapter(deviceAdapter)
	deviceList.SetBubbleOverlay(overlay)

	packageList := theme.CreateList()
	packageList.SetAdapter(packageAdapter)

	deviceList.OnSelectionChanged(func(sel gxui.AdapterItem) {
		device := sel.(*adb.Device)
		pkgs, _ := device.InstalledPackages()
		packageAdapter.SetItems(pkgs)
		packageAdapter.SetSize(math.Size{W: math.MaxSize.W, H: 16})
	})

	packageList.OnDoubleClick(func(gxui.MouseEvent) {
		if sel := packageList.Selected(); sel != nil {
			pkg := sel.(*adb.InstalledPackage)

			actions, _ := pkg.Actions()
			for _, action := range actions {
				for _, category := range action.Categories {
					if category == "android.intent.category.LAUNCHER" {
						dev := pkg.Device
						go func() {
							status("Disabling SELinux enforcing...")
							dev.SetSELinuxEnforcing(false)

							status("Setting LD_PRELOAD...")
							pkg.SetWrapProperties("LD_PRELOAD=/data/spy.so")

							status("Forwarding port...")
							dev.Forward(adb.TCPPort(*spyport), adb.NamedAbstractSocket("gfxspy"))

							status("Starting activity...")
							dev.StartActivity(action)

							launched()
						}()
						window.Close()
						return
					}
				}
			}
		}
	})

	layout := theme.CreateLinearLayout()
	layout.AddChild(deviceList)
	layout.AddChild(packageList)
	layout.AddChild(overlay)

	window.AddChild(layout)

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

	status := theme.CreateLabel()
	top.AddChild(status)

	updateStatus := func(s string) {
		theme.Driver().Call(func() {
			status.SetText(s)
		})
	}

	launch.OnClick(func(ev gxui.MouseEvent) {
		CreateLaunchAndroidDialog(theme, updateStatus, func() { button.Click(ev) })
	})

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

	clickSubscription = button.OnClick(func(gxui.MouseEvent) {
		clickSubscription.Unlisten()
		stop := make(signal)
		button.SetText("Stop")
		button.OnClick(func(gxui.MouseEvent) { stop.raise() })

		updates := make(chan tcUpdate, 8)
		go takeCapture(appCtx, stop, updates)
		go func() {
			var data []byte
			for update := range updates {
				updateStatus(update.msg)
				if update.data != nil {
					data = update.data
					break
				}
			}

			if data != nil {
				updateStatus("Importing...")
				id, err := appCtx.Rpc().Import(appCtx.Logger(), name.Text(), data)
				if err != nil {
					panic(err)
				}

				updateStatus("Loading...")
				appCtx.LoadCapture(id, true)

				theme.Driver().Call(func() {
					window.Close()
				})
			}
		}()
	})
}

type signal chan struct{}

func (s signal) raise() { close(s) }
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

func takeCapture(appCtx *ApplicationContext, stop signal, updates chan tcUpdate) {
	var conn net.Conn
	var err error

	updates <- tcUpdate{msg: fmt.Sprintf("Waiting for connection to localhost:%d...", *spyport)}

waiting:
	for {
		time.Sleep(500 * time.Millisecond)
		conn, err = net.Dial("tcp", fmt.Sprintf("localhost:%d", *spyport))
		if err == nil {
			buf := &bytes.Buffer{}
			bytesWritten := int64(0)
			for {
				n, err := io.CopyN(buf, conn, 1024*32)
				switch err {
				case nil:
					bytesWritten += n
					updates <- tcUpdate{msg: fmt.Sprintf("Capturing...\n%v bytes", bytesWritten)}

				case io.EOF:
					if len(buf.Bytes()) == 0 {
						// ADB has an annoying tendancy to insta-close forwarded sockets when
						// there's no application waiting for the connection. Treat this as
						// another waiting-for-connection case.
						continue waiting
					}

					vle.Writer(buf).Uint16(0xffff) // EOS
					updates <- tcUpdate{msg: "Done", data: buf.Bytes()}
					close(updates)
					return

				default:
					updates <- tcUpdate{msg: fmt.Sprintf("Connection error: %v", err)}
					close(updates)
					return
				}
			}
		}
		if stop.signaled() {
			return
		}
	}
	defer conn.Close()

}
