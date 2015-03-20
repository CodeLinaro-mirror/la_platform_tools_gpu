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

//go:generate codergen -go

// Package memory contains types used for representing and simulating memory
// observed in the capture.
package memory

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/interval"
	"android.googlesource.com/platform/tools/gpu/log"
)

// Data is the interface for a data source that can be resolved to a byte slice
// with Get.
//
// Size returns the number of bytes that would be returned by calling Get.
type Data interface {
	Get(db database.Database, logger log.Logger) ([]byte, error)
	Size() uint64
}

// DataSlicer extends the Data interface with aditional support for slicing.
//
// Slice returns a new DataSlicer referencing a subset range of the data.
// The range r is relative to the base of the DataSlicer. For example a slice of
// [0, 4] would return a DataSlicer referencing the first 5 bytes of this
// DataSlicer. Attempting to slice outside the range of this DataSlicer will
// result in a panic.
type DataSlicer interface {
	Data
	Slice(r Range) DataSlicer
}

// DataSliceWriter is similar to DataSlicer, but also has the Write method for
// modifying the data that the DataSliceWriter refers to.
//
// Write will replace the data in this DataSliceWriter with d.
// If d is shorter in length than the slice, then only the range [0, d.Size()-1]
// bytes will be replaced.
// A sliced DataSliceWriter shares the same data from which it was sliced - i.e.
// any mutations to a DataSliceWriter will also affect the Data it was sliced
// from.
type DataSliceWriter interface {
	Data
	Slice(r Range) DataSliceWriter
	Write(d Data)
}

type dataSlice struct {
	src Data
	rng Range
}

func slice(src Data, rng Range) DataSlicer {
	if rng.Base < 0 || uint64(rng.Last()) > src.Size() {
		panic(fmt.Errorf("Slice range (%s) out of bounds (%s)", rng, Range{Base: 0, Size: src.Size()}))
	}
	return dataSlice{src, rng}
}

func (s dataSlice) Get(db database.Database, logger log.Logger) ([]byte, error) {
	if data, err := s.src.Get(db, logger); err == nil {
		return data[s.rng.First() : s.rng.Last()+1], nil
	} else {
		return nil, err
	}
}

func (s dataSlice) Slice(rng Range) DataSlicer { return slice(s.src, rng) }
func (s dataSlice) Size() uint64               { return s.rng.Size }

type memorySlice struct {
	mem *Memory
	rng Range
}

func (m memorySlice) Get(db database.Database, logger log.Logger) ([]byte, error) {
	buffer := make([]byte, m.rng.Size)
	remaining := RangeList{m.rng}
	for i := len(m.mem.writes) - 1; i >= 0; i-- {
		w := m.mem.writes[i]
		if s, c := interval.Intersect(&remaining, w.dst.Span()); c > 0 {
			src, err := w.src.Get(db, logger)
			if err != nil {
				return nil, err
			}
			for i := s; i < s+c; i++ {
				intersect := w.dst.Intersect(remaining[i])
				src := src[intersect.Base-w.dst.Base:][:intersect.Size]
				dst := buffer[intersect.Base-m.rng.Base:][:intersect.Size]
				copy(dst, src)
			}
			interval.Remove(&remaining, w.dst.Span())
			if len(remaining) == 0 {
				return buffer, nil
			}
		}
	}
	return buffer, nil
}

func (m memorySlice) Slice(rng Range) DataSliceWriter {
	if rng.Base < 0 || uint64(rng.Last()) > m.rng.Size {
		panic(fmt.Errorf("Slice range (%s) out of bounds (%s)", rng, Range{Base: 0, Size: m.rng.Size}))
	}
	rng.Base += m.rng.Base
	return memorySlice{m.mem, rng}
}

func (m memorySlice) Size() uint64 {
	return m.rng.Size
}

func (m memorySlice) Write(src Data) {
	m.mem.writes = append(m.mem.writes, memoryWrite{m.rng, src})
}

type memoryWrite struct {
	dst Range
	src Data
}

// Memory represents an unbounded and isolated memory space. Memory can be used
// to represent the application address space, or hidden GPU memory.
//
// Memory can be sliced into smaller regions which can be read or written to.
// All writes to Memory or its slices do not actually perform binary data
// copies, but instead all writes are stored as lightweight records. Only when a
// Memory slice has Get called will any resolving, loading or copying of binary
// data occur.
type Memory struct {
	writes []memoryWrite
}

// Slice returns a DataSliceWriter referencing the subset of the Memory range.
func (m *Memory) Slice(rng Range) DataSliceWriter {
	return memorySlice{m, rng}
}

// Write copies d to the Memory slice [0, d.Size()-1].
func (m *Memory) Write(d Data) {
	m.writes = append(m.writes, memoryWrite{Range{Base: 0, Size: d.Size()}, d})
}

func (m *Memory) Encode(e binary.Encoder) error { return nil /* Currently not persisted */ }
func (m *Memory) Decode(d binary.Decoder) error { return nil /* Currently not persisted */ }
