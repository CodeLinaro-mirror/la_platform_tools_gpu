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

	"android.googlesource.com/platform/tools/gpu/service"
	"android.googlesource.com/platform/tools/gpu/service/path"
	"android.googlesource.com/platform/tools/gpu/task"
	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

////////////////////////////////////////////////////////////////////////////////
// MemoryImageAdapter
////////////////////////////////////////////////////////////////////////////////
type MemoryImageAdapter struct {
	gxui.AdapterBase
	appCtx        *ApplicationContext
	path          *path.MemoryRange
	pixelsPerLine uint64
	pixelType     PixelType
}

func CreateMemoryImageAdapter(appCtx *ApplicationContext) *MemoryImageAdapter {
	return &MemoryImageAdapter{
		appCtx:        appCtx,
		pixelsPerLine: 32,
		pixelType:     RGBA8888(0),
	}
}

func (a *MemoryImageAdapter) Update(p *path.MemoryRange) {
	a.path = p
	a.DataReplaced()
}

func (a *MemoryImageAdapter) IndexOfAddress(addr uint64) int {
	if a.path != nil {
		return int(addr/a.pixelsPerLine - a.path.Address)
	} else {
		return 0
	}
}

func (a *MemoryImageAdapter) AddressAtIndex(index int) uint64 {
	if a.path != nil {
		bytesPerPixel := uint64(a.pixelType.SizeBytes())
		bytesPerLine := a.pixelsPerLine * bytesPerPixel
		return a.path.Address + uint64(index)*bytesPerLine
	} else {
		return 0
	}
}

func (a *MemoryImageAdapter) SetPixelType(pixelType PixelType) {
	a.pixelType = pixelType
	a.DataReplaced()
}

func (a *MemoryImageAdapter) Size(theme gxui.Theme) math.Size {
	return math.Size{W: math.MaxSize.W, H: 14}
}

func (a *MemoryImageAdapter) Count() int {
	if a.path != nil {
		return 10000 //Who knows?
	} else {
		return 0
	}
}

func (a *MemoryImageAdapter) ItemAt(index int) gxui.AdapterItem {
	return a.AddressAtIndex(index)
}

func (a *MemoryImageAdapter) ItemIndex(item gxui.AdapterItem) int {
	addr := item.(uint64)
	return a.IndexOfAddress(addr)
}

func (a *MemoryImageAdapter) Create(theme gxui.Theme, index int) gxui.Control {
	bytesPerPixel := uint64(a.pixelType.SizeBytes())
	base, size := a.AddressAtIndex(index), a.pixelsPerLine*bytesPerPixel

	ll := theme.CreateLinearLayout()
	ll.SetDirection(gxui.LeftToRight)

	t := task.New()
	update := func() {
		p := a.path.Clone().(*path.MemoryRange)
		p.Address = base
		p.Size = size
		t.Run(requestMemory{a.appCtx, p, func(info service.MemoryInfo) {
			ll.RemoveAll()
			ll.AddChild(createLabel(a.appCtx, fmt.Sprintf("%.16x ", base), LINE_NUMBER_COLOR))
			addr := base
			data := info.Data
			pixelSizeDips := a.Size(theme).H
			for uint64(len(data)) >= bytesPerPixel {
				color := a.pixelType.Read(data)
				img := theme.CreateImage()
				img.SetExplicitSize(math.Size{W: pixelSizeDips, H: pixelSizeDips})
				img.SetBackgroundBrush(gxui.Brush{Color: color})
				ll.AddChild(img)
				addr += bytesPerPixel
				data = data[bytesPerPixel:]
			}
		}})
	}

	ll.OnAttach(update)
	ll.OnDetach(t.Cancel)
	return ll
}
