package gles

import (
	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/atom/transform"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/replay/builder"
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

func (a api) ReplayWriter(b *builder.Builder) replay.Writer {
	return newReplayWriter(b)
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
						req.out <- gfxapi.Image{Data: data.([]byte)}
					} else {
						req.out <- gfxapi.Image{Error: err}
					}
				}), readFramebufferDepth{},
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
			uncompressTextures(ctx.CaptureID, db, logger),
		)
	}

	if c, ok := config.(drawConfig); ok && c.wireframe {
		transforms.Add(wireframe(db, logger))
	}

	// Cleanup
	transforms.Add(destroyResourcesAtEOS())

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
func halfFloatOESToHalfFloatARB() atom.Transform {
	// https://www.opengl.org/registry/specs/ARB/half_float_pixel.txt
	const GL_HALF_FLOAT_ARB = 0x140B

	return func(id atom.ID, a atom.Atom, out atom.Writer) {
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
	}
}

// destroyResourcesAtEOS returns a transform that destroys all textures,
// framebuffers, buffers, shaders, programs and vertex-arrays that were not
// destroyed by EOS.
func destroyResourcesAtEOS() atom.Transform {
	mutator := &StateMutator{State: initialState()}

	return func(id atom.ID, a atom.Atom, out atom.Writer) {
		switch a.(type) {
		default:
			mutator.Write(id, a)
		case *atom.EOS:
			state := mutator.State

			// Delete all Renderbuffers.
			renderbuffers := RenderbufferIdArray{}
			for renderbufferId := range state.Instances.Renderbuffers {
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
			for textureId := range state.Instances.Textures {
				textures = append(textures, textureId)
			}
			if len(textures) > 0 {
				out.Write(id, NewGlDeleteTextures(int32(len(textures)), textures))
			}

			// Delete all Framebuffers.
			framebuffers := FramebufferIdArray{}
			for framebufferId := range state.Instances.Framebuffers {
				framebuffers = append(framebuffers, framebufferId)
			}
			if len(framebuffers) > 0 {
				out.Write(id, NewGlDeleteFramebuffers(int32(len(framebuffers)), framebuffers))
			}

			// Delete all Buffers.
			buffers := BufferIdArray{}
			for bufferId := range state.Instances.Buffers {
				buffers = append(buffers, bufferId)
			}
			if len(buffers) > 0 {
				out.Write(id, NewGlDeleteBuffers(int32(len(buffers)), buffers))
			}

			// Delete all Shaders.
			for shaderId := range state.Instances.Shaders {
				out.Write(id, NewGlDeleteShader(shaderId))
			}

			// Delete all Programs.
			for programId := range state.Instances.Programs {
				out.Write(id, NewGlDeleteProgram(programId))
			}

			// Delete all VertexArrays.
			vertexArrays := VertexArrayIdArray{}
			for vertexArrayId := range state.Instances.VertexArrays {
				vertexArrays = append(vertexArrays, vertexArrayId)
			}
			if len(vertexArrays) > 0 {
				out.Write(id, NewGlDeleteVertexArraysOES(int32(len(vertexArrays)), vertexArrays))
			}

			// Delete all SyncObjects. TODO: Uncomment when added to API file.
			// for syncObjectId := range state.Instances.SyncObjects {
			// 	out.Write(id, NewGlDeleteSync(syncObjectId))
			// }
		}
		out.Write(id, a)
	}
}
