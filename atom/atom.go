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

// Package atom provides the fundamental types used to describe a capture stream.
package atom

// binary: cpp = atom
// binary: java.source = adt/idea/android/src
// binary: java.package = com.android.tools.idea.editors.gfxtrace.service.atom
// binary: java.indent = "  "
// binary: java.member_prefix = my
// binary: java.disable.ID = true

import (
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
)

// Atom is the interface implemented by all objects that describe an single
// event in a capture stream. Typical implementations of Atom describe an
// application's call to a graphics API function or provide meta-data describing
// observed memory or state at the time of capture.
//
// Each implementation of Atom should have a unique and stable Signature to ensure
// binary compatibility with old capture formats. Any change to the Atom's
// binary format should also result in a new Signature.
type Atom interface {
	binary.Object

	// API returns the graphics API id this atom belongs to.
	API() gfxapi.ID

	// Flags returns the flags of the atom.
	Flags() Flags

	// Observations returns all the memory observations made by the atom.
	Observations() *Observations

	// Mutate mutates the State using the atom.
	Mutate(*gfxapi.State, database.Database, log.Logger) error
}

// ID is the index of an atom in an atom stream.
type ID uint64

// NoID is used when you have to pass an ID, but don't have one to use.
const NoID = ID(1<<63 - 1) // use max int64 for the benefit of java

// AtomCast is automatically called by the generated decoders.
func AtomCast(obj binary.Object) Atom {
	return obj.(Atom)
}
