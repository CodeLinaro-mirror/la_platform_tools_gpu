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
	postback replay.Postback,
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
			injector.Inject(
				req.after,
				postback(func(data interface{}, err error) {
					if err == nil {
						req.out <- replay.Image{Data: data.([]byte)}
					} else {
						req.out <- replay.Image{Error: err}
					}
				}), readFramebufferColor{
					width:  req.width,
					height: req.height,
				},
			)

		case depthBufferRequest:
			earlyTerminator.Add(req.after)
			skipDrawCalls.Draw(req.after)
			injector.Inject(
				req.after,
				postback(func(data interface{}, err error) {
					if err == nil {
						req.out <- replay.Image{Data: data.([]byte)}
					} else {
						req.out <- replay.Image{Error: err}
					}
				}), readFramebufferDepth{
					database: db,
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
	transforms.Add(&destroyResourcesAtEOS{state: &gfxapi.State{}})

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

// halfFloatOESToHalfFloatARB returns a transform that converts all
// vertex streams declared of type GL_HALF_FLOAT_OES to GL_HALF_FLOAT_ARB.
func halfFloatOESToHalfFloatARB() atom.Transformer {
	// https://www.opengl.org/registry/specs/ARB/half_float_pixel.txt
	const GL_HALF_FLOAT_ARB = 0x140B

	return atom.Transform("HalfFloatOESToHalfFloatARB", func(id atom.ID, a atom.Atom, out atom.Writer) {
		if cmd, ok := a.(*GlVertexAttribPointer); ok &&
			cmd.Type == VertexAttribType_GL_HALF_FLOAT_OES {
			out.Write(id, &GlVertexAttribPointer{
				Location:   cmd.Location,
				Size:       cmd.Size,
				Type:       GL_HALF_FLOAT_ARB,
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
	state *gfxapi.State
}

func (t *destroyResourcesAtEOS) Transform(id atom.ID, a atom.Atom, out atom.Writer) {
	a.Mutate(t.state)
	out.Write(id, a)
}

func (t *destroyResourcesAtEOS) Flush(out atom.Writer) {
	id := atom.NoID
	c := getContext(t.state)
	if c == nil {
		return
	}

	// Delete all Renderbuffers.
	renderbuffers := RenderbufferIdArray{}
	for renderbufferId := range c.Instances.Renderbuffers {
		// Skip virtual renderbuffers: backbuffer_color(-1), backbuffer_depth(-2), backbuffer_stencil(-3).
		if renderbufferId < 0xf0000000 {
			renderbuffers = append(renderbuffers, renderbufferId)
		}
	}
	if len(renderbuffers) > 0 {
		out.Write(id, NewGlDeleteRenderbuffers(int32(len(renderbuffers)), renderbuffers))
	}

	// Delete all Textures.
	textures := TextureIdArray{}
	for textureId := range c.Instances.Textures {
		textures = append(textures, textureId)
	}
	if len(textures) > 0 {
		out.Write(id, NewGlDeleteTextures(int32(len(textures)), textures))
	}

	// Delete all Framebuffers.
	framebuffers := FramebufferIdArray{}
	for framebufferId := range c.Instances.Framebuffers {
		framebuffers = append(framebuffers, framebufferId)
	}
	if len(framebuffers) > 0 {
		out.Write(id, NewGlDeleteFramebuffers(int32(len(framebuffers)), framebuffers))
	}

	// Delete all Buffers.
	buffers := BufferIdArray{}
	for bufferId := range c.Instances.Buffers {
		buffers = append(buffers, bufferId)
	}
	if len(buffers) > 0 {
		out.Write(id, NewGlDeleteBuffers(int32(len(buffers)), buffers))
	}

	// Delete all Shaders.
	for shaderId := range c.Instances.Shaders {
		out.Write(id, NewGlDeleteShader(shaderId))
	}

	// Delete all Programs.
	for programId := range c.Instances.Programs {
		out.Write(id, NewGlDeleteProgram(programId))
	}

	// Delete all VertexArrays.
	vertexArrays := VertexArrayIdArray{}
	for vertexArrayId := range c.Instances.VertexArrays {
		vertexArrays = append(vertexArrays, vertexArrayId)
	}
	if len(vertexArrays) > 0 {
		out.Write(id, NewGlDeleteVertexArraysOES(int32(len(vertexArrays)), vertexArrays))
	}

	// Delete all SyncObjects. TODO: Uncomment when added to API file.
	// for syncObjectId := range c.Instances.SyncObjects {
	// 	out.Write(id, NewGlDeleteSync(syncObjectId))
	// }
}
