////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

package gles

import (
	"fmt"
	"strings"

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
	case GLenum_GL_PRIMITIVES_GENERATED:
		minRequiredVersion_37_major := uint32(3) // u32
		minRequiredVersion_37_minor := uint32(2) // u32
		_, _ = minRequiredVersion_37_major, minRequiredVersion_37_minor
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
	minRequiredVersion_39_major := uint32(3)                               // u32
	minRequiredVersion_39_minor := uint32(0)                               // u32
	q := ϟa.Queries.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // QueryIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                           // Contextʳ
	GetContext_40_result := context                                        // Contextʳ
	ctx := GetContext_40_result                                            // Contextʳ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.Queries, q.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_39_major, minRequiredVersion_39_minor, q, context, GetContext_40_result, ctx
	return nil
}
func (ϟa *GlEndQuery) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_41_major := uint32(3) // u32
	minRequiredVersion_41_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ANY_SAMPLES_PASSED, GLenum_GL_ANY_SAMPLES_PASSED_CONSERVATIVE, GLenum_GL_TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN:
	case GLenum_GL_PRIMITIVES_GENERATED:
		minRequiredVersion_42_major := uint32(3) // u32
		minRequiredVersion_42_minor := uint32(2) // u32
		_, _ = minRequiredVersion_42_major, minRequiredVersion_42_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_41_major, minRequiredVersion_41_minor
	return nil
}
func (ϟa *GlGenQueries) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_44_major := uint32(3)                               // u32
	minRequiredVersion_44_minor := uint32(0)                               // u32
	q := ϟa.Queries.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // QueryIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                           // Contextʳ
	GetContext_45_result := context                                        // Contextʳ
	ctx := GetContext_45_result                                            // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := QueryId(ϟa.Queries.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // QueryId
		ctx.Instances.Queries[id] = &Query{}
		q.Index(uint64(i), ϟs).Write(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = minRequiredVersion_44_major, minRequiredVersion_44_minor, q, context, GetContext_45_result, ctx
	return nil
}
func (ϟa *GlGetQueryObjectuiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_46_major := uint32(3) // u32
	minRequiredVersion_46_minor := uint32(0) // u32
	switch ϟa.Parameter {
	case GLenum_GL_QUERY_RESULT, GLenum_GL_QUERY_RESULT_AVAILABLE:
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_46_major, minRequiredVersion_46_minor
	return nil
}
func (ϟa *GlGetQueryiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_48_major := uint32(3) // u32
	minRequiredVersion_48_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ANY_SAMPLES_PASSED, GLenum_GL_ANY_SAMPLES_PASSED_CONSERVATIVE, GLenum_GL_TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN:
	case GLenum_GL_PRIMITIVES_GENERATED:
		minRequiredVersion_49_major := uint32(3) // u32
		minRequiredVersion_49_minor := uint32(2) // u32
		_, _ = minRequiredVersion_49_major, minRequiredVersion_49_minor
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
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_48_major, minRequiredVersion_48_minor
	return nil
}
func (ϟa *GlIsQuery) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_52_major := uint32(3)     // u32
	minRequiredVersion_52_minor := uint32(0)     // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_53_result := context              // Contextʳ
	ctx := GetContext_53_result                  // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.Queries.Contains(ϟa.Query) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = minRequiredVersion_52_major, minRequiredVersion_52_minor, context, GetContext_53_result, ctx
	return nil
}
func (ϟa *GlBindBuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_54_major := uint32(2) // u32
	minRequiredVersion_54_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER:
	case GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
		minRequiredVersion_55_major := uint32(3) // u32
		minRequiredVersion_55_minor := uint32(0) // u32
		_, _ = minRequiredVersion_55_major, minRequiredVersion_55_minor
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_56_major := uint32(3) // u32
		minRequiredVersion_56_minor := uint32(1) // u32
		_, _ = minRequiredVersion_56_major, minRequiredVersion_56_minor
	case GLenum_GL_TEXTURE_BUFFER:
		minRequiredVersion_57_major := uint32(3) // u32
		minRequiredVersion_57_minor := uint32(2) // u32
		_, _ = minRequiredVersion_57_major, minRequiredVersion_57_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_59_result := context              // Contextʳ
	ctx := GetContext_59_result                  // Contextʳ
	if !(ctx.Instances.Buffers.Contains(ϟa.Buffer)) {
		ctx.Instances.Buffers[ϟa.Buffer] = &Buffer{Size: GLsizeiptr(int32(0)), Usage: GLenum_GL_STATIC_DRAW}
	}
	ctx.BoundBuffers[ϟa.Target] = ϟa.Buffer
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_54_major, minRequiredVersion_54_minor, context, GetContext_59_result, ctx
	return nil
}
func (ϟa *GlBindBufferBase) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_60_major := uint32(3) // u32
	minRequiredVersion_60_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_61_major := uint32(3) // u32
		minRequiredVersion_61_minor := uint32(1) // u32
		_, _ = minRequiredVersion_61_major, minRequiredVersion_61_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_60_major, minRequiredVersion_60_minor
	return nil
}
func (ϟa *GlBindBufferRange) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_63_major := uint32(3) // u32
	minRequiredVersion_63_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_64_major := uint32(3) // u32
		minRequiredVersion_64_minor := uint32(1) // u32
		_, _ = minRequiredVersion_64_major, minRequiredVersion_64_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_63_major, minRequiredVersion_63_minor
	return nil
}
func (ϟa *GlBufferData) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_66_major := uint32(2) // u32
	minRequiredVersion_66_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER:
	case GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
		minRequiredVersion_67_major := uint32(3) // u32
		minRequiredVersion_67_minor := uint32(0) // u32
		_, _ = minRequiredVersion_67_major, minRequiredVersion_67_minor
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_68_major := uint32(3) // u32
		minRequiredVersion_68_minor := uint32(1) // u32
		_, _ = minRequiredVersion_68_major, minRequiredVersion_68_minor
	case GLenum_GL_TEXTURE_BUFFER:
		minRequiredVersion_69_major := uint32(3) // u32
		minRequiredVersion_69_minor := uint32(2) // u32
		_, _ = minRequiredVersion_69_major, minRequiredVersion_69_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Usage {
	case GLenum_GL_DYNAMIC_DRAW, GLenum_GL_STATIC_DRAW, GLenum_GL_STREAM_DRAW:
	case GLenum_GL_DYNAMIC_COPY, GLenum_GL_DYNAMIC_READ, GLenum_GL_STATIC_COPY, GLenum_GL_STATIC_READ, GLenum_GL_STREAM_COPY, GLenum_GL_STREAM_READ:
		minRequiredVersion_71_major := uint32(3) // u32
		minRequiredVersion_71_minor := uint32(0) // u32
		_, _ = minRequiredVersion_71_major, minRequiredVersion_71_minor
	default:
		v := ϟa.Usage
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_73_result := context              // Contextʳ
	ctx := GetContext_73_result                  // Contextʳ
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
	_, _, _, _, _, _, _ = minRequiredVersion_66_major, minRequiredVersion_66_minor, context, GetContext_73_result, ctx, id, b
	return nil
}
func (ϟa *GlBufferSubData) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_74_major := uint32(2) // u32
	minRequiredVersion_74_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER:
	case GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
		minRequiredVersion_75_major := uint32(3) // u32
		minRequiredVersion_75_minor := uint32(0) // u32
		_, _ = minRequiredVersion_75_major, minRequiredVersion_75_minor
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_76_major := uint32(3) // u32
		minRequiredVersion_76_minor := uint32(1) // u32
		_, _ = minRequiredVersion_76_major, minRequiredVersion_76_minor
	case GLenum_GL_TEXTURE_BUFFER:
		minRequiredVersion_77_major := uint32(3) // u32
		minRequiredVersion_77_minor := uint32(2) // u32
		_, _ = minRequiredVersion_77_major, minRequiredVersion_77_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.Data.Slice(uint64(GLsizeiptr(int32(0))), uint64(ϟa.Size), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_74_major, minRequiredVersion_74_minor
	return nil
}
func (ϟa *GlCopyBufferSubData) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_79_major := uint32(3) // u32
	minRequiredVersion_79_minor := uint32(0) // u32
	switch ϟa.ReadTarget {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER, GLenum_GL_TEXTURE_BUFFER:
		minRequiredVersion_80_major := uint32(3) // u32
		minRequiredVersion_80_minor := uint32(2) // u32
		_, _ = minRequiredVersion_80_major, minRequiredVersion_80_minor
	default:
		v := ϟa.ReadTarget
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.WriteTarget {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER, GLenum_GL_TEXTURE_BUFFER:
		minRequiredVersion_82_major := uint32(3) // u32
		minRequiredVersion_82_minor := uint32(2) // u32
		_, _ = minRequiredVersion_82_major, minRequiredVersion_82_minor
	default:
		v := ϟa.WriteTarget
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_79_major, minRequiredVersion_79_minor
	return nil
}
func (ϟa *GlDeleteBuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_84_major := uint32(2)                               // u32
	minRequiredVersion_84_minor := uint32(0)                               // u32
	b := ϟa.Buffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // BufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                           // Contextʳ
	GetContext_85_result := context                                        // Contextʳ
	ctx := GetContext_85_result                                            // Contextʳ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.Buffers, b.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_84_major, minRequiredVersion_84_minor, b, context, GetContext_85_result, ctx
	return nil
}
func (ϟa *GlGenBuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_86_major := uint32(2)                               // u32
	minRequiredVersion_86_minor := uint32(0)                               // u32
	b := ϟa.Buffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // BufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                           // Contextʳ
	GetContext_87_result := context                                        // Contextʳ
	ctx := GetContext_87_result                                            // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := BufferId(ϟa.Buffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // BufferId
		ctx.Instances.Buffers[id] = &Buffer{Size: GLsizeiptr(int32(0)), Usage: GLenum_GL_STATIC_DRAW}
		b.Index(uint64(i), ϟs).Write(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = minRequiredVersion_86_major, minRequiredVersion_86_minor, b, context, GetContext_87_result, ctx
	return nil
}
func (ϟa *GlGetBufferParameteri64v) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_88_major := uint32(3) // u32
	minRequiredVersion_88_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER, GLenum_GL_TEXTURE_BUFFER:
		minRequiredVersion_89_major := uint32(3) // u32
		minRequiredVersion_89_minor := uint32(2) // u32
		_, _ = minRequiredVersion_89_major, minRequiredVersion_89_minor
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
	_, _ = minRequiredVersion_88_major, minRequiredVersion_88_minor
	return nil
}
func (ϟa *GlGetBufferParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_92_major := uint32(2) // u32
	minRequiredVersion_92_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER:
	case GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
		minRequiredVersion_93_major := uint32(3) // u32
		minRequiredVersion_93_minor := uint32(0) // u32
		_, _ = minRequiredVersion_93_major, minRequiredVersion_93_minor
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER, GLenum_GL_TEXTURE_BUFFER:
		minRequiredVersion_94_major := uint32(3) // u32
		minRequiredVersion_94_minor := uint32(2) // u32
		_, _ = minRequiredVersion_94_major, minRequiredVersion_94_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Parameter {
	case GLenum_GL_BUFFER_SIZE, GLenum_GL_BUFFER_USAGE:
	case GLenum_GL_BUFFER_ACCESS_FLAGS, GLenum_GL_BUFFER_MAPPED, GLenum_GL_BUFFER_MAP_LENGTH, GLenum_GL_BUFFER_MAP_OFFSET:
		minRequiredVersion_96_major := uint32(3) // u32
		minRequiredVersion_96_minor := uint32(0) // u32
		_, _ = minRequiredVersion_96_major, minRequiredVersion_96_minor
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_98_result := context              // Contextʳ
	ctx := GetContext_98_result                  // Contextʳ
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
	_, _, _, _, _, _, _ = minRequiredVersion_92_major, minRequiredVersion_92_minor, context, GetContext_98_result, ctx, id, b
	return nil
}
func (ϟa *GlGetBufferPointerv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_99_major := uint32(3) // u32
	minRequiredVersion_99_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER, GLenum_GL_TEXTURE_BUFFER:
		minRequiredVersion_100_major := uint32(3) // u32
		minRequiredVersion_100_minor := uint32(2) // u32
		_, _ = minRequiredVersion_100_major, minRequiredVersion_100_minor
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
	_, _ = minRequiredVersion_99_major, minRequiredVersion_99_minor
	return nil
}
func (ϟa *GlIsBuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_103_major := uint32(2)    // u32
	minRequiredVersion_103_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_104_result := context             // Contextʳ
	ctx := GetContext_104_result                 // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.Buffers.Contains(ϟa.Buffer) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = minRequiredVersion_103_major, minRequiredVersion_103_minor, context, GetContext_104_result, ctx
	return nil
}
func (ϟa *GlMapBufferRange) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_105_major := uint32(3) // u32
	minRequiredVersion_105_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_106_major := uint32(3) // u32
		minRequiredVersion_106_minor := uint32(1) // u32
		_, _ = minRequiredVersion_106_major, minRequiredVersion_106_minor
	case GLenum_GL_TEXTURE_BUFFER:
		minRequiredVersion_107_major := uint32(3) // u32
		minRequiredVersion_107_minor := uint32(2) // u32
		_, _ = minRequiredVersion_107_major, minRequiredVersion_107_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	supportsBits_109_seenBits := ϟa.Access                                                                                                                                                                                                                                      // GLbitfield
	supportsBits_109_validBits := (GLbitfield_GL_MAP_FLUSH_EXPLICIT_BIT) | ((GLbitfield_GL_MAP_INVALIDATE_BUFFER_BIT) | ((GLbitfield_GL_MAP_INVALIDATE_RANGE_BIT) | ((GLbitfield_GL_MAP_READ_BIT) | ((GLbitfield_GL_MAP_UNSYNCHRONIZED_BIT) | (GLbitfield_GL_MAP_WRITE_BIT))))) // GLbitfield
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
	GetContext_110_result := context                                // Contextʳ
	ctx := GetContext_110_result                                    // Contextʳ
	b := ctx.Instances.Buffers.Get(ctx.BoundBuffers.Get(ϟa.Target)) // Bufferʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ptr := U8ᵖ(ϟa.Result) // U8ᵖ
	b.MappingAccess = ϟa.Access
	b.MappingData = ptr.Slice(uint64(GLsizeiptr(int32(0))), uint64(ϟa.Length), ϟs)
	externs{ϟa, ϟs, ϟd, ϟl, ϟb}.mapMemory(b.MappingData)
	if (GLbitfield_GL_MAP_READ_BIT)&(ϟa.Access) != 0 {
		src := b.Data.Slice(uint64(ϟa.Offset), uint64((ϟa.Offset)+(GLintptr(ϟa.Length))), ϟs) // U8ˢ
		dst := b.MappingData                                                                  // U8ˢ
		dst.OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = src, dst
	}
	ϟa.Result = Voidᵖ(ptr)
	_, _, _, _, _, _, _, _, _ = minRequiredVersion_105_major, minRequiredVersion_105_minor, supportsBits_109_seenBits, supportsBits_109_validBits, context, GetContext_110_result, ctx, b, ptr
	return nil
}
func (ϟa *GlUnmapBuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_111_major := uint32(3) // u32
	minRequiredVersion_111_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_112_major := uint32(3) // u32
		minRequiredVersion_112_minor := uint32(1) // u32
		_, _ = minRequiredVersion_112_major, minRequiredVersion_112_minor
	case GLenum_GL_TEXTURE_BUFFER:
		minRequiredVersion_113_major := uint32(3) // u32
		minRequiredVersion_113_minor := uint32(2) // u32
		_, _ = minRequiredVersion_113_major, minRequiredVersion_113_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                    // Contextʳ
	GetContext_115_result := context                                // Contextʳ
	ctx := GetContext_115_result                                    // Contextʳ
	b := ctx.Instances.Buffers.Get(ctx.BoundBuffers.Get(ϟa.Target)) // Bufferʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	b.Data.Slice(uint64(b.MappingOffset), uint64((b.MappingOffset)+(int32(b.MappingData.Count))), ϟs).Copy(b.MappingData, ϟa, ϟs, ϟd, ϟl, ϟb)
	externs{ϟa, ϟs, ϟd, ϟl, ϟb}.unmapMemory(b.MappingData)
	b.MappingOffset = int32(0)
	b.MappingData = U8ˢ{}
	ϟa.Result = ϟa.Result
	_, _, _, _, _, _ = minRequiredVersion_111_major, minRequiredVersion_111_minor, context, GetContext_115_result, ctx, b
	return nil
}
func (ϟa *GlDrawArrays) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_116_major := uint32(2) // u32
	minRequiredVersion_116_minor := uint32(0) // u32
	switch ϟa.DrawMode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	case GLenum_GL_LINES_ADJACENCY, GLenum_GL_LINE_STRIP_ADJACENCY, GLenum_GL_PATCHES, GLenum_GL_TRIANGLES_ADJACENCY, GLenum_GL_TRIANGLE_STRIP_ADJACENCY:
		minRequiredVersion_117_major := uint32(3) // u32
		minRequiredVersion_117_minor := uint32(2) // u32
		_, _ = minRequiredVersion_117_major, minRequiredVersion_117_minor
	default:
		v := ϟa.DrawMode
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                    // Contextʳ
	GetContext_119_result := context                                                // Contextʳ
	ctx := GetContext_119_result                                                    // Contextʳ
	last_index := (uint32(ϟa.FirstIndex)) + ((uint32(ϟa.IndexCount)) - (uint32(1))) // u32
	ReadVertexArrays_120_ctx := ctx                                                 // Contextʳ
	ReadVertexArrays_120_first_index := uint32(ϟa.FirstIndex)                       // u32
	ReadVertexArrays_120_last_index := last_index                                   // u32
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
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _, _ = minRequiredVersion_116_major, minRequiredVersion_116_minor, context, GetContext_119_result, ctx, last_index, ReadVertexArrays_120_ctx, ReadVertexArrays_120_first_index, ReadVertexArrays_120_last_index
	return nil
}
func (ϟa *GlDrawArraysIndirect) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_122_major := uint32(3) // u32
	minRequiredVersion_122_minor := uint32(1) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	case GLenum_GL_LINES_ADJACENCY, GLenum_GL_LINE_STRIP_ADJACENCY, GLenum_GL_PATCHES, GLenum_GL_TRIANGLES_ADJACENCY, GLenum_GL_TRIANGLE_STRIP_ADJACENCY:
		minRequiredVersion_123_major := uint32(3) // u32
		minRequiredVersion_123_minor := uint32(2) // u32
		_, _ = minRequiredVersion_123_major, minRequiredVersion_123_minor
	default:
		v := ϟa.Mode
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_122_major, minRequiredVersion_122_minor
	return nil
}
func (ϟa *GlDrawArraysInstanced) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_125_major := uint32(3) // u32
	minRequiredVersion_125_minor := uint32(0) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	case GLenum_GL_LINES_ADJACENCY, GLenum_GL_LINE_STRIP_ADJACENCY, GLenum_GL_PATCHES, GLenum_GL_TRIANGLES_ADJACENCY, GLenum_GL_TRIANGLE_STRIP_ADJACENCY:
		minRequiredVersion_126_major := uint32(3) // u32
		minRequiredVersion_126_minor := uint32(2) // u32
		_, _ = minRequiredVersion_126_major, minRequiredVersion_126_minor
	default:
		v := ϟa.Mode
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_125_major, minRequiredVersion_125_minor
	return nil
}
func (ϟa *GlDrawBuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_128_major := uint32(3) // u32
	minRequiredVersion_128_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_128_major, minRequiredVersion_128_minor
	return nil
}
func (ϟa *GlDrawElements) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_129_major := uint32(2) // u32
	minRequiredVersion_129_minor := uint32(0) // u32
	switch ϟa.DrawMode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	case GLenum_GL_LINES_ADJACENCY, GLenum_GL_LINE_STRIP_ADJACENCY, GLenum_GL_PATCHES, GLenum_GL_TRIANGLES_ADJACENCY, GLenum_GL_TRIANGLE_STRIP_ADJACENCY:
		minRequiredVersion_130_major := uint32(3) // u32
		minRequiredVersion_130_minor := uint32(2) // u32
		_, _ = minRequiredVersion_130_major, minRequiredVersion_130_minor
	default:
		v := ϟa.DrawMode
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.IndicesType {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_SHORT:
	case GLenum_GL_UNSIGNED_INT:
		minRequiredVersion_132_major := uint32(3) // u32
		minRequiredVersion_132_minor := uint32(0) // u32
		_, _ = minRequiredVersion_132_major, minRequiredVersion_132_minor
	default:
		v := ϟa.IndicesType
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)               // Contextʳ
	GetContext_134_result := context                           // Contextʳ
	ctx := GetContext_134_result                               // Contextʳ
	count := uint32(ϟa.ElementCount)                           // u32
	id := ctx.BoundBuffers.Get(GLenum_GL_ELEMENT_ARRAY_BUFFER) // BufferId
	if (id) != (BufferId(uint32(0))) {
		index_data := ctx.Instances.Buffers.Get(id).Data                                                           // U8ˢ
		offset := uint32(uint64(ϟa.Indices.Address))                                                               // u32
		first := externs{ϟa, ϟs, ϟd, ϟl, ϟb}.minIndex(U8ᵖ(index_data.Index(0, ϟs)), ϟa.IndicesType, offset, count) // u32
		last := externs{ϟa, ϟs, ϟd, ϟl, ϟb}.maxIndex(U8ᵖ(index_data.Index(0, ϟs)), ϟa.IndicesType, offset, count)  // u32
		ReadVertexArrays_135_ctx := ctx                                                                            // Contextʳ
		ReadVertexArrays_135_first_index := first                                                                  // u32
		ReadVertexArrays_135_last_index := last                                                                    // u32
		for i := int32(int32(0)); i < int32(len(ReadVertexArrays_135_ctx.VertexAttributeArrays)); i++ {
			arr := ReadVertexArrays_135_ctx.VertexAttributeArrays.Get(AttributeLocation(i)) // VertexAttributeArrayʳ
			if (arr.Enabled) && ((arr.Buffer) == (BufferId(uint32(0)))) {
				vertexAttribTypeSize_136_t := arr.Type // GLenum
				vertexAttribTypeSize_136_result := func() (result uint32) {
					switch vertexAttribTypeSize_136_t {
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
						panic(fmt.Errorf("Unmatched switch(%v) in atom %T", vertexAttribTypeSize_136_t, ϟa))
						return result
					}
				}() // u32
				elsize := (vertexAttribTypeSize_136_result) * (arr.Size) // u32
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
				for v := uint32(ReadVertexArrays_135_first_index); v < (ReadVertexArrays_135_last_index)+(uint32(1)); v++ {
					offset := (elstride) * (v) // u32
					arr.Pointer.Slice(uint64(offset), uint64((offset)+(elsize)), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
					_ = offset
				}
				_, _, _, _ = vertexAttribTypeSize_136_t, vertexAttribTypeSize_136_result, elsize, elstride
			}
			_ = arr
		}
		_, _, _, _, _, _, _ = index_data, offset, first, last, ReadVertexArrays_135_ctx, ReadVertexArrays_135_first_index, ReadVertexArrays_135_last_index
	} else {
		index_data := U8ᵖ(ϟa.Indices)                                                               // U8ᵖ
		first := externs{ϟa, ϟs, ϟd, ϟl, ϟb}.minIndex(index_data, ϟa.IndicesType, uint32(0), count) // u32
		last := externs{ϟa, ϟs, ϟd, ϟl, ϟb}.maxIndex(index_data, ϟa.IndicesType, uint32(0), count)  // u32
		ReadVertexArrays_137_ctx := ctx                                                             // Contextʳ
		ReadVertexArrays_137_first_index := first                                                   // u32
		ReadVertexArrays_137_last_index := last                                                     // u32
		for i := int32(int32(0)); i < int32(len(ReadVertexArrays_137_ctx.VertexAttributeArrays)); i++ {
			arr := ReadVertexArrays_137_ctx.VertexAttributeArrays.Get(AttributeLocation(i)) // VertexAttributeArrayʳ
			if (arr.Enabled) && ((arr.Buffer) == (BufferId(uint32(0)))) {
				vertexAttribTypeSize_138_t := arr.Type // GLenum
				vertexAttribTypeSize_138_result := func() (result uint32) {
					switch vertexAttribTypeSize_138_t {
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
						panic(fmt.Errorf("Unmatched switch(%v) in atom %T", vertexAttribTypeSize_138_t, ϟa))
						return result
					}
				}() // u32
				elsize := (vertexAttribTypeSize_138_result) * (arr.Size) // u32
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
				for v := uint32(ReadVertexArrays_137_first_index); v < (ReadVertexArrays_137_last_index)+(uint32(1)); v++ {
					offset := (elstride) * (v) // u32
					arr.Pointer.Slice(uint64(offset), uint64((offset)+(elsize)), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
					_ = offset
				}
				_, _, _, _ = vertexAttribTypeSize_138_t, vertexAttribTypeSize_138_result, elsize, elstride
			}
			_ = arr
		}
		IndexSize_139_indices_type := ϟa.IndicesType // GLenum
		IndexSize_139_result := func() (result uint32) {
			switch IndexSize_139_indices_type {
			case GLenum_GL_UNSIGNED_BYTE:
				return uint32(1)
			case GLenum_GL_UNSIGNED_SHORT:
				return uint32(2)
			case GLenum_GL_UNSIGNED_INT:
				return uint32(4)
			default:
				// TODO: better unmatched handling
				panic(fmt.Errorf("Unmatched switch(%v) in atom %T", IndexSize_139_indices_type, ϟa))
				return result
			}
		}() // u32
		index_data.Slice(uint64(uint32(0)), uint64((uint32(ϟa.ElementCount))*(IndexSize_139_result)), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _, _, _, _, _, _, _ = index_data, first, last, ReadVertexArrays_137_ctx, ReadVertexArrays_137_first_index, ReadVertexArrays_137_last_index, IndexSize_139_indices_type, IndexSize_139_result
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = minRequiredVersion_129_major, minRequiredVersion_129_minor, context, GetContext_134_result, ctx, count, id
	return nil
}
func (ϟa *GlDrawElementsBaseVertex) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_140_major := uint32(3) // u32
	minRequiredVersion_140_minor := uint32(2) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINES_ADJACENCY, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_LINE_STRIP_ADJACENCY, GLenum_GL_PATCHES, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLES_ADJACENCY, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP, GLenum_GL_TRIANGLE_STRIP_ADJACENCY:
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
	_, _ = minRequiredVersion_140_major, minRequiredVersion_140_minor
	return nil
}
func (ϟa *GlDrawElementsIndirect) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_143_major := uint32(3) // u32
	minRequiredVersion_143_minor := uint32(1) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	case GLenum_GL_LINES_ADJACENCY, GLenum_GL_LINE_STRIP_ADJACENCY, GLenum_GL_PATCHES, GLenum_GL_TRIANGLES_ADJACENCY, GLenum_GL_TRIANGLE_STRIP_ADJACENCY:
		minRequiredVersion_144_major := uint32(3) // u32
		minRequiredVersion_144_minor := uint32(2) // u32
		_, _ = minRequiredVersion_144_major, minRequiredVersion_144_minor
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
	_, _ = minRequiredVersion_143_major, minRequiredVersion_143_minor
	return nil
}
func (ϟa *GlDrawElementsInstanced) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_147_major := uint32(3) // u32
	minRequiredVersion_147_minor := uint32(0) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	case GLenum_GL_LINES_ADJACENCY, GLenum_GL_LINE_STRIP_ADJACENCY, GLenum_GL_PATCHES, GLenum_GL_TRIANGLES_ADJACENCY, GLenum_GL_TRIANGLE_STRIP_ADJACENCY:
		minRequiredVersion_148_major := uint32(3) // u32
		minRequiredVersion_148_minor := uint32(2) // u32
		_, _ = minRequiredVersion_148_major, minRequiredVersion_148_minor
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
	_, _ = minRequiredVersion_147_major, minRequiredVersion_147_minor
	return nil
}
func (ϟa *GlDrawElementsInstancedBaseVertex) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_151_major := uint32(3) // u32
	minRequiredVersion_151_minor := uint32(2) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINES_ADJACENCY, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_LINE_STRIP_ADJACENCY, GLenum_GL_PATCHES, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLES_ADJACENCY, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP, GLenum_GL_TRIANGLE_STRIP_ADJACENCY:
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
	_, _ = minRequiredVersion_151_major, minRequiredVersion_151_minor
	return nil
}
func (ϟa *GlDrawRangeElements) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_154_major := uint32(3) // u32
	minRequiredVersion_154_minor := uint32(0) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	case GLenum_GL_LINES_ADJACENCY, GLenum_GL_LINE_STRIP_ADJACENCY, GLenum_GL_PATCHES, GLenum_GL_TRIANGLES_ADJACENCY, GLenum_GL_TRIANGLE_STRIP_ADJACENCY:
		minRequiredVersion_155_major := uint32(3) // u32
		minRequiredVersion_155_minor := uint32(2) // u32
		_, _ = minRequiredVersion_155_major, minRequiredVersion_155_minor
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
	_, _ = minRequiredVersion_154_major, minRequiredVersion_154_minor
	return nil
}
func (ϟa *GlDrawRangeElementsBaseVertex) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_158_major := uint32(3) // u32
	minRequiredVersion_158_minor := uint32(2) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINES_ADJACENCY, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_LINE_STRIP_ADJACENCY, GLenum_GL_PATCHES, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLES_ADJACENCY, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP, GLenum_GL_TRIANGLE_STRIP_ADJACENCY:
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
	_, _ = minRequiredVersion_158_major, minRequiredVersion_158_minor
	return nil
}
func (ϟa *GlPatchParameteri) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_161_major := uint32(3) // u32
	minRequiredVersion_161_minor := uint32(2) // u32
	switch ϟa.Pname {
	case GLenum_GL_PATCH_VERTICES:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_161_major, minRequiredVersion_161_minor
	return nil
}
func (ϟa *GlPrimitiveBoundingBox) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_163_major := uint32(3) // u32
	minRequiredVersion_163_minor := uint32(2) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_163_major, minRequiredVersion_163_minor
	return nil
}
func (ϟa *GlActiveShaderProgramEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_164_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_164_ext
	return nil
}
func (ϟa *GlAlphaFuncQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_165_ext := ExtensionId_GL_QCOM_alpha_test // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_165_ext
	return nil
}
func (ϟa *GlApplyFramebufferAttachmentCMAAINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_166_ext := ExtensionId_GL_INTEL_framebuffer_CMAA // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_166_ext
	return nil
}
func (ϟa *GlBeginConditionalRenderNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_167_ext := ExtensionId_GL_NV_conditional_render // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_167_ext
	return nil
}
func (ϟa *GlBeginPerfMonitorAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_168_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_168_ext
	return nil
}
func (ϟa *GlBeginPerfQueryINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_169_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_169_ext
	return nil
}
func (ϟa *GlBeginQueryEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_170_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_171_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = requiresExtension_170_ext, requiresExtension_171_ext
	return nil
}
func (ϟa *GlBindProgramPipelineEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_172_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_172_ext
	return nil
}
func (ϟa *GlBindVertexArrayOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_173_ext := ExtensionId_GL_OES_vertex_array_object // ExtensionId
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                        // Contextʳ
	GetContext_174_result := context                                    // Contextʳ
	ctx := GetContext_174_result                                        // Contextʳ
	if !(ctx.Instances.VertexArrays.Contains(ϟa.Array)) {
		ctx.Instances.VertexArrays[ϟa.Array] = &VertexArray{}
	}
	ctx.BoundVertexArray = ϟa.Array
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = requiresExtension_173_ext, context, GetContext_174_result, ctx
	return nil
}
func (ϟa *GlBlendBarrierNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_175_ext := ExtensionId_GL_NV_blend_equation_advanced // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_175_ext
	return nil
}
func (ϟa *GlBlendEquationSeparateiOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_176_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_176_ext
	return nil
}
func (ϟa *GlBlendEquationiOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_177_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_177_ext
	return nil
}
func (ϟa *GlBlendFuncSeparateiOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_178_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_178_ext
	return nil
}
func (ϟa *GlBlendFunciOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_179_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_179_ext
	return nil
}
func (ϟa *GlBlendParameteriNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_180_ext := ExtensionId_GL_NV_blend_equation_advanced // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_180_ext
	return nil
}
func (ϟa *GlBlitFramebufferANGLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_181_ext := ExtensionId_GL_ANGLE_framebuffer_blit // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_181_ext
	return nil
}
func (ϟa *GlBlitFramebufferNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_182_ext := ExtensionId_GL_NV_framebuffer_blit // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_182_ext
	return nil
}
func (ϟa *GlBufferStorageEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_183_ext := ExtensionId_GL_EXT_buffer_storage // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_183_ext
	return nil
}
func (ϟa *GlClientWaitSyncAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_184_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_184_ext
	return nil
}
func (ϟa *GlColorMaskiOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_185_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_185_ext
	return nil
}
func (ϟa *GlCompressedTexImage3DOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_186_ext := ExtensionId_GL_OES_texture_3D // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_186_ext
	return nil
}
func (ϟa *GlCompressedTexSubImage3DOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_187_ext := ExtensionId_GL_OES_texture_3D // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_187_ext
	return nil
}
func (ϟa *GlCopyBufferSubDataNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_188_ext := ExtensionId_GL_NV_copy_buffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_188_ext
	return nil
}
func (ϟa *GlCopyImageSubDataOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_189_ext := ExtensionId_GL_OES_copy_image // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_189_ext
	return nil
}
func (ϟa *GlCopyPathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_190_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_190_ext
	return nil
}
func (ϟa *GlCopyTexSubImage3DOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_191_ext := ExtensionId_GL_OES_texture_3D // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_191_ext
	return nil
}
func (ϟa *GlCopyTextureLevelsAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_192_ext := ExtensionId_GL_APPLE_copy_texture_levels // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_192_ext
	return nil
}
func (ϟa *GlCoverFillPathInstancedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_193_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_193_ext
	return nil
}
func (ϟa *GlCoverFillPathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_194_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_194_ext
	return nil
}
func (ϟa *GlCoverStrokePathInstancedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_195_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_195_ext
	return nil
}
func (ϟa *GlCoverStrokePathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_196_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_196_ext
	return nil
}
func (ϟa *GlCoverageMaskNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_197_ext := ExtensionId_GL_NV_coverage_sample // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_197_ext
	return nil
}
func (ϟa *GlCoverageModulationNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_198_ext := ExtensionId_GL_NV_framebuffer_mixed_samples // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_198_ext
	return nil
}
func (ϟa *GlCoverageModulationTableNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_199_ext := ExtensionId_GL_NV_framebuffer_mixed_samples // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_199_ext
	return nil
}
func (ϟa *GlCoverageOperationNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_200_ext := ExtensionId_GL_NV_coverage_sample // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_200_ext
	return nil
}
func (ϟa *GlCreatePerfQueryINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_201_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_201_ext
	return nil
}
func (ϟa *GlCreateShaderProgramvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_202_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_202_ext
	return nil
}
func (ϟa *GlDeleteFencesNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_203_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_203_ext
	return nil
}
func (ϟa *GlDeletePathsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_204_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_204_ext
	return nil
}
func (ϟa *GlDeletePerfMonitorsAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_205_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_205_ext
	return nil
}
func (ϟa *GlDeletePerfQueryINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_206_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_206_ext
	return nil
}
func (ϟa *GlDeleteProgramPipelinesEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_207_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_207_ext
	return nil
}
func (ϟa *GlDeleteQueriesEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_208_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_209_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	q := ϟa.Queries.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs)  // QueryIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                            // Contextʳ
	GetContext_210_result := context                                        // Contextʳ
	ctx := GetContext_210_result                                            // Contextʳ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.Queries, q.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = requiresExtension_208_ext, requiresExtension_209_ext, q, context, GetContext_210_result, ctx
	return nil
}
func (ϟa *GlDeleteSyncAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_211_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_211_ext
	return nil
}
func (ϟa *GlDeleteVertexArraysOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_212_ext := ExtensionId_GL_OES_vertex_array_object   // ExtensionId
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                          // Contextʳ
	GetContext_213_result := context                                      // Contextʳ
	ctx := GetContext_213_result                                          // Contextʳ
	a := ϟa.Arrays.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // VertexArrayIdˢ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.VertexArrays, a.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = requiresExtension_212_ext, context, GetContext_213_result, ctx, a
	return nil
}
func (ϟa *GlDepthRangeArrayfvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_214_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_214_ext
	return nil
}
func (ϟa *GlDepthRangeIndexedfNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_215_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_215_ext
	return nil
}
func (ϟa *GlDisableDriverControlQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_216_ext := ExtensionId_GL_QCOM_driver_control // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_216_ext
	return nil
}
func (ϟa *GlDisableiNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_217_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_217_ext
	return nil
}
func (ϟa *GlDisableiOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_218_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_218_ext
	return nil
}
func (ϟa *GlDiscardFramebufferEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_219_ext := ExtensionId_GL_EXT_discard_framebuffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_219_ext
	return nil
}
func (ϟa *GlDrawArraysInstancedANGLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_220_ext := ExtensionId_GL_ANGLE_instanced_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_220_ext
	return nil
}
func (ϟa *GlDrawArraysInstancedBaseInstanceEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_221_ext := ExtensionId_GL_EXT_base_instance // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_221_ext
	return nil
}
func (ϟa *GlDrawArraysInstancedEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_222_ext := ExtensionId_GL_EXT_draw_instanced   // ExtensionId
	requiresExtension_223_ext := ExtensionId_GL_EXT_instanced_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = requiresExtension_222_ext, requiresExtension_223_ext
	return nil
}
func (ϟa *GlDrawArraysInstancedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_224_ext := ExtensionId_GL_NV_draw_instanced // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_224_ext
	return nil
}
func (ϟa *GlDrawBuffersEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_225_ext := ExtensionId_GL_EXT_draw_buffers // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_225_ext
	return nil
}
func (ϟa *GlDrawBuffersIndexedEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_226_ext := ExtensionId_GL_EXT_multiview_draw_buffers // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_226_ext
	return nil
}
func (ϟa *GlDrawBuffersNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_227_ext := ExtensionId_GL_NV_draw_buffers // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_227_ext
	return nil
}
func (ϟa *GlDrawElementsBaseVertexEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_228_ext := ExtensionId_GL_EXT_draw_elements_base_vertex // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_228_ext
	return nil
}
func (ϟa *GlDrawElementsBaseVertexOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_229_ext := ExtensionId_GL_OES_draw_elements_base_vertex // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_229_ext
	return nil
}
func (ϟa *GlDrawElementsInstancedANGLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_230_ext := ExtensionId_GL_ANGLE_instanced_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_230_ext
	return nil
}
func (ϟa *GlDrawElementsInstancedBaseInstanceEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_231_ext := ExtensionId_GL_EXT_base_instance // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_231_ext
	return nil
}
func (ϟa *GlDrawElementsInstancedBaseVertexBaseInstanceEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_232_ext := ExtensionId_GL_EXT_base_instance // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_232_ext
	return nil
}
func (ϟa *GlDrawElementsInstancedBaseVertexEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_233_ext := ExtensionId_GL_EXT_draw_elements_base_vertex // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_233_ext
	return nil
}
func (ϟa *GlDrawElementsInstancedBaseVertexOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_234_ext := ExtensionId_GL_OES_draw_elements_base_vertex // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_234_ext
	return nil
}
func (ϟa *GlDrawElementsInstancedEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_235_ext := ExtensionId_GL_EXT_draw_instanced   // ExtensionId
	requiresExtension_236_ext := ExtensionId_GL_EXT_instanced_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = requiresExtension_235_ext, requiresExtension_236_ext
	return nil
}
func (ϟa *GlDrawElementsInstancedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_237_ext := ExtensionId_GL_NV_draw_instanced // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_237_ext
	return nil
}
func (ϟa *GlDrawRangeElementsBaseVertexEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_238_ext := ExtensionId_GL_EXT_draw_elements_base_vertex // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_238_ext
	return nil
}
func (ϟa *GlDrawRangeElementsBaseVertexOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_239_ext := ExtensionId_GL_OES_draw_elements_base_vertex // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_239_ext
	return nil
}
func (ϟa *GlEGLImageTargetRenderbufferStorageOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_240_ext := ExtensionId_GL_OES_EGL_image // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_240_ext
	return nil
}
func (ϟa *GlEGLImageTargetTexture2DOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_241_ext := ExtensionId_GL_OES_EGL_image // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_241_ext
	return nil
}
func (ϟa *GlEnableDriverControlQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_242_ext := ExtensionId_GL_QCOM_driver_control // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_242_ext
	return nil
}
func (ϟa *GlEnableiNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_243_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_243_ext
	return nil
}
func (ϟa *GlEnableiOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_244_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_244_ext
	return nil
}
func (ϟa *GlEndConditionalRenderNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_245_ext := ExtensionId_GL_NV_conditional_render // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_245_ext
	return nil
}
func (ϟa *GlEndPerfMonitorAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_246_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_246_ext
	return nil
}
func (ϟa *GlEndPerfQueryINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_247_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_247_ext
	return nil
}
func (ϟa *GlEndQueryEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_248_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_249_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = requiresExtension_248_ext, requiresExtension_249_ext
	return nil
}
func (ϟa *GlEndTilingQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_250_ext := ExtensionId_GL_QCOM_tiled_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_250_ext
	return nil
}
func (ϟa *GlExtGetBufferPointervQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_251_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_251_ext
	return nil
}
func (ϟa *GlExtGetBuffersQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_252_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_252_ext
	return nil
}
func (ϟa *GlExtGetFramebuffersQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_253_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_253_ext
	return nil
}
func (ϟa *GlExtGetProgramBinarySourceQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_254_ext := ExtensionId_GL_QCOM_extended_get2 // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_254_ext
	return nil
}
func (ϟa *GlExtGetProgramsQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_255_ext := ExtensionId_GL_QCOM_extended_get2 // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_255_ext
	return nil
}
func (ϟa *GlExtGetRenderbuffersQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_256_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_256_ext
	return nil
}
func (ϟa *GlExtGetShadersQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_257_ext := ExtensionId_GL_QCOM_extended_get2 // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_257_ext
	return nil
}
func (ϟa *GlExtGetTexLevelParameterivQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_258_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_258_ext
	return nil
}
func (ϟa *GlExtGetTexSubImageQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_259_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_259_ext
	return nil
}
func (ϟa *GlExtGetTexturesQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_260_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_260_ext
	return nil
}
func (ϟa *GlExtIsProgramBinaryQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_261_ext := ExtensionId_GL_QCOM_extended_get2 // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_261_ext
	return nil
}
func (ϟa *GlExtTexObjectStateOverrideiQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_262_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_262_ext
	return nil
}
func (ϟa *GlFenceSyncAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_263_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_263_ext
	return nil
}
func (ϟa *GlFinishFenceNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_264_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_264_ext
	return nil
}
func (ϟa *GlFlushMappedBufferRangeEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_265_ext := ExtensionId_GL_EXT_map_buffer_range // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_265_ext
	return nil
}
func (ϟa *GlFragmentCoverageColorNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_266_ext := ExtensionId_GL_NV_fragment_coverage_to_color // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_266_ext
	return nil
}
func (ϟa *GlFramebufferSampleLocationsfvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_267_ext := ExtensionId_GL_NV_sample_locations // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_267_ext
	return nil
}
func (ϟa *GlFramebufferTexture2DMultisampleEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_268_ext := ExtensionId_GL_EXT_multisampled_render_to_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_268_ext
	return nil
}
func (ϟa *GlFramebufferTexture2DMultisampleIMG) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_269_ext := ExtensionId_GL_IMG_multisampled_render_to_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_269_ext
	return nil
}
func (ϟa *GlFramebufferTexture3DOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_270_ext := ExtensionId_GL_OES_texture_3D // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_270_ext
	return nil
}
func (ϟa *GlFramebufferTextureMultiviewOVR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_271_ext := ExtensionId_GL_OVR_multiview // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_271_ext
	return nil
}
func (ϟa *GlFramebufferTextureOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_272_ext := ExtensionId_GL_OES_geometry_shader // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_272_ext
	return nil
}
func (ϟa *GlGenFencesNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_273_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_273_ext
	return nil
}
func (ϟa *GlGenPathsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_274_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_274_ext
	return nil
}
func (ϟa *GlGenPerfMonitorsAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_275_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_275_ext
	return nil
}
func (ϟa *GlGenProgramPipelinesEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_276_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_276_ext
	return nil
}
func (ϟa *GlGenQueriesEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_277_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_278_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	q := ϟa.Queries.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs)  // QueryIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                            // Contextʳ
	GetContext_279_result := context                                        // Contextʳ
	ctx := GetContext_279_result                                            // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := QueryId(ϟa.Queries.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // QueryId
		ctx.Instances.Queries[id] = &Query{}
		q.Index(uint64(i), ϟs).Write(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = requiresExtension_277_ext, requiresExtension_278_ext, q, context, GetContext_279_result, ctx
	return nil
}
func (ϟa *GlGenVertexArraysOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_280_ext := ExtensionId_GL_OES_vertex_array_object   // ExtensionId
	a := ϟa.Arrays.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // VertexArrayIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                          // Contextʳ
	GetContext_281_result := context                                      // Contextʳ
	ctx := GetContext_281_result                                          // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := VertexArrayId(ϟa.Arrays.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // VertexArrayId
		ctx.Instances.VertexArrays[id] = &VertexArray{}
		a.Index(uint64(i), ϟs).Write(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _ = requiresExtension_280_ext, a, context, GetContext_281_result, ctx
	return nil
}
func (ϟa *GlGetBufferPointervOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_282_ext := ExtensionId_GL_OES_mapbuffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_282_ext
	return nil
}
func (ϟa *GlGetCoverageModulationTableNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_283_ext := ExtensionId_GL_NV_framebuffer_mixed_samples // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_283_ext
	return nil
}
func (ϟa *GlGetDriverControlStringQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_284_ext := ExtensionId_GL_QCOM_driver_control // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_284_ext
	return nil
}
func (ϟa *GlGetDriverControlsQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_285_ext := ExtensionId_GL_QCOM_driver_control // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_285_ext
	return nil
}
func (ϟa *GlGetFenceivNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_286_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_286_ext
	return nil
}
func (ϟa *GlGetFirstPerfQueryIdINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_287_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_287_ext
	return nil
}
func (ϟa *GlGetFloati_vNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_288_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_288_ext
	return nil
}
func (ϟa *GlGetGraphicsResetStatusEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_289_ext := ExtensionId_GL_EXT_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_289_ext
	return nil
}
func (ϟa *GlGetGraphicsResetStatusKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_290_ext := ExtensionId_GL_KHR_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_290_ext
	return nil
}
func (ϟa *GlGetImageHandleNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_291_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_291_ext
	return nil
}
func (ϟa *GlGetInteger64vAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_292_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_292_ext
	return nil
}
func (ϟa *GlGetIntegeri_vEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_293_ext := ExtensionId_GL_EXT_multiview_draw_buffers // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_293_ext
	return nil
}
func (ϟa *GlGetInternalformatSampleivNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_294_ext := ExtensionId_GL_NV_internalformat_sample_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_294_ext
	return nil
}
func (ϟa *GlGetNextPerfQueryIdINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_295_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_295_ext
	return nil
}
func (ϟa *GlGetObjectLabelEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_296_ext := ExtensionId_GL_EXT_debug_label // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_296_ext
	return nil
}
func (ϟa *GlGetPathCommandsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_297_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_297_ext
	return nil
}
func (ϟa *GlGetPathCoordsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_298_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_298_ext
	return nil
}
func (ϟa *GlGetPathDashArrayNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_299_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_299_ext
	return nil
}
func (ϟa *GlGetPathLengthNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_300_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_300_ext
	return nil
}
func (ϟa *GlGetPathMetricRangeNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_301_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_301_ext
	return nil
}
func (ϟa *GlGetPathMetricsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_302_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_302_ext
	return nil
}
func (ϟa *GlGetPathParameterfvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_303_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_303_ext
	return nil
}
func (ϟa *GlGetPathParameterivNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_304_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_304_ext
	return nil
}
func (ϟa *GlGetPathSpacingNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_305_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_305_ext
	return nil
}
func (ϟa *GlGetPerfCounterInfoINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_306_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_306_ext
	return nil
}
func (ϟa *GlGetPerfMonitorCounterDataAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_307_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_307_ext
	return nil
}
func (ϟa *GlGetPerfMonitorCounterInfoAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_308_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_308_ext
	return nil
}
func (ϟa *GlGetPerfMonitorCounterStringAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_309_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_309_ext
	return nil
}
func (ϟa *GlGetPerfMonitorCountersAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_310_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_310_ext
	return nil
}
func (ϟa *GlGetPerfMonitorGroupStringAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_311_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_311_ext
	return nil
}
func (ϟa *GlGetPerfMonitorGroupsAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_312_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_312_ext
	return nil
}
func (ϟa *GlGetPerfQueryDataINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_313_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_313_ext
	return nil
}
func (ϟa *GlGetPerfQueryIdByNameINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_314_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_314_ext
	return nil
}
func (ϟa *GlGetPerfQueryInfoINTEL) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_315_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_315_ext
	return nil
}
func (ϟa *GlGetProgramBinaryOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_316_ext := ExtensionId_GL_OES_get_program_binary // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	l := GLsizei(ϟa.BytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // GLsizei
	ϟa.BytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(l, ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.BinaryFormat.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.BinaryFormat.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Binary.Slice(uint64(GLsizei(int32(0))), uint64(l), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = requiresExtension_316_ext, l
	return nil
}
func (ϟa *GlGetProgramPipelineInfoLogEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_317_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_317_ext
	return nil
}
func (ϟa *GlGetProgramPipelineivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_318_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_318_ext
	return nil
}
func (ϟa *GlGetProgramResourcefvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_319_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_319_ext
	return nil
}
func (ϟa *GlGetQueryObjecti64vEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_320_ext := ExtensionId_GL_EXT_disjoint_timer_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_ = requiresExtension_320_ext
	return nil
}
func (ϟa *GlGetQueryObjectivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_321_ext := ExtensionId_GL_EXT_disjoint_timer_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_ = requiresExtension_321_ext
	return nil
}
func (ϟa *GlGetQueryObjectui64vEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_322_ext := ExtensionId_GL_EXT_disjoint_timer_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_ = requiresExtension_322_ext
	return nil
}
func (ϟa *GlGetQueryObjectuivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_323_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_324_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = requiresExtension_323_ext, requiresExtension_324_ext
	return nil
}
func (ϟa *GlGetQueryivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_325_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_326_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = requiresExtension_325_ext, requiresExtension_326_ext
	return nil
}
func (ϟa *GlGetSamplerParameterIivOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_327_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_327_ext
	return nil
}
func (ϟa *GlGetSamplerParameterIuivOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_328_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_328_ext
	return nil
}
func (ϟa *GlGetSyncivAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_329_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_329_ext
	return nil
}
func (ϟa *GlGetTexParameterIivOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_330_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_330_ext
	return nil
}
func (ϟa *GlGetTexParameterIuivOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_331_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_331_ext
	return nil
}
func (ϟa *GlGetTextureHandleNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_332_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_332_ext
	return nil
}
func (ϟa *GlGetTextureSamplerHandleNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_333_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_333_ext
	return nil
}
func (ϟa *GlGetTranslatedShaderSourceANGLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_334_ext := ExtensionId_GL_ANGLE_translated_shader_source // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_334_ext
	return nil
}
func (ϟa *GlGetnUniformfvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_335_ext := ExtensionId_GL_EXT_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_335_ext
	return nil
}
func (ϟa *GlGetnUniformfvKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_336_ext := ExtensionId_GL_KHR_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_336_ext
	return nil
}
func (ϟa *GlGetnUniformivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_337_ext := ExtensionId_GL_EXT_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_337_ext
	return nil
}
func (ϟa *GlGetnUniformivKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_338_ext := ExtensionId_GL_KHR_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_338_ext
	return nil
}
func (ϟa *GlGetnUniformuivKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_339_ext := ExtensionId_GL_KHR_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_339_ext
	return nil
}
func (ϟa *GlInsertEventMarkerEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_340_ext := ExtensionId_GL_EXT_debug_marker // ExtensionId
	if (ϟa.Length) > (GLsizei(int32(0))) {
		ϟa.Marker.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Length), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	} else {
		_ = strings.TrimRight(string(Charᵖ(ϟa.Marker).StringSlice(ϟs, ϟd, ϟl).Read(ϟa, ϟs, ϟd, ϟl, ϟb)), "\x00")
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_340_ext
	return nil
}
func (ϟa *GlInterpolatePathsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_341_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_341_ext
	return nil
}
func (ϟa *GlIsEnablediNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_342_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_342_ext
	return nil
}
func (ϟa *GlIsEnablediOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_343_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_343_ext
	return nil
}
func (ϟa *GlIsFenceNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_344_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_344_ext
	return nil
}
func (ϟa *GlIsImageHandleResidentNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_345_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_345_ext
	return nil
}
func (ϟa *GlIsPathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_346_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_346_ext
	return nil
}
func (ϟa *GlIsPointInFillPathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_347_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_347_ext
	return nil
}
func (ϟa *GlIsPointInStrokePathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_348_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_348_ext
	return nil
}
func (ϟa *GlIsProgramPipelineEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_349_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_349_ext
	return nil
}
func (ϟa *GlIsQueryEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_350_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_351_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                            // Contextʳ
	GetContext_352_result := context                                        // Contextʳ
	ctx := GetContext_352_result                                            // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.Queries.Contains(ϟa.Query) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = requiresExtension_350_ext, requiresExtension_351_ext, context, GetContext_352_result, ctx
	return nil
}
func (ϟa *GlIsSyncAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_353_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_353_ext
	return nil
}
func (ϟa *GlIsTextureHandleResidentNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_354_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_354_ext
	return nil
}
func (ϟa *GlIsVertexArrayOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_355_ext := ExtensionId_GL_OES_vertex_array_object // ExtensionId
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                        // Contextʳ
	GetContext_356_result := context                                    // Contextʳ
	ctx := GetContext_356_result                                        // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.VertexArrays.Contains(ϟa.Array) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _ = requiresExtension_355_ext, context, GetContext_356_result, ctx
	return nil
}
func (ϟa *GlLabelObjectEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_357_ext := ExtensionId_GL_EXT_debug_label // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_357_ext
	return nil
}
func (ϟa *GlMakeImageHandleNonResidentNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_358_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_358_ext
	return nil
}
func (ϟa *GlMakeImageHandleResidentNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_359_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_359_ext
	return nil
}
func (ϟa *GlMakeTextureHandleNonResidentNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_360_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_360_ext
	return nil
}
func (ϟa *GlMakeTextureHandleResidentNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_361_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_361_ext
	return nil
}
func (ϟa *GlMapBufferOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_362_ext := ExtensionId_GL_OES_mapbuffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_362_ext
	return nil
}
func (ϟa *GlMapBufferRangeEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_363_ext := ExtensionId_GL_EXT_map_buffer_range // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_363_ext
	return nil
}
func (ϟa *GlMatrixLoad3x2fNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_364_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_364_ext
	return nil
}
func (ϟa *GlMatrixLoad3x3fNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_365_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_365_ext
	return nil
}
func (ϟa *GlMatrixLoadTranspose3x3fNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_366_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_366_ext
	return nil
}
func (ϟa *GlMatrixMult3x2fNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_367_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_367_ext
	return nil
}
func (ϟa *GlMatrixMult3x3fNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_368_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_368_ext
	return nil
}
func (ϟa *GlMatrixMultTranspose3x3fNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_369_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_369_ext
	return nil
}
func (ϟa *GlMultiDrawArraysEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_370_ext := ExtensionId_GL_EXT_multi_draw_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_370_ext
	return nil
}
func (ϟa *GlMultiDrawArraysIndirectEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_371_ext := ExtensionId_GL_EXT_multi_draw_indirect // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_371_ext
	return nil
}
func (ϟa *GlMultiDrawElementsBaseVertexEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_372_ext := ExtensionId_GL_EXT_draw_elements_base_vertex // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_372_ext
	return nil
}
func (ϟa *GlMultiDrawElementsBaseVertexOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_373_ext := ExtensionId_GL_OES_draw_elements_base_vertex // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_373_ext
	return nil
}
func (ϟa *GlMultiDrawElementsEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_374_ext := ExtensionId_GL_EXT_multi_draw_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_374_ext
	return nil
}
func (ϟa *GlMultiDrawElementsIndirectEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_375_ext := ExtensionId_GL_EXT_multi_draw_indirect // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_375_ext
	return nil
}
func (ϟa *GlNamedFramebufferSampleLocationsfvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_376_ext := ExtensionId_GL_NV_sample_locations // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_376_ext
	return nil
}
func (ϟa *GlPatchParameteriOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_377_ext := ExtensionId_GL_OES_tessellation_shader // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_377_ext
	return nil
}
func (ϟa *GlPathCommandsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_378_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_378_ext
	return nil
}
func (ϟa *GlPathCoordsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_379_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_379_ext
	return nil
}
func (ϟa *GlPathCoverDepthFuncNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_380_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_380_ext
	return nil
}
func (ϟa *GlPathDashArrayNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_381_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_381_ext
	return nil
}
func (ϟa *GlPathGlyphIndexArrayNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_382_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_382_ext
	return nil
}
func (ϟa *GlPathGlyphIndexRangeNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_383_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_383_ext
	return nil
}
func (ϟa *GlPathGlyphRangeNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_384_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_384_ext
	return nil
}
func (ϟa *GlPathGlyphsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_385_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_385_ext
	return nil
}
func (ϟa *GlPathMemoryGlyphIndexArrayNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_386_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_386_ext
	return nil
}
func (ϟa *GlPathParameterfNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_387_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_387_ext
	return nil
}
func (ϟa *GlPathParameterfvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_388_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_388_ext
	return nil
}
func (ϟa *GlPathParameteriNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_389_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_389_ext
	return nil
}
func (ϟa *GlPathParameterivNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_390_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_390_ext
	return nil
}
func (ϟa *GlPathStencilDepthOffsetNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_391_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_391_ext
	return nil
}
func (ϟa *GlPathStencilFuncNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_392_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_392_ext
	return nil
}
func (ϟa *GlPathStringNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_393_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_393_ext
	return nil
}
func (ϟa *GlPathSubCommandsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_394_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_394_ext
	return nil
}
func (ϟa *GlPathSubCoordsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_395_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_395_ext
	return nil
}
func (ϟa *GlPointAlongPathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_396_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_396_ext
	return nil
}
func (ϟa *GlPolygonModeNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_397_ext := ExtensionId_GL_NV_polygon_mode // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_397_ext
	return nil
}
func (ϟa *GlPopGroupMarkerEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_398_ext := ExtensionId_GL_EXT_debug_marker // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_398_ext
	return nil
}
func (ϟa *GlPrimitiveBoundingBoxOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_399_ext := ExtensionId_GL_OES_primitive_bounding_box // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_399_ext
	return nil
}
func (ϟa *GlProgramBinaryOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_400_ext := ExtensionId_GL_OES_get_program_binary // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_400_ext
	return nil
}
func (ϟa *GlProgramParameteriEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_401_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_401_ext
	return nil
}
func (ϟa *GlProgramPathFragmentInputGenNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_402_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_402_ext
	return nil
}
func (ϟa *GlProgramUniform1fEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_403_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_403_ext
	return nil
}
func (ϟa *GlProgramUniform1fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_404_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_404_ext
	return nil
}
func (ϟa *GlProgramUniform1iEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_405_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_405_ext
	return nil
}
func (ϟa *GlProgramUniform1ivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_406_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_406_ext
	return nil
}
func (ϟa *GlProgramUniform1uiEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_407_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_407_ext
	return nil
}
func (ϟa *GlProgramUniform1uivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_408_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_408_ext
	return nil
}
func (ϟa *GlProgramUniform2fEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_409_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_409_ext
	return nil
}
func (ϟa *GlProgramUniform2fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_410_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_410_ext
	return nil
}
func (ϟa *GlProgramUniform2iEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_411_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_411_ext
	return nil
}
func (ϟa *GlProgramUniform2ivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_412_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_412_ext
	return nil
}
func (ϟa *GlProgramUniform2uiEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_413_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_413_ext
	return nil
}
func (ϟa *GlProgramUniform2uivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_414_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_414_ext
	return nil
}
func (ϟa *GlProgramUniform3fEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_415_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_415_ext
	return nil
}
func (ϟa *GlProgramUniform3fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_416_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_416_ext
	return nil
}
func (ϟa *GlProgramUniform3iEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_417_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_417_ext
	return nil
}
func (ϟa *GlProgramUniform3ivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_418_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_418_ext
	return nil
}
func (ϟa *GlProgramUniform3uiEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_419_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_419_ext
	return nil
}
func (ϟa *GlProgramUniform3uivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_420_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_420_ext
	return nil
}
func (ϟa *GlProgramUniform4fEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_421_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_421_ext
	return nil
}
func (ϟa *GlProgramUniform4fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_422_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_422_ext
	return nil
}
func (ϟa *GlProgramUniform4iEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_423_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_423_ext
	return nil
}
func (ϟa *GlProgramUniform4ivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_424_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_424_ext
	return nil
}
func (ϟa *GlProgramUniform4uiEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_425_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_425_ext
	return nil
}
func (ϟa *GlProgramUniform4uivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_426_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_426_ext
	return nil
}
func (ϟa *GlProgramUniformHandleui64NV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_427_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_427_ext
	return nil
}
func (ϟa *GlProgramUniformHandleui64vNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_428_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_428_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix2fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_429_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_429_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix2x3fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_430_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_430_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix2x4fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_431_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_431_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix3fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_432_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_432_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix3x2fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_433_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_433_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix3x4fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_434_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_434_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix4fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_435_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_435_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix4x2fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_436_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_436_ext
	return nil
}
func (ϟa *GlProgramUniformMatrix4x3fvEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_437_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_437_ext
	return nil
}
func (ϟa *GlPushGroupMarkerEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_438_ext := ExtensionId_GL_EXT_debug_marker // ExtensionId
	if (ϟa.Length) > (GLsizei(int32(0))) {
		ϟa.Marker.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Length), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	} else {
		_ = strings.TrimRight(string(Charᵖ(ϟa.Marker).StringSlice(ϟs, ϟd, ϟl).Read(ϟa, ϟs, ϟd, ϟl, ϟb)), "\x00")
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_438_ext
	return nil
}
func (ϟa *GlQueryCounterEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_439_ext := ExtensionId_GL_EXT_disjoint_timer_query // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_439_ext
	return nil
}
func (ϟa *GlRasterSamplesEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_440_ext := ExtensionId_GL_EXT_raster_multisample       // ExtensionId
	requiresExtension_441_ext := ExtensionId_GL_EXT_texture_filter_minmax    // ExtensionId
	requiresExtension_442_ext := ExtensionId_GL_NV_framebuffer_mixed_samples // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = requiresExtension_440_ext, requiresExtension_441_ext, requiresExtension_442_ext
	return nil
}
func (ϟa *GlReadBufferIndexedEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_443_ext := ExtensionId_GL_EXT_multiview_draw_buffers // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_443_ext
	return nil
}
func (ϟa *GlReadBufferNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_444_ext := ExtensionId_GL_NV_read_buffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_444_ext
	return nil
}
func (ϟa *GlReadnPixelsEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_445_ext := ExtensionId_GL_EXT_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_445_ext
	return nil
}
func (ϟa *GlReadnPixelsKHR) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_446_ext := ExtensionId_GL_KHR_robustness // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_446_ext
	return nil
}
func (ϟa *GlRenderbufferStorageMultisampleANGLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_447_ext := ExtensionId_GL_ANGLE_framebuffer_multisample // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_447_ext
	return nil
}
func (ϟa *GlRenderbufferStorageMultisampleAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_448_ext := ExtensionId_GL_APPLE_framebuffer_multisample // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_448_ext
	return nil
}
func (ϟa *GlRenderbufferStorageMultisampleEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_449_ext := ExtensionId_GL_EXT_multisampled_render_to_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_449_ext
	return nil
}
func (ϟa *GlRenderbufferStorageMultisampleIMG) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_450_ext := ExtensionId_GL_IMG_multisampled_render_to_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_450_ext
	return nil
}
func (ϟa *GlRenderbufferStorageMultisampleNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_451_ext := ExtensionId_GL_NV_framebuffer_multisample // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_451_ext
	return nil
}
func (ϟa *GlResolveDepthValuesNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_452_ext := ExtensionId_GL_NV_sample_locations // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_452_ext
	return nil
}
func (ϟa *GlResolveMultisampleFramebufferAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_453_ext := ExtensionId_GL_APPLE_framebuffer_multisample // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_453_ext
	return nil
}
func (ϟa *GlSamplerParameterIivOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_454_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_454_ext
	return nil
}
func (ϟa *GlSamplerParameterIuivOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_455_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_455_ext
	return nil
}
func (ϟa *GlScissorArrayvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_456_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_456_ext
	return nil
}
func (ϟa *GlScissorIndexedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_457_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_457_ext
	return nil
}
func (ϟa *GlScissorIndexedvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_458_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_458_ext
	return nil
}
func (ϟa *GlSelectPerfMonitorCountersAMD) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_459_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_459_ext
	return nil
}
func (ϟa *GlSetFenceNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_460_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_460_ext
	return nil
}
func (ϟa *GlStartTilingQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_461_ext := ExtensionId_GL_QCOM_tiled_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_461_ext
	return nil
}
func (ϟa *GlStencilFillPathInstancedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_462_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_462_ext
	return nil
}
func (ϟa *GlStencilFillPathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_463_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_463_ext
	return nil
}
func (ϟa *GlStencilStrokePathInstancedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_464_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_464_ext
	return nil
}
func (ϟa *GlStencilStrokePathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_465_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_465_ext
	return nil
}
func (ϟa *GlStencilThenCoverFillPathInstancedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_466_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_466_ext
	return nil
}
func (ϟa *GlStencilThenCoverFillPathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_467_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_467_ext
	return nil
}
func (ϟa *GlStencilThenCoverStrokePathInstancedNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_468_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_468_ext
	return nil
}
func (ϟa *GlStencilThenCoverStrokePathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_469_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_469_ext
	return nil
}
func (ϟa *GlSubpixelPrecisionBiasNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_470_ext := ExtensionId_GL_NV_conservative_raster // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_470_ext
	return nil
}
func (ϟa *GlTestFenceNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_471_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_471_ext
	return nil
}
func (ϟa *GlTexBufferOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_472_ext := ExtensionId_GL_OES_texture_buffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_472_ext
	return nil
}
func (ϟa *GlTexBufferRangeOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_473_ext := ExtensionId_GL_OES_texture_buffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_473_ext
	return nil
}
func (ϟa *GlTexImage3DOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_474_ext := ExtensionId_GL_OES_texture_3D // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_474_ext
	return nil
}
func (ϟa *GlTexPageCommitmentARB) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_475_ext := ExtensionId_GL_EXT_sparse_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_475_ext
	return nil
}
func (ϟa *GlTexParameterIivOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_476_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_476_ext
	return nil
}
func (ϟa *GlTexParameterIuivOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_477_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_477_ext
	return nil
}
func (ϟa *GlTexStorage1DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_478_ext := ExtensionId_GL_EXT_texture_storage // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_478_ext
	return nil
}
func (ϟa *GlTexStorage2DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_479_ext := ExtensionId_GL_EXT_texture_storage // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_479_ext
	return nil
}
func (ϟa *GlTexStorage3DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_480_ext := ExtensionId_GL_EXT_texture_storage // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_480_ext
	return nil
}
func (ϟa *GlTexSubImage3DOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_481_ext := ExtensionId_GL_OES_texture_3D // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_481_ext
	return nil
}
func (ϟa *GlTextureStorage1DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_482_ext := ExtensionId_GL_EXT_texture_storage // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_482_ext
	return nil
}
func (ϟa *GlTextureStorage2DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_483_ext := ExtensionId_GL_EXT_texture_storage // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_483_ext
	return nil
}
func (ϟa *GlTextureStorage3DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_484_ext := ExtensionId_GL_EXT_texture_storage // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_484_ext
	return nil
}
func (ϟa *GlTextureViewEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_485_ext := ExtensionId_GL_EXT_texture_view // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_485_ext
	return nil
}
func (ϟa *GlTextureViewOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_486_ext := ExtensionId_GL_OES_texture_view // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_486_ext
	return nil
}
func (ϟa *GlTransformPathNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_487_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_487_ext
	return nil
}
func (ϟa *GlUniformHandleui64NV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_488_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_488_ext
	return nil
}
func (ϟa *GlUniformHandleui64vNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_489_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_489_ext
	return nil
}
func (ϟa *GlUniformMatrix2x3fvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_490_ext := ExtensionId_GL_NV_non_square_matrices // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_490_ext
	return nil
}
func (ϟa *GlUniformMatrix2x4fvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_491_ext := ExtensionId_GL_NV_non_square_matrices // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_491_ext
	return nil
}
func (ϟa *GlUniformMatrix3x2fvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_492_ext := ExtensionId_GL_NV_non_square_matrices // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_492_ext
	return nil
}
func (ϟa *GlUniformMatrix3x4fvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_493_ext := ExtensionId_GL_NV_non_square_matrices // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_493_ext
	return nil
}
func (ϟa *GlUniformMatrix4x2fvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_494_ext := ExtensionId_GL_NV_non_square_matrices // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_494_ext
	return nil
}
func (ϟa *GlUniformMatrix4x3fvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_495_ext := ExtensionId_GL_NV_non_square_matrices // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_495_ext
	return nil
}
func (ϟa *GlUnmapBufferOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_496_ext := ExtensionId_GL_OES_mapbuffer // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = requiresExtension_496_ext
	return nil
}
func (ϟa *GlUseProgramStagesEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_497_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_497_ext
	return nil
}
func (ϟa *GlValidateProgramPipelineEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_498_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_498_ext
	return nil
}
func (ϟa *GlVertexAttribDivisorANGLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_499_ext := ExtensionId_GL_ANGLE_instanced_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_499_ext
	return nil
}
func (ϟa *GlVertexAttribDivisorEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_500_ext := ExtensionId_GL_EXT_instanced_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_500_ext
	return nil
}
func (ϟa *GlVertexAttribDivisorNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_501_ext := ExtensionId_GL_NV_instanced_arrays // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_501_ext
	return nil
}
func (ϟa *GlViewportArrayvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_502_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_502_ext
	return nil
}
func (ϟa *GlViewportIndexedfNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_503_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_503_ext
	return nil
}
func (ϟa *GlViewportIndexedfvNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_504_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_504_ext
	return nil
}
func (ϟa *GlWaitSyncAPPLE) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_505_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_505_ext
	return nil
}
func (ϟa *GlWeightPathsNV) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_506_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_506_ext
	return nil
}
func (ϟa *GlBlendBarrier) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_507_major := uint32(3) // u32
	minRequiredVersion_507_minor := uint32(2) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_507_major, minRequiredVersion_507_minor
	return nil
}
func (ϟa *GlBlendColor) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_508_major := uint32(2)    // u32
	minRequiredVersion_508_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_509_result := context             // Contextʳ
	ctx := GetContext_509_result                 // Contextʳ
	ctx.Blending.BlendColor = Color{Red: ϟa.Red, Green: ϟa.Green, Blue: ϟa.Blue, Alpha: ϟa.Alpha}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_508_major, minRequiredVersion_508_minor, context, GetContext_509_result, ctx
	return nil
}
func (ϟa *GlBlendEquation) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_510_major := uint32(2) // u32
	minRequiredVersion_510_minor := uint32(0) // u32
	switch ϟa.Equation {
	case GLenum_GL_FUNC_ADD, GLenum_GL_FUNC_REVERSE_SUBTRACT, GLenum_GL_FUNC_SUBTRACT:
	case GLenum_GL_MAX, GLenum_GL_MIN:
		minRequiredVersion_511_major := uint32(3) // u32
		minRequiredVersion_511_minor := uint32(0) // u32
		_, _ = minRequiredVersion_511_major, minRequiredVersion_511_minor
	default:
		v := ϟa.Equation
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_513_result := context             // Contextʳ
	ctx := GetContext_513_result                 // Contextʳ
	ctx.Blending.BlendEquationRgb = ϟa.Equation
	ctx.Blending.BlendEquationAlpha = ϟa.Equation
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_510_major, minRequiredVersion_510_minor, context, GetContext_513_result, ctx
	return nil
}
func (ϟa *GlBlendEquationSeparate) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_514_major := uint32(2) // u32
	minRequiredVersion_514_minor := uint32(0) // u32
	switch ϟa.Rgb {
	case GLenum_GL_FUNC_ADD, GLenum_GL_FUNC_REVERSE_SUBTRACT, GLenum_GL_FUNC_SUBTRACT:
	case GLenum_GL_MAX, GLenum_GL_MIN:
		minRequiredVersion_515_major := uint32(3) // u32
		minRequiredVersion_515_minor := uint32(0) // u32
		_, _ = minRequiredVersion_515_major, minRequiredVersion_515_minor
	default:
		v := ϟa.Rgb
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Alpha {
	case GLenum_GL_FUNC_ADD, GLenum_GL_FUNC_REVERSE_SUBTRACT, GLenum_GL_FUNC_SUBTRACT:
	case GLenum_GL_MAX, GLenum_GL_MIN:
		minRequiredVersion_517_major := uint32(3) // u32
		minRequiredVersion_517_minor := uint32(0) // u32
		_, _ = minRequiredVersion_517_major, minRequiredVersion_517_minor
	default:
		v := ϟa.Alpha
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_519_result := context             // Contextʳ
	ctx := GetContext_519_result                 // Contextʳ
	ctx.Blending.BlendEquationRgb = ϟa.Rgb
	ctx.Blending.BlendEquationAlpha = ϟa.Alpha
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_514_major, minRequiredVersion_514_minor, context, GetContext_519_result, ctx
	return nil
}
func (ϟa *GlBlendEquationSeparatei) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_520_major := uint32(3) // u32
	minRequiredVersion_520_minor := uint32(2) // u32
	switch ϟa.ModeRGB {
	case GLenum_GL_FUNC_ADD, GLenum_GL_FUNC_REVERSE_SUBTRACT, GLenum_GL_FUNC_SUBTRACT, GLenum_GL_MAX, GLenum_GL_MIN:
	default:
		v := ϟa.ModeRGB
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.ModeAlpha {
	case GLenum_GL_FUNC_ADD, GLenum_GL_FUNC_REVERSE_SUBTRACT, GLenum_GL_FUNC_SUBTRACT, GLenum_GL_MAX, GLenum_GL_MIN:
	default:
		v := ϟa.ModeAlpha
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_520_major, minRequiredVersion_520_minor
	return nil
}
func (ϟa *GlBlendEquationi) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_523_major := uint32(3) // u32
	minRequiredVersion_523_minor := uint32(2) // u32
	switch ϟa.Mode {
	case GLenum_GL_FUNC_ADD, GLenum_GL_FUNC_REVERSE_SUBTRACT, GLenum_GL_FUNC_SUBTRACT, GLenum_GL_MAX, GLenum_GL_MIN:
	default:
		v := ϟa.Mode
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_523_major, minRequiredVersion_523_minor
	return nil
}
func (ϟa *GlBlendFunc) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_525_major := uint32(2) // u32
	minRequiredVersion_525_minor := uint32(0) // u32
	switch ϟa.SrcFactor {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		v := ϟa.SrcFactor
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.DstFactor {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		v := ϟa.DstFactor
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_528_result := context             // Contextʳ
	ctx := GetContext_528_result                 // Contextʳ
	ctx.Blending.SrcRgbBlendFactor = ϟa.SrcFactor
	ctx.Blending.SrcAlphaBlendFactor = ϟa.SrcFactor
	ctx.Blending.DstRgbBlendFactor = ϟa.DstFactor
	ctx.Blending.DstAlphaBlendFactor = ϟa.DstFactor
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_525_major, minRequiredVersion_525_minor, context, GetContext_528_result, ctx
	return nil
}
func (ϟa *GlBlendFuncSeparate) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_529_major := uint32(2) // u32
	minRequiredVersion_529_minor := uint32(0) // u32
	switch ϟa.SrcFactorRgb {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		v := ϟa.SrcFactorRgb
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.DstFactorRgb {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
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
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		v := ϟa.DstFactorAlpha
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_534_result := context             // Contextʳ
	ctx := GetContext_534_result                 // Contextʳ
	ctx.Blending.SrcRgbBlendFactor = ϟa.SrcFactorRgb
	ctx.Blending.DstRgbBlendFactor = ϟa.DstFactorRgb
	ctx.Blending.SrcAlphaBlendFactor = ϟa.SrcFactorAlpha
	ctx.Blending.DstAlphaBlendFactor = ϟa.DstFactorAlpha
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_529_major, minRequiredVersion_529_minor, context, GetContext_534_result, ctx
	return nil
}
func (ϟa *GlBlendFuncSeparatei) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_535_major := uint32(3) // u32
	minRequiredVersion_535_minor := uint32(2) // u32
	switch ϟa.SrcRGB {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		v := ϟa.SrcRGB
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.DstRGB {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		v := ϟa.DstRGB
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.SrcAlpha {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		v := ϟa.SrcAlpha
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.DstAlpha {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		v := ϟa.DstAlpha
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_535_major, minRequiredVersion_535_minor
	return nil
}
func (ϟa *GlBlendFunci) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_540_major := uint32(3) // u32
	minRequiredVersion_540_minor := uint32(2) // u32
	switch ϟa.Src {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		v := ϟa.Src
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Dst {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		v := ϟa.Dst
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_540_major, minRequiredVersion_540_minor
	return nil
}
func (ϟa *GlDepthFunc) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_543_major := uint32(2) // u32
	minRequiredVersion_543_minor := uint32(0) // u32
	switch ϟa.Function {
	case GLenum_GL_ALWAYS, GLenum_GL_EQUAL, GLenum_GL_GEQUAL, GLenum_GL_GREATER, GLenum_GL_LEQUAL, GLenum_GL_LESS, GLenum_GL_NEVER, GLenum_GL_NOTEQUAL:
	default:
		v := ϟa.Function
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_545_result := context             // Contextʳ
	ctx := GetContext_545_result                 // Contextʳ
	ctx.Rasterizing.DepthTestFunction = ϟa.Function
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_543_major, minRequiredVersion_543_minor, context, GetContext_545_result, ctx
	return nil
}
func (ϟa *GlSampleCoverage) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_546_major := uint32(2)    // u32
	minRequiredVersion_546_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_547_result := context             // Contextʳ
	ctx := GetContext_547_result                 // Contextʳ
	ctx.Rasterizing.SampleCoverageValue = ϟa.Value
	ctx.Rasterizing.SampleCoverageInvert = ϟa.Invert
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_546_major, minRequiredVersion_546_minor, context, GetContext_547_result, ctx
	return nil
}
func (ϟa *GlSampleMaski) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_548_major := uint32(3) // u32
	minRequiredVersion_548_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_548_major, minRequiredVersion_548_minor
	return nil
}
func (ϟa *GlScissor) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_549_major := uint32(2)    // u32
	minRequiredVersion_549_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_550_result := context             // Contextʳ
	ctx := GetContext_550_result                 // Contextʳ
	ctx.Rasterizing.Scissor = Rect{X: ϟa.X, Y: ϟa.Y, Width: ϟa.Width, Height: ϟa.Height}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_549_major, minRequiredVersion_549_minor, context, GetContext_550_result, ctx
	return nil
}
func (ϟa *GlStencilFunc) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_551_major := uint32(2) // u32
	minRequiredVersion_551_minor := uint32(0) // u32
	switch ϟa.Func {
	case GLenum_GL_ALWAYS, GLenum_GL_EQUAL, GLenum_GL_GEQUAL, GLenum_GL_GREATER, GLenum_GL_LEQUAL, GLenum_GL_LESS, GLenum_GL_NEVER, GLenum_GL_NOTEQUAL:
	default:
		v := ϟa.Func
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_551_major, minRequiredVersion_551_minor
	return nil
}
func (ϟa *GlStencilFuncSeparate) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_553_major := uint32(2) // u32
	minRequiredVersion_553_minor := uint32(0) // u32
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
	_, _ = minRequiredVersion_553_major, minRequiredVersion_553_minor
	return nil
}
func (ϟa *GlStencilOp) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_556_major := uint32(2) // u32
	minRequiredVersion_556_minor := uint32(0) // u32
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
	_, _ = minRequiredVersion_556_major, minRequiredVersion_556_minor
	return nil
}
func (ϟa *GlStencilOpSeparate) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_560_major := uint32(2) // u32
	minRequiredVersion_560_minor := uint32(0) // u32
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
	_, _ = minRequiredVersion_560_major, minRequiredVersion_560_minor
	return nil
}
func (ϟa *GlBindFramebuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_565_major := uint32(2) // u32
	minRequiredVersion_565_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_FRAMEBUFFER:
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
		minRequiredVersion_566_major := uint32(3) // u32
		minRequiredVersion_566_minor := uint32(0) // u32
		_, _ = minRequiredVersion_566_major, minRequiredVersion_566_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_568_result := context             // Contextʳ
	ctx := GetContext_568_result                 // Contextʳ
	if !(ctx.Instances.Framebuffers.Contains(ϟa.Framebuffer)) {
		ctx.Instances.Framebuffers[ϟa.Framebuffer] = &Framebuffer{Attachments: GLenumːFramebufferAttachmentInfoᵐ{}}
	}
	if (ϟa.Target) == (GLenum_GL_FRAMEBUFFER) {
		ctx.BoundFramebuffers[GLenum_GL_READ_FRAMEBUFFER] = ϟa.Framebuffer
		ctx.BoundFramebuffers[GLenum_GL_DRAW_FRAMEBUFFER] = ϟa.Framebuffer
	} else {
		ctx.BoundFramebuffers[ϟa.Target] = ϟa.Framebuffer
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_565_major, minRequiredVersion_565_minor, context, GetContext_568_result, ctx
	return nil
}
func (ϟa *GlBindRenderbuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_569_major := uint32(2) // u32
	minRequiredVersion_569_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_RENDERBUFFER:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_571_result := context             // Contextʳ
	ctx := GetContext_571_result                 // Contextʳ
	if !(ctx.Instances.Renderbuffers.Contains(ϟa.Renderbuffer)) {
		ctx.Instances.Renderbuffers[ϟa.Renderbuffer] = &Renderbuffer{}
	}
	ctx.BoundRenderbuffers[ϟa.Target] = ϟa.Renderbuffer
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_569_major, minRequiredVersion_569_minor, context, GetContext_571_result, ctx
	return nil
}
func (ϟa *GlBlitFramebuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_572_major := uint32(3)                                                                                                // u32
	minRequiredVersion_572_minor := uint32(0)                                                                                                // u32
	supportsBits_573_seenBits := ϟa.Mask                                                                                                     // GLbitfield
	supportsBits_573_validBits := (GLbitfield_GL_COLOR_BUFFER_BIT) | ((GLbitfield_GL_DEPTH_BUFFER_BIT) | (GLbitfield_GL_STENCIL_BUFFER_BIT)) // GLbitfield
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
	_, _, _, _ = minRequiredVersion_572_major, minRequiredVersion_572_minor, supportsBits_573_seenBits, supportsBits_573_validBits
	return nil
}
func (ϟa *GlCheckFramebufferStatus) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_575_major := uint32(2) // u32
	minRequiredVersion_575_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_FRAMEBUFFER:
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
		minRequiredVersion_576_major := uint32(3) // u32
		minRequiredVersion_576_minor := uint32(0) // u32
		_, _ = minRequiredVersion_576_major, minRequiredVersion_576_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_575_major, minRequiredVersion_575_minor
	return nil
}
func (ϟa *GlClear) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_578_major := uint32(2)                                                                                                // u32
	minRequiredVersion_578_minor := uint32(0)                                                                                                // u32
	supportsBits_579_seenBits := ϟa.Mask                                                                                                     // GLbitfield
	supportsBits_579_validBits := (GLbitfield_GL_COLOR_BUFFER_BIT) | ((GLbitfield_GL_DEPTH_BUFFER_BIT) | (GLbitfield_GL_STENCIL_BUFFER_BIT)) // GLbitfield
	if (GLbitfield_GL_COLOR_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	if (GLbitfield_GL_DEPTH_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	if (GLbitfield_GL_STENCIL_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	if (GLbitfield_GL_COLOR_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = minRequiredVersion_578_major, minRequiredVersion_578_minor, supportsBits_579_seenBits, supportsBits_579_validBits
	return nil
}
func (ϟa *GlClearBufferfi) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_580_major := uint32(3) // u32
	minRequiredVersion_580_minor := uint32(0) // u32
	switch ϟa.Buffer {
	case GLenum_GL_DEPTH_STENCIL:
	default:
		v := ϟa.Buffer
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_580_major, minRequiredVersion_580_minor
	return nil
}
func (ϟa *GlClearBufferfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_582_major := uint32(3) // u32
	minRequiredVersion_582_minor := uint32(0) // u32
	switch ϟa.Buffer {
	case GLenum_GL_COLOR, GLenum_GL_DEPTH:
	default:
		v := ϟa.Buffer
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_582_major, minRequiredVersion_582_minor
	return nil
}
func (ϟa *GlClearBufferiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_584_major := uint32(3) // u32
	minRequiredVersion_584_minor := uint32(0) // u32
	switch ϟa.Buffer {
	case GLenum_GL_COLOR, GLenum_GL_STENCIL:
	default:
		v := ϟa.Buffer
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_584_major, minRequiredVersion_584_minor
	return nil
}
func (ϟa *GlClearBufferuiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_586_major := uint32(3) // u32
	minRequiredVersion_586_minor := uint32(0) // u32
	switch ϟa.Buffer {
	case GLenum_GL_COLOR:
	default:
		v := ϟa.Buffer
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_586_major, minRequiredVersion_586_minor
	return nil
}
func (ϟa *GlClearColor) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_588_major := uint32(2)    // u32
	minRequiredVersion_588_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_589_result := context             // Contextʳ
	ctx := GetContext_589_result                 // Contextʳ
	ctx.Clearing.ClearColor = Color{Red: ϟa.R, Green: ϟa.G, Blue: ϟa.B, Alpha: ϟa.A}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_588_major, minRequiredVersion_588_minor, context, GetContext_589_result, ctx
	return nil
}
func (ϟa *GlClearDepthf) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_590_major := uint32(2)    // u32
	minRequiredVersion_590_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_591_result := context             // Contextʳ
	ctx := GetContext_591_result                 // Contextʳ
	ctx.Clearing.ClearDepth = ϟa.Depth
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_590_major, minRequiredVersion_590_minor, context, GetContext_591_result, ctx
	return nil
}
func (ϟa *GlClearStencil) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_592_major := uint32(2)    // u32
	minRequiredVersion_592_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_593_result := context             // Contextʳ
	ctx := GetContext_593_result                 // Contextʳ
	ctx.Clearing.ClearStencil = ϟa.Stencil
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_592_major, minRequiredVersion_592_minor, context, GetContext_593_result, ctx
	return nil
}
func (ϟa *GlColorMask) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_594_major := uint32(2)    // u32
	minRequiredVersion_594_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_595_result := context             // Contextʳ
	ctx := GetContext_595_result                 // Contextʳ
	ctx.Rasterizing.ColorMaskRed = ϟa.Red
	ctx.Rasterizing.ColorMaskGreen = ϟa.Green
	ctx.Rasterizing.ColorMaskBlue = ϟa.Blue
	ctx.Rasterizing.ColorMaskAlpha = ϟa.Alpha
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_594_major, minRequiredVersion_594_minor, context, GetContext_595_result, ctx
	return nil
}
func (ϟa *GlColorMaski) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_596_major := uint32(3) // u32
	minRequiredVersion_596_minor := uint32(2) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_596_major, minRequiredVersion_596_minor
	return nil
}
func (ϟa *GlDeleteFramebuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_597_major := uint32(2)                                   // u32
	minRequiredVersion_597_minor := uint32(0)                                   // u32
	f := ϟa.Framebuffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // FramebufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                // Contextʳ
	GetContext_598_result := context                                            // Contextʳ
	ctx := GetContext_598_result                                                // Contextʳ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.Framebuffers, f.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_597_major, minRequiredVersion_597_minor, f, context, GetContext_598_result, ctx
	return nil
}
func (ϟa *GlDeleteRenderbuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_599_major := uint32(2)                                    // u32
	minRequiredVersion_599_minor := uint32(0)                                    // u32
	r := ϟa.Renderbuffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // RenderbufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                 // Contextʳ
	GetContext_600_result := context                                             // Contextʳ
	ctx := GetContext_600_result                                                 // Contextʳ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.Renderbuffers, r.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_599_major, minRequiredVersion_599_minor, r, context, GetContext_600_result, ctx
	return nil
}
func (ϟa *GlDepthMask) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_601_major := uint32(2)    // u32
	minRequiredVersion_601_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_602_result := context             // Contextʳ
	ctx := GetContext_602_result                 // Contextʳ
	ctx.Rasterizing.DepthMask = ϟa.Enabled
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_601_major, minRequiredVersion_601_minor, context, GetContext_602_result, ctx
	return nil
}
func (ϟa *GlFramebufferParameteri) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_603_major := uint32(3) // u32
	minRequiredVersion_603_minor := uint32(1) // u32
	switch ϟa.Target {
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Pname {
	case GLenum_GL_FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS, GLenum_GL_FRAMEBUFFER_DEFAULT_HEIGHT, GLenum_GL_FRAMEBUFFER_DEFAULT_SAMPLES, GLenum_GL_FRAMEBUFFER_DEFAULT_WIDTH:
	case GLenum_GL_FRAMEBUFFER_DEFAULT_LAYERS:
		minRequiredVersion_605_major := uint32(3) // u32
		minRequiredVersion_605_minor := uint32(2) // u32
		_, _ = minRequiredVersion_605_major, minRequiredVersion_605_minor
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_603_major, minRequiredVersion_603_minor
	return nil
}
func (ϟa *GlFramebufferRenderbuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_607_major := uint32(2) // u32
	minRequiredVersion_607_minor := uint32(0) // u32
	switch ϟa.FramebufferTarget {
	case GLenum_GL_FRAMEBUFFER:
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
		minRequiredVersion_608_major := uint32(3) // u32
		minRequiredVersion_608_minor := uint32(0) // u32
		_, _ = minRequiredVersion_608_major, minRequiredVersion_608_minor
	default:
		v := ϟa.FramebufferTarget
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.FramebufferAttachment {
	case GLenum_GL_COLOR_ATTACHMENT0, GLenum_GL_DEPTH_ATTACHMENT, GLenum_GL_STENCIL_ATTACHMENT:
	case GLenum_GL_COLOR_ATTACHMENT1, GLenum_GL_COLOR_ATTACHMENT10, GLenum_GL_COLOR_ATTACHMENT11, GLenum_GL_COLOR_ATTACHMENT12, GLenum_GL_COLOR_ATTACHMENT13, GLenum_GL_COLOR_ATTACHMENT14, GLenum_GL_COLOR_ATTACHMENT15, GLenum_GL_COLOR_ATTACHMENT2, GLenum_GL_COLOR_ATTACHMENT3, GLenum_GL_COLOR_ATTACHMENT4, GLenum_GL_COLOR_ATTACHMENT5, GLenum_GL_COLOR_ATTACHMENT6, GLenum_GL_COLOR_ATTACHMENT7, GLenum_GL_COLOR_ATTACHMENT8, GLenum_GL_COLOR_ATTACHMENT9, GLenum_GL_DEPTH_STENCIL_ATTACHMENT:
		minRequiredVersion_610_major := uint32(3) // u32
		minRequiredVersion_610_minor := uint32(0) // u32
		_, _ = minRequiredVersion_610_major, minRequiredVersion_610_minor
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
	GetContext_613_result := context             // Contextʳ
	ctx := GetContext_613_result                 // Contextʳ
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
	_, _, _, _, _, _, _, _, _ = minRequiredVersion_607_major, minRequiredVersion_607_minor, context, GetContext_613_result, ctx, target, framebufferId, framebuffer, attachment
	return nil
}
func (ϟa *GlFramebufferTexture) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_614_major := uint32(3) // u32
	minRequiredVersion_614_minor := uint32(2) // u32
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
	_, _ = minRequiredVersion_614_major, minRequiredVersion_614_minor
	return nil
}
func (ϟa *GlFramebufferTexture2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_617_major := uint32(2) // u32
	minRequiredVersion_617_minor := uint32(0) // u32
	switch ϟa.FramebufferTarget {
	case GLenum_GL_FRAMEBUFFER:
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
		minRequiredVersion_618_major := uint32(3) // u32
		minRequiredVersion_618_minor := uint32(0) // u32
		_, _ = minRequiredVersion_618_major, minRequiredVersion_618_minor
	default:
		v := ϟa.FramebufferTarget
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.FramebufferAttachment {
	case GLenum_GL_COLOR_ATTACHMENT0, GLenum_GL_DEPTH_ATTACHMENT, GLenum_GL_STENCIL_ATTACHMENT:
	case GLenum_GL_COLOR_ATTACHMENT1, GLenum_GL_COLOR_ATTACHMENT10, GLenum_GL_COLOR_ATTACHMENT11, GLenum_GL_COLOR_ATTACHMENT12, GLenum_GL_COLOR_ATTACHMENT13, GLenum_GL_COLOR_ATTACHMENT14, GLenum_GL_COLOR_ATTACHMENT15, GLenum_GL_COLOR_ATTACHMENT2, GLenum_GL_COLOR_ATTACHMENT3, GLenum_GL_COLOR_ATTACHMENT4, GLenum_GL_COLOR_ATTACHMENT5, GLenum_GL_COLOR_ATTACHMENT6, GLenum_GL_COLOR_ATTACHMENT7, GLenum_GL_COLOR_ATTACHMENT8, GLenum_GL_COLOR_ATTACHMENT9, GLenum_GL_DEPTH_STENCIL_ATTACHMENT:
		minRequiredVersion_620_major := uint32(3) // u32
		minRequiredVersion_620_minor := uint32(0) // u32
		_, _ = minRequiredVersion_620_major, minRequiredVersion_620_minor
	default:
		v := ϟa.FramebufferAttachment
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.TextureTarget {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_622_major := uint32(3) // u32
		minRequiredVersion_622_minor := uint32(1) // u32
		_, _ = minRequiredVersion_622_major, minRequiredVersion_622_minor
	default:
		v := ϟa.TextureTarget
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_624_result := context             // Contextʳ
	ctx := GetContext_624_result                 // Contextʳ
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
	_, _, _, _, _, _, _, _, _ = minRequiredVersion_617_major, minRequiredVersion_617_minor, context, GetContext_624_result, ctx, target, framebufferId, framebuffer, attachment
	return nil
}
func (ϟa *GlFramebufferTextureLayer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_625_major := uint32(3) // u32
	minRequiredVersion_625_minor := uint32(0) // u32
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
	_, _ = minRequiredVersion_625_major, minRequiredVersion_625_minor
	return nil
}
func (ϟa *GlGenFramebuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_628_major := uint32(2)                                   // u32
	minRequiredVersion_628_minor := uint32(0)                                   // u32
	f := ϟa.Framebuffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // FramebufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                // Contextʳ
	GetContext_629_result := context                                            // Contextʳ
	ctx := GetContext_629_result                                                // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := FramebufferId(ϟa.Framebuffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // FramebufferId
		ctx.Instances.Framebuffers[id] = &Framebuffer{Attachments: GLenumːFramebufferAttachmentInfoᵐ{}}
		f.Index(uint64(i), ϟs).Write(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = minRequiredVersion_628_major, minRequiredVersion_628_minor, f, context, GetContext_629_result, ctx
	return nil
}
func (ϟa *GlGenRenderbuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_630_major := uint32(2)                                    // u32
	minRequiredVersion_630_minor := uint32(0)                                    // u32
	r := ϟa.Renderbuffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // RenderbufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                 // Contextʳ
	GetContext_631_result := context                                             // Contextʳ
	ctx := GetContext_631_result                                                 // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := RenderbufferId(ϟa.Renderbuffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // RenderbufferId
		ctx.Instances.Renderbuffers[id] = &Renderbuffer{}
		r.Index(uint64(i), ϟs).Write(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = minRequiredVersion_630_major, minRequiredVersion_630_minor, r, context, GetContext_631_result, ctx
	return nil
}
func (ϟa *GlGetFramebufferAttachmentParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_632_major := uint32(2) // u32
	minRequiredVersion_632_minor := uint32(0) // u32
	switch ϟa.FramebufferTarget {
	case GLenum_GL_FRAMEBUFFER:
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
		minRequiredVersion_633_major := uint32(3) // u32
		minRequiredVersion_633_minor := uint32(0) // u32
		_, _ = minRequiredVersion_633_major, minRequiredVersion_633_minor
	default:
		v := ϟa.FramebufferTarget
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Attachment {
	case GLenum_GL_COLOR_ATTACHMENT0, GLenum_GL_DEPTH_ATTACHMENT, GLenum_GL_STENCIL_ATTACHMENT:
	case GLenum_GL_BACK, GLenum_GL_COLOR_ATTACHMENT1, GLenum_GL_COLOR_ATTACHMENT10, GLenum_GL_COLOR_ATTACHMENT11, GLenum_GL_COLOR_ATTACHMENT12, GLenum_GL_COLOR_ATTACHMENT13, GLenum_GL_COLOR_ATTACHMENT14, GLenum_GL_COLOR_ATTACHMENT15, GLenum_GL_COLOR_ATTACHMENT2, GLenum_GL_COLOR_ATTACHMENT3, GLenum_GL_COLOR_ATTACHMENT4, GLenum_GL_COLOR_ATTACHMENT5, GLenum_GL_COLOR_ATTACHMENT6, GLenum_GL_COLOR_ATTACHMENT7, GLenum_GL_COLOR_ATTACHMENT8, GLenum_GL_COLOR_ATTACHMENT9, GLenum_GL_DEPTH, GLenum_GL_DEPTH_STENCIL_ATTACHMENT, GLenum_GL_STENCIL:
		minRequiredVersion_635_major := uint32(3) // u32
		minRequiredVersion_635_minor := uint32(0) // u32
		_, _ = minRequiredVersion_635_major, minRequiredVersion_635_minor
	default:
		v := ϟa.Attachment
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Parameter {
	case GLenum_GL_FRAMEBUFFER_ATTACHMENT_OBJECT_NAME, GLenum_GL_FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL:
	case GLenum_GL_FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_BLUE_SIZE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING, GLenum_GL_FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_GREEN_SIZE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_RED_SIZE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER:
		minRequiredVersion_637_major := uint32(3) // u32
		minRequiredVersion_637_minor := uint32(0) // u32
		_, _ = minRequiredVersion_637_major, minRequiredVersion_637_minor
	case GLenum_GL_FRAMEBUFFER_ATTACHMENT_LAYERED:
		minRequiredVersion_638_major := uint32(3) // u32
		minRequiredVersion_638_minor := uint32(2) // u32
		_, _ = minRequiredVersion_638_major, minRequiredVersion_638_minor
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_640_result := context             // Contextʳ
	ctx := GetContext_640_result                 // Contextʳ
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
	_, _, _, _, _, _, _, _, _ = minRequiredVersion_632_major, minRequiredVersion_632_minor, context, GetContext_640_result, ctx, target, framebufferId, framebuffer, a
	return nil
}
func (ϟa *GlGetFramebufferParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_641_major := uint32(3) // u32
	minRequiredVersion_641_minor := uint32(1) // u32
	switch ϟa.Target {
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Pname {
	case GLenum_GL_FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS, GLenum_GL_FRAMEBUFFER_DEFAULT_HEIGHT, GLenum_GL_FRAMEBUFFER_DEFAULT_SAMPLES, GLenum_GL_FRAMEBUFFER_DEFAULT_WIDTH:
	case GLenum_GL_FRAMEBUFFER_DEFAULT_LAYERS:
		minRequiredVersion_643_major := uint32(3) // u32
		minRequiredVersion_643_minor := uint32(2) // u32
		_, _ = minRequiredVersion_643_major, minRequiredVersion_643_minor
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_641_major, minRequiredVersion_641_minor
	return nil
}
func (ϟa *GlGetRenderbufferParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_645_major := uint32(2) // u32
	minRequiredVersion_645_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_RENDERBUFFER:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Parameter {
	case GLenum_GL_RENDERBUFFER_ALPHA_SIZE, GLenum_GL_RENDERBUFFER_BLUE_SIZE, GLenum_GL_RENDERBUFFER_DEPTH_SIZE, GLenum_GL_RENDERBUFFER_GREEN_SIZE, GLenum_GL_RENDERBUFFER_HEIGHT, GLenum_GL_RENDERBUFFER_INTERNAL_FORMAT, GLenum_GL_RENDERBUFFER_RED_SIZE, GLenum_GL_RENDERBUFFER_STENCIL_SIZE, GLenum_GL_RENDERBUFFER_WIDTH:
	case GLenum_GL_RENDERBUFFER_SAMPLES:
		minRequiredVersion_647_major := uint32(3) // u32
		minRequiredVersion_647_minor := uint32(0) // u32
		_, _ = minRequiredVersion_647_major, minRequiredVersion_647_minor
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_649_result := context             // Contextʳ
	ctx := GetContext_649_result                 // Contextʳ
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
	_, _, _, _, _, _, _ = minRequiredVersion_645_major, minRequiredVersion_645_minor, context, GetContext_649_result, ctx, id, rb
	return nil
}
func (ϟa *GlInvalidateFramebuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_650_major := uint32(3) // u32
	minRequiredVersion_650_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_FRAMEBUFFER:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_650_major, minRequiredVersion_650_minor
	return nil
}
func (ϟa *GlInvalidateSubFramebuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_652_major := uint32(3) // u32
	minRequiredVersion_652_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_FRAMEBUFFER:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_652_major, minRequiredVersion_652_minor
	return nil
}
func (ϟa *GlIsFramebuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_654_major := uint32(2)    // u32
	minRequiredVersion_654_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_655_result := context             // Contextʳ
	ctx := GetContext_655_result                 // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.Framebuffers.Contains(ϟa.Framebuffer) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = minRequiredVersion_654_major, minRequiredVersion_654_minor, context, GetContext_655_result, ctx
	return nil
}
func (ϟa *GlIsRenderbuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_656_major := uint32(2)    // u32
	minRequiredVersion_656_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_657_result := context             // Contextʳ
	ctx := GetContext_657_result                 // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.Renderbuffers.Contains(ϟa.Renderbuffer) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = minRequiredVersion_656_major, minRequiredVersion_656_minor, context, GetContext_657_result, ctx
	return nil
}
func (ϟa *GlReadBuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_658_major := uint32(3) // u32
	minRequiredVersion_658_minor := uint32(0) // u32
	switch ϟa.Src {
	case GLenum_GL_BACK, GLenum_GL_COLOR_ATTACHMENT0, GLenum_GL_COLOR_ATTACHMENT1, GLenum_GL_COLOR_ATTACHMENT10, GLenum_GL_COLOR_ATTACHMENT11, GLenum_GL_COLOR_ATTACHMENT12, GLenum_GL_COLOR_ATTACHMENT13, GLenum_GL_COLOR_ATTACHMENT14, GLenum_GL_COLOR_ATTACHMENT15, GLenum_GL_COLOR_ATTACHMENT2, GLenum_GL_COLOR_ATTACHMENT3, GLenum_GL_COLOR_ATTACHMENT4, GLenum_GL_COLOR_ATTACHMENT5, GLenum_GL_COLOR_ATTACHMENT6, GLenum_GL_COLOR_ATTACHMENT7, GLenum_GL_COLOR_ATTACHMENT8, GLenum_GL_COLOR_ATTACHMENT9, GLenum_GL_NONE:
	default:
		v := ϟa.Src
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_658_major, minRequiredVersion_658_minor
	return nil
}
func (ϟa *GlReadPixels) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_660_major := uint32(2) // u32
	minRequiredVersion_660_minor := uint32(0) // u32
	switch ϟa.Format {
	case GLenum_GL_RGBA, GLenum_GL_RGBA_INTEGER:
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Type {
	case GLenum_GL_FLOAT, GLenum_GL_INT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT:
	case GLenum_GL_UNSIGNED_INT_2_10_10_10_REV:
		minRequiredVersion_662_major := uint32(3) // u32
		minRequiredVersion_662_minor := uint32(1) // u32
		_, _ = minRequiredVersion_662_major, minRequiredVersion_662_minor
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Data.Slice(uint64(uint32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.imageSize(uint32(ϟa.Width), uint32(ϟa.Height), ϟa.Format, ϟa.Type)), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_660_major, minRequiredVersion_660_minor
	return nil
}
func (ϟa *GlReadnPixels) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_664_major := uint32(3) // u32
	minRequiredVersion_664_minor := uint32(2) // u32
	switch ϟa.Format {
	case GLenum_GL_RGBA, GLenum_GL_RGBA_INTEGER:
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Type {
	case GLenum_GL_FLOAT, GLenum_GL_INT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV:
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_664_major, minRequiredVersion_664_minor
	return nil
}
func (ϟa *GlRenderbufferStorage) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_667_major := uint32(2) // u32
	minRequiredVersion_667_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_RENDERBUFFER:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Format {
	case GLenum_GL_DEPTH_COMPONENT16, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGBA4, GLenum_GL_STENCIL_INDEX8:
	case GLenum_GL_DEPTH24_STENCIL8, GLenum_GL_DEPTH32F_STENCIL8, GLenum_GL_DEPTH_COMPONENT24, GLenum_GL_DEPTH_COMPONENT32F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB8, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_SRGB8_ALPHA8:
		minRequiredVersion_669_major := uint32(3) // u32
		minRequiredVersion_669_minor := uint32(0) // u32
		_, _ = minRequiredVersion_669_major, minRequiredVersion_669_minor
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_671_result := context             // Contextʳ
	ctx := GetContext_671_result                 // Contextʳ
	id := ctx.BoundRenderbuffers.Get(ϟa.Target)  // RenderbufferId
	rb := ctx.Instances.Renderbuffers.Get(id)    // Renderbufferʳ
	rb.Format = ϟa.Format
	rb.Width = ϟa.Width
	rb.Height = ϟa.Height
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = minRequiredVersion_667_major, minRequiredVersion_667_minor, context, GetContext_671_result, ctx, id, rb
	return nil
}
func (ϟa *GlRenderbufferStorageMultisample) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_672_major := uint32(3) // u32
	minRequiredVersion_672_minor := uint32(0) // u32
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
	_, _ = minRequiredVersion_672_major, minRequiredVersion_672_minor
	return nil
}
func (ϟa *GlStencilMask) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_675_major := uint32(2)    // u32
	minRequiredVersion_675_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_676_result := context             // Contextʳ
	ctx := GetContext_676_result                 // Contextʳ
	ctx.Rasterizing.StencilMask[GLenum_GL_FRONT] = ϟa.Mask
	ctx.Rasterizing.StencilMask[GLenum_GL_BACK] = ϟa.Mask
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_675_major, minRequiredVersion_675_minor, context, GetContext_676_result, ctx
	return nil
}
func (ϟa *GlStencilMaskSeparate) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_677_major := uint32(2) // u32
	minRequiredVersion_677_minor := uint32(0) // u32
	switch ϟa.Face {
	case GLenum_GL_BACK, GLenum_GL_FRONT, GLenum_GL_FRONT_AND_BACK:
	default:
		v := ϟa.Face
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_679_result := context             // Contextʳ
	ctx := GetContext_679_result                 // Contextʳ
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
	_, _, _, _, _ = minRequiredVersion_677_major, minRequiredVersion_677_minor, context, GetContext_679_result, ctx
	return nil
}
func (ϟa *GlDisable) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_680_major := uint32(2) // u32
	minRequiredVersion_680_minor := uint32(0) // u32
	switch ϟa.Capability {
	case GLenum_GL_BLEND, GLenum_GL_CULL_FACE, GLenum_GL_DEPTH_TEST, GLenum_GL_DITHER, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SCISSOR_TEST, GLenum_GL_STENCIL_TEST:
	case GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_RASTERIZER_DISCARD:
		minRequiredVersion_681_major := uint32(3) // u32
		minRequiredVersion_681_minor := uint32(0) // u32
		_, _ = minRequiredVersion_681_major, minRequiredVersion_681_minor
	case GLenum_GL_SAMPLE_MASK:
		minRequiredVersion_682_major := uint32(3) // u32
		minRequiredVersion_682_minor := uint32(1) // u32
		_, _ = minRequiredVersion_682_major, minRequiredVersion_682_minor
	case GLenum_GL_DEBUG_OUTPUT, GLenum_GL_DEBUG_OUTPUT_SYNCHRONOUS:
		minRequiredVersion_683_major := uint32(3) // u32
		minRequiredVersion_683_minor := uint32(2) // u32
		_, _ = minRequiredVersion_683_major, minRequiredVersion_683_minor
	default:
		v := ϟa.Capability
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_685_result := context             // Contextʳ
	ctx := GetContext_685_result                 // Contextʳ
	ctx.Capabilities[ϟa.Capability] = false
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_680_major, minRequiredVersion_680_minor, context, GetContext_685_result, ctx
	return nil
}
func (ϟa *GlDisablei) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_686_major := uint32(3) // u32
	minRequiredVersion_686_minor := uint32(2) // u32
	switch ϟa.Target {
	case GLenum_GL_BLEND, GLenum_GL_CULL_FACE, GLenum_GL_DEBUG_OUTPUT, GLenum_GL_DEBUG_OUTPUT_SYNCHRONOUS, GLenum_GL_DEPTH_TEST, GLenum_GL_DITHER, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_MASK, GLenum_GL_SCISSOR_TEST, GLenum_GL_STENCIL_TEST:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_686_major, minRequiredVersion_686_minor
	return nil
}
func (ϟa *GlEnable) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_688_major := uint32(2) // u32
	minRequiredVersion_688_minor := uint32(0) // u32
	switch ϟa.Capability {
	case GLenum_GL_BLEND, GLenum_GL_CULL_FACE, GLenum_GL_DEPTH_TEST, GLenum_GL_DITHER, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SCISSOR_TEST, GLenum_GL_STENCIL_TEST:
	case GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_RASTERIZER_DISCARD:
		minRequiredVersion_689_major := uint32(3) // u32
		minRequiredVersion_689_minor := uint32(0) // u32
		_, _ = minRequiredVersion_689_major, minRequiredVersion_689_minor
	case GLenum_GL_SAMPLE_MASK:
		minRequiredVersion_690_major := uint32(3) // u32
		minRequiredVersion_690_minor := uint32(1) // u32
		_, _ = minRequiredVersion_690_major, minRequiredVersion_690_minor
	case GLenum_GL_DEBUG_OUTPUT, GLenum_GL_DEBUG_OUTPUT_SYNCHRONOUS:
		minRequiredVersion_691_major := uint32(3) // u32
		minRequiredVersion_691_minor := uint32(2) // u32
		_, _ = minRequiredVersion_691_major, minRequiredVersion_691_minor
	default:
		v := ϟa.Capability
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_693_result := context             // Contextʳ
	ctx := GetContext_693_result                 // Contextʳ
	ctx.Capabilities[ϟa.Capability] = true
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_688_major, minRequiredVersion_688_minor, context, GetContext_693_result, ctx
	return nil
}
func (ϟa *GlEnablei) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_694_major := uint32(3) // u32
	minRequiredVersion_694_minor := uint32(2) // u32
	switch ϟa.Target {
	case GLenum_GL_BLEND, GLenum_GL_CULL_FACE, GLenum_GL_DEBUG_OUTPUT, GLenum_GL_DEBUG_OUTPUT_SYNCHRONOUS, GLenum_GL_DEPTH_TEST, GLenum_GL_DITHER, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_MASK, GLenum_GL_SCISSOR_TEST, GLenum_GL_STENCIL_TEST:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_694_major, minRequiredVersion_694_minor
	return nil
}
func (ϟa *GlFinish) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_696_major := uint32(2) // u32
	minRequiredVersion_696_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_696_major, minRequiredVersion_696_minor
	return nil
}
func (ϟa *GlFlush) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_697_major := uint32(2) // u32
	minRequiredVersion_697_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_697_major, minRequiredVersion_697_minor
	return nil
}
func (ϟa *GlFlushMappedBufferRange) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_698_major := uint32(3) // u32
	minRequiredVersion_698_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER, GLenum_GL_TEXTURE_BUFFER:
		minRequiredVersion_699_major := uint32(3) // u32
		minRequiredVersion_699_minor := uint32(2) // u32
		_, _ = minRequiredVersion_699_major, minRequiredVersion_699_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_698_major, minRequiredVersion_698_minor
	return nil
}
func (ϟa *GlGetError) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_701_major := uint32(2) // u32
	minRequiredVersion_701_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_701_major, minRequiredVersion_701_minor
	return nil
}
func (ϟa *GlGetGraphicsResetStatus) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_702_major := uint32(3) // u32
	minRequiredVersion_702_minor := uint32(2) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_702_major, minRequiredVersion_702_minor
	return nil
}
func (ϟa *GlHint) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_703_major := uint32(2) // u32
	minRequiredVersion_703_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_GENERATE_MIPMAP_HINT:
	case GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT:
		minRequiredVersion_704_major := uint32(3) // u32
		minRequiredVersion_704_minor := uint32(0) // u32
		_, _ = minRequiredVersion_704_major, minRequiredVersion_704_minor
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
	GetContext_707_result := context             // Contextʳ
	ctx := GetContext_707_result                 // Contextʳ
	ctx.GenerateMipmapHint = ϟa.Mode
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_703_major, minRequiredVersion_703_minor, context, GetContext_707_result, ctx
	return nil
}
func (ϟa *GlActiveShaderProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_708_major := uint32(3) // u32
	minRequiredVersion_708_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_708_major, minRequiredVersion_708_minor
	return nil
}
func (ϟa *GlAttachShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_709_major := uint32(2)    // u32
	minRequiredVersion_709_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_710_result := context             // Contextʳ
	ctx := GetContext_710_result                 // Contextʳ
	p := ctx.Instances.Programs.Get(ϟa.Program)  // Programʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // Shaderʳ
	p.Shaders[s.Type] = ϟa.Shader
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = minRequiredVersion_709_major, minRequiredVersion_709_minor, context, GetContext_710_result, ctx, p, s
	return nil
}
func (ϟa *GlBindAttribLocation) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_711_major := uint32(2)    // u32
	minRequiredVersion_711_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_712_result := context             // Contextʳ
	ctx := GetContext_712_result                 // Contextʳ
	p := ctx.Instances.Programs.Get(ϟa.Program)  // Programʳ
	p.AttributeBindings[ϟa.Name] = ϟa.Location
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_711_major, minRequiredVersion_711_minor, context, GetContext_712_result, ctx, p
	return nil
}
func (ϟa *GlBindProgramPipeline) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_713_major := uint32(3) // u32
	minRequiredVersion_713_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_713_major, minRequiredVersion_713_minor
	return nil
}
func (ϟa *GlCompileShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_714_major := uint32(2) // u32
	minRequiredVersion_714_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_714_major, minRequiredVersion_714_minor
	return nil
}
func (ϟa *GlCreateProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_715_major := uint32(2)    // u32
	minRequiredVersion_715_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_716_result := context             // Contextʳ
	ctx := GetContext_716_result                 // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	id := ProgramId(ϟa.Result) // ProgramId
	ctx.Instances.Programs[id] = &Program{Shaders: GLenumːShaderIdᵐ{}, AttributeBindings: StringːAttributeLocationᵐ{}, Attributes: S32ːVertexAttributeᵐ{}, Uniforms: UniformLocationːUniformᵐ{}}
	ϟa.Result = id
	_, _, _, _, _, _ = minRequiredVersion_715_major, minRequiredVersion_715_minor, context, GetContext_716_result, ctx, id
	return nil
}
func (ϟa *GlCreateShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_717_major := uint32(2) // u32
	minRequiredVersion_717_minor := uint32(0) // u32
	switch ϟa.Type {
	case GLenum_GL_FRAGMENT_SHADER, GLenum_GL_VERTEX_SHADER:
	case GLenum_GL_COMPUTE_SHADER:
		minRequiredVersion_718_major := uint32(3) // u32
		minRequiredVersion_718_minor := uint32(1) // u32
		_, _ = minRequiredVersion_718_major, minRequiredVersion_718_minor
	case GLenum_GL_GEOMETRY_SHADER, GLenum_GL_TESS_CONTROL_SHADER, GLenum_GL_TESS_EVALUATION_SHADER:
		minRequiredVersion_719_major := uint32(3) // u32
		minRequiredVersion_719_minor := uint32(2) // u32
		_, _ = minRequiredVersion_719_major, minRequiredVersion_719_minor
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_721_result := context             // Contextʳ
	ctx := GetContext_721_result                 // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	id := ShaderId(ϟa.Result) // ShaderId
	ctx.Instances.Shaders[id] = &Shader{Compiled: false, Deletable: false}
	s := ctx.Instances.Shaders.Get(id) // Shaderʳ
	s.Type = ϟa.Type
	ϟa.Result = id
	_, _, _, _, _, _, _ = minRequiredVersion_717_major, minRequiredVersion_717_minor, context, GetContext_721_result, ctx, id, s
	return nil
}
func (ϟa *GlCreateShaderProgramv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_722_major := uint32(3) // u32
	minRequiredVersion_722_minor := uint32(1) // u32
	switch ϟa.Type {
	case GLenum_GL_COMPUTE_SHADER, GLenum_GL_FRAGMENT_SHADER, GLenum_GL_VERTEX_SHADER:
	case GLenum_GL_GEOMETRY_SHADER, GLenum_GL_TESS_CONTROL_SHADER, GLenum_GL_TESS_EVALUATION_SHADER:
		minRequiredVersion_723_major := uint32(3) // u32
		minRequiredVersion_723_minor := uint32(2) // u32
		_, _ = minRequiredVersion_723_major, minRequiredVersion_723_minor
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_722_major, minRequiredVersion_722_minor
	return nil
}
func (ϟa *GlDeleteProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_725_major := uint32(2)    // u32
	minRequiredVersion_725_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_726_result := context             // Contextʳ
	ctx := GetContext_726_result                 // Contextʳ
	delete(ctx.Instances.Programs, ϟa.Program)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_725_major, minRequiredVersion_725_minor, context, GetContext_726_result, ctx
	return nil
}
func (ϟa *GlDeleteProgramPipelines) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_727_major := uint32(3) // u32
	minRequiredVersion_727_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_727_major, minRequiredVersion_727_minor
	return nil
}
func (ϟa *GlDeleteShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_728_major := uint32(2)    // u32
	minRequiredVersion_728_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_729_result := context             // Contextʳ
	ctx := GetContext_729_result                 // Contextʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // Shaderʳ
	s.Deletable = true
	delete(ctx.Instances.Shaders, ϟa.Shader)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_728_major, minRequiredVersion_728_minor, context, GetContext_729_result, ctx, s
	return nil
}
func (ϟa *GlDetachShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_730_major := uint32(2)    // u32
	minRequiredVersion_730_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_731_result := context             // Contextʳ
	ctx := GetContext_731_result                 // Contextʳ
	p := ctx.Instances.Programs.Get(ϟa.Program)  // Programʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // Shaderʳ
	delete(p.Shaders, s.Type)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = minRequiredVersion_730_major, minRequiredVersion_730_minor, context, GetContext_731_result, ctx, p, s
	return nil
}
func (ϟa *GlDispatchCompute) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_732_major := uint32(3) // u32
	minRequiredVersion_732_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_732_major, minRequiredVersion_732_minor
	return nil
}
func (ϟa *GlDispatchComputeIndirect) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_733_major := uint32(3) // u32
	minRequiredVersion_733_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_733_major, minRequiredVersion_733_minor
	return nil
}
func (ϟa *GlGenProgramPipelines) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_734_major := uint32(3) // u32
	minRequiredVersion_734_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_734_major, minRequiredVersion_734_minor
	return nil
}
func (ϟa *GlGetActiveAttrib) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_735_major := uint32(2) // u32
	minRequiredVersion_735_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	if (ϟa.BufferBytesWritten) != (GLsizeiᵖ{}) {
		l := GLsizei(ϟa.BufferBytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // GLsizei
		ϟa.BufferBytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(l, ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟa.Name.Slice(uint64(GLsizei(int32(0))), uint64(l), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = l
	} else {
		ϟa.Name.Slice(uint64(0), uint64(256), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	}
	ϟa.VectorCount.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(GLint(ϟa.VectorCount.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Type.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(GLenum(ϟa.Type.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_735_major, minRequiredVersion_735_minor
	return nil
}
func (ϟa *GlGetActiveUniform) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_736_major := uint32(2) // u32
	minRequiredVersion_736_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	if (ϟa.BufferBytesWritten) != (GLsizeiᵖ{}) {
		l := GLsizei(ϟa.BufferBytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // GLsizei
		ϟa.BufferBytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(l, ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟa.Name.Slice(uint64(GLsizei(int32(0))), uint64(l), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = l
	} else {
		ϟa.Name.Slice(uint64(0), uint64(256), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	}
	ϟa.VectorCount.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(GLint(ϟa.VectorCount.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Type.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(GLenum(ϟa.Type.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_736_major, minRequiredVersion_736_minor
	return nil
}
func (ϟa *GlGetActiveUniformBlockName) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_737_major := uint32(3) // u32
	minRequiredVersion_737_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	l := GLsizei(ϟa.BufferBytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // GLsizei
	ϟa.BufferBytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(l, ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Name.Slice(uint64(GLsizei(int32(0))), uint64(l), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _ = minRequiredVersion_737_major, minRequiredVersion_737_minor, l
	return nil
}
func (ϟa *GlGetActiveUniformBlockiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_738_major := uint32(3) // u32
	minRequiredVersion_738_minor := uint32(0) // u32
	switch ϟa.ParameterName {
	case GLenum_GL_UNIFORM_BLOCK_ACTIVE_UNIFORMS, GLenum_GL_UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES, GLenum_GL_UNIFORM_BLOCK_BINDING, GLenum_GL_UNIFORM_BLOCK_DATA_SIZE, GLenum_GL_UNIFORM_BLOCK_NAME_LENGTH, GLenum_GL_UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER, GLenum_GL_UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER:
	default:
		v := ϟa.ParameterName
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Parameters.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Parameters.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_738_major, minRequiredVersion_738_minor
	return nil
}
func (ϟa *GlGetActiveUniformsiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_740_major := uint32(3) // u32
	minRequiredVersion_740_minor := uint32(0) // u32
	switch ϟa.ParameterName {
	case GLenum_GL_UNIFORM_ARRAY_STRIDE, GLenum_GL_UNIFORM_BLOCK_INDEX, GLenum_GL_UNIFORM_IS_ROW_MAJOR, GLenum_GL_UNIFORM_MATRIX_STRIDE, GLenum_GL_UNIFORM_NAME_LENGTH, GLenum_GL_UNIFORM_OFFSET, GLenum_GL_UNIFORM_SIZE, GLenum_GL_UNIFORM_TYPE:
	default:
		v := ϟa.ParameterName
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.UniformIndices.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.UniformCount), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Parameters.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.UniformCount), ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_740_major, minRequiredVersion_740_minor
	return nil
}
func (ϟa *GlGetAttachedShaders) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_742_major := uint32(2)    // u32
	minRequiredVersion_742_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_743_result := context             // Contextʳ
	ctx := GetContext_743_result                 // Contextʳ
	p := ctx.Instances.Programs.Get(ϟa.Program)  // Programʳ
	min_744_a := int32(ϟa.BufferLength)          // s32
	min_744_b := int32(len(p.Shaders))           // s32
	min_744_result := func() (result int32) {
		switch (min_744_a) < (min_744_b) {
		case true:
			return min_744_a
		case false:
			return min_744_b
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (min_744_a) < (min_744_b), ϟa))
			return result
		}
	}() // s32
	l := min_744_result // s32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.ShadersLengthWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(GLsizei(l), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _, _, _, _, _ = minRequiredVersion_742_major, minRequiredVersion_742_minor, context, GetContext_743_result, ctx, p, min_744_a, min_744_b, min_744_result, l
	return nil
}
func (ϟa *GlGetAttribLocation) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_745_major := uint32(2) // u32
	minRequiredVersion_745_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_745_major, minRequiredVersion_745_minor
	return nil
}
func (ϟa *GlGetFragDataLocation) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_746_major := uint32(3) // u32
	minRequiredVersion_746_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_746_major, minRequiredVersion_746_minor
	return nil
}
func (ϟa *GlGetProgramBinary) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_747_major := uint32(3) // u32
	minRequiredVersion_747_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_747_major, minRequiredVersion_747_minor
	return nil
}
func (ϟa *GlGetProgramInfoLog) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_748_major := uint32(2)    // u32
	minRequiredVersion_748_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_749_result := context             // Contextʳ
	ctx := GetContext_749_result                 // Contextʳ
	p := ctx.Instances.Programs.Get(ϟa.Program)  // Programʳ
	min_750_a := int32(ϟa.BufferLength)          // s32
	min_750_b := int32(p.InfoLog.Count)          // s32
	min_750_result := func() (result int32) {
		switch (min_750_a) < (min_750_b) {
		case true:
			return min_750_a
		case false:
			return min_750_b
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (min_750_a) < (min_750_b), ϟa))
			return result
		}
	}() // s32
	l := min_750_result // s32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Info.Slice(uint64(int32(0)), uint64(l), ϟs).Copy(p.InfoLog.Slice(uint64(int32(0)), uint64(l), ϟs), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.StringLengthWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(GLsizei(l), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _, _, _, _, _ = minRequiredVersion_748_major, minRequiredVersion_748_minor, context, GetContext_749_result, ctx, p, min_750_a, min_750_b, min_750_result, l
	return nil
}
func (ϟa *GlGetProgramInterfaceiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_751_major := uint32(3) // u32
	minRequiredVersion_751_minor := uint32(1) // u32
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
	_, _ = minRequiredVersion_751_major, minRequiredVersion_751_minor
	return nil
}
func (ϟa *GlGetProgramPipelineInfoLog) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_754_major := uint32(3) // u32
	minRequiredVersion_754_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_754_major, minRequiredVersion_754_minor
	return nil
}
func (ϟa *GlGetProgramPipelineiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_755_major := uint32(3) // u32
	minRequiredVersion_755_minor := uint32(1) // u32
	switch ϟa.Pname {
	case GLenum_GL_ACTIVE_PROGRAM, GLenum_GL_COMPUTE_SHADER, GLenum_GL_FRAGMENT_SHADER, GLenum_GL_INFO_LOG_LENGTH, GLenum_GL_VALIDATE_STATUS, GLenum_GL_VERTEX_SHADER:
	case GLenum_GL_GEOMETRY_SHADER, GLenum_GL_TESS_CONTROL_SHADER, GLenum_GL_TESS_EVALUATION_SHADER:
		minRequiredVersion_756_major := uint32(3) // u32
		minRequiredVersion_756_minor := uint32(2) // u32
		_, _ = minRequiredVersion_756_major, minRequiredVersion_756_minor
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_755_major, minRequiredVersion_755_minor
	return nil
}
func (ϟa *GlGetProgramResourceIndex) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_758_major := uint32(3) // u32
	minRequiredVersion_758_minor := uint32(1) // u32
	switch ϟa.ProgramInterface {
	case GLenum_GL_BUFFER_VARIABLE, GLenum_GL_PROGRAM_INPUT, GLenum_GL_PROGRAM_OUTPUT, GLenum_GL_SHADER_STORAGE_BLOCK, GLenum_GL_TRANSFORM_FEEDBACK_VARYING, GLenum_GL_UNIFORM, GLenum_GL_UNIFORM_BLOCK:
	default:
		v := ϟa.ProgramInterface
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_758_major, minRequiredVersion_758_minor
	return nil
}
func (ϟa *GlGetProgramResourceLocation) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_760_major := uint32(3) // u32
	minRequiredVersion_760_minor := uint32(1) // u32
	switch ϟa.ProgramInterface {
	case GLenum_GL_PROGRAM_INPUT, GLenum_GL_PROGRAM_OUTPUT, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM:
	default:
		v := ϟa.ProgramInterface
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_760_major, minRequiredVersion_760_minor
	return nil
}
func (ϟa *GlGetProgramResourceName) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_762_major := uint32(3) // u32
	minRequiredVersion_762_minor := uint32(1) // u32
	switch ϟa.ProgramInterface {
	case GLenum_GL_BUFFER_VARIABLE, GLenum_GL_PROGRAM_INPUT, GLenum_GL_PROGRAM_OUTPUT, GLenum_GL_SHADER_STORAGE_BLOCK, GLenum_GL_TRANSFORM_FEEDBACK_VARYING, GLenum_GL_UNIFORM, GLenum_GL_UNIFORM_BLOCK:
	default:
		v := ϟa.ProgramInterface
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_762_major, minRequiredVersion_762_minor
	return nil
}
func (ϟa *GlGetProgramResourceiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_764_major := uint32(3) // u32
	minRequiredVersion_764_minor := uint32(1) // u32
	switch ϟa.ProgramInterface {
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_BUFFER_VARIABLE, GLenum_GL_PROGRAM_INPUT, GLenum_GL_PROGRAM_OUTPUT, GLenum_GL_SHADER_STORAGE_BLOCK, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_VARYING, GLenum_GL_UNIFORM, GLenum_GL_UNIFORM_BLOCK:
	default:
		v := ϟa.ProgramInterface
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_764_major, minRequiredVersion_764_minor
	return nil
}
func (ϟa *GlGetProgramiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_766_major := uint32(2) // u32
	minRequiredVersion_766_minor := uint32(0) // u32
	switch ϟa.Parameter {
	case GLenum_GL_ACTIVE_ATTRIBUTES, GLenum_GL_ACTIVE_ATTRIBUTE_MAX_LENGTH, GLenum_GL_ACTIVE_UNIFORMS, GLenum_GL_ACTIVE_UNIFORM_MAX_LENGTH, GLenum_GL_ATTACHED_SHADERS, GLenum_GL_DELETE_STATUS, GLenum_GL_INFO_LOG_LENGTH, GLenum_GL_LINK_STATUS, GLenum_GL_VALIDATE_STATUS:
	case GLenum_GL_ACTIVE_UNIFORM_BLOCKS, GLenum_GL_ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH, GLenum_GL_PROGRAM_BINARY_RETRIEVABLE_HINT, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_MODE, GLenum_GL_TRANSFORM_FEEDBACK_VARYINGS, GLenum_GL_TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH:
		minRequiredVersion_767_major := uint32(3) // u32
		minRequiredVersion_767_minor := uint32(0) // u32
		_, _ = minRequiredVersion_767_major, minRequiredVersion_767_minor
	case GLenum_GL_ACTIVE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_PROGRAM_SEPARABLE:
		minRequiredVersion_768_major := uint32(3) // u32
		minRequiredVersion_768_minor := uint32(1) // u32
		_, _ = minRequiredVersion_768_major, minRequiredVersion_768_minor
	case GLenum_GL_GEOMETRY_INPUT_TYPE, GLenum_GL_GEOMETRY_OUTPUT_TYPE, GLenum_GL_GEOMETRY_VERTICES_OUT, GLenum_GL_GEOMETRY_SHADER_INVOCATIONS, GLenum_GL_TESS_CONTROL_OUTPUT_VERTICES, GLenum_GL_TESS_GEN_MODE, GLenum_GL_TESS_GEN_POINT_MODE, GLenum_GL_TESS_GEN_SPACING, GLenum_GL_TESS_GEN_VERTEX_ORDER:
		minRequiredVersion_769_major := uint32(3) // u32
		minRequiredVersion_769_minor := uint32(2) // u32
		_, _ = minRequiredVersion_769_major, minRequiredVersion_769_minor
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_766_major, minRequiredVersion_766_minor
	return nil
}
func (ϟa *GlGetShaderInfoLog) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_771_major := uint32(2)    // u32
	minRequiredVersion_771_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_772_result := context             // Contextʳ
	ctx := GetContext_772_result                 // Contextʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // Shaderʳ
	min_773_a := int32(ϟa.BufferLength)          // s32
	min_773_b := int32(s.InfoLog.Count)          // s32
	min_773_result := func() (result int32) {
		switch (min_773_a) < (min_773_b) {
		case true:
			return min_773_a
		case false:
			return min_773_b
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (min_773_a) < (min_773_b), ϟa))
			return result
		}
	}() // s32
	l := min_773_result // s32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Info.Slice(uint64(int32(0)), uint64(l), ϟs).Copy(s.InfoLog.Slice(uint64(int32(0)), uint64(l), ϟs), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.StringLengthWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(GLsizei(l), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _, _, _, _, _ = minRequiredVersion_771_major, minRequiredVersion_771_minor, context, GetContext_772_result, ctx, s, min_773_a, min_773_b, min_773_result, l
	return nil
}
func (ϟa *GlGetShaderPrecisionFormat) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_774_major := uint32(2) // u32
	minRequiredVersion_774_minor := uint32(0) // u32
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
	ϟa.Precision.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Precision.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_774_major, minRequiredVersion_774_minor
	return nil
}
func (ϟa *GlGetShaderSource) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_777_major := uint32(2)    // u32
	minRequiredVersion_777_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_778_result := context             // Contextʳ
	ctx := GetContext_778_result                 // Contextʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // Shaderʳ
	min_779_a := int32(ϟa.BufferLength)          // s32
	min_779_b := int32(len(s.Source))            // s32
	min_779_result := func() (result int32) {
		switch (min_779_a) < (min_779_b) {
		case true:
			return min_779_a
		case false:
			return min_779_b
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (min_779_a) < (min_779_b), ϟa))
			return result
		}
	}() // s32
	l := min_779_result // s32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	Charᵖ(ϟa.Source).Slice(uint64(int32(0)), uint64(l), ϟs).Copy(MakeCharˢFromString(s.Source, ϟs).Slice(uint64(int32(0)), uint64(l), ϟs), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.StringLengthWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(GLsizei(l), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _, _, _, _, _ = minRequiredVersion_777_major, minRequiredVersion_777_minor, context, GetContext_778_result, ctx, s, min_779_a, min_779_b, min_779_result, l
	return nil
}
func (ϟa *GlGetShaderiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_780_major := uint32(2) // u32
	minRequiredVersion_780_minor := uint32(0) // u32
	switch ϟa.Parameter {
	case GLenum_GL_COMPILE_STATUS, GLenum_GL_DELETE_STATUS, GLenum_GL_INFO_LOG_LENGTH, GLenum_GL_SHADER_SOURCE_LENGTH, GLenum_GL_SHADER_TYPE:
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_782_result := context             // Contextʳ
	ctx := GetContext_782_result                 // Contextʳ
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
	_, _, _, _, _, _ = minRequiredVersion_780_major, minRequiredVersion_780_minor, context, GetContext_782_result, ctx, s
	return nil
}
func (ϟa *GlGetUniformBlockIndex) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_783_major := uint32(3) // u32
	minRequiredVersion_783_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_783_major, minRequiredVersion_783_minor
	return nil
}
func (ϟa *GlGetUniformIndices) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_784_major := uint32(3) // u32
	minRequiredVersion_784_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_784_major, minRequiredVersion_784_minor
	return nil
}
func (ϟa *GlGetUniformLocation) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_785_major := uint32(2) // u32
	minRequiredVersion_785_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_785_major, minRequiredVersion_785_minor
	return nil
}
func (ϟa *GlGetUniformfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_786_major := uint32(2) // u32
	minRequiredVersion_786_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_786_major, minRequiredVersion_786_minor
	return nil
}
func (ϟa *GlGetUniformiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_787_major := uint32(2) // u32
	minRequiredVersion_787_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_787_major, minRequiredVersion_787_minor
	return nil
}
func (ϟa *GlGetUniformuiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_788_major := uint32(3) // u32
	minRequiredVersion_788_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_788_major, minRequiredVersion_788_minor
	return nil
}
func (ϟa *GlGetnUniformfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_789_major := uint32(3) // u32
	minRequiredVersion_789_minor := uint32(2) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_789_major, minRequiredVersion_789_minor
	return nil
}
func (ϟa *GlGetnUniformiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_790_major := uint32(3) // u32
	minRequiredVersion_790_minor := uint32(2) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_790_major, minRequiredVersion_790_minor
	return nil
}
func (ϟa *GlGetnUniformuiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_791_major := uint32(3) // u32
	minRequiredVersion_791_minor := uint32(2) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_791_major, minRequiredVersion_791_minor
	return nil
}
func (ϟa *GlIsProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_792_major := uint32(2)    // u32
	minRequiredVersion_792_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_793_result := context             // Contextʳ
	ctx := GetContext_793_result                 // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.Programs.Contains(ϟa.Program) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = minRequiredVersion_792_major, minRequiredVersion_792_minor, context, GetContext_793_result, ctx
	return nil
}
func (ϟa *GlIsProgramPipeline) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_794_major := uint32(3) // u32
	minRequiredVersion_794_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_794_major, minRequiredVersion_794_minor
	return nil
}
func (ϟa *GlIsShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_795_major := uint32(2)    // u32
	minRequiredVersion_795_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_796_result := context             // Contextʳ
	ctx := GetContext_796_result                 // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.Shaders.Contains(ϟa.Shader) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = minRequiredVersion_795_major, minRequiredVersion_795_minor, context, GetContext_796_result, ctx
	return nil
}
func (ϟa *GlLinkProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_797_major := uint32(2) // u32
	minRequiredVersion_797_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_797_major, minRequiredVersion_797_minor
	return nil
}
func (ϟa *GlMemoryBarrier) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_798_major := uint32(3)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               // u32
	minRequiredVersion_798_minor := uint32(1)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               // u32
	supportsBits_799_seenBits := ϟa.Barriers                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                // GLbitfield
	supportsBits_799_validBits := (GLbitfield_GL_ALL_BARRIER_BITS) | ((GLbitfield_GL_ATOMIC_COUNTER_BARRIER_BIT) | ((GLbitfield_GL_BUFFER_UPDATE_BARRIER_BIT) | ((GLbitfield_GL_COMMAND_BARRIER_BIT) | ((GLbitfield_GL_ELEMENT_ARRAY_BARRIER_BIT) | ((GLbitfield_GL_FRAMEBUFFER_BARRIER_BIT) | ((GLbitfield_GL_PIXEL_BUFFER_BARRIER_BIT) | ((GLbitfield_GL_SHADER_IMAGE_ACCESS_BARRIER_BIT) | ((GLbitfield_GL_SHADER_STORAGE_BARRIER_BIT) | ((GLbitfield_GL_TEXTURE_FETCH_BARRIER_BIT) | ((GLbitfield_GL_TEXTURE_UPDATE_BARRIER_BIT) | ((GLbitfield_GL_TRANSFORM_FEEDBACK_BARRIER_BIT) | ((GLbitfield_GL_UNIFORM_BARRIER_BIT) | (GLbitfield_GL_VERTEX_ATTRIB_ARRAY_BARRIER_BIT))))))))))))) // GLbitfield
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
	_, _, _, _ = minRequiredVersion_798_major, minRequiredVersion_798_minor, supportsBits_799_seenBits, supportsBits_799_validBits
	return nil
}
func (ϟa *GlMemoryBarrierByRegion) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_800_major := uint32(3)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               // u32
	minRequiredVersion_800_minor := uint32(1)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               // u32
	supportsBits_801_seenBits := ϟa.Barriers                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                // GLbitfield
	supportsBits_801_validBits := (GLbitfield_GL_ALL_BARRIER_BITS) | ((GLbitfield_GL_ATOMIC_COUNTER_BARRIER_BIT) | ((GLbitfield_GL_BUFFER_UPDATE_BARRIER_BIT) | ((GLbitfield_GL_COMMAND_BARRIER_BIT) | ((GLbitfield_GL_ELEMENT_ARRAY_BARRIER_BIT) | ((GLbitfield_GL_FRAMEBUFFER_BARRIER_BIT) | ((GLbitfield_GL_PIXEL_BUFFER_BARRIER_BIT) | ((GLbitfield_GL_SHADER_IMAGE_ACCESS_BARRIER_BIT) | ((GLbitfield_GL_SHADER_STORAGE_BARRIER_BIT) | ((GLbitfield_GL_TEXTURE_FETCH_BARRIER_BIT) | ((GLbitfield_GL_TEXTURE_UPDATE_BARRIER_BIT) | ((GLbitfield_GL_TRANSFORM_FEEDBACK_BARRIER_BIT) | ((GLbitfield_GL_UNIFORM_BARRIER_BIT) | (GLbitfield_GL_VERTEX_ATTRIB_ARRAY_BARRIER_BIT))))))))))))) // GLbitfield
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
	_, _, _, _ = minRequiredVersion_800_major, minRequiredVersion_800_minor, supportsBits_801_seenBits, supportsBits_801_validBits
	return nil
}
func (ϟa *GlProgramBinary) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_802_major := uint32(3) // u32
	minRequiredVersion_802_minor := uint32(0) // u32
	switch ϟa.BinaryFormat {
	default:
		v := ϟa.BinaryFormat
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_802_major, minRequiredVersion_802_minor
	return nil
}
func (ϟa *GlProgramParameteri) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_804_major := uint32(3) // u32
	minRequiredVersion_804_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_PROGRAM_BINARY_RETRIEVABLE_HINT:
	case GLenum_GL_PROGRAM_SEPARABLE:
		minRequiredVersion_805_major := uint32(3) // u32
		minRequiredVersion_805_minor := uint32(1) // u32
		_, _ = minRequiredVersion_805_major, minRequiredVersion_805_minor
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_804_major, minRequiredVersion_804_minor
	return nil
}
func (ϟa *GlProgramUniform1f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_807_major := uint32(3) // u32
	minRequiredVersion_807_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_807_major, minRequiredVersion_807_minor
	return nil
}
func (ϟa *GlProgramUniform1fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_808_major := uint32(3) // u32
	minRequiredVersion_808_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_808_major, minRequiredVersion_808_minor
	return nil
}
func (ϟa *GlProgramUniform1i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_809_major := uint32(3) // u32
	minRequiredVersion_809_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_809_major, minRequiredVersion_809_minor
	return nil
}
func (ϟa *GlProgramUniform1iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_810_major := uint32(3) // u32
	minRequiredVersion_810_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_810_major, minRequiredVersion_810_minor
	return nil
}
func (ϟa *GlProgramUniform1ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_811_major := uint32(3) // u32
	minRequiredVersion_811_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_811_major, minRequiredVersion_811_minor
	return nil
}
func (ϟa *GlProgramUniform1uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_812_major := uint32(3) // u32
	minRequiredVersion_812_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_812_major, minRequiredVersion_812_minor
	return nil
}
func (ϟa *GlProgramUniform2f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_813_major := uint32(3) // u32
	minRequiredVersion_813_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_813_major, minRequiredVersion_813_minor
	return nil
}
func (ϟa *GlProgramUniform2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_814_major := uint32(3) // u32
	minRequiredVersion_814_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_814_major, minRequiredVersion_814_minor
	return nil
}
func (ϟa *GlProgramUniform2i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_815_major := uint32(3) // u32
	minRequiredVersion_815_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_815_major, minRequiredVersion_815_minor
	return nil
}
func (ϟa *GlProgramUniform2iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_816_major := uint32(3) // u32
	minRequiredVersion_816_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_816_major, minRequiredVersion_816_minor
	return nil
}
func (ϟa *GlProgramUniform2ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_817_major := uint32(3) // u32
	minRequiredVersion_817_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_817_major, minRequiredVersion_817_minor
	return nil
}
func (ϟa *GlProgramUniform2uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_818_major := uint32(3) // u32
	minRequiredVersion_818_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_818_major, minRequiredVersion_818_minor
	return nil
}
func (ϟa *GlProgramUniform3f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_819_major := uint32(3) // u32
	minRequiredVersion_819_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_819_major, minRequiredVersion_819_minor
	return nil
}
func (ϟa *GlProgramUniform3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_820_major := uint32(3) // u32
	minRequiredVersion_820_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_820_major, minRequiredVersion_820_minor
	return nil
}
func (ϟa *GlProgramUniform3i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_821_major := uint32(3) // u32
	minRequiredVersion_821_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_821_major, minRequiredVersion_821_minor
	return nil
}
func (ϟa *GlProgramUniform3iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_822_major := uint32(3) // u32
	minRequiredVersion_822_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_822_major, minRequiredVersion_822_minor
	return nil
}
func (ϟa *GlProgramUniform3ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_823_major := uint32(3) // u32
	minRequiredVersion_823_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_823_major, minRequiredVersion_823_minor
	return nil
}
func (ϟa *GlProgramUniform3uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_824_major := uint32(3) // u32
	minRequiredVersion_824_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_824_major, minRequiredVersion_824_minor
	return nil
}
func (ϟa *GlProgramUniform4f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_825_major := uint32(3) // u32
	minRequiredVersion_825_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_825_major, minRequiredVersion_825_minor
	return nil
}
func (ϟa *GlProgramUniform4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_826_major := uint32(3) // u32
	minRequiredVersion_826_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_826_major, minRequiredVersion_826_minor
	return nil
}
func (ϟa *GlProgramUniform4i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_827_major := uint32(3) // u32
	minRequiredVersion_827_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_827_major, minRequiredVersion_827_minor
	return nil
}
func (ϟa *GlProgramUniform4iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_828_major := uint32(3) // u32
	minRequiredVersion_828_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_828_major, minRequiredVersion_828_minor
	return nil
}
func (ϟa *GlProgramUniform4ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_829_major := uint32(3) // u32
	minRequiredVersion_829_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_829_major, minRequiredVersion_829_minor
	return nil
}
func (ϟa *GlProgramUniform4uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_830_major := uint32(3) // u32
	minRequiredVersion_830_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_830_major, minRequiredVersion_830_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_831_major := uint32(3) // u32
	minRequiredVersion_831_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_831_major, minRequiredVersion_831_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix2x3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_832_major := uint32(3) // u32
	minRequiredVersion_832_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_832_major, minRequiredVersion_832_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix2x4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_833_major := uint32(3) // u32
	minRequiredVersion_833_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_833_major, minRequiredVersion_833_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_834_major := uint32(3) // u32
	minRequiredVersion_834_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_834_major, minRequiredVersion_834_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix3x2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_835_major := uint32(3) // u32
	minRequiredVersion_835_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_835_major, minRequiredVersion_835_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix3x4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_836_major := uint32(3) // u32
	minRequiredVersion_836_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_836_major, minRequiredVersion_836_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_837_major := uint32(3) // u32
	minRequiredVersion_837_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_837_major, minRequiredVersion_837_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix4x2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_838_major := uint32(3) // u32
	minRequiredVersion_838_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_838_major, minRequiredVersion_838_minor
	return nil
}
func (ϟa *GlProgramUniformMatrix4x3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_839_major := uint32(3) // u32
	minRequiredVersion_839_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_839_major, minRequiredVersion_839_minor
	return nil
}
func (ϟa *GlReleaseShaderCompiler) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_840_major := uint32(2) // u32
	minRequiredVersion_840_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_840_major, minRequiredVersion_840_minor
	return nil
}
func (ϟa *GlShaderBinary) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_841_major := uint32(2) // u32
	minRequiredVersion_841_minor := uint32(0) // u32
	switch ϟa.BinaryFormat {
	default:
		v := ϟa.BinaryFormat
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_841_major, minRequiredVersion_841_minor
	return nil
}
func (ϟa *GlShaderSource) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_843_major := uint32(2)                                   // u32
	minRequiredVersion_843_minor := uint32(0)                                   // u32
	sources := ϟa.Source.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // GLcharᶜᵖˢ
	lengths := ϟa.Length.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // GLintˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                // Contextʳ
	GetContext_844_result := context                                            // Contextʳ
	ctx := GetContext_844_result                                                // Contextʳ
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
	_, _, _, _, _, _, _, _ = minRequiredVersion_843_major, minRequiredVersion_843_minor, sources, lengths, context, GetContext_844_result, ctx, s
	return nil
}
func (ϟa *GlUniform1f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_845_major := uint32(2)    // u32
	minRequiredVersion_845_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_846_result := context             // Contextʳ
	ctx := GetContext_846_result                 // Contextʳ
	v := MakeGLfloatˢ(uint64(1), ϟs)             // GLfloatˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).Write(ϟa.Value, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_FLOAT
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_845_major, minRequiredVersion_845_minor, context, GetContext_846_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform1fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_847_major := uint32(2)                             // u32
	minRequiredVersion_847_minor := uint32(0)                             // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                          // Contextʳ
	GetContext_848_result := context                                      // Contextʳ
	ctx := GetContext_848_result                                          // Contextʳ
	v := ϟa.Values.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // GLfloatˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)               // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                          // Uniform
	uniform.Type = GLenum_GL_FLOAT
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_847_major, minRequiredVersion_847_minor, context, GetContext_848_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform1i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_849_major := uint32(2)    // u32
	minRequiredVersion_849_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_850_result := context             // Contextʳ
	ctx := GetContext_850_result                 // Contextʳ
	v := MakeGLintˢ(uint64(1), ϟs)               // GLintˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).Write(ϟa.Value, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_INT
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_849_major, minRequiredVersion_849_minor, context, GetContext_850_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform1iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_851_major := uint32(2)                             // u32
	minRequiredVersion_851_minor := uint32(0)                             // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                          // Contextʳ
	GetContext_852_result := context                                      // Contextʳ
	ctx := GetContext_852_result                                          // Contextʳ
	v := ϟa.Values.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // GLintˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)               // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                          // Uniform
	uniform.Type = GLenum_GL_INT
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_851_major, minRequiredVersion_851_minor, context, GetContext_852_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform1ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_853_major := uint32(3) // u32
	minRequiredVersion_853_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_853_major, minRequiredVersion_853_minor
	return nil
}
func (ϟa *GlUniform1uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_854_major := uint32(3) // u32
	minRequiredVersion_854_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_854_major, minRequiredVersion_854_minor
	return nil
}
func (ϟa *GlUniform2f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_855_major := uint32(2)    // u32
	minRequiredVersion_855_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_856_result := context             // Contextʳ
	ctx := GetContext_856_result                 // Contextʳ
	v := MakeVec2fˢ(uint64(1), ϟs)               // Vec2fˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).Write(Vec2f{Elements: [2]GLfloat{ϟa.Value0, ϟa.Value1}}, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_FLOAT_VEC2
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_855_major, minRequiredVersion_855_minor, context, GetContext_856_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_857_major := uint32(2)                                     // u32
	minRequiredVersion_857_minor := uint32(0)                                     // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                  // Contextʳ
	GetContext_858_result := context                                              // Contextʳ
	ctx := GetContext_858_result                                                  // Contextʳ
	v := Vec2fᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Vec2fˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_FLOAT_VEC2
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_857_major, minRequiredVersion_857_minor, context, GetContext_858_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform2i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_859_major := uint32(2)    // u32
	minRequiredVersion_859_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_860_result := context             // Contextʳ
	ctx := GetContext_860_result                 // Contextʳ
	v := MakeVec2iˢ(uint64(1), ϟs)               // Vec2iˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).Write(Vec2i{Elements: [2]GLint{ϟa.Value0, ϟa.Value1}}, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_INT_VEC2
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_859_major, minRequiredVersion_859_minor, context, GetContext_860_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform2iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_861_major := uint32(2)                                     // u32
	minRequiredVersion_861_minor := uint32(0)                                     // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                  // Contextʳ
	GetContext_862_result := context                                              // Contextʳ
	ctx := GetContext_862_result                                                  // Contextʳ
	v := Vec2iᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Vec2iˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_INT_VEC2
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_861_major, minRequiredVersion_861_minor, context, GetContext_862_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform2ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_863_major := uint32(3) // u32
	minRequiredVersion_863_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_863_major, minRequiredVersion_863_minor
	return nil
}
func (ϟa *GlUniform2uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_864_major := uint32(3) // u32
	minRequiredVersion_864_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_864_major, minRequiredVersion_864_minor
	return nil
}
func (ϟa *GlUniform3f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_865_major := uint32(2)    // u32
	minRequiredVersion_865_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_866_result := context             // Contextʳ
	ctx := GetContext_866_result                 // Contextʳ
	v := MakeVec3fˢ(uint64(1), ϟs)               // Vec3fˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).Write(Vec3f{Elements: [3]GLfloat{ϟa.Value0, ϟa.Value1, ϟa.Value2}}, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_FLOAT_VEC3
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_865_major, minRequiredVersion_865_minor, context, GetContext_866_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_867_major := uint32(2)                                     // u32
	minRequiredVersion_867_minor := uint32(0)                                     // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                  // Contextʳ
	GetContext_868_result := context                                              // Contextʳ
	ctx := GetContext_868_result                                                  // Contextʳ
	v := Vec3fᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Vec3fˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_FLOAT_VEC3
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_867_major, minRequiredVersion_867_minor, context, GetContext_868_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform3i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_869_major := uint32(2)    // u32
	minRequiredVersion_869_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_870_result := context             // Contextʳ
	ctx := GetContext_870_result                 // Contextʳ
	v := MakeVec3iˢ(uint64(1), ϟs)               // Vec3iˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).Write(Vec3i{Elements: [3]GLint{ϟa.Value0, ϟa.Value1, ϟa.Value2}}, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_INT_VEC3
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_869_major, minRequiredVersion_869_minor, context, GetContext_870_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform3iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_871_major := uint32(2)                                     // u32
	minRequiredVersion_871_minor := uint32(0)                                     // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                  // Contextʳ
	GetContext_872_result := context                                              // Contextʳ
	ctx := GetContext_872_result                                                  // Contextʳ
	v := Vec3iᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Vec3iˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_INT_VEC3
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_871_major, minRequiredVersion_871_minor, context, GetContext_872_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform3ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_873_major := uint32(3) // u32
	minRequiredVersion_873_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_873_major, minRequiredVersion_873_minor
	return nil
}
func (ϟa *GlUniform3uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_874_major := uint32(3) // u32
	minRequiredVersion_874_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_874_major, minRequiredVersion_874_minor
	return nil
}
func (ϟa *GlUniform4f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_875_major := uint32(2)    // u32
	minRequiredVersion_875_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_876_result := context             // Contextʳ
	ctx := GetContext_876_result                 // Contextʳ
	v := MakeVec4fˢ(uint64(1), ϟs)               // Vec4fˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).Write(Vec4f{Elements: [4]GLfloat{ϟa.Value0, ϟa.Value1, ϟa.Value2, ϟa.Value3}}, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_FLOAT_VEC4
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_875_major, minRequiredVersion_875_minor, context, GetContext_876_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_877_major := uint32(2)                                     // u32
	minRequiredVersion_877_minor := uint32(0)                                     // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                  // Contextʳ
	GetContext_878_result := context                                              // Contextʳ
	ctx := GetContext_878_result                                                  // Contextʳ
	v := Vec4fᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Vec4fˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_FLOAT_VEC4
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_877_major, minRequiredVersion_877_minor, context, GetContext_878_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform4i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_879_major := uint32(2)    // u32
	minRequiredVersion_879_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_880_result := context             // Contextʳ
	ctx := GetContext_880_result                 // Contextʳ
	v := MakeVec4iˢ(uint64(1), ϟs)               // Vec4iˢ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).Write(Vec4i{Elements: [4]GLint{ϟa.Value0, ϟa.Value1, ϟa.Value2, ϟa.Value3}}, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_INT_VEC4
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_879_major, minRequiredVersion_879_minor, context, GetContext_880_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform4iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_881_major := uint32(2)                                     // u32
	minRequiredVersion_881_minor := uint32(0)                                     // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                  // Contextʳ
	GetContext_882_result := context                                              // Contextʳ
	ctx := GetContext_882_result                                                  // Contextʳ
	v := Vec4iᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Vec4iˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_INT_VEC4
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_881_major, minRequiredVersion_881_minor, context, GetContext_882_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniform4ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_883_major := uint32(3) // u32
	minRequiredVersion_883_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_883_major, minRequiredVersion_883_minor
	return nil
}
func (ϟa *GlUniform4uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_884_major := uint32(3) // u32
	minRequiredVersion_884_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_884_major, minRequiredVersion_884_minor
	return nil
}
func (ϟa *GlUniformBlockBinding) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_885_major := uint32(3) // u32
	minRequiredVersion_885_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_885_major, minRequiredVersion_885_minor
	return nil
}
func (ϟa *GlUniformMatrix2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_886_major := uint32(2)                                     // u32
	minRequiredVersion_886_minor := uint32(0)                                     // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                  // Contextʳ
	GetContext_887_result := context                                              // Contextʳ
	ctx := GetContext_887_result                                                  // Contextʳ
	v := Mat2fᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Mat2fˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_FLOAT_MAT2
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_886_major, minRequiredVersion_886_minor, context, GetContext_887_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniformMatrix2x3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_888_major := uint32(3) // u32
	minRequiredVersion_888_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_888_major, minRequiredVersion_888_minor
	return nil
}
func (ϟa *GlUniformMatrix2x4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_889_major := uint32(3) // u32
	minRequiredVersion_889_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_889_major, minRequiredVersion_889_minor
	return nil
}
func (ϟa *GlUniformMatrix3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_890_major := uint32(2)                                     // u32
	minRequiredVersion_890_minor := uint32(0)                                     // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                  // Contextʳ
	GetContext_891_result := context                                              // Contextʳ
	ctx := GetContext_891_result                                                  // Contextʳ
	v := Mat3fᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Mat3fˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_FLOAT_MAT3
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_890_major, minRequiredVersion_890_minor, context, GetContext_891_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniformMatrix3x2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_892_major := uint32(3) // u32
	minRequiredVersion_892_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_892_major, minRequiredVersion_892_minor
	return nil
}
func (ϟa *GlUniformMatrix3x4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_893_major := uint32(3) // u32
	minRequiredVersion_893_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_893_major, minRequiredVersion_893_minor
	return nil
}
func (ϟa *GlUniformMatrix4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_894_major := uint32(2)                                     // u32
	minRequiredVersion_894_minor := uint32(0)                                     // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                  // Contextʳ
	GetContext_895_result := context                                              // Contextʳ
	ctx := GetContext_895_result                                                  // Contextʳ
	v := Mat4fᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Mat4fˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Value = AsU8ˢ(v, ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
	program.Uniforms[ϟa.Location] = uniform
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_894_major, minRequiredVersion_894_minor, context, GetContext_895_result, ctx, v, program, uniform
	return nil
}
func (ϟa *GlUniformMatrix4x2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_896_major := uint32(3) // u32
	minRequiredVersion_896_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_896_major, minRequiredVersion_896_minor
	return nil
}
func (ϟa *GlUniformMatrix4x3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_897_major := uint32(3) // u32
	minRequiredVersion_897_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_897_major, minRequiredVersion_897_minor
	return nil
}
func (ϟa *GlUseProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_898_major := uint32(2)    // u32
	minRequiredVersion_898_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_899_result := context             // Contextʳ
	ctx := GetContext_899_result                 // Contextʳ
	ctx.BoundProgram = ϟa.Program
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_898_major, minRequiredVersion_898_minor, context, GetContext_899_result, ctx
	return nil
}
func (ϟa *GlUseProgramStages) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_900_major := uint32(3)                                                                                                                                        // u32
	minRequiredVersion_900_minor := uint32(1)                                                                                                                                        // u32
	supportsBits_901_seenBits := ϟa.Stages                                                                                                                                           // GLbitfield
	supportsBits_901_validBits := (GLbitfield_GL_ALL_SHADER_BITS) | ((GLbitfield_GL_COMPUTE_SHADER_BIT) | ((GLbitfield_GL_FRAGMENT_SHADER_BIT) | (GLbitfield_GL_VERTEX_SHADER_BIT))) // GLbitfield
	if (GLbitfield_GL_ALL_SHADER_BITS)&(ϟa.Stages) != 0 {
	}
	if (GLbitfield_GL_COMPUTE_SHADER_BIT)&(ϟa.Stages) != 0 {
	}
	if (GLbitfield_GL_FRAGMENT_SHADER_BIT)&(ϟa.Stages) != 0 {
	}
	if (GLbitfield_GL_VERTEX_SHADER_BIT)&(ϟa.Stages) != 0 {
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = minRequiredVersion_900_major, minRequiredVersion_900_minor, supportsBits_901_seenBits, supportsBits_901_validBits
	return nil
}
func (ϟa *GlValidateProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_902_major := uint32(2) // u32
	minRequiredVersion_902_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_902_major, minRequiredVersion_902_minor
	return nil
}
func (ϟa *GlValidateProgramPipeline) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_903_major := uint32(3) // u32
	minRequiredVersion_903_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_903_major, minRequiredVersion_903_minor
	return nil
}
func (ϟa *GlCullFace) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_904_major := uint32(2) // u32
	minRequiredVersion_904_minor := uint32(0) // u32
	switch ϟa.Mode {
	case GLenum_GL_BACK, GLenum_GL_FRONT, GLenum_GL_FRONT_AND_BACK:
	default:
		v := ϟa.Mode
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_906_result := context             // Contextʳ
	ctx := GetContext_906_result                 // Contextʳ
	ctx.Rasterizing.CullFace = ϟa.Mode
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_904_major, minRequiredVersion_904_minor, context, GetContext_906_result, ctx
	return nil
}
func (ϟa *GlDepthRangef) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_907_major := uint32(2)    // u32
	minRequiredVersion_907_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_908_result := context             // Contextʳ
	ctx := GetContext_908_result                 // Contextʳ
	ctx.Rasterizing.DepthNear = ϟa.Near
	ctx.Rasterizing.DepthFar = ϟa.Far
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_907_major, minRequiredVersion_907_minor, context, GetContext_908_result, ctx
	return nil
}
func (ϟa *GlFrontFace) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_909_major := uint32(2) // u32
	minRequiredVersion_909_minor := uint32(0) // u32
	switch ϟa.Orientation {
	case GLenum_GL_CCW, GLenum_GL_CW:
	default:
		v := ϟa.Orientation
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_911_result := context             // Contextʳ
	ctx := GetContext_911_result                 // Contextʳ
	ctx.Rasterizing.FrontFace = ϟa.Orientation
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_909_major, minRequiredVersion_909_minor, context, GetContext_911_result, ctx
	return nil
}
func (ϟa *GlGetMultisamplefv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_912_major := uint32(3) // u32
	minRequiredVersion_912_minor := uint32(1) // u32
	switch ϟa.Pname {
	case GLenum_GL_SAMPLE_POSITION:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_912_major, minRequiredVersion_912_minor
	return nil
}
func (ϟa *GlLineWidth) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_914_major := uint32(2)    // u32
	minRequiredVersion_914_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_915_result := context             // Contextʳ
	ctx := GetContext_915_result                 // Contextʳ
	ctx.Rasterizing.LineWidth = ϟa.Width
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_914_major, minRequiredVersion_914_minor, context, GetContext_915_result, ctx
	return nil
}
func (ϟa *GlMinSampleShading) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_916_major := uint32(3) // u32
	minRequiredVersion_916_minor := uint32(2) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_916_major, minRequiredVersion_916_minor
	return nil
}
func (ϟa *GlPolygonOffset) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_917_major := uint32(2)    // u32
	minRequiredVersion_917_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_918_result := context             // Contextʳ
	ctx := GetContext_918_result                 // Contextʳ
	ctx.Rasterizing.PolygonOffsetUnits = ϟa.Units
	ctx.Rasterizing.PolygonOffsetFactor = ϟa.ScaleFactor
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_917_major, minRequiredVersion_917_minor, context, GetContext_918_result, ctx
	return nil
}
func (ϟa *GlViewport) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_919_major := uint32(2)    // u32
	minRequiredVersion_919_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_920_result := context             // Contextʳ
	ctx := GetContext_920_result                 // Contextʳ
	ctx.Rasterizing.Viewport = Rect{X: ϟa.X, Y: ϟa.Y, Width: ϟa.Width, Height: ϟa.Height}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_919_major, minRequiredVersion_919_minor, context, GetContext_920_result, ctx
	return nil
}
func (ϟa *GlGetBooleani_v) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_921_major := uint32(3) // u32
	minRequiredVersion_921_minor := uint32(1) // u32
	switch ϟa.Target {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_DITHER, GLenum_GL_DRAW_BUFFER, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE, GLenum_GL_VIEWPORT:
	case GLenum_GL_CONTEXT_FLAGS, GLenum_GL_CONTEXT_ROBUST_ACCESS, GLenum_GL_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_DEBUG_LOGGED_MESSAGES, GLenum_GL_DEBUG_NEXT_LOGGED_MESSAGE_LENGTH, GLenum_GL_FRAGMENT_INTERPOLATION_OFFSET_BITS, GLenum_GL_LAYER_PROVOKING_VERTEX, GLenum_GL_MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_IMAGE_UNIFORMS, GLenum_GL_MAX_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_MAX_DEBUG_LOGGED_MESSAGES, GLenum_GL_MAX_DEBUG_MESSAGE_LENGTH, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_FRAGMENT_IMAGE_UNIFORMS, GLenum_GL_MAX_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MAX_FRAMEBUFFER_LAYERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_GEOMETRY_IMAGE_UNIFORMS, GLenum_GL_MAX_GEOMETRY_INPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_VERTICES, GLenum_GL_MAX_GEOMETRY_SHADER_INVOCATIONS, GLenum_GL_MAX_GEOMETRY_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_GEOMETRY_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_UNIFORM_BLOCKS, GLenum_GL_MAX_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_LABEL_LENGTH, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_CONTROL_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_CONTROL_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_EVALUATION_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_EVALUATION_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_GEN_LEVEL, GLenum_GL_MAX_TESS_PATCH_COMPONENTS, GLenum_GL_MAX_TEXTURE_BUFFER_SIZE, GLenum_GL_MAX_VERTEX_IMAGE_UNIFORMS, GLenum_GL_MIN_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MIN_SAMPLE_SHADING_VALUE, GLenum_GL_MULTISAMPLE_LINE_WIDTH_RANGE, GLenum_GL_PATCH_VERTICES, GLenum_GL_PRIMITIVE_BOUNDING_BOX, GLenum_GL_RESET_NOTIFICATION_STRATEGY, GLenum_GL_SAMPLE_SHADING, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_BINDING_BUFFER, GLenum_GL_TEXTURE_BINDING_CUBE_MAP_ARRAY, GLenum_GL_TEXTURE_BUFFER_BINDING, GLenum_GL_TEXTURE_BUFFER_OFFSET_ALIGNMENT:
		minRequiredVersion_922_major := uint32(3) // u32
		minRequiredVersion_922_minor := uint32(2) // u32
		_, _ = minRequiredVersion_922_major, minRequiredVersion_922_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_921_major, minRequiredVersion_921_minor
	return nil
}
func (ϟa *GlGetBooleanv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_924_major := uint32(2) // u32
	minRequiredVersion_924_minor := uint32(0) // u32
	switch ϟa.Param {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DITHER, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_VIEWPORT:
	case GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_DRAW_BUFFER, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING:
		minRequiredVersion_925_major := uint32(3) // u32
		minRequiredVersion_925_minor := uint32(0) // u32
		_, _ = minRequiredVersion_925_major, minRequiredVersion_925_minor
	case GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE:
		minRequiredVersion_926_major := uint32(3) // u32
		minRequiredVersion_926_minor := uint32(1) // u32
		_, _ = minRequiredVersion_926_major, minRequiredVersion_926_minor
	case GLenum_GL_CONTEXT_FLAGS, GLenum_GL_CONTEXT_ROBUST_ACCESS, GLenum_GL_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_DEBUG_LOGGED_MESSAGES, GLenum_GL_DEBUG_NEXT_LOGGED_MESSAGE_LENGTH, GLenum_GL_FRAGMENT_INTERPOLATION_OFFSET_BITS, GLenum_GL_LAYER_PROVOKING_VERTEX, GLenum_GL_MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_IMAGE_UNIFORMS, GLenum_GL_MAX_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_MAX_DEBUG_LOGGED_MESSAGES, GLenum_GL_MAX_DEBUG_MESSAGE_LENGTH, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_FRAGMENT_IMAGE_UNIFORMS, GLenum_GL_MAX_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MAX_FRAMEBUFFER_LAYERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_GEOMETRY_IMAGE_UNIFORMS, GLenum_GL_MAX_GEOMETRY_INPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_VERTICES, GLenum_GL_MAX_GEOMETRY_SHADER_INVOCATIONS, GLenum_GL_MAX_GEOMETRY_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_GEOMETRY_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_UNIFORM_BLOCKS, GLenum_GL_MAX_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_LABEL_LENGTH, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_CONTROL_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_CONTROL_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_EVALUATION_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_EVALUATION_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_GEN_LEVEL, GLenum_GL_MAX_TESS_PATCH_COMPONENTS, GLenum_GL_MAX_TEXTURE_BUFFER_SIZE, GLenum_GL_MAX_VERTEX_IMAGE_UNIFORMS, GLenum_GL_MIN_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MIN_SAMPLE_SHADING_VALUE, GLenum_GL_MULTISAMPLE_LINE_WIDTH_RANGE, GLenum_GL_PATCH_VERTICES, GLenum_GL_PRIMITIVE_BOUNDING_BOX, GLenum_GL_RESET_NOTIFICATION_STRATEGY, GLenum_GL_SAMPLE_SHADING, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_BINDING_BUFFER, GLenum_GL_TEXTURE_BINDING_CUBE_MAP_ARRAY, GLenum_GL_TEXTURE_BUFFER_BINDING, GLenum_GL_TEXTURE_BUFFER_OFFSET_ALIGNMENT:
		minRequiredVersion_927_major := uint32(3) // u32
		minRequiredVersion_927_minor := uint32(2) // u32
		_, _ = minRequiredVersion_927_major, minRequiredVersion_927_minor
	default:
		v := ϟa.Param
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	v := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs) // GLbooleanˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                                                // Contextʳ
	GetContext_929_result := context                                                                            // Contextʳ
	ctx := GetContext_929_result                                                                                // Contextʳ
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
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	default:
		v := ϟa.Param
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	_, _, _, _, _, _ = minRequiredVersion_924_major, minRequiredVersion_924_minor, v, context, GetContext_929_result, ctx
	return nil
}
func (ϟa *GlGetFloatv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_930_major := uint32(2) // u32
	minRequiredVersion_930_minor := uint32(0) // u32
	switch ϟa.Param {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DITHER, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_VIEWPORT:
	case GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_DRAW_BUFFER, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING:
		minRequiredVersion_931_major := uint32(3) // u32
		minRequiredVersion_931_minor := uint32(0) // u32
		_, _ = minRequiredVersion_931_major, minRequiredVersion_931_minor
	case GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE:
		minRequiredVersion_932_major := uint32(3) // u32
		minRequiredVersion_932_minor := uint32(1) // u32
		_, _ = minRequiredVersion_932_major, minRequiredVersion_932_minor
	case GLenum_GL_CONTEXT_FLAGS, GLenum_GL_CONTEXT_ROBUST_ACCESS, GLenum_GL_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_DEBUG_LOGGED_MESSAGES, GLenum_GL_DEBUG_NEXT_LOGGED_MESSAGE_LENGTH, GLenum_GL_FRAGMENT_INTERPOLATION_OFFSET_BITS, GLenum_GL_LAYER_PROVOKING_VERTEX, GLenum_GL_MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_IMAGE_UNIFORMS, GLenum_GL_MAX_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_MAX_DEBUG_LOGGED_MESSAGES, GLenum_GL_MAX_DEBUG_MESSAGE_LENGTH, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_FRAGMENT_IMAGE_UNIFORMS, GLenum_GL_MAX_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MAX_FRAMEBUFFER_LAYERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_GEOMETRY_IMAGE_UNIFORMS, GLenum_GL_MAX_GEOMETRY_INPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_VERTICES, GLenum_GL_MAX_GEOMETRY_SHADER_INVOCATIONS, GLenum_GL_MAX_GEOMETRY_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_GEOMETRY_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_UNIFORM_BLOCKS, GLenum_GL_MAX_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_LABEL_LENGTH, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_CONTROL_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_CONTROL_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_EVALUATION_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_EVALUATION_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_GEN_LEVEL, GLenum_GL_MAX_TESS_PATCH_COMPONENTS, GLenum_GL_MAX_TEXTURE_BUFFER_SIZE, GLenum_GL_MAX_VERTEX_IMAGE_UNIFORMS, GLenum_GL_MIN_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MIN_SAMPLE_SHADING_VALUE, GLenum_GL_MULTISAMPLE_LINE_WIDTH_RANGE, GLenum_GL_PATCH_VERTICES, GLenum_GL_PRIMITIVE_BOUNDING_BOX, GLenum_GL_RESET_NOTIFICATION_STRATEGY, GLenum_GL_SAMPLE_SHADING, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_BINDING_BUFFER, GLenum_GL_TEXTURE_BINDING_CUBE_MAP_ARRAY, GLenum_GL_TEXTURE_BUFFER_BINDING, GLenum_GL_TEXTURE_BUFFER_OFFSET_ALIGNMENT:
		minRequiredVersion_933_major := uint32(3) // u32
		minRequiredVersion_933_minor := uint32(2) // u32
		_, _ = minRequiredVersion_933_major, minRequiredVersion_933_minor
	default:
		v := ϟa.Param
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	v := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs) // GLfloatˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                                                // Contextʳ
	GetContext_935_result := context                                                                            // Contextʳ
	ctx := GetContext_935_result                                                                                // Contextʳ
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
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(1), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(1), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_ALIASED_POINT_SIZE_RANGE:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(1), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(1), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	default:
		v := ϟa.Param
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	_, _, _, _, _, _ = minRequiredVersion_930_major, minRequiredVersion_930_minor, v, context, GetContext_935_result, ctx
	return nil
}
func (ϟa *GlGetInteger64i_v) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_936_major := uint32(3) // u32
	minRequiredVersion_936_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DITHER, GLenum_GL_DRAW_BUFFER, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING, GLenum_GL_VIEWPORT:
	case GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE:
		minRequiredVersion_937_major := uint32(3) // u32
		minRequiredVersion_937_minor := uint32(1) // u32
		_, _ = minRequiredVersion_937_major, minRequiredVersion_937_minor
	case GLenum_GL_CONTEXT_FLAGS, GLenum_GL_CONTEXT_ROBUST_ACCESS, GLenum_GL_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_DEBUG_LOGGED_MESSAGES, GLenum_GL_DEBUG_NEXT_LOGGED_MESSAGE_LENGTH, GLenum_GL_FRAGMENT_INTERPOLATION_OFFSET_BITS, GLenum_GL_LAYER_PROVOKING_VERTEX, GLenum_GL_MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_IMAGE_UNIFORMS, GLenum_GL_MAX_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_MAX_DEBUG_LOGGED_MESSAGES, GLenum_GL_MAX_DEBUG_MESSAGE_LENGTH, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_FRAGMENT_IMAGE_UNIFORMS, GLenum_GL_MAX_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MAX_FRAMEBUFFER_LAYERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_GEOMETRY_IMAGE_UNIFORMS, GLenum_GL_MAX_GEOMETRY_INPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_VERTICES, GLenum_GL_MAX_GEOMETRY_SHADER_INVOCATIONS, GLenum_GL_MAX_GEOMETRY_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_GEOMETRY_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_UNIFORM_BLOCKS, GLenum_GL_MAX_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_LABEL_LENGTH, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_CONTROL_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_CONTROL_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_EVALUATION_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_EVALUATION_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_GEN_LEVEL, GLenum_GL_MAX_TESS_PATCH_COMPONENTS, GLenum_GL_MAX_TEXTURE_BUFFER_SIZE, GLenum_GL_MAX_VERTEX_IMAGE_UNIFORMS, GLenum_GL_MIN_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MIN_SAMPLE_SHADING_VALUE, GLenum_GL_MULTISAMPLE_LINE_WIDTH_RANGE, GLenum_GL_PATCH_VERTICES, GLenum_GL_PRIMITIVE_BOUNDING_BOX, GLenum_GL_RESET_NOTIFICATION_STRATEGY, GLenum_GL_SAMPLE_SHADING, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_BINDING_BUFFER, GLenum_GL_TEXTURE_BINDING_CUBE_MAP_ARRAY, GLenum_GL_TEXTURE_BUFFER_BINDING, GLenum_GL_TEXTURE_BUFFER_OFFSET_ALIGNMENT:
		minRequiredVersion_938_major := uint32(3) // u32
		minRequiredVersion_938_minor := uint32(2) // u32
		_, _ = minRequiredVersion_938_major, minRequiredVersion_938_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_936_major, minRequiredVersion_936_minor
	return nil
}
func (ϟa *GlGetInteger64v) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_940_major := uint32(3) // u32
	minRequiredVersion_940_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DITHER, GLenum_GL_DRAW_BUFFER, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING, GLenum_GL_VIEWPORT:
	case GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE:
		minRequiredVersion_941_major := uint32(3) // u32
		minRequiredVersion_941_minor := uint32(1) // u32
		_, _ = minRequiredVersion_941_major, minRequiredVersion_941_minor
	case GLenum_GL_CONTEXT_FLAGS, GLenum_GL_CONTEXT_ROBUST_ACCESS, GLenum_GL_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_DEBUG_LOGGED_MESSAGES, GLenum_GL_DEBUG_NEXT_LOGGED_MESSAGE_LENGTH, GLenum_GL_FRAGMENT_INTERPOLATION_OFFSET_BITS, GLenum_GL_LAYER_PROVOKING_VERTEX, GLenum_GL_MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_IMAGE_UNIFORMS, GLenum_GL_MAX_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_MAX_DEBUG_LOGGED_MESSAGES, GLenum_GL_MAX_DEBUG_MESSAGE_LENGTH, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_FRAGMENT_IMAGE_UNIFORMS, GLenum_GL_MAX_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MAX_FRAMEBUFFER_LAYERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_GEOMETRY_IMAGE_UNIFORMS, GLenum_GL_MAX_GEOMETRY_INPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_VERTICES, GLenum_GL_MAX_GEOMETRY_SHADER_INVOCATIONS, GLenum_GL_MAX_GEOMETRY_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_GEOMETRY_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_UNIFORM_BLOCKS, GLenum_GL_MAX_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_LABEL_LENGTH, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_CONTROL_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_CONTROL_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_EVALUATION_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_EVALUATION_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_GEN_LEVEL, GLenum_GL_MAX_TESS_PATCH_COMPONENTS, GLenum_GL_MAX_TEXTURE_BUFFER_SIZE, GLenum_GL_MAX_VERTEX_IMAGE_UNIFORMS, GLenum_GL_MIN_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MIN_SAMPLE_SHADING_VALUE, GLenum_GL_MULTISAMPLE_LINE_WIDTH_RANGE, GLenum_GL_PATCH_VERTICES, GLenum_GL_PRIMITIVE_BOUNDING_BOX, GLenum_GL_RESET_NOTIFICATION_STRATEGY, GLenum_GL_SAMPLE_SHADING, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_BINDING_BUFFER, GLenum_GL_TEXTURE_BINDING_CUBE_MAP_ARRAY, GLenum_GL_TEXTURE_BUFFER_BINDING, GLenum_GL_TEXTURE_BUFFER_OFFSET_ALIGNMENT:
		minRequiredVersion_942_major := uint32(3) // u32
		minRequiredVersion_942_minor := uint32(2) // u32
		_, _ = minRequiredVersion_942_major, minRequiredVersion_942_minor
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_940_major, minRequiredVersion_940_minor
	return nil
}
func (ϟa *GlGetIntegeri_v) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_944_major := uint32(3) // u32
	minRequiredVersion_944_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DITHER, GLenum_GL_DRAW_BUFFER, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING, GLenum_GL_VIEWPORT:
	case GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE:
		minRequiredVersion_945_major := uint32(3) // u32
		minRequiredVersion_945_minor := uint32(1) // u32
		_, _ = minRequiredVersion_945_major, minRequiredVersion_945_minor
	case GLenum_GL_CONTEXT_FLAGS, GLenum_GL_CONTEXT_ROBUST_ACCESS, GLenum_GL_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_DEBUG_LOGGED_MESSAGES, GLenum_GL_DEBUG_NEXT_LOGGED_MESSAGE_LENGTH, GLenum_GL_FRAGMENT_INTERPOLATION_OFFSET_BITS, GLenum_GL_LAYER_PROVOKING_VERTEX, GLenum_GL_MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_IMAGE_UNIFORMS, GLenum_GL_MAX_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_MAX_DEBUG_LOGGED_MESSAGES, GLenum_GL_MAX_DEBUG_MESSAGE_LENGTH, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_FRAGMENT_IMAGE_UNIFORMS, GLenum_GL_MAX_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MAX_FRAMEBUFFER_LAYERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_GEOMETRY_IMAGE_UNIFORMS, GLenum_GL_MAX_GEOMETRY_INPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_VERTICES, GLenum_GL_MAX_GEOMETRY_SHADER_INVOCATIONS, GLenum_GL_MAX_GEOMETRY_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_GEOMETRY_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_UNIFORM_BLOCKS, GLenum_GL_MAX_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_LABEL_LENGTH, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_CONTROL_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_CONTROL_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_EVALUATION_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_EVALUATION_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_GEN_LEVEL, GLenum_GL_MAX_TESS_PATCH_COMPONENTS, GLenum_GL_MAX_TEXTURE_BUFFER_SIZE, GLenum_GL_MAX_VERTEX_IMAGE_UNIFORMS, GLenum_GL_MIN_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MIN_SAMPLE_SHADING_VALUE, GLenum_GL_MULTISAMPLE_LINE_WIDTH_RANGE, GLenum_GL_PATCH_VERTICES, GLenum_GL_PRIMITIVE_BOUNDING_BOX, GLenum_GL_RESET_NOTIFICATION_STRATEGY, GLenum_GL_SAMPLE_SHADING, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_BINDING_BUFFER, GLenum_GL_TEXTURE_BINDING_CUBE_MAP_ARRAY, GLenum_GL_TEXTURE_BUFFER_BINDING, GLenum_GL_TEXTURE_BUFFER_OFFSET_ALIGNMENT:
		minRequiredVersion_946_major := uint32(3) // u32
		minRequiredVersion_946_minor := uint32(2) // u32
		_, _ = minRequiredVersion_946_major, minRequiredVersion_946_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_944_major, minRequiredVersion_944_minor
	return nil
}
func (ϟa *GlGetIntegerv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_948_major := uint32(2) // u32
	minRequiredVersion_948_minor := uint32(0) // u32
	switch ϟa.Param {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DITHER, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_VIEWPORT:
	case GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_DRAW_BUFFER, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING:
		minRequiredVersion_949_major := uint32(3) // u32
		minRequiredVersion_949_minor := uint32(0) // u32
		_, _ = minRequiredVersion_949_major, minRequiredVersion_949_minor
	case GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE:
		minRequiredVersion_950_major := uint32(3) // u32
		minRequiredVersion_950_minor := uint32(1) // u32
		_, _ = minRequiredVersion_950_major, minRequiredVersion_950_minor
	case GLenum_GL_CONTEXT_FLAGS, GLenum_GL_CONTEXT_ROBUST_ACCESS, GLenum_GL_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_DEBUG_LOGGED_MESSAGES, GLenum_GL_DEBUG_NEXT_LOGGED_MESSAGE_LENGTH, GLenum_GL_FRAGMENT_INTERPOLATION_OFFSET_BITS, GLenum_GL_LAYER_PROVOKING_VERTEX, GLenum_GL_MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_IMAGE_UNIFORMS, GLenum_GL_MAX_DEBUG_GROUP_STACK_DEPTH, GLenum_GL_MAX_DEBUG_LOGGED_MESSAGES, GLenum_GL_MAX_DEBUG_MESSAGE_LENGTH, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_FRAGMENT_IMAGE_UNIFORMS, GLenum_GL_MAX_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MAX_FRAMEBUFFER_LAYERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTERS, GLenum_GL_MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_GEOMETRY_IMAGE_UNIFORMS, GLenum_GL_MAX_GEOMETRY_INPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_OUTPUT_VERTICES, GLenum_GL_MAX_GEOMETRY_SHADER_INVOCATIONS, GLenum_GL_MAX_GEOMETRY_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_GEOMETRY_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_GEOMETRY_UNIFORM_BLOCKS, GLenum_GL_MAX_GEOMETRY_UNIFORM_COMPONENTS, GLenum_GL_MAX_LABEL_LENGTH, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_CONTROL_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_CONTROL_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_CONTROL_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTERS, GLenum_GL_MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_TESS_EVALUATION_IMAGE_UNIFORMS, GLenum_GL_MAX_TESS_EVALUATION_INPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_OUTPUT_COMPONENTS, GLenum_GL_MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_BLOCKS, GLenum_GL_MAX_TESS_EVALUATION_UNIFORM_COMPONENTS, GLenum_GL_MAX_TESS_GEN_LEVEL, GLenum_GL_MAX_TESS_PATCH_COMPONENTS, GLenum_GL_MAX_TEXTURE_BUFFER_SIZE, GLenum_GL_MAX_VERTEX_IMAGE_UNIFORMS, GLenum_GL_MIN_FRAGMENT_INTERPOLATION_OFFSET, GLenum_GL_MIN_SAMPLE_SHADING_VALUE, GLenum_GL_MULTISAMPLE_LINE_WIDTH_RANGE, GLenum_GL_PATCH_VERTICES, GLenum_GL_PRIMITIVE_BOUNDING_BOX, GLenum_GL_RESET_NOTIFICATION_STRATEGY, GLenum_GL_SAMPLE_SHADING, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_BINDING_BUFFER, GLenum_GL_TEXTURE_BINDING_CUBE_MAP_ARRAY, GLenum_GL_TEXTURE_BUFFER_BINDING, GLenum_GL_TEXTURE_BUFFER_OFFSET_ALIGNMENT:
		minRequiredVersion_951_major := uint32(3) // u32
		minRequiredVersion_951_minor := uint32(2) // u32
		_, _ = minRequiredVersion_951_major, minRequiredVersion_951_minor
	default:
		v := ϟa.Param
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	v := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs) // GLintˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                                                // Contextʳ
	GetContext_953_result := context                                                                            // Contextʳ
	ctx := GetContext_953_result                                                                                // Contextʳ
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
		v.Index(uint64(0), ϟs).Write(GLint(ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Bindings.Get(GLenum_GL_TEXTURE_2D)), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_TEXTURE_BINDING_CUBE_MAP:
		v.Index(uint64(0), ϟs).Write(GLint(ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Bindings.Get(GLenum_GL_TEXTURE_CUBE_MAP)), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_GENERATE_MIPMAP_HINT:
		v.Index(uint64(0), ϟs).Write(GLint(ctx.GenerateMipmapHint), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil) // any
		v.Index(uint64(0), ϟs).Write(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil) // any
		v.Index(uint64(0), ϟs).Write(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil) // any
		v.Index(uint64(0), ϟs).Write(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_MAX_RENDERBUFFER_SIZE:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil) // any
		v.Index(uint64(0), ϟs).Write(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_MAX_TEXTURE_IMAGE_UNITS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil) // any
		v.Index(uint64(0), ϟs).Write(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_MAX_TEXTURE_SIZE:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil) // any
		v.Index(uint64(0), ϟs).Write(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_MAX_VARYING_VECTORS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil) // any
		v.Index(uint64(0), ϟs).Write(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_MAX_VERTEX_ATTRIBS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil) // any
		v.Index(uint64(0), ϟs).Write(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil) // any
		v.Index(uint64(0), ϟs).Write(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_MAX_VIEWPORT_DIMS:
		max_width := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)  // any
		max_height := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(1), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil) // any
		v.Index(uint64(0), ϟs).Write(max_width, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(1), ϟs).Write(max_height, ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = max_width, max_height
	case GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil) // any
		v.Index(uint64(0), ϟs).Write(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_NUM_SHADER_BINARY_FORMATS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil) // any
		v.Index(uint64(0), ϟs).Write(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_PACK_ALIGNMENT:
		v.Index(uint64(0), ϟs).Write(ctx.PixelStorage.Get(GLenum_GL_PACK_ALIGNMENT), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_UNPACK_ALIGNMENT:
		v.Index(uint64(0), ϟs).Write(ctx.PixelStorage.Get(GLenum_GL_UNPACK_ALIGNMENT), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_ALPHA_BITS:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_BLUE_BITS:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_GREEN_BITS:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_RED_BITS:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_DEPTH_BITS:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_SAMPLE_BUFFERS:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_SAMPLES:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_STENCIL_BITS:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_SUBPIXEL_BITS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil) // any
		v.Index(uint64(0), ϟs).Write(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_GPU_DISJOINT_EXT:
		v.Index(uint64(0), ϟs).Write(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟa, ϟs, ϟd, ϟl, ϟb}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	default:
		v := ϟa.Param
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	_, _, _, _, _, _ = minRequiredVersion_948_major, minRequiredVersion_948_minor, v, context, GetContext_953_result, ctx
	return nil
}
func (ϟa *GlGetInternalformativ) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_954_major := uint32(3) // u32
	minRequiredVersion_954_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_RENDERBUFFER:
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_955_major := uint32(3) // u32
		minRequiredVersion_955_minor := uint32(1) // u32
		_, _ = minRequiredVersion_955_major, minRequiredVersion_955_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY:
		minRequiredVersion_956_major := uint32(3) // u32
		minRequiredVersion_956_minor := uint32(2) // u32
		_, _ = minRequiredVersion_956_major, minRequiredVersion_956_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Internalformat {
	case GLenum_GL_DEPTH24_STENCIL8, GLenum_GL_DEPTH32F_STENCIL8, GLenum_GL_DEPTH_COMPONENT16, GLenum_GL_DEPTH_COMPONENT24, GLenum_GL_DEPTH_COMPONENT32F, GLenum_GL_R11F_G11F_B10F, GLenum_GL_R16F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32F, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_R8_SNORM, GLenum_GL_RG16F, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32F, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RG8_SNORM, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB16F, GLenum_GL_RGB16I, GLenum_GL_RGB16UI, GLenum_GL_RGB32F, GLenum_GL_RGB32I, GLenum_GL_RGB32UI, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGB8, GLenum_GL_RGB8I, GLenum_GL_RGB8UI, GLenum_GL_RGB8_SNORM, GLenum_GL_RGB9_E5, GLenum_GL_RGBA16F, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32F, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA4, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_RGBA8_SNORM, GLenum_GL_SRGB8, GLenum_GL_SRGB8_ALPHA8:
	case GLenum_GL_STENCIL_INDEX8:
		minRequiredVersion_958_major := uint32(3) // u32
		minRequiredVersion_958_minor := uint32(2) // u32
		_, _ = minRequiredVersion_958_major, minRequiredVersion_958_minor
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
	_, _ = minRequiredVersion_954_major, minRequiredVersion_954_minor
	return nil
}
func (ϟa *GlGetString) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_961_major := uint32(2) // u32
	minRequiredVersion_961_minor := uint32(0) // u32
	switch ϟa.Param {
	case GLenum_GL_EXTENSIONS, GLenum_GL_RENDERER, GLenum_GL_SHADING_LANGUAGE_VERSION, GLenum_GL_VENDOR, GLenum_GL_VERSION:
	default:
		v := ϟa.Param
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = GLubyteᶜᵖ{}
	_, _ = minRequiredVersion_961_major, minRequiredVersion_961_minor
	return nil
}
func (ϟa *GlGetStringi) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_963_major := uint32(3) // u32
	minRequiredVersion_963_minor := uint32(0) // u32
	switch ϟa.Name {
	case GLenum_GL_EXTENSIONS:
	default:
		v := ϟa.Name
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_963_major, minRequiredVersion_963_minor
	return nil
}
func (ϟa *GlIsEnabled) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_965_major := uint32(2) // u32
	minRequiredVersion_965_minor := uint32(0) // u32
	switch ϟa.Capability {
	case GLenum_GL_BLEND, GLenum_GL_CULL_FACE, GLenum_GL_DEPTH_TEST, GLenum_GL_DITHER, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SCISSOR_TEST, GLenum_GL_STENCIL_TEST:
	case GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_RASTERIZER_DISCARD:
		minRequiredVersion_966_major := uint32(3) // u32
		minRequiredVersion_966_minor := uint32(0) // u32
		_, _ = minRequiredVersion_966_major, minRequiredVersion_966_minor
	case GLenum_GL_DEBUG_OUTPUT, GLenum_GL_DEBUG_OUTPUT_SYNCHRONOUS, GLenum_GL_SAMPLE_MASK:
		minRequiredVersion_967_major := uint32(3) // u32
		minRequiredVersion_967_minor := uint32(2) // u32
		_, _ = minRequiredVersion_967_major, minRequiredVersion_967_minor
	default:
		v := ϟa.Capability
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_969_result := context             // Contextʳ
	ctx := GetContext_969_result                 // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Capabilities.Get(ϟa.Capability) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = minRequiredVersion_965_major, minRequiredVersion_965_minor, context, GetContext_969_result, ctx
	return nil
}
func (ϟa *GlIsEnabledi) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_970_major := uint32(3) // u32
	minRequiredVersion_970_minor := uint32(2) // u32
	switch ϟa.Target {
	case GLenum_GL_BLEND, GLenum_GL_CULL_FACE, GLenum_GL_DEBUG_OUTPUT, GLenum_GL_DEBUG_OUTPUT_SYNCHRONOUS, GLenum_GL_DEPTH_TEST, GLenum_GL_DITHER, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_MASK, GLenum_GL_SCISSOR_TEST, GLenum_GL_STENCIL_TEST:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_970_major, minRequiredVersion_970_minor
	return nil
}
func (ϟa *GlClientWaitSync) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_972_major := uint32(3)                           // u32
	minRequiredVersion_972_minor := uint32(0)                           // u32
	supportsBits_973_seenBits := ϟa.SyncFlags                           // GLbitfield
	supportsBits_973_validBits := GLbitfield_GL_SYNC_FLUSH_COMMANDS_BIT // GLbitfield
	if (GLbitfield_GL_SYNC_FLUSH_COMMANDS_BIT)&(ϟa.SyncFlags) != 0 {
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _, _, _ = minRequiredVersion_972_major, minRequiredVersion_972_minor, supportsBits_973_seenBits, supportsBits_973_validBits
	return nil
}
func (ϟa *GlDeleteSync) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_974_major := uint32(3) // u32
	minRequiredVersion_974_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_974_major, minRequiredVersion_974_minor
	return nil
}
func (ϟa *GlFenceSync) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_975_major := uint32(3) // u32
	minRequiredVersion_975_minor := uint32(0) // u32
	switch ϟa.Condition {
	case GLenum_GL_SYNC_GPU_COMMANDS_COMPLETE:
	default:
		v := ϟa.Condition
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_975_major, minRequiredVersion_975_minor
	return nil
}
func (ϟa *GlGetSynciv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_977_major := uint32(3) // u32
	minRequiredVersion_977_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_OBJECT_TYPE, GLenum_GL_SYNC_CONDITION, GLenum_GL_SYNC_FLAGS, GLenum_GL_SYNC_STATUS:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_977_major, minRequiredVersion_977_minor
	return nil
}
func (ϟa *GlIsSync) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_979_major := uint32(3) // u32
	minRequiredVersion_979_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_979_major, minRequiredVersion_979_minor
	return nil
}
func (ϟa *GlWaitSync) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_980_major := uint32(3) // u32
	minRequiredVersion_980_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_980_major, minRequiredVersion_980_minor
	return nil
}
func (ϟa *GlActiveTexture) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_981_major := uint32(2) // u32
	minRequiredVersion_981_minor := uint32(0) // u32
	switch ϟa.Unit {
	case GLenum_GL_TEXTURE0, GLenum_GL_TEXTURE1, GLenum_GL_TEXTURE10, GLenum_GL_TEXTURE11, GLenum_GL_TEXTURE12, GLenum_GL_TEXTURE13, GLenum_GL_TEXTURE14, GLenum_GL_TEXTURE15, GLenum_GL_TEXTURE16, GLenum_GL_TEXTURE17, GLenum_GL_TEXTURE18, GLenum_GL_TEXTURE19, GLenum_GL_TEXTURE2, GLenum_GL_TEXTURE20, GLenum_GL_TEXTURE21, GLenum_GL_TEXTURE22, GLenum_GL_TEXTURE23, GLenum_GL_TEXTURE24, GLenum_GL_TEXTURE25, GLenum_GL_TEXTURE26, GLenum_GL_TEXTURE27, GLenum_GL_TEXTURE28, GLenum_GL_TEXTURE29, GLenum_GL_TEXTURE3, GLenum_GL_TEXTURE30, GLenum_GL_TEXTURE31, GLenum_GL_TEXTURE4, GLenum_GL_TEXTURE5, GLenum_GL_TEXTURE6, GLenum_GL_TEXTURE7, GLenum_GL_TEXTURE8, GLenum_GL_TEXTURE9:
	default:
		v := ϟa.Unit
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_983_result := context             // Contextʳ
	ctx := GetContext_983_result                 // Contextʳ
	ctx.ActiveTextureUnit = ϟa.Unit
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_981_major, minRequiredVersion_981_minor, context, GetContext_983_result, ctx
	return nil
}
func (ϟa *GlBindImageTexture) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_984_major := uint32(3) // u32
	minRequiredVersion_984_minor := uint32(1) // u32
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
	_, _ = minRequiredVersion_984_major, minRequiredVersion_984_minor
	return nil
}
func (ϟa *GlBindSampler) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_987_major := uint32(3) // u32
	minRequiredVersion_987_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_987_major, minRequiredVersion_987_minor
	return nil
}
func (ϟa *GlBindTexture) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_988_major := uint32(2) // u32
	minRequiredVersion_988_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_989_major := uint32(3) // u32
		minRequiredVersion_989_minor := uint32(0) // u32
		_, _ = minRequiredVersion_989_major, minRequiredVersion_989_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_990_major := uint32(3) // u32
		minRequiredVersion_990_minor := uint32(1) // u32
		_, _ = minRequiredVersion_990_major, minRequiredVersion_990_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_BUFFER, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_991_major := uint32(3) // u32
		minRequiredVersion_991_minor := uint32(2) // u32
		_, _ = minRequiredVersion_991_major, minRequiredVersion_991_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_993_result := context             // Contextʳ
	ctx := GetContext_993_result                 // Contextʳ
	if !(ctx.Instances.Textures.Contains(ϟa.Texture)) {
		ctx.Instances.Textures[ϟa.Texture] = (&Texture{ID: ϟa.Texture, Texture2D: GLintːImageᵐ{}, Cubemap: GLintːCubemapLevelᵐ{}, MagFilter: GLenum_GL_LINEAR, MinFilter: GLenum_GL_NEAREST_MIPMAP_LINEAR, WrapS: GLenum_GL_REPEAT, WrapT: GLenum_GL_REPEAT, SwizzleR: GLenum_GL_RED, SwizzleG: GLenum_GL_GREEN, SwizzleB: GLenum_GL_BLUE, SwizzleA: GLenum_GL_ALPHA, MaxAnisotropy: float32(1)}).OnCreate(ϟs)
	}
	tu := ctx.TextureUnits.Get(ctx.ActiveTextureUnit) // TextureUnitʳ
	tu.Bindings[ϟa.Target] = ϟa.Texture
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_988_major, minRequiredVersion_988_minor, context, GetContext_993_result, ctx, tu
	return nil
}
func (ϟa *GlCompressedTexImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_994_major := uint32(2) // u32
	minRequiredVersion_994_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Format {
	case GLenum_GL_COMPRESSED_R11_EAC, GLenum_GL_COMPRESSED_RG11_EAC, GLenum_GL_COMPRESSED_RGB8_ETC2, GLenum_GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_COMPRESSED_RGBA8_ETC2_EAC, GLenum_GL_COMPRESSED_SIGNED_R11_EAC, GLenum_GL_COMPRESSED_SIGNED_RG11_EAC, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC, GLenum_GL_COMPRESSED_SRGB8_ETC2, GLenum_GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2:
		minRequiredVersion_996_major := uint32(3) // u32
		minRequiredVersion_996_minor := uint32(0) // u32
		_, _ = minRequiredVersion_996_major, minRequiredVersion_996_minor
	case GLenum_GL_ATC_RGB_AMD, GLenum_GL_ATC_RGBA_EXPLICIT_ALPHA_AMD, GLenum_GL_ATC_RGBA_INTERPOLATED_ALPHA_AMD:
		requiresExtension_997_ext := ExtensionId_GL_AMD_compressed_ATC_texture // ExtensionId
		_ = requiresExtension_997_ext
	case GLenum_GL_COMPRESSED_RGBA_ASTC_10x10, GLenum_GL_COMPRESSED_RGBA_ASTC_10x5, GLenum_GL_COMPRESSED_RGBA_ASTC_10x6, GLenum_GL_COMPRESSED_RGBA_ASTC_10x8, GLenum_GL_COMPRESSED_RGBA_ASTC_12x10, GLenum_GL_COMPRESSED_RGBA_ASTC_12x12, GLenum_GL_COMPRESSED_RGBA_ASTC_4x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x5, GLenum_GL_COMPRESSED_RGBA_ASTC_8x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x12, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_4x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x8:
		minRequiredVersion_998_major := uint32(3) // u32
		minRequiredVersion_998_minor := uint32(2) // u32
		_, _ = minRequiredVersion_998_major, minRequiredVersion_998_minor
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)      // Contextʳ
	GetContext_1000_result := context                 // Contextʳ
	ctx := GetContext_1000_result                     // Contextʳ
	tu := ctx.TextureUnits.Get(ctx.ActiveTextureUnit) // TextureUnitʳ
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D:
		id := tu.Bindings.Get(GLenum_GL_TEXTURE_2D)                                                   // TextureId
		t := ctx.Instances.Textures.Get(id)                                                           // Textureʳ
		l := Image{Width: ϟa.Width, Height: ϟa.Height, Size: uint32(ϟa.ImageSize), Format: ϟa.Format} // Image
		if ((ctx.BoundBuffers.Get(GLenum_GL_PIXEL_UNPACK_BUFFER)) == (BufferId(uint32(0)))) && ((ϟa.Data) != (TexturePointer(Voidᶜᵖ{}))) {
			l.Data = U8ᵖ(ϟa.Data).Slice(uint64(uint32(0)), uint64(l.Size), ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
		}
		t.OnAccess(ϟs).Texture2D[ϟa.Level] = l
		t.OnAccess(ϟs).Kind = TextureKind_TEXTURE2D
		t.OnAccess(ϟs).Format = ϟa.Format
		_, _, _ = id, t, l
	case GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z:
		id := tu.Bindings.Get(GLenum_GL_TEXTURE_CUBE_MAP)                                             // TextureId
		t := ctx.Instances.Textures.Get(id)                                                           // Textureʳ
		l := Image{Width: ϟa.Width, Height: ϟa.Height, Size: uint32(ϟa.ImageSize), Format: ϟa.Format} // Image
		if ((ctx.BoundBuffers.Get(GLenum_GL_PIXEL_UNPACK_BUFFER)) == (BufferId(uint32(0)))) && ((ϟa.Data) != (TexturePointer(Voidᶜᵖ{}))) {
			l.Data = U8ᵖ(ϟa.Data).Slice(uint64(uint32(0)), uint64(l.Size), ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
		}
		cube := t.OnAccess(ϟs).Cubemap.Get(ϟa.Level) // CubemapLevel
		cube.Faces[ϟa.Target] = l
		t.OnAccess(ϟs).Cubemap[ϟa.Level] = cube
		t.OnAccess(ϟs).Kind = TextureKind_CUBEMAP
		t.OnAccess(ϟs).Format = ϟa.Format
		_, _, _, _ = id, t, l, cube
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_994_major, minRequiredVersion_994_minor, context, GetContext_1000_result, ctx, tu
	return nil
}
func (ϟa *GlCompressedTexImage3D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1001_major := uint32(3) // u32
	minRequiredVersion_1001_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
	case GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1002_major := uint32(3) // u32
		minRequiredVersion_1002_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1002_major, minRequiredVersion_1002_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Internalformat {
	case GLenum_GL_COMPRESSED_R11_EAC, GLenum_GL_COMPRESSED_RG11_EAC, GLenum_GL_COMPRESSED_RGB8_ETC2, GLenum_GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_COMPRESSED_RGBA8_ETC2_EAC, GLenum_GL_COMPRESSED_SIGNED_R11_EAC, GLenum_GL_COMPRESSED_SIGNED_RG11_EAC, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC, GLenum_GL_COMPRESSED_SRGB8_ETC2, GLenum_GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2:
	case GLenum_GL_COMPRESSED_RGBA_ASTC_10x10, GLenum_GL_COMPRESSED_RGBA_ASTC_10x5, GLenum_GL_COMPRESSED_RGBA_ASTC_10x6, GLenum_GL_COMPRESSED_RGBA_ASTC_10x8, GLenum_GL_COMPRESSED_RGBA_ASTC_12x10, GLenum_GL_COMPRESSED_RGBA_ASTC_12x12, GLenum_GL_COMPRESSED_RGBA_ASTC_4x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x5, GLenum_GL_COMPRESSED_RGBA_ASTC_8x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x12, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_4x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x8:
		minRequiredVersion_1004_major := uint32(3) // u32
		minRequiredVersion_1004_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1004_major, minRequiredVersion_1004_minor
	default:
		v := ϟa.Internalformat
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1001_major, minRequiredVersion_1001_minor
	return nil
}
func (ϟa *GlCompressedTexSubImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1006_major := uint32(2) // u32
	minRequiredVersion_1006_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Format {
	case GLenum_GL_COMPRESSED_R11_EAC, GLenum_GL_COMPRESSED_RG11_EAC, GLenum_GL_COMPRESSED_RGB8_ETC2, GLenum_GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_COMPRESSED_RGBA8_ETC2_EAC, GLenum_GL_COMPRESSED_SIGNED_R11_EAC, GLenum_GL_COMPRESSED_SIGNED_RG11_EAC, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC, GLenum_GL_COMPRESSED_SRGB8_ETC2, GLenum_GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2:
		minRequiredVersion_1008_major := uint32(3) // u32
		minRequiredVersion_1008_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1008_major, minRequiredVersion_1008_minor
	case GLenum_GL_COMPRESSED_RGBA_ASTC_10x10, GLenum_GL_COMPRESSED_RGBA_ASTC_10x5, GLenum_GL_COMPRESSED_RGBA_ASTC_10x6, GLenum_GL_COMPRESSED_RGBA_ASTC_10x8, GLenum_GL_COMPRESSED_RGBA_ASTC_12x10, GLenum_GL_COMPRESSED_RGBA_ASTC_12x12, GLenum_GL_COMPRESSED_RGBA_ASTC_4x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x5, GLenum_GL_COMPRESSED_RGBA_ASTC_8x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x12, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_4x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x8:
		minRequiredVersion_1009_major := uint32(3) // u32
		minRequiredVersion_1009_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1009_major, minRequiredVersion_1009_minor
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1006_major, minRequiredVersion_1006_minor
	return nil
}
func (ϟa *GlCompressedTexSubImage3D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1011_major := uint32(3) // u32
	minRequiredVersion_1011_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
	case GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1012_major := uint32(3) // u32
		minRequiredVersion_1012_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1012_major, minRequiredVersion_1012_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Format {
	case GLenum_GL_COMPRESSED_R11_EAC, GLenum_GL_COMPRESSED_RG11_EAC, GLenum_GL_COMPRESSED_RGB8_ETC2, GLenum_GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_COMPRESSED_RGBA8_ETC2_EAC, GLenum_GL_COMPRESSED_SIGNED_R11_EAC, GLenum_GL_COMPRESSED_SIGNED_RG11_EAC, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC, GLenum_GL_COMPRESSED_SRGB8_ETC2, GLenum_GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2:
	case GLenum_GL_COMPRESSED_RGBA_ASTC_10x10, GLenum_GL_COMPRESSED_RGBA_ASTC_10x5, GLenum_GL_COMPRESSED_RGBA_ASTC_10x6, GLenum_GL_COMPRESSED_RGBA_ASTC_10x8, GLenum_GL_COMPRESSED_RGBA_ASTC_12x10, GLenum_GL_COMPRESSED_RGBA_ASTC_12x12, GLenum_GL_COMPRESSED_RGBA_ASTC_4x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x5, GLenum_GL_COMPRESSED_RGBA_ASTC_8x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x12, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_4x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x8:
		minRequiredVersion_1014_major := uint32(3) // u32
		minRequiredVersion_1014_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1014_major, minRequiredVersion_1014_minor
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1011_major, minRequiredVersion_1011_minor
	return nil
}
func (ϟa *GlCopyImageSubData) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1016_major := uint32(3) // u32
	minRequiredVersion_1016_minor := uint32(2) // u32
	switch ϟa.SrcTarget {
	case GLenum_GL_RENDERBUFFER, GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_2D_MULTISAMPLE, GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_3D, GLenum_GL_TEXTURE_CUBE_MAP, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
	default:
		v := ϟa.SrcTarget
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.DstTarget {
	case GLenum_GL_RENDERBUFFER, GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_2D_MULTISAMPLE, GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_3D, GLenum_GL_TEXTURE_CUBE_MAP, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
	default:
		v := ϟa.DstTarget
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1016_major, minRequiredVersion_1016_minor
	return nil
}
func (ϟa *GlCopyTexImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1019_major := uint32(2) // u32
	minRequiredVersion_1019_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Format {
	case GLenum_GL_ALPHA, GLenum_GL_LUMINANCE, GLenum_GL_LUMINANCE_ALPHA, GLenum_GL_RGB, GLenum_GL_RGBA:
	case GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGB8, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA4, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_SRGB8, GLenum_GL_SRGB8_ALPHA8:
		minRequiredVersion_1021_major := uint32(3) // u32
		minRequiredVersion_1021_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1021_major, minRequiredVersion_1021_minor
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1019_major, minRequiredVersion_1019_minor
	return nil
}
func (ϟa *GlCopyTexSubImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1023_major := uint32(2) // u32
	minRequiredVersion_1023_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1023_major, minRequiredVersion_1023_minor
	return nil
}
func (ϟa *GlCopyTexSubImage3D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1025_major := uint32(3) // u32
	minRequiredVersion_1025_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
	case GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1026_major := uint32(3) // u32
		minRequiredVersion_1026_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1026_major, minRequiredVersion_1026_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1025_major, minRequiredVersion_1025_minor
	return nil
}
func (ϟa *GlDeleteSamplers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1028_major := uint32(3) // u32
	minRequiredVersion_1028_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1028_major, minRequiredVersion_1028_minor
	return nil
}
func (ϟa *GlDeleteTextures) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1029_major := uint32(2)                              // u32
	minRequiredVersion_1029_minor := uint32(0)                              // u32
	t := ϟa.Textures.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // TextureIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                            // Contextʳ
	GetContext_1030_result := context                                       // Contextʳ
	ctx := GetContext_1030_result                                           // Contextʳ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.Textures, t.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_1029_major, minRequiredVersion_1029_minor, t, context, GetContext_1030_result, ctx
	return nil
}
func (ϟa *GlGenSamplers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1031_major := uint32(3) // u32
	minRequiredVersion_1031_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1031_major, minRequiredVersion_1031_minor
	return nil
}
func (ϟa *GlGenTextures) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1032_major := uint32(2)                              // u32
	minRequiredVersion_1032_minor := uint32(0)                              // u32
	t := ϟa.Textures.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // TextureIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                            // Contextʳ
	GetContext_1033_result := context                                       // Contextʳ
	ctx := GetContext_1033_result                                           // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := TextureId(ϟa.Textures.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // TextureId
		ctx.Instances.Textures[id] = (&Texture{ID: id, Texture2D: GLintːImageᵐ{}, Cubemap: GLintːCubemapLevelᵐ{}, MagFilter: GLenum_GL_LINEAR, MinFilter: GLenum_GL_NEAREST_MIPMAP_LINEAR, WrapS: GLenum_GL_REPEAT, WrapT: GLenum_GL_REPEAT, SwizzleR: GLenum_GL_RED, SwizzleG: GLenum_GL_GREEN, SwizzleB: GLenum_GL_BLUE, SwizzleA: GLenum_GL_ALPHA, MaxAnisotropy: float32(1)}).OnCreate(ϟs)
		t.Index(uint64(i), ϟs).Write(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = minRequiredVersion_1032_major, minRequiredVersion_1032_minor, t, context, GetContext_1033_result, ctx
	return nil
}
func (ϟa *GlGenerateMipmap) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1034_major := uint32(2) // u32
	minRequiredVersion_1034_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_1035_major := uint32(3) // u32
		minRequiredVersion_1035_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1035_major, minRequiredVersion_1035_minor
	case GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1036_major := uint32(3) // u32
		minRequiredVersion_1036_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1036_major, minRequiredVersion_1036_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1034_major, minRequiredVersion_1034_minor
	return nil
}
func (ϟa *GlGetSamplerParameterIiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1038_major := uint32(3) // u32
	minRequiredVersion_1038_minor := uint32(2) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_BORDER_COLOR, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1038_major, minRequiredVersion_1038_minor
	return nil
}
func (ϟa *GlGetSamplerParameterIuiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1040_major := uint32(3) // u32
	minRequiredVersion_1040_minor := uint32(2) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_BORDER_COLOR, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1040_major, minRequiredVersion_1040_minor
	return nil
}
func (ϟa *GlGetSamplerParameterfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1042_major := uint32(3) // u32
	minRequiredVersion_1042_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_1043_major := uint32(3) // u32
		minRequiredVersion_1043_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1043_major, minRequiredVersion_1043_minor
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1042_major, minRequiredVersion_1042_minor
	return nil
}
func (ϟa *GlGetSamplerParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1045_major := uint32(3) // u32
	minRequiredVersion_1045_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_1046_major := uint32(3) // u32
		minRequiredVersion_1046_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1046_major, minRequiredVersion_1046_minor
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1045_major, minRequiredVersion_1045_minor
	return nil
}
func (ϟa *GlGetTexLevelParameterfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1048_major := uint32(3) // u32
	minRequiredVersion_1048_minor := uint32(1) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_2D_MULTISAMPLE, GLenum_GL_TEXTURE_3D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_BUFFER, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1049_major := uint32(3) // u32
		minRequiredVersion_1049_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1049_major, minRequiredVersion_1049_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_ALPHA_SIZE, GLenum_GL_TEXTURE_ALPHA_TYPE, GLenum_GL_TEXTURE_BLUE_SIZE, GLenum_GL_TEXTURE_BLUE_TYPE, GLenum_GL_TEXTURE_COMPRESSED, GLenum_GL_TEXTURE_DEPTH, GLenum_GL_TEXTURE_DEPTH_SIZE, GLenum_GL_TEXTURE_DEPTH_TYPE, GLenum_GL_TEXTURE_FIXED_SAMPLE_LOCATIONS, GLenum_GL_TEXTURE_GREEN_SIZE, GLenum_GL_TEXTURE_GREEN_TYPE, GLenum_GL_TEXTURE_HEIGHT, GLenum_GL_TEXTURE_INTERNAL_FORMAT, GLenum_GL_TEXTURE_RED_SIZE, GLenum_GL_TEXTURE_RED_TYPE, GLenum_GL_TEXTURE_SAMPLES, GLenum_GL_TEXTURE_SHARED_SIZE, GLenum_GL_TEXTURE_STENCIL_SIZE, GLenum_GL_TEXTURE_WIDTH:
	case GLenum_GL_TEXTURE_BUFFER_DATA_STORE_BINDING, GLenum_GL_TEXTURE_BUFFER_OFFSET, GLenum_GL_TEXTURE_BUFFER_SIZE:
		minRequiredVersion_1051_major := uint32(3) // u32
		minRequiredVersion_1051_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1051_major, minRequiredVersion_1051_minor
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1048_major, minRequiredVersion_1048_minor
	return nil
}
func (ϟa *GlGetTexLevelParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1053_major := uint32(3) // u32
	minRequiredVersion_1053_minor := uint32(1) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_2D_MULTISAMPLE, GLenum_GL_TEXTURE_3D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_BUFFER, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1054_major := uint32(3) // u32
		minRequiredVersion_1054_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1054_major, minRequiredVersion_1054_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_ALPHA_SIZE, GLenum_GL_TEXTURE_ALPHA_TYPE, GLenum_GL_TEXTURE_BLUE_SIZE, GLenum_GL_TEXTURE_BLUE_TYPE, GLenum_GL_TEXTURE_COMPRESSED, GLenum_GL_TEXTURE_DEPTH, GLenum_GL_TEXTURE_DEPTH_SIZE, GLenum_GL_TEXTURE_DEPTH_TYPE, GLenum_GL_TEXTURE_FIXED_SAMPLE_LOCATIONS, GLenum_GL_TEXTURE_GREEN_SIZE, GLenum_GL_TEXTURE_GREEN_TYPE, GLenum_GL_TEXTURE_HEIGHT, GLenum_GL_TEXTURE_INTERNAL_FORMAT, GLenum_GL_TEXTURE_RED_SIZE, GLenum_GL_TEXTURE_RED_TYPE, GLenum_GL_TEXTURE_SAMPLES, GLenum_GL_TEXTURE_SHARED_SIZE, GLenum_GL_TEXTURE_STENCIL_SIZE, GLenum_GL_TEXTURE_WIDTH:
	case GLenum_GL_TEXTURE_BUFFER_DATA_STORE_BINDING, GLenum_GL_TEXTURE_BUFFER_OFFSET, GLenum_GL_TEXTURE_BUFFER_SIZE:
		minRequiredVersion_1056_major := uint32(3) // u32
		minRequiredVersion_1056_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1056_major, minRequiredVersion_1056_minor
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1053_major, minRequiredVersion_1053_minor
	return nil
}
func (ϟa *GlGetTexParameterIiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1058_major := uint32(3) // u32
	minRequiredVersion_1058_minor := uint32(2) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_2D_MULTISAMPLE, GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_3D, GLenum_GL_TEXTURE_CUBE_MAP, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Pname {
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE, GLenum_GL_IMAGE_FORMAT_COMPATIBILITY_TYPE, GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_BORDER_COLOR, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_IMMUTABLE_FORMAT, GLenum_GL_TEXTURE_IMMUTABLE_LEVELS, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1058_major, minRequiredVersion_1058_minor
	return nil
}
func (ϟa *GlGetTexParameterIuiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1061_major := uint32(3) // u32
	minRequiredVersion_1061_minor := uint32(2) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_2D_MULTISAMPLE, GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_3D, GLenum_GL_TEXTURE_CUBE_MAP, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Pname {
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE, GLenum_GL_IMAGE_FORMAT_COMPATIBILITY_TYPE, GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_BORDER_COLOR, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_IMMUTABLE_FORMAT, GLenum_GL_TEXTURE_IMMUTABLE_LEVELS, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1061_major, minRequiredVersion_1061_minor
	return nil
}
func (ϟa *GlGetTexParameterfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1064_major := uint32(2) // u32
	minRequiredVersion_1064_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_1065_major := uint32(3) // u32
		minRequiredVersion_1065_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1065_major, minRequiredVersion_1065_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_1066_major := uint32(3) // u32
		minRequiredVersion_1066_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1066_major, minRequiredVersion_1066_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1067_major := uint32(3) // u32
		minRequiredVersion_1067_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1067_major, minRequiredVersion_1067_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Parameter {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_IMMUTABLE_FORMAT, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_1069_major := uint32(3) // u32
		minRequiredVersion_1069_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1069_major, minRequiredVersion_1069_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE, GLenum_GL_IMAGE_FORMAT_COMPATIBILITY_TYPE, GLenum_GL_TEXTURE_IMMUTABLE_LEVELS:
		minRequiredVersion_1070_major := uint32(3) // u32
		minRequiredVersion_1070_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1070_major, minRequiredVersion_1070_minor
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_1071_major := uint32(3) // u32
		minRequiredVersion_1071_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1071_major, minRequiredVersion_1071_minor
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)      // Contextʳ
	GetContext_1073_result := context                 // Contextʳ
	ctx := GetContext_1073_result                     // Contextʳ
	tu := ctx.TextureUnits.Get(ctx.ActiveTextureUnit) // TextureUnitʳ
	id := tu.Bindings.Get(ϟa.Target)                  // TextureId
	t := ctx.Instances.Textures.Get(id)               // Textureʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Values.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(func() (result GLfloat) {
		switch ϟa.Parameter {
		case GLenum_GL_TEXTURE_MAG_FILTER:
			return GLfloat(t.OnAccess(ϟs).MagFilter)
		case GLenum_GL_TEXTURE_MIN_FILTER:
			return GLfloat(t.OnAccess(ϟs).MinFilter)
		case GLenum_GL_TEXTURE_WRAP_S:
			return GLfloat(t.OnAccess(ϟs).WrapS)
		case GLenum_GL_TEXTURE_WRAP_T:
			return GLfloat(t.OnAccess(ϟs).WrapT)
		case GLenum_GL_TEXTURE_MAX_ANISOTROPY_EXT:
			return GLfloat(t.OnAccess(ϟs).MaxAnisotropy)
		case GLenum_GL_TEXTURE_SWIZZLE_R:
			return GLfloat(t.OnAccess(ϟs).SwizzleR)
		case GLenum_GL_TEXTURE_SWIZZLE_G:
			return GLfloat(t.OnAccess(ϟs).SwizzleG)
		case GLenum_GL_TEXTURE_SWIZZLE_B:
			return GLfloat(t.OnAccess(ϟs).SwizzleB)
		case GLenum_GL_TEXTURE_SWIZZLE_A:
			return GLfloat(t.OnAccess(ϟs).SwizzleA)
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.Parameter, ϟa))
			return result
		}
	}(), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _, _, _ = minRequiredVersion_1064_major, minRequiredVersion_1064_minor, context, GetContext_1073_result, ctx, tu, id, t
	return nil
}
func (ϟa *GlGetTexParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1074_major := uint32(2) // u32
	minRequiredVersion_1074_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_1075_major := uint32(3) // u32
		minRequiredVersion_1075_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1075_major, minRequiredVersion_1075_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_1076_major := uint32(3) // u32
		minRequiredVersion_1076_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1076_major, minRequiredVersion_1076_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1077_major := uint32(3) // u32
		minRequiredVersion_1077_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1077_major, minRequiredVersion_1077_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Parameter {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_IMMUTABLE_FORMAT, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_1079_major := uint32(3) // u32
		minRequiredVersion_1079_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1079_major, minRequiredVersion_1079_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE, GLenum_GL_IMAGE_FORMAT_COMPATIBILITY_TYPE, GLenum_GL_TEXTURE_IMMUTABLE_LEVELS:
		minRequiredVersion_1080_major := uint32(3) // u32
		minRequiredVersion_1080_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1080_major, minRequiredVersion_1080_minor
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_1081_major := uint32(3) // u32
		minRequiredVersion_1081_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1081_major, minRequiredVersion_1081_minor
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)      // Contextʳ
	GetContext_1083_result := context                 // Contextʳ
	ctx := GetContext_1083_result                     // Contextʳ
	tu := ctx.TextureUnits.Get(ctx.ActiveTextureUnit) // TextureUnitʳ
	id := tu.Bindings.Get(ϟa.Target)                  // TextureId
	t := ctx.Instances.Textures.Get(id)               // Textureʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Values.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(func() (result GLint) {
		switch ϟa.Parameter {
		case GLenum_GL_TEXTURE_MAG_FILTER:
			return GLint(t.OnAccess(ϟs).MagFilter)
		case GLenum_GL_TEXTURE_MIN_FILTER:
			return GLint(t.OnAccess(ϟs).MinFilter)
		case GLenum_GL_TEXTURE_WRAP_S:
			return GLint(t.OnAccess(ϟs).WrapS)
		case GLenum_GL_TEXTURE_WRAP_T:
			return GLint(t.OnAccess(ϟs).WrapT)
		case GLenum_GL_TEXTURE_MAX_ANISOTROPY_EXT:
			return GLint(t.OnAccess(ϟs).MaxAnisotropy)
		case GLenum_GL_TEXTURE_SWIZZLE_R:
			return GLint(t.OnAccess(ϟs).SwizzleR)
		case GLenum_GL_TEXTURE_SWIZZLE_G:
			return GLint(t.OnAccess(ϟs).SwizzleG)
		case GLenum_GL_TEXTURE_SWIZZLE_B:
			return GLint(t.OnAccess(ϟs).SwizzleB)
		case GLenum_GL_TEXTURE_SWIZZLE_A:
			return GLint(t.OnAccess(ϟs).SwizzleA)
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.Parameter, ϟa))
			return result
		}
	}(), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _, _, _ = minRequiredVersion_1074_major, minRequiredVersion_1074_minor, context, GetContext_1083_result, ctx, tu, id, t
	return nil
}
func (ϟa *GlIsSampler) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1084_major := uint32(3) // u32
	minRequiredVersion_1084_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_1084_major, minRequiredVersion_1084_minor
	return nil
}
func (ϟa *GlIsTexture) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1085_major := uint32(2)   // u32
	minRequiredVersion_1085_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_1086_result := context            // Contextʳ
	ctx := GetContext_1086_result                // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = func() GLboolean {
		if ctx.Instances.Textures.Contains(ϟa.Texture) {
			return 1
		} else {
			return 0
		}
	}()
	_, _, _, _, _ = minRequiredVersion_1085_major, minRequiredVersion_1085_minor, context, GetContext_1086_result, ctx
	return nil
}
func (ϟa *GlPixelStorei) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1087_major := uint32(2) // u32
	minRequiredVersion_1087_minor := uint32(0) // u32
	switch ϟa.Parameter {
	case GLenum_GL_PACK_ALIGNMENT, GLenum_GL_UNPACK_ALIGNMENT:
	case GLenum_GL_PACK_IMAGE_HEIGHT, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_IMAGES, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS:
		minRequiredVersion_1088_major := uint32(3) // u32
		minRequiredVersion_1088_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1088_major, minRequiredVersion_1088_minor
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_1090_result := context            // Contextʳ
	ctx := GetContext_1090_result                // Contextʳ
	ctx.PixelStorage[ϟa.Parameter] = ϟa.Value
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1087_major, minRequiredVersion_1087_minor, context, GetContext_1090_result, ctx
	return nil
}
func (ϟa *GlSamplerParameterIiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1091_major := uint32(3) // u32
	minRequiredVersion_1091_minor := uint32(2) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_BORDER_COLOR, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1091_major, minRequiredVersion_1091_minor
	return nil
}
func (ϟa *GlSamplerParameterIuiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1093_major := uint32(3) // u32
	minRequiredVersion_1093_minor := uint32(2) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_BORDER_COLOR, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1093_major, minRequiredVersion_1093_minor
	return nil
}
func (ϟa *GlSamplerParameterf) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1095_major := uint32(3) // u32
	minRequiredVersion_1095_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1095_major, minRequiredVersion_1095_minor
	return nil
}
func (ϟa *GlSamplerParameterfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1097_major := uint32(3) // u32
	minRequiredVersion_1097_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_1098_major := uint32(3) // u32
		minRequiredVersion_1098_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1098_major, minRequiredVersion_1098_minor
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1097_major, minRequiredVersion_1097_minor
	return nil
}
func (ϟa *GlSamplerParameteri) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1100_major := uint32(3) // u32
	minRequiredVersion_1100_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1100_major, minRequiredVersion_1100_minor
	return nil
}
func (ϟa *GlSamplerParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1102_major := uint32(3) // u32
	minRequiredVersion_1102_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_1103_major := uint32(3) // u32
		minRequiredVersion_1103_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1103_major, minRequiredVersion_1103_minor
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1102_major, minRequiredVersion_1102_minor
	return nil
}
func (ϟa *GlTexBuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1105_major := uint32(3) // u32
	minRequiredVersion_1105_minor := uint32(2) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_BUFFER:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Internalformat {
	case GLenum_GL_R16, GLenum_GL_R16F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32F, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_RG16, GLenum_GL_RG16F, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32F, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RGB32F, GLenum_GL_RGB32I, GLenum_GL_RGB32UI, GLenum_GL_RGBA16, GLenum_GL_RGBA16F, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32F, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI:
	default:
		v := ϟa.Internalformat
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1105_major, minRequiredVersion_1105_minor
	return nil
}
func (ϟa *GlTexBufferRange) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1108_major := uint32(3) // u32
	minRequiredVersion_1108_minor := uint32(2) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_BUFFER:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Internalformat {
	case GLenum_GL_R16, GLenum_GL_R16F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32F, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_RG16, GLenum_GL_RG16F, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32F, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RGB32F, GLenum_GL_RGB32I, GLenum_GL_RGB32UI, GLenum_GL_RGBA16, GLenum_GL_RGBA16F, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32F, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI:
	default:
		v := ϟa.Internalformat
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1108_major, minRequiredVersion_1108_minor
	return nil
}
func (ϟa *GlTexImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1111_major := uint32(2) // u32
	minRequiredVersion_1111_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Format {
	case GLenum_GL_ALPHA, GLenum_GL_LUMINANCE, GLenum_GL_LUMINANCE_ALPHA, GLenum_GL_RGB, GLenum_GL_RGBA:
	case GLenum_GL_DEPTH_COMPONENT, GLenum_GL_DEPTH_STENCIL, GLenum_GL_RED, GLenum_GL_RED_INTEGER, GLenum_GL_RG, GLenum_GL_RGBA_INTEGER, GLenum_GL_RGB_INTEGER, GLenum_GL_RG_INTEGER:
		minRequiredVersion_1113_major := uint32(3) // u32
		minRequiredVersion_1113_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1113_major, minRequiredVersion_1113_minor
	case GLenum_GL_STENCIL_INDEX:
		minRequiredVersion_1114_major := uint32(3) // u32
		minRequiredVersion_1114_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1114_major, minRequiredVersion_1114_minor
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Type {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_SHORT_4_4_4_4, GLenum_GL_UNSIGNED_SHORT_5_5_5_1, GLenum_GL_UNSIGNED_SHORT_5_6_5:
	case GLenum_GL_HALF_FLOAT_OES:
		requiresExtension_1116_ext := ExtensionId_GL_OES_texture_half_float // ExtensionId
		_ = requiresExtension_1116_ext
	case GLenum_GL_BYTE, GLenum_GL_FLOAT, GLenum_GL_FLOAT_32_UNSIGNED_INT_24_8_REV, GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_10F_11F_11F_REV, GLenum_GL_UNSIGNED_INT_24_8, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_INT_5_9_9_9_REV, GLenum_GL_UNSIGNED_SHORT:
		minRequiredVersion_1117_major := uint32(3) // u32
		minRequiredVersion_1117_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1117_major, minRequiredVersion_1117_minor
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)      // Contextʳ
	GetContext_1119_result := context                 // Contextʳ
	ctx := GetContext_1119_result                     // Contextʳ
	tu := ctx.TextureUnits.Get(ctx.ActiveTextureUnit) // TextureUnitʳ
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D:
		id := tu.Bindings.Get(GLenum_GL_TEXTURE_2D)                                                                                                                             // TextureId
		t := ctx.Instances.Textures.Get(id)                                                                                                                                     // Textureʳ
		l := Image{Width: ϟa.Width, Height: ϟa.Height, Size: externs{ϟa, ϟs, ϟd, ϟl, ϟb}.imageSize(uint32(ϟa.Width), uint32(ϟa.Height), ϟa.Format, ϟa.Type), Format: ϟa.Format} // Image
		if (ϟa.Data) != (TexturePointer(Voidᶜᵖ{})) {
			if (ctx.BoundBuffers.Get(GLenum_GL_PIXEL_UNPACK_BUFFER)) == (BufferId(uint32(0))) {
				l.Data = U8ᵖ(ϟa.Data).Slice(uint64(uint32(0)), uint64(l.Size), ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
			}
		} else {
			l.Data = MakeU8ˢ(uint64(l.Size), ϟs)
		}
		t.OnAccess(ϟs).Texture2D[ϟa.Level] = l
		t.OnAccess(ϟs).Kind = TextureKind_TEXTURE2D
		t.OnAccess(ϟs).Format = ϟa.Format
		_, _, _ = id, t, l
	case GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z:
		id := tu.Bindings.Get(GLenum_GL_TEXTURE_CUBE_MAP)                                                                                                                       // TextureId
		t := ctx.Instances.Textures.Get(id)                                                                                                                                     // Textureʳ
		l := Image{Width: ϟa.Width, Height: ϟa.Height, Size: externs{ϟa, ϟs, ϟd, ϟl, ϟb}.imageSize(uint32(ϟa.Width), uint32(ϟa.Height), ϟa.Format, ϟa.Type), Format: ϟa.Format} // Image
		if (ϟa.Data) != (TexturePointer(Voidᶜᵖ{})) {
			if (ctx.BoundBuffers.Get(GLenum_GL_PIXEL_UNPACK_BUFFER)) == (BufferId(uint32(0))) {
				l.Data = U8ᵖ(ϟa.Data).Slice(uint64(uint32(0)), uint64(l.Size), ϟs).Clone(ϟa, ϟs, ϟd, ϟl, ϟb)
			}
		} else {
			l.Data = MakeU8ˢ(uint64(l.Size), ϟs)
		}
		cube := t.OnAccess(ϟs).Cubemap.Get(ϟa.Level) // CubemapLevel
		cube.Faces[ϟa.Target] = l
		t.OnAccess(ϟs).Cubemap[ϟa.Level] = cube
		t.OnAccess(ϟs).Kind = TextureKind_CUBEMAP
		t.OnAccess(ϟs).Format = ϟa.Format
		_, _, _, _ = id, t, l, cube
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_1111_major, minRequiredVersion_1111_minor, context, GetContext_1119_result, ctx, tu
	return nil
}
func (ϟa *GlTexImage3D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1120_major := uint32(3) // u32
	minRequiredVersion_1120_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
	case GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1121_major := uint32(3) // u32
		minRequiredVersion_1121_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1121_major, minRequiredVersion_1121_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Format {
	case GLenum_GL_ALPHA, GLenum_GL_DEPTH_COMPONENT, GLenum_GL_DEPTH_STENCIL, GLenum_GL_LUMINANCE, GLenum_GL_LUMINANCE_ALPHA, GLenum_GL_RED, GLenum_GL_RED_INTEGER, GLenum_GL_RG, GLenum_GL_RGB, GLenum_GL_RGBA, GLenum_GL_RGBA_INTEGER, GLenum_GL_RGB_INTEGER, GLenum_GL_RG_INTEGER:
	case GLenum_GL_STENCIL_INDEX:
		minRequiredVersion_1123_major := uint32(3) // u32
		minRequiredVersion_1123_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1123_major, minRequiredVersion_1123_minor
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Type {
	case GLenum_GL_HALF_FLOAT_OES:
		requiresExtension_1125_ext := ExtensionId_GL_OES_texture_half_float // ExtensionId
		_ = requiresExtension_1125_ext
	case GLenum_GL_BYTE, GLenum_GL_FLOAT, GLenum_GL_FLOAT_32_UNSIGNED_INT_24_8_REV, GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_10F_11F_11F_REV, GLenum_GL_UNSIGNED_INT_24_8, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_INT_5_9_9_9_REV, GLenum_GL_UNSIGNED_SHORT, GLenum_GL_UNSIGNED_SHORT_4_4_4_4, GLenum_GL_UNSIGNED_SHORT_5_5_5_1, GLenum_GL_UNSIGNED_SHORT_5_6_5:
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1120_major, minRequiredVersion_1120_minor
	return nil
}
func (ϟa *GlTexParameterIiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1127_major := uint32(3) // u32
	minRequiredVersion_1127_minor := uint32(2) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_2D_MULTISAMPLE, GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_3D, GLenum_GL_TEXTURE_CUBE_MAP, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Pname {
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE, GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_BORDER_COLOR, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1127_major, minRequiredVersion_1127_minor
	return nil
}
func (ϟa *GlTexParameterIuiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1130_major := uint32(3) // u32
	minRequiredVersion_1130_minor := uint32(2) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_2D_MULTISAMPLE, GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_3D, GLenum_GL_TEXTURE_CUBE_MAP, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Pname {
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE, GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_BORDER_COLOR, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1130_major, minRequiredVersion_1130_minor
	return nil
}
func (ϟa *GlTexParameterf) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1133_major := uint32(2) // u32
	minRequiredVersion_1133_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_1134_major := uint32(3) // u32
		minRequiredVersion_1134_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1134_major, minRequiredVersion_1134_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_1135_major := uint32(3) // u32
		minRequiredVersion_1135_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1135_major, minRequiredVersion_1135_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1136_major := uint32(3) // u32
		minRequiredVersion_1136_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1136_major, minRequiredVersion_1136_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Parameter {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_1138_major := uint32(3) // u32
		minRequiredVersion_1138_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1138_major, minRequiredVersion_1138_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE:
		minRequiredVersion_1139_major := uint32(3) // u32
		minRequiredVersion_1139_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1139_major, minRequiredVersion_1139_minor
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)      // Contextʳ
	GetContext_1141_result := context                 // Contextʳ
	ctx := GetContext_1141_result                     // Contextʳ
	tu := ctx.TextureUnits.Get(ctx.ActiveTextureUnit) // TextureUnitʳ
	id := tu.Bindings.Get(ϟa.Target)                  // TextureId
	t := ctx.Instances.Textures.Get(id)               // Textureʳ
	switch ϟa.Parameter {
	case GLenum_GL_TEXTURE_MAG_FILTER:
		t.OnAccess(ϟs).MagFilter = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_MIN_FILTER:
		t.OnAccess(ϟs).MinFilter = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_WRAP_S:
		t.OnAccess(ϟs).WrapS = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_WRAP_T:
		t.OnAccess(ϟs).WrapT = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_MAX_ANISOTROPY_EXT:
		t.OnAccess(ϟs).MaxAnisotropy = float32(ϟa.Value)
	case GLenum_GL_TEXTURE_SWIZZLE_R:
		t.OnAccess(ϟs).SwizzleR = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_SWIZZLE_G:
		t.OnAccess(ϟs).SwizzleG = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_SWIZZLE_B:
		t.OnAccess(ϟs).SwizzleB = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_SWIZZLE_A:
		t.OnAccess(ϟs).SwizzleA = GLenum(ϟa.Value)
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_1133_major, minRequiredVersion_1133_minor, context, GetContext_1141_result, ctx, tu, id, t
	return nil
}
func (ϟa *GlTexParameterfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1142_major := uint32(2) // u32
	minRequiredVersion_1142_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_1143_major := uint32(3) // u32
		minRequiredVersion_1143_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1143_major, minRequiredVersion_1143_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_1144_major := uint32(3) // u32
		minRequiredVersion_1144_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1144_major, minRequiredVersion_1144_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1145_major := uint32(3) // u32
		minRequiredVersion_1145_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1145_major, minRequiredVersion_1145_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_1147_major := uint32(3) // u32
		minRequiredVersion_1147_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1147_major, minRequiredVersion_1147_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE:
		minRequiredVersion_1148_major := uint32(3) // u32
		minRequiredVersion_1148_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1148_major, minRequiredVersion_1148_minor
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_1149_major := uint32(3) // u32
		minRequiredVersion_1149_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1149_major, minRequiredVersion_1149_minor
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1142_major, minRequiredVersion_1142_minor
	return nil
}
func (ϟa *GlTexParameteri) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1151_major := uint32(2) // u32
	minRequiredVersion_1151_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_1152_major := uint32(3) // u32
		minRequiredVersion_1152_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1152_major, minRequiredVersion_1152_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_1153_major := uint32(3) // u32
		minRequiredVersion_1153_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1153_major, minRequiredVersion_1153_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1154_major := uint32(3) // u32
		minRequiredVersion_1154_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1154_major, minRequiredVersion_1154_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Parameter {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_1156_major := uint32(3) // u32
		minRequiredVersion_1156_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1156_major, minRequiredVersion_1156_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE:
		minRequiredVersion_1157_major := uint32(3) // u32
		minRequiredVersion_1157_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1157_major, minRequiredVersion_1157_minor
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)      // Contextʳ
	GetContext_1159_result := context                 // Contextʳ
	ctx := GetContext_1159_result                     // Contextʳ
	tu := ctx.TextureUnits.Get(ctx.ActiveTextureUnit) // TextureUnitʳ
	id := tu.Bindings.Get(ϟa.Target)                  // TextureId
	t := ctx.Instances.Textures.Get(id)               // Textureʳ
	switch ϟa.Parameter {
	case GLenum_GL_TEXTURE_MAG_FILTER:
		t.OnAccess(ϟs).MagFilter = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_MIN_FILTER:
		t.OnAccess(ϟs).MinFilter = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_WRAP_S:
		t.OnAccess(ϟs).WrapS = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_WRAP_T:
		t.OnAccess(ϟs).WrapT = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_MAX_ANISOTROPY_EXT:
		t.OnAccess(ϟs).MaxAnisotropy = float32(ϟa.Value)
	case GLenum_GL_TEXTURE_SWIZZLE_R:
		t.OnAccess(ϟs).SwizzleR = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_SWIZZLE_G:
		t.OnAccess(ϟs).SwizzleG = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_SWIZZLE_B:
		t.OnAccess(ϟs).SwizzleB = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_SWIZZLE_A:
		t.OnAccess(ϟs).SwizzleA = GLenum(ϟa.Value)
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_1151_major, minRequiredVersion_1151_minor, context, GetContext_1159_result, ctx, tu, id, t
	return nil
}
func (ϟa *GlTexParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1160_major := uint32(2) // u32
	minRequiredVersion_1160_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_1161_major := uint32(3) // u32
		minRequiredVersion_1161_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1161_major, minRequiredVersion_1161_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_1162_major := uint32(3) // u32
		minRequiredVersion_1162_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1162_major, minRequiredVersion_1162_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY, GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1163_major := uint32(3) // u32
		minRequiredVersion_1163_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1163_major, minRequiredVersion_1163_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_1165_major := uint32(3) // u32
		minRequiredVersion_1165_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1165_major, minRequiredVersion_1165_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE:
		minRequiredVersion_1166_major := uint32(3) // u32
		minRequiredVersion_1166_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1166_major, minRequiredVersion_1166_minor
	case GLenum_GL_TEXTURE_BORDER_COLOR:
		minRequiredVersion_1167_major := uint32(3) // u32
		minRequiredVersion_1167_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1167_major, minRequiredVersion_1167_minor
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1160_major, minRequiredVersion_1160_minor
	return nil
}
func (ϟa *GlTexStorage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1169_major := uint32(3) // u32
	minRequiredVersion_1169_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Internalformat {
	case GLenum_GL_COMPRESSED_R11_EAC, GLenum_GL_COMPRESSED_RG11_EAC, GLenum_GL_COMPRESSED_RGB8_ETC2, GLenum_GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_COMPRESSED_RGBA8_ETC2_EAC, GLenum_GL_COMPRESSED_SIGNED_R11_EAC, GLenum_GL_COMPRESSED_SIGNED_RG11_EAC, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC, GLenum_GL_COMPRESSED_SRGB8_ETC2, GLenum_GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_DEPTH24_STENCIL8, GLenum_GL_DEPTH32F_STENCIL8, GLenum_GL_DEPTH_COMPONENT16, GLenum_GL_DEPTH_COMPONENT24, GLenum_GL_DEPTH_COMPONENT32F, GLenum_GL_R11F_G11F_B10F, GLenum_GL_R16F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32F, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_R8_SNORM, GLenum_GL_RG16F, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32F, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RG8_SNORM, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB16F, GLenum_GL_RGB16I, GLenum_GL_RGB16UI, GLenum_GL_RGB32F, GLenum_GL_RGB32I, GLenum_GL_RGB32UI, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGB8, GLenum_GL_RGB8I, GLenum_GL_RGB8UI, GLenum_GL_RGB8_SNORM, GLenum_GL_RGB9_E5, GLenum_GL_RGBA16F, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32F, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA4, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_RGBA8_SNORM, GLenum_GL_SRGB8, GLenum_GL_SRGB8_ALPHA8:
	case GLenum_GL_COMPRESSED_RGBA_ASTC_10x10, GLenum_GL_COMPRESSED_RGBA_ASTC_10x5, GLenum_GL_COMPRESSED_RGBA_ASTC_10x6, GLenum_GL_COMPRESSED_RGBA_ASTC_10x8, GLenum_GL_COMPRESSED_RGBA_ASTC_12x10, GLenum_GL_COMPRESSED_RGBA_ASTC_12x12, GLenum_GL_COMPRESSED_RGBA_ASTC_4x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x5, GLenum_GL_COMPRESSED_RGBA_ASTC_8x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x12, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_4x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x8, GLenum_GL_STENCIL_INDEX8:
		minRequiredVersion_1171_major := uint32(3) // u32
		minRequiredVersion_1171_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1171_major, minRequiredVersion_1171_minor
	default:
		v := ϟa.Internalformat
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1169_major, minRequiredVersion_1169_minor
	return nil
}
func (ϟa *GlTexStorage2DMultisample) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1173_major := uint32(3) // u32
	minRequiredVersion_1173_minor := uint32(1) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Internalformat {
	case GLenum_GL_DEPTH24_STENCIL8, GLenum_GL_DEPTH32F_STENCIL8, GLenum_GL_DEPTH_COMPONENT16, GLenum_GL_DEPTH_COMPONENT24, GLenum_GL_DEPTH_COMPONENT32F, GLenum_GL_R11F_G11F_B10F, GLenum_GL_R16F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32F, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_R8_SNORM, GLenum_GL_RG16F, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32F, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RG8_SNORM, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB16F, GLenum_GL_RGB16I, GLenum_GL_RGB16UI, GLenum_GL_RGB32F, GLenum_GL_RGB32I, GLenum_GL_RGB32UI, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGB8, GLenum_GL_RGB8I, GLenum_GL_RGB8UI, GLenum_GL_RGB8_SNORM, GLenum_GL_RGB9_E5, GLenum_GL_RGBA16F, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32F, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA4, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_RGBA8_SNORM, GLenum_GL_SRGB8, GLenum_GL_SRGB8_ALPHA8:
	case GLenum_GL_STENCIL_INDEX8:
		minRequiredVersion_1175_major := uint32(3) // u32
		minRequiredVersion_1175_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1175_major, minRequiredVersion_1175_minor
	default:
		v := ϟa.Internalformat
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1173_major, minRequiredVersion_1173_minor
	return nil
}
func (ϟa *GlTexStorage3D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1177_major := uint32(3) // u32
	minRequiredVersion_1177_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
	case GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1178_major := uint32(3) // u32
		minRequiredVersion_1178_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1178_major, minRequiredVersion_1178_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Internalformat {
	case GLenum_GL_COMPRESSED_R11_EAC, GLenum_GL_COMPRESSED_RG11_EAC, GLenum_GL_COMPRESSED_RGB8_ETC2, GLenum_GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_COMPRESSED_RGBA8_ETC2_EAC, GLenum_GL_COMPRESSED_SIGNED_R11_EAC, GLenum_GL_COMPRESSED_SIGNED_RG11_EAC, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC, GLenum_GL_COMPRESSED_SRGB8_ETC2, GLenum_GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_DEPTH24_STENCIL8, GLenum_GL_DEPTH32F_STENCIL8, GLenum_GL_DEPTH_COMPONENT16, GLenum_GL_DEPTH_COMPONENT24, GLenum_GL_DEPTH_COMPONENT32F, GLenum_GL_R11F_G11F_B10F, GLenum_GL_R16F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32F, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_R8_SNORM, GLenum_GL_RG16F, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32F, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RG8_SNORM, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB16F, GLenum_GL_RGB16I, GLenum_GL_RGB16UI, GLenum_GL_RGB32F, GLenum_GL_RGB32I, GLenum_GL_RGB32UI, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGB8, GLenum_GL_RGB8I, GLenum_GL_RGB8UI, GLenum_GL_RGB8_SNORM, GLenum_GL_RGB9_E5, GLenum_GL_RGBA16F, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32F, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA4, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_RGBA8_SNORM, GLenum_GL_SRGB8, GLenum_GL_SRGB8_ALPHA8:
	case GLenum_GL_COMPRESSED_RGBA_ASTC_10x10, GLenum_GL_COMPRESSED_RGBA_ASTC_10x5, GLenum_GL_COMPRESSED_RGBA_ASTC_10x6, GLenum_GL_COMPRESSED_RGBA_ASTC_10x8, GLenum_GL_COMPRESSED_RGBA_ASTC_12x10, GLenum_GL_COMPRESSED_RGBA_ASTC_12x12, GLenum_GL_COMPRESSED_RGBA_ASTC_4x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x4, GLenum_GL_COMPRESSED_RGBA_ASTC_5x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x5, GLenum_GL_COMPRESSED_RGBA_ASTC_6x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x5, GLenum_GL_COMPRESSED_RGBA_ASTC_8x6, GLenum_GL_COMPRESSED_RGBA_ASTC_8x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x8, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x10, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x12, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_4x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x4, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x5, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x6, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x8, GLenum_GL_STENCIL_INDEX8:
		minRequiredVersion_1180_major := uint32(3) // u32
		minRequiredVersion_1180_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1180_major, minRequiredVersion_1180_minor
	default:
		v := ϟa.Internalformat
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1177_major, minRequiredVersion_1177_minor
	return nil
}
func (ϟa *GlTexStorage3DMultisample) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1182_major := uint32(3) // u32
	minRequiredVersion_1182_minor := uint32(2) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE_ARRAY:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Internalformat {
	case GLenum_GL_DEPTH24_STENCIL8, GLenum_GL_DEPTH32F_STENCIL8, GLenum_GL_DEPTH_COMPONENT16, GLenum_GL_DEPTH_COMPONENT24, GLenum_GL_DEPTH_COMPONENT32F, GLenum_GL_R11F_G11F_B10F, GLenum_GL_R16F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32F, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_R8_SNORM, GLenum_GL_RG16F, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32F, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RG8_SNORM, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB16F, GLenum_GL_RGB16I, GLenum_GL_RGB16UI, GLenum_GL_RGB32F, GLenum_GL_RGB32I, GLenum_GL_RGB32UI, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGB8, GLenum_GL_RGB8I, GLenum_GL_RGB8UI, GLenum_GL_RGB8_SNORM, GLenum_GL_RGB9_E5, GLenum_GL_RGBA16F, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32F, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA4, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_RGBA8_SNORM, GLenum_GL_SRGB8, GLenum_GL_SRGB8_ALPHA8, GLenum_GL_STENCIL_INDEX8:
	default:
		v := ϟa.Internalformat
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1182_major, minRequiredVersion_1182_minor
	return nil
}
func (ϟa *GlTexSubImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1185_major := uint32(2) // u32
	minRequiredVersion_1185_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Format {
	case GLenum_GL_ALPHA, GLenum_GL_LUMINANCE, GLenum_GL_LUMINANCE_ALPHA, GLenum_GL_RGB, GLenum_GL_RGBA:
	case GLenum_GL_DEPTH_COMPONENT, GLenum_GL_DEPTH_STENCIL, GLenum_GL_RED, GLenum_GL_RED_INTEGER, GLenum_GL_RG, GLenum_GL_RGBA_INTEGER, GLenum_GL_RGB_INTEGER, GLenum_GL_RG_INTEGER:
		minRequiredVersion_1187_major := uint32(3) // u32
		minRequiredVersion_1187_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1187_major, minRequiredVersion_1187_minor
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Type {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_SHORT_4_4_4_4, GLenum_GL_UNSIGNED_SHORT_5_5_5_1, GLenum_GL_UNSIGNED_SHORT_5_6_5:
	case GLenum_GL_HALF_FLOAT_OES:
		requiresExtension_1189_ext := ExtensionId_GL_OES_texture_half_float // ExtensionId
		_ = requiresExtension_1189_ext
	case GLenum_GL_BYTE, GLenum_GL_FLOAT, GLenum_GL_FLOAT_32_UNSIGNED_INT_24_8_REV, GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_10F_11F_11F_REV, GLenum_GL_UNSIGNED_INT_24_8, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_INT_5_9_9_9_REV, GLenum_GL_UNSIGNED_SHORT:
		minRequiredVersion_1190_major := uint32(3) // u32
		minRequiredVersion_1190_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1190_major, minRequiredVersion_1190_minor
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)      // Contextʳ
	GetContext_1192_result := context                 // Contextʳ
	ctx := GetContext_1192_result                     // Contextʳ
	tu := ctx.TextureUnits.Get(ctx.ActiveTextureUnit) // TextureUnitʳ
	image := func() (result Image) {
		switch ϟa.Target {
		case GLenum_GL_TEXTURE_2D:
			return ctx.Instances.Textures.Get(tu.Bindings.Get(GLenum_GL_TEXTURE_2D)).OnAccess(ϟs).Texture2D.Get(ϟa.Level)
		case GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z:
			return ctx.Instances.Textures.Get(tu.Bindings.Get(GLenum_GL_TEXTURE_CUBE_MAP)).OnAccess(ϟs).Cubemap.Get(ϟa.Level).Faces.Get(ϟa.Target)
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.Target, ϟa))
			return result
		}
	}() // Image
	pbo := ctx.BoundBuffers.Get(GLenum_GL_PIXEL_UNPACK_BUFFER) // BufferId
	url := ctx.PixelStorage.Get(GLenum_GL_UNPACK_ROW_LENGTH)   // GLint
	src_width := func() (result uint32) {
		switch (url) == (GLint(int32(0))) {
		case true:
			return uint32(ϟa.Width)
		case false:
			return uint32(url)
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (url) == (GLint(int32(0))), ϟa))
			return result
		}
	}() // u32
	src_stride := externs{ϟa, ϟs, ϟd, ϟl, ϟb}.imageSize(src_width, uint32(1), ϟa.Format, ϟa.Type)                                                    // u32
	src_size := (src_stride) * (uint32(ϟa.Height))                                                                                                   // u32
	dst_stride := externs{ϟa, ϟs, ϟd, ϟl, ϟb}.imageSize(uint32(image.Width), uint32(1), ϟa.Format, ϟa.Type)                                          // u32
	dst_offset := (externs{ϟa, ϟs, ϟd, ϟl, ϟb}.imageSize(uint32(ϟa.Xoffset), uint32(1), ϟa.Format, ϟa.Type)) + ((dst_stride) * (uint32(ϟa.Yoffset))) // u32
	src_data := func() (result U8ˢ) {
		switch (pbo) == (BufferId(uint32(0))) {
		case true:
			return U8ᵖ(ϟa.Data).Slice(uint64(uint32(0)), uint64(src_size), ϟs)
		case false:
			return U8ᵖ(ctx.Instances.Buffers.Get(pbo).Data.Index(0, ϟs)).Slice(uint64(ϟa.Data.Address), (uint64(ϟa.Data.Address))+(uint64(src_size)), ϟs)
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (pbo) == (BufferId(uint32(0))), ϟa))
			return result
		}
	}() // U8ˢ
	line_bytes := externs{ϟa, ϟs, ϟd, ϟl, ϟb}.imageSize(uint32(ϟa.Width), uint32(1), ϟa.Format, ϟa.Type) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for y := uint32(uint32(0)); y < uint32(ϟa.Height); y++ {
		src := (src_stride) * (y)                  // u32
		dst := ((dst_stride) * (y)) + (dst_offset) // u32
		image.Data.Slice(uint64(dst), uint64((dst)+(line_bytes)), ϟs).Copy(src_data.Slice(uint64(src), uint64((src)+(line_bytes)), ϟs), ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = src, dst
	}
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = minRequiredVersion_1185_major, minRequiredVersion_1185_minor, context, GetContext_1192_result, ctx, tu, image, pbo, url, src_width, src_stride, src_size, dst_stride, dst_offset, src_data, line_bytes
	return nil
}
func (ϟa *GlTexSubImage3D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1193_major := uint32(3) // u32
	minRequiredVersion_1193_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
	case GLenum_GL_TEXTURE_CUBE_MAP_ARRAY:
		minRequiredVersion_1194_major := uint32(3) // u32
		minRequiredVersion_1194_minor := uint32(2) // u32
		_, _ = minRequiredVersion_1194_major, minRequiredVersion_1194_minor
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
		requiresExtension_1197_ext := ExtensionId_GL_OES_texture_half_float // ExtensionId
		_ = requiresExtension_1197_ext
	case GLenum_GL_BYTE, GLenum_GL_FLOAT, GLenum_GL_FLOAT_32_UNSIGNED_INT_24_8_REV, GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_10F_11F_11F_REV, GLenum_GL_UNSIGNED_INT_24_8, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_INT_5_9_9_9_REV, GLenum_GL_UNSIGNED_SHORT, GLenum_GL_UNSIGNED_SHORT_4_4_4_4, GLenum_GL_UNSIGNED_SHORT_5_5_5_1, GLenum_GL_UNSIGNED_SHORT_5_6_5:
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1193_major, minRequiredVersion_1193_minor
	return nil
}
func (ϟa *GlBeginTransformFeedback) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1199_major := uint32(3) // u32
	minRequiredVersion_1199_minor := uint32(0) // u32
	switch ϟa.PrimitiveMode {
	case GLenum_GL_LINES, GLenum_GL_POINTS, GLenum_GL_TRIANGLES:
	default:
		v := ϟa.PrimitiveMode
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1199_major, minRequiredVersion_1199_minor
	return nil
}
func (ϟa *GlBindTransformFeedback) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1201_major := uint32(3) // u32
	minRequiredVersion_1201_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TRANSFORM_FEEDBACK:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1201_major, minRequiredVersion_1201_minor
	return nil
}
func (ϟa *GlDeleteTransformFeedbacks) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1203_major := uint32(3) // u32
	minRequiredVersion_1203_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1203_major, minRequiredVersion_1203_minor
	return nil
}
func (ϟa *GlEndTransformFeedback) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1204_major := uint32(3) // u32
	minRequiredVersion_1204_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1204_major, minRequiredVersion_1204_minor
	return nil
}
func (ϟa *GlGenTransformFeedbacks) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1205_major := uint32(3) // u32
	minRequiredVersion_1205_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1205_major, minRequiredVersion_1205_minor
	return nil
}
func (ϟa *GlGetTransformFeedbackVarying) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1206_major := uint32(3) // u32
	minRequiredVersion_1206_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1206_major, minRequiredVersion_1206_minor
	return nil
}
func (ϟa *GlIsTransformFeedback) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1207_major := uint32(3) // u32
	minRequiredVersion_1207_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_1207_major, minRequiredVersion_1207_minor
	return nil
}
func (ϟa *GlPauseTransformFeedback) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1208_major := uint32(3) // u32
	minRequiredVersion_1208_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1208_major, minRequiredVersion_1208_minor
	return nil
}
func (ϟa *GlResumeTransformFeedback) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1209_major := uint32(3) // u32
	minRequiredVersion_1209_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1209_major, minRequiredVersion_1209_minor
	return nil
}
func (ϟa *GlTransformFeedbackVaryings) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1210_major := uint32(3) // u32
	minRequiredVersion_1210_minor := uint32(0) // u32
	switch ϟa.BufferMode {
	case GLenum_GL_INTERLEAVED_ATTRIBS, GLenum_GL_SEPARATE_ATTRIBS:
	default:
		v := ϟa.BufferMode
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1210_major, minRequiredVersion_1210_minor
	return nil
}
func (ϟa *GlBindVertexArray) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1212_major := uint32(3)   // u32
	minRequiredVersion_1212_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_1213_result := context            // Contextʳ
	ctx := GetContext_1213_result                // Contextʳ
	if !(ctx.Instances.VertexArrays.Contains(ϟa.Array)) {
		ctx.Instances.VertexArrays[ϟa.Array] = &VertexArray{}
	}
	ctx.BoundVertexArray = ϟa.Array
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1212_major, minRequiredVersion_1212_minor, context, GetContext_1213_result, ctx
	return nil
}
func (ϟa *GlBindVertexBuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1214_major := uint32(3) // u32
	minRequiredVersion_1214_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1214_major, minRequiredVersion_1214_minor
	return nil
}
func (ϟa *GlDeleteVertexArrays) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1215_major := uint32(3)                            // u32
	minRequiredVersion_1215_minor := uint32(0)                            // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                          // Contextʳ
	GetContext_1216_result := context                                     // Contextʳ
	ctx := GetContext_1216_result                                         // Contextʳ
	a := ϟa.Arrays.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // VertexArrayIdˢ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.VertexArrays, a.Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_1215_major, minRequiredVersion_1215_minor, context, GetContext_1216_result, ctx, a
	return nil
}
func (ϟa *GlDisableVertexAttribArray) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1217_major := uint32(2)   // u32
	minRequiredVersion_1217_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_1218_result := context            // Contextʳ
	ctx := GetContext_1218_result                // Contextʳ
	ctx.VertexAttributeArrays.Get(ϟa.Location).Enabled = false
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1217_major, minRequiredVersion_1217_minor, context, GetContext_1218_result, ctx
	return nil
}
func (ϟa *GlEnableVertexAttribArray) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1219_major := uint32(2)   // u32
	minRequiredVersion_1219_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_1220_result := context            // Contextʳ
	ctx := GetContext_1220_result                // Contextʳ
	ctx.VertexAttributeArrays.Get(ϟa.Location).Enabled = true
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1219_major, minRequiredVersion_1219_minor, context, GetContext_1220_result, ctx
	return nil
}
func (ϟa *GlGenVertexArrays) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1221_major := uint32(3)                            // u32
	minRequiredVersion_1221_minor := uint32(0)                            // u32
	a := ϟa.Arrays.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // VertexArrayIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                          // Contextʳ
	GetContext_1222_result := context                                     // Contextʳ
	ctx := GetContext_1222_result                                         // Contextʳ
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := VertexArrayId(ϟa.Arrays.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // VertexArrayId
		ctx.Instances.VertexArrays[id] = &VertexArray{}
		a.Index(uint64(i), ϟs).Write(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = minRequiredVersion_1221_major, minRequiredVersion_1221_minor, a, context, GetContext_1222_result, ctx
	return nil
}
func (ϟa *GlGetVertexAttribIiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1223_major := uint32(3) // u32
	minRequiredVersion_1223_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_CURRENT_VERTEX_ATTRIB, GLenum_GL_VERTEX_ATTRIB_ARRAY_BUFFER_BINDING, GLenum_GL_VERTEX_ATTRIB_ARRAY_DIVISOR, GLenum_GL_VERTEX_ATTRIB_ARRAY_ENABLED, GLenum_GL_VERTEX_ATTRIB_ARRAY_INTEGER, GLenum_GL_VERTEX_ATTRIB_ARRAY_NORMALIZED, GLenum_GL_VERTEX_ATTRIB_ARRAY_SIZE, GLenum_GL_VERTEX_ATTRIB_ARRAY_STRIDE, GLenum_GL_VERTEX_ATTRIB_ARRAY_TYPE:
	case GLenum_GL_VERTEX_ATTRIB_BINDING:
		minRequiredVersion_1224_major := uint32(3) // u32
		minRequiredVersion_1224_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1224_major, minRequiredVersion_1224_minor
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1223_major, minRequiredVersion_1223_minor
	return nil
}
func (ϟa *GlGetVertexAttribIuiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1226_major := uint32(3) // u32
	minRequiredVersion_1226_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_CURRENT_VERTEX_ATTRIB, GLenum_GL_VERTEX_ATTRIB_ARRAY_BUFFER_BINDING, GLenum_GL_VERTEX_ATTRIB_ARRAY_DIVISOR, GLenum_GL_VERTEX_ATTRIB_ARRAY_ENABLED, GLenum_GL_VERTEX_ATTRIB_ARRAY_INTEGER, GLenum_GL_VERTEX_ATTRIB_ARRAY_NORMALIZED, GLenum_GL_VERTEX_ATTRIB_ARRAY_SIZE, GLenum_GL_VERTEX_ATTRIB_ARRAY_STRIDE, GLenum_GL_VERTEX_ATTRIB_ARRAY_TYPE:
	case GLenum_GL_VERTEX_ATTRIB_BINDING:
		minRequiredVersion_1227_major := uint32(3) // u32
		minRequiredVersion_1227_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1227_major, minRequiredVersion_1227_minor
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1226_major, minRequiredVersion_1226_minor
	return nil
}
func (ϟa *GlGetVertexAttribPointerv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1229_major := uint32(2) // u32
	minRequiredVersion_1229_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_VERTEX_ATTRIB_ARRAY_POINTER:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1229_major, minRequiredVersion_1229_minor
	return nil
}
func (ϟa *GlGetVertexAttribfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1231_major := uint32(2) // u32
	minRequiredVersion_1231_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_CURRENT_VERTEX_ATTRIB, GLenum_GL_VERTEX_ATTRIB_ARRAY_BUFFER_BINDING, GLenum_GL_VERTEX_ATTRIB_ARRAY_ENABLED, GLenum_GL_VERTEX_ATTRIB_ARRAY_NORMALIZED, GLenum_GL_VERTEX_ATTRIB_ARRAY_SIZE, GLenum_GL_VERTEX_ATTRIB_ARRAY_STRIDE, GLenum_GL_VERTEX_ATTRIB_ARRAY_TYPE:
	case GLenum_GL_VERTEX_ATTRIB_ARRAY_DIVISOR, GLenum_GL_VERTEX_ATTRIB_ARRAY_INTEGER:
		minRequiredVersion_1232_major := uint32(3) // u32
		minRequiredVersion_1232_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1232_major, minRequiredVersion_1232_minor
	case GLenum_GL_VERTEX_ATTRIB_BINDING:
		minRequiredVersion_1233_major := uint32(3) // u32
		minRequiredVersion_1233_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1233_major, minRequiredVersion_1233_minor
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1231_major, minRequiredVersion_1231_minor
	return nil
}
func (ϟa *GlGetVertexAttribiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1235_major := uint32(2) // u32
	minRequiredVersion_1235_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_CURRENT_VERTEX_ATTRIB, GLenum_GL_VERTEX_ATTRIB_ARRAY_BUFFER_BINDING, GLenum_GL_VERTEX_ATTRIB_ARRAY_ENABLED, GLenum_GL_VERTEX_ATTRIB_ARRAY_NORMALIZED, GLenum_GL_VERTEX_ATTRIB_ARRAY_SIZE, GLenum_GL_VERTEX_ATTRIB_ARRAY_STRIDE, GLenum_GL_VERTEX_ATTRIB_ARRAY_TYPE:
	case GLenum_GL_VERTEX_ATTRIB_ARRAY_DIVISOR, GLenum_GL_VERTEX_ATTRIB_ARRAY_INTEGER:
		minRequiredVersion_1236_major := uint32(3) // u32
		minRequiredVersion_1236_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1236_major, minRequiredVersion_1236_minor
	case GLenum_GL_VERTEX_ATTRIB_BINDING:
		minRequiredVersion_1237_major := uint32(3) // u32
		minRequiredVersion_1237_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1237_major, minRequiredVersion_1237_minor
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1235_major, minRequiredVersion_1235_minor
	return nil
}
func (ϟa *GlIsVertexArray) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1239_major := uint32(3) // u32
	minRequiredVersion_1239_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_, _ = minRequiredVersion_1239_major, minRequiredVersion_1239_minor
	return nil
}
func (ϟa *GlVertexAttrib1f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1240_major := uint32(2) // u32
	minRequiredVersion_1240_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1240_major, minRequiredVersion_1240_minor
	return nil
}
func (ϟa *GlVertexAttrib1fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1241_major := uint32(2) // u32
	minRequiredVersion_1241_minor := uint32(0) // u32
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1241_major, minRequiredVersion_1241_minor
	return nil
}
func (ϟa *GlVertexAttrib2f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1242_major := uint32(2) // u32
	minRequiredVersion_1242_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1242_major, minRequiredVersion_1242_minor
	return nil
}
func (ϟa *GlVertexAttrib2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1243_major := uint32(2) // u32
	minRequiredVersion_1243_minor := uint32(0) // u32
	ϟa.Value.Slice(uint64(0), uint64(2), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1243_major, minRequiredVersion_1243_minor
	return nil
}
func (ϟa *GlVertexAttrib3f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1244_major := uint32(2) // u32
	minRequiredVersion_1244_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1244_major, minRequiredVersion_1244_minor
	return nil
}
func (ϟa *GlVertexAttrib3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1245_major := uint32(2) // u32
	minRequiredVersion_1245_minor := uint32(0) // u32
	ϟa.Value.Slice(uint64(0), uint64(3), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1245_major, minRequiredVersion_1245_minor
	return nil
}
func (ϟa *GlVertexAttrib4f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1246_major := uint32(2) // u32
	minRequiredVersion_1246_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1246_major, minRequiredVersion_1246_minor
	return nil
}
func (ϟa *GlVertexAttrib4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1247_major := uint32(2) // u32
	minRequiredVersion_1247_minor := uint32(0) // u32
	ϟa.Value.Slice(uint64(0), uint64(4), ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1247_major, minRequiredVersion_1247_minor
	return nil
}
func (ϟa *GlVertexAttribBinding) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1248_major := uint32(3) // u32
	minRequiredVersion_1248_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1248_major, minRequiredVersion_1248_minor
	return nil
}
func (ϟa *GlVertexAttribDivisor) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1249_major := uint32(3) // u32
	minRequiredVersion_1249_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1249_major, minRequiredVersion_1249_minor
	return nil
}
func (ϟa *GlVertexAttribFormat) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1250_major := uint32(3) // u32
	minRequiredVersion_1250_minor := uint32(1) // u32
	switch ϟa.Type {
	case GLenum_GL_BYTE, GLenum_GL_FIXED, GLenum_GL_FLOAT, GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_INT_2_10_10_10_REV, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_SHORT:
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1250_major, minRequiredVersion_1250_minor
	return nil
}
func (ϟa *GlVertexAttribI4i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1252_major := uint32(3) // u32
	minRequiredVersion_1252_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1252_major, minRequiredVersion_1252_minor
	return nil
}
func (ϟa *GlVertexAttribI4iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1253_major := uint32(3) // u32
	minRequiredVersion_1253_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1253_major, minRequiredVersion_1253_minor
	return nil
}
func (ϟa *GlVertexAttribI4ui) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1254_major := uint32(3) // u32
	minRequiredVersion_1254_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1254_major, minRequiredVersion_1254_minor
	return nil
}
func (ϟa *GlVertexAttribI4uiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1255_major := uint32(3) // u32
	minRequiredVersion_1255_minor := uint32(0) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1255_major, minRequiredVersion_1255_minor
	return nil
}
func (ϟa *GlVertexAttribIFormat) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1256_major := uint32(3) // u32
	minRequiredVersion_1256_minor := uint32(1) // u32
	switch ϟa.Type {
	case GLenum_GL_BYTE, GLenum_GL_FIXED, GLenum_GL_FLOAT, GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_INT_2_10_10_10_REV, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_SHORT:
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1256_major, minRequiredVersion_1256_minor
	return nil
}
func (ϟa *GlVertexAttribIPointer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1258_major := uint32(3) // u32
	minRequiredVersion_1258_minor := uint32(0) // u32
	switch ϟa.Type {
	case GLenum_GL_HALF_FLOAT_OES:
		requiresExtension_1259_ext := ExtensionId_GL_OES_vertex_half_float // ExtensionId
		_ = requiresExtension_1259_ext
	case GLenum_GL_BYTE, GLenum_GL_INT, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_SHORT:
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1258_major, minRequiredVersion_1258_minor
	return nil
}
func (ϟa *GlVertexAttribPointer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1261_major := uint32(2) // u32
	minRequiredVersion_1261_minor := uint32(0) // u32
	switch ϟa.Type {
	case GLenum_GL_BYTE, GLenum_GL_FIXED, GLenum_GL_FLOAT, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_SHORT:
	case GLenum_GL_HALF_FLOAT_OES:
		requiresExtension_1262_ext := ExtensionId_GL_OES_vertex_half_float // ExtensionId
		_ = requiresExtension_1262_ext
	case GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV:
		minRequiredVersion_1263_major := uint32(3) // u32
		minRequiredVersion_1263_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1263_major, minRequiredVersion_1263_minor
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)    // Contextʳ
	GetContext_1265_result := context               // Contextʳ
	ctx := GetContext_1265_result                   // Contextʳ
	a := ctx.VertexAttributeArrays.Get(ϟa.Location) // VertexAttributeArrayʳ
	a.Size = uint32(ϟa.Size)
	a.Type = ϟa.Type
	a.Normalized = ϟa.Normalized
	a.Stride = ϟa.Stride
	a.Pointer = ϟa.Data
	a.Buffer = ctx.BoundBuffers.Get(GLenum_GL_ARRAY_BUFFER)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_1261_major, minRequiredVersion_1261_minor, context, GetContext_1265_result, ctx, a
	return nil
}
func (ϟa *GlVertexBindingDivisor) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1266_major := uint32(3) // u32
	minRequiredVersion_1266_minor := uint32(1) // u32
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1266_major, minRequiredVersion_1266_minor
	return nil
}
func (ϟa *EglInitialize) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	if (ϟa.Major) != (EGLintᵖ{}) {
		ϟa.Major.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Major.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	}
	if (ϟa.Major) != (EGLintᵖ{}) {
		ϟa.Minor.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Minor.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
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
	ctx := &Context{Blending: BlendState{SrcRgbBlendFactor: GLenum_GL_ONE, SrcAlphaBlendFactor: GLenum_GL_ZERO, DstRgbBlendFactor: GLenum_GL_ONE, DstAlphaBlendFactor: GLenum_GL_ZERO, BlendEquationRgb: GLenum_GL_FUNC_ADD, BlendEquationAlpha: GLenum_GL_FUNC_ADD}, Rasterizing: RasterizerState{DepthMask: GLboolean(uint8(1)), DepthTestFunction: GLenum_GL_LESS, DepthNear: GLfloat(float32(0)), DepthFar: GLfloat(float32(1)), ColorMaskRed: GLboolean(uint8(1)), ColorMaskGreen: GLboolean(uint8(1)), ColorMaskBlue: GLboolean(uint8(1)), ColorMaskAlpha: GLboolean(uint8(1)), StencilMask: GLenumːGLuintᵐ{}, FrontFace: GLenum_GL_CCW, CullFace: GLenum_GL_BACK, LineWidth: GLfloat(float32(1)), SampleCoverageValue: GLfloat(float32(1))}, Clearing: ClearState{ClearDepth: GLfloat(float32(1))}, BoundFramebuffers: GLenumːFramebufferIdᵐ{}, BoundRenderbuffers: GLenumːRenderbufferIdᵐ{}, BoundBuffers: GLenumːBufferIdᵐ{}, VertexAttributeArrays: AttributeLocationːVertexAttributeArrayʳᵐ{}, TextureUnits: GLenumːTextureUnitʳᵐ{}, ActiveTextureUnit: GLenum_GL_TEXTURE0, Capabilities: GLenumːboolᵐ{}, GenerateMipmapHint: GLenum_GL_DONT_CARE, PixelStorage: GLenumːGLintᵐ{}, Instances: Objects{Renderbuffers: RenderbufferIdːRenderbufferʳᵐ{}, Textures: TextureIdːTextureʳᵐ{}, Framebuffers: FramebufferIdːFramebufferʳᵐ{}, Buffers: BufferIdːBufferʳᵐ{}, Shaders: ShaderIdːShaderʳᵐ{}, Programs: ProgramIdːProgramʳᵐ{}, VertexArrays: VertexArrayIdːVertexArrayʳᵐ{}, Queries: QueryIdːQueryʳᵐ{}}} // Contextʳ
	ctx.Identifier = identifier
	ctx.Instances.Buffers[BufferId(uint32(0))] = &Buffer{Size: GLsizeiptr(int32(0)), Usage: GLenum_GL_STATIC_DRAW}
	ctx.Instances.Textures[TextureId(uint32(0))] = (&Texture{Texture2D: GLintːImageᵐ{}, Cubemap: GLintːCubemapLevelᵐ{}, MagFilter: GLenum_GL_LINEAR, MinFilter: GLenum_GL_NEAREST_MIPMAP_LINEAR, WrapS: GLenum_GL_REPEAT, WrapT: GLenum_GL_REPEAT, SwizzleR: GLenum_GL_RED, SwizzleG: GLenum_GL_GREEN, SwizzleB: GLenum_GL_BLUE, SwizzleA: GLenum_GL_ALPHA, MaxAnisotropy: float32(1)}).OnCreate(ϟs)
	ctx.Instances.Renderbuffers[RenderbufferId(uint32(0))] = &Renderbuffer{}
	color_id := RenderbufferId(uint32(4294967295))   // RenderbufferId
	depth_id := RenderbufferId(uint32(4294967294))   // RenderbufferId
	stencil_id := RenderbufferId(uint32(4294967293)) // RenderbufferId
	ctx.Instances.Renderbuffers[color_id] = &Renderbuffer{}
	ctx.Instances.Renderbuffers[depth_id] = &Renderbuffer{}
	ctx.Instances.Renderbuffers[stencil_id] = &Renderbuffer{}
	backbuffer := &Framebuffer{Attachments: GLenumːFramebufferAttachmentInfoᵐ{}} // Framebufferʳ
	backbuffer.Attachments[GLenum_GL_COLOR_ATTACHMENT0] = FramebufferAttachmentInfo{Object: uint32(color_id), Type: GLenum_GL_RENDERBUFFER, CubeMapFace: GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X}
	backbuffer.Attachments[GLenum_GL_DEPTH_ATTACHMENT] = FramebufferAttachmentInfo{Object: uint32(depth_id), Type: GLenum_GL_RENDERBUFFER, CubeMapFace: GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X}
	backbuffer.Attachments[GLenum_GL_STENCIL_ATTACHMENT] = FramebufferAttachmentInfo{Object: uint32(stencil_id), Type: GLenum_GL_RENDERBUFFER, CubeMapFace: GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X}
	ctx.Instances.Framebuffers[FramebufferId(uint32(0))] = backbuffer
	ctx.BoundFramebuffers[GLenum_GL_DRAW_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.BoundFramebuffers[GLenum_GL_READ_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.Rasterizing.StencilMask[GLenum_GL_FRONT] = GLuint(uint32(4294967295))
	ctx.Rasterizing.StencilMask[GLenum_GL_BACK] = GLuint(uint32(4294967295))
	ctx.PixelStorage[GLenum_GL_PACK_ALIGNMENT] = GLint(int32(4))
	ctx.PixelStorage[GLenum_GL_UNPACK_ALIGNMENT] = GLint(int32(4))
	for i := int32(int32(0)); i < int32(64); i++ {
		ctx.VertexAttributeArrays[AttributeLocation(i)] = &VertexAttributeArray{Enabled: false, Size: uint32(4), Type: GLenum_GL_FLOAT, Normalized: GLboolean(uint8(0)), Stride: GLsizei(int32(0)), Buffer: BufferId(uint32(0))}
	}
	for i := int32(int32(0)); i < int32(64); i++ {
		ctx.TextureUnits[(GLenum_GL_TEXTURE0)+(GLenum(i))] = &TextureUnit{Bindings: GLenumːTextureIdᵐ{}}
	}
	CreateContext_1267_result := ctx // Contextʳ
	ϟc.EGLContexts[context] = CreateContext_1267_result
	ϟa.Result = context
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_1267_result
	return nil
}
func (ϟa *EglMakeCurrent) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	SetContext_1268_context := ϟc.EGLContexts.Get(ϟa.Context) // Contextʳ
	ϟc.Contexts[ϟc.CurrentThread] = SetContext_1268_context
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = SetContext_1268_context
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
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
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
	ctx := &Context{Blending: BlendState{SrcRgbBlendFactor: GLenum_GL_ONE, SrcAlphaBlendFactor: GLenum_GL_ZERO, DstRgbBlendFactor: GLenum_GL_ONE, DstAlphaBlendFactor: GLenum_GL_ZERO, BlendEquationRgb: GLenum_GL_FUNC_ADD, BlendEquationAlpha: GLenum_GL_FUNC_ADD}, Rasterizing: RasterizerState{DepthMask: GLboolean(uint8(1)), DepthTestFunction: GLenum_GL_LESS, DepthNear: GLfloat(float32(0)), DepthFar: GLfloat(float32(1)), ColorMaskRed: GLboolean(uint8(1)), ColorMaskGreen: GLboolean(uint8(1)), ColorMaskBlue: GLboolean(uint8(1)), ColorMaskAlpha: GLboolean(uint8(1)), StencilMask: GLenumːGLuintᵐ{}, FrontFace: GLenum_GL_CCW, CullFace: GLenum_GL_BACK, LineWidth: GLfloat(float32(1)), SampleCoverageValue: GLfloat(float32(1))}, Clearing: ClearState{ClearDepth: GLfloat(float32(1))}, BoundFramebuffers: GLenumːFramebufferIdᵐ{}, BoundRenderbuffers: GLenumːRenderbufferIdᵐ{}, BoundBuffers: GLenumːBufferIdᵐ{}, VertexAttributeArrays: AttributeLocationːVertexAttributeArrayʳᵐ{}, TextureUnits: GLenumːTextureUnitʳᵐ{}, ActiveTextureUnit: GLenum_GL_TEXTURE0, Capabilities: GLenumːboolᵐ{}, GenerateMipmapHint: GLenum_GL_DONT_CARE, PixelStorage: GLenumːGLintᵐ{}, Instances: Objects{Renderbuffers: RenderbufferIdːRenderbufferʳᵐ{}, Textures: TextureIdːTextureʳᵐ{}, Framebuffers: FramebufferIdːFramebufferʳᵐ{}, Buffers: BufferIdːBufferʳᵐ{}, Shaders: ShaderIdːShaderʳᵐ{}, Programs: ProgramIdːProgramʳᵐ{}, VertexArrays: VertexArrayIdːVertexArrayʳᵐ{}, Queries: QueryIdːQueryʳᵐ{}}} // Contextʳ
	ctx.Identifier = identifier
	ctx.Instances.Buffers[BufferId(uint32(0))] = &Buffer{Size: GLsizeiptr(int32(0)), Usage: GLenum_GL_STATIC_DRAW}
	ctx.Instances.Textures[TextureId(uint32(0))] = (&Texture{Texture2D: GLintːImageᵐ{}, Cubemap: GLintːCubemapLevelᵐ{}, MagFilter: GLenum_GL_LINEAR, MinFilter: GLenum_GL_NEAREST_MIPMAP_LINEAR, WrapS: GLenum_GL_REPEAT, WrapT: GLenum_GL_REPEAT, SwizzleR: GLenum_GL_RED, SwizzleG: GLenum_GL_GREEN, SwizzleB: GLenum_GL_BLUE, SwizzleA: GLenum_GL_ALPHA, MaxAnisotropy: float32(1)}).OnCreate(ϟs)
	ctx.Instances.Renderbuffers[RenderbufferId(uint32(0))] = &Renderbuffer{}
	color_id := RenderbufferId(uint32(4294967295))   // RenderbufferId
	depth_id := RenderbufferId(uint32(4294967294))   // RenderbufferId
	stencil_id := RenderbufferId(uint32(4294967293)) // RenderbufferId
	ctx.Instances.Renderbuffers[color_id] = &Renderbuffer{}
	ctx.Instances.Renderbuffers[depth_id] = &Renderbuffer{}
	ctx.Instances.Renderbuffers[stencil_id] = &Renderbuffer{}
	backbuffer := &Framebuffer{Attachments: GLenumːFramebufferAttachmentInfoᵐ{}} // Framebufferʳ
	backbuffer.Attachments[GLenum_GL_COLOR_ATTACHMENT0] = FramebufferAttachmentInfo{Object: uint32(color_id), Type: GLenum_GL_RENDERBUFFER, CubeMapFace: GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X}
	backbuffer.Attachments[GLenum_GL_DEPTH_ATTACHMENT] = FramebufferAttachmentInfo{Object: uint32(depth_id), Type: GLenum_GL_RENDERBUFFER, CubeMapFace: GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X}
	backbuffer.Attachments[GLenum_GL_STENCIL_ATTACHMENT] = FramebufferAttachmentInfo{Object: uint32(stencil_id), Type: GLenum_GL_RENDERBUFFER, CubeMapFace: GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X}
	ctx.Instances.Framebuffers[FramebufferId(uint32(0))] = backbuffer
	ctx.BoundFramebuffers[GLenum_GL_DRAW_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.BoundFramebuffers[GLenum_GL_READ_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.Rasterizing.StencilMask[GLenum_GL_FRONT] = GLuint(uint32(4294967295))
	ctx.Rasterizing.StencilMask[GLenum_GL_BACK] = GLuint(uint32(4294967295))
	ctx.PixelStorage[GLenum_GL_PACK_ALIGNMENT] = GLint(int32(4))
	ctx.PixelStorage[GLenum_GL_UNPACK_ALIGNMENT] = GLint(int32(4))
	for i := int32(int32(0)); i < int32(64); i++ {
		ctx.VertexAttributeArrays[AttributeLocation(i)] = &VertexAttributeArray{Enabled: false, Size: uint32(4), Type: GLenum_GL_FLOAT, Normalized: GLboolean(uint8(0)), Stride: GLsizei(int32(0)), Buffer: BufferId(uint32(0))}
	}
	for i := int32(int32(0)); i < int32(64); i++ {
		ctx.TextureUnits[(GLenum_GL_TEXTURE0)+(GLenum(i))] = &TextureUnit{Bindings: GLenumːTextureIdᵐ{}}
	}
	CreateContext_1269_result := ctx // Contextʳ
	ϟc.GLXContexts[context] = CreateContext_1269_result
	ϟa.Result = context
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_1269_result
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
	ctx := &Context{Blending: BlendState{SrcRgbBlendFactor: GLenum_GL_ONE, SrcAlphaBlendFactor: GLenum_GL_ZERO, DstRgbBlendFactor: GLenum_GL_ONE, DstAlphaBlendFactor: GLenum_GL_ZERO, BlendEquationRgb: GLenum_GL_FUNC_ADD, BlendEquationAlpha: GLenum_GL_FUNC_ADD}, Rasterizing: RasterizerState{DepthMask: GLboolean(uint8(1)), DepthTestFunction: GLenum_GL_LESS, DepthNear: GLfloat(float32(0)), DepthFar: GLfloat(float32(1)), ColorMaskRed: GLboolean(uint8(1)), ColorMaskGreen: GLboolean(uint8(1)), ColorMaskBlue: GLboolean(uint8(1)), ColorMaskAlpha: GLboolean(uint8(1)), StencilMask: GLenumːGLuintᵐ{}, FrontFace: GLenum_GL_CCW, CullFace: GLenum_GL_BACK, LineWidth: GLfloat(float32(1)), SampleCoverageValue: GLfloat(float32(1))}, Clearing: ClearState{ClearDepth: GLfloat(float32(1))}, BoundFramebuffers: GLenumːFramebufferIdᵐ{}, BoundRenderbuffers: GLenumːRenderbufferIdᵐ{}, BoundBuffers: GLenumːBufferIdᵐ{}, VertexAttributeArrays: AttributeLocationːVertexAttributeArrayʳᵐ{}, TextureUnits: GLenumːTextureUnitʳᵐ{}, ActiveTextureUnit: GLenum_GL_TEXTURE0, Capabilities: GLenumːboolᵐ{}, GenerateMipmapHint: GLenum_GL_DONT_CARE, PixelStorage: GLenumːGLintᵐ{}, Instances: Objects{Renderbuffers: RenderbufferIdːRenderbufferʳᵐ{}, Textures: TextureIdːTextureʳᵐ{}, Framebuffers: FramebufferIdːFramebufferʳᵐ{}, Buffers: BufferIdːBufferʳᵐ{}, Shaders: ShaderIdːShaderʳᵐ{}, Programs: ProgramIdːProgramʳᵐ{}, VertexArrays: VertexArrayIdːVertexArrayʳᵐ{}, Queries: QueryIdːQueryʳᵐ{}}} // Contextʳ
	ctx.Identifier = identifier
	ctx.Instances.Buffers[BufferId(uint32(0))] = &Buffer{Size: GLsizeiptr(int32(0)), Usage: GLenum_GL_STATIC_DRAW}
	ctx.Instances.Textures[TextureId(uint32(0))] = (&Texture{Texture2D: GLintːImageᵐ{}, Cubemap: GLintːCubemapLevelᵐ{}, MagFilter: GLenum_GL_LINEAR, MinFilter: GLenum_GL_NEAREST_MIPMAP_LINEAR, WrapS: GLenum_GL_REPEAT, WrapT: GLenum_GL_REPEAT, SwizzleR: GLenum_GL_RED, SwizzleG: GLenum_GL_GREEN, SwizzleB: GLenum_GL_BLUE, SwizzleA: GLenum_GL_ALPHA, MaxAnisotropy: float32(1)}).OnCreate(ϟs)
	ctx.Instances.Renderbuffers[RenderbufferId(uint32(0))] = &Renderbuffer{}
	color_id := RenderbufferId(uint32(4294967295))   // RenderbufferId
	depth_id := RenderbufferId(uint32(4294967294))   // RenderbufferId
	stencil_id := RenderbufferId(uint32(4294967293)) // RenderbufferId
	ctx.Instances.Renderbuffers[color_id] = &Renderbuffer{}
	ctx.Instances.Renderbuffers[depth_id] = &Renderbuffer{}
	ctx.Instances.Renderbuffers[stencil_id] = &Renderbuffer{}
	backbuffer := &Framebuffer{Attachments: GLenumːFramebufferAttachmentInfoᵐ{}} // Framebufferʳ
	backbuffer.Attachments[GLenum_GL_COLOR_ATTACHMENT0] = FramebufferAttachmentInfo{Object: uint32(color_id), Type: GLenum_GL_RENDERBUFFER, CubeMapFace: GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X}
	backbuffer.Attachments[GLenum_GL_DEPTH_ATTACHMENT] = FramebufferAttachmentInfo{Object: uint32(depth_id), Type: GLenum_GL_RENDERBUFFER, CubeMapFace: GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X}
	backbuffer.Attachments[GLenum_GL_STENCIL_ATTACHMENT] = FramebufferAttachmentInfo{Object: uint32(stencil_id), Type: GLenum_GL_RENDERBUFFER, CubeMapFace: GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X}
	ctx.Instances.Framebuffers[FramebufferId(uint32(0))] = backbuffer
	ctx.BoundFramebuffers[GLenum_GL_DRAW_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.BoundFramebuffers[GLenum_GL_READ_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.Rasterizing.StencilMask[GLenum_GL_FRONT] = GLuint(uint32(4294967295))
	ctx.Rasterizing.StencilMask[GLenum_GL_BACK] = GLuint(uint32(4294967295))
	ctx.PixelStorage[GLenum_GL_PACK_ALIGNMENT] = GLint(int32(4))
	ctx.PixelStorage[GLenum_GL_UNPACK_ALIGNMENT] = GLint(int32(4))
	for i := int32(int32(0)); i < int32(64); i++ {
		ctx.VertexAttributeArrays[AttributeLocation(i)] = &VertexAttributeArray{Enabled: false, Size: uint32(4), Type: GLenum_GL_FLOAT, Normalized: GLboolean(uint8(0)), Stride: GLsizei(int32(0)), Buffer: BufferId(uint32(0))}
	}
	for i := int32(int32(0)); i < int32(64); i++ {
		ctx.TextureUnits[(GLenum_GL_TEXTURE0)+(GLenum(i))] = &TextureUnit{Bindings: GLenumːTextureIdᵐ{}}
	}
	CreateContext_1270_result := ctx // Contextʳ
	ϟc.GLXContexts[context] = CreateContext_1270_result
	ϟa.Result = context
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_1270_result
	return nil
}
func (ϟa *GlXMakeContextCurrent) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	SetContext_1271_context := ϟc.GLXContexts.Get(ϟa.Ctx) // Contextʳ
	ϟc.Contexts[ϟc.CurrentThread] = SetContext_1271_context
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = SetContext_1271_context
	return nil
}
func (ϟa *GlXMakeCurrent) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	SetContext_1272_context := ϟc.GLXContexts.Get(ϟa.Ctx) // Contextʳ
	ϟc.Contexts[ϟc.CurrentThread] = SetContext_1272_context
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = SetContext_1272_context
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
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
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
	ctx := &Context{Blending: BlendState{SrcRgbBlendFactor: GLenum_GL_ONE, SrcAlphaBlendFactor: GLenum_GL_ZERO, DstRgbBlendFactor: GLenum_GL_ONE, DstAlphaBlendFactor: GLenum_GL_ZERO, BlendEquationRgb: GLenum_GL_FUNC_ADD, BlendEquationAlpha: GLenum_GL_FUNC_ADD}, Rasterizing: RasterizerState{DepthMask: GLboolean(uint8(1)), DepthTestFunction: GLenum_GL_LESS, DepthNear: GLfloat(float32(0)), DepthFar: GLfloat(float32(1)), ColorMaskRed: GLboolean(uint8(1)), ColorMaskGreen: GLboolean(uint8(1)), ColorMaskBlue: GLboolean(uint8(1)), ColorMaskAlpha: GLboolean(uint8(1)), StencilMask: GLenumːGLuintᵐ{}, FrontFace: GLenum_GL_CCW, CullFace: GLenum_GL_BACK, LineWidth: GLfloat(float32(1)), SampleCoverageValue: GLfloat(float32(1))}, Clearing: ClearState{ClearDepth: GLfloat(float32(1))}, BoundFramebuffers: GLenumːFramebufferIdᵐ{}, BoundRenderbuffers: GLenumːRenderbufferIdᵐ{}, BoundBuffers: GLenumːBufferIdᵐ{}, VertexAttributeArrays: AttributeLocationːVertexAttributeArrayʳᵐ{}, TextureUnits: GLenumːTextureUnitʳᵐ{}, ActiveTextureUnit: GLenum_GL_TEXTURE0, Capabilities: GLenumːboolᵐ{}, GenerateMipmapHint: GLenum_GL_DONT_CARE, PixelStorage: GLenumːGLintᵐ{}, Instances: Objects{Renderbuffers: RenderbufferIdːRenderbufferʳᵐ{}, Textures: TextureIdːTextureʳᵐ{}, Framebuffers: FramebufferIdːFramebufferʳᵐ{}, Buffers: BufferIdːBufferʳᵐ{}, Shaders: ShaderIdːShaderʳᵐ{}, Programs: ProgramIdːProgramʳᵐ{}, VertexArrays: VertexArrayIdːVertexArrayʳᵐ{}, Queries: QueryIdːQueryʳᵐ{}}} // Contextʳ
	ctx.Identifier = identifier
	ctx.Instances.Buffers[BufferId(uint32(0))] = &Buffer{Size: GLsizeiptr(int32(0)), Usage: GLenum_GL_STATIC_DRAW}
	ctx.Instances.Textures[TextureId(uint32(0))] = (&Texture{Texture2D: GLintːImageᵐ{}, Cubemap: GLintːCubemapLevelᵐ{}, MagFilter: GLenum_GL_LINEAR, MinFilter: GLenum_GL_NEAREST_MIPMAP_LINEAR, WrapS: GLenum_GL_REPEAT, WrapT: GLenum_GL_REPEAT, SwizzleR: GLenum_GL_RED, SwizzleG: GLenum_GL_GREEN, SwizzleB: GLenum_GL_BLUE, SwizzleA: GLenum_GL_ALPHA, MaxAnisotropy: float32(1)}).OnCreate(ϟs)
	ctx.Instances.Renderbuffers[RenderbufferId(uint32(0))] = &Renderbuffer{}
	color_id := RenderbufferId(uint32(4294967295))   // RenderbufferId
	depth_id := RenderbufferId(uint32(4294967294))   // RenderbufferId
	stencil_id := RenderbufferId(uint32(4294967293)) // RenderbufferId
	ctx.Instances.Renderbuffers[color_id] = &Renderbuffer{}
	ctx.Instances.Renderbuffers[depth_id] = &Renderbuffer{}
	ctx.Instances.Renderbuffers[stencil_id] = &Renderbuffer{}
	backbuffer := &Framebuffer{Attachments: GLenumːFramebufferAttachmentInfoᵐ{}} // Framebufferʳ
	backbuffer.Attachments[GLenum_GL_COLOR_ATTACHMENT0] = FramebufferAttachmentInfo{Object: uint32(color_id), Type: GLenum_GL_RENDERBUFFER, CubeMapFace: GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X}
	backbuffer.Attachments[GLenum_GL_DEPTH_ATTACHMENT] = FramebufferAttachmentInfo{Object: uint32(depth_id), Type: GLenum_GL_RENDERBUFFER, CubeMapFace: GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X}
	backbuffer.Attachments[GLenum_GL_STENCIL_ATTACHMENT] = FramebufferAttachmentInfo{Object: uint32(stencil_id), Type: GLenum_GL_RENDERBUFFER, CubeMapFace: GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X}
	ctx.Instances.Framebuffers[FramebufferId(uint32(0))] = backbuffer
	ctx.BoundFramebuffers[GLenum_GL_DRAW_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.BoundFramebuffers[GLenum_GL_READ_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.Rasterizing.StencilMask[GLenum_GL_FRONT] = GLuint(uint32(4294967295))
	ctx.Rasterizing.StencilMask[GLenum_GL_BACK] = GLuint(uint32(4294967295))
	ctx.PixelStorage[GLenum_GL_PACK_ALIGNMENT] = GLint(int32(4))
	ctx.PixelStorage[GLenum_GL_UNPACK_ALIGNMENT] = GLint(int32(4))
	for i := int32(int32(0)); i < int32(64); i++ {
		ctx.VertexAttributeArrays[AttributeLocation(i)] = &VertexAttributeArray{Enabled: false, Size: uint32(4), Type: GLenum_GL_FLOAT, Normalized: GLboolean(uint8(0)), Stride: GLsizei(int32(0)), Buffer: BufferId(uint32(0))}
	}
	for i := int32(int32(0)); i < int32(64); i++ {
		ctx.TextureUnits[(GLenum_GL_TEXTURE0)+(GLenum(i))] = &TextureUnit{Bindings: GLenumːTextureIdᵐ{}}
	}
	CreateContext_1273_result := ctx // Contextʳ
	ϟc.WGLContexts[context] = CreateContext_1273_result
	ϟa.Result = context
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_1273_result
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
	ctx := &Context{Blending: BlendState{SrcRgbBlendFactor: GLenum_GL_ONE, SrcAlphaBlendFactor: GLenum_GL_ZERO, DstRgbBlendFactor: GLenum_GL_ONE, DstAlphaBlendFactor: GLenum_GL_ZERO, BlendEquationRgb: GLenum_GL_FUNC_ADD, BlendEquationAlpha: GLenum_GL_FUNC_ADD}, Rasterizing: RasterizerState{DepthMask: GLboolean(uint8(1)), DepthTestFunction: GLenum_GL_LESS, DepthNear: GLfloat(float32(0)), DepthFar: GLfloat(float32(1)), ColorMaskRed: GLboolean(uint8(1)), ColorMaskGreen: GLboolean(uint8(1)), ColorMaskBlue: GLboolean(uint8(1)), ColorMaskAlpha: GLboolean(uint8(1)), StencilMask: GLenumːGLuintᵐ{}, FrontFace: GLenum_GL_CCW, CullFace: GLenum_GL_BACK, LineWidth: GLfloat(float32(1)), SampleCoverageValue: GLfloat(float32(1))}, Clearing: ClearState{ClearDepth: GLfloat(float32(1))}, BoundFramebuffers: GLenumːFramebufferIdᵐ{}, BoundRenderbuffers: GLenumːRenderbufferIdᵐ{}, BoundBuffers: GLenumːBufferIdᵐ{}, VertexAttributeArrays: AttributeLocationːVertexAttributeArrayʳᵐ{}, TextureUnits: GLenumːTextureUnitʳᵐ{}, ActiveTextureUnit: GLenum_GL_TEXTURE0, Capabilities: GLenumːboolᵐ{}, GenerateMipmapHint: GLenum_GL_DONT_CARE, PixelStorage: GLenumːGLintᵐ{}, Instances: Objects{Renderbuffers: RenderbufferIdːRenderbufferʳᵐ{}, Textures: TextureIdːTextureʳᵐ{}, Framebuffers: FramebufferIdːFramebufferʳᵐ{}, Buffers: BufferIdːBufferʳᵐ{}, Shaders: ShaderIdːShaderʳᵐ{}, Programs: ProgramIdːProgramʳᵐ{}, VertexArrays: VertexArrayIdːVertexArrayʳᵐ{}, Queries: QueryIdːQueryʳᵐ{}}} // Contextʳ
	ctx.Identifier = identifier
	ctx.Instances.Buffers[BufferId(uint32(0))] = &Buffer{Size: GLsizeiptr(int32(0)), Usage: GLenum_GL_STATIC_DRAW}
	ctx.Instances.Textures[TextureId(uint32(0))] = (&Texture{Texture2D: GLintːImageᵐ{}, Cubemap: GLintːCubemapLevelᵐ{}, MagFilter: GLenum_GL_LINEAR, MinFilter: GLenum_GL_NEAREST_MIPMAP_LINEAR, WrapS: GLenum_GL_REPEAT, WrapT: GLenum_GL_REPEAT, SwizzleR: GLenum_GL_RED, SwizzleG: GLenum_GL_GREEN, SwizzleB: GLenum_GL_BLUE, SwizzleA: GLenum_GL_ALPHA, MaxAnisotropy: float32(1)}).OnCreate(ϟs)
	ctx.Instances.Renderbuffers[RenderbufferId(uint32(0))] = &Renderbuffer{}
	color_id := RenderbufferId(uint32(4294967295))   // RenderbufferId
	depth_id := RenderbufferId(uint32(4294967294))   // RenderbufferId
	stencil_id := RenderbufferId(uint32(4294967293)) // RenderbufferId
	ctx.Instances.Renderbuffers[color_id] = &Renderbuffer{}
	ctx.Instances.Renderbuffers[depth_id] = &Renderbuffer{}
	ctx.Instances.Renderbuffers[stencil_id] = &Renderbuffer{}
	backbuffer := &Framebuffer{Attachments: GLenumːFramebufferAttachmentInfoᵐ{}} // Framebufferʳ
	backbuffer.Attachments[GLenum_GL_COLOR_ATTACHMENT0] = FramebufferAttachmentInfo{Object: uint32(color_id), Type: GLenum_GL_RENDERBUFFER, CubeMapFace: GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X}
	backbuffer.Attachments[GLenum_GL_DEPTH_ATTACHMENT] = FramebufferAttachmentInfo{Object: uint32(depth_id), Type: GLenum_GL_RENDERBUFFER, CubeMapFace: GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X}
	backbuffer.Attachments[GLenum_GL_STENCIL_ATTACHMENT] = FramebufferAttachmentInfo{Object: uint32(stencil_id), Type: GLenum_GL_RENDERBUFFER, CubeMapFace: GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X}
	ctx.Instances.Framebuffers[FramebufferId(uint32(0))] = backbuffer
	ctx.BoundFramebuffers[GLenum_GL_DRAW_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.BoundFramebuffers[GLenum_GL_READ_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.Rasterizing.StencilMask[GLenum_GL_FRONT] = GLuint(uint32(4294967295))
	ctx.Rasterizing.StencilMask[GLenum_GL_BACK] = GLuint(uint32(4294967295))
	ctx.PixelStorage[GLenum_GL_PACK_ALIGNMENT] = GLint(int32(4))
	ctx.PixelStorage[GLenum_GL_UNPACK_ALIGNMENT] = GLint(int32(4))
	for i := int32(int32(0)); i < int32(64); i++ {
		ctx.VertexAttributeArrays[AttributeLocation(i)] = &VertexAttributeArray{Enabled: false, Size: uint32(4), Type: GLenum_GL_FLOAT, Normalized: GLboolean(uint8(0)), Stride: GLsizei(int32(0)), Buffer: BufferId(uint32(0))}
	}
	for i := int32(int32(0)); i < int32(64); i++ {
		ctx.TextureUnits[(GLenum_GL_TEXTURE0)+(GLenum(i))] = &TextureUnit{Bindings: GLenumːTextureIdᵐ{}}
	}
	CreateContext_1274_result := ctx // Contextʳ
	ϟc.WGLContexts[context] = CreateContext_1274_result
	ϟa.Result = context
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_1274_result
	return nil
}
func (ϟa *WglMakeCurrent) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	SetContext_1275_context := ϟc.WGLContexts.Get(ϟa.Hglrc) // Contextʳ
	ϟc.Contexts[ϟc.CurrentThread] = SetContext_1275_context
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = SetContext_1275_context
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
	context := CGLContextObj(ϟa.Ctx.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil)) // CGLContextObj
	identifier := ϟc.NextContextID                                                                                  // ContextID
	ϟc.NextContextID = (ϟc.NextContextID) + (ContextID(uint32(1)))
	ctx := &Context{Blending: BlendState{SrcRgbBlendFactor: GLenum_GL_ONE, SrcAlphaBlendFactor: GLenum_GL_ZERO, DstRgbBlendFactor: GLenum_GL_ONE, DstAlphaBlendFactor: GLenum_GL_ZERO, BlendEquationRgb: GLenum_GL_FUNC_ADD, BlendEquationAlpha: GLenum_GL_FUNC_ADD}, Rasterizing: RasterizerState{DepthMask: GLboolean(uint8(1)), DepthTestFunction: GLenum_GL_LESS, DepthNear: GLfloat(float32(0)), DepthFar: GLfloat(float32(1)), ColorMaskRed: GLboolean(uint8(1)), ColorMaskGreen: GLboolean(uint8(1)), ColorMaskBlue: GLboolean(uint8(1)), ColorMaskAlpha: GLboolean(uint8(1)), StencilMask: GLenumːGLuintᵐ{}, FrontFace: GLenum_GL_CCW, CullFace: GLenum_GL_BACK, LineWidth: GLfloat(float32(1)), SampleCoverageValue: GLfloat(float32(1))}, Clearing: ClearState{ClearDepth: GLfloat(float32(1))}, BoundFramebuffers: GLenumːFramebufferIdᵐ{}, BoundRenderbuffers: GLenumːRenderbufferIdᵐ{}, BoundBuffers: GLenumːBufferIdᵐ{}, VertexAttributeArrays: AttributeLocationːVertexAttributeArrayʳᵐ{}, TextureUnits: GLenumːTextureUnitʳᵐ{}, ActiveTextureUnit: GLenum_GL_TEXTURE0, Capabilities: GLenumːboolᵐ{}, GenerateMipmapHint: GLenum_GL_DONT_CARE, PixelStorage: GLenumːGLintᵐ{}, Instances: Objects{Renderbuffers: RenderbufferIdːRenderbufferʳᵐ{}, Textures: TextureIdːTextureʳᵐ{}, Framebuffers: FramebufferIdːFramebufferʳᵐ{}, Buffers: BufferIdːBufferʳᵐ{}, Shaders: ShaderIdːShaderʳᵐ{}, Programs: ProgramIdːProgramʳᵐ{}, VertexArrays: VertexArrayIdːVertexArrayʳᵐ{}, Queries: QueryIdːQueryʳᵐ{}}} // Contextʳ
	ctx.Identifier = identifier
	ctx.Instances.Buffers[BufferId(uint32(0))] = &Buffer{Size: GLsizeiptr(int32(0)), Usage: GLenum_GL_STATIC_DRAW}
	ctx.Instances.Textures[TextureId(uint32(0))] = (&Texture{Texture2D: GLintːImageᵐ{}, Cubemap: GLintːCubemapLevelᵐ{}, MagFilter: GLenum_GL_LINEAR, MinFilter: GLenum_GL_NEAREST_MIPMAP_LINEAR, WrapS: GLenum_GL_REPEAT, WrapT: GLenum_GL_REPEAT, SwizzleR: GLenum_GL_RED, SwizzleG: GLenum_GL_GREEN, SwizzleB: GLenum_GL_BLUE, SwizzleA: GLenum_GL_ALPHA, MaxAnisotropy: float32(1)}).OnCreate(ϟs)
	ctx.Instances.Renderbuffers[RenderbufferId(uint32(0))] = &Renderbuffer{}
	color_id := RenderbufferId(uint32(4294967295))   // RenderbufferId
	depth_id := RenderbufferId(uint32(4294967294))   // RenderbufferId
	stencil_id := RenderbufferId(uint32(4294967293)) // RenderbufferId
	ctx.Instances.Renderbuffers[color_id] = &Renderbuffer{}
	ctx.Instances.Renderbuffers[depth_id] = &Renderbuffer{}
	ctx.Instances.Renderbuffers[stencil_id] = &Renderbuffer{}
	backbuffer := &Framebuffer{Attachments: GLenumːFramebufferAttachmentInfoᵐ{}} // Framebufferʳ
	backbuffer.Attachments[GLenum_GL_COLOR_ATTACHMENT0] = FramebufferAttachmentInfo{Object: uint32(color_id), Type: GLenum_GL_RENDERBUFFER, CubeMapFace: GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X}
	backbuffer.Attachments[GLenum_GL_DEPTH_ATTACHMENT] = FramebufferAttachmentInfo{Object: uint32(depth_id), Type: GLenum_GL_RENDERBUFFER, CubeMapFace: GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X}
	backbuffer.Attachments[GLenum_GL_STENCIL_ATTACHMENT] = FramebufferAttachmentInfo{Object: uint32(stencil_id), Type: GLenum_GL_RENDERBUFFER, CubeMapFace: GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X}
	ctx.Instances.Framebuffers[FramebufferId(uint32(0))] = backbuffer
	ctx.BoundFramebuffers[GLenum_GL_DRAW_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.BoundFramebuffers[GLenum_GL_READ_FRAMEBUFFER] = FramebufferId(uint32(0))
	ctx.Rasterizing.StencilMask[GLenum_GL_FRONT] = GLuint(uint32(4294967295))
	ctx.Rasterizing.StencilMask[GLenum_GL_BACK] = GLuint(uint32(4294967295))
	ctx.PixelStorage[GLenum_GL_PACK_ALIGNMENT] = GLint(int32(4))
	ctx.PixelStorage[GLenum_GL_UNPACK_ALIGNMENT] = GLint(int32(4))
	for i := int32(int32(0)); i < int32(64); i++ {
		ctx.VertexAttributeArrays[AttributeLocation(i)] = &VertexAttributeArray{Enabled: false, Size: uint32(4), Type: GLenum_GL_FLOAT, Normalized: GLboolean(uint8(0)), Stride: GLsizei(int32(0)), Buffer: BufferId(uint32(0))}
	}
	for i := int32(int32(0)); i < int32(64); i++ {
		ctx.TextureUnits[(GLenum_GL_TEXTURE0)+(GLenum(i))] = &TextureUnit{Bindings: GLenumːTextureIdᵐ{}}
	}
	CreateContext_1276_result := ctx // Contextʳ
	ϟc.CGLContexts[context] = CreateContext_1276_result
	ϟa.Ctx.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(context, ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Result = ϟa.Result
	_, _, _, _, _, _, _, _ = context, identifier, ctx, color_id, depth_id, stencil_id, backbuffer, CreateContext_1276_result
	return nil
}
func (ϟa *CGLSetCurrentContext) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	SetContext_1277_context := ϟc.CGLContexts.Get(ϟa.Ctx) // Contextʳ
	ϟc.Contexts[ϟc.CurrentThread] = SetContext_1277_context
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Result = ϟa.Result
	_ = SetContext_1277_context
	return nil
}
func (ϟa *CGLGetSurface) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Cid.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Cid.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Wid.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Wid.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Sid.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Sid.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
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
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}
func (ϟa *GlGetQueryObjectui64v) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Write(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).Read(ϟa, ϟs, ϟd, ϟl, nil), ϟa, ϟs, ϟd, ϟl, ϟb)
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
func (ϟa *ContextInfo) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	ϟc, ϟb := getState(ϟs), (*builder.Builder)(nil)
	_, _ = ϟc, ϟb
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_1278_result := context            // Contextʳ
	ctx := GetContext_1278_result                // Contextʳ
	ctx.Info.Name = ϟa.Name
	ctx.Info.Vendor = ϟa.Vendor
	ctx.Info.Extensions = ϟa.Extensions
	ctx.Info.Version = ϟa.Version
	ctx.Info.PreserveBuffersOnSwap = ϟa.PreserveBuffersOnSwap
	backbuffer := ctx.Instances.Framebuffers.Get(FramebufferId(uint32(0)))                        // Framebufferʳ
	color_id := RenderbufferId(backbuffer.Attachments.Get(GLenum_GL_COLOR_ATTACHMENT0).Object)    // RenderbufferId
	color_buffer := ctx.Instances.Renderbuffers.Get(color_id)                                     // Renderbufferʳ
	depth_id := RenderbufferId(backbuffer.Attachments.Get(GLenum_GL_DEPTH_ATTACHMENT).Object)     // RenderbufferId
	depth_buffer := ctx.Instances.Renderbuffers.Get(depth_id)                                     // Renderbufferʳ
	stencil_id := RenderbufferId(backbuffer.Attachments.Get(GLenum_GL_STENCIL_ATTACHMENT).Object) // RenderbufferId
	stencil_buffer := ctx.Instances.Renderbuffers.Get(stencil_id)                                 // Renderbufferʳ
	color_buffer.Width = ϟa.BackbufferWidth
	color_buffer.Height = ϟa.BackbufferHeight
	color_buffer.Format = ϟa.BackbufferColorFmt
	depth_buffer.Width = ϟa.BackbufferWidth
	depth_buffer.Height = ϟa.BackbufferHeight
	depth_buffer.Format = ϟa.BackbufferDepthFmt
	stencil_buffer.Width = ϟa.BackbufferWidth
	stencil_buffer.Height = ϟa.BackbufferHeight
	stencil_buffer.Format = ϟa.BackbufferStencilFmt
	if ϟa.ResetViewportScissor {
		ctx.Rasterizing.Scissor.Width = ϟa.BackbufferWidth
		ctx.Rasterizing.Scissor.Height = ϟa.BackbufferHeight
		ctx.Rasterizing.Viewport.Width = ϟa.BackbufferWidth
		ctx.Rasterizing.Viewport.Height = ϟa.BackbufferHeight
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _, _, _ = context, GetContext_1278_result, ctx, backbuffer, color_id, color_buffer, depth_id, depth_buffer, stencil_id, stencil_buffer
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
