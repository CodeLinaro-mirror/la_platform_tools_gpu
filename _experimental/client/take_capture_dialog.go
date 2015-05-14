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

	"android.googlesource.com/platform/tools/gpu/binary/vle"
	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

var (
	spy = flag.String("i", "localhost:9286", "gapii TCP host:port to connect to")
)

func CreateTakeCaptureDialog(appCtx *ApplicationContext) {
	theme := appCtx.Theme()
	window := theme.CreateWindow(500, 200, "Take capture")

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

	bottom := theme.CreateLinearLayout()
	bottom.SetDirection(gxui.RightToLeft)
	bottom.AddChild(button)

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
				theme.Driver().Call(func() {
					status.SetText(update.msg)
				})
				if update.data != nil {
					data = update.data
					break
				}
			}

			if data != nil {
				theme.Driver().Call(func() {
					status.SetText("Importing...")
				})
				id, err := appCtx.Rpc().Import(appCtx.Logger(), name.Text(), data)
				if err != nil {
					panic(err)
				}
				appCtx.LoadCapture(id, true)
			}

			theme.Driver().Call(func() {
				window.Close()
			})
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

	updates <- tcUpdate{msg: fmt.Sprintf("Waiting for connection to %s...", *spy)}

	for {
		conn, err = net.Dial("tcp", *spy)
		if err == nil {
			break
		}
		if stop.signaled() {
			return
		}
		time.Sleep(500 * time.Millisecond)
	}
	defer conn.Close()

	updates <- tcUpdate{msg: "Capturing..."}

	buf := &bytes.Buffer{}
	bytesWritten := int64(0)
	for {
		if n, err := io.CopyN(buf, conn, 1024*32); err == nil {
			bytesWritten += n
			updates <- tcUpdate{msg: fmt.Sprintf("Capturing...\n%v bytes", bytesWritten)}
		} else {
			break
		}
	}
	vle.Writer(buf).Uint16(0xffff) // EOS

	updates <- tcUpdate{msg: "Done", data: buf.Bytes()}
	close(updates)
}
