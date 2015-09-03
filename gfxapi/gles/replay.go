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
	"android.googlesource.com/platform/tools/gpu/atom/transform"
	"android.googlesource.com/platform/tools/gpu/config"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/service"
	"android.googlesource.com/platform/tools/gpu/service/path"
)

var (
	// Interface compliance tests
	_ = replay.QueryColorBuffer(api{})
	_ = replay.QueryDepthBuffer(api{})
	_ = replay.QueryCallDurations(api{})
)

// drawConfig is a replay Config used by colorBufferRequest and
// depthBufferRequests.
type drawConfig struct {
	wireframeMode      replay.WireframeMode
	wireframeOverlayID atom.ID // used when wireframeMode == WireframeOverlay
}

// uniqueConfig returns a replay.Config that is guaranteed to be unique.
// Any requests made with a Config returned from uniqueConfig will not be
// batched with any other request.
func uniqueConfig() replay.Config {
	return &struct{}{}
}

// colorBufferRequest requests a postback of the framebuffer's color attachment.
type colorBufferRequest struct {
	after            atom.ID
	width, height    uint32
	out              chan replay.Image
	wireframeOverlay bool
}

// colorBufferRequest requests a postback of the framebuffer's depth attachment.
type depthBufferRequest struct {
	after atom.ID
	out   chan replay.Image
}

// timeCallsRequest requests a postback of atom timing information.
type timeCallsRequest struct {
	out   chan replay.CallTiming
	flags service.TimingFlags
}

func (a api) Replay(
	ctx replay.Context,
	cfg replay.Config,
	requests []replay.Request,
	device *service.Device,
	atoms atom.List,
	out atom.Writer,
	d database.Database,
	l log.Logger) error {

	transforms := atom.Transforms{}

	// Terminate after all atoms of interest.
	earlyTerminator := &transform.EarlyTerminator{}

	// Skip unnecessary draw calls.
	skipDrawCalls := &transform.SkipDrawCalls{}

	// Injector of new atoms.
	injector := &transform.Injector{}

	// Transform for all framebuffer reads.
	readFramebuffer := NewReadFramebuffer(d, l)

	profiling := false

	for _, req := range requests {
		switch req := req.(type) {
		case colorBufferRequest:
			earlyTerminator.Add(req.after)
			skipDrawCalls.Draw(req.after)
			readFramebuffer.Color(req.after, req.width, req.height, req.out)

			cfg := cfg.(drawConfig)
			switch cfg.wireframeMode {
			case replay.AllWireframe:
				transforms.Add(wireframe(d, l))
			case replay.WireframeOverlay:
				transforms.Add(wireframeOverlay(req.after, d, l))
			}

		case depthBufferRequest:
			earlyTerminator.Add(req.after)
			skipDrawCalls.Draw(req.after)
			readFramebuffer.Depth(req.after, device, req.out)

		case timeCallsRequest:
			profiling = true
			transforms.Add(timingInfo(req.flags, req.out, device, d, l))
		}
	}

	if !profiling {
		// Not profiling. Add optimisation transforms.
		transforms.Add(earlyTerminator)

		// Check to see if any contexts use the 'preserveBuffersOnSwap' flag.
		// If it is used, we can't skip draw calls or display the
		// undefined-framebuffer pattern.
		preserveBuffersOnSwap := false
		for _, a := range atoms.Atoms {
			if b, ok := a.(*ContextInfo); ok && b.PreserveBuffersOnSwap {
				preserveBuffersOnSwap = true
				break
			}
		}

		if !preserveBuffersOnSwap {
			transforms.Add(skipDrawCalls, undefinedFramebuffer(d, l))
		}
	}

	transforms.Add(
		readFramebuffer,
		injector,
		remapAttributes(),
	)

	// Device-dependent transforms.
	transforms.Add(
		decompressTextures(device, &path.Capture{ID: ctx.Capture}, d, l),
		halfFloatOESToHalfFloatARB(device))

	if c, err := compat(device, d, l); err == nil {
		transforms.Add(c)
	} else {
		log.E(l, "Failed to create compatability transform: %v", err)
	}

	// Cleanup
	transforms.Add(&destroyResourcesAtEOS{
		state:  gfxapi.NewState(),
		db:     d,
		logger: l,
	})

	if config.DebugReplay {
		log.Infof(l, "Replaying %d atoms using transform chain:", len(atoms.Atoms))
		for i, t := range transforms {
			log.Infof(l, "(%d) %#v", i, t)
		}
	}

	transforms.Transform(atoms, out)

	return nil
}

func (a api) QueryColorBuffer(
	ctx replay.Context,
	mgr *replay.Manager,
	after atom.ID,
	width, height uint32,
	wireframeMode replay.WireframeMode) <-chan replay.Image {

	out := make(chan replay.Image, 1)
	c := drawConfig{wireframeMode: wireframeMode}
	if wireframeMode == replay.WireframeOverlay {
		c.wireframeOverlayID = after
	}
	r := colorBufferRequest{after: after, width: width, height: height, out: out}
	if err := mgr.Replay(ctx, c, r, a); err != nil {
		out <- replay.Image{Error: err}
	}
	return out
}

func (a api) QueryDepthBuffer(ctx replay.Context, mgr *replay.Manager, after atom.ID) <-chan replay.Image {
	out := make(chan replay.Image, 1)
	c := drawConfig{}
	r := depthBufferRequest{after: after, out: out}
	if err := mgr.Replay(ctx, c, r, a); err != nil {
		out <- replay.Image{Error: err}
	}
	return out
}

