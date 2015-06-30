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
	"fmt"
	"image"
	"net"
	"os"
	"os/exec"
	"time"

	"android.googlesource.com/platform/tools/gpu/atexit"
	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary/cyclic"
	"android.googlesource.com/platform/tools/gpu/binary/registry"
	"android.googlesource.com/platform/tools/gpu/binary/schema"
	"android.googlesource.com/platform/tools/gpu/binary/vle"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/multiplexer"
	"android.googlesource.com/platform/tools/gpu/service"
	"github.com/google/gxui"
	"github.com/google/gxui/gxfont"
)

const mtu = 1024

var (
	rpcResourceRetryDelay = time.Millisecond * 250
)

type ApplicationContext struct {
	Config
	theme               gxui.Theme
	monospace           gxui.Font
	logger              *log.Splitter
	rpc                 service.Client
	captureID           service.CaptureId
	capture             service.Capture
	dropDownOverlay     gxui.BubbleOverlay
	toolTipOverlay      gxui.BubbleOverlay
	toolTipController   *gxui.ToolTipController
	onAtomSelected      gxui.Event
	onObjectSelected    gxui.Event
	onAddressSelected   gxui.Event
	onColorBufferUpdate gxui.Event
	onDepthBufferUpdate gxui.Event
	onRequestReplay     gxui.Event
	onWireframeChanged  gxui.Event
	onDeviceSelected    gxui.Event
	onAtomsUpdated      gxui.Event
	onHierarchyUpdated  gxui.Event
	onReportUpdated     gxui.Event
	onStateUpdated      gxui.Event
	onTimingInfoUpdated gxui.Event
	schema              service.Schema
	atoms               []Atom
	state               *schema.Object
	hierarchy           atom.Group
	report              service.Report
	selectedAtomID      atom.ID
	selectedAddress     memory.Pointer
	selectedObject      interface{}
	selectedDevice      service.DeviceId
	wireframe           bool
	colorBuffer         gxui.Texture
	depthBuffer         gxui.Texture
	timingInfo          service.TimingInfo
	timingPerCommand    map[uint64]uint64
	namespace           *registry.Namespace // The namespace to use in coders
	schemaNamespace     *registry.Namespace // The namespace that holds the schema classes
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
	// We are running a server,  shut it down when we are done
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
	// Start the client with the global namespace. This will be replaced when
	// the schema has been fetched from the server.
	rpc := service.NewClient(multiplexer.New(rpcSocket, rpcSocket, mtu, nil), nil)

	monospace, _ := theme.Driver().CreateFont(gxfont.Monospace, 12)

	appCtx := &ApplicationContext{
		Config:              config,
		theme:               theme,
		monospace:           monospace,
		logger:              &log.Splitter{},
		rpc:                 rpc,
		dropDownOverlay:     dropDownOverlay,
		toolTipOverlay:      toolTipOverlay,
		toolTipController:   gxui.CreateToolTipController(toolTipOverlay, theme.Driver()),
		onAtomSelected:      gxui.CreateEvent(func() {}),
		onObjectSelected:    gxui.CreateEvent(func() {}),
		onAddressSelected:   gxui.CreateEvent(func() {}),
		onColorBufferUpdate: gxui.CreateEvent(func() {}),
		onDepthBufferUpdate: gxui.CreateEvent(func() {}),
		onRequestReplay:     gxui.CreateEvent(func() {}),
		onWireframeChanged:  gxui.CreateEvent(func() {}),
		onDeviceSelected:    gxui.CreateEvent(func() {}),
		onAtomsUpdated:      gxui.CreateEvent(func() {}),
		onHierarchyUpdated:  gxui.CreateEvent(func() {}),
		onReportUpdated:     gxui.CreateEvent(func() {}),
		onStateUpdated:      gxui.CreateEvent(func() {}),
		onTimingInfoUpdated: gxui.CreateEvent(func() {}),
		selectedAtomID:      InvalidAtomID,
	}
	appCtx.schemaNamespace = registry.NewNamespace()
	// make the decoder namespace try the global namespace before the schema one
	appCtx.namespace = registry.NewNamespace(registry.Global, appCtx.schemaNamespace)
	return appCtx, nil
}

