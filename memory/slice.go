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
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/log"
)

// Slice is the interface for a data source that can be resolved to a byte slice
// with Get, or 'sliced' to a subset of the data source.
type Slice interface {
	// Get resolves all the bytes representing the slice.
	Get(d database.Database, l log.Logger) ([]byte, error)

	// ResourceID returns the identifier of the resource representing the slice,
	// creating a new resource if it isn't already backed by one.
	ResourceID(d database.Database, l log.Logger) (binary.ID, error)

	// Size returns the number of bytes that would be returned by calling Get.
	Size() uint64

	// Slice returns a new Slice referencing a subset range of the data.
	// The range r is relative to the base of the Slice. For example a slice of
	// [0, 4] would return a Slice referencing the first 5 bytes of this Slice.
	// Attempting to slice outside the range of this Slice will result in a
	// panic.
	Slice(r Range) Slice

	// ValidRanges returns the list of slice-relative memory ranges that contain
	// valid (non-zero) data that can be read with Get.
	ValidRanges() RangeList
}
