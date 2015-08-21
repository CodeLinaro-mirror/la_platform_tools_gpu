////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

package gles

import (
	"fmt"
	"strings"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/replay/builder"
)

func getState(s *gfxapi.State) *State {
	api := API()
	if state, ok := s.APIs[api].(*State); ok {
		return state
	} else {
		if s.APIs == nil {
			s.APIs = make(map[gfxapi.API]binary.Object)
		}
		state = &State{}
		state.Init()
		s.APIs[api] = state
		return state
	}
}

func (ϟa *GlBlendBarrierKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_0_ext := ExtensionId_GL_KHR_blend_equation_advanced // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_0_ext
	return nil
}
func (ϟa *GlBlendEquationSeparateiEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_1_ext := ExtensionId_GL_EXT_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_1_ext
	return nil
}
func (ϟa *GlBlendEquationiEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_2_ext := ExtensionId_GL_EXT_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_2_ext
	return nil
}
func (ϟa *GlBlendFuncSeparateiEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_3_ext := ExtensionId_GL_EXT_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_3_ext
	return nil
}
func (ϟa *GlBlendFunciEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_4_ext := ExtensionId_GL_EXT_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_4_ext
	return nil
}
func (ϟa *GlColorMaskiEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_5_ext := ExtensionId_GL_EXT_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_5_ext
	return nil
}
func (ϟa *GlCopyImageSubDataEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_6_ext := ExtensionId_GL_EXT_copy_image // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_6_ext
	return nil
}
func (ϟa *GlDebugMessageCallbackKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_7_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_7_ext
	return nil
}
func (ϟa *GlDebugMessageControlKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_8_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_8_ext
	return nil
}
func (ϟa *GlDebugMessageInsertKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_9_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_9_ext
	return nil
}
func (ϟa *GlDisableiEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_10_ext := ExtensionId_GL_EXT_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_10_ext
	return nil
}
func (ϟa *GlEnableiEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_11_ext := ExtensionId_GL_EXT_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_11_ext
	return nil
}
func (ϟa *GlFramebufferTextureEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_12_ext := ExtensionId_GL_EXT_geometry_shader // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_12_ext
	return nil
}
func (ϟa *GlGetDebugMessageLogKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_13_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_13_ext
	return nil
}
func (ϟa *GlGetObjectLabelKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_14_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_14_ext
	return nil
}
func (ϟa *GlGetObjectPtrLabelKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_15_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_15_ext
	return nil
}
func (ϟa *GlGetPointervKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_16_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_16_ext
	return nil
}
func (ϟa *GlGetSamplerParameterIivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_17_ext := ExtensionId_GL_EXT_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_17_ext
	return nil
}
func (ϟa *GlGetSamplerParameterIuivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_18_ext := ExtensionId_GL_EXT_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_18_ext
	return nil
}
func (ϟa *GlGetTexParameterIivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_19_ext := ExtensionId_GL_EXT_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_19_ext
	return nil
}
func (ϟa *GlGetTexParameterIuivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_20_ext := ExtensionId_GL_EXT_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_20_ext
	return nil
}
func (ϟa *GlIsEnablediEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_21_ext := ExtensionId_GL_EXT_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_21_ext
	return nil
}
func (ϟa *GlMinSampleShadingOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_22_ext := ExtensionId_GL_OES_sample_shading // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_22_ext
	return nil
}
func (ϟa *GlObjectLabelKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_23_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_23_ext
	return nil
}
func (ϟa *GlObjectPtrLabelKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_24_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_24_ext
	return nil
}
func (ϟa *GlPatchParameteriEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_25_ext := ExtensionId_GL_EXT_tessellation_shader // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_25_ext
	return nil
}
func (ϟa *GlPopDebugGroupKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_26_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_26_ext
	return nil
}
func (ϟa *GlPrimitiveBoundingBoxEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_27_ext := ExtensionId_GL_EXT_primitive_bounding_box // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_27_ext
	return nil
}
func (ϟa *GlPushDebugGroupKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_28_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_28_ext
	return nil
}
func (ϟa *GlSamplerParameterIivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_29_ext := ExtensionId_GL_EXT_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_29_ext
	return nil
}
func (ϟa *GlSamplerParameterIuivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_30_ext := ExtensionId_GL_EXT_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_30_ext
	return nil
}
func (ϟa *GlTexBufferEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_31_ext := ExtensionId_GL_EXT_texture_buffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_31_ext
	return nil
}
func (ϟa *GlTexBufferRangeEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_32_ext := ExtensionId_GL_EXT_texture_buffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_32_ext
	return nil
}
func (ϟa *GlTexParameterIivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_33_ext := ExtensionId_GL_EXT_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_33_ext
	return nil
}
func (ϟa *GlTexParameterIuivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_34_ext := ExtensionId_GL_EXT_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_34_ext
	return nil
}
func (ϟa *GlTexStorage3DMultisampleOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_35_ext := ExtensionId_GL_OES_texture_storage_multisample_2d_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_35_ext
	return nil
}
func (ϟa *GlBeginQuery) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_36_major := uint32(3) // u32
	minRequiredVersion_36_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ANY_SAMPLES_PASSED, GLenum_GL_ANY_SAMPLES_PASSED_CONSERVATIVE, GLenum_GL_TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_36_major, minRequiredVersion_36_minor
	return nil
}
func (ϟa *GlDeleteQueries) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_38_major := uint32(3)                               // u32
	minRequiredVersion_38_minor := uint32(0)                               // u32
	q := ϟa.Queries.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // QueryIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                           // Contextʳ
	GetContext_39_result := context                                        // Contextʳ
	ctx := GetContext_39_result                                            // Contextʳ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.Queries, q.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_38_major, minRequiredVersion_38_minor, q, context, GetContext_39_result, ctx
	return nil
}
func (ϟa *GlEndQuery) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_40_major := uint32(3) // u32
	minRequiredVersion_40_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ANY_SAMPLES_PASSED, GLenum_GL_ANY_SAMPLES_PASSED_CONSERVATIVE, GLenum_GL_TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_40_major, minRequiredVersion_40_minor
	return nil
}
func (ϟa *GlGenQueries) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_42_major := uint32(3)                               // u32
	minRequiredVersion_42_minor := uint32(0)                               // u32
	q := ϟa.Queries.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // QueryIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                           // Contextʳ
	GetContext_43_result := context                                        // Contextʳ
	ctx := GetContext_43_result                                            // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := QueryId(ϟa.Queries.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil)) // QueryId
		ctx.Instances.Queries[id] = func() *Query {
			s := &Query{}
			s.Init()
			return s
		}()
		q.Index(uint64(i), ϟs).Write(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = minRequiredVersion_42_major, minRequiredVersion_42_minor, q, context, GetContext_43_result, ctx
	return nil
}
func (ϟa *GlGetQueryObjectuiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_44_major := uint32(3) // u32
	minRequiredVersion_44_minor := uint32(0) // u32
	switch ϟa.Parameter {
	case GLenum_GL_QUERY_RESULT, GLenum_GL_QUERY_RESULT_AVAILABLE:
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_44_major, minRequiredVersion_44_minor
	return nil
}
func (ϟa *GlGetQueryiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_46_major := uint32(3) // u32
	minRequiredVersion_46_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ANY_SAMPLES_PASSED, GLenum_GL_ANY_SAMPLES_PASSED_CONSERVATIVE, GLenum_GL_TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Parameter {
	case GLenum_GL_CURRENT_QUERY:
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_46_major, minRequiredVersion_46_minor
	return nil
}
func (ϟa *GlIsQuery) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_49_major := uint32(3)     // u32
	minRequiredVersion_49_minor := uint32(0)     // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_50_result := context              // Contextʳ
	ctx := GetContext_50_result                  // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.Queries.Contains(ϟa.Query) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = minRequiredVersion_49_major, minRequiredVersion_49_minor, context, GetContext_50_result, ctx
	return nil
}
func (ϟa *GlBindBuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_51_major := uint32(2) // u32
	minRequiredVersion_51_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER:
	case GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
		minRequiredVersion_52_major := uint32(3) // u32
		minRequiredVersion_52_minor := uint32(0) // u32
		_, _ = minRequiredVersion_52_major, minRequiredVersion_52_minor
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_53_major := uint32(3) // u32
		minRequiredVersion_53_minor := uint32(1) // u32
		_, _ = minRequiredVersion_53_major, minRequiredVersion_53_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_55_result := context              // Contextʳ
	ctx := GetContext_55_result                  // Contextʳ
	if !(ctx.Instances.Buffers.Contains(ϟa.Buffer)) {
		ctx.Instances.Buffers[ϟa.Buffer] = func() *Buffer {
			s := &Buffer{}
			s.Init()
			return s
		}()
	}
	ctx.BoundBuffers[ϟa.Target] = ϟa.Buffer
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_51_major, minRequiredVersion_51_minor, context, GetContext_55_result, ctx
	return nil
}
func (ϟa *GlBindBufferBase) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_56_major := uint32(3) // u32
	minRequiredVersion_56_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_57_major := uint32(3) // u32
		minRequiredVersion_57_minor := uint32(1) // u32
		_, _ = minRequiredVersion_57_major, minRequiredVersion_57_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_56_major, minRequiredVersion_56_minor
	return nil
}
func (ϟa *GlBindBufferRange) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_59_major := uint32(3) // u32
	minRequiredVersion_59_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_60_major := uint32(3) // u32
		minRequiredVersion_60_minor := uint32(1) // u32
		_, _ = minRequiredVersion_60_major, minRequiredVersion_60_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_59_major, minRequiredVersion_59_minor
	return nil
}
func (ϟa *GlBufferData) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_62_major := uint32(2) // u32
	minRequiredVersion_62_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER:
	case GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
		minRequiredVersion_63_major := uint32(3) // u32
		minRequiredVersion_63_minor := uint32(0) // u32
		_, _ = minRequiredVersion_63_major, minRequiredVersion_63_minor
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_64_major := uint32(3) // u32
		minRequiredVersion_64_minor := uint32(1) // u32
		_, _ = minRequiredVersion_64_major, minRequiredVersion_64_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Usage {
	case GLenum_GL_DYNAMIC_DRAW, GLenum_GL_STATIC_DRAW, GLenum_GL_STREAM_DRAW:
	case GLenum_GL_DYNAMIC_COPY, GLenum_GL_DYNAMIC_READ, GLenum_GL_STATIC_COPY, GLenum_GL_STATIC_READ, GLenum_GL_STREAM_COPY, GLenum_GL_STREAM_READ:
		minRequiredVersion_66_major := uint32(3) // u32
		minRequiredVersion_66_minor := uint32(0) // u32
		_, _ = minRequiredVersion_66_major, minRequiredVersion_66_minor
	default:
		v := ϟa.Usage
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_68_result := context              // Contextʳ
	ctx := GetContext_68_result                  // Contextʳ
	id := ctx.BoundBuffers.Get(ϟa.Target)        // BufferId
	b := ctx.Instances.Buffers.Get(id)           // Bufferʳ
	b.Data = func() (result U8ˢ) {
		switch (ϟa.Data) != (BufferDataPointer(Voidᶜᵖ{})) {
		case true:
			return U8ᵖ(ϟa.Data).Slice(uint64(GLsizeiptr(int32(0))), uint64(ϟa.Size), ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
		case false:
			return MakeU8ˢ(uint64(ϟa.Size), ϟs)
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (ϟa.Data) != (BufferDataPointer(Voidᶜᵖ{})), ϟa))
			return result
		}
	}()
	b.Size = ϟa.Size
	b.Usage = ϟa.Usage
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = minRequiredVersion_62_major, minRequiredVersion_62_minor, context, GetContext_68_result, ctx, id, b
	return nil
}
func (ϟa *GlBufferSubData) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_69_major := uint32(2) // u32
	minRequiredVersion_69_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER:
	case GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
		minRequiredVersion_70_major := uint32(3) // u32
		minRequiredVersion_70_minor := uint32(0) // u32
		_, _ = minRequiredVersion_70_major, minRequiredVersion_70_minor
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_71_major := uint32(3) // u32
		minRequiredVersion_71_minor := uint32(1) // u32
		_, _ = minRequiredVersion_71_major, minRequiredVersion_71_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.Data.Slice(uint64(GLsizeiptr(int32(0))), uint64(ϟa.Size), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_69_major, minRequiredVersion_69_minor
	return nil
}
func (ϟa *GlCopyBufferSubData) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_73_major := uint32(3) // u32
	minRequiredVersion_73_minor := uint32(0) // u32
	switch ϟa.ReadTarget {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	default:
		v := ϟa.ReadTarget
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.WriteTarget {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	default:
		v := ϟa.WriteTarget
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_73_major, minRequiredVersion_73_minor
	return nil
}
func (ϟa *GlDeleteBuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_76_major := uint32(2)                               // u32
	minRequiredVersion_76_minor := uint32(0)                               // u32
	b := ϟa.Buffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // BufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                           // Contextʳ
	GetContext_77_result := context                                        // Contextʳ
	ctx := GetContext_77_result                                            // Contextʳ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.Buffers, b.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_76_major, minRequiredVersion_76_minor, b, context, GetContext_77_result, ctx
	return nil
}
func (ϟa *GlGenBuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_78_major := uint32(2)                               // u32
	minRequiredVersion_78_minor := uint32(0)                               // u32
	b := ϟa.Buffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // BufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                           // Contextʳ
	GetContext_79_result := context                                        // Contextʳ
	ctx := GetContext_79_result                                            // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := BufferId(ϟa.Buffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil)) // BufferId
		ctx.Instances.Buffers[id] = func() *Buffer {
			s := &Buffer{}
			s.Init()
			return s
		}()
		b.Index(uint64(i), ϟs).Write(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = minRequiredVersion_78_major, minRequiredVersion_78_minor, b, context, GetContext_79_result, ctx
	return nil
}
func (ϟa *GlGetBufferParameteri64v) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_80_major := uint32(3) // u32
	minRequiredVersion_80_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Pname {
	case GLenum_GL_BUFFER_ACCESS_FLAGS, GLenum_GL_BUFFER_MAPPED, GLenum_GL_BUFFER_MAP_LENGTH, GLenum_GL_BUFFER_MAP_OFFSET, GLenum_GL_BUFFER_SIZE, GLenum_GL_BUFFER_USAGE:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_80_major, minRequiredVersion_80_minor
	return nil
}
func (ϟa *GlGetBufferParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_83_major := uint32(2) // u32
	minRequiredVersion_83_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER:
	case GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
		minRequiredVersion_84_major := uint32(3) // u32
		minRequiredVersion_84_minor := uint32(0) // u32
		_, _ = minRequiredVersion_84_major, minRequiredVersion_84_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Parameter {
	case GLenum_GL_BUFFER_SIZE, GLenum_GL_BUFFER_USAGE:
	case GLenum_GL_BUFFER_ACCESS_FLAGS, GLenum_GL_BUFFER_MAPPED, GLenum_GL_BUFFER_MAP_LENGTH, GLenum_GL_BUFFER_MAP_OFFSET:
		minRequiredVersion_86_major := uint32(3) // u32
		minRequiredVersion_86_minor := uint32(0) // u32
		_, _ = minRequiredVersion_86_major, minRequiredVersion_86_minor
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_88_result := context              // Contextʳ
	ctx := GetContext_88_result                  // Contextʳ
	id := ctx.BoundBuffers.Get(ϟa.Target)        // BufferId
	b := ctx.Instances.Buffers.Get(id)           // Bufferʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(func() (result GLint) {
		switch ϟa.Parameter {
		case GLenum_GL_BUFFER_SIZE:
			return GLint(b.Size)
		case GLenum_GL_BUFFER_USAGE:
			return GLint(b.Usage)
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.Parameter, ϟa))
			return result
		}
	}(), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _, _ = minRequiredVersion_83_major, minRequiredVersion_83_minor, context, GetContext_88_result, ctx, id, b
	return nil
}
func (ϟa *GlGetBufferPointerv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_89_major := uint32(3) // u32
	minRequiredVersion_89_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Pname {
	case GLenum_GL_BUFFER_MAP_POINTER:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_89_major, minRequiredVersion_89_minor
	return nil
}
func (ϟa *GlIsBuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_92_major := uint32(2)     // u32
	minRequiredVersion_92_minor := uint32(0)     // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_93_result := context              // Contextʳ
	ctx := GetContext_93_result                  // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.Buffers.Contains(ϟa.Buffer) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = minRequiredVersion_92_major, minRequiredVersion_92_minor, context, GetContext_93_result, ctx
	return nil
}
func (ϟa *GlMapBufferRange) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_94_major := uint32(3) // u32
	minRequiredVersion_94_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_95_major := uint32(3) // u32
		minRequiredVersion_95_minor := uint32(1) // u32
		_, _ = minRequiredVersion_95_major, minRequiredVersion_95_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	supportsBits_97_seenBits := ϟa.Access                                                                                                                                                                                                                                      // GLbitfield
	supportsBits_97_validBits := (GLbitfield_GL_MAP_FLUSH_EXPLICIT_BIT) | ((GLbitfield_GL_MAP_INVALIDATE_BUFFER_BIT) | ((GLbitfield_GL_MAP_INVALIDATE_RANGE_BIT) | ((GLbitfield_GL_MAP_READ_BIT) | ((GLbitfield_GL_MAP_UNSYNCHRONIZED_BIT) | (GLbitfield_GL_MAP_WRITE_BIT))))) // GLbitfield
	if (GLbitfield_GL_MAP_FLUSH_EXPLICIT_BIT)&(ϟa.Access) != 0 {
	}
	if (GLbitfield_GL_MAP_INVALIDATE_BUFFER_BIT)&(ϟa.Access) != 0 {
	}
	if (GLbitfield_GL_MAP_INVALIDATE_RANGE_BIT)&(ϟa.Access) != 0 {
	}
	if (GLbitfield_GL_MAP_READ_BIT)&(ϟa.Access) != 0 {
	}
	if (GLbitfield_GL_MAP_UNSYNCHRONIZED_BIT)&(ϟa.Access) != 0 {
	}
	if (GLbitfield_GL_MAP_WRITE_BIT)&(ϟa.Access) != 0 {
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                    // Contextʳ
	GetContext_98_result := context                                 // Contextʳ
	ctx := GetContext_98_result                                     // Contextʳ
	b := ctx.Instances.Buffers.Get(ctx.BoundBuffers.Get(ϟa.Target)) // Bufferʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ptr := U8ᵖ(ϟa.Result) // U8ᵖ
	b.MappingAccess = ϟa.Access
	b.MappingData = ptr.Slice(uint64(GLsizeiptr(int32(0))), uint64(ϟa.Length), ϟs)
	if (GLbitfield_GL_MAP_READ_BIT)&(ϟa.Access) != 0 {
		src := b.Data.Slice(uint64(ϟa.Offset), uint64((ϟa.Offset)+(GLintptr(ϟa.Length))), ϟs) // U8ˢ
		dst := b.MappingData                                                                  // U8ˢ
		dst.OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = src, dst
	}
	ϟa.Result = Voidᵖ(ptr)
	_, _, _, _, _, _, _, _, _ = minRequiredVersion_94_major, minRequiredVersion_94_minor, supportsBits_97_seenBits, supportsBits_97_validBits, context, GetContext_98_result, ctx, b, ptr
	return nil
}
func (ϟa *GlUnmapBuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_99_major := uint32(3) // u32
	minRequiredVersion_99_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_100_major := uint32(3) // u32
		minRequiredVersion_100_minor := uint32(1) // u32
		_, _ = minRequiredVersion_100_major, minRequiredVersion_100_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                    // Contextʳ
	GetContext_102_result := context                                // Contextʳ
	ctx := GetContext_102_result                                    // Contextʳ
	b := ctx.Instances.Buffers.Get(ctx.BoundBuffers.Get(ϟa.Target)) // Bufferʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	b.Data.Slice(uint64(b.MappingOffset), uint64((b.MappingOffset)+(int32(b.MappingData.Count))), ϟs).Copy(b.MappingData, ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Result = ϟa.Result
	_, _, _, _, _, _ = minRequiredVersion_99_major, minRequiredVersion_99_minor, context, GetContext_102_result, ctx, b
	return nil
}
func (ϟa *GlDrawArrays) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_103_major := uint32(2) // u32
	minRequiredVersion_103_minor := uint32(0) // u32
	switch ϟa.DrawMode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	default:
		v := ϟa.DrawMode
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                    // Contextʳ
	GetContext_105_result := context                                                // Contextʳ
	ctx := GetContext_105_result                                                    // Contextʳ
	last_index := (uint32(ϟa.FirstIndex)) + ((uint32(ϟa.IndexCount)) - (uint32(1))) // u32
	ReadVertexArrays_106_ctx := ctx                                                 // Contextʳ
	ReadVertexArrays_106_first_index := uint32(ϟa.FirstIndex)                       // u32
	ReadVertexArrays_106_last_index := last_index                                   // u32
	for i := int32(int32(0)); i < int32(len(ReadVertexArrays_106_ctx.VertexAttributeArrays)); i++ {
		arr := ReadVertexArrays_106_ctx.VertexAttributeArrays.Get(AttributeLocation(i)) // VertexAttributeArrayʳ
		if (arr.Enabled) && ((arr.Buffer) == (BufferId(uint32(0)))) {
			vertexAttribTypeSize_107_t := arr.Type // GLenum
			vertexAttribTypeSize_107_result := func() (result uint32) {
				switch vertexAttribTypeSize_107_t {
				case GLenum_GL_BYTE:
					return uint32(1)
				case GLenum_GL_UNSIGNED_BYTE:
					return uint32(1)
				case GLenum_GL_SHORT:
					return uint32(2)
				case GLenum_GL_UNSIGNED_SHORT:
					return uint32(2)
				case GLenum_GL_FIXED:
					return uint32(4)
				case GLenum_GL_FLOAT:
					return uint32(4)
				case GLenum_GL_HALF_FLOAT_ARB:
					return uint32(2)
				case GLenum_GL_HALF_FLOAT_OES:
					return uint32(2)
				default:
					// TODO: better unmatched handling
					panic(fmt.Errorf("Unmatched switch(%v) in atom %T", vertexAttribTypeSize_107_t, ϟa))
					return result
				}
			}() // u32
			elsize := (vertexAttribTypeSize_107_result) * (arr.Size) // u32
			elstride := func() (result uint32) {
				switch (arr.Stride) == (GLsizei(int32(0))) {
				case true:
					return elsize
				case false:
					return uint32(arr.Stride)
				default:
					// TODO: better unmatched handling
					panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (arr.Stride) == (GLsizei(int32(0))), ϟa))
					return result
				}
			}() // u32
			for v := uint32(ReadVertexArrays_106_first_index); v < (ReadVertexArrays_106_last_index)+(uint32(1)); v++ {
				offset := (elstride) * (v) // u32
				arr.Pointer.Slice(uint64(offset), uint64((offset)+(elsize)), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
				_ = offset
			}
			_, _, _, _ = vertexAttribTypeSize_107_t, vertexAttribTypeSize_107_result, elsize, elstride
		}
		_ = arr
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _, _ = minRequiredVersion_103_major, minRequiredVersion_103_minor, context, GetContext_105_result, ctx, last_index, ReadVertexArrays_106_ctx, ReadVertexArrays_106_first_index, ReadVertexArrays_106_last_index
	return nil
}
func (ϟa *GlDrawArraysIndirect) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_108_major := uint32(3) // u32
	minRequiredVersion_108_minor := uint32(1) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	default:
		v := ϟa.Mode
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_108_major, minRequiredVersion_108_minor
	return nil
}
func (ϟa *GlDrawArraysInstanced) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_110_major := uint32(3) // u32
	minRequiredVersion_110_minor := uint32(0) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	default:
		v := ϟa.Mode
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_110_major, minRequiredVersion_110_minor
	return nil
}
func (ϟa *GlDrawBuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_112_major := uint32(3) // u32
	minRequiredVersion_112_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_112_major, minRequiredVersion_112_minor
	return nil
}
func (ϟa *GlDrawElements) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_113_major := uint32(2) // u32
	minRequiredVersion_113_minor := uint32(0) // u32
	switch ϟa.DrawMode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	default:
		v := ϟa.DrawMode
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.IndicesType {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_SHORT:
	case GLenum_GL_UNSIGNED_INT:
		minRequiredVersion_115_major := uint32(3) // u32
		minRequiredVersion_115_minor := uint32(0) // u32
		_, _ = minRequiredVersion_115_major, minRequiredVersion_115_minor
	default:
		v := ϟa.IndicesType
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)               // Contextʳ
	GetContext_117_result := context                           // Contextʳ
	ctx := GetContext_117_result                               // Contextʳ
	count := uint32(ϟa.ElementCount)                           // u32
	id := ctx.BoundBuffers.Get(GLenum_GL_ELEMENT_ARRAY_BUFFER) // BufferId
	if (id) != (BufferId(uint32(0))) {
		index_data := ctx.Instances.Buffers.Get(id).Data                                                           // U8ˢ
		offset := uint32(uint64(ϟa.Indices.Address))                                                               // u32
		first := externs{ϟa, ϟs, ϟd, ϟl, ϟb}.minIndex(U8ᵖ(index_data.Index(0, ϟs)), ϟa.IndicesType, offset, count) // u32
		last := externs{ϟa, ϟs, ϟd, ϟl, ϟb}.maxIndex(U8ᵖ(index_data.Index(0, ϟs)), ϟa.IndicesType, offset, count)  // u32
		ReadVertexArrays_118_ctx := ctx                                                                            // Contextʳ
		ReadVertexArrays_118_first_index := first                                                                  // u32
		ReadVertexArrays_118_last_index := last                                                                    // u32
		for i := int32(int32(0)); i < int32(len(ReadVertexArrays_118_ctx.VertexAttributeArrays)); i++ {
			arr := ReadVertexArrays_118_ctx.VertexAttributeArrays.Get(AttributeLocation(i)) // VertexAttributeArrayʳ
			if (arr.Enabled) && ((arr.Buffer) == (BufferId(uint32(0)))) {
				vertexAttribTypeSize_119_t := arr.Type // GLenum
				vertexAttribTypeSize_119_result := func() (result uint32) {
					switch vertexAttribTypeSize_119_t {
					case GLenum_GL_BYTE:
						return uint32(1)
					case GLenum_GL_UNSIGNED_BYTE:
						return uint32(1)
					case GLenum_GL_SHORT:
						return uint32(2)
					case GLenum_GL_UNSIGNED_SHORT:
						return uint32(2)
					case GLenum_GL_FIXED:
						return uint32(4)
					case GLenum_GL_FLOAT:
						return uint32(4)
					case GLenum_GL_HALF_FLOAT_ARB:
						return uint32(2)
					case GLenum_GL_HALF_FLOAT_OES:
						return uint32(2)
					default:
						// TODO: better unmatched handling
						panic(fmt.Errorf("Unmatched switch(%v) in atom %T", vertexAttribTypeSize_119_t, ϟa))
						return result
					}
				}() // u32
				elsize := (vertexAttribTypeSize_119_result) * (arr.Size) // u32
				elstride := func() (result uint32) {
					switch (arr.Stride) == (GLsizei(int32(0))) {
					case true:
						return elsize
					case false:
						return uint32(arr.Stride)
					default:
						// TODO: better unmatched handling
						panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (arr.Stride) == (GLsizei(int32(0))), ϟa))
						return result
					}
				}() // u32
				for v := uint32(ReadVertexArrays_118_first_index); v < (ReadVertexArrays_118_last_index)+(uint32(1)); v++ {
					offset := (elstride) * (v) // u32
					arr.Pointer.Slice(uint64(offset), uint64((offset)+(elsize)), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
					_ = offset
				}
				_, _, _, _ = vertexAttribTypeSize_119_t, vertexAttribTypeSize_119_result, elsize, elstride
			}
			_ = arr
		}
		_, _, _, _, _, _, _ = index_data, offset, first, last, ReadVertexArrays_118_ctx, ReadVertexArrays_118_first_index, ReadVertexArrays_118_last_index
	} else {
		index_data := U8ᵖ(ϟa.Indices)                                                               // U8ᵖ
		first := externs{ϟa, ϟs, ϟd, ϟl, ϟb}.minIndex(index_data, ϟa.IndicesType, uint32(0), count) // u32
		last := externs{ϟa, ϟs, ϟd, ϟl, ϟb}.maxIndex(index_data, ϟa.IndicesType, uint32(0), count)  // u32
		ReadVertexArrays_120_ctx := ctx                                                             // Contextʳ
		ReadVertexArrays_120_first_index := first                                                   // u32
		ReadVertexArrays_120_last_index := last                                                     // u32
		for i := int32(int32(0)); i < int32(len(ReadVertexArrays_120_ctx.VertexAttributeArrays)); i++ {
			arr := ReadVertexArrays_120_ctx.VertexAttributeArrays.Get(AttributeLocation(i)) // VertexAttributeArrayʳ
			if (arr.Enabled) && ((arr.Buffer) == (BufferId(uint32(0)))) {
				vertexAttribTypeSize_121_t := arr.Type // GLenum
				vertexAttribTypeSize_121_result := func() (result uint32) {
					switch vertexAttribTypeSize_121_t {
					case GLenum_GL_BYTE:
						return uint32(1)
					case GLenum_GL_UNSIGNED_BYTE:
						return uint32(1)
					case GLenum_GL_SHORT:
						return uint32(2)
					case GLenum_GL_UNSIGNED_SHORT:
						return uint32(2)
					case GLenum_GL_FIXED:
						return uint32(4)
					case GLenum_GL_FLOAT:
						return uint32(4)
					case GLenum_GL_HALF_FLOAT_ARB:
						return uint32(2)
					case GLenum_GL_HALF_FLOAT_OES:
						return uint32(2)
					default:
						// TODO: better unmatched handling
						panic(fmt.Errorf("Unmatched switch(%v) in atom %T", vertexAttribTypeSize_121_t, ϟa))
						return result
					}
				}() // u32
				elsize := (vertexAttribTypeSize_121_result) * (arr.Size) // u32
				elstride := func() (result uint32) {
					switch (arr.Stride) == (GLsizei(int32(0))) {
					case true:
						return elsize
					case false:
						return uint32(arr.Stride)
					default:
						// TODO: better unmatched handling
						panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (arr.Stride) == (GLsizei(int32(0))), ϟa))
						return result
					}
				}() // u32
				for v := uint32(ReadVertexArrays_120_first_index); v < (ReadVertexArrays_120_last_index)+(uint32(1)); v++ {
					offset := (elstride) * (v) // u32
					arr.Pointer.Slice(uint64(offset), uint64((offset)+(elsize)), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
					_ = offset
				}
				_, _, _, _ = vertexAttribTypeSize_121_t, vertexAttribTypeSize_121_result, elsize, elstride
			}
			_ = arr
		}
		IndexSize_122_indices_type := ϟa.IndicesType // GLenum
		IndexSize_122_result := func() (result uint32) {
			switch IndexSize_122_indices_type {
			case GLenum_GL_UNSIGNED_BYTE:
				return uint32(1)
			case GLenum_GL_UNSIGNED_SHORT:
				return uint32(2)
			case GLenum_GL_UNSIGNED_INT:
				return uint32(4)
			default:
				// TODO: better unmatched handling
				panic(fmt.Errorf("Unmatched switch(%v) in atom %T", IndexSize_122_indices_type, ϟa))
				return result
			}
		}() // u32
		index_data.Slice(uint64(uint32(0)), uint64((uint32(ϟa.ElementCount))*(IndexSize_122_result)), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _, _, _, _, _, _, _ = index_data, first, last, ReadVertexArrays_120_ctx, ReadVertexArrays_120_first_index, ReadVertexArrays_120_last_index, IndexSize_122_indices_type, IndexSize_122_result
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = minRequiredVersion_113_major, minRequiredVersion_113_minor, context, GetContext_117_result, ctx, count, id
	return nil
}
func (ϟa *GlDrawElementsIndirect) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_123_major := uint32(3) // u32
	minRequiredVersion_123_minor := uint32(1) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	default:
		v := ϟa.Mode
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Type {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_SHORT:
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_123_major, minRequiredVersion_123_minor
	return nil
}
func (ϟa *GlDrawElementsInstanced) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_126_major := uint32(3) // u32
	minRequiredVersion_126_minor := uint32(0) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	default:
		v := ϟa.Mode
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Type {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_SHORT:
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_126_major, minRequiredVersion_126_minor
	return nil
}
func (ϟa *GlDrawRangeElements) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_129_major := uint32(3) // u32
	minRequiredVersion_129_minor := uint32(0) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	default:
		v := ϟa.Mode
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Type {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_SHORT:
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_129_major, minRequiredVersion_129_minor
	return nil
}
func (ϟa *GlActiveShaderProgramEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_132_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_132_ext
	return nil
}
func (ϟa *GlAlphaFuncQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_133_ext := ExtensionId_GL_QCOM_alpha_test // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_133_ext
	return nil
}
func (ϟa *GlBeginConditionalRenderNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_134_ext := ExtensionId_GL_NV_conditional_render // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_134_ext
	return nil
}
func (ϟa *GlBeginPerfMonitorAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_135_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_135_ext
	return nil
}
func (ϟa *GlBeginPerfQueryINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_136_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_136_ext
	return nil
}
func (ϟa *GlBeginQueryEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_137_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_138_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = requiresExtension_137_ext, requiresExtension_138_ext
	return nil
}
func (ϟa *GlBindProgramPipelineEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_139_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_139_ext
	return nil
}
func (ϟa *GlBindVertexArrayOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_140_ext := ExtensionId_GL_OES_vertex_array_object // ExtensionId
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                        // Contextʳ
	GetContext_141_result := context                                    // Contextʳ
	ctx := GetContext_141_result                                        // Contextʳ
	if !(ctx.Instances.VertexArrays.Contains(ϟa.Array)) {
		ctx.Instances.VertexArrays[ϟa.Array] = func() *VertexArray {
			s := &VertexArray{}
			s.Init()
			return s
		}()
	}
	ctx.BoundVertexArray = ϟa.Array
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = requiresExtension_140_ext, context, GetContext_141_result, ctx
	return nil
}
func (ϟa *GlBlendBarrierNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_142_ext := ExtensionId_GL_NV_blend_equation_advanced // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_142_ext
	return nil
}
func (ϟa *GlBlendEquationSeparateiOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_143_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_143_ext
	return nil
}
func (ϟa *GlBlendEquationiOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_144_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_144_ext
	return nil
}
func (ϟa *GlBlendFuncSeparateiOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_145_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_145_ext
	return nil
}
func (ϟa *GlBlendFunciOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_146_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_146_ext
	return nil
}
func (ϟa *GlBlendParameteriNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_147_ext := ExtensionId_GL_NV_blend_equation_advanced // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_147_ext
	return nil
}
func (ϟa *GlBlitFramebufferANGLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_148_ext := ExtensionId_GL_ANGLE_framebuffer_blit // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_148_ext
	return nil
}
func (ϟa *GlBlitFramebufferNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_149_ext := ExtensionId_GL_NV_framebuffer_blit // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_149_ext
	return nil
}
func (ϟa *GlBufferStorageEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_150_ext := ExtensionId_GL_EXT_buffer_storage // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_150_ext
	return nil
}
func (ϟa *GlClientWaitSyncAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_151_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_151_ext
	return nil
}
func (ϟa *GlColorMaskiOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_152_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_152_ext
	return nil
}
func (ϟa *GlCompressedTexImage3DOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_153_ext := ExtensionId_GL_OES_texture_3D // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_153_ext
	return nil
}
func (ϟa *GlCompressedTexSubImage3DOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_154_ext := ExtensionId_GL_OES_texture_3D // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_154_ext
	return nil
}
func (ϟa *GlCopyBufferSubDataNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_155_ext := ExtensionId_GL_NV_copy_buffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_155_ext
	return nil
}
func (ϟa *GlCopyImageSubDataOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_156_ext := ExtensionId_GL_OES_copy_image // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_156_ext
	return nil
}
func (ϟa *GlCopyPathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_157_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_157_ext
	return nil
}
func (ϟa *GlCopyTexSubImage3DOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_158_ext := ExtensionId_GL_OES_texture_3D // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_158_ext
	return nil
}
func (ϟa *GlCopyTextureLevelsAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_159_ext := ExtensionId_GL_APPLE_copy_texture_levels // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_159_ext
	return nil
}
func (ϟa *GlCoverFillPathInstancedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_160_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_160_ext
	return nil
}
func (ϟa *GlCoverFillPathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_161_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_161_ext
	return nil
}
func (ϟa *GlCoverStrokePathInstancedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_162_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_162_ext
	return nil
}
func (ϟa *GlCoverStrokePathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_163_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_163_ext
	return nil
}
func (ϟa *GlCoverageMaskNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_164_ext := ExtensionId_GL_NV_coverage_sample // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_164_ext
	return nil
}
func (ϟa *GlCoverageOperationNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_165_ext := ExtensionId_GL_NV_coverage_sample // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_165_ext
	return nil
}
func (ϟa *GlCreatePerfQueryINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_166_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_166_ext
	return nil
}
func (ϟa *GlCreateShaderProgramvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_167_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_167_ext
	return nil
}
func (ϟa *GlDeleteFencesNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_168_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_168_ext
	return nil
}
func (ϟa *GlDeletePathsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_169_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_169_ext
	return nil
}
func (ϟa *GlDeletePerfMonitorsAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_170_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_170_ext
	return nil
}
func (ϟa *GlDeletePerfQueryINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_171_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_171_ext
	return nil
}
func (ϟa *GlDeleteProgramPipelinesEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_172_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_172_ext
	return nil
}
func (ϟa *GlDeleteQueriesEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_173_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_174_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	q := ϟa.Queries.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs)  // QueryIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                            // Contextʳ
	GetContext_175_result := context                                        // Contextʳ
	ctx := GetContext_175_result                                            // Contextʳ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.Queries, q.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = requiresExtension_173_ext, requiresExtension_174_ext, q, context, GetContext_175_result, ctx
	return nil
}
func (ϟa *GlDeleteSyncAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_176_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_176_ext
	return nil
}
func (ϟa *GlDeleteVertexArraysOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_177_ext := ExtensionId_GL_OES_vertex_array_object   // ExtensionId
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                          // Contextʳ
	GetContext_178_result := context                                      // Contextʳ
	ctx := GetContext_178_result                                          // Contextʳ
	a := ϟa.Arrays.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // VertexArrayIdˢ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.VertexArrays, a.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = requiresExtension_177_ext, context, GetContext_178_result, ctx, a
	return nil
}
func (ϟa *GlDepthRangeArrayfvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_179_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_179_ext
	return nil
}
func (ϟa *GlDepthRangeIndexedfNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_180_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_180_ext
	return nil
}
func (ϟa *GlDisableDriverControlQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_181_ext := ExtensionId_GL_QCOM_driver_control // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_181_ext
	return nil
}
func (ϟa *GlDisableiNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_182_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_182_ext
	return nil
}
func (ϟa *GlDisableiOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_183_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_183_ext
	return nil
}
func (ϟa *GlDiscardFramebufferEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_184_ext := ExtensionId_GL_EXT_discard_framebuffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_184_ext
	return nil
}
func (ϟa *GlDrawArraysInstancedANGLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_185_ext := ExtensionId_GL_ANGLE_instanced_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_185_ext
	return nil
}
func (ϟa *GlDrawArraysInstancedBaseInstanceEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_186_ext := ExtensionId_GL_EXT_base_instance // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_186_ext
	return nil
}
func (ϟa *GlDrawArraysInstancedEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_187_ext := ExtensionId_GL_EXT_draw_instanced   // ExtensionId
	requiresExtension_188_ext := ExtensionId_GL_EXT_instanced_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = requiresExtension_187_ext, requiresExtension_188_ext
	return nil
}
func (ϟa *GlDrawArraysInstancedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_189_ext := ExtensionId_GL_NV_draw_instanced // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_189_ext
	return nil
}
func (ϟa *GlDrawBuffersEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_190_ext := ExtensionId_GL_EXT_draw_buffers // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_190_ext
	return nil
}
func (ϟa *GlDrawBuffersIndexedEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_191_ext := ExtensionId_GL_EXT_multiview_draw_buffers // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_191_ext
	return nil
}
func (ϟa *GlDrawBuffersNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_192_ext := ExtensionId_GL_NV_draw_buffers // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_192_ext
	return nil
}
func (ϟa *GlDrawElementsBaseVertexEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_193_ext := ExtensionId_GL_EXT_draw_elements_base_vertex // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_193_ext
	return nil
}
func (ϟa *GlDrawElementsBaseVertexOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_194_ext := ExtensionId_GL_OES_draw_elements_base_vertex // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_194_ext
	return nil
}
func (ϟa *GlDrawElementsInstancedANGLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_195_ext := ExtensionId_GL_ANGLE_instanced_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_195_ext
	return nil
}
func (ϟa *GlDrawElementsInstancedBaseInstanceEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_196_ext := ExtensionId_GL_EXT_base_instance // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_196_ext
	return nil
}
func (ϟa *GlDrawElementsInstancedBaseVertexEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_197_ext := ExtensionId_GL_EXT_draw_elements_base_vertex // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_197_ext
	return nil
}
func (ϟa *GlDrawElementsInstancedBaseVertexOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_198_ext := ExtensionId_GL_OES_draw_elements_base_vertex // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_198_ext
	return nil
}
func (ϟa *GlDrawElementsInstancedBaseVertexBaseInstanceEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_199_ext := ExtensionId_GL_EXT_base_instance // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_199_ext
	return nil
}
func (ϟa *GlDrawElementsInstancedEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_200_ext := ExtensionId_GL_EXT_draw_instanced   // ExtensionId
	requiresExtension_201_ext := ExtensionId_GL_EXT_instanced_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = requiresExtension_200_ext, requiresExtension_201_ext
	return nil
}
func (ϟa *GlDrawElementsInstancedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_202_ext := ExtensionId_GL_NV_draw_instanced // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_202_ext
	return nil
}
func (ϟa *GlDrawRangeElementsBaseVertexEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_203_ext := ExtensionId_GL_EXT_draw_elements_base_vertex // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_203_ext
	return nil
}
func (ϟa *GlDrawRangeElementsBaseVertexOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_204_ext := ExtensionId_GL_OES_draw_elements_base_vertex // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_204_ext
	return nil
}
func (ϟa *GlEGLImageTargetRenderbufferStorageOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_205_ext := ExtensionId_GL_OES_EGL_image // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_205_ext
	return nil
}
func (ϟa *GlEGLImageTargetTexture2DOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_206_ext := ExtensionId_GL_OES_EGL_image // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_206_ext
	return nil
}
func (ϟa *GlEnableDriverControlQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_207_ext := ExtensionId_GL_QCOM_driver_control // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_207_ext
	return nil
}
func (ϟa *GlEnableiNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_208_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_208_ext
	return nil
}
func (ϟa *GlEnableiOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_209_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_209_ext
	return nil
}
func (ϟa *GlEndConditionalRenderNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_210_ext := ExtensionId_GL_NV_conditional_render // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_210_ext
	return nil
}
func (ϟa *GlEndPerfMonitorAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_211_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_211_ext
	return nil
}
func (ϟa *GlEndPerfQueryINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_212_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_212_ext
	return nil
}
func (ϟa *GlEndQueryEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_213_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_214_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = requiresExtension_213_ext, requiresExtension_214_ext
	return nil
}
func (ϟa *GlEndTilingQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_215_ext := ExtensionId_GL_QCOM_tiled_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_215_ext
	return nil
}
func (ϟa *GlExtGetBufferPointervQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_216_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_216_ext
	return nil
}
func (ϟa *GlExtGetBuffersQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_217_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_217_ext
	return nil
}
func (ϟa *GlExtGetFramebuffersQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_218_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_218_ext
	return nil
}
func (ϟa *GlExtGetProgramBinarySourceQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_219_ext := ExtensionId_GL_QCOM_extended_get2 // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_219_ext
	return nil
}
func (ϟa *GlExtGetProgramsQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_220_ext := ExtensionId_GL_QCOM_extended_get2 // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_220_ext
	return nil
}
func (ϟa *GlExtGetRenderbuffersQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_221_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_221_ext
	return nil
}
func (ϟa *GlExtGetShadersQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_222_ext := ExtensionId_GL_QCOM_extended_get2 // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_222_ext
	return nil
}
func (ϟa *GlExtGetTexLevelParameterivQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_223_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_223_ext
	return nil
}
func (ϟa *GlExtGetTexSubImageQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_224_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_224_ext
	return nil
}
func (ϟa *GlExtGetTexturesQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_225_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_225_ext
	return nil
}
func (ϟa *GlExtIsProgramBinaryQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_226_ext := ExtensionId_GL_QCOM_extended_get2 // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_226_ext
	return nil
}
func (ϟa *GlExtTexObjectStateOverrideiQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_227_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_227_ext
	return nil
}
func (ϟa *GlFenceSyncAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_228_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_228_ext
	return nil
}
func (ϟa *GlFinishFenceNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_229_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_229_ext
	return nil
}
func (ϟa *GlFlushMappedBufferRangeEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_230_ext := ExtensionId_GL_EXT_map_buffer_range // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_230_ext
	return nil
}
func (ϟa *GlFramebufferTexture2DMultisampleEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_231_ext := ExtensionId_GL_EXT_multisampled_render_to_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_231_ext
	return nil
}
func (ϟa *GlFramebufferTexture2DMultisampleIMG) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_232_ext := ExtensionId_GL_IMG_multisampled_render_to_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_232_ext
	return nil
}
func (ϟa *GlFramebufferTexture3DOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_233_ext := ExtensionId_GL_OES_texture_3D // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_233_ext
	return nil
}
func (ϟa *GlFramebufferTextureOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_234_ext := ExtensionId_GL_OES_geometry_shader // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_234_ext
	return nil
}
func (ϟa *GlFramebufferTextureMultiviewOVR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_235_ext := ExtensionId_GL_OVR_multiview // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_235_ext
	return nil
}
func (ϟa *GlGenFencesNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_236_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_236_ext
	return nil
}
func (ϟa *GlGenPathsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_237_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_237_ext
	return nil
}
func (ϟa *GlGenPerfMonitorsAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_238_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_238_ext
	return nil
}
func (ϟa *GlGenProgramPipelinesEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_239_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_239_ext
	return nil
}
func (ϟa *GlGenQueriesEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_240_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_241_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	q := ϟa.Queries.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs)  // QueryIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                            // Contextʳ
	GetContext_242_result := context                                        // Contextʳ
	ctx := GetContext_242_result                                            // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := QueryId(ϟa.Queries.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil)) // QueryId
		ctx.Instances.Queries[id] = func() *Query {
			s := &Query{}
			s.Init()
			return s
		}()
		q.Index(uint64(i), ϟs).Write(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = requiresExtension_240_ext, requiresExtension_241_ext, q, context, GetContext_242_result, ctx
	return nil
}
func (ϟa *GlGenVertexArraysOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_243_ext := ExtensionId_GL_OES_vertex_array_object   // ExtensionId
	a := ϟa.Arrays.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // VertexArrayIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                          // Contextʳ
	GetContext_244_result := context                                      // Contextʳ
	ctx := GetContext_244_result                                          // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := VertexArrayId(ϟa.Arrays.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil)) // VertexArrayId
		ctx.Instances.VertexArrays[id] = func() *VertexArray {
			s := &VertexArray{}
			s.Init()
			return s
		}()
		a.Index(uint64(i), ϟs).Write(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _ = requiresExtension_243_ext, a, context, GetContext_244_result, ctx
	return nil
}
func (ϟa *GlGetBufferPointervOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_245_ext := ExtensionId_GL_OES_mapbuffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_245_ext
	return nil
}
func (ϟa *GlGetDriverControlStringQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_246_ext := ExtensionId_GL_QCOM_driver_control // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_246_ext
	return nil
}
func (ϟa *GlGetDriverControlsQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_247_ext := ExtensionId_GL_QCOM_driver_control // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_247_ext
	return nil
}
func (ϟa *GlGetFenceivNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_248_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_248_ext
	return nil
}
func (ϟa *GlGetFirstPerfQueryIdINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_249_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_249_ext
	return nil
}
func (ϟa *GlGetFloati_vNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_250_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_250_ext
	return nil
}
func (ϟa *GlGetGraphicsResetStatusEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_251_ext := ExtensionId_GL_EXT_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_251_ext
	return nil
}
func (ϟa *GlGetGraphicsResetStatusKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_252_ext := ExtensionId_GL_KHR_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_252_ext
	return nil
}
func (ϟa *GlGetImageHandleNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_253_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_253_ext
	return nil
}
func (ϟa *GlGetInteger64vAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_254_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_254_ext
	return nil
}
func (ϟa *GlGetIntegeri_vEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_255_ext := ExtensionId_GL_EXT_multiview_draw_buffers // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_255_ext
	return nil
}
func (ϟa *GlGetInternalformatSampleivNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_256_ext := ExtensionId_GL_NV_internalformat_sample_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_256_ext
	return nil
}
func (ϟa *GlGetNextPerfQueryIdINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_257_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_257_ext
	return nil
}
func (ϟa *GlGetObjectLabelEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_258_ext := ExtensionId_GL_EXT_debug_label // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_258_ext
	return nil
}
func (ϟa *GlGetPathCommandsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_259_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_259_ext
	return nil
}
func (ϟa *GlGetPathCoordsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_260_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_260_ext
	return nil
}
func (ϟa *GlGetPathDashArrayNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_261_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_261_ext
	return nil
}
func (ϟa *GlGetPathLengthNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_262_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_262_ext
	return nil
}
func (ϟa *GlGetPathMetricRangeNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_263_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_263_ext
	return nil
}
func (ϟa *GlGetPathMetricsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_264_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_264_ext
	return nil
}
func (ϟa *GlGetPathParameterfvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_265_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_265_ext
	return nil
}
func (ϟa *GlGetPathParameterivNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_266_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_266_ext
	return nil
}
func (ϟa *GlGetPathSpacingNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_267_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_267_ext
	return nil
}
func (ϟa *GlGetPerfCounterInfoINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_268_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_268_ext
	return nil
}
func (ϟa *GlGetPerfMonitorCounterDataAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_269_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_269_ext
	return nil
}
func (ϟa *GlGetPerfMonitorCounterInfoAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_270_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_270_ext
	return nil
}
func (ϟa *GlGetPerfMonitorCounterStringAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_271_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_271_ext
	return nil
}
func (ϟa *GlGetPerfMonitorCountersAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_272_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_272_ext
	return nil
}
func (ϟa *GlGetPerfMonitorGroupStringAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_273_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_273_ext
	return nil
}
func (ϟa *GlGetPerfMonitorGroupsAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_274_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_274_ext
	return nil
}
func (ϟa *GlGetPerfQueryDataINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_275_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_275_ext
	return nil
}
func (ϟa *GlGetPerfQueryIdByNameINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_276_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_276_ext
	return nil
}
func (ϟa *GlGetPerfQueryInfoINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_277_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_277_ext
	return nil
}
func (ϟa *GlGetProgramBinaryOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_278_ext := ExtensionId_GL_OES_get_program_binary // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	l := GLsizei(ϟa.BytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil)) // GLsizei
	ϟa.BytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(l, ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.BinaryFormat.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.BinaryFormat.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Binary.Slice(uint64(GLsizei(int32(0))), uint64(l), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = requiresExtension_278_ext, l
	return nil
}
func (ϟa *GlGetProgramPipelineInfoLogEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_279_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_279_ext
	return nil
}
func (ϟa *GlGetProgramPipelineivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_280_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_280_ext
	return nil
}
func (ϟa *GlGetProgramResourcefvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_281_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_281_ext
	return nil
}
func (ϟa *GlGetQueryObjecti64vEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_282_ext := ExtensionId_GL_EXT_disjoint_timer_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_ = requiresExtension_282_ext
	return nil
}
func (ϟa *GlGetQueryObjectivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_283_ext := ExtensionId_GL_EXT_disjoint_timer_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_ = requiresExtension_283_ext
	return nil
}
func (ϟa *GlGetQueryObjectui64vEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_284_ext := ExtensionId_GL_EXT_disjoint_timer_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_ = requiresExtension_284_ext
	return nil
}
func (ϟa *GlGetQueryObjectuivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_285_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_286_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = requiresExtension_285_ext, requiresExtension_286_ext
	return nil
}
func (ϟa *GlGetQueryivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_287_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_288_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = requiresExtension_287_ext, requiresExtension_288_ext
	return nil
}
func (ϟa *GlGetSamplerParameterIivOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_289_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_289_ext
	return nil
}
func (ϟa *GlGetSamplerParameterIuivOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_290_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_290_ext
	return nil
}
func (ϟa *GlGetSyncivAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_291_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_291_ext
	return nil
}
func (ϟa *GlGetTexParameterIivOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_292_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_292_ext
	return nil
}
func (ϟa *GlGetTexParameterIuivOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_293_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_293_ext
	return nil
}
func (ϟa *GlGetTextureHandleNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_294_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_294_ext
	return nil
}
func (ϟa *GlGetTextureSamplerHandleNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_295_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_295_ext
	return nil
}
func (ϟa *GlGetTranslatedShaderSourceANGLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_296_ext := ExtensionId_GL_ANGLE_translated_shader_source // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_296_ext
	return nil
}
func (ϟa *GlGetnUniformfvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_297_ext := ExtensionId_GL_EXT_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_297_ext
	return nil
}
func (ϟa *GlGetnUniformfvKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_298_ext := ExtensionId_GL_KHR_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_298_ext
	return nil
}
func (ϟa *GlGetnUniformivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_299_ext := ExtensionId_GL_EXT_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_299_ext
	return nil
}
func (ϟa *GlGetnUniformivKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_300_ext := ExtensionId_GL_KHR_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_300_ext
	return nil
}
func (ϟa *GlGetnUniformuivKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_301_ext := ExtensionId_GL_KHR_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_301_ext
	return nil
}
func (ϟa *GlInsertEventMarkerEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_302_ext := ExtensionId_GL_EXT_debug_marker // ExtensionId
	if (ϟa.Length) > (GLsizei(int32(0))) {
		ϟa.Marker.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Length), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	} else {
		_ = strings.TrimRight(string(Charᵖ(ϟa.Marker).StringSlice(ϟs, ϟd, ϟl).Read(ϟa, ϟs, ϟd, ϟl, ϟb)), "\x00")
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_302_ext
	return nil
}
func (ϟa *GlInterpolatePathsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_303_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_303_ext
	return nil
}
func (ϟa *GlIsEnablediOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_304_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_304_ext
	return nil
}
func (ϟa *GlIsEnablediNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_305_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_305_ext
	return nil
}
func (ϟa *GlIsFenceNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_306_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_306_ext
	return nil
}
func (ϟa *GlIsImageHandleResidentNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_307_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_307_ext
	return nil
}
func (ϟa *GlIsPathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_308_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_308_ext
	return nil
}
func (ϟa *GlIsPointInFillPathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_309_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_309_ext
	return nil
}
func (ϟa *GlIsPointInStrokePathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_310_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_310_ext
	return nil
}
func (ϟa *GlIsProgramPipelineEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_311_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_311_ext
	return nil
}
func (ϟa *GlIsQueryEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_312_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_313_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                            // Contextʳ
	GetContext_314_result := context                                        // Contextʳ
	ctx := GetContext_314_result                                            // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.Queries.Contains(ϟa.Query) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = requiresExtension_312_ext, requiresExtension_313_ext, context, GetContext_314_result, ctx
	return nil
}
func (ϟa *GlIsSyncAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_315_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_315_ext
	return nil
}
func (ϟa *GlIsTextureHandleResidentNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_316_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_316_ext
	return nil
}
func (ϟa *GlIsVertexArrayOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_317_ext := ExtensionId_GL_OES_vertex_array_object // ExtensionId
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                        // Contextʳ
	GetContext_318_result := context                                    // Contextʳ
	ctx := GetContext_318_result                                        // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.VertexArrays.Contains(ϟa.Array) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _ = requiresExtension_317_ext, context, GetContext_318_result, ctx
	return nil
}
func (ϟa *GlLabelObjectEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_319_ext := ExtensionId_GL_EXT_debug_label // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_319_ext
	return nil
}
func (ϟa *GlMakeImageHandleNonResidentNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_320_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_320_ext
	return nil
}
func (ϟa *GlMakeImageHandleResidentNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_321_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_321_ext
	return nil
}
func (ϟa *GlMakeTextureHandleNonResidentNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_322_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_322_ext
	return nil
}
func (ϟa *GlMakeTextureHandleResidentNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_323_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_323_ext
	return nil
}
func (ϟa *GlMapBufferOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_324_ext := ExtensionId_GL_OES_mapbuffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_324_ext
	return nil
}
func (ϟa *GlMapBufferRangeEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_325_ext := ExtensionId_GL_EXT_map_buffer_range // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_325_ext
	return nil
}
func (ϟa *GlMatrixLoad3x2fNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_326_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_326_ext
	return nil
}
func (ϟa *GlMatrixLoad3x3fNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_327_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_327_ext
	return nil
}
func (ϟa *GlMatrixLoadTranspose3x3fNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_328_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_328_ext
	return nil
}
func (ϟa *GlMatrixMult3x2fNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_329_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_329_ext
	return nil
}
func (ϟa *GlMatrixMult3x3fNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_330_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_330_ext
	return nil
}
func (ϟa *GlMatrixMultTranspose3x3fNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_331_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_331_ext
	return nil
}
func (ϟa *GlMultiDrawArraysEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_332_ext := ExtensionId_GL_EXT_multi_draw_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_332_ext
	return nil
}
func (ϟa *GlMultiDrawArraysIndirectEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_333_ext := ExtensionId_GL_EXT_multi_draw_indirect // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_333_ext
	return nil
}
func (ϟa *GlMultiDrawElementsBaseVertexEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_334_ext := ExtensionId_GL_EXT_draw_elements_base_vertex // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_334_ext
	return nil
}
func (ϟa *GlMultiDrawElementsBaseVertexOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_335_ext := ExtensionId_GL_OES_draw_elements_base_vertex // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_335_ext
	return nil
}
func (ϟa *GlMultiDrawElementsEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_336_ext := ExtensionId_GL_EXT_multi_draw_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_336_ext
	return nil
}
func (ϟa *GlMultiDrawElementsIndirectEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_337_ext := ExtensionId_GL_EXT_multi_draw_indirect // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_337_ext
	return nil
}
func (ϟa *GlPatchParameteriOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_338_ext := ExtensionId_GL_OES_tessellation_shader // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_338_ext
	return nil
}
func (ϟa *GlPathCommandsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_339_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_339_ext
	return nil
}
func (ϟa *GlPathCoordsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_340_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_340_ext
	return nil
}
func (ϟa *GlPathCoverDepthFuncNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_341_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_341_ext
	return nil
}
func (ϟa *GlPathDashArrayNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_342_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_342_ext
	return nil
}
func (ϟa *GlPathGlyphIndexArrayNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_343_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_343_ext
	return nil
}
func (ϟa *GlPathGlyphIndexRangeNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_344_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_344_ext
	return nil
}
func (ϟa *GlPathGlyphRangeNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_345_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_345_ext
	return nil
}
func (ϟa *GlPathGlyphsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_346_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_346_ext
	return nil
}
func (ϟa *GlPathMemoryGlyphIndexArrayNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_347_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_347_ext
	return nil
}
func (ϟa *GlPathParameterfNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_348_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_348_ext
	return nil
}
func (ϟa *GlPathParameterfvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_349_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_349_ext
	return nil
}
func (ϟa *GlPathParameteriNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_350_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_350_ext
	return nil
}
func (ϟa *GlPathParameterivNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_351_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_351_ext
	return nil
}
func (ϟa *GlPathStencilDepthOffsetNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_352_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_352_ext
	return nil
}
func (ϟa *GlPathStencilFuncNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_353_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_353_ext
	return nil
}
func (ϟa *GlPathStringNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_354_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_354_ext
	return nil
}
func (ϟa *GlPathSubCommandsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_355_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_355_ext
	return nil
}
func (ϟa *GlPathSubCoordsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_356_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_356_ext
	return nil
}
func (ϟa *GlPointAlongPathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_357_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_357_ext
	return nil
}
func (ϟa *GlPolygonModeNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_358_ext := ExtensionId_GL_NV_polygon_mode // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_358_ext
	return nil
}
func (ϟa *GlPopGroupMarkerEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_359_ext := ExtensionId_GL_EXT_debug_marker // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_359_ext
	return nil
}
func (ϟa *GlPrimitiveBoundingBoxOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_360_ext := ExtensionId_GL_OES_primitive_bounding_box // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_360_ext
	return nil
}
func (ϟa *GlProgramBinaryOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_361_ext := ExtensionId_GL_OES_get_program_binary // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_361_ext
	return nil
}
func (ϟa *GlProgramParameteriEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_362_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_362_ext
	return nil
}
func (ϟa *GlProgramPathFragmentInputGenNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_363_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_363_ext
	return nil
}
func (ϟa *GlProgramUniform1fEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_364_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_364_ext
	return nil
}
func (ϟa *GlProgramUniform1fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_365_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_365_ext
	return nil
}
func (ϟa *GlProgramUniform1iEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_366_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_366_ext
	return nil
}
func (ϟa *GlProgramUniform1ivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_367_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_367_ext
	return nil
}
func (ϟa *GlProgramUniform1uiEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_368_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_368_ext
	return nil
}
func (ϟa *GlProgramUniform1uivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_369_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_369_ext
	return nil
}
func (ϟa *GlProgramUniform2fEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_370_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_370_ext
	return nil
}
func (ϟa *GlProgramUniform2fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_371_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_371_ext
	return nil
}
func (ϟa *GlProgramUniform2iEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_372_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_372_ext
	return nil
}
func (ϟa *GlProgramUniform2ivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_373_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_373_ext
	return nil
}
func (ϟa *GlProgramUniform2uiEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_374_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_374_ext
	return nil
}
func (ϟa *GlProgramUniform2uivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_375_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_375_ext
	return nil
}
func (ϟa *GlProgramUniform3fEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_376_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_376_ext
	return nil
}
func (ϟa *GlProgramUniform3fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_377_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_377_ext
	return nil
}
func (ϟa *GlProgramUniform3iEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_378_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_378_ext
	return nil
}
func (ϟa *GlProgramUniform3ivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_379_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_379_ext
	return nil
}
func (ϟa *GlProgramUniform3uiEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_380_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_380_ext
	return nil
}
func (ϟa *GlProgramUniform3uivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_381_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_381_ext
	return nil
}
func (ϟa *GlProgramUniform4fEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_382_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_382_ext
	return nil
}
func (ϟa *GlProgramUniform4fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_383_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_383_ext
	return nil
}
func (ϟa *GlProgramUniform4iEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_384_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_384_ext
	return nil
}
func (ϟa *GlProgramUniform4ivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_385_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_385_ext
	return nil
}
func (ϟa *GlProgramUniform4uiEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_386_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_386_ext
	return nil
}
func (ϟa *GlProgramUniform4uivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_387_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_387_ext
	return nil
}
func (ϟa *GlProgramUniformHandleui64NV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_388_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_388_ext
	return nil
}
func (ϟa *GlProgramUniformHandleui64vNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_389_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_389_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix2fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_390_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_390_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix2x3fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_391_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_391_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix2x4fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_392_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_392_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix3fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_393_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_393_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix3x2fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_394_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_394_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix3x4fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_395_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_395_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix4fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_396_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_396_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix4x2fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_397_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_397_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix4x3fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_398_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_398_ext
	return nil
}
func (ϟa *GlPushGroupMarkerEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_399_ext := ExtensionId_GL_EXT_debug_marker // ExtensionId
	if (ϟa.Length) > (GLsizei(int32(0))) {
		ϟa.Marker.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Length), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	} else {
		_ = strings.TrimRight(string(Charᵖ(ϟa.Marker).StringSlice(ϟs, ϟd, ϟl).Read(ϟa, ϟs, ϟd, ϟl, ϟb)), "\x00")
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_399_ext
	return nil
}
func (ϟa *GlQueryCounterEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_400_ext := ExtensionId_GL_EXT_disjoint_timer_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_400_ext
	return nil
}
func (ϟa *GlReadBufferIndexedEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_401_ext := ExtensionId_GL_EXT_multiview_draw_buffers // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_401_ext
	return nil
}
func (ϟa *GlReadBufferNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_402_ext := ExtensionId_GL_NV_read_buffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_402_ext
	return nil
}
func (ϟa *GlReadnPixelsEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_403_ext := ExtensionId_GL_EXT_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_403_ext
	return nil
}
func (ϟa *GlReadnPixelsKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_404_ext := ExtensionId_GL_KHR_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_404_ext
	return nil
}
func (ϟa *GlRenderbufferStorageMultisampleANGLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_405_ext := ExtensionId_GL_ANGLE_framebuffer_multisample // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_405_ext
	return nil
}
func (ϟa *GlRenderbufferStorageMultisampleAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_406_ext := ExtensionId_GL_APPLE_framebuffer_multisample // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_406_ext
	return nil
}
func (ϟa *GlRenderbufferStorageMultisampleEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_407_ext := ExtensionId_GL_EXT_multisampled_render_to_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_407_ext
	return nil
}
func (ϟa *GlRenderbufferStorageMultisampleIMG) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_408_ext := ExtensionId_GL_IMG_multisampled_render_to_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_408_ext
	return nil
}
func (ϟa *GlRenderbufferStorageMultisampleNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_409_ext := ExtensionId_GL_NV_framebuffer_multisample // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_409_ext
	return nil
}
func (ϟa *GlResolveMultisampleFramebufferAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_410_ext := ExtensionId_GL_APPLE_framebuffer_multisample // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_410_ext
	return nil
}
func (ϟa *GlSamplerParameterIivOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_411_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_411_ext
	return nil
}
func (ϟa *GlSamplerParameterIuivOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_412_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_412_ext
	return nil
}
func (ϟa *GlScissorArrayvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_413_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_413_ext
	return nil
}
func (ϟa *GlScissorIndexedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_414_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_414_ext
	return nil
}
func (ϟa *GlScissorIndexedvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_415_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_415_ext
	return nil
}
func (ϟa *GlSelectPerfMonitorCountersAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_416_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_416_ext
	return nil
}
func (ϟa *GlSetFenceNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_417_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_417_ext
	return nil
}
func (ϟa *GlStartTilingQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_418_ext := ExtensionId_GL_QCOM_tiled_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_418_ext
	return nil
}
func (ϟa *GlStencilFillPathInstancedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_419_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_419_ext
	return nil
}
func (ϟa *GlStencilFillPathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_420_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_420_ext
	return nil
}
func (ϟa *GlStencilStrokePathInstancedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_421_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_421_ext
	return nil
}
func (ϟa *GlStencilStrokePathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_422_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_422_ext
	return nil
}
func (ϟa *GlStencilThenCoverFillPathInstancedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_423_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_423_ext
	return nil
}
func (ϟa *GlStencilThenCoverFillPathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_424_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_424_ext
	return nil
}
func (ϟa *GlStencilThenCoverStrokePathInstancedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_425_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_425_ext
	return nil
}
func (ϟa *GlStencilThenCoverStrokePathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_426_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_426_ext
	return nil
}
func (ϟa *GlTestFenceNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_427_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_427_ext
	return nil
}
func (ϟa *GlTexBufferOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_428_ext := ExtensionId_GL_OES_texture_buffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_428_ext
	return nil
}
func (ϟa *GlTexBufferRangeOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_429_ext := ExtensionId_GL_OES_texture_buffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_429_ext
	return nil
}
func (ϟa *GlTexImage3DOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_430_ext := ExtensionId_GL_OES_texture_3D // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_430_ext
	return nil
}
func (ϟa *GlTexPageCommitmentARB) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_431_ext := ExtensionId_GL_EXT_sparse_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_431_ext
	return nil
}
func (ϟa *GlTexParameterIivOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_432_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_432_ext
	return nil
}
func (ϟa *GlTexParameterIuivOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_433_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_433_ext
	return nil
}
func (ϟa *GlTexStorage1DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_434_ext := ExtensionId_GL_EXT_texture_storage // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_434_ext
	return nil
}
func (ϟa *GlTexStorage2DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_435_ext := ExtensionId_GL_EXT_texture_storage // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_435_ext
	return nil
}
func (ϟa *GlTexStorage3DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_436_ext := ExtensionId_GL_EXT_texture_storage // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_436_ext
	return nil
}
func (ϟa *GlTexSubImage3DOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_437_ext := ExtensionId_GL_OES_texture_3D // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_437_ext
	return nil
}
func (ϟa *GlTextureStorage1DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_438_ext := ExtensionId_GL_EXT_texture_storage // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_438_ext
	return nil
}
func (ϟa *GlTextureStorage2DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_439_ext := ExtensionId_GL_EXT_texture_storage // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_439_ext
	return nil
}
func (ϟa *GlTextureStorage3DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_440_ext := ExtensionId_GL_EXT_texture_storage // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_440_ext
	return nil
}
func (ϟa *GlTextureViewEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_441_ext := ExtensionId_GL_EXT_texture_view // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_441_ext
	return nil
}
func (ϟa *GlTextureViewOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_442_ext := ExtensionId_GL_OES_texture_view // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_442_ext
	return nil
}
func (ϟa *GlTransformPathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_443_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_443_ext
	return nil
}
func (ϟa *GlUniformHandleui64NV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_444_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_444_ext
	return nil
}
func (ϟa *GlUniformHandleui64vNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_445_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_445_ext
	return nil
}
func (ϟa *GlUniformMatrix2x3fvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_446_ext := ExtensionId_GL_NV_non_square_matrices // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_446_ext
	return nil
}
func (ϟa *GlUniformMatrix2x4fvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_447_ext := ExtensionId_GL_NV_non_square_matrices // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_447_ext
	return nil
}
func (ϟa *GlUniformMatrix3x2fvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_448_ext := ExtensionId_GL_NV_non_square_matrices // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_448_ext
	return nil
}
func (ϟa *GlUniformMatrix3x4fvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_449_ext := ExtensionId_GL_NV_non_square_matrices // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_449_ext
	return nil
}
func (ϟa *GlUniformMatrix4x2fvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_450_ext := ExtensionId_GL_NV_non_square_matrices // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_450_ext
	return nil
}
func (ϟa *GlUniformMatrix4x3fvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_451_ext := ExtensionId_GL_NV_non_square_matrices // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_451_ext
	return nil
}
func (ϟa *GlUnmapBufferOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_452_ext := ExtensionId_GL_OES_mapbuffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_452_ext
	return nil
}
func (ϟa *GlUseProgramStagesEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_453_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_453_ext
	return nil
}
func (ϟa *GlValidateProgramPipelineEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_454_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_454_ext
	return nil
}
func (ϟa *GlVertexAttribDivisorANGLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_455_ext := ExtensionId_GL_ANGLE_instanced_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_455_ext
	return nil
}
func (ϟa *GlVertexAttribDivisorEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_456_ext := ExtensionId_GL_EXT_instanced_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_456_ext
	return nil
}
func (ϟa *GlVertexAttribDivisorNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_457_ext := ExtensionId_GL_NV_instanced_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_457_ext
	return nil
}
func (ϟa *GlViewportArrayvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_458_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_458_ext
	return nil
}
func (ϟa *GlViewportIndexedfNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_459_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_459_ext
	return nil
}
func (ϟa *GlViewportIndexedfvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_460_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_460_ext
	return nil
}
func (ϟa *GlWaitSyncAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_461_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_461_ext
	return nil
}
func (ϟa *GlWeightPathsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_462_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_462_ext
	return nil
}
func (ϟa *GlCoverageModulationNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_463_ext := ExtensionId_GL_NV_framebuffer_mixed_samples // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_463_ext
	return nil
}
func (ϟa *GlCoverageModulationTableNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_464_ext := ExtensionId_GL_NV_framebuffer_mixed_samples // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_464_ext
	return nil
}
func (ϟa *GlFragmentCoverageColorNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_465_ext := ExtensionId_GL_NV_fragment_coverage_to_color // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_465_ext
	return nil
}
func (ϟa *GlFramebufferSampleLocationsfvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_466_ext := ExtensionId_GL_NV_sample_locations // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_466_ext
	return nil
}
func (ϟa *GlGetCoverageModulationTableNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_467_ext := ExtensionId_GL_NV_framebuffer_mixed_samples // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_467_ext
	return nil
}
func (ϟa *GlNamedFramebufferSampleLocationsfvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_468_ext := ExtensionId_GL_NV_sample_locations // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_468_ext
	return nil
}
func (ϟa *GlRasterSamplesEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_469_ext := ExtensionId_GL_EXT_raster_multisample       // ExtensionId
	requiresExtension_470_ext := ExtensionId_GL_EXT_texture_filter_minmax    // ExtensionId
	requiresExtension_471_ext := ExtensionId_GL_NV_framebuffer_mixed_samples // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = requiresExtension_469_ext, requiresExtension_470_ext, requiresExtension_471_ext
	return nil
}
func (ϟa *GlResolveDepthValuesNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_472_ext := ExtensionId_GL_NV_sample_locations // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_472_ext
	return nil
}
func (ϟa *GlSubpixelPrecisionBiasNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_473_ext := ExtensionId_GL_NV_conservative_raster // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_473_ext
	return nil
}
func (ϟa *GlBlendColor) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_474_major := uint32(2)    // u32
	minRequiredVersion_474_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_475_result := context             // Contextʳ
	ctx := GetContext_475_result                 // Contextʳ
	ctx.Blending.BlendColor = func() Color {
		s := Color{}
		s.Init()
		s.Red = ϟa.Red
		s.Green = ϟa.Green
		s.Blue = ϟa.Blue
		s.Alpha = ϟa.Alpha
		return s
	}()
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_474_major, minRequiredVersion_474_minor, context, GetContext_475_result, ctx
	return nil
}
func (ϟa *GlBlendEquation) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_476_major := uint32(2) // u32
	minRequiredVersion_476_minor := uint32(0) // u32
	switch ϟa.Equation {
	case GLenum_GL_FUNC_ADD, GLenum_GL_FUNC_REVERSE_SUBTRACT, GLenum_GL_FUNC_SUBTRACT:
	case GLenum_GL_MAX, GLenum_GL_MIN:
		minRequiredVersion_477_major := uint32(3) // u32
		minRequiredVersion_477_minor := uint32(0) // u32
		_, _ = minRequiredVersion_477_major, minRequiredVersion_477_minor
	default:
		v := ϟa.Equation
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_479_result := context             // Contextʳ
	ctx := GetContext_479_result                 // Contextʳ
	ctx.Blending.BlendEquationRgb = ϟa.Equation
	ctx.Blending.BlendEquationAlpha = ϟa.Equation
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_476_major, minRequiredVersion_476_minor, context, GetContext_479_result, ctx
	return nil
}
func (ϟa *GlBlendEquationSeparate) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_480_major := uint32(2) // u32
	minRequiredVersion_480_minor := uint32(0) // u32
	switch ϟa.Rgb {
	case GLenum_GL_FUNC_ADD, GLenum_GL_FUNC_REVERSE_SUBTRACT, GLenum_GL_FUNC_SUBTRACT:
	case GLenum_GL_MAX, GLenum_GL_MIN:
		minRequiredVersion_481_major := uint32(3) // u32
		minRequiredVersion_481_minor := uint32(0) // u32
		_, _ = minRequiredVersion_481_major, minRequiredVersion_481_minor
	default:
		v := ϟa.Rgb
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Alpha {
	case GLenum_GL_FUNC_ADD, GLenum_GL_FUNC_REVERSE_SUBTRACT, GLenum_GL_FUNC_SUBTRACT:
	case GLenum_GL_MAX, GLenum_GL_MIN:
		minRequiredVersion_483_major := uint32(3) // u32
		minRequiredVersion_483_minor := uint32(0) // u32
		_, _ = minRequiredVersion_483_major, minRequiredVersion_483_minor
	default:
		v := ϟa.Alpha
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_485_result := context             // Contextʳ
	ctx := GetContext_485_result                 // Contextʳ
	ctx.Blending.BlendEquationRgb = ϟa.Rgb
	ctx.Blending.BlendEquationAlpha = ϟa.Alpha
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_480_major, minRequiredVersion_480_minor, context, GetContext_485_result, ctx
	return nil
}
func (ϟa *GlBlendFunc) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_486_major := uint32(2) // u32
	minRequiredVersion_486_minor := uint32(0) // u32
	switch ϟa.SrcFactor {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		v := ϟa.SrcFactor
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.DstFactor {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		v := ϟa.DstFactor
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_489_result := context             // Contextʳ
	ctx := GetContext_489_result                 // Contextʳ
	ctx.Blending.SrcRgbBlendFactor = ϟa.SrcFactor
	ctx.Blending.SrcAlphaBlendFactor = ϟa.SrcFactor
	ctx.Blending.DstRgbBlendFactor = ϟa.DstFactor
	ctx.Blending.DstAlphaBlendFactor = ϟa.DstFactor
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_486_major, minRequiredVersion_486_minor, context, GetContext_489_result, ctx
	return nil
}
func (ϟa *GlBlendFuncSeparate) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_490_major := uint32(2) // u32
	minRequiredVersion_490_minor := uint32(0) // u32
	switch ϟa.SrcFactorRgb {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		v := ϟa.SrcFactorRgb
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.DstFactorRgb {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		v := ϟa.DstFactorRgb
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.SrcFactorAlpha {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		v := ϟa.SrcFactorAlpha
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.DstFactorAlpha {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		v := ϟa.DstFactorAlpha
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_495_result := context             // Contextʳ
	ctx := GetContext_495_result                 // Contextʳ
	ctx.Blending.SrcRgbBlendFactor = ϟa.SrcFactorRgb
	ctx.Blending.DstRgbBlendFactor = ϟa.DstFactorRgb
	ctx.Blending.SrcAlphaBlendFactor = ϟa.SrcFactorAlpha
	ctx.Blending.DstAlphaBlendFactor = ϟa.DstFactorAlpha
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_490_major, minRequiredVersion_490_minor, context, GetContext_495_result, ctx
	return nil
}
func (ϟa *GlDepthFunc) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_496_major := uint32(2) // u32
	minRequiredVersion_496_minor := uint32(0) // u32
	switch ϟa.Function {
	case GLenum_GL_ALWAYS, GLenum_GL_EQUAL, GLenum_GL_GEQUAL, GLenum_GL_GREATER, GLenum_GL_LEQUAL, GLenum_GL_LESS, GLenum_GL_NEVER, GLenum_GL_NOTEQUAL:
	default:
		v := ϟa.Function
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_498_result := context             // Contextʳ
	ctx := GetContext_498_result                 // Contextʳ
	ctx.Rasterizing.DepthTestFunction = ϟa.Function
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_496_major, minRequiredVersion_496_minor, context, GetContext_498_result, ctx
	return nil
}
func (ϟa *GlSampleCoverage) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_499_major := uint32(2)    // u32
	minRequiredVersion_499_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_500_result := context             // Contextʳ
	ctx := GetContext_500_result                 // Contextʳ
	ctx.Rasterizing.SampleCoverageValue = ϟa.Value
	ctx.Rasterizing.SampleCoverageInvert = ϟa.Invert
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_499_major, minRequiredVersion_499_minor, context, GetContext_500_result, ctx
	return nil
}
func (ϟa *GlSampleMaski) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_501_major := uint32(3) // u32
	minRequiredVersion_501_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_501_major, minRequiredVersion_501_minor
	return nil
}
func (ϟa *GlScissor) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_502_major := uint32(2)    // u32
	minRequiredVersion_502_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_503_result := context             // Contextʳ
	ctx := GetContext_503_result                 // Contextʳ
	ctx.Rasterizing.Scissor = func() Rect {
		s := Rect{}
		s.Init()
		s.X = ϟa.X
		s.Y = ϟa.Y
		s.Width = ϟa.Width
		s.Height = ϟa.Height
		return s
	}()
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_502_major, minRequiredVersion_502_minor, context, GetContext_503_result, ctx
	return nil
}
func (ϟa *GlStencilFunc) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_504_major := uint32(2) // u32
	minRequiredVersion_504_minor := uint32(0) // u32
	switch ϟa.Func {
	case GLenum_GL_ALWAYS, GLenum_GL_EQUAL, GLenum_GL_GEQUAL, GLenum_GL_GREATER, GLenum_GL_LEQUAL, GLenum_GL_LESS, GLenum_GL_NEVER, GLenum_GL_NOTEQUAL:
	default:
		v := ϟa.Func
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_504_major, minRequiredVersion_504_minor
	return nil
}
func (ϟa *GlStencilFuncSeparate) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_506_major := uint32(2) // u32
	minRequiredVersion_506_minor := uint32(0) // u32
	switch ϟa.Face {
	case GLenum_GL_BACK, GLenum_GL_FRONT, GLenum_GL_FRONT_AND_BACK:
	default:
		v := ϟa.Face
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Function {
	case GLenum_GL_ALWAYS, GLenum_GL_EQUAL, GLenum_GL_GEQUAL, GLenum_GL_GREATER, GLenum_GL_LEQUAL, GLenum_GL_LESS, GLenum_GL_NEVER, GLenum_GL_NOTEQUAL:
	default:
		v := ϟa.Function
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_506_major, minRequiredVersion_506_minor
	return nil
}
func (ϟa *GlStencilOp) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_509_major := uint32(2) // u32
	minRequiredVersion_509_minor := uint32(0) // u32
	switch ϟa.Fail {
	case GLenum_GL_DECR, GLenum_GL_DECR_WRAP, GLenum_GL_INCR, GLenum_GL_INCR_WRAP, GLenum_GL_INVERT, GLenum_GL_KEEP, GLenum_GL_REPLACE, GLenum_GL_ZERO:
	default:
		v := ϟa.Fail
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Zfail {
	case GLenum_GL_DECR, GLenum_GL_DECR_WRAP, GLenum_GL_INCR, GLenum_GL_INCR_WRAP, GLenum_GL_INVERT, GLenum_GL_KEEP, GLenum_GL_REPLACE, GLenum_GL_ZERO:
	default:
		v := ϟa.Zfail
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Zpass {
	case GLenum_GL_DECR, GLenum_GL_DECR_WRAP, GLenum_GL_INCR, GLenum_GL_INCR_WRAP, GLenum_GL_INVERT, GLenum_GL_KEEP, GLenum_GL_REPLACE, GLenum_GL_ZERO:
	default:
		v := ϟa.Zpass
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_509_major, minRequiredVersion_509_minor
	return nil
}
func (ϟa *GlStencilOpSeparate) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_513_major := uint32(2) // u32
	minRequiredVersion_513_minor := uint32(0) // u32
	switch ϟa.Face {
	case GLenum_GL_BACK, GLenum_GL_FRONT, GLenum_GL_FRONT_AND_BACK:
	default:
		v := ϟa.Face
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.StencilFail {
	case GLenum_GL_DECR, GLenum_GL_DECR_WRAP, GLenum_GL_INCR, GLenum_GL_INCR_WRAP, GLenum_GL_INVERT, GLenum_GL_KEEP, GLenum_GL_REPLACE, GLenum_GL_ZERO:
	default:
		v := ϟa.StencilFail
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.StencilPassDepthFail {
	case GLenum_GL_DECR, GLenum_GL_DECR_WRAP, GLenum_GL_INCR, GLenum_GL_INCR_WRAP, GLenum_GL_INVERT, GLenum_GL_KEEP, GLenum_GL_REPLACE, GLenum_GL_ZERO:
	default:
		v := ϟa.StencilPassDepthFail
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.StencilPassDepthPass {
	case GLenum_GL_DECR, GLenum_GL_DECR_WRAP, GLenum_GL_INCR, GLenum_GL_INCR_WRAP, GLenum_GL_INVERT, GLenum_GL_KEEP, GLenum_GL_REPLACE, GLenum_GL_ZERO:
	default:
		v := ϟa.StencilPassDepthPass
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_513_major, minRequiredVersion_513_minor
	return nil
}
func (ϟa *GlBindFramebuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_518_major := uint32(2) // u32
	minRequiredVersion_518_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_FRAMEBUFFER:
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
		minRequiredVersion_519_major := uint32(3) // u32
		minRequiredVersion_519_minor := uint32(0) // u32
		_, _ = minRequiredVersion_519_major, minRequiredVersion_519_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_521_result := context             // Contextʳ
	ctx := GetContext_521_result                 // Contextʳ
	if !(ctx.Instances.Framebuffers.Contains(ϟa.Framebuffer)) {
		ctx.Instances.Framebuffers[ϟa.Framebuffer] = func() *Framebuffer {
			s := &Framebuffer{}
			s.Init()
			return s
		}()
	}
	if (ϟa.Target) == (GLenum_GL_FRAMEBUFFER) {
		ctx.BoundFramebuffers[GLenum_GL_READ_FRAMEBUFFER] = ϟa.Framebuffer
		ctx.BoundFramebuffers[GLenum_GL_DRAW_FRAMEBUFFER] = ϟa.Framebuffer
	} else {
		ctx.BoundFramebuffers[ϟa.Target] = ϟa.Framebuffer
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_518_major, minRequiredVersion_518_minor, context, GetContext_521_result, ctx
	return nil
}
func (ϟa *GlBindRenderbuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_522_major := uint32(2) // u32
	minRequiredVersion_522_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_RENDERBUFFER:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_524_result := context             // Contextʳ
	ctx := GetContext_524_result                 // Contextʳ
	if !(ctx.Instances.Renderbuffers.Contains(ϟa.Renderbuffer)) {
		ctx.Instances.Renderbuffers[ϟa.Renderbuffer] = func() *Renderbuffer {
			s := &Renderbuffer{}
			s.Init()
			return s
		}()
	}
	ctx.BoundRenderbuffers[ϟa.Target] = ϟa.Renderbuffer
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_522_major, minRequiredVersion_522_minor, context, GetContext_524_result, ctx
	return nil
}
func (ϟa *GlBlitFramebuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_525_major := uint32(3)                                                                                                // u32
	minRequiredVersion_525_minor := uint32(0)                                                                                                // u32
	supportsBits_526_seenBits := ϟa.Mask                                                                                                     // GLbitfield
	supportsBits_526_validBits := (GLbitfield_GL_COLOR_BUFFER_BIT) | ((GLbitfield_GL_DEPTH_BUFFER_BIT) | (GLbitfield_GL_STENCIL_BUFFER_BIT)) // GLbitfield
	if (GLbitfield_GL_COLOR_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	if (GLbitfield_GL_DEPTH_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	if (GLbitfield_GL_STENCIL_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	switch ϟa.Filter {
	case GLenum_GL_LINEAR, GLenum_GL_NEAREST:
	default:
		v := ϟa.Filter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = minRequiredVersion_525_major, minRequiredVersion_525_minor, supportsBits_526_seenBits, supportsBits_526_validBits
	return nil
}
func (ϟa *GlCheckFramebufferStatus) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_528_major := uint32(2) // u32
	minRequiredVersion_528_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_FRAMEBUFFER:
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
		minRequiredVersion_529_major := uint32(3) // u32
		minRequiredVersion_529_minor := uint32(0) // u32
		_, _ = minRequiredVersion_529_major, minRequiredVersion_529_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_528_major, minRequiredVersion_528_minor
	return nil
}
func (ϟa *GlClear) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_531_major := uint32(2)                                                                                                // u32
	minRequiredVersion_531_minor := uint32(0)                                                                                                // u32
	supportsBits_532_seenBits := ϟa.Mask                                                                                                     // GLbitfield
	supportsBits_532_validBits := (GLbitfield_GL_COLOR_BUFFER_BIT) | ((GLbitfield_GL_DEPTH_BUFFER_BIT) | (GLbitfield_GL_STENCIL_BUFFER_BIT)) // GLbitfield
	if (GLbitfield_GL_COLOR_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	if (GLbitfield_GL_DEPTH_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	if (GLbitfield_GL_STENCIL_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	if (GLbitfield_GL_COLOR_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = minRequiredVersion_531_major, minRequiredVersion_531_minor, supportsBits_532_seenBits, supportsBits_532_validBits
	return nil
}
func (ϟa *GlClearBufferfi) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_533_major := uint32(3) // u32
	minRequiredVersion_533_minor := uint32(0) // u32
	switch ϟa.Buffer {
	case GLenum_GL_DEPTH_STENCIL:
	default:
		v := ϟa.Buffer
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_533_major, minRequiredVersion_533_minor
	return nil
}
func (ϟa *GlClearBufferfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_535_major := uint32(3) // u32
	minRequiredVersion_535_minor := uint32(0) // u32
	switch ϟa.Buffer {
	case GLenum_GL_COLOR, GLenum_GL_DEPTH:
	default:
		v := ϟa.Buffer
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_535_major, minRequiredVersion_535_minor
	return nil
}
func (ϟa *GlClearBufferiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_537_major := uint32(3) // u32
	minRequiredVersion_537_minor := uint32(0) // u32
	switch ϟa.Buffer {
	case GLenum_GL_COLOR, GLenum_GL_STENCIL:
	default:
		v := ϟa.Buffer
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_537_major, minRequiredVersion_537_minor
	return nil
}
func (ϟa *GlClearBufferuiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_539_major := uint32(3) // u32
	minRequiredVersion_539_minor := uint32(0) // u32
	switch ϟa.Buffer {
	case GLenum_GL_COLOR:
	default:
		v := ϟa.Buffer
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_539_major, minRequiredVersion_539_minor
	return nil
}
func (ϟa *GlClearColor) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_541_major := uint32(2)    // u32
	minRequiredVersion_541_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_542_result := context             // Contextʳ
	ctx := GetContext_542_result                 // Contextʳ
	ctx.Clearing.ClearColor = func() Color {
		s := Color{}
		s.Init()
		s.Red = ϟa.R
		s.Green = ϟa.G
		s.Blue = ϟa.B
		s.Alpha = ϟa.A
		return s
	}()
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_541_major, minRequiredVersion_541_minor, context, GetContext_542_result, ctx
	return nil
}
func (ϟa *GlClearDepthf) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_543_major := uint32(2)    // u32
	minRequiredVersion_543_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_544_result := context             // Contextʳ
	ctx := GetContext_544_result                 // Contextʳ
	ctx.Clearing.ClearDepth = ϟa.Depth
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_543_major, minRequiredVersion_543_minor, context, GetContext_544_result, ctx
	return nil
}
func (ϟa *GlClearStencil) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_545_major := uint32(2)    // u32
	minRequiredVersion_545_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_546_result := context             // Contextʳ
	ctx := GetContext_546_result                 // Contextʳ
	ctx.Clearing.ClearStencil = ϟa.Stencil
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_545_major, minRequiredVersion_545_minor, context, GetContext_546_result, ctx
	return nil
}
func (ϟa *GlColorMask) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_547_major := uint32(2)    // u32
	minRequiredVersion_547_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_548_result := context             // Contextʳ
	ctx := GetContext_548_result                 // Contextʳ
	ctx.Rasterizing.ColorMaskRed = ϟa.Red
	ctx.Rasterizing.ColorMaskGreen = ϟa.Green
	ctx.Rasterizing.ColorMaskBlue = ϟa.Blue
	ctx.Rasterizing.ColorMaskAlpha = ϟa.Alpha
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_547_major, minRequiredVersion_547_minor, context, GetContext_548_result, ctx
	return nil
}
func (ϟa *GlDeleteFramebuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_549_major := uint32(2)                                   // u32
	minRequiredVersion_549_minor := uint32(0)                                   // u32
	f := ϟa.Framebuffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // FramebufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                // Contextʳ
	GetContext_550_result := context                                            // Contextʳ
	ctx := GetContext_550_result                                                // Contextʳ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.Framebuffers, f.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_549_major, minRequiredVersion_549_minor, f, context, GetContext_550_result, ctx
	return nil
}
func (ϟa *GlDeleteRenderbuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_551_major := uint32(2)                                    // u32
	minRequiredVersion_551_minor := uint32(0)                                    // u32
	r := ϟa.Renderbuffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // RenderbufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                 // Contextʳ
	GetContext_552_result := context                                             // Contextʳ
	ctx := GetContext_552_result                                                 // Contextʳ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.Renderbuffers, r.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_551_major, minRequiredVersion_551_minor, r, context, GetContext_552_result, ctx
	return nil
}
func (ϟa *GlDepthMask) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_553_major := uint32(2)    // u32
	minRequiredVersion_553_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_554_result := context             // Contextʳ
	ctx := GetContext_554_result                 // Contextʳ
	ctx.Rasterizing.DepthMask = ϟa.Enabled
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_553_major, minRequiredVersion_553_minor, context, GetContext_554_result, ctx
	return nil
}
func (ϟa *GlFramebufferParameteri) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_555_major := uint32(3) // u32
	minRequiredVersion_555_minor := uint32(1) // u32
	switch ϟa.Target {
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Pname {
	case GLenum_GL_FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS, GLenum_GL_FRAMEBUFFER_DEFAULT_HEIGHT, GLenum_GL_FRAMEBUFFER_DEFAULT_SAMPLES, GLenum_GL_FRAMEBUFFER_DEFAULT_WIDTH:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_555_major, minRequiredVersion_555_minor
	return nil
}
func (ϟa *GlFramebufferRenderbuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_558_major := uint32(2) // u32
	minRequiredVersion_558_minor := uint32(0) // u32
	switch ϟa.FramebufferTarget {
	case GLenum_GL_FRAMEBUFFER:
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
		minRequiredVersion_559_major := uint32(3) // u32
		minRequiredVersion_559_minor := uint32(0) // u32
		_, _ = minRequiredVersion_559_major, minRequiredVersion_559_minor
	default:
		v := ϟa.FramebufferTarget
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.FramebufferAttachment {
	case GLenum_GL_COLOR_ATTACHMENT0, GLenum_GL_DEPTH_ATTACHMENT, GLenum_GL_STENCIL_ATTACHMENT:
	case GLenum_GL_COLOR_ATTACHMENT1, GLenum_GL_COLOR_ATTACHMENT10, GLenum_GL_COLOR_ATTACHMENT11, GLenum_GL_COLOR_ATTACHMENT12, GLenum_GL_COLOR_ATTACHMENT13, GLenum_GL_COLOR_ATTACHMENT14, GLenum_GL_COLOR_ATTACHMENT15, GLenum_GL_COLOR_ATTACHMENT2, GLenum_GL_COLOR_ATTACHMENT3, GLenum_GL_COLOR_ATTACHMENT4, GLenum_GL_COLOR_ATTACHMENT5, GLenum_GL_COLOR_ATTACHMENT6, GLenum_GL_COLOR_ATTACHMENT7, GLenum_GL_COLOR_ATTACHMENT8, GLenum_GL_COLOR_ATTACHMENT9, GLenum_GL_DEPTH_STENCIL_ATTACHMENT:
		minRequiredVersion_561_major := uint32(3) // u32
		minRequiredVersion_561_minor := uint32(0) // u32
		_, _ = minRequiredVersion_561_major, minRequiredVersion_561_minor
	default:
		v := ϟa.FramebufferAttachment
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.RenderbufferTarget {
	case GLenum_GL_RENDERBUFFER:
	default:
		v := ϟa.RenderbufferTarget
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_564_result := context             // Contextʳ
	ctx := GetContext_564_result                 // Contextʳ
	target := func() (result GLenum) {
		switch ϟa.FramebufferTarget {
		case GLenum_GL_FRAMEBUFFER:
			return GLenum_GL_DRAW_FRAMEBUFFER
		case GLenum_GL_DRAW_FRAMEBUFFER:
			return GLenum_GL_DRAW_FRAMEBUFFER
		case GLenum_GL_READ_FRAMEBUFFER:
			return GLenum_GL_READ_FRAMEBUFFER
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.FramebufferTarget, ϟa))
			return result
		}
	}() // GLenum
	framebufferId := ctx.BoundFramebuffers.Get(target)                  // FramebufferId
	framebuffer := ctx.Instances.Framebuffers.Get(framebufferId)        // Framebufferʳ
	attachment := framebuffer.Attachments.Get(ϟa.FramebufferAttachment) // FramebufferAttachmentInfo
	if (ϟa.Renderbuffer) == (RenderbufferId(uint32(0))) {
		attachment.Type = GLenum_GL_NONE
	} else {
		attachment.Type = GLenum_GL_RENDERBUFFER
	}
	attachment.Object = uint32(ϟa.Renderbuffer)
	attachment.TextureLevel = GLint(int32(0))
	attachment.CubeMapFace = GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X
	framebuffer.Attachments[ϟa.FramebufferAttachment] = attachment
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _, _ = minRequiredVersion_558_major, minRequiredVersion_558_minor, context, GetContext_564_result, ctx, target, framebufferId, framebuffer, attachment
	return nil
}
func (ϟa *GlFramebufferTexture2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_565_major := uint32(2) // u32
	minRequiredVersion_565_minor := uint32(0) // u32
	switch ϟa.FramebufferTarget {
	case GLenum_GL_FRAMEBUFFER:
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
		minRequiredVersion_566_major := uint32(3) // u32
		minRequiredVersion_566_minor := uint32(0) // u32
		_, _ = minRequiredVersion_566_major, minRequiredVersion_566_minor
	default:
		v := ϟa.FramebufferTarget
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.FramebufferAttachment {
	case GLenum_GL_COLOR_ATTACHMENT0, GLenum_GL_DEPTH_ATTACHMENT, GLenum_GL_STENCIL_ATTACHMENT:
	case GLenum_GL_COLOR_ATTACHMENT1, GLenum_GL_COLOR_ATTACHMENT10, GLenum_GL_COLOR_ATTACHMENT11, GLenum_GL_COLOR_ATTACHMENT12, GLenum_GL_COLOR_ATTACHMENT13, GLenum_GL_COLOR_ATTACHMENT14, GLenum_GL_COLOR_ATTACHMENT15, GLenum_GL_COLOR_ATTACHMENT2, GLenum_GL_COLOR_ATTACHMENT3, GLenum_GL_COLOR_ATTACHMENT4, GLenum_GL_COLOR_ATTACHMENT5, GLenum_GL_COLOR_ATTACHMENT6, GLenum_GL_COLOR_ATTACHMENT7, GLenum_GL_COLOR_ATTACHMENT8, GLenum_GL_COLOR_ATTACHMENT9, GLenum_GL_DEPTH_STENCIL_ATTACHMENT:
		minRequiredVersion_568_major := uint32(3) // u32
		minRequiredVersion_568_minor := uint32(0) // u32
		_, _ = minRequiredVersion_568_major, minRequiredVersion_568_minor
	default:
		v := ϟa.FramebufferAttachment
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.TextureTarget {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_570_major := uint32(3) // u32
		minRequiredVersion_570_minor := uint32(1) // u32
		_, _ = minRequiredVersion_570_major, minRequiredVersion_570_minor
	default:
		v := ϟa.TextureTarget
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_572_result := context             // Contextʳ
	ctx := GetContext_572_result                 // Contextʳ
	target := func() (result GLenum) {
		switch ϟa.FramebufferTarget {
		case GLenum_GL_FRAMEBUFFER:
			return GLenum_GL_DRAW_FRAMEBUFFER
		case GLenum_GL_DRAW_FRAMEBUFFER:
			return GLenum_GL_DRAW_FRAMEBUFFER
		case GLenum_GL_READ_FRAMEBUFFER:
			return GLenum_GL_READ_FRAMEBUFFER
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.FramebufferTarget, ϟa))
			return result
		}
	}() // GLenum
	framebufferId := ctx.BoundFramebuffers.Get(target)                  // FramebufferId
	framebuffer := ctx.Instances.Framebuffers.Get(framebufferId)        // Framebufferʳ
	attachment := framebuffer.Attachments.Get(ϟa.FramebufferAttachment) // FramebufferAttachmentInfo
	if (ϟa.Texture) == (TextureId(uint32(0))) {
		attachment.Type = GLenum_GL_NONE
		attachment.Object = uint32(0)
		attachment.TextureLevel = GLint(int32(0))
		attachment.CubeMapFace = GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X
	} else {
		attachment.Type = GLenum_GL_TEXTURE
		attachment.Object = uint32(ϟa.Texture)
		attachment.TextureLevel = ϟa.Level
		attachment.CubeMapFace = func() (result GLenum) {
			switch ϟa.TextureTarget {
			case GLenum_GL_TEXTURE_2D:
				return GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X
			case GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X:
				return GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X
			case GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y:
				return GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y
			case GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
				return GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z
			case GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X:
				return GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X
			case GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y:
				return GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y
			case GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z:
				return GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z
			default:
				// TODO: better unmatched handling
				panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.TextureTarget, ϟa))
				return result
			}
		}()
	}
	framebuffer.Attachments[ϟa.FramebufferAttachment] = attachment
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _, _ = minRequiredVersion_565_major, minRequiredVersion_565_minor, context, GetContext_572_result, ctx, target, framebufferId, framebuffer, attachment
	return nil
}
func (ϟa *GlFramebufferTextureLayer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_573_major := uint32(3) // u32
	minRequiredVersion_573_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Attachment {
	case GLenum_GL_COLOR_ATTACHMENT0, GLenum_GL_COLOR_ATTACHMENT1, GLenum_GL_COLOR_ATTACHMENT10, GLenum_GL_COLOR_ATTACHMENT11, GLenum_GL_COLOR_ATTACHMENT12, GLenum_GL_COLOR_ATTACHMENT13, GLenum_GL_COLOR_ATTACHMENT14, GLenum_GL_COLOR_ATTACHMENT15, GLenum_GL_COLOR_ATTACHMENT2, GLenum_GL_COLOR_ATTACHMENT3, GLenum_GL_COLOR_ATTACHMENT4, GLenum_GL_COLOR_ATTACHMENT5, GLenum_GL_COLOR_ATTACHMENT6, GLenum_GL_COLOR_ATTACHMENT7, GLenum_GL_COLOR_ATTACHMENT8, GLenum_GL_COLOR_ATTACHMENT9, GLenum_GL_DEPTH_ATTACHMENT, GLenum_GL_DEPTH_STENCIL_ATTACHMENT, GLenum_GL_STENCIL_ATTACHMENT:
	default:
		v := ϟa.Attachment
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_573_major, minRequiredVersion_573_minor
	return nil
}
func (ϟa *GlGenFramebuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_576_major := uint32(2)                                   // u32
	minRequiredVersion_576_minor := uint32(0)                                   // u32
	f := ϟa.Framebuffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // FramebufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                // Contextʳ
	GetContext_577_result := context                                            // Contextʳ
	ctx := GetContext_577_result                                                // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := FramebufferId(ϟa.Framebuffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil)) // FramebufferId
		ctx.Instances.Framebuffers[id] = func() *Framebuffer {
			s := &Framebuffer{}
			s.Init()
			return s
		}()
		f.Index(uint64(i), ϟs).Write(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = minRequiredVersion_576_major, minRequiredVersion_576_minor, f, context, GetContext_577_result, ctx
	return nil
}
func (ϟa *GlGenRenderbuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_578_major := uint32(2)                                    // u32
	minRequiredVersion_578_minor := uint32(0)                                    // u32
	r := ϟa.Renderbuffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // RenderbufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                 // Contextʳ
	GetContext_579_result := context                                             // Contextʳ
	ctx := GetContext_579_result                                                 // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := RenderbufferId(ϟa.Renderbuffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil)) // RenderbufferId
		ctx.Instances.Renderbuffers[id] = func() *Renderbuffer {
			s := &Renderbuffer{}
			s.Init()
			return s
		}()
		r.Index(uint64(i), ϟs).Write(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = minRequiredVersion_578_major, minRequiredVersion_578_minor, r, context, GetContext_579_result, ctx
	return nil
}
func (ϟa *GlGetFramebufferAttachmentParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_580_major := uint32(2) // u32
	minRequiredVersion_580_minor := uint32(0) // u32
	switch ϟa.FramebufferTarget {
	case GLenum_GL_FRAMEBUFFER:
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
		minRequiredVersion_581_major := uint32(3) // u32
		minRequiredVersion_581_minor := uint32(0) // u32
		_, _ = minRequiredVersion_581_major, minRequiredVersion_581_minor
	default:
		v := ϟa.FramebufferTarget
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Attachment {
	case GLenum_GL_COLOR_ATTACHMENT0, GLenum_GL_DEPTH_ATTACHMENT, GLenum_GL_STENCIL_ATTACHMENT:
	case GLenum_GL_BACK, GLenum_GL_COLOR_ATTACHMENT1, GLenum_GL_COLOR_ATTACHMENT10, GLenum_GL_COLOR_ATTACHMENT11, GLenum_GL_COLOR_ATTACHMENT12, GLenum_GL_COLOR_ATTACHMENT13, GLenum_GL_COLOR_ATTACHMENT14, GLenum_GL_COLOR_ATTACHMENT15, GLenum_GL_COLOR_ATTACHMENT2, GLenum_GL_COLOR_ATTACHMENT3, GLenum_GL_COLOR_ATTACHMENT4, GLenum_GL_COLOR_ATTACHMENT5, GLenum_GL_COLOR_ATTACHMENT6, GLenum_GL_COLOR_ATTACHMENT7, GLenum_GL_COLOR_ATTACHMENT8, GLenum_GL_COLOR_ATTACHMENT9, GLenum_GL_DEPTH, GLenum_GL_DEPTH_STENCIL_ATTACHMENT, GLenum_GL_STENCIL:
		minRequiredVersion_583_major := uint32(3) // u32
		minRequiredVersion_583_minor := uint32(0) // u32
		_, _ = minRequiredVersion_583_major, minRequiredVersion_583_minor
	default:
		v := ϟa.Attachment
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Parameter {
	case GLenum_GL_FRAMEBUFFER_ATTACHMENT_OBJECT_NAME, GLenum_GL_FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL:
	case GLenum_GL_FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_BLUE_SIZE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING, GLenum_GL_FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_GREEN_SIZE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_RED_SIZE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER:
		minRequiredVersion_585_major := uint32(3) // u32
		minRequiredVersion_585_minor := uint32(0) // u32
		_, _ = minRequiredVersion_585_major, minRequiredVersion_585_minor
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_587_result := context             // Contextʳ
	ctx := GetContext_587_result                 // Contextʳ
	target := func() (result GLenum) {
		switch ϟa.FramebufferTarget {
		case GLenum_GL_FRAMEBUFFER:
			return GLenum_GL_DRAW_FRAMEBUFFER
		case GLenum_GL_DRAW_FRAMEBUFFER:
			return GLenum_GL_DRAW_FRAMEBUFFER
		case GLenum_GL_READ_FRAMEBUFFER:
			return GLenum_GL_READ_FRAMEBUFFER
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.FramebufferTarget, ϟa))
			return result
		}
	}() // GLenum
	framebufferId := ctx.BoundFramebuffers.Get(target)           // FramebufferId
	framebuffer := ctx.Instances.Framebuffers.Get(framebufferId) // Framebufferʳ
	a := framebuffer.Attachments.Get(ϟa.Attachment)              // FramebufferAttachmentInfo
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(func() (result GLint) {
		switch ϟa.Parameter {
		case GLenum_GL_FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE:
			return GLint(a.Type)
		case GLenum_GL_FRAMEBUFFER_ATTACHMENT_OBJECT_NAME:
			return GLint(a.Object)
		case GLenum_GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL:
			return a.TextureLevel
		case GLenum_GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE:
			return GLint(a.CubeMapFace)
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.Parameter, ϟa))
			return result
		}
	}(), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _, _, _, _ = minRequiredVersion_580_major, minRequiredVersion_580_minor, context, GetContext_587_result, ctx, target, framebufferId, framebuffer, a
	return nil
}
func (ϟa *GlGetFramebufferParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_588_major := uint32(3) // u32
	minRequiredVersion_588_minor := uint32(1) // u32
	switch ϟa.Target {
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Pname {
	case GLenum_GL_FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS, GLenum_GL_FRAMEBUFFER_DEFAULT_HEIGHT, GLenum_GL_FRAMEBUFFER_DEFAULT_SAMPLES, GLenum_GL_FRAMEBUFFER_DEFAULT_WIDTH:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_588_major, minRequiredVersion_588_minor
	return nil
}
func (ϟa *GlGetRenderbufferParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_591_major := uint32(2) // u32
	minRequiredVersion_591_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_RENDERBUFFER:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Parameter {
	case GLenum_GL_RENDERBUFFER_ALPHA_SIZE, GLenum_GL_RENDERBUFFER_BLUE_SIZE, GLenum_GL_RENDERBUFFER_DEPTH_SIZE, GLenum_GL_RENDERBUFFER_GREEN_SIZE, GLenum_GL_RENDERBUFFER_HEIGHT, GLenum_GL_RENDERBUFFER_INTERNAL_FORMAT, GLenum_GL_RENDERBUFFER_RED_SIZE, GLenum_GL_RENDERBUFFER_STENCIL_SIZE, GLenum_GL_RENDERBUFFER_WIDTH:
	case GLenum_GL_RENDERBUFFER_SAMPLES:
		minRequiredVersion_593_major := uint32(3) // u32
		minRequiredVersion_593_minor := uint32(0) // u32
		_, _ = minRequiredVersion_593_major, minRequiredVersion_593_minor
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_595_result := context             // Contextʳ
	ctx := GetContext_595_result                 // Contextʳ
	id := ctx.BoundRenderbuffers.Get(ϟa.Target)  // RenderbufferId
	rb := ctx.Instances.Renderbuffers.Get(id)    // Renderbufferʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Values.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(func() (result GLint) {
		switch ϟa.Parameter {
		case GLenum_GL_RENDERBUFFER_WIDTH:
			return GLint(rb.Width)
		case GLenum_GL_RENDERBUFFER_HEIGHT:
			return GLint(rb.Height)
		case GLenum_GL_RENDERBUFFER_INTERNAL_FORMAT:
			return GLint(rb.Format)
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.Parameter, ϟa))
			return result
		}
	}(), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _, _ = minRequiredVersion_591_major, minRequiredVersion_591_minor, context, GetContext_595_result, ctx, id, rb
	return nil
}
func (ϟa *GlInvalidateFramebuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_596_major := uint32(3) // u32
	minRequiredVersion_596_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_FRAMEBUFFER:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_596_major, minRequiredVersion_596_minor
	return nil
}
func (ϟa *GlInvalidateSubFramebuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_598_major := uint32(3) // u32
	minRequiredVersion_598_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_FRAMEBUFFER:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_598_major, minRequiredVersion_598_minor
	return nil
}
func (ϟa *GlIsFramebuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_600_major := uint32(2)    // u32
	minRequiredVersion_600_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_601_result := context             // Contextʳ
	ctx := GetContext_601_result                 // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.Framebuffers.Contains(ϟa.Framebuffer) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = minRequiredVersion_600_major, minRequiredVersion_600_minor, context, GetContext_601_result, ctx
	return nil
}
func (ϟa *GlIsRenderbuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_602_major := uint32(2)    // u32
	minRequiredVersion_602_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_603_result := context             // Contextʳ
	ctx := GetContext_603_result                 // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.Renderbuffers.Contains(ϟa.Renderbuffer) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = minRequiredVersion_602_major, minRequiredVersion_602_minor, context, GetContext_603_result, ctx
	return nil
}
func (ϟa *GlReadBuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_604_major := uint32(3) // u32
	minRequiredVersion_604_minor := uint32(0) // u32
	switch ϟa.Src {
	case GLenum_GL_BACK, GLenum_GL_COLOR_ATTACHMENT0, GLenum_GL_COLOR_ATTACHMENT1, GLenum_GL_COLOR_ATTACHMENT10, GLenum_GL_COLOR_ATTACHMENT11, GLenum_GL_COLOR_ATTACHMENT12, GLenum_GL_COLOR_ATTACHMENT13, GLenum_GL_COLOR_ATTACHMENT14, GLenum_GL_COLOR_ATTACHMENT15, GLenum_GL_COLOR_ATTACHMENT2, GLenum_GL_COLOR_ATTACHMENT3, GLenum_GL_COLOR_ATTACHMENT4, GLenum_GL_COLOR_ATTACHMENT5, GLenum_GL_COLOR_ATTACHMENT6, GLenum_GL_COLOR_ATTACHMENT7, GLenum_GL_COLOR_ATTACHMENT8, GLenum_GL_COLOR_ATTACHMENT9, GLenum_GL_NONE:
	default:
		v := ϟa.Src
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_604_major, minRequiredVersion_604_minor
	return nil
}
func (ϟa *GlReadPixels) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_606_major := uint32(2) // u32
	minRequiredVersion_606_minor := uint32(0) // u32
	switch ϟa.Format {
	case GLenum_GL_ALPHA, GLenum_GL_RGB, GLenum_GL_RGBA:
	case GLenum_GL_LUMINANCE, GLenum_GL_LUMINANCE_ALPHA, GLenum_GL_RED, GLenum_GL_RED_INTEGER, GLenum_GL_RG, GLenum_GL_RGBA_INTEGER, GLenum_GL_RGB_INTEGER, GLenum_GL_RG_INTEGER:
		minRequiredVersion_607_major := uint32(3) // u32
		minRequiredVersion_607_minor := uint32(0) // u32
		_, _ = minRequiredVersion_607_major, minRequiredVersion_607_minor
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Type {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_SHORT_4_4_4_4, GLenum_GL_UNSIGNED_SHORT_5_5_5_1, GLenum_GL_UNSIGNED_SHORT_5_6_5:
	case GLenum_GL_BYTE, GLenum_GL_FLOAT, GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_10F_11F_11F_REV, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_INT_5_9_9_9_REV:
		minRequiredVersion_609_major := uint32(3) // u32
		minRequiredVersion_609_minor := uint32(0) // u32
		_, _ = minRequiredVersion_609_major, minRequiredVersion_609_minor
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Data.Slice(uint64(uint32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.imageSize(uint32(ϟa.Width), uint32(ϟa.Height), ϟa.Format, ϟa.Type)), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_606_major, minRequiredVersion_606_minor
	return nil
}
func (ϟa *GlRenderbufferStorage) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_611_major := uint32(2) // u32
	minRequiredVersion_611_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_RENDERBUFFER:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Format {
	case GLenum_GL_DEPTH_COMPONENT16, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGBA4, GLenum_GL_STENCIL_INDEX8:
	case GLenum_GL_DEPTH24_STENCIL8, GLenum_GL_DEPTH32F_STENCIL8, GLenum_GL_DEPTH_COMPONENT24, GLenum_GL_DEPTH_COMPONENT32F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB8, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_SRGB8_ALPHA8:
		minRequiredVersion_613_major := uint32(3) // u32
		minRequiredVersion_613_minor := uint32(0) // u32
		_, _ = minRequiredVersion_613_major, minRequiredVersion_613_minor
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_615_result := context             // Contextʳ
	ctx := GetContext_615_result                 // Contextʳ
	id := ctx.BoundRenderbuffers.Get(ϟa.Target)  // RenderbufferId
	rb := ctx.Instances.Renderbuffers.Get(id)    // Renderbufferʳ
	rb.Format = ϟa.Format
	rb.Width = ϟa.Width
	rb.Height = ϟa.Height
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = minRequiredVersion_611_major, minRequiredVersion_611_minor, context, GetContext_615_result, ctx, id, rb
	return nil
}
func (ϟa *GlRenderbufferStorageMultisample) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_616_major := uint32(3) // u32
	minRequiredVersion_616_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_RENDERBUFFER:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Format {
	case GLenum_GL_DEPTH24_STENCIL8, GLenum_GL_DEPTH32F_STENCIL8, GLenum_GL_DEPTH_COMPONENT16, GLenum_GL_DEPTH_COMPONENT24, GLenum_GL_DEPTH_COMPONENT32F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGB8, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA4, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_SRGB8_ALPHA8, GLenum_GL_STENCIL_INDEX8:
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_616_major, minRequiredVersion_616_minor
	return nil
}
func (ϟa *GlStencilMask) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_619_major := uint32(2)    // u32
	minRequiredVersion_619_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_620_result := context             // Contextʳ
	ctx := GetContext_620_result                 // Contextʳ
	ctx.Rasterizing.StencilMask[GLenum_GL_FRONT] = ϟa.Mask
	ctx.Rasterizing.StencilMask[GLenum_GL_BACK] = ϟa.Mask
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_619_major, minRequiredVersion_619_minor, context, GetContext_620_result, ctx
	return nil
}
func (ϟa *GlStencilMaskSeparate) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_621_major := uint32(2) // u32
	minRequiredVersion_621_minor := uint32(0) // u32
	switch ϟa.Face {
	case GLenum_GL_BACK, GLenum_GL_FRONT, GLenum_GL_FRONT_AND_BACK:
	default:
		v := ϟa.Face
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_623_result := context             // Contextʳ
	ctx := GetContext_623_result                 // Contextʳ
	switch ϟa.Face {
	case GLenum_GL_FRONT:
		ctx.Rasterizing.StencilMask[GLenum_GL_FRONT] = ϟa.Mask
	case GLenum_GL_BACK:
		ctx.Rasterizing.StencilMask[GLenum_GL_BACK] = ϟa.Mask
	case GLenum_GL_FRONT_AND_BACK:
		ctx.Rasterizing.StencilMask[GLenum_GL_FRONT] = ϟa.Mask
		ctx.Rasterizing.StencilMask[GLenum_GL_BACK] = ϟa.Mask
	default:
		v := ϟa.Face
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_621_major, minRequiredVersion_621_minor, context, GetContext_623_result, ctx
	return nil
}
func (ϟa *GlDisable) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_624_major := uint32(2) // u32
	minRequiredVersion_624_minor := uint32(0) // u32
	switch ϟa.Capability {
	case GLenum_GL_BLEND, GLenum_GL_CULL_FACE, GLenum_GL_DEPTH_TEST, GLenum_GL_DITHER, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SCISSOR_TEST, GLenum_GL_STENCIL_TEST:
	case GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_RASTERIZER_DISCARD:
		minRequiredVersion_625_major := uint32(3) // u32
		minRequiredVersion_625_minor := uint32(0) // u32
		_, _ = minRequiredVersion_625_major, minRequiredVersion_625_minor
	case GLenum_GL_SAMPLE_MASK:
		minRequiredVersion_626_major := uint32(3) // u32
		minRequiredVersion_626_minor := uint32(1) // u32
		_, _ = minRequiredVersion_626_major, minRequiredVersion_626_minor
	default:
		v := ϟa.Capability
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_628_result := context             // Contextʳ
	ctx := GetContext_628_result                 // Contextʳ
	ctx.Capabilities[ϟa.Capability] = false
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_624_major, minRequiredVersion_624_minor, context, GetContext_628_result, ctx
	return nil
}
func (ϟa *GlEnable) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_629_major := uint32(2) // u32
	minRequiredVersion_629_minor := uint32(0) // u32
	switch ϟa.Capability {
	case GLenum_GL_BLEND, GLenum_GL_CULL_FACE, GLenum_GL_DEPTH_TEST, GLenum_GL_DITHER, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SCISSOR_TEST, GLenum_GL_STENCIL_TEST:
	case GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_RASTERIZER_DISCARD:
		minRequiredVersion_630_major := uint32(3) // u32
		minRequiredVersion_630_minor := uint32(0) // u32
		_, _ = minRequiredVersion_630_major, minRequiredVersion_630_minor
	case GLenum_GL_SAMPLE_MASK:
		minRequiredVersion_631_major := uint32(3) // u32
		minRequiredVersion_631_minor := uint32(1) // u32
		_, _ = minRequiredVersion_631_major, minRequiredVersion_631_minor
	default:
		v := ϟa.Capability
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_633_result := context             // Contextʳ
	ctx := GetContext_633_result                 // Contextʳ
	ctx.Capabilities[ϟa.Capability] = true
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_629_major, minRequiredVersion_629_minor, context, GetContext_633_result, ctx
	return nil
}
func (ϟa *GlFinish) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_634_major := uint32(2) // u32
	minRequiredVersion_634_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_634_major, minRequiredVersion_634_minor
	return nil
}
func (ϟa *GlFlush) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_635_major := uint32(2) // u32
	minRequiredVersion_635_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_635_major, minRequiredVersion_635_minor
	return nil
}
func (ϟa *GlFlushMappedBufferRange) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_636_major := uint32(3) // u32
	minRequiredVersion_636_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_636_major, minRequiredVersion_636_minor
	return nil
}
func (ϟa *GlGetError) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_638_major := uint32(2) // u32
	minRequiredVersion_638_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_638_major, minRequiredVersion_638_minor
	return nil
}
func (ϟa *GlHint) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_639_major := uint32(2) // u32
	minRequiredVersion_639_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_GENERATE_MIPMAP_HINT:
	case GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT:
		minRequiredVersion_640_major := uint32(3) // u32
		minRequiredVersion_640_minor := uint32(0) // u32
		_, _ = minRequiredVersion_640_major, minRequiredVersion_640_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Mode {
	case GLenum_GL_DONT_CARE, GLenum_GL_FASTEST, GLenum_GL_NICEST:
	default:
		v := ϟa.Mode
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_643_result := context             // Contextʳ
	ctx := GetContext_643_result                 // Contextʳ
	ctx.GenerateMipmapHint = ϟa.Mode
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_639_major, minRequiredVersion_639_minor, context, GetContext_643_result, ctx
	return nil
}
func (ϟa *GlActiveShaderProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_644_major := uint32(3) // u32
	minRequiredVersion_644_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_644_major, minRequiredVersion_644_minor
	return nil
}
func (ϟa *GlAttachShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_645_major := uint32(2)    // u32
	minRequiredVersion_645_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_646_result := context             // Contextʳ
	ctx := GetContext_646_result                 // Contextʳ
	p := ctx.Instances.Programs.Get(ϟa.Program)  // Programʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // Shaderʳ
	p.Shaders[s.Type] = ϟa.Shader
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = minRequiredVersion_645_major, minRequiredVersion_645_minor, context, GetContext_646_result, ctx, p, s
	return nil
}
func (ϟa *GlBindAttribLocation) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_647_major := uint32(2)    // u32
	minRequiredVersion_647_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_648_result := context             // Contextʳ
	ctx := GetContext_648_result                 // Contextʳ
	p := ctx.Instances.Programs.Get(ϟa.Program)  // Programʳ
	p.AttributeBindings[ϟa.Name] = ϟa.Location
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_647_major, minRequiredVersion_647_minor, context, GetContext_648_result, ctx, p
	return nil
}
func (ϟa *GlBindProgramPipeline) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_649_major := uint32(3) // u32
	minRequiredVersion_649_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_649_major, minRequiredVersion_649_minor
	return nil
}
func (ϟa *GlCompileShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_650_major := uint32(2) // u32
	minRequiredVersion_650_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_650_major, minRequiredVersion_650_minor
	return nil
}
func (ϟa *GlCreateProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_651_major := uint32(2)    // u32
	minRequiredVersion_651_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_652_result := context             // Contextʳ
	ctx := GetContext_652_result                 // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	id := ProgramId(ϟa.Result) // ProgramId
	ctx.Instances.Programs[id] = func() *Program {
		s := &Program{}
		s.Init()
		return s
	}()
	ϟa.Result = id
	_, _, _, _, _, _ = minRequiredVersion_651_major, minRequiredVersion_651_minor, context, GetContext_652_result, ctx, id
	return nil
}
func (ϟa *GlCreateShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_653_major := uint32(2) // u32
	minRequiredVersion_653_minor := uint32(0) // u32
	switch ϟa.Type {
	case GLenum_GL_FRAGMENT_SHADER, GLenum_GL_VERTEX_SHADER:
	case GLenum_GL_COMPUTE_SHADER:
		minRequiredVersion_654_major := uint32(3) // u32
		minRequiredVersion_654_minor := uint32(1) // u32
		_, _ = minRequiredVersion_654_major, minRequiredVersion_654_minor
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_656_result := context             // Contextʳ
	ctx := GetContext_656_result                 // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	id := ShaderId(ϟa.Result) // ShaderId
	ctx.Instances.Shaders[id] = func() *Shader {
		s := &Shader{}
		s.Init()
		return s
	}()
	s := ctx.Instances.Shaders.Get(id) // Shaderʳ
	s.Type = ϟa.Type
	ϟa.Result = id
	_, _, _, _, _, _, _ = minRequiredVersion_653_major, minRequiredVersion_653_minor, context, GetContext_656_result, ctx, id, s
	return nil
}
func (ϟa *GlCreateShaderProgramv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_657_major := uint32(3) // u32
	minRequiredVersion_657_minor := uint32(1) // u32
	switch ϟa.Type {
	case GLenum_GL_COMPUTE_SHADER, GLenum_GL_FRAGMENT_SHADER, GLenum_GL_VERTEX_SHADER:
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_657_major, minRequiredVersion_657_minor
	return nil
}
func (ϟa *GlDeleteProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_659_major := uint32(2)    // u32
	minRequiredVersion_659_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_660_result := context             // Contextʳ
	ctx := GetContext_660_result                 // Contextʳ
	delete(ctx.Instances.Programs, ϟa.Program)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_659_major, minRequiredVersion_659_minor, context, GetContext_660_result, ctx
	return nil
}
func (ϟa *GlDeleteProgramPipelines) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_661_major := uint32(3) // u32
	minRequiredVersion_661_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_661_major, minRequiredVersion_661_minor
	return nil
}
func (ϟa *GlDeleteShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_662_major := uint32(2)    // u32
	minRequiredVersion_662_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_663_result := context             // Contextʳ
	ctx := GetContext_663_result                 // Contextʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // Shaderʳ
	s.Deletable = true
	delete(ctx.Instances.Shaders, ϟa.Shader)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_662_major, minRequiredVersion_662_minor, context, GetContext_663_result, ctx, s
	return nil
}
func (ϟa *GlDetachShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_664_major := uint32(2)    // u32
	minRequiredVersion_664_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_665_result := context             // Contextʳ
	ctx := GetContext_665_result                 // Contextʳ
	p := ctx.Instances.Programs.Get(ϟa.Program)  // Programʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // Shaderʳ
	delete(p.Shaders, s.Type)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = minRequiredVersion_664_major, minRequiredVersion_664_minor, context, GetContext_665_result, ctx, p, s
	return nil
}
func (ϟa *GlDispatchCompute) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_666_major := uint32(3) // u32
	minRequiredVersion_666_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_666_major, minRequiredVersion_666_minor
	return nil
}
func (ϟa *GlDispatchComputeIndirect) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_667_major := uint32(3) // u32
	minRequiredVersion_667_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_667_major, minRequiredVersion_667_minor
	return nil
}
func (ϟa *GlGenProgramPipelines) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_668_major := uint32(3) // u32
	minRequiredVersion_668_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_668_major, minRequiredVersion_668_minor
	return nil
}
func (ϟa *GlGetActiveAttrib) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_669_major := uint32(2) // u32
	minRequiredVersion_669_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	if (ϟa.BufferBytesWritten) != (GLsizeiᵖ{}) {
		l := GLsizei(ϟa.BufferBytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil)) // GLsizei
		ϟa.BufferBytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(l, ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟa.Name.Slice(uint64(GLsizei(int32(0))), uint64(l), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = l
	} else {
		ϟa.Name.Slice(uint64(0), uint64(256), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	}
	ϟa.VectorCount.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(GLint(ϟa.VectorCount.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil)), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Type.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(GLenum(ϟa.Type.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil)), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_669_major, minRequiredVersion_669_minor
	return nil
}
func (ϟa *GlGetActiveUniform) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_670_major := uint32(2) // u32
	minRequiredVersion_670_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	if (ϟa.BufferBytesWritten) != (GLsizeiᵖ{}) {
		l := GLsizei(ϟa.BufferBytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil)) // GLsizei
		ϟa.BufferBytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(l, ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟa.Name.Slice(uint64(GLsizei(int32(0))), uint64(l), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = l
	} else {
		ϟa.Name.Slice(uint64(0), uint64(256), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	}
	ϟa.VectorCount.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(GLint(ϟa.VectorCount.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil)), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Type.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(GLenum(ϟa.Type.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil)), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_670_major, minRequiredVersion_670_minor
	return nil
}
func (ϟa *GlGetActiveUniformBlockName) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_671_major := uint32(3) // u32
	minRequiredVersion_671_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	l := GLsizei(ϟa.BufferBytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil)) // GLsizei
	ϟa.BufferBytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(l, ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Name.Slice(uint64(GLsizei(int32(0))), uint64(l), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _ = minRequiredVersion_671_major, minRequiredVersion_671_minor, l
	return nil
}
func (ϟa *GlGetActiveUniformBlockiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_672_major := uint32(3) // u32
	minRequiredVersion_672_minor := uint32(0) // u32
	switch ϟa.ParameterName {
	case GLenum_GL_UNIFORM_BLOCK_ACTIVE_UNIFORMS, GLenum_GL_UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES, GLenum_GL_UNIFORM_BLOCK_BINDING, GLenum_GL_UNIFORM_BLOCK_DATA_SIZE, GLenum_GL_UNIFORM_BLOCK_NAME_LENGTH, GLenum_GL_UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER, GLenum_GL_UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER:
	default:
		v := ϟa.ParameterName
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Parameters.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Parameters.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_672_major, minRequiredVersion_672_minor
	return nil
}
func (ϟa *GlGetActiveUniformsiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_674_major := uint32(3) // u32
	minRequiredVersion_674_minor := uint32(0) // u32
	switch ϟa.ParameterName {
	case GLenum_GL_UNIFORM_ARRAY_STRIDE, GLenum_GL_UNIFORM_BLOCK_INDEX, GLenum_GL_UNIFORM_IS_ROW_MAJOR, GLenum_GL_UNIFORM_MATRIX_STRIDE, GLenum_GL_UNIFORM_NAME_LENGTH, GLenum_GL_UNIFORM_OFFSET, GLenum_GL_UNIFORM_SIZE, GLenum_GL_UNIFORM_TYPE:
	default:
		v := ϟa.ParameterName
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.UniformIndices.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.UniformCount), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Parameters.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.UniformCount), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_674_major, minRequiredVersion_674_minor
	return nil
}
func (ϟa *GlGetAttachedShaders) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_676_major := uint32(2)    // u32
	minRequiredVersion_676_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_677_result := context             // Contextʳ
	ctx := GetContext_677_result                 // Contextʳ
	p := ctx.Instances.Programs.Get(ϟa.Program)  // Programʳ
	min_678_a := int32(ϟa.BufferLength)          // s32
	min_678_b := int32(len(p.Shaders))           // s32
	min_678_result := func() (result int32) {
		switch (min_678_a) < (min_678_b) {
		case true:
			return min_678_a
		case false:
			return min_678_b
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (min_678_a) < (min_678_b), ϟa))
			return result
		}
	}() // s32
	l := min_678_result // s32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.ShadersLengthWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(GLsizei(l), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _, _, _, _, _ = minRequiredVersion_676_major, minRequiredVersion_676_minor, context, GetContext_677_result, ctx, p, min_678_a, min_678_b, min_678_result, l
	return nil
}
func (ϟa *GlGetAttribLocation) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_679_major := uint32(2) // u32
	minRequiredVersion_679_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_679_major, minRequiredVersion_679_minor
	return nil
}
func (ϟa *GlGetFragDataLocation) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_680_major := uint32(3) // u32
	minRequiredVersion_680_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_680_major, minRequiredVersion_680_minor
	return nil
}
func (ϟa *GlGetProgramBinary) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_681_major := uint32(3) // u32
	minRequiredVersion_681_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_681_major, minRequiredVersion_681_minor
	return nil
}
func (ϟa *GlGetProgramInfoLog) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_682_major := uint32(2)    // u32
	minRequiredVersion_682_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_683_result := context             // Contextʳ
	ctx := GetContext_683_result                 // Contextʳ
	p := ctx.Instances.Programs.Get(ϟa.Program)  // Programʳ
	min_684_a := int32(ϟa.BufferLength)          // s32
	min_684_b := int32(p.InfoLog.Count)          // s32
	min_684_result := func() (result int32) {
		switch (min_684_a) < (min_684_b) {
		case true:
			return min_684_a
		case false:
			return min_684_b
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (min_684_a) < (min_684_b), ϟa))
			return result
		}
	}() // s32
	l := min_684_result // s32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Info.Slice(uint64(int32(0)), uint64(l), ϟs).Copy(p.InfoLog.Slice(uint64(int32(0)), uint64(l), ϟs), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.StringLengthWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(GLsizei(l), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _, _, _, _, _ = minRequiredVersion_682_major, minRequiredVersion_682_minor, context, GetContext_683_result, ctx, p, min_684_a, min_684_b, min_684_result, l
	return nil
}
func (ϟa *GlGetProgramInterfaceiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_685_major := uint32(3) // u32
	minRequiredVersion_685_minor := uint32(1) // u32
	switch ϟa.ProgramInterface {
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_BUFFER_VARIABLE, GLenum_GL_PROGRAM_INPUT, GLenum_GL_PROGRAM_OUTPUT, GLenum_GL_SHADER_STORAGE_BLOCK, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_VARYING, GLenum_GL_UNIFORM, GLenum_GL_UNIFORM_BLOCK:
	default:
		v := ϟa.ProgramInterface
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Pname {
	case GLenum_GL_ACTIVE_RESOURCES, GLenum_GL_MAX_NAME_LENGTH, GLenum_GL_MAX_NUM_ACTIVE_VARIABLES:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_685_major, minRequiredVersion_685_minor
	return nil
}
func (ϟa *GlGetProgramPipelineInfoLog) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_688_major := uint32(3) // u32
	minRequiredVersion_688_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_688_major, minRequiredVersion_688_minor
	return nil
}
func (ϟa *GlGetProgramPipelineiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_689_major := uint32(3) // u32
	minRequiredVersion_689_minor := uint32(1) // u32
	switch ϟa.Pname {
	case GLenum_GL_ACTIVE_PROGRAM, GLenum_GL_COMPUTE_SHADER, GLenum_GL_FRAGMENT_SHADER, GLenum_GL_INFO_LOG_LENGTH, GLenum_GL_VALIDATE_STATUS, GLenum_GL_VERTEX_SHADER:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_689_major, minRequiredVersion_689_minor
	return nil
}
func (ϟa *GlGetProgramResourceIndex) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_691_major := uint32(3) // u32
	minRequiredVersion_691_minor := uint32(1) // u32
	switch ϟa.ProgramInterface {
	case GLenum_GL_BUFFER_VARIABLE, GLenum_GL_PROGRAM_INPUT, GLenum_GL_PROGRAM_OUTPUT, GLenum_GL_SHADER_STORAGE_BLOCK, GLenum_GL_TRANSFORM_FEEDBACK_VARYING, GLenum_GL_UNIFORM, GLenum_GL_UNIFORM_BLOCK:
	default:
		v := ϟa.ProgramInterface
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_691_major, minRequiredVersion_691_minor
	return nil
}
func (ϟa *GlGetProgramResourceLocation) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_693_major := uint32(3) // u32
	minRequiredVersion_693_minor := uint32(1) // u32
	switch ϟa.ProgramInterface {
	case GLenum_GL_PROGRAM_INPUT, GLenum_GL_PROGRAM_OUTPUT, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM:
	default:
		v := ϟa.ProgramInterface
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_693_major, minRequiredVersion_693_minor
	return nil
}
func (ϟa *GlGetProgramResourceName) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_695_major := uint32(3) // u32
	minRequiredVersion_695_minor := uint32(1) // u32
	switch ϟa.ProgramInterface {
	case GLenum_GL_BUFFER_VARIABLE, GLenum_GL_PROGRAM_INPUT, GLenum_GL_PROGRAM_OUTPUT, GLenum_GL_SHADER_STORAGE_BLOCK, GLenum_GL_TRANSFORM_FEEDBACK_VARYING, GLenum_GL_UNIFORM, GLenum_GL_UNIFORM_BLOCK:
	default:
		v := ϟa.ProgramInterface
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_695_major, minRequiredVersion_695_minor
	return nil
}
func (ϟa *GlGetProgramResourceiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_697_major := uint32(3) // u32
	minRequiredVersion_697_minor := uint32(1) // u32
	switch ϟa.ProgramInterface {
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_BUFFER_VARIABLE, GLenum_GL_PROGRAM_INPUT, GLenum_GL_PROGRAM_OUTPUT, GLenum_GL_SHADER_STORAGE_BLOCK, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_VARYING, GLenum_GL_UNIFORM, GLenum_GL_UNIFORM_BLOCK:
	default:
		v := ϟa.ProgramInterface
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_697_major, minRequiredVersion_697_minor
	return nil
}
func (ϟa *GlGetProgramiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_699_major := uint32(2) // u32
	minRequiredVersion_699_minor := uint32(0) // u32
	switch ϟa.Parameter {
	case GLenum_GL_ACTIVE_ATTRIBUTES, GLenum_GL_ACTIVE_ATTRIBUTE_MAX_LENGTH, GLenum_GL_ACTIVE_UNIFORMS, GLenum_GL_ACTIVE_UNIFORM_MAX_LENGTH, GLenum_GL_ATTACHED_SHADERS, GLenum_GL_DELETE_STATUS, GLenum_GL_INFO_LOG_LENGTH, GLenum_GL_LINK_STATUS, GLenum_GL_VALIDATE_STATUS:
	case GLenum_GL_ACTIVE_UNIFORM_BLOCKS, GLenum_GL_ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH, GLenum_GL_PROGRAM_BINARY_RETRIEVABLE_HINT, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_MODE, GLenum_GL_TRANSFORM_FEEDBACK_VARYINGS, GLenum_GL_TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH:
		minRequiredVersion_700_major := uint32(3) // u32
		minRequiredVersion_700_minor := uint32(0) // u32
		_, _ = minRequiredVersion_700_major, minRequiredVersion_700_minor
	case GLenum_GL_ACTIVE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_PROGRAM_SEPARABLE:
		minRequiredVersion_701_major := uint32(3) // u32
		minRequiredVersion_701_minor := uint32(1) // u32
		_, _ = minRequiredVersion_701_major, minRequiredVersion_701_minor
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_699_major, minRequiredVersion_699_minor
	return nil
}
func (ϟa *GlGetShaderInfoLog) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_703_major := uint32(2)    // u32
	minRequiredVersion_703_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_704_result := context             // Contextʳ
	ctx := GetContext_704_result                 // Contextʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // Shaderʳ
	min_705_a := int32(ϟa.BufferLength)          // s32
	min_705_b := int32(s.InfoLog.Count)          // s32
	min_705_result := func() (result int32) {
		switch (min_705_a) < (min_705_b) {
		case true:
			return min_705_a
		case false:
			return min_705_b
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (min_705_a) < (min_705_b), ϟa))
			return result
		}
	}() // s32
	l := min_705_result // s32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Info.Slice(uint64(int32(0)), uint64(l), ϟs).Copy(s.InfoLog.Slice(uint64(int32(0)), uint64(l), ϟs), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.StringLengthWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(GLsizei(l), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _, _, _, _, _ = minRequiredVersion_703_major, minRequiredVersion_703_minor, context, GetContext_704_result, ctx, s, min_705_a, min_705_b, min_705_result, l
	return nil
}
func (ϟa *GlGetShaderPrecisionFormat) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_706_major := uint32(2) // u32
	minRequiredVersion_706_minor := uint32(0) // u32
	switch ϟa.ShaderType {
	case GLenum_GL_FRAGMENT_SHADER, GLenum_GL_VERTEX_SHADER:
	default:
		v := ϟa.ShaderType
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.PrecisionType {
	case GLenum_GL_HIGH_FLOAT, GLenum_GL_HIGH_INT, GLenum_GL_LOW_FLOAT, GLenum_GL_LOW_INT, GLenum_GL_MEDIUM_FLOAT, GLenum_GL_MEDIUM_INT:
	default:
		v := ϟa.PrecisionType
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Range.Slice(uint64(0), uint64(2), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Precision.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Precision.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_706_major, minRequiredVersion_706_minor
	return nil
}
func (ϟa *GlGetShaderSource) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_709_major := uint32(2)    // u32
	minRequiredVersion_709_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_710_result := context             // Contextʳ
	ctx := GetContext_710_result                 // Contextʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // Shaderʳ
	min_711_a := int32(ϟa.BufferLength)          // s32
	min_711_b := int32(len(s.Source))            // s32
	min_711_result := func() (result int32) {
		switch (min_711_a) < (min_711_b) {
		case true:
			return min_711_a
		case false:
			return min_711_b
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (min_711_a) < (min_711_b), ϟa))
			return result
		}
	}() // s32
	l := min_711_result // s32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	Charᵖ(ϟa.Source).Slice(uint64(int32(0)), uint64(l), ϟs).Copy(MakeCharˢFromString(s.Source, ϟs).Slice(uint64(int32(0)), uint64(l), ϟs), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.StringLengthWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(GLsizei(l), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _, _, _, _, _ = minRequiredVersion_709_major, minRequiredVersion_709_minor, context, GetContext_710_result, ctx, s, min_711_a, min_711_b, min_711_result, l
	return nil
}
func (ϟa *GlGetShaderiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_712_major := uint32(2) // u32
	minRequiredVersion_712_minor := uint32(0) // u32
	switch ϟa.Parameter {
	case GLenum_GL_COMPILE_STATUS, GLenum_GL_DELETE_STATUS, GLenum_GL_INFO_LOG_LENGTH, GLenum_GL_SHADER_SOURCE_LENGTH, GLenum_GL_SHADER_TYPE:
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_714_result := context             // Contextʳ
	ctx := GetContext_714_result                 // Contextʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // Shaderʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(func() (result GLint) {
		switch ϟa.Parameter {
		case GLenum_GL_SHADER_TYPE:
			return GLint(s.Type)
		case GLenum_GL_DELETE_STATUS:
			return func() (result GLint) {
				switch s.Deletable {
				case true:
					return GLint(int32(1))
				case false:
					return GLint(int32(0))
				default:
					// TODO: better unmatched handling
					panic(fmt.Errorf("Unmatched switch(%v) in atom %T", s.Deletable, ϟa))
					return result
				}
			}()
		case GLenum_GL_COMPILE_STATUS:
			return func() (result GLint) {
				switch s.Compiled {
				case true:
					return GLint(int32(1))
				case false:
					return GLint(int32(0))
				default:
					// TODO: better unmatched handling
					panic(fmt.Errorf("Unmatched switch(%v) in atom %T", s.Compiled, ϟa))
					return result
				}
			}()
		case GLenum_GL_INFO_LOG_LENGTH:
			return GLint(int32(s.InfoLog.Count))
		case GLenum_GL_SHADER_SOURCE_LENGTH:
			return GLint(int32(len(s.Source)))
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.Parameter, ϟa))
			return result
		}
	}(), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _ = minRequiredVersion_712_major, minRequiredVersion_712_minor, context, GetContext_714_result, ctx, s
	return nil
}
func (ϟa *GlGetUniformBlockIndex) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_715_major := uint32(3) // u32
	minRequiredVersion_715_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_715_major, minRequiredVersion_715_minor
	return nil
}
func (ϟa *GlGetUniformIndices) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_716_major := uint32(3) // u32
	minRequiredVersion_716_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_716_major, minRequiredVersion_716_minor
	return nil
}
func (ϟa *GlGetUniformLocation) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_717_major := uint32(2) // u32
	minRequiredVersion_717_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_717_major, minRequiredVersion_717_minor
	return nil
}
func (ϟa *GlGetUniformfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_718_major := uint32(2) // u32
	minRequiredVersion_718_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_718_major, minRequiredVersion_718_minor
	return nil
}
func (ϟa *GlGetUniformiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_719_major := uint32(2) // u32
	minRequiredVersion_719_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_719_major, minRequiredVersion_719_minor
	return nil
}
func (ϟa *GlGetUniformuiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_720_major := uint32(3) // u32
	minRequiredVersion_720_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_720_major, minRequiredVersion_720_minor
	return nil
}
func (ϟa *GlIsProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_721_major := uint32(2)    // u32
	minRequiredVersion_721_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_722_result := context             // Contextʳ
	ctx := GetContext_722_result                 // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.Programs.Contains(ϟa.Program) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = minRequiredVersion_721_major, minRequiredVersion_721_minor, context, GetContext_722_result, ctx
	return nil
}
func (ϟa *GlIsProgramPipeline) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_723_major := uint32(3) // u32
	minRequiredVersion_723_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_723_major, minRequiredVersion_723_minor
	return nil
}
func (ϟa *GlIsShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_724_major := uint32(2)    // u32
	minRequiredVersion_724_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_725_result := context             // Contextʳ
	ctx := GetContext_725_result                 // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.Shaders.Contains(ϟa.Shader) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = minRequiredVersion_724_major, minRequiredVersion_724_minor, context, GetContext_725_result, ctx
	return nil
}
func (ϟa *GlLinkProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_726_major := uint32(2) // u32
	minRequiredVersion_726_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_726_major, minRequiredVersion_726_minor
	return nil
}
func (ϟa *GlMemoryBarrier) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_727_major := uint32(3)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               // u32
	minRequiredVersion_727_minor := uint32(1)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               // u32
	supportsBits_728_seenBits := ϟa.Barriers                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                // GLbitfield
	supportsBits_728_validBits := (GLbitfield_GL_ALL_BARRIER_BITS) | ((GLbitfield_GL_ATOMIC_COUNTER_BARRIER_BIT) | ((GLbitfield_GL_BUFFER_UPDATE_BARRIER_BIT) | ((GLbitfield_GL_COMMAND_BARRIER_BIT) | ((GLbitfield_GL_ELEMENT_ARRAY_BARRIER_BIT) | ((GLbitfield_GL_FRAMEBUFFER_BARRIER_BIT) | ((GLbitfield_GL_PIXEL_BUFFER_BARRIER_BIT) | ((GLbitfield_GL_SHADER_IMAGE_ACCESS_BARRIER_BIT) | ((GLbitfield_GL_SHADER_STORAGE_BARRIER_BIT) | ((GLbitfield_GL_TEXTURE_FETCH_BARRIER_BIT) | ((GLbitfield_GL_TEXTURE_UPDATE_BARRIER_BIT) | ((GLbitfield_GL_TRANSFORM_FEEDBACK_BARRIER_BIT) | ((GLbitfield_GL_UNIFORM_BARRIER_BIT) | (GLbitfield_GL_VERTEX_ATTRIB_ARRAY_BARRIER_BIT))))))))))))) // GLbitfield
	if (GLbitfield_GL_ALL_BARRIER_BITS)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_ATOMIC_COUNTER_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_BUFFER_UPDATE_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_COMMAND_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_ELEMENT_ARRAY_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_FRAMEBUFFER_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_PIXEL_BUFFER_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_SHADER_IMAGE_ACCESS_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_SHADER_STORAGE_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_TEXTURE_FETCH_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_TEXTURE_UPDATE_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_TRANSFORM_FEEDBACK_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_UNIFORM_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_VERTEX_ATTRIB_ARRAY_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = minRequiredVersion_727_major, minRequiredVersion_727_minor, supportsBits_728_seenBits, supportsBits_728_validBits
	return nil
}
func (ϟa *GlMemoryBarrierByRegion) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_729_major := uint32(3)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               // u32
	minRequiredVersion_729_minor := uint32(1)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               // u32
	supportsBits_730_seenBits := ϟa.Barriers                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                // GLbitfield
	supportsBits_730_validBits := (GLbitfield_GL_ALL_BARRIER_BITS) | ((GLbitfield_GL_ATOMIC_COUNTER_BARRIER_BIT) | ((GLbitfield_GL_BUFFER_UPDATE_BARRIER_BIT) | ((GLbitfield_GL_COMMAND_BARRIER_BIT) | ((GLbitfield_GL_ELEMENT_ARRAY_BARRIER_BIT) | ((GLbitfield_GL_FRAMEBUFFER_BARRIER_BIT) | ((GLbitfield_GL_PIXEL_BUFFER_BARRIER_BIT) | ((GLbitfield_GL_SHADER_IMAGE_ACCESS_BARRIER_BIT) | ((GLbitfield_GL_SHADER_STORAGE_BARRIER_BIT) | ((GLbitfield_GL_TEXTURE_FETCH_BARRIER_BIT) | ((GLbitfield_GL_TEXTURE_UPDATE_BARRIER_BIT) | ((GLbitfield_GL_TRANSFORM_FEEDBACK_BARRIER_BIT) | ((GLbitfield_GL_UNIFORM_BARRIER_BIT) | (GLbitfield_GL_VERTEX_ATTRIB_ARRAY_BARRIER_BIT))))))))))))) // GLbitfield
	if (GLbitfield_GL_ALL_BARRIER_BITS)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_ATOMIC_COUNTER_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_BUFFER_UPDATE_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_COMMAND_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_ELEMENT_ARRAY_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_FRAMEBUFFER_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_PIXEL_BUFFER_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_SHADER_IMAGE_ACCESS_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_SHADER_STORAGE_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_TEXTURE_FETCH_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_TEXTURE_UPDATE_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_TRANSFORM_FEEDBACK_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_UNIFORM_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_VERTEX_ATTRIB_ARRAY_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = minRequiredVersion_729_major, minRequiredVersion_729_minor, supportsBits_730_seenBits, supportsBits_730_validBits
	return nil
}
func (ϟa *GlProgramBinary) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_731_major := uint32(3) // u32
	minRequiredVersion_731_minor := uint32(0) // u32
	switch ϟa.BinaryFormat {
	default:
		v := ϟa.BinaryFormat
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_731_major, minRequiredVersion_731_minor
	return nil
}
func (ϟa *GlProgramParameteri) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_733_major := uint32(3) // u32
	minRequiredVersion_733_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_PROGRAM_BINARY_RETRIEVABLE_HINT:
	case GLenum_GL_PROGRAM_SEPARABLE:
		minRequiredVersion_734_major := uint32(3) // u32
		minRequiredVersion_734_minor := uint32(1) // u32
		_, _ = minRequiredVersion_734_major, minRequiredVersion_734_minor
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_733_major, minRequiredVersion_733_minor
	return nil
}
func (ϟa *GlProgramUniform1f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_736_major := uint32(3) // u32
	minRequiredVersion_736_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_736_major, minRequiredVersion_736_minor
	return nil
}
func (ϟa *GlProgramUniform1fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_737_major := uint32(3) // u32
	minRequiredVersion_737_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_737_major, minRequiredVersion_737_minor
	return nil
}
func (ϟa *GlProgramUniform1i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_738_major := uint32(3) // u32
	minRequiredVersion_738_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_738_major, minRequiredVersion_738_minor
	return nil
}
func (ϟa *GlProgramUniform1iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_739_major := uint32(3) // u32
	minRequiredVersion_739_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_739_major, minRequiredVersion_739_minor
	return nil
}
func (ϟa *GlProgramUniform1ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_740_major := uint32(3) // u32
	minRequiredVersion_740_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_740_major, minRequiredVersion_740_minor
	return nil
}
func (ϟa *GlProgramUniform1uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_741_major := uint32(3) // u32
	minRequiredVersion_741_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_741_major, minRequiredVersion_741_minor
	return nil
}
func (ϟa *GlProgramUniform2f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_742_major := uint32(3) // u32
	minRequiredVersion_742_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_742_major, minRequiredVersion_742_minor
	return nil
}
func (ϟa *GlProgramUniform2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_743_major := uint32(3) // u32
	minRequiredVersion_743_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_743_major, minRequiredVersion_743_minor
	return nil
}
func (ϟa *GlProgramUniform2i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_744_major := uint32(3) // u32
	minRequiredVersion_744_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_744_major, minRequiredVersion_744_minor
	return nil
}
func (ϟa *GlProgramUniform2iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_745_major := uint32(3) // u32
	minRequiredVersion_745_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_745_major, minRequiredVersion_745_minor
	return nil
}
func (ϟa *GlProgramUniform2ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_746_major := uint32(3) // u32
	minRequiredVersion_746_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_746_major, minRequiredVersion_746_minor
	return nil
}
func (ϟa *GlProgramUniform2uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_747_major := uint32(3) // u32
	minRequiredVersion_747_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_747_major, minRequiredVersion_747_minor
	return nil
}
func (ϟa *GlProgramUniform3f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_748_major := uint32(3) // u32
	minRequiredVersion_748_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_748_major, minRequiredVersion_748_minor
	return nil
}
func (ϟa *GlProgramUniform3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_749_major := uint32(3) // u32
	minRequiredVersion_749_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_749_major, minRequiredVersion_749_minor
	return nil
}
func (ϟa *GlProgramUniform3i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_750_major := uint32(3) // u32
	minRequiredVersion_750_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_750_major, minRequiredVersion_750_minor
	return nil
}
func (ϟa *GlProgramUniform3iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_751_major := uint32(3) // u32
	minRequiredVersion_751_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_751_major, minRequiredVersion_751_minor
	return nil
}
func (ϟa *GlProgramUniform3ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_752_major := uint32(3) // u32
	minRequiredVersion_752_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_752_major, minRequiredVersion_752_minor
	return nil
}
func (ϟa *GlProgramUniform3uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_753_major := uint32(3) // u32
	minRequiredVersion_753_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_753_major, minRequiredVersion_753_minor
	return nil
}
func (ϟa *GlProgramUniform4f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_754_major := uint32(3) // u32
	minRequiredVersion_754_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_754_major, minRequiredVersion_754_minor
	return nil
}
func (ϟa *GlProgramUniform4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_755_major := uint32(3) // u32
	minRequiredVersion_755_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_755_major, minRequiredVersion_755_minor
	return nil
}
func (ϟa *GlProgramUniform4i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_756_major := uint32(3) // u32
	minRequiredVersion_756_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_756_major, minRequiredVersion_756_minor
	return nil
}
func (ϟa *GlProgramUniform4iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_757_major := uint32(3) // u32
	minRequiredVersion_757_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_757_major, minRequiredVersion_757_minor
	return nil
}
func (ϟa *GlProgramUniform4ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_758_major := uint32(3) // u32
	minRequiredVersion_758_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_758_major, minRequiredVersion_758_minor
	return nil
}
func (ϟa *GlProgramUniform4uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_759_major := uint32(3) // u32
	minRequiredVersion_759_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_759_major, minRequiredVersion_759_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_760_major := uint32(3) // u32
	minRequiredVersion_760_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_760_major, minRequiredVersion_760_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix2x3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_761_major := uint32(3) // u32
	minRequiredVersion_761_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_761_major, minRequiredVersion_761_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix2x4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_762_major := uint32(3) // u32
	minRequiredVersion_762_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_762_major, minRequiredVersion_762_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_763_major := uint32(3) // u32
	minRequiredVersion_763_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_763_major, minRequiredVersion_763_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix3x2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_764_major := uint32(3) // u32
	minRequiredVersion_764_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_764_major, minRequiredVersion_764_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix3x4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_765_major := uint32(3) // u32
	minRequiredVersion_765_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_765_major, minRequiredVersion_765_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_766_major := uint32(3) // u32
	minRequiredVersion_766_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_766_major, minRequiredVersion_766_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix4x2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_767_major := uint32(3) // u32
	minRequiredVersion_767_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_767_major, minRequiredVersion_767_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix4x3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_768_major := uint32(3) // u32
	minRequiredVersion_768_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_768_major, minRequiredVersion_768_minor
	return nil
}
func (ϟa *GlReleaseShaderCompiler) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_769_major := uint32(2) // u32
	minRequiredVersion_769_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_769_major, minRequiredVersion_769_minor
	return nil
}
func (ϟa *GlShaderBinary) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_770_major := uint32(2) // u32
	minRequiredVersion_770_minor := uint32(0) // u32
	switch ϟa.BinaryFormat {
	default:
		v := ϟa.BinaryFormat
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_770_major, minRequiredVersion_770_minor
	return nil
}
func (ϟa *GlShaderSource) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_772_major := uint32(2)                                   // u32
	minRequiredVersion_772_minor := uint32(0)                                   // u32
	sources := ϟa.Source.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // GLcharᶜᵖˢ
	lengths := ϟa.Length.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // GLintˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                // Contextʳ
	GetContext_773_result := context                                            // Contextʳ
	ctx := GetContext_773_result                                                // Contextʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)                                   // Shaderʳ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		str := func() (result string) {
			switch ((ϟa.Length) == (GLintᶜᵖ{})) || ((lengths.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb)) < (GLint(int32(0)))) {
			case true:
				return strings.TrimRight(string(Charᵖ(sources.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb)).StringSlice(ϟs, ϟd, ϟl).Read(ϟa, ϟs, ϟd, ϟl, ϟb)), "\x00")
			case false:
				return string(Charᵖ(sources.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb)).Slice(uint64(GLint(int32(0))), uint64(lengths.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb)), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
			default:
				// TODO: better unmatched handling
				panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ((ϟa.Length) == (GLintᶜᵖ{})) || ((lengths.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb)) < (GLint(int32(0)))), ϟa))
				return result
			}
		}() // string
		s.Source += str
		_ = str
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_772_major, minRequiredVersion_772_minor, sources, lengths, context, GetContext_773_result, ctx, s
	return nil
}
func (ϟa *GlUniform1f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_774_major := uint32(2)    // u32
	minRequiredVersion_774_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_775_result := context             // Contextʳ
	ctx := GetContext_775_result                 // Contextʳ
	v := MakeGLfloatˢ(uint64(1), ϟs)             // GLfloatˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).Write(ϟa.Value, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_FLOAT
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_774_major, minRequiredVersion_774_minor, context, GetContext_775_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform1fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_776_major := uint32(2)                             // u32
	minRequiredVersion_776_minor := uint32(0)                             // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                          // Contextʳ
	GetContext_777_result := context                                      // Contextʳ
	ctx := GetContext_777_result                                          // Contextʳ
	v := ϟa.Values.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // GLfloatˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)               // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                          // Uniform
	uniform.Type = GLenum_GL_FLOAT
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_776_major, minRequiredVersion_776_minor, context, GetContext_777_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform1i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_778_major := uint32(2)    // u32
	minRequiredVersion_778_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_779_result := context             // Contextʳ
	ctx := GetContext_779_result                 // Contextʳ
	v := MakeGLintˢ(uint64(1), ϟs)               // GLintˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).Write(ϟa.Value, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_INT
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_778_major, minRequiredVersion_778_minor, context, GetContext_779_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform1iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_780_major := uint32(2)                             // u32
	minRequiredVersion_780_minor := uint32(0)                             // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                          // Contextʳ
	GetContext_781_result := context                                      // Contextʳ
	ctx := GetContext_781_result                                          // Contextʳ
	v := ϟa.Values.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // GLintˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)               // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                          // Uniform
	uniform.Type = GLenum_GL_INT
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_780_major, minRequiredVersion_780_minor, context, GetContext_781_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform1ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_782_major := uint32(3) // u32
	minRequiredVersion_782_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_782_major, minRequiredVersion_782_minor
	return nil
}
func (ϟa *GlUniform1uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_783_major := uint32(3) // u32
	minRequiredVersion_783_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_783_major, minRequiredVersion_783_minor
	return nil
}
func (ϟa *GlUniform2f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_784_major := uint32(2)    // u32
	minRequiredVersion_784_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_785_result := context             // Contextʳ
	ctx := GetContext_785_result                 // Contextʳ
	v := MakeVec2fˢ(uint64(1), ϟs)               // Vec2fˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).Write(Vec2f{Elements: [2]GLfloat{ϟa.Value0, ϟa.Value1}}, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_FLOAT_VEC2
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_784_major, minRequiredVersion_784_minor, context, GetContext_785_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_786_major := uint32(2)                                     // u32
	minRequiredVersion_786_minor := uint32(0)                                     // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                  // Contextʳ
	GetContext_787_result := context                                              // Contextʳ
	ctx := GetContext_787_result                                                  // Contextʳ
	v := Vec2fᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Vec2fˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_FLOAT_VEC2
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_786_major, minRequiredVersion_786_minor, context, GetContext_787_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform2i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_788_major := uint32(2)    // u32
	minRequiredVersion_788_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_789_result := context             // Contextʳ
	ctx := GetContext_789_result                 // Contextʳ
	v := MakeVec2iˢ(uint64(1), ϟs)               // Vec2iˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).Write(Vec2i{Elements: [2]GLint{ϟa.Value0, ϟa.Value1}}, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_INT_VEC2
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_788_major, minRequiredVersion_788_minor, context, GetContext_789_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform2iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_790_major := uint32(2)                                     // u32
	minRequiredVersion_790_minor := uint32(0)                                     // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                  // Contextʳ
	GetContext_791_result := context                                              // Contextʳ
	ctx := GetContext_791_result                                                  // Contextʳ
	v := Vec2iᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Vec2iˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_INT_VEC2
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_790_major, minRequiredVersion_790_minor, context, GetContext_791_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform2ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_792_major := uint32(3) // u32
	minRequiredVersion_792_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_792_major, minRequiredVersion_792_minor
	return nil
}
func (ϟa *GlUniform2uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_793_major := uint32(3) // u32
	minRequiredVersion_793_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_793_major, minRequiredVersion_793_minor
	return nil
}
func (ϟa *GlUniform3f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_794_major := uint32(2)    // u32
	minRequiredVersion_794_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_795_result := context             // Contextʳ
	ctx := GetContext_795_result                 // Contextʳ
	v := MakeVec3fˢ(uint64(1), ϟs)               // Vec3fˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).Write(Vec3f{Elements: [3]GLfloat{ϟa.Value0, ϟa.Value1, ϟa.Value2}}, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_FLOAT_VEC3
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_794_major, minRequiredVersion_794_minor, context, GetContext_795_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_796_major := uint32(2)                                     // u32
	minRequiredVersion_796_minor := uint32(0)                                     // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                  // Contextʳ
	GetContext_797_result := context                                              // Contextʳ
	ctx := GetContext_797_result                                                  // Contextʳ
	v := Vec3fᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Vec3fˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_FLOAT_VEC3
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_796_major, minRequiredVersion_796_minor, context, GetContext_797_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform3i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_798_major := uint32(2)    // u32
	minRequiredVersion_798_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_799_result := context             // Contextʳ
	ctx := GetContext_799_result                 // Contextʳ
	v := MakeVec3iˢ(uint64(1), ϟs)               // Vec3iˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).Write(Vec3i{Elements: [3]GLint{ϟa.Value0, ϟa.Value1, ϟa.Value2}}, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_INT_VEC3
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_798_major, minRequiredVersion_798_minor, context, GetContext_799_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform3iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_800_major := uint32(2)                                     // u32
	minRequiredVersion_800_minor := uint32(0)                                     // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                  // Contextʳ
	GetContext_801_result := context                                              // Contextʳ
	ctx := GetContext_801_result                                                  // Contextʳ
	v := Vec3iᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Vec3iˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_INT_VEC3
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_800_major, minRequiredVersion_800_minor, context, GetContext_801_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform3ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_802_major := uint32(3) // u32
	minRequiredVersion_802_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_802_major, minRequiredVersion_802_minor
	return nil
}
func (ϟa *GlUniform3uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_803_major := uint32(3) // u32
	minRequiredVersion_803_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_803_major, minRequiredVersion_803_minor
	return nil
}
func (ϟa *GlUniform4f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_804_major := uint32(2)    // u32
	minRequiredVersion_804_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_805_result := context             // Contextʳ
	ctx := GetContext_805_result                 // Contextʳ
	v := MakeVec4fˢ(uint64(1), ϟs)               // Vec4fˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).Write(Vec4f{Elements: [4]GLfloat{ϟa.Value0, ϟa.Value1, ϟa.Value2, ϟa.Value3}}, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_FLOAT_VEC4
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_804_major, minRequiredVersion_804_minor, context, GetContext_805_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_806_major := uint32(2)                                     // u32
	minRequiredVersion_806_minor := uint32(0)                                     // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                  // Contextʳ
	GetContext_807_result := context                                              // Contextʳ
	ctx := GetContext_807_result                                                  // Contextʳ
	v := Vec4fᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Vec4fˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_FLOAT_VEC4
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_806_major, minRequiredVersion_806_minor, context, GetContext_807_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform4i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_808_major := uint32(2)    // u32
	minRequiredVersion_808_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_809_result := context             // Contextʳ
	ctx := GetContext_809_result                 // Contextʳ
	v := MakeVec4iˢ(uint64(1), ϟs)               // Vec4iˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).Write(Vec4i{Elements: [4]GLint{ϟa.Value0, ϟa.Value1, ϟa.Value2, ϟa.Value3}}, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_INT_VEC4
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_808_major, minRequiredVersion_808_minor, context, GetContext_809_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform4iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_810_major := uint32(2)                                     // u32
	minRequiredVersion_810_minor := uint32(0)                                     // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                  // Contextʳ
	GetContext_811_result := context                                              // Contextʳ
	ctx := GetContext_811_result                                                  // Contextʳ
	v := Vec4iᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Vec4iˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_INT_VEC4
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_810_major, minRequiredVersion_810_minor, context, GetContext_811_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform4ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_812_major := uint32(3) // u32
	minRequiredVersion_812_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_812_major, minRequiredVersion_812_minor
	return nil
}
func (ϟa *GlUniform4uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_813_major := uint32(3) // u32
	minRequiredVersion_813_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_813_major, minRequiredVersion_813_minor
	return nil
}
func (ϟa *GlUniformBlockBinding) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_814_major := uint32(3) // u32
	minRequiredVersion_814_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_814_major, minRequiredVersion_814_minor
	return nil
}
func (ϟa *GlUniformMatrix2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_815_major := uint32(2)                                     // u32
	minRequiredVersion_815_minor := uint32(0)                                     // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                  // Contextʳ
	GetContext_816_result := context                                              // Contextʳ
	ctx := GetContext_816_result                                                  // Contextʳ
	v := Mat2fᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Mat2fˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_FLOAT_MAT2
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_815_major, minRequiredVersion_815_minor, context, GetContext_816_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniformMatrix2x3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_817_major := uint32(3) // u32
	minRequiredVersion_817_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_817_major, minRequiredVersion_817_minor
	return nil
}
func (ϟa *GlUniformMatrix2x4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_818_major := uint32(3) // u32
	minRequiredVersion_818_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_818_major, minRequiredVersion_818_minor
	return nil
}
func (ϟa *GlUniformMatrix3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_819_major := uint32(2)                                     // u32
	minRequiredVersion_819_minor := uint32(0)                                     // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                  // Contextʳ
	GetContext_820_result := context                                              // Contextʳ
	ctx := GetContext_820_result                                                  // Contextʳ
	v := Mat3fᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Mat3fˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_FLOAT_MAT3
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_819_major, minRequiredVersion_819_minor, context, GetContext_820_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniformMatrix3x2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_821_major := uint32(3) // u32
	minRequiredVersion_821_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_821_major, minRequiredVersion_821_minor
	return nil
}
func (ϟa *GlUniformMatrix3x4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_822_major := uint32(3) // u32
	minRequiredVersion_822_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_822_major, minRequiredVersion_822_minor
	return nil
}
func (ϟa *GlUniformMatrix4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_823_major := uint32(2)                                     // u32
	minRequiredVersion_823_minor := uint32(0)                                     // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                  // Contextʳ
	GetContext_824_result := context                                              // Contextʳ
	ctx := GetContext_824_result                                                  // Contextʳ
	v := Mat4fᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Mat4fˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_823_major, minRequiredVersion_823_minor, context, GetContext_824_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniformMatrix4x2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_825_major := uint32(3) // u32
	minRequiredVersion_825_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_825_major, minRequiredVersion_825_minor
	return nil
}
func (ϟa *GlUniformMatrix4x3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_826_major := uint32(3) // u32
	minRequiredVersion_826_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_826_major, minRequiredVersion_826_minor
	return nil
}
func (ϟa *GlUseProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_827_major := uint32(2)    // u32
	minRequiredVersion_827_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_828_result := context             // Contextʳ
	ctx := GetContext_828_result                 // Contextʳ
	ctx.BoundProgram = ϟa.Program
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_827_major, minRequiredVersion_827_minor, context, GetContext_828_result, ctx
	return nil
}
func (ϟa *GlUseProgramStages) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_829_major := uint32(3)                                                                                                                                        // u32
	minRequiredVersion_829_minor := uint32(1)                                                                                                                                        // u32
	supportsBits_830_seenBits := ϟa.Stages                                                                                                                                           // GLbitfield
	supportsBits_830_validBits := (GLbitfield_GL_ALL_SHADER_BITS) | ((GLbitfield_GL_COMPUTE_SHADER_BIT) | ((GLbitfield_GL_FRAGMENT_SHADER_BIT) | (GLbitfield_GL_VERTEX_SHADER_BIT))) // GLbitfield
	if (GLbitfield_GL_ALL_SHADER_BITS)&(ϟa.Stages) != 0 {
	}
	if (GLbitfield_GL_COMPUTE_SHADER_BIT)&(ϟa.Stages) != 0 {
	}
	if (GLbitfield_GL_FRAGMENT_SHADER_BIT)&(ϟa.Stages) != 0 {
	}
	if (GLbitfield_GL_VERTEX_SHADER_BIT)&(ϟa.Stages) != 0 {
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = minRequiredVersion_829_major, minRequiredVersion_829_minor, supportsBits_830_seenBits, supportsBits_830_validBits
	return nil
}
func (ϟa *GlValidateProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_831_major := uint32(2) // u32
	minRequiredVersion_831_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_831_major, minRequiredVersion_831_minor
	return nil
}
func (ϟa *GlValidateProgramPipeline) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_832_major := uint32(3) // u32
	minRequiredVersion_832_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_832_major, minRequiredVersion_832_minor
	return nil
}
func (ϟa *GlCullFace) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_833_major := uint32(2) // u32
	minRequiredVersion_833_minor := uint32(0) // u32
	switch ϟa.Mode {
	case GLenum_GL_BACK, GLenum_GL_FRONT, GLenum_GL_FRONT_AND_BACK:
	default:
		v := ϟa.Mode
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_835_result := context             // Contextʳ
	ctx := GetContext_835_result                 // Contextʳ
	ctx.Rasterizing.CullFace = ϟa.Mode
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_833_major, minRequiredVersion_833_minor, context, GetContext_835_result, ctx
	return nil
}
func (ϟa *GlDepthRangef) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_836_major := uint32(2)    // u32
	minRequiredVersion_836_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_837_result := context             // Contextʳ
	ctx := GetContext_837_result                 // Contextʳ
	ctx.Rasterizing.DepthNear = ϟa.Near
	ctx.Rasterizing.DepthFar = ϟa.Far
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_836_major, minRequiredVersion_836_minor, context, GetContext_837_result, ctx
	return nil
}
func (ϟa *GlFrontFace) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_838_major := uint32(2) // u32
	minRequiredVersion_838_minor := uint32(0) // u32
	switch ϟa.Orientation {
	case GLenum_GL_CCW, GLenum_GL_CW:
	default:
		v := ϟa.Orientation
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_840_result := context             // Contextʳ
	ctx := GetContext_840_result                 // Contextʳ
	ctx.Rasterizing.FrontFace = ϟa.Orientation
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_838_major, minRequiredVersion_838_minor, context, GetContext_840_result, ctx
	return nil
}
func (ϟa *GlGetMultisamplefv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_841_major := uint32(3) // u32
	minRequiredVersion_841_minor := uint32(1) // u32
	switch ϟa.Pname {
	case GLenum_GL_SAMPLE_POSITION:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_841_major, minRequiredVersion_841_minor
	return nil
}
func (ϟa *GlLineWidth) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_843_major := uint32(2)    // u32
	minRequiredVersion_843_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_844_result := context             // Contextʳ
	ctx := GetContext_844_result                 // Contextʳ
	ctx.Rasterizing.LineWidth = ϟa.Width
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_843_major, minRequiredVersion_843_minor, context, GetContext_844_result, ctx
	return nil
}
func (ϟa *GlPolygonOffset) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_845_major := uint32(2)    // u32
	minRequiredVersion_845_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_846_result := context             // Contextʳ
	ctx := GetContext_846_result                 // Contextʳ
	ctx.Rasterizing.PolygonOffsetUnits = ϟa.Units
	ctx.Rasterizing.PolygonOffsetFactor = ϟa.ScaleFactor
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_845_major, minRequiredVersion_845_minor, context, GetContext_846_result, ctx
	return nil
}
func (ϟa *GlViewport) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_847_major := uint32(2)    // u32
	minRequiredVersion_847_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_848_result := context             // Contextʳ
	ctx := GetContext_848_result                 // Contextʳ
	ctx.Rasterizing.Viewport = func() Rect {
		s := Rect{}
		s.Init()
		s.X = ϟa.X
		s.Y = ϟa.Y
		s.Width = ϟa.Width
		s.Height = ϟa.Height
		return s
	}()
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_847_major, minRequiredVersion_847_minor, context, GetContext_848_result, ctx
	return nil
}
func (ϟa *GlGetBooleani_v) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_849_major := uint32(3) // u32
	minRequiredVersion_849_minor := uint32(1) // u32
	switch ϟa.Target {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_DITHER, GLenum_GL_DRAW_BUFFER, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE, GLenum_GL_VIEWPORT:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_849_major, minRequiredVersion_849_minor
	return nil
}
func (ϟa *GlGetBooleanv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_851_major := uint32(2) // u32
	minRequiredVersion_851_minor := uint32(0) // u32
	switch ϟa.Param {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DITHER, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_VIEWPORT:
	case GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_DRAW_BUFFER, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING:
		minRequiredVersion_852_major := uint32(3) // u32
		minRequiredVersion_852_minor := uint32(0) // u32
		_, _ = minRequiredVersion_852_major, minRequiredVersion_852_minor
	case GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE:
		minRequiredVersion_853_major := uint32(3) // u32
		minRequiredVersion_853_minor := uint32(1) // u32
		_, _ = minRequiredVersion_853_major, minRequiredVersion_853_minor
	default:
		v := ϟa.Param
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	v := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs) // GLbooleanˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                                                // Contextʳ
	GetContext_855_result := context                                                                            // Contextʳ
	ctx := GetContext_855_result                                                                                // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	switch ϟa.Param {
	case GLenum_GL_BLEND:
		v.Index(uint64(0), ϟs).Write(func() GLboolean {
			if ctx.Capabilities.Get(GLenum_GL_BLEND) {
				return 1
			} else {
				return 0
			}
		}(), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_CULL_FACE:
		v.Index(uint64(0), ϟs).Write(func() GLboolean {
			if ctx.Capabilities.Get(GLenum_GL_CULL_FACE) {
				return 1
			} else {
				return 0
			}
		}(), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_DEPTH_TEST:
		v.Index(uint64(0), ϟs).Write(func() GLboolean {
			if ctx.Capabilities.Get(GLenum_GL_DEPTH_TEST) {
				return 1
			} else {
				return 0
			}
		}(), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_DITHER:
		v.Index(uint64(0), ϟs).Write(func() GLboolean {
			if ctx.Capabilities.Get(GLenum_GL_DITHER) {
				return 1
			} else {
				return 0
			}
		}(), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_POLYGON_OFFSET_FILL:
		v.Index(uint64(0), ϟs).Write(func() GLboolean {
			if ctx.Capabilities.Get(GLenum_GL_POLYGON_OFFSET_FILL) {
				return 1
			} else {
				return 0
			}
		}(), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE:
		v.Index(uint64(0), ϟs).Write(func() GLboolean {
			if ctx.Capabilities.Get(GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE) {
				return 1
			} else {
				return 0
			}
		}(), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_SAMPLE_COVERAGE:
		v.Index(uint64(0), ϟs).Write(func() GLboolean {
			if ctx.Capabilities.Get(GLenum_GL_SAMPLE_COVERAGE) {
				return 1
			} else {
				return 0
			}
		}(), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_SCISSOR_TEST:
		v.Index(uint64(0), ϟs).Write(func() GLboolean {
			if ctx.Capabilities.Get(GLenum_GL_SCISSOR_TEST) {
				return 1
			} else {
				return 0
			}
		}(), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_STENCIL_TEST:
		v.Index(uint64(0), ϟs).Write(func() GLboolean {
			if ctx.Capabilities.Get(GLenum_GL_STENCIL_TEST) {
				return 1
			} else {
				return 0
			}
		}(), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_DEPTH_WRITEMASK:
		v.Index(uint64(0), ϟs).Write(ctx.Rasterizing.DepthMask, ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_COLOR_WRITEMASK:
		v.Index(uint64(0), ϟs).Write(ctx.Rasterizing.ColorMaskRed, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(1), ϟs).Write(ctx.Rasterizing.ColorMaskGreen, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(2), ϟs).Write(ctx.Rasterizing.ColorMaskBlue, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(3), ϟs).Write(ctx.Rasterizing.ColorMaskAlpha, ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_SAMPLE_COVERAGE_INVERT:
		v.Index(uint64(0), ϟs).Write(ctx.Rasterizing.SampleCoverageInvert, ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_SHADER_COMPILER:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	default:
		v := ϟa.Param
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	_, _, _, _, _, _ = minRequiredVersion_851_major, minRequiredVersion_851_minor, v, context, GetContext_855_result, ctx
	return nil
}
func (ϟa *GlGetFloatv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_856_major := uint32(2) // u32
	minRequiredVersion_856_minor := uint32(0) // u32
	switch ϟa.Param {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DITHER, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_VIEWPORT:
	case GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_DRAW_BUFFER, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING:
		minRequiredVersion_857_major := uint32(3) // u32
		minRequiredVersion_857_minor := uint32(0) // u32
		_, _ = minRequiredVersion_857_major, minRequiredVersion_857_minor
	case GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE:
		minRequiredVersion_858_major := uint32(3) // u32
		minRequiredVersion_858_minor := uint32(1) // u32
		_, _ = minRequiredVersion_858_major, minRequiredVersion_858_minor
	default:
		v := ϟa.Param
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	v := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs) // GLfloatˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                                                // Contextʳ
	GetContext_860_result := context                                                                            // Contextʳ
	ctx := GetContext_860_result                                                                                // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	switch ϟa.Param {
	case GLenum_GL_DEPTH_RANGE:
		v.Index(uint64(0), ϟs).Write(ctx.Rasterizing.DepthNear, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(1), ϟs).Write(ctx.Rasterizing.DepthFar, ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_LINE_WIDTH:
		v.Index(uint64(0), ϟs).Write(ctx.Rasterizing.LineWidth, ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_POLYGON_OFFSET_FACTOR:
		v.Index(uint64(0), ϟs).Write(ctx.Rasterizing.PolygonOffsetFactor, ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_POLYGON_OFFSET_UNITS:
		v.Index(uint64(0), ϟs).Write(ctx.Rasterizing.PolygonOffsetUnits, ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_SAMPLE_COVERAGE_VALUE:
		v.Index(uint64(0), ϟs).Write(ctx.Rasterizing.SampleCoverageValue, ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_COLOR_CLEAR_VALUE:
		v.Index(uint64(0), ϟs).Write(ctx.Clearing.ClearColor.Red, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(1), ϟs).Write(ctx.Clearing.ClearColor.Green, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(2), ϟs).Write(ctx.Clearing.ClearColor.Blue, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(3), ϟs).Write(ctx.Clearing.ClearColor.Alpha, ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_DEPTH_CLEAR_VALUE:
		v.Index(uint64(0), ϟs).Write(ctx.Clearing.ClearDepth, ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_ALIASED_LINE_WIDTH_RANGE:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(1), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(1), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_ALIASED_POINT_SIZE_RANGE:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(1), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(1), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	default:
		v := ϟa.Param
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	_, _, _, _, _, _ = minRequiredVersion_856_major, minRequiredVersion_856_minor, v, context, GetContext_860_result, ctx
	return nil
}
func (ϟa *GlGetInteger64i_v) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_861_major := uint32(3) // u32
	minRequiredVersion_861_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DITHER, GLenum_GL_DRAW_BUFFER, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING, GLenum_GL_VIEWPORT:
	case GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE:
		minRequiredVersion_862_major := uint32(3) // u32
		minRequiredVersion_862_minor := uint32(1) // u32
		_, _ = minRequiredVersion_862_major, minRequiredVersion_862_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_861_major, minRequiredVersion_861_minor
	return nil
}
func (ϟa *GlGetInteger64v) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_864_major := uint32(3) // u32
	minRequiredVersion_864_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DITHER, GLenum_GL_DRAW_BUFFER, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING, GLenum_GL_VIEWPORT:
	case GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE:
		minRequiredVersion_865_major := uint32(3) // u32
		minRequiredVersion_865_minor := uint32(1) // u32
		_, _ = minRequiredVersion_865_major, minRequiredVersion_865_minor
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_864_major, minRequiredVersion_864_minor
	return nil
}
func (ϟa *GlGetIntegeri_v) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_867_major := uint32(3) // u32
	minRequiredVersion_867_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DITHER, GLenum_GL_DRAW_BUFFER, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING, GLenum_GL_VIEWPORT:
	case GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE:
		minRequiredVersion_868_major := uint32(3) // u32
		minRequiredVersion_868_minor := uint32(1) // u32
		_, _ = minRequiredVersion_868_major, minRequiredVersion_868_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_867_major, minRequiredVersion_867_minor
	return nil
}
func (ϟa *GlGetIntegerv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_870_major := uint32(2) // u32
	minRequiredVersion_870_minor := uint32(0) // u32
	switch ϟa.Param {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DITHER, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_VIEWPORT:
	case GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_DRAW_BUFFER, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING:
		minRequiredVersion_871_major := uint32(3) // u32
		minRequiredVersion_871_minor := uint32(0) // u32
		_, _ = minRequiredVersion_871_major, minRequiredVersion_871_minor
	case GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE:
		minRequiredVersion_872_major := uint32(3) // u32
		minRequiredVersion_872_minor := uint32(1) // u32
		_, _ = minRequiredVersion_872_major, minRequiredVersion_872_minor
	default:
		v := ϟa.Param
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	v := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs) // GLintˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                                                // Contextʳ
	GetContext_874_result := context                                                                            // Contextʳ
	ctx := GetContext_874_result                                                                                // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	switch ϟa.Param {
	case GLenum_GL_ACTIVE_TEXTURE:
		v.Index(uint64(0), ϟs).Write(GLint(ctx.ActiveTextureUnit), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_ARRAY_BUFFER_BINDING:
		v.Index(uint64(0), ϟs).Write(GLint(ctx.BoundBuffers.Get(GLenum_GL_ARRAY_BUFFER)), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING:
		v.Index(uint64(0), ϟs).Write(GLint(ctx.BoundBuffers.Get(GLenum_GL_ELEMENT_ARRAY_BUFFER)), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_BLEND_SRC_ALPHA:
		v.Index(uint64(0), ϟs).Write(GLint(ctx.Blending.SrcAlphaBlendFactor), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_BLEND_SRC_RGB:
		v.Index(uint64(0), ϟs).Write(GLint(ctx.Blending.SrcRgbBlendFactor), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_BLEND_DST_ALPHA:
		v.Index(uint64(0), ϟs).Write(GLint(ctx.Blending.DstAlphaBlendFactor), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_BLEND_DST_RGB:
		v.Index(uint64(0), ϟs).Write(GLint(ctx.Blending.DstRgbBlendFactor), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_BLEND_EQUATION_RGB:
		v.Index(uint64(0), ϟs).Write(GLint(ctx.Blending.BlendEquationRgb), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_BLEND_EQUATION_ALPHA:
		v.Index(uint64(0), ϟs).Write(GLint(ctx.Blending.BlendEquationAlpha), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_BLEND_COLOR:
		v.Index(uint64(0), ϟs).Write(GLint(ctx.Blending.BlendColor.Red), ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(1), ϟs).Write(GLint(ctx.Blending.BlendColor.Green), ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(2), ϟs).Write(GLint(ctx.Blending.BlendColor.Blue), ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(3), ϟs).Write(GLint(ctx.Blending.BlendColor.Alpha), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_DEPTH_FUNC:
		v.Index(uint64(0), ϟs).Write(GLint(ctx.Rasterizing.DepthTestFunction), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_DEPTH_CLEAR_VALUE:
		v.Index(uint64(0), ϟs).Write(GLint(ctx.Clearing.ClearDepth), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_STENCIL_WRITEMASK:
		v.Index(uint64(0), ϟs).Write(GLint(ctx.Rasterizing.StencilMask.Get(GLenum_GL_FRONT)), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_STENCIL_BACK_WRITEMASK:
		v.Index(uint64(0), ϟs).Write(GLint(ctx.Rasterizing.StencilMask.Get(GLenum_GL_BACK)), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_VIEWPORT:
		v.Index(uint64(0), ϟs).Write(ctx.Rasterizing.Viewport.X, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(1), ϟs).Write(ctx.Rasterizing.Viewport.Y, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(2), ϟs).Write(GLint(ctx.Rasterizing.Viewport.Width), ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(3), ϟs).Write(GLint(ctx.Rasterizing.Viewport.Height), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_SCISSOR_BOX:
		v.Index(uint64(0), ϟs).Write(ctx.Rasterizing.Scissor.X, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(1), ϟs).Write(ctx.Rasterizing.Scissor.Y, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(2), ϟs).Write(GLint(ctx.Rasterizing.Scissor.Width), ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(3), ϟs).Write(GLint(ctx.Rasterizing.Scissor.Height), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_FRONT_FACE:
		v.Index(uint64(0), ϟs).Write(GLint(ctx.Rasterizing.FrontFace), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_CULL_FACE_MODE:
		v.Index(uint64(0), ϟs).Write(GLint(ctx.Rasterizing.CullFace), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_STENCIL_CLEAR_VALUE:
		v.Index(uint64(0), ϟs).Write(ctx.Clearing.ClearStencil, ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_DRAW_FRAMEBUFFER_BINDING:
		v.Index(uint64(0), ϟs).Write(GLint(ctx.BoundFramebuffers.Get(GLenum_GL_FRAMEBUFFER)), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_READ_FRAMEBUFFER_BINDING:
		v.Index(uint64(0), ϟs).Write(GLint(ctx.BoundFramebuffers.Get(GLenum_GL_READ_FRAMEBUFFER)), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_RENDERBUFFER_BINDING:
		v.Index(uint64(0), ϟs).Write(GLint(ctx.BoundRenderbuffers.Get(GLenum_GL_RENDERBUFFER)), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_CURRENT_PROGRAM:
		v.Index(uint64(0), ϟs).Write(GLint(ctx.BoundProgram), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_TEXTURE_BINDING_2D:
		v.Index(uint64(0), ϟs).Write(GLint(ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(GLenum_GL_TEXTURE_2D)), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_TEXTURE_BINDING_CUBE_MAP:
		v.Index(uint64(0), ϟs).Write(GLint(ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(GLenum_GL_TEXTURE_CUBE_MAP)), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_GENERATE_MIPMAP_HINT:
		v.Index(uint64(0), ϟs).Write(GLint(ctx.GenerateMipmapHint), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil) // any
		v.Index(uint64(0), ϟs).Write(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil) // any
		v.Index(uint64(0), ϟs).Write(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil) // any
		v.Index(uint64(0), ϟs).Write(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_MAX_RENDERBUFFER_SIZE:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil) // any
		v.Index(uint64(0), ϟs).Write(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_MAX_TEXTURE_IMAGE_UNITS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil) // any
		v.Index(uint64(0), ϟs).Write(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_MAX_TEXTURE_SIZE:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil) // any
		v.Index(uint64(0), ϟs).Write(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_MAX_VARYING_VECTORS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil) // any
		v.Index(uint64(0), ϟs).Write(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_MAX_VERTEX_ATTRIBS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil) // any
		v.Index(uint64(0), ϟs).Write(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil) // any
		v.Index(uint64(0), ϟs).Write(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_MAX_VIEWPORT_DIMS:
		max_width := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil)  // any
		max_height := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(1), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil) // any
		v.Index(uint64(0), ϟs).Write(max_width, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(1), ϟs).Write(max_height, ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = max_width, max_height
	case GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil) // any
		v.Index(uint64(0), ϟs).Write(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_NUM_SHADER_BINARY_FORMATS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil) // any
		v.Index(uint64(0), ϟs).Write(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_PACK_ALIGNMENT:
		v.Index(uint64(0), ϟs).Write(ctx.PixelStorage.Get(GLenum_GL_PACK_ALIGNMENT), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_UNPACK_ALIGNMENT:
		v.Index(uint64(0), ϟs).Write(ctx.PixelStorage.Get(GLenum_GL_UNPACK_ALIGNMENT), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_ALPHA_BITS:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_BLUE_BITS:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_GREEN_BITS:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_RED_BITS:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_DEPTH_BITS:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_SAMPLE_BUFFERS:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_SAMPLES:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_STENCIL_BITS:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_SUBPIXEL_BITS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil) // any
		v.Index(uint64(0), ϟs).Write(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_GPU_DISJOINT_EXT:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	default:
		v := ϟa.Param
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	_, _, _, _, _, _ = minRequiredVersion_870_major, minRequiredVersion_870_minor, v, context, GetContext_874_result, ctx
	return nil
}
func (ϟa *GlGetInternalformativ) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_875_major := uint32(3) // u32
	minRequiredVersion_875_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_RENDERBUFFER:
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_876_major := uint32(3) // u32
		minRequiredVersion_876_minor := uint32(1) // u32
		_, _ = minRequiredVersion_876_major, minRequiredVersion_876_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Internalformat {
	case GLenum_GL_DEPTH24_STENCIL8, GLenum_GL_DEPTH32F_STENCIL8, GLenum_GL_DEPTH_COMPONENT16, GLenum_GL_DEPTH_COMPONENT24, GLenum_GL_DEPTH_COMPONENT32F, GLenum_GL_R11F_G11F_B10F, GLenum_GL_R16F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32F, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_R8_SNORM, GLenum_GL_RG16F, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32F, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RG8_SNORM, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB16F, GLenum_GL_RGB16I, GLenum_GL_RGB16UI, GLenum_GL_RGB32F, GLenum_GL_RGB32I, GLenum_GL_RGB32UI, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGB8, GLenum_GL_RGB8I, GLenum_GL_RGB8UI, GLenum_GL_RGB8_SNORM, GLenum_GL_RGB9_E5, GLenum_GL_RGBA16F, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32F, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA4, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_RGBA8_SNORM, GLenum_GL_SRGB8, GLenum_GL_SRGB8_ALPHA8:
	default:
		v := ϟa.Internalformat
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Pname {
	case GLenum_GL_NUM_SAMPLE_COUNTS, GLenum_GL_SAMPLES:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_875_major, minRequiredVersion_875_minor
	return nil
}
func (ϟa *GlGetString) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_880_major := uint32(2) // u32
	minRequiredVersion_880_minor := uint32(0) // u32
	switch ϟa.Param {
	case GLenum_GL_EXTENSIONS, GLenum_GL_RENDERER, GLenum_GL_SHADING_LANGUAGE_VERSION, GLenum_GL_VENDOR, GLenum_GL_VERSION:
	default:
		v := ϟa.Param
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = GLubyteᶜᵖ{}
	_, _ = minRequiredVersion_880_major, minRequiredVersion_880_minor
	return nil
}
func (ϟa *GlGetStringi) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_882_major := uint32(3) // u32
	minRequiredVersion_882_minor := uint32(0) // u32
	switch ϟa.Name {
	case GLenum_GL_EXTENSIONS:
	default:
		v := ϟa.Name
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_882_major, minRequiredVersion_882_minor
	return nil
}
func (ϟa *GlIsEnabled) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_884_major := uint32(2) // u32
	minRequiredVersion_884_minor := uint32(0) // u32
	switch ϟa.Capability {
	case GLenum_GL_BLEND, GLenum_GL_CULL_FACE, GLenum_GL_DEPTH_TEST, GLenum_GL_DITHER, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SCISSOR_TEST, GLenum_GL_STENCIL_TEST:
	case GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_RASTERIZER_DISCARD:
		minRequiredVersion_885_major := uint32(3) // u32
		minRequiredVersion_885_minor := uint32(0) // u32
		_, _ = minRequiredVersion_885_major, minRequiredVersion_885_minor
	default:
		v := ϟa.Capability
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_887_result := context             // Contextʳ
	ctx := GetContext_887_result                 // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Capabilities.Get(ϟa.Capability) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = minRequiredVersion_884_major, minRequiredVersion_884_minor, context, GetContext_887_result, ctx
	return nil
}
func (ϟa *GlClientWaitSync) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_888_major := uint32(3)                           // u32
	minRequiredVersion_888_minor := uint32(0)                           // u32
	supportsBits_889_seenBits := ϟa.SyncFlags                           // GLbitfield
	supportsBits_889_validBits := GLbitfield_GL_SYNC_FLUSH_COMMANDS_BIT // GLbitfield
	if (GLbitfield_GL_SYNC_FLUSH_COMMANDS_BIT)&(ϟa.SyncFlags) != 0 {
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _, _, _ = minRequiredVersion_888_major, minRequiredVersion_888_minor, supportsBits_889_seenBits, supportsBits_889_validBits
	return nil
}
func (ϟa *GlDeleteSync) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_890_major := uint32(3) // u32
	minRequiredVersion_890_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_890_major, minRequiredVersion_890_minor
	return nil
}
func (ϟa *GlFenceSync) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_891_major := uint32(3) // u32
	minRequiredVersion_891_minor := uint32(0) // u32
	switch ϟa.Condition {
	case GLenum_GL_SYNC_GPU_COMMANDS_COMPLETE:
	default:
		v := ϟa.Condition
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_891_major, minRequiredVersion_891_minor
	return nil
}
func (ϟa *GlGetSynciv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_893_major := uint32(3) // u32
	minRequiredVersion_893_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_OBJECT_TYPE, GLenum_GL_SYNC_CONDITION, GLenum_GL_SYNC_FLAGS, GLenum_GL_SYNC_STATUS:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_893_major, minRequiredVersion_893_minor
	return nil
}
func (ϟa *GlIsSync) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_895_major := uint32(3) // u32
	minRequiredVersion_895_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_895_major, minRequiredVersion_895_minor
	return nil
}
func (ϟa *GlWaitSync) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_896_major := uint32(3) // u32
	minRequiredVersion_896_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_896_major, minRequiredVersion_896_minor
	return nil
}
func (ϟa *GlActiveTexture) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_897_major := uint32(2) // u32
	minRequiredVersion_897_minor := uint32(0) // u32
	switch ϟa.Unit {
	case GLenum_GL_TEXTURE0, GLenum_GL_TEXTURE1, GLenum_GL_TEXTURE10, GLenum_GL_TEXTURE11, GLenum_GL_TEXTURE12, GLenum_GL_TEXTURE13, GLenum_GL_TEXTURE14, GLenum_GL_TEXTURE15, GLenum_GL_TEXTURE16, GLenum_GL_TEXTURE17, GLenum_GL_TEXTURE18, GLenum_GL_TEXTURE19, GLenum_GL_TEXTURE2, GLenum_GL_TEXTURE20, GLenum_GL_TEXTURE21, GLenum_GL_TEXTURE22, GLenum_GL_TEXTURE23, GLenum_GL_TEXTURE24, GLenum_GL_TEXTURE25, GLenum_GL_TEXTURE26, GLenum_GL_TEXTURE27, GLenum_GL_TEXTURE28, GLenum_GL_TEXTURE29, GLenum_GL_TEXTURE3, GLenum_GL_TEXTURE30, GLenum_GL_TEXTURE31, GLenum_GL_TEXTURE4, GLenum_GL_TEXTURE5, GLenum_GL_TEXTURE6, GLenum_GL_TEXTURE7, GLenum_GL_TEXTURE8, GLenum_GL_TEXTURE9:
	default:
		v := ϟa.Unit
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_899_result := context             // Contextʳ
	ctx := GetContext_899_result                 // Contextʳ
	ctx.ActiveTextureUnit = ϟa.Unit
	if !(ctx.TextureUnits.Contains(ϟa.Unit)) {
		ctx.TextureUnits[ϟa.Unit] = ctx.TextureUnits.Get(ϟa.Unit)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_897_major, minRequiredVersion_897_minor, context, GetContext_899_result, ctx
	return nil
}
func (ϟa *GlBindImageTexture) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_900_major := uint32(3) // u32
	minRequiredVersion_900_minor := uint32(1) // u32
	switch ϟa.Access {
	case GLenum_GL_READ_ONLY, GLenum_GL_READ_WRITE, GLenum_GL_WRITE_ONLY:
	default:
		v := ϟa.Access
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Format {
	case GLenum_GL_R32F, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_RGBA16F, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32F, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_RGBA8_SNORM:
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_900_major, minRequiredVersion_900_minor
	return nil
}
func (ϟa *GlBindSampler) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_903_major := uint32(3) // u32
	minRequiredVersion_903_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_903_major, minRequiredVersion_903_minor
	return nil
}
func (ϟa *GlBindTexture) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_904_major := uint32(2) // u32
	minRequiredVersion_904_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_905_major := uint32(3) // u32
		minRequiredVersion_905_minor := uint32(0) // u32
		_, _ = minRequiredVersion_905_major, minRequiredVersion_905_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_906_major := uint32(3) // u32
		minRequiredVersion_906_minor := uint32(1) // u32
		_, _ = minRequiredVersion_906_major, minRequiredVersion_906_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_908_result := context             // Contextʳ
	ctx := GetContext_908_result                 // Contextʳ
	if !(ctx.Instances.Textures.Contains(ϟa.Texture)) {
		ctx.Instances.Textures[ϟa.Texture] = func() *Texture {
			s := &Texture{}
			s.Init()
			s.ID = ϟa.Texture
			return s
		}()
	}
	ctx.TextureUnits.Get(ctx.ActiveTextureUnit)[ϟa.Target] = ϟa.Texture
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_904_major, minRequiredVersion_904_minor, context, GetContext_908_result, ctx
	return nil
}
func (ϟa *GlCompressedTexImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_909_major := uint32(2) // u32
	minRequiredVersion_909_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Format {
	case GLenum_GL_COMPRESSED_R11_EAC, GLenum_GL_COMPRESSED_RG11_EAC, GLenum_GL_COMPRESSED_RGB8_ETC2, GLenum_GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_COMPRESSED_RGBA8_ETC2_EAC, GLenum_GL_COMPRESSED_SIGNED_R11_EAC, GLenum_GL_COMPRESSED_SIGNED_RG11_EAC, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC, GLenum_GL_COMPRESSED_SRGB8_ETC2, GLenum_GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2:
		minRequiredVersion_911_major := uint32(3) // u32
		minRequiredVersion_911_minor := uint32(0) // u32
		_, _ = minRequiredVersion_911_major, minRequiredVersion_911_minor
	case GLenum_GL_ATC_RGB_AMD, GLenum_GL_ATC_RGBA_EXPLICIT_ALPHA_AMD, GLenum_GL_ATC_RGBA_INTERPOLATED_ALPHA_AMD:
		requiresExtension_912_ext := ExtensionId_GL_AMD_compressed_ATC_texture // ExtensionId
		_ = requiresExtension_912_ext
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_914_result := context             // Contextʳ
	ctx := GetContext_914_result                 // Contextʳ
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D:
		id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(GLenum_GL_TEXTURE_2D) // TextureId
		t := ctx.Instances.Textures.Get(id)                                         // Textureʳ
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.Width
			s.Height = ϟa.Height
			s.Size = uint32(ϟa.ImageSize)
			s.Format = ϟa.Format
			return s
		}() // Image
		if ((ctx.BoundBuffers.Get(GLenum_GL_PIXEL_UNPACK_BUFFER)) == (BufferId(uint32(0)))) && ((ϟa.Data) != (TexturePointer(Voidᶜᵖ{}))) {
			l.Data = U8ᵖ(ϟa.Data).Slice(uint64(uint32(0)), uint64(l.Size), ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
		}
		t.Texture2D[ϟa.Level] = l
		t.Kind = TextureKind_TEXTURE2D
		t.Format = ϟa.Format
		_, _, _ = id, t, l
	case GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z:
		id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(GLenum_GL_TEXTURE_CUBE_MAP) // TextureId
		t := ctx.Instances.Textures.Get(id)                                               // Textureʳ
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.Width
			s.Height = ϟa.Height
			s.Size = uint32(ϟa.ImageSize)
			s.Format = ϟa.Format
			return s
		}() // Image
		if ((ctx.BoundBuffers.Get(GLenum_GL_PIXEL_UNPACK_BUFFER)) == (BufferId(uint32(0)))) && ((ϟa.Data) != (TexturePointer(Voidᶜᵖ{}))) {
			l.Data = U8ᵖ(ϟa.Data).Slice(uint64(uint32(0)), uint64(l.Size), ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
		}
		cube := t.Cubemap.Get(ϟa.Level) // CubemapLevel
		cube.Faces[ϟa.Target] = l
		t.Cubemap[ϟa.Level] = cube
		t.Kind = TextureKind_CUBEMAP
		t.Format = ϟa.Format
		_, _, _, _ = id, t, l, cube
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_909_major, minRequiredVersion_909_minor, context, GetContext_914_result, ctx
	return nil
}
func (ϟa *GlCompressedTexImage3D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_915_major := uint32(3) // u32
	minRequiredVersion_915_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Internalformat {
	case GLenum_GL_COMPRESSED_R11_EAC, GLenum_GL_COMPRESSED_RG11_EAC, GLenum_GL_COMPRESSED_RGB8_ETC2, GLenum_GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_COMPRESSED_RGBA8_ETC2_EAC, GLenum_GL_COMPRESSED_SIGNED_R11_EAC, GLenum_GL_COMPRESSED_SIGNED_RG11_EAC, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC, GLenum_GL_COMPRESSED_SRGB8_ETC2, GLenum_GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2:
	default:
		v := ϟa.Internalformat
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_915_major, minRequiredVersion_915_minor
	return nil
}
func (ϟa *GlCompressedTexSubImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_918_major := uint32(2) // u32
	minRequiredVersion_918_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Format {
	case GLenum_GL_COMPRESSED_R11_EAC, GLenum_GL_COMPRESSED_RG11_EAC, GLenum_GL_COMPRESSED_RGB8_ETC2, GLenum_GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_COMPRESSED_RGBA8_ETC2_EAC, GLenum_GL_COMPRESSED_SIGNED_R11_EAC, GLenum_GL_COMPRESSED_SIGNED_RG11_EAC, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC, GLenum_GL_COMPRESSED_SRGB8_ETC2, GLenum_GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2:
		minRequiredVersion_920_major := uint32(3) // u32
		minRequiredVersion_920_minor := uint32(0) // u32
		_, _ = minRequiredVersion_920_major, minRequiredVersion_920_minor
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_918_major, minRequiredVersion_918_minor
	return nil
}
func (ϟa *GlCompressedTexSubImage3D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_922_major := uint32(3) // u32
	minRequiredVersion_922_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Format {
	case GLenum_GL_COMPRESSED_R11_EAC, GLenum_GL_COMPRESSED_RG11_EAC, GLenum_GL_COMPRESSED_RGB8_ETC2, GLenum_GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_COMPRESSED_RGBA8_ETC2_EAC, GLenum_GL_COMPRESSED_SIGNED_R11_EAC, GLenum_GL_COMPRESSED_SIGNED_RG11_EAC, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC, GLenum_GL_COMPRESSED_SRGB8_ETC2, GLenum_GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2:
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_922_major, minRequiredVersion_922_minor
	return nil
}
func (ϟa *GlCopyTexImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_925_major := uint32(2) // u32
	minRequiredVersion_925_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Format {
	case GLenum_GL_ALPHA, GLenum_GL_LUMINANCE, GLenum_GL_LUMINANCE_ALPHA, GLenum_GL_RGB, GLenum_GL_RGBA:
	case GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGB8, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA4, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_SRGB8, GLenum_GL_SRGB8_ALPHA8:
		minRequiredVersion_927_major := uint32(3) // u32
		minRequiredVersion_927_minor := uint32(0) // u32
		_, _ = minRequiredVersion_927_major, minRequiredVersion_927_minor
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_925_major, minRequiredVersion_925_minor
	return nil
}
func (ϟa *GlCopyTexSubImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_929_major := uint32(2) // u32
	minRequiredVersion_929_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_929_major, minRequiredVersion_929_minor
	return nil
}
func (ϟa *GlCopyTexSubImage3D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_931_major := uint32(3) // u32
	minRequiredVersion_931_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_931_major, minRequiredVersion_931_minor
	return nil
}
func (ϟa *GlDeleteSamplers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_933_major := uint32(3) // u32
	minRequiredVersion_933_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_933_major, minRequiredVersion_933_minor
	return nil
}
func (ϟa *GlDeleteTextures) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_934_major := uint32(2)                               // u32
	minRequiredVersion_934_minor := uint32(0)                               // u32
	t := ϟa.Textures.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // TextureIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                            // Contextʳ
	GetContext_935_result := context                                        // Contextʳ
	ctx := GetContext_935_result                                            // Contextʳ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.Textures, t.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_934_major, minRequiredVersion_934_minor, t, context, GetContext_935_result, ctx
	return nil
}
func (ϟa *GlGenSamplers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_936_major := uint32(3) // u32
	minRequiredVersion_936_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_936_major, minRequiredVersion_936_minor
	return nil
}
func (ϟa *GlGenTextures) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_937_major := uint32(2)                               // u32
	minRequiredVersion_937_minor := uint32(0)                               // u32
	t := ϟa.Textures.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // TextureIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                            // Contextʳ
	GetContext_938_result := context                                        // Contextʳ
	ctx := GetContext_938_result                                            // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := TextureId(ϟa.Textures.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil)) // TextureId
		ctx.Instances.Textures[id] = func() *Texture {
			s := &Texture{}
			s.Init()
			s.ID = id
			return s
		}()
		t.Index(uint64(i), ϟs).Write(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = minRequiredVersion_937_major, minRequiredVersion_937_minor, t, context, GetContext_938_result, ctx
	return nil
}
func (ϟa *GlGenerateMipmap) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_939_major := uint32(2) // u32
	minRequiredVersion_939_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_940_major := uint32(3) // u32
		minRequiredVersion_940_minor := uint32(0) // u32
		_, _ = minRequiredVersion_940_major, minRequiredVersion_940_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_939_major, minRequiredVersion_939_minor
	return nil
}
func (ϟa *GlGetSamplerParameterfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_942_major := uint32(3) // u32
	minRequiredVersion_942_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_942_major, minRequiredVersion_942_minor
	return nil
}
func (ϟa *GlGetSamplerParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_944_major := uint32(3) // u32
	minRequiredVersion_944_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_944_major, minRequiredVersion_944_minor
	return nil
}
func (ϟa *GlGetTexLevelParameterfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_946_major := uint32(3) // u32
	minRequiredVersion_946_minor := uint32(1) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_2D_MULTISAMPLE, GLenum_GL_TEXTURE_3D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_ALPHA_SIZE, GLenum_GL_TEXTURE_ALPHA_TYPE, GLenum_GL_TEXTURE_BLUE_SIZE, GLenum_GL_TEXTURE_BLUE_TYPE, GLenum_GL_TEXTURE_COMPRESSED, GLenum_GL_TEXTURE_DEPTH, GLenum_GL_TEXTURE_DEPTH_SIZE, GLenum_GL_TEXTURE_DEPTH_TYPE, GLenum_GL_TEXTURE_FIXED_SAMPLE_LOCATIONS, GLenum_GL_TEXTURE_GREEN_SIZE, GLenum_GL_TEXTURE_GREEN_TYPE, GLenum_GL_TEXTURE_HEIGHT, GLenum_GL_TEXTURE_INTERNAL_FORMAT, GLenum_GL_TEXTURE_RED_SIZE, GLenum_GL_TEXTURE_RED_TYPE, GLenum_GL_TEXTURE_SAMPLES, GLenum_GL_TEXTURE_SHARED_SIZE, GLenum_GL_TEXTURE_STENCIL_SIZE, GLenum_GL_TEXTURE_WIDTH:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_946_major, minRequiredVersion_946_minor
	return nil
}
func (ϟa *GlGetTexLevelParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_949_major := uint32(3) // u32
	minRequiredVersion_949_minor := uint32(1) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_2D_MULTISAMPLE, GLenum_GL_TEXTURE_3D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_ALPHA_SIZE, GLenum_GL_TEXTURE_ALPHA_TYPE, GLenum_GL_TEXTURE_BLUE_SIZE, GLenum_GL_TEXTURE_BLUE_TYPE, GLenum_GL_TEXTURE_COMPRESSED, GLenum_GL_TEXTURE_DEPTH, GLenum_GL_TEXTURE_DEPTH_SIZE, GLenum_GL_TEXTURE_DEPTH_TYPE, GLenum_GL_TEXTURE_FIXED_SAMPLE_LOCATIONS, GLenum_GL_TEXTURE_GREEN_SIZE, GLenum_GL_TEXTURE_GREEN_TYPE, GLenum_GL_TEXTURE_HEIGHT, GLenum_GL_TEXTURE_INTERNAL_FORMAT, GLenum_GL_TEXTURE_RED_SIZE, GLenum_GL_TEXTURE_RED_TYPE, GLenum_GL_TEXTURE_SAMPLES, GLenum_GL_TEXTURE_SHARED_SIZE, GLenum_GL_TEXTURE_STENCIL_SIZE, GLenum_GL_TEXTURE_WIDTH:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_949_major, minRequiredVersion_949_minor
	return nil
}
func (ϟa *GlGetTexParameterfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_952_major := uint32(2) // u32
	minRequiredVersion_952_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_953_major := uint32(3) // u32
		minRequiredVersion_953_minor := uint32(0) // u32
		_, _ = minRequiredVersion_953_major, minRequiredVersion_953_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_954_major := uint32(3) // u32
		minRequiredVersion_954_minor := uint32(1) // u32
		_, _ = minRequiredVersion_954_major, minRequiredVersion_954_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Parameter {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_IMMUTABLE_FORMAT, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_956_major := uint32(3) // u32
		minRequiredVersion_956_minor := uint32(0) // u32
		_, _ = minRequiredVersion_956_major, minRequiredVersion_956_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE, GLenum_GL_IMAGE_FORMAT_COMPATIBILITY_TYPE, GLenum_GL_TEXTURE_IMMUTABLE_LEVELS:
		minRequiredVersion_957_major := uint32(3) // u32
		minRequiredVersion_957_minor := uint32(1) // u32
		_, _ = minRequiredVersion_957_major, minRequiredVersion_957_minor
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                     // Contextʳ
	GetContext_959_result := context                                 // Contextʳ
	ctx := GetContext_959_result                                     // Contextʳ
	id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(ϟa.Target) // TextureId
	t := ctx.Instances.Textures.Get(id)                              // Textureʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Values.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(func() (result GLfloat) {
		switch ϟa.Parameter {
		case GLenum_GL_TEXTURE_MAG_FILTER:
			return GLfloat(t.MagFilter)
		case GLenum_GL_TEXTURE_MIN_FILTER:
			return GLfloat(t.MinFilter)
		case GLenum_GL_TEXTURE_WRAP_S:
			return GLfloat(t.WrapS)
		case GLenum_GL_TEXTURE_WRAP_T:
			return GLfloat(t.WrapT)
		case GLenum_GL_TEXTURE_MAX_ANISOTROPY_EXT:
			return GLfloat(t.MaxAnisotropy)
		case GLenum_GL_TEXTURE_SWIZZLE_R:
			return GLfloat(t.SwizzleR)
		case GLenum_GL_TEXTURE_SWIZZLE_G:
			return GLfloat(t.SwizzleG)
		case GLenum_GL_TEXTURE_SWIZZLE_B:
			return GLfloat(t.SwizzleB)
		case GLenum_GL_TEXTURE_SWIZZLE_A:
			return GLfloat(t.SwizzleA)
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.Parameter, ϟa))
			return result
		}
	}(), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _, _ = minRequiredVersion_952_major, minRequiredVersion_952_minor, context, GetContext_959_result, ctx, id, t
	return nil
}
func (ϟa *GlGetTexParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_960_major := uint32(2) // u32
	minRequiredVersion_960_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_961_major := uint32(3) // u32
		minRequiredVersion_961_minor := uint32(0) // u32
		_, _ = minRequiredVersion_961_major, minRequiredVersion_961_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_962_major := uint32(3) // u32
		minRequiredVersion_962_minor := uint32(1) // u32
		_, _ = minRequiredVersion_962_major, minRequiredVersion_962_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Parameter {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_IMMUTABLE_FORMAT, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_964_major := uint32(3) // u32
		minRequiredVersion_964_minor := uint32(0) // u32
		_, _ = minRequiredVersion_964_major, minRequiredVersion_964_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE, GLenum_GL_IMAGE_FORMAT_COMPATIBILITY_TYPE, GLenum_GL_TEXTURE_IMMUTABLE_LEVELS:
		minRequiredVersion_965_major := uint32(3) // u32
		minRequiredVersion_965_minor := uint32(1) // u32
		_, _ = minRequiredVersion_965_major, minRequiredVersion_965_minor
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                     // Contextʳ
	GetContext_967_result := context                                 // Contextʳ
	ctx := GetContext_967_result                                     // Contextʳ
	id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(ϟa.Target) // TextureId
	t := ctx.Instances.Textures.Get(id)                              // Textureʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Values.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(func() (result GLint) {
		switch ϟa.Parameter {
		case GLenum_GL_TEXTURE_MAG_FILTER:
			return GLint(t.MagFilter)
		case GLenum_GL_TEXTURE_MIN_FILTER:
			return GLint(t.MinFilter)
		case GLenum_GL_TEXTURE_WRAP_S:
			return GLint(t.WrapS)
		case GLenum_GL_TEXTURE_WRAP_T:
			return GLint(t.WrapT)
		case GLenum_GL_TEXTURE_MAX_ANISOTROPY_EXT:
			return GLint(t.MaxAnisotropy)
		case GLenum_GL_TEXTURE_SWIZZLE_R:
			return GLint(t.SwizzleR)
		case GLenum_GL_TEXTURE_SWIZZLE_G:
			return GLint(t.SwizzleG)
		case GLenum_GL_TEXTURE_SWIZZLE_B:
			return GLint(t.SwizzleB)
		case GLenum_GL_TEXTURE_SWIZZLE_A:
			return GLint(t.SwizzleA)
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.Parameter, ϟa))
			return result
		}
	}(), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _, _ = minRequiredVersion_960_major, minRequiredVersion_960_minor, context, GetContext_967_result, ctx, id, t
	return nil
}
func (ϟa *GlIsSampler) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_968_major := uint32(3) // u32
	minRequiredVersion_968_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_968_major, minRequiredVersion_968_minor
	return nil
}
func (ϟa *GlIsTexture) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_969_major := uint32(2)    // u32
	minRequiredVersion_969_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_970_result := context             // Contextʳ
	ctx := GetContext_970_result                 // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.Textures.Contains(ϟa.Texture) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = minRequiredVersion_969_major, minRequiredVersion_969_minor, context, GetContext_970_result, ctx
	return nil
}
func (ϟa *GlPixelStorei) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_971_major := uint32(2) // u32
	minRequiredVersion_971_minor := uint32(0) // u32
	switch ϟa.Parameter {
	case GLenum_GL_PACK_ALIGNMENT, GLenum_GL_UNPACK_ALIGNMENT:
	case GLenum_GL_PACK_IMAGE_HEIGHT, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_IMAGES, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS:
		minRequiredVersion_972_major := uint32(3) // u32
		minRequiredVersion_972_minor := uint32(0) // u32
		_, _ = minRequiredVersion_972_major, minRequiredVersion_972_minor
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_974_result := context             // Contextʳ
	ctx := GetContext_974_result                 // Contextʳ
	ctx.PixelStorage[ϟa.Parameter] = ϟa.Value
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_971_major, minRequiredVersion_971_minor, context, GetContext_974_result, ctx
	return nil
}
func (ϟa *GlSamplerParameterf) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_975_major := uint32(3) // u32
	minRequiredVersion_975_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_975_major, minRequiredVersion_975_minor
	return nil
}
func (ϟa *GlSamplerParameterfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_977_major := uint32(3) // u32
	minRequiredVersion_977_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_977_major, minRequiredVersion_977_minor
	return nil
}
func (ϟa *GlSamplerParameteri) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_979_major := uint32(3) // u32
	minRequiredVersion_979_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_979_major, minRequiredVersion_979_minor
	return nil
}
func (ϟa *GlSamplerParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_981_major := uint32(3) // u32
	minRequiredVersion_981_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_981_major, minRequiredVersion_981_minor
	return nil
}
func (ϟa *GlTexImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_983_major := uint32(2) // u32
	minRequiredVersion_983_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Format {
	case GLenum_GL_ALPHA, GLenum_GL_LUMINANCE, GLenum_GL_LUMINANCE_ALPHA, GLenum_GL_RGB, GLenum_GL_RGBA:
	case GLenum_GL_DEPTH_COMPONENT, GLenum_GL_DEPTH_STENCIL, GLenum_GL_RED, GLenum_GL_RED_INTEGER, GLenum_GL_RG, GLenum_GL_RGBA_INTEGER, GLenum_GL_RGB_INTEGER, GLenum_GL_RG_INTEGER:
		minRequiredVersion_985_major := uint32(3) // u32
		minRequiredVersion_985_minor := uint32(0) // u32
		_, _ = minRequiredVersion_985_major, minRequiredVersion_985_minor
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Type {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_SHORT_4_4_4_4, GLenum_GL_UNSIGNED_SHORT_5_5_5_1, GLenum_GL_UNSIGNED_SHORT_5_6_5:
	case GLenum_GL_HALF_FLOAT_OES:
		requiresExtension_987_ext := ExtensionId_GL_OES_texture_half_float // ExtensionId
		_ = requiresExtension_987_ext
	case GLenum_GL_BYTE, GLenum_GL_FLOAT, GLenum_GL_FLOAT_32_UNSIGNED_INT_24_8_REV, GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_10F_11F_11F_REV, GLenum_GL_UNSIGNED_INT_24_8, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_INT_5_9_9_9_REV, GLenum_GL_UNSIGNED_SHORT:
		minRequiredVersion_988_major := uint32(3) // u32
		minRequiredVersion_988_minor := uint32(0) // u32
		_, _ = minRequiredVersion_988_major, minRequiredVersion_988_minor
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_990_result := context             // Contextʳ
	ctx := GetContext_990_result                 // Contextʳ
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D:
		id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(GLenum_GL_TEXTURE_2D) // TextureId
		t := ctx.Instances.Textures.Get(id)                                         // Textureʳ
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.Width
			s.Height = ϟa.Height
			s.Size = externs{ϟa, ϟs, ϟd, ϟl, ϟb}.imageSize(uint32(ϟa.Width), uint32(ϟa.Height), ϟa.Format, ϟa.Type)
			s.Format = ϟa.Format
			return s
		}() // Image
		if (ϟa.Data) != (TexturePointer(Voidᶜᵖ{})) {
			if (ctx.BoundBuffers.Get(GLenum_GL_PIXEL_UNPACK_BUFFER)) == (BufferId(uint32(0))) {
				l.Data = U8ᵖ(ϟa.Data).Slice(uint64(uint32(0)), uint64(l.Size), ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
			}
		} else {
			l.Data = MakeU8ˢ(uint64(l.Size), ϟs)
		}
		t.Texture2D[ϟa.Level] = l
		t.Kind = TextureKind_TEXTURE2D
		t.Format = ϟa.Format
		_, _, _ = id, t, l
	case GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z:
		id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(GLenum_GL_TEXTURE_CUBE_MAP) // TextureId
		t := ctx.Instances.Textures.Get(id)                                               // Textureʳ
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.Width
			s.Height = ϟa.Height
			s.Size = externs{ϟa, ϟs, ϟd, ϟl, ϟb}.imageSize(uint32(ϟa.Width), uint32(ϟa.Height), ϟa.Format, ϟa.Type)
			s.Format = ϟa.Format
			return s
		}() // Image
		if (ϟa.Data) != (TexturePointer(Voidᶜᵖ{})) {
			if (ctx.BoundBuffers.Get(GLenum_GL_PIXEL_UNPACK_BUFFER)) == (BufferId(uint32(0))) {
				l.Data = U8ᵖ(ϟa.Data).Slice(uint64(uint32(0)), uint64(l.Size), ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
			}
		} else {
			l.Data = MakeU8ˢ(uint64(l.Size), ϟs)
		}
		cube := t.Cubemap.Get(ϟa.Level) // CubemapLevel
		cube.Faces[ϟa.Target] = l
		t.Cubemap[ϟa.Level] = cube
		t.Kind = TextureKind_CUBEMAP
		t.Format = ϟa.Format
		_, _, _, _ = id, t, l, cube
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_983_major, minRequiredVersion_983_minor, context, GetContext_990_result, ctx
	return nil
}
func (ϟa *GlTexImage3D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_991_major := uint32(3) // u32
	minRequiredVersion_991_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Format {
	case GLenum_GL_ALPHA, GLenum_GL_DEPTH_COMPONENT, GLenum_GL_DEPTH_STENCIL, GLenum_GL_LUMINANCE, GLenum_GL_LUMINANCE_ALPHA, GLenum_GL_RED, GLenum_GL_RED_INTEGER, GLenum_GL_RG, GLenum_GL_RGB, GLenum_GL_RGBA, GLenum_GL_RGBA_INTEGER, GLenum_GL_RGB_INTEGER, GLenum_GL_RG_INTEGER:
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Type {
	case GLenum_GL_HALF_FLOAT_OES:
		requiresExtension_994_ext := ExtensionId_GL_OES_texture_half_float // ExtensionId
		_ = requiresExtension_994_ext
	case GLenum_GL_BYTE, GLenum_GL_FLOAT, GLenum_GL_FLOAT_32_UNSIGNED_INT_24_8_REV, GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_10F_11F_11F_REV, GLenum_GL_UNSIGNED_INT_24_8, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_INT_5_9_9_9_REV, GLenum_GL_UNSIGNED_SHORT, GLenum_GL_UNSIGNED_SHORT_4_4_4_4, GLenum_GL_UNSIGNED_SHORT_5_5_5_1, GLenum_GL_UNSIGNED_SHORT_5_6_5:
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_991_major, minRequiredVersion_991_minor
	return nil
}
func (ϟa *GlTexParameterf) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_996_major := uint32(2) // u32
	minRequiredVersion_996_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_997_major := uint32(3) // u32
		minRequiredVersion_997_minor := uint32(0) // u32
		_, _ = minRequiredVersion_997_major, minRequiredVersion_997_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_998_major := uint32(3) // u32
		minRequiredVersion_998_minor := uint32(1) // u32
		_, _ = minRequiredVersion_998_major, minRequiredVersion_998_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Parameter {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_1000_major := uint32(3) // u32
		minRequiredVersion_1000_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1000_major, minRequiredVersion_1000_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE:
		minRequiredVersion_1001_major := uint32(3) // u32
		minRequiredVersion_1001_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1001_major, minRequiredVersion_1001_minor
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                     // Contextʳ
	GetContext_1003_result := context                                // Contextʳ
	ctx := GetContext_1003_result                                    // Contextʳ
	id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(ϟa.Target) // TextureId
	t := ctx.Instances.Textures.Get(id)                              // Textureʳ
	switch ϟa.Parameter {
	case GLenum_GL_TEXTURE_MAG_FILTER:
		t.MagFilter = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_MIN_FILTER:
		t.MinFilter = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_WRAP_S:
		t.WrapS = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_WRAP_T:
		t.WrapT = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_MAX_ANISOTROPY_EXT:
		t.MaxAnisotropy = float32(ϟa.Value)
	case GLenum_GL_TEXTURE_SWIZZLE_R:
		t.SwizzleR = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_SWIZZLE_G:
		t.SwizzleG = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_SWIZZLE_B:
		t.SwizzleB = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_SWIZZLE_A:
		t.SwizzleA = GLenum(ϟa.Value)
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = minRequiredVersion_996_major, minRequiredVersion_996_minor, context, GetContext_1003_result, ctx, id, t
	return nil
}
func (ϟa *GlTexParameterfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1004_major := uint32(2) // u32
	minRequiredVersion_1004_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_1005_major := uint32(3) // u32
		minRequiredVersion_1005_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1005_major, minRequiredVersion_1005_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_1006_major := uint32(3) // u32
		minRequiredVersion_1006_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1006_major, minRequiredVersion_1006_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_1008_major := uint32(3) // u32
		minRequiredVersion_1008_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1008_major, minRequiredVersion_1008_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE:
		minRequiredVersion_1009_major := uint32(3) // u32
		minRequiredVersion_1009_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1009_major, minRequiredVersion_1009_minor
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1004_major, minRequiredVersion_1004_minor
	return nil
}
func (ϟa *GlTexParameteri) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1011_major := uint32(2) // u32
	minRequiredVersion_1011_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_1012_major := uint32(3) // u32
		minRequiredVersion_1012_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1012_major, minRequiredVersion_1012_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_1013_major := uint32(3) // u32
		minRequiredVersion_1013_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1013_major, minRequiredVersion_1013_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Parameter {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_1015_major := uint32(3) // u32
		minRequiredVersion_1015_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1015_major, minRequiredVersion_1015_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE:
		minRequiredVersion_1016_major := uint32(3) // u32
		minRequiredVersion_1016_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1016_major, minRequiredVersion_1016_minor
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                     // Contextʳ
	GetContext_1018_result := context                                // Contextʳ
	ctx := GetContext_1018_result                                    // Contextʳ
	id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(ϟa.Target) // TextureId
	t := ctx.Instances.Textures.Get(id)                              // Textureʳ
	switch ϟa.Parameter {
	case GLenum_GL_TEXTURE_MAG_FILTER:
		t.MagFilter = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_MIN_FILTER:
		t.MinFilter = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_WRAP_S:
		t.WrapS = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_WRAP_T:
		t.WrapT = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_MAX_ANISOTROPY_EXT:
		t.MaxAnisotropy = float32(ϟa.Value)
	case GLenum_GL_TEXTURE_SWIZZLE_R:
		t.SwizzleR = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_SWIZZLE_G:
		t.SwizzleG = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_SWIZZLE_B:
		t.SwizzleB = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_SWIZZLE_A:
		t.SwizzleA = GLenum(ϟa.Value)
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = minRequiredVersion_1011_major, minRequiredVersion_1011_minor, context, GetContext_1018_result, ctx, id, t
	return nil
}
func (ϟa *GlTexParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1019_major := uint32(2) // u32
	minRequiredVersion_1019_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_1020_major := uint32(3) // u32
		minRequiredVersion_1020_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1020_major, minRequiredVersion_1020_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_1021_major := uint32(3) // u32
		minRequiredVersion_1021_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1021_major, minRequiredVersion_1021_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_1023_major := uint32(3) // u32
		minRequiredVersion_1023_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1023_major, minRequiredVersion_1023_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE:
		minRequiredVersion_1024_major := uint32(3) // u32
		minRequiredVersion_1024_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1024_major, minRequiredVersion_1024_minor
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1019_major, minRequiredVersion_1019_minor
	return nil
}
func (ϟa *GlTexStorage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1026_major := uint32(3) // u32
	minRequiredVersion_1026_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Internalformat {
	case GLenum_GL_COMPRESSED_R11_EAC, GLenum_GL_COMPRESSED_RG11_EAC, GLenum_GL_COMPRESSED_RGB8_ETC2, GLenum_GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_COMPRESSED_RGBA8_ETC2_EAC, GLenum_GL_COMPRESSED_SIGNED_R11_EAC, GLenum_GL_COMPRESSED_SIGNED_RG11_EAC, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC, GLenum_GL_COMPRESSED_SRGB8_ETC2, GLenum_GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_DEPTH24_STENCIL8, GLenum_GL_DEPTH32F_STENCIL8, GLenum_GL_DEPTH_COMPONENT16, GLenum_GL_DEPTH_COMPONENT24, GLenum_GL_DEPTH_COMPONENT32F, GLenum_GL_R11F_G11F_B10F, GLenum_GL_R16F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32F, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_R8_SNORM, GLenum_GL_RG16F, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32F, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RG8_SNORM, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB16F, GLenum_GL_RGB16I, GLenum_GL_RGB16UI, GLenum_GL_RGB32F, GLenum_GL_RGB32I, GLenum_GL_RGB32UI, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGB8, GLenum_GL_RGB8I, GLenum_GL_RGB8UI, GLenum_GL_RGB8_SNORM, GLenum_GL_RGB9_E5, GLenum_GL_RGBA16F, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32F, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA4, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_RGBA8_SNORM, GLenum_GL_SRGB8, GLenum_GL_SRGB8_ALPHA8:
	default:
		v := ϟa.Internalformat
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1026_major, minRequiredVersion_1026_minor
	return nil
}
func (ϟa *GlTexStorage2DMultisample) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1029_major := uint32(3) // u32
	minRequiredVersion_1029_minor := uint32(1) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Internalformat {
	case GLenum_GL_DEPTH24_STENCIL8, GLenum_GL_DEPTH32F_STENCIL8, GLenum_GL_DEPTH_COMPONENT16, GLenum_GL_DEPTH_COMPONENT24, GLenum_GL_DEPTH_COMPONENT32F, GLenum_GL_R11F_G11F_B10F, GLenum_GL_R16F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32F, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_R8_SNORM, GLenum_GL_RG16F, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32F, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RG8_SNORM, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB16F, GLenum_GL_RGB16I, GLenum_GL_RGB16UI, GLenum_GL_RGB32F, GLenum_GL_RGB32I, GLenum_GL_RGB32UI, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGB8, GLenum_GL_RGB8I, GLenum_GL_RGB8UI, GLenum_GL_RGB8_SNORM, GLenum_GL_RGB9_E5, GLenum_GL_RGBA16F, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32F, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA4, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_RGBA8_SNORM, GLenum_GL_SRGB8, GLenum_GL_SRGB8_ALPHA8:
	default:
		v := ϟa.Internalformat
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1029_major, minRequiredVersion_1029_minor
	return nil
}
func (ϟa *GlTexStorage3D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1032_major := uint32(3) // u32
	minRequiredVersion_1032_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Internalformat {
	case GLenum_GL_COMPRESSED_R11_EAC, GLenum_GL_COMPRESSED_RG11_EAC, GLenum_GL_COMPRESSED_RGB8_ETC2, GLenum_GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_COMPRESSED_RGBA8_ETC2_EAC, GLenum_GL_COMPRESSED_SIGNED_R11_EAC, GLenum_GL_COMPRESSED_SIGNED_RG11_EAC, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC, GLenum_GL_COMPRESSED_SRGB8_ETC2, GLenum_GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_DEPTH24_STENCIL8, GLenum_GL_DEPTH32F_STENCIL8, GLenum_GL_DEPTH_COMPONENT16, GLenum_GL_DEPTH_COMPONENT24, GLenum_GL_DEPTH_COMPONENT32F, GLenum_GL_R11F_G11F_B10F, GLenum_GL_R16F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32F, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_R8_SNORM, GLenum_GL_RG16F, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32F, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RG8_SNORM, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB16F, GLenum_GL_RGB16I, GLenum_GL_RGB16UI, GLenum_GL_RGB32F, GLenum_GL_RGB32I, GLenum_GL_RGB32UI, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGB8, GLenum_GL_RGB8I, GLenum_GL_RGB8UI, GLenum_GL_RGB8_SNORM, GLenum_GL_RGB9_E5, GLenum_GL_RGBA16F, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32F, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA4, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_RGBA8_SNORM, GLenum_GL_SRGB8, GLenum_GL_SRGB8_ALPHA8:
	default:
		v := ϟa.Internalformat
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1032_major, minRequiredVersion_1032_minor
	return nil
}
func (ϟa *GlTexSubImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1035_major := uint32(2) // u32
	minRequiredVersion_1035_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Format {
	case GLenum_GL_ALPHA, GLenum_GL_LUMINANCE, GLenum_GL_LUMINANCE_ALPHA, GLenum_GL_RGB, GLenum_GL_RGBA:
	case GLenum_GL_DEPTH_COMPONENT, GLenum_GL_DEPTH_STENCIL, GLenum_GL_RED, GLenum_GL_RED_INTEGER, GLenum_GL_RG, GLenum_GL_RGBA_INTEGER, GLenum_GL_RGB_INTEGER, GLenum_GL_RG_INTEGER:
		minRequiredVersion_1037_major := uint32(3) // u32
		minRequiredVersion_1037_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1037_major, minRequiredVersion_1037_minor
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Type {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_SHORT_4_4_4_4, GLenum_GL_UNSIGNED_SHORT_5_5_5_1, GLenum_GL_UNSIGNED_SHORT_5_6_5:
	case GLenum_GL_HALF_FLOAT_OES:
		requiresExtension_1039_ext := ExtensionId_GL_OES_texture_half_float // ExtensionId
		_ = requiresExtension_1039_ext
	case GLenum_GL_BYTE, GLenum_GL_FLOAT, GLenum_GL_FLOAT_32_UNSIGNED_INT_24_8_REV, GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_10F_11F_11F_REV, GLenum_GL_UNSIGNED_INT_24_8, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_INT_5_9_9_9_REV, GLenum_GL_UNSIGNED_SHORT:
		minRequiredVersion_1040_major := uint32(3) // u32
		minRequiredVersion_1040_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1040_major, minRequiredVersion_1040_minor
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_1042_result := context            // Contextʳ
	ctx := GetContext_1042_result                // Contextʳ
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D:
		id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(GLenum_GL_TEXTURE_2D) // TextureId
		t := ctx.Instances.Textures.Get(id)                                         // Textureʳ
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.Width
			s.Height = ϟa.Height
			s.Size = externs{ϟa, ϟs, ϟd, ϟl, ϟb}.imageSize(uint32(ϟa.Width), uint32(ϟa.Height), ϟa.Format, ϟa.Type)
			s.Format = ϟa.Format
			return s
		}() // Image
		if ((ctx.BoundBuffers.Get(GLenum_GL_PIXEL_UNPACK_BUFFER)) == (BufferId(uint32(0)))) && ((ϟa.Data) != (TexturePointer(Voidᶜᵖ{}))) {
			l.Data = U8ᵖ(ϟa.Data).Slice(uint64(uint32(0)), uint64(l.Size), ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
		}
		t.Texture2D[ϟa.Level] = l
		t.Kind = TextureKind_TEXTURE2D
		t.Format = ϟa.Format
		_, _, _ = id, t, l
	case GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z:
		id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(GLenum_GL_TEXTURE_CUBE_MAP) // TextureId
		t := ctx.Instances.Textures.Get(id)                                               // Textureʳ
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.Width
			s.Height = ϟa.Height
			s.Size = externs{ϟa, ϟs, ϟd, ϟl, ϟb}.imageSize(uint32(ϟa.Width), uint32(ϟa.Height), ϟa.Format, ϟa.Type)
			s.Format = ϟa.Format
			return s
		}() // Image
		if ((ctx.BoundBuffers.Get(GLenum_GL_PIXEL_UNPACK_BUFFER)) == (BufferId(uint32(0)))) && ((ϟa.Data) != (TexturePointer(Voidᶜᵖ{}))) {
			l.Data = U8ᵖ(ϟa.Data).Slice(uint64(uint32(0)), uint64(l.Size), ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
		}
		cube := t.Cubemap.Get(ϟa.Level) // CubemapLevel
		cube.Faces[ϟa.Target] = l
		t.Cubemap[ϟa.Level] = cube
		t.Kind = TextureKind_CUBEMAP
		t.Format = ϟa.Format
		_, _, _, _ = id, t, l, cube
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1035_major, minRequiredVersion_1035_minor, context, GetContext_1042_result, ctx
	return nil
}
func (ϟa *GlTexSubImage3D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1043_major := uint32(3) // u32
	minRequiredVersion_1043_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Format {
	case GLenum_GL_ALPHA, GLenum_GL_DEPTH_COMPONENT, GLenum_GL_DEPTH_STENCIL, GLenum_GL_LUMINANCE, GLenum_GL_LUMINANCE_ALPHA, GLenum_GL_RED, GLenum_GL_RED_INTEGER, GLenum_GL_RG, GLenum_GL_RGB, GLenum_GL_RGBA, GLenum_GL_RGBA_INTEGER, GLenum_GL_RGB_INTEGER, GLenum_GL_RG_INTEGER:
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Type {
	case GLenum_GL_HALF_FLOAT_OES:
		requiresExtension_1046_ext := ExtensionId_GL_OES_texture_half_float // ExtensionId
		_ = requiresExtension_1046_ext
	case GLenum_GL_BYTE, GLenum_GL_FLOAT, GLenum_GL_FLOAT_32_UNSIGNED_INT_24_8_REV, GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_10F_11F_11F_REV, GLenum_GL_UNSIGNED_INT_24_8, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_INT_5_9_9_9_REV, GLenum_GL_UNSIGNED_SHORT, GLenum_GL_UNSIGNED_SHORT_4_4_4_4, GLenum_GL_UNSIGNED_SHORT_5_5_5_1, GLenum_GL_UNSIGNED_SHORT_5_6_5:
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1043_major, minRequiredVersion_1043_minor
	return nil
}
func (ϟa *GlBeginTransformFeedback) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1048_major := uint32(3) // u32
	minRequiredVersion_1048_minor := uint32(0) // u32
	switch ϟa.PrimitiveMode {
	case GLenum_GL_LINES, GLenum_GL_POINTS, GLenum_GL_TRIANGLES:
	default:
		v := ϟa.PrimitiveMode
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1048_major, minRequiredVersion_1048_minor
	return nil
}
func (ϟa *GlBindTransformFeedback) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1050_major := uint32(3) // u32
	minRequiredVersion_1050_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TRANSFORM_FEEDBACK:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1050_major, minRequiredVersion_1050_minor
	return nil
}
func (ϟa *GlDeleteTransformFeedbacks) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1052_major := uint32(3) // u32
	minRequiredVersion_1052_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1052_major, minRequiredVersion_1052_minor
	return nil
}
func (ϟa *GlEndTransformFeedback) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1053_major := uint32(3) // u32
	minRequiredVersion_1053_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1053_major, minRequiredVersion_1053_minor
	return nil
}
func (ϟa *GlGenTransformFeedbacks) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1054_major := uint32(3) // u32
	minRequiredVersion_1054_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1054_major, minRequiredVersion_1054_minor
	return nil
}
func (ϟa *GlGetTransformFeedbackVarying) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1055_major := uint32(3) // u32
	minRequiredVersion_1055_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1055_major, minRequiredVersion_1055_minor
	return nil
}
func (ϟa *GlIsTransformFeedback) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1056_major := uint32(3) // u32
	minRequiredVersion_1056_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_1056_major, minRequiredVersion_1056_minor
	return nil
}
func (ϟa *GlPauseTransformFeedback) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1057_major := uint32(3) // u32
	minRequiredVersion_1057_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1057_major, minRequiredVersion_1057_minor
	return nil
}
func (ϟa *GlResumeTransformFeedback) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1058_major := uint32(3) // u32
	minRequiredVersion_1058_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1058_major, minRequiredVersion_1058_minor
	return nil
}
func (ϟa *GlTransformFeedbackVaryings) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1059_major := uint32(3) // u32
	minRequiredVersion_1059_minor := uint32(0) // u32
	switch ϟa.BufferMode {
	case GLenum_GL_INTERLEAVED_ATTRIBS, GLenum_GL_SEPARATE_ATTRIBS:
	default:
		v := ϟa.BufferMode
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1059_major, minRequiredVersion_1059_minor
	return nil
}
func (ϟa *GlBindVertexArray) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1061_major := uint32(3)   // u32
	minRequiredVersion_1061_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_1062_result := context            // Contextʳ
	ctx := GetContext_1062_result                // Contextʳ
	if !(ctx.Instances.VertexArrays.Contains(ϟa.Array)) {
		ctx.Instances.VertexArrays[ϟa.Array] = func() *VertexArray {
			s := &VertexArray{}
			s.Init()
			return s
		}()
	}
	ctx.BoundVertexArray = ϟa.Array
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1061_major, minRequiredVersion_1061_minor, context, GetContext_1062_result, ctx
	return nil
}
func (ϟa *GlBindVertexBuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1063_major := uint32(3) // u32
	minRequiredVersion_1063_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1063_major, minRequiredVersion_1063_minor
	return nil
}
func (ϟa *GlDeleteVertexArrays) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1064_major := uint32(3)                            // u32
	minRequiredVersion_1064_minor := uint32(0)                            // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                          // Contextʳ
	GetContext_1065_result := context                                     // Contextʳ
	ctx := GetContext_1065_result                                         // Contextʳ
	a := ϟa.Arrays.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // VertexArrayIdˢ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.VertexArrays, a.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_1064_major, minRequiredVersion_1064_minor, context, GetContext_1065_result, ctx, a
	return nil
}
func (ϟa *GlDisableVertexAttribArray) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1066_major := uint32(2)   // u32
	minRequiredVersion_1066_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_1067_result := context            // Contextʳ
	ctx := GetContext_1067_result                // Contextʳ
	ctx.VertexAttributeArrays.Get(ϟa.Location).Enabled = false
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1066_major, minRequiredVersion_1066_minor, context, GetContext_1067_result, ctx
	return nil
}
func (ϟa *GlEnableVertexAttribArray) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1068_major := uint32(2)   // u32
	minRequiredVersion_1068_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_1069_result := context            // Contextʳ
	ctx := GetContext_1069_result                // Contextʳ
	ctx.VertexAttributeArrays.Get(ϟa.Location).Enabled = true
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1068_major, minRequiredVersion_1068_minor, context, GetContext_1069_result, ctx
	return nil
}
func (ϟa *GlGenVertexArrays) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1070_major := uint32(3)                            // u32
	minRequiredVersion_1070_minor := uint32(0)                            // u32
	a := ϟa.Arrays.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // VertexArrayIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                          // Contextʳ
	GetContext_1071_result := context                                     // Contextʳ
	ctx := GetContext_1071_result                                         // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := VertexArrayId(ϟa.Arrays.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil)) // VertexArrayId
		ctx.Instances.VertexArrays[id] = func() *VertexArray {
			s := &VertexArray{}
			s.Init()
			return s
		}()
		a.Index(uint64(i), ϟs).Write(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = minRequiredVersion_1070_major, minRequiredVersion_1070_minor, a, context, GetContext_1071_result, ctx
	return nil
}
func (ϟa *GlGetVertexAttribIiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1072_major := uint32(3) // u32
	minRequiredVersion_1072_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_CURRENT_VERTEX_ATTRIB, GLenum_GL_VERTEX_ATTRIB_ARRAY_BUFFER_BINDING, GLenum_GL_VERTEX_ATTRIB_ARRAY_DIVISOR, GLenum_GL_VERTEX_ATTRIB_ARRAY_ENABLED, GLenum_GL_VERTEX_ATTRIB_ARRAY_INTEGER, GLenum_GL_VERTEX_ATTRIB_ARRAY_NORMALIZED, GLenum_GL_VERTEX_ATTRIB_ARRAY_SIZE, GLenum_GL_VERTEX_ATTRIB_ARRAY_STRIDE, GLenum_GL_VERTEX_ATTRIB_ARRAY_TYPE:
	case GLenum_GL_VERTEX_ATTRIB_BINDING:
		minRequiredVersion_1073_major := uint32(3) // u32
		minRequiredVersion_1073_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1073_major, minRequiredVersion_1073_minor
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1072_major, minRequiredVersion_1072_minor
	return nil
}
func (ϟa *GlGetVertexAttribIuiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1075_major := uint32(3) // u32
	minRequiredVersion_1075_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_CURRENT_VERTEX_ATTRIB, GLenum_GL_VERTEX_ATTRIB_ARRAY_BUFFER_BINDING, GLenum_GL_VERTEX_ATTRIB_ARRAY_DIVISOR, GLenum_GL_VERTEX_ATTRIB_ARRAY_ENABLED, GLenum_GL_VERTEX_ATTRIB_ARRAY_INTEGER, GLenum_GL_VERTEX_ATTRIB_ARRAY_NORMALIZED, GLenum_GL_VERTEX_ATTRIB_ARRAY_SIZE, GLenum_GL_VERTEX_ATTRIB_ARRAY_STRIDE, GLenum_GL_VERTEX_ATTRIB_ARRAY_TYPE:
	case GLenum_GL_VERTEX_ATTRIB_BINDING:
		minRequiredVersion_1076_major := uint32(3) // u32
		minRequiredVersion_1076_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1076_major, minRequiredVersion_1076_minor
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1075_major, minRequiredVersion_1075_minor
	return nil
}
func (ϟa *GlGetVertexAttribPointerv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1078_major := uint32(2) // u32
	minRequiredVersion_1078_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_VERTEX_ATTRIB_ARRAY_POINTER:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1078_major, minRequiredVersion_1078_minor
	return nil
}
func (ϟa *GlGetVertexAttribfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1080_major := uint32(2) // u32
	minRequiredVersion_1080_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_CURRENT_VERTEX_ATTRIB, GLenum_GL_VERTEX_ATTRIB_ARRAY_BUFFER_BINDING, GLenum_GL_VERTEX_ATTRIB_ARRAY_ENABLED, GLenum_GL_VERTEX_ATTRIB_ARRAY_NORMALIZED, GLenum_GL_VERTEX_ATTRIB_ARRAY_SIZE, GLenum_GL_VERTEX_ATTRIB_ARRAY_STRIDE, GLenum_GL_VERTEX_ATTRIB_ARRAY_TYPE:
	case GLenum_GL_VERTEX_ATTRIB_ARRAY_DIVISOR, GLenum_GL_VERTEX_ATTRIB_ARRAY_INTEGER:
		minRequiredVersion_1081_major := uint32(3) // u32
		minRequiredVersion_1081_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1081_major, minRequiredVersion_1081_minor
	case GLenum_GL_VERTEX_ATTRIB_BINDING:
		minRequiredVersion_1082_major := uint32(3) // u32
		minRequiredVersion_1082_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1082_major, minRequiredVersion_1082_minor
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1080_major, minRequiredVersion_1080_minor
	return nil
}
func (ϟa *GlGetVertexAttribiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1084_major := uint32(2) // u32
	minRequiredVersion_1084_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_CURRENT_VERTEX_ATTRIB, GLenum_GL_VERTEX_ATTRIB_ARRAY_BUFFER_BINDING, GLenum_GL_VERTEX_ATTRIB_ARRAY_ENABLED, GLenum_GL_VERTEX_ATTRIB_ARRAY_NORMALIZED, GLenum_GL_VERTEX_ATTRIB_ARRAY_SIZE, GLenum_GL_VERTEX_ATTRIB_ARRAY_STRIDE, GLenum_GL_VERTEX_ATTRIB_ARRAY_TYPE:
	case GLenum_GL_VERTEX_ATTRIB_ARRAY_DIVISOR, GLenum_GL_VERTEX_ATTRIB_ARRAY_INTEGER:
		minRequiredVersion_1085_major := uint32(3) // u32
		minRequiredVersion_1085_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1085_major, minRequiredVersion_1085_minor
	case GLenum_GL_VERTEX_ATTRIB_BINDING:
		minRequiredVersion_1086_major := uint32(3) // u32
		minRequiredVersion_1086_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1086_major, minRequiredVersion_1086_minor
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1084_major, minRequiredVersion_1084_minor
	return nil
}
func (ϟa *GlIsVertexArray) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1088_major := uint32(3) // u32
	minRequiredVersion_1088_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_1088_major, minRequiredVersion_1088_minor
	return nil
}
func (ϟa *GlVertexAttrib1f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1089_major := uint32(2) // u32
	minRequiredVersion_1089_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1089_major, minRequiredVersion_1089_minor
	return nil
}
func (ϟa *GlVertexAttrib1fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1090_major := uint32(2) // u32
	minRequiredVersion_1090_minor := uint32(0) // u32
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1090_major, minRequiredVersion_1090_minor
	return nil
}
func (ϟa *GlVertexAttrib2f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1091_major := uint32(2) // u32
	minRequiredVersion_1091_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1091_major, minRequiredVersion_1091_minor
	return nil
}
func (ϟa *GlVertexAttrib2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1092_major := uint32(2) // u32
	minRequiredVersion_1092_minor := uint32(0) // u32
	ϟa.Value.Slice(uint64(0), uint64(2), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1092_major, minRequiredVersion_1092_minor
	return nil
}
func (ϟa *GlVertexAttrib3f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1093_major := uint32(2) // u32
	minRequiredVersion_1093_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1093_major, minRequiredVersion_1093_minor
	return nil
}
func (ϟa *GlVertexAttrib3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1094_major := uint32(2) // u32
	minRequiredVersion_1094_minor := uint32(0) // u32
	ϟa.Value.Slice(uint64(0), uint64(3), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1094_major, minRequiredVersion_1094_minor
	return nil
}
func (ϟa *GlVertexAttrib4f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1095_major := uint32(2) // u32
	minRequiredVersion_1095_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1095_major, minRequiredVersion_1095_minor
	return nil
}
func (ϟa *GlVertexAttrib4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1096_major := uint32(2) // u32
	minRequiredVersion_1096_minor := uint32(0) // u32
	ϟa.Value.Slice(uint64(0), uint64(4), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1096_major, minRequiredVersion_1096_minor
	return nil
}
func (ϟa *GlVertexAttribBinding) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1097_major := uint32(3) // u32
	minRequiredVersion_1097_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1097_major, minRequiredVersion_1097_minor
	return nil
}
func (ϟa *GlVertexAttribDivisor) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1098_major := uint32(3) // u32
	minRequiredVersion_1098_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1098_major, minRequiredVersion_1098_minor
	return nil
}
func (ϟa *GlVertexAttribFormat) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1099_major := uint32(3) // u32
	minRequiredVersion_1099_minor := uint32(1) // u32
	switch ϟa.Type {
	case GLenum_GL_BYTE, GLenum_GL_FIXED, GLenum_GL_FLOAT, GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_INT_2_10_10_10_REV, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_SHORT:
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1099_major, minRequiredVersion_1099_minor
	return nil
}
func (ϟa *GlVertexAttribI4i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1101_major := uint32(3) // u32
	minRequiredVersion_1101_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1101_major, minRequiredVersion_1101_minor
	return nil
}
func (ϟa *GlVertexAttribI4iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1102_major := uint32(3) // u32
	minRequiredVersion_1102_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1102_major, minRequiredVersion_1102_minor
	return nil
}
func (ϟa *GlVertexAttribI4ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1103_major := uint32(3) // u32
	minRequiredVersion_1103_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1103_major, minRequiredVersion_1103_minor
	return nil
}
func (ϟa *GlVertexAttribI4uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1104_major := uint32(3) // u32
	minRequiredVersion_1104_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1104_major, minRequiredVersion_1104_minor
	return nil
}
func (ϟa *GlVertexAttribIFormat) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1105_major := uint32(3) // u32
	minRequiredVersion_1105_minor := uint32(1) // u32
	switch ϟa.Type {
	case GLenum_GL_BYTE, GLenum_GL_FIXED, GLenum_GL_FLOAT, GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_INT_2_10_10_10_REV, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_SHORT:
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1105_major, minRequiredVersion_1105_minor
	return nil
}
func (ϟa *GlVertexAttribIPointer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1107_major := uint32(3) // u32
	minRequiredVersion_1107_minor := uint32(0) // u32
	switch ϟa.Type {
	case GLenum_GL_HALF_FLOAT_OES:
		requiresExtension_1108_ext := ExtensionId_GL_OES_vertex_half_float // ExtensionId
		_ = requiresExtension_1108_ext
	case GLenum_GL_BYTE, GLenum_GL_INT, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_SHORT:
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1107_major, minRequiredVersion_1107_minor
	return nil
}
func (ϟa *GlVertexAttribPointer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1110_major := uint32(2) // u32
	minRequiredVersion_1110_minor := uint32(0) // u32
	switch ϟa.Type {
	case GLenum_GL_BYTE, GLenum_GL_FIXED, GLenum_GL_FLOAT, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_SHORT:
	case GLenum_GL_HALF_FLOAT_OES:
		requiresExtension_1111_ext := ExtensionId_GL_OES_vertex_half_float // ExtensionId
		_ = requiresExtension_1111_ext
	case GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV:
		minRequiredVersion_1112_major := uint32(3) // u32
		minRequiredVersion_1112_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1112_major, minRequiredVersion_1112_minor
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)    // Contextʳ
	GetContext_1114_result := context               // Contextʳ
	ctx := GetContext_1114_result                   // Contextʳ
	a := ctx.VertexAttributeArrays.Get(ϟa.Location) // VertexAttributeArrayʳ
	a.Size = uint32(ϟa.Size)
	a.Type = ϟa.Type
	a.Normalized = ϟa.Normalized
	a.Stride = ϟa.Stride
	a.Pointer = ϟa.Data
	a.Buffer = ctx.BoundBuffers.Get(GLenum_GL_ARRAY_BUFFER)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_1110_major, minRequiredVersion_1110_minor, context, GetContext_1114_result, ctx, a
	return nil
}
func (ϟa *GlVertexBindingDivisor) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1115_major := uint32(3) // u32
	minRequiredVersion_1115_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1115_major, minRequiredVersion_1115_minor
	return nil
}
func (ϟa *EglInitialize) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	if (ϟa.Major) != (EGLintᵖ{}) {
		ϟa.Major.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Major.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	}
	if (ϟa.Major) != (EGLintᵖ{}) {
		ϟa.Minor.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Minor.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	}
	ϟa.Result = ϟa.Result
	return nil
}
func (ϟa *EglCreateContext) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	context := EGLContext(ϟa.Result) // EGLContext
	identifier := ϟc.NextContextID   // ContextID
	ϟc.NextContextID = (ϟc.NextContextID) + (ContextID(uint32(1)))
	ctx := func() *Context {
		s := &Context{}
		s.Init()
		return s
	}() // Contextʳ
	ctx.Identifier = identifier
	ctx.Instances.Buffers[BufferId(uint32(0))] = func() *Buffer {
		s := &Buffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Textures[TextureId(uint32(0))] = func() *Texture {
		s := &Texture{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[RenderbufferId(uint32(0))] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	color_id := RenderbufferId(uint32(4294967295))   // RenderbufferId
	depth_id := RenderbufferId(uint32(4294967294))   // RenderbufferId
	stencil_id := RenderbufferId(uint32(4294967293)) // RenderbufferId
	ctx.Instances.Renderbuffers[color_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[depth_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[stencil_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	backbuffer := func() *Framebuffer {
		s := &Framebuffer{}
		s.Init()
		return s
	}() // Framebufferʳ
	backbuffer.Attachments[GLenum_GL_COLOR_ATTACHMENT0] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(color_id)
		s.Type = GLenum_GL_RENDERBUFFER
		s.CubeMapFace = GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	backbuffer.Attachments[GLenum_GL_DEPTH_ATTACHMENT] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(depth_id)
		s.Type = GLenum_GL_RENDERBUFFER
		s.CubeMapFace = GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	backbuffer.Attachments[GLenum_GL_STENCIL_ATTACHMENT] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(stencil_id)
		s.Type = GLenum_GL_RENDERBUFFER
		s.CubeMapFace = GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	ctx.Instances.Framebuffers[FramebufferId(uint32(0))] = backbuffer
	ctx.BoundFramebuffers[GLenum_GL_DRAW_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.BoundFramebuffers[GLenum_GL_READ_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.Rasterizing.StencilMask[GLenum_GL_FRONT] = GLuint(uint32(4294967295))
	ctx.Rasterizing.StencilMask[GLenum_GL_BACK] = GLuint(uint32(4294967295))
	ctx.PixelStorage[GLenum_GL_PACK_ALIGNMENT] = GLint(int32(4))
	ctx.PixelStorage[GLenum_GL_UNPACK_ALIGNMENT] = GLint(int32(4))
	for i := int32(int32(0)); i < int32(64); i++ {
		ctx.VertexAttributeArrays[AttributeLocation(i)] = func() *VertexAttributeArray {
			s := &VertexAttributeArray{}
			s.Init()
			return s
		}()
	}
	CreateContext_1116_result := ctx // Contextʳ
	ϟc.EGLContexts[context] = CreateContext_1116_result
	ϟa.Result = context
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_1116_result
	return nil
}
func (ϟa *EglMakeCurrent) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	SetContext_1117_context := ϟc.EGLContexts.Get(ϟa.Context) // Contextʳ
	ϟc.Contexts[ϟc.CurrentThread] = SetContext_1117_context
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = SetContext_1117_context
	return nil
}
func (ϟa *EglSwapBuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	return nil
}
func (ϟa *EglQuerySurface) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Result = ϟa.Result
	return nil
}
func (ϟa *GlXCreateContext) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	context := GLXContext(ϟa.Result) // GLXContext
	identifier := ϟc.NextContextID   // ContextID
	ϟc.NextContextID = (ϟc.NextContextID) + (ContextID(uint32(1)))
	ctx := func() *Context {
		s := &Context{}
		s.Init()
		return s
	}() // Contextʳ
	ctx.Identifier = identifier
	ctx.Instances.Buffers[BufferId(uint32(0))] = func() *Buffer {
		s := &Buffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Textures[TextureId(uint32(0))] = func() *Texture {
		s := &Texture{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[RenderbufferId(uint32(0))] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	color_id := RenderbufferId(uint32(4294967295))   // RenderbufferId
	depth_id := RenderbufferId(uint32(4294967294))   // RenderbufferId
	stencil_id := RenderbufferId(uint32(4294967293)) // RenderbufferId
	ctx.Instances.Renderbuffers[color_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[depth_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[stencil_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	backbuffer := func() *Framebuffer {
		s := &Framebuffer{}
		s.Init()
		return s
	}() // Framebufferʳ
	backbuffer.Attachments[GLenum_GL_COLOR_ATTACHMENT0] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(color_id)
		s.Type = GLenum_GL_RENDERBUFFER
		s.CubeMapFace = GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	backbuffer.Attachments[GLenum_GL_DEPTH_ATTACHMENT] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(depth_id)
		s.Type = GLenum_GL_RENDERBUFFER
		s.CubeMapFace = GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	backbuffer.Attachments[GLenum_GL_STENCIL_ATTACHMENT] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(stencil_id)
		s.Type = GLenum_GL_RENDERBUFFER
		s.CubeMapFace = GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	ctx.Instances.Framebuffers[FramebufferId(uint32(0))] = backbuffer
	ctx.BoundFramebuffers[GLenum_GL_DRAW_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.BoundFramebuffers[GLenum_GL_READ_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.Rasterizing.StencilMask[GLenum_GL_FRONT] = GLuint(uint32(4294967295))
	ctx.Rasterizing.StencilMask[GLenum_GL_BACK] = GLuint(uint32(4294967295))
	ctx.PixelStorage[GLenum_GL_PACK_ALIGNMENT] = GLint(int32(4))
	ctx.PixelStorage[GLenum_GL_UNPACK_ALIGNMENT] = GLint(int32(4))
	for i := int32(int32(0)); i < int32(64); i++ {
		ctx.VertexAttributeArrays[AttributeLocation(i)] = func() *VertexAttributeArray {
			s := &VertexAttributeArray{}
			s.Init()
			return s
		}()
	}
	CreateContext_1118_result := ctx // Contextʳ
	ϟc.GLXContexts[context] = CreateContext_1118_result
	ϟa.Result = context
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_1118_result
	return nil
}
func (ϟa *GlXCreateNewContext) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	context := GLXContext(ϟa.Result) // GLXContext
	identifier := ϟc.NextContextID   // ContextID
	ϟc.NextContextID = (ϟc.NextContextID) + (ContextID(uint32(1)))
	ctx := func() *Context {
		s := &Context{}
		s.Init()
		return s
	}() // Contextʳ
	ctx.Identifier = identifier
	ctx.Instances.Buffers[BufferId(uint32(0))] = func() *Buffer {
		s := &Buffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Textures[TextureId(uint32(0))] = func() *Texture {
		s := &Texture{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[RenderbufferId(uint32(0))] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	color_id := RenderbufferId(uint32(4294967295))   // RenderbufferId
	depth_id := RenderbufferId(uint32(4294967294))   // RenderbufferId
	stencil_id := RenderbufferId(uint32(4294967293)) // RenderbufferId
	ctx.Instances.Renderbuffers[color_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[depth_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[stencil_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	backbuffer := func() *Framebuffer {
		s := &Framebuffer{}
		s.Init()
		return s
	}() // Framebufferʳ
	backbuffer.Attachments[GLenum_GL_COLOR_ATTACHMENT0] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(color_id)
		s.Type = GLenum_GL_RENDERBUFFER
		s.CubeMapFace = GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	backbuffer.Attachments[GLenum_GL_DEPTH_ATTACHMENT] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(depth_id)
		s.Type = GLenum_GL_RENDERBUFFER
		s.CubeMapFace = GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	backbuffer.Attachments[GLenum_GL_STENCIL_ATTACHMENT] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(stencil_id)
		s.Type = GLenum_GL_RENDERBUFFER
		s.CubeMapFace = GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	ctx.Instances.Framebuffers[FramebufferId(uint32(0))] = backbuffer
	ctx.BoundFramebuffers[GLenum_GL_DRAW_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.BoundFramebuffers[GLenum_GL_READ_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.Rasterizing.StencilMask[GLenum_GL_FRONT] = GLuint(uint32(4294967295))
	ctx.Rasterizing.StencilMask[GLenum_GL_BACK] = GLuint(uint32(4294967295))
	ctx.PixelStorage[GLenum_GL_PACK_ALIGNMENT] = GLint(int32(4))
	ctx.PixelStorage[GLenum_GL_UNPACK_ALIGNMENT] = GLint(int32(4))
	for i := int32(int32(0)); i < int32(64); i++ {
		ctx.VertexAttributeArrays[AttributeLocation(i)] = func() *VertexAttributeArray {
			s := &VertexAttributeArray{}
			s.Init()
			return s
		}()
	}
	CreateContext_1119_result := ctx // Contextʳ
	ϟc.GLXContexts[context] = CreateContext_1119_result
	ϟa.Result = context
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_1119_result
	return nil
}
func (ϟa *GlXMakeContextCurrent) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	SetContext_1120_context := ϟc.GLXContexts.Get(ϟa.Ctx) // Contextʳ
	ϟc.Contexts[ϟc.CurrentThread] = SetContext_1120_context
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = SetContext_1120_context
	return nil
}
func (ϟa *GlXMakeCurrent) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	SetContext_1121_context := ϟc.GLXContexts.Get(ϟa.Ctx) // Contextʳ
	ϟc.Contexts[ϟc.CurrentThread] = SetContext_1121_context
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = SetContext_1121_context
	return nil
}
func (ϟa *GlXSwapBuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *GlXQueryDrawable) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Result = ϟa.Result
	return nil
}
func (ϟa *WglCreateContext) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	context := HGLRC(ϟa.Result)    // HGLRC
	identifier := ϟc.NextContextID // ContextID
	ϟc.NextContextID = (ϟc.NextContextID) + (ContextID(uint32(1)))
	ctx := func() *Context {
		s := &Context{}
		s.Init()
		return s
	}() // Contextʳ
	ctx.Identifier = identifier
	ctx.Instances.Buffers[BufferId(uint32(0))] = func() *Buffer {
		s := &Buffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Textures[TextureId(uint32(0))] = func() *Texture {
		s := &Texture{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[RenderbufferId(uint32(0))] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	color_id := RenderbufferId(uint32(4294967295))   // RenderbufferId
	depth_id := RenderbufferId(uint32(4294967294))   // RenderbufferId
	stencil_id := RenderbufferId(uint32(4294967293)) // RenderbufferId
	ctx.Instances.Renderbuffers[color_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[depth_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[stencil_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	backbuffer := func() *Framebuffer {
		s := &Framebuffer{}
		s.Init()
		return s
	}() // Framebufferʳ
	backbuffer.Attachments[GLenum_GL_COLOR_ATTACHMENT0] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(color_id)
		s.Type = GLenum_GL_RENDERBUFFER
		s.CubeMapFace = GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	backbuffer.Attachments[GLenum_GL_DEPTH_ATTACHMENT] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(depth_id)
		s.Type = GLenum_GL_RENDERBUFFER
		s.CubeMapFace = GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	backbuffer.Attachments[GLenum_GL_STENCIL_ATTACHMENT] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(stencil_id)
		s.Type = GLenum_GL_RENDERBUFFER
		s.CubeMapFace = GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	ctx.Instances.Framebuffers[FramebufferId(uint32(0))] = backbuffer
	ctx.BoundFramebuffers[GLenum_GL_DRAW_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.BoundFramebuffers[GLenum_GL_READ_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.Rasterizing.StencilMask[GLenum_GL_FRONT] = GLuint(uint32(4294967295))
	ctx.Rasterizing.StencilMask[GLenum_GL_BACK] = GLuint(uint32(4294967295))
	ctx.PixelStorage[GLenum_GL_PACK_ALIGNMENT] = GLint(int32(4))
	ctx.PixelStorage[GLenum_GL_UNPACK_ALIGNMENT] = GLint(int32(4))
	for i := int32(int32(0)); i < int32(64); i++ {
		ctx.VertexAttributeArrays[AttributeLocation(i)] = func() *VertexAttributeArray {
			s := &VertexAttributeArray{}
			s.Init()
			return s
		}()
	}
	CreateContext_1122_result := ctx // Contextʳ
	ϟc.WGLContexts[context] = CreateContext_1122_result
	ϟa.Result = context
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_1122_result
	return nil
}
func (ϟa *WglCreateContextAttribsARB) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	context := HGLRC(ϟa.Result)    // HGLRC
	identifier := ϟc.NextContextID // ContextID
	ϟc.NextContextID = (ϟc.NextContextID) + (ContextID(uint32(1)))
	ctx := func() *Context {
		s := &Context{}
		s.Init()
		return s
	}() // Contextʳ
	ctx.Identifier = identifier
	ctx.Instances.Buffers[BufferId(uint32(0))] = func() *Buffer {
		s := &Buffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Textures[TextureId(uint32(0))] = func() *Texture {
		s := &Texture{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[RenderbufferId(uint32(0))] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	color_id := RenderbufferId(uint32(4294967295))   // RenderbufferId
	depth_id := RenderbufferId(uint32(4294967294))   // RenderbufferId
	stencil_id := RenderbufferId(uint32(4294967293)) // RenderbufferId
	ctx.Instances.Renderbuffers[color_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[depth_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[stencil_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	backbuffer := func() *Framebuffer {
		s := &Framebuffer{}
		s.Init()
		return s
	}() // Framebufferʳ
	backbuffer.Attachments[GLenum_GL_COLOR_ATTACHMENT0] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(color_id)
		s.Type = GLenum_GL_RENDERBUFFER
		s.CubeMapFace = GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	backbuffer.Attachments[GLenum_GL_DEPTH_ATTACHMENT] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(depth_id)
		s.Type = GLenum_GL_RENDERBUFFER
		s.CubeMapFace = GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	backbuffer.Attachments[GLenum_GL_STENCIL_ATTACHMENT] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(stencil_id)
		s.Type = GLenum_GL_RENDERBUFFER
		s.CubeMapFace = GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	ctx.Instances.Framebuffers[FramebufferId(uint32(0))] = backbuffer
	ctx.BoundFramebuffers[GLenum_GL_DRAW_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.BoundFramebuffers[GLenum_GL_READ_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.Rasterizing.StencilMask[GLenum_GL_FRONT] = GLuint(uint32(4294967295))
	ctx.Rasterizing.StencilMask[GLenum_GL_BACK] = GLuint(uint32(4294967295))
	ctx.PixelStorage[GLenum_GL_PACK_ALIGNMENT] = GLint(int32(4))
	ctx.PixelStorage[GLenum_GL_UNPACK_ALIGNMENT] = GLint(int32(4))
	for i := int32(int32(0)); i < int32(64); i++ {
		ctx.VertexAttributeArrays[AttributeLocation(i)] = func() *VertexAttributeArray {
			s := &VertexAttributeArray{}
			s.Init()
			return s
		}()
	}
	CreateContext_1123_result := ctx // Contextʳ
	ϟc.WGLContexts[context] = CreateContext_1123_result
	ϟa.Result = context
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_1123_result
	return nil
}
func (ϟa *WglMakeCurrent) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	SetContext_1124_context := ϟc.WGLContexts.Get(ϟa.Hglrc) // Contextʳ
	ϟc.Contexts[ϟc.CurrentThread] = SetContext_1124_context
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = SetContext_1124_context
	return nil
}
func (ϟa *WglSwapBuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *CGLCreateContext) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	context := CGLContextObj(ϟa.Ctx.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil)) // CGLContextObj
	identifier := ϟc.NextContextID                                                                                                                // ContextID
	ϟc.NextContextID = (ϟc.NextContextID) + (ContextID(uint32(1)))
	ctx := func() *Context {
		s := &Context{}
		s.Init()
		return s
	}() // Contextʳ
	ctx.Identifier = identifier
	ctx.Instances.Buffers[BufferId(uint32(0))] = func() *Buffer {
		s := &Buffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Textures[TextureId(uint32(0))] = func() *Texture {
		s := &Texture{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[RenderbufferId(uint32(0))] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	color_id := RenderbufferId(uint32(4294967295))   // RenderbufferId
	depth_id := RenderbufferId(uint32(4294967294))   // RenderbufferId
	stencil_id := RenderbufferId(uint32(4294967293)) // RenderbufferId
	ctx.Instances.Renderbuffers[color_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[depth_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	ctx.Instances.Renderbuffers[stencil_id] = func() *Renderbuffer {
		s := &Renderbuffer{}
		s.Init()
		return s
	}()
	backbuffer := func() *Framebuffer {
		s := &Framebuffer{}
		s.Init()
		return s
	}() // Framebufferʳ
	backbuffer.Attachments[GLenum_GL_COLOR_ATTACHMENT0] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(color_id)
		s.Type = GLenum_GL_RENDERBUFFER
		s.CubeMapFace = GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	backbuffer.Attachments[GLenum_GL_DEPTH_ATTACHMENT] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(depth_id)
		s.Type = GLenum_GL_RENDERBUFFER
		s.CubeMapFace = GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	backbuffer.Attachments[GLenum_GL_STENCIL_ATTACHMENT] = func() FramebufferAttachmentInfo {
		s := FramebufferAttachmentInfo{}
		s.Init()
		s.Object = uint32(stencil_id)
		s.Type = GLenum_GL_RENDERBUFFER
		s.CubeMapFace = GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X
		return s
	}()
	ctx.Instances.Framebuffers[FramebufferId(uint32(0))] = backbuffer
	ctx.BoundFramebuffers[GLenum_GL_DRAW_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.BoundFramebuffers[GLenum_GL_READ_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.Rasterizing.StencilMask[GLenum_GL_FRONT] = GLuint(uint32(4294967295))
	ctx.Rasterizing.StencilMask[GLenum_GL_BACK] = GLuint(uint32(4294967295))
	ctx.PixelStorage[GLenum_GL_PACK_ALIGNMENT] = GLint(int32(4))
	ctx.PixelStorage[GLenum_GL_UNPACK_ALIGNMENT] = GLint(int32(4))
	for i := int32(int32(0)); i < int32(64); i++ {
		ctx.VertexAttributeArrays[AttributeLocation(i)] = func() *VertexAttributeArray {
			s := &VertexAttributeArray{}
			s.Init()
			return s
		}()
	}
	CreateContext_1125_result := ctx // Contextʳ
	ϟc.CGLContexts[context] = CreateContext_1125_result
	ϟa.Ctx.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(context, ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Result = ϟa.Result
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_1125_result
	return nil
}
func (ϟa *CGLSetCurrentContext) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	SetContext_1126_context := ϟc.CGLContexts.Get(ϟa.Ctx) // Contextʳ
	ϟc.Contexts[ϟc.CurrentThread] = SetContext_1126_context
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = SetContext_1126_context
	return nil
}
func (ϟa *CGLGetSurface) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Cid.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Cid.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Wid.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Wid.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Sid.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Sid.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Result = ϟa.Result
	return nil
}
func (ϟa *CGSGetSurfaceBounds) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Bounds.Slice(uint64(0), uint64(4), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Result = ϟa.Result
	return nil
}
func (ϟa *CGLFlushDrawable) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	return nil
}
func (ϟa *GlGetQueryObjecti64v) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}
func (ϟa *GlGetQueryObjectui64v) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).MapMemory(ϟa, ϟs, ϟd, ϟl, ϟb).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}
func (ϟa *ReplayCreateRenderer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *ReplayBindRenderer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *SwitchThread) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟc.CurrentThread = ϟa.ThreadID
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *BackbufferInfo) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_1127_result := context            // Contextʳ
	ctx := GetContext_1127_result                // Contextʳ
	ctx.PreserveBuffersOnSwap = ϟa.PreserveBuffersOnSwap
	backbuffer := ctx.Instances.Framebuffers.Get(FramebufferId(uint32(0)))                        // Framebufferʳ
	color_id := RenderbufferId(backbuffer.Attachments.Get(GLenum_GL_COLOR_ATTACHMENT0).Object)    // RenderbufferId
	color_buffer := ctx.Instances.Renderbuffers.Get(color_id)                                     // Renderbufferʳ
	depth_id := RenderbufferId(backbuffer.Attachments.Get(GLenum_GL_DEPTH_ATTACHMENT).Object)     // RenderbufferId
	depth_buffer := ctx.Instances.Renderbuffers.Get(depth_id)                                     // Renderbufferʳ
	stencil_id := RenderbufferId(backbuffer.Attachments.Get(GLenum_GL_STENCIL_ATTACHMENT).Object) // RenderbufferId
	stencil_buffer := ctx.Instances.Renderbuffers.Get(stencil_id)                                 // Renderbufferʳ
	color_buffer.Width = ϟa.Width
	color_buffer.Height = ϟa.Height
	color_buffer.Format = ϟa.ColorFmt
	depth_buffer.Width = ϟa.Width
	depth_buffer.Height = ϟa.Height
	depth_buffer.Format = ϟa.DepthFmt
	stencil_buffer.Width = ϟa.Width
	stencil_buffer.Height = ϟa.Height
	stencil_buffer.Format = ϟa.StencilFmt
	if ϟa.ResetViewportScissor {
		ctx.Rasterizing.Scissor.Width = ϟa.Width
		ctx.Rasterizing.Scissor.Height = ϟa.Height
		ctx.Rasterizing.Viewport.Width = ϟa.Width
		ctx.Rasterizing.Viewport.Height = ϟa.Height
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _, _, _ = context, GetContext_1127_result, ctx, backbuffer, color_id, color_buffer, depth_id, depth_buffer, stencil_id, stencil_buffer
	return nil
}
func (ϟa *StartTimer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (ϟa *StopTimer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	return nil
}
func (ϟa *FlushPostBuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
