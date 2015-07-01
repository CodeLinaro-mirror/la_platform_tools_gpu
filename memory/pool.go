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

package memory

import (
	"fmt"
	"strings"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/interval"
	"android.googlesource.com/platform/tools/gpu/log"
)

// Pool represents an unbounded and isolated memory space. Pool can be used
// to represent the application address space, or hidden GPU Pool.
//
// Pool can be sliced into smaller regions which can be read or written to.
// All writes to Pool or its slices do not actually perform binary data
// copies, but instead all writes are stored as lightweight records. Only when a
// Pool slice has Get called will any resolving, loading or copying of binary
// data occur.
type Pool struct {
	binary.Generate `disable:"true"`
	writes          []write
	OnRead          func(Range)
	OnWrite         func(Range)
}

// PoolID is an indentifier of a Pool.
type PoolID uint32

// ApplicationPool is the PoolID of Pool representing the application's memory
// address space.
const ApplicationPool PoolID = 0

// Slice returns a Slice referencing the subset of the Pool range.
func (m *Pool) Slice(rng Range) Slice {
	for i := len(m.writes) - 1; i >= 0; i-- {
		w := m.writes[i]
		if w.dst.Overlaps(rng) {
			if w.dst == rng {
				return w.src
			} else {
				return poolSlice{m, rng, len(m.writes)}
			}
		}
	}
	return nullSlice(rng.Size)
}

// At returns an unbounded Slice starting at p.
func (m *Pool) At(addr uint64) Slice {
	return m.Slice(Range{Base: addr, Size: ^uint64(0) - addr})
}

// Write copies the slice src to the address dst.
func (m *Pool) Write(dst uint64, src Slice) {
	rng := Range{Base: dst, Size: src.Size()}
	m.writes = append(m.writes, write{rng, src})
}

// String returns the full history of writes performed to this pool.
func (m *Pool) String() string {
	l := make([]string, len(m.writes)+1)
	l[0] = fmt.Sprintf("Pool(%p):", m)
	for i, w := range m.writes {
		l[i+1] = fmt.Sprintf("(%d) %v <- %v", i, w.dst, w.src)
	}
	return strings.Join(l, "\n")
}

type write struct {
	dst Range
	src Slice
}

type poolSlice struct {
	pool *Pool // The pool that this slice belongs to.
	rng  Range // The memory range of the slice.
	at   int   // The index in the pool's writes slice when the slice was created.
}

func (m poolSlice) Get(d database.Database, l log.Logger) ([]byte, error) {
	buffer := make([]byte, m.rng.Size)
	remaining := RangeList{m.rng}
	for i := m.at - 1; i >= 0; i-- {
		w := m.pool.writes[i]
		if s, c := interval.Intersect(&remaining, w.dst.Span()); c > 0 {
			src, err := w.src.Get(d, l)
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

func (m poolSlice) ResourceID(d database.Database, l log.Logger) (binary.ID, error) {
	bytes, err := m.Get(d, l)
	if err != nil {
		return binary.ID{}, err
	}
	return database.StoreBlob(bytes, d, l)
}

func (m poolSlice) Slice(rng Range) Slice {
	if uint64(rng.Last()) > m.rng.Size {
		panic(fmt.Errorf("%v.Slice(%v) - out of bounds", m.String(), rng))
	}
	rng.Base += m.rng.Base
	return poolSlice{m.pool, rng, len(m.pool.writes)}
}

func (m poolSlice) ValidRanges() RangeList {
	// build a list of unknown ranges
	unknown := RangeList{m.rng}
	for i := m.at - 1; i >= 0 && len(unknown) > 0; i-- {
		w := m.pool.writes[i]
		interval.Remove(&unknown, w.dst.Span())
	}
	// invert the unknowns to a list of knowns
	knowns := RangeList{m.rng}
	for _, u := range unknown {
		interval.Remove(&knowns, u.Span())
	}
	// return slice-relative addresses
	valid := make(RangeList, len(knowns))
	for i, k := range knowns {
		valid[i] = Range{Base: k.Base - m.rng.Base, Size: k.Size}
	}
	return valid
}

func (m poolSlice) Size() uint64 {
	return m.rng.Size
}

func (m poolSlice) String() string {
	return fmt.Sprintf("Pool(%p)@%d%v", m.pool, m.at, m.rng)
}