func (a api) QueryCallDurations(ctx replay.Context, mgr *replay.Manager, flags service.TimingFlags) <-chan replay.CallTiming {
	out := make(chan replay.CallTiming, 1)
	c := uniqueConfig()
	r := timeCallsRequest{flags: flags, out: out}
	if err := mgr.Replay(ctx, c, r, a); err != nil {
		out <- replay.CallTiming{Error: err}
	}
	return out
}

// halfFloatOESToHalfFloatARB returns a transform that converts all vertex streams
// declared of type GL_HALF_FLOAT_OES to GL_HALF_FLOAT_ARB, if unsupported by the target device.
func halfFloatOESToHalfFloatARB(device *service.Device) atom.Transformer {
	if v, err := ParseVersion(device.Version); err == nil {
		if v.IsES && device.HasExtension("GL_OES_vertex_half_float") {
			return nil
		}
	}
	// TODO: fallback to full GL_FLOAT unpacking if GL_ARB_half_float_vertex isn't supported.
	return atom.Transform("HalfFloatOESToHalfFloatARB", func(id atom.ID, a atom.Atom, out atom.Writer) {
		if cmd, ok := a.(*GlVertexAttribPointer); ok &&
			cmd.Type == GLenum_GL_HALF_FLOAT_OES {
			out.Write(id, &GlVertexAttribPointer{
				Location:   cmd.Location,
				Size:       cmd.Size,
				Type:       GLenum_GL_HALF_FLOAT_ARB,
				Normalized: cmd.Normalized,
				Stride:     cmd.Stride,
				Data:       cmd.Data,
			})
		} else {
			out.Write(id, a)
		}
	})
}

// destroyResourcesAtEOS is a transform that destroys all textures,
// framebuffers, buffers, shaders, programs and vertex-arrays that were not
// destroyed by EOS.
type destroyResourcesAtEOS struct {
	state  *gfxapi.State
	db     database.Database
	logger log.Logger
}

func (t *destroyResourcesAtEOS) Transform(id atom.ID, a atom.Atom, out atom.Writer) {
	a.Mutate(t.state, t.db, t.logger)
	out.Write(id, a)
}

func (t *destroyResourcesAtEOS) Flush(out atom.Writer) {
	id := atom.NoID
	c := getContext(t.state)
	if c == nil {
		return
	}

	a, d, l := t.state.Architecture, t.db, t.logger

	// Delete all Renderbuffers.
	renderbuffers := make([]RenderbufferId, 0, len(c.Instances.Renderbuffers)-3)
	for renderbufferId := range c.Instances.Renderbuffers {
		// Skip virtual renderbuffers: backbuffer_color(-1), backbuffer_depth(-2), backbuffer_stencil(-3).
		if renderbufferId < 0xf0000000 {
			renderbuffers = append(renderbuffers, renderbufferId)
		}
	}
	if len(renderbuffers) > 0 {
		out.Write(id,
			NewGlDeleteRenderbuffers(GLsizei(len(renderbuffers)), memory.Tmp).
				AddRead(atom.Data(a, d, l, memory.Tmp, renderbuffers)))
	}

	// Delete all Textures.
	textures := make([]TextureId, 0, len(c.Instances.Textures))
	for textureId := range c.Instances.Textures {
		textures = append(textures, textureId)
	}
	if len(textures) > 0 {
		out.Write(id,
			NewGlDeleteTextures(GLsizei(len(textures)), memory.Tmp).
				AddRead(atom.Data(a, d, l, memory.Tmp, textures)))
	}

	// Delete all Framebuffers.
	framebuffers := make([]FramebufferId, 0, len(c.Instances.Framebuffers))
	for framebufferId := range c.Instances.Framebuffers {
		framebuffers = append(framebuffers, framebufferId)
	}
	if len(framebuffers) > 0 {
		out.Write(id,
			NewGlDeleteFramebuffers(GLsizei(len(framebuffers)), memory.Tmp).
				AddRead(atom.Data(a, d, l, memory.Tmp, framebuffers)))
	}

	// Delete all Buffers.
	buffers := make([]BufferId, 0, len(c.Instances.Buffers))
	for bufferId := range c.Instances.Buffers {
		buffers = append(buffers, bufferId)
	}
	if len(buffers) > 0 {
		out.Write(id,
			NewGlDeleteBuffers(GLsizei(len(buffers)), memory.Tmp).
				AddRead(atom.Data(a, d, l, memory.Tmp, buffers)))
	}

	// Delete all VertexArrays.
	vertexArrays := make([]VertexArrayId, 0, len(c.Instances.VertexArrays))
	for vertexArrayId := range c.Instances.VertexArrays {
		vertexArrays = append(vertexArrays, vertexArrayId)
	}
	if len(vertexArrays) > 0 {
		out.Write(id,
			NewGlDeleteVertexArraysOES(GLsizei(len(vertexArrays)), memory.Tmp).
				AddRead(atom.Data(a, d, l, memory.Tmp, vertexArrays)))
	}

	// Delete all Shaders.
	for shaderId := range c.Instances.Shaders {
		out.Write(id, NewGlDeleteShader(shaderId))
	}

	// Delete all Programs.
	for programId := range c.Instances.Programs {
		out.Write(id, NewGlDeleteProgram(programId))
	}

	// Delete all Queries.
	queries := make([]QueryId, 0, len(c.Instances.Queries))
	for queryId := range c.Instances.Queries {
		queries = append(queries, queryId)
	}
	if len(queries) > 0 {
		out.Write(id,
			NewGlDeleteQueries(GLsizei(len(queries)), memory.Tmp).
				AddRead(atom.Data(a, d, l, memory.Tmp, queries)))
	}
}
