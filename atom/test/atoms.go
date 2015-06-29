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
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
)

type AtomA struct {
	binary.Generate `id:"AtomAID"`
	ID              atom.ID
	AtomFlags       atom.Flags
}

func (a *AtomA) API() gfxapi.API                                           { return nil }
func (a *AtomA) Flags() atom.Flags                                         { return a.AtomFlags }
func (a *AtomA) Observations() *atom.Observations                          { return &atom.Observations{} }
func (a *AtomA) Mutate(*gfxapi.State, database.Database, log.Logger) error { return nil }

type AtomB struct {
	binary.Generate `id:"AtomBID"`
	ID              atom.ID
	Bool            bool
}

func (a *AtomB) API() gfxapi.API                                           { return nil }
func (a *AtomB) Flags() atom.Flags                                         { return 0 }
func (a *AtomB) Observations() *atom.Observations                          { return &atom.Observations{} }
func (a *AtomB) Mutate(*gfxapi.State, database.Database, log.Logger) error { return nil }

type AtomC struct {
	binary.Generate `id:"AtomCID"`
	String          string
}

func (a *AtomC) API() gfxapi.API                                           { return nil }
func (a *AtomC) Flags() atom.Flags                                         { return 0 }
func (a *AtomC) Observations() *atom.Observations                          { return &atom.Observations{} }
func (a *AtomC) Mutate(*gfxapi.State, database.Database, log.Logger) error { return nil }
