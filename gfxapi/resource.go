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

package gfxapi

import (
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/log"
)

// Resource represents an asset in a capture.
type Resource interface {
	// ResourceName returns the UI name for the resource.
	ResourceName() string

	// ResourceType returns the type of this resource.
	ResourceType() ResourceType

	// ResourceData returns the resource data given the current state.
	ResourceData(s *State, d database.Database, l log.Logger) (interface{}, error)
}

// ResourceType is an enumerator of resource types.
type ResourceType int

const (
	// TypeTexture represents the Texture resource type
	TypeTexture = ResourceType(iota)
)
