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
	"fmt"

	"android.googlesource.com/platform/tools/gpu/binary"
)

// FramebufferAttachment values indicate the type of frame buffer attachment.
type FramebufferAttachment uint32

const (
	FramebufferAttachmentColor FramebufferAttachment = iota
	FramebufferAttachmentDepth
	FramebufferAttachmentStencil
)

// API is the common interface to a graphics programming api.
type API interface {
	// Name returns the official name of the api.
	Name() string

	// ID returns the unique API identifier.
	ID() ID

	// GetFramebufferAttachmentSize returns the width and height of the framebuffer at the given attachment.
	GetFramebufferAttachmentSize(state *State, attachment FramebufferAttachment) (width uint32, height uint32, err error)
}

// ID is an API identifier
type ID binary.ID

// Valid returns true if the id is not the default zero value.
func (i ID) Valid() bool { return binary.ID(i).Valid() }

var apis map[ID]API = make(map[ID]API)

// Register adds an api to the understood set.
// It is illegal to register the same name twice.
func Register(api API) {
	id := api.ID()
	if _, present := apis[id]; present {
		panic(fmt.Errorf("API %s registered more than once", id))
	}
	apis[id] = api
}

// Find looks up a graphics API by identifier.
// If the id has not been registered, it returns nil.
func Find(id ID) API {
	return apis[id]
}
