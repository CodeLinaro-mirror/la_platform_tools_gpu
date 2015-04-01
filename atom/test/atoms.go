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

// Package test provides testing helpers for the atom package.
package test

import (
	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
)

const AtomIDA = atom.TypeID(1)
const AtomIDB = atom.TypeID(2)
const AtomIDC = atom.TypeID(3)

type AtomA struct {
	binary.Generate
	ID        atom.ID
	Context   atom.ContextID
	AtomFlags atom.Flags
}

func (AtomA) TypeID() atom.TypeID          { return AtomIDA }
func (a *AtomA) ContextID() atom.ContextID { return a.Context }
func (a *AtomA) Info() string              { return "" }
func (a *AtomA) Flags() atom.Flags         { return a.AtomFlags }

type AtomB struct {
	binary.Generate
	ID      atom.ID
	Context atom.ContextID
	Bool    bool
}

func (AtomB) TypeID() atom.TypeID          { return AtomIDB }
func (a *AtomB) ContextID() atom.ContextID { return a.Context }
func (a *AtomB) Info() string              { return "" }
func (a *AtomB) Flags() atom.Flags         { return 0 }

type AtomC struct {
	binary.Generate
	Context atom.ContextID
	String  string
}

func (AtomC) TypeID() atom.TypeID          { return AtomIDC }
func (a *AtomC) ContextID() atom.ContextID { return a.Context }
func (a *AtomC) Info() string              { return "" }
func (a *AtomC) Flags() atom.Flags         { return 0 }

func init() {
	atom.Register(atom.TypeInfo{ID: AtomIDA, New: func() atom.Atom { return &AtomA{} }})
	atom.Register(atom.TypeInfo{ID: AtomIDB, New: func() atom.Atom { return &AtomB{} }})
	atom.Register(atom.TypeInfo{ID: AtomIDC, New: func() atom.Atom { return &AtomC{} }})
}
