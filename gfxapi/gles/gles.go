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

//go:generate apic validate gles.api

//go:generate apic template gles.api ../templates/api.go.tmpl
//go:generate apic template gles.api ../templates/replay_writer.go.tmpl
//go:generate apic template gles.api ../templates/schema.go.tmpl
//go:generate apic template gles.api ../templates/state_mutator.go.tmpl

//go:generate apic template --dir ../../cc/replayd/src gles.api ../templates/GfxApi.cpp.tmpl
//go:generate apic template --dir ../../cc/replayd/src gles.api ../templates/GfxApi.h.tmpl

//go:generate apic template --dir ../../cc/gfxspy2/src gles.api ../templates/spy.h.tmpl
//go:generate apic template --dir ../../cc/gfxspy2/src gles.api ../templates/state.h.tmpl
//go:generate apic template --dir ../../cc/gfxspy2/src gles.api ../templates/types.h.tmpl

// Package gles implementes the API interface for the OpenGL ES graphics library.
package gles
