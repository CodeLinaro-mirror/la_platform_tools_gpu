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
	"net"
	"os"
	"os/exec"
	"time"

	"android.googlesource.com/platform/tools/gpu/atexit"
	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary/schema"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/multiplexer"
	"android.googlesource.com/platform/tools/gpu/service"
	"android.googlesource.com/platform/tools/gpu/service/path"
	"github.com/google/gxui"
	"github.com/google/gxui/gxfont"
)

const mtu = 1024

var (
	rpcResourceRetryDelay = time.Millisecond * 250
)

type ApplicationContext struct {
	Config
	theme             gxui.Theme
	monospace         gxui.Font
	logger            *log.Splitter
	rpc               rpc
	dropDownOverlay   gxui.BubbleOverlay
	toolTipOverlay    gxui.BubbleOverlay
	toolTipController *gxui.ToolTipController
	events            Events
	atoms             []atom.Atom
	device            *path.Device
	constants         map[string]schema.ConstantSet
}

func connectServer(config Config) (net.Conn, error) {
	conn, err := net.Dial("tcp", config.Gapis)
	if err == nil {
		return conn, nil
	}
	// connection failed, was it localhost?
	host, _, err2 := net.SplitHostPort(config.Gapis)
	if host != "localhost" {
		return nil, err
	}
	if err2 != nil {
		return nil, err2
	}
	// try to run the server ourselves
	null, err := os.OpenFile(os.DevNull, os.O_RDWR, 0)
	if err != nil {
		return nil, err
	}
	path, err := exec.LookPath("gapis")
	if err != nil {
		return nil, err
	}
	args := []string{path,
		"--rpc", config.Gapis,
		"--data", config.DataPath,
	}
	proc, err := os.StartProcess(path, args, &os.ProcAttr{Files: []*os.File{null, null, null}})
	if err != nil {
		return nil, err
	}
	// We are running a server, shut it down when we are done
	atexit.Register(func() {
		proc.Kill()
		proc.Wait()
	}, time.Second)
	// and try to connect to it (for a while)
	for i := 0; i < 30; i++ {
		conn, err = net.Dial("tcp", config.Gapis)
		if err == nil {
			return conn, nil
		}
		time.Sleep(100 * time.Millisecond)
	}
	return nil, fmt.Errorf("Failed to spawn %v in time", path)
}

func CreateApplicationContext(theme gxui.Theme, config Config) (*ApplicationContext, error) {
	dropDownOverlay := theme.CreateBubbleOverlay()
	toolTipOverlay := theme.CreateBubbleOverlay()

	rpcSocket, err := connectServer(config)
	if err != nil {
		return nil, err
	}

	monospace, _ := theme.Driver().CreateFont(gxfont.Monospace, 12)

	appCtx := &ApplicationContext{
		Config:            config,
		theme:             theme,
		monospace:         monospace,
		logger:            &log.Splitter{},
		dropDownOverlay:   dropDownOverlay,
		toolTipOverlay:    toolTipOverlay,
		toolTipController: gxui.CreateToolTipController(toolTipOverlay, theme.Driver()),
		constants:         map[string]schema.ConstantSet{},
	}

	client := service.NewClient(multiplexer.New(rpcSocket, rpcSocket, mtu, nil), nil)
	appCtx.rpc.init(log.Enter(appCtx.logger, "rpc"), client, appCtx.constants)
	appCtx.events.Init()
	return appCtx, nil
}

// Run enqueues f to be called on the UI go-routine.
// Run can return before f is called.
func (c *ApplicationContext) Run(f func()) bool {
	return c.theme.Driver().Call(f)
}

// RunSync calls f on the UI go-routine, blocking until f has returned.
func (c *ApplicationContext) RunSync(f func()) bool {
	return c.theme.Driver().CallSync(f)
}

// Change modifies the value at p to v, and selects the new path.
// The call is blocking.
func (c *ApplicationContext) Change(p path.Path, v interface{}) error {
	n, err := c.rpc.Change(p, v)
	if err != nil {
		return err
	}

	if p != n {
		c.events.Select(n)
	}

	return nil
}
