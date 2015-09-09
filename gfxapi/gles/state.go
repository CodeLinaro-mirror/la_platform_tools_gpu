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
	"fmt"

	"android.googlesource.com/platform/tools/gpu/gfxapi"
)

func getContext(s *gfxapi.State) *Context {
	return getState(s).getContext()
}

func (s *State) getContext() *Context {
	return s.Contexts[s.CurrentThread]
}

// TODO: When gfx api macros produce functions instead of inlining, move this logic
// to the gles.api file.
func (s *State) getFramebufferAttachmentSize(att gfxapi.FramebufferAttachment) (width, height uint32, err error) {
	c := s.getContext()
	if c == nil {
		return 0, 0, fmt.Errorf("No context bound")
	}

	framebufferID := c.BoundFramebuffers[GLenum_GL_READ_FRAMEBUFFER]

	framebuffer, ok := c.Instances.Framebuffers[framebufferID]
	if !ok {
		return 0, 0, fmt.Errorf("No GL_FRAMEBUFFER bound")
	}

	var attachment GLenum
	switch att {
	case gfxapi.FramebufferAttachmentColor:
		attachment = GLenum_GL_COLOR_ATTACHMENT0
	case gfxapi.FramebufferAttachmentDepth:
		attachment = GLenum_GL_DEPTH_ATTACHMENT
	case gfxapi.FramebufferAttachmentStencil:
		attachment = GLenum_GL_STENCIL_ATTACHMENT
	default:
		return 0, 0, fmt.Errorf("Framebuffer attachment %v unsupported by gles", att)
	}

	a, ok := framebuffer.Attachments[attachment]
	if !ok {
		return 0, 0, fmt.Errorf("%s is not bound", attachment)
	}

	switch a.Type {
	case GLenum_GL_TEXTURE:
		id := TextureId(a.Object)
		t := c.Instances.Textures[id]
		switch t.Kind {
		case TextureKind_TEXTURE2D:
			l := t.Texture2D[a.TextureLevel]
			return uint32(l.Width), uint32(l.Height), nil
		case TextureKind_CUBEMAP:
			l := t.Cubemap[a.TextureLevel]
			f := l.Faces[a.CubeMapFace]
			return uint32(f.Width), uint32(f.Height), nil
		default:
			return 0, 0, fmt.Errorf("Unknown texture kind %v", t.Kind)
		}
	case GLenum_GL_RENDERBUFFER:
		id := RenderbufferId(a.Object)
		r := c.Instances.Renderbuffers[id]
		return uint32(r.Width), uint32(r.Height), nil
	default:
		return 0, 0, fmt.Errorf("Unknown framebuffer attachment type %T", a.Type)
	}
}
