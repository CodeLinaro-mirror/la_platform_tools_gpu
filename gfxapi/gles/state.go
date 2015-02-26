package gles

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/memory"
)

type state struct {
	Globals
	Mem memory.Memory
}

func initialState() *state {
	s := state{}
	s.Init()
	return &s
}

func (s *state) Memory() *memory.Memory {
	return &s.Mem
}

func (s *state) GetFramebufferAttachmentSize(att gfxapi.FramebufferAttachment) (width, height uint32, err error) {
	framebufferID := s.BoundFramebuffers.GetOrError(FramebufferTarget_GL_FRAMEBUFFER)
	if !s.Instances.Framebuffers.Contains(framebufferID) {
		return 0, 0, fmt.Errorf("No GL_FRAMEBUFFER bound")
	}
	framebuffer := s.Instances.Framebuffers.Get(framebufferID, nil)

	var attachment FramebufferAttachment
	attachment, ok := map[gfxapi.FramebufferAttachment]FramebufferAttachment{
		gfxapi.FramebufferAttachmentColor:   FramebufferAttachment_GL_COLOR_ATTACHMENT0,
		gfxapi.FramebufferAttachmentDepth:   FramebufferAttachment_GL_DEPTH_ATTACHMENT,
		gfxapi.FramebufferAttachmentStencil: FramebufferAttachment_GL_STENCIL_ATTACHMENT,
	}[att]
	if !ok {
		return 0, 0, fmt.Errorf("Framebuffer attachment %v unsupported by gles", att)
	}

	if !framebuffer.Attachments.Contains(attachment) {
		return 0, 0, fmt.Errorf("%s is not bound", attachment)
	}

	a := framebuffer.Attachments.GetOrError(attachment)
	switch a.Type {
	case FramebufferAttachmentType_GL_TEXTURE:
		id := TextureId(a.Object)
		t := s.Instances.Textures.GetOrError(id)
		switch t.Kind {
		case TextureKind_TEXTURE2D:
			l := t.Texture2D.GetOrError(a.TextureLevel)
			return uint32(l.Width), uint32(l.Height), nil
		case TextureKind_CUBEMAP:
			l := t.Cubemap.GetOrError(a.TextureLevel)
			f := l.Faces.GetOrError(a.CubeMapFace)
			return uint32(f.Width), uint32(f.Height), nil
		default:
			return 0, 0, fmt.Errorf("Unknown texture kind %v", t.Kind)
		}
	case FramebufferAttachmentType_GL_RENDERBUFFER:
		id := RenderbufferId(a.Object)
		r := s.Instances.Renderbuffers.GetOrError(id)
		return uint32(r.Width), uint32(r.Height), nil
	default:
		return 0, 0, fmt.Errorf("Unknown framebuffer attachment type %T", a.Type)
	}
}
