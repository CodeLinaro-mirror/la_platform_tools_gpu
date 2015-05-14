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

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/service"
	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

////////////////////////////////////////////////////////////////////////////////
// MemoryImageAdapter
////////////////////////////////////////////////////////////////////////////////
type MemoryImageAdapter struct {
	gxui.AdapterBase
	appCtx        *ApplicationContext
	commandID     atom.ID
	baseAddress   memory.Pointer
	pixelsPerLine int
	pixelType     PixelType
}

func CreateMemoryImageAdapter(appCtx *ApplicationContext) *MemoryImageAdapter {
	return &MemoryImageAdapter{
		appCtx:        appCtx,
		pixelsPerLine: 32,
		pixelType:     RGBA8888(0),
	}
}

func (a *MemoryImageAdapter) IndexOfAddress(addr memory.Pointer) int {
	return int(addr/memory.Pointer(a.pixelsPerLine) - a.baseAddress)
}

func (a *MemoryImageAdapter) AddressAtIndex(index int) memory.Pointer {
	bytesPerPixel := a.pixelType.SizeBytes()
	bytesPerLine := a.pixelsPerLine * bytesPerPixel
	return a.baseAddress + memory.Pointer(index*bytesPerLine)
}

func (a *MemoryImageAdapter) SetData(atomID atom.ID, baseAddress memory.Pointer) {
	if atomID != InvalidAtomID {
		a.commandID = atomID
	}
	a.baseAddress = baseAddress
	a.DataReplaced()
}

func (a *MemoryImageAdapter) SetPixelType(pixelType PixelType) {
	a.pixelType = pixelType
	a.DataReplaced()
}

func (a *MemoryImageAdapter) Size(theme gxui.Theme) math.Size {
	return math.Size{W: math.MaxSize.W, H: 14}
}

func (a *MemoryImageAdapter) Count() int {
	if a.appCtx.CaptureID().Valid() {
		return 10000 //Who knows?
	} else {
		return 0
	}
}

func (a *MemoryImageAdapter) ItemAt(index int) gxui.AdapterItem {
	return a.AddressAtIndex(index)
}

func (a *MemoryImageAdapter) ItemIndex(item gxui.AdapterItem) int {
	addr := item.(memory.Pointer)
	return a.IndexOfAddress(addr)
}

func (a *MemoryImageAdapter) Create(t gxui.Theme, index int) gxui.Control {
	bytesPerPixel := a.pixelType.SizeBytes()
	bytesPerLine := a.pixelsPerLine * bytesPerPixel
	base := a.AddressAtIndex(index)

	ll := t.CreateLinearLayout()
	ll.SetDirection(gxui.LeftToRight)
	ll.AddChild(CreateLabel(t, fmt.Sprintf("%.16x ", base), LINE_NUMBER_COLOR, true))

	var cancel chan<- struct{}
	ll.OnAttach(func() {
		cancel = a.appCtx.RequestMemory(a.commandID, base, uint64(bytesPerLine), func(info service.MemoryInfo) {
			addr := base
			data := info.Data
			pixelSizeDips := a.Size(t).H
			for len(data) >= bytesPerPixel {
				color := a.pixelType.Read(data)
				img := t.CreateImage()
				img.SetExplicitSize(math.Size{W: pixelSizeDips, H: pixelSizeDips})
				img.SetBackgroundBrush(gxui.Brush{Color: color})
				ll.AddChild(img)
				addr += memory.Pointer(bytesPerPixel)
				data = data[bytesPerPixel:]
			}
		})
	})
	ll.OnDetach(func() { close(cancel) })
	return ll
}
