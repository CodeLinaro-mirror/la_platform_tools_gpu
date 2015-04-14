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

// Package test is the integration test suite for the api compiler and templates.
package test

import (
	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/gfxapi/state"
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/service"
)

type State struct {
	Globals
	ValidateOutput bool
}

func (s *State) GetFramebufferAttachmentSize(att state.FramebufferAttachment) (uint32, uint32, error) {
	return 0, 0, nil
}

func (i remapped) remap(a atom.Atom, s *state.State) (interface{}, bool) {
	return i, true
}

func (a api) ColorBuffer(ctx *replay.Context, mgr *replay.Manager, after atom.ID, width, height uint32, wireframe bool) <-chan gfxapi.Image {
	return nil
}

func (a api) DepthBuffer(ctx *replay.Context, mgr *replay.Manager, after atom.ID) <-chan gfxapi.Image {
	return nil
}

func (a api) TimeCalls(ctx *replay.Context, mgr *replay.Manager, mask service.TimingMask) <-chan gfxapi.CallTiming {
	return nil
}