func (c *ApplicationContext) UpdateSchema() {
	go func() {
		classes, err := c.rpc.GetSchema(c.logger)
		if err != nil {
			log.Errorf(c.logger, "Error resolving schema: %v", err)
			return
		}
		log.Infof(c.logger, "Schema with %d classes", len(classes))
		for _, class := range classes {
			c.schemaNamespace.Add(class)
		}
		// Replace the current RPC
		c.rpc = service.NewClient(c.rpc.Multiplexer(), c.namespace)
	}()
}

func (c *ApplicationContext) Run(f func()) {
	c.theme.Driver().Call(f)
}

func (c *ApplicationContext) SelectAtom(id atom.ID) {
	if c.selectedAtomID != id {
		log.Infof(c.logger, "SelectAtom(%v)", id)
		c.selectedAtomID = id
		c.onAtomSelected.Fire()
	}
}

func (c *ApplicationContext) SelectAddress(address memory.Pointer) {
	if c.selectedAddress != address {
		log.Infof(c.logger, "SelectAddress(%v)", address)
		c.selectedAddress = address
		c.onAddressSelected.Fire()
	}
}

func (c *ApplicationContext) SelectObject(object interface{}) {
	if c.selectedObject != object {
		log.Infof(c.logger, "SelectObject(%v)", object)
		c.selectedObject = object
		c.onObjectSelected.Fire()
	}
}

func (c *ApplicationContext) SelectDevice(device service.DeviceId) {
	if c.selectedDevice != device {
		log.Infof(c.logger, "SelectDevice(%v)", device)
		c.selectedDevice = device
		c.onDeviceSelected.Fire()
	}
}

func (c *ApplicationContext) SetWireframe(value bool) {
	if c.wireframe != value {
		log.Infof(c.logger, "SetWireframe(%v)", value)
		c.wireframe = value
		c.onWireframeChanged.Fire()
	}
}

func (c *ApplicationContext) LoadCapture(captureID service.CaptureId, resetSelected bool) {
	if captureID == c.captureID {
		return
	}

	l := c.logger.Fork().Enter("LoadCapture")
	log.Infof(l, "(capture: %v)", captureID)

	go func() {
		var err error

		capture, err := c.rpc.ResolveCapture(captureID, l)
		if err != nil {
			log.Errorf(l, "Error resolving capture: %v", err)
			return
		}

		atoms, err := c.rpc.ResolveAtomStream(capture.Atoms, l)
		if err != nil {
			log.Errorf(l, "Error resolving capture: %v", err)
			return
		}

		s, err := c.rpc.ResolveSchema(capture.Schema, l)
		if err != nil {
			log.Errorf(l, "Error resolving capture: %v", err)
			return
		}

		c.Run(func() {
			atoms, err := c.DecodeAtoms(atoms, s)
			if err != nil {
				panic(err)
			}
			c.captureID = captureID
			c.capture = capture
			c.schema = s
			c.atoms = atoms
			if resetSelected {
				c.selectedAtomID = InvalidAtomID
				c.selectedAddress = 0
				c.selectedObject = nil
			}
			c.onAtomsUpdated.Fire()
			c.RequestReplay()
			log.Infof(l, "Capture '%s' loaded: %d atoms", c.capture.GetName(), len(atoms))
		})
	}()
}

func (c *ApplicationContext) LoadHierarchy() {
	captureID := c.captureID
	l := c.logger.Fork().Enter("LoadHierarchy")
	log.Infof(l, "(capture: %v)", captureID)

	go func() {
		hierarchy, err := c.rpc.GetHierarchy(captureID, l)
		if err != nil {
			return
		}
		root, err := c.rpc.ResolveHierarchy(hierarchy, l)
		if err != nil {
			return
		}
		c.Run(func() {
			c.hierarchy = atom.Group{}
			root.Root.Unpack(&c.hierarchy)
			c.onHierarchyUpdated.Fire()
			log.Infof(l, "Hierarchy loaded")
		})
	}()
}

func (c *ApplicationContext) LoadReport() {
	captureID := c.captureID
	l := c.logger.Fork().Enter("LoadReport")
	log.Infof(l, "(capture: %v)", captureID)

	go func() {
		report, err := c.rpc.ResolveReport(c.capture.Report, l)
		if err != nil {
			return
		}
		c.Run(func() {
			c.report = report
			c.onReportUpdated.Fire()
			log.Infof(l, "Report loaded")
		})
	}()
}

