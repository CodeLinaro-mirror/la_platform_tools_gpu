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

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/interval"
)

func min(a, b uint64) uint64 {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b uint64) uint64 {
	if a > b {
		return a
	} else {
		return b
	}
}

// Range represents a region of memory.
type Range struct {
	binary.Generate
	Base Pointer // A pointer to the first byte in the memory range.
	Size uint64  // The size in bytes of the memory range.
}

// Expand returns a new Range that is grown to include the pointer p.
func (i Range) Expand(p Pointer) Range {
	if i.Base > p {
		i.Base = p
	}
	if i.Last() < p {
		i.Size += uint64(p - i.Last())
	}
	return i
}

// Contains returns true if the pointer p is within the Range.
func (i Range) Contains(p Pointer) bool {
	return i.First() <= p && p <= i.Last()
}

// Overlaps returns true if other overlaps this memory range.
func (i Range) Overlaps(other Range) bool {
	a, b := i.Span(), other.Span()
	s, e := max(a.Start, b.Start), min(a.End, b.End)
	return s < e
}

// Intersect returns the Range that is common between this Range and other.
// If the two memory ranges do not intersect, then this function panics.
func (i Range) Intersect(other Range) Range {
	a, b := i.Span(), other.Span()
	s, e := max(a.Start, b.Start), min(a.End, b.End)
	if e < s {
		panic(fmt.Errorf("Intervals %v and %v do not intersect", i, other))
	}
	return Range{Base: Pointer(s), Size: e - s}
}

// First returns a Pointer to the first byte in the Range.
func (i Range) First() Pointer {
	return i.Base
}

// Last returns a Pointer to the last byte in the Range.
func (i Range) Last() Pointer {
	return i.End() - 1
}

// End returns a Pointer to one byte beyond the end of the Range.
func (i Range) End() Pointer {
	return i.Base + Pointer(i.Size)
}

// Span returns the Range as a U64Span.
func (i Range) Span() interval.U64Span {
	return interval.U64Span{
		Start: uint64(i.Base),
		End:   uint64(i.Base) + i.Size,
	}
}

func (i Range) String() string {
	return fmt.Sprintf("[%v-%v]", i.First(), i.Last())
}
