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
	"time"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary/schema"
	"android.googlesource.com/platform/tools/gpu/client/gapis"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service/path"
	"github.com/google/gxui"
	"github.com/google/gxui/gxfont"
)

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

func CreateApplicationContext(theme gxui.Theme, config Config) (*ApplicationContext, error) {
	dropDownOverlay := theme.CreateBubbleOverlay()
	toolTipOverlay := theme.CreateBubbleOverlay()

	logger := &log.Splitter{}

	monospace, _ := theme.Driver().CreateFont(gxfont.Monospace, 12)

	appCtx := &ApplicationContext{
		Config:            config,
		theme:             theme,
		monospace:         monospace,
		logger:            logger,
		dropDownOverlay:   dropDownOverlay,
		toolTipOverlay:    toolTipOverlay,
		toolTipController: gxui.CreateToolTipController(toolTipOverlay, theme.Driver()),
		constants:         map[string]schema.ConstantSet{},
	}

	ready := make(chan struct{})
	appCtx.rpc = rpc{logger: logger, ready: ready}
	appCtx.events.Init()

	go func() {
		client, schema, err := gapis.Connect(config.Gapis, logger)
		if err != nil {
			log.E(logger, "Failed to connect to GAPIS: %v", err)
			return
		}

		for _, s := range schema.Constants {
			appCtx.constants[s.Type.String()] = s
		}

		appCtx.rpc.client = client
		close(ready)
	}()

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
