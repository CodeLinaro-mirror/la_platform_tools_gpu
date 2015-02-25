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

//go:generate apic template gfxapi_test.api ../templates/api.go.tmpl
//go:generate apic template gfxapi_test.api ../templates/replay_writer.go.tmpl
//go:generate apic template gfxapi_test.api ../templates/schema.go.tmpl
//go:generate apic template gfxapi_test.api ../templates/state_mutator.go.tmpl

// Package test is the integration test suite for the api compiler and templates.
package test

import (
	"gaze/gfxapi"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/service"
)

type state struct {
	Globals
	Mem memory.Memory
}

func (s *state) Memory() *memory.Memory {
	return &s.Mem
}

func (s *state) GetFramebufferAttachmentSize(att gfxapi.FramebufferAttachment) (uint32, uint32, error) {
	return 0, 0, nil
}

func initialState() *state {
	return &state{}
}

func (i remapped) remap(a atom.Atom, s *state) (interface{}, bool) { return i, true }

func (a api) ColorBuffer(ctx *replay.Context, mgr *replay.Manager, after atom.Id, width, height uint32, wireframe bool) <-chan gfxapi.Image {
	return nil
}
func (a api) DepthBuffer(ctx *replay.Context, mgr *replay.Manager, after atom.Id) <-chan gfxapi.Image {
	return nil
}
func (a api) TimeCalls(ctx *replay.Context, mgr *replay.Manager, mask service.TimingMask) <-chan gfxapi.CallTiming {
	return nil
}
