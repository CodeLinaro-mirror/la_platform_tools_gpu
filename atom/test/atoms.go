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

// Package test provides testing helpers for the atom package.
package test

import (
	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
)

const AtomIDA = atom.TypeID(1)
const AtomIDB = atom.TypeID(2)
const AtomIDC = atom.TypeID(3)

type AtomA struct {
	binary.Generate
	ID        atom.ID
	AtomFlags atom.Flags
}

func (a *AtomA) API() gfxapi.API            { return nil }
func (a *AtomA) TypeID() atom.TypeID        { return AtomIDA }
func (a *AtomA) Flags() atom.Flags          { return a.AtomFlags }
func (a *AtomA) Mutate(*gfxapi.State) error { return nil }

type AtomB struct {
	binary.Generate
	ID   atom.ID
	Bool bool
}

func (a *AtomB) API() gfxapi.API            { return nil }
func (a *AtomB) TypeID() atom.TypeID        { return AtomIDB }
func (a *AtomB) Flags() atom.Flags          { return 0 }
func (a *AtomB) Mutate(*gfxapi.State) error { return nil }

type AtomC struct {
	binary.Generate
	String string
}

func (a *AtomC) API() gfxapi.API            { return nil }
func (a *AtomC) TypeID() atom.TypeID        { return AtomIDC }
func (a *AtomC) Flags() atom.Flags          { return 0 }
func (a *AtomC) Mutate(*gfxapi.State) error { return nil }

func init() {
	atom.Register(atom.TypeInfo{ID: AtomIDA, New: func() atom.Atom { return &AtomA{} }})
	atom.Register(atom.TypeInfo{ID: AtomIDB, New: func() atom.Atom { return &AtomB{} }})
	atom.Register(atom.TypeInfo{ID: AtomIDC, New: func() atom.Atom { return &AtomC{} }})
}
