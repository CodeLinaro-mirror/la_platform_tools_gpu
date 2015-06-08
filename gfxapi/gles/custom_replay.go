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
	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/replay/builder"
	"android.googlesource.com/platform/tools/gpu/replay/value"
)

func (i BufferId) remap(a atom.Atom, s *gfxapi.State) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = getContext(s).Instances.Buffers[i]
	}
	return
}

func (i FramebufferId) remap(a atom.Atom, s *gfxapi.State) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = getContext(s).Instances.Framebuffers[i]
	}
	return
}

func (i RenderbufferId) remap(a atom.Atom, s *gfxapi.State) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = getContext(s).Instances.Renderbuffers[i]
	}
	return
}

func (i ProgramId) remap(a atom.Atom, s *gfxapi.State) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = getContext(s).Instances.Programs[i]
	}
	return
}

func (i ShaderId) remap(a atom.Atom, s *gfxapi.State) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = getContext(s).Instances.Shaders[i]
	}
	return
}

func (i TextureId) remap(a atom.Atom, s *gfxapi.State) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = getContext(s).Instances.Textures[i]
	}
	return
}

func (i VertexArrayId) remap(a atom.Atom, s *gfxapi.State) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = getContext(s).Instances.VertexArrays[i]
	}
	return
}

func (i QueryId) remap(a atom.Atom, s *gfxapi.State) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = getContext(s).Instances.Queries[i]
	}
	return
}

func (i UniformLocation) remap(a atom.Atom, s *gfxapi.State) (key interface{}, remap bool) {
	state := getContext(s)
	program := state.BoundProgram
	switch a := a.(type) {
	case *GlGetActiveUniform:
		program = a.Program
	case *GlGetUniformLocation:
		program = a.Program
	}
	return struct {
		p *Program
		l UniformLocation
	}{
		state.Instances.Programs[program], i,
	}, true
}

func (i IndicesPointer) value(b *builder.Builder, a atom.Atom, s *gfxapi.State) value.Value {
	if getContext(s).BoundBuffers[BufferTarget_GL_ELEMENT_ARRAY_BUFFER] != 0 {
		return value.AbsolutePointer(i.Address)
	} else {
		return value.VolatileCapturePointer(i.Address)
	}
}

func (i VertexPointer) value(b *builder.Builder, a atom.Atom, s *gfxapi.State) value.Value {
	if getContext(s).BoundBuffers[BufferTarget_GL_ARRAY_BUFFER] != 0 {
		return value.AbsolutePointer(i.Address)
	} else {
		return value.VolatileCapturePointer(i.Address)
	}
}

func (i TexturePointer) value(b *builder.Builder, a atom.Atom, s *gfxapi.State) value.Value {
	if i.Address == 0 {
		return value.AbsolutePointer(i.Address)
	} else {
		return value.VolatileCapturePointer(i.Address)
	}
}

func (i BufferDataPointer) value(b *builder.Builder, a atom.Atom, s *gfxapi.State) value.Value {
	if i.Address == 0 {
		return value.AbsolutePointer(i.Address)
	} else {
		return value.VolatileCapturePointer(i.Address)
	}
}

func (i ImageOES) value(b *builder.Builder, a atom.Atom, s *gfxapi.State) value.Value {
	return value.AbsolutePointer(i.Address)
}

// AttributeLocations cannot be remapped like UniformLocations as the VertexAttributeArrays are
// shared between different programs. Instead, simply force the location to match what was recorded
// in the capture using glBindAttribLocation.
// TODO: This implementation currently calls glLinkProgram for every call to glGetAttribLocation!
//       This is obviously not ideal, and we should be doing this once at glLinkProgram once the
//       spy emits location hinting information.
func (ω *GlGetAttribLocation) Replay(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error {
	if ω.Result >= 0 {
		NewGlBindAttribLocation(ω.Program, ω.Result, ω.Name).Replay(i, s, d, l, b)
		NewGlLinkProgram(ω.Program).Replay(i, s, d, l, b)
	}
	return nil
}

func (ω *EglCreateContext) Replay(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error {
	ω.Mutate(s, d, l)
	ctxID := uint32(getState(s).EGLContexts[ω.Result].Identifier)
	return NewReplayCreateRenderer(ctxID).Replay(i, s, d, l, b)
}

func (ω *EglMakeCurrent) Replay(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error {
	ω.Mutate(s, d, l)
	if ω.Context.Address == 0 {
		return nil
	}
	ctxID := uint32(getState(s).EGLContexts[ω.Context].Identifier)
	return NewReplayBindRenderer(ctxID).Replay(i, s, d, l, b)
}

func (ω *WglCreateContext) Replay(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error {
	ω.Mutate(s, d, l)
	ctxID := uint32(getState(s).WGLContexts[ω.Result].Identifier)
	return NewReplayCreateRenderer(ctxID).Replay(i, s, d, l, b)
}

func (ω *WglCreateContextAttribsARB) Replay(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error {
	ω.Mutate(s, d, l)
	ctxID := uint32(getState(s).WGLContexts[ω.Result].Identifier)
	return NewReplayCreateRenderer(ctxID).Replay(i, s, d, l, b)
}

func (ω *WglMakeCurrent) Replay(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error {
	ω.Mutate(s, d, l)
	if ω.Hglrc.Address == 0 {
		return nil
	}
	ctxID := uint32(getState(s).WGLContexts[ω.Hglrc].Identifier)
	return NewReplayBindRenderer(ctxID).Replay(i, s, d, l, b)
}

func (ω *CGLCreateContext) Replay(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error {
	ω.Mutate(s, d, l)
	ctxID := uint32(getState(s).CGLContexts[ω.Ctx.Read(s, d, l)].Identifier)
	return NewReplayCreateRenderer(ctxID).Replay(i, s, d, l, b)
}

func (ω *CGLSetCurrentContext) Replay(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error {
	ω.Mutate(s, d, l)
	if ω.Ctx.Address == 0 {
		return nil
	}
	ctxID := uint32(getState(s).CGLContexts[ω.Ctx].Identifier)
	return NewReplayBindRenderer(ctxID).Replay(i, s, d, l, b)
}

func (ω *GlXCreateContext) Replay(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error {
	ω.Mutate(s, d, l)
	ctxID := uint32(getState(s).GLXContexts[ω.Result].Identifier)
	return NewReplayCreateRenderer(ctxID).Replay(i, s, d, l, b)
}

func (ω *GlXCreateNewContext) Replay(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error {
	ω.Mutate(s, d, l)
	ctxID := uint32(getState(s).GLXContexts[ω.Result].Identifier)
	return NewReplayCreateRenderer(ctxID).Replay(i, s, d, l, b)
}

func (ω *GlXMakeContextCurrent) Replay(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error {
	ω.Mutate(s, d, l)
	if ω.Ctx.Address == 0 {
		return nil
	}
	ctxID := uint32(getState(s).GLXContexts[ω.Ctx].Identifier)
	return NewReplayBindRenderer(ctxID).Replay(i, s, d, l, b)
}
