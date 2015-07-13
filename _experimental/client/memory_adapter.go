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

	"android.googlesource.com/platform/tools/gpu/interval"
	"android.googlesource.com/platform/tools/gpu/service"
	"android.googlesource.com/platform/tools/gpu/service/path"
	"android.googlesource.com/platform/tools/gpu/task"
	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

type MemoryAdapter struct {
	gxui.AdapterBase
	appCtx       *ApplicationContext
	after        *path.Atom
	baseAddress  uint64
	bytesPerLine uint64
	dataType     DataType
}

func CreateMemoryAdapter(appCtx *ApplicationContext) *MemoryAdapter {
	return &MemoryAdapter{
		appCtx:       appCtx,
		bytesPerLine: 32,
		dataType:     U8(0),
	}
}

func (a *MemoryAdapter) Update(after *path.Atom, baseAddress uint64) {
	a.after = after
	a.baseAddress = baseAddress
	a.DataReplaced()
}

func (a *MemoryAdapter) IndexOfAddress(addr uint64) int {
	return int(addr/a.bytesPerLine - a.baseAddress)
}

func (a *MemoryAdapter) AddressAtIndex(index int) uint64 {
	return a.baseAddress + uint64(index)*a.bytesPerLine
}

func (a *MemoryAdapter) SetDataType(dataType DataType) {
	a.dataType = dataType
	a.DataReplaced()
}

func (a *MemoryAdapter) Size(theme gxui.Theme) math.Size {
	return math.Size{W: math.MaxSize.W, H: 14}
}

func (a *MemoryAdapter) Count() int {
	if a.after != nil {
		return 10000 //Who knows?
	} else {
		return 0
	}
}

func (a *MemoryAdapter) ItemAt(index int) gxui.AdapterItem {
	return a.AddressAtIndex(index)
}

func (a *MemoryAdapter) ItemIndex(item gxui.AdapterItem) int {
	addr := item.(uint64)
	return a.IndexOfAddress(addr)
}

func (a *MemoryAdapter) Create(theme gxui.Theme, index int) gxui.Control {
	ll := theme.CreateLinearLayout()
	ll.SetDirection(gxui.LeftToRight)
	base := a.AddressAtIndex(index)

	t := task.New()
	update := func() {
		t.Run(requestMemory{a.appCtx, a.after, base, a.bytesPerLine, func(info service.MemoryInfo) {
			ll.RemoveAll()
			ll.AddChild(CreateLabel(theme, fmt.Sprintf("%.16x ", base), LINE_NUMBER_COLOR, true))
			offset := uint64(0)
			data := info.Data
			dataType := a.dataType
			dataTypeSize := dataType.SizeBytes()
			for len(data) >= dataTypeSize {
				switch {
				case interval.Contains(&info.Writes, offset):
					ll.AddChild(CreateMonospaceLabel(a.appCtx, dataType.Read(data).String()+" ", WRITE_MEMORY_COLOR, true))
				case interval.Contains(&info.Reads, offset):
					ll.AddChild(CreateMonospaceLabel(a.appCtx, dataType.Read(data).String()+" ", READ_MEMORY_COLOR, true))
				case interval.Contains(&info.Observed, offset):
					ll.AddChild(CreateMonospaceLabel(a.appCtx, dataType.Read(data).String()+" ", STALE_MEMORY_COLOR, true))
				default:
					ll.AddChild(CreateMonospaceLabel(a.appCtx, dataType.Unknown()+" ", STALE_MEMORY_COLOR, true))
				}
				offset += uint64(dataTypeSize)
				data = data[dataTypeSize:]
			}
		}})
	}

	ll.OnAttach(update)
	ll.OnDetach(t.Cancel)
	return ll
}