func (c *ApplicationContext) LoadState() {
	captureID := c.captureID
	after := c.selectedAtomID
	apiID := c.atoms[after].Api()
	l := c.logger.Fork().Enter("LoadState")

	go func() {
		id, err := c.rpc.GetState(captureID, apiID, uint64(after), l)
		if err != nil {
			log.E(l, "%v", err)
			return
		}
		state, err := c.rpc.ResolveState(id, l)
		if err != nil {
			log.E(l, "%v", err)
			return
		}
		c.Run(func() {
			c.state = state.(*schema.Object)
			c.onStateUpdated.Fire()
		})
	}()
}

func (c *ApplicationContext) RequestReplay() {
	c.onRequestReplay.Fire()
}

type ImageCallback func(gxui.Texture)

func isClosed(c <-chan struct{}) bool {
	select {
	case <-c:
		return true
	default:
		return false
	}
}

func (c *ApplicationContext) RequestThumbnail(after atom.ID, maxWidth, maxHeight uint32, callback ImageCallback) chan<- struct{} {
	l := c.logger.Fork().Enter("RequestThumbnail")
	log.Infof(l, "(device: %v, after: %v, max size: %dx%d)", c.selectedDevice, after, maxWidth, maxHeight)

	cancel := make(chan struct{})
	device := c.selectedDevice
	captureID := c.captureID
	settings := service.RenderSettings{
		MaxWidth:  maxWidth,
		MaxHeight: maxHeight,
		Wireframe: false,
	}

	if !device.Valid() {
		log.Warningf(l, "No device selected")
		return nil
	}

	apiID := c.atoms[after].Api()

	go func() {
		imageID, err := c.rpc.GetFramebufferColor(device, captureID, apiID, uint64(after), settings, l)
		if err != nil {
			return
		}
		if isClosed(cancel) {
			log.Infof(l, "Request cancelled")
			return
		}

		imageInfo, err := c.rpc.ResolveImageInfo(imageID, l)
		if err != nil {
			return
		}
		if isClosed(cancel) {
			log.Infof(l, "Request cancelled")
			return
		}

		log.Infof(l, "Image info resolved")
		imageData, err := c.rpc.ResolveBinary(imageInfo.Data, l)
		if err != nil {
			return
		}
		if isClosed(cancel) {
			log.Infof(l, "Request cancelled")
			return
		}

		log.Infof(l, "Image %dx%d resolved", imageInfo.Width, imageInfo.Height)
		if imageInfo.Width > 0 && imageInfo.Height > 0 {
			img := image.NewRGBA(image.Rect(0, 0, int(imageInfo.Width), int(imageInfo.Height)))
			img.Pix = []byte(imageData.Data)
			c.Run(func() {
				tex := c.theme.Driver().CreateTexture(img, 1)
				tex.SetFlipY(true)
				callback(tex)
			})
		}
	}()

	return cancel
}

type MemoryCallback func(service.MemoryInfo)

func (c *ApplicationContext) RequestMemory(after atom.ID, base memory.Pointer, size uint64, callback MemoryCallback) chan<- struct{} {
	l := c.logger.Fork().Enter("RequestMemory")
	log.Infof(l, "(after: %v, base: 0x%x, size: 0x%x)", after, base, size)

	cancel := make(chan struct{})
	captureID := c.captureID
	if c.captureID.Valid() {
		go func() {
			rng := service.MemoryRange{Base: uint64(base), Size: size}
			id, err := c.rpc.GetMemoryInfo(captureID, uint64(after), rng, l)
			if err != nil {
				return
			}
			if isClosed(cancel) {
				log.Infof(l, "Request cancelled")
				return
			}

			info, err := c.rpc.ResolveMemoryInfo(id, l)
			if err != nil {
				return
			}
			if isClosed(cancel) {
				log.Infof(l, "Request cancelled")
				return
			}

			c.Run(func() {
				callback(info)
			})
		}()
	}
	return cancel
}

