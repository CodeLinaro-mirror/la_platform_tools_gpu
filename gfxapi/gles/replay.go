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
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/service"
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
	wireframe bool
}

// uniqueConfig returns a replay.Config that is guaranteed to be unique.
// Any requests made with a Config returned from uniqueConfig will not be
// batched with any other request.
func uniqueConfig() replay.Config {
	return &struct{}{}
}

// colorBufferRequest requests a postback of the framebuffer's color attachment.
type colorBufferRequest struct {
	after         atom.ID
	width, height uint32
	out           chan replay.Image
}

// colorBufferRequest requests a postback of the framebuffer's depth attachment.
type depthBufferRequest struct {
	after atom.ID
	out   chan replay.Image
}

// timeCallsRequest requests a postback of atom timing information.
type timeCallsRequest struct {
	out  chan replay.CallTiming
	mask service.TimingMask
}

func (a api) ReplayTransforms(
	ctx replay.Context,
	config replay.Config,
	requests []replay.Request,
	device *service.Device,
	db database.Database,
	logger log.Logger) atom.Transforms {

	transforms := atom.Transforms{}

	// Terminate after all atoms of interest.
	earlyTerminator := &transform.EarlyTerminator{}

	// Skip unnecessary draw calls.
	skipDrawCalls := &transform.SkipDrawCalls{}

	// Injector of new atoms.
	injector := &transform.Injector{}

	profiling := false

	for _, req := range requests {
		switch req := req.(type) {
		case colorBufferRequest:
			earlyTerminator.Add(req.after)
			skipDrawCalls.Draw(req.after)
			injector.Inject(req.after, readFramebufferColor(req.width, req.height, req.out))

		case depthBufferRequest:
			earlyTerminator.Add(req.after)
			skipDrawCalls.Draw(req.after)
			injector.Inject(req.after, readFramebufferDepth(req.out))

		case timeCallsRequest:
			profiling = true
			transforms.Add(&timingInfoTransform{
				out:          req.out,
				perCommand:   (req.mask & service.TimingMaskTimingPerCommand) != 0,
				perDrawCall:  (req.mask & service.TimingMaskTimingPerDrawCall) != 0,
				perFrame:     (req.mask & service.TimingMaskTimingPerFrame) != 0,
				timerStartId: make(map[uint8]atom.ID),
			})
		}
	}

	if !profiling {
		// Not profiling. Add optimisation transforms.
		transforms.Add(earlyTerminator, skipDrawCalls)
	}

	transforms.Add(injector)

	// Device-dependent transforms.
	transforms.Add(
		decompressTextures(device, ctx.CaptureID, db, logger),
		precisionStrip(device, db, logger),
		halfFloatOESToHalfFloatARB(device))

	if c, ok := config.(drawConfig); ok && c.wireframe {
		transforms.Add(wireframe(db, logger))
	}

	// Cleanup
	transforms.Add(&destroyResourcesAtEOS{
		state:  gfxapi.NewState(),
		db:     db,
		logger: logger,
	})

	return transforms
}

func (a api) QueryColorBuffer(ctx *replay.Context, mgr *replay.Manager, after atom.ID, width, height uint32, wireframe bool) <-chan replay.Image {
	out := make(chan replay.Image, 1)
	c := drawConfig{wireframe: wireframe}
	r := colorBufferRequest{after: after, width: width, height: height, out: out}
	if err := mgr.Replay(ctx, c, r, a); err != nil {
		out <- replay.Image{Error: err}
	}
	return out
}

func (a api) QueryDepthBuffer(ctx *replay.Context, mgr *replay.Manager, after atom.ID) <-chan replay.Image {
	out := make(chan replay.Image, 1)
	c := drawConfig{}
	r := depthBufferRequest{after: after, out: out}
	if err := mgr.Replay(ctx, c, r, a); err != nil {
		out <- replay.Image{Error: err}
	}
	return out
}

func (a api) QueryCallDurations(ctx *replay.Context, mgr *replay.Manager, mask service.TimingMask) <-chan replay.CallTiming {
	out := make(chan replay.CallTiming, 1)
	c := uniqueConfig()
	r := timeCallsRequest{mask: mask, out: out}
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
			cmd.Type == VertexAttribType_GL_HALF_FLOAT_OES {
			out.Write(id, &GlVertexAttribPointer{
				Location:   cmd.Location,
				Size:       cmd.Size,
				Type:       VertexAttribType_GL_HALF_FLOAT_ARB,
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
	renderbuffers := []RenderbufferId{}
	for renderbufferId := range c.Instances.Renderbuffers {
		// Skip virtual renderbuffers: backbuffer_color(-1), backbuffer_depth(-2), backbuffer_stencil(-3).
		if renderbufferId < 0xf0000000 {
			renderbuffers = append(renderbuffers, renderbufferId)
		}
	}
	if len(renderbuffers) > 0 {
		out.Write(id,
			NewGlDeleteRenderbuffers(int32(len(renderbuffers)), memory.Tmp.Base).
				AddRead(atom.Data(a, d, l, memory.Tmp.Base, renderbuffers)))
	}

	// Delete all Textures.
	textures := []TextureId{}
	for textureId := range c.Instances.Textures {
		textures = append(textures, textureId)
	}
	if len(textures) > 0 {
		out.Write(id,
			NewGlDeleteTextures(int32(len(textures)), memory.Tmp.Base).
				AddRead(atom.Data(a, d, l, memory.Tmp.Base, textures)))
	}

	// Delete all Framebuffers.
	framebuffers := []FramebufferId{}
	for framebufferId := range c.Instances.Framebuffers {
		framebuffers = append(framebuffers, framebufferId)
	}
	if len(framebuffers) > 0 {
		out.Write(id,
			NewGlDeleteFramebuffers(int32(len(framebuffers)), memory.Tmp.Base).
				AddRead(atom.Data(a, d, l, memory.Tmp.Base, framebuffers)))
	}

	// Delete all Buffers.
	buffers := []BufferId{}
	for bufferId := range c.Instances.Buffers {
		buffers = append(buffers, bufferId)
	}
	if len(buffers) > 0 {
		out.Write(id,
			NewGlDeleteBuffers(int32(len(buffers)), memory.Tmp.Base).
				AddRead(atom.Data(a, d, l, memory.Tmp.Base, buffers)))
	}

	// Delete all VertexArrays.
	vertexArrays := []VertexArrayId{}
	for vertexArrayId := range c.Instances.VertexArrays {
		vertexArrays = append(vertexArrays, vertexArrayId)
	}
	if len(vertexArrays) > 0 {
		out.Write(id,
			NewGlDeleteVertexArraysOES(int32(len(vertexArrays)), memory.Tmp.Base).
				AddRead(atom.Data(a, d, l, memory.Tmp.Base, vertexArrays)))
	}

	// Delete all Shaders.
	for shaderId := range c.Instances.Shaders {
		out.Write(id, NewGlDeleteShader(shaderId))
	}

	// Delete all Programs.
	for programId := range c.Instances.Programs {
		out.Write(id, NewGlDeleteProgram(programId))
	}
}
