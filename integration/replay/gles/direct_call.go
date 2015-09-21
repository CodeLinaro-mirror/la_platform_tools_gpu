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

package gles

import (
	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/replay/builder"
)

// directCall is an atom-wrapper that replaces calls to the the wrapped atom's
// Replay() method with Call(), preventing any state-mutation or memory
// observations to be performed. This is used by the tests that check for GL
// errors that would otherwise be caught by the state-mutator.
type directCall struct {
	binary.Generate
	atom caller // The wrapped atom.
}

type caller interface {
	atom.Atom
	Call(s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder)
}

func (c directCall) Replay(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error {
	c.atom.Call(s, d, l, b)
	return nil
}

// atom.Atom compliance
func (c directCall) API() gfxapi.ID                   { return c.atom.API() }
func (c directCall) Flags() atom.Flags                { return c.atom.Flags() }
func (c directCall) Observations() *atom.Observations { return c.atom.Observations() }
func (c directCall) Mutate(s *gfxapi.State, d database.Database, l log.Logger) error {
	return c.atom.Mutate(s, d, l)
}
