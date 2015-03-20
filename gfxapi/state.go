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
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/memory"
)

// FramebufferAttachment values indicate the type of frame buffer attachment.
type FramebufferAttachment uint32

const (
	FramebufferAttachmentColor   FramebufferAttachment = iota
	FramebufferAttachmentDepth   FramebufferAttachment = iota
	FramebufferAttachmentStencil FramebufferAttachment = iota
)

// State represents the common interface to managed state of a graphics system.
type State interface {
	// Memory returns the memory mapping for the current state.
	Memory() *memory.Memory
	// Encode writes the state to a binary stream.
	Encode(binary.Encoder) error
	// GetFramebufferAttachmentSize returns the width and height of the framebuffer at the given attachment.
	GetFramebufferAttachmentSize(attachment FramebufferAttachment) (uint32, uint32, error)
}
