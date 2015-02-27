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

package atom

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/interval"
)

// Range describes an interval of atoms in a stream.
type Range struct {
	Start ID // The first atom within the range.
	End   ID // One past the last atom within the range.
}

// String returns a string representing the range.
func (i Range) String() string {
	return fmt.Sprintf("[%.6d-%.6d]", i.Start, i.End-1)
}

// Contains returns true if atomID is within the range, otherwise false.
func (i Range) Contains(atomID ID) bool {
	return atomID >= i.Start && atomID < i.End
}

// Length returns the number of atoms in the range.
func (i Range) Length() uint64 {
	return uint64(i.End - i.Start)
}

// Range returns the start and end of the range.
func (i Range) Range() (start, end ID) {
	return i.Start, i.End
}

// First returns the first atom identifier within the range.
func (i Range) First() ID {
	return i.Start
}

// Last returns the last atom identifier within the range.
func (i Range) Last() ID {
	return i.End - 1
}

// Span returns the start and end of the range as a U64Span.
func (i Range) Span() interval.U64Span {
	return interval.U64Span{Start: uint64(i.Start), End: uint64(i.End)}
}

// SetSpan sets the start and end range using a U64Span.
func (i *Range) SetSpan(span interval.U64Span) {
	i.Start = ID(span.Start)
	i.End = ID(span.End)
}

// Encode encodes the Range using the specified encoder.
func (i Range) Encode(e *binary.Encoder) error {
	if err := i.Start.Encode(e); err != nil {
		return err
	}
	if err := i.End.Encode(e); err != nil {
		return err
	}
	return nil
}

// Decode decodes the Range using the specified decoder.
func (i *Range) Decode(d *binary.Decoder) error {
	if err := i.Start.Decode(d); err != nil {
		return err
	}
	if err := i.End.Decode(d); err != nil {
		return err
	}
	return nil
}
