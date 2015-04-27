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
	"android.googlesource.com/platform/tools/gpu/gfxapi/state"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/service"
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
	out           chan gfxapi.Image
}

// colorBufferRequest requests a postback of the framebuffer's depth attachment.
type depthBufferRequest struct {
	after atom.ID
	out   chan gfxapi.Image
}

// timeCallsRequest requests a postback of atom timing information.
type timeCallsRequest struct {
	out  chan gfxapi.CallTiming
	mask service.TimingMask
}

func (a api) ReplayTransforms(
	ctx replay.Context,
	config replay.Config,
	requests []replay.Request,
	postback replay.Postback,
	device *service.Device,
	db database.Database,
	logger log.Logger) atom.Transforms {

	transforms := atom.Transforms{
		// Pre-filter all atoms to the given context.
		transform.ContextFilter(ctx.ContextID),
	}

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
			injector.Inject(
				req.after,
				postback(func(data interface{}, err error) {
					if err == nil {
						req.out <- gfxapi.Image{Data: data.([]byte)}
					} else {
						req.out <- gfxapi.Image{Error: err}
					}
				}), readFramebufferColor{
					contextID: ctx.ContextID,
					width:     req.width,
					height:    req.height,
				},
			)

		case depthBufferRequest:
			earlyTerminator.Add(req.after)
			skipDrawCalls.Draw(req.after)
			injector.Inject(
				req.after,
				postback(func(data interface{}, err error) {
					if err == nil {
						req.out <- gfxapi.Image{Data: data.([]byte)}
					} else {
						req.out <- gfxapi.Image{Error: err}
					}
				}), readFramebufferDepth{
					contextID: ctx.ContextID,
					database:  db,
				},
			)

		case timeCallsRequest:
			profiling = true
			transforms.Add(&timingInfoTransform{
				postback:     postback,
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

	// TODO: These features should be testing capabilites returned by extensions.
	if device.RequiresShaderPatching {
		transforms.Add(
			precisionStrip(),
			halfFloatOESToHalfFloatARB(),
			decompressTextures(ctx.CaptureID, db, logger),
		)
	}

	if c, ok := config.(drawConfig); ok && c.wireframe {
		transforms.Add(wireframe(db, logger))
	}

	// Cleanup
	transforms.Add(&destroyResourcesAtEOS{state: state.New()})

	return transforms
}

func (a api) ColorBuffer(ctx *replay.Context, mgr *replay.Manager, after atom.ID, width, height uint32, wireframe bool) <-chan gfxapi.Image {
	out := make(chan gfxapi.Image, 1)
	c := drawConfig{wireframe: wireframe}
	r := colorBufferRequest{after: after, width: width, height: height, out: out}
	if err := mgr.Replay(ctx, c, r, a); err != nil {
		out <- gfxapi.Image{Error: err}
	}
	return out
}

func (a api) DepthBuffer(ctx *replay.Context, mgr *replay.Manager, after atom.ID) <-chan gfxapi.Image {
	out := make(chan gfxapi.Image, 1)
	c := drawConfig{}
	r := depthBufferRequest{after: after, out: out}
	if err := mgr.Replay(ctx, c, r, a); err != nil {
		out <- gfxapi.Image{Error: err}
	}
	return out
}

func (a api) TimeCalls(ctx *replay.Context, mgr *replay.Manager, mask service.TimingMask) <-chan gfxapi.CallTiming {
	out := make(chan gfxapi.CallTiming, 1)
	c := uniqueConfig()
	r := timeCallsRequest{mask: mask, out: out}
	if err := mgr.Replay(ctx, c, r, a); err != nil {
		out <- gfxapi.CallTiming{Error: err}
	}
	return out
}

// halfFloatOESToHalfFloatARB returns a transform that converts all
// vertex streams declared of type GL_HALF_FLOAT_OES to GL_HALF_FLOAT_ARB.
func halfFloatOESToHalfFloatARB() atom.Transformer {
	// https://www.opengl.org/registry/specs/ARB/half_float_pixel.txt
	const GL_HALF_FLOAT_ARB = 0x140B

	return atom.Transform("HalfFloatOESToHalfFloatARB", func(id atom.ID, a atom.Atom, out atom.Writer) {
		if cmd, ok := a.(*GlVertexAttribPointer); ok &&
			cmd.In.Type == VertexAttribType_GL_HALF_FLOAT_OES {
			out.Write(id, &GlVertexAttribPointer{
				In: GlVertexAttribPointer_In{
					Location:   cmd.In.Location,
					Size:       cmd.In.Size,
					Type:       GL_HALF_FLOAT_ARB,
					Normalized: cmd.In.Normalized,
					Stride:     cmd.In.Stride,
					Data:       cmd.In.Data,
				},
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
	state *state.State
}

func (t *destroyResourcesAtEOS) Transform(id atom.ID, a atom.Atom, out atom.Writer) {
	if m, ok := a.(state.Mutator); ok {
		m.Mutate(t.state)
	}
	out.Write(id, a)
}

func (t *destroyResourcesAtEOS) Flush(out atom.Writer) {
	id := atom.NoID
	for cid, state := range t.state.Contexts {
		s, ok := state.(*State)
		if !ok {
			return
		}

		// Delete all Renderbuffers.
		renderbuffers := RenderbufferIdArray{}
		for renderbufferId := range s.Instances.Renderbuffers {
			// Skip virtual renderbuffers: backbuffer_color(-1), backbuffer_depth(-2), backbuffer_stencil(-3).
			if renderbufferId < 0xf0000000 {
				renderbuffers = append(renderbuffers, renderbufferId)
			}
		}
		if len(renderbuffers) > 0 {
			out.Write(id, NewGlDeleteRenderbuffers(cid, int32(len(renderbuffers)), renderbuffers))
		}

		// Delete all Textures.
		textures := TextureIdArray{}
		for textureId := range s.Instances.Textures {
			textures = append(textures, textureId)
		}
		if len(textures) > 0 {
			out.Write(id, NewGlDeleteTextures(cid, int32(len(textures)), textures))
		}

		// Delete all Framebuffers.
		framebuffers := FramebufferIdArray{}
		for framebufferId := range s.Instances.Framebuffers {
			framebuffers = append(framebuffers, framebufferId)
		}
		if len(framebuffers) > 0 {
			out.Write(id, NewGlDeleteFramebuffers(cid, int32(len(framebuffers)), framebuffers))
		}

		// Delete all Buffers.
		buffers := BufferIdArray{}
		for bufferId := range s.Instances.Buffers {
			buffers = append(buffers, bufferId)
		}
		if len(buffers) > 0 {
			out.Write(id, NewGlDeleteBuffers(cid, int32(len(buffers)), buffers))
		}

		// Delete all Shaders.
		for shaderId := range s.Instances.Shaders {
			out.Write(id, NewGlDeleteShader(cid, shaderId))
		}

		// Delete all Programs.
		for programId := range s.Instances.Programs {
			out.Write(id, NewGlDeleteProgram(cid, programId))
		}

		// Delete all VertexArrays.
		vertexArrays := VertexArrayIdArray{}
		for vertexArrayId := range s.Instances.VertexArrays {
			vertexArrays = append(vertexArrays, vertexArrayId)
		}
		if len(vertexArrays) > 0 {
			out.Write(id, NewGlDeleteVertexArraysOES(cid, int32(len(vertexArrays)), vertexArrays))
		}

		// Delete all SyncObjects. TODO: Uncomment when added to API file.
		// for syncObjectId := range s.Instances.SyncObjects {
		// 	out.Write(id, NewGlDeleteSync(cid, syncObjectId))
		// }
	}
}
