package gles

import (
	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/replay/builder"
	"android.googlesource.com/platform/tools/gpu/replay/value"
)

func (i BufferId) remap(a atom.Atom, s *state) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = s.Instances.Buffers[i]
	}
	return
}

func (i FramebufferId) remap(a atom.Atom, s *state) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = s.Instances.Framebuffers[i]
	}
	return
}

func (i RenderbufferId) remap(a atom.Atom, s *state) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = s.Instances.Renderbuffers[i]
	}
	return
}

func (i ProgramId) remap(a atom.Atom, s *state) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = s.Instances.Programs[i]
	}
	return
}

func (i ShaderId) remap(a atom.Atom, s *state) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = s.Instances.Shaders[i]
	}
	return
}

func (i TextureId) remap(a atom.Atom, s *state) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = s.Instances.Textures[i]
	}
	return
}

func (i VertexArrayId) remap(a atom.Atom, s *state) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = s.Instances.VertexArrays[i]
	}
	return
}

func (i QueryId) remap(a atom.Atom, s *state) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = s.Instances.Queries[i]
	}
	return
}

func (i UniformLocation) remap(a atom.Atom, s *state) (key interface{}, remap bool) {
	program := s.BoundProgram
	switch a := a.(type) {
	case *GlGetActiveUniform:
		program = a.In.Program
	case *GlGetUniformLocation:
		program = a.In.Program
	}
	return struct {
		p *Program
		l UniformLocation
	}{
		s.Instances.Programs[program], i,
	}, true
}

func (i IndicesPointer) value(b *builder.Builder, a atom.Atom, s *state) value.Value {
	if s.BoundBuffers[BufferTarget_GL_ELEMENT_ARRAY_BUFFER] != 0 {
		return value.AbsolutePointer(i)
	} else {
		return value.VolatileCapturePointer(i)
	}
}

func (i VertexPointer) value(b *builder.Builder, a atom.Atom, s *state) value.Value {
	if s.BoundBuffers[BufferTarget_GL_ARRAY_BUFFER] != 0 {
		return value.AbsolutePointer(i)
	} else {
		return value.VolatileCapturePointer(i)
	}
}

func (i TexturePointer) value(b *builder.Builder, a atom.Atom, s *state) value.Value {
	if i == 0 {
		return value.AbsolutePointer(i)
	} else {
		return value.VolatileCapturePointer(i)
	}
}

func (i BufferDataPointer) value(b *builder.Builder, a atom.Atom, s *state) value.Value {
	if i == 0 {
		return value.AbsolutePointer(i)
	} else {
		return value.VolatileCapturePointer(i)
	}
}

func (i ImageOES) value(b *builder.Builder, a atom.Atom, s *state) value.Value {
	return value.AbsolutePointer(i)
}

// AttributeLocations cannot be remapped like UniformLocations as the VertexAttributeArrays are
// shared between different programs. Instead, simply force the location to match what was recorded
// in the capture using glBindAttribLocation.
// TODO: This implementation currently calls glLinkProgram for every call to glGetAttribLocation!
//       This is obviously not ideal, and we should be doing this once at glLinkProgram once the
//       spy emits location hinting information.
func (ω *GlGetAttribLocation) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {
	if ω.Out.Result >= 0 {
		NewGlBindAttribLocation(ω.In.Program, ω.Out.Result, ω.In.Name).replay(id, s, b, false)
		NewGlLinkProgram(ω.In.Program).replay(id, s, b, false)
	}
}

// These are not called in replay
func (ω *EglCreateContext) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool) {}
func (ω *EglMakeCurrent) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool)   {}
func (ω *EglSwapBuffers) replay(id atom.ID, s *state, b *builder.Builder, wantOutput bool)   {}
