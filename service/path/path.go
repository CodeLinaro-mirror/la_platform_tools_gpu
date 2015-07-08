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

// Package path contains types that represent data references.
package path

import "android.googlesource.com/platform/tools/gpu/binary"

// binary: java.source = adt/idea/android/src
// binary: java.package = com.android.tools.idea.editors.gfxtrace
// binary: java.indent = "  "
// binary: java.member_prefix = my

// Path is the interface for types that represent a reference to a capture,
// atom list, single atom, memory, state or sub-object. A path can be
// passed between client and server using RPCs in order to describe some data
// in a capture.
type Path interface {
	binary.Object

	// Path returns the string representation of the path.
	// The returned string must be consistent for equal paths.
	Path() string

	// Base returns the path that this path derives from.
	// If this path is a root, then Base returns nil.
	Base() Path

	// Clone returns a deep-copy of the path.
	Clone() Path

	// Validate checks the path for correctness, returning an error if any
	// issues are found.
	Validate() error
}

// Value is the expanded Path interface for types that represent a reference to
// a value type.
// The value referenced by this path may be a struct, array, slice, map or POD
// type.
type Value interface {
	// Value extends the Path interface.
	Path

	// Field returns the path to the field value with the specified name on the
	// struct object represented by this path.
	// The represented value type must be of type struct, otherwise the returned
	// path is invalid.
	Field(name string) *Field

	// Slice returns the path to the sliced subset of this array or slice
	// represented by this path.
	// The represented value type must be of type array or slice, otherwise the
	// returned path is invalid.
	Slice(start, end uint64) *Slice

	// ArrayIndex returns the path to the i'th element on the array or slice
	// represented by this path.
	// The represented value type must be of type array or slice, otherwise the
	// returned path is invalid.
	ArrayIndex(i uint64) *ArrayIndex

	// MapIndex returns the path to the map element with key k on the map object
	// represented by this path.
	// The represented value type must be of type map, otherwise the returned path
	// is invalid.
	MapIndex(k interface{}) *MapIndex
}
