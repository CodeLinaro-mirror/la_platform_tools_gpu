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
	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/interval"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/service"
	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

type MemoryAdapter struct {
	gxui.AdapterBase
	appCtx       *ApplicationContext
	commandID    atom.ID
	baseAddress  memory.Pointer
	bytesPerLine int
	dataType     DataType
}

func CreateMemoryAdapter(appCtx *ApplicationContext) *MemoryAdapter {
	return &MemoryAdapter{
		appCtx:       appCtx,
		bytesPerLine: 32,
		dataType:     U8(0),
	}
}

func (a *MemoryAdapter) IndexOfAddress(addr memory.Pointer) int {
	return int(addr/memory.Pointer(a.bytesPerLine) - a.baseAddress)
}

func (a *MemoryAdapter) AddressAtIndex(index int) memory.Pointer {
	return a.baseAddress + memory.Pointer(index*a.bytesPerLine)
}

func (a *MemoryAdapter) SetData(atomID atom.ID, baseAddress memory.Pointer) {
	if atomID != InvalidAtomID {
		a.commandID = atomID
	}
	a.baseAddress = baseAddress
	a.DataReplaced()
}

func (a *MemoryAdapter) SetDataType(dataType DataType) {
	a.dataType = dataType
	a.DataReplaced()
}

func (a *MemoryAdapter) Size(theme gxui.Theme) math.Size {
	return math.Size{W: math.MaxSize.W, H: 14}
}

func (a *MemoryAdapter) Count() int {
	if a.appCtx.CaptureID().Valid() {
		return 10000 //Who knows?
	} else {
		return 0
	}
}

func (a *MemoryAdapter) ItemAt(index int) gxui.AdapterItem {
	return a.AddressAtIndex(index)
}

func (a *MemoryAdapter) ItemIndex(item gxui.AdapterItem) int {
	addr := item.(memory.Pointer)
	return a.IndexOfAddress(addr)
}

func (a *MemoryAdapter) Create(t gxui.Theme, index int) gxui.Control {
	ll := t.CreateLinearLayout()
	ll.SetDirection(gxui.LeftToRight)
	base := a.AddressAtIndex(index)
	ll.AddChild(CreateLabel(t, base.String(), LINE_NUMBER_COLOR, true))

	var cancel chan<- struct{}
	ll.OnAttach(func() {
		cancel = a.appCtx.RequestMemory(a.commandID, base, uint64(a.bytesPerLine), func(info service.MemoryInfo) {
			var current, stale memory.RangeList
			info.Current.Unpack(&current)
			info.Stale.Unpack(&stale)

			addr := base
			data := info.Data
			dataType := a.dataType
			dataTypeSize := dataType.SizeBytes()
			for len(data) >= dataTypeSize {
				switch {
				case interval.Contains(&current, uint64(addr)):
					ll.AddChild(CreateLabel(t, dataType.Read(data).String()+" ", MEMORY_COLOR, true))
				case interval.Contains(&stale, uint64(addr)):
					ll.AddChild(CreateLabel(t, dataType.Read(data).String()+" ", STALE_MEMORY_COLOR, true))
				default:
					ll.AddChild(CreateLabel(t, dataType.Unknown()+" ", STALE_MEMORY_COLOR, true))
				}
				addr += memory.Pointer(dataTypeSize)
				data = data[dataTypeSize:]
			}
		})
	})
	ll.OnDetach(func() { close(cancel) })
	return ll
}