func (c *ApplicationContext) ReplaceAtom(a Atom, id atom.ID) {
	l := c.logger.Enter("ReplaceAtom")
	buf := &bytes.Buffer{}
	enc := cyclic.Encoder(vle.Writer(buf))
	err := enc.Value(a.object)
	if err != nil {
		panic(err)
	}
	b := service.Binary{
		Data: buf.Bytes(),
	}
	capture, err := c.rpc.ReplaceAtom(c.captureID, uint64(id), b, l)
	if err != nil {
		panic(err)
	}
	c.LoadCapture(capture, false)
}

func (c *ApplicationContext) Theme() gxui.Theme                          { return c.theme }
func (c *ApplicationContext) Logger() *log.Splitter                      { return c.logger }
func (c *ApplicationContext) Rpc() service.RPC                           { return c.rpc }
func (c *ApplicationContext) DropDownOverlay() gxui.BubbleOverlay        { return c.dropDownOverlay }
func (c *ApplicationContext) ToolTipOverlay() gxui.BubbleOverlay         { return c.toolTipOverlay }
func (c *ApplicationContext) ToolTipController() *gxui.ToolTipController { return c.toolTipController }
func (c *ApplicationContext) Schema() service.Schema                     { return c.schema }
func (c *ApplicationContext) Atoms() []Atom                              { return c.atoms }
func (c *ApplicationContext) Hierarchy() atom.Group                      { return c.hierarchy }

//func (c *ApplicationContext) State() schema.Struct                       { return c.state }
func (c *ApplicationContext) SelectedAtomID() atom.ID          { return c.selectedAtomID }
func (c *ApplicationContext) SelectedAddress() memory.Pointer  { return c.selectedAddress }
func (c *ApplicationContext) SelectedObject() interface{}      { return c.selectedObject }
func (c *ApplicationContext) SelectedDevice() service.DeviceId { return c.selectedDevice }
func (c *ApplicationContext) Wireframe() bool                  { return c.wireframe }
func (c *ApplicationContext) ColorBuffer() gxui.Texture        { return c.colorBuffer }
func (c *ApplicationContext) DepthBuffer() gxui.Texture        { return c.depthBuffer }
func (c *ApplicationContext) CaptureID() service.CaptureId     { return c.captureID }
func (c *ApplicationContext) Capture() service.Capture         { return c.capture }

func (c *ApplicationContext) OnAtomSelected(f func()) gxui.EventSubscription {
	return c.onAtomSelected.Listen(f)
}

func (c *ApplicationContext) OnObjectSelected(f func()) gxui.EventSubscription {
	return c.onObjectSelected.Listen(f)
}

func (c *ApplicationContext) OnAddressSelected(f func()) gxui.EventSubscription {
	return c.onAddressSelected.Listen(f)
}

func (c *ApplicationContext) OnColorBufferUpdate(f func()) gxui.EventSubscription {
	return c.onColorBufferUpdate.Listen(f)
}

func (c *ApplicationContext) OnDepthBufferUpdate(f func()) gxui.EventSubscription {
	return c.onDepthBufferUpdate.Listen(f)
}

func (c *ApplicationContext) OnRequestReplay(f func()) gxui.EventSubscription {
	return c.onRequestReplay.Listen(f)
}

func (c *ApplicationContext) OnWireframeChanged(f func()) gxui.EventSubscription {
	return c.onWireframeChanged.Listen(f)
}

func (c *ApplicationContext) OnDeviceSelected(f func()) gxui.EventSubscription {
	return c.onDeviceSelected.Listen(f)
}

func (c *ApplicationContext) OnAtomsUpdated(f func()) gxui.EventSubscription {
	return c.onAtomsUpdated.Listen(f)
}

func (c *ApplicationContext) OnHierarchyUpdated(f func()) gxui.EventSubscription {
	return c.onHierarchyUpdated.Listen(f)
}

func (c *ApplicationContext) OnReportUpdated(f func()) gxui.EventSubscription {
	return c.onReportUpdated.Listen(f)
}

func (c *ApplicationContext) OnStateUpdated(f func()) gxui.EventSubscription {
	return c.onStateUpdated.Listen(f)
}

func (c *ApplicationContext) OnTimingInfoUpdated(f func()) gxui.EventSubscription {
	return c.onTimingInfoUpdated.Listen(f)
}
