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

// Package atom provides the fundamental types used to describe a capture stream.
package atom

import "android.googlesource.com/platform/tools/gpu/binary"

// Atom is the interface implemented by all objects that describe an single
// event in a capture stream. Typical implementations of Atom describe an
// application's call to a graphics API function or provide meta-data describing
// observed memory or state at the time of capture.
//
// Each Atom instance is assigned a ContextID. The meaning of this identifier
// is API dependent, but is typically used to identify which thread or rendering
// context the atom belongs to. For example the ContextID is used to associate
// atoms of the same GLES context together.
//
// Each implementation of Atom should have a unique and stable TypeID to ensure
// binary compatibility with old capture formats. Any change to the Atom's
// binary format should also result in a new TypeID.
type Atom interface {
	binary.Object
	TypeID() TypeID
	ContextID() ContextID
	Flags() Flags
}

// ID is the index of an atom in an atom stream.
type ID uint64

// NoID is used when you have to pass an ID, but don't have one to use.
const NoID = ^ID(0)

// ContextIDs are used to identify which thread or rendering context each atom
// belongs to. The numerical value is meaningless aside from equality testing.
type ContextID uint32

// NoContext is used when the ContextID is not really a valid context.
const NoContextID = ^ContextID(0)
