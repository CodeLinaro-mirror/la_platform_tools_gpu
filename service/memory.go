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

package service

import "android.googlesource.com/platform/tools/gpu/memory"

// Pack packs the memory Range o into the RPC-friendly MemoryRange structure.
func (r *MemoryRange) Pack(o memory.Range) {
	r.Base = o.Base
	r.Size = o.Size
}

// Unpack unpacks the RPC-friendly MemoryRange structure into the memory Range
// o.
func (r MemoryRange) Unpack(o *memory.Range) {
	(*o) = memory.Range{Base: r.Base, Size: r.Size}
}

// Pack packs the memory RangeList o into the RPC-friendly MemoryRangeArray
// structure.
func (l *MemoryRangeArray) Pack(o memory.RangeList) {
	*l = make(MemoryRangeArray, len(o))
	for i := range o {
		(*l)[i].Pack(o[i])
	}
}

// Unpack unpacks the RPC-friendly MemoryRangeArray structure into the memory
// RangeList o.
func (l MemoryRangeArray) Unpack(o *memory.RangeList) {
	*o = make(memory.RangeList, len(l))
	for i := range l {
		l[i].Unpack(&(*o)[i])
	}
}
