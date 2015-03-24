package gles

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/gfxapi/state"
)

type State struct {
	Globals
	ValidateOutput bool
}

func (s *State) GetFramebufferAttachmentSize(att state.FramebufferAttachment) (width, height uint32, err error) {
	framebufferID := s.BoundFramebuffers[FramebufferTarget_GL_READ_FRAMEBUFFER]

	framebuffer, ok := s.Instances.Framebuffers[framebufferID]
	if !ok {
		return 0, 0, fmt.Errorf("No GL_FRAMEBUFFER bound")
	}

	var attachment FramebufferAttachment
	switch att {
	case state.FramebufferAttachmentColor:
		attachment = FramebufferAttachment_GL_COLOR_ATTACHMENT0
	case state.FramebufferAttachmentDepth:
		attachment = FramebufferAttachment_GL_DEPTH_ATTACHMENT
	case state.FramebufferAttachmentStencil:
		attachment = FramebufferAttachment_GL_STENCIL_ATTACHMENT
	default:
		return 0, 0, fmt.Errorf("Framebuffer attachment %v unsupported by gles", att)
	}

	a, ok := framebuffer.Attachments[attachment]
	if !ok {
		return 0, 0, fmt.Errorf("%s is not bound", attachment)
	}

	switch a.Type {
	case FramebufferAttachmentType_GL_TEXTURE:
		id := TextureId(a.Object)
		t := s.Instances.Textures[id]
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
	case FramebufferAttachmentType_GL_RENDERBUFFER:
		id := RenderbufferId(a.Object)
		r := s.Instances.Renderbuffers[id]
		return uint32(r.Width), uint32(r.Height), nil
	default:
		return 0, 0, fmt.Errorf("Unknown framebuffer attachment type %T", a.Type)
	}
}
